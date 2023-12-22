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
func (r *AccountLoadBalancerMonitorService) Get(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *MonitorSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *AccountLoadBalancerMonitorService) Update(ctx context.Context, accountIdentifier string, identifier interface{}, body AccountLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *MonitorSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *AccountLoadBalancerMonitorService) Delete(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *IDResponse4120Y2Od, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a configured monitor.
func (r *AccountLoadBalancerMonitorService) AccountLoadBalancerMonitorsNewMonitor(ctx context.Context, accountIdentifier string, body AccountLoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *MonitorSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured monitors for an account.
func (r *AccountLoadBalancerMonitorService) AccountLoadBalancerMonitorsListMonitors(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MonitorResponseCollection2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IDResponse4120Y2Od struct {
	Errors   []IDResponse4120Y2OdError   `json:"errors"`
	Messages []IDResponse4120Y2OdMessage `json:"messages"`
	Result   IDResponse4120Y2OdResult    `json:"result"`
	// Whether the API call was successful
	Success IDResponse4120Y2OdSuccess `json:"success"`
	JSON    idResponse4120Y2OdJSON    `json:"-"`
}

// idResponse4120Y2OdJSON contains the JSON metadata for the struct
// [IDResponse4120Y2Od]
type idResponse4120Y2OdJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponse4120Y2Od) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponse4120Y2OdError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    idResponse4120Y2OdErrorJSON `json:"-"`
}

// idResponse4120Y2OdErrorJSON contains the JSON metadata for the struct
// [IDResponse4120Y2OdError]
type idResponse4120Y2OdErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponse4120Y2OdError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponse4120Y2OdMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    idResponse4120Y2OdMessageJSON `json:"-"`
}

// idResponse4120Y2OdMessageJSON contains the JSON metadata for the struct
// [IDResponse4120Y2OdMessage]
type idResponse4120Y2OdMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponse4120Y2OdMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponse4120Y2OdResult struct {
	ID   interface{}                  `json:"id"`
	JSON idResponse4120Y2OdResultJSON `json:"-"`
}

// idResponse4120Y2OdResultJSON contains the JSON metadata for the struct
// [IDResponse4120Y2OdResult]
type idResponse4120Y2OdResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponse4120Y2OdResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IDResponse4120Y2OdSuccess bool

const (
	IDResponse4120Y2OdSuccessTrue IDResponse4120Y2OdSuccess = true
)

type MonitorResponseCollection2 struct {
	Errors     []MonitorResponseCollection2Error    `json:"errors"`
	Messages   []MonitorResponseCollection2Message  `json:"messages"`
	Result     []MonitorResponseCollection2Result   `json:"result"`
	ResultInfo MonitorResponseCollection2ResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success MonitorResponseCollection2Success `json:"success"`
	JSON    monitorResponseCollection2JSON    `json:"-"`
}

// monitorResponseCollection2JSON contains the JSON metadata for the struct
// [MonitorResponseCollection2]
type monitorResponseCollection2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollection2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResponseCollection2Error struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    monitorResponseCollection2ErrorJSON `json:"-"`
}

// monitorResponseCollection2ErrorJSON contains the JSON metadata for the struct
// [MonitorResponseCollection2Error]
type monitorResponseCollection2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollection2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResponseCollection2Message struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    monitorResponseCollection2MessageJSON `json:"-"`
}

// monitorResponseCollection2MessageJSON contains the JSON metadata for the struct
// [MonitorResponseCollection2Message]
type monitorResponseCollection2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollection2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorResponseCollection2Result struct {
	ID interface{} `json:"id"`
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
	// The expected HTTP response codes or code ranges of the health check,
	// comma-separated. This parameter is only valid for HTTP and HTTPS monitors.
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
	Type MonitorResponseCollection2ResultType `json:"type"`
	JSON monitorResponseCollection2ResultJSON `json:"-"`
}

// monitorResponseCollection2ResultJSON contains the JSON metadata for the struct
// [MonitorResponseCollection2Result]
type monitorResponseCollection2ResultJSON struct {
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

func (r *MonitorResponseCollection2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorResponseCollection2ResultType string

const (
	MonitorResponseCollection2ResultTypeHTTP     MonitorResponseCollection2ResultType = "http"
	MonitorResponseCollection2ResultTypeHTTPs    MonitorResponseCollection2ResultType = "https"
	MonitorResponseCollection2ResultTypeTcp      MonitorResponseCollection2ResultType = "tcp"
	MonitorResponseCollection2ResultTypeUdpIcmp  MonitorResponseCollection2ResultType = "udp_icmp"
	MonitorResponseCollection2ResultTypeIcmpPing MonitorResponseCollection2ResultType = "icmp_ping"
	MonitorResponseCollection2ResultTypeSmtp     MonitorResponseCollection2ResultType = "smtp"
)

type MonitorResponseCollection2ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       monitorResponseCollection2ResultInfoJSON `json:"-"`
}

// monitorResponseCollection2ResultInfoJSON contains the JSON metadata for the
// struct [MonitorResponseCollection2ResultInfo]
type monitorResponseCollection2ResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorResponseCollection2ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MonitorResponseCollection2Success bool

const (
	MonitorResponseCollection2SuccessTrue MonitorResponseCollection2Success = true
)

type MonitorSingleResponse struct {
	Errors   []MonitorSingleResponseError   `json:"errors"`
	Messages []MonitorSingleResponseMessage `json:"messages"`
	Result   MonitorSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success MonitorSingleResponseSuccess `json:"success"`
	JSON    monitorSingleResponseJSON    `json:"-"`
}

// monitorSingleResponseJSON contains the JSON metadata for the struct
// [MonitorSingleResponse]
type monitorSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorSingleResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    monitorSingleResponseErrorJSON `json:"-"`
}

// monitorSingleResponseErrorJSON contains the JSON metadata for the struct
// [MonitorSingleResponseError]
type monitorSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorSingleResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    monitorSingleResponseMessageJSON `json:"-"`
}

// monitorSingleResponseMessageJSON contains the JSON metadata for the struct
// [MonitorSingleResponseMessage]
type monitorSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorSingleResponseResult struct {
	ID interface{} `json:"id"`
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
	// The expected HTTP response codes or code ranges of the health check,
	// comma-separated. This parameter is only valid for HTTP and HTTPS monitors.
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
	Type MonitorSingleResponseResultType `json:"type"`
	JSON monitorSingleResponseResultJSON `json:"-"`
}

// monitorSingleResponseResultJSON contains the JSON metadata for the struct
// [MonitorSingleResponseResult]
type monitorSingleResponseResultJSON struct {
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

func (r *MonitorSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorSingleResponseResultType string

const (
	MonitorSingleResponseResultTypeHTTP     MonitorSingleResponseResultType = "http"
	MonitorSingleResponseResultTypeHTTPs    MonitorSingleResponseResultType = "https"
	MonitorSingleResponseResultTypeTcp      MonitorSingleResponseResultType = "tcp"
	MonitorSingleResponseResultTypeUdpIcmp  MonitorSingleResponseResultType = "udp_icmp"
	MonitorSingleResponseResultTypeIcmpPing MonitorSingleResponseResultType = "icmp_ping"
	MonitorSingleResponseResultTypeSmtp     MonitorSingleResponseResultType = "smtp"
)

// Whether the API call was successful
type MonitorSingleResponseSuccess bool

const (
	MonitorSingleResponseSuccessTrue MonitorSingleResponseSuccess = true
)

type AccountLoadBalancerMonitorUpdateParams struct {
	// The expected HTTP response codes or code ranges of the health check,
	// comma-separated. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
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
	// The expected HTTP response codes or code ranges of the health check,
	// comma-separated. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
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
