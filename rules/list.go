// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ListService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewListService] method instead.
type ListService struct {
	Options        []option.RequestOption
	BulkOperations *ListBulkOperationService
	Items          *ListItemService
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r *ListService) {
	r = &ListService{}
	r.Options = opts
	r.BulkOperations = NewListBulkOperationService(opts...)
	r.Items = NewListItemService(opts...)
	return
}

// Creates a new list of the specified type.
func (r *ListService) New(ctx context.Context, params ListNewParams, opts ...option.RequestOption) (res *ListsList, err error) {
	opts = append(r.Options[:], opts...)
	var env ListNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the description of a list.
func (r *ListService) Update(ctx context.Context, listID string, params ListUpdateParams, opts ...option.RequestOption) (res *ListsList, err error) {
	opts = append(r.Options[:], opts...)
	var env ListUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all lists in the account.
func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ListsList], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches all lists in the account.
func (r *ListService) ListAutoPaging(ctx context.Context, query ListListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ListsList] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific list and all its items.
func (r *ListService) Delete(ctx context.Context, listID string, params ListDeleteParams, opts ...option.RequestOption) (res *ListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ListDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a list.
func (r *ListService) Get(ctx context.Context, listID string, query ListGetParams, opts ...option.RequestOption) (res *ListsList, err error) {
	opts = append(r.Options[:], opts...)
	var env ListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ListsList struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListsListKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64       `json:"num_referencing_filters"`
	JSON                  listsListJSON `json:"-"`
}

// listsListJSON contains the JSON metadata for the struct [ListsList]
type listsListJSON struct {
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

func (r *ListsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listsListJSON) RawJSON() string {
	return r.raw
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListsListKind string

const (
	ListsListKindIP       ListsListKind = "ip"
	ListsListKindRedirect ListsListKind = "redirect"
	ListsListKindHostname ListsListKind = "hostname"
	ListsListKindASN      ListsListKind = "asn"
)

func (r ListsListKind) IsKnown() bool {
	switch r {
	case ListsListKindIP, ListsListKindRedirect, ListsListKindHostname, ListsListKindASN:
		return true
	}
	return false
}

type ListNewResponse = interface{}

type ListUpdateResponse = interface{}

type ListDeleteResponse struct {
	// The unique ID of the item in the List.
	ID   string                 `json:"id"`
	JSON listDeleteResponseJSON `json:"-"`
}

// listDeleteResponseJSON contains the JSON metadata for the struct
// [ListDeleteResponse]
type listDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ListGetResponse = interface{}

type ListNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind param.Field[ListNewParamsKind] `json:"kind,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name param.Field[string] `json:"name,required"`
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r ListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListNewParamsKind string

const (
	ListNewParamsKindIP       ListNewParamsKind = "ip"
	ListNewParamsKindRedirect ListNewParamsKind = "redirect"
	ListNewParamsKindHostname ListNewParamsKind = "hostname"
	ListNewParamsKindASN      ListNewParamsKind = "asn"
)

func (r ListNewParamsKind) IsKnown() bool {
	switch r {
	case ListNewParamsKindIP, ListNewParamsKindRedirect, ListNewParamsKindHostname, ListNewParamsKindASN:
		return true
	}
	return false
}

type ListNewResponseEnvelope struct {
	Errors   []ListNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ListsList                         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    listNewResponseEnvelopeJSON    `json:"-"`
}

// listNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListNewResponseEnvelope]
type listNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    listNewResponseEnvelopeErrorsJSON `json:"-"`
}

// listNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ListNewResponseEnvelopeErrors]
type listNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    listNewResponseEnvelopeMessagesJSON `json:"-"`
}

// listNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ListNewResponseEnvelopeMessages]
type listNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListNewResponseEnvelopeSuccess bool

const (
	ListNewResponseEnvelopeSuccessTrue ListNewResponseEnvelopeSuccess = true
)

func (r ListNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r ListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListUpdateResponseEnvelope struct {
	Errors   []ListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ListsList                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    listUpdateResponseEnvelopeJSON    `json:"-"`
}

// listUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListUpdateResponseEnvelope]
type listUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    listUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// listUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ListUpdateResponseEnvelopeErrors]
type listUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    listUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// listUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ListUpdateResponseEnvelopeMessages]
type listUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListUpdateResponseEnvelopeSuccess bool

const (
	ListUpdateResponseEnvelopeSuccessTrue ListUpdateResponseEnvelopeSuccess = true
)

func (r ListUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ListDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListDeleteResponseEnvelope struct {
	Errors   []ListDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ListDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    listDeleteResponseEnvelopeJSON    `json:"-"`
}

// listDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListDeleteResponseEnvelope]
type listDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    listDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// listDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ListDeleteResponseEnvelopeErrors]
type listDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    listDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// listDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ListDeleteResponseEnvelopeMessages]
type listDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListDeleteResponseEnvelopeSuccess bool

const (
	ListDeleteResponseEnvelopeSuccessTrue ListDeleteResponseEnvelopeSuccess = true
)

func (r ListDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListGetResponseEnvelope struct {
	Errors   []ListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ListsList                         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    listGetResponseEnvelopeJSON    `json:"-"`
}

// listGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListGetResponseEnvelope]
type listGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    listGetResponseEnvelopeErrorsJSON `json:"-"`
}

// listGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ListGetResponseEnvelopeErrors]
type listGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    listGetResponseEnvelopeMessagesJSON `json:"-"`
}

// listGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ListGetResponseEnvelopeMessages]
type listGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListGetResponseEnvelopeSuccess bool

const (
	ListGetResponseEnvelopeSuccessTrue ListGetResponseEnvelopeSuccess = true
)

func (r ListGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
