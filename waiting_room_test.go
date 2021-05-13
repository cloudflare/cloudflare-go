package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

var testTimestampWaitingRoom = time.Now().UTC()
var waitingRoomJSON = fmt.Sprintf(`
		{
      "id": "699d98642c564d2e855e9661899b7252",
      "created_on": "%s",
      "modified_on": "%s",
      "name": "production_webinar",
      "description": "Production - DO NOT MODIFY",
      "suspended": false,
      "host": "shop.example.com",
      "path": "/shop/checkout",
      "queue_all": true,
      "new_users_per_minute": 600,
      "total_active_users": 1000,
      "session_duration": 10,
      "disable_session_renewal": false,
      "custom_page_html": "{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"
    }
   `, testTimestampWaitingRoom.Format(time.RFC3339Nano), testTimestampWaitingRoom.Format(time.RFC3339Nano))

var waitingRoom = WaitingRoom{
	ID:                    "699d98642c564d2e855e9661899b7252",
	CreatedOn:             testTimestampWaitingRoom,
	ModifiedOn:            testTimestampWaitingRoom,
	Name:                  "production_webinar",
	Description:           "Production - DO NOT MODIFY",
	Suspended:             false,
	Host:                  "shop.example.com",
	Path:                  "/shop/checkout",
	QueueAll:              true,
	NewUsersPerMinute:     600,
	TotalActiveUsers:      1000,
	SessionDuration:       10,
	DisableSessionRenewal: false,
	CustomPageHTML:        "{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}",
}

func TestListWaitingRooms(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": [
			    %s
			  ]
			}
		`, waitingRoomJSON)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms", handler)
	want := []WaitingRoom{waitingRoom}

	actual, err := client.ListWaitingRooms(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListWaitingRoomsNoResult(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": []
			}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms", handler)
	want := []WaitingRoom{}

	actual, err := client.ListWaitingRooms(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestWaitingRoom(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": %s
			}
		`, waitingRoomJSON)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms/699d98642c564d2e855e9661899b7252", handler)
	want := waitingRoom

	actual, err := client.WaitingRoom(context.Background(), testZoneID, "699d98642c564d2e855e9661899b7252")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestWaitingRoomNotFound(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			  "success": null,
			  "errors": [
			  	{
      			"code": 1001,
      			"message": "Object not found."
    			}
    		],
			  "messages": []
			}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms/699d98642c564d2e855e9661899b7252", handler)

	_, err := client.WaitingRoom(context.Background(), testZoneID, "699d98642c564d2e855e9661899b7252")
	assert.NotNil(t, err)
}

func TestCreateWaitingRoom(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": %s
			}
		`, waitingRoomJSON)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms", handler)
	want := &waitingRoom

	actual, err := client.CreateWaitingRoom(context.Background(), testZoneID, waitingRoom)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateWaitingRoomError(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			 "success": false,
			 "errors": [
			  	{
      			"code": 1002,
      			"message": "new_users_per_minute must be in range [200, total_active_users]: invalid data"
    			}
    		],
			  "messages": []
			}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms", handler)

	_, err := client.CreateWaitingRoom(context.Background(), testZoneID, waitingRoom)
	assert.NotNil(t, err)
}

func TestUpdateWaitingRoom(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": %s
			}
		`, waitingRoomJSON)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms/699d98642c564d2e855e9661899b7252", handler)
	want := waitingRoom

	actual, err := client.UpdateWaitingRoom(context.Background(), testZoneID, waitingRoom)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestChangeWaitingRoom(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": %s
			}
		`, waitingRoomJSON)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms/699d98642c564d2e855e9661899b7252", handler)
	want := waitingRoom

	actual, err := client.ChangeWaitingRoom(context.Background(), testZoneID, "699d98642c564d2e855e9661899b7252", WaitingRoom{TotalActiveUsers: 400})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteWaitingRoom(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": {
			    "id": "699d98642c564d2e855e9661899b7252"
			  }
			}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/waiting_rooms/699d98642c564d2e855e9661899b7252", handler)

	err := client.DeleteWaitingRoom(context.Background(), testZoneID, "699d98642c564d2e855e9661899b7252")
	assert.NoError(t, err)
}
