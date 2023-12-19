package cloudflare

import (
	"context"
	"fmt"
	"io"
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
			"debugKey": "cheese",
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
		"debugKey":      "cheese",
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

	payload := "realtime"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		body, _ := io.ReadAll(r.Body)
		bodyString := string(body)
		assert.Equal(t, fmt.Sprintf("\"%s\"", payload), bodyString)

		w.Header().Set("Content-Type", "text/plain")
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/workflow", handler)

	err := client.UpdateZarazWorkflow(context.Background(), ZoneIdentifier(testZoneID), payload)

	require.NoError(t, err)

}

func TestPublishZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	payload := "test description"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		body, _ := io.ReadAll(r.Body)
		bodyString := string(body)
		assert.Equal(t, fmt.Sprintf("\"%s\"", payload), bodyString)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/publish", handler)

	err := client.PublishZarazConfig(context.Background(), ZoneIdentifier(testZoneID), payload)

	require.NoError(t, err)
}

func TestGetZarazConfigHistory(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"count": 1,
			"data": [
				{
					"createdAt": "2023-12-19T19:24:42.779683Z",
					"description": "Moving to Preview & Publish workflow",
					"id": 1005135,
					"updatedAt": "2023-12-19T19:24:42.779683Z",
					"userId": "9ceddf6f117afe04c64716c83468d3a4"
				}
			]
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/history", handler)

	limit := int64(1)
	offset := int64(10)
	_, err := client.GetZarazConfigHistory(context.Background(), ZoneIdentifier(testZoneID), limit, offset)
	require.NoError(t, err)
}

func TestGetZarazConfigsById(t *testing.T) {
	setup()
	defer teardown()

	ids := "1234,5678"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		params := r.URL.Query()
		expectedIds := params.Get("ids")
		assert.Equal(t, ids, expectedIds)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"prop1": {
				"testKey": "cheese"
			},
			"prop2": {
				"testKey": "hedgehog"
			}
		}`)
	}

	uri := "/zones/" + testZoneID + "/settings/zaraz/history/configs?ids=" + ids
	mux.HandleFunc(uri, handler)

	_, err := client.GetZarazConfigsById(context.Background(), ZoneIdentifier(testZoneID), ids)
	require.NoError(t, err)
}

func TestGetDefaultZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "text/plain")
		fmt.Fprint(w, `{
			"someTestKeyThatRepsTheConfig": "test"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/default", handler)

	_, err := client.GetDefaultZarazConfig(context.Background(), ZoneIdentifier(testZoneID))
	require.NoError(t, err)
}

func TestExportZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "text/plain")
		fmt.Fprint(w, `{
			"someTestKeyThatRepsTheConfig": "test"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/export", handler)

	err := client.ExportZarazConfig(context.Background(), ZoneIdentifier(testZoneID))
	require.NoError(t, err)

}
