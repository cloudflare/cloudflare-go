// File generated from our OpenAPI spec by Stainless.

package rules

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ListItemService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewListItemService] method instead.
type ListItemService struct {
	Options []option.RequestOption
}

// NewListItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListItemService(opts ...option.RequestOption) (r *ListItemService) {
	r = &ListItemService{}
	r.Options = opts
	return
}

// Appends new items to the list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *ListItemService) New(ctx context.Context, listID string, params ListItemNewParams, opts ...option.RequestOption) (res *ListItemNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ListItemNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes all existing items from the list and adds the provided items to the
// list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *ListItemService) Update(ctx context.Context, listID string, params ListItemUpdateParams, opts ...option.RequestOption) (res *ListItemUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ListItemUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the items in the list.
func (r *ListItemService) List(ctx context.Context, listID string, params ListItemListParams, opts ...option.RequestOption) (res *shared.CursorPagination[ListItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches all the items in the list.
func (r *ListItemService) ListAutoPaging(ctx context.Context, listID string, params ListItemListParams, opts ...option.RequestOption) *shared.CursorPaginationAutoPager[ListItemListResponse] {
	return shared.NewCursorPaginationAutoPager(r.List(ctx, listID, params, opts...))
}

// Removes one or more items from a list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *ListItemService) Delete(ctx context.Context, listID string, params ListItemDeleteParams, opts ...option.RequestOption) (res *ListItemDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ListItemDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list item in the list.
func (r *ListItemService) Get(ctx context.Context, accountIdentifier string, listID string, itemID string, opts ...option.RequestOption) (res *ListItemGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ListItemGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items/%s", accountIdentifier, listID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ListItemNewResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                  `json:"operation_id"`
	JSON        listItemNewResponseJSON `json:"-"`
}

// listItemNewResponseJSON contains the JSON metadata for the struct
// [ListItemNewResponse]
type listItemNewResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemUpdateResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                     `json:"operation_id"`
	JSON        listItemUpdateResponseJSON `json:"-"`
}

// listItemUpdateResponseJSON contains the JSON metadata for the struct
// [ListItemUpdateResponse]
type listItemUpdateResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemListResponse = interface{}

type ListItemDeleteResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                     `json:"operation_id"`
	JSON        listItemDeleteResponseJSON `json:"-"`
}

// listItemDeleteResponseJSON contains the JSON metadata for the struct
// [ListItemDeleteResponse]
type listItemDeleteResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
//
// Union satisfied by [shared.UnionString],
// [rules.ListItemGetResponseListsItemRedirect],
// [rules.ListItemGetResponseListsItemHostname] or [shared.UnionInt].
type ListItemGetResponse interface {
	ImplementsRulesListItemGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListItemGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemGetResponseListsItemRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemGetResponseListsItemHostname{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// The definition of the redirect.
type ListItemGetResponseListsItemRedirect struct {
	SourceURL           string                                         `json:"source_url,required"`
	TargetURL           string                                         `json:"target_url,required"`
	IncludeSubdomains   bool                                           `json:"include_subdomains"`
	PreservePathSuffix  bool                                           `json:"preserve_path_suffix"`
	PreserveQueryString bool                                           `json:"preserve_query_string"`
	StatusCode          ListItemGetResponseListsItemRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                           `json:"subpath_matching"`
	JSON                listItemGetResponseListsItemRedirectJSON       `json:"-"`
}

// listItemGetResponseListsItemRedirectJSON contains the JSON metadata for the
// struct [ListItemGetResponseListsItemRedirect]
type listItemGetResponseListsItemRedirectJSON struct {
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

func (r *ListItemGetResponseListsItemRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseListsItemRedirectJSON) RawJSON() string {
	return r.raw
}

func (r ListItemGetResponseListsItemRedirect) ImplementsRulesListItemGetResponse() {}

type ListItemGetResponseListsItemRedirectStatusCode int64

const (
	ListItemGetResponseListsItemRedirectStatusCode301 ListItemGetResponseListsItemRedirectStatusCode = 301
	ListItemGetResponseListsItemRedirectStatusCode302 ListItemGetResponseListsItemRedirectStatusCode = 302
	ListItemGetResponseListsItemRedirectStatusCode307 ListItemGetResponseListsItemRedirectStatusCode = 307
	ListItemGetResponseListsItemRedirectStatusCode308 ListItemGetResponseListsItemRedirectStatusCode = 308
)

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemGetResponseListsItemHostname struct {
	URLHostname string                                   `json:"url_hostname,required"`
	JSON        listItemGetResponseListsItemHostnameJSON `json:"-"`
}

// listItemGetResponseListsItemHostnameJSON contains the JSON metadata for the
// struct [ListItemGetResponseListsItemHostname]
type listItemGetResponseListsItemHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemGetResponseListsItemHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseListsItemHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ListItemGetResponseListsItemHostname) ImplementsRulesListItemGetResponse() {}

type ListItemNewParams struct {
	// Identifier
	AccountID param.Field[string]                  `path:"account_id,required"`
	Body      param.Field[[]ListItemNewParamsBody] `json:"body,required"`
}

func (r ListItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListItemNewParamsBody struct {
	// A non-negative 32 bit integer
	ASN param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[ListItemNewParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[ListItemNewParamsBodyRedirect] `json:"redirect"`
}

func (r ListItemNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemNewParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r ListItemNewParamsBodyHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type ListItemNewParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                  `json:"source_url,required"`
	TargetURL           param.Field[string]                                  `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                    `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                    `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                    `json:"preserve_query_string"`
	StatusCode          param.Field[ListItemNewParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                    `json:"subpath_matching"`
}

func (r ListItemNewParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemNewParamsBodyRedirectStatusCode int64

const (
	ListItemNewParamsBodyRedirectStatusCode301 ListItemNewParamsBodyRedirectStatusCode = 301
	ListItemNewParamsBodyRedirectStatusCode302 ListItemNewParamsBodyRedirectStatusCode = 302
	ListItemNewParamsBodyRedirectStatusCode307 ListItemNewParamsBodyRedirectStatusCode = 307
	ListItemNewParamsBodyRedirectStatusCode308 ListItemNewParamsBodyRedirectStatusCode = 308
)

type ListItemNewResponseEnvelope struct {
	Errors   []ListItemNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListItemNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ListItemNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListItemNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemNewResponseEnvelopeJSON    `json:"-"`
}

// listItemNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemNewResponseEnvelope]
type listItemNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListItemNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    listItemNewResponseEnvelopeErrorsJSON `json:"-"`
}

// listItemNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ListItemNewResponseEnvelopeErrors]
type listItemNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListItemNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    listItemNewResponseEnvelopeMessagesJSON `json:"-"`
}

// listItemNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ListItemNewResponseEnvelopeMessages]
type listItemNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListItemNewResponseEnvelopeSuccess bool

const (
	ListItemNewResponseEnvelopeSuccessTrue ListItemNewResponseEnvelopeSuccess = true
)

type ListItemUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                     `path:"account_id,required"`
	Body      param.Field[[]ListItemUpdateParamsBody] `json:"body,required"`
}

func (r ListItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListItemUpdateParamsBody struct {
	// A non-negative 32 bit integer
	ASN param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[ListItemUpdateParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[ListItemUpdateParamsBodyRedirect] `json:"redirect"`
}

func (r ListItemUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemUpdateParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r ListItemUpdateParamsBodyHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type ListItemUpdateParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                     `json:"source_url,required"`
	TargetURL           param.Field[string]                                     `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                       `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                       `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                       `json:"preserve_query_string"`
	StatusCode          param.Field[ListItemUpdateParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                       `json:"subpath_matching"`
}

func (r ListItemUpdateParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemUpdateParamsBodyRedirectStatusCode int64

const (
	ListItemUpdateParamsBodyRedirectStatusCode301 ListItemUpdateParamsBodyRedirectStatusCode = 301
	ListItemUpdateParamsBodyRedirectStatusCode302 ListItemUpdateParamsBodyRedirectStatusCode = 302
	ListItemUpdateParamsBodyRedirectStatusCode307 ListItemUpdateParamsBodyRedirectStatusCode = 307
	ListItemUpdateParamsBodyRedirectStatusCode308 ListItemUpdateParamsBodyRedirectStatusCode = 308
)

type ListItemUpdateResponseEnvelope struct {
	Errors   []ListItemUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListItemUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ListItemUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListItemUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemUpdateResponseEnvelopeJSON    `json:"-"`
}

// listItemUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemUpdateResponseEnvelope]
type listItemUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListItemUpdateResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    listItemUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// listItemUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ListItemUpdateResponseEnvelopeErrors]
type listItemUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListItemUpdateResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    listItemUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// listItemUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ListItemUpdateResponseEnvelopeMessages]
type listItemUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListItemUpdateResponseEnvelopeSuccess bool

const (
	ListItemUpdateResponseEnvelopeSuccessTrue ListItemUpdateResponseEnvelopeSuccess = true
)

type ListItemListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

// URLQuery serializes [ListItemListParams]'s query parameters as `url.Values`.
func (r ListItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListItemDeleteParams struct {
	// Identifier
	AccountID param.Field[string]                     `path:"account_id,required"`
	Items     param.Field[[]ListItemDeleteParamsItem] `json:"items"`
}

func (r ListItemDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemDeleteParamsItem struct {
	// The unique ID of the item in the List.
	ID param.Field[string] `json:"id"`
}

func (r ListItemDeleteParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemDeleteResponseEnvelope struct {
	Errors   []ListItemDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListItemDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ListItemDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListItemDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemDeleteResponseEnvelopeJSON    `json:"-"`
}

// listItemDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemDeleteResponseEnvelope]
type listItemDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListItemDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    listItemDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// listItemDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ListItemDeleteResponseEnvelopeErrors]
type listItemDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListItemDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    listItemDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// listItemDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ListItemDeleteResponseEnvelopeMessages]
type listItemDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListItemDeleteResponseEnvelopeSuccess bool

const (
	ListItemDeleteResponseEnvelopeSuccessTrue ListItemDeleteResponseEnvelopeSuccess = true
)

type ListItemGetResponseEnvelope struct {
	Errors   []ListItemGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListItemGetResponseEnvelopeMessages `json:"messages,required"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	Result ListItemGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListItemGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemGetResponseEnvelopeJSON    `json:"-"`
}

// listItemGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemGetResponseEnvelope]
type listItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListItemGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    listItemGetResponseEnvelopeErrorsJSON `json:"-"`
}

// listItemGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ListItemGetResponseEnvelopeErrors]
type listItemGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListItemGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    listItemGetResponseEnvelopeMessagesJSON `json:"-"`
}

// listItemGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ListItemGetResponseEnvelopeMessages]
type listItemGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListItemGetResponseEnvelopeSuccess bool

const (
	ListItemGetResponseEnvelopeSuccessTrue ListItemGetResponseEnvelopeSuccess = true
)
