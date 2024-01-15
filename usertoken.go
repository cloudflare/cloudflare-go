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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *UserTokenService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *UserTokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing token.
func (r *UserTokenService) Update(ctx context.Context, identifier interface{}, body UserTokenUpdateParams, opts ...option.RequestOption) (res *UserTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Destroy a token.
func (r *UserTokenService) Delete(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *UserTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new access token.
func (r *UserTokenService) UserAPITokensNewToken(ctx context.Context, body UserTokenUserAPITokensNewTokenParams, opts ...option.RequestOption) (res *UserTokenUserAPITokensNewTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all access tokens you created.
func (r *UserTokenService) UserAPITokensListTokens(ctx context.Context, query UserTokenUserAPITokensListTokensParams, opts ...option.RequestOption) (res *shared.Page[UserTokenUserAPITokensListTokensResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/tokens"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

type UserTokenGetResponse struct {
	Errors   []UserTokenGetResponseError   `json:"errors"`
	Messages []UserTokenGetResponseMessage `json:"messages"`
	Result   interface{}                   `json:"result"`
	// Whether the API call was successful
	Success UserTokenGetResponseSuccess `json:"success"`
	JSON    userTokenGetResponseJSON    `json:"-"`
}

// userTokenGetResponseJSON contains the JSON metadata for the struct
// [UserTokenGetResponse]
type userTokenGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    userTokenGetResponseErrorJSON `json:"-"`
}

// userTokenGetResponseErrorJSON contains the JSON metadata for the struct
// [UserTokenGetResponseError]
type userTokenGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    userTokenGetResponseMessageJSON `json:"-"`
}

// userTokenGetResponseMessageJSON contains the JSON metadata for the struct
// [UserTokenGetResponseMessage]
type userTokenGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenGetResponseSuccess bool

const (
	UserTokenGetResponseSuccessTrue UserTokenGetResponseSuccess = true
)

type UserTokenUpdateResponse struct {
	Errors   []UserTokenUpdateResponseError   `json:"errors"`
	Messages []UserTokenUpdateResponseMessage `json:"messages"`
	Result   interface{}                      `json:"result"`
	// Whether the API call was successful
	Success UserTokenUpdateResponseSuccess `json:"success"`
	JSON    userTokenUpdateResponseJSON    `json:"-"`
}

// userTokenUpdateResponseJSON contains the JSON metadata for the struct
// [UserTokenUpdateResponse]
type userTokenUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUpdateResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    userTokenUpdateResponseErrorJSON `json:"-"`
}

// userTokenUpdateResponseErrorJSON contains the JSON metadata for the struct
// [UserTokenUpdateResponseError]
type userTokenUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUpdateResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    userTokenUpdateResponseMessageJSON `json:"-"`
}

// userTokenUpdateResponseMessageJSON contains the JSON metadata for the struct
// [UserTokenUpdateResponseMessage]
type userTokenUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenUpdateResponseSuccess bool

const (
	UserTokenUpdateResponseSuccessTrue UserTokenUpdateResponseSuccess = true
)

type UserTokenDeleteResponse struct {
	Errors   []UserTokenDeleteResponseError   `json:"errors"`
	Messages []UserTokenDeleteResponseMessage `json:"messages"`
	Result   UserTokenDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success UserTokenDeleteResponseSuccess `json:"success"`
	JSON    userTokenDeleteResponseJSON    `json:"-"`
}

// userTokenDeleteResponseJSON contains the JSON metadata for the struct
// [UserTokenDeleteResponse]
type userTokenDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenDeleteResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    userTokenDeleteResponseErrorJSON `json:"-"`
}

// userTokenDeleteResponseErrorJSON contains the JSON metadata for the struct
// [UserTokenDeleteResponseError]
type userTokenDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenDeleteResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    userTokenDeleteResponseMessageJSON `json:"-"`
}

// userTokenDeleteResponseMessageJSON contains the JSON metadata for the struct
// [UserTokenDeleteResponseMessage]
type userTokenDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenDeleteResponseResult struct {
	// Identifier
	ID   string                            `json:"id,required"`
	JSON userTokenDeleteResponseResultJSON `json:"-"`
}

// userTokenDeleteResponseResultJSON contains the JSON metadata for the struct
// [UserTokenDeleteResponseResult]
type userTokenDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenDeleteResponseSuccess bool

const (
	UserTokenDeleteResponseSuccessTrue UserTokenDeleteResponseSuccess = true
)

type UserTokenUserAPITokensNewTokenResponse struct {
	Errors   []UserTokenUserAPITokensNewTokenResponseError   `json:"errors"`
	Messages []UserTokenUserAPITokensNewTokenResponseMessage `json:"messages"`
	Result   UserTokenUserAPITokensNewTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserTokenUserAPITokensNewTokenResponseSuccess `json:"success"`
	JSON    userTokenUserAPITokensNewTokenResponseJSON    `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseJSON contains the JSON metadata for the
// struct [UserTokenUserAPITokensNewTokenResponse]
type userTokenUserAPITokensNewTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensNewTokenResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    userTokenUserAPITokensNewTokenResponseErrorJSON `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseErrorJSON contains the JSON metadata for
// the struct [UserTokenUserAPITokensNewTokenResponseError]
type userTokenUserAPITokensNewTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensNewTokenResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userTokenUserAPITokensNewTokenResponseMessageJSON `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseMessageJSON contains the JSON metadata for
// the struct [UserTokenUserAPITokensNewTokenResponseMessage]
type userTokenUserAPITokensNewTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensNewTokenResponseResult struct {
	// The token value.
	Value string                                           `json:"value"`
	JSON  userTokenUserAPITokensNewTokenResponseResultJSON `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseResultJSON contains the JSON metadata for
// the struct [UserTokenUserAPITokensNewTokenResponseResult]
type userTokenUserAPITokensNewTokenResponseResultJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenUserAPITokensNewTokenResponseSuccess bool

const (
	UserTokenUserAPITokensNewTokenResponseSuccessTrue UserTokenUserAPITokensNewTokenResponseSuccess = true
)

type UserTokenUserAPITokensListTokensResponse = interface{}

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
