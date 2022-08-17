package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	listRouteEntResponseData = `{
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
            "pattern": "app3.example.com/*"
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
		return ioutil.ReadAll(file)
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
	Script      string
	BindingMeta map[string]workerBindingMeta
}

func parseMultipartUpload(r *http.Request) (multipartUpload, error) {
	// Parse the metadata
	mdBytes, err := getFormValue(r, "metadata")
	if err != nil {
		return multipartUpload{}, err
	}

	var metadata struct {
		BodyPart   string              `json:"body_part,omitempty"`
		MainModule string              `json:"main_module,omitempty"`
		Bindings   []workerBindingMeta `json:"bindings"`
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
		Script:      string(script),
		BindingMeta: bindingMeta,
	}, nil
}

func TestWorkers_DeleteWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, deleteWorkerResponseData) //nolint
	})
	res, err := client.DeleteWorker(context.Background(), &WorkerRequestParams{ZoneID: "foo"})
	want := WorkerScriptResponse{
		Response: successResponse,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}

func TestWorkers_DeleteWorkerWithName(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, deleteWorkerResponseData) //nolint
	})
	res, err := client.DeleteWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"})
	want := WorkerScriptResponse{
		Response: successResponse,
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}

func TestWorkers_DeleteWorkerWithNameErrorsWithoutAccountId(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.DeleteWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"})
	assert.Error(t, err)
}

func TestWorkers_DownloadWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, workerScript) //nolint
	})
	res, err := client.DownloadWorker(context.Background(), &WorkerRequestParams{ZoneID: "foo"})
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

func TestWorkers_DownloadWorkerWithName(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, workerScript) //nolint
	})
	res, err := client.DownloadWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"})
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

func TestWorkers_DownloadWorkerWithNameErrorsWithoutAccountId(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.DownloadWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"})
	assert.Error(t, err)
}

func TestWorkers_DownloadWorkerModule(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "multipart/form-data; boundary=workermodulescriptdownload")
		fmt.Fprintf(w, workerModuleScriptDownloadResponse) //nolint
	})
	res, err := client.DownloadWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"})
	want := WorkerScriptResponse{
		successResponse,
		true,
		WorkerScript{
			Script: workerModuleScript,
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Script, res.Script)
	}
}

func TestWorkers_ListWorkerScripts(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, listWorkersResponseData) //nolint
	})

	res, err := client.ListWorkerScripts(context.Background())
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

func TestWorkers_UploadWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		contentTypeHeader := r.Header.Get("content-type")
		assert.Equal(t, "application/javascript", contentTypeHeader, "Expected content-type request header to be 'application/javascript', got %s", contentTypeHeader)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	})
	res, err := client.UploadWorker(context.Background(), &WorkerRequestParams{ZoneID: "foo"}, &WorkerScriptParams{Script: workerScript})
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

func TestWorkers_UploadWorkerAsModule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, uploadWorkerModuleResponseData) //nolint
	})
	res, err := client.UploadWorker(context.Background(), &WorkerRequestParams{ZoneID: "foo"}, &WorkerScriptParams{Script: workerModuleScript, Module: true})
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

func TestWorkers_UploadWorkerWithName(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		contentTypeHeader := r.Header.Get("content-type")
		assert.Equal(t, "application/javascript", contentTypeHeader, "Expected content-type request header to be 'application/javascript', got %s", contentTypeHeader)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	})
	res, err := client.UploadWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &WorkerScriptParams{Script: workerScript})
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

func TestWorkers_UploadWorkerSingleScriptWithAccount(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		contentTypeHeader := r.Header.Get("content-type")
		assert.Equal(t, "application/javascript", contentTypeHeader, "Expected content-type request header to be 'application/javascript', got %s", contentTypeHeader)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	})
	res, err := client.UploadWorker(context.Background(), &WorkerRequestParams{ZoneID: "foo"}, &WorkerScriptParams{Script: workerScript})
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

func TestWorkers_UploadWorkerWithNameErrorsWithoutAccountId(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UploadWorker(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &WorkerScriptParams{Script: workerScript})
	assert.Error(t, err)
}

func TestWorkers_UploadWorkerWithDurableObjectBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerDurableObjectBinding{
				ClassName:  "TheClass",
				ScriptName: "the_script",
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_UploadWorkerWithInheritBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/zones/foo/workers/script", handler)
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerInheritBinding{},
			"b2": WorkerInheritBinding{
				OldName: "old_binding_name",
			},
		},
	}

	// Expected response
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

	// Test single-script
	res, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ZoneID: "foo"}, &scriptParams)
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}

	// Test multi-script
	res, err = client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_UploadWorkerWithKVBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerKvNamespaceBinding{
				NamespaceID: "test-namespace",
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_UploadWorkerWithWasmBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerWebAssemblyBinding{
				Module: strings.NewReader("fake-wasm"),
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_UploadWorkerWithPlainTextBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerPlainTextBinding{
				Text: "plain text value",
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_UploadWorkerAsModuleWithPlainTextBinding(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, uploadWorkerModuleResponseData) //nolint
	})

	scriptParams := WorkerScriptParams{
		Script: workerModuleScript,
		Module: true,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerPlainTextBinding{
				Text: "plain text value",
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_UploadWorkerWithSecretTextBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerSecretTextBinding{
				Text: "secret text value",
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_UploadWorkerWithServiceBinding(t *testing.T) {
	setup(UsingAccount("foo"))
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
		fmt.Fprintf(w, uploadWorkerResponseData) //nolint
	}
	mux.HandleFunc("/accounts/foo/workers/scripts/bar", handler)

	environment := "the_environment"
	scriptParams := WorkerScriptParams{
		Script: workerScript,
		Bindings: map[string]WorkerBinding{
			"b1": WorkerServiceBinding{
				Service: "the_service",
			},
			"b2": WorkerServiceBinding{
				Service:     "the_service",
				Environment: &environment,
			},
		},
	}
	_, err := client.UploadWorkerWithBindings(context.Background(), &WorkerRequestParams{ScriptName: "bar"}, &scriptParams)
	assert.NoError(t, err)
}

func TestWorkers_CreateWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
	})
	route := WorkerRoute{Pattern: "app1.example.com/*", Enabled: true}
	res, err := client.CreateWorkerRoute(context.Background(), "foo", route)
	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_CreateWorkerRouteEnt(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
	})
	route := WorkerRoute{Pattern: "app1.example.com/*", Script: "test_script"}
	res, err := client.CreateWorkerRoute(context.Background(), "foo", route)
	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_CreateWorkerRouteSingleScriptWithAccount(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
	})
	route := WorkerRoute{Pattern: "app1.example.com/*", Enabled: true}
	res, err := client.CreateWorkerRoute(context.Background(), "foo", route)
	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_CreateWorkerRouteErrorsWhenMixingSingleAndMultiScriptProperties(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	route := WorkerRoute{Pattern: "app1.example.com/*", Script: "test_script", Enabled: true}
	_, err := client.CreateWorkerRoute(context.Background(), "foo", route)
	assert.EqualError(t, err, "Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
}

func TestWorkers_CreateWorkerRouteWithNoScript(t *testing.T) {
	setup(UsingAccount("foo"))

	mux.HandleFunc("/zones/foo/workers/routes", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
	})

	route := WorkerRoute{Pattern: "app1.example.com/*"}
	_, err := client.CreateWorkerRoute(context.Background(), "foo", route)
	assert.NoError(t, err)
}

func TestWorkers_DeleteWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, deleteWorkerRouteResponseData) //nolint
	})
	res, err := client.DeleteWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1")
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID: "e7a57d8746e74ae49c25994dadb421b1",
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_DeleteWorkerRouteEnt(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, deleteWorkerRouteResponseData) //nolint
	})
	res, err := client.DeleteWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1")
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID: "e7a57d8746e74ae49c25994dadb421b1",
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_ListWorkerRoutes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, listRouteResponseData) //nolint
	})

	res, err := client.ListWorkerRoutes(context.Background(), "foo")
	want := WorkerRoutesResponse{successResponse,
		[]WorkerRoute{
			{ID: "e7a57d8746e74ae49c25994dadb421b1", Pattern: "app1.example.com/*", Enabled: true},
			{ID: "f8b68e9857f85bf59c25994dadb421b1", Pattern: "app2.example.com/*", Enabled: false},
		},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_ListWorkerRoutesEnt(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, listRouteEntResponseData) //nolint
	})

	res, err := client.ListWorkerRoutes(context.Background(), "foo")
	want := WorkerRoutesResponse{successResponse,
		[]WorkerRoute{
			{ID: "e7a57d8746e74ae49c25994dadb421b1", Pattern: "app1.example.com/*", Script: "test_script_1", Enabled: true},
			{ID: "f8b68e9857f85bf59c25994dadb421b1", Pattern: "app2.example.com/*", Script: "test_script_2", Enabled: true},
			{ID: "2b5bf4240cd34c77852fac70b1bf745a", Pattern: "app3.example.com/*", Script: "", Enabled: false},
		},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_GetWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes/1234", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, getRouteResponseData) //nolint
	})

	res, err := client.GetWorkerRoute(context.Background(), "foo", "1234")
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID:      "e7a57d8746e74ae49c25994dadb421b1",
			Pattern: "app1.example.com/*",
			Script:  "script-name"},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_UpdateWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, updateWorkerRouteResponse) //nolint
	})
	route := WorkerRoute{Pattern: "app3.example.com/*", Enabled: true}
	res, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID:      "e7a57d8746e74ae49c25994dadb421b1",
			Pattern: "app3.example.com/*",
			Enabled: true,
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_UpdateWorkerRouteEnt(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, updateWorkerRouteEntResponse) //nolint
	})
	route := WorkerRoute{Pattern: "app3.example.com/*", Script: "test_script_1"}
	res, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID:      "e7a57d8746e74ae49c25994dadb421b1",
			Pattern: "app3.example.com/*",
			Script:  "test_script_1",
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_UpdateWorkerRouteSingleScriptWithAccount(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, updateWorkerRouteEntResponse) //nolint
	})
	route := WorkerRoute{Pattern: "app3.example.com/*", Enabled: true}
	res, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID:      "e7a57d8746e74ae49c25994dadb421b1",
			Pattern: "app3.example.com/*",
			Script:  "test_script_1",
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_ListWorkerBindingsMultiScript(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/my-script/bindings", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, listBindingsResponseData) //nolint
	})

	mux.HandleFunc("/accounts/foo/workers/scripts/my-script/bindings/MY_WASM/content", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/wasm")
		_, _ = w.Write([]byte("mock multi-script wasm"))
	})

	res, err := client.ListWorkerBindings(context.Background(), &WorkerRequestParams{
		ScriptName: "my-script",
	})
	assert.NoError(t, err)

	assert.Equal(t, successResponse, res.Response)
	assert.Equal(t, 7, len(res.BindingList))

	assert.Equal(t, res.BindingList[0], WorkerBindingListItem{
		Name: "MY_KV",
		Binding: WorkerKvNamespaceBinding{
			NamespaceID: "89f5f8fd93f94cb98473f6f421aa3b65",
		},
	})
	assert.Equal(t, WorkerKvNamespaceBindingType, res.BindingList[0].Binding.Type())

	assert.Equal(t, "MY_WASM", res.BindingList[1].Name)
	wasmBinding := res.BindingList[1].Binding.(WorkerWebAssemblyBinding)
	wasmModuleContent, err := ioutil.ReadAll(wasmBinding.Module)
	assert.NoError(t, err)
	assert.Equal(t, []byte("mock multi-script wasm"), wasmModuleContent)
	assert.Equal(t, WorkerWebAssemblyBindingType, res.BindingList[1].Binding.Type())

	assert.Equal(t, res.BindingList[2], WorkerBindingListItem{
		Name: "MY_PLAIN_TEXT",
		Binding: WorkerPlainTextBinding{
			Text: "text",
		},
	})
	assert.Equal(t, WorkerPlainTextBindingType, res.BindingList[2].Binding.Type())

	assert.Equal(t, res.BindingList[3], WorkerBindingListItem{
		Name:    "MY_SECRET_TEXT",
		Binding: WorkerSecretTextBinding{},
	})
	assert.Equal(t, WorkerSecretTextBindingType, res.BindingList[3].Binding.Type())

	environment := "MY_ENVIRONMENT"
	assert.Equal(t, res.BindingList[4], WorkerBindingListItem{
		Name: "MY_SERVICE_BINDING",
		Binding: WorkerServiceBinding{
			Service:     "MY_SERVICE",
			Environment: &environment,
		},
	})
	assert.Equal(t, WorkerServiceBindingType, res.BindingList[4].Binding.Type())

	assert.Equal(t, res.BindingList[5], WorkerBindingListItem{
		Name:    "MY_NEW_BINDING",
		Binding: WorkerInheritBinding{},
	})
	assert.Equal(t, WorkerInheritBindingType, res.BindingList[5].Binding.Type())

	assert.Equal(t, res.BindingList[6], WorkerBindingListItem{
		Name: "MY_BUCKET",
		Binding: WorkerR2BucketBinding{
			BucketName: "bucket",
		},
	})
	assert.Equal(t, WorkerR2BucketBindingType, res.BindingList[6].Binding.Type())
}

func TestWorkers_UpdateWorkerRouteErrorsWhenMixingSingleAndMultiScriptProperties(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	route := WorkerRoute{Pattern: "app1.example.com/*", Script: "test_script", Enabled: true}
	_, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
	assert.EqualError(t, err, "Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
}

func TestWorkers_UpdateWorkerRouteWithNoScript(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, updateWorkerRouteEntResponse) //nolint
	})

	route := WorkerRoute{Pattern: "app1.example.com/*"}
	_, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
	assert.NoError(t, err)
}

func TestWorkers_AttachWorkerToDomain(t *testing.T) {
	setup(UsingAccount(testAccountID))
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/domains", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, attachWorkerToDomainResponse) //nolint
	})
	res, err := client.AttachWorkerToDomain(context.Background(), AccountIdentifier(testAccountID), &WorkerDomainParams{
		ZoneID:      testZoneID,
		Hostname:    "app4.example.com",
		Service:     "test_script_1",
		Environment: "production",
	})
	want := WorkerDomainResponse{
		successResponse,
		WorkerDomainResult{
			ID:          "e7a57d8746e74ae49c25994dadb421b1",
			ZoneID:      testZoneID,
			Service:     "test_script_1",
			Hostname:    "api4.example.com",
			Environment: "production",
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}
