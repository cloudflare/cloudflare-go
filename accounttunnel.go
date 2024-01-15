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

// AccountTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountTunnelService] method
// instead.
type AccountTunnelService struct {
	Options     []option.RequestOption
	Connections *AccountTunnelConnectionService
}

// NewAccountTunnelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTunnelService(opts ...option.RequestOption) (r *AccountTunnelService) {
	r = &AccountTunnelService{}
	r.Options = opts
	r.Connections = NewAccountTunnelConnectionService(opts...)
	return
}

// Fetches a single Argo Tunnel.
func (r *AccountTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *AccountTunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Argo Tunnel from an account.
func (r *AccountTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountTunnelDeleteParams, opts ...option.RequestOption) (res *AccountTunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Creates a new Argo Tunnel in an account.
func (r *AccountTunnelService) ArgoTunnelNewAnArgoTunnel(ctx context.Context, accountIdentifier string, body AccountTunnelArgoTunnelNewAnArgoTunnelParams, opts ...option.RequestOption) (res *AccountTunnelArgoTunnelNewAnArgoTunnelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *AccountTunnelService) ArgoTunnelListArgoTunnels(ctx context.Context, accountIdentifier string, query AccountTunnelArgoTunnelListArgoTunnelsParams, opts ...option.RequestOption) (res *shared.Page[AccountTunnelArgoTunnelListArgoTunnelsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountIdentifier)
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

type AccountTunnelGetResponse struct {
	Errors   []AccountTunnelGetResponseError   `json:"errors"`
	Messages []AccountTunnelGetResponseMessage `json:"messages"`
	Result   AccountTunnelGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTunnelGetResponseSuccess `json:"success"`
	JSON    accountTunnelGetResponseJSON    `json:"-"`
}

// accountTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountTunnelGetResponse]
type accountTunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accountTunnelGetResponseErrorJSON `json:"-"`
}

// accountTunnelGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountTunnelGetResponseError]
type accountTunnelGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountTunnelGetResponseMessageJSON `json:"-"`
}

// accountTunnelGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountTunnelGetResponseMessage]
type accountTunnelGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelGetResponseResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountTunnelGetResponseResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                          `json:"deleted_at,nullable" format:"date-time"`
	JSON      accountTunnelGetResponseResultJSON `json:"-"`
}

// accountTunnelGetResponseResultJSON contains the JSON metadata for the struct
// [AccountTunnelGetResponseResult]
type accountTunnelGetResponseResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelGetResponseResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                       `json:"uuid"`
	JSON accountTunnelGetResponseResultConnectionJSON `json:"-"`
}

// accountTunnelGetResponseResultConnectionJSON contains the JSON metadata for the
// struct [AccountTunnelGetResponseResultConnection]
type accountTunnelGetResponseResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountTunnelGetResponseResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTunnelGetResponseSuccess bool

const (
	AccountTunnelGetResponseSuccessTrue AccountTunnelGetResponseSuccess = true
)

type AccountTunnelDeleteResponse struct {
	Errors   []AccountTunnelDeleteResponseError   `json:"errors"`
	Messages []AccountTunnelDeleteResponseMessage `json:"messages"`
	Result   AccountTunnelDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTunnelDeleteResponseSuccess `json:"success"`
	JSON    accountTunnelDeleteResponseJSON    `json:"-"`
}

// accountTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [AccountTunnelDeleteResponse]
type accountTunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelDeleteResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountTunnelDeleteResponseErrorJSON `json:"-"`
}

// accountTunnelDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountTunnelDeleteResponseError]
type accountTunnelDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelDeleteResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountTunnelDeleteResponseMessageJSON `json:"-"`
}

// accountTunnelDeleteResponseMessageJSON contains the JSON metadata for the struct
// [AccountTunnelDeleteResponseMessage]
type accountTunnelDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelDeleteResponseResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountTunnelDeleteResponseResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                             `json:"deleted_at,nullable" format:"date-time"`
	JSON      accountTunnelDeleteResponseResultJSON `json:"-"`
}

// accountTunnelDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccountTunnelDeleteResponseResult]
type accountTunnelDeleteResponseResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelDeleteResponseResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                          `json:"uuid"`
	JSON accountTunnelDeleteResponseResultConnectionJSON `json:"-"`
}

// accountTunnelDeleteResponseResultConnectionJSON contains the JSON metadata for
// the struct [AccountTunnelDeleteResponseResultConnection]
type accountTunnelDeleteResponseResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountTunnelDeleteResponseResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTunnelDeleteResponseSuccess bool

const (
	AccountTunnelDeleteResponseSuccessTrue AccountTunnelDeleteResponseSuccess = true
)

type AccountTunnelArgoTunnelNewAnArgoTunnelResponse struct {
	Errors   []AccountTunnelArgoTunnelNewAnArgoTunnelResponseError   `json:"errors"`
	Messages []AccountTunnelArgoTunnelNewAnArgoTunnelResponseMessage `json:"messages"`
	Result   AccountTunnelArgoTunnelNewAnArgoTunnelResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountTunnelArgoTunnelNewAnArgoTunnelResponseSuccess `json:"success"`
	JSON    accountTunnelArgoTunnelNewAnArgoTunnelResponseJSON    `json:"-"`
}

// accountTunnelArgoTunnelNewAnArgoTunnelResponseJSON contains the JSON metadata
// for the struct [AccountTunnelArgoTunnelNewAnArgoTunnelResponse]
type accountTunnelArgoTunnelNewAnArgoTunnelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelArgoTunnelNewAnArgoTunnelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelArgoTunnelNewAnArgoTunnelResponseError struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountTunnelArgoTunnelNewAnArgoTunnelResponseErrorJSON `json:"-"`
}

// accountTunnelArgoTunnelNewAnArgoTunnelResponseErrorJSON contains the JSON
// metadata for the struct [AccountTunnelArgoTunnelNewAnArgoTunnelResponseError]
type accountTunnelArgoTunnelNewAnArgoTunnelResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelArgoTunnelNewAnArgoTunnelResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelArgoTunnelNewAnArgoTunnelResponseMessage struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountTunnelArgoTunnelNewAnArgoTunnelResponseMessageJSON `json:"-"`
}

// accountTunnelArgoTunnelNewAnArgoTunnelResponseMessageJSON contains the JSON
// metadata for the struct [AccountTunnelArgoTunnelNewAnArgoTunnelResponseMessage]
type accountTunnelArgoTunnelNewAnArgoTunnelResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelArgoTunnelNewAnArgoTunnelResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelArgoTunnelNewAnArgoTunnelResponseResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                                                `json:"deleted_at,nullable" format:"date-time"`
	JSON      accountTunnelArgoTunnelNewAnArgoTunnelResponseResultJSON `json:"-"`
}

// accountTunnelArgoTunnelNewAnArgoTunnelResponseResultJSON contains the JSON
// metadata for the struct [AccountTunnelArgoTunnelNewAnArgoTunnelResponseResult]
type accountTunnelArgoTunnelNewAnArgoTunnelResponseResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelArgoTunnelNewAnArgoTunnelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                             `json:"uuid"`
	JSON accountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnectionJSON `json:"-"`
}

// accountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnectionJSON contains the
// JSON metadata for the struct
// [AccountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnection]
type accountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountTunnelArgoTunnelNewAnArgoTunnelResponseResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTunnelArgoTunnelNewAnArgoTunnelResponseSuccess bool

const (
	AccountTunnelArgoTunnelNewAnArgoTunnelResponseSuccessTrue AccountTunnelArgoTunnelNewAnArgoTunnelResponseSuccess = true
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnel] or
// [AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnel].
type AccountTunnelArgoTunnelListArgoTunnelsResponse interface {
	implementsAccountTunnelArgoTunnelListArgoTunnelsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountTunnelArgoTunnelListArgoTunnelsResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnection `json:"connections"`
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
	TunType AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType `json:"tun_type"`
	JSON    accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelJSON    `json:"-"`
}

// accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelJSON contains the
// JSON metadata for the struct
// [AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnel]
type accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelJSON struct {
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

func (r *AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnel) implementsAccountTunnelArgoTunnelListArgoTunnelsResponse() {
}

type AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnection struct {
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
	JSON accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnectionJSON `json:"-"`
}

// accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnection]
type accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnectionJSON struct {
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

func (r *AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType string

const (
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunTypeCfdTunnel     AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType = "cfd_tunnel"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunTypeWarpConnector AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType = "warp_connector"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunTypeIPSec         AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType = "ip_sec"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunTypeGre           AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType = "gre"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunTypeCni           AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnection `json:"connections"`
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
	TunType AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelJSON    `json:"-"`
}

// accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelJSON
// contains the JSON metadata for the struct
// [AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnel]
type accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelJSON struct {
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

func (r *AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnel) implementsAccountTunnelArgoTunnelListArgoTunnelsResponse() {
}

type AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                                  `json:"uuid"`
	JSON accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnectionJSON `json:"-"`
}

// accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnection]
type accountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnectionJSON struct {
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

func (r *AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType string

const (
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeCfdTunnel     AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "cfd_tunnel"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeWarpConnector AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "warp_connector"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeIPSec         AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "ip_sec"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeGre           AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "gre"
	AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunTypeCni           AccountTunnelArgoTunnelListArgoTunnelsResponseXk6JhoizWarpConnectorTunnelTunType = "cni"
)

type AccountTunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountTunnelArgoTunnelNewAnArgoTunnelParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Sets the password required to run the tunnel. Must be at least 32 bytes and
	// encoded as a base64 string.
	TunnelSecret param.Field[interface{}] `json:"tunnel_secret,required"`
}

func (r AccountTunnelArgoTunnelNewAnArgoTunnelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTunnelArgoTunnelListArgoTunnelsParams struct {
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
	PerPage param.Field[float64] `query:"per_page"`
	// The types of tunnels to filter separated by a comma.
	TunTypes      param.Field[string]    `query:"tun_types"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [AccountTunnelArgoTunnelListArgoTunnelsParams]'s query
// parameters as `url.Values`.
func (r AccountTunnelArgoTunnelListArgoTunnelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
