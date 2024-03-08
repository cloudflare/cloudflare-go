// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZoneSettingWAFService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingWAFService] method
// instead.
type ZoneSettingWAFService struct {
	Options []option.RequestOption
}

// NewZoneSettingWAFService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingWAFService(opts ...option.RequestOption) (r *ZoneSettingWAFService) {
	r = &ZoneSettingWAFService{}
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
func (r *ZoneSettingWAFService) Edit(ctx context.Context, params ZoneSettingWAFEditParams, opts ...option.RequestOption) (res *ZoneSettingWAFEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingWAFEditResponseEnvelope
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
func (r *ZoneSettingWAFService) Get(ctx context.Context, query ZoneSettingWAFGetParams, opts ...option.RequestOption) (res *ZoneSettingWAFGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingWAFGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/waf", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type ZoneSettingWAFEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingWAFEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingWAFEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWAFEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWAFEditResponseJSON `json:"-"`
}

// zoneSettingWAFEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWAFEditResponse]
type zoneSettingWAFEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingWAFEditResponseID string

const (
	ZoneSettingWAFEditResponseIDWAF ZoneSettingWAFEditResponseID = "waf"
)

// Current value of the zone setting.
type ZoneSettingWAFEditResponseValue string

const (
	ZoneSettingWAFEditResponseValueOn  ZoneSettingWAFEditResponseValue = "on"
	ZoneSettingWAFEditResponseValueOff ZoneSettingWAFEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWAFEditResponseEditable bool

const (
	ZoneSettingWAFEditResponseEditableTrue  ZoneSettingWAFEditResponseEditable = true
	ZoneSettingWAFEditResponseEditableFalse ZoneSettingWAFEditResponseEditable = false
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
type ZoneSettingWAFGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingWAFGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingWAFGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWAFGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWAFGetResponseJSON `json:"-"`
}

// zoneSettingWAFGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWAFGetResponse]
type zoneSettingWAFGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingWAFGetResponseID string

const (
	ZoneSettingWAFGetResponseIDWAF ZoneSettingWAFGetResponseID = "waf"
)

// Current value of the zone setting.
type ZoneSettingWAFGetResponseValue string

const (
	ZoneSettingWAFGetResponseValueOn  ZoneSettingWAFGetResponseValue = "on"
	ZoneSettingWAFGetResponseValueOff ZoneSettingWAFGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWAFGetResponseEditable bool

const (
	ZoneSettingWAFGetResponseEditableTrue  ZoneSettingWAFGetResponseEditable = true
	ZoneSettingWAFGetResponseEditableFalse ZoneSettingWAFGetResponseEditable = false
)

type ZoneSettingWAFEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingWAFEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingWAFEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWAFEditParamsValue string

const (
	ZoneSettingWAFEditParamsValueOn  ZoneSettingWAFEditParamsValue = "on"
	ZoneSettingWAFEditParamsValueOff ZoneSettingWAFEditParamsValue = "off"
)

type ZoneSettingWAFEditResponseEnvelope struct {
	Errors   []ZoneSettingWAFEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingWAFEditResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZoneSettingWAFEditResponse             `json:"result"`
	JSON   zoneSettingWAFEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingWAFEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingWAFEditResponseEnvelope]
type zoneSettingWAFEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWAFEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingWAFEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingWAFEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingWAFEditResponseEnvelopeErrors]
type zoneSettingWAFEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWAFEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingWAFEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingWAFEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingWAFEditResponseEnvelopeMessages]
type zoneSettingWAFEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWAFGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingWAFGetResponseEnvelope struct {
	Errors   []ZoneSettingWAFGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingWAFGetResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZoneSettingWAFGetResponse             `json:"result"`
	JSON   zoneSettingWAFGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingWAFGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingWAFGetResponseEnvelope]
type zoneSettingWAFGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWAFGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingWAFGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingWAFGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingWAFGetResponseEnvelopeErrors]
type zoneSettingWAFGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWAFGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingWAFGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingWAFGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneSettingWAFGetResponseEnvelopeMessages]
type zoneSettingWAFGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWAFGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
