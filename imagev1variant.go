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

// Specify variants that allow you to resize images for different use cases.
func (r *ImageV1VariantService) New(ctx context.Context, params ImageV1VariantNewParams, opts ...option.RequestOption) (res *ImageVariant, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing variants.
func (r *ImageV1VariantService) List(ctx context.Context, query ImageV1VariantListParams, opts ...option.RequestOption) (res *ImageVariants, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *ImageV1VariantService) Delete(ctx context.Context, variantID interface{}, body ImageV1VariantDeleteParams, opts ...option.RequestOption) (res *ImageV1VariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", body.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *ImageV1VariantService) Edit(ctx context.Context, variantID interface{}, params ImageV1VariantEditParams, opts ...option.RequestOption) (res *ImageVariant, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", params.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single variant.
func (r *ImageV1VariantService) Get(ctx context.Context, variantID interface{}, query ImageV1VariantGetParams, opts ...option.RequestOption) (res *ImageVariant, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1VariantGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", query.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageVariant struct {
	Variant ImageVariantVariant `json:"variant"`
	JSON    imageVariantJSON    `json:"-"`
}

// imageVariantJSON contains the JSON metadata for the struct [ImageVariant]
type imageVariantJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageVariantVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageVariantVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                    `json:"neverRequireSignedURLs"`
	JSON                   imageVariantVariantJSON `json:"-"`
}

// imageVariantVariantJSON contains the JSON metadata for the struct
// [ImageVariantVariant]
type imageVariantVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageVariantVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageVariantVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageVariantVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageVariantVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                        `json:"width,required"`
	JSON  imageVariantVariantOptionsJSON `json:"-"`
}

// imageVariantVariantOptionsJSON contains the JSON metadata for the struct
// [ImageVariantVariantOptions]
type imageVariantVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageVariantVariantOptionsFit string

const (
	ImageVariantVariantOptionsFitScaleDown ImageVariantVariantOptionsFit = "scale-down"
	ImageVariantVariantOptionsFitContain   ImageVariantVariantOptionsFit = "contain"
	ImageVariantVariantOptionsFitCover     ImageVariantVariantOptionsFit = "cover"
	ImageVariantVariantOptionsFitCrop      ImageVariantVariantOptionsFit = "crop"
	ImageVariantVariantOptionsFitPad       ImageVariantVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageVariantVariantOptionsMetadata string

const (
	ImageVariantVariantOptionsMetadataKeep      ImageVariantVariantOptionsMetadata = "keep"
	ImageVariantVariantOptionsMetadataCopyright ImageVariantVariantOptionsMetadata = "copyright"
	ImageVariantVariantOptionsMetadataNone      ImageVariantVariantOptionsMetadata = "none"
)

type ImageVariants struct {
	Variants ImageVariantsVariants `json:"variants"`
	JSON     imageVariantsJSON     `json:"-"`
}

// imageVariantsJSON contains the JSON metadata for the struct [ImageVariants]
type imageVariantsJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageVariantsVariants struct {
	Hero ImageVariantsVariantsHero `json:"hero"`
	JSON imageVariantsVariantsJSON `json:"-"`
}

// imageVariantsVariantsJSON contains the JSON metadata for the struct
// [ImageVariantsVariants]
type imageVariantsVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantsVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageVariantsVariantsHero struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageVariantsVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                          `json:"neverRequireSignedURLs"`
	JSON                   imageVariantsVariantsHeroJSON `json:"-"`
}

// imageVariantsVariantsHeroJSON contains the JSON metadata for the struct
// [ImageVariantsVariantsHero]
type imageVariantsVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageVariantsVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageVariantsVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageVariantsVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageVariantsVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                              `json:"width,required"`
	JSON  imageVariantsVariantsHeroOptionsJSON `json:"-"`
}

// imageVariantsVariantsHeroOptionsJSON contains the JSON metadata for the struct
// [ImageVariantsVariantsHeroOptions]
type imageVariantsVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantsVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageVariantsVariantsHeroOptionsFit string

const (
	ImageVariantsVariantsHeroOptionsFitScaleDown ImageVariantsVariantsHeroOptionsFit = "scale-down"
	ImageVariantsVariantsHeroOptionsFitContain   ImageVariantsVariantsHeroOptionsFit = "contain"
	ImageVariantsVariantsHeroOptionsFitCover     ImageVariantsVariantsHeroOptionsFit = "cover"
	ImageVariantsVariantsHeroOptionsFitCrop      ImageVariantsVariantsHeroOptionsFit = "crop"
	ImageVariantsVariantsHeroOptionsFitPad       ImageVariantsVariantsHeroOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageVariantsVariantsHeroOptionsMetadata string

const (
	ImageVariantsVariantsHeroOptionsMetadataKeep      ImageVariantsVariantsHeroOptionsMetadata = "keep"
	ImageVariantsVariantsHeroOptionsMetadataCopyright ImageVariantsVariantsHeroOptionsMetadata = "copyright"
	ImageVariantsVariantsHeroOptionsMetadataNone      ImageVariantsVariantsHeroOptionsMetadata = "none"
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

type ImageV1VariantNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string]      `path:"account_id,required"`
	ID        param.Field[interface{}] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[ImageV1VariantNewParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r ImageV1VariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantNewParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[ImageV1VariantNewParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[ImageV1VariantNewParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r ImageV1VariantNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantNewParamsOptionsFit string

const (
	ImageV1VariantNewParamsOptionsFitScaleDown ImageV1VariantNewParamsOptionsFit = "scale-down"
	ImageV1VariantNewParamsOptionsFitContain   ImageV1VariantNewParamsOptionsFit = "contain"
	ImageV1VariantNewParamsOptionsFitCover     ImageV1VariantNewParamsOptionsFit = "cover"
	ImageV1VariantNewParamsOptionsFitCrop      ImageV1VariantNewParamsOptionsFit = "crop"
	ImageV1VariantNewParamsOptionsFitPad       ImageV1VariantNewParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantNewParamsOptionsMetadata string

const (
	ImageV1VariantNewParamsOptionsMetadataKeep      ImageV1VariantNewParamsOptionsMetadata = "keep"
	ImageV1VariantNewParamsOptionsMetadataCopyright ImageV1VariantNewParamsOptionsMetadata = "copyright"
	ImageV1VariantNewParamsOptionsMetadataNone      ImageV1VariantNewParamsOptionsMetadata = "none"
)

type ImageV1VariantNewResponseEnvelope struct {
	Errors   []ImageV1VariantNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariant                                `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantNewResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1VariantNewResponseEnvelope]
type imageV1VariantNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    imageV1VariantNewResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1VariantNewResponseEnvelopeErrors]
type imageV1VariantNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    imageV1VariantNewResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImageV1VariantNewResponseEnvelopeMessages]
type imageV1VariantNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantNewResponseEnvelopeSuccess bool

const (
	ImageV1VariantNewResponseEnvelopeSuccessTrue ImageV1VariantNewResponseEnvelopeSuccess = true
)

type ImageV1VariantListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ImageV1VariantListResponseEnvelope struct {
	Errors   []ImageV1VariantListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantListResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariants                                `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantListResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantListResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1VariantListResponseEnvelope]
type imageV1VariantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    imageV1VariantListResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1VariantListResponseEnvelopeErrors]
type imageV1VariantListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    imageV1VariantListResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ImageV1VariantListResponseEnvelopeMessages]
type imageV1VariantListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantListResponseEnvelopeSuccess bool

const (
	ImageV1VariantListResponseEnvelopeSuccessTrue ImageV1VariantListResponseEnvelopeSuccess = true
)

type ImageV1VariantDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

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

type ImageV1VariantEditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[ImageV1VariantEditParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r ImageV1VariantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantEditParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[ImageV1VariantEditParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[ImageV1VariantEditParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r ImageV1VariantEditParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantEditParamsOptionsFit string

const (
	ImageV1VariantEditParamsOptionsFitScaleDown ImageV1VariantEditParamsOptionsFit = "scale-down"
	ImageV1VariantEditParamsOptionsFitContain   ImageV1VariantEditParamsOptionsFit = "contain"
	ImageV1VariantEditParamsOptionsFitCover     ImageV1VariantEditParamsOptionsFit = "cover"
	ImageV1VariantEditParamsOptionsFitCrop      ImageV1VariantEditParamsOptionsFit = "crop"
	ImageV1VariantEditParamsOptionsFitPad       ImageV1VariantEditParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantEditParamsOptionsMetadata string

const (
	ImageV1VariantEditParamsOptionsMetadataKeep      ImageV1VariantEditParamsOptionsMetadata = "keep"
	ImageV1VariantEditParamsOptionsMetadataCopyright ImageV1VariantEditParamsOptionsMetadata = "copyright"
	ImageV1VariantEditParamsOptionsMetadataNone      ImageV1VariantEditParamsOptionsMetadata = "none"
)

type ImageV1VariantEditResponseEnvelope struct {
	Errors   []ImageV1VariantEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariant                                 `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1VariantEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1VariantEditResponseEnvelopeJSON    `json:"-"`
}

// imageV1VariantEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1VariantEditResponseEnvelope]
type imageV1VariantEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    imageV1VariantEditResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1VariantEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1VariantEditResponseEnvelopeErrors]
type imageV1VariantEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1VariantEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    imageV1VariantEditResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1VariantEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ImageV1VariantEditResponseEnvelopeMessages]
type imageV1VariantEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1VariantEditResponseEnvelopeSuccess bool

const (
	ImageV1VariantEditResponseEnvelopeSuccessTrue ImageV1VariantEditResponseEnvelopeSuccess = true
)

type ImageV1VariantGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ImageV1VariantGetResponseEnvelope struct {
	Errors   []ImageV1VariantGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1VariantGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariant                                `json:"result,required"`
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
