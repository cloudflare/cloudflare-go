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

// SettingSecurityLevelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingSecurityLevelService]
// method instead.
type SettingSecurityLevelService struct {
	Options []option.RequestOption
}

// NewSettingSecurityLevelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingSecurityLevelService(opts ...option.RequestOption) (r *SettingSecurityLevelService) {
	r = &SettingSecurityLevelService{}
	r.Options = opts
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *SettingSecurityLevelService) Edit(ctx context.Context, params SettingSecurityLevelEditParams, opts ...option.RequestOption) (res *ZoneSettingSecurityLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityLevelEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *SettingSecurityLevelService) Get(ctx context.Context, query SettingSecurityLevelGetParams, opts ...option.RequestOption) (res *ZoneSettingSecurityLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityLevelJSON `json:"-"`
}

// zoneSettingSecurityLevelJSON contains the JSON metadata for the struct
// [ZoneSettingSecurityLevel]
type zoneSettingSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingSecurityLevel) implementsZonesSettingEditResponse() {}

func (r ZoneSettingSecurityLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingSecurityLevelID string

const (
	ZoneSettingSecurityLevelIDSecurityLevel ZoneSettingSecurityLevelID = "security_level"
)

func (r ZoneSettingSecurityLevelID) IsKnown() bool {
	switch r {
	case ZoneSettingSecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingSecurityLevelValue string

const (
	ZoneSettingSecurityLevelValueOff            ZoneSettingSecurityLevelValue = "off"
	ZoneSettingSecurityLevelValueEssentiallyOff ZoneSettingSecurityLevelValue = "essentially_off"
	ZoneSettingSecurityLevelValueLow            ZoneSettingSecurityLevelValue = "low"
	ZoneSettingSecurityLevelValueMedium         ZoneSettingSecurityLevelValue = "medium"
	ZoneSettingSecurityLevelValueHigh           ZoneSettingSecurityLevelValue = "high"
	ZoneSettingSecurityLevelValueUnderAttack    ZoneSettingSecurityLevelValue = "under_attack"
)

func (r ZoneSettingSecurityLevelValue) IsKnown() bool {
	switch r {
	case ZoneSettingSecurityLevelValueOff, ZoneSettingSecurityLevelValueEssentiallyOff, ZoneSettingSecurityLevelValueLow, ZoneSettingSecurityLevelValueMedium, ZoneSettingSecurityLevelValueHigh, ZoneSettingSecurityLevelValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityLevelEditable bool

const (
	ZoneSettingSecurityLevelEditableTrue  ZoneSettingSecurityLevelEditable = true
	ZoneSettingSecurityLevelEditableFalse ZoneSettingSecurityLevelEditable = false
)

func (r ZoneSettingSecurityLevelEditable) IsKnown() bool {
	switch r {
	case ZoneSettingSecurityLevelEditableTrue, ZoneSettingSecurityLevelEditableFalse:
		return true
	}
	return false
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingSecurityLevelParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingSecurityLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingSecurityLevelValue] `json:"value,required"`
}

func (r ZoneSettingSecurityLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingSecurityLevelParam) implementsZonesSettingEditParamsItem() {}

type SettingSecurityLevelEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingSecurityLevelEditParamsValue] `json:"value,required"`
}

func (r SettingSecurityLevelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingSecurityLevelEditParamsValue string

const (
	SettingSecurityLevelEditParamsValueOff            SettingSecurityLevelEditParamsValue = "off"
	SettingSecurityLevelEditParamsValueEssentiallyOff SettingSecurityLevelEditParamsValue = "essentially_off"
	SettingSecurityLevelEditParamsValueLow            SettingSecurityLevelEditParamsValue = "low"
	SettingSecurityLevelEditParamsValueMedium         SettingSecurityLevelEditParamsValue = "medium"
	SettingSecurityLevelEditParamsValueHigh           SettingSecurityLevelEditParamsValue = "high"
	SettingSecurityLevelEditParamsValueUnderAttack    SettingSecurityLevelEditParamsValue = "under_attack"
)

func (r SettingSecurityLevelEditParamsValue) IsKnown() bool {
	switch r {
	case SettingSecurityLevelEditParamsValueOff, SettingSecurityLevelEditParamsValueEssentiallyOff, SettingSecurityLevelEditParamsValueLow, SettingSecurityLevelEditParamsValueMedium, SettingSecurityLevelEditParamsValueHigh, SettingSecurityLevelEditParamsValueUnderAttack:
		return true
	}
	return false
}

type SettingSecurityLevelEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result ZoneSettingSecurityLevel                     `json:"result"`
	JSON   settingSecurityLevelEditResponseEnvelopeJSON `json:"-"`
}

// settingSecurityLevelEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityLevelEditResponseEnvelope]
type settingSecurityLevelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityLevelGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSecurityLevelGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result ZoneSettingSecurityLevel                    `json:"result"`
	JSON   settingSecurityLevelGetResponseEnvelopeJSON `json:"-"`
}

// settingSecurityLevelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityLevelGetResponseEnvelope]
type settingSecurityLevelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
