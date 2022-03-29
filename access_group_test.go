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
			map[string]interface{}{"gateway": map[string]interface{}{"integration_uid": "9ca005c0-a695-402d-ba96-18d8d09c4fbf"}},
		},
		Exclude: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
			map[string]interface{}{"gateway": map[string]interface{}{"integration_uid": "9ca005c0-a695-402d-ba96-18d8d09c4fbf"}},
		},
		Require: []interface{}{
			map[string]interface{}{"email": map[string]interface{}{"email": "test@example.com"}},
			map[string]interface{}{"gateway": map[string]interface{}{"integration_uid": "9ca005c0-a695-402d-ba96-18d8d09c4fbf"}},
		},
	}

	result = `{
		"id": "699d98642c564d2e855e9661899b7252",
		"created_at": "2014-01-01T05:20:00.12345Z",
		"updated_at": "2014-01-01T05:20:00.12345Z",
		"name": "Allow devs",
		"include": [
			{
				"email": {
					"email": "test@example.com"
				}
			},
			{
				"gateway": {
					"integration_uid": "9ca005c0-a695-402d-ba96-18d8d09c4fbf"
				}
			}
		],
		"exclude": [
			{
				"email": {
					"email": "test@example.com"
				}
			},
			{
				"gateway": {
					"integration_uid": "9ca005c0-a695-402d-ba96-18d8d09c4fbf"
				}
			}
		],
		"require": [
			{
				"email": {
					"email": "test@example.com"
				}
			},
			{
				"gateway": {
					"integration_uid": "9ca005c0-a695-402d-ba96-18d8d09c4fbf"
				}
			}
		]
	}`
)

func TestAccessGroups(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf(`{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				%s
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000
			}
		}
		`, result))
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/groups", handler)

	actual, _, err := client.AccessGroups(context.Background(), testAccountID, pageOptions)

	if assert.NoError(t, err) {
		assert.Equal(t, []AccessGroup{expectedAccessGroup}, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/groups", handler)

	actual, _, err = client.ZoneLevelAccessGroups(context.Background(), testZoneID, pageOptions)

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
		fmt.Fprintf(w, fmt.Sprintf(`{
			"success": true,
			"errors": [],
			"messages": [],
			"result": %s
		}
		`, result))
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/groups/"+accessGroupID, handler)

	actual, err := client.AccessGroup(context.Background(), testAccountID, accessGroupID)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/groups/"+accessGroupID, handler)

	actual, err = client.ZoneLevelAccessGroup(context.Background(), testZoneID, accessGroupID)

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
		fmt.Fprintf(w, fmt.Sprintf(`{
			"success": true,
			"errors": [],
			"messages": [],
			"result": %s
		}`, result))
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/groups", handler)

	accessGroup := AccessGroup{
		Name: "Allow devs",
		Include: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
			AccessGroupGateway{struct {
				ID string `json:"integration_uid"`
			}{ID: "9ca005c0-a695-402d-ba96-18d8d09c4fbf"}},
		},
		Exclude: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
			AccessGroupGateway{struct {
				ID string `json:"integration_uid"`
			}{ID: "9ca005c0-a695-402d-ba96-18d8d09c4fbf"}},
		},
		Require: []interface{}{
			AccessGroupEmail{struct {
				Email string `json:"email"`
			}{Email: "test@example.com"}},
			AccessGroupGateway{struct {
				ID string `json:"integration_uid"`
			}{ID: "9ca005c0-a695-402d-ba96-18d8d09c4fbf"}},
		},
	}

	actual, err := client.CreateAccessGroup(context.Background(), testAccountID, accessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/groups", handler)

	actual, err = client.CreateZoneLevelAccessGroup(context.Background(), testZoneID, accessGroup)

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
		fmt.Fprintf(w, fmt.Sprintf(`{
			"success": true,
			"errors": [],
			"messages": [],
			"result": %s
		}`, result))
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/groups/"+accessGroupID, handler)
	actual, err := client.UpdateAccessGroup(context.Background(), testAccountID, expectedAccessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/groups/"+accessGroupID, handler)
	actual, err = client.UpdateZoneLevelAccessGroup(context.Background(), testZoneID, expectedAccessGroup)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedAccessGroup, actual)
	}
}

func TestUpdateAccessGroupWithMissingID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessGroup(context.Background(), testAccountID, AccessGroup{})
	assert.EqualError(t, err, "access group ID cannot be empty")

	_, err = client.UpdateZoneLevelAccessGroup(context.Background(), testZoneID, AccessGroup{})
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

	mux.HandleFunc("/accounts/"+testAccountID+"/access/groups/"+accessGroupID, handler)
	err := client.DeleteAccessGroup(context.Background(), testAccountID, accessGroupID)

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+testZoneID+"/access/groups/"+accessGroupID, handler)
	err = client.DeleteZoneLevelAccessGroup(context.Background(), testZoneID, accessGroupID)

	assert.NoError(t, err)
}
