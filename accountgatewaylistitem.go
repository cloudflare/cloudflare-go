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

// Fetch all items of a single Zero Trust List.
func (r *AccountGatewayListItemService) ZeroTrustListsZeroTrustListItems(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *ListItemResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s/items", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ListItemResponseCollection struct {
	Errors     []ListItemResponseCollectionError    `json:"errors"`
	Messages   []ListItemResponseCollectionMessage  `json:"messages"`
	Result     [][]ListItemResponseCollectionResult `json:"result"`
	ResultInfo ListItemResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ListItemResponseCollectionSuccess `json:"success"`
	JSON    listItemResponseCollectionJSON    `json:"-"`
}

// listItemResponseCollectionJSON contains the JSON metadata for the struct
// [ListItemResponseCollection]
type listItemResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListItemResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    listItemResponseCollectionErrorJSON `json:"-"`
}

// listItemResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ListItemResponseCollectionError]
type listItemResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListItemResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    listItemResponseCollectionMessageJSON `json:"-"`
}

// listItemResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ListItemResponseCollectionMessage]
type listItemResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListItemResponseCollectionResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a List.
	Value string                               `json:"value"`
	JSON  listItemResponseCollectionResultJSON `json:"-"`
}

// listItemResponseCollectionResultJSON contains the JSON metadata for the struct
// [ListItemResponseCollectionResult]
type listItemResponseCollectionResultJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListItemResponseCollectionResultInfo struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                  `json:"total_count"`
	JSON       listItemResponseCollectionResultInfoJSON `json:"-"`
}

// listItemResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [ListItemResponseCollectionResultInfo]
type listItemResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ListItemResponseCollectionSuccess bool

const (
	ListItemResponseCollectionSuccessTrue ListItemResponseCollectionSuccess = true
)
