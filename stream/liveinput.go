// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *LiveInputService) New(ctx context.Context, params LiveInputNewParams, opts ...option.RequestOption) (res *StreamLiveInput, err error) {
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
func (r *LiveInputService) Update(ctx context.Context, liveInputIdentifier string, params LiveInputUpdateParams, opts ...option.RequestOption) (res *StreamLiveInput, err error) {
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
func (r *LiveInputService) Delete(ctx context.Context, liveInputIdentifier string, params LiveInputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", params.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves details of an existing live input.
func (r *LiveInputService) Get(ctx context.Context, liveInputIdentifier string, query LiveInputGetParams, opts ...option.RequestOption) (res *StreamLiveInput, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", query.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about a live input.
type StreamLiveInput struct {
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
	Recording StreamLiveInputRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps StreamLiveInputRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback StreamLiveInputRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt StreamLiveInputSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback StreamLiveInputSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status StreamLiveInputStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc StreamLiveInputWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback StreamLiveInputWebRtcPlayback `json:"webRTCPlayback"`
	JSON           streamLiveInputJSON           `json:"-"`
}

// streamLiveInputJSON contains the JSON metadata for the struct [StreamLiveInput]
type streamLiveInputJSON struct {
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

func (r *StreamLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputJSON) RawJSON() string {
	return r.raw
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode StreamLiveInputRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                        `json:"timeoutSeconds"`
	JSON           streamLiveInputRecordingJSON `json:"-"`
}

// streamLiveInputRecordingJSON contains the JSON metadata for the struct
// [StreamLiveInputRecording]
type streamLiveInputRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamLiveInputRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputRecordingJSON) RawJSON() string {
	return r.raw
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputRecordingMode string

const (
	StreamLiveInputRecordingModeOff       StreamLiveInputRecordingMode = "off"
	StreamLiveInputRecordingModeAutomatic StreamLiveInputRecordingMode = "automatic"
)

func (r StreamLiveInputRecordingMode) IsKnown() bool {
	switch r {
	case StreamLiveInputRecordingModeOff, StreamLiveInputRecordingModeAutomatic:
		return true
	}
	return false
}

// Details for streaming to an live input using RTMPS.
type StreamLiveInputRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                   `json:"url"`
	JSON streamLiveInputRtmpsJSON `json:"-"`
}

// streamLiveInputRtmpsJSON contains the JSON metadata for the struct
// [StreamLiveInputRtmps]
type streamLiveInputRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputRtmpsJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using RTMPS.
type StreamLiveInputRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                           `json:"url"`
	JSON streamLiveInputRtmpsPlaybackJSON `json:"-"`
}

// streamLiveInputRtmpsPlaybackJSON contains the JSON metadata for the struct
// [StreamLiveInputRtmpsPlayback]
type streamLiveInputRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputRtmpsPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to a live input using SRT.
type StreamLiveInputSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                 `json:"url"`
	JSON streamLiveInputSrtJSON `json:"-"`
}

// streamLiveInputSrtJSON contains the JSON metadata for the struct
// [StreamLiveInputSrt]
type streamLiveInputSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputSrtJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using SRT.
type StreamLiveInputSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                         `json:"url"`
	JSON streamLiveInputSrtPlaybackJSON `json:"-"`
}

// streamLiveInputSrtPlaybackJSON contains the JSON metadata for the struct
// [StreamLiveInputSrtPlayback]
type streamLiveInputSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputSrtPlaybackJSON) RawJSON() string {
	return r.raw
}

// The connection status of a live input.
type StreamLiveInputStatus string

const (
	StreamLiveInputStatusConnected                StreamLiveInputStatus = "connected"
	StreamLiveInputStatusReconnected              StreamLiveInputStatus = "reconnected"
	StreamLiveInputStatusReconnecting             StreamLiveInputStatus = "reconnecting"
	StreamLiveInputStatusClientDisconnect         StreamLiveInputStatus = "client_disconnect"
	StreamLiveInputStatusTTLExceeded              StreamLiveInputStatus = "ttl_exceeded"
	StreamLiveInputStatusFailedToConnect          StreamLiveInputStatus = "failed_to_connect"
	StreamLiveInputStatusFailedToReconnect        StreamLiveInputStatus = "failed_to_reconnect"
	StreamLiveInputStatusNewConfigurationAccepted StreamLiveInputStatus = "new_configuration_accepted"
)

func (r StreamLiveInputStatus) IsKnown() bool {
	switch r {
	case StreamLiveInputStatusConnected, StreamLiveInputStatusReconnected, StreamLiveInputStatusReconnecting, StreamLiveInputStatusClientDisconnect, StreamLiveInputStatusTTLExceeded, StreamLiveInputStatusFailedToConnect, StreamLiveInputStatusFailedToReconnect, StreamLiveInputStatusNewConfigurationAccepted:
		return true
	}
	return false
}

// Details for streaming to a live input using WebRTC.
type StreamLiveInputWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                    `json:"url"`
	JSON streamLiveInputWebRtcJSON `json:"-"`
}

// streamLiveInputWebRtcJSON contains the JSON metadata for the struct
// [StreamLiveInputWebRtc]
type streamLiveInputWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputWebRtcJSON) RawJSON() string {
	return r.raw
}

// Details for playback from a live input using WebRTC.
type StreamLiveInputWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                            `json:"url"`
	JSON streamLiveInputWebRtcPlaybackJSON `json:"-"`
}

// streamLiveInputWebRtcPlaybackJSON contains the JSON metadata for the struct
// [StreamLiveInputWebRtcPlayback]
type streamLiveInputWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamLiveInputWebRtcPlaybackJSON) RawJSON() string {
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

func (r LiveInputNewParamsRecordingMode) IsKnown() bool {
	switch r {
	case LiveInputNewParamsRecordingModeOff, LiveInputNewParamsRecordingModeAutomatic:
		return true
	}
	return false
}

type LiveInputNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInput `json:"result,required"`
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

// Whether the API call was successful
type LiveInputNewResponseEnvelopeSuccess bool

const (
	LiveInputNewResponseEnvelopeSuccessTrue LiveInputNewResponseEnvelopeSuccess = true
)

func (r LiveInputNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r LiveInputUpdateParamsRecordingMode) IsKnown() bool {
	switch r {
	case LiveInputUpdateParamsRecordingModeOff, LiveInputUpdateParamsRecordingModeAutomatic:
		return true
	}
	return false
}

type LiveInputUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInput `json:"result,required"`
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

// Whether the API call was successful
type LiveInputUpdateResponseEnvelopeSuccess bool

const (
	LiveInputUpdateResponseEnvelopeSuccessTrue LiveInputUpdateResponseEnvelopeSuccess = true
)

func (r LiveInputUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LiveInputListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LiveInputListResponse `json:"result,required"`
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

// Whether the API call was successful
type LiveInputListResponseEnvelopeSuccess bool

const (
	LiveInputListResponseEnvelopeSuccessTrue LiveInputListResponseEnvelopeSuccess = true
)

func (r LiveInputListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LiveInputDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r LiveInputDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LiveInputGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInput `json:"result,required"`
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

// Whether the API call was successful
type LiveInputGetResponseEnvelopeSuccess bool

const (
	LiveInputGetResponseEnvelopeSuccessTrue LiveInputGetResponseEnvelopeSuccess = true
)

func (r LiveInputGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
