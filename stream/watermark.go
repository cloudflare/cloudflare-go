// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// WatermarkService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWatermarkService] method instead.
type WatermarkService struct {
	Options []option.RequestOption
}

// NewWatermarkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWatermarkService(opts ...option.RequestOption) (r *WatermarkService) {
	r = &WatermarkService{}
	r.Options = opts
	return
}

// Creates watermark profiles using a single `HTTP POST multipart/form-data`
// request.
func (r *WatermarkService) New(ctx context.Context, params WatermarkNewParams, opts ...option.RequestOption) (res *Watermark, err error) {
	opts = append(r.Options[:], opts...)
	var env WatermarkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all watermark profiles for an account.
func (r *WatermarkService) List(ctx context.Context, query WatermarkListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Watermark], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all watermark profiles for an account.
func (r *WatermarkService) ListAutoPaging(ctx context.Context, query WatermarkListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Watermark] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a watermark profile.
func (r *WatermarkService) Delete(ctx context.Context, identifier string, body WatermarkDeleteParams, opts ...option.RequestOption) (res *WatermarkDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env WatermarkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves details for a single watermark profile.
func (r *WatermarkService) Get(ctx context.Context, identifier string, query WatermarkGetParams, opts ...option.RequestOption) (res *Watermark, err error) {
	opts = append(r.Options[:], opts...)
	var env WatermarkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Watermark struct {
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
	UID string `json:"uid"`
	// The width of the image in pixels.
	Width int64         `json:"width"`
	JSON  watermarkJSON `json:"-"`
}

// watermarkJSON contains the JSON metadata for the struct [Watermark]
type watermarkJSON struct {
	Created        apijson.Field
	DownloadedFrom apijson.Field
	Height         apijson.Field
	Name           apijson.Field
	Opacity        apijson.Field
	Padding        apijson.Field
	Position       apijson.Field
	Scale          apijson.Field
	Size           apijson.Field
	UID            apijson.Field
	Width          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Watermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarkJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [stream.WatermarkDeleteResponseUnknown] or
// [shared.UnionString].
type WatermarkDeleteResponseUnion interface {
	ImplementsStreamWatermarkDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WatermarkDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WatermarkNewParams struct {
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

func (r WatermarkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatermarkNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success WatermarkNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Watermark                           `json:"result"`
	JSON    watermarkNewResponseEnvelopeJSON    `json:"-"`
}

// watermarkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WatermarkNewResponseEnvelope]
type watermarkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WatermarkNewResponseEnvelopeSuccess bool

const (
	WatermarkNewResponseEnvelopeSuccessTrue WatermarkNewResponseEnvelopeSuccess = true
)

func (r WatermarkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WatermarkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WatermarkListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WatermarkDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WatermarkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success WatermarkDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  WatermarkDeleteResponseUnion           `json:"result"`
	JSON    watermarkDeleteResponseEnvelopeJSON    `json:"-"`
}

// watermarkDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WatermarkDeleteResponseEnvelope]
type watermarkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WatermarkDeleteResponseEnvelopeSuccess bool

const (
	WatermarkDeleteResponseEnvelopeSuccessTrue WatermarkDeleteResponseEnvelopeSuccess = true
)

func (r WatermarkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WatermarkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WatermarkGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WatermarkGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success WatermarkGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Watermark                           `json:"result"`
	JSON    watermarkGetResponseEnvelopeJSON    `json:"-"`
}

// watermarkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WatermarkGetResponseEnvelope]
type watermarkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WatermarkGetResponseEnvelopeSuccess bool

const (
	WatermarkGetResponseEnvelopeSuccessTrue WatermarkGetResponseEnvelopeSuccess = true
)

func (r WatermarkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WatermarkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
