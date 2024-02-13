// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// StreamWatermarkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamWatermarkService] method
// instead.
type StreamWatermarkService struct {
	Options []option.RequestOption
}

// NewStreamWatermarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamWatermarkService(opts ...option.RequestOption) (r *StreamWatermarkService) {
	r = &StreamWatermarkService{}
	r.Options = opts
	return
}

// Deletes a watermark profile.
func (r *StreamWatermarkService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamWatermarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves details for a single watermark profile.
func (r *StreamWatermarkService) Get(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamWatermarkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates watermark profiles using a single `HTTP POST multipart/form-data`
// request.
func (r *StreamWatermarkService) StreamWatermarkProfileNewWatermarkProfilesViaBasicUpload(ctx context.Context, accountID string, body StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams, opts ...option.RequestOption) (res *StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all watermark profiles for an account.
func (r *StreamWatermarkService) StreamWatermarkProfileListWatermarkProfiles(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StreamWatermarkDeleteResponseUnknown] or
// [shared.UnionString].
type StreamWatermarkDeleteResponse interface {
	ImplementsStreamWatermarkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWatermarkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamWatermarkGetResponseUnknown] or [shared.UnionString].
type StreamWatermarkGetResponse interface {
	ImplementsStreamWatermarkGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWatermarkGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseUnknown]
// or [shared.UnionString].
type StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse interface {
	ImplementsStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse struct {
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
	Width int64                                                                  `json:"width"`
	JSON  streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseJSON `json:"-"`
}

// streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseJSON contains
// the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse]
type streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseJSON struct {
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

func (r *StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkDeleteResponseEnvelope struct {
	Errors   []StreamWatermarkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWatermarkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWatermarkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWatermarkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWatermarkDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamWatermarkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamWatermarkDeleteResponseEnvelope]
type streamWatermarkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamWatermarkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWatermarkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamWatermarkDeleteResponseEnvelopeErrors]
type streamWatermarkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    streamWatermarkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWatermarkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWatermarkDeleteResponseEnvelopeMessages]
type streamWatermarkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWatermarkDeleteResponseEnvelopeSuccess bool

const (
	StreamWatermarkDeleteResponseEnvelopeSuccessTrue StreamWatermarkDeleteResponseEnvelopeSuccess = true
)

type StreamWatermarkGetResponseEnvelope struct {
	Errors   []StreamWatermarkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWatermarkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWatermarkGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWatermarkGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWatermarkGetResponseEnvelopeJSON    `json:"-"`
}

// streamWatermarkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamWatermarkGetResponseEnvelope]
type streamWatermarkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamWatermarkGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWatermarkGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamWatermarkGetResponseEnvelopeErrors]
type streamWatermarkGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamWatermarkGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWatermarkGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWatermarkGetResponseEnvelopeMessages]
type streamWatermarkGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWatermarkGetResponseEnvelopeSuccess bool

const (
	StreamWatermarkGetResponseEnvelopeSuccessTrue StreamWatermarkGetResponseEnvelopeSuccess = true
)

type StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams struct {
	// The image file to upload.
	File param.Field[string] `json:"file,required"`
	// A short description of the watermark profile.
	Name param.Field[string] `json:"name"`
	// The translucency of the image. A value of `0.0` makes the image completely
	// transparent, and `1.0` makes the image completely opaque. Note that if the image
	// is already semi-transparent, setting this to `1.0` will not make the image
	// completely opaque.
	Opacity param.Field[float64] `json:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video
	// and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded
	// video width or length, as determined by the algorithm.
	Padding param.Field[float64] `json:"padding"`
	// The location of the image. Valid positions are: `upperRight`, `upperLeft`,
	// `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the
	// `padding` parameter.
	Position param.Field[string] `json:"position"`
	// The size of the image relative to the overall size of the video. This parameter
	// will adapt to horizontal and vertical videos automatically. `0.0` indicates no
	// scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	Scale param.Field[float64] `json:"scale"`
}

func (r StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelope struct {
	Errors   []StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeJSON    `json:"-"`
}

// streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelope]
type streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrors struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrors]
type streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessages struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessages]
type streamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeSuccess bool

const (
	StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeSuccessTrue StreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseEnvelopeSuccess = true
)

type StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelope struct {
	Errors   []StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeJSON    `json:"-"`
}

// streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelope]
type streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrors]
type streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessages]
type streamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeSuccess bool

const (
	StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeSuccessTrue StreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseEnvelopeSuccess = true
)
