package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//List Tunnel Routes (GET accounts/:account_identifier/teamnet/routes)
func TestAPI_ListTunnelRoutes(t *testing.T) {
	setup(UsingAccount(testAccountID))
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
			    "network": "ff01::/32",
				"tunnel_id": "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
				"tunnel_name": "blog",
				"comment": "Example comment for this route",
				"created_at": "2021-01-25T18:22:34.317854Z",
				"deleted_at": "2021-01-25T18:22:34.317854Z"
              }
            ]
          }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/teamnet/routes", handler)

	ts, _ := time.Parse(time.RFC3339Nano, "2021-01-25T18:22:34.317854Z")
	want := []TunnelRoute{
		{
			"ff01::/32",
			"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
			"blog",
			"Example comment for this route",
			ts,
			ts,
		},
	}

	got, err := client.TunnelRoutes(context.Background())

	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestAPI_TunnelRouteForIp(t *testing.T) {
	setup(UsingAccount(testAccountID))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "network": "ff01::/32",
			  "tunnel_id": "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
			  "tunnel_name": "blog",
			  "comment": "Example comment for this route",
			  "created_at": "2021-01-25T18:22:34.317854Z",
			  "deleted_at": "2021-01-25T18:22:34.317854Z"
            }
          }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/teamnet/routes/ip/10.1.0.137", handler)

	ts, _ := time.Parse(time.RFC3339Nano, "2021-01-25T18:22:34.317854Z")
	want := TunnelRoute{
		"ff01::/32",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		"blog",
		"Example comment for this route",
		ts,
		ts,
	}

	got, err := client.TunnelRouteForIp(context.Background(), "10.1.0.137")

	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestAPI_CreateTunnelRouteWithComment(t *testing.T) {
	setup(UsingAccount(testAccountID))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		var responseContent map[string]string
		json.Unmarshal(body, &responseContent)
		assert.Equal(t, testTunnelID, responseContent["tunnel_id"])
		assert.Equal(t, "foo", responseContent["comment"])

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "network": "ff01::/32",
			  "tunnel_id": "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
			  "tunnel_name": "blog",
			  "comment": "Example comment for this route",
			  "created_at": "2021-01-25T18:22:34.317854Z",
			  "deleted_at": "2021-01-25T18:22:34.317854Z"
            }
          }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/teamnet/routes/network/10.0.0.0/16", handler)
	_, err := client.CreateTunnelRouteWithComment(context.Background(), testTunnelID, "10.0.0.0/16", "foo")
	assert.NoError(t, err)
}

func TestAPI_UpdateTunnelRouteWithComment(t *testing.T) {
	setup(UsingAccount(testAccountID))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		var responseContent map[string]string
		json.Unmarshal(body, &responseContent)
		assert.Equal(t, testTunnelID, responseContent["tunnel_id"])
		assert.Equal(t, "foo", responseContent["comment"])
		assert.Equal(t, "192.168.0.0/16", responseContent["new_network"])

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "network": "ff01::/32",
			  "tunnel_id": "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
			  "tunnel_name": "blog",
			  "comment": "Example comment for this route",
			  "created_at": "2021-01-25T18:22:34.317854Z",
			  "deleted_at": "2021-01-25T18:22:34.317854Z"
            }
          }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/teamnet/routes/network/10.0.0.0/16", handler)
	_, err := client.UpdateTunnelRouteWithComment(context.Background(), testTunnelID, "10.0.0.0/16", "192.168.0.0/16", "foo")
	assert.NoError(t, err)
}

//Delete Route (DELETE accounts/:account_identifier/teamnet/routes/network/:ip_network)
