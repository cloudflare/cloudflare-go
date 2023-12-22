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

// ZoneCachVariantService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCachVariantService] method
// instead.
type ZoneCachVariantService struct {
	Options []option.RequestOption
}

// NewZoneCachVariantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCachVariantService(opts ...option.RequestOption) (r *ZoneCachVariantService) {
	r = &ZoneCachVariantService{}
	r.Options = opts
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *ZoneCachVariantService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCachVariantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/variants", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *ZoneCachVariantService) Delete(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCachVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/variants", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *ZoneCachVariantService) ZoneCacheSettingsChangeVariantsSetting(ctx context.Context, zoneIdentifier string, body ZoneCachVariantZoneCacheSettingsChangeVariantsSettingParams, opts ...option.RequestOption) (res *ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/variants", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneCachVariantListResponse struct {
	Errors   []ZoneCachVariantListResponseError   `json:"errors"`
	Messages []ZoneCachVariantListResponseMessage `json:"messages"`
	Result   ZoneCachVariantListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCachVariantListResponseSuccess `json:"success"`
	JSON    zoneCachVariantListResponseJSON    `json:"-"`
}

// zoneCachVariantListResponseJSON contains the JSON metadata for the struct
// [ZoneCachVariantListResponse]
type zoneCachVariantListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneCachVariantListResponseErrorJSON `json:"-"`
}

// zoneCachVariantListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneCachVariantListResponseError]
type zoneCachVariantListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneCachVariantListResponseMessageJSON `json:"-"`
}

// zoneCachVariantListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneCachVariantListResponseMessage]
type zoneCachVariantListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantListResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachVariantListResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneCachVariantListResponseResultValue `json:"value"`
	JSON  zoneCachVariantListResponseResultJSON  `json:"-"`
}

// zoneCachVariantListResponseResultJSON contains the JSON metadata for the struct
// [ZoneCachVariantListResponseResult]
type zoneCachVariantListResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachVariantListResponseResultID string

const (
	ZoneCachVariantListResponseResultIDVariants ZoneCachVariantListResponseResultID = "variants"
)

// Value of the zone setting.
type ZoneCachVariantListResponseResultValue struct {
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
	Webp []interface{}                              `json:"webp"`
	JSON zoneCachVariantListResponseResultValueJSON `json:"-"`
}

// zoneCachVariantListResponseResultValueJSON contains the JSON metadata for the
// struct [ZoneCachVariantListResponseResultValue]
type zoneCachVariantListResponseResultValueJSON struct {
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

func (r *ZoneCachVariantListResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCachVariantListResponseSuccess bool

const (
	ZoneCachVariantListResponseSuccessTrue ZoneCachVariantListResponseSuccess = true
)

type ZoneCachVariantDeleteResponse struct {
	Errors   []ZoneCachVariantDeleteResponseError   `json:"errors"`
	Messages []ZoneCachVariantDeleteResponseMessage `json:"messages"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result ZoneCachVariantDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneCachVariantDeleteResponseSuccess `json:"success"`
	JSON    zoneCachVariantDeleteResponseJSON    `json:"-"`
}

// zoneCachVariantDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneCachVariantDeleteResponse]
type zoneCachVariantDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneCachVariantDeleteResponseErrorJSON `json:"-"`
}

// zoneCachVariantDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneCachVariantDeleteResponseError]
type zoneCachVariantDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneCachVariantDeleteResponseMessageJSON `json:"-"`
}

// zoneCachVariantDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneCachVariantDeleteResponseMessage]
type zoneCachVariantDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type ZoneCachVariantDeleteResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachVariantDeleteResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCachVariantDeleteResponseResultJSON `json:"-"`
}

// zoneCachVariantDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneCachVariantDeleteResponseResult]
type zoneCachVariantDeleteResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachVariantDeleteResponseResultID string

const (
	ZoneCachVariantDeleteResponseResultIDVariants ZoneCachVariantDeleteResponseResultID = "variants"
)

// Whether the API call was successful
type ZoneCachVariantDeleteResponseSuccess bool

const (
	ZoneCachVariantDeleteResponseSuccessTrue ZoneCachVariantDeleteResponseSuccess = true
)

type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponse struct {
	Errors   []ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseError   `json:"errors"`
	Messages []ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessage `json:"messages"`
	Result   ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseSuccess `json:"success"`
	JSON    zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseJSON    `json:"-"`
}

// zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseJSON contains the
// JSON metadata for the struct
// [ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponse]
type zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseError struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseErrorJSON `json:"-"`
}

// zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseError]
type zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessage struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessageJSON `json:"-"`
}

// zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessage]
type zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValue `json:"value"`
	JSON  zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultJSON  `json:"-"`
}

// zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResult]
type zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultID string

const (
	ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultIDVariants ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultID = "variants"
)

// Value of the zone setting.
type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValue struct {
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
	Webp []interface{}                                                                `json:"webp"`
	JSON zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValueJSON `json:"-"`
}

// zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValueJSON
// contains the JSON metadata for the struct
// [ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValue]
type zoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValueJSON struct {
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

func (r *ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseSuccess bool

const (
	ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseSuccessTrue ZoneCachVariantZoneCacheSettingsChangeVariantsSettingResponseSuccess = true
)

type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneCachVariantZoneCacheSettingsChangeVariantsSettingParamsValue] `json:"value,required"`
}

func (r ZoneCachVariantZoneCacheSettingsChangeVariantsSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneCachVariantZoneCacheSettingsChangeVariantsSettingParamsValue struct {
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

func (r ZoneCachVariantZoneCacheSettingsChangeVariantsSettingParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
