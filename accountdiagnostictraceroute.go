// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDiagnosticTracerouteService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDiagnosticTracerouteService] method instead.
type AccountDiagnosticTracerouteService struct {
	Options []option.RequestOption
}

// NewAccountDiagnosticTracerouteService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDiagnosticTracerouteService(opts ...option.RequestOption) (r *AccountDiagnosticTracerouteService) {
	r = &AccountDiagnosticTracerouteService{}
	r.Options = opts
	return
}

// Run traceroutes from Cloudflare colos.
func (r *AccountDiagnosticTracerouteService) DiagnosticsTraceroute(ctx context.Context, accountIdentifier string, body AccountDiagnosticTracerouteDiagnosticsTracerouteParams, opts ...option.RequestOption) (res *AccountDiagnosticTracerouteDiagnosticsTracerouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/diagnostics/traceroute", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponse struct {
	Errors     []AccountDiagnosticTracerouteDiagnosticsTracerouteResponseError    `json:"errors"`
	Messages   []AccountDiagnosticTracerouteDiagnosticsTracerouteResponseMessage  `json:"messages"`
	Result     []AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResult   `json:"result"`
	ResultInfo AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountDiagnosticTracerouteDiagnosticsTracerouteResponseSuccess `json:"success"`
	JSON    accountDiagnosticTracerouteDiagnosticsTracerouteResponseJSON    `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseJSON contains the JSON
// metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponse]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountDiagnosticTracerouteDiagnosticsTracerouteResponseErrorJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseError]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountDiagnosticTracerouteDiagnosticsTracerouteResponseMessageJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseMessage]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResult struct {
	Colos []AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColo `json:"colos"`
	// The target hostname, IPv6, or IPv6 address.
	Target string                                                             `json:"target"`
	JSON   accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultJSON contains the
// JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResult]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultJSON struct {
	Colos       apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColo struct {
	Colo AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColo `json:"colo"`
	// Errors resulting from collecting traceroute from colo to target.
	Error AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError `json:"error"`
	Hops  []AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHop `json:"hops"`
	// Aggregated statistics from all hops about the target.
	TargetSummary interface{} `json:"target_summary"`
	// Total time of traceroute in ms.
	TracerouteTimeMs int64                                                                  `json:"traceroute_time_ms"`
	JSON             accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColoJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColoJSON contains
// the JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColo]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColoJSON struct {
	Colo             apijson.Field
	Error            apijson.Field
	Hops             apijson.Field
	TargetSummary    apijson.Field
	TracerouteTimeMs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColo struct {
	// Source colo city.
	City string `json:"city"`
	// Source colo name.
	Name string                                                                      `json:"name"`
	JSON accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColoJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColoJSON
// contains the JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColo]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColoJSON struct {
	City        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Errors resulting from collecting traceroute from colo to target.
type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError string

const (
	AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosErrorEmpty                             AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError = ""
	AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode1 AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError = "Could not gather traceroute data: Code 1"
	AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode2 AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError = "Could not gather traceroute data: Code 2"
	AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode3 AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError = "Could not gather traceroute data: Code 3"
	AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode4 AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosError = "Could not gather traceroute data: Code 4"
)

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHop struct {
	// An array of node objects.
	Nodes []AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNode `json:"nodes"`
	// Number of packets where no response was received.
	PacketsLost int64 `json:"packets_lost"`
	// Number of packets sent with specified TTL.
	PacketsSent int64 `json:"packets_sent"`
	// The time to live (TTL).
	PacketsTtl int64                                                                      `json:"packets_ttl"`
	JSON       accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopJSON
// contains the JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHop]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopJSON struct {
	Nodes       apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	PacketsTtl  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNode struct {
	// AS number associated with the node object.
	ASN string `json:"asn"`
	// IP address of the node.
	IP string `json:"ip"`
	// Field appears if there is an additional annotation printed when the probe
	// returns. Field also appears when running a GRE+ICMP traceroute to denote which
	// traceroute a node comes from.
	Labels []string `json:"labels"`
	// Maximum RTT in ms.
	MaxRttMs float64 `json:"max_rtt_ms"`
	// Mean RTT in ms.
	MeanRttMs float64 `json:"mean_rtt_ms"`
	// Minimum RTT in ms.
	MinRttMs float64 `json:"min_rtt_ms"`
	// Host name of the address, this may be the same as the IP address.
	Name string `json:"name"`
	// Number of packets with a response from this node.
	PacketCount int64 `json:"packet_count"`
	// Standard deviation of the RTTs in ms.
	StdDevRttMs float64                                                                         `json:"std_dev_rtt_ms"`
	JSON        accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNodeJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNodeJSON
// contains the JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNode]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNodeJSON struct {
	ASN         apijson.Field
	IP          apijson.Field
	Labels      apijson.Field
	MaxRttMs    apijson.Field
	MeanRttMs   apijson.Field
	MinRttMs    apijson.Field
	Name        apijson.Field
	PacketCount apijson.Field
	StdDevRttMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultColosHopsNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfoJSON `json:"-"`
}

// accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfo]
type accountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticTracerouteDiagnosticsTracerouteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDiagnosticTracerouteDiagnosticsTracerouteResponseSuccess bool

const (
	AccountDiagnosticTracerouteDiagnosticsTracerouteResponseSuccessTrue AccountDiagnosticTracerouteDiagnosticsTracerouteResponseSuccess = true
)

type AccountDiagnosticTracerouteDiagnosticsTracerouteParams struct {
	Targets param.Field[[]string] `json:"targets,required"`
	// If no source colo names specified, all colos will be used. China colos are
	// unavailable for traceroutes.
	Colos   param.Field[[]string]                                                      `json:"colos"`
	Options param.Field[AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptions] `json:"options"`
}

func (r AccountDiagnosticTracerouteDiagnosticsTracerouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptions struct {
	// Max TTL.
	MaxTtl param.Field[int64] `json:"max_ttl"`
	// Type of packet sent.
	PacketType param.Field[AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType] `json:"packet_type"`
	// Number of packets sent at each TTL.
	PacketsPerTtl param.Field[int64] `json:"packets_per_ttl"`
	// For UDP and TCP, specifies the destination port. For ICMP, specifies the initial
	// ICMP sequence value. Default value 0 will choose the best value to use for each
	// protocol.
	Port param.Field[int64] `json:"port"`
	// Set the time (in seconds) to wait for a response to a probe.
	WaitTime param.Field[int64] `json:"wait_time"`
}

func (r AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of packet sent.
type AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType string

const (
	AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeIcmp    AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "icmp"
	AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeTcp     AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "tcp"
	AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeUdp     AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "udp"
	AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeGre     AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "gre"
	AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeGreIcmp AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "gre+icmp"
)
