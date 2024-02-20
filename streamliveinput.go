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
func (r *StreamLiveInputService) New(ctx context.Context, accountID string, body StreamLiveInputNewParams, opts ...option.RequestOption) (res *StreamLiveInputNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *StreamLiveInputService) List(ctx context.Context, accountID string, query StreamLiveInputListParams, opts ...option.RequestOption) (res *StreamLiveInputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

// Details about a live input.
type StreamLiveInputNewResponse struct {
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
	Recording StreamLiveInputNewResponseRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps StreamLiveInputNewResponseRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback StreamLiveInputNewResponseRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt StreamLiveInputNewResponseSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback StreamLiveInputNewResponseSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status StreamLiveInputNewResponseStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc StreamLiveInputNewResponseWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback StreamLiveInputNewResponseWebRtcPlayback `json:"webRTCPlayback"`
	JSON           streamLiveInputNewResponseJSON           `json:"-"`
}

// streamLiveInputNewResponseJSON contains the JSON metadata for the struct
// [StreamLiveInputNewResponse]
type streamLiveInputNewResponseJSON struct {
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

func (r *StreamLiveInputNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type StreamLiveInputNewResponseRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode StreamLiveInputNewResponseRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                   `json:"timeoutSeconds"`
	JSON           streamLiveInputNewResponseRecordingJSON `json:"-"`
}

// streamLiveInputNewResponseRecordingJSON contains the JSON metadata for the
// struct [StreamLiveInputNewResponseRecording]
type streamLiveInputNewResponseRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type StreamLiveInputNewResponseRecordingMode string

const (
	StreamLiveInputNewResponseRecordingModeOff       StreamLiveInputNewResponseRecordingMode = "off"
	StreamLiveInputNewResponseRecordingModeAutomatic StreamLiveInputNewResponseRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type StreamLiveInputNewResponseRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                              `json:"url"`
	JSON streamLiveInputNewResponseRtmpsJSON `json:"-"`
}

// streamLiveInputNewResponseRtmpsJSON contains the JSON metadata for the struct
// [StreamLiveInputNewResponseRtmps]
type streamLiveInputNewResponseRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type StreamLiveInputNewResponseRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                      `json:"url"`
	JSON streamLiveInputNewResponseRtmpsPlaybackJSON `json:"-"`
}

// streamLiveInputNewResponseRtmpsPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputNewResponseRtmpsPlayback]
type streamLiveInputNewResponseRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type StreamLiveInputNewResponseSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                            `json:"url"`
	JSON streamLiveInputNewResponseSrtJSON `json:"-"`
}

// streamLiveInputNewResponseSrtJSON contains the JSON metadata for the struct
// [StreamLiveInputNewResponseSrt]
type streamLiveInputNewResponseSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type StreamLiveInputNewResponseSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                    `json:"url"`
	JSON streamLiveInputNewResponseSrtPlaybackJSON `json:"-"`
}

// streamLiveInputNewResponseSrtPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputNewResponseSrtPlayback]
type streamLiveInputNewResponseSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type StreamLiveInputNewResponseStatus string

const (
	StreamLiveInputNewResponseStatusConnected                StreamLiveInputNewResponseStatus = "connected"
	StreamLiveInputNewResponseStatusReconnected              StreamLiveInputNewResponseStatus = "reconnected"
	StreamLiveInputNewResponseStatusReconnecting             StreamLiveInputNewResponseStatus = "reconnecting"
	StreamLiveInputNewResponseStatusClientDisconnect         StreamLiveInputNewResponseStatus = "client_disconnect"
	StreamLiveInputNewResponseStatusTTLExceeded              StreamLiveInputNewResponseStatus = "ttl_exceeded"
	StreamLiveInputNewResponseStatusFailedToConnect          StreamLiveInputNewResponseStatus = "failed_to_connect"
	StreamLiveInputNewResponseStatusFailedToReconnect        StreamLiveInputNewResponseStatus = "failed_to_reconnect"
	StreamLiveInputNewResponseStatusNewConfigurationAccepted StreamLiveInputNewResponseStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type StreamLiveInputNewResponseWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                               `json:"url"`
	JSON streamLiveInputNewResponseWebRtcJSON `json:"-"`
}

// streamLiveInputNewResponseWebRtcJSON contains the JSON metadata for the struct
// [StreamLiveInputNewResponseWebRtc]
type streamLiveInputNewResponseWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type StreamLiveInputNewResponseWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                       `json:"url"`
	JSON streamLiveInputNewResponseWebRtcPlaybackJSON `json:"-"`
}

// streamLiveInputNewResponseWebRtcPlaybackJSON contains the JSON metadata for the
// struct [StreamLiveInputNewResponseWebRtcPlayback]
type streamLiveInputNewResponseWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputNewResponseWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type StreamLiveInputNewParams struct {
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
	Result StreamLiveInputNewResponse `json:"result,required"`
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

type StreamLiveInputListParams struct {
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
