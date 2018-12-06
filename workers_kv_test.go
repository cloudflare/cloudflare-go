package cloudflare

import (
	"context"
	"fmt"
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

	mux.HandleFunc("/accounts/foo/workers/namespaces", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.CreateStorageNamespace(context.Background(), &StorageNamespaceRequest{Title: "Namespace"})
	want := StorageNamespaceResponse{
		successResponse,
		"3aeaxxxxee014exxxx4cf66xxxxc0448",
		"test_namespace",
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}

func TestWorkersKV_DeleteStorageNamespace(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	namespace := "3aeaxxxxee014exxxx4cf66xxxxc0448"
	response := `{
		"result": null,
		"success": true,
		"errors": [],
		"messages": []
	}`

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/workers/namespaces/%s", namespace), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.DeleteStorageNamespace(context.Background(), namespace)
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersKV_ListStorageNamespaces(t *testing.T) {
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

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/workers/namespaces"), func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc(fmt.Sprintf("/accounts/foo/workers/namespaces/%s", namespace), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, response)
	})

	res, err := client.UpdateStorageNamespace(context.Background(), namespace, &StorageNamespaceRequest{Title: "Namespace"})
	want := successResponse

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}
