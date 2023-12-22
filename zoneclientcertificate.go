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
func (r *ZoneClientCertificateService) Get(ctx context.Context, zoneIdentifier string, clientCertificateIdentifier string, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneIdentifier, clientCertificateIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// If a API Shield mTLS Client Certificate is in a pending_revocation state, you
// may reactivate it with this endpoint.
func (r *ZoneClientCertificateService) Update(ctx context.Context, zoneIdentifier string, clientCertificateIdentifier string, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneIdentifier, clientCertificateIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Set a API Shield mTLS Client Certificate to pending_revocation status for
// processing to revoked status.
func (r *ZoneClientCertificateService) Delete(ctx context.Context, zoneIdentifier string, clientCertificateIdentifier string, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneIdentifier, clientCertificateIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new API Shield mTLS Client Certificate
func (r *ZoneClientCertificateService) ClientCertificateForAZoneNewClientCertificate(ctx context.Context, zoneIdentifier string, body ZoneClientCertificateClientCertificateForAZoneNewClientCertificateParams, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ZoneClientCertificateService) ClientCertificateForAZoneListClientCertificates(ctx context.Context, zoneIdentifier string, query ZoneClientCertificateClientCertificateForAZoneListClientCertificatesParams, opts ...option.RequestOption) (res *ClientCertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/client_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ClientCertificateResponseCollection struct {
	Errors     []ClientCertificateResponseCollectionError    `json:"errors"`
	Messages   []ClientCertificateResponseCollectionMessage  `json:"messages"`
	Result     []ClientCertificateResponseCollectionResult   `json:"result"`
	ResultInfo ClientCertificateResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ClientCertificateResponseCollectionSuccess `json:"success"`
	JSON    clientCertificateResponseCollectionJSON    `json:"-"`
}

// clientCertificateResponseCollectionJSON contains the JSON metadata for the
// struct [ClientCertificateResponseCollection]
type clientCertificateResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateResponseCollectionError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    clientCertificateResponseCollectionErrorJSON `json:"-"`
}

// clientCertificateResponseCollectionErrorJSON contains the JSON metadata for the
// struct [ClientCertificateResponseCollectionError]
type clientCertificateResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateResponseCollectionMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    clientCertificateResponseCollectionMessageJSON `json:"-"`
}

// clientCertificateResponseCollectionMessageJSON contains the JSON metadata for
// the struct [ClientCertificateResponseCollectionMessage]
type clientCertificateResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateResponseCollectionResult struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateResponseCollectionResultCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateResponseCollectionResultStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                         `json:"validity_days"`
	JSON         clientCertificateResponseCollectionResultJSON `json:"-"`
}

// clientCertificateResponseCollectionResultJSON contains the JSON metadata for the
// struct [ClientCertificateResponseCollectionResult]
type clientCertificateResponseCollectionResultJSON struct {
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

func (r *ClientCertificateResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateResponseCollectionResultCertificateAuthority struct {
	ID   string                                                            `json:"id"`
	Name string                                                            `json:"name"`
	JSON clientCertificateResponseCollectionResultCertificateAuthorityJSON `json:"-"`
}

// clientCertificateResponseCollectionResultCertificateAuthorityJSON contains the
// JSON metadata for the struct
// [ClientCertificateResponseCollectionResultCertificateAuthority]
type clientCertificateResponseCollectionResultCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseCollectionResultCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateResponseCollectionResultStatus string

const (
	ClientCertificateResponseCollectionResultStatusActive              ClientCertificateResponseCollectionResultStatus = "active"
	ClientCertificateResponseCollectionResultStatusPendingReactivation ClientCertificateResponseCollectionResultStatus = "pending_reactivation"
	ClientCertificateResponseCollectionResultStatusPendingRevocation   ClientCertificateResponseCollectionResultStatus = "pending_revocation"
	ClientCertificateResponseCollectionResultStatusRevoked             ClientCertificateResponseCollectionResultStatus = "revoked"
)

type ClientCertificateResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       clientCertificateResponseCollectionResultInfoJSON `json:"-"`
}

// clientCertificateResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [ClientCertificateResponseCollectionResultInfo]
type clientCertificateResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ClientCertificateResponseCollectionSuccess bool

const (
	ClientCertificateResponseCollectionSuccessTrue ClientCertificateResponseCollectionSuccess = true
)

type ClientCertificateResponseSingle struct {
	Errors   []ClientCertificateResponseSingleError   `json:"errors"`
	Messages []ClientCertificateResponseSingleMessage `json:"messages"`
	Result   ClientCertificateResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success ClientCertificateResponseSingleSuccess `json:"success"`
	JSON    clientCertificateResponseSingleJSON    `json:"-"`
}

// clientCertificateResponseSingleJSON contains the JSON metadata for the struct
// [ClientCertificateResponseSingle]
type clientCertificateResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateResponseSingleError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    clientCertificateResponseSingleErrorJSON `json:"-"`
}

// clientCertificateResponseSingleErrorJSON contains the JSON metadata for the
// struct [ClientCertificateResponseSingleError]
type clientCertificateResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateResponseSingleMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    clientCertificateResponseSingleMessageJSON `json:"-"`
}

// clientCertificateResponseSingleMessageJSON contains the JSON metadata for the
// struct [ClientCertificateResponseSingleMessage]
type clientCertificateResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateResponseSingleResult struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateResponseSingleResultCertificateAuthority `json:"certificate_authority"`
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
	Status ClientCertificateResponseSingleResultStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                                     `json:"validity_days"`
	JSON         clientCertificateResponseSingleResultJSON `json:"-"`
}

// clientCertificateResponseSingleResultJSON contains the JSON metadata for the
// struct [ClientCertificateResponseSingleResult]
type clientCertificateResponseSingleResultJSON struct {
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

func (r *ClientCertificateResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateResponseSingleResultCertificateAuthority struct {
	ID   string                                                        `json:"id"`
	Name string                                                        `json:"name"`
	JSON clientCertificateResponseSingleResultCertificateAuthorityJSON `json:"-"`
}

// clientCertificateResponseSingleResultCertificateAuthorityJSON contains the JSON
// metadata for the struct
// [ClientCertificateResponseSingleResultCertificateAuthority]
type clientCertificateResponseSingleResultCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseSingleResultCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateResponseSingleResultStatus string

const (
	ClientCertificateResponseSingleResultStatusActive              ClientCertificateResponseSingleResultStatus = "active"
	ClientCertificateResponseSingleResultStatusPendingReactivation ClientCertificateResponseSingleResultStatus = "pending_reactivation"
	ClientCertificateResponseSingleResultStatusPendingRevocation   ClientCertificateResponseSingleResultStatus = "pending_revocation"
	ClientCertificateResponseSingleResultStatusRevoked             ClientCertificateResponseSingleResultStatus = "revoked"
)

// Whether the API call was successful
type ClientCertificateResponseSingleSuccess bool

const (
	ClientCertificateResponseSingleSuccessTrue ClientCertificateResponseSingleSuccess = true
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
