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

// WaitingRoomEventService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomEventService] method
// instead.
type WaitingRoomEventService struct {
	Options []option.RequestOption
	Details *WaitingRoomEventDetailService
}

// NewWaitingRoomEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaitingRoomEventService(opts ...option.RequestOption) (r *WaitingRoomEventService) {
	r = &WaitingRoomEventService{}
	r.Options = opts
	r.Details = NewWaitingRoomEventDetailService(opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Creates an event for
// a waiting room. An event takes place during a specified period of time,
// temporarily changing the behavior of a waiting room. While the event is active,
// some of the properties in the event's configuration may either override or
// inherit from the waiting room's configuration. Note that events cannot overlap
// with each other, so only one event can be active at a time.
func (r *WaitingRoomEventService) New(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomEventNewParams, opts ...option.RequestOption) (res *WaitingRoomEventNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured event for a waiting room.
func (r *WaitingRoomEventService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, body WaitingRoomEventUpdateParams, opts ...option.RequestOption) (res *WaitingRoomEventUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists events for a waiting room.
func (r *WaitingRoomEventService) List(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *[]WaitingRoomEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventListResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an event for a waiting room.
func (r *WaitingRoomEventService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, opts ...option.RequestOption) (res *WaitingRoomEventDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches a configured event for a waiting room.
func (r *WaitingRoomEventService) Edit(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, body WaitingRoomEventEditParams, opts ...option.RequestOption) (res *WaitingRoomEventEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single configured event for a waiting room.
func (r *WaitingRoomEventService) Get(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, eventID interface{}, opts ...option.RequestOption) (res *WaitingRoomEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEventGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/events/%v", zoneIdentifier, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingRoomEventNewResponse struct {
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
	TotalActiveUsers int64                           `json:"total_active_users,nullable"`
	JSON             waitingRoomEventNewResponseJSON `json:"-"`
}

// waitingRoomEventNewResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventNewResponse]
type waitingRoomEventNewResponseJSON struct {
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

func (r *WaitingRoomEventNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventNewResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventUpdateResponse struct {
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
	TotalActiveUsers int64                              `json:"total_active_users,nullable"`
	JSON             waitingRoomEventUpdateResponseJSON `json:"-"`
}

// waitingRoomEventUpdateResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventUpdateResponse]
type waitingRoomEventUpdateResponseJSON struct {
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

func (r *WaitingRoomEventUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventListResponse struct {
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
	TotalActiveUsers int64                            `json:"total_active_users,nullable"`
	JSON             waitingRoomEventListResponseJSON `json:"-"`
}

// waitingRoomEventListResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventListResponse]
type waitingRoomEventListResponseJSON struct {
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

func (r *WaitingRoomEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventListResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventDeleteResponse struct {
	ID   interface{}                        `json:"id"`
	JSON waitingRoomEventDeleteResponseJSON `json:"-"`
}

// waitingRoomEventDeleteResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventDeleteResponse]
type waitingRoomEventDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventEditResponse struct {
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
	TotalActiveUsers int64                            `json:"total_active_users,nullable"`
	JSON             waitingRoomEventEditResponseJSON `json:"-"`
}

// waitingRoomEventEditResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventEditResponse]
type waitingRoomEventEditResponseJSON struct {
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

func (r *WaitingRoomEventEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventEditResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventGetResponse struct {
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
	TotalActiveUsers int64                           `json:"total_active_users,nullable"`
	JSON             waitingRoomEventGetResponseJSON `json:"-"`
}

// waitingRoomEventGetResponseJSON contains the JSON metadata for the struct
// [WaitingRoomEventGetResponse]
type waitingRoomEventGetResponseJSON struct {
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

func (r *WaitingRoomEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventNewParams struct {
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

func (r WaitingRoomEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaitingRoomEventNewResponseEnvelope struct {
	Result WaitingRoomEventNewResponse             `json:"result,required"`
	JSON   waitingRoomEventNewResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEventNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventNewResponseEnvelope]
type waitingRoomEventNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventUpdateParams struct {
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

func (r WaitingRoomEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaitingRoomEventUpdateResponseEnvelope struct {
	Result WaitingRoomEventUpdateResponse             `json:"result,required"`
	JSON   waitingRoomEventUpdateResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEventUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventUpdateResponseEnvelope]
type waitingRoomEventUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventListResponseEnvelope struct {
	Errors   []WaitingRoomEventListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomEventListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomEventListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomEventListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomEventListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomEventListResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomEventListResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventListResponseEnvelope]
type waitingRoomEventListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    waitingRoomEventListResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomEventListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WaitingRoomEventListResponseEnvelopeErrors]
type waitingRoomEventListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    waitingRoomEventListResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomEventListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WaitingRoomEventListResponseEnvelopeMessages]
type waitingRoomEventListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WaitingRoomEventListResponseEnvelopeSuccess bool

const (
	WaitingRoomEventListResponseEnvelopeSuccessTrue WaitingRoomEventListResponseEnvelopeSuccess = true
)

type WaitingRoomEventListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       waitingRoomEventListResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomEventListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [WaitingRoomEventListResponseEnvelopeResultInfo]
type waitingRoomEventListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventDeleteResponseEnvelope struct {
	Result WaitingRoomEventDeleteResponse             `json:"result,required"`
	JSON   waitingRoomEventDeleteResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEventDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventDeleteResponseEnvelope]
type waitingRoomEventDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventEditParams struct {
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

func (r WaitingRoomEventEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaitingRoomEventEditResponseEnvelope struct {
	Result WaitingRoomEventEditResponse             `json:"result,required"`
	JSON   waitingRoomEventEditResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEventEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventEditResponseEnvelope]
type waitingRoomEventEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEventGetResponseEnvelope struct {
	Result WaitingRoomEventGetResponse             `json:"result,required"`
	JSON   waitingRoomEventGetResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEventGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomEventGetResponseEnvelope]
type waitingRoomEventGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEventGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEventGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
