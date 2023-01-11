package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	deleteWorkerResponseData = `{
    "result": null,
    "success": true,
    "errors": [],
    "messages": []
}`
	uploadWorkerResponseData = `{
    "result": {
        "script": "addEventListener('fetch', event => {\n  event.passThroughOnException()\n  event.respondWith(handleRequest(event.request))\n})\n\nasync function handleRequest(request) {\n  return fetch(request)\n}",
        "etag": "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
        "size": 191,
        "modified_on": "2018-06-09T15:17:01.989141Z"
    },
    "success": true,
    "errors": [],
    "messages": []
}`

	uploadWorkerModuleResponseData = `{
    "result": {
        "script": "export default {\n  async fetch(request, env, event) {\n    event.passThroughOnException()\n    return fetch(request)\n  }\n}",
        "etag": "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
        "size": 191,
        "modified_on": "2018-06-09T15:17:01.989141Z"
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	updateWorkerRouteResponse = `{
    "result": {
        "id": "e7a57d8746e74ae49c25994dadb421b1",
        "pattern": "app3.example.com/*",
        "enabled": true
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	updateWorkerRouteEntResponse = `{
    "result": {
        "id": "e7a57d8746e74ae49c25994dadb421b1",
        "pattern": "app3.example.com/*",
        "script": "test_script_1"
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	createWorkerRouteResponse = `{
    "result": {
        "id": "e7a57d8746e74ae49c25994dadb421b1"
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	listRouteResponseData = `{
    "result": [
        {
            "id": "e7a57d8746e74ae49c25994dadb421b1",
            "pattern": "app1.example.com/*",
            "enabled": true
        },
        {
            "id": "f8b68e9857f85bf59c25994dadb421b1",
            "pattern": "app2.example.com/*",
            "enabled": false
        }
    ],
    "success": true,
    "errors": [],
    "messages": []
}`
	listWorkerRouteResponse = `{
    "result": [
        {
            "id": "e7a57d8746e74ae49c25994dadb421b1",
            "pattern": "app1.example.com/*",
            "script": "test_script_1"
        },
        {
            "id": "f8b68e9857f85bf59c25994dadb421b1",
            "pattern": "app2.example.com/*",
            "script": "test_script_2"
        },
        {
            "id": "2b5bf4240cd34c77852fac70b1bf745a",
            "pattern": "app3.example.com/*",
			"script": "test_script_3"
        }
    ],
    "success": true,
    "errors": [],
    "messages": []
}`
	getRouteResponseData = `{
    "result": {
       "id": "e7a57d8746e74ae49c25994dadb421b1",
       "pattern": "app1.example.com/*",
       "script": "script-name"
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	listBindingsResponseData = `{
		"result": [
			{
				"name": "MY_KV",
				"namespace_id": "89f5f8fd93f94cb98473f6f421aa3b65",
				"type": "kv_namespace"
			},
			{
				"name": "MY_WASM",
				"type": "wasm_module"
			},
			{
				"name": "MY_PLAIN_TEXT",
				"type": "plain_text",
				"text": "text"
			},
			{
				"name": "MY_SECRET_TEXT",
				"type": "secret_text"
			},
			{
				"name": "MY_SERVICE_BINDING",
				"type": "service",
				"service": "MY_SERVICE",
				"environment": "MY_ENVIRONMENT"
			},
			{
				"name": "MY_NEW_BINDING",
				"type": "some_imaginary_new_binding_type"
			},
			{
				"name": "MY_BUCKET",
				"type": "r2_bucket",
				"bucket_name": "bucket"
			},
			{
				"name": "MY_DATASET",
				"type": "analytics_engine",
				"dataset": "my_dataset"
			}
		],
		"success": true,
		"errors": [],
		"messages": []
	}`
	listWorkersResponseData = `{
  "result": [
    {
      "id": "bar",
      "created_on": "2018-04-22T17:10:48.938097Z",
      "modified_on": "2018-04-22T17:10:48.938097Z",
      "etag": "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a"
    },
    {
      "id": "baz",
      "created_on": "2018-04-22T17:10:48.938097Z",
      "modified_on": "2018-04-22T17:10:48.938097Z",
      "etag": "380dg51e97e80b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43088b"
    }
  ],
  "success": true,
  "errors": [],
  "messages": []
}`
	workerScript = `addEventListener('fetch', event => {
  event.passThroughOnException()
  event.respondWith(handleRequest(event.request))
})

async function handleRequest(request) {
  return fetch(request)
}`
	workerModuleScript = `export default {
  async fetch(request, env, event) {
    event.passThroughOnException()
    return fetch(request)
  }
}`
	workerModuleScriptDownloadResponse = `
--workermodulescriptdownload
Content-Disposition: form-data; name="worker.js"

export default {
  async fetch(request, env, event) {
    event.passThroughOnException()
    return fetch(request)
  }
}
--workermodulescriptdownload--
`
)

var (
	successResponse               = Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}}
	deleteWorkerRouteResponseData = createWorkerRouteResponse
	attachWorkerToDomainResponse  = fmt.Sprintf(`{
    "result": {
        "id": "e7a57d8746e74ae49c25994dadb421b1",
	"zone_id": "%s",
	"service":"test_script_1",
	"hostname":"api4.example.com",
	"environment":"production"
    },
    "success": true,
    "errors": [],
    "messages": []
}`, testZoneID)
)

func getFormValue(r *http.Request, key string) ([]byte, error) {
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		return nil, err
	}

	// In Go 1.10 there was a bug where field values with a content-type
	// but without a filename would end up in Form.File but in versions
	// before and after 1.10 they would be in form.Value. Here we check
	// both in order to handle both scenarios
	// https://golang.org/doc/go1.11#mime/multipart

	// pre/post v1.10
	if values, ok := r.MultipartForm.Value[key]; ok {
		return []byte(values[0]), nil
	}

	// v1.10
	if fileHeaders, ok := r.MultipartForm.File[key]; ok {
		file, err := fileHeaders[0].Open()
		if err != nil {
			return nil, err
		}
		return io.ReadAll(file)
	}

	return nil, fmt.Errorf("no value found for key %v", key)
}

func getFileDetails(r *http.Request, key string) (*multipart.FileHeader, error) {
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		return nil, err
	}

	fileHeaders := r.MultipartForm.File[key]

	if len(fileHeaders) > 0 {
		return fileHeaders[0], nil
	}

	return nil, fmt.Errorf("no value found for key %v", key)
}

type multipartUpload struct {
	Script             string
	BindingMeta        map[string]workerBindingMeta
	Logpush            *bool
	CompatibilityDate  string
	CompatibilityFlags []string
}

func parseMultipartUpload(r *http.Request) (multipartUpload, error) {
	// Parse the metadata
	mdBytes, err := getFormValue(r, "metadata")
	if err != nil {
		return multipartUpload{}, err
	}

	var metadata struct {
		BodyPart           string              `json:"body_part,omitempty"`
		MainModule         string              `json:"main_module,omitempty"`
		Bindings           []workerBindingMeta `json:"bindings"`
		Logpush            *bool               `json:"logpush,omitempty"`
		CompatibilityDate  string              `json:"compatibility_date,omitempty"`
		CompatibilityFlags []string            `json:"compatibility_flags,omitempty"`
	}
	err = json.Unmarshal(mdBytes, &metadata)
	if err != nil {
		return multipartUpload{}, err
	}

	// Get the script
	script, err := getFormValue(r, metadata.BodyPart)
	if err != nil {
		script, err = getFormValue(r, metadata.MainModule)

		if err != nil {
			return multipartUpload{}, err
		}
	}

	// Since bindings are specified in the Go API as a map but are uploaded as a
	// JSON array, the ordering of uploaded bindings is non-deterministic. To make
	// it easier to compare for equality without running into ordering issues, we
	// convert it back to a map
	bindingMeta := make(map[string]workerBindingMeta)
	for _, binding := range metadata.Bindings {
		bindingMeta[binding["name"].(string)] = binding
	}

	return multipartUpload{
		Script:             string(script),
		BindingMeta:        bindingMeta,
		Logpush:            metadata.Logpush,
		CompatibilityDate:  metadata.CompatibilityDate,
		CompatibilityFlags: metadata.CompatibilityFlags,
	}, nil
}

func TestDeleteWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprint(w, deleteWorkerResponseData)
	})

	err := client.DeleteWorker(context.Background(), AccountIdentifier(testAccountID), DeleteWorkerParams{ScriptName: "bar"})
	assert.NoError(t, err)
}

func TestGetWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/foo", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprint(w, workerScript)
	})
	res, err := client.GetWorker(context.Background(), AccountIdentifier(testAccountID), "foo")
	want := WorkerScriptResponse{
		successResponse,
		false,
		WorkerScript{
			Script: workerScript,
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Script, res.Script)
	}
}

func TestGetWorker_Module(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/foo", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "multipart/form-data; boundary=workermodulescriptdownload")
		fmt.Fprint(w, workerModuleScriptDownloadResponse)
	})

	res, err := client.GetWorker(context.Background(), AccountIdentifier(testAccountID), "foo")
	want := WorkerScriptResponse{
		successResponse,
		true,
		WorkerScript{
			Script: workerModuleScript,
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want.Script, res.Script)
	}
}

func TestListWorkers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, listWorkersResponseData)
	})

	res, _, err := client.ListWorkers(context.Background(), AccountIdentifier(testAccountID), ListWorkersParams{})
	sampleDate, _ := time.Parse(time.RFC3339Nano, "2018-04-22T17:10:48.938097Z")
	want := []WorkerMetaData{
		{
			ID:         "bar",
			ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
			CreatedOn:  sampleDate,
			ModifiedOn: sampleDate,
		},
		{
			ID:         "baz",
			ETAG:       "380dg51e97e80b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43088b",
			CreatedOn:  sampleDate,
			ModifiedOn: sampleDate,
		},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res.WorkerList)
	}
}

func TestUploadWorker_Basic(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/foo", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		contentTypeHeader := r.Header.Get("content-type")
		assert.Equal(t, "application/javascript", contentTypeHeader, "Expected content-type request header to be 'application/javascript', got %s", contentTypeHeader)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	})
	res, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{ScriptName: "foo", Script: workerScript})
	formattedTime, _ := time.Parse(time.RFC3339Nano, "2018-06-09T15:17:01.989141Z")
	want := WorkerScriptResponse{
		successResponse,
		false,
		WorkerScript{
			Script: workerScript,
			WorkerMetaData: WorkerMetaData{
				ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
				Size:       191,
				ModifiedOn: formattedTime,
			},
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestUploadWorker_Module(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/foo", func(w http.ResponseWriter, r *http.Request) {
		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		assert.Equal(t, workerModuleScript, mpUpload.Script)

		workerFileDetails, err := getFileDetails(r, "worker.mjs")
		if !assert.NoError(t, err) {
			assert.FailNow(t, "worker file not found in multipart form body")
		}
		contentTypeHeader := workerFileDetails.Header.Get("content-type")
		expectedContentType := "application/javascript+module"
		assert.Equal(t, expectedContentType, contentTypeHeader, "Expected content-type request header to be %s, got %s", expectedContentType, contentTypeHeader)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerModuleResponseData)
	})
	res, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{ScriptName: "foo", Script: workerModuleScript, Module: true})
	formattedTime, _ := time.Parse(time.RFC3339Nano, "2018-06-09T15:17:01.989141Z")
	want := WorkerScriptResponse{
		successResponse,
		false,
		WorkerScript{
			Script: workerModuleScript,
			WorkerMetaData: WorkerMetaData{
				ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
				Size:       191,
				ModifiedOn: formattedTime,
			},
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestUploadWorker_WithDurableObjectBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name":        "b1",
				"type":        "durable_object_namespace",
				"class_name":  "TheClass",
				"script_name": "the_script",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerDurableObjectBinding{
				ClassName:  "TheClass",
				ScriptName: "the_script",
			},
		},
	})

	assert.NoError(t, err)
}

func TestUploadWorker_WithInheritBinding(t *testing.T) {
	setup()
	defer teardown()

	// Setup route handler for both single-script and multi-script
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name": "b1",
				"type": "inherit",
			},
			"b2": {
				"name":     "b2",
				"type":     "inherit",
				"old_name": "old_binding_name",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	formattedTime, _ := time.Parse(time.RFC3339Nano, "2018-06-09T15:17:01.989141Z")
	want := WorkerScriptResponse{
		successResponse,
		false,
		WorkerScript{
			Script: workerScript,
			WorkerMetaData: WorkerMetaData{
				ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
				Size:       191,
				ModifiedOn: formattedTime,
			},
		}}

	res, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerInheritBinding{},
			"b2": WorkerInheritBinding{
				OldName: "old_binding_name",
			},
		}})
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestUploadWorker_WithKVBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name":         "b1",
				"type":         "kv_namespace",
				"namespace_id": "test-namespace",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerKvNamespaceBinding{
				NamespaceID: "test-namespace",
			},
		}})
	assert.NoError(t, err)
}

func TestUploadWorker_WithWasmBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		partName := mpUpload.BindingMeta["b1"]["part"].(string)
		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name": "b1",
				"type": "wasm_module",
				"part": partName,
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		wasmContent, err := getFormValue(r, partName)
		assert.NoError(t, err)
		assert.Equal(t, []byte("fake-wasm"), wasmContent)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerWebAssemblyBinding{
				Module: strings.NewReader("fake-wasm"),
			},
		},
	})

	assert.NoError(t, err)
}

func TestUploadWorker_WithPlainTextBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name": "b1",
				"type": "plain_text",
				"text": "plain text value",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerPlainTextBinding{
				Text: "plain text value",
			},
		},
	})

	assert.NoError(t, err)
}

func TestUploadWorker_ModuleWithPlainTextBinding(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name": "b1",
				"type": "plain_text",
				"text": "plain text value",
			},
		}
		assert.Equal(t, workerModuleScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		workerFileDetails, err := getFileDetails(r, "worker.mjs")
		if !assert.NoError(t, err) {
			assert.FailNow(t, "worker file not found in multipart form body")
		}
		contentDispositonHeader := workerFileDetails.Header.Get("content-disposition")
		expectedContentDisposition := fmt.Sprintf(`form-data; name="%s"; filename="%[1]s"`, "worker.mjs")
		assert.Equal(t, expectedContentDisposition, contentDispositonHeader, "Expected content-disposition request header to be %s, got %s", expectedContentDisposition, contentDispositonHeader)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerModuleResponseData)
	})

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerModuleScript,
		Module:     true,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerPlainTextBinding{
				Text: "plain text value",
			},
		},
	})

	assert.NoError(t, err)
}

func TestUploadWorker_WithSecretTextBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name": "b1",
				"type": "secret_text",
				"text": "secret text value",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerSecretTextBinding{
				Text: "secret text value",
			},
		},
	})
	assert.NoError(t, err)
}

func TestUploadWorker_WithServiceBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name":    "b1",
				"type":    "service",
				"service": "the_service",
			},
			"b2": {
				"name":        "b2",
				"type":        "service",
				"service":     "the_service",
				"environment": "the_environment",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerServiceBinding{
				Service: "the_service",
			},
			"b2": WorkerServiceBinding{
				Service:     "the_service",
				Environment: StringPtr("the_environment"),
			},
		},
	})
	assert.NoError(t, err)
}

func TestUploadWorker_WithLogpush(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/foo", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expected := true
		assert.Equal(t, &expected, mpUpload.Logpush)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	})
	res, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{ScriptName: "foo", Script: workerScript, Logpush: BoolPtr(true)})
	formattedTime, _ := time.Parse(time.RFC3339Nano, "2018-06-09T15:17:01.989141Z")
	want := WorkerScriptResponse{
		successResponse,
		false,
		WorkerScript{
			Script: workerScript,
			WorkerMetaData: WorkerMetaData{
				ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
				Size:       191,
				ModifiedOn: formattedTime,
			},
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestUploadWorker_WithCompatibilityFlags(t *testing.T) {
	setup()
	defer teardown()

	compatibilityDate := time.Now().Format("2006-01-02")
	compatibilityFlags := []string{"formdata_parser_supports_files"}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, compatibilityDate, mpUpload.CompatibilityDate)
		assert.Equal(t, compatibilityFlags, mpUpload.CompatibilityFlags)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName:         "bar",
		Script:             workerScript,
		CompatibilityDate:  compatibilityDate,
		CompatibilityFlags: compatibilityFlags,
	})
	assert.NoError(t, err)
}

func TestUploadWorker_WithQueueBinding(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		mpUpload, err := parseMultipartUpload(r)
		assert.NoError(t, err)

		expectedBindings := map[string]workerBindingMeta{
			"b1": {
				"name":       "b1",
				"type":       "queue",
				"queue_name": "test-queue",
			},
		}
		assert.Equal(t, workerScript, mpUpload.Script)
		assert.Equal(t, expectedBindings, mpUpload.BindingMeta)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, uploadWorkerResponseData)
	}
	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/bar", handler)

	_, err := client.UploadWorker(context.Background(), AccountIdentifier(testAccountID), CreateWorkerParams{
		ScriptName: "bar",
		Script:     workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerQueueBinding{
				Binding: "b1",
				Queue:   "test-queue",
			},
		}})
	assert.NoError(t, err)
}
