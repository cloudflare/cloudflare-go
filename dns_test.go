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

func Test_toUTS46ASCII(t *testing.T) {
	tests := map[string]struct {
		domain   string
		expected string
	}{
		"empty stays empty": {
			domain:   "",
			expected: "",
		},
		"unicode gets encoded": {
			domain:   "😺.com",
			expected: "xn--138h.com",
		},
		"unicode gets mapped and encoded": {
			domain:   "ÖBB.at",
			expected: "xn--bb-eka.at",
		},
		"punycode stays punycode": {
			domain:   "xn--138h.com",
			expected: "xn--138h.com",
		},
		"hyphens are not checked": {
			domain:   "s3--s4.com",
			expected: "s3--s4.com",
		},
		"STD3 rules are not enforced": {
			domain:   "℀.com",
			expected: "a/c.com",
		},
		"bidi check is disabled": {
			domain:   "englishﻋﺮﺑﻲ.com",
			expected: "xn--english-gqjzfwd1j.com",
		},
		"invalid joiners are allowed": {
			domain:   "a\u200cb.com",
			expected: "xn--ab-j1t.com",
		},
		"partial results are used despite errors": {
			domain:   "xn--:D.xn--.😺.com",
			expected: "xn--:d..xn--138h.com",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := toUTS46ASCII(tt.domain)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCreateDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	priority := uint16(10)
	proxied := false
	asciiInput := DNSRecord{
		Type:     "A",
		Name:     "xn--138h.example.com",
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
		assert.Equal(t, asciiInput, v)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"type": "A",
				"name": "xn--138h.example.com",
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

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := &DNSRecordResponse{
		Result: DNSRecord{
			ID:         "372e67954025e0ba6aaa6d586b9e0b59",
			Type:       asciiInput.Type,
			Name:       asciiInput.Name,
			Content:    asciiInput.Content,
			Proxiable:  true,
			Proxied:    asciiInput.Proxied,
			TTL:        asciiInput.TTL,
			ZoneID:     testZoneID,
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

	_, err := client.CreateDNSRecord(context.Background(), ZoneIdentifier(""), CreateDNSRecordParams{})
	assert.Equal(t, ErrMissingZoneID, err)

	actual, err := client.CreateDNSRecord(context.Background(), ZoneIdentifier(testZoneID), CreateDNSRecordParams{
		Type:     "A",
		Name:     "😺.example.com",
		Content:  "198.51.100.4",
		TTL:      120,
		Priority: &priority,
		Proxied:  &proxied})
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestDNSRecords(t *testing.T) {
	setup()
	defer teardown()

	asciiInput := DNSRecord{
		Name:    "xn--138h.example.com",
		Type:    "A",
		Content: "198.51.100.4",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, asciiInput.Name, r.URL.Query().Get("name"))
		assert.Equal(t, asciiInput.Type, r.URL.Query().Get("type"))
		assert.Equal(t, asciiInput.Content, r.URL.Query().Get("content"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "372e67954025e0ba6aaa6d586b9e0b59",
					"type": "A",
					"name": "xn--138h.example.com",
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
				"count": 1,
				"page": 1,
				"per_page": 20,
				"total_count": 2000
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records", handler)

	proxied := false
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := []DNSRecord{{
		ID:         "372e67954025e0ba6aaa6d586b9e0b59",
		Type:       "A",
		Name:       asciiInput.Name,
		Content:    asciiInput.Content,
		Proxiable:  true,
		Proxied:    &proxied,
		TTL:        120,
		Locked:     false,
		ZoneID:     testZoneID,
		ZoneName:   "example.com",
		CreatedOn:  createdOn,
		ModifiedOn: modifiedOn,
		Data:       map[string]interface{}{},
		Meta: map[string]interface{}{
			"auto_added": true,
			"source":     "primary",
		},
	}}

	_, _, err := client.ListDNSRecords(context.Background(), ZoneIdentifier(""), ListDNSRecordsParams{})
	assert.Equal(t, ErrMissingZoneID, err)

	actual, _, err := client.ListDNSRecords(context.Background(), ZoneIdentifier(testZoneID), ListDNSRecordsParams{
		Name:    "😺.example.com",
		Type:    "A",
		Content: "198.51.100.4",
	})
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestDNSRecordsSearch(t *testing.T) {
	setup()
	defer teardown()

	recordInput := DNSRecord{
		Name:    "example.com",
		Type:    "A",
		Content: "198.51.100.4",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, recordInput.Name, r.URL.Query().Get("name"))
		assert.Equal(t, recordInput.Type, r.URL.Query().Get("type"))
		assert.Equal(t, recordInput.Content, r.URL.Query().Get("content"))
		assert.Equal(t, "all", r.URL.Query().Get("match"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "type", r.URL.Query().Get("order"))
		assert.Equal(t, "asc", r.URL.Query().Get("direction"))

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
					"proxied": true,
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
					},
					"tags": ["tag1", "tag2extended"]
				}
			],
			"result_info": {
				"count": 1,
				"page": 1,
				"per_page": 20,
				"total_count": 2000
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records", handler)

	proxied := true
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := []DNSRecord{{
		ID:         "372e67954025e0ba6aaa6d586b9e0b59",
		Type:       "A",
		Name:       recordInput.Name,
		Content:    recordInput.Content,
		Proxiable:  true,
		Proxied:    &proxied,
		TTL:        120,
		Locked:     false,
		ZoneID:     testZoneID,
		ZoneName:   "example.com",
		CreatedOn:  createdOn,
		ModifiedOn: modifiedOn,
		Data:       map[string]interface{}{},
		Meta: map[string]interface{}{
			"auto_added": true,
			"source":     "primary",
		},
		Tags: []string{"tag1", "tag2extended"},
	}}

	actual, resultInfo, err := client.ListDNSRecords(context.Background(), ZoneIdentifier(testZoneID), ListDNSRecordsParams{
		ResultInfo: ResultInfo{
			Page: 1,
		},
		Match:     "all",
		Order:     "type",
		Direction: ListDirectionAsc,
		Name:      "example.com",
		Type:      "A",
		Content:   "198.51.100.4",
	})
	require.NoError(t, err)
	assert.Equal(t, 2000, resultInfo.Total)

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
				},
				"comment": "This is a comment",
				"tags": ["tag1", "tag2"]
			}
		}`)
	}

	dnsRecordID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records/"+dnsRecordID, handler)

	proxied := false
	comment := DNSRecordComment("This is a comment")
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
		ZoneID:     testZoneID,
		ZoneName:   "example.com",
		CreatedOn:  createdOn,
		ModifiedOn: modifiedOn,
		Data:       map[string]interface{}{},
		Meta: map[string]interface{}{
			"auto_added": true,
			"source":     "primary",
		},
		Comment: &comment,
		Tags:    []string{"tag1", "tag2"},
	}

	_, err := client.GetDNSRecord(context.Background(), ZoneIdentifier(""), dnsRecordID)
	assert.Equal(t, ErrMissingZoneID, err)

	actual, err := client.GetDNSRecord(context.Background(), ZoneIdentifier(testZoneID), dnsRecordID)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestUpdateDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	proxied := false
	input := DNSRecord{
		ID:      "372e67954025e0ba6aaa6d586b9e0b59",
		Type:    "A",
		Name:    "example.com",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
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

	dnsRecordID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records/"+dnsRecordID, handler)

	err := client.UpdateDNSRecord(context.Background(), ZoneIdentifier(""), UpdateDNSRecordParams{})
	assert.Equal(t, ErrMissingZoneID, err)

	err = client.UpdateDNSRecord(context.Background(), ZoneIdentifier(testZoneID), UpdateDNSRecordParams{
		ID:      dnsRecordID,
		Type:    "A",
		Name:    "example.com",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
	})
	require.NoError(t, err)
}

func TestUpdateDNSRecordWithoutName(t *testing.T) {
	setup()
	defer teardown()

	proxied := false

	asciiInput := DNSRecord{
		ID:      "372e67954025e0ba6aaa6d586b9e0b59",
		Name:    "xn--138h.example.com",
		Type:    "A",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
	}

	handleUpdateDNSRecord := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		var v DNSRecord
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, asciiInput, v)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"type": "A",
				"name": "xn--138h.example.com",
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
				"name": "xn--138h.example.com",
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

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records/"+dnsRecordID, handler)

	err := client.UpdateDNSRecord(context.Background(), ZoneIdentifier(""), UpdateDNSRecordParams{})
	assert.Equal(t, ErrMissingZoneID, err)

	err = client.UpdateDNSRecord(context.Background(), ZoneIdentifier(testZoneID), UpdateDNSRecordParams{
		ID:      dnsRecordID,
		Type:    "A",
		Name:    "xn--138h.example.com",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
	})
	require.NoError(t, err)
}

func TestUpdateDNSRecordWithoutType(t *testing.T) {
	setup()
	defer teardown()

	proxied := false

	completedASCIIInput := DNSRecord{
		Name:    "xn--138h.example.com",
		Type:    "A",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
		ID:      "372e67954025e0ba6aaa6d586b9e0b59",
	}

	handleUpdateDNSRecord := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		var v DNSRecord
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, completedASCIIInput, v)

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

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records/"+dnsRecordID, handler)

	err := client.UpdateDNSRecord(context.Background(), ZoneIdentifier(testZoneID), UpdateDNSRecordParams{
		ID:      dnsRecordID,
		Name:    "😺.example.com",
		Content: "198.51.100.4",
		TTL:     120,
		Proxied: &proxied,
	})
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

	mux.HandleFunc("/zones/"+testZoneID+"/dns_records/"+dnsRecordID, handler)

	err := client.DeleteDNSRecord(context.Background(), ZoneIdentifier(""), dnsRecordID)
	assert.Equal(t, ErrMissingZoneID, err)

	err = client.DeleteDNSRecord(context.Background(), ZoneIdentifier(testZoneID), dnsRecordID)
	require.NoError(t, err)
}
