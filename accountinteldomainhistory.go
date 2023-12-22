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
func (r *AccountIntelDomainHistoryService) List(ctx context.Context, accountIdentifier string, query AccountIntelDomainHistoryListParams, opts ...option.RequestOption) (res *DomainHistory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/domain-history", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DomainHistory struct {
	Errors     []DomainHistoryError    `json:"errors"`
	Messages   []DomainHistoryMessage  `json:"messages"`
	Result     []DomainHistoryResult   `json:"result"`
	ResultInfo DomainHistoryResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DomainHistorySuccess `json:"success"`
	JSON    domainHistoryJSON    `json:"-"`
}

// domainHistoryJSON contains the JSON metadata for the struct [DomainHistory]
type domainHistoryJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainHistoryError struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    domainHistoryErrorJSON `json:"-"`
}

// domainHistoryErrorJSON contains the JSON metadata for the struct
// [DomainHistoryError]
type domainHistoryErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainHistoryMessage struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    domainHistoryMessageJSON `json:"-"`
}

// domainHistoryMessageJSON contains the JSON metadata for the struct
// [DomainHistoryMessage]
type domainHistoryMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainHistoryResult struct {
	Categorizations []DomainHistoryResultCategorization `json:"categorizations"`
	Domain          string                              `json:"domain"`
	JSON            domainHistoryResultJSON             `json:"-"`
}

// domainHistoryResultJSON contains the JSON metadata for the struct
// [DomainHistoryResult]
type domainHistoryResultJSON struct {
	Categorizations apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainHistoryResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainHistoryResultCategorization struct {
	Categories interface{}                           `json:"categories"`
	End        time.Time                             `json:"end" format:"date"`
	Start      time.Time                             `json:"start" format:"date"`
	JSON       domainHistoryResultCategorizationJSON `json:"-"`
}

// domainHistoryResultCategorizationJSON contains the JSON metadata for the struct
// [DomainHistoryResultCategorization]
type domainHistoryResultCategorizationJSON struct {
	Categories  apijson.Field
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryResultCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainHistoryResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                     `json:"total_count"`
	JSON       domainHistoryResultInfoJSON `json:"-"`
}

// domainHistoryResultInfoJSON contains the JSON metadata for the struct
// [DomainHistoryResultInfo]
type domainHistoryResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DomainHistorySuccess bool

const (
	DomainHistorySuccessTrue DomainHistorySuccess = true
)

type AccountIntelDomainHistoryListParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [AccountIntelDomainHistoryListParams]'s query parameters as
// `url.Values`.
func (r AccountIntelDomainHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
