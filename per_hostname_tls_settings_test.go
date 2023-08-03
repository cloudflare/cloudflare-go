package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListHostnameTLSSettingsMinTLSVersion(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"hostname": "app.example.com",
					"value": "1.2",
					"status": "active",
					"created_at": "2023-07-26T21:12:55.56942Z",
					"updated_at": "2023-07-31T22:06:44.739794Z"
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 50,
				"count": 1,
				"total_count": 1,
				"total_pages": 1
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/hostnames/settings/min_tls_version", handler)
	createdAt, _ := time.Parse(time.RFC3339, "2023-07-26T21:12:55.56942Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-07-31T22:06:44.739794Z")
	want := []HostnameTLSSetting{
		{
			Hostname:  "app.example.com",
			Value:     "1.2",
			Status:    "active",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	actual, _, err := client.ListHostnameTLSSettings(context.Background(), ZoneIdentifier(testZoneID), "min_tls_version", ListHostnameTLSSettingsParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListHostnameTLSSettingsCiphers(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"hostname": "app.example.com",
					"value": [
						"AES128-GCM-SHA256",
						"ECDHE-RSA-AES128-GCM-SHA256"
					],
					"status": "active",
					"created_at": "2023-07-26T21:12:55.56942Z",
					"updated_at": "2023-07-31T22:06:44.739794Z"
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 50,
				"count": 1,
				"total_count": 1,
				"total_pages": 1
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/hostnames/settings/ciphers", handler)
	createdAt, _ := time.Parse(time.RFC3339, "2023-07-26T21:12:55.56942Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-07-31T22:06:44.739794Z")

	strs := []string{"AES128-GCM-SHA256", "ECDHE-RSA-AES128-GCM-SHA256"}
	cipherssuites := make([]interface{}, len(strs))
	for i, s := range strs {
		cipherssuites[i] = s
	}
	want := []HostnameTLSSetting{
		{
			Hostname:  "app.example.com",
			Value:     cipherssuites,
			Status:    "active",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	actual, _, err := client.ListHostnameTLSSettings(context.Background(), ZoneIdentifier(testZoneID), "ciphers", ListHostnameTLSSettingsParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestEditHostnameTLSSetting(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"hostname": "app.example.com",
				"value": [
					"AES128-GCM-SHA256",
					"ECDHE-RSA-AES128-GCM-SHA256"
				],
				"status": "active",
				"created_at": "2023-07-26T21:12:55.56942Z",
				"updated_at": "2023-07-31T22:06:44.739794Z"
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/hostnames/settings/ciphers/app.example.com", handler)
	createdAt, _ := time.Parse(time.RFC3339, "2023-07-26T21:12:55.56942Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-07-31T22:06:44.739794Z")

	strs := []string{"AES128-GCM-SHA256", "ECDHE-RSA-AES128-GCM-SHA256"}
	cipherssuites := make([]interface{}, len(strs))
	for i, s := range strs {
		cipherssuites[i] = s
	}
	want := HostnameTLSSetting{
		Hostname:  "app.example.com",
		Value:     cipherssuites,
		Status:    "active",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	params := EditHostnameTLSSettingParams{
		Value: []string{"AES128-GCM-SHA256", "ECDHE-RSA-AES128-GCM-SHA256"},
	}
	actual, err := client.EditHostnameTLSSetting(context.Background(), ZoneIdentifier(testZoneID), "ciphers", "app.example.com", params)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteHostnameTLSSetting(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"hostname": "app.example.com",
				"value": "",
				"status": "active",
				"created_at": "2023-07-26T21:12:55.56942Z",
				"updated_at": "2023-07-31T22:06:44.739794Z"
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/hostnames/settings/ciphers/app.example.com", handler)
	createdAt, _ := time.Parse(time.RFC3339, "2023-07-26T21:12:55.56942Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-07-31T22:06:44.739794Z")
	want := HostnameTLSSetting{
		Hostname:  "app.example.com",
		Value:     "",
		Status:    "active",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	actual, err := client.DeleteHostnameTLSSetting(context.Background(), ZoneIdentifier(testZoneID), "ciphers", "app.example.com")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
