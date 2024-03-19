package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListHostnameAssociations(t *testing.T) {
	setup()
	defer teardown()

	params := HostnameAssociationParams{
		CertificateID: "af7149d5-1ca0-4768-8bb1-d50a51c7d845",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"mtls_certificate_id": "af7149d5-1ca0-4768-8bb1-d50a51c7d845",
				"hostnames": [
					"test.example.com"
				]
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	want := HostnameAssociation{
		CertificateID: "af7149d5-1ca0-4768-8bb1-d50a51c7d845",
		Hostnames:     []string{"test.example.com"},
	}

	actual, err := client.ListHostnameAssociations(context.Background(), ZoneIdentifier(testZoneID), params)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListHostnameAssociationsWithoutID(t *testing.T) {
	setup()
	defer teardown()

	params := HostnameAssociationParams{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"hostnames": [
					"test.example.com"
				]
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/certificate_authorities/hostname_associations", handler)

	want := HostnameAssociation{
		Hostnames: []string{"test.example.com"},
	}

	actual, err := client.ListHostnameAssociations(context.Background(), ZoneIdentifier(testZoneID), params)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
