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

// ZoneSslVerificationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSslVerificationService]
// method instead.
type ZoneSslVerificationService struct {
	Options []option.RequestOption
}

// NewZoneSslVerificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslVerificationService(opts ...option.RequestOption) (r *ZoneSslVerificationService) {
	r = &ZoneSslVerificationService{}
	r.Options = opts
	return
}

// Edit SSL validation method for a certificate pack. A PATCH request will request
// an immediate validation check on any certificate, and return the updated status.
// If a validation method is provided, the validation will be immediately attempted
// using that method.
func (r *ZoneSslVerificationService) Update(ctx context.Context, zoneIdentifier string, certPackUuid string, body ZoneSslVerificationUpdateParams, opts ...option.RequestOption) (res *ZoneSslVerificationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", zoneIdentifier, certPackUuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get SSL Verification Info for a Zone.
func (r *ZoneSslVerificationService) SslVerificationSslVerificationDetails(ctx context.Context, zoneIdentifier string, query ZoneSslVerificationSslVerificationSslVerificationDetailsParams, opts ...option.RequestOption) (res *ZoneSslVerificationSslVerificationSslVerificationDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/verification", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSslVerificationUpdateResponse struct {
	Errors   []ZoneSslVerificationUpdateResponseError   `json:"errors"`
	Messages []ZoneSslVerificationUpdateResponseMessage `json:"messages"`
	Result   ZoneSslVerificationUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslVerificationUpdateResponseSuccess `json:"success"`
	JSON    zoneSslVerificationUpdateResponseJSON    `json:"-"`
}

// zoneSslVerificationUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSslVerificationUpdateResponse]
type zoneSslVerificationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslVerificationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslVerificationUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSslVerificationUpdateResponseErrorJSON `json:"-"`
}

// zoneSslVerificationUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSslVerificationUpdateResponseError]
type zoneSslVerificationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslVerificationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslVerificationUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSslVerificationUpdateResponseMessageJSON `json:"-"`
}

// zoneSslVerificationUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSslVerificationUpdateResponseMessage]
type zoneSslVerificationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslVerificationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslVerificationUpdateResponseResult struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod ZoneSslVerificationUpdateResponseResultValidationMethod `json:"validation_method"`
	JSON             zoneSslVerificationUpdateResponseResultJSON             `json:"-"`
}

// zoneSslVerificationUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSslVerificationUpdateResponseResult]
type zoneSslVerificationUpdateResponseResultJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneSslVerificationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Desired validation method.
type ZoneSslVerificationUpdateResponseResultValidationMethod string

const (
	ZoneSslVerificationUpdateResponseResultValidationMethodHTTP  ZoneSslVerificationUpdateResponseResultValidationMethod = "http"
	ZoneSslVerificationUpdateResponseResultValidationMethodCname ZoneSslVerificationUpdateResponseResultValidationMethod = "cname"
	ZoneSslVerificationUpdateResponseResultValidationMethodTxt   ZoneSslVerificationUpdateResponseResultValidationMethod = "txt"
	ZoneSslVerificationUpdateResponseResultValidationMethodEmail ZoneSslVerificationUpdateResponseResultValidationMethod = "email"
)

// Whether the API call was successful
type ZoneSslVerificationUpdateResponseSuccess bool

const (
	ZoneSslVerificationUpdateResponseSuccessTrue ZoneSslVerificationUpdateResponseSuccess = true
)

type ZoneSslVerificationSslVerificationSslVerificationDetailsResponse struct {
	Result []ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResult `json:"result"`
	JSON   zoneSslVerificationSslVerificationSslVerificationDetailsResponseJSON     `json:"-"`
}

// zoneSslVerificationSslVerificationSslVerificationDetailsResponseJSON contains
// the JSON metadata for the struct
// [ZoneSslVerificationSslVerificationSslVerificationDetailsResponse]
type zoneSslVerificationSslVerificationSslVerificationDetailsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslVerificationSslVerificationSslVerificationDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResult struct {
	// Current status of certificate.
	CertificateStatus ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUuid string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationType `json:"verification_type"`
	JSON             zoneSslVerificationSslVerificationSslVerificationDetailsResponseResultJSON             `json:"-"`
}

// zoneSslVerificationSslVerificationSslVerificationDetailsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResult]
type zoneSslVerificationSslVerificationSslVerificationDetailsResponseResultJSON struct {
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

func (r *ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of certificate.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus string

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusInitializing      ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "initializing"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusAuthorizing       ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "authorizing"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusActive            ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "active"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusExpired           ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "expired"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusIssuing           ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "issuing"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusTimingOut         ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "timing_out"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatusPendingDeployment ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultCertificateStatus = "pending_deployment"
)

// Certificate's signature algorithm.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignature string

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignatureEcdsaWithSha256 ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignature = "ECDSAWithSHA256"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignatureSha1WithRsa     ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignature = "SHA1WithRSA"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignatureSha256WithRsa   ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethod string

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethodHTTP  ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethod = "http"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethodCname ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethod = "cname"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethodTxt   ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultValidationMethod = "txt"
)

// Certificate's required verification information.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfo struct {
	// Name of CNAME record.
	RecordName ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTarget `json:"record_target"`
	JSON         zoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoJSON         `json:"-"`
}

// zoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoJSON
// contains the JSON metadata for the struct
// [ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfo]
type zoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Name of CNAME record.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordName string

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordNameRecordName ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordName = "record_name"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordNameHTTPURL    ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordName = "http_url"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordNameCname      ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordName = "cname"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordNameTxtName    ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordName = "txt_name"
)

// Target of CNAME record.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTarget string

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTargetRecordValue ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTarget = "record_value"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTargetHTTPBody    ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTarget = "http_body"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTargetCnameTarget ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTarget = "cname_target"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTargetTxtValue    ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationInfoRecordTarget = "txt_value"
)

// Method of verification.
type ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationType string

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationTypeCname   ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationType = "cname"
	ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationTypeMetaTag ZoneSslVerificationSslVerificationSslVerificationDetailsResponseResultVerificationType = "meta tag"
)

type ZoneSslVerificationUpdateParams struct {
	// Desired validation method.
	ValidationMethod param.Field[ZoneSslVerificationUpdateParamsValidationMethod] `json:"validation_method,required"`
}

func (r ZoneSslVerificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired validation method.
type ZoneSslVerificationUpdateParamsValidationMethod string

const (
	ZoneSslVerificationUpdateParamsValidationMethodHTTP  ZoneSslVerificationUpdateParamsValidationMethod = "http"
	ZoneSslVerificationUpdateParamsValidationMethodCname ZoneSslVerificationUpdateParamsValidationMethod = "cname"
	ZoneSslVerificationUpdateParamsValidationMethodTxt   ZoneSslVerificationUpdateParamsValidationMethod = "txt"
	ZoneSslVerificationUpdateParamsValidationMethodEmail ZoneSslVerificationUpdateParamsValidationMethod = "email"
)

type ZoneSslVerificationSslVerificationSslVerificationDetailsParams struct {
	// Immediately retry SSL Verification.
	Retry param.Field[ZoneSslVerificationSslVerificationSslVerificationDetailsParamsRetry] `query:"retry"`
}

// URLQuery serializes
// [ZoneSslVerificationSslVerificationSslVerificationDetailsParams]'s query
// parameters as `url.Values`.
func (r ZoneSslVerificationSslVerificationSslVerificationDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Immediately retry SSL Verification.
type ZoneSslVerificationSslVerificationSslVerificationDetailsParamsRetry bool

const (
	ZoneSslVerificationSslVerificationSslVerificationDetailsParamsRetryTrue ZoneSslVerificationSslVerificationSslVerificationDetailsParamsRetry = true
)
