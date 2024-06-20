// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// GRETunnelService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGRETunnelService] method instead.
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
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
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
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", query.AccountID)
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

type GRETunnelNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
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
