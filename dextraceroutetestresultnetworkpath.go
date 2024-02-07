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
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", accountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DexTracerouteTestResultNetworkPathListResponse struct {
	Errors   []DexTracerouteTestResultNetworkPathListResponseError   `json:"errors"`
	Messages []DexTracerouteTestResultNetworkPathListResponseMessage `json:"messages"`
	Result   DexTracerouteTestResultNetworkPathListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexTracerouteTestResultNetworkPathListResponseSuccess `json:"success"`
	JSON    dexTracerouteTestResultNetworkPathListResponseJSON    `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseJSON contains the JSON metadata
// for the struct [DexTracerouteTestResultNetworkPathListResponse]
type dexTracerouteTestResultNetworkPathListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseError struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseErrorJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseErrorJSON contains the JSON
// metadata for the struct [DexTracerouteTestResultNetworkPathListResponseError]
type dexTracerouteTestResultNetworkPathListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseMessage struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseMessageJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseMessageJSON contains the JSON
// metadata for the struct [DexTracerouteTestResultNetworkPathListResponseMessage]
type dexTracerouteTestResultNetworkPathListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseResult struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []DexTracerouteTestResultNetworkPathListResponseResultHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                                   `json:"testName"`
	JSON     dexTracerouteTestResultNetworkPathListResponseResultJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseResultJSON contains the JSON
// metadata for the struct [DexTracerouteTestResultNetworkPathListResponseResult]
type dexTracerouteTestResultNetworkPathListResponseResultJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseResultHop struct {
	TTL           int64                                                            `json:"ttl,required"`
	Asn           int64                                                            `json:"asn,nullable"`
	Aso           string                                                           `json:"aso,nullable"`
	IPAddress     string                                                           `json:"ipAddress,nullable"`
	Location      DexTracerouteTestResultNetworkPathListResponseResultHopsLocation `json:"location,nullable"`
	Mile          DexTracerouteTestResultNetworkPathListResponseResultHopsMile     `json:"mile,nullable"`
	Name          string                                                           `json:"name,nullable"`
	PacketLossPct float64                                                          `json:"packetLossPct,nullable"`
	RttMs         int64                                                            `json:"rttMs,nullable"`
	JSON          dexTracerouteTestResultNetworkPathListResponseResultHopJSON      `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseResultHopJSON contains the JSON
// metadata for the struct
// [DexTracerouteTestResultNetworkPathListResponseResultHop]
type dexTracerouteTestResultNetworkPathListResponseResultHopJSON struct {
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

func (r *DexTracerouteTestResultNetworkPathListResponseResultHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseResultHopsLocation struct {
	City  string                                                               `json:"city,nullable"`
	State string                                                               `json:"state,nullable"`
	Zip   string                                                               `json:"zip,nullable"`
	JSON  dexTracerouteTestResultNetworkPathListResponseResultHopsLocationJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseResultHopsLocationJSON contains
// the JSON metadata for the struct
// [DexTracerouteTestResultNetworkPathListResponseResultHopsLocation]
type dexTracerouteTestResultNetworkPathListResponseResultHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTracerouteTestResultNetworkPathListResponseResultHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTracerouteTestResultNetworkPathListResponseResultHopsMile string

const (
	DexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToApp       DexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-app"
	DexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToCfEgress  DexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-cf-egress"
	DexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToCfIngress DexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-cf-ingress"
	DexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToIsp       DexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-isp"
)

// Whether the API call was successful
type DexTracerouteTestResultNetworkPathListResponseSuccess bool

const (
	DexTracerouteTestResultNetworkPathListResponseSuccessTrue DexTracerouteTestResultNetworkPathListResponseSuccess = true
)
