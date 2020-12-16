package cloudflare

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessApplications(t *testing.T) {
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
					"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
					"created_at": "2014-01-01T05:20:00.12345Z",
					"updated_at": "2014-01-01T05:20:00.12345Z",
					"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
					"name": "Admin Site",
					"domain": "test.example.com/admin",
					"session_duration": "24h",
					"allowed_idps": ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"],
					"auto_redirect_to_identity": false,
					"enable_binding_cookie": false,
					"custom_deny_url": "https://www.cloudflare.com",
					"custom_deny_message": "denied!"
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

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := []AccessApplication{{
		ID:                     "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		CreatedAt:              &createdAt,
		UpdatedAt:              &updatedAt,
		AUD:                    "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		Name:                   "Admin Site",
		Domain:                 "test.example.com/admin",
		SessionDuration:        "24h",
		AllowedIdps:            []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"},
		AutoRedirectToIdentity: false,
		EnableBindingCookie:    false,
		CustomDenyMessage:      "denied!",
		CustomDenyURL:          "https://www.cloudflare.com",
	}}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps", handler)

	actual, _, err := client.AccessApplications(accountID, PaginationOptions{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps", handler)

	actual, _, err = client.ZoneLevelAccessApplications(zoneID, PaginationOptions{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccessApplication(t *testing.T) {
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
				"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
				"name": "Admin Site",
				"domain": "test.example.com/admin",
				"session_duration": "24h",
				"allowed_idps": ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"],
				"auto_redirect_to_identity": false,
				"enable_binding_cookie": false,
				"custom_deny_url": "https://www.cloudflare.com",
				"custom_deny_message": "denied!"
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := AccessApplication{
		ID:                     "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		CreatedAt:              &createdAt,
		UpdatedAt:              &updatedAt,
		AUD:                    "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		Name:                   "Admin Site",
		Domain:                 "test.example.com/admin",
		SessionDuration:        "24h",
		AllowedIdps:            []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"},
		AutoRedirectToIdentity: false,
		EnableBindingCookie:    false,
		CustomDenyMessage:      "denied!",
		CustomDenyURL:          "https://www.cloudflare.com",
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.AccessApplication(accountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err = client.ZoneLevelAccessApplication(zoneID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessApplications(t *testing.T) {
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
				"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
				"name": "Admin Site",
				"domain": "test.example.com/admin",
				"session_duration": "24h",
				"allowed_idps": ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"],
				"auto_redirect_to_identity": false,
				"enable_binding_cookie": false,
				"custom_deny_url": "https://www.cloudflare.com",
				"custom_deny_message": "denied!"
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	fullAccessApplication := AccessApplication{
		ID:                     "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:                   "Admin Site",
		Domain:                 "test.example.com/admin",
		SessionDuration:        "24h",
		AUD:                    "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		AllowedIdps:            []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"},
		AutoRedirectToIdentity: false,
		EnableBindingCookie:    false,
		CustomDenyMessage:      "denied!",
		CustomDenyURL:          "https://www.cloudflare.com",
		CreatedAt:              &createdAt,
		UpdatedAt:              &updatedAt,
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps", handler)

	actual, err := client.CreateAccessApplication(
		accountID,
		AccessApplication{
			Name:            "Admin Site",
			Domain:          "test.example.com/admin",
			SessionDuration: "24h",
		},
	)

	if assert.NoError(t, err) {
		assert.Equal(t, fullAccessApplication, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps", handler)

	actual, err = client.CreateZoneLevelAccessApplication(
		zoneID,
		AccessApplication{
			Name:            "Admin Site",
			Domain:          "test.example.com/admin",
			SessionDuration: "24h",
		},
	)

	if assert.NoError(t, err) {
		assert.Equal(t, fullAccessApplication, actual)
	}
}

func TestUpdateAccessApplication(t *testing.T) {
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
				"id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
				"name": "Admin Site",
				"domain": "test.example.com/admin",
				"session_duration": "24h",
				"allowed_idps": ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"],
				"auto_redirect_to_identity": false,
				"enable_binding_cookie": false,
				"custom_deny_url": "https://www.cloudflare.com",
				"custom_deny_message": "denied!"
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	fullAccessApplication := AccessApplication{
		ID:                     "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		Name:                   "Admin Site",
		Domain:                 "test.example.com/admin",
		SessionDuration:        "24h",
		AUD:                    "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		AllowedIdps:            []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"},
		AutoRedirectToIdentity: false,
		EnableBindingCookie:    false,
		CustomDenyMessage:      "denied!",
		CustomDenyURL:          "https://www.cloudflare.com",
		CreatedAt:              &createdAt,
		UpdatedAt:              &updatedAt,
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.UpdateAccessApplication(
		accountID,
		fullAccessApplication,
	)

	if assert.NoError(t, err) {
		assert.Equal(t, fullAccessApplication, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err = client.UpdateZoneLevelAccessApplication(
		zoneID,
		fullAccessApplication,
	)

	if assert.NoError(t, err) {
		assert.Equal(t, fullAccessApplication, actual)
	}
}

func TestUpdateAccessApplicationWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessApplication(zoneID, AccessApplication{})
	assert.EqualError(t, err, "access application ID cannot be empty")

	_, err = client.UpdateZoneLevelAccessApplication(zoneID, AccessApplication{})
	assert.EqualError(t, err, "access application ID cannot be empty")
}

func TestDeleteAccessApplication(t *testing.T) {
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
        "id": "699d98642c564d2e855e9661899b7252"
      }
    }
    `)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)
	err := client.DeleteAccessApplication(accountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)
	err = client.DeleteZoneLevelAccessApplication(zoneID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	assert.NoError(t, err)
}

func TestRevokeAccessApplicationTokens(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
      "success": true,
      "errors": [],
      "messages": []
    }
    `)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db/revoke-tokens", handler)
	err := client.RevokeAccessApplicationTokens(accountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db/revoke-tokens", handler)
	err = client.RevokeZoneLevelAccessApplicationTokens(zoneID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	assert.NoError(t, err)
}

func TestAccessApplicationWithCORS(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
      "success": true,
      "errors": [],
      "messages": [

      ],
      "result":{
        "id": "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
        "created_at": "2014-01-01T05:20:00.12345Z",
        "updated_at": "2014-01-01T05:20:00.12345Z",
        "aud": "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
        "name": "Admin Site",
        "domain": "test.example.com/admin",
        "session_duration": "24h",
        "cors_headers": {
          "allowed_methods": [
            "GET"
          ],
          "allowed_origins": [
            "https://example.com"
          ],
          "allow_all_headers": true,
          "max_age": -1
        }
      }
    }
    `)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := AccessApplication{
		ID:              "480f4f69-1a28-4fdd-9240-1ed29f0ac1db",
		CreatedAt:       &createdAt,
		UpdatedAt:       &updatedAt,
		AUD:             "737646a56ab1df6ec9bddc7e5ca84eaf3b0768850f3ffb5d74f1534911fe3893",
		Name:            "Admin Site",
		Domain:          "test.example.com/admin",
		SessionDuration: "24h",
		CorsHeaders: &AccessApplicationCorsHeaders{
			AllowedMethods:  []string{"GET"},
			AllowedOrigins:  []string{"https://example.com"},
			AllowAllHeaders: true,
			MaxAge:          -1,
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err := client.AccessApplication(accountID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/480f4f69-1a28-4fdd-9240-1ed29f0ac1db", handler)

	actual, err = client.ZoneLevelAccessApplication(zoneID, "480f4f69-1a28-4fdd-9240-1ed29f0ac1db")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
