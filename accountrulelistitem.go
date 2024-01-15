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
func (r *AccountRuleListItemService) Get(ctx context.Context, accountIdentifier string, listID string, itemID string, opts ...option.RequestOption) (res *AccountRuleListItemGetResponse, err error) {
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
func (r *AccountRuleListItemService) Delete(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListItemDeleteParams, opts ...option.RequestOption) (res *AccountRuleListItemDeleteResponse, err error) {
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
func (r *AccountRuleListItemService) ListsNewListItems(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListItemListsNewListItemsParams, opts ...option.RequestOption) (res *AccountRuleListItemListsNewListItemsResponse, err error) {
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
func (r *AccountRuleListItemService) ListsUpdateAllListItems(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListItemListsUpdateAllListItemsParams, opts ...option.RequestOption) (res *AccountRuleListItemListsUpdateAllListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountRuleListItemGetResponse struct {
	Errors   []AccountRuleListItemGetResponseError   `json:"errors"`
	Messages []AccountRuleListItemGetResponseMessage `json:"messages"`
	Result   AccountRuleListItemGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListItemGetResponseSuccess `json:"success"`
	JSON    accountRuleListItemGetResponseJSON    `json:"-"`
}

// accountRuleListItemGetResponseJSON contains the JSON metadata for the struct
// [AccountRuleListItemGetResponse]
type accountRuleListItemGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountRuleListItemGetResponseErrorJSON `json:"-"`
}

// accountRuleListItemGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountRuleListItemGetResponseError]
type accountRuleListItemGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountRuleListItemGetResponseMessageJSON `json:"-"`
}

// accountRuleListItemGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountRuleListItemGetResponseMessage]
type accountRuleListItemGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// A non-negative 32 bit integer
	ASN int64 `json:"asn"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname AccountRuleListItemGetResponseResultHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemGetResponseResultRedirect `json:"redirect"`
	JSON     accountRuleListItemGetResponseResultJSON     `json:"-"`
}

// accountRuleListItemGetResponseResultJSON contains the JSON metadata for the
// struct [AccountRuleListItemGetResponseResult]
type accountRuleListItemGetResponseResultJSON struct {
	ID          apijson.Field
	ASN         apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedOn  apijson.Field
	Redirect    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemGetResponseResultHostname struct {
	URLHostname string                                           `json:"url_hostname,required"`
	JSON        accountRuleListItemGetResponseResultHostnameJSON `json:"-"`
}

// accountRuleListItemGetResponseResultHostnameJSON contains the JSON metadata for
// the struct [AccountRuleListItemGetResponseResultHostname]
type accountRuleListItemGetResponseResultHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseResultHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemGetResponseResultRedirect struct {
	SourceURL           string                                                 `json:"source_url,required"`
	TargetURL           string                                                 `json:"target_url,required"`
	IncludeSubdomains   bool                                                   `json:"include_subdomains"`
	PreservePathSuffix  bool                                                   `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                   `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemGetResponseResultRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                   `json:"subpath_matching"`
	JSON                accountRuleListItemGetResponseResultRedirectJSON       `json:"-"`
}

// accountRuleListItemGetResponseResultRedirectJSON contains the JSON metadata for
// the struct [AccountRuleListItemGetResponseResultRedirect]
type accountRuleListItemGetResponseResultRedirectJSON struct {
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

func (r *AccountRuleListItemGetResponseResultRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseResultRedirectStatusCode int64

const (
	AccountRuleListItemGetResponseResultRedirectStatusCode301 AccountRuleListItemGetResponseResultRedirectStatusCode = 301
	AccountRuleListItemGetResponseResultRedirectStatusCode302 AccountRuleListItemGetResponseResultRedirectStatusCode = 302
	AccountRuleListItemGetResponseResultRedirectStatusCode307 AccountRuleListItemGetResponseResultRedirectStatusCode = 307
	AccountRuleListItemGetResponseResultRedirectStatusCode308 AccountRuleListItemGetResponseResultRedirectStatusCode = 308
)

// Whether the API call was successful
type AccountRuleListItemGetResponseSuccess bool

const (
	AccountRuleListItemGetResponseSuccessTrue AccountRuleListItemGetResponseSuccess = true
)

type AccountRuleListItemDeleteResponse struct {
	Errors   []AccountRuleListItemDeleteResponseError   `json:"errors"`
	Messages []AccountRuleListItemDeleteResponseMessage `json:"messages"`
	Result   AccountRuleListItemDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListItemDeleteResponseSuccess `json:"success"`
	JSON    accountRuleListItemDeleteResponseJSON    `json:"-"`
}

// accountRuleListItemDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRuleListItemDeleteResponse]
type accountRuleListItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountRuleListItemDeleteResponseErrorJSON `json:"-"`
}

// accountRuleListItemDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountRuleListItemDeleteResponseError]
type accountRuleListItemDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountRuleListItemDeleteResponseMessageJSON `json:"-"`
}

// accountRuleListItemDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountRuleListItemDeleteResponseMessage]
type accountRuleListItemDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemDeleteResponseResult struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                      `json:"operation_id"`
	JSON        accountRuleListItemDeleteResponseResultJSON `json:"-"`
}

// accountRuleListItemDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountRuleListItemDeleteResponseResult]
type accountRuleListItemDeleteResponseResultJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRuleListItemDeleteResponseSuccess bool

const (
	AccountRuleListItemDeleteResponseSuccessTrue AccountRuleListItemDeleteResponseSuccess = true
)

type AccountRuleListItemListsNewListItemsResponse struct {
	Errors   []AccountRuleListItemListsNewListItemsResponseError   `json:"errors"`
	Messages []AccountRuleListItemListsNewListItemsResponseMessage `json:"messages"`
	Result   AccountRuleListItemListsNewListItemsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListItemListsNewListItemsResponseSuccess `json:"success"`
	JSON    accountRuleListItemListsNewListItemsResponseJSON    `json:"-"`
}

// accountRuleListItemListsNewListItemsResponseJSON contains the JSON metadata for
// the struct [AccountRuleListItemListsNewListItemsResponse]
type accountRuleListItemListsNewListItemsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsNewListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsNewListItemsResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountRuleListItemListsNewListItemsResponseErrorJSON `json:"-"`
}

// accountRuleListItemListsNewListItemsResponseErrorJSON contains the JSON metadata
// for the struct [AccountRuleListItemListsNewListItemsResponseError]
type accountRuleListItemListsNewListItemsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsNewListItemsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsNewListItemsResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountRuleListItemListsNewListItemsResponseMessageJSON `json:"-"`
}

// accountRuleListItemListsNewListItemsResponseMessageJSON contains the JSON
// metadata for the struct [AccountRuleListItemListsNewListItemsResponseMessage]
type accountRuleListItemListsNewListItemsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsNewListItemsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsNewListItemsResponseResult struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                                 `json:"operation_id"`
	JSON        accountRuleListItemListsNewListItemsResponseResultJSON `json:"-"`
}

// accountRuleListItemListsNewListItemsResponseResultJSON contains the JSON
// metadata for the struct [AccountRuleListItemListsNewListItemsResponseResult]
type accountRuleListItemListsNewListItemsResponseResultJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsNewListItemsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRuleListItemListsNewListItemsResponseSuccess bool

const (
	AccountRuleListItemListsNewListItemsResponseSuccessTrue AccountRuleListItemListsNewListItemsResponseSuccess = true
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
	// A non-negative 32 bit integer
	ASN int64 `json:"asn"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname AccountRuleListItemListsGetListItemsResponseResultHostname `json:"hostname"`
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
	ASN         apijson.Field
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

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsGetListItemsResponseResultHostname struct {
	URLHostname string                                                         `json:"url_hostname,required"`
	JSON        accountRuleListItemListsGetListItemsResponseResultHostnameJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultHostnameJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultHostname]
type accountRuleListItemListsGetListItemsResponseResultHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultHostname) UnmarshalJSON(data []byte) (err error) {
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

type AccountRuleListItemListsUpdateAllListItemsResponse struct {
	Errors   []AccountRuleListItemListsUpdateAllListItemsResponseError   `json:"errors"`
	Messages []AccountRuleListItemListsUpdateAllListItemsResponseMessage `json:"messages"`
	Result   AccountRuleListItemListsUpdateAllListItemsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListItemListsUpdateAllListItemsResponseSuccess `json:"success"`
	JSON    accountRuleListItemListsUpdateAllListItemsResponseJSON    `json:"-"`
}

// accountRuleListItemListsUpdateAllListItemsResponseJSON contains the JSON
// metadata for the struct [AccountRuleListItemListsUpdateAllListItemsResponse]
type accountRuleListItemListsUpdateAllListItemsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsUpdateAllListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsUpdateAllListItemsResponseError struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountRuleListItemListsUpdateAllListItemsResponseErrorJSON `json:"-"`
}

// accountRuleListItemListsUpdateAllListItemsResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemListsUpdateAllListItemsResponseError]
type accountRuleListItemListsUpdateAllListItemsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsUpdateAllListItemsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsUpdateAllListItemsResponseMessage struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountRuleListItemListsUpdateAllListItemsResponseMessageJSON `json:"-"`
}

// accountRuleListItemListsUpdateAllListItemsResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemListsUpdateAllListItemsResponseMessage]
type accountRuleListItemListsUpdateAllListItemsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsUpdateAllListItemsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsUpdateAllListItemsResponseResult struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                                       `json:"operation_id"`
	JSON        accountRuleListItemListsUpdateAllListItemsResponseResultJSON `json:"-"`
}

// accountRuleListItemListsUpdateAllListItemsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemListsUpdateAllListItemsResponseResult]
type accountRuleListItemListsUpdateAllListItemsResponseResultJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsUpdateAllListItemsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRuleListItemListsUpdateAllListItemsResponseSuccess bool

const (
	AccountRuleListItemListsUpdateAllListItemsResponseSuccessTrue AccountRuleListItemListsUpdateAllListItemsResponseSuccess = true
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
	// A non-negative 32 bit integer
	ASN param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[AccountRuleListItemListsNewListItemsParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[AccountRuleListItemListsNewListItemsParamsBodyRedirect] `json:"redirect"`
}

func (r AccountRuleListItemListsNewListItemsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsNewListItemsParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r AccountRuleListItemListsNewListItemsParamsBodyHostname) MarshalJSON() (data []byte, err error) {
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
	// Amount of results to include in each paginated response. A non-negative 32 bit
	// integer.
	PerPage param.Field[int64] `query:"per_page"`
	// A search query to filter returned items. Its meaning depends on the list type:
	// IP addresses must start with the provided string, hostnames and bulk redirects
	// must contain the string, and ASNs must match the string exactly.
	Search param.Field[string] `query:"search"`
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
	// A non-negative 32 bit integer
	ASN param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[AccountRuleListItemListsUpdateAllListItemsParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[AccountRuleListItemListsUpdateAllListItemsParamsBodyRedirect] `json:"redirect"`
}

func (r AccountRuleListItemListsUpdateAllListItemsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsUpdateAllListItemsParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r AccountRuleListItemListsUpdateAllListItemsParamsBodyHostname) MarshalJSON() (data []byte, err error) {
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
