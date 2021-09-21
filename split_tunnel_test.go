package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitTunnelIncludeHost(t *testing.T) {
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
					"host": "*.example.com",
					"description": "default"
			 }
			]
		}
		`)
	}

	want := []SplitTunnel{{
		Host: "*.example.com",
		Description: "default",
	}}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/include", handler)

	actual, err := client.ListSplitTunnel(context.Background(), testAccountID, "include")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSplitTunnelIncludeAddress(t *testing.T) {
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
					"address": "10.0.0.0/8",
					"description": "default address"
			 }
			]
		}
		`)
	}

	want := []SplitTunnel{{
		Address: "10.0.0.0/8",
		Description: "default address",
	}}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/include", handler)

	actual, err := client.ListSplitTunnel(context.Background(), testAccountID, "include")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateSplitTunnelInclude(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"address": "10.0.0.0/8",
					"description": "default address"
			 },
			 {
				  "address": "100.64.0.0/10",
				  "description": "second default address"
		   },
			 {
				  "host": "*.example.com",
				  "description": "example host name"
		   }
			]
		}
		`)
	}

	tunnels := []SplitTunnel{
		{
		Address: "10.0.0.0/8",
		Description: "default address",
	  },
		{
			Address: "100.64.0.0/10",
			Description: "second default address",
		},
		{
			Host: "*.example.com",
			Description: "example host name",
		},
}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/include", handler)

	actual, err := client.UpdateSplitTunnel(context.Background(), testAccountID, "include", tunnels)

	if assert.NoError(t, err) {
		assert.Equal(t, tunnels, actual)
	}
}

func TestSplitTunnelExcludeHost(t *testing.T) {
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
					"host": "*.example.com",
					"description": "default"
			 }
			]
		}
		`)
	}

	want := []SplitTunnel{{
		Host: "*.example.com",
		Description: "default",
	}}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/exclude", handler)

	actual, err := client.ListSplitTunnel(context.Background(), testAccountID, "exclude")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSplitTunnelExcludeAddress(t *testing.T) {
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
					"address": "10.0.0.0/8",
					"description": "default address"
			 }
			]
		}
		`)
	}

	want := []SplitTunnel{{
		Address: "10.0.0.0/8",
		Description: "default address",
	}}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/exclude", handler)

	actual, err := client.ListSplitTunnel(context.Background(), testAccountID, "exclude")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateSplitTunnelExclude(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"address": "10.0.0.0/8",
					"description": "default address"
			 },
			 {
				  "address": "100.64.0.0/10",
				  "description": "second default address"
		   },
			 {
				  "host": "*.example.com",
				  "description": "example host name"
		   }
			]
		}
		`)
	}

	tunnels := []SplitTunnel{
		{
		Address: "10.0.0.0/8",
		Description: "default address",
	  },
		{
			Address: "100.64.0.0/10",
			Description: "second default address",
		},
		{
			Host: "*.example.com",
			Description: "example host name",
		},
}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/exclude", handler)

	actual, err := client.UpdateSplitTunnel(context.Background(), testAccountID, "exclude", tunnels)

	if assert.NoError(t, err) {
		assert.Equal(t, tunnels, actual)
	}
}
