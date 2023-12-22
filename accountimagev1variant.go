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
func (r *AccountImageV1VariantService) Get(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *VariantSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *AccountImageV1VariantService) Update(ctx context.Context, accountIdentifier string, identifier interface{}, body AccountImageV1VariantUpdateParams, opts ...option.RequestOption) (res *VariantSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *AccountImageV1VariantService) Delete(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *DeletedResponseFkozR92k, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *AccountImageV1VariantService) CloudflareImagesVariantsNewAVariant(ctx context.Context, accountIdentifier string, body AccountImageV1VariantCloudflareImagesVariantsNewAVariantParams, opts ...option.RequestOption) (res *VariantSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing variants.
func (r *AccountImageV1VariantService) CloudflareImagesVariantsListVariants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *VariantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VariantListResponse struct {
	Errors     []VariantListResponseError    `json:"errors"`
	Messages   []VariantListResponseMessage  `json:"messages"`
	Result     VariantListResponseResult     `json:"result"`
	ResultInfo VariantListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success VariantListResponseSuccess `json:"success"`
	JSON    variantListResponseJSON    `json:"-"`
}

// variantListResponseJSON contains the JSON metadata for the struct
// [VariantListResponse]
type variantListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    variantListResponseErrorJSON `json:"-"`
}

// variantListResponseErrorJSON contains the JSON metadata for the struct
// [VariantListResponseError]
type variantListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    variantListResponseMessageJSON `json:"-"`
}

// variantListResponseMessageJSON contains the JSON metadata for the struct
// [VariantListResponseMessage]
type variantListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantListResponseResult struct {
	Variants VariantListResponseResultVariants `json:"variants"`
	JSON     variantListResponseResultJSON     `json:"-"`
}

// variantListResponseResultJSON contains the JSON metadata for the struct
// [VariantListResponseResult]
type variantListResponseResultJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantListResponseResultVariants struct {
	Hero interface{}                           `json:"hero"`
	JSON variantListResponseResultVariantsJSON `json:"-"`
}

// variantListResponseResultVariantsJSON contains the JSON metadata for the struct
// [VariantListResponseResultVariants]
type variantListResponseResultVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantListResponseResultVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                           `json:"total_count"`
	JSON       variantListResponseResultInfoJSON `json:"-"`
}

// variantListResponseResultInfoJSON contains the JSON metadata for the struct
// [VariantListResponseResultInfo]
type variantListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VariantListResponseSuccess bool

const (
	VariantListResponseSuccessTrue VariantListResponseSuccess = true
)

type VariantSimpleResponse struct {
	Errors   []VariantSimpleResponseError   `json:"errors"`
	Messages []VariantSimpleResponseMessage `json:"messages"`
	Result   VariantSimpleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success VariantSimpleResponseSuccess `json:"success"`
	JSON    variantSimpleResponseJSON    `json:"-"`
}

// variantSimpleResponseJSON contains the JSON metadata for the struct
// [VariantSimpleResponse]
type variantSimpleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantSimpleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantSimpleResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    variantSimpleResponseErrorJSON `json:"-"`
}

// variantSimpleResponseErrorJSON contains the JSON metadata for the struct
// [VariantSimpleResponseError]
type variantSimpleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantSimpleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantSimpleResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    variantSimpleResponseMessageJSON `json:"-"`
}

// variantSimpleResponseMessageJSON contains the JSON metadata for the struct
// [VariantSimpleResponseMessage]
type variantSimpleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantSimpleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VariantSimpleResponseResult struct {
	Variant interface{}                     `json:"variant"`
	JSON    variantSimpleResponseResultJSON `json:"-"`
}

// variantSimpleResponseResultJSON contains the JSON metadata for the struct
// [VariantSimpleResponseResult]
type variantSimpleResponseResultJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantSimpleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VariantSimpleResponseSuccess bool

const (
	VariantSimpleResponseSuccessTrue VariantSimpleResponseSuccess = true
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
