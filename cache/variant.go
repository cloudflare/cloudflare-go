// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
)

// VariantService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVariantService] method instead.
type VariantService struct {
	Options []option.RequestOption
}

// NewVariantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVariantService(opts ...option.RequestOption) (r *VariantService) {
	r = &VariantService{}
	r.Options = opts
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *VariantService) Delete(ctx context.Context, body VariantDeleteParams, opts ...option.RequestOption) (res *CacheVariant, err error) {
	var env VariantDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/variants", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *VariantService) Edit(ctx context.Context, params VariantEditParams, opts ...option.RequestOption) (res *VariantEditResponse, err error) {
	var env VariantEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/variants", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *VariantService) Get(ctx context.Context, query VariantGetParams, opts ...option.RequestOption) (res *VariantGetResponse, err error) {
	var env VariantGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/variants", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type CacheVariant struct {
	// ID of the zone setting.
	ID CacheVariantIdentifier `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,required,nullable" format:"date-time"`
	JSON       cacheVariantJSON `json:"-"`
}

// cacheVariantJSON contains the JSON metadata for the struct [CacheVariant]
type cacheVariantJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheVariantJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheVariantIdentifier string

const (
	CacheVariantIdentifierVariants CacheVariantIdentifier = "variants"
)

func (r CacheVariantIdentifier) IsKnown() bool {
	switch r {
	case CacheVariantIdentifierVariants:
		return true
	}
	return false
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type VariantEditResponse struct {
	// ID of the zone setting.
	ID CacheVariantIdentifier `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the zone setting.
	Value VariantEditResponseValue `json:"value,required"`
	JSON  variantEditResponseJSON  `json:"-"`
}

// variantEditResponseJSON contains the JSON metadata for the struct
// [VariantEditResponse]
type variantEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantEditResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type VariantEditResponseValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	AVIF []string `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	BMP []string `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	GIF []string `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	JP2 []string `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	JPEG []string `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	JPG []string `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	JPG2 []string `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	PNG []string `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	TIF []string `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	TIFF []string `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	WebP []string                     `json:"webp"`
	JSON variantEditResponseValueJSON `json:"-"`
}

// variantEditResponseValueJSON contains the JSON metadata for the struct
// [VariantEditResponseValue]
type variantEditResponseValueJSON struct {
	AVIF        apijson.Field
	BMP         apijson.Field
	GIF         apijson.Field
	JP2         apijson.Field
	JPEG        apijson.Field
	JPG         apijson.Field
	JPG2        apijson.Field
	PNG         apijson.Field
	TIF         apijson.Field
	TIFF        apijson.Field
	WebP        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type VariantGetResponse struct {
	// ID of the zone setting.
	ID CacheVariantIdentifier `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the zone setting.
	Value VariantGetResponseValue `json:"value,required"`
	JSON  variantGetResponseJSON  `json:"-"`
}

// variantGetResponseJSON contains the JSON metadata for the struct
// [VariantGetResponse]
type variantGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantGetResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type VariantGetResponseValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	AVIF []string `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	BMP []string `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	GIF []string `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	JP2 []string `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	JPEG []string `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	JPG []string `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	JPG2 []string `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	PNG []string `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	TIF []string `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	TIFF []string `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	WebP []string                    `json:"webp"`
	JSON variantGetResponseValueJSON `json:"-"`
}

// variantGetResponseValueJSON contains the JSON metadata for the struct
// [VariantGetResponseValue]
type variantGetResponseValueJSON struct {
	AVIF        apijson.Field
	BMP         apijson.Field
	GIF         apijson.Field
	JP2         apijson.Field
	JPEG        apijson.Field
	JPG         apijson.Field
	JPG2        apijson.Field
	PNG         apijson.Field
	TIF         apijson.Field
	TIFF        apijson.Field
	WebP        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantGetResponseValueJSON) RawJSON() string {
	return r.raw
}

type VariantDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type VariantDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CacheVariant `json:"result,required"`
	// Whether the API call was successful
	Success VariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    variantDeleteResponseEnvelopeJSON    `json:"-"`
}

// variantDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [VariantDeleteResponseEnvelope]
type variantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VariantDeleteResponseEnvelopeSuccess bool

const (
	VariantDeleteResponseEnvelopeSuccessTrue VariantDeleteResponseEnvelopeSuccess = true
)

func (r VariantDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VariantDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type VariantEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[VariantEditParamsValue] `json:"value,required"`
}

func (r VariantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type VariantEditParamsValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	AVIF param.Field[[]string] `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	BMP param.Field[[]string] `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	GIF param.Field[[]string] `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	JP2 param.Field[[]string] `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	JPEG param.Field[[]string] `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	JPG param.Field[[]string] `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	JPG2 param.Field[[]string] `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	PNG param.Field[[]string] `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	TIF param.Field[[]string] `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	TIFF param.Field[[]string] `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	WebP param.Field[[]string] `json:"webp"`
}

func (r VariantEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VariantEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result VariantEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success VariantEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    variantEditResponseEnvelopeJSON    `json:"-"`
}

// variantEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [VariantEditResponseEnvelope]
type variantEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VariantEditResponseEnvelopeSuccess bool

const (
	VariantEditResponseEnvelopeSuccessTrue VariantEditResponseEnvelopeSuccess = true
)

func (r VariantEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VariantEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type VariantGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type VariantGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result VariantGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success VariantGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    variantGetResponseEnvelopeJSON    `json:"-"`
}

// variantGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VariantGetResponseEnvelope]
type variantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VariantGetResponseEnvelopeSuccess bool

const (
	VariantGetResponseEnvelopeSuccessTrue VariantGetResponseEnvelopeSuccess = true
)

func (r VariantGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VariantGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
