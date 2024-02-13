// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CachVariantService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCachVariantService] method
// instead.
type CachVariantService struct {
	Options []option.RequestOption
}

// NewCachVariantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCachVariantService(opts ...option.RequestOption) (r *CachVariantService) {
	r = &CachVariantService{}
	r.Options = opts
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *CachVariantService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CachVariantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachVariantListResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
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
func (r *CachVariantService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CachVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachVariantDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
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
func (r *CachVariantService) ZoneCacheSettingsChangeVariantsSetting(ctx context.Context, zoneID string, body CachVariantZoneCacheSettingsChangeVariantsSettingParams, opts ...option.RequestOption) (res *CachVariantZoneCacheSettingsChangeVariantsSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
type CachVariantListResponse struct {
	// ID of the zone setting.
	ID CachVariantListResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the zone setting.
	Value CachVariantListResponseValue `json:"value,required"`
	JSON  cachVariantListResponseJSON  `json:"-"`
}

// cachVariantListResponseJSON contains the JSON metadata for the struct
// [CachVariantListResponse]
type cachVariantListResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CachVariantListResponseID string

const (
	CachVariantListResponseIDVariants CachVariantListResponseID = "variants"
)

// Value of the zone setting.
type CachVariantListResponseValue struct {
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
	Webp []interface{}                    `json:"webp"`
	JSON cachVariantListResponseValueJSON `json:"-"`
}

// cachVariantListResponseValueJSON contains the JSON metadata for the struct
// [CachVariantListResponseValue]
type cachVariantListResponseValueJSON struct {
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

func (r *CachVariantListResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type CachVariantDeleteResponse struct {
	// ID of the zone setting.
	ID CachVariantDeleteResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,required,nullable" format:"date-time"`
	JSON       cachVariantDeleteResponseJSON `json:"-"`
}

// cachVariantDeleteResponseJSON contains the JSON metadata for the struct
// [CachVariantDeleteResponse]
type cachVariantDeleteResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CachVariantDeleteResponseID string

const (
	CachVariantDeleteResponseIDVariants CachVariantDeleteResponseID = "variants"
)

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type CachVariantZoneCacheSettingsChangeVariantsSettingResponse struct {
	// ID of the zone setting.
	ID CachVariantZoneCacheSettingsChangeVariantsSettingResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the zone setting.
	Value CachVariantZoneCacheSettingsChangeVariantsSettingResponseValue `json:"value,required"`
	JSON  cachVariantZoneCacheSettingsChangeVariantsSettingResponseJSON  `json:"-"`
}

// cachVariantZoneCacheSettingsChangeVariantsSettingResponseJSON contains the JSON
// metadata for the struct
// [CachVariantZoneCacheSettingsChangeVariantsSettingResponse]
type cachVariantZoneCacheSettingsChangeVariantsSettingResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantZoneCacheSettingsChangeVariantsSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CachVariantZoneCacheSettingsChangeVariantsSettingResponseID string

const (
	CachVariantZoneCacheSettingsChangeVariantsSettingResponseIDVariants CachVariantZoneCacheSettingsChangeVariantsSettingResponseID = "variants"
)

// Value of the zone setting.
type CachVariantZoneCacheSettingsChangeVariantsSettingResponseValue struct {
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
	Webp []interface{}                                                      `json:"webp"`
	JSON cachVariantZoneCacheSettingsChangeVariantsSettingResponseValueJSON `json:"-"`
}

// cachVariantZoneCacheSettingsChangeVariantsSettingResponseValueJSON contains the
// JSON metadata for the struct
// [CachVariantZoneCacheSettingsChangeVariantsSettingResponseValue]
type cachVariantZoneCacheSettingsChangeVariantsSettingResponseValueJSON struct {
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

func (r *CachVariantZoneCacheSettingsChangeVariantsSettingResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantListResponseEnvelope struct {
	Errors   []CachVariantListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachVariantListResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CachVariantListResponse `json:"result,required"`
	// Whether the API call was successful
	Success CachVariantListResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachVariantListResponseEnvelopeJSON    `json:"-"`
}

// cachVariantListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CachVariantListResponseEnvelope]
type cachVariantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    cachVariantListResponseEnvelopeErrorsJSON `json:"-"`
}

// cachVariantListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CachVariantListResponseEnvelopeErrors]
type cachVariantListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cachVariantListResponseEnvelopeMessagesJSON `json:"-"`
}

// cachVariantListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CachVariantListResponseEnvelopeMessages]
type cachVariantListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachVariantListResponseEnvelopeSuccess bool

const (
	CachVariantListResponseEnvelopeSuccessTrue CachVariantListResponseEnvelopeSuccess = true
)

type CachVariantDeleteResponseEnvelope struct {
	Errors   []CachVariantDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachVariantDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CachVariantDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success CachVariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachVariantDeleteResponseEnvelopeJSON    `json:"-"`
}

// cachVariantDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CachVariantDeleteResponseEnvelope]
type cachVariantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cachVariantDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cachVariantDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CachVariantDeleteResponseEnvelopeErrors]
type cachVariantDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    cachVariantDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cachVariantDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CachVariantDeleteResponseEnvelopeMessages]
type cachVariantDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachVariantDeleteResponseEnvelopeSuccess bool

const (
	CachVariantDeleteResponseEnvelopeSuccessTrue CachVariantDeleteResponseEnvelopeSuccess = true
)

type CachVariantZoneCacheSettingsChangeVariantsSettingParams struct {
	// Value of the zone setting.
	Value param.Field[CachVariantZoneCacheSettingsChangeVariantsSettingParamsValue] `json:"value,required"`
}

func (r CachVariantZoneCacheSettingsChangeVariantsSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type CachVariantZoneCacheSettingsChangeVariantsSettingParamsValue struct {
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

func (r CachVariantZoneCacheSettingsChangeVariantsSettingParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelope struct {
	Errors   []CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CachVariantZoneCacheSettingsChangeVariantsSettingResponse `json:"result,required"`
	// Whether the API call was successful
	Success CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeJSON    `json:"-"`
}

// cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelope]
type cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrors]
type cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessages]
type cachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeSuccess bool

const (
	CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeSuccessTrue CachVariantZoneCacheSettingsChangeVariantsSettingResponseEnvelopeSuccess = true
)
