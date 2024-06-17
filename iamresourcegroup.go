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

// IamResourceGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIamResourceGroupService] method instead.
type IamResourceGroupService struct {
	Options []option.RequestOption
}

// NewIamResourceGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIamResourceGroupService(opts ...option.RequestOption) (r *IamResourceGroupService) {
	r = &IamResourceGroupService{}
	r.Options = opts
	return
}

// Create a new Resource Group under the specified account.
func (r *IamResourceGroupService) New(ctx context.Context, params IamResourceGroupNewParams, opts ...option.RequestOption) (res *IamResourceGroupNewResponse, err error) {
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
func (r *IamResourceGroupService) Update(ctx context.Context, resourceGroupID string, params IamResourceGroupUpdateParams, opts ...option.RequestOption) (res *IamResourceGroupUpdateResponse, err error) {
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
func (r *IamResourceGroupService) List(ctx context.Context, params IamResourceGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[IamResourceGroupListResponse], err error) {
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
func (r *IamResourceGroupService) ListAutoPaging(ctx context.Context, params IamResourceGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[IamResourceGroupListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a resource group from an account.
func (r *IamResourceGroupService) Delete(ctx context.Context, resourceGroupID string, body IamResourceGroupDeleteParams, opts ...option.RequestOption) (res *IamResourceGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IamResourceGroupDeleteResponseEnvelope
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
func (r *IamResourceGroupService) Get(ctx context.Context, resourceGroupID string, query IamResourceGroupGetParams, opts ...option.RequestOption) (res *IamResourceGroupGetResponse, err error) {
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
type IamResourceGroupNewResponse struct {
	// Identifier of the group.
	ID string `json:"id"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// A scope is a combination of scope objects which provides additional context.
	Scope IamResourceGroupNewResponseScope `json:"scope"`
	JSON  iamResourceGroupNewResponseJSON  `json:"-"`
}

// iamResourceGroupNewResponseJSON contains the JSON metadata for the struct
// [IamResourceGroupNewResponse]
type iamResourceGroupNewResponseJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IamResourceGroupNewResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects []IamResourceGroupNewResponseScopeObject `json:"objects,required"`
	JSON    iamResourceGroupNewResponseScopeJSON     `json:"-"`
}

// iamResourceGroupNewResponseScopeJSON contains the JSON metadata for the struct
// [IamResourceGroupNewResponseScope]
type iamResourceGroupNewResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupNewResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupNewResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamResourceGroupNewResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                     `json:"key,required"`
	JSON iamResourceGroupNewResponseScopeObjectJSON `json:"-"`
}

// iamResourceGroupNewResponseScopeObjectJSON contains the JSON metadata for the
// struct [IamResourceGroupNewResponseScopeObject]
type iamResourceGroupNewResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupNewResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupNewResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type IamResourceGroupUpdateResponse struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []IamResourceGroupUpdateResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                             `json:"name"`
	JSON iamResourceGroupUpdateResponseJSON `json:"-"`
}

// iamResourceGroupUpdateResponseJSON contains the JSON metadata for the struct
// [IamResourceGroupUpdateResponse]
type iamResourceGroupUpdateResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IamResourceGroupUpdateResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []IamResourceGroupUpdateResponseScopeObject `json:"objects,required"`
	JSON    iamResourceGroupUpdateResponseScopeJSON     `json:"-"`
}

// iamResourceGroupUpdateResponseScopeJSON contains the JSON metadata for the
// struct [IamResourceGroupUpdateResponseScope]
type iamResourceGroupUpdateResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupUpdateResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupUpdateResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamResourceGroupUpdateResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                        `json:"key,required"`
	JSON iamResourceGroupUpdateResponseScopeObjectJSON `json:"-"`
}

// iamResourceGroupUpdateResponseScopeObjectJSON contains the JSON metadata for the
// struct [IamResourceGroupUpdateResponseScopeObject]
type iamResourceGroupUpdateResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupUpdateResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupUpdateResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

type IamResourceGroupListResponse = interface{}

type IamResourceGroupDeleteResponse struct {
	// Identifier
	ID   string                             `json:"id,required"`
	JSON iamResourceGroupDeleteResponseJSON `json:"-"`
}

// iamResourceGroupDeleteResponseJSON contains the JSON metadata for the struct
// [IamResourceGroupDeleteResponse]
type iamResourceGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type IamResourceGroupGetResponse struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []IamResourceGroupGetResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                          `json:"name"`
	JSON iamResourceGroupGetResponseJSON `json:"-"`
}

// iamResourceGroupGetResponseJSON contains the JSON metadata for the struct
// [IamResourceGroupGetResponse]
type iamResourceGroupGetResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IamResourceGroupGetResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []IamResourceGroupGetResponseScopeObject `json:"objects,required"`
	JSON    iamResourceGroupGetResponseScopeJSON     `json:"-"`
}

// iamResourceGroupGetResponseScopeJSON contains the JSON metadata for the struct
// [IamResourceGroupGetResponseScope]
type iamResourceGroupGetResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupGetResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupGetResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamResourceGroupGetResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                     `json:"key,required"`
	JSON iamResourceGroupGetResponseScopeObjectJSON `json:"-"`
}

// iamResourceGroupGetResponseScopeObjectJSON contains the JSON metadata for the
// struct [IamResourceGroupGetResponseScopeObject]
type iamResourceGroupGetResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupGetResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupGetResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

type IamResourceGroupNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[IamResourceGroupNewParamsScope] `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta param.Field[interface{}] `json:"meta"`
}

func (r IamResourceGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope is a combination of scope objects which provides additional context.
type IamResourceGroupNewParamsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]IamResourceGroupNewParamsScopeObject] `json:"objects,required"`
}

func (r IamResourceGroupNewParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamResourceGroupNewParamsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r IamResourceGroupNewParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamResourceGroupUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[IamResourceGroupUpdateParamsScope] `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta param.Field[interface{}] `json:"meta"`
}

func (r IamResourceGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope is a combination of scope objects which provides additional context.
type IamResourceGroupUpdateParamsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]IamResourceGroupUpdateParamsScopeObject] `json:"objects,required"`
}

func (r IamResourceGroupUpdateParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamResourceGroupUpdateParamsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r IamResourceGroupUpdateParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamResourceGroupListParams struct {
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

// URLQuery serializes [IamResourceGroupListParams]'s query parameters as
// `url.Values`.
func (r IamResourceGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IamResourceGroupDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type IamResourceGroupDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IamResourceGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  IamResourceGroupDeleteResponse                `json:"result,nullable"`
	JSON    iamResourceGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// iamResourceGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [IamResourceGroupDeleteResponseEnvelope]
type iamResourceGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IamResourceGroupDeleteResponseEnvelopeSuccess bool

const (
	IamResourceGroupDeleteResponseEnvelopeSuccessTrue IamResourceGroupDeleteResponseEnvelopeSuccess = true
)

func (r IamResourceGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IamResourceGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IamResourceGroupGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
