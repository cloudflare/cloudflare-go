// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ImageV1Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewImageV1Service] method instead.
type ImageV1Service struct {
	Options  []option.RequestOption
	Keys     *ImageV1KeyService
	Stats    *ImageV1StatService
	Variants *ImageV1VariantService
	Blobs    *ImageV1BlobService
}

// NewImageV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageV1Service(opts ...option.RequestOption) (r *ImageV1Service) {
	r = &ImageV1Service{}
	r.Options = opts
	r.Keys = NewImageV1KeyService(opts...)
	r.Stats = NewImageV1StatService(opts...)
	r.Variants = NewImageV1VariantService(opts...)
	r.Blobs = NewImageV1BlobService(opts...)
	return
}

// Upload an image with up to 10 Megabytes using a single HTTP POST
// (multipart/form-data) request. An image can be uploaded by sending an image file
// or passing an accessible to an API url.
func (r *ImageV1Service) New(ctx context.Context, accountID string, body ImageV1NewParams, opts ...option.RequestOption) (res *ImageV1NewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1NewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update image access control. On access control change, all copies of the image
// are purged from cache.
func (r *ImageV1Service) Update(ctx context.Context, accountID string, imageID string, body ImageV1UpdateParams, opts ...option.RequestOption) (res *ImageV1UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1UpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
func (r *ImageV1Service) List(ctx context.Context, accountID string, query ImageV1ListParams, opts ...option.RequestOption) (res *shared.V4PagePagination[ImageV1ListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/images/v1", accountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *ImageV1Service) ListAutoPaging(ctx context.Context, accountID string, query ImageV1ListParams, opts ...option.RequestOption) *shared.V4PagePaginationAutoPager[ImageV1ListResponse] {
	return shared.NewV4PagePaginationAutoPager(r.List(ctx, accountID, query, opts...))
}

// Delete an image on Cloudflare Images. On success, all copies of the image are
// deleted and purged from cache.
func (r *ImageV1Service) Delete(ctx context.Context, accountID string, imageID string, opts ...option.RequestOption) (res *ImageV1DeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1DeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single image.
func (r *ImageV1Service) Get(ctx context.Context, accountID string, imageID string, opts ...option.RequestOption) (res *ImageV1GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1GetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV1NewResponse struct {
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
	Variants []ImageV1NewResponseVariant `json:"variants" format:"uri"`
	JSON     imageV1NewResponseJSON      `json:"-"`
}

// imageV1NewResponseJSON contains the JSON metadata for the struct
// [ImageV1NewResponse]
type imageV1NewResponseJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageV1NewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImageV1NewResponseVariant interface {
	ImplementsImageV1NewResponseVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV1NewResponseVariant)(nil)).Elem(),
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

type ImageV1UpdateResponse struct {
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
	Variants []ImageV1UpdateResponseVariant `json:"variants" format:"uri"`
	JSON     imageV1UpdateResponseJSON      `json:"-"`
}

// imageV1UpdateResponseJSON contains the JSON metadata for the struct
// [ImageV1UpdateResponse]
type imageV1UpdateResponseJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageV1UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImageV1UpdateResponseVariant interface {
	ImplementsImageV1UpdateResponseVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV1UpdateResponseVariant)(nil)).Elem(),
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

type ImageV1ListResponse struct {
	Errors   []ImageV1ListResponseError   `json:"errors,required"`
	Messages []ImageV1ListResponseMessage `json:"messages,required"`
	Result   ImageV1ListResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1ListResponseSuccess `json:"success,required"`
	JSON    imageV1ListResponseJSON    `json:"-"`
}

// imageV1ListResponseJSON contains the JSON metadata for the struct
// [ImageV1ListResponse]
type imageV1ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1ListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    imageV1ListResponseErrorJSON `json:"-"`
}

// imageV1ListResponseErrorJSON contains the JSON metadata for the struct
// [ImageV1ListResponseError]
type imageV1ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1ListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    imageV1ListResponseMessageJSON `json:"-"`
}

// imageV1ListResponseMessageJSON contains the JSON metadata for the struct
// [ImageV1ListResponseMessage]
type imageV1ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1ListResponseResult struct {
	Images []ImageV1ListResponseResultImage `json:"images"`
	JSON   imageV1ListResponseResultJSON    `json:"-"`
}

// imageV1ListResponseResultJSON contains the JSON metadata for the struct
// [ImageV1ListResponseResult]
type imageV1ListResponseResultJSON struct {
	Images      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1ListResponseResultImage struct {
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
	Variants []ImageV1ListResponseResultImagesVariant `json:"variants" format:"uri"`
	JSON     imageV1ListResponseResultImageJSON       `json:"-"`
}

// imageV1ListResponseResultImageJSON contains the JSON metadata for the struct
// [ImageV1ListResponseResultImage]
type imageV1ListResponseResultImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageV1ListResponseResultImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImageV1ListResponseResultImagesVariant interface {
	ImplementsImageV1ListResponseResultImagesVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV1ListResponseResultImagesVariant)(nil)).Elem(),
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

// Whether the API call was successful
type ImageV1ListResponseSuccess bool

const (
	ImageV1ListResponseSuccessTrue ImageV1ListResponseSuccess = true
)

// Union satisfied by [ImageV1DeleteResponseUnknown] or [shared.UnionString].
type ImageV1DeleteResponse interface {
	ImplementsImageV1DeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV1DeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ImageV1GetResponse struct {
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
	Variants []ImageV1GetResponseVariant `json:"variants" format:"uri"`
	JSON     imageV1GetResponseJSON      `json:"-"`
}

// imageV1GetResponseJSON contains the JSON metadata for the struct
// [ImageV1GetResponse]
type imageV1GetResponseJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageV1GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImageV1GetResponseVariant interface {
	ImplementsImageV1GetResponseVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV1GetResponseVariant)(nil)).Elem(),
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

// This interface is a union satisfied by one of the following:
// [ImageV1NewParamsImagesImageUploadViaFile],
// [ImageV1NewParamsImagesImageUploadViaURL].
type ImageV1NewParams interface {
	ImplementsImageV1NewParams()
}

type ImageV1NewParamsImagesImageUploadViaFile struct {
	// An image binary data.
	File param.Field[interface{}] `json:"file,required"`
}

func (r ImageV1NewParamsImagesImageUploadViaFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ImageV1NewParamsImagesImageUploadViaFile) ImplementsImageV1NewParams() {

}

type ImageV1NewParamsImagesImageUploadViaURL struct {
	// A URL to fetch an image from origin.
	URL param.Field[string] `json:"url,required"`
}

func (r ImageV1NewParamsImagesImageUploadViaURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ImageV1NewParamsImagesImageUploadViaURL) ImplementsImageV1NewParams() {

}

type ImageV1NewResponseEnvelope struct {
	Errors   []ImageV1NewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1NewResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1NewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1NewResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1NewResponseEnvelopeJSON    `json:"-"`
}

// imageV1NewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1NewResponseEnvelope]
type imageV1NewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1NewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1NewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    imageV1NewResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1NewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ImageV1NewResponseEnvelopeErrors]
type imageV1NewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1NewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1NewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    imageV1NewResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1NewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ImageV1NewResponseEnvelopeMessages]
type imageV1NewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1NewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1NewResponseEnvelopeSuccess bool

const (
	ImageV1NewResponseEnvelopeSuccessTrue ImageV1NewResponseEnvelopeSuccess = true
)

type ImageV1UpdateParams struct {
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. No change if not specified.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image can be accessed using only its UID. If set to
	// `true`, a signed token needs to be generated with a signing key to view the
	// image. Returns a new UID on a change. No change if not specified.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r ImageV1UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageV1UpdateResponseEnvelope struct {
	Errors   []ImageV1UpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1UpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1UpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1UpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1UpdateResponseEnvelopeJSON    `json:"-"`
}

// imageV1UpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1UpdateResponseEnvelope]
type imageV1UpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1UpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1UpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    imageV1UpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1UpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1UpdateResponseEnvelopeErrors]
type imageV1UpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1UpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1UpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    imageV1UpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1UpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImageV1UpdateResponseEnvelopeMessages]
type imageV1UpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1UpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1UpdateResponseEnvelopeSuccess bool

const (
	ImageV1UpdateResponseEnvelopeSuccessTrue ImageV1UpdateResponseEnvelopeSuccess = true
)

type ImageV1ListParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ImageV1ListParams]'s query parameters as `url.Values`.
func (r ImageV1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImageV1DeleteResponseEnvelope struct {
	Errors   []ImageV1DeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1DeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1DeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1DeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1DeleteResponseEnvelopeJSON    `json:"-"`
}

// imageV1DeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1DeleteResponseEnvelope]
type imageV1DeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1DeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1DeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    imageV1DeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1DeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1DeleteResponseEnvelopeErrors]
type imageV1DeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1DeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1DeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    imageV1DeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1DeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImageV1DeleteResponseEnvelopeMessages]
type imageV1DeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1DeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1DeleteResponseEnvelopeSuccess bool

const (
	ImageV1DeleteResponseEnvelopeSuccessTrue ImageV1DeleteResponseEnvelopeSuccess = true
)

type ImageV1GetResponseEnvelope struct {
	Errors   []ImageV1GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1GetResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1GetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1GetResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1GetResponseEnvelopeJSON    `json:"-"`
}

// imageV1GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1GetResponseEnvelope]
type imageV1GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1GetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    imageV1GetResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1GetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ImageV1GetResponseEnvelopeErrors]
type imageV1GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1GetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    imageV1GetResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1GetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ImageV1GetResponseEnvelopeMessages]
type imageV1GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1GetResponseEnvelopeSuccess bool

const (
	ImageV1GetResponseEnvelopeSuccessTrue ImageV1GetResponseEnvelopeSuccess = true
)
