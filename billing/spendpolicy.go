// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// SpendPolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendPolicyService] method instead.
type SpendPolicyService struct {
	Options  []option.RequestOption
	Audit    *SpendPolicyAuditService
	Activity *SpendPolicyActivityService
}

// NewSpendPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpendPolicyService(opts ...option.RequestOption) (r *SpendPolicyService) {
	r = &SpendPolicyService{}
	r.Options = opts
	r.Audit = NewSpendPolicyAuditService(opts...)
	r.Activity = NewSpendPolicyActivityService(opts...)
	return
}

// Creates a spend policy. Each account can have only one enabled `usage_block`
// policy per product. A second one returns `409`. There is no limit on
// `notification` policies.
func (r *SpendPolicyService) New(ctx context.Context, params SpendPolicyNewParams, opts ...option.RequestOption) (res *SpendPolicyNewResponse, err error) {
	var env SpendPolicyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the whole spend policy. Every field is required, as on create; there
// are no partial updates and `PATCH` is not supported. The `If-Match` header must
// equal the policy's current `version`. On success, `version` goes up by one and
// the new value is returned in `ETag`.
func (r *SpendPolicyService) Update(ctx context.Context, policyID string, params SpendPolicyUpdateParams, opts ...option.RequestOption) (res *SpendPolicyUpdateResponse, err error) {
	var env SpendPolicyUpdateResponseEnvelope
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns every spend policy on the account that has not been deleted, including
// disabled ones. Results are sorted by `name`, then `id`. There is no pagination.
func (r *SpendPolicyService) List(ctx context.Context, query SpendPolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SpendPolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies", query.AccountID)
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

// Returns every spend policy on the account that has not been deleted, including
// disabled ones. Results are sorted by `name`, then `id`. There is no pagination.
func (r *SpendPolicyService) ListAutoPaging(ctx context.Context, query SpendPolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SpendPolicyListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Soft-deletes the spend policy. The policy is disabled and no longer appears in
// list, get or audit responses. The `If-Match` header must equal the policy's
// current `version`.
func (r *SpendPolicyService) Delete(ctx context.Context, policyID string, params SpendPolicyDeleteParams, opts ...option.RequestOption) (res *SpendPolicyDeleteResponse, err error) {
	var env SpendPolicyDeleteResponseEnvelope
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns one spend policy. The `ETag` response header holds the current
// `version`; send it as `If-Match` on a later update or delete.
func (r *SpendPolicyService) Get(ctx context.Context, policyID string, query SpendPolicyGetParams, opts ...option.RequestOption) (res *SpendPolicyGetResponse, err error) {
	var env SpendPolicyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A spend policy.
type SpendPolicyNewResponse struct {
	// The unique identifier of the spend policy.
	ID string `json:"id" api:"required" format:"uuid"`
	// Internal numeric ID of the account that owns the policy. This is not the account
	// tag used in the path.
	AccountID int64 `json:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyNewResponseAction `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled bool `json:"enabled" api:"required"`
	// Display name of the policy.
	Name string `json:"name" api:"required"`
	// Email address that receives notifications. An empty string when none is set.
	NotificationEmail string `json:"notification_email" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID SpendPolicyNewResponseProductID `json:"product_id" api:"required"`
	// The spend threshold for the current billing period, in US cents.
	ThresholdCents int64 `json:"threshold_cents" api:"required"`
	// Optimistic-concurrency version. It starts at 1 and goes up by one on every
	// update or delete. This value is also returned in `ETag`.
	Version int64 `json:"version" api:"required"`
	// Read-only state on list/get responses, evaluated from the latest owning
	// projector snapshot at request time, independently of the API replica.
	// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
	// history remains unchanged. Missing during a rolling upgrade or on mutation
	// responses; clients must treat missing state as unavailable. Balances are
	// asynchronous estimates, not a synchronous usage read.
	CurrentState SpendPolicyNewResponseCurrentState `json:"current_state"`
	JSON         spendPolicyNewResponseJSON         `json:"-"`
}

// spendPolicyNewResponseJSON contains the JSON metadata for the struct
// [SpendPolicyNewResponse]
type spendPolicyNewResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	Action            apijson.Field
	Enabled           apijson.Field
	Name              apijson.Field
	NotificationEmail apijson.Field
	ProductID         apijson.Field
	ThresholdCents    apijson.Field
	Version           apijson.Field
	CurrentState      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyNewResponseAction string

const (
	SpendPolicyNewResponseActionNotification SpendPolicyNewResponseAction = "notification"
	SpendPolicyNewResponseActionUsageBlock   SpendPolicyNewResponseAction = "usage_block"
)

func (r SpendPolicyNewResponseAction) IsKnown() bool {
	switch r {
	case SpendPolicyNewResponseActionNotification, SpendPolicyNewResponseActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyNewResponseProductID string

const (
	SpendPolicyNewResponseProductIDWorkers        SpendPolicyNewResponseProductID = "workers"
	SpendPolicyNewResponseProductIDDurableObjects SpendPolicyNewResponseProductID = "durable_objects"
)

func (r SpendPolicyNewResponseProductID) IsKnown() bool {
	switch r {
	case SpendPolicyNewResponseProductIDWorkers, SpendPolicyNewResponseProductIDDurableObjects:
		return true
	}
	return false
}

// Read-only state on list/get responses, evaluated from the latest owning
// projector snapshot at request time, independently of the API replica.
// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
// history remains unchanged. Missing during a rolling upgrade or on mutation
// responses; clients must treat missing state as unavailable. Balances are
// asynchronous estimates, not a synchronous usage read.
type SpendPolicyNewResponseCurrentState struct {
	BalanceCents       int64     `json:"balance_cents" api:"required,nullable"`
	BillingPeriodEnd   time.Time `json:"billing_period_end" api:"required,nullable" format:"date-time"`
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required,nullable" format:"date-time"`
	// Durable desired cap and delivery bookkeeping for this policy, or null if no
	// owned effect exists. Separate from threshold evaluation: pending/leased/retry
	// can lag the new billing period. Succeeded with equal generations means delivery
	// completed, not an EWC readback.
	Enforcement SpendPolicyNewResponseCurrentStateEnforcement `json:"enforcement" api:"required,nullable"`
	// Evaluation time; not the timestamp of the latest usage event.
	EvaluatedAt time.Time `json:"evaluated_at" api:"required" format:"date-time"`
	// Threshold evaluation, not proof of an edge cap. No current active period yields
	// no_current_period for enabled policies; failed aggregate reads yield unavailable
	// for enabled policies. Disabled policies report disabled. Neither
	// no_current_period nor unavailable implies a zero balance or below-limit state.
	// Operators may disable current-state reads, in which case current_state is
	// omitted.
	Status SpendPolicyNewResponseCurrentStateStatus `json:"status" api:"required"`
	JSON   spendPolicyNewResponseCurrentStateJSON   `json:"-"`
}

// spendPolicyNewResponseCurrentStateJSON contains the JSON metadata for the struct
// [SpendPolicyNewResponseCurrentState]
type spendPolicyNewResponseCurrentStateJSON struct {
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	Enforcement        apijson.Field
	EvaluatedAt        apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyNewResponseCurrentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyNewResponseCurrentStateJSON) RawJSON() string {
	return r.raw
}

// Durable desired cap and delivery bookkeeping for this policy, or null if no
// owned effect exists. Separate from threshold evaluation: pending/leased/retry
// can lag the new billing period. Succeeded with equal generations means delivery
// completed, not an EWC readback.
type SpendPolicyNewResponseCurrentStateEnforcement struct {
	AppliedGeneration int64                                                       `json:"applied_generation" api:"required"`
	DeliveryStatus    SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus `json:"delivery_status" api:"required"`
	DesiredGeneration int64                                                       `json:"desired_generation" api:"required"`
	DesiredState      SpendPolicyNewResponseCurrentStateEnforcementDesiredState   `json:"desired_state" api:"required"`
	ExpiresAt         time.Time                                                   `json:"expires_at" api:"required,nullable" format:"date-time"`
	JSON              spendPolicyNewResponseCurrentStateEnforcementJSON           `json:"-"`
}

// spendPolicyNewResponseCurrentStateEnforcementJSON contains the JSON metadata for
// the struct [SpendPolicyNewResponseCurrentStateEnforcement]
type spendPolicyNewResponseCurrentStateEnforcementJSON struct {
	AppliedGeneration apijson.Field
	DeliveryStatus    apijson.Field
	DesiredGeneration apijson.Field
	DesiredState      apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyNewResponseCurrentStateEnforcement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyNewResponseCurrentStateEnforcementJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus string

const (
	SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusPending   SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus = "pending"
	SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusLeased    SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus = "leased"
	SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusRetry     SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus = "retry"
	SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusSucceeded SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus = "succeeded"
)

func (r SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatus) IsKnown() bool {
	switch r {
	case SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusPending, SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusLeased, SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusRetry, SpendPolicyNewResponseCurrentStateEnforcementDeliveryStatusSucceeded:
		return true
	}
	return false
}

type SpendPolicyNewResponseCurrentStateEnforcementDesiredState string

const (
	SpendPolicyNewResponseCurrentStateEnforcementDesiredStateActive   SpendPolicyNewResponseCurrentStateEnforcementDesiredState = "active"
	SpendPolicyNewResponseCurrentStateEnforcementDesiredStateInactive SpendPolicyNewResponseCurrentStateEnforcementDesiredState = "inactive"
)

func (r SpendPolicyNewResponseCurrentStateEnforcementDesiredState) IsKnown() bool {
	switch r {
	case SpendPolicyNewResponseCurrentStateEnforcementDesiredStateActive, SpendPolicyNewResponseCurrentStateEnforcementDesiredStateInactive:
		return true
	}
	return false
}

// Threshold evaluation, not proof of an edge cap. No current active period yields
// no_current_period for enabled policies; failed aggregate reads yield unavailable
// for enabled policies. Disabled policies report disabled. Neither
// no_current_period nor unavailable implies a zero balance or below-limit state.
// Operators may disable current-state reads, in which case current_state is
// omitted.
type SpendPolicyNewResponseCurrentStateStatus string

const (
	SpendPolicyNewResponseCurrentStateStatusDisabled        SpendPolicyNewResponseCurrentStateStatus = "disabled"
	SpendPolicyNewResponseCurrentStateStatusNoCurrentPeriod SpendPolicyNewResponseCurrentStateStatus = "no_current_period"
	SpendPolicyNewResponseCurrentStateStatusUnavailable     SpendPolicyNewResponseCurrentStateStatus = "unavailable"
	SpendPolicyNewResponseCurrentStateStatusBelowLimit      SpendPolicyNewResponseCurrentStateStatus = "below_limit"
	SpendPolicyNewResponseCurrentStateStatusLimitReached    SpendPolicyNewResponseCurrentStateStatus = "limit_reached"
)

func (r SpendPolicyNewResponseCurrentStateStatus) IsKnown() bool {
	switch r {
	case SpendPolicyNewResponseCurrentStateStatusDisabled, SpendPolicyNewResponseCurrentStateStatusNoCurrentPeriod, SpendPolicyNewResponseCurrentStateStatusUnavailable, SpendPolicyNewResponseCurrentStateStatusBelowLimit, SpendPolicyNewResponseCurrentStateStatusLimitReached:
		return true
	}
	return false
}

// A spend policy.
type SpendPolicyUpdateResponse struct {
	// The unique identifier of the spend policy.
	ID string `json:"id" api:"required" format:"uuid"`
	// Internal numeric ID of the account that owns the policy. This is not the account
	// tag used in the path.
	AccountID int64 `json:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyUpdateResponseAction `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled bool `json:"enabled" api:"required"`
	// Display name of the policy.
	Name string `json:"name" api:"required"`
	// Email address that receives notifications. An empty string when none is set.
	NotificationEmail string `json:"notification_email" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID SpendPolicyUpdateResponseProductID `json:"product_id" api:"required"`
	// The spend threshold for the current billing period, in US cents.
	ThresholdCents int64 `json:"threshold_cents" api:"required"`
	// Optimistic-concurrency version. It starts at 1 and goes up by one on every
	// update or delete. This value is also returned in `ETag`.
	Version int64 `json:"version" api:"required"`
	// Read-only state on list/get responses, evaluated from the latest owning
	// projector snapshot at request time, independently of the API replica.
	// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
	// history remains unchanged. Missing during a rolling upgrade or on mutation
	// responses; clients must treat missing state as unavailable. Balances are
	// asynchronous estimates, not a synchronous usage read.
	CurrentState SpendPolicyUpdateResponseCurrentState `json:"current_state"`
	JSON         spendPolicyUpdateResponseJSON         `json:"-"`
}

// spendPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [SpendPolicyUpdateResponse]
type spendPolicyUpdateResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	Action            apijson.Field
	Enabled           apijson.Field
	Name              apijson.Field
	NotificationEmail apijson.Field
	ProductID         apijson.Field
	ThresholdCents    apijson.Field
	Version           apijson.Field
	CurrentState      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyUpdateResponseAction string

const (
	SpendPolicyUpdateResponseActionNotification SpendPolicyUpdateResponseAction = "notification"
	SpendPolicyUpdateResponseActionUsageBlock   SpendPolicyUpdateResponseAction = "usage_block"
)

func (r SpendPolicyUpdateResponseAction) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateResponseActionNotification, SpendPolicyUpdateResponseActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyUpdateResponseProductID string

const (
	SpendPolicyUpdateResponseProductIDWorkers        SpendPolicyUpdateResponseProductID = "workers"
	SpendPolicyUpdateResponseProductIDDurableObjects SpendPolicyUpdateResponseProductID = "durable_objects"
)

func (r SpendPolicyUpdateResponseProductID) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateResponseProductIDWorkers, SpendPolicyUpdateResponseProductIDDurableObjects:
		return true
	}
	return false
}

// Read-only state on list/get responses, evaluated from the latest owning
// projector snapshot at request time, independently of the API replica.
// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
// history remains unchanged. Missing during a rolling upgrade or on mutation
// responses; clients must treat missing state as unavailable. Balances are
// asynchronous estimates, not a synchronous usage read.
type SpendPolicyUpdateResponseCurrentState struct {
	BalanceCents       int64     `json:"balance_cents" api:"required,nullable"`
	BillingPeriodEnd   time.Time `json:"billing_period_end" api:"required,nullable" format:"date-time"`
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required,nullable" format:"date-time"`
	// Durable desired cap and delivery bookkeeping for this policy, or null if no
	// owned effect exists. Separate from threshold evaluation: pending/leased/retry
	// can lag the new billing period. Succeeded with equal generations means delivery
	// completed, not an EWC readback.
	Enforcement SpendPolicyUpdateResponseCurrentStateEnforcement `json:"enforcement" api:"required,nullable"`
	// Evaluation time; not the timestamp of the latest usage event.
	EvaluatedAt time.Time `json:"evaluated_at" api:"required" format:"date-time"`
	// Threshold evaluation, not proof of an edge cap. No current active period yields
	// no_current_period for enabled policies; failed aggregate reads yield unavailable
	// for enabled policies. Disabled policies report disabled. Neither
	// no_current_period nor unavailable implies a zero balance or below-limit state.
	// Operators may disable current-state reads, in which case current_state is
	// omitted.
	Status SpendPolicyUpdateResponseCurrentStateStatus `json:"status" api:"required"`
	JSON   spendPolicyUpdateResponseCurrentStateJSON   `json:"-"`
}

// spendPolicyUpdateResponseCurrentStateJSON contains the JSON metadata for the
// struct [SpendPolicyUpdateResponseCurrentState]
type spendPolicyUpdateResponseCurrentStateJSON struct {
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	Enforcement        apijson.Field
	EvaluatedAt        apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyUpdateResponseCurrentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyUpdateResponseCurrentStateJSON) RawJSON() string {
	return r.raw
}

// Durable desired cap and delivery bookkeeping for this policy, or null if no
// owned effect exists. Separate from threshold evaluation: pending/leased/retry
// can lag the new billing period. Succeeded with equal generations means delivery
// completed, not an EWC readback.
type SpendPolicyUpdateResponseCurrentStateEnforcement struct {
	AppliedGeneration int64                                                          `json:"applied_generation" api:"required"`
	DeliveryStatus    SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus `json:"delivery_status" api:"required"`
	DesiredGeneration int64                                                          `json:"desired_generation" api:"required"`
	DesiredState      SpendPolicyUpdateResponseCurrentStateEnforcementDesiredState   `json:"desired_state" api:"required"`
	ExpiresAt         time.Time                                                      `json:"expires_at" api:"required,nullable" format:"date-time"`
	JSON              spendPolicyUpdateResponseCurrentStateEnforcementJSON           `json:"-"`
}

// spendPolicyUpdateResponseCurrentStateEnforcementJSON contains the JSON metadata
// for the struct [SpendPolicyUpdateResponseCurrentStateEnforcement]
type spendPolicyUpdateResponseCurrentStateEnforcementJSON struct {
	AppliedGeneration apijson.Field
	DeliveryStatus    apijson.Field
	DesiredGeneration apijson.Field
	DesiredState      apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyUpdateResponseCurrentStateEnforcement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyUpdateResponseCurrentStateEnforcementJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus string

const (
	SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusPending   SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus = "pending"
	SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusLeased    SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus = "leased"
	SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusRetry     SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus = "retry"
	SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusSucceeded SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus = "succeeded"
)

func (r SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatus) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusPending, SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusLeased, SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusRetry, SpendPolicyUpdateResponseCurrentStateEnforcementDeliveryStatusSucceeded:
		return true
	}
	return false
}

type SpendPolicyUpdateResponseCurrentStateEnforcementDesiredState string

const (
	SpendPolicyUpdateResponseCurrentStateEnforcementDesiredStateActive   SpendPolicyUpdateResponseCurrentStateEnforcementDesiredState = "active"
	SpendPolicyUpdateResponseCurrentStateEnforcementDesiredStateInactive SpendPolicyUpdateResponseCurrentStateEnforcementDesiredState = "inactive"
)

func (r SpendPolicyUpdateResponseCurrentStateEnforcementDesiredState) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateResponseCurrentStateEnforcementDesiredStateActive, SpendPolicyUpdateResponseCurrentStateEnforcementDesiredStateInactive:
		return true
	}
	return false
}

// Threshold evaluation, not proof of an edge cap. No current active period yields
// no_current_period for enabled policies; failed aggregate reads yield unavailable
// for enabled policies. Disabled policies report disabled. Neither
// no_current_period nor unavailable implies a zero balance or below-limit state.
// Operators may disable current-state reads, in which case current_state is
// omitted.
type SpendPolicyUpdateResponseCurrentStateStatus string

const (
	SpendPolicyUpdateResponseCurrentStateStatusDisabled        SpendPolicyUpdateResponseCurrentStateStatus = "disabled"
	SpendPolicyUpdateResponseCurrentStateStatusNoCurrentPeriod SpendPolicyUpdateResponseCurrentStateStatus = "no_current_period"
	SpendPolicyUpdateResponseCurrentStateStatusUnavailable     SpendPolicyUpdateResponseCurrentStateStatus = "unavailable"
	SpendPolicyUpdateResponseCurrentStateStatusBelowLimit      SpendPolicyUpdateResponseCurrentStateStatus = "below_limit"
	SpendPolicyUpdateResponseCurrentStateStatusLimitReached    SpendPolicyUpdateResponseCurrentStateStatus = "limit_reached"
)

func (r SpendPolicyUpdateResponseCurrentStateStatus) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateResponseCurrentStateStatusDisabled, SpendPolicyUpdateResponseCurrentStateStatusNoCurrentPeriod, SpendPolicyUpdateResponseCurrentStateStatusUnavailable, SpendPolicyUpdateResponseCurrentStateStatusBelowLimit, SpendPolicyUpdateResponseCurrentStateStatusLimitReached:
		return true
	}
	return false
}

// A spend policy.
type SpendPolicyListResponse struct {
	// The unique identifier of the spend policy.
	ID string `json:"id" api:"required" format:"uuid"`
	// Internal numeric ID of the account that owns the policy. This is not the account
	// tag used in the path.
	AccountID int64 `json:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyListResponseAction `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled bool `json:"enabled" api:"required"`
	// Display name of the policy.
	Name string `json:"name" api:"required"`
	// Email address that receives notifications. An empty string when none is set.
	NotificationEmail string `json:"notification_email" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID SpendPolicyListResponseProductID `json:"product_id" api:"required"`
	// The spend threshold for the current billing period, in US cents.
	ThresholdCents int64 `json:"threshold_cents" api:"required"`
	// Optimistic-concurrency version. It starts at 1 and goes up by one on every
	// update or delete. This value is also returned in `ETag`.
	Version int64 `json:"version" api:"required"`
	// Read-only state on list/get responses, evaluated from the latest owning
	// projector snapshot at request time, independently of the API replica.
	// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
	// history remains unchanged. Missing during a rolling upgrade or on mutation
	// responses; clients must treat missing state as unavailable. Balances are
	// asynchronous estimates, not a synchronous usage read.
	CurrentState SpendPolicyListResponseCurrentState `json:"current_state"`
	JSON         spendPolicyListResponseJSON         `json:"-"`
}

// spendPolicyListResponseJSON contains the JSON metadata for the struct
// [SpendPolicyListResponse]
type spendPolicyListResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	Action            apijson.Field
	Enabled           apijson.Field
	Name              apijson.Field
	NotificationEmail apijson.Field
	ProductID         apijson.Field
	ThresholdCents    apijson.Field
	Version           apijson.Field
	CurrentState      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyListResponseAction string

const (
	SpendPolicyListResponseActionNotification SpendPolicyListResponseAction = "notification"
	SpendPolicyListResponseActionUsageBlock   SpendPolicyListResponseAction = "usage_block"
)

func (r SpendPolicyListResponseAction) IsKnown() bool {
	switch r {
	case SpendPolicyListResponseActionNotification, SpendPolicyListResponseActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyListResponseProductID string

const (
	SpendPolicyListResponseProductIDWorkers        SpendPolicyListResponseProductID = "workers"
	SpendPolicyListResponseProductIDDurableObjects SpendPolicyListResponseProductID = "durable_objects"
)

func (r SpendPolicyListResponseProductID) IsKnown() bool {
	switch r {
	case SpendPolicyListResponseProductIDWorkers, SpendPolicyListResponseProductIDDurableObjects:
		return true
	}
	return false
}

// Read-only state on list/get responses, evaluated from the latest owning
// projector snapshot at request time, independently of the API replica.
// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
// history remains unchanged. Missing during a rolling upgrade or on mutation
// responses; clients must treat missing state as unavailable. Balances are
// asynchronous estimates, not a synchronous usage read.
type SpendPolicyListResponseCurrentState struct {
	BalanceCents       int64     `json:"balance_cents" api:"required,nullable"`
	BillingPeriodEnd   time.Time `json:"billing_period_end" api:"required,nullable" format:"date-time"`
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required,nullable" format:"date-time"`
	// Durable desired cap and delivery bookkeeping for this policy, or null if no
	// owned effect exists. Separate from threshold evaluation: pending/leased/retry
	// can lag the new billing period. Succeeded with equal generations means delivery
	// completed, not an EWC readback.
	Enforcement SpendPolicyListResponseCurrentStateEnforcement `json:"enforcement" api:"required,nullable"`
	// Evaluation time; not the timestamp of the latest usage event.
	EvaluatedAt time.Time `json:"evaluated_at" api:"required" format:"date-time"`
	// Threshold evaluation, not proof of an edge cap. No current active period yields
	// no_current_period for enabled policies; failed aggregate reads yield unavailable
	// for enabled policies. Disabled policies report disabled. Neither
	// no_current_period nor unavailable implies a zero balance or below-limit state.
	// Operators may disable current-state reads, in which case current_state is
	// omitted.
	Status SpendPolicyListResponseCurrentStateStatus `json:"status" api:"required"`
	JSON   spendPolicyListResponseCurrentStateJSON   `json:"-"`
}

// spendPolicyListResponseCurrentStateJSON contains the JSON metadata for the
// struct [SpendPolicyListResponseCurrentState]
type spendPolicyListResponseCurrentStateJSON struct {
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	Enforcement        apijson.Field
	EvaluatedAt        apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyListResponseCurrentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyListResponseCurrentStateJSON) RawJSON() string {
	return r.raw
}

// Durable desired cap and delivery bookkeeping for this policy, or null if no
// owned effect exists. Separate from threshold evaluation: pending/leased/retry
// can lag the new billing period. Succeeded with equal generations means delivery
// completed, not an EWC readback.
type SpendPolicyListResponseCurrentStateEnforcement struct {
	AppliedGeneration int64                                                        `json:"applied_generation" api:"required"`
	DeliveryStatus    SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus `json:"delivery_status" api:"required"`
	DesiredGeneration int64                                                        `json:"desired_generation" api:"required"`
	DesiredState      SpendPolicyListResponseCurrentStateEnforcementDesiredState   `json:"desired_state" api:"required"`
	ExpiresAt         time.Time                                                    `json:"expires_at" api:"required,nullable" format:"date-time"`
	JSON              spendPolicyListResponseCurrentStateEnforcementJSON           `json:"-"`
}

// spendPolicyListResponseCurrentStateEnforcementJSON contains the JSON metadata
// for the struct [SpendPolicyListResponseCurrentStateEnforcement]
type spendPolicyListResponseCurrentStateEnforcementJSON struct {
	AppliedGeneration apijson.Field
	DeliveryStatus    apijson.Field
	DesiredGeneration apijson.Field
	DesiredState      apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyListResponseCurrentStateEnforcement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyListResponseCurrentStateEnforcementJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus string

const (
	SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusPending   SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus = "pending"
	SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusLeased    SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus = "leased"
	SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusRetry     SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus = "retry"
	SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusSucceeded SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus = "succeeded"
)

func (r SpendPolicyListResponseCurrentStateEnforcementDeliveryStatus) IsKnown() bool {
	switch r {
	case SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusPending, SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusLeased, SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusRetry, SpendPolicyListResponseCurrentStateEnforcementDeliveryStatusSucceeded:
		return true
	}
	return false
}

type SpendPolicyListResponseCurrentStateEnforcementDesiredState string

const (
	SpendPolicyListResponseCurrentStateEnforcementDesiredStateActive   SpendPolicyListResponseCurrentStateEnforcementDesiredState = "active"
	SpendPolicyListResponseCurrentStateEnforcementDesiredStateInactive SpendPolicyListResponseCurrentStateEnforcementDesiredState = "inactive"
)

func (r SpendPolicyListResponseCurrentStateEnforcementDesiredState) IsKnown() bool {
	switch r {
	case SpendPolicyListResponseCurrentStateEnforcementDesiredStateActive, SpendPolicyListResponseCurrentStateEnforcementDesiredStateInactive:
		return true
	}
	return false
}

// Threshold evaluation, not proof of an edge cap. No current active period yields
// no_current_period for enabled policies; failed aggregate reads yield unavailable
// for enabled policies. Disabled policies report disabled. Neither
// no_current_period nor unavailable implies a zero balance or below-limit state.
// Operators may disable current-state reads, in which case current_state is
// omitted.
type SpendPolicyListResponseCurrentStateStatus string

const (
	SpendPolicyListResponseCurrentStateStatusDisabled        SpendPolicyListResponseCurrentStateStatus = "disabled"
	SpendPolicyListResponseCurrentStateStatusNoCurrentPeriod SpendPolicyListResponseCurrentStateStatus = "no_current_period"
	SpendPolicyListResponseCurrentStateStatusUnavailable     SpendPolicyListResponseCurrentStateStatus = "unavailable"
	SpendPolicyListResponseCurrentStateStatusBelowLimit      SpendPolicyListResponseCurrentStateStatus = "below_limit"
	SpendPolicyListResponseCurrentStateStatusLimitReached    SpendPolicyListResponseCurrentStateStatus = "limit_reached"
)

func (r SpendPolicyListResponseCurrentStateStatus) IsKnown() bool {
	switch r {
	case SpendPolicyListResponseCurrentStateStatusDisabled, SpendPolicyListResponseCurrentStateStatusNoCurrentPeriod, SpendPolicyListResponseCurrentStateStatusUnavailable, SpendPolicyListResponseCurrentStateStatusBelowLimit, SpendPolicyListResponseCurrentStateStatusLimitReached:
		return true
	}
	return false
}

type SpendPolicyDeleteResponse struct {
	// The unique identifier of the spend policy.
	ID   string                        `json:"id" api:"required" format:"uuid"`
	JSON spendPolicyDeleteResponseJSON `json:"-"`
}

// spendPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [SpendPolicyDeleteResponse]
type spendPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A spend policy.
type SpendPolicyGetResponse struct {
	// The unique identifier of the spend policy.
	ID string `json:"id" api:"required" format:"uuid"`
	// Internal numeric ID of the account that owns the policy. This is not the account
	// tag used in the path.
	AccountID int64 `json:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyGetResponseAction `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled bool `json:"enabled" api:"required"`
	// Display name of the policy.
	Name string `json:"name" api:"required"`
	// Email address that receives notifications. An empty string when none is set.
	NotificationEmail string `json:"notification_email" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID SpendPolicyGetResponseProductID `json:"product_id" api:"required"`
	// The spend threshold for the current billing period, in US cents.
	ThresholdCents int64 `json:"threshold_cents" api:"required"`
	// Optimistic-concurrency version. It starts at 1 and goes up by one on every
	// update or delete. This value is also returned in `ETag`.
	Version int64 `json:"version" api:"required"`
	// Read-only state on list/get responses, evaluated from the latest owning
	// projector snapshot at request time, independently of the API replica.
	// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
	// history remains unchanged. Missing during a rolling upgrade or on mutation
	// responses; clients must treat missing state as unavailable. Balances are
	// asynchronous estimates, not a synchronous usage read.
	CurrentState SpendPolicyGetResponseCurrentState `json:"current_state"`
	JSON         spendPolicyGetResponseJSON         `json:"-"`
}

// spendPolicyGetResponseJSON contains the JSON metadata for the struct
// [SpendPolicyGetResponse]
type spendPolicyGetResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	Action            apijson.Field
	Enabled           apijson.Field
	Name              apijson.Field
	NotificationEmail apijson.Field
	ProductID         apijson.Field
	ThresholdCents    apijson.Field
	Version           apijson.Field
	CurrentState      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyGetResponseAction string

const (
	SpendPolicyGetResponseActionNotification SpendPolicyGetResponseAction = "notification"
	SpendPolicyGetResponseActionUsageBlock   SpendPolicyGetResponseAction = "usage_block"
)

func (r SpendPolicyGetResponseAction) IsKnown() bool {
	switch r {
	case SpendPolicyGetResponseActionNotification, SpendPolicyGetResponseActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyGetResponseProductID string

const (
	SpendPolicyGetResponseProductIDWorkers        SpendPolicyGetResponseProductID = "workers"
	SpendPolicyGetResponseProductIDDurableObjects SpendPolicyGetResponseProductID = "durable_objects"
)

func (r SpendPolicyGetResponseProductID) IsKnown() bool {
	switch r {
	case SpendPolicyGetResponseProductIDWorkers, SpendPolicyGetResponseProductIDDurableObjects:
		return true
	}
	return false
}

// Read-only state on list/get responses, evaluated from the latest owning
// projector snapshot at request time, independently of the API replica.
// Configuration enabled/version/ETag do not change when a cycle rolls over. Audit
// history remains unchanged. Missing during a rolling upgrade or on mutation
// responses; clients must treat missing state as unavailable. Balances are
// asynchronous estimates, not a synchronous usage read.
type SpendPolicyGetResponseCurrentState struct {
	BalanceCents       int64     `json:"balance_cents" api:"required,nullable"`
	BillingPeriodEnd   time.Time `json:"billing_period_end" api:"required,nullable" format:"date-time"`
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required,nullable" format:"date-time"`
	// Durable desired cap and delivery bookkeeping for this policy, or null if no
	// owned effect exists. Separate from threshold evaluation: pending/leased/retry
	// can lag the new billing period. Succeeded with equal generations means delivery
	// completed, not an EWC readback.
	Enforcement SpendPolicyGetResponseCurrentStateEnforcement `json:"enforcement" api:"required,nullable"`
	// Evaluation time; not the timestamp of the latest usage event.
	EvaluatedAt time.Time `json:"evaluated_at" api:"required" format:"date-time"`
	// Threshold evaluation, not proof of an edge cap. No current active period yields
	// no_current_period for enabled policies; failed aggregate reads yield unavailable
	// for enabled policies. Disabled policies report disabled. Neither
	// no_current_period nor unavailable implies a zero balance or below-limit state.
	// Operators may disable current-state reads, in which case current_state is
	// omitted.
	Status SpendPolicyGetResponseCurrentStateStatus `json:"status" api:"required"`
	JSON   spendPolicyGetResponseCurrentStateJSON   `json:"-"`
}

// spendPolicyGetResponseCurrentStateJSON contains the JSON metadata for the struct
// [SpendPolicyGetResponseCurrentState]
type spendPolicyGetResponseCurrentStateJSON struct {
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	Enforcement        apijson.Field
	EvaluatedAt        apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyGetResponseCurrentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyGetResponseCurrentStateJSON) RawJSON() string {
	return r.raw
}

// Durable desired cap and delivery bookkeeping for this policy, or null if no
// owned effect exists. Separate from threshold evaluation: pending/leased/retry
// can lag the new billing period. Succeeded with equal generations means delivery
// completed, not an EWC readback.
type SpendPolicyGetResponseCurrentStateEnforcement struct {
	AppliedGeneration int64                                                       `json:"applied_generation" api:"required"`
	DeliveryStatus    SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus `json:"delivery_status" api:"required"`
	DesiredGeneration int64                                                       `json:"desired_generation" api:"required"`
	DesiredState      SpendPolicyGetResponseCurrentStateEnforcementDesiredState   `json:"desired_state" api:"required"`
	ExpiresAt         time.Time                                                   `json:"expires_at" api:"required,nullable" format:"date-time"`
	JSON              spendPolicyGetResponseCurrentStateEnforcementJSON           `json:"-"`
}

// spendPolicyGetResponseCurrentStateEnforcementJSON contains the JSON metadata for
// the struct [SpendPolicyGetResponseCurrentStateEnforcement]
type spendPolicyGetResponseCurrentStateEnforcementJSON struct {
	AppliedGeneration apijson.Field
	DeliveryStatus    apijson.Field
	DesiredGeneration apijson.Field
	DesiredState      apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyGetResponseCurrentStateEnforcement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyGetResponseCurrentStateEnforcementJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus string

const (
	SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusPending   SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus = "pending"
	SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusLeased    SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus = "leased"
	SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusRetry     SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus = "retry"
	SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusSucceeded SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus = "succeeded"
)

func (r SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatus) IsKnown() bool {
	switch r {
	case SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusPending, SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusLeased, SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusRetry, SpendPolicyGetResponseCurrentStateEnforcementDeliveryStatusSucceeded:
		return true
	}
	return false
}

type SpendPolicyGetResponseCurrentStateEnforcementDesiredState string

const (
	SpendPolicyGetResponseCurrentStateEnforcementDesiredStateActive   SpendPolicyGetResponseCurrentStateEnforcementDesiredState = "active"
	SpendPolicyGetResponseCurrentStateEnforcementDesiredStateInactive SpendPolicyGetResponseCurrentStateEnforcementDesiredState = "inactive"
)

func (r SpendPolicyGetResponseCurrentStateEnforcementDesiredState) IsKnown() bool {
	switch r {
	case SpendPolicyGetResponseCurrentStateEnforcementDesiredStateActive, SpendPolicyGetResponseCurrentStateEnforcementDesiredStateInactive:
		return true
	}
	return false
}

// Threshold evaluation, not proof of an edge cap. No current active period yields
// no_current_period for enabled policies; failed aggregate reads yield unavailable
// for enabled policies. Disabled policies report disabled. Neither
// no_current_period nor unavailable implies a zero balance or below-limit state.
// Operators may disable current-state reads, in which case current_state is
// omitted.
type SpendPolicyGetResponseCurrentStateStatus string

const (
	SpendPolicyGetResponseCurrentStateStatusDisabled        SpendPolicyGetResponseCurrentStateStatus = "disabled"
	SpendPolicyGetResponseCurrentStateStatusNoCurrentPeriod SpendPolicyGetResponseCurrentStateStatus = "no_current_period"
	SpendPolicyGetResponseCurrentStateStatusUnavailable     SpendPolicyGetResponseCurrentStateStatus = "unavailable"
	SpendPolicyGetResponseCurrentStateStatusBelowLimit      SpendPolicyGetResponseCurrentStateStatus = "below_limit"
	SpendPolicyGetResponseCurrentStateStatusLimitReached    SpendPolicyGetResponseCurrentStateStatus = "limit_reached"
)

func (r SpendPolicyGetResponseCurrentStateStatus) IsKnown() bool {
	switch r {
	case SpendPolicyGetResponseCurrentStateStatusDisabled, SpendPolicyGetResponseCurrentStateStatusNoCurrentPeriod, SpendPolicyGetResponseCurrentStateStatusUnavailable, SpendPolicyGetResponseCurrentStateStatusBelowLimit, SpendPolicyGetResponseCurrentStateStatusLimitReached:
		return true
	}
	return false
}

type SpendPolicyNewParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action param.Field[SpendPolicyNewParamsAction] `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Display name of the policy. After surrounding whitespace is trimmed, it must be
	// 1 to 200 bytes long. The value is stored as sent, untrimmed.
	Name param.Field[string] `json:"name" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID param.Field[SpendPolicyNewParamsProductID] `json:"product_id" api:"required"`
	// The spend threshold for the current billing period, in US cents. Must be greater
	// than zero.
	ThresholdCents param.Field[int64] `json:"threshold_cents" api:"required"`
	// Email address that receives notifications. Required when `action` is
	// `notification`. When set, it must be a bare address (`user@example.com`);
	// display-name forms such as `Name <user@example.com>` are rejected.
	NotificationEmail param.Field[string] `json:"notification_email" format:"email"`
}

func (r SpendPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyNewParamsAction string

const (
	SpendPolicyNewParamsActionNotification SpendPolicyNewParamsAction = "notification"
	SpendPolicyNewParamsActionUsageBlock   SpendPolicyNewParamsAction = "usage_block"
)

func (r SpendPolicyNewParamsAction) IsKnown() bool {
	switch r {
	case SpendPolicyNewParamsActionNotification, SpendPolicyNewParamsActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyNewParamsProductID string

const (
	SpendPolicyNewParamsProductIDWorkers        SpendPolicyNewParamsProductID = "workers"
	SpendPolicyNewParamsProductIDDurableObjects SpendPolicyNewParamsProductID = "durable_objects"
)

func (r SpendPolicyNewParamsProductID) IsKnown() bool {
	switch r {
	case SpendPolicyNewParamsProductIDWorkers, SpendPolicyNewParamsProductIDDurableObjects:
		return true
	}
	return false
}

// Represents a successful API response.
type SpendPolicyNewResponseEnvelope struct {
	// Always empty on success.
	Errors []SpendPolicyNewResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational notices about the response.
	Messages []SpendPolicyNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// A spend policy.
	Result SpendPolicyNewResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success SpendPolicyNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    spendPolicyNewResponseEnvelopeJSON    `json:"-"`
}

// spendPolicyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpendPolicyNewResponseEnvelope]
type spendPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyNewResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                    `json:"code"`
	JSON spendPolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// spendPolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpendPolicyNewResponseEnvelopeErrors]
type spendPolicyNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyNewResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                      `json:"code"`
	JSON spendPolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// spendPolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpendPolicyNewResponseEnvelopeMessages]
type spendPolicyNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type SpendPolicyNewResponseEnvelopeSuccess bool

const (
	SpendPolicyNewResponseEnvelopeSuccessTrue SpendPolicyNewResponseEnvelopeSuccess = true
)

func (r SpendPolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SpendPolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SpendPolicyUpdateParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action param.Field[SpendPolicyUpdateParamsAction] `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Display name of the policy. After surrounding whitespace is trimmed, it must be
	// 1 to 200 bytes long. The value is stored as sent, untrimmed.
	Name param.Field[string] `json:"name" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID param.Field[SpendPolicyUpdateParamsProductID] `json:"product_id" api:"required"`
	// The spend threshold for the current billing period, in US cents. Must be greater
	// than zero.
	ThresholdCents param.Field[int64]  `json:"threshold_cents" api:"required"`
	IfMatch        param.Field[string] `header:"If-Match" api:"required"`
	// Email address that receives notifications. Required when `action` is
	// `notification`. When set, it must be a bare address (`user@example.com`);
	// display-name forms such as `Name <user@example.com>` are rejected.
	NotificationEmail param.Field[string] `json:"notification_email" format:"email"`
}

func (r SpendPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyUpdateParamsAction string

const (
	SpendPolicyUpdateParamsActionNotification SpendPolicyUpdateParamsAction = "notification"
	SpendPolicyUpdateParamsActionUsageBlock   SpendPolicyUpdateParamsAction = "usage_block"
)

func (r SpendPolicyUpdateParamsAction) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateParamsActionNotification, SpendPolicyUpdateParamsActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyUpdateParamsProductID string

const (
	SpendPolicyUpdateParamsProductIDWorkers        SpendPolicyUpdateParamsProductID = "workers"
	SpendPolicyUpdateParamsProductIDDurableObjects SpendPolicyUpdateParamsProductID = "durable_objects"
)

func (r SpendPolicyUpdateParamsProductID) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateParamsProductIDWorkers, SpendPolicyUpdateParamsProductIDDurableObjects:
		return true
	}
	return false
}

// Represents a successful API response.
type SpendPolicyUpdateResponseEnvelope struct {
	// Always empty on success.
	Errors []SpendPolicyUpdateResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational notices about the response.
	Messages []SpendPolicyUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// A spend policy.
	Result SpendPolicyUpdateResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success SpendPolicyUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    spendPolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// spendPolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpendPolicyUpdateResponseEnvelope]
type spendPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyUpdateResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                       `json:"code"`
	JSON spendPolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// spendPolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpendPolicyUpdateResponseEnvelopeErrors]
type spendPolicyUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyUpdateResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                         `json:"code"`
	JSON spendPolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// spendPolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpendPolicyUpdateResponseEnvelopeMessages]
type spendPolicyUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type SpendPolicyUpdateResponseEnvelopeSuccess bool

const (
	SpendPolicyUpdateResponseEnvelopeSuccessTrue SpendPolicyUpdateResponseEnvelopeSuccess = true
)

func (r SpendPolicyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SpendPolicyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SpendPolicyListParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SpendPolicyDeleteParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	IfMatch   param.Field[string] `header:"If-Match" api:"required"`
}

// Represents a successful API response.
type SpendPolicyDeleteResponseEnvelope struct {
	// Always empty on success.
	Errors []SpendPolicyDeleteResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational notices about the response.
	Messages []SpendPolicyDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SpendPolicyDeleteResponse                   `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success SpendPolicyDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    spendPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// spendPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpendPolicyDeleteResponseEnvelope]
type spendPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyDeleteResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                       `json:"code"`
	JSON spendPolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// spendPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpendPolicyDeleteResponseEnvelopeErrors]
type spendPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyDeleteResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                         `json:"code"`
	JSON spendPolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// spendPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpendPolicyDeleteResponseEnvelopeMessages]
type spendPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type SpendPolicyDeleteResponseEnvelopeSuccess bool

const (
	SpendPolicyDeleteResponseEnvelopeSuccessTrue SpendPolicyDeleteResponseEnvelopeSuccess = true
)

func (r SpendPolicyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SpendPolicyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SpendPolicyGetParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Represents a successful API response.
type SpendPolicyGetResponseEnvelope struct {
	// Always empty on success.
	Errors []SpendPolicyGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational notices about the response.
	Messages []SpendPolicyGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A spend policy.
	Result SpendPolicyGetResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success SpendPolicyGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    spendPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// spendPolicyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpendPolicyGetResponseEnvelope]
type spendPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyGetResponseEnvelopeErrors struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                    `json:"code"`
	JSON spendPolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// spendPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpendPolicyGetResponseEnvelopeErrors]
type spendPolicyGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an API notice or error detail.
type SpendPolicyGetResponseEnvelopeMessages struct {
	// Describes the error or notice.
	Message string `json:"message" api:"required"`
	// Identifies the error or notice type.
	Code int64                                      `json:"code"`
	JSON spendPolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// spendPolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpendPolicyGetResponseEnvelopeMessages]
type spendPolicyGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type SpendPolicyGetResponseEnvelopeSuccess bool

const (
	SpendPolicyGetResponseEnvelopeSuccessTrue SpendPolicyGetResponseEnvelopeSuccess = true
)

func (r SpendPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SpendPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
