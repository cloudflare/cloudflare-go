package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCertificateAuthoritiesHostnameAssociations(t *testing.T) {
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
				"admin.example.com",
				"foobar.example.com"
			]
		}`)
	}

	hostnameAssociations := ListCertificateAuthoritiesHostnameAssociationsParams{
		MTLSCertificateID: "72ef4d06-4752-4493-a60a-7421470fd585",
	}

	want := []HostnameAssociation{
		"admin.example.com",
		"foobar.example.com",
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	actual, err := client.ListCertificateAuthoritiesHostnameAssociations(context.Background(), testZoneRC, hostnameAssociations)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateCertificateAuthoritiesHostnameAssociations(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				"admin.example.com",
				"foobar.example.com"
			]
		}`)
	}

	hostnameAssociations := UpdateCertificateAuthoritiesHostnameAssociationsParams{
		Hostnames: []HostnameAssociation{
			"admin.example.com",
			"foobar.example.com",
		},
	}

	want := []HostnameAssociation{
		"admin.example.com",
		"foobar.example.com",
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	actual, err := client.UpdateCertificateAuthoritiesHostnameAssociations(context.Background(), testZoneRC, hostnameAssociations)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
