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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
func (r *SettingWAFService) Edit(ctx context.Context, params SettingWAFEditParams, opts ...option.RequestOption) (res *WAF, err error) {
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
func (r *SettingWAFService) Get(ctx context.Context, query SettingWAFGetParams, opts ...option.RequestOption) (res *WAF, err error) {
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
type WAF struct {
	// ID of the zone setting.
	ID WAFID `json:"id,required"`
	// Current value of the zone setting.
	Value WAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       wafJSON   `json:"-"`
}

// wafJSON contains the JSON metadata for the struct [WAF]
type wafJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type WAFID string

const (
	WAFIDWAF WAFID = "waf"
)

func (r WAFID) IsKnown() bool {
	switch r {
	case WAFIDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type WAFValue string

const (
	WAFValueOn  WAFValue = "on"
	WAFValueOff WAFValue = "off"
)

func (r WAFValue) IsKnown() bool {
	switch r {
	case WAFValueOn, WAFValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type WAFEditable bool

const (
	WAFEditableTrue  WAFEditable = true
	WAFEditableFalse WAFEditable = false
)

func (r WAFEditable) IsKnown() bool {
	switch r {
	case WAFEditableTrue, WAFEditableFalse:
		return true
	}
	return false
}

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
	Result WAF                                `json:"result"`
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
	Result WAF                               `json:"result"`
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
