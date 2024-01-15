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
func (r *ZoneOriginTlsClientAuthHostnameService) Get(ctx context.Context, zoneIdentifier string, hostname string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameGetResponse, err error) {
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
func (r *ZoneOriginTlsClientAuthHostnameService) PerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthentication(ctx context.Context, zoneIdentifier string, body ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationParams, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneOriginTlsClientAuthHostnameGetResponse struct {
	Errors   []ZoneOriginTlsClientAuthHostnameGetResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthHostnameGetResponseMessage `json:"messages"`
	Result   ZoneOriginTlsClientAuthHostnameGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthHostnameGetResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthHostnameGetResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnameGetResponseJSON contains the JSON metadata for
// the struct [ZoneOriginTlsClientAuthHostnameGetResponse]
type zoneOriginTlsClientAuthHostnameGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameGetResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameGetResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameGetResponseErrorJSON contains the JSON metadata
// for the struct [ZoneOriginTlsClientAuthHostnameGetResponseError]
type zoneOriginTlsClientAuthHostnameGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameGetResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameGetResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameGetResponseMessageJSON contains the JSON metadata
// for the struct [ZoneOriginTlsClientAuthHostnameGetResponseMessage]
type zoneOriginTlsClientAuthHostnameGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameGetResponseResult struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus `json:"cert_status"`
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
	Status ZoneOriginTlsClientAuthHostnameGetResponseResultStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                                            `json:"updated_at" format:"date-time"`
	JSON      zoneOriginTlsClientAuthHostnameGetResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameGetResponseResultJSON contains the JSON metadata
// for the struct [ZoneOriginTlsClientAuthHostnameGetResponseResult]
type zoneOriginTlsClientAuthHostnameGetResponseResultJSON struct {
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

func (r *ZoneOriginTlsClientAuthHostnameGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus string

const (
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusInitializing       ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusActive             ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "active"
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusDeleted            ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameGetResponseResultCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameGetResponseResultStatus string

const (
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusInitializing       ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusActive             ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "active"
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusDeleted            ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameGetResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameGetResponseResultStatus = "deletion_timed_out"
)

// Whether the API call was successful
type ZoneOriginTlsClientAuthHostnameGetResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameGetResponseSuccessTrue ZoneOriginTlsClientAuthHostnameGetResponseSuccess = true
)

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponse struct {
	Errors     []ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseError    `json:"errors"`
	Messages   []ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessage  `json:"messages"`
	Result     []ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResult   `json:"result"`
	ResultInfo ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponse]
type zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseError struct {
	Code    int64                                                                                                                             `json:"code,required"`
	Message string                                                                                                                            `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseError]
type zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessage struct {
	Code    int64                                                                                                                               `json:"code,required"`
	Message string                                                                                                                              `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessage]
type zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResult struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus `json:"cert_status"`
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
	Status ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                                                                                                                          `json:"updated_at" format:"date-time"`
	JSON      zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResult]
type zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultJSON struct {
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

func (r *ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus string

const (
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusInitializing       ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "initializing"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusPendingDeployment  ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusPendingDeletion    ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusActive             ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "active"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusDeleted            ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "deleted"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus string

const (
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusInitializing       ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusActive             ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "active"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusDeleted            ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultStatus = "deletion_timed_out"
)

type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                                `json:"total_count"`
	JSON       zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfoJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfo]
type zoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseSuccessTrue ZoneOriginTlsClientAuthHostnamePerHostnameAuthenticatedOriginPullEnableOrDisableAHostnameForClientAuthenticationResponseSuccess = true
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
