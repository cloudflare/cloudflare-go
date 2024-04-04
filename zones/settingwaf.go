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

// SettingWAFService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingWAFService] method instead.
type SettingWAFService struct {
	Options []option.RequestOption
}

// NewSettingWAFService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingWAFService(opts ...option.RequestOption) (r *SettingWAFService) {
	r = &SettingWAFService{}
	r.Options = opts
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
func (r *SettingWAFService) Edit(ctx context.Context, params SettingWAFEditParams, opts ...option.RequestOption) (res *ZoneSettingWAF, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWAFEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/waf", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
func (r *SettingWAFService) Get(ctx context.Context, query SettingWAFGetParams, opts ...option.RequestOption) (res *ZoneSettingWAF, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWAFGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/waf", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type ZoneSettingWAF struct {
	// ID of the zone setting.
	ID ZoneSettingWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWAFJSON `json:"-"`
}

// zoneSettingWAFJSON contains the JSON metadata for the struct [ZoneSettingWAF]
type zoneSettingWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingWAF) implementsZonesSettingEditResponse() {}

func (r ZoneSettingWAF) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingWAFID string

const (
	ZoneSettingWAFIDWAF ZoneSettingWAFID = "waf"
)

func (r ZoneSettingWAFID) IsKnown() bool {
	switch r {
	case ZoneSettingWAFIDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingWAFValue string

const (
	ZoneSettingWAFValueOn  ZoneSettingWAFValue = "on"
	ZoneSettingWAFValueOff ZoneSettingWAFValue = "off"
)

func (r ZoneSettingWAFValue) IsKnown() bool {
	switch r {
	case ZoneSettingWAFValueOn, ZoneSettingWAFValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWAFEditable bool

const (
	ZoneSettingWAFEditableTrue  ZoneSettingWAFEditable = true
	ZoneSettingWAFEditableFalse ZoneSettingWAFEditable = false
)

func (r ZoneSettingWAFEditable) IsKnown() bool {
	switch r {
	case ZoneSettingWAFEditableTrue, ZoneSettingWAFEditableFalse:
		return true
	}
	return false
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type ZoneSettingWAFParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingWAFID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingWAFValue] `json:"value,required"`
}

func (r ZoneSettingWAFParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingWAFParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingWAFEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingWAFEditParamsValue] `json:"value,required"`
}

func (r SettingWAFEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingWAFEditParamsValue string

const (
	SettingWAFEditParamsValueOn  SettingWAFEditParamsValue = "on"
	SettingWAFEditParamsValueOff SettingWAFEditParamsValue = "off"
)

func (r SettingWAFEditParamsValue) IsKnown() bool {
	switch r {
	case SettingWAFEditParamsValueOn, SettingWAFEditParamsValueOff:
		return true
	}
	return false
}

type SettingWAFEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The WAF examines HTTP requests to your website. It inspects both GET and POST
	// requests and applies rules to help filter out illegitimate traffic from
	// legitimate website visitors. The Cloudflare WAF inspects website addresses or
	// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
	// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
	// page that asks them to submit a CAPTCHA successfully to continue their action.
	// If the challenge is failed, the action will be stopped. What this means is that
	// Cloudflare's WAF will block any traffic identified as illegitimate before it
	// reaches your origin web server.
	// (https://support.cloudflare.com/hc/en-us/articles/200172016).
	Result ZoneSettingWAF                     `json:"result"`
	JSON   settingWAFEditResponseEnvelopeJSON `json:"-"`
}

// settingWAFEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWAFEditResponseEnvelope]
type settingWAFEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWAFEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingWAFGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWAFGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The WAF examines HTTP requests to your website. It inspects both GET and POST
	// requests and applies rules to help filter out illegitimate traffic from
	// legitimate website visitors. The Cloudflare WAF inspects website addresses or
	// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
	// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
	// page that asks them to submit a CAPTCHA successfully to continue their action.
	// If the challenge is failed, the action will be stopped. What this means is that
	// Cloudflare's WAF will block any traffic identified as illegitimate before it
	// reaches your origin web server.
	// (https://support.cloudflare.com/hc/en-us/articles/200172016).
	Result ZoneSettingWAF                    `json:"result"`
	JSON   settingWAFGetResponseEnvelopeJSON `json:"-"`
}

// settingWAFGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWAFGetResponseEnvelope]
type settingWAFGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWAFGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
