// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/filters"
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

// An object configuring the rule's logging behavior.
type Logging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool        `json:"enabled,required"`
	JSON    loggingJSON `json:"-"`
}

// loggingJSON contains the JSON metadata for the struct [Logging]
type loggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Logging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loggingJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type LoggingParam struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r LoggingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseInfo struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    responseInfoJSON `json:"-"`
}

// responseInfoJSON contains the JSON metadata for the struct [ResponseInfo]
type responseInfoJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseInfoJSON) RawJSON() string {
	return r.raw
}

type ResponseInfoParam struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r ResponseInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRef10 struct {
	// URL that was excluded.
	URL  string                 `json:"url"`
	JSON unnamedSchemaRef10JSON `json:"-"`
}

// unnamedSchemaRef10JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef10]
type unnamedSchemaRef10JSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef10) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef10JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef102 struct {
	CreatedOn  interface{}             `json:"created_on"`
	Cron       interface{}             `json:"cron"`
	ModifiedOn interface{}             `json:"modified_on"`
	JSON       unnamedSchemaRef102JSON `json:"-"`
}

// unnamedSchemaRef102JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef102]
type unnamedSchemaRef102JSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef102) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef102JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [filters.FirewallFilter] or
// [shared.UnnamedSchemaRef107LegacyJhsDeletedFilter].
type UnnamedSchemaRef107 interface {
	implementsSharedUnnamedSchemaRef107()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef107)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(filters.FirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef107LegacyJhsDeletedFilter{}),
		},
	)
}

type UnnamedSchemaRef107LegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                          `json:"deleted,required"`
	JSON    unnamedSchemaRef107LegacyJhsDeletedFilterJSON `json:"-"`
}

// unnamedSchemaRef107LegacyJhsDeletedFilterJSON contains the JSON metadata for the
// struct [UnnamedSchemaRef107LegacyJhsDeletedFilter]
type unnamedSchemaRef107LegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef107LegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef107LegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r UnnamedSchemaRef107LegacyJhsDeletedFilter) implementsSharedUnnamedSchemaRef107() {}

type UnnamedSchemaRef11 struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                  `json:"url_id"`
	JSON  unnamedSchemaRef11JSON `json:"-"`
}

// unnamedSchemaRef11JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef11]
type unnamedSchemaRef11JSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef11) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef11JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnnamedSchemaRef116Unknown],
// [shared.UnnamedSchemaRef116Array] or [shared.UnionString].
type UnnamedSchemaRef116 interface {
	ImplementsSharedUnnamedSchemaRef116()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef116)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef116Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef116Array []interface{}

func (r UnnamedSchemaRef116Array) ImplementsSharedUnnamedSchemaRef116() {}

// Operating system
type UnnamedSchemaRef119 string

const (
	UnnamedSchemaRef119Windows UnnamedSchemaRef119 = "windows"
	UnnamedSchemaRef119Linux   UnnamedSchemaRef119 = "linux"
	UnnamedSchemaRef119Mac     UnnamedSchemaRef119 = "mac"
)

func (r UnnamedSchemaRef119) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef119Windows, UnnamedSchemaRef119Linux, UnnamedSchemaRef119Mac:
		return true
	}
	return false
}

type UnnamedSchemaRef12 struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                  `json:"url_id"`
	JSON  unnamedSchemaRef12JSON `json:"-"`
}

// unnamedSchemaRef12JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef12]
type unnamedSchemaRef12JSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef12) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef12JSON) RawJSON() string {
	return r.raw
}

// The rule group to which the current WAF rule belongs.
type UnnamedSchemaRef120 struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                  `json:"name"`
	JSON unnamedSchemaRef120JSON `json:"-"`
}

// unnamedSchemaRef120JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef120]
type unnamedSchemaRef120JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef120) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef120JSON) RawJSON() string {
	return r.raw
}

// Satisfied by [shared.UnionString], [shared.UnnamedSchemaRef121ArrayParam].
type UnnamedSchemaRef121Param interface {
	ImplementsSharedUnnamedSchemaRef121Param()
}

type UnnamedSchemaRef121ArrayParam []string

func (r UnnamedSchemaRef121ArrayParam) ImplementsSharedUnnamedSchemaRef121() {}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.UnnamedSchemaRef123TunnelCfdTunnel] or
// [shared.UnnamedSchemaRef123TunnelWARPConnectorTunnel].
type UnnamedSchemaRef123 interface {
	implementsSharedUnnamedSchemaRef123()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef123)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef123TunnelCfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef123TunnelWARPConnectorTunnel{}),
		},
	)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type UnnamedSchemaRef123TunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []UnnamedSchemaRef123TunnelCfdTunnelConnection `json:"connections"`
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
	TunType UnnamedSchemaRef123TunnelCfdTunnelTunType `json:"tun_type"`
	JSON    unnamedSchemaRef123TunnelCfdTunnelJSON    `json:"-"`
}

// unnamedSchemaRef123TunnelCfdTunnelJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef123TunnelCfdTunnel]
type unnamedSchemaRef123TunnelCfdTunnelJSON struct {
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

func (r *UnnamedSchemaRef123TunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef123TunnelCfdTunnelJSON) RawJSON() string {
	return r.raw
}

func (r UnnamedSchemaRef123TunnelCfdTunnel) implementsSharedUnnamedSchemaRef123() {}

type UnnamedSchemaRef123TunnelCfdTunnelConnection struct {
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
	UUID string                                           `json:"uuid"`
	JSON unnamedSchemaRef123TunnelCfdTunnelConnectionJSON `json:"-"`
}

// unnamedSchemaRef123TunnelCfdTunnelConnectionJSON contains the JSON metadata for
// the struct [UnnamedSchemaRef123TunnelCfdTunnelConnection]
type unnamedSchemaRef123TunnelCfdTunnelConnectionJSON struct {
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

func (r *UnnamedSchemaRef123TunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef123TunnelCfdTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type UnnamedSchemaRef123TunnelCfdTunnelTunType string

const (
	UnnamedSchemaRef123TunnelCfdTunnelTunTypeCfdTunnel     UnnamedSchemaRef123TunnelCfdTunnelTunType = "cfd_tunnel"
	UnnamedSchemaRef123TunnelCfdTunnelTunTypeWARPConnector UnnamedSchemaRef123TunnelCfdTunnelTunType = "warp_connector"
	UnnamedSchemaRef123TunnelCfdTunnelTunTypeIPSec         UnnamedSchemaRef123TunnelCfdTunnelTunType = "ip_sec"
	UnnamedSchemaRef123TunnelCfdTunnelTunTypeGRE           UnnamedSchemaRef123TunnelCfdTunnelTunType = "gre"
	UnnamedSchemaRef123TunnelCfdTunnelTunTypeCni           UnnamedSchemaRef123TunnelCfdTunnelTunType = "cni"
)

func (r UnnamedSchemaRef123TunnelCfdTunnelTunType) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef123TunnelCfdTunnelTunTypeCfdTunnel, UnnamedSchemaRef123TunnelCfdTunnelTunTypeWARPConnector, UnnamedSchemaRef123TunnelCfdTunnelTunTypeIPSec, UnnamedSchemaRef123TunnelCfdTunnelTunTypeGRE, UnnamedSchemaRef123TunnelCfdTunnelTunTypeCni:
		return true
	}
	return false
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type UnnamedSchemaRef123TunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []UnnamedSchemaRef123TunnelWARPConnectorTunnelConnection `json:"connections"`
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
	TunType UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    unnamedSchemaRef123TunnelWARPConnectorTunnelJSON    `json:"-"`
}

// unnamedSchemaRef123TunnelWARPConnectorTunnelJSON contains the JSON metadata for
// the struct [UnnamedSchemaRef123TunnelWARPConnectorTunnel]
type unnamedSchemaRef123TunnelWARPConnectorTunnelJSON struct {
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

func (r *UnnamedSchemaRef123TunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef123TunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r UnnamedSchemaRef123TunnelWARPConnectorTunnel) implementsSharedUnnamedSchemaRef123() {}

type UnnamedSchemaRef123TunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                     `json:"uuid"`
	JSON unnamedSchemaRef123TunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// unnamedSchemaRef123TunnelWARPConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct [UnnamedSchemaRef123TunnelWARPConnectorTunnelConnection]
type unnamedSchemaRef123TunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *UnnamedSchemaRef123TunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef123TunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The type of tunnel.
type UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType string

const (
	UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeCfdTunnel     UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeWARPConnector UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType = "warp_connector"
	UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeIPSec         UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType = "ip_sec"
	UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeGRE           UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType = "gre"
	UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeCni           UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType = "cni"
)

func (r UnnamedSchemaRef123TunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeCfdTunnel, UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeWARPConnector, UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeIPSec, UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeGRE, UnnamedSchemaRef123TunnelWARPConnectorTunnelTunTypeCni:
		return true
	}
	return false
}

// account settings.
type UnnamedSchemaRef125 struct {
	// Activity log settings.
	ActivityLog UnnamedSchemaRef125ActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus UnnamedSchemaRef125Antivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage UnnamedSchemaRef125BlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning UnnamedSchemaRef125BodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation UnnamedSchemaRef125BrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate UnnamedSchemaRef125CustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching UnnamedSchemaRef125ExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips UnnamedSchemaRef125Fips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection UnnamedSchemaRef125ProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt UnnamedSchemaRef125TLSDecrypt `json:"tls_decrypt"`
	JSON       unnamedSchemaRef125JSON       `json:"-"`
}

// unnamedSchemaRef125JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125]
type unnamedSchemaRef125JSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	ProtocolDetection     apijson.Field
	TLSDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *UnnamedSchemaRef125) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125JSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type UnnamedSchemaRef125ActivityLog struct {
	// Enable activity logging.
	Enabled bool                               `json:"enabled"`
	JSON    unnamedSchemaRef125ActivityLogJSON `json:"-"`
}

// unnamedSchemaRef125ActivityLogJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125ActivityLog]
type unnamedSchemaRef125ActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef125ActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125ActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type UnnamedSchemaRef125Antivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings UnnamedSchemaRef125AntivirusNotificationSettings `json:"notification_settings"`
	JSON                 unnamedSchemaRef125AntivirusJSON                 `json:"-"`
}

// unnamedSchemaRef125AntivirusJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125Antivirus]
type unnamedSchemaRef125AntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UnnamedSchemaRef125Antivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125AntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type UnnamedSchemaRef125AntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                               `json:"support_url"`
	JSON       unnamedSchemaRef125AntivirusNotificationSettingsJSON `json:"-"`
}

// unnamedSchemaRef125AntivirusNotificationSettingsJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef125AntivirusNotificationSettings]
type unnamedSchemaRef125AntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef125AntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125AntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type UnnamedSchemaRef125BlockPage struct {
	// Block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled"`
	// Block page footer text.
	FooterText string `json:"footer_text"`
	// Block page header text.
	HeaderText string `json:"header_text"`
	// Full URL to the logo file.
	LogoPath string `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress string `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject string `json:"mailto_subject"`
	// Block page title.
	Name string `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter bool                             `json:"suppress_footer"`
	JSON           unnamedSchemaRef125BlockPageJSON `json:"-"`
}

// unnamedSchemaRef125BlockPageJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125BlockPage]
type unnamedSchemaRef125BlockPageJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Name            apijson.Field
	SuppressFooter  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UnnamedSchemaRef125BlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125BlockPageJSON) RawJSON() string {
	return r.raw
}

// DLP body scanning settings.
type UnnamedSchemaRef125BodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                              `json:"inspection_mode"`
	JSON           unnamedSchemaRef125BodyScanningJSON `json:"-"`
}

// unnamedSchemaRef125BodyScanningJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125BodyScanning]
type unnamedSchemaRef125BodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UnnamedSchemaRef125BodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125BodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type UnnamedSchemaRef125BrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                    `json:"url_browser_isolation_enabled"`
	JSON                       unnamedSchemaRef125BrowserIsolationJSON `json:"-"`
}

// unnamedSchemaRef125BrowserIsolationJSON contains the JSON metadata for the
// struct [UnnamedSchemaRef125BrowserIsolation]
type unnamedSchemaRef125BrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UnnamedSchemaRef125BrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125BrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI.
type UnnamedSchemaRef125CustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                   `json:"binding_status"`
	UpdatedAt     time.Time                                `json:"updated_at" format:"date-time"`
	JSON          unnamedSchemaRef125CustomCertificateJSON `json:"-"`
}

// unnamedSchemaRef125CustomCertificateJSON contains the JSON metadata for the
// struct [UnnamedSchemaRef125CustomCertificate]
type unnamedSchemaRef125CustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UnnamedSchemaRef125CustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125CustomCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type UnnamedSchemaRef125ExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                         `json:"enabled"`
	JSON    unnamedSchemaRef125ExtendedEmailMatchingJSON `json:"-"`
}

// unnamedSchemaRef125ExtendedEmailMatchingJSON contains the JSON metadata for the
// struct [UnnamedSchemaRef125ExtendedEmailMatching]
type unnamedSchemaRef125ExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef125ExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125ExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type UnnamedSchemaRef125Fips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                        `json:"tls"`
	JSON unnamedSchemaRef125FipsJSON `json:"-"`
}

// unnamedSchemaRef125FipsJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125Fips]
type unnamedSchemaRef125FipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef125Fips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125FipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type UnnamedSchemaRef125ProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                     `json:"enabled"`
	JSON    unnamedSchemaRef125ProtocolDetectionJSON `json:"-"`
}

// unnamedSchemaRef125ProtocolDetectionJSON contains the JSON metadata for the
// struct [UnnamedSchemaRef125ProtocolDetection]
type unnamedSchemaRef125ProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef125ProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125ProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type UnnamedSchemaRef125TLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                              `json:"enabled"`
	JSON    unnamedSchemaRef125TLSDecryptJSON `json:"-"`
}

// unnamedSchemaRef125TLSDecryptJSON contains the JSON metadata for the struct
// [UnnamedSchemaRef125TLSDecrypt]
type unnamedSchemaRef125TLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef125TLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef125TLSDecryptJSON) RawJSON() string {
	return r.raw
}

// account settings.
type UnnamedSchemaRef125Param struct {
	// Activity log settings.
	ActivityLog param.Field[UnnamedSchemaRef125ActivityLogParam] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[UnnamedSchemaRef125AntivirusParam] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[UnnamedSchemaRef125BlockPageParam] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[UnnamedSchemaRef125BodyScanningParam] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[UnnamedSchemaRef125BrowserIsolationParam] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[UnnamedSchemaRef125CustomCertificateParam] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[UnnamedSchemaRef125ExtendedEmailMatchingParam] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[UnnamedSchemaRef125FipsParam] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[UnnamedSchemaRef125ProtocolDetectionParam] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[UnnamedSchemaRef125TLSDecryptParam] `json:"tls_decrypt"`
}

func (r UnnamedSchemaRef125Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type UnnamedSchemaRef125ActivityLogParam struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef125ActivityLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type UnnamedSchemaRef125AntivirusParam struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[UnnamedSchemaRef125AntivirusNotificationSettingsParam] `json:"notification_settings"`
}

func (r UnnamedSchemaRef125AntivirusParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type UnnamedSchemaRef125AntivirusNotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r UnnamedSchemaRef125AntivirusNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type UnnamedSchemaRef125BlockPageParam struct {
	// Block page background color in #rrggbb format.
	BackgroundColor param.Field[string] `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled param.Field[bool] `json:"enabled"`
	// Block page footer text.
	FooterText param.Field[string] `json:"footer_text"`
	// Block page header text.
	HeaderText param.Field[string] `json:"header_text"`
	// Full URL to the logo file.
	LogoPath param.Field[string] `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress param.Field[string] `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject param.Field[string] `json:"mailto_subject"`
	// Block page title.
	Name param.Field[string] `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter param.Field[bool] `json:"suppress_footer"`
}

func (r UnnamedSchemaRef125BlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type UnnamedSchemaRef125BodyScanningParam struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r UnnamedSchemaRef125BodyScanningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type UnnamedSchemaRef125BrowserIsolationParam struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r UnnamedSchemaRef125BrowserIsolationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type UnnamedSchemaRef125CustomCertificateParam struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r UnnamedSchemaRef125CustomCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type UnnamedSchemaRef125ExtendedEmailMatchingParam struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef125ExtendedEmailMatchingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type UnnamedSchemaRef125FipsParam struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r UnnamedSchemaRef125FipsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type UnnamedSchemaRef125ProtocolDetectionParam struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef125ProtocolDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type UnnamedSchemaRef125TLSDecryptParam struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef125TLSDecryptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRef13 struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                 `json:"verification_status"`
	JSON               unnamedSchemaRef13JSON `json:"-"`
}

// unnamedSchemaRef13JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef13]
type unnamedSchemaRef13JSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UnnamedSchemaRef13) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef13JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef131 struct {
	// The ID of a previously created Access group.
	ID   string                  `json:"id,required"`
	JSON unnamedSchemaRef131JSON `json:"-"`
}

// unnamedSchemaRef131JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef131]
type unnamedSchemaRef131JSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef131) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef131JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef131Param struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r UnnamedSchemaRef131Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRef135 struct {
	Label string                  `json:"label"`
	Score float64                 `json:"score"`
	JSON  unnamedSchemaRef135JSON `json:"-"`
}

// unnamedSchemaRef135JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef135]
type unnamedSchemaRef135JSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef135) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef135JSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type UnnamedSchemaRef139 string

const (
	UnnamedSchemaRef139CacheReserveClear UnnamedSchemaRef139 = "cache_reserve_clear"
)

func (r UnnamedSchemaRef139) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef139CacheReserveClear:
		return true
	}
	return false
}

type UnnamedSchemaRef14 struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                `json:"model_score"`
	JSON       unnamedSchemaRef14JSON `json:"-"`
}

// unnamedSchemaRef14JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef14]
type unnamedSchemaRef14JSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef14) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef14JSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type UnnamedSchemaRef140 string

const (
	UnnamedSchemaRef140TcRegional UnnamedSchemaRef140 = "tc_regional"
)

func (r UnnamedSchemaRef140) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef140TcRegional:
		return true
	}
	return false
}

// Value of the zone setting.
type UnnamedSchemaRef142 string

const (
	UnnamedSchemaRef142OriginMaxHTTPVersion UnnamedSchemaRef142 = "origin_max_http_version"
)

func (r UnnamedSchemaRef142) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef142OriginMaxHTTPVersion:
		return true
	}
	return false
}

// ID of the zone setting.
type UnnamedSchemaRef144 string

const (
	UnnamedSchemaRef144CacheReserve UnnamedSchemaRef144 = "cache_reserve"
)

func (r UnnamedSchemaRef144) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef144CacheReserve:
		return true
	}
	return false
}

type UnnamedSchemaRef146 string

const (
	UnnamedSchemaRef146Star     UnnamedSchemaRef146 = "*"
	UnnamedSchemaRef146Referral UnnamedSchemaRef146 = "referral"
	UnnamedSchemaRef146Referrer UnnamedSchemaRef146 = "referrer"
)

func (r UnnamedSchemaRef146) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef146Star, UnnamedSchemaRef146Referral, UnnamedSchemaRef146Referrer:
		return true
	}
	return false
}

type UnnamedSchemaRef148 struct {
	// List of metrics returned by the query.
	Metrics []interface{}           `json:"metrics,required"`
	JSON    unnamedSchemaRef148JSON `json:"-"`
}

// unnamedSchemaRef148JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef148]
type unnamedSchemaRef148JSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef148) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef148JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef15 struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                 `json:"name"`
	JSON unnamedSchemaRef15JSON `json:"-"`
}

// unnamedSchemaRef15JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef15]
type unnamedSchemaRef15JSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef15) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef15JSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type UnnamedSchemaRef151 string

const (
	UnnamedSchemaRef151Variants UnnamedSchemaRef151 = "variants"
)

func (r UnnamedSchemaRef151) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef151Variants:
		return true
	}
	return false
}

type UnnamedSchemaRef155 struct {
	Name  string                  `json:"name,required"`
	Value string                  `json:"value,required"`
	JSON  unnamedSchemaRef155JSON `json:"-"`
}

// unnamedSchemaRef155JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef155]
type unnamedSchemaRef155JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef155) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef155JSON) RawJSON() string {
	return r.raw
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type UnnamedSchemaRef158 struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                    `json:"user_deprovision"`
	JSON            unnamedSchemaRef158JSON `json:"-"`
}

// unnamedSchemaRef158JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef158]
type unnamedSchemaRef158JSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *UnnamedSchemaRef158) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef158JSON) RawJSON() string {
	return r.raw
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type UnnamedSchemaRef158Param struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r UnnamedSchemaRef158Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the most recent scan found.
type UnnamedSchemaRef16 struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                  `json:"submission_id"`
	JSON         unnamedSchemaRef16JSON `json:"-"`
}

// unnamedSchemaRef16JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef16]
type unnamedSchemaRef16JSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UnnamedSchemaRef16) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef16JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnnamedSchemaRef161Unknown],
// [shared.UnnamedSchemaRef161Array] or [shared.UnionString].
type UnnamedSchemaRef161 interface {
	ImplementsSharedUnnamedSchemaRef161()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef161)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef161Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef161Array []interface{}

func (r UnnamedSchemaRef161Array) ImplementsSharedUnnamedSchemaRef161() {}

// Extra Cloudflare-specific information about the record.
type UnnamedSchemaRef162 struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                  `json:"source"`
	JSON   unnamedSchemaRef162JSON `json:"-"`
}

// unnamedSchemaRef162JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef162]
type unnamedSchemaRef162JSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef162) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef162JSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type UnnamedSchemaRef163 string

const (
	UnnamedSchemaRef163Onetimepin UnnamedSchemaRef163 = "onetimepin"
	UnnamedSchemaRef163AzureAd    UnnamedSchemaRef163 = "azureAD"
	UnnamedSchemaRef163Saml       UnnamedSchemaRef163 = "saml"
	UnnamedSchemaRef163Centrify   UnnamedSchemaRef163 = "centrify"
	UnnamedSchemaRef163Facebook   UnnamedSchemaRef163 = "facebook"
	UnnamedSchemaRef163GitHub     UnnamedSchemaRef163 = "github"
	UnnamedSchemaRef163GoogleApps UnnamedSchemaRef163 = "google-apps"
	UnnamedSchemaRef163Google     UnnamedSchemaRef163 = "google"
	UnnamedSchemaRef163Linkedin   UnnamedSchemaRef163 = "linkedin"
	UnnamedSchemaRef163Oidc       UnnamedSchemaRef163 = "oidc"
	UnnamedSchemaRef163Okta       UnnamedSchemaRef163 = "okta"
	UnnamedSchemaRef163Onelogin   UnnamedSchemaRef163 = "onelogin"
	UnnamedSchemaRef163Pingone    UnnamedSchemaRef163 = "pingone"
	UnnamedSchemaRef163Yandex     UnnamedSchemaRef163 = "yandex"
)

func (r UnnamedSchemaRef163) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef163Onetimepin, UnnamedSchemaRef163AzureAd, UnnamedSchemaRef163Saml, UnnamedSchemaRef163Centrify, UnnamedSchemaRef163Facebook, UnnamedSchemaRef163GitHub, UnnamedSchemaRef163GoogleApps, UnnamedSchemaRef163Google, UnnamedSchemaRef163Linkedin, UnnamedSchemaRef163Oidc, UnnamedSchemaRef163Okta, UnnamedSchemaRef163Onelogin, UnnamedSchemaRef163Pingone, UnnamedSchemaRef163Yandex:
		return true
	}
	return false
}

// Union satisfied by [shared.UnnamedSchemaRef167Unknown],
// [shared.UnnamedSchemaRef167Array] or [shared.UnionString].
type UnnamedSchemaRef167 interface {
	ImplementsSharedUnnamedSchemaRef167()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef167)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef167Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef167Array []interface{}

func (r UnnamedSchemaRef167Array) ImplementsSharedUnnamedSchemaRef167() {}

// Union satisfied by [shared.UnnamedSchemaRef169Unknown] or [shared.UnionString].
type UnnamedSchemaRef169 interface {
	ImplementsSharedUnnamedSchemaRef169()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef169)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

// Union satisfied by [shared.UnnamedSchemaRef171Unknown],
// [shared.UnnamedSchemaRef171Array] or [shared.UnionString].
type UnnamedSchemaRef171 interface {
	ImplementsSharedUnnamedSchemaRef171()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef171)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef171Array{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef171Array []interface{}

func (r UnnamedSchemaRef171Array) ImplementsSharedUnnamedSchemaRef171() {}

// Union satisfied by [shared.UnnamedSchemaRef173Unknown] or [shared.UnionString].
type UnnamedSchemaRef173 interface {
	ImplementsSharedUnnamedSchemaRef173()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef173)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

type UnnamedSchemaRef22 struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64              `json:"metrics,required"`
	JSON    unnamedSchemaRef22JSON `json:"-"`
}

// unnamedSchemaRef22JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef22]
type unnamedSchemaRef22JSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef22) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef22JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef23 []interface{}

type UnnamedSchemaRef24 struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta UnnamedSchemaRef24TimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string               `json:"sort"`
	JSON unnamedSchemaRef24JSON `json:"-"`
}

// unnamedSchemaRef24JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef24]
type unnamedSchemaRef24JSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef24) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef24JSON) RawJSON() string {
	return r.raw
}

// Unit of time to group data by.
type UnnamedSchemaRef24TimeDelta string

const (
	UnnamedSchemaRef24TimeDeltaAll        UnnamedSchemaRef24TimeDelta = "all"
	UnnamedSchemaRef24TimeDeltaAuto       UnnamedSchemaRef24TimeDelta = "auto"
	UnnamedSchemaRef24TimeDeltaYear       UnnamedSchemaRef24TimeDelta = "year"
	UnnamedSchemaRef24TimeDeltaQuarter    UnnamedSchemaRef24TimeDelta = "quarter"
	UnnamedSchemaRef24TimeDeltaMonth      UnnamedSchemaRef24TimeDelta = "month"
	UnnamedSchemaRef24TimeDeltaWeek       UnnamedSchemaRef24TimeDelta = "week"
	UnnamedSchemaRef24TimeDeltaDay        UnnamedSchemaRef24TimeDelta = "day"
	UnnamedSchemaRef24TimeDeltaHour       UnnamedSchemaRef24TimeDelta = "hour"
	UnnamedSchemaRef24TimeDeltaDekaminute UnnamedSchemaRef24TimeDelta = "dekaminute"
	UnnamedSchemaRef24TimeDeltaMinute     UnnamedSchemaRef24TimeDelta = "minute"
)

func (r UnnamedSchemaRef24TimeDelta) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef24TimeDeltaAll, UnnamedSchemaRef24TimeDeltaAuto, UnnamedSchemaRef24TimeDeltaYear, UnnamedSchemaRef24TimeDeltaQuarter, UnnamedSchemaRef24TimeDeltaMonth, UnnamedSchemaRef24TimeDeltaWeek, UnnamedSchemaRef24TimeDeltaDay, UnnamedSchemaRef24TimeDeltaHour, UnnamedSchemaRef24TimeDeltaDekaminute, UnnamedSchemaRef24TimeDeltaMinute:
		return true
	}
	return false
}

// Logging settings by rule type.
type UnnamedSchemaRef28 struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}            `json:"l4"`
	JSON unnamedSchemaRef28JSON `json:"-"`
}

// unnamedSchemaRef28JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef28]
type unnamedSchemaRef28JSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef28) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef28JSON) RawJSON() string {
	return r.raw
}

// Logging settings by rule type.
type UnnamedSchemaRef28Param struct {
	// Logging settings for DNS firewall.
	DNS param.Field[interface{}] `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP param.Field[interface{}] `json:"http"`
	// Logging settings for Network firewall.
	L4 param.Field[interface{}] `json:"l4"`
}

func (r UnnamedSchemaRef28Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRef3 struct {
	// Subscription identifier tag.
	SubscriptionID string                `json:"subscription_id"`
	JSON           unnamedSchemaRef3JSON `json:"-"`
}

// unnamedSchemaRef3JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef3]
type unnamedSchemaRef3JSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UnnamedSchemaRef3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef3JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef35 struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname interface{} `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time              `json:"last_seen" format:"date"`
	JSON     unnamedSchemaRef35JSON `json:"-"`
}

// unnamedSchemaRef35JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef35]
type unnamedSchemaRef35JSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef35) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef35JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef41 struct {
	// The maximum number of messages to include in a batch
	BatchSize     float64                `json:"batch_size"`
	MaxRetries    float64                `json:"max_retries"`
	MaxWaitTimeMs float64                `json:"max_wait_time_ms"`
	JSON          unnamedSchemaRef41JSON `json:"-"`
}

// unnamedSchemaRef41JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef41]
type unnamedSchemaRef41JSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UnnamedSchemaRef41) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef41JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef44 struct {
	After  string                 `json:"after"`
	Before string                 `json:"before"`
	JSON   unnamedSchemaRef44JSON `json:"-"`
}

// unnamedSchemaRef44JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef44]
type unnamedSchemaRef44JSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef44) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef44JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnnamedSchemaRef49Unknown] or [shared.UnionString].
type UnnamedSchemaRef49 interface {
	ImplementsSharedUnnamedSchemaRef49()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef49)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type UnnamedSchemaRef51Param struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r UnnamedSchemaRef51Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRef62 struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time              `json:"start_time" format:"date-time"`
	JSON      unnamedSchemaRef62JSON `json:"-"`
}

// unnamedSchemaRef62JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef62]
type unnamedSchemaRef62JSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef62) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef62JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef65 struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value UnnamedSchemaRef65Value `json:"value"`
	JSON  unnamedSchemaRef65JSON  `json:"-"`
}

// unnamedSchemaRef65JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef65]
type unnamedSchemaRef65JSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef65) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef65JSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [shared.UnnamedSchemaRef65ValueArray].
type UnnamedSchemaRef65Value interface {
	ImplementsSharedUnnamedSchemaRef65Value()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef65Value)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UnnamedSchemaRef65ValueArray{}),
		},
	)
}

type UnnamedSchemaRef65ValueArray []string

func (r UnnamedSchemaRef65ValueArray) ImplementsSharedUnnamedSchemaRef65Value() {}

// A globally unique name for an identity or service provider.
type UnnamedSchemaRef75 string

const (
	UnnamedSchemaRef75UrnOasisNamesTcSaml2_0AttrnameFormatUnspecified UnnamedSchemaRef75 = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	UnnamedSchemaRef75UrnOasisNamesTcSaml2_0AttrnameFormatBasic       UnnamedSchemaRef75 = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	UnnamedSchemaRef75UrnOasisNamesTcSaml2_0AttrnameFormatURI         UnnamedSchemaRef75 = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r UnnamedSchemaRef75) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef75UrnOasisNamesTcSaml2_0AttrnameFormatUnspecified, UnnamedSchemaRef75UrnOasisNamesTcSaml2_0AttrnameFormatBasic, UnnamedSchemaRef75UrnOasisNamesTcSaml2_0AttrnameFormatURI:
		return true
	}
	return false
}

type UnnamedSchemaRef76 struct {
	// The name of the IdP attribute.
	Name string                 `json:"name"`
	JSON unnamedSchemaRef76JSON `json:"-"`
}

// unnamedSchemaRef76JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef76]
type unnamedSchemaRef76JSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef76) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef76JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef76Param struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r UnnamedSchemaRef76Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type UnnamedSchemaRef77 string

const (
	UnnamedSchemaRef77ID    UnnamedSchemaRef77 = "id"
	UnnamedSchemaRef77Email UnnamedSchemaRef77 = "email"
)

func (r UnnamedSchemaRef77) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef77ID, UnnamedSchemaRef77Email:
		return true
	}
	return false
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type UnnamedSchemaRef78 string

const (
	UnnamedSchemaRef78Ubiquitous UnnamedSchemaRef78 = "ubiquitous"
	UnnamedSchemaRef78Optimal    UnnamedSchemaRef78 = "optimal"
	UnnamedSchemaRef78Force      UnnamedSchemaRef78 = "force"
)

func (r UnnamedSchemaRef78) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef78Ubiquitous, UnnamedSchemaRef78Optimal, UnnamedSchemaRef78Force:
		return true
	}
	return false
}

// Domain control validation (DCV) method used for this hostname.
type UnnamedSchemaRef79 string

const (
	UnnamedSchemaRef79HTTP  UnnamedSchemaRef79 = "http"
	UnnamedSchemaRef79TXT   UnnamedSchemaRef79 = "txt"
	UnnamedSchemaRef79Email UnnamedSchemaRef79 = "email"
)

func (r UnnamedSchemaRef79) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef79HTTP, UnnamedSchemaRef79TXT, UnnamedSchemaRef79Email:
		return true
	}
	return false
}

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type UnnamedSchemaRef80 string

const (
	UnnamedSchemaRef80Dv UnnamedSchemaRef80 = "dv"
)

func (r UnnamedSchemaRef80) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef80Dv:
		return true
	}
	return false
}

// Whether or not Early Hints is enabled.
type UnnamedSchemaRef81 string

const (
	UnnamedSchemaRef81On  UnnamedSchemaRef81 = "on"
	UnnamedSchemaRef81Off UnnamedSchemaRef81 = "off"
)

func (r UnnamedSchemaRef81) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef81On, UnnamedSchemaRef81Off:
		return true
	}
	return false
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type UnnamedSchemaRef82 string

const (
	UnnamedSchemaRef82Tcp  UnnamedSchemaRef82 = "tcp"
	UnnamedSchemaRef82Udp  UnnamedSchemaRef82 = "udp"
	UnnamedSchemaRef82Icmp UnnamedSchemaRef82 = "icmp"
)

func (r UnnamedSchemaRef82) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef82Tcp, UnnamedSchemaRef82Udp, UnnamedSchemaRef82Icmp:
		return true
	}
	return false
}

// How frequent the health check is run. The default value is `mid`.
type UnnamedSchemaRef83 string

const (
	UnnamedSchemaRef83Low  UnnamedSchemaRef83 = "low"
	UnnamedSchemaRef83Mid  UnnamedSchemaRef83 = "mid"
	UnnamedSchemaRef83High UnnamedSchemaRef83 = "high"
)

func (r UnnamedSchemaRef83) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef83Low, UnnamedSchemaRef83Mid, UnnamedSchemaRef83High:
		return true
	}
	return false
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type UnnamedSchemaRef84 string

const (
	UnnamedSchemaRef84Reply   UnnamedSchemaRef84 = "reply"
	UnnamedSchemaRef84Request UnnamedSchemaRef84 = "request"
)

func (r UnnamedSchemaRef84) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef84Reply, UnnamedSchemaRef84Request:
		return true
	}
	return false
}

type UnnamedSchemaRef85 string

const (
	UnnamedSchemaRef85R2 UnnamedSchemaRef85 = "r2"
)

func (r UnnamedSchemaRef85) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef85R2:
		return true
	}
	return false
}

// Operator
type UnnamedSchemaRef87 string

const (
	UnnamedSchemaRef87Less            UnnamedSchemaRef87 = "<"
	UnnamedSchemaRef87LessOrEquals    UnnamedSchemaRef87 = "<="
	UnnamedSchemaRef87Greater         UnnamedSchemaRef87 = ">"
	UnnamedSchemaRef87GreaterOrEquals UnnamedSchemaRef87 = ">="
	UnnamedSchemaRef87Equals          UnnamedSchemaRef87 = "=="
)

func (r UnnamedSchemaRef87) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef87Less, UnnamedSchemaRef87LessOrEquals, UnnamedSchemaRef87Greater, UnnamedSchemaRef87GreaterOrEquals, UnnamedSchemaRef87Equals:
		return true
	}
	return false
}

// Version Operator
type UnnamedSchemaRef88 string

const (
	UnnamedSchemaRef88Less            UnnamedSchemaRef88 = "<"
	UnnamedSchemaRef88LessOrEquals    UnnamedSchemaRef88 = "<="
	UnnamedSchemaRef88Greater         UnnamedSchemaRef88 = ">"
	UnnamedSchemaRef88GreaterOrEquals UnnamedSchemaRef88 = ">="
	UnnamedSchemaRef88Equals          UnnamedSchemaRef88 = "=="
)

func (r UnnamedSchemaRef88) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef88Less, UnnamedSchemaRef88LessOrEquals, UnnamedSchemaRef88Greater, UnnamedSchemaRef88GreaterOrEquals, UnnamedSchemaRef88Equals:
		return true
	}
	return false
}

// Automatically minify all CSS files for your website.
type UnnamedSchemaRef92 string

const (
	UnnamedSchemaRef92On  UnnamedSchemaRef92 = "on"
	UnnamedSchemaRef92Off UnnamedSchemaRef92 = "off"
)

func (r UnnamedSchemaRef92) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef92On, UnnamedSchemaRef92Off:
		return true
	}
	return false
}

type UnnamedSchemaRef94 struct {
	End   float64                `json:"end"`
	Start float64                `json:"start"`
	Word  string                 `json:"word"`
	JSON  unnamedSchemaRef94JSON `json:"-"`
}

// unnamedSchemaRef94JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef94]
type unnamedSchemaRef94JSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef94) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef94JSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type UnnamedSchemaRef96 string

const (
	UnnamedSchemaRef96Custom UnnamedSchemaRef96 = "custom"
)

func (r UnnamedSchemaRef96) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef96Custom:
		return true
	}
	return false
}

// The type of the profile.
type UnnamedSchemaRef97 string

const (
	UnnamedSchemaRef97Predefined UnnamedSchemaRef97 = "predefined"
)

func (r UnnamedSchemaRef97) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef97Predefined:
		return true
	}
	return false
}
