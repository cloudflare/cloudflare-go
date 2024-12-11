// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package diagnostics

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// TracerouteService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTracerouteService] method instead.
type TracerouteService struct {
	Options []option.RequestOption
}

// NewTracerouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTracerouteService(opts ...option.RequestOption) (r *TracerouteService) {
	r = &TracerouteService{}
	r.Options = opts
	return
}

// Run traceroutes from Cloudflare colos.
func (r *TracerouteService) New(ctx context.Context, params TracerouteNewParams, opts ...option.RequestOption) (res *[]Traceroute, err error) {
	var env TracerouteNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/traceroute", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Traceroute struct {
	Colos []TracerouteColo `json:"colos"`
	// The target hostname, IPv6, or IPv6 address.
	Target string         `json:"target"`
	JSON   tracerouteJSON `json:"-"`
}

// tracerouteJSON contains the JSON metadata for the struct [Traceroute]
type tracerouteJSON struct {
	Colos       apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Traceroute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteJSON) RawJSON() string {
	return r.raw
}

type TracerouteColo struct {
	Colo TracerouteColosColo `json:"colo"`
	// Errors resulting from collecting traceroute from colo to target.
	Error TracerouteColosError `json:"error"`
	Hops  []TracerouteColosHop `json:"hops"`
	// Aggregated statistics from all hops about the target.
	TargetSummary interface{} `json:"target_summary"`
	// Total time of traceroute in ms.
	TracerouteTimeMs int64              `json:"traceroute_time_ms"`
	JSON             tracerouteColoJSON `json:"-"`
}

// tracerouteColoJSON contains the JSON metadata for the struct [TracerouteColo]
type tracerouteColoJSON struct {
	Colo             apijson.Field
	Error            apijson.Field
	Hops             apijson.Field
	TargetSummary    apijson.Field
	TracerouteTimeMs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TracerouteColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteColoJSON) RawJSON() string {
	return r.raw
}

type TracerouteColosColo struct {
	// Source colo city.
	City string `json:"city"`
	// Source colo name.
	Name string                  `json:"name"`
	JSON tracerouteColosColoJSON `json:"-"`
}

// tracerouteColosColoJSON contains the JSON metadata for the struct
// [TracerouteColosColo]
type tracerouteColosColoJSON struct {
	City        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteColosColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteColosColoJSON) RawJSON() string {
	return r.raw
}

// Errors resulting from collecting traceroute from colo to target.
type TracerouteColosError string

const (
	TracerouteColosErrorEmpty                             TracerouteColosError = ""
	TracerouteColosErrorCouldNotGatherTracerouteDataCode1 TracerouteColosError = "Could not gather traceroute data: Code 1"
	TracerouteColosErrorCouldNotGatherTracerouteDataCode2 TracerouteColosError = "Could not gather traceroute data: Code 2"
	TracerouteColosErrorCouldNotGatherTracerouteDataCode3 TracerouteColosError = "Could not gather traceroute data: Code 3"
	TracerouteColosErrorCouldNotGatherTracerouteDataCode4 TracerouteColosError = "Could not gather traceroute data: Code 4"
)

func (r TracerouteColosError) IsKnown() bool {
	switch r {
	case TracerouteColosErrorEmpty, TracerouteColosErrorCouldNotGatherTracerouteDataCode1, TracerouteColosErrorCouldNotGatherTracerouteDataCode2, TracerouteColosErrorCouldNotGatherTracerouteDataCode3, TracerouteColosErrorCouldNotGatherTracerouteDataCode4:
		return true
	}
	return false
}

type TracerouteColosHop struct {
	// An array of node objects.
	Nodes []TracerouteColosHopsNode `json:"nodes"`
	// Number of packets where no response was received.
	PacketsLost int64 `json:"packets_lost"`
	// Number of packets sent with specified TTL.
	PacketsSent int64 `json:"packets_sent"`
	// The time to live (TTL).
	PacketsTTL int64                  `json:"packets_ttl"`
	JSON       tracerouteColosHopJSON `json:"-"`
}

// tracerouteColosHopJSON contains the JSON metadata for the struct
// [TracerouteColosHop]
type tracerouteColosHopJSON struct {
	Nodes       apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	PacketsTTL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteColosHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteColosHopJSON) RawJSON() string {
	return r.raw
}

type TracerouteColosHopsNode struct {
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
	StdDevRTTMs float64                     `json:"std_dev_rtt_ms"`
	JSON        tracerouteColosHopsNodeJSON `json:"-"`
}

// tracerouteColosHopsNodeJSON contains the JSON metadata for the struct
// [TracerouteColosHopsNode]
type tracerouteColosHopsNodeJSON struct {
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

func (r *TracerouteColosHopsNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteColosHopsNodeJSON) RawJSON() string {
	return r.raw
}

type TracerouteNewParams struct {
	// Identifier
	AccountID param.Field[string]   `path:"account_id,required"`
	Targets   param.Field[[]string] `json:"targets,required"`
	// If no source colo names specified, all colos will be used. China colos are
	// unavailable for traceroutes.
	Colos   param.Field[[]string]                   `json:"colos"`
	Options param.Field[TracerouteNewParamsOptions] `json:"options"`
}

func (r TracerouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TracerouteNewParamsOptions struct {
	// Max TTL.
	MaxTTL param.Field[int64] `json:"max_ttl"`
	// Type of packet sent.
	PacketType param.Field[TracerouteNewParamsOptionsPacketType] `json:"packet_type"`
	// Number of packets sent at each TTL.
	PacketsPerTTL param.Field[int64] `json:"packets_per_ttl"`
	// For UDP and TCP, specifies the destination port. For ICMP, specifies the initial
	// ICMP sequence value. Default value 0 will choose the best value to use for each
	// protocol.
	Port param.Field[int64] `json:"port"`
	// Set the time (in seconds) to wait for a response to a probe.
	WaitTime param.Field[int64] `json:"wait_time"`
}

func (r TracerouteNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of packet sent.
type TracerouteNewParamsOptionsPacketType string

const (
	TracerouteNewParamsOptionsPacketTypeIcmp    TracerouteNewParamsOptionsPacketType = "icmp"
	TracerouteNewParamsOptionsPacketTypeTCP     TracerouteNewParamsOptionsPacketType = "tcp"
	TracerouteNewParamsOptionsPacketTypeUdp     TracerouteNewParamsOptionsPacketType = "udp"
	TracerouteNewParamsOptionsPacketTypeGRE     TracerouteNewParamsOptionsPacketType = "gre"
	TracerouteNewParamsOptionsPacketTypeGREIcmp TracerouteNewParamsOptionsPacketType = "gre+icmp"
)

func (r TracerouteNewParamsOptionsPacketType) IsKnown() bool {
	switch r {
	case TracerouteNewParamsOptionsPacketTypeIcmp, TracerouteNewParamsOptionsPacketTypeTCP, TracerouteNewParamsOptionsPacketTypeUdp, TracerouteNewParamsOptionsPacketTypeGRE, TracerouteNewParamsOptionsPacketTypeGREIcmp:
		return true
	}
	return false
}

type TracerouteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []Traceroute          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TracerouteNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TracerouteNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tracerouteNewResponseEnvelopeJSON       `json:"-"`
}

// tracerouteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TracerouteNewResponseEnvelope]
type tracerouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TracerouteNewResponseEnvelopeSuccess bool

const (
	TracerouteNewResponseEnvelopeSuccessTrue TracerouteNewResponseEnvelopeSuccess = true
)

func (r TracerouteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TracerouteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TracerouteNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       tracerouteNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// tracerouteNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [TracerouteNewResponseEnvelopeResultInfo]
type tracerouteNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TracerouteNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tracerouteNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
