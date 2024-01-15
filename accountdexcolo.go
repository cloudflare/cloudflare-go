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

// AccountDexColoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDexColoService] method
// instead.
type AccountDexColoService struct {
	Options []option.RequestOption
}

// NewAccountDexColoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDexColoService(opts ...option.RequestOption) (r *AccountDexColoService) {
	r = &AccountDexColoService{}
	r.Options = opts
	return
}

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *AccountDexColoService) List(ctx context.Context, accountIdentifier string, query AccountDexColoListParams, opts ...option.RequestOption) (res *AccountDexColoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/colos", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexColoListResponse struct {
	Errors   []AccountDexColoListResponseError   `json:"errors"`
	Messages []AccountDexColoListResponseMessage `json:"messages"`
	// array of colos.
	Result     []AccountDexColoListResponseResult   `json:"result"`
	ResultInfo AccountDexColoListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountDexColoListResponseSuccess `json:"success"`
	JSON    accountDexColoListResponseJSON    `json:"-"`
}

// accountDexColoListResponseJSON contains the JSON metadata for the struct
// [AccountDexColoListResponse]
type accountDexColoListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexColoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexColoListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountDexColoListResponseErrorJSON `json:"-"`
}

// accountDexColoListResponseErrorJSON contains the JSON metadata for the struct
// [AccountDexColoListResponseError]
type accountDexColoListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexColoListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexColoListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountDexColoListResponseMessageJSON `json:"-"`
}

// accountDexColoListResponseMessageJSON contains the JSON metadata for the struct
// [AccountDexColoListResponseMessage]
type accountDexColoListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexColoListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexColoListResponseResult struct {
	// Airport code
	AirportCode string `json:"airportCode,required"`
	// City
	City string `json:"city,required"`
	// Country code
	CountryCode string                               `json:"countryCode,required"`
	JSON        accountDexColoListResponseResultJSON `json:"-"`
}

// accountDexColoListResponseResultJSON contains the JSON metadata for the struct
// [AccountDexColoListResponseResult]
type accountDexColoListResponseResultJSON struct {
	AirportCode apijson.Field
	City        apijson.Field
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexColoListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexColoListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       accountDexColoListResponseResultInfoJSON `json:"-"`
}

// accountDexColoListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDexColoListResponseResultInfo]
type accountDexColoListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexColoListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexColoListResponseSuccess bool

const (
	AccountDexColoListResponseSuccessTrue AccountDexColoListResponseSuccess = true
)

type AccountDexColoListParams struct {
	// End time for connection period in RFC3339 (ISO 8601) format.
	TimeEnd param.Field[string] `query:"timeEnd,required"`
	// Start time for connection period in RFC3339 (ISO 8601) format.
	TimeStart param.Field[string] `query:"timeStart,required"`
	// Type of usage that colos should be sorted by. If unspecified, returns all
	// Cloudflare colos sorted alphabetically.
	SortBy param.Field[AccountDexColoListParamsSortBy] `query:"sortBy"`
}

// URLQuery serializes [AccountDexColoListParams]'s query parameters as
// `url.Values`.
func (r AccountDexColoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of usage that colos should be sorted by. If unspecified, returns all
// Cloudflare colos sorted alphabetically.
type AccountDexColoListParamsSortBy string

const (
	AccountDexColoListParamsSortByFleetStatusUsage      AccountDexColoListParamsSortBy = "fleet-status-usage"
	AccountDexColoListParamsSortByApplicationTestsUsage AccountDexColoListParamsSortBy = "application-tests-usage"
)
