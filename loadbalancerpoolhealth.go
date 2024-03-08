// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// LoadBalancerPoolHealthService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerPoolHealthService]
// method instead.
type LoadBalancerPoolHealthService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPoolHealthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolHealthService(opts ...option.RequestOption) (r *LoadBalancerPoolHealthService) {
	r = &LoadBalancerPoolHealthService{}
	r.Options = opts
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *LoadBalancerPoolHealthService) New(ctx context.Context, poolID string, params LoadBalancerPoolHealthNewParams, opts ...option.RequestOption) (res *LoadBalancerPoolHealthNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolHealthNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/preview", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the latest pool health status for a single pool.
func (r *LoadBalancerPoolHealthService) Get(ctx context.Context, poolID string, query LoadBalancerPoolHealthGetParams, opts ...option.RequestOption) (res *LoadBalancerPoolHealthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolHealthGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/health", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerPoolHealthNewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                     `json:"pools"`
	PreviewID string                                `json:"preview_id"`
	JSON      loadBalancerPoolHealthNewResponseJSON `json:"-"`
}

// loadBalancerPoolHealthNewResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolHealthNewResponse]
type loadBalancerPoolHealthNewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthNewResponseJSON) RawJSON() string {
	return r.raw
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
//
// Union satisfied by [LoadBalancerPoolHealthGetResponseUnknown] or
// [shared.UnionString].
type LoadBalancerPoolHealthGetResponse interface {
	ImplementsLoadBalancerPoolHealthGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancerPoolHealthGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LoadBalancerPoolHealthNewParams struct {
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
	Type param.Field[LoadBalancerPoolHealthNewParamsType] `json:"type"`
}

func (r LoadBalancerPoolHealthNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerPoolHealthNewParamsType string

const (
	LoadBalancerPoolHealthNewParamsTypeHTTP     LoadBalancerPoolHealthNewParamsType = "http"
	LoadBalancerPoolHealthNewParamsTypeHTTPS    LoadBalancerPoolHealthNewParamsType = "https"
	LoadBalancerPoolHealthNewParamsTypeTcp      LoadBalancerPoolHealthNewParamsType = "tcp"
	LoadBalancerPoolHealthNewParamsTypeUdpIcmp  LoadBalancerPoolHealthNewParamsType = "udp_icmp"
	LoadBalancerPoolHealthNewParamsTypeIcmpPing LoadBalancerPoolHealthNewParamsType = "icmp_ping"
	LoadBalancerPoolHealthNewParamsTypeSmtp     LoadBalancerPoolHealthNewParamsType = "smtp"
)

type LoadBalancerPoolHealthNewResponseEnvelope struct {
	Errors   []LoadBalancerPoolHealthNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolHealthNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerPoolHealthNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolHealthNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolHealthNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolHealthNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolHealthNewResponseEnvelope]
type loadBalancerPoolHealthNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolHealthNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerPoolHealthNewResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolHealthNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LoadBalancerPoolHealthNewResponseEnvelopeErrors]
type loadBalancerPoolHealthNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolHealthNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    loadBalancerPoolHealthNewResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolHealthNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerPoolHealthNewResponseEnvelopeMessages]
type loadBalancerPoolHealthNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolHealthNewResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolHealthNewResponseEnvelopeSuccessTrue LoadBalancerPoolHealthNewResponseEnvelopeSuccess = true
)

type LoadBalancerPoolHealthGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerPoolHealthGetResponseEnvelope struct {
	Errors   []LoadBalancerPoolHealthGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolHealthGetResponseEnvelopeMessages `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result LoadBalancerPoolHealthGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolHealthGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolHealthGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolHealthGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolHealthGetResponseEnvelope]
type loadBalancerPoolHealthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolHealthGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerPoolHealthGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolHealthGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LoadBalancerPoolHealthGetResponseEnvelopeErrors]
type loadBalancerPoolHealthGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolHealthGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    loadBalancerPoolHealthGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolHealthGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerPoolHealthGetResponseEnvelopeMessages]
type loadBalancerPoolHealthGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolHealthGetResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolHealthGetResponseEnvelopeSuccessTrue LoadBalancerPoolHealthGetResponseEnvelopeSuccess = true
)
