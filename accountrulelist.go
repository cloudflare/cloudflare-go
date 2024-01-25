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
func (r *AccountRuleListService) Get(ctx context.Context, accountIdentifier string, listID string, opts ...option.RequestOption) (res *AccountRuleListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the description of a list.
func (r *AccountRuleListService) Update(ctx context.Context, accountIdentifier string, listID string, body AccountRuleListUpdateParams, opts ...option.RequestOption) (res *AccountRuleListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a specific list and all its items.
func (r *AccountRuleListService) Delete(ctx context.Context, accountIdentifier string, listID string, opts ...option.RequestOption) (res *AccountRuleListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountIdentifier, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new list of the specified type.
func (r *AccountRuleListService) ListsNewAList(ctx context.Context, accountIdentifier string, body AccountRuleListListsNewAListParams, opts ...option.RequestOption) (res *AccountRuleListListsNewAListResponse, err error) {
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

type AccountRuleListGetResponse struct {
	Errors   []AccountRuleListGetResponseError   `json:"errors"`
	Messages []AccountRuleListGetResponseMessage `json:"messages"`
	Result   AccountRuleListGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListGetResponseSuccess `json:"success"`
	JSON    accountRuleListGetResponseJSON    `json:"-"`
}

// accountRuleListGetResponseJSON contains the JSON metadata for the struct
// [AccountRuleListGetResponse]
type accountRuleListGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountRuleListGetResponseErrorJSON `json:"-"`
}

// accountRuleListGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountRuleListGetResponseError]
type accountRuleListGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountRuleListGetResponseMessageJSON `json:"-"`
}

// accountRuleListGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountRuleListGetResponseMessage]
type accountRuleListGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListGetResponseResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind AccountRuleListGetResponseResultKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](#filters) referencing the list.
	NumReferencingFilters float64                              `json:"num_referencing_filters"`
	JSON                  accountRuleListGetResponseResultJSON `json:"-"`
}

// accountRuleListGetResponseResultJSON contains the JSON metadata for the struct
// [AccountRuleListGetResponseResult]
type accountRuleListGetResponseResultJSON struct {
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

func (r *AccountRuleListGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type AccountRuleListGetResponseResultKind string

const (
	AccountRuleListGetResponseResultKindIP       AccountRuleListGetResponseResultKind = "ip"
	AccountRuleListGetResponseResultKindRedirect AccountRuleListGetResponseResultKind = "redirect"
	AccountRuleListGetResponseResultKindHostname AccountRuleListGetResponseResultKind = "hostname"
	AccountRuleListGetResponseResultKindASN      AccountRuleListGetResponseResultKind = "asn"
)

// Whether the API call was successful
type AccountRuleListGetResponseSuccess bool

const (
	AccountRuleListGetResponseSuccessTrue AccountRuleListGetResponseSuccess = true
)

type AccountRuleListUpdateResponse struct {
	Errors   []AccountRuleListUpdateResponseError   `json:"errors"`
	Messages []AccountRuleListUpdateResponseMessage `json:"messages"`
	Result   AccountRuleListUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListUpdateResponseSuccess `json:"success"`
	JSON    accountRuleListUpdateResponseJSON    `json:"-"`
}

// accountRuleListUpdateResponseJSON contains the JSON metadata for the struct
// [AccountRuleListUpdateResponse]
type accountRuleListUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountRuleListUpdateResponseErrorJSON `json:"-"`
}

// accountRuleListUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountRuleListUpdateResponseError]
type accountRuleListUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountRuleListUpdateResponseMessageJSON `json:"-"`
}

// accountRuleListUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountRuleListUpdateResponseMessage]
type accountRuleListUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListUpdateResponseResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind AccountRuleListUpdateResponseResultKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](#filters) referencing the list.
	NumReferencingFilters float64                                 `json:"num_referencing_filters"`
	JSON                  accountRuleListUpdateResponseResultJSON `json:"-"`
}

// accountRuleListUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountRuleListUpdateResponseResult]
type accountRuleListUpdateResponseResultJSON struct {
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

func (r *AccountRuleListUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type AccountRuleListUpdateResponseResultKind string

const (
	AccountRuleListUpdateResponseResultKindIP       AccountRuleListUpdateResponseResultKind = "ip"
	AccountRuleListUpdateResponseResultKindRedirect AccountRuleListUpdateResponseResultKind = "redirect"
	AccountRuleListUpdateResponseResultKindHostname AccountRuleListUpdateResponseResultKind = "hostname"
	AccountRuleListUpdateResponseResultKindASN      AccountRuleListUpdateResponseResultKind = "asn"
)

// Whether the API call was successful
type AccountRuleListUpdateResponseSuccess bool

const (
	AccountRuleListUpdateResponseSuccessTrue AccountRuleListUpdateResponseSuccess = true
)

type AccountRuleListDeleteResponse struct {
	Errors   []AccountRuleListDeleteResponseError   `json:"errors"`
	Messages []AccountRuleListDeleteResponseMessage `json:"messages"`
	Result   AccountRuleListDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListDeleteResponseSuccess `json:"success"`
	JSON    accountRuleListDeleteResponseJSON    `json:"-"`
}

// accountRuleListDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRuleListDeleteResponse]
type accountRuleListDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountRuleListDeleteResponseErrorJSON `json:"-"`
}

// accountRuleListDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountRuleListDeleteResponseError]
type accountRuleListDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountRuleListDeleteResponseMessageJSON `json:"-"`
}

// accountRuleListDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountRuleListDeleteResponseMessage]
type accountRuleListDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListDeleteResponseResult struct {
	// The unique ID of the item in the List.
	ID   string                                  `json:"id"`
	JSON accountRuleListDeleteResponseResultJSON `json:"-"`
}

// accountRuleListDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountRuleListDeleteResponseResult]
type accountRuleListDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRuleListDeleteResponseSuccess bool

const (
	AccountRuleListDeleteResponseSuccessTrue AccountRuleListDeleteResponseSuccess = true
)

type AccountRuleListListsNewAListResponse struct {
	Errors   []AccountRuleListListsNewAListResponseError   `json:"errors"`
	Messages []AccountRuleListListsNewAListResponseMessage `json:"messages"`
	Result   AccountRuleListListsNewAListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListListsNewAListResponseSuccess `json:"success"`
	JSON    accountRuleListListsNewAListResponseJSON    `json:"-"`
}

// accountRuleListListsNewAListResponseJSON contains the JSON metadata for the
// struct [AccountRuleListListsNewAListResponse]
type accountRuleListListsNewAListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsNewAListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListListsNewAListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountRuleListListsNewAListResponseErrorJSON `json:"-"`
}

// accountRuleListListsNewAListResponseErrorJSON contains the JSON metadata for the
// struct [AccountRuleListListsNewAListResponseError]
type accountRuleListListsNewAListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsNewAListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListListsNewAListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountRuleListListsNewAListResponseMessageJSON `json:"-"`
}

// accountRuleListListsNewAListResponseMessageJSON contains the JSON metadata for
// the struct [AccountRuleListListsNewAListResponseMessage]
type accountRuleListListsNewAListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListsNewAListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListListsNewAListResponseResult struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind AccountRuleListListsNewAListResponseResultKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](#filters) referencing the list.
	NumReferencingFilters float64                                        `json:"num_referencing_filters"`
	JSON                  accountRuleListListsNewAListResponseResultJSON `json:"-"`
}

// accountRuleListListsNewAListResponseResultJSON contains the JSON metadata for
// the struct [AccountRuleListListsNewAListResponseResult]
type accountRuleListListsNewAListResponseResultJSON struct {
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

func (r *AccountRuleListListsNewAListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type AccountRuleListListsNewAListResponseResultKind string

const (
	AccountRuleListListsNewAListResponseResultKindIP       AccountRuleListListsNewAListResponseResultKind = "ip"
	AccountRuleListListsNewAListResponseResultKindRedirect AccountRuleListListsNewAListResponseResultKind = "redirect"
	AccountRuleListListsNewAListResponseResultKindHostname AccountRuleListListsNewAListResponseResultKind = "hostname"
	AccountRuleListListsNewAListResponseResultKindASN      AccountRuleListListsNewAListResponseResultKind = "asn"
)

// Whether the API call was successful
type AccountRuleListListsNewAListResponseSuccess bool

const (
	AccountRuleListListsNewAListResponseSuccessTrue AccountRuleListListsNewAListResponseSuccess = true
)

type AccountRuleListListsGetListsResponse struct {
	Errors   []AccountRuleListListsGetListsResponseError   `json:"errors"`
	Messages []AccountRuleListListsGetListsResponseMessage `json:"messages"`
	Result   []AccountRuleListListsGetListsResponseResult  `json:"result"`
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
	// ASNs, hostnames or redirects).
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
// ASNs, hostnames or redirects).
type AccountRuleListListsGetListsResponseResultKind string

const (
	AccountRuleListListsGetListsResponseResultKindIP       AccountRuleListListsGetListsResponseResultKind = "ip"
	AccountRuleListListsGetListsResponseResultKindRedirect AccountRuleListListsGetListsResponseResultKind = "redirect"
	AccountRuleListListsGetListsResponseResultKindHostname AccountRuleListListsGetListsResponseResultKind = "hostname"
	AccountRuleListListsGetListsResponseResultKindASN      AccountRuleListListsGetListsResponseResultKind = "asn"
)

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
	// ASNs, hostnames or redirects).
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
// ASNs, hostnames or redirects).
type AccountRuleListListsNewAListParamsKind string

const (
	AccountRuleListListsNewAListParamsKindIP       AccountRuleListListsNewAListParamsKind = "ip"
	AccountRuleListListsNewAListParamsKindRedirect AccountRuleListListsNewAListParamsKind = "redirect"
	AccountRuleListListsNewAListParamsKindHostname AccountRuleListListsNewAListParamsKind = "hostname"
	AccountRuleListListsNewAListParamsKindASN      AccountRuleListListsNewAListParamsKind = "asn"
)
