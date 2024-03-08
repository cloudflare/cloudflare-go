// File generated from our OpenAPI spec by Stainless.

package images

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

// V1VariantService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewV1VariantService] method instead.
type V1VariantService struct {
	Options []option.RequestOption
}

// NewV1VariantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1VariantService(opts ...option.RequestOption) (r *V1VariantService) {
	r = &V1VariantService{}
	r.Options = opts
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *V1VariantService) New(ctx context.Context, params V1VariantNewParams, opts ...option.RequestOption) (res *V1VariantNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing variants.
func (r *V1VariantService) List(ctx context.Context, query V1VariantListParams, opts ...option.RequestOption) (res *V1VariantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *V1VariantService) Delete(ctx context.Context, variantID interface{}, body V1VariantDeleteParams, opts ...option.RequestOption) (res *V1VariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", body.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *V1VariantService) Edit(ctx context.Context, variantID interface{}, params V1VariantEditParams, opts ...option.RequestOption) (res *V1VariantEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", params.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single variant.
func (r *V1VariantService) Get(ctx context.Context, variantID interface{}, query V1VariantGetParams, opts ...option.RequestOption) (res *V1VariantGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%v", query.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V1VariantNewResponse struct {
	Variant V1VariantNewResponseVariant `json:"variant"`
	JSON    v1VariantNewResponseJSON    `json:"-"`
}

// v1VariantNewResponseJSON contains the JSON metadata for the struct
// [V1VariantNewResponse]
type v1VariantNewResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantNewResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantNewResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                            `json:"neverRequireSignedURLs"`
	JSON                   v1VariantNewResponseVariantJSON `json:"-"`
}

// v1VariantNewResponseVariantJSON contains the JSON metadata for the struct
// [V1VariantNewResponseVariant]
type v1VariantNewResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantNewResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantNewResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantNewResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantNewResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                `json:"width,required"`
	JSON  v1VariantNewResponseVariantOptionsJSON `json:"-"`
}

// v1VariantNewResponseVariantOptionsJSON contains the JSON metadata for the struct
// [V1VariantNewResponseVariantOptions]
type v1VariantNewResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantNewResponseVariantOptionsFit string

const (
	V1VariantNewResponseVariantOptionsFitScaleDown V1VariantNewResponseVariantOptionsFit = "scale-down"
	V1VariantNewResponseVariantOptionsFitContain   V1VariantNewResponseVariantOptionsFit = "contain"
	V1VariantNewResponseVariantOptionsFitCover     V1VariantNewResponseVariantOptionsFit = "cover"
	V1VariantNewResponseVariantOptionsFitCrop      V1VariantNewResponseVariantOptionsFit = "crop"
	V1VariantNewResponseVariantOptionsFitPad       V1VariantNewResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type V1VariantNewResponseVariantOptionsMetadata string

const (
	V1VariantNewResponseVariantOptionsMetadataKeep      V1VariantNewResponseVariantOptionsMetadata = "keep"
	V1VariantNewResponseVariantOptionsMetadataCopyright V1VariantNewResponseVariantOptionsMetadata = "copyright"
	V1VariantNewResponseVariantOptionsMetadataNone      V1VariantNewResponseVariantOptionsMetadata = "none"
)

type V1VariantListResponse struct {
	Variants V1VariantListResponseVariants `json:"variants"`
	JSON     v1VariantListResponseJSON     `json:"-"`
}

// v1VariantListResponseJSON contains the JSON metadata for the struct
// [V1VariantListResponse]
type v1VariantListResponseJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantListResponseVariants struct {
	Hero V1VariantListResponseVariantsHero `json:"hero"`
	JSON v1VariantListResponseVariantsJSON `json:"-"`
}

// v1VariantListResponseVariantsJSON contains the JSON metadata for the struct
// [V1VariantListResponseVariants]
type v1VariantListResponseVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseVariantsJSON) RawJSON() string {
	return r.raw
}

type V1VariantListResponseVariantsHero struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantListResponseVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                  `json:"neverRequireSignedURLs"`
	JSON                   v1VariantListResponseVariantsHeroJSON `json:"-"`
}

// v1VariantListResponseVariantsHeroJSON contains the JSON metadata for the struct
// [V1VariantListResponseVariantsHero]
type v1VariantListResponseVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantListResponseVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseVariantsHeroJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantListResponseVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantListResponseVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantListResponseVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                      `json:"width,required"`
	JSON  v1VariantListResponseVariantsHeroOptionsJSON `json:"-"`
}

// v1VariantListResponseVariantsHeroOptionsJSON contains the JSON metadata for the
// struct [V1VariantListResponseVariantsHeroOptions]
type v1VariantListResponseVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseVariantsHeroOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantListResponseVariantsHeroOptionsFit string

const (
	V1VariantListResponseVariantsHeroOptionsFitScaleDown V1VariantListResponseVariantsHeroOptionsFit = "scale-down"
	V1VariantListResponseVariantsHeroOptionsFitContain   V1VariantListResponseVariantsHeroOptionsFit = "contain"
	V1VariantListResponseVariantsHeroOptionsFitCover     V1VariantListResponseVariantsHeroOptionsFit = "cover"
	V1VariantListResponseVariantsHeroOptionsFitCrop      V1VariantListResponseVariantsHeroOptionsFit = "crop"
	V1VariantListResponseVariantsHeroOptionsFitPad       V1VariantListResponseVariantsHeroOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type V1VariantListResponseVariantsHeroOptionsMetadata string

const (
	V1VariantListResponseVariantsHeroOptionsMetadataKeep      V1VariantListResponseVariantsHeroOptionsMetadata = "keep"
	V1VariantListResponseVariantsHeroOptionsMetadataCopyright V1VariantListResponseVariantsHeroOptionsMetadata = "copyright"
	V1VariantListResponseVariantsHeroOptionsMetadataNone      V1VariantListResponseVariantsHeroOptionsMetadata = "none"
)

// Union satisfied by [images.V1VariantDeleteResponseUnknown] or
// [shared.UnionString].
type V1VariantDeleteResponse interface {
	ImplementsImagesV1VariantDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1VariantDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V1VariantEditResponse struct {
	Variant V1VariantEditResponseVariant `json:"variant"`
	JSON    v1VariantEditResponseJSON    `json:"-"`
}

// v1VariantEditResponseJSON contains the JSON metadata for the struct
// [V1VariantEditResponse]
type v1VariantEditResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantEditResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantEditResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                             `json:"neverRequireSignedURLs"`
	JSON                   v1VariantEditResponseVariantJSON `json:"-"`
}

// v1VariantEditResponseVariantJSON contains the JSON metadata for the struct
// [V1VariantEditResponseVariant]
type v1VariantEditResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantEditResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantEditResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantEditResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantEditResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                 `json:"width,required"`
	JSON  v1VariantEditResponseVariantOptionsJSON `json:"-"`
}

// v1VariantEditResponseVariantOptionsJSON contains the JSON metadata for the
// struct [V1VariantEditResponseVariantOptions]
type v1VariantEditResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantEditResponseVariantOptionsFit string

const (
	V1VariantEditResponseVariantOptionsFitScaleDown V1VariantEditResponseVariantOptionsFit = "scale-down"
	V1VariantEditResponseVariantOptionsFitContain   V1VariantEditResponseVariantOptionsFit = "contain"
	V1VariantEditResponseVariantOptionsFitCover     V1VariantEditResponseVariantOptionsFit = "cover"
	V1VariantEditResponseVariantOptionsFitCrop      V1VariantEditResponseVariantOptionsFit = "crop"
	V1VariantEditResponseVariantOptionsFitPad       V1VariantEditResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type V1VariantEditResponseVariantOptionsMetadata string

const (
	V1VariantEditResponseVariantOptionsMetadataKeep      V1VariantEditResponseVariantOptionsMetadata = "keep"
	V1VariantEditResponseVariantOptionsMetadataCopyright V1VariantEditResponseVariantOptionsMetadata = "copyright"
	V1VariantEditResponseVariantOptionsMetadataNone      V1VariantEditResponseVariantOptionsMetadata = "none"
)

type V1VariantGetResponse struct {
	Variant V1VariantGetResponseVariant `json:"variant"`
	JSON    v1VariantGetResponseJSON    `json:"-"`
}

// v1VariantGetResponseJSON contains the JSON metadata for the struct
// [V1VariantGetResponse]
type v1VariantGetResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantGetResponseVariant struct {
	ID interface{} `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantGetResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                            `json:"neverRequireSignedURLs"`
	JSON                   v1VariantGetResponseVariantJSON `json:"-"`
}

// v1VariantGetResponseVariantJSON contains the JSON metadata for the struct
// [V1VariantGetResponseVariant]
type v1VariantGetResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantGetResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantGetResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantGetResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantGetResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                `json:"width,required"`
	JSON  v1VariantGetResponseVariantOptionsJSON `json:"-"`
}

// v1VariantGetResponseVariantOptionsJSON contains the JSON metadata for the struct
// [V1VariantGetResponseVariantOptions]
type v1VariantGetResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantGetResponseVariantOptionsFit string

const (
	V1VariantGetResponseVariantOptionsFitScaleDown V1VariantGetResponseVariantOptionsFit = "scale-down"
	V1VariantGetResponseVariantOptionsFitContain   V1VariantGetResponseVariantOptionsFit = "contain"
	V1VariantGetResponseVariantOptionsFitCover     V1VariantGetResponseVariantOptionsFit = "cover"
	V1VariantGetResponseVariantOptionsFitCrop      V1VariantGetResponseVariantOptionsFit = "crop"
	V1VariantGetResponseVariantOptionsFitPad       V1VariantGetResponseVariantOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type V1VariantGetResponseVariantOptionsMetadata string

const (
	V1VariantGetResponseVariantOptionsMetadataKeep      V1VariantGetResponseVariantOptionsMetadata = "keep"
	V1VariantGetResponseVariantOptionsMetadataCopyright V1VariantGetResponseVariantOptionsMetadata = "copyright"
	V1VariantGetResponseVariantOptionsMetadataNone      V1VariantGetResponseVariantOptionsMetadata = "none"
)

type V1VariantNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string]      `path:"account_id,required"`
	ID        param.Field[interface{}] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[V1VariantNewParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r V1VariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantNewParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[V1VariantNewParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[V1VariantNewParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r V1VariantNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantNewParamsOptionsFit string

const (
	V1VariantNewParamsOptionsFitScaleDown V1VariantNewParamsOptionsFit = "scale-down"
	V1VariantNewParamsOptionsFitContain   V1VariantNewParamsOptionsFit = "contain"
	V1VariantNewParamsOptionsFitCover     V1VariantNewParamsOptionsFit = "cover"
	V1VariantNewParamsOptionsFitCrop      V1VariantNewParamsOptionsFit = "crop"
	V1VariantNewParamsOptionsFitPad       V1VariantNewParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type V1VariantNewParamsOptionsMetadata string

const (
	V1VariantNewParamsOptionsMetadataKeep      V1VariantNewParamsOptionsMetadata = "keep"
	V1VariantNewParamsOptionsMetadataCopyright V1VariantNewParamsOptionsMetadata = "copyright"
	V1VariantNewParamsOptionsMetadataNone      V1VariantNewParamsOptionsMetadata = "none"
)

type V1VariantNewResponseEnvelope struct {
	Errors   []V1VariantNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantNewResponseEnvelopeMessages `json:"messages,required"`
	Result   V1VariantNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantNewResponseEnvelopeJSON    `json:"-"`
}

// v1VariantNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantNewResponseEnvelope]
type v1VariantNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    v1VariantNewResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1VariantNewResponseEnvelopeErrors]
type v1VariantNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v1VariantNewResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantNewResponseEnvelopeMessages]
type v1VariantNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantNewResponseEnvelopeSuccess bool

const (
	V1VariantNewResponseEnvelopeSuccessTrue V1VariantNewResponseEnvelopeSuccess = true
)

type V1VariantListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantListResponseEnvelope struct {
	Errors   []V1VariantListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantListResponseEnvelopeMessages `json:"messages,required"`
	Result   V1VariantListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantListResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantListResponseEnvelopeJSON    `json:"-"`
}

// v1VariantListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantListResponseEnvelope]
type v1VariantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v1VariantListResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V1VariantListResponseEnvelopeErrors]
type v1VariantListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    v1VariantListResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantListResponseEnvelopeMessages]
type v1VariantListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantListResponseEnvelopeSuccess bool

const (
	V1VariantListResponseEnvelopeSuccessTrue V1VariantListResponseEnvelopeSuccess = true
)

type V1VariantDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantDeleteResponseEnvelope struct {
	Errors   []V1VariantDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V1VariantDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantDeleteResponseEnvelopeJSON    `json:"-"`
}

// v1VariantDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantDeleteResponseEnvelope]
type v1VariantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    v1VariantDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V1VariantDeleteResponseEnvelopeErrors]
type v1VariantDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    v1VariantDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantDeleteResponseEnvelopeMessages]
type v1VariantDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantDeleteResponseEnvelopeSuccess bool

const (
	V1VariantDeleteResponseEnvelopeSuccessTrue V1VariantDeleteResponseEnvelopeSuccess = true
)

type V1VariantEditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[V1VariantEditParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r V1VariantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantEditParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[V1VariantEditParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[V1VariantEditParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r V1VariantEditParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantEditParamsOptionsFit string

const (
	V1VariantEditParamsOptionsFitScaleDown V1VariantEditParamsOptionsFit = "scale-down"
	V1VariantEditParamsOptionsFitContain   V1VariantEditParamsOptionsFit = "contain"
	V1VariantEditParamsOptionsFitCover     V1VariantEditParamsOptionsFit = "cover"
	V1VariantEditParamsOptionsFitCrop      V1VariantEditParamsOptionsFit = "crop"
	V1VariantEditParamsOptionsFitPad       V1VariantEditParamsOptionsFit = "pad"
)

// What EXIF data should be preserved in the output image.
type V1VariantEditParamsOptionsMetadata string

const (
	V1VariantEditParamsOptionsMetadataKeep      V1VariantEditParamsOptionsMetadata = "keep"
	V1VariantEditParamsOptionsMetadataCopyright V1VariantEditParamsOptionsMetadata = "copyright"
	V1VariantEditParamsOptionsMetadataNone      V1VariantEditParamsOptionsMetadata = "none"
)

type V1VariantEditResponseEnvelope struct {
	Errors   []V1VariantEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantEditResponseEnvelopeMessages `json:"messages,required"`
	Result   V1VariantEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantEditResponseEnvelopeJSON    `json:"-"`
}

// v1VariantEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantEditResponseEnvelope]
type v1VariantEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantEditResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v1VariantEditResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V1VariantEditResponseEnvelopeErrors]
type v1VariantEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantEditResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    v1VariantEditResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantEditResponseEnvelopeMessages]
type v1VariantEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantEditResponseEnvelopeSuccess bool

const (
	V1VariantEditResponseEnvelopeSuccessTrue V1VariantEditResponseEnvelopeSuccess = true
)

type V1VariantGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantGetResponseEnvelope struct {
	Errors   []V1VariantGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantGetResponseEnvelopeMessages `json:"messages,required"`
	Result   V1VariantGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantGetResponseEnvelopeJSON    `json:"-"`
}

// v1VariantGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantGetResponseEnvelope]
type v1VariantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    v1VariantGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1VariantGetResponseEnvelopeErrors]
type v1VariantGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v1VariantGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantGetResponseEnvelopeMessages]
type v1VariantGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantGetResponseEnvelopeSuccess bool

const (
	V1VariantGetResponseEnvelopeSuccessTrue V1VariantGetResponseEnvelopeSuccess = true
)
