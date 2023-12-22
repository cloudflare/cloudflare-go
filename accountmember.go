// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMemberService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMemberService] method
// instead.
type AccountMemberService struct {
	Options []option.RequestOption
}

// NewAccountMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMemberService(opts ...option.RequestOption) (r *AccountMemberService) {
	r = &AccountMemberService{}
	r.Options = opts
	return
}

// Get information about a specific member of an account.
func (r *AccountMemberService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify an account member.
func (r *AccountMemberService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountMemberUpdateParams, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove a member from an account.
func (r *AccountMemberService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a user to the list of members for this account.
func (r *AccountMemberService) AccountMembersAddMember(ctx context.Context, accountIdentifier interface{}, body AccountMemberAccountMembersAddMemberParams, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all members of an account.
func (r *AccountMemberService) AccountMembersListMembers(ctx context.Context, accountIdentifier interface{}, query AccountMemberAccountMembersListMembersParams, opts ...option.RequestOption) (res *ResponseCollectionYgt6DzoC, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountMemberUpdateParams struct {
	// Roles assigned to this member.
	Roles param.Field[[]AccountMemberUpdateParamsRole] `json:"roles,required"`
}

func (r AccountMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRole struct {
	Permissions param.Field[AccountMemberUpdateParamsRolesPermissions] `json:"permissions,required"`
}

func (r AccountMemberUpdateParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissions struct {
	Analytics    param.Field[AccountMemberUpdateParamsRolesPermissionsAnalytics]    `json:"analytics"`
	Billing      param.Field[AccountMemberUpdateParamsRolesPermissionsBilling]      `json:"billing"`
	CachePurge   param.Field[AccountMemberUpdateParamsRolesPermissionsCachePurge]   `json:"cache_purge"`
	DNS          param.Field[AccountMemberUpdateParamsRolesPermissionsDNS]          `json:"dns"`
	DNSRecords   param.Field[AccountMemberUpdateParamsRolesPermissionsDNSRecords]   `json:"dns_records"`
	Lb           param.Field[AccountMemberUpdateParamsRolesPermissionsLb]           `json:"lb"`
	Logs         param.Field[AccountMemberUpdateParamsRolesPermissionsLogs]         `json:"logs"`
	Organization param.Field[AccountMemberUpdateParamsRolesPermissionsOrganization] `json:"organization"`
	Ssl          param.Field[AccountMemberUpdateParamsRolesPermissionsSsl]          `json:"ssl"`
	Waf          param.Field[AccountMemberUpdateParamsRolesPermissionsWaf]          `json:"waf"`
	ZoneSettings param.Field[AccountMemberUpdateParamsRolesPermissionsZoneSettings] `json:"zone_settings"`
	Zones        param.Field[AccountMemberUpdateParamsRolesPermissionsZones]        `json:"zones"`
}

func (r AccountMemberUpdateParamsRolesPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsAnalytics struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsAnalytics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsBilling struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsCachePurge struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsCachePurge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsDNS struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsDNSRecords struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsDNSRecords) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsLb struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsLb) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsLogs struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsOrganization struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsSsl struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsSsl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsWaf struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsWaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsZoneSettings struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsZoneSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsZones struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsZones) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberAccountMembersAddMemberParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]                                         `json:"roles,required"`
	Status param.Field[AccountMemberAccountMembersAddMemberParamsStatus] `json:"status"`
}

func (r AccountMemberAccountMembersAddMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberAccountMembersAddMemberParamsStatus string

const (
	AccountMemberAccountMembersAddMemberParamsStatusAccepted AccountMemberAccountMembersAddMemberParamsStatus = "accepted"
	AccountMemberAccountMembersAddMemberParamsStatusPending  AccountMemberAccountMembersAddMemberParamsStatus = "pending"
)

type AccountMemberAccountMembersListMembersParams struct {
	// Direction to order results.
	Direction param.Field[AccountMemberAccountMembersListMembersParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountMemberAccountMembersListMembersParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[AccountMemberAccountMembersListMembersParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountMemberAccountMembersListMembersParams]'s query
// parameters as `url.Values`.
func (r AccountMemberAccountMembersListMembersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountMemberAccountMembersListMembersParamsDirection string

const (
	AccountMemberAccountMembersListMembersParamsDirectionAsc  AccountMemberAccountMembersListMembersParamsDirection = "asc"
	AccountMemberAccountMembersListMembersParamsDirectionDesc AccountMemberAccountMembersListMembersParamsDirection = "desc"
)

// Field to order results by.
type AccountMemberAccountMembersListMembersParamsOrder string

const (
	AccountMemberAccountMembersListMembersParamsOrderUserFirstName AccountMemberAccountMembersListMembersParamsOrder = "user.first_name"
	AccountMemberAccountMembersListMembersParamsOrderUserLastName  AccountMemberAccountMembersListMembersParamsOrder = "user.last_name"
	AccountMemberAccountMembersListMembersParamsOrderUserEmail     AccountMemberAccountMembersListMembersParamsOrder = "user.email"
	AccountMemberAccountMembersListMembersParamsOrderStatus        AccountMemberAccountMembersListMembersParamsOrder = "status"
)

// A member's status in the account.
type AccountMemberAccountMembersListMembersParamsStatus string

const (
	AccountMemberAccountMembersListMembersParamsStatusAccepted AccountMemberAccountMembersListMembersParamsStatus = "accepted"
	AccountMemberAccountMembersListMembersParamsStatusPending  AccountMemberAccountMembersListMembersParamsStatus = "pending"
	AccountMemberAccountMembersListMembersParamsStatusRejected AccountMemberAccountMembersListMembersParamsStatus = "rejected"
)
