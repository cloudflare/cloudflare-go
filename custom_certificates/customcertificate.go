// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_certificates

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/keyless_certificates"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// CustomCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomCertificateService] method
// instead.
type CustomCertificateService struct {
	Options    []option.RequestOption
	Prioritize *PrioritizeService
}

// NewCustomCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomCertificateService(opts ...option.RequestOption) (r *CustomCertificateService) {
	r = &CustomCertificateService{}
	r.Options = opts
	r.Prioritize = NewPrioritizeService(opts...)
	return
}

// Upload a new SSL certificate for a zone.
func (r *CustomCertificateService) New(ctx context.Context, params CustomCertificateNewParams, opts ...option.RequestOption) (res *CustomCertificateNewResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, and filter all of your custom SSL certificates. The higher
// priority will break ties across overlapping 'legacy_custom' certificates, but
// 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
func (r *CustomCertificateService) List(ctx context.Context, params CustomCertificateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CustomCertificate], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates", params.ZoneID)
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

// List, search, and filter all of your custom SSL certificates. The higher
// priority will break ties across overlapping 'legacy_custom' certificates, but
// 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
func (r *CustomCertificateService) ListAutoPaging(ctx context.Context, params CustomCertificateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CustomCertificate] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a SSL certificate from a zone.
func (r *CustomCertificateService) Delete(ctx context.Context, customCertificateID string, body CustomCertificateDeleteParams, opts ...option.RequestOption) (res *CustomCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", body.ZoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing
// a configuration for sni_custom certificates will result in a new resource id
// being returned, and the previous one being deleted.
func (r *CustomCertificateService) Edit(ctx context.Context, customCertificateID string, params CustomCertificateEditParams, opts ...option.RequestOption) (res *CustomCertificateEditResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", params.ZoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// SSL Configuration Details
func (r *CustomCertificateService) Get(ctx context.Context, customCertificateID string, query CustomCertificateGetParams, opts ...option.RequestOption) (res *CustomCertificateGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", query.ZoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomCertificate struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod custom_hostnames.BundleMethod `json:"bundle_method,required"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on,required" format:"date-time"`
	Hosts     []string  `json:"hosts,required"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer,required"`
	// When the certificate was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority float64 `json:"priority,required"`
	// The type of hash used for the certificate.
	Signature string `json:"signature,required"`
	// Status of the zone's custom SSL.
	Status CustomCertificateStatus `json:"status,required"`
	// When the certificate was uploaded to Cloudflare.
	UploadedOn time.Time `json:"uploaded_on,required" format:"date-time"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions GeoRestrictions                         `json:"geo_restrictions"`
	KeylessServer   keyless_certificates.KeylessCertificate `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                `json:"policy"`
	JSON   customCertificateJSON `json:"-"`
}

// customCertificateJSON contains the JSON metadata for the struct
// [CustomCertificate]
type customCertificateJSON struct {
	ID              apijson.Field
	BundleMethod    apijson.Field
	ExpiresOn       apijson.Field
	Hosts           apijson.Field
	Issuer          apijson.Field
	ModifiedOn      apijson.Field
	Priority        apijson.Field
	Signature       apijson.Field
	Status          apijson.Field
	UploadedOn      apijson.Field
	ZoneID          apijson.Field
	GeoRestrictions apijson.Field
	KeylessServer   apijson.Field
	Policy          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateJSON) RawJSON() string {
	return r.raw
}

// Status of the zone's custom SSL.
type CustomCertificateStatus string

const (
	CustomCertificateStatusActive       CustomCertificateStatus = "active"
	CustomCertificateStatusExpired      CustomCertificateStatus = "expired"
	CustomCertificateStatusDeleted      CustomCertificateStatus = "deleted"
	CustomCertificateStatusPending      CustomCertificateStatus = "pending"
	CustomCertificateStatusInitializing CustomCertificateStatus = "initializing"
)

func (r CustomCertificateStatus) IsKnown() bool {
	switch r {
	case CustomCertificateStatusActive, CustomCertificateStatusExpired, CustomCertificateStatusDeleted, CustomCertificateStatusPending, CustomCertificateStatusInitializing:
		return true
	}
	return false
}

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type GeoRestrictions struct {
	Label GeoRestrictionsLabel `json:"label"`
	JSON  geoRestrictionsJSON  `json:"-"`
}

// geoRestrictionsJSON contains the JSON metadata for the struct [GeoRestrictions]
type geoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geoRestrictionsJSON) RawJSON() string {
	return r.raw
}

type GeoRestrictionsLabel string

const (
	GeoRestrictionsLabelUs              GeoRestrictionsLabel = "us"
	GeoRestrictionsLabelEu              GeoRestrictionsLabel = "eu"
	GeoRestrictionsLabelHighestSecurity GeoRestrictionsLabel = "highest_security"
)

func (r GeoRestrictionsLabel) IsKnown() bool {
	switch r {
	case GeoRestrictionsLabelUs, GeoRestrictionsLabelEu, GeoRestrictionsLabelHighestSecurity:
		return true
	}
	return false
}

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type GeoRestrictionsParam struct {
	Label param.Field[GeoRestrictionsLabel] `json:"label"`
}

func (r GeoRestrictionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [custom_certificates.CustomCertificateNewResponseUnknown] or
// [shared.UnionString].
type CustomCertificateNewResponseUnion interface {
	ImplementsCustomCertificatesCustomCertificateNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomCertificateNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomCertificateDeleteResponse struct {
	// Identifier
	ID   string                              `json:"id"`
	JSON customCertificateDeleteResponseJSON `json:"-"`
}

// customCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [CustomCertificateDeleteResponse]
type customCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [custom_certificates.CustomCertificateEditResponseUnknown] or
// [shared.UnionString].
type CustomCertificateEditResponseUnion interface {
	ImplementsCustomCertificatesCustomCertificateEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomCertificateEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [custom_certificates.CustomCertificateGetResponseUnknown] or
// [shared.UnionString].
type CustomCertificateGetResponseUnion interface {
	ImplementsCustomCertificatesCustomCertificateGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomCertificateGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomCertificateNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[custom_hostnames.BundleMethod] `json:"bundle_method"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[GeoRestrictionsParam] `json:"geo_restrictions"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy param.Field[string] `json:"policy"`
	// The type 'legacy_custom' enables support for legacy clients which do not include
	// SNI in the TLS handshake.
	Type param.Field[CustomCertificateNewParamsType] `json:"type"`
}

func (r CustomCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type 'legacy_custom' enables support for legacy clients which do not include
// SNI in the TLS handshake.
type CustomCertificateNewParamsType string

const (
	CustomCertificateNewParamsTypeLegacyCustom CustomCertificateNewParamsType = "legacy_custom"
	CustomCertificateNewParamsTypeSNICustom    CustomCertificateNewParamsType = "sni_custom"
)

func (r CustomCertificateNewParamsType) IsKnown() bool {
	switch r {
	case CustomCertificateNewParamsTypeLegacyCustom, CustomCertificateNewParamsTypeSNICustom:
		return true
	}
	return false
}

type CustomCertificateNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CustomCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomCertificateNewResponseUnion           `json:"result"`
	JSON    customCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// customCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateNewResponseEnvelope]
type customCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomCertificateNewResponseEnvelopeSuccess bool

const (
	CustomCertificateNewResponseEnvelopeSuccessTrue CustomCertificateNewResponseEnvelopeSuccess = true
)

func (r CustomCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomCertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[CustomCertificateListParamsMatch] `query:"match"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Status of the zone's custom SSL.
	Status param.Field[CustomCertificateListParamsStatus] `query:"status"`
}

// URLQuery serializes [CustomCertificateListParams]'s query parameters as
// `url.Values`.
func (r CustomCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all search requirements or at least one (any).
type CustomCertificateListParamsMatch string

const (
	CustomCertificateListParamsMatchAny CustomCertificateListParamsMatch = "any"
	CustomCertificateListParamsMatchAll CustomCertificateListParamsMatch = "all"
)

func (r CustomCertificateListParamsMatch) IsKnown() bool {
	switch r {
	case CustomCertificateListParamsMatchAny, CustomCertificateListParamsMatchAll:
		return true
	}
	return false
}

// Status of the zone's custom SSL.
type CustomCertificateListParamsStatus string

const (
	CustomCertificateListParamsStatusActive       CustomCertificateListParamsStatus = "active"
	CustomCertificateListParamsStatusExpired      CustomCertificateListParamsStatus = "expired"
	CustomCertificateListParamsStatusDeleted      CustomCertificateListParamsStatus = "deleted"
	CustomCertificateListParamsStatusPending      CustomCertificateListParamsStatus = "pending"
	CustomCertificateListParamsStatusInitializing CustomCertificateListParamsStatus = "initializing"
)

func (r CustomCertificateListParamsStatus) IsKnown() bool {
	switch r {
	case CustomCertificateListParamsStatusActive, CustomCertificateListParamsStatusExpired, CustomCertificateListParamsStatusDeleted, CustomCertificateListParamsStatusPending, CustomCertificateListParamsStatusInitializing:
		return true
	}
	return false
}

type CustomCertificateDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomCertificateDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CustomCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomCertificateDeleteResponse                `json:"result"`
	JSON    customCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// customCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateDeleteResponseEnvelope]
type customCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomCertificateDeleteResponseEnvelopeSuccess bool

const (
	CustomCertificateDeleteResponseEnvelopeSuccessTrue CustomCertificateDeleteResponseEnvelopeSuccess = true
)

func (r CustomCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomCertificateEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[custom_hostnames.BundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[GeoRestrictionsParam] `json:"geo_restrictions"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy param.Field[string] `json:"policy"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key"`
}

func (r CustomCertificateEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificateEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CustomCertificateEditResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomCertificateEditResponseUnion           `json:"result"`
	JSON    customCertificateEditResponseEnvelopeJSON    `json:"-"`
}

// customCertificateEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateEditResponseEnvelope]
type customCertificateEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomCertificateEditResponseEnvelopeSuccess bool

const (
	CustomCertificateEditResponseEnvelopeSuccessTrue CustomCertificateEditResponseEnvelopeSuccess = true
)

func (r CustomCertificateEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCertificateEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomCertificateGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomCertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CustomCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomCertificateGetResponseUnion           `json:"result"`
	JSON    customCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// customCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateGetResponseEnvelope]
type customCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomCertificateGetResponseEnvelopeSuccess bool

const (
	CustomCertificateGetResponseEnvelopeSuccessTrue CustomCertificateGetResponseEnvelopeSuccess = true
)

func (r CustomCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
