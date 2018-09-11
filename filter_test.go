package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pageOpts = PaginationOptions{
	PerPage: 25,
	Page:    1,
}

func TestFilter(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
				"id": "b7ff25282d394be7b945e23c7106ce8a",
				"paused": false,
				"description": "Login from office",
				"expression": "ip.src eq 127.0.0.1"
			},
			"success": true,
			"errors": null,
			"messages": null
		}
		`)
	}

	mux.HandleFunc("/zones/d56084adb405e0b7e32c52321bf07be6/filters/b7ff25282d394be7b945e23c7106ce8a", handler)
	want := Filter{
		ID:          "b7ff25282d394be7b945e23c7106ce8a",
		Paused:      false,
		Description: "Login from office",
		Expression:  "ip.src eq 127.0.0.1",
	}

	actual, err := client.Filter("d56084adb405e0b7e32c52321bf07be6", "b7ff25282d394be7b945e23c7106ce8a")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

