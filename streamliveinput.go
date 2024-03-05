// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// StreamLiveInputService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamLiveInputService] method
// instead.
type StreamLiveInputService struct {
	Options []option.RequestOption
	Outputs *StreamLiveInputOutputService
}

// NewStreamLiveInputService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamLiveInputService(opts ...option.RequestOption) (r *StreamLiveInputService) {
	r = &StreamLiveInputService{}
	r.Options = opts
	r.Outputs = NewStreamLiveInputOutputService(opts...)
	return
}

// Creates a live input, and returns credentials that you or your users can use to
// stream live video to Cloudflare Stream.
func (r *StreamLiveInputService) New(ctx context.Context, params StreamLiveInputNewParams, opts ...option.RequestOption) (res *StreamLiveInput, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specified live input.
func (r *StreamLiveInputService) Update(ctx context.Context, liveInputIdentifier string, params StreamLiveInputUpdateParams, opts ...option.RequestOption) (res *StreamLiveInput, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputUpdateResponseEnvelope
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
func (r *StreamLiveInputService) List(ctx context.Context, params StreamLiveInputListParams, opts ...option.RequestOption) (res *StreamLiveInputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputListResponseEnvelope
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
func (r *StreamLiveInputService) Delete(ctx context.Context, liveInputIdentifier string, body StreamLiveInputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", body.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Retrieves details of an existing live input.
func (r *StreamLiveInputService) Get(ctx context.Context, liveInputIdentifier string, query StreamLiveInputGetParams, opts ...option.RequestOption) (res *StreamLiveInput, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", query.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputRecordingMode string

const (
	StreamLiveInputRecordingModeOff       StreamLiveInputRecordingMode = "off"
	StreamLiveInputRecordingModeAutomatic StreamLiveInputRecordingMode = "automatic"
)

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

type StreamLiveInputListResponse struct {
	LiveInputs []StreamLiveInputListResponseLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                           `json:"total"`
	JSON  streamLiveInputListResponseJSON `json:"-"`
}

// streamLiveInputListResponseJSON contains the JSON metadata for the struct
// [StreamLiveInputListResponse]
type streamLiveInputListResponseJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputListResponseLiveInput struct {
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
	Uid  string                                   `json:"uid"`
	JSON streamLiveInputListResponseLiveInputJSON `json:"-"`
}

// streamLiveInputListResponseLiveInputJSON contains the JSON metadata for the
// struct [StreamLiveInputListResponseLiveInput]
type streamLiveInputListResponseLiveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Uid                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *StreamLiveInputListResponseLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputNewParams struct {
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
	Recording param.Field[StreamLiveInputNewParamsRecording] `json:"recording"`
}

func (r StreamLiveInputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputNewParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[StreamLiveInputNewParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r StreamLiveInputNewParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputNewParamsRecordingMode string

const (
	StreamLiveInputNewParamsRecordingModeOff       StreamLiveInputNewParamsRecordingMode = "off"
	StreamLiveInputNewParamsRecordingModeAutomatic StreamLiveInputNewParamsRecordingMode = "automatic"
)

type StreamLiveInputNewResponseEnvelope struct {
	Errors   []StreamLiveInputNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputNewResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInput `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputNewResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamLiveInputNewResponseEnvelope]
type streamLiveInputNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamLiveInputNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamLiveInputNewResponseEnvelopeErrors]
type streamLiveInputNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamLiveInputNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamLiveInputNewResponseEnvelopeMessages]
type streamLiveInputNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputNewResponseEnvelopeSuccess bool

const (
	StreamLiveInputNewResponseEnvelopeSuccessTrue StreamLiveInputNewResponseEnvelopeSuccess = true
)

type StreamLiveInputUpdateParams struct {
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
	Recording param.Field[StreamLiveInputUpdateParamsRecording] `json:"recording"`
}

func (r StreamLiveInputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputUpdateParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[StreamLiveInputUpdateParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r StreamLiveInputUpdateParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputUpdateParamsRecordingMode string

const (
	StreamLiveInputUpdateParamsRecordingModeOff       StreamLiveInputUpdateParamsRecordingMode = "off"
	StreamLiveInputUpdateParamsRecordingModeAutomatic StreamLiveInputUpdateParamsRecordingMode = "automatic"
)

type StreamLiveInputUpdateResponseEnvelope struct {
	Errors   []StreamLiveInputUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInput `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputUpdateResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamLiveInputUpdateResponseEnvelope]
type streamLiveInputUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamLiveInputUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamLiveInputUpdateResponseEnvelopeErrors]
type streamLiveInputUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    streamLiveInputUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamLiveInputUpdateResponseEnvelopeMessages]
type streamLiveInputUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputUpdateResponseEnvelopeSuccess bool

const (
	StreamLiveInputUpdateResponseEnvelopeSuccessTrue StreamLiveInputUpdateResponseEnvelopeSuccess = true
)

type StreamLiveInputListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
}

// URLQuery serializes [StreamLiveInputListParams]'s query parameters as
// `url.Values`.
func (r StreamLiveInputListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StreamLiveInputListResponseEnvelope struct {
	Errors   []StreamLiveInputListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputListResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamLiveInputListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputListResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputListResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamLiveInputListResponseEnvelope]
type streamLiveInputListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamLiveInputListResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamLiveInputListResponseEnvelopeErrors]
type streamLiveInputListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamLiveInputListResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamLiveInputListResponseEnvelopeMessages]
type streamLiveInputListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputListResponseEnvelopeSuccess bool

const (
	StreamLiveInputListResponseEnvelopeSuccessTrue StreamLiveInputListResponseEnvelopeSuccess = true
)

type StreamLiveInputDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamLiveInputGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamLiveInputGetResponseEnvelope struct {
	Errors   []StreamLiveInputGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInput `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputGetResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamLiveInputGetResponseEnvelope]
type streamLiveInputGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamLiveInputGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamLiveInputGetResponseEnvelopeErrors]
type streamLiveInputGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamLiveInputGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamLiveInputGetResponseEnvelopeMessages]
type streamLiveInputGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputGetResponseEnvelopeSuccess bool

const (
	StreamLiveInputGetResponseEnvelopeSuccessTrue StreamLiveInputGetResponseEnvelopeSuccess = true
)
