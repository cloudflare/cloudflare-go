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

	want := []AccessIdentityProvider{
		{
			ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			Name: "Widget Corps OTP",
			Type: "github",
			Config: AccessIdentityProviderConfiguration{
				ClientID:     "example_id",
				ClientSecret: "a-secret-key",
			},
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/identity_providers", handler)

	actual, err := client.AccessIdentityProviders(accountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers", handler)

	actual, err = client.ZoneLevelAccessIdentityProviders(zoneID)

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

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessIdentityProviderConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc841", handler)

	actual, err := client.AccessIdentityProviderDetails(accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc841", handler)

	actual, err = client.ZoneLevelAccessIdentityProviderDetails(zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

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

	newIdentityProvider := AccessIdentityProvider{
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessIdentityProviderConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessIdentityProviderConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/identity_providers", handler)

	actual, err := client.CreateAccessIdentityProvider(accountID, newIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers", handler)

	actual, err = client.CreateZoneLevelAccessIdentityProvider(zoneID, newIdentityProvider)

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

	updatedIdentityProvider := AccessIdentityProvider{
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessIdentityProviderConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessIdentityProviderConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err := client.UpdateAccessIdentityProvider(accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.UpdateZoneLevelAccessIdentityProvider(zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

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

	want := AccessIdentityProvider{
		ID:   "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name: "Widget Corps OTP",
		Type: "github",
		Config: AccessIdentityProviderConfiguration{
			ClientID:     "example_id",
			ClientSecret: "a-secret-key",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err := client.DeleteAccessIdentityProvider("01a7362d577a6c3019a474fd6f485823", "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.DeleteZoneLevelAccessIdentityProvider(zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
