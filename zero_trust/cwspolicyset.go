// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CWSPolicySetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSPolicySetService] method instead.
type CWSPolicySetService struct {
	Options []option.RequestOption
}

// NewCWSPolicySetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSPolicySetService(opts ...option.RequestOption) (r *CWSPolicySetService) {
	r = &CWSPolicySetService{}
	r.Options = opts
	return
}

// Creates a policy set containing the specified policies and workload targets.
func (r *CWSPolicySetService) New(ctx context.Context, params CWSPolicySetNewParams, opts ...option.RequestOption) (res *PolicySet, err error) {
	var env CWSPolicySetNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policy-sets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the policy set's configuration. Omit tags or use {} to clear them.
func (r *CWSPolicySetService) Update(ctx context.Context, policySetID string, params CWSPolicySetUpdateParams, opts ...option.RequestOption) (res *PolicySet, err error) {
	var env CWSPolicySetUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policy-sets/%s", params.AccountID, policySetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of policy sets in the account.
func (r *CWSPolicySetService) List(ctx context.Context, params CWSPolicySetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[PolicySet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policy-sets", params.AccountID)
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

// Returns a paginated list of policy sets in the account.
func (r *CWSPolicySetService) ListAutoPaging(ctx context.Context, params CWSPolicySetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[PolicySet] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the policy set and cascades its deployment attachment and policy and
// workload membership rows. The referenced deployments, policies, and workloads
// remain.
func (r *CWSPolicySetService) Delete(ctx context.Context, policySetID string, body CWSPolicySetDeleteParams, opts ...option.RequestOption) (res *CWSPolicySetDeleteResponse, err error) {
	var env CWSPolicySetDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policy-sets/%s", body.AccountID, policySetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the policy set with the specified ID.
func (r *CWSPolicySetService) Get(ctx context.Context, policySetID string, query CWSPolicySetGetParams, opts ...option.RequestOption) (res *PolicySet, err error) {
	var env CWSPolicySetGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policy-sets/%s", query.AccountID, policySetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type PolicySet struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// IDs of workloads that exclude matching requests. Exclusions take precedence.
	ExcludeWorkloadIDs []string `json:"exclude_workload_ids" api:"required,nullable" format:"uuid"`
	// IDs of workloads that include matching requests. Empty includes all requests
	// before exclusions.
	IncludeWorkloadIDs []string `json:"include_workload_ids" api:"required,nullable" format:"uuid"`
	// Human-readable name; unique per account
	Name string `json:"name" api:"required"`
	// Member policies and their per-membership enforcement mode
	Policies []PolicySetPolicy `json:"policies" api:"required,nullable"`
	// Tags attached to this policy set
	Tags      map[string]string `json:"tags" api:"required"`
	UpdatedAt time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// Optional free-form description.
	Description string        `json:"description"`
	JSON        policySetJSON `json:"-"`
}

// policySetJSON contains the JSON metadata for the struct [PolicySet]
type policySetJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	ExcludeWorkloadIDs apijson.Field
	IncludeWorkloadIDs apijson.Field
	Name               apijson.Field
	Policies           apijson.Field
	Tags               apijson.Field
	UpdatedAt          apijson.Field
	Description        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PolicySet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policySetJSON) RawJSON() string {
	return r.raw
}

type PolicySetPolicy struct {
	// Whether this policy should enforce when it matches. False keeps the policy
	// attached for non-enforcing evaluation.
	Enforce bool `json:"enforce" api:"required"`
	// ID of the member policy
	PolicyID string              `json:"policy_id" api:"required" format:"uuid"`
	JSON     policySetPolicyJSON `json:"-"`
}

// policySetPolicyJSON contains the JSON metadata for the struct [PolicySetPolicy]
type policySetPolicyJSON struct {
	Enforce     apijson.Field
	PolicyID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicySetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policySetPolicyJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetDeleteResponse struct {
	ID   string                         `json:"id" api:"required" format:"uuid"`
	JSON cwsPolicySetDeleteResponseJSON `json:"-"`
}

// cwsPolicySetDeleteResponseJSON contains the JSON metadata for the struct
// [CWSPolicySetDeleteResponse]
type cwsPolicySetDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Name of the policy set
	Name param.Field[string] `json:"name" api:"required"`
	// Policies to include with per-membership enforcement mode
	Policies param.Field[[]CWSPolicySetNewParamsPolicy] `json:"policies" api:"required"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// Workload IDs that exclude matching requests. Exclusions take precedence.
	ExcludeWorkloadIDs param.Field[[]string] `json:"exclude_workload_ids" format:"uuid"`
	// Workload IDs that include matching requests. Empty or omitted includes all
	// requests before exclusions.
	IncludeWorkloadIDs param.Field[[]string] `json:"include_workload_ids" format:"uuid"`
	// Tags to attach. Keys are limited to 256 UTF-8 bytes and values to 1024 UTF-8
	// bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSPolicySetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSPolicySetNewParamsPolicy struct {
	// ID of the member policy
	PolicyID param.Field[string] `json:"policy_id" api:"required" format:"uuid"`
	// Whether this policy should enforce when it matches. Defaults to true; false
	// keeps the policy attached for non-enforcing evaluation.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r CWSPolicySetNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSPolicySetNewResponseEnvelope struct {
	Errors   []CWSPolicySetNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicySetNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   PolicySet                                 `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     cwsPolicySetNewResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicySetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicySetNewResponseEnvelope]
type cwsPolicySetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	Source  CWSPolicySetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicySetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicySetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSPolicySetNewResponseEnvelopeErrors]
type cwsPolicySetNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetNewResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    cwsPolicySetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicySetNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSPolicySetNewResponseEnvelopeErrorsSource]
type cwsPolicySetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSPolicySetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicySetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicySetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSPolicySetNewResponseEnvelopeMessages]
type cwsPolicySetNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetNewResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsPolicySetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicySetNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSPolicySetNewResponseEnvelopeMessagesSource]
type cwsPolicySetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Name of the policy set
	Name param.Field[string] `json:"name" api:"required"`
	// Replacement set of policies with per-membership enforcement mode
	Policies param.Field[[]CWSPolicySetUpdateParamsPolicy] `json:"policies" api:"required"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// Replacement set of excluded workload IDs. Exclusions take precedence.
	ExcludeWorkloadIDs param.Field[[]string] `json:"exclude_workload_ids" format:"uuid"`
	// Replacement set of included workload IDs. Empty or omitted includes all requests
	// before exclusions.
	IncludeWorkloadIDs param.Field[[]string] `json:"include_workload_ids" format:"uuid"`
	// Replacement tag set. Omit or use {} to clear. Keys are limited to 256 UTF-8
	// bytes and values to 1024 UTF-8 bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSPolicySetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSPolicySetUpdateParamsPolicy struct {
	// ID of the member policy
	PolicyID param.Field[string] `json:"policy_id" api:"required" format:"uuid"`
	// Whether this policy should enforce when it matches. Defaults to true; false
	// keeps the policy attached for non-enforcing evaluation.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r CWSPolicySetUpdateParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSPolicySetUpdateResponseEnvelope struct {
	Errors   []CWSPolicySetUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicySetUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   PolicySet                                    `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     cwsPolicySetUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicySetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicySetUpdateResponseEnvelope]
type cwsPolicySetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSPolicySetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicySetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicySetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSPolicySetUpdateResponseEnvelopeErrors]
type cwsPolicySetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsPolicySetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicySetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSPolicySetUpdateResponseEnvelopeErrorsSource]
type cwsPolicySetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	Source  CWSPolicySetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicySetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicySetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSPolicySetUpdateResponseEnvelopeMessages]
type cwsPolicySetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    cwsPolicySetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicySetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSPolicySetUpdateResponseEnvelopeMessagesSource]
type cwsPolicySetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Filter to policy sets attached to this deployment
	DeploymentID param.Field[string] `query:"deployment_id" format:"uuid"`
	// Sort direction (default: asc)
	Direction param.Field[CWSPolicySetListParamsDirection] `query:"direction"`
	// Filter by policy set name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSPolicySetListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSPolicySetListParams]'s query parameters as `url.Values`.
func (r CWSPolicySetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSPolicySetListParamsDirection string

const (
	CWSPolicySetListParamsDirectionAsc  CWSPolicySetListParamsDirection = "asc"
	CWSPolicySetListParamsDirectionDesc CWSPolicySetListParamsDirection = "desc"
)

func (r CWSPolicySetListParamsDirection) IsKnown() bool {
	switch r {
	case CWSPolicySetListParamsDirectionAsc, CWSPolicySetListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSPolicySetListParamsOrder string

const (
	CWSPolicySetListParamsOrderName      CWSPolicySetListParamsOrder = "name"
	CWSPolicySetListParamsOrderCreatedAt CWSPolicySetListParamsOrder = "created_at"
	CWSPolicySetListParamsOrderUpdatedAt CWSPolicySetListParamsOrder = "updated_at"
)

func (r CWSPolicySetListParamsOrder) IsKnown() bool {
	switch r {
	case CWSPolicySetListParamsOrderName, CWSPolicySetListParamsOrderCreatedAt, CWSPolicySetListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSPolicySetDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSPolicySetDeleteResponseEnvelope struct {
	Errors   []CWSPolicySetDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicySetDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSPolicySetDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     cwsPolicySetDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicySetDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicySetDeleteResponseEnvelope]
type cwsPolicySetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSPolicySetDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicySetDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicySetDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSPolicySetDeleteResponseEnvelopeErrors]
type cwsPolicySetDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsPolicySetDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicySetDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSPolicySetDeleteResponseEnvelopeErrorsSource]
type cwsPolicySetDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	Source  CWSPolicySetDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicySetDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicySetDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSPolicySetDeleteResponseEnvelopeMessages]
type cwsPolicySetDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    cwsPolicySetDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicySetDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSPolicySetDeleteResponseEnvelopeMessagesSource]
type cwsPolicySetDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSPolicySetGetResponseEnvelope struct {
	Errors   []CWSPolicySetGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicySetGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   PolicySet                                 `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     cwsPolicySetGetResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicySetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicySetGetResponseEnvelope]
type cwsPolicySetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	Source  CWSPolicySetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicySetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicySetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSPolicySetGetResponseEnvelopeErrors]
type cwsPolicySetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    cwsPolicySetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicySetGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSPolicySetGetResponseEnvelopeErrorsSource]
type cwsPolicySetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSPolicySetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicySetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicySetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSPolicySetGetResponseEnvelopeMessages]
type cwsPolicySetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicySetGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsPolicySetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicySetGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSPolicySetGetResponseEnvelopeMessagesSource]
type cwsPolicySetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicySetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicySetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
