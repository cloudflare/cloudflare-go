// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserTokenService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewUserTokenService] method instead.
type UserTokenService struct {
	Options          []option.RequestOption
	PermissionGroups *UserTokenPermissionGroupService
	Verifies         *UserTokenVerifyService
	Values           *UserTokenValueService
}

// NewUserTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenService(opts ...option.RequestOption) (r *UserTokenService) {
	r = &UserTokenService{}
	r.Options = opts
	r.PermissionGroups = NewUserTokenPermissionGroupService(opts...)
	r.Verifies = NewUserTokenVerifyService(opts...)
	r.Values = NewUserTokenValueService(opts...)
	return
}

// Get information about a specific token.
func (r *UserTokenService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing token.
func (r *UserTokenService) Update(ctx context.Context, identifier interface{}, body UserTokenUpdateParams, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Destroy a token.
func (r *UserTokenService) Delete(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new access token.
func (r *UserTokenService) UserAPITokensNewToken(ctx context.Context, body UserTokenUserAPITokensNewTokenParams, opts ...option.RequestOption) (res *ResponseCreate, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all access tokens you created.
func (r *UserTokenService) UserAPITokensListTokens(ctx context.Context, query UserTokenUserAPITokensListTokensParams, opts ...option.RequestOption) (res *ResponseCollectionYgt6DzoC, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ResponseCreate struct {
	Errors   []ResponseCreateError   `json:"errors"`
	Messages []ResponseCreateMessage `json:"messages"`
	Result   ResponseCreateResult    `json:"result"`
	// Whether the API call was successful
	Success ResponseCreateSuccess `json:"success"`
	JSON    responseCreateJSON    `json:"-"`
}

// responseCreateJSON contains the JSON metadata for the struct [ResponseCreate]
type responseCreateJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCreate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCreateError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    responseCreateErrorJSON `json:"-"`
}

// responseCreateErrorJSON contains the JSON metadata for the struct
// [ResponseCreateError]
type responseCreateErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCreateError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCreateMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    responseCreateMessageJSON `json:"-"`
}

// responseCreateMessageJSON contains the JSON metadata for the struct
// [ResponseCreateMessage]
type responseCreateMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCreateMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCreateResult struct {
	// The token value.
	Value string                   `json:"value"`
	JSON  responseCreateResultJSON `json:"-"`
}

// responseCreateResultJSON contains the JSON metadata for the struct
// [ResponseCreateResult]
type responseCreateResultJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCreateResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCreateSuccess bool

const (
	ResponseCreateSuccessTrue ResponseCreateSuccess = true
)

type UserTokenUpdateParams struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies param.Field[[]UserTokenUpdateParamsPolicy] `json:"policies,required"`
	// Status of the token.
	Status    param.Field[UserTokenUpdateParamsStatus]    `json:"status,required"`
	Condition param.Field[UserTokenUpdateParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r UserTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenUpdateParamsPolicy struct {
	// Allow or deny operations against the resources.
	Effect param.Field[UserTokenUpdateParamsPoliciesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]UserTokenUpdateParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[interface{}] `json:"resources,required"`
}

func (r UserTokenUpdateParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type UserTokenUpdateParamsPoliciesEffect string

const (
	UserTokenUpdateParamsPoliciesEffectAllow UserTokenUpdateParamsPoliciesEffect = "allow"
	UserTokenUpdateParamsPoliciesEffectDeny  UserTokenUpdateParamsPoliciesEffect = "deny"
)

// A named group of permissions that map to a group of operations against
// resources.
type UserTokenUpdateParamsPoliciesPermissionGroup struct {
}

func (r UserTokenUpdateParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the token.
type UserTokenUpdateParamsStatus string

const (
	UserTokenUpdateParamsStatusActive   UserTokenUpdateParamsStatus = "active"
	UserTokenUpdateParamsStatusDisabled UserTokenUpdateParamsStatus = "disabled"
	UserTokenUpdateParamsStatusExpired  UserTokenUpdateParamsStatus = "expired"
)

type UserTokenUpdateParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[UserTokenUpdateParamsConditionRequestIP] `json:"request.ip"`
}

func (r UserTokenUpdateParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type UserTokenUpdateParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r UserTokenUpdateParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenUserAPITokensNewTokenParams struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]UserTokenUserAPITokensNewTokenParamsPolicy]  `json:"policies,required"`
	Condition param.Field[UserTokenUserAPITokensNewTokenParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r UserTokenUserAPITokensNewTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenUserAPITokensNewTokenParamsPolicy struct {
	// Allow or deny operations against the resources.
	Effect param.Field[UserTokenUserAPITokensNewTokenParamsPoliciesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]UserTokenUserAPITokensNewTokenParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[interface{}] `json:"resources,required"`
}

func (r UserTokenUserAPITokensNewTokenParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type UserTokenUserAPITokensNewTokenParamsPoliciesEffect string

const (
	UserTokenUserAPITokensNewTokenParamsPoliciesEffectAllow UserTokenUserAPITokensNewTokenParamsPoliciesEffect = "allow"
	UserTokenUserAPITokensNewTokenParamsPoliciesEffectDeny  UserTokenUserAPITokensNewTokenParamsPoliciesEffect = "deny"
)

// A named group of permissions that map to a group of operations against
// resources.
type UserTokenUserAPITokensNewTokenParamsPoliciesPermissionGroup struct {
}

func (r UserTokenUserAPITokensNewTokenParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenUserAPITokensNewTokenParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[UserTokenUserAPITokensNewTokenParamsConditionRequestIP] `json:"request.ip"`
}

func (r UserTokenUserAPITokensNewTokenParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type UserTokenUserAPITokensNewTokenParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r UserTokenUserAPITokensNewTokenParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenUserAPITokensListTokensParams struct {
	// Direction to order results.
	Direction param.Field[UserTokenUserAPITokensListTokensParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserTokenUserAPITokensListTokensParams]'s query parameters
// as `url.Values`.
func (r UserTokenUserAPITokensListTokensParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type UserTokenUserAPITokensListTokensParamsDirection string

const (
	UserTokenUserAPITokensListTokensParamsDirectionAsc  UserTokenUserAPITokensListTokensParamsDirection = "asc"
	UserTokenUserAPITokensListTokensParamsDirectionDesc UserTokenUserAPITokensListTokensParamsDirection = "desc"
)
