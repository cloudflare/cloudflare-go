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
func (r *SettingRocketLoaderService) Update(ctx context.Context, zoneID string, body SettingRocketLoaderUpdateParams, opts ...option.RequestOption) (res *SettingRocketLoaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingRocketLoaderUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingRocketLoaderService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingRocketLoaderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingRocketLoaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type SettingRocketLoaderUpdateResponse struct {
	// ID of the zone setting.
	ID SettingRocketLoaderUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingRocketLoaderUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingRocketLoaderUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingRocketLoaderUpdateResponseJSON `json:"-"`
}

// settingRocketLoaderUpdateResponseJSON contains the JSON metadata for the struct
// [SettingRocketLoaderUpdateResponse]
type settingRocketLoaderUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingRocketLoaderUpdateResponseID string

const (
	SettingRocketLoaderUpdateResponseIDRocketLoader SettingRocketLoaderUpdateResponseID = "rocket_loader"
)

// Current value of the zone setting.
type SettingRocketLoaderUpdateResponseValue string

const (
	SettingRocketLoaderUpdateResponseValueOn  SettingRocketLoaderUpdateResponseValue = "on"
	SettingRocketLoaderUpdateResponseValueOff SettingRocketLoaderUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingRocketLoaderUpdateResponseEditable bool

const (
	SettingRocketLoaderUpdateResponseEditableTrue  SettingRocketLoaderUpdateResponseEditable = true
	SettingRocketLoaderUpdateResponseEditableFalse SettingRocketLoaderUpdateResponseEditable = false
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

type SettingRocketLoaderUpdateParams struct {
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
	Value param.Field[SettingRocketLoaderUpdateParamsValue] `json:"value,required"`
}

func (r SettingRocketLoaderUpdateParams) MarshalJSON() (data []byte, err error) {
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
type SettingRocketLoaderUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingRocketLoaderUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingRocketLoaderUpdateParamsValueValue] `json:"value,required"`
}

func (r SettingRocketLoaderUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingRocketLoaderUpdateParamsValueID string

const (
	SettingRocketLoaderUpdateParamsValueIDRocketLoader SettingRocketLoaderUpdateParamsValueID = "rocket_loader"
)

// Current value of the zone setting.
type SettingRocketLoaderUpdateParamsValueValue string

const (
	SettingRocketLoaderUpdateParamsValueValueOn  SettingRocketLoaderUpdateParamsValueValue = "on"
	SettingRocketLoaderUpdateParamsValueValueOff SettingRocketLoaderUpdateParamsValueValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingRocketLoaderUpdateParamsValueEditable bool

const (
	SettingRocketLoaderUpdateParamsValueEditableTrue  SettingRocketLoaderUpdateParamsValueEditable = true
	SettingRocketLoaderUpdateParamsValueEditableFalse SettingRocketLoaderUpdateParamsValueEditable = false
)

type SettingRocketLoaderUpdateResponseEnvelope struct {
	Errors   []SettingRocketLoaderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingRocketLoaderUpdateResponseEnvelopeMessages `json:"messages,required"`
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
	Result SettingRocketLoaderUpdateResponse             `json:"result"`
	JSON   settingRocketLoaderUpdateResponseEnvelopeJSON `json:"-"`
}

// settingRocketLoaderUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingRocketLoaderUpdateResponseEnvelope]
type settingRocketLoaderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingRocketLoaderUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingRocketLoaderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingRocketLoaderUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingRocketLoaderUpdateResponseEnvelopeErrors]
type settingRocketLoaderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingRocketLoaderUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingRocketLoaderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingRocketLoaderUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingRocketLoaderUpdateResponseEnvelopeMessages]
type settingRocketLoaderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingRocketLoaderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
