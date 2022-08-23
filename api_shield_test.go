package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIShield(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
	"success" : true,
	"errors": [],
	"messages": [],
	"result": {
		"auth_id_characteristics": [
			{
				"type": "header",
				"name": "test-header"
			},
			{
				"type": "cookie",
				"name": "test-cookie"
			}
		]
	}
}
		`)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/api_gateway/configuration", handler)

	var authChars []AuthIdCharacteristics
	authChars = append(authChars, AuthIdCharacteristics{Type: "header", Name: "test-header"})
	authChars = append(authChars, AuthIdCharacteristics{Type: "cookie", Name: "test-cookie"})

	want := APIShield{
		AuthIdCharacteristics: authChars,
	}

	actual, _, err := client.GetAPIShieldConfiguration(context.Background(), "01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
