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

// AccountStreamDirectUploadService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountStreamDirectUploadService] method instead.
type AccountStreamDirectUploadService struct {
	Options []option.RequestOption
}

// NewAccountStreamDirectUploadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStreamDirectUploadService(opts ...option.RequestOption) (r *AccountStreamDirectUploadService) {
	r = &AccountStreamDirectUploadService{}
	r.Options = opts
	return
}

// Creates a direct upload that allows video uploads without an API key.
func (r *AccountStreamDirectUploadService) StreamVideosUploadVideosViaDirectUploadURLs(ctx context.Context, accountIdentifier string, params AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParams, opts ...option.RequestOption) (res *DirectUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/direct_upload", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type DirectUploadResponse struct {
	Errors   []DirectUploadResponseError   `json:"errors"`
	Messages []DirectUploadResponseMessage `json:"messages"`
	Result   DirectUploadResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DirectUploadResponseSuccess `json:"success"`
	JSON    directUploadResponseJSON    `json:"-"`
}

// directUploadResponseJSON contains the JSON metadata for the struct
// [DirectUploadResponse]
type directUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DirectUploadResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    directUploadResponseErrorJSON `json:"-"`
}

// directUploadResponseErrorJSON contains the JSON metadata for the struct
// [DirectUploadResponseError]
type directUploadResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectUploadResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DirectUploadResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    directUploadResponseMessageJSON `json:"-"`
}

// directUploadResponseMessageJSON contains the JSON metadata for the struct
// [DirectUploadResponseMessage]
type directUploadResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectUploadResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DirectUploadResponseResult struct {
	// A Cloudflare-generated unique identifier for a media item.
	Uid string `json:"uid"`
	// The URL an unauthenticated upload can use for a single
	// `HTTP POST multipart/form-data` request.
	UploadURL string                              `json:"uploadURL"`
	Watermark DirectUploadResponseResultWatermark `json:"watermark"`
	JSON      directUploadResponseResultJSON      `json:"-"`
}

// directUploadResponseResultJSON contains the JSON metadata for the struct
// [DirectUploadResponseResult]
type directUploadResponseResultJSON struct {
	Uid         apijson.Field
	UploadURL   apijson.Field
	Watermark   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DirectUploadResponseResultWatermark struct {
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
	Width int64                                   `json:"width"`
	JSON  directUploadResponseResultWatermarkJSON `json:"-"`
}

// directUploadResponseResultWatermarkJSON contains the JSON metadata for the
// struct [DirectUploadResponseResultWatermark]
type directUploadResponseResultWatermarkJSON struct {
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

func (r *DirectUploadResponseResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DirectUploadResponseSuccess bool

const (
	DirectUploadResponseSuccessTrue DirectUploadResponseSuccess = true
)

type AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParams struct {
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds,required"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The date and time after upload when videos will not be accepted.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                                                                             `json:"thumbnailTimestampPct"`
	Watermark             param.Field[AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParamsWatermark] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
}

func (r AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
