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

func TestCreateDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	priority := uint16(10)
	proxied := false
	input := DNSRecord{
		Type:     "A",
		Name:     "example.com",
		Content:  "198.51.100.4",
		TTL:      120,
		Priority: &priority,
		Proxied:  &proxied,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		var v DNSRecord
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, input, v)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"type": "A",
				"name": "example.com",
				"content": "198.51.100.4",
				"proxiable": true,
				"proxied": false,
				"ttl": 120,
				"locked": false,
				"zone_id": "d56084adb405e0b7e32c52321bf07be6",
				"zone_name": "example.com",
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"data": {},
				"meta": {
					"auto_added": true,
					"source": "primary"
				}
			}
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/dns_records", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := &DNSRecordResponse{
		Result: DNSRecord{
			ID:         "372e67954025e0ba6aaa6d586b9e0b59",
			Type:       input.Type,
			Name:       input.Name,
			Content:    input.Content,
			Proxiable:  true,
			Proxied:    input.Proxied,
			TTL:        input.TTL,
			ZoneID:     zoneID,
			ZoneName:   "example.com",
			CreatedOn:  createdOn,
			ModifiedOn: modifiedOn,
			Data:       map[string]interface{}{},
			Meta: map[string]interface{}{
				"auto_added": true,
				"source":     "primary",
			},
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	actual, err := client.CreateDNSRecord(context.Background(), zoneID, input)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestDNSRecords(t *testing.T) {
	setup()
	defer teardown()

	input := DNSRecord{
		Name:    "example.com",
		Type:    "A",
		Content: "198.51.100.4",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "100", r.URL.Query().Get("per_page"))
		assert.Equal(t, input.Name, r.URL.Query().Get("name"))
		assert.Equal(t, input.Type, r.URL.Query().Get("type"))
		assert.Equal(t, input.Content, r.URL.Query().Get("content"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "372e67954025e0ba6aaa6d586b9e0b59",
					"type": "A",
					"name": "example.com",
					"content": "198.51.100.4",
					"proxiable": true,
					"proxied": false,
					"ttl": 120,
					"locked": false,
					"zone_id": "d56084adb405e0b7e32c52321bf07be6",
					"zone_name": "example.com",
					"created_on": "2014-01-01T05:20:00Z",
					"modified_on": "2014-01-01T05:20:00Z",
					"data": {},
					"meta": {
						"auto_added": true,
						"source": "primary"
					}
				}
			],
			"result_info": {
				"page": 1,
				"total_pages": 1
			}
		}`)
	}

	mux.HandleFunc("/zones/"+zoneID+"/dns_records", handler)

	proxied := false
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := []DNSRecord{{
		ID:         "372e67954025e0ba6aaa6d586b9e0b59",
		Type:       "A",
		Name:       input.Name,
		Content:    input.Content,
		Proxiable:  true,
		Proxied:    &proxied,
		TTL:        120,
		Locked:     false,
		ZoneID:     zoneID,
		ZoneName:   "example.com",
		CreatedOn:  createdOn,
		ModifiedOn: modifiedOn,
		Data:       map[string]interface{}{},
		Meta: map[string]interface{}{
			"auto_added": true,
			"source":     "primary",
		},
	}}

	actual, err := client.DNSRecords(context.Background(), zoneID, input)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"type": "A",
				"name": "example.com",
				"content": "198.51.100.4",
				"proxiable": true,
				"proxied": false,
				"ttl": 120,
				"locked": false,
				"zone_id": "d56084adb405e0b7e32c52321bf07be6",
				"zone_name": "example.com",
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"data": {},
				"meta": {
					"auto_added": true,
					"source": "primary"
				}
			}
		}`)
	}

	dnsRecordID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+zoneID+"/dns_records/"+dnsRecordID, handler)

	proxied := false
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := DNSRecord{
		ID:         dnsRecordID,
		Type:       "A",
		Name:       "example.com",
		Content:    "198.51.100.4",
		Proxiable:  true,
		Proxied:    &proxied,
		TTL:        120,
		ZoneID:     zoneID,
		ZoneName:   "example.com",
		CreatedOn:  createdOn,
		ModifiedOn: modifiedOn,
		Data:       map[string]interface{}{},
		Meta: map[string]interface{}{
			"auto_added": true,
			"source":     "primary",
		},
	}

	actual, err := client.DNSRecord(context.Background(), zoneID, dnsRecordID)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestUpdateDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	proxied := false
	input := DNSRecord{
		Type:    "A",
		Name:    "example.com",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
	}

	handleUpdateDNSRecord := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		var v DNSRecord
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, input, v)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"type": "A",
				"name": "example.com",
				"content": "198.51.100.4",
				"proxiable": true,
				"proxied": false,
				"ttl": 120,
				"locked": false,
				"zone_id": "d56084adb405e0b7e32c52321bf07be6",
				"zone_name": "example.com",
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"data": {},
				"meta": {
					"auto_added": true,
					"source": "primary"
				}
			}
		}`)
	}

	handleDNSRecord := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"type": "A",
				"name": "example.com",
				"content": "198.51.100.4",
				"proxiable": true,
				"proxied": false,
				"ttl": 120,
				"locked": false,
				"zone_id": "d56084adb405e0b7e32c52321bf07be6",
				"zone_name": "example.com",
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"data": {},
				"meta": {
					"auto_added": true,
					"source": "primary"
				}
			}
		}`)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleDNSRecord(w, r)
			return
		}

		if r.Method == http.MethodPatch {
			handleUpdateDNSRecord(w, r)
			return
		}

		assert.Failf(t, "Expected method 'GET' or `PATCH`, got %s", r.Method)
	}

	dnsRecordID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+zoneID+"/dns_records/"+dnsRecordID, handler)

	err := client.UpdateDNSRecord(context.Background(), zoneID, dnsRecordID, input)
	require.NoError(t, err)
}

func TestDeleteDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59"
			}
		}`)
	}

	dnsRecordID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+zoneID+"/dns_records/"+dnsRecordID, handler)

	err := client.DeleteDNSRecord(context.Background(), zoneID, dnsRecordID)
	require.NoError(t, err)
}
