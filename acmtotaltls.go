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

// AcmTotalTLSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAcmTotalTLSService] method
// instead.
type AcmTotalTLSService struct {
	Options []option.RequestOption
}

// NewAcmTotalTLSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAcmTotalTLSService(opts ...option.RequestOption) (r *AcmTotalTLSService) {
	r = &AcmTotalTLSService{}
	r.Options = opts
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *AcmTotalTLSService) Update(ctx context.Context, zoneID string, body AcmTotalTLSUpdateParams, opts ...option.RequestOption) (res *AcmTotalTLSUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AcmTotalTLSUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Total TLS Settings for a Zone.
func (r *AcmTotalTLSService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *AcmTotalTLSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AcmTotalTLSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AcmTotalTLSUpdateResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority AcmTotalTLSUpdateResponseCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays AcmTotalTLSUpdateResponseValidityDays `json:"validity_days"`
	JSON         acmTotalTLSUpdateResponseJSON         `json:"-"`
}

// acmTotalTLSUpdateResponseJSON contains the JSON metadata for the struct
// [AcmTotalTLSUpdateResponse]
type acmTotalTLSUpdateResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AcmTotalTLSUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type AcmTotalTLSUpdateResponseCertificateAuthority string

const (
	AcmTotalTLSUpdateResponseCertificateAuthorityGoogle      AcmTotalTLSUpdateResponseCertificateAuthority = "google"
	AcmTotalTLSUpdateResponseCertificateAuthorityLetsEncrypt AcmTotalTLSUpdateResponseCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type AcmTotalTLSUpdateResponseValidityDays int64

const (
	AcmTotalTLSUpdateResponseValidityDays90 AcmTotalTLSUpdateResponseValidityDays = 90
)

type AcmTotalTLSGetResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority AcmTotalTLSGetResponseCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays AcmTotalTLSGetResponseValidityDays `json:"validity_days"`
	JSON         acmTotalTLSGetResponseJSON         `json:"-"`
}

// acmTotalTLSGetResponseJSON contains the JSON metadata for the struct
// [AcmTotalTLSGetResponse]
type acmTotalTLSGetResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AcmTotalTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type AcmTotalTLSGetResponseCertificateAuthority string

const (
	AcmTotalTLSGetResponseCertificateAuthorityGoogle      AcmTotalTLSGetResponseCertificateAuthority = "google"
	AcmTotalTLSGetResponseCertificateAuthorityLetsEncrypt AcmTotalTLSGetResponseCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type AcmTotalTLSGetResponseValidityDays int64

const (
	AcmTotalTLSGetResponseValidityDays90 AcmTotalTLSGetResponseValidityDays = 90
)

type AcmTotalTLSUpdateParams struct {
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[AcmTotalTLSUpdateParamsCertificateAuthority] `json:"certificate_authority"`
}

func (r AcmTotalTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type AcmTotalTLSUpdateParamsCertificateAuthority string

const (
	AcmTotalTLSUpdateParamsCertificateAuthorityGoogle      AcmTotalTLSUpdateParamsCertificateAuthority = "google"
	AcmTotalTLSUpdateParamsCertificateAuthorityLetsEncrypt AcmTotalTLSUpdateParamsCertificateAuthority = "lets_encrypt"
)

type AcmTotalTLSUpdateResponseEnvelope struct {
	Errors   []AcmTotalTLSUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AcmTotalTLSUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AcmTotalTLSUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AcmTotalTLSUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    acmTotalTLSUpdateResponseEnvelopeJSON    `json:"-"`
}

// acmTotalTLSUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AcmTotalTLSUpdateResponseEnvelope]
type acmTotalTLSUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcmTotalTLSUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcmTotalTLSUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    acmTotalTLSUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// acmTotalTLSUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AcmTotalTLSUpdateResponseEnvelopeErrors]
type acmTotalTLSUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcmTotalTLSUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcmTotalTLSUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    acmTotalTLSUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// acmTotalTLSUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AcmTotalTLSUpdateResponseEnvelopeMessages]
type acmTotalTLSUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcmTotalTLSUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AcmTotalTLSUpdateResponseEnvelopeSuccess bool

const (
	AcmTotalTLSUpdateResponseEnvelopeSuccessTrue AcmTotalTLSUpdateResponseEnvelopeSuccess = true
)

type AcmTotalTLSGetResponseEnvelope struct {
	Errors   []AcmTotalTLSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AcmTotalTLSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AcmTotalTLSGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AcmTotalTLSGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    acmTotalTLSGetResponseEnvelopeJSON    `json:"-"`
}

// acmTotalTLSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AcmTotalTLSGetResponseEnvelope]
type acmTotalTLSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcmTotalTLSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcmTotalTLSGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    acmTotalTLSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// acmTotalTLSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AcmTotalTLSGetResponseEnvelopeErrors]
type acmTotalTLSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcmTotalTLSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AcmTotalTLSGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    acmTotalTLSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// acmTotalTLSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AcmTotalTLSGetResponseEnvelopeMessages]
type acmTotalTLSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AcmTotalTLSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AcmTotalTLSGetResponseEnvelopeSuccess bool

const (
	AcmTotalTLSGetResponseEnvelopeSuccessTrue AcmTotalTLSGetResponseEnvelopeSuccess = true
)
