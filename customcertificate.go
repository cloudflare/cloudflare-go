// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CustomCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomCertificateService] method
// instead.
type CustomCertificateService struct {
	Options    []option.RequestOption
	Prioritize *CustomCertificatePrioritizeService
}

// NewCustomCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomCertificateService(opts ...option.RequestOption) (r *CustomCertificateService) {
	r = &CustomCertificateService{}
	r.Options = opts
	r.Prioritize = NewCustomCertificatePrioritizeService(opts...)
	return
}

// Upload a new SSL certificate for a zone.
func (r *CustomCertificateService) New(ctx context.Context, zoneID string, body CustomCertificateNewParams, opts ...option.RequestOption) (res *CustomCertificateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, and filter all of your custom SSL certificates. The higher
// priority will break ties across overlapping 'legacy_custom' certificates, but
// 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
func (r *CustomCertificateService) List(ctx context.Context, zoneID string, query CustomCertificateListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[CustomCertificateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *CustomCertificateService) ListAutoPaging(ctx context.Context, zoneID string, query CustomCertificateListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[CustomCertificateListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneID, query, opts...))
}

// Remove a SSL certificate from a zone.
func (r *CustomCertificateService) Delete(ctx context.Context, zoneID string, customCertificateID string, opts ...option.RequestOption) (res *CustomCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneID, customCertificateID)
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
func (r *CustomCertificateService) Edit(ctx context.Context, zoneID string, customCertificateID string, body CustomCertificateEditParams, opts ...option.RequestOption) (res *CustomCertificateEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// SSL Configuration Details
func (r *CustomCertificateService) Get(ctx context.Context, zoneID string, customCertificateID string, opts ...option.RequestOption) (res *CustomCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CustomCertificateNewResponseUnknown] or
// [shared.UnionString].
type CustomCertificateNewResponse interface {
	ImplementsCustomCertificateNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomCertificateNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomCertificateListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomCertificateListResponseBundleMethod `json:"bundle_method,required"`
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
	Status CustomCertificateListResponseStatus `json:"status,required"`
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
	GeoRestrictions CustomCertificateListResponseGeoRestrictions `json:"geo_restrictions"`
	KeylessServer   CustomCertificateListResponseKeylessServer   `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                            `json:"policy"`
	JSON   customCertificateListResponseJSON `json:"-"`
}

// customCertificateListResponseJSON contains the JSON metadata for the struct
// [CustomCertificateListResponse]
type customCertificateListResponseJSON struct {
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

func (r *CustomCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomCertificateListResponseBundleMethod string

const (
	CustomCertificateListResponseBundleMethodUbiquitous CustomCertificateListResponseBundleMethod = "ubiquitous"
	CustomCertificateListResponseBundleMethodOptimal    CustomCertificateListResponseBundleMethod = "optimal"
	CustomCertificateListResponseBundleMethodForce      CustomCertificateListResponseBundleMethod = "force"
)

// Status of the zone's custom SSL.
type CustomCertificateListResponseStatus string

const (
	CustomCertificateListResponseStatusActive       CustomCertificateListResponseStatus = "active"
	CustomCertificateListResponseStatusExpired      CustomCertificateListResponseStatus = "expired"
	CustomCertificateListResponseStatusDeleted      CustomCertificateListResponseStatus = "deleted"
	CustomCertificateListResponseStatusPending      CustomCertificateListResponseStatus = "pending"
	CustomCertificateListResponseStatusInitializing CustomCertificateListResponseStatus = "initializing"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type CustomCertificateListResponseGeoRestrictions struct {
	Label CustomCertificateListResponseGeoRestrictionsLabel `json:"label"`
	JSON  customCertificateListResponseGeoRestrictionsJSON  `json:"-"`
}

// customCertificateListResponseGeoRestrictionsJSON contains the JSON metadata for
// the struct [CustomCertificateListResponseGeoRestrictions]
type customCertificateListResponseGeoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateListResponseGeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateListResponseGeoRestrictionsLabel string

const (
	CustomCertificateListResponseGeoRestrictionsLabelUs              CustomCertificateListResponseGeoRestrictionsLabel = "us"
	CustomCertificateListResponseGeoRestrictionsLabelEu              CustomCertificateListResponseGeoRestrictionsLabel = "eu"
	CustomCertificateListResponseGeoRestrictionsLabelHighestSecurity CustomCertificateListResponseGeoRestrictionsLabel = "highest_security"
)

type CustomCertificateListResponseKeylessServer struct {
	// Keyless certificate identifier tag.
	ID string `json:"id,required"`
	// When the Keyless SSL was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Whether or not the Keyless SSL is on or off.
	Enabled bool `json:"enabled,required"`
	// The keyless SSL name.
	Host string `json:"host,required" format:"hostname"`
	// When the Keyless SSL was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The keyless SSL name.
	Name string `json:"name,required"`
	// Available permissions for the Keyless SSL for the current user requesting the
	// item.
	Permissions []interface{} `json:"permissions,required"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port float64 `json:"port,required"`
	// Status of the Keyless SSL.
	Status CustomCertificateListResponseKeylessServerStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel CustomCertificateListResponseKeylessServerTunnel `json:"tunnel"`
	JSON   customCertificateListResponseKeylessServerJSON   `json:"-"`
}

// customCertificateListResponseKeylessServerJSON contains the JSON metadata for
// the struct [CustomCertificateListResponseKeylessServer]
type customCertificateListResponseKeylessServerJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	Host        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Port        apijson.Field
	Status      apijson.Field
	Tunnel      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateListResponseKeylessServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type CustomCertificateListResponseKeylessServerStatus string

const (
	CustomCertificateListResponseKeylessServerStatusActive  CustomCertificateListResponseKeylessServerStatus = "active"
	CustomCertificateListResponseKeylessServerStatusDeleted CustomCertificateListResponseKeylessServerStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type CustomCertificateListResponseKeylessServerTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                               `json:"vnet_id,required"`
	JSON   customCertificateListResponseKeylessServerTunnelJSON `json:"-"`
}

// customCertificateListResponseKeylessServerTunnelJSON contains the JSON metadata
// for the struct [CustomCertificateListResponseKeylessServerTunnel]
type customCertificateListResponseKeylessServerTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateListResponseKeylessServerTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Union satisfied by [CustomCertificateEditResponseUnknown] or
// [shared.UnionString].
type CustomCertificateEditResponse interface {
	ImplementsCustomCertificateEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomCertificateEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CustomCertificateGetResponseUnknown] or
// [shared.UnionString].
type CustomCertificateGetResponse interface {
	ImplementsCustomCertificateGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomCertificateGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomCertificateNewParams struct {
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[CustomCertificateNewParamsBundleMethod] `json:"bundle_method"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[CustomCertificateNewParamsGeoRestrictions] `json:"geo_restrictions"`
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

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomCertificateNewParamsBundleMethod string

const (
	CustomCertificateNewParamsBundleMethodUbiquitous CustomCertificateNewParamsBundleMethod = "ubiquitous"
	CustomCertificateNewParamsBundleMethodOptimal    CustomCertificateNewParamsBundleMethod = "optimal"
	CustomCertificateNewParamsBundleMethodForce      CustomCertificateNewParamsBundleMethod = "force"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type CustomCertificateNewParamsGeoRestrictions struct {
	Label param.Field[CustomCertificateNewParamsGeoRestrictionsLabel] `json:"label"`
}

func (r CustomCertificateNewParamsGeoRestrictions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificateNewParamsGeoRestrictionsLabel string

const (
	CustomCertificateNewParamsGeoRestrictionsLabelUs              CustomCertificateNewParamsGeoRestrictionsLabel = "us"
	CustomCertificateNewParamsGeoRestrictionsLabelEu              CustomCertificateNewParamsGeoRestrictionsLabel = "eu"
	CustomCertificateNewParamsGeoRestrictionsLabelHighestSecurity CustomCertificateNewParamsGeoRestrictionsLabel = "highest_security"
)

// The type 'legacy_custom' enables support for legacy clients which do not include
// SNI in the TLS handshake.
type CustomCertificateNewParamsType string

const (
	CustomCertificateNewParamsTypeLegacyCustom CustomCertificateNewParamsType = "legacy_custom"
	CustomCertificateNewParamsTypeSniCustom    CustomCertificateNewParamsType = "sni_custom"
)

type CustomCertificateNewResponseEnvelope struct {
	Errors   []CustomCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomCertificateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    customCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// customCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateNewResponseEnvelope]
type customCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    customCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// customCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomCertificateNewResponseEnvelopeErrors]
type customCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    customCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// customCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomCertificateNewResponseEnvelopeMessages]
type customCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomCertificateNewResponseEnvelopeSuccess bool

const (
	CustomCertificateNewResponseEnvelopeSuccessTrue CustomCertificateNewResponseEnvelopeSuccess = true
)

type CustomCertificateListParams struct {
	// Whether to match all search requirements or at least one (any).
	Match param.Field[CustomCertificateListParamsMatch] `query:"match"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [CustomCertificateListParams]'s query parameters as
// `url.Values`.
func (r CustomCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all search requirements or at least one (any).
type CustomCertificateListParamsMatch string

const (
	CustomCertificateListParamsMatchAny CustomCertificateListParamsMatch = "any"
	CustomCertificateListParamsMatchAll CustomCertificateListParamsMatch = "all"
)

type CustomCertificateDeleteResponseEnvelope struct {
	Errors   []CustomCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    customCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// customCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateDeleteResponseEnvelope]
type customCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    customCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// customCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomCertificateDeleteResponseEnvelopeErrors]
type customCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    customCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// customCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CustomCertificateDeleteResponseEnvelopeMessages]
type customCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomCertificateDeleteResponseEnvelopeSuccess bool

const (
	CustomCertificateDeleteResponseEnvelopeSuccessTrue CustomCertificateDeleteResponseEnvelopeSuccess = true
)

type CustomCertificateEditParams struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[CustomCertificateEditParamsBundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[CustomCertificateEditParamsGeoRestrictions] `json:"geo_restrictions"`
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

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomCertificateEditParamsBundleMethod string

const (
	CustomCertificateEditParamsBundleMethodUbiquitous CustomCertificateEditParamsBundleMethod = "ubiquitous"
	CustomCertificateEditParamsBundleMethodOptimal    CustomCertificateEditParamsBundleMethod = "optimal"
	CustomCertificateEditParamsBundleMethodForce      CustomCertificateEditParamsBundleMethod = "force"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type CustomCertificateEditParamsGeoRestrictions struct {
	Label param.Field[CustomCertificateEditParamsGeoRestrictionsLabel] `json:"label"`
}

func (r CustomCertificateEditParamsGeoRestrictions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificateEditParamsGeoRestrictionsLabel string

const (
	CustomCertificateEditParamsGeoRestrictionsLabelUs              CustomCertificateEditParamsGeoRestrictionsLabel = "us"
	CustomCertificateEditParamsGeoRestrictionsLabelEu              CustomCertificateEditParamsGeoRestrictionsLabel = "eu"
	CustomCertificateEditParamsGeoRestrictionsLabelHighestSecurity CustomCertificateEditParamsGeoRestrictionsLabel = "highest_security"
)

type CustomCertificateEditResponseEnvelope struct {
	Errors   []CustomCertificateEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomCertificateEditResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomCertificateEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomCertificateEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    customCertificateEditResponseEnvelopeJSON    `json:"-"`
}

// customCertificateEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateEditResponseEnvelope]
type customCertificateEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    customCertificateEditResponseEnvelopeErrorsJSON `json:"-"`
}

// customCertificateEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomCertificateEditResponseEnvelopeErrors]
type customCertificateEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    customCertificateEditResponseEnvelopeMessagesJSON `json:"-"`
}

// customCertificateEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomCertificateEditResponseEnvelopeMessages]
type customCertificateEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomCertificateEditResponseEnvelopeSuccess bool

const (
	CustomCertificateEditResponseEnvelopeSuccessTrue CustomCertificateEditResponseEnvelopeSuccess = true
)

type CustomCertificateGetResponseEnvelope struct {
	Errors   []CustomCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomCertificateGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    customCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// customCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomCertificateGetResponseEnvelope]
type customCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    customCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// customCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomCertificateGetResponseEnvelopeErrors]
type customCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    customCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// customCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomCertificateGetResponseEnvelopeMessages]
type customCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomCertificateGetResponseEnvelopeSuccess bool

const (
	CustomCertificateGetResponseEnvelopeSuccessTrue CustomCertificateGetResponseEnvelopeSuccess = true
)
