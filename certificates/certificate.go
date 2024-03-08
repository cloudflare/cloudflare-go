// File generated from our OpenAPI spec by Stainless.

package certificates

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
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
	var env CertificateNewResponseEnvelope
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all existing Origin CA certificates for a given zone. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) List(ctx context.Context, query CertificateListParams, opts ...option.RequestOption) (res *[]OriginCACertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateListResponseEnvelope
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revoke an existing Origin CA certificate by its serial number. Use your Origin
// CA Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) Delete(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *CertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateDeleteResponseEnvelope
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an existing Origin CA certificate by its serial number. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) Get(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *CertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateGetResponseEnvelope
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginCACertificate struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr,required"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames []interface{} `json:"hostnames,required"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType OriginCACertificateRequestType `json:"request_type,required"`
	// The number of days for which the certificate should be valid.
	RequestedValidity OriginCACertificateRequestedValidity `json:"requested_validity,required"`
	// Identifier
	ID string `json:"id"`
	// The Origin CA certificate. Will be newline-encoded.
	Certificate string `json:"certificate"`
	// When the certificate will expire.
	ExpiresOn time.Time               `json:"expires_on" format:"date-time"`
	JSON      originCACertificateJSON `json:"-"`
}

// originCACertificateJSON contains the JSON metadata for the struct
// [OriginCACertificate]
type originCACertificateJSON struct {
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

func (r *OriginCACertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateJSON) RawJSON() string {
	return r.raw
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type OriginCACertificateRequestType string

const (
	OriginCACertificateRequestTypeOriginRsa          OriginCACertificateRequestType = "origin-rsa"
	OriginCACertificateRequestTypeOriginEcc          OriginCACertificateRequestType = "origin-ecc"
	OriginCACertificateRequestTypeKeylessCertificate OriginCACertificateRequestType = "keyless-certificate"
)

// The number of days for which the certificate should be valid.
type OriginCACertificateRequestedValidity float64

const (
	OriginCACertificateRequestedValidity7    OriginCACertificateRequestedValidity = 7
	OriginCACertificateRequestedValidity30   OriginCACertificateRequestedValidity = 30
	OriginCACertificateRequestedValidity90   OriginCACertificateRequestedValidity = 90
	OriginCACertificateRequestedValidity365  OriginCACertificateRequestedValidity = 365
	OriginCACertificateRequestedValidity730  OriginCACertificateRequestedValidity = 730
	OriginCACertificateRequestedValidity1095 OriginCACertificateRequestedValidity = 1095
	OriginCACertificateRequestedValidity5475 OriginCACertificateRequestedValidity = 5475
)

// Union satisfied by [certificates.CertificateNewResponseUnknown] or
// [shared.UnionString].
type CertificateNewResponse interface {
	ImplementsCertificatesCertificateNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CertificateNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CertificateDeleteResponse struct {
	// Identifier
	ID   string                        `json:"id"`
	JSON certificateDeleteResponseJSON `json:"-"`
}

// certificateDeleteResponseJSON contains the JSON metadata for the struct
// [CertificateDeleteResponse]
type certificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [certificates.CertificateGetResponseUnknown] or
// [shared.UnionString].
type CertificateGetResponse interface {
	ImplementsCertificatesCertificateGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CertificateGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

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

type CertificateNewResponseEnvelope struct {
	Errors   []CertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateNewResponseEnvelopeJSON    `json:"-"`
}

// certificateNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificateNewResponseEnvelope]
type certificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificateNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    certificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificateNewResponseEnvelopeErrors]
type certificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificateNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    certificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CertificateNewResponseEnvelopeMessages]
type certificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificateNewResponseEnvelopeSuccess bool

const (
	CertificateNewResponseEnvelopeSuccessTrue CertificateNewResponseEnvelopeSuccess = true
)

type CertificateListParams struct {
}

type CertificateListResponseEnvelope struct {
	Errors   []CertificateListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateListResponseEnvelopeMessages `json:"messages,required"`
	Result   []OriginCACertificate                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CertificateListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CertificateListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       certificateListResponseEnvelopeJSON       `json:"-"`
}

// certificateListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificateListResponseEnvelope]
type certificateListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificateListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    certificateListResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificateListResponseEnvelopeErrors]
type certificateListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificateListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    certificateListResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CertificateListResponseEnvelopeMessages]
type certificateListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificateListResponseEnvelopeSuccess bool

const (
	CertificateListResponseEnvelopeSuccessTrue CertificateListResponseEnvelopeSuccess = true
)

type CertificateListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       certificateListResponseEnvelopeResultInfoJSON `json:"-"`
}

// certificateListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [CertificateListResponseEnvelopeResultInfo]
type certificateListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type CertificateDeleteResponseEnvelope struct {
	Errors   []CertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// certificateDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificateDeleteResponseEnvelope]
type certificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    certificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificateDeleteResponseEnvelopeErrors]
type certificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    certificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CertificateDeleteResponseEnvelopeMessages]
type certificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificateDeleteResponseEnvelopeSuccess bool

const (
	CertificateDeleteResponseEnvelopeSuccessTrue CertificateDeleteResponseEnvelopeSuccess = true
)

type CertificateGetResponseEnvelope struct {
	Errors   []CertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateGetResponseEnvelopeJSON    `json:"-"`
}

// certificateGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificateGetResponseEnvelope]
type certificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificateGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    certificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificateGetResponseEnvelopeErrors]
type certificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificateGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    certificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CertificateGetResponseEnvelopeMessages]
type certificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificateGetResponseEnvelopeSuccess bool

const (
	CertificateGetResponseEnvelopeSuccessTrue CertificateGetResponseEnvelopeSuccess = true
)
