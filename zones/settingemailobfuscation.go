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

// SettingEmailObfuscationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingEmailObfuscationService] method instead.
type SettingEmailObfuscationService struct {
	Options []option.RequestOption
}

// NewSettingEmailObfuscationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingEmailObfuscationService(opts ...option.RequestOption) (r *SettingEmailObfuscationService) {
	r = &SettingEmailObfuscationService{}
	r.Options = opts
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *SettingEmailObfuscationService) Edit(ctx context.Context, params SettingEmailObfuscationEditParams, opts ...option.RequestOption) (res *ZoneSettingEmailObfuscation, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEmailObfuscationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *SettingEmailObfuscationService) Get(ctx context.Context, query SettingEmailObfuscationGetParams, opts ...option.RequestOption) (res *ZoneSettingEmailObfuscation, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEmailObfuscationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEmailObfuscationJSON `json:"-"`
}

// zoneSettingEmailObfuscationJSON contains the JSON metadata for the struct
// [ZoneSettingEmailObfuscation]
type zoneSettingEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingEmailObfuscation) implementsZonesSettingEditResponse() {}

func (r ZoneSettingEmailObfuscation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingEmailObfuscationID string

const (
	ZoneSettingEmailObfuscationIDEmailObfuscation ZoneSettingEmailObfuscationID = "email_obfuscation"
)

func (r ZoneSettingEmailObfuscationID) IsKnown() bool {
	switch r {
	case ZoneSettingEmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingEmailObfuscationValue string

const (
	ZoneSettingEmailObfuscationValueOn  ZoneSettingEmailObfuscationValue = "on"
	ZoneSettingEmailObfuscationValueOff ZoneSettingEmailObfuscationValue = "off"
)

func (r ZoneSettingEmailObfuscationValue) IsKnown() bool {
	switch r {
	case ZoneSettingEmailObfuscationValueOn, ZoneSettingEmailObfuscationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEmailObfuscationEditable bool

const (
	ZoneSettingEmailObfuscationEditableTrue  ZoneSettingEmailObfuscationEditable = true
	ZoneSettingEmailObfuscationEditableFalse ZoneSettingEmailObfuscationEditable = false
)

func (r ZoneSettingEmailObfuscationEditable) IsKnown() bool {
	switch r {
	case ZoneSettingEmailObfuscationEditableTrue, ZoneSettingEmailObfuscationEditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEmailObfuscationParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEmailObfuscationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEmailObfuscationValue] `json:"value,required"`
}

func (r ZoneSettingEmailObfuscationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEmailObfuscationParam) implementsZonesSettingEditParamsItem() {}

type SettingEmailObfuscationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingEmailObfuscationEditParamsValue] `json:"value,required"`
}

func (r SettingEmailObfuscationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingEmailObfuscationEditParamsValue string

const (
	SettingEmailObfuscationEditParamsValueOn  SettingEmailObfuscationEditParamsValue = "on"
	SettingEmailObfuscationEditParamsValueOff SettingEmailObfuscationEditParamsValue = "off"
)

func (r SettingEmailObfuscationEditParamsValue) IsKnown() bool {
	switch r {
	case SettingEmailObfuscationEditParamsValueOn, SettingEmailObfuscationEditParamsValueOff:
		return true
	}
	return false
}

type SettingEmailObfuscationEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result ZoneSettingEmailObfuscation                     `json:"result"`
	JSON   settingEmailObfuscationEditResponseEnvelopeJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingEmailObfuscationEditResponseEnvelope]
type settingEmailObfuscationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingEmailObfuscationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingEmailObfuscationGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result ZoneSettingEmailObfuscation                    `json:"result"`
	JSON   settingEmailObfuscationGetResponseEnvelopeJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingEmailObfuscationGetResponseEnvelope]
type settingEmailObfuscationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
