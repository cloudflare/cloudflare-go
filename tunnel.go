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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// TunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTunnelService] method instead.
type TunnelService struct {
	Options     []option.RequestOption
	Connections *TunnelConnectionService
}

// NewTunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTunnelService(opts ...option.RequestOption) (r *TunnelService) {
	r = &TunnelService{}
	r.Options = opts
	r.Connections = NewTunnelConnectionService(opts...)
	return
}

// Fetches a single Argo Tunnel.
func (r *TunnelService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Argo Tunnel from an account.
func (r *TunnelService) Delete(ctx context.Context, accountID string, tunnelID string, body TunnelDeleteParams, opts ...option.RequestOption) (res *TunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Creates a new Argo Tunnel in an account.
func (r *TunnelService) ArgoTunnelNewAnArgoTunnel(ctx context.Context, accountID string, body TunnelArgoTunnelNewAnArgoTunnelParams, opts ...option.RequestOption) (res *TunnelArgoTunnelNewAnArgoTunnelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) ArgoTunnelListArgoTunnels(ctx context.Context, accountID string, query TunnelArgoTunnelListArgoTunnelsParams, opts ...option.RequestOption) (res *TunnelArgoTunnelListArgoTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TunnelGetResponse struct {
	Errors   []TunnelGetResponseError   `json:"errors"`
	Messages []TunnelGetResponseMessage `json:"messages"`
	Result   TunnelGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelGetResponseSuccess `json:"success"`
	JSON    tunnelGetResponseJSON    `json:"-"`
}

// tunnelGetResponseJSON contains the JSON metadata for the struct
// [TunnelGetResponse]
type tunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    tunnelGetResponseErrorJSON `json:"-"`
}

// tunnelGetResponseErrorJSON contains the JSON metadata for the struct
// [TunnelGetResponseError]
type tunnelGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    tunnelGetResponseMessageJSON `json:"-"`
}

// tunnelGetResponseMessageJSON contains the JSON metadata for the struct
// [TunnelGetResponseMessage]
type tunnelGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelGetResponseResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                   `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelGetResponseResultJSON `json:"-"`
}

// tunnelGetResponseResultJSON contains the JSON metadata for the struct
// [TunnelGetResponseResult]
type tunnelGetResponseResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                `json:"uuid"`
	JSON tunnelGetResponseResultConnectionJSON `json:"-"`
}

// tunnelGetResponseResultConnectionJSON contains the JSON metadata for the struct
// [TunnelGetResponseResultConnection]
type tunnelGetResponseResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelGetResponseResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelGetResponseSuccess bool

const (
	TunnelGetResponseSuccessTrue TunnelGetResponseSuccess = true
)

type TunnelDeleteResponse struct {
	Errors   []TunnelDeleteResponseError   `json:"errors"`
	Messages []TunnelDeleteResponseMessage `json:"messages"`
	Result   TunnelDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelDeleteResponseSuccess `json:"success"`
	JSON    tunnelDeleteResponseJSON    `json:"-"`
}

// tunnelDeleteResponseJSON contains the JSON metadata for the struct
// [TunnelDeleteResponse]
type tunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    tunnelDeleteResponseErrorJSON `json:"-"`
}

// tunnelDeleteResponseErrorJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseError]
type tunnelDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    tunnelDeleteResponseMessageJSON `json:"-"`
}

// tunnelDeleteResponseMessageJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseMessage]
type tunnelDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelDeleteResponseResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                      `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelDeleteResponseResultJSON `json:"-"`
}

// tunnelDeleteResponseResultJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseResult]
type tunnelDeleteResponseResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                   `json:"uuid"`
	JSON tunnelDeleteResponseResultConnectionJSON `json:"-"`
}

// tunnelDeleteResponseResultConnectionJSON contains the JSON metadata for the
// struct [TunnelDeleteResponseResultConnection]
type tunnelDeleteResponseResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelDeleteResponseResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelDeleteResponseSuccess bool

const (
	TunnelDeleteResponseSuccessTrue TunnelDeleteResponseSuccess = true
)

type TunnelArgoTunnelNewAnArgoTunnelResponse struct {
	Errors   []TunnelArgoTunnelNewAnArgoTunnelResponseError   `json:"errors"`
	Messages []TunnelArgoTunnelNewAnArgoTunnelResponseMessage `json:"messages"`
	Result   TunnelArgoTunnelNewAnArgoTunnelResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TunnelArgoTunnelNewAnArgoTunnelResponseSuccess `json:"success"`
	JSON    tunnelArgoTunnelNewAnArgoTunnelResponseJSON    `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseJSON contains the JSON metadata for the
// struct [TunnelArgoTunnelNewAnArgoTunnelResponse]
type tunnelArgoTunnelNewAnArgoTunnelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    tunnelArgoTunnelNewAnArgoTunnelResponseErrorJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseErrorJSON contains the JSON metadata for
// the struct [TunnelArgoTunnelNewAnArgoTunnelResponseError]
type tunnelArgoTunnelNewAnArgoTunnelResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    tunnelArgoTunnelNewAnArgoTunnelResponseMessageJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseMessageJSON contains the JSON metadata
// for the struct [TunnelArgoTunnelNewAnArgoTunnelResponseMessage]
type tunnelArgoTunnelNewAnArgoTunnelResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelArgoTunnelNewAnArgoTunnelResponseResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                                         `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelArgoTunnelNewAnArgoTunnelResponseResultJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseResultJSON contains the JSON metadata for
// the struct [TunnelArgoTunnelNewAnArgoTunnelResponseResult]
type tunnelArgoTunnelNewAnArgoTunnelResponseResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                      `json:"uuid"`
	JSON tunnelArgoTunnelNewAnArgoTunnelResponseResultConnectionJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseResultConnectionJSON contains the JSON
// metadata for the struct
// [TunnelArgoTunnelNewAnArgoTunnelResponseResultConnection]
type tunnelArgoTunnelNewAnArgoTunnelResponseResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelArgoTunnelNewAnArgoTunnelResponseSuccess bool

const (
	TunnelArgoTunnelNewAnArgoTunnelResponseSuccessTrue TunnelArgoTunnelNewAnArgoTunnelResponseSuccess = true
)

type TunnelArgoTunnelListArgoTunnelsResponse struct {
	Errors     []TunnelArgoTunnelListArgoTunnelsResponseError    `json:"errors"`
	Messages   []TunnelArgoTunnelListArgoTunnelsResponseMessage  `json:"messages"`
	Result     []TunnelArgoTunnelListArgoTunnelsResponseResult   `json:"result"`
	ResultInfo TunnelArgoTunnelListArgoTunnelsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success TunnelArgoTunnelListArgoTunnelsResponseSuccess `json:"success"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseJSON    `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseJSON contains the JSON metadata for the
// struct [TunnelArgoTunnelListArgoTunnelsResponse]
type tunnelArgoTunnelListArgoTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelListArgoTunnelsResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseErrorJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseErrorJSON contains the JSON metadata for
// the struct [TunnelArgoTunnelListArgoTunnelsResponseError]
type tunnelArgoTunnelListArgoTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelListArgoTunnelsResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseMessageJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseMessageJSON contains the JSON metadata
// for the struct [TunnelArgoTunnelListArgoTunnelsResponseMessage]
type tunnelArgoTunnelListArgoTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnel] or
// [TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnel].
type TunnelArgoTunnelListArgoTunnelsResponseResult interface {
	implementsTunnelArgoTunnelListArgoTunnelsResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TunnelArgoTunnelListArgoTunnelsResponseResult)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnection `json:"connections"`
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
	TunType TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType `json:"tun_type"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelJSON    `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelJSON contains the
// JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnel]
type tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelJSON struct {
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

func (r *TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnel) implementsTunnelArgoTunnelListArgoTunnelsResponseResult() {
}

type TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnection struct {
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
	Uuid string                                                                       `json:"uuid"`
	JSON tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnectionJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnection]
type tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnectionJSON struct {
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

func (r *TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType string

const (
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunTypeCfdTunnel     TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType = "cfd_tunnel"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunTypeWarpConnector TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType = "warp_connector"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunTypeIPSec         TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType = "ip_sec"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunTypeGre           TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType = "gre"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunTypeCni           TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnection `json:"connections"`
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
	TunType TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelJSON    `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelJSON
// contains the JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnel]
type tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelJSON struct {
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

func (r *TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnel) implementsTunnelArgoTunnelListArgoTunnelsResponseResult() {
}

type TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnection struct {
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
	Uuid string                                                                                 `json:"uuid"`
	JSON tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnection]
type tunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnectionJSON struct {
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

func (r *TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType string

const (
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunTypeCfdTunnel     TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType = "cfd_tunnel"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunTypeWarpConnector TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType = "warp_connector"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunTypeIPSec         TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType = "ip_sec"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunTypeGre           TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType = "gre"
	TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunTypeCni           TunnelArgoTunnelListArgoTunnelsResponseResultXlFXheKkWarpConnectorTunnelTunType = "cni"
)

type TunnelArgoTunnelListArgoTunnelsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       tunnelArgoTunnelListArgoTunnelsResponseResultInfoJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseResultInfoJSON contains the JSON metadata
// for the struct [TunnelArgoTunnelListArgoTunnelsResponseResultInfo]
type tunnelArgoTunnelListArgoTunnelsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelArgoTunnelListArgoTunnelsResponseSuccess bool

const (
	TunnelArgoTunnelListArgoTunnelsResponseSuccessTrue TunnelArgoTunnelListArgoTunnelsResponseSuccess = true
)

type TunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TunnelArgoTunnelNewAnArgoTunnelParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Sets the password required to run the tunnel. Must be at least 32 bytes and
	// encoded as a base64 string.
	TunnelSecret param.Field[interface{}] `json:"tunnel_secret,required"`
}

func (r TunnelArgoTunnelNewAnArgoTunnelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelArgoTunnelListArgoTunnelsParams struct {
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

// URLQuery serializes [TunnelArgoTunnelListArgoTunnelsParams]'s query parameters
// as `url.Values`.
func (r TunnelArgoTunnelListArgoTunnelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
