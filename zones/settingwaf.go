// File generated from our OpenAPI spec by Stainless.

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
func (r *SettingWAFService) Edit(ctx context.Context, params SettingWAFEditParams, opts ...option.RequestOption) (res *ZonesWAF, err error) {
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
func (r *SettingWAFService) Get(ctx context.Context, query SettingWAFGetParams, opts ...option.RequestOption) (res *ZonesWAF, err error) {
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
type ZonesWAF struct {
	// ID of the zone setting.
	ID ZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time    `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesWAFJSON `json:"-"`
}

// zonesWAFJSON contains the JSON metadata for the struct [ZonesWAF]
type zonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesWAFJSON) RawJSON() string {
	return r.raw
}

func (r ZonesWAF) implementsZonesSettingEditResponse() {}

func (r ZonesWAF) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesWAFID string

const (
	ZonesWAFIDWAF ZonesWAFID = "waf"
)

// Current value of the zone setting.
type ZonesWAFValue string

const (
	ZonesWAFValueOn  ZonesWAFValue = "on"
	ZonesWAFValueOff ZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesWAFEditable bool

const (
	ZonesWAFEditableTrue  ZonesWAFEditable = true
	ZonesWAFEditableFalse ZonesWAFEditable = false
)

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
type ZonesWAFParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWAFID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesWAFValue] `json:"value,required"`
}

func (r ZonesWAFParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWAFParam) implementsZonesSettingEditParamsItem() {}

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

type SettingWAFEditResponseEnvelope struct {
	Errors   []SettingWAFEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWAFEditResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZonesWAF                           `json:"result"`
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

type SettingWAFEditResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingWAFEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWAFEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWAFEditResponseEnvelopeErrors]
type settingWAFEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWAFEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWAFEditResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingWAFEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWAFEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWAFEditResponseEnvelopeMessages]
type settingWAFEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWAFEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingWAFGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWAFGetResponseEnvelope struct {
	Errors   []SettingWAFGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWAFGetResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZonesWAF                          `json:"result"`
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

type SettingWAFGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingWAFGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWAFGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWAFGetResponseEnvelopeErrors]
type settingWAFGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWAFGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWAFGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingWAFGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWAFGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWAFGetResponseEnvelopeMessages]
type settingWAFGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWAFGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
