package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateKeyless(t *testing.T) {
	setup()
	defer teardown()

	input := KeylessSSLCreateRequest{
		Host:         "example.com",
		Port:         24008,
		Certificate:  "-----BEGIN CERTIFICATE----- MIIDtTCCAp2g1v2tdw= -----END CERTIFICATE-----",
		Name:         "example.com Keyless SSL",
		BundleMethod: "optimal",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		var v KeylessSSLCreateRequest
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, input, v)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "4d2844d2ce78891c34d0b6c0535a291e",
				"name": "example.com Keyless SSL",
				"host": "example.com",
				"port": 24008,
				"status": "active",
				"enabled": false,
				"permissions": [
					"#ssl:read",
					"#ssl:edit"
				],
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z"
			}
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/keyless_certificates", handler)

	actual, err := client.CreateKeyless(context.Background(), zoneID, input)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := KeylessSSL{
		ID:          "4d2844d2ce78891c34d0b6c0535a291e",
		Name:        input.Name,
		Host:        input.Host,
		Port:        input.Port,
		Status:      "active",
		Enabled:     false,
		Permissions: []string{"#ssl:read", "#ssl:edit"},
		CreatedOn:   createdOn,
		ModifiedOn:  modifiedOn,
	}
	assert.Equal(t, want, actual)
}

func TestListKeyless(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":[
			   {
				  "id":"4d2844d2ce78891c34d0b6c0535a291e",
				  "name":"example.com Keyless SSL",
				  "host":"example.com",
				  "port":24008,
				  "status":"active",
				  "enabled":false,
				  "permissions":[
					 "#ssl:read",
					 "#ssl:edit"
				  ],
				  "created_on":"2014-01-01T05:20:00Z",
				  "modified_on":"2014-01-01T05:20:00Z"
			   }
			]
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/keyless_certificates", handler)

	actual, err := client.ListKeyless(context.Background(), zoneID)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := []KeylessSSL{{
		ID:          "4d2844d2ce78891c34d0b6c0535a291e",
		Name:        "example.com Keyless SSL",
		Host:        "example.com",
		Port:        24008,
		Status:      "active",
		Enabled:     false,
		Permissions: []string{"#ssl:read", "#ssl:edit"},
		CreatedOn:   createdOn,
		ModifiedOn:  modifiedOn,
	}}
	assert.Equal(t, want, actual)
}

func TestKeyless(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":{
			   "id":"4d2844d2ce78891c34d0b6c0535a291e",
			   "name":"example.com Keyless SSL",
			   "host":"example.com",
			   "port":24008,
			   "status":"active",
			   "enabled":false,
			   "permissions":[
				  "#ssl:read",
				  "#ssl:edit"
			   ],
			   "created_on":"2014-01-01T05:20:00Z",
			   "modified_on":"2014-01-01T05:20:00Z"
			}
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/keyless_certificates/"+"4d2844d2ce78891c34d0b6c0535a291e", handler)

	actual, err := client.Keyless(context.Background(), zoneID, "4d2844d2ce78891c34d0b6c0535a291e")
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := KeylessSSL{
		ID:          "4d2844d2ce78891c34d0b6c0535a291e",
		Name:        "example.com Keyless SSL",
		Host:        "example.com",
		Port:        24008,
		Status:      "active",
		Enabled:     false,
		Permissions: []string{"#ssl:read", "#ssl:edit"},
		CreatedOn:   createdOn,
		ModifiedOn:  modifiedOn,
	}
	assert.Equal(t, want, actual)
}

func TestUpdateKeyless(t *testing.T) {
	setup()
	defer teardown()

	enabled := false
	input := KeylessSSLUpdateRequest{
		Host:    "example.com",
		Name:    "example.com Keyless SSL",
		Port:    24008,
		Enabled: &enabled,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		var v KeylessSSLUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, input, v)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":{
			   "id":"4d2844d2ce78891c34d0b6c0535a291e",
			   "name":"example.com Keyless SSL",
			   "host":"example.com",
			   "port":24008,
			   "status":"active",
			   "enabled":false,
			   "permissions":[
				  "#ssl:read",
				  "#ssl:edit"
			   ],
			   "created_on":"2014-01-01T05:20:00Z",
			   "modified_on":"2014-01-01T05:20:00Z"
			}
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/keyless_certificates/"+"4d2844d2ce78891c34d0b6c0535a291e", handler)

	actual, err := client.UpdateKeyless(context.Background(), zoneID, "4d2844d2ce78891c34d0b6c0535a291e", input)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := KeylessSSL{
		ID:          "4d2844d2ce78891c34d0b6c0535a291e",
		Name:        input.Name,
		Host:        input.Host,
		Port:        input.Port,
		Status:      "active",
		Enabled:     enabled,
		Permissions: []string{"#ssl:read", "#ssl:edit"},
		CreatedOn:   createdOn,
		ModifiedOn:  modifiedOn,
	}
	assert.Equal(t, want, actual)
}

func TestDeleteKeyless(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":{
			   "id":"4d2844d2ce78891c34d0b6c0535a291e"
			}
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/keyless_certificates/"+"4d2844d2ce78891c34d0b6c0535a291e", handler)

	err := client.DeleteKeyless(context.Background(), zoneID, "4d2844d2ce78891c34d0b6c0535a291e")
	require.NoError(t, err)
}
