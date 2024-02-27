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
func (r *ZoneSettingSecurityLevelService) Edit(ctx context.Context, params ZoneSettingSecurityLevelEditParams, opts ...option.RequestOption) (res *ZoneSettingSecurityLevelEditResponse, err error) {
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
func (r *ZoneSettingSecurityLevelService) Get(ctx context.Context, query ZoneSettingSecurityLevelGetParams, opts ...option.RequestOption) (res *ZoneSettingSecurityLevelGetResponse, err error) {
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
type ZoneSettingSecurityLevelEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingSecurityLevelEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSecurityLevelEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityLevelEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityLevelEditResponseJSON `json:"-"`
}

// zoneSettingSecurityLevelEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityLevelEditResponse]
type zoneSettingSecurityLevelEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSecurityLevelEditResponseID string

const (
	ZoneSettingSecurityLevelEditResponseIDSecurityLevel ZoneSettingSecurityLevelEditResponseID = "security_level"
)

// Current value of the zone setting.
type ZoneSettingSecurityLevelEditResponseValue string

const (
	ZoneSettingSecurityLevelEditResponseValueOff            ZoneSettingSecurityLevelEditResponseValue = "off"
	ZoneSettingSecurityLevelEditResponseValueEssentiallyOff ZoneSettingSecurityLevelEditResponseValue = "essentially_off"
	ZoneSettingSecurityLevelEditResponseValueLow            ZoneSettingSecurityLevelEditResponseValue = "low"
	ZoneSettingSecurityLevelEditResponseValueMedium         ZoneSettingSecurityLevelEditResponseValue = "medium"
	ZoneSettingSecurityLevelEditResponseValueHigh           ZoneSettingSecurityLevelEditResponseValue = "high"
	ZoneSettingSecurityLevelEditResponseValueUnderAttack    ZoneSettingSecurityLevelEditResponseValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityLevelEditResponseEditable bool

const (
	ZoneSettingSecurityLevelEditResponseEditableTrue  ZoneSettingSecurityLevelEditResponseEditable = true
	ZoneSettingSecurityLevelEditResponseEditableFalse ZoneSettingSecurityLevelEditResponseEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingSecurityLevelGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingSecurityLevelGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSecurityLevelGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityLevelGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityLevelGetResponseJSON `json:"-"`
}

// zoneSettingSecurityLevelGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityLevelGetResponse]
type zoneSettingSecurityLevelGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSecurityLevelGetResponseID string

const (
	ZoneSettingSecurityLevelGetResponseIDSecurityLevel ZoneSettingSecurityLevelGetResponseID = "security_level"
)

// Current value of the zone setting.
type ZoneSettingSecurityLevelGetResponseValue string

const (
	ZoneSettingSecurityLevelGetResponseValueOff            ZoneSettingSecurityLevelGetResponseValue = "off"
	ZoneSettingSecurityLevelGetResponseValueEssentiallyOff ZoneSettingSecurityLevelGetResponseValue = "essentially_off"
	ZoneSettingSecurityLevelGetResponseValueLow            ZoneSettingSecurityLevelGetResponseValue = "low"
	ZoneSettingSecurityLevelGetResponseValueMedium         ZoneSettingSecurityLevelGetResponseValue = "medium"
	ZoneSettingSecurityLevelGetResponseValueHigh           ZoneSettingSecurityLevelGetResponseValue = "high"
	ZoneSettingSecurityLevelGetResponseValueUnderAttack    ZoneSettingSecurityLevelGetResponseValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityLevelGetResponseEditable bool

const (
	ZoneSettingSecurityLevelGetResponseEditableTrue  ZoneSettingSecurityLevelGetResponseEditable = true
	ZoneSettingSecurityLevelGetResponseEditableFalse ZoneSettingSecurityLevelGetResponseEditable = false
)

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
	Result ZoneSettingSecurityLevelEditResponse             `json:"result"`
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
	Result ZoneSettingSecurityLevelGetResponse             `json:"result"`
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
