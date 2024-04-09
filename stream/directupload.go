// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DirectUploadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDirectUploadService] method
// instead.
type DirectUploadService struct {
	Options []option.RequestOption
}

// NewDirectUploadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirectUploadService(opts ...option.RequestOption) (r *DirectUploadService) {
	r = &DirectUploadService{}
	r.Options = opts
	return
}

// Creates a direct upload that allows video uploads without an API key.
func (r *DirectUploadService) New(ctx context.Context, params DirectUploadNewParams, opts ...option.RequestOption) (res *DirectUploadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DirectUploadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/direct_upload", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DirectUploadNewResponse struct {
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion time.Time `json:"scheduledDeletion" format:"date-time"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid string `json:"uid"`
	// The URL an unauthenticated upload can use for a single
	// `HTTP POST multipart/form-data` request.
	UploadURL string                      `json:"uploadURL"`
	Watermark Watermaks                   `json:"watermark"`
	JSON      directUploadNewResponseJSON `json:"-"`
}

// directUploadNewResponseJSON contains the JSON metadata for the struct
// [DirectUploadNewResponse]
type directUploadNewResponseJSON struct {
	ScheduledDeletion apijson.Field
	Uid               apijson.Field
	UploadURL         apijson.Field
	Watermark         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DirectUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type DirectUploadNewParams struct {
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
	AllowedOrigins param.Field[[]AllowedOriginsItemParam] `json:"allowedOrigins"`
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
	ThumbnailTimestampPct param.Field[float64]                        `json:"thumbnailTimestampPct"`
	Watermark             param.Field[DirectUploadNewParamsWatermark] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
}

func (r DirectUploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectUploadNewParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r DirectUploadNewParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectUploadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   DirectUploadNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success DirectUploadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    directUploadNewResponseEnvelopeJSON    `json:"-"`
}

// directUploadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DirectUploadNewResponseEnvelope]
type directUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DirectUploadNewResponseEnvelopeSuccess bool

const (
	DirectUploadNewResponseEnvelopeSuccessTrue DirectUploadNewResponseEnvelopeSuccess = true
)

func (r DirectUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DirectUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
