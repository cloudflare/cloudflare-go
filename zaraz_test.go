package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"dataLayer": true,
			"debugKey": "cluu08t8nne8gq9g3160",
			"dlp": [],
			"historyChange": true,
			"settings": {
				"autoInjectScript": true
			},
			"tools": {
				"PBQr": {
					"blockingTriggers": [],
					"component": "html",
					"defaultFields": {},
					"enabled": true,
					"mode": {
					  "cloud": false,
					  "ignoreSPA": true,
					  "light": false,
					  "sample": false,
					  "segment": {
						"end": 100,
						"start": 0
					  },
					  "trigger": "pageload"
					},
					"name": "Custom HTML",
					"neoEvents": [],
					"permissions": [
					  "execute_unsafe_scripts"
					],
					"settings": {},
					"type": "component"
				  }
			},
			"triggers": {
				"Pageview": {
					"clientRules": [],
					"description": "All page loads",
					"excludeRules": [],
					"loadRules": [
					  {
						"match": "{{ client.__zarazTrack }}",
						"op": "EQUALS",
						"value": "Pageview"
					  }
					],
					"name": "Pageview",
					"system": "pageload"
				}
			},
			"variables": {},
			"zarazVersion": 44
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/config", handler)
	want := ZarazConfig{
		"dataLayer":     true,
		"debugKey":      "cluu08t8nne8gq9g3160",
		"dlp":           []interface{}{},
		"historyChange": true,
		"settings": map[string]interface{}{
			"autoInjectScript": true,
		},
		"tools": map[string]interface{}{
			"PBQr": map[string]interface{}{
				"blockingTriggers": []interface{}{},
				"component":        "html",
				"defaultFields":    map[string]interface{}{},
				"enabled":          true,
				"mode": map[string]interface{}{
					"cloud":     false,
					"ignoreSPA": true,
					"light":     false,
					"sample":    false,
					"segment": map[string]interface{}{
						"end":   float64(100),
						"start": float64(0),
					},
					"trigger": "pageload",
				},
				"name":        "Custom HTML",
				"neoEvents":   []interface{}{},
				"permissions": []interface{}{"execute_unsafe_scripts"},
				"settings":    map[string]interface{}{},
				"type":        "component",
			},
		},
		"triggers": map[string]interface{}{
			"Pageview": map[string]interface{}{
				"clientRules": []interface{}{},
				"description": "All page loads",
				"loadRules": []interface{}{map[string]interface{}{
					"match": "{{ client.__zarazTrack }}",
					"op":    "EQUALS",
					"value": "Pageview",
				}},
				"excludeRules": []interface{}{},
				"name":         "Pageview",
				"system":       "pageload",
			},
		},
		"variables":    map[string]interface{}{},
		"zarazVersion": float64(44),
	}

	actual, err := client.GetZarazConfig(context.Background(), ZoneIdentifier(testZoneID))
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestUpdateZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"dataLayer": true,
			"debugKey": "cluu08t8nne8gq9g3160",
			"dlp": [],
			"historyChange": true,
			"settings": {
				"autoInjectScript": true
			},
			"tools": {
				"PBQr": {
					"blockingTriggers": [],
					"component": "html",
					"defaultFields": {},
					"enabled": false,
					"mode": {
					  "cloud": false,
					  "ignoreSPA": true,
					  "light": false,
					  "sample": false,
					  "segment": {
						"end": 100,
						"start": 0
					  },
					  "trigger": "pageload"
					},
					"name": "Custom HTML",
					"neoEvents": [],
					"permissions": [
					  "execute_unsafe_scripts"
					],
					"settings": {},
					"type": "component"
				  }
			},
			"triggers": {
				"Pageview": {
					"clientRules": [],
					"description": "All page loads",
					"excludeRules": [],
					"loadRules": [
					  {
						"match": "{{ client.__zarazTrack }}",
						"op": "EQUALS",
						"value": "Pageview"
					  }
					],
					"name": "Pageview",
					"system": "pageload"
				}
			},
			"variables": {},
			"zarazVersion": 44
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/config", handler)
	payload := ZarazConfig{
		"dataLayer":     true,
		"debugKey":      "cluu08t8nne8gq9g3160",
		"dlp":           []interface{}{},
		"historyChange": true,
		"settings": map[string]interface{}{
			"autoInjectScript": true,
		},
		"tools": map[string]interface{}{
			"PBQr": map[string]interface{}{
				"blockingTriggers": []interface{}{},
				"component":        "html",
				"defaultFields":    map[string]interface{}{},
				"enabled":          true,
				"mode": map[string]interface{}{
					"cloud":     false,
					"ignoreSPA": true,
					"light":     false,
					"sample":    false,
					"segment": map[string]interface{}{
						"end":   float64(100),
						"start": float64(0),
					},
					"trigger": "pageload",
				},
				"name":        "Custom HTML",
				"neoEvents":   []interface{}{},
				"permissions": []interface{}{"execute_unsafe_scripts"},
				"settings":    map[string]interface{}{},
				"type":        "component",
			},
		},
		"triggers": map[string]interface{}{
			"Pageview": map[string]interface{}{
				"clientRules": []interface{}{},
				"description": "All page loads",
				"loadRules": []interface{}{map[string]interface{}{
					"match": "{{ client.__zarazTrack }}",
					"op":    "EQUALS",
					"value": "Pageview",
				}},
				"excludeRules": []interface{}{},
				"name":         "Pageview",
				"system":       "pageload",
			},
		},
		"variables":    map[string]interface{}{},
		"zarazVersion": float64(44),
	}

	err := client.UpdateZarazConfig(context.Background(), ZoneIdentifier(testZoneID), payload)
	require.NoError(t, err)
}

func TestGetZarazWorkflow(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, `realtime`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/workflow", handler)
	want := "realtime"

	actual, err := client.GetZarazWorkflow(context.Background(), ZoneIdentifier(testZoneID))

	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestUpdateZarazWorkflow(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		w.Header().Set("Content-Type", "text/plain")
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/workflow", handler)
	payload := "realtime"

	err := client.UpdateZarazWorkflow(context.Background(), ZoneIdentifier(testZoneID), payload)

	require.NoError(t, err)

}

func TestPublishZarazConfig(t *testing.T) {}

func TestGetZarazConfigHistory(t *testing.T) {}

func TestGetZarazConfigHistoryDiff(t *testing.T) {}

func TestGetDefaultZarazConfig(t *testing.T) {}

func TestExportZarazConfig(t *testing.T) {}
