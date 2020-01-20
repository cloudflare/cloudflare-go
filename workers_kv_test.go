package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"sort"
	"testing"
)

func TestWorkersKV_CreateWorkersKVNamespace(t *testing.T) {
	setup(UsingAccount("foo"))
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

	res, err := client.CreateWorkersKVNamespace(context.Background(), &WorkersKVNamespaceRequest{Title: "Namespace"})
	want := WorkersKVNamespaceResponse{
		successResponse,
		WorkersKVNamespace{
			ID:    "3aeaxxxxee014exxxx4cf66xxxxc0448",
			Title: "test_namespace",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}

func TestWorkersKV_DeleteWorkersKVNamespace(t *testing.T) {
	setup(UsingAccount("foo"))
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

	res, err := client.DeleteWorkersKVNamespace(context.Background(), namespace)
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_ListWorkersKVNamespace(t *testing.T) {
	setup(UsingAccount("foo"))
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
			"per_page": 100,
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

	res, err := client.ListWorkersKVNamespaces(context.Background())
	want := []WorkersKVNamespace{
		WorkersKVNamespace {
			ID:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Title: "test_namespace_1",
		},
		WorkersKVNamespace {
			ID:    "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			Title: "test_namespace_2",
		},
	}

	if assert.NoError(t, err) {
		sort.Slice(res, func(i, j int) bool {
			return res[i].ID < res[j].ID
		})
		sort.Slice(want, func(i, j int) bool {
			return want[i].ID < want[j].ID
		})
		assert.Equal(t, res, want)
	}
}

func TestWorkersKV_ListWorkersKVNamespaceMultiplePages(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	response1 := `{
		"result": [
			{"id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"title": "test_namespace_1"
			}
		],
		"success": true,
		"errors": [],
		"messages": [],
		"result_info": {
			"page": 1,
			"per_page": 100,
			"count": 1,
			"total_count": 2,
			"total_pages": 2
		}
	}`

	response2 := `{
		"result": [
			{"id": "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			"title": "test_namespace_2"
			}
		],
		"success": true,
		"errors": [],
		"messages": [],
		"result_info": {
			"page": 2,
			"per_page": 100,
			"count": 1,
			"total_count": 2,
			"total_pages": 2
		}
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")

		if r.URL.Query().Get("page") == "1" {
			fmt.Fprintf(w, response1)
			return
		} else if r.URL.Query().Get("page") == "2" {
			fmt.Fprintf(w, response2)
			return
		} else {
			panic(errors.New("Got a request for an unexpected page"))
		}
	})

	res, err := client.ListWorkersKVNamespaces(context.Background())
	want := []WorkersKVNamespace{
		WorkersKVNamespace {
			ID:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Title: "test_namespace_1",
		},
		WorkersKVNamespace {
			ID:    "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			Title: "test_namespace_2",
		},
	}

	if assert.NoError(t, err) {
		sort.Slice(res, func(i, j int) bool {
			return res[i].ID < res[j].ID
		})
		sort.Slice(want, func(i, j int) bool {
			return want[i].ID < want[j].ID
		})
		assert.Equal(t, res, want)
	}
}

func TestWorkersKV_UpdateWorkersKVNamespace(t *testing.T) {
	setup(UsingAccount("foo"))
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

	res, err := client.UpdateWorkersKVNamespace(context.Background(), namespace, &WorkersKVNamespaceRequest{Title: "Namespace"})
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_WriteWorkersKV(t *testing.T) {
	setup(UsingAccount("foo"))
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
		w.Header().Set("content-type", "application/octet-stream")
		fmt.Fprintf(w, response)
	})

	want := successResponse
	res, err := client.WriteWorkersKV(context.Background(), namespace, key, value)

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_WriteWorkersKVBulk(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	kvs := WorkersKVBulkWriteRequest{{Key: "key1", Value: "value1"}, {Key: "key2", Value: "value2"}}

	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	response := `{
		"result": null,
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/bulk", namespace), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, response)
	})

	want := successResponse
	res, err := client.WriteWorkersKVBulk(context.Background(), namespace, kvs)
	require.NoError(t, err)
	assert.Equal(t, want, res)
}

func TestWorkersKV_ReadWorkersKV(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	key := "test_key"
	namespace := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/storage/kv/namespaces/%s/values/%s", namespace, key), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "text/plain")
		fmt.Fprintf(w, "test_value")
	})

	res, err := client.ReadWorkersKV(context.Background(), namespace, key)
	want := []byte("test_value")

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_DeleteWorkersKV(t *testing.T) {
	setup(UsingAccount("foo"))
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

	res, err := client.DeleteWorkersKV(context.Background(), namespace, key)
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_ListStorageKeys(t *testing.T) {
	setup(UsingAccount("foo"))
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

	res, err := client.ListWorkersKVs(context.Background(), namespace)

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
		assert.Equal(t, want.Result, res.Result)
	}
}
