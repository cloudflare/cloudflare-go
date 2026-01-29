// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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

// Create a new Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) New(ctx context.Context, params GatewayProxyEndpointNewParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	var env GatewayProxyEndpointNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) List(ctx context.Context, query GatewayProxyEndpointListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProxyEndpoint], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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

// List all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) ListAutoPaging(ctx context.Context, query GatewayProxyEndpointListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProxyEndpoint] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Delete(ctx context.Context, proxyEndpointID string, body GatewayProxyEndpointDeleteParams, opts ...option.RequestOption) (res *GatewayProxyEndpointDeleteResponse, err error) {
	var env GatewayProxyEndpointDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", body.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Edit(ctx context.Context, proxyEndpointID string, params GatewayProxyEndpointEditParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	var env GatewayProxyEndpointEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", params.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Get(ctx context.Context, proxyEndpointID string, query GatewayProxyEndpointGetParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	var env GatewayProxyEndpointGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
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
	// Specify the name of the proxy endpoint.
	Name      string    `json:"name,required"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of [[]GatewayIPs].
	IPs interface{} `json:"ips"`
	// The proxy endpoint kind
	Kind ProxyEndpointKind `json:"kind"`
	// Specify the subdomain to use as the destination in the proxy client.
	Subdomain string            `json:"subdomain"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointJSON `json:"-"`
	union     ProxyEndpointUnion
}

// proxyEndpointJSON contains the JSON metadata for the struct [ProxyEndpoint]
type proxyEndpointJSON struct {
	Name        apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Kind        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r proxyEndpointJSON) RawJSON() string {
	return r.raw
}

func (r *ProxyEndpoint) UnmarshalJSON(data []byte) (err error) {
	*r = ProxyEndpoint{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProxyEndpointUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ProxyEndpointZeroTrustGatewayProxyEndpointIP],
// [ProxyEndpointZeroTrustGatewayProxyEndpointIdentity].
func (r ProxyEndpoint) AsUnion() ProxyEndpointUnion {
	return r.union
}

// Union satisfied by [ProxyEndpointZeroTrustGatewayProxyEndpointIP] or
// [ProxyEndpointZeroTrustGatewayProxyEndpointIdentity].
type ProxyEndpointUnion interface {
	implementsProxyEndpoint()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProxyEndpointUnion)(nil)).Elem(),
		"kind",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProxyEndpointZeroTrustGatewayProxyEndpointIP{}),
			DiscriminatorValue: "ip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProxyEndpointZeroTrustGatewayProxyEndpointIdentity{}),
			DiscriminatorValue: "identity",
		},
	)
}

type ProxyEndpointZeroTrustGatewayProxyEndpointIP struct {
	// Specify the list of CIDRs to restrict ingress connections.
	IPs []GatewayIPs `json:"ips,required"`
	// Specify the name of the proxy endpoint.
	Name      string    `json:"name,required"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The proxy endpoint kind
	Kind ProxyEndpointZeroTrustGatewayProxyEndpointIPKind `json:"kind"`
	// Specify the subdomain to use as the destination in the proxy client.
	Subdomain string                                           `json:"subdomain"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointZeroTrustGatewayProxyEndpointIPJSON `json:"-"`
}

// proxyEndpointZeroTrustGatewayProxyEndpointIPJSON contains the JSON metadata for
// the struct [ProxyEndpointZeroTrustGatewayProxyEndpointIP]
type proxyEndpointZeroTrustGatewayProxyEndpointIPJSON struct {
	IPs         apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	Kind        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointZeroTrustGatewayProxyEndpointIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxyEndpointZeroTrustGatewayProxyEndpointIPJSON) RawJSON() string {
	return r.raw
}

func (r ProxyEndpointZeroTrustGatewayProxyEndpointIP) implementsProxyEndpoint() {}

// The proxy endpoint kind
type ProxyEndpointZeroTrustGatewayProxyEndpointIPKind string

const (
	ProxyEndpointZeroTrustGatewayProxyEndpointIPKindIP ProxyEndpointZeroTrustGatewayProxyEndpointIPKind = "ip"
)

func (r ProxyEndpointZeroTrustGatewayProxyEndpointIPKind) IsKnown() bool {
	switch r {
	case ProxyEndpointZeroTrustGatewayProxyEndpointIPKindIP:
		return true
	}
	return false
}

type ProxyEndpointZeroTrustGatewayProxyEndpointIdentity struct {
	// The proxy endpoint kind
	Kind ProxyEndpointZeroTrustGatewayProxyEndpointIdentityKind `json:"kind,required"`
	// Specify the name of the proxy endpoint.
	Name      string    `json:"name,required"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specify the subdomain to use as the destination in the proxy client.
	Subdomain string                                                 `json:"subdomain"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointZeroTrustGatewayProxyEndpointIdentityJSON `json:"-"`
}

// proxyEndpointZeroTrustGatewayProxyEndpointIdentityJSON contains the JSON
// metadata for the struct [ProxyEndpointZeroTrustGatewayProxyEndpointIdentity]
type proxyEndpointZeroTrustGatewayProxyEndpointIdentityJSON struct {
	Kind        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointZeroTrustGatewayProxyEndpointIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxyEndpointZeroTrustGatewayProxyEndpointIdentityJSON) RawJSON() string {
	return r.raw
}

func (r ProxyEndpointZeroTrustGatewayProxyEndpointIdentity) implementsProxyEndpoint() {}

// The proxy endpoint kind
type ProxyEndpointZeroTrustGatewayProxyEndpointIdentityKind string

const (
	ProxyEndpointZeroTrustGatewayProxyEndpointIdentityKindIdentity ProxyEndpointZeroTrustGatewayProxyEndpointIdentityKind = "identity"
)

func (r ProxyEndpointZeroTrustGatewayProxyEndpointIdentityKind) IsKnown() bool {
	switch r {
	case ProxyEndpointZeroTrustGatewayProxyEndpointIdentityKindIdentity:
		return true
	}
	return false
}

// The proxy endpoint kind
type ProxyEndpointKind string

const (
	ProxyEndpointKindIP       ProxyEndpointKind = "ip"
	ProxyEndpointKindIdentity ProxyEndpointKind = "identity"
)

func (r ProxyEndpointKind) IsKnown() bool {
	switch r {
	case ProxyEndpointKindIP, ProxyEndpointKindIdentity:
		return true
	}
	return false
}

type GatewayProxyEndpointDeleteResponse = interface{}

type GatewayProxyEndpointNewParams struct {
	AccountID param.Field[string]                    `path:"account_id,required"`
	Body      GatewayProxyEndpointNewParamsBodyUnion `json:"body,required"`
}

func (r GatewayProxyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GatewayProxyEndpointNewParamsBody struct {
	// Specify the name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
	// The proxy endpoint kind
	Kind param.Field[GatewayProxyEndpointNewParamsBodyKind] `json:"kind"`
}

func (r GatewayProxyEndpointNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GatewayProxyEndpointNewParamsBody) implementsGatewayProxyEndpointNewParamsBodyUnion() {}

// Satisfied by
// [zero_trust.GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreate],
// [zero_trust.GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreate],
// [GatewayProxyEndpointNewParamsBody].
type GatewayProxyEndpointNewParamsBodyUnion interface {
	implementsGatewayProxyEndpointNewParamsBodyUnion()
}

type GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreate struct {
	// Specify the name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
	// The proxy endpoint kind
	Kind param.Field[GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreateKind] `json:"kind"`
}

func (r GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreate) implementsGatewayProxyEndpointNewParamsBodyUnion() {
}

// The proxy endpoint kind
type GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreateKind string

const (
	GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreateKindIP GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreateKind = "ip"
)

func (r GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreateKind) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIPCreateKindIP:
		return true
	}
	return false
}

type GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreate struct {
	// The proxy endpoint kind
	Kind param.Field[GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreateKind] `json:"kind,required"`
	// Specify the name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
}

func (r GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreate) implementsGatewayProxyEndpointNewParamsBodyUnion() {
}

// The proxy endpoint kind
type GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreateKind string

const (
	GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreateKindIdentity GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreateKind = "identity"
)

func (r GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreateKind) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointNewParamsBodyZeroTrustGatewayProxyEndpointIdentityCreateKindIdentity:
		return true
	}
	return false
}

// The proxy endpoint kind
type GatewayProxyEndpointNewParamsBodyKind string

const (
	GatewayProxyEndpointNewParamsBodyKindIP       GatewayProxyEndpointNewParamsBodyKind = "ip"
	GatewayProxyEndpointNewParamsBodyKindIdentity GatewayProxyEndpointNewParamsBodyKind = "identity"
)

func (r GatewayProxyEndpointNewParamsBodyKind) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointNewParamsBodyKindIP, GatewayProxyEndpointNewParamsBodyKindIdentity:
		return true
	}
	return false
}

type GatewayProxyEndpointNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
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

// Indicate whether the API call was successful.
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
	// Indicate whether the API call was successful.
	Success GatewayProxyEndpointDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayProxyEndpointDeleteResponse                `json:"result"`
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

// Indicate whether the API call was successful.
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
	// Specify the list of CIDRs to restrict ingress connections.
	IPs param.Field[[]GatewayIPsParam] `json:"ips"`
	// Specify the name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
}

func (r GatewayProxyEndpointEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
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

// Indicate whether the API call was successful.
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
	// Indicate whether the API call was successful.
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

// Indicate whether the API call was successful.
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
