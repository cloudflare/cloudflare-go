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
func (r *SettingEarlyHintService) Edit(ctx context.Context, params SettingEarlyHintEditParams, opts ...option.RequestOption) (res *ZonesEarlyHints, err error) {
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
func (r *SettingEarlyHintService) Get(ctx context.Context, query SettingEarlyHintGetParams, opts ...option.RequestOption) (res *ZonesEarlyHints, err error) {
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
type ZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesEarlyHintsJSON `json:"-"`
}

// zonesEarlyHintsJSON contains the JSON metadata for the struct [ZonesEarlyHints]
type zonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesEarlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r ZonesEarlyHints) implementsZonesSettingEditResponse() {}

func (r ZonesEarlyHints) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesEarlyHintsID string

const (
	ZonesEarlyHintsIDEarlyHints ZonesEarlyHintsID = "early_hints"
)

func (r ZonesEarlyHintsID) IsKnown() bool {
	switch r {
	case ZonesEarlyHintsIDEarlyHints:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesEarlyHintsValue string

const (
	ZonesEarlyHintsValueOn  ZonesEarlyHintsValue = "on"
	ZonesEarlyHintsValueOff ZonesEarlyHintsValue = "off"
)

func (r ZonesEarlyHintsValue) IsKnown() bool {
	switch r {
	case ZonesEarlyHintsValueOn, ZonesEarlyHintsValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesEarlyHintsEditable bool

const (
	ZonesEarlyHintsEditableTrue  ZonesEarlyHintsEditable = true
	ZonesEarlyHintsEditableFalse ZonesEarlyHintsEditable = false
)

func (r ZonesEarlyHintsEditable) IsKnown() bool {
	switch r {
	case ZonesEarlyHintsEditableTrue, ZonesEarlyHintsEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZonesEarlyHintsParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesEarlyHintsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesEarlyHintsValue] `json:"value,required"`
}

func (r ZonesEarlyHintsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesEarlyHintsParam) implementsZonesSettingEditParamsItem() {}

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
	Errors   []SettingEarlyHintEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEarlyHintEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZonesEarlyHints                          `json:"result"`
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

type SettingEarlyHintEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingEarlyHintEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEarlyHintEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingEarlyHintEditResponseEnvelopeErrors]
type settingEarlyHintEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingEarlyHintEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEarlyHintEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingEarlyHintEditResponseEnvelopeMessages]
type settingEarlyHintEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingEarlyHintGetResponseEnvelope struct {
	Errors   []SettingEarlyHintGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEarlyHintGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZonesEarlyHints                         `json:"result"`
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

type SettingEarlyHintGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingEarlyHintGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEarlyHintGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingEarlyHintGetResponseEnvelopeErrors]
type settingEarlyHintGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingEarlyHintGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEarlyHintGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingEarlyHintGetResponseEnvelopeMessages]
type settingEarlyHintGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
