package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentScanningEnable(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": null
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/content-upload-scan/enable", handler)

	var want ContentScanningEnableResponse
	want.Success = true
	want.Result = nil
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}

	actual, err := client.ContentScanningEnable(context.Background(), ZoneIdentifier(testZoneID), ContentScanningEnableParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestContentScanningDisable(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": null
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/content-upload-scan/disable", handler)

	var want ContentScanningDisableResponse
	want.Success = true
	want.Result = nil
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}

	actual, err := client.ContentScanningDisable(context.Background(), ZoneIdentifier(testZoneID), ContentScanningDisableParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestContentScanningStatus(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"value": "disabled",
    			"modified": "2024-12-03T13:34:38.376312Z"
			}
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/content-upload-scan/settings", handler)

	var want ContentScanningStatusResponse
	want.Success = true
	want.Result = &ContentScanningStatusResult{
		Value:    "disabled",
		Modified: "2024-12-03T13:34:38.376312Z",
	}
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}

	actual, err := client.ContentScanningStatus(context.Background(), ZoneIdentifier(testZoneID), ContentScanningStatusParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestContentScanningListCustomExpressions(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "a350a054caa840c9becd89c3b4f0195b",
					"payload": "lookup_json_string(http.request.body.raw, \"file\")"
				},
				{
					"id": "4e3b841073914278b7413d9c1e6581e4",
					"payload": "lookup_json_string(http.request.body.raw, \"document\")"
				}
			]
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/content-upload-scan/payloads", handler)

	want := []ContentScanningCustomExpression{
		{
			ID:      "a350a054caa840c9becd89c3b4f0195b",
			Payload: "lookup_json_string(http.request.body.raw, \"file\")",
		},
		{
			ID:      "4e3b841073914278b7413d9c1e6581e4",
			Payload: "lookup_json_string(http.request.body.raw, \"document\")",
		},
	}

	actual, err := client.ContentScanningListCustomExpressions(context.Background(), ZoneIdentifier(testZoneID), ContentScanningListCustomExpressionsParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestContentScanningAddCustomExpressions(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "a350a054caa840c9becd89c3b4f0195b",
					"payload": "lookup_json_string(http.request.body.raw, \"alreadyExisting\")"
				},
				{
					"id": "4e3b841073914278b7413d9c1e6581e4",
					"payload": "lookup_json_string(http.request.body.raw, \"OneNew\")"
				},
				{
					"id": "8f7403465db64e47bbb34bd7191186c6",
					"payload": "lookup_json_string(http.request.body.raw, \"AnotherNew\")"
				}
			]
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/content-upload-scan/payloads", handler)

	want := []ContentScanningCustomExpression{
		{
			ID:      "a350a054caa840c9becd89c3b4f0195b",
			Payload: "lookup_json_string(http.request.body.raw, \"alreadyExisting\")",
		},
		{
			ID:      "4e3b841073914278b7413d9c1e6581e4",
			Payload: "lookup_json_string(http.request.body.raw, \"OneNew\")",
		},
		{
			ID:      "8f7403465db64e47bbb34bd7191186c6",
			Payload: "lookup_json_string(http.request.body.raw, \"AnotherNew\")",
		},
	}
	params := ContentScanningAddCustomExpressionsParams{
		Payloads: []ContentScanningCustomPayload{
			{
				Payload: "lookup_json_string(http.request.body.raw, \"OneNew\")",
			},
			{
				Payload: "lookup_json_string(http.request.body.raw, \"AnotherNew\")",
			},
		},
	}
	actual, err := client.ContentScanningAddCustomExpressions(context.Background(), ZoneIdentifier(testZoneID), params)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestContentScanningDeleteCustomExpression(t *testing.T) {
	setup()
	payloadId := "cafb3307c5cc4c029d6bbd557b9e223a"
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "a350a054caa840c9becd89c3b4f0195b",
					"payload": "lookup_json_string(http.request.body.raw, \"file\")"
				}
			]
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/content-upload-scan/payloads/"+payloadId, handler)

	want := []ContentScanningCustomExpression{
		{
			ID:      "a350a054caa840c9becd89c3b4f0195b",
			Payload: "lookup_json_string(http.request.body.raw, \"file\")",
		},
	}

	actual, err := client.ContentScanningDeleteCustomExpression(context.Background(), ZoneIdentifier(testZoneID), ContentScanningDeleteCustomExpressionsParams{ID: payloadId})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
