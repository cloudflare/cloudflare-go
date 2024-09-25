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
		Aud:       "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
		CreatedAt: infrastructureAccessTargetCreatedOn.String(),
		UpdatedAt: infrastructureApplicationCreatedOn.String(),
		TargetContexts: []InfraTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests"},
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
		Aud:       "6570722c13470d8c3e7aed619fc66cdf948002efbb06601166b8f2dad85fae34",
		CreatedAt: infrastructureAccessTargetCreatedOn.String(),
		UpdatedAt: infrastructureApplicationUpdatedOn.String(),
		TargetContexts: []InfraTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests"},
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
				"id": "0a20bd6b-26bb-4e9f-a835-f81f4e5f4427",
				"uid": "0a20bd6b-26bb-4e9f-a835-f81f4e5f4427",
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
								"usernames": [ "devuser" ]
							}
						},
						"uid": "8dd6af47-a80b-48cb-89c1-f39299ebf096",
						"updated_at": "2024-09-24T17:58:44Z",
						"reusable": false,
						"precedence": 1
					}
				]
			},
		}`)
	})

	// We expect this to parse through the Access Application response and only return
	// Infrastructure Application built from that response
	out, err := client.CreateAccessApplication(context.Background(), AccountIdentifier(testAccountID), CreateAccessApplicationParams{
		Name: "infrastructure_application_test",
		Type: "infrastructure",
		TargetContexts: []InfraTargetContext{
			{
				TargetAttributes: map[string][]string{
					"hostname": {"cfgo-acc-tests"},
				},
				Port:     24,
				Protocol: "SSH",
			},
		},
		Policies: []string{
			"8dd6af47-a80b-48cb-89c1-f39299ebf096",
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureApplication, out, "create infrastructure_application structs not equal")
	}
}
