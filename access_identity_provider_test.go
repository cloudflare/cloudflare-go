package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessIdentityProviders(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
					"name": "Widget Corps OTP",
					"type": "github",
					"config": {
						"client_id": "example_id",
						"client_secret": "a-secret-key"
					}
				}
			]
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/identity_providers", handler)

	want := []AccessIdentityProvider{AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: map[string]interface{}{
			"client_id":     "example_id",
			"client_secret": "a-secret-key",
		},
	}}

	actual, err := client.AccessIdentityProviders("01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccessIdentityProviderDetails(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name": "Widget Corps OTP",
				"type": "github",
				"config": {
					"client_id": "example_id",
					"client_secret": "a-secret-key"
				}
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc841", handler)

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: map[string]interface{}{
			"client_id":     "example_id",
			"client_secret": "a-secret-key",
		},
	}

	actual, err := client.AccessIdentityProviderDetails("01a7362d577a6c3019a474fd6f485823", "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessIdentityProvider(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name": "Widget Corps OTP",
				"type": "github",
				"config": {
					"client_id": "example_id",
					"client_secret": "a-secret-key"
				}
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/identity_providers", handler)

	newIdentityProvider := AccessIdentityProvider{
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessGitHubConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: map[string]interface{}{
			"client_id":     "example_id",
			"client_secret": "a-secret-key",
		},
	}

	actual, err := client.CreateAccessIdentityProvider("01a7362d577a6c3019a474fd6f485823", newIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
func TestUpdateAccessIdentityProvider(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "PUT", "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name": "Widget Corps OTP",
				"type": "github",
				"config": {
					"client_id": "example_id",
					"client_secret": "a-secret-key"
				}
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	updatedIdentityProvider := AccessIdentityProvider{
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessGitHubConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: map[string]interface{}{
			"client_id":     "example_id",
			"client_secret": "a-secret-key",
		},
	}

	actual, err := client.UpdateAccessIdentityProvider("01a7362d577a6c3019a474fd6f485823", "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteAccessIdentityProvider(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name": "Widget Corps OTP",
				"type": "github",
				"config": {
					"client_id": "example_id",
					"client_secret": "a-secret-key"
				}
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: map[string]interface{}{
			"client_id":     "example_id",
			"client_secret": "a-secret-key",
		},
	}

	actual, err := client.DeleteAccessIdentityProvider("01a7362d577a6c3019a474fd6f485823", "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
