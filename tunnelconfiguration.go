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

// TunnelConfigurationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelConfigurationService]
// method instead.
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
func (r *TunnelConfigurationService) Update(ctx context.Context, accountID string, tunnelID string, body TunnelConfigurationUpdateParams, opts ...option.RequestOption) (res *TunnelConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConfigurationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *TunnelConfigurationService) List(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelConfigurationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConfigurationListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [TunnelConfigurationUpdateResponseUnknown],
// [TunnelConfigurationUpdateResponseArray] or [shared.UnionString].
type TunnelConfigurationUpdateResponse interface {
	ImplementsTunnelConfigurationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConfigurationUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConfigurationUpdateResponseArray []interface{}

func (r TunnelConfigurationUpdateResponseArray) ImplementsTunnelConfigurationUpdateResponse() {}

// Union satisfied by [TunnelConfigurationListResponseUnknown],
// [TunnelConfigurationListResponseArray] or [shared.UnionString].
type TunnelConfigurationListResponse interface {
	ImplementsTunnelConfigurationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConfigurationListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConfigurationListResponseArray []interface{}

func (r TunnelConfigurationListResponseArray) ImplementsTunnelConfigurationListResponse() {}

type TunnelConfigurationUpdateParams struct {
	// The tunnel configuration and ingress rules.
	Config param.Field[TunnelConfigurationUpdateParamsConfig] `json:"config"`
}

func (r TunnelConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type TunnelConfigurationUpdateParamsConfig struct {
	// List of public hostname definitions
	Ingress param.Field[[]TunnelConfigurationUpdateParamsConfigIngress] `json:"ingress"`
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[TunnelConfigurationUpdateParamsConfigOriginRequest] `json:"originRequest"`
	// Enable private network access from WARP users to private network routes
	WARPRouting param.Field[TunnelConfigurationUpdateParamsConfigWARPRouting] `json:"warp-routing"`
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
	// Configuration parameters of connection between cloudflared and origin server.
	OriginRequest param.Field[TunnelConfigurationUpdateParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r TunnelConfigurationUpdateParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
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
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
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
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters of connection between cloudflared and origin server.
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
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
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
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelConfigurationUpdateParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes
type TunnelConfigurationUpdateParamsConfigWARPRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TunnelConfigurationUpdateParamsConfigWARPRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelConfigurationUpdateResponseEnvelope struct {
	Errors   []TunnelConfigurationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConfigurationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelConfigurationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// tunnelConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConfigurationUpdateResponseEnvelope]
type tunnelConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConfigurationUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    tunnelConfigurationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [TunnelConfigurationUpdateResponseEnvelopeErrors]
type tunnelConfigurationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConfigurationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    tunnelConfigurationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelConfigurationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TunnelConfigurationUpdateResponseEnvelopeMessages]
type tunnelConfigurationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConfigurationUpdateResponseEnvelopeSuccess bool

const (
	TunnelConfigurationUpdateResponseEnvelopeSuccessTrue TunnelConfigurationUpdateResponseEnvelopeSuccess = true
)

type TunnelConfigurationListResponseEnvelope struct {
	Errors   []TunnelConfigurationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConfigurationListResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelConfigurationListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConfigurationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConfigurationListResponseEnvelopeJSON    `json:"-"`
}

// tunnelConfigurationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConfigurationListResponseEnvelope]
type tunnelConfigurationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConfigurationListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    tunnelConfigurationListResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelConfigurationListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TunnelConfigurationListResponseEnvelopeErrors]
type tunnelConfigurationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConfigurationListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    tunnelConfigurationListResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelConfigurationListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TunnelConfigurationListResponseEnvelopeMessages]
type tunnelConfigurationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConfigurationListResponseEnvelopeSuccess bool

const (
	TunnelConfigurationListResponseEnvelopeSuccessTrue TunnelConfigurationListResponseEnvelopeSuccess = true
)
