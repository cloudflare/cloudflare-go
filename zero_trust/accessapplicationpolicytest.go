// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessApplicationPolicyTestService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationPolicyTestService] method instead.
type AccessApplicationPolicyTestService struct {
	Options []option.RequestOption
	Users   *AccessApplicationPolicyTestUserService
}

// NewAccessApplicationPolicyTestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationPolicyTestService(opts ...option.RequestOption) (r *AccessApplicationPolicyTestService) {
	r = &AccessApplicationPolicyTestService{}
	r.Options = opts
	r.Users = NewAccessApplicationPolicyTestUserService(opts...)
	return
}

// Starts an Access policy test.
func (r *AccessApplicationPolicyTestService) New(ctx context.Context, params AccessApplicationPolicyTestNewParams, opts ...option.RequestOption) (res *AccessApplicationPolicyTestNewResponse, err error) {
	var env AccessApplicationPolicyTestNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current status of a given Access policy test.
func (r *AccessApplicationPolicyTestService) Get(ctx context.Context, policyTestID string, params AccessApplicationPolicyTestGetParams, opts ...option.RequestOption) (res *AccessApplicationPolicyTestGetResponse, err error) {
	var env AccessApplicationPolicyTestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s", params.AccountID, policyTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessApplicationPolicyTestNewResponse struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The status of the policy test request.
	Status AccessApplicationPolicyTestNewResponseStatus `json:"status"`
	JSON   accessApplicationPolicyTestNewResponseJSON   `json:"-"`
}

// accessApplicationPolicyTestNewResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyTestNewResponse]
type accessApplicationPolicyTestNewResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test request.
type AccessApplicationPolicyTestNewResponseStatus string

const (
	AccessApplicationPolicyTestNewResponseStatusSuccess AccessApplicationPolicyTestNewResponseStatus = "success"
)

func (r AccessApplicationPolicyTestNewResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestNewResponseStatusSuccess:
		return true
	}
	return false
}

type AccessApplicationPolicyTestGetResponse struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The number of pages of (processed) users.
	PagesProcessed int64 `json:"pages_processed"`
	// The percentage of (processed) users approved based on policy evaluation results.
	PercentApproved int64 `json:"percent_approved"`
	// The percentage of (processed) users blocked based on policy evaluation results.
	PercentBlocked int64 `json:"percent_blocked"`
	// The percentage of users processed so far (of the entire user base).
	PercentUsersProcessed int64 `json:"percent_users_processed"`
	// The status of the policy test.
	Status AccessApplicationPolicyTestGetResponseStatus `json:"status"`
	// The total number of users in the user base.
	TotalUsers int64 `json:"total_users"`
	// The number of (processed) users approved based on policy evaluation results.
	UsersApproved int64 `json:"users_approved"`
	// The number of (processed) users blocked based on policy evaluation results.
	UsersBlocked int64                                      `json:"users_blocked"`
	JSON         accessApplicationPolicyTestGetResponseJSON `json:"-"`
}

// accessApplicationPolicyTestGetResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyTestGetResponse]
type accessApplicationPolicyTestGetResponseJSON struct {
	ID                    apijson.Field
	PagesProcessed        apijson.Field
	PercentApproved       apijson.Field
	PercentBlocked        apijson.Field
	PercentUsersProcessed apijson.Field
	Status                apijson.Field
	TotalUsers            apijson.Field
	UsersApproved         apijson.Field
	UsersBlocked          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test.
type AccessApplicationPolicyTestGetResponseStatus string

const (
	AccessApplicationPolicyTestGetResponseStatusBlocked    AccessApplicationPolicyTestGetResponseStatus = "blocked"
	AccessApplicationPolicyTestGetResponseStatusProcessing AccessApplicationPolicyTestGetResponseStatus = "processing"
	AccessApplicationPolicyTestGetResponseStatusComplete   AccessApplicationPolicyTestGetResponseStatus = "complete"
)

func (r AccessApplicationPolicyTestGetResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestGetResponseStatusBlocked, AccessApplicationPolicyTestGetResponseStatusProcessing, AccessApplicationPolicyTestGetResponseStatusComplete:
		return true
	}
	return false
}

type AccessApplicationPolicyTestNewParams struct {
	// Identifier
	AccountID param.Field[string]                                       `path:"account_id,required"`
	Policies  param.Field[[]AccessApplicationPolicyTestNewParamsPolicy] `json:"policies"`
}

func (r AccessApplicationPolicyTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyTestNewParamsPolicy struct {
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyTestNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyTestNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyTestNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationPolicyTestNewResponse                `json:"result"`
	JSON    accessApplicationPolicyTestNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyTestNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyTestNewResponseEnvelope]
type accessApplicationPolicyTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationPolicyTestNewResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyTestNewResponseEnvelopeSuccessTrue AccessApplicationPolicyTestNewResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyTestNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationPolicyTestGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
}

// URLQuery serializes [AccessApplicationPolicyTestGetParams]'s query parameters as
// `url.Values`.
func (r AccessApplicationPolicyTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessApplicationPolicyTestGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyTestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationPolicyTestGetResponse                `json:"result"`
	JSON    accessApplicationPolicyTestGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyTestGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyTestGetResponseEnvelope]
type accessApplicationPolicyTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationPolicyTestGetResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyTestGetResponseEnvelopeSuccessTrue AccessApplicationPolicyTestGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
