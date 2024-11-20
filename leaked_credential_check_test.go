package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeakedCredentialCheckGetStatus(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "enabled": true
			}
		  }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/leaked-credential-checks", handler)

	want := LeakedCredentialCheckStatus{
		Enabled: BoolPtr(true),
	}
	actual, err := client.LeakedCredentialCheckGetStatus(context.Background(), ZoneIdentifier(testZoneID), LeakedCredentialCheckGetStatusParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLeakedCredentialCheckSetStatus(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
			  "enabled": false
			}
		  }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/leaked-credential-checks", handler)

	want := LeakedCredentialCheckStatus{
		Enabled: BoolPtr(false),
	}
	actual, err := client.LeakedCredentialCheckSetStatus(context.Background(), ZoneIdentifier(testZoneID), LeakCredentialCheckSetStatusParams{BoolPtr(false)})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLeakedCredentialCheckListDetections(t *testing.T) {
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
				"id": "18a14bafaa8eb1df04ce683ec18c765e",
				"password": "lookup_json_string(http.request.body.raw, \"secret\")",
				"username": "lookup_json_string(http.request.body.raw, \"user\")"
				}
			]
		  }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/leaked-credential-checks/detections", handler)

	want := []LeakedCredentialCheckDetectionEntry{
		{
			ID:       "18a14bafaa8eb1df04ce683ec18c765e",
			Password: "lookup_json_string(http.request.body.raw, \"secret\")",
			Username: "lookup_json_string(http.request.body.raw, \"user\")",
		},
	}
	actual, err := client.LeakedCredentialCheckListDetections(context.Background(), ZoneIdentifier(testZoneID), LeakedCredentialCheckListDetectionsParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLeakedCredentialCheckCreateDetection(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "18a14bafaa8eb1df04ce683ec18c765e",
				"password": "lookup_json_string(http.request.body.raw, \"secret\")",
				"username": "lookup_json_string(http.request.body.raw, \"user\")"
			}
		  }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/leaked-credential-checks/detections", handler)

	want := LeakedCredentialCheckDetectionEntry{
		ID:       "18a14bafaa8eb1df04ce683ec18c765e",
		Password: "lookup_json_string(http.request.body.raw, \"secret\")",
		Username: "lookup_json_string(http.request.body.raw, \"user\")",
	}
	// POST data
	detectionPattern := LeakedCredentialCheckCreateDetectionParams{
		Username: "lookup_json_string(http.request.body.raw, \"secret\")",
		Password: "lookup_json_string(http.request.body.raw, \"user\")",
	}
	actual, err := client.LeakedCredentialCheckCreateDetection(context.Background(), ZoneIdentifier(testZoneID), detectionPattern)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLeakedCredentialCheckDeleteDetection(t *testing.T) {
	setup()
	detectionId := "cafb3307c5cc4c029d6bbd557b9e223a"
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": []
		  }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/leaked-credential-checks/detections/"+detectionId, handler)

	want := LeakedCredentialCheckDeleteDetectionResponse{}
	want.Success = true
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}
	want.Result = []struct{}{}

	actual, err := client.LeakedCredentialCheckDeleteDetection(context.Background(), ZoneIdentifier(testZoneID), LeakedCredentialCheckDeleteDetectionParams{detectionId})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLeakedCredentialCheckUpdateDetection(t *testing.T) {
	setup()
	detectionId := "18a14bafaa8eb1df04ce683ec18c765e"
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "%s",
				"password": "lookup_json_string(http.request.body.raw, \"secret\")",
				"username": "lookup_json_string(http.request.body.raw, \"user\")"
			}
		  }`, detectionId)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/leaked-credential-checks/detections/"+detectionId, handler)

	want := LeakedCredentialCheckDetectionEntry{
		ID:       detectionId,
		Password: "lookup_json_string(http.request.body.raw, \"secret\")",
		Username: "lookup_json_string(http.request.body.raw, \"user\")",
	}
	actual, err := client.LeakedCredentialCheckUpdateDetection(context.Background(), ZoneIdentifier(testZoneID), LeakedCredentialCheckUpdateDetectionParams{want})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
