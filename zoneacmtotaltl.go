// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAcmTotalTlService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAcmTotalTlService] method
// instead.
type ZoneAcmTotalTlService struct {
	Options []option.RequestOption
}

// NewZoneAcmTotalTlService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAcmTotalTlService(opts ...option.RequestOption) (r *ZoneAcmTotalTlService) {
	r = &ZoneAcmTotalTlService{}
	r.Options = opts
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *ZoneAcmTotalTlService) TotalTlsEnableOrDisableTotalTls(ctx context.Context, zoneIdentifier string, body ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParams, opts ...option.RequestOption) (res *TotalTlsSettingsResponseRJujzw74, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Total TLS Settings for a Zone.
func (r *ZoneAcmTotalTlService) TotalTlsTotalTlsSettingsDetails(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *TotalTlsSettingsResponseRJujzw74, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TotalTlsSettingsResponseRJujzw74 struct {
	Errors   []TotalTlsSettingsResponseRJujzw74Error   `json:"errors"`
	Messages []TotalTlsSettingsResponseRJujzw74Message `json:"messages"`
	Result   TotalTlsSettingsResponseRJujzw74Result    `json:"result"`
	// Whether the API call was successful
	Success TotalTlsSettingsResponseRJujzw74Success `json:"success"`
	JSON    totalTlsSettingsResponseRJujzw74JSON    `json:"-"`
}

// totalTlsSettingsResponseRJujzw74JSON contains the JSON metadata for the struct
// [TotalTlsSettingsResponseRJujzw74]
type totalTlsSettingsResponseRJujzw74JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTlsSettingsResponseRJujzw74) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TotalTlsSettingsResponseRJujzw74Error struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    totalTlsSettingsResponseRJujzw74ErrorJSON `json:"-"`
}

// totalTlsSettingsResponseRJujzw74ErrorJSON contains the JSON metadata for the
// struct [TotalTlsSettingsResponseRJujzw74Error]
type totalTlsSettingsResponseRJujzw74ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTlsSettingsResponseRJujzw74Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TotalTlsSettingsResponseRJujzw74Message struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    totalTlsSettingsResponseRJujzw74MessageJSON `json:"-"`
}

// totalTlsSettingsResponseRJujzw74MessageJSON contains the JSON metadata for the
// struct [TotalTlsSettingsResponseRJujzw74Message]
type totalTlsSettingsResponseRJujzw74MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTlsSettingsResponseRJujzw74Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TotalTlsSettingsResponseRJujzw74Result struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority TotalTlsSettingsResponseRJujzw74ResultCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays TotalTlsSettingsResponseRJujzw74ResultValidityDays `json:"validity_days"`
	JSON         totalTlsSettingsResponseRJujzw74ResultJSON         `json:"-"`
}

// totalTlsSettingsResponseRJujzw74ResultJSON contains the JSON metadata for the
// struct [TotalTlsSettingsResponseRJujzw74Result]
type totalTlsSettingsResponseRJujzw74ResultJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTlsSettingsResponseRJujzw74Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type TotalTlsSettingsResponseRJujzw74ResultCertificateAuthority string

const (
	TotalTlsSettingsResponseRJujzw74ResultCertificateAuthorityGoogle      TotalTlsSettingsResponseRJujzw74ResultCertificateAuthority = "google"
	TotalTlsSettingsResponseRJujzw74ResultCertificateAuthorityLetsEncrypt TotalTlsSettingsResponseRJujzw74ResultCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type TotalTlsSettingsResponseRJujzw74ResultValidityDays int64

const (
	TotalTlsSettingsResponseRJujzw74ResultValidityDays90 TotalTlsSettingsResponseRJujzw74ResultValidityDays = 90
)

// Whether the API call was successful
type TotalTlsSettingsResponseRJujzw74Success bool

const (
	TotalTlsSettingsResponseRJujzw74SuccessTrue TotalTlsSettingsResponseRJujzw74Success = true
)

type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParams struct {
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParamsCertificateAuthority] `json:"certificate_authority"`
}

func (r ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParamsCertificateAuthority string

const (
	ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParamsCertificateAuthorityGoogle      ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParamsCertificateAuthority = "google"
	ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParamsCertificateAuthorityLetsEncrypt ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParamsCertificateAuthority = "lets_encrypt"
)
