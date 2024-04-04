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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *VerificationService) Get(ctx context.Context, params VerificationGetParams, opts ...option.RequestOption) (res *[]TLSVerificationSetting, err error) {
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

type TLSVerificationSetting struct {
	// Current status of certificate.
	CertificateStatus TLSVerificationSettingCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CERTPackUUID string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature TLSVerificationSettingSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod TLSVerificationSettingValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo TLSVerificationSettingVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType TLSVerificationSettingVerificationType `json:"verification_type"`
	JSON             tlsVerificationSettingJSON             `json:"-"`
}

// tlsVerificationSettingJSON contains the JSON metadata for the struct
// [TLSVerificationSetting]
type tlsVerificationSettingJSON struct {
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

func (r *TLSVerificationSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsVerificationSettingJSON) RawJSON() string {
	return r.raw
}

// Current status of certificate.
type TLSVerificationSettingCertificateStatus string

const (
	TLSVerificationSettingCertificateStatusInitializing      TLSVerificationSettingCertificateStatus = "initializing"
	TLSVerificationSettingCertificateStatusAuthorizing       TLSVerificationSettingCertificateStatus = "authorizing"
	TLSVerificationSettingCertificateStatusActive            TLSVerificationSettingCertificateStatus = "active"
	TLSVerificationSettingCertificateStatusExpired           TLSVerificationSettingCertificateStatus = "expired"
	TLSVerificationSettingCertificateStatusIssuing           TLSVerificationSettingCertificateStatus = "issuing"
	TLSVerificationSettingCertificateStatusTimingOut         TLSVerificationSettingCertificateStatus = "timing_out"
	TLSVerificationSettingCertificateStatusPendingDeployment TLSVerificationSettingCertificateStatus = "pending_deployment"
)

func (r TLSVerificationSettingCertificateStatus) IsKnown() bool {
	switch r {
	case TLSVerificationSettingCertificateStatusInitializing, TLSVerificationSettingCertificateStatusAuthorizing, TLSVerificationSettingCertificateStatusActive, TLSVerificationSettingCertificateStatusExpired, TLSVerificationSettingCertificateStatusIssuing, TLSVerificationSettingCertificateStatusTimingOut, TLSVerificationSettingCertificateStatusPendingDeployment:
		return true
	}
	return false
}

// Certificate's signature algorithm.
type TLSVerificationSettingSignature string

const (
	TLSVerificationSettingSignatureEcdsaWithSha256 TLSVerificationSettingSignature = "ECDSAWithSHA256"
	TLSVerificationSettingSignatureSha1WithRsa     TLSVerificationSettingSignature = "SHA1WithRSA"
	TLSVerificationSettingSignatureSha256WithRsa   TLSVerificationSettingSignature = "SHA256WithRSA"
)

func (r TLSVerificationSettingSignature) IsKnown() bool {
	switch r {
	case TLSVerificationSettingSignatureEcdsaWithSha256, TLSVerificationSettingSignatureSha1WithRsa, TLSVerificationSettingSignatureSha256WithRsa:
		return true
	}
	return false
}

// Validation method in use for a certificate pack order.
type TLSVerificationSettingValidationMethod string

const (
	TLSVerificationSettingValidationMethodHTTP  TLSVerificationSettingValidationMethod = "http"
	TLSVerificationSettingValidationMethodCNAME TLSVerificationSettingValidationMethod = "cname"
	TLSVerificationSettingValidationMethodTXT   TLSVerificationSettingValidationMethod = "txt"
)

func (r TLSVerificationSettingValidationMethod) IsKnown() bool {
	switch r {
	case TLSVerificationSettingValidationMethodHTTP, TLSVerificationSettingValidationMethodCNAME, TLSVerificationSettingValidationMethodTXT:
		return true
	}
	return false
}

// Certificate's required verification information.
type TLSVerificationSettingVerificationInfo struct {
	// Name of CNAME record.
	RecordName TLSVerificationSettingVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget TLSVerificationSettingVerificationInfoRecordTarget `json:"record_target"`
	JSON         tlsVerificationSettingVerificationInfoJSON         `json:"-"`
}

// tlsVerificationSettingVerificationInfoJSON contains the JSON metadata for the
// struct [TLSVerificationSettingVerificationInfo]
type tlsVerificationSettingVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TLSVerificationSettingVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsVerificationSettingVerificationInfoJSON) RawJSON() string {
	return r.raw
}

// Name of CNAME record.
type TLSVerificationSettingVerificationInfoRecordName string

const (
	TLSVerificationSettingVerificationInfoRecordNameRecordName TLSVerificationSettingVerificationInfoRecordName = "record_name"
	TLSVerificationSettingVerificationInfoRecordNameHTTPURL    TLSVerificationSettingVerificationInfoRecordName = "http_url"
	TLSVerificationSettingVerificationInfoRecordNameCNAME      TLSVerificationSettingVerificationInfoRecordName = "cname"
	TLSVerificationSettingVerificationInfoRecordNameTXTName    TLSVerificationSettingVerificationInfoRecordName = "txt_name"
)

func (r TLSVerificationSettingVerificationInfoRecordName) IsKnown() bool {
	switch r {
	case TLSVerificationSettingVerificationInfoRecordNameRecordName, TLSVerificationSettingVerificationInfoRecordNameHTTPURL, TLSVerificationSettingVerificationInfoRecordNameCNAME, TLSVerificationSettingVerificationInfoRecordNameTXTName:
		return true
	}
	return false
}

// Target of CNAME record.
type TLSVerificationSettingVerificationInfoRecordTarget string

const (
	TLSVerificationSettingVerificationInfoRecordTargetRecordValue TLSVerificationSettingVerificationInfoRecordTarget = "record_value"
	TLSVerificationSettingVerificationInfoRecordTargetHTTPBody    TLSVerificationSettingVerificationInfoRecordTarget = "http_body"
	TLSVerificationSettingVerificationInfoRecordTargetCNAMETarget TLSVerificationSettingVerificationInfoRecordTarget = "cname_target"
	TLSVerificationSettingVerificationInfoRecordTargetTXTValue    TLSVerificationSettingVerificationInfoRecordTarget = "txt_value"
)

func (r TLSVerificationSettingVerificationInfoRecordTarget) IsKnown() bool {
	switch r {
	case TLSVerificationSettingVerificationInfoRecordTargetRecordValue, TLSVerificationSettingVerificationInfoRecordTargetHTTPBody, TLSVerificationSettingVerificationInfoRecordTargetCNAMETarget, TLSVerificationSettingVerificationInfoRecordTargetTXTValue:
		return true
	}
	return false
}

// Method of verification.
type TLSVerificationSettingVerificationType string

const (
	TLSVerificationSettingVerificationTypeCNAME   TLSVerificationSettingVerificationType = "cname"
	TLSVerificationSettingVerificationTypeMetaTag TLSVerificationSettingVerificationType = "meta tag"
)

func (r TLSVerificationSettingVerificationType) IsKnown() bool {
	switch r {
	case TLSVerificationSettingVerificationTypeCNAME, TLSVerificationSettingVerificationTypeMetaTag:
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
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   VerificationEditResponse `json:"result,required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
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
	Result []TLSVerificationSetting            `json:"result"`
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
