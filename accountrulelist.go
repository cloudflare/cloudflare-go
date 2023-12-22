// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRuleListService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRuleListService] method
// instead.
type AccountRuleListService struct {
	Options        []option.RequestOption
	BulkOperations *AccountRuleListBulkOperationService
	Items          *AccountRuleListItemService
}

// NewAccountRuleListService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRuleListService(opts ...option.RequestOption) (r *AccountRuleListService) {
	r = &AccountRuleListService{}
	r.Options = opts
	r.BulkOperations = NewAccountRuleListBulkOperationService(opts...)
	r.Items = NewAccountRuleListItemService(opts...)
	return
}

// Fetches the details of a list.
func (r *AccountRuleListService) Get(ctx context.Context, accountIdentifier string, listID string, opts ...option.RequestOption) (res *ListResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the description of a list.
func (r *AccountRuleListService) Update(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListUpdateParams, opts ...option.RequestOption) (res *ListResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a specific list and all its items.
func (r *AccountRuleListService) Delete(ctx context.Context, accountIdentifier string, listID string, opts ...option.RequestOption) (res *ListDeleteResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new list of the specified type.
func (r *AccountRuleListService) ListsNewAList(ctx context.Context, accountIdentifier string, body AccountRuleListListsNewAListParams, opts ...option.RequestOption) (res *ListResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all lists in the account.
func (r *AccountRuleListService) ListsGetLists(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountRuleListListsGetListsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ListDeleteResponseCollection struct {
	Errors     []ListDeleteResponseCollectionError    `json:"errors"`
	Messages   []ListDeleteResponseCollectionMessage  `json:"messages"`
	Result     ListDeleteResponseCollectionResult     `json:"result"`
	ResultInfo ListDeleteResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ListDeleteResponseCollectionSuccess `json:"success"`
	JSON    listDeleteResponseCollectionJSON    `json:"-"`
}

// listDeleteResponseCollectionJSON contains the JSON metadata for the struct
// [ListDeleteResponseCollection]
type listDeleteResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListDeleteResponseCollectionError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    listDeleteResponseCollectionErrorJSON `json:"-"`
}

// listDeleteResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ListDeleteResponseCollectionError]
type listDeleteResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListDeleteResponseCollectionMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    listDeleteResponseCollectionMessageJSON `json:"-"`
}

// listDeleteResponseCollectionMessageJSON contains the JSON metadata for the
// struct [ListDeleteResponseCollectionMessage]
type listDeleteResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListDeleteResponseCollectionResult struct {
	// The unique ID of the item in the List.
	ID   string                                 `json:"id"`
	JSON listDeleteResponseCollectionResultJSON `json:"-"`
}

// listDeleteResponseCollectionResultJSON contains the JSON metadata for the struct
// [ListDeleteResponseCollectionResult]
type listDeleteResponseCollectionResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListDeleteResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       listDeleteResponseCollectionResultInfoJSON `json:"-"`
}

// listDeleteResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [ListDeleteResponseCollectionResultInfo]
type listDeleteResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ListDeleteResponseCollectionSuccess bool

const (
	ListDeleteResponseCollectionSuccessTrue ListDeleteResponseCollectionSuccess = true
)

type ListResponseCollection struct {
	Errors     []ListResponseCollectionError    `json:"errors"`
	Messages   []ListResponseCollectionMessage  `json:"messages"`
	Result     ListResponseCollectionResult     `json:"result"`
	ResultInfo ListResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ListResponseCollectionSuccess `json:"success"`
	JSON    listResponseCollectionJSON    `json:"-"`
}

// listResponseCollectionJSON contains the JSON metadata for the struct
// [ListResponseCollection]
type listResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    listResponseCollectionErrorJSON `json:"-"`
}

// listResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ListResponseCollectionError]
type listResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    listResponseCollectionMessageJSON `json:"-"`
}

// listResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ListResponseCollectionMessage]
type listResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseCollectionResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// hostnames or redirects).
	Kind ListResponseCollectionResultKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](#filters) referencing the list.
	NumReferencingFilters float64                          `json:"num_referencing_filters"`
	JSON                  listResponseCollectionResultJSON `json:"-"`
}

// listResponseCollectionResultJSON contains the JSON metadata for the struct
// [ListResponseCollectionResult]
type listResponseCollectionResultJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	Kind                  apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NumItems              apijson.Field
	NumReferencingFilters apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ListResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// hostnames or redirects).
type ListResponseCollectionResultKind string

const (
	ListResponseCollectionResultKindIP       ListResponseCollectionResultKind = "ip"
	ListResponseCollectionResultKindRedirect ListResponseCollectionResultKind = "redirect"
	ListResponseCollectionResultKindHostname ListResponseCollectionResultKind = "hostname"
)

type ListResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       listResponseCollectionResultInfoJSON `json:"-"`
}

// listResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [ListResponseCollectionResultInfo]
type listResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ListResponseCollectionSuccess bool

const (
	ListResponseCollectionSuccessTrue ListResponseCollectionSuccess = true
)

type AccountRuleListListsGetListsResponse struct {
	Errors     []AccountRuleListListsGetListsResponseError    `json:"errors"`
	Messages   []AccountRuleListListsGetListsResponseMessage  `json:"messages"`
	Result     []AccountRuleListListsGetListsResponseResult   `json:"result"`
	ResultInfo AccountRuleListListsGetListsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountRuleListListsGetListsResponseSuccess `json:"success"`
	JSON    accountRuleListListsGetListsResponseJSON    `json:"-"`
}

// accountRuleListListsGetListsResponseJSON contains the JSON metadata for the
// struct [AccountRuleListListsGetListsResponse]
type accountRuleListListsGetListsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsGetListsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListListsGetListsResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountRuleListListsGetListsResponseErrorJSON `json:"-"`
}

// accountRuleListListsGetListsResponseErrorJSON contains the JSON metadata for the
// struct [AccountRuleListListsGetListsResponseError]
type accountRuleListListsGetListsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsGetListsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListListsGetListsResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountRuleListListsGetListsResponseMessageJSON `json:"-"`
}

// accountRuleListListsGetListsResponseMessageJSON contains the JSON metadata for
// the struct [AccountRuleListListsGetListsResponseMessage]
type accountRuleListListsGetListsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsGetListsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListListsGetListsResponseResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// hostnames or redirects).
	Kind AccountRuleListListsGetListsResponseResultKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](#filters) referencing the list.
	NumReferencingFilters float64                                        `json:"num_referencing_filters"`
	JSON                  accountRuleListListsGetListsResponseResultJSON `json:"-"`
}

// accountRuleListListsGetListsResponseResultJSON contains the JSON metadata for
// the struct [AccountRuleListListsGetListsResponseResult]
type accountRuleListListsGetListsResponseResultJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	Kind                  apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NumItems              apijson.Field
	NumReferencingFilters apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountRuleListListsGetListsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// hostnames or redirects).
type AccountRuleListListsGetListsResponseResultKind string

const (
	AccountRuleListListsGetListsResponseResultKindIP       AccountRuleListListsGetListsResponseResultKind = "ip"
	AccountRuleListListsGetListsResponseResultKindRedirect AccountRuleListListsGetListsResponseResultKind = "redirect"
	AccountRuleListListsGetListsResponseResultKindHostname AccountRuleListListsGetListsResponseResultKind = "hostname"
)

type AccountRuleListListsGetListsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       accountRuleListListsGetListsResponseResultInfoJSON `json:"-"`
}

// accountRuleListListsGetListsResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountRuleListListsGetListsResponseResultInfo]
type accountRuleListListsGetListsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsGetListsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRuleListListsGetListsResponseSuccess bool

const (
	AccountRuleListListsGetListsResponseSuccessTrue AccountRuleListListsGetListsResponseSuccess = true
)

type AccountRuleListUpdateParams struct {
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r AccountRuleListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListListsNewAListParams struct {
	// The type of the list. Each type supports specific list items (IP addresses,
	// hostnames or redirects).
	Kind param.Field[AccountRuleListListsNewAListParamsKind] `json:"kind,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name param.Field[string] `json:"name,required"`
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r AccountRuleListListsNewAListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// hostnames or redirects).
type AccountRuleListListsNewAListParamsKind string

const (
	AccountRuleListListsNewAListParamsKindIP       AccountRuleListListsNewAListParamsKind = "ip"
	AccountRuleListListsNewAListParamsKindRedirect AccountRuleListListsNewAListParamsKind = "redirect"
	AccountRuleListListsNewAListParamsKindHostname AccountRuleListListsNewAListParamsKind = "hostname"
)
