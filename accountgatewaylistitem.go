// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayListItemService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayListItemService]
// method instead.
type AccountGatewayListItemService struct {
	Options []option.RequestOption
}

// NewAccountGatewayListItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayListItemService(opts ...option.RequestOption) (r *AccountGatewayListItemService) {
	r = &AccountGatewayListItemService{}
	r.Options = opts
	return
}

// Fetches all items in a single Zero Trust list.
func (r *AccountGatewayListItemService) ZeroTrustListsZeroTrustListItems(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s/items", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponse struct {
	Errors     []AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseError    `json:"errors"`
	Messages   []AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessage  `json:"messages"`
	Result     [][]AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResult `json:"result"`
	ResultInfo AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseSuccess `json:"success"`
	JSON    accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseJSON    `json:"-"`
}

// accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseJSON contains the
// JSON metadata for the struct
// [AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponse]
type accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseErrorJSON `json:"-"`
}

// accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseError]
type accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessageJSON `json:"-"`
}

// accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessage]
type accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                                                                   `json:"value"`
	JSON  accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultJSON `json:"-"`
}

// accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResult]
type accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfo struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                                      `json:"total_count"`
	JSON       accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfoJSON `json:"-"`
}

// accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfo]
type accountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseSuccess bool

const (
	AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseSuccessTrue AccountGatewayListItemZeroTrustListsZeroTrustListItemsResponseSuccess = true
)
