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

// AccountRuleListItemService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRuleListItemService]
// method instead.
type AccountRuleListItemService struct {
	Options []option.RequestOption
}

// NewAccountRuleListItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRuleListItemService(opts ...option.RequestOption) (r *AccountRuleListItemService) {
	r = &AccountRuleListItemService{}
	r.Options = opts
	return
}

// Fetches a list item in the list.
func (r *AccountRuleListItemService) Get(ctx context.Context, accountIdentifier string, listID string, itemID string, opts ...option.RequestOption) (res *ItemResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items/%s", accountIdentifier, listID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes one or more items from a list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](#lists-get-bulk-operation-status) endpoint with the
// returned `operation_id`.
func (r *AccountRuleListItemService) Delete(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListItemDeleteParams, opts ...option.RequestOption) (res *ListsAsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Appends new items to the list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](#lists-get-bulk-operation-status) endpoint with the
// returned `operation_id`.
func (r *AccountRuleListItemService) ListsNewListItems(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListItemListsNewListItemsParams, opts ...option.RequestOption) (res *ListsAsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all the items in the list.
func (r *AccountRuleListItemService) ListsGetListItems(ctx context.Context, accountIdentifier string, listID string, query AccountRuleListItemListsGetListItemsParams, opts ...option.RequestOption) (res *AccountRuleListItemListsGetListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes all existing items from the list and adds the provided items to the
// list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](#lists-get-bulk-operation-status) endpoint with the
// returned `operation_id`.
func (r *AccountRuleListItemService) ListsUpdateAllListItems(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListItemListsUpdateAllListItemsParams, opts ...option.RequestOption) (res *ListsAsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ItemResponseCollection struct {
	Errors     []ItemResponseCollectionError    `json:"errors"`
	Messages   []ItemResponseCollectionMessage  `json:"messages"`
	Result     ItemResponseCollectionResult     `json:"result"`
	ResultInfo ItemResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ItemResponseCollectionSuccess `json:"success"`
	JSON    itemResponseCollectionJSON    `json:"-"`
}

// itemResponseCollectionJSON contains the JSON metadata for the struct
// [ItemResponseCollection]
type itemResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    itemResponseCollectionErrorJSON `json:"-"`
}

// itemResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ItemResponseCollectionError]
type itemResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    itemResponseCollectionMessageJSON `json:"-"`
}

// itemResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ItemResponseCollectionMessage]
type itemResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemResponseCollectionResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname string `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect ItemResponseCollectionResultRedirect `json:"redirect"`
	JSON     itemResponseCollectionResultJSON     `json:"-"`
}

// itemResponseCollectionResultJSON contains the JSON metadata for the struct
// [ItemResponseCollectionResult]
type itemResponseCollectionResultJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedOn  apijson.Field
	Redirect    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type ItemResponseCollectionResultRedirect struct {
	SourceURL           string                                         `json:"source_url,required"`
	TargetURL           string                                         `json:"target_url,required"`
	IncludeSubdomains   bool                                           `json:"include_subdomains"`
	PreservePathSuffix  bool                                           `json:"preserve_path_suffix"`
	PreserveQueryString bool                                           `json:"preserve_query_string"`
	StatusCode          ItemResponseCollectionResultRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                           `json:"subpath_matching"`
	JSON                itemResponseCollectionResultRedirectJSON       `json:"-"`
}

// itemResponseCollectionResultRedirectJSON contains the JSON metadata for the
// struct [ItemResponseCollectionResultRedirect]
type itemResponseCollectionResultRedirectJSON struct {
	SourceURL           apijson.Field
	TargetURL           apijson.Field
	IncludeSubdomains   apijson.Field
	PreservePathSuffix  apijson.Field
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	SubpathMatching     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ItemResponseCollectionResultRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemResponseCollectionResultRedirectStatusCode int64

const (
	ItemResponseCollectionResultRedirectStatusCode301 ItemResponseCollectionResultRedirectStatusCode = 301
	ItemResponseCollectionResultRedirectStatusCode302 ItemResponseCollectionResultRedirectStatusCode = 302
	ItemResponseCollectionResultRedirectStatusCode307 ItemResponseCollectionResultRedirectStatusCode = 307
	ItemResponseCollectionResultRedirectStatusCode308 ItemResponseCollectionResultRedirectStatusCode = 308
)

type ItemResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       itemResponseCollectionResultInfoJSON `json:"-"`
}

// itemResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [ItemResponseCollectionResultInfo]
type itemResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ItemResponseCollectionSuccess bool

const (
	ItemResponseCollectionSuccessTrue ItemResponseCollectionSuccess = true
)

type ListsAsyncResponse struct {
	Errors     []ListsAsyncResponseError    `json:"errors"`
	Messages   []ListsAsyncResponseMessage  `json:"messages"`
	Result     ListsAsyncResponseResult     `json:"result"`
	ResultInfo ListsAsyncResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ListsAsyncResponseSuccess `json:"success"`
	JSON    listsAsyncResponseJSON    `json:"-"`
}

// listsAsyncResponseJSON contains the JSON metadata for the struct
// [ListsAsyncResponse]
type listsAsyncResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListsAsyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListsAsyncResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    listsAsyncResponseErrorJSON `json:"-"`
}

// listsAsyncResponseErrorJSON contains the JSON metadata for the struct
// [ListsAsyncResponseError]
type listsAsyncResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListsAsyncResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListsAsyncResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    listsAsyncResponseMessageJSON `json:"-"`
}

// listsAsyncResponseMessageJSON contains the JSON metadata for the struct
// [ListsAsyncResponseMessage]
type listsAsyncResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListsAsyncResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListsAsyncResponseResult struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                       `json:"operation_id"`
	JSON        listsAsyncResponseResultJSON `json:"-"`
}

// listsAsyncResponseResultJSON contains the JSON metadata for the struct
// [ListsAsyncResponseResult]
type listsAsyncResponseResultJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListsAsyncResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListsAsyncResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                          `json:"total_count"`
	JSON       listsAsyncResponseResultInfoJSON `json:"-"`
}

// listsAsyncResponseResultInfoJSON contains the JSON metadata for the struct
// [ListsAsyncResponseResultInfo]
type listsAsyncResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListsAsyncResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ListsAsyncResponseSuccess bool

const (
	ListsAsyncResponseSuccessTrue ListsAsyncResponseSuccess = true
)

type AccountRuleListItemListsGetListItemsResponse struct {
	Errors     []AccountRuleListItemListsGetListItemsResponseError    `json:"errors"`
	Messages   []AccountRuleListItemListsGetListItemsResponseMessage  `json:"messages"`
	Result     []AccountRuleListItemListsGetListItemsResponseResult   `json:"result"`
	ResultInfo AccountRuleListItemListsGetListItemsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountRuleListItemListsGetListItemsResponseSuccess `json:"success"`
	JSON    accountRuleListItemListsGetListItemsResponseJSON    `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseJSON contains the JSON metadata for
// the struct [AccountRuleListItemListsGetListItemsResponse]
type accountRuleListItemListsGetListItemsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountRuleListItemListsGetListItemsResponseErrorJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseErrorJSON contains the JSON metadata
// for the struct [AccountRuleListItemListsGetListItemsResponseError]
type accountRuleListItemListsGetListItemsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountRuleListItemListsGetListItemsResponseMessageJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseMessageJSON contains the JSON
// metadata for the struct [AccountRuleListItemListsGetListItemsResponseMessage]
type accountRuleListItemListsGetListItemsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname string `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemListsGetListItemsResponseResultRedirect `json:"redirect"`
	JSON     accountRuleListItemListsGetListItemsResponseResultJSON     `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultJSON contains the JSON
// metadata for the struct [AccountRuleListItemListsGetListItemsResponseResult]
type accountRuleListItemListsGetListItemsResponseResultJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedOn  apijson.Field
	Redirect    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemListsGetListItemsResponseResultRedirect struct {
	SourceURL           string                                                               `json:"source_url,required"`
	TargetURL           string                                                               `json:"target_url,required"`
	IncludeSubdomains   bool                                                                 `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                 `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                 `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                 `json:"subpath_matching"`
	JSON                accountRuleListItemListsGetListItemsResponseResultRedirectJSON       `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultRedirectJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultRedirect]
type accountRuleListItemListsGetListItemsResponseResultRedirectJSON struct {
	SourceURL           apijson.Field
	TargetURL           apijson.Field
	IncludeSubdomains   apijson.Field
	PreservePathSuffix  apijson.Field
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	SubpathMatching     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode int64

const (
	AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode301 AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode = 301
	AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode302 AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode = 302
	AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode307 AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode = 307
	AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode308 AccountRuleListItemListsGetListItemsResponseResultRedirectStatusCode = 308
)

type AccountRuleListItemListsGetListItemsResponseResultInfo struct {
	Cursors AccountRuleListItemListsGetListItemsResponseResultInfoCursors `json:"cursors"`
	JSON    accountRuleListItemListsGetListItemsResponseResultInfoJSON    `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountRuleListItemListsGetListItemsResponseResultInfo]
type accountRuleListItemListsGetListItemsResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResultInfoCursors struct {
	After  string                                                            `json:"after"`
	Before string                                                            `json:"before"`
	JSON   accountRuleListItemListsGetListItemsResponseResultInfoCursorsJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultInfoCursorsJSON contains the
// JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultInfoCursors]
type accountRuleListItemListsGetListItemsResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRuleListItemListsGetListItemsResponseSuccess bool

const (
	AccountRuleListItemListsGetListItemsResponseSuccessTrue AccountRuleListItemListsGetListItemsResponseSuccess = true
)

type AccountRuleListItemDeleteParams struct {
	Items param.Field[[]AccountRuleListItemDeleteParamsItem] `json:"items"`
}

func (r AccountRuleListItemDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListItemDeleteParamsItem struct {
	// The unique ID of the item in the List.
	ID param.Field[string] `json:"id"`
}

func (r AccountRuleListItemDeleteParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListItemListsNewListItemsParams struct {
	Body param.Field[[]AccountRuleListItemListsNewListItemsParamsBody] `json:"body,required"`
}

func (r AccountRuleListItemListsNewListItemsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRuleListItemListsNewListItemsParamsBody struct {
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[AccountRuleListItemListsNewListItemsParamsBodyRedirect] `json:"redirect"`
}

func (r AccountRuleListItemListsNewListItemsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type AccountRuleListItemListsNewListItemsParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                                           `json:"source_url,required"`
	TargetURL           param.Field[string]                                                           `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                                             `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                                             `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                                             `json:"preserve_query_string"`
	StatusCode          param.Field[AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                                             `json:"subpath_matching"`
}

func (r AccountRuleListItemListsNewListItemsParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode int64

const (
	AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode301 AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 301
	AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode302 AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 302
	AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode307 AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 307
	AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode308 AccountRuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 308
)

type AccountRuleListItemListsGetListItemsParams struct {
	// The pagination cursor. An opaque string token indicating the position from which
	// to continue when requesting the next/previous set of records. Cursor values are
	// provided under `result_info.cursors` in the response. You should make no
	// assumptions about a cursor's content or length.
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [AccountRuleListItemListsGetListItemsParams]'s query
// parameters as `url.Values`.
func (r AccountRuleListItemListsGetListItemsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountRuleListItemListsUpdateAllListItemsParams struct {
	Body param.Field[[]AccountRuleListItemListsUpdateAllListItemsParamsBody] `json:"body,required"`
}

func (r AccountRuleListItemListsUpdateAllListItemsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRuleListItemListsUpdateAllListItemsParamsBody struct {
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect] `json:"redirect"`
}

func (r AccountRuleListItemListsUpdateAllListItemsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                                                 `json:"source_url,required"`
	TargetURL           param.Field[string]                                                                 `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                                                   `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                                                   `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                                                   `json:"preserve_query_string"`
	StatusCode          param.Field[AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                                                   `json:"subpath_matching"`
}

func (r AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode int64

const (
	AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301 AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 301
	AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode302 AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 302
	AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode307 AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 307
	AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode308 AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 308
)
