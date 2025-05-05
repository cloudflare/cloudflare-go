// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
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
func (r *ListService) New(ctx context.Context, params ListNewParams, opts ...option.RequestOption) (res *ListNewResponse, err error) {
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
func (r *ListService) Update(ctx context.Context, listID string, params ListUpdateParams, opts ...option.RequestOption) (res *ListUpdateResponse, err error) {
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
func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *[]interface{}, err error) {
	var env ListListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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
func (r *ListService) Get(ctx context.Context, listID string, query ListGetParams, opts ...option.RequestOption) (res *ListGetResponse, err error) {
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

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type HostnameParam struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
}

func (r HostnameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type ListNewResponse struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListNewResponseKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64             `json:"num_referencing_filters"`
	JSON                  listNewResponseJSON `json:"-"`
	union                 ListNewResponseUnion
}

// listNewResponseJSON contains the JSON metadata for the struct [ListNewResponse]
type listNewResponseJSON struct {
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

func (r listNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [rules.ListNewResponseObject],
// [rules.ListNewResponseObject].
func (r ListNewResponse) AsUnion() ListNewResponseUnion {
	return r.union
}

// Union satisfied by [rules.ListNewResponseObject] or
// [rules.ListNewResponseObject].
type ListNewResponseUnion interface {
	implementsListNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListNewResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListNewResponseObject{}),
		},
	)
}

type ListNewResponseObject struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListNewResponseObjectKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64                   `json:"num_referencing_filters"`
	JSON                  listNewResponseObjectJSON `json:"-"`
}

// listNewResponseObjectJSON contains the JSON metadata for the struct
// [ListNewResponseObject]
type listNewResponseObjectJSON struct {
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

func (r *ListNewResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listNewResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r ListNewResponseObject) implementsListNewResponse() {}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListNewResponseObjectKind string

const (
	ListNewResponseObjectKindIP       ListNewResponseObjectKind = "ip"
	ListNewResponseObjectKindRedirect ListNewResponseObjectKind = "redirect"
	ListNewResponseObjectKindHostname ListNewResponseObjectKind = "hostname"
	ListNewResponseObjectKindASN      ListNewResponseObjectKind = "asn"
)

func (r ListNewResponseObjectKind) IsKnown() bool {
	switch r {
	case ListNewResponseObjectKindIP, ListNewResponseObjectKindRedirect, ListNewResponseObjectKindHostname, ListNewResponseObjectKindASN:
		return true
	}
	return false
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListNewResponseKind string

const (
	ListNewResponseKindIP       ListNewResponseKind = "ip"
	ListNewResponseKindRedirect ListNewResponseKind = "redirect"
	ListNewResponseKindHostname ListNewResponseKind = "hostname"
	ListNewResponseKindASN      ListNewResponseKind = "asn"
)

func (r ListNewResponseKind) IsKnown() bool {
	switch r {
	case ListNewResponseKindIP, ListNewResponseKindRedirect, ListNewResponseKindHostname, ListNewResponseKindASN:
		return true
	}
	return false
}

type ListUpdateResponse struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListUpdateResponseKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64                `json:"num_referencing_filters"`
	JSON                  listUpdateResponseJSON `json:"-"`
	union                 ListUpdateResponseUnion
}

// listUpdateResponseJSON contains the JSON metadata for the struct
// [ListUpdateResponse]
type listUpdateResponseJSON struct {
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

func (r listUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListUpdateResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [rules.ListUpdateResponseObject],
// [rules.ListUpdateResponseObject].
func (r ListUpdateResponse) AsUnion() ListUpdateResponseUnion {
	return r.union
}

// Union satisfied by [rules.ListUpdateResponseObject] or
// [rules.ListUpdateResponseObject].
type ListUpdateResponseUnion interface {
	implementsListUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListUpdateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListUpdateResponseObject{}),
		},
	)
}

type ListUpdateResponseObject struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListUpdateResponseObjectKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64                      `json:"num_referencing_filters"`
	JSON                  listUpdateResponseObjectJSON `json:"-"`
}

// listUpdateResponseObjectJSON contains the JSON metadata for the struct
// [ListUpdateResponseObject]
type listUpdateResponseObjectJSON struct {
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

func (r *ListUpdateResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listUpdateResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r ListUpdateResponseObject) implementsListUpdateResponse() {}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListUpdateResponseObjectKind string

const (
	ListUpdateResponseObjectKindIP       ListUpdateResponseObjectKind = "ip"
	ListUpdateResponseObjectKindRedirect ListUpdateResponseObjectKind = "redirect"
	ListUpdateResponseObjectKindHostname ListUpdateResponseObjectKind = "hostname"
	ListUpdateResponseObjectKindASN      ListUpdateResponseObjectKind = "asn"
)

func (r ListUpdateResponseObjectKind) IsKnown() bool {
	switch r {
	case ListUpdateResponseObjectKindIP, ListUpdateResponseObjectKindRedirect, ListUpdateResponseObjectKindHostname, ListUpdateResponseObjectKindASN:
		return true
	}
	return false
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListUpdateResponseKind string

const (
	ListUpdateResponseKindIP       ListUpdateResponseKind = "ip"
	ListUpdateResponseKindRedirect ListUpdateResponseKind = "redirect"
	ListUpdateResponseKindHostname ListUpdateResponseKind = "hostname"
	ListUpdateResponseKindASN      ListUpdateResponseKind = "asn"
)

func (r ListUpdateResponseKind) IsKnown() bool {
	switch r {
	case ListUpdateResponseKindIP, ListUpdateResponseKindRedirect, ListUpdateResponseKindHostname, ListUpdateResponseKindASN:
		return true
	}
	return false
}

type ListDeleteResponse struct {
	// Defines the unique ID of the item in the List.
	ID    string                 `json:"id"`
	JSON  listDeleteResponseJSON `json:"-"`
	union ListDeleteResponseUnion
}

// listDeleteResponseJSON contains the JSON metadata for the struct
// [ListDeleteResponse]
type listDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r listDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListDeleteResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [rules.ListDeleteResponseID],
// [rules.ListDeleteResponseID].
func (r ListDeleteResponse) AsUnion() ListDeleteResponseUnion {
	return r.union
}

// Union satisfied by [rules.ListDeleteResponseID] or [rules.ListDeleteResponseID].
type ListDeleteResponseUnion interface {
	implementsListDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListDeleteResponseID{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListDeleteResponseID{}),
		},
	)
}

type ListDeleteResponseID struct {
	// Defines the unique ID of the item in the List.
	ID   string                   `json:"id"`
	JSON listDeleteResponseIDJSON `json:"-"`
}

// listDeleteResponseIDJSON contains the JSON metadata for the struct
// [ListDeleteResponseID]
type listDeleteResponseIDJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListDeleteResponseID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listDeleteResponseIDJSON) RawJSON() string {
	return r.raw
}

func (r ListDeleteResponseID) implementsListDeleteResponse() {}

type ListGetResponse struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListGetResponseKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64             `json:"num_referencing_filters"`
	JSON                  listGetResponseJSON `json:"-"`
	union                 ListGetResponseUnion
}

// listGetResponseJSON contains the JSON metadata for the struct [ListGetResponse]
type listGetResponseJSON struct {
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

func (r listGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ListGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ListGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [rules.ListGetResponseObject],
// [rules.ListGetResponseObject].
func (r ListGetResponse) AsUnion() ListGetResponseUnion {
	return r.union
}

// Union satisfied by [rules.ListGetResponseObject] or
// [rules.ListGetResponseObject].
type ListGetResponseUnion interface {
	implementsListGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListGetResponseObject{}),
		},
	)
}

type ListGetResponseObject struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListGetResponseObjectKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64                   `json:"num_referencing_filters"`
	JSON                  listGetResponseObjectJSON `json:"-"`
}

// listGetResponseObjectJSON contains the JSON metadata for the struct
// [ListGetResponseObject]
type listGetResponseObjectJSON struct {
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

func (r *ListGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r ListGetResponseObject) implementsListGetResponse() {}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListGetResponseObjectKind string

const (
	ListGetResponseObjectKindIP       ListGetResponseObjectKind = "ip"
	ListGetResponseObjectKindRedirect ListGetResponseObjectKind = "redirect"
	ListGetResponseObjectKindHostname ListGetResponseObjectKind = "hostname"
	ListGetResponseObjectKindASN      ListGetResponseObjectKind = "asn"
)

func (r ListGetResponseObjectKind) IsKnown() bool {
	switch r {
	case ListGetResponseObjectKindIP, ListGetResponseObjectKindRedirect, ListGetResponseObjectKindHostname, ListGetResponseObjectKindASN:
		return true
	}
	return false
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListGetResponseKind string

const (
	ListGetResponseKindIP       ListGetResponseKind = "ip"
	ListGetResponseKindRedirect ListGetResponseKind = "redirect"
	ListGetResponseKindHostname ListGetResponseKind = "hostname"
	ListGetResponseKindASN      ListGetResponseKind = "asn"
)

func (r ListGetResponseKind) IsKnown() bool {
	switch r {
	case ListGetResponseKindIP, ListGetResponseKindRedirect, ListGetResponseKindHostname, ListGetResponseKindASN:
		return true
	}
	return false
}

type ListNewParams struct {
	// Defines an identifier.
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
	Result   ListNewResponse       `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
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
	// Defines an identifier.
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
	Result   ListUpdateResponse    `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
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
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []interface{}         `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListListResponseEnvelopeSuccess `json:"success,required"`
	JSON    listListResponseEnvelopeJSON    `json:"-"`
}

// listListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListListResponseEnvelope]
type listListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListListResponseEnvelopeSuccess bool

const (
	ListListResponseEnvelopeSuccessTrue ListListResponseEnvelopeSuccess = true
)

func (r ListListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListDeleteParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListDeleteResponse    `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
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
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListGetResponse       `json:"result,required"`
	// Defines whether the API call was successful.
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

// Defines whether the API call was successful.
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
