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
)

// IAMPermissionGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIAMPermissionGroupService] method instead.
type IAMPermissionGroupService struct {
	Options []option.RequestOption
}

// NewIAMPermissionGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIAMPermissionGroupService(opts ...option.RequestOption) (r *IAMPermissionGroupService) {
	r = &IAMPermissionGroupService{}
	r.Options = opts
	return
}

// List all the permissions groups for an account.
func (r *IAMPermissionGroupService) List(ctx context.Context, params IAMPermissionGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[IAMPermissionGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/iam/permission_groups", params.AccountID)
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

// List all the permissions groups for an account.
func (r *IAMPermissionGroupService) ListAutoPaging(ctx context.Context, params IAMPermissionGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[IAMPermissionGroupListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Get information about a specific permission group in an account.
func (r *IAMPermissionGroupService) Get(ctx context.Context, permissionGroupID string, query IAMPermissionGroupGetParams, opts ...option.RequestOption) (res *IAMPermissionGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if permissionGroupID == "" {
		err = errors.New("missing required permission_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/permission_groups/%s", query.AccountID, permissionGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IAMPermissionGroupListResponse = interface{}

// A named group of permissions that map to a group of operations against
// resources.
type IAMPermissionGroupGetResponse struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta interface{} `json:"meta"`
	// Name of the group.
	Name string                            `json:"name"`
	JSON iamPermissionGroupGetResponseJSON `json:"-"`
}

// iamPermissionGroupGetResponseJSON contains the JSON metadata for the struct
// [IAMPermissionGroupGetResponse]
type iamPermissionGroupGetResponseJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IAMPermissionGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type IAMPermissionGroupListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// ID of the permission group to be fetched.
	ID param.Field[string] `query:"id"`
	// Label of the permission group to be fetched.
	Label param.Field[string] `query:"label"`
	// Name of the permission group to be fetched.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [IAMPermissionGroupListParams]'s query parameters as
// `url.Values`.
func (r IAMPermissionGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IAMPermissionGroupGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
