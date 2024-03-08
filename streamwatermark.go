// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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

// Creates watermark profiles using a single `HTTP POST multipart/form-data`
// request.
func (r *StreamWatermarkService) New(ctx context.Context, params StreamWatermarkNewParams, opts ...option.RequestOption) (res *StreamWatermarkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all watermark profiles for an account.
func (r *StreamWatermarkService) List(ctx context.Context, query StreamWatermarkListParams, opts ...option.RequestOption) (res *[]StreamWatermarkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a watermark profile.
func (r *StreamWatermarkService) Delete(ctx context.Context, identifier string, body StreamWatermarkDeleteParams, opts ...option.RequestOption) (res *StreamWatermarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves details for a single watermark profile.
func (r *StreamWatermarkService) Get(ctx context.Context, identifier string, query StreamWatermarkGetParams, opts ...option.RequestOption) (res *StreamWatermarkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWatermarkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StreamWatermarkNewResponseUnknown] or [shared.UnionString].
type StreamWatermarkNewResponse interface {
	ImplementsStreamWatermarkNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWatermarkNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamWatermarkListResponse struct {
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
	Width int64                           `json:"width"`
	JSON  streamWatermarkListResponseJSON `json:"-"`
}

// streamWatermarkListResponseJSON contains the JSON metadata for the struct
// [StreamWatermarkListResponse]
type streamWatermarkListResponseJSON struct {
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

func (r *StreamWatermarkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkListResponseJSON) RawJSON() string {
	return r.raw
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

type StreamWatermarkNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r StreamWatermarkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamWatermarkNewResponseEnvelope struct {
	Errors   []StreamWatermarkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWatermarkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWatermarkNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWatermarkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWatermarkNewResponseEnvelopeJSON    `json:"-"`
}

// streamWatermarkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamWatermarkNewResponseEnvelope]
type streamWatermarkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamWatermarkNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamWatermarkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWatermarkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamWatermarkNewResponseEnvelopeErrors]
type streamWatermarkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamWatermarkNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamWatermarkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWatermarkNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWatermarkNewResponseEnvelopeMessages]
type streamWatermarkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWatermarkNewResponseEnvelopeSuccess bool

const (
	StreamWatermarkNewResponseEnvelopeSuccessTrue StreamWatermarkNewResponseEnvelopeSuccess = true
)

type StreamWatermarkListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamWatermarkListResponseEnvelope struct {
	Errors   []StreamWatermarkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWatermarkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamWatermarkListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamWatermarkListResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWatermarkListResponseEnvelopeJSON    `json:"-"`
}

// streamWatermarkListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamWatermarkListResponseEnvelope]
type streamWatermarkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamWatermarkListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamWatermarkListResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWatermarkListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamWatermarkListResponseEnvelopeErrors]
type streamWatermarkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamWatermarkListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamWatermarkListResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWatermarkListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWatermarkListResponseEnvelopeMessages]
type streamWatermarkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWatermarkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWatermarkListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWatermarkListResponseEnvelopeSuccess bool

const (
	StreamWatermarkListResponseEnvelopeSuccessTrue StreamWatermarkListResponseEnvelopeSuccess = true
)

type StreamWatermarkDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r streamWatermarkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r streamWatermarkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r streamWatermarkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWatermarkDeleteResponseEnvelopeSuccess bool

const (
	StreamWatermarkDeleteResponseEnvelopeSuccessTrue StreamWatermarkDeleteResponseEnvelopeSuccess = true
)

type StreamWatermarkGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

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

func (r streamWatermarkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r streamWatermarkGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r streamWatermarkGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWatermarkGetResponseEnvelopeSuccess bool

const (
	StreamWatermarkGetResponseEnvelopeSuccessTrue StreamWatermarkGetResponseEnvelopeSuccess = true
)
