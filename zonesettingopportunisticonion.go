// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ZoneSettingOpportunisticOnionService) Edit(ctx context.Context, params ZoneSettingOpportunisticOnionEditParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticOnionEditResponse, err error) {
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
func (r *ZoneSettingOpportunisticOnionService) Get(ctx context.Context, query ZoneSettingOpportunisticOnionGetParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticOnionGetResponse, err error) {
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
type ZoneSettingOpportunisticOnionEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticOnionEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingOpportunisticOnionEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticOnionEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOpportunisticOnionEditResponseJSON `json:"-"`
}

// zoneSettingOpportunisticOnionEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOpportunisticOnionEditResponse]
type zoneSettingOpportunisticOnionEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOpportunisticOnionEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingOpportunisticOnionEditResponseID string

const (
	ZoneSettingOpportunisticOnionEditResponseIDOpportunisticOnion ZoneSettingOpportunisticOnionEditResponseID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZoneSettingOpportunisticOnionEditResponseValue string

const (
	ZoneSettingOpportunisticOnionEditResponseValueOn  ZoneSettingOpportunisticOnionEditResponseValue = "on"
	ZoneSettingOpportunisticOnionEditResponseValueOff ZoneSettingOpportunisticOnionEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticOnionEditResponseEditable bool

const (
	ZoneSettingOpportunisticOnionEditResponseEditableTrue  ZoneSettingOpportunisticOnionEditResponseEditable = true
	ZoneSettingOpportunisticOnionEditResponseEditableFalse ZoneSettingOpportunisticOnionEditResponseEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingOpportunisticOnionGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticOnionGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingOpportunisticOnionGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticOnionGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOpportunisticOnionGetResponseJSON `json:"-"`
}

// zoneSettingOpportunisticOnionGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOpportunisticOnionGetResponse]
type zoneSettingOpportunisticOnionGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOpportunisticOnionGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingOpportunisticOnionGetResponseID string

const (
	ZoneSettingOpportunisticOnionGetResponseIDOpportunisticOnion ZoneSettingOpportunisticOnionGetResponseID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZoneSettingOpportunisticOnionGetResponseValue string

const (
	ZoneSettingOpportunisticOnionGetResponseValueOn  ZoneSettingOpportunisticOnionGetResponseValue = "on"
	ZoneSettingOpportunisticOnionGetResponseValueOff ZoneSettingOpportunisticOnionGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticOnionGetResponseEditable bool

const (
	ZoneSettingOpportunisticOnionGetResponseEditableTrue  ZoneSettingOpportunisticOnionGetResponseEditable = true
	ZoneSettingOpportunisticOnionGetResponseEditableFalse ZoneSettingOpportunisticOnionGetResponseEditable = false
)

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
	Result ZoneSettingOpportunisticOnionEditResponse             `json:"result"`
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

func (r zoneSettingOpportunisticOnionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingOpportunisticOnionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingOpportunisticOnionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingOpportunisticOnionGetResponse             `json:"result"`
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

func (r zoneSettingOpportunisticOnionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingOpportunisticOnionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingOpportunisticOnionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
