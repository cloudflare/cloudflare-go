// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/tidwall/gjson"
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
	var env GRETunnelNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// Updates a specific GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *GRETunnelService) Update(ctx context.Context, greTunnelID string, params GRETunnelUpdateParams, opts ...option.RequestOption) (res *GRETunnelUpdateResponse, err error) {
	var env GRETunnelUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if greTunnelID == "" {
		err = errors.New("missing required gre_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", params.AccountID, greTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists GRE tunnels associated with an account.
func (r *GRETunnelService) List(ctx context.Context, params GRETunnelListParams, opts ...option.RequestOption) (res *GRETunnelListResponse, err error) {
	var env GRETunnelListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as
// an optional query parameter to only run validation without persisting changes.
func (r *GRETunnelService) Delete(ctx context.Context, greTunnelID string, params GRETunnelDeleteParams, opts ...option.RequestOption) (res *GRETunnelDeleteResponse, err error) {
	var env GRETunnelDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if greTunnelID == "" {
		err = errors.New("missing required gre_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", params.AccountID, greTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *GRETunnelService) Get(ctx context.Context, greTunnelID string, params GRETunnelGetParams, opts ...option.RequestOption) (res *GRETunnelGetResponse, err error) {
	var env GRETunnelGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if greTunnelID == "" {
		err = errors.New("missing required gre_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", params.AccountID, greTunnelID)
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
	Description string                                    `json:"description"`
	HealthCheck GRETunnelNewResponseGRETunnelsHealthCheck `json:"health_check"`
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

type GRETunnelNewResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction GRETunnelNewResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target GRETunnelNewResponseGRETunnelsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                               `json:"type"`
	JSON greTunnelNewResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// greTunnelNewResponseGRETunnelsHealthCheckJSON contains the JSON metadata for the
// struct [GRETunnelNewResponseGRETunnelsHealthCheck]
type greTunnelNewResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseGRETunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type GRETunnelNewResponseGRETunnelsHealthCheckDirection string

const (
	GRETunnelNewResponseGRETunnelsHealthCheckDirectionUnidirectional GRETunnelNewResponseGRETunnelsHealthCheckDirection = "unidirectional"
	GRETunnelNewResponseGRETunnelsHealthCheckDirectionBidirectional  GRETunnelNewResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

func (r GRETunnelNewResponseGRETunnelsHealthCheckDirection) IsKnown() bool {
	switch r {
	case GRETunnelNewResponseGRETunnelsHealthCheckDirectionUnidirectional, GRETunnelNewResponseGRETunnelsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.GRETunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type GRETunnelNewResponseGRETunnelsHealthCheckTargetUnion interface {
	ImplementsMagicTransitGRETunnelNewResponseGRETunnelsHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GRETunnelNewResponseGRETunnelsHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GRETunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type GRETunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                    `json:"saved"`
	JSON  greTunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// greTunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [GRETunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget]
type greTunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r GRETunnelNewResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitGRETunnelNewResponseGRETunnelsHealthCheckTargetUnion() {
}

type GRETunnelUpdateResponse struct {
	Modified          bool                                     `json:"modified"`
	ModifiedGRETunnel GRETunnelUpdateResponseModifiedGRETunnel `json:"modified_gre_tunnel"`
	JSON              greTunnelUpdateResponseJSON              `json:"-"`
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

type GRETunnelUpdateResponseModifiedGRETunnel struct {
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
	Description string                                              `json:"description"`
	HealthCheck GRETunnelUpdateResponseModifiedGRETunnelHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                        `json:"ttl"`
	JSON greTunnelUpdateResponseModifiedGRETunnelJSON `json:"-"`
}

// greTunnelUpdateResponseModifiedGRETunnelJSON contains the JSON metadata for the
// struct [GRETunnelUpdateResponseModifiedGRETunnel]
type greTunnelUpdateResponseModifiedGRETunnelJSON struct {
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

func (r *GRETunnelUpdateResponseModifiedGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseModifiedGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelUpdateResponseModifiedGRETunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                         `json:"type"`
	JSON greTunnelUpdateResponseModifiedGRETunnelHealthCheckJSON `json:"-"`
}

// greTunnelUpdateResponseModifiedGRETunnelHealthCheckJSON contains the JSON
// metadata for the struct [GRETunnelUpdateResponseModifiedGRETunnelHealthCheck]
type greTunnelUpdateResponseModifiedGRETunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelUpdateResponseModifiedGRETunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseModifiedGRETunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirection string

const (
	GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirectionUnidirectional GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirection = "unidirectional"
	GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirectionBidirectional  GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirection = "bidirectional"
)

func (r GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirectionUnidirectional, GRETunnelUpdateResponseModifiedGRETunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion interface {
	ImplementsMagicTransitGRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                              `json:"saved"`
	JSON  greTunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// greTunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTarget]
type greTunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r GRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitGRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion() {
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
	Description string                                     `json:"description"`
	HealthCheck GRETunnelListResponseGRETunnelsHealthCheck `json:"health_check"`
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

type GRETunnelListResponseGRETunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction GRETunnelListResponseGRETunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target GRETunnelListResponseGRETunnelsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                `json:"type"`
	JSON greTunnelListResponseGRETunnelsHealthCheckJSON `json:"-"`
}

// greTunnelListResponseGRETunnelsHealthCheckJSON contains the JSON metadata for
// the struct [GRETunnelListResponseGRETunnelsHealthCheck]
type greTunnelListResponseGRETunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseGRETunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseGRETunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type GRETunnelListResponseGRETunnelsHealthCheckDirection string

const (
	GRETunnelListResponseGRETunnelsHealthCheckDirectionUnidirectional GRETunnelListResponseGRETunnelsHealthCheckDirection = "unidirectional"
	GRETunnelListResponseGRETunnelsHealthCheckDirectionBidirectional  GRETunnelListResponseGRETunnelsHealthCheckDirection = "bidirectional"
)

func (r GRETunnelListResponseGRETunnelsHealthCheckDirection) IsKnown() bool {
	switch r {
	case GRETunnelListResponseGRETunnelsHealthCheckDirectionUnidirectional, GRETunnelListResponseGRETunnelsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.GRETunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type GRETunnelListResponseGRETunnelsHealthCheckTargetUnion interface {
	ImplementsMagicTransitGRETunnelListResponseGRETunnelsHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GRETunnelListResponseGRETunnelsHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GRETunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type GRETunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                     `json:"saved"`
	JSON  greTunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// greTunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [GRETunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget]
type greTunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r GRETunnelListResponseGRETunnelsHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitGRETunnelListResponseGRETunnelsHealthCheckTargetUnion() {
}

type GRETunnelDeleteResponse struct {
	Deleted          bool                                    `json:"deleted"`
	DeletedGRETunnel GRETunnelDeleteResponseDeletedGRETunnel `json:"deleted_gre_tunnel"`
	JSON             greTunnelDeleteResponseJSON             `json:"-"`
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

type GRETunnelDeleteResponseDeletedGRETunnel struct {
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
	Description string                                             `json:"description"`
	HealthCheck GRETunnelDeleteResponseDeletedGRETunnelHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                                       `json:"ttl"`
	JSON greTunnelDeleteResponseDeletedGRETunnelJSON `json:"-"`
}

// greTunnelDeleteResponseDeletedGRETunnelJSON contains the JSON metadata for the
// struct [GRETunnelDeleteResponseDeletedGRETunnel]
type greTunnelDeleteResponseDeletedGRETunnelJSON struct {
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

func (r *GRETunnelDeleteResponseDeletedGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseDeletedGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelDeleteResponseDeletedGRETunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                        `json:"type"`
	JSON greTunnelDeleteResponseDeletedGRETunnelHealthCheckJSON `json:"-"`
}

// greTunnelDeleteResponseDeletedGRETunnelHealthCheckJSON contains the JSON
// metadata for the struct [GRETunnelDeleteResponseDeletedGRETunnelHealthCheck]
type greTunnelDeleteResponseDeletedGRETunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelDeleteResponseDeletedGRETunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseDeletedGRETunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirection string

const (
	GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirectionUnidirectional GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirection = "unidirectional"
	GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirectionBidirectional  GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirection = "bidirectional"
)

func (r GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirectionUnidirectional, GRETunnelDeleteResponseDeletedGRETunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion interface {
	ImplementsMagicTransitGRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                             `json:"saved"`
	JSON  greTunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// greTunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTarget]
type greTunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r GRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitGRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion() {
}

type GRETunnelGetResponse struct {
	GRETunnel GRETunnelGetResponseGRETunnel `json:"gre_tunnel"`
	JSON      greTunnelGetResponseJSON      `json:"-"`
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

type GRETunnelGetResponseGRETunnel struct {
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
	Description string                                   `json:"description"`
	HealthCheck GRETunnelGetResponseGRETunnelHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL  int64                             `json:"ttl"`
	JSON greTunnelGetResponseGRETunnelJSON `json:"-"`
}

// greTunnelGetResponseGRETunnelJSON contains the JSON metadata for the struct
// [GRETunnelGetResponseGRETunnel]
type greTunnelGetResponseGRETunnelJSON struct {
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

func (r *GRETunnelGetResponseGRETunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseGRETunnelJSON) RawJSON() string {
	return r.raw
}

type GRETunnelGetResponseGRETunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction GRETunnelGetResponseGRETunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target GRETunnelGetResponseGRETunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                              `json:"type"`
	JSON greTunnelGetResponseGRETunnelHealthCheckJSON `json:"-"`
}

// greTunnelGetResponseGRETunnelHealthCheckJSON contains the JSON metadata for the
// struct [GRETunnelGetResponseGRETunnelHealthCheck]
type greTunnelGetResponseGRETunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponseGRETunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseGRETunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type GRETunnelGetResponseGRETunnelHealthCheckDirection string

const (
	GRETunnelGetResponseGRETunnelHealthCheckDirectionUnidirectional GRETunnelGetResponseGRETunnelHealthCheckDirection = "unidirectional"
	GRETunnelGetResponseGRETunnelHealthCheckDirectionBidirectional  GRETunnelGetResponseGRETunnelHealthCheckDirection = "bidirectional"
)

func (r GRETunnelGetResponseGRETunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case GRETunnelGetResponseGRETunnelHealthCheckDirectionUnidirectional, GRETunnelGetResponseGRETunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.GRETunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type GRETunnelGetResponseGRETunnelHealthCheckTargetUnion interface {
	ImplementsMagicTransitGRETunnelGetResponseGRETunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GRETunnelGetResponseGRETunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GRETunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type GRETunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                   `json:"saved"`
	JSON  greTunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// greTunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [GRETunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTarget]
type greTunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GRETunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r greTunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r GRETunnelGetResponseGRETunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitGRETunnelGetResponseGRETunnelHealthCheckTargetUnion() {
}

type GRETunnelNewParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	Body              interface{}         `json:"body,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
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
	Description param.Field[string]                           `json:"description"`
	HealthCheck param.Field[GRETunnelUpdateParamsHealthCheck] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL               param.Field[int64] `json:"ttl"`
	XMagicNewHcTarget param.Field[bool]  `header:"x-magic-new-hc-target"`
}

func (r GRETunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GRETunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction param.Field[GRETunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[GRETunnelUpdateParamsHealthCheckTargetUnion] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[HealthCheckType] `json:"type"`
}

func (r GRETunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type GRETunnelUpdateParamsHealthCheckDirection string

const (
	GRETunnelUpdateParamsHealthCheckDirectionUnidirectional GRETunnelUpdateParamsHealthCheckDirection = "unidirectional"
	GRETunnelUpdateParamsHealthCheckDirectionBidirectional  GRETunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

func (r GRETunnelUpdateParamsHealthCheckDirection) IsKnown() bool {
	switch r {
	case GRETunnelUpdateParamsHealthCheckDirectionUnidirectional, GRETunnelUpdateParamsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by
// [magic_transit.GRETunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget],
// [shared.UnionString].
type GRETunnelUpdateParamsHealthCheckTargetUnion interface {
	ImplementsMagicTransitGRETunnelUpdateParamsHealthCheckTargetUnion()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type GRETunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r GRETunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GRETunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitGRETunnelUpdateParamsHealthCheckTargetUnion() {
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
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
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
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
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
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
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
