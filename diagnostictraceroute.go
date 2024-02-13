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

// DiagnosticTracerouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDiagnosticTracerouteService]
// method instead.
type DiagnosticTracerouteService struct {
	Options []option.RequestOption
}

// NewDiagnosticTracerouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDiagnosticTracerouteService(opts ...option.RequestOption) (r *DiagnosticTracerouteService) {
	r = &DiagnosticTracerouteService{}
	r.Options = opts
	return
}

// Run traceroutes from Cloudflare colos.
func (r *DiagnosticTracerouteService) DiagnosticsTraceroute(ctx context.Context, accountIdentifier string, body DiagnosticTracerouteDiagnosticsTracerouteParams, opts ...option.RequestOption) (res *[]DiagnosticTracerouteDiagnosticsTracerouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/diagnostics/traceroute", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DiagnosticTracerouteDiagnosticsTracerouteResponse struct {
	Colos []DiagnosticTracerouteDiagnosticsTracerouteResponseColo `json:"colos"`
	// The target hostname, IPv6, or IPv6 address.
	Target string                                                `json:"target"`
	JSON   diagnosticTracerouteDiagnosticsTracerouteResponseJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseJSON contains the JSON metadata
// for the struct [DiagnosticTracerouteDiagnosticsTracerouteResponse]
type diagnosticTracerouteDiagnosticsTracerouteResponseJSON struct {
	Colos       apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DiagnosticTracerouteDiagnosticsTracerouteResponseColo struct {
	Colo DiagnosticTracerouteDiagnosticsTracerouteResponseColosColo `json:"colo"`
	// Errors resulting from collecting traceroute from colo to target.
	Error DiagnosticTracerouteDiagnosticsTracerouteResponseColosError `json:"error"`
	Hops  []DiagnosticTracerouteDiagnosticsTracerouteResponseColosHop `json:"hops"`
	// Aggregated statistics from all hops about the target.
	TargetSummary interface{} `json:"target_summary"`
	// Total time of traceroute in ms.
	TracerouteTimeMs int64                                                     `json:"traceroute_time_ms"`
	JSON             diagnosticTracerouteDiagnosticsTracerouteResponseColoJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseColoJSON contains the JSON
// metadata for the struct [DiagnosticTracerouteDiagnosticsTracerouteResponseColo]
type diagnosticTracerouteDiagnosticsTracerouteResponseColoJSON struct {
	Colo             apijson.Field
	Error            apijson.Field
	Hops             apijson.Field
	TargetSummary    apijson.Field
	TracerouteTimeMs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DiagnosticTracerouteDiagnosticsTracerouteResponseColosColo struct {
	// Source colo city.
	City string `json:"city"`
	// Source colo name.
	Name string                                                         `json:"name"`
	JSON diagnosticTracerouteDiagnosticsTracerouteResponseColosColoJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseColosColoJSON contains the JSON
// metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseColosColo]
type diagnosticTracerouteDiagnosticsTracerouteResponseColosColoJSON struct {
	City        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseColosColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Errors resulting from collecting traceroute from colo to target.
type DiagnosticTracerouteDiagnosticsTracerouteResponseColosError string

const (
	DiagnosticTracerouteDiagnosticsTracerouteResponseColosErrorEmpty                             DiagnosticTracerouteDiagnosticsTracerouteResponseColosError = ""
	DiagnosticTracerouteDiagnosticsTracerouteResponseColosErrorCouldNotGatherTracerouteDataCode1 DiagnosticTracerouteDiagnosticsTracerouteResponseColosError = "Could not gather traceroute data: Code 1"
	DiagnosticTracerouteDiagnosticsTracerouteResponseColosErrorCouldNotGatherTracerouteDataCode2 DiagnosticTracerouteDiagnosticsTracerouteResponseColosError = "Could not gather traceroute data: Code 2"
	DiagnosticTracerouteDiagnosticsTracerouteResponseColosErrorCouldNotGatherTracerouteDataCode3 DiagnosticTracerouteDiagnosticsTracerouteResponseColosError = "Could not gather traceroute data: Code 3"
	DiagnosticTracerouteDiagnosticsTracerouteResponseColosErrorCouldNotGatherTracerouteDataCode4 DiagnosticTracerouteDiagnosticsTracerouteResponseColosError = "Could not gather traceroute data: Code 4"
)

type DiagnosticTracerouteDiagnosticsTracerouteResponseColosHop struct {
	// An array of node objects.
	Nodes []DiagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNode `json:"nodes"`
	// Number of packets where no response was received.
	PacketsLost int64 `json:"packets_lost"`
	// Number of packets sent with specified TTL.
	PacketsSent int64 `json:"packets_sent"`
	// The time to live (TTL).
	PacketsTTL int64                                                         `json:"packets_ttl"`
	JSON       diagnosticTracerouteDiagnosticsTracerouteResponseColosHopJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseColosHopJSON contains the JSON
// metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseColosHop]
type diagnosticTracerouteDiagnosticsTracerouteResponseColosHopJSON struct {
	Nodes       apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	PacketsTTL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseColosHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DiagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNode struct {
	// AS number associated with the node object.
	Asn string `json:"asn"`
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
	StdDevRttMs float64                                                            `json:"std_dev_rtt_ms"`
	JSON        diagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNodeJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNodeJSON contains the
// JSON metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNode]
type diagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNodeJSON struct {
	Asn         apijson.Field
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

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseColosHopsNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DiagnosticTracerouteDiagnosticsTracerouteParams struct {
	Targets param.Field[[]string] `json:"targets,required"`
	// If no source colo names specified, all colos will be used. China colos are
	// unavailable for traceroutes.
	Colos   param.Field[[]string]                                               `json:"colos"`
	Options param.Field[DiagnosticTracerouteDiagnosticsTracerouteParamsOptions] `json:"options"`
}

func (r DiagnosticTracerouteDiagnosticsTracerouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiagnosticTracerouteDiagnosticsTracerouteParamsOptions struct {
	// Max TTL.
	MaxTTL param.Field[int64] `json:"max_ttl"`
	// Type of packet sent.
	PacketType param.Field[DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType] `json:"packet_type"`
	// Number of packets sent at each TTL.
	PacketsPerTTL param.Field[int64] `json:"packets_per_ttl"`
	// For UDP and TCP, specifies the destination port. For ICMP, specifies the initial
	// ICMP sequence value. Default value 0 will choose the best value to use for each
	// protocol.
	Port param.Field[int64] `json:"port"`
	// Set the time (in seconds) to wait for a response to a probe.
	WaitTime param.Field[int64] `json:"wait_time"`
}

func (r DiagnosticTracerouteDiagnosticsTracerouteParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of packet sent.
type DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType string

const (
	DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeIcmp    DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "icmp"
	DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeTcp     DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "tcp"
	DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeUdp     DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "udp"
	DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeGre     DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "gre"
	DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeGreIcmp DiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketType = "gre+icmp"
)

type DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelope struct {
	Errors   []DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessages `json:"messages,required"`
	Result   []DiagnosticTracerouteDiagnosticsTracerouteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeJSON       `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelope]
type diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrorsJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrors]
type diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessagesJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessages]
type diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeSuccess bool

const (
	DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeSuccessTrue DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeSuccess = true
)

type DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfoJSON `json:"-"`
}

// diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfo]
type diagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteDiagnosticsTracerouteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
