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

// OriginTLSClientAuthHostnameService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewOriginTLSClientAuthHostnameService] method instead.
type OriginTLSClientAuthHostnameService struct {
	Options      []option.RequestOption
	Certificates *OriginTLSClientAuthHostnameCertificateService
}

// NewOriginTLSClientAuthHostnameService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOriginTLSClientAuthHostnameService(opts ...option.RequestOption) (r *OriginTLSClientAuthHostnameService) {
	r = &OriginTLSClientAuthHostnameService{}
	r.Options = opts
	r.Certificates = NewOriginTLSClientAuthHostnameCertificateService(opts...)
	return
}

// Associate a hostname to a certificate and enable, disable or invalidate the
// association. If disabled, client certificate will not be sent to the hostname
// even if activated at the zone level. 100 maximum associations on a single
// certificate are allowed. Note: Use a null value for parameter _enabled_ to
// invalidate the association.
func (r *OriginTLSClientAuthHostnameService) Update(ctx context.Context, params OriginTLSClientAuthHostnameUpdateParams, opts ...option.RequestOption) (res *[]OriginTLSClientAuthHostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthHostnameUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Hostname Status for Client Authentication
func (r *OriginTLSClientAuthHostnameService) Get(ctx context.Context, hostname string, query OriginTLSClientAuthHostnameGetParams, opts ...option.RequestOption) (res *OriginTLSClientAuthHostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthHostnameGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/%s", query.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginTLSClientAuthHostnameUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	CertID string `json:"cert_id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The hostname certificate's private key.
	PrivateKey string                                        `json:"private_key"`
	JSON       originTLSClientAuthHostnameUpdateResponseJSON `json:"-"`
}

// originTLSClientAuthHostnameUpdateResponseJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthHostnameUpdateResponse]
type originTLSClientAuthHostnameUpdateResponseJSON struct {
	ID          apijson.Field
	CertID      apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	Hostname    apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthHostnameGetResponse struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus OriginTLSClientAuthHostnameGetResponseCertStatus `json:"cert_status"`
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
	Status OriginTLSClientAuthHostnameGetResponseStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      originTLSClientAuthHostnameGetResponseJSON `json:"-"`
}

// originTLSClientAuthHostnameGetResponseJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthHostnameGetResponse]
type originTLSClientAuthHostnameGetResponseJSON struct {
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

func (r *OriginTLSClientAuthHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the certificate or the association.
type OriginTLSClientAuthHostnameGetResponseCertStatus string

const (
	OriginTLSClientAuthHostnameGetResponseCertStatusInitializing       OriginTLSClientAuthHostnameGetResponseCertStatus = "initializing"
	OriginTLSClientAuthHostnameGetResponseCertStatusPendingDeployment  OriginTLSClientAuthHostnameGetResponseCertStatus = "pending_deployment"
	OriginTLSClientAuthHostnameGetResponseCertStatusPendingDeletion    OriginTLSClientAuthHostnameGetResponseCertStatus = "pending_deletion"
	OriginTLSClientAuthHostnameGetResponseCertStatusActive             OriginTLSClientAuthHostnameGetResponseCertStatus = "active"
	OriginTLSClientAuthHostnameGetResponseCertStatusDeleted            OriginTLSClientAuthHostnameGetResponseCertStatus = "deleted"
	OriginTLSClientAuthHostnameGetResponseCertStatusDeploymentTimedOut OriginTLSClientAuthHostnameGetResponseCertStatus = "deployment_timed_out"
	OriginTLSClientAuthHostnameGetResponseCertStatusDeletionTimedOut   OriginTLSClientAuthHostnameGetResponseCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type OriginTLSClientAuthHostnameGetResponseStatus string

const (
	OriginTLSClientAuthHostnameGetResponseStatusInitializing       OriginTLSClientAuthHostnameGetResponseStatus = "initializing"
	OriginTLSClientAuthHostnameGetResponseStatusPendingDeployment  OriginTLSClientAuthHostnameGetResponseStatus = "pending_deployment"
	OriginTLSClientAuthHostnameGetResponseStatusPendingDeletion    OriginTLSClientAuthHostnameGetResponseStatus = "pending_deletion"
	OriginTLSClientAuthHostnameGetResponseStatusActive             OriginTLSClientAuthHostnameGetResponseStatus = "active"
	OriginTLSClientAuthHostnameGetResponseStatusDeleted            OriginTLSClientAuthHostnameGetResponseStatus = "deleted"
	OriginTLSClientAuthHostnameGetResponseStatusDeploymentTimedOut OriginTLSClientAuthHostnameGetResponseStatus = "deployment_timed_out"
	OriginTLSClientAuthHostnameGetResponseStatusDeletionTimedOut   OriginTLSClientAuthHostnameGetResponseStatus = "deletion_timed_out"
)

type OriginTLSClientAuthHostnameUpdateParams struct {
	// Identifier
	ZoneID param.Field[string]                                          `path:"zone_id,required"`
	Config param.Field[[]OriginTLSClientAuthHostnameUpdateParamsConfig] `json:"config,required"`
}

func (r OriginTLSClientAuthHostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthHostnameUpdateParamsConfig struct {
	// Certificate identifier tag.
	CertID param.Field[string] `json:"cert_id"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname param.Field[string] `json:"hostname"`
}

func (r OriginTLSClientAuthHostnameUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthHostnameUpdateResponseEnvelope struct {
	Errors   []OriginTLSClientAuthHostnameUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthHostnameUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []OriginTLSClientAuthHostnameUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    OriginTLSClientAuthHostnameUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo OriginTLSClientAuthHostnameUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       originTLSClientAuthHostnameUpdateResponseEnvelopeJSON       `json:"-"`
}

// originTLSClientAuthHostnameUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthHostnameUpdateResponseEnvelope]
type originTLSClientAuthHostnameUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthHostnameUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    originTLSClientAuthHostnameUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthHostnameUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthHostnameUpdateResponseEnvelopeErrors]
type originTLSClientAuthHostnameUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthHostnameUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    originTLSClientAuthHostnameUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthHostnameUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthHostnameUpdateResponseEnvelopeMessages]
type originTLSClientAuthHostnameUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthHostnameUpdateResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthHostnameUpdateResponseEnvelopeSuccessTrue OriginTLSClientAuthHostnameUpdateResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthHostnameUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       originTLSClientAuthHostnameUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// originTLSClientAuthHostnameUpdateResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [OriginTLSClientAuthHostnameUpdateResponseEnvelopeResultInfo]
type originTLSClientAuthHostnameUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthHostnameGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthHostnameGetResponseEnvelope struct {
	Errors   []OriginTLSClientAuthHostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthHostnameGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthHostnameGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthHostnameGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthHostnameGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthHostnameGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthHostnameGetResponseEnvelope]
type originTLSClientAuthHostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthHostnameGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    originTLSClientAuthHostnameGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthHostnameGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthHostnameGetResponseEnvelopeErrors]
type originTLSClientAuthHostnameGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthHostnameGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    originTLSClientAuthHostnameGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthHostnameGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthHostnameGetResponseEnvelopeMessages]
type originTLSClientAuthHostnameGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthHostnameGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthHostnameGetResponseEnvelopeSuccessTrue OriginTLSClientAuthHostnameGetResponseEnvelopeSuccess = true
)
