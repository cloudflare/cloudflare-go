package cloudflare

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
		CreatedOn:  &createdOn,
		ModifiedOn: &modifiedOn,
		TwoFA:      true,
		Betas:      []string{"mirage_forever"},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, user, want)
	}
}

func TestUser_UpdateUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method, "Expected method 'PATCH', got %s", r.Method)
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{"country":"US","first_name":"John","username":"cfuser12345","email":"user@example.com",
                       "last_name": "Appleseed","telephone": "+1 123-123-1234","zipcode": "12345"}`, string(b), "JSON not equal")
		}
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "7c5dae5552338874e5053f2534d2767a",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Appleseed",
    "username": "cfuser12345",
    "telephone": "+1 123-123-1234",
    "country": "US",
    "zipcode": "12345",
    "created_on": "2014-01-01T05:20:00Z",
    "modified_on": "2014-01-01T05:20:00Z",
    "two_factor_authentication_enabled": false
  }
}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00Z")

	userIn := User{
		Email:     "user@example.com",
		FirstName: "John",
		LastName:  "Appleseed",
		Username:  "cfuser12345",
		Telephone: "+1 123-123-1234",
		Country:   "US",
		Zipcode:   "12345",
		TwoFA:     false,
	}

	userOut, err := client.UpdateUser(&userIn)

	want := User{
		ID:         "7c5dae5552338874e5053f2534d2767a",
		Email:      "user@example.com",
		FirstName:  "John",
		LastName:   "Appleseed",
		Username:   "cfuser12345",
		Telephone:  "+1 123-123-1234",
		Country:    "US",
		Zipcode:    "12345",
		CreatedOn:  &createdOn,
		ModifiedOn: &modifiedOn,
		TwoFA:      false,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, userOut, want, "structs not equal")
	}
}
