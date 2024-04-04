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

// LoadBalancerMonitorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerMonitorService]
// method instead.
type LoadBalancerMonitorService struct {
	Options []option.RequestOption
}

// NewLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerMonitorService(opts ...option.RequestOption) (r *LoadBalancerMonitorService) {
	r = &LoadBalancerMonitorService{}
	r.Options = opts
	return
}

// Create a configured monitor.
func (r *LoadBalancerMonitorService) New(ctx context.Context, body LoadBalancerMonitorNewParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorNewResponseEnvelope
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor.
func (r *LoadBalancerMonitorService) Update(ctx context.Context, monitorID string, body LoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorUpdateResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for a user.
func (r *LoadBalancerMonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[LoadBalancingMonitor], err error) {
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
func (r *LoadBalancerMonitorService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LoadBalancingMonitor] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// Delete a configured monitor.
func (r *LoadBalancerMonitorService) Delete(ctx context.Context, monitorID string, body LoadBalancerMonitorDeleteParams, opts ...option.RequestOption) (res *LoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorDeleteResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *LoadBalancerMonitorService) Edit(ctx context.Context, monitorID string, body LoadBalancerMonitorEditParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorEditResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single configured monitor for a user.
func (r *LoadBalancerMonitorService) Get(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorGetResponseEnvelope
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
func (r *LoadBalancerMonitorService) Preview(ctx context.Context, monitorID string, body LoadBalancerMonitorPreviewParams, opts ...option.RequestOption) (res *LoadBalancerMonitorPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorPreviewResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/preview", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the list of resources that reference the provided monitor.
func (r *LoadBalancerMonitorService) References(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *[]LoadBalancerMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/references", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingMonitor struct {
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
	Type LoadBalancingMonitorType `json:"type"`
	JSON loadBalancingMonitorJSON `json:"-"`
}

// loadBalancingMonitorJSON contains the JSON metadata for the struct
// [LoadBalancingMonitor]
type loadBalancingMonitorJSON struct {
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

func (r *LoadBalancingMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMonitorJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorType string

const (
	LoadBalancingMonitorTypeHTTP     LoadBalancingMonitorType = "http"
	LoadBalancingMonitorTypeHTTPS    LoadBalancingMonitorType = "https"
	LoadBalancingMonitorTypeTcp      LoadBalancingMonitorType = "tcp"
	LoadBalancingMonitorTypeUdpIcmp  LoadBalancingMonitorType = "udp_icmp"
	LoadBalancingMonitorTypeIcmpPing LoadBalancingMonitorType = "icmp_ping"
	LoadBalancingMonitorTypeSmtp     LoadBalancingMonitorType = "smtp"
)

func (r LoadBalancingMonitorType) IsKnown() bool {
	switch r {
	case LoadBalancingMonitorTypeHTTP, LoadBalancingMonitorTypeHTTPS, LoadBalancingMonitorTypeTcp, LoadBalancingMonitorTypeUdpIcmp, LoadBalancingMonitorTypeIcmpPing, LoadBalancingMonitorTypeSmtp:
		return true
	}
	return false
}

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

func (r loadBalancerMonitorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerMonitorPreviewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                      `json:"pools"`
	PreviewID string                                 `json:"preview_id"`
	JSON      loadBalancerMonitorPreviewResponseJSON `json:"-"`
}

// loadBalancerMonitorPreviewResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorPreviewResponse]
type loadBalancerMonitorPreviewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorPreviewResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerMonitorReferencesResponse struct {
	ReferenceType shared.UnnamedSchemaRef146                `json:"reference_type"`
	ResourceID    string                                    `json:"resource_id"`
	ResourceName  string                                    `json:"resource_name"`
	ResourceType  string                                    `json:"resource_type"`
	JSON          loadBalancerMonitorReferencesResponseJSON `json:"-"`
}

// loadBalancerMonitorReferencesResponseJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorReferencesResponse]
type loadBalancerMonitorReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferencesResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerMonitorNewParams struct {
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
	Type param.Field[LoadBalancerMonitorNewParamsType] `json:"type"`
}

func (r LoadBalancerMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorNewParamsType string

const (
	LoadBalancerMonitorNewParamsTypeHTTP     LoadBalancerMonitorNewParamsType = "http"
	LoadBalancerMonitorNewParamsTypeHTTPS    LoadBalancerMonitorNewParamsType = "https"
	LoadBalancerMonitorNewParamsTypeTcp      LoadBalancerMonitorNewParamsType = "tcp"
	LoadBalancerMonitorNewParamsTypeUdpIcmp  LoadBalancerMonitorNewParamsType = "udp_icmp"
	LoadBalancerMonitorNewParamsTypeIcmpPing LoadBalancerMonitorNewParamsType = "icmp_ping"
	LoadBalancerMonitorNewParamsTypeSmtp     LoadBalancerMonitorNewParamsType = "smtp"
)

func (r LoadBalancerMonitorNewParamsType) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorNewParamsTypeHTTP, LoadBalancerMonitorNewParamsTypeHTTPS, LoadBalancerMonitorNewParamsTypeTcp, LoadBalancerMonitorNewParamsTypeUdpIcmp, LoadBalancerMonitorNewParamsTypeIcmpPing, LoadBalancerMonitorNewParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancerMonitorNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancingMonitor  `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorNewResponseEnvelope]
type loadBalancerMonitorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorNewResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorNewResponseEnvelopeSuccessTrue LoadBalancerMonitorNewResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorUpdateParams struct {
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
	LoadBalancerMonitorUpdateParamsTypeHTTPS    LoadBalancerMonitorUpdateParamsType = "https"
	LoadBalancerMonitorUpdateParamsTypeTcp      LoadBalancerMonitorUpdateParamsType = "tcp"
	LoadBalancerMonitorUpdateParamsTypeUdpIcmp  LoadBalancerMonitorUpdateParamsType = "udp_icmp"
	LoadBalancerMonitorUpdateParamsTypeIcmpPing LoadBalancerMonitorUpdateParamsType = "icmp_ping"
	LoadBalancerMonitorUpdateParamsTypeSmtp     LoadBalancerMonitorUpdateParamsType = "smtp"
)

func (r LoadBalancerMonitorUpdateParamsType) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorUpdateParamsTypeHTTP, LoadBalancerMonitorUpdateParamsTypeHTTPS, LoadBalancerMonitorUpdateParamsTypeTcp, LoadBalancerMonitorUpdateParamsTypeUdpIcmp, LoadBalancerMonitorUpdateParamsTypeIcmpPing, LoadBalancerMonitorUpdateParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancerMonitorUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancingMonitor  `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorUpdateResponseEnvelopeSuccess `json:"success,required"`
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

func (r loadBalancerMonitorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorUpdateResponseEnvelopeSuccessTrue LoadBalancerMonitorUpdateResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LoadBalancerMonitorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LoadBalancerMonitorDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo             `json:"errors,required"`
	Messages []shared.ResponseInfo             `json:"messages,required"`
	Result   LoadBalancerMonitorDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorDeleteResponseEnvelopeSuccess `json:"success,required"`
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

func (r loadBalancerMonitorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorDeleteResponseEnvelopeSuccessTrue LoadBalancerMonitorDeleteResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorEditParams struct {
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
	Type param.Field[LoadBalancerMonitorEditParamsType] `json:"type"`
}

func (r LoadBalancerMonitorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorEditParamsType string

const (
	LoadBalancerMonitorEditParamsTypeHTTP     LoadBalancerMonitorEditParamsType = "http"
	LoadBalancerMonitorEditParamsTypeHTTPS    LoadBalancerMonitorEditParamsType = "https"
	LoadBalancerMonitorEditParamsTypeTcp      LoadBalancerMonitorEditParamsType = "tcp"
	LoadBalancerMonitorEditParamsTypeUdpIcmp  LoadBalancerMonitorEditParamsType = "udp_icmp"
	LoadBalancerMonitorEditParamsTypeIcmpPing LoadBalancerMonitorEditParamsType = "icmp_ping"
	LoadBalancerMonitorEditParamsTypeSmtp     LoadBalancerMonitorEditParamsType = "smtp"
)

func (r LoadBalancerMonitorEditParamsType) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorEditParamsTypeHTTP, LoadBalancerMonitorEditParamsTypeHTTPS, LoadBalancerMonitorEditParamsTypeTcp, LoadBalancerMonitorEditParamsTypeUdpIcmp, LoadBalancerMonitorEditParamsTypeIcmpPing, LoadBalancerMonitorEditParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancerMonitorEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancingMonitor  `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorEditResponseEnvelope]
type loadBalancerMonitorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorEditResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorEditResponseEnvelopeSuccessTrue LoadBalancerMonitorEditResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancingMonitor  `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorGetResponseEnvelopeSuccess `json:"success,required"`
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

func (r loadBalancerMonitorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorGetResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorGetResponseEnvelopeSuccessTrue LoadBalancerMonitorGetResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorPreviewParams struct {
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
	Type param.Field[LoadBalancerMonitorPreviewParamsType] `json:"type"`
}

func (r LoadBalancerMonitorPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorPreviewParamsType string

const (
	LoadBalancerMonitorPreviewParamsTypeHTTP     LoadBalancerMonitorPreviewParamsType = "http"
	LoadBalancerMonitorPreviewParamsTypeHTTPS    LoadBalancerMonitorPreviewParamsType = "https"
	LoadBalancerMonitorPreviewParamsTypeTcp      LoadBalancerMonitorPreviewParamsType = "tcp"
	LoadBalancerMonitorPreviewParamsTypeUdpIcmp  LoadBalancerMonitorPreviewParamsType = "udp_icmp"
	LoadBalancerMonitorPreviewParamsTypeIcmpPing LoadBalancerMonitorPreviewParamsType = "icmp_ping"
	LoadBalancerMonitorPreviewParamsTypeSmtp     LoadBalancerMonitorPreviewParamsType = "smtp"
)

func (r LoadBalancerMonitorPreviewParamsType) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorPreviewParamsTypeHTTP, LoadBalancerMonitorPreviewParamsTypeHTTPS, LoadBalancerMonitorPreviewParamsTypeTcp, LoadBalancerMonitorPreviewParamsTypeUdpIcmp, LoadBalancerMonitorPreviewParamsTypeIcmpPing, LoadBalancerMonitorPreviewParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancerMonitorPreviewResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   LoadBalancerMonitorPreviewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorPreviewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorPreviewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorPreviewResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorPreviewResponseEnvelope]
type loadBalancerMonitorPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorPreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorPreviewResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorPreviewResponseEnvelopeSuccessTrue LoadBalancerMonitorPreviewResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorPreviewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorPreviewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorReferencesResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []LoadBalancerMonitorReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerMonitorReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerMonitorReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerMonitorReferencesResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerMonitorReferencesResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorReferencesResponseEnvelope]
type loadBalancerMonitorReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferencesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorReferencesResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorReferencesResponseEnvelopeSuccessTrue LoadBalancerMonitorReferencesResponseEnvelopeSuccess = true
)

func (r LoadBalancerMonitorReferencesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerMonitorReferencesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerMonitorReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       loadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [LoadBalancerMonitorReferencesResponseEnvelopeResultInfo]
type loadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
