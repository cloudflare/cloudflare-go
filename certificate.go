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
func (r *CertificateService) New(ctx context.Context, body CertificateNewParams, opts ...option.RequestOption) (res *SchemasCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing Origin CA certificate by its serial number. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SchemasCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all existing Origin CA certificates for a given zone. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) List(ctx context.Context, query CertificateListParams, opts ...option.RequestOption) (res *SchemasCertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Revoke an existing Origin CA certificate by its serial number. Use your Origin
// CA Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *CertificateResponseSingleID9CkVsmTj, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CertificateResponseSingleID9CkVsmTj struct {
	Errors   []CertificateResponseSingleID9CkVsmTjError   `json:"errors"`
	Messages []CertificateResponseSingleID9CkVsmTjMessage `json:"messages"`
	Result   CertificateResponseSingleID9CkVsmTjResult    `json:"result"`
	// Whether the API call was successful
	Success CertificateResponseSingleID9CkVsmTjSuccess `json:"success"`
	JSON    certificateResponseSingleId9CkVsmTjJSON    `json:"-"`
}

// certificateResponseSingleId9CkVsmTjJSON contains the JSON metadata for the
// struct [CertificateResponseSingleID9CkVsmTj]
type certificateResponseSingleId9CkVsmTjJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleID9CkVsmTj) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSingleID9CkVsmTjError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    certificateResponseSingleId9CkVsmTjErrorJSON `json:"-"`
}

// certificateResponseSingleId9CkVsmTjErrorJSON contains the JSON metadata for the
// struct [CertificateResponseSingleID9CkVsmTjError]
type certificateResponseSingleId9CkVsmTjErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleID9CkVsmTjError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSingleID9CkVsmTjMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    certificateResponseSingleId9CkVsmTjMessageJSON `json:"-"`
}

// certificateResponseSingleId9CkVsmTjMessageJSON contains the JSON metadata for
// the struct [CertificateResponseSingleID9CkVsmTjMessage]
type certificateResponseSingleId9CkVsmTjMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleID9CkVsmTjMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSingleID9CkVsmTjResult struct {
	// Identifier
	ID   string                                        `json:"id"`
	JSON certificateResponseSingleId9CkVsmTjResultJSON `json:"-"`
}

// certificateResponseSingleId9CkVsmTjResultJSON contains the JSON metadata for the
// struct [CertificateResponseSingleID9CkVsmTjResult]
type certificateResponseSingleId9CkVsmTjResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleID9CkVsmTjResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateResponseSingleID9CkVsmTjSuccess bool

const (
	CertificateResponseSingleID9CkVsmTjSuccessTrue CertificateResponseSingleID9CkVsmTjSuccess = true
)

type SchemasCertificateResponseCollection struct {
	Errors     []SchemasCertificateResponseCollectionError    `json:"errors"`
	Messages   []SchemasCertificateResponseCollectionMessage  `json:"messages"`
	Result     []SchemasCertificateResponseCollectionResult   `json:"result"`
	ResultInfo SchemasCertificateResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasCertificateResponseCollectionSuccess `json:"success"`
	JSON    schemasCertificateResponseCollectionJSON    `json:"-"`
}

// schemasCertificateResponseCollectionJSON contains the JSON metadata for the
// struct [SchemasCertificateResponseCollection]
type schemasCertificateResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCertificateResponseCollectionError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    schemasCertificateResponseCollectionErrorJSON `json:"-"`
}

// schemasCertificateResponseCollectionErrorJSON contains the JSON metadata for the
// struct [SchemasCertificateResponseCollectionError]
type schemasCertificateResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCertificateResponseCollectionMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    schemasCertificateResponseCollectionMessageJSON `json:"-"`
}

// schemasCertificateResponseCollectionMessageJSON contains the JSON metadata for
// the struct [SchemasCertificateResponseCollectionMessage]
type schemasCertificateResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCertificateResponseCollectionResult struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr,required"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames []interface{} `json:"hostnames,required"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType SchemasCertificateResponseCollectionResultRequestType `json:"request_type,required"`
	// The number of days for which the certificate should be valid.
	RequestedValidity SchemasCertificateResponseCollectionResultRequestedValidity `json:"requested_validity,required"`
	// Identifier
	ID string `json:"id"`
	// The Origin CA certificate. Will be newline-encoded.
	Certificate string `json:"certificate"`
	// When the certificate will expire.
	ExpiresOn time.Time                                      `json:"expires_on" format:"date-time"`
	JSON      schemasCertificateResponseCollectionResultJSON `json:"-"`
}

// schemasCertificateResponseCollectionResultJSON contains the JSON metadata for
// the struct [SchemasCertificateResponseCollectionResult]
type schemasCertificateResponseCollectionResultJSON struct {
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

func (r *SchemasCertificateResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type SchemasCertificateResponseCollectionResultRequestType string

const (
	SchemasCertificateResponseCollectionResultRequestTypeOriginRsa          SchemasCertificateResponseCollectionResultRequestType = "origin-rsa"
	SchemasCertificateResponseCollectionResultRequestTypeOriginEcc          SchemasCertificateResponseCollectionResultRequestType = "origin-ecc"
	SchemasCertificateResponseCollectionResultRequestTypeKeylessCertificate SchemasCertificateResponseCollectionResultRequestType = "keyless-certificate"
)

// The number of days for which the certificate should be valid.
type SchemasCertificateResponseCollectionResultRequestedValidity float64

const (
	SchemasCertificateResponseCollectionResultRequestedValidity7    SchemasCertificateResponseCollectionResultRequestedValidity = 7
	SchemasCertificateResponseCollectionResultRequestedValidity30   SchemasCertificateResponseCollectionResultRequestedValidity = 30
	SchemasCertificateResponseCollectionResultRequestedValidity90   SchemasCertificateResponseCollectionResultRequestedValidity = 90
	SchemasCertificateResponseCollectionResultRequestedValidity365  SchemasCertificateResponseCollectionResultRequestedValidity = 365
	SchemasCertificateResponseCollectionResultRequestedValidity730  SchemasCertificateResponseCollectionResultRequestedValidity = 730
	SchemasCertificateResponseCollectionResultRequestedValidity1095 SchemasCertificateResponseCollectionResultRequestedValidity = 1095
	SchemasCertificateResponseCollectionResultRequestedValidity5475 SchemasCertificateResponseCollectionResultRequestedValidity = 5475
)

type SchemasCertificateResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       schemasCertificateResponseCollectionResultInfoJSON `json:"-"`
}

// schemasCertificateResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [SchemasCertificateResponseCollectionResultInfo]
type schemasCertificateResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasCertificateResponseCollectionSuccess bool

const (
	SchemasCertificateResponseCollectionSuccessTrue SchemasCertificateResponseCollectionSuccess = true
)

type SchemasCertificateResponseSingle struct {
	Errors   []SchemasCertificateResponseSingleError   `json:"errors"`
	Messages []SchemasCertificateResponseSingleMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success SchemasCertificateResponseSingleSuccess `json:"success"`
	JSON    schemasCertificateResponseSingleJSON    `json:"-"`
}

// schemasCertificateResponseSingleJSON contains the JSON metadata for the struct
// [SchemasCertificateResponseSingle]
type schemasCertificateResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCertificateResponseSingleError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    schemasCertificateResponseSingleErrorJSON `json:"-"`
}

// schemasCertificateResponseSingleErrorJSON contains the JSON metadata for the
// struct [SchemasCertificateResponseSingleError]
type schemasCertificateResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCertificateResponseSingleMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    schemasCertificateResponseSingleMessageJSON `json:"-"`
}

// schemasCertificateResponseSingleMessageJSON contains the JSON metadata for the
// struct [SchemasCertificateResponseSingleMessage]
type schemasCertificateResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificateResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasCertificateResponseSingleSuccess bool

const (
	SchemasCertificateResponseSingleSuccessTrue SchemasCertificateResponseSingleSuccess = true
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
