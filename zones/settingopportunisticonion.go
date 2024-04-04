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

// SettingOpportunisticOnionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingOpportunisticOnionService] method instead.
type SettingOpportunisticOnionService struct {
	Options []option.RequestOption
}

// NewSettingOpportunisticOnionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOpportunisticOnionService(opts ...option.RequestOption) (r *SettingOpportunisticOnionService) {
	r = &SettingOpportunisticOnionService{}
	r.Options = opts
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *SettingOpportunisticOnionService) Edit(ctx context.Context, params SettingOpportunisticOnionEditParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *SettingOpportunisticOnionService) Get(ctx context.Context, query SettingOpportunisticOnionGetParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOpportunisticOnionJSON `json:"-"`
}

// zoneSettingOpportunisticOnionJSON contains the JSON metadata for the struct
// [ZoneSettingOpportunisticOnion]
type zoneSettingOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOpportunisticOnionJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingOpportunisticOnionID string

const (
	ZoneSettingOpportunisticOnionIDOpportunisticOnion ZoneSettingOpportunisticOnionID = "opportunistic_onion"
)

func (r ZoneSettingOpportunisticOnionID) IsKnown() bool {
	switch r {
	case ZoneSettingOpportunisticOnionIDOpportunisticOnion:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingOpportunisticOnionValue string

const (
	ZoneSettingOpportunisticOnionValueOn  ZoneSettingOpportunisticOnionValue = "on"
	ZoneSettingOpportunisticOnionValueOff ZoneSettingOpportunisticOnionValue = "off"
)

func (r ZoneSettingOpportunisticOnionValue) IsKnown() bool {
	switch r {
	case ZoneSettingOpportunisticOnionValueOn, ZoneSettingOpportunisticOnionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticOnionEditable bool

const (
	ZoneSettingOpportunisticOnionEditableTrue  ZoneSettingOpportunisticOnionEditable = true
	ZoneSettingOpportunisticOnionEditableFalse ZoneSettingOpportunisticOnionEditable = false
)

func (r ZoneSettingOpportunisticOnionEditable) IsKnown() bool {
	switch r {
	case ZoneSettingOpportunisticOnionEditableTrue, ZoneSettingOpportunisticOnionEditableFalse:
		return true
	}
	return false
}

type SettingOpportunisticOnionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingOpportunisticOnionEditParamsValue] `json:"value,required"`
}

func (r SettingOpportunisticOnionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingOpportunisticOnionEditParamsValue string

const (
	SettingOpportunisticOnionEditParamsValueOn  SettingOpportunisticOnionEditParamsValue = "on"
	SettingOpportunisticOnionEditParamsValueOff SettingOpportunisticOnionEditParamsValue = "off"
)

func (r SettingOpportunisticOnionEditParamsValue) IsKnown() bool {
	switch r {
	case SettingOpportunisticOnionEditParamsValueOn, SettingOpportunisticOnionEditParamsValueOff:
		return true
	}
	return false
}

type SettingOpportunisticOnionEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZoneSettingOpportunisticOnion                     `json:"result"`
	JSON   settingOpportunisticOnionEditResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticOnionEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingOpportunisticOnionEditResponseEnvelope]
type settingOpportunisticOnionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticOnionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOpportunisticOnionGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZoneSettingOpportunisticOnion                    `json:"result"`
	JSON   settingOpportunisticOnionGetResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingOpportunisticOnionGetResponseEnvelope]
type settingOpportunisticOnionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
