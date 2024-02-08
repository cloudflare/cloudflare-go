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

// DexColoService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDexColoService] method instead.
type DexColoService struct {
	Options []option.RequestOption
}

// NewDexColoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDexColoService(opts ...option.RequestOption) (r *DexColoService) {
	r = &DexColoService{}
	r.Options = opts
	return
}

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *DexColoService) List(ctx context.Context, accountID string, query DexColoListParams, opts ...option.RequestOption) (res *[]DexColoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DexColoListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/colos", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexColoListResponse = interface{}

type DexColoListParams struct {
	// End time for connection period in RFC3339 (ISO 8601) format.
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for connection period in RFC3339 (ISO 8601) format.
	TimeStart param.Field[string] `query:"timeStart,required"`
	// Type of usage that colos should be sorted by. If unspecified, returns all
	// Cloudflare colos sorted alphabetically.
	SortBy param.Field[DexColoListParamsSortBy] `query:"sortBy"`
}

// URLQuery serializes [DexColoListParams]'s query parameters as `url.Values`.
func (r DexColoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of usage that colos should be sorted by. If unspecified, returns all
// Cloudflare colos sorted alphabetically.
type DexColoListParamsSortBy string

const (
	DexColoListParamsSortByFleetStatusUsage      DexColoListParamsSortBy = "fleet-status-usage"
	DexColoListParamsSortByApplicationTestsUsage DexColoListParamsSortBy = "application-tests-usage"
)

type DexColoListResponseEnvelope struct {
	Errors   []DexColoListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexColoListResponseEnvelopeMessages `json:"messages,required"`
	// array of colos.
	Result []DexColoListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DexColoListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DexColoListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexColoListResponseEnvelopeJSON       `json:"-"`
}

// dexColoListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DexColoListResponseEnvelope]
type dexColoListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexColoListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexColoListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dexColoListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexColoListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DexColoListResponseEnvelopeErrors]
type dexColoListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexColoListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexColoListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dexColoListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexColoListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DexColoListResponseEnvelopeMessages]
type dexColoListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexColoListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexColoListResponseEnvelopeSuccess bool

const (
	DexColoListResponseEnvelopeSuccessTrue DexColoListResponseEnvelopeSuccess = true
)

type DexColoListResponseEnvelopeResultInfo struct {
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
// struct [DexColoListResponseEnvelopeResultInfo]
type dexColoListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexColoListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
