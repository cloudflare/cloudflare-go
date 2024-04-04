// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/tidwall/gjson"
)

type ErrorData struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	JSON    errorDataJSON `json:"-"`
}

// errorDataJSON contains the JSON metadata for the struct [ErrorData]
type errorDataJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorDataJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9 struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  unnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9JSON `json:"-"`
}

// unnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9]
type unnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef2173d81a0b2d332c9e2ac46900fe8bb9JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    unnamedSchemaRef3248f24329456e19dfa042fff9986f72JSON `json:"-"`
}

// unnamedSchemaRef3248f24329456e19dfa042fff9986f72JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef3248f24329456e19dfa042fff9986f72]
type unnamedSchemaRef3248f24329456e19dfa042fff9986f72JSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef3248f24329456e19dfa042fff9986f72) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef3248f24329456e19dfa042fff9986f72JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8 struct {
	// Cloudflare account ID
	AccountTag  string      `json:"account_tag"`
	Connections interface{} `json:"connections,required"`
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
	// UUID of the tunnel.
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata,required"`
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
	TunType UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType `json:"tun_type"`
	JSON    unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8JSON    `json:"-"`
	union   UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8Union
}

// unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8]
type unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8JSON struct {
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	ID              apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8JSON) RawJSON() string {
	return r.raw
}

func (r *UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8) AsUnion() UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8Union {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by
// [shared.UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnel] or
// [shared.UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnel].
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8Union interface {
	implementsSharedUnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnel{}),
		},
	)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnection `json:"connections"`
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
	TunType UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType `json:"tun_type"`
	JSON    unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelJSON    `json:"-"`
}

// unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelJSON contains the
// JSON metadata for the struct
// [UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnel]
type unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelJSON struct {
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

func (r *UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelJSON) RawJSON() string {
	return r.raw
}

func (r UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnel) implementsSharedUnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8() {
}

type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnection struct {
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
	UUID string                                                                        `json:"uuid"`
	JSON unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnectionJSON `json:"-"`
}

// unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnectionJSON
// contains the JSON metadata for the struct
// [UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnection]
type unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType string

const (
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeCfdTunnel     UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType = "cfd_tunnel"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeWARPConnector UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType = "warp_connector"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeIPSec         UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType = "ip_sec"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeGRE           UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType = "gre"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeCni           UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType = "cni"
)

func (r UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunType) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeCfdTunnel, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeWARPConnector, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeIPSec, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeGRE, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelCfdTunnelTunTypeCni:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnection `json:"connections"`
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
	TunType UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelJSON    `json:"-"`
}

// unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelJSON
// contains the JSON metadata for the struct
// [UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnel]
type unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelJSON struct {
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

func (r *UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnel) implementsSharedUnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8() {
}

type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                                                  `json:"uuid"`
	JSON unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnection]
type unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType string

const (
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeCfdTunnel     UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeWARPConnector UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType = "warp_connector"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeIPSec         UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType = "ip_sec"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeGRE           UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType = "gre"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeCni           UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType = "cni"
)

func (r UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeCfdTunnel, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeWARPConnector, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeIPSec, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeGRE, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunnelWARPConnectorTunnelTunTypeCni:
		return true
	}
	return false
}

// The type of tunnel.
type UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType string

const (
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeCfdTunnel     UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType = "cfd_tunnel"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeWARPConnector UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType = "warp_connector"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeIPSec         UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType = "ip_sec"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeGRE           UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType = "gre"
	UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeCni           UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType = "cni"
)

func (r UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunType) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeCfdTunnel, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeWARPConnector, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeIPSec, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeGRE, UnnamedSchemaRef413ab4522f0bb93f63444799121fe2f8TunTypeCni:
		return true
	}
	return false
}

// Union satisfied by
// [shared.UnnamedSchemaRef602dd5f63eab958d53da61434dec08f0Unknown] or
// [shared.UnionString].
type UnnamedSchemaRef602dd5f63eab958d53da61434dec08f0Union interface {
	ImplementsSharedUnnamedSchemaRef602dd5f63eab958d53da61434dec08f0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef602dd5f63eab958d53da61434dec08f0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

// Union satisfied by
// [shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Unknown],
// [shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Array] or
// [shared.UnionString].
type UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union interface {
	ImplementsSharedUnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Array []interface{}

func (r UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Array) ImplementsSharedUnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union() {
}

// Union satisfied by
// [shared.UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Unknown],
// [shared.UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Array] or
// [shared.UnionString].
type UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Union interface {
	ImplementsSharedUnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Array []interface{}

func (r UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Array) ImplementsSharedUnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Union() {
}

// An object configuring the rule's logging behavior.
type UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                 `json:"enabled,required"`
	JSON    unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON `json:"-"`
}

// unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c]
type unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837 struct {
	// Identifier
	ID   string                                               `json:"id"`
	JSON unnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837JSON `json:"-"`
}

// unnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837]
type unnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837JSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [shared.UnnamedSchemaRef8d6a37a1e4190f86652802244d29525fUnknown] or
// [shared.UnionString].
type UnnamedSchemaRef8d6a37a1e4190f86652802244d29525fUnion interface {
	ImplementsSharedUnnamedSchemaRef8d6a37a1e4190f86652802244d29525fUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef8d6a37a1e4190f86652802244d29525fUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

// Union satisfied by
// [shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnknown] or
// [shared.UnionString].
type UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion interface {
	ImplementsSharedUnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854 struct {
	Valid bool                                                 `json:"valid"`
	JSON  unnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854JSON `json:"-"`
}

// unnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854]
type unnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854JSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefCc2ac1a037e5d6702fc77b3bcb527854JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691b string

const (
	UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691bStar     UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691b = "*"
	UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691bReferral UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691b = "referral"
	UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691bReferrer UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691b = "referrer"
)

func (r UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691b) IsKnown() bool {
	switch r {
	case UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691bStar, UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691bReferral, UnnamedSchemaRefD8600eb4758b3ae35607a0327bcd691bReferrer:
		return true
	}
	return false
}

// Union satisfied by
// [shared.UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Unknown],
// [shared.UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Array] or
// [shared.UnionString].
type UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Union interface {
	ImplementsSharedUnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Array []interface{}

func (r UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Array) ImplementsSharedUnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Union() {
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type UnnamedSchemaRefEe1e79edcb234d14c4dd266880f2fd24Param struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r UnnamedSchemaRefEe1e79edcb234d14c4dd266880f2fd24Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
