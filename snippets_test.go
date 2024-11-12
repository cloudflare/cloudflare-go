package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSnippets(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": [
		{
			"snippet_name": "some_id_1",
			"created_on":"0001-01-01T00:00:00Z",
			"modified_on":"0001-01-01T00:00:00Z"
	  	},
		{
			"snippet_name": "some_id_2",
			"created_on":"0001-01-01T00:00:00Z",
			"modified_on":"0001-01-01T00:00:00Z"
	  	}
	  ],
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/snippets", handler)

	want := []Snippet{
		{
			SnippetName: "some_id_1",
			CreatedOn:   &time.Time{},
			ModifiedOn:  &time.Time{},
		},
		{
			SnippetName: "some_id_2",
			CreatedOn:   &time.Time{},
			ModifiedOn:  &time.Time{},
		},
	}

	zoneActual, err := client.ListZoneSnippets(context.Background(), ZoneIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}
}

func TestGetSnippet(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": {
			"snippet_name": "some_id_1",
			"created_on":"0001-01-01T00:00:00Z",
			"modified_on":"0001-01-01T00:00:00Z"
		},
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/snippets/some_id_1", handler)

	want := Snippet{
		SnippetName: "some_id_1",
		CreatedOn:   &time.Time{},
		ModifiedOn:  &time.Time{},
	}

	zoneActual, err := client.GetZoneSnippet(context.Background(), ZoneIdentifier(testZoneID), "some_id_1")
	if assert.NoError(t, err) {
		assert.Equal(t, want, *zoneActual)
	}
}

func TestDeleteSnippet(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": null,
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/snippets/some_id_1", handler)

	err := client.DeleteZoneSnippet(context.Background(), ZoneIdentifier(testZoneID), "some_id_1")
	assert.NoError(t, err)
}

func TestUpdateSnippets(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result":
		{
			"snippet_name": "some_id_1",
			"created_on":"0001-01-01T00:00:00Z",
			"modified_on":"0001-01-01T00:00:00Z"
	  	},
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/snippets/some_id_1", handler)
	toUpdate := SnippetRequest{
		SnippetName: "some_id_1",
		MainFile:    "file1.js",
		Files: []SnippetFile{
			{
				FileName: "file1.js",
				Content:  "test1.json",
			},
			{
				FileName: "file2.js",
				Content:  "test2.json",
			},
		},
	}
	want := &Snippet{
		SnippetName: "some_id_1",
		CreatedOn:   &time.Time{},
		ModifiedOn:  &time.Time{},
	}
	zoneActual, err := client.UpdateZoneSnippet(context.Background(), ZoneIdentifier(testZoneID), toUpdate)
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}
}
