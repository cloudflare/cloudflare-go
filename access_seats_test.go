package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testAccessGroupSeatUID = "access-group-seat-uid"

func TestUpdateAccessUserSeat_ZoneIsNotSupported(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessUserSeat(context.Background(), testZoneRC, UpdateAccessUserSeatParams{})
	assert.EqualError(t, err, fmt.Sprintf(errInvalidResourceContainerAccess, ZoneRouteLevel))
}

func TestUpdateAccessUserSeat_MissingUID(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.UpdateAccessUserSeat(context.Background(), testAccountRC, UpdateAccessUserSeatParams{})
	assert.EqualError(t, err, "missing required access seat UID")
}

func TestUpdateAccessUserSeat(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"errors": [],
			"messages": [],
			"result": [
			  {
				"access_seat": false,
				"created_at": "2014-01-01T05:20:00.12345Z",
				"gateway_seat": false,
				"seat_uid": null,
				"updated_at": "2014-01-01T05:20:00.12345Z"
			  }
			],
			"success": true,
			"result_info": {
			  "count": 1,
			  "page": 1,
			  "per_page": 20,
			  "total_count": 2000
			}
		  }
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/access/seats", handler)
	createdAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	want := []AccessUpdateAccessUserSeatResult{
		{
			AccessSeat:  BoolPtr(false),
			CreatedAt:   &createdAt,
			GatewaySeat: BoolPtr(false),
			SeatUID:     "",
			UpdatedAt:   &updatedAt,
		},
	}

	actual, err := client.UpdateAccessUserSeat(context.Background(), testAccountRC, UpdateAccessUserSeatParams{
		SeatUID:     testAccessGroupSeatUID,
		AccessSeat:  BoolPtr(false),
		GatewaySeat: BoolPtr(false),
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
