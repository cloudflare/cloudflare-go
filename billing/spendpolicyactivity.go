// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

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

// SpendPolicyActivityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendPolicyActivityService] method instead.
type SpendPolicyActivityService struct {
	Options []option.RequestOption
}

// NewSpendPolicyActivityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpendPolicyActivityService(opts ...option.RequestOption) (r *SpendPolicyActivityService) {
	r = &SpendPolicyActivityService{}
	r.Options = opts
	return
}

// Returns configuration history from immutable policy revisions for the account,
// including deleted policies. Entries contain before/after values and verified
// actor attribution. Pages use descending revision IDs, not event or commit order;
// refresh the first page for new entries. This cursor is not an incremental
// consumer checkpoint. Defaults to `period=current_billing_cycle`: each revision's
// timestamp must be within its product's current subscription cycle (inclusive
// start, exclusive end). Each product's cycle is resolved independently from
// canonical aggregates, including for disabled or deleted policies. Set
// `period=all` for history across cycles. Filtering happens before pagination. No
// current cycles returns an empty page; unavailable or disabled cycle resolution
// returns `503` with code `1013`. Cursors pin resolved cycle bounds across pages
// and must be used with the same period selection. Cursors obtained before this
// filter was introduced must be refreshed.
func (r *SpendPolicyActivityService) List(ctx context.Context, params SpendPolicyActivityListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[SpendPolicyActivityListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies/activity", params.AccountID)
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

// Returns configuration history from immutable policy revisions for the account,
// including deleted policies. Entries contain before/after values and verified
// actor attribution. Pages use descending revision IDs, not event or commit order;
// refresh the first page for new entries. This cursor is not an incremental
// consumer checkpoint. Defaults to `period=current_billing_cycle`: each revision's
// timestamp must be within its product's current subscription cycle (inclusive
// start, exclusive end). Each product's cycle is resolved independently from
// canonical aggregates, including for disabled or deleted policies. Set
// `period=all` for history across cycles. Filtering happens before pagination. No
// current cycles returns an empty page; unavailable or disabled cycle resolution
// returns `503` with code `1013`. Cursors pin resolved cycle bounds across pages
// and must be used with the same period selection. Cursors obtained before this
// filter was introduced must be refreshed.
func (r *SpendPolicyActivityService) ListAutoPaging(ctx context.Context, params SpendPolicyActivityListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[SpendPolicyActivityListResponse] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, params, opts...))
}

type SpendPolicyActivityListResponse struct {
	ID     string                                `json:"id" api:"required"`
	Action SpendPolicyActivityListResponseAction `json:"action" api:"required"`
	// Verified actor identifier, unknown when absent, or trusted internal actor.
	Actor      string                                 `json:"actor" api:"required"`
	Details    SpendPolicyActivityListResponseDetails `json:"details" api:"required"`
	OccurredAt time.Time                              `json:"occurred_at" api:"required" format:"date-time"`
	// The unique identifier of the spend policy.
	PolicyID      string                              `json:"policy_id" api:"required" format:"uuid"`
	PolicyVersion int64                               `json:"policy_version" api:"required"`
	ProductID     string                              `json:"product_id" api:"required"`
	JSON          spendPolicyActivityListResponseJSON `json:"-"`
}

// spendPolicyActivityListResponseJSON contains the JSON metadata for the struct
// [SpendPolicyActivityListResponse]
type spendPolicyActivityListResponseJSON struct {
	ID            apijson.Field
	Action        apijson.Field
	Actor         apijson.Field
	Details       apijson.Field
	OccurredAt    apijson.Field
	PolicyID      apijson.Field
	PolicyVersion apijson.Field
	ProductID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SpendPolicyActivityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyActivityListResponseAction string

const (
	SpendPolicyActivityListResponseActionCreated  SpendPolicyActivityListResponseAction = "created"
	SpendPolicyActivityListResponseActionUpdated  SpendPolicyActivityListResponseAction = "updated"
	SpendPolicyActivityListResponseActionEnabled  SpendPolicyActivityListResponseAction = "enabled"
	SpendPolicyActivityListResponseActionDisabled SpendPolicyActivityListResponseAction = "disabled"
	SpendPolicyActivityListResponseActionDeleted  SpendPolicyActivityListResponseAction = "deleted"
)

func (r SpendPolicyActivityListResponseAction) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseActionCreated, SpendPolicyActivityListResponseActionUpdated, SpendPolicyActivityListResponseActionEnabled, SpendPolicyActivityListResponseActionDisabled, SpendPolicyActivityListResponseActionDeleted:
		return true
	}
	return false
}

type SpendPolicyActivityListResponseDetails struct {
	// A spend policy.
	After SpendPolicyActivityListResponseDetailsAfter `json:"after"`
	// Prior configuration; null on creation or if a previous revision is unavailable.
	Before SpendPolicyActivityListResponseDetailsBefore `json:"before" api:"nullable"`
	JSON   spendPolicyActivityListResponseDetailsJSON   `json:"-"`
}

// spendPolicyActivityListResponseDetailsJSON contains the JSON metadata for the
// struct [SpendPolicyActivityListResponseDetails]
type spendPolicyActivityListResponseDetailsJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpendPolicyActivityListResponseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsJSON) RawJSON() string {
	return r.raw
}

// A spend policy.
type SpendPolicyActivityListResponseDetailsAfter struct {
	// The unique identifier of the spend policy.
	ID string `json:"id" api:"required" format:"uuid"`
	// Internal numeric ID of the account that owns the policy. This is not the account
	// tag used in the path.
	AccountID int64 `json:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyActivityListResponseDetailsAfterAction `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled bool `json:"enabled" api:"required"`
	// Display name of the policy.
	Name string `json:"name" api:"required"`
	// Email address that receives notifications. An empty string when none is set.
	NotificationEmail string `json:"notification_email" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID SpendPolicyActivityListResponseDetailsAfterProductID `json:"product_id" api:"required"`
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
	CurrentState SpendPolicyActivityListResponseDetailsAfterCurrentState `json:"current_state"`
	JSON         spendPolicyActivityListResponseDetailsAfterJSON         `json:"-"`
}

// spendPolicyActivityListResponseDetailsAfterJSON contains the JSON metadata for
// the struct [SpendPolicyActivityListResponseDetailsAfter]
type spendPolicyActivityListResponseDetailsAfterJSON struct {
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

func (r *SpendPolicyActivityListResponseDetailsAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsAfterJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyActivityListResponseDetailsAfterAction string

const (
	SpendPolicyActivityListResponseDetailsAfterActionNotification SpendPolicyActivityListResponseDetailsAfterAction = "notification"
	SpendPolicyActivityListResponseDetailsAfterActionUsageBlock   SpendPolicyActivityListResponseDetailsAfterAction = "usage_block"
)

func (r SpendPolicyActivityListResponseDetailsAfterAction) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsAfterActionNotification, SpendPolicyActivityListResponseDetailsAfterActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyActivityListResponseDetailsAfterProductID string

const (
	SpendPolicyActivityListResponseDetailsAfterProductIDWorkers        SpendPolicyActivityListResponseDetailsAfterProductID = "workers"
	SpendPolicyActivityListResponseDetailsAfterProductIDDurableObjects SpendPolicyActivityListResponseDetailsAfterProductID = "durable_objects"
)

func (r SpendPolicyActivityListResponseDetailsAfterProductID) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsAfterProductIDWorkers, SpendPolicyActivityListResponseDetailsAfterProductIDDurableObjects:
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
type SpendPolicyActivityListResponseDetailsAfterCurrentState struct {
	BalanceCents       int64     `json:"balance_cents" api:"required,nullable"`
	BillingPeriodEnd   time.Time `json:"billing_period_end" api:"required,nullable" format:"date-time"`
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required,nullable" format:"date-time"`
	// Durable desired cap and delivery bookkeeping for this policy, or null if no
	// owned effect exists. Separate from threshold evaluation: pending/leased/retry
	// can lag the new billing period. Succeeded with equal generations means delivery
	// completed, not an EWC readback.
	Enforcement SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcement `json:"enforcement" api:"required,nullable"`
	// Evaluation time; not the timestamp of the latest usage event.
	EvaluatedAt time.Time `json:"evaluated_at" api:"required" format:"date-time"`
	// Threshold evaluation, not proof of an edge cap. No current active period yields
	// no_current_period for enabled policies; failed aggregate reads yield unavailable
	// for enabled policies. Disabled policies report disabled. Neither
	// no_current_period nor unavailable implies a zero balance or below-limit state.
	// Operators may disable current-state reads, in which case current_state is
	// omitted.
	Status SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus `json:"status" api:"required"`
	JSON   spendPolicyActivityListResponseDetailsAfterCurrentStateJSON   `json:"-"`
}

// spendPolicyActivityListResponseDetailsAfterCurrentStateJSON contains the JSON
// metadata for the struct
// [SpendPolicyActivityListResponseDetailsAfterCurrentState]
type spendPolicyActivityListResponseDetailsAfterCurrentStateJSON struct {
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	Enforcement        apijson.Field
	EvaluatedAt        apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyActivityListResponseDetailsAfterCurrentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsAfterCurrentStateJSON) RawJSON() string {
	return r.raw
}

// Durable desired cap and delivery bookkeeping for this policy, or null if no
// owned effect exists. Separate from threshold evaluation: pending/leased/retry
// can lag the new billing period. Succeeded with equal generations means delivery
// completed, not an EWC readback.
type SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcement struct {
	AppliedGeneration int64                                                                            `json:"applied_generation" api:"required"`
	DeliveryStatus    SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus `json:"delivery_status" api:"required"`
	DesiredGeneration int64                                                                            `json:"desired_generation" api:"required"`
	DesiredState      SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredState   `json:"desired_state" api:"required"`
	ExpiresAt         time.Time                                                                        `json:"expires_at" api:"required,nullable" format:"date-time"`
	JSON              spendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementJSON           `json:"-"`
}

// spendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementJSON contains
// the JSON metadata for the struct
// [SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcement]
type spendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementJSON struct {
	AppliedGeneration apijson.Field
	DeliveryStatus    apijson.Field
	DesiredGeneration apijson.Field
	DesiredState      apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus string

const (
	SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusPending   SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus = "pending"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusLeased    SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus = "leased"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusRetry     SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus = "retry"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusSucceeded SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus = "succeeded"
)

func (r SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatus) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusPending, SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusLeased, SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusRetry, SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDeliveryStatusSucceeded:
		return true
	}
	return false
}

type SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredState string

const (
	SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredStateActive   SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredState = "active"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredStateInactive SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredState = "inactive"
)

func (r SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredState) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredStateActive, SpendPolicyActivityListResponseDetailsAfterCurrentStateEnforcementDesiredStateInactive:
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
type SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus string

const (
	SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusDisabled        SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus = "disabled"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusNoCurrentPeriod SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus = "no_current_period"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusUnavailable     SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus = "unavailable"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusBelowLimit      SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus = "below_limit"
	SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusLimitReached    SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus = "limit_reached"
)

func (r SpendPolicyActivityListResponseDetailsAfterCurrentStateStatus) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusDisabled, SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusNoCurrentPeriod, SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusUnavailable, SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusBelowLimit, SpendPolicyActivityListResponseDetailsAfterCurrentStateStatusLimitReached:
		return true
	}
	return false
}

// Prior configuration; null on creation or if a previous revision is unavailable.
type SpendPolicyActivityListResponseDetailsBefore struct {
	// The unique identifier of the spend policy.
	ID string `json:"id" api:"required" format:"uuid"`
	// Internal numeric ID of the account that owns the policy. This is not the account
	// tag used in the path.
	AccountID int64 `json:"account_id" api:"required"`
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyActivityListResponseDetailsBeforeAction `json:"action" api:"required"`
	// Whether the policy is evaluated.
	Enabled bool `json:"enabled" api:"required"`
	// Display name of the policy.
	Name string `json:"name" api:"required"`
	// Email address that receives notifications. An empty string when none is set.
	NotificationEmail string `json:"notification_email" api:"required"`
	// The product the policy applies to: `workers` or `durable_objects`. Each
	// product's balance and policies are separate. Both products support
	// `notification` and `usage_block` policies on create and update.
	ProductID SpendPolicyActivityListResponseDetailsBeforeProductID `json:"product_id" api:"required"`
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
	CurrentState SpendPolicyActivityListResponseDetailsBeforeCurrentState `json:"current_state"`
	JSON         spendPolicyActivityListResponseDetailsBeforeJSON         `json:"-"`
}

// spendPolicyActivityListResponseDetailsBeforeJSON contains the JSON metadata for
// the struct [SpendPolicyActivityListResponseDetailsBefore]
type spendPolicyActivityListResponseDetailsBeforeJSON struct {
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

func (r *SpendPolicyActivityListResponseDetailsBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsBeforeJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyActivityListResponseDetailsBeforeAction string

const (
	SpendPolicyActivityListResponseDetailsBeforeActionNotification SpendPolicyActivityListResponseDetailsBeforeAction = "notification"
	SpendPolicyActivityListResponseDetailsBeforeActionUsageBlock   SpendPolicyActivityListResponseDetailsBeforeAction = "usage_block"
)

func (r SpendPolicyActivityListResponseDetailsBeforeAction) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsBeforeActionNotification, SpendPolicyActivityListResponseDetailsBeforeActionUsageBlock:
		return true
	}
	return false
}

// The product the policy applies to: `workers` or `durable_objects`. Each
// product's balance and policies are separate. Both products support
// `notification` and `usage_block` policies on create and update.
type SpendPolicyActivityListResponseDetailsBeforeProductID string

const (
	SpendPolicyActivityListResponseDetailsBeforeProductIDWorkers        SpendPolicyActivityListResponseDetailsBeforeProductID = "workers"
	SpendPolicyActivityListResponseDetailsBeforeProductIDDurableObjects SpendPolicyActivityListResponseDetailsBeforeProductID = "durable_objects"
)

func (r SpendPolicyActivityListResponseDetailsBeforeProductID) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsBeforeProductIDWorkers, SpendPolicyActivityListResponseDetailsBeforeProductIDDurableObjects:
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
type SpendPolicyActivityListResponseDetailsBeforeCurrentState struct {
	BalanceCents       int64     `json:"balance_cents" api:"required,nullable"`
	BillingPeriodEnd   time.Time `json:"billing_period_end" api:"required,nullable" format:"date-time"`
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required,nullable" format:"date-time"`
	// Durable desired cap and delivery bookkeeping for this policy, or null if no
	// owned effect exists. Separate from threshold evaluation: pending/leased/retry
	// can lag the new billing period. Succeeded with equal generations means delivery
	// completed, not an EWC readback.
	Enforcement SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcement `json:"enforcement" api:"required,nullable"`
	// Evaluation time; not the timestamp of the latest usage event.
	EvaluatedAt time.Time `json:"evaluated_at" api:"required" format:"date-time"`
	// Threshold evaluation, not proof of an edge cap. No current active period yields
	// no_current_period for enabled policies; failed aggregate reads yield unavailable
	// for enabled policies. Disabled policies report disabled. Neither
	// no_current_period nor unavailable implies a zero balance or below-limit state.
	// Operators may disable current-state reads, in which case current_state is
	// omitted.
	Status SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus `json:"status" api:"required"`
	JSON   spendPolicyActivityListResponseDetailsBeforeCurrentStateJSON   `json:"-"`
}

// spendPolicyActivityListResponseDetailsBeforeCurrentStateJSON contains the JSON
// metadata for the struct
// [SpendPolicyActivityListResponseDetailsBeforeCurrentState]
type spendPolicyActivityListResponseDetailsBeforeCurrentStateJSON struct {
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	Enforcement        apijson.Field
	EvaluatedAt        apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyActivityListResponseDetailsBeforeCurrentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsBeforeCurrentStateJSON) RawJSON() string {
	return r.raw
}

// Durable desired cap and delivery bookkeeping for this policy, or null if no
// owned effect exists. Separate from threshold evaluation: pending/leased/retry
// can lag the new billing period. Succeeded with equal generations means delivery
// completed, not an EWC readback.
type SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcement struct {
	AppliedGeneration int64                                                                             `json:"applied_generation" api:"required"`
	DeliveryStatus    SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus `json:"delivery_status" api:"required"`
	DesiredGeneration int64                                                                             `json:"desired_generation" api:"required"`
	DesiredState      SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredState   `json:"desired_state" api:"required"`
	ExpiresAt         time.Time                                                                         `json:"expires_at" api:"required,nullable" format:"date-time"`
	JSON              spendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementJSON           `json:"-"`
}

// spendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementJSON contains
// the JSON metadata for the struct
// [SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcement]
type spendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementJSON struct {
	AppliedGeneration apijson.Field
	DeliveryStatus    apijson.Field
	DesiredGeneration apijson.Field
	DesiredState      apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementJSON) RawJSON() string {
	return r.raw
}

type SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus string

const (
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusPending   SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus = "pending"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusLeased    SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus = "leased"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusRetry     SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus = "retry"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusSucceeded SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus = "succeeded"
)

func (r SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatus) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusPending, SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusLeased, SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusRetry, SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDeliveryStatusSucceeded:
		return true
	}
	return false
}

type SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredState string

const (
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredStateActive   SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredState = "active"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredStateInactive SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredState = "inactive"
)

func (r SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredState) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredStateActive, SpendPolicyActivityListResponseDetailsBeforeCurrentStateEnforcementDesiredStateInactive:
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
type SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus string

const (
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusDisabled        SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus = "disabled"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusNoCurrentPeriod SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus = "no_current_period"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusUnavailable     SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus = "unavailable"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusBelowLimit      SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus = "below_limit"
	SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusLimitReached    SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus = "limit_reached"
)

func (r SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatus) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusDisabled, SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusNoCurrentPeriod, SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusUnavailable, SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusBelowLimit, SpendPolicyActivityListResponseDetailsBeforeCurrentStateStatusLimitReached:
		return true
	}
	return false
}

type SpendPolicyActivityListParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque token from the preceding response's result_info.cursor. Pass it unchanged
	// to retrieve older entries.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of audit entries to return. Defaults to 50. Values above 100 are
	// capped at 100 rather than rejected.
	Limit param.Field[int64] `query:"limit"`
	// Select current subscription billing-cycle history or history across all cycles.
	Period param.Field[SpendPolicyActivityListParamsPeriod] `query:"period"`
}

// URLQuery serializes [SpendPolicyActivityListParams]'s query parameters as
// `url.Values`.
func (r SpendPolicyActivityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Select current subscription billing-cycle history or history across all cycles.
type SpendPolicyActivityListParamsPeriod string

const (
	SpendPolicyActivityListParamsPeriodCurrentBillingCycle SpendPolicyActivityListParamsPeriod = "current_billing_cycle"
	SpendPolicyActivityListParamsPeriodAll                 SpendPolicyActivityListParamsPeriod = "all"
)

func (r SpendPolicyActivityListParamsPeriod) IsKnown() bool {
	switch r {
	case SpendPolicyActivityListParamsPeriodCurrentBillingCycle, SpendPolicyActivityListParamsPeriodAll:
		return true
	}
	return false
}
