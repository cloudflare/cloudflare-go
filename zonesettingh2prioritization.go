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
func (r *ZoneSettingH2PrioritizationService) Edit(ctx context.Context, params ZoneSettingH2PrioritizationEditParams, opts ...option.RequestOption) (res *ZoneSettingH2PrioritizationEditResponse, err error) {
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
func (r *ZoneSettingH2PrioritizationService) Get(ctx context.Context, query ZoneSettingH2PrioritizationGetParams, opts ...option.RequestOption) (res *ZoneSettingH2PrioritizationGetResponse, err error) {
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
type ZoneSettingH2PrioritizationEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingH2PrioritizationEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingH2PrioritizationEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingH2PrioritizationEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingH2PrioritizationEditResponseJSON `json:"-"`
}

// zoneSettingH2PrioritizationEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingH2PrioritizationEditResponse]
type zoneSettingH2PrioritizationEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingH2PrioritizationEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingH2PrioritizationEditResponseID string

const (
	ZoneSettingH2PrioritizationEditResponseIDH2Prioritization ZoneSettingH2PrioritizationEditResponseID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingH2PrioritizationEditResponseValue string

const (
	ZoneSettingH2PrioritizationEditResponseValueOn     ZoneSettingH2PrioritizationEditResponseValue = "on"
	ZoneSettingH2PrioritizationEditResponseValueOff    ZoneSettingH2PrioritizationEditResponseValue = "off"
	ZoneSettingH2PrioritizationEditResponseValueCustom ZoneSettingH2PrioritizationEditResponseValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationEditResponseEditable bool

const (
	ZoneSettingH2PrioritizationEditResponseEditableTrue  ZoneSettingH2PrioritizationEditResponseEditable = true
	ZoneSettingH2PrioritizationEditResponseEditableFalse ZoneSettingH2PrioritizationEditResponseEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingH2PrioritizationGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingH2PrioritizationGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingH2PrioritizationGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingH2PrioritizationGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingH2PrioritizationGetResponseJSON `json:"-"`
}

// zoneSettingH2PrioritizationGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingH2PrioritizationGetResponse]
type zoneSettingH2PrioritizationGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingH2PrioritizationGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingH2PrioritizationGetResponseID string

const (
	ZoneSettingH2PrioritizationGetResponseIDH2Prioritization ZoneSettingH2PrioritizationGetResponseID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingH2PrioritizationGetResponseValue string

const (
	ZoneSettingH2PrioritizationGetResponseValueOn     ZoneSettingH2PrioritizationGetResponseValue = "on"
	ZoneSettingH2PrioritizationGetResponseValueOff    ZoneSettingH2PrioritizationGetResponseValue = "off"
	ZoneSettingH2PrioritizationGetResponseValueCustom ZoneSettingH2PrioritizationGetResponseValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationGetResponseEditable bool

const (
	ZoneSettingH2PrioritizationGetResponseEditableTrue  ZoneSettingH2PrioritizationGetResponseEditable = true
	ZoneSettingH2PrioritizationGetResponseEditableFalse ZoneSettingH2PrioritizationGetResponseEditable = false
)

type ZoneSettingH2PrioritizationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[ZoneSettingH2PrioritizationEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingH2PrioritizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingH2PrioritizationEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingH2PrioritizationEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingH2PrioritizationEditParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingH2PrioritizationEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingH2PrioritizationEditParamsValueID string

const (
	ZoneSettingH2PrioritizationEditParamsValueIDH2Prioritization ZoneSettingH2PrioritizationEditParamsValueID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingH2PrioritizationEditParamsValueValue string

const (
	ZoneSettingH2PrioritizationEditParamsValueValueOn     ZoneSettingH2PrioritizationEditParamsValueValue = "on"
	ZoneSettingH2PrioritizationEditParamsValueValueOff    ZoneSettingH2PrioritizationEditParamsValueValue = "off"
	ZoneSettingH2PrioritizationEditParamsValueValueCustom ZoneSettingH2PrioritizationEditParamsValueValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationEditParamsValueEditable bool

const (
	ZoneSettingH2PrioritizationEditParamsValueEditableTrue  ZoneSettingH2PrioritizationEditParamsValueEditable = true
	ZoneSettingH2PrioritizationEditParamsValueEditableFalse ZoneSettingH2PrioritizationEditParamsValueEditable = false
)

type ZoneSettingH2PrioritizationEditResponseEnvelope struct {
	Errors   []ZoneSettingH2PrioritizationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingH2PrioritizationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result ZoneSettingH2PrioritizationEditResponse             `json:"result"`
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

func (r zoneSettingH2PrioritizationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingH2PrioritizationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingH2PrioritizationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingH2PrioritizationGetResponse             `json:"result"`
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

func (r zoneSettingH2PrioritizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingH2PrioritizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingH2PrioritizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
