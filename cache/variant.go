// File generated from our OpenAPI spec by Stainless.

package cache

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// VariantService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewVariantService] method instead.
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
func (r *VariantService) Delete(ctx context.Context, body VariantDeleteParams, opts ...option.RequestOption) (res *VariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VariantDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
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
	opts = append(r.Options[:], opts...)
	var env VariantEditResponseEnvelope
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
	opts = append(r.Options[:], opts...)
	var env VariantGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type VariantDeleteResponse struct {
	// ID of the zone setting.
	ID VariantDeleteResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,required,nullable" format:"date-time"`
	JSON       variantDeleteResponseJSON `json:"-"`
}

// variantDeleteResponseJSON contains the JSON metadata for the struct
// [VariantDeleteResponse]
type variantDeleteResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type VariantDeleteResponseID string

const (
	VariantDeleteResponseIDVariants VariantDeleteResponseID = "variants"
)

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type VariantEditResponse struct {
	// ID of the zone setting.
	ID VariantEditResponseID `json:"id,required"`
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

// ID of the zone setting.
type VariantEditResponseID string

const (
	VariantEditResponseIDVariants VariantEditResponseID = "variants"
)

// Value of the zone setting.
type VariantEditResponseValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif []interface{} `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp []interface{} `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif []interface{} `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 []interface{} `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg []interface{} `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg []interface{} `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 []interface{} `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png []interface{} `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif []interface{} `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff []interface{} `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp []interface{}                `json:"webp"`
	JSON variantEditResponseValueJSON `json:"-"`
}

// variantEditResponseValueJSON contains the JSON metadata for the struct
// [VariantEditResponseValue]
type variantEditResponseValueJSON struct {
	Avif        apijson.Field
	Bmp         apijson.Field
	Gif         apijson.Field
	Jp2         apijson.Field
	Jpeg        apijson.Field
	Jpg         apijson.Field
	Jpg2        apijson.Field
	Png         apijson.Field
	Tif         apijson.Field
	Tiff        apijson.Field
	Webp        apijson.Field
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
	ID VariantGetResponseID `json:"id,required"`
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

// ID of the zone setting.
type VariantGetResponseID string

const (
	VariantGetResponseIDVariants VariantGetResponseID = "variants"
)

// Value of the zone setting.
type VariantGetResponseValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif []interface{} `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp []interface{} `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif []interface{} `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 []interface{} `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg []interface{} `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg []interface{} `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 []interface{} `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png []interface{} `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif []interface{} `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff []interface{} `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp []interface{}               `json:"webp"`
	JSON variantGetResponseValueJSON `json:"-"`
}

// variantGetResponseValueJSON contains the JSON metadata for the struct
// [VariantGetResponseValue]
type variantGetResponseValueJSON struct {
	Avif        apijson.Field
	Bmp         apijson.Field
	Gif         apijson.Field
	Jp2         apijson.Field
	Jpeg        apijson.Field
	Jpg         apijson.Field
	Jpg2        apijson.Field
	Png         apijson.Field
	Tif         apijson.Field
	Tiff        apijson.Field
	Webp        apijson.Field
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
	Errors   []VariantDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VariantDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result VariantDeleteResponse `json:"result,required"`
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

type VariantDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    variantDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// variantDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VariantDeleteResponseEnvelopeErrors]
type variantDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VariantDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    variantDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// variantDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VariantDeleteResponseEnvelopeMessages]
type variantDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VariantDeleteResponseEnvelopeSuccess bool

const (
	VariantDeleteResponseEnvelopeSuccessTrue VariantDeleteResponseEnvelopeSuccess = true
)

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
	Avif param.Field[[]interface{}] `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp param.Field[[]interface{}] `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif param.Field[[]interface{}] `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 param.Field[[]interface{}] `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg param.Field[[]interface{}] `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg param.Field[[]interface{}] `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 param.Field[[]interface{}] `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png param.Field[[]interface{}] `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif param.Field[[]interface{}] `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff param.Field[[]interface{}] `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp param.Field[[]interface{}] `json:"webp"`
}

func (r VariantEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VariantEditResponseEnvelope struct {
	Errors   []VariantEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VariantEditResponseEnvelopeMessages `json:"messages,required"`
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

type VariantEditResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    variantEditResponseEnvelopeErrorsJSON `json:"-"`
}

// variantEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VariantEditResponseEnvelopeErrors]
type variantEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VariantEditResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    variantEditResponseEnvelopeMessagesJSON `json:"-"`
}

// variantEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VariantEditResponseEnvelopeMessages]
type variantEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VariantEditResponseEnvelopeSuccess bool

const (
	VariantEditResponseEnvelopeSuccessTrue VariantEditResponseEnvelopeSuccess = true
)

type VariantGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type VariantGetResponseEnvelope struct {
	Errors   []VariantGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VariantGetResponseEnvelopeMessages `json:"messages,required"`
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

type VariantGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    variantGetResponseEnvelopeErrorsJSON `json:"-"`
}

// variantGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VariantGetResponseEnvelopeErrors]
type variantGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VariantGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    variantGetResponseEnvelopeMessagesJSON `json:"-"`
}

// variantGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [VariantGetResponseEnvelopeMessages]
type variantGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VariantGetResponseEnvelopeSuccess bool

const (
	VariantGetResponseEnvelopeSuccessTrue VariantGetResponseEnvelopeSuccess = true
)
