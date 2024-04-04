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
func (r *SettingPseudoIPV4Service) Edit(ctx context.Context, params SettingPseudoIPV4EditParams, opts ...option.RequestOption) (res *PseudoIPV4, err error) {
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
func (r *SettingPseudoIPV4Service) Get(ctx context.Context, query SettingPseudoIPV4GetParams, opts ...option.RequestOption) (res *PseudoIPV4, err error) {
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
type PseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID PseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value PseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       pseudoIPV4JSON `json:"-"`
}

// pseudoIPV4JSON contains the JSON metadata for the struct [PseudoIPV4]
type pseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pseudoIPV4JSON) RawJSON() string {
	return r.raw
}

// Value of the Pseudo IPv4 setting.
type PseudoIPV4ID string

const (
	PseudoIPV4IDPseudoIPV4 PseudoIPV4ID = "pseudo_ipv4"
)

func (r PseudoIPV4ID) IsKnown() bool {
	switch r {
	case PseudoIPV4IDPseudoIPV4:
		return true
	}
	return false
}

// Current value of the zone setting.
type PseudoIPV4Value string

const (
	PseudoIPV4ValueOff             PseudoIPV4Value = "off"
	PseudoIPV4ValueAddHeader       PseudoIPV4Value = "add_header"
	PseudoIPV4ValueOverwriteHeader PseudoIPV4Value = "overwrite_header"
)

func (r PseudoIPV4Value) IsKnown() bool {
	switch r {
	case PseudoIPV4ValueOff, PseudoIPV4ValueAddHeader, PseudoIPV4ValueOverwriteHeader:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type PseudoIPV4Editable bool

const (
	PseudoIPV4EditableTrue  PseudoIPV4Editable = true
	PseudoIPV4EditableFalse PseudoIPV4Editable = false
)

func (r PseudoIPV4Editable) IsKnown() bool {
	switch r {
	case PseudoIPV4EditableTrue, PseudoIPV4EditableFalse:
		return true
	}
	return false
}

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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result PseudoIPV4                                `json:"result"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result PseudoIPV4                               `json:"result"`
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
