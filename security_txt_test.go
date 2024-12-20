package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestSecurityTXT_Get(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "result": {
			  "acknowledgments": [
				"https://example.com/hall-of-fame.html"
			  ],
			  "canonical": [
				"https://www.example.com/.well-known/security.txt"
			  ],
			  "contact": [
				"mailto:security@example.com",
				"tel:+1-201-555-0123",
				"https://example.com/security-contact.html"
			  ],
			  "enabled": true,
			  "encryption": [
				"https://example.com/pgp-key.txt",
				"dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY",
				"openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f"
			  ],
			  "expires": "2020-12-02T20:24:07.776073Z",
			  "hiring": [
				"https://example.com/jobs.html"
			  ],
			  "policy": [
				"https://example.com/disclosure-policy.html"
			  ],
			  "preferredLanguages": "en, es, fr"
			},
		  "success": true,
		  "errors": [],
		  "messages": []
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/security-center/securitytxt", handler)

	expires, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	want := SecurityTXT{
		Acknowledgements: []string{"https://example.com/hall-of-fame.html"},
		Canonical:        []string{"https://www.example.com/.well-known/security.txt"},
		Contact: []string{
			"mailto:security@example.com",
			"tel:+1-201-555-0123",
			"https://example.com/security-contact.html",
		},
		Enabled: true,
		Encryption: []string{
			"https://example.com/pgp-key.txt",
			"dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY",
			"openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f",
		},
		Expires:            &expires,
		Hiring:             []string{"https://example.com/jobs.html"},
		Policy:             []string{"https://example.com/disclosure-policy.html"},
		PreferredLanguages: "en, es, fr",
	}

	_, err := client.GetSecurityTXT(context.Background(), ZoneIdentifier(""))
	assert.Error(t, err, "Expected error for missing Zone ID")

	actual, err := client.GetSecurityTXT(context.Background(), ZoneIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSecurityTXT_Update(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result":
        {
			"acknowledgments": [
			  "https://example.com/hall-of-fame.html"
			],
			"canonical": [
			  "https://www.example.com/.well-known/security.txt"
			],
			"contact": [
			  "mailto:security@example.com",
			  "tel:+1-201-555-0123",
			  "https://example.com/security-contact.html"
			],
			"enabled": false,
			"encryption": [
			  "https://example.com/pgp-key.txt",
			  "dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY",
			  "openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f"
			],
			"expires": "2020-12-02T20:24:07.776073Z",
			"hiring": [
			  "https://example.com/jobs.html"
			],
			"policy": [
			  "https://example.com/disclosure-policy.html"
			],
			"preferredLanguages": "en, es, fr"
		  },
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/security-center/securitytxt", handler)

	expires, _ := time.Parse(time.RFC3339, "2020-12-02T20:24:07.776073Z")

	want := SecurityTXT{
		Acknowledgements: []string{"https://example.com/hall-of-fame.html"},
		Canonical:        []string{"https://www.example.com/.well-known/security.txt"},
		Contact: []string{
			"mailto:security@example.com",
			"tel:+1-201-555-0123",
			"https://example.com/security-contact.html",
		},
		Enabled: false,
		Encryption: []string{
			"https://example.com/pgp-key.txt",
			"dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY",
			"openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f",
		},
		Expires:            &expires,
		Hiring:             []string{"https://example.com/jobs.html"},
		Policy:             []string{"https://example.com/disclosure-policy.html"},
		PreferredLanguages: "en, es, fr",
	}

	_, err := client.UpdateSecurityTXT(context.Background(), ZoneIdentifier(""), want)
	assert.Error(t, err, "Expected error for missing Zone ID")

	actual, err := client.UpdateSecurityTXT(context.Background(), ZoneIdentifier(testZoneID), want)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSecurityTXT_DELETE(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
			{
			  "errors": [],
			  "messages": [],
			  "success": true
			}
			`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/security-center/securitytxt", handler)
	err := client.DeleteSecurityTXT(context.Background(), ZoneIdentifier(""))
	assert.Error(t, err, "Expected error for missing Zone ID")

	err = client.DeleteSecurityTXT(context.Background(), ZoneIdentifier(testZoneID))
	assert.NoError(t, err)
}
