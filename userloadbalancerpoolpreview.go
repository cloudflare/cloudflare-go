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

// UserLoadBalancerPoolPreviewService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerPoolPreviewService] method instead.
type UserLoadBalancerPoolPreviewService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPoolPreviewService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolPreviewService(opts ...option.RequestOption) (r *UserLoadBalancerPoolPreviewService) {
	r = &UserLoadBalancerPoolPreviewService{}
	r.Options = opts
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *UserLoadBalancerPoolPreviewService) LoadBalancerPoolsPreviewPool(ctx context.Context, identifier interface{}, body UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParams, opts ...option.RequestOption) (res *PreviewResponsePfp0bPtX, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%v/preview", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
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
	// Port number to connect to for the health check. Required for TCP, UDP, and SMTP
	// checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType] `json:"type"`
}

func (r UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType string

const (
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeHTTP     UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "http"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeHTTPs    UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "https"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeTcp      UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "tcp"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeUdpIcmp  UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "udp_icmp"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeIcmpPing UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "icmp_ping"
	UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsTypeSmtp     UserLoadBalancerPoolPreviewLoadBalancerPoolsPreviewPoolParamsType = "smtp"
)
