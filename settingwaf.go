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
func (r *SettingWAFService) Update(ctx context.Context, zoneID string, body SettingWAFUpdateParams, opts ...option.RequestOption) (res *SettingWAFUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWAFUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/waf", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingWAFService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingWAFGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWAFGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/waf", zoneID)
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
type SettingWAFUpdateResponse struct {
	// ID of the zone setting.
	ID SettingWAFUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingWAFUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingWAFUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingWAFUpdateResponseJSON `json:"-"`
}

// settingWAFUpdateResponseJSON contains the JSON metadata for the struct
// [SettingWAFUpdateResponse]
type settingWAFUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingWAFUpdateResponseID string

const (
	SettingWAFUpdateResponseIDWAF SettingWAFUpdateResponseID = "waf"
)

// Current value of the zone setting.
type SettingWAFUpdateResponseValue string

const (
	SettingWAFUpdateResponseValueOn  SettingWAFUpdateResponseValue = "on"
	SettingWAFUpdateResponseValueOff SettingWAFUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingWAFUpdateResponseEditable bool

const (
	SettingWAFUpdateResponseEditableTrue  SettingWAFUpdateResponseEditable = true
	SettingWAFUpdateResponseEditableFalse SettingWAFUpdateResponseEditable = false
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
type SettingWAFGetResponse struct {
	// ID of the zone setting.
	ID SettingWAFGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingWAFGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingWAFGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingWAFGetResponseJSON `json:"-"`
}

// settingWAFGetResponseJSON contains the JSON metadata for the struct
// [SettingWAFGetResponse]
type settingWAFGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingWAFGetResponseID string

const (
	SettingWAFGetResponseIDWAF SettingWAFGetResponseID = "waf"
)

// Current value of the zone setting.
type SettingWAFGetResponseValue string

const (
	SettingWAFGetResponseValueOn  SettingWAFGetResponseValue = "on"
	SettingWAFGetResponseValueOff SettingWAFGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingWAFGetResponseEditable bool

const (
	SettingWAFGetResponseEditableTrue  SettingWAFGetResponseEditable = true
	SettingWAFGetResponseEditableFalse SettingWAFGetResponseEditable = false
)

type SettingWAFUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingWAFUpdateParamsValue] `json:"value,required"`
}

func (r SettingWAFUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingWAFUpdateParamsValue string

const (
	SettingWAFUpdateParamsValueOn  SettingWAFUpdateParamsValue = "on"
	SettingWAFUpdateParamsValueOff SettingWAFUpdateParamsValue = "off"
)

type SettingWAFUpdateResponseEnvelope struct {
	Errors   []SettingWAFUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWAFUpdateResponseEnvelopeMessages `json:"messages,required"`
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
	Result SettingWAFUpdateResponse             `json:"result"`
	JSON   settingWAFUpdateResponseEnvelopeJSON `json:"-"`
}

// settingWAFUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWAFUpdateResponseEnvelope]
type settingWAFUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWAFUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingWAFUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWAFUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWAFUpdateResponseEnvelopeErrors]
type settingWAFUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWAFUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingWAFUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWAFUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWAFUpdateResponseEnvelopeMessages]
type settingWAFUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWAFUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Result SettingWAFGetResponse             `json:"result"`
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
