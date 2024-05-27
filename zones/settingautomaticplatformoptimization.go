// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingAutomaticPlatformOptimizationService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAutomaticPlatformOptimizationService] method instead.
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
func (r *SettingAutomaticPlatformOptimizationService) Edit(ctx context.Context, params SettingAutomaticPlatformOptimizationEditParams, opts ...option.RequestOption) (res *AutomaticPlatformOptimization, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticPlatformOptimizationEditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
func (r *SettingAutomaticPlatformOptimizationService) Get(ctx context.Context, query SettingAutomaticPlatformOptimizationGetParams, opts ...option.RequestOption) (res *AutomaticPlatformOptimization, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticPlatformOptimizationGetResponseEnvelope
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AutomaticPlatformOptimization struct {
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
	WpPlugin bool                              `json:"wp_plugin,required"`
	JSON     automaticPlatformOptimizationJSON `json:"-"`
}

// automaticPlatformOptimizationJSON contains the JSON metadata for the struct
// [AutomaticPlatformOptimization]
type automaticPlatformOptimizationJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
}

type AutomaticPlatformOptimizationParam struct {
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

func (r AutomaticPlatformOptimizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAutomaticPlatformOptimizationEditParams struct {
	// Identifier
	ZoneID param.Field[string]                             `path:"zone_id,required"`
	Value  param.Field[AutomaticPlatformOptimizationParam] `json:"value,required"`
}

func (r SettingAutomaticPlatformOptimizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAutomaticPlatformOptimizationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                         `json:"success,required"`
	Result  AutomaticPlatformOptimization                                `json:"result"`
	JSON    settingAutomaticPlatformOptimizationEditResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticPlatformOptimizationEditResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SettingAutomaticPlatformOptimizationEditResponseEnvelope]
type settingAutomaticPlatformOptimizationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticPlatformOptimizationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticPlatformOptimizationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAutomaticPlatformOptimizationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAutomaticPlatformOptimizationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                        `json:"success,required"`
	Result  AutomaticPlatformOptimization                               `json:"result"`
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

func (r settingAutomaticPlatformOptimizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
