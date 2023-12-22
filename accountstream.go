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
	return
}

// Fetches details for a single video.
func (r *AccountStreamService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *VideoResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update details for a single video.
func (r *AccountStreamService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountStreamUpdateParams, opts ...option.RequestOption) (res *VideoResponseSingle, err error) {
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
func (r *AccountStreamService) StreamVideosListVideos(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *VideoResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VideoResponseCollection struct {
	Errors   []VideoResponseCollectionError   `json:"errors"`
	Messages []VideoResponseCollectionMessage `json:"messages"`
	// The total number of remaining videos based on cursor position.
	Range      int64                             `json:"range"`
	Result     []VideoResponseCollectionResult   `json:"result"`
	ResultInfo VideoResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success VideoResponseCollectionSuccess `json:"success"`
	// The total number of videos that match the provided filters.
	Total int64                       `json:"total"`
	JSON  videoResponseCollectionJSON `json:"-"`
}

// videoResponseCollectionJSON contains the JSON metadata for the struct
// [VideoResponseCollection]
type videoResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Range       apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseCollectionError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    videoResponseCollectionErrorJSON `json:"-"`
}

// videoResponseCollectionErrorJSON contains the JSON metadata for the struct
// [VideoResponseCollectionError]
type videoResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseCollectionMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    videoResponseCollectionMessageJSON `json:"-"`
}

// videoResponseCollectionMessageJSON contains the JSON metadata for the struct
// [VideoResponseCollectionMessage]
type videoResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseCollectionResult struct {
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
	Duration int64                              `json:"duration"`
	Input    VideoResponseCollectionResultInput `json:"input"`
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
	Modified time.Time                             `json:"modified" format:"date-time"`
	Playback VideoResponseCollectionResultPlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video is ready for viewing.
	ReadyToStream bool `json:"readyToStream"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// The size of the media item in bytes.
	Size float64 `json:"size"`
	// Specifies a detailed status for a video. If the `state` is `inprogress` or
	// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
	// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
	// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
	// and `errorReasonText` provide additional details.
	Status VideoResponseCollectionResultStatus `json:"status"`
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
	UploadExpiry time.Time                              `json:"uploadExpiry" format:"date-time"`
	Watermark    VideoResponseCollectionResultWatermark `json:"watermark"`
	JSON         videoResponseCollectionResultJSON      `json:"-"`
}

// videoResponseCollectionResultJSON contains the JSON metadata for the struct
// [VideoResponseCollectionResult]
type videoResponseCollectionResultJSON struct {
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
	RequireSignedURLs     apijson.Field
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

func (r *VideoResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseCollectionResultInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64                                  `json:"width"`
	JSON  videoResponseCollectionResultInputJSON `json:"-"`
}

// videoResponseCollectionResultInputJSON contains the JSON metadata for the struct
// [VideoResponseCollectionResultInput]
type videoResponseCollectionResultInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseCollectionResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseCollectionResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                                    `json:"hls"`
	JSON videoResponseCollectionResultPlaybackJSON `json:"-"`
}

// videoResponseCollectionResultPlaybackJSON contains the JSON metadata for the
// struct [VideoResponseCollectionResultPlayback]
type videoResponseCollectionResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseCollectionResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type VideoResponseCollectionResultStatus struct {
	// Specifies why the video failed to encode. This field is empty if the video is
	// not in an `error` state. Preferred for programmatic use.
	ErrorReasonCode string `json:"errorReasonCode"`
	// Specifies why the video failed to encode using a human readable error message in
	// English. This field is empty if the video is not in an `error` state.
	ErrorReasonText string `json:"errorReasonText"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	PctComplete string `json:"pctComplete"`
	// Specifies the processing status of the video.
	State VideoResponseCollectionResultStatusState `json:"state"`
	JSON  videoResponseCollectionResultStatusJSON  `json:"-"`
}

// videoResponseCollectionResultStatusJSON contains the JSON metadata for the
// struct [VideoResponseCollectionResultStatus]
type videoResponseCollectionResultStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VideoResponseCollectionResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type VideoResponseCollectionResultStatusState string

const (
	VideoResponseCollectionResultStatusStatePendingupload VideoResponseCollectionResultStatusState = "pendingupload"
	VideoResponseCollectionResultStatusStateDownloading   VideoResponseCollectionResultStatusState = "downloading"
	VideoResponseCollectionResultStatusStateQueued        VideoResponseCollectionResultStatusState = "queued"
	VideoResponseCollectionResultStatusStateInprogress    VideoResponseCollectionResultStatusState = "inprogress"
	VideoResponseCollectionResultStatusStateReady         VideoResponseCollectionResultStatusState = "ready"
	VideoResponseCollectionResultStatusStateError         VideoResponseCollectionResultStatusState = "error"
)

type VideoResponseCollectionResultWatermark struct {
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
	Width int64                                      `json:"width"`
	JSON  videoResponseCollectionResultWatermarkJSON `json:"-"`
}

// videoResponseCollectionResultWatermarkJSON contains the JSON metadata for the
// struct [VideoResponseCollectionResultWatermark]
type videoResponseCollectionResultWatermarkJSON struct {
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

func (r *VideoResponseCollectionResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       videoResponseCollectionResultInfoJSON `json:"-"`
}

// videoResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [VideoResponseCollectionResultInfo]
type videoResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VideoResponseCollectionSuccess bool

const (
	VideoResponseCollectionSuccessTrue VideoResponseCollectionSuccess = true
)

type VideoResponseSingle struct {
	Errors   []VideoResponseSingleError   `json:"errors"`
	Messages []VideoResponseSingleMessage `json:"messages"`
	Result   VideoResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success VideoResponseSingleSuccess `json:"success"`
	JSON    videoResponseSingleJSON    `json:"-"`
}

// videoResponseSingleJSON contains the JSON metadata for the struct
// [VideoResponseSingle]
type videoResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseSingleError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    videoResponseSingleErrorJSON `json:"-"`
}

// videoResponseSingleErrorJSON contains the JSON metadata for the struct
// [VideoResponseSingleError]
type videoResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseSingleMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    videoResponseSingleMessageJSON `json:"-"`
}

// videoResponseSingleMessageJSON contains the JSON metadata for the struct
// [VideoResponseSingleMessage]
type videoResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseSingleResult struct {
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
	Duration int64                          `json:"duration"`
	Input    VideoResponseSingleResultInput `json:"input"`
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
	Modified time.Time                         `json:"modified" format:"date-time"`
	Playback VideoResponseSingleResultPlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video is ready for viewing.
	ReadyToStream bool `json:"readyToStream"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// The size of the media item in bytes.
	Size float64 `json:"size"`
	// Specifies a detailed status for a video. If the `state` is `inprogress` or
	// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
	// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
	// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
	// and `errorReasonText` provide additional details.
	Status VideoResponseSingleResultStatus `json:"status"`
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
	UploadExpiry time.Time                          `json:"uploadExpiry" format:"date-time"`
	Watermark    VideoResponseSingleResultWatermark `json:"watermark"`
	JSON         videoResponseSingleResultJSON      `json:"-"`
}

// videoResponseSingleResultJSON contains the JSON metadata for the struct
// [VideoResponseSingleResult]
type videoResponseSingleResultJSON struct {
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
	RequireSignedURLs     apijson.Field
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

func (r *VideoResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseSingleResultInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64                              `json:"width"`
	JSON  videoResponseSingleResultInputJSON `json:"-"`
}

// videoResponseSingleResultInputJSON contains the JSON metadata for the struct
// [VideoResponseSingleResultInput]
type videoResponseSingleResultInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseSingleResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VideoResponseSingleResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                                `json:"hls"`
	JSON videoResponseSingleResultPlaybackJSON `json:"-"`
}

// videoResponseSingleResultPlaybackJSON contains the JSON metadata for the struct
// [VideoResponseSingleResultPlayback]
type videoResponseSingleResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseSingleResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type VideoResponseSingleResultStatus struct {
	// Specifies why the video failed to encode. This field is empty if the video is
	// not in an `error` state. Preferred for programmatic use.
	ErrorReasonCode string `json:"errorReasonCode"`
	// Specifies why the video failed to encode using a human readable error message in
	// English. This field is empty if the video is not in an `error` state.
	ErrorReasonText string `json:"errorReasonText"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	PctComplete string `json:"pctComplete"`
	// Specifies the processing status of the video.
	State VideoResponseSingleResultStatusState `json:"state"`
	JSON  videoResponseSingleResultStatusJSON  `json:"-"`
}

// videoResponseSingleResultStatusJSON contains the JSON metadata for the struct
// [VideoResponseSingleResultStatus]
type videoResponseSingleResultStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VideoResponseSingleResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type VideoResponseSingleResultStatusState string

const (
	VideoResponseSingleResultStatusStatePendingupload VideoResponseSingleResultStatusState = "pendingupload"
	VideoResponseSingleResultStatusStateDownloading   VideoResponseSingleResultStatusState = "downloading"
	VideoResponseSingleResultStatusStateQueued        VideoResponseSingleResultStatusState = "queued"
	VideoResponseSingleResultStatusStateInprogress    VideoResponseSingleResultStatusState = "inprogress"
	VideoResponseSingleResultStatusStateReady         VideoResponseSingleResultStatusState = "ready"
	VideoResponseSingleResultStatusStateError         VideoResponseSingleResultStatusState = "error"
)

type VideoResponseSingleResultWatermark struct {
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
	Width int64                                  `json:"width"`
	JSON  videoResponseSingleResultWatermarkJSON `json:"-"`
}

// videoResponseSingleResultWatermarkJSON contains the JSON metadata for the struct
// [VideoResponseSingleResultWatermark]
type videoResponseSingleResultWatermarkJSON struct {
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

func (r *VideoResponseSingleResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VideoResponseSingleSuccess bool

const (
	VideoResponseSingleSuccessTrue VideoResponseSingleSuccess = true
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
}

// Specifies the TUS protocol version. This value must be included in every upload
// request. Notes: The only supported version of TUS protocol is 1.0.0.
type AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable string

const (
	AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable1_0_0 AccountStreamStreamVideosInitiateVideoUploadsUsingTusParamsTusResumable = "1.0.0"
)
