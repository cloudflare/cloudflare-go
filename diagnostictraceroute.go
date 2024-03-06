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
func (r *DiagnosticTracerouteService) New(ctx context.Context, params DiagnosticTracerouteNewParams, opts ...option.RequestOption) (res *[]DiagnosticTracerouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DiagnosticTracerouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/diagnostics/traceroute", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DiagnosticTracerouteNewResponse struct {
	Colos []DiagnosticTracerouteNewResponseColo `json:"colos"`
	// The target hostname, IPv6, or IPv6 address.
	Target string                              `json:"target"`
	JSON   diagnosticTracerouteNewResponseJSON `json:"-"`
}

// diagnosticTracerouteNewResponseJSON contains the JSON metadata for the struct
// [DiagnosticTracerouteNewResponse]
type diagnosticTracerouteNewResponseJSON struct {
	Colos       apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseJSON) RawJSON() string {
	return r.raw
}

type DiagnosticTracerouteNewResponseColo struct {
	Colo DiagnosticTracerouteNewResponseColosColo `json:"colo"`
	// Errors resulting from collecting traceroute from colo to target.
	Error DiagnosticTracerouteNewResponseColosError `json:"error"`
	Hops  []DiagnosticTracerouteNewResponseColosHop `json:"hops"`
	// Aggregated statistics from all hops about the target.
	TargetSummary interface{} `json:"target_summary"`
	// Total time of traceroute in ms.
	TracerouteTimeMs int64                                   `json:"traceroute_time_ms"`
	JSON             diagnosticTracerouteNewResponseColoJSON `json:"-"`
}

// diagnosticTracerouteNewResponseColoJSON contains the JSON metadata for the
// struct [DiagnosticTracerouteNewResponseColo]
type diagnosticTracerouteNewResponseColoJSON struct {
	Colo             apijson.Field
	Error            apijson.Field
	Hops             apijson.Field
	TargetSummary    apijson.Field
	TracerouteTimeMs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseColoJSON) RawJSON() string {
	return r.raw
}

type DiagnosticTracerouteNewResponseColosColo struct {
	// Source colo city.
	City string `json:"city"`
	// Source colo name.
	Name string                                       `json:"name"`
	JSON diagnosticTracerouteNewResponseColosColoJSON `json:"-"`
}

// diagnosticTracerouteNewResponseColosColoJSON contains the JSON metadata for the
// struct [DiagnosticTracerouteNewResponseColosColo]
type diagnosticTracerouteNewResponseColosColoJSON struct {
	City        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseColosColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseColosColoJSON) RawJSON() string {
	return r.raw
}

// Errors resulting from collecting traceroute from colo to target.
type DiagnosticTracerouteNewResponseColosError string

const (
	DiagnosticTracerouteNewResponseColosErrorEmpty                             DiagnosticTracerouteNewResponseColosError = ""
	DiagnosticTracerouteNewResponseColosErrorCouldNotGatherTracerouteDataCode1 DiagnosticTracerouteNewResponseColosError = "Could not gather traceroute data: Code 1"
	DiagnosticTracerouteNewResponseColosErrorCouldNotGatherTracerouteDataCode2 DiagnosticTracerouteNewResponseColosError = "Could not gather traceroute data: Code 2"
	DiagnosticTracerouteNewResponseColosErrorCouldNotGatherTracerouteDataCode3 DiagnosticTracerouteNewResponseColosError = "Could not gather traceroute data: Code 3"
	DiagnosticTracerouteNewResponseColosErrorCouldNotGatherTracerouteDataCode4 DiagnosticTracerouteNewResponseColosError = "Could not gather traceroute data: Code 4"
)

type DiagnosticTracerouteNewResponseColosHop struct {
	// An array of node objects.
	Nodes []DiagnosticTracerouteNewResponseColosHopsNode `json:"nodes"`
	// Number of packets where no response was received.
	PacketsLost int64 `json:"packets_lost"`
	// Number of packets sent with specified TTL.
	PacketsSent int64 `json:"packets_sent"`
	// The time to live (TTL).
	PacketsTTL int64                                       `json:"packets_ttl"`
	JSON       diagnosticTracerouteNewResponseColosHopJSON `json:"-"`
}

// diagnosticTracerouteNewResponseColosHopJSON contains the JSON metadata for the
// struct [DiagnosticTracerouteNewResponseColosHop]
type diagnosticTracerouteNewResponseColosHopJSON struct {
	Nodes       apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	PacketsTTL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseColosHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseColosHopJSON) RawJSON() string {
	return r.raw
}

type DiagnosticTracerouteNewResponseColosHopsNode struct {
	// AS number associated with the node object.
	ASN string `json:"asn"`
	// IP address of the node.
	IP string `json:"ip"`
	// Field appears if there is an additional annotation printed when the probe
	// returns. Field also appears when running a GRE+ICMP traceroute to denote which
	// traceroute a node comes from.
	Labels []string `json:"labels"`
	// Maximum RTT in ms.
	MaxRTTMs float64 `json:"max_rtt_ms"`
	// Mean RTT in ms.
	MeanRTTMs float64 `json:"mean_rtt_ms"`
	// Minimum RTT in ms.
	MinRTTMs float64 `json:"min_rtt_ms"`
	// Host name of the address, this may be the same as the IP address.
	Name string `json:"name"`
	// Number of packets with a response from this node.
	PacketCount int64 `json:"packet_count"`
	// Standard deviation of the RTTs in ms.
	StdDevRTTMs float64                                          `json:"std_dev_rtt_ms"`
	JSON        diagnosticTracerouteNewResponseColosHopsNodeJSON `json:"-"`
}

// diagnosticTracerouteNewResponseColosHopsNodeJSON contains the JSON metadata for
// the struct [DiagnosticTracerouteNewResponseColosHopsNode]
type diagnosticTracerouteNewResponseColosHopsNodeJSON struct {
	ASN         apijson.Field
	IP          apijson.Field
	Labels      apijson.Field
	MaxRTTMs    apijson.Field
	MeanRTTMs   apijson.Field
	MinRTTMs    apijson.Field
	Name        apijson.Field
	PacketCount apijson.Field
	StdDevRTTMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseColosHopsNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseColosHopsNodeJSON) RawJSON() string {
	return r.raw
}

type DiagnosticTracerouteNewParams struct {
	// Identifier
	AccountID param.Field[string]   `path:"account_id,required"`
	Targets   param.Field[[]string] `json:"targets,required"`
	// If no source colo names specified, all colos will be used. China colos are
	// unavailable for traceroutes.
	Colos   param.Field[[]string]                             `json:"colos"`
	Options param.Field[DiagnosticTracerouteNewParamsOptions] `json:"options"`
}

func (r DiagnosticTracerouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiagnosticTracerouteNewParamsOptions struct {
	// Max TTL.
	MaxTTL param.Field[int64] `json:"max_ttl"`
	// Type of packet sent.
	PacketType param.Field[DiagnosticTracerouteNewParamsOptionsPacketType] `json:"packet_type"`
	// Number of packets sent at each TTL.
	PacketsPerTTL param.Field[int64] `json:"packets_per_ttl"`
	// For UDP and TCP, specifies the destination port. For ICMP, specifies the initial
	// ICMP sequence value. Default value 0 will choose the best value to use for each
	// protocol.
	Port param.Field[int64] `json:"port"`
	// Set the time (in seconds) to wait for a response to a probe.
	WaitTime param.Field[int64] `json:"wait_time"`
}

func (r DiagnosticTracerouteNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of packet sent.
type DiagnosticTracerouteNewParamsOptionsPacketType string

const (
	DiagnosticTracerouteNewParamsOptionsPacketTypeIcmp    DiagnosticTracerouteNewParamsOptionsPacketType = "icmp"
	DiagnosticTracerouteNewParamsOptionsPacketTypeTcp     DiagnosticTracerouteNewParamsOptionsPacketType = "tcp"
	DiagnosticTracerouteNewParamsOptionsPacketTypeUdp     DiagnosticTracerouteNewParamsOptionsPacketType = "udp"
	DiagnosticTracerouteNewParamsOptionsPacketTypeGRE     DiagnosticTracerouteNewParamsOptionsPacketType = "gre"
	DiagnosticTracerouteNewParamsOptionsPacketTypeGREIcmp DiagnosticTracerouteNewParamsOptionsPacketType = "gre+icmp"
)

type DiagnosticTracerouteNewResponseEnvelope struct {
	Errors   []DiagnosticTracerouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DiagnosticTracerouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []DiagnosticTracerouteNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DiagnosticTracerouteNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DiagnosticTracerouteNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       diagnosticTracerouteNewResponseEnvelopeJSON       `json:"-"`
}

// diagnosticTracerouteNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DiagnosticTracerouteNewResponseEnvelope]
type diagnosticTracerouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DiagnosticTracerouteNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    diagnosticTracerouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// diagnosticTracerouteNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DiagnosticTracerouteNewResponseEnvelopeErrors]
type diagnosticTracerouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DiagnosticTracerouteNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    diagnosticTracerouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// diagnosticTracerouteNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DiagnosticTracerouteNewResponseEnvelopeMessages]
type diagnosticTracerouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DiagnosticTracerouteNewResponseEnvelopeSuccess bool

const (
	DiagnosticTracerouteNewResponseEnvelopeSuccessTrue DiagnosticTracerouteNewResponseEnvelopeSuccess = true
)

type DiagnosticTracerouteNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       diagnosticTracerouteNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// diagnosticTracerouteNewResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DiagnosticTracerouteNewResponseEnvelopeResultInfo]
type diagnosticTracerouteNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiagnosticTracerouteNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diagnosticTracerouteNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
