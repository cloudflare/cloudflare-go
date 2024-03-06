// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WaitingRoomEventDetailService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomEventDetailService]
// method instead.
type WaitingRoomEventDetailService struct {
	Options []option.RequestOption
}

// NewWaitingRoomEventDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaitingRoomEventDetailService(opts ...option.RequestOption) (r *WaitingRoomEventDetailService) {
	r = &WaitingRoomEventDetailService{}
	r.Options = opts
	return
}

// Previews an event's configuration as if it was active. Inherited fields from the
// waiting room will be displayed with their current values.
func (r *WaitingRoomEventDetailService) Get(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, opts ...option.RequestOption) (res *WaitingRoomEventDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventDetailGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v/details", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingRoomEventDetailGetResponse struct {
	ID             interface{} `json:"id"`
	CreatedOn      time.Time   `json:"created_on" format:"date-time"`
	CustomPageHTML string      `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description           string `json:"description"`
	DisableSessionRenewal bool   `json:"disable_session_renewal"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name              string `json:"name"`
	NewUsersPerMinute int64  `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	QueueingMethod    string `json:"queueing_method"`
	SessionDuration   int64  `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended        bool                                  `json:"suspended"`
	TotalActiveUsers int64                                 `json:"total_active_users"`
	JSON             waitingRoomEventDetailGetResponseJSON `json:"-"`
}

// waitingRoomEventDetailGetResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventDetailGetResponse]
type waitingRoomEventDetailGetResponseJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	CustomPageHTML        apijson.Field
	Description           apijson.Field
	DisableSessionRenewal apijson.Field
	EventEndTime          apijson.Field
	EventStartTime        apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NewUsersPerMinute     apijson.Field
	PrequeueStartTime     apijson.Field
	QueueingMethod        apijson.Field
	SessionDuration       apijson.Field
	ShuffleAtEventStart   apijson.Field
	Suspended             apijson.Field
	TotalActiveUsers      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WaitingRoomEventDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventDetailGetResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventDetailGetResponseEnvelope struct {
	Result WaitingRoomEventDetailGetResponse             `json:"result,required"`
	JSON   waitingRoomEventDetailGetResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEventDetailGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventDetailGetResponseEnvelope]
type waitingRoomEventDetailGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventDetailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventDetailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
