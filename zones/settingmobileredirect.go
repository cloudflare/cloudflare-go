// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingMobileRedirectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMobileRedirectService]
// method instead.
type SettingMobileRedirectService struct {
	Options []option.RequestOption
}

// NewSettingMobileRedirectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingMobileRedirectService(opts ...option.RequestOption) (r *SettingMobileRedirectService) {
	r = &SettingMobileRedirectService{}
	r.Options = opts
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *SettingMobileRedirectService) Edit(ctx context.Context, params SettingMobileRedirectEditParams, opts ...option.RequestOption) (res *ZoneSettingMobileRedirect, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectEditResponseEnvelope
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
func (r *SettingMobileRedirectService) Get(ctx context.Context, query SettingMobileRedirectGetParams, opts ...option.RequestOption) (res *ZoneSettingMobileRedirect, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type ZoneSettingMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMobileRedirectJSON `json:"-"`
}

// zoneSettingMobileRedirectJSON contains the JSON metadata for the struct
// [ZoneSettingMobileRedirect]
type zoneSettingMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMobileRedirectJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingMobileRedirect) implementsZonesSettingEditResponse() {}

func (r ZoneSettingMobileRedirect) implementsZonesSettingGetResponse() {}

// Identifier of the zone setting.
type ZoneSettingMobileRedirectID string

const (
	ZoneSettingMobileRedirectIDMobileRedirect ZoneSettingMobileRedirectID = "mobile_redirect"
)

func (r ZoneSettingMobileRedirectID) IsKnown() bool {
	switch r {
	case ZoneSettingMobileRedirectIDMobileRedirect:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI bool                               `json:"strip_uri"`
	JSON     zoneSettingMobileRedirectValueJSON `json:"-"`
}

// zoneSettingMobileRedirectValueJSON contains the JSON metadata for the struct
// [ZoneSettingMobileRedirectValue]
type zoneSettingMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMobileRedirectValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectValueStatus string

const (
	ZoneSettingMobileRedirectValueStatusOn  ZoneSettingMobileRedirectValueStatus = "on"
	ZoneSettingMobileRedirectValueStatusOff ZoneSettingMobileRedirectValueStatus = "off"
)

func (r ZoneSettingMobileRedirectValueStatus) IsKnown() bool {
	switch r {
	case ZoneSettingMobileRedirectValueStatusOn, ZoneSettingMobileRedirectValueStatusOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMobileRedirectEditable bool

const (
	ZoneSettingMobileRedirectEditableTrue  ZoneSettingMobileRedirectEditable = true
	ZoneSettingMobileRedirectEditableFalse ZoneSettingMobileRedirectEditable = false
)

func (r ZoneSettingMobileRedirectEditable) IsKnown() bool {
	switch r {
	case ZoneSettingMobileRedirectEditableTrue, ZoneSettingMobileRedirectEditableFalse:
		return true
	}
	return false
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingMobileRedirectParam struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingMobileRedirectID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingMobileRedirectValueParam] `json:"value,required"`
}

func (r ZoneSettingMobileRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingMobileRedirectParam) implementsZonesSettingEditParamsItemUnion() {}

// Current value of the zone setting.
type ZoneSettingMobileRedirectValueParam struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[ZoneSettingMobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI param.Field[bool] `json:"strip_uri"`
}

func (r ZoneSettingMobileRedirectValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingMobileRedirectEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMobileRedirectEditParamsValue] `json:"value,required"`
}

func (r SettingMobileRedirectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMobileRedirectEditParamsValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[SettingMobileRedirectEditParamsValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI param.Field[bool] `json:"strip_uri"`
}

func (r SettingMobileRedirectEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectEditParamsValueStatus string

const (
	SettingMobileRedirectEditParamsValueStatusOn  SettingMobileRedirectEditParamsValueStatus = "on"
	SettingMobileRedirectEditParamsValueStatusOff SettingMobileRedirectEditParamsValueStatus = "off"
)

func (r SettingMobileRedirectEditParamsValueStatus) IsKnown() bool {
	switch r {
	case SettingMobileRedirectEditParamsValueStatusOn, SettingMobileRedirectEditParamsValueStatusOff:
		return true
	}
	return false
}

type SettingMobileRedirectEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result ZoneSettingMobileRedirect                     `json:"result"`
	JSON   settingMobileRedirectEditResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMobileRedirectEditResponseEnvelope]
type settingMobileRedirectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMobileRedirectGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMobileRedirectGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result ZoneSettingMobileRedirect                    `json:"result"`
	JSON   settingMobileRedirectGetResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMobileRedirectGetResponseEnvelope]
type settingMobileRedirectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
