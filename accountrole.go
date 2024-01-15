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

// AccountRoleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRoleService] method
// instead.
type AccountRoleService struct {
	Options []option.RequestOption
}

// NewAccountRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRoleService(opts ...option.RequestOption) (r *AccountRoleService) {
	r = &AccountRoleService{}
	r.Options = opts
	return
}

// Get information about a specific role for an account.
func (r *AccountRoleService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountRoleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/roles/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all available roles for an account.
func (r *AccountRoleService) AccountRolesListRoles(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountRoleAccountRolesListRolesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/roles", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRoleGetResponse struct {
	Errors   []AccountRoleGetResponseError   `json:"errors"`
	Messages []AccountRoleGetResponseMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success AccountRoleGetResponseSuccess `json:"success"`
	JSON    accountRoleGetResponseJSON    `json:"-"`
}

// accountRoleGetResponseJSON contains the JSON metadata for the struct
// [AccountRoleGetResponse]
type accountRoleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    accountRoleGetResponseErrorJSON `json:"-"`
}

// accountRoleGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountRoleGetResponseError]
type accountRoleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accountRoleGetResponseMessageJSON `json:"-"`
}

// accountRoleGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountRoleGetResponseMessage]
type accountRoleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRoleGetResponseSuccess bool

const (
	AccountRoleGetResponseSuccessTrue AccountRoleGetResponseSuccess = true
)

type AccountRoleAccountRolesListRolesResponse struct {
	Errors     []AccountRoleAccountRolesListRolesResponseError    `json:"errors"`
	Messages   []AccountRoleAccountRolesListRolesResponseMessage  `json:"messages"`
	Result     []AccountRoleAccountRolesListRolesResponseResult   `json:"result"`
	ResultInfo AccountRoleAccountRolesListRolesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountRoleAccountRolesListRolesResponseSuccess `json:"success"`
	JSON    accountRoleAccountRolesListRolesResponseJSON    `json:"-"`
}

// accountRoleAccountRolesListRolesResponseJSON contains the JSON metadata for the
// struct [AccountRoleAccountRolesListRolesResponse]
type accountRoleAccountRolesListRolesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleAccountRolesListRolesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleAccountRolesListRolesResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountRoleAccountRolesListRolesResponseErrorJSON `json:"-"`
}

// accountRoleAccountRolesListRolesResponseErrorJSON contains the JSON metadata for
// the struct [AccountRoleAccountRolesListRolesResponseError]
type accountRoleAccountRolesListRolesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleAccountRolesListRolesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleAccountRolesListRolesResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountRoleAccountRolesListRolesResponseMessageJSON `json:"-"`
}

// accountRoleAccountRolesListRolesResponseMessageJSON contains the JSON metadata
// for the struct [AccountRoleAccountRolesListRolesResponseMessage]
type accountRoleAccountRolesListRolesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleAccountRolesListRolesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleAccountRolesListRolesResponseResult struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                           `json:"permissions,required"`
	JSON        accountRoleAccountRolesListRolesResponseResultJSON `json:"-"`
}

// accountRoleAccountRolesListRolesResponseResultJSON contains the JSON metadata
// for the struct [AccountRoleAccountRolesListRolesResponseResult]
type accountRoleAccountRolesListRolesResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleAccountRolesListRolesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleAccountRolesListRolesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       accountRoleAccountRolesListRolesResponseResultInfoJSON `json:"-"`
}

// accountRoleAccountRolesListRolesResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountRoleAccountRolesListRolesResponseResultInfo]
type accountRoleAccountRolesListRolesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleAccountRolesListRolesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRoleAccountRolesListRolesResponseSuccess bool

const (
	AccountRoleAccountRolesListRolesResponseSuccessTrue AccountRoleAccountRolesListRolesResponseSuccess = true
)
