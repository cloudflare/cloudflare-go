// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// SubscriptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSubscriptionService] method
// instead.
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

// Updates an account subscription.
func (r *SubscriptionService) Update(ctx context.Context, accountIdentifier string, subscriptionIdentifier string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", accountIdentifier, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an account's subscription.
func (r *SubscriptionService) Delete(ctx context.Context, accountIdentifier string, subscriptionIdentifier string, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", accountIdentifier, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates an account subscription.
func (r *SubscriptionService) AccountSubscriptionsNewSubscription(ctx context.Context, accountIdentifier string, body SubscriptionAccountSubscriptionsNewSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionAccountSubscriptionsNewSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelope
	path := fmt.Sprintf("accounts/%s/subscriptions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of an account's subscriptions.
func (r *SubscriptionService) AccountSubscriptionsListSubscriptions(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]SubscriptionAccountSubscriptionsListSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/subscriptions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a zone subscription, either plan or add-ons.
func (r *SubscriptionService) ZoneSubscriptionNewZoneSubscription(ctx context.Context, identifier string, body SubscriptionZoneSubscriptionNewZoneSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionZoneSubscriptionNewZoneSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *SubscriptionService) ZoneSubscriptionUpdateZoneSubscription(ctx context.Context, identifier string, body SubscriptionZoneSubscriptionUpdateZoneSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists zone subscription details.
func (r *SubscriptionService) ZoneSubscriptionZoneSubscriptionDetails(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [SubscriptionUpdateResponseUnknown] or [shared.UnionString].
type SubscriptionUpdateResponse interface {
	ImplementsSubscriptionUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionUpdateResponse)(nil)).Elem(),
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

// Union satisfied by
// [SubscriptionAccountSubscriptionsNewSubscriptionResponseUnknown] or
// [shared.UnionString].
type SubscriptionAccountSubscriptionsNewSubscriptionResponse interface {
	ImplementsSubscriptionAccountSubscriptionsNewSubscriptionResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionAccountSubscriptionsNewSubscriptionResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionAccountSubscriptionsListSubscriptionsResponse struct {
	// Subscription identifier tag.
	ID  string                                                       `json:"id"`
	App SubscriptionAccountSubscriptionsListSubscriptionsResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []SubscriptionAccountSubscriptionsListSubscriptionsResponseComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan SubscriptionAccountSubscriptionsListSubscriptionsResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionAccountSubscriptionsListSubscriptionsResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone SubscriptionAccountSubscriptionsListSubscriptionsResponseZone `json:"zone"`
	JSON subscriptionAccountSubscriptionsListSubscriptionsResponseJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseJSON contains the JSON
// metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponse]
type subscriptionAccountSubscriptionsListSubscriptionsResponseJSON struct {
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

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionAccountSubscriptionsListSubscriptionsResponseApp struct {
	// app install id.
	InstallID string                                                           `json:"install_id"`
	JSON      subscriptionAccountSubscriptionsListSubscriptionsResponseAppJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseAppJSON contains the
// JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseApp]
type subscriptionAccountSubscriptionsListSubscriptionsResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A component value for a subscription.
type SubscriptionAccountSubscriptionsListSubscriptionsResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                                                     `json:"value"`
	JSON  subscriptionAccountSubscriptionsListSubscriptionsResponseComponentValueJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseComponentValueJSON
// contains the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseComponentValue]
type subscriptionAccountSubscriptionsListSubscriptionsResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription is renewed automatically.
type SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequency string

const (
	SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequencyWeekly    SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequency = "weekly"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequencyMonthly   SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequency = "monthly"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequencyQuarterly SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequency = "quarterly"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequencyYearly    SubscriptionAccountSubscriptionsListSubscriptionsResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionAccountSubscriptionsListSubscriptionsResponseRatePlan struct {
	// The ID of the rate plan.
	ID interface{} `json:"id"`
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
	Sets []string                                                              `json:"sets"`
	JSON subscriptionAccountSubscriptionsListSubscriptionsResponseRatePlanJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseRatePlanJSON contains
// the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseRatePlan]
type subscriptionAccountSubscriptionsListSubscriptionsResponseRatePlanJSON struct {
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

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the subscription is in.
type SubscriptionAccountSubscriptionsListSubscriptionsResponseState string

const (
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStateTrial           SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "Trial"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStateProvisioned     SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "Provisioned"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStatePaid            SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "Paid"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStateAwaitingPayment SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "AwaitingPayment"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStateCancelled       SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "Cancelled"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStateFailed          SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "Failed"
	SubscriptionAccountSubscriptionsListSubscriptionsResponseStateExpired         SubscriptionAccountSubscriptionsListSubscriptionsResponseState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionAccountSubscriptionsListSubscriptionsResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                                                            `json:"name"`
	JSON subscriptionAccountSubscriptionsListSubscriptionsResponseZoneJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseZoneJSON contains the
// JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseZone]
type subscriptionAccountSubscriptionsListSubscriptionsResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [SubscriptionZoneSubscriptionNewZoneSubscriptionResponseUnknown] or
// [shared.UnionString].
type SubscriptionZoneSubscriptionNewZoneSubscriptionResponse interface {
	ImplementsSubscriptionZoneSubscriptionNewZoneSubscriptionResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionZoneSubscriptionNewZoneSubscriptionResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseUnknown] or
// [shared.UnionString].
type SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse interface {
	ImplementsSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseUnknown] or
// [shared.UnionString].
type SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse interface {
	ImplementsSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionUpdateParams struct {
	App param.Field[SubscriptionUpdateParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionUpdateParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionUpdateParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionUpdateParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionUpdateParamsZone] `json:"zone"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionUpdateParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionUpdateParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionUpdateParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionUpdateParamsFrequency string

const (
	SubscriptionUpdateParamsFrequencyWeekly    SubscriptionUpdateParamsFrequency = "weekly"
	SubscriptionUpdateParamsFrequencyMonthly   SubscriptionUpdateParamsFrequency = "monthly"
	SubscriptionUpdateParamsFrequencyQuarterly SubscriptionUpdateParamsFrequency = "quarterly"
	SubscriptionUpdateParamsFrequencyYearly    SubscriptionUpdateParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionUpdateParamsRatePlan struct {
	// The ID of the rate plan.
	ID param.Field[interface{}] `json:"id"`
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

func (r SubscriptionUpdateParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionUpdateParamsZone struct {
}

func (r SubscriptionUpdateParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []SubscriptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionUpdateResponse                   `json:"result,required"`
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

type SubscriptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    subscriptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionUpdateResponseEnvelopeErrors]
type subscriptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    subscriptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SubscriptionUpdateResponseEnvelopeMessages]
type subscriptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

type SubscriptionDeleteResponseEnvelope struct {
	Errors   []SubscriptionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionDeleteResponseEnvelopeJSON    `json:"-"`
}

// subscriptionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponseEnvelope]
type subscriptionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    subscriptionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionDeleteResponseEnvelopeErrors]
type subscriptionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    subscriptionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SubscriptionDeleteResponseEnvelopeMessages]
type subscriptionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionDeleteResponseEnvelopeSuccess bool

const (
	SubscriptionDeleteResponseEnvelopeSuccessTrue SubscriptionDeleteResponseEnvelopeSuccess = true
)

type SubscriptionAccountSubscriptionsNewSubscriptionParams struct {
	App param.Field[SubscriptionAccountSubscriptionsNewSubscriptionParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionAccountSubscriptionsNewSubscriptionParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionAccountSubscriptionsNewSubscriptionParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionAccountSubscriptionsNewSubscriptionParamsZone] `json:"zone"`
}

func (r SubscriptionAccountSubscriptionsNewSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionAccountSubscriptionsNewSubscriptionParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionAccountSubscriptionsNewSubscriptionParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionAccountSubscriptionsNewSubscriptionParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionAccountSubscriptionsNewSubscriptionParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency string

const (
	SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyWeekly    SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "weekly"
	SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyMonthly   SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "monthly"
	SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyQuarterly SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "quarterly"
	SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyYearly    SubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionAccountSubscriptionsNewSubscriptionParamsRatePlan struct {
	// The ID of the rate plan.
	ID param.Field[interface{}] `json:"id"`
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

func (r SubscriptionAccountSubscriptionsNewSubscriptionParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionAccountSubscriptionsNewSubscriptionParamsZone struct {
}

func (r SubscriptionAccountSubscriptionsNewSubscriptionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelope struct {
	Errors   []SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionAccountSubscriptionsNewSubscriptionResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeJSON    `json:"-"`
}

// subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelope]
type subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrors]
type subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessages]
type subscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeSuccess bool

const (
	SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeSuccessTrue SubscriptionAccountSubscriptionsNewSubscriptionResponseEnvelopeSuccess = true
)

type SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelope struct {
	Errors   []SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessages `json:"messages,required"`
	Result   []SubscriptionAccountSubscriptionsListSubscriptionsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeJSON       `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelope]
type subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrors]
type subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessages]
type subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeSuccess bool

const (
	SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeSuccessTrue SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeSuccess = true
)

type SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfoJSON `json:"-"`
}

// subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfo]
type subscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAccountSubscriptionsListSubscriptionsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionNewZoneSubscriptionParams struct {
	App param.Field[SubscriptionZoneSubscriptionNewZoneSubscriptionParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionZoneSubscriptionNewZoneSubscriptionParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionZoneSubscriptionNewZoneSubscriptionParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionZoneSubscriptionNewZoneSubscriptionParamsZone] `json:"zone"`
}

func (r SubscriptionZoneSubscriptionNewZoneSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionZoneSubscriptionNewZoneSubscriptionParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionZoneSubscriptionNewZoneSubscriptionParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionZoneSubscriptionNewZoneSubscriptionParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionZoneSubscriptionNewZoneSubscriptionParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency string

const (
	SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyWeekly    SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "weekly"
	SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyMonthly   SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "monthly"
	SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyQuarterly SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "quarterly"
	SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyYearly    SubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionZoneSubscriptionNewZoneSubscriptionParamsRatePlan struct {
	// The ID of the rate plan.
	ID param.Field[interface{}] `json:"id"`
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

func (r SubscriptionZoneSubscriptionNewZoneSubscriptionParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionZoneSubscriptionNewZoneSubscriptionParamsZone struct {
}

func (r SubscriptionZoneSubscriptionNewZoneSubscriptionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelope struct {
	Errors   []SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionZoneSubscriptionNewZoneSubscriptionResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeJSON    `json:"-"`
}

// subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelope]
type subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrors]
type subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessages]
type subscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeSuccess bool

const (
	SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeSuccessTrue SubscriptionZoneSubscriptionNewZoneSubscriptionResponseEnvelopeSuccess = true
)

type SubscriptionZoneSubscriptionUpdateZoneSubscriptionParams struct {
	App param.Field[SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsZone] `json:"zone"`
}

func (r SubscriptionZoneSubscriptionUpdateZoneSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency string

const (
	SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyWeekly    SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "weekly"
	SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyMonthly   SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "monthly"
	SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyQuarterly SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "quarterly"
	SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyYearly    SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsRatePlan struct {
	// The ID of the rate plan.
	ID param.Field[interface{}] `json:"id"`
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

func (r SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsZone struct {
}

func (r SubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelope struct {
	Errors   []SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeJSON    `json:"-"`
}

// subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelope]
type subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrors struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrors]
type subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessages struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessages]
type subscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeSuccess bool

const (
	SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeSuccessTrue SubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseEnvelopeSuccess = true
)

type SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelope struct {
	Errors   []SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeJSON    `json:"-"`
}

// subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelope]
type subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrors]
type subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessages]
type subscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeSuccess bool

const (
	SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeSuccessTrue SubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseEnvelopeSuccess = true
)
