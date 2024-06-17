// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// IAMResourceGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIAMResourceGroupService] method instead.
type IAMResourceGroupService struct {
	Options []option.RequestOption
}

// NewIAMResourceGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIAMResourceGroupService(opts ...option.RequestOption) (r *IAMResourceGroupService) {
	r = &IAMResourceGroupService{}
	r.Options = opts
	return
}

// Create a new Resource Group under the specified account.
func (r *IAMResourceGroupService) New(ctx context.Context, params IAMResourceGroupNewParams, opts ...option.RequestOption) (res *IAMResourceGroupNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Modify an existing resource group.
func (r *IAMResourceGroupService) Update(ctx context.Context, resourceGroupID string, params IAMResourceGroupUpdateParams, opts ...option.RequestOption) (res *IAMResourceGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", params.AccountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List all the resource groups for an account.
func (r *IAMResourceGroupService) List(ctx context.Context, params IAMResourceGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[IAMResourceGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/iam/resource_groups", params.AccountID)
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

// List all the resource groups for an account.
func (r *IAMResourceGroupService) ListAutoPaging(ctx context.Context, params IAMResourceGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[IAMResourceGroupListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a resource group from an account.
func (r *IAMResourceGroupService) Delete(ctx context.Context, resourceGroupID string, body IAMResourceGroupDeleteParams, opts ...option.RequestOption) (res *IAMResourceGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IAMResourceGroupDeleteResponseEnvelope
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", body.AccountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific resource group in an account.
func (r *IAMResourceGroupService) Get(ctx context.Context, resourceGroupID string, query IAMResourceGroupGetParams, opts ...option.RequestOption) (res *IAMResourceGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", query.AccountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A group of scoped resources.
type IAMResourceGroupNewResponse struct {
	// Identifier of the group.
	ID string `json:"id"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// A scope is a combination of scope objects which provides additional context.
	Scope IAMResourceGroupNewResponseScope `json:"scope"`
	JSON  iamResourceGroupNewResponseJSON  `json:"-"`
}

// iamResourceGroupNewResponseJSON contains the JSON metadata for the struct
// [IAMResourceGroupNewResponse]
type iamResourceGroupNewResponseJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IAMResourceGroupNewResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects []IAMResourceGroupNewResponseScopeObject `json:"objects,required"`
	JSON    iamResourceGroupNewResponseScopeJSON     `json:"-"`
}

// iamResourceGroupNewResponseScopeJSON contains the JSON metadata for the struct
// [IAMResourceGroupNewResponseScope]
type iamResourceGroupNewResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupNewResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupNewResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IAMResourceGroupNewResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                     `json:"key,required"`
	JSON iamResourceGroupNewResponseScopeObjectJSON `json:"-"`
}

// iamResourceGroupNewResponseScopeObjectJSON contains the JSON metadata for the
// struct [IAMResourceGroupNewResponseScopeObject]
type iamResourceGroupNewResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupNewResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupNewResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type IAMResourceGroupUpdateResponse struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []IAMResourceGroupUpdateResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                             `json:"name"`
	JSON iamResourceGroupUpdateResponseJSON `json:"-"`
}

// iamResourceGroupUpdateResponseJSON contains the JSON metadata for the struct
// [IAMResourceGroupUpdateResponse]
type iamResourceGroupUpdateResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IAMResourceGroupUpdateResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []IAMResourceGroupUpdateResponseScopeObject `json:"objects,required"`
	JSON    iamResourceGroupUpdateResponseScopeJSON     `json:"-"`
}

// iamResourceGroupUpdateResponseScopeJSON contains the JSON metadata for the
// struct [IAMResourceGroupUpdateResponseScope]
type iamResourceGroupUpdateResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupUpdateResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupUpdateResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IAMResourceGroupUpdateResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                        `json:"key,required"`
	JSON iamResourceGroupUpdateResponseScopeObjectJSON `json:"-"`
}

// iamResourceGroupUpdateResponseScopeObjectJSON contains the JSON metadata for the
// struct [IAMResourceGroupUpdateResponseScopeObject]
type iamResourceGroupUpdateResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupUpdateResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupUpdateResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

type IAMResourceGroupListResponse = interface{}

type IAMResourceGroupDeleteResponse struct {
	// Identifier
	ID   string                             `json:"id,required"`
	JSON iamResourceGroupDeleteResponseJSON `json:"-"`
}

// iamResourceGroupDeleteResponseJSON contains the JSON metadata for the struct
// [IAMResourceGroupDeleteResponse]
type iamResourceGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type IAMResourceGroupGetResponse struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []IAMResourceGroupGetResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                          `json:"name"`
	JSON iamResourceGroupGetResponseJSON `json:"-"`
}

// iamResourceGroupGetResponseJSON contains the JSON metadata for the struct
// [IAMResourceGroupGetResponse]
type iamResourceGroupGetResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IAMResourceGroupGetResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []IAMResourceGroupGetResponseScopeObject `json:"objects,required"`
	JSON    iamResourceGroupGetResponseScopeJSON     `json:"-"`
}

// iamResourceGroupGetResponseScopeJSON contains the JSON metadata for the struct
// [IAMResourceGroupGetResponseScope]
type iamResourceGroupGetResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupGetResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupGetResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IAMResourceGroupGetResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                     `json:"key,required"`
	JSON iamResourceGroupGetResponseScopeObjectJSON `json:"-"`
}

// iamResourceGroupGetResponseScopeObjectJSON contains the JSON metadata for the
// struct [IAMResourceGroupGetResponseScopeObject]
type iamResourceGroupGetResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupGetResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupGetResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

type IAMResourceGroupNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[IAMResourceGroupNewParamsScope] `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta param.Field[interface{}] `json:"meta"`
}

func (r IAMResourceGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope is a combination of scope objects which provides additional context.
type IAMResourceGroupNewParamsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]IAMResourceGroupNewParamsScopeObject] `json:"objects,required"`
}

func (r IAMResourceGroupNewParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type IAMResourceGroupNewParamsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r IAMResourceGroupNewParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IAMResourceGroupUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[IAMResourceGroupUpdateParamsScope] `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta param.Field[interface{}] `json:"meta"`
}

func (r IAMResourceGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope is a combination of scope objects which provides additional context.
type IAMResourceGroupUpdateParamsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]IAMResourceGroupUpdateParamsScopeObject] `json:"objects,required"`
}

func (r IAMResourceGroupUpdateParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type IAMResourceGroupUpdateParamsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r IAMResourceGroupUpdateParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IAMResourceGroupListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// ID of the resource group to be fetched.
	ID param.Field[string] `query:"id"`
	// Name of the resource group to be fetched.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [IAMResourceGroupListParams]'s query parameters as
// `url.Values`.
func (r IAMResourceGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IAMResourceGroupDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type IAMResourceGroupDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IAMResourceGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  IAMResourceGroupDeleteResponse                `json:"result,nullable"`
	JSON    iamResourceGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// iamResourceGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [IAMResourceGroupDeleteResponseEnvelope]
type iamResourceGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMResourceGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IAMResourceGroupDeleteResponseEnvelopeSuccess bool

const (
	IAMResourceGroupDeleteResponseEnvelopeSuccessTrue IAMResourceGroupDeleteResponseEnvelopeSuccess = true
)

func (r IAMResourceGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IAMResourceGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IAMResourceGroupGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
