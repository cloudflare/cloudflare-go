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
func (r *ZoneSettingAutomaticPlatformOptimizationService) Edit(ctx context.Context, params ZoneSettingAutomaticPlatformOptimizationEditParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticPlatformOptimizationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *ZoneSettingAutomaticPlatformOptimizationService) Get(ctx context.Context, query ZoneSettingAutomaticPlatformOptimizationGetParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticPlatformOptimizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneSettingAutomaticPlatformOptimizationEditResponse struct {
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
	WpPlugin bool                                                     `json:"wp_plugin,required"`
	JSON     zoneSettingAutomaticPlatformOptimizationEditResponseJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationEditResponseJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticPlatformOptimizationEditResponse]
type zoneSettingAutomaticPlatformOptimizationEditResponseJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationGetResponse struct {
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
	WpPlugin bool                                                    `json:"wp_plugin,required"`
	JSON     zoneSettingAutomaticPlatformOptimizationGetResponseJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationGetResponseJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticPlatformOptimizationGetResponse]
type zoneSettingAutomaticPlatformOptimizationGetResponseJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationEditParams struct {
	// Identifier
	ZoneID param.Field[string]                                                  `path:"zone_id,required"`
	Value  param.Field[ZoneSettingAutomaticPlatformOptimizationEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingAutomaticPlatformOptimizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingAutomaticPlatformOptimizationEditParamsValue struct {
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

func (r ZoneSettingAutomaticPlatformOptimizationEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelope struct {
	Errors   []ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                             `json:"success,required"`
	Result  ZoneSettingAutomaticPlatformOptimizationEditResponse             `json:"result"`
	JSON    zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelope]
type zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrors]
type zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessages]
type zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelope struct {
	Errors   []ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                            `json:"success,required"`
	Result  ZoneSettingAutomaticPlatformOptimizationGetResponse             `json:"result"`
	JSON    zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelope]
type zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors]
type zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages]
type zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
