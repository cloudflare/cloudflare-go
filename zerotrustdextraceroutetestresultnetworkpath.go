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

// ZeroTrustDEXTracerouteTestResultNetworkPathService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZeroTrustDEXTracerouteTestResultNetworkPathService] method instead.
type ZeroTrustDEXTracerouteTestResultNetworkPathService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXTracerouteTestResultNetworkPathService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZeroTrustDEXTracerouteTestResultNetworkPathService(opts ...option.RequestOption) (r *ZeroTrustDEXTracerouteTestResultNetworkPathService) {
	r = &ZeroTrustDEXTracerouteTestResultNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *ZeroTrustDEXTracerouteTestResultNetworkPathService) List(ctx context.Context, testResultID string, query ZeroTrustDEXTracerouteTestResultNetworkPathListParams, opts ...option.RequestOption) (res *ZeroTrustDEXTracerouteTestResultNetworkPathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", query.AccountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponse struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                                      `json:"testName"`
	JSON     zeroTrustDEXTracerouteTestResultNetworkPathListResponseJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestResultNetworkPathListResponseJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTracerouteTestResultNetworkPathListResponse]
type zeroTrustDEXTracerouteTestResultNetworkPathListResponseJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestResultNetworkPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHop struct {
	TTL           int64                                                               `json:"ttl,required"`
	ASN           int64                                                               `json:"asn,nullable"`
	Aso           string                                                              `json:"aso,nullable"`
	IPAddress     string                                                              `json:"ipAddress,nullable"`
	Location      ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocation `json:"location,nullable"`
	Mile          ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMile     `json:"mile,nullable"`
	Name          string                                                              `json:"name,nullable"`
	PacketLossPct float64                                                             `json:"packetLossPct,nullable"`
	RTTMs         int64                                                               `json:"rttMs,nullable"`
	JSON          zeroTrustDEXTracerouteTestResultNetworkPathListResponseHopJSON      `json:"-"`
}

// zeroTrustDEXTracerouteTestResultNetworkPathListResponseHopJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHop]
type zeroTrustDEXTracerouteTestResultNetworkPathListResponseHopJSON struct {
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

func (r *ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocation struct {
	City  string                                                                  `json:"city,nullable"`
	State string                                                                  `json:"state,nullable"`
	Zip   string                                                                  `json:"zip,nullable"`
	JSON  zeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocationJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocationJSON contains
// the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocation]
type zeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMile string

const (
	ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMileClientToApp       ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-app"
	ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMileClientToCfEgress  ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-cf-egress"
	ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMileClientToCfIngress ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-cf-ingress"
	ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMileClientToIsp       ZeroTrustDEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-isp"
)

type ZeroTrustDEXTracerouteTestResultNetworkPathListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelope struct {
	Errors   []ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDEXTracerouteTestResultNetworkPathListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelope]
type zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors]
type zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages]
type zeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccessTrue ZeroTrustDEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess = true
)
