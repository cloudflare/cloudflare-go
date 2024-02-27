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

// ZoneSettingMobileRedirectService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingMobileRedirectService] method instead.
type ZoneSettingMobileRedirectService struct {
	Options []option.RequestOption
}

// NewZoneSettingMobileRedirectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingMobileRedirectService(opts ...option.RequestOption) (r *ZoneSettingMobileRedirectService) {
	r = &ZoneSettingMobileRedirectService{}
	r.Options = opts
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *ZoneSettingMobileRedirectService) Edit(ctx context.Context, params ZoneSettingMobileRedirectEditParams, opts ...option.RequestOption) (res *ZoneSettingMobileRedirectEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMobileRedirectEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *ZoneSettingMobileRedirectService) Get(ctx context.Context, query ZoneSettingMobileRedirectGetParams, opts ...option.RequestOption) (res *ZoneSettingMobileRedirectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMobileRedirectGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingMobileRedirectEditResponse struct {
	// Identifier of the zone setting.
	ID ZoneSettingMobileRedirectEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMobileRedirectEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMobileRedirectEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMobileRedirectEditResponseJSON `json:"-"`
}

// zoneSettingMobileRedirectEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMobileRedirectEditResponse]
type zoneSettingMobileRedirectEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type ZoneSettingMobileRedirectEditResponseID string

const (
	ZoneSettingMobileRedirectEditResponseIDMobileRedirect ZoneSettingMobileRedirectEditResponseID = "mobile_redirect"
)

// Current value of the zone setting.
type ZoneSettingMobileRedirectEditResponseValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingMobileRedirectEditResponseValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                           `json:"strip_uri"`
	JSON     zoneSettingMobileRedirectEditResponseValueJSON `json:"-"`
}

// zoneSettingMobileRedirectEditResponseValueJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectEditResponseValue]
type zoneSettingMobileRedirectEditResponseValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectEditResponseValueStatus string

const (
	ZoneSettingMobileRedirectEditResponseValueStatusOn  ZoneSettingMobileRedirectEditResponseValueStatus = "on"
	ZoneSettingMobileRedirectEditResponseValueStatusOff ZoneSettingMobileRedirectEditResponseValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMobileRedirectEditResponseEditable bool

const (
	ZoneSettingMobileRedirectEditResponseEditableTrue  ZoneSettingMobileRedirectEditResponseEditable = true
	ZoneSettingMobileRedirectEditResponseEditableFalse ZoneSettingMobileRedirectEditResponseEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingMobileRedirectGetResponse struct {
	// Identifier of the zone setting.
	ID ZoneSettingMobileRedirectGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMobileRedirectGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMobileRedirectGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMobileRedirectGetResponseJSON `json:"-"`
}

// zoneSettingMobileRedirectGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMobileRedirectGetResponse]
type zoneSettingMobileRedirectGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type ZoneSettingMobileRedirectGetResponseID string

const (
	ZoneSettingMobileRedirectGetResponseIDMobileRedirect ZoneSettingMobileRedirectGetResponseID = "mobile_redirect"
)

// Current value of the zone setting.
type ZoneSettingMobileRedirectGetResponseValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingMobileRedirectGetResponseValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                          `json:"strip_uri"`
	JSON     zoneSettingMobileRedirectGetResponseValueJSON `json:"-"`
}

// zoneSettingMobileRedirectGetResponseValueJSON contains the JSON metadata for the
// struct [ZoneSettingMobileRedirectGetResponseValue]
type zoneSettingMobileRedirectGetResponseValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectGetResponseValueStatus string

const (
	ZoneSettingMobileRedirectGetResponseValueStatusOn  ZoneSettingMobileRedirectGetResponseValueStatus = "on"
	ZoneSettingMobileRedirectGetResponseValueStatusOff ZoneSettingMobileRedirectGetResponseValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMobileRedirectGetResponseEditable bool

const (
	ZoneSettingMobileRedirectGetResponseEditableTrue  ZoneSettingMobileRedirectGetResponseEditable = true
	ZoneSettingMobileRedirectGetResponseEditableFalse ZoneSettingMobileRedirectGetResponseEditable = false
)

type ZoneSettingMobileRedirectEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingMobileRedirectEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingMobileRedirectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMobileRedirectEditParamsValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[ZoneSettingMobileRedirectEditParamsValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r ZoneSettingMobileRedirectEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectEditParamsValueStatus string

const (
	ZoneSettingMobileRedirectEditParamsValueStatusOn  ZoneSettingMobileRedirectEditParamsValueStatus = "on"
	ZoneSettingMobileRedirectEditParamsValueStatusOff ZoneSettingMobileRedirectEditParamsValueStatus = "off"
)

type ZoneSettingMobileRedirectEditResponseEnvelope struct {
	Errors   []ZoneSettingMobileRedirectEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMobileRedirectEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result ZoneSettingMobileRedirectEditResponse             `json:"result"`
	JSON   zoneSettingMobileRedirectEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMobileRedirectEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectEditResponseEnvelope]
type zoneSettingMobileRedirectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingMobileRedirectEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMobileRedirectEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingMobileRedirectEditResponseEnvelopeErrors]
type zoneSettingMobileRedirectEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingMobileRedirectEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMobileRedirectEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingMobileRedirectEditResponseEnvelopeMessages]
type zoneSettingMobileRedirectEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingMobileRedirectGetResponseEnvelope struct {
	Errors   []ZoneSettingMobileRedirectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMobileRedirectGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result ZoneSettingMobileRedirectGetResponse             `json:"result"`
	JSON   zoneSettingMobileRedirectGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMobileRedirectGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectGetResponseEnvelope]
type zoneSettingMobileRedirectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingMobileRedirectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMobileRedirectGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingMobileRedirectGetResponseEnvelopeErrors]
type zoneSettingMobileRedirectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingMobileRedirectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMobileRedirectGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingMobileRedirectGetResponseEnvelopeMessages]
type zoneSettingMobileRedirectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
