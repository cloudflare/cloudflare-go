// File generated from our OpenAPI spec by Stainless.

package custom_certificates

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// PrioritizeService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPrioritizeService] method instead.
type PrioritizeService struct {
	Options []option.RequestOption
}

// NewPrioritizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrioritizeService(opts ...option.RequestOption) (r *PrioritizeService) {
	r = &PrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *PrioritizeService) Update(ctx context.Context, params PrioritizeUpdateParams, opts ...option.RequestOption) (res *[]PrioritizeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrioritizeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrioritizeUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod PrioritizeUpdateResponseBundleMethod `json:"bundle_method,required"`
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
	Status PrioritizeUpdateResponseStatus `json:"status,required"`
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
	GeoRestrictions PrioritizeUpdateResponseGeoRestrictions `json:"geo_restrictions"`
	KeylessServer   PrioritizeUpdateResponseKeylessServer   `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                       `json:"policy"`
	JSON   prioritizeUpdateResponseJSON `json:"-"`
}

// prioritizeUpdateResponseJSON contains the JSON metadata for the struct
// [PrioritizeUpdateResponse]
type prioritizeUpdateResponseJSON struct {
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

func (r *PrioritizeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type PrioritizeUpdateResponseBundleMethod string

const (
	PrioritizeUpdateResponseBundleMethodUbiquitous PrioritizeUpdateResponseBundleMethod = "ubiquitous"
	PrioritizeUpdateResponseBundleMethodOptimal    PrioritizeUpdateResponseBundleMethod = "optimal"
	PrioritizeUpdateResponseBundleMethodForce      PrioritizeUpdateResponseBundleMethod = "force"
)

// Status of the zone's custom SSL.
type PrioritizeUpdateResponseStatus string

const (
	PrioritizeUpdateResponseStatusActive       PrioritizeUpdateResponseStatus = "active"
	PrioritizeUpdateResponseStatusExpired      PrioritizeUpdateResponseStatus = "expired"
	PrioritizeUpdateResponseStatusDeleted      PrioritizeUpdateResponseStatus = "deleted"
	PrioritizeUpdateResponseStatusPending      PrioritizeUpdateResponseStatus = "pending"
	PrioritizeUpdateResponseStatusInitializing PrioritizeUpdateResponseStatus = "initializing"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type PrioritizeUpdateResponseGeoRestrictions struct {
	Label PrioritizeUpdateResponseGeoRestrictionsLabel `json:"label"`
	JSON  prioritizeUpdateResponseGeoRestrictionsJSON  `json:"-"`
}

// prioritizeUpdateResponseGeoRestrictionsJSON contains the JSON metadata for the
// struct [PrioritizeUpdateResponseGeoRestrictions]
type prioritizeUpdateResponseGeoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseGeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseGeoRestrictionsJSON) RawJSON() string {
	return r.raw
}

type PrioritizeUpdateResponseGeoRestrictionsLabel string

const (
	PrioritizeUpdateResponseGeoRestrictionsLabelUs              PrioritizeUpdateResponseGeoRestrictionsLabel = "us"
	PrioritizeUpdateResponseGeoRestrictionsLabelEu              PrioritizeUpdateResponseGeoRestrictionsLabel = "eu"
	PrioritizeUpdateResponseGeoRestrictionsLabelHighestSecurity PrioritizeUpdateResponseGeoRestrictionsLabel = "highest_security"
)

type PrioritizeUpdateResponseKeylessServer struct {
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
	Status PrioritizeUpdateResponseKeylessServerStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel PrioritizeUpdateResponseKeylessServerTunnel `json:"tunnel"`
	JSON   prioritizeUpdateResponseKeylessServerJSON   `json:"-"`
}

// prioritizeUpdateResponseKeylessServerJSON contains the JSON metadata for the
// struct [PrioritizeUpdateResponseKeylessServer]
type prioritizeUpdateResponseKeylessServerJSON struct {
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

func (r *PrioritizeUpdateResponseKeylessServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseKeylessServerJSON) RawJSON() string {
	return r.raw
}

// Status of the Keyless SSL.
type PrioritizeUpdateResponseKeylessServerStatus string

const (
	PrioritizeUpdateResponseKeylessServerStatusActive  PrioritizeUpdateResponseKeylessServerStatus = "active"
	PrioritizeUpdateResponseKeylessServerStatusDeleted PrioritizeUpdateResponseKeylessServerStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type PrioritizeUpdateResponseKeylessServerTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                          `json:"vnet_id,required"`
	JSON   prioritizeUpdateResponseKeylessServerTunnelJSON `json:"-"`
}

// prioritizeUpdateResponseKeylessServerTunnelJSON contains the JSON metadata for
// the struct [PrioritizeUpdateResponseKeylessServerTunnel]
type prioritizeUpdateResponseKeylessServerTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseKeylessServerTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseKeylessServerTunnelJSON) RawJSON() string {
	return r.raw
}

type PrioritizeUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Array of ordered certificates.
	Certificates param.Field[[]PrioritizeUpdateParamsCertificate] `json:"certificates,required"`
}

func (r PrioritizeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrioritizeUpdateParamsCertificate struct {
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r PrioritizeUpdateParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrioritizeUpdateResponseEnvelope struct {
	Errors   []PrioritizeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrioritizeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []PrioritizeUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PrioritizeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PrioritizeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       prioritizeUpdateResponseEnvelopeJSON       `json:"-"`
}

// prioritizeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrioritizeUpdateResponseEnvelope]
type prioritizeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrioritizeUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    prioritizeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// prioritizeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrioritizeUpdateResponseEnvelopeErrors]
type prioritizeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrioritizeUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    prioritizeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// prioritizeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PrioritizeUpdateResponseEnvelopeMessages]
type prioritizeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrioritizeUpdateResponseEnvelopeSuccess bool

const (
	PrioritizeUpdateResponseEnvelopeSuccessTrue PrioritizeUpdateResponseEnvelopeSuccess = true
)

type PrioritizeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       prioritizeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// prioritizeUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PrioritizeUpdateResponseEnvelopeResultInfo]
type prioritizeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
