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
func (r *AccountStreamLiveInputService) Get(ctx context.Context, accountIdentifier string, liveInputIdentifier string, opts ...option.RequestOption) (res *LiveInputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specified live input.
func (r *AccountStreamLiveInputService) Update(ctx context.Context, accountIdentifier string, liveInputIdentifier string, body AccountStreamLiveInputUpdateParams, opts ...option.RequestOption) (res *LiveInputResponseSingle, err error) {
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
func (r *AccountStreamLiveInputService) StreamLiveInputsNewALiveInput(ctx context.Context, accountIdentifier string, body AccountStreamLiveInputStreamLiveInputsNewALiveInputParams, opts ...option.RequestOption) (res *LiveInputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *AccountStreamLiveInputService) StreamLiveInputsListLiveInputs(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *LiveInputResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LiveInputResponseCollection struct {
	Errors     []LiveInputResponseCollectionError    `json:"errors"`
	Messages   []LiveInputResponseCollectionMessage  `json:"messages"`
	Result     LiveInputResponseCollectionResult     `json:"result"`
	ResultInfo LiveInputResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LiveInputResponseCollectionSuccess `json:"success"`
	JSON    liveInputResponseCollectionJSON    `json:"-"`
}

// liveInputResponseCollectionJSON contains the JSON metadata for the struct
// [LiveInputResponseCollection]
type liveInputResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseCollectionError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    liveInputResponseCollectionErrorJSON `json:"-"`
}

// liveInputResponseCollectionErrorJSON contains the JSON metadata for the struct
// [LiveInputResponseCollectionError]
type liveInputResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseCollectionMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    liveInputResponseCollectionMessageJSON `json:"-"`
}

// liveInputResponseCollectionMessageJSON contains the JSON metadata for the struct
// [LiveInputResponseCollectionMessage]
type liveInputResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseCollectionResult struct {
	LiveInputs []LiveInputResponseCollectionResultLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                                 `json:"total"`
	JSON  liveInputResponseCollectionResultJSON `json:"-"`
}

// liveInputResponseCollectionResultJSON contains the JSON metadata for the struct
// [LiveInputResponseCollectionResult]
type liveInputResponseCollectionResultJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseCollectionResultLiveInput struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A unique identifier for a live input.
	Uid  string                                         `json:"uid"`
	JSON liveInputResponseCollectionResultLiveInputJSON `json:"-"`
}

// liveInputResponseCollectionResultLiveInputJSON contains the JSON metadata for
// the struct [LiveInputResponseCollectionResultLiveInput]
type liveInputResponseCollectionResultLiveInputJSON struct {
	Created     apijson.Field
	Meta        apijson.Field
	Modified    apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseCollectionResultLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       liveInputResponseCollectionResultInfoJSON `json:"-"`
}

// liveInputResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [LiveInputResponseCollectionResultInfo]
type liveInputResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LiveInputResponseCollectionSuccess bool

const (
	LiveInputResponseCollectionSuccessTrue LiveInputResponseCollectionSuccess = true
)

type LiveInputResponseSingle struct {
	Errors   []LiveInputResponseSingleError   `json:"errors"`
	Messages []LiveInputResponseSingleMessage `json:"messages"`
	// Details about a live input.
	Result LiveInputResponseSingleResult `json:"result"`
	// Whether the API call was successful
	Success LiveInputResponseSingleSuccess `json:"success"`
	JSON    liveInputResponseSingleJSON    `json:"-"`
}

// liveInputResponseSingleJSON contains the JSON metadata for the struct
// [LiveInputResponseSingle]
type liveInputResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseSingleError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    liveInputResponseSingleErrorJSON `json:"-"`
}

// liveInputResponseSingleErrorJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleError]
type liveInputResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LiveInputResponseSingleMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    liveInputResponseSingleMessageJSON `json:"-"`
}

// liveInputResponseSingleMessageJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleMessage]
type liveInputResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a live input.
type LiveInputResponseSingleResult struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording LiveInputResponseSingleResultRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps LiveInputResponseSingleResultRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback LiveInputResponseSingleResultRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt LiveInputResponseSingleResultSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback LiveInputResponseSingleResultSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status LiveInputResponseSingleResultStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc LiveInputResponseSingleResultWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback LiveInputResponseSingleResultWebRtcPlayback `json:"webRTCPlayback"`
	JSON           liveInputResponseSingleResultJSON           `json:"-"`
}

// liveInputResponseSingleResultJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleResult]
type liveInputResponseSingleResultJSON struct {
	Created        apijson.Field
	Meta           apijson.Field
	Modified       apijson.Field
	Recording      apijson.Field
	Rtmps          apijson.Field
	RtmpsPlayback  apijson.Field
	Srt            apijson.Field
	SrtPlayback    apijson.Field
	Status         apijson.Field
	Uid            apijson.Field
	WebRtc         apijson.Field
	WebRtcPlayback apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LiveInputResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputResponseSingleResultRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode LiveInputResponseSingleResultRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                                      `json:"timeoutSeconds"`
	JSON           liveInputResponseSingleResultRecordingJSON `json:"-"`
}

// liveInputResponseSingleResultRecordingJSON contains the JSON metadata for the
// struct [LiveInputResponseSingleResultRecording]
type liveInputResponseSingleResultRecordingJSON struct {
	AllowedOrigins    apijson.Field
	Mode              apijson.Field
	RequireSignedURLs apijson.Field
	TimeoutSeconds    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputResponseSingleResultRecordingMode string

const (
	LiveInputResponseSingleResultRecordingModeOff       LiveInputResponseSingleResultRecordingMode = "off"
	LiveInputResponseSingleResultRecordingModeAutomatic LiveInputResponseSingleResultRecordingMode = "automatic"
)

// Details for streaming to an live input using RTMPS.
type LiveInputResponseSingleResultRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                 `json:"url"`
	JSON liveInputResponseSingleResultRtmpsJSON `json:"-"`
}

// liveInputResponseSingleResultRtmpsJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleResultRtmps]
type liveInputResponseSingleResultRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using RTMPS.
type LiveInputResponseSingleResultRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                         `json:"url"`
	JSON liveInputResponseSingleResultRtmpsPlaybackJSON `json:"-"`
}

// liveInputResponseSingleResultRtmpsPlaybackJSON contains the JSON metadata for
// the struct [LiveInputResponseSingleResultRtmpsPlayback]
type liveInputResponseSingleResultRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for streaming to a live input using SRT.
type LiveInputResponseSingleResultSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                               `json:"url"`
	JSON liveInputResponseSingleResultSrtJSON `json:"-"`
}

// liveInputResponseSingleResultSrtJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleResultSrt]
type liveInputResponseSingleResultSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from an live input using SRT.
type LiveInputResponseSingleResultSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                       `json:"url"`
	JSON liveInputResponseSingleResultSrtPlaybackJSON `json:"-"`
}

// liveInputResponseSingleResultSrtPlaybackJSON contains the JSON metadata for the
// struct [LiveInputResponseSingleResultSrtPlayback]
type liveInputResponseSingleResultSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of a live input.
type LiveInputResponseSingleResultStatus string

const (
	LiveInputResponseSingleResultStatusConnected                LiveInputResponseSingleResultStatus = "connected"
	LiveInputResponseSingleResultStatusReconnected              LiveInputResponseSingleResultStatus = "reconnected"
	LiveInputResponseSingleResultStatusReconnecting             LiveInputResponseSingleResultStatus = "reconnecting"
	LiveInputResponseSingleResultStatusClientDisconnect         LiveInputResponseSingleResultStatus = "client_disconnect"
	LiveInputResponseSingleResultStatusTtlExceeded              LiveInputResponseSingleResultStatus = "ttl_exceeded"
	LiveInputResponseSingleResultStatusFailedToConnect          LiveInputResponseSingleResultStatus = "failed_to_connect"
	LiveInputResponseSingleResultStatusFailedToReconnect        LiveInputResponseSingleResultStatus = "failed_to_reconnect"
	LiveInputResponseSingleResultStatusNewConfigurationAccepted LiveInputResponseSingleResultStatus = "new_configuration_accepted"
)

// Details for streaming to a live input using WebRTC.
type LiveInputResponseSingleResultWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                  `json:"url"`
	JSON liveInputResponseSingleResultWebRtcJSON `json:"-"`
}

// liveInputResponseSingleResultWebRtcJSON contains the JSON metadata for the
// struct [LiveInputResponseSingleResultWebRtc]
type liveInputResponseSingleResultWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details for playback from a live input using WebRTC.
type LiveInputResponseSingleResultWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                          `json:"url"`
	JSON liveInputResponseSingleResultWebRtcPlaybackJSON `json:"-"`
}

// liveInputResponseSingleResultWebRtcPlaybackJSON contains the JSON metadata for
// the struct [LiveInputResponseSingleResultWebRtcPlayback]
type liveInputResponseSingleResultWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LiveInputResponseSingleSuccess bool

const (
	LiveInputResponseSingleSuccessTrue LiveInputResponseSingleSuccess = true
)

type AccountStreamLiveInputUpdateParams struct {
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
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
