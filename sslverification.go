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

// Get SSL Verification Info for a Zone.
func (r *SSLVerificationService) List(ctx context.Context, zoneID string, query SSLVerificationListParams, opts ...option.RequestOption) (res *[]SSLVerificationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLVerificationListResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit SSL validation method for a certificate pack. A PATCH request will request
// an immediate validation check on any certificate, and return the updated status.
// If a validation method is provided, the validation will be immediately attempted
// using that method.
func (r *SSLVerificationService) Edit(ctx context.Context, zoneID string, certificatePackID string, body SSLVerificationEditParams, opts ...option.RequestOption) (res *SSLVerificationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLVerificationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLVerificationListResponse struct {
	// Current status of certificate.
	CertificateStatus SSLVerificationListResponseCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUUID string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature SSLVerificationListResponseSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod SSLVerificationListResponseValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo SSLVerificationListResponseVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType SSLVerificationListResponseVerificationType `json:"verification_type"`
	JSON             sslVerificationListResponseJSON             `json:"-"`
}

// sslVerificationListResponseJSON contains the JSON metadata for the struct
// [SSLVerificationListResponse]
type sslVerificationListResponseJSON struct {
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

func (r *SSLVerificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of certificate.
type SSLVerificationListResponseCertificateStatus string

const (
	SSLVerificationListResponseCertificateStatusInitializing      SSLVerificationListResponseCertificateStatus = "initializing"
	SSLVerificationListResponseCertificateStatusAuthorizing       SSLVerificationListResponseCertificateStatus = "authorizing"
	SSLVerificationListResponseCertificateStatusActive            SSLVerificationListResponseCertificateStatus = "active"
	SSLVerificationListResponseCertificateStatusExpired           SSLVerificationListResponseCertificateStatus = "expired"
	SSLVerificationListResponseCertificateStatusIssuing           SSLVerificationListResponseCertificateStatus = "issuing"
	SSLVerificationListResponseCertificateStatusTimingOut         SSLVerificationListResponseCertificateStatus = "timing_out"
	SSLVerificationListResponseCertificateStatusPendingDeployment SSLVerificationListResponseCertificateStatus = "pending_deployment"
)

// Certificate's signature algorithm.
type SSLVerificationListResponseSignature string

const (
	SSLVerificationListResponseSignatureEcdsaWithSha256 SSLVerificationListResponseSignature = "ECDSAWithSHA256"
	SSLVerificationListResponseSignatureSha1WithRsa     SSLVerificationListResponseSignature = "SHA1WithRSA"
	SSLVerificationListResponseSignatureSha256WithRsa   SSLVerificationListResponseSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type SSLVerificationListResponseValidationMethod string

const (
	SSLVerificationListResponseValidationMethodHTTP  SSLVerificationListResponseValidationMethod = "http"
	SSLVerificationListResponseValidationMethodCname SSLVerificationListResponseValidationMethod = "cname"
	SSLVerificationListResponseValidationMethodTxt   SSLVerificationListResponseValidationMethod = "txt"
)

// Certificate's required verification information.
type SSLVerificationListResponseVerificationInfo struct {
	// Name of CNAME record.
	RecordName SSLVerificationListResponseVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget SSLVerificationListResponseVerificationInfoRecordTarget `json:"record_target"`
	JSON         sslVerificationListResponseVerificationInfoJSON         `json:"-"`
}

// sslVerificationListResponseVerificationInfoJSON contains the JSON metadata for
// the struct [SSLVerificationListResponseVerificationInfo]
type sslVerificationListResponseVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SSLVerificationListResponseVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Name of CNAME record.
type SSLVerificationListResponseVerificationInfoRecordName string

const (
	SSLVerificationListResponseVerificationInfoRecordNameRecordName SSLVerificationListResponseVerificationInfoRecordName = "record_name"
	SSLVerificationListResponseVerificationInfoRecordNameHTTPURL    SSLVerificationListResponseVerificationInfoRecordName = "http_url"
	SSLVerificationListResponseVerificationInfoRecordNameCname      SSLVerificationListResponseVerificationInfoRecordName = "cname"
	SSLVerificationListResponseVerificationInfoRecordNameTxtName    SSLVerificationListResponseVerificationInfoRecordName = "txt_name"
)

// Target of CNAME record.
type SSLVerificationListResponseVerificationInfoRecordTarget string

const (
	SSLVerificationListResponseVerificationInfoRecordTargetRecordValue SSLVerificationListResponseVerificationInfoRecordTarget = "record_value"
	SSLVerificationListResponseVerificationInfoRecordTargetHTTPBody    SSLVerificationListResponseVerificationInfoRecordTarget = "http_body"
	SSLVerificationListResponseVerificationInfoRecordTargetCnameTarget SSLVerificationListResponseVerificationInfoRecordTarget = "cname_target"
	SSLVerificationListResponseVerificationInfoRecordTargetTxtValue    SSLVerificationListResponseVerificationInfoRecordTarget = "txt_value"
)

// Method of verification.
type SSLVerificationListResponseVerificationType string

const (
	SSLVerificationListResponseVerificationTypeCname   SSLVerificationListResponseVerificationType = "cname"
	SSLVerificationListResponseVerificationTypeMetaTag SSLVerificationListResponseVerificationType = "meta tag"
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

type SSLVerificationListParams struct {
	// Immediately retry SSL Verification.
	Retry param.Field[SSLVerificationListParamsRetry] `query:"retry"`
}

// URLQuery serializes [SSLVerificationListParams]'s query parameters as
// `url.Values`.
func (r SSLVerificationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Immediately retry SSL Verification.
type SSLVerificationListParamsRetry bool

const (
	SSLVerificationListParamsRetryTrue SSLVerificationListParamsRetry = true
)

type SSLVerificationListResponseEnvelope struct {
	Result []SSLVerificationListResponse           `json:"result"`
	JSON   sslVerificationListResponseEnvelopeJSON `json:"-"`
}

// sslVerificationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLVerificationListResponseEnvelope]
type sslVerificationListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLVerificationEditParams struct {
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
