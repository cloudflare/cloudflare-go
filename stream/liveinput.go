// File generated from our OpenAPI spec by Stainless.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// LiveInputService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewLiveInputService] method instead.
type LiveInputService struct {
	Options []option.RequestOption
	Outputs *LiveInputOutputService
}

// NewLiveInputService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLiveInputService(opts ...option.RequestOption) (r *LiveInputService) {
	r = &LiveInputService{}
	r.Options = opts
	r.Outputs = NewLiveInputOutputService(opts...)
	return
}

// Creates a live input, and returns credentials that you or your users can use to
// stream live video to Cloudflare Stream.
func (r *LiveInputService) New(ctx context.Context, params LiveInputNewParams, opts ...option.RequestOption) (res *LiveInputNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specified live input.
func (r *LiveInputService) Update(ctx context.Context, liveInputIdentifier string, params LiveInputUpdateParams, opts ...option.RequestOption) (res *LiveInputUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", params.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *LiveInputService) List(ctx context.Context, params LiveInputListParams, opts ...option.RequestOption) (res *LiveInputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Prevents a live input from being streamed to and makes the live input
// inaccessible to any future API calls.
func (r *LiveInputService) Delete(ctx context.Context, liveInputIdentifier string, body LiveInputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", body.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Retrieves details of an existing live input.
func (r *LiveInputService) Get(ctx context.Context, liveInputIdentifier string, query LiveInputGetParams, opts ...option.RequestOption) (res *LiveInputGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", query.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about a live input.
type LiveInputNewResponse struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording LiveInputNewResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps LiveInputNewResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback LiveInputNewResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt LiveInputNewResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback LiveInputNewResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status LiveInputNewResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc LiveInputNewResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback LiveInputNewResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           liveInputNewResponseJSON           `json:"-"`
}

// liveInputNewResponseJSON contains the JSON metadata for the struct
// [LiveInputNewResponse]
type liveInputNewResponseJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Recording                apijson.Field
	Rtmps                    apijson.Field
	RtmpsPlayback            apijson.Field
	Srt                      apijson.Field
	SrtPlayback              apijson.Field
	Status                   apijson.Field
	Uid                      apijson.Field
	WebRtc                   apijson.Field
	WebRtcPlayback           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInputNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseJSON) RawJSON() string {
	return r.raw
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputNewResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode LiveInputNewResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                             `json:"timeoutSeconds"`
	JSON           liveInputNewResponseRecordingJSON `json:"-"`
}

// liveInputNewResponseRecordingJSON contains the JSON metadata for the struct
// [LiveInputNewResponseRecording]
type liveInputNewResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LiveInputNewResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseRecordingJSON) RawJSON() string {
	return r.raw
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputNewResponseRecordingMode string

const (
	LiveInputNewResponseRecordingModeOff       LiveInputNewResponseRecordingMode = "off"
	LiveInputNewResponseRecordingModeAutomatic LiveInputNewResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type LiveInputNewResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                        `json:"url"`
	JSON liveInputNewResponseRtmpsJSON `json:"-"`
}

// liveInputNewResponseRtmpsJSON contains the JSON metadata for the struct
// [LiveInputNewResponseRtmps]
type liveInputNewResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseRtmpsJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using RTMPS.
type LiveInputNewResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                `json:"url"`
	JSON liveInputNewResponseRtmpsPlaybackJSON `json:"-"`
}

// liveInputNewResponseRtmpsPlaybackJSON contains the JSON metadata for the struct
// [LiveInputNewResponseRtmpsPlayback]
type liveInputNewResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseRtmpsPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to a live input using SRT.
type LiveInputNewResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                      `json:"url"`
	JSON liveInputNewResponseSrtJSON `json:"-"`
}

// liveInputNewResponseSrtJSON contains the JSON metadata for the struct
// [LiveInputNewResponseSrt]
type liveInputNewResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseSrtJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using SRT.
type LiveInputNewResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                              `json:"url"`
	JSON liveInputNewResponseSrtPlaybackJSON `json:"-"`
}

// liveInputNewResponseSrtPlaybackJSON contains the JSON metadata for the struct
// [LiveInputNewResponseSrtPlayback]
type liveInputNewResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseSrtPlaybackJSON) RawJSON() string {
	return r.raw
}

// The connection status of a live input.
type LiveInputNewResponseStatus string

const (
	LiveInputNewResponseStatusConnected                LiveInputNewResponseStatus = "connected"
	LiveInputNewResponseStatusReconnected              LiveInputNewResponseStatus = "reconnected"
	LiveInputNewResponseStatusReconnecting             LiveInputNewResponseStatus = "reconnecting"
	LiveInputNewResponseStatusClientDisconnect         LiveInputNewResponseStatus = "client_disconnect"
	LiveInputNewResponseStatusTTLExceeded              LiveInputNewResponseStatus = "ttl_exceeded"
	LiveInputNewResponseStatusFailedToConnect          LiveInputNewResponseStatus = "failed_to_connect"
	LiveInputNewResponseStatusFailedToReconnect        LiveInputNewResponseStatus = "failed_to_reconnect"
	LiveInputNewResponseStatusNewConfigurationAccepted LiveInputNewResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type LiveInputNewResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                         `json:"url"`
	JSON liveInputNewResponseWebRtcJSON `json:"-"`
}

// liveInputNewResponseWebRtcJSON contains the JSON metadata for the struct
// [LiveInputNewResponseWebRtc]
type liveInputNewResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseWebRtcJSON) RawJSON() string {
	return r.raw
}

// Details for playback from a live input using WebRTC.
type LiveInputNewResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                 `json:"url"`
	JSON liveInputNewResponseWebRtcPlaybackJSON `json:"-"`
}

// liveInputNewResponseWebRtcPlaybackJSON contains the JSON metadata for the struct
// [LiveInputNewResponseWebRtcPlayback]
type liveInputNewResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseWebRtcPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details about a live input.
type LiveInputUpdateResponse struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording LiveInputUpdateResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps LiveInputUpdateResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback LiveInputUpdateResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt LiveInputUpdateResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback LiveInputUpdateResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status LiveInputUpdateResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc LiveInputUpdateResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback LiveInputUpdateResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           liveInputUpdateResponseJSON           `json:"-"`
}

// liveInputUpdateResponseJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponse]
type liveInputUpdateResponseJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Recording                apijson.Field
	Rtmps                    apijson.Field
	RtmpsPlayback            apijson.Field
	Srt                      apijson.Field
	SrtPlayback              apijson.Field
	Status                   apijson.Field
	Uid                      apijson.Field
	WebRtc                   apijson.Field
	WebRtcPlayback           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInputUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputUpdateResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode LiveInputUpdateResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                `json:"timeoutSeconds"`
	JSON           liveInputUpdateResponseRecordingJSON `json:"-"`
}

// liveInputUpdateResponseRecordingJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseRecording]
type liveInputUpdateResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LiveInputUpdateResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseRecordingJSON) RawJSON() string {
	return r.raw
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputUpdateResponseRecordingMode string

const (
	LiveInputUpdateResponseRecordingModeOff       LiveInputUpdateResponseRecordingMode = "off"
	LiveInputUpdateResponseRecordingModeAutomatic LiveInputUpdateResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type LiveInputUpdateResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                           `json:"url"`
	JSON liveInputUpdateResponseRtmpsJSON `json:"-"`
}

// liveInputUpdateResponseRtmpsJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseRtmps]
type liveInputUpdateResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseRtmpsJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using RTMPS.
type LiveInputUpdateResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                   `json:"url"`
	JSON liveInputUpdateResponseRtmpsPlaybackJSON `json:"-"`
}

// liveInputUpdateResponseRtmpsPlaybackJSON contains the JSON metadata for the
// struct [LiveInputUpdateResponseRtmpsPlayback]
type liveInputUpdateResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseRtmpsPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to a live input using SRT.
type LiveInputUpdateResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                         `json:"url"`
	JSON liveInputUpdateResponseSrtJSON `json:"-"`
}

// liveInputUpdateResponseSrtJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseSrt]
type liveInputUpdateResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseSrtJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using SRT.
type LiveInputUpdateResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                 `json:"url"`
	JSON liveInputUpdateResponseSrtPlaybackJSON `json:"-"`
}

// liveInputUpdateResponseSrtPlaybackJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseSrtPlayback]
type liveInputUpdateResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseSrtPlaybackJSON) RawJSON() string {
	return r.raw
}

// The connection status of a live input.
type LiveInputUpdateResponseStatus string

const (
	LiveInputUpdateResponseStatusConnected                LiveInputUpdateResponseStatus = "connected"
	LiveInputUpdateResponseStatusReconnected              LiveInputUpdateResponseStatus = "reconnected"
	LiveInputUpdateResponseStatusReconnecting             LiveInputUpdateResponseStatus = "reconnecting"
	LiveInputUpdateResponseStatusClientDisconnect         LiveInputUpdateResponseStatus = "client_disconnect"
	LiveInputUpdateResponseStatusTTLExceeded              LiveInputUpdateResponseStatus = "ttl_exceeded"
	LiveInputUpdateResponseStatusFailedToConnect          LiveInputUpdateResponseStatus = "failed_to_connect"
	LiveInputUpdateResponseStatusFailedToReconnect        LiveInputUpdateResponseStatus = "failed_to_reconnect"
	LiveInputUpdateResponseStatusNewConfigurationAccepted LiveInputUpdateResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type LiveInputUpdateResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                            `json:"url"`
	JSON liveInputUpdateResponseWebRtcJSON `json:"-"`
}

// liveInputUpdateResponseWebRtcJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseWebRtc]
type liveInputUpdateResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseWebRtcJSON) RawJSON() string {
	return r.raw
}

// Details for playback from a live input using WebRTC.
type LiveInputUpdateResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                    `json:"url"`
	JSON liveInputUpdateResponseWebRtcPlaybackJSON `json:"-"`
}

// liveInputUpdateResponseWebRtcPlaybackJSON contains the JSON metadata for the
// struct [LiveInputUpdateResponseWebRtcPlayback]
type liveInputUpdateResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseWebRtcPlaybackJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponse struct {
	LiveInputs []LiveInputListResponseLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                     `json:"total"`
	JSON  liveInputListResponseJSON `json:"-"`
}

// liveInputListResponseJSON contains the JSON metadata for the struct
// [LiveInputListResponse]
type liveInputListResponseJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseLiveInput struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A unique identifier for a live input.
	Uid  string                             `json:"uid"`
	JSON liveInputListResponseLiveInputJSON `json:"-"`
}

// liveInputListResponseLiveInputJSON contains the JSON metadata for the struct
// [LiveInputListResponseLiveInput]
type liveInputListResponseLiveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Uid                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInputListResponseLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseLiveInputJSON) RawJSON() string {
	return r.raw
}

// Details about a live input.
type LiveInputGetResponse struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording LiveInputGetResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps LiveInputGetResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback LiveInputGetResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt LiveInputGetResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback LiveInputGetResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status LiveInputGetResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc LiveInputGetResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback LiveInputGetResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           liveInputGetResponseJSON           `json:"-"`
}

// liveInputGetResponseJSON contains the JSON metadata for the struct
// [LiveInputGetResponse]
type liveInputGetResponseJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Recording                apijson.Field
	Rtmps                    apijson.Field
	RtmpsPlayback            apijson.Field
	Srt                      apijson.Field
	SrtPlayback              apijson.Field
	Status                   apijson.Field
	Uid                      apijson.Field
	WebRtc                   apijson.Field
	WebRtcPlayback           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInputGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseJSON) RawJSON() string {
	return r.raw
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputGetResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode LiveInputGetResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                             `json:"timeoutSeconds"`
	JSON           liveInputGetResponseRecordingJSON `json:"-"`
}

// liveInputGetResponseRecordingJSON contains the JSON metadata for the struct
// [LiveInputGetResponseRecording]
type liveInputGetResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LiveInputGetResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseRecordingJSON) RawJSON() string {
	return r.raw
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputGetResponseRecordingMode string

const (
	LiveInputGetResponseRecordingModeOff       LiveInputGetResponseRecordingMode = "off"
	LiveInputGetResponseRecordingModeAutomatic LiveInputGetResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type LiveInputGetResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                        `json:"url"`
	JSON liveInputGetResponseRtmpsJSON `json:"-"`
}

// liveInputGetResponseRtmpsJSON contains the JSON metadata for the struct
// [LiveInputGetResponseRtmps]
type liveInputGetResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseRtmpsJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using RTMPS.
type LiveInputGetResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                `json:"url"`
	JSON liveInputGetResponseRtmpsPlaybackJSON `json:"-"`
}

// liveInputGetResponseRtmpsPlaybackJSON contains the JSON metadata for the struct
// [LiveInputGetResponseRtmpsPlayback]
type liveInputGetResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseRtmpsPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to a live input using SRT.
type LiveInputGetResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                      `json:"url"`
	JSON liveInputGetResponseSrtJSON `json:"-"`
}

// liveInputGetResponseSrtJSON contains the JSON metadata for the struct
// [LiveInputGetResponseSrt]
type liveInputGetResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseSrtJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using SRT.
type LiveInputGetResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                              `json:"url"`
	JSON liveInputGetResponseSrtPlaybackJSON `json:"-"`
}

// liveInputGetResponseSrtPlaybackJSON contains the JSON metadata for the struct
// [LiveInputGetResponseSrtPlayback]
type liveInputGetResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseSrtPlaybackJSON) RawJSON() string {
	return r.raw
}

// The connection status of a live input.
type LiveInputGetResponseStatus string

const (
	LiveInputGetResponseStatusConnected                LiveInputGetResponseStatus = "connected"
	LiveInputGetResponseStatusReconnected              LiveInputGetResponseStatus = "reconnected"
	LiveInputGetResponseStatusReconnecting             LiveInputGetResponseStatus = "reconnecting"
	LiveInputGetResponseStatusClientDisconnect         LiveInputGetResponseStatus = "client_disconnect"
	LiveInputGetResponseStatusTTLExceeded              LiveInputGetResponseStatus = "ttl_exceeded"
	LiveInputGetResponseStatusFailedToConnect          LiveInputGetResponseStatus = "failed_to_connect"
	LiveInputGetResponseStatusFailedToReconnect        LiveInputGetResponseStatus = "failed_to_reconnect"
	LiveInputGetResponseStatusNewConfigurationAccepted LiveInputGetResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type LiveInputGetResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                         `json:"url"`
	JSON liveInputGetResponseWebRtcJSON `json:"-"`
}

// liveInputGetResponseWebRtcJSON contains the JSON metadata for the struct
// [LiveInputGetResponseWebRtc]
type liveInputGetResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseWebRtcJSON) RawJSON() string {
	return r.raw
}

// Details for playback from a live input using WebRTC.
type LiveInputGetResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                 `json:"url"`
	JSON liveInputGetResponseWebRtcPlaybackJSON `json:"-"`
}

// liveInputGetResponseWebRtcPlaybackJSON contains the JSON metadata for the struct
// [LiveInputGetResponseWebRtcPlayback]
type liveInputGetResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseWebRtcPlaybackJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays param.Field[float64] `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta param.Field[interface{}] `json:"meta"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording param.Field[LiveInputNewParamsRecording] `json:"recording"`
}

func (r LiveInputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputNewParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[LiveInputNewParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r LiveInputNewParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputNewParamsRecordingMode string

const (
	LiveInputNewParamsRecordingModeOff       LiveInputNewParamsRecordingMode = "off"
	LiveInputNewParamsRecordingModeAutomatic LiveInputNewParamsRecordingMode = "automatic"
)

type LiveInputNewResponseEnvelope struct {
	Errors   []LiveInputNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputNewResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result LiveInputNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputNewResponseEnvelopeJSON    `json:"-"`
}

// liveInputNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputNewResponseEnvelope]
type liveInputNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    liveInputNewResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LiveInputNewResponseEnvelopeErrors]
type liveInputNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    liveInputNewResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputNewResponseEnvelopeMessages]
type liveInputNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputNewResponseEnvelopeSuccess bool

const (
	LiveInputNewResponseEnvelopeSuccessTrue LiveInputNewResponseEnvelopeSuccess = true
)

type LiveInputUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays param.Field[float64] `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta param.Field[interface{}] `json:"meta"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording param.Field[LiveInputUpdateParamsRecording] `json:"recording"`
}

func (r LiveInputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputUpdateParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[LiveInputUpdateParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r LiveInputUpdateParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputUpdateParamsRecordingMode string

const (
	LiveInputUpdateParamsRecordingModeOff       LiveInputUpdateParamsRecordingMode = "off"
	LiveInputUpdateParamsRecordingModeAutomatic LiveInputUpdateParamsRecordingMode = "automatic"
)

type LiveInputUpdateResponseEnvelope struct {
	Errors   []LiveInputUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result LiveInputUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputUpdateResponseEnvelopeJSON    `json:"-"`
}

// liveInputUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseEnvelope]
type liveInputUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    liveInputUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputUpdateResponseEnvelopeErrors]
type liveInputUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    liveInputUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputUpdateResponseEnvelopeMessages]
type liveInputUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputUpdateResponseEnvelopeSuccess bool

const (
	LiveInputUpdateResponseEnvelopeSuccessTrue LiveInputUpdateResponseEnvelopeSuccess = true
)

type LiveInputListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
}

// URLQuery serializes [LiveInputListParams]'s query parameters as `url.Values`.
func (r LiveInputListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LiveInputListResponseEnvelope struct {
	Errors   []LiveInputListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputListResponseEnvelopeMessages `json:"messages,required"`
	Result   LiveInputListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputListResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputListResponseEnvelopeJSON    `json:"-"`
}

// liveInputListResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputListResponseEnvelope]
type liveInputListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    liveInputListResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputListResponseEnvelopeErrors]
type liveInputListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    liveInputListResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputListResponseEnvelopeMessages]
type liveInputListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputListResponseEnvelopeSuccess bool

const (
	LiveInputListResponseEnvelopeSuccessTrue LiveInputListResponseEnvelopeSuccess = true
)

type LiveInputDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputGetResponseEnvelope struct {
	Errors   []LiveInputGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result LiveInputGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputGetResponseEnvelopeJSON    `json:"-"`
}

// liveInputGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputGetResponseEnvelope]
type liveInputGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    liveInputGetResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LiveInputGetResponseEnvelopeErrors]
type liveInputGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    liveInputGetResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputGetResponseEnvelopeMessages]
type liveInputGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputGetResponseEnvelopeSuccess bool

const (
	LiveInputGetResponseEnvelopeSuccessTrue LiveInputGetResponseEnvelopeSuccess = true
)
