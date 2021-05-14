package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessMutualTLSCertificates(t *testing.T) {
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
					"created_at": "2014-01-01T05:20:00.12345Z",
					"updated_at": "2014-01-01T05:20:00.12345Z",
					"expires_on": "2014-01-01T05:20:00.12345Z",
					"name": "Allow devs",
					"fingerprint": "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
					"associated_hostnames": [
						"admin.example.com"
					]
				}
			]
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	expiresOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := []AccessMutualTLSCertificate{{
		ID:                  "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		ExpiresOn:           expiresOn,
		Name:                "Allow devs",
		Fingerprint:         "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
		AssociatedHostnames: []string{"admin.example.com"},
	}}

	mux.HandleFunc("/accounts/"+accountID+"/access/certificates", handler)

	actual, err := client.AccessMutualTLSCertificates(context.Background(), accountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/certificates", handler)

	actual, err = client.ZoneAccessMutualTLSCertificates(context.Background(), zoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccessMutualTLSCertificate(t *testing.T) {
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
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"expires_on": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"fingerprint": "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
				"associated_hostnames": [
					"admin.example.com"
				]
			}
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	expiresOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := AccessMutualTLSCertificate{
		ID:                  "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		ExpiresOn:           expiresOn,
		Name:                "Allow devs",
		Fingerprint:         "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
		AssociatedHostnames: []string{"admin.example.com"},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/certificates/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err := client.AccessMutualTLSCertificate(context.Background(), accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/certificates/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.ZoneAccessMutualTLSCertificate(context.Background(), zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessMutualTLSCertificate(t *testing.T) {
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
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"expires_on": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"fingerprint": "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
				"associated_hostnames": [
					"admin.example.com"
				]
			}
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	expiresOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	certificate := AccessMutualTLSCertificate{
		Name:        "Allow devs",
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10\nDrUtmu/B\n-----END CERTIFICATE-----",
	}

	want := AccessMutualTLSCertificate{
		ID:                  "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		ExpiresOn:           expiresOn,
		Name:                "Allow devs",
		Fingerprint:         "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
		AssociatedHostnames: []string{"admin.example.com"},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/certificates", handler)

	actual, err := client.CreateAccessMutualTLSCertificate(context.Background(), accountID, certificate)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/certificates", handler)

	actual, err = client.CreateZoneAccessMutualTLSCertificate(context.Background(), zoneID, certificate)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateAccessMutualTLSCertificate(t *testing.T) {
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
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"expires_on": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"fingerprint": "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
				"associated_hostnames": [
					"admin.example.com"
				]
			}
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	expiresOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	certificate := AccessMutualTLSCertificate{
		Name:        "Allow devs",
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10\nDrUtmu/B\n-----END CERTIFICATE-----",
	}

	want := AccessMutualTLSCertificate{
		ID:                  "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		ExpiresOn:           expiresOn,
		Name:                "Allow devs",
		Fingerprint:         "MD5 Fingerprint=1E:80:0F:7A:FD:31:55:96:DE:D5:CB:E2:F0:91:F6:91",
		AssociatedHostnames: []string{"admin.example.com"},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/certificates/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err := client.UpdateAccessMutualTLSCertificate(context.Background(), accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", certificate)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/certificates/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	actual, err = client.UpdateZoneAccessMutualTLSCertificate(context.Background(), zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", certificate)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteAccessMutualTLSCertificate(t *testing.T) {
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
				"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
			}
		}`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/certificates/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	err := client.DeleteAccessMutualTLSCertificate(context.Background(), accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+zoneID+"/access/certificates/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	err = client.DeleteZoneAccessMutualTLSCertificate(context.Background(), zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	assert.NoError(t, err)
}
