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

// SettingAutomaticPlatformOptimizationService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingAutomaticPlatformOptimizationService] method instead.
type SettingAutomaticPlatformOptimizationService struct {
	Options []option.RequestOption
}

// NewSettingAutomaticPlatformOptimizationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSettingAutomaticPlatformOptimizationService(opts ...option.RequestOption) (r *SettingAutomaticPlatformOptimizationService) {
	r = &SettingAutomaticPlatformOptimizationService{}
	r.Options = opts
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *SettingAutomaticPlatformOptimizationService) Update(ctx context.Context, zoneID string, body SettingAutomaticPlatformOptimizationUpdateParams, opts ...option.RequestOption) (res *SettingAutomaticPlatformOptimizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticPlatformOptimizationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *SettingAutomaticPlatformOptimizationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingAutomaticPlatformOptimizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticPlatformOptimizationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingAutomaticPlatformOptimizationUpdateResponse struct {
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
	WpPlugin bool                                                   `json:"wp_plugin,required"`
	JSON     settingAutomaticPlatformOptimizationUpdateResponseJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationUpdateResponseJSON contains the JSON
// metadata for the struct [SettingAutomaticPlatformOptimizationUpdateResponse]
type settingAutomaticPlatformOptimizationUpdateResponseJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationGetResponse struct {
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
	WpPlugin bool                                                `json:"wp_plugin,required"`
	JSON     settingAutomaticPlatformOptimizationGetResponseJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationGetResponseJSON contains the JSON metadata
// for the struct [SettingAutomaticPlatformOptimizationGetResponse]
type settingAutomaticPlatformOptimizationGetResponseJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationUpdateParams struct {
	Value param.Field[SettingAutomaticPlatformOptimizationUpdateParamsValue] `json:"value,required"`
}

func (r SettingAutomaticPlatformOptimizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAutomaticPlatformOptimizationUpdateParamsValue struct {
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

func (r SettingAutomaticPlatformOptimizationUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAutomaticPlatformOptimizationUpdateResponseEnvelope struct {
	Errors   []SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                           `json:"success,required"`
	Result  SettingAutomaticPlatformOptimizationUpdateResponse             `json:"result"`
	JSON    settingAutomaticPlatformOptimizationUpdateResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SettingAutomaticPlatformOptimizationUpdateResponseEnvelope]
type settingAutomaticPlatformOptimizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    settingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrors]
type settingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    settingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessages]
type settingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationGetResponseEnvelope struct {
	Errors   []SettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                        `json:"success,required"`
	Result  SettingAutomaticPlatformOptimizationGetResponse             `json:"result"`
	JSON    settingAutomaticPlatformOptimizationGetResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SettingAutomaticPlatformOptimizationGetResponseEnvelope]
type settingAutomaticPlatformOptimizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    settingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors]
type settingAutomaticPlatformOptimizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    settingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages]
type settingAutomaticPlatformOptimizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
