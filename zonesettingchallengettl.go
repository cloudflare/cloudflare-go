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
func (r *ZoneSettingChallengeTTLService) Edit(ctx context.Context, params ZoneSettingChallengeTTLEditParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTTLEditResponse, err error) {
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
func (r *ZoneSettingChallengeTTLService) Get(ctx context.Context, query ZoneSettingChallengeTTLGetParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTTLGetResponse, err error) {
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
type ZoneSettingChallengeTTLEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingChallengeTTLEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTTLEditResponseJSON `json:"-"`
}

// zoneSettingChallengeTTLEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLEditResponse]
type zoneSettingChallengeTTLEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingChallengeTTLEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingChallengeTTLEditResponseID string

const (
	ZoneSettingChallengeTTLEditResponseIDChallengeTTL ZoneSettingChallengeTTLEditResponseID = "challenge_ttl"
)

// Current value of the zone setting.
type ZoneSettingChallengeTTLEditResponseValue float64

const (
	ZoneSettingChallengeTTLEditResponseValue300      ZoneSettingChallengeTTLEditResponseValue = 300
	ZoneSettingChallengeTTLEditResponseValue900      ZoneSettingChallengeTTLEditResponseValue = 900
	ZoneSettingChallengeTTLEditResponseValue1800     ZoneSettingChallengeTTLEditResponseValue = 1800
	ZoneSettingChallengeTTLEditResponseValue2700     ZoneSettingChallengeTTLEditResponseValue = 2700
	ZoneSettingChallengeTTLEditResponseValue3600     ZoneSettingChallengeTTLEditResponseValue = 3600
	ZoneSettingChallengeTTLEditResponseValue7200     ZoneSettingChallengeTTLEditResponseValue = 7200
	ZoneSettingChallengeTTLEditResponseValue10800    ZoneSettingChallengeTTLEditResponseValue = 10800
	ZoneSettingChallengeTTLEditResponseValue14400    ZoneSettingChallengeTTLEditResponseValue = 14400
	ZoneSettingChallengeTTLEditResponseValue28800    ZoneSettingChallengeTTLEditResponseValue = 28800
	ZoneSettingChallengeTTLEditResponseValue57600    ZoneSettingChallengeTTLEditResponseValue = 57600
	ZoneSettingChallengeTTLEditResponseValue86400    ZoneSettingChallengeTTLEditResponseValue = 86400
	ZoneSettingChallengeTTLEditResponseValue604800   ZoneSettingChallengeTTLEditResponseValue = 604800
	ZoneSettingChallengeTTLEditResponseValue2592000  ZoneSettingChallengeTTLEditResponseValue = 2592000
	ZoneSettingChallengeTTLEditResponseValue31536000 ZoneSettingChallengeTTLEditResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLEditResponseEditable bool

const (
	ZoneSettingChallengeTTLEditResponseEditableTrue  ZoneSettingChallengeTTLEditResponseEditable = true
	ZoneSettingChallengeTTLEditResponseEditableFalse ZoneSettingChallengeTTLEditResponseEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingChallengeTTLGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingChallengeTTLGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTTLGetResponseJSON `json:"-"`
}

// zoneSettingChallengeTTLGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingChallengeTTLGetResponse]
type zoneSettingChallengeTTLGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingChallengeTTLGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingChallengeTTLGetResponseID string

const (
	ZoneSettingChallengeTTLGetResponseIDChallengeTTL ZoneSettingChallengeTTLGetResponseID = "challenge_ttl"
)

// Current value of the zone setting.
type ZoneSettingChallengeTTLGetResponseValue float64

const (
	ZoneSettingChallengeTTLGetResponseValue300      ZoneSettingChallengeTTLGetResponseValue = 300
	ZoneSettingChallengeTTLGetResponseValue900      ZoneSettingChallengeTTLGetResponseValue = 900
	ZoneSettingChallengeTTLGetResponseValue1800     ZoneSettingChallengeTTLGetResponseValue = 1800
	ZoneSettingChallengeTTLGetResponseValue2700     ZoneSettingChallengeTTLGetResponseValue = 2700
	ZoneSettingChallengeTTLGetResponseValue3600     ZoneSettingChallengeTTLGetResponseValue = 3600
	ZoneSettingChallengeTTLGetResponseValue7200     ZoneSettingChallengeTTLGetResponseValue = 7200
	ZoneSettingChallengeTTLGetResponseValue10800    ZoneSettingChallengeTTLGetResponseValue = 10800
	ZoneSettingChallengeTTLGetResponseValue14400    ZoneSettingChallengeTTLGetResponseValue = 14400
	ZoneSettingChallengeTTLGetResponseValue28800    ZoneSettingChallengeTTLGetResponseValue = 28800
	ZoneSettingChallengeTTLGetResponseValue57600    ZoneSettingChallengeTTLGetResponseValue = 57600
	ZoneSettingChallengeTTLGetResponseValue86400    ZoneSettingChallengeTTLGetResponseValue = 86400
	ZoneSettingChallengeTTLGetResponseValue604800   ZoneSettingChallengeTTLGetResponseValue = 604800
	ZoneSettingChallengeTTLGetResponseValue2592000  ZoneSettingChallengeTTLGetResponseValue = 2592000
	ZoneSettingChallengeTTLGetResponseValue31536000 ZoneSettingChallengeTTLGetResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLGetResponseEditable bool

const (
	ZoneSettingChallengeTTLGetResponseEditableTrue  ZoneSettingChallengeTTLGetResponseEditable = true
	ZoneSettingChallengeTTLGetResponseEditableFalse ZoneSettingChallengeTTLGetResponseEditable = false
)

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
	Result ZoneSettingChallengeTTLEditResponse             `json:"result"`
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

func (r zoneSettingChallengeTTLEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingChallengeTTLEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingChallengeTTLEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingChallengeTTLGetResponse             `json:"result"`
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

func (r zoneSettingChallengeTTLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingChallengeTTLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingChallengeTTLGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
