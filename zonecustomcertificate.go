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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *ZoneCustomCertificateService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneCustomCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing
// a configuration for sni_custom certificates will result in a new resource id
// being returned, and the previous one being deleted.
func (r *ZoneCustomCertificateService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneCustomCertificateUpdateParams, opts ...option.RequestOption) (res *ZoneCustomCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Remove a SSL certificate from a zone.
func (r *ZoneCustomCertificateService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneCustomCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload a new SSL certificate for a zone.
func (r *ZoneCustomCertificateService) CustomSslForAZoneNewSslConfiguration(ctx context.Context, zoneIdentifier string, body ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationParams, opts ...option.RequestOption) (res *ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, and filter all of your custom SSL certificates. The higher
// priority will break ties across overlapping 'legacy_custom' certificates, but
// 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
func (r *ZoneCustomCertificateService) CustomSslForAZoneListSslConfigurations(ctx context.Context, zoneIdentifier string, query ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsParams, opts ...option.RequestOption) (res *shared.Page[ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneIdentifier)
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

type ZoneCustomCertificateGetResponse struct {
	Errors   []ZoneCustomCertificateGetResponseError   `json:"errors"`
	Messages []ZoneCustomCertificateGetResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomCertificateGetResponseSuccess `json:"success"`
	JSON    zoneCustomCertificateGetResponseJSON    `json:"-"`
}

// zoneCustomCertificateGetResponseJSON contains the JSON metadata for the struct
// [ZoneCustomCertificateGetResponse]
type zoneCustomCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneCustomCertificateGetResponseErrorJSON `json:"-"`
}

// zoneCustomCertificateGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateGetResponseError]
type zoneCustomCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneCustomCertificateGetResponseMessageJSON `json:"-"`
}

// zoneCustomCertificateGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateGetResponseMessage]
type zoneCustomCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomCertificateGetResponseSuccess bool

const (
	ZoneCustomCertificateGetResponseSuccessTrue ZoneCustomCertificateGetResponseSuccess = true
)

type ZoneCustomCertificateUpdateResponse struct {
	Errors   []ZoneCustomCertificateUpdateResponseError   `json:"errors"`
	Messages []ZoneCustomCertificateUpdateResponseMessage `json:"messages"`
	Result   interface{}                                  `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomCertificateUpdateResponseSuccess `json:"success"`
	JSON    zoneCustomCertificateUpdateResponseJSON    `json:"-"`
}

// zoneCustomCertificateUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateUpdateResponse]
type zoneCustomCertificateUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneCustomCertificateUpdateResponseErrorJSON `json:"-"`
}

// zoneCustomCertificateUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateUpdateResponseError]
type zoneCustomCertificateUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneCustomCertificateUpdateResponseMessageJSON `json:"-"`
}

// zoneCustomCertificateUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneCustomCertificateUpdateResponseMessage]
type zoneCustomCertificateUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomCertificateUpdateResponseSuccess bool

const (
	ZoneCustomCertificateUpdateResponseSuccessTrue ZoneCustomCertificateUpdateResponseSuccess = true
)

type ZoneCustomCertificateDeleteResponse struct {
	Errors   []ZoneCustomCertificateDeleteResponseError   `json:"errors"`
	Messages []ZoneCustomCertificateDeleteResponseMessage `json:"messages"`
	Result   ZoneCustomCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomCertificateDeleteResponseSuccess `json:"success"`
	JSON    zoneCustomCertificateDeleteResponseJSON    `json:"-"`
}

// zoneCustomCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateDeleteResponse]
type zoneCustomCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneCustomCertificateDeleteResponseErrorJSON `json:"-"`
}

// zoneCustomCertificateDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateDeleteResponseError]
type zoneCustomCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneCustomCertificateDeleteResponseMessageJSON `json:"-"`
}

// zoneCustomCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneCustomCertificateDeleteResponseMessage]
type zoneCustomCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateDeleteResponseResult struct {
	// Identifier
	ID   string                                        `json:"id"`
	JSON zoneCustomCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneCustomCertificateDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateDeleteResponseResult]
type zoneCustomCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomCertificateDeleteResponseSuccess bool

const (
	ZoneCustomCertificateDeleteResponseSuccessTrue ZoneCustomCertificateDeleteResponseSuccess = true
)

type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponse struct {
	Errors   []ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseError   `json:"errors"`
	Messages []ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessage `json:"messages"`
	Result   interface{}                                                                `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseSuccess `json:"success"`
	JSON    zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseJSON    `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseJSON contains
// the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponse]
type zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseErrorJSON `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseError]
type zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessageJSON `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessage]
type zoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseSuccess bool

const (
	ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseSuccessTrue ZoneCustomCertificateCustomSslForAZoneNewSslConfigurationResponseSuccess = true
)

type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethod `json:"bundle_method,required"`
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
	Status ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus `json:"status,required"`
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
	GeoRestrictions ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictions `json:"geo_restrictions"`
	KeylessServer   ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServer   `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                                                                  `json:"policy"`
	JSON   zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseJSON `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseJSON contains
// the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponse]
type zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseJSON struct {
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

func (r *ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethod string

const (
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethodUbiquitous ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethod = "ubiquitous"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethodOptimal    ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethod = "optimal"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethodForce      ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseBundleMethod = "force"
)

// Status of the zone's custom SSL.
type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus string

const (
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatusActive       ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus = "active"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatusExpired      ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus = "expired"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatusDeleted      ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus = "deleted"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatusPending      ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus = "pending"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatusInitializing ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseStatus = "initializing"
)

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictions struct {
	Label ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabel `json:"label"`
	JSON  zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsJSON  `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictions]
type zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabel string

const (
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabelUs              ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabel = "us"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabelEu              ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabel = "eu"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabelHighestSecurity ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseGeoRestrictionsLabel = "highest_security"
)

type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServer struct {
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
	Status ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnel `json:"tunnel"`
	JSON   zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerJSON   `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServer]
type zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerJSON struct {
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

func (r *ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerStatus string

const (
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerStatusActive  ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerStatus = "active"
	ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerStatusDeleted ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                                                                     `json:"vnet_id,required"`
	JSON   zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnelJSON `json:"-"`
}

// zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnelJSON
// contains the JSON metadata for the struct
// [ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnel]
type zoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateCustomSslForAZoneListSslConfigurationsResponseKeylessServerTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
