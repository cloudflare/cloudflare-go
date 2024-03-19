// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificates

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ClientCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewClientCertificateService] method
// instead.
type ClientCertificateService struct {
	Options []option.RequestOption
}

// NewClientCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientCertificateService(opts ...option.RequestOption) (r *ClientCertificateService) {
	r = &ClientCertificateService{}
	r.Options = opts
	return
}

// Create a new API Shield mTLS Client Certificate
func (r *ClientCertificateService) New(ctx context.Context, params ClientCertificateNewParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ClientCertificateService) List(ctx context.Context, params ClientCertificateListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[TLSCertificatesAndHostnamesClientCertificate], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/client_certificates", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ClientCertificateService) ListAutoPaging(ctx context.Context, params ClientCertificateListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[TLSCertificatesAndHostnamesClientCertificate] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Set a API Shield mTLS Client Certificate to pending_revocation status for
// processing to revoked status.
func (r *ClientCertificateService) Delete(ctx context.Context, clientCertificateID string, body ClientCertificateDeleteParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", body.ZoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If a API Shield mTLS Client Certificate is in a pending_revocation state, you
// may reactivate it with this endpoint.
func (r *ClientCertificateService) Edit(ctx context.Context, clientCertificateID string, body ClientCertificateEditParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", body.ZoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Details for a single mTLS API Shield Client Certificate
func (r *ClientCertificateService) Get(ctx context.Context, clientCertificateID string, query ClientCertificateGetParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", query.ZoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TLSCertificatesAndHostnamesClientCertificate struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority TLSCertificatesAndHostnamesClientCertificateCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status TLSCertificatesAndHostnamesClientCertificateStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                            `json:"validity_days"`
	JSON         tlsCertificatesAndHostnamesClientCertificateJSON `json:"-"`
}

// tlsCertificatesAndHostnamesClientCertificateJSON contains the JSON metadata for
// the struct [TLSCertificatesAndHostnamesClientCertificate]
type tlsCertificatesAndHostnamesClientCertificateJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesClientCertificateJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type TLSCertificatesAndHostnamesClientCertificateCertificateAuthority struct {
	ID   string                                                               `json:"id"`
	Name string                                                               `json:"name"`
	JSON tlsCertificatesAndHostnamesClientCertificateCertificateAuthorityJSON `json:"-"`
}

// tlsCertificatesAndHostnamesClientCertificateCertificateAuthorityJSON contains
// the JSON metadata for the struct
// [TLSCertificatesAndHostnamesClientCertificateCertificateAuthority]
type tlsCertificatesAndHostnamesClientCertificateCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesClientCertificateCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesClientCertificateCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type TLSCertificatesAndHostnamesClientCertificateStatus string

const (
	TLSCertificatesAndHostnamesClientCertificateStatusActive              TLSCertificatesAndHostnamesClientCertificateStatus = "active"
	TLSCertificatesAndHostnamesClientCertificateStatusPendingReactivation TLSCertificatesAndHostnamesClientCertificateStatus = "pending_reactivation"
	TLSCertificatesAndHostnamesClientCertificateStatusPendingRevocation   TLSCertificatesAndHostnamesClientCertificateStatus = "pending_revocation"
	TLSCertificatesAndHostnamesClientCertificateStatusRevoked             TLSCertificatesAndHostnamesClientCertificateStatus = "revoked"
)

func (r TLSCertificatesAndHostnamesClientCertificateStatus) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesClientCertificateStatusActive, TLSCertificatesAndHostnamesClientCertificateStatusPendingReactivation, TLSCertificatesAndHostnamesClientCertificateStatusPendingRevocation, TLSCertificatesAndHostnamesClientCertificateStatusRevoked:
		return true
	}
	return false
}

type ClientCertificateNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr,required"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays param.Field[int64] `json:"validity_days,required"`
}

func (r ClientCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClientCertificateNewResponseEnvelope struct {
	Errors   []ClientCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesClientCertificate   `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateNewResponseEnvelope]
type clientCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    clientCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateNewResponseEnvelopeErrors]
type clientCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    clientCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ClientCertificateNewResponseEnvelopeMessages]
type clientCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ClientCertificateNewResponseEnvelopeSuccess bool

const (
	ClientCertificateNewResponseEnvelopeSuccessTrue ClientCertificateNewResponseEnvelopeSuccess = true
)

func (r ClientCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ClientCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ClientCertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Limit to the number of records returned.
	Limit param.Field[int64] `query:"limit"`
	// Offset the results
	Offset param.Field[int64] `query:"offset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Client Certitifcate Status to filter results by.
	Status param.Field[ClientCertificateListParamsStatus] `query:"status"`
}

// URLQuery serializes [ClientCertificateListParams]'s query parameters as
// `url.Values`.
func (r ClientCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Client Certitifcate Status to filter results by.
type ClientCertificateListParamsStatus string

const (
	ClientCertificateListParamsStatusAll                 ClientCertificateListParamsStatus = "all"
	ClientCertificateListParamsStatusActive              ClientCertificateListParamsStatus = "active"
	ClientCertificateListParamsStatusPendingReactivation ClientCertificateListParamsStatus = "pending_reactivation"
	ClientCertificateListParamsStatusPendingRevocation   ClientCertificateListParamsStatus = "pending_revocation"
	ClientCertificateListParamsStatusRevoked             ClientCertificateListParamsStatus = "revoked"
)

func (r ClientCertificateListParamsStatus) IsKnown() bool {
	switch r {
	case ClientCertificateListParamsStatusAll, ClientCertificateListParamsStatusActive, ClientCertificateListParamsStatusPendingReactivation, ClientCertificateListParamsStatusPendingRevocation, ClientCertificateListParamsStatusRevoked:
		return true
	}
	return false
}

type ClientCertificateDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ClientCertificateDeleteResponseEnvelope struct {
	Errors   []ClientCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesClientCertificate      `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateDeleteResponseEnvelope]
type clientCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    clientCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateDeleteResponseEnvelopeErrors]
type clientCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    clientCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ClientCertificateDeleteResponseEnvelopeMessages]
type clientCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ClientCertificateDeleteResponseEnvelopeSuccess bool

const (
	ClientCertificateDeleteResponseEnvelopeSuccessTrue ClientCertificateDeleteResponseEnvelopeSuccess = true
)

func (r ClientCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ClientCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ClientCertificateEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ClientCertificateEditResponseEnvelope struct {
	Errors   []ClientCertificateEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesClientCertificate    `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateEditResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateEditResponseEnvelope]
type clientCertificateEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    clientCertificateEditResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateEditResponseEnvelopeErrors]
type clientCertificateEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    clientCertificateEditResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ClientCertificateEditResponseEnvelopeMessages]
type clientCertificateEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ClientCertificateEditResponseEnvelopeSuccess bool

const (
	ClientCertificateEditResponseEnvelopeSuccessTrue ClientCertificateEditResponseEnvelopeSuccess = true
)

func (r ClientCertificateEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ClientCertificateEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ClientCertificateGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ClientCertificateGetResponseEnvelope struct {
	Errors   []ClientCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesClientCertificate   `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateGetResponseEnvelope]
type clientCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    clientCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateGetResponseEnvelopeErrors]
type clientCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ClientCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    clientCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ClientCertificateGetResponseEnvelopeMessages]
type clientCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ClientCertificateGetResponseEnvelopeSuccess bool

const (
	ClientCertificateGetResponseEnvelopeSuccessTrue ClientCertificateGetResponseEnvelopeSuccess = true
)

func (r ClientCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ClientCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
