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

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture", handler)

	actual, _, err := client.DevicePostureRules(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDevicePostureFileRule(t *testing.T) {
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
					"path": "/tmp/test",
					"exists": true,
					"sha256": "42b4daec3962691f5893a966245e5ea30f9f8df7254e7b7af43a171e3e29c857"
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
		Input: DevicePostureRuleInput{
			Path:   "/tmp/test",
			Exists: true,
			Sha256: "42b4daec3962691f5893a966245e5ea30f9f8df7254e7b7af43a171e3e29c857",
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.DevicePostureRule(context.Background(), testAccountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDevicePostureDiskEncryptionRule(t *testing.T) {
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
				"type": "disk_encryption",
				"name": "My rule name",
				"description": "My description",
				"match": [
					{
						"platform": "ios"
					}
				],
				"input": {
					"requireAll": true
				}
			}
		}
		`)
	}

	want := DevicePostureRule{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "disk_encryption",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input: DevicePostureRuleInput{
			RequireAll: true,
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.DevicePostureRule(context.Background(), testAccountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDevicePostureOsVersionRule(t *testing.T) {
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
				"type": "os_version",
				"name": "My rule name",
				"description": "My description",
				"match": [
					{
						"platform": "ios"
					}
				],
				"input": {
					"version": "10.0.1",
					"operator": ">="
				}
			}
		}
		`)
	}

	want := DevicePostureRule{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "os_version",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input: DevicePostureRuleInput{
			Version:  "10.0.1",
			Operator: ">=",
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.DevicePostureRule(context.Background(), testAccountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDevicePostureDomainJoinedRule(t *testing.T) {
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
				"type": "domain_joined",
				"name": "My rule name",
				"description": "My description",
				"match": [
					{
						"platform": "ios"
					}
				],
				"input": {
					"domain": "example.com"
				}
			}
		}
		`)
	}

	want := DevicePostureRule{
		ID:          "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:        "My rule name",
		Description: "My description",
		Type:        "domain_joined",
		Schedule:    "1h",
		Match:       []DevicePostureRuleMatch{{Platform: "ios"}},
		Input: DevicePostureRuleInput{
			Domain: "example.com",
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.DevicePostureRule(context.Background(), testAccountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

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

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture", handler)

	actual, err := client.CreateDevicePostureRule(context.Background(), testAccountID, DevicePostureRule{
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

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.UpdateDevicePostureRule(context.Background(), testAccountID, rule)

	if assert.NoError(t, err) {
		assert.Equal(t, rule, actual)
	}
}

func TestUpdateDevicePostureRuleWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateDevicePostureRule(context.Background(), testZoneID, DevicePostureRule{})
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

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/posture/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)
	err := client.DeleteDevicePostureRule(context.Background(), testAccountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	assert.NoError(t, err)
}
