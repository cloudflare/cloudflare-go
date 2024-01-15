// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCfdTunnelConfigurationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountCfdTunnelConfigurationService] method instead.
type AccountCfdTunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelConfigurationService(opts ...option.RequestOption) (r *AccountCfdTunnelConfigurationService) {
	r = &AccountCfdTunnelConfigurationService{}
	r.Options = opts
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *AccountCfdTunnelConfigurationService) CloudflareTunnelConfigurationGetConfiguration(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *AccountCfdTunnelConfigurationService) CloudflareTunnelConfigurationPutConfiguration(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams, opts ...option.RequestOption) (res *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse struct {
	Errors   []AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseError   `json:"errors"`
	Messages []AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessage `json:"messages"`
	Result   interface{}                                                                                 `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseSuccess `json:"success"`
	JSON    accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseJSON    `json:"-"`
}

// accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse]
type accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseErrorJSON `json:"-"`
}

// accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseError]
type accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessageJSON `json:"-"`
}

// accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessage]
type accountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseSuccess bool

const (
	AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseSuccessTrue AccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseSuccess = true
)

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse struct {
	Errors   []AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseError   `json:"errors"`
	Messages []AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessage `json:"messages"`
	Result   interface{}                                                                                 `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseSuccess `json:"success"`
	JSON    accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseJSON    `json:"-"`
}

// accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse]
type accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseErrorJSON `json:"-"`
}

// accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseError]
type accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessageJSON `json:"-"`
}

// accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessage]
type accountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseSuccess bool

const (
	AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseSuccessTrue AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseSuccess = true
)

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams struct {
	// The tunnel configuration and ingress rules.
	Config param.Field[AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig] `json:"config"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig struct {
	// List of public hostname definitions
	Ingress param.Field[[]AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress] `json:"ingress"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest] `json:"originRequest"`
	// Enable private network access from WARP users to private network routes
	WarpRouting param.Field[AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting] `json:"warp-routing"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CaPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	Http2Origin param.Field[bool] `json:"http2Origin"`
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
	NoTlsVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TlsTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CaPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	Http2Origin param.Field[bool] `json:"http2Origin"`
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
	NoTlsVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TlsTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes
type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
