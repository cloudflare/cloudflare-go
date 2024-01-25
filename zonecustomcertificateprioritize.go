// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneCustomCertificatePrioritizeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneCustomCertificatePrioritizeService] method instead.
type ZoneCustomCertificatePrioritizeService struct {
	Options []option.RequestOption
}

// NewZoneCustomCertificatePrioritizeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneCustomCertificatePrioritizeService(opts ...option.RequestOption) (r *ZoneCustomCertificatePrioritizeService) {
	r = &ZoneCustomCertificatePrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *ZoneCustomCertificatePrioritizeService) CustomSslForAZoneRePrioritizeSslCertificates(ctx context.Context, zoneIdentifier string, body ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParams, opts ...option.RequestOption) (res *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponse struct {
	Errors     []ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseError    `json:"errors"`
	Messages   []ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessage  `json:"messages"`
	Result     []ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResult   `json:"result"`
	ResultInfo ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseSuccess `json:"success"`
	JSON    zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseJSON    `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponse]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseErrorJSON `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseError]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessageJSON `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessage]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethod `json:"bundle_method,required"`
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
	Status ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus `json:"status,required"`
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
	GeoRestrictions ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictions `json:"geo_restrictions"`
	KeylessServer   ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServer   `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                                                                                        `json:"policy"`
	JSON   zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultJSON `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResult]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultJSON struct {
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

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethod string

const (
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethodUbiquitous ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethod = "ubiquitous"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethodOptimal    ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethod = "optimal"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethodForce      ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultBundleMethod = "force"
)

// Status of the zone's custom SSL.
type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus string

const (
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatusActive       ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus = "active"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatusExpired      ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus = "expired"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatusDeleted      ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus = "deleted"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatusPending      ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus = "pending"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatusInitializing ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultStatus = "initializing"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictions struct {
	Label ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabel `json:"label"`
	JSON  zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsJSON  `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictions]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabel string

const (
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabelUs              ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabel = "us"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabelEu              ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabel = "eu"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabelHighestSecurity ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultGeoRestrictionsLabel = "highest_security"
)

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServer struct {
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
	Status ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnel `json:"tunnel"`
	JSON   zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerJSON   `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServer]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerJSON struct {
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

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerStatus string

const (
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerStatusActive  ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerStatus = "active"
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerStatusDeleted ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                                                                                           `json:"vnet_id,required"`
	JSON   zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnelJSON `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnelJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnel]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultKeylessServerTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                           `json:"total_count"`
	JSON       zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfoJSON `json:"-"`
}

// zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfo]
type zoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseSuccess bool

const (
	ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseSuccessTrue ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesResponseSuccess = true
)

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParams struct {
	// Array of ordered certificates.
	Certificates param.Field[[]ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParamsCertificate] `json:"certificates,required"`
}

func (r ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParamsCertificate struct {
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
