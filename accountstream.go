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

// AccountStreamService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamService] method
// instead.
type AccountStreamService struct {
	Options       []option.RequestOption
	Clips         *AccountStreamClipService
	Copies        *AccountStreamCopyService
	DirectUploads *AccountStreamDirectUploadService
	Keys          *AccountStreamKeyService
	LiveInputs    *AccountStreamLiveInputService
	Watermarks    *AccountStreamWatermarkService
	Webhooks      *AccountStreamWebhookService
	Captions      *AccountStreamCaptionService
	Downloads     *AccountStreamDownloadService
	Embeds        *AccountStreamEmbedService
	Tokens        *AccountStreamTokenService
	Audios        *AccountStreamAudioService
	StorageUsages *AccountStreamStorageUsageService
}

// NewAccountStreamService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountStreamService(opts ...option.RequestOption) (r *AccountStreamService) {
	r = &AccountStreamService{}
	r.Options = opts
	r.Clips = NewAccountStreamClipService(opts...)
	r.Copies = NewAccountStreamCopyService(opts...)
	r.DirectUploads = NewAccountStreamDirectUploadService(opts...)
	r.Keys = NewAccountStreamKeyService(opts...)
	r.LiveInputs = NewAccountStreamLiveInputService(opts...)
	r.Watermarks = NewAccountStreamWatermarkService(opts...)
	r.Webhooks = NewAccountStreamWebhookService(opts...)
	r.Captions = NewAccountStreamCaptionService(opts...)
	r.Downloads = NewAccountStreamDownloadService(opts...)
	r.Embeds = NewAccountStreamEmbedService(opts...)
	r.Tokens = NewAccountStreamTokenService(opts...)
	r.Audios = NewAccountStreamAudioService(opts...)
	r.StorageUsages = NewAccountStreamStorageUsageService(opts...)
	return
}

// Fetches details for a single video.
func (r *AccountStreamService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit details for a single video.
func (r *AccountStreamService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountStreamUpdateParams, opts ...option.RequestOption) (res *AccountStreamUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a video and its copies from Cloudflare Stream.
func (r *AccountStreamService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Initiates a video upload using the TUS protocol. On success, the server responds
// with a status code 201 (created) and includes a `location` header to indicate
// where the content should be uploaded. Refer to https://tus.io for protocol
// details.
func (r *AccountStreamService) StreamVideosInitiateVideoUploadsUsingTus(ctx context.Context, accountIdentifier string, body AccountStreamStreamVideosInitiateVideoUploadsUsingTusParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists up to 1000 videos from a single request. For a specific range, refer to
// the optional parameters.
func (r *AccountStreamService) StreamVideosListVideos(ctx context.Context, accountIdentifier string, query AccountStreamStreamVideosListVideosParams, opts ...option.RequestOption) (res *AccountStreamStreamVideosListVideosResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountStreamGetResponse struct {
	Errors   []AccountStreamGetResponseError   `json:"errors"`
	Messages []AccountStreamGetResponseMessage `json:"messages"`
	Result   AccountStreamGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamGetResponseSuccess `json:"success"`
	JSON    accountStreamGetResponseJSON    `json:"-"`
}

// accountStreamGetResponseJSON contains the JSON metadata for the struct
// [AccountStreamGetResponse]
type accountStreamGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accountStreamGetResponseErrorJSON `json:"-"`
}

// accountStreamGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountStreamGetResponseError]
type accountStreamGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountStreamGetResponseMessageJSON `json:"-"`
}

// accountStreamGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountStreamGetResponseMessage]
type accountStreamGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamGetResponseResult struct {
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
	Duration float64                             `json:"duration"`
	Input    AccountStreamGetResponseResultInput `json:"input"`
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
	Modified time.Time                              `json:"modified" format:"date-time"`
	Playback AccountStreamGetResponseResultPlayback `json:"playback"`
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
	Status AccountStreamGetResponseResultStatus `json:"status"`
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
	UploadExpiry time.Time                               `json:"uploadExpiry" format:"date-time"`
	Watermark    AccountStreamGetResponseResultWatermark `json:"watermark"`
	JSON         accountStreamGetResponseResultJSON      `json:"-"`
}

// accountStreamGetResponseResultJSON contains the JSON metadata for the struct
// [AccountStreamGetResponseResult]
type accountStreamGetResponseResultJSON struct {
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

func (r *AccountStreamGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamGetResponseResultInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64                                   `json:"width"`
	JSON  accountStreamGetResponseResultInputJSON `json:"-"`
}

// accountStreamGetResponseResultInputJSON contains the JSON metadata for the
// struct [AccountStreamGetResponseResultInput]
type accountStreamGetResponseResultInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamGetResponseResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamGetResponseResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                                     `json:"hls"`
	JSON accountStreamGetResponseResultPlaybackJSON `json:"-"`
}

// accountStreamGetResponseResultPlaybackJSON contains the JSON metadata for the
// struct [AccountStreamGetResponseResultPlayback]
type accountStreamGetResponseResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamGetResponseResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type AccountStreamGetResponseResultStatus struct {
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
	State AccountStreamGetResponseResultStatusState `json:"state"`
	JSON  accountStreamGetResponseResultStatusJSON  `json:"-"`
}

// accountStreamGetResponseResultStatusJSON contains the JSON metadata for the
// struct [AccountStreamGetResponseResultStatus]
type accountStreamGetResponseResultStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountStreamGetResponseResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status for all quality levels for a video.
type AccountStreamGetResponseResultStatusState string

const (
	AccountStreamGetResponseResultStatusStatePendingupload AccountStreamGetResponseResultStatusState = "pendingupload"
	AccountStreamGetResponseResultStatusStateDownloading   AccountStreamGetResponseResultStatusState = "downloading"
	AccountStreamGetResponseResultStatusStateQueued        AccountStreamGetResponseResultStatusState = "queued"
	AccountStreamGetResponseResultStatusStateInprogress    AccountStreamGetResponseResultStatusState = "inprogress"
	AccountStreamGetResponseResultStatusStateReady         AccountStreamGetResponseResultStatusState = "ready"
	AccountStreamGetResponseResultStatusStateError         AccountStreamGetResponseResultStatusState = "error"
)

type AccountStreamGetResponseResultWatermark struct {
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
	Width int64                                       `json:"width"`
	JSON  accountStreamGetResponseResultWatermarkJSON `json:"-"`
}

// accountStreamGetResponseResultWatermarkJSON contains the JSON metadata for the
// struct [AccountStreamGetResponseResultWatermark]
type accountStreamGetResponseResultWatermarkJSON struct {
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

func (r *AccountStreamGetResponseResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamGetResponseSuccess bool

const (
	AccountStreamGetResponseSuccessTrue AccountStreamGetResponseSuccess = true
)

type AccountStreamUpdateResponse struct {
	Errors   []AccountStreamUpdateResponseError   `json:"errors"`
	Messages []AccountStreamUpdateResponseMessage `json:"messages"`
	Result   AccountStreamUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamUpdateResponseSuccess `json:"success"`
	JSON    accountStreamUpdateResponseJSON    `json:"-"`
}

// accountStreamUpdateResponseJSON contains the JSON metadata for the struct
// [AccountStreamUpdateResponse]
type accountStreamUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamUpdateResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountStreamUpdateResponseErrorJSON `json:"-"`
}

// accountStreamUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountStreamUpdateResponseError]
type accountStreamUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamUpdateResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountStreamUpdateResponseMessageJSON `json:"-"`
}

// accountStreamUpdateResponseMessageJSON contains the JSON metadata for the struct
// [AccountStreamUpdateResponseMessage]
type accountStreamUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamUpdateResponseResult struct {
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
	Duration float64                                `json:"duration"`
	Input    AccountStreamUpdateResponseResultInput `json:"input"`
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
	Modified time.Time                                 `json:"modified" format:"date-time"`
	Playback AccountStreamUpdateResponseResultPlayback `json:"playback"`
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
	Status AccountStreamUpdateResponseResultStatus `json:"status"`
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
	UploadExpiry time.Time                                  `json:"uploadExpiry" format:"date-time"`
	Watermark    AccountStreamUpdateResponseResultWatermark `json:"watermark"`
	JSON         accountStreamUpdateResponseResultJSON      `json:"-"`
}

// accountStreamUpdateResponseResultJSON contains the JSON metadata for the struct
// [AccountStreamUpdateResponseResult]
type accountStreamUpdateResponseResultJSON struct {
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

func (r *AccountStreamUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamUpdateResponseResultInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64                                      `json:"width"`
	JSON  accountStreamUpdateResponseResultInputJSON `json:"-"`
}

// accountStreamUpdateResponseResultInputJSON contains the JSON metadata for the
// struct [AccountStreamUpdateResponseResultInput]
type accountStreamUpdateResponseResultInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamUpdateResponseResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamUpdateResponseResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                                        `json:"hls"`
	JSON accountStreamUpdateResponseResultPlaybackJSON `json:"-"`
}

// accountStreamUpdateResponseResultPlaybackJSON contains the JSON metadata for the
// struct [AccountStreamUpdateResponseResultPlayback]
type accountStreamUpdateResponseResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamUpdateResponseResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type AccountStreamUpdateResponseResultStatus struct {
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
	State AccountStreamUpdateResponseResultStatusState `json:"state"`
	JSON  accountStreamUpdateResponseResultStatusJSON  `json:"-"`
}

// accountStreamUpdateResponseResultStatusJSON contains the JSON metadata for the
// struct [AccountStreamUpdateResponseResultStatus]
type accountStreamUpdateResponseResultStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountStreamUpdateResponseResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status for all quality levels for a video.
type AccountStreamUpdateResponseResultStatusState string

const (
	AccountStreamUpdateResponseResultStatusStatePendingupload AccountStreamUpdateResponseResultStatusState = "pendingupload"
	AccountStreamUpdateResponseResultStatusStateDownloading   AccountStreamUpdateResponseResultStatusState = "downloading"
	AccountStreamUpdateResponseResultStatusStateQueued        AccountStreamUpdateResponseResultStatusState = "queued"
	AccountStreamUpdateResponseResultStatusStateInprogress    AccountStreamUpdateResponseResultStatusState = "inprogress"
	AccountStreamUpdateResponseResultStatusStateReady         AccountStreamUpdateResponseResultStatusState = "ready"
	AccountStreamUpdateResponseResultStatusStateError         AccountStreamUpdateResponseResultStatusState = "error"
)

type AccountStreamUpdateResponseResultWatermark struct {
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
	Width int64                                          `json:"width"`
	JSON  accountStreamUpdateResponseResultWatermarkJSON `json:"-"`
}

// accountStreamUpdateResponseResultWatermarkJSON contains the JSON metadata for
// the struct [AccountStreamUpdateResponseResultWatermark]
type accountStreamUpdateResponseResultWatermarkJSON struct {
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

func (r *AccountStreamUpdateResponseResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamUpdateResponseSuccess bool

const (
	AccountStreamUpdateResponseSuccessTrue AccountStreamUpdateResponseSuccess = true
)

type AccountStreamStreamVideosListVideosResponse struct {
	Errors   []AccountStreamStreamVideosListVideosResponseError   `json:"errors"`
	Messages []AccountStreamStreamVideosListVideosResponseMessage `json:"messages"`
	// The total number of remaining videos based on cursor position.
	Range  int64                                               `json:"range"`
	Result []AccountStreamStreamVideosListVideosResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountStreamStreamVideosListVideosResponseSuccess `json:"success"`
	// The total number of videos that match the provided filters.
	Total int64                                           `json:"total"`
	JSON  accountStreamStreamVideosListVideosResponseJSON `json:"-"`
}

// accountStreamStreamVideosListVideosResponseJSON contains the JSON metadata for
// the struct [AccountStreamStreamVideosListVideosResponse]
type accountStreamStreamVideosListVideosResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Range       apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStreamVideosListVideosResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStreamVideosListVideosResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountStreamStreamVideosListVideosResponseErrorJSON `json:"-"`
}

// accountStreamStreamVideosListVideosResponseErrorJSON contains the JSON metadata
// for the struct [AccountStreamStreamVideosListVideosResponseError]
type accountStreamStreamVideosListVideosResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStreamVideosListVideosResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStreamVideosListVideosResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountStreamStreamVideosListVideosResponseMessageJSON `json:"-"`
}

// accountStreamStreamVideosListVideosResponseMessageJSON contains the JSON
// metadata for the struct [AccountStreamStreamVideosListVideosResponseMessage]
type accountStreamStreamVideosListVideosResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStreamVideosListVideosResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStreamVideosListVideosResponseResult struct {
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
	Duration float64                                                `json:"duration"`
	Input    AccountStreamStreamVideosListVideosResponseResultInput `json:"input"`
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
	Modified time.Time                                                 `json:"modified" format:"date-time"`
	Playback AccountStreamStreamVideosListVideosResponseResultPlayback `json:"playback"`
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
	Status AccountStreamStreamVideosListVideosResponseResultStatus `json:"status"`
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
	UploadExpiry time.Time                                                  `json:"uploadExpiry" format:"date-time"`
	Watermark    AccountStreamStreamVideosListVideosResponseResultWatermark `json:"watermark"`
	JSON         accountStreamStreamVideosListVideosResponseResultJSON      `json:"-"`
}

// accountStreamStreamVideosListVideosResponseResultJSON contains the JSON metadata
// for the struct [AccountStreamStreamVideosListVideosResponseResult]
type accountStreamStreamVideosListVideosResponseResultJSON struct {
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

func (r *AccountStreamStreamVideosListVideosResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStreamVideosListVideosResponseResultInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64                                                      `json:"width"`
	JSON  accountStreamStreamVideosListVideosResponseResultInputJSON `json:"-"`
}

// accountStreamStreamVideosListVideosResponseResultInputJSON contains the JSON
// metadata for the struct [AccountStreamStreamVideosListVideosResponseResultInput]
type accountStreamStreamVideosListVideosResponseResultInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStreamVideosListVideosResponseResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStreamVideosListVideosResponseResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                                                        `json:"hls"`
	JSON accountStreamStreamVideosListVideosResponseResultPlaybackJSON `json:"-"`
}

// accountStreamStreamVideosListVideosResponseResultPlaybackJSON contains the JSON
// metadata for the struct
// [AccountStreamStreamVideosListVideosResponseResultPlayback]
type accountStreamStreamVideosListVideosResponseResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStreamVideosListVideosResponseResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type AccountStreamStreamVideosListVideosResponseResultStatus struct {
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
	State AccountStreamStreamVideosListVideosResponseResultStatusState `json:"state"`
	JSON  accountStreamStreamVideosListVideosResponseResultStatusJSON  `json:"-"`
}

// accountStreamStreamVideosListVideosResponseResultStatusJSON contains the JSON
// metadata for the struct
// [AccountStreamStreamVideosListVideosResponseResultStatus]
type accountStreamStreamVideosListVideosResponseResultStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountStreamStreamVideosListVideosResponseResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status for all quality levels for a video.
type AccountStreamStreamVideosListVideosResponseResultStatusState string

const (
	AccountStreamStreamVideosListVideosResponseResultStatusStatePendingupload AccountStreamStreamVideosListVideosResponseResultStatusState = "pendingupload"
	AccountStreamStreamVideosListVideosResponseResultStatusStateDownloading   AccountStreamStreamVideosListVideosResponseResultStatusState = "downloading"
	AccountStreamStreamVideosListVideosResponseResultStatusStateQueued        AccountStreamStreamVideosListVideosResponseResultStatusState = "queued"
	AccountStreamStreamVideosListVideosResponseResultStatusStateInprogress    AccountStreamStreamVideosListVideosResponseResultStatusState = "inprogress"
	AccountStreamStreamVideosListVideosResponseResultStatusStateReady         AccountStreamStreamVideosListVideosResponseResultStatusState = "ready"
	AccountStreamStreamVideosListVideosResponseResultStatusStateError         AccountStreamStreamVideosListVideosResponseResultStatusState = "error"
)

type AccountStreamStreamVideosListVideosResponseResultWatermark struct {
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
	Width int64                                                          `json:"width"`
	JSON  accountStreamStreamVideosListVideosResponseResultWatermarkJSON `json:"-"`
}

// accountStreamStreamVideosListVideosResponseResultWatermarkJSON contains the JSON
// metadata for the struct
// [AccountStreamStreamVideosListVideosResponseResultWatermark]
type accountStreamStreamVideosListVideosResponseResultWatermarkJSON struct {
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

func (r *AccountStreamStreamVideosListVideosResponseResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamStreamVideosListVideosResponseSuccess bool

const (
	AccountStreamStreamVideosListVideosResponseSuccessTrue AccountStreamStreamVideosListVideosResponseSuccess = true
)

type AccountStreamUpdateParams struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds"`
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
	ThumbnailTimestampPct param.Field[float64] `json:"thumbnailTimestampPct"`
	// The date and time when the video upload URL is no longer valid for direct user
	// uploads.
	UploadExpiry param.Field[time.Time] `json:"uploadExpiry" format:"date-time"`
}

func (r AccountStreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamStreamVideosInitiateVideoUploadsUsingTusParams struct {
	// Specifies the TUS protocol version. This value must be included in every upload
	// request. Notes: The only supported version of TUS protocol is 1.0.0.
	TusResumable param.Field[AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable] `header:"Tus-Resumable,required"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	UploadLength param.Field[int64] `header:"Upload-Length,required"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
	// Comma-separated key-value pairs following the TUS protocol specification. Values
	// are Base-64 encoded. Supported keys: `name`, `requiresignedurls`,
	// `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`.
	UploadMetadata param.Field[string] `header:"Upload-Metadata"`
}

// Specifies the TUS protocol version. This value must be included in every upload
// request. Notes: The only supported version of TUS protocol is 1.0.0.
type AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable string

const (
	AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable1_0_0 AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable = "1.0.0"
)

type AccountStreamStreamVideosListVideosParams struct {
	// Lists videos in ascending order of creation.
	Asc param.Field[bool] `query:"asc"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
	// Lists videos created before the specified date.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
	// Searches over the `name` key in the `meta` field. This field can be set with or
	// after the upload request.
	Search param.Field[string] `query:"search"`
	// Lists videos created after the specified date.
	Start param.Field[time.Time] `query:"start" format:"date-time"`
	// Specifies the processing status for all quality levels for a video.
	Status param.Field[AccountStreamStreamVideosListVideosParamsStatus] `query:"status"`
	// Specifies whether the video is `vod` or `live`.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [AccountStreamStreamVideosListVideosParams]'s query
// parameters as `url.Values`.
func (r AccountStreamStreamVideosListVideosParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the processing status for all quality levels for a video.
type AccountStreamStreamVideosListVideosParamsStatus string

const (
	AccountStreamStreamVideosListVideosParamsStatusPendingupload AccountStreamStreamVideosListVideosParamsStatus = "pendingupload"
	AccountStreamStreamVideosListVideosParamsStatusDownloading   AccountStreamStreamVideosListVideosParamsStatus = "downloading"
	AccountStreamStreamVideosListVideosParamsStatusQueued        AccountStreamStreamVideosListVideosParamsStatus = "queued"
	AccountStreamStreamVideosListVideosParamsStatusInprogress    AccountStreamStreamVideosListVideosParamsStatus = "inprogress"
	AccountStreamStreamVideosListVideosParamsStatusReady         AccountStreamStreamVideosListVideosParamsStatus = "ready"
	AccountStreamStreamVideosListVideosParamsStatusError         AccountStreamStreamVideosListVideosParamsStatus = "error"
)
