package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRevokeAccessUserTokens(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"result": true
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/organizations/revoke_user", handler)

	AccessUserEmail := AccessUserEmail{Email: "test@example.com"}

	err := client.RevokeAccessUserTokens(context.Background(), testAccountID, AccessUserEmail)

	assert.NoError(t, err)
}

func TestRevokeZoneLevelAccessUserTokens(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"result": true
		}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/access/organizations/revoke_user", handler)

	AccessUserEmail := AccessUserEmail{Email: "test@example.com"}

	err := client.RevokeZoneLevelAccessUserTokens(context.Background(), testZoneID, AccessUserEmail)

	assert.NoError(t, err)
}
