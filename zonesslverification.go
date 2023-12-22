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
func (r *ZoneSslVerificationService) Update(ctx context.Context, zoneIdentifier string, certPackUuid string, body ZoneSslVerificationUpdateParams, opts ...option.RequestOption) (res *SslValidationMethodResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", zoneIdentifier, certPackUuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get SSL Verification Info for a Zone.
func (r *ZoneSslVerificationService) SslVerificationSslVerificationDetails(ctx context.Context, zoneIdentifier string, query ZoneSslVerificationSslVerificationSslVerificationDetailsParams, opts ...option.RequestOption) (res *SslVerificationResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/verification", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SslValidationMethodResponseCollection struct {
	Errors   []SslValidationMethodResponseCollectionError   `json:"errors"`
	Messages []SslValidationMethodResponseCollectionMessage `json:"messages"`
	Result   SslValidationMethodResponseCollectionResult    `json:"result"`
	// Whether the API call was successful
	Success SslValidationMethodResponseCollectionSuccess `json:"success"`
	JSON    sslValidationMethodResponseCollectionJSON    `json:"-"`
}

// sslValidationMethodResponseCollectionJSON contains the JSON metadata for the
// struct [SslValidationMethodResponseCollection]
type sslValidationMethodResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslValidationMethodResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslValidationMethodResponseCollectionError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    sslValidationMethodResponseCollectionErrorJSON `json:"-"`
}

// sslValidationMethodResponseCollectionErrorJSON contains the JSON metadata for
// the struct [SslValidationMethodResponseCollectionError]
type sslValidationMethodResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslValidationMethodResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslValidationMethodResponseCollectionMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    sslValidationMethodResponseCollectionMessageJSON `json:"-"`
}

// sslValidationMethodResponseCollectionMessageJSON contains the JSON metadata for
// the struct [SslValidationMethodResponseCollectionMessage]
type sslValidationMethodResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslValidationMethodResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslValidationMethodResponseCollectionResult struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod SslValidationMethodResponseCollectionResultValidationMethod `json:"validation_method"`
	JSON             sslValidationMethodResponseCollectionResultJSON             `json:"-"`
}

// sslValidationMethodResponseCollectionResultJSON contains the JSON metadata for
// the struct [SslValidationMethodResponseCollectionResult]
type sslValidationMethodResponseCollectionResultJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SslValidationMethodResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Desired validation method.
type SslValidationMethodResponseCollectionResultValidationMethod string

const (
	SslValidationMethodResponseCollectionResultValidationMethodHTTP  SslValidationMethodResponseCollectionResultValidationMethod = "http"
	SslValidationMethodResponseCollectionResultValidationMethodCname SslValidationMethodResponseCollectionResultValidationMethod = "cname"
	SslValidationMethodResponseCollectionResultValidationMethodTxt   SslValidationMethodResponseCollectionResultValidationMethod = "txt"
	SslValidationMethodResponseCollectionResultValidationMethodEmail SslValidationMethodResponseCollectionResultValidationMethod = "email"
)

// Whether the API call was successful
type SslValidationMethodResponseCollectionSuccess bool

const (
	SslValidationMethodResponseCollectionSuccessTrue SslValidationMethodResponseCollectionSuccess = true
)

type SslVerificationResponseCollection struct {
	Result []SslVerificationResponseCollectionResult `json:"result"`
	JSON   sslVerificationResponseCollectionJSON     `json:"-"`
}

// sslVerificationResponseCollectionJSON contains the JSON metadata for the struct
// [SslVerificationResponseCollection]
type sslVerificationResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslVerificationResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslVerificationResponseCollectionResult struct {
	// Current status of certificate.
	CertificateStatus SslVerificationResponseCollectionResultCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUuid string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature SslVerificationResponseCollectionResultSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod SslVerificationResponseCollectionResultValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo SslVerificationResponseCollectionResultVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType SslVerificationResponseCollectionResultVerificationType `json:"verification_type"`
	JSON             sslVerificationResponseCollectionResultJSON             `json:"-"`
}

// sslVerificationResponseCollectionResultJSON contains the JSON metadata for the
// struct [SslVerificationResponseCollectionResult]
type sslVerificationResponseCollectionResultJSON struct {
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

func (r *SslVerificationResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of certificate.
type SslVerificationResponseCollectionResultCertificateStatus string

const (
	SslVerificationResponseCollectionResultCertificateStatusInitializing      SslVerificationResponseCollectionResultCertificateStatus = "initializing"
	SslVerificationResponseCollectionResultCertificateStatusAuthorizing       SslVerificationResponseCollectionResultCertificateStatus = "authorizing"
	SslVerificationResponseCollectionResultCertificateStatusActive            SslVerificationResponseCollectionResultCertificateStatus = "active"
	SslVerificationResponseCollectionResultCertificateStatusExpired           SslVerificationResponseCollectionResultCertificateStatus = "expired"
	SslVerificationResponseCollectionResultCertificateStatusIssuing           SslVerificationResponseCollectionResultCertificateStatus = "issuing"
	SslVerificationResponseCollectionResultCertificateStatusTimingOut         SslVerificationResponseCollectionResultCertificateStatus = "timing_out"
	SslVerificationResponseCollectionResultCertificateStatusPendingDeployment SslVerificationResponseCollectionResultCertificateStatus = "pending_deployment"
)

// Certificate's signature algorithm.
type SslVerificationResponseCollectionResultSignature string

const (
	SslVerificationResponseCollectionResultSignatureEcdsaWithSha256 SslVerificationResponseCollectionResultSignature = "ECDSAWithSHA256"
	SslVerificationResponseCollectionResultSignatureSha1WithRsa     SslVerificationResponseCollectionResultSignature = "SHA1WithRSA"
	SslVerificationResponseCollectionResultSignatureSha256WithRsa   SslVerificationResponseCollectionResultSignature = "SHA256WithRSA"
)

// Validation method in use for a certificate pack order.
type SslVerificationResponseCollectionResultValidationMethod string

const (
	SslVerificationResponseCollectionResultValidationMethodHTTP  SslVerificationResponseCollectionResultValidationMethod = "http"
	SslVerificationResponseCollectionResultValidationMethodCname SslVerificationResponseCollectionResultValidationMethod = "cname"
	SslVerificationResponseCollectionResultValidationMethodTxt   SslVerificationResponseCollectionResultValidationMethod = "txt"
)

// Certificate's required verification information.
type SslVerificationResponseCollectionResultVerificationInfo struct {
	// Name of CNAME record.
	RecordName string `json:"record_name" format:"hostname"`
	// Target of CNAME record.
	RecordTarget string                                                      `json:"record_target" format:"hostname"`
	JSON         sslVerificationResponseCollectionResultVerificationInfoJSON `json:"-"`
}

// sslVerificationResponseCollectionResultVerificationInfoJSON contains the JSON
// metadata for the struct
// [SslVerificationResponseCollectionResultVerificationInfo]
type sslVerificationResponseCollectionResultVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SslVerificationResponseCollectionResultVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Method of verification.
type SslVerificationResponseCollectionResultVerificationType string

const (
	SslVerificationResponseCollectionResultVerificationTypeCname   SslVerificationResponseCollectionResultVerificationType = "cname"
	SslVerificationResponseCollectionResultVerificationTypeMetaTag SslVerificationResponseCollectionResultVerificationType = "meta tag"
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
