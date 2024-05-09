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

// GatewayProxyEndpointService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayProxyEndpointService] method instead.
type GatewayProxyEndpointService struct {
	Options []option.RequestOption
}

// NewGatewayProxyEndpointService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayProxyEndpointService(opts ...option.RequestOption) (r *GatewayProxyEndpointService) {
	r = &GatewayProxyEndpointService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) New(ctx context.Context, params GatewayProxyEndpointNewParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) List(ctx context.Context, query GatewayProxyEndpointListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProxyEndpoint], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", query.AccountID)
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

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) ListAutoPaging(ctx context.Context, query GatewayProxyEndpointListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProxyEndpoint] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Delete(ctx context.Context, proxyEndpointID string, body GatewayProxyEndpointDeleteParams, opts ...option.RequestOption) (res *GatewayProxyEndpointDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", body.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Edit(ctx context.Context, proxyEndpointID string, params GatewayProxyEndpointEditParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", params.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) Get(ctx context.Context, proxyEndpointID string, query GatewayProxyEndpointGetParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", query.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayIPs = string

type GatewayIPsParam = string

type ProxyEndpoint struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []GatewayIPs `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string            `json:"subdomain"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointJSON `json:"-"`
}

// proxyEndpointJSON contains the JSON metadata for the struct [ProxyEndpoint]
type proxyEndpointJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxyEndpointJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.GatewayProxyEndpointDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayProxyEndpointDeleteResponseUnion interface {
	ImplementsZeroTrustGatewayProxyEndpointDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayProxyEndpointDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayProxyEndpointNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]GatewayIPsParam] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
}

func (r GatewayProxyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ProxyEndpoint                                  `json:"result"`
	JSON    gatewayProxyEndpointNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointNewResponseEnvelope]
type gatewayProxyEndpointNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointNewResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointNewResponseEnvelopeSuccessTrue GatewayProxyEndpointNewResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayProxyEndpointDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayProxyEndpointDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayProxyEndpointDeleteResponseUnion           `json:"result"`
	JSON    gatewayProxyEndpointDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayProxyEndpointDeleteResponseEnvelope]
type gatewayProxyEndpointDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointDeleteResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue GatewayProxyEndpointDeleteResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]GatewayIPsParam] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
}

func (r GatewayProxyEndpointEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ProxyEndpoint                                   `json:"result"`
	JSON    gatewayProxyEndpointEditResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointEditResponseEnvelope]
type gatewayProxyEndpointEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointEditResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointEditResponseEnvelopeSuccessTrue GatewayProxyEndpointEditResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayProxyEndpointGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ProxyEndpoint                                  `json:"result"`
	JSON    gatewayProxyEndpointGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointGetResponseEnvelope]
type gatewayProxyEndpointGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointGetResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointGetResponseEnvelopeSuccessTrue GatewayProxyEndpointGetResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
