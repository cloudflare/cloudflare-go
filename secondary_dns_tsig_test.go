package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecondaryDNSTSIG(t *testing.T) {
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
				"id": "69cd1e104af3e6ed3cb344f263fd0d5a",
				"name": "tsig.customer.cf.",
				"secret": "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
				"algo": "hmac-sha512."
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/secondary_dns/tsigs/69cd1e104af3e6ed3cb344f263fd0d5a", handler)

	want := SecondaryZoneDNSTSIG{
		ID:     "69cd1e104af3e6ed3cb344f263fd0d5a",
		Name:   "tsig.customer.cf.",
		Secret: "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
		Algo:   "hmac-sha512.",
	}

	actual, err := client.GetSecondaryDNSTSIG(context.Background(), "01a7362d577a6c3019a474fd6f485823", "69cd1e104af3e6ed3cb344f263fd0d5a")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSecondaryDNSTSIGs(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "69cd1e104af3e6ed3cb344f263fd0d5a",
					"name": "tsig.customer.cf.",
					"secret": "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
					"algo": "hmac-sha512."
				}
			]
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/secondary_dns/tsigs", handler)

	want := []SecondaryZoneDNSTSIG{{
		ID:     "69cd1e104af3e6ed3cb344f263fd0d5a",
		Name:   "tsig.customer.cf.",
		Secret: "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
		Algo:   "hmac-sha512.",
	}}

	actual, err := client.SecondaryDNSTSIGs(context.Background(), "01a7362d577a6c3019a474fd6f485823")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateSecondaryDNSZoneTSIG(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "69cd1e104af3e6ed3cb344f263fd0d5a",
				"name": "tsig.customer.cf.",
				"secret": "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
				"algo": "hmac-sha512."
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/secondary_dns/tsigs", handler)

	want := SecondaryZoneDNSTSIG{
		ID:     "69cd1e104af3e6ed3cb344f263fd0d5a",
		Name:   "tsig.customer.cf.",
		Secret: "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
		Algo:   "hmac-sha512.",
	}

	actual, err := client.CreateSecondaryDNSZoneTSIG(context.Background(), "01a7362d577a6c3019a474fd6f485823", SecondaryZoneDNSTSIG{
		Name:   "tsig.customer.cf.",
		Secret: "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c",
		Algo:   "hmac-sha512.",
	})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

