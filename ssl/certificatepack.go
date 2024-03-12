// File generated from our OpenAPI spec by Stainless.

package ssl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// CertificatePackService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCertificatePackService] method
// instead.
type CertificatePackService struct {
	Options []option.RequestOption
	Order   *CertificatePackOrderService
	Quota   *CertificatePackQuotaService
}

// NewCertificatePackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificatePackService(opts ...option.RequestOption) (r *CertificatePackService) {
	r = &CertificatePackService{}
	r.Options = opts
	r.Order = NewCertificatePackOrderService(opts...)
	r.Quota = NewCertificatePackQuotaService(opts...)
	return
}

// For a given zone, list all active certificate packs.
func (r *CertificatePackService) List(ctx context.Context, params CertificatePackListParams, opts ...option.RequestOption) (res *[]CertificatePackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackListResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, delete an advanced certificate pack.
func (r *CertificatePackService) Delete(ctx context.Context, certificatePackID string, body CertificatePackDeleteParams, opts ...option.RequestOption) (res *CertificatePackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", body.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, restart validation for an advanced certificate pack. This is
// only a validation operation for a Certificate Pack in a validation_timed_out
// status.
func (r *CertificatePackService) Edit(ctx context.Context, certificatePackID string, body CertificatePackEditParams, opts ...option.RequestOption) (res *CertificatePackEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", body.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, get a certificate pack.
func (r *CertificatePackService) Get(ctx context.Context, certificatePackID string, query CertificatePackGetParams, opts ...option.RequestOption) (res *CertificatePackGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", query.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificatePackListResponse = interface{}

type CertificatePackDeleteResponse struct {
	// Identifier
	ID   string                            `json:"id"`
	JSON certificatePackDeleteResponseJSON `json:"-"`
}

// certificatePackDeleteResponseJSON contains the JSON metadata for the struct
// [CertificatePackDeleteResponse]
type certificatePackDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponse struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority CertificatePackEditResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status CertificatePackEditResponseStatus `json:"status"`
	// Type of certificate pack.
	Type CertificatePackEditResponseType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod CertificatePackEditResponseValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays CertificatePackEditResponseValidityDays `json:"validity_days"`
	JSON         certificatePackEditResponseJSON         `json:"-"`
}

// certificatePackEditResponseJSON contains the JSON metadata for the struct
// [CertificatePackEditResponse]
type certificatePackEditResponseJSON struct {
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

func (r *CertificatePackEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackEditResponseCertificateAuthority string

const (
	CertificatePackEditResponseCertificateAuthorityGoogle      CertificatePackEditResponseCertificateAuthority = "google"
	CertificatePackEditResponseCertificateAuthorityLetsEncrypt CertificatePackEditResponseCertificateAuthority = "lets_encrypt"
)

// Status of certificate pack.
type CertificatePackEditResponseStatus string

const (
	CertificatePackEditResponseStatusInitializing         CertificatePackEditResponseStatus = "initializing"
	CertificatePackEditResponseStatusPendingValidation    CertificatePackEditResponseStatus = "pending_validation"
	CertificatePackEditResponseStatusDeleted              CertificatePackEditResponseStatus = "deleted"
	CertificatePackEditResponseStatusPendingIssuance      CertificatePackEditResponseStatus = "pending_issuance"
	CertificatePackEditResponseStatusPendingDeployment    CertificatePackEditResponseStatus = "pending_deployment"
	CertificatePackEditResponseStatusPendingDeletion      CertificatePackEditResponseStatus = "pending_deletion"
	CertificatePackEditResponseStatusPendingExpiration    CertificatePackEditResponseStatus = "pending_expiration"
	CertificatePackEditResponseStatusExpired              CertificatePackEditResponseStatus = "expired"
	CertificatePackEditResponseStatusActive               CertificatePackEditResponseStatus = "active"
	CertificatePackEditResponseStatusInitializingTimedOut CertificatePackEditResponseStatus = "initializing_timed_out"
	CertificatePackEditResponseStatusValidationTimedOut   CertificatePackEditResponseStatus = "validation_timed_out"
	CertificatePackEditResponseStatusIssuanceTimedOut     CertificatePackEditResponseStatus = "issuance_timed_out"
	CertificatePackEditResponseStatusDeploymentTimedOut   CertificatePackEditResponseStatus = "deployment_timed_out"
	CertificatePackEditResponseStatusDeletionTimedOut     CertificatePackEditResponseStatus = "deletion_timed_out"
	CertificatePackEditResponseStatusPendingCleanup       CertificatePackEditResponseStatus = "pending_cleanup"
	CertificatePackEditResponseStatusStagingDeployment    CertificatePackEditResponseStatus = "staging_deployment"
	CertificatePackEditResponseStatusStagingActive        CertificatePackEditResponseStatus = "staging_active"
	CertificatePackEditResponseStatusDeactivating         CertificatePackEditResponseStatus = "deactivating"
	CertificatePackEditResponseStatusInactive             CertificatePackEditResponseStatus = "inactive"
	CertificatePackEditResponseStatusBackupIssued         CertificatePackEditResponseStatus = "backup_issued"
	CertificatePackEditResponseStatusHoldingDeployment    CertificatePackEditResponseStatus = "holding_deployment"
)

// Type of certificate pack.
type CertificatePackEditResponseType string

const (
	CertificatePackEditResponseTypeAdvanced CertificatePackEditResponseType = "advanced"
)

// Validation Method selected for the order.
type CertificatePackEditResponseValidationMethod string

const (
	CertificatePackEditResponseValidationMethodTXT   CertificatePackEditResponseValidationMethod = "txt"
	CertificatePackEditResponseValidationMethodHTTP  CertificatePackEditResponseValidationMethod = "http"
	CertificatePackEditResponseValidationMethodEmail CertificatePackEditResponseValidationMethod = "email"
)

// Validity Days selected for the order.
type CertificatePackEditResponseValidityDays int64

const (
	CertificatePackEditResponseValidityDays14  CertificatePackEditResponseValidityDays = 14
	CertificatePackEditResponseValidityDays30  CertificatePackEditResponseValidityDays = 30
	CertificatePackEditResponseValidityDays90  CertificatePackEditResponseValidityDays = 90
	CertificatePackEditResponseValidityDays365 CertificatePackEditResponseValidityDays = 365
)

// Union satisfied by [ssl.CertificatePackGetResponseUnknown] or
// [shared.UnionString].
type CertificatePackGetResponse interface {
	ImplementsSSLCertificatePackGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CertificatePackGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CertificatePackListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Include Certificate Packs of all statuses, not just active ones.
	Status param.Field[CertificatePackListParamsStatus] `query:"status"`
}

// URLQuery serializes [CertificatePackListParams]'s query parameters as
// `url.Values`.
func (r CertificatePackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include Certificate Packs of all statuses, not just active ones.
type CertificatePackListParamsStatus string

const (
	CertificatePackListParamsStatusAll CertificatePackListParamsStatus = "all"
)

type CertificatePackListResponseEnvelope struct {
	Errors   []CertificatePackListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackListResponseEnvelopeMessages `json:"messages,required"`
	Result   []CertificatePackListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CertificatePackListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CertificatePackListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       certificatePackListResponseEnvelopeJSON       `json:"-"`
}

// certificatePackListResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackListResponseEnvelope]
type certificatePackListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    certificatePackListResponseEnvelopeErrorsJSON `json:"-"`
}

// certificatePackListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificatePackListResponseEnvelopeErrors]
type certificatePackListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    certificatePackListResponseEnvelopeMessagesJSON `json:"-"`
}

// certificatePackListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackListResponseEnvelopeMessages]
type certificatePackListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackListResponseEnvelopeSuccess bool

const (
	CertificatePackListResponseEnvelopeSuccessTrue CertificatePackListResponseEnvelopeSuccess = true
)

type CertificatePackListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       certificatePackListResponseEnvelopeResultInfoJSON `json:"-"`
}

// certificatePackListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [CertificatePackListResponseEnvelopeResultInfo]
type certificatePackListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackDeleteResponseEnvelope struct {
	Errors   []CertificatePackDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificatePackDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificatePackDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificatePackDeleteResponseEnvelopeJSON    `json:"-"`
}

// certificatePackDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackDeleteResponseEnvelope]
type certificatePackDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    certificatePackDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// certificatePackDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CertificatePackDeleteResponseEnvelopeErrors]
type certificatePackDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    certificatePackDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// certificatePackDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackDeleteResponseEnvelopeMessages]
type certificatePackDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackDeleteResponseEnvelopeSuccess bool

const (
	CertificatePackDeleteResponseEnvelopeSuccessTrue CertificatePackDeleteResponseEnvelopeSuccess = true
)

type CertificatePackEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackEditResponseEnvelope struct {
	Errors   []CertificatePackEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackEditResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificatePackEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificatePackEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificatePackEditResponseEnvelopeJSON    `json:"-"`
}

// certificatePackEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackEditResponseEnvelope]
type certificatePackEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    certificatePackEditResponseEnvelopeErrorsJSON `json:"-"`
}

// certificatePackEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificatePackEditResponseEnvelopeErrors]
type certificatePackEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    certificatePackEditResponseEnvelopeMessagesJSON `json:"-"`
}

// certificatePackEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackEditResponseEnvelopeMessages]
type certificatePackEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackEditResponseEnvelopeSuccess bool

const (
	CertificatePackEditResponseEnvelopeSuccessTrue CertificatePackEditResponseEnvelopeSuccess = true
)

type CertificatePackGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackGetResponseEnvelope struct {
	Errors   []CertificatePackGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificatePackGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificatePackGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificatePackGetResponseEnvelopeJSON    `json:"-"`
}

// certificatePackGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificatePackGetResponseEnvelope]
type certificatePackGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    certificatePackGetResponseEnvelopeErrorsJSON `json:"-"`
}

// certificatePackGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificatePackGetResponseEnvelopeErrors]
type certificatePackGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    certificatePackGetResponseEnvelopeMessagesJSON `json:"-"`
}

// certificatePackGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackGetResponseEnvelopeMessages]
type certificatePackGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackGetResponseEnvelopeSuccess bool

const (
	CertificatePackGetResponseEnvelopeSuccessTrue CertificatePackGetResponseEnvelopeSuccess = true
)
