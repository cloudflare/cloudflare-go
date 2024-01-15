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

// AccountIntelIPListService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelIPListService] method
// instead.
type AccountIntelIPListService struct {
	Options []option.RequestOption
}

// NewAccountIntelIPListService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIntelIPListService(opts ...option.RequestOption) (r *AccountIntelIPListService) {
	r = &AccountIntelIPListService{}
	r.Options = opts
	return
}

// Get IP Lists
func (r *AccountIntelIPListService) IPListGetIPLists(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountIntelIPListIPListGetIPListsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/ip-list", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountIntelIPListIPListGetIPListsResponse struct {
	Errors     []AccountIntelIPListIPListGetIPListsResponseError    `json:"errors"`
	Messages   []AccountIntelIPListIPListGetIPListsResponseMessage  `json:"messages"`
	Result     []AccountIntelIPListIPListGetIPListsResponseResult   `json:"result"`
	ResultInfo AccountIntelIPListIPListGetIPListsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountIntelIPListIPListGetIPListsResponseSuccess `json:"success"`
	JSON    accountIntelIPListIPListGetIPListsResponseJSON    `json:"-"`
}

// accountIntelIPListIPListGetIPListsResponseJSON contains the JSON metadata for
// the struct [AccountIntelIPListIPListGetIPListsResponse]
type accountIntelIPListIPListGetIPListsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIPListIPListGetIPListsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIPListIPListGetIPListsResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountIntelIPListIPListGetIPListsResponseErrorJSON `json:"-"`
}

// accountIntelIPListIPListGetIPListsResponseErrorJSON contains the JSON metadata
// for the struct [AccountIntelIPListIPListGetIPListsResponseError]
type accountIntelIPListIPListGetIPListsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIPListIPListGetIPListsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIPListIPListGetIPListsResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountIntelIPListIPListGetIPListsResponseMessageJSON `json:"-"`
}

// accountIntelIPListIPListGetIPListsResponseMessageJSON contains the JSON metadata
// for the struct [AccountIntelIPListIPListGetIPListsResponseMessage]
type accountIntelIPListIPListGetIPListsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIPListIPListGetIPListsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIPListIPListGetIPListsResponseResult struct {
	ID          int64                                                `json:"id"`
	Description string                                               `json:"description"`
	Name        string                                               `json:"name"`
	JSON        accountIntelIPListIPListGetIPListsResponseResultJSON `json:"-"`
}

// accountIntelIPListIPListGetIPListsResponseResultJSON contains the JSON metadata
// for the struct [AccountIntelIPListIPListGetIPListsResponseResult]
type accountIntelIPListIPListGetIPListsResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIPListIPListGetIPListsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIPListIPListGetIPListsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       accountIntelIPListIPListGetIPListsResponseResultInfoJSON `json:"-"`
}

// accountIntelIPListIPListGetIPListsResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountIntelIPListIPListGetIPListsResponseResultInfo]
type accountIntelIPListIPListGetIPListsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIPListIPListGetIPListsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelIPListIPListGetIPListsResponseSuccess bool

const (
	AccountIntelIPListIPListGetIPListsResponseSuccessTrue AccountIntelIPListIPListGetIPListsResponseSuccess = true
)
