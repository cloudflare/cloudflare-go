// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// LivestreamService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLivestreamService] method instead.
type LivestreamService struct {
	Options []option.RequestOption
}

// NewLivestreamService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLivestreamService(opts ...option.RequestOption) (r *LivestreamService) {
	r = &LivestreamService{}
	r.Options = opts
	return
}

// Returns details of all active livestreams for the given livestream ID. Retreive
// the livestream ID using the `Start livestreaming a meeting` API.
func (r *LivestreamService) GetActiveLivestreamsForLivestreamID(ctx context.Context, appID string, livestreamID string, query LivestreamGetActiveLivestreamsForLivestreamIDParams, opts ...option.RequestOption) (res *LivestreamGetActiveLivestreamsForLivestreamIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if livestreamID == "" {
		err = errors.New("missing required livestream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams/%s/active-livestream-session", query.AccountID, appID, livestreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns details of livestreams associated with the given App ID. It includes
// livestreams created by your App and RealtimeKit meetings that are livestreamed
// by your App. If you only want details of livestreams created by your App and not
// RealtimeKit meetings, you can use the `exclude_meetings` query parameter.
func (r *LivestreamService) GetAllLivestreams(ctx context.Context, appID string, params LivestreamGetAllLivestreamsParams, opts ...option.RequestOption) (res *LivestreamGetAllLivestreamsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns livestream analytics for the specified time range.
func (r *LivestreamService) GetLivestreamAnalyticsComplete(ctx context.Context, appID string, params LivestreamGetLivestreamAnalyticsCompleteParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamAnalyticsCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/analytics/livestreams/overall", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns day-wise livestream analytics for the specified time range.
func (r *LivestreamService) GetLivestreamAnalyticsDaywise(ctx context.Context, appID string, params LivestreamGetLivestreamAnalyticsDaywiseParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamAnalyticsDaywiseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/analytics/livestreams/daywise", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns livestream session details for the given meeting ID. Retreive the
// meeting ID using the `Create a meeting` API.
func (r *LivestreamService) GetLivestreamSessionDetailsForMeeting(ctx context.Context, appID string, meetingID string, params LivestreamGetLivestreamSessionDetailsForMeetingParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamSessionDetailsForMeetingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/livestream", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns livestream session details for the given livestream session ID. Retrieve
// the `livestream_session_id`using the
// `Fetch livestream session details using a session ID` API.
func (r *LivestreamService) GetLivestreamSessionDetailsForSessionID(ctx context.Context, appID string, livestreamSessionID string, query LivestreamGetLivestreamSessionDetailsForSessionIDParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamSessionDetailsForSessionIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if livestreamSessionID == "" {
		err = errors.New("missing required livestream-session-id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams/sessions/%s", query.AccountID, appID, livestreamSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns details of a livestream with sessions for the given livestream ID.
// Retreive the livestream ID using the `Start livestreaming a meeting` API.
func (r *LivestreamService) GetLivestreamSessionForLivestreamID(ctx context.Context, appID string, livestreamID string, params LivestreamGetLivestreamSessionForLivestreamIDParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamSessionForLivestreamIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if livestreamID == "" {
		err = errors.New("missing required livestream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams/%s", params.AccountID, appID, livestreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns livestream session details for the given session ID. Retreive the
// session ID using the `Fetch all sessions of an App` API.
func (r *LivestreamService) GetLivestreamSessionsForSession(ctx context.Context, appID string, sessionID string, params LivestreamGetLivestreamSessionsForSessionParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamSessionsForSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/livestream-sessions", params.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns details of all active livestreams for the given meeting ID.
func (r *LivestreamService) GetMeetingActiveLivestreams(ctx context.Context, appID string, meetingID string, query LivestreamGetMeetingActiveLivestreamsParams, opts ...option.RequestOption) (res *LivestreamGetMeetingActiveLivestreamsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-livestream", query.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns day-wise session and recording analytics data of an App for the
// specified time range start_date to end_date. If start_date and end_date are not
// provided, the default time range is set from 30 days ago to the current date.
func (r *LivestreamService) GetOrgAnalytics(ctx context.Context, appID string, params LivestreamGetOrgAnalyticsParams, opts ...option.RequestOption) (res *LivestreamGetOrgAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/analytics/daywise", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Starts livestream of a meeting associated with the given meeting ID. Retreive
// the meeting ID using the `Create a meeting` API.
func (r *LivestreamService) StartLivestreamingAMeeting(ctx context.Context, appID string, meetingID string, params LivestreamStartLivestreamingAMeetingParams, opts ...option.RequestOption) (res *LivestreamStartLivestreamingAMeetingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/livestreams", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Stops the active livestream of a meeting associated with the given meeting ID.
// Retreive the meeting ID using the `Create a meeting` API.
func (r *LivestreamService) StopLivestreamingAMeeting(ctx context.Context, appID string, meetingID string, body LivestreamStopLivestreamingAMeetingParams, opts ...option.RequestOption) (res *LivestreamStopLivestreamingAMeetingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-livestream/stop", body.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type LivestreamGetActiveLivestreamsForLivestreamIDResponse struct {
	Data    LivestreamGetActiveLivestreamsForLivestreamIDResponseData `json:"data" api:"required"`
	Success bool                                                      `json:"success" api:"required"`
	JSON    livestreamGetActiveLivestreamsForLivestreamIDResponseJSON `json:"-"`
}

// livestreamGetActiveLivestreamsForLivestreamIDResponseJSON contains the JSON
// metadata for the struct [LivestreamGetActiveLivestreamsForLivestreamIDResponse]
type livestreamGetActiveLivestreamsForLivestreamIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetActiveLivestreamsForLivestreamIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetActiveLivestreamsForLivestreamIDResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetActiveLivestreamsForLivestreamIDResponseData struct {
	Livestream LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestream `json:"livestream" api:"required"`
	Session    LivestreamGetActiveLivestreamsForLivestreamIDResponseDataSession    `json:"session" api:"required"`
	JSON       livestreamGetActiveLivestreamsForLivestreamIDResponseDataJSON       `json:"-"`
}

// livestreamGetActiveLivestreamsForLivestreamIDResponseDataJSON contains the JSON
// metadata for the struct
// [LivestreamGetActiveLivestreamsForLivestreamIDResponseData]
type livestreamGetActiveLivestreamsForLivestreamIDResponseDataJSON struct {
	Livestream  apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetActiveLivestreamsForLivestreamIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetActiveLivestreamsForLivestreamIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestream struct {
	// ID of the livestream.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Specifies if the livestream is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// ID of the underlying Cloudflare Stream live input.
	ExternalID string `json:"external_id" api:"required"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server" api:"required"`
	// Specifies if the livestream is ingested manually by an RTMP encoder.
	ManualIngest bool `json:"manual_ingest" api:"required"`
	// ID of the meeting being livestreamed, if any.
	MeetingID string `json:"meeting_id" api:"required,nullable"`
	// Name of the livestream.
	Name string `json:"name" api:"required,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string `json:"playback_url" api:"required"`
	// Status of the livestream.
	Status LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus `json:"status" api:"required"`
	// Unique key for accessing each livestream. Empty if the livestream is disabled.
	StreamKey string `json:"stream_key" api:"required"`
	// Timestamp the livestream was last updated at, in ISO format.
	UpdatedAt time.Time                                                               `json:"updated_at" api:"required" format:"date-time"`
	JSON      livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON `json:"-"`
}

// livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON contains
// the JSON metadata for the struct
// [LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestream]
type livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	ExternalID   apijson.Field
	IngestServer apijson.Field
	ManualIngest apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON) RawJSON() string {
	return r.raw
}

// Status of the livestream.
type LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus string

const (
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusLive         LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "LIVE"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusIdle         LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "IDLE"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusErrored      LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "ERRORED"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusInvoked      LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "INVOKED"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusMetricsReady LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "METRICS_READY"
)

func (r LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus) IsKnown() bool {
	switch r {
	case LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusLive, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusIdle, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusErrored, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusInvoked, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusMetricsReady:
		return true
	}
	return false
}

type LivestreamGetActiveLivestreamsForLivestreamIDResponseDataSession struct {
	// ID of the livestream session.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream session was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error message, if the livestream session failed.
	ErrMessage string `json:"err_message" api:"required,nullable"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds float64 `json:"ingest_seconds" api:"required"`
	// Timestamp the livestream session was invoked at, in ISO format.
	InvokedTime time.Time `json:"invoked_time" api:"required,nullable" format:"date-time"`
	// ID of the livestream.
	LivestreamID string `json:"livestream_id" api:"required" format:"uuid"`
	// Timestamp the livestream session started at, in ISO format.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session stopped at, in ISO format.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session was last updated at, in ISO format.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds float64                                                              `json:"viewer_seconds" api:"required"`
	JSON          livestreamGetActiveLivestreamsForLivestreamIDResponseDataSessionJSON `json:"-"`
}

// livestreamGetActiveLivestreamsForLivestreamIDResponseDataSessionJSON contains
// the JSON metadata for the struct
// [LivestreamGetActiveLivestreamsForLivestreamIDResponseDataSession]
type livestreamGetActiveLivestreamsForLivestreamIDResponseDataSessionJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ErrMessage    apijson.Field
	IngestSeconds apijson.Field
	InvokedTime   apijson.Field
	LivestreamID  apijson.Field
	StartedTime   apijson.Field
	StoppedTime   apijson.Field
	UpdatedAt     apijson.Field
	ViewerSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LivestreamGetActiveLivestreamsForLivestreamIDResponseDataSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetActiveLivestreamsForLivestreamIDResponseDataSessionJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetAllLivestreamsResponse struct {
	Data    []LivestreamGetAllLivestreamsResponseData `json:"data" api:"required"`
	Paging  LivestreamGetAllLivestreamsResponsePaging `json:"paging" api:"required"`
	Success bool                                      `json:"success" api:"required"`
	JSON    livestreamGetAllLivestreamsResponseJSON   `json:"-"`
}

// livestreamGetAllLivestreamsResponseJSON contains the JSON metadata for the
// struct [LivestreamGetAllLivestreamsResponse]
type livestreamGetAllLivestreamsResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetAllLivestreamsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetAllLivestreamsResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetAllLivestreamsResponseData struct {
	// ID of the livestream.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Specifies if the livestream is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// ID of the underlying Cloudflare Stream live input.
	ExternalID string `json:"external_id" api:"required"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server" api:"required"`
	// Specifies if the livestream is ingested manually by an RTMP encoder.
	ManualIngest bool `json:"manual_ingest" api:"required"`
	// ID of the meeting being livestreamed, if any.
	MeetingID string `json:"meeting_id" api:"required,nullable"`
	// Name of the livestream.
	Name string `json:"name" api:"required,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string `json:"playback_url" api:"required"`
	// Status of the livestream.
	Status LivestreamGetAllLivestreamsResponseDataStatus `json:"status" api:"required"`
	// Unique key for accessing each livestream. Empty if the livestream is disabled.
	StreamKey string `json:"stream_key" api:"required"`
	// Timestamp the livestream was last updated at, in ISO format.
	UpdatedAt time.Time                                   `json:"updated_at" api:"required" format:"date-time"`
	JSON      livestreamGetAllLivestreamsResponseDataJSON `json:"-"`
}

// livestreamGetAllLivestreamsResponseDataJSON contains the JSON metadata for the
// struct [LivestreamGetAllLivestreamsResponseData]
type livestreamGetAllLivestreamsResponseDataJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	ExternalID   apijson.Field
	IngestServer apijson.Field
	ManualIngest apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamGetAllLivestreamsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetAllLivestreamsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status of the livestream.
type LivestreamGetAllLivestreamsResponseDataStatus string

const (
	LivestreamGetAllLivestreamsResponseDataStatusLive         LivestreamGetAllLivestreamsResponseDataStatus = "LIVE"
	LivestreamGetAllLivestreamsResponseDataStatusIdle         LivestreamGetAllLivestreamsResponseDataStatus = "IDLE"
	LivestreamGetAllLivestreamsResponseDataStatusErrored      LivestreamGetAllLivestreamsResponseDataStatus = "ERRORED"
	LivestreamGetAllLivestreamsResponseDataStatusInvoked      LivestreamGetAllLivestreamsResponseDataStatus = "INVOKED"
	LivestreamGetAllLivestreamsResponseDataStatusMetricsReady LivestreamGetAllLivestreamsResponseDataStatus = "METRICS_READY"
)

func (r LivestreamGetAllLivestreamsResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamGetAllLivestreamsResponseDataStatusLive, LivestreamGetAllLivestreamsResponseDataStatusIdle, LivestreamGetAllLivestreamsResponseDataStatusErrored, LivestreamGetAllLivestreamsResponseDataStatusInvoked, LivestreamGetAllLivestreamsResponseDataStatusMetricsReady:
		return true
	}
	return false
}

type LivestreamGetAllLivestreamsResponsePaging struct {
	EndOffset   float64                                       `json:"end_offset" api:"required"`
	StartOffset float64                                       `json:"start_offset" api:"required"`
	TotalCount  float64                                       `json:"total_count" api:"required"`
	JSON        livestreamGetAllLivestreamsResponsePagingJSON `json:"-"`
}

// livestreamGetAllLivestreamsResponsePagingJSON contains the JSON metadata for the
// struct [LivestreamGetAllLivestreamsResponsePaging]
type livestreamGetAllLivestreamsResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetAllLivestreamsResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetAllLivestreamsResponsePagingJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamAnalyticsCompleteResponse struct {
	Data    LivestreamGetLivestreamAnalyticsCompleteResponseData `json:"data"`
	Success bool                                                 `json:"success"`
	JSON    livestreamGetLivestreamAnalyticsCompleteResponseJSON `json:"-"`
}

// livestreamGetLivestreamAnalyticsCompleteResponseJSON contains the JSON metadata
// for the struct [LivestreamGetLivestreamAnalyticsCompleteResponse]
type livestreamGetLivestreamAnalyticsCompleteResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamAnalyticsCompleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamAnalyticsCompleteResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamAnalyticsCompleteResponseData struct {
	// Count of total livestreams.
	Count int64 `json:"count"`
	// Total time duration for which the input was given or the meeting was streamed.
	TotalIngestSeconds float64 `json:"total_ingest_seconds"`
	// Total view time for which the viewers watched the stream.
	TotalViewerSeconds float64                                                  `json:"total_viewer_seconds"`
	JSON               livestreamGetLivestreamAnalyticsCompleteResponseDataJSON `json:"-"`
}

// livestreamGetLivestreamAnalyticsCompleteResponseDataJSON contains the JSON
// metadata for the struct [LivestreamGetLivestreamAnalyticsCompleteResponseData]
type livestreamGetLivestreamAnalyticsCompleteResponseDataJSON struct {
	Count              apijson.Field
	TotalIngestSeconds apijson.Field
	TotalViewerSeconds apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LivestreamGetLivestreamAnalyticsCompleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamAnalyticsCompleteResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamAnalyticsDaywiseResponse struct {
	Data    []LivestreamGetLivestreamAnalyticsDaywiseResponseData `json:"data"`
	Success bool                                                  `json:"success"`
	JSON    livestreamGetLivestreamAnalyticsDaywiseResponseJSON   `json:"-"`
}

// livestreamGetLivestreamAnalyticsDaywiseResponseJSON contains the JSON metadata
// for the struct [LivestreamGetLivestreamAnalyticsDaywiseResponse]
type livestreamGetLivestreamAnalyticsDaywiseResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamAnalyticsDaywiseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamAnalyticsDaywiseResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamAnalyticsDaywiseResponseData struct {
	// Count of total livestream sessions.
	Count int64 `json:"count"`
	// Analytics date.
	Date string `json:"date" api:"nullable"`
	// Total time duration for which the input was given or the meeting was streamed.
	TotalIngestSeconds float64 `json:"total_ingest_seconds"`
	// Total view time for which the viewers watched the stream.
	TotalViewerSeconds float64                                                 `json:"total_viewer_seconds"`
	JSON               livestreamGetLivestreamAnalyticsDaywiseResponseDataJSON `json:"-"`
}

// livestreamGetLivestreamAnalyticsDaywiseResponseDataJSON contains the JSON
// metadata for the struct [LivestreamGetLivestreamAnalyticsDaywiseResponseData]
type livestreamGetLivestreamAnalyticsDaywiseResponseDataJSON struct {
	Count              apijson.Field
	Date               apijson.Field
	TotalIngestSeconds apijson.Field
	TotalViewerSeconds apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LivestreamGetLivestreamAnalyticsDaywiseResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamAnalyticsDaywiseResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForMeetingResponse struct {
	Data    LivestreamGetLivestreamSessionDetailsForMeetingResponseData `json:"data" api:"required"`
	Success bool                                                        `json:"success" api:"required"`
	JSON    livestreamGetLivestreamSessionDetailsForMeetingResponseJSON `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForMeetingResponseJSON contains the JSON
// metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForMeetingResponse]
type livestreamGetLivestreamSessionDetailsForMeetingResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForMeetingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForMeetingResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForMeetingResponseData struct {
	Data   LivestreamGetLivestreamSessionDetailsForMeetingResponseDataData   `json:"data" api:"required"`
	Paging LivestreamGetLivestreamSessionDetailsForMeetingResponseDataPaging `json:"paging" api:"required"`
	JSON   livestreamGetLivestreamSessionDetailsForMeetingResponseDataJSON   `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForMeetingResponseDataJSON contains the
// JSON metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForMeetingResponseData]
type livestreamGetLivestreamSessionDetailsForMeetingResponseDataJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForMeetingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForMeetingResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForMeetingResponseDataData struct {
	Livestream LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestream `json:"livestream" api:"required"`
	Sessions   []LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSession  `json:"sessions" api:"required"`
	JSON       livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataJSON       `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataJSON contains the
// JSON metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForMeetingResponseDataData]
type livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataJSON struct {
	Livestream  apijson.Field
	Sessions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForMeetingResponseDataData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestream struct {
	// ID of the livestream.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Specifies if the livestream is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// ID of the underlying Cloudflare Stream live input.
	ExternalID string `json:"external_id" api:"required"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server" api:"required"`
	// Specifies if the livestream is ingested manually by an RTMP encoder.
	ManualIngest bool `json:"manual_ingest" api:"required"`
	// ID of the meeting being livestreamed, if any.
	MeetingID string `json:"meeting_id" api:"required,nullable"`
	// Name of the livestream.
	Name string `json:"name" api:"required,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string `json:"playback_url" api:"required"`
	// Status of the livestream.
	Status LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus `json:"status" api:"required"`
	// Unique key for accessing each livestream. Empty if the livestream is disabled.
	StreamKey string `json:"stream_key" api:"required"`
	// Timestamp the livestream was last updated at, in ISO format.
	UpdatedAt time.Time                                                                     `json:"updated_at" api:"required" format:"date-time"`
	JSON      livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamJSON `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamJSON
// contains the JSON metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestream]
type livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	ExternalID   apijson.Field
	IngestServer apijson.Field
	ManualIngest apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamJSON) RawJSON() string {
	return r.raw
}

// Status of the livestream.
type LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus string

const (
	LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusLive         LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus = "LIVE"
	LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusIdle         LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus = "IDLE"
	LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusErrored      LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus = "ERRORED"
	LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusInvoked      LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus = "INVOKED"
	LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusMetricsReady LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus = "METRICS_READY"
)

func (r LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatus) IsKnown() bool {
	switch r {
	case LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusLive, LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusIdle, LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusErrored, LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusInvoked, LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataLivestreamStatusMetricsReady:
		return true
	}
	return false
}

type LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSession struct {
	// ID of the livestream session.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream session was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error message, if the livestream session failed.
	ErrMessage string `json:"err_message" api:"required,nullable"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds float64 `json:"ingest_seconds" api:"required"`
	// Timestamp the livestream session was invoked at, in ISO format.
	InvokedTime time.Time `json:"invoked_time" api:"required,nullable" format:"date-time"`
	// ID of the livestream.
	LivestreamID string `json:"livestream_id" api:"required" format:"uuid"`
	// Timestamp the livestream session started at, in ISO format.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session stopped at, in ISO format.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session was last updated at, in ISO format.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds float64                                                                    `json:"viewer_seconds" api:"required"`
	JSON          livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSessionJSON `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSessionJSON
// contains the JSON metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSession]
type livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSessionJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ErrMessage    apijson.Field
	IngestSeconds apijson.Field
	InvokedTime   apijson.Field
	LivestreamID  apijson.Field
	StartedTime   apijson.Field
	StoppedTime   apijson.Field
	UpdatedAt     apijson.Field
	ViewerSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForMeetingResponseDataDataSessionJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForMeetingResponseDataPaging struct {
	EndOffset   float64                                                               `json:"end_offset" api:"required"`
	StartOffset float64                                                               `json:"start_offset" api:"required"`
	TotalCount  float64                                                               `json:"total_count" api:"required"`
	JSON        livestreamGetLivestreamSessionDetailsForMeetingResponseDataPagingJSON `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForMeetingResponseDataPagingJSON contains
// the JSON metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForMeetingResponseDataPaging]
type livestreamGetLivestreamSessionDetailsForMeetingResponseDataPagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForMeetingResponseDataPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForMeetingResponseDataPagingJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForSessionIDResponse struct {
	Data    LivestreamGetLivestreamSessionDetailsForSessionIDResponseData `json:"data" api:"required"`
	Success bool                                                          `json:"success" api:"required"`
	JSON    livestreamGetLivestreamSessionDetailsForSessionIDResponseJSON `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForSessionIDResponseJSON contains the JSON
// metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForSessionIDResponse]
type livestreamGetLivestreamSessionDetailsForSessionIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForSessionIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForSessionIDResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionDetailsForSessionIDResponseData struct {
	// ID of the livestream session.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream session was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error message, if the livestream session failed.
	ErrMessage string `json:"err_message" api:"required,nullable"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds float64 `json:"ingest_seconds" api:"required"`
	// Timestamp the livestream session was invoked at, in ISO format.
	InvokedTime time.Time `json:"invoked_time" api:"required,nullable" format:"date-time"`
	// ID of the livestream.
	LivestreamID string `json:"livestream_id" api:"required" format:"uuid"`
	// Timestamp the livestream session started at, in ISO format.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session stopped at, in ISO format.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session was last updated at, in ISO format.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds float64                                                           `json:"viewer_seconds" api:"required"`
	JSON          livestreamGetLivestreamSessionDetailsForSessionIDResponseDataJSON `json:"-"`
}

// livestreamGetLivestreamSessionDetailsForSessionIDResponseDataJSON contains the
// JSON metadata for the struct
// [LivestreamGetLivestreamSessionDetailsForSessionIDResponseData]
type livestreamGetLivestreamSessionDetailsForSessionIDResponseDataJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ErrMessage    apijson.Field
	IngestSeconds apijson.Field
	InvokedTime   apijson.Field
	LivestreamID  apijson.Field
	StartedTime   apijson.Field
	StoppedTime   apijson.Field
	UpdatedAt     apijson.Field
	ViewerSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionDetailsForSessionIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionDetailsForSessionIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionForLivestreamIDResponse struct {
	Data    LivestreamGetLivestreamSessionForLivestreamIDResponseData   `json:"data" api:"required"`
	Paging  LivestreamGetLivestreamSessionForLivestreamIDResponsePaging `json:"paging" api:"required"`
	Success bool                                                        `json:"success" api:"required"`
	JSON    livestreamGetLivestreamSessionForLivestreamIDResponseJSON   `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseJSON contains the JSON
// metadata for the struct [LivestreamGetLivestreamSessionForLivestreamIDResponse]
type livestreamGetLivestreamSessionForLivestreamIDResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionForLivestreamIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionForLivestreamIDResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionForLivestreamIDResponseData struct {
	Livestream LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestream `json:"livestream" api:"required"`
	Sessions   []LivestreamGetLivestreamSessionForLivestreamIDResponseDataSession  `json:"sessions" api:"required"`
	JSON       livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON       `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON contains the JSON
// metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponseData]
type livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON struct {
	Livestream  apijson.Field
	Sessions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionForLivestreamIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestream struct {
	// ID of the livestream.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Specifies if the livestream is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// ID of the underlying Cloudflare Stream live input.
	ExternalID string `json:"external_id" api:"required"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server" api:"required"`
	// Specifies if the livestream is ingested manually by an RTMP encoder.
	ManualIngest bool `json:"manual_ingest" api:"required"`
	// ID of the meeting being livestreamed, if any.
	MeetingID string `json:"meeting_id" api:"required,nullable"`
	// Name of the livestream.
	Name string `json:"name" api:"required,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string `json:"playback_url" api:"required"`
	// Status of the livestream.
	Status LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus `json:"status" api:"required"`
	// Unique key for accessing each livestream. Empty if the livestream is disabled.
	StreamKey string `json:"stream_key" api:"required"`
	// Timestamp the livestream was last updated at, in ISO format.
	UpdatedAt time.Time                                                               `json:"updated_at" api:"required" format:"date-time"`
	JSON      livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON contains
// the JSON metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestream]
type livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	ExternalID   apijson.Field
	IngestServer apijson.Field
	ManualIngest apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON) RawJSON() string {
	return r.raw
}

// Status of the livestream.
type LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus string

const (
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusLive         LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "LIVE"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusIdle         LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "IDLE"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusErrored      LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "ERRORED"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusInvoked      LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "INVOKED"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusMetricsReady LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "METRICS_READY"
)

func (r LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus) IsKnown() bool {
	switch r {
	case LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusLive, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusIdle, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusErrored, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusInvoked, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusMetricsReady:
		return true
	}
	return false
}

type LivestreamGetLivestreamSessionForLivestreamIDResponseDataSession struct {
	// ID of the livestream session.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream session was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error message, if the livestream session failed.
	ErrMessage string `json:"err_message" api:"required,nullable"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds float64 `json:"ingest_seconds" api:"required"`
	// Timestamp the livestream session was invoked at, in ISO format.
	InvokedTime time.Time `json:"invoked_time" api:"required,nullable" format:"date-time"`
	// ID of the livestream.
	LivestreamID string `json:"livestream_id" api:"required" format:"uuid"`
	// Timestamp the livestream session started at, in ISO format.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session stopped at, in ISO format.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session was last updated at, in ISO format.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds float64                                                              `json:"viewer_seconds" api:"required"`
	JSON          livestreamGetLivestreamSessionForLivestreamIDResponseDataSessionJSON `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseDataSessionJSON contains
// the JSON metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponseDataSession]
type livestreamGetLivestreamSessionForLivestreamIDResponseDataSessionJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ErrMessage    apijson.Field
	IngestSeconds apijson.Field
	InvokedTime   apijson.Field
	LivestreamID  apijson.Field
	StartedTime   apijson.Field
	StoppedTime   apijson.Field
	UpdatedAt     apijson.Field
	ViewerSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionForLivestreamIDResponseDataSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionForLivestreamIDResponseDataSessionJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionForLivestreamIDResponsePaging struct {
	EndOffset   float64                                                         `json:"end_offset" api:"required"`
	StartOffset float64                                                         `json:"start_offset" api:"required"`
	TotalCount  float64                                                         `json:"total_count" api:"required"`
	JSON        livestreamGetLivestreamSessionForLivestreamIDResponsePagingJSON `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponsePagingJSON contains the
// JSON metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponsePaging]
type livestreamGetLivestreamSessionForLivestreamIDResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionForLivestreamIDResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionForLivestreamIDResponsePagingJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionsForSessionResponse struct {
	Data    LivestreamGetLivestreamSessionsForSessionResponseData `json:"data" api:"required"`
	Success bool                                                  `json:"success" api:"required"`
	JSON    livestreamGetLivestreamSessionsForSessionResponseJSON `json:"-"`
}

// livestreamGetLivestreamSessionsForSessionResponseJSON contains the JSON metadata
// for the struct [LivestreamGetLivestreamSessionsForSessionResponse]
type livestreamGetLivestreamSessionsForSessionResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionsForSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionsForSessionResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionsForSessionResponseData struct {
	Data   []LivestreamGetLivestreamSessionsForSessionResponseDataData `json:"data" api:"required"`
	Paging LivestreamGetLivestreamSessionsForSessionResponseDataPaging `json:"paging" api:"required"`
	JSON   livestreamGetLivestreamSessionsForSessionResponseDataJSON   `json:"-"`
}

// livestreamGetLivestreamSessionsForSessionResponseDataJSON contains the JSON
// metadata for the struct [LivestreamGetLivestreamSessionsForSessionResponseData]
type livestreamGetLivestreamSessionsForSessionResponseDataJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionsForSessionResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionsForSessionResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionsForSessionResponseDataData struct {
	// ID of the livestream session.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream session was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error message, if the livestream session failed.
	ErrMessage string `json:"err_message" api:"required,nullable"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds float64 `json:"ingest_seconds" api:"required"`
	// Timestamp the livestream session was invoked at, in ISO format.
	InvokedTime time.Time `json:"invoked_time" api:"required,nullable" format:"date-time"`
	// ID of the livestream.
	LivestreamID string `json:"livestream_id" api:"required" format:"uuid"`
	// Timestamp the livestream session started at, in ISO format.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session stopped at, in ISO format.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Timestamp the livestream session was last updated at, in ISO format.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds float64                                                       `json:"viewer_seconds" api:"required"`
	JSON          livestreamGetLivestreamSessionsForSessionResponseDataDataJSON `json:"-"`
}

// livestreamGetLivestreamSessionsForSessionResponseDataDataJSON contains the JSON
// metadata for the struct
// [LivestreamGetLivestreamSessionsForSessionResponseDataData]
type livestreamGetLivestreamSessionsForSessionResponseDataDataJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ErrMessage    apijson.Field
	IngestSeconds apijson.Field
	InvokedTime   apijson.Field
	LivestreamID  apijson.Field
	StartedTime   apijson.Field
	StoppedTime   apijson.Field
	UpdatedAt     apijson.Field
	ViewerSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionsForSessionResponseDataData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionsForSessionResponseDataDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionsForSessionResponseDataPaging struct {
	EndOffset   float64                                                         `json:"end_offset" api:"required"`
	StartOffset float64                                                         `json:"start_offset" api:"required"`
	TotalCount  float64                                                         `json:"total_count" api:"required"`
	JSON        livestreamGetLivestreamSessionsForSessionResponseDataPagingJSON `json:"-"`
}

// livestreamGetLivestreamSessionsForSessionResponseDataPagingJSON contains the
// JSON metadata for the struct
// [LivestreamGetLivestreamSessionsForSessionResponseDataPaging]
type livestreamGetLivestreamSessionsForSessionResponseDataPagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionsForSessionResponseDataPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionsForSessionResponseDataPagingJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetMeetingActiveLivestreamsResponse struct {
	Data    LivestreamGetMeetingActiveLivestreamsResponseData `json:"data" api:"required"`
	Success bool                                              `json:"success" api:"required"`
	JSON    livestreamGetMeetingActiveLivestreamsResponseJSON `json:"-"`
}

// livestreamGetMeetingActiveLivestreamsResponseJSON contains the JSON metadata for
// the struct [LivestreamGetMeetingActiveLivestreamsResponse]
type livestreamGetMeetingActiveLivestreamsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetMeetingActiveLivestreamsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetMeetingActiveLivestreamsResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetMeetingActiveLivestreamsResponseData struct {
	// ID of the livestream.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the livestream was created at, in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Specifies if the livestream is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// ID of the underlying Cloudflare Stream live input.
	ExternalID string `json:"external_id" api:"required"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server" api:"required"`
	// Specifies if the livestream is ingested manually by an RTMP encoder.
	ManualIngest bool `json:"manual_ingest" api:"required"`
	// ID of the meeting being livestreamed, if any.
	MeetingID string `json:"meeting_id" api:"required,nullable"`
	// Name of the livestream.
	Name string `json:"name" api:"required,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string `json:"playback_url" api:"required"`
	// Status of the livestream.
	Status LivestreamGetMeetingActiveLivestreamsResponseDataStatus `json:"status" api:"required"`
	// Unique key for accessing each livestream. Empty if the livestream is disabled.
	StreamKey string `json:"stream_key" api:"required"`
	// Timestamp the livestream was last updated at, in ISO format.
	UpdatedAt time.Time                                             `json:"updated_at" api:"required" format:"date-time"`
	JSON      livestreamGetMeetingActiveLivestreamsResponseDataJSON `json:"-"`
}

// livestreamGetMeetingActiveLivestreamsResponseDataJSON contains the JSON metadata
// for the struct [LivestreamGetMeetingActiveLivestreamsResponseData]
type livestreamGetMeetingActiveLivestreamsResponseDataJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	ExternalID   apijson.Field
	IngestServer apijson.Field
	ManualIngest apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamGetMeetingActiveLivestreamsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetMeetingActiveLivestreamsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status of the livestream.
type LivestreamGetMeetingActiveLivestreamsResponseDataStatus string

const (
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusLive         LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "LIVE"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusIdle         LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "IDLE"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusErrored      LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "ERRORED"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusInvoked      LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "INVOKED"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusMetricsReady LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "METRICS_READY"
)

func (r LivestreamGetMeetingActiveLivestreamsResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamGetMeetingActiveLivestreamsResponseDataStatusLive, LivestreamGetMeetingActiveLivestreamsResponseDataStatusIdle, LivestreamGetMeetingActiveLivestreamsResponseDataStatusErrored, LivestreamGetMeetingActiveLivestreamsResponseDataStatusInvoked, LivestreamGetMeetingActiveLivestreamsResponseDataStatusMetricsReady:
		return true
	}
	return false
}

type LivestreamGetOrgAnalyticsResponse struct {
	Data    LivestreamGetOrgAnalyticsResponseData `json:"data"`
	Success bool                                  `json:"success"`
	JSON    livestreamGetOrgAnalyticsResponseJSON `json:"-"`
}

// livestreamGetOrgAnalyticsResponseJSON contains the JSON metadata for the struct
// [LivestreamGetOrgAnalyticsResponse]
type livestreamGetOrgAnalyticsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetOrgAnalyticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetOrgAnalyticsResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetOrgAnalyticsResponseData struct {
	// Recording statistics of an App during the range specified
	RecordingStats LivestreamGetOrgAnalyticsResponseDataRecordingStats `json:"recording_stats"`
	// Session statistics of an App during the range specified
	SessionStats LivestreamGetOrgAnalyticsResponseDataSessionStats `json:"session_stats"`
	JSON         livestreamGetOrgAnalyticsResponseDataJSON         `json:"-"`
}

// livestreamGetOrgAnalyticsResponseDataJSON contains the JSON metadata for the
// struct [LivestreamGetOrgAnalyticsResponseData]
type livestreamGetOrgAnalyticsResponseDataJSON struct {
	RecordingStats apijson.Field
	SessionStats   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LivestreamGetOrgAnalyticsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetOrgAnalyticsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Recording statistics of an App during the range specified
type LivestreamGetOrgAnalyticsResponseDataRecordingStats struct {
	// Day wise recording stats
	DayStats []LivestreamGetOrgAnalyticsResponseDataRecordingStatsDayStat `json:"day_stats"`
	// Total number of recordings during the range specified
	RecordingCount int64 `json:"recording_count"`
	// Total recording minutes during the range specified
	RecordingMinutesConsumed float64                                                 `json:"recording_minutes_consumed"`
	JSON                     livestreamGetOrgAnalyticsResponseDataRecordingStatsJSON `json:"-"`
}

// livestreamGetOrgAnalyticsResponseDataRecordingStatsJSON contains the JSON
// metadata for the struct [LivestreamGetOrgAnalyticsResponseDataRecordingStats]
type livestreamGetOrgAnalyticsResponseDataRecordingStatsJSON struct {
	DayStats                 apijson.Field
	RecordingCount           apijson.Field
	RecordingMinutesConsumed apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LivestreamGetOrgAnalyticsResponseDataRecordingStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetOrgAnalyticsResponseDataRecordingStatsJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetOrgAnalyticsResponseDataRecordingStatsDayStat struct {
	Day string `json:"day"`
	// Total recording minutes for a specific day
	TotalRecordingMinutes int64 `json:"total_recording_minutes"`
	// Total number of recordings for a specific day
	TotalRecordings int64                                                          `json:"total_recordings"`
	JSON            livestreamGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON `json:"-"`
}

// livestreamGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON contains the JSON
// metadata for the struct
// [LivestreamGetOrgAnalyticsResponseDataRecordingStatsDayStat]
type livestreamGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON struct {
	Day                   apijson.Field
	TotalRecordingMinutes apijson.Field
	TotalRecordings       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *LivestreamGetOrgAnalyticsResponseDataRecordingStatsDayStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON) RawJSON() string {
	return r.raw
}

// Session statistics of an App during the range specified
type LivestreamGetOrgAnalyticsResponseDataSessionStats struct {
	// Day wise session stats
	DayStats []LivestreamGetOrgAnalyticsResponseDataSessionStatsDayStat `json:"day_stats"`
	// Total number of sessions during the range specified
	SessionsCount int64 `json:"sessions_count"`
	// Total session minutes during the range specified
	SessionsMinutesConsumed float64                                               `json:"sessions_minutes_consumed"`
	JSON                    livestreamGetOrgAnalyticsResponseDataSessionStatsJSON `json:"-"`
}

// livestreamGetOrgAnalyticsResponseDataSessionStatsJSON contains the JSON metadata
// for the struct [LivestreamGetOrgAnalyticsResponseDataSessionStats]
type livestreamGetOrgAnalyticsResponseDataSessionStatsJSON struct {
	DayStats                apijson.Field
	SessionsCount           apijson.Field
	SessionsMinutesConsumed apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LivestreamGetOrgAnalyticsResponseDataSessionStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetOrgAnalyticsResponseDataSessionStatsJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetOrgAnalyticsResponseDataSessionStatsDayStat struct {
	Day string `json:"day"`
	// Total session minutes for a specific day
	TotalSessionMinutes float64 `json:"total_session_minutes"`
	// Total number of sessions for a specific day
	TotalSessions int64                                                        `json:"total_sessions"`
	JSON          livestreamGetOrgAnalyticsResponseDataSessionStatsDayStatJSON `json:"-"`
}

// livestreamGetOrgAnalyticsResponseDataSessionStatsDayStatJSON contains the JSON
// metadata for the struct
// [LivestreamGetOrgAnalyticsResponseDataSessionStatsDayStat]
type livestreamGetOrgAnalyticsResponseDataSessionStatsDayStatJSON struct {
	Day                 apijson.Field
	TotalSessionMinutes apijson.Field
	TotalSessions       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LivestreamGetOrgAnalyticsResponseDataSessionStatsDayStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetOrgAnalyticsResponseDataSessionStatsDayStatJSON) RawJSON() string {
	return r.raw
}

type LivestreamStartLivestreamingAMeetingResponse struct {
	Data    LivestreamStartLivestreamingAMeetingResponseData `json:"data" api:"required"`
	Success bool                                             `json:"success" api:"required"`
	JSON    livestreamStartLivestreamingAMeetingResponseJSON `json:"-"`
}

// livestreamStartLivestreamingAMeetingResponseJSON contains the JSON metadata for
// the struct [LivestreamStartLivestreamingAMeetingResponse]
type livestreamStartLivestreamingAMeetingResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamStartLivestreamingAMeetingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamStartLivestreamingAMeetingResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamStartLivestreamingAMeetingResponseData struct {
	// ID of the livestream.
	ID string `json:"id" api:"required" format:"uuid"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server" api:"required"`
	// ID of the meeting being livestreamed, if any.
	MeetingID string `json:"meeting_id" api:"required,nullable"`
	// Name of the livestream.
	Name string `json:"name" api:"required,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string `json:"playback_url" api:"required"`
	// Status of the livestream.
	Status LivestreamStartLivestreamingAMeetingResponseDataStatus `json:"status" api:"required"`
	// Unique key for accessing each livestream. Empty if the livestream is disabled.
	StreamKey string                                               `json:"stream_key" api:"required"`
	JSON      livestreamStartLivestreamingAMeetingResponseDataJSON `json:"-"`
}

// livestreamStartLivestreamingAMeetingResponseDataJSON contains the JSON metadata
// for the struct [LivestreamStartLivestreamingAMeetingResponseData]
type livestreamStartLivestreamingAMeetingResponseDataJSON struct {
	ID           apijson.Field
	IngestServer apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamStartLivestreamingAMeetingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamStartLivestreamingAMeetingResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status of the livestream.
type LivestreamStartLivestreamingAMeetingResponseDataStatus string

const (
	LivestreamStartLivestreamingAMeetingResponseDataStatusLive         LivestreamStartLivestreamingAMeetingResponseDataStatus = "LIVE"
	LivestreamStartLivestreamingAMeetingResponseDataStatusIdle         LivestreamStartLivestreamingAMeetingResponseDataStatus = "IDLE"
	LivestreamStartLivestreamingAMeetingResponseDataStatusErrored      LivestreamStartLivestreamingAMeetingResponseDataStatus = "ERRORED"
	LivestreamStartLivestreamingAMeetingResponseDataStatusInvoked      LivestreamStartLivestreamingAMeetingResponseDataStatus = "INVOKED"
	LivestreamStartLivestreamingAMeetingResponseDataStatusMetricsReady LivestreamStartLivestreamingAMeetingResponseDataStatus = "METRICS_READY"
)

func (r LivestreamStartLivestreamingAMeetingResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamStartLivestreamingAMeetingResponseDataStatusLive, LivestreamStartLivestreamingAMeetingResponseDataStatusIdle, LivestreamStartLivestreamingAMeetingResponseDataStatusErrored, LivestreamStartLivestreamingAMeetingResponseDataStatusInvoked, LivestreamStartLivestreamingAMeetingResponseDataStatusMetricsReady:
		return true
	}
	return false
}

type LivestreamStopLivestreamingAMeetingResponse struct {
	Data    LivestreamStopLivestreamingAMeetingResponseData `json:"data"`
	Success bool                                            `json:"success"`
	JSON    livestreamStopLivestreamingAMeetingResponseJSON `json:"-"`
}

// livestreamStopLivestreamingAMeetingResponseJSON contains the JSON metadata for
// the struct [LivestreamStopLivestreamingAMeetingResponse]
type livestreamStopLivestreamingAMeetingResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamStopLivestreamingAMeetingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamStopLivestreamingAMeetingResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamStopLivestreamingAMeetingResponseData struct {
	Message string                                              `json:"message"`
	JSON    livestreamStopLivestreamingAMeetingResponseDataJSON `json:"-"`
}

// livestreamStopLivestreamingAMeetingResponseDataJSON contains the JSON metadata
// for the struct [LivestreamStopLivestreamingAMeetingResponseData]
type livestreamStopLivestreamingAMeetingResponseDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamStopLivestreamingAMeetingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamStopLivestreamingAMeetingResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetActiveLivestreamsForLivestreamIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type LivestreamGetAllLivestreamsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Specify the end time range in ISO format to access the live stream.
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// Exclude the RealtimeKit meetings that are livestreamed.
	ExcludeMeetings param.Field[bool] `query:"exclude_meetings"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[int64] `query:"page_no"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Specifies the sorting order for the results.
	SortOrder param.Field[LivestreamGetAllLivestreamsParamsSortOrder] `query:"sort_order"`
	// Specify the start time range in ISO format to access the live stream.
	StartTime param.Field[time.Time] `query:"start_time" format:"date-time"`
	// Specifies the status of the operation.
	Status param.Field[LivestreamGetAllLivestreamsParamsStatus] `query:"status"`
}

// URLQuery serializes [LivestreamGetAllLivestreamsParams]'s query parameters as
// `url.Values`.
func (r LivestreamGetAllLivestreamsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the sorting order for the results.
type LivestreamGetAllLivestreamsParamsSortOrder string

const (
	LivestreamGetAllLivestreamsParamsSortOrderAsc LivestreamGetAllLivestreamsParamsSortOrder = "ASC"
	LivestreamGetAllLivestreamsParamsSortOrderDsc LivestreamGetAllLivestreamsParamsSortOrder = "DSC"
)

func (r LivestreamGetAllLivestreamsParamsSortOrder) IsKnown() bool {
	switch r {
	case LivestreamGetAllLivestreamsParamsSortOrderAsc, LivestreamGetAllLivestreamsParamsSortOrderDsc:
		return true
	}
	return false
}

// Specifies the status of the operation.
type LivestreamGetAllLivestreamsParamsStatus string

const (
	LivestreamGetAllLivestreamsParamsStatusLive    LivestreamGetAllLivestreamsParamsStatus = "LIVE"
	LivestreamGetAllLivestreamsParamsStatusIdle    LivestreamGetAllLivestreamsParamsStatus = "IDLE"
	LivestreamGetAllLivestreamsParamsStatusErrored LivestreamGetAllLivestreamsParamsStatus = "ERRORED"
	LivestreamGetAllLivestreamsParamsStatusInvoked LivestreamGetAllLivestreamsParamsStatus = "INVOKED"
)

func (r LivestreamGetAllLivestreamsParamsStatus) IsKnown() bool {
	switch r {
	case LivestreamGetAllLivestreamsParamsStatusLive, LivestreamGetAllLivestreamsParamsStatusIdle, LivestreamGetAllLivestreamsParamsStatusErrored, LivestreamGetAllLivestreamsParamsStatusInvoked:
		return true
	}
	return false
}

type LivestreamGetLivestreamAnalyticsCompleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Specify the end time as a Unix timestamp in seconds to access the livestream
	// analytics.
	EndTime param.Field[int64] `query:"end_time"`
	// Optional filters for livestream analytics.
	Filters param.Field[string] `query:"filters"`
	// Specify the start time as a Unix timestamp in seconds to access the livestream
	// analytics.
	StartTime param.Field[int64] `query:"start_time"`
}

// URLQuery serializes [LivestreamGetLivestreamAnalyticsCompleteParams]'s query
// parameters as `url.Values`.
func (r LivestreamGetLivestreamAnalyticsCompleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamGetLivestreamAnalyticsDaywiseParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Specify the end time as a Unix timestamp in seconds to access the livestream
	// analytics.
	EndTime param.Field[int64] `query:"end_time"`
	// Optional filters for livestream analytics.
	Filters param.Field[string] `query:"filters"`
	// Specify the start time as a Unix timestamp in seconds to access the livestream
	// analytics.
	StartTime param.Field[int64] `query:"start_time"`
}

// URLQuery serializes [LivestreamGetLivestreamAnalyticsDaywiseParams]'s query
// parameters as `url.Values`.
func (r LivestreamGetLivestreamAnalyticsDaywiseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamGetLivestreamSessionDetailsForMeetingParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[int64] `query:"page_no"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [LivestreamGetLivestreamSessionDetailsForMeetingParams]'s
// query parameters as `url.Values`.
func (r LivestreamGetLivestreamSessionDetailsForMeetingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamGetLivestreamSessionDetailsForSessionIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type LivestreamGetLivestreamSessionForLivestreamIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[int64] `query:"page_no"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [LivestreamGetLivestreamSessionForLivestreamIDParams]'s
// query parameters as `url.Values`.
func (r LivestreamGetLivestreamSessionForLivestreamIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamGetLivestreamSessionsForSessionParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [LivestreamGetLivestreamSessionsForSessionParams]'s query
// parameters as `url.Values`.
func (r LivestreamGetLivestreamSessionsForSessionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamGetMeetingActiveLivestreamsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type LivestreamGetOrgAnalyticsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// end date in YYYY-MM-DD format
	EndDate param.Field[string] `query:"end_date"`
	// start date in YYYY-MM-DD format
	StartDate param.Field[string] `query:"start_date"`
}

// URLQuery serializes [LivestreamGetOrgAnalyticsParams]'s query parameters as
// `url.Values`.
func (r LivestreamGetOrgAnalyticsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamStartLivestreamingAMeetingParams struct {
	// The account identifier tag.
	AccountID   param.Field[string]                                                `path:"account_id" api:"required"`
	Name        param.Field[string]                                                `json:"name"`
	VideoConfig param.Field[LivestreamStartLivestreamingAMeetingParamsVideoConfig] `json:"video_config"`
}

func (r LivestreamStartLivestreamingAMeetingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LivestreamStartLivestreamingAMeetingParamsVideoConfig struct {
	// Height of the livestreaming video in pixels
	Height param.Field[int64] `json:"height"`
	// Width of the livestreaming video in pixels
	Width param.Field[int64] `json:"width"`
}

func (r LivestreamStartLivestreamingAMeetingParamsVideoConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LivestreamStopLivestreamingAMeetingParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
