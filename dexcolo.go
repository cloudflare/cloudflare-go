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

// DEXColoService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDEXColoService] method instead.
type DEXColoService struct {
	Options []option.RequestOption
}

// NewDEXColoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXColoService(opts ...option.RequestOption) (r *DEXColoService) {
	r = &DEXColoService{}
	r.Options = opts
	return
}

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *DEXColoService) List(ctx context.Context, params DEXColoListParams, opts ...option.RequestOption) (res *[]DEXColoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXColoListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/colos", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXColoListResponse = interface{}

type DEXColoListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// End time for connection period in RFC3339 (ISO 8601) format.
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for connection period in RFC3339 (ISO 8601) format.
	TimeStart param.Field[string] `query:"timeStart,required"`
	// Type of usage that colos should be sorted by. If unspecified, returns all
	// Cloudflare colos sorted alphabetically.
	SortBy param.Field[DEXColoListParamsSortBy] `query:"sortBy"`
}

// URLQuery serializes [DEXColoListParams]'s query parameters as `url.Values`.
func (r DEXColoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of usage that colos should be sorted by. If unspecified, returns all
// Cloudflare colos sorted alphabetically.
type DEXColoListParamsSortBy string

const (
	DEXColoListParamsSortByFleetStatusUsage      DEXColoListParamsSortBy = "fleet-status-usage"
	DEXColoListParamsSortByApplicationTestsUsage DEXColoListParamsSortBy = "application-tests-usage"
)

type DEXColoListResponseEnvelope struct {
	Errors   []DEXColoListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXColoListResponseEnvelopeMessages `json:"messages,required"`
	// array of colos.
	Result []DEXColoListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DEXColoListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DEXColoListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexColoListResponseEnvelopeJSON       `json:"-"`
}

// dexColoListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXColoListResponseEnvelope]
type dexColoListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXColoListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXColoListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dexColoListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexColoListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DEXColoListResponseEnvelopeErrors]
type dexColoListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXColoListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXColoListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dexColoListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexColoListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DEXColoListResponseEnvelopeMessages]
type dexColoListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXColoListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DEXColoListResponseEnvelopeSuccess bool

const (
	DEXColoListResponseEnvelopeSuccessTrue DEXColoListResponseEnvelopeSuccess = true
)

type DEXColoListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       dexColoListResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexColoListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DEXColoListResponseEnvelopeResultInfo]
type dexColoListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXColoListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
