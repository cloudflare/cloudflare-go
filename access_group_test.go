package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	accessGroupID = "699d98642c564d2e855e9661899b7252"

	expectedAccessGroup = AccessGroup{
		ID:        "699d98642c564d2e855e9661899b7252",
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Name:      "Allow devs",
		Include: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
		},
		Exclude: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
		},
		Require: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
		},
	}
)

func TestAccessGroups(t *testing.T) {
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
				{
					"id": "699d98642c564d2e855e9661899b7252",
					"created_at": "2014-01-01T05:20:00.12345Z",
					"updated_at": "2014-01-01T05:20:00.12345Z",
					"name": "Allow devs",
					"include": [
						{
							"email": {
								"email": "test@example.com"
							}
						}
					],
					"exclude": [
						{
							"email": {
								"email": "test@example.com"
							}
						}
					],
					"require": [
						{
							"email": {
								"email": "test@example.com"
							}
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
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/groups", handler)

	actual, _, err := client.AccessGroups(context.Background(), accountID, pageOptions)

	if assert.NoError(t, err) {
		assert.Equal(t, []AccessGroup{expectedAccessGroup}, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/groups", handler)

	actual, _, err = client.ZoneLevelAccessGroups(context.Background(), zoneID, pageOptions)

	if assert.NoError(t, err) {
		assert.Equal(t, []AccessGroup{expectedAccessGroup}, actual)
	}
}

func TestAccessGroup(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "699d98642c564d2e855e9661899b7252",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"exclude": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"require": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/groups/"+accessGroupID, handler)

	actual, err := client.AccessGroup(context.Background(), accountID, accessGroupID)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/groups/"+accessGroupID, handler)

	actual, err = client.ZoneLevelAccessGroup(context.Background(), zoneID, accessGroupID)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}
}

func TestCreateAccessGroup(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "699d98642c564d2e855e9661899b7252",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"exclude": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"require": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/groups", handler)

	accessGroup := AccessGroup{
		Name: "Allow devs",
		Include: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
		},
		Exclude: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
		},
		Require: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
		},
	}

	actual, err := client.CreateAccessGroup(context.Background(), accountID, accessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/groups", handler)

	actual, err = client.CreateZoneLevelAccessGroup(context.Background(), zoneID, accessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}
}

func TestUpdateAccessGroup(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "699d98642c564d2e855e9661899b7252",
				"created_at": "2014-01-01T05:20:00.12345Z",
				"updated_at": "2014-01-01T05:20:00.12345Z",
				"name": "Allow devs",
				"include": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"exclude": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				],
				"require": [
					{
						"email": {
							"email": "test@example.com"
						}
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/groups/"+accessGroupID, handler)
	actual, err := client.UpdateAccessGroup(context.Background(), accountID, expectedAccessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/groups/"+accessGroupID, handler)
	actual, err = client.UpdateZoneLevelAccessGroup(context.Background(), zoneID, expectedAccessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}
}

func TestUpdateAccessGroupWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessGroup(context.Background(), accountID, AccessGroup{})
	assert.EqualError(t, err, "access group ID cannot be empty")

	_, err = client.UpdateZoneLevelAccessGroup(context.Background(), zoneID, AccessGroup{})
	assert.EqualError(t, err, "access group ID cannot be empty")
}

func TestDeleteAccessGroup(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "699d98642c564d2e855e9661899b7252"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/groups/"+accessGroupID, handler)
	err := client.DeleteAccessGroup(context.Background(), accountID, accessGroupID)

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+zoneID+"/access/groups/"+accessGroupID, handler)
	err = client.DeleteZoneLevelAccessGroup(context.Background(), zoneID, accessGroupID)

	assert.NoError(t, err)
}
