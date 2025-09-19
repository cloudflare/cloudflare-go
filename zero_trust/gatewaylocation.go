// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// GatewayLocationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayLocationService] method instead.
type GatewayLocationService struct {
	Options []option.RequestOption
}

// NewGatewayLocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayLocationService(opts ...option.RequestOption) (r *GatewayLocationService) {
	r = &GatewayLocationService{}
	r.Options = opts
	return
}

// Create a new Zero Trust Gateway location.
func (r *GatewayLocationService) New(ctx context.Context, params GatewayLocationNewParams, opts ...option.RequestOption) (res *Location, err error) {
	var env GatewayLocationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Update(ctx context.Context, locationID string, params GatewayLocationUpdateParams, opts ...option.RequestOption) (res *Location, err error) {
	var env GatewayLocationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations/%s", params.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Zero Trust Gateway locations for an account.
func (r *GatewayLocationService) List(ctx context.Context, query GatewayLocationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Location], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// List Zero Trust Gateway locations for an account.
func (r *GatewayLocationService) ListAutoPaging(ctx context.Context, query GatewayLocationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Location] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Delete(ctx context.Context, locationID string, body GatewayLocationDeleteParams, opts ...option.RequestOption) (res *GatewayLocationDeleteResponse, err error) {
	var env GatewayLocationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations/%s", body.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single Zero Trust Gateway location.
func (r *GatewayLocationService) Get(ctx context.Context, locationID string, query GatewayLocationGetParams, opts ...option.RequestOption) (res *Location, err error) {
	var env GatewayLocationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations/%s", query.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DOHEndpoint struct {
	// Indicate whether the DOH endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// Specify the list of allowed source IP network ranges for this endpoint. When the
	// list is empty, the endpoint allows all source IPs. The list takes effect only if
	// the endpoint is enabled for this location.
	Networks []IPNetwork `json:"networks,nullable"`
	// Specify whether the DOH endpoint requires user identity authentication.
	RequireToken bool            `json:"require_token"`
	JSON         dohEndpointJSON `json:"-"`
}

// dohEndpointJSON contains the JSON metadata for the struct [DOHEndpoint]
type dohEndpointJSON struct {
	Enabled      apijson.Field
	Networks     apijson.Field
	RequireToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DOHEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dohEndpointJSON) RawJSON() string {
	return r.raw
}

type DOHEndpointParam struct {
	// Indicate whether the DOH endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// Specify the list of allowed source IP network ranges for this endpoint. When the
	// list is empty, the endpoint allows all source IPs. The list takes effect only if
	// the endpoint is enabled for this location.
	Networks param.Field[[]IPNetworkParam] `json:"networks"`
	// Specify whether the DOH endpoint requires user identity authentication.
	RequireToken param.Field[bool] `json:"require_token"`
}

func (r DOHEndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DOTEndpoint struct {
	// Indicate whether the DOT endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// Specify the list of allowed source IP network ranges for this endpoint. When the
	// list is empty, the endpoint allows all source IPs. The list takes effect only if
	// the endpoint is enabled for this location.
	Networks []IPNetwork     `json:"networks,nullable"`
	JSON     dotEndpointJSON `json:"-"`
}

// dotEndpointJSON contains the JSON metadata for the struct [DOTEndpoint]
type dotEndpointJSON struct {
	Enabled     apijson.Field
	Networks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DOTEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dotEndpointJSON) RawJSON() string {
	return r.raw
}

type DOTEndpointParam struct {
	// Indicate whether the DOT endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// Specify the list of allowed source IP network ranges for this endpoint. When the
	// list is empty, the endpoint allows all source IPs. The list takes effect only if
	// the endpoint is enabled for this location.
	Networks param.Field[[]IPNetworkParam] `json:"networks"`
}

func (r DOTEndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure the destination endpoints for this location.
type Endpoint struct {
	DOH  DOHEndpoint  `json:"doh,required"`
	DOT  DOTEndpoint  `json:"dot,required"`
	IPV4 IPV4Endpoint `json:"ipv4,required"`
	IPV6 IPV6Endpoint `json:"ipv6,required"`
	JSON endpointJSON `json:"-"`
}

// endpointJSON contains the JSON metadata for the struct [Endpoint]
type endpointJSON struct {
	DOH         apijson.Field
	DOT         apijson.Field
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Endpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointJSON) RawJSON() string {
	return r.raw
}

// Configure the destination endpoints for this location.
type EndpointParam struct {
	DOH  param.Field[DOHEndpointParam]  `json:"doh,required"`
	DOT  param.Field[DOTEndpointParam]  `json:"dot,required"`
	IPV4 param.Field[IPV4EndpointParam] `json:"ipv4,required"`
	IPV6 param.Field[IPV6EndpointParam] `json:"ipv6,required"`
}

func (r EndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPNetwork struct {
	// Specify the IP address or IP CIDR.
	Network string        `json:"network,required"`
	JSON    ipNetworkJSON `json:"-"`
}

// ipNetworkJSON contains the JSON metadata for the struct [IPNetwork]
type ipNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipNetworkJSON) RawJSON() string {
	return r.raw
}

type IPNetworkParam struct {
	// Specify the IP address or IP CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r IPNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPV4Endpoint struct {
	// Indicate whether the IPv4 endpoint is enabled for this location.
	Enabled bool             `json:"enabled"`
	JSON    ipv4EndpointJSON `json:"-"`
}

// ipv4EndpointJSON contains the JSON metadata for the struct [IPV4Endpoint]
type ipv4EndpointJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPV4Endpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipv4EndpointJSON) RawJSON() string {
	return r.raw
}

type IPV4EndpointParam struct {
	// Indicate whether the IPv4 endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r IPV4EndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPV6Endpoint struct {
	// Indicate whether the IPV6 endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// Specify the list of allowed source IPv6 network ranges for this endpoint. When
	// the list is empty, the endpoint allows all source IPs. The list takes effect
	// only if the endpoint is enabled for this location.
	Networks []IPV6Network    `json:"networks,nullable"`
	JSON     ipv6EndpointJSON `json:"-"`
}

// ipv6EndpointJSON contains the JSON metadata for the struct [IPV6Endpoint]
type ipv6EndpointJSON struct {
	Enabled     apijson.Field
	Networks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPV6Endpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipv6EndpointJSON) RawJSON() string {
	return r.raw
}

type IPV6EndpointParam struct {
	// Indicate whether the IPV6 endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// Specify the list of allowed source IPv6 network ranges for this endpoint. When
	// the list is empty, the endpoint allows all source IPs. The list takes effect
	// only if the endpoint is enabled for this location.
	Networks param.Field[[]IPV6NetworkParam] `json:"networks"`
}

func (r IPV6EndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPV6Network struct {
	// Specify the IPv6 address or IPv6 CIDR.
	Network string          `json:"network,required"`
	JSON    ipv6NetworkJSON `json:"-"`
}

// ipv6NetworkJSON contains the JSON metadata for the struct [IPV6Network]
type ipv6NetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPV6Network) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipv6NetworkJSON) RawJSON() string {
	return r.raw
}

type IPV6NetworkParam struct {
	// Specify the IPv6 address or IPv6 CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r IPV6NetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Location struct {
	ID string `json:"id"`
	// Indicate whether this location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// Indicate the identifier of the pair of IPv4 addresses assigned to this location.
	DNSDestinationIPsID string `json:"dns_destination_ips_id"`
	// Specify the UUID of the IPv6 block brought to the gateway so that this
	// location's IPv6 address is allocated from the Bring Your Own IPv6 (BYOIPv6)
	// block rather than the standard Cloudflare IPv6 block.
	DNSDestinationIPV6BlockID string `json:"dns_destination_ipv6_block_id,nullable"`
	// Specify the DNS over HTTPS domain that receives DNS requests. Gateway
	// automatically generates this value.
	DOHSubdomain string `json:"doh_subdomain"`
	// Indicate whether the location must resolve EDNS queries.
	ECSSupport bool `json:"ecs_support"`
	// Configure the destination endpoints for this location.
	Endpoints Endpoint `json:"endpoints,nullable"`
	// Defines the automatically generated IPv6 destination IP assigned to this
	// location. Gateway counts all DNS requests sent to this IP as requests under this
	// location.
	IP string `json:"ip"`
	// Show the primary destination IPv4 address from the pair identified
	// dns_destination_ips_id. This field read-only.
	IPV4Destination string `json:"ipv4_destination"`
	// Show the backup destination IPv4 address from the pair identified
	// dns_destination_ips_id. This field read-only.
	IPV4DestinationBackup string `json:"ipv4_destination_backup"`
	// Specify the location name.
	Name string `json:"name"`
	// Specify the list of network ranges from which requests at this location
	// originate. The list takes effect only if it is non-empty and the IPv4 endpoint
	// is enabled for this location.
	Networks  []LocationNetwork `json:"networks,nullable"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      locationJSON      `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	ID                        apijson.Field
	ClientDefault             apijson.Field
	CreatedAt                 apijson.Field
	DNSDestinationIPsID       apijson.Field
	DNSDestinationIPV6BlockID apijson.Field
	DOHSubdomain              apijson.Field
	ECSSupport                apijson.Field
	Endpoints                 apijson.Field
	IP                        apijson.Field
	IPV4Destination           apijson.Field
	IPV4DestinationBackup     apijson.Field
	Name                      apijson.Field
	Networks                  apijson.Field
	UpdatedAt                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationJSON) RawJSON() string {
	return r.raw
}

type LocationNetwork struct {
	// Specify the IPv4 address or IPv4 CIDR. Limit IPv4 CIDRs to a maximum of /24.
	Network string              `json:"network,required"`
	JSON    locationNetworkJSON `json:"-"`
}

// locationNetworkJSON contains the JSON metadata for the struct [LocationNetwork]
type locationNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationNetworkJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationDeleteResponse = interface{}

type GatewayLocationNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Specify the location name.
	Name param.Field[string] `json:"name,required"`
	// Indicate whether this location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// Specify the identifier of the pair of IPv4 addresses assigned to this location.
	// When creating a location, if this field is absent or set to null, the pair of
	// shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned.
	// When updating a location, if this field is absent or set to null, the
	// pre-assigned pair remains unchanged.
	DNSDestinationIPsID param.Field[string] `json:"dns_destination_ips_id"`
	// Indicate whether the location must resolve EDNS queries.
	ECSSupport param.Field[bool] `json:"ecs_support"`
	// Configure the destination endpoints for this location.
	Endpoints param.Field[EndpointParam] `json:"endpoints"`
	// Specify the list of network ranges from which requests at this location
	// originate. The list takes effect only if it is non-empty and the IPv4 endpoint
	// is enabled for this location.
	Networks param.Field[[]GatewayLocationNewParamsNetwork] `json:"networks"`
}

func (r GatewayLocationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsNetwork struct {
	// Specify the IPv4 address or IPv4 CIDR. Limit IPv4 CIDRs to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayLocationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Location                                  `json:"result"`
	JSON    gatewayLocationNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLocationNewResponseEnvelope]
type gatewayLocationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayLocationNewResponseEnvelopeSuccess bool

const (
	GatewayLocationNewResponseEnvelopeSuccessTrue GatewayLocationNewResponseEnvelopeSuccess = true
)

func (r GatewayLocationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayLocationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayLocationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Specify the location name.
	Name param.Field[string] `json:"name,required"`
	// Indicate whether this location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// Specify the identifier of the pair of IPv4 addresses assigned to this location.
	// When creating a location, if this field is absent or set to null, the pair of
	// shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned.
	// When updating a location, if this field is absent or set to null, the
	// pre-assigned pair remains unchanged.
	DNSDestinationIPsID param.Field[string] `json:"dns_destination_ips_id"`
	// Indicate whether the location must resolve EDNS queries.
	ECSSupport param.Field[bool] `json:"ecs_support"`
	// Configure the destination endpoints for this location.
	Endpoints param.Field[EndpointParam] `json:"endpoints"`
	// Specify the list of network ranges from which requests at this location
	// originate. The list takes effect only if it is non-empty and the IPv4 endpoint
	// is enabled for this location.
	Networks param.Field[[]GatewayLocationUpdateParamsNetwork] `json:"networks"`
}

func (r GatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsNetwork struct {
	// Specify the IPv4 address or IPv4 CIDR. Limit IPv4 CIDRs to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayLocationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Location                                     `json:"result"`
	JSON    gatewayLocationUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationUpdateResponseEnvelope]
type gatewayLocationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayLocationUpdateResponseEnvelopeSuccess bool

const (
	GatewayLocationUpdateResponseEnvelopeSuccessTrue GatewayLocationUpdateResponseEnvelopeSuccess = true
)

func (r GatewayLocationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayLocationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayLocationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayLocationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayLocationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayLocationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayLocationDeleteResponse                `json:"result"`
	JSON    gatewayLocationDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationDeleteResponseEnvelope]
type gatewayLocationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayLocationDeleteResponseEnvelopeSuccess bool

const (
	GatewayLocationDeleteResponseEnvelopeSuccessTrue GatewayLocationDeleteResponseEnvelopeSuccess = true
)

func (r GatewayLocationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayLocationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayLocationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayLocationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayLocationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Location                                  `json:"result"`
	JSON    gatewayLocationGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLocationGetResponseEnvelope]
type gatewayLocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayLocationGetResponseEnvelopeSuccess bool

const (
	GatewayLocationGetResponseEnvelopeSuccessTrue GatewayLocationGetResponseEnvelopeSuccess = true
)

func (r GatewayLocationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayLocationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
