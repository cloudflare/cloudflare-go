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

// ZoneKeylessCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneKeylessCertificateService]
// method instead.
type ZoneKeylessCertificateService struct {
	Options []option.RequestOption
}

// NewZoneKeylessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneKeylessCertificateService(opts ...option.RequestOption) (r *ZoneKeylessCertificateService) {
	r = &ZoneKeylessCertificateService{}
	r.Options = opts
	return
}

// Get details for one Keyless SSL configuration.
func (r *ZoneKeylessCertificateService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *KeylessResponseSingle1Ulr3eBh, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This will update attributes of a Keyless SSL. Consists of one or more of the
// following: host,name,port.
func (r *ZoneKeylessCertificateService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneKeylessCertificateUpdateParams, opts ...option.RequestOption) (res *KeylessResponseSingle1Ulr3eBh, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete Keyless SSL Configuration
func (r *ZoneKeylessCertificateService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *KeylessResponseSingleIDSi9DlyBw, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Keyless SSL Configuration
func (r *ZoneKeylessCertificateService) KeylessSslForAZoneNewKeylessSslConfiguration(ctx context.Context, zoneIdentifier string, body ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParams, opts ...option.RequestOption) (res *KeylessResponseSingle1Ulr3eBh, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all Keyless SSL configurations for a given zone.
func (r *ZoneKeylessCertificateService) KeylessSslForAZoneListKeylessSslConfigurations(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *KeylessResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type KeylessResponseCollection struct {
	Errors     []KeylessResponseCollectionError    `json:"errors"`
	Messages   []KeylessResponseCollectionMessage  `json:"messages"`
	Result     []KeylessResponseCollectionResult   `json:"result"`
	ResultInfo KeylessResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success KeylessResponseCollectionSuccess `json:"success"`
	JSON    keylessResponseCollectionJSON    `json:"-"`
}

// keylessResponseCollectionJSON contains the JSON metadata for the struct
// [KeylessResponseCollection]
type keylessResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseCollectionError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    keylessResponseCollectionErrorJSON `json:"-"`
}

// keylessResponseCollectionErrorJSON contains the JSON metadata for the struct
// [KeylessResponseCollectionError]
type keylessResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseCollectionMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    keylessResponseCollectionMessageJSON `json:"-"`
}

// keylessResponseCollectionMessageJSON contains the JSON metadata for the struct
// [KeylessResponseCollectionMessage]
type keylessResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseCollectionResult struct {
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
	Status KeylessResponseCollectionResultStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel KeylessResponseCollectionResultTunnel `json:"tunnel"`
	JSON   keylessResponseCollectionResultJSON   `json:"-"`
}

// keylessResponseCollectionResultJSON contains the JSON metadata for the struct
// [KeylessResponseCollectionResult]
type keylessResponseCollectionResultJSON struct {
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

func (r *KeylessResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type KeylessResponseCollectionResultStatus string

const (
	KeylessResponseCollectionResultStatusActive  KeylessResponseCollectionResultStatus = "active"
	KeylessResponseCollectionResultStatusDeleted KeylessResponseCollectionResultStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type KeylessResponseCollectionResultTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                    `json:"vnet_id,required"`
	JSON   keylessResponseCollectionResultTunnelJSON `json:"-"`
}

// keylessResponseCollectionResultTunnelJSON contains the JSON metadata for the
// struct [KeylessResponseCollectionResultTunnel]
type keylessResponseCollectionResultTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseCollectionResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       keylessResponseCollectionResultInfoJSON `json:"-"`
}

// keylessResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [KeylessResponseCollectionResultInfo]
type keylessResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KeylessResponseCollectionSuccess bool

const (
	KeylessResponseCollectionSuccessTrue KeylessResponseCollectionSuccess = true
)

type KeylessResponseSingle1Ulr3eBh struct {
	Errors   []KeylessResponseSingle1Ulr3eBhError   `json:"errors"`
	Messages []KeylessResponseSingle1Ulr3eBhMessage `json:"messages"`
	Result   KeylessResponseSingle1Ulr3eBhResult    `json:"result"`
	// Whether the API call was successful
	Success KeylessResponseSingle1Ulr3eBhSuccess `json:"success"`
	JSON    keylessResponseSingle1Ulr3eBhJSON    `json:"-"`
}

// keylessResponseSingle1Ulr3eBhJSON contains the JSON metadata for the struct
// [KeylessResponseSingle1Ulr3eBh]
type keylessResponseSingle1Ulr3eBhJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingle1Ulr3eBh) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseSingle1Ulr3eBhError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    keylessResponseSingle1Ulr3eBhErrorJSON `json:"-"`
}

// keylessResponseSingle1Ulr3eBhErrorJSON contains the JSON metadata for the struct
// [KeylessResponseSingle1Ulr3eBhError]
type keylessResponseSingle1Ulr3eBhErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingle1Ulr3eBhError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseSingle1Ulr3eBhMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    keylessResponseSingle1Ulr3eBhMessageJSON `json:"-"`
}

// keylessResponseSingle1Ulr3eBhMessageJSON contains the JSON metadata for the
// struct [KeylessResponseSingle1Ulr3eBhMessage]
type keylessResponseSingle1Ulr3eBhMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingle1Ulr3eBhMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseSingle1Ulr3eBhResult struct {
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
	Status KeylessResponseSingle1Ulr3eBhResultStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel KeylessResponseSingle1Ulr3eBhResultTunnel `json:"tunnel"`
	JSON   keylessResponseSingle1Ulr3eBhResultJSON   `json:"-"`
}

// keylessResponseSingle1Ulr3eBhResultJSON contains the JSON metadata for the
// struct [KeylessResponseSingle1Ulr3eBhResult]
type keylessResponseSingle1Ulr3eBhResultJSON struct {
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

func (r *KeylessResponseSingle1Ulr3eBhResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type KeylessResponseSingle1Ulr3eBhResultStatus string

const (
	KeylessResponseSingle1Ulr3eBhResultStatusActive  KeylessResponseSingle1Ulr3eBhResultStatus = "active"
	KeylessResponseSingle1Ulr3eBhResultStatusDeleted KeylessResponseSingle1Ulr3eBhResultStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type KeylessResponseSingle1Ulr3eBhResultTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                        `json:"vnet_id,required"`
	JSON   keylessResponseSingle1Ulr3eBhResultTunnelJSON `json:"-"`
}

// keylessResponseSingle1Ulr3eBhResultTunnelJSON contains the JSON metadata for the
// struct [KeylessResponseSingle1Ulr3eBhResultTunnel]
type keylessResponseSingle1Ulr3eBhResultTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingle1Ulr3eBhResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KeylessResponseSingle1Ulr3eBhSuccess bool

const (
	KeylessResponseSingle1Ulr3eBhSuccessTrue KeylessResponseSingle1Ulr3eBhSuccess = true
)

type KeylessResponseSingleIDSi9DlyBw struct {
	Errors   []KeylessResponseSingleIDSi9DlyBwError   `json:"errors"`
	Messages []KeylessResponseSingleIDSi9DlyBwMessage `json:"messages"`
	Result   KeylessResponseSingleIDSi9DlyBwResult    `json:"result"`
	// Whether the API call was successful
	Success KeylessResponseSingleIDSi9DlyBwSuccess `json:"success"`
	JSON    keylessResponseSingleIDSi9DlyBwJSON    `json:"-"`
}

// keylessResponseSingleIDSi9DlyBwJSON contains the JSON metadata for the struct
// [KeylessResponseSingleIDSi9DlyBw]
type keylessResponseSingleIDSi9DlyBwJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingleIDSi9DlyBw) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseSingleIDSi9DlyBwError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    keylessResponseSingleIDSi9DlyBwErrorJSON `json:"-"`
}

// keylessResponseSingleIDSi9DlyBwErrorJSON contains the JSON metadata for the
// struct [KeylessResponseSingleIDSi9DlyBwError]
type keylessResponseSingleIDSi9DlyBwErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingleIDSi9DlyBwError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseSingleIDSi9DlyBwMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    keylessResponseSingleIDSi9DlyBwMessageJSON `json:"-"`
}

// keylessResponseSingleIDSi9DlyBwMessageJSON contains the JSON metadata for the
// struct [KeylessResponseSingleIDSi9DlyBwMessage]
type keylessResponseSingleIDSi9DlyBwMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingleIDSi9DlyBwMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeylessResponseSingleIDSi9DlyBwResult struct {
	// Identifier
	ID   string                                    `json:"id"`
	JSON keylessResponseSingleIDSi9DlyBwResultJSON `json:"-"`
}

// keylessResponseSingleIDSi9DlyBwResultJSON contains the JSON metadata for the
// struct [KeylessResponseSingleIDSi9DlyBwResult]
type keylessResponseSingleIDSi9DlyBwResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingleIDSi9DlyBwResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KeylessResponseSingleIDSi9DlyBwSuccess bool

const (
	KeylessResponseSingleIDSi9DlyBwSuccessTrue KeylessResponseSingleIDSi9DlyBwSuccess = true
)

type ZoneKeylessCertificateUpdateParams struct {
	// Whether or not the Keyless SSL is on or off.
	Enabled param.Field[bool] `json:"enabled"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host" format:"hostname"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[ZoneKeylessCertificateUpdateParamsTunnel] `json:"tunnel"`
}

func (r ZoneKeylessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateUpdateParamsTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r ZoneKeylessCertificateUpdateParamsTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParams struct {
	// The zone's SSL certificate or SSL certificate and intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod] `json:"bundle_method"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsTunnel] `json:"tunnel"`
}

func (r ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod string

const (
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethodUbiquitous ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod = "ubiquitous"
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethodOptimal    ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod = "optimal"
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethodForce      ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod = "force"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
