// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CfdTunnelConfigurationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCfdTunnelConfigurationService]
// method instead.
type CfdTunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewCfdTunnelConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCfdTunnelConfigurationService(opts ...option.RequestOption) (r *CfdTunnelConfigurationService) {
	r = &CfdTunnelConfigurationService{}
	r.Options = opts
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *CfdTunnelConfigurationService) CloudflareTunnelConfigurationGetConfiguration(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *CfdTunnelConfigurationService) CloudflareTunnelConfigurationPutConfiguration(ctx context.Context, accountID string, tunnelID string, body CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams, opts ...option.RequestOption) (res *CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseUnknown],
// [CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseArray]
// or [shared.UnionString].
type CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse interface {
	ImplementsCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseArray []interface{}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseArray) ImplementsCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse() {
}

// Union satisfied by
// [CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseUnknown],
// [CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseArray]
// or [shared.UnionString].
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse interface {
	ImplementsCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseArray []interface{}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseArray) ImplementsCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse() {
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelope struct {
	Errors   []CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelope]
type cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrors]
type cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessages]
type cfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeSuccess bool

const (
	CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeSuccessTrue CfdTunnelConfigurationCloudflareTunnelConfigurationGetConfigurationResponseEnvelopeSuccess = true
)

type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams struct {
	// The tunnel configuration and ingress rules.
	Config param.Field[CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig] `json:"config"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig struct {
	// List of public hostname definitions
	Ingress param.Field[[]CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress] `json:"ingress"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest] `json:"originRequest"`
	// Enable private network access from WARP users to private network routes
	WarpRouting param.Field[CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting] `json:"warp-routing"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CaPool param.Field[string] `json:"caPool"`
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
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CaPool param.Field[string] `json:"caPool"`
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
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelope struct {
	Errors   []CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelope]
type cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrors]
type cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessages]
type cfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeSuccess bool

const (
	CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeSuccessTrue CfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationResponseEnvelopeSuccess = true
)
