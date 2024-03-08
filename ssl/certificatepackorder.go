// File generated from our OpenAPI spec by Stainless.

package ssl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CertificatePackOrderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCertificatePackOrderService]
// method instead.
type CertificatePackOrderService struct {
	Options []option.RequestOption
}

// NewCertificatePackOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCertificatePackOrderService(opts ...option.RequestOption) (r *CertificatePackOrderService) {
	r = &CertificatePackOrderService{}
	r.Options = opts
	return
}

// For a given zone, order an advanced certificate pack.
func (r *CertificatePackOrderService) New(ctx context.Context, params CertificatePackOrderNewParams, opts ...option.RequestOption) (res *CertificatePackOrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackOrderNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificatePackOrderNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority CertificatePackOrderNewResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status CertificatePackOrderNewResponseStatus `json:"status"`
	// Type of certificate pack.
	Type CertificatePackOrderNewResponseType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod CertificatePackOrderNewResponseValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays CertificatePackOrderNewResponseValidityDays `json:"validity_days"`
	JSON         certificatePackOrderNewResponseJSON         `json:"-"`
}

// certificatePackOrderNewResponseJSON contains the JSON metadata for the struct
// [CertificatePackOrderNewResponse]
type certificatePackOrderNewResponseJSON struct {
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

func (r *CertificatePackOrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackOrderNewResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackOrderNewResponseCertificateAuthority string

const (
	CertificatePackOrderNewResponseCertificateAuthorityGoogle      CertificatePackOrderNewResponseCertificateAuthority = "google"
	CertificatePackOrderNewResponseCertificateAuthorityLetsEncrypt CertificatePackOrderNewResponseCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type CertificatePackOrderNewResponseStatus string

const (
	CertificatePackOrderNewResponseStatusInitializing         CertificatePackOrderNewResponseStatus = "initializing"
	CertificatePackOrderNewResponseStatusPendingValidation    CertificatePackOrderNewResponseStatus = "pending_validation"
	CertificatePackOrderNewResponseStatusDeleted              CertificatePackOrderNewResponseStatus = "deleted"
	CertificatePackOrderNewResponseStatusPendingIssuance      CertificatePackOrderNewResponseStatus = "pending_issuance"
	CertificatePackOrderNewResponseStatusPendingDeployment    CertificatePackOrderNewResponseStatus = "pending_deployment"
	CertificatePackOrderNewResponseStatusPendingDeletion      CertificatePackOrderNewResponseStatus = "pending_deletion"
	CertificatePackOrderNewResponseStatusPendingExpiration    CertificatePackOrderNewResponseStatus = "pending_expiration"
	CertificatePackOrderNewResponseStatusExpired              CertificatePackOrderNewResponseStatus = "expired"
	CertificatePackOrderNewResponseStatusActive               CertificatePackOrderNewResponseStatus = "active"
	CertificatePackOrderNewResponseStatusInitializingTimedOut CertificatePackOrderNewResponseStatus = "initializing_timed_out"
	CertificatePackOrderNewResponseStatusValidationTimedOut   CertificatePackOrderNewResponseStatus = "validation_timed_out"
	CertificatePackOrderNewResponseStatusIssuanceTimedOut     CertificatePackOrderNewResponseStatus = "issuance_timed_out"
	CertificatePackOrderNewResponseStatusDeploymentTimedOut   CertificatePackOrderNewResponseStatus = "deployment_timed_out"
	CertificatePackOrderNewResponseStatusDeletionTimedOut     CertificatePackOrderNewResponseStatus = "deletion_timed_out"
	CertificatePackOrderNewResponseStatusPendingCleanup       CertificatePackOrderNewResponseStatus = "pending_cleanup"
	CertificatePackOrderNewResponseStatusStagingDeployment    CertificatePackOrderNewResponseStatus = "staging_deployment"
	CertificatePackOrderNewResponseStatusStagingActive        CertificatePackOrderNewResponseStatus = "staging_active"
	CertificatePackOrderNewResponseStatusDeactivating         CertificatePackOrderNewResponseStatus = "deactivating"
	CertificatePackOrderNewResponseStatusInactive             CertificatePackOrderNewResponseStatus = "inactive"
	CertificatePackOrderNewResponseStatusBackupIssued         CertificatePackOrderNewResponseStatus = "backup_issued"
	CertificatePackOrderNewResponseStatusHoldingDeployment    CertificatePackOrderNewResponseStatus = "holding_deployment"
)

// Type of certificate pack.
type CertificatePackOrderNewResponseType string

const (
	CertificatePackOrderNewResponseTypeAdvanced CertificatePackOrderNewResponseType = "advanced"
)

// Validation Method selected for the order.
type CertificatePackOrderNewResponseValidationMethod string

const (
	CertificatePackOrderNewResponseValidationMethodTXT   CertificatePackOrderNewResponseValidationMethod = "txt"
	CertificatePackOrderNewResponseValidationMethodHTTP  CertificatePackOrderNewResponseValidationMethod = "http"
	CertificatePackOrderNewResponseValidationMethodEmail CertificatePackOrderNewResponseValidationMethod = "email"
)

// Validity Days selected for the order.
type CertificatePackOrderNewResponseValidityDays int64

const (
	CertificatePackOrderNewResponseValidityDays14  CertificatePackOrderNewResponseValidityDays = 14
	CertificatePackOrderNewResponseValidityDays30  CertificatePackOrderNewResponseValidityDays = 30
	CertificatePackOrderNewResponseValidityDays90  CertificatePackOrderNewResponseValidityDays = 90
	CertificatePackOrderNewResponseValidityDays365 CertificatePackOrderNewResponseValidityDays = 365
)

type CertificatePackOrderNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[CertificatePackOrderNewParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]string] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[CertificatePackOrderNewParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[CertificatePackOrderNewParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[CertificatePackOrderNewParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r CertificatePackOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackOrderNewParamsCertificateAuthority string

const (
	CertificatePackOrderNewParamsCertificateAuthorityGoogle      CertificatePackOrderNewParamsCertificateAuthority = "google"
	CertificatePackOrderNewParamsCertificateAuthorityLetsEncrypt CertificatePackOrderNewParamsCertificateAuthority = "lets_encrypt"
)

// Type of certificate pack.
type CertificatePackOrderNewParamsType string

const (
	CertificatePackOrderNewParamsTypeAdvanced CertificatePackOrderNewParamsType = "advanced"
)

// Validation Method selected for the order.
type CertificatePackOrderNewParamsValidationMethod string

const (
	CertificatePackOrderNewParamsValidationMethodTXT   CertificatePackOrderNewParamsValidationMethod = "txt"
	CertificatePackOrderNewParamsValidationMethodHTTP  CertificatePackOrderNewParamsValidationMethod = "http"
	CertificatePackOrderNewParamsValidationMethodEmail CertificatePackOrderNewParamsValidationMethod = "email"
)

// Validity Days selected for the order.
type CertificatePackOrderNewParamsValidityDays int64

const (
	CertificatePackOrderNewParamsValidityDays14  CertificatePackOrderNewParamsValidityDays = 14
	CertificatePackOrderNewParamsValidityDays30  CertificatePackOrderNewParamsValidityDays = 30
	CertificatePackOrderNewParamsValidityDays90  CertificatePackOrderNewParamsValidityDays = 90
	CertificatePackOrderNewParamsValidityDays365 CertificatePackOrderNewParamsValidityDays = 365
)

type CertificatePackOrderNewResponseEnvelope struct {
	Errors   []CertificatePackOrderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackOrderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificatePackOrderNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificatePackOrderNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificatePackOrderNewResponseEnvelopeJSON    `json:"-"`
}

// certificatePackOrderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackOrderNewResponseEnvelope]
type certificatePackOrderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackOrderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackOrderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackOrderNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    certificatePackOrderNewResponseEnvelopeErrorsJSON `json:"-"`
}

// certificatePackOrderNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CertificatePackOrderNewResponseEnvelopeErrors]
type certificatePackOrderNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackOrderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackOrderNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackOrderNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    certificatePackOrderNewResponseEnvelopeMessagesJSON `json:"-"`
}

// certificatePackOrderNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CertificatePackOrderNewResponseEnvelopeMessages]
type certificatePackOrderNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackOrderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackOrderNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackOrderNewResponseEnvelopeSuccess bool

const (
	CertificatePackOrderNewResponseEnvelopeSuccessTrue CertificatePackOrderNewResponseEnvelopeSuccess = true
)
