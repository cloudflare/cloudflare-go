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
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingOpportunisticEncryptionService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingOpportunisticEncryptionService] method instead.
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
func (r *SettingOpportunisticEncryptionService) Edit(ctx context.Context, params SettingOpportunisticEncryptionEditParams, opts ...option.RequestOption) (res *OpportunisticEncryption, err error) {
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
func (r *SettingOpportunisticEncryptionService) Get(ctx context.Context, query SettingOpportunisticEncryptionGetParams, opts ...option.RequestOption) (res *OpportunisticEncryption, err error) {
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
type OpportunisticEncryption struct {
	// ID of the zone setting.
	ID OpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value OpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       opportunisticEncryptionJSON `json:"-"`
}

// opportunisticEncryptionJSON contains the JSON metadata for the struct
// [OpportunisticEncryption]
type opportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r opportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type OpportunisticEncryptionID string

const (
	OpportunisticEncryptionIDOpportunisticEncryption OpportunisticEncryptionID = "opportunistic_encryption"
)

func (r OpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case OpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type OpportunisticEncryptionValue string

const (
	OpportunisticEncryptionValueOn  OpportunisticEncryptionValue = "on"
	OpportunisticEncryptionValueOff OpportunisticEncryptionValue = "off"
)

func (r OpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case OpportunisticEncryptionValueOn, OpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OpportunisticEncryptionEditable bool

const (
	OpportunisticEncryptionEditableTrue  OpportunisticEncryptionEditable = true
	OpportunisticEncryptionEditableFalse OpportunisticEncryptionEditable = false
)

func (r OpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case OpportunisticEncryptionEditableTrue, OpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result OpportunisticEncryption                                `json:"result"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result OpportunisticEncryption                               `json:"result"`
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
