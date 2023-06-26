package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	pageOptions         = PaginationOptions{}
	accessApplicationID = "6e1c88f1-6b06-4b8a-a9e9-3ec7da2ee0c1"
	accessPolicyID      = "699d98642c564d2e855e9661899b7252"

	createdAt, _ = time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ = time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	expiresAt, _ = time.Parse(time.RFC3339, "2015-01-01T05:20:00.12345Z")

	isolationRequired            = true
	purposeJustificationRequired = true
	purposeJustificationPrompt   = "Please provide a business reason for your need to access before continuing."
	approvalRequired             = true

	expectedAccessPolicy = AccessPolicy{
		ID:         "699d98642c564d2e855e9661899b7252",
		Precedence: 1,
		Decision:   "allow",
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
		Name:       "Allow devs",
		Include: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
		},
		Exclude: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
		},
		Require: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
		},
		IsolationRequired:            &isolationRequired,
		PurposeJustificationRequired: &purposeJustificationRequired,
		ApprovalRequired:             &approvalRequired,
		PurposeJustificationPrompt:   &purposeJustificationPrompt,
		ApprovalGroups: []AccessApprovalGroup{
			{
				EmailListUuid:   "2413b6d7-bbe5-48bd-8fbb-e52069c85561",
				ApprovalsNeeded: 3,
			},
			{
				EmailAddresses:  []string{"email1@example.com", "email2@example.com"},
				ApprovalsNeeded: 1,
			},
		},
	}
)

func TestAccessPolicies(t *testing.T) {
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
				"id": "699d98642c564d2e855e9661899b7252",
				"precedence": 1,
				"decision": "allow",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
				  {
					"email": {
					  "email": "test@example.com"
					}
				  }
				],
				"exclude": [
				  {
					"email": {
					  "email": "test@example.com"
					}
				  }
				],
				"require": [
				  {
					"email": {
					  "email": "test@example.com"
					}
				  }
				],
				"isolation_required": true,
				"purpose_justification_required": true,
				"purpose_justification_prompt": "Please provide a business reason for your need to access before continuing.",
				"approval_required": true,
				"approval_groups": [
				  {
					"email_list_uuid": "2413b6d7-bbe5-48bd-8fbb-e52069c85561",
					"approvals_needed": 3
				  },
				  {
					"email_addresses": [
					  "email1@example.com",
					  "email2@example.com"
					],
					"approvals_needed": 1
				  }
				]
			  }
			],
			"result_info": {
			  "page": 1,
			  "per_page": 20,
			  "count": 1,
			  "total_count": 20
			}
		  }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, _, err := client.ListAccessPolicies(context.Background(), testAccountRC, ListAccessPoliciesParams{ApplicationID: accessApplicationID})

	if assert.NoError(t, err) {
		assert.Equal(t, []AccessPolicy{expectedAccessPolicy}, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, _, err = client.ListAccessPolicies(context.Background(), testZoneRC, ListAccessPoliciesParams{ApplicationID: accessApplicationID})

	if assert.NoError(t, err) {
		assert.Equal(t, []AccessPolicy{expectedAccessPolicy}, actual)
	}
}

func TestAccessPolicy(t *testing.T) {
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
				"id": "699d98642c564d2e855e9661899b7252",
				"precedence": 1,
				"decision": "allow",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"exclude": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"require": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"isolation_required": true,
				"purpose_justification_required": true,
				"purpose_justification_prompt": "Please provide a business reason for your need to access before continuing.",
				"approval_required": true,
				"approval_groups": [
					{
						"email_list_uuid": "2413b6d7-bbe5-48bd-8fbb-e52069c85561",
						"approvals_needed": 3
					},
					{
						"email_addresses": ["email1@example.com", "email2@example.com"],
						"approvals_needed": 1
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)

	actual, err := client.GetAccessPolicy(context.Background(), testAccountRC, GetAccessPolicyParams{ApplicationID: accessApplicationID, PolicyID: accessPolicyID})

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)

	actual, err = client.GetAccessPolicy(context.Background(), testZoneRC, GetAccessPolicyParams{ApplicationID: accessApplicationID, PolicyID: accessPolicyID})

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}
}

func TestCreateAccessPolicy(t *testing.T) {
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
				"id": "699d98642c564d2e855e9661899b7252",
				"precedence": 1,
				"decision": "allow",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"exclude": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"require": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"isolation_required": true,
				"purpose_justification_required": true,
				"purpose_justification_prompt": "Please provide a business reason for your need to access before continuing.",
				"approval_required": true,
				"approval_groups": [
					{
						"email_list_uuid": "2413b6d7-bbe5-48bd-8fbb-e52069c85561",
						"approvals_needed": 3
					},
					{
						"email_addresses": ["email1@example.com", "email2@example.com"],
						"approvals_needed": 1
					}
				]
			}
		}
		`)
	}

	accessPolicy := CreateAccessPolicyParams{
		ApplicationID: accessApplicationID,
		Name:          "Allow devs",
		Include: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
		},
		Exclude: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
		},
		Require: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
		},
		Decision:                     "allow",
		PurposeJustificationRequired: &purposeJustificationRequired,
		PurposeJustificationPrompt:   &purposeJustificationPrompt,
		ApprovalGroups: []AccessApprovalGroup{
			{
				EmailListUuid:   "2413b6d7-bbe5-48bd-8fbb-e52069c85561",
				ApprovalsNeeded: 3,
			},
			{
				EmailAddresses:  []string{"email1@example.com", "email2@example.com"},
				ApprovalsNeeded: 1,
			},
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, err := client.CreateAccessPolicy(context.Background(), testAccountRC, accessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, err = client.CreateAccessPolicy(context.Background(), testZoneRC, accessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}
}

func TestUpdateAccessPolicy(t *testing.T) {
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
				"id": "699d98642c564d2e855e9661899b7252",
				"precedence": 1,
				"decision": "allow",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"exclude": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"require": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"isolation_required": true,
				"purpose_justification_required": true,
				"purpose_justification_prompt": "Please provide a business reason for your need to access before continuing.",
				"approval_required": true,
				"approval_groups": [
					{
						"email_list_uuid": "2413b6d7-bbe5-48bd-8fbb-e52069c85561",
						"approvals_needed": 3
					},
					{
						"email_addresses": ["email1@example.com", "email2@example.com"],
						"approvals_needed": 1
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	actual, err := client.UpdateAccessPolicy(context.Background(), testAccountRC, accessApplicationID, expectedAccessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	actual, err = client.UpdateAccessPolicy(context.Background(), testZoneRC, accessApplicationID, expectedAccessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}
}

func TestUpdateAccessPolicyWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessPolicy(context.Background(), testAccountRC, accessApplicationID, AccessPolicy{})
	assert.EqualError(t, err, "access policy ID cannot be empty")

	_, err = client.UpdateAccessPolicy(context.Background(), testZoneRC, accessApplicationID, AccessPolicy{})
	assert.EqualError(t, err, "access policy ID cannot be empty")
}

func TestDeleteAccessPolicy(t *testing.T) {
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
				"id": "699d98642c564d2e855e9661899b7252"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	err := client.DeleteAccessPolicy(context.Background(), testAccountRC, DeleteAccessPolicyParams{ApplicationID: accessApplicationID, PolicyID: accessPolicyID})

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+testZoneID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	err = client.DeleteAccessPolicy(context.Background(), testZoneRC, DeleteAccessPolicyParams{ApplicationID: accessApplicationID, PolicyID: accessPolicyID})

	assert.NoError(t, err)
}
