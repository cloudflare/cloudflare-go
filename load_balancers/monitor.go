// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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
func (r *MonitorService) New(ctx context.Context, params MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponse, err error) {
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
func (r *MonitorService) Update(ctx context.Context, monitorID string, params MonitorUpdateParams, opts ...option.RequestOption) (res *MonitorUpdateResponse, err error) {
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
func (r *MonitorService) List(ctx context.Context, query MonitorListParams, opts ...option.RequestOption) (res *pagination.SinglePage[MonitorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List configured monitors for an account.
func (r *MonitorService) ListAutoPaging(ctx context.Context, query MonitorListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[MonitorListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured monitor.
func (r *MonitorService) Delete(ctx context.Context, monitorID string, params MonitorDeleteParams, opts ...option.RequestOption) (res *MonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", params.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *MonitorService) Edit(ctx context.Context, monitorID string, params MonitorEditParams, opts ...option.RequestOption) (res *MonitorEditResponse, err error) {
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
func (r *MonitorService) Get(ctx context.Context, monitorID string, query MonitorGetParams, opts ...option.RequestOption) (res *MonitorGetResponse, err error) {
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

type MonitorNewResponse struct {
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
	Type MonitorNewResponseType `json:"type"`
	JSON monitorNewResponseJSON `json:"-"`
}

// monitorNewResponseJSON contains the JSON metadata for the struct
// [MonitorNewResponse]
type monitorNewResponseJSON struct {
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

func (r *MonitorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorNewResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorNewResponseType string

const (
	MonitorNewResponseTypeHTTP     MonitorNewResponseType = "http"
	MonitorNewResponseTypeHTTPS    MonitorNewResponseType = "https"
	MonitorNewResponseTypeTCP      MonitorNewResponseType = "tcp"
	MonitorNewResponseTypeUdpIcmp  MonitorNewResponseType = "udp_icmp"
	MonitorNewResponseTypeIcmpPing MonitorNewResponseType = "icmp_ping"
	MonitorNewResponseTypeSmtp     MonitorNewResponseType = "smtp"
)

func (r MonitorNewResponseType) IsKnown() bool {
	switch r {
	case MonitorNewResponseTypeHTTP, MonitorNewResponseTypeHTTPS, MonitorNewResponseTypeTCP, MonitorNewResponseTypeUdpIcmp, MonitorNewResponseTypeIcmpPing, MonitorNewResponseTypeSmtp:
		return true
	}
	return false
}

type MonitorUpdateResponse struct {
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
	Type MonitorUpdateResponseType `json:"type"`
	JSON monitorUpdateResponseJSON `json:"-"`
}

// monitorUpdateResponseJSON contains the JSON metadata for the struct
// [MonitorUpdateResponse]
type monitorUpdateResponseJSON struct {
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

func (r *MonitorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorUpdateResponseType string

const (
	MonitorUpdateResponseTypeHTTP     MonitorUpdateResponseType = "http"
	MonitorUpdateResponseTypeHTTPS    MonitorUpdateResponseType = "https"
	MonitorUpdateResponseTypeTCP      MonitorUpdateResponseType = "tcp"
	MonitorUpdateResponseTypeUdpIcmp  MonitorUpdateResponseType = "udp_icmp"
	MonitorUpdateResponseTypeIcmpPing MonitorUpdateResponseType = "icmp_ping"
	MonitorUpdateResponseTypeSmtp     MonitorUpdateResponseType = "smtp"
)

func (r MonitorUpdateResponseType) IsKnown() bool {
	switch r {
	case MonitorUpdateResponseTypeHTTP, MonitorUpdateResponseTypeHTTPS, MonitorUpdateResponseTypeTCP, MonitorUpdateResponseTypeUdpIcmp, MonitorUpdateResponseTypeIcmpPing, MonitorUpdateResponseTypeSmtp:
		return true
	}
	return false
}

type MonitorListResponse struct {
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
	Type MonitorListResponseType `json:"type"`
	JSON monitorListResponseJSON `json:"-"`
}

// monitorListResponseJSON contains the JSON metadata for the struct
// [MonitorListResponse]
type monitorListResponseJSON struct {
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

func (r *MonitorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorListResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorListResponseType string

const (
	MonitorListResponseTypeHTTP     MonitorListResponseType = "http"
	MonitorListResponseTypeHTTPS    MonitorListResponseType = "https"
	MonitorListResponseTypeTCP      MonitorListResponseType = "tcp"
	MonitorListResponseTypeUdpIcmp  MonitorListResponseType = "udp_icmp"
	MonitorListResponseTypeIcmpPing MonitorListResponseType = "icmp_ping"
	MonitorListResponseTypeSmtp     MonitorListResponseType = "smtp"
)

func (r MonitorListResponseType) IsKnown() bool {
	switch r {
	case MonitorListResponseTypeHTTP, MonitorListResponseTypeHTTPS, MonitorListResponseTypeTCP, MonitorListResponseTypeUdpIcmp, MonitorListResponseTypeIcmpPing, MonitorListResponseTypeSmtp:
		return true
	}
	return false
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

type MonitorEditResponse struct {
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
	Type MonitorEditResponseType `json:"type"`
	JSON monitorEditResponseJSON `json:"-"`
}

// monitorEditResponseJSON contains the JSON metadata for the struct
// [MonitorEditResponse]
type monitorEditResponseJSON struct {
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

func (r *MonitorEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEditResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorEditResponseType string

const (
	MonitorEditResponseTypeHTTP     MonitorEditResponseType = "http"
	MonitorEditResponseTypeHTTPS    MonitorEditResponseType = "https"
	MonitorEditResponseTypeTCP      MonitorEditResponseType = "tcp"
	MonitorEditResponseTypeUdpIcmp  MonitorEditResponseType = "udp_icmp"
	MonitorEditResponseTypeIcmpPing MonitorEditResponseType = "icmp_ping"
	MonitorEditResponseTypeSmtp     MonitorEditResponseType = "smtp"
)

func (r MonitorEditResponseType) IsKnown() bool {
	switch r {
	case MonitorEditResponseTypeHTTP, MonitorEditResponseTypeHTTPS, MonitorEditResponseTypeTCP, MonitorEditResponseTypeUdpIcmp, MonitorEditResponseTypeIcmpPing, MonitorEditResponseTypeSmtp:
		return true
	}
	return false
}

type MonitorGetResponse struct {
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
	Type MonitorGetResponseType `json:"type"`
	JSON monitorGetResponseJSON `json:"-"`
}

// monitorGetResponseJSON contains the JSON metadata for the struct
// [MonitorGetResponse]
type monitorGetResponseJSON struct {
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

func (r *MonitorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorGetResponseType string

const (
	MonitorGetResponseTypeHTTP     MonitorGetResponseType = "http"
	MonitorGetResponseTypeHTTPS    MonitorGetResponseType = "https"
	MonitorGetResponseTypeTCP      MonitorGetResponseType = "tcp"
	MonitorGetResponseTypeUdpIcmp  MonitorGetResponseType = "udp_icmp"
	MonitorGetResponseTypeIcmpPing MonitorGetResponseType = "icmp_ping"
	MonitorGetResponseTypeSmtp     MonitorGetResponseType = "smtp"
)

func (r MonitorGetResponseType) IsKnown() bool {
	switch r {
	case MonitorGetResponseTypeHTTP, MonitorGetResponseTypeHTTPS, MonitorGetResponseTypeTCP, MonitorGetResponseTypeUdpIcmp, MonitorGetResponseTypeIcmpPing, MonitorGetResponseTypeSmtp:
		return true
	}
	return false
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
	MonitorNewParamsTypeTCP      MonitorNewParamsType = "tcp"
	MonitorNewParamsTypeUdpIcmp  MonitorNewParamsType = "udp_icmp"
	MonitorNewParamsTypeIcmpPing MonitorNewParamsType = "icmp_ping"
	MonitorNewParamsTypeSmtp     MonitorNewParamsType = "smtp"
)

func (r MonitorNewParamsType) IsKnown() bool {
	switch r {
	case MonitorNewParamsTypeHTTP, MonitorNewParamsTypeHTTPS, MonitorNewParamsTypeTCP, MonitorNewParamsTypeUdpIcmp, MonitorNewParamsTypeIcmpPing, MonitorNewParamsTypeSmtp:
		return true
	}
	return false
}

type MonitorNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorNewResponse    `json:"result,required"`
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

// Whether the API call was successful
type MonitorNewResponseEnvelopeSuccess bool

const (
	MonitorNewResponseEnvelopeSuccessTrue MonitorNewResponseEnvelopeSuccess = true
)

func (r MonitorNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	MonitorUpdateParamsTypeTCP      MonitorUpdateParamsType = "tcp"
	MonitorUpdateParamsTypeUdpIcmp  MonitorUpdateParamsType = "udp_icmp"
	MonitorUpdateParamsTypeIcmpPing MonitorUpdateParamsType = "icmp_ping"
	MonitorUpdateParamsTypeSmtp     MonitorUpdateParamsType = "smtp"
)

func (r MonitorUpdateParamsType) IsKnown() bool {
	switch r {
	case MonitorUpdateParamsTypeHTTP, MonitorUpdateParamsTypeHTTPS, MonitorUpdateParamsTypeTCP, MonitorUpdateParamsTypeUdpIcmp, MonitorUpdateParamsTypeIcmpPing, MonitorUpdateParamsTypeSmtp:
		return true
	}
	return false
}

type MonitorUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorUpdateResponse `json:"result,required"`
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

// Whether the API call was successful
type MonitorUpdateResponseEnvelopeSuccess bool

const (
	MonitorUpdateResponseEnvelopeSuccessTrue MonitorUpdateResponseEnvelopeSuccess = true
)

func (r MonitorUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r MonitorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MonitorDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorDeleteResponse `json:"result,required"`
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

// Whether the API call was successful
type MonitorDeleteResponseEnvelopeSuccess bool

const (
	MonitorDeleteResponseEnvelopeSuccessTrue MonitorDeleteResponseEnvelopeSuccess = true
)

func (r MonitorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	MonitorEditParamsTypeTCP      MonitorEditParamsType = "tcp"
	MonitorEditParamsTypeUdpIcmp  MonitorEditParamsType = "udp_icmp"
	MonitorEditParamsTypeIcmpPing MonitorEditParamsType = "icmp_ping"
	MonitorEditParamsTypeSmtp     MonitorEditParamsType = "smtp"
)

func (r MonitorEditParamsType) IsKnown() bool {
	switch r {
	case MonitorEditParamsTypeHTTP, MonitorEditParamsTypeHTTPS, MonitorEditParamsTypeTCP, MonitorEditParamsTypeUdpIcmp, MonitorEditParamsTypeIcmpPing, MonitorEditParamsTypeSmtp:
		return true
	}
	return false
}

type MonitorEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorEditResponse   `json:"result,required"`
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

// Whether the API call was successful
type MonitorEditResponseEnvelopeSuccess bool

const (
	MonitorEditResponseEnvelopeSuccessTrue MonitorEditResponseEnvelopeSuccess = true
)

func (r MonitorEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorGetResponse    `json:"result,required"`
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

// Whether the API call was successful
type MonitorGetResponseEnvelopeSuccess bool

const (
	MonitorGetResponseEnvelopeSuccessTrue MonitorGetResponseEnvelopeSuccess = true
)

func (r MonitorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
