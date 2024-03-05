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

// ZoneSettingTrueClientIPHeaderService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingTrueClientIPHeaderService] method instead.
type ZoneSettingTrueClientIPHeaderService struct {
	Options []option.RequestOption
}

// NewZoneSettingTrueClientIPHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingTrueClientIPHeaderService(opts ...option.RequestOption) (r *ZoneSettingTrueClientIPHeaderService) {
	r = &ZoneSettingTrueClientIPHeaderService{}
	r.Options = opts
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *ZoneSettingTrueClientIPHeaderService) Edit(ctx context.Context, params ZoneSettingTrueClientIPHeaderEditParams, opts ...option.RequestOption) (res *ZonesTrueClientIPHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingTrueClientIPHeaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *ZoneSettingTrueClientIPHeaderService) Get(ctx context.Context, query ZoneSettingTrueClientIPHeaderGetParams, opts ...option.RequestOption) (res *ZonesTrueClientIPHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingTrueClientIPHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesTrueClientIPHeaderJSON `json:"-"`
}

// zonesTrueClientIPHeaderJSON contains the JSON metadata for the struct
// [ZonesTrueClientIPHeader]
type zonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesTrueClientIPHeader) implementsZoneSettingEditResponse() {}

func (r ZonesTrueClientIPHeader) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesTrueClientIPHeaderID string

const (
	ZonesTrueClientIPHeaderIDTrueClientIPHeader ZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type ZonesTrueClientIPHeaderValue string

const (
	ZonesTrueClientIPHeaderValueOn  ZonesTrueClientIPHeaderValue = "on"
	ZonesTrueClientIPHeaderValueOff ZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesTrueClientIPHeaderEditable bool

const (
	ZonesTrueClientIPHeaderEditableTrue  ZonesTrueClientIPHeaderEditable = true
	ZonesTrueClientIPHeaderEditableFalse ZonesTrueClientIPHeaderEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZonesTrueClientIPHeaderParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesTrueClientIPHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesTrueClientIPHeaderValue] `json:"value,required"`
}

func (r ZonesTrueClientIPHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTrueClientIPHeaderParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingTrueClientIPHeaderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingTrueClientIPHeaderEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingTrueClientIPHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingTrueClientIPHeaderEditParamsValue string

const (
	ZoneSettingTrueClientIPHeaderEditParamsValueOn  ZoneSettingTrueClientIPHeaderEditParamsValue = "on"
	ZoneSettingTrueClientIPHeaderEditParamsValueOff ZoneSettingTrueClientIPHeaderEditParamsValue = "off"
)

type ZoneSettingTrueClientIPHeaderEditResponseEnvelope struct {
	Errors   []ZoneSettingTrueClientIPHeaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingTrueClientIPHeaderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result ZonesTrueClientIPHeader                               `json:"result"`
	JSON   zoneSettingTrueClientIPHeaderEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderEditResponseEnvelope]
type zoneSettingTrueClientIPHeaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderEditResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZoneSettingTrueClientIPHeaderEditResponseEnvelopeErrors]
type zoneSettingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderEditResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingTrueClientIPHeaderEditResponseEnvelopeMessages]
type zoneSettingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingTrueClientIPHeaderGetResponseEnvelope struct {
	Errors   []ZoneSettingTrueClientIPHeaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingTrueClientIPHeaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result ZonesTrueClientIPHeader                              `json:"result"`
	JSON   zoneSettingTrueClientIPHeaderGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderGetResponseEnvelope]
type zoneSettingTrueClientIPHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingTrueClientIPHeaderGetResponseEnvelopeErrors]
type zoneSettingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingTrueClientIPHeaderGetResponseEnvelopeMessages]
type zoneSettingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
