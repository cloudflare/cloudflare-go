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

// ZoneSslCertificatePackOrderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSslCertificatePackOrderService] method instead.
type ZoneSslCertificatePackOrderService struct {
	Options []option.RequestOption
}

// NewZoneSslCertificatePackOrderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSslCertificatePackOrderService(opts ...option.RequestOption) (r *ZoneSslCertificatePackOrderService) {
	r = &ZoneSslCertificatePackOrderService{}
	r.Options = opts
	return
}

// For a given zone, order an advanced certificate pack.
func (r *ZoneSslCertificatePackOrderService) CertificatePacksOrderAdvancedCertificateManagerCertificatePack(ctx context.Context, zoneIdentifier string, body ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams, opts ...option.RequestOption) (res *ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse struct {
	Errors   []ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseError   `json:"errors"`
	Messages []ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessage `json:"messages"`
	Result   ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseSuccess `json:"success"`
	JSON    zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseJSON    `json:"-"`
}

// zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse]
type zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseError struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseErrorJSON `json:"-"`
}

// zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseError]
type zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessage struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessageJSON `json:"-"`
}

// zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessage]
type zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus `json:"status"`
	// Type of certificate pack.
	Type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays `json:"validity_days"`
	JSON         zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultJSON         `json:"-"`
}

// zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResult]
type zoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultJSON struct {
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

func (r *ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultCertificateAuthority string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultCertificateAuthorityGoogle      ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultCertificateAuthority = "google"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultCertificateAuthorityLetsEncrypt ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusInitializing         ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "initializing"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusPendingValidation    ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "pending_validation"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusDeleted              ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "deleted"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusPendingIssuance      ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "pending_issuance"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusPendingDeployment    ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "pending_deployment"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusPendingDeletion      ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "pending_deletion"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusPendingExpiration    ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "pending_expiration"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusExpired              ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "expired"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusActive               ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "active"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusInitializingTimedOut ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "initializing_timed_out"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusValidationTimedOut   ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "validation_timed_out"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusIssuanceTimedOut     ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "issuance_timed_out"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusDeploymentTimedOut   ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "deployment_timed_out"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusDeletionTimedOut     ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "deletion_timed_out"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusPendingCleanup       ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "pending_cleanup"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusStagingDeployment    ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "staging_deployment"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusStagingActive        ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "staging_active"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusDeactivating         ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "deactivating"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusInactive             ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "inactive"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusBackupIssued         ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "backup_issued"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatusHoldingDeployment    ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultStatus = "holding_deployment"
)

// Type of certificate pack.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultType string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultTypeAdvanced ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultType = "advanced"
)

// Validation Method selected for the order.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethod string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethodTxt   ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethod = "txt"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethodHTTP  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethod = "http"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethodEmail ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidationMethod = "email"
)

// Validity Days selected for the order.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays int64

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays14  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays = 14
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays30  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays = 30
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays90  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays = 90
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays365 ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseResultValidityDays = 365
)

// Whether the API call was successful
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseSuccess bool

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseSuccessTrue ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackResponseSuccess = true
)

type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams struct {
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]string] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityGoogle      ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "google"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityLetsEncrypt ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "lets_encrypt"
)

// Type of certificate pack.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsTypeAdvanced ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType = "advanced"
)

// Validation Method selected for the order.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodTxt   ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "txt"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodHTTP  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "http"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodEmail ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "email"
)

// Validity Days selected for the order.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays int64

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays14  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 14
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays30  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 30
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays90  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 90
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays365 ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 365
)
