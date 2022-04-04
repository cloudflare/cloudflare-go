package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListWorkerCronTriggers(t *testing.T) {
	setup(UsingAccount("9a7806061c88ada191ed06f989cc3dac"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"schedules": [
					{
						"cron": "*/30 * * * *",
						"created_on": "2017-01-01T00:00:00Z",
						"modified_on": "2017-01-01T00:00:00Z"
					}
				]
			}
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/example-script/schedules", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	want := []WorkerCronTrigger{{
		Cron:       "*/30 * * * *",
		ModifiedOn: &modifiedOn,
		CreatedOn:  &createdOn,
	}}

	actual, err := client.ListWorkerCronTriggers(context.Background(), testAccountID, "example-script")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateWorkerCronTriggers(t *testing.T) {
	setup(UsingAccount("9a7806061c88ada191ed06f989cc3dac"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"schedules": [
					{
						"cron": "*/30 * * * *",
						"created_on": "2017-01-01T00:00:00Z",
						"modified_on": "2017-01-01T00:00:00Z"
					}
				]
			}
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/example-script/schedules", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	want := []WorkerCronTrigger{{
		Cron:       "*/30 * * * *",
		ModifiedOn: &modifiedOn,
		CreatedOn:  &createdOn,
	}}

	actual, err := client.UpdateWorkerCronTriggers(context.Background(), testAccountID, "example-script", want)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
