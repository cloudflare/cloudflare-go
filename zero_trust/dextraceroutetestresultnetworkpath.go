// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
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
func (r *DEXTracerouteTestResultNetworkPathService) List(ctx context.Context, testResultID string, query DEXTracerouteTestResultNetworkPathListParams, opts ...option.RequestOption) (res *DigitalExperienceMonitoringTracerouteTestResultNetworkPath, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestResultNetworkPathListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", query.AccountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DigitalExperienceMonitoringTracerouteTestResultNetworkPath struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []DigitalExperienceMonitoringTracerouteTestResultNetworkPathHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                                         `json:"testName"`
	JSON     digitalExperienceMonitoringTracerouteTestResultNetworkPathJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestResultNetworkPathJSON contains the JSON
// metadata for the struct
// [DigitalExperienceMonitoringTracerouteTestResultNetworkPath]
type digitalExperienceMonitoringTracerouteTestResultNetworkPathJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteTestResultNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestResultNetworkPathJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestResultNetworkPathHop struct {
	TTL           int64                                                                  `json:"ttl,required"`
	ASN           int64                                                                  `json:"asn,nullable"`
	Aso           string                                                                 `json:"aso,nullable"`
	IPAddress     string                                                                 `json:"ipAddress,nullable"`
	Location      DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocation `json:"location,nullable"`
	Mile          DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMile     `json:"mile,nullable"`
	Name          string                                                                 `json:"name,nullable"`
	PacketLossPct float64                                                                `json:"packetLossPct,nullable"`
	RTTMs         int64                                                                  `json:"rttMs,nullable"`
	JSON          digitalExperienceMonitoringTracerouteTestResultNetworkPathHopJSON      `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestResultNetworkPathHopJSON contains the
// JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteTestResultNetworkPathHop]
type digitalExperienceMonitoringTracerouteTestResultNetworkPathHopJSON struct {
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

func (r *DigitalExperienceMonitoringTracerouteTestResultNetworkPathHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestResultNetworkPathHopJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocation struct {
	City  string                                                                     `json:"city,nullable"`
	State string                                                                     `json:"state,nullable"`
	Zip   string                                                                     `json:"zip,nullable"`
	JSON  digitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocationJSON `json:"-"`
}

// digitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocationJSON
// contains the JSON metadata for the struct
// [DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocation]
type digitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitoringTracerouteTestResultNetworkPathHopsLocationJSON) RawJSON() string {
	return r.raw
}

type DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMile string

const (
	DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMileClientToApp       DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMile = "client-to-app"
	DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMileClientToCfEgress  DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMile = "client-to-cf-egress"
	DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMileClientToCfIngress DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMile = "client-to-cf-ingress"
	DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMileClientToIsp       DigitalExperienceMonitoringTracerouteTestResultNetworkPathHopsMile = "client-to-isp"
)

type DEXTracerouteTestResultNetworkPathListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DEXTracerouteTestResultNetworkPathListResponseEnvelope struct {
	Errors   []DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages `json:"messages,required"`
	Result   DigitalExperienceMonitoringTracerouteTestResultNetworkPath       `json:"result,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON contains the JSON
// metadata for the struct [DEXTracerouteTestResultNetworkPathListResponseEnvelope]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccessTrue DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess = true
)
