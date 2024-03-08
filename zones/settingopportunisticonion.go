// File generated from our OpenAPI spec by Stainless.

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
func (r *SettingOpportunisticOnionService) Edit(ctx context.Context, params SettingOpportunisticOnionEditParams, opts ...option.RequestOption) (res *ZonesOpportunisticOnion, err error) {
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
func (r *SettingOpportunisticOnionService) Get(ctx context.Context, query SettingOpportunisticOnionGetParams, opts ...option.RequestOption) (res *ZonesOpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesOpportunisticOnionJSON `json:"-"`
}

// zonesOpportunisticOnionJSON contains the JSON metadata for the struct
// [ZonesOpportunisticOnion]
type zonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesOpportunisticOnionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesOpportunisticOnion) implementsZonesSettingEditResponse() {}

func (r ZonesOpportunisticOnion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesOpportunisticOnionID string

const (
	ZonesOpportunisticOnionIDOpportunisticOnion ZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZonesOpportunisticOnionValue string

const (
	ZonesOpportunisticOnionValueOn  ZonesOpportunisticOnionValue = "on"
	ZonesOpportunisticOnionValueOff ZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesOpportunisticOnionEditable bool

const (
	ZonesOpportunisticOnionEditableTrue  ZonesOpportunisticOnionEditable = true
	ZonesOpportunisticOnionEditableFalse ZonesOpportunisticOnionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZonesOpportunisticOnionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOpportunisticOnionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesOpportunisticOnionValue] `json:"value,required"`
}

func (r ZonesOpportunisticOnionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOpportunisticOnionParam) implementsZonesSettingEditParamsItem() {}

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

type SettingOpportunisticOnionEditResponseEnvelope struct {
	Errors   []SettingOpportunisticOnionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticOnionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZonesOpportunisticOnion                           `json:"result"`
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

type SettingOpportunisticOnionEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingOpportunisticOnionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticOnionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionEditResponseEnvelopeErrors]
type settingOpportunisticOnionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticOnionEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingOpportunisticOnionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticOnionEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionEditResponseEnvelopeMessages]
type settingOpportunisticOnionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticOnionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOpportunisticOnionGetResponseEnvelope struct {
	Errors   []SettingOpportunisticOnionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticOnionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZonesOpportunisticOnion                          `json:"result"`
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

type SettingOpportunisticOnionGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingOpportunisticOnionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionGetResponseEnvelopeErrors]
type settingOpportunisticOnionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticOnionGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingOpportunisticOnionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionGetResponseEnvelopeMessages]
type settingOpportunisticOnionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
