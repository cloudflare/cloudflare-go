package cloudflare

import (
	"context" //nolint:gosec
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTieredCacheSmartTopology(t *testing.T) {
	setup()
	defer teardown()
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
					"id": "tiered_cache_smart_topology_enable",
					"modified_on": "2022-06-01T17:29:22.334114Z",
					"value": "on"
			},
			"success": true,
			"errors": [],
			"messages": []
	}`)
	}
	mux.HandleFunc("/zones/foo/cache/tiered_cache_smart_topology_enable", handler)
	s, err := client.TieredCacheSmartTopology(context.Background(), "foo")
	if assert.NoError(t, err) {
		assert.Equal(t, s.ID, "tiered_cache_smart_topology_enable")
		assert.Equal(t, s.Value, TieredCacheSmartTopologyOn)
		assert.Equal(t, s.ModifiedOn, "2022-06-01T17:29:22.334114Z")
	}
}

func TestUpdateTieredCacheSmartTopology(t *testing.T) {
	setup()
	defer teardown()
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
					"id": "tiered_cache_smart_topology_enable",
					"modified_on": "2022-06-01T17:29:22.334114Z",
					"value": "off"
			},
			"success": true,
			"errors": [],
			"messages": []
	}`)
	}
	mux.HandleFunc("/zones/foo/cache/tiered_cache_smart_topology_enable", handler)
	s, err := client.UpdateTieredCacheSmartTopology(context.Background(), "foo", TieredCacheSmartTopologyUpdateOptions{Value: TieredCacheSmartTopologyOff})
	if assert.NoError(t, err) {
		assert.Equal(t, s.ID, "tiered_cache_smart_topology_enable")
		assert.Equal(t, s.Value, TieredCacheSmartTopologyOff)
		assert.Equal(t, s.ModifiedOn, "2022-06-01T17:29:22.334114Z")
	}
}
