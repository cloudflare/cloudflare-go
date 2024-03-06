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

// Appends new items to the list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *RuleListItemService) New(ctx context.Context, listID string, params RuleListItemNewParams, opts ...option.RequestOption) (res *RuleListItemNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemNewResponseEnvelope
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
func (r *RuleListItemService) Update(ctx context.Context, listID string, params RuleListItemUpdateParams, opts ...option.RequestOption) (res *RuleListItemUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the items in the list.
func (r *RuleListItemService) List(ctx context.Context, listID string, params RuleListItemListParams, opts ...option.RequestOption) (res *[]RuleListItemListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes one or more items from a list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
func (r *RuleListItemService) Delete(ctx context.Context, listID string, params RuleListItemDeleteParams, opts ...option.RequestOption) (res *RuleListItemDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListItemDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
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

type RuleListItemNewResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                      `json:"operation_id"`
	JSON        ruleListItemNewResponseJSON `json:"-"`
}

// ruleListItemNewResponseJSON contains the JSON metadata for the struct
// [RuleListItemNewResponse]
type ruleListItemNewResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemNewResponseJSON) RawJSON() string {
	return r.raw
}

type RuleListItemUpdateResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                         `json:"operation_id"`
	JSON        ruleListItemUpdateResponseJSON `json:"-"`
}

// ruleListItemUpdateResponseJSON contains the JSON metadata for the struct
// [RuleListItemUpdateResponse]
type ruleListItemUpdateResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RuleListItemListResponse = interface{}

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

func (r ruleListItemDeleteResponseJSON) RawJSON() string {
	return r.raw
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
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleListItemGetResponseListsItemRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleListItemGetResponseListsItemHostname{}),
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

func (r ruleListItemGetResponseListsItemRedirectJSON) RawJSON() string {
	return r.raw
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

func (r ruleListItemGetResponseListsItemHostnameJSON) RawJSON() string {
	return r.raw
}

func (r RuleListItemGetResponseListsItemHostname) ImplementsRuleListItemGetResponse() {}

type RuleListItemNewParams struct {
	// Identifier
	AccountID param.Field[string]                      `path:"account_id,required"`
	Body      param.Field[[]RuleListItemNewParamsBody] `json:"body,required"`
}

func (r RuleListItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleListItemNewParamsBody struct {
	// A non-negative 32 bit integer
	ASN param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[RuleListItemNewParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RuleListItemNewParamsBodyRedirect] `json:"redirect"`
}

func (r RuleListItemNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type RuleListItemNewParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r RuleListItemNewParamsBodyHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type RuleListItemNewParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                      `json:"source_url,required"`
	TargetURL           param.Field[string]                                      `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                        `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                        `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                        `json:"preserve_query_string"`
	StatusCode          param.Field[RuleListItemNewParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                        `json:"subpath_matching"`
}

func (r RuleListItemNewParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListItemNewParamsBodyRedirectStatusCode int64

const (
	RuleListItemNewParamsBodyRedirectStatusCode301 RuleListItemNewParamsBodyRedirectStatusCode = 301
	RuleListItemNewParamsBodyRedirectStatusCode302 RuleListItemNewParamsBodyRedirectStatusCode = 302
	RuleListItemNewParamsBodyRedirectStatusCode307 RuleListItemNewParamsBodyRedirectStatusCode = 307
	RuleListItemNewParamsBodyRedirectStatusCode308 RuleListItemNewParamsBodyRedirectStatusCode = 308
)

type RuleListItemNewResponseEnvelope struct {
	Errors   []RuleListItemNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleListItemNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListItemNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListItemNewResponseEnvelopeJSON    `json:"-"`
}

// ruleListItemNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListItemNewResponseEnvelope]
type ruleListItemNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleListItemNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    ruleListItemNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListItemNewResponseEnvelopeErrors]
type ruleListItemNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleListItemNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    ruleListItemNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListItemNewResponseEnvelopeMessages]
type ruleListItemNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleListItemNewResponseEnvelopeSuccess bool

const (
	RuleListItemNewResponseEnvelopeSuccessTrue RuleListItemNewResponseEnvelopeSuccess = true
)

type RuleListItemUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                         `path:"account_id,required"`
	Body      param.Field[[]RuleListItemUpdateParamsBody] `json:"body,required"`
}

func (r RuleListItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleListItemUpdateParamsBody struct {
	// A non-negative 32 bit integer
	ASN param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[RuleListItemUpdateParamsBodyHostname] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RuleListItemUpdateParamsBodyRedirect] `json:"redirect"`
}

func (r RuleListItemUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type RuleListItemUpdateParamsBodyHostname struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r RuleListItemUpdateParamsBodyHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type RuleListItemUpdateParamsBodyRedirect struct {
	SourceURL           param.Field[string]                                         `json:"source_url,required"`
	TargetURL           param.Field[string]                                         `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                                           `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                                           `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                                           `json:"preserve_query_string"`
	StatusCode          param.Field[RuleListItemUpdateParamsBodyRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                                           `json:"subpath_matching"`
}

func (r RuleListItemUpdateParamsBodyRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListItemUpdateParamsBodyRedirectStatusCode int64

const (
	RuleListItemUpdateParamsBodyRedirectStatusCode301 RuleListItemUpdateParamsBodyRedirectStatusCode = 301
	RuleListItemUpdateParamsBodyRedirectStatusCode302 RuleListItemUpdateParamsBodyRedirectStatusCode = 302
	RuleListItemUpdateParamsBodyRedirectStatusCode307 RuleListItemUpdateParamsBodyRedirectStatusCode = 307
	RuleListItemUpdateParamsBodyRedirectStatusCode308 RuleListItemUpdateParamsBodyRedirectStatusCode = 308
)

type RuleListItemUpdateResponseEnvelope struct {
	Errors   []RuleListItemUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleListItemUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListItemUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListItemUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleListItemUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListItemUpdateResponseEnvelope]
type ruleListItemUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleListItemUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    ruleListItemUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListItemUpdateResponseEnvelopeErrors]
type ruleListItemUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleListItemUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    ruleListItemUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RuleListItemUpdateResponseEnvelopeMessages]
type ruleListItemUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleListItemUpdateResponseEnvelopeSuccess bool

const (
	RuleListItemUpdateResponseEnvelopeSuccessTrue RuleListItemUpdateResponseEnvelopeSuccess = true
)

type RuleListItemListParams struct {
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

// URLQuery serializes [RuleListItemListParams]'s query parameters as `url.Values`.
func (r RuleListItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleListItemListResponseEnvelope struct {
	Errors   []RuleListItemListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListItemListResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListItemListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleListItemListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleListItemListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleListItemListResponseEnvelopeJSON       `json:"-"`
}

// ruleListItemListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListItemListResponseEnvelope]
type ruleListItemListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleListItemListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    ruleListItemListResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListItemListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListItemListResponseEnvelopeErrors]
type ruleListItemListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleListItemListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    ruleListItemListResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListItemListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListItemListResponseEnvelopeMessages]
type ruleListItemListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleListItemListResponseEnvelopeSuccess bool

const (
	RuleListItemListResponseEnvelopeSuccessTrue RuleListItemListResponseEnvelopeSuccess = true
)

type RuleListItemListResponseEnvelopeResultInfo struct {
	Cursors RuleListItemListResponseEnvelopeResultInfoCursors `json:"cursors"`
	JSON    ruleListItemListResponseEnvelopeResultInfoJSON    `json:"-"`
}

// ruleListItemListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [RuleListItemListResponseEnvelopeResultInfo]
type ruleListItemListResponseEnvelopeResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleListItemListResponseEnvelopeResultInfoCursors struct {
	After  string                                                `json:"after"`
	Before string                                                `json:"before"`
	JSON   ruleListItemListResponseEnvelopeResultInfoCursorsJSON `json:"-"`
}

// ruleListItemListResponseEnvelopeResultInfoCursorsJSON contains the JSON metadata
// for the struct [RuleListItemListResponseEnvelopeResultInfoCursors]
type ruleListItemListResponseEnvelopeResultInfoCursorsJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListItemListResponseEnvelopeResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListItemListResponseEnvelopeResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type RuleListItemDeleteParams struct {
	// Identifier
	AccountID param.Field[string]                         `path:"account_id,required"`
	Items     param.Field[[]RuleListItemDeleteParamsItem] `json:"items"`
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

func (r ruleListItemDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r ruleListItemDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r ruleListItemDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r ruleListItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r ruleListItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r ruleListItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleListItemGetResponseEnvelopeSuccess bool

const (
	RuleListItemGetResponseEnvelopeSuccessTrue RuleListItemGetResponseEnvelopeSuccess = true
)
