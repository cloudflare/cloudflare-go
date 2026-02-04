// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// TotalTLSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTotalTLSService] method instead.
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
func (r *TotalTLSService) Update(ctx context.Context, params TotalTLSUpdateParams, opts ...option.RequestOption) (res *TotalTLSUpdateResponse, err error) {
	var env TotalTLSUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/total_tls", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *TotalTLSService) Edit(ctx context.Context, params TotalTLSEditParams, opts ...option.RequestOption) (res *TotalTLSEditResponse, err error) {
	var env TotalTLSEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
	var env TotalTLSGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
	CertificateAuthoritySSLCom      CertificateAuthority = "ssl_com"
)

func (r CertificateAuthority) IsKnown() bool {
	switch r {
	case CertificateAuthorityGoogle, CertificateAuthorityLetsEncrypt, CertificateAuthoritySSLCom:
		return true
	}
	return false
}

type TotalTLSUpdateResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityPeriod TotalTLSUpdateResponseValidityPeriod `json:"validity_period"`
	JSON           totalTLSUpdateResponseJSON           `json:"-"`
}

// totalTLSUpdateResponseJSON contains the JSON metadata for the struct
// [TotalTLSUpdateResponse]
type totalTLSUpdateResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityPeriod       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTLSUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type TotalTLSUpdateResponseValidityPeriod int64

const (
	TotalTLSUpdateResponseValidityPeriod90 TotalTLSUpdateResponseValidityPeriod = 90
)

func (r TotalTLSUpdateResponseValidityPeriod) IsKnown() bool {
	switch r {
	case TotalTLSUpdateResponseValidityPeriod90:
		return true
	}
	return false
}

type TotalTLSEditResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityPeriod TotalTLSEditResponseValidityPeriod `json:"validity_period"`
	JSON           totalTLSEditResponseJSON           `json:"-"`
}

// totalTLSEditResponseJSON contains the JSON metadata for the struct
// [TotalTLSEditResponse]
type totalTLSEditResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityPeriod       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTLSEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSEditResponseJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type TotalTLSEditResponseValidityPeriod int64

const (
	TotalTLSEditResponseValidityPeriod90 TotalTLSEditResponseValidityPeriod = 90
)

func (r TotalTLSEditResponseValidityPeriod) IsKnown() bool {
	switch r {
	case TotalTLSEditResponseValidityPeriod90:
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
	ValidityPeriod TotalTLSGetResponseValidityPeriod `json:"validity_period"`
	JSON           totalTLSGetResponseJSON           `json:"-"`
}

// totalTLSGetResponseJSON contains the JSON metadata for the struct
// [TotalTLSGetResponse]
type totalTLSGetResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityPeriod       apijson.Field
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
type TotalTLSGetResponseValidityPeriod int64

const (
	TotalTLSGetResponseValidityPeriod90 TotalTLSGetResponseValidityPeriod = 90
)

func (r TotalTLSGetResponseValidityPeriod) IsKnown() bool {
	switch r {
	case TotalTLSGetResponseValidityPeriod90:
		return true
	}
	return false
}

type TotalTLSUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[CertificateAuthority] `json:"certificate_authority"`
}

func (r TotalTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TotalTLSUpdateResponseEnvelope struct {
	Errors   []TotalTLSUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TotalTLSUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TotalTLSUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  TotalTLSUpdateResponse                `json:"result"`
	JSON    totalTLSUpdateResponseEnvelopeJSON    `json:"-"`
}

// totalTLSUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TotalTLSUpdateResponseEnvelope]
type totalTLSUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TotalTLSUpdateResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           TotalTLSUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             totalTLSUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// totalTLSUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TotalTLSUpdateResponseEnvelopeErrors]
type totalTLSUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TotalTLSUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    totalTLSUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// totalTLSUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [TotalTLSUpdateResponseEnvelopeErrorsSource]
type totalTLSUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TotalTLSUpdateResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           TotalTLSUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             totalTLSUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// totalTLSUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TotalTLSUpdateResponseEnvelopeMessages]
type totalTLSUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TotalTLSUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    totalTLSUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// totalTLSUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [TotalTLSUpdateResponseEnvelopeMessagesSource]
type totalTLSUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TotalTLSUpdateResponseEnvelopeSuccess bool

const (
	TotalTLSUpdateResponseEnvelopeSuccessTrue TotalTLSUpdateResponseEnvelopeSuccess = true
)

func (r TotalTLSUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TotalTLSUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TotalTLSEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[CertificateAuthority] `json:"certificate_authority"`
}

func (r TotalTLSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TotalTLSEditResponseEnvelope struct {
	Errors   []TotalTLSEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TotalTLSEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TotalTLSEditResponseEnvelopeSuccess `json:"success,required"`
	Result  TotalTLSEditResponse                `json:"result"`
	JSON    totalTLSEditResponseEnvelopeJSON    `json:"-"`
}

// totalTLSEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [TotalTLSEditResponseEnvelope]
type totalTLSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TotalTLSEditResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           TotalTLSEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             totalTLSEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// totalTLSEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TotalTLSEditResponseEnvelopeErrors]
type totalTLSEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TotalTLSEditResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    totalTLSEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// totalTLSEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TotalTLSEditResponseEnvelopeErrorsSource]
type totalTLSEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TotalTLSEditResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           TotalTLSEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             totalTLSEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// totalTLSEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TotalTLSEditResponseEnvelopeMessages]
type totalTLSEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TotalTLSEditResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    totalTLSEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// totalTLSEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [TotalTLSEditResponseEnvelopeMessagesSource]
type totalTLSEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TotalTLSEditResponseEnvelopeSuccess bool

const (
	TotalTLSEditResponseEnvelopeSuccessTrue TotalTLSEditResponseEnvelopeSuccess = true
)

func (r TotalTLSEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TotalTLSEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TotalTLSGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TotalTLSGetResponseEnvelope struct {
	Errors   []TotalTLSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TotalTLSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
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

type TotalTLSGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           TotalTLSGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             totalTLSGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// totalTLSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TotalTLSGetResponseEnvelopeErrors]
type totalTLSGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    totalTLSGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// totalTLSGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TotalTLSGetResponseEnvelopeErrorsSource]
type totalTLSGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           TotalTLSGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             totalTLSGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// totalTLSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TotalTLSGetResponseEnvelopeMessages]
type totalTLSGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    totalTLSGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// totalTLSGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TotalTLSGetResponseEnvelopeMessagesSource]
type totalTLSGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
