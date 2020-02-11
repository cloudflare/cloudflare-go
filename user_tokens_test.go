package cloudflare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	userTokenID          = "47aa30b6eb97ecae0518b750d6b142b6"
	userTokenValue       = "very_secret_token_value"
	userTokenName        = "Test UserToken"
	userTokenDescription = `{
		"id": "%s",
		"name": "%s",
		%s
		"status": "active",
		"issued_on": "%s",
		"modified_on": "%s",
		"policies": [%s]
	}`
	userTokenPolicyDescription = `{
		"id": "%s",
		"effect": "allow",
		"resources": {
			"com.cloudflare.api.account.zone.%s": "*"
		},
		"permission_groups": [{
			"id": "%s",
			"name": "DNS Read"
		}]
	}`
)

var (
	expectedUserTokenPolicyStruct = UserTokenPolicy{
		ID:     userTokenID,
		Effect: "allow",
		Resources: map[string]string{
			"com.cloudflare.api.account.zone." + userTokenID: "*",
		},
		PermissionGroups: []UserTokensPermissionGroup{{
			ID:   userTokenID,
			Name: "DNS Read",
		}},
	}

	expectedUserTokenStruct = UserToken{
		ID:         userTokenID,
		Name:       userTokenName,
		Status:     "active",
		IssuedOn:   &testTimestamp,
		ModifiedOn: &testTimestamp,
		Policies:   []UserTokenPolicy{expectedUserTokenPolicyStruct},
	}
)

func TestUserTokensList(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": [
			%s
		  ],
		  "success": true,
		  "errors": [],
		  "messages": []
		}
		`, fmt.Sprintf(userTokenDescription, userTokenID, userTokenName, "",
			testTimestamp.Format(time.RFC3339Nano), testTimestamp.Format(time.RFC3339Nano),
			fmt.Sprintf(userTokenPolicyDescription, userTokenID, userTokenID, userTokenID)))
	}

	mux.HandleFunc("/user/tokens", handler)
	want := []UserToken{expectedUserTokenStruct}

	actual, err := client.ListUserTokens()
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUserTokensGet(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": [],
		  "messages": []
		}
		`, fmt.Sprintf(userTokenDescription, userTokenID, userTokenName, "",
			testTimestamp.Format(time.RFC3339Nano), testTimestamp.Format(time.RFC3339Nano),
			fmt.Sprintf(userTokenPolicyDescription, userTokenID, userTokenID, userTokenID)))
	}

	mux.HandleFunc("/user/tokens/"+userTokenID, handler)
	want := expectedUserTokenStruct

	actual, err := client.UserToken(userTokenID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUserTokensCreate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		defer r.Body.Close()

		userTokenReq, _ := ioutil.ReadAll(r.Body)
		ut := UserToken{}
		json.Unmarshal(userTokenReq, &ut)
		p, _ := json.Marshal(ut.Policies[0])

		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": [],
		  "messages": []
		}
		`, fmt.Sprintf(userTokenDescription, userTokenID, userTokenName, fmt.Sprintf("\"value\": %q,", userTokenValue),
			testTimestamp.Format(time.RFC3339Nano), testTimestamp.Format(time.RFC3339Nano), string(p)))
	}

	mux.HandleFunc("/user/tokens", handler)
	want := expectedUserTokenStruct
	want.Value = userTokenValue

	actual, err := client.CreateUserToken(userTokenName, []UserTokenPolicy{expectedUserTokenPolicyStruct})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUserTokensDelete(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": {
			"id": %q
          },
		  "success": true,
		  "errors": [],
		  "messages": []
		}
		`, userTokenID)
	}

	mux.HandleFunc("/user/tokens/"+userTokenID, handler)

	want := UserTokenID{userTokenID}

	actual, err := client.DeleteUserToken(userTokenID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUserTokensUpdate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "PUT", "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		p := fmt.Sprintf(userTokenPolicyDescription, userTokenID, userTokenID, userTokenID)
		p = fmt.Sprintf("%s, %s", p, p)

		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": [],
		  "messages": []
		}
		`, fmt.Sprintf(userTokenDescription, userTokenID, userTokenName, "",
			testTimestamp.Format(time.RFC3339Nano), testTimestamp.Format(time.RFC3339Nano),
			p))
	}

	mux.HandleFunc("/user/tokens/"+userTokenID, handler)

	want := expectedUserTokenStruct
	want.Policies = []UserTokenPolicy{expectedUserTokenPolicyStruct, expectedUserTokenPolicyStruct}

	actual, err := client.UpdateUserToken(want)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
