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

// SettingEarlyHintService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingEarlyHintService] method instead.
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
func (r *SettingEarlyHintService) Edit(ctx context.Context, params SettingEarlyHintEditParams, opts ...option.RequestOption) (res *EarlyHints, err error) {
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
func (r *SettingEarlyHintService) Get(ctx context.Context, query SettingEarlyHintGetParams, opts ...option.RequestOption) (res *EarlyHints, err error) {
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
type EarlyHints struct {
	// ID of the zone setting.
	ID EarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value EarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       earlyHintsJSON `json:"-"`
}

// earlyHintsJSON contains the JSON metadata for the struct [EarlyHints]
type earlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r earlyHintsJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type EarlyHintsID string

const (
	EarlyHintsIDEarlyHints EarlyHintsID = "early_hints"
)

func (r EarlyHintsID) IsKnown() bool {
	switch r {
	case EarlyHintsIDEarlyHints:
		return true
	}
	return false
}

// Current value of the zone setting.
type EarlyHintsValue string

const (
	EarlyHintsValueOn  EarlyHintsValue = "on"
	EarlyHintsValueOff EarlyHintsValue = "off"
)

func (r EarlyHintsValue) IsKnown() bool {
	switch r {
	case EarlyHintsValueOn, EarlyHintsValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type EarlyHintsEditable bool

const (
	EarlyHintsEditableTrue  EarlyHintsEditable = true
	EarlyHintsEditableFalse EarlyHintsEditable = false
)

func (r EarlyHintsEditable) IsKnown() bool {
	switch r {
	case EarlyHintsEditableTrue, EarlyHintsEditableFalse:
		return true
	}
	return false
}

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
	Result EarlyHints                               `json:"result"`
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
	Result EarlyHints                              `json:"result"`
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
