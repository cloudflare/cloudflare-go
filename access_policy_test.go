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
					]
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

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, _, err := client.AccessPolicies(context.Background(), accountID, accessApplicationID, pageOptions)

	if assert.NoError(t, err) {
		assert.Equal(t, []AccessPolicy{expectedAccessPolicy}, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, _, err = client.ZoneLevelAccessPolicies(context.Background(), zoneID, accessApplicationID, pageOptions)

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
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)

	actual, err := client.AccessPolicy(context.Background(), accountID, accessApplicationID, accessPolicyID)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)

	actual, err = client.ZoneLevelAccessPolicy(context.Background(), zoneID, accessApplicationID, accessPolicyID)

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
				]
			}
		}
		`)
	}

	accessPolicy := AccessPolicy{
		Name: "Allow devs",
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
		Decision: "allow",
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, err := client.CreateAccessPolicy(context.Background(), accountID, accessApplicationID, accessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/"+accessApplicationID+"/policies", handler)

	actual, err = client.CreateZoneLevelAccessPolicy(context.Background(), zoneID, accessApplicationID, accessPolicy)

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
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	actual, err := client.UpdateAccessPolicy(context.Background(), accountID, accessApplicationID, expectedAccessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	actual, err = client.UpdateZoneLevelAccessPolicy(context.Background(), zoneID, accessApplicationID, expectedAccessPolicy)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessPolicy, actual)
	}
}

func TestUpdateAccessPolicyWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessPolicy(context.Background(), accountID, accessApplicationID, AccessPolicy{})
	assert.EqualError(t, err, "access policy ID cannot be empty")

	_, err = client.UpdateZoneLevelAccessPolicy(context.Background(), zoneID, accessApplicationID, AccessPolicy{})
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

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	err := client.DeleteAccessPolicy(context.Background(), accountID, accessApplicationID, accessPolicyID)

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/"+accessApplicationID+"/policies/"+accessPolicyID, handler)
	err = client.DeleteZoneLevelAccessPolicy(context.Background(), zoneID, accessApplicationID, accessPolicyID)

	assert.NoError(t, err)
}
