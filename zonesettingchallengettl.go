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

// ZoneSettingChallengeTTLService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingChallengeTTLService] method instead.
type ZoneSettingChallengeTTLService struct {
	Options []option.RequestOption
}

// NewZoneSettingChallengeTTLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingChallengeTTLService(opts ...option.RequestOption) (r *ZoneSettingChallengeTTLService) {
	r = &ZoneSettingChallengeTTLService{}
	r.Options = opts
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZoneSettingChallengeTTLService) Edit(ctx context.Context, params ZoneSettingChallengeTTLEditParams, opts ...option.RequestOption) (res *ZonesChallengeTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingChallengeTTLEditResponseEnvelope
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
func (r *ZoneSettingChallengeTTLService) Get(ctx context.Context, query ZoneSettingChallengeTTLGetParams, opts ...option.RequestOption) (res *ZonesChallengeTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingChallengeTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r ZonesChallengeTTL) implementsZoneSettingEditResponse() {}

func (r ZonesChallengeTTL) implementsZoneSettingGetResponse() {}

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

func (r ZonesChallengeTTLParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingChallengeTTLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingChallengeTTLEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingChallengeTTLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingChallengeTTLEditParamsValue float64

const (
	ZoneSettingChallengeTTLEditParamsValue300      ZoneSettingChallengeTTLEditParamsValue = 300
	ZoneSettingChallengeTTLEditParamsValue900      ZoneSettingChallengeTTLEditParamsValue = 900
	ZoneSettingChallengeTTLEditParamsValue1800     ZoneSettingChallengeTTLEditParamsValue = 1800
	ZoneSettingChallengeTTLEditParamsValue2700     ZoneSettingChallengeTTLEditParamsValue = 2700
	ZoneSettingChallengeTTLEditParamsValue3600     ZoneSettingChallengeTTLEditParamsValue = 3600
	ZoneSettingChallengeTTLEditParamsValue7200     ZoneSettingChallengeTTLEditParamsValue = 7200
	ZoneSettingChallengeTTLEditParamsValue10800    ZoneSettingChallengeTTLEditParamsValue = 10800
	ZoneSettingChallengeTTLEditParamsValue14400    ZoneSettingChallengeTTLEditParamsValue = 14400
	ZoneSettingChallengeTTLEditParamsValue28800    ZoneSettingChallengeTTLEditParamsValue = 28800
	ZoneSettingChallengeTTLEditParamsValue57600    ZoneSettingChallengeTTLEditParamsValue = 57600
	ZoneSettingChallengeTTLEditParamsValue86400    ZoneSettingChallengeTTLEditParamsValue = 86400
	ZoneSettingChallengeTTLEditParamsValue604800   ZoneSettingChallengeTTLEditParamsValue = 604800
	ZoneSettingChallengeTTLEditParamsValue2592000  ZoneSettingChallengeTTLEditParamsValue = 2592000
	ZoneSettingChallengeTTLEditParamsValue31536000 ZoneSettingChallengeTTLEditParamsValue = 31536000
)

type ZoneSettingChallengeTTLEditResponseEnvelope struct {
	Errors   []ZoneSettingChallengeTTLEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingChallengeTTLEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZonesChallengeTTL                               `json:"result"`
	JSON   zoneSettingChallengeTTLEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingChallengeTTLEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLEditResponseEnvelope]
type zoneSettingChallengeTTLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingChallengeTTLEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingChallengeTTLEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingChallengeTTLEditResponseEnvelopeErrors]
type zoneSettingChallengeTTLEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingChallengeTTLEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingChallengeTTLEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingChallengeTTLEditResponseEnvelopeMessages]
type zoneSettingChallengeTTLEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingChallengeTTLGetResponseEnvelope struct {
	Errors   []ZoneSettingChallengeTTLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingChallengeTTLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZonesChallengeTTL                              `json:"result"`
	JSON   zoneSettingChallengeTTLGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingChallengeTTLGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLGetResponseEnvelope]
type zoneSettingChallengeTTLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingChallengeTTLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingChallengeTTLGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingChallengeTTLGetResponseEnvelopeErrors]
type zoneSettingChallengeTTLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingChallengeTTLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingChallengeTTLGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingChallengeTTLGetResponseEnvelopeMessages]
type zoneSettingChallengeTTLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
