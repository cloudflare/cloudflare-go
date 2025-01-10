// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// TunnelConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelConfigurationService] method instead.
type TunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewTunnelConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelConfigurationService(opts ...option.RequestOption) (r *TunnelConfigurationService) {
	r = &TunnelConfigurationService{}
	r.Options = opts
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *TunnelConfigurationService) Update(ctx context.Context, tunnelID string, params TunnelConfigurationUpdateParams, opts ...option.RequestOption) (res *TunnelConfigurationUpdateResponse, err error) {
	var env TunnelConfigurationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *TunnelConfigurationService) Get(ctx context.Context, tunnelID string, query TunnelConfigurationGetParams, opts ...option.RequestOption) (res *TunnelConfigurationGetResponse, err error) {
	var env TunnelConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare Tunnel configuration
type TunnelConfigurationUpdateResponse struct {
	// Identifier
	AccountID string `json:"account_id"`
	// The tunnel configuration and ingress rules.
	Config    TunnelConfigurationUpdateResponseConfig `json:"config"`
	CreatedAt time.Time                               `json:"created_at" format:"date-time"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel's configuration on the Zero Trust dashboard.
	Source TunnelConfigurationUpdateResponseSource `json:"source"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// The version of the Tunnel Configuration.
	Version int64                                 `json:"version"`
	JSON    tunnelConfigurationUpdateResponseJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseJSON contains the JSON metadata for the struct
// [TunnelConfigurationUpdateResponse]
type tunnelConfigurationUpdateResponseJSON struct {
	AccountID   apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	Source      apijson.Field
	TunnelID    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The tunnel configuration and ingress rules.
type TunnelConfigurationUpdateResponseConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress []TunnelConfigurationUpdateResponseConfigIngress `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelConfigurationUpdateResponseConfigOriginRequest `json:"originRequest"`
	// Enable private network access from WARP users to private network routes. This is
	// enabled if the tunnel has an assigned route.
	WARPRouting TunnelConfigurationUpdateResponseConfigWARPRouting `json:"warp-routing"`
	JSON        tunnelConfigurationUpdateResponseConfigJSON        `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigJSON contains the JSON metadata for the
// struct [TunnelConfigurationUpdateResponseConfig]
type tunnelConfigurationUpdateResponseConfigJSON struct {
	Ingress       apijson.Field
	OriginRequest apijson.Field
	WARPRouting   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Public hostname
type TunnelConfigurationUpdateResponseConfigIngress struct {
	// Public hostname for this service.
	Hostname string `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service string `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelConfigurationUpdateResponseConfigIngressOriginRequest `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path string                                             `json:"path"`
	JSON tunnelConfigurationUpdateResponseConfigIngressJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigIngressJSON contains the JSON metadata
// for the struct [TunnelConfigurationUpdateResponseConfigIngress]
type tunnelConfigurationUpdateResponseConfigIngressJSON struct {
	Hostname      apijson.Field
	Service       apijson.Field
	OriginRequest apijson.Field
	Path          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfigIngress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigIngressJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationUpdateResponseConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelConfigurationUpdateResponseConfigIngressOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                           `json:"tlsTimeout"`
	JSON       tunnelConfigurationUpdateResponseConfigIngressOriginRequestJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigIngressOriginRequestJSON contains the
// JSON metadata for the struct
// [TunnelConfigurationUpdateResponseConfigIngressOriginRequest]
type tunnelConfigurationUpdateResponseConfigIngressOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfigIngressOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigIngressOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationUpdateResponseConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                                  `json:"required"`
	JSON     tunnelConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON contains
// the JSON metadata for the struct
// [TunnelConfigurationUpdateResponseConfigIngressOriginRequestAccess]
type tunnelConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfigIngressOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationUpdateResponseConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelConfigurationUpdateResponseConfigOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                    `json:"tlsTimeout"`
	JSON       tunnelConfigurationUpdateResponseConfigOriginRequestJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigOriginRequestJSON contains the JSON
// metadata for the struct [TunnelConfigurationUpdateResponseConfigOriginRequest]
type tunnelConfigurationUpdateResponseConfigOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfigOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationUpdateResponseConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                           `json:"required"`
	JSON     tunnelConfigurationUpdateResponseConfigOriginRequestAccessJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigOriginRequestAccessJSON contains the JSON
// metadata for the struct
// [TunnelConfigurationUpdateResponseConfigOriginRequestAccess]
type tunnelConfigurationUpdateResponseConfigOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfigOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelConfigurationUpdateResponseConfigWARPRouting struct {
	Enabled bool                                                   `json:"enabled"`
	JSON    tunnelConfigurationUpdateResponseConfigWARPRoutingJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseConfigWARPRoutingJSON contains the JSON
// metadata for the struct [TunnelConfigurationUpdateResponseConfigWARPRouting]
type tunnelConfigurationUpdateResponseConfigWARPRoutingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseConfigWARPRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseConfigWARPRoutingJSON) RawJSON() string {
	return r.raw
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel's configuration on the Zero Trust dashboard.
type TunnelConfigurationUpdateResponseSource string

const (
	TunnelConfigurationUpdateResponseSourceLocal      TunnelConfigurationUpdateResponseSource = "local"
	TunnelConfigurationUpdateResponseSourceCloudflare TunnelConfigurationUpdateResponseSource = "cloudflare"
)

func (r TunnelConfigurationUpdateResponseSource) IsKnown() bool {
	switch r {
	case TunnelConfigurationUpdateResponseSourceLocal, TunnelConfigurationUpdateResponseSourceCloudflare:
		return true
	}
	return false
}

// Cloudflare Tunnel configuration
type TunnelConfigurationGetResponse struct {
	// Identifier
	AccountID string `json:"account_id"`
	// The tunnel configuration and ingress rules.
	Config    TunnelConfigurationGetResponseConfig `json:"config"`
	CreatedAt time.Time                            `json:"created_at" format:"date-time"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel's configuration on the Zero Trust dashboard.
	Source TunnelConfigurationGetResponseSource `json:"source"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// The version of the Tunnel Configuration.
	Version int64                              `json:"version"`
	JSON    tunnelConfigurationGetResponseJSON `json:"-"`
}

// tunnelConfigurationGetResponseJSON contains the JSON metadata for the struct
// [TunnelConfigurationGetResponse]
type tunnelConfigurationGetResponseJSON struct {
	AccountID   apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	Source      apijson.Field
	TunnelID    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

// The tunnel configuration and ingress rules.
type TunnelConfigurationGetResponseConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress []TunnelConfigurationGetResponseConfigIngress `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelConfigurationGetResponseConfigOriginRequest `json:"originRequest"`
	// Enable private network access from WARP users to private network routes. This is
	// enabled if the tunnel has an assigned route.
	WARPRouting TunnelConfigurationGetResponseConfigWARPRouting `json:"warp-routing"`
	JSON        tunnelConfigurationGetResponseConfigJSON        `json:"-"`
}

// tunnelConfigurationGetResponseConfigJSON contains the JSON metadata for the
// struct [TunnelConfigurationGetResponseConfig]
type tunnelConfigurationGetResponseConfigJSON struct {
	Ingress       apijson.Field
	OriginRequest apijson.Field
	WARPRouting   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Public hostname
type TunnelConfigurationGetResponseConfigIngress struct {
	// Public hostname for this service.
	Hostname string `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service string `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelConfigurationGetResponseConfigIngressOriginRequest `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path string                                          `json:"path"`
	JSON tunnelConfigurationGetResponseConfigIngressJSON `json:"-"`
}

// tunnelConfigurationGetResponseConfigIngressJSON contains the JSON metadata for
// the struct [TunnelConfigurationGetResponseConfigIngress]
type tunnelConfigurationGetResponseConfigIngressJSON struct {
	Hostname      apijson.Field
	Service       apijson.Field
	OriginRequest apijson.Field
	Path          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfigIngress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigIngressJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationGetResponseConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelConfigurationGetResponseConfigIngressOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                        `json:"tlsTimeout"`
	JSON       tunnelConfigurationGetResponseConfigIngressOriginRequestJSON `json:"-"`
}

// tunnelConfigurationGetResponseConfigIngressOriginRequestJSON contains the JSON
// metadata for the struct
// [TunnelConfigurationGetResponseConfigIngressOriginRequest]
type tunnelConfigurationGetResponseConfigIngressOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfigIngressOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigIngressOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationGetResponseConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                               `json:"required"`
	JSON     tunnelConfigurationGetResponseConfigIngressOriginRequestAccessJSON `json:"-"`
}

// tunnelConfigurationGetResponseConfigIngressOriginRequestAccessJSON contains the
// JSON metadata for the struct
// [TunnelConfigurationGetResponseConfigIngressOriginRequestAccess]
type tunnelConfigurationGetResponseConfigIngressOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfigIngressOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigIngressOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationGetResponseConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelConfigurationGetResponseConfigOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                 `json:"tlsTimeout"`
	JSON       tunnelConfigurationGetResponseConfigOriginRequestJSON `json:"-"`
}

// tunnelConfigurationGetResponseConfigOriginRequestJSON contains the JSON metadata
// for the struct [TunnelConfigurationGetResponseConfigOriginRequest]
type tunnelConfigurationGetResponseConfigOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfigOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationGetResponseConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                        `json:"required"`
	JSON     tunnelConfigurationGetResponseConfigOriginRequestAccessJSON `json:"-"`
}

// tunnelConfigurationGetResponseConfigOriginRequestAccessJSON contains the JSON
// metadata for the struct
// [TunnelConfigurationGetResponseConfigOriginRequestAccess]
type tunnelConfigurationGetResponseConfigOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfigOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelConfigurationGetResponseConfigWARPRouting struct {
	Enabled bool                                                `json:"enabled"`
	JSON    tunnelConfigurationGetResponseConfigWARPRoutingJSON `json:"-"`
}

// tunnelConfigurationGetResponseConfigWARPRoutingJSON contains the JSON metadata
// for the struct [TunnelConfigurationGetResponseConfigWARPRouting]
type tunnelConfigurationGetResponseConfigWARPRoutingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseConfigWARPRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseConfigWARPRoutingJSON) RawJSON() string {
	return r.raw
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel's configuration on the Zero Trust dashboard.
type TunnelConfigurationGetResponseSource string

const (
	TunnelConfigurationGetResponseSourceLocal      TunnelConfigurationGetResponseSource = "local"
	TunnelConfigurationGetResponseSourceCloudflare TunnelConfigurationGetResponseSource = "cloudflare"
)

func (r TunnelConfigurationGetResponseSource) IsKnown() bool {
	switch r {
	case TunnelConfigurationGetResponseSourceLocal, TunnelConfigurationGetResponseSourceCloudflare:
		return true
	}
	return false
}

type TunnelConfigurationUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The tunnel configuration and ingress rules.
	Config param.Field[TunnelConfigurationUpdateParamsConfig] `json:"config"`
}

func (r TunnelConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type TunnelConfigurationUpdateParamsConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress param.Field[[]TunnelConfigurationUpdateParamsConfigIngress] `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[TunnelConfigurationUpdateParamsConfigOriginRequest] `json:"originRequest"`
}

func (r TunnelConfigurationUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type TunnelConfigurationUpdateParamsConfigIngress struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[TunnelConfigurationUpdateParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r TunnelConfigurationUpdateParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationUpdateParamsConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r TunnelConfigurationUpdateParamsConfigIngressOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationUpdateParamsConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[TunnelConfigurationUpdateParamsConfigOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r TunnelConfigurationUpdateParamsConfigOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationUpdateParamsConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelConfigurationUpdateParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelConfigurationUpdateParamsConfigWARPRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TunnelConfigurationUpdateParamsConfigWARPRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TunnelConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Cloudflare Tunnel configuration
	Result TunnelConfigurationUpdateResponse             `json:"result"`
	JSON   tunnelConfigurationUpdateResponseEnvelopeJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConfigurationUpdateResponseEnvelope]
type tunnelConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConfigurationUpdateResponseEnvelopeSuccess bool

const (
	TunnelConfigurationUpdateResponseEnvelopeSuccessTrue TunnelConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r TunnelConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelConfigurationGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TunnelConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	// Cloudflare Tunnel configuration
	Result TunnelConfigurationGetResponse             `json:"result"`
	JSON   tunnelConfigurationGetResponseEnvelopeJSON `json:"-"`
}

// tunnelConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConfigurationGetResponseEnvelope]
type tunnelConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConfigurationGetResponseEnvelopeSuccess bool

const (
	TunnelConfigurationGetResponseEnvelopeSuccessTrue TunnelConfigurationGetResponseEnvelopeSuccess = true
)

func (r TunnelConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
