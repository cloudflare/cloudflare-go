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
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingSecurityHeaderService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingSecurityHeaderService] method instead.
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
func (r *SettingSecurityHeaderService) Edit(ctx context.Context, params SettingSecurityHeaderEditParams, opts ...option.RequestOption) (res *SecurityHeaders, err error) {
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
func (r *SettingSecurityHeaderService) Get(ctx context.Context, query SettingSecurityHeaderGetParams, opts ...option.RequestOption) (res *SecurityHeaders, err error) {
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
type SecurityHeaders struct {
	// ID of the zone's security header.
	ID SecurityHeadersID `json:"id,required"`
	// Current value of the zone setting.
	Value SecurityHeadersValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SecurityHeadersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       securityHeadersJSON `json:"-"`
}

// securityHeadersJSON contains the JSON metadata for the struct [SecurityHeaders]
type securityHeadersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityHeadersJSON) RawJSON() string {
	return r.raw
}

// ID of the zone's security header.
type SecurityHeadersID string

const (
	SecurityHeadersIDSecurityHeader SecurityHeadersID = "security_header"
)

func (r SecurityHeadersID) IsKnown() bool {
	switch r {
	case SecurityHeadersIDSecurityHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SecurityHeadersValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SecurityHeadersValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    securityHeadersValueJSON                    `json:"-"`
}

// securityHeadersValueJSON contains the JSON metadata for the struct
// [SecurityHeadersValue]
type securityHeadersValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SecurityHeadersValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityHeadersValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type SecurityHeadersValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                            `json:"nosniff"`
	JSON    securityHeadersValueStrictTransportSecurityJSON `json:"-"`
}

// securityHeadersValueStrictTransportSecurityJSON contains the JSON metadata for
// the struct [SecurityHeadersValueStrictTransportSecurity]
type securityHeadersValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SecurityHeadersValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityHeadersValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SecurityHeadersEditable bool

const (
	SecurityHeadersEditableTrue  SecurityHeadersEditable = true
	SecurityHeadersEditableFalse SecurityHeadersEditable = false
)

func (r SecurityHeadersEditable) IsKnown() bool {
	switch r {
	case SecurityHeadersEditableTrue, SecurityHeadersEditableFalse:
		return true
	}
	return false
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result SecurityHeaders                               `json:"result"`
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

type SettingSecurityHeaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSecurityHeaderGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result SecurityHeaders                              `json:"result"`
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
