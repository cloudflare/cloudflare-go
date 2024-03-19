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

// SettingSecurityHeaderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingSecurityHeaderService]
// method instead.
type SettingSecurityHeaderService struct {
	Options []option.RequestOption
}

// NewSettingSecurityHeaderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingSecurityHeaderService(opts ...option.RequestOption) (r *SettingSecurityHeaderService) {
	r = &SettingSecurityHeaderService{}
	r.Options = opts
	return
}

// Cloudflare security header for a zone.
func (r *SettingSecurityHeaderService) Edit(ctx context.Context, params SettingSecurityHeaderEditParams, opts ...option.RequestOption) (res *ZonesSecurityHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityHeaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_header", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare security header for a zone.
func (r *SettingSecurityHeaderService) Get(ctx context.Context, query SettingSecurityHeaderGetParams, opts ...option.RequestOption) (res *ZonesSecurityHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_header", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r zonesSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSecurityHeader) implementsZonesSettingEditResponse() {}

func (r ZonesSecurityHeader) implementsZonesSettingGetResponse() {}

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

func (r zonesSecurityHeaderValueJSON) RawJSON() string {
	return r.raw
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

func (r zonesSecurityHeaderValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
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

func (r ZonesSecurityHeaderParam) implementsZonesSettingEditParamsItem() {}

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

type SettingSecurityHeaderEditParams struct {
	// Identifier
	ZoneID param.Field[string]                               `path:"zone_id,required"`
	Value  param.Field[SettingSecurityHeaderEditParamsValue] `json:"value,required"`
}

func (r SettingSecurityHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSecurityHeaderEditParamsValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[SettingSecurityHeaderEditParamsValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r SettingSecurityHeaderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type SettingSecurityHeaderEditParamsValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r SettingSecurityHeaderEditParamsValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSecurityHeaderEditResponseEnvelope struct {
	Errors   []SettingSecurityHeaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityHeaderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result ZonesSecurityHeader                           `json:"result"`
	JSON   settingSecurityHeaderEditResponseEnvelopeJSON `json:"-"`
}

// settingSecurityHeaderEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityHeaderEditResponseEnvelope]
type settingSecurityHeaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityHeaderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityHeaderEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingSecurityHeaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityHeaderEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderEditResponseEnvelopeErrors]
type settingSecurityHeaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityHeaderEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityHeaderEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingSecurityHeaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityHeaderEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderEditResponseEnvelopeMessages]
type settingSecurityHeaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityHeaderEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityHeaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSecurityHeaderGetResponseEnvelope struct {
	Errors   []SettingSecurityHeaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityHeaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result ZonesSecurityHeader                          `json:"result"`
	JSON   settingSecurityHeaderGetResponseEnvelopeJSON `json:"-"`
}

// settingSecurityHeaderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityHeaderGetResponseEnvelope]
type settingSecurityHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityHeaderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityHeaderGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingSecurityHeaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityHeaderGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderGetResponseEnvelopeErrors]
type settingSecurityHeaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityHeaderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityHeaderGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingSecurityHeaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityHeaderGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderGetResponseEnvelopeMessages]
type settingSecurityHeaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityHeaderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
