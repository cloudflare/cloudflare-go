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

// ACMTotalTLSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewACMTotalTLSService] method
// instead.
type ACMTotalTLSService struct {
	Options []option.RequestOption
}

// NewACMTotalTLSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewACMTotalTLSService(opts ...option.RequestOption) (r *ACMTotalTLSService) {
	r = &ACMTotalTLSService{}
	r.Options = opts
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *ACMTotalTLSService) New(ctx context.Context, zoneID string, body ACMTotalTLSNewParams, opts ...option.RequestOption) (res *ACMTotalTLSNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ACMTotalTLSNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Total TLS Settings for a Zone.
func (r *ACMTotalTLSService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ACMTotalTLSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ACMTotalTLSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ACMTotalTLSNewResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority ACMTotalTLSNewResponseCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays ACMTotalTLSNewResponseValidityDays `json:"validity_days"`
	JSON         acmTotalTLSNewResponseJSON         `json:"-"`
}

// acmTotalTLSNewResponseJSON contains the JSON metadata for the struct
// [ACMTotalTLSNewResponse]
type acmTotalTLSNewResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ACMTotalTLSNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type ACMTotalTLSNewResponseCertificateAuthority string

const (
	ACMTotalTLSNewResponseCertificateAuthorityGoogle      ACMTotalTLSNewResponseCertificateAuthority = "google"
	ACMTotalTLSNewResponseCertificateAuthorityLetsEncrypt ACMTotalTLSNewResponseCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type ACMTotalTLSNewResponseValidityDays int64

const (
	ACMTotalTLSNewResponseValidityDays90 ACMTotalTLSNewResponseValidityDays = 90
)

type ACMTotalTLSGetResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority ACMTotalTLSGetResponseCertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays ACMTotalTLSGetResponseValidityDays `json:"validity_days"`
	JSON         acmTotalTLSGetResponseJSON         `json:"-"`
}

// acmTotalTLSGetResponseJSON contains the JSON metadata for the struct
// [ACMTotalTLSGetResponse]
type acmTotalTLSGetResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ACMTotalTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type ACMTotalTLSGetResponseCertificateAuthority string

const (
	ACMTotalTLSGetResponseCertificateAuthorityGoogle      ACMTotalTLSGetResponseCertificateAuthority = "google"
	ACMTotalTLSGetResponseCertificateAuthorityLetsEncrypt ACMTotalTLSGetResponseCertificateAuthority = "lets_encrypt"
)

// The validity period in days for the certificates ordered via Total TLS.
type ACMTotalTLSGetResponseValidityDays int64

const (
	ACMTotalTLSGetResponseValidityDays90 ACMTotalTLSGetResponseValidityDays = 90
)

type ACMTotalTLSNewParams struct {
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[ACMTotalTLSNewParamsCertificateAuthority] `json:"certificate_authority"`
}

func (r ACMTotalTLSNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Certificate Authority that Total TLS certificates will be issued through.
type ACMTotalTLSNewParamsCertificateAuthority string

const (
	ACMTotalTLSNewParamsCertificateAuthorityGoogle      ACMTotalTLSNewParamsCertificateAuthority = "google"
	ACMTotalTLSNewParamsCertificateAuthorityLetsEncrypt ACMTotalTLSNewParamsCertificateAuthority = "lets_encrypt"
)

type ACMTotalTLSNewResponseEnvelope struct {
	Errors   []ACMTotalTLSNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ACMTotalTLSNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ACMTotalTLSNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ACMTotalTLSNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    acmTotalTLSNewResponseEnvelopeJSON    `json:"-"`
}

// acmTotalTLSNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACMTotalTLSNewResponseEnvelope]
type acmTotalTLSNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACMTotalTLSNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACMTotalTLSNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    acmTotalTLSNewResponseEnvelopeErrorsJSON `json:"-"`
}

// acmTotalTLSNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ACMTotalTLSNewResponseEnvelopeErrors]
type acmTotalTLSNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACMTotalTLSNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACMTotalTLSNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    acmTotalTLSNewResponseEnvelopeMessagesJSON `json:"-"`
}

// acmTotalTLSNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ACMTotalTLSNewResponseEnvelopeMessages]
type acmTotalTLSNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACMTotalTLSNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ACMTotalTLSNewResponseEnvelopeSuccess bool

const (
	ACMTotalTLSNewResponseEnvelopeSuccessTrue ACMTotalTLSNewResponseEnvelopeSuccess = true
)

type ACMTotalTLSGetResponseEnvelope struct {
	Errors   []ACMTotalTLSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ACMTotalTLSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ACMTotalTLSGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ACMTotalTLSGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    acmTotalTLSGetResponseEnvelopeJSON    `json:"-"`
}

// acmTotalTLSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACMTotalTLSGetResponseEnvelope]
type acmTotalTLSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACMTotalTLSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACMTotalTLSGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    acmTotalTLSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// acmTotalTLSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ACMTotalTLSGetResponseEnvelopeErrors]
type acmTotalTLSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACMTotalTLSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACMTotalTLSGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    acmTotalTLSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// acmTotalTLSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ACMTotalTLSGetResponseEnvelopeMessages]
type acmTotalTLSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACMTotalTLSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ACMTotalTLSGetResponseEnvelopeSuccess bool

const (
	ACMTotalTLSGetResponseEnvelopeSuccessTrue ACMTotalTLSGetResponseEnvelopeSuccess = true
)
