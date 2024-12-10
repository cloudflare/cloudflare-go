// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessPolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessPolicyService] method instead.
type AccessPolicyService struct {
	Options []option.RequestOption
}

// NewAccessPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessPolicyService(opts ...option.RequestOption) (r *AccessPolicyService) {
	r = &AccessPolicyService{}
	r.Options = opts
	return
}

// Creates a new Access reusable policy.
func (r *AccessPolicyService) New(ctx context.Context, params AccessPolicyNewParams, opts ...option.RequestOption) (res *AccessPolicyNewResponse, err error) {
	var env AccessPolicyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a Access reusable policy.
func (r *AccessPolicyService) Update(ctx context.Context, policyID string, params AccessPolicyUpdateParams, opts ...option.RequestOption) (res *AccessPolicyUpdateResponse, err error) {
	var env AccessPolicyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Access reusable policies.
func (r *AccessPolicyService) List(ctx context.Context, query AccessPolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessPolicyListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists Access reusable policies.
func (r *AccessPolicyService) ListAutoPaging(ctx context.Context, query AccessPolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessPolicyListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an Access reusable policy.
func (r *AccessPolicyService) Delete(ctx context.Context, policyID string, body AccessPolicyDeleteParams, opts ...option.RequestOption) (res *AccessPolicyDeleteResponse, err error) {
	var env AccessPolicyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access reusable policy.
func (r *AccessPolicyService) Get(ctx context.Context, policyID string, query AccessPolicyGetParams, opts ...option.RequestOption) (res *AccessPolicyGetResponse, err error) {
	var env AccessPolicyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A group of email addresses that can approve a temporary authentication request.
type ApprovalGroupParam struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]string] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID param.Field[string] `json:"email_list_uuid"`
}

func (r ApprovalGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessPolicyNewResponse struct {
	// The UUID of the policy
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// The name of the Access policy.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule                `json:"require"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      accessPolicyNewResponseJSON `json:"-"`
}

// accessPolicyNewResponseJSON contains the JSON metadata for the struct
// [AccessPolicyNewResponse]
type accessPolicyNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Decision    apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessPolicyUpdateResponse struct {
	// The UUID of the policy
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// The name of the Access policy.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule                   `json:"require"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      accessPolicyUpdateResponseJSON `json:"-"`
}

// accessPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [AccessPolicyUpdateResponse]
type accessPolicyUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Decision    apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessPolicyListResponse struct {
	// The UUID of the policy
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// The name of the Access policy.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule                 `json:"require"`
	UpdatedAt time.Time                    `json:"updated_at" format:"date-time"`
	JSON      accessPolicyListResponseJSON `json:"-"`
}

// accessPolicyListResponseJSON contains the JSON metadata for the struct
// [AccessPolicyListResponse]
type accessPolicyListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Decision    apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessPolicyDeleteResponse struct {
	// The UUID of the policy
	ID   string                         `json:"id"`
	JSON accessPolicyDeleteResponseJSON `json:"-"`
}

// accessPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [AccessPolicyDeleteResponse]
type accessPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessPolicyGetResponse struct {
	// The UUID of the policy
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// The name of the Access policy.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule                `json:"require"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      accessPolicyGetResponseJSON `json:"-"`
}

// accessPolicyGetResponseJSON contains the JSON metadata for the struct
// [AccessPolicyGetResponse]
type accessPolicyGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Decision    apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessPolicyNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessPolicyNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessPolicyNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessPolicyNewResponse                `json:"result"`
	JSON    accessPolicyNewResponseEnvelopeJSON    `json:"-"`
}

// accessPolicyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessPolicyNewResponseEnvelope]
type accessPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessPolicyNewResponseEnvelopeSuccess bool

const (
	AccessPolicyNewResponseEnvelopeSuccessTrue AccessPolicyNewResponseEnvelopeSuccess = true
)

func (r AccessPolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessPolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessPolicyUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessPolicyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessPolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessPolicyUpdateResponse                `json:"result"`
	JSON    accessPolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessPolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessPolicyUpdateResponseEnvelope]
type accessPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessPolicyUpdateResponseEnvelopeSuccess bool

const (
	AccessPolicyUpdateResponseEnvelopeSuccessTrue AccessPolicyUpdateResponseEnvelopeSuccess = true
)

func (r AccessPolicyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessPolicyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessPolicyListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessPolicyDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessPolicyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessPolicyDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessPolicyDeleteResponse                `json:"result"`
	JSON    accessPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessPolicyDeleteResponseEnvelope]
type accessPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessPolicyDeleteResponseEnvelopeSuccess bool

const (
	AccessPolicyDeleteResponseEnvelopeSuccessTrue AccessPolicyDeleteResponseEnvelopeSuccess = true
)

func (r AccessPolicyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessPolicyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessPolicyGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessPolicyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessPolicyGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessPolicyGetResponse                `json:"result"`
	JSON    accessPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// accessPolicyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessPolicyGetResponseEnvelope]
type accessPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessPolicyGetResponseEnvelopeSuccess bool

const (
	AccessPolicyGetResponseEnvelopeSuccessTrue AccessPolicyGetResponseEnvelopeSuccess = true
)

func (r AccessPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
