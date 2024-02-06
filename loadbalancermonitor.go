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

// LoadBalancerMonitorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerMonitorService]
// method instead.
type LoadBalancerMonitorService struct {
	Options    []option.RequestOption
	Previews   *LoadBalancerMonitorPreviewService
	References *LoadBalancerMonitorReferenceService
}

// NewLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerMonitorService(opts ...option.RequestOption) (r *LoadBalancerMonitorService) {
	r = &LoadBalancerMonitorService{}
	r.Options = opts
	r.Previews = NewLoadBalancerMonitorPreviewService(opts...)
	r.References = NewLoadBalancerMonitorReferenceService(opts...)
	return
}

// List a single configured monitor for an account.
func (r *LoadBalancerMonitorService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerMonitorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *LoadBalancerMonitorService) Update(ctx context.Context, accountIdentifier string, identifier string, body LoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *LoadBalancerMonitorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *LoadBalancerMonitorService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a configured monitor.
func (r *LoadBalancerMonitorService) AccountLoadBalancerMonitorsNewMonitor(ctx context.Context, accountIdentifier string, body LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured monitors for an account.
func (r *LoadBalancerMonitorService) AccountLoadBalancerMonitorsListMonitors(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerMonitorGetResponse struct {
	Errors   []LoadBalancerMonitorGetResponseError   `json:"errors"`
	Messages []LoadBalancerMonitorGetResponseMessage `json:"messages"`
	Result   LoadBalancerMonitorGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorGetResponseSuccess `json:"success"`
	JSON    loadBalancerMonitorGetResponseJSON    `json:"-"`
}

// loadBalancerMonitorGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorGetResponse]
type loadBalancerMonitorGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    loadBalancerMonitorGetResponseErrorJSON `json:"-"`
}

// loadBalancerMonitorGetResponseErrorJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorGetResponseError]
type loadBalancerMonitorGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    loadBalancerMonitorGetResponseMessageJSON `json:"-"`
}

// loadBalancerMonitorGetResponseMessageJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorGetResponseMessage]
type loadBalancerMonitorGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseResult struct {
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
	Type LoadBalancerMonitorGetResponseResultType `json:"type"`
	JSON loadBalancerMonitorGetResponseResultJSON `json:"-"`
}

// loadBalancerMonitorGetResponseResultJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorGetResponseResult]
type loadBalancerMonitorGetResponseResultJSON struct {
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

func (r *LoadBalancerMonitorGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorGetResponseResultType string

const (
	LoadBalancerMonitorGetResponseResultTypeHTTP     LoadBalancerMonitorGetResponseResultType = "http"
	LoadBalancerMonitorGetResponseResultTypeHTTPs    LoadBalancerMonitorGetResponseResultType = "https"
	LoadBalancerMonitorGetResponseResultTypeTcp      LoadBalancerMonitorGetResponseResultType = "tcp"
	LoadBalancerMonitorGetResponseResultTypeUdpIcmp  LoadBalancerMonitorGetResponseResultType = "udp_icmp"
	LoadBalancerMonitorGetResponseResultTypeIcmpPing LoadBalancerMonitorGetResponseResultType = "icmp_ping"
	LoadBalancerMonitorGetResponseResultTypeSmtp     LoadBalancerMonitorGetResponseResultType = "smtp"
)

// Whether the API call was successful
type LoadBalancerMonitorGetResponseSuccess bool

const (
	LoadBalancerMonitorGetResponseSuccessTrue LoadBalancerMonitorGetResponseSuccess = true
)

type LoadBalancerMonitorUpdateResponse struct {
	Errors   []LoadBalancerMonitorUpdateResponseError   `json:"errors"`
	Messages []LoadBalancerMonitorUpdateResponseMessage `json:"messages"`
	Result   LoadBalancerMonitorUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorUpdateResponseSuccess `json:"success"`
	JSON    loadBalancerMonitorUpdateResponseJSON    `json:"-"`
}

// loadBalancerMonitorUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorUpdateResponse]
type loadBalancerMonitorUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    loadBalancerMonitorUpdateResponseErrorJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseErrorJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorUpdateResponseError]
type loadBalancerMonitorUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    loadBalancerMonitorUpdateResponseMessageJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseMessageJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorUpdateResponseMessage]
type loadBalancerMonitorUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseResult struct {
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
	Type LoadBalancerMonitorUpdateResponseResultType `json:"type"`
	JSON loadBalancerMonitorUpdateResponseResultJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseResultJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorUpdateResponseResult]
type loadBalancerMonitorUpdateResponseResultJSON struct {
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

func (r *LoadBalancerMonitorUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorUpdateResponseResultType string

const (
	LoadBalancerMonitorUpdateResponseResultTypeHTTP     LoadBalancerMonitorUpdateResponseResultType = "http"
	LoadBalancerMonitorUpdateResponseResultTypeHTTPs    LoadBalancerMonitorUpdateResponseResultType = "https"
	LoadBalancerMonitorUpdateResponseResultTypeTcp      LoadBalancerMonitorUpdateResponseResultType = "tcp"
	LoadBalancerMonitorUpdateResponseResultTypeUdpIcmp  LoadBalancerMonitorUpdateResponseResultType = "udp_icmp"
	LoadBalancerMonitorUpdateResponseResultTypeIcmpPing LoadBalancerMonitorUpdateResponseResultType = "icmp_ping"
	LoadBalancerMonitorUpdateResponseResultTypeSmtp     LoadBalancerMonitorUpdateResponseResultType = "smtp"
)

// Whether the API call was successful
type LoadBalancerMonitorUpdateResponseSuccess bool

const (
	LoadBalancerMonitorUpdateResponseSuccessTrue LoadBalancerMonitorUpdateResponseSuccess = true
)

type LoadBalancerMonitorDeleteResponse struct {
	Errors   []LoadBalancerMonitorDeleteResponseError   `json:"errors"`
	Messages []LoadBalancerMonitorDeleteResponseMessage `json:"messages"`
	Result   LoadBalancerMonitorDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorDeleteResponseSuccess `json:"success"`
	JSON    loadBalancerMonitorDeleteResponseJSON    `json:"-"`
}

// loadBalancerMonitorDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorDeleteResponse]
type loadBalancerMonitorDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    loadBalancerMonitorDeleteResponseErrorJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseErrorJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorDeleteResponseError]
type loadBalancerMonitorDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    loadBalancerMonitorDeleteResponseMessageJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseMessageJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorDeleteResponseMessage]
type loadBalancerMonitorDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseResult struct {
	ID   string                                      `json:"id"`
	JSON loadBalancerMonitorDeleteResponseResultJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseResultJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorDeleteResponseResult]
type loadBalancerMonitorDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorDeleteResponseSuccess bool

const (
	LoadBalancerMonitorDeleteResponseSuccessTrue LoadBalancerMonitorDeleteResponseSuccess = true
)

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse struct {
	Errors   []LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseError   `json:"errors"`
	Messages []LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessage `json:"messages"`
	Result   LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseSuccess `json:"success"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseJSON    `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseJSON contains
// the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseErrorJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseErrorJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseError]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessageJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessageJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessage]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResult struct {
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
	Type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType `json:"type"`
	JSON loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResult]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultJSON struct {
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

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType string

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultTypeHTTP     LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType = "http"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultTypeHTTPs    LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType = "https"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultTypeTcp      LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType = "tcp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultTypeUdpIcmp  LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType = "udp_icmp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultTypeIcmpPing LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType = "icmp_ping"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultTypeSmtp     LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseResultType = "smtp"
)

// Whether the API call was successful
type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseSuccess bool

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseSuccessTrue LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseSuccess = true
)

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse struct {
	Errors     []LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError    `json:"errors"`
	Messages   []LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage  `json:"messages"`
	Result     []LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult   `json:"result"`
	ResultInfo LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccess `json:"success"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON    `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON contains
// the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseErrorJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseErrorJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessageJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessageJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult struct {
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
	Type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType `json:"type"`
	JSON loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultJSON struct {
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

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType string

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeHTTP     LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "http"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeHTTPs    LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "https"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeTcp      LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "tcp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeUdpIcmp  LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "udp_icmp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeIcmpPing LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "icmp_ping"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeSmtp     LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "smtp"
)

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfoJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccess bool

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccessTrue LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccess = true
)

type LoadBalancerMonitorUpdateParams struct {
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
	Type param.Field[LoadBalancerMonitorUpdateParamsType] `json:"type"`
}

func (r LoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorUpdateParamsType string

const (
	LoadBalancerMonitorUpdateParamsTypeHTTP     LoadBalancerMonitorUpdateParamsType = "http"
	LoadBalancerMonitorUpdateParamsTypeHTTPs    LoadBalancerMonitorUpdateParamsType = "https"
	LoadBalancerMonitorUpdateParamsTypeTcp      LoadBalancerMonitorUpdateParamsType = "tcp"
	LoadBalancerMonitorUpdateParamsTypeUdpIcmp  LoadBalancerMonitorUpdateParamsType = "udp_icmp"
	LoadBalancerMonitorUpdateParamsTypeIcmpPing LoadBalancerMonitorUpdateParamsType = "icmp_ping"
	LoadBalancerMonitorUpdateParamsTypeSmtp     LoadBalancerMonitorUpdateParamsType = "smtp"
)

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams struct {
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
	Type param.Field[LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType] `json:"type"`
}

func (r LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType string

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeHTTP     LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "http"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeHTTPs    LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "https"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeTcp      LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "tcp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeUdpIcmp  LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "udp_icmp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeIcmpPing LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "icmp_ping"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeSmtp     LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "smtp"
)
