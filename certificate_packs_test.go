package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	uploadedOn, _ = time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	expiresOn, _  = time.Parse(time.RFC3339, "2016-01-01T05:20:00Z")

	desiredCertificatePack = CertificatePack{
		ID:                 "3822ff90-ea29-44df-9e55-21300bb9419b",
		Type:               "custom",
		Hosts:              []string{"example.com", "*.example.com", "www.example.com"},
		PrimaryCertificate: 12345678,
		Certificates: []CertificatePackCertificate{{
			ID:              12345678,
			Hosts:           []string{"example.com"},
			Issuer:          "GlobalSign",
			Signature:       "SHA256WithRSA",
			Status:          "active",
			BundleMethod:    "ubiquitous",
			GeoRestrictions: CertificatePackGeoRestrictions{Label: "us"},
			ZoneID:          "023e105f4ecef8ad9ca31a8372d0c353",
			UploadedOn:      uploadedOn,
			ModifiedOn:      uploadedOn,
			ExpiresOn:       expiresOn,
			Priority:        1,
		}},
	}
)

func TestListCertificatePacks(t *testing.T) {
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
      "id": "3822ff90-ea29-44df-9e55-21300bb9419b",
      "type": "custom",
      "hosts": [
        "example.com",
        "*.example.com",
        "www.example.com"
      ],
      "certificates": [
        {
          "id": 12345678,
          "hosts": [
            "example.com"
          ],
          "issuer": "GlobalSign",
          "signature": "SHA256WithRSA",
          "status": "active",
          "bundle_method": "ubiquitous",
          "geo_restrictions": {
            "label": "us"
          },
          "zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
          "uploaded_on": "2014-01-01T05:20:00Z",
          "modified_on": "2014-01-01T05:20:00Z",
          "expires_on": "2016-01-01T05:20:00Z",
          "priority": 1
        }
      ],
      "primary_certificate": 12345678
    }
  ]
}
		`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/certificate_packs", handler)

	want := []CertificatePack{desiredCertificatePack}
	actual, err := client.ListCertificatePacks(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListCertificatePack(t *testing.T) {
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
    "id": "3822ff90-ea29-44df-9e55-21300bb9419b",
    "type": "custom",
    "hosts": [
      "example.com",
      "*.example.com",
      "www.example.com"
    ],
    "certificates": [
      {
        "id": 12345678,
        "hosts": [
          "example.com"
        ],
        "issuer": "GlobalSign",
        "signature": "SHA256WithRSA",
        "status": "active",
        "bundle_method": "ubiquitous",
        "geo_restrictions": {
          "label": "us"
        },
        "zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
        "uploaded_on": "2014-01-01T05:20:00Z",
        "modified_on": "2014-01-01T05:20:00Z",
        "expires_on": "2016-01-01T05:20:00Z",
        "priority": 1
      }
    ],
    "primary_certificate": 12345678
  }
}
		`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/certificate_packs/3822ff90-ea29-44df-9e55-21300bb9419b", handler)

	actual, err := client.CertificatePack(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353", "3822ff90-ea29-44df-9e55-21300bb9419b")

	if assert.NoError(t, err) {
		assert.Equal(t, desiredCertificatePack, actual)
	}
}

func TestCreateCertificatePack(t *testing.T) {
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
    "id": "3822ff90-ea29-44df-9e55-21300bb9419b",
    "type": "custom",
    "hosts": [
      "example.com",
      "*.example.com",
      "www.example.com"
    ],
    "certificates": [
      {
        "id": 12345678,
        "hosts": [
          "example.com"
        ],
        "issuer": "GlobalSign",
        "signature": "SHA256WithRSA",
        "status": "active",
        "bundle_method": "ubiquitous",
        "geo_restrictions": {
          "label": "us"
        },
        "zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
        "uploaded_on": "2014-01-01T05:20:00Z",
        "modified_on": "2014-01-01T05:20:00Z",
        "expires_on": "2016-01-01T05:20:00Z",
        "priority": 1
      }
    ],
    "primary_certificate": 12345678
  }
}
		`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/certificate_packs", handler)

	certificate := CertificatePackRequest{Type: "custom", Hosts: []string{"example.com", "*.example.com", "www.example.com"}}
	actual, err := client.CreateCertificatePack(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353", certificate)

	if assert.NoError(t, err) {
		assert.Equal(t, desiredCertificatePack, actual)
	}
}

func TestCreateAdvancedCertificatePack(t *testing.T) {
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
    "id": "3822ff90-ea29-44df-9e55-21300bb9419b",
    "type": "advanced",
    "hosts": [
      "example.com",
      "*.example.com",
      "www.example.com"
    ],
    "status": "initializing",
    "validation_method": "txt",
    "validity_days": 365,
    "certificate_authority": "digicert",
    "cloudflare_branding": false
  }
}`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/certificate_packs/order", handler)

	certificate := CertificatePackAdvancedCertificate{
		ID:                   "3822ff90-ea29-44df-9e55-21300bb9419b",
		Type:                 "advanced",
		Hosts:                []string{"example.com", "*.example.com", "www.example.com"},
		ValidityDays:         365,
		ValidationMethod:     "txt",
		CertificateAuthority: "digicert",
		CloudflareBranding:   false,
	}

	actual, err := client.CreateAdvancedCertificatePack(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353", certificate)

	if assert.NoError(t, err) {
		assert.Equal(t, certificate, actual)
	}
}

func TestRestartAdvancedCertificateValidation(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "3822ff90-ea29-44df-9e55-21300bb9419b",
    "type": "advanced",
    "hosts": [
      "example.com",
      "*.example.com",
      "www.example.com"
    ],
    "status": "initializing",
    "validation_method": "txt",
    "validity_days": 365,
    "certificate_authority": "digicert",
    "cloudflare_branding": false
  }
}`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/certificate_packs/3822ff90-ea29-44df-9e55-21300bb9419b", handler)

	certificate := CertificatePackAdvancedCertificate{
		ID:                   "3822ff90-ea29-44df-9e55-21300bb9419b",
		Type:                 "advanced",
		Hosts:                []string{"example.com", "*.example.com", "www.example.com"},
		ValidityDays:         365,
		ValidationMethod:     "txt",
		CertificateAuthority: "digicert",
		CloudflareBranding:   false,
	}

	actual, err := client.RestartAdvancedCertificateValidation(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353", "3822ff90-ea29-44df-9e55-21300bb9419b")

	if assert.NoError(t, err) {
		assert.Equal(t, certificate, actual)
	}
}

func TestDeleteCertificatePack(t *testing.T) {
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
    "id": "3822ff90-ea29-44df-9e55-21300bb9419b"
  }
}
		`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/certificate_packs/3822ff90-ea29-44df-9e55-21300bb9419b", handler)

	err := client.DeleteCertificatePack(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353", "3822ff90-ea29-44df-9e55-21300bb9419b")

	assert.NoError(t, err)
}
