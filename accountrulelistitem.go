// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

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
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	Result AccountRuleListItemGetResponseResult `json:"result"`
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

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
//
// Union satisfied by [AccountRuleListItemGetResponseResultPeciSksBItemIP],
// [AccountRuleListItemGetResponseResultPeciSksBItemRedirect],
// [AccountRuleListItemGetResponseResultPeciSksBItemHostname] or
// [AccountRuleListItemGetResponseResultPeciSksBItemASN].
type AccountRuleListItemGetResponseResult interface {
	implementsAccountRuleListItemGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountRuleListItemGetResponseResult)(nil)).Elem(), "")
}

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
type AccountRuleListItemGetResponseResultPeciSksBItemIP struct {
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
	Hostname AccountRuleListItemGetResponseResultPeciSksBItemIPHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemGetResponseResultPeciSksBItemIPRedirect `json:"redirect"`
	JSON     accountRuleListItemGetResponseResultPeciSksBItemIPJSON     `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemIPJSON contains the JSON
// metadata for the struct [AccountRuleListItemGetResponseResultPeciSksBItemIP]
type accountRuleListItemGetResponseResultPeciSksBItemIPJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemGetResponseResultPeciSksBItemIP) implementsAccountRuleListItemGetResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemGetResponseResultPeciSksBItemIPHostname struct {
	URLHostname string                                                         `json:"url_hostname,required"`
	JSON        accountRuleListItemGetResponseResultPeciSksBItemIPHostnameJSON `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemIPHostnameJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemIPHostname]
type accountRuleListItemGetResponseResultPeciSksBItemIPHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseResultPeciSksBItemIPHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemGetResponseResultPeciSksBItemIPRedirect struct {
	SourceURL           string                                                               `json:"source_url,required"`
	TargetURL           string                                                               `json:"target_url,required"`
	IncludeSubdomains   bool                                                                 `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                 `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                 `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                 `json:"subpath_matching"`
	JSON                accountRuleListItemGetResponseResultPeciSksBItemIPRedirectJSON       `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemIPRedirectJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemIPRedirect]
type accountRuleListItemGetResponseResultPeciSksBItemIPRedirectJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemIPRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode int64

const (
	AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode301 AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode = 301
	AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode302 AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode = 302
	AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode307 AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode = 307
	AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode308 AccountRuleListItemGetResponseResultPeciSksBItemIPRedirectStatusCode = 308
)

// The definition of the redirect.
type AccountRuleListItemGetResponseResultPeciSksBItemRedirect struct {
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
	Hostname AccountRuleListItemGetResponseResultPeciSksBItemRedirectHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirect `json:"redirect"`
	JSON     accountRuleListItemGetResponseResultPeciSksBItemRedirectJSON     `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemRedirectJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemRedirect]
type accountRuleListItemGetResponseResultPeciSksBItemRedirectJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemGetResponseResultPeciSksBItemRedirect) implementsAccountRuleListItemGetResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemGetResponseResultPeciSksBItemRedirectHostname struct {
	URLHostname string                                                               `json:"url_hostname,required"`
	JSON        accountRuleListItemGetResponseResultPeciSksBItemRedirectHostnameJSON `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemRedirectHostnameJSON contains
// the JSON metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemRedirectHostname]
type accountRuleListItemGetResponseResultPeciSksBItemRedirectHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseResultPeciSksBItemRedirectHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirect struct {
	SourceURL           string                                                                     `json:"source_url,required"`
	TargetURL           string                                                                     `json:"target_url,required"`
	IncludeSubdomains   bool                                                                       `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                       `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                       `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                       `json:"subpath_matching"`
	JSON                accountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectJSON       `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectJSON contains
// the JSON metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirect]
type accountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode int64

const (
	AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode301 AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode = 301
	AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode302 AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode = 302
	AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode307 AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode = 307
	AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode308 AccountRuleListItemGetResponseResultPeciSksBItemRedirectRedirectStatusCode = 308
)

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemGetResponseResultPeciSksBItemHostname struct {
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
	Hostname AccountRuleListItemGetResponseResultPeciSksBItemHostnameHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirect `json:"redirect"`
	JSON     accountRuleListItemGetResponseResultPeciSksBItemHostnameJSON     `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemHostnameJSON contains the JSON
// metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemHostname]
type accountRuleListItemGetResponseResultPeciSksBItemHostnameJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemGetResponseResultPeciSksBItemHostname) implementsAccountRuleListItemGetResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemGetResponseResultPeciSksBItemHostnameHostname struct {
	URLHostname string                                                               `json:"url_hostname,required"`
	JSON        accountRuleListItemGetResponseResultPeciSksBItemHostnameHostnameJSON `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemHostnameHostnameJSON contains
// the JSON metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemHostnameHostname]
type accountRuleListItemGetResponseResultPeciSksBItemHostnameHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseResultPeciSksBItemHostnameHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirect struct {
	SourceURL           string                                                                     `json:"source_url,required"`
	TargetURL           string                                                                     `json:"target_url,required"`
	IncludeSubdomains   bool                                                                       `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                       `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                       `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                       `json:"subpath_matching"`
	JSON                accountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectJSON       `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectJSON contains
// the JSON metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirect]
type accountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode int64

const (
	AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode301 AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode = 301
	AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode302 AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode = 302
	AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode307 AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode = 307
	AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode308 AccountRuleListItemGetResponseResultPeciSksBItemHostnameRedirectStatusCode = 308
)

// A non-negative 32 bit integer
type AccountRuleListItemGetResponseResultPeciSksBItemASN struct {
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
	Hostname AccountRuleListItemGetResponseResultPeciSksBItemASNHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemGetResponseResultPeciSksBItemASNRedirect `json:"redirect"`
	JSON     accountRuleListItemGetResponseResultPeciSksBItemASNJSON     `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemASNJSON contains the JSON
// metadata for the struct [AccountRuleListItemGetResponseResultPeciSksBItemASN]
type accountRuleListItemGetResponseResultPeciSksBItemASNJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemGetResponseResultPeciSksBItemASN) implementsAccountRuleListItemGetResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemGetResponseResultPeciSksBItemASNHostname struct {
	URLHostname string                                                          `json:"url_hostname,required"`
	JSON        accountRuleListItemGetResponseResultPeciSksBItemASNHostnameJSON `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemASNHostnameJSON contains the
// JSON metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemASNHostname]
type accountRuleListItemGetResponseResultPeciSksBItemASNHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponseResultPeciSksBItemASNHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemGetResponseResultPeciSksBItemASNRedirect struct {
	SourceURL           string                                                                `json:"source_url,required"`
	TargetURL           string                                                                `json:"target_url,required"`
	IncludeSubdomains   bool                                                                  `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                  `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                  `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                  `json:"subpath_matching"`
	JSON                accountRuleListItemGetResponseResultPeciSksBItemASNRedirectJSON       `json:"-"`
}

// accountRuleListItemGetResponseResultPeciSksBItemASNRedirectJSON contains the
// JSON metadata for the struct
// [AccountRuleListItemGetResponseResultPeciSksBItemASNRedirect]
type accountRuleListItemGetResponseResultPeciSksBItemASNRedirectJSON struct {
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

func (r *AccountRuleListItemGetResponseResultPeciSksBItemASNRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode int64

const (
	AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode301 AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode = 301
	AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode302 AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode = 302
	AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode307 AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode = 307
	AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode308 AccountRuleListItemGetResponseResultPeciSksBItemASNRedirectStatusCode = 308
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

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
//
// Union satisfied by
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIP],
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirect],
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostname] or
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASN].
type AccountRuleListItemListsGetListItemsResponseResult interface {
	implementsAccountRuleListItemListsGetListItemsResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountRuleListItemListsGetListItemsResponseResult)(nil)).Elem(), "")
}

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIP struct {
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
	Hostname AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirect `json:"redirect"`
	JSON     accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPJSON     `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPJSON contains
// the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIP]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIP) implementsAccountRuleListItemListsGetListItemsResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostname struct {
	URLHostname string                                                                       `json:"url_hostname,required"`
	JSON        accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostnameJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostnameJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostname]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirect struct {
	SourceURL           string                                                                             `json:"source_url,required"`
	TargetURL           string                                                                             `json:"target_url,required"`
	IncludeSubdomains   bool                                                                               `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                               `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                               `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                               `json:"subpath_matching"`
	JSON                accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectJSON       `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirect]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode int64

const (
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode301 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode = 301
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode302 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode = 302
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode307 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode = 307
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode308 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemIPRedirectStatusCode = 308
)

// The definition of the redirect.
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirect struct {
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
	Hostname AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirect `json:"redirect"`
	JSON     accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectJSON     `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirect]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirect) implementsAccountRuleListItemListsGetListItemsResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostname struct {
	URLHostname string                                                                             `json:"url_hostname,required"`
	JSON        accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostnameJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostnameJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostname]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirect struct {
	SourceURL           string                                                                                   `json:"source_url,required"`
	TargetURL           string                                                                                   `json:"target_url,required"`
	IncludeSubdomains   bool                                                                                     `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                                     `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                                     `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                                     `json:"subpath_matching"`
	JSON                accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectJSON       `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirect]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode int64

const (
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode301 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode = 301
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode302 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode = 302
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode307 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode = 307
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode308 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemRedirectRedirectStatusCode = 308
)

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostname struct {
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
	Hostname AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirect `json:"redirect"`
	JSON     accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameJSON     `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostname]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostname) implementsAccountRuleListItemListsGetListItemsResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostname struct {
	URLHostname string                                                                             `json:"url_hostname,required"`
	JSON        accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostnameJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostnameJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostname]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirect struct {
	SourceURL           string                                                                                   `json:"source_url,required"`
	TargetURL           string                                                                                   `json:"target_url,required"`
	IncludeSubdomains   bool                                                                                     `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                                     `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                                     `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                                     `json:"subpath_matching"`
	JSON                accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectJSON       `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirect]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode int64

const (
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode301 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode = 301
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode302 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode = 302
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode307 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode = 307
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode308 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemHostnameRedirectStatusCode = 308
)

// A non-negative 32 bit integer
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASN struct {
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
	Hostname AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirect `json:"redirect"`
	JSON     accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNJSON     `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNJSON contains
// the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASN]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASN) implementsAccountRuleListItemListsGetListItemsResponseResult() {
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostname struct {
	URLHostname string                                                                        `json:"url_hostname,required"`
	JSON        accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostnameJSON `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostnameJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostname]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of the redirect.
type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirect struct {
	SourceURL           string                                                                              `json:"source_url,required"`
	TargetURL           string                                                                              `json:"target_url,required"`
	IncludeSubdomains   bool                                                                                `json:"include_subdomains"`
	PreservePathSuffix  bool                                                                                `json:"preserve_path_suffix"`
	PreserveQueryString bool                                                                                `json:"preserve_query_string"`
	StatusCode          AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                                                                `json:"subpath_matching"`
	JSON                accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectJSON       `json:"-"`
}

// accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectJSON
// contains the JSON metadata for the struct
// [AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirect]
type accountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectJSON struct {
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

func (r *AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode int64

const (
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode301 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode = 301
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode302 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode = 302
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode307 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode = 307
	AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode308 AccountRuleListItemListsGetListItemsResponseResultPeciSksBItemASNRedirectStatusCode = 308
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
