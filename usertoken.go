// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Update an existing token.
func (r *UserTokenService) Update(ctx context.Context, tokenID interface{}, body UserTokenUpdateParams, opts ...option.RequestOption) (res *UserTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenUpdateResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Destroy a token.
func (r *UserTokenService) Delete(ctx context.Context, tokenID interface{}, opts ...option.RequestOption) (res *UserTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenDeleteResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific token.
func (r *UserTokenService) Get(ctx context.Context, tokenID interface{}, opts ...option.RequestOption) (res *UserTokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenGetResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new access token.
func (r *UserTokenService) UserAPITokensNewToken(ctx context.Context, body UserTokenUserAPITokensNewTokenParams, opts ...option.RequestOption) (res *UserTokenUserAPITokensNewTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenUserAPITokensNewTokenResponseEnvelope
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all access tokens you created.
func (r *UserTokenService) UserAPITokensListTokens(ctx context.Context, query UserTokenUserAPITokensListTokensParams, opts ...option.RequestOption) (res *[]UserTokenUserAPITokensListTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenUserAPITokensListTokensResponseEnvelope
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserTokenUpdateResponseUnknown] or [shared.UnionString].
type UserTokenUpdateResponse interface {
	ImplementsUserTokenUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserTokenUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserTokenDeleteResponse struct {
	// Identifier
	ID   string                      `json:"id,required"`
	JSON userTokenDeleteResponseJSON `json:"-"`
}

// userTokenDeleteResponseJSON contains the JSON metadata for the struct
// [UserTokenDeleteResponse]
type userTokenDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [UserTokenGetResponseUnknown] or [shared.UnionString].
type UserTokenGetResponse interface {
	ImplementsUserTokenGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserTokenGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserTokenUserAPITokensNewTokenResponse struct {
	// The token value.
	Value string                                     `json:"value"`
	JSON  userTokenUserAPITokensNewTokenResponseJSON `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseJSON contains the JSON metadata for the
// struct [UserTokenUserAPITokensNewTokenResponse]
type userTokenUserAPITokensNewTokenResponseJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	RequestIP param.Field[UserTokenUpdateParamsConditionRequestIP] `json:"request_ip"`
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

type UserTokenUpdateResponseEnvelope struct {
	Errors   []UserTokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// userTokenUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserTokenUpdateResponseEnvelope]
type userTokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    userTokenUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserTokenUpdateResponseEnvelopeErrors]
type userTokenUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    userTokenUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserTokenUpdateResponseEnvelopeMessages]
type userTokenUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenUpdateResponseEnvelopeSuccess bool

const (
	UserTokenUpdateResponseEnvelopeSuccessTrue UserTokenUpdateResponseEnvelopeSuccess = true
)

type UserTokenDeleteResponseEnvelope struct {
	Errors   []UserTokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserTokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// userTokenDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserTokenDeleteResponseEnvelope]
type userTokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    userTokenDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserTokenDeleteResponseEnvelopeErrors]
type userTokenDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    userTokenDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserTokenDeleteResponseEnvelopeMessages]
type userTokenDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenDeleteResponseEnvelopeSuccess bool

const (
	UserTokenDeleteResponseEnvelopeSuccessTrue UserTokenDeleteResponseEnvelopeSuccess = true
)

type UserTokenGetResponseEnvelope struct {
	Errors   []UserTokenGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenGetResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenGetResponseEnvelopeJSON    `json:"-"`
}

// userTokenGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserTokenGetResponseEnvelope]
type userTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    userTokenGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserTokenGetResponseEnvelopeErrors]
type userTokenGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    userTokenGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserTokenGetResponseEnvelopeMessages]
type userTokenGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenGetResponseEnvelopeSuccess bool

const (
	UserTokenGetResponseEnvelopeSuccessTrue UserTokenGetResponseEnvelopeSuccess = true
)

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
	RequestIP param.Field[UserTokenUserAPITokensNewTokenParamsConditionRequestIP] `json:"request_ip"`
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

type UserTokenUserAPITokensNewTokenResponseEnvelope struct {
	Errors   []UserTokenUserAPITokensNewTokenResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenUserAPITokensNewTokenResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenUserAPITokensNewTokenResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenUserAPITokensNewTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenUserAPITokensNewTokenResponseEnvelopeJSON    `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserTokenUserAPITokensNewTokenResponseEnvelope]
type userTokenUserAPITokensNewTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensNewTokenResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userTokenUserAPITokensNewTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserTokenUserAPITokensNewTokenResponseEnvelopeErrors]
type userTokenUserAPITokensNewTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensNewTokenResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userTokenUserAPITokensNewTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenUserAPITokensNewTokenResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserTokenUserAPITokensNewTokenResponseEnvelopeMessages]
type userTokenUserAPITokensNewTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensNewTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenUserAPITokensNewTokenResponseEnvelopeSuccess bool

const (
	UserTokenUserAPITokensNewTokenResponseEnvelopeSuccessTrue UserTokenUserAPITokensNewTokenResponseEnvelopeSuccess = true
)

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

type UserTokenUserAPITokensListTokensResponseEnvelope struct {
	Errors   []UserTokenUserAPITokensListTokensResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenUserAPITokensListTokensResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserTokenUserAPITokensListTokensResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserTokenUserAPITokensListTokensResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserTokenUserAPITokensListTokensResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userTokenUserAPITokensListTokensResponseEnvelopeJSON       `json:"-"`
}

// userTokenUserAPITokensListTokensResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserTokenUserAPITokensListTokensResponseEnvelope]
type userTokenUserAPITokensListTokensResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensListTokensResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensListTokensResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userTokenUserAPITokensListTokensResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenUserAPITokensListTokensResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserTokenUserAPITokensListTokensResponseEnvelopeErrors]
type userTokenUserAPITokensListTokensResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensListTokensResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenUserAPITokensListTokensResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    userTokenUserAPITokensListTokensResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenUserAPITokensListTokensResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [UserTokenUserAPITokensListTokensResponseEnvelopeMessages]
type userTokenUserAPITokensListTokensResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensListTokensResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenUserAPITokensListTokensResponseEnvelopeSuccess bool

const (
	UserTokenUserAPITokensListTokensResponseEnvelopeSuccessTrue UserTokenUserAPITokensListTokensResponseEnvelopeSuccess = true
)

type UserTokenUserAPITokensListTokensResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       userTokenUserAPITokensListTokensResponseEnvelopeResultInfoJSON `json:"-"`
}

// userTokenUserAPITokensListTokensResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [UserTokenUserAPITokensListTokensResponseEnvelopeResultInfo]
type userTokenUserAPITokensListTokensResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenUserAPITokensListTokensResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
