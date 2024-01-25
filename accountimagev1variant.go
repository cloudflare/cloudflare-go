// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountImageV1VariantService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1VariantService]
// method instead.
type AccountImageV1VariantService struct {
	Options []option.RequestOption
}

// NewAccountImageV1VariantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1VariantService(opts ...option.RequestOption) (r *AccountImageV1VariantService) {
	r = &AccountImageV1VariantService{}
	r.Options = opts
	return
}

// Fetch details for a single variant.
func (r *AccountImageV1VariantService) Get(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *AccountImageV1VariantGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *AccountImageV1VariantService) Update(ctx context.Context, accountIdentifier string, identifier interface{}, body AccountImageV1VariantUpdateParams, opts ...option.RequestOption) (res *AccountImageV1VariantUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *AccountImageV1VariantService) Delete(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *AccountImageV1VariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *AccountImageV1VariantService) CloudflareImagesVariantsNewAVariant(ctx context.Context, accountIdentifier string, body AccountImageV1VariantCloudflareImagesVariantsNewAVariantParams, opts ...option.RequestOption) (res *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing variants.
func (r *AccountImageV1VariantService) CloudflareImagesVariantsListVariants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountImageV1VariantGetResponse struct {
	Errors   []AccountImageV1VariantGetResponseError   `json:"errors"`
	Messages []AccountImageV1VariantGetResponseMessage `json:"messages"`
	Result   AccountImageV1VariantGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1VariantGetResponseSuccess `json:"success"`
	JSON    accountImageV1VariantGetResponseJSON    `json:"-"`
}

// accountImageV1VariantGetResponseJSON contains the JSON metadata for the struct
// [AccountImageV1VariantGetResponse]
type accountImageV1VariantGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountImageV1VariantGetResponseErrorJSON `json:"-"`
}

// accountImageV1VariantGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountImageV1VariantGetResponseError]
type accountImageV1VariantGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountImageV1VariantGetResponseMessageJSON `json:"-"`
}

// accountImageV1VariantGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountImageV1VariantGetResponseMessage]
type accountImageV1VariantGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantGetResponseResult struct {
	Variant AccountImageV1VariantGetResponseResultVariant `json:"variant"`
	JSON    accountImageV1VariantGetResponseResultJSON    `json:"-"`
}

// accountImageV1VariantGetResponseResultJSON contains the JSON metadata for the
// struct [AccountImageV1VariantGetResponseResult]
type accountImageV1VariantGetResponseResultJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantGetResponseResultVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options AccountImageV1VariantGetResponseResultVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                              `json:"neverRequireSignedURLs"`
	JSON                   accountImageV1VariantGetResponseResultVariantJSON `json:"-"`
}

// accountImageV1VariantGetResponseResultVariantJSON contains the JSON metadata for
// the struct [AccountImageV1VariantGetResponseResultVariant]
type accountImageV1VariantGetResponseResultVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountImageV1VariantGetResponseResultVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type AccountImageV1VariantGetResponseResultVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit AccountImageV1VariantGetResponseResultVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata AccountImageV1VariantGetResponseResultVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                                  `json:"width,required"`
	JSON  accountImageV1VariantGetResponseResultVariantOptionsJSON `json:"-"`
}

// accountImageV1VariantGetResponseResultVariantOptionsJSON contains the JSON
// metadata for the struct [AccountImageV1VariantGetResponseResultVariantOptions]
type accountImageV1VariantGetResponseResultVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantGetResponseResultVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type AccountImageV1VariantGetResponseResultVariantOptionsFit string

const (
	AccountImageV1VariantGetResponseResultVariantOptionsFitScaleDown AccountImageV1VariantGetResponseResultVariantOptionsFit = "scale-down"
	AccountImageV1VariantGetResponseResultVariantOptionsFitContain   AccountImageV1VariantGetResponseResultVariantOptionsFit = "contain"
	AccountImageV1VariantGetResponseResultVariantOptionsFitCover     AccountImageV1VariantGetResponseResultVariantOptionsFit = "cover"
	AccountImageV1VariantGetResponseResultVariantOptionsFitCrop      AccountImageV1VariantGetResponseResultVariantOptionsFit = "crop"
	AccountImageV1VariantGetResponseResultVariantOptionsFitPad       AccountImageV1VariantGetResponseResultVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type AccountImageV1VariantGetResponseResultVariantOptionsMetadata string

const (
	AccountImageV1VariantGetResponseResultVariantOptionsMetadataKeep      AccountImageV1VariantGetResponseResultVariantOptionsMetadata = "keep"
	AccountImageV1VariantGetResponseResultVariantOptionsMetadataCopyright AccountImageV1VariantGetResponseResultVariantOptionsMetadata = "copyright"
	AccountImageV1VariantGetResponseResultVariantOptionsMetadataNone      AccountImageV1VariantGetResponseResultVariantOptionsMetadata = "none"
)

// Whether the API call was successful
type AccountImageV1VariantGetResponseSuccess bool

const (
	AccountImageV1VariantGetResponseSuccessTrue AccountImageV1VariantGetResponseSuccess = true
)

type AccountImageV1VariantUpdateResponse struct {
	Errors   []AccountImageV1VariantUpdateResponseError   `json:"errors"`
	Messages []AccountImageV1VariantUpdateResponseMessage `json:"messages"`
	Result   AccountImageV1VariantUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1VariantUpdateResponseSuccess `json:"success"`
	JSON    accountImageV1VariantUpdateResponseJSON    `json:"-"`
}

// accountImageV1VariantUpdateResponseJSON contains the JSON metadata for the
// struct [AccountImageV1VariantUpdateResponse]
type accountImageV1VariantUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountImageV1VariantUpdateResponseErrorJSON `json:"-"`
}

// accountImageV1VariantUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountImageV1VariantUpdateResponseError]
type accountImageV1VariantUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountImageV1VariantUpdateResponseMessageJSON `json:"-"`
}

// accountImageV1VariantUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountImageV1VariantUpdateResponseMessage]
type accountImageV1VariantUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantUpdateResponseResult struct {
	Variant AccountImageV1VariantUpdateResponseResultVariant `json:"variant"`
	JSON    accountImageV1VariantUpdateResponseResultJSON    `json:"-"`
}

// accountImageV1VariantUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountImageV1VariantUpdateResponseResult]
type accountImageV1VariantUpdateResponseResultJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantUpdateResponseResultVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options AccountImageV1VariantUpdateResponseResultVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                                 `json:"neverRequireSignedURLs"`
	JSON                   accountImageV1VariantUpdateResponseResultVariantJSON `json:"-"`
}

// accountImageV1VariantUpdateResponseResultVariantJSON contains the JSON metadata
// for the struct [AccountImageV1VariantUpdateResponseResultVariant]
type accountImageV1VariantUpdateResponseResultVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountImageV1VariantUpdateResponseResultVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type AccountImageV1VariantUpdateResponseResultVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit AccountImageV1VariantUpdateResponseResultVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata AccountImageV1VariantUpdateResponseResultVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                                     `json:"width,required"`
	JSON  accountImageV1VariantUpdateResponseResultVariantOptionsJSON `json:"-"`
}

// accountImageV1VariantUpdateResponseResultVariantOptionsJSON contains the JSON
// metadata for the struct
// [AccountImageV1VariantUpdateResponseResultVariantOptions]
type accountImageV1VariantUpdateResponseResultVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantUpdateResponseResultVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type AccountImageV1VariantUpdateResponseResultVariantOptionsFit string

const (
	AccountImageV1VariantUpdateResponseResultVariantOptionsFitScaleDown AccountImageV1VariantUpdateResponseResultVariantOptionsFit = "scale-down"
	AccountImageV1VariantUpdateResponseResultVariantOptionsFitContain   AccountImageV1VariantUpdateResponseResultVariantOptionsFit = "contain"
	AccountImageV1VariantUpdateResponseResultVariantOptionsFitCover     AccountImageV1VariantUpdateResponseResultVariantOptionsFit = "cover"
	AccountImageV1VariantUpdateResponseResultVariantOptionsFitCrop      AccountImageV1VariantUpdateResponseResultVariantOptionsFit = "crop"
	AccountImageV1VariantUpdateResponseResultVariantOptionsFitPad       AccountImageV1VariantUpdateResponseResultVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type AccountImageV1VariantUpdateResponseResultVariantOptionsMetadata string

const (
	AccountImageV1VariantUpdateResponseResultVariantOptionsMetadataKeep      AccountImageV1VariantUpdateResponseResultVariantOptionsMetadata = "keep"
	AccountImageV1VariantUpdateResponseResultVariantOptionsMetadataCopyright AccountImageV1VariantUpdateResponseResultVariantOptionsMetadata = "copyright"
	AccountImageV1VariantUpdateResponseResultVariantOptionsMetadataNone      AccountImageV1VariantUpdateResponseResultVariantOptionsMetadata = "none"
)

// Whether the API call was successful
type AccountImageV1VariantUpdateResponseSuccess bool

const (
	AccountImageV1VariantUpdateResponseSuccessTrue AccountImageV1VariantUpdateResponseSuccess = true
)

type AccountImageV1VariantDeleteResponse struct {
	Errors   []AccountImageV1VariantDeleteResponseError   `json:"errors"`
	Messages []AccountImageV1VariantDeleteResponseMessage `json:"messages"`
	Result   interface{}                                  `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1VariantDeleteResponseSuccess `json:"success"`
	JSON    accountImageV1VariantDeleteResponseJSON    `json:"-"`
}

// accountImageV1VariantDeleteResponseJSON contains the JSON metadata for the
// struct [AccountImageV1VariantDeleteResponse]
type accountImageV1VariantDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountImageV1VariantDeleteResponseErrorJSON `json:"-"`
}

// accountImageV1VariantDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountImageV1VariantDeleteResponseError]
type accountImageV1VariantDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountImageV1VariantDeleteResponseMessageJSON `json:"-"`
}

// accountImageV1VariantDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountImageV1VariantDeleteResponseMessage]
type accountImageV1VariantDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountImageV1VariantDeleteResponseSuccess bool

const (
	AccountImageV1VariantDeleteResponseSuccessTrue AccountImageV1VariantDeleteResponseSuccess = true
)

type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponse struct {
	Errors   []AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseError   `json:"errors"`
	Messages []AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessage `json:"messages"`
	Result   AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseSuccess `json:"success"`
	JSON    accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseJSON    `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseJSON contains
// the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponse]
type accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseErrorJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseError]
type accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessageJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessage]
type accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResult struct {
	Variant AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariant `json:"variant"`
	JSON    accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultJSON    `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResult]
type accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                                                              `json:"neverRequireSignedURLs"`
	JSON                   accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariant]
type accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                                                                  `json:"width,required"`
	JSON  accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptions]
type accountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit string

const (
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFitScaleDown AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit = "scale-down"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFitContain   AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit = "contain"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFitCover     AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit = "cover"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFitCrop      AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit = "crop"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFitPad       AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadata string

const (
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadataKeep      AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadata = "keep"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadataCopyright AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadata = "copyright"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadataNone      AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseResultVariantOptionsMetadata = "none"
)

// Whether the API call was successful
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseSuccess bool

const (
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseSuccessTrue AccountImageV1VariantCloudflareImagesVariantsNewAVariantResponseSuccess = true
)

type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponse struct {
	Errors   []AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseError   `json:"errors"`
	Messages []AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessage `json:"messages"`
	Result   AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseSuccess `json:"success"`
	JSON    accountImageV1VariantCloudflareImagesVariantsListVariantsResponseJSON    `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseJSON contains
// the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponse]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountImageV1VariantCloudflareImagesVariantsListVariantsResponseErrorJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseError]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessageJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessage]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResult struct {
	Variants AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariants `json:"variants"`
	JSON     accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultJSON     `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResult]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariants struct {
	Hero AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHero `json:"hero"`
	JSON accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariants]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHero struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                                                                    `json:"neverRequireSignedURLs"`
	JSON                   accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHero]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                                                                        `json:"width,required"`
	JSON  accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsJSON `json:"-"`
}

// accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsJSON
// contains the JSON metadata for the struct
// [AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptions]
type accountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit string

const (
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFitScaleDown AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit = "scale-down"
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFitContain   AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit = "contain"
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFitCover     AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit = "cover"
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFitCrop      AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit = "crop"
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFitPad       AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadata string

const (
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadataKeep      AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadata = "keep"
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadataCopyright AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadata = "copyright"
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadataNone      AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseResultVariantsHeroOptionsMetadata = "none"
)

// Whether the API call was successful
type AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseSuccess bool

const (
	AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseSuccessTrue AccountImageV1VariantCloudflareImagesVariantsListVariantsResponseSuccess = true
)

type AccountImageV1VariantUpdateParams struct {
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[AccountImageV1VariantUpdateParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r AccountImageV1VariantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type AccountImageV1VariantUpdateParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[AccountImageV1VariantUpdateParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[AccountImageV1VariantUpdateParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r AccountImageV1VariantUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type AccountImageV1VariantUpdateParamsOptionsFit string

const (
	AccountImageV1VariantUpdateParamsOptionsFitScaleDown AccountImageV1VariantUpdateParamsOptionsFit = "scale-down"
	AccountImageV1VariantUpdateParamsOptionsFitContain   AccountImageV1VariantUpdateParamsOptionsFit = "contain"
	AccountImageV1VariantUpdateParamsOptionsFitCover     AccountImageV1VariantUpdateParamsOptionsFit = "cover"
	AccountImageV1VariantUpdateParamsOptionsFitCrop      AccountImageV1VariantUpdateParamsOptionsFit = "crop"
	AccountImageV1VariantUpdateParamsOptionsFitPad       AccountImageV1VariantUpdateParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type AccountImageV1VariantUpdateParamsOptionsMetadata string

const (
	AccountImageV1VariantUpdateParamsOptionsMetadataKeep      AccountImageV1VariantUpdateParamsOptionsMetadata = "keep"
	AccountImageV1VariantUpdateParamsOptionsMetadataCopyright AccountImageV1VariantUpdateParamsOptionsMetadata = "copyright"
	AccountImageV1VariantUpdateParamsOptionsMetadataNone      AccountImageV1VariantUpdateParamsOptionsMetadata = "none"
)

type AccountImageV1VariantCloudflareImagesVariantsNewAVariantParams struct {
	ID param.Field[interface{}] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r AccountImageV1VariantCloudflareImagesVariantsNewAVariantParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit string

const (
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitScaleDown AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "scale-down"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitContain   AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "contain"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitCover     AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "cover"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitCrop      AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "crop"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitPad       AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata string

const (
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadataKeep      AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata = "keep"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadataCopyright AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata = "copyright"
	AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadataNone      AccountImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata = "none"
)
