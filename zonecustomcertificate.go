// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneCustomCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCustomCertificateService]
// method instead.
type ZoneCustomCertificateService struct {
	Options     []option.RequestOption
	Prioritizes *ZoneCustomCertificatePrioritizeService
}

// NewZoneCustomCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCustomCertificateService(opts ...option.RequestOption) (r *ZoneCustomCertificateService) {
	r = &ZoneCustomCertificateService{}
	r.Options = opts
	r.Prioritizes = NewZoneCustomCertificatePrioritizeService(opts...)
	return
}

// SSL Configuration Details
func (r *ZoneCustomCertificateService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *CertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing
// a configuration for sni_custom certificates will result in a new resource id
// being returned, and the previous one being deleted.
func (r *ZoneCustomCertificateService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneCustomCertificateUpdateParams, opts ...option.RequestOption) (res *CertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Remove a SSL certificate from a zone.
func (r *ZoneCustomCertificateService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *CertificateResponseIDOnlyTMCaIs8i, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload a new SSL certificate for a zone.
func (r *ZoneCustomCertificateService) CustomSslForAZoneNewSslConfiguration(ctx context.Context, zoneIdentifier string, body ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParams, opts ...option.RequestOption) (res *CertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, and filter all of your custom SSL certificates. The higher
// priority will break ties across overlapping 'legacy_custom' certificates, but
// 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
func (r *ZoneCustomCertificateService) CustomSslForAZoneListSslConfigurations(ctx context.Context, zoneIdentifier string, query ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParams, opts ...option.RequestOption) (res *CertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CertificateResponseCollection struct {
	Errors     []CertificateResponseCollectionError    `json:"errors"`
	Messages   []CertificateResponseCollectionMessage  `json:"messages"`
	Result     []CertificateResponseCollectionResult   `json:"result"`
	ResultInfo CertificateResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CertificateResponseCollectionSuccess `json:"success"`
	JSON    certificateResponseCollectionJSON    `json:"-"`
}

// certificateResponseCollectionJSON contains the JSON metadata for the struct
// [CertificateResponseCollection]
type certificateResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    certificateResponseCollectionErrorJSON `json:"-"`
}

// certificateResponseCollectionErrorJSON contains the JSON metadata for the struct
// [CertificateResponseCollectionError]
type certificateResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    certificateResponseCollectionMessageJSON `json:"-"`
}

// certificateResponseCollectionMessageJSON contains the JSON metadata for the
// struct [CertificateResponseCollectionMessage]
type certificateResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseCollectionResult struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CertificateResponseCollectionResultBundleMethod `json:"bundle_method,required"`
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
	Status CertificateResponseCollectionResultStatus `json:"status,required"`
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
	GeoRestrictions CertificateResponseCollectionResultGeoRestrictions `json:"geo_restrictions"`
	KeylessServer   CertificateResponseCollectionResultKeylessServer   `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                                  `json:"policy"`
	JSON   certificateResponseCollectionResultJSON `json:"-"`
}

// certificateResponseCollectionResultJSON contains the JSON metadata for the
// struct [CertificateResponseCollectionResult]
type certificateResponseCollectionResultJSON struct {
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

func (r *CertificateResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CertificateResponseCollectionResultBundleMethod string

const (
	CertificateResponseCollectionResultBundleMethodUbiquitous CertificateResponseCollectionResultBundleMethod = "ubiquitous"
	CertificateResponseCollectionResultBundleMethodOptimal    CertificateResponseCollectionResultBundleMethod = "optimal"
	CertificateResponseCollectionResultBundleMethodForce      CertificateResponseCollectionResultBundleMethod = "force"
)

// Status of the zone's custom SSL.
type CertificateResponseCollectionResultStatus string

const (
	CertificateResponseCollectionResultStatusActive       CertificateResponseCollectionResultStatus = "active"
	CertificateResponseCollectionResultStatusExpired      CertificateResponseCollectionResultStatus = "expired"
	CertificateResponseCollectionResultStatusDeleted      CertificateResponseCollectionResultStatus = "deleted"
	CertificateResponseCollectionResultStatusPending      CertificateResponseCollectionResultStatus = "pending"
	CertificateResponseCollectionResultStatusInitializing CertificateResponseCollectionResultStatus = "initializing"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type CertificateResponseCollectionResultGeoRestrictions struct {
	Label CertificateResponseCollectionResultGeoRestrictionsLabel `json:"label"`
	JSON  certificateResponseCollectionResultGeoRestrictionsJSON  `json:"-"`
}

// certificateResponseCollectionResultGeoRestrictionsJSON contains the JSON
// metadata for the struct [CertificateResponseCollectionResultGeoRestrictions]
type certificateResponseCollectionResultGeoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollectionResultGeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseCollectionResultGeoRestrictionsLabel string

const (
	CertificateResponseCollectionResultGeoRestrictionsLabelUs              CertificateResponseCollectionResultGeoRestrictionsLabel = "us"
	CertificateResponseCollectionResultGeoRestrictionsLabelEu              CertificateResponseCollectionResultGeoRestrictionsLabel = "eu"
	CertificateResponseCollectionResultGeoRestrictionsLabelHighestSecurity CertificateResponseCollectionResultGeoRestrictionsLabel = "highest_security"
)

type CertificateResponseCollectionResultKeylessServer struct {
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
	Status CertificateResponseCollectionResultKeylessServerStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel CertificateResponseCollectionResultKeylessServerTunnel `json:"tunnel"`
	JSON   certificateResponseCollectionResultKeylessServerJSON   `json:"-"`
}

// certificateResponseCollectionResultKeylessServerJSON contains the JSON metadata
// for the struct [CertificateResponseCollectionResultKeylessServer]
type certificateResponseCollectionResultKeylessServerJSON struct {
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

func (r *CertificateResponseCollectionResultKeylessServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type CertificateResponseCollectionResultKeylessServerStatus string

const (
	CertificateResponseCollectionResultKeylessServerStatusActive  CertificateResponseCollectionResultKeylessServerStatus = "active"
	CertificateResponseCollectionResultKeylessServerStatusDeleted CertificateResponseCollectionResultKeylessServerStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type CertificateResponseCollectionResultKeylessServerTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                                     `json:"vnet_id,required"`
	JSON   certificateResponseCollectionResultKeylessServerTunnelJSON `json:"-"`
}

// certificateResponseCollectionResultKeylessServerTunnelJSON contains the JSON
// metadata for the struct [CertificateResponseCollectionResultKeylessServerTunnel]
type certificateResponseCollectionResultKeylessServerTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollectionResultKeylessServerTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       certificateResponseCollectionResultInfoJSON `json:"-"`
}

// certificateResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [CertificateResponseCollectionResultInfo]
type certificateResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateResponseCollectionSuccess bool

const (
	CertificateResponseCollectionSuccessTrue CertificateResponseCollectionSuccess = true
)

type CertificateResponseIDOnlyTMCaIs8i struct {
	Errors   []CertificateResponseIDOnlyTMCaIs8iError   `json:"errors"`
	Messages []CertificateResponseIDOnlyTMCaIs8iMessage `json:"messages"`
	Result   CertificateResponseIDOnlyTMCaIs8iResult    `json:"result"`
	// Whether the API call was successful
	Success CertificateResponseIDOnlyTMCaIs8iSuccess `json:"success"`
	JSON    certificateResponseIDOnlyTmCaIs8iJSON    `json:"-"`
}

// certificateResponseIDOnlyTmCaIs8iJSON contains the JSON metadata for the struct
// [CertificateResponseIDOnlyTMCaIs8i]
type certificateResponseIDOnlyTmCaIs8iJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseIDOnlyTMCaIs8i) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseIDOnlyTMCaIs8iError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    certificateResponseIDOnlyTmCaIs8iErrorJSON `json:"-"`
}

// certificateResponseIDOnlyTmCaIs8iErrorJSON contains the JSON metadata for the
// struct [CertificateResponseIDOnlyTMCaIs8iError]
type certificateResponseIDOnlyTmCaIs8iErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseIDOnlyTMCaIs8iError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseIDOnlyTMCaIs8iMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    certificateResponseIDOnlyTmCaIs8iMessageJSON `json:"-"`
}

// certificateResponseIDOnlyTmCaIs8iMessageJSON contains the JSON metadata for the
// struct [CertificateResponseIDOnlyTMCaIs8iMessage]
type certificateResponseIDOnlyTmCaIs8iMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseIDOnlyTMCaIs8iMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseIDOnlyTMCaIs8iResult struct {
	// Identifier
	ID   string                                      `json:"id"`
	JSON certificateResponseIDOnlyTmCaIs8iResultJSON `json:"-"`
}

// certificateResponseIDOnlyTmCaIs8iResultJSON contains the JSON metadata for the
// struct [CertificateResponseIDOnlyTMCaIs8iResult]
type certificateResponseIDOnlyTmCaIs8iResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseIDOnlyTMCaIs8iResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateResponseIDOnlyTMCaIs8iSuccess bool

const (
	CertificateResponseIDOnlyTMCaIs8iSuccessTrue CertificateResponseIDOnlyTMCaIs8iSuccess = true
)

type CertificateResponseSingle struct {
	Errors   []CertificateResponseSingleError   `json:"errors"`
	Messages []CertificateResponseSingleMessage `json:"messages"`
	Result   interface{}                        `json:"result"`
	// Whether the API call was successful
	Success CertificateResponseSingleSuccess `json:"success"`
	JSON    certificateResponseSingleJSON    `json:"-"`
}

// certificateResponseSingleJSON contains the JSON metadata for the struct
// [CertificateResponseSingle]
type certificateResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSingleError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    certificateResponseSingleErrorJSON `json:"-"`
}

// certificateResponseSingleErrorJSON contains the JSON metadata for the struct
// [CertificateResponseSingleError]
type certificateResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateResponseSingleMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    certificateResponseSingleMessageJSON `json:"-"`
}

// certificateResponseSingleMessageJSON contains the JSON metadata for the struct
// [CertificateResponseSingleMessage]
type certificateResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateResponseSingleSuccess bool

const (
	CertificateResponseSingleSuccessTrue CertificateResponseSingleSuccess = true
)

type ZoneCustomCertificateUpdateParams struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[ZoneCustomCertificateUpdateParamsBundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[ZoneCustomCertificateUpdateParamsGeoRestrictions] `json:"geo_restrictions"`
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

func (r ZoneCustomCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomCertificateUpdateParamsBundleMethod string

const (
	ZoneCustomCertificateUpdateParamsBundleMethodUbiquitous ZoneCustomCertificateUpdateParamsBundleMethod = "ubiquitous"
	ZoneCustomCertificateUpdateParamsBundleMethodOptimal    ZoneCustomCertificateUpdateParamsBundleMethod = "optimal"
	ZoneCustomCertificateUpdateParamsBundleMethodForce      ZoneCustomCertificateUpdateParamsBundleMethod = "force"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type ZoneCustomCertificateUpdateParamsGeoRestrictions struct {
	Label param.Field[ZoneCustomCertificateUpdateParamsGeoRestrictionsLabel] `json:"label"`
}

func (r ZoneCustomCertificateUpdateParamsGeoRestrictions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificateUpdateParamsGeoRestrictionsLabel string

const (
	ZoneCustomCertificateUpdateParamsGeoRestrictionsLabelUs              ZoneCustomCertificateUpdateParamsGeoRestrictionsLabel = "us"
	ZoneCustomCertificateUpdateParamsGeoRestrictionsLabelEu              ZoneCustomCertificateUpdateParamsGeoRestrictionsLabel = "eu"
	ZoneCustomCertificateUpdateParamsGeoRestrictionsLabelHighestSecurity ZoneCustomCertificateUpdateParamsGeoRestrictionsLabel = "highest_security"
)

type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParams struct {
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethod] `json:"bundle_method"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictions] `json:"geo_restrictions"`
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
	Type param.Field[ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsType] `json:"type"`
}

func (r ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethod string

const (
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethodUbiquitous ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethod = "ubiquitous"
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethodOptimal    ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethod = "optimal"
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethodForce      ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsBundleMethod = "force"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictions struct {
	Label param.Field[ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabel] `json:"label"`
}

func (r ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabel string

const (
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabelUs              ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabel = "us"
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabelEu              ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabel = "eu"
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabelHighestSecurity ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsGeoRestrictionsLabel = "highest_security"
)

// The type 'legacy_custom' enables support for legacy clients which do not include
// SNI in the TLS handshake.
type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsType string

const (
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsTypeLegacyCustom ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsType = "legacy_custom"
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsTypeSniCustom    ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParamsType = "sni_custom"
)

type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParams struct {
	// Whether to match all search requirements or at least one (any).
	Match param.Field[ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParamsMatch] `query:"match"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParams]'s query
// parameters as `url.Values`.
func (r ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all search requirements or at least one (any).
type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParamsMatch string

const (
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParamsMatchAny ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParamsMatch = "any"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParamsMatchAll ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParamsMatch = "all"
)
