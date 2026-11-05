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

// SpendPolicyAuditService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendPolicyAuditService] method instead.
type SpendPolicyAuditService struct {
	Options []option.RequestOption
}

// NewSpendPolicyAuditService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpendPolicyAuditService(opts ...option.RequestOption) (r *SpendPolicyAuditService) {
	r = &SpendPolicyAuditService{}
	r.Options = opts
	return
}

// Returns the policy's evaluation history, newest first. At most one
// threshold-crossing entry (`limit_reached`) is recorded per policy version and
// billing period, so a policy update can record a new entry without any change in
// balance. Disabled policies still return their history; deleted policies return
// `404`. There is no pagination: use `limit` to set how many entries come back.
// Defaults to `period=current_billing_cycle`, resolved from the product's
// canonical aggregates, not calendar months. Crossings match their recorded
// billing-period bounds, even if recorded late. Set `period=all` for history
// across cycles. Filtering happens before LIMIT. No current cycle returns an empty
// array; unavailable or disabled cycle resolution returns `503` with code `1013`,
// never an empty success.
func (r *SpendPolicyAuditService) List(ctx context.Context, policyID string, params SpendPolicyAuditListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SpendPolicyAuditListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/spend-policies/%s/audit", params.AccountID, policyID)
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

// Returns the policy's evaluation history, newest first. At most one
// threshold-crossing entry (`limit_reached`) is recorded per policy version and
// billing period, so a policy update can record a new entry without any change in
// balance. Disabled policies still return their history; deleted policies return
// `404`. There is no pagination: use `limit` to set how many entries come back.
// Defaults to `period=current_billing_cycle`, resolved from the product's
// canonical aggregates, not calendar months. Crossings match their recorded
// billing-period bounds, even if recorded late. Set `period=all` for history
// across cycles. Filtering happens before LIMIT. No current cycle returns an empty
// array; unavailable or disabled cycle resolution returns `503` with code `1013`,
// never an empty success.
func (r *SpendPolicyAuditService) ListAutoPaging(ctx context.Context, policyID string, params SpendPolicyAuditListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SpendPolicyAuditListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, policyID, params, opts...))
}

// A state transition recorded when the policy was evaluated against the account's
// balance.
type SpendPolicyAuditListResponse struct {
	// The action taken when the threshold is reached. `notification` sends an email to
	// `notification_email`. `usage_block` blocks further usage of the product until
	// the end of the billing period. Raising the threshold above the current balance,
	// disabling the policy, or deleting it lifts the block early.
	Action SpendPolicyAuditListResponseAction `json:"action" api:"required"`
	// The account's balance for the product at evaluation time, in US cents.
	BalanceCents int64 `json:"balance_cents" api:"required"`
	// Exclusive end of the billing period that was evaluated.
	BillingPeriodEnd time.Time `json:"billing_period_end" api:"required" format:"date-time"`
	// Inclusive start of the billing period that was evaluated.
	BillingPeriodStart time.Time `json:"billing_period_start" api:"required" format:"date-time"`
	// When the entry was recorded.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The policy `version` that was evaluated.
	PolicyVersion int64 `json:"policy_version" api:"required"`
	// A human-readable explanation of the trigger.
	Reason string `json:"reason" api:"required"`
	// What caused the evaluation. `kafka` means a balance update. `policy_revision`
	// means a change to the policy.
	Source SpendPolicyAuditListResponseSource `json:"source" api:"required"`
	// The policy threshold at evaluation time, in US cents.
	ThresholdCents int64 `json:"threshold_cents" api:"required"`
	// The trigger recorded at evaluation time. New evaluations record `limit_reached`
	// when the balance reaches or exceeds the configured threshold. Earlier entries
	// retain their recorded trigger.
	Trigger string                           `json:"trigger" api:"required"`
	JSON    spendPolicyAuditListResponseJSON `json:"-"`
}

// spendPolicyAuditListResponseJSON contains the JSON metadata for the struct
// [SpendPolicyAuditListResponse]
type spendPolicyAuditListResponseJSON struct {
	Action             apijson.Field
	BalanceCents       apijson.Field
	BillingPeriodEnd   apijson.Field
	BillingPeriodStart apijson.Field
	CreatedAt          apijson.Field
	PolicyVersion      apijson.Field
	Reason             apijson.Field
	Source             apijson.Field
	ThresholdCents     apijson.Field
	Trigger            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendPolicyAuditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendPolicyAuditListResponseJSON) RawJSON() string {
	return r.raw
}

// The action taken when the threshold is reached. `notification` sends an email to
// `notification_email`. `usage_block` blocks further usage of the product until
// the end of the billing period. Raising the threshold above the current balance,
// disabling the policy, or deleting it lifts the block early.
type SpendPolicyAuditListResponseAction string

const (
	SpendPolicyAuditListResponseActionNotification SpendPolicyAuditListResponseAction = "notification"
	SpendPolicyAuditListResponseActionUsageBlock   SpendPolicyAuditListResponseAction = "usage_block"
)

func (r SpendPolicyAuditListResponseAction) IsKnown() bool {
	switch r {
	case SpendPolicyAuditListResponseActionNotification, SpendPolicyAuditListResponseActionUsageBlock:
		return true
	}
	return false
}

// What caused the evaluation. `kafka` means a balance update. `policy_revision`
// means a change to the policy.
type SpendPolicyAuditListResponseSource string

const (
	SpendPolicyAuditListResponseSourceKafka          SpendPolicyAuditListResponseSource = "kafka"
	SpendPolicyAuditListResponseSourcePolicyRevision SpendPolicyAuditListResponseSource = "policy_revision"
)

func (r SpendPolicyAuditListResponseSource) IsKnown() bool {
	switch r {
	case SpendPolicyAuditListResponseSourceKafka, SpendPolicyAuditListResponseSourcePolicyRevision:
		return true
	}
	return false
}

type SpendPolicyAuditListParams struct {
	// Represents a Cloudflare resource identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Maximum number of audit entries to return. Defaults to 50. Values above 100 are
	// capped at 100 rather than rejected.
	Limit param.Field[int64] `query:"limit"`
	// Select current subscription billing-cycle history or history across all cycles.
	Period param.Field[SpendPolicyAuditListParamsPeriod] `query:"period"`
}

// URLQuery serializes [SpendPolicyAuditListParams]'s query parameters as
// `url.Values`.
func (r SpendPolicyAuditListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Select current subscription billing-cycle history or history across all cycles.
type SpendPolicyAuditListParamsPeriod string

const (
	SpendPolicyAuditListParamsPeriodCurrentBillingCycle SpendPolicyAuditListParamsPeriod = "current_billing_cycle"
	SpendPolicyAuditListParamsPeriodAll                 SpendPolicyAuditListParamsPeriod = "all"
)

func (r SpendPolicyAuditListParamsPeriod) IsKnown() bool {
	switch r {
	case SpendPolicyAuditListParamsPeriodCurrentBillingCycle, SpendPolicyAuditListParamsPeriodAll:
		return true
	}
	return false
}
