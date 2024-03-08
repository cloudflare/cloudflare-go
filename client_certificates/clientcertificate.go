// File generated from our OpenAPI spec by Stainless.

package client_certificates

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ClientCertificateService) New(ctx context.Context, params ClientCertificateNewParams, opts ...option.RequestOption) (res *ClientCertificateNewResponse, err error) {
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
func (r *ClientCertificateService) List(ctx context.Context, params ClientCertificateListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[ClientCertificateListResponse], err error) {
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
func (r *ClientCertificateService) ListAutoPaging(ctx context.Context, params ClientCertificateListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[ClientCertificateListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Set a API Shield mTLS Client Certificate to pending_revocation status for
// processing to revoked status.
func (r *ClientCertificateService) Delete(ctx context.Context, clientCertificateID string, body ClientCertificateDeleteParams, opts ...option.RequestOption) (res *ClientCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", body.ZoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If a API Shield mTLS Client Certificate is in a pending_revocation state, you
// may reactivate it with this endpoint.
func (r *ClientCertificateService) Edit(ctx context.Context, clientCertificateID string, body ClientCertificateEditParams, opts ...option.RequestOption) (res *ClientCertificateEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", body.ZoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Details for a single mTLS API Shield Client Certificate
func (r *ClientCertificateService) Get(ctx context.Context, clientCertificateID string, query ClientCertificateGetParams, opts ...option.RequestOption) (res *ClientCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", query.ZoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ClientCertificateNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateNewResponseCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateNewResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                            `json:"validity_days"`
	JSON         clientCertificateNewResponseJSON `json:"-"`
}

// clientCertificateNewResponseJSON contains the JSON metadata for the struct
// [ClientCertificateNewResponse]
type clientCertificateNewResponseJSON struct {
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

func (r *ClientCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateNewResponseCertificateAuthority struct {
	ID   string                                               `json:"id"`
	Name string                                               `json:"name"`
	JSON clientCertificateNewResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateNewResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateNewResponseCertificateAuthority]
type clientCertificateNewResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateNewResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateNewResponseStatus string

const (
	ClientCertificateNewResponseStatusActive              ClientCertificateNewResponseStatus = "active"
	ClientCertificateNewResponseStatusPendingReactivation ClientCertificateNewResponseStatus = "pending_reactivation"
	ClientCertificateNewResponseStatusPendingRevocation   ClientCertificateNewResponseStatus = "pending_revocation"
	ClientCertificateNewResponseStatusRevoked             ClientCertificateNewResponseStatus = "revoked"
)

type ClientCertificateListResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateListResponseCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateListResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                             `json:"validity_days"`
	JSON         clientCertificateListResponseJSON `json:"-"`
}

// clientCertificateListResponseJSON contains the JSON metadata for the struct
// [ClientCertificateListResponse]
type clientCertificateListResponseJSON struct {
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

func (r *ClientCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateListResponseCertificateAuthority struct {
	ID   string                                                `json:"id"`
	Name string                                                `json:"name"`
	JSON clientCertificateListResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateListResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateListResponseCertificateAuthority]
type clientCertificateListResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateListResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateListResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateListResponseStatus string

const (
	ClientCertificateListResponseStatusActive              ClientCertificateListResponseStatus = "active"
	ClientCertificateListResponseStatusPendingReactivation ClientCertificateListResponseStatus = "pending_reactivation"
	ClientCertificateListResponseStatusPendingRevocation   ClientCertificateListResponseStatus = "pending_revocation"
	ClientCertificateListResponseStatusRevoked             ClientCertificateListResponseStatus = "revoked"
)

type ClientCertificateDeleteResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateDeleteResponseCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateDeleteResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                               `json:"validity_days"`
	JSON         clientCertificateDeleteResponseJSON `json:"-"`
}

// clientCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [ClientCertificateDeleteResponse]
type clientCertificateDeleteResponseJSON struct {
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

func (r *ClientCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateDeleteResponseCertificateAuthority struct {
	ID   string                                                  `json:"id"`
	Name string                                                  `json:"name"`
	JSON clientCertificateDeleteResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateDeleteResponseCertificateAuthorityJSON contains the JSON
// metadata for the struct [ClientCertificateDeleteResponseCertificateAuthority]
type clientCertificateDeleteResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateDeleteResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateDeleteResponseStatus string

const (
	ClientCertificateDeleteResponseStatusActive              ClientCertificateDeleteResponseStatus = "active"
	ClientCertificateDeleteResponseStatusPendingReactivation ClientCertificateDeleteResponseStatus = "pending_reactivation"
	ClientCertificateDeleteResponseStatusPendingRevocation   ClientCertificateDeleteResponseStatus = "pending_revocation"
	ClientCertificateDeleteResponseStatusRevoked             ClientCertificateDeleteResponseStatus = "revoked"
)

type ClientCertificateEditResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateEditResponseCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateEditResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                             `json:"validity_days"`
	JSON         clientCertificateEditResponseJSON `json:"-"`
}

// clientCertificateEditResponseJSON contains the JSON metadata for the struct
// [ClientCertificateEditResponse]
type clientCertificateEditResponseJSON struct {
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

func (r *ClientCertificateEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateEditResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateEditResponseCertificateAuthority struct {
	ID   string                                                `json:"id"`
	Name string                                                `json:"name"`
	JSON clientCertificateEditResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateEditResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateEditResponseCertificateAuthority]
type clientCertificateEditResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateEditResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateEditResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateEditResponseStatus string

const (
	ClientCertificateEditResponseStatusActive              ClientCertificateEditResponseStatus = "active"
	ClientCertificateEditResponseStatusPendingReactivation ClientCertificateEditResponseStatus = "pending_reactivation"
	ClientCertificateEditResponseStatusPendingRevocation   ClientCertificateEditResponseStatus = "pending_revocation"
	ClientCertificateEditResponseStatusRevoked             ClientCertificateEditResponseStatus = "revoked"
)

type ClientCertificateGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateGetResponseCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateGetResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                            `json:"validity_days"`
	JSON         clientCertificateGetResponseJSON `json:"-"`
}

// clientCertificateGetResponseJSON contains the JSON metadata for the struct
// [ClientCertificateGetResponse]
type clientCertificateGetResponseJSON struct {
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

func (r *ClientCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateGetResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateGetResponseCertificateAuthority struct {
	ID   string                                               `json:"id"`
	Name string                                               `json:"name"`
	JSON clientCertificateGetResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateGetResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateGetResponseCertificateAuthority]
type clientCertificateGetResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateGetResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateGetResponseStatus string

const (
	ClientCertificateGetResponseStatusActive              ClientCertificateGetResponseStatus = "active"
	ClientCertificateGetResponseStatusPendingReactivation ClientCertificateGetResponseStatus = "pending_reactivation"
	ClientCertificateGetResponseStatusPendingRevocation   ClientCertificateGetResponseStatus = "pending_revocation"
	ClientCertificateGetResponseStatusRevoked             ClientCertificateGetResponseStatus = "revoked"
)

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
	Result   ClientCertificateNewResponse                   `json:"result,required"`
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

type ClientCertificateDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ClientCertificateDeleteResponseEnvelope struct {
	Errors   []ClientCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateDeleteResponse                   `json:"result,required"`
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

type ClientCertificateEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ClientCertificateEditResponseEnvelope struct {
	Errors   []ClientCertificateEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateEditResponse                   `json:"result,required"`
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

type ClientCertificateGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ClientCertificateGetResponseEnvelope struct {
	Errors   []ClientCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateGetResponse                   `json:"result,required"`
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
