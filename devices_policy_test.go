package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateDeviceClientCertificatesZone(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": {"enabled": true}
		}`)
	}

	want := DeviceClientCertificatesZone{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: Enabled{true},
	}

	mux.HandleFunc("/zones/"+testZoneID+"/devices/policy/certificates", handler)

	actual, err := client.UpdateDeviceClientCertificatesZone(context.Background(), testZoneID, true)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetDeviceClientCertificatesZone(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": {"enabled": false}
		}`)
	}

	want := DeviceClientCertificatesZone{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: Enabled{false},
	}

	mux.HandleFunc("/zones/"+testZoneID+"/devices/policy/certificates", handler)

	actual, err := client.GetDeviceClientCertificatesZone(context.Background(), testZoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
