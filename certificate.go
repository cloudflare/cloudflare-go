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

// CertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCertificateService] method
// instead.
type CertificateService struct {
	Options []option.RequestOption
}

// NewCertificateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificateService(opts ...option.RequestOption) (r *CertificateService) {
	r = &CertificateService{}
	r.Options = opts
	return
}

// Create an Origin CA certificate. Use your Origin CA Key as your User Service Key
// when calling this endpoint ([see above](#requests)).
func (r *CertificateService) New(ctx context.Context, body CertificateNewParams, opts ...option.RequestOption) (res *CertificateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing Origin CA certificate by its serial number. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *CertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all existing Origin CA certificates for a given zone. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) List(ctx context.Context, query CertificateListParams, opts ...option.RequestOption) (res *CertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Revoke an existing Origin CA certificate by its serial number. Use your Origin
// CA Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *CertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CertificateNewResponse struct {
	Errors   []CertificateNewResponseError   `json:"errors"`
	Messages []CertificateNewResponseMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success CertificateNewResponseSuccess `json:"success"`
	JSON    certificateNewResponseJSON    `json:"-"`
}

// certificateNewResponseJSON contains the JSON metadata for the struct
// [CertificateNewResponse]
type certificateNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateNewResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    certificateNewResponseErrorJSON `json:"-"`
}

// certificateNewResponseErrorJSON contains the JSON metadata for the struct
// [CertificateNewResponseError]
type certificateNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateNewResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    certificateNewResponseMessageJSON `json:"-"`
}

// certificateNewResponseMessageJSON contains the JSON metadata for the struct
// [CertificateNewResponseMessage]
type certificateNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateNewResponseSuccess bool

const (
	CertificateNewResponseSuccessTrue CertificateNewResponseSuccess = true
)

type CertificateGetResponse struct {
	Errors   []CertificateGetResponseError   `json:"errors"`
	Messages []CertificateGetResponseMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success CertificateGetResponseSuccess `json:"success"`
	JSON    certificateGetResponseJSON    `json:"-"`
}

// certificateGetResponseJSON contains the JSON metadata for the struct
// [CertificateGetResponse]
type certificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    certificateGetResponseErrorJSON `json:"-"`
}

// certificateGetResponseErrorJSON contains the JSON metadata for the struct
// [CertificateGetResponseError]
type certificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    certificateGetResponseMessageJSON `json:"-"`
}

// certificateGetResponseMessageJSON contains the JSON metadata for the struct
// [CertificateGetResponseMessage]
type certificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateGetResponseSuccess bool

const (
	CertificateGetResponseSuccessTrue CertificateGetResponseSuccess = true
)

type CertificateListResponse struct {
	Errors     []CertificateListResponseError    `json:"errors"`
	Messages   []CertificateListResponseMessage  `json:"messages"`
	Result     []CertificateListResponseResult   `json:"result"`
	ResultInfo CertificateListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CertificateListResponseSuccess `json:"success"`
	JSON    certificateListResponseJSON    `json:"-"`
}

// certificateListResponseJSON contains the JSON metadata for the struct
// [CertificateListResponse]
type certificateListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateListResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    certificateListResponseErrorJSON `json:"-"`
}

// certificateListResponseErrorJSON contains the JSON metadata for the struct
// [CertificateListResponseError]
type certificateListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateListResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    certificateListResponseMessageJSON `json:"-"`
}

// certificateListResponseMessageJSON contains the JSON metadata for the struct
// [CertificateListResponseMessage]
type certificateListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateListResponseResult struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr,required"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames []interface{} `json:"hostnames,required"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType CertificateListResponseResultRequestType `json:"request_type,required"`
	// The number of days for which the certificate should be valid.
	RequestedValidity CertificateListResponseResultRequestedValidity `json:"requested_validity,required"`
	// Identifier
	ID string `json:"id"`
	// The Origin CA certificate. Will be newline-encoded.
	Certificate string `json:"certificate"`
	// When the certificate will expire.
	ExpiresOn time.Time                         `json:"expires_on" format:"date-time"`
	JSON      certificateListResponseResultJSON `json:"-"`
}

// certificateListResponseResultJSON contains the JSON metadata for the struct
// [CertificateListResponseResult]
type certificateListResponseResultJSON struct {
	Csr               apijson.Field
	Hostnames         apijson.Field
	RequestType       apijson.Field
	RequestedValidity apijson.Field
	ID                apijson.Field
	Certificate       apijson.Field
	ExpiresOn         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type CertificateListResponseResultRequestType string

const (
	CertificateListResponseResultRequestTypeOriginRsa          CertificateListResponseResultRequestType = "origin-rsa"
	CertificateListResponseResultRequestTypeOriginEcc          CertificateListResponseResultRequestType = "origin-ecc"
	CertificateListResponseResultRequestTypeKeylessCertificate CertificateListResponseResultRequestType = "keyless-certificate"
)

// The number of days for which the certificate should be valid.
type CertificateListResponseResultRequestedValidity float64

const (
	CertificateListResponseResultRequestedValidity7    CertificateListResponseResultRequestedValidity = 7
	CertificateListResponseResultRequestedValidity30   CertificateListResponseResultRequestedValidity = 30
	CertificateListResponseResultRequestedValidity90   CertificateListResponseResultRequestedValidity = 90
	CertificateListResponseResultRequestedValidity365  CertificateListResponseResultRequestedValidity = 365
	CertificateListResponseResultRequestedValidity730  CertificateListResponseResultRequestedValidity = 730
	CertificateListResponseResultRequestedValidity1095 CertificateListResponseResultRequestedValidity = 1095
	CertificateListResponseResultRequestedValidity5475 CertificateListResponseResultRequestedValidity = 5475
)

type CertificateListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       certificateListResponseResultInfoJSON `json:"-"`
}

// certificateListResponseResultInfoJSON contains the JSON metadata for the struct
// [CertificateListResponseResultInfo]
type certificateListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateListResponseSuccess bool

const (
	CertificateListResponseSuccessTrue CertificateListResponseSuccess = true
)

type CertificateDeleteResponse struct {
	Errors   []CertificateDeleteResponseError   `json:"errors"`
	Messages []CertificateDeleteResponseMessage `json:"messages"`
	Result   CertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CertificateDeleteResponseSuccess `json:"success"`
	JSON    certificateDeleteResponseJSON    `json:"-"`
}

// certificateDeleteResponseJSON contains the JSON metadata for the struct
// [CertificateDeleteResponse]
type certificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateDeleteResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    certificateDeleteResponseErrorJSON `json:"-"`
}

// certificateDeleteResponseErrorJSON contains the JSON metadata for the struct
// [CertificateDeleteResponseError]
type certificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateDeleteResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    certificateDeleteResponseMessageJSON `json:"-"`
}

// certificateDeleteResponseMessageJSON contains the JSON metadata for the struct
// [CertificateDeleteResponseMessage]
type certificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateDeleteResponseResult struct {
	// Identifier
	ID   string                              `json:"id"`
	JSON certificateDeleteResponseResultJSON `json:"-"`
}

// certificateDeleteResponseResultJSON contains the JSON metadata for the struct
// [CertificateDeleteResponseResult]
type certificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateDeleteResponseSuccess bool

const (
	CertificateDeleteResponseSuccessTrue CertificateDeleteResponseSuccess = true
)

type CertificateNewParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames param.Field[[]interface{}] `json:"hostnames"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType param.Field[CertificateNewParamsRequestType] `json:"request_type"`
	// The number of days for which the certificate should be valid.
	RequestedValidity param.Field[CertificateNewParamsRequestedValidity] `json:"requested_validity"`
}

func (r CertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type CertificateNewParamsRequestType string

const (
	CertificateNewParamsRequestTypeOriginRsa          CertificateNewParamsRequestType = "origin-rsa"
	CertificateNewParamsRequestTypeOriginEcc          CertificateNewParamsRequestType = "origin-ecc"
	CertificateNewParamsRequestTypeKeylessCertificate CertificateNewParamsRequestType = "keyless-certificate"
)

// The number of days for which the certificate should be valid.
type CertificateNewParamsRequestedValidity float64

const (
	CertificateNewParamsRequestedValidity7    CertificateNewParamsRequestedValidity = 7
	CertificateNewParamsRequestedValidity30   CertificateNewParamsRequestedValidity = 30
	CertificateNewParamsRequestedValidity90   CertificateNewParamsRequestedValidity = 90
	CertificateNewParamsRequestedValidity365  CertificateNewParamsRequestedValidity = 365
	CertificateNewParamsRequestedValidity730  CertificateNewParamsRequestedValidity = 730
	CertificateNewParamsRequestedValidity1095 CertificateNewParamsRequestedValidity = 1095
	CertificateNewParamsRequestedValidity5475 CertificateNewParamsRequestedValidity = 5475
)

type CertificateListParams struct {
}
