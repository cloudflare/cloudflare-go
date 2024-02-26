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
func (r *SSLCertificatePackOrderService) New(ctx context.Context, params SSLCertificatePackOrderNewParams, opts ...option.RequestOption) (res *SSLCertificatePackOrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackOrderNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLCertificatePackOrderNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority SSLCertificatePackOrderNewResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status SSLCertificatePackOrderNewResponseStatus `json:"status"`
	// Type of certificate pack.
	Type SSLCertificatePackOrderNewResponseType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod SSLCertificatePackOrderNewResponseValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays SSLCertificatePackOrderNewResponseValidityDays `json:"validity_days"`
	JSON         sslCertificatePackOrderNewResponseJSON         `json:"-"`
}

// sslCertificatePackOrderNewResponseJSON contains the JSON metadata for the struct
// [SSLCertificatePackOrderNewResponse]
type sslCertificatePackOrderNewResponseJSON struct {
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

func (r *SSLCertificatePackOrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type SSLCertificatePackOrderNewResponseCertificateAuthority string

const (
	SSLCertificatePackOrderNewResponseCertificateAuthorityGoogle      SSLCertificatePackOrderNewResponseCertificateAuthority = "google"
	SSLCertificatePackOrderNewResponseCertificateAuthorityLetsEncrypt SSLCertificatePackOrderNewResponseCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type SSLCertificatePackOrderNewResponseStatus string

const (
	SSLCertificatePackOrderNewResponseStatusInitializing         SSLCertificatePackOrderNewResponseStatus = "initializing"
	SSLCertificatePackOrderNewResponseStatusPendingValidation    SSLCertificatePackOrderNewResponseStatus = "pending_validation"
	SSLCertificatePackOrderNewResponseStatusDeleted              SSLCertificatePackOrderNewResponseStatus = "deleted"
	SSLCertificatePackOrderNewResponseStatusPendingIssuance      SSLCertificatePackOrderNewResponseStatus = "pending_issuance"
	SSLCertificatePackOrderNewResponseStatusPendingDeployment    SSLCertificatePackOrderNewResponseStatus = "pending_deployment"
	SSLCertificatePackOrderNewResponseStatusPendingDeletion      SSLCertificatePackOrderNewResponseStatus = "pending_deletion"
	SSLCertificatePackOrderNewResponseStatusPendingExpiration    SSLCertificatePackOrderNewResponseStatus = "pending_expiration"
	SSLCertificatePackOrderNewResponseStatusExpired              SSLCertificatePackOrderNewResponseStatus = "expired"
	SSLCertificatePackOrderNewResponseStatusActive               SSLCertificatePackOrderNewResponseStatus = "active"
	SSLCertificatePackOrderNewResponseStatusInitializingTimedOut SSLCertificatePackOrderNewResponseStatus = "initializing_timed_out"
	SSLCertificatePackOrderNewResponseStatusValidationTimedOut   SSLCertificatePackOrderNewResponseStatus = "validation_timed_out"
	SSLCertificatePackOrderNewResponseStatusIssuanceTimedOut     SSLCertificatePackOrderNewResponseStatus = "issuance_timed_out"
	SSLCertificatePackOrderNewResponseStatusDeploymentTimedOut   SSLCertificatePackOrderNewResponseStatus = "deployment_timed_out"
	SSLCertificatePackOrderNewResponseStatusDeletionTimedOut     SSLCertificatePackOrderNewResponseStatus = "deletion_timed_out"
	SSLCertificatePackOrderNewResponseStatusPendingCleanup       SSLCertificatePackOrderNewResponseStatus = "pending_cleanup"
	SSLCertificatePackOrderNewResponseStatusStagingDeployment    SSLCertificatePackOrderNewResponseStatus = "staging_deployment"
	SSLCertificatePackOrderNewResponseStatusStagingActive        SSLCertificatePackOrderNewResponseStatus = "staging_active"
	SSLCertificatePackOrderNewResponseStatusDeactivating         SSLCertificatePackOrderNewResponseStatus = "deactivating"
	SSLCertificatePackOrderNewResponseStatusInactive             SSLCertificatePackOrderNewResponseStatus = "inactive"
	SSLCertificatePackOrderNewResponseStatusBackupIssued         SSLCertificatePackOrderNewResponseStatus = "backup_issued"
	SSLCertificatePackOrderNewResponseStatusHoldingDeployment    SSLCertificatePackOrderNewResponseStatus = "holding_deployment"
)

// Type of certificate pack.
type SSLCertificatePackOrderNewResponseType string

const (
	SSLCertificatePackOrderNewResponseTypeAdvanced SSLCertificatePackOrderNewResponseType = "advanced"
)

// Validation Method selected for the order.
type SSLCertificatePackOrderNewResponseValidationMethod string

const (
	SSLCertificatePackOrderNewResponseValidationMethodTxt   SSLCertificatePackOrderNewResponseValidationMethod = "txt"
	SSLCertificatePackOrderNewResponseValidationMethodHTTP  SSLCertificatePackOrderNewResponseValidationMethod = "http"
	SSLCertificatePackOrderNewResponseValidationMethodEmail SSLCertificatePackOrderNewResponseValidationMethod = "email"
)

// Validity Days selected for the order.
type SSLCertificatePackOrderNewResponseValidityDays int64

const (
	SSLCertificatePackOrderNewResponseValidityDays14  SSLCertificatePackOrderNewResponseValidityDays = 14
	SSLCertificatePackOrderNewResponseValidityDays30  SSLCertificatePackOrderNewResponseValidityDays = 30
	SSLCertificatePackOrderNewResponseValidityDays90  SSLCertificatePackOrderNewResponseValidityDays = 90
	SSLCertificatePackOrderNewResponseValidityDays365 SSLCertificatePackOrderNewResponseValidityDays = 365
)

type SSLCertificatePackOrderNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[SSLCertificatePackOrderNewParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]string] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[SSLCertificatePackOrderNewParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[SSLCertificatePackOrderNewParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[SSLCertificatePackOrderNewParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r SSLCertificatePackOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type SSLCertificatePackOrderNewParamsCertificateAuthority string

const (
	SSLCertificatePackOrderNewParamsCertificateAuthorityGoogle      SSLCertificatePackOrderNewParamsCertificateAuthority = "google"
	SSLCertificatePackOrderNewParamsCertificateAuthorityLetsEncrypt SSLCertificatePackOrderNewParamsCertificateAuthority = "lets_encrypt"
)

// Type of certificate pack.
type SSLCertificatePackOrderNewParamsType string

const (
	SSLCertificatePackOrderNewParamsTypeAdvanced SSLCertificatePackOrderNewParamsType = "advanced"
)

// Validation Method selected for the order.
type SSLCertificatePackOrderNewParamsValidationMethod string

const (
	SSLCertificatePackOrderNewParamsValidationMethodTxt   SSLCertificatePackOrderNewParamsValidationMethod = "txt"
	SSLCertificatePackOrderNewParamsValidationMethodHTTP  SSLCertificatePackOrderNewParamsValidationMethod = "http"
	SSLCertificatePackOrderNewParamsValidationMethodEmail SSLCertificatePackOrderNewParamsValidationMethod = "email"
)

// Validity Days selected for the order.
type SSLCertificatePackOrderNewParamsValidityDays int64

const (
	SSLCertificatePackOrderNewParamsValidityDays14  SSLCertificatePackOrderNewParamsValidityDays = 14
	SSLCertificatePackOrderNewParamsValidityDays30  SSLCertificatePackOrderNewParamsValidityDays = 30
	SSLCertificatePackOrderNewParamsValidityDays90  SSLCertificatePackOrderNewParamsValidityDays = 90
	SSLCertificatePackOrderNewParamsValidityDays365 SSLCertificatePackOrderNewParamsValidityDays = 365
)

type SSLCertificatePackOrderNewResponseEnvelope struct {
	Errors   []SSLCertificatePackOrderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackOrderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackOrderNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackOrderNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackOrderNewResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackOrderNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [SSLCertificatePackOrderNewResponseEnvelope]
type sslCertificatePackOrderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackOrderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackOrderNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    sslCertificatePackOrderNewResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackOrderNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SSLCertificatePackOrderNewResponseEnvelopeErrors]
type sslCertificatePackOrderNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackOrderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackOrderNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    sslCertificatePackOrderNewResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackOrderNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SSLCertificatePackOrderNewResponseEnvelopeMessages]
type sslCertificatePackOrderNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackOrderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackOrderNewResponseEnvelopeSuccess bool

const (
	SSLCertificatePackOrderNewResponseEnvelopeSuccessTrue SSLCertificatePackOrderNewResponseEnvelopeSuccess = true
)
