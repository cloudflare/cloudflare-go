// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// LoadBalancingMonitorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancingMonitorService]
// method instead.
type LoadBalancingMonitorService struct {
	Options []option.RequestOption
}

// NewLoadBalancingMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancingMonitorService(opts ...option.RequestOption) (r *LoadBalancingMonitorService) {
	r = &LoadBalancingMonitorService{}
	r.Options = opts
	return
}

// Create a configured monitor.
func (r *LoadBalancingMonitorService) New(ctx context.Context, body LoadBalancingMonitorNewParams, opts ...option.RequestOption) (res *LoadBalancingMonitorNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorNewResponseEnvelope
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor.
func (r *LoadBalancingMonitorService) Update(ctx context.Context, monitorID string, body LoadBalancingMonitorUpdateParams, opts ...option.RequestOption) (res *LoadBalancingMonitorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorUpdateResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for a user.
func (r *LoadBalancingMonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[LoadBalancingMonitorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/load_balancers/monitors"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List configured monitors for a user.
func (r *LoadBalancingMonitorService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LoadBalancingMonitorListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// Delete a configured monitor.
func (r *LoadBalancingMonitorService) Delete(ctx context.Context, monitorID string, body LoadBalancingMonitorDeleteParams, opts ...option.RequestOption) (res *LoadBalancingMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorDeleteResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *LoadBalancingMonitorService) Edit(ctx context.Context, monitorID string, body LoadBalancingMonitorEditParams, opts ...option.RequestOption) (res *LoadBalancingMonitorEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorEditResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single configured monitor for a user.
func (r *LoadBalancingMonitorService) Get(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *LoadBalancingMonitorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Preview pools using the specified monitor with provided monitor details. The
// returned preview_id can be used in the preview endpoint to retrieve the results.
func (r *LoadBalancingMonitorService) Preview(ctx context.Context, monitorID string, body LoadBalancingMonitorPreviewParams, opts ...option.RequestOption) (res *LoadBalancingMonitorPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorPreviewResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/preview", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the list of resources that reference the provided monitor.
func (r *LoadBalancingMonitorService) References(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *[]LoadBalancingMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingMonitorReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/references", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingMonitorNewResponse struct {
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
	Type LoadBalancingMonitorNewResponseType `json:"type"`
	JSON loadBalancingMonitorNewResponseJSON `json:"-"`
}

// loadBalancingMonitorNewResponseJSON contains the JSON metadata for the struct
// [LoadBalancingMonitorNewResponse]
type loadBalancingMonitorNewResponseJSON struct {
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

func (r *LoadBalancingMonitorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorNewResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorNewResponseType string

const (
	LoadBalancingMonitorNewResponseTypeHTTP     LoadBalancingMonitorNewResponseType = "http"
	LoadBalancingMonitorNewResponseTypeHTTPS    LoadBalancingMonitorNewResponseType = "https"
	LoadBalancingMonitorNewResponseTypeTCP      LoadBalancingMonitorNewResponseType = "tcp"
	LoadBalancingMonitorNewResponseTypeUdpIcmp  LoadBalancingMonitorNewResponseType = "udp_icmp"
	LoadBalancingMonitorNewResponseTypeIcmpPing LoadBalancingMonitorNewResponseType = "icmp_ping"
	LoadBalancingMonitorNewResponseTypeSmtp     LoadBalancingMonitorNewResponseType = "smtp"
)

func (r LoadBalancingMonitorNewResponseType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorNewResponseTypeHTTP, LoadBalancingMonitorNewResponseTypeHTTPS, LoadBalancingMonitorNewResponseTypeTCP, LoadBalancingMonitorNewResponseTypeUdpIcmp, LoadBalancingMonitorNewResponseTypeIcmpPing, LoadBalancingMonitorNewResponseTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorUpdateResponse struct {
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
	Type LoadBalancingMonitorUpdateResponseType `json:"type"`
	JSON loadBalancingMonitorUpdateResponseJSON `json:"-"`
}

// loadBalancingMonitorUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancingMonitorUpdateResponse]
type loadBalancingMonitorUpdateResponseJSON struct {
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

func (r *LoadBalancingMonitorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorUpdateResponseType string

const (
	LoadBalancingMonitorUpdateResponseTypeHTTP     LoadBalancingMonitorUpdateResponseType = "http"
	LoadBalancingMonitorUpdateResponseTypeHTTPS    LoadBalancingMonitorUpdateResponseType = "https"
	LoadBalancingMonitorUpdateResponseTypeTCP      LoadBalancingMonitorUpdateResponseType = "tcp"
	LoadBalancingMonitorUpdateResponseTypeUdpIcmp  LoadBalancingMonitorUpdateResponseType = "udp_icmp"
	LoadBalancingMonitorUpdateResponseTypeIcmpPing LoadBalancingMonitorUpdateResponseType = "icmp_ping"
	LoadBalancingMonitorUpdateResponseTypeSmtp     LoadBalancingMonitorUpdateResponseType = "smtp"
)

func (r LoadBalancingMonitorUpdateResponseType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorUpdateResponseTypeHTTP, LoadBalancingMonitorUpdateResponseTypeHTTPS, LoadBalancingMonitorUpdateResponseTypeTCP, LoadBalancingMonitorUpdateResponseTypeUdpIcmp, LoadBalancingMonitorUpdateResponseTypeIcmpPing, LoadBalancingMonitorUpdateResponseTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorListResponse struct {
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
	Type LoadBalancingMonitorListResponseType `json:"type"`
	JSON loadBalancingMonitorListResponseJSON `json:"-"`
}

// loadBalancingMonitorListResponseJSON contains the JSON metadata for the struct
// [LoadBalancingMonitorListResponse]
type loadBalancingMonitorListResponseJSON struct {
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

func (r *LoadBalancingMonitorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorListResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorListResponseType string

const (
	LoadBalancingMonitorListResponseTypeHTTP     LoadBalancingMonitorListResponseType = "http"
	LoadBalancingMonitorListResponseTypeHTTPS    LoadBalancingMonitorListResponseType = "https"
	LoadBalancingMonitorListResponseTypeTCP      LoadBalancingMonitorListResponseType = "tcp"
	LoadBalancingMonitorListResponseTypeUdpIcmp  LoadBalancingMonitorListResponseType = "udp_icmp"
	LoadBalancingMonitorListResponseTypeIcmpPing LoadBalancingMonitorListResponseType = "icmp_ping"
	LoadBalancingMonitorListResponseTypeSmtp     LoadBalancingMonitorListResponseType = "smtp"
)

func (r LoadBalancingMonitorListResponseType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorListResponseTypeHTTP, LoadBalancingMonitorListResponseTypeHTTPS, LoadBalancingMonitorListResponseTypeTCP, LoadBalancingMonitorListResponseTypeUdpIcmp, LoadBalancingMonitorListResponseTypeIcmpPing, LoadBalancingMonitorListResponseTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorDeleteResponse struct {
	ID   string                                 `json:"id"`
	JSON loadBalancingMonitorDeleteResponseJSON `json:"-"`
}

// loadBalancingMonitorDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancingMonitorDeleteResponse]
type loadBalancingMonitorDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancingMonitorEditResponse struct {
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
	Type LoadBalancingMonitorEditResponseType `json:"type"`
	JSON loadBalancingMonitorEditResponseJSON `json:"-"`
}

// loadBalancingMonitorEditResponseJSON contains the JSON metadata for the struct
// [LoadBalancingMonitorEditResponse]
type loadBalancingMonitorEditResponseJSON struct {
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

func (r *LoadBalancingMonitorEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorEditResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorEditResponseType string

const (
	LoadBalancingMonitorEditResponseTypeHTTP     LoadBalancingMonitorEditResponseType = "http"
	LoadBalancingMonitorEditResponseTypeHTTPS    LoadBalancingMonitorEditResponseType = "https"
	LoadBalancingMonitorEditResponseTypeTCP      LoadBalancingMonitorEditResponseType = "tcp"
	LoadBalancingMonitorEditResponseTypeUdpIcmp  LoadBalancingMonitorEditResponseType = "udp_icmp"
	LoadBalancingMonitorEditResponseTypeIcmpPing LoadBalancingMonitorEditResponseType = "icmp_ping"
	LoadBalancingMonitorEditResponseTypeSmtp     LoadBalancingMonitorEditResponseType = "smtp"
)

func (r LoadBalancingMonitorEditResponseType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorEditResponseTypeHTTP, LoadBalancingMonitorEditResponseTypeHTTPS, LoadBalancingMonitorEditResponseTypeTCP, LoadBalancingMonitorEditResponseTypeUdpIcmp, LoadBalancingMonitorEditResponseTypeIcmpPing, LoadBalancingMonitorEditResponseTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorGetResponse struct {
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
	Type LoadBalancingMonitorGetResponseType `json:"type"`
	JSON loadBalancingMonitorGetResponseJSON `json:"-"`
}

// loadBalancingMonitorGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancingMonitorGetResponse]
type loadBalancingMonitorGetResponseJSON struct {
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

func (r *LoadBalancingMonitorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorGetResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorGetResponseType string

const (
	LoadBalancingMonitorGetResponseTypeHTTP     LoadBalancingMonitorGetResponseType = "http"
	LoadBalancingMonitorGetResponseTypeHTTPS    LoadBalancingMonitorGetResponseType = "https"
	LoadBalancingMonitorGetResponseTypeTCP      LoadBalancingMonitorGetResponseType = "tcp"
	LoadBalancingMonitorGetResponseTypeUdpIcmp  LoadBalancingMonitorGetResponseType = "udp_icmp"
	LoadBalancingMonitorGetResponseTypeIcmpPing LoadBalancingMonitorGetResponseType = "icmp_ping"
	LoadBalancingMonitorGetResponseTypeSmtp     LoadBalancingMonitorGetResponseType = "smtp"
)

func (r LoadBalancingMonitorGetResponseType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorGetResponseTypeHTTP, LoadBalancingMonitorGetResponseTypeHTTPS, LoadBalancingMonitorGetResponseTypeTCP, LoadBalancingMonitorGetResponseTypeUdpIcmp, LoadBalancingMonitorGetResponseTypeIcmpPing, LoadBalancingMonitorGetResponseTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorPreviewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                       `json:"pools"`
	PreviewID string                                  `json:"preview_id"`
	JSON      loadBalancingMonitorPreviewResponseJSON `json:"-"`
}

// loadBalancingMonitorPreviewResponseJSON contains the JSON metadata for the
// struct [LoadBalancingMonitorPreviewResponse]
type loadBalancingMonitorPreviewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorPreviewResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancingMonitorReferencesResponse struct {
	ReferenceType LoadBalancingMonitorReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                              `json:"resource_id"`
	ResourceName  string                                              `json:"resource_name"`
	ResourceType  string                                              `json:"resource_type"`
	JSON          loadBalancingMonitorReferencesResponseJSON          `json:"-"`
}

// loadBalancingMonitorReferencesResponseJSON contains the JSON metadata for the
// struct [LoadBalancingMonitorReferencesResponse]
type loadBalancingMonitorReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancingMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorReferencesResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancingMonitorReferencesResponseReferenceType string

const (
	LoadBalancingMonitorReferencesResponseReferenceTypeStar     LoadBalancingMonitorReferencesResponseReferenceType = "*"
	LoadBalancingMonitorReferencesResponseReferenceTypeReferral LoadBalancingMonitorReferencesResponseReferenceType = "referral"
	LoadBalancingMonitorReferencesResponseReferenceTypeReferrer LoadBalancingMonitorReferencesResponseReferenceType = "referrer"
)

func (r LoadBalancingMonitorReferencesResponseReferenceType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorReferencesResponseReferenceTypeStar, LoadBalancingMonitorReferencesResponseReferenceTypeReferral, LoadBalancingMonitorReferencesResponseReferenceTypeReferrer:
		return true
	}
	return false
}

type LoadBalancingMonitorNewParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
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
	Type param.Field[LoadBalancingMonitorNewParamsType] `json:"type"`
}

func (r LoadBalancingMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorNewParamsType string

const (
	LoadBalancingMonitorNewParamsTypeHTTP     LoadBalancingMonitorNewParamsType = "http"
	LoadBalancingMonitorNewParamsTypeHTTPS    LoadBalancingMonitorNewParamsType = "https"
	LoadBalancingMonitorNewParamsTypeTCP      LoadBalancingMonitorNewParamsType = "tcp"
	LoadBalancingMonitorNewParamsTypeUdpIcmp  LoadBalancingMonitorNewParamsType = "udp_icmp"
	LoadBalancingMonitorNewParamsTypeIcmpPing LoadBalancingMonitorNewParamsType = "icmp_ping"
	LoadBalancingMonitorNewParamsTypeSmtp     LoadBalancingMonitorNewParamsType = "smtp"
)

func (r LoadBalancingMonitorNewParamsType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorNewParamsTypeHTTP, LoadBalancingMonitorNewParamsTypeHTTPS, LoadBalancingMonitorNewParamsTypeTCP, LoadBalancingMonitorNewParamsTypeUdpIcmp, LoadBalancingMonitorNewParamsTypeIcmpPing, LoadBalancingMonitorNewParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   LoadBalancingMonitorNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingMonitorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingMonitorNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingMonitorNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingMonitorNewResponseEnvelope]
type loadBalancingMonitorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorNewResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorNewResponseEnvelopeSuccessTrue LoadBalancingMonitorNewResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorUpdateParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
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
	Type param.Field[LoadBalancingMonitorUpdateParamsType] `json:"type"`
}

func (r LoadBalancingMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorUpdateParamsType string

const (
	LoadBalancingMonitorUpdateParamsTypeHTTP     LoadBalancingMonitorUpdateParamsType = "http"
	LoadBalancingMonitorUpdateParamsTypeHTTPS    LoadBalancingMonitorUpdateParamsType = "https"
	LoadBalancingMonitorUpdateParamsTypeTCP      LoadBalancingMonitorUpdateParamsType = "tcp"
	LoadBalancingMonitorUpdateParamsTypeUdpIcmp  LoadBalancingMonitorUpdateParamsType = "udp_icmp"
	LoadBalancingMonitorUpdateParamsTypeIcmpPing LoadBalancingMonitorUpdateParamsType = "icmp_ping"
	LoadBalancingMonitorUpdateParamsTypeSmtp     LoadBalancingMonitorUpdateParamsType = "smtp"
)

func (r LoadBalancingMonitorUpdateParamsType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorUpdateParamsTypeHTTP, LoadBalancingMonitorUpdateParamsTypeHTTPS, LoadBalancingMonitorUpdateParamsTypeTCP, LoadBalancingMonitorUpdateParamsTypeUdpIcmp, LoadBalancingMonitorUpdateParamsTypeIcmpPing, LoadBalancingMonitorUpdateParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   LoadBalancingMonitorUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingMonitorUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingMonitorUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingMonitorUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancingMonitorUpdateResponseEnvelope]
type loadBalancingMonitorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorUpdateResponseEnvelopeSuccessTrue LoadBalancingMonitorUpdateResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LoadBalancingMonitorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LoadBalancingMonitorDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   LoadBalancingMonitorDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingMonitorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingMonitorDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingMonitorDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancingMonitorDeleteResponseEnvelope]
type loadBalancingMonitorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorDeleteResponseEnvelopeSuccessTrue LoadBalancingMonitorDeleteResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorEditParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
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
	Type param.Field[LoadBalancingMonitorEditParamsType] `json:"type"`
}

func (r LoadBalancingMonitorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorEditParamsType string

const (
	LoadBalancingMonitorEditParamsTypeHTTP     LoadBalancingMonitorEditParamsType = "http"
	LoadBalancingMonitorEditParamsTypeHTTPS    LoadBalancingMonitorEditParamsType = "https"
	LoadBalancingMonitorEditParamsTypeTCP      LoadBalancingMonitorEditParamsType = "tcp"
	LoadBalancingMonitorEditParamsTypeUdpIcmp  LoadBalancingMonitorEditParamsType = "udp_icmp"
	LoadBalancingMonitorEditParamsTypeIcmpPing LoadBalancingMonitorEditParamsType = "icmp_ping"
	LoadBalancingMonitorEditParamsTypeSmtp     LoadBalancingMonitorEditParamsType = "smtp"
)

func (r LoadBalancingMonitorEditParamsType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorEditParamsTypeHTTP, LoadBalancingMonitorEditParamsTypeHTTPS, LoadBalancingMonitorEditParamsTypeTCP, LoadBalancingMonitorEditParamsTypeUdpIcmp, LoadBalancingMonitorEditParamsTypeIcmpPing, LoadBalancingMonitorEditParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   LoadBalancingMonitorEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingMonitorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingMonitorEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingMonitorEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingMonitorEditResponseEnvelope]
type loadBalancingMonitorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorEditResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorEditResponseEnvelopeSuccessTrue LoadBalancingMonitorEditResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   LoadBalancingMonitorGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingMonitorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingMonitorGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingMonitorGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingMonitorGetResponseEnvelope]
type loadBalancingMonitorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorGetResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorGetResponseEnvelopeSuccessTrue LoadBalancingMonitorGetResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorPreviewParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
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
	Type param.Field[LoadBalancingMonitorPreviewParamsType] `json:"type"`
}

func (r LoadBalancingMonitorPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorPreviewParamsType string

const (
	LoadBalancingMonitorPreviewParamsTypeHTTP     LoadBalancingMonitorPreviewParamsType = "http"
	LoadBalancingMonitorPreviewParamsTypeHTTPS    LoadBalancingMonitorPreviewParamsType = "https"
	LoadBalancingMonitorPreviewParamsTypeTCP      LoadBalancingMonitorPreviewParamsType = "tcp"
	LoadBalancingMonitorPreviewParamsTypeUdpIcmp  LoadBalancingMonitorPreviewParamsType = "udp_icmp"
	LoadBalancingMonitorPreviewParamsTypeIcmpPing LoadBalancingMonitorPreviewParamsType = "icmp_ping"
	LoadBalancingMonitorPreviewParamsTypeSmtp     LoadBalancingMonitorPreviewParamsType = "smtp"
)

func (r LoadBalancingMonitorPreviewParamsType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorPreviewParamsTypeHTTP, LoadBalancingMonitorPreviewParamsTypeHTTPS, LoadBalancingMonitorPreviewParamsTypeTCP, LoadBalancingMonitorPreviewParamsTypeUdpIcmp, LoadBalancingMonitorPreviewParamsTypeIcmpPing, LoadBalancingMonitorPreviewParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingMonitorPreviewResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors,required"`
	Messages []shared.ResponseInfo               `json:"messages,required"`
	Result   LoadBalancingMonitorPreviewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingMonitorPreviewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingMonitorPreviewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingMonitorPreviewResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancingMonitorPreviewResponseEnvelope]
type loadBalancingMonitorPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorPreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorPreviewResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorPreviewResponseEnvelopeSuccessTrue LoadBalancingMonitorPreviewResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorPreviewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorPreviewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorReferencesResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []LoadBalancingMonitorReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancingMonitorReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancingMonitorReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancingMonitorReferencesResponseEnvelopeJSON       `json:"-"`
}

// loadBalancingMonitorReferencesResponseEnvelopeJSON contains the JSON metadata
// for the struct [LoadBalancingMonitorReferencesResponseEnvelope]
type loadBalancingMonitorReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorReferencesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingMonitorReferencesResponseEnvelopeSuccess bool

const (
	LoadBalancingMonitorReferencesResponseEnvelopeSuccessTrue LoadBalancingMonitorReferencesResponseEnvelopeSuccess = true
)

func (r LoadBalancingMonitorReferencesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorReferencesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMonitorReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       loadBalancingMonitorReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancingMonitorReferencesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [LoadBalancingMonitorReferencesResponseEnvelopeResultInfo]
type loadBalancingMonitorReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMonitorReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorReferencesResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
