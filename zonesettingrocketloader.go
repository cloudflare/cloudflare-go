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
func (r *ZoneSettingRocketLoaderService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingRocketLoaderUpdateParams, opts ...option.RequestOption) (res *ZoneSettingRocketLoaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
func (r *ZoneSettingRocketLoaderService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingRocketLoaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingRocketLoaderUpdateResponse struct {
	Errors   []ZoneSettingRocketLoaderUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingRocketLoaderUpdateResponseMessage `json:"messages"`
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
	Result ZoneSettingRocketLoaderUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                      `json:"success"`
	JSON    zoneSettingRocketLoaderUpdateResponseJSON `json:"-"`
}

// zoneSettingRocketLoaderUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingRocketLoaderUpdateResponse]
type zoneSettingRocketLoaderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingRocketLoaderUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingRocketLoaderUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingRocketLoaderUpdateResponseError]
type zoneSettingRocketLoaderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingRocketLoaderUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingRocketLoaderUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingRocketLoaderUpdateResponseMessage]
type zoneSettingRocketLoaderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type ZoneSettingRocketLoaderUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingRocketLoaderUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingRocketLoaderUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingRocketLoaderUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingRocketLoaderUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingRocketLoaderUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingRocketLoaderUpdateResponseResult]
type zoneSettingRocketLoaderUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingRocketLoaderUpdateResponseResultID string

const (
	ZoneSettingRocketLoaderUpdateResponseResultIDRocketLoader ZoneSettingRocketLoaderUpdateResponseResultID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingRocketLoaderUpdateResponseResultEditable bool

const (
	ZoneSettingRocketLoaderUpdateResponseResultEditableTrue  ZoneSettingRocketLoaderUpdateResponseResultEditable = true
	ZoneSettingRocketLoaderUpdateResponseResultEditableFalse ZoneSettingRocketLoaderUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingRocketLoaderUpdateResponseResultValue string

const (
	ZoneSettingRocketLoaderUpdateResponseResultValueOn  ZoneSettingRocketLoaderUpdateResponseResultValue = "on"
	ZoneSettingRocketLoaderUpdateResponseResultValueOff ZoneSettingRocketLoaderUpdateResponseResultValue = "off"
)

type ZoneSettingRocketLoaderListResponse struct {
	Errors   []ZoneSettingRocketLoaderListResponseError   `json:"errors"`
	Messages []ZoneSettingRocketLoaderListResponseMessage `json:"messages"`
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
	Result ZoneSettingRocketLoaderListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    zoneSettingRocketLoaderListResponseJSON `json:"-"`
}

// zoneSettingRocketLoaderListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingRocketLoaderListResponse]
type zoneSettingRocketLoaderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingRocketLoaderListResponseErrorJSON `json:"-"`
}

// zoneSettingRocketLoaderListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingRocketLoaderListResponseError]
type zoneSettingRocketLoaderListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingRocketLoaderListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingRocketLoaderListResponseMessageJSON `json:"-"`
}

// zoneSettingRocketLoaderListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingRocketLoaderListResponseMessage]
type zoneSettingRocketLoaderListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type ZoneSettingRocketLoaderListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingRocketLoaderListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingRocketLoaderListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingRocketLoaderListResponseResultValue `json:"value"`
	JSON  zoneSettingRocketLoaderListResponseResultJSON  `json:"-"`
}

// zoneSettingRocketLoaderListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingRocketLoaderListResponseResult]
type zoneSettingRocketLoaderListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingRocketLoaderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingRocketLoaderListResponseResultID string

const (
	ZoneSettingRocketLoaderListResponseResultIDRocketLoader ZoneSettingRocketLoaderListResponseResultID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingRocketLoaderListResponseResultEditable bool

const (
	ZoneSettingRocketLoaderListResponseResultEditableTrue  ZoneSettingRocketLoaderListResponseResultEditable = true
	ZoneSettingRocketLoaderListResponseResultEditableFalse ZoneSettingRocketLoaderListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingRocketLoaderListResponseResultValue string

const (
	ZoneSettingRocketLoaderListResponseResultValueOn  ZoneSettingRocketLoaderListResponseResultValue = "on"
	ZoneSettingRocketLoaderListResponseResultValueOff ZoneSettingRocketLoaderListResponseResultValue = "off"
)

type ZoneSettingRocketLoaderUpdateParams struct {
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
	Value param.Field[ZoneSettingRocketLoaderUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingRocketLoaderUpdateParams) MarshalJSON() (data []byte, err error) {
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
type ZoneSettingRocketLoaderUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingRocketLoaderUpdateParamsValueID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingRocketLoaderUpdateParamsValueValue] `json:"value"`
}

func (r ZoneSettingRocketLoaderUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingRocketLoaderUpdateParamsValueID string

const (
	ZoneSettingRocketLoaderUpdateParamsValueIDRocketLoader ZoneSettingRocketLoaderUpdateParamsValueID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingRocketLoaderUpdateParamsValueEditable bool

const (
	ZoneSettingRocketLoaderUpdateParamsValueEditableTrue  ZoneSettingRocketLoaderUpdateParamsValueEditable = true
	ZoneSettingRocketLoaderUpdateParamsValueEditableFalse ZoneSettingRocketLoaderUpdateParamsValueEditable = false
)

// Value of the zone setting.
type ZoneSettingRocketLoaderUpdateParamsValueValue string

const (
	ZoneSettingRocketLoaderUpdateParamsValueValueOn  ZoneSettingRocketLoaderUpdateParamsValueValue = "on"
	ZoneSettingRocketLoaderUpdateParamsValueValueOff ZoneSettingRocketLoaderUpdateParamsValueValue = "off"
)
