package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSSLVerification(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": [
				{
					"certificate_status": "pending_validation",
					"cert_pack_uuid": "4449b368-d10f-4a73-9d27-f3f2752443df",
					"validation_method": "http",
					"validation_type": "dv",
					"verification_info": {
						"http_url": "http:\/\/example.com\/.well-known\/pki-validation\/ca3-0fdf5b3dc4d0428fa81ac381a8c4e333.txt",
						"http_body": "ca3-ff2d513629514e3b90675a982f9f0644"
					},
					"hostname": "example.com"
				}
			],
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/ssl/verification", handler)

	want := make([]ZoneSSLVerification, 1, 4)
	want[0] = ZoneSSLVerification{
		CertificateStatus: "pending_validation",
		CertPackUUID:      "4449b368-d10f-4a73-9d27-f3f2752443df",
		ValidationMethod:  "http",
		ValidationType:    "dv",
		VerificationInfo: struct {
			HTTPURL  string `json:"http_url"`
			HTTPBody string `json:"http_body"`
		}{"http://example.com/.well-known/pki-validation/ca3-0fdf5b3dc4d0428fa81ac381a8c4e333.txt", "ca3-ff2d513629514e3b90675a982f9f0644"},
		Hostname: "example.com",
	}

	actual, err := client.ListSSLVerification("023e105f4ecef8ad9ca31a8372d0c353")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	_, err = client.ListSSLVerification("bar")
	assert.Error(t, err)
}
