package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var regionalTieredCacheTimestampString = "2019-02-20T22:37:07.107449Z"
var regionalTieredCacheTimestamp, _ = time.Parse(time.RFC3339Nano, regionalTieredCacheTimestampString)

func TestRegionalTieredCache(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "regional_tiered_cache",
				"value": "on",
				"modified_on": "%s"
			}
		}
		`, regionalTieredCacheTimestampString)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/cache/regional_tiered_cache", handler)
	want := RegionalTieredCache{
		ID:         "regional_tiered_cache",
		Value:      "on",
		ModifiedOn: regionalTieredCacheTimestamp,
	}

	actual, err := client.GetRegionalTieredCache(
		context.Background(),
		ZoneIdentifier("01a7362d577a6c3019a474fd6f485823"),
		GetRegionalTieredCacheParams{},
	)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateRegionalTieredCache(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "regional_tiered_cache",
				"value": "off",
				"modified_on": "%s"
			}
		}
		`, regionalTieredCacheTimestampString)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/cache/regional_tiered_cache", handler)
	want := RegionalTieredCache{
		ID:         "regional_tiered_cache",
		Value:      "off",
		ModifiedOn: regionalTieredCacheTimestamp,
	}

	actual, err := client.UpdateRegionalTieredCache(
		context.Background(),
		ZoneIdentifier("01a7362d577a6c3019a474fd6f485823"),
		UpdateRegionalTieredCacheParams{Value: "off"},
	)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
