// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
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

// Updates a user's subscriptions.
func (r *SubscriptionService) Update(ctx context.Context, identifier string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionUpdateResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionUpdateResponseEnvelope
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a user's subscription.
func (r *SubscriptionService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *SubscriptionService) Edit(ctx context.Context, identifier string, body SubscriptionEditParams, opts ...option.RequestOption) (res *SubscriptionEditResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionEditResponseEnvelope
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of a user's subscriptions.
func (r *SubscriptionService) Get(ctx context.Context, opts ...option.RequestOption) (res *[]Subscription, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionGetResponseEnvelope
	path := "user/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The rate plan applied to the subscription.
type RatePlan struct {
	// The ID of the rate plan.
	ID string `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency string `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged bool `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract bool `json:"is_contract"`
	// The full name of the rate plan.
	PublicName string `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope string `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets []string     `json:"sets"`
	JSON ratePlanJSON `json:"-"`
}

// ratePlanJSON contains the JSON metadata for the struct [RatePlan]
type ratePlanJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	IsContract        apijson.Field
	PublicName        apijson.Field
	Scope             apijson.Field
	Sets              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanJSON) RawJSON() string {
	return r.raw
}

// The rate plan applied to the subscription.
type RatePlanParam struct {
	// The ID of the rate plan.
	ID param.Field[string] `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency param.Field[string] `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged param.Field[bool] `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract param.Field[bool] `json:"is_contract"`
	// The full name of the rate plan.
	PublicName param.Field[string] `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope param.Field[string] `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets param.Field[[]string] `json:"sets"`
}

func (r RatePlanParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Subscription struct {
	// Subscription identifier tag.
	ID  string          `json:"id"`
	App SubscriptionApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []SubscriptionComponent `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan RatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone SubscriptionZone `json:"zone"`
	JSON subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID                 apijson.Field
	App                apijson.Field
	ComponentValues    apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	Zone               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionApp struct {
	// app install id.
	InstallID string              `json:"install_id"`
	JSON      subscriptionAppJSON `json:"-"`
}

// subscriptionAppJSON contains the JSON metadata for the struct [SubscriptionApp]
type subscriptionAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAppJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionFrequency string

const (
	SubscriptionFrequencyWeekly    SubscriptionFrequency = "weekly"
	SubscriptionFrequencyMonthly   SubscriptionFrequency = "monthly"
	SubscriptionFrequencyQuarterly SubscriptionFrequency = "quarterly"
	SubscriptionFrequencyYearly    SubscriptionFrequency = "yearly"
)

func (r SubscriptionFrequency) IsKnown() bool {
	switch r {
	case SubscriptionFrequencyWeekly, SubscriptionFrequencyMonthly, SubscriptionFrequencyQuarterly, SubscriptionFrequencyYearly:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionState string

const (
	SubscriptionStateTrial           SubscriptionState = "Trial"
	SubscriptionStateProvisioned     SubscriptionState = "Provisioned"
	SubscriptionStatePaid            SubscriptionState = "Paid"
	SubscriptionStateAwaitingPayment SubscriptionState = "AwaitingPayment"
	SubscriptionStateCancelled       SubscriptionState = "Cancelled"
	SubscriptionStateFailed          SubscriptionState = "Failed"
	SubscriptionStateExpired         SubscriptionState = "Expired"
)

func (r SubscriptionState) IsKnown() bool {
	switch r {
	case SubscriptionStateTrial, SubscriptionStateProvisioned, SubscriptionStatePaid, SubscriptionStateAwaitingPayment, SubscriptionStateCancelled, SubscriptionStateFailed, SubscriptionStateExpired:
		return true
	}
	return false
}

type SubscriptionParam struct {
	App param.Field[SubscriptionAppParam] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionComponentParam] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[RatePlanParam] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionZoneParam] `json:"zone"`
}

func (r SubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionAppParam struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionComponent struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                   `json:"value"`
	JSON  subscriptionComponentJSON `json:"-"`
}

// subscriptionComponentJSON contains the JSON metadata for the struct
// [SubscriptionComponent]
type subscriptionComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionComponentJSON) RawJSON() string {
	return r.raw
}

// A component value for a subscription.
type SubscriptionComponentParam struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionComponentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string               `json:"name"`
	JSON subscriptionZoneJSON `json:"-"`
}

// subscriptionZoneJSON contains the JSON metadata for the struct
// [SubscriptionZone]
type subscriptionZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionZoneJSON) RawJSON() string {
	return r.raw
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionZoneParam struct {
}

func (r SubscriptionZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [user.SubscriptionUpdateResponseUnknown] or
// [shared.UnionString].
type SubscriptionUpdateResponseUnion interface {
	ImplementsUserSubscriptionUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                         `json:"subscription_id"`
	JSON           subscriptionDeleteResponseJSON `json:"-"`
}

// subscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponse]
type subscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [user.SubscriptionEditResponseUnknown] or
// [shared.UnionString].
type SubscriptionEditResponseUnion interface {
	ImplementsUserSubscriptionEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionUpdateParams struct {
	Subscription SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   SubscriptionUpdateResponseUnion `json:"result,required"`
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

type SubscriptionEditParams struct {
	Subscription SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   SubscriptionEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionEditResponseEnvelopeJSON    `json:"-"`
}

// subscriptionEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionEditResponseEnvelope]
type subscriptionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionEditResponseEnvelopeSuccess bool

const (
	SubscriptionEditResponseEnvelopeSuccessTrue SubscriptionEditResponseEnvelopeSuccess = true
)

func (r SubscriptionEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []Subscription        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SubscriptionGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SubscriptionGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       subscriptionGetResponseEnvelopeJSON       `json:"-"`
}

// subscriptionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseEnvelope]
type subscriptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
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

type SubscriptionGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       subscriptionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// subscriptionGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [SubscriptionGetResponseEnvelopeResultInfo]
type subscriptionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
