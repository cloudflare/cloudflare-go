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
func (r *SettingOpportunisticEncryptionService) Edit(ctx context.Context, params SettingOpportunisticEncryptionEditParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticEncryption, err error) {
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
func (r *SettingOpportunisticEncryptionService) Get(ctx context.Context, query SettingOpportunisticEncryptionGetParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticEncryption, err error) {
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
type ZoneSettingOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOpportunisticEncryptionJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionJSON contains the JSON metadata for the struct
// [ZoneSettingOpportunisticEncryption]
type zoneSettingOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingOpportunisticEncryption) implementsZonesSettingEditResponse() {}

func (r ZoneSettingOpportunisticEncryption) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingOpportunisticEncryptionID string

const (
	ZoneSettingOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r ZoneSettingOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case ZoneSettingOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingOpportunisticEncryptionValue string

const (
	ZoneSettingOpportunisticEncryptionValueOn  ZoneSettingOpportunisticEncryptionValue = "on"
	ZoneSettingOpportunisticEncryptionValueOff ZoneSettingOpportunisticEncryptionValue = "off"
)

func (r ZoneSettingOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case ZoneSettingOpportunisticEncryptionValueOn, ZoneSettingOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticEncryptionEditable bool

const (
	ZoneSettingOpportunisticEncryptionEditableTrue  ZoneSettingOpportunisticEncryptionEditable = true
	ZoneSettingOpportunisticEncryptionEditableFalse ZoneSettingOpportunisticEncryptionEditable = false
)

func (r ZoneSettingOpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case ZoneSettingOpportunisticEncryptionEditableTrue, ZoneSettingOpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingOpportunisticEncryptionParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingOpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingOpportunisticEncryptionValue] `json:"value,required"`
}

func (r ZoneSettingOpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingOpportunisticEncryptionParam) implementsZonesSettingEditParamsItemUnion() {}

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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZoneSettingOpportunisticEncryption                     `json:"result"`
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

type SettingOpportunisticEncryptionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOpportunisticEncryptionGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZoneSettingOpportunisticEncryption                    `json:"result"`
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
