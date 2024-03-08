// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// WaitingRoomStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomStatusService] method
// instead.
type WaitingRoomStatusService struct {
	Options []option.RequestOption
}

// NewWaitingRoomStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaitingRoomStatusService(opts ...option.RequestOption) (r *WaitingRoomStatusService) {
	r = &WaitingRoomStatusService{}
	r.Options = opts
	return
}

// Fetches the status of a configured waiting room. Response fields include:
//
//  1. `status`: String indicating the status of the waiting room. The possible
//     status are:
//     - **not_queueing** indicates that the configured thresholds have not been met
//     and all users are going through to the origin.
//     - **queueing** indicates that the thresholds have been met and some users are
//     held in the waiting room.
//     - **event_prequeueing** indicates that an event is active and is currently
//     prequeueing users before it starts.
//  2. `event_id`: String of the current event's `id` if an event is active,
//     otherwise an empty string.
//  3. `estimated_queued_users`: Integer of the estimated number of users currently
//     waiting in the queue.
//  4. `estimated_total_active_users`: Integer of the estimated number of users
//     currently active on the origin.
//  5. `max_estimated_time_minutes`: Integer of the maximum estimated time currently
//     presented to the users.
func (r *WaitingRoomStatusService) Get(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *WaitingRoomStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomStatusGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/status", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingRoomStatusGetResponse struct {
	EstimatedQueuedUsers      int64                              `json:"estimated_queued_users"`
	EstimatedTotalActiveUsers int64                              `json:"estimated_total_active_users"`
	EventID                   string                             `json:"event_id"`
	MaxEstimatedTimeMinutes   int64                              `json:"max_estimated_time_minutes"`
	Status                    WaitingRoomStatusGetResponseStatus `json:"status"`
	JSON                      waitingRoomStatusGetResponseJSON   `json:"-"`
}

// waitingRoomStatusGetResponseJSON contains the JSON metadata for the struct
// [WaitingRoomStatusGetResponse]
type waitingRoomStatusGetResponseJSON struct {
	EstimatedQueuedUsers      apijson.Field
	EstimatedTotalActiveUsers apijson.Field
	EventID                   apijson.Field
	MaxEstimatedTimeMinutes   apijson.Field
	Status                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WaitingRoomStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomStatusGetResponseStatus string

const (
	WaitingRoomStatusGetResponseStatusEventPrequeueing WaitingRoomStatusGetResponseStatus = "event_prequeueing"
	WaitingRoomStatusGetResponseStatusNotQueueing      WaitingRoomStatusGetResponseStatus = "not_queueing"
	WaitingRoomStatusGetResponseStatusQueueing         WaitingRoomStatusGetResponseStatus = "queueing"
)

type WaitingRoomStatusGetResponseEnvelope struct {
	Result WaitingRoomStatusGetResponse             `json:"result,required"`
	JSON   waitingRoomStatusGetResponseEnvelopeJSON `json:"-"`
}

// waitingRoomStatusGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomStatusGetResponseEnvelope]
type waitingRoomStatusGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
