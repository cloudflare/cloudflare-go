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

// SettingPseudoIPV4Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingPseudoIPV4Service] method
// instead.
type SettingPseudoIPV4Service struct {
	Options []option.RequestOption
}

// NewSettingPseudoIPV4Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingPseudoIPV4Service(opts ...option.RequestOption) (r *SettingPseudoIPV4Service) {
	r = &SettingPseudoIPV4Service{}
	r.Options = opts
	return
}

// Value of the Pseudo IPv4 setting.
func (r *SettingPseudoIPV4Service) Edit(ctx context.Context, params SettingPseudoIPV4EditParams, opts ...option.RequestOption) (res *ZoneSettingPseudoIPV4, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPseudoIPV4EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the Pseudo IPv4 setting.
func (r *SettingPseudoIPV4Service) Get(ctx context.Context, query SettingPseudoIPV4GetParams, opts ...option.RequestOption) (res *ZoneSettingPseudoIPV4, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPseudoIPV4GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The value set for the Pseudo IPv4 setting.
type ZoneSettingPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPseudoIPV4JSON `json:"-"`
}

// zoneSettingPseudoIPV4JSON contains the JSON metadata for the struct
// [ZoneSettingPseudoIPV4]
type zoneSettingPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPseudoIPV4JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingPseudoIPV4) implementsZonesSettingEditResponse() {}

func (r ZoneSettingPseudoIPV4) implementsZonesSettingGetResponse() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIPV4ID string

const (
	ZoneSettingPseudoIPV4IDPseudoIPV4 ZoneSettingPseudoIPV4ID = "pseudo_ipv4"
)

func (r ZoneSettingPseudoIPV4ID) IsKnown() bool {
	switch r {
	case ZoneSettingPseudoIPV4IDPseudoIPV4:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingPseudoIPV4Value string

const (
	ZoneSettingPseudoIPV4ValueOff             ZoneSettingPseudoIPV4Value = "off"
	ZoneSettingPseudoIPV4ValueAddHeader       ZoneSettingPseudoIPV4Value = "add_header"
	ZoneSettingPseudoIPV4ValueOverwriteHeader ZoneSettingPseudoIPV4Value = "overwrite_header"
)

func (r ZoneSettingPseudoIPV4Value) IsKnown() bool {
	switch r {
	case ZoneSettingPseudoIPV4ValueOff, ZoneSettingPseudoIPV4ValueAddHeader, ZoneSettingPseudoIPV4ValueOverwriteHeader:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPseudoIPV4Editable bool

const (
	ZoneSettingPseudoIPV4EditableTrue  ZoneSettingPseudoIPV4Editable = true
	ZoneSettingPseudoIPV4EditableFalse ZoneSettingPseudoIPV4Editable = false
)

func (r ZoneSettingPseudoIPV4Editable) IsKnown() bool {
	switch r {
	case ZoneSettingPseudoIPV4EditableTrue, ZoneSettingPseudoIPV4EditableFalse:
		return true
	}
	return false
}

// The value set for the Pseudo IPv4 setting.
type ZoneSettingPseudoIPV4Param struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[ZoneSettingPseudoIPV4ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingPseudoIPV4Value] `json:"value,required"`
}

func (r ZoneSettingPseudoIPV4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingPseudoIPV4Param) implementsZonesSettingEditParamsItem() {}

type SettingPseudoIPV4EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Pseudo IPv4 setting.
	Value param.Field[SettingPseudoIPV4EditParamsValue] `json:"value,required"`
}

func (r SettingPseudoIPV4EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Pseudo IPv4 setting.
type SettingPseudoIPV4EditParamsValue string

const (
	SettingPseudoIPV4EditParamsValueOff             SettingPseudoIPV4EditParamsValue = "off"
	SettingPseudoIPV4EditParamsValueAddHeader       SettingPseudoIPV4EditParamsValue = "add_header"
	SettingPseudoIPV4EditParamsValueOverwriteHeader SettingPseudoIPV4EditParamsValue = "overwrite_header"
)

func (r SettingPseudoIPV4EditParamsValue) IsKnown() bool {
	switch r {
	case SettingPseudoIPV4EditParamsValueOff, SettingPseudoIPV4EditParamsValueAddHeader, SettingPseudoIPV4EditParamsValueOverwriteHeader:
		return true
	}
	return false
}

type SettingPseudoIPV4EditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result ZoneSettingPseudoIPV4                     `json:"result"`
	JSON   settingPseudoIPV4EditResponseEnvelopeJSON `json:"-"`
}

// settingPseudoIPV4EditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPseudoIPV4EditResponseEnvelope]
type settingPseudoIPV4EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPseudoIPV4EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingPseudoIPV4GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingPseudoIPV4GetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result ZoneSettingPseudoIPV4                    `json:"result"`
	JSON   settingPseudoIPV4GetResponseEnvelopeJSON `json:"-"`
}

// settingPseudoIPV4GetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPseudoIPV4GetResponseEnvelope]
type settingPseudoIPV4GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPseudoIPV4GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
