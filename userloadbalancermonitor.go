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
func (r *UserLoadBalancerMonitorService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SingleResponseA6vzVv6O, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *UserLoadBalancerMonitorService) Update(ctx context.Context, identifier interface{}, body UserLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *SingleResponseA6vzVv6O, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *UserLoadBalancerMonitorService) Delete(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *IDResponse4120Y2Od, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a configured monitor.
func (r *UserLoadBalancerMonitorService) LoadBalancerMonitorsNewMonitor(ctx context.Context, body UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *SingleResponseA6vzVv6O, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured monitors for a user.
func (r *UserLoadBalancerMonitorService) LoadBalancerMonitorsListMonitors(ctx context.Context, opts ...option.RequestOption) (res *MonitorResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MonitorResponseCollection struct {
	Errors     []MonitorResponseCollectionError    `json:"errors"`
	Messages   []MonitorResponseCollectionMessage  `json:"messages"`
	Result     []MonitorResponseCollectionResult   `json:"result"`
	ResultInfo MonitorResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success MonitorResponseCollectionSuccess `json:"success"`
	JSON    monitorResponseCollectionJSON    `json:"-"`
}

// monitorResponseCollectionJSON contains the JSON metadata for the struct
// [MonitorResponseCollection]
type monitorResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResponseCollectionError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    monitorResponseCollectionErrorJSON `json:"-"`
}

// monitorResponseCollectionErrorJSON contains the JSON metadata for the struct
// [MonitorResponseCollectionError]
type monitorResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResponseCollectionMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    monitorResponseCollectionMessageJSON `json:"-"`
}

// monitorResponseCollectionMessageJSON contains the JSON metadata for the struct
// [MonitorResponseCollectionMessage]
type monitorResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResponseCollectionResult struct {
	ID interface{} `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool      `json:"allow_insecure"`
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
	// Port number to connect to for the health check. Required for TCP, UDP, and SMTP
	// checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type MonitorResponseCollectionResultType `json:"type"`
	JSON monitorResponseCollectionResultJSON `json:"-"`
}

// monitorResponseCollectionResultJSON contains the JSON metadata for the struct
// [MonitorResponseCollectionResult]
type monitorResponseCollectionResultJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
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
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MonitorResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorResponseCollectionResultType string

const (
	MonitorResponseCollectionResultTypeHTTP     MonitorResponseCollectionResultType = "http"
	MonitorResponseCollectionResultTypeHTTPs    MonitorResponseCollectionResultType = "https"
	MonitorResponseCollectionResultTypeTcp      MonitorResponseCollectionResultType = "tcp"
	MonitorResponseCollectionResultTypeUdpIcmp  MonitorResponseCollectionResultType = "udp_icmp"
	MonitorResponseCollectionResultTypeIcmpPing MonitorResponseCollectionResultType = "icmp_ping"
	MonitorResponseCollectionResultTypeSmtp     MonitorResponseCollectionResultType = "smtp"
)

type MonitorResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       monitorResponseCollectionResultInfoJSON `json:"-"`
}

// monitorResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [MonitorResponseCollectionResultInfo]
type monitorResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MonitorResponseCollectionSuccess bool

const (
	MonitorResponseCollectionSuccessTrue MonitorResponseCollectionSuccess = true
)

type SingleResponseA6vzVv6O struct {
	Errors   []SingleResponseA6vzVv6OError   `json:"errors"`
	Messages []SingleResponseA6vzVv6OMessage `json:"messages"`
	Result   SingleResponseA6vzVv6OResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseA6vzVv6OSuccess `json:"success"`
	JSON    singleResponseA6vzVv6OJSON    `json:"-"`
}

// singleResponseA6vzVv6OJSON contains the JSON metadata for the struct
// [SingleResponseA6vzVv6O]
type singleResponseA6vzVv6OJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseA6vzVv6O) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseA6vzVv6OError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseA6vzVv6OErrorJSON `json:"-"`
}

// singleResponseA6vzVv6OErrorJSON contains the JSON metadata for the struct
// [SingleResponseA6vzVv6OError]
type singleResponseA6vzVv6OErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseA6vzVv6OError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseA6vzVv6OMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseA6vzVv6OMessageJSON `json:"-"`
}

// singleResponseA6vzVv6OMessageJSON contains the JSON metadata for the struct
// [SingleResponseA6vzVv6OMessage]
type singleResponseA6vzVv6OMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseA6vzVv6OMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseA6vzVv6OResult struct {
	ID interface{} `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool      `json:"allow_insecure"`
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
	// Port number to connect to for the health check. Required for TCP, UDP, and SMTP
	// checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type SingleResponseA6vzVv6OResultType `json:"type"`
	JSON singleResponseA6vzVv6OResultJSON `json:"-"`
}

// singleResponseA6vzVv6OResultJSON contains the JSON metadata for the struct
// [SingleResponseA6vzVv6OResult]
type singleResponseA6vzVv6OResultJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
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
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SingleResponseA6vzVv6OResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type SingleResponseA6vzVv6OResultType string

const (
	SingleResponseA6vzVv6OResultTypeHTTP     SingleResponseA6vzVv6OResultType = "http"
	SingleResponseA6vzVv6OResultTypeHTTPs    SingleResponseA6vzVv6OResultType = "https"
	SingleResponseA6vzVv6OResultTypeTcp      SingleResponseA6vzVv6OResultType = "tcp"
	SingleResponseA6vzVv6OResultTypeUdpIcmp  SingleResponseA6vzVv6OResultType = "udp_icmp"
	SingleResponseA6vzVv6OResultTypeIcmpPing SingleResponseA6vzVv6OResultType = "icmp_ping"
	SingleResponseA6vzVv6OResultTypeSmtp     SingleResponseA6vzVv6OResultType = "smtp"
)

// Whether the API call was successful
type SingleResponseA6vzVv6OSuccess bool

const (
	SingleResponseA6vzVv6OSuccessTrue SingleResponseA6vzVv6OSuccess = true
)

type UserLoadBalancerMonitorUpdateParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// Object description.
	Description param.Field[string] `json:"description"`
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
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
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
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// Object description.
	Description param.Field[string] `json:"description"`
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
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
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
