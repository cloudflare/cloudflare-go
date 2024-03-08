// File generated from our OpenAPI spec by Stainless.

package zones

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

// SettingRocketLoaderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingRocketLoaderService]
// method instead.
type SettingRocketLoaderService struct {
	Options []option.RequestOption
}

// NewSettingRocketLoaderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingRocketLoaderService(opts ...option.RequestOption) (r *SettingRocketLoaderService) {
	r = &SettingRocketLoaderService{}
	r.Options = opts
	return
}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
func (r *SettingRocketLoaderService) Edit(ctx context.Context, params SettingRocketLoaderEditParams, opts ...option.RequestOption) (res *SettingRocketLoaderEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingRocketLoaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
func (r *SettingRocketLoaderService) Get(ctx context.Context, query SettingRocketLoaderGetParams, opts ...option.RequestOption) (res *SettingRocketLoaderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingRocketLoaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type SettingRocketLoaderEditResponse struct {
	// ID of the zone setting.
	ID SettingRocketLoaderEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingRocketLoaderEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingRocketLoaderEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingRocketLoaderEditResponseJSON `json:"-"`
}

// settingRocketLoaderEditResponseJSON contains the JSON metadata for the struct
// [SettingRocketLoaderEditResponse]
type settingRocketLoaderEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingRocketLoaderEditResponseID string

const (
	SettingRocketLoaderEditResponseIDRocketLoader SettingRocketLoaderEditResponseID = "rocket_loader"
)

// Current value of the zone setting.
type SettingRocketLoaderEditResponseValue string

const (
	SettingRocketLoaderEditResponseValueOn  SettingRocketLoaderEditResponseValue = "on"
	SettingRocketLoaderEditResponseValueOff SettingRocketLoaderEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingRocketLoaderEditResponseEditable bool

const (
	SettingRocketLoaderEditResponseEditableTrue  SettingRocketLoaderEditResponseEditable = true
	SettingRocketLoaderEditResponseEditableFalse SettingRocketLoaderEditResponseEditable = false
)

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type SettingRocketLoaderGetResponse struct {
	// ID of the zone setting.
	ID SettingRocketLoaderGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingRocketLoaderGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingRocketLoaderGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingRocketLoaderGetResponseJSON `json:"-"`
}

// settingRocketLoaderGetResponseJSON contains the JSON metadata for the struct
// [SettingRocketLoaderGetResponse]
type settingRocketLoaderGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingRocketLoaderGetResponseID string

const (
	SettingRocketLoaderGetResponseIDRocketLoader SettingRocketLoaderGetResponseID = "rocket_loader"
)

// Current value of the zone setting.
type SettingRocketLoaderGetResponseValue string

const (
	SettingRocketLoaderGetResponseValueOn  SettingRocketLoaderGetResponseValue = "on"
	SettingRocketLoaderGetResponseValueOff SettingRocketLoaderGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingRocketLoaderGetResponseEditable bool

const (
	SettingRocketLoaderGetResponseEditableTrue  SettingRocketLoaderGetResponseEditable = true
	SettingRocketLoaderGetResponseEditableFalse SettingRocketLoaderGetResponseEditable = false
)

type SettingRocketLoaderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
	// prioritises rendering your content while loading your site's Javascript
	// asynchronously. Turning on Rocket Loader will immediately improve a web page's
	// rendering time sometimes measured as Time to First Paint (TTFP), and also the
	// `window.onload` time (assuming there is JavaScript on the page). This can have a
	// positive impact on your Google search ranking. When turned on, Rocket Loader
	// will automatically defer the loading of all Javascript referenced in your HTML,
	// with no configuration required. Refer to
	// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
	// for more information.
	Value param.Field[SettingRocketLoaderEditParamsValue] `json:"value,required"`
}

func (r SettingRocketLoaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type SettingRocketLoaderEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingRocketLoaderEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingRocketLoaderEditParamsValueValue] `json:"value,required"`
}

func (r SettingRocketLoaderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingRocketLoaderEditParamsValueID string

const (
	SettingRocketLoaderEditParamsValueIDRocketLoader SettingRocketLoaderEditParamsValueID = "rocket_loader"
)

// Current value of the zone setting.
type SettingRocketLoaderEditParamsValueValue string

const (
	SettingRocketLoaderEditParamsValueValueOn  SettingRocketLoaderEditParamsValueValue = "on"
	SettingRocketLoaderEditParamsValueValueOff SettingRocketLoaderEditParamsValueValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingRocketLoaderEditParamsValueEditable bool

const (
	SettingRocketLoaderEditParamsValueEditableTrue  SettingRocketLoaderEditParamsValueEditable = true
	SettingRocketLoaderEditParamsValueEditableFalse SettingRocketLoaderEditParamsValueEditable = false
)

type SettingRocketLoaderEditResponseEnvelope struct {
	Errors   []SettingRocketLoaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingRocketLoaderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
	// prioritises rendering your content while loading your site's Javascript
	// asynchronously. Turning on Rocket Loader will immediately improve a web page's
	// rendering time sometimes measured as Time to First Paint (TTFP), and also the
	// `window.onload` time (assuming there is JavaScript on the page). This can have a
	// positive impact on your Google search ranking. When turned on, Rocket Loader
	// will automatically defer the loading of all Javascript referenced in your HTML,
	// with no configuration required. Refer to
	// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
	// for more information.
	Result SettingRocketLoaderEditResponse             `json:"result"`
	JSON   settingRocketLoaderEditResponseEnvelopeJSON `json:"-"`
}

// settingRocketLoaderEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingRocketLoaderEditResponseEnvelope]
type settingRocketLoaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingRocketLoaderEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingRocketLoaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingRocketLoaderEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingRocketLoaderEditResponseEnvelopeErrors]
type settingRocketLoaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingRocketLoaderEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingRocketLoaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingRocketLoaderEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingRocketLoaderEditResponseEnvelopeMessages]
type settingRocketLoaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingRocketLoaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingRocketLoaderGetResponseEnvelope struct {
	Errors   []SettingRocketLoaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingRocketLoaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
	// prioritises rendering your content while loading your site's Javascript
	// asynchronously. Turning on Rocket Loader will immediately improve a web page's
	// rendering time sometimes measured as Time to First Paint (TTFP), and also the
	// `window.onload` time (assuming there is JavaScript on the page). This can have a
	// positive impact on your Google search ranking. When turned on, Rocket Loader
	// will automatically defer the loading of all Javascript referenced in your HTML,
	// with no configuration required. Refer to
	// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
	// for more information.
	Result SettingRocketLoaderGetResponse             `json:"result"`
	JSON   settingRocketLoaderGetResponseEnvelopeJSON `json:"-"`
}

// settingRocketLoaderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingRocketLoaderGetResponseEnvelope]
type settingRocketLoaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingRocketLoaderGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingRocketLoaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingRocketLoaderGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingRocketLoaderGetResponseEnvelopeErrors]
type settingRocketLoaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingRocketLoaderGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingRocketLoaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingRocketLoaderGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingRocketLoaderGetResponseEnvelopeMessages]
type settingRocketLoaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingRocketLoaderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
