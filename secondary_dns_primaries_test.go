package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecondaryDNSPrimary(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "23ff594956f20c2a721606e94745a8aa",
				"ip": "192.0.2.53",
				"port": 53,
				"ixfr_enable": false,
				"tsig_id": "69cd1e104af3e6ed3cb344f263fd0d5a",
				"name": "my-primary-1"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/secondary_dns/primaries/23ff594956f20c2a721606e94745a8aa", handler)
	want := SecondaryDNSPrimary{
		ID:         "23ff594956f20c2a721606e94745a8aa",
		IP:         "192.0.2.53",
		Port:       53,
		IxfrEnable: false,
		TsigID:     "69cd1e104af3e6ed3cb344f263fd0d5a",
		Name:       "my-primary-1",
	}

	actual, err := client.GetSecondaryDNSPrimary(context.Background(), "01a7362d577a6c3019a474fd6f485823", "23ff594956f20c2a721606e94745a8aa")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListSecondaryDNSPrimaries(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [{
				"id": "23ff594956f20c2a721606e94745a8aa",
				"ip": "192.0.2.53",
				"port": 53,
				"ixfr_enable": false,
				"tsig_id": "69cd1e104af3e6ed3cb344f263fd0d5a",
				"name": "my-primary-1"
			}]
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/secondary_dns/primaries", handler)
	want := []SecondaryDNSPrimary{{
		ID:         "23ff594956f20c2a721606e94745a8aa",
		IP:         "192.0.2.53",
		Port:       53,
		IxfrEnable: false,
		TsigID:     "69cd1e104af3e6ed3cb344f263fd0d5a",
		Name:       "my-primary-1",
	}}

	actual, err := client.ListSecondaryDNSPrimaries(context.Background(), "01a7362d577a6c3019a474fd6f485823")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

