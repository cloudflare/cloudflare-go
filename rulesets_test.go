package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListRulesets(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": [
        {
          "id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
          "name": "my example ruleset",
          "description": "Test magic transit ruleset",
          "kind": "root",
          "version": "1",
          "last_updated": "2020-12-02T20:24:07.776073Z",
          "phase": "magic_transit"
        }
      ],
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets", handler)
	mux.HandleFunc("/zones/"+testZoneID+"/rulesets", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	want := []Ruleset{
		{
			ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
			Name:        "my example ruleset",
			Description: "Test magic transit ruleset",
			Kind:        "root",
			Version:     "1",
			LastUpdated: &lastUpdated,
			Phase:       RulesetPhaseMagicTransit,
		},
	}

	zoneActual, err := client.ListZoneRulesets(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}

	accountActual, err := client.ListAccountRulesets(context.Background(), testAccountID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, accountActual)
	}
}

func TestGetRuleset_MagicTransit(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": {
        "id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
        "name": "my example ruleset",
        "description": "Test magic transit ruleset",
        "kind": "root",
        "version": "1",
        "last_updated": "2020-12-02T20:24:07.776073Z",
        "phase": "magic_transit"
      },
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)
	mux.HandleFunc("/zones/"+testZoneID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	want := Ruleset{
		ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:        "my example ruleset",
		Description: "Test magic transit ruleset",
		Kind:        "root",
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       RulesetPhaseMagicTransit,
	}

	zoneActual, err := client.GetZoneRuleset(context.Background(), testZoneID, "2c0fc9fa937b11eaa1b71c4d701ab86e")
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}

	accountActual, err := client.GetAccountRuleset(context.Background(), testAccountID, "2c0fc9fa937b11eaa1b71c4d701ab86e")
	if assert.NoError(t, err) {
		assert.Equal(t, want, accountActual)
	}
}

func TestGetRuleset_WAF(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": {
        "id": "70339d97bdb34195bbf054b1ebe81f76",
        "name": "Cloudflare Normalization Ruleset",
        "description": "Created by the Cloudflare security team, this ruleset provides normalization on the URL path",
        "kind": "managed",
        "version": "1",
        "rules": [
          {
            "id": "78723a9e0c7c4c6dbec5684cb766231d",
            "version": "1",
            "action": "rewrite",
            "action_parameters": {
              "uri": {
                "path": {
                  "expression": "normalize_url_path(raw.http.request.uri.path)"
                },
                "origin": false
              }
            },
            "description": "Normalization on the URL path, without propagating it to the origin",
            "last_updated": "2020-12-18T09:28:09.655749Z",
            "ref": "272936dc447b41fe976255ff6b768ec0",
            "enabled": true
          }
        ],
        "last_updated": "2020-12-18T09:28:09.655749Z",
        "phase": "http_request_sanitize"
      },
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets/b232b534beea4e00a21dcbb7a8a545e9", handler)
	mux.HandleFunc("/zones/"+testZoneID+"/rulesets/b232b534beea4e00a21dcbb7a8a545e9", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-18T09:28:09.655749Z")

	rules := []RulesetRule{{
		ID:      "78723a9e0c7c4c6dbec5684cb766231d",
		Version: "1",
		Action:  RulesetRuleActionRewrite,
		ActionParameters: &RulesetRuleActionParameters{
			URI: RulesetRuleActionParametersURI{
				Path: RulesetRuleActionParametersURIPath{
					Expression: "normalize_url_path(raw.http.request.uri.path)",
				},
				Origin: false,
			},
		},
		Description: "Normalization on the URL path, without propagating it to the origin",
		LastUpdated: &lastUpdated,
		Ref:         "272936dc447b41fe976255ff6b768ec0",
		Enabled:     true,
	}}

	want := Ruleset{
		ID:          "70339d97bdb34195bbf054b1ebe81f76",
		Name:        "Cloudflare Normalization Ruleset",
		Description: "Created by the Cloudflare security team, this ruleset provides normalization on the URL path",
		Kind:        RulesetKindManaged,
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       RulesetPhaseHTTPRequestSanitize,
		Rules:       rules,
	}

	zoneActual, err := client.GetZoneRuleset(context.Background(), testZoneID, "b232b534beea4e00a21dcbb7a8a545e9")
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}

	accountActual, err := client.GetAccountRuleset(context.Background(), testAccountID, "b232b534beea4e00a21dcbb7a8a545e9")
	if assert.NoError(t, err) {
		assert.Equal(t, want, accountActual)
	}
}

func TestCreateRuleset(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": {
        "id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
        "name": "my example ruleset",
        "description": "Test magic transit ruleset",
        "kind": "root",
        "version": "1",
        "last_updated": "2020-12-02T20:24:07.776073Z",
        "phase": "magic_transit",
        "rules": [
          {
            "id": "62449e2e0de149619edb35e59c10d801",
            "version": "1",
            "action": "skip",
            "action_parameters":{
              "ruleset":"current"
            },
            "expression": "tcp.dstport in { 32768..65535 }",
            "description": "Allow TCP Ephemeral Ports",
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
	mux.HandleFunc("/zones/"+testZoneID+"/rulesets", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	rules := []RulesetRule{{
		ID:      "62449e2e0de149619edb35e59c10d801",
		Version: "1",
		Action:  RulesetRuleActionSkip,
		ActionParameters: &RulesetRuleActionParameters{
			Ruleset: "current",
		},
		Expression:  "tcp.dstport in { 32768..65535 }",
		Description: "Allow TCP Ephemeral Ports",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}}

	newRuleset := Ruleset{
		Name:        "my example ruleset",
		Description: "Test magic transit ruleset",
		Kind:        "root",
		Phase:       RulesetPhaseMagicTransit,
		Rules:       rules,
	}

	want := Ruleset{
		ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:        "my example ruleset",
		Description: "Test magic transit ruleset",
		Kind:        "root",
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       RulesetPhaseMagicTransit,
		Rules:       rules,
	}

	zoneActual, err := client.CreateZoneRuleset(context.Background(), testZoneID, newRuleset)
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}

	accountActual, err := client.CreateAccountRuleset(context.Background(), testAccountID, newRuleset)
	if assert.NoError(t, err) {
		assert.Equal(t, want, accountActual)
	}
}

func TestDeleteRuleset(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, ``)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)
	mux.HandleFunc("/zones/"+testZoneID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	zErr := client.DeleteZoneRuleset(context.Background(), testZoneID, "2c0fc9fa937b11eaa1b71c4d701ab86e")
	assert.NoError(t, zErr)

	aErr := client.DeleteAccountRuleset(context.Background(), testAccountID, "2c0fc9fa937b11eaa1b71c4d701ab86e")
	assert.NoError(t, aErr)
}

func TestUpdateRuleset(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": {
        "id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
        "name": "ruleset1",
        "description": "Test Firewall Ruleset Update",
        "kind": "root",
        "version": "1",
        "last_updated": "2020-12-02T20:24:07.776073Z",
        "phase": "magic_transit",
        "rules": [
          {
            "id": "62449e2e0de149619edb35e59c10d801",
            "version": "1",
            "action": "skip",
            "action_parameters":{
              "ruleset":"current"
            },
            "expression": "tcp.dstport in { 32768..65535 }",
            "description": "Allow TCP Ephemeral Ports",
            "last_updated": "2020-12-02T20:24:07.776073Z",
            "ref": "72449e2e0de149619edb35e59c10d801",
            "enabled": true
          },
          {
            "id": "62449e2e0de149619edb35e59c10d802",
            "version": "1",
            "action": "skip",
            "action_parameters":{
              "ruleset":"current"
            },
            "expression": "udp.dstport in { 32768..65535 }",
            "description": "Allow UDP Ephemeral Ports",
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
	mux.HandleFunc("/zones/"+testZoneID+"/rulesets/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	lastUpdated, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	rules := []RulesetRule{{
		ID:      "62449e2e0de149619edb35e59c10d801",
		Version: "1",
		Action:  RulesetRuleActionSkip,
		ActionParameters: &RulesetRuleActionParameters{
			Ruleset: "current",
		},
		Expression:  "tcp.dstport in { 32768..65535 }",
		Description: "Allow TCP Ephemeral Ports",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}, {
		ID:      "62449e2e0de149619edb35e59c10d802",
		Version: "1",
		Action:  RulesetRuleActionSkip,
		ActionParameters: &RulesetRuleActionParameters{
			Ruleset: "current",
		},
		Expression:  "udp.dstport in { 32768..65535 }",
		Description: "Allow UDP Ephemeral Ports",
		LastUpdated: &lastUpdated,
		Ref:         "72449e2e0de149619edb35e59c10d801",
		Enabled:     true,
	}}

	want := Ruleset{
		ID:          "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:        "ruleset1",
		Description: "Test Firewall Ruleset Update",
		Kind:        "root",
		Version:     "1",
		LastUpdated: &lastUpdated,
		Phase:       RulesetPhaseMagicTransit,
		Rules:       rules,
	}

	zoneActual, err := client.UpdateZoneRuleset(context.Background(), testZoneID, "2c0fc9fa937b11eaa1b71c4d701ab86e", "Test Firewall Ruleset Update", rules)
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}

	accountActual, err := client.UpdateAccountRuleset(context.Background(), testAccountID, "2c0fc9fa937b11eaa1b71c4d701ab86e", "Test Firewall Ruleset Update", rules)
	if assert.NoError(t, err) {
		assert.Equal(t, want, accountActual)
	}
}
