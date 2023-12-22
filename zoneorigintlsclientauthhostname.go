// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneOriginTlsClientAuthHostnameService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneOriginTlsClientAuthHostnameService] method instead.
type ZoneOriginTlsClientAuthHostnameService struct {
	Options      []option.RequestOption
	Certificates *ZoneOriginTlsClientAuthHostnameCertificateService
}

// NewZoneOriginTlsClientAuthHostnameService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneOriginTlsClientAuthHostnameService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthHostnameService) {
	r = &ZoneOriginTlsClientAuthHostnameService{}
	r.Options = opts
	r.Certificates = NewZoneOriginTlsClientAuthHostnameCertificateService(opts...)
	return
}

// Get the Hostname Status for Client Authentication
func (r *ZoneOriginTlsClientAuthHostnameService) Get(ctx context.Context, zoneIdentifier string, hostname string, opts ...option.RequestOption) (res *HostnameAopSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/%s", zoneIdentifier, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Associate a hostname to a certificate and enable, disable or invalidate the
// association. If disabled, client certificate will not be sent to the hostname
// even if activated at the zone level. 100 maximum associations on a single
// certificate are allowed. Note: Use a null value for parameter _enabled_ to
// invalidate the association.
func (r *ZoneOriginTlsClientAuthHostnameService) PerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthentication(ctx context.Context, zoneIdentifier string, body ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParams, opts ...option.RequestOption) (res *HostnameAopResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type HostnameAopResponseCollection struct {
	Errors     []HostnameAopResponseCollectionError    `json:"errors"`
	Messages   []HostnameAopResponseCollectionMessage  `json:"messages"`
	Result     []HostnameAopResponseCollectionResult   `json:"result"`
	ResultInfo HostnameAopResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success HostnameAopResponseCollectionSuccess `json:"success"`
	JSON    hostnameAopResponseCollectionJSON    `json:"-"`
}

// hostnameAopResponseCollectionJSON contains the JSON metadata for the struct
// [HostnameAopResponseCollection]
type hostnameAopResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAopResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hostnameAopResponseCollectionErrorJSON `json:"-"`
}

// hostnameAopResponseCollectionErrorJSON contains the JSON metadata for the struct
// [HostnameAopResponseCollectionError]
type hostnameAopResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAopResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameAopResponseCollectionMessageJSON `json:"-"`
}

// hostnameAopResponseCollectionMessageJSON contains the JSON metadata for the
// struct [HostnameAopResponseCollectionMessage]
type hostnameAopResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAopResponseCollectionResult struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus HostnameAopResponseCollectionResultCertStatus `json:"cert_status"`
	// The time when the certificate was updated.
	CertUpdatedAt time.Time `json:"cert_updated_at" format:"date-time"`
	// The time when the certificate was uploaded.
	CertUploadedOn time.Time `json:"cert_uploaded_on" format:"date-time"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The time when the certificate was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status HostnameAopResponseCollectionResultStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      hostnameAopResponseCollectionResultJSON `json:"-"`
}

// hostnameAopResponseCollectionResultJSON contains the JSON metadata for the
// struct [HostnameAopResponseCollectionResult]
type hostnameAopResponseCollectionResultJSON struct {
	CertID         apijson.Field
	CertStatus     apijson.Field
	CertUpdatedAt  apijson.Field
	CertUploadedOn apijson.Field
	Certificate    apijson.Field
	CreatedAt      apijson.Field
	Enabled        apijson.Field
	ExpiresOn      apijson.Field
	Hostname       apijson.Field
	Issuer         apijson.Field
	SerialNumber   apijson.Field
	Signature      apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HostnameAopResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type HostnameAopResponseCollectionResultCertStatus string

const (
	HostnameAopResponseCollectionResultCertStatusInitializing       HostnameAopResponseCollectionResultCertStatus = "initializing"
	HostnameAopResponseCollectionResultCertStatusPendingDeployment  HostnameAopResponseCollectionResultCertStatus = "pending_deployment"
	HostnameAopResponseCollectionResultCertStatusPendingDeletion    HostnameAopResponseCollectionResultCertStatus = "pending_deletion"
	HostnameAopResponseCollectionResultCertStatusActive             HostnameAopResponseCollectionResultCertStatus = "active"
	HostnameAopResponseCollectionResultCertStatusDeleted            HostnameAopResponseCollectionResultCertStatus = "deleted"
	HostnameAopResponseCollectionResultCertStatusDeploymentTimedOut HostnameAopResponseCollectionResultCertStatus = "deployment_timed_out"
	HostnameAopResponseCollectionResultCertStatusDeletionTimedOut   HostnameAopResponseCollectionResultCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type HostnameAopResponseCollectionResultStatus string

const (
	HostnameAopResponseCollectionResultStatusInitializing       HostnameAopResponseCollectionResultStatus = "initializing"
	HostnameAopResponseCollectionResultStatusPendingDeployment  HostnameAopResponseCollectionResultStatus = "pending_deployment"
	HostnameAopResponseCollectionResultStatusPendingDeletion    HostnameAopResponseCollectionResultStatus = "pending_deletion"
	HostnameAopResponseCollectionResultStatusActive             HostnameAopResponseCollectionResultStatus = "active"
	HostnameAopResponseCollectionResultStatusDeleted            HostnameAopResponseCollectionResultStatus = "deleted"
	HostnameAopResponseCollectionResultStatusDeploymentTimedOut HostnameAopResponseCollectionResultStatus = "deployment_timed_out"
	HostnameAopResponseCollectionResultStatusDeletionTimedOut   HostnameAopResponseCollectionResultStatus = "deletion_timed_out"
)

type HostnameAopResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       hostnameAopResponseCollectionResultInfoJSON `json:"-"`
}

// hostnameAopResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [HostnameAopResponseCollectionResultInfo]
type hostnameAopResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HostnameAopResponseCollectionSuccess bool

const (
	HostnameAopResponseCollectionSuccessTrue HostnameAopResponseCollectionSuccess = true
)

type HostnameAopSingleResponse struct {
	Errors   []HostnameAopSingleResponseError   `json:"errors"`
	Messages []HostnameAopSingleResponseMessage `json:"messages"`
	Result   HostnameAopSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HostnameAopSingleResponseSuccess `json:"success"`
	JSON    hostnameAopSingleResponseJSON    `json:"-"`
}

// hostnameAopSingleResponseJSON contains the JSON metadata for the struct
// [HostnameAopSingleResponse]
type hostnameAopSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAopSingleResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    hostnameAopSingleResponseErrorJSON `json:"-"`
}

// hostnameAopSingleResponseErrorJSON contains the JSON metadata for the struct
// [HostnameAopSingleResponseError]
type hostnameAopSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAopSingleResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    hostnameAopSingleResponseMessageJSON `json:"-"`
}

// hostnameAopSingleResponseMessageJSON contains the JSON metadata for the struct
// [HostnameAopSingleResponseMessage]
type hostnameAopSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAopSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAopSingleResponseResult struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus HostnameAopSingleResponseResultCertStatus `json:"cert_status"`
	// The time when the certificate was updated.
	CertUpdatedAt time.Time `json:"cert_updated_at" format:"date-time"`
	// The time when the certificate was uploaded.
	CertUploadedOn time.Time `json:"cert_uploaded_on" format:"date-time"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The time when the certificate was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status HostnameAopSingleResponseResultStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      hostnameAopSingleResponseResultJSON `json:"-"`
}

// hostnameAopSingleResponseResultJSON contains the JSON metadata for the struct
// [HostnameAopSingleResponseResult]
type hostnameAopSingleResponseResultJSON struct {
	CertID         apijson.Field
	CertStatus     apijson.Field
	CertUpdatedAt  apijson.Field
	CertUploadedOn apijson.Field
	Certificate    apijson.Field
	CreatedAt      apijson.Field
	Enabled        apijson.Field
	ExpiresOn      apijson.Field
	Hostname       apijson.Field
	Issuer         apijson.Field
	SerialNumber   apijson.Field
	Signature      apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HostnameAopSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type HostnameAopSingleResponseResultCertStatus string

const (
	HostnameAopSingleResponseResultCertStatusInitializing       HostnameAopSingleResponseResultCertStatus = "initializing"
	HostnameAopSingleResponseResultCertStatusPendingDeployment  HostnameAopSingleResponseResultCertStatus = "pending_deployment"
	HostnameAopSingleResponseResultCertStatusPendingDeletion    HostnameAopSingleResponseResultCertStatus = "pending_deletion"
	HostnameAopSingleResponseResultCertStatusActive             HostnameAopSingleResponseResultCertStatus = "active"
	HostnameAopSingleResponseResultCertStatusDeleted            HostnameAopSingleResponseResultCertStatus = "deleted"
	HostnameAopSingleResponseResultCertStatusDeploymentTimedOut HostnameAopSingleResponseResultCertStatus = "deployment_timed_out"
	HostnameAopSingleResponseResultCertStatusDeletionTimedOut   HostnameAopSingleResponseResultCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type HostnameAopSingleResponseResultStatus string

const (
	HostnameAopSingleResponseResultStatusInitializing       HostnameAopSingleResponseResultStatus = "initializing"
	HostnameAopSingleResponseResultStatusPendingDeployment  HostnameAopSingleResponseResultStatus = "pending_deployment"
	HostnameAopSingleResponseResultStatusPendingDeletion    HostnameAopSingleResponseResultStatus = "pending_deletion"
	HostnameAopSingleResponseResultStatusActive             HostnameAopSingleResponseResultStatus = "active"
	HostnameAopSingleResponseResultStatusDeleted            HostnameAopSingleResponseResultStatus = "deleted"
	HostnameAopSingleResponseResultStatusDeploymentTimedOut HostnameAopSingleResponseResultStatus = "deployment_timed_out"
	HostnameAopSingleResponseResultStatusDeletionTimedOut   HostnameAopSingleResponseResultStatus = "deletion_timed_out"
)

// Whether the API call was successful
type HostnameAopSingleResponseSuccess bool

const (
	HostnameAopSingleResponseSuccessTrue HostnameAopSingleResponseSuccess = true
)

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParams struct {
	Config param.Field[[]ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParamsConfig] `json:"config,required"`
}

func (r ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParamsConfig struct {
	// Certificate identifier tag.
	CertID param.Field[string] `json:"cert_id"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname param.Field[string] `json:"hostname"`
}

func (r ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
