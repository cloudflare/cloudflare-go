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

// SettingAutomaticHTTPSRewriteService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingAutomaticHTTPSRewriteService] method instead.
type SettingAutomaticHTTPSRewriteService struct {
	Options []option.RequestOption
}

// NewSettingAutomaticHTTPSRewriteService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingAutomaticHTTPSRewriteService(opts ...option.RequestOption) (r *SettingAutomaticHTTPSRewriteService) {
	r = &SettingAutomaticHTTPSRewriteService{}
	r.Options = opts
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPSRewriteService) Edit(ctx context.Context, params SettingAutomaticHTTPSRewriteEditParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticHTTPSRewrites, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPSRewriteEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPSRewriteService) Get(ctx context.Context, query SettingAutomaticHTTPSRewriteGetParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticHTTPSRewrites, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPSRewriteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAutomaticHTTPSRewritesJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewritesJSON contains the JSON metadata for the struct
// [ZoneSettingAutomaticHTTPSRewrites]
type zoneSettingAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingAutomaticHTTPSRewrites) implementsZonesSettingEditResponse() {}

func (r ZoneSettingAutomaticHTTPSRewrites) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingAutomaticHTTPSRewritesID string

const (
	ZoneSettingAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZoneSettingAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r ZoneSettingAutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case ZoneSettingAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingAutomaticHTTPSRewritesValue string

const (
	ZoneSettingAutomaticHTTPSRewritesValueOn  ZoneSettingAutomaticHTTPSRewritesValue = "on"
	ZoneSettingAutomaticHTTPSRewritesValueOff ZoneSettingAutomaticHTTPSRewritesValue = "off"
)

func (r ZoneSettingAutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case ZoneSettingAutomaticHTTPSRewritesValueOn, ZoneSettingAutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAutomaticHTTPSRewritesEditable bool

const (
	ZoneSettingAutomaticHTTPSRewritesEditableTrue  ZoneSettingAutomaticHTTPSRewritesEditable = true
	ZoneSettingAutomaticHTTPSRewritesEditableFalse ZoneSettingAutomaticHTTPSRewritesEditable = false
)

func (r ZoneSettingAutomaticHTTPSRewritesEditable) IsKnown() bool {
	switch r {
	case ZoneSettingAutomaticHTTPSRewritesEditableTrue, ZoneSettingAutomaticHTTPSRewritesEditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingAutomaticHTTPSRewritesParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingAutomaticHTTPSRewritesID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingAutomaticHTTPSRewritesValue] `json:"value,required"`
}

func (r ZoneSettingAutomaticHTTPSRewritesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingAutomaticHTTPSRewritesParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingAutomaticHTTPSRewriteEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingAutomaticHTTPSRewriteEditParamsValue] `json:"value,required"`
}

func (r SettingAutomaticHTTPSRewriteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingAutomaticHTTPSRewriteEditParamsValue string

const (
	SettingAutomaticHTTPSRewriteEditParamsValueOn  SettingAutomaticHTTPSRewriteEditParamsValue = "on"
	SettingAutomaticHTTPSRewriteEditParamsValueOff SettingAutomaticHTTPSRewriteEditParamsValue = "off"
)

func (r SettingAutomaticHTTPSRewriteEditParamsValue) IsKnown() bool {
	switch r {
	case SettingAutomaticHTTPSRewriteEditParamsValueOn, SettingAutomaticHTTPSRewriteEditParamsValueOff:
		return true
	}
	return false
}

type SettingAutomaticHTTPSRewriteEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result ZoneSettingAutomaticHTTPSRewrites                    `json:"result"`
	JSON   settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAutomaticHTTPSRewriteEditResponseEnvelope]
type settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAutomaticHTTPSRewriteGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAutomaticHTTPSRewriteGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result ZoneSettingAutomaticHTTPSRewrites                   `json:"result"`
	JSON   settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAutomaticHTTPSRewriteGetResponseEnvelope]
type settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
