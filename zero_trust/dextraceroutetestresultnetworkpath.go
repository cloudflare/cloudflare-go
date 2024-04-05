// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DEXTracerouteTestResultNetworkPathService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDEXTracerouteTestResultNetworkPathService] method instead.
type DEXTracerouteTestResultNetworkPathService struct {
	Options []option.RequestOption
}

// NewDEXTracerouteTestResultNetworkPathService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDEXTracerouteTestResultNetworkPathService(opts ...option.RequestOption) (r *DEXTracerouteTestResultNetworkPathService) {
	r = &DEXTracerouteTestResultNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *DEXTracerouteTestResultNetworkPathService) Get(ctx context.Context, testResultID string, query DEXTracerouteTestResultNetworkPathGetParams, opts ...option.RequestOption) (res *NetworkPath, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestResultNetworkPathGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", query.AccountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkPath struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []NetworkPathHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string          `json:"testName"`
	JSON     networkPathJSON `json:"-"`
}

// networkPathJSON contains the JSON metadata for the struct [NetworkPath]
type networkPathJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathJSON) RawJSON() string {
	return r.raw
}

type NetworkPathHop struct {
	TTL           int64                   `json:"ttl,required"`
	ASN           int64                   `json:"asn,nullable"`
	Aso           string                  `json:"aso,nullable"`
	IPAddress     string                  `json:"ipAddress,nullable"`
	Location      NetworkPathHopsLocation `json:"location,nullable"`
	Mile          NetworkPathHopsMile     `json:"mile,nullable"`
	Name          string                  `json:"name,nullable"`
	PacketLossPct float64                 `json:"packetLossPct,nullable"`
	RTTMs         int64                   `json:"rttMs,nullable"`
	JSON          networkPathHopJSON      `json:"-"`
}

// networkPathHopJSON contains the JSON metadata for the struct [NetworkPathHop]
type networkPathHopJSON struct {
	TTL           apijson.Field
	ASN           apijson.Field
	Aso           apijson.Field
	IPAddress     apijson.Field
	Location      apijson.Field
	Mile          apijson.Field
	Name          apijson.Field
	PacketLossPct apijson.Field
	RTTMs         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NetworkPathHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathHopJSON) RawJSON() string {
	return r.raw
}

type NetworkPathHopsLocation struct {
	City  string                      `json:"city,nullable"`
	State string                      `json:"state,nullable"`
	Zip   string                      `json:"zip,nullable"`
	JSON  networkPathHopsLocationJSON `json:"-"`
}

// networkPathHopsLocationJSON contains the JSON metadata for the struct
// [NetworkPathHopsLocation]
type networkPathHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkPathHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathHopsLocationJSON) RawJSON() string {
	return r.raw
}

type NetworkPathHopsMile string

const (
	NetworkPathHopsMileClientToApp       NetworkPathHopsMile = "client-to-app"
	NetworkPathHopsMileClientToCfEgress  NetworkPathHopsMile = "client-to-cf-egress"
	NetworkPathHopsMileClientToCfIngress NetworkPathHopsMile = "client-to-cf-ingress"
	NetworkPathHopsMileClientToIsp       NetworkPathHopsMile = "client-to-isp"
)

func (r NetworkPathHopsMile) IsKnown() bool {
	switch r {
	case NetworkPathHopsMileClientToApp, NetworkPathHopsMileClientToCfEgress, NetworkPathHopsMileClientToCfIngress, NetworkPathHopsMileClientToIsp:
		return true
	}
	return false
}

type DEXTracerouteTestResultNetworkPathGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DEXTracerouteTestResultNetworkPathGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   NetworkPath                                               `json:"result,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [DEXTracerouteTestResultNetworkPathGetResponseEnvelope]
type dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccessTrue DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess = true
)

func (r DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
