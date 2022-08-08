package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWorkersAccountSettings_CreateAccountSettings(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/account-settings", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "default_usage_model": "unbound",
	"green_compute": false
  }
}`) //nolint
	})

	// Make sure missing account ID is thrown
	_, err := client.CreateWorkersAccountSettings(context.Background(), CreateWorkersAccountSettingsParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	res, err := client.CreateWorkersAccountSettings(context.Background(), CreateWorkersAccountSettingsParameters{AccountID: testAccountID, DefaultUsageModel: "unbound", GreenCompute: false})
	want := WorkersAccountSettings{
		DefaultUsageModel: "unbound",
		GreenCompute:      false,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersAccountSettings_GetAccountSettings(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/account-settings", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "default_usage_model": "unbound",
	"green_compute": true
  }
}`) //nolint
	})

	// Make sure missing account ID is thrown
	_, err := client.GetWorkersAccountSettings(context.Background(), GetWorkersAccountSettingsParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	res, err := client.GetWorkersAccountSettings(context.Background(), GetWorkersAccountSettingsParameters{AccountID: testAccountID})
	want := WorkersAccountSettings{
		DefaultUsageModel: "unbound",
		GreenCompute:      true,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}
