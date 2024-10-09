package cloudflare

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSnippets(t *testing.T) {
	setup()
	defer teardown()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "result": [
    {
      "created_on": "2024-01-12T17:40:39Z",
      "modified_on": "2024-01-12T17:40:39Z",
      "snippet_name": "hello_world"
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 25,
    "total_pages": 1,
    "count": 1,
    "total_count": 1
  }
}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/snippets", handler)

	t.Run("Lists snippets found in the response", func(t *testing.T) {
		params := ListSnippetsParams{}
		response, _, err := client.ListSnippets(context.Background(), ZoneIdentifier(testZoneID), params)
		assert.NoError(t, err)

		assert.Equal(t, "hello_world", response.Snippets[0].SnippetName)
	})
}

func TestDeleteSnippets(t *testing.T) {
	setup()
	defer teardown()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "errors":[],
  "messages":[],
  "success": true
 }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/snippets/hello_world", handler)

	t.Run("Deletes a snippet", func(t *testing.T) {
		params := DeleteSnippetParams{
			SnippetName: "hello_world",
		}
		response, err := client.DeleteSnippet(context.Background(), ZoneIdentifier(testZoneID), params)
		assert.NoError(t, err)

		assert.Equal(t, true, response.Success)
	})
}

func TestGetSnippet(t *testing.T) {
	setup()
	defer teardown()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "result": {
    "created_on": "2024-01-12T17:40:39Z",
    "modified_on": "2024-01-12T17:40:39Z",
    "snippet_name": "hello_world"
  }
}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/snippets/hello_world", handler)

	t.Run("Gets the hello_world snippet", func(t *testing.T) {
		params := GetSnippetParams{
			SnippetName: "hello_world",
		}
		response, err := client.GetSnippet(context.Background(), ZoneIdentifier(testZoneID), params)
		assert.NoError(t, err)

		assert.Equal(t, true, response.Success)
		assert.Equal(t, "hello_world", response.Snippet.SnippetName)
	})
}

func TestUploadSnippet(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		err := r.ParseMultipartForm(1024 * 1024 * 10)
		assert.NoError(t, err)

		assert.Equal(t, "{\"main_module\":\"main.js\"}", r.MultipartForm.Value["metadata"][0])

		fileHeader, ok := r.MultipartForm.File["main.js"]
		assert.True(t, ok)

		file, err := fileHeader[0].Open()
		assert.NoError(t, err)
		defer file.Close()

		content, err := io.ReadAll(file)
		assert.NoError(t, err)

		assert.Equal(t, "main.js", fileHeader[0].Filename)
		assert.Equal(t, "console.log('main.js')", string(content))

		fileHeader, ok = r.MultipartForm.File["worker.js"]
		assert.True(t, ok)

		file, err = fileHeader[0].Open()
		assert.NoError(t, err)
		defer file.Close()

		content, err = io.ReadAll(file)
		assert.NoError(t, err)

		assert.Equal(t, "worker.js", fileHeader[0].Filename)
		assert.Equal(t, "console.log('worker.js')", string(content))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "errors": [],
  "messages": [],
  "success": true,
  "result": {
    "created_on": "2023-07-24-00:00:00",
    "modified_on": "2023-07-24-00:00:00",
    "snippet_name": "test_snippet"
  }
}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/snippets/test_snippet", handler)

	t.Run("Test multiples files are encoded in the request body", func(t *testing.T) {
		params := PutSnippetParams{
			SnippetName: "test_snippet",
			MainModule:  "main.js",
			Files: map[string]string{
				"main.js":   "console.log('main.js')",
				"worker.js": "console.log('worker.js')",
			},
		}
		response, err := client.PutSnippet(context.Background(), ZoneIdentifier(testZoneID), params)
		assert.NoError(t, err)

		assert.Equal(t, "test_snippet", response.Snippet.SnippetName)
	})
}
