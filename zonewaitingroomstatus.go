// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWaitingRoomStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWaitingRoomStatusService]
// method instead.
type ZoneWaitingRoomStatusService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomStatusService(opts ...option.RequestOption) (r *ZoneWaitingRoomStatusService) {
	r = &ZoneWaitingRoomStatusService{}
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
func (r *ZoneWaitingRoomStatusService) WaitingRoomGetWaitingRoomStatus(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *StatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/status", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StatusResponse struct {
	Errors   []StatusResponseError   `json:"errors"`
	Messages []StatusResponseMessage `json:"messages"`
	Result   StatusResponseResult    `json:"result"`
	// Whether the API call was successful
	Success StatusResponseSuccess `json:"success"`
	JSON    statusResponseJSON    `json:"-"`
}

// statusResponseJSON contains the JSON metadata for the struct [StatusResponse]
type statusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StatusResponseError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    statusResponseErrorJSON `json:"-"`
}

// statusResponseErrorJSON contains the JSON metadata for the struct
// [StatusResponseError]
type statusResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StatusResponseMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    statusResponseMessageJSON `json:"-"`
}

// statusResponseMessageJSON contains the JSON metadata for the struct
// [StatusResponseMessage]
type statusResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StatusResponseResult struct {
	EstimatedQueuedUsers      int64                      `json:"estimated_queued_users"`
	EstimatedTotalActiveUsers int64                      `json:"estimated_total_active_users"`
	EventID                   string                     `json:"event_id"`
	MaxEstimatedTimeMinutes   int64                      `json:"max_estimated_time_minutes"`
	Status                    StatusResponseResultStatus `json:"status"`
	JSON                      statusResponseResultJSON   `json:"-"`
}

// statusResponseResultJSON contains the JSON metadata for the struct
// [StatusResponseResult]
type statusResponseResultJSON struct {
	EstimatedQueuedUsers      apijson.Field
	EstimatedTotalActiveUsers apijson.Field
	EventID                   apijson.Field
	MaxEstimatedTimeMinutes   apijson.Field
	Status                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *StatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StatusResponseResultStatus string

const (
	StatusResponseResultStatusEventPrequeueing StatusResponseResultStatus = "event_prequeueing"
	StatusResponseResultStatusNotQueueing      StatusResponseResultStatus = "not_queueing"
	StatusResponseResultStatusQueueing         StatusResponseResultStatus = "queueing"
)

// Whether the API call was successful
type StatusResponseSuccess bool

const (
	StatusResponseSuccessTrue StatusResponseSuccess = true
)
