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
func (r *AccountMtlsCertificateService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountMtlsCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the mTLS certificate unless the certificate is in use by one or more
// Cloudflare services.
func (r *AccountMtlsCertificateService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountMtlsCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists all mTLS certificates.
func (r *AccountMtlsCertificateService) MTlsCertificateManagementListMTlsCertificates(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a certificate that you want to use with mTLS-enabled Cloudflare services.
func (r *AccountMtlsCertificateService) MTlsCertificateManagementUploadMTlsCertificate(ctx context.Context, accountIdentifier string, body AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateParams, opts ...option.RequestOption) (res *AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountMtlsCertificateGetResponse struct {
	Errors   []AccountMtlsCertificateGetResponseError   `json:"errors"`
	Messages []AccountMtlsCertificateGetResponseMessage `json:"messages"`
	Result   AccountMtlsCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMtlsCertificateGetResponseSuccess `json:"success"`
	JSON    accountMtlsCertificateGetResponseJSON    `json:"-"`
}

// accountMtlsCertificateGetResponseJSON contains the JSON metadata for the struct
// [AccountMtlsCertificateGetResponse]
type accountMtlsCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountMtlsCertificateGetResponseErrorJSON `json:"-"`
}

// accountMtlsCertificateGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountMtlsCertificateGetResponseError]
type accountMtlsCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountMtlsCertificateGetResponseMessageJSON `json:"-"`
}

// accountMtlsCertificateGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountMtlsCertificateGetResponseMessage]
type accountMtlsCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateGetResponseResult struct {
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
	UploadedOn time.Time                                   `json:"uploaded_on" format:"date-time"`
	JSON       accountMtlsCertificateGetResponseResultJSON `json:"-"`
}

// accountMtlsCertificateGetResponseResultJSON contains the JSON metadata for the
// struct [AccountMtlsCertificateGetResponseResult]
type accountMtlsCertificateGetResponseResultJSON struct {
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

func (r *AccountMtlsCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMtlsCertificateGetResponseSuccess bool

const (
	AccountMtlsCertificateGetResponseSuccessTrue AccountMtlsCertificateGetResponseSuccess = true
)

type AccountMtlsCertificateDeleteResponse struct {
	Errors   []AccountMtlsCertificateDeleteResponseError   `json:"errors"`
	Messages []AccountMtlsCertificateDeleteResponseMessage `json:"messages"`
	Result   AccountMtlsCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMtlsCertificateDeleteResponseSuccess `json:"success"`
	JSON    accountMtlsCertificateDeleteResponseJSON    `json:"-"`
}

// accountMtlsCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMtlsCertificateDeleteResponse]
type accountMtlsCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountMtlsCertificateDeleteResponseErrorJSON `json:"-"`
}

// accountMtlsCertificateDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountMtlsCertificateDeleteResponseError]
type accountMtlsCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountMtlsCertificateDeleteResponseMessageJSON `json:"-"`
}

// accountMtlsCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountMtlsCertificateDeleteResponseMessage]
type accountMtlsCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateDeleteResponseResult struct {
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
	UploadedOn time.Time                                      `json:"uploaded_on" format:"date-time"`
	JSON       accountMtlsCertificateDeleteResponseResultJSON `json:"-"`
}

// accountMtlsCertificateDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountMtlsCertificateDeleteResponseResult]
type accountMtlsCertificateDeleteResponseResultJSON struct {
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

func (r *AccountMtlsCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMtlsCertificateDeleteResponseSuccess bool

const (
	AccountMtlsCertificateDeleteResponseSuccessTrue AccountMtlsCertificateDeleteResponseSuccess = true
)

type AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponse struct {
	Errors     []AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseError    `json:"errors"`
	Messages   []AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessage  `json:"messages"`
	Result     []AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResult   `json:"result"`
	ResultInfo AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseSuccess `json:"success"`
	JSON    accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseJSON    `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponse]
type accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseErrorJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseError]
type accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessageJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessage]
type accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResult struct {
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
	UploadedOn time.Time                                                                             `json:"uploaded_on" format:"date-time"`
	JSON       accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResult]
type accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultJSON struct {
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

func (r *AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfo struct {
	Count      interface{}                                                                               `json:"count"`
	Page       interface{}                                                                               `json:"page"`
	PerPage    interface{}                                                                               `json:"per_page"`
	TotalCount interface{}                                                                               `json:"total_count"`
	TotalPages interface{}                                                                               `json:"total_pages"`
	JSON       accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfoJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfo]
type accountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseSuccess bool

const (
	AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseSuccessTrue AccountMtlsCertificateMTlsCertificateManagementListMTlsCertificatesResponseSuccess = true
)

type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponse struct {
	Errors   []AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseError   `json:"errors"`
	Messages []AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessage `json:"messages"`
	Result   AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseSuccess `json:"success"`
	JSON    accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseJSON    `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponse]
type accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseErrorJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseError]
type accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessageJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessage]
type accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResult struct {
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
	UploadedOn time.Time                                                                              `json:"uploaded_on" format:"date-time"`
	JSON       accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResultJSON `json:"-"`
}

// accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResult]
type accountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResultJSON struct {
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

func (r *AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseSuccess bool

const (
	AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseSuccessTrue AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateResponseSuccess = true
)

type AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateParams struct {
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca param.Field[bool] `json:"ca,required"`
	// The uploaded root CA certificate.
	Certificates param.Field[string] `json:"certificates,required"`
	// Optional unique name for the certificate. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// The private key for the certificate
	PrivateKey param.Field[string] `json:"private_key"`
}

func (r AccountMtlsCertificateMTlsCertificateManagementUploadMTlsCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
