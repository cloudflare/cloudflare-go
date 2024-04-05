// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// PoolHealthService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPoolHealthService] method instead.
type PoolHealthService struct {
	Options []option.RequestOption
}

// NewPoolHealthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPoolHealthService(opts ...option.RequestOption) (r *PoolHealthService) {
	r = &PoolHealthService{}
	r.Options = opts
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *PoolHealthService) New(ctx context.Context, poolID string, params PoolHealthNewParams, opts ...option.RequestOption) (res *PoolHealthNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolHealthNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/preview", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the latest pool health status for a single pool.
func (r *PoolHealthService) Get(ctx context.Context, poolID string, query PoolHealthGetParams, opts ...option.RequestOption) (res *PoolHealthGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolHealthGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/health", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PoolHealthNewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string         `json:"pools"`
	PreviewID string                    `json:"preview_id"`
	JSON      poolHealthNewResponseJSON `json:"-"`
}

// poolHealthNewResponseJSON contains the JSON metadata for the struct
// [PoolHealthNewResponse]
type poolHealthNewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolHealthNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolHealthNewResponseJSON) RawJSON() string {
	return r.raw
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
//
// Union satisfied by [load_balancers.PoolHealthGetResponseUnknown] or
// [shared.UnionString].
type PoolHealthGetResponseUnion interface {
	ImplementsLoadBalancersPoolHealthGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PoolHealthGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PoolHealthNewParams struct {
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
	Type param.Field[PoolHealthNewParamsType] `json:"type"`
}

func (r PoolHealthNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type PoolHealthNewParamsType string

const (
	PoolHealthNewParamsTypeHTTP     PoolHealthNewParamsType = "http"
	PoolHealthNewParamsTypeHTTPS    PoolHealthNewParamsType = "https"
	PoolHealthNewParamsTypeTCP      PoolHealthNewParamsType = "tcp"
	PoolHealthNewParamsTypeUdpIcmp  PoolHealthNewParamsType = "udp_icmp"
	PoolHealthNewParamsTypeIcmpPing PoolHealthNewParamsType = "icmp_ping"
	PoolHealthNewParamsTypeSmtp     PoolHealthNewParamsType = "smtp"
)

func (r PoolHealthNewParamsType) IsKnown() bool {
	switch r {
	case PoolHealthNewParamsTypeHTTP, PoolHealthNewParamsTypeHTTPS, PoolHealthNewParamsTypeTCP, PoolHealthNewParamsTypeUdpIcmp, PoolHealthNewParamsTypeIcmpPing, PoolHealthNewParamsTypeSmtp:
		return true
	}
	return false
}

type PoolHealthNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   PoolHealthNewResponse                                     `json:"result,required"`
	// Whether the API call was successful
	Success PoolHealthNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolHealthNewResponseEnvelopeJSON    `json:"-"`
}

// poolHealthNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolHealthNewResponseEnvelope]
type poolHealthNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolHealthNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolHealthNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolHealthNewResponseEnvelopeSuccess bool

const (
	PoolHealthNewResponseEnvelopeSuccessTrue PoolHealthNewResponseEnvelopeSuccess = true
)

func (r PoolHealthNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolHealthNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PoolHealthGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PoolHealthGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result PoolHealthGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success PoolHealthGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolHealthGetResponseEnvelopeJSON    `json:"-"`
}

// poolHealthGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolHealthGetResponseEnvelope]
type poolHealthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolHealthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolHealthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolHealthGetResponseEnvelopeSuccess bool

const (
	PoolHealthGetResponseEnvelopeSuccessTrue PoolHealthGetResponseEnvelopeSuccess = true
)

func (r PoolHealthGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolHealthGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
