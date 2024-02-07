// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DexTracerouteTestResultNetworkPathService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDexTracerouteTestResultNetworkPathService] method instead.
type DexTracerouteTestResultNetworkPathService struct {
	Options []option.RequestOption
}

// NewDexTracerouteTestResultNetworkPathService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDexTracerouteTestResultNetworkPathService(opts ...option.RequestOption) (r *DexTracerouteTestResultNetworkPathService) {
	r = &DexTracerouteTestResultNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *DexTracerouteTestResultNetworkPathService) List(ctx context.Context, accountID string, testResultID string, opts ...option.RequestOption) (res *DexTracerouteTestResultNetworkPathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexTracerouteTestResultNetworkPathListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", accountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexTracerouteTestResultNetworkPathListResponse struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []DexTracerouteTestResultNetworkPathListResponseHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                             `json:"testName"`
	JSON     dexTracerouteTestResultNetworkPathListResponseJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseJSON contains the JSON metadata
// for the struct [DexTracerouteTestResultNetworkPathListResponse]
type dexTracerouteTestResultNetworkPathListResponseJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseHop struct {
	TTL           int64                                                      `json:"ttl,required"`
	Asn           int64                                                      `json:"asn,nullable"`
	Aso           string                                                     `json:"aso,nullable"`
	IPAddress     string                                                     `json:"ipAddress,nullable"`
	Location      DexTracerouteTestResultNetworkPathListResponseHopsLocation `json:"location,nullable"`
	Mile          DexTracerouteTestResultNetworkPathListResponseHopsMile     `json:"mile,nullable"`
	Name          string                                                     `json:"name,nullable"`
	PacketLossPct float64                                                    `json:"packetLossPct,nullable"`
	RttMs         int64                                                      `json:"rttMs,nullable"`
	JSON          dexTracerouteTestResultNetworkPathListResponseHopJSON      `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseHopJSON contains the JSON metadata
// for the struct [DexTracerouteTestResultNetworkPathListResponseHop]
type dexTracerouteTestResultNetworkPathListResponseHopJSON struct {
	TTL           apijson.Field
	Asn           apijson.Field
	Aso           apijson.Field
	IPAddress     apijson.Field
	Location      apijson.Field
	Mile          apijson.Field
	Name          apijson.Field
	PacketLossPct apijson.Field
	RttMs         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseHopsLocation struct {
	City  string                                                         `json:"city,nullable"`
	State string                                                         `json:"state,nullable"`
	Zip   string                                                         `json:"zip,nullable"`
	JSON  dexTracerouteTestResultNetworkPathListResponseHopsLocationJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseHopsLocationJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestResultNetworkPathListResponseHopsLocation]
type dexTracerouteTestResultNetworkPathListResponseHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseHopsMile string

const (
	DexTracerouteTestResultNetworkPathListResponseHopsMileClientToApp       DexTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-app"
	DexTracerouteTestResultNetworkPathListResponseHopsMileClientToCfEgress  DexTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-cf-egress"
	DexTracerouteTestResultNetworkPathListResponseHopsMileClientToCfIngress DexTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-cf-ingress"
	DexTracerouteTestResultNetworkPathListResponseHopsMileClientToIsp       DexTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-isp"
)

type DexTracerouteTestResultNetworkPathListResponseEnvelope struct {
	Errors   []DexTracerouteTestResultNetworkPathListResponseEnvelopeErrors   `json:"errors"`
	Messages []DexTracerouteTestResultNetworkPathListResponseEnvelopeMessages `json:"messages"`
	Result   DexTracerouteTestResultNetworkPathListResponse                   `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestResultNetworkPathListResponseEnvelopeSuccess `json:"success"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON contains the JSON
// metadata for the struct [DexTracerouteTestResultNetworkPathListResponseEnvelope]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestResultNetworkPathListResponseEnvelopeErrors]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DexTracerouteTestResultNetworkPathListResponseEnvelopeMessages]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTracerouteTestResultNetworkPathListResponseEnvelopeSuccess bool

const (
	DexTracerouteTestResultNetworkPathListResponseEnvelopeSuccessTrue DexTracerouteTestResultNetworkPathListResponseEnvelopeSuccess = true
)
