package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTeamsRules(t *testing.T) {
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
				  "id": "7559a944-3dd7-41bf-b183-360a814a8c36",
				  "name": "rule1",
				  "description": "rule description",
				  "precedence": 1000,
				  "enabled": false,
				  "action": "isolate",
				  "filters": [
					"http"
				  ],
				  "created_at": "2014-01-01T05:20:00.12345Z",
				  "updated_at": "2014-01-01T05:20:00.12345Z",
				  "deleted_at": null,
				  "traffic": "http.host == \"example.com\"",
				  "identity": "",
				  "version": 1,
				  "rule_settings": {
					"block_page_enabled": false,
					"block_reason": "",
					"override_ips": null,
					"override_host": "",
					"l4override": null,
					"biso_admin_controls": null,
					"add_headers": null,
					"check_session": {
						"enforce": true,
						"duration": "15m0s"
					},
                    "insecure_disable_dnssec_validation": false
				  }
				},
				{
				  "id": "9ae57318-f32e-46b3-b889-48dd6dcc49af",
				  "name": "rule2",
				  "description": "rule description 2",
				  "precedence": 2000,
				  "enabled": true,
				  "action": "block",
				  "filters": [
					"http"
				  ],
				  "created_at": "2014-01-01T05:20:00.12345Z",
				  "updated_at": "2014-01-01T05:20:00.12345Z",
				  "deleted_at": null,
				  "traffic": "http.host == \"abcd.com\"",
				  "identity": "",
				  "version": 1,
				  "rule_settings": {
					"block_page_enabled": true,
					"block_reason": "",
					"override_ips": null,
					"override_host": "",
					"l4override": null,
					"biso_admin_controls": null,
					"add_headers": null,
					"check_session": null,
                    "insecure_disable_dnssec_validation": true
				  }
				}
			]
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := []TeamsRule{{
		ID:            "7559a944-3dd7-41bf-b183-360a814a8c36",
		Name:          "rule1",
		Description:   "rule description",
		Precedence:    1000,
		Enabled:       false,
		Action:        Isolate,
		Filters:       []TeamsFilterType{HttpFilter},
		Traffic:       `http.host == "example.com"`,
		DevicePosture: "",
		Identity:      "",
		Version:       1,
		RuleSettings: TeamsRuleSettings{
			BlockPageEnabled:  false,
			BlockReason:       "",
			OverrideIPs:       nil,
			OverrideHost:      "",
			L4Override:        nil,
			AddHeaders:        nil,
			BISOAdminControls: nil,
			CheckSession: &TeamsCheckSessionSettings{
				Enforce:  true,
				Duration: Duration{900 * time.Second},
			},
			InsecureDisableDNSSECValidation: false,
		},
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: nil,
	},
		{
			ID:            "9ae57318-f32e-46b3-b889-48dd6dcc49af",
			Name:          "rule2",
			Description:   "rule description 2",
			Precedence:    2000,
			Enabled:       true,
			Action:        Block,
			Filters:       []TeamsFilterType{HttpFilter},
			Traffic:       `http.host == "abcd.com"`,
			Identity:      "",
			DevicePosture: "",
			Version:       1,
			RuleSettings: TeamsRuleSettings{
				BlockPageEnabled:  true,
				BlockReason:       "",
				OverrideIPs:       nil,
				OverrideHost:      "",
				L4Override:        nil,
				AddHeaders:        nil,
				BISOAdminControls: nil,
				CheckSession:      nil,
				// setting is invalid for block rules, just testing serialization here
				InsecureDisableDNSSECValidation: true,
			},
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
			DeletedAt: nil,
		}}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/rules", handler)

	actual, err := client.TeamsRules(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestTeamsRule(t *testing.T) {
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
				"id": "7559a944-3dd7-41bf-b183-360a814a8c36",
				"name": "rule1",
				"description": "rule description",
				"precedence": 1000,
				"enabled": false,
				"action": "isolate",
				"filters": [
					"http"
				],
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"deleted_at": null,
				"traffic": "http.host == \"abcd.com\"",
				"identity": "",
				"version": 1,
				"rule_settings": {
					"block_page_enabled": false,
					"block_reason": "",
					"override_ips": null,
					"override_host": "",
					"l4override": null,
					"biso_admin_controls": null,
					"add_headers": null,
					"check_session": {
						"enforce": true,
						"duration": "15m0s"
					},
                    "insecure_disable_dnssec_validation": false
				}
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := TeamsRule{
		ID:            "7559a944-3dd7-41bf-b183-360a814a8c36",
		Name:          "rule1",
		Description:   "rule description",
		Precedence:    1000,
		Enabled:       false,
		Action:        Isolate,
		Filters:       []TeamsFilterType{HttpFilter},
		Traffic:       `http.host == "abcd.com"`,
		Identity:      "",
		DevicePosture: "",
		Version:       1,
		RuleSettings: TeamsRuleSettings{
			BlockPageEnabled:  false,
			BlockReason:       "",
			OverrideIPs:       nil,
			OverrideHost:      "",
			L4Override:        nil,
			AddHeaders:        nil,
			BISOAdminControls: nil,
			CheckSession: &TeamsCheckSessionSettings{
				Enforce:  true,
				Duration: Duration{900 * time.Second},
			},
			InsecureDisableDNSSECValidation: false,
		},
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: nil,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/rules/7559a9443dd741bfb183360a814a8c36", handler)

	actual, err := client.TeamsRule(context.Background(), testAccountID, "7559a9443dd741bfb183360a814a8c36")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestTeamsCreateRule(t *testing.T) {
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
				"name": "rule1",
				"description": "rule description",
				"precedence": 1000,
				"enabled": false,
				"action": "isolate",
				"filters": [
					"http"
				],
				"traffic": "http.host == \"abcd.com\"",
				"identity": "",
				"rule_settings": {
					"block_page_enabled": false,
					"block_reason": "",
					"override_ips": null,
					"override_host": "",
					"l4override": null,
					"biso_admin_controls": null,
					"add_headers": {
						"X-Test": ["abcd"]
					},
					"check_session": {
						"enforce": true,
						"duration": "5m0s"
					},
                    "insecure_disable_dnssec_validation": false
				}
			}
		}
		`)
	}

	want := TeamsRule{
		Name:          "rule1",
		Description:   "rule description",
		Precedence:    1000,
		Enabled:       false,
		Action:        Isolate,
		Filters:       []TeamsFilterType{HttpFilter},
		Traffic:       `http.host == "abcd.com"`,
		Identity:      "",
		DevicePosture: "",
		RuleSettings: TeamsRuleSettings{
			BlockPageEnabled:  false,
			BlockReason:       "",
			OverrideIPs:       nil,
			OverrideHost:      "",
			L4Override:        nil,
			AddHeaders:        http.Header{"X-Test": []string{"abcd"}},
			BISOAdminControls: nil,
			CheckSession: &TeamsCheckSessionSettings{
				Enforce:  true,
				Duration: Duration{300 * time.Second},
			},
			InsecureDisableDNSSECValidation: false,
		},
		DeletedAt: nil,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/rules", handler)

	actual, err := client.TeamsCreateRule(context.Background(), testAccountID, want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestTeamsUpdateRule(t *testing.T) {
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
				"id": "7559a944-3dd7-41bf-b183-360a814a8c36",
				"name": "rule_name_change",
				"description": "rule new description",
				"precedence": 3000,
				"enabled": true,
				"action": "block",
				"filters": [
					"http"
				],
				"traffic": "",
				"identity": "",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-02-01T05:20:00.12345Z",
				"rule_settings": {
					"block_page_enabled": false,
					"block_reason": "",
					"override_ips": null,
					"override_host": "",
					"l4override": null,
					"biso_admin_controls": null,
					"add_headers": null,
					"check_session": null,
                    "insecure_disable_dnssec_validation": false
				}
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")

	want := TeamsRule{
		ID:            "7559a944-3dd7-41bf-b183-360a814a8c36",
		Name:          "rule_name_change",
		Description:   "rule new description",
		Precedence:    3000,
		Enabled:       true,
		Action:        Block,
		Filters:       []TeamsFilterType{HttpFilter},
		Traffic:       "",
		Identity:      "",
		DevicePosture: "",
		RuleSettings: TeamsRuleSettings{
			BlockPageEnabled:                false,
			BlockReason:                     "",
			OverrideIPs:                     nil,
			OverrideHost:                    "",
			L4Override:                      nil,
			AddHeaders:                      nil,
			BISOAdminControls:               nil,
			CheckSession:                    nil,
			InsecureDisableDNSSECValidation: false,
		},
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: nil,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/rules/7559a9443dd741bfb183360a814a8c36", handler)

	actual, err := client.TeamsUpdateRule(context.Background(), testAccountID, "7559a9443dd741bfb183360a814a8c36", want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestTeamsPatchRule(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'Patch', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"name": "rule_name_change",
				"description": "rule new description",
				"precedence": 3000,
				"enabled": true,
				"action": "block",
				"rule_settings": {
					"block_page_enabled": false,
					"block_reason": "",
					"override_ips": null,
					"override_host": "",
					"l4override": null,
					"biso_admin_controls": null,
					"add_headers": null,
					"check_session": null,
                    "insecure_disable_dnssec_validation": false
				}
			}
		}
		`)
	}

	reqBody := TeamsRulePatchRequest{
		Name:        "rule_name_change",
		Description: "rule new description",
		Precedence:  3000,
		Enabled:     true,
		Action:      Block,
		RuleSettings: TeamsRuleSettings{
			BlockPageEnabled:                false,
			BlockReason:                     "",
			OverrideIPs:                     nil,
			OverrideHost:                    "",
			L4Override:                      nil,
			AddHeaders:                      nil,
			BISOAdminControls:               nil,
			CheckSession:                    nil,
			InsecureDisableDNSSECValidation: false,
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/rules/7559a9443dd741bfb183360a814a8c36", handler)

	actual, err := client.TeamsPatchRule(context.Background(), testAccountID, "7559a9443dd741bfb183360a814a8c36", reqBody)

	if assert.NoError(t, err) {
		assert.Equal(t, actual.Name, "rule_name_change")
		assert.Equal(t, actual.Description, "rule new description")
		assert.Equal(t, actual.Precedence, uint64(3000))
		assert.Equal(t, actual.Action, Block)
		assert.Equal(t, actual.Enabled, true)
	}
}

func TestTeamsDeleteRule(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'Delete', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": nil
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/rules/7559a9443dd741bfb183360a814a8c36", handler)

	err := client.TeamsDeleteRule(context.Background(), testAccountID, "7559a9443dd741bfb183360a814a8c36")

	assert.NoError(t, err)
}
