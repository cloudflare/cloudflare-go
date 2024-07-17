// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
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

// Creates a new Zero Trust Gateway location.
func (r *GatewayLocationService) New(ctx context.Context, params GatewayLocationNewParams, opts ...option.RequestOption) (res *Location, err error) {
	var env GatewayLocationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// Updates a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Update(ctx context.Context, locationID string, params GatewayLocationUpdateParams, opts ...option.RequestOption) (res *Location, err error) {
	var env GatewayLocationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// Fetches Zero Trust Gateway locations for an account.
func (r *GatewayLocationService) List(ctx context.Context, query GatewayLocationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Location], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// Fetches Zero Trust Gateway locations for an account.
func (r *GatewayLocationService) ListAutoPaging(ctx context.Context, query GatewayLocationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Location] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Delete(ctx context.Context, locationID string, body GatewayLocationDeleteParams, opts ...option.RequestOption) (res *GatewayLocationDeleteResponseUnion, err error) {
	var env GatewayLocationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// Fetches a single Zero Trust Gateway location.
func (r *GatewayLocationService) Get(ctx context.Context, locationID string, query GatewayLocationGetParams, opts ...option.RequestOption) (res *Location, err error) {
	var env GatewayLocationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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

type Location struct {
	ID string `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The identifier of the pair of IPv4 addresses assigned to this location.
	DNSDestinationIPsID string `json:"dns_destination_ips_id"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DOHSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	ECSSupport bool `json:"ecs_support"`
	// The destination endpoints configured for this location. When updating a
	// location, if this field is absent or set with null, the endpoints configuration
	// remains unchanged.
	Endpoints LocationEndpoints `json:"endpoints"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The primary destination IPv4 address from the pair identified by the
	// dns_destination_ips_id. This field is read-only.
	IPV4Destination string `json:"ipv4_destination"`
	// The backup destination IPv4 address from the pair identified by the
	// dns_destination_ips_id. This field is read-only.
	IPV4DestinationBackup string `json:"ipv4_destination_backup"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this
	// location.
	Networks  []LocationNetwork `json:"networks"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      locationJSON      `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	ID                    apijson.Field
	ClientDefault         apijson.Field
	CreatedAt             apijson.Field
	DNSDestinationIPsID   apijson.Field
	DOHSubdomain          apijson.Field
	ECSSupport            apijson.Field
	Endpoints             apijson.Field
	IP                    apijson.Field
	IPV4Destination       apijson.Field
	IPV4DestinationBackup apijson.Field
	Name                  apijson.Field
	Networks              apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationJSON) RawJSON() string {
	return r.raw
}

// The destination endpoints configured for this location. When updating a
// location, if this field is absent or set with null, the endpoints configuration
// remains unchanged.
type LocationEndpoints struct {
	DOH  LocationEndpointsDOH  `json:"doh"`
	DOT  LocationEndpointsDOT  `json:"dot"`
	IPV4 LocationEndpointsIPV4 `json:"ipv4"`
	IPV6 LocationEndpointsIPV6 `json:"ipv6"`
	JSON locationEndpointsJSON `json:"-"`
}

// locationEndpointsJSON contains the JSON metadata for the struct
// [LocationEndpoints]
type locationEndpointsJSON struct {
	DOH         apijson.Field
	DOT         apijson.Field
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpoints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsJSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsDOH struct {
	// True if the endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks []LocationEndpointsDOHNetwork `json:"networks"`
	// True if the endpoint requires
	// [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user)
	// authentication.
	RequireToken bool                     `json:"require_token"`
	JSON         locationEndpointsDOHJSON `json:"-"`
}

// locationEndpointsDOHJSON contains the JSON metadata for the struct
// [LocationEndpointsDOH]
type locationEndpointsDOHJSON struct {
	Enabled      apijson.Field
	Networks     apijson.Field
	RequireToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LocationEndpointsDOH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsDOHJSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsDOHNetwork struct {
	// The IP address or IP CIDR.
	Network string                          `json:"network,required"`
	JSON    locationEndpointsDOHNetworkJSON `json:"-"`
}

// locationEndpointsDOHNetworkJSON contains the JSON metadata for the struct
// [LocationEndpointsDOHNetwork]
type locationEndpointsDOHNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpointsDOHNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsDOHNetworkJSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsDOT struct {
	// True if the endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks []LocationEndpointsDOTNetwork `json:"networks"`
	JSON     locationEndpointsDOTJSON      `json:"-"`
}

// locationEndpointsDOTJSON contains the JSON metadata for the struct
// [LocationEndpointsDOT]
type locationEndpointsDOTJSON struct {
	Enabled     apijson.Field
	Networks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpointsDOT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsDOTJSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsDOTNetwork struct {
	// The IP address or IP CIDR.
	Network string                          `json:"network,required"`
	JSON    locationEndpointsDOTNetworkJSON `json:"-"`
}

// locationEndpointsDOTNetworkJSON contains the JSON metadata for the struct
// [LocationEndpointsDOTNetwork]
type locationEndpointsDOTNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpointsDOTNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsDOTNetworkJSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsIPV4 struct {
	// True if the endpoint is enabled for this location.
	Enabled bool                      `json:"enabled"`
	JSON    locationEndpointsIPV4JSON `json:"-"`
}

// locationEndpointsIPV4JSON contains the JSON metadata for the struct
// [LocationEndpointsIPV4]
type locationEndpointsIPV4JSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpointsIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsIPV4JSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsIPV6 struct {
	// True if the endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// A list of allowed source IPv6 network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks []LocationEndpointsIPV6Network `json:"networks"`
	JSON     locationEndpointsIPV6JSON      `json:"-"`
}

// locationEndpointsIPV6JSON contains the JSON metadata for the struct
// [LocationEndpointsIPV6]
type locationEndpointsIPV6JSON struct {
	Enabled     apijson.Field
	Networks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpointsIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsIPV6JSON) RawJSON() string {
	return r.raw
}

type LocationEndpointsIPV6Network struct {
	// The IPv6 address or IPv6 CIDR.
	Network string                           `json:"network,required"`
	JSON    locationEndpointsIPV6NetworkJSON `json:"-"`
}

// locationEndpointsIPV6NetworkJSON contains the JSON metadata for the struct
// [LocationEndpointsIPV6Network]
type locationEndpointsIPV6NetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationEndpointsIPV6Network) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationEndpointsIPV6NetworkJSON) RawJSON() string {
	return r.raw
}

type LocationNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
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

// Union satisfied by [zero_trust.GatewayLocationDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayLocationDeleteResponseUnion interface {
	ImplementsZeroTrustGatewayLocationDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayLocationDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayLocationNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// The identifier of the pair of IPv4 addresses assigned to this location. When
	// creating a location, if this field is absent or set with null, the pair of
	// shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned.
	// When updating a location, if the field is absent or set with null, the
	// pre-assigned pair remains unchanged.
	DNSDestinationIPsID param.Field[string] `json:"dns_destination_ips_id"`
	// True if the location needs to resolve EDNS queries.
	ECSSupport param.Field[bool] `json:"ecs_support"`
	// The destination endpoints configured for this location. When updating a
	// location, if this field is absent or set with null, the endpoints configuration
	// remains unchanged.
	Endpoints param.Field[GatewayLocationNewParamsEndpoints] `json:"endpoints"`
	// A list of network ranges that requests from this location would originate from.
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this
	// location.
	Networks param.Field[[]GatewayLocationNewParamsNetwork] `json:"networks"`
}

func (r GatewayLocationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination endpoints configured for this location. When updating a
// location, if this field is absent or set with null, the endpoints configuration
// remains unchanged.
type GatewayLocationNewParamsEndpoints struct {
	DOH  param.Field[GatewayLocationNewParamsEndpointsDOH]  `json:"doh"`
	DOT  param.Field[GatewayLocationNewParamsEndpointsDOT]  `json:"dot"`
	IPV4 param.Field[GatewayLocationNewParamsEndpointsIPV4] `json:"ipv4"`
	IPV6 param.Field[GatewayLocationNewParamsEndpointsIPV6] `json:"ipv6"`
}

func (r GatewayLocationNewParamsEndpoints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsDOH struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]GatewayLocationNewParamsEndpointsDOHNetwork] `json:"networks"`
	// True if the endpoint requires
	// [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user)
	// authentication.
	RequireToken param.Field[bool] `json:"require_token"`
}

func (r GatewayLocationNewParamsEndpointsDOH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsDOHNetwork struct {
	// The IP address or IP CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationNewParamsEndpointsDOHNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsDOT struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]GatewayLocationNewParamsEndpointsDOTNetwork] `json:"networks"`
}

func (r GatewayLocationNewParamsEndpointsDOT) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsDOTNetwork struct {
	// The IP address or IP CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationNewParamsEndpointsDOTNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsIPV4 struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayLocationNewParamsEndpointsIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsIPV6 struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IPv6 network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]GatewayLocationNewParamsEndpointsIPV6Network] `json:"networks"`
}

func (r GatewayLocationNewParamsEndpointsIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsEndpointsIPV6Network struct {
	// The IPv6 address or IPv6 CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationNewParamsEndpointsIPV6Network) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// The identifier of the pair of IPv4 addresses assigned to this location. When
	// creating a location, if this field is absent or set with null, the pair of
	// shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned.
	// When updating a location, if the field is absent or set with null, the
	// pre-assigned pair remains unchanged.
	DNSDestinationIPsID param.Field[string] `json:"dns_destination_ips_id"`
	// True if the location needs to resolve EDNS queries.
	ECSSupport param.Field[bool] `json:"ecs_support"`
	// The destination endpoints configured for this location. When updating a
	// location, if this field is absent or set with null, the endpoints configuration
	// remains unchanged.
	Endpoints param.Field[GatewayLocationUpdateParamsEndpoints] `json:"endpoints"`
	// A list of network ranges that requests from this location would originate from.
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this
	// location.
	Networks param.Field[[]GatewayLocationUpdateParamsNetwork] `json:"networks"`
}

func (r GatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination endpoints configured for this location. When updating a
// location, if this field is absent or set with null, the endpoints configuration
// remains unchanged.
type GatewayLocationUpdateParamsEndpoints struct {
	DOH  param.Field[GatewayLocationUpdateParamsEndpointsDOH]  `json:"doh"`
	DOT  param.Field[GatewayLocationUpdateParamsEndpointsDOT]  `json:"dot"`
	IPV4 param.Field[GatewayLocationUpdateParamsEndpointsIPV4] `json:"ipv4"`
	IPV6 param.Field[GatewayLocationUpdateParamsEndpointsIPV6] `json:"ipv6"`
}

func (r GatewayLocationUpdateParamsEndpoints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsDOH struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]GatewayLocationUpdateParamsEndpointsDOHNetwork] `json:"networks"`
	// True if the endpoint requires
	// [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user)
	// authentication.
	RequireToken param.Field[bool] `json:"require_token"`
}

func (r GatewayLocationUpdateParamsEndpointsDOH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsDOHNetwork struct {
	// The IP address or IP CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsEndpointsDOHNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsDOT struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]GatewayLocationUpdateParamsEndpointsDOTNetwork] `json:"networks"`
}

func (r GatewayLocationUpdateParamsEndpointsDOT) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsDOTNetwork struct {
	// The IP address or IP CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsEndpointsDOTNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsIPV4 struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayLocationUpdateParamsEndpointsIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsIPV6 struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IPv6 network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]GatewayLocationUpdateParamsEndpointsIPV6Network] `json:"networks"`
}

func (r GatewayLocationUpdateParamsEndpointsIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsEndpointsIPV6Network struct {
	// The IPv6 address or IPv6 CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsEndpointsIPV6Network) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
	// Whether the API call was successful
	Success GatewayLocationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayLocationDeleteResponseUnion           `json:"result"`
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

// Whether the API call was successful
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
	// Whether the API call was successful
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

// Whether the API call was successful
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
