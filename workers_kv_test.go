package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkersKV_CreateStorageNamespace(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	response := `{
		"result": {
			"id" : "3aeaxxxxee014exxxx4cf66xxxxc0448",
			"title": "test_namespace"
		},
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc("/accounts/foo/storage/kv/namespaces", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.CreateStorageNamespace(context.Background(), &StorageNamespaceRequest{Title: "Namespace"})
	want := StorageNamespaceResponse{
		successResponse,
		StorageNamespace{
			ID:    "3aeaxxxxee014exxxx4cf66xxxxc0448",
			Title: "test_namespace",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_DeleteStorageNamespace(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	namespace := "3aeaxxxxee014exxxx4cf66xxxxc0448"
	response := `{
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s", namespace), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.DeleteStorageNamespace(context.Background(), namespace)

	if assert.NoError(t, err) {
		assert.Equal(t, successResponse, res)
	}
}

func TestWorkersKV_ListStorageNamespace(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	response := `{
		"result": [
			{"id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"title": "test_namespace_1"
			},
			{"id": "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			"title": "test_namespace_2"
			}
		],
		"success": true,
		"errors": [],
		"messages": [],
		"result_info": {
			"page": 1,
			"per_page": 20,
			"count": 2,
			"total_count": 2,
			"total_pages": 1
		}
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.ListStorageNamespaces(context.Background())
	want := ListStorageNamespacesResponse{
		successResponse,
		[]StorageNamespace{
			{
				ID:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				Title: "test_namespace_1",
			},
			{
				ID:    "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
				Title: "test_namespace_2",
			},
		},
		ResultInfo{
			Page:       1,
			PerPage:    20,
			Count:      2,
			TotalPages: 1,
			Total:      2,
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
		assert.Equal(t, want.ResultInfo, res.ResultInfo)

		sort.Slice(res.Result, func(i, j int) bool {
			return res.Result[i].ID < res.Result[j].ID
		})
		sort.Slice(want.Result, func(i, j int) bool {
			return want.Result[i].ID < want.Result[j].ID
		})
		assert.Equal(t, res.Result, want.Result)
	}
}

func TestWorkersKV_UpdateStorageNamespace(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	response := `{
		"result": null,
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s", namespace), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.UpdateStorageNamespace(context.Background(), namespace, &StorageNamespaceRequest{Title: "Namespace"})
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_CreateStorageKV(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	key := "test_key"
	value := []byte("test_value")
	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	response := `{
		"result": null,
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/values/%s", namespace, key), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "binary/octet-stream")
		fmt.Fprintf(w, response)
	})

	want := successResponse
	res, err := client.CreateStorageKV(context.Background(), namespace, key, bytes.NewReader(value))

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_CreateStorageKVRetryResetsBody(t *testing.T) {
	setup(UsingOrganization("foo"), UsingRetryPolicy(1, 0, 0))
	defer teardown()

	key := "test_key"
	value := []byte("test_value")
	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	response := `{
		"result": null,
		"success": true,
		"errors": [],
		"messages": []
	}`

	try := 0
	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/values/%s", namespace, key), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)

		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, value, res)

		// Fail the first time after reading the request payload
		if try < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			try++
		}

		w.Header().Set("content-type", "binary/octet-stream")
		fmt.Fprintf(w, response)
	})

	reader := bytes.NewReader(value)

	res, err := client.CreateStorageKV(context.Background(), namespace, key, reader)

	if assert.NoError(t, err) {
		assert.Equal(t, successResponse, res)
	}
}

type noopReader struct {
}

func (noopReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func TestWorkersKV_CreateStorageKVRetryUnsupportedSeek(t *testing.T) {
	setup(UsingOrganization("foo"), UsingRetryPolicy(1, 0, 0))
	defer teardown()

	key := "test_key"
	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/values/%s", namespace, key), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateStorageKV(context.Background(), namespace, key, noopReader{})

	assert.EqualError(t, err, "error from makeRequest: Request body does not support io.Seek and can't be retried")
}

func TestWorkersKV_ReadStorageKV(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	key := "test_key"
	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/values/%s", namespace, key), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "text/plain")
		fmt.Fprintf(w, "test_value")
	})

	res, err := client.ReadStorageKV(context.Background(), namespace, key)
	want := []byte("test_value")

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_DeleteStorageKV(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	key := "test_key"
	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	response := `{
		"result": null,
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/values/%s", namespace, key), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.DeleteStorageKV(context.Background(), namespace, key)
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_ListStorageKeys(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	response := `{
		"result": [
			{"name": "test_key_1"},
			{"name": "test_key_2"}
		],
		"success": true,
		"errors": [],
		"messages": [],
		"result_info": {
			"page": 1,
			"per_page": 20,
			"count": 2,
			"total_count": 2,
			"total_pages": 1
		}
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/keys", namespace), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.ListStorageKeys(context.Background(), namespace)

	want := ListStorageKeysResponse{
		successResponse,
		[]StorageKey{
			{Name: "test_key_1"},
			{Name: "test_key_2"},
		},
		ResultInfo{
			Page:       1,
			PerPage:    20,
			Count:      2,
			TotalPages: 1,
			Total:      2,
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
		assert.Equal(t, want.ResultInfo, res.ResultInfo)

		sort.Slice(res.Result, func(i, j int) bool {
			return res.Result[i].Name < res.Result[j].Name
		})

		sort.Slice(want.Result, func(i, j int) bool {
			return want.Result[i].Name < want.Result[j].Name
		})
		assert.Equal(t, want, res)
	}
}
