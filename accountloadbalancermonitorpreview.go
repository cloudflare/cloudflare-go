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

// AccountLoadBalancerMonitorPreviewService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLoadBalancerMonitorPreviewService] method instead.
type AccountLoadBalancerMonitorPreviewService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerMonitorPreviewService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerMonitorPreviewService(opts ...option.RequestOption) (r *AccountLoadBalancerMonitorPreviewService) {
	r = &AccountLoadBalancerMonitorPreviewService{}
	r.Options = opts
	return
}

// Preview pools using the specified monitor with provided monitor details. The
// returned preview_id can be used in the preview endpoint to retrieve the results.
func (r *AccountLoadBalancerMonitorPreviewService) AccountLoadBalancerMonitorsPreviewMonitor(ctx context.Context, accountIdentifier string, identifier interface{}, body AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParams, opts ...option.RequestOption) (res *PreviewResponsePfp0bPtX, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%v/preview", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PreviewResponsePfp0bPtX struct {
	Errors   []PreviewResponsePfp0bPtXError   `json:"errors"`
	Messages []PreviewResponsePfp0bPtXMessage `json:"messages"`
	Result   PreviewResponsePfp0bPtXResult    `json:"result"`
	// Whether the API call was successful
	Success PreviewResponsePfp0bPtXSuccess `json:"success"`
	JSON    previewResponsePfp0bPtXJSON    `json:"-"`
}

// previewResponsePfp0bPtXJSON contains the JSON metadata for the struct
// [PreviewResponsePfp0bPtX]
type previewResponsePfp0bPtXJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponsePfp0bPtX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResponsePfp0bPtXError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    previewResponsePfp0bPtXErrorJSON `json:"-"`
}

// previewResponsePfp0bPtXErrorJSON contains the JSON metadata for the struct
// [PreviewResponsePfp0bPtXError]
type previewResponsePfp0bPtXErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponsePfp0bPtXError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResponsePfp0bPtXMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    previewResponsePfp0bPtXMessageJSON `json:"-"`
}

// previewResponsePfp0bPtXMessageJSON contains the JSON metadata for the struct
// [PreviewResponsePfp0bPtXMessage]
type previewResponsePfp0bPtXMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponsePfp0bPtXMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResponsePfp0bPtXResult struct {
	Pools     interface{}                       `json:"pools"`
	PreviewID interface{}                       `json:"preview_id"`
	JSON      previewResponsePfp0bPtXResultJSON `json:"-"`
}

// previewResponsePfp0bPtXResultJSON contains the JSON metadata for the struct
// [PreviewResponsePfp0bPtXResult]
type previewResponsePfp0bPtXResultJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponsePfp0bPtXResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PreviewResponsePfp0bPtXSuccess bool

const (
	PreviewResponsePfp0bPtXSuccessTrue PreviewResponsePfp0bPtXSuccess = true
)

type AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParams struct {
	// The expected HTTP response codes or code ranges of the health check,
	// comma-separated. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType] `json:"type"`
}

func (r AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType string

const (
	AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsTypeHTTP     AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType = "http"
	AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsTypeHTTPs    AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType = "https"
	AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsTypeTcp      AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType = "tcp"
	AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsTypeUdpIcmp  AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType = "udp_icmp"
	AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsTypeIcmpPing AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType = "icmp_ping"
	AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsTypeSmtp     AccountLoadBalancerMonitorPreviewAccountLoadBalancerMonitorsPreviewMonitorParamsType = "smtp"
)
