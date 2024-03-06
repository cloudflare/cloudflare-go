// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDEXColoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDEXColoService] method
// instead.
type ZeroTrustDEXColoService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXColoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDEXColoService(opts ...option.RequestOption) (r *ZeroTrustDEXColoService) {
	r = &ZeroTrustDEXColoService{}
	r.Options = opts
	return
}

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *ZeroTrustDEXColoService) List(ctx context.Context, params ZeroTrustDEXColoListParams, opts ...option.RequestOption) (res *[]ZeroTrustDEXColoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXColoListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/colos", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDEXColoListResponse = interface{}

type ZeroTrustDEXColoListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// End time for connection period in RFC3339 (ISO 8601) format.
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for connection period in RFC3339 (ISO 8601) format.
	TimeStart param.Field[string] `query:"timeStart,required"`
	// Type of usage that colos should be sorted by. If unspecified, returns all
	// Cloudflare colos sorted alphabetically.
	SortBy param.Field[ZeroTrustDEXColoListParamsSortBy] `query:"sortBy"`
}

// URLQuery serializes [ZeroTrustDEXColoListParams]'s query parameters as
// `url.Values`.
func (r ZeroTrustDEXColoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of usage that colos should be sorted by. If unspecified, returns all
// Cloudflare colos sorted alphabetically.
type ZeroTrustDEXColoListParamsSortBy string

const (
	ZeroTrustDEXColoListParamsSortByFleetStatusUsage      ZeroTrustDEXColoListParamsSortBy = "fleet-status-usage"
	ZeroTrustDEXColoListParamsSortByApplicationTestsUsage ZeroTrustDEXColoListParamsSortBy = "application-tests-usage"
)

type ZeroTrustDEXColoListResponseEnvelope struct {
	Errors   []ZeroTrustDEXColoListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXColoListResponseEnvelopeMessages `json:"messages,required"`
	// array of colos.
	Result []ZeroTrustDEXColoListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustDEXColoListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDEXColoListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDEXColoListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDEXColoListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDEXColoListResponseEnvelope]
type zeroTrustDEXColoListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXColoListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXColoListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXColoListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zeroTrustDEXColoListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXColoListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDEXColoListResponseEnvelopeErrors]
type zeroTrustDEXColoListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXColoListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXColoListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXColoListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustDEXColoListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXColoListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustDEXColoListResponseEnvelopeMessages]
type zeroTrustDEXColoListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXColoListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXColoListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDEXColoListResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXColoListResponseEnvelopeSuccessTrue ZeroTrustDEXColoListResponseEnvelopeSuccess = true
)

type ZeroTrustDEXColoListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       zeroTrustDEXColoListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDEXColoListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustDEXColoListResponseEnvelopeResultInfo]
type zeroTrustDEXColoListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXColoListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXColoListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
