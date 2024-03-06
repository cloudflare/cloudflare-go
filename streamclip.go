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

// StreamClipService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewStreamClipService] method instead.
type StreamClipService struct {
	Options []option.RequestOption
}

// NewStreamClipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamClipService(opts ...option.RequestOption) (r *StreamClipService) {
	r = &StreamClipService{}
	r.Options = opts
	return
}

// Clips a video based on the specified start and end times provided in seconds.
func (r *StreamClipService) New(ctx context.Context, params StreamClipNewParams, opts ...option.RequestOption) (res *StreamClipNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamClipNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/clip", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamClipNewResponse struct {
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
	Modified time.Time                     `json:"modified" format:"date-time"`
	Playback StreamClipNewResponsePlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds int64 `json:"startTimeSeconds"`
	// Specifies the processing status for all quality levels for a video.
	Status StreamClipNewResponseStatus `json:"status"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64                        `json:"thumbnailTimestampPct"`
	Watermark             StreamClipNewResponseWatermark `json:"watermark"`
	JSON                  streamClipNewResponseJSON      `json:"-"`
}

// streamClipNewResponseJSON contains the JSON metadata for the struct
// [StreamClipNewResponse]
type streamClipNewResponseJSON struct {
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

func (r *StreamClipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamClipNewResponsePlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                            `json:"hls"`
	JSON streamClipNewResponsePlaybackJSON `json:"-"`
}

// streamClipNewResponsePlaybackJSON contains the JSON metadata for the struct
// [StreamClipNewResponsePlayback]
type streamClipNewResponsePlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClipNewResponsePlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status for all quality levels for a video.
type StreamClipNewResponseStatus string

const (
	StreamClipNewResponseStatusPendingupload StreamClipNewResponseStatus = "pendingupload"
	StreamClipNewResponseStatusDownloading   StreamClipNewResponseStatus = "downloading"
	StreamClipNewResponseStatusQueued        StreamClipNewResponseStatus = "queued"
	StreamClipNewResponseStatusInprogress    StreamClipNewResponseStatus = "inprogress"
	StreamClipNewResponseStatusReady         StreamClipNewResponseStatus = "ready"
	StreamClipNewResponseStatusError         StreamClipNewResponseStatus = "error"
)

type StreamClipNewResponseWatermark struct {
	// The unique identifier for the watermark profile.
	Uid  string                             `json:"uid"`
	JSON streamClipNewResponseWatermarkJSON `json:"-"`
}

// streamClipNewResponseWatermarkJSON contains the JSON metadata for the struct
// [StreamClipNewResponseWatermark]
type streamClipNewResponseWatermarkJSON struct {
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClipNewResponseWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamClipNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique video identifier (UID).
	ClippedFromVideoUid param.Field[string] `json:"clippedFromVideoUID,required"`
	// Specifies the end time for the video clip in seconds.
	EndTimeSeconds param.Field[int64] `json:"endTimeSeconds,required"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds param.Field[int64] `json:"startTimeSeconds,required"`
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
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                      `json:"thumbnailTimestampPct"`
	Watermark             param.Field[StreamClipNewParamsWatermark] `json:"watermark"`
}

func (r StreamClipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamClipNewParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r StreamClipNewParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamClipNewResponseEnvelope struct {
	Errors   []StreamClipNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamClipNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamClipNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamClipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamClipNewResponseEnvelopeJSON    `json:"-"`
}

// streamClipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamClipNewResponseEnvelope]
type streamClipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamClipNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    streamClipNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamClipNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamClipNewResponseEnvelopeErrors]
type streamClipNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClipNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamClipNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    streamClipNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamClipNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamClipNewResponseEnvelopeMessages]
type streamClipNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClipNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamClipNewResponseEnvelopeSuccess bool

const (
	StreamClipNewResponseEnvelopeSuccessTrue StreamClipNewResponseEnvelopeSuccess = true
)
