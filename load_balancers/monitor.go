// File generated from our OpenAPI spec by Stainless.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/user"
)

// MonitorService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMonitorService] method instead.
type MonitorService struct {
	Options    []option.RequestOption
	Previews   *MonitorPreviewService
	References *MonitorReferenceService
}

// NewMonitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMonitorService(opts ...option.RequestOption) (r *MonitorService) {
	r = &MonitorService{}
	r.Options = opts
	r.Previews = NewMonitorPreviewService(opts...)
	r.References = NewMonitorReferenceService(opts...)
	return
}

// Create a configured monitor.
func (r *MonitorService) New(ctx context.Context, params MonitorNewParams, opts ...option.RequestOption) (res *user.LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor.
func (r *MonitorService) Update(ctx context.Context, monitorID string, params MonitorUpdateParams, opts ...option.RequestOption) (res *user.LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", params.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for an account.
func (r *MonitorService) List(ctx context.Context, query MonitorListParams, opts ...option.RequestOption) (res *[]user.LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured monitor.
func (r *MonitorService) Delete(ctx context.Context, monitorID string, body MonitorDeleteParams, opts ...option.RequestOption) (res *MonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", body.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *MonitorService) Edit(ctx context.Context, monitorID string, params MonitorEditParams, opts ...option.RequestOption) (res *user.LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", params.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single configured monitor for an account.
func (r *MonitorService) Get(ctx context.Context, monitorID string, query MonitorGetParams, opts ...option.RequestOption) (res *user.LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", query.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MonitorDeleteResponse struct {
	ID   string                    `json:"id"`
	JSON monitorDeleteResponseJSON `json:"-"`
}

// monitorDeleteResponseJSON contains the JSON metadata for the struct
// [MonitorDeleteResponse]
type monitorDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Type param.Field[MonitorNewParamsType] `json:"type"`
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorNewParamsType string

const (
	MonitorNewParamsTypeHTTP     MonitorNewParamsType = "http"
	MonitorNewParamsTypeHTTPS    MonitorNewParamsType = "https"
	MonitorNewParamsTypeTcp      MonitorNewParamsType = "tcp"
	MonitorNewParamsTypeUdpIcmp  MonitorNewParamsType = "udp_icmp"
	MonitorNewParamsTypeIcmpPing MonitorNewParamsType = "icmp_ping"
	MonitorNewParamsTypeSmtp     MonitorNewParamsType = "smtp"
)

type MonitorNewResponseEnvelope struct {
	Errors   []MonitorNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MonitorNewResponseEnvelopeMessages `json:"messages,required"`
	Result   user.LoadBalancingMonitor            `json:"result,required"`
	// Whether the API call was successful
	Success MonitorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorNewResponseEnvelopeJSON    `json:"-"`
}

// monitorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorNewResponseEnvelope]
type monitorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MonitorNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    monitorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// monitorNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MonitorNewResponseEnvelopeErrors]
type monitorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MonitorNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    monitorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// monitorNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MonitorNewResponseEnvelopeMessages]
type monitorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorNewResponseEnvelopeSuccess bool

const (
	MonitorNewResponseEnvelopeSuccessTrue MonitorNewResponseEnvelopeSuccess = true
)

type MonitorUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Type param.Field[MonitorUpdateParamsType] `json:"type"`
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorUpdateParamsType string

const (
	MonitorUpdateParamsTypeHTTP     MonitorUpdateParamsType = "http"
	MonitorUpdateParamsTypeHTTPS    MonitorUpdateParamsType = "https"
	MonitorUpdateParamsTypeTcp      MonitorUpdateParamsType = "tcp"
	MonitorUpdateParamsTypeUdpIcmp  MonitorUpdateParamsType = "udp_icmp"
	MonitorUpdateParamsTypeIcmpPing MonitorUpdateParamsType = "icmp_ping"
	MonitorUpdateParamsTypeSmtp     MonitorUpdateParamsType = "smtp"
)

type MonitorUpdateResponseEnvelope struct {
	Errors   []MonitorUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MonitorUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   user.LoadBalancingMonitor               `json:"result,required"`
	// Whether the API call was successful
	Success MonitorUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorUpdateResponseEnvelopeJSON    `json:"-"`
}

// monitorUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorUpdateResponseEnvelope]
type monitorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MonitorUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    monitorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// monitorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MonitorUpdateResponseEnvelopeErrors]
type monitorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MonitorUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    monitorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// monitorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MonitorUpdateResponseEnvelopeMessages]
type monitorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorUpdateResponseEnvelopeSuccess bool

const (
	MonitorUpdateResponseEnvelopeSuccessTrue MonitorUpdateResponseEnvelopeSuccess = true
)

type MonitorListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorListResponseEnvelope struct {
	Errors   []MonitorListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MonitorListResponseEnvelopeMessages `json:"messages,required"`
	Result   []user.LoadBalancingMonitor           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MonitorListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MonitorListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       monitorListResponseEnvelopeJSON       `json:"-"`
}

// monitorListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorListResponseEnvelope]
type monitorListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MonitorListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    monitorListResponseEnvelopeErrorsJSON `json:"-"`
}

// monitorListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MonitorListResponseEnvelopeErrors]
type monitorListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MonitorListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    monitorListResponseEnvelopeMessagesJSON `json:"-"`
}

// monitorListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MonitorListResponseEnvelopeMessages]
type monitorListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorListResponseEnvelopeSuccess bool

const (
	MonitorListResponseEnvelopeSuccessTrue MonitorListResponseEnvelopeSuccess = true
)

type MonitorListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       monitorListResponseEnvelopeResultInfoJSON `json:"-"`
}

// monitorListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [MonitorListResponseEnvelopeResultInfo]
type monitorListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type MonitorDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorDeleteResponseEnvelope struct {
	Errors   []MonitorDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MonitorDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MonitorDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MonitorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorDeleteResponseEnvelopeJSON    `json:"-"`
}

// monitorDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorDeleteResponseEnvelope]
type monitorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MonitorDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    monitorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// monitorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MonitorDeleteResponseEnvelopeErrors]
type monitorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MonitorDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    monitorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// monitorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MonitorDeleteResponseEnvelopeMessages]
type monitorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorDeleteResponseEnvelopeSuccess bool

const (
	MonitorDeleteResponseEnvelopeSuccessTrue MonitorDeleteResponseEnvelopeSuccess = true
)

type MonitorEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Type param.Field[MonitorEditParamsType] `json:"type"`
}

func (r MonitorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorEditParamsType string

const (
	MonitorEditParamsTypeHTTP     MonitorEditParamsType = "http"
	MonitorEditParamsTypeHTTPS    MonitorEditParamsType = "https"
	MonitorEditParamsTypeTcp      MonitorEditParamsType = "tcp"
	MonitorEditParamsTypeUdpIcmp  MonitorEditParamsType = "udp_icmp"
	MonitorEditParamsTypeIcmpPing MonitorEditParamsType = "icmp_ping"
	MonitorEditParamsTypeSmtp     MonitorEditParamsType = "smtp"
)

type MonitorEditResponseEnvelope struct {
	Errors   []MonitorEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MonitorEditResponseEnvelopeMessages `json:"messages,required"`
	Result   user.LoadBalancingMonitor             `json:"result,required"`
	// Whether the API call was successful
	Success MonitorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorEditResponseEnvelopeJSON    `json:"-"`
}

// monitorEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorEditResponseEnvelope]
type monitorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MonitorEditResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    monitorEditResponseEnvelopeErrorsJSON `json:"-"`
}

// monitorEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MonitorEditResponseEnvelopeErrors]
type monitorEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MonitorEditResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    monitorEditResponseEnvelopeMessagesJSON `json:"-"`
}

// monitorEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MonitorEditResponseEnvelopeMessages]
type monitorEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorEditResponseEnvelopeSuccess bool

const (
	MonitorEditResponseEnvelopeSuccessTrue MonitorEditResponseEnvelopeSuccess = true
)

type MonitorGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorGetResponseEnvelope struct {
	Errors   []MonitorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MonitorGetResponseEnvelopeMessages `json:"messages,required"`
	Result   user.LoadBalancingMonitor            `json:"result,required"`
	// Whether the API call was successful
	Success MonitorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorGetResponseEnvelopeJSON    `json:"-"`
}

// monitorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorGetResponseEnvelope]
type monitorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MonitorGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    monitorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// monitorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MonitorGetResponseEnvelopeErrors]
type monitorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MonitorGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    monitorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// monitorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MonitorGetResponseEnvelopeMessages]
type monitorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorGetResponseEnvelopeSuccess bool

const (
	MonitorGetResponseEnvelopeSuccessTrue MonitorGetResponseEnvelopeSuccess = true
)
