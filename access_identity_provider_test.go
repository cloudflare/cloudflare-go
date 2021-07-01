package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessIdentityProviders(t *testing.T) {
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

	mux.HandleFunc("/accounts/"+testAccountID+"/access/identity_providers", handler)

	actual, err := client.AccessIdentityProviders(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/identity_providers", handler)

	actual, err = client.ZoneLevelAccessIdentityProviders(context.Background(), testZoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccessIdentityProviderDetails(t *testing.T) {
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

	mux.HandleFunc("/accounts/"+testAccountID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc841", handler)

	actual, err := client.AccessIdentityProviderDetails(context.Background(), testAccountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc841", handler)

	actual, err = client.ZoneLevelAccessIdentityProviderDetails(context.Background(), testZoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessIdentityProvider(t *testing.T) {
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

	mux.HandleFunc("/accounts/"+testAccountID+"/access/identity_providers", handler)

	actual, err := client.CreateAccessIdentityProvider(context.Background(), testAccountID, newIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/identity_providers", handler)

	actual, err = client.CreateZoneLevelAccessIdentityProvider(context.Background(), testZoneID, newIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
func TestUpdateAccessIdentityProvider(t *testing.T) {
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

	mux.HandleFunc("/accounts/"+testAccountID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err := client.UpdateAccessIdentityProvider(context.Background(), testAccountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.UpdateZoneLevelAccessIdentityProvider(context.Background(), testZoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteAccessIdentityProvider(t *testing.T) {
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

	mux.HandleFunc("/accounts/"+testAccountID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err := client.DeleteAccessIdentityProvider(context.Background(), "01a7362d577a6c3019a474fd6f485823", "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.DeleteZoneLevelAccessIdentityProvider(context.Background(), testZoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
