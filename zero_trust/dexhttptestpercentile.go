// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DEXHTTPTestPercentileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXHTTPTestPercentileService] method instead.
type DEXHTTPTestPercentileService struct {
	Options []option.RequestOption
}

// NewDEXHTTPTestPercentileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXHTTPTestPercentileService(opts ...option.RequestOption) (r *DEXHTTPTestPercentileService) {
	r = &DEXHTTPTestPercentileService{}
	r.Options = opts
	return
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *DEXHTTPTestPercentileService) Get(ctx context.Context, testID string, params DEXHTTPTestPercentileGetParams, opts ...option.RequestOption) (res *HTTPDetailsPercentiles, err error) {
	opts = append(r.Options[:], opts...)
	var env DexhttpTestPercentileGetResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPDetailsPercentiles struct {
	DNSResponseTimeMs    Percentiles                `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  Percentiles                `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs Percentiles                `json:"serverResponseTimeMs"`
	JSON                 httpDetailsPercentilesJSON `json:"-"`
}

// httpDetailsPercentilesJSON contains the JSON metadata for the struct
// [HTTPDetailsPercentiles]
type httpDetailsPercentilesJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HTTPDetailsPercentiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsPercentilesJSON) RawJSON() string {
	return r.raw
}

type TestStatOverTime struct {
	Slots []TestStatOverTimeSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                `json:"min,nullable"`
	JSON testStatOverTimeJSON `json:"-"`
}

// testStatOverTimeJSON contains the JSON metadata for the struct
// [TestStatOverTime]
type testStatOverTimeJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestStatOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testStatOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestStatOverTimeSlot struct {
	Timestamp string                   `json:"timestamp,required"`
	Value     int64                    `json:"value,required"`
	JSON      testStatOverTimeSlotJSON `json:"-"`
}

// testStatOverTimeSlotJSON contains the JSON metadata for the struct
// [TestStatOverTimeSlot]
type testStatOverTimeSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestStatOverTimeSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testStatOverTimeSlotJSON) RawJSON() string {
	return r.raw
}

type DEXHTTPTestPercentileGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Start time for aggregate metrics in ISO format
	From param.Field[string] `query:"from,required"`
	// End time for aggregate metrics in ISO format
	To param.Field[string] `query:"to,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [DEXHTTPTestPercentileGetParams]'s query parameters as
// `url.Values`.
func (r DEXHTTPTestPercentileGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DexhttpTestPercentileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DexhttpTestPercentileGetResponseEnvelopeSuccess `json:"success,required"`
	Result  HTTPDetailsPercentiles                          `json:"result"`
	JSON    dexhttpTestPercentileGetResponseEnvelopeJSON    `json:"-"`
}

// dexhttpTestPercentileGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexhttpTestPercentileGetResponseEnvelope]
type dexhttpTestPercentileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestPercentileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestPercentileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DexhttpTestPercentileGetResponseEnvelopeSuccess bool

const (
	DexhttpTestPercentileGetResponseEnvelopeSuccessTrue DexhttpTestPercentileGetResponseEnvelopeSuccess = true
)

func (r DexhttpTestPercentileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DexhttpTestPercentileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
