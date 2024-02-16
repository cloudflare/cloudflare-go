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

// MtlsCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMtlsCertificateService] method
// instead.
type MtlsCertificateService struct {
	Options      []option.RequestOption
	Associations *MtlsCertificateAssociationService
}

// NewMtlsCertificateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMtlsCertificateService(opts ...option.RequestOption) (r *MtlsCertificateService) {
	r = &MtlsCertificateService{}
	r.Options = opts
	r.Associations = NewMtlsCertificateAssociationService(opts...)
	return
}

// Upload a certificate that you want to use with mTLS-enabled Cloudflare services.
func (r *MtlsCertificateService) Update(ctx context.Context, accountID string, body MtlsCertificateUpdateParams, opts ...option.RequestOption) (res *MtlsCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MtlsCertificateUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all mTLS certificates.
func (r *MtlsCertificateService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]MtlsCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MtlsCertificateListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the mTLS certificate unless the certificate is in use by one or more
// Cloudflare services.
func (r *MtlsCertificateService) Delete(ctx context.Context, accountID string, mtlsCertificateID string, opts ...option.RequestOption) (res *MtlsCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MtlsCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single mTLS certificate.
func (r *MtlsCertificateService) Get(ctx context.Context, accountID string, mtlsCertificateID string, opts ...option.RequestOption) (res *MtlsCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MtlsCertificateGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MtlsCertificateUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                         `json:"uploaded_on" format:"date-time"`
	JSON       mtlsCertificateUpdateResponseJSON `json:"-"`
}

// mtlsCertificateUpdateResponseJSON contains the JSON metadata for the struct
// [MtlsCertificateUpdateResponse]
type mtlsCertificateUpdateResponseJSON struct {
	ID           apijson.Field
	Ca           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UpdatedAt    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MtlsCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                       `json:"uploaded_on" format:"date-time"`
	JSON       mtlsCertificateListResponseJSON `json:"-"`
}

// mtlsCertificateListResponseJSON contains the JSON metadata for the struct
// [MtlsCertificateListResponse]
type mtlsCertificateListResponseJSON struct {
	ID           apijson.Field
	Ca           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MtlsCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateDeleteResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                         `json:"uploaded_on" format:"date-time"`
	JSON       mtlsCertificateDeleteResponseJSON `json:"-"`
}

// mtlsCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [MtlsCertificateDeleteResponse]
type mtlsCertificateDeleteResponseJSON struct {
	ID           apijson.Field
	Ca           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MtlsCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                      `json:"uploaded_on" format:"date-time"`
	JSON       mtlsCertificateGetResponseJSON `json:"-"`
}

// mtlsCertificateGetResponseJSON contains the JSON metadata for the struct
// [MtlsCertificateGetResponse]
type mtlsCertificateGetResponseJSON struct {
	ID           apijson.Field
	Ca           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MtlsCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateUpdateParams struct {
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca param.Field[bool] `json:"ca,required"`
	// The uploaded root CA certificate.
	Certificates param.Field[string] `json:"certificates,required"`
	// Optional unique name for the certificate. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// The private key for the certificate
	PrivateKey param.Field[string] `json:"private_key"`
}

func (r MtlsCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MtlsCertificateUpdateResponseEnvelope struct {
	Errors   []MtlsCertificateUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MtlsCertificateUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MtlsCertificateUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MtlsCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    mtlsCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// mtlsCertificateUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MtlsCertificateUpdateResponseEnvelope]
type mtlsCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    mtlsCertificateUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// mtlsCertificateUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MtlsCertificateUpdateResponseEnvelopeErrors]
type mtlsCertificateUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    mtlsCertificateUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// mtlsCertificateUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MtlsCertificateUpdateResponseEnvelopeMessages]
type mtlsCertificateUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsCertificateUpdateResponseEnvelopeSuccess bool

const (
	MtlsCertificateUpdateResponseEnvelopeSuccessTrue MtlsCertificateUpdateResponseEnvelopeSuccess = true
)

type MtlsCertificateListResponseEnvelope struct {
	Errors   []MtlsCertificateListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MtlsCertificateListResponseEnvelopeMessages `json:"messages,required"`
	Result   []MtlsCertificateListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MtlsCertificateListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MtlsCertificateListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       mtlsCertificateListResponseEnvelopeJSON       `json:"-"`
}

// mtlsCertificateListResponseEnvelopeJSON contains the JSON metadata for the
// struct [MtlsCertificateListResponseEnvelope]
type mtlsCertificateListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    mtlsCertificateListResponseEnvelopeErrorsJSON `json:"-"`
}

// mtlsCertificateListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MtlsCertificateListResponseEnvelopeErrors]
type mtlsCertificateListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    mtlsCertificateListResponseEnvelopeMessagesJSON `json:"-"`
}

// mtlsCertificateListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MtlsCertificateListResponseEnvelopeMessages]
type mtlsCertificateListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsCertificateListResponseEnvelopeSuccess bool

const (
	MtlsCertificateListResponseEnvelopeSuccessTrue MtlsCertificateListResponseEnvelopeSuccess = true
)

type MtlsCertificateListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	TotalPages interface{}                                       `json:"total_pages"`
	JSON       mtlsCertificateListResponseEnvelopeResultInfoJSON `json:"-"`
}

// mtlsCertificateListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [MtlsCertificateListResponseEnvelopeResultInfo]
type mtlsCertificateListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateDeleteResponseEnvelope struct {
	Errors   []MtlsCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MtlsCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MtlsCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MtlsCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    mtlsCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// mtlsCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MtlsCertificateDeleteResponseEnvelope]
type mtlsCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    mtlsCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// mtlsCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MtlsCertificateDeleteResponseEnvelopeErrors]
type mtlsCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    mtlsCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// mtlsCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MtlsCertificateDeleteResponseEnvelopeMessages]
type mtlsCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsCertificateDeleteResponseEnvelopeSuccess bool

const (
	MtlsCertificateDeleteResponseEnvelopeSuccessTrue MtlsCertificateDeleteResponseEnvelopeSuccess = true
)

type MtlsCertificateGetResponseEnvelope struct {
	Errors   []MtlsCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MtlsCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MtlsCertificateGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MtlsCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    mtlsCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// mtlsCertificateGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MtlsCertificateGetResponseEnvelope]
type mtlsCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    mtlsCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// mtlsCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MtlsCertificateGetResponseEnvelopeErrors]
type mtlsCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    mtlsCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// mtlsCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MtlsCertificateGetResponseEnvelopeMessages]
type mtlsCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsCertificateGetResponseEnvelopeSuccess bool

const (
	MtlsCertificateGetResponseEnvelopeSuccessTrue MtlsCertificateGetResponseEnvelopeSuccess = true
)
