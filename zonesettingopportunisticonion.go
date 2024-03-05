// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingOpportunisticOnionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOpportunisticOnionService] method instead.
type ZoneSettingOpportunisticOnionService struct {
	Options []option.RequestOption
}

// NewZoneSettingOpportunisticOnionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOpportunisticOnionService(opts ...option.RequestOption) (r *ZoneSettingOpportunisticOnionService) {
	r = &ZoneSettingOpportunisticOnionService{}
	r.Options = opts
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *ZoneSettingOpportunisticOnionService) Edit(ctx context.Context, params ZoneSettingOpportunisticOnionEditParams, opts ...option.RequestOption) (res *ZonesOpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOpportunisticOnionEditResponseEnvelope
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
func (r *ZoneSettingOpportunisticOnionService) Get(ctx context.Context, query ZoneSettingOpportunisticOnionGetParams, opts ...option.RequestOption) (res *ZonesOpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOpportunisticOnionGetResponseEnvelope
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

func (r ZonesOpportunisticOnion) implementsZoneSettingEditResponse() {}

func (r ZonesOpportunisticOnion) implementsZoneSettingGetResponse() {}

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

func (r ZonesOpportunisticOnionParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingOpportunisticOnionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingOpportunisticOnionEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingOpportunisticOnionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticOnionEditParamsValue string

const (
	ZoneSettingOpportunisticOnionEditParamsValueOn  ZoneSettingOpportunisticOnionEditParamsValue = "on"
	ZoneSettingOpportunisticOnionEditParamsValueOff ZoneSettingOpportunisticOnionEditParamsValue = "off"
)

type ZoneSettingOpportunisticOnionEditResponseEnvelope struct {
	Errors   []ZoneSettingOpportunisticOnionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOpportunisticOnionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZonesOpportunisticOnion                               `json:"result"`
	JSON   zoneSettingOpportunisticOnionEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOpportunisticOnionEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionEditResponseEnvelope]
type zoneSettingOpportunisticOnionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionEditResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOpportunisticOnionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZoneSettingOpportunisticOnionEditResponseEnvelopeErrors]
type zoneSettingOpportunisticOnionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionEditResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOpportunisticOnionEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingOpportunisticOnionEditResponseEnvelopeMessages]
type zoneSettingOpportunisticOnionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingOpportunisticOnionGetResponseEnvelope struct {
	Errors   []ZoneSettingOpportunisticOnionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOpportunisticOnionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZonesOpportunisticOnion                              `json:"result"`
	JSON   zoneSettingOpportunisticOnionGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOpportunisticOnionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionGetResponseEnvelope]
type zoneSettingOpportunisticOnionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOpportunisticOnionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticOnionGetResponseEnvelopeErrors]
type zoneSettingOpportunisticOnionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOpportunisticOnionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingOpportunisticOnionGetResponseEnvelopeMessages]
type zoneSettingOpportunisticOnionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
