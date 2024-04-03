// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// HostnameCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHostnameCertificateService]
// method instead.
type HostnameCertificateService struct {
	Options []option.RequestOption
}

// NewHostnameCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHostnameCertificateService(opts ...option.RequestOption) (r *HostnameCertificateService) {
	r = &HostnameCertificateService{}
	r.Options = opts
	return
}

// Upload a certificate to be used for client authentication on a hostname. 10
// hostname certificates per zone are allowed.
func (r *HostnameCertificateService) New(ctx context.Context, params HostnameCertificateNewParams, opts ...option.RequestOption) (res *OriginTLSClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *HostnameCertificateService) List(ctx context.Context, query HostnameCertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OriginTLSClientCertificateID], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", query.ZoneID)
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

// List Certificates
func (r *HostnameCertificateService) ListAutoPaging(ctx context.Context, query HostnameCertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OriginTLSClientCertificateID] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Hostname Client Certificate
func (r *HostnameCertificateService) Delete(ctx context.Context, certificateID string, params HostnameCertificateDeleteParams, opts ...option.RequestOption) (res *OriginTLSClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", params.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the certificate by ID to be used for client authentication on a hostname.
func (r *HostnameCertificateService) Get(ctx context.Context, certificateID string, query HostnameCertificateGetParams, opts ...option.RequestOption) (res *OriginTLSClientCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", query.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginTLSClientCertificate struct {
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
	Status OriginTLSClientCertificateStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                      `json:"uploaded_on" format:"date-time"`
	JSON       originTLSClientCertificateJSON `json:"-"`
}

// originTLSClientCertificateJSON contains the JSON metadata for the struct
// [OriginTLSClientCertificate]
type originTLSClientCertificateJSON struct {
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

func (r *OriginTLSClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientCertificateJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type OriginTLSClientCertificateStatus string

const (
	OriginTLSClientCertificateStatusInitializing       OriginTLSClientCertificateStatus = "initializing"
	OriginTLSClientCertificateStatusPendingDeployment  OriginTLSClientCertificateStatus = "pending_deployment"
	OriginTLSClientCertificateStatusPendingDeletion    OriginTLSClientCertificateStatus = "pending_deletion"
	OriginTLSClientCertificateStatusActive             OriginTLSClientCertificateStatus = "active"
	OriginTLSClientCertificateStatusDeleted            OriginTLSClientCertificateStatus = "deleted"
	OriginTLSClientCertificateStatusDeploymentTimedOut OriginTLSClientCertificateStatus = "deployment_timed_out"
	OriginTLSClientCertificateStatusDeletionTimedOut   OriginTLSClientCertificateStatus = "deletion_timed_out"
)

func (r OriginTLSClientCertificateStatus) IsKnown() bool {
	switch r {
	case OriginTLSClientCertificateStatusInitializing, OriginTLSClientCertificateStatusPendingDeployment, OriginTLSClientCertificateStatusPendingDeletion, OriginTLSClientCertificateStatusActive, OriginTLSClientCertificateStatusDeleted, OriginTLSClientCertificateStatusDeploymentTimedOut, OriginTLSClientCertificateStatusDeletionTimedOut:
		return true
	}
	return false
}

type HostnameCertificateNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The hostname certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The hostname certificate's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r HostnameCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameCertificateNewResponseEnvelope struct {
	Errors   []HostnameCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientCertificate                       `json:"result,required"`
	// Whether the API call was successful
	Success HostnameCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameCertificateNewResponseEnvelope]
type hostnameCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hostnameCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HostnameCertificateNewResponseEnvelopeErrors]
type hostnameCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hostnameCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameCertificateNewResponseEnvelopeMessages]
type hostnameCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameCertificateNewResponseEnvelopeSuccess bool

const (
	HostnameCertificateNewResponseEnvelopeSuccessTrue HostnameCertificateNewResponseEnvelopeSuccess = true
)

func (r HostnameCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameCertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameCertificateDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r HostnameCertificateDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type HostnameCertificateDeleteResponseEnvelope struct {
	Errors   []HostnameCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientCertificate                          `json:"result,required"`
	// Whether the API call was successful
	Success HostnameCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameCertificateDeleteResponseEnvelope]
type hostnameCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    hostnameCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [HostnameCertificateDeleteResponseEnvelopeErrors]
type hostnameCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    hostnameCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameCertificateDeleteResponseEnvelopeMessages]
type hostnameCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameCertificateDeleteResponseEnvelopeSuccess bool

const (
	HostnameCertificateDeleteResponseEnvelopeSuccessTrue HostnameCertificateDeleteResponseEnvelopeSuccess = true
)

func (r HostnameCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameCertificateGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameCertificateGetResponseEnvelope struct {
	Errors   []HostnameCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientCertificate                       `json:"result,required"`
	// Whether the API call was successful
	Success HostnameCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameCertificateGetResponseEnvelope]
type hostnameCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hostnameCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HostnameCertificateGetResponseEnvelopeErrors]
type hostnameCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hostnameCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameCertificateGetResponseEnvelopeMessages]
type hostnameCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameCertificateGetResponseEnvelopeSuccess bool

const (
	HostnameCertificateGetResponseEnvelopeSuccessTrue HostnameCertificateGetResponseEnvelopeSuccess = true
)

func (r HostnameCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
