// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// SSLCertificatePackService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSSLCertificatePackService] method
// instead.
type SSLCertificatePackService struct {
	Options []option.RequestOption
	Order   *SSLCertificatePackOrderService
	Quota   *SSLCertificatePackQuotaService
}

// NewSSLCertificatePackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSSLCertificatePackService(opts ...option.RequestOption) (r *SSLCertificatePackService) {
	r = &SSLCertificatePackService{}
	r.Options = opts
	r.Order = NewSSLCertificatePackOrderService(opts...)
	r.Quota = NewSSLCertificatePackQuotaService(opts...)
	return
}

// For a given zone, restart validation for an advanced certificate pack. This is
// only a validation operation for a Certificate Pack in a validation_timed_out
// status.
func (r *SSLCertificatePackService) Update(ctx context.Context, zoneID string, certificatePackID string, opts ...option.RequestOption) (res *SSLCertificatePackUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, list all active certificate packs.
func (r *SSLCertificatePackService) List(ctx context.Context, zoneID string, query SSLCertificatePackListParams, opts ...option.RequestOption) (res *[]SSLCertificatePackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackListResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, delete an advanced certificate pack.
func (r *SSLCertificatePackService) Delete(ctx context.Context, zoneID string, certificatePackID string, opts ...option.RequestOption) (res *SSLCertificatePackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, get a certificate pack.
func (r *SSLCertificatePackService) Get(ctx context.Context, zoneID string, certificatePackID string, opts ...option.RequestOption) (res *SSLCertificatePackGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLCertificatePackUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority SSLCertificatePackUpdateResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status SSLCertificatePackUpdateResponseStatus `json:"status"`
	// Type of certificate pack.
	Type SSLCertificatePackUpdateResponseType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod SSLCertificatePackUpdateResponseValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays SSLCertificatePackUpdateResponseValidityDays `json:"validity_days"`
	JSON         sslCertificatePackUpdateResponseJSON         `json:"-"`
}

// sslCertificatePackUpdateResponseJSON contains the JSON metadata for the struct
// [SSLCertificatePackUpdateResponse]
type sslCertificatePackUpdateResponseJSON struct {
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

func (r *SSLCertificatePackUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type SSLCertificatePackUpdateResponseCertificateAuthority string

const (
	SSLCertificatePackUpdateResponseCertificateAuthorityGoogle      SSLCertificatePackUpdateResponseCertificateAuthority = "google"
	SSLCertificatePackUpdateResponseCertificateAuthorityLetsEncrypt SSLCertificatePackUpdateResponseCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type SSLCertificatePackUpdateResponseStatus string

const (
	SSLCertificatePackUpdateResponseStatusInitializing         SSLCertificatePackUpdateResponseStatus = "initializing"
	SSLCertificatePackUpdateResponseStatusPendingValidation    SSLCertificatePackUpdateResponseStatus = "pending_validation"
	SSLCertificatePackUpdateResponseStatusDeleted              SSLCertificatePackUpdateResponseStatus = "deleted"
	SSLCertificatePackUpdateResponseStatusPendingIssuance      SSLCertificatePackUpdateResponseStatus = "pending_issuance"
	SSLCertificatePackUpdateResponseStatusPendingDeployment    SSLCertificatePackUpdateResponseStatus = "pending_deployment"
	SSLCertificatePackUpdateResponseStatusPendingDeletion      SSLCertificatePackUpdateResponseStatus = "pending_deletion"
	SSLCertificatePackUpdateResponseStatusPendingExpiration    SSLCertificatePackUpdateResponseStatus = "pending_expiration"
	SSLCertificatePackUpdateResponseStatusExpired              SSLCertificatePackUpdateResponseStatus = "expired"
	SSLCertificatePackUpdateResponseStatusActive               SSLCertificatePackUpdateResponseStatus = "active"
	SSLCertificatePackUpdateResponseStatusInitializingTimedOut SSLCertificatePackUpdateResponseStatus = "initializing_timed_out"
	SSLCertificatePackUpdateResponseStatusValidationTimedOut   SSLCertificatePackUpdateResponseStatus = "validation_timed_out"
	SSLCertificatePackUpdateResponseStatusIssuanceTimedOut     SSLCertificatePackUpdateResponseStatus = "issuance_timed_out"
	SSLCertificatePackUpdateResponseStatusDeploymentTimedOut   SSLCertificatePackUpdateResponseStatus = "deployment_timed_out"
	SSLCertificatePackUpdateResponseStatusDeletionTimedOut     SSLCertificatePackUpdateResponseStatus = "deletion_timed_out"
	SSLCertificatePackUpdateResponseStatusPendingCleanup       SSLCertificatePackUpdateResponseStatus = "pending_cleanup"
	SSLCertificatePackUpdateResponseStatusStagingDeployment    SSLCertificatePackUpdateResponseStatus = "staging_deployment"
	SSLCertificatePackUpdateResponseStatusStagingActive        SSLCertificatePackUpdateResponseStatus = "staging_active"
	SSLCertificatePackUpdateResponseStatusDeactivating         SSLCertificatePackUpdateResponseStatus = "deactivating"
	SSLCertificatePackUpdateResponseStatusInactive             SSLCertificatePackUpdateResponseStatus = "inactive"
	SSLCertificatePackUpdateResponseStatusBackupIssued         SSLCertificatePackUpdateResponseStatus = "backup_issued"
	SSLCertificatePackUpdateResponseStatusHoldingDeployment    SSLCertificatePackUpdateResponseStatus = "holding_deployment"
)

// Type of certificate pack.
type SSLCertificatePackUpdateResponseType string

const (
	SSLCertificatePackUpdateResponseTypeAdvanced SSLCertificatePackUpdateResponseType = "advanced"
)

// Validation Method selected for the order.
type SSLCertificatePackUpdateResponseValidationMethod string

const (
	SSLCertificatePackUpdateResponseValidationMethodTxt   SSLCertificatePackUpdateResponseValidationMethod = "txt"
	SSLCertificatePackUpdateResponseValidationMethodHTTP  SSLCertificatePackUpdateResponseValidationMethod = "http"
	SSLCertificatePackUpdateResponseValidationMethodEmail SSLCertificatePackUpdateResponseValidationMethod = "email"
)

// Validity Days selected for the order.
type SSLCertificatePackUpdateResponseValidityDays int64

const (
	SSLCertificatePackUpdateResponseValidityDays14  SSLCertificatePackUpdateResponseValidityDays = 14
	SSLCertificatePackUpdateResponseValidityDays30  SSLCertificatePackUpdateResponseValidityDays = 30
	SSLCertificatePackUpdateResponseValidityDays90  SSLCertificatePackUpdateResponseValidityDays = 90
	SSLCertificatePackUpdateResponseValidityDays365 SSLCertificatePackUpdateResponseValidityDays = 365
)

type SSLCertificatePackListResponse = interface{}

type SSLCertificatePackDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON sslCertificatePackDeleteResponseJSON `json:"-"`
}

// sslCertificatePackDeleteResponseJSON contains the JSON metadata for the struct
// [SSLCertificatePackDeleteResponse]
type sslCertificatePackDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [SSLCertificatePackGetResponseUnknown] or
// [shared.UnionString].
type SSLCertificatePackGetResponse interface {
	ImplementsSSLCertificatePackGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SSLCertificatePackGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SSLCertificatePackUpdateResponseEnvelope struct {
	Errors   []SSLCertificatePackUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackUpdateResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLCertificatePackUpdateResponseEnvelope]
type sslCertificatePackUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    sslCertificatePackUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SSLCertificatePackUpdateResponseEnvelopeErrors]
type sslCertificatePackUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    sslCertificatePackUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SSLCertificatePackUpdateResponseEnvelopeMessages]
type sslCertificatePackUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackUpdateResponseEnvelopeSuccess bool

const (
	SSLCertificatePackUpdateResponseEnvelopeSuccessTrue SSLCertificatePackUpdateResponseEnvelopeSuccess = true
)

type SSLCertificatePackListParams struct {
	// Include Certificate Packs of all statuses, not just active ones.
	Status param.Field[SSLCertificatePackListParamsStatus] `query:"status"`
}

// URLQuery serializes [SSLCertificatePackListParams]'s query parameters as
// `url.Values`.
func (r SSLCertificatePackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include Certificate Packs of all statuses, not just active ones.
type SSLCertificatePackListParamsStatus string

const (
	SSLCertificatePackListParamsStatusAll SSLCertificatePackListParamsStatus = "all"
)

type SSLCertificatePackListResponseEnvelope struct {
	Errors   []SSLCertificatePackListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SSLCertificatePackListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SSLCertificatePackListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SSLCertificatePackListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       sslCertificatePackListResponseEnvelopeJSON       `json:"-"`
}

// sslCertificatePackListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLCertificatePackListResponseEnvelope]
type sslCertificatePackListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    sslCertificatePackListResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLCertificatePackListResponseEnvelopeErrors]
type sslCertificatePackListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    sslCertificatePackListResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SSLCertificatePackListResponseEnvelopeMessages]
type sslCertificatePackListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackListResponseEnvelopeSuccess bool

const (
	SSLCertificatePackListResponseEnvelopeSuccessTrue SSLCertificatePackListResponseEnvelopeSuccess = true
)

type SSLCertificatePackListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       sslCertificatePackListResponseEnvelopeResultInfoJSON `json:"-"`
}

// sslCertificatePackListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [SSLCertificatePackListResponseEnvelopeResultInfo]
type sslCertificatePackListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackDeleteResponseEnvelope struct {
	Errors   []SSLCertificatePackDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackDeleteResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLCertificatePackDeleteResponseEnvelope]
type sslCertificatePackDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    sslCertificatePackDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SSLCertificatePackDeleteResponseEnvelopeErrors]
type sslCertificatePackDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    sslCertificatePackDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SSLCertificatePackDeleteResponseEnvelopeMessages]
type sslCertificatePackDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackDeleteResponseEnvelopeSuccess bool

const (
	SSLCertificatePackDeleteResponseEnvelopeSuccessTrue SSLCertificatePackDeleteResponseEnvelopeSuccess = true
)

type SSLCertificatePackGetResponseEnvelope struct {
	Errors   []SSLCertificatePackGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackGetResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLCertificatePackGetResponseEnvelope]
type sslCertificatePackGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    sslCertificatePackGetResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLCertificatePackGetResponseEnvelopeErrors]
type sslCertificatePackGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    sslCertificatePackGetResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SSLCertificatePackGetResponseEnvelopeMessages]
type sslCertificatePackGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackGetResponseEnvelopeSuccess bool

const (
	SSLCertificatePackGetResponseEnvelopeSuccessTrue SSLCertificatePackGetResponseEnvelopeSuccess = true
)
