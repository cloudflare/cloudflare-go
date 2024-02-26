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
func (r *SettingChallengeTTLService) Edit(ctx context.Context, params SettingChallengeTTLEditParams, opts ...option.RequestOption) (res *SettingChallengeTTLEditResponse, err error) {
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
func (r *SettingChallengeTTLService) Get(ctx context.Context, query SettingChallengeTTLGetParams, opts ...option.RequestOption) (res *SettingChallengeTTLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingChallengeTTLGetResponseEnvelope
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
type SettingChallengeTTLEditResponse struct {
	// ID of the zone setting.
	ID SettingChallengeTTLEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingChallengeTTLEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingChallengeTTLEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingChallengeTTLEditResponseJSON `json:"-"`
}

// settingChallengeTTLEditResponseJSON contains the JSON metadata for the struct
// [SettingChallengeTTLEditResponse]
type settingChallengeTTLEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingChallengeTTLEditResponseID string

const (
	SettingChallengeTTLEditResponseIDChallengeTTL SettingChallengeTTLEditResponseID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingChallengeTTLEditResponseValue float64

const (
	SettingChallengeTTLEditResponseValue300      SettingChallengeTTLEditResponseValue = 300
	SettingChallengeTTLEditResponseValue900      SettingChallengeTTLEditResponseValue = 900
	SettingChallengeTTLEditResponseValue1800     SettingChallengeTTLEditResponseValue = 1800
	SettingChallengeTTLEditResponseValue2700     SettingChallengeTTLEditResponseValue = 2700
	SettingChallengeTTLEditResponseValue3600     SettingChallengeTTLEditResponseValue = 3600
	SettingChallengeTTLEditResponseValue7200     SettingChallengeTTLEditResponseValue = 7200
	SettingChallengeTTLEditResponseValue10800    SettingChallengeTTLEditResponseValue = 10800
	SettingChallengeTTLEditResponseValue14400    SettingChallengeTTLEditResponseValue = 14400
	SettingChallengeTTLEditResponseValue28800    SettingChallengeTTLEditResponseValue = 28800
	SettingChallengeTTLEditResponseValue57600    SettingChallengeTTLEditResponseValue = 57600
	SettingChallengeTTLEditResponseValue86400    SettingChallengeTTLEditResponseValue = 86400
	SettingChallengeTTLEditResponseValue604800   SettingChallengeTTLEditResponseValue = 604800
	SettingChallengeTTLEditResponseValue2592000  SettingChallengeTTLEditResponseValue = 2592000
	SettingChallengeTTLEditResponseValue31536000 SettingChallengeTTLEditResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingChallengeTTLEditResponseEditable bool

const (
	SettingChallengeTTLEditResponseEditableTrue  SettingChallengeTTLEditResponseEditable = true
	SettingChallengeTTLEditResponseEditableFalse SettingChallengeTTLEditResponseEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type SettingChallengeTTLGetResponse struct {
	// ID of the zone setting.
	ID SettingChallengeTTLGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingChallengeTTLGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingChallengeTTLGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingChallengeTTLGetResponseJSON `json:"-"`
}

// settingChallengeTTLGetResponseJSON contains the JSON metadata for the struct
// [SettingChallengeTTLGetResponse]
type settingChallengeTTLGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingChallengeTTLGetResponseID string

const (
	SettingChallengeTTLGetResponseIDChallengeTTL SettingChallengeTTLGetResponseID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingChallengeTTLGetResponseValue float64

const (
	SettingChallengeTTLGetResponseValue300      SettingChallengeTTLGetResponseValue = 300
	SettingChallengeTTLGetResponseValue900      SettingChallengeTTLGetResponseValue = 900
	SettingChallengeTTLGetResponseValue1800     SettingChallengeTTLGetResponseValue = 1800
	SettingChallengeTTLGetResponseValue2700     SettingChallengeTTLGetResponseValue = 2700
	SettingChallengeTTLGetResponseValue3600     SettingChallengeTTLGetResponseValue = 3600
	SettingChallengeTTLGetResponseValue7200     SettingChallengeTTLGetResponseValue = 7200
	SettingChallengeTTLGetResponseValue10800    SettingChallengeTTLGetResponseValue = 10800
	SettingChallengeTTLGetResponseValue14400    SettingChallengeTTLGetResponseValue = 14400
	SettingChallengeTTLGetResponseValue28800    SettingChallengeTTLGetResponseValue = 28800
	SettingChallengeTTLGetResponseValue57600    SettingChallengeTTLGetResponseValue = 57600
	SettingChallengeTTLGetResponseValue86400    SettingChallengeTTLGetResponseValue = 86400
	SettingChallengeTTLGetResponseValue604800   SettingChallengeTTLGetResponseValue = 604800
	SettingChallengeTTLGetResponseValue2592000  SettingChallengeTTLGetResponseValue = 2592000
	SettingChallengeTTLGetResponseValue31536000 SettingChallengeTTLGetResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingChallengeTTLGetResponseEditable bool

const (
	SettingChallengeTTLGetResponseEditableTrue  SettingChallengeTTLGetResponseEditable = true
	SettingChallengeTTLGetResponseEditableFalse SettingChallengeTTLGetResponseEditable = false
)

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
	Result SettingChallengeTTLEditResponse             `json:"result"`
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
	Result SettingChallengeTTLGetResponse             `json:"result"`
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
