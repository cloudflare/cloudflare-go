package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTunnels(t *testing.T) {
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
					"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
					"name": "blog",
					"created_at": "2009-11-10T23:00:00Z",
					"deleted_at": "2009-11-10T23:00:00Z",
					"connections": [
						{
							"colo_name": "DFW",
							"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
							"is_pending_reconnect": false,
							"client_id": "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
							"client_version": "2022.2.0",
							"opened_at": "2021-01-25T18:22:34.317854Z",
							"origin_ip": "85.12.78.6"
						}
					]
				}
			]
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel", handler)

	createdAt, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	deletedAt, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	want := []Tunnel{{
		ID:        "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name:      "blog",
		CreatedAt: &createdAt,
		DeletedAt: &deletedAt,
		Connections: []TunnelConnection{{
			ColoName:           "DFW",
			ID:                 "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			IsPendingReconnect: false,
			ClientID:           "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
			ClientVersion:      "2022.2.0",
			OpenedAt:           "2021-01-25T18:22:34.317854Z",
			OriginIP:           "85.12.78.6",
		}},
	}}

	actual, err := client.Tunnels(context.Background(), TunnelListParams{AccountID: testAccountID, UUID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestTunnel(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":{
				"id":"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name":"blog",
				"created_at":"2009-11-10T23:00:00Z",
				"deleted_at":"2009-11-10T23:00:00Z",
				"connections": [
					{
						"colo_name": "DFW",
						"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
						"is_pending_reconnect": false,
						"client_id": "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
						"client_version": "2022.2.0",
						"opened_at": "2021-01-25T18:22:34.317854Z",
						"origin_ip": "85.12.78.6"
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	createdAt, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	deletedAt, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	want := Tunnel{
		ID:        "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name:      "blog",
		CreatedAt: &createdAt,
		DeletedAt: &deletedAt,
		Connections: []TunnelConnection{{
			ColoName:           "DFW",
			ID:                 "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			IsPendingReconnect: false,
			ClientID:           "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
			ClientVersion:      "2022.2.0",
			OpenedAt:           "2021-01-25T18:22:34.317854Z",
			OriginIP:           "85.12.78.6",
		}},
	}

	actual, err := client.Tunnel(context.Background(), TunnelParams{AccountID: testAccountID, ID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateTunnel(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":{
				"id":"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name":"blog",
				"created_at":"2009-11-10T23:00:00Z",
				"deleted_at":"2009-11-10T23:00:00Z",
				"connections": [
					{
						"colo_name": "DFW",
						"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
						"is_pending_reconnect": false,
						"client_id": "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
						"client_version": "2022.2.0",
						"opened_at": "2021-01-25T18:22:34.317854Z",
						"origin_ip": "85.12.78.6"
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel", handler)

	createdAt, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	deletedAt, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	want := Tunnel{
		ID:        "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Name:      "blog",
		CreatedAt: &createdAt,
		DeletedAt: &deletedAt,
		Connections: []TunnelConnection{{
			ColoName:           "DFW",
			ID:                 "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			IsPendingReconnect: false,
			ClientID:           "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
			ClientVersion:      "2022.2.0",
			OpenedAt:           "2021-01-25T18:22:34.317854Z",
			OriginIP:           "85.12.78.6",
		}},
	}

	actual, err := client.CreateTunnel(context.Background(), TunnelCreateParams{AccountID: testAccountID, Name: "blog", Secret: "notarealsecret"})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateTunnelConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method '%s', got %s", http.MethodPut, r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "config": "{\"warp-routing\": {\"enabled\": true},  \"originRequest\" : {\"connectTimeout\": 10}, \"ingress\" : [ {\"hostname\": \"test.example.com\", \"service\": \"https://localhost:8000\" } , {\"service\": \"http_status:404\"} ]}",
			  "tunnel_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			  "version": 5,
			  "created_at": "2021-01-25T18:22:34.317854Z"
			}
		  }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/configurations", handler)
	want := TunnelConfigurationResult{
		TunnelID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Version:  5,
		Config: TunnelConfiguration{
			Ingress: []UnvalidatedIngressRule{
				{
					Hostname: "test.example.com",
					Service:  "https://localhost:8000",
				},
				{
					Service: "http_status:404",
				},
			},
			WarpRouting: &WarpRoutingConfig{
				Enabled: true,
			},
			OriginRequest: OriginRequestConfig{
				ConnectTimeout: DurationPtr(10),
			},
		}}

	actual, err := client.UpdateTunnelConfiguration(context.Background(), TunnelConfigurationParams{
		AccountID: testAccountID,
		TunnelID:  "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Config: TunnelConfiguration{
			Ingress: []UnvalidatedIngressRule{
				{
					Hostname: "test.example.com",
					Service:  "https://localhost:8000",
				},
				{
					Service: "http_status:404",
				},
			},
			WarpRouting: &WarpRoutingConfig{
				Enabled: true,
			},
			OriginRequest: OriginRequestConfig{
				ConnectTimeout: DurationPtr(10 * time.Second),
			},
		},
	})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetTunnelConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method '%s', got %s", http.MethodGet, r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "config": "{\"warp-routing\": {\"enabled\": true},  \"originRequest\" : {\"connectTimeout\": 10}, \"ingress\" : [ {\"hostname\": \"test.example.com\", \"service\": \"https://localhost:8000\" } , {\"service\": \"http_status:404\"} ]}",
			  "tunnel_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			  "version": 5,
			  "created_at": "2021-01-25T18:22:34.317854Z"
			}
		  }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/configurations", handler)
	want := TunnelConfigurationResult{
		TunnelID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		Version:  5,
		Config: TunnelConfiguration{
			Ingress: []UnvalidatedIngressRule{
				{
					Hostname: "test.example.com",
					Service:  "https://localhost:8000",
				},
				{
					Service: "http_status:404",
				},
			},
			WarpRouting: &WarpRoutingConfig{
				Enabled: true,
			},
			OriginRequest: OriginRequestConfig{
				ConnectTimeout: DurationPtr(10),
			},
		}}

	actual, err := client.GetTunnelConfiguration(context.Background(), GetTunnelConfigurationParams{AccountID: testAccountID, TunnelID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteTunnel(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success":true,
			"errors":[],
			"messages":[],
			"result":{
				"id":"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
				"name":"blog",
				"created_at":"2009-11-10T23:00:00Z",
				"deleted_at":"2009-11-10T23:00:00Z",
				"connections": [
					{
						"colo_name": "DFW",
						"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
						"is_pending_reconnect": false,
						"client_id": "dc6472cc-f1ae-44a0-b795-6b8a0ce29f90",
						"client_version": "2022.2.0",
						"opened_at": "2021-01-25T18:22:34.317854Z",
						"origin_ip": "85.12.78.6"
					}
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel/f174e90a-fafe-4643-bbbc-4a0ed4fc8415", handler)

	err := client.DeleteTunnel(context.Background(), TunnelDeleteParams{AccountID: testAccountID, ID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})
	assert.NoError(t, err)
}

func TestCleanupTunnelConnections(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": []
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/connections", handler)

	err := client.CleanupTunnelConnections(context.Background(), TunnelCleanupParams{AccountID: testAccountID, ID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})
	assert.NoError(t, err)
}

func TestTunnelToken(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": "ZHNraGdhc2RraGFza2hqZGFza2poZGFza2poYXNrZGpoYWtzamRoa2FzZGpoa2FzamRoa2Rhc2po\na2FzamRoa2FqCg=="
		  }
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/cfd_tunnel/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/token", handler)

	token, err := client.TunnelToken(context.Background(), TunnelTokenParams{AccountID: testAccountID, ID: "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})
	assert.NoError(t, err)
	assert.Equal(t, "ZHNraGdhc2RraGFza2hqZGFza2poZGFza2poYXNrZGpoYWtzamRoa2FzZGpoa2FzamRoa2Rhc2po\na2FzamRoa2FqCg==", token)
}
