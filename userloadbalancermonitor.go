// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerMonitorService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerMonitorService] method instead.
type UserLoadBalancerMonitorService struct {
	Options    []option.RequestOption
	Previews   *UserLoadBalancerMonitorPreviewService
	References *UserLoadBalancerMonitorReferenceService
}

// NewUserLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorService) {
	r = &UserLoadBalancerMonitorService{}
	r.Options = opts
	r.Previews = NewUserLoadBalancerMonitorPreviewService(opts...)
	r.References = NewUserLoadBalancerMonitorReferenceService(opts...)
	return
}

// List a single configured monitor for a user.
func (r *UserLoadBalancerMonitorService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *UserLoadBalancerMonitorService) Update(ctx context.Context, identifier string, body UserLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *UserLoadBalancerMonitorService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserLoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a configured monitor.
func (r *UserLoadBalancerMonitorService) LoadBalancerMonitorsNewMonitor(ctx context.Context, body UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured monitors for a user.
func (r *UserLoadBalancerMonitorService) LoadBalancerMonitorsListMonitors(ctx context.Context, opts ...option.RequestOption) (res *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *UserLoadBalancerMonitorService) Patch(ctx context.Context, identifier string, body UserLoadBalancerMonitorPatchParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type UserLoadBalancerMonitorDeleteResponse struct {
	Errors   []UserLoadBalancerMonitorDeleteResponseError   `json:"errors"`
	Messages []UserLoadBalancerMonitorDeleteResponseMessage `json:"messages"`
	Result   UserLoadBalancerMonitorDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorDeleteResponseSuccess `json:"success"`
	JSON    userLoadBalancerMonitorDeleteResponseJSON    `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseJSON contains the JSON metadata for the
// struct [UserLoadBalancerMonitorDeleteResponse]
type userLoadBalancerMonitorDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    userLoadBalancerMonitorDeleteResponseErrorJSON `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseErrorJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorDeleteResponseError]
type userLoadBalancerMonitorDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userLoadBalancerMonitorDeleteResponseMessageJSON `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseMessageJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorDeleteResponseMessage]
type userLoadBalancerMonitorDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorDeleteResponseResult struct {
	ID   string                                          `json:"id"`
	JSON userLoadBalancerMonitorDeleteResponseResultJSON `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseResultJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorDeleteResponseResult]
type userLoadBalancerMonitorDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorDeleteResponseSuccess bool

const (
	UserLoadBalancerMonitorDeleteResponseSuccessTrue UserLoadBalancerMonitorDeleteResponseSuccess = true
)

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse struct {
	Errors     []UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseError    `json:"errors"`
	Messages   []UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessage  `json:"messages"`
	Result     []UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResult   `json:"result"`
	ResultInfo UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseSuccess `json:"success"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseJSON    `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseErrorJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseErrorJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseError]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessageJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessageJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessage]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResult struct {
	ID string `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown int64 `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp int64     `json:"consecutive_up"`
	CreatedOn     time.Time `json:"created_on" format:"date-time"`
	// Object description.
	Description string `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes string `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval int64 `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method     string    `json:"method"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path string `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone string `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType `json:"type"`
	JSON userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResult]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
	ConsecutiveDown apijson.Field
	ConsecutiveUp   apijson.Field
	CreatedOn       apijson.Field
	Description     apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Interval        apijson.Field
	Method          apijson.Field
	ModifiedOn      apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	ProbeZone       apijson.Field
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType string

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultTypeHTTP     UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType = "http"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultTypeHTTPs    UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType = "https"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultTypeTcp      UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType = "tcp"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultTypeUdpIcmp  UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType = "udp_icmp"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultTypeIcmpPing UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType = "icmp_ping"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultTypeSmtp     UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultType = "smtp"
)

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfoJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfo]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseSuccess bool

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseSuccessTrue UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseSuccess = true
)

type UserLoadBalancerMonitorUpdateParams struct {
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
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
	Type param.Field[UserLoadBalancerMonitorUpdateParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorUpdateParamsType string

const (
	UserLoadBalancerMonitorUpdateParamsTypeHTTP     UserLoadBalancerMonitorUpdateParamsType = "http"
	UserLoadBalancerMonitorUpdateParamsTypeHTTPs    UserLoadBalancerMonitorUpdateParamsType = "https"
	UserLoadBalancerMonitorUpdateParamsTypeTcp      UserLoadBalancerMonitorUpdateParamsType = "tcp"
	UserLoadBalancerMonitorUpdateParamsTypeUdpIcmp  UserLoadBalancerMonitorUpdateParamsType = "udp_icmp"
	UserLoadBalancerMonitorUpdateParamsTypeIcmpPing UserLoadBalancerMonitorUpdateParamsType = "icmp_ping"
	UserLoadBalancerMonitorUpdateParamsTypeSmtp     UserLoadBalancerMonitorUpdateParamsType = "smtp"
)

type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams struct {
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
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
	Type param.Field[UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType string

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeHTTP     UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "http"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeHTTPs    UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "https"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeTcp      UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "tcp"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeUdpIcmp  UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "udp_icmp"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeIcmpPing UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "icmp_ping"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeSmtp     UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "smtp"
)

type UserLoadBalancerMonitorPatchParams struct {
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
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
	Type param.Field[UserLoadBalancerMonitorPatchParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorPatchParamsType string

const (
	UserLoadBalancerMonitorPatchParamsTypeHTTP     UserLoadBalancerMonitorPatchParamsType = "http"
	UserLoadBalancerMonitorPatchParamsTypeHTTPs    UserLoadBalancerMonitorPatchParamsType = "https"
	UserLoadBalancerMonitorPatchParamsTypeTcp      UserLoadBalancerMonitorPatchParamsType = "tcp"
	UserLoadBalancerMonitorPatchParamsTypeUdpIcmp  UserLoadBalancerMonitorPatchParamsType = "udp_icmp"
	UserLoadBalancerMonitorPatchParamsTypeIcmpPing UserLoadBalancerMonitorPatchParamsType = "icmp_ping"
	UserLoadBalancerMonitorPatchParamsTypeSmtp     UserLoadBalancerMonitorPatchParamsType = "smtp"
)
