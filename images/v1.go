// File generated from our OpenAPI spec by Stainless.

package images

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// V1Service contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewV1Service] method instead.
type V1Service struct {
	Options  []option.RequestOption
	Keys     *V1KeyService
	Stats    *V1StatService
	Variants *V1VariantService
	Blobs    *V1BlobService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r *V1Service) {
	r = &V1Service{}
	r.Options = opts
	r.Keys = NewV1KeyService(opts...)
	r.Stats = NewV1StatService(opts...)
	r.Variants = NewV1VariantService(opts...)
	r.Blobs = NewV1BlobService(opts...)
	return
}

// Upload an image with up to 10 Megabytes using a single HTTP POST
// (multipart/form-data) request. An image can be uploaded by sending an image file
// or passing an accessible to an API url.
func (r *V1Service) New(ctx context.Context, params V1NewParams, opts ...option.RequestOption) (res *ImagesImage, err error) {
	opts = append(r.Options[:], opts...)
	var env V1NewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
func (r *V1Service) List(ctx context.Context, params V1ListParams, opts ...option.RequestOption) (res *shared.V4PagePagination[V1ListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/images/v1", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
func (r *V1Service) ListAutoPaging(ctx context.Context, params V1ListParams, opts ...option.RequestOption) *shared.V4PagePaginationAutoPager[V1ListResponse] {
	return shared.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Delete an image on Cloudflare Images. On success, all copies of the image are
// deleted and purged from cache.
func (r *V1Service) Delete(ctx context.Context, imageID string, body V1DeleteParams, opts ...option.RequestOption) (res *V1DeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1DeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/%s", body.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update image access control. On access control change, all copies of the image
// are purged from cache.
func (r *V1Service) Edit(ctx context.Context, imageID string, params V1EditParams, opts ...option.RequestOption) (res *ImagesImage, err error) {
	opts = append(r.Options[:], opts...)
	var env V1EditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/%s", params.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single image.
func (r *V1Service) Get(ctx context.Context, imageID string, query V1GetParams, opts ...option.RequestOption) (res *ImagesImage, err error) {
	opts = append(r.Options[:], opts...)
	var env V1GetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/%s", query.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImagesImage struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Meta interface{} `json:"meta"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []ImagesImageVariant `json:"variants" format:"uri"`
	JSON     imagesImageJSON      `json:"-"`
}

// imagesImageJSON contains the JSON metadata for the struct [ImagesImage]
type imagesImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImagesImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagesImageJSON) RawJSON() string {
	return r.raw
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImagesImageVariant interface {
	ImplementsImagesImagesImageVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImagesImageVariant)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V1ListResponse struct {
	Errors   []V1ListResponseError   `json:"errors,required"`
	Messages []V1ListResponseMessage `json:"messages,required"`
	Result   V1ListResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success V1ListResponseSuccess `json:"success,required"`
	JSON    v1ListResponseJSON    `json:"-"`
}

// v1ListResponseJSON contains the JSON metadata for the struct [V1ListResponse]
type v1ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ListResponseJSON) RawJSON() string {
	return r.raw
}

type V1ListResponseError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    v1ListResponseErrorJSON `json:"-"`
}

// v1ListResponseErrorJSON contains the JSON metadata for the struct
// [V1ListResponseError]
type v1ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type V1ListResponseMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    v1ListResponseMessageJSON `json:"-"`
}

// v1ListResponseMessageJSON contains the JSON metadata for the struct
// [V1ListResponseMessage]
type v1ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type V1ListResponseResult struct {
	Images []ImagesImage            `json:"images"`
	JSON   v1ListResponseResultJSON `json:"-"`
}

// v1ListResponseResultJSON contains the JSON metadata for the struct
// [V1ListResponseResult]
type v1ListResponseResultJSON struct {
	Images      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1ListResponseSuccess bool

const (
	V1ListResponseSuccessTrue V1ListResponseSuccess = true
)

// Union satisfied by [images.V1DeleteResponseUnknown] or [shared.UnionString].
type V1DeleteResponse interface {
	ImplementsImagesV1DeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1DeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V1NewParams struct {
	// Account identifier tag.
	AccountID param.Field[string]      `path:"account_id,required"`
	Metadata  param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token for the access.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r V1NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1NewResponseEnvelope struct {
	Errors   []V1NewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1NewResponseEnvelopeMessages `json:"messages,required"`
	Result   ImagesImage                     `json:"result,required"`
	// Whether the API call was successful
	Success V1NewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1NewResponseEnvelopeJSON    `json:"-"`
}

// v1NewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1NewResponseEnvelope]
type v1NewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1NewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1NewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1NewResponseEnvelopeErrors struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    v1NewResponseEnvelopeErrorsJSON `json:"-"`
}

// v1NewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1NewResponseEnvelopeErrors]
type v1NewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1NewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1NewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1NewResponseEnvelopeMessages struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    v1NewResponseEnvelopeMessagesJSON `json:"-"`
}

// v1NewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V1NewResponseEnvelopeMessages]
type v1NewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1NewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1NewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1NewResponseEnvelopeSuccess bool

const (
	V1NewResponseEnvelopeSuccessTrue V1NewResponseEnvelopeSuccess = true
)

type V1ListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [V1ListParams]'s query parameters as `url.Values`.
func (r V1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1DeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1DeleteResponseEnvelope struct {
	Errors   []V1DeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1DeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V1DeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1DeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1DeleteResponseEnvelopeJSON    `json:"-"`
}

// v1DeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1DeleteResponseEnvelope]
type v1DeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1DeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1DeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1DeleteResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    v1DeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v1DeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1DeleteResponseEnvelopeErrors]
type v1DeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1DeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1DeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1DeleteResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    v1DeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v1DeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V1DeleteResponseEnvelopeMessages]
type v1DeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1DeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1DeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1DeleteResponseEnvelopeSuccess bool

const (
	V1DeleteResponseEnvelopeSuccessTrue V1DeleteResponseEnvelopeSuccess = true
)

type V1EditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. No change if not specified.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image can be accessed using only its UID. If set to
	// `true`, a signed token needs to be generated with a signing key to view the
	// image. Returns a new UID on a change. No change if not specified.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r V1EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1EditResponseEnvelope struct {
	Errors   []V1EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1EditResponseEnvelopeMessages `json:"messages,required"`
	Result   ImagesImage                      `json:"result,required"`
	// Whether the API call was successful
	Success V1EditResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1EditResponseEnvelopeJSON    `json:"-"`
}

// v1EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1EditResponseEnvelope]
type v1EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1EditResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    v1EditResponseEnvelopeErrorsJSON `json:"-"`
}

// v1EditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1EditResponseEnvelopeErrors]
type v1EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1EditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1EditResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    v1EditResponseEnvelopeMessagesJSON `json:"-"`
}

// v1EditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V1EditResponseEnvelopeMessages]
type v1EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1EditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1EditResponseEnvelopeSuccess bool

const (
	V1EditResponseEnvelopeSuccessTrue V1EditResponseEnvelopeSuccess = true
)

type V1GetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1GetResponseEnvelope struct {
	Errors   []V1GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1GetResponseEnvelopeMessages `json:"messages,required"`
	Result   ImagesImage                     `json:"result,required"`
	// Whether the API call was successful
	Success V1GetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1GetResponseEnvelopeJSON    `json:"-"`
}

// v1GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1GetResponseEnvelope]
type v1GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1GetResponseEnvelopeErrors struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    v1GetResponseEnvelopeErrorsJSON `json:"-"`
}

// v1GetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1GetResponseEnvelopeErrors]
type v1GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1GetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1GetResponseEnvelopeMessages struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    v1GetResponseEnvelopeMessagesJSON `json:"-"`
}

// v1GetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V1GetResponseEnvelopeMessages]
type v1GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1GetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1GetResponseEnvelopeSuccess bool

const (
	V1GetResponseEnvelopeSuccessTrue V1GetResponseEnvelopeSuccess = true
)
