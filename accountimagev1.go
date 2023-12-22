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

// AccountImageV1Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1Service] method
// instead.
type AccountImageV1Service struct {
	Options  []option.RequestOption
	Keys     *AccountImageV1KeyService
	Stats    *AccountImageV1StatService
	Variants *AccountImageV1VariantService
	Blobs    *AccountImageV1BlobService
}

// NewAccountImageV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountImageV1Service(opts ...option.RequestOption) (r *AccountImageV1Service) {
	r = &AccountImageV1Service{}
	r.Options = opts
	r.Keys = NewAccountImageV1KeyService(opts...)
	r.Stats = NewAccountImageV1StatService(opts...)
	r.Variants = NewAccountImageV1VariantService(opts...)
	r.Blobs = NewAccountImageV1BlobService(opts...)
	return
}

// Fetch details for a single image.
func (r *AccountImageV1Service) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *ImageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update image access control. On access control change, all copies of the image
// are purged from cache.
func (r *AccountImageV1Service) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountImageV1UpdateParams, opts ...option.RequestOption) (res *ImageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete an image on Cloudflare Images. On success, all copies of the image are
// deleted and purged from cache.
func (r *AccountImageV1Service) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *DeletedResponseFkozR92k, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
func (r *AccountImageV1Service) CloudflareImagesListImages(ctx context.Context, accountIdentifier string, query AccountImageV1CloudflareImagesListImagesParams, opts ...option.RequestOption) (res *ImageResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload an image via URL with up to 10 Megabytes using a single HTTP POST
// (multipart/form-data) request.
func (r *AccountImageV1Service) CloudflareImagesUploadAnImageViaURL(ctx context.Context, accountIdentifier string, body AccountImageV1CloudflareImagesUploadAnImageViaURLParams, opts ...option.RequestOption) (res *ImageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ImageResponseCollection struct {
	Errors     []ImageResponseCollectionError    `json:"errors"`
	Messages   []ImageResponseCollectionMessage  `json:"messages"`
	Result     []ImageResponseCollectionResult   `json:"result"`
	ResultInfo ImageResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ImageResponseCollectionSuccess `json:"success"`
	JSON    imageResponseCollectionJSON    `json:"-"`
}

// imageResponseCollectionJSON contains the JSON metadata for the struct
// [ImageResponseCollection]
type imageResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageResponseCollectionError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    imageResponseCollectionErrorJSON `json:"-"`
}

// imageResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ImageResponseCollectionError]
type imageResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageResponseCollectionMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    imageResponseCollectionMessageJSON `json:"-"`
}

// imageResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ImageResponseCollectionMessage]
type imageResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageResponseCollectionResult struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Metadata interface{} `json:"metadata"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []ImageResponseCollectionResultVariant `json:"variants" format:"uri"`
	JSON     imageResponseCollectionResultJSON      `json:"-"`
}

// imageResponseCollectionResultJSON contains the JSON metadata for the struct
// [ImageResponseCollectionResult]
type imageResponseCollectionResultJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Metadata          apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImageResponseCollectionResultVariant interface {
	ImplementsImageResponseCollectionResultVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageResponseCollectionResultVariant)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ImageResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       imageResponseCollectionResultInfoJSON `json:"-"`
}

// imageResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [ImageResponseCollectionResultInfo]
type imageResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageResponseCollectionSuccess bool

const (
	ImageResponseCollectionSuccessTrue ImageResponseCollectionSuccess = true
)

type ImageResponseSingle struct {
	Errors   []ImageResponseSingleError   `json:"errors"`
	Messages []ImageResponseSingleMessage `json:"messages"`
	Result   interface{}                  `json:"result"`
	// Whether the API call was successful
	Success ImageResponseSingleSuccess `json:"success"`
	JSON    imageResponseSingleJSON    `json:"-"`
}

// imageResponseSingleJSON contains the JSON metadata for the struct
// [ImageResponseSingle]
type imageResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageResponseSingleError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    imageResponseSingleErrorJSON `json:"-"`
}

// imageResponseSingleErrorJSON contains the JSON metadata for the struct
// [ImageResponseSingleError]
type imageResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageResponseSingleMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    imageResponseSingleMessageJSON `json:"-"`
}

// imageResponseSingleMessageJSON contains the JSON metadata for the struct
// [ImageResponseSingleMessage]
type imageResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageResponseSingleSuccess bool

const (
	ImageResponseSingleSuccessTrue ImageResponseSingleSuccess = true
)

type DeletedResponseFkozR92k struct {
	Errors   []DeletedResponseFkozR92kError   `json:"errors"`
	Messages []DeletedResponseFkozR92kMessage `json:"messages"`
	Result   interface{}                      `json:"result"`
	// Whether the API call was successful
	Success DeletedResponseFkozR92kSuccess `json:"success"`
	JSON    deletedResponseFkozR92kJSON    `json:"-"`
}

// deletedResponseFkozR92kJSON contains the JSON metadata for the struct
// [DeletedResponseFkozR92k]
type deletedResponseFkozR92kJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedResponseFkozR92k) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeletedResponseFkozR92kError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    deletedResponseFkozR92kErrorJSON `json:"-"`
}

// deletedResponseFkozR92kErrorJSON contains the JSON metadata for the struct
// [DeletedResponseFkozR92kError]
type deletedResponseFkozR92kErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedResponseFkozR92kError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeletedResponseFkozR92kMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    deletedResponseFkozR92kMessageJSON `json:"-"`
}

// deletedResponseFkozR92kMessageJSON contains the JSON metadata for the struct
// [DeletedResponseFkozR92kMessage]
type deletedResponseFkozR92kMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedResponseFkozR92kMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeletedResponseFkozR92kSuccess bool

const (
	DeletedResponseFkozR92kSuccessTrue DeletedResponseFkozR92kSuccess = true
)

type AccountImageV1UpdateParams struct {
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. No change if not specified.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image can be accessed using only its UID. If set to
	// `true`, a signed token needs to be generated with a signing key to view the
	// image. Returns a new UID on a change. No change if not specified.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r AccountImageV1UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountImageV1CloudflareImagesListImagesParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountImageV1CloudflareImagesListImagesParams]'s query
// parameters as `url.Values`.
func (r AccountImageV1CloudflareImagesListImagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountImageV1CloudflareImagesUploadAnImageViaURLParams struct {
	// A URL to fetch an image from origin.
	URL param.Field[string] `json:"url,required"`
	// User modifiable key-value store. Can use used for keeping references to another
	// system of record for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token for the access.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r AccountImageV1CloudflareImagesUploadAnImageViaURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
