// File generated from our OpenAPI spec by Stainless.

package user

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Updates a user's subscriptions.
func (r *SubscriptionService) Update(ctx context.Context, identifier string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionUpdateResponseEnvelope
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
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *SubscriptionService) Edit(ctx context.Context, identifier string, body SubscriptionEditParams, opts ...option.RequestOption) (res *SubscriptionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of a user's subscriptions.
func (r *SubscriptionService) Get(ctx context.Context, opts ...option.RequestOption) (res *[]SubscriptionGetResponse, err error) {
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

// Union satisfied by [user.SubscriptionUpdateResponseUnknown] or
// [shared.UnionString].
type SubscriptionUpdateResponse interface {
	ImplementsUserSubscriptionUpdateResponse()
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

func (r subscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [user.SubscriptionEditResponseUnknown] or
// [shared.UnionString].
type SubscriptionEditResponse interface {
	ImplementsUserSubscriptionEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionGetResponse struct {
	// Subscription identifier tag.
	ID  string                     `json:"id"`
	App SubscriptionGetResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []SubscriptionGetResponseComponentValue `json:"component_values"`
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
	RatePlan SubscriptionGetResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionGetResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone SubscriptionGetResponseZone `json:"zone"`
	JSON subscriptionGetResponseJSON `json:"-"`
}

// subscriptionGetResponseJSON contains the JSON metadata for the struct
// [SubscriptionGetResponse]
type subscriptionGetResponseJSON struct {
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

func (r *SubscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetResponseApp struct {
	// app install id.
	InstallID string                         `json:"install_id"`
	JSON      subscriptionGetResponseAppJSON `json:"-"`
}

// subscriptionGetResponseAppJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseApp]
type subscriptionGetResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseAppJSON) RawJSON() string {
	return r.raw
}

// A component value for a subscription.
type SubscriptionGetResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                   `json:"value"`
	JSON  subscriptionGetResponseComponentValueJSON `json:"-"`
}

// subscriptionGetResponseComponentValueJSON contains the JSON metadata for the
// struct [SubscriptionGetResponseComponentValue]
type subscriptionGetResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseComponentValueJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionGetResponseFrequency string

const (
	SubscriptionGetResponseFrequencyWeekly    SubscriptionGetResponseFrequency = "weekly"
	SubscriptionGetResponseFrequencyMonthly   SubscriptionGetResponseFrequency = "monthly"
	SubscriptionGetResponseFrequencyQuarterly SubscriptionGetResponseFrequency = "quarterly"
	SubscriptionGetResponseFrequencyYearly    SubscriptionGetResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionGetResponseRatePlan struct {
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
	Sets []string                            `json:"sets"`
	JSON subscriptionGetResponseRatePlanJSON `json:"-"`
}

// subscriptionGetResponseRatePlanJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseRatePlan]
type subscriptionGetResponseRatePlanJSON struct {
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

func (r *SubscriptionGetResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseRatePlanJSON) RawJSON() string {
	return r.raw
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

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionGetResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                          `json:"name"`
	JSON subscriptionGetResponseZoneJSON `json:"-"`
}

// subscriptionGetResponseZoneJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseZone]
type subscriptionGetResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseZoneJSON) RawJSON() string {
	return r.raw
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

func (r subscriptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r subscriptionUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r subscriptionUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

type SubscriptionEditParams struct {
	App param.Field[SubscriptionEditParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionEditParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionEditParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionEditParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionEditParamsZone] `json:"zone"`
}

func (r SubscriptionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionEditParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionEditParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionEditParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionEditParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionEditParamsFrequency string

const (
	SubscriptionEditParamsFrequencyWeekly    SubscriptionEditParamsFrequency = "weekly"
	SubscriptionEditParamsFrequencyMonthly   SubscriptionEditParamsFrequency = "monthly"
	SubscriptionEditParamsFrequencyQuarterly SubscriptionEditParamsFrequency = "quarterly"
	SubscriptionEditParamsFrequencyYearly    SubscriptionEditParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionEditParamsRatePlan struct {
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

func (r SubscriptionEditParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionEditParamsZone struct {
}

func (r SubscriptionEditParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionEditResponseEnvelope struct {
	Errors   []SubscriptionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionEditResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionEditResponse                   `json:"result,required"`
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

type SubscriptionEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    subscriptionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionEditResponseEnvelopeErrors]
type subscriptionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    subscriptionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubscriptionEditResponseEnvelopeMessages]
type subscriptionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionEditResponseEnvelopeSuccess bool

const (
	SubscriptionEditResponseEnvelopeSuccessTrue SubscriptionEditResponseEnvelopeSuccess = true
)

type SubscriptionGetResponseEnvelope struct {
	Errors   []SubscriptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []SubscriptionGetResponse                 `json:"result,required,nullable"`
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

type SubscriptionGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    subscriptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionGetResponseEnvelopeErrors]
type subscriptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    subscriptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubscriptionGetResponseEnvelopeMessages]
type subscriptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionGetResponseEnvelopeSuccess bool

const (
	SubscriptionGetResponseEnvelopeSuccessTrue SubscriptionGetResponseEnvelopeSuccess = true
)

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
