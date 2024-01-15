// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCfdTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCfdTunnelService] method
// instead.
type AccountCfdTunnelService struct {
	Options        []option.RequestOption
	Configurations *AccountCfdTunnelConfigurationService
	Connections    *AccountCfdTunnelConnectionService
	Connectors     *AccountCfdTunnelConnectorService
	Management     *AccountCfdTunnelManagementService
	Tokens         *AccountCfdTunnelTokenService
}

// NewAccountCfdTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelService(opts ...option.RequestOption) (r *AccountCfdTunnelService) {
	r = &AccountCfdTunnelService{}
	r.Options = opts
	r.Configurations = NewAccountCfdTunnelConfigurationService(opts...)
	r.Connections = NewAccountCfdTunnelConnectionService(opts...)
	r.Connectors = NewAccountCfdTunnelConnectorService(opts...)
	r.Management = NewAccountCfdTunnelManagementService(opts...)
	r.Tokens = NewAccountCfdTunnelTokenService(opts...)
	return
}

// Fetches a single Cloudflare Tunnel.
func (r *AccountCfdTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *AccountCfdTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *AccountCfdTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelUpdateParams, opts ...option.RequestOption) (res *AccountCfdTunnelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a Cloudflare Tunnel from an account.
func (r *AccountCfdTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelDeleteParams, opts ...option.RequestOption) (res *AccountCfdTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Creates a new Cloudflare Tunnel in an account.
func (r *AccountCfdTunnelService) CloudflareTunnelNewACloudflareTunnel(ctx context.Context, accountIdentifier string, body AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParams, opts ...option.RequestOption) (res *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters Cloudflare Tunnels in an account.
func (r *AccountCfdTunnelService) CloudflareTunnelListCloudflareTunnels(ctx context.Context, accountIdentifier string, query AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsParams, opts ...option.RequestOption) (res *shared.Page[AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountIdentifier)
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

type AccountCfdTunnelGetResponse struct {
	Errors   []AccountCfdTunnelGetResponseError   `json:"errors"`
	Messages []AccountCfdTunnelGetResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountCfdTunnelGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelGetResponseSuccess `json:"success"`
	JSON    accountCfdTunnelGetResponseJSON    `json:"-"`
}

// accountCfdTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountCfdTunnelGetResponse]
type accountCfdTunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountCfdTunnelGetResponseErrorJSON `json:"-"`
}

// accountCfdTunnelGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountCfdTunnelGetResponseError]
type accountCfdTunnelGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountCfdTunnelGetResponseMessageJSON `json:"-"`
}

// accountCfdTunnelGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountCfdTunnelGetResponseMessage]
type accountCfdTunnelGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnel] or
// [AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnel].
type AccountCfdTunnelGetResponseResult interface {
	implementsAccountCfdTunnelGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountCfdTunnelGetResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelGetResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountCfdTunnelGetResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct [AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnel]
type accountCfdTunnelGetResponseResultXk6JhoizCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnel) implementsAccountCfdTunnelGetResponseResult() {
}

type AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                           `json:"uuid"`
	JSON accountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnectionJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnection]
type accountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountCfdTunnelGetResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnel]
type accountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountCfdTunnelGetResponseResult() {
}

type AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                     `json:"uuid"`
	JSON accountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountCfdTunnelGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountCfdTunnelGetResponseSuccess bool

const (
	AccountCfdTunnelGetResponseSuccessTrue AccountCfdTunnelGetResponseSuccess = true
)

type AccountCfdTunnelUpdateResponse struct {
	Errors   []AccountCfdTunnelUpdateResponseError   `json:"errors"`
	Messages []AccountCfdTunnelUpdateResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountCfdTunnelUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelUpdateResponseSuccess `json:"success"`
	JSON    accountCfdTunnelUpdateResponseJSON    `json:"-"`
}

// accountCfdTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [AccountCfdTunnelUpdateResponse]
type accountCfdTunnelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountCfdTunnelUpdateResponseErrorJSON `json:"-"`
}

// accountCfdTunnelUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountCfdTunnelUpdateResponseError]
type accountCfdTunnelUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountCfdTunnelUpdateResponseMessageJSON `json:"-"`
}

// accountCfdTunnelUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountCfdTunnelUpdateResponseMessage]
type accountCfdTunnelUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnel] or
// [AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnel].
type AccountCfdTunnelUpdateResponseResult interface {
	implementsAccountCfdTunnelUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountCfdTunnelUpdateResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct [AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnel]
type accountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnel) implementsAccountCfdTunnelUpdateResponseResult() {
}

type AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                              `json:"uuid"`
	JSON accountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnectionJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnection]
type accountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountCfdTunnelUpdateResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnel]
type accountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountCfdTunnelUpdateResponseResult() {
}

type AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                        `json:"uuid"`
	JSON accountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountCfdTunnelUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountCfdTunnelUpdateResponseSuccess bool

const (
	AccountCfdTunnelUpdateResponseSuccessTrue AccountCfdTunnelUpdateResponseSuccess = true
)

type AccountCfdTunnelDeleteResponse struct {
	Errors   []AccountCfdTunnelDeleteResponseError   `json:"errors"`
	Messages []AccountCfdTunnelDeleteResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountCfdTunnelDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelDeleteResponseSuccess `json:"success"`
	JSON    accountCfdTunnelDeleteResponseJSON    `json:"-"`
}

// accountCfdTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [AccountCfdTunnelDeleteResponse]
type accountCfdTunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountCfdTunnelDeleteResponseErrorJSON `json:"-"`
}

// accountCfdTunnelDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountCfdTunnelDeleteResponseError]
type accountCfdTunnelDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountCfdTunnelDeleteResponseMessageJSON `json:"-"`
}

// accountCfdTunnelDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountCfdTunnelDeleteResponseMessage]
type accountCfdTunnelDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnel] or
// [AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnel].
type AccountCfdTunnelDeleteResponseResult interface {
	implementsAccountCfdTunnelDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountCfdTunnelDeleteResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct [AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnel]
type accountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnel) implementsAccountCfdTunnelDeleteResponseResult() {
}

type AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                              `json:"uuid"`
	JSON accountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnectionJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnection]
type accountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountCfdTunnelDeleteResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnel]
type accountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountCfdTunnelDeleteResponseResult() {
}

type AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                        `json:"uuid"`
	JSON accountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountCfdTunnelDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountCfdTunnelDeleteResponseSuccess bool

const (
	AccountCfdTunnelDeleteResponseSuccessTrue AccountCfdTunnelDeleteResponseSuccess = true
)

type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse struct {
	Errors   []AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseError   `json:"errors"`
	Messages []AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseSuccess `json:"success"`
	JSON    accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseJSON    `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseJSON contains the
// JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseErrorJSON `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseError]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessageJSON `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessage]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnel]
// or
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnel].
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResult interface {
	implementsAccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnel]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnel) implementsAccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResult() {
}

type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                                            `json:"uuid"`
	JSON accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnection]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnel]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResult() {
}

type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                                                      `json:"uuid"`
	JSON accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseSuccess bool

const (
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseSuccessTrue AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelResponseSuccess = true
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnel]
// or
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnel].
type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse interface {
	implementsAccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnel]
type accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnel) implementsAccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse() {
}

type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                                       `json:"uuid"`
	JSON accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnection]
type accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType string

const (
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunTypeWarpConnector AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunTypeIPSec         AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunTypeGre           AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType = "gre"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunTypeCni           AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnel]
type accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnel) implementsAccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponse() {
}

type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                                                 `json:"uuid"`
	JSON accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnection]
type accountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeGre           AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeCni           AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "cni"
)

type AccountCfdTunnelUpdateParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r AccountCfdTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCfdTunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountCfdTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard or using the
	// [Cloudflare Tunnel configuration](https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties)
	// endpoint.
	ConfigSrc param.Field[AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc] `json:"config_src"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard or using the
// [Cloudflare Tunnel configuration](https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties)
// endpoint.
type AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc string

const (
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrcLocal      AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc = "local"
	AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrcCloudflare AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc = "cloudflare"
)

type AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsParams struct {
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt     param.Field[time.Time] `query:"existed_at" format:"date-time"`
	IncludePrefix param.Field[string]    `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage       param.Field[float64]   `query:"per_page"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes
// [AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsParams]'s query parameters
// as `url.Values`.
func (r AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
