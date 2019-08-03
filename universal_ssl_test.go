package cloudflare

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniversalSSLSettingDetails(t *testing.T) {
	setup()
	defer teardown()

	testZoneID := "abcd123"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "enabled": true
			}
		  }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/ssl/universal/settings", handler)

	want := UniversalSSLSetting{
		Enabled: true,
	}

	got, err := client.UniversalSSLSettingDetails(testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestEditUniversalSSLSetting(t *testing.T) {
	setup()
	defer teardown()

	testZoneID := "abcd123"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method, "Expected method 'PATCH', got %s", r.Method)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		assert.Equal(t, `{"enabled":true}`, string(body))

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "enabled": true
			}
		  }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/ssl/universal/settings", handler)

	want := UniversalSSLSetting{
		Enabled: true,
	}

	got, err := client.EditUniversalSSLSetting(testZoneID, want)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestUniversalSSLVerificationDetails(t *testing.T) {
	setup()
	defer teardown()

	testZoneID := "abcd123"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": [{
				"certificate_status": "active",
				"verification_type": "cname",
				"verification_status": true,
				"verification_info": {
					"record_name": "b3b90cfedd89a3e487d3e383c56c4267.example.com",
					"record_target": "6979be7e4cfc9e5c603e31df7efac9cc60fee82d.comodoca.com"
				},
				"brand_check": false,
				"validation_method": "txt",
				"cert_pack_uuid": "a77f8bd7-3b47-46b4-a6f1-75cf98109948"
			}]
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/ssl/verification", handler)

	want := []UniversalSSLVerificationDetails{
		{
			CertificateStatus:  "active",
			VerificationType:   "cname",
			ValidationMethod:   "txt",
			CertPackUUID:       "a77f8bd7-3b47-46b4-a6f1-75cf98109948",
			VerificationStatus: true,
			BrandCheck:         false,
			VerificationInfo: UniversalSSLVerificationInfo{
				RecordName:   "b3b90cfedd89a3e487d3e383c56c4267.example.com",
				RecordTarget: "6979be7e4cfc9e5c603e31df7efac9cc60fee82d.comodoca.com",
			},
		},
	}

	got, err := client.UniversalSSLVerificationDetails(testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
