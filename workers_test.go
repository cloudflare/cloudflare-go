package cloudflare

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	deleteWorkerResponseData = `{
    "result": null,
    "success": true,
    "errors": [],
    "messages": []
}`
	uploadWorkerResponseData = `{
    "result": {
        "script": "addEventListener('fetch', event => {\n    event.passThroughOnException()\nevent.respondWith(handleRequest(event.request))\n})\n\nasync function handleRequest(request) {\n    return fetch(request)\n}",
        "etag": "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
        "size": 191,
        "modified_on": "2018-06-09T15:17:01.989141Z"
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	updateWorkerRouteResponse = `{
    "result": {
        "id": "e7a57d8746e74ae49c25994dadb421b1",
        "pattern": "app3.example.com/*",
        "enabled": true
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	createWorkerRouteResponse = `{
    "result": {
        "id": "e7a57d8746e74ae49c25994dadb421b1"
    },
    "success": true,
    "errors": [],
    "messages": []
}`
	listRouteResponseData = `{
    "result": [
        {
            "id": "e7a57d8746e74ae49c25994dadb421b1",
            "pattern": "app1.example.com/*",
            "enabled": true
        },
        {
            "id": "f8b68e9857f85bf59c25994dadb421b1",
            "pattern": "app2.example.com/*",
            "enabled": false
        }
    ],
    "success": true,
    "errors": [],
    "messages": []
}`
	listWorkersResponseData = `{
  "result": [
    {
      "id": "bar",
      "created_on": "2018-04-22T17:10:48.938097Z",
      "modified_on": "2018-04-22T17:10:48.938097Z",
      "etag": "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a"
    },
    {
      "id": "baz",
      "created_on": "2018-04-22T17:10:48.938097Z",
      "modified_on": "2018-04-22T17:10:48.938097Z",
      "etag": "380dg51e97e80b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43088b"
    }
  ],
  "success": true,
  "errors": [],
  "messages": []
}`
)

var (
	successResponse               = Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}}
	workerScript                  = "addEventListener('fetch', event => {\n    event.passThroughOnException()\nevent.respondWith(handleRequest(event.request))\n})\n\nasync function handleRequest(request) {\n    return fetch(request)\n}"
	deleteWorkerRouteResponseData = createWorkerRouteResponse
)

func TestWorkers_DeleteWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, deleteWorkerResponseData)
	})
	res, err := client.DeleteWorker("foo")
	want := WorkerScriptResponse{
		successResponse,
		WorkerScript{}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}

func TestWorkers_DeleteWorkerWithName(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, deleteWorkerResponseData)
	})
	res, err := client.DeleteWorkerWithName("bar")
	want := WorkerScriptResponse{
		successResponse,
		WorkerScript{}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Response, res.Response)
	}
}

func TestWorkers_DownloadWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, workerScript)
	})
	res, err := client.DownloadWorker("foo")
	want := WorkerScriptResponse{
		successResponse,
		WorkerScript{
			Script: workerScript,
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Script, res.Script)
	}
}

func TestWorkers_DownloadWorkerWithName(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, workerScript)
	})
	res, err := client.DownloadWorkerWithName("bar")
	want := WorkerScriptResponse{
		successResponse,
		WorkerScript{
			Script: workerScript,
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want.Script, res.Script)
	}
}

func TestWorkers_ListWorkerScripts(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application-json")
		fmt.Fprintf(w, listWorkersResponseData)
	})

	res, err := client.ListWorkerScripts()
	sampleDate, _ := time.Parse(time.RFC3339Nano, "2018-04-22T17:10:48.938097Z")
	want := []WorkerMetaData{
		{
			ID:         "bar",
			ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
			CreatedOn:  sampleDate,
			ModifiedOn: sampleDate,
		},
		{
			ID:         "baz",
			ETAG:       "380dg51e97e80b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43088b",
			CreatedOn:  sampleDate,
			ModifiedOn: sampleDate,
		},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res.WorkerList)
	}

}

func TestWorkers_UploadWorker(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/script", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, uploadWorkerResponseData)
	})
	res, err := client.UploadWorker("foo", workerScript)
	formattedTime, _ := time.Parse(time.RFC3339Nano, "2018-06-09T15:17:01.989141Z")
	want := WorkerScriptResponse{
		successResponse,
		WorkerScript{
			Script: workerScript,
			WorkerMetaData: WorkerMetaData{
				ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
				Size:       191,
				ModifiedOn: formattedTime,
			},
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}

}

func TestWorkers_UploadWorkerWithName(t *testing.T) {
	setup(UsingOrganization("foo"))
	defer teardown()

	mux.HandleFunc("/accounts/foo/workers/scripts/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, uploadWorkerResponseData)
	})
	res, err := client.UploadWorkerWithName("bar", workerScript)
	formattedTime, _ := time.Parse(time.RFC3339Nano, "2018-06-09T15:17:01.989141Z")
	want := WorkerScriptResponse{
		successResponse,
		WorkerScript{
			Script: workerScript,
			WorkerMetaData: WorkerMetaData{
				ETAG:       "279cf40d86d70b82f6cd3ba90a646b3ad995912da446836d7371c21c6a43977a",
				Size:       191,
				ModifiedOn: formattedTime,
			},
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_CreateWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application-json")
		fmt.Fprintf(w, createWorkerRouteResponse)
	})
	route := WorkerRoute{Pattern: "app1.example.com/*", Enabled: true}
	res, err := client.CreateWorkerRoute("foo", route)
	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}

}

func TestWorkers_DeleteWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application-json")
		fmt.Fprintf(w, deleteWorkerRouteResponseData)
	})
	res, err := client.DeleteWorkerRoute("foo", "e7a57d8746e74ae49c25994dadb421b1")
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID: "e7a57d8746e74ae49c25994dadb421b1",
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}

}

func TestWorkers_ListWorkerRoutes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application-json")
		fmt.Fprintf(w, listRouteResponseData)
	})

	res, err := client.ListWorkerRoutes("foo")
	want := WorkerRoutesResponse{successResponse,
		[]WorkerRoute{
			{ID: "e7a57d8746e74ae49c25994dadb421b1", Pattern: "app1.example.com/*", Enabled: true},
			{ID: "f8b68e9857f85bf59c25994dadb421b1", Pattern: "app2.example.com/*", Enabled: false},
		},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkers_UpdateWorkerRoute(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/workers/filters/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application-json")
		fmt.Fprintf(w, updateWorkerRouteResponse)
	})
	route := WorkerRoute{Pattern: "app3.example.com/*", Enabled: true}
	res, err := client.UpdateWorkerRoute("foo", "e7a57d8746e74ae49c25994dadb421b1", route)
	want := WorkerRouteResponse{successResponse,
		WorkerRoute{
			ID:      "e7a57d8746e74ae49c25994dadb421b1",
			Pattern: "app3.example.com/*",
			Enabled: true,
		}}
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}

}
