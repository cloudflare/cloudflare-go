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

// AccountStreamCopyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamCopyService] method
// instead.
type AccountStreamCopyService struct {
	Options []option.RequestOption
}

// NewAccountStreamCopyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamCopyService(opts ...option.RequestOption) (r *AccountStreamCopyService) {
	r = &AccountStreamCopyService{}
	r.Options = opts
	return
}

// Uploads a video to Stream from a provided URL.
func (r *AccountStreamCopyService) StreamVideosUploadVideosFromAURL(ctx context.Context, accountIdentifier string, params AccountStreamCopyStreamVideosUploadVideosFromAURLParams, opts ...option.RequestOption) (res *AccountStreamCopyStreamVideosUploadVideosFromAurlResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/copy", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponse struct {
	Errors   []AccountStreamCopyStreamVideosUploadVideosFromAurlResponseError   `json:"errors"`
	Messages []AccountStreamCopyStreamVideosUploadVideosFromAurlResponseMessage `json:"messages"`
	Result   AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamCopyStreamVideosUploadVideosFromAurlResponseSuccess `json:"success"`
	JSON    accountStreamCopyStreamVideosUploadVideosFromAurlResponseJSON    `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseJSON contains the JSON
// metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponse]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountStreamCopyStreamVideosUploadVideosFromAurlResponseErrorJSON `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseError]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountStreamCopyStreamVideosUploadVideosFromAurlResponseMessageJSON `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseMessage]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResult struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// The date and time the media item was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The duration of the video in seconds. A value of `-1` means the duration is
	// unknown. The duration becomes available after the upload and before the video is
	// ready.
	Duration float64                                                              `json:"duration"`
	Input    AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInput `json:"input"`
	// The live input ID used to upload a video with Stream Live.
	LiveInput string `json:"liveInput"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds int64 `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta interface{} `json:"meta"`
	// The date and time the media item was last modified.
	Modified time.Time                                                               `json:"modified" format:"date-time"`
	Playback AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video is playable. The field is empty if the video is not
	// ready for viewing or the live stream is still in progress.
	ReadyToStream bool `json:"readyToStream"`
	// Indicates the time at which the video became playable. The field is empty if the
	// video is not ready for viewing or the live stream is still in progress.
	ReadyToStreamAt time.Time `json:"readyToStreamAt" format:"date-time"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion time.Time `json:"scheduledDeletion" format:"date-time"`
	// The size of the media item in bytes.
	Size float64 `json:"size"`
	// Specifies a detailed status for a video. If the `state` is `inprogress` or
	// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
	// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
	// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
	// and `errorReasonText` provide additional details.
	Status AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatus `json:"status"`
	// The media item's thumbnail URI. This field is omitted until encoding is
	// complete.
	Thumbnail string `json:"thumbnail" format:"uri"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64 `json:"thumbnailTimestampPct"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid string `json:"uid"`
	// The date and time the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// The date and time when the video upload URL is no longer valid for direct user
	// uploads.
	UploadExpiry time.Time                                                                `json:"uploadExpiry" format:"date-time"`
	Watermark    AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermark `json:"watermark"`
	JSON         accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultJSON      `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultJSON contains the
// JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResult]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultJSON struct {
	AllowedOrigins        apijson.Field
	Created               apijson.Field
	Creator               apijson.Field
	Duration              apijson.Field
	Input                 apijson.Field
	LiveInput             apijson.Field
	MaxDurationSeconds    apijson.Field
	Meta                  apijson.Field
	Modified              apijson.Field
	Playback              apijson.Field
	Preview               apijson.Field
	ReadyToStream         apijson.Field
	ReadyToStreamAt       apijson.Field
	RequireSignedURLs     apijson.Field
	ScheduledDeletion     apijson.Field
	Size                  apijson.Field
	Status                apijson.Field
	Thumbnail             apijson.Field
	ThumbnailTimestampPct apijson.Field
	Uid                   apijson.Field
	Uploaded              apijson.Field
	UploadExpiry          apijson.Field
	Watermark             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64                                                                    `json:"width"`
	JSON  accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInputJSON `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInputJSON
// contains the JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInput]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                                                                      `json:"hls"`
	JSON accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlaybackJSON `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlaybackJSON
// contains the JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlayback]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatus struct {
	// Specifies why the video failed to encode. This field is empty if the video is
	// not in an `error` state. Preferred for programmatic use.
	ErrorReasonCode string `json:"errorReasonCode"`
	// Specifies why the video failed to encode using a human readable error message in
	// English. This field is empty if the video is not in an `error` state.
	ErrorReasonText string `json:"errorReasonText"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	PctComplete string `json:"pctComplete"`
	// Specifies the processing status for all quality levels for a video.
	State AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState `json:"state"`
	JSON  accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusJSON  `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusJSON
// contains the JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatus]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status for all quality levels for a video.
type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState string

const (
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusStatePendingupload AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState = "pendingupload"
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusStateDownloading   AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState = "downloading"
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusStateQueued        AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState = "queued"
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusStateInprogress    AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState = "inprogress"
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusStateReady         AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState = "ready"
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusStateError         AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultStatusState = "error"
)

type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermark struct {
	// The date and a time a watermark profile was created.
	Created time.Time `json:"created" format:"date-time"`
	// The source URL for a downloaded image. If the watermark profile was created via
	// direct upload, this field is null.
	DownloadedFrom string `json:"downloadedFrom"`
	// The height of the image in pixels.
	Height int64 `json:"height"`
	// A short description of the watermark profile.
	Name string `json:"name"`
	// The translucency of the image. A value of `0.0` makes the image completely
	// transparent, and `1.0` makes the image completely opaque. Note that if the image
	// is already semi-transparent, setting this to `1.0` will not make the image
	// completely opaque.
	Opacity float64 `json:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video
	// and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded
	// video width or length, as determined by the algorithm.
	Padding float64 `json:"padding"`
	// The location of the image. Valid positions are: `upperRight`, `upperLeft`,
	// `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the
	// `padding` parameter.
	Position string `json:"position"`
	// The size of the image relative to the overall size of the video. This parameter
	// will adapt to horizontal and vertical videos automatically. `0.0` indicates no
	// scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	Scale float64 `json:"scale"`
	// The size of the image in bytes.
	Size float64 `json:"size"`
	// The unique identifier for a watermark profile.
	Uid string `json:"uid"`
	// The width of the image in pixels.
	Width int64                                                                        `json:"width"`
	JSON  accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermarkJSON `json:"-"`
}

// accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermarkJSON
// contains the JSON metadata for the struct
// [AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermark]
type accountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermarkJSON struct {
	Created        apijson.Field
	DownloadedFrom apijson.Field
	Height         apijson.Field
	Name           apijson.Field
	Opacity        apijson.Field
	Padding        apijson.Field
	Position       apijson.Field
	Scale          apijson.Field
	Size           apijson.Field
	Uid            apijson.Field
	Width          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountStreamCopyStreamVideosUploadVideosFromAurlResponseResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamCopyStreamVideosUploadVideosFromAurlResponseSuccess bool

const (
	AccountStreamCopyStreamVideosUploadVideosFromAurlResponseSuccessTrue AccountStreamCopyStreamVideosUploadVideosFromAurlResponseSuccess = true
)

type AccountStreamCopyStreamVideosUploadVideosFromAURLParams struct {
	// A video's URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta param.Field[interface{}] `json:"meta"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion param.Field[time.Time] `json:"scheduledDeletion" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                                                          `json:"thumbnailTimestampPct"`
	Watermark             param.Field[AccountStreamCopyStreamVideosUploadVideosFromAurlParamsWatermark] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
	// Comma-separated key-value pairs following the TUS protocol specification. Values
	// are Base-64 encoded. Supported keys: `name`, `requiresignedurls`,
	// `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`.
	UploadMetadata param.Field[string] `header:"Upload-Metadata"`
}

func (r AccountStreamCopyStreamVideosUploadVideosFromAURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamCopyStreamVideosUploadVideosFromAurlParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r AccountStreamCopyStreamVideosUploadVideosFromAurlParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
