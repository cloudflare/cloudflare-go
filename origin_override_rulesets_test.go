package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListOriginOverrideRulesets(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": [
				{
					"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
					"name": "ruleset1",
					"description": "Test Origin Override Ruleset",
					"kind": "root",
					"version": "1",
					"last_updated": "2020-12-02T20:24:07.776073Z",
					"phase": "http_request_origin"
				}
			],
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	want := []OriginOverrideRuleset{
		{
			ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
			Name:        "ruleset1",
			Description: "Test Origin Override Ruleset",
			Kind:        "root",
			Version:     "1",
			LastUpdated: &lastUpdated,
			Phase:       OriginOverrideRulesetPhaseOrigin,
		},
	}

	actual, err := client.ListOriginOverrideRulesets(context.Background(), testAccountID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetOriginOverrideRuleset(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"name": "ruleset1",
				"description": "Test Origin Override Ruleset",
				"kind": "root",
				"version": "1",
				"last_updated": "2020-12-02T20:24:07.776073Z",
				"phase": "http_request_origin",
				"rules": [
					{
						"id": "62449e2e0de149619edb35e59c10d801",
						"version": "1",
						"action": "route",
						"action_parameters":{
							"host_header": "my_host",
							"origin": {
								"host": "test.host.com",
								"port": 80,
							}
						},
						"expression": "http.host eq \"fake.net\"",
						"description": "Change origin destination and header",
						"last_updated": "2020-12-02T20:24:07.776073Z",
						"ref": "72449e2e0de149619edb35e59c10d801",
						"enabled": true
					}
				]
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")
	rules := []OriginOverrideRulesetRule{{
		ID:      "62449e2e0de149619edb35e59c10d801",
		Version: "1",
		Action:  OriginOverrideRulesetRuleActionRoute,
		ActionParameters: &OriginOverrideRulesetRuleActionParameters{
			HostHeader: "test.host.com",
			Origin: OriginOverrideRulesetRuleActionOriginParameters{
				Host: "test.host.com",
				Port: 80,
			},
		},
		Expression:  "http.host eq \"fake.net\"",
		Description: "Change origin destination and header",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}}

	want := OriginOverrideRuleset{
		ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:        "ruleset1",
		Description: "Test Origin Override Ruleset",
		Kind:        "root",
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       OriginOverrideRulesetPhaseOrigin,
		Rules:       rules,
	}

	actual, err := client.GetOriginOverrideRuleset(context.Background(), testAccountID, "2c0fc9fa937b11eaa1b71c4d701ab86e")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateOriginOverrideRuleset(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"name": "ruleset1",
				"description": "Test Origin Override Ruleset",
				"kind": "root",
				"version": "1",
				"last_updated": "2020-12-02T20:24:07.776073Z",
				"phase": "http_request_origin",
				"rules": [
					{
						"id": "62449e2e0de149619edb35e59c10d801",
						"version": "1",
						"action": "route",
						"action_parameters":{
							"host_header": "my_host",
							"origin": {
								"host": "test.host.com",
								"port": 80,
							}
						},
						"expression": "http.host eq \"fake.net\"",
						"description": "Change origin destination and header",
						"last_updated": "2020-12-02T20:24:07.776073Z",
						"ref": "72449e2e0de149619edb35e59c10d801",
						"enabled": true
					}
				]
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")
	rules := []OriginOverrideRulesetRule{{
		ID:      "62449e2e0de149619edb35e59c10d801",
		Version: "1",
		Action:  OriginOverrideRulesetRuleActionRoute,
		ActionParameters: &OriginOverrideRulesetRuleActionParameters{
			HostHeader: "my_host",
			Origin: OriginOverrideRulesetRuleActionOriginParameters{
				Host: "test.host.com",
				Port: 80,
			},
		},
		Expression:  "http.host eq \"fake.net\"",
		Description: "Change origin destination and header",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}}

	want := OriginOverrideRuleset{
		ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:        "ruleset1",
		Description: "Test Origin Override Ruleset",
		Kind:        "root",
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       OriginOverrideRulesetPhaseOrigin,
		Rules:       rules,
	}

	actual, err := client.CreateOriginOverrideRuleset(context.Background(), testAccountID, "ruleset1", "Test Origin Override Ruleset", rules)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateOriginOverrideRuleset(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"name": "ruleset1",
				"description": "Origin Override Ruleset Update",
				"kind": "root",
				"version": "1",
				"last_updated": "2020-12-02T20:24:07.776073Z",
				"phase": "http_request_origin",
				"rules": [
					{
						"id": "62449e2e0de149619edb35e59c10d801",
						"version": "1",
						"action": "route",
						"action_parameters":{
							"host_header": "my_host",
							"origin": {
								"host": "test.host.com",
								"port": 80,
							}
						},
						"expression": "http.host eq \"fake.net\"",
						"description": "Change origin destination and header",
						"last_updated": "2020-12-02T20:24:07.776073Z",
						"ref": "72449e2e0de149619edb35e59c10d801",
						"enabled": true
					},
					{
						"id": "62449e2e0de149619edb35e59c10d802",
						"version": "1",
						"action": "route",
						"action_parameters":{
							"host_header": "my_host2",
							"origin": {
								"host": "test2.host.com",
								"port": 80,
							}
						},
						"expression": "http.host eq \"fake2.net\"",
						"description": "Change origin destination and header",
						"last_updated": "2020-12-02T20:24:07.776073Z",
						"ref": "72449e2e0de149619edb35e59c10d801",
						"enabled": true
					}
				]
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")
	rules := []OriginOverrideRulesetRule{{
		ID:      "62449e2e0de149619edb35e59c10d801",
		Version: "1",
		Action:  OriginOverrideRulesetRuleActionRoute,
		ActionParameters: &OriginOverrideRulesetRuleActionParameters{
			HostHeader: "my_host",
			Origin: OriginOverrideRulesetRuleActionOriginParameters{
				Host: "test.host.com",
				Port: 80,
			},
		},
		Expression:  "http.host eq \"fake.net\"",
		Description: "Change origin destination and header",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}, {
		ID:      "62449e2e0de149619edb35e59c10d802",
		Version: "1",
		Action:  OriginOverrideRulesetRuleActionRoute,
		ActionParameters: &OriginOverrideRulesetRuleActionParameters{
			HostHeader: "my_host2",
			Origin: OriginOverrideRulesetRuleActionOriginParameters{
				Host: "test2.host.com",
				Port: 80,
			},
		},
		Expression:  "http.host eq \"fake2.net\"",
		Description: "Change origin destination and header",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}}

	want := OriginOverrideRuleset{
		ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:        "ruleset1",
		Description: "Test Origin Override Ruleset Update",
		Kind:        "root",
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       OriginOverrideRulesetPhaseOrigin,
		Rules:       rules,
	}

	actual, err := client.UpdateOriginOverrideRuleset(context.Background(), testAccountID, "2c0fc9fa937b11eaa1b71c4d701ab86e", "Test Origin Override Ruleset Update", rules)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

// Firewall API is not implementing the standard response blob but returns an empty response (204) in case
// of a success. So we are checking for the response body size here
// TODO, This is going to be changed by MFW-63.
func TestDeleteOriginOverrideRuleset(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, ``)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	err := client.DeleteOriginOverrideRuleset(context.Background(), testAccountID, "2c0fc9fa937b11eaa1b71c4d701ab86e")
	assert.NoError(t, err)
}
