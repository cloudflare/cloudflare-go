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

// Updates a specified live input.
func (r *StreamLiveInputService) Update(ctx context.Context, accountID string, liveInputIdentifier string, body StreamLiveInputUpdateParams, opts ...option.RequestOption) (res *StreamLiveInputUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Prevents a live input from being streamed to and makes the live input
// inaccessible to any future API calls.
func (r *StreamLiveInputService) Delete(ctx context.Context, accountID string, liveInputIdentifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves details of an existing live input.
func (r *StreamLiveInputService) Get(ctx context.Context, accountID string, liveInputIdentifier string, opts ...option.RequestOption) (res *StreamLiveInputGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a live input, and returns credentials that you or your users can use to
// stream live video to Cloudflare Stream.
func (r *StreamLiveInputService) StreamLiveInputsNewALiveInput(ctx context.Context, accountID string, body StreamLiveInputStreamLiveInputsNewALiveInputParams, opts ...option.RequestOption) (res *StreamLiveInputStreamLiveInputsNewALiveInputResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *StreamLiveInputService) StreamLiveInputsListLiveInputs(ctx context.Context, accountID string, query StreamLiveInputStreamLiveInputsListLiveInputsParams, opts ...option.RequestOption) (res *StreamLiveInputStreamLiveInputsListLiveInputsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about a live input.
type StreamLiveInputUpdateResponse struct {
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
	Recording StreamLiveInputUpdateResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps StreamLiveInputUpdateResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback StreamLiveInputUpdateResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt StreamLiveInputUpdateResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback StreamLiveInputUpdateResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status StreamLiveInputUpdateResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc StreamLiveInputUpdateResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback StreamLiveInputUpdateResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           streamLiveInputUpdateResponseJSON           `json:"-"`
}

// streamLiveInputUpdateResponseJSON contains the JSON metadata for the struct
// [StreamLiveInputUpdateResponse]
type streamLiveInputUpdateResponseJSON struct {
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

func (r *StreamLiveInputUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputUpdateResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode StreamLiveInputUpdateResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                      `json:"timeoutSeconds"`
	JSON           streamLiveInputUpdateResponseRecordingJSON `json:"-"`
}

// streamLiveInputUpdateResponseRecordingJSON contains the JSON metadata for the
// struct [StreamLiveInputUpdateResponseRecording]
type streamLiveInputUpdateResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputUpdateResponseRecordingMode string

const (
	StreamLiveInputUpdateResponseRecordingModeOff       StreamLiveInputUpdateResponseRecordingMode = "off"
	StreamLiveInputUpdateResponseRecordingModeAutomatic StreamLiveInputUpdateResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type StreamLiveInputUpdateResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                 `json:"url"`
	JSON streamLiveInputUpdateResponseRtmpsJSON `json:"-"`
}

// streamLiveInputUpdateResponseRtmpsJSON contains the JSON metadata for the struct
// [StreamLiveInputUpdateResponseRtmps]
type streamLiveInputUpdateResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type StreamLiveInputUpdateResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                         `json:"url"`
	JSON streamLiveInputUpdateResponseRtmpsPlaybackJSON `json:"-"`
}

// streamLiveInputUpdateResponseRtmpsPlaybackJSON contains the JSON metadata for
// the struct [StreamLiveInputUpdateResponseRtmpsPlayback]
type streamLiveInputUpdateResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type StreamLiveInputUpdateResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                               `json:"url"`
	JSON streamLiveInputUpdateResponseSrtJSON `json:"-"`
}

// streamLiveInputUpdateResponseSrtJSON contains the JSON metadata for the struct
// [StreamLiveInputUpdateResponseSrt]
type streamLiveInputUpdateResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type StreamLiveInputUpdateResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                       `json:"url"`
	JSON streamLiveInputUpdateResponseSrtPlaybackJSON `json:"-"`
}

// streamLiveInputUpdateResponseSrtPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputUpdateResponseSrtPlayback]
type streamLiveInputUpdateResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type StreamLiveInputUpdateResponseStatus string

const (
	StreamLiveInputUpdateResponseStatusConnected                StreamLiveInputUpdateResponseStatus = "connected"
	StreamLiveInputUpdateResponseStatusReconnected              StreamLiveInputUpdateResponseStatus = "reconnected"
	StreamLiveInputUpdateResponseStatusReconnecting             StreamLiveInputUpdateResponseStatus = "reconnecting"
	StreamLiveInputUpdateResponseStatusClientDisconnect         StreamLiveInputUpdateResponseStatus = "client_disconnect"
	StreamLiveInputUpdateResponseStatusTTLExceeded              StreamLiveInputUpdateResponseStatus = "ttl_exceeded"
	StreamLiveInputUpdateResponseStatusFailedToConnect          StreamLiveInputUpdateResponseStatus = "failed_to_connect"
	StreamLiveInputUpdateResponseStatusFailedToReconnect        StreamLiveInputUpdateResponseStatus = "failed_to_reconnect"
	StreamLiveInputUpdateResponseStatusNewConfigurationAccepted StreamLiveInputUpdateResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type StreamLiveInputUpdateResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                  `json:"url"`
	JSON streamLiveInputUpdateResponseWebRtcJSON `json:"-"`
}

// streamLiveInputUpdateResponseWebRtcJSON contains the JSON metadata for the
// struct [StreamLiveInputUpdateResponseWebRtc]
type streamLiveInputUpdateResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type StreamLiveInputUpdateResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                          `json:"url"`
	JSON streamLiveInputUpdateResponseWebRtcPlaybackJSON `json:"-"`
}

// streamLiveInputUpdateResponseWebRtcPlaybackJSON contains the JSON metadata for
// the struct [StreamLiveInputUpdateResponseWebRtcPlayback]
type streamLiveInputUpdateResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputUpdateResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a live input.
type StreamLiveInputGetResponse struct {
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
	Recording StreamLiveInputGetResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps StreamLiveInputGetResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback StreamLiveInputGetResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt StreamLiveInputGetResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback StreamLiveInputGetResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status StreamLiveInputGetResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc StreamLiveInputGetResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback StreamLiveInputGetResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           streamLiveInputGetResponseJSON           `json:"-"`
}

// streamLiveInputGetResponseJSON contains the JSON metadata for the struct
// [StreamLiveInputGetResponse]
type streamLiveInputGetResponseJSON struct {
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

func (r *StreamLiveInputGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputGetResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode StreamLiveInputGetResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                   `json:"timeoutSeconds"`
	JSON           streamLiveInputGetResponseRecordingJSON `json:"-"`
}

// streamLiveInputGetResponseRecordingJSON contains the JSON metadata for the
// struct [StreamLiveInputGetResponseRecording]
type streamLiveInputGetResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputGetResponseRecordingMode string

const (
	StreamLiveInputGetResponseRecordingModeOff       StreamLiveInputGetResponseRecordingMode = "off"
	StreamLiveInputGetResponseRecordingModeAutomatic StreamLiveInputGetResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type StreamLiveInputGetResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                              `json:"url"`
	JSON streamLiveInputGetResponseRtmpsJSON `json:"-"`
}

// streamLiveInputGetResponseRtmpsJSON contains the JSON metadata for the struct
// [StreamLiveInputGetResponseRtmps]
type streamLiveInputGetResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type StreamLiveInputGetResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                      `json:"url"`
	JSON streamLiveInputGetResponseRtmpsPlaybackJSON `json:"-"`
}

// streamLiveInputGetResponseRtmpsPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputGetResponseRtmpsPlayback]
type streamLiveInputGetResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type StreamLiveInputGetResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                            `json:"url"`
	JSON streamLiveInputGetResponseSrtJSON `json:"-"`
}

// streamLiveInputGetResponseSrtJSON contains the JSON metadata for the struct
// [StreamLiveInputGetResponseSrt]
type streamLiveInputGetResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type StreamLiveInputGetResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                    `json:"url"`
	JSON streamLiveInputGetResponseSrtPlaybackJSON `json:"-"`
}

// streamLiveInputGetResponseSrtPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputGetResponseSrtPlayback]
type streamLiveInputGetResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type StreamLiveInputGetResponseStatus string

const (
	StreamLiveInputGetResponseStatusConnected                StreamLiveInputGetResponseStatus = "connected"
	StreamLiveInputGetResponseStatusReconnected              StreamLiveInputGetResponseStatus = "reconnected"
	StreamLiveInputGetResponseStatusReconnecting             StreamLiveInputGetResponseStatus = "reconnecting"
	StreamLiveInputGetResponseStatusClientDisconnect         StreamLiveInputGetResponseStatus = "client_disconnect"
	StreamLiveInputGetResponseStatusTTLExceeded              StreamLiveInputGetResponseStatus = "ttl_exceeded"
	StreamLiveInputGetResponseStatusFailedToConnect          StreamLiveInputGetResponseStatus = "failed_to_connect"
	StreamLiveInputGetResponseStatusFailedToReconnect        StreamLiveInputGetResponseStatus = "failed_to_reconnect"
	StreamLiveInputGetResponseStatusNewConfigurationAccepted StreamLiveInputGetResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type StreamLiveInputGetResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                               `json:"url"`
	JSON streamLiveInputGetResponseWebRtcJSON `json:"-"`
}

// streamLiveInputGetResponseWebRtcJSON contains the JSON metadata for the struct
// [StreamLiveInputGetResponseWebRtc]
type streamLiveInputGetResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type StreamLiveInputGetResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                       `json:"url"`
	JSON streamLiveInputGetResponseWebRtcPlaybackJSON `json:"-"`
}

// streamLiveInputGetResponseWebRtcPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputGetResponseWebRtcPlayback]
type streamLiveInputGetResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputGetResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a live input.
type StreamLiveInputStreamLiveInputsNewALiveInputResponse struct {
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
	Recording StreamLiveInputStreamLiveInputsNewALiveInputResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt StreamLiveInputStreamLiveInputsNewALiveInputResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback StreamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           streamLiveInputStreamLiveInputsNewALiveInputResponseJSON           `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseJSON contains the JSON
// metadata for the struct [StreamLiveInputStreamLiveInputsNewALiveInputResponse]
type streamLiveInputStreamLiveInputsNewALiveInputResponseJSON struct {
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

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode StreamLiveInputStreamLiveInputsNewALiveInputResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                                             `json:"timeoutSeconds"`
	JSON           streamLiveInputStreamLiveInputsNewALiveInputResponseRecordingJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseRecordingJSON contains the
// JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseRecording]
type streamLiveInputStreamLiveInputsNewALiveInputResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseRecordingMode string

const (
	StreamLiveInputStreamLiveInputsNewALiveInputResponseRecordingModeOff       StreamLiveInputStreamLiveInputsNewALiveInputResponseRecordingMode = "off"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseRecordingModeAutomatic StreamLiveInputStreamLiveInputsNewALiveInputResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                                        `json:"url"`
	JSON streamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsJSON contains the JSON
// metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmps]
type streamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                                                `json:"url"`
	JSON streamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlaybackJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlaybackJSON contains
// the JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlayback]
type streamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                                                      `json:"url"`
	JSON streamLiveInputStreamLiveInputsNewALiveInputResponseSrtJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseSrtJSON contains the JSON
// metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseSrt]
type streamLiveInputStreamLiveInputsNewALiveInputResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                                              `json:"url"`
	JSON streamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlaybackJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlaybackJSON contains the
// JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlayback]
type streamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus string

const (
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusConnected                StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "connected"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusReconnected              StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "reconnected"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusReconnecting             StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "reconnecting"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusClientDisconnect         StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "client_disconnect"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusTTLExceeded              StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "ttl_exceeded"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusFailedToConnect          StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "failed_to_connect"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusFailedToReconnect        StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "failed_to_reconnect"
	StreamLiveInputStreamLiveInputsNewALiveInputResponseStatusNewConfigurationAccepted StreamLiveInputStreamLiveInputsNewALiveInputResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                                         `json:"url"`
	JSON streamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcJSON contains the JSON
// metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtc]
type streamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                                                 `json:"url"`
	JSON streamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlaybackJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlaybackJSON contains
// the JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlayback]
type streamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputStreamLiveInputsListLiveInputsResponse struct {
	LiveInputs []StreamLiveInputStreamLiveInputsListLiveInputsResponseLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                                                     `json:"total"`
	JSON  streamLiveInputStreamLiveInputsListLiveInputsResponseJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsListLiveInputsResponseJSON contains the JSON
// metadata for the struct [StreamLiveInputStreamLiveInputsListLiveInputsResponse]
type streamLiveInputStreamLiveInputsListLiveInputsResponseJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsListLiveInputsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputStreamLiveInputsListLiveInputsResponseLiveInput struct {
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
	Uid  string                                                             `json:"uid"`
	JSON streamLiveInputStreamLiveInputsListLiveInputsResponseLiveInputJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsListLiveInputsResponseLiveInputJSON contains the
// JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsListLiveInputsResponseLiveInput]
type streamLiveInputStreamLiveInputsListLiveInputsResponseLiveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Uid                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsListLiveInputsResponseLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputUpdateParams struct {
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
	Result StreamLiveInputUpdateResponse `json:"result,required"`
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

type StreamLiveInputGetResponseEnvelope struct {
	Errors   []StreamLiveInputGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInputGetResponse `json:"result,required"`
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

type StreamLiveInputStreamLiveInputsNewALiveInputParams struct {
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
	Recording param.Field[StreamLiveInputStreamLiveInputsNewALiveInputParamsRecording] `json:"recording"`
}

func (r StreamLiveInputStreamLiveInputsNewALiveInputParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputStreamLiveInputsNewALiveInputParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[StreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r StreamLiveInputStreamLiveInputsNewALiveInputParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode string

const (
	StreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingModeOff       StreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode = "off"
	StreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingModeAutomatic StreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode = "automatic"
)

type StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelope struct {
	Errors   []StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessages `json:"messages,required"`
	// Details about a live input.
	Result StreamLiveInputStreamLiveInputsNewALiveInputResponse `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelope]
type streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrors]
type streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessages]
type streamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeSuccess bool

const (
	StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeSuccessTrue StreamLiveInputStreamLiveInputsNewALiveInputResponseEnvelopeSuccess = true
)

type StreamLiveInputStreamLiveInputsListLiveInputsParams struct {
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
}

// URLQuery serializes [StreamLiveInputStreamLiveInputsListLiveInputsParams]'s
// query parameters as `url.Values`.
func (r StreamLiveInputStreamLiveInputsListLiveInputsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelope struct {
	Errors   []StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamLiveInputStreamLiveInputsListLiveInputsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelope]
type streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrors]
type streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessages]
type streamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeSuccess bool

const (
	StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeSuccessTrue StreamLiveInputStreamLiveInputsListLiveInputsResponseEnvelopeSuccess = true
)
