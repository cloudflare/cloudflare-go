package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestR2_ListObjects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/r2/buckets/%s/objects", testAccountID, testBucketName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
	{
		"key": "cf1.txt",
		"etag": "d41d8cd98f00b204e9800998ecf8427e",
		"last_modified": "2022-06-24T19:58:49.477Z",
		"size": 104,
		"http_metadata": {"contentType":"text/plain"},
		"custom_metadata": {}
	},
	{
		"key": "cf2.txt",
		"etag": "d41d8cd98f00b204e9800998ecf8427e",
		"last_modified": "2022-06-24T19:58:49.477Z",
		"size": 104,
		"http_metadata": {"contentType":"text/plain"},
		"custom_metadata": {}
	}
  ],
  "result_info": {
	"per_page": 2,
	"cursor": "some-text",
	"is_truncated": true
  }
}`)
	})

	lastModified, _ := time.Parse(time.RFC3339, "2022-06-24T19:58:49.477Z")
	want := []R2Object{
		{
			Key:          "cf1.txt",
			Etag:         "d41d8cd98f00b204e9800998ecf8427e",
			LastModified: lastModified,
			Size:         104,
			HTTPMetadata: map[string]any{
				"contentType": "text/plain",
			},
			CustomMetadata: map[string]string{},
		},
		{
			Key:          "cf2.txt",
			Etag:         "d41d8cd98f00b204e9800998ecf8427e",
			LastModified: lastModified,
			Size:         104,
			HTTPMetadata: map[string]any{
				"contentType": "text/plain",
			},
			CustomMetadata: map[string]string{},
		},
	}
	wantInfo := &ResultInfo{
		PerPage:     2,
		Cursor:      "some-text",
		IsTruncated: true,
	}
	actual, info, err := client.ListR2Objects(context.Background(), AccountIdentifier(testAccountID), testBucketName, ListR2ObjectsParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
		assert.Equal(t, wantInfo, info)
	}
}
