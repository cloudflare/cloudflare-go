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

// ZoneOriginTlsClientAuthService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneOriginTlsClientAuthService] method instead.
type ZoneOriginTlsClientAuthService struct {
	Options   []option.RequestOption
	Hostnames *ZoneOriginTlsClientAuthHostnameService
	Settings  *ZoneOriginTlsClientAuthSettingService
}

// NewZoneOriginTlsClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneOriginTlsClientAuthService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthService) {
	r = &ZoneOriginTlsClientAuthService{}
	r.Options = opts
	r.Hostnames = NewZoneOriginTlsClientAuthHostnameService(opts...)
	r.Settings = NewZoneOriginTlsClientAuthSettingService(opts...)
	return
}

// Get Certificate Details
func (r *ZoneOriginTlsClientAuthService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Certificate
func (r *ZoneOriginTlsClientAuthService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List Certificates
func (r *ZoneOriginTlsClientAuthService) ZoneLevelAuthenticatedOriginPullsListCertificates(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *ZoneOriginTlsClientAuthService) ZoneLevelAuthenticatedOriginPullsUploadCertificate(ctx context.Context, zoneIdentifier string, body ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateParams, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneOriginTlsClientAuthGetResponse struct {
	Errors   []ZoneOriginTlsClientAuthGetResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthGetResponseMessage `json:"messages"`
	Result   interface{}                                 `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthGetResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthGetResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthGetResponseJSON contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthGetResponse]
type zoneOriginTlsClientAuthGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneOriginTlsClientAuthGetResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneOriginTlsClientAuthGetResponseError]
type zoneOriginTlsClientAuthGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneOriginTlsClientAuthGetResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneOriginTlsClientAuthGetResponseMessage]
type zoneOriginTlsClientAuthGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthGetResponseSuccess bool

const (
	ZoneOriginTlsClientAuthGetResponseSuccessTrue ZoneOriginTlsClientAuthGetResponseSuccess = true
)

type ZoneOriginTlsClientAuthDeleteResponse struct {
	Errors   []ZoneOriginTlsClientAuthDeleteResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthDeleteResponseMessage `json:"messages"`
	Result   interface{}                                    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthDeleteResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthDeleteResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneOriginTlsClientAuthDeleteResponse]
type zoneOriginTlsClientAuthDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneOriginTlsClientAuthDeleteResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthDeleteResponseErrorJSON contains the JSON metadata for
// the struct [ZoneOriginTlsClientAuthDeleteResponseError]
type zoneOriginTlsClientAuthDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneOriginTlsClientAuthDeleteResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneOriginTlsClientAuthDeleteResponseMessage]
type zoneOriginTlsClientAuthDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthDeleteResponseSuccess bool

const (
	ZoneOriginTlsClientAuthDeleteResponseSuccessTrue ZoneOriginTlsClientAuthDeleteResponseSuccess = true
)

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponse struct {
	Errors     []ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseError    `json:"errors"`
	Messages   []ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessage  `json:"messages"`
	Result     []ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResult   `json:"result"`
	ResultInfo ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponse]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseError struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseError]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessage struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessage]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate activation.
	Status ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus `json:"status"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                                                                                  `json:"uploaded_on" format:"date-time"`
	JSON       zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResult]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	ExpiresOn   apijson.Field
	Issuer      apijson.Field
	Signature   apijson.Field
	Status      apijson.Field
	UploadedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate activation.
type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus string

const (
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusInitializing       ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "initializing"
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusPendingDeployment  ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "pending_deployment"
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusPendingDeletion    ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "pending_deletion"
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusActive             ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "active"
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusDeleted            ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "deleted"
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusDeploymentTimedOut ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "deployment_timed_out"
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatusDeletionTimedOut   ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultStatus = "deletion_timed_out"
)

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                        `json:"total_count"`
	JSON       zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfoJSON `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfo]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseSuccess bool

const (
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseSuccessTrue ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsListCertificatesResponseSuccess = true
)

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponse struct {
	Errors   []ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessage `json:"messages"`
	Result   interface{}                                                                                `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponse]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseError]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessage]
type zoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseSuccess bool

const (
	ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseSuccessTrue ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateResponseSuccess = true
)

type ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateParams struct {
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r ZoneOriginTlsClientAuthZoneLevelAuthenticatedOriginPullsUploadCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
