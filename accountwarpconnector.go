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

// AccountWarpConnectorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountWarpConnectorService]
// method instead.
type AccountWarpConnectorService struct {
	Options []option.RequestOption
}

// NewAccountWarpConnectorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWarpConnectorService(opts ...option.RequestOption) (r *AccountWarpConnectorService) {
	r = &AccountWarpConnectorService{}
	r.Options = opts
	return
}

// Creates a new Warp Connector Tunnel in an account.
func (r *AccountWarpConnectorService) New(ctx context.Context, accountIdentifier string, body AccountWarpConnectorNewParams, opts ...option.RequestOption) (res *AccountWarpConnectorNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Warp Connector Tunnel.
func (r *AccountWarpConnectorService) Get(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *AccountWarpConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Warp Connector Tunnel.
func (r *AccountWarpConnectorService) Update(ctx context.Context, accountIdentifier string, tunnelID string, body AccountWarpConnectorUpdateParams, opts ...option.RequestOption) (res *AccountWarpConnectorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *AccountWarpConnectorService) List(ctx context.Context, accountIdentifier string, query AccountWarpConnectorListParams, opts ...option.RequestOption) (res *shared.Page[AccountWarpConnectorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector", accountIdentifier)
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

// Deletes a Warp Connector Tunnel from an account.
func (r *AccountWarpConnectorService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountWarpConnectorDeleteParams, opts ...option.RequestOption) (res *AccountWarpConnectorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Gets the token used to associate warp device with a specific Warp Connector
// tunnel.
func (r *AccountWarpConnectorService) Token(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *AccountWarpConnectorTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/token", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWarpConnectorNewResponse struct {
	Errors   []AccountWarpConnectorNewResponseError   `json:"errors"`
	Messages []AccountWarpConnectorNewResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountWarpConnectorNewResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountWarpConnectorNewResponseSuccess `json:"success"`
	JSON    accountWarpConnectorNewResponseJSON    `json:"-"`
}

// accountWarpConnectorNewResponseJSON contains the JSON metadata for the struct
// [AccountWarpConnectorNewResponse]
type accountWarpConnectorNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorNewResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountWarpConnectorNewResponseErrorJSON `json:"-"`
}

// accountWarpConnectorNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountWarpConnectorNewResponseError]
type accountWarpConnectorNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorNewResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountWarpConnectorNewResponseMessageJSON `json:"-"`
}

// accountWarpConnectorNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountWarpConnectorNewResponseMessage]
type accountWarpConnectorNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnel] or
// [AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnel].
type AccountWarpConnectorNewResponseResult interface {
	implementsAccountWarpConnectorNewResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWarpConnectorNewResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorNewResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountWarpConnectorNewResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct [AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnel]
type accountWarpConnectorNewResponseResultXk6JhoizCfdTunnelJSON struct {
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

func (r *AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnel) implementsAccountWarpConnectorNewResponseResult() {
}

type AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnection struct {
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
	Uuid string                                                               `json:"uuid"`
	JSON accountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnectionJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnection]
type accountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountWarpConnectorNewResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnel]
type accountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
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

func (r *AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountWarpConnectorNewResponseResult() {
}

type AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                         `json:"uuid"`
	JSON accountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountWarpConnectorNewResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountWarpConnectorNewResponseSuccess bool

const (
	AccountWarpConnectorNewResponseSuccessTrue AccountWarpConnectorNewResponseSuccess = true
)

type AccountWarpConnectorGetResponse struct {
	Errors   []AccountWarpConnectorGetResponseError   `json:"errors"`
	Messages []AccountWarpConnectorGetResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountWarpConnectorGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountWarpConnectorGetResponseSuccess `json:"success"`
	JSON    accountWarpConnectorGetResponseJSON    `json:"-"`
}

// accountWarpConnectorGetResponseJSON contains the JSON metadata for the struct
// [AccountWarpConnectorGetResponse]
type accountWarpConnectorGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountWarpConnectorGetResponseErrorJSON `json:"-"`
}

// accountWarpConnectorGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountWarpConnectorGetResponseError]
type accountWarpConnectorGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountWarpConnectorGetResponseMessageJSON `json:"-"`
}

// accountWarpConnectorGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountWarpConnectorGetResponseMessage]
type accountWarpConnectorGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnel] or
// [AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnel].
type AccountWarpConnectorGetResponseResult interface {
	implementsAccountWarpConnectorGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWarpConnectorGetResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorGetResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountWarpConnectorGetResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct [AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnel]
type accountWarpConnectorGetResponseResultXk6JhoizCfdTunnelJSON struct {
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

func (r *AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnel) implementsAccountWarpConnectorGetResponseResult() {
}

type AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnection struct {
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
	Uuid string                                                               `json:"uuid"`
	JSON accountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnectionJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnection]
type accountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountWarpConnectorGetResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnel]
type accountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
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

func (r *AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountWarpConnectorGetResponseResult() {
}

type AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                         `json:"uuid"`
	JSON accountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountWarpConnectorGetResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountWarpConnectorGetResponseSuccess bool

const (
	AccountWarpConnectorGetResponseSuccessTrue AccountWarpConnectorGetResponseSuccess = true
)

type AccountWarpConnectorUpdateResponse struct {
	Errors   []AccountWarpConnectorUpdateResponseError   `json:"errors"`
	Messages []AccountWarpConnectorUpdateResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountWarpConnectorUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountWarpConnectorUpdateResponseSuccess `json:"success"`
	JSON    accountWarpConnectorUpdateResponseJSON    `json:"-"`
}

// accountWarpConnectorUpdateResponseJSON contains the JSON metadata for the struct
// [AccountWarpConnectorUpdateResponse]
type accountWarpConnectorUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWarpConnectorUpdateResponseErrorJSON `json:"-"`
}

// accountWarpConnectorUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountWarpConnectorUpdateResponseError]
type accountWarpConnectorUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountWarpConnectorUpdateResponseMessageJSON `json:"-"`
}

// accountWarpConnectorUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountWarpConnectorUpdateResponseMessage]
type accountWarpConnectorUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnel]
// or [AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnel].
type AccountWarpConnectorUpdateResponseResult interface {
	implementsAccountWarpConnectorUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWarpConnectorUpdateResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct
// [AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnel]
type accountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelJSON struct {
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

func (r *AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnel) implementsAccountWarpConnectorUpdateResponseResult() {
}

type AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnection struct {
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
	Uuid string                                                                  `json:"uuid"`
	JSON accountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnectionJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnection]
type accountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountWarpConnectorUpdateResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnel]
type accountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
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

func (r *AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountWarpConnectorUpdateResponseResult() {
}

type AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                            `json:"uuid"`
	JSON accountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountWarpConnectorUpdateResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountWarpConnectorUpdateResponseSuccess bool

const (
	AccountWarpConnectorUpdateResponseSuccessTrue AccountWarpConnectorUpdateResponseSuccess = true
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountWarpConnectorListResponseXk6JhoizCfdTunnel] or
// [AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnel].
type AccountWarpConnectorListResponse interface {
	implementsAccountWarpConnectorListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWarpConnectorListResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorListResponseXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorListResponseXk6JhoizCfdTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorListResponseXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountWarpConnectorListResponseXk6JhoizCfdTunnelJSON contains the JSON metadata
// for the struct [AccountWarpConnectorListResponseXk6JhoizCfdTunnel]
type accountWarpConnectorListResponseXk6JhoizCfdTunnelJSON struct {
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

func (r *AccountWarpConnectorListResponseXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorListResponseXk6JhoizCfdTunnel) implementsAccountWarpConnectorListResponse() {
}

type AccountWarpConnectorListResponseXk6JhoizCfdTunnelConnection struct {
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
	Uuid string                                                          `json:"uuid"`
	JSON accountWarpConnectorListResponseXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorListResponseXk6JhoizCfdTunnelConnectionJSON contains the
// JSON metadata for the struct
// [AccountWarpConnectorListResponseXk6JhoizCfdTunnelConnection]
type accountWarpConnectorListResponseXk6JhoizCfdTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorListResponseXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType string

const (
	AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunTypeWarpConnector AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunTypeIPSec         AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunTypeGre           AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType = "gre"
	AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunTypeCni           AccountWarpConnectorListResponseXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelJSON contains the
// JSON metadata for the struct
// [AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnel]
type accountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelJSON struct {
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

func (r *AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnel) implementsAccountWarpConnectorListResponse() {
}

type AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                    `json:"uuid"`
	JSON accountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnection]
type accountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunTypeGre           AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunTypeCni           AccountWarpConnectorListResponseXk6JhoizWarpConnectorTunnelTunType = "cni"
)

type AccountWarpConnectorDeleteResponse struct {
	Errors   []AccountWarpConnectorDeleteResponseError   `json:"errors"`
	Messages []AccountWarpConnectorDeleteResponseMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result AccountWarpConnectorDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountWarpConnectorDeleteResponseSuccess `json:"success"`
	JSON    accountWarpConnectorDeleteResponseJSON    `json:"-"`
}

// accountWarpConnectorDeleteResponseJSON contains the JSON metadata for the struct
// [AccountWarpConnectorDeleteResponse]
type accountWarpConnectorDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWarpConnectorDeleteResponseErrorJSON `json:"-"`
}

// accountWarpConnectorDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountWarpConnectorDeleteResponseError]
type accountWarpConnectorDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountWarpConnectorDeleteResponseMessageJSON `json:"-"`
}

// accountWarpConnectorDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountWarpConnectorDeleteResponseMessage]
type accountWarpConnectorDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnel]
// or [AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnel].
type AccountWarpConnectorDeleteResponseResult interface {
	implementsAccountWarpConnectorDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWarpConnectorDeleteResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelJSON contains the JSON
// metadata for the struct
// [AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnel]
type accountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelJSON struct {
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

func (r *AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnel) implementsAccountWarpConnectorDeleteResponseResult() {
}

type AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnection struct {
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
	Uuid string                                                                  `json:"uuid"`
	JSON accountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnectionJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnection]
type accountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType string

const (
	AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunTypeWarpConnector AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunTypeIPSec         AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunTypeGre           AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType = "gre"
	AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunTypeCni           AccountWarpConnectorDeleteResponseResultXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
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
	TunType AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelJSON contains
// the JSON metadata for the struct
// [AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnel]
type accountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelJSON struct {
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

func (r *AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnel) implementsAccountWarpConnectorDeleteResponseResult() {
}

type AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                            `json:"uuid"`
	JSON accountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection]
type accountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnectionJSON struct {
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

func (r *AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeGre           AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunTypeCni           AccountWarpConnectorDeleteResponseResultXk6JhoizWarpConnectorTunnelTunType = "cni"
)

// Whether the API call was successful
type AccountWarpConnectorDeleteResponseSuccess bool

const (
	AccountWarpConnectorDeleteResponseSuccessTrue AccountWarpConnectorDeleteResponseSuccess = true
)

type AccountWarpConnectorTokenResponse struct {
	Errors   []AccountWarpConnectorTokenResponseError   `json:"errors"`
	Messages []AccountWarpConnectorTokenResponseMessage `json:"messages"`
	Result   string                                     `json:"result"`
	// Whether the API call was successful
	Success AccountWarpConnectorTokenResponseSuccess `json:"success"`
	JSON    accountWarpConnectorTokenResponseJSON    `json:"-"`
}

// accountWarpConnectorTokenResponseJSON contains the JSON metadata for the struct
// [AccountWarpConnectorTokenResponse]
type accountWarpConnectorTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorTokenResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountWarpConnectorTokenResponseErrorJSON `json:"-"`
}

// accountWarpConnectorTokenResponseErrorJSON contains the JSON metadata for the
// struct [AccountWarpConnectorTokenResponseError]
type accountWarpConnectorTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWarpConnectorTokenResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountWarpConnectorTokenResponseMessageJSON `json:"-"`
}

// accountWarpConnectorTokenResponseMessageJSON contains the JSON metadata for the
// struct [AccountWarpConnectorTokenResponseMessage]
type accountWarpConnectorTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWarpConnectorTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWarpConnectorTokenResponseSuccess bool

const (
	AccountWarpConnectorTokenResponseSuccessTrue AccountWarpConnectorTokenResponseSuccess = true
)

type AccountWarpConnectorNewParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountWarpConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWarpConnectorUpdateParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r AccountWarpConnectorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWarpConnectorListParams struct {
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

// URLQuery serializes [AccountWarpConnectorListParams]'s query parameters as
// `url.Values`.
func (r AccountWarpConnectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWarpConnectorDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWarpConnectorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
