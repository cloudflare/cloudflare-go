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

// ZoneSettingAutomaticPlatformOptimizationService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingAutomaticPlatformOptimizationService] method instead.
type ZoneSettingAutomaticPlatformOptimizationService struct {
	Options []option.RequestOption
}

// NewZoneSettingAutomaticPlatformOptimizationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSettingAutomaticPlatformOptimizationService(opts ...option.RequestOption) (r *ZoneSettingAutomaticPlatformOptimizationService) {
	r = &ZoneSettingAutomaticPlatformOptimizationService{}
	r.Options = opts
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *ZoneSettingAutomaticPlatformOptimizationService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingAutomaticPlatformOptimizationUpdateParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticPlatformOptimizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *ZoneSettingAutomaticPlatformOptimizationService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAutomaticPlatformOptimizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingAutomaticPlatformOptimizationUpdateResponse struct {
	Errors   []ZoneSettingAutomaticPlatformOptimizationUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingAutomaticPlatformOptimizationUpdateResponseMessage `json:"messages"`
	Result   ZoneSettingAutomaticPlatformOptimizationUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                                       `json:"success"`
	JSON    zoneSettingAutomaticPlatformOptimizationUpdateResponseJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationUpdateResponseJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticPlatformOptimizationUpdateResponse]
type zoneSettingAutomaticPlatformOptimizationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationUpdateResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationUpdateResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationUpdateResponseError]
type zoneSettingAutomaticPlatformOptimizationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationUpdateResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationUpdateResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationUpdateResponseMessage]
type zoneSettingAutomaticPlatformOptimizationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationUpdateResponseResult struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType bool `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf bool `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled bool `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames []string `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress bool `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin bool                                                             `json:"wp_plugin,required"`
	JSON     zoneSettingAutomaticPlatformOptimizationUpdateResponseResultJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationUpdateResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationUpdateResponseResult]
type zoneSettingAutomaticPlatformOptimizationUpdateResponseResultJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationListResponse struct {
	Errors   []ZoneSettingAutomaticPlatformOptimizationListResponseError   `json:"errors"`
	Messages []ZoneSettingAutomaticPlatformOptimizationListResponseMessage `json:"messages"`
	Result   ZoneSettingAutomaticPlatformOptimizationListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                                     `json:"success"`
	JSON    zoneSettingAutomaticPlatformOptimizationListResponseJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationListResponseJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticPlatformOptimizationListResponse]
type zoneSettingAutomaticPlatformOptimizationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationListResponseError struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationListResponseErrorJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationListResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationListResponseError]
type zoneSettingAutomaticPlatformOptimizationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationListResponseMessage struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationListResponseMessageJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationListResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationListResponseMessage]
type zoneSettingAutomaticPlatformOptimizationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationListResponseResult struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType bool `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf bool `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled bool `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames []string `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress bool `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin bool                                                           `json:"wp_plugin,required"`
	JSON     zoneSettingAutomaticPlatformOptimizationListResponseResultJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationListResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationListResponseResult]
type zoneSettingAutomaticPlatformOptimizationListResponseResultJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticPlatformOptimizationUpdateParams struct {
	Value param.Field[ZoneSettingAutomaticPlatformOptimizationUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingAutomaticPlatformOptimizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingAutomaticPlatformOptimizationUpdateParamsValue struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf param.Field[bool] `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames param.Field[[]string] `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress param.Field[bool] `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin param.Field[bool] `json:"wp_plugin,required"`
}

func (r ZoneSettingAutomaticPlatformOptimizationUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
