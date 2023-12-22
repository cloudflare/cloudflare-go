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

// ZoneSslCertificatePackService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSslCertificatePackService]
// method instead.
type ZoneSslCertificatePackService struct {
	Options []option.RequestOption
	Orders  *ZoneSslCertificatePackOrderService
	Quotas  *ZoneSslCertificatePackQuotaService
}

// NewZoneSslCertificatePackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslCertificatePackService(opts ...option.RequestOption) (r *ZoneSslCertificatePackService) {
	r = &ZoneSslCertificatePackService{}
	r.Options = opts
	r.Orders = NewZoneSslCertificatePackOrderService(opts...)
	r.Quotas = NewZoneSslCertificatePackQuotaService(opts...)
	return
}

// For a given zone, get a certificate pack.
func (r *ZoneSslCertificatePackService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *CertificatePackResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For a given zone, restart validation for an advanced certificate pack. This is
// only a validation operation for a Certificate Pack in a validation_timed_out
// status.
func (r *ZoneSslCertificatePackService) Update(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *AdvancedCertificatePackResponseSingleQ8qWa5qi, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// For a given zone, delete an advanced certificate pack.
func (r *ZoneSslCertificatePackService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DeleteAdvancedCertificatePackResponseSingleNs3TjUge, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// For a given zone, list all active certificate packs.
func (r *ZoneSslCertificatePackService) CertificatePacksListCertificatePacks(ctx context.Context, zoneIdentifier string, query ZoneSslCertificatePackCertificatePacksListCertificatePacksParams, opts ...option.RequestOption) (res *CertificatePackResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdvancedCertificatePackResponseSingleQ8qWa5qi struct {
	Errors   []AdvancedCertificatePackResponseSingleQ8qWa5qiError   `json:"errors"`
	Messages []AdvancedCertificatePackResponseSingleQ8qWa5qiMessage `json:"messages"`
	Result   AdvancedCertificatePackResponseSingleQ8qWa5qiResult    `json:"result"`
	// Whether the API call was successful
	Success AdvancedCertificatePackResponseSingleQ8qWa5qiSuccess `json:"success"`
	JSON    advancedCertificatePackResponseSingleQ8qWa5qiJSON    `json:"-"`
}

// advancedCertificatePackResponseSingleQ8qWa5qiJSON contains the JSON metadata for
// the struct [AdvancedCertificatePackResponseSingleQ8qWa5qi]
type advancedCertificatePackResponseSingleQ8qWa5qiJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedCertificatePackResponseSingleQ8qWa5qi) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedCertificatePackResponseSingleQ8qWa5qiError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    advancedCertificatePackResponseSingleQ8qWa5qiErrorJSON `json:"-"`
}

// advancedCertificatePackResponseSingleQ8qWa5qiErrorJSON contains the JSON
// metadata for the struct [AdvancedCertificatePackResponseSingleQ8qWa5qiError]
type advancedCertificatePackResponseSingleQ8qWa5qiErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedCertificatePackResponseSingleQ8qWa5qiError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedCertificatePackResponseSingleQ8qWa5qiMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    advancedCertificatePackResponseSingleQ8qWa5qiMessageJSON `json:"-"`
}

// advancedCertificatePackResponseSingleQ8qWa5qiMessageJSON contains the JSON
// metadata for the struct [AdvancedCertificatePackResponseSingleQ8qWa5qiMessage]
type advancedCertificatePackResponseSingleQ8qWa5qiMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedCertificatePackResponseSingleQ8qWa5qiMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedCertificatePackResponseSingleQ8qWa5qiResult struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. Selecting Let's Encrypt will
	// reduce customization of other fields: validation_method must be 'txt',
	// validity_days must be 90, cloudflare_branding must be omitted, and hosts must
	// contain only 2 entries, one for the zone name and one for the subdomain wildcard
	// of the zone name (e.g. example.com, \*.example.com).
	CertificateAuthority AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus `json:"status"`
	// Type of certificate pack.
	Type AdvancedCertificatePackResponseSingleQ8qWa5qiResultType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays `json:"validity_days"`
	JSON         advancedCertificatePackResponseSingleQ8qWa5qiResultJSON         `json:"-"`
}

// advancedCertificatePackResponseSingleQ8qWa5qiResultJSON contains the JSON
// metadata for the struct [AdvancedCertificatePackResponseSingleQ8qWa5qiResult]
type advancedCertificatePackResponseSingleQ8qWa5qiResultJSON struct {
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

func (r *AdvancedCertificatePackResponseSingleQ8qWa5qiResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority selected for the order. Selecting Let's Encrypt will
// reduce customization of other fields: validation_method must be 'txt',
// validity_days must be 90, cloudflare_branding must be omitted, and hosts must
// contain only 2 entries, one for the zone name and one for the subdomain wildcard
// of the zone name (e.g. example.com, \*.example.com).
type AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthority string

const (
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthorityDigicert    AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthority = "digicert"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthorityGoogle      AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthority = "google"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthorityLetsEncrypt AdvancedCertificatePackResponseSingleQ8qWa5qiResultCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus string

const (
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusInitializing         AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "initializing"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusPendingValidation    AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "pending_validation"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusDeleted              AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "deleted"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusPendingIssuance      AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "pending_issuance"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusPendingDeployment    AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "pending_deployment"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusPendingDeletion      AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "pending_deletion"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusPendingExpiration    AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "pending_expiration"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusExpired              AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "expired"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusActive               AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "active"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusInitializingTimedOut AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "initializing_timed_out"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusValidationTimedOut   AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "validation_timed_out"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusIssuanceTimedOut     AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "issuance_timed_out"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusDeploymentTimedOut   AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "deployment_timed_out"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusDeletionTimedOut     AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "deletion_timed_out"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusPendingCleanup       AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "pending_cleanup"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusStagingDeployment    AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "staging_deployment"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusStagingActive        AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "staging_active"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusDeactivating         AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "deactivating"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusInactive             AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "inactive"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusBackupIssued         AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "backup_issued"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatusHoldingDeployment    AdvancedCertificatePackResponseSingleQ8qWa5qiResultStatus = "holding_deployment"
)

// Type of certificate pack.
type AdvancedCertificatePackResponseSingleQ8qWa5qiResultType string

const (
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultTypeAdvanced AdvancedCertificatePackResponseSingleQ8qWa5qiResultType = "advanced"
)

// Validation Method selected for the order.
type AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethod string

const (
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethodTxt   AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethod = "txt"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethodHTTP  AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethod = "http"
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethodEmail AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidationMethod = "email"
)

// Validity Days selected for the order.
type AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays int64

const (
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays14  AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays = 14
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays30  AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays = 30
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays90  AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays = 90
	AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays365 AdvancedCertificatePackResponseSingleQ8qWa5qiResultValidityDays = 365
)

// Whether the API call was successful
type AdvancedCertificatePackResponseSingleQ8qWa5qiSuccess bool

const (
	AdvancedCertificatePackResponseSingleQ8qWa5qiSuccessTrue AdvancedCertificatePackResponseSingleQ8qWa5qiSuccess = true
)

type CertificatePackResponseCollection struct {
	Errors     []CertificatePackResponseCollectionError    `json:"errors"`
	Messages   []CertificatePackResponseCollectionMessage  `json:"messages"`
	Result     []interface{}                               `json:"result"`
	ResultInfo CertificatePackResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CertificatePackResponseCollectionSuccess `json:"success"`
	JSON    certificatePackResponseCollectionJSON    `json:"-"`
}

// certificatePackResponseCollectionJSON contains the JSON metadata for the struct
// [CertificatePackResponseCollection]
type certificatePackResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackResponseCollectionError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    certificatePackResponseCollectionErrorJSON `json:"-"`
}

// certificatePackResponseCollectionErrorJSON contains the JSON metadata for the
// struct [CertificatePackResponseCollectionError]
type certificatePackResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackResponseCollectionMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    certificatePackResponseCollectionMessageJSON `json:"-"`
}

// certificatePackResponseCollectionMessageJSON contains the JSON metadata for the
// struct [CertificatePackResponseCollectionMessage]
type certificatePackResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       certificatePackResponseCollectionResultInfoJSON `json:"-"`
}

// certificatePackResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [CertificatePackResponseCollectionResultInfo]
type certificatePackResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificatePackResponseCollectionSuccess bool

const (
	CertificatePackResponseCollectionSuccessTrue CertificatePackResponseCollectionSuccess = true
)

type CertificatePackResponseSingle struct {
	Errors   []CertificatePackResponseSingleError   `json:"errors"`
	Messages []CertificatePackResponseSingleMessage `json:"messages"`
	Result   interface{}                            `json:"result"`
	// Whether the API call was successful
	Success CertificatePackResponseSingleSuccess `json:"success"`
	JSON    certificatePackResponseSingleJSON    `json:"-"`
}

// certificatePackResponseSingleJSON contains the JSON metadata for the struct
// [CertificatePackResponseSingle]
type certificatePackResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackResponseSingleError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    certificatePackResponseSingleErrorJSON `json:"-"`
}

// certificatePackResponseSingleErrorJSON contains the JSON metadata for the struct
// [CertificatePackResponseSingleError]
type certificatePackResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackResponseSingleMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    certificatePackResponseSingleMessageJSON `json:"-"`
}

// certificatePackResponseSingleMessageJSON contains the JSON metadata for the
// struct [CertificatePackResponseSingleMessage]
type certificatePackResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificatePackResponseSingleSuccess bool

const (
	CertificatePackResponseSingleSuccessTrue CertificatePackResponseSingleSuccess = true
)

type DeleteAdvancedCertificatePackResponseSingleNs3TjUge struct {
	Errors   []DeleteAdvancedCertificatePackResponseSingleNs3TjUgeError   `json:"errors"`
	Messages []DeleteAdvancedCertificatePackResponseSingleNs3TjUgeMessage `json:"messages"`
	Result   DeleteAdvancedCertificatePackResponseSingleNs3TjUgeResult    `json:"result"`
	// Whether the API call was successful
	Success DeleteAdvancedCertificatePackResponseSingleNs3TjUgeSuccess `json:"success"`
	JSON    deleteAdvancedCertificatePackResponseSingleNs3TjUgeJSON    `json:"-"`
}

// deleteAdvancedCertificatePackResponseSingleNs3TjUgeJSON contains the JSON
// metadata for the struct [DeleteAdvancedCertificatePackResponseSingleNs3TjUge]
type deleteAdvancedCertificatePackResponseSingleNs3TjUgeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeleteAdvancedCertificatePackResponseSingleNs3TjUge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteAdvancedCertificatePackResponseSingleNs3TjUgeError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    deleteAdvancedCertificatePackResponseSingleNs3TjUgeErrorJSON `json:"-"`
}

// deleteAdvancedCertificatePackResponseSingleNs3TjUgeErrorJSON contains the JSON
// metadata for the struct
// [DeleteAdvancedCertificatePackResponseSingleNs3TjUgeError]
type deleteAdvancedCertificatePackResponseSingleNs3TjUgeErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeleteAdvancedCertificatePackResponseSingleNs3TjUgeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteAdvancedCertificatePackResponseSingleNs3TjUgeMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    deleteAdvancedCertificatePackResponseSingleNs3TjUgeMessageJSON `json:"-"`
}

// deleteAdvancedCertificatePackResponseSingleNs3TjUgeMessageJSON contains the JSON
// metadata for the struct
// [DeleteAdvancedCertificatePackResponseSingleNs3TjUgeMessage]
type deleteAdvancedCertificatePackResponseSingleNs3TjUgeMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeleteAdvancedCertificatePackResponseSingleNs3TjUgeMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteAdvancedCertificatePackResponseSingleNs3TjUgeResult struct {
	// Identifier
	ID   string                                                        `json:"id"`
	JSON deleteAdvancedCertificatePackResponseSingleNs3TjUgeResultJSON `json:"-"`
}

// deleteAdvancedCertificatePackResponseSingleNs3TjUgeResultJSON contains the JSON
// metadata for the struct
// [DeleteAdvancedCertificatePackResponseSingleNs3TjUgeResult]
type deleteAdvancedCertificatePackResponseSingleNs3TjUgeResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeleteAdvancedCertificatePackResponseSingleNs3TjUgeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeleteAdvancedCertificatePackResponseSingleNs3TjUgeSuccess bool

const (
	DeleteAdvancedCertificatePackResponseSingleNs3TjUgeSuccessTrue DeleteAdvancedCertificatePackResponseSingleNs3TjUgeSuccess = true
)

type ZoneSslCertificatePackCertificatePacksListCertificatePacksParams struct {
	// Include Certificate Packs of all statuses, not just active ones.
	Status param.Field[ZoneSslCertificatePackCertificatePacksListCertificatePacksParamsStatus] `query:"status"`
}

// URLQuery serializes
// [ZoneSslCertificatePackCertificatePacksListCertificatePacksParams]'s query
// parameters as `url.Values`.
func (r ZoneSslCertificatePackCertificatePacksListCertificatePacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include Certificate Packs of all statuses, not just active ones.
type ZoneSslCertificatePackCertificatePacksListCertificatePacksParamsStatus string

const (
	ZoneSslCertificatePackCertificatePacksListCertificatePacksParamsStatusAll ZoneSslCertificatePackCertificatePacksListCertificatePacksParamsStatus = "all"
)
