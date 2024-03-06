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

// ZoneSettingRocketLoaderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingRocketLoaderService] method instead.
type ZoneSettingRocketLoaderService struct {
	Options []option.RequestOption
}

// NewZoneSettingRocketLoaderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingRocketLoaderService(opts ...option.RequestOption) (r *ZoneSettingRocketLoaderService) {
	r = &ZoneSettingRocketLoaderService{}
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
func (r *ZoneSettingRocketLoaderService) Edit(ctx context.Context, params ZoneSettingRocketLoaderEditParams, opts ...option.RequestOption) (res *ZoneSettingRocketLoaderEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingRocketLoaderEditResponseEnvelope
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
func (r *ZoneSettingRocketLoaderService) Get(ctx context.Context, query ZoneSettingRocketLoaderGetParams, opts ...option.RequestOption) (res *ZoneSettingRocketLoaderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingRocketLoaderGetResponseEnvelope
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
type ZoneSettingRocketLoaderEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingRocketLoaderEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingRocketLoaderEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingRocketLoaderEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingRocketLoaderEditResponseJSON `json:"-"`
}

// zoneSettingRocketLoaderEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingRocketLoaderEditResponse]
type zoneSettingRocketLoaderEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingRocketLoaderEditResponseID string

const (
	ZoneSettingRocketLoaderEditResponseIDRocketLoader ZoneSettingRocketLoaderEditResponseID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingRocketLoaderEditResponseValue string

const (
	ZoneSettingRocketLoaderEditResponseValueOn  ZoneSettingRocketLoaderEditResponseValue = "on"
	ZoneSettingRocketLoaderEditResponseValueOff ZoneSettingRocketLoaderEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingRocketLoaderEditResponseEditable bool

const (
	ZoneSettingRocketLoaderEditResponseEditableTrue  ZoneSettingRocketLoaderEditResponseEditable = true
	ZoneSettingRocketLoaderEditResponseEditableFalse ZoneSettingRocketLoaderEditResponseEditable = false
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
type ZoneSettingRocketLoaderGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingRocketLoaderGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingRocketLoaderGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingRocketLoaderGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingRocketLoaderGetResponseJSON `json:"-"`
}

// zoneSettingRocketLoaderGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingRocketLoaderGetResponse]
type zoneSettingRocketLoaderGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingRocketLoaderGetResponseID string

const (
	ZoneSettingRocketLoaderGetResponseIDRocketLoader ZoneSettingRocketLoaderGetResponseID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingRocketLoaderGetResponseValue string

const (
	ZoneSettingRocketLoaderGetResponseValueOn  ZoneSettingRocketLoaderGetResponseValue = "on"
	ZoneSettingRocketLoaderGetResponseValueOff ZoneSettingRocketLoaderGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingRocketLoaderGetResponseEditable bool

const (
	ZoneSettingRocketLoaderGetResponseEditableTrue  ZoneSettingRocketLoaderGetResponseEditable = true
	ZoneSettingRocketLoaderGetResponseEditableFalse ZoneSettingRocketLoaderGetResponseEditable = false
)

type ZoneSettingRocketLoaderEditParams struct {
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
	Value param.Field[ZoneSettingRocketLoaderEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingRocketLoaderEditParams) MarshalJSON() (data []byte, err error) {
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
type ZoneSettingRocketLoaderEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingRocketLoaderEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingRocketLoaderEditParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingRocketLoaderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingRocketLoaderEditParamsValueID string

const (
	ZoneSettingRocketLoaderEditParamsValueIDRocketLoader ZoneSettingRocketLoaderEditParamsValueID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingRocketLoaderEditParamsValueValue string

const (
	ZoneSettingRocketLoaderEditParamsValueValueOn  ZoneSettingRocketLoaderEditParamsValueValue = "on"
	ZoneSettingRocketLoaderEditParamsValueValueOff ZoneSettingRocketLoaderEditParamsValueValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingRocketLoaderEditParamsValueEditable bool

const (
	ZoneSettingRocketLoaderEditParamsValueEditableTrue  ZoneSettingRocketLoaderEditParamsValueEditable = true
	ZoneSettingRocketLoaderEditParamsValueEditableFalse ZoneSettingRocketLoaderEditParamsValueEditable = false
)

type ZoneSettingRocketLoaderEditResponseEnvelope struct {
	Errors   []ZoneSettingRocketLoaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingRocketLoaderEditResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZoneSettingRocketLoaderEditResponse             `json:"result"`
	JSON   zoneSettingRocketLoaderEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingRocketLoaderEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingRocketLoaderEditResponseEnvelope]
type zoneSettingRocketLoaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingRocketLoaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingRocketLoaderEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingRocketLoaderEditResponseEnvelopeErrors]
type zoneSettingRocketLoaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingRocketLoaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingRocketLoaderEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingRocketLoaderEditResponseEnvelopeMessages]
type zoneSettingRocketLoaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingRocketLoaderGetResponseEnvelope struct {
	Errors   []ZoneSettingRocketLoaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingRocketLoaderGetResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZoneSettingRocketLoaderGetResponse             `json:"result"`
	JSON   zoneSettingRocketLoaderGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingRocketLoaderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingRocketLoaderGetResponseEnvelope]
type zoneSettingRocketLoaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingRocketLoaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingRocketLoaderGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingRocketLoaderGetResponseEnvelopeErrors]
type zoneSettingRocketLoaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingRocketLoaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingRocketLoaderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingRocketLoaderGetResponseEnvelopeMessages]
type zoneSettingRocketLoaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
