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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingOpportunisticEncryptionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingOpportunisticEncryptionService] method instead.
type SettingOpportunisticEncryptionService struct {
	Options []option.RequestOption
}

// NewSettingOpportunisticEncryptionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOpportunisticEncryptionService(opts ...option.RequestOption) (r *SettingOpportunisticEncryptionService) {
	r = &SettingOpportunisticEncryptionService{}
	r.Options = opts
	return
}

// Changes Opportunistic Encryption setting.
func (r *SettingOpportunisticEncryptionService) Edit(ctx context.Context, params SettingOpportunisticEncryptionEditParams, opts ...option.RequestOption) (res *ZonesOpportunisticEncryption, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticEncryptionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Opportunistic Encryption setting.
func (r *SettingOpportunisticEncryptionService) Get(ctx context.Context, query SettingOpportunisticEncryptionGetParams, opts ...option.RequestOption) (res *ZonesOpportunisticEncryption, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticEncryptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables the Opportunistic Encryption feature for a zone.
type ZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesOpportunisticEncryptionJSON `json:"-"`
}

// zonesOpportunisticEncryptionJSON contains the JSON metadata for the struct
// [ZonesOpportunisticEncryption]
type zonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesOpportunisticEncryption) implementsZonesSettingEditResponse() {}

func (r ZonesOpportunisticEncryption) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesOpportunisticEncryptionID string

const (
	ZonesOpportunisticEncryptionIDOpportunisticEncryption ZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r ZonesOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case ZonesOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesOpportunisticEncryptionValue string

const (
	ZonesOpportunisticEncryptionValueOn  ZonesOpportunisticEncryptionValue = "on"
	ZonesOpportunisticEncryptionValueOff ZonesOpportunisticEncryptionValue = "off"
)

func (r ZonesOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case ZonesOpportunisticEncryptionValueOn, ZonesOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesOpportunisticEncryptionEditable bool

const (
	ZonesOpportunisticEncryptionEditableTrue  ZonesOpportunisticEncryptionEditable = true
	ZonesOpportunisticEncryptionEditableFalse ZonesOpportunisticEncryptionEditable = false
)

func (r ZonesOpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case ZonesOpportunisticEncryptionEditableTrue, ZonesOpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type ZonesOpportunisticEncryptionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesOpportunisticEncryptionValue] `json:"value,required"`
}

func (r ZonesOpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOpportunisticEncryptionParam) implementsZonesSettingEditParamsItem() {}

type SettingOpportunisticEncryptionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingOpportunisticEncryptionEditParamsValue] `json:"value,required"`
}

func (r SettingOpportunisticEncryptionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingOpportunisticEncryptionEditParamsValue string

const (
	SettingOpportunisticEncryptionEditParamsValueOn  SettingOpportunisticEncryptionEditParamsValue = "on"
	SettingOpportunisticEncryptionEditParamsValueOff SettingOpportunisticEncryptionEditParamsValue = "off"
)

func (r SettingOpportunisticEncryptionEditParamsValue) IsKnown() bool {
	switch r {
	case SettingOpportunisticEncryptionEditParamsValueOn, SettingOpportunisticEncryptionEditParamsValueOff:
		return true
	}
	return false
}

type SettingOpportunisticEncryptionEditResponseEnvelope struct {
	Errors   []SettingOpportunisticEncryptionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticEncryptionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZonesOpportunisticEncryption                           `json:"result"`
	JSON   settingOpportunisticEncryptionEditResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticEncryptionEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingOpportunisticEncryptionEditResponseEnvelope]
type settingOpportunisticEncryptionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticEncryptionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticEncryptionEditResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    settingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionEditResponseEnvelopeErrors]
type settingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticEncryptionEditResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionEditResponseEnvelopeMessages]
type settingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticEncryptionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOpportunisticEncryptionGetResponseEnvelope struct {
	Errors   []SettingOpportunisticEncryptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticEncryptionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZonesOpportunisticEncryption                          `json:"result"`
	JSON   settingOpportunisticEncryptionGetResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOpportunisticEncryptionGetResponseEnvelope]
type settingOpportunisticEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticEncryptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticEncryptionGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionGetResponseEnvelopeErrors]
type settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticEncryptionGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionGetResponseEnvelopeMessages]
type settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
