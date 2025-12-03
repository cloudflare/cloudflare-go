// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ResourceGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceGroupService] method instead.
type ResourceGroupService struct {
	Options []option.RequestOption
}

// NewResourceGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceGroupService(opts ...option.RequestOption) (r *ResourceGroupService) {
	r = &ResourceGroupService{}
	r.Options = opts
	return
}

// Create a new Resource Group under the specified account.
func (r *ResourceGroupService) New(ctx context.Context, params ResourceGroupNewParams, opts ...option.RequestOption) (res *ResourceGroupNewResponse, err error) {
	var env ResourceGroupNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an existing resource group.
func (r *ResourceGroupService) Update(ctx context.Context, resourceGroupID string, params ResourceGroupUpdateParams, opts ...option.RequestOption) (res *ResourceGroupUpdateResponse, err error) {
	var env ResourceGroupUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", params.AccountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all the resource groups for an account.
func (r *ResourceGroupService) List(ctx context.Context, params ResourceGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ResourceGroupListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
func (r *ResourceGroupService) ListAutoPaging(ctx context.Context, params ResourceGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ResourceGroupListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Remove a resource group from an account.
func (r *ResourceGroupService) Delete(ctx context.Context, resourceGroupID string, body ResourceGroupDeleteParams, opts ...option.RequestOption) (res *ResourceGroupDeleteResponse, err error) {
	var env ResourceGroupDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
func (r *ResourceGroupService) Get(ctx context.Context, resourceGroupID string, query ResourceGroupGetParams, opts ...option.RequestOption) (res *ResourceGroupGetResponse, err error) {
	var env ResourceGroupGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", query.AccountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A group of scoped resources.
type ResourceGroupNewResponse struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []ResourceGroupNewResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta ResourceGroupNewResponseMeta `json:"meta"`
	// Name of the resource group.
	Name string                       `json:"name"`
	JSON resourceGroupNewResponseJSON `json:"-"`
}

// resourceGroupNewResponseJSON contains the JSON metadata for the struct
// [ResourceGroupNewResponse]
type resourceGroupNewResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type ResourceGroupNewResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []ResourceGroupNewResponseScopeObject `json:"objects,required"`
	JSON    resourceGroupNewResponseScopeJSON     `json:"-"`
}

// resourceGroupNewResponseScopeJSON contains the JSON metadata for the struct
// [ResourceGroupNewResponseScope]
type resourceGroupNewResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type ResourceGroupNewResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                  `json:"key,required"`
	JSON resourceGroupNewResponseScopeObjectJSON `json:"-"`
}

// resourceGroupNewResponseScopeObjectJSON contains the JSON metadata for the
// struct [ResourceGroupNewResponseScopeObject]
type resourceGroupNewResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type ResourceGroupNewResponseMeta struct {
	Key   string                           `json:"key"`
	Value string                           `json:"value"`
	JSON  resourceGroupNewResponseMetaJSON `json:"-"`
}

// resourceGroupNewResponseMetaJSON contains the JSON metadata for the struct
// [ResourceGroupNewResponseMeta]
type resourceGroupNewResponseMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type ResourceGroupUpdateResponse struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []ResourceGroupUpdateResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta ResourceGroupUpdateResponseMeta `json:"meta"`
	// Name of the resource group.
	Name string                          `json:"name"`
	JSON resourceGroupUpdateResponseJSON `json:"-"`
}

// resourceGroupUpdateResponseJSON contains the JSON metadata for the struct
// [ResourceGroupUpdateResponse]
type resourceGroupUpdateResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type ResourceGroupUpdateResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []ResourceGroupUpdateResponseScopeObject `json:"objects,required"`
	JSON    resourceGroupUpdateResponseScopeJSON     `json:"-"`
}

// resourceGroupUpdateResponseScopeJSON contains the JSON metadata for the struct
// [ResourceGroupUpdateResponseScope]
type resourceGroupUpdateResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type ResourceGroupUpdateResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                     `json:"key,required"`
	JSON resourceGroupUpdateResponseScopeObjectJSON `json:"-"`
}

// resourceGroupUpdateResponseScopeObjectJSON contains the JSON metadata for the
// struct [ResourceGroupUpdateResponseScopeObject]
type resourceGroupUpdateResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type ResourceGroupUpdateResponseMeta struct {
	Key   string                              `json:"key"`
	Value string                              `json:"value"`
	JSON  resourceGroupUpdateResponseMetaJSON `json:"-"`
}

// resourceGroupUpdateResponseMetaJSON contains the JSON metadata for the struct
// [ResourceGroupUpdateResponseMeta]
type resourceGroupUpdateResponseMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type ResourceGroupListResponse struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []ResourceGroupListResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta ResourceGroupListResponseMeta `json:"meta"`
	// Name of the resource group.
	Name string                        `json:"name"`
	JSON resourceGroupListResponseJSON `json:"-"`
}

// resourceGroupListResponseJSON contains the JSON metadata for the struct
// [ResourceGroupListResponse]
type resourceGroupListResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupListResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type ResourceGroupListResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []ResourceGroupListResponseScopeObject `json:"objects,required"`
	JSON    resourceGroupListResponseScopeJSON     `json:"-"`
}

// resourceGroupListResponseScopeJSON contains the JSON metadata for the struct
// [ResourceGroupListResponseScope]
type resourceGroupListResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupListResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupListResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type ResourceGroupListResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                   `json:"key,required"`
	JSON resourceGroupListResponseScopeObjectJSON `json:"-"`
}

// resourceGroupListResponseScopeObjectJSON contains the JSON metadata for the
// struct [ResourceGroupListResponseScopeObject]
type resourceGroupListResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupListResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupListResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type ResourceGroupListResponseMeta struct {
	Key   string                            `json:"key"`
	Value string                            `json:"value"`
	JSON  resourceGroupListResponseMetaJSON `json:"-"`
}

// resourceGroupListResponseMetaJSON contains the JSON metadata for the struct
// [ResourceGroupListResponseMeta]
type resourceGroupListResponseMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupListResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupDeleteResponse struct {
	// Identifier
	ID   string                          `json:"id,required"`
	JSON resourceGroupDeleteResponseJSON `json:"-"`
}

// resourceGroupDeleteResponseJSON contains the JSON metadata for the struct
// [ResourceGroupDeleteResponse]
type resourceGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type ResourceGroupGetResponse struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []ResourceGroupGetResponseScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta ResourceGroupGetResponseMeta `json:"meta"`
	// Name of the resource group.
	Name string                       `json:"name"`
	JSON resourceGroupGetResponseJSON `json:"-"`
}

// resourceGroupGetResponseJSON contains the JSON metadata for the struct
// [ResourceGroupGetResponse]
type resourceGroupGetResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type ResourceGroupGetResponseScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []ResourceGroupGetResponseScopeObject `json:"objects,required"`
	JSON    resourceGroupGetResponseScopeJSON     `json:"-"`
}

// resourceGroupGetResponseScopeJSON contains the JSON metadata for the struct
// [ResourceGroupGetResponseScope]
type resourceGroupGetResponseScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type ResourceGroupGetResponseScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                  `json:"key,required"`
	JSON resourceGroupGetResponseScopeObjectJSON `json:"-"`
}

// resourceGroupGetResponseScopeObjectJSON contains the JSON metadata for the
// struct [ResourceGroupGetResponseScopeObject]
type resourceGroupGetResponseScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type ResourceGroupGetResponseMeta struct {
	Key   string                           `json:"key"`
	Value string                           `json:"value"`
	JSON  resourceGroupGetResponseMetaJSON `json:"-"`
}

// resourceGroupGetResponseMetaJSON contains the JSON metadata for the struct
// [ResourceGroupGetResponseMeta]
type resourceGroupGetResponseMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the resource group
	Name param.Field[string] `json:"name,required"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[ResourceGroupNewParamsScope] `json:"scope,required"`
}

func (r ResourceGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope is a combination of scope objects which provides additional context.
type ResourceGroupNewParamsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]ResourceGroupNewParamsScopeObject] `json:"objects,required"`
}

func (r ResourceGroupNewParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type ResourceGroupNewParamsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r ResourceGroupNewParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceGroupNewResponseEnvelope struct {
	Errors   []ResourceGroupNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ResourceGroupNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ResourceGroupNewResponseEnvelopeSuccess `json:"success,required"`
	// A group of scoped resources.
	Result ResourceGroupNewResponse             `json:"result"`
	JSON   resourceGroupNewResponseEnvelopeJSON `json:"-"`
}

// resourceGroupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceGroupNewResponseEnvelope]
type resourceGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ResourceGroupNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceGroupNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ResourceGroupNewResponseEnvelopeErrors]
type resourceGroupNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    resourceGroupNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceGroupNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ResourceGroupNewResponseEnvelopeErrorsSource]
type resourceGroupNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ResourceGroupNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceGroupNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ResourceGroupNewResponseEnvelopeMessages]
type resourceGroupNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    resourceGroupNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceGroupNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ResourceGroupNewResponseEnvelopeMessagesSource]
type resourceGroupNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResourceGroupNewResponseEnvelopeSuccess bool

const (
	ResourceGroupNewResponseEnvelopeSuccessTrue ResourceGroupNewResponseEnvelopeSuccess = true
)

func (r ResourceGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ResourceGroupUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the resource group
	Name param.Field[string] `json:"name"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[ResourceGroupUpdateParamsScope] `json:"scope"`
}

func (r ResourceGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope is a combination of scope objects which provides additional context.
type ResourceGroupUpdateParamsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]ResourceGroupUpdateParamsScopeObject] `json:"objects,required"`
}

func (r ResourceGroupUpdateParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type ResourceGroupUpdateParamsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r ResourceGroupUpdateParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceGroupUpdateResponseEnvelope struct {
	Errors   []ResourceGroupUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ResourceGroupUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ResourceGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	// A group of scoped resources.
	Result ResourceGroupUpdateResponse             `json:"result"`
	JSON   resourceGroupUpdateResponseEnvelopeJSON `json:"-"`
}

// resourceGroupUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourceGroupUpdateResponseEnvelope]
type resourceGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ResourceGroupUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceGroupUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceGroupUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ResourceGroupUpdateResponseEnvelopeErrors]
type resourceGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    resourceGroupUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceGroupUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ResourceGroupUpdateResponseEnvelopeErrorsSource]
type resourceGroupUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ResourceGroupUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceGroupUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceGroupUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ResourceGroupUpdateResponseEnvelopeMessages]
type resourceGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    resourceGroupUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceGroupUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ResourceGroupUpdateResponseEnvelopeMessagesSource]
type resourceGroupUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResourceGroupUpdateResponseEnvelopeSuccess bool

const (
	ResourceGroupUpdateResponseEnvelopeSuccessTrue ResourceGroupUpdateResponseEnvelopeSuccess = true
)

func (r ResourceGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ResourceGroupListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// ID of the resource group to be fetched.
	ID param.Field[string] `query:"id"`
	// Name of the resource group to be fetched.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [ResourceGroupListParams]'s query parameters as
// `url.Values`.
func (r ResourceGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceGroupDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceGroupDeleteResponseEnvelope struct {
	Errors   []ResourceGroupDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ResourceGroupDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ResourceGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ResourceGroupDeleteResponse                `json:"result,nullable"`
	JSON    resourceGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// resourceGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourceGroupDeleteResponseEnvelope]
type resourceGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupDeleteResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ResourceGroupDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceGroupDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ResourceGroupDeleteResponseEnvelopeErrors]
type resourceGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    resourceGroupDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceGroupDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ResourceGroupDeleteResponseEnvelopeErrorsSource]
type resourceGroupDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupDeleteResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ResourceGroupDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceGroupDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceGroupDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ResourceGroupDeleteResponseEnvelopeMessages]
type resourceGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    resourceGroupDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceGroupDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ResourceGroupDeleteResponseEnvelopeMessagesSource]
type resourceGroupDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResourceGroupDeleteResponseEnvelopeSuccess bool

const (
	ResourceGroupDeleteResponseEnvelopeSuccessTrue ResourceGroupDeleteResponseEnvelopeSuccess = true
)

func (r ResourceGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ResourceGroupGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceGroupGetResponseEnvelope struct {
	Errors   []ResourceGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ResourceGroupGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ResourceGroupGetResponseEnvelopeSuccess `json:"success,required"`
	// A group of scoped resources.
	Result ResourceGroupGetResponse             `json:"result"`
	JSON   resourceGroupGetResponseEnvelopeJSON `json:"-"`
}

// resourceGroupGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceGroupGetResponseEnvelope]
type resourceGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ResourceGroupGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceGroupGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ResourceGroupGetResponseEnvelopeErrors]
type resourceGroupGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    resourceGroupGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceGroupGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ResourceGroupGetResponseEnvelopeErrorsSource]
type resourceGroupGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ResourceGroupGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceGroupGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ResourceGroupGetResponseEnvelopeMessages]
type resourceGroupGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    resourceGroupGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceGroupGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ResourceGroupGetResponseEnvelopeMessagesSource]
type resourceGroupGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGroupGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResourceGroupGetResponseEnvelopeSuccess bool

const (
	ResourceGroupGetResponseEnvelopeSuccessTrue ResourceGroupGetResponseEnvelopeSuccess = true
)

func (r ResourceGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
