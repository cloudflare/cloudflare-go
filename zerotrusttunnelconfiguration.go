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

// ZeroTrustTunnelConfigurationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustTunnelConfigurationService] method instead.
type ZeroTrustTunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewZeroTrustTunnelConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustTunnelConfigurationService(opts ...option.RequestOption) (r *ZeroTrustTunnelConfigurationService) {
	r = &ZeroTrustTunnelConfigurationService{}
	r.Options = opts
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *ZeroTrustTunnelConfigurationService) Update(ctx context.Context, tunnelID string, params ZeroTrustTunnelConfigurationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustTunnelConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelConfigurationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *ZeroTrustTunnelConfigurationService) Get(ctx context.Context, tunnelID string, query ZeroTrustTunnelConfigurationGetParams, opts ...option.RequestOption) (res *ZeroTrustTunnelConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelConfigurationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustTunnelConfigurationUpdateResponseUnknown],
// [ZeroTrustTunnelConfigurationUpdateResponseArray] or [shared.UnionString].
type ZeroTrustTunnelConfigurationUpdateResponse interface {
	ImplementsZeroTrustTunnelConfigurationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelConfigurationUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustTunnelConfigurationUpdateResponseArray []interface{}

func (r ZeroTrustTunnelConfigurationUpdateResponseArray) ImplementsZeroTrustTunnelConfigurationUpdateResponse() {
}

// Union satisfied by [ZeroTrustTunnelConfigurationGetResponseUnknown],
// [ZeroTrustTunnelConfigurationGetResponseArray] or [shared.UnionString].
type ZeroTrustTunnelConfigurationGetResponse interface {
	ImplementsZeroTrustTunnelConfigurationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelConfigurationGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustTunnelConfigurationGetResponseArray []interface{}

func (r ZeroTrustTunnelConfigurationGetResponseArray) ImplementsZeroTrustTunnelConfigurationGetResponse() {
}

type ZeroTrustTunnelConfigurationUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The tunnel configuration and ingress rules.
	Config param.Field[ZeroTrustTunnelConfigurationUpdateParamsConfig] `json:"config"`
}

func (r ZeroTrustTunnelConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type ZeroTrustTunnelConfigurationUpdateParamsConfig struct {
	// List of public hostname definitions
	Ingress param.Field[[]ZeroTrustTunnelConfigurationUpdateParamsConfigIngress] `json:"ingress"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[ZeroTrustTunnelConfigurationUpdateParamsConfigOriginRequest] `json:"originRequest"`
	// Enable private network access from WARP users to private network routes
	WARPRouting param.Field[ZeroTrustTunnelConfigurationUpdateParamsConfigWARPRouting] `json:"warp-routing"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type ZeroTrustTunnelConfigurationUpdateParamsConfigIngress struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[ZeroTrustTunnelConfigurationUpdateParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
type ZeroTrustTunnelConfigurationUpdateParamsConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[ZeroTrustTunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess] `json:"access"`
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
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfigIngressOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type ZeroTrustTunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
type ZeroTrustTunnelConfigurationUpdateParamsConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[ZeroTrustTunnelConfigurationUpdateParamsConfigOriginRequestAccess] `json:"access"`
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
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfigOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type ZeroTrustTunnelConfigurationUpdateParamsConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes
type ZeroTrustTunnelConfigurationUpdateParamsConfigWARPRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustTunnelConfigurationUpdateParamsConfigWARPRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustTunnelConfigurationUpdateResponseEnvelope struct {
	Errors   []ZeroTrustTunnelConfigurationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelConfigurationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelConfigurationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelConfigurationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConfigurationUpdateResponseEnvelope]
type zeroTrustTunnelConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustTunnelConfigurationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustTunnelConfigurationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelConfigurationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustTunnelConfigurationUpdateResponseEnvelopeErrors]
type zeroTrustTunnelConfigurationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConfigurationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustTunnelConfigurationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustTunnelConfigurationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelConfigurationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustTunnelConfigurationUpdateResponseEnvelopeMessages]
type zeroTrustTunnelConfigurationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConfigurationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustTunnelConfigurationUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelConfigurationUpdateResponseEnvelopeSuccessTrue ZeroTrustTunnelConfigurationUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustTunnelConfigurationGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustTunnelConfigurationGetResponseEnvelope struct {
	Errors   []ZeroTrustTunnelConfigurationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelConfigurationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelConfigurationGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelConfigurationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustTunnelConfigurationGetResponseEnvelope]
type zeroTrustTunnelConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustTunnelConfigurationGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustTunnelConfigurationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelConfigurationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConfigurationGetResponseEnvelopeErrors]
type zeroTrustTunnelConfigurationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConfigurationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustTunnelConfigurationGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustTunnelConfigurationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelConfigurationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustTunnelConfigurationGetResponseEnvelopeMessages]
type zeroTrustTunnelConfigurationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConfigurationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustTunnelConfigurationGetResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelConfigurationGetResponseEnvelopeSuccessTrue ZeroTrustTunnelConfigurationGetResponseEnvelopeSuccess = true
)
