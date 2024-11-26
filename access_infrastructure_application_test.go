package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testInfrastructureApplicationId = "f52d1f8f-f194-4ac1-845f-e47203bf4dbb"

var (
	accessPolicyCreatedOn, _              = time.Parse(time.RFC3339, "2024-09-24T17:58:44Z")
	accessPolicyUpdatedOn, _              = time.Parse(time.RFC3339, "2024-09-24T17:58:44Z")
	infrastructureApplicationCreatedOn, _ = time.Parse(time.RFC3339, "2024-09-14T05:13:43Z")
	infrastructureApplicationUpdatedOn, _ = time.Parse(time.RFC3339, "2024-09-24T17:58:44Z")
	expectedInfraAppPolicy                = AccessPolicy{
		ID:         "8dd6af47-a80b-48cb-89c1-f39299ebf096",
		Precedence: 1,
		Reusable:   BoolPtr(false),
		Decision:   "allow",
		CreatedAt:  &accessPolicyCreatedOn,
		UpdatedAt:  &accessPolicyUpdatedOn,
		Name:       "my infra SSH policy",
		Include: []any{
			map[string]interface{}{"email": map[string]interface{}{"email": "devuser@cloudflare.com"}},
		},
		Exclude: []any{},
		Require: []any{},
		InfrastructureConnectionRules: &AccessInfrastructureConnectionRules{
			SSH: &AccessInfrastructureConnectionRulesSSH{
				Usernames:       []string{"devuser"},
				AllowEmailAlias: BoolPtr(true),
			},
		},
	}
	expectedInfrastructureApplication = AccessApplication{
		ID:                       testInfrastructureApplicationId,
		CreatedAt:                &infrastructureApplicationCreatedOn,
		UpdatedAt:                &infrastructureApplicationUpdatedOn,
		AUD:                      "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
		Name:                     "infrastructure_application_test",
		Domain:                   "",
		Type:                     "infrastructure",
		SessionDuration:          "",
		AllowedIdps:              nil,
		AutoRedirectToIdentity:   nil,
		EnableBindingCookie:      nil,
		AppLauncherVisible:       nil,
		ServiceAuth401Redirect:   nil,
		CustomDenyMessage:        "",
		CustomDenyURL:            "",
		SameSiteCookieAttribute:  "",
		HttpOnlyCookieAttribute:  nil,
		LogoURL:                  "",
		SkipInterstitial:         nil,
		PathCookieAttribute:      nil,
		CustomPages:              nil,
		Tags:                     nil,
		CustomNonIdentityDenyURL: "",
		AllowAuthenticateViaWarp: nil,
		OptionsPreflightBypass:   nil,
		TargetContexts: &[]AccessInfrastructureTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests"},
				},
				Port:     22,
				Protocol: "SSH",
			},
		},
		Policies: []AccessPolicy{
			expectedInfraAppPolicy,
		},
	}

	expectedInfrastructureApplicationUpdated = AccessApplication{
		ID:                       testInfrastructureApplicationId,
		CreatedAt:                &infrastructureApplicationCreatedOn,
		UpdatedAt:                &infrastructureApplicationUpdatedOn,
		AUD:                      "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
		Name:                     "infrastructure_application_test_updated",
		Domain:                   "",
		Type:                     "infrastructure",
		SessionDuration:          "",
		AllowedIdps:              nil,
		AutoRedirectToIdentity:   nil,
		EnableBindingCookie:      nil,
		AppLauncherVisible:       nil,
		ServiceAuth401Redirect:   nil,
		CustomDenyMessage:        "",
		CustomDenyURL:            "",
		SameSiteCookieAttribute:  "",
		HttpOnlyCookieAttribute:  nil,
		LogoURL:                  "",
		SkipInterstitial:         nil,
		PathCookieAttribute:      nil,
		CustomPages:              nil,
		Tags:                     nil,
		CustomNonIdentityDenyURL: "",
		AllowAuthenticateViaWarp: nil,
		OptionsPreflightBypass:   nil,
		TargetContexts: &[]AccessInfrastructureTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests", "cfgo-acc-tests-duplicate"},
				},
				Port:     22,
				Protocol: "SSH",
			},
		},
		Policies: []AccessPolicy{
			expectedInfraAppPolicy,
		},
	}
)

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
				"id": "f52d1f8f-f194-4ac1-845f-e47203bf4dbb",
				"type": "infrastructure",
				"name": "infrastructure_application_test",
				"aud": "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
				"created_at": "2024-09-14T05:13:43Z",
				"updated_at": "2024-09-24T17:58:44Z",
				"target_criteria": [
					{
						"target_attributes": {
							"hostname": [ "cfgo-acc-tests" ]
						},
						"port": 22,
						"protocol": "SSH"
					}
				],
				"policies": [
					{
						"created_at": "2024-09-24T17:58:44Z",
						"decision": "allow",
						"exclude": [],
						"id": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"include": [
							{
								"email": {
									"email": "devuser@cloudflare.com"
								}
							}
						],
						"name": "my infra SSH policy",
						"require": [],
						"connection_rules": {
							"ssh": {
								"usernames": [ "devuser" ],
								"allow_email_alias": true
							}
						},
						"uid": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"updated_at": "2024-09-24T17:58:44Z",
						"reusable": false,
						"precedence": 1
					}
				]
			}
		}`)
	})

	actual, err := client.CreateAccessApplication(context.Background(), AccountIdentifier(testAccountID), CreateAccessApplicationParams{
		Name: "infrastructure_application_test",
		Type: "infrastructure",
		TargetContexts: &[]AccessInfrastructureTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests"},
				},
				Port:     22,
				Protocol: "SSH",
			},
		},
		Policies: []string{
			"8dd6af47-a80b-48cb-89c1-f39299ebf096",
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplication, actual, "create infrastructure_application structs not equal")
	}
}

func TestInfrastructureApplication_Update(t *testing.T) {
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
				"id": "f52d1f8f-f194-4ac1-845f-e47203bf4dbb",
				"type": "infrastructure",
				"name": "infrastructure_application_test_updated",
				"aud": "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
				"created_at": "2024-09-14T05:13:43Z",
				"updated_at": "2024-09-24T17:58:44Z",
				"target_criteria": [
					{
						"target_attributes": {
							"hostname": [ "cfgo-acc-tests", "cfgo-acc-tests-duplicate" ]
						},
						"port": 22,
						"protocol": "SSH"
					}
				],
				"policies": [
					{
						"created_at": "2024-09-24T17:58:44Z",
						"decision": "allow",
						"exclude": [],
						"id": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"include": [
							{
								"email": {
									"email": "devuser@cloudflare.com"
								}
							}
						],
						"name": "my infra SSH policy",
						"require": [],
						"connection_rules": {
							"ssh": {
								"usernames": [ "devuser" ],
								"allow_email_alias": true
							}
						},
						"uid": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"updated_at": "2024-09-24T17:58:44Z",
						"reusable": false,
						"precedence": 1
					}
				]
			}
		}`)
	})

	actual, err := client.CreateAccessApplication(context.Background(), AccountIdentifier(testAccountID), CreateAccessApplicationParams{
		Name: "infrastructure_application_test_updated",
		Type: "infrastructure",
		TargetContexts: &[]AccessInfrastructureTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests", "cfgo-acc-tests-duplicate"},
				},
				Port:     22,
				Protocol: "SSH",
			},
		},
		Policies: []string{
			"8dd6af47-a80b-48cb-89c1-f39299ebf096",
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplicationUpdated, actual, "create infrastructure_application structs not equal")
	}
}

func TestInfrastructureApplication_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/f52d1f8f-f194-4ac1-845f-e47203bf4dbb", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "f52d1f8f-f194-4ac1-845f-e47203bf4dbb",
				"type": "infrastructure",
				"name": "infrastructure_application_test",
				"aud": "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
				"created_at": "2024-09-14T05:13:43Z",
				"updated_at": "2024-09-24T17:58:44Z",
				"target_criteria": [
					{
						"target_attributes": {
							"hostname": [ "cfgo-acc-tests" ]
						},
						"port": 22,
						"protocol": "SSH"
					}
				],
				"policies": [
					{
						"created_at": "2024-09-24T17:58:44Z",
						"decision": "allow",
						"exclude": [],
						"id": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"include": [
							{
								"email": {
									"email": "devuser@cloudflare.com"
								}
							}
						],
						"name": "my infra SSH policy",
						"require": [],
						"connection_rules": {
							"ssh": {
								"usernames": [ "devuser" ],
								"allow_email_alias": true
							}
						},
						"uid": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"updated_at": "2024-09-24T17:58:44Z",
						"reusable": false,
						"precedence": 1
					}
				]
			}
		}`)
	})

	actual, err := client.GetAccessApplication(context.Background(), AccountIdentifier(testAccountID), "f52d1f8f-f194-4ac1-845f-e47203bf4dbb")
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplication, actual, "get infrastructure_application structs not equal")
	}
}

func TestInfrastructureApplication_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/f52d1f8f-f194-4ac1-845f-e47203bf4dbb", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "success": true,
      "errors": [],
      "messages": [],
      "result": {
        "id": "f52d1f8f-f194-4ac1-845f-e47203bf4dbb"
      }
    }`)
	})

	err := client.DeleteAccessApplication(context.Background(), AccountIdentifier(testAccountID), "f52d1f8f-f194-4ac1-845f-e47203bf4dbb")
	assert.NoError(t, err)
}
