// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SSLCertificatePackOrderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSSLCertificatePackOrderService] method instead.
type SSLCertificatePackOrderService struct {
	Options []option.RequestOption
}

// NewSSLCertificatePackOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSSLCertificatePackOrderService(opts ...option.RequestOption) (r *SSLCertificatePackOrderService) {
	r = &SSLCertificatePackOrderService{}
	r.Options = opts
	return
}

// For a given zone, order an advanced certificate pack.
func (r *SSLCertificatePackOrderService) CertificatePacksOrderAdvancedCertificateManagerCertificatePack(ctx context.Context, zoneID string, body SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams, opts ...option.RequestOption) (res *SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus `json:"status"`
	// Type of certificate pack.
	Type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays `json:"validity_days"`
	JSON         sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseJSON         `json:"-"`
}

// sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse]
type sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CloudflareBranding   apijson.Field
	Hosts                apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	ValidationMethod     apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseCertificateAuthority string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseCertificateAuthorityGoogle      SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseCertificateAuthority = "google"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseCertificateAuthorityLetsEncrypt SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusInitializing         SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "initializing"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusPendingValidation    SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "pending_validation"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusDeleted              SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "deleted"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusPendingIssuance      SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "pending_issuance"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusPendingDeployment    SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "pending_deployment"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusPendingDeletion      SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "pending_deletion"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusPendingExpiration    SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "pending_expiration"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusExpired              SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "expired"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusActive               SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "active"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusInitializingTimedOut SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "initializing_timed_out"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusValidationTimedOut   SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "validation_timed_out"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusIssuanceTimedOut     SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "issuance_timed_out"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusDeploymentTimedOut   SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "deployment_timed_out"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusDeletionTimedOut     SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "deletion_timed_out"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusPendingCleanup       SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "pending_cleanup"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusStagingDeployment    SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "staging_deployment"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusStagingActive        SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "staging_active"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusDeactivating         SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "deactivating"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusInactive             SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "inactive"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusBackupIssued         SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "backup_issued"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatusHoldingDeployment    SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseStatus = "holding_deployment"
)

// Type of certificate pack.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseType string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseTypeAdvanced SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseType = "advanced"
)

// Validation Method selected for the order.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethod string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethodTxt   SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethod = "txt"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethodHTTP  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethod = "http"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethodEmail SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidationMethod = "email"
)

// Validity Days selected for the order.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays int64

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays14  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays = 14
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays30  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays = 30
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays90  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays = 90
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays365 SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseValidityDays = 365
)

type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams struct {
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]string] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityGoogle      SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "google"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityLetsEncrypt SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "lets_encrypt"
)

// Type of certificate pack.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsTypeAdvanced SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType = "advanced"
)

// Validation Method selected for the order.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod string

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodTxt   SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "txt"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodHTTP  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "http"
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodEmail SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "email"
)

// Validity Days selected for the order.
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays int64

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays14  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 14
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays30  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 30
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays90  SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 90
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays365 SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 365
)

type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelope struct {
	Errors   []SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelope]
type sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrors struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrors]
type sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessages struct {
	Code    int64                                                                                                             `json:"code,required"`
	Message string                                                                                                            `json:"message,required"`
	JSON    sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessages]
type sslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeSuccess bool

const (
	SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeSuccessTrue SSLCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseEnvelopeSuccess = true
)
