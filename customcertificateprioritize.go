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

// CustomCertificatePrioritizeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewCustomCertificatePrioritizeService] method instead.
type CustomCertificatePrioritizeService struct {
	Options []option.RequestOption
}

// NewCustomCertificatePrioritizeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomCertificatePrioritizeService(opts ...option.RequestOption) (r *CustomCertificatePrioritizeService) {
	r = &CustomCertificatePrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *CustomCertificatePrioritizeService) Replace(ctx context.Context, zoneID string, body CustomCertificatePrioritizeReplaceParams, opts ...option.RequestOption) (res *[]CustomCertificatePrioritizeReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificatePrioritizeReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomCertificatePrioritizeReplaceResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomCertificatePrioritizeReplaceResponseBundleMethod `json:"bundle_method,required"`
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
	Status CustomCertificatePrioritizeReplaceResponseStatus `json:"status,required"`
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
	GeoRestrictions CustomCertificatePrioritizeReplaceResponseGeoRestrictions `json:"geo_restrictions"`
	KeylessServer   CustomCertificatePrioritizeReplaceResponseKeylessServer   `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                                         `json:"policy"`
	JSON   customCertificatePrioritizeReplaceResponseJSON `json:"-"`
}

// customCertificatePrioritizeReplaceResponseJSON contains the JSON metadata for
// the struct [CustomCertificatePrioritizeReplaceResponse]
type customCertificatePrioritizeReplaceResponseJSON struct {
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

func (r *CustomCertificatePrioritizeReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomCertificatePrioritizeReplaceResponseBundleMethod string

const (
	CustomCertificatePrioritizeReplaceResponseBundleMethodUbiquitous CustomCertificatePrioritizeReplaceResponseBundleMethod = "ubiquitous"
	CustomCertificatePrioritizeReplaceResponseBundleMethodOptimal    CustomCertificatePrioritizeReplaceResponseBundleMethod = "optimal"
	CustomCertificatePrioritizeReplaceResponseBundleMethodForce      CustomCertificatePrioritizeReplaceResponseBundleMethod = "force"
)

// Status of the zone's custom SSL.
type CustomCertificatePrioritizeReplaceResponseStatus string

const (
	CustomCertificatePrioritizeReplaceResponseStatusActive       CustomCertificatePrioritizeReplaceResponseStatus = "active"
	CustomCertificatePrioritizeReplaceResponseStatusExpired      CustomCertificatePrioritizeReplaceResponseStatus = "expired"
	CustomCertificatePrioritizeReplaceResponseStatusDeleted      CustomCertificatePrioritizeReplaceResponseStatus = "deleted"
	CustomCertificatePrioritizeReplaceResponseStatusPending      CustomCertificatePrioritizeReplaceResponseStatus = "pending"
	CustomCertificatePrioritizeReplaceResponseStatusInitializing CustomCertificatePrioritizeReplaceResponseStatus = "initializing"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type CustomCertificatePrioritizeReplaceResponseGeoRestrictions struct {
	Label CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabel `json:"label"`
	JSON  customCertificatePrioritizeReplaceResponseGeoRestrictionsJSON  `json:"-"`
}

// customCertificatePrioritizeReplaceResponseGeoRestrictionsJSON contains the JSON
// metadata for the struct
// [CustomCertificatePrioritizeReplaceResponseGeoRestrictions]
type customCertificatePrioritizeReplaceResponseGeoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeReplaceResponseGeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabel string

const (
	CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabelUs              CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabel = "us"
	CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabelEu              CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabel = "eu"
	CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabelHighestSecurity CustomCertificatePrioritizeReplaceResponseGeoRestrictionsLabel = "highest_security"
)

type CustomCertificatePrioritizeReplaceResponseKeylessServer struct {
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
	Status CustomCertificatePrioritizeReplaceResponseKeylessServerStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel CustomCertificatePrioritizeReplaceResponseKeylessServerTunnel `json:"tunnel"`
	JSON   customCertificatePrioritizeReplaceResponseKeylessServerJSON   `json:"-"`
}

// customCertificatePrioritizeReplaceResponseKeylessServerJSON contains the JSON
// metadata for the struct
// [CustomCertificatePrioritizeReplaceResponseKeylessServer]
type customCertificatePrioritizeReplaceResponseKeylessServerJSON struct {
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

func (r *CustomCertificatePrioritizeReplaceResponseKeylessServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type CustomCertificatePrioritizeReplaceResponseKeylessServerStatus string

const (
	CustomCertificatePrioritizeReplaceResponseKeylessServerStatusActive  CustomCertificatePrioritizeReplaceResponseKeylessServerStatus = "active"
	CustomCertificatePrioritizeReplaceResponseKeylessServerStatusDeleted CustomCertificatePrioritizeReplaceResponseKeylessServerStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type CustomCertificatePrioritizeReplaceResponseKeylessServerTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                                            `json:"vnet_id,required"`
	JSON   customCertificatePrioritizeReplaceResponseKeylessServerTunnelJSON `json:"-"`
}

// customCertificatePrioritizeReplaceResponseKeylessServerTunnelJSON contains the
// JSON metadata for the struct
// [CustomCertificatePrioritizeReplaceResponseKeylessServerTunnel]
type customCertificatePrioritizeReplaceResponseKeylessServerTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeReplaceResponseKeylessServerTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificatePrioritizeReplaceParams struct {
	// Array of ordered certificates.
	Certificates param.Field[[]CustomCertificatePrioritizeReplaceParamsCertificate] `json:"certificates,required"`
}

func (r CustomCertificatePrioritizeReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificatePrioritizeReplaceParamsCertificate struct {
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r CustomCertificatePrioritizeReplaceParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificatePrioritizeReplaceResponseEnvelope struct {
	Errors   []CustomCertificatePrioritizeReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomCertificatePrioritizeReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomCertificatePrioritizeReplaceResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomCertificatePrioritizeReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomCertificatePrioritizeReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customCertificatePrioritizeReplaceResponseEnvelopeJSON       `json:"-"`
}

// customCertificatePrioritizeReplaceResponseEnvelopeJSON contains the JSON
// metadata for the struct [CustomCertificatePrioritizeReplaceResponseEnvelope]
type customCertificatePrioritizeReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificatePrioritizeReplaceResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    customCertificatePrioritizeReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// customCertificatePrioritizeReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CustomCertificatePrioritizeReplaceResponseEnvelopeErrors]
type customCertificatePrioritizeReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificatePrioritizeReplaceResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    customCertificatePrioritizeReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// customCertificatePrioritizeReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CustomCertificatePrioritizeReplaceResponseEnvelopeMessages]
type customCertificatePrioritizeReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomCertificatePrioritizeReplaceResponseEnvelopeSuccess bool

const (
	CustomCertificatePrioritizeReplaceResponseEnvelopeSuccessTrue CustomCertificatePrioritizeReplaceResponseEnvelopeSuccess = true
)

type CustomCertificatePrioritizeReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       customCertificatePrioritizeReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// customCertificatePrioritizeReplaceResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [CustomCertificatePrioritizeReplaceResponseEnvelopeResultInfo]
type customCertificatePrioritizeReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
