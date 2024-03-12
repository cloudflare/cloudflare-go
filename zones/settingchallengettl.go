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

// SettingChallengeTTLService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingChallengeTTLService]
// method instead.
type SettingChallengeTTLService struct {
	Options []option.RequestOption
}

// NewSettingChallengeTTLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingChallengeTTLService(opts ...option.RequestOption) (r *SettingChallengeTTLService) {
	r = &SettingChallengeTTLService{}
	r.Options = opts
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *SettingChallengeTTLService) Edit(ctx context.Context, params SettingChallengeTTLEditParams, opts ...option.RequestOption) (res *ZonesChallengeTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingChallengeTTLEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *SettingChallengeTTLService) Get(ctx context.Context, query SettingChallengeTTLGetParams, opts ...option.RequestOption) (res *ZonesChallengeTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingChallengeTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZonesChallengeTTL struct {
	// ID of the zone setting.
	ID ZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesChallengeTTLJSON `json:"-"`
}

// zonesChallengeTTLJSON contains the JSON metadata for the struct
// [ZonesChallengeTTL]
type zonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesChallengeTTLJSON) RawJSON() string {
	return r.raw
}

func (r ZonesChallengeTTL) implementsZonesSettingEditResponse() {}

func (r ZonesChallengeTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesChallengeTTLID string

const (
	ZonesChallengeTTLIDChallengeTTL ZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type ZonesChallengeTTLValue float64

const (
	ZonesChallengeTTLValue300      ZonesChallengeTTLValue = 300
	ZonesChallengeTTLValue900      ZonesChallengeTTLValue = 900
	ZonesChallengeTTLValue1800     ZonesChallengeTTLValue = 1800
	ZonesChallengeTTLValue2700     ZonesChallengeTTLValue = 2700
	ZonesChallengeTTLValue3600     ZonesChallengeTTLValue = 3600
	ZonesChallengeTTLValue7200     ZonesChallengeTTLValue = 7200
	ZonesChallengeTTLValue10800    ZonesChallengeTTLValue = 10800
	ZonesChallengeTTLValue14400    ZonesChallengeTTLValue = 14400
	ZonesChallengeTTLValue28800    ZonesChallengeTTLValue = 28800
	ZonesChallengeTTLValue57600    ZonesChallengeTTLValue = 57600
	ZonesChallengeTTLValue86400    ZonesChallengeTTLValue = 86400
	ZonesChallengeTTLValue604800   ZonesChallengeTTLValue = 604800
	ZonesChallengeTTLValue2592000  ZonesChallengeTTLValue = 2592000
	ZonesChallengeTTLValue31536000 ZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesChallengeTTLEditable bool

const (
	ZonesChallengeTTLEditableTrue  ZonesChallengeTTLEditable = true
	ZonesChallengeTTLEditableFalse ZonesChallengeTTLEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZonesChallengeTTLParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesChallengeTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesChallengeTTLValue] `json:"value,required"`
}

func (r ZonesChallengeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesChallengeTTLParam) implementsZonesSettingEditParamsItem() {}

type SettingChallengeTTLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingChallengeTTLEditParamsValue] `json:"value,required"`
}

func (r SettingChallengeTTLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingChallengeTTLEditParamsValue float64

const (
	SettingChallengeTTLEditParamsValue300      SettingChallengeTTLEditParamsValue = 300
	SettingChallengeTTLEditParamsValue900      SettingChallengeTTLEditParamsValue = 900
	SettingChallengeTTLEditParamsValue1800     SettingChallengeTTLEditParamsValue = 1800
	SettingChallengeTTLEditParamsValue2700     SettingChallengeTTLEditParamsValue = 2700
	SettingChallengeTTLEditParamsValue3600     SettingChallengeTTLEditParamsValue = 3600
	SettingChallengeTTLEditParamsValue7200     SettingChallengeTTLEditParamsValue = 7200
	SettingChallengeTTLEditParamsValue10800    SettingChallengeTTLEditParamsValue = 10800
	SettingChallengeTTLEditParamsValue14400    SettingChallengeTTLEditParamsValue = 14400
	SettingChallengeTTLEditParamsValue28800    SettingChallengeTTLEditParamsValue = 28800
	SettingChallengeTTLEditParamsValue57600    SettingChallengeTTLEditParamsValue = 57600
	SettingChallengeTTLEditParamsValue86400    SettingChallengeTTLEditParamsValue = 86400
	SettingChallengeTTLEditParamsValue604800   SettingChallengeTTLEditParamsValue = 604800
	SettingChallengeTTLEditParamsValue2592000  SettingChallengeTTLEditParamsValue = 2592000
	SettingChallengeTTLEditParamsValue31536000 SettingChallengeTTLEditParamsValue = 31536000
)

type SettingChallengeTTLEditResponseEnvelope struct {
	Errors   []SettingChallengeTTLEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingChallengeTTLEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZonesChallengeTTL                           `json:"result"`
	JSON   settingChallengeTTLEditResponseEnvelopeJSON `json:"-"`
}

// settingChallengeTTLEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingChallengeTTLEditResponseEnvelope]
type settingChallengeTTLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingChallengeTTLEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingChallengeTTLEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingChallengeTTLEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingChallengeTTLEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingChallengeTTLEditResponseEnvelopeErrors]
type settingChallengeTTLEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingChallengeTTLEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingChallengeTTLEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingChallengeTTLEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingChallengeTTLEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingChallengeTTLEditResponseEnvelopeMessages]
type settingChallengeTTLEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingChallengeTTLEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingChallengeTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingChallengeTTLGetResponseEnvelope struct {
	Errors   []SettingChallengeTTLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingChallengeTTLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZonesChallengeTTL                          `json:"result"`
	JSON   settingChallengeTTLGetResponseEnvelopeJSON `json:"-"`
}

// settingChallengeTTLGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingChallengeTTLGetResponseEnvelope]
type settingChallengeTTLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingChallengeTTLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingChallengeTTLGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingChallengeTTLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingChallengeTTLGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingChallengeTTLGetResponseEnvelopeErrors]
type settingChallengeTTLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingChallengeTTLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingChallengeTTLGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingChallengeTTLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingChallengeTTLGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingChallengeTTLGetResponseEnvelopeMessages]
type settingChallengeTTLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingChallengeTTLGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
