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

func TestCreateZoneLockdown(t *testing.T) {
	setup()
	defer teardown()

	input := ZoneLockdown{
		Description: "Restrict access to these endpoints to requests from a known IP address",
		URLs:        []string{"api.mysite.com/some/endpoint*"},
		Configurations: []ZoneLockdownConfig{{
			Target: "ip",
			Value:  "198.51.100.4",
		}},
		Paused: false,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s")

		var v ZoneLockdown
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
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"paused": false,
				"description": "Restrict access to these endpoints to requests from a known IP address",
				"urls": [
					"api.mysite.com/some/endpoint*"
				],
				"configurations": [
				  {
					  "target": "ip",
					  "value": "198.51.100.4"
				  }
				]
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/firewall/lockdowns", handler)

	actual, err := client.CreateZoneLockdown(context.Background(), testZoneID, input)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := &ZoneLockdownResponse{
		Result: ZoneLockdown{
			ID:             "372e67954025e0ba6aaa6d586b9e0b59",
			Description:    input.Description,
			URLs:           input.URLs,
			Configurations: input.Configurations,
			Paused:         input.Paused,
			CreatedOn:      &createdOn,
			ModifiedOn:     &modifiedOn,
		},
		Response: Response{
			Success:  true,
			Errors:   []ResponseInfo{},
			Messages: []ResponseInfo{},
		},
	}
	assert.Equal(t, want, actual)
}

func TestUpdateZoneLockdownt(t *testing.T) {
	setup()
	defer teardown()

	input := ZoneLockdown{
		Description: "Restrict access to these endpoints to requests from a known IP address",
		URLs:        []string{"api.mysite.com/some/endpoint*"},
		Configurations: []ZoneLockdownConfig{{
			Target: "ip",
			Value:  "198.51.100.4",
		}},
		Paused: false,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s")

		var v ZoneLockdown
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
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"paused": false,
				"description": "Restrict access to these endpoints to requests from a known IP address",
				"urls": [
					"api.mysite.com/some/endpoint*"
				],
				"configurations": [
				  {
					  "target": "ip",
					  "value": "198.51.100.4"
				  }
				]
			}
		}`)
	}

	zoneLockdownID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+testZoneID+"/firewall/lockdowns/"+zoneLockdownID, handler)

	actual, err := client.UpdateZoneLockdown(context.Background(), testZoneID, zoneLockdownID, input)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := &ZoneLockdownResponse{
		Result: ZoneLockdown{
			ID:             zoneLockdownID,
			Description:    input.Description,
			URLs:           input.URLs,
			Configurations: input.Configurations,
			Paused:         input.Paused,
			CreatedOn:      &createdOn,
			ModifiedOn:     &modifiedOn,
		},
		Response: Response{
			Success:  true,
			Errors:   []ResponseInfo{},
			Messages: []ResponseInfo{},
		},
	}
	assert.Equal(t, want, actual)
}

func TestDeleteZoneLockdown(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s")

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

	zoneLockdownID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+testZoneID+"/firewall/lockdowns/"+zoneLockdownID, handler)

	actual, err := client.DeleteZoneLockdown(context.Background(), testZoneID, zoneLockdownID)
	require.NoError(t, err)

	want := &ZoneLockdownResponse{
		Result: ZoneLockdown{
			ID: zoneLockdownID,
		},
		Response: Response{
			Success:  true,
			Errors:   []ResponseInfo{},
			Messages: []ResponseInfo{},
		},
	}
	assert.Equal(t, want, actual)
}

func TestZoneLockdown(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s")

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "372e67954025e0ba6aaa6d586b9e0b59",
				"created_on": "2014-01-01T05:20:00Z",
				"modified_on": "2014-01-01T05:20:00Z",
				"paused": false,
				"description": "Restrict access to these endpoints to requests from a known IP address",
				"urls": [
					"api.mysite.com/some/endpoint*"
				],
				"configurations": [
				  {
					  "target": "ip",
					  "value": "198.51.100.4"
				  }
				]
			}
		}`)
	}

	zoneLockdownID := "372e67954025e0ba6aaa6d586b9e0b59"

	mux.HandleFunc("/zones/"+testZoneID+"/firewall/lockdowns/"+zoneLockdownID, handler)

	actual, err := client.ZoneLockdown(context.Background(), testZoneID, zoneLockdownID)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := &ZoneLockdownResponse{
		Result: ZoneLockdown{
			ID:          zoneLockdownID,
			Description: "Restrict access to these endpoints to requests from a known IP address",
			URLs:        []string{"api.mysite.com/some/endpoint*"},
			Configurations: []ZoneLockdownConfig{{
				Target: "ip",
				Value:  "198.51.100.4",
			}},
			Paused:     false,
			CreatedOn:  &createdOn,
			ModifiedOn: &modifiedOn,
		},
		Response: Response{
			Success:  true,
			Errors:   []ResponseInfo{},
			Messages: []ResponseInfo{},
		},
	}
	assert.Equal(t, want, actual)
}

func TestListZoneLockdowns(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s")
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
			  {
				  "id": "372e67954025e0ba6aaa6d586b9e0b59",
				  "created_on": "2014-01-01T05:20:00Z",
				  "modified_on": "2014-01-01T05:20:00Z",
				  "paused": false,
				  "description": "Restrict access to these endpoints to requests from a known IP address",
				  "urls": [
					  "api.mysite.com/some/endpoint*"
				  ],
				  "configurations": [
					{
						"target": "ip",
						"value": "198.51.100.4"
					}
				  ]
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

	mux.HandleFunc("/zones/"+testZoneID+"/firewall/lockdowns", handler)

	actual, err := client.ListZoneLockdowns(context.Background(), testZoneID, 1)
	require.NoError(t, err)

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	want := &ZoneLockdownListResponse{
		Result: []ZoneLockdown{{
			ID:          "372e67954025e0ba6aaa6d586b9e0b59",
			Description: "Restrict access to these endpoints to requests from a known IP address",
			URLs:        []string{"api.mysite.com/some/endpoint*"},
			Configurations: []ZoneLockdownConfig{{
				Target: "ip",
				Value:  "198.51.100.4",
			}},
			Paused:     false,
			CreatedOn:  &createdOn,
			ModifiedOn: &modifiedOn,
		}},
		Response: Response{
			Success:  true,
			Errors:   []ResponseInfo{},
			Messages: []ResponseInfo{},
		},
		ResultInfo: ResultInfo{
			Page:    1,
			PerPage: 20,
			Count:   1,
			Total:   2000,
		},
	}
	assert.Equal(t, want, actual)
}
