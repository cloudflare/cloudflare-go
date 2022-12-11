package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAlwaysUseHTTPS(t *testing.T) {
    setup()
    defer teardown()

    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprintf(w, `{
          "success": true,
          "errors": [],
          "messages": [],
          "result": {
            "id": "always_use_https",
            "value": "on",
            "editable": true,
            "modified_on": "2022-12-11T03:20:00.12345Z"
          }
        }`)
    }
    mux.HandleFunc("/zones/4e2d6095be5f10bbec901abc5136415d/settings/always_use_https", handler)

    modifiedOn, _ := time.Parse(time.RFC3339, "2022-12-11T03:20:00.12345Z")
    want := AlwaysUseHTTPS{
        ID: "always_use_https",
        Value: "on",
        Editable: true,
        ModifiedOn: modifiedOn,
    }

    actual, err := client.GetAlwaysUseHTTPS(context.TODO(), "4e2d6095be5f10bbec901abc5136415d")
    if assert.NoError(t, err) {
        assert.Equal(t, want, actual)
    }
}

func TestSetAlwaysUseHTTPSEnabled(t *testing.T) {
    setup()
    defer teardown()

    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprintf(w, `{
          "success": true,
          "errors": [],
          "messages": [],
          "result": {
            "id": "always_use_https",
            "value": "on",
            "editable": true,
            "modified_on": "2022-12-11T03:20:00.12345Z"
          }
        }`)
    }
    mux.HandleFunc("/zones/4e2d6095be5f10bbec901abc5136415d/settings/always_use_https", handler)

    modifiedOn, _ := time.Parse(time.RFC3339, "2022-12-11T03:20:00.12345Z")
    want := AlwaysUseHTTPS{
        ID: "always_use_https",
        Value: "on",
        Editable: true,
        ModifiedOn: modifiedOn,
    }

    actual, err := client.SetAlwaysUseHTTPS(context.TODO(), "4e2d6095be5f10bbec901abc5136415d", true)
    if assert.NoError(t, err) {
        assert.Equal(t, want, actual)
    }
}

func TestSetAlwaysUseHTTPSDisabled(t *testing.T) {
    setup()
    defer teardown()

    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprintf(w, `{
          "success": true,
          "errors": [],
          "messages": [],
          "result": {
            "id": "always_use_https",
            "value": "off",
            "editable": true,
            "modified_on": "2022-12-11T03:20:00.12345Z"
          }
        }`)
    }
    mux.HandleFunc("/zones/4e2d6095be5f10bbec901abc5136415d/settings/always_use_https", handler)

    modifiedOn, _ := time.Parse(time.RFC3339, "2022-12-11T03:20:00.12345Z")
    want := AlwaysUseHTTPS{
        ID: "always_use_https",
        Value: "off",
        Editable: true,
        ModifiedOn: modifiedOn,
    }

    actual, err := client.SetAlwaysUseHTTPS(context.TODO(), "4e2d6095be5f10bbec901abc5136415d", false)
    if assert.NoError(t, err) {
        assert.Equal(t, want, actual)
    }
}

