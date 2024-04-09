// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// GRETunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewGRETunnelService] method instead.
type GRETunnelService struct {
	Options []option.RequestOption
}

// NewGRETunnelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGRETunnelService(opts ...option.RequestOption) (r *GRETunnelService) {
	r = &GRETunnelService{}
	r.Options = opts
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *GRETunnelService) New(ctx context.Context, params GRETunnelNewParams, opts ...option.RequestOption) (res *GRETunnelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specific GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *GRETunnelService) Update(ctx context.Context, tunnelIdentifier string, params GRETunnelUpdateParams, opts ...option.RequestOption) (res *GRETunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", params.AccountID, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *GRETunnelService) List(ctx context.Context, query GRETunnelListParams, opts ...option.RequestOption) (res *GRETunnelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as
// an optional query parameter to only run validation without persisting changes.
func (r *GRETunnelService) Delete(ctx context.Context, tunnelIdentifier string, params GRETunnelDeleteParams, opts ...option.RequestOption) (res *GRETunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", params.AccountID, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *GRETunnelService) Get(ctx context.Context, tunnelIdentifier string, query GRETunnelGetParams, opts ...option.RequestOption) (res *GRETunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GRETunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", query.AccountID, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GRETunnelNewResponse struct {
	GRETunnels []GRETunnelNewResponseGRETunnel `json:"gre_tunnels"`
	JSON       greTunnelNewResponseJSON        `json:"-"`
}

// greTunnelNewResponseJSON contains the JSON metadata for the struct
// [GRETunnelNewResponse]
type greTunnelNewResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewResponseGRETunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGREEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGREEndpoint string `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the GRE tunnel.
	Description string      `json:"description"`
	HealthCheck HealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                             `json:"ttl"`
	JSON greTunnelNewResponseGRETunnelJSON `json:"-"`
}

// greTunnelNewResponseGRETunnelJSON contains the JSON metadata for the struct
// [GRETunnelNewResponseGRETunnel]
type greTunnelNewResponseGRETunnelJSON struct {
	CloudflareGREEndpoint apijson.Field
	CustomerGREEndpoint   apijson.Field
	InterfaceAddress      apijson.Field
	Name                  apijson.Field
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	HealthCheck           apijson.Field
	ModifiedOn            apijson.Field
	Mtu                   apijson.Field
	TTL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GRETunnelNewResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelUpdateResponse struct {
	Modified          bool                        `json:"modified"`
	ModifiedGRETunnel interface{}                 `json:"modified_gre_tunnel"`
	JSON              greTunnelUpdateResponseJSON `json:"-"`
}

// greTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [GRETunnelUpdateResponse]
type greTunnelUpdateResponseJSON struct {
	Modified          apijson.Field
	ModifiedGRETunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GRETunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponse struct {
	GRETunnels []GRETunnelListResponseGRETunnel `json:"gre_tunnels"`
	JSON       greTunnelListResponseJSON        `json:"-"`
}

// greTunnelListResponseJSON contains the JSON metadata for the struct
// [GRETunnelListResponse]
type greTunnelListResponseJSON struct {
	GRETunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelListResponseGRETunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGREEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGREEndpoint string `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the GRE tunnel.
	Description string      `json:"description"`
	HealthCheck HealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                              `json:"ttl"`
	JSON greTunnelListResponseGRETunnelJSON `json:"-"`
}

// greTunnelListResponseGRETunnelJSON contains the JSON metadata for the struct
// [GRETunnelListResponseGRETunnel]
type greTunnelListResponseGRETunnelJSON struct {
	CloudflareGREEndpoint apijson.Field
	CustomerGREEndpoint   apijson.Field
	InterfaceAddress      apijson.Field
	Name                  apijson.Field
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	HealthCheck           apijson.Field
	ModifiedOn            apijson.Field
	Mtu                   apijson.Field
	TTL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GRETunnelListResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelDeleteResponse struct {
	Deleted          bool                        `json:"deleted"`
	DeletedGRETunnel interface{}                 `json:"deleted_gre_tunnel"`
	JSON             greTunnelDeleteResponseJSON `json:"-"`
}

// greTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [GRETunnelDeleteResponse]
type greTunnelDeleteResponseJSON struct {
	Deleted          apijson.Field
	DeletedGRETunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GRETunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelGetResponse struct {
	GRETunnel interface{}              `json:"gre_tunnel"`
	JSON      greTunnelGetResponseJSON `json:"-"`
}

// greTunnelGetResponseJSON contains the JSON metadata for the struct
// [GRETunnelGetResponse]
type greTunnelGetResponseJSON struct {
	GRETunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type GRETunnelNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r GRETunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GRETunnelNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   GRETunnelNewResponse  `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// greTunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelNewResponseEnvelope]
type greTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelNewResponseEnvelopeSuccess bool

const (
	GRETunnelNewResponseEnvelopeSuccessTrue GRETunnelNewResponseEnvelopeSuccess = true
)

func (r GRETunnelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GRETunnelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GRETunnelUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGREEndpoint param.Field[string] `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGREEndpoint param.Field[string] `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name param.Field[string] `json:"name,required"`
	// An optional description of the GRE tunnel.
	Description param.Field[string]           `json:"description"`
	HealthCheck param.Field[HealthCheckParam] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL param.Field[int64] `json:"ttl"`
}

func (r GRETunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GRETunnelUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   GRETunnelUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// greTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelUpdateResponseEnvelope]
type greTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelUpdateResponseEnvelopeSuccess bool

const (
	GRETunnelUpdateResponseEnvelopeSuccessTrue GRETunnelUpdateResponseEnvelopeSuccess = true
)

func (r GRETunnelUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GRETunnelUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GRETunnelListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type GRETunnelListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   GRETunnelListResponse `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelListResponseEnvelopeJSON    `json:"-"`
}

// greTunnelListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelListResponseEnvelope]
type greTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelListResponseEnvelopeSuccess bool

const (
	GRETunnelListResponseEnvelopeSuccessTrue GRETunnelListResponseEnvelopeSuccess = true
)

func (r GRETunnelListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GRETunnelListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GRETunnelDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r GRETunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GRETunnelDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   GRETunnelDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// greTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelDeleteResponseEnvelope]
type greTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelDeleteResponseEnvelopeSuccess bool

const (
	GRETunnelDeleteResponseEnvelopeSuccessTrue GRETunnelDeleteResponseEnvelopeSuccess = true
)

func (r GRETunnelDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GRETunnelDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GRETunnelGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type GRETunnelGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   GRETunnelGetResponse  `json:"result,required"`
	// Whether the API call was successful
	Success GRETunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    greTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// greTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GRETunnelGetResponseEnvelope]
type greTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GRETunnelGetResponseEnvelopeSuccess bool

const (
	GRETunnelGetResponseEnvelopeSuccessTrue GRETunnelGetResponseEnvelopeSuccess = true
)

func (r GRETunnelGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GRETunnelGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
