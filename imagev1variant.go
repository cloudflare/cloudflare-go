// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ImageV1VariantService) New(ctx context.Context, params ImageV1VariantNewParams, opts ...option.RequestOption) (res *ImageV1VariantNewResponse, err error) {
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
func (r *ImageV1VariantService) List(ctx context.Context, query ImageV1VariantListParams, opts ...option.RequestOption) (res *ImageV1VariantListResponse, err error) {
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
func (r *ImageV1VariantService) Edit(ctx context.Context, variantID interface{}, params ImageV1VariantEditParams, opts ...option.RequestOption) (res *ImageV1VariantEditResponse, err error) {
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
func (r *ImageV1VariantService) Get(ctx context.Context, variantID interface{}, query ImageV1VariantGetParams, opts ...option.RequestOption) (res *ImageV1VariantGetResponse, err error) {
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

type ImageV1VariantNewResponse struct {
	Variant ImageV1VariantNewResponseVariant `json:"variant"`
	JSON    imageV1VariantNewResponseJSON    `json:"-"`
}

// imageV1VariantNewResponseJSON contains the JSON metadata for the struct
// [ImageV1VariantNewResponse]
type imageV1VariantNewResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantNewResponseJSON) RawJSON() string {
	return r.raw
}

type ImageV1VariantNewResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantNewResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                 `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantNewResponseVariantJSON `json:"-"`
}

// imageV1VariantNewResponseVariantJSON contains the JSON metadata for the struct
// [ImageV1VariantNewResponseVariant]
type imageV1VariantNewResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantNewResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantNewResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantNewResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantNewResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantNewResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                     `json:"width,required"`
	JSON  imageV1VariantNewResponseVariantOptionsJSON `json:"-"`
}

// imageV1VariantNewResponseVariantOptionsJSON contains the JSON metadata for the
// struct [ImageV1VariantNewResponseVariantOptions]
type imageV1VariantNewResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantNewResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantNewResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantNewResponseVariantOptionsFit string

const (
	ImageV1VariantNewResponseVariantOptionsFitScaleDown ImageV1VariantNewResponseVariantOptionsFit = "scale-down"
	ImageV1VariantNewResponseVariantOptionsFitContain   ImageV1VariantNewResponseVariantOptionsFit = "contain"
	ImageV1VariantNewResponseVariantOptionsFitCover     ImageV1VariantNewResponseVariantOptionsFit = "cover"
	ImageV1VariantNewResponseVariantOptionsFitCrop      ImageV1VariantNewResponseVariantOptionsFit = "crop"
	ImageV1VariantNewResponseVariantOptionsFitPad       ImageV1VariantNewResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantNewResponseVariantOptionsMetadata string

const (
	ImageV1VariantNewResponseVariantOptionsMetadataKeep      ImageV1VariantNewResponseVariantOptionsMetadata = "keep"
	ImageV1VariantNewResponseVariantOptionsMetadataCopyright ImageV1VariantNewResponseVariantOptionsMetadata = "copyright"
	ImageV1VariantNewResponseVariantOptionsMetadataNone      ImageV1VariantNewResponseVariantOptionsMetadata = "none"
)

type ImageV1VariantListResponse struct {
	Variants ImageV1VariantListResponseVariants `json:"variants"`
	JSON     imageV1VariantListResponseJSON     `json:"-"`
}

// imageV1VariantListResponseJSON contains the JSON metadata for the struct
// [ImageV1VariantListResponse]
type imageV1VariantListResponseJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantListResponseJSON) RawJSON() string {
	return r.raw
}

type ImageV1VariantListResponseVariants struct {
	Hero ImageV1VariantListResponseVariantsHero `json:"hero"`
	JSON imageV1VariantListResponseVariantsJSON `json:"-"`
}

// imageV1VariantListResponseVariantsJSON contains the JSON metadata for the struct
// [ImageV1VariantListResponseVariants]
type imageV1VariantListResponseVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantListResponseVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantListResponseVariantsJSON) RawJSON() string {
	return r.raw
}

type ImageV1VariantListResponseVariantsHero struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantListResponseVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                       `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantListResponseVariantsHeroJSON `json:"-"`
}

// imageV1VariantListResponseVariantsHeroJSON contains the JSON metadata for the
// struct [ImageV1VariantListResponseVariantsHero]
type imageV1VariantListResponseVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantListResponseVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantListResponseVariantsHeroJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantListResponseVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantListResponseVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantListResponseVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                           `json:"width,required"`
	JSON  imageV1VariantListResponseVariantsHeroOptionsJSON `json:"-"`
}

// imageV1VariantListResponseVariantsHeroOptionsJSON contains the JSON metadata for
// the struct [ImageV1VariantListResponseVariantsHeroOptions]
type imageV1VariantListResponseVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantListResponseVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantListResponseVariantsHeroOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantListResponseVariantsHeroOptionsFit string

const (
	ImageV1VariantListResponseVariantsHeroOptionsFitScaleDown ImageV1VariantListResponseVariantsHeroOptionsFit = "scale-down"
	ImageV1VariantListResponseVariantsHeroOptionsFitContain   ImageV1VariantListResponseVariantsHeroOptionsFit = "contain"
	ImageV1VariantListResponseVariantsHeroOptionsFitCover     ImageV1VariantListResponseVariantsHeroOptionsFit = "cover"
	ImageV1VariantListResponseVariantsHeroOptionsFitCrop      ImageV1VariantListResponseVariantsHeroOptionsFit = "crop"
	ImageV1VariantListResponseVariantsHeroOptionsFitPad       ImageV1VariantListResponseVariantsHeroOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantListResponseVariantsHeroOptionsMetadata string

const (
	ImageV1VariantListResponseVariantsHeroOptionsMetadataKeep      ImageV1VariantListResponseVariantsHeroOptionsMetadata = "keep"
	ImageV1VariantListResponseVariantsHeroOptionsMetadataCopyright ImageV1VariantListResponseVariantsHeroOptionsMetadata = "copyright"
	ImageV1VariantListResponseVariantsHeroOptionsMetadataNone      ImageV1VariantListResponseVariantsHeroOptionsMetadata = "none"
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

type ImageV1VariantEditResponse struct {
	Variant ImageV1VariantEditResponseVariant `json:"variant"`
	JSON    imageV1VariantEditResponseJSON    `json:"-"`
}

// imageV1VariantEditResponseJSON contains the JSON metadata for the struct
// [ImageV1VariantEditResponse]
type imageV1VariantEditResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantEditResponseJSON) RawJSON() string {
	return r.raw
}

type ImageV1VariantEditResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageV1VariantEditResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                  `json:"neverRequireSignedURLs"`
	JSON                   imageV1VariantEditResponseVariantJSON `json:"-"`
}

// imageV1VariantEditResponseVariantJSON contains the JSON metadata for the struct
// [ImageV1VariantEditResponseVariant]
type imageV1VariantEditResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageV1VariantEditResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantEditResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type ImageV1VariantEditResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageV1VariantEditResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageV1VariantEditResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                      `json:"width,required"`
	JSON  imageV1VariantEditResponseVariantOptionsJSON `json:"-"`
}

// imageV1VariantEditResponseVariantOptionsJSON contains the JSON metadata for the
// struct [ImageV1VariantEditResponseVariantOptions]
type imageV1VariantEditResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1VariantEditResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1VariantEditResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageV1VariantEditResponseVariantOptionsFit string

const (
	ImageV1VariantEditResponseVariantOptionsFitScaleDown ImageV1VariantEditResponseVariantOptionsFit = "scale-down"
	ImageV1VariantEditResponseVariantOptionsFitContain   ImageV1VariantEditResponseVariantOptionsFit = "contain"
	ImageV1VariantEditResponseVariantOptionsFitCover     ImageV1VariantEditResponseVariantOptionsFit = "cover"
	ImageV1VariantEditResponseVariantOptionsFitCrop      ImageV1VariantEditResponseVariantOptionsFit = "crop"
	ImageV1VariantEditResponseVariantOptionsFitPad       ImageV1VariantEditResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type ImageV1VariantEditResponseVariantOptionsMetadata string

const (
	ImageV1VariantEditResponseVariantOptionsMetadataKeep      ImageV1VariantEditResponseVariantOptionsMetadata = "keep"
	ImageV1VariantEditResponseVariantOptionsMetadataCopyright ImageV1VariantEditResponseVariantOptionsMetadata = "copyright"
	ImageV1VariantEditResponseVariantOptionsMetadataNone      ImageV1VariantEditResponseVariantOptionsMetadata = "none"
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

func (r imageV1VariantGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantGetResponseVariantJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantGetResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
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
	Result   ImageV1VariantNewResponse                   `json:"result,required"`
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

func (r imageV1VariantNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ImageV1VariantListResponse                   `json:"result,required"`
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

func (r imageV1VariantListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ImageV1VariantEditResponse                   `json:"result,required"`
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

func (r imageV1VariantEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r imageV1VariantGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ImageV1VariantGetResponseEnvelopeSuccess bool

const (
	ImageV1VariantGetResponseEnvelopeSuccessTrue ImageV1VariantGetResponseEnvelopeSuccess = true
)
