package cloudflare

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockPageShieldPolicies = []PageShieldPolicy{
	{
		Action:      "allow",
		Description: "Checkout page CSP policy",
		Enabled:     true,
		Expression:  "ends_with(http.request.uri.path, \"/checkout\")",
		ID:          "c9ef84a6bf5e47138c75d95e2f933e8f",
		Value:       "script-src 'none';",
	},
	{
		Action:      "allow",
		Description: "Checkout page CSP policy",
		Enabled:     true,
		Expression:  "ends_with(http.request.uri.path, \"/login\")",
		ID:          "c9ef84a6bf5e47138c75d95e2f933e82",
		Value:       "script-src 'none';",
	},
	{
		Action:      "allow",
		Description: "Checkout page CSP policy",
		Enabled:     true,
		Expression:  "ends_with(http.request.uri.path, \"/logout\")",
		ID:          "c9ef84a6bf5e47138c75d95e2f933e83",
		Value:       "script-src 'none';",
	},
}

func TestListPageShieldPolicies(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/testzone/page_shield/policies", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		response := ListPageShieldPoliciesResponse{
			Result: mockPageShieldPolicies,
		}
		json.NewEncoder(w).Encode(response)
	})

	result, _, err := client.ListPageShieldPolicies(context.Background(), &ResourceContainer{Identifier: "testzone"})
	assert.NoError(t, err)
	assert.Equal(t, mockPageShieldPolicies, result)
}

func TestCreatePageShieldPolicy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/testzone/page_shield/policies", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		var params PageShieldPolicy
		err := json.NewDecoder(r.Body).Decode(&params)
		assert.NoError(t, err)
		params.ID = "newPolicyID"
		json.NewEncoder(w).Encode(params)
	})

	newPolicy := PageShieldPolicy{
		Action:      "block",
		Description: "New policy",
		Enabled:     true,
		Expression:  "ends_with(http.request.uri.path, \"/new\")",
		Value:       "script-src 'self';",
	}

	result, err := client.CreatePageShieldPolicy(context.Background(), &ResourceContainer{Identifier: "testzone"}, newPolicy)
	assert.NoError(t, err)
	assert.Equal(t, "newPolicyID", result.ID)
}

func TestDeletePageShieldPolicy(t *testing.T) {
	setup()
	defer teardown()

	policyID := "c9ef84a6bf5e47138c75d95e2f933e8f"
	mux.HandleFunc(fmt.Sprintf("/zones/testzone/page_shield/policies/%s", policyID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.WriteHeader(http.StatusOK) // Assuming successful deletion returns 200 OK
	})

	err := client.DeletePageShieldPolicy(context.Background(), &ResourceContainer{Identifier: "testzone"}, policyID)
	assert.NoError(t, err)
}

func TestGetPageShieldPolicy(t *testing.T) {
	setup()
	defer teardown()

	policyID := "c9ef84a6bf5e47138c75d95e2f933e8f"
	mux.HandleFunc(fmt.Sprintf("/zones/testzone/page_shield/policies/%s", policyID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(mockPageShieldPolicies[0]) // Assuming the first mock policy
	})

	result, err := client.GetPageShieldPolicy(context.Background(), &ResourceContainer{Identifier: "testzone"}, policyID)
	assert.NoError(t, err)
	assert.Equal(t, &mockPageShieldPolicies[0], result)
}

func TestUpdatePageShieldPolicy(t *testing.T) {
	setup()
	defer teardown()

	policyID := "c9ef84a6bf5e47138c75d95e2f933e8f"
	mux.HandleFunc(fmt.Sprintf("/zones/testzone/page_shield/policies/%s", policyID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		var params PageShieldPolicy
		err := json.NewDecoder(r.Body).Decode(&params)
		assert.NoError(t, err)
		params.ID = policyID
		json.NewEncoder(w).Encode(params)
	})

	updatedPolicy := PageShieldPolicy{
		Action:      "block",
		Description: "Updated policy",
		Enabled:     false,
		Expression:  "ends_with(http.request.uri.path, \"/updated\")",
		Value:       "script-src 'self';",
	}

	result, err := client.UpdatePageShieldPolicy(context.Background(), &ResourceContainer{Identifier: "testzone"}, policyID, updatedPolicy)
	assert.NoError(t, err)
	assert.Equal(t, policyID, result.ID)
	assert.Equal(t, "Updated policy", result.Description)
}
