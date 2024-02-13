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
func (r *SSLVerificationService) Update(ctx context.Context, zoneID string, certificatePackID string, body SSLVerificationUpdateParams, opts ...option.RequestOption) (res *SSLVerificationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLVerificationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get SSL Verification Info for a Zone.
func (r *SSLVerificationService) SSLVerificationSSLVerificationDetails(ctx context.Context, zoneID string, query SSLVerificationSSLVerificationSSLVerificationDetailsParams, opts ...option.RequestOption) (res *[]SSLVerificationSSLVerificationSSLVerificationDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLVerificationSSLVerificationSSLVerificationDetailsResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/verification", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLVerificationUpdateResponse struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod SSLVerificationUpdateResponseValidationMethod `json:"validation_method"`
	JSON             sslVerificationUpdateResponseJSON             `json:"-"`
}

// sslVerificationUpdateResponseJSON contains the JSON metadata for the struct
// [SSLVerificationUpdateResponse]
type sslVerificationUpdateResponseJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSLVerificationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Desired validation method.
type SSLVerificationUpdateResponseValidationMethod string

const (
	SSLVerificationUpdateResponseValidationMethodHTTP  SSLVerificationUpdateResponseValidationMethod = "http"
	SSLVerificationUpdateResponseValidationMethodCname SSLVerificationUpdateResponseValidationMethod = "cname"
	SSLVerificationUpdateResponseValidationMethodTxt   SSLVerificationUpdateResponseValidationMethod = "txt"
	SSLVerificationUpdateResponseValidationMethodEmail SSLVerificationUpdateResponseValidationMethod = "email"
)

type SSLVerificationSSLVerificationSSLVerificationDetailsResponse struct {
	// Current status of certificate.
	CertificateStatus SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUuid string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationType `json:"verification_type"`
	JSON             sslVerificationSSLVerificationSSLVerificationDetailsResponseJSON             `json:"-"`
}

// sslVerificationSSLVerificationSSLVerificationDetailsResponseJSON contains the
// JSON metadata for the struct
// [SSLVerificationSSLVerificationSSLVerificationDetailsResponse]
type sslVerificationSSLVerificationSSLVerificationDetailsResponseJSON struct {
	CertificateStatus  apijson.Field
	BrandCheck         apijson.Field
	CertPackUuid       apijson.Field
	Signature          apijson.Field
	ValidationMethod   apijson.Field
	VerificationInfo   apijson.Field
	VerificationStatus apijson.Field
	VerificationType   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SSLVerificationSSLVerificationSSLVerificationDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of certificate.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus string

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusInitializing      SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "initializing"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusAuthorizing       SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "authorizing"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusActive            SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "active"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusExpired           SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "expired"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusIssuing           SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "issuing"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusTimingOut         SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "timing_out"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatusPendingDeployment SSLVerificationSSLVerificationSSLVerificationDetailsResponseCertificateStatus = "pending_deployment"
)

// Certificate's signature algorithm.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignature string

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignatureEcdsaWithSha256 SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignature = "ECDSAWithSHA256"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignatureSha1WithRsa     SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignature = "SHA1WithRSA"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignatureSha256WithRsa   SSLVerificationSSLVerificationSSLVerificationDetailsResponseSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethod string

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethodHTTP  SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethod = "http"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethodCname SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethod = "cname"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethodTxt   SSLVerificationSSLVerificationSSLVerificationDetailsResponseValidationMethod = "txt"
)

// Certificate's required verification information.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfo struct {
	// Name of CNAME record.
	RecordName SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTarget `json:"record_target"`
	JSON         sslVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoJSON         `json:"-"`
}

// sslVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoJSON
// contains the JSON metadata for the struct
// [SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfo]
type sslVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Name of CNAME record.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordName string

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordNameRecordName SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordName = "record_name"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordNameHTTPURL    SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordName = "http_url"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordNameCname      SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordName = "cname"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordNameTxtName    SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordName = "txt_name"
)

// Target of CNAME record.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTarget string

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTargetRecordValue SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTarget = "record_value"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTargetHTTPBody    SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTarget = "http_body"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTargetCnameTarget SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTarget = "cname_target"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTargetTxtValue    SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationInfoRecordTarget = "txt_value"
)

// Method of verification.
type SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationType string

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationTypeCname   SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationType = "cname"
	SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationTypeMetaTag SSLVerificationSSLVerificationSSLVerificationDetailsResponseVerificationType = "meta tag"
)

type SSLVerificationUpdateParams struct {
	// Desired validation method.
	ValidationMethod param.Field[SSLVerificationUpdateParamsValidationMethod] `json:"validation_method,required"`
}

func (r SSLVerificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired validation method.
type SSLVerificationUpdateParamsValidationMethod string

const (
	SSLVerificationUpdateParamsValidationMethodHTTP  SSLVerificationUpdateParamsValidationMethod = "http"
	SSLVerificationUpdateParamsValidationMethodCname SSLVerificationUpdateParamsValidationMethod = "cname"
	SSLVerificationUpdateParamsValidationMethodTxt   SSLVerificationUpdateParamsValidationMethod = "txt"
	SSLVerificationUpdateParamsValidationMethodEmail SSLVerificationUpdateParamsValidationMethod = "email"
)

type SSLVerificationUpdateResponseEnvelope struct {
	Errors   []SSLVerificationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLVerificationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLVerificationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLVerificationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslVerificationUpdateResponseEnvelopeJSON    `json:"-"`
}

// sslVerificationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLVerificationUpdateResponseEnvelope]
type sslVerificationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLVerificationUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    sslVerificationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// sslVerificationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLVerificationUpdateResponseEnvelopeErrors]
type sslVerificationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLVerificationUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    sslVerificationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// sslVerificationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SSLVerificationUpdateResponseEnvelopeMessages]
type sslVerificationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLVerificationUpdateResponseEnvelopeSuccess bool

const (
	SSLVerificationUpdateResponseEnvelopeSuccessTrue SSLVerificationUpdateResponseEnvelopeSuccess = true
)

type SSLVerificationSSLVerificationSSLVerificationDetailsParams struct {
	// Immediately retry SSL Verification.
	Retry param.Field[SSLVerificationSSLVerificationSSLVerificationDetailsParamsRetry] `query:"retry"`
}

// URLQuery serializes
// [SSLVerificationSSLVerificationSSLVerificationDetailsParams]'s query parameters
// as `url.Values`.
func (r SSLVerificationSSLVerificationSSLVerificationDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Immediately retry SSL Verification.
type SSLVerificationSSLVerificationSSLVerificationDetailsParamsRetry bool

const (
	SSLVerificationSSLVerificationSSLVerificationDetailsParamsRetryTrue SSLVerificationSSLVerificationSSLVerificationDetailsParamsRetry = true
)

type SSLVerificationSSLVerificationSSLVerificationDetailsResponseEnvelope struct {
	Result []SSLVerificationSSLVerificationSSLVerificationDetailsResponse           `json:"result"`
	JSON   sslVerificationSSLVerificationSSLVerificationDetailsResponseEnvelopeJSON `json:"-"`
}

// sslVerificationSSLVerificationSSLVerificationDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SSLVerificationSSLVerificationSSLVerificationDetailsResponseEnvelope]
type sslVerificationSSLVerificationSSLVerificationDetailsResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLVerificationSSLVerificationSSLVerificationDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
