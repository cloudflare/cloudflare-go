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

// AccountStreamLiveInputService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamLiveInputService]
// method instead.
type AccountStreamLiveInputService struct {
	Options []option.RequestOption
	Outputs *AccountStreamLiveInputOutputService
}

// NewAccountStreamLiveInputService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamLiveInputService(opts ...option.RequestOption) (r *AccountStreamLiveInputService) {
	r = &AccountStreamLiveInputService{}
	r.Options = opts
	r.Outputs = NewAccountStreamLiveInputOutputService(opts...)
	return
}

// Retrieves details of an existing live input.
func (r *AccountStreamLiveInputService) Get(ctx context.Context, accountIdentifier string, liveInputIdentifier string, opts ...option.RequestOption) (res *AccountStreamLiveInputGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specified live input.
func (r *AccountStreamLiveInputService) Update(ctx context.Context, accountIdentifier string, liveInputIdentifier string, body AccountStreamLiveInputUpdateParams, opts ...option.RequestOption) (res *AccountStreamLiveInputUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Prevents a live input from being streamed to and makes the live input
// inaccessible to any future API calls.
func (r *AccountStreamLiveInputService) Delete(ctx context.Context, accountIdentifier string, liveInputIdentifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a live input, and returns credentials that you or your users can use to
// stream live video to Cloudflare Stream.
func (r *AccountStreamLiveInputService) StreamLiveInputsNewALiveInput(ctx context.Context, accountIdentifier string, body AccountStreamLiveInputStreamLiveInputsNewALiveInputParams, opts ...option.RequestOption) (res *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *AccountStreamLiveInputService) StreamLiveInputsListLiveInputs(ctx context.Context, accountIdentifier string, query AccountStreamLiveInputStreamLiveInputsListLiveInputsParams, opts ...option.RequestOption) (res *AccountStreamLiveInputStreamLiveInputsListLiveInputsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountStreamLiveInputGetResponse struct {
	Errors   []AccountStreamLiveInputGetResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputGetResponseMessage `json:"messages"`
	// Details about a live input.
	Result AccountStreamLiveInputGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputGetResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputGetResponseJSON    `json:"-"`
}

// accountStreamLiveInputGetResponseJSON contains the JSON metadata for the struct
// [AccountStreamLiveInputGetResponse]
type accountStreamLiveInputGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountStreamLiveInputGetResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputGetResponseError]
type accountStreamLiveInputGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountStreamLiveInputGetResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputGetResponseMessage]
type accountStreamLiveInputGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a live input.
type AccountStreamLiveInputGetResponseResult struct {
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
	Recording AccountStreamLiveInputGetResponseResultRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps AccountStreamLiveInputGetResponseResultRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback AccountStreamLiveInputGetResponseResultRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt AccountStreamLiveInputGetResponseResultSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback AccountStreamLiveInputGetResponseResultSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status AccountStreamLiveInputGetResponseResultStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc AccountStreamLiveInputGetResponseResultWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback AccountStreamLiveInputGetResponseResultWebRtcPlayback `json:"webRTCPlayback"`
	JSON           accountStreamLiveInputGetResponseResultJSON           `json:"-"`
}

// accountStreamLiveInputGetResponseResultJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputGetResponseResult]
type accountStreamLiveInputGetResponseResultJSON struct {
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

func (r *AccountStreamLiveInputGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type AccountStreamLiveInputGetResponseResultRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode AccountStreamLiveInputGetResponseResultRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                                `json:"timeoutSeconds"`
	JSON           accountStreamLiveInputGetResponseResultRecordingJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultRecordingJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputGetResponseResultRecording]
type accountStreamLiveInputGetResponseResultRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type AccountStreamLiveInputGetResponseResultRecordingMode string

const (
	AccountStreamLiveInputGetResponseResultRecordingModeOff       AccountStreamLiveInputGetResponseResultRecordingMode = "off"
	AccountStreamLiveInputGetResponseResultRecordingModeAutomatic AccountStreamLiveInputGetResponseResultRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type AccountStreamLiveInputGetResponseResultRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                           `json:"url"`
	JSON accountStreamLiveInputGetResponseResultRtmpsJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultRtmpsJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputGetResponseResultRtmps]
type accountStreamLiveInputGetResponseResultRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type AccountStreamLiveInputGetResponseResultRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                                   `json:"url"`
	JSON accountStreamLiveInputGetResponseResultRtmpsPlaybackJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultRtmpsPlaybackJSON contains the JSON
// metadata for the struct [AccountStreamLiveInputGetResponseResultRtmpsPlayback]
type accountStreamLiveInputGetResponseResultRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type AccountStreamLiveInputGetResponseResultSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                                         `json:"url"`
	JSON accountStreamLiveInputGetResponseResultSrtJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultSrtJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputGetResponseResultSrt]
type accountStreamLiveInputGetResponseResultSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type AccountStreamLiveInputGetResponseResultSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                                 `json:"url"`
	JSON accountStreamLiveInputGetResponseResultSrtPlaybackJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultSrtPlaybackJSON contains the JSON
// metadata for the struct [AccountStreamLiveInputGetResponseResultSrtPlayback]
type accountStreamLiveInputGetResponseResultSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type AccountStreamLiveInputGetResponseResultStatus string

const (
	AccountStreamLiveInputGetResponseResultStatusConnected                AccountStreamLiveInputGetResponseResultStatus = "connected"
	AccountStreamLiveInputGetResponseResultStatusReconnected              AccountStreamLiveInputGetResponseResultStatus = "reconnected"
	AccountStreamLiveInputGetResponseResultStatusReconnecting             AccountStreamLiveInputGetResponseResultStatus = "reconnecting"
	AccountStreamLiveInputGetResponseResultStatusClientDisconnect         AccountStreamLiveInputGetResponseResultStatus = "client_disconnect"
	AccountStreamLiveInputGetResponseResultStatusTtlExceeded              AccountStreamLiveInputGetResponseResultStatus = "ttl_exceeded"
	AccountStreamLiveInputGetResponseResultStatusFailedToConnect          AccountStreamLiveInputGetResponseResultStatus = "failed_to_connect"
	AccountStreamLiveInputGetResponseResultStatusFailedToReconnect        AccountStreamLiveInputGetResponseResultStatus = "failed_to_reconnect"
	AccountStreamLiveInputGetResponseResultStatusNewConfigurationAccepted AccountStreamLiveInputGetResponseResultStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type AccountStreamLiveInputGetResponseResultWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                            `json:"url"`
	JSON accountStreamLiveInputGetResponseResultWebRtcJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultWebRtcJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputGetResponseResultWebRtc]
type accountStreamLiveInputGetResponseResultWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type AccountStreamLiveInputGetResponseResultWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                                    `json:"url"`
	JSON accountStreamLiveInputGetResponseResultWebRtcPlaybackJSON `json:"-"`
}

// accountStreamLiveInputGetResponseResultWebRtcPlaybackJSON contains the JSON
// metadata for the struct [AccountStreamLiveInputGetResponseResultWebRtcPlayback]
type accountStreamLiveInputGetResponseResultWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputGetResponseResultWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputGetResponseSuccess bool

const (
	AccountStreamLiveInputGetResponseSuccessTrue AccountStreamLiveInputGetResponseSuccess = true
)

type AccountStreamLiveInputUpdateResponse struct {
	Errors   []AccountStreamLiveInputUpdateResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputUpdateResponseMessage `json:"messages"`
	// Details about a live input.
	Result AccountStreamLiveInputUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputUpdateResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputUpdateResponseJSON    `json:"-"`
}

// accountStreamLiveInputUpdateResponseJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputUpdateResponse]
type accountStreamLiveInputUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountStreamLiveInputUpdateResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputUpdateResponseError]
type accountStreamLiveInputUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountStreamLiveInputUpdateResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputUpdateResponseMessage]
type accountStreamLiveInputUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a live input.
type AccountStreamLiveInputUpdateResponseResult struct {
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
	Recording AccountStreamLiveInputUpdateResponseResultRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps AccountStreamLiveInputUpdateResponseResultRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback AccountStreamLiveInputUpdateResponseResultRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt AccountStreamLiveInputUpdateResponseResultSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback AccountStreamLiveInputUpdateResponseResultSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status AccountStreamLiveInputUpdateResponseResultStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc AccountStreamLiveInputUpdateResponseResultWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback AccountStreamLiveInputUpdateResponseResultWebRtcPlayback `json:"webRTCPlayback"`
	JSON           accountStreamLiveInputUpdateResponseResultJSON           `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputUpdateResponseResult]
type accountStreamLiveInputUpdateResponseResultJSON struct {
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

func (r *AccountStreamLiveInputUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type AccountStreamLiveInputUpdateResponseResultRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode AccountStreamLiveInputUpdateResponseResultRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                                   `json:"timeoutSeconds"`
	JSON           accountStreamLiveInputUpdateResponseResultRecordingJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultRecordingJSON contains the JSON
// metadata for the struct [AccountStreamLiveInputUpdateResponseResultRecording]
type accountStreamLiveInputUpdateResponseResultRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type AccountStreamLiveInputUpdateResponseResultRecordingMode string

const (
	AccountStreamLiveInputUpdateResponseResultRecordingModeOff       AccountStreamLiveInputUpdateResponseResultRecordingMode = "off"
	AccountStreamLiveInputUpdateResponseResultRecordingModeAutomatic AccountStreamLiveInputUpdateResponseResultRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type AccountStreamLiveInputUpdateResponseResultRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                              `json:"url"`
	JSON accountStreamLiveInputUpdateResponseResultRtmpsJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultRtmpsJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputUpdateResponseResultRtmps]
type accountStreamLiveInputUpdateResponseResultRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type AccountStreamLiveInputUpdateResponseResultRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                                      `json:"url"`
	JSON accountStreamLiveInputUpdateResponseResultRtmpsPlaybackJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultRtmpsPlaybackJSON contains the JSON
// metadata for the struct
// [AccountStreamLiveInputUpdateResponseResultRtmpsPlayback]
type accountStreamLiveInputUpdateResponseResultRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type AccountStreamLiveInputUpdateResponseResultSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                                            `json:"url"`
	JSON accountStreamLiveInputUpdateResponseResultSrtJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultSrtJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputUpdateResponseResultSrt]
type accountStreamLiveInputUpdateResponseResultSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type AccountStreamLiveInputUpdateResponseResultSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                                    `json:"url"`
	JSON accountStreamLiveInputUpdateResponseResultSrtPlaybackJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultSrtPlaybackJSON contains the JSON
// metadata for the struct [AccountStreamLiveInputUpdateResponseResultSrtPlayback]
type accountStreamLiveInputUpdateResponseResultSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type AccountStreamLiveInputUpdateResponseResultStatus string

const (
	AccountStreamLiveInputUpdateResponseResultStatusConnected                AccountStreamLiveInputUpdateResponseResultStatus = "connected"
	AccountStreamLiveInputUpdateResponseResultStatusReconnected              AccountStreamLiveInputUpdateResponseResultStatus = "reconnected"
	AccountStreamLiveInputUpdateResponseResultStatusReconnecting             AccountStreamLiveInputUpdateResponseResultStatus = "reconnecting"
	AccountStreamLiveInputUpdateResponseResultStatusClientDisconnect         AccountStreamLiveInputUpdateResponseResultStatus = "client_disconnect"
	AccountStreamLiveInputUpdateResponseResultStatusTtlExceeded              AccountStreamLiveInputUpdateResponseResultStatus = "ttl_exceeded"
	AccountStreamLiveInputUpdateResponseResultStatusFailedToConnect          AccountStreamLiveInputUpdateResponseResultStatus = "failed_to_connect"
	AccountStreamLiveInputUpdateResponseResultStatusFailedToReconnect        AccountStreamLiveInputUpdateResponseResultStatus = "failed_to_reconnect"
	AccountStreamLiveInputUpdateResponseResultStatusNewConfigurationAccepted AccountStreamLiveInputUpdateResponseResultStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type AccountStreamLiveInputUpdateResponseResultWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                               `json:"url"`
	JSON accountStreamLiveInputUpdateResponseResultWebRtcJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultWebRtcJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputUpdateResponseResultWebRtc]
type accountStreamLiveInputUpdateResponseResultWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type AccountStreamLiveInputUpdateResponseResultWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                                       `json:"url"`
	JSON accountStreamLiveInputUpdateResponseResultWebRtcPlaybackJSON `json:"-"`
}

// accountStreamLiveInputUpdateResponseResultWebRtcPlaybackJSON contains the JSON
// metadata for the struct
// [AccountStreamLiveInputUpdateResponseResultWebRtcPlayback]
type accountStreamLiveInputUpdateResponseResultWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputUpdateResponseResultWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputUpdateResponseSuccess bool

const (
	AccountStreamLiveInputUpdateResponseSuccessTrue AccountStreamLiveInputUpdateResponseSuccess = true
)

type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponse struct {
	Errors   []AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessage `json:"messages"`
	// Details about a live input.
	Result AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputStreamLiveInputsNewALiveInputResponseJSON    `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseJSON contains the
// JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponse]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseError struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountStreamLiveInputStreamLiveInputsNewALiveInputResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseError]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessage struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    accountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessage]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a live input.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResult struct {
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
	Recording AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlayback `json:"webRTCPlayback"`
	JSON           accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultJSON           `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultJSON contains
// the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResult]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultJSON struct {
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

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                                                          `json:"timeoutSeconds"`
	JSON           accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecording]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingMode string

const (
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingModeOff       AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingMode = "off"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingModeAutomatic AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                                                     `json:"url"`
	JSON accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmps]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                                                             `json:"url"`
	JSON accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlaybackJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlaybackJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlayback]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                                                                   `json:"url"`
	JSON accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrt]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                                                           `json:"url"`
	JSON accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlaybackJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlaybackJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlayback]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus string

const (
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusConnected                AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "connected"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusReconnected              AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "reconnected"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusReconnecting             AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "reconnecting"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusClientDisconnect         AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "client_disconnect"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusTtlExceeded              AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "ttl_exceeded"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusFailedToConnect          AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "failed_to_connect"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusFailedToReconnect        AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "failed_to_reconnect"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatusNewConfigurationAccepted AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                                                      `json:"url"`
	JSON accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtc]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                                                              `json:"url"`
	JSON accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlaybackJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlaybackJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlayback]
type accountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseResultWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseSuccess bool

const (
	AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseSuccessTrue AccountStreamLiveInputStreamLiveInputsNewALiveInputResponseSuccess = true
)

type AccountStreamLiveInputStreamLiveInputsListLiveInputsResponse struct {
	Errors   []AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessage `json:"messages"`
	Result   AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputStreamLiveInputsListLiveInputsResponseJSON    `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsListLiveInputsResponseJSON contains the
// JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsListLiveInputsResponse]
type accountStreamLiveInputStreamLiveInputsListLiveInputsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsListLiveInputsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountStreamLiveInputStreamLiveInputsListLiveInputsResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsListLiveInputsResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseError]
type accountStreamLiveInputStreamLiveInputsListLiveInputsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessage]
type accountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResult struct {
	LiveInputs []AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                                                                  `json:"total"`
	JSON  accountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultJSON contains
// the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResult]
type accountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInput struct {
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
	Uid  string                                                                          `json:"uid"`
	JSON accountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInputJSON `json:"-"`
}

// accountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInputJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInput]
type accountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Uid                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseResultLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseSuccess bool

const (
	AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseSuccessTrue AccountStreamLiveInputStreamLiveInputsListLiveInputsResponseSuccess = true
)

type AccountStreamLiveInputUpdateParams struct {
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
	Recording param.Field[AccountStreamLiveInputUpdateParamsRecording] `json:"recording"`
}

func (r AccountStreamLiveInputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type AccountStreamLiveInputUpdateParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[AccountStreamLiveInputUpdateParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r AccountStreamLiveInputUpdateParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type AccountStreamLiveInputUpdateParamsRecordingMode string

const (
	AccountStreamLiveInputUpdateParamsRecordingModeOff       AccountStreamLiveInputUpdateParamsRecordingMode = "off"
	AccountStreamLiveInputUpdateParamsRecordingModeAutomatic AccountStreamLiveInputUpdateParamsRecordingMode = "automatic"
)

type AccountStreamLiveInputStreamLiveInputsNewALiveInputParams struct {
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
	Recording param.Field[AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecording] `json:"recording"`
}

func (r AccountStreamLiveInputStreamLiveInputsNewALiveInputParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode string

const (
	AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingModeOff       AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode = "off"
	AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingModeAutomatic AccountStreamLiveInputStreamLiveInputsNewALiveInputParamsRecordingMode = "automatic"
)

type AccountStreamLiveInputStreamLiveInputsListLiveInputsParams struct {
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
}

// URLQuery serializes
// [AccountStreamLiveInputStreamLiveInputsListLiveInputsParams]'s query parameters
// as `url.Values`.
func (r AccountStreamLiveInputStreamLiveInputsListLiveInputsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
