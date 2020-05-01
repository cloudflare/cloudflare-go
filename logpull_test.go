package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogpullRentionFlag(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		"errors": [],
		"messages": [],
		"result": {
			"flag": true
		},
		"success": true
	}`)
	}

	mux.HandleFunc("/zones/d56084adb405e0b7e32c52321bf07be6/logs/control/retention/flag", handler)
	want := &LogpullRentionConfiguration{Flag: true}

	actual, err := client.GetLogpullRentionFlag("d56084adb405e0b7e32c52321bf07be6")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSetLogpullRentionFlag(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		"errors": [],
		"messages": [],
		"result": {
			"flag": false
		},
		"success": true
	}`)
	}

	mux.HandleFunc("/zones/d56084adb405e0b7e32c52321bf07be6/logs/control/retention/flag", handler)
	want := &LogpullRentionConfiguration{Flag: false}

	actual, err := client.SetLogpullRentionFlag("d56084adb405e0b7e32c52321bf07be6", false)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
