package cloudflare

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessApplications(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
					"created_at": "2014-01-01T05:20:00.12345Z",
					"updated_at": "2014-01-01T05:20:00.12345Z",
					"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
					"name": "Admin Site",
					"domain": "test.example.com/admin",
					"session_duration": "24h"
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000
			}
		}
		`)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/access/apps", handler)
	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := []AccessApplication{AccessApplication{
		ID:              "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		CreatedAt:       &createdAt,
		UpdatedAt:       &updatedAt,
		AUD:             "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		Name:            "Admin Site",
		Domain:          "test.example.com/admin",
		SessionDuration: "24h",
	}}

	actual, _, err := client.AccessApplications("01a7362d577a6c3019a474fd6f485823", PaginationOptions{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
