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

// AccountStreamClipService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamClipService] method
// instead.
type AccountStreamClipService struct {
	Options []option.RequestOption
}

// NewAccountStreamClipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamClipService(opts ...option.RequestOption) (r *AccountStreamClipService) {
	r = &AccountStreamClipService{}
	r.Options = opts
	return
}

// Clips a video based on the specified start and end times provided in seconds.
func (r *AccountStreamClipService) StreamVideoClippingClipVideosGivenAStartAndEndTime(ctx context.Context, accountIdentifier string, body AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParams, opts ...option.RequestOption) (res *ClipResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/clip", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ClipResponseSingle struct {
	Errors     []ClipResponseSingleError    `json:"errors"`
	Messages   []ClipResponseSingleMessage  `json:"messages"`
	Result     ClipResponseSingleResult     `json:"result"`
	ResultInfo ClipResponseSingleResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ClipResponseSingleSuccess `json:"success"`
	JSON    clipResponseSingleJSON    `json:"-"`
}

// clipResponseSingleJSON contains the JSON metadata for the struct
// [ClipResponseSingle]
type clipResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClipResponseSingleError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    clipResponseSingleErrorJSON `json:"-"`
}

// clipResponseSingleErrorJSON contains the JSON metadata for the struct
// [ClipResponseSingleError]
type clipResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClipResponseSingleMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    clipResponseSingleMessageJSON `json:"-"`
}

// clipResponseSingleMessageJSON contains the JSON metadata for the struct
// [ClipResponseSingleMessage]
type clipResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClipResponseSingleResult struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// The unique video identifier (UID).
	ClippedFromVideoUid string `json:"clippedFromVideoUID"`
	// The date and time the clip was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// Specifies the end time for the video clip in seconds.
	EndTimeSeconds int64 `json:"endTimeSeconds"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds int64 `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time                        `json:"modified" format:"date-time"`
	Playback ClipResponseSingleResultPlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds int64 `json:"startTimeSeconds"`
	// Specifies the processing status of the video.
	Status ClipResponseSingleResultStatus `json:"status"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64                           `json:"thumbnailTimestampPct"`
	Watermark             ClipResponseSingleResultWatermark `json:"watermark"`
	JSON                  clipResponseSingleResultJSON      `json:"-"`
}

// clipResponseSingleResultJSON contains the JSON metadata for the struct
// [ClipResponseSingleResult]
type clipResponseSingleResultJSON struct {
	AllowedOrigins        apijson.Field
	ClippedFromVideoUid   apijson.Field
	Created               apijson.Field
	Creator               apijson.Field
	EndTimeSeconds        apijson.Field
	MaxDurationSeconds    apijson.Field
	Meta                  apijson.Field
	Modified              apijson.Field
	Playback              apijson.Field
	Preview               apijson.Field
	RequireSignedURLs     apijson.Field
	StartTimeSeconds      apijson.Field
	Status                apijson.Field
	ThumbnailTimestampPct apijson.Field
	Watermark             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ClipResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClipResponseSingleResultPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                               `json:"hls"`
	JSON clipResponseSingleResultPlaybackJSON `json:"-"`
}

// clipResponseSingleResultPlaybackJSON contains the JSON metadata for the struct
// [ClipResponseSingleResultPlayback]
type clipResponseSingleResultPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipResponseSingleResultPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type ClipResponseSingleResultStatus string

const (
	ClipResponseSingleResultStatusPendingupload ClipResponseSingleResultStatus = "pendingupload"
	ClipResponseSingleResultStatusDownloading   ClipResponseSingleResultStatus = "downloading"
	ClipResponseSingleResultStatusQueued        ClipResponseSingleResultStatus = "queued"
	ClipResponseSingleResultStatusInprogress    ClipResponseSingleResultStatus = "inprogress"
	ClipResponseSingleResultStatusReady         ClipResponseSingleResultStatus = "ready"
	ClipResponseSingleResultStatusError         ClipResponseSingleResultStatus = "error"
)

type ClipResponseSingleResultWatermark struct {
	// The unique identifier for the watermark profile.
	Uid  string                                `json:"uid"`
	JSON clipResponseSingleResultWatermarkJSON `json:"-"`
}

// clipResponseSingleResultWatermarkJSON contains the JSON metadata for the struct
// [ClipResponseSingleResultWatermark]
type clipResponseSingleResultWatermarkJSON struct {
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipResponseSingleResultWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClipResponseSingleResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                          `json:"total_count"`
	JSON       clipResponseSingleResultInfoJSON `json:"-"`
}

// clipResponseSingleResultInfoJSON contains the JSON metadata for the struct
// [ClipResponseSingleResultInfo]
type clipResponseSingleResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipResponseSingleResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ClipResponseSingleSuccess bool

const (
	ClipResponseSingleSuccessTrue ClipResponseSingleSuccess = true
)

type AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParams struct {
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
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                                                                            `json:"thumbnailTimestampPct"`
	Watermark             param.Field[AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParamsWatermark] `json:"watermark"`
}

func (r AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r AccountStreamClipStreamVideoClippingClipVideosGivenAStartAndEndTimeParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
