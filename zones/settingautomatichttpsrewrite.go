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
func (r *SettingAutomaticHTTPSRewriteService) Edit(ctx context.Context, params SettingAutomaticHTTPSRewriteEditParams, opts ...option.RequestOption) (res *AutomaticHTTPSRewrites, err error) {
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
func (r *SettingAutomaticHTTPSRewriteService) Get(ctx context.Context, query SettingAutomaticHTTPSRewriteGetParams, opts ...option.RequestOption) (res *AutomaticHTTPSRewrites, err error) {
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
type AutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID AutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value AutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       automaticHTTPSRewritesJSON `json:"-"`
}

// automaticHTTPSRewritesJSON contains the JSON metadata for the struct
// [AutomaticHTTPSRewrites]
type automaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type AutomaticHTTPSRewritesID string

const (
	AutomaticHTTPSRewritesIDAutomaticHTTPSRewrites AutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r AutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case AutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type AutomaticHTTPSRewritesValue string

const (
	AutomaticHTTPSRewritesValueOn  AutomaticHTTPSRewritesValue = "on"
	AutomaticHTTPSRewritesValueOff AutomaticHTTPSRewritesValue = "off"
)

func (r AutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case AutomaticHTTPSRewritesValueOn, AutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AutomaticHTTPSRewritesEditable bool

const (
	AutomaticHTTPSRewritesEditableTrue  AutomaticHTTPSRewritesEditable = true
	AutomaticHTTPSRewritesEditableFalse AutomaticHTTPSRewritesEditable = false
)

func (r AutomaticHTTPSRewritesEditable) IsKnown() bool {
	switch r {
	case AutomaticHTTPSRewritesEditableTrue, AutomaticHTTPSRewritesEditableFalse:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result AutomaticHTTPSRewrites                               `json:"result"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result AutomaticHTTPSRewrites                              `json:"result"`
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
