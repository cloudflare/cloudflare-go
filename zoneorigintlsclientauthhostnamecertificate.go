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

// ZoneOriginTlsClientAuthHostnameCertificateService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneOriginTlsClientAuthHostnameCertificateService] method instead.
type ZoneOriginTlsClientAuthHostnameCertificateService struct {
	Options []option.RequestOption
}

// NewZoneOriginTlsClientAuthHostnameCertificateService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneOriginTlsClientAuthHostnameCertificateService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthHostnameCertificateService) {
	r = &ZoneOriginTlsClientAuthHostnameCertificateService{}
	r.Options = opts
	return
}

// Get the certificate by ID to be used for client authentication on a hostname.
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Hostname Client Certificate
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List Certificates
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) PerHostnameAuthenticatedOriginPullListCertificates(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a certificate to be used for client authentication on a hostname. 10
// hostname certificates per zone are allowed.
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) PerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificate(ctx context.Context, zoneIdentifier string, body ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateParams, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneOriginTlsClientAuthHostnameCertificateGetResponse struct {
	Errors   []ZoneOriginTlsClientAuthHostnameCertificateGetResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthHostnameCertificateGetResponseMessage `json:"messages"`
	Result   ZoneOriginTlsClientAuthHostnameCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthHostnameCertificateGetResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthHostnameCertificateGetResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateGetResponseJSON contains the JSON
// metadata for the struct [ZoneOriginTlsClientAuthHostnameCertificateGetResponse]
type zoneOriginTlsClientAuthHostnameCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificateGetResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificateGetResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateGetResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateGetResponseError]
type zoneOriginTlsClientAuthHostnameCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificateGetResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificateGetResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateGetResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateGetResponseMessage]
type zoneOriginTlsClientAuthHostnameCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificateGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                                       `json:"uploaded_on" format:"date-time"`
	JSON       zoneOriginTlsClientAuthHostnameCertificateGetResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateGetResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateGetResponseResult]
type zoneOriginTlsClientAuthHostnameCertificateGetResponseResultJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus string

const (
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusInitializing       ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusActive             ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "active"
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusDeleted            ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameCertificateGetResponseResultStatus = "deletion_timed_out"
)

// Whether the API call was successful
type ZoneOriginTlsClientAuthHostnameCertificateGetResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameCertificateGetResponseSuccessTrue ZoneOriginTlsClientAuthHostnameCertificateGetResponseSuccess = true
)

type ZoneOriginTlsClientAuthHostnameCertificateDeleteResponse struct {
	Errors   []ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessage `json:"messages"`
	Result   ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthHostnameCertificateDeleteResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateDeleteResponseJSON contains the JSON
// metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateDeleteResponse]
type zoneOriginTlsClientAuthHostnameCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificateDeleteResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateDeleteResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseError]
type zoneOriginTlsClientAuthHostnameCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessage]
type zoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                                          `json:"uploaded_on" format:"date-time"`
	JSON       zoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResult]
type zoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus string

const (
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusInitializing       ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusActive             ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "active"
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusDeleted            ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseResultStatus = "deletion_timed_out"
)

// Whether the API call was successful
type ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseSuccessTrue ZoneOriginTlsClientAuthHostnameCertificateDeleteResponseSuccess = true
)

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponse struct {
	Errors     []ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseError    `json:"errors"`
	Messages   []ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessage  `json:"messages"`
	Result     []ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResult   `json:"result"`
	ResultInfo ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponse]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseError struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseError]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessage struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessage]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResult struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus `json:"cert_status"`
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
	Status ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                                                                                                      `json:"updated_at" format:"date-time"`
	JSON      zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResult]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultJSON struct {
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

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus string

const (
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusInitializing       ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusActive             ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "active"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusDeleted            ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus string

const (
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusInitializing       ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusActive             ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "active"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusDeleted            ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultStatus = "deletion_timed_out"
)

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                            `json:"total_count"`
	JSON       zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfoJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfo]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseSuccessTrue ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullListCertificatesResponseSuccess = true
)

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponse struct {
	Errors   []ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessage `json:"messages"`
	Result   ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponse]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseError struct {
	Code    int64                                                                                                                         `json:"code,required"`
	Message string                                                                                                                        `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseError]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessage struct {
	Code    int64                                                                                                                           `json:"code,required"`
	Message string                                                                                                                          `json:"message,required"`
	JSON    zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessage]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                                                                                                      `json:"uploaded_on" format:"date-time"`
	JSON       zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResult]
type zoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus string

const (
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusInitializing       ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusActive             ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "active"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusDeleted            ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseResultStatus = "deletion_timed_out"
)

// Whether the API call was successful
type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseSuccessTrue ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateResponseSuccess = true
)

type ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateParams struct {
	// The hostname certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The hostname certificate's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r ZoneOriginTlsClientAuthHostnameCertificatePerHostnameAuthenticatedOriginPullUploadAHostnameClientCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
