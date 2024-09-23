package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeedBrainSettings(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
		    "id": "speed_brain",
		    "value": "on",
		    "editable": true,
		    "modified_on": "2014-01-01T05:20:00.12345Z"
		  }
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/settings/speed_brain", handler)

	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	editable := true
	want := SpeedBrainSetting{
		ID:         "speed_brain",
		Value:      "on",
		Editable:   &editable,
		ModifiedOn: &modifiedOn,
	}

	actual, err := client.GetSpeedBrainSettings(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestPatchSpeedBrainSettings(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
		    "id": "speed_brain",
		    "value": "off",
		    "editable": true,
		    "modified_on": "2014-01-01T05:20:00.12345Z"
		  }
		}`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/settings/speed_brain", handler)

	modifiedOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	editable := true
	want := SpeedBrainSetting{
		ID:         "speed_brain",
		Value:      "off",
		Editable:   &editable,
		ModifiedOn: &modifiedOn,
	}

	actual, err := client.PatchSpeedBrainSettings(context.Background(), testZoneID, "off")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
