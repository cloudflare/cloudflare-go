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
func (r *ZoneSettingSecurityHeaderService) Edit(ctx context.Context, params ZoneSettingSecurityHeaderEditParams, opts ...option.RequestOption) (res *ZoneSettingSecurityHeaderEditResponse, err error) {
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
func (r *ZoneSettingSecurityHeaderService) Get(ctx context.Context, query ZoneSettingSecurityHeaderGetParams, opts ...option.RequestOption) (res *ZoneSettingSecurityHeaderGetResponse, err error) {
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
type ZoneSettingSecurityHeaderEditResponse struct {
	// ID of the zone's security header.
	ID ZoneSettingSecurityHeaderEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSecurityHeaderEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityHeaderEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityHeaderEditResponseJSON `json:"-"`
}

// zoneSettingSecurityHeaderEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityHeaderEditResponse]
type zoneSettingSecurityHeaderEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityHeaderEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone's security header.
type ZoneSettingSecurityHeaderEditResponseID string

const (
	ZoneSettingSecurityHeaderEditResponseIDSecurityHeader ZoneSettingSecurityHeaderEditResponseID = "security_header"
)

// Current value of the zone setting.
type ZoneSettingSecurityHeaderEditResponseValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingSecurityHeaderEditResponseValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingSecurityHeaderEditResponseValueJSON                    `json:"-"`
}

// zoneSettingSecurityHeaderEditResponseValueJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderEditResponseValue]
type zoneSettingSecurityHeaderEditResponseValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityHeaderEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type ZoneSettingSecurityHeaderEditResponseValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                  `json:"nosniff"`
	JSON    zoneSettingSecurityHeaderEditResponseValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingSecurityHeaderEditResponseValueStrictTransportSecurityJSON contains
// the JSON metadata for the struct
// [ZoneSettingSecurityHeaderEditResponseValueStrictTransportSecurity]
type zoneSettingSecurityHeaderEditResponseValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderEditResponseValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityHeaderEditResponseValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityHeaderEditResponseEditable bool

const (
	ZoneSettingSecurityHeaderEditResponseEditableTrue  ZoneSettingSecurityHeaderEditResponseEditable = true
	ZoneSettingSecurityHeaderEditResponseEditableFalse ZoneSettingSecurityHeaderEditResponseEditable = false
)

// Cloudflare security header for a zone.
type ZoneSettingSecurityHeaderGetResponse struct {
	// ID of the zone's security header.
	ID ZoneSettingSecurityHeaderGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSecurityHeaderGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityHeaderGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityHeaderGetResponseJSON `json:"-"`
}

// zoneSettingSecurityHeaderGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityHeaderGetResponse]
type zoneSettingSecurityHeaderGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityHeaderGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone's security header.
type ZoneSettingSecurityHeaderGetResponseID string

const (
	ZoneSettingSecurityHeaderGetResponseIDSecurityHeader ZoneSettingSecurityHeaderGetResponseID = "security_header"
)

// Current value of the zone setting.
type ZoneSettingSecurityHeaderGetResponseValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingSecurityHeaderGetResponseValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingSecurityHeaderGetResponseValueJSON                    `json:"-"`
}

// zoneSettingSecurityHeaderGetResponseValueJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityHeaderGetResponseValue]
type zoneSettingSecurityHeaderGetResponseValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityHeaderGetResponseValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type ZoneSettingSecurityHeaderGetResponseValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                 `json:"nosniff"`
	JSON    zoneSettingSecurityHeaderGetResponseValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingSecurityHeaderGetResponseValueStrictTransportSecurityJSON contains
// the JSON metadata for the struct
// [ZoneSettingSecurityHeaderGetResponseValueStrictTransportSecurity]
type zoneSettingSecurityHeaderGetResponseValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderGetResponseValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSecurityHeaderGetResponseValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityHeaderGetResponseEditable bool

const (
	ZoneSettingSecurityHeaderGetResponseEditableTrue  ZoneSettingSecurityHeaderGetResponseEditable = true
	ZoneSettingSecurityHeaderGetResponseEditableFalse ZoneSettingSecurityHeaderGetResponseEditable = false
)

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
	Result ZoneSettingSecurityHeaderEditResponse             `json:"result"`
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

func (r zoneSettingSecurityHeaderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingSecurityHeaderEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingSecurityHeaderEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingSecurityHeaderGetResponse             `json:"result"`
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

func (r zoneSettingSecurityHeaderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingSecurityHeaderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingSecurityHeaderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
