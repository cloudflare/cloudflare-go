// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// ListService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
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
	var env ListNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
	var env ListUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *ListService) Delete(ctx context.Context, listID string, body ListDeleteParams, opts ...option.RequestOption) (res *ListDeleteResponse, err error) {
	var env ListDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", body.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a list.
func (r *ListService) Get(ctx context.Context, listID string, query ListGetParams, opts ...option.RequestOption) (res *ListsList, err error) {
	var env ListGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type Hostname struct {
	URLHostname string       `json:"url_hostname,required"`
	JSON        hostnameJSON `json:"-"`
}

// hostnameJSON contains the JSON metadata for the struct [Hostname]
type hostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Hostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameJSON) RawJSON() string {
	return r.raw
}

func (r Hostname) ImplementsRulesListItemGetResponseUnion() {}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type HostnameParam struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r HostnameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// The definition of the redirect.
type Redirect struct {
	SourceURL           string             `json:"source_url,required"`
	TargetURL           string             `json:"target_url,required"`
	IncludeSubdomains   bool               `json:"include_subdomains"`
	PreservePathSuffix  bool               `json:"preserve_path_suffix"`
	PreserveQueryString bool               `json:"preserve_query_string"`
	StatusCode          RedirectStatusCode `json:"status_code"`
	SubpathMatching     bool               `json:"subpath_matching"`
	JSON                redirectJSON       `json:"-"`
}

// redirectJSON contains the JSON metadata for the struct [Redirect]
type redirectJSON struct {
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

func (r *Redirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectJSON) RawJSON() string {
	return r.raw
}

func (r Redirect) ImplementsRulesListItemGetResponseUnion() {}

type RedirectStatusCode int64

const (
	RedirectStatusCode301 RedirectStatusCode = 301
	RedirectStatusCode302 RedirectStatusCode = 302
	RedirectStatusCode307 RedirectStatusCode = 307
	RedirectStatusCode308 RedirectStatusCode = 308
)

func (r RedirectStatusCode) IsKnown() bool {
	switch r {
	case RedirectStatusCode301, RedirectStatusCode302, RedirectStatusCode307, RedirectStatusCode308:
		return true
	}
	return false
}

// The definition of the redirect.
type RedirectParam struct {
	SourceURL           param.Field[string]             `json:"source_url,required"`
	TargetURL           param.Field[string]             `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]               `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]               `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]               `json:"preserve_query_string"`
	StatusCode          param.Field[RedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]               `json:"subpath_matching"`
}

func (r RedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListsList             `json:"result,required,nullable"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListsList             `json:"result,required,nullable"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListDeleteResponse    `json:"result,required,nullable"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListsList             `json:"result,required,nullable"`
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
