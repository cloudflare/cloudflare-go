package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDevicePostureRules(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
					"schedule": "1h",
					"type": "file",
					"name": "My rule name",
					"description": "My description",
					"match": [
						{
							"platform": "ios"
						}
					],
					"input": {
						"id": "9e597887-345e-4a32-a09c-68811b129768",
						"path": "/tmp/data.zta",
						"exists": true,
						"thumbprint": "asdfasdfasdfasdf",
						"sha256": "D75398FC796D659DEB4170569DCFEC63E3897C71E3AE8642FD3139A554AEE21E",
						"running": true
					}
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

	want := []DevicePostureRule{{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "file",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input: DevicePostureRuleInput{
			ID:         "9e597887-345e-4a32-a09c-68811b129768",
			Path:       "/tmp/data.zta",
			Exists:     true,
			Thumbprint: "asdfasdfasdfasdf",
			Sha256:     "D75398FC796D659DEB4170569DCFEC63E3897C71E3AE8642FD3139A554AEE21E",
			Running:    true,
		},
	}}

	mux.HandleFunc("/accounts/"+accountID+"/devices/posture", handler)

	actual, _, err := client.DevicePostureRules(context.Background(), accountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDevicePostureRule(t *testing.T) {
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
				"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
				"schedule": "1h",
				"type": "file",
				"name": "My rule name",
				"description": "My description",
				"match": [
					{
						"platform": "ios"
					}
				],
				"input": {
					"id": "9e597887-345e-4a32-a09c-68811b129768"
				}
			}
		}
		`)
	}

	want := DevicePostureRule{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "file",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input:       DevicePostureRuleInput{ID: "9e597887-345e-4a32-a09c-68811b129768"},
	}

	mux.HandleFunc("/accounts/"+accountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.DevicePostureRule(context.Background(), accountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateDevicePostureRule(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
				"schedule": "1h",
				"type": "file",
				"name": "My rule name",
				"description": "My description",
				"match": [
					{
						"platform": "ios"
					}
				],
				"input": {
					"id": "9e597887-345e-4a32-a09c-68811b129768"
				}
			}
		}
		`)
	}

	rule := DevicePostureRule{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "file",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input:       DevicePostureRuleInput{ID: "9e597887-345e-4a32-a09c-68811b129768"},
	}

	mux.HandleFunc("/accounts/"+accountID+"/devices/posture", handler)

	actual, err := client.CreateDevicePostureRule(context.Background(), accountID, DevicePostureRule{
		Name:        "My rule name",
		Description: "My description",
		Type:        "file",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input:       DevicePostureRuleInput{ID: "9e597887-345e-4a32-a09c-68811b129768"},
	})

	if assert.NoError(t, err) {
		assert.Equal(t, rule, actual)
	}
}

func TestUpdateDevicePostureRule(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
				"schedule": "1h",
				"type": "file",
				"name": "My rule name",
				"description": "My description",
				"match": [
					{
						"platform": "ios"
					}
				],
				"input": {
					"id": "9e597887-345e-4a32-a09c-68811b129768"
				}
			}
		}
		`)
	}

	rule := DevicePostureRule{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "file",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input:       DevicePostureRuleInput{ID: "9e597887-345e-4a32-a09c-68811b129768"},
	}

	mux.HandleFunc("/accounts/"+accountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.UpdateDevicePostureRule(context.Background(), accountID, rule)

	if assert.NoError(t, err) {
		assert.Equal(t, rule, actual)
	}
}

func TestUpdateDevicePostureRuleWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateDevicePostureRule(context.Background(), zoneID, DevicePostureRule{})
	assert.EqualError(t, err, "device posture rule ID cannot be empty")
}

func TestDeleteDevicePostureRule(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
      "success": true,
      "errors": [],
      "messages": [],
      "result": {
        "id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db"
      }
    }
    `)
	}

	mux.HandleFunc("/accounts/"+accountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)
	err := client.DeleteDevicePostureRule(context.Background(), accountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	assert.NoError(t, err)
}
