package cloudflare

import (
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
    "enabled": false,
	"name": "example.com",
    "logpull_options": "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339",
	"destination_conf": "s3://mybucket/logs?region=us-west-2",
	"last_complete": "%[2]s",
	"last_error": "%[2]s",
	"error_message": "test"
  }
`
)

var testLogpushTimestamp = time.Now().UTC()
var expectedLogpushJobStruct = LogpushJob{
	ID:              jobID,
	Enabled:         false,
	Name:            "example.com",
	LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339",
	DestinationConf: "s3://mybucket/logs?region=us-west-2",
	LastComplete:    &testLogpushTimestamp,
	LastError:       &testLogpushTimestamp,
	ErrorMessage:    "test",
}

func TestListLogpushJobs(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
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

	actual, err := client.ListLogpushJobs(testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetLogpushJob(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
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

	actual, err := client.LogpushJob(testZoneID, jobID)
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
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
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

	actual, err := client.CreateLogpushJob(testZoneID, newJob)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteLogpushJob(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Expected method 'DELETE', got %s", r.Method)
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

	err := client.DeleteLogpushJob(testZoneID, jobID)
	assert.NoError(t, err)
}
