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
func (r *ZoneSslCertificatePackService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneSslCertificatePackGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For a given zone, restart validation for an advanced certificate pack. This is
// only a validation operation for a Certificate Pack in a validation_timed_out
// status.
func (r *ZoneSslCertificatePackService) Update(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneSslCertificatePackUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// For a given zone, delete an advanced certificate pack.
func (r *ZoneSslCertificatePackService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneSslCertificatePackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// For a given zone, list all active certificate packs.
func (r *ZoneSslCertificatePackService) CertificatePacksListCertificatePacks(ctx context.Context, zoneIdentifier string, query ZoneSslCertificatePackCertificatePacksListCertificatePacksParams, opts ...option.RequestOption) (res *ZoneSslCertificatePackCertificatePacksListCertificatePacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSslCertificatePackGetResponse struct {
	Errors   []ZoneSslCertificatePackGetResponseError   `json:"errors"`
	Messages []ZoneSslCertificatePackGetResponseMessage `json:"messages"`
	Result   interface{}                                `json:"result"`
	// Whether the API call was successful
	Success ZoneSslCertificatePackGetResponseSuccess `json:"success"`
	JSON    zoneSslCertificatePackGetResponseJSON    `json:"-"`
}

// zoneSslCertificatePackGetResponseJSON contains the JSON metadata for the struct
// [ZoneSslCertificatePackGetResponse]
type zoneSslCertificatePackGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSslCertificatePackGetResponseErrorJSON `json:"-"`
}

// zoneSslCertificatePackGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackGetResponseError]
type zoneSslCertificatePackGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSslCertificatePackGetResponseMessageJSON `json:"-"`
}

// zoneSslCertificatePackGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackGetResponseMessage]
type zoneSslCertificatePackGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSslCertificatePackGetResponseSuccess bool

const (
	ZoneSslCertificatePackGetResponseSuccessTrue ZoneSslCertificatePackGetResponseSuccess = true
)

type ZoneSslCertificatePackUpdateResponse struct {
	Errors   []ZoneSslCertificatePackUpdateResponseError   `json:"errors"`
	Messages []ZoneSslCertificatePackUpdateResponseMessage `json:"messages"`
	Result   ZoneSslCertificatePackUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslCertificatePackUpdateResponseSuccess `json:"success"`
	JSON    zoneSslCertificatePackUpdateResponseJSON    `json:"-"`
}

// zoneSslCertificatePackUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackUpdateResponse]
type zoneSslCertificatePackUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSslCertificatePackUpdateResponseErrorJSON `json:"-"`
}

// zoneSslCertificatePackUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackUpdateResponseError]
type zoneSslCertificatePackUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSslCertificatePackUpdateResponseMessageJSON `json:"-"`
}

// zoneSslCertificatePackUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSslCertificatePackUpdateResponseMessage]
type zoneSslCertificatePackUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackUpdateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority ZoneSslCertificatePackUpdateResponseResultCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status ZoneSslCertificatePackUpdateResponseResultStatus `json:"status"`
	// Type of certificate pack.
	Type ZoneSslCertificatePackUpdateResponseResultType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod ZoneSslCertificatePackUpdateResponseResultValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays ZoneSslCertificatePackUpdateResponseResultValidityDays `json:"validity_days"`
	JSON         zoneSslCertificatePackUpdateResponseResultJSON         `json:"-"`
}

// zoneSslCertificatePackUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSslCertificatePackUpdateResponseResult]
type zoneSslCertificatePackUpdateResponseResultJSON struct {
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

func (r *ZoneSslCertificatePackUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type ZoneSslCertificatePackUpdateResponseResultCertificateAuthority string

const (
	ZoneSslCertificatePackUpdateResponseResultCertificateAuthorityGoogle      ZoneSslCertificatePackUpdateResponseResultCertificateAuthority = "google"
	ZoneSslCertificatePackUpdateResponseResultCertificateAuthorityLetsEncrypt ZoneSslCertificatePackUpdateResponseResultCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type ZoneSslCertificatePackUpdateResponseResultStatus string

const (
	ZoneSslCertificatePackUpdateResponseResultStatusInitializing         ZoneSslCertificatePackUpdateResponseResultStatus = "initializing"
	ZoneSslCertificatePackUpdateResponseResultStatusPendingValidation    ZoneSslCertificatePackUpdateResponseResultStatus = "pending_validation"
	ZoneSslCertificatePackUpdateResponseResultStatusDeleted              ZoneSslCertificatePackUpdateResponseResultStatus = "deleted"
	ZoneSslCertificatePackUpdateResponseResultStatusPendingIssuance      ZoneSslCertificatePackUpdateResponseResultStatus = "pending_issuance"
	ZoneSslCertificatePackUpdateResponseResultStatusPendingDeployment    ZoneSslCertificatePackUpdateResponseResultStatus = "pending_deployment"
	ZoneSslCertificatePackUpdateResponseResultStatusPendingDeletion      ZoneSslCertificatePackUpdateResponseResultStatus = "pending_deletion"
	ZoneSslCertificatePackUpdateResponseResultStatusPendingExpiration    ZoneSslCertificatePackUpdateResponseResultStatus = "pending_expiration"
	ZoneSslCertificatePackUpdateResponseResultStatusExpired              ZoneSslCertificatePackUpdateResponseResultStatus = "expired"
	ZoneSslCertificatePackUpdateResponseResultStatusActive               ZoneSslCertificatePackUpdateResponseResultStatus = "active"
	ZoneSslCertificatePackUpdateResponseResultStatusInitializingTimedOut ZoneSslCertificatePackUpdateResponseResultStatus = "initializing_timed_out"
	ZoneSslCertificatePackUpdateResponseResultStatusValidationTimedOut   ZoneSslCertificatePackUpdateResponseResultStatus = "validation_timed_out"
	ZoneSslCertificatePackUpdateResponseResultStatusIssuanceTimedOut     ZoneSslCertificatePackUpdateResponseResultStatus = "issuance_timed_out"
	ZoneSslCertificatePackUpdateResponseResultStatusDeploymentTimedOut   ZoneSslCertificatePackUpdateResponseResultStatus = "deployment_timed_out"
	ZoneSslCertificatePackUpdateResponseResultStatusDeletionTimedOut     ZoneSslCertificatePackUpdateResponseResultStatus = "deletion_timed_out"
	ZoneSslCertificatePackUpdateResponseResultStatusPendingCleanup       ZoneSslCertificatePackUpdateResponseResultStatus = "pending_cleanup"
	ZoneSslCertificatePackUpdateResponseResultStatusStagingDeployment    ZoneSslCertificatePackUpdateResponseResultStatus = "staging_deployment"
	ZoneSslCertificatePackUpdateResponseResultStatusStagingActive        ZoneSslCertificatePackUpdateResponseResultStatus = "staging_active"
	ZoneSslCertificatePackUpdateResponseResultStatusDeactivating         ZoneSslCertificatePackUpdateResponseResultStatus = "deactivating"
	ZoneSslCertificatePackUpdateResponseResultStatusInactive             ZoneSslCertificatePackUpdateResponseResultStatus = "inactive"
	ZoneSslCertificatePackUpdateResponseResultStatusBackupIssued         ZoneSslCertificatePackUpdateResponseResultStatus = "backup_issued"
	ZoneSslCertificatePackUpdateResponseResultStatusHoldingDeployment    ZoneSslCertificatePackUpdateResponseResultStatus = "holding_deployment"
)

// Type of certificate pack.
type ZoneSslCertificatePackUpdateResponseResultType string

const (
	ZoneSslCertificatePackUpdateResponseResultTypeAdvanced ZoneSslCertificatePackUpdateResponseResultType = "advanced"
)

// Validation Method selected for the order.
type ZoneSslCertificatePackUpdateResponseResultValidationMethod string

const (
	ZoneSslCertificatePackUpdateResponseResultValidationMethodTxt   ZoneSslCertificatePackUpdateResponseResultValidationMethod = "txt"
	ZoneSslCertificatePackUpdateResponseResultValidationMethodHTTP  ZoneSslCertificatePackUpdateResponseResultValidationMethod = "http"
	ZoneSslCertificatePackUpdateResponseResultValidationMethodEmail ZoneSslCertificatePackUpdateResponseResultValidationMethod = "email"
)

// Validity Days selected for the order.
type ZoneSslCertificatePackUpdateResponseResultValidityDays int64

const (
	ZoneSslCertificatePackUpdateResponseResultValidityDays14  ZoneSslCertificatePackUpdateResponseResultValidityDays = 14
	ZoneSslCertificatePackUpdateResponseResultValidityDays30  ZoneSslCertificatePackUpdateResponseResultValidityDays = 30
	ZoneSslCertificatePackUpdateResponseResultValidityDays90  ZoneSslCertificatePackUpdateResponseResultValidityDays = 90
	ZoneSslCertificatePackUpdateResponseResultValidityDays365 ZoneSslCertificatePackUpdateResponseResultValidityDays = 365
)

// Whether the API call was successful
type ZoneSslCertificatePackUpdateResponseSuccess bool

const (
	ZoneSslCertificatePackUpdateResponseSuccessTrue ZoneSslCertificatePackUpdateResponseSuccess = true
)

type ZoneSslCertificatePackDeleteResponse struct {
	Errors   []ZoneSslCertificatePackDeleteResponseError   `json:"errors"`
	Messages []ZoneSslCertificatePackDeleteResponseMessage `json:"messages"`
	Result   ZoneSslCertificatePackDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslCertificatePackDeleteResponseSuccess `json:"success"`
	JSON    zoneSslCertificatePackDeleteResponseJSON    `json:"-"`
}

// zoneSslCertificatePackDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackDeleteResponse]
type zoneSslCertificatePackDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSslCertificatePackDeleteResponseErrorJSON `json:"-"`
}

// zoneSslCertificatePackDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackDeleteResponseError]
type zoneSslCertificatePackDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSslCertificatePackDeleteResponseMessageJSON `json:"-"`
}

// zoneSslCertificatePackDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSslCertificatePackDeleteResponseMessage]
type zoneSslCertificatePackDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackDeleteResponseResult struct {
	// Identifier
	ID   string                                         `json:"id"`
	JSON zoneSslCertificatePackDeleteResponseResultJSON `json:"-"`
}

// zoneSslCertificatePackDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneSslCertificatePackDeleteResponseResult]
type zoneSslCertificatePackDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSslCertificatePackDeleteResponseSuccess bool

const (
	ZoneSslCertificatePackDeleteResponseSuccessTrue ZoneSslCertificatePackDeleteResponseSuccess = true
)

type ZoneSslCertificatePackCertificatePacksListCertificatePacksResponse struct {
	Errors     []ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseError    `json:"errors"`
	Messages   []ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessage  `json:"messages"`
	Result     []interface{}                                                                `json:"result"`
	ResultInfo ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseSuccess `json:"success"`
	JSON    zoneSslCertificatePackCertificatePacksListCertificatePacksResponseJSON    `json:"-"`
}

// zoneSslCertificatePackCertificatePacksListCertificatePacksResponseJSON contains
// the JSON metadata for the struct
// [ZoneSslCertificatePackCertificatePacksListCertificatePacksResponse]
type zoneSslCertificatePackCertificatePacksListCertificatePacksResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackCertificatePacksListCertificatePacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    zoneSslCertificatePackCertificatePacksListCertificatePacksResponseErrorJSON `json:"-"`
}

// zoneSslCertificatePackCertificatePacksListCertificatePacksResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseError]
type zoneSslCertificatePackCertificatePacksListCertificatePacksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    zoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessageJSON `json:"-"`
}

// zoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessage]
type zoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       zoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfoJSON `json:"-"`
}

// zoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfo]
type zoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseSuccess bool

const (
	ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseSuccessTrue ZoneSslCertificatePackCertificatePacksListCertificatePacksResponseSuccess = true
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
