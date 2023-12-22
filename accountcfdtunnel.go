// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
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
	r.Tokens = NewAccountCfdTunnelTokenService(opts...)
	return
}

// Fetches a single Cloudflare Tunnel.
func (r *AccountCfdTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *SchemasTunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *AccountCfdTunnelService) Update(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelUpdateParams, opts ...option.RequestOption) (res *SchemasTunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a Cloudflare Tunnel from an account.
func (r *AccountCfdTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelDeleteParams, opts ...option.RequestOption) (res *SchemasTunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Creates a new Cloudflare Tunnel in an account.
func (r *AccountCfdTunnelService) CloudflareTunnelNewACloudflareTunnel(ctx context.Context, accountIdentifier string, body AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParams, opts ...option.RequestOption) (res *SchemasTunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters Cloudflare Tunnels in an account.
func (r *AccountCfdTunnelService) CloudflareTunnelListCloudflareTunnels(ctx context.Context, accountIdentifier string, query AccountCfdTunnelCloudflareTunnelListCloudflareTunnelsParams, opts ...option.RequestOption) (res *SchemasTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SchemasTunnelResponseCollection struct {
	Errors     []SchemasTunnelResponseCollectionError    `json:"errors"`
	Messages   []SchemasTunnelResponseCollectionMessage  `json:"messages"`
	Result     []SchemasTunnelResponseCollectionResult   `json:"result"`
	ResultInfo SchemasTunnelResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasTunnelResponseCollectionSuccess `json:"success"`
	JSON    schemasTunnelResponseCollectionJSON    `json:"-"`
}

// schemasTunnelResponseCollectionJSON contains the JSON metadata for the struct
// [SchemasTunnelResponseCollection]
type schemasTunnelResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelResponseCollectionError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasTunnelResponseCollectionErrorJSON `json:"-"`
}

// schemasTunnelResponseCollectionErrorJSON contains the JSON metadata for the
// struct [SchemasTunnelResponseCollectionError]
type schemasTunnelResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelResponseCollectionMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasTunnelResponseCollectionMessageJSON `json:"-"`
}

// schemasTunnelResponseCollectionMessageJSON contains the JSON metadata for the
// struct [SchemasTunnelResponseCollectionMessage]
type schemasTunnelResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type SchemasTunnelResponseCollectionResult struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []SchemasTunnelResponseCollectionResultConnection `json:"connections"`
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
	TunType SchemasTunnelResponseCollectionResultTunType `json:"tun_type"`
	JSON    schemasTunnelResponseCollectionResultJSON    `json:"-"`
}

// schemasTunnelResponseCollectionResultJSON contains the JSON metadata for the
// struct [SchemasTunnelResponseCollectionResult]
type schemasTunnelResponseCollectionResultJSON struct {
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

func (r *SchemasTunnelResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelResponseCollectionResultConnection struct {
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
	Uuid string                                              `json:"uuid"`
	JSON schemasTunnelResponseCollectionResultConnectionJSON `json:"-"`
}

// schemasTunnelResponseCollectionResultConnectionJSON contains the JSON metadata
// for the struct [SchemasTunnelResponseCollectionResultConnection]
type schemasTunnelResponseCollectionResultConnectionJSON struct {
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

func (r *SchemasTunnelResponseCollectionResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type SchemasTunnelResponseCollectionResultTunType string

const (
	SchemasTunnelResponseCollectionResultTunTypeCfdTunnel SchemasTunnelResponseCollectionResultTunType = "cfd_tunnel"
)

type SchemasTunnelResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       schemasTunnelResponseCollectionResultInfoJSON `json:"-"`
}

// schemasTunnelResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [SchemasTunnelResponseCollectionResultInfo]
type schemasTunnelResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasTunnelResponseCollectionSuccess bool

const (
	SchemasTunnelResponseCollectionSuccessTrue SchemasTunnelResponseCollectionSuccess = true
)

type SchemasTunnelResponseSingle struct {
	Errors   []SchemasTunnelResponseSingleError   `json:"errors"`
	Messages []SchemasTunnelResponseSingleMessage `json:"messages"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result SchemasTunnelResponseSingleResult `json:"result"`
	// Whether the API call was successful
	Success SchemasTunnelResponseSingleSuccess `json:"success"`
	JSON    schemasTunnelResponseSingleJSON    `json:"-"`
}

// schemasTunnelResponseSingleJSON contains the JSON metadata for the struct
// [SchemasTunnelResponseSingle]
type schemasTunnelResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelResponseSingleError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasTunnelResponseSingleErrorJSON `json:"-"`
}

// schemasTunnelResponseSingleErrorJSON contains the JSON metadata for the struct
// [SchemasTunnelResponseSingleError]
type schemasTunnelResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelResponseSingleMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasTunnelResponseSingleMessageJSON `json:"-"`
}

// schemasTunnelResponseSingleMessageJSON contains the JSON metadata for the struct
// [SchemasTunnelResponseSingleMessage]
type schemasTunnelResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasTunnelResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type SchemasTunnelResponseSingleResult struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []SchemasTunnelResponseSingleResultConnection `json:"connections"`
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
	TunType SchemasTunnelResponseSingleResultTunType `json:"tun_type"`
	JSON    schemasTunnelResponseSingleResultJSON    `json:"-"`
}

// schemasTunnelResponseSingleResultJSON contains the JSON metadata for the struct
// [SchemasTunnelResponseSingleResult]
type schemasTunnelResponseSingleResultJSON struct {
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

func (r *SchemasTunnelResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasTunnelResponseSingleResultConnection struct {
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
	Uuid string                                          `json:"uuid"`
	JSON schemasTunnelResponseSingleResultConnectionJSON `json:"-"`
}

// schemasTunnelResponseSingleResultConnectionJSON contains the JSON metadata for
// the struct [SchemasTunnelResponseSingleResultConnection]
type schemasTunnelResponseSingleResultConnectionJSON struct {
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

func (r *SchemasTunnelResponseSingleResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type SchemasTunnelResponseSingleResultTunType string

const (
	SchemasTunnelResponseSingleResultTunTypeCfdTunnel SchemasTunnelResponseSingleResultTunType = "cfd_tunnel"
)

// Whether the API call was successful
type SchemasTunnelResponseSingleSuccess bool

const (
	SchemasTunnelResponseSingleSuccessTrue SchemasTunnelResponseSingleSuccess = true
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
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret,required"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard or using the
	// [Cloudflare Tunnel configuration](https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties)
	// endpoint.
	ConfigSrc param.Field[AccountCfdTunnelCloudflareTunnelNewACloudflareTunnelParamsConfigSrc] `json:"config_src"`
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
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt param.Field[time.Time] `query:"existed_at" format:"date-time"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// A user-friendly name for the tunnel.
	TunnelName param.Field[string] `query:"tunnel_name"`
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
