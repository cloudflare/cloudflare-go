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

// ZoneSettingH2PrioritizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingH2PrioritizationService] method instead.
type ZoneSettingH2PrioritizationService struct {
	Options []option.RequestOption
}

// NewZoneSettingH2PrioritizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingH2PrioritizationService(opts ...option.RequestOption) (r *ZoneSettingH2PrioritizationService) {
	r = &ZoneSettingH2PrioritizationService{}
	r.Options = opts
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *ZoneSettingH2PrioritizationService) Edit(ctx context.Context, params ZoneSettingH2PrioritizationEditParams, opts ...option.RequestOption) (res *ZonesH2Prioritization, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingH2PrioritizationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *ZoneSettingH2PrioritizationService) Get(ctx context.Context, query ZoneSettingH2PrioritizationGetParams, opts ...option.RequestOption) (res *ZonesH2Prioritization, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingH2PrioritizationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesH2PrioritizationJSON `json:"-"`
}

// zonesH2PrioritizationJSON contains the JSON metadata for the struct
// [ZonesH2Prioritization]
type zonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesH2Prioritization) implementsZoneSettingEditResponse() {}

func (r ZonesH2Prioritization) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesH2PrioritizationID string

const (
	ZonesH2PrioritizationIDH2Prioritization ZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type ZonesH2PrioritizationValue string

const (
	ZonesH2PrioritizationValueOn     ZonesH2PrioritizationValue = "on"
	ZonesH2PrioritizationValueOff    ZonesH2PrioritizationValue = "off"
	ZonesH2PrioritizationValueCustom ZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesH2PrioritizationEditable bool

const (
	ZonesH2PrioritizationEditableTrue  ZonesH2PrioritizationEditable = true
	ZonesH2PrioritizationEditableFalse ZonesH2PrioritizationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZonesH2PrioritizationParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesH2PrioritizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesH2PrioritizationValue] `json:"value,required"`
}

func (r ZonesH2PrioritizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesH2PrioritizationParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingH2PrioritizationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[ZonesH2PrioritizationParam] `json:"value,required"`
}

func (r ZoneSettingH2PrioritizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingH2PrioritizationEditResponseEnvelope struct {
	Errors   []ZoneSettingH2PrioritizationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingH2PrioritizationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result ZonesH2Prioritization                               `json:"result"`
	JSON   zoneSettingH2PrioritizationEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingH2PrioritizationEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingH2PrioritizationEditResponseEnvelope]
type zoneSettingH2PrioritizationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationEditResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingH2PrioritizationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingH2PrioritizationEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingH2PrioritizationEditResponseEnvelopeErrors]
type zoneSettingH2PrioritizationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationEditResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingH2PrioritizationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingH2PrioritizationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingH2PrioritizationEditResponseEnvelopeMessages]
type zoneSettingH2PrioritizationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingH2PrioritizationGetResponseEnvelope struct {
	Errors   []ZoneSettingH2PrioritizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingH2PrioritizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result ZonesH2Prioritization                              `json:"result"`
	JSON   zoneSettingH2PrioritizationGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingH2PrioritizationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingH2PrioritizationGetResponseEnvelope]
type zoneSettingH2PrioritizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingH2PrioritizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingH2PrioritizationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingH2PrioritizationGetResponseEnvelopeErrors]
type zoneSettingH2PrioritizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingH2PrioritizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingH2PrioritizationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingH2PrioritizationGetResponseEnvelopeMessages]
type zoneSettingH2PrioritizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
