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
func (r *SettingChallengeTTLService) Edit(ctx context.Context, params SettingChallengeTTLEditParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTTL, err error) {
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
func (r *SettingChallengeTTLService) Get(ctx context.Context, query SettingChallengeTTLGetParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTTL, err error) {
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
type ZoneSettingChallengeTTL struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTTLJSON `json:"-"`
}

// zoneSettingChallengeTTLJSON contains the JSON metadata for the struct
// [ZoneSettingChallengeTTL]
type zoneSettingChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingChallengeTTLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingChallengeTTL) implementsZonesSettingEditResponse() {}

func (r ZoneSettingChallengeTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingChallengeTTLID string

const (
	ZoneSettingChallengeTTLIDChallengeTTL ZoneSettingChallengeTTLID = "challenge_ttl"
)

func (r ZoneSettingChallengeTTLID) IsKnown() bool {
	switch r {
	case ZoneSettingChallengeTTLIDChallengeTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingChallengeTTLValue float64

const (
	ZoneSettingChallengeTTLValue300      ZoneSettingChallengeTTLValue = 300
	ZoneSettingChallengeTTLValue900      ZoneSettingChallengeTTLValue = 900
	ZoneSettingChallengeTTLValue1800     ZoneSettingChallengeTTLValue = 1800
	ZoneSettingChallengeTTLValue2700     ZoneSettingChallengeTTLValue = 2700
	ZoneSettingChallengeTTLValue3600     ZoneSettingChallengeTTLValue = 3600
	ZoneSettingChallengeTTLValue7200     ZoneSettingChallengeTTLValue = 7200
	ZoneSettingChallengeTTLValue10800    ZoneSettingChallengeTTLValue = 10800
	ZoneSettingChallengeTTLValue14400    ZoneSettingChallengeTTLValue = 14400
	ZoneSettingChallengeTTLValue28800    ZoneSettingChallengeTTLValue = 28800
	ZoneSettingChallengeTTLValue57600    ZoneSettingChallengeTTLValue = 57600
	ZoneSettingChallengeTTLValue86400    ZoneSettingChallengeTTLValue = 86400
	ZoneSettingChallengeTTLValue604800   ZoneSettingChallengeTTLValue = 604800
	ZoneSettingChallengeTTLValue2592000  ZoneSettingChallengeTTLValue = 2592000
	ZoneSettingChallengeTTLValue31536000 ZoneSettingChallengeTTLValue = 31536000
)

func (r ZoneSettingChallengeTTLValue) IsKnown() bool {
	switch r {
	case ZoneSettingChallengeTTLValue300, ZoneSettingChallengeTTLValue900, ZoneSettingChallengeTTLValue1800, ZoneSettingChallengeTTLValue2700, ZoneSettingChallengeTTLValue3600, ZoneSettingChallengeTTLValue7200, ZoneSettingChallengeTTLValue10800, ZoneSettingChallengeTTLValue14400, ZoneSettingChallengeTTLValue28800, ZoneSettingChallengeTTLValue57600, ZoneSettingChallengeTTLValue86400, ZoneSettingChallengeTTLValue604800, ZoneSettingChallengeTTLValue2592000, ZoneSettingChallengeTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLEditable bool

const (
	ZoneSettingChallengeTTLEditableTrue  ZoneSettingChallengeTTLEditable = true
	ZoneSettingChallengeTTLEditableFalse ZoneSettingChallengeTTLEditable = false
)

func (r ZoneSettingChallengeTTLEditable) IsKnown() bool {
	switch r {
	case ZoneSettingChallengeTTLEditableTrue, ZoneSettingChallengeTTLEditableFalse:
		return true
	}
	return false
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingChallengeTTLParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingChallengeTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingChallengeTTLValue] `json:"value,required"`
}

func (r ZoneSettingChallengeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingChallengeTTLParam) implementsZonesSettingEditParamsItem() {}

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

func (r SettingChallengeTTLEditParamsValue) IsKnown() bool {
	switch r {
	case SettingChallengeTTLEditParamsValue300, SettingChallengeTTLEditParamsValue900, SettingChallengeTTLEditParamsValue1800, SettingChallengeTTLEditParamsValue2700, SettingChallengeTTLEditParamsValue3600, SettingChallengeTTLEditParamsValue7200, SettingChallengeTTLEditParamsValue10800, SettingChallengeTTLEditParamsValue14400, SettingChallengeTTLEditParamsValue28800, SettingChallengeTTLEditParamsValue57600, SettingChallengeTTLEditParamsValue86400, SettingChallengeTTLEditParamsValue604800, SettingChallengeTTLEditParamsValue2592000, SettingChallengeTTLEditParamsValue31536000:
		return true
	}
	return false
}

type SettingChallengeTTLEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZoneSettingChallengeTTL                     `json:"result"`
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

type SettingChallengeTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingChallengeTTLGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZoneSettingChallengeTTL                    `json:"result"`
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
