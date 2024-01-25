// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWaitingRoomEventService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWaitingRoomEventService]
// method instead.
type ZoneWaitingRoomEventService struct {
	Options []option.RequestOption
	Details *ZoneWaitingRoomEventDetailService
}

// NewZoneWaitingRoomEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomEventService(opts ...option.RequestOption) (r *ZoneWaitingRoomEventService) {
	r = &ZoneWaitingRoomEventService{}
	r.Options = opts
	r.Details = NewZoneWaitingRoomEventDetailService(opts...)
	return
}

// Fetches a single configured event for a waiting room.
func (r *ZoneWaitingRoomEventService) Get(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured event for a waiting room.
func (r *ZoneWaitingRoomEventService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, body ZoneWaitingRoomEventUpdateParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an event for a waiting room.
func (r *ZoneWaitingRoomEventService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, opts ...option.RequestOption) (res *ZoneWaitingRoomEventDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patches a configured event for a waiting room.
func (r *ZoneWaitingRoomEventService) Patch(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, body ZoneWaitingRoomEventPatchParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Creates an event for
// a waiting room. An event takes place during a specified period of time,
// temporarily changing the behavior of a waiting room. While the event is active,
// some of the properties in the event's configuration may either override or
// inherit from the waiting room's configuration. Note that events cannot overlap
// with each other, so only one event can be active at a time.
func (r *ZoneWaitingRoomEventService) WaitingRoomNewEvent(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomEventWaitingRoomNewEventParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists events for a waiting room.
func (r *ZoneWaitingRoomEventService) WaitingRoomListEvents(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *ZoneWaitingRoomEventWaitingRoomListEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EventResponse struct {
	Result EventResponseResult `json:"result"`
	JSON   eventResponseJSON   `json:"-"`
}

// eventResponseJSON contains the JSON metadata for the struct [EventResponse]
type eventResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedOn time.Time   `json:"created_on" format:"date-time"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML string `json:"custom_page_html,nullable"`
	// A note that you can use to add more details about the event.
	Description string `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal bool `json:"disable_session_renewal,nullable"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name string `json:"name"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute int64 `json:"new_users_per_minute,nullable"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod string `json:"queueing_method,nullable"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration int64 `json:"session_duration,nullable"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended bool `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers int64                   `json:"total_active_users,nullable"`
	JSON             eventResponseResultJSON `json:"-"`
}

// eventResponseResultJSON contains the JSON metadata for the struct
// [EventResponseResult]
type eventResponseResultJSON struct {
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

func (r *EventResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventDeleteResponse struct {
	Result ZoneWaitingRoomEventDeleteResponseResult `json:"result"`
	JSON   zoneWaitingRoomEventDeleteResponseJSON   `json:"-"`
}

// zoneWaitingRoomEventDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomEventDeleteResponse]
type zoneWaitingRoomEventDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventDeleteResponseResult struct {
	ID   interface{}                                  `json:"id"`
	JSON zoneWaitingRoomEventDeleteResponseResultJSON `json:"-"`
}

// zoneWaitingRoomEventDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomEventDeleteResponseResult]
type zoneWaitingRoomEventDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventWaitingRoomListEventsResponse struct {
	Errors     []ZoneWaitingRoomEventWaitingRoomListEventsResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomEventWaitingRoomListEventsResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomEventWaitingRoomListEventsResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomEventWaitingRoomListEventsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomEventWaitingRoomListEventsResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomEventWaitingRoomListEventsResponseJSON    `json:"-"`
}

// zoneWaitingRoomEventWaitingRoomListEventsResponseJSON contains the JSON metadata
// for the struct [ZoneWaitingRoomEventWaitingRoomListEventsResponse]
type zoneWaitingRoomEventWaitingRoomListEventsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventWaitingRoomListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventWaitingRoomListEventsResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneWaitingRoomEventWaitingRoomListEventsResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomEventWaitingRoomListEventsResponseErrorJSON contains the JSON
// metadata for the struct [ZoneWaitingRoomEventWaitingRoomListEventsResponseError]
type zoneWaitingRoomEventWaitingRoomListEventsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventWaitingRoomListEventsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventWaitingRoomListEventsResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneWaitingRoomEventWaitingRoomListEventsResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomEventWaitingRoomListEventsResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomEventWaitingRoomListEventsResponseMessage]
type zoneWaitingRoomEventWaitingRoomListEventsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventWaitingRoomListEventsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventWaitingRoomListEventsResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedOn time.Time   `json:"created_on" format:"date-time"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML string `json:"custom_page_html,nullable"`
	// A note that you can use to add more details about the event.
	Description string `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal bool `json:"disable_session_renewal,nullable"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name string `json:"name"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute int64 `json:"new_users_per_minute,nullable"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod string `json:"queueing_method,nullable"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration int64 `json:"session_duration,nullable"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended bool `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers int64                                                       `json:"total_active_users,nullable"`
	JSON             zoneWaitingRoomEventWaitingRoomListEventsResponseResultJSON `json:"-"`
}

// zoneWaitingRoomEventWaitingRoomListEventsResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomEventWaitingRoomListEventsResponseResult]
type zoneWaitingRoomEventWaitingRoomListEventsResponseResultJSON struct {
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

func (r *ZoneWaitingRoomEventWaitingRoomListEventsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomEventWaitingRoomListEventsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       zoneWaitingRoomEventWaitingRoomListEventsResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomEventWaitingRoomListEventsResponseResultInfoJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomEventWaitingRoomListEventsResponseResultInfo]
type zoneWaitingRoomEventWaitingRoomListEventsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventWaitingRoomListEventsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomEventWaitingRoomListEventsResponseSuccess bool

const (
	ZoneWaitingRoomEventWaitingRoomListEventsResponseSuccessTrue ZoneWaitingRoomEventWaitingRoomListEventsResponseSuccess = true
)

type ZoneWaitingRoomEventUpdateParams struct {
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime param.Field[string] `json:"event_end_time,required"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime param.Field[string] `json:"event_start_time,required"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description param.Field[string] `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime param.Field[string] `json:"prequeue_start_time"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod param.Field[string] `json:"queueing_method"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart param.Field[bool] `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended param.Field[bool] `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers param.Field[int64] `json:"total_active_users"`
}

func (r ZoneWaitingRoomEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWaitingRoomEventPatchParams struct {
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime param.Field[string] `json:"event_end_time,required"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime param.Field[string] `json:"event_start_time,required"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description param.Field[string] `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime param.Field[string] `json:"prequeue_start_time"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod param.Field[string] `json:"queueing_method"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart param.Field[bool] `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended param.Field[bool] `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers param.Field[int64] `json:"total_active_users"`
}

func (r ZoneWaitingRoomEventPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWaitingRoomEventWaitingRoomNewEventParams struct {
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime param.Field[string] `json:"event_end_time,required"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime param.Field[string] `json:"event_start_time,required"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description param.Field[string] `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime param.Field[string] `json:"prequeue_start_time"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod param.Field[string] `json:"queueing_method"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart param.Field[bool] `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended param.Field[bool] `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers param.Field[int64] `json:"total_active_users"`
}

func (r ZoneWaitingRoomEventWaitingRoomNewEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
