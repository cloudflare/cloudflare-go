// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SSLVerificationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSSLVerificationService] method
// instead.
type SSLVerificationService struct {
	Options []option.RequestOption
}

// NewSSLVerificationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSSLVerificationService(opts ...option.RequestOption) (r *SSLVerificationService) {
	r = &SSLVerificationService{}
	r.Options = opts
	return
}

// Edit SSL validation method for a certificate pack. A PATCH request will request
// an immediate validation check on any certificate, and return the updated status.
// If a validation method is provided, the validation will be immediately attempted
// using that method.
func (r *SSLVerificationService) Edit(ctx context.Context, certificatePackID string, params SSLVerificationEditParams, opts ...option.RequestOption) (res *SSLVerificationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLVerificationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", params.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get SSL Verification Info for a Zone.
func (r *SSLVerificationService) Get(ctx context.Context, params SSLVerificationGetParams, opts ...option.RequestOption) (res *[]TLSCertificatesAndHostnamesVerification, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLVerificationGetResponseEnvelope
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
	CertPackUUID string `json:"cert_pack_uuid"`
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
	CertPackUUID       apijson.Field
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

// Certificate's signature algorithm.
type TLSCertificatesAndHostnamesVerificationSignature string

const (
	TLSCertificatesAndHostnamesVerificationSignatureEcdsaWithSha256 TLSCertificatesAndHostnamesVerificationSignature = "ECDSAWithSHA256"
	TLSCertificatesAndHostnamesVerificationSignatureSha1WithRsa     TLSCertificatesAndHostnamesVerificationSignature = "SHA1WithRSA"
	TLSCertificatesAndHostnamesVerificationSignatureSha256WithRsa   TLSCertificatesAndHostnamesVerificationSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type TLSCertificatesAndHostnamesVerificationValidationMethod string

const (
	TLSCertificatesAndHostnamesVerificationValidationMethodHTTP  TLSCertificatesAndHostnamesVerificationValidationMethod = "http"
	TLSCertificatesAndHostnamesVerificationValidationMethodCname TLSCertificatesAndHostnamesVerificationValidationMethod = "cname"
	TLSCertificatesAndHostnamesVerificationValidationMethodTxt   TLSCertificatesAndHostnamesVerificationValidationMethod = "txt"
)

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

// Name of CNAME record.
type TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName string

const (
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameRecordName TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "record_name"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameHTTPURL    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "http_url"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameCname      TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "cname"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordNameTxtName    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordName = "txt_name"
)

// Target of CNAME record.
type TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget string

const (
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetRecordValue TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "record_value"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetHTTPBody    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "http_body"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetCnameTarget TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "cname_target"
	TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTargetTxtValue    TLSCertificatesAndHostnamesVerificationVerificationInfoRecordTarget = "txt_value"
)

// Method of verification.
type TLSCertificatesAndHostnamesVerificationVerificationType string

const (
	TLSCertificatesAndHostnamesVerificationVerificationTypeCname   TLSCertificatesAndHostnamesVerificationVerificationType = "cname"
	TLSCertificatesAndHostnamesVerificationVerificationTypeMetaTag TLSCertificatesAndHostnamesVerificationVerificationType = "meta tag"
)

type SSLVerificationEditResponse struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod SSLVerificationEditResponseValidationMethod `json:"validation_method"`
	JSON             sslVerificationEditResponseJSON             `json:"-"`
}

// sslVerificationEditResponseJSON contains the JSON metadata for the struct
// [SSLVerificationEditResponse]
type sslVerificationEditResponseJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSLVerificationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Desired validation method.
type SSLVerificationEditResponseValidationMethod string

const (
	SSLVerificationEditResponseValidationMethodHTTP  SSLVerificationEditResponseValidationMethod = "http"
	SSLVerificationEditResponseValidationMethodCname SSLVerificationEditResponseValidationMethod = "cname"
	SSLVerificationEditResponseValidationMethodTxt   SSLVerificationEditResponseValidationMethod = "txt"
	SSLVerificationEditResponseValidationMethodEmail SSLVerificationEditResponseValidationMethod = "email"
)

type SSLVerificationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Desired validation method.
	ValidationMethod param.Field[SSLVerificationEditParamsValidationMethod] `json:"validation_method,required"`
}

func (r SSLVerificationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired validation method.
type SSLVerificationEditParamsValidationMethod string

const (
	SSLVerificationEditParamsValidationMethodHTTP  SSLVerificationEditParamsValidationMethod = "http"
	SSLVerificationEditParamsValidationMethodCname SSLVerificationEditParamsValidationMethod = "cname"
	SSLVerificationEditParamsValidationMethodTxt   SSLVerificationEditParamsValidationMethod = "txt"
	SSLVerificationEditParamsValidationMethodEmail SSLVerificationEditParamsValidationMethod = "email"
)

type SSLVerificationEditResponseEnvelope struct {
	Errors   []SSLVerificationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLVerificationEditResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLVerificationEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLVerificationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslVerificationEditResponseEnvelopeJSON    `json:"-"`
}

// sslVerificationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLVerificationEditResponseEnvelope]
type sslVerificationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLVerificationEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    sslVerificationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// sslVerificationEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SSLVerificationEditResponseEnvelopeErrors]
type sslVerificationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLVerificationEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    sslVerificationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// sslVerificationEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SSLVerificationEditResponseEnvelopeMessages]
type sslVerificationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLVerificationEditResponseEnvelopeSuccess bool

const (
	SSLVerificationEditResponseEnvelopeSuccessTrue SSLVerificationEditResponseEnvelopeSuccess = true
)

type SSLVerificationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Immediately retry SSL Verification.
	Retry param.Field[SSLVerificationGetParamsRetry] `query:"retry"`
}

// URLQuery serializes [SSLVerificationGetParams]'s query parameters as
// `url.Values`.
func (r SSLVerificationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Immediately retry SSL Verification.
type SSLVerificationGetParamsRetry bool

const (
	SSLVerificationGetParamsRetryTrue SSLVerificationGetParamsRetry = true
)

type SSLVerificationGetResponseEnvelope struct {
	Result []TLSCertificatesAndHostnamesVerification `json:"result"`
	JSON   sslVerificationGetResponseEnvelopeJSON    `json:"-"`
}

// sslVerificationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SSLVerificationGetResponseEnvelope]
type sslVerificationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
