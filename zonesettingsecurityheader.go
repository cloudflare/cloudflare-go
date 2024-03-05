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

// ZoneSettingSecurityHeaderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingSecurityHeaderService] method instead.
type ZoneSettingSecurityHeaderService struct {
	Options []option.RequestOption
}

// NewZoneSettingSecurityHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSecurityHeaderService(opts ...option.RequestOption) (r *ZoneSettingSecurityHeaderService) {
	r = &ZoneSettingSecurityHeaderService{}
	r.Options = opts
	return
}

// Cloudflare security header for a zone.
func (r *ZoneSettingSecurityHeaderService) Edit(ctx context.Context, params ZoneSettingSecurityHeaderEditParams, opts ...option.RequestOption) (res *ZonesSecurityHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSecurityHeaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_header", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare security header for a zone.
func (r *ZoneSettingSecurityHeaderService) Get(ctx context.Context, query ZoneSettingSecurityHeaderGetParams, opts ...option.RequestOption) (res *ZonesSecurityHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSecurityHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_header", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare security header for a zone.
type ZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID ZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time               `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesSecurityHeaderJSON `json:"-"`
}

// zonesSecurityHeaderJSON contains the JSON metadata for the struct
// [ZonesSecurityHeader]
type zonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesSecurityHeader) implementsZoneSettingEditResponse() {}

func (r ZonesSecurityHeader) implementsZoneSettingGetResponse() {}

// ID of the zone's security header.
type ZonesSecurityHeaderID string

const (
	ZonesSecurityHeaderIDSecurityHeader ZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type ZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zonesSecurityHeaderValueJSON                    `json:"-"`
}

// zonesSecurityHeaderValueJSON contains the JSON metadata for the struct
// [ZonesSecurityHeaderValue]
type zonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                `json:"nosniff"`
	JSON    zonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zonesSecurityHeaderValueStrictTransportSecurityJSON contains the JSON metadata
// for the struct [ZonesSecurityHeaderValueStrictTransportSecurity]
type zonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesSecurityHeaderEditable bool

const (
	ZonesSecurityHeaderEditableTrue  ZonesSecurityHeaderEditable = true
	ZonesSecurityHeaderEditableFalse ZonesSecurityHeaderEditable = false
)

// Cloudflare security header for a zone.
type ZonesSecurityHeaderParam struct {
	// ID of the zone's security header.
	ID param.Field[ZonesSecurityHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesSecurityHeaderValueParam] `json:"value,required"`
}

func (r ZonesSecurityHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSecurityHeaderParam) implementsZoneSettingEditParamsItem() {}

// Current value of the zone setting.
type ZonesSecurityHeaderValueParam struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[ZonesSecurityHeaderValueStrictTransportSecurityParam] `json:"strict_transport_security"`
}

func (r ZonesSecurityHeaderValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type ZonesSecurityHeaderValueStrictTransportSecurityParam struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r ZonesSecurityHeaderValueStrictTransportSecurityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingSecurityHeaderEditParams struct {
	// Identifier
	ZoneID param.Field[string]                                   `path:"zone_id,required"`
	Value  param.Field[ZoneSettingSecurityHeaderEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingSecurityHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingSecurityHeaderEditParamsValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[ZoneSettingSecurityHeaderEditParamsValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r ZoneSettingSecurityHeaderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type ZoneSettingSecurityHeaderEditParamsValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r ZoneSettingSecurityHeaderEditParamsValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingSecurityHeaderEditResponseEnvelope struct {
	Errors   []ZoneSettingSecurityHeaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSecurityHeaderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result ZonesSecurityHeader                               `json:"result"`
	JSON   zoneSettingSecurityHeaderEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSecurityHeaderEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderEditResponseEnvelope]
type zoneSettingSecurityHeaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingSecurityHeaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSecurityHeaderEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityHeaderEditResponseEnvelopeErrors]
type zoneSettingSecurityHeaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingSecurityHeaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSecurityHeaderEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityHeaderEditResponseEnvelopeMessages]
type zoneSettingSecurityHeaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingSecurityHeaderGetResponseEnvelope struct {
	Errors   []ZoneSettingSecurityHeaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSecurityHeaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result ZonesSecurityHeader                              `json:"result"`
	JSON   zoneSettingSecurityHeaderGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSecurityHeaderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderGetResponseEnvelope]
type zoneSettingSecurityHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingSecurityHeaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSecurityHeaderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityHeaderGetResponseEnvelopeErrors]
type zoneSettingSecurityHeaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingSecurityHeaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSecurityHeaderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityHeaderGetResponseEnvelopeMessages]
type zoneSettingSecurityHeaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
