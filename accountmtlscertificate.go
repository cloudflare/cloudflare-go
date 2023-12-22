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

// AccountMtlsCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMtlsCertificateService]
// method instead.
type AccountMtlsCertificateService struct {
	Options      []option.RequestOption
	Associations *AccountMtlsCertificateAssociationService
}

// NewAccountMtlsCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMtlsCertificateService(opts ...option.RequestOption) (r *AccountMtlsCertificateService) {
	r = &AccountMtlsCertificateService{}
	r.Options = opts
	r.Associations = NewAccountMtlsCertificateAssociationService(opts...)
	return
}

// Fetches a single mTLS certificate.
func (r *AccountMtlsCertificateService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *MtlsManagementCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the mTLS certificate unless the certificate is in use by one or more
// Cloudflare services.
func (r *AccountMtlsCertificateService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *MtlsManagementCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists all mTLS certificates.
func (r *AccountMtlsCertificateService) MTlsCertificateManagementListMTlsCertificates(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MtlsManagementCertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a certificate that you want to use with mTLS-enabled Cloudflare services.
func (r *AccountMtlsCertificateService) MTlsCertificateManagementUploadMTlsCertificate(ctx context.Context, accountIdentifier string, body AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateParams, opts ...option.RequestOption) (res *CertificateResponseSinglePost, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CertificateResponseSinglePost struct {
	Errors   []CertificateResponseSinglePostError   `json:"errors"`
	Messages []CertificateResponseSinglePostMessage `json:"messages"`
	Result   CertificateResponseSinglePostResult    `json:"result"`
	// Whether the API call was successful
	Success CertificateResponseSinglePostSuccess `json:"success"`
	JSON    certificateResponseSinglePostJSON    `json:"-"`
}

// certificateResponseSinglePostJSON contains the JSON metadata for the struct
// [CertificateResponseSinglePost]
type certificateResponseSinglePostJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSinglePost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSinglePostError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    certificateResponseSinglePostErrorJSON `json:"-"`
}

// certificateResponseSinglePostErrorJSON contains the JSON metadata for the struct
// [CertificateResponseSinglePostError]
type certificateResponseSinglePostErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSinglePostError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSinglePostMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    certificateResponseSinglePostMessageJSON `json:"-"`
}

// certificateResponseSinglePostMessageJSON contains the JSON metadata for the
// struct [CertificateResponseSinglePostMessage]
type certificateResponseSinglePostMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSinglePostMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSinglePostResult struct {
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
	UploadedOn time.Time                               `json:"uploaded_on" format:"date-time"`
	JSON       certificateResponseSinglePostResultJSON `json:"-"`
}

// certificateResponseSinglePostResultJSON contains the JSON metadata for the
// struct [CertificateResponseSinglePostResult]
type certificateResponseSinglePostResultJSON struct {
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

func (r *CertificateResponseSinglePostResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateResponseSinglePostSuccess bool

const (
	CertificateResponseSinglePostSuccessTrue CertificateResponseSinglePostSuccess = true
)

type MtlsManagementCertificateResponseCollection struct {
	Errors     []MtlsManagementCertificateResponseCollectionError    `json:"errors"`
	Messages   []MtlsManagementCertificateResponseCollectionMessage  `json:"messages"`
	Result     []MtlsManagementCertificateResponseCollectionResult   `json:"result"`
	ResultInfo MtlsManagementCertificateResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success MtlsManagementCertificateResponseCollectionSuccess `json:"success"`
	JSON    mtlsManagementCertificateResponseCollectionJSON    `json:"-"`
}

// mtlsManagementCertificateResponseCollectionJSON contains the JSON metadata for
// the struct [MtlsManagementCertificateResponseCollection]
type mtlsManagementCertificateResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseCollectionError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    mtlsManagementCertificateResponseCollectionErrorJSON `json:"-"`
}

// mtlsManagementCertificateResponseCollectionErrorJSON contains the JSON metadata
// for the struct [MtlsManagementCertificateResponseCollectionError]
type mtlsManagementCertificateResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseCollectionMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    mtlsManagementCertificateResponseCollectionMessageJSON `json:"-"`
}

// mtlsManagementCertificateResponseCollectionMessageJSON contains the JSON
// metadata for the struct [MtlsManagementCertificateResponseCollectionMessage]
type mtlsManagementCertificateResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseCollectionResult struct {
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
	UploadedOn time.Time                                             `json:"uploaded_on" format:"date-time"`
	JSON       mtlsManagementCertificateResponseCollectionResultJSON `json:"-"`
}

// mtlsManagementCertificateResponseCollectionResultJSON contains the JSON metadata
// for the struct [MtlsManagementCertificateResponseCollectionResult]
type mtlsManagementCertificateResponseCollectionResultJSON struct {
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

func (r *MtlsManagementCertificateResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseCollectionResultInfo struct {
	Count      interface{}                                               `json:"count"`
	Page       interface{}                                               `json:"page"`
	PerPage    interface{}                                               `json:"per_page"`
	TotalCount interface{}                                               `json:"total_count"`
	TotalPages interface{}                                               `json:"total_pages"`
	JSON       mtlsManagementCertificateResponseCollectionResultInfoJSON `json:"-"`
}

// mtlsManagementCertificateResponseCollectionResultInfoJSON contains the JSON
// metadata for the struct [MtlsManagementCertificateResponseCollectionResultInfo]
type mtlsManagementCertificateResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsManagementCertificateResponseCollectionSuccess bool

const (
	MtlsManagementCertificateResponseCollectionSuccessTrue MtlsManagementCertificateResponseCollectionSuccess = true
)

type MtlsManagementCertificateResponseSingle struct {
	Errors   []MtlsManagementCertificateResponseSingleError   `json:"errors"`
	Messages []MtlsManagementCertificateResponseSingleMessage `json:"messages"`
	Result   MtlsManagementCertificateResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success MtlsManagementCertificateResponseSingleSuccess `json:"success"`
	JSON    mtlsManagementCertificateResponseSingleJSON    `json:"-"`
}

// mtlsManagementCertificateResponseSingleJSON contains the JSON metadata for the
// struct [MtlsManagementCertificateResponseSingle]
type mtlsManagementCertificateResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseSingleError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    mtlsManagementCertificateResponseSingleErrorJSON `json:"-"`
}

// mtlsManagementCertificateResponseSingleErrorJSON contains the JSON metadata for
// the struct [MtlsManagementCertificateResponseSingleError]
type mtlsManagementCertificateResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseSingleMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    mtlsManagementCertificateResponseSingleMessageJSON `json:"-"`
}

// mtlsManagementCertificateResponseSingleMessageJSON contains the JSON metadata
// for the struct [MtlsManagementCertificateResponseSingleMessage]
type mtlsManagementCertificateResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsManagementCertificateResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsManagementCertificateResponseSingleResult struct {
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
	UploadedOn time.Time                                         `json:"uploaded_on" format:"date-time"`
	JSON       mtlsManagementCertificateResponseSingleResultJSON `json:"-"`
}

// mtlsManagementCertificateResponseSingleResultJSON contains the JSON metadata for
// the struct [MtlsManagementCertificateResponseSingleResult]
type mtlsManagementCertificateResponseSingleResultJSON struct {
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

func (r *MtlsManagementCertificateResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsManagementCertificateResponseSingleSuccess bool

const (
	MtlsManagementCertificateResponseSingleSuccessTrue MtlsManagementCertificateResponseSingleSuccess = true
)

type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateParams struct {
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca param.Field[bool] `json:"ca,required"`
	// The uploaded root CA certificate.
	Certificates param.Field[string] `json:"certificates,required"`
	// The private key for the certificate
	PrivateKey param.Field[string] `json:"private_key,required"`
	// Optional unique name for the certificate. Only used for human readability.
	Name param.Field[string] `json:"name"`
}

func (r AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
