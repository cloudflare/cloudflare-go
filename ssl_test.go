package cloudflare

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListSSL(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
          "success": true,
          "errors": [],
          "messages": [],
          "result": [
            {
              "id": "7e7b8deba8538af625850b7b2530034c",
              "hosts": [
                "example.com"
              ],
              "issuer": "GlobalSign",
              "signature": "SHA256WithRSA",
              "status": "active",
              "bundle_method": "ubiquitous",
              "zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
              "uploaded_on": "2014-01-01T05:20:00Z",
              "modified_on": "2014-01-01T05:20:00Z",
              "expires_on": "2016-01-01T05:20:00Z",
              "priority": 1
            }
          ],
          "result_info": {
            "page": 1,
            "per_page": 20,
            "count": 1,
            "total_count": 2000
          }
        }`)
	}

	mux.HandleFunc("/zones/foo/custom_certificates", handler)

	hosts := make([]string, 1, 4)
	hosts[0] = "example.com"
	uploadedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	expiresOn, _ := time.Parse(time.RFC3339, "2016-01-01T05:20:00Z")

	want := make([]ZoneCustomSSL, 1, 4)
	want[0] = ZoneCustomSSL{
		ID:           "7e7b8deba8538af625850b7b2530034c",
		Hosts:        hosts,
		Issuer:       "GlobalSign",
		Signature:    "SHA256WithRSA",
		Status:       "active",
		BundleMethod: "ubiquitous",
		ZoneID:       "023e105f4ecef8ad9ca31a8372d0c353",
		UploadedOn:   uploadedOn,
		ModifiedOn:   modifiedOn,
		ExpiresOn:    expiresOn,
		Priority:     1,
	}

	actual, err := client.ListSSL("foo")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	_, err = client.ListSSL("bar")
	assert.Error(t, err)
}
