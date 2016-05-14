package cloudflare

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_UserDetails(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
"success": true,
"errors": [],
"messages": [],
"result": {
    "id": "1",
    "email": "cloudflare@example.com",
    "first_name": "Jane",
    "last_name": "Smith",
    "username": "cloudflare12345",
    "telephone": "+1 (650) 319 8930",
    "country": "US",
    "zipcode": "94107",
    "created_on": "2009-07-01T00:00:00Z",
    "modified_on": "2016-05-06T20:32:00Z",
    "two_factor_authentication_enabled": true,
    "betas": ["mirage_forever"]
  }
}`)
	})

	user, err := client.UserDetails()

	createdOn, _ := time.Parse(time.RFC3339, "2009-07-01T00:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2016-05-06T20:32:00Z")

	want := User{
		ID:         "1",
		Email:      "cloudflare@example.com",
		FirstName:  "Jane",
		LastName:   "Smith",
		Username:   "cloudflare12345",
		Telephone:  "+1 (650) 319 8930",
		Country:    "US",
		Zipcode:    "94107",
		CreatedOn:  createdOn,
		ModifiedOn: modifiedOn,
		TwoFA:      true,
		Betas:      []string{"mirage_forever"},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, user, want)
	}
}
