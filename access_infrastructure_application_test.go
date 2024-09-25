package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testInfrastructureApplicationId = "019205b5-97d7-7272-b00e-0ea05e61a124"

var (
	accessPolicyCreatedOn, _              = time.Parse(time.RFC3339, "2020-08-25T05:00:22Z")
	accessPolicyUpdatedOn, _              = time.Parse(time.RFC3339, "2020-08-25T05:00:22Z")
	infrastructureApplicationCreatedOn, _ = time.Parse(time.RFC3339, "2024-08-25T05:00:22Z")
	infrastructureApplicationUpdatedOn, _ = time.Parse(time.RFC3339, "2024-08-25T14:30:22Z")
	expectedInfraAppPolicy                = AccessPolicy{
		ID:         "699d98642c564d2e855e9661899b7252",
		Precedence: 1,
		Reusable:   BoolPtr(true),
		Decision:   "allow",
		CreatedAt:  &accessPolicyCreatedOn,
		UpdatedAt:  &accessPolicyUpdatedOn,
		Name:       "my infra SSH policy",
		Include: []any{
			map[string]interface{}{"email": map[string]interface{}{"email": "devuser@example.com"}},
		},
		InfraConnectionRules: &InfraConnectionRules{
			SSH: &InfraConnectionRulesSSH{
				Usernames: []string{"devuser"},
			},
		},
	}
	expectedInfrastructureApplication = InfrastructureApplication{
		ID:        testInfrastructureApplicationId,
		Name:      "infrastructure_application_test",
		Type:      "infrastructure",
		Aud:       "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		CreatedAt: infrastructureAccessTargetCreatedOn.String(),
		UpdatedAt: infrastructureApplicationCreatedOn.String(),
		TargetContexts: []InfraTargetContext{
			{
				TargetAttributes: map[string][]string{
					"name": {"vault 420", "vault 421"},
					"env":  {"dev", "staging"},
				},
				Port:     24,
				Protocol: "SSH",
			},
		},
		Policies: []AccessPolicy{
			expectedInfraAppPolicy,
		},
	}

	expectedInfrastructureApplicationModified = InfrastructureApplication{
		ID:        testInfrastructureApplicationId,
		Name:      "infrastructure_application_test_updated",
		Type:      "infrastructure",
		Aud:       "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		CreatedAt: infrastructureAccessTargetCreatedOn.String(),
		UpdatedAt: infrastructureApplicationUpdatedOn.String(),
		TargetContexts: []InfraTargetContext{
			{
				TargetAttributes: map[string][]string{
					"name": {"vault 425", "vault 430"},
					"env":  {"staging"},
				},
				Port:     24,
				Protocol: "SSH",
			},
		},
		Policies: []AccessPolicy{
			expectedInfraAppPolicy,
		},
	}
)

func TestInfrastructureApplication_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps?app_type=infrastructure", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "019205b5-97d7-7272-b00e-0ea05e61a124",
					"created_at": "2024-08-25T05:00:22Z",
					"updated_at": "2024-08-25T05:00:22Z",
					"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
					"name": "infrastructure_application_test",
					"domain": "test.example.com/admin",
					"type": "infrastructure",
					"session_duration": "24h",
					"auto_redirect_to_identity": false,
					"enable_binding_cookie": false,
					"custom_deny_url": "https://www.example.com",
					"custom_non_identity_deny_url": "https://blocked.com",
					"custom_deny_message": "denied!",
					"http_only_cookie_attribute": true,
					"same_site_cookie_attribute": "strict",
					"logo_url": "https://www.example.com/example.png",
					"skip_interstitial": true,
					"app_launcher_visible": true,
					"service_auth_401_redirect": true,
					"path_cookie_attribute": true,
					"custom_pages": ["480f4f69-1a28-4fdd-9240-1ed29f0ac1dc"],
					"allow_authenticate_via_warp": true,
					"options_preflight_bypass": false,
					"target_criteria": [
						{
							"target_attributes": {
								"name": [ "vault 420", "vault 421" ],
								"env": [ "dev", "staging" ]
							},
							"port": 24,
							"protocol": "SSH"
						}
					], 
					"policies": [
						{
							"id": "699d98642c564d2e855e9661899b7252",
							"precedence": 1,
							"reusable": true,
							"decision": "allow",
							"created_at": "2020-08-25T05:00:22Z",
							"updated_at": "2020-08-25T05:00:22Z",
							"name": "my infra SSH policy",
							"include": [
								{
									"email": {
										"email": "devuser@example.com"
									}
								}
							],
							"connection_rules": {
								"ssh": {
									"usernames": ["devuser"]
								}
							}
						}
					]
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 1
			}
		}`)
	})

	_, _, err := client.ListAccessInfrastructureApplications(context.Background(), AccountIdentifier(""), ListAccessApplicationsParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, results, err := client.ListAccessInfrastructureApplications(context.Background(), AccountIdentifier(testAccountID), ListAccessApplicationsParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(out), "expected 1 infrastructure_application")
		assert.Equal(t, 20, results.PerPage, "expected 20 per page")
		assert.Equal(t, expectedInfrastructureApplication, out[0], "list infrastructure_application structs not equal")
	}
}

func TestInfrastructureApplication_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "019205b5-97d7-7272-b00e-0ea05e61a124",
				"created_at": "2024-08-25T05:00:22Z",
				"updated_at": "2024-08-25T05:00:22Z",
				"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
				"name": "infrastructure_application_test",
				"domain": "test.example.com/admin",
				"type": "infrastructure",
				"session_duration": "24h",
				"auto_redirect_to_identity": false,
				"enable_binding_cookie": false,
				"custom_deny_url": "https://www.example.com",
				"custom_non_identity_deny_url": "https://blocked.com",
				"custom_deny_message": "denied!",
				"http_only_cookie_attribute": true,
				"same_site_cookie_attribute": "strict",
				"logo_url": "https://www.example.com/example.png",
				"skip_interstitial": true,
				"app_launcher_visible": true,
				"service_auth_401_redirect": true,
				"path_cookie_attribute": true,
				"custom_pages": ["480f4f69-1a28-4fdd-9240-1ed29f0ac1dc"],
				"allow_authenticate_via_warp": true,
				"options_preflight_bypass": false,
				"target_criteria": [
					{
						"target_attributes": {
							"name": [ "vault 420", "vault 421" ],
							"env": [ "dev", "staging" ]
						},
						"port": 24,
						"protocol": "SSH"
					}
				], 
				"policies": [
					{
						"id": "699d98642c564d2e855e9661899b7252",
						"precedence": 1,
						"reusable": true,
						"decision": "allow",
						"created_at": "2020-08-25T05:00:22Z",
						"updated_at": "2020-08-25T05:00:22Z",
						"name": "my infra SSH policy",
						"include": [
							{
								"email": {
									"email": "devuser@example.com"
								}
							}
						],
						"connection_rules": {
							"ssh": {
								"usernames": ["devuser"]
							}
						}
					}
				]
			}
		}`)
	})

	// Make sure missing account ID is thrown
	_, err := client.CreateAccessInfrastructureApplication(context.Background(), AccountIdentifier(""), CreateInfrastructureApplicationParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	// We expect this to parse through the Access Application response and only return
	// Infrastructure Application built from that response
	out, err := client.CreateAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), CreateInfrastructureApplicationParams{
		InfrastructureApplicationParams: InfrastructureApplicationParams{
			Name: "infrastructure_application_test",
			Type: "infrastructure",
			TargetContexts: []InfraTargetContext{
				{
					TargetAttributes: map[string][]string{
						"name": {"vault 420", "vault 421"},
						"env":  {"dev", "staging"},
					},
					Port:     24,
					Protocol: "SSH",
				},
			},
			Policies: []AccessPolicy{
				expectedInfraAppPolicy,
			},
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplication, out, "create infrastructure_application structs not equal")
	}
}

func TestInfrastructureApplication_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+testInfrastructureApplicationId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "019205b5-97d7-7272-b00e-0ea05e61a124",
				"created_at": "2024-08-25T05:00:22Z",
				"updated_at": "2024-08-25T14:30:22Z",
				"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
				"name": "infrastructure_application_test_updated",
				"domain": "test.example.com/admin",
				"type": "infrastructure",
				"session_duration": "24h",
				"auto_redirect_to_identity": false,
				"enable_binding_cookie": false,
				"custom_deny_url": "https://www.example.com",
				"custom_non_identity_deny_url": "https://blocked.com",
				"custom_deny_message": "denied!",
				"http_only_cookie_attribute": true,
				"same_site_cookie_attribute": "strict",
				"logo_url": "https://www.example.com/example.png",
				"skip_interstitial": true,
				"app_launcher_visible": true,
				"service_auth_401_redirect": true,
				"path_cookie_attribute": true,
				"custom_pages": ["480f4f69-1a28-4fdd-9240-1ed29f0ac1dc"],
				"allow_authenticate_via_warp": true,
				"options_preflight_bypass": false,
				"target_criteria": [
					{
						"target_attributes": {
							"name": [ "vault 425", "vault 430" ],
							"env": [ "staging" ]
						},
						"port": 24,
						"protocol": "SSH"
					}
				], 
				"policies": [
					{
						"id": "699d98642c564d2e855e9661899b7252",
						"precedence": 1,
						"reusable": true,
						"decision": "allow",
						"created_at": "2020-08-25T05:00:22Z",
						"updated_at": "2020-08-25T05:00:22Z",
						"name": "my infra SSH policy",
						"include": [
							{
								"email": {
									"email": "devuser@example.com"
								}
							}
						],
						"connection_rules": {
							"ssh": {
								"usernames": ["devuser"]
							}
						}
					}
				]
			}
		}`)
	})

	_, err := client.UpdateAccessInfrastructureApplication(context.Background(), AccountIdentifier(""), UpdateInfrastructureApplicationParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.UpdateAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), UpdateInfrastructureApplicationParams{
		ID:                              "",
		InfrastructureApplicationParams: InfrastructureApplicationParams{},
	})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	out, err := client.UpdateAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), UpdateInfrastructureApplicationParams{
		ID: testInfrastructureApplicationId,
		InfrastructureApplicationParams: InfrastructureApplicationParams{
			Name: "infrastructure_application_test_updated",
			Type: "infrastructure",
			TargetContexts: []InfraTargetContext{
				{
					TargetAttributes: map[string][]string{
						"name": {"vault 420", "vault 430"},
						"env":  {"staging"},
					},
					Port:     24,
					Protocol: "SSH",
				},
			},
			Policies: []AccessPolicy{
				expectedInfraAppPolicy,
			},
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplicationModified, out, "update infrastructure_application structs not equal")
	}
}

func TestInfrastructureApplication_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+testInfrastructureApplicationId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "019205b5-97d7-7272-b00e-0ea05e61a124",
				"created_at": "2024-08-25T05:00:22Z",
				"updated_at": "2024-08-25T05:00:22Z",
				"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
				"name": "infrastructure_application_test",
				"domain": "test.example.com/admin",
				"type": "infrastructure",
				"session_duration": "24h",
				"auto_redirect_to_identity": false,
				"enable_binding_cookie": false,
				"custom_deny_url": "https://www.example.com",
				"custom_non_identity_deny_url": "https://blocked.com",
				"custom_deny_message": "denied!",
				"http_only_cookie_attribute": true,
				"same_site_cookie_attribute": "strict",
				"logo_url": "https://www.example.com/example.png",
				"skip_interstitial": true,
				"app_launcher_visible": true,
				"service_auth_401_redirect": true,
				"path_cookie_attribute": true,
				"custom_pages": ["480f4f69-1a28-4fdd-9240-1ed29f0ac1dc"],
				"allow_authenticate_via_warp": true,
				"options_preflight_bypass": false,
				"target_criteria": [
					{
						"target_attributes": {
							"name": [ "vault 420", "vault 421" ],
							"env": [ "dev", "staging" ]
						},
						"port": 24,
						"protocol": "SSH"
					}
				], 
				"policies": [
					{
						"id": "699d98642c564d2e855e9661899b7252",
						"precedence": 1,
						"reusable": true,
						"decision": "allow",
						"created_at": "2020-08-25T05:00:22Z",
						"updated_at": "2020-08-25T05:00:22Z",
						"name": "my infra SSH policy",
						"include": [
							{
								"email": {
									"email": "devuser@example.com"
								}
							}
						],
						"connection_rules": {
							"ssh": {
								"usernames": ["devuser"]
							}
						}
					}
				]
			}
		}`)
	})

	_, err := client.GetAccessInfrastructureApplication(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.GetAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	out, err := client.GetAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), testInfrastructureApplicationId)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplication, out, "get infrastructure_target not equal to expected")
	}
}

func TestInfrastructureApplication_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+testInfrastructureApplicationId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, ``)
	})

	// Make sure missing account ID is thrown
	err := client.DeleteAccessInfrastructureApplication(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.DeleteAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	err = client.DeleteAccessInfrastructureApplication(context.Background(), AccountIdentifier(testAccountID), testInfrastructureApplicationId)
	assert.NoError(t, err)
}
