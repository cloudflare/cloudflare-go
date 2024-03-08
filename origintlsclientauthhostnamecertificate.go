// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// OriginTLSClientAuthHostnameCertificateService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginTLSClientAuthHostnameCertificateService] method instead.
type OriginTLSClientAuthHostnameCertificateService struct {
	Options []option.RequestOption
}

// NewOriginTLSClientAuthHostnameCertificateService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOriginTLSClientAuthHostnameCertificateService(opts ...option.RequestOption) (r *OriginTLSClientAuthHostnameCertificateService) {
	r = &OriginTLSClientAuthHostnameCertificateService{}
	r.Options = opts
	return
}

// Upload a certificate to be used for client authentication on a hostname. 10
// hostname certificates per zone are allowed.
func (r *OriginTLSClientAuthHostnameCertificateService) New(ctx context.Context, params OriginTLSClientAuthHostnameCertificateNewParams, opts ...option.RequestOption) (res *OriginTLSClientAuthHostnameCertificateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthHostnameCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *OriginTLSClientAuthHostnameCertificateService) List(ctx context.Context, query OriginTLSClientAuthHostnameCertificateListParams, opts ...option.RequestOption) (res *[]OriginTLSClientAuthHostnameCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthHostnameCertificateListResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Hostname Client Certificate
func (r *OriginTLSClientAuthHostnameCertificateService) Delete(ctx context.Context, certificateID string, body OriginTLSClientAuthHostnameCertificateDeleteParams, opts ...option.RequestOption) (res *OriginTLSClientAuthHostnameCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", body.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the certificate by ID to be used for client authentication on a hostname.
func (r *OriginTLSClientAuthHostnameCertificateService) Get(ctx context.Context, certificateID string, query OriginTLSClientAuthHostnameCertificateGetParams, opts ...option.RequestOption) (res *OriginTLSClientAuthHostnameCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthHostnameCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", query.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginTLSClientAuthHostnameCertificateNewResponse struct {
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
	Status OriginTLSClientAuthHostnameCertificateNewResponseStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                             `json:"uploaded_on" format:"date-time"`
	JSON       originTLSClientAuthHostnameCertificateNewResponseJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateNewResponseJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthHostnameCertificateNewResponse]
type originTLSClientAuthHostnameCertificateNewResponseJSON struct {
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

func (r *OriginTLSClientAuthHostnameCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type OriginTLSClientAuthHostnameCertificateNewResponseStatus string

const (
	OriginTLSClientAuthHostnameCertificateNewResponseStatusInitializing       OriginTLSClientAuthHostnameCertificateNewResponseStatus = "initializing"
	OriginTLSClientAuthHostnameCertificateNewResponseStatusPendingDeployment  OriginTLSClientAuthHostnameCertificateNewResponseStatus = "pending_deployment"
	OriginTLSClientAuthHostnameCertificateNewResponseStatusPendingDeletion    OriginTLSClientAuthHostnameCertificateNewResponseStatus = "pending_deletion"
	OriginTLSClientAuthHostnameCertificateNewResponseStatusActive             OriginTLSClientAuthHostnameCertificateNewResponseStatus = "active"
	OriginTLSClientAuthHostnameCertificateNewResponseStatusDeleted            OriginTLSClientAuthHostnameCertificateNewResponseStatus = "deleted"
	OriginTLSClientAuthHostnameCertificateNewResponseStatusDeploymentTimedOut OriginTLSClientAuthHostnameCertificateNewResponseStatus = "deployment_timed_out"
	OriginTLSClientAuthHostnameCertificateNewResponseStatusDeletionTimedOut   OriginTLSClientAuthHostnameCertificateNewResponseStatus = "deletion_timed_out"
)

type OriginTLSClientAuthHostnameCertificateListResponse struct {
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
	PrivateKey string                                                 `json:"private_key"`
	JSON       originTLSClientAuthHostnameCertificateListResponseJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateListResponseJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthHostnameCertificateListResponse]
type originTLSClientAuthHostnameCertificateListResponseJSON struct {
	ID          apijson.Field
	CertID      apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	Hostname    apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateDeleteResponse struct {
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
	Status OriginTLSClientAuthHostnameCertificateDeleteResponseStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                                `json:"uploaded_on" format:"date-time"`
	JSON       originTLSClientAuthHostnameCertificateDeleteResponseJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateDeleteResponseJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthHostnameCertificateDeleteResponse]
type originTLSClientAuthHostnameCertificateDeleteResponseJSON struct {
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

func (r *OriginTLSClientAuthHostnameCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type OriginTLSClientAuthHostnameCertificateDeleteResponseStatus string

const (
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusInitializing       OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "initializing"
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusPendingDeployment  OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "pending_deployment"
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusPendingDeletion    OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "pending_deletion"
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusActive             OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "active"
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusDeleted            OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "deleted"
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusDeploymentTimedOut OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "deployment_timed_out"
	OriginTLSClientAuthHostnameCertificateDeleteResponseStatusDeletionTimedOut   OriginTLSClientAuthHostnameCertificateDeleteResponseStatus = "deletion_timed_out"
)

type OriginTLSClientAuthHostnameCertificateGetResponse struct {
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
	Status OriginTLSClientAuthHostnameCertificateGetResponseStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                             `json:"uploaded_on" format:"date-time"`
	JSON       originTLSClientAuthHostnameCertificateGetResponseJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateGetResponseJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthHostnameCertificateGetResponse]
type originTLSClientAuthHostnameCertificateGetResponseJSON struct {
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

func (r *OriginTLSClientAuthHostnameCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type OriginTLSClientAuthHostnameCertificateGetResponseStatus string

const (
	OriginTLSClientAuthHostnameCertificateGetResponseStatusInitializing       OriginTLSClientAuthHostnameCertificateGetResponseStatus = "initializing"
	OriginTLSClientAuthHostnameCertificateGetResponseStatusPendingDeployment  OriginTLSClientAuthHostnameCertificateGetResponseStatus = "pending_deployment"
	OriginTLSClientAuthHostnameCertificateGetResponseStatusPendingDeletion    OriginTLSClientAuthHostnameCertificateGetResponseStatus = "pending_deletion"
	OriginTLSClientAuthHostnameCertificateGetResponseStatusActive             OriginTLSClientAuthHostnameCertificateGetResponseStatus = "active"
	OriginTLSClientAuthHostnameCertificateGetResponseStatusDeleted            OriginTLSClientAuthHostnameCertificateGetResponseStatus = "deleted"
	OriginTLSClientAuthHostnameCertificateGetResponseStatusDeploymentTimedOut OriginTLSClientAuthHostnameCertificateGetResponseStatus = "deployment_timed_out"
	OriginTLSClientAuthHostnameCertificateGetResponseStatusDeletionTimedOut   OriginTLSClientAuthHostnameCertificateGetResponseStatus = "deletion_timed_out"
)

type OriginTLSClientAuthHostnameCertificateNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The hostname certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The hostname certificate's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r OriginTLSClientAuthHostnameCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthHostnameCertificateNewResponseEnvelope struct {
	Errors   []OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthHostnameCertificateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthHostnameCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthHostnameCertificateNewResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthHostnameCertificateNewResponseEnvelope]
type originTLSClientAuthHostnameCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateNewResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeErrors]
type originTLSClientAuthHostnameCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateNewResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeMessages]
type originTLSClientAuthHostnameCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeSuccessTrue OriginTLSClientAuthHostnameCertificateNewResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthHostnameCertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthHostnameCertificateListResponseEnvelope struct {
	Errors   []OriginTLSClientAuthHostnameCertificateListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthHostnameCertificateListResponseEnvelopeMessages `json:"messages,required"`
	Result   []OriginTLSClientAuthHostnameCertificateListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    OriginTLSClientAuthHostnameCertificateListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo OriginTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       originTLSClientAuthHostnameCertificateListResponseEnvelopeJSON       `json:"-"`
}

// originTLSClientAuthHostnameCertificateListResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthHostnameCertificateListResponseEnvelope]
type originTLSClientAuthHostnameCertificateListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateListResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateListResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateListResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateListResponseEnvelopeErrors]
type originTLSClientAuthHostnameCertificateListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateListResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateListResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateListResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateListResponseEnvelopeMessages]
type originTLSClientAuthHostnameCertificateListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthHostnameCertificateListResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthHostnameCertificateListResponseEnvelopeSuccessTrue OriginTLSClientAuthHostnameCertificateListResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       originTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfoJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfo]
type originTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelope struct {
	Errors   []OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthHostnameCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelope]
type originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrors]
type originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessages]
type originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeSuccessTrue OriginTLSClientAuthHostnameCertificateDeleteResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthHostnameCertificateGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthHostnameCertificateGetResponseEnvelope struct {
	Errors   []OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthHostnameCertificateGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthHostnameCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthHostnameCertificateGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthHostnameCertificateGetResponseEnvelope]
type originTLSClientAuthHostnameCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeErrors]
type originTLSClientAuthHostnameCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    originTLSClientAuthHostnameCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthHostnameCertificateGetResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeMessages]
type originTLSClientAuthHostnameCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthHostnameCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeSuccessTrue OriginTLSClientAuthHostnameCertificateGetResponseEnvelopeSuccess = true
)
