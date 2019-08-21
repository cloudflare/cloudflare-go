package cloudflare

import (
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
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
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

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/organizations", handler)
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

	actual, _, err := client.AccessOrganization("01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccessOrganization(t *testing.T) {
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

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/organizations", handler)
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

	actual, err := client.CreateAccessOrganization("01a7362d577a6c3019a474fd6f485823", want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateAccessOrganization(t *testing.T) {
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

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/organizations", handler)
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

	actual, err := client.UpdateAccessOrganization("01a7362d577a6c3019a474fd6f485823", want)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
