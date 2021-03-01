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
		assert.Equal(t, r.Method, http.MethodGet, "Expected method 'GET', got %s", r.Method)
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

	actual, err := client.AccessIdentityProviders(context.TODO(), accountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers", handler)

	actual, err = client.ZoneLevelAccessIdentityProviders(context.TODO(), zoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccessIdentityProviderDetails(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet, "Expected method 'GET', got %s", r.Method)
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

	actual, err := client.AccessIdentityProviderDetails(context.TODO(), accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc841", handler)

	actual, err = client.ZoneLevelAccessIdentityProviderDetails(context.TODO(), zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc841")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessIdentityProvider(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost, "Expected method 'POST', got %s", r.Method)
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

	actual, err := client.CreateAccessIdentityProvider(context.TODO(), accountID, newIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers", handler)

	actual, err = client.CreateZoneLevelAccessIdentityProvider(context.TODO(), zoneID, newIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
func TestUpdateAccessIdentityProvider(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPut, "Expected method 'PUT', got %s", r.Method)
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

	actual, err := client.UpdateAccessIdentityProvider(context.TODO(), accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.UpdateZoneLevelAccessIdentityProvider(context.TODO(), zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", updatedIdentityProvider)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteAccessIdentityProvider(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodDelete, "Expected method 'DELETE', got %s", r.Method)
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

	actual, err := client.DeleteAccessIdentityProvider(context.TODO(), "01a7362d577a6c3019a474fd6f485823", "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/identity_providers/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.DeleteZoneLevelAccessIdentityProvider(context.TODO(), zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
