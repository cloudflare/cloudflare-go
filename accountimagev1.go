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
func (r *AccountImageV1Service) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountImageV1GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update image access control. On access control change, all copies of the image
// are purged from cache.
func (r *AccountImageV1Service) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountImageV1UpdateParams, opts ...option.RequestOption) (res *AccountImageV1UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete an image on Cloudflare Images. On success, all copies of the image are
// deleted and purged from cache.
func (r *AccountImageV1Service) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountImageV1DeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
func (r *AccountImageV1Service) CloudflareImagesListImages(ctx context.Context, accountIdentifier string, query AccountImageV1CloudflareImagesListImagesParams, opts ...option.RequestOption) (res *AccountImageV1CloudflareImagesListImagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload an image with up to 10 Megabytes using a single HTTP POST
// (multipart/form-data) request. An image can be uploaded by sending an image file
// or passing an accessible to an API url.
func (r *AccountImageV1Service) CloudflareImagesUploadAnImageViaURL(ctx context.Context, accountIdentifier string, body AccountImageV1CloudflareImagesUploadAnImageViaURLParams, opts ...option.RequestOption) (res *AccountImageV1CloudflareImagesUploadAnImageViaURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountImageV1GetResponse struct {
	Errors   []AccountImageV1GetResponseError   `json:"errors"`
	Messages []AccountImageV1GetResponseMessage `json:"messages"`
	Result   AccountImageV1GetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1GetResponseSuccess `json:"success"`
	JSON    accountImageV1GetResponseJSON    `json:"-"`
}

// accountImageV1GetResponseJSON contains the JSON metadata for the struct
// [AccountImageV1GetResponse]
type accountImageV1GetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1GetResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accountImageV1GetResponseErrorJSON `json:"-"`
}

// accountImageV1GetResponseErrorJSON contains the JSON metadata for the struct
// [AccountImageV1GetResponseError]
type accountImageV1GetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1GetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1GetResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountImageV1GetResponseMessageJSON `json:"-"`
}

// accountImageV1GetResponseMessageJSON contains the JSON metadata for the struct
// [AccountImageV1GetResponseMessage]
type accountImageV1GetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1GetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1GetResponseResult struct {
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
	Variants []AccountImageV1GetResponseResultVariant `json:"variants" format:"uri"`
	JSON     accountImageV1GetResponseResultJSON      `json:"-"`
}

// accountImageV1GetResponseResultJSON contains the JSON metadata for the struct
// [AccountImageV1GetResponseResult]
type accountImageV1GetResponseResultJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountImageV1GetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type AccountImageV1GetResponseResultVariant interface {
	ImplementsAccountImageV1GetResponseResultVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountImageV1GetResponseResultVariant)(nil)).Elem(),
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

// Whether the API call was successful
type AccountImageV1GetResponseSuccess bool

const (
	AccountImageV1GetResponseSuccessTrue AccountImageV1GetResponseSuccess = true
)

type AccountImageV1UpdateResponse struct {
	Errors   []AccountImageV1UpdateResponseError   `json:"errors"`
	Messages []AccountImageV1UpdateResponseMessage `json:"messages"`
	Result   AccountImageV1UpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1UpdateResponseSuccess `json:"success"`
	JSON    accountImageV1UpdateResponseJSON    `json:"-"`
}

// accountImageV1UpdateResponseJSON contains the JSON metadata for the struct
// [AccountImageV1UpdateResponse]
type accountImageV1UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1UpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountImageV1UpdateResponseErrorJSON `json:"-"`
}

// accountImageV1UpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountImageV1UpdateResponseError]
type accountImageV1UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1UpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountImageV1UpdateResponseMessageJSON `json:"-"`
}

// accountImageV1UpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountImageV1UpdateResponseMessage]
type accountImageV1UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1UpdateResponseResult struct {
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
	Variants []AccountImageV1UpdateResponseResultVariant `json:"variants" format:"uri"`
	JSON     accountImageV1UpdateResponseResultJSON      `json:"-"`
}

// accountImageV1UpdateResponseResultJSON contains the JSON metadata for the struct
// [AccountImageV1UpdateResponseResult]
type accountImageV1UpdateResponseResultJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountImageV1UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type AccountImageV1UpdateResponseResultVariant interface {
	ImplementsAccountImageV1UpdateResponseResultVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountImageV1UpdateResponseResultVariant)(nil)).Elem(),
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

// Whether the API call was successful
type AccountImageV1UpdateResponseSuccess bool

const (
	AccountImageV1UpdateResponseSuccessTrue AccountImageV1UpdateResponseSuccess = true
)

type AccountImageV1DeleteResponse struct {
	Errors   []AccountImageV1DeleteResponseError   `json:"errors"`
	Messages []AccountImageV1DeleteResponseMessage `json:"messages"`
	Result   interface{}                           `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1DeleteResponseSuccess `json:"success"`
	JSON    accountImageV1DeleteResponseJSON    `json:"-"`
}

// accountImageV1DeleteResponseJSON contains the JSON metadata for the struct
// [AccountImageV1DeleteResponse]
type accountImageV1DeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1DeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1DeleteResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountImageV1DeleteResponseErrorJSON `json:"-"`
}

// accountImageV1DeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountImageV1DeleteResponseError]
type accountImageV1DeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1DeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1DeleteResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountImageV1DeleteResponseMessageJSON `json:"-"`
}

// accountImageV1DeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountImageV1DeleteResponseMessage]
type accountImageV1DeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1DeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountImageV1DeleteResponseSuccess bool

const (
	AccountImageV1DeleteResponseSuccessTrue AccountImageV1DeleteResponseSuccess = true
)

type AccountImageV1CloudflareImagesListImagesResponse struct {
	Errors   []AccountImageV1CloudflareImagesListImagesResponseError   `json:"errors"`
	Messages []AccountImageV1CloudflareImagesListImagesResponseMessage `json:"messages"`
	Result   AccountImageV1CloudflareImagesListImagesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1CloudflareImagesListImagesResponseSuccess `json:"success"`
	JSON    accountImageV1CloudflareImagesListImagesResponseJSON    `json:"-"`
}

// accountImageV1CloudflareImagesListImagesResponseJSON contains the JSON metadata
// for the struct [AccountImageV1CloudflareImagesListImagesResponse]
type accountImageV1CloudflareImagesListImagesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesListImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesListImagesResponseError struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountImageV1CloudflareImagesListImagesResponseErrorJSON `json:"-"`
}

// accountImageV1CloudflareImagesListImagesResponseErrorJSON contains the JSON
// metadata for the struct [AccountImageV1CloudflareImagesListImagesResponseError]
type accountImageV1CloudflareImagesListImagesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesListImagesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesListImagesResponseMessage struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountImageV1CloudflareImagesListImagesResponseMessageJSON `json:"-"`
}

// accountImageV1CloudflareImagesListImagesResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountImageV1CloudflareImagesListImagesResponseMessage]
type accountImageV1CloudflareImagesListImagesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesListImagesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesListImagesResponseResult struct {
	Images []AccountImageV1CloudflareImagesListImagesResponseResultImage `json:"images"`
	JSON   accountImageV1CloudflareImagesListImagesResponseResultJSON    `json:"-"`
}

// accountImageV1CloudflareImagesListImagesResponseResultJSON contains the JSON
// metadata for the struct [AccountImageV1CloudflareImagesListImagesResponseResult]
type accountImageV1CloudflareImagesListImagesResponseResultJSON struct {
	Images      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesListImagesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesListImagesResponseResultImage struct {
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
	Variants []AccountImageV1CloudflareImagesListImagesResponseResultImagesVariant `json:"variants" format:"uri"`
	JSON     accountImageV1CloudflareImagesListImagesResponseResultImageJSON       `json:"-"`
}

// accountImageV1CloudflareImagesListImagesResponseResultImageJSON contains the
// JSON metadata for the struct
// [AccountImageV1CloudflareImagesListImagesResponseResultImage]
type accountImageV1CloudflareImagesListImagesResponseResultImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesListImagesResponseResultImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type AccountImageV1CloudflareImagesListImagesResponseResultImagesVariant interface {
	ImplementsAccountImageV1CloudflareImagesListImagesResponseResultImagesVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountImageV1CloudflareImagesListImagesResponseResultImagesVariant)(nil)).Elem(),
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

// Whether the API call was successful
type AccountImageV1CloudflareImagesListImagesResponseSuccess bool

const (
	AccountImageV1CloudflareImagesListImagesResponseSuccessTrue AccountImageV1CloudflareImagesListImagesResponseSuccess = true
)

type AccountImageV1CloudflareImagesUploadAnImageViaURLResponse struct {
	Errors   []AccountImageV1CloudflareImagesUploadAnImageViaURLResponseError   `json:"errors"`
	Messages []AccountImageV1CloudflareImagesUploadAnImageViaURLResponseMessage `json:"messages"`
	Result   AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1CloudflareImagesUploadAnImageViaURLResponseSuccess `json:"success"`
	JSON    accountImageV1CloudflareImagesUploadAnImageViaURLResponseJSON    `json:"-"`
}

// accountImageV1CloudflareImagesUploadAnImageViaURLResponseJSON contains the JSON
// metadata for the struct
// [AccountImageV1CloudflareImagesUploadAnImageViaURLResponse]
type accountImageV1CloudflareImagesUploadAnImageViaURLResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesUploadAnImageViaURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesUploadAnImageViaURLResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountImageV1CloudflareImagesUploadAnImageViaURLResponseErrorJSON `json:"-"`
}

// accountImageV1CloudflareImagesUploadAnImageViaURLResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountImageV1CloudflareImagesUploadAnImageViaURLResponseError]
type accountImageV1CloudflareImagesUploadAnImageViaURLResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesUploadAnImageViaURLResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesUploadAnImageViaURLResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountImageV1CloudflareImagesUploadAnImageViaURLResponseMessageJSON `json:"-"`
}

// accountImageV1CloudflareImagesUploadAnImageViaURLResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountImageV1CloudflareImagesUploadAnImageViaURLResponseMessage]
type accountImageV1CloudflareImagesUploadAnImageViaURLResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesUploadAnImageViaURLResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResult struct {
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
	Variants []AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResultVariant `json:"variants" format:"uri"`
	JSON     accountImageV1CloudflareImagesUploadAnImageViaURLResponseResultJSON      `json:"-"`
}

// accountImageV1CloudflareImagesUploadAnImageViaURLResponseResultJSON contains the
// JSON metadata for the struct
// [AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResult]
type accountImageV1CloudflareImagesUploadAnImageViaURLResponseResultJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResultVariant interface {
	ImplementsAccountImageV1CloudflareImagesUploadAnImageViaURLResponseResultVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountImageV1CloudflareImagesUploadAnImageViaURLResponseResultVariant)(nil)).Elem(),
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

// Whether the API call was successful
type AccountImageV1CloudflareImagesUploadAnImageViaURLResponseSuccess bool

const (
	AccountImageV1CloudflareImagesUploadAnImageViaURLResponseSuccessTrue AccountImageV1CloudflareImagesUploadAnImageViaURLResponseSuccess = true
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

// This interface is a union satisfied by one of the following:
// [AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaFile],
// [AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaURL].
type AccountImageV1CloudflareImagesUploadAnImageViaURLParams interface {
	ImplementsAccountImageV1CloudflareImagesUploadAnImageViaURLParams()
}

type AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaFile struct {
	// An image binary data.
	File param.Field[interface{}] `json:"file,required"`
}

func (r AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaFile) ImplementsAccountImageV1CloudflareImagesUploadAnImageViaURLParams() {

}

type AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaURL struct {
	// A URL to fetch an image from origin.
	URL param.Field[string] `json:"url,required"`
}

func (r AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountImageV1CloudflareImagesUploadAnImageViaURLParamsImagesImageUploadViaURL) ImplementsAccountImageV1CloudflareImagesUploadAnImageViaURLParams() {

}
