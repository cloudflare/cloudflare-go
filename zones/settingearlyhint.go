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

// SettingEarlyHintService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingEarlyHintService] method
// instead.
type SettingEarlyHintService struct {
	Options []option.RequestOption
}

// NewSettingEarlyHintService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingEarlyHintService(opts ...option.RequestOption) (r *SettingEarlyHintService) {
	r = &SettingEarlyHintService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *SettingEarlyHintService) Edit(ctx context.Context, params SettingEarlyHintEditParams, opts ...option.RequestOption) (res *ZoneSettingEarlyHints, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEarlyHintEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *SettingEarlyHintService) Get(ctx context.Context, query SettingEarlyHintGetParams, opts ...option.RequestOption) (res *ZoneSettingEarlyHints, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEarlyHintGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEarlyHintsJSON `json:"-"`
}

// zoneSettingEarlyHintsJSON contains the JSON metadata for the struct
// [ZoneSettingEarlyHints]
type zoneSettingEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingEarlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingEarlyHints) implementsZonesSettingEditResponse() {}

func (r ZoneSettingEarlyHints) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingEarlyHintsID string

const (
	ZoneSettingEarlyHintsIDEarlyHints ZoneSettingEarlyHintsID = "early_hints"
)

func (r ZoneSettingEarlyHintsID) IsKnown() bool {
	switch r {
	case ZoneSettingEarlyHintsIDEarlyHints:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingEarlyHintsValue string

const (
	ZoneSettingEarlyHintsValueOn  ZoneSettingEarlyHintsValue = "on"
	ZoneSettingEarlyHintsValueOff ZoneSettingEarlyHintsValue = "off"
)

func (r ZoneSettingEarlyHintsValue) IsKnown() bool {
	switch r {
	case ZoneSettingEarlyHintsValueOn, ZoneSettingEarlyHintsValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEarlyHintsEditable bool

const (
	ZoneSettingEarlyHintsEditableTrue  ZoneSettingEarlyHintsEditable = true
	ZoneSettingEarlyHintsEditableFalse ZoneSettingEarlyHintsEditable = false
)

func (r ZoneSettingEarlyHintsEditable) IsKnown() bool {
	switch r {
	case ZoneSettingEarlyHintsEditableTrue, ZoneSettingEarlyHintsEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEarlyHintsParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEarlyHintsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEarlyHintsValue] `json:"value,required"`
}

func (r ZoneSettingEarlyHintsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEarlyHintsParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingEarlyHintEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingEarlyHintEditParamsValue] `json:"value,required"`
}

func (r SettingEarlyHintEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingEarlyHintEditParamsValue string

const (
	SettingEarlyHintEditParamsValueOn  SettingEarlyHintEditParamsValue = "on"
	SettingEarlyHintEditParamsValueOff SettingEarlyHintEditParamsValue = "off"
)

func (r SettingEarlyHintEditParamsValue) IsKnown() bool {
	switch r {
	case SettingEarlyHintEditParamsValueOn, SettingEarlyHintEditParamsValueOff:
		return true
	}
	return false
}

type SettingEarlyHintEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZoneSettingEarlyHints                    `json:"result"`
	JSON   settingEarlyHintEditResponseEnvelopeJSON `json:"-"`
}

// settingEarlyHintEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingEarlyHintEditResponseEnvelope]
type settingEarlyHintEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingEarlyHintGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZoneSettingEarlyHints                   `json:"result"`
	JSON   settingEarlyHintGetResponseEnvelopeJSON `json:"-"`
}

// settingEarlyHintGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingEarlyHintGetResponseEnvelope]
type settingEarlyHintGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
