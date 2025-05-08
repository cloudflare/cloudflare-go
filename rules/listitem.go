// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// ListItemService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListItemService] method instead.
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
	var env ListItemNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
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
	var env ListItemUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the items in the list.
func (r *ListItemService) List(ctx context.Context, listID string, params ListItemListParams, opts ...option.RequestOption) (res *[]interface{}, err error) {
	var env ListItemListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
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
func (r *ListItemService) Delete(ctx context.Context, listID string, body ListItemDeleteParams, opts ...option.RequestOption) (res *ListItemDeleteResponse, err error) {
	var env ListItemDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", body.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list item in the list.
func (r *ListItemService) Get(ctx context.Context, listID string, itemID string, query ListItemGetParams, opts ...option.RequestOption) (res *ListItemGetResponse, err error) {
	var env ListItemGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items/%s", query.AccountID, listID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ListCursor struct {
	After  string         `json:"after"`
	Before string         `json:"before"`
	JSON   listCursorJSON `json:"-"`
}

// listCursorJSON contains the JSON metadata for the struct [ListCursor]
type listCursorJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListCursor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listCursorJSON) RawJSON() string {
	return r.raw
}

type ListItemNewResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                  `json:"operation_id"`
	JSON        listItemNewResponseJSON `json:"-"`
	union       ListItemNewResponseUnion
}

// listItemNewResponseJSON contains the JSON metadata for the struct
// [ListItemNewResponse]
type listItemNewResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r listItemNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListItemNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListItemNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ListItemNewResponseOperationID],
// [ListItemNewResponseOperationID].
func (r ListItemNewResponse) AsUnion() ListItemNewResponseUnion {
	return r.union
}

// Union satisfied by [ListItemNewResponseOperationID] or
// [ListItemNewResponseOperationID].
type ListItemNewResponseUnion interface {
	implementsListItemNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListItemNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemNewResponseOperationID{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemNewResponseOperationID{}),
		},
	)
}

type ListItemNewResponseOperationID struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                             `json:"operation_id"`
	JSON        listItemNewResponseOperationIDJSON `json:"-"`
}

// listItemNewResponseOperationIDJSON contains the JSON metadata for the struct
// [ListItemNewResponseOperationID]
type listItemNewResponseOperationIDJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponseOperationID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseOperationIDJSON) RawJSON() string {
	return r.raw
}

func (r ListItemNewResponseOperationID) implementsListItemNewResponse() {}

type ListItemUpdateResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                     `json:"operation_id"`
	JSON        listItemUpdateResponseJSON `json:"-"`
	union       ListItemUpdateResponseUnion
}

// listItemUpdateResponseJSON contains the JSON metadata for the struct
// [ListItemUpdateResponse]
type listItemUpdateResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r listItemUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListItemUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListItemUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListItemUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ListItemUpdateResponseOperationID],
// [ListItemUpdateResponseOperationID].
func (r ListItemUpdateResponse) AsUnion() ListItemUpdateResponseUnion {
	return r.union
}

// Union satisfied by [ListItemUpdateResponseOperationID] or
// [ListItemUpdateResponseOperationID].
type ListItemUpdateResponseUnion interface {
	implementsListItemUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListItemUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemUpdateResponseOperationID{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemUpdateResponseOperationID{}),
		},
	)
}

type ListItemUpdateResponseOperationID struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                `json:"operation_id"`
	JSON        listItemUpdateResponseOperationIDJSON `json:"-"`
}

// listItemUpdateResponseOperationIDJSON contains the JSON metadata for the struct
// [ListItemUpdateResponseOperationID]
type listItemUpdateResponseOperationIDJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponseOperationID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseOperationIDJSON) RawJSON() string {
	return r.raw
}

func (r ListItemUpdateResponseOperationID) implementsListItemUpdateResponse() {}

type ListItemDeleteResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                     `json:"operation_id"`
	JSON        listItemDeleteResponseJSON `json:"-"`
	union       ListItemDeleteResponseUnion
}

// listItemDeleteResponseJSON contains the JSON metadata for the struct
// [ListItemDeleteResponse]
type listItemDeleteResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r listItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListItemDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListItemDeleteResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ListItemDeleteResponseOperationID],
// [ListItemDeleteResponseOperationID].
func (r ListItemDeleteResponse) AsUnion() ListItemDeleteResponseUnion {
	return r.union
}

// Union satisfied by [ListItemDeleteResponseOperationID] or
// [ListItemDeleteResponseOperationID].
type ListItemDeleteResponseUnion interface {
	implementsListItemDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListItemDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemDeleteResponseOperationID{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemDeleteResponseOperationID{}),
		},
	)
}

type ListItemDeleteResponseOperationID struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                                `json:"operation_id"`
	JSON        listItemDeleteResponseOperationIDJSON `json:"-"`
}

// listItemDeleteResponseOperationIDJSON contains the JSON metadata for the struct
// [ListItemDeleteResponseOperationID]
type listItemDeleteResponseOperationIDJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponseOperationID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseOperationIDJSON) RawJSON() string {
	return r.raw
}

func (r ListItemDeleteResponseOperationID) implementsListItemDeleteResponse() {}

type ListItemGetResponse struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// Defines a non-negative 32 bit integer.
	ASN int64 `json:"asn"`
	// Defines an informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname Hostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect Redirect                `json:"redirect"`
	JSON     listItemGetResponseJSON `json:"-"`
	union    ListItemGetResponseUnion
}

// listItemGetResponseJSON contains the JSON metadata for the struct
// [ListItemGetResponse]
type listItemGetResponseJSON struct {
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

func (r listItemGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListItemGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListItemGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ListItemGetResponseObject],
// [ListItemGetResponseObject].
func (r ListItemGetResponse) AsUnion() ListItemGetResponseUnion {
	return r.union
}

// Union satisfied by [ListItemGetResponseObject] or [ListItemGetResponseObject].
type ListItemGetResponseUnion interface {
	implementsListItemGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListItemGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemGetResponseObject{}),
		},
	)
}

type ListItemGetResponseObject struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// Defines a non-negative 32 bit integer.
	ASN int64 `json:"asn"`
	// Defines an informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname Hostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect Redirect                      `json:"redirect"`
	JSON     listItemGetResponseObjectJSON `json:"-"`
}

// listItemGetResponseObjectJSON contains the JSON metadata for the struct
// [ListItemGetResponseObject]
type listItemGetResponseObjectJSON struct {
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

func (r *ListItemGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r ListItemGetResponseObject) implementsListItemGetResponse() {}

type ListItemNewParams struct {
	// Defines an identifier.
	AccountID param.Field[string]     `path:"account_id,required"`
	Body      []ListItemNewParamsBody `json:"body,required"`
}

func (r ListItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListItemNewParamsBody struct {
	// Defines a non-negative 32 bit integer.
	ASN param.Field[int64] `json:"asn"`
	// Defines an informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[HostnameParam] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RedirectParam] `json:"redirect"`
}

func (r ListItemNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListItemNewResponse   `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
type ListItemNewResponseEnvelopeSuccess bool

const (
	ListItemNewResponseEnvelopeSuccessTrue ListItemNewResponseEnvelopeSuccess = true
)

func (r ListItemNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemUpdateParams struct {
	// Defines an identifier.
	AccountID param.Field[string]        `path:"account_id,required"`
	Body      []ListItemUpdateParamsBody `json:"body,required"`
}

func (r ListItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListItemUpdateParamsBody struct {
	// Defines a non-negative 32 bit integer.
	ASN param.Field[int64] `json:"asn"`
	// Defines an informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[HostnameParam] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RedirectParam] `json:"redirect"`
}

func (r ListItemUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   ListItemUpdateResponse `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
type ListItemUpdateResponseEnvelopeSuccess bool

const (
	ListItemUpdateResponseEnvelopeSuccessTrue ListItemUpdateResponseEnvelopeSuccess = true
)

func (r ListItemUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemListParams struct {
	// Defines an identifier.
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ListItemListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []interface{}         `json:"result,required"`
	// Defines whether the API call was successful.
	Success    ListItemListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ListItemListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       listItemListResponseEnvelopeJSON       `json:"-"`
}

// listItemListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemListResponseEnvelope]
type listItemListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListItemListResponseEnvelopeSuccess bool

const (
	ListItemListResponseEnvelopeSuccessTrue ListItemListResponseEnvelopeSuccess = true
)

func (r ListItemListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemListResponseEnvelopeResultInfo struct {
	Cursors ListCursor                                 `json:"cursors"`
	JSON    listItemListResponseEnvelopeResultInfoJSON `json:"-"`
}

// listItemListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ListItemListResponseEnvelopeResultInfo]
type listItemListResponseEnvelopeResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ListItemDeleteParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListItemDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   ListItemDeleteResponse `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
type ListItemDeleteResponseEnvelopeSuccess bool

const (
	ListItemDeleteResponseEnvelopeSuccessTrue ListItemDeleteResponseEnvelopeSuccess = true
)

func (r ListItemDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemGetParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListItemGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListItemGetResponse   `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
type ListItemGetResponseEnvelopeSuccess bool

const (
	ListItemGetResponseEnvelopeSuccessTrue ListItemGetResponseEnvelopeSuccess = true
)

func (r ListItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
