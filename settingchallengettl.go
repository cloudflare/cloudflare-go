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
func (r *SettingChallengeTTLService) Update(ctx context.Context, zoneID string, body SettingChallengeTTLUpdateParams, opts ...option.RequestOption) (res *SettingChallengeTTLUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingChallengeTTLUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingChallengeTTLService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingChallengeTTLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingChallengeTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zoneID)
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
type SettingChallengeTTLUpdateResponse struct {
	// ID of the zone setting.
	ID SettingChallengeTTLUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingChallengeTTLUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingChallengeTTLUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingChallengeTTLUpdateResponseJSON `json:"-"`
}

// settingChallengeTTLUpdateResponseJSON contains the JSON metadata for the struct
// [SettingChallengeTTLUpdateResponse]
type settingChallengeTTLUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingChallengeTTLUpdateResponseID string

const (
	SettingChallengeTTLUpdateResponseIDChallengeTTL SettingChallengeTTLUpdateResponseID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingChallengeTTLUpdateResponseValue float64

const (
	SettingChallengeTTLUpdateResponseValue300      SettingChallengeTTLUpdateResponseValue = 300
	SettingChallengeTTLUpdateResponseValue900      SettingChallengeTTLUpdateResponseValue = 900
	SettingChallengeTTLUpdateResponseValue1800     SettingChallengeTTLUpdateResponseValue = 1800
	SettingChallengeTTLUpdateResponseValue2700     SettingChallengeTTLUpdateResponseValue = 2700
	SettingChallengeTTLUpdateResponseValue3600     SettingChallengeTTLUpdateResponseValue = 3600
	SettingChallengeTTLUpdateResponseValue7200     SettingChallengeTTLUpdateResponseValue = 7200
	SettingChallengeTTLUpdateResponseValue10800    SettingChallengeTTLUpdateResponseValue = 10800
	SettingChallengeTTLUpdateResponseValue14400    SettingChallengeTTLUpdateResponseValue = 14400
	SettingChallengeTTLUpdateResponseValue28800    SettingChallengeTTLUpdateResponseValue = 28800
	SettingChallengeTTLUpdateResponseValue57600    SettingChallengeTTLUpdateResponseValue = 57600
	SettingChallengeTTLUpdateResponseValue86400    SettingChallengeTTLUpdateResponseValue = 86400
	SettingChallengeTTLUpdateResponseValue604800   SettingChallengeTTLUpdateResponseValue = 604800
	SettingChallengeTTLUpdateResponseValue2592000  SettingChallengeTTLUpdateResponseValue = 2592000
	SettingChallengeTTLUpdateResponseValue31536000 SettingChallengeTTLUpdateResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingChallengeTTLUpdateResponseEditable bool

const (
	SettingChallengeTTLUpdateResponseEditableTrue  SettingChallengeTTLUpdateResponseEditable = true
	SettingChallengeTTLUpdateResponseEditableFalse SettingChallengeTTLUpdateResponseEditable = false
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

type SettingChallengeTTLUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingChallengeTTLUpdateParamsValue] `json:"value,required"`
}

func (r SettingChallengeTTLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingChallengeTTLUpdateParamsValue float64

const (
	SettingChallengeTTLUpdateParamsValue300      SettingChallengeTTLUpdateParamsValue = 300
	SettingChallengeTTLUpdateParamsValue900      SettingChallengeTTLUpdateParamsValue = 900
	SettingChallengeTTLUpdateParamsValue1800     SettingChallengeTTLUpdateParamsValue = 1800
	SettingChallengeTTLUpdateParamsValue2700     SettingChallengeTTLUpdateParamsValue = 2700
	SettingChallengeTTLUpdateParamsValue3600     SettingChallengeTTLUpdateParamsValue = 3600
	SettingChallengeTTLUpdateParamsValue7200     SettingChallengeTTLUpdateParamsValue = 7200
	SettingChallengeTTLUpdateParamsValue10800    SettingChallengeTTLUpdateParamsValue = 10800
	SettingChallengeTTLUpdateParamsValue14400    SettingChallengeTTLUpdateParamsValue = 14400
	SettingChallengeTTLUpdateParamsValue28800    SettingChallengeTTLUpdateParamsValue = 28800
	SettingChallengeTTLUpdateParamsValue57600    SettingChallengeTTLUpdateParamsValue = 57600
	SettingChallengeTTLUpdateParamsValue86400    SettingChallengeTTLUpdateParamsValue = 86400
	SettingChallengeTTLUpdateParamsValue604800   SettingChallengeTTLUpdateParamsValue = 604800
	SettingChallengeTTLUpdateParamsValue2592000  SettingChallengeTTLUpdateParamsValue = 2592000
	SettingChallengeTTLUpdateParamsValue31536000 SettingChallengeTTLUpdateParamsValue = 31536000
)

type SettingChallengeTTLUpdateResponseEnvelope struct {
	Errors   []SettingChallengeTTLUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingChallengeTTLUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result SettingChallengeTTLUpdateResponse             `json:"result"`
	JSON   settingChallengeTTLUpdateResponseEnvelopeJSON `json:"-"`
}

// settingChallengeTTLUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingChallengeTTLUpdateResponseEnvelope]
type settingChallengeTTLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingChallengeTTLUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingChallengeTTLUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingChallengeTTLUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingChallengeTTLUpdateResponseEnvelopeErrors]
type settingChallengeTTLUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingChallengeTTLUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingChallengeTTLUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingChallengeTTLUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingChallengeTTLUpdateResponseEnvelopeMessages]
type settingChallengeTTLUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingChallengeTTLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
