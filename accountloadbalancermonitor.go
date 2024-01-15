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

// AccountLoadBalancerMonitorService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLoadBalancerMonitorService] method instead.
type AccountLoadBalancerMonitorService struct {
	Options    []option.RequestOption
	Previews   *AccountLoadBalancerMonitorPreviewService
	References *AccountLoadBalancerMonitorReferenceService
}

// NewAccountLoadBalancerMonitorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerMonitorService(opts ...option.RequestOption) (r *AccountLoadBalancerMonitorService) {
	r = &AccountLoadBalancerMonitorService{}
	r.Options = opts
	r.Previews = NewAccountLoadBalancerMonitorPreviewService(opts...)
	r.References = NewAccountLoadBalancerMonitorReferenceService(opts...)
	return
}

// List a single configured monitor for an account.
func (r *AccountLoadBalancerMonitorService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *AccountLoadBalancerMonitorService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *AccountLoadBalancerMonitorService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountLoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a configured monitor.
func (r *AccountLoadBalancerMonitorService) AccountLoadBalancerMonitorsNewMonitor(ctx context.Context, accountIdentifier string, body AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured monitors for an account.
func (r *AccountLoadBalancerMonitorService) AccountLoadBalancerMonitorsListMonitors(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *AccountLoadBalancerMonitorService) Patch(ctx context.Context, accountIdentifier string, identifier string, body AccountLoadBalancerMonitorPatchParams, opts ...option.RequestOption) (res *Monitor, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Monitor struct {
	Errors   []MonitorError   `json:"errors"`
	Messages []MonitorMessage `json:"messages"`
	Result   MonitorResult    `json:"result"`
	// Whether the API call was successful
	Success MonitorSuccess `json:"success"`
	JSON    monitorJSON    `json:"-"`
}

// monitorJSON contains the JSON metadata for the struct [Monitor]
type monitorJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Monitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorError struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    monitorErrorJSON `json:"-"`
}

// monitorErrorJSON contains the JSON metadata for the struct [MonitorError]
type monitorErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorMessage struct {
	Code    int64              `json:"code,required"`
	Message string             `json:"message,required"`
	JSON    monitorMessageJSON `json:"-"`
}

// monitorMessageJSON contains the JSON metadata for the struct [MonitorMessage]
type monitorMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResult struct {
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
	Type MonitorResultType `json:"type"`
	JSON monitorResultJSON `json:"-"`
}

// monitorResultJSON contains the JSON metadata for the struct [MonitorResult]
type monitorResultJSON struct {
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

func (r *MonitorResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorResultType string

const (
	MonitorResultTypeHTTP     MonitorResultType = "http"
	MonitorResultTypeHTTPs    MonitorResultType = "https"
	MonitorResultTypeTcp      MonitorResultType = "tcp"
	MonitorResultTypeUdpIcmp  MonitorResultType = "udp_icmp"
	MonitorResultTypeIcmpPing MonitorResultType = "icmp_ping"
	MonitorResultTypeSmtp     MonitorResultType = "smtp"
)

// Whether the API call was successful
type MonitorSuccess bool

const (
	MonitorSuccessTrue MonitorSuccess = true
)

type AccountLoadBalancerMonitorDeleteResponse struct {
	Errors   []AccountLoadBalancerMonitorDeleteResponseError   `json:"errors"`
	Messages []AccountLoadBalancerMonitorDeleteResponseMessage `json:"messages"`
	Result   AccountLoadBalancerMonitorDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountLoadBalancerMonitorDeleteResponseSuccess `json:"success"`
	JSON    accountLoadBalancerMonitorDeleteResponseJSON    `json:"-"`
}

// accountLoadBalancerMonitorDeleteResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerMonitorDeleteResponse]
type accountLoadBalancerMonitorDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountLoadBalancerMonitorDeleteResponseErrorJSON `json:"-"`
}

// accountLoadBalancerMonitorDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountLoadBalancerMonitorDeleteResponseError]
type accountLoadBalancerMonitorDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountLoadBalancerMonitorDeleteResponseMessageJSON `json:"-"`
}

// accountLoadBalancerMonitorDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountLoadBalancerMonitorDeleteResponseMessage]
type accountLoadBalancerMonitorDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorDeleteResponseResult struct {
	ID   string                                             `json:"id"`
	JSON accountLoadBalancerMonitorDeleteResponseResultJSON `json:"-"`
}

// accountLoadBalancerMonitorDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountLoadBalancerMonitorDeleteResponseResult]
type accountLoadBalancerMonitorDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerMonitorDeleteResponseSuccess bool

const (
	AccountLoadBalancerMonitorDeleteResponseSuccessTrue AccountLoadBalancerMonitorDeleteResponseSuccess = true
)

type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse struct {
	Errors     []AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError    `json:"errors"`
	Messages   []AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage  `json:"messages"`
	Result     []AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult   `json:"result"`
	ResultInfo AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccess `json:"success"`
	JSON    accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON    `json:"-"`
}

// accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse]
type accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseErrorJSON `json:"-"`
}

// accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError]
type accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessageJSON `json:"-"`
}

// accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage]
type accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult struct {
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
	Type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType `json:"type"`
	JSON accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultJSON `json:"-"`
}

// accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult]
type accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultJSON struct {
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

func (r *AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType string

const (
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeHTTP     AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "http"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeHTTPs    AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "https"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeTcp      AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "tcp"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeUdpIcmp  AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "udp_icmp"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeIcmpPing AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "icmp_ping"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultTypeSmtp     AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultType = "smtp"
)

type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                 `json:"total_count"`
	JSON       accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfoJSON `json:"-"`
}

// accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo]
type accountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccess bool

const (
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccessTrue AccountLoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseSuccess = true
)

type AccountLoadBalancerMonitorUpdateParams struct {
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
	Type param.Field[AccountLoadBalancerMonitorUpdateParamsType] `json:"type"`
}

func (r AccountLoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type AccountLoadBalancerMonitorUpdateParamsType string

const (
	AccountLoadBalancerMonitorUpdateParamsTypeHTTP     AccountLoadBalancerMonitorUpdateParamsType = "http"
	AccountLoadBalancerMonitorUpdateParamsTypeHTTPs    AccountLoadBalancerMonitorUpdateParamsType = "https"
	AccountLoadBalancerMonitorUpdateParamsTypeTcp      AccountLoadBalancerMonitorUpdateParamsType = "tcp"
	AccountLoadBalancerMonitorUpdateParamsTypeUdpIcmp  AccountLoadBalancerMonitorUpdateParamsType = "udp_icmp"
	AccountLoadBalancerMonitorUpdateParamsTypeIcmpPing AccountLoadBalancerMonitorUpdateParamsType = "icmp_ping"
	AccountLoadBalancerMonitorUpdateParamsTypeSmtp     AccountLoadBalancerMonitorUpdateParamsType = "smtp"
)

type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams struct {
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
	Type param.Field[AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType] `json:"type"`
}

func (r AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType string

const (
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeHTTP     AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "http"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeHTTPs    AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "https"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeTcp      AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "tcp"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeUdpIcmp  AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "udp_icmp"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeIcmpPing AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "icmp_ping"
	AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsTypeSmtp     AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParamsType = "smtp"
)

type AccountLoadBalancerMonitorPatchParams struct {
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
	Type param.Field[AccountLoadBalancerMonitorPatchParamsType] `json:"type"`
}

func (r AccountLoadBalancerMonitorPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type AccountLoadBalancerMonitorPatchParamsType string

const (
	AccountLoadBalancerMonitorPatchParamsTypeHTTP     AccountLoadBalancerMonitorPatchParamsType = "http"
	AccountLoadBalancerMonitorPatchParamsTypeHTTPs    AccountLoadBalancerMonitorPatchParamsType = "https"
	AccountLoadBalancerMonitorPatchParamsTypeTcp      AccountLoadBalancerMonitorPatchParamsType = "tcp"
	AccountLoadBalancerMonitorPatchParamsTypeUdpIcmp  AccountLoadBalancerMonitorPatchParamsType = "udp_icmp"
	AccountLoadBalancerMonitorPatchParamsTypeIcmpPing AccountLoadBalancerMonitorPatchParamsType = "icmp_ping"
	AccountLoadBalancerMonitorPatchParamsTypeSmtp     AccountLoadBalancerMonitorPatchParamsType = "smtp"
)
