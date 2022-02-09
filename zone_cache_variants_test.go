package cloudflare

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZoneSingleCacheSetting(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		// JSON data from: https://api.cloudflare.com/#zone-cache-settings-get-variants-setting
		_, _ = fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "id": "variants",
			  "modified_on": "2014-01-01T05:20:00.12345Z",
			  "value": {
				"avif": [
				  "image/webp",
				  "image/jpeg"
				],
				"bmp": [
				  "image/webp",
				  "image/jpeg"
				]
			  }
			}
		  }`)
	}

	testZoneId := "foo"
	mux.HandleFunc("/zones/"+testZoneId+"/cache/variants", handler)

	want := ZoneCacheVariants{
		ID:         "variants",
		ModifiedOn: "2014-01-01T05:20:00.12345Z",
		Value: map[string][]string{
			"avif": {
				"image/webp",
				"image/jpeg",
			},
			"bmp": {
				"image/webp",
				"image/jpeg",
			},
		},
	}

	actual, err := client.ZoneCacheVariants(context.Background(), testZoneId)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateZoneSingleCacheSetting(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if assert.NoError(t, err) {
			assert.JSONEq(t, `{"id":"","value":{"avif":["image/webp","image/jpeg"],"bmp":["image/webp","image/jpeg"]}}`, string(b))
		}

		w.Header().Set("content-type", "application/json")
		// JSON data from: https://api.cloudflare.com/#zone-cache-settings-change-variants-setting
		_, _ = fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "id": "variants",
			  "modified_on": "2014-01-01T05:20:00.12345Z",
			  "value": {
				"avif": [
				  "image/webp",
				  "image/jpeg"
				],
				"bmp": [
				  "image/webp",
				  "image/jpeg"
				]
			  }
			}
		  }`)
	}

	testZoneId := "foo"
	mux.HandleFunc("/zones/"+testZoneId+"/cache/variants", handler)

	want := ZoneCacheVariants{
		ID:         "variants",
		ModifiedOn: "2014-01-01T05:20:00.12345Z",
		Value: map[string][]string{
			"avif": {
				"image/webp",
				"image/jpeg",
			},
			"bmp": {
				"image/webp",
				"image/jpeg",
			},
		},
	}

	zoneCacheSetting := ZoneCacheVariants{Value: map[string][]string{
		"avif": {
			"image/webp",
			"image/jpeg",
		},
		"bmp": {
			"image/webp",
			"image/jpeg",
		}}}

	actual, err := client.UpdateZoneCacheVariants(context.Background(), testZoneId, zoneCacheSetting)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteZoneSingleCacheSetting(t *testing.T) {
	setup()
	defer teardown()

	apiCalled := false

	handler := func(w http.ResponseWriter, r *http.Request) {
		apiCalled = true
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		// JSON data from: https://api.cloudflare.com/#zone-cache-settings-delete-variants-setting
		_, _ = fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "id": "variants",
			  "modified_on": "2014-01-01T05:20:00.12345Z"
			}
		  }`)
	}

	testZoneId := "foo"
	mux.HandleFunc("/zones/"+testZoneId+"/cache/variants", handler)

	err := client.DeleteZoneCacheVariants(context.Background(), testZoneId)

	assert.NoError(t, err)
	assert.True(t, apiCalled)
}
