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
func (r *ZoneAcmTotalTlService) TotalTlsEnableOrDisableTotalTls(ctx context.Context, zoneIdentifier string, body ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsParams, opts ...option.RequestOption) (res *ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Total TLS Settings for a Zone.
func (r *ZoneAcmTotalTlService) TotalTlsTotalTlsSettingsDetails(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponse struct {
	Errors   []ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseError   `json:"errors"`
	Messages []ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessage `json:"messages"`
	Result   ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseSuccess `json:"success"`
	JSON    zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseJSON    `json:"-"`
}

// zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseJSON contains the JSON
// metadata for the struct [ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponse]
type zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseErrorJSON `json:"-"`
}

// zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseError]
type zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessageJSON `json:"-"`
}

// zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessage]
type zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResult struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultValidityDays `json:"validity_days"`
	JSON         zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultJSON         `json:"-"`
}

// zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResult]
type zoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultCertificateAuthority string

const (
	ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultCertificateAuthorityGoogle      ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultCertificateAuthority = "google"
	ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultCertificateAuthorityLetsEncrypt ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultValidityDays int64

const (
	ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultValidityDays90 ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseResultValidityDays = 90
)

// Whether the API call was successful
type ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseSuccess bool

const (
	ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseSuccessTrue ZoneAcmTotalTlTotalTlsEnableOrDisableTotalTlsResponseSuccess = true
)

type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponse struct {
	Errors   []ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseError   `json:"errors"`
	Messages []ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessage `json:"messages"`
	Result   ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseSuccess `json:"success"`
	JSON    zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseJSON    `json:"-"`
}

// zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseJSON contains the JSON
// metadata for the struct [ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponse]
type zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseErrorJSON `json:"-"`
}

// zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseError]
type zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessageJSON `json:"-"`
}

// zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessage]
type zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResult struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultValidityDays `json:"validity_days"`
	JSON         zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultJSON         `json:"-"`
}

// zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResult]
type zoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultCertificateAuthority string

const (
	ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultCertificateAuthorityGoogle      ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultCertificateAuthority = "google"
	ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultCertificateAuthorityLetsEncrypt ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultValidityDays int64

const (
	ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultValidityDays90 ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseResultValidityDays = 90
)

// Whether the API call was successful
type ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseSuccess bool

const (
	ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseSuccessTrue ZoneAcmTotalTlTotalTlsTotalTlsSettingsDetailsResponseSuccess = true
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
