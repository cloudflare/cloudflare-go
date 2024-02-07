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
	var env LoadBalancerMonitorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor.
func (r *LoadBalancerMonitorService) Update(ctx context.Context, accountIdentifier string, identifier string, body LoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *LoadBalancerMonitorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured monitor.
func (r *LoadBalancerMonitorService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a configured monitor.
func (r *LoadBalancerMonitorService) AccountLoadBalancerMonitorsNewMonitor(ctx context.Context, accountIdentifier string, body LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for an account.
func (r *LoadBalancerMonitorService) AccountLoadBalancerMonitorsListMonitors(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerMonitorGetResponse struct {
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
	Type LoadBalancerMonitorGetResponseType `json:"type"`
	JSON loadBalancerMonitorGetResponseJSON `json:"-"`
}

// loadBalancerMonitorGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorGetResponse]
type loadBalancerMonitorGetResponseJSON struct {
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

func (r *LoadBalancerMonitorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorGetResponseType string

const (
	LoadBalancerMonitorGetResponseTypeHTTP     LoadBalancerMonitorGetResponseType = "http"
	LoadBalancerMonitorGetResponseTypeHTTPs    LoadBalancerMonitorGetResponseType = "https"
	LoadBalancerMonitorGetResponseTypeTcp      LoadBalancerMonitorGetResponseType = "tcp"
	LoadBalancerMonitorGetResponseTypeUdpIcmp  LoadBalancerMonitorGetResponseType = "udp_icmp"
	LoadBalancerMonitorGetResponseTypeIcmpPing LoadBalancerMonitorGetResponseType = "icmp_ping"
	LoadBalancerMonitorGetResponseTypeSmtp     LoadBalancerMonitorGetResponseType = "smtp"
)

type LoadBalancerMonitorUpdateResponse struct {
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
	Type LoadBalancerMonitorUpdateResponseType `json:"type"`
	JSON loadBalancerMonitorUpdateResponseJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorUpdateResponse]
type loadBalancerMonitorUpdateResponseJSON struct {
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

func (r *LoadBalancerMonitorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorUpdateResponseType string

const (
	LoadBalancerMonitorUpdateResponseTypeHTTP     LoadBalancerMonitorUpdateResponseType = "http"
	LoadBalancerMonitorUpdateResponseTypeHTTPs    LoadBalancerMonitorUpdateResponseType = "https"
	LoadBalancerMonitorUpdateResponseTypeTcp      LoadBalancerMonitorUpdateResponseType = "tcp"
	LoadBalancerMonitorUpdateResponseTypeUdpIcmp  LoadBalancerMonitorUpdateResponseType = "udp_icmp"
	LoadBalancerMonitorUpdateResponseTypeIcmpPing LoadBalancerMonitorUpdateResponseType = "icmp_ping"
	LoadBalancerMonitorUpdateResponseTypeSmtp     LoadBalancerMonitorUpdateResponseType = "smtp"
)

type LoadBalancerMonitorDeleteResponse struct {
	ID   string                                `json:"id"`
	JSON loadBalancerMonitorDeleteResponseJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorDeleteResponse]
type loadBalancerMonitorDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse struct {
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
	Type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType `json:"type"`
	JSON loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseJSON contains
// the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseJSON struct {
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

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType string

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseTypeHTTP     LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType = "http"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseTypeHTTPs    LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType = "https"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseTypeTcp      LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType = "tcp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseTypeUdpIcmp  LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType = "udp_icmp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseTypeIcmpPing LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType = "icmp_ping"
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseTypeSmtp     LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseType = "smtp"
)

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse struct {
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
	Type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType `json:"type"`
	JSON loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON contains
// the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseJSON struct {
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

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType string

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseTypeHTTP     LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType = "http"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseTypeHTTPs    LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType = "https"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseTypeTcp      LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType = "tcp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseTypeUdpIcmp  LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType = "udp_icmp"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseTypeIcmpPing LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType = "icmp_ping"
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseTypeSmtp     LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseType = "smtp"
)

type LoadBalancerMonitorGetResponseEnvelope struct {
	Errors   []LoadBalancerMonitorGetResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerMonitorGetResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerMonitorGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorGetResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerMonitorGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorGetResponseEnvelope]
type loadBalancerMonitorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerMonitorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorGetResponseEnvelopeErrors]
type loadBalancerMonitorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerMonitorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorGetResponseEnvelopeMessages]
type loadBalancerMonitorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorGetResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorGetResponseEnvelopeSuccessTrue LoadBalancerMonitorGetResponseEnvelopeSuccess = true
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

type LoadBalancerMonitorUpdateResponseEnvelope struct {
	Errors   []LoadBalancerMonitorUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerMonitorUpdateResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerMonitorUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerMonitorUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorUpdateResponseEnvelope]
type loadBalancerMonitorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerMonitorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorUpdateResponseEnvelopeErrors]
type loadBalancerMonitorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    loadBalancerMonitorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorUpdateResponseEnvelopeMessages]
type loadBalancerMonitorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorUpdateResponseEnvelopeSuccessTrue LoadBalancerMonitorUpdateResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorDeleteResponseEnvelope struct {
	Errors   []LoadBalancerMonitorDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerMonitorDeleteResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerMonitorDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerMonitorDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorDeleteResponseEnvelope]
type loadBalancerMonitorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerMonitorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorDeleteResponseEnvelopeErrors]
type loadBalancerMonitorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    loadBalancerMonitorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorDeleteResponseEnvelopeMessages]
type loadBalancerMonitorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorDeleteResponseEnvelopeSuccessTrue LoadBalancerMonitorDeleteResponseEnvelopeSuccess = true
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

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelope struct {
	Errors   []LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelope]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages]
type loadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccessTrue LoadBalancerMonitorAccountLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelope struct {
	Errors     []LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors   `json:"errors"`
	Messages   []LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages `json:"messages"`
	Result     []LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponse                 `json:"result"`
	ResultInfo LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelope]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                  `json:"total_count"`
	JSON       loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo]
type loadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccessTrue LoadBalancerMonitorAccountLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccess = true
)
