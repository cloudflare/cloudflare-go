// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// SubscriptionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Create a zone subscription, either plan or add-ons.
func (r *SubscriptionService) New(ctx context.Context, params SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	var env SubscriptionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *SubscriptionService) Update(ctx context.Context, params SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionUpdateResponse, err error) {
	var env SubscriptionUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists zone subscription details.
func (r *SubscriptionService) Get(ctx context.Context, query SubscriptionGetParams, opts ...option.RequestOption) (res *SubscriptionGetResponse, err error) {
	var env SubscriptionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SubscriptionNewResponse struct {
	// Subscription identifier tag.
	ID string `json:"id"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionNewResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan shared.RatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionNewResponseState `json:"state"`
	JSON  subscriptionNewResponseJSON  `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	ID                 apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionNewResponseFrequency string

const (
	SubscriptionNewResponseFrequencyWeekly        SubscriptionNewResponseFrequency = "weekly"
	SubscriptionNewResponseFrequencyMonthly       SubscriptionNewResponseFrequency = "monthly"
	SubscriptionNewResponseFrequencyQuarterly     SubscriptionNewResponseFrequency = "quarterly"
	SubscriptionNewResponseFrequencyYearly        SubscriptionNewResponseFrequency = "yearly"
	SubscriptionNewResponseFrequencyNotApplicable SubscriptionNewResponseFrequency = "not-applicable"
)

func (r SubscriptionNewResponseFrequency) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseFrequencyWeekly, SubscriptionNewResponseFrequencyMonthly, SubscriptionNewResponseFrequencyQuarterly, SubscriptionNewResponseFrequencyYearly, SubscriptionNewResponseFrequencyNotApplicable:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionNewResponseState string

const (
	SubscriptionNewResponseStateTrial           SubscriptionNewResponseState = "Trial"
	SubscriptionNewResponseStateProvisioned     SubscriptionNewResponseState = "Provisioned"
	SubscriptionNewResponseStatePaid            SubscriptionNewResponseState = "Paid"
	SubscriptionNewResponseStateAwaitingPayment SubscriptionNewResponseState = "AwaitingPayment"
	SubscriptionNewResponseStateCancelled       SubscriptionNewResponseState = "Cancelled"
	SubscriptionNewResponseStateFailed          SubscriptionNewResponseState = "Failed"
	SubscriptionNewResponseStateExpired         SubscriptionNewResponseState = "Expired"
)

func (r SubscriptionNewResponseState) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseStateTrial, SubscriptionNewResponseStateProvisioned, SubscriptionNewResponseStatePaid, SubscriptionNewResponseStateAwaitingPayment, SubscriptionNewResponseStateCancelled, SubscriptionNewResponseStateFailed, SubscriptionNewResponseStateExpired:
		return true
	}
	return false
}

type SubscriptionUpdateResponse struct {
	// Subscription identifier tag.
	ID string `json:"id"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionUpdateResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan shared.RatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionUpdateResponseState `json:"state"`
	JSON  subscriptionUpdateResponseJSON  `json:"-"`
}

// subscriptionUpdateResponseJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponse]
type subscriptionUpdateResponseJSON struct {
	ID                 apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionUpdateResponseFrequency string

const (
	SubscriptionUpdateResponseFrequencyWeekly        SubscriptionUpdateResponseFrequency = "weekly"
	SubscriptionUpdateResponseFrequencyMonthly       SubscriptionUpdateResponseFrequency = "monthly"
	SubscriptionUpdateResponseFrequencyQuarterly     SubscriptionUpdateResponseFrequency = "quarterly"
	SubscriptionUpdateResponseFrequencyYearly        SubscriptionUpdateResponseFrequency = "yearly"
	SubscriptionUpdateResponseFrequencyNotApplicable SubscriptionUpdateResponseFrequency = "not-applicable"
)

func (r SubscriptionUpdateResponseFrequency) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseFrequencyWeekly, SubscriptionUpdateResponseFrequencyMonthly, SubscriptionUpdateResponseFrequencyQuarterly, SubscriptionUpdateResponseFrequencyYearly, SubscriptionUpdateResponseFrequencyNotApplicable:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionUpdateResponseState string

const (
	SubscriptionUpdateResponseStateTrial           SubscriptionUpdateResponseState = "Trial"
	SubscriptionUpdateResponseStateProvisioned     SubscriptionUpdateResponseState = "Provisioned"
	SubscriptionUpdateResponseStatePaid            SubscriptionUpdateResponseState = "Paid"
	SubscriptionUpdateResponseStateAwaitingPayment SubscriptionUpdateResponseState = "AwaitingPayment"
	SubscriptionUpdateResponseStateCancelled       SubscriptionUpdateResponseState = "Cancelled"
	SubscriptionUpdateResponseStateFailed          SubscriptionUpdateResponseState = "Failed"
	SubscriptionUpdateResponseStateExpired         SubscriptionUpdateResponseState = "Expired"
)

func (r SubscriptionUpdateResponseState) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseStateTrial, SubscriptionUpdateResponseStateProvisioned, SubscriptionUpdateResponseStatePaid, SubscriptionUpdateResponseStateAwaitingPayment, SubscriptionUpdateResponseStateCancelled, SubscriptionUpdateResponseStateFailed, SubscriptionUpdateResponseStateExpired:
		return true
	}
	return false
}

type SubscriptionGetResponse struct {
	// Subscription identifier tag.
	ID string `json:"id"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionGetResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan shared.RatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionGetResponseState `json:"state"`
	JSON  subscriptionGetResponseJSON  `json:"-"`
}

// subscriptionGetResponseJSON contains the JSON metadata for the struct
// [SubscriptionGetResponse]
type subscriptionGetResponseJSON struct {
	ID                 apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionGetResponseFrequency string

const (
	SubscriptionGetResponseFrequencyWeekly        SubscriptionGetResponseFrequency = "weekly"
	SubscriptionGetResponseFrequencyMonthly       SubscriptionGetResponseFrequency = "monthly"
	SubscriptionGetResponseFrequencyQuarterly     SubscriptionGetResponseFrequency = "quarterly"
	SubscriptionGetResponseFrequencyYearly        SubscriptionGetResponseFrequency = "yearly"
	SubscriptionGetResponseFrequencyNotApplicable SubscriptionGetResponseFrequency = "not-applicable"
)

func (r SubscriptionGetResponseFrequency) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseFrequencyWeekly, SubscriptionGetResponseFrequencyMonthly, SubscriptionGetResponseFrequencyQuarterly, SubscriptionGetResponseFrequencyYearly, SubscriptionGetResponseFrequencyNotApplicable:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionGetResponseState string

const (
	SubscriptionGetResponseStateTrial           SubscriptionGetResponseState = "Trial"
	SubscriptionGetResponseStateProvisioned     SubscriptionGetResponseState = "Provisioned"
	SubscriptionGetResponseStatePaid            SubscriptionGetResponseState = "Paid"
	SubscriptionGetResponseStateAwaitingPayment SubscriptionGetResponseState = "AwaitingPayment"
	SubscriptionGetResponseStateCancelled       SubscriptionGetResponseState = "Cancelled"
	SubscriptionGetResponseStateFailed          SubscriptionGetResponseState = "Failed"
	SubscriptionGetResponseStateExpired         SubscriptionGetResponseState = "Expired"
)

func (r SubscriptionGetResponseState) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseStateTrial, SubscriptionGetResponseStateProvisioned, SubscriptionGetResponseStatePaid, SubscriptionGetResponseStateAwaitingPayment, SubscriptionGetResponseStateCancelled, SubscriptionGetResponseStateFailed, SubscriptionGetResponseStateExpired:
		return true
	}
	return false
}

type SubscriptionNewParams struct {
	// Identifier
	ZoneID       param.Field[string]      `path:"zone_id,required"`
	Subscription shared.SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   SubscriptionNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionNewResponseEnvelopeJSON    `json:"-"`
}

// subscriptionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseEnvelope]
type subscriptionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionNewResponseEnvelopeSuccess bool

const (
	SubscriptionNewResponseEnvelopeSuccessTrue SubscriptionNewResponseEnvelopeSuccess = true
)

func (r SubscriptionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionUpdateParams struct {
	// Identifier
	ZoneID       param.Field[string]      `path:"zone_id,required"`
	Subscription shared.SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   SubscriptionUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// subscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponseEnvelope]
type subscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

func (r SubscriptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SubscriptionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   SubscriptionGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionGetResponseEnvelopeJSON    `json:"-"`
}

// subscriptionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseEnvelope]
type subscriptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionGetResponseEnvelopeSuccess bool

const (
	SubscriptionGetResponseEnvelopeSuccessTrue SubscriptionGetResponseEnvelopeSuccess = true
)

func (r SubscriptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
