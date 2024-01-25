// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneClientCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneClientCertificateService]
// method instead.
type ZoneClientCertificateService struct {
	Options []option.RequestOption
}

// NewZoneClientCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneClientCertificateService(opts ...option.RequestOption) (r *ZoneClientCertificateService) {
	r = &ZoneClientCertificateService{}
	r.Options = opts
	return
}

// Get Details for a single mTLS API Shield Client Certificate
func (r *ZoneClientCertificateService) Get(ctx context.Context, zoneIdentifier string, clientCertificateIdentifier string, opts ...option.RequestOption) (res *ZoneClientCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneIdentifier, clientCertificateIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// If a API Shield mTLS Client Certificate is in a pending_revocation state, you
// may reactivate it with this endpoint.
func (r *ZoneClientCertificateService) Update(ctx context.Context, zoneIdentifier string, clientCertificateIdentifier string, opts ...option.RequestOption) (res *ZoneClientCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneIdentifier, clientCertificateIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Set a API Shield mTLS Client Certificate to pending_revocation status for
// processing to revoked status.
func (r *ZoneClientCertificateService) Delete(ctx context.Context, zoneIdentifier string, clientCertificateIdentifier string, opts ...option.RequestOption) (res *ZoneClientCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneIdentifier, clientCertificateIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new API Shield mTLS Client Certificate
func (r *ZoneClientCertificateService) ClientCertificateForAZoneNewClientCertificate(ctx context.Context, zoneIdentifier string, body ZoneClientCertificateClientCertificateForAZoneNewClientCertificateParams, opts ...option.RequestOption) (res *ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ZoneClientCertificateService) ClientCertificateForAZoneListClientCertificates(ctx context.Context, zoneIdentifier string, query ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParams, opts ...option.RequestOption) (res *shared.Page[ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/client_certificates", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

type ZoneClientCertificateGetResponse struct {
	Errors   []ZoneClientCertificateGetResponseError   `json:"errors"`
	Messages []ZoneClientCertificateGetResponseMessage `json:"messages"`
	Result   ZoneClientCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneClientCertificateGetResponseSuccess `json:"success"`
	JSON    zoneClientCertificateGetResponseJSON    `json:"-"`
}

// zoneClientCertificateGetResponseJSON contains the JSON metadata for the struct
// [ZoneClientCertificateGetResponse]
type zoneClientCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneClientCertificateGetResponseErrorJSON `json:"-"`
}

// zoneClientCertificateGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneClientCertificateGetResponseError]
type zoneClientCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneClientCertificateGetResponseMessageJSON `json:"-"`
}

// zoneClientCertificateGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneClientCertificateGetResponseMessage]
type zoneClientCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ZoneClientCertificateGetResponseResultCertificateAuthority `json:"certificate_authority"`
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
	Status ZoneClientCertificateGetResponseResultStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                      `json:"validity_days"`
	JSON         zoneClientCertificateGetResponseResultJSON `json:"-"`
}

// zoneClientCertificateGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneClientCertificateGetResponseResult]
type zoneClientCertificateGetResponseResultJSON struct {
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

func (r *ZoneClientCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ZoneClientCertificateGetResponseResultCertificateAuthority struct {
	ID   string                                                         `json:"id"`
	Name string                                                         `json:"name"`
	JSON zoneClientCertificateGetResponseResultCertificateAuthorityJSON `json:"-"`
}

// zoneClientCertificateGetResponseResultCertificateAuthorityJSON contains the JSON
// metadata for the struct
// [ZoneClientCertificateGetResponseResultCertificateAuthority]
type zoneClientCertificateGetResponseResultCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateGetResponseResultCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ZoneClientCertificateGetResponseResultStatus string

const (
	ZoneClientCertificateGetResponseResultStatusActive              ZoneClientCertificateGetResponseResultStatus = "active"
	ZoneClientCertificateGetResponseResultStatusPendingReactivation ZoneClientCertificateGetResponseResultStatus = "pending_reactivation"
	ZoneClientCertificateGetResponseResultStatusPendingRevocation   ZoneClientCertificateGetResponseResultStatus = "pending_revocation"
	ZoneClientCertificateGetResponseResultStatusRevoked             ZoneClientCertificateGetResponseResultStatus = "revoked"
)

// Whether the API call was successful
type ZoneClientCertificateGetResponseSuccess bool

const (
	ZoneClientCertificateGetResponseSuccessTrue ZoneClientCertificateGetResponseSuccess = true
)

type ZoneClientCertificateUpdateResponse struct {
	Errors   []ZoneClientCertificateUpdateResponseError   `json:"errors"`
	Messages []ZoneClientCertificateUpdateResponseMessage `json:"messages"`
	Result   ZoneClientCertificateUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneClientCertificateUpdateResponseSuccess `json:"success"`
	JSON    zoneClientCertificateUpdateResponseJSON    `json:"-"`
}

// zoneClientCertificateUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneClientCertificateUpdateResponse]
type zoneClientCertificateUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneClientCertificateUpdateResponseErrorJSON `json:"-"`
}

// zoneClientCertificateUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneClientCertificateUpdateResponseError]
type zoneClientCertificateUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneClientCertificateUpdateResponseMessageJSON `json:"-"`
}

// zoneClientCertificateUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneClientCertificateUpdateResponseMessage]
type zoneClientCertificateUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateUpdateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ZoneClientCertificateUpdateResponseResultCertificateAuthority `json:"certificate_authority"`
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
	Status ZoneClientCertificateUpdateResponseResultStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                         `json:"validity_days"`
	JSON         zoneClientCertificateUpdateResponseResultJSON `json:"-"`
}

// zoneClientCertificateUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneClientCertificateUpdateResponseResult]
type zoneClientCertificateUpdateResponseResultJSON struct {
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

func (r *ZoneClientCertificateUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ZoneClientCertificateUpdateResponseResultCertificateAuthority struct {
	ID   string                                                            `json:"id"`
	Name string                                                            `json:"name"`
	JSON zoneClientCertificateUpdateResponseResultCertificateAuthorityJSON `json:"-"`
}

// zoneClientCertificateUpdateResponseResultCertificateAuthorityJSON contains the
// JSON metadata for the struct
// [ZoneClientCertificateUpdateResponseResultCertificateAuthority]
type zoneClientCertificateUpdateResponseResultCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateUpdateResponseResultCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ZoneClientCertificateUpdateResponseResultStatus string

const (
	ZoneClientCertificateUpdateResponseResultStatusActive              ZoneClientCertificateUpdateResponseResultStatus = "active"
	ZoneClientCertificateUpdateResponseResultStatusPendingReactivation ZoneClientCertificateUpdateResponseResultStatus = "pending_reactivation"
	ZoneClientCertificateUpdateResponseResultStatusPendingRevocation   ZoneClientCertificateUpdateResponseResultStatus = "pending_revocation"
	ZoneClientCertificateUpdateResponseResultStatusRevoked             ZoneClientCertificateUpdateResponseResultStatus = "revoked"
)

// Whether the API call was successful
type ZoneClientCertificateUpdateResponseSuccess bool

const (
	ZoneClientCertificateUpdateResponseSuccessTrue ZoneClientCertificateUpdateResponseSuccess = true
)

type ZoneClientCertificateDeleteResponse struct {
	Errors   []ZoneClientCertificateDeleteResponseError   `json:"errors"`
	Messages []ZoneClientCertificateDeleteResponseMessage `json:"messages"`
	Result   ZoneClientCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneClientCertificateDeleteResponseSuccess `json:"success"`
	JSON    zoneClientCertificateDeleteResponseJSON    `json:"-"`
}

// zoneClientCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneClientCertificateDeleteResponse]
type zoneClientCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneClientCertificateDeleteResponseErrorJSON `json:"-"`
}

// zoneClientCertificateDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneClientCertificateDeleteResponseError]
type zoneClientCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneClientCertificateDeleteResponseMessageJSON `json:"-"`
}

// zoneClientCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneClientCertificateDeleteResponseMessage]
type zoneClientCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateDeleteResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ZoneClientCertificateDeleteResponseResultCertificateAuthority `json:"certificate_authority"`
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
	Status ZoneClientCertificateDeleteResponseResultStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                         `json:"validity_days"`
	JSON         zoneClientCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneClientCertificateDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneClientCertificateDeleteResponseResult]
type zoneClientCertificateDeleteResponseResultJSON struct {
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

func (r *ZoneClientCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ZoneClientCertificateDeleteResponseResultCertificateAuthority struct {
	ID   string                                                            `json:"id"`
	Name string                                                            `json:"name"`
	JSON zoneClientCertificateDeleteResponseResultCertificateAuthorityJSON `json:"-"`
}

// zoneClientCertificateDeleteResponseResultCertificateAuthorityJSON contains the
// JSON metadata for the struct
// [ZoneClientCertificateDeleteResponseResultCertificateAuthority]
type zoneClientCertificateDeleteResponseResultCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateDeleteResponseResultCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ZoneClientCertificateDeleteResponseResultStatus string

const (
	ZoneClientCertificateDeleteResponseResultStatusActive              ZoneClientCertificateDeleteResponseResultStatus = "active"
	ZoneClientCertificateDeleteResponseResultStatusPendingReactivation ZoneClientCertificateDeleteResponseResultStatus = "pending_reactivation"
	ZoneClientCertificateDeleteResponseResultStatusPendingRevocation   ZoneClientCertificateDeleteResponseResultStatus = "pending_revocation"
	ZoneClientCertificateDeleteResponseResultStatusRevoked             ZoneClientCertificateDeleteResponseResultStatus = "revoked"
)

// Whether the API call was successful
type ZoneClientCertificateDeleteResponseSuccess bool

const (
	ZoneClientCertificateDeleteResponseSuccessTrue ZoneClientCertificateDeleteResponseSuccess = true
)

type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponse struct {
	Errors   []ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseError   `json:"errors"`
	Messages []ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessage `json:"messages"`
	Result   ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseSuccess `json:"success"`
	JSON    zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseJSON    `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponse]
type zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseErrorJSON `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseError]
type zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessageJSON `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessage]
type zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthority `json:"certificate_authority"`
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
	Status ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                                                                `json:"validity_days"`
	JSON         zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultJSON `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResult]
type zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultJSON struct {
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

func (r *ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthority struct {
	ID   string                                                                                                   `json:"id"`
	Name string                                                                                                   `json:"name"`
	JSON zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthorityJSON `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthorityJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthority]
type zoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatus string

const (
	ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatusActive              ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatus = "active"
	ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatusPendingReactivation ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatus = "pending_reactivation"
	ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatusPendingRevocation   ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatus = "pending_revocation"
	ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatusRevoked             ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseResultStatus = "revoked"
)

// Whether the API call was successful
type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseSuccess bool

const (
	ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseSuccessTrue ZoneClientCertificateClientCertificateForAZoneNewClientCertificateResponseSuccess = true
)

type ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthority `json:"certificate_authority"`
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
	Status ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                                                            `json:"validity_days"`
	JSON         zoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseJSON `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponse]
type zoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseJSON struct {
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

func (r *ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthority struct {
	ID   string                                                                                               `json:"id"`
	Name string                                                                                               `json:"name"`
	JSON zoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthorityJSON `json:"-"`
}

// zoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthorityJSON
// contains the JSON metadata for the struct
// [ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthority]
type zoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatus string

const (
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatusActive              ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatus = "active"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatusPendingReactivation ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatus = "pending_reactivation"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatusPendingRevocation   ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatus = "pending_revocation"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatusRevoked             ZoneClientCertificateClientCertificateForAZoneListClientCertificatesResponseStatus = "revoked"
)

type ZoneClientCertificateClientCertificateForAZoneNewClientCertificateParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr,required"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays param.Field[int64] `json:"validity_days,required"`
}

func (r ZoneClientCertificateClientCertificateForAZoneNewClientCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParams struct {
	// Limit to the number of records returned.
	Limit param.Field[int64] `query:"limit"`
	// Offset the results
	Offset param.Field[int64] `query:"offset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Client Certitifcate Status to filter results by.
	Status param.Field[ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus] `query:"status"`
}

// URLQuery serializes
// [ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParams]'s
// query parameters as `url.Values`.
func (r ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Client Certitifcate Status to filter results by.
type ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus string

const (
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatusAll                 ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus = "all"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatusActive              ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus = "active"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatusPendingReactivation ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus = "pending_reactivation"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatusPendingRevocation   ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus = "pending_revocation"
	ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatusRevoked             ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParamsStatus = "revoked"
)
