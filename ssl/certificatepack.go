// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
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
func (r *CertificatePackService) List(ctx context.Context, params CertificatePackListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CertificatePackListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// For a given zone, list all active certificate packs.
func (r *CertificatePackService) ListAutoPaging(ctx context.Context, params CertificatePackListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CertificatePackListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
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
func (r *CertificatePackService) Edit(ctx context.Context, certificatePackID string, params CertificatePackEditParams, opts ...option.RequestOption) (res *CertificatePackEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", params.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, get a certificate pack.
func (r *CertificatePackService) Get(ctx context.Context, certificatePackID string, query CertificatePackGetParams, opts ...option.RequestOption) (res *CertificatePackGetResponseUnion, err error) {
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

type Host = string

type HostParam = string

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
	Hosts []Host `json:"hosts"`
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

func (r CertificatePackEditResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseCertificateAuthorityGoogle, CertificatePackEditResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

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

func (r CertificatePackEditResponseStatus) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseStatusInitializing, CertificatePackEditResponseStatusPendingValidation, CertificatePackEditResponseStatusDeleted, CertificatePackEditResponseStatusPendingIssuance, CertificatePackEditResponseStatusPendingDeployment, CertificatePackEditResponseStatusPendingDeletion, CertificatePackEditResponseStatusPendingExpiration, CertificatePackEditResponseStatusExpired, CertificatePackEditResponseStatusActive, CertificatePackEditResponseStatusInitializingTimedOut, CertificatePackEditResponseStatusValidationTimedOut, CertificatePackEditResponseStatusIssuanceTimedOut, CertificatePackEditResponseStatusDeploymentTimedOut, CertificatePackEditResponseStatusDeletionTimedOut, CertificatePackEditResponseStatusPendingCleanup, CertificatePackEditResponseStatusStagingDeployment, CertificatePackEditResponseStatusStagingActive, CertificatePackEditResponseStatusDeactivating, CertificatePackEditResponseStatusInactive, CertificatePackEditResponseStatusBackupIssued, CertificatePackEditResponseStatusHoldingDeployment:
		return true
	}
	return false
}

// Type of certificate pack.
type CertificatePackEditResponseType string

const (
	CertificatePackEditResponseTypeAdvanced CertificatePackEditResponseType = "advanced"
)

func (r CertificatePackEditResponseType) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseTypeAdvanced:
		return true
	}
	return false
}

// Validation Method selected for the order.
type CertificatePackEditResponseValidationMethod string

const (
	CertificatePackEditResponseValidationMethodTXT   CertificatePackEditResponseValidationMethod = "txt"
	CertificatePackEditResponseValidationMethodHTTP  CertificatePackEditResponseValidationMethod = "http"
	CertificatePackEditResponseValidationMethodEmail CertificatePackEditResponseValidationMethod = "email"
)

func (r CertificatePackEditResponseValidationMethod) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseValidationMethodTXT, CertificatePackEditResponseValidationMethodHTTP, CertificatePackEditResponseValidationMethodEmail:
		return true
	}
	return false
}

// Validity Days selected for the order.
type CertificatePackEditResponseValidityDays int64

const (
	CertificatePackEditResponseValidityDays14  CertificatePackEditResponseValidityDays = 14
	CertificatePackEditResponseValidityDays30  CertificatePackEditResponseValidityDays = 30
	CertificatePackEditResponseValidityDays90  CertificatePackEditResponseValidityDays = 90
	CertificatePackEditResponseValidityDays365 CertificatePackEditResponseValidityDays = 365
)

func (r CertificatePackEditResponseValidityDays) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseValidityDays14, CertificatePackEditResponseValidityDays30, CertificatePackEditResponseValidityDays90, CertificatePackEditResponseValidityDays365:
		return true
	}
	return false
}

// Union satisfied by [ssl.CertificatePackGetResponseUnknown] or
// [shared.UnionString].
type CertificatePackGetResponseUnion interface {
	ImplementsSSLCertificatePackGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CertificatePackGetResponseUnion)(nil)).Elem(),
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include Certificate Packs of all statuses, not just active ones.
type CertificatePackListParamsStatus string

const (
	CertificatePackListParamsStatusAll CertificatePackListParamsStatus = "all"
)

func (r CertificatePackListParamsStatus) IsKnown() bool {
	switch r {
	case CertificatePackListParamsStatusAll:
		return true
	}
	return false
}

type CertificatePackDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   CertificatePackDeleteResponse `json:"result,required"`
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

// Whether the API call was successful
type CertificatePackDeleteResponseEnvelopeSuccess bool

const (
	CertificatePackDeleteResponseEnvelopeSuccessTrue CertificatePackDeleteResponseEnvelopeSuccess = true
)

func (r CertificatePackDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CertificatePackEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r CertificatePackEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CertificatePackEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   CertificatePackEditResponse `json:"result,required"`
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

// Whether the API call was successful
type CertificatePackEditResponseEnvelopeSuccess bool

const (
	CertificatePackEditResponseEnvelopeSuccessTrue CertificatePackEditResponseEnvelopeSuccess = true
)

func (r CertificatePackEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CertificatePackGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   CertificatePackGetResponseUnion `json:"result,required"`
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

// Whether the API call was successful
type CertificatePackGetResponseEnvelopeSuccess bool

const (
	CertificatePackGetResponseEnvelopeSuccessTrue CertificatePackGetResponseEnvelopeSuccess = true
)

func (r CertificatePackGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
