package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

const (
	jobID                       = 1
	serverLogpushJobDescription = `{
		"id": %d,
		"dataset": "http_requests",
    "enabled": false,
	"name": "example.com",
    "logpull_options": "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339",
	"destination_conf": "s3://mybucket/logs?region=us-west-2",
	"last_complete": "%[2]s",
	"last_error": "%[2]s",
	"error_message": "test"
  }
`
	serverLogpushGetOwnershipChallengeDescription = `{
    "filename": "logs/challenge-filename.txt",
	"valid": true,
	"message": ""
  }
`
	serverLogpushGetOwnershipChallengeInvalidResponseDescription = `{
    "filename": "logs/challenge-filename.txt",
	"valid": false,
	"message": "destination is invalid"
  }
`
)

var (
	testLogpushTimestamp     = time.Now().UTC()
	expectedLogpushJobStruct = LogpushJob{
		ID:              jobID,
		Dataset:         "http_requests",
		Enabled:         false,
		Name:            "example.com",
		LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339",
		DestinationConf: "s3://mybucket/logs?region=us-west-2",
		LastComplete:    &testLogpushTimestamp,
		LastError:       &testLogpushTimestamp,
		ErrorMessage:    "test",
	}
	expectedLogpushGetOwnershipChallengeStruct = LogpushGetOwnershipChallenge{
		Filename: "logs/challenge-filename.txt",
		Valid:    true,
		Message:  "",
	}
	expectedLogpushGetOwnershipChallengeInvalidResponseStruct = LogpushGetOwnershipChallenge{
		Filename: "logs/challenge-filename.txt",
		Valid:    false,
		Message:  "destination is invalid",
	}
	expectedUpdatedLogpushJobStruct = LogpushJob{
		ID:              jobID,
		Dataset:         "http_requests",
		Enabled:         true,
		Name:            "updated.com",
		LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp",
		DestinationConf: "gs://mybucket/logs",
		LastComplete:    &testLogpushTimestamp,
		LastError:       &testLogpushTimestamp,
		ErrorMessage:    "test",
	}
)

func TestLogpushJobs(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": [
			%s
		  ],
		  "success": true,
		  "errors": null,
		  "messages": null,
		  "result_info": {
			"page": 1,
			"per_page": 25,
			"count": 1,
			"total_count": 1
		  }
		}
		`, fmt.Sprintf(serverLogpushJobDescription, jobID, testLogpushTimestamp.Format(time.RFC3339Nano)))
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/jobs", handler)
	want := []LogpushJob{expectedLogpushJobStruct}

	actual, err := client.LogpushJobs(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetLogpushJob(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": null,
		  "messages": null
		}
		`, fmt.Sprintf(serverLogpushJobDescription, jobID, testLogpushTimestamp.Format(time.RFC3339Nano)))
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/jobs/"+strconv.Itoa(jobID), handler)
	want := expectedLogpushJobStruct

	actual, err := client.LogpushJob(context.Background(), testZoneID, jobID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateLogpushJob(t *testing.T) {
	setup()
	defer teardown()
	newJob := LogpushJob{
		Enabled:         false,
		Name:            "example.com",
		LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339",
		DestinationConf: "s3://mybucket/logs?region=us-west-2",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": null,
		  "messages": null
		}
		`, fmt.Sprintf(serverLogpushJobDescription, jobID, testLogpushTimestamp.Format(time.RFC3339Nano)))
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/jobs", handler)
	want := &expectedLogpushJobStruct

	actual, err := client.CreateLogpushJob(context.Background(), testZoneID, newJob)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateLogpushJob(t *testing.T) {
	setup()
	defer teardown()
	updatedJob := LogpushJob{
		Enabled:         true,
		Name:            "updated.com",
		LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp",
		DestinationConf: "gs://mybucket/logs",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": null,
		  "messages": null
		}
		`, fmt.Sprintf(serverLogpushJobDescription, jobID, testLogpushTimestamp.Format(time.RFC3339Nano)))
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/jobs/"+strconv.Itoa(jobID), handler)

	err := client.UpdateLogpushJob(context.Background(), testZoneID, jobID, updatedJob)
	assert.NoError(t, err)
}

func TestDeleteLogpushJob(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "result": null,
		  "success": true,
		  "errors": null,
		  "messages": null
		}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/jobs/"+strconv.Itoa(jobID), handler)

	err := client.DeleteLogpushJob(context.Background(), testZoneID, jobID)
	assert.NoError(t, err)
}

func TestGetLogpushOwnershipChallenge(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": null,
		  "messages": null
		}
		`, serverLogpushGetOwnershipChallengeDescription)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/ownership", handler)

	want := &expectedLogpushGetOwnershipChallengeStruct

	actual, err := client.GetLogpushOwnershipChallenge(context.Background(), testZoneID, "destination_conf")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetLogpushOwnershipChallengeWithInvalidResponse(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "result": %s,
		  "success": true,
		  "errors": null,
		  "messages": null
		}
		`, serverLogpushGetOwnershipChallengeInvalidResponseDescription)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logpush/ownership", handler)
	_, err := client.GetLogpushOwnershipChallenge(context.Background(), testZoneID, "destination_conf")

	assert.Error(t, err)
}

func TestValidateLogpushOwnershipChallenge(t *testing.T) {
	testCases := map[string]struct {
		isValid bool
	}{
		"ownership is valid": {
			isValid: true,
		},
		"ownership is not valid": {
			isValid: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			setup()
			defer teardown()

			handler := func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				fmt.Fprintf(w, `{
				  "result": {
					"valid": %v
				  },
				  "success": true,
				  "errors": null,
				  "messages": null
				}
				`, tc.isValid)
			}

			mux.HandleFunc("/zones/"+testZoneID+"/logpush/ownership/validate", handler)

			actual, err := client.ValidateLogpushOwnershipChallenge(context.Background(), testZoneID, "destination_conf", "ownership_challenge")
			if assert.NoError(t, err) {
				assert.Equal(t, tc.isValid, actual)
			}
		})
	}
}

func TestCheckLogpushDestinationExists(t *testing.T) {
	testCases := map[string]struct {
		exists bool
	}{
		"destination exists": {
			exists: true,
		},
		"destination does not exists": {
			exists: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			setup()
			defer teardown()

			handler := func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				fmt.Fprintf(w, `{
				  "result": {
					"exists": %v
				  },
				  "success": true,
				  "errors": null,
				  "messages": null
				}
				`, tc.exists)
			}

			mux.HandleFunc("/zones/"+testZoneID+"/logpush/validate/destination/exists", handler)

			actual, err := client.CheckLogpushDestinationExists(context.Background(), testZoneID, "destination_conf")
			if assert.NoError(t, err) {
				assert.Equal(t, tc.exists, actual)
			}
		})
	}
}
