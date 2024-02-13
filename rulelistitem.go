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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// RuleListItemService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRuleListItemService] method
// instead.
type RuleListItemService struct {
	Options []option.RequestOption
}

// NewRuleListItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuleListItemService(opts ...option.RequestOption) (r *RuleListItemService) {
	r = &RuleListItemService{}
	r.Options = opts
	return
}

// Removes one or more items from a list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *RuleListItemService) Delete(ctx context.Context, accountID string, listID string, body RuleListItemDeleteParams, opts ...option.RequestOption) (res *RuleListItemDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list item in the list.
func (r *RuleListItemService) Get(ctx context.Context, accountIdentifier string, listID string, itemID string, opts ...option.RequestOption) (res *RuleListItemGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items/%s", accountIdentifier, listID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Appends new items to the list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *RuleListItemService) ListsNewListItems(ctx context.Context, accountID string, listID string, body RuleListItemListsNewListItemsParams, opts ...option.RequestOption) (res *RuleListItemListsNewListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemListsNewListItemsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the items in the list.
func (r *RuleListItemService) ListsGetListItems(ctx context.Context, accountID string, listID string, query RuleListItemListsGetListItemsParams, opts ...option.RequestOption) (res *[]RuleListItemListsGetListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemListsGetListItemsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
func (r *RuleListItemService) ListsUpdateAllListItems(ctx context.Context, accountID string, listID string, body RuleListItemListsUpdateAllListItemsParams, opts ...option.RequestOption) (res *RuleListItemListsUpdateAllListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemListsUpdateAllListItemsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleListItemDeleteResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                         `json:"operation_id"`
	JSON        ruleListItemDeleteResponseJSON `json:"-"`
}

// ruleListItemDeleteResponseJSON contains the JSON metadata for the struct
// [RuleListItemDeleteResponse]
type ruleListItemDeleteResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
//
// Union satisfied by [shared.UnionString],
// [RuleListItemGetResponseListsItemRedirect],
// [RuleListItemGetResponseListsItemHostname] or [shared.UnionInt].
type RuleListItemGetResponse interface {
	ImplementsRuleListItemGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleListItemGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// The definition of the redirect.
type RuleListItemGetResponseListsItemRedirect struct {
	SourceURL           string                                             `json:"source_url,required"`
	TargetURL           string                                             `json:"target_url,required"`
	IncludeSubdomains   bool                                               `json:"include_subdomains"`
	PreservePathSuffix  bool                                               `json:"preserve_path_suffix"`
	PreserveQueryString bool                                               `json:"preserve_query_string"`
	StatusCode          RuleListItemGetResponseListsItemRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                                               `json:"subpath_matching"`
	JSON                ruleListItemGetResponseListsItemRedirectJSON       `json:"-"`
}

// ruleListItemGetResponseListsItemRedirectJSON contains the JSON metadata for the
// struct [RuleListItemGetResponseListsItemRedirect]
type ruleListItemGetResponseListsItemRedirectJSON struct {
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

func (r *RuleListItemGetResponseListsItemRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleListItemGetResponseListsItemRedirect) ImplementsRuleListItemGetResponse() {}

type RuleListItemGetResponseListsItemRedirectStatusCode int64

const (
	RuleListItemGetResponseListsItemRedirectStatusCode301 RuleListItemGetResponseListsItemRedirectStatusCode = 301
	RuleListItemGetResponseListsItemRedirectStatusCode302 RuleListItemGetResponseListsItemRedirectStatusCode = 302
	RuleListItemGetResponseListsItemRedirectStatusCode307 RuleListItemGetResponseListsItemRedirectStatusCode = 307
	RuleListItemGetResponseListsItemRedirectStatusCode308 RuleListItemGetResponseListsItemRedirectStatusCode = 308
)

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type RuleListItemGetResponseListsItemHostname struct {
	URLHostname string                                       `json:"url_hostname,required"`
	JSON        ruleListItemGetResponseListsItemHostnameJSON `json:"-"`
}

// ruleListItemGetResponseListsItemHostnameJSON contains the JSON metadata for the
// struct [RuleListItemGetResponseListsItemHostname]
type ruleListItemGetResponseListsItemHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemGetResponseListsItemHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleListItemGetResponseListsItemHostname) ImplementsRuleListItemGetResponse() {}

type RuleListItemListsNewListItemsResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                    `json:"operation_id"`
	JSON        ruleListItemListsNewListItemsResponseJSON `json:"-"`
}

// ruleListItemListsNewListItemsResponseJSON contains the JSON metadata for the
// struct [RuleListItemListsNewListItemsResponse]
type ruleListItemListsNewListItemsResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsNewListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsGetListItemsResponse = interface{}

type RuleListItemListsUpdateAllListItemsResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                          `json:"operation_id"`
	JSON        ruleListItemListsUpdateAllListItemsResponseJSON `json:"-"`
}

// ruleListItemListsUpdateAllListItemsResponseJSON contains the JSON metadata for
// the struct [RuleListItemListsUpdateAllListItemsResponse]
type ruleListItemListsUpdateAllListItemsResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsUpdateAllListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemDeleteParams struct {
	Items param.Field[[]RuleListItemDeleteParamsItem] `json:"items"`
}

func (r RuleListItemDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListItemDeleteParamsItem struct {
	// The unique ID of the item in the List.
	ID param.Field[string] `json:"id"`
}

func (r RuleListItemDeleteParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListItemDeleteResponseEnvelope struct {
	Errors   []RuleListItemDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleListItemDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListItemDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListItemDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleListItemDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListItemDeleteResponseEnvelope]
type ruleListItemDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    ruleListItemDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListItemDeleteResponseEnvelopeErrors]
type ruleListItemDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    ruleListItemDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RuleListItemDeleteResponseEnvelopeMessages]
type ruleListItemDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListItemDeleteResponseEnvelopeSuccess bool

const (
	RuleListItemDeleteResponseEnvelopeSuccessTrue RuleListItemDeleteResponseEnvelopeSuccess = true
)

type RuleListItemGetResponseEnvelope struct {
	Errors   []RuleListItemGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemGetResponseEnvelopeMessages `json:"messages,required"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	Result RuleListItemGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListItemGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListItemGetResponseEnvelopeJSON    `json:"-"`
}

// ruleListItemGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListItemGetResponseEnvelope]
type ruleListItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    ruleListItemGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListItemGetResponseEnvelopeErrors]
type ruleListItemGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    ruleListItemGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListItemGetResponseEnvelopeMessages]
type ruleListItemGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListItemGetResponseEnvelopeSuccess bool

const (
	RuleListItemGetResponseEnvelopeSuccessTrue RuleListItemGetResponseEnvelopeSuccess = true
)

type RuleListItemListsNewListItemsParams struct {
	Body param.Field[[]RuleListItemListsNewListItemsParamsBody] `json:"body,required"`
}

func (r RuleListItemListsNewListItemsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleListItemListsNewListItemsParamsBody struct {
	// A non-negative 32 bit integer
	Asn param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[RuleListItemListsNewListItemsParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RuleListItemListsNewListItemsParamsBodyRedirect] `json:"redirect"`
}

func (r RuleListItemListsNewListItemsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type RuleListItemListsNewListItemsParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r RuleListItemListsNewListItemsParamsBodyHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type RuleListItemListsNewListItemsParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                                    `json:"source_url,required"`
	TargetURL           param.Field[string]                                                    `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                                      `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                                      `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                                      `json:"preserve_query_string"`
	StatusCode          param.Field[RuleListItemListsNewListItemsParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                                      `json:"subpath_matching"`
}

func (r RuleListItemListsNewListItemsParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListItemListsNewListItemsParamsBodyRedirectStatusCode int64

const (
	RuleListItemListsNewListItemsParamsBodyRedirectStatusCode301 RuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 301
	RuleListItemListsNewListItemsParamsBodyRedirectStatusCode302 RuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 302
	RuleListItemListsNewListItemsParamsBodyRedirectStatusCode307 RuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 307
	RuleListItemListsNewListItemsParamsBodyRedirectStatusCode308 RuleListItemListsNewListItemsParamsBodyRedirectStatusCode = 308
)

type RuleListItemListsNewListItemsResponseEnvelope struct {
	Errors   []RuleListItemListsNewListItemsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemListsNewListItemsResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleListItemListsNewListItemsResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListItemListsNewListItemsResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListItemListsNewListItemsResponseEnvelopeJSON    `json:"-"`
}

// ruleListItemListsNewListItemsResponseEnvelopeJSON contains the JSON metadata for
// the struct [RuleListItemListsNewListItemsResponseEnvelope]
type ruleListItemListsNewListItemsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsNewListItemsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsNewListItemsResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    ruleListItemListsNewListItemsResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemListsNewListItemsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [RuleListItemListsNewListItemsResponseEnvelopeErrors]
type ruleListItemListsNewListItemsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsNewListItemsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsNewListItemsResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    ruleListItemListsNewListItemsResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemListsNewListItemsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [RuleListItemListsNewListItemsResponseEnvelopeMessages]
type ruleListItemListsNewListItemsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsNewListItemsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListItemListsNewListItemsResponseEnvelopeSuccess bool

const (
	RuleListItemListsNewListItemsResponseEnvelopeSuccessTrue RuleListItemListsNewListItemsResponseEnvelopeSuccess = true
)

type RuleListItemListsGetListItemsParams struct {
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

// URLQuery serializes [RuleListItemListsGetListItemsParams]'s query parameters as
// `url.Values`.
func (r RuleListItemListsGetListItemsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleListItemListsGetListItemsResponseEnvelope struct {
	Errors   []RuleListItemListsGetListItemsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemListsGetListItemsResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListItemListsGetListItemsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleListItemListsGetListItemsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleListItemListsGetListItemsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleListItemListsGetListItemsResponseEnvelopeJSON       `json:"-"`
}

// ruleListItemListsGetListItemsResponseEnvelopeJSON contains the JSON metadata for
// the struct [RuleListItemListsGetListItemsResponseEnvelope]
type ruleListItemListsGetListItemsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsGetListItemsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsGetListItemsResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    ruleListItemListsGetListItemsResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemListsGetListItemsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [RuleListItemListsGetListItemsResponseEnvelopeErrors]
type ruleListItemListsGetListItemsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsGetListItemsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsGetListItemsResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    ruleListItemListsGetListItemsResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemListsGetListItemsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [RuleListItemListsGetListItemsResponseEnvelopeMessages]
type ruleListItemListsGetListItemsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsGetListItemsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListItemListsGetListItemsResponseEnvelopeSuccess bool

const (
	RuleListItemListsGetListItemsResponseEnvelopeSuccessTrue RuleListItemListsGetListItemsResponseEnvelopeSuccess = true
)

type RuleListItemListsGetListItemsResponseEnvelopeResultInfo struct {
	Cursors RuleListItemListsGetListItemsResponseEnvelopeResultInfoCursors `json:"cursors"`
	JSON    ruleListItemListsGetListItemsResponseEnvelopeResultInfoJSON    `json:"-"`
}

// ruleListItemListsGetListItemsResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [RuleListItemListsGetListItemsResponseEnvelopeResultInfo]
type ruleListItemListsGetListItemsResponseEnvelopeResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsGetListItemsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsGetListItemsResponseEnvelopeResultInfoCursors struct {
	After  string                                                             `json:"after"`
	Before string                                                             `json:"before"`
	JSON   ruleListItemListsGetListItemsResponseEnvelopeResultInfoCursorsJSON `json:"-"`
}

// ruleListItemListsGetListItemsResponseEnvelopeResultInfoCursorsJSON contains the
// JSON metadata for the struct
// [RuleListItemListsGetListItemsResponseEnvelopeResultInfoCursors]
type ruleListItemListsGetListItemsResponseEnvelopeResultInfoCursorsJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsGetListItemsResponseEnvelopeResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsUpdateAllListItemsParams struct {
	Body param.Field[[]RuleListItemListsUpdateAllListItemsParamsBody] `json:"body,required"`
}

func (r RuleListItemListsUpdateAllListItemsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleListItemListsUpdateAllListItemsParamsBody struct {
	// A non-negative 32 bit integer
	Asn param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[RuleListItemListsUpdateAllListItemsParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RuleListItemListsUpdateAllListItemsParamsBodyRedirect] `json:"redirect"`
}

func (r RuleListItemListsUpdateAllListItemsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type RuleListItemListsUpdateAllListItemsParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r RuleListItemListsUpdateAllListItemsParamsBodyHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type RuleListItemListsUpdateAllListItemsParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                                          `json:"source_url,required"`
	TargetURL           param.Field[string]                                                          `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                                            `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                                            `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                                            `json:"preserve_query_string"`
	StatusCode          param.Field[RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                                            `json:"subpath_matching"`
}

func (r RuleListItemListsUpdateAllListItemsParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode int64

const (
	RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode301 RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 301
	RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode302 RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 302
	RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode307 RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 307
	RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode308 RuleListItemListsUpdateAllListItemsParamsBodyRedirectStatusCode = 308
)

type RuleListItemListsUpdateAllListItemsResponseEnvelope struct {
	Errors   []RuleListItemListsUpdateAllListItemsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemListsUpdateAllListItemsResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleListItemListsUpdateAllListItemsResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListItemListsUpdateAllListItemsResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListItemListsUpdateAllListItemsResponseEnvelopeJSON    `json:"-"`
}

// ruleListItemListsUpdateAllListItemsResponseEnvelopeJSON contains the JSON
// metadata for the struct [RuleListItemListsUpdateAllListItemsResponseEnvelope]
type ruleListItemListsUpdateAllListItemsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsUpdateAllListItemsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsUpdateAllListItemsResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    ruleListItemListsUpdateAllListItemsResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemListsUpdateAllListItemsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [RuleListItemListsUpdateAllListItemsResponseEnvelopeErrors]
type ruleListItemListsUpdateAllListItemsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsUpdateAllListItemsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListItemListsUpdateAllListItemsResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    ruleListItemListsUpdateAllListItemsResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemListsUpdateAllListItemsResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [RuleListItemListsUpdateAllListItemsResponseEnvelopeMessages]
type ruleListItemListsUpdateAllListItemsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListsUpdateAllListItemsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListItemListsUpdateAllListItemsResponseEnvelopeSuccess bool

const (
	RuleListItemListsUpdateAllListItemsResponseEnvelopeSuccessTrue RuleListItemListsUpdateAllListItemsResponseEnvelopeSuccess = true
)
