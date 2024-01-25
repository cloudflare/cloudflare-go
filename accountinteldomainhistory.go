// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountIntelDomainHistoryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountIntelDomainHistoryService] method instead.
type AccountIntelDomainHistoryService struct {
	Options []option.RequestOption
}

// NewAccountIntelDomainHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountIntelDomainHistoryService(opts ...option.RequestOption) (r *AccountIntelDomainHistoryService) {
	r = &AccountIntelDomainHistoryService{}
	r.Options = opts
	return
}

// Get Domain History
func (r *AccountIntelDomainHistoryService) List(ctx context.Context, accountIdentifier string, query AccountIntelDomainHistoryListParams, opts ...option.RequestOption) (res *AccountIntelDomainHistoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/domain-history", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountIntelDomainHistoryListResponse struct {
	Errors     []AccountIntelDomainHistoryListResponseError    `json:"errors"`
	Messages   []AccountIntelDomainHistoryListResponseMessage  `json:"messages"`
	Result     []AccountIntelDomainHistoryListResponseResult   `json:"result"`
	ResultInfo AccountIntelDomainHistoryListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountIntelDomainHistoryListResponseSuccess `json:"success"`
	JSON    accountIntelDomainHistoryListResponseJSON    `json:"-"`
}

// accountIntelDomainHistoryListResponseJSON contains the JSON metadata for the
// struct [AccountIntelDomainHistoryListResponse]
type accountIntelDomainHistoryListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainHistoryListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountIntelDomainHistoryListResponseErrorJSON `json:"-"`
}

// accountIntelDomainHistoryListResponseErrorJSON contains the JSON metadata for
// the struct [AccountIntelDomainHistoryListResponseError]
type accountIntelDomainHistoryListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainHistoryListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainHistoryListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountIntelDomainHistoryListResponseMessageJSON `json:"-"`
}

// accountIntelDomainHistoryListResponseMessageJSON contains the JSON metadata for
// the struct [AccountIntelDomainHistoryListResponseMessage]
type accountIntelDomainHistoryListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainHistoryListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainHistoryListResponseResult struct {
	Categorizations []AccountIntelDomainHistoryListResponseResultCategorization `json:"categorizations"`
	Domain          string                                                      `json:"domain"`
	JSON            accountIntelDomainHistoryListResponseResultJSON             `json:"-"`
}

// accountIntelDomainHistoryListResponseResultJSON contains the JSON metadata for
// the struct [AccountIntelDomainHistoryListResponseResult]
type accountIntelDomainHistoryListResponseResultJSON struct {
	Categorizations apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelDomainHistoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainHistoryListResponseResultCategorization struct {
	Categories interface{}                                                   `json:"categories"`
	End        time.Time                                                     `json:"end" format:"date"`
	Start      time.Time                                                     `json:"start" format:"date"`
	JSON       accountIntelDomainHistoryListResponseResultCategorizationJSON `json:"-"`
}

// accountIntelDomainHistoryListResponseResultCategorizationJSON contains the JSON
// metadata for the struct
// [AccountIntelDomainHistoryListResponseResultCategorization]
type accountIntelDomainHistoryListResponseResultCategorizationJSON struct {
	Categories  apijson.Field
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainHistoryListResponseResultCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainHistoryListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       accountIntelDomainHistoryListResponseResultInfoJSON `json:"-"`
}

// accountIntelDomainHistoryListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountIntelDomainHistoryListResponseResultInfo]
type accountIntelDomainHistoryListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainHistoryListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelDomainHistoryListResponseSuccess bool

const (
	AccountIntelDomainHistoryListResponseSuccessTrue AccountIntelDomainHistoryListResponseSuccess = true
)

type AccountIntelDomainHistoryListParams struct {
	Domain param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes [AccountIntelDomainHistoryListParams]'s query parameters as
// `url.Values`.
func (r AccountIntelDomainHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
