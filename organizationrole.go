// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationRoleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationRoleService] method
// instead.
type OrganizationRoleService struct {
	Options []option.RequestOption
}

// NewOrganizationRoleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationRoleService(opts ...option.RequestOption) (r *OrganizationRoleService) {
	r = &OrganizationRoleService{}
	r.Options = opts
	return
}

// Get information about a specific role for an organization.
func (r *OrganizationRoleService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *SingleRoleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/roles/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all available roles for an organization.
func (r *OrganizationRoleService) OrganizationRolesListRoles(ctx context.Context, organizationIdentifier string, opts ...option.RequestOption) (res *CollectionRoleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/roles", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionRoleResponse struct {
	Errors     []CollectionRoleResponseError    `json:"errors"`
	Messages   []CollectionRoleResponseMessage  `json:"messages"`
	Result     []CollectionRoleResponseResult   `json:"result"`
	ResultInfo CollectionRoleResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CollectionRoleResponseSuccess `json:"success"`
	JSON    collectionRoleResponseJSON    `json:"-"`
}

// collectionRoleResponseJSON contains the JSON metadata for the struct
// [CollectionRoleResponse]
type collectionRoleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionRoleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionRoleResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    collectionRoleResponseErrorJSON `json:"-"`
}

// collectionRoleResponseErrorJSON contains the JSON metadata for the struct
// [CollectionRoleResponseError]
type collectionRoleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionRoleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionRoleResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    collectionRoleResponseMessageJSON `json:"-"`
}

// collectionRoleResponseMessageJSON contains the JSON metadata for the struct
// [CollectionRoleResponseMessage]
type collectionRoleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionRoleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionRoleResponseResult struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                         `json:"permissions,required"`
	JSON        collectionRoleResponseResultJSON `json:"-"`
}

// collectionRoleResponseResultJSON contains the JSON metadata for the struct
// [CollectionRoleResponseResult]
type collectionRoleResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionRoleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionRoleResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       collectionRoleResponseResultInfoJSON `json:"-"`
}

// collectionRoleResponseResultInfoJSON contains the JSON metadata for the struct
// [CollectionRoleResponseResultInfo]
type collectionRoleResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionRoleResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CollectionRoleResponseSuccess bool

const (
	CollectionRoleResponseSuccessTrue CollectionRoleResponseSuccess = true
)

type SingleRoleResponse struct {
	Errors   []SingleRoleResponseError   `json:"errors"`
	Messages []SingleRoleResponseMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success SingleRoleResponseSuccess `json:"success"`
	JSON    singleRoleResponseJSON    `json:"-"`
}

// singleRoleResponseJSON contains the JSON metadata for the struct
// [SingleRoleResponse]
type singleRoleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleRoleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleRoleResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    singleRoleResponseErrorJSON `json:"-"`
}

// singleRoleResponseErrorJSON contains the JSON metadata for the struct
// [SingleRoleResponseError]
type singleRoleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleRoleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleRoleResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    singleRoleResponseMessageJSON `json:"-"`
}

// singleRoleResponseMessageJSON contains the JSON metadata for the struct
// [SingleRoleResponseMessage]
type singleRoleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleRoleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleRoleResponseSuccess bool

const (
	SingleRoleResponseSuccessTrue SingleRoleResponseSuccess = true
)
