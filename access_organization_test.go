package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessOrganization(t *testing.T) {
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
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Widget Corps Internal Applications",
				"auth_domain": "test.cloudflareaccess.com",
				"login_design": {
					"background_color": "#c5ed1b",
					"text_color": "#c5ed1b",
					"logo_path": "https://example.com/logo.png"
				}
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := AccessOrganization{
		Name:       "Widget Corps Internal Applications",
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
		AuthDomain: "test.cloudflareaccess.com",
		LoginDesign: AccessOrganizationLoginDesign{
			BackgroundColor: "#c5ed1b",
			TextColor:       "#c5ed1b",
			LogoPath:        "https://example.com/logo.png",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/organizations", handler)

	actual, _, err := client.AccessOrganization(context.Background(), accountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/organizations", handler)

	actual, _, err = client.ZoneLevelAccessOrganization(context.Background(), zoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessOrganization(t *testing.T) {
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
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Widget Corps Internal Applications",
				"auth_domain": "test.cloudflareaccess.com",
				"login_design": {
					"background_color": "#c5ed1b",
					"text_color": "#c5ed1b",
					"logo_path": "https://example.com/logo.png"
				}
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := AccessOrganization{
		Name:       "Widget Corps Internal Applications",
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
		AuthDomain: "test.cloudflareaccess.com",
		LoginDesign: AccessOrganizationLoginDesign{
			BackgroundColor: "#c5ed1b",
			TextColor:       "#c5ed1b",
			LogoPath:        "https://example.com/logo.png",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/organizations", handler)

	actual, err := client.CreateAccessOrganization(context.Background(), accountID, want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/organizations", handler)

	actual, err = client.CreateZoneLevelAccessOrganization(context.Background(), zoneID, want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateAccessOrganization(t *testing.T) {
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
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Widget Corps Internal Applications",
				"auth_domain": "test.cloudflareaccess.com",
				"login_design": {
					"background_color": "#c5ed1b",
					"text_color": "#c5ed1b",
					"logo_path": "https://example.com/logo.png"
				}
			}
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")

	want := AccessOrganization{
		Name:       "Widget Corps Internal Applications",
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
		AuthDomain: "test.cloudflareaccess.com",
		LoginDesign: AccessOrganizationLoginDesign{
			BackgroundColor: "#c5ed1b",
			TextColor:       "#c5ed1b",
			LogoPath:        "https://example.com/logo.png",
		},
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/organizations", handler)

	actual, err := client.UpdateAccessOrganization(context.Background(), accountID, want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/organizations", handler)

	actual, err = client.UpdateZoneLevelAccessOrganization(context.Background(), zoneID, want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
