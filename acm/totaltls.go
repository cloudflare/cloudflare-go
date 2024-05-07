// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// TotalTLSService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTotalTLSService] method instead.
type TotalTLSService struct {
	Options []option.RequestOption
}

// NewTotalTLSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTotalTLSService(opts ...option.RequestOption) (r *TotalTLSService) {
	r = &TotalTLSService{}
	r.Options = opts
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *TotalTLSService) New(ctx context.Context, params TotalTLSNewParams, opts ...option.RequestOption) (res *TotalTLSNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TotalTLSNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/acm/total_tls", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Total TLS Settings for a Zone.
func (r *TotalTLSService) Get(ctx context.Context, query TotalTLSGetParams, opts ...option.RequestOption) (res *TotalTLSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TotalTLSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/acm/total_tls", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The Certificate Authority that Total TLS certificates will be issued through.
type CertificateAuthority string

const (
	CertificateAuthorityGoogle      CertificateAuthority = "google"
	CertificateAuthorityLetsEncrypt CertificateAuthority = "lets_encrypt"
)

func (r CertificateAuthority) IsKnown() bool {
	switch r {
	case CertificateAuthorityGoogle, CertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type TotalTLSNewResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays TotalTLSNewResponseValidityDays `json:"validity_days"`
	JSON         totalTLSNewResponseJSON         `json:"-"`
}

// totalTLSNewResponseJSON contains the JSON metadata for the struct
// [TotalTLSNewResponse]
type totalTLSNewResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTLSNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type TotalTLSNewResponseValidityDays int64

const (
	TotalTLSNewResponseValidityDays90 TotalTLSNewResponseValidityDays = 90
)

func (r TotalTLSNewResponseValidityDays) IsKnown() bool {
	switch r {
	case TotalTLSNewResponseValidityDays90:
		return true
	}
	return false
}

type TotalTLSGetResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityDays TotalTLSGetResponseValidityDays `json:"validity_days"`
	JSON         totalTLSGetResponseJSON         `json:"-"`
}

// totalTLSGetResponseJSON contains the JSON metadata for the struct
// [TotalTLSGetResponse]
type totalTLSGetResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type TotalTLSGetResponseValidityDays int64

const (
	TotalTLSGetResponseValidityDays90 TotalTLSGetResponseValidityDays = 90
)

func (r TotalTLSGetResponseValidityDays) IsKnown() bool {
	switch r {
	case TotalTLSGetResponseValidityDays90:
		return true
	}
	return false
}

type TotalTLSNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[CertificateAuthority] `json:"certificate_authority"`
}

func (r TotalTLSNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TotalTLSNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TotalTLSNewResponseEnvelopeSuccess `json:"success,required"`
	Result  TotalTLSNewResponse                `json:"result"`
	JSON    totalTLSNewResponseEnvelopeJSON    `json:"-"`
}

// totalTLSNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TotalTLSNewResponseEnvelope]
type totalTLSNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TotalTLSNewResponseEnvelopeSuccess bool

const (
	TotalTLSNewResponseEnvelopeSuccessTrue TotalTLSNewResponseEnvelopeSuccess = true
)

func (r TotalTLSNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TotalTLSNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TotalTLSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TotalTLSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TotalTLSGetResponseEnvelopeSuccess `json:"success,required"`
	Result  TotalTLSGetResponse                `json:"result"`
	JSON    totalTLSGetResponseEnvelopeJSON    `json:"-"`
}

// totalTLSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TotalTLSGetResponseEnvelope]
type totalTLSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TotalTLSGetResponseEnvelopeSuccess bool

const (
	TotalTLSGetResponseEnvelopeSuccessTrue TotalTLSGetResponseEnvelopeSuccess = true
)

func (r TotalTLSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TotalTLSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
