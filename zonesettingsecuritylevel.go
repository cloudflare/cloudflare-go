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

// ZoneSettingSecurityLevelService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingSecurityLevelService] method instead.
type ZoneSettingSecurityLevelService struct {
	Options []option.RequestOption
}

// NewZoneSettingSecurityLevelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSecurityLevelService(opts ...option.RequestOption) (r *ZoneSettingSecurityLevelService) {
	r = &ZoneSettingSecurityLevelService{}
	r.Options = opts
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *ZoneSettingSecurityLevelService) Edit(ctx context.Context, params ZoneSettingSecurityLevelEditParams, opts ...option.RequestOption) (res *ZonesSecurityLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSecurityLevelEditResponseEnvelope
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
func (r *ZoneSettingSecurityLevelService) Get(ctx context.Context, query ZoneSettingSecurityLevelGetParams, opts ...option.RequestOption) (res *ZonesSecurityLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSecurityLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type ZonesSecurityLevel struct {
	// ID of the zone setting.
	ID ZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesSecurityLevelJSON `json:"-"`
}

// zonesSecurityLevelJSON contains the JSON metadata for the struct
// [ZonesSecurityLevel]
type zonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesSecurityLevel) implementsZoneSettingEditResponse() {}

func (r ZonesSecurityLevel) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesSecurityLevelID string

const (
	ZonesSecurityLevelIDSecurityLevel ZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type ZonesSecurityLevelValue string

const (
	ZonesSecurityLevelValueOff            ZonesSecurityLevelValue = "off"
	ZonesSecurityLevelValueEssentiallyOff ZonesSecurityLevelValue = "essentially_off"
	ZonesSecurityLevelValueLow            ZonesSecurityLevelValue = "low"
	ZonesSecurityLevelValueMedium         ZonesSecurityLevelValue = "medium"
	ZonesSecurityLevelValueHigh           ZonesSecurityLevelValue = "high"
	ZonesSecurityLevelValueUnderAttack    ZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesSecurityLevelEditable bool

const (
	ZonesSecurityLevelEditableTrue  ZonesSecurityLevelEditable = true
	ZonesSecurityLevelEditableFalse ZonesSecurityLevelEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZonesSecurityLevelParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSecurityLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesSecurityLevelValue] `json:"value,required"`
}

func (r ZonesSecurityLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSecurityLevelParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingSecurityLevelEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingSecurityLevelEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingSecurityLevelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingSecurityLevelEditParamsValue string

const (
	ZoneSettingSecurityLevelEditParamsValueOff            ZoneSettingSecurityLevelEditParamsValue = "off"
	ZoneSettingSecurityLevelEditParamsValueEssentiallyOff ZoneSettingSecurityLevelEditParamsValue = "essentially_off"
	ZoneSettingSecurityLevelEditParamsValueLow            ZoneSettingSecurityLevelEditParamsValue = "low"
	ZoneSettingSecurityLevelEditParamsValueMedium         ZoneSettingSecurityLevelEditParamsValue = "medium"
	ZoneSettingSecurityLevelEditParamsValueHigh           ZoneSettingSecurityLevelEditParamsValue = "high"
	ZoneSettingSecurityLevelEditParamsValueUnderAttack    ZoneSettingSecurityLevelEditParamsValue = "under_attack"
)

type ZoneSettingSecurityLevelEditResponseEnvelope struct {
	Errors   []ZoneSettingSecurityLevelEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSecurityLevelEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result ZonesSecurityLevel                               `json:"result"`
	JSON   zoneSettingSecurityLevelEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSecurityLevelEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelEditResponseEnvelope]
type zoneSettingSecurityLevelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingSecurityLevelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSecurityLevelEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityLevelEditResponseEnvelopeErrors]
type zoneSettingSecurityLevelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingSecurityLevelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSecurityLevelEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityLevelEditResponseEnvelopeMessages]
type zoneSettingSecurityLevelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingSecurityLevelGetResponseEnvelope struct {
	Errors   []ZoneSettingSecurityLevelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSecurityLevelGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result ZonesSecurityLevel                              `json:"result"`
	JSON   zoneSettingSecurityLevelGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSecurityLevelGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelGetResponseEnvelope]
type zoneSettingSecurityLevelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingSecurityLevelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSecurityLevelGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingSecurityLevelGetResponseEnvelopeErrors]
type zoneSettingSecurityLevelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingSecurityLevelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSecurityLevelGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityLevelGetResponseEnvelopeMessages]
type zoneSettingSecurityLevelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
