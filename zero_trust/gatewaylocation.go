// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
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
	opts = append(r.Options[:], opts...)
	var env GatewayLocationNewResponseEnvelope
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
	opts = append(r.Options[:], opts...)
	var env GatewayLocationUpdateResponseEnvelope
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
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	opts = append(r.Options[:], opts...)
	var env GatewayLocationDeleteResponseEnvelope
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
	opts = append(r.Options[:], opts...)
	var env GatewayLocationGetResponseEnvelope
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
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []LocationNetwork `json:"networks"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      locationJSON      `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationJSON) RawJSON() string {
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

type LocationNetworkParam struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r LocationNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]LocationNetworkParam] `json:"networks"`
}

func (r GatewayLocationNewParams) MarshalJSON() (data []byte, err error) {
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
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]LocationNetworkParam] `json:"networks"`
}

func (r GatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
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
