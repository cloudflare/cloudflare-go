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
func (r *OrganizationRoleService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationRoleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/roles/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all available roles for an organization.
func (r *OrganizationRoleService) OrganizationRolesListRoles(ctx context.Context, organizationIdentifier string, opts ...option.RequestOption) (res *OrganizationRoleOrganizationRolesListRolesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/roles", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrganizationRoleGetResponse struct {
	Errors   []OrganizationRoleGetResponseError   `json:"errors"`
	Messages []OrganizationRoleGetResponseMessage `json:"messages"`
	Result   interface{}                          `json:"result"`
	// Whether the API call was successful
	Success OrganizationRoleGetResponseSuccess `json:"success"`
	JSON    organizationRoleGetResponseJSON    `json:"-"`
}

// organizationRoleGetResponseJSON contains the JSON metadata for the struct
// [OrganizationRoleGetResponse]
type organizationRoleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRoleGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    organizationRoleGetResponseErrorJSON `json:"-"`
}

// organizationRoleGetResponseErrorJSON contains the JSON metadata for the struct
// [OrganizationRoleGetResponseError]
type organizationRoleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRoleGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    organizationRoleGetResponseMessageJSON `json:"-"`
}

// organizationRoleGetResponseMessageJSON contains the JSON metadata for the struct
// [OrganizationRoleGetResponseMessage]
type organizationRoleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRoleGetResponseSuccess bool

const (
	OrganizationRoleGetResponseSuccessTrue OrganizationRoleGetResponseSuccess = true
)

type OrganizationRoleOrganizationRolesListRolesResponse struct {
	Errors     []OrganizationRoleOrganizationRolesListRolesResponseError    `json:"errors"`
	Messages   []OrganizationRoleOrganizationRolesListRolesResponseMessage  `json:"messages"`
	Result     []OrganizationRoleOrganizationRolesListRolesResponseResult   `json:"result"`
	ResultInfo OrganizationRoleOrganizationRolesListRolesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OrganizationRoleOrganizationRolesListRolesResponseSuccess `json:"success"`
	JSON    organizationRoleOrganizationRolesListRolesResponseJSON    `json:"-"`
}

// organizationRoleOrganizationRolesListRolesResponseJSON contains the JSON
// metadata for the struct [OrganizationRoleOrganizationRolesListRolesResponse]
type organizationRoleOrganizationRolesListRolesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleOrganizationRolesListRolesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRoleOrganizationRolesListRolesResponseError struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    organizationRoleOrganizationRolesListRolesResponseErrorJSON `json:"-"`
}

// organizationRoleOrganizationRolesListRolesResponseErrorJSON contains the JSON
// metadata for the struct
// [OrganizationRoleOrganizationRolesListRolesResponseError]
type organizationRoleOrganizationRolesListRolesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleOrganizationRolesListRolesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRoleOrganizationRolesListRolesResponseMessage struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    organizationRoleOrganizationRolesListRolesResponseMessageJSON `json:"-"`
}

// organizationRoleOrganizationRolesListRolesResponseMessageJSON contains the JSON
// metadata for the struct
// [OrganizationRoleOrganizationRolesListRolesResponseMessage]
type organizationRoleOrganizationRolesListRolesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleOrganizationRolesListRolesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRoleOrganizationRolesListRolesResponseResult struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                                     `json:"permissions,required"`
	JSON        organizationRoleOrganizationRolesListRolesResponseResultJSON `json:"-"`
}

// organizationRoleOrganizationRolesListRolesResponseResultJSON contains the JSON
// metadata for the struct
// [OrganizationRoleOrganizationRolesListRolesResponseResult]
type organizationRoleOrganizationRolesListRolesResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleOrganizationRolesListRolesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRoleOrganizationRolesListRolesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       organizationRoleOrganizationRolesListRolesResponseResultInfoJSON `json:"-"`
}

// organizationRoleOrganizationRolesListRolesResponseResultInfoJSON contains the
// JSON metadata for the struct
// [OrganizationRoleOrganizationRolesListRolesResponseResultInfo]
type organizationRoleOrganizationRolesListRolesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRoleOrganizationRolesListRolesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRoleOrganizationRolesListRolesResponseSuccess bool

const (
	OrganizationRoleOrganizationRolesListRolesResponseSuccessTrue OrganizationRoleOrganizationRolesListRolesResponseSuccess = true
)
