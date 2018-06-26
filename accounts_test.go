package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedAccountStruct = Account{
	ID:   "01a7362d577a6c3019a474fd6f485823",
	Name: "Cloudflare Demo",
	Settings: AccountSettings{
		EnforceTwoFactor: false,
	},
}

func TestListAccounts(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "01a7362d577a6c3019a474fd6f485823",
					"name": "Cloudflare Demo",
					"settings": {
						"enforce_twofactor": false
					}
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

	mux.HandleFunc("/accounts", handler)
	want := []Account{expectedAccountStruct}

	actual, _, err := client.ListAccounts(PaginationOptions{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccount(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "01a7362d577a6c3019a474fd6f485823",
				"name": "Cloudflare Demo",
				"settings": {
					"enforce_twofactor": false
				}
			},
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823", handler)
	want := expectedAccountStruct

	actual, _, err := client.Account("01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
