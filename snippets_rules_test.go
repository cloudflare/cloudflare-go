package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnippetsRules(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": [
		{
			"id": "some_id_1",
			"expression": "true",
			"enabled": true,
			"description": "some description",
			"snippet_name": "snippet_1"
	  	},
		{
			"id": "some_id_2",
			"expression": "true",
			"enabled": true,
			"description": "some description",
			"snippet_name": "snippet_2"
	  	}
	  ],
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/snippets/rules", handler)

	want := []SnippetRule{
		{
			ID:          "some_id_1",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			SnippetName: "snippet_1",
		},
		{
			ID:          "some_id_2",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			SnippetName: "snippet_2",
		},
	}

	zoneActual, err := client.ListZoneSnippetsRules(context.Background(), ZoneIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}
}

func TestUpdateSnippetsRules(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": [
		{
			"id": "some_id_1",
			"expression": "true",
			"enabled": false,
			"description": "some description",
			"snippet_name": "snippet_1"
	  	},
		{
			"id": "some_id_2",
			"expression": "true",
			"enabled": false,
			"description": "some description",
			"snippet_name": "snippet_2"
	  	}
	  ],
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/snippets/rules", handler)
	toUpdate := []SnippetRule{
		{
			Expression:  "true",
			Enabled:     BoolPtr(false),
			Description: "some description",
			SnippetName: "snippet_1",
		},
		{
			Expression:  "true",
			Enabled:     BoolPtr(false),
			Description: "some description",
			SnippetName: "snippet_2",
		},
	}
	want := []SnippetRule{
		{
			ID:          "some_id_1",
			Expression:  "true",
			Enabled:     BoolPtr(false),
			Description: "some description",
			SnippetName: "snippet_1",
		},
		{
			ID:          "some_id_2",
			Expression:  "true",
			Enabled:     BoolPtr(false),
			Description: "some description",
			SnippetName: "snippet_2",
		},
	}
	zoneActual, err := client.UpdateZoneSnippetsRules(context.Background(), ZoneIdentifier(testZoneID), toUpdate)
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}
}
