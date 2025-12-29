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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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

// Creates a livestream for the given App ID and returns ingest server, stream key,
// and playback URL. You can pass custom input to the ingest server and stream key,
// and freely distribute the content using the playback URL on any player that
// supports HLS/LHLS.
func (r *LivestreamService) NewIndependentLivestream(ctx context.Context, appID string, params LivestreamNewIndependentLivestreamParams, opts ...option.RequestOption) (res *LivestreamNewIndependentLivestreamResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns details of all active livestreams for the given livestream ID. Retreive
// the livestream ID using the `Start livestreaming a meeting` API.
func (r *LivestreamService) GetActiveLivestreamsForLivestreamID(ctx context.Context, appID string, livestreamID string, query LivestreamGetActiveLivestreamsForLivestreamIDParams, opts ...option.RequestOption) (res *LivestreamGetActiveLivestreamsForLivestreamIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if livestreamID == "" {
		err = errors.New("missing required livestream_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams/%s/active-livestream-session", query.AccountID, appID, livestreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns details of livestreams associated with the given App ID. It includes
// livestreams created by your App and RealtimeKit meetings that are livestreamed
// by your App. If you only want details of livestreams created by your App and not
// RealtimeKit meetings, you can use the `exclude_meetings` query parameter.
func (r *LivestreamService) GetAllLivestreams(ctx context.Context, appID string, params LivestreamGetAllLivestreamsParams, opts ...option.RequestOption) (res *LivestreamGetAllLivestreamsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns livestream analytics for the specified time range.
func (r *LivestreamService) GetLivestreamAnalyticsComplete(ctx context.Context, appID string, params LivestreamGetLivestreamAnalyticsCompleteParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamAnalyticsCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/analytics/livestreams/overall", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns livestream session details for the given livestream session ID. Retrieve
// the `livestream_session_id`using the
// `Fetch livestream session details using a session ID` API.
func (r *LivestreamService) GetLivestreamSessionDetailsForSessionID(ctx context.Context, appID string, livestreamSessionID string, query LivestreamGetLivestreamSessionDetailsForSessionIDParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamSessionDetailsForSessionIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if livestreamSessionID == "" {
		err = errors.New("missing required livestream-session-id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams/sessions/%s", query.AccountID, appID, livestreamSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns details of a livestream with sessions for the given livestream ID.
// Retreive the livestream ID using the `Start livestreaming a meeting` API.
func (r *LivestreamService) GetLivestreamSessionForLivestreamID(ctx context.Context, appID string, livestreamID string, params LivestreamGetLivestreamSessionForLivestreamIDParams, opts ...option.RequestOption) (res *LivestreamGetLivestreamSessionForLivestreamIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if livestreamID == "" {
		err = errors.New("missing required livestream_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/livestreams/%s", params.AccountID, appID, livestreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns details of all active livestreams for the given meeting ID.
func (r *LivestreamService) GetMeetingActiveLivestreams(ctx context.Context, appID string, meetingID string, query LivestreamGetMeetingActiveLivestreamsParams, opts ...option.RequestOption) (res *LivestreamGetMeetingActiveLivestreamsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-livestream", query.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns day-wise session and recording analytics data of an App for the
// specified time range start_date to end_date. If start_date and end_date are not
// provided, the default time range is set from 30 days ago to the current date.
func (r *LivestreamService) GetOrgAnalytics(ctx context.Context, appID string, params LivestreamGetOrgAnalyticsParams, opts ...option.RequestOption) (res *LivestreamGetOrgAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/analytics/daywise", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Starts livestream of a meeting associated with the given meeting ID. Retreive
// the meeting ID using the `Create a meeting` API.
func (r *LivestreamService) StartLivestreamingAMeeting(ctx context.Context, appID string, meetingID string, params LivestreamStartLivestreamingAMeetingParams, opts ...option.RequestOption) (res *LivestreamStartLivestreamingAMeetingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/livestreams", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Stops the active livestream of a meeting associated with the given meeting ID.
// Retreive the meeting ID using the `Create a meeting` API.
func (r *LivestreamService) StopLivestreamingAMeeting(ctx context.Context, appID string, meetingID string, body LivestreamStopLivestreamingAMeetingParams, opts ...option.RequestOption) (res *LivestreamStopLivestreamingAMeetingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-livestream/stop", body.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type LivestreamNewIndependentLivestreamResponse struct {
	Data    LivestreamNewIndependentLivestreamResponseData `json:"data"`
	Success bool                                           `json:"success"`
	JSON    livestreamNewIndependentLivestreamResponseJSON `json:"-"`
}

// livestreamNewIndependentLivestreamResponseJSON contains the JSON metadata for
// the struct [LivestreamNewIndependentLivestreamResponse]
type livestreamNewIndependentLivestreamResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamNewIndependentLivestreamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamNewIndependentLivestreamResponseJSON) RawJSON() string {
	return r.raw
}

type LivestreamNewIndependentLivestreamResponseData struct {
	// The livestream ID.
	ID string `json:"id"`
	// Specifies if the livestream was disabled.
	Disabled bool `json:"disabled"`
	// The server URL to which the RTMP encoder should send the video and audio data.
	IngestServer string `json:"ingest_server"`
	MeetingID    string `json:"meeting_id,nullable"`
	Name         string `json:"name"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string                                               `json:"playback_url"`
	Status      LivestreamNewIndependentLivestreamResponseDataStatus `json:"status"`
	// Unique key for accessing each livestream.
	StreamKey string                                             `json:"stream_key"`
	JSON      livestreamNewIndependentLivestreamResponseDataJSON `json:"-"`
}

// livestreamNewIndependentLivestreamResponseDataJSON contains the JSON metadata
// for the struct [LivestreamNewIndependentLivestreamResponseData]
type livestreamNewIndependentLivestreamResponseDataJSON struct {
	ID           apijson.Field
	Disabled     apijson.Field
	IngestServer apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	PlaybackURL  apijson.Field
	Status       apijson.Field
	StreamKey    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LivestreamNewIndependentLivestreamResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamNewIndependentLivestreamResponseDataJSON) RawJSON() string {
	return r.raw
}

type LivestreamNewIndependentLivestreamResponseDataStatus string

const (
	LivestreamNewIndependentLivestreamResponseDataStatusLive    LivestreamNewIndependentLivestreamResponseDataStatus = "LIVE"
	LivestreamNewIndependentLivestreamResponseDataStatusIdle    LivestreamNewIndependentLivestreamResponseDataStatus = "IDLE"
	LivestreamNewIndependentLivestreamResponseDataStatusErrored LivestreamNewIndependentLivestreamResponseDataStatus = "ERRORED"
	LivestreamNewIndependentLivestreamResponseDataStatusInvoked LivestreamNewIndependentLivestreamResponseDataStatus = "INVOKED"
)

func (r LivestreamNewIndependentLivestreamResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamNewIndependentLivestreamResponseDataStatusLive, LivestreamNewIndependentLivestreamResponseDataStatusIdle, LivestreamNewIndependentLivestreamResponseDataStatusErrored, LivestreamNewIndependentLivestreamResponseDataStatusInvoked:
		return true
	}
	return false
}

type LivestreamGetActiveLivestreamsForLivestreamIDResponse struct {
	Data    LivestreamGetActiveLivestreamsForLivestreamIDResponseData `json:"data"`
	Success bool                                                      `json:"success"`
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
	Livestream LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestream `json:"livestream"`
	Session    LivestreamGetActiveLivestreamsForLivestreamIDResponseDataSession    `json:"session"`
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
	ID string `json:"id"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies if the livestream was disabled.
	Disabled string `json:"disabled"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server"`
	// ID of the meeting.
	MeetingID string `json:"meeting_id"`
	// Name of the livestream.
	Name string `json:"name"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string                                                                    `json:"playback_url"`
	Status      LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus `json:"status"`
	// Unique key for accessing each livestream.
	StreamKey string `json:"stream_key"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time                                                               `json:"updated_at" format:"date-time"`
	JSON      livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON `json:"-"`
}

// livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON contains
// the JSON metadata for the struct
// [LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestream]
type livestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	IngestServer apijson.Field
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

type LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus string

const (
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusLive    LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "LIVE"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusIdle    LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "IDLE"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusErrored LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "ERRORED"
	LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusInvoked LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus = "INVOKED"
)

func (r LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatus) IsKnown() bool {
	switch r {
	case LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusLive, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusIdle, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusErrored, LivestreamGetActiveLivestreamsForLivestreamIDResponseDataLivestreamStatusInvoked:
		return true
	}
	return false
}

type LivestreamGetActiveLivestreamsForLivestreamIDResponseDataSession struct {
	ID string `json:"id"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	ErrMessage string    `json:"err_message"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds string `json:"ingest_seconds"`
	// Timestamp the object was invoked. The time is returned in ISO format.
	InvokedTime  time.Time `json:"invoked_time" format:"date-time"`
	LivestreamID string    `json:"livestream_id"`
	// Timestamp the object was started. The time is returned in ISO format.
	StartedTime time.Time `json:"started_time" format:"date-time"`
	// Timestamp the object was stopped. The time is returned in ISO format.
	StoppedTime time.Time `json:"stopped_time" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds string                                                               `json:"viewer_seconds"`
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
	Data    LivestreamGetAllLivestreamsResponseData `json:"data"`
	Success bool                                    `json:"success"`
	JSON    livestreamGetAllLivestreamsResponseJSON `json:"-"`
}

// livestreamGetAllLivestreamsResponseJSON contains the JSON metadata for the
// struct [LivestreamGetAllLivestreamsResponse]
type livestreamGetAllLivestreamsResponseJSON struct {
	Data        apijson.Field
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
	// The ID of the livestream.
	ID string `json:"id" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies if the livestream was disabled.
	Disabled string `json:"disabled"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server"`
	// ID of the meeting.
	MeetingID string `json:"meeting_id"`
	// Name of the livestream.
	Name   string                                        `json:"name"`
	Paging LivestreamGetAllLivestreamsResponseDataPaging `json:"paging"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string                                        `json:"playback_url"`
	Status      LivestreamGetAllLivestreamsResponseDataStatus `json:"status"`
	// Unique key for accessing each livestream.
	StreamKey string `json:"stream_key"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      livestreamGetAllLivestreamsResponseDataJSON `json:"-"`
}

// livestreamGetAllLivestreamsResponseDataJSON contains the JSON metadata for the
// struct [LivestreamGetAllLivestreamsResponseData]
type livestreamGetAllLivestreamsResponseDataJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	IngestServer apijson.Field
	MeetingID    apijson.Field
	Name         apijson.Field
	Paging       apijson.Field
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

type LivestreamGetAllLivestreamsResponseDataPaging struct {
	EndOffset   int64                                             `json:"end_offset"`
	StartOffset int64                                             `json:"start_offset"`
	TotalCount  int64                                             `json:"total_count"`
	JSON        livestreamGetAllLivestreamsResponseDataPagingJSON `json:"-"`
}

// livestreamGetAllLivestreamsResponseDataPagingJSON contains the JSON metadata for
// the struct [LivestreamGetAllLivestreamsResponseDataPaging]
type livestreamGetAllLivestreamsResponseDataPagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetAllLivestreamsResponseDataPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetAllLivestreamsResponseDataPagingJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetAllLivestreamsResponseDataStatus string

const (
	LivestreamGetAllLivestreamsResponseDataStatusLive    LivestreamGetAllLivestreamsResponseDataStatus = "LIVE"
	LivestreamGetAllLivestreamsResponseDataStatusIdle    LivestreamGetAllLivestreamsResponseDataStatus = "IDLE"
	LivestreamGetAllLivestreamsResponseDataStatusErrored LivestreamGetAllLivestreamsResponseDataStatus = "ERRORED"
	LivestreamGetAllLivestreamsResponseDataStatusInvoked LivestreamGetAllLivestreamsResponseDataStatus = "INVOKED"
)

func (r LivestreamGetAllLivestreamsResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamGetAllLivestreamsResponseDataStatusLive, LivestreamGetAllLivestreamsResponseDataStatusIdle, LivestreamGetAllLivestreamsResponseDataStatusErrored, LivestreamGetAllLivestreamsResponseDataStatusInvoked:
		return true
	}
	return false
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
	TotalIngestSeconds int64 `json:"total_ingest_seconds"`
	// Total view time for which the viewers watched the stream.
	TotalViewerSeconds int64                                                    `json:"total_viewer_seconds"`
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

type LivestreamGetLivestreamSessionDetailsForSessionIDResponse struct {
	Data    LivestreamGetLivestreamSessionDetailsForSessionIDResponseData `json:"data"`
	Success bool                                                          `json:"success"`
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
	// The livestream ID.
	ID string `json:"id"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	ErrMessage string `json:"err_message"`
	// Name of the livestream.
	IngestSeconds int64  `json:"ingest_seconds"`
	LivestreamID  string `json:"livestream_id"`
	// Unique key for accessing each livestream.
	StartedTime string `json:"started_time"`
	// The web address that viewers can use to watch the livestream.
	StoppedTime string `json:"stopped_time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt string `json:"updated_at"`
	// Specifies if the livestream was disabled.
	ViewerSeconds int64                                                             `json:"viewer_seconds"`
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
	Data    LivestreamGetLivestreamSessionForLivestreamIDResponseData `json:"data"`
	Success bool                                                      `json:"success"`
	JSON    livestreamGetLivestreamSessionForLivestreamIDResponseJSON `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseJSON contains the JSON
// metadata for the struct [LivestreamGetLivestreamSessionForLivestreamIDResponse]
type livestreamGetLivestreamSessionForLivestreamIDResponseJSON struct {
	Data        apijson.Field
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
	Livestream LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestream `json:"livestream"`
	Paging     LivestreamGetLivestreamSessionForLivestreamIDResponseDataPaging     `json:"paging"`
	Session    LivestreamGetLivestreamSessionForLivestreamIDResponseDataSession    `json:"session"`
	JSON       livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON       `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON contains the JSON
// metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponseData]
type livestreamGetLivestreamSessionForLivestreamIDResponseDataJSON struct {
	Livestream  apijson.Field
	Paging      apijson.Field
	Session     apijson.Field
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
	ID string `json:"id"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt string `json:"created_at"`
	// Specifies if the livestream was disabled.
	Disabled string `json:"disabled"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server"`
	// The ID of the meeting.
	MeetingID string `json:"meeting_id"`
	// Name of the livestream.
	Name string `json:"name"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string                                                                    `json:"playback_url"`
	Status      LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus `json:"status"`
	// Unique key for accessing each livestream.
	StreamKey string `json:"stream_key"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt string                                                                  `json:"updated_at"`
	JSON      livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON contains
// the JSON metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestream]
type livestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	IngestServer apijson.Field
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

type LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus string

const (
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusLive    LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "LIVE"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusIdle    LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "IDLE"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusErrored LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "ERRORED"
	LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusInvoked LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus = "INVOKED"
)

func (r LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatus) IsKnown() bool {
	switch r {
	case LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusLive, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusIdle, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusErrored, LivestreamGetLivestreamSessionForLivestreamIDResponseDataLivestreamStatusInvoked:
		return true
	}
	return false
}

type LivestreamGetLivestreamSessionForLivestreamIDResponseDataPaging struct {
	EndOffset   int64                                                               `json:"end_offset"`
	StartOffset int64                                                               `json:"start_offset"`
	TotalCount  int64                                                               `json:"total_count"`
	JSON        livestreamGetLivestreamSessionForLivestreamIDResponseDataPagingJSON `json:"-"`
}

// livestreamGetLivestreamSessionForLivestreamIDResponseDataPagingJSON contains the
// JSON metadata for the struct
// [LivestreamGetLivestreamSessionForLivestreamIDResponseDataPaging]
type livestreamGetLivestreamSessionForLivestreamIDResponseDataPagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LivestreamGetLivestreamSessionForLivestreamIDResponseDataPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r livestreamGetLivestreamSessionForLivestreamIDResponseDataPagingJSON) RawJSON() string {
	return r.raw
}

type LivestreamGetLivestreamSessionForLivestreamIDResponseDataSession struct {
	// ID of the session.
	ID string `json:"id"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	ErrMessage string    `json:"err_message"`
	// The time duration for which the input was given or the meeting was streamed.
	IngestSeconds float64 `json:"ingest_seconds"`
	// Timestamp the object was invoked. The time is returned in ISO format.
	InvokedTime  time.Time `json:"invoked_time" format:"date-time"`
	LivestreamID string    `json:"livestream_id"`
	// Timestamp the object was started. The time is returned in ISO format.
	StartedTime time.Time `json:"started_time" format:"date-time"`
	// Timestamp the object was stopped. The time is returned in ISO format.
	StoppedTime time.Time `json:"stopped_time" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The total view time for which the viewers watched the stream.
	ViewerSeconds float64                                                              `json:"viewer_seconds"`
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

type LivestreamGetMeetingActiveLivestreamsResponse struct {
	Data    LivestreamGetMeetingActiveLivestreamsResponseData `json:"data"`
	Success bool                                              `json:"success"`
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
	// The livestream ID.
	ID string `json:"id"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies if the livestream was disabled.
	Disabled string `json:"disabled"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server"`
	MeetingID    string `json:"meeting_id"`
	// Name of the livestream.
	Name string `json:"name,nullable"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string                                                  `json:"playback_url"`
	Status      LivestreamGetMeetingActiveLivestreamsResponseDataStatus `json:"status"`
	// Unique key for accessing each livestream.
	StreamKey string `json:"stream_key"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      livestreamGetMeetingActiveLivestreamsResponseDataJSON `json:"-"`
}

// livestreamGetMeetingActiveLivestreamsResponseDataJSON contains the JSON metadata
// for the struct [LivestreamGetMeetingActiveLivestreamsResponseData]
type livestreamGetMeetingActiveLivestreamsResponseDataJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Disabled     apijson.Field
	IngestServer apijson.Field
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

type LivestreamGetMeetingActiveLivestreamsResponseDataStatus string

const (
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusLive    LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "LIVE"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusIdle    LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "IDLE"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusErrored LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "ERRORED"
	LivestreamGetMeetingActiveLivestreamsResponseDataStatusInvoked LivestreamGetMeetingActiveLivestreamsResponseDataStatus = "INVOKED"
)

func (r LivestreamGetMeetingActiveLivestreamsResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamGetMeetingActiveLivestreamsResponseDataStatusLive, LivestreamGetMeetingActiveLivestreamsResponseDataStatusIdle, LivestreamGetMeetingActiveLivestreamsResponseDataStatusErrored, LivestreamGetMeetingActiveLivestreamsResponseDataStatusInvoked:
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
	Data    LivestreamStartLivestreamingAMeetingResponseData `json:"data"`
	Success bool                                             `json:"success"`
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
	// The livestream ID.
	ID string `json:"id"`
	// The server URL to which the RTMP encoder sends the video and audio data.
	IngestServer string `json:"ingest_server"`
	// The web address that viewers can use to watch the livestream.
	PlaybackURL string                                                 `json:"playback_url"`
	Status      LivestreamStartLivestreamingAMeetingResponseDataStatus `json:"status"`
	// Unique key for accessing each livestream.
	StreamKey string                                               `json:"stream_key"`
	JSON      livestreamStartLivestreamingAMeetingResponseDataJSON `json:"-"`
}

// livestreamStartLivestreamingAMeetingResponseDataJSON contains the JSON metadata
// for the struct [LivestreamStartLivestreamingAMeetingResponseData]
type livestreamStartLivestreamingAMeetingResponseDataJSON struct {
	ID           apijson.Field
	IngestServer apijson.Field
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

type LivestreamStartLivestreamingAMeetingResponseDataStatus string

const (
	LivestreamStartLivestreamingAMeetingResponseDataStatusLive    LivestreamStartLivestreamingAMeetingResponseDataStatus = "LIVE"
	LivestreamStartLivestreamingAMeetingResponseDataStatusIdle    LivestreamStartLivestreamingAMeetingResponseDataStatus = "IDLE"
	LivestreamStartLivestreamingAMeetingResponseDataStatusErrored LivestreamStartLivestreamingAMeetingResponseDataStatus = "ERRORED"
	LivestreamStartLivestreamingAMeetingResponseDataStatusInvoked LivestreamStartLivestreamingAMeetingResponseDataStatus = "INVOKED"
)

func (r LivestreamStartLivestreamingAMeetingResponseDataStatus) IsKnown() bool {
	switch r {
	case LivestreamStartLivestreamingAMeetingResponseDataStatusLive, LivestreamStartLivestreamingAMeetingResponseDataStatusIdle, LivestreamStartLivestreamingAMeetingResponseDataStatusErrored, LivestreamStartLivestreamingAMeetingResponseDataStatusInvoked:
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

type LivestreamNewIndependentLivestreamParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the livestream
	Name param.Field[string] `json:"name"`
}

func (r LivestreamNewIndependentLivestreamParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LivestreamGetActiveLivestreamsForLivestreamIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type LivestreamGetAllLivestreamsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	// Specify the end time range in ISO format to access the livestream analytics.
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// Specify the start time range in ISO format to access the livestream analytics.
	StartTime param.Field[time.Time] `query:"start_time" format:"date-time"`
}

// URLQuery serializes [LivestreamGetLivestreamAnalyticsCompleteParams]'s query
// parameters as `url.Values`.
func (r LivestreamGetLivestreamAnalyticsCompleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LivestreamGetLivestreamSessionDetailsForSessionIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type LivestreamGetLivestreamSessionForLivestreamIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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

type LivestreamGetMeetingActiveLivestreamsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type LivestreamGetOrgAnalyticsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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
	AccountID   param.Field[string]                                                `path:"account_id,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}
