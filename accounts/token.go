// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/cloudflare/cloudflare-go/v3/user"
)

// TokenService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options          []option.RequestOption
	PermissionGroups *TokenPermissionGroupService
	Value            *TokenValueService
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	r.PermissionGroups = NewTokenPermissionGroupService(opts...)
	r.Value = NewTokenValueService(opts...)
	return
}

// Create a new Account Owned API token.
func (r *TokenService) New(ctx context.Context, params TokenNewParams, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	var env TokenNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing token.
func (r *TokenService) Update(ctx context.Context, tokenID string, params TokenUpdateParams, opts ...option.RequestOption) (res *user.Token, err error) {
	var env TokenUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", params.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Account Owned API tokens created for this account.
func (r *TokenService) List(ctx context.Context, params TokenListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[user.Token], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens", params.AccountID)
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

// List all Account Owned API tokens created for this account.
func (r *TokenService) ListAutoPaging(ctx context.Context, params TokenListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[user.Token] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Destroy an Account Owned API token.
func (r *TokenService) Delete(ctx context.Context, tokenID string, body TokenDeleteParams, opts ...option.RequestOption) (res *TokenDeleteResponse, err error) {
	var env TokenDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", body.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific Account Owned API token.
func (r *TokenService) Get(ctx context.Context, tokenID string, query TokenGetParams, opts ...option.RequestOption) (res *user.Token, err error) {
	var env TokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", query.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test whether a token works.
func (r *TokenService) Verify(ctx context.Context, query TokenVerifyParams, opts ...option.RequestOption) (res *TokenVerifyResponse, err error) {
	var env TokenVerifyResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/verify", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenNewResponse struct {
	// Token identifier tag.
	ID        string                    `json:"id"`
	Condition TokenNewResponseCondition `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time on which the token was created.
	IssuedOn time.Time `json:"issued_on" format:"date-time"`
	// Last time the token was used.
	LastUsedOn time.Time `json:"last_used_on" format:"date-time"`
	// Last time the token was modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Token name.
	Name string `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies []user.Policy `json:"policies"`
	// Status of the token.
	Status TokenNewResponseStatus `json:"status"`
	// The token value.
	Value shared.TokenValue    `json:"value"`
	JSON  tokenNewResponseJSON `json:"-"`
}

// tokenNewResponseJSON contains the JSON metadata for the struct
// [TokenNewResponse]
type tokenNewResponseJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	ExpiresOn   apijson.Field
	IssuedOn    apijson.Field
	LastUsedOn  apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	NotBefore   apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseCondition struct {
	// Client IP restrictions.
	RequestIP TokenNewResponseConditionRequestIP `json:"request.ip"`
	JSON      tokenNewResponseConditionJSON      `json:"-"`
}

// tokenNewResponseConditionJSON contains the JSON metadata for the struct
// [TokenNewResponseCondition]
type tokenNewResponseConditionJSON struct {
	RequestIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseConditionJSON) RawJSON() string {
	return r.raw
}

// Client IP restrictions.
type TokenNewResponseConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In []user.CIDRList `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn []user.CIDRList                        `json:"not_in"`
	JSON  tokenNewResponseConditionRequestIPJSON `json:"-"`
}

// tokenNewResponseConditionRequestIPJSON contains the JSON metadata for the struct
// [TokenNewResponseConditionRequestIP]
type tokenNewResponseConditionRequestIPJSON struct {
	In          apijson.Field
	NotIn       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseConditionRequestIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseConditionRequestIPJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type TokenNewResponseStatus string

const (
	TokenNewResponseStatusActive   TokenNewResponseStatus = "active"
	TokenNewResponseStatusDisabled TokenNewResponseStatus = "disabled"
	TokenNewResponseStatusExpired  TokenNewResponseStatus = "expired"
)

func (r TokenNewResponseStatus) IsKnown() bool {
	switch r {
	case TokenNewResponseStatusActive, TokenNewResponseStatusDisabled, TokenNewResponseStatusExpired:
		return true
	}
	return false
}

type TokenDeleteResponse struct {
	// Identifier
	ID   string                  `json:"id,required"`
	JSON tokenDeleteResponseJSON `json:"-"`
}

// tokenDeleteResponseJSON contains the JSON metadata for the struct
// [TokenDeleteResponse]
type tokenDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponse struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status TokenVerifyResponseStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time               `json:"not_before" format:"date-time"`
	JSON      tokenVerifyResponseJSON `json:"-"`
}

// tokenVerifyResponseJSON contains the JSON metadata for the struct
// [TokenVerifyResponse]
type tokenVerifyResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type TokenVerifyResponseStatus string

const (
	TokenVerifyResponseStatusActive   TokenVerifyResponseStatus = "active"
	TokenVerifyResponseStatusDisabled TokenVerifyResponseStatus = "disabled"
	TokenVerifyResponseStatusExpired  TokenVerifyResponseStatus = "expired"
)

func (r TokenVerifyResponseStatus) IsKnown() bool {
	switch r {
	case TokenVerifyResponseStatusActive, TokenVerifyResponseStatusDisabled, TokenVerifyResponseStatusExpired:
		return true
	}
	return false
}

type TokenNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]user.PolicyParam]      `json:"policies,required"`
	Condition param.Field[TokenNewParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[TokenNewParamsConditionRequestIP] `json:"request.ip"`
}

func (r TokenNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type TokenNewParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]user.CIDRListParam] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]user.CIDRListParam] `json:"not_in"`
}

func (r TokenNewParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TokenNewResponseEnvelopeSuccess `json:"success,required"`
	Result  TokenNewResponse                `json:"result"`
	JSON    tokenNewResponseEnvelopeJSON    `json:"-"`
}

// tokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelope]
type tokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenNewResponseEnvelopeSuccess bool

const (
	TokenNewResponseEnvelopeSuccessTrue TokenNewResponseEnvelopeSuccess = true
)

func (r TokenNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Token     user.TokenParam     `json:"token,required"`
}

func (r TokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Token)
}

type TokenUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  user.Token                         `json:"result"`
	JSON    tokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// tokenUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenUpdateResponseEnvelope]
type tokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenUpdateResponseEnvelopeSuccess bool

const (
	TokenUpdateResponseEnvelopeSuccessTrue TokenUpdateResponseEnvelopeSuccess = true
)

func (r TokenUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order results.
	Direction param.Field[TokenListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [TokenListParams]'s query parameters as `url.Values`.
func (r TokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type TokenListParamsDirection string

const (
	TokenListParamsDirectionAsc  TokenListParamsDirection = "asc"
	TokenListParamsDirectionDesc TokenListParamsDirection = "desc"
)

func (r TokenListParamsDirection) IsKnown() bool {
	switch r {
	case TokenListParamsDirectionAsc, TokenListParamsDirectionDesc:
		return true
	}
	return false
}

type TokenDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  TokenDeleteResponse                `json:"result,nullable"`
	JSON    tokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// tokenDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenDeleteResponseEnvelope]
type tokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenDeleteResponseEnvelopeSuccess bool

const (
	TokenDeleteResponseEnvelopeSuccessTrue TokenDeleteResponseEnvelopeSuccess = true
)

func (r TokenDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TokenGetResponseEnvelopeSuccess `json:"success,required"`
	Result  user.Token                      `json:"result"`
	JSON    tokenGetResponseEnvelopeJSON    `json:"-"`
}

// tokenGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelope]
type tokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenGetResponseEnvelopeSuccess bool

const (
	TokenGetResponseEnvelopeSuccessTrue TokenGetResponseEnvelopeSuccess = true
)

func (r TokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenVerifyParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenVerifyResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TokenVerifyResponseEnvelopeSuccess `json:"success,required"`
	Result  TokenVerifyResponse                `json:"result"`
	JSON    tokenVerifyResponseEnvelopeJSON    `json:"-"`
}

// tokenVerifyResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenVerifyResponseEnvelope]
type tokenVerifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenVerifyResponseEnvelopeSuccess bool

const (
	TokenVerifyResponseEnvelopeSuccessTrue TokenVerifyResponseEnvelopeSuccess = true
)

func (r TokenVerifyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenVerifyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
