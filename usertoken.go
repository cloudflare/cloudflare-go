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
	Values           *UserTokenValueService
}

// NewUserTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenService(opts ...option.RequestOption) (r *UserTokenService) {
	r = &UserTokenService{}
	r.Options = opts
	r.PermissionGroups = NewUserTokenPermissionGroupService(opts...)
	r.Values = NewUserTokenValueService(opts...)
	return
}

// Create a new access token.
func (r *UserTokenService) New(ctx context.Context, body UserTokenNewParams, opts ...option.RequestOption) (res *UserTokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenNewResponseEnvelope
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all access tokens you created.
func (r *UserTokenService) List(ctx context.Context, query UserTokenListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[UserTokenListResponse], err error) {
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

// List all access tokens you created.
func (r *UserTokenService) ListAutoPaging(ctx context.Context, query UserTokenListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[UserTokenListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
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

// Update an existing token.
func (r *UserTokenService) Replace(ctx context.Context, tokenID interface{}, body UserTokenReplaceParams, opts ...option.RequestOption) (res *UserTokenReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenReplaceResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test whether a token works.
func (r *UserTokenService) Verify(ctx context.Context, opts ...option.RequestOption) (res *UserTokenVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenVerifyResponseEnvelope
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserTokenNewResponse struct {
	// The token value.
	Value string                   `json:"value"`
	JSON  userTokenNewResponseJSON `json:"-"`
}

// userTokenNewResponseJSON contains the JSON metadata for the struct
// [UserTokenNewResponse]
type userTokenNewResponseJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenListResponse = interface{}

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

// Union satisfied by [UserTokenReplaceResponseUnknown] or [shared.UnionString].
type UserTokenReplaceResponse interface {
	ImplementsUserTokenReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserTokenReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserTokenVerifyResponse struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status UserTokenVerifyResponseStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                   `json:"not_before" format:"date-time"`
	JSON      userTokenVerifyResponseJSON `json:"-"`
}

// userTokenVerifyResponseJSON contains the JSON metadata for the struct
// [UserTokenVerifyResponse]
type userTokenVerifyResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the token.
type UserTokenVerifyResponseStatus string

const (
	UserTokenVerifyResponseStatusActive   UserTokenVerifyResponseStatus = "active"
	UserTokenVerifyResponseStatusDisabled UserTokenVerifyResponseStatus = "disabled"
	UserTokenVerifyResponseStatusExpired  UserTokenVerifyResponseStatus = "expired"
)

type UserTokenNewParams struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]UserTokenNewParamsPolicy]  `json:"policies,required"`
	Condition param.Field[UserTokenNewParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r UserTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenNewParamsPolicy struct {
	// Allow or deny operations against the resources.
	Effect param.Field[UserTokenNewParamsPoliciesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]UserTokenNewParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[interface{}] `json:"resources,required"`
}

func (r UserTokenNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type UserTokenNewParamsPoliciesEffect string

const (
	UserTokenNewParamsPoliciesEffectAllow UserTokenNewParamsPoliciesEffect = "allow"
	UserTokenNewParamsPoliciesEffectDeny  UserTokenNewParamsPoliciesEffect = "deny"
)

// A named group of permissions that map to a group of operations against
// resources.
type UserTokenNewParamsPoliciesPermissionGroup struct {
}

func (r UserTokenNewParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenNewParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[UserTokenNewParamsConditionRequestIP] `json:"request_ip"`
}

func (r UserTokenNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type UserTokenNewParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r UserTokenNewParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenNewResponseEnvelope struct {
	Errors   []UserTokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenNewResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenNewResponseEnvelopeJSON    `json:"-"`
}

// userTokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserTokenNewResponseEnvelope]
type userTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    userTokenNewResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserTokenNewResponseEnvelopeErrors]
type userTokenNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    userTokenNewResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserTokenNewResponseEnvelopeMessages]
type userTokenNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenNewResponseEnvelopeSuccess bool

const (
	UserTokenNewResponseEnvelopeSuccessTrue UserTokenNewResponseEnvelopeSuccess = true
)

type UserTokenListParams struct {
	// Direction to order results.
	Direction param.Field[UserTokenListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserTokenListParams]'s query parameters as `url.Values`.
func (r UserTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type UserTokenListParamsDirection string

const (
	UserTokenListParamsDirectionAsc  UserTokenListParamsDirection = "asc"
	UserTokenListParamsDirectionDesc UserTokenListParamsDirection = "desc"
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

type UserTokenReplaceParams struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies param.Field[[]UserTokenReplaceParamsPolicy] `json:"policies,required"`
	// Status of the token.
	Status    param.Field[UserTokenReplaceParamsStatus]    `json:"status,required"`
	Condition param.Field[UserTokenReplaceParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r UserTokenReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenReplaceParamsPolicy struct {
	// Allow or deny operations against the resources.
	Effect param.Field[UserTokenReplaceParamsPoliciesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]UserTokenReplaceParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[interface{}] `json:"resources,required"`
}

func (r UserTokenReplaceParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type UserTokenReplaceParamsPoliciesEffect string

const (
	UserTokenReplaceParamsPoliciesEffectAllow UserTokenReplaceParamsPoliciesEffect = "allow"
	UserTokenReplaceParamsPoliciesEffectDeny  UserTokenReplaceParamsPoliciesEffect = "deny"
)

// A named group of permissions that map to a group of operations against
// resources.
type UserTokenReplaceParamsPoliciesPermissionGroup struct {
}

func (r UserTokenReplaceParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the token.
type UserTokenReplaceParamsStatus string

const (
	UserTokenReplaceParamsStatusActive   UserTokenReplaceParamsStatus = "active"
	UserTokenReplaceParamsStatusDisabled UserTokenReplaceParamsStatus = "disabled"
	UserTokenReplaceParamsStatusExpired  UserTokenReplaceParamsStatus = "expired"
)

type UserTokenReplaceParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[UserTokenReplaceParamsConditionRequestIP] `json:"request_ip"`
}

func (r UserTokenReplaceParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type UserTokenReplaceParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r UserTokenReplaceParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserTokenReplaceResponseEnvelope struct {
	Errors   []UserTokenReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenReplaceResponseEnvelopeJSON    `json:"-"`
}

// userTokenReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserTokenReplaceResponseEnvelope]
type userTokenReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenReplaceResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    userTokenReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserTokenReplaceResponseEnvelopeErrors]
type userTokenReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenReplaceResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    userTokenReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserTokenReplaceResponseEnvelopeMessages]
type userTokenReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenReplaceResponseEnvelopeSuccess bool

const (
	UserTokenReplaceResponseEnvelopeSuccessTrue UserTokenReplaceResponseEnvelopeSuccess = true
)

type UserTokenVerifyResponseEnvelope struct {
	Errors   []UserTokenVerifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenVerifyResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenVerifyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenVerifyResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenVerifyResponseEnvelopeJSON    `json:"-"`
}

// userTokenVerifyResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserTokenVerifyResponseEnvelope]
type userTokenVerifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    userTokenVerifyResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenVerifyResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserTokenVerifyResponseEnvelopeErrors]
type userTokenVerifyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    userTokenVerifyResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenVerifyResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserTokenVerifyResponseEnvelopeMessages]
type userTokenVerifyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenVerifyResponseEnvelopeSuccess bool

const (
	UserTokenVerifyResponseEnvelopeSuccessTrue UserTokenVerifyResponseEnvelopeSuccess = true
)
