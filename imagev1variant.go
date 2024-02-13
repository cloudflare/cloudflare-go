// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ImageV1VariantService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewImageV1VariantService] method
// instead.
type ImageV1VariantService struct {
	Options []option.RequestOption
}

// NewImageV1VariantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageV1VariantService(opts ...option.RequestOption) (r *ImageV1VariantService) {
	r = &ImageV1VariantService{}
	r.Options = opts
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *ImageV1VariantService) Update(ctx context.Context, accountID string, variantID interface{}, body ImageV1VariantUpdateParams, opts ...option.RequestOption) (res *ImageV1VariantUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *ImageV1VariantService) Delete(ctx context.Context, accountID string, variantID interface{}, opts ...option.RequestOption) (res *ImageV1VariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *ImageV1VariantService) CloudflareImagesVariantsNewAVariant(ctx context.Context, accountID string, body ImageV1VariantCloudflareImagesVariantsNewAVariantParams, opts ...option.RequestOption) (res *ImageV1VariantCloudflareImagesVariantsNewAVariantResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing variants.
func (r *ImageV1VariantService) CloudflareImagesVariantsListVariants(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ImageV1VariantCloudflareImagesVariantsListVariantsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single variant.
func (r *ImageV1VariantService) Get(ctx context.Context, accountID string, variantID interface{}, opts ...option.RequestOption) (res *ImageV1VariantGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", accountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV1VariantUpdateResponse struct {
	Variant ImageV1VariantUpdateResponseVariant `json:"variant"`
	JSON    imageV1VariantUpdateResponseJSON    `json:"-"`
}

// imageV1VariantUpdateResponseJSON contains the JSON metadata for the struct
// [ImageV1VariantUpdateResponse]
type imageV1VariantUpdateResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantUpdateResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantUpdateResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                    `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantUpdateResponseVariantJSON `json:"-"`
}

// imageV1VariantUpdateResponseVariantJSON contains the JSON metadata for the
// struct [ImageV1VariantUpdateResponseVariant]
type imageV1VariantUpdateResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantUpdateResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantUpdateResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantUpdateResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantUpdateResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                        `json:"width,required"`
	JSON  imageV1VariantUpdateResponseVariantOptionsJSON `json:"-"`
}

// imageV1VariantUpdateResponseVariantOptionsJSON contains the JSON metadata for
// the struct [ImageV1VariantUpdateResponseVariantOptions]
type imageV1VariantUpdateResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantUpdateResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantUpdateResponseVariantOptionsFit string

const (
	ImageV1VariantUpdateResponseVariantOptionsFitScaleDown ImageV1VariantUpdateResponseVariantOptionsFit = "scale-down"
	ImageV1VariantUpdateResponseVariantOptionsFitContain   ImageV1VariantUpdateResponseVariantOptionsFit = "contain"
	ImageV1VariantUpdateResponseVariantOptionsFitCover     ImageV1VariantUpdateResponseVariantOptionsFit = "cover"
	ImageV1VariantUpdateResponseVariantOptionsFitCrop      ImageV1VariantUpdateResponseVariantOptionsFit = "crop"
	ImageV1VariantUpdateResponseVariantOptionsFitPad       ImageV1VariantUpdateResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantUpdateResponseVariantOptionsMetadata string

const (
	ImageV1VariantUpdateResponseVariantOptionsMetadataKeep      ImageV1VariantUpdateResponseVariantOptionsMetadata = "keep"
	ImageV1VariantUpdateResponseVariantOptionsMetadataCopyright ImageV1VariantUpdateResponseVariantOptionsMetadata = "copyright"
	ImageV1VariantUpdateResponseVariantOptionsMetadataNone      ImageV1VariantUpdateResponseVariantOptionsMetadata = "none"
)

// Union satisfied by [ImageV1VariantDeleteResponseUnknown] or
// [shared.UnionString].
type ImageV1VariantDeleteResponse interface {
	ImplementsImageV1VariantDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV1VariantDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ImageV1VariantCloudflareImagesVariantsNewAVariantResponse struct {
	Variant ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariant `json:"variant"`
	JSON    imageV1VariantCloudflareImagesVariantsNewAVariantResponseJSON    `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsNewAVariantResponseJSON contains the JSON
// metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsNewAVariantResponse]
type imageV1VariantCloudflareImagesVariantsNewAVariantResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsNewAVariantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                                                 `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantJSON contains
// the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariant]
type imageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                                                     `json:"width,required"`
	JSON  imageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptions]
type imageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit string

const (
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFitScaleDown ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit = "scale-down"
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFitContain   ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit = "contain"
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFitCover     ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit = "cover"
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFitCrop      ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit = "crop"
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFitPad       ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadata string

const (
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadataKeep      ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadata = "keep"
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadataCopyright ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadata = "copyright"
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadataNone      ImageV1VariantCloudflareImagesVariantsNewAVariantResponseVariantOptionsMetadata = "none"
)

type ImageV1VariantCloudflareImagesVariantsListVariantsResponse struct {
	Variants ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariants `json:"variants"`
	JSON     imageV1VariantCloudflareImagesVariantsListVariantsResponseJSON     `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseJSON contains the JSON
// metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponse]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariants struct {
	Hero ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHero `json:"hero"`
	JSON imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsJSON contains
// the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariants]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHero struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                                                       `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHero]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                                                           `json:"width,required"`
	JSON  imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptions]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit string

const (
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFitScaleDown ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit = "scale-down"
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFitContain   ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit = "contain"
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFitCover     ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit = "cover"
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFitCrop      ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit = "crop"
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFitPad       ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadata string

const (
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadataKeep      ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadata = "keep"
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadataCopyright ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadata = "copyright"
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadataNone      ImageV1VariantCloudflareImagesVariantsListVariantsResponseVariantsHeroOptionsMetadata = "none"
)

type ImageV1VariantGetResponse struct {
	Variant ImageV1VariantGetResponseVariant `json:"variant"`
	JSON    imageV1VariantGetResponseJSON    `json:"-"`
}

// imageV1VariantGetResponseJSON contains the JSON metadata for the struct
// [ImageV1VariantGetResponse]
type imageV1VariantGetResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantGetResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantGetResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                 `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantGetResponseVariantJSON `json:"-"`
}

// imageV1VariantGetResponseVariantJSON contains the JSON metadata for the struct
// [ImageV1VariantGetResponseVariant]
type imageV1VariantGetResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantGetResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantGetResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantGetResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantGetResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                     `json:"width,required"`
	JSON  imageV1VariantGetResponseVariantOptionsJSON `json:"-"`
}

// imageV1VariantGetResponseVariantOptionsJSON contains the JSON metadata for the
// struct [ImageV1VariantGetResponseVariantOptions]
type imageV1VariantGetResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantGetResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantGetResponseVariantOptionsFit string

const (
	ImageV1VariantGetResponseVariantOptionsFitScaleDown ImageV1VariantGetResponseVariantOptionsFit = "scale-down"
	ImageV1VariantGetResponseVariantOptionsFitContain   ImageV1VariantGetResponseVariantOptionsFit = "contain"
	ImageV1VariantGetResponseVariantOptionsFitCover     ImageV1VariantGetResponseVariantOptionsFit = "cover"
	ImageV1VariantGetResponseVariantOptionsFitCrop      ImageV1VariantGetResponseVariantOptionsFit = "crop"
	ImageV1VariantGetResponseVariantOptionsFitPad       ImageV1VariantGetResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantGetResponseVariantOptionsMetadata string

const (
	ImageV1VariantGetResponseVariantOptionsMetadataKeep      ImageV1VariantGetResponseVariantOptionsMetadata = "keep"
	ImageV1VariantGetResponseVariantOptionsMetadataCopyright ImageV1VariantGetResponseVariantOptionsMetadata = "copyright"
	ImageV1VariantGetResponseVariantOptionsMetadataNone      ImageV1VariantGetResponseVariantOptionsMetadata = "none"
)

type ImageV1VariantUpdateParams struct {
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[ImageV1VariantUpdateParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r ImageV1VariantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantUpdateParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[ImageV1VariantUpdateParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[ImageV1VariantUpdateParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r ImageV1VariantUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantUpdateParamsOptionsFit string

const (
	ImageV1VariantUpdateParamsOptionsFitScaleDown ImageV1VariantUpdateParamsOptionsFit = "scale-down"
	ImageV1VariantUpdateParamsOptionsFitContain   ImageV1VariantUpdateParamsOptionsFit = "contain"
	ImageV1VariantUpdateParamsOptionsFitCover     ImageV1VariantUpdateParamsOptionsFit = "cover"
	ImageV1VariantUpdateParamsOptionsFitCrop      ImageV1VariantUpdateParamsOptionsFit = "crop"
	ImageV1VariantUpdateParamsOptionsFitPad       ImageV1VariantUpdateParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantUpdateParamsOptionsMetadata string

const (
	ImageV1VariantUpdateParamsOptionsMetadataKeep      ImageV1VariantUpdateParamsOptionsMetadata = "keep"
	ImageV1VariantUpdateParamsOptionsMetadataCopyright ImageV1VariantUpdateParamsOptionsMetadata = "copyright"
	ImageV1VariantUpdateParamsOptionsMetadataNone      ImageV1VariantUpdateParamsOptionsMetadata = "none"
)

type ImageV1VariantUpdateResponseEnvelope struct {
	Errors   []ImageV1VariantUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1VariantUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantUpdateResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ImageV1VariantUpdateResponseEnvelope]
type imageV1VariantUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    imageV1VariantUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ImageV1VariantUpdateResponseEnvelopeErrors]
type imageV1VariantUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    imageV1VariantUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ImageV1VariantUpdateResponseEnvelopeMessages]
type imageV1VariantUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantUpdateResponseEnvelopeSuccess bool

const (
	ImageV1VariantUpdateResponseEnvelopeSuccessTrue ImageV1VariantUpdateResponseEnvelopeSuccess = true
)

type ImageV1VariantDeleteResponseEnvelope struct {
	Errors   []ImageV1VariantDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1VariantDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantDeleteResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ImageV1VariantDeleteResponseEnvelope]
type imageV1VariantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    imageV1VariantDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ImageV1VariantDeleteResponseEnvelopeErrors]
type imageV1VariantDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    imageV1VariantDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ImageV1VariantDeleteResponseEnvelopeMessages]
type imageV1VariantDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantDeleteResponseEnvelopeSuccess bool

const (
	ImageV1VariantDeleteResponseEnvelopeSuccessTrue ImageV1VariantDeleteResponseEnvelopeSuccess = true
)

type ImageV1VariantCloudflareImagesVariantsNewAVariantParams struct {
	ID param.Field[interface{}] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r ImageV1VariantCloudflareImagesVariantsNewAVariantParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit string

const (
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitScaleDown ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "scale-down"
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitContain   ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "contain"
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitCover     ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "cover"
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitCrop      ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "crop"
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFitPad       ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata string

const (
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadataKeep      ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata = "keep"
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadataCopyright ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata = "copyright"
	ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadataNone      ImageV1VariantCloudflareImagesVariantsNewAVariantParamsOptionsMetadata = "none"
)

type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelope struct {
	Errors   []ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1VariantCloudflareImagesVariantsNewAVariantResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelope]
type imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrors]
type imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessages]
type imageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeSuccess bool

const (
	ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeSuccessTrue ImageV1VariantCloudflareImagesVariantsNewAVariantResponseEnvelopeSuccess = true
)

type ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelope struct {
	Errors   []ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1VariantCloudflareImagesVariantsListVariantsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelope]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrors struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrors]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessages struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessages]
type imageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeSuccess bool

const (
	ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeSuccessTrue ImageV1VariantCloudflareImagesVariantsListVariantsResponseEnvelopeSuccess = true
)

type ImageV1VariantGetResponseEnvelope struct {
	Errors   []ImageV1VariantGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1VariantGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantGetResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1VariantGetResponseEnvelope]
type imageV1VariantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    imageV1VariantGetResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1VariantGetResponseEnvelopeErrors]
type imageV1VariantGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    imageV1VariantGetResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImageV1VariantGetResponseEnvelopeMessages]
type imageV1VariantGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantGetResponseEnvelopeSuccess bool

const (
	ImageV1VariantGetResponseEnvelopeSuccessTrue ImageV1VariantGetResponseEnvelopeSuccess = true
)
