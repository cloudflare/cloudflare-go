// File generated from our OpenAPI spec by Stainless.

package ssl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// VerificationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVerificationService] method
// instead.
type VerificationService struct {
	Options []option.RequestOption
}

// NewVerificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerificationService(opts ...option.RequestOption) (r *VerificationService) {
	r = &VerificationService{}
	r.Options = opts
	return
}

// Edit SSL validation method for a certificate pack. A PATCH request will request
// an immediate validation check on any certificate, and return the updated status.
// If a validation method is provided, the validation will be immediately attempted
// using that method.
func (r *VerificationService) Edit(ctx context.Context, certificatePackID string, params VerificationEditParams, opts ...option.RequestOption) (res *VerificationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VerificationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", params.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get SSL Verification Info for a Zone.
func (r *VerificationService) Get(ctx context.Context, params VerificationGetParams, opts ...option.RequestOption) (res *[]VerificationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VerificationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VerificationEditResponse struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod VerificationEditResponseValidationMethod `json:"validation_method"`
	JSON             verificationEditResponseJSON             `json:"-"`
}

// verificationEditResponseJSON contains the JSON metadata for the struct
// [VerificationEditResponse]
type verificationEditResponseJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VerificationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseJSON) RawJSON() string {
	return r.raw
}

// Desired validation method.
type VerificationEditResponseValidationMethod string

const (
	VerificationEditResponseValidationMethodHTTP  VerificationEditResponseValidationMethod = "http"
	VerificationEditResponseValidationMethodCNAME VerificationEditResponseValidationMethod = "cname"
	VerificationEditResponseValidationMethodTXT   VerificationEditResponseValidationMethod = "txt"
	VerificationEditResponseValidationMethodEmail VerificationEditResponseValidationMethod = "email"
)

type VerificationGetResponse struct {
	// Current status of certificate.
	CertificateStatus VerificationGetResponseCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUUID string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature VerificationGetResponseSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod VerificationGetResponseValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo VerificationGetResponseVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType VerificationGetResponseVerificationType `json:"verification_type"`
	JSON             verificationGetResponseJSON             `json:"-"`
}

// verificationGetResponseJSON contains the JSON metadata for the struct
// [VerificationGetResponse]
type verificationGetResponseJSON struct {
	CertificateStatus  apijson.Field
	BrandCheck         apijson.Field
	CertPackUUID       apijson.Field
	Signature          apijson.Field
	ValidationMethod   apijson.Field
	VerificationInfo   apijson.Field
	VerificationStatus apijson.Field
	VerificationType   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VerificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of certificate.
type VerificationGetResponseCertificateStatus string

const (
	VerificationGetResponseCertificateStatusInitializing      VerificationGetResponseCertificateStatus = "initializing"
	VerificationGetResponseCertificateStatusAuthorizing       VerificationGetResponseCertificateStatus = "authorizing"
	VerificationGetResponseCertificateStatusActive            VerificationGetResponseCertificateStatus = "active"
	VerificationGetResponseCertificateStatusExpired           VerificationGetResponseCertificateStatus = "expired"
	VerificationGetResponseCertificateStatusIssuing           VerificationGetResponseCertificateStatus = "issuing"
	VerificationGetResponseCertificateStatusTimingOut         VerificationGetResponseCertificateStatus = "timing_out"
	VerificationGetResponseCertificateStatusPendingDeployment VerificationGetResponseCertificateStatus = "pending_deployment"
)

// Certificate's signature algorithm.
type VerificationGetResponseSignature string

const (
	VerificationGetResponseSignatureEcdsaWithSha256 VerificationGetResponseSignature = "ECDSAWithSHA256"
	VerificationGetResponseSignatureSha1WithRsa     VerificationGetResponseSignature = "SHA1WithRSA"
	VerificationGetResponseSignatureSha256WithRsa   VerificationGetResponseSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type VerificationGetResponseValidationMethod string

const (
	VerificationGetResponseValidationMethodHTTP  VerificationGetResponseValidationMethod = "http"
	VerificationGetResponseValidationMethodCNAME VerificationGetResponseValidationMethod = "cname"
	VerificationGetResponseValidationMethodTXT   VerificationGetResponseValidationMethod = "txt"
)

// Certificate's required verification information.
type VerificationGetResponseVerificationInfo struct {
	// Name of CNAME record.
	RecordName VerificationGetResponseVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget VerificationGetResponseVerificationInfoRecordTarget `json:"record_target"`
	JSON         verificationGetResponseVerificationInfoJSON         `json:"-"`
}

// verificationGetResponseVerificationInfoJSON contains the JSON metadata for the
// struct [VerificationGetResponseVerificationInfo]
type verificationGetResponseVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VerificationGetResponseVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationGetResponseVerificationInfoJSON) RawJSON() string {
	return r.raw
}

// Name of CNAME record.
type VerificationGetResponseVerificationInfoRecordName string

const (
	VerificationGetResponseVerificationInfoRecordNameRecordName VerificationGetResponseVerificationInfoRecordName = "record_name"
	VerificationGetResponseVerificationInfoRecordNameHTTPURL    VerificationGetResponseVerificationInfoRecordName = "http_url"
	VerificationGetResponseVerificationInfoRecordNameCNAME      VerificationGetResponseVerificationInfoRecordName = "cname"
	VerificationGetResponseVerificationInfoRecordNameTXTName    VerificationGetResponseVerificationInfoRecordName = "txt_name"
)

// Target of CNAME record.
type VerificationGetResponseVerificationInfoRecordTarget string

const (
	VerificationGetResponseVerificationInfoRecordTargetRecordValue VerificationGetResponseVerificationInfoRecordTarget = "record_value"
	VerificationGetResponseVerificationInfoRecordTargetHTTPBody    VerificationGetResponseVerificationInfoRecordTarget = "http_body"
	VerificationGetResponseVerificationInfoRecordTargetCNAMETarget VerificationGetResponseVerificationInfoRecordTarget = "cname_target"
	VerificationGetResponseVerificationInfoRecordTargetTXTValue    VerificationGetResponseVerificationInfoRecordTarget = "txt_value"
)

// Method of verification.
type VerificationGetResponseVerificationType string

const (
	VerificationGetResponseVerificationTypeCNAME   VerificationGetResponseVerificationType = "cname"
	VerificationGetResponseVerificationTypeMetaTag VerificationGetResponseVerificationType = "meta tag"
)

type VerificationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Desired validation method.
	ValidationMethod param.Field[VerificationEditParamsValidationMethod] `json:"validation_method,required"`
}

func (r VerificationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired validation method.
type VerificationEditParamsValidationMethod string

const (
	VerificationEditParamsValidationMethodHTTP  VerificationEditParamsValidationMethod = "http"
	VerificationEditParamsValidationMethodCNAME VerificationEditParamsValidationMethod = "cname"
	VerificationEditParamsValidationMethodTXT   VerificationEditParamsValidationMethod = "txt"
	VerificationEditParamsValidationMethodEmail VerificationEditParamsValidationMethod = "email"
)

type VerificationEditResponseEnvelope struct {
	Errors   []VerificationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VerificationEditResponseEnvelopeMessages `json:"messages,required"`
	Result   VerificationEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success VerificationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    verificationEditResponseEnvelopeJSON    `json:"-"`
}

// verificationEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [VerificationEditResponseEnvelope]
type verificationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type VerificationEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    verificationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// verificationEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VerificationEditResponseEnvelopeErrors]
type verificationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VerificationEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    verificationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// verificationEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VerificationEditResponseEnvelopeMessages]
type verificationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VerificationEditResponseEnvelopeSuccess bool

const (
	VerificationEditResponseEnvelopeSuccessTrue VerificationEditResponseEnvelopeSuccess = true
)

type VerificationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Immediately retry SSL Verification.
	Retry param.Field[VerificationGetParamsRetry] `query:"retry"`
}

// URLQuery serializes [VerificationGetParams]'s query parameters as `url.Values`.
func (r VerificationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Immediately retry SSL Verification.
type VerificationGetParamsRetry bool

const (
	VerificationGetParamsRetryTrue VerificationGetParamsRetry = true
)

type VerificationGetResponseEnvelope struct {
	Result []VerificationGetResponse           `json:"result"`
	JSON   verificationGetResponseEnvelopeJSON `json:"-"`
}

// verificationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VerificationGetResponseEnvelope]
type verificationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
