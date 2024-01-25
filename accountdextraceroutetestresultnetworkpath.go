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

// AccountDexTracerouteTestResultNetworkPathService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexTracerouteTestResultNetworkPathService] method instead.
type AccountDexTracerouteTestResultNetworkPathService struct {
	Options []option.RequestOption
}

// NewAccountDexTracerouteTestResultNetworkPathService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountDexTracerouteTestResultNetworkPathService(opts ...option.RequestOption) (r *AccountDexTracerouteTestResultNetworkPathService) {
	r = &AccountDexTracerouteTestResultNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *AccountDexTracerouteTestResultNetworkPathService) List(ctx context.Context, accountIdentifier string, testResultID string, opts ...option.RequestOption) (res *AccountDexTracerouteTestResultNetworkPathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", accountIdentifier, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDexTracerouteTestResultNetworkPathListResponse struct {
	Errors   []AccountDexTracerouteTestResultNetworkPathListResponseError   `json:"errors"`
	Messages []AccountDexTracerouteTestResultNetworkPathListResponseMessage `json:"messages"`
	Result   AccountDexTracerouteTestResultNetworkPathListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexTracerouteTestResultNetworkPathListResponseSuccess `json:"success"`
	JSON    accountDexTracerouteTestResultNetworkPathListResponseJSON    `json:"-"`
}

// accountDexTracerouteTestResultNetworkPathListResponseJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestResultNetworkPathListResponse]
type accountDexTracerouteTestResultNetworkPathListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultNetworkPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestResultNetworkPathListResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountDexTracerouteTestResultNetworkPathListResponseErrorJSON `json:"-"`
}

// accountDexTracerouteTestResultNetworkPathListResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountDexTracerouteTestResultNetworkPathListResponseError]
type accountDexTracerouteTestResultNetworkPathListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultNetworkPathListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestResultNetworkPathListResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountDexTracerouteTestResultNetworkPathListResponseMessageJSON `json:"-"`
}

// accountDexTracerouteTestResultNetworkPathListResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestResultNetworkPathListResponseMessage]
type accountDexTracerouteTestResultNetworkPathListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultNetworkPathListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestResultNetworkPathListResponseResult struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []AccountDexTracerouteTestResultNetworkPathListResponseResultHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                                          `json:"testName"`
	JSON     accountDexTracerouteTestResultNetworkPathListResponseResultJSON `json:"-"`
}

// accountDexTracerouteTestResultNetworkPathListResponseResultJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestResultNetworkPathListResponseResult]
type accountDexTracerouteTestResultNetworkPathListResponseResultJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultNetworkPathListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestResultNetworkPathListResponseResultHop struct {
	Ttl           int64                                                                   `json:"ttl,required"`
	ASN           int64                                                                   `json:"asn,nullable"`
	Aso           string                                                                  `json:"aso,nullable"`
	IPAddress     string                                                                  `json:"ipAddress,nullable"`
	Location      AccountDexTracerouteTestResultNetworkPathListResponseResultHopsLocation `json:"location,nullable"`
	Mile          AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMile     `json:"mile,nullable"`
	Name          string                                                                  `json:"name,nullable"`
	PacketLossPct float64                                                                 `json:"packetLossPct,nullable"`
	RttMs         int64                                                                   `json:"rttMs,nullable"`
	JSON          accountDexTracerouteTestResultNetworkPathListResponseResultHopJSON      `json:"-"`
}

// accountDexTracerouteTestResultNetworkPathListResponseResultHopJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestResultNetworkPathListResponseResultHop]
type accountDexTracerouteTestResultNetworkPathListResponseResultHopJSON struct {
	Ttl           apijson.Field
	ASN           apijson.Field
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

func (r *AccountDexTracerouteTestResultNetworkPathListResponseResultHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestResultNetworkPathListResponseResultHopsLocation struct {
	City  string                                                                      `json:"city,nullable"`
	State string                                                                      `json:"state,nullable"`
	Zip   string                                                                      `json:"zip,nullable"`
	JSON  accountDexTracerouteTestResultNetworkPathListResponseResultHopsLocationJSON `json:"-"`
}

// accountDexTracerouteTestResultNetworkPathListResponseResultHopsLocationJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestResultNetworkPathListResponseResultHopsLocation]
type accountDexTracerouteTestResultNetworkPathListResponseResultHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultNetworkPathListResponseResultHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMile string

const (
	AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToApp       AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-app"
	AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToCfEgress  AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-cf-egress"
	AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToCfIngress AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-cf-ingress"
	AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMileClientToIsp       AccountDexTracerouteTestResultNetworkPathListResponseResultHopsMile = "client-to-isp"
)

// Whether the API call was successful
type AccountDexTracerouteTestResultNetworkPathListResponseSuccess bool

const (
	AccountDexTracerouteTestResultNetworkPathListResponseSuccessTrue AccountDexTracerouteTestResultNetworkPathListResponseSuccess = true
)
