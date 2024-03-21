// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *VerificationService) Get(ctx context.Context, params VerificationGetParams, opts ...option.RequestOption) (res *[]TLSCertificatesAndHostnamesVerification, err error) {
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

type TLSCertificatesAndHostnamesVerification struct {
	// Current status of certificate.
	CertificateStatus TLSCertificatesAndHostnamesVerificationCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CERTPackUUID string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature TLSCertificatesAndHostnamesVerificationSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod TLSCertificatesAndHostnamesVerificationValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo TLSCertificatesAndHostnamesVerificationVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType TLSCertificatesAndHostnamesVerificationVerificationType `json:"verification_type"`
	JSON             tlsCertificatesAndHostnamesVerificationJSON             `json:"-"`
}

// tlsCertificatesAndHostnamesVerificationJSON contains the JSON metadata for the
// struct [TLSCertificatesAndHostnamesVerification]
type tlsCertificatesAndHostnamesVerificationJSON struct {
	CertificateStatus  apijson.Field
	BrandCheck         apijson.Field
	CERTPackUUID       apijson.Field
	Signature          apijson.Field
	ValidationMethod   apijson.Field
	VerificationInfo   apijson.Field
	VerificationStatus apijson.Field
	VerificationType   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesVerificationJSON) RawJSON() string {
	return r.raw
}

// Current status of certificate.
type TLSCertificatesAndHostnamesVerificationCertificateStatus string

const (
	TLSCertificatesAndHostnamesVerificationCertificateStatusInitializing      TLSCertificatesAndHostnamesVerificationCertificateStatus = "initializing"
	TLSCertificatesAndHostnamesVerificationCertificateStatusAuthorizing       TLSCertificatesAndHostnamesVerificationCertificateStatus = "authorizing"
	TLSCertificatesAndHostnamesVerificationCertificateStatusActive            TLSCertificatesAndHostnamesVerificationCertificateStatus = "active"
	TLSCertificatesAndHostnamesVerificationCertificateStatusExpired           TLSCertificatesAndHostnamesVerificationCertificateStatus = "expired"
	TLSCertificatesAndHostnamesVerificationCertificateStatusIssuing           TLSCertificatesAndHostnamesVerificationCertificateStatus = "issuing"
	TLSCertificatesAndHostnamesVerificationCertificateStatusTimingOut         TLSCertificatesAndHostnamesVerificationCertificateStatus = "timing_out"
	TLSCertificatesAndHostnamesVerificationCertificateStatusPendingDeployment TLSCertificatesAndHostnamesVerificationCertificateStatus = "pending_deployment"
)

func (r TLSCertificatesAndHostnamesVerificationCertificateStatus) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesVerificationCertificateStatusInitializing, TLSCertificatesAndHostnamesVerificationCertificateStatusAuthorizing, TLSCertificatesAndHostnamesVerificationCertificateStatusActive, TLSCertificatesAndHostnamesVerificationCertificateStatusExpired, TLSCertificatesAndHostnamesVerificationCertificateStatusIssuing, TLSCertificatesAndHostnamesVerificationCertificateStatusTimingOut, TLSCertificatesAndHostnamesVerificationCertificateStatusPendingDeployment:
		return true
	}
	return false
}

// Certificate's signature algorithm.
type TLSCertificatesAndHostnamesVerificationSignature string

const (
	TLSCertificatesAndHostnamesVerificationSignatureEcdsaWithSha256 TLSCertificatesAndHostnamesVerificationSignature = "ECDSAWithSHA256"
	TLSCertificatesAndHostnamesVerificationSignatureSha1WithRsa     TLSCertificatesAndHostnamesVerificationSignature = "SHA1WithRSA"
	TLSCertificatesAndHostnamesVerificationSignatureSha256WithRsa   TLSCertificatesAndHostnamesVerificationSignature = "SHA256WithRSA"
)

func (r TLSCertificatesAndHostnamesVerificationSignature) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesVerificationSignatureEcdsaWithSha256, TLSCertificatesAndHostnamesVerificationSignatureSha1WithRsa, TLSCertificatesAndHostnamesVerificationSignatureSha256WithRsa:
		return true
	}
	return false
}

// Validation method in use for a certificate pack order.
type TLSCertificatesAndHostnamesVerificationValidationMethod string

const (
	TLSCertificatesAndHostnamesVerificationValidationMethodHTTP  TLSCertificatesAndHostnamesVerificationValidationMethod = "http"
	TLSCertificatesAndHostnamesVerificationValidationMethodCNAME TLSCertificatesAndHostnamesVerificationValidationMethod = "cname"
	TLSCertificatesAndHostnamesVerificationValidationMethodTXT   TLSCertificatesAndHostnamesVerificationValidationMethod = "txt"
)

func (r TLSCertificatesAndHostnamesVerificationValidationMethod) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesVerificationValidationMethodHTTP, TLSCertificatesAndHostnamesVerificationValidationMethodCNAME, TLSCertificatesAndHostnamesVerificationValidationMethodTXT:
		return true
	}
	return false
}

// Certificate's required verification information.
type TLSCertificatesAndHostnamesVerificationVerificationInfo struct {
	// Name of CNAME record.
	RecordName TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget `json:"record_target"`
	JSON         tlsCertificatesAndHostnamesVerificationVerificationInfoJSON         `json:"-"`
}

// tlsCertificatesAndHostnamesVerificationVerificationInfoJSON contains the JSON
// metadata for the struct
// [TLSCertificatesAndHostnamesVerificationVerificationInfo]
type tlsCertificatesAndHostnamesVerificationVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesVerificationVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesVerificationVerificationInfoJSON) RawJSON() string {
	return r.raw
}

// Name of CNAME record.
type TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName string

const (
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameRecordName TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "record_name"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameHTTPURL    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "http_url"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameCNAME      TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "cname"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameTXTName    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "txt_name"
)

func (r TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameRecordName, TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameHTTPURL, TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameCNAME, TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameTXTName:
		return true
	}
	return false
}

// Target of CNAME record.
type TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget string

const (
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetRecordValue TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "record_value"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetHTTPBody    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "http_body"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetCNAMETarget TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "cname_target"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetTXTValue    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "txt_value"
)

func (r TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetRecordValue, TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetHTTPBody, TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetCNAMETarget, TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetTXTValue:
		return true
	}
	return false
}

// Method of verification.
type TLSCertificatesAndHostnamesVerificationVerificationType string

const (
	TLSCertificatesAndHostnamesVerificationVerificationTypeCNAME   TLSCertificatesAndHostnamesVerificationVerificationType = "cname"
	TLSCertificatesAndHostnamesVerificationVerificationTypeMetaTag TLSCertificatesAndHostnamesVerificationVerificationType = "meta tag"
)

func (r TLSCertificatesAndHostnamesVerificationVerificationType) IsKnown() bool {
	switch r {
	case TLSCertificatesAndHostnamesVerificationVerificationTypeCNAME, TLSCertificatesAndHostnamesVerificationVerificationTypeMetaTag:
		return true
	}
	return false
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

func (r VerificationEditResponseValidationMethod) IsKnown() bool {
	switch r {
	case VerificationEditResponseValidationMethodHTTP, VerificationEditResponseValidationMethodCNAME, VerificationEditResponseValidationMethodTXT, VerificationEditResponseValidationMethodEmail:
		return true
	}
	return false
}

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

func (r VerificationEditParamsValidationMethod) IsKnown() bool {
	switch r {
	case VerificationEditParamsValidationMethodHTTP, VerificationEditParamsValidationMethodCNAME, VerificationEditParamsValidationMethodTXT, VerificationEditParamsValidationMethodEmail:
		return true
	}
	return false
}

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

func (r VerificationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VerificationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r VerificationGetParamsRetry) IsKnown() bool {
	switch r {
	case VerificationGetParamsRetryTrue:
		return true
	}
	return false
}

type VerificationGetResponseEnvelope struct {
	Result []TLSCertificatesAndHostnamesVerification `json:"result"`
	JSON   verificationGetResponseEnvelopeJSON       `json:"-"`
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
