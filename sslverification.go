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
func (r *SSLVerificationService) Get(ctx context.Context, params SSLVerificationGetParams, opts ...option.RequestOption) (res *[]SSLVerificationGetResponse, err error) {
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

func (r sslVerificationEditResponseJSON) RawJSON() string {
	return r.raw
}

// Desired validation method.
type SSLVerificationEditResponseValidationMethod string

const (
	SSLVerificationEditResponseValidationMethodHTTP  SSLVerificationEditResponseValidationMethod = "http"
	SSLVerificationEditResponseValidationMethodCNAME SSLVerificationEditResponseValidationMethod = "cname"
	SSLVerificationEditResponseValidationMethodTXT   SSLVerificationEditResponseValidationMethod = "txt"
	SSLVerificationEditResponseValidationMethodEmail SSLVerificationEditResponseValidationMethod = "email"
)

type SSLVerificationGetResponse struct {
	// Current status of certificate.
	CertificateStatus SSLVerificationGetResponseCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUUID string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature SSLVerificationGetResponseSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod SSLVerificationGetResponseValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo SSLVerificationGetResponseVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType SSLVerificationGetResponseVerificationType `json:"verification_type"`
	JSON             sslVerificationGetResponseJSON             `json:"-"`
}

// sslVerificationGetResponseJSON contains the JSON metadata for the struct
// [SSLVerificationGetResponse]
type sslVerificationGetResponseJSON struct {
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

func (r *SSLVerificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslVerificationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of certificate.
type SSLVerificationGetResponseCertificateStatus string

const (
	SSLVerificationGetResponseCertificateStatusInitializing      SSLVerificationGetResponseCertificateStatus = "initializing"
	SSLVerificationGetResponseCertificateStatusAuthorizing       SSLVerificationGetResponseCertificateStatus = "authorizing"
	SSLVerificationGetResponseCertificateStatusActive            SSLVerificationGetResponseCertificateStatus = "active"
	SSLVerificationGetResponseCertificateStatusExpired           SSLVerificationGetResponseCertificateStatus = "expired"
	SSLVerificationGetResponseCertificateStatusIssuing           SSLVerificationGetResponseCertificateStatus = "issuing"
	SSLVerificationGetResponseCertificateStatusTimingOut         SSLVerificationGetResponseCertificateStatus = "timing_out"
	SSLVerificationGetResponseCertificateStatusPendingDeployment SSLVerificationGetResponseCertificateStatus = "pending_deployment"
)

// Certificate's signature algorithm.
type SSLVerificationGetResponseSignature string

const (
	SSLVerificationGetResponseSignatureEcdsaWithSha256 SSLVerificationGetResponseSignature = "ECDSAWithSHA256"
	SSLVerificationGetResponseSignatureSha1WithRsa     SSLVerificationGetResponseSignature = "SHA1WithRSA"
	SSLVerificationGetResponseSignatureSha256WithRsa   SSLVerificationGetResponseSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type SSLVerificationGetResponseValidationMethod string

const (
	SSLVerificationGetResponseValidationMethodHTTP  SSLVerificationGetResponseValidationMethod = "http"
	SSLVerificationGetResponseValidationMethodCNAME SSLVerificationGetResponseValidationMethod = "cname"
	SSLVerificationGetResponseValidationMethodTXT   SSLVerificationGetResponseValidationMethod = "txt"
)

// Certificate's required verification information.
type SSLVerificationGetResponseVerificationInfo struct {
	// Name of CNAME record.
	RecordName SSLVerificationGetResponseVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget SSLVerificationGetResponseVerificationInfoRecordTarget `json:"record_target"`
	JSON         sslVerificationGetResponseVerificationInfoJSON         `json:"-"`
}

// sslVerificationGetResponseVerificationInfoJSON contains the JSON metadata for
// the struct [SSLVerificationGetResponseVerificationInfo]
type sslVerificationGetResponseVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SSLVerificationGetResponseVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslVerificationGetResponseVerificationInfoJSON) RawJSON() string {
	return r.raw
}

// Name of CNAME record.
type SSLVerificationGetResponseVerificationInfoRecordName string

const (
	SSLVerificationGetResponseVerificationInfoRecordNameRecordName SSLVerificationGetResponseVerificationInfoRecordName = "record_name"
	SSLVerificationGetResponseVerificationInfoRecordNameHTTPURL    SSLVerificationGetResponseVerificationInfoRecordName = "http_url"
	SSLVerificationGetResponseVerificationInfoRecordNameCNAME      SSLVerificationGetResponseVerificationInfoRecordName = "cname"
	SSLVerificationGetResponseVerificationInfoRecordNameTXTName    SSLVerificationGetResponseVerificationInfoRecordName = "txt_name"
)

// Target of CNAME record.
type SSLVerificationGetResponseVerificationInfoRecordTarget string

const (
	SSLVerificationGetResponseVerificationInfoRecordTargetRecordValue SSLVerificationGetResponseVerificationInfoRecordTarget = "record_value"
	SSLVerificationGetResponseVerificationInfoRecordTargetHTTPBody    SSLVerificationGetResponseVerificationInfoRecordTarget = "http_body"
	SSLVerificationGetResponseVerificationInfoRecordTargetCNAMETarget SSLVerificationGetResponseVerificationInfoRecordTarget = "cname_target"
	SSLVerificationGetResponseVerificationInfoRecordTargetTXTValue    SSLVerificationGetResponseVerificationInfoRecordTarget = "txt_value"
)

// Method of verification.
type SSLVerificationGetResponseVerificationType string

const (
	SSLVerificationGetResponseVerificationTypeCNAME   SSLVerificationGetResponseVerificationType = "cname"
	SSLVerificationGetResponseVerificationTypeMetaTag SSLVerificationGetResponseVerificationType = "meta tag"
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
	SSLVerificationEditParamsValidationMethodCNAME SSLVerificationEditParamsValidationMethod = "cname"
	SSLVerificationEditParamsValidationMethodTXT   SSLVerificationEditParamsValidationMethod = "txt"
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

func (r sslVerificationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r sslVerificationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r sslVerificationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result []SSLVerificationGetResponse           `json:"result"`
	JSON   sslVerificationGetResponseEnvelopeJSON `json:"-"`
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

func (r sslVerificationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
