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

// StreamDirectUploadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamDirectUploadService] method
// instead.
type StreamDirectUploadService struct {
	Options []option.RequestOption
}

// NewStreamDirectUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStreamDirectUploadService(opts ...option.RequestOption) (r *StreamDirectUploadService) {
	r = &StreamDirectUploadService{}
	r.Options = opts
	return
}

// Creates a direct upload that allows video uploads without an API key.
func (r *StreamDirectUploadService) New(ctx context.Context, params StreamDirectUploadNewParams, opts ...option.RequestOption) (res *StreamDirectUploadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDirectUploadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/direct_upload", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamDirectUploadNewResponse struct {
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion time.Time `json:"scheduledDeletion" format:"date-time"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid string `json:"uid"`
	// The URL an unauthenticated upload can use for a single
	// `HTTP POST multipart/form-data` request.
	UploadURL string                            `json:"uploadURL"`
	Watermark StreamWatermarks                  `json:"watermark"`
	JSON      streamDirectUploadNewResponseJSON `json:"-"`
}

// streamDirectUploadNewResponseJSON contains the JSON metadata for the struct
// [StreamDirectUploadNewResponse]
type streamDirectUploadNewResponseJSON struct {
	ScheduledDeletion apijson.Field
	Uid               apijson.Field
	UploadURL         apijson.Field
	Watermark         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamDirectUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDirectUploadNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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
	ThumbnailTimestampPct param.Field[float64]                              `json:"thumbnailTimestampPct"`
	Watermark             param.Field[StreamDirectUploadNewParamsWatermark] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
}

func (r StreamDirectUploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamDirectUploadNewParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r StreamDirectUploadNewParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamDirectUploadNewResponseEnvelope struct {
	Errors   []StreamDirectUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDirectUploadNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDirectUploadNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDirectUploadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDirectUploadNewResponseEnvelopeJSON    `json:"-"`
}

// streamDirectUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamDirectUploadNewResponseEnvelope]
type streamDirectUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDirectUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDirectUploadNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamDirectUploadNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDirectUploadNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamDirectUploadNewResponseEnvelopeErrors]
type streamDirectUploadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDirectUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDirectUploadNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    streamDirectUploadNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDirectUploadNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamDirectUploadNewResponseEnvelopeMessages]
type streamDirectUploadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDirectUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDirectUploadNewResponseEnvelopeSuccess bool

const (
	StreamDirectUploadNewResponseEnvelopeSuccessTrue StreamDirectUploadNewResponseEnvelopeSuccess = true
)
