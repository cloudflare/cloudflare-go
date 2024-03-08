// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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

// Create a zone subscription, either plan or add-ons.
func (r *SubscriptionService) New(ctx context.Context, identifier string, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of an account's subscriptions.
func (r *SubscriptionService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]SubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/subscriptions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists zone subscription details.
func (r *SubscriptionService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SubscriptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SubscriptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zones.SubscriptionNewResponseUnknown] or
// [shared.UnionString].
type SubscriptionNewResponse interface {
	ImplementsZonesSubscriptionNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionListResponse struct {
	// Subscription identifier tag.
	ID  string                      `json:"id"`
	App SubscriptionListResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []SubscriptionListResponseComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionListResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan SubscriptionListResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionListResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone SubscriptionListResponseZone `json:"zone"`
	JSON subscriptionListResponseJSON `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
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

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponseApp struct {
	// app install id.
	InstallID string                          `json:"install_id"`
	JSON      subscriptionListResponseAppJSON `json:"-"`
}

// subscriptionListResponseAppJSON contains the JSON metadata for the struct
// [SubscriptionListResponseApp]
type subscriptionListResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseAppJSON) RawJSON() string {
	return r.raw
}

// A component value for a subscription.
type SubscriptionListResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                    `json:"value"`
	JSON  subscriptionListResponseComponentValueJSON `json:"-"`
}

// subscriptionListResponseComponentValueJSON contains the JSON metadata for the
// struct [SubscriptionListResponseComponentValue]
type subscriptionListResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseComponentValueJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionListResponseFrequency string

const (
	SubscriptionListResponseFrequencyWeekly    SubscriptionListResponseFrequency = "weekly"
	SubscriptionListResponseFrequencyMonthly   SubscriptionListResponseFrequency = "monthly"
	SubscriptionListResponseFrequencyQuarterly SubscriptionListResponseFrequency = "quarterly"
	SubscriptionListResponseFrequencyYearly    SubscriptionListResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionListResponseRatePlan struct {
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
	Sets []string                             `json:"sets"`
	JSON subscriptionListResponseRatePlanJSON `json:"-"`
}

// subscriptionListResponseRatePlanJSON contains the JSON metadata for the struct
// [SubscriptionListResponseRatePlan]
type subscriptionListResponseRatePlanJSON struct {
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

func (r *SubscriptionListResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseRatePlanJSON) RawJSON() string {
	return r.raw
}

// The state that the subscription is in.
type SubscriptionListResponseState string

const (
	SubscriptionListResponseStateTrial           SubscriptionListResponseState = "Trial"
	SubscriptionListResponseStateProvisioned     SubscriptionListResponseState = "Provisioned"
	SubscriptionListResponseStatePaid            SubscriptionListResponseState = "Paid"
	SubscriptionListResponseStateAwaitingPayment SubscriptionListResponseState = "AwaitingPayment"
	SubscriptionListResponseStateCancelled       SubscriptionListResponseState = "Cancelled"
	SubscriptionListResponseStateFailed          SubscriptionListResponseState = "Failed"
	SubscriptionListResponseStateExpired         SubscriptionListResponseState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionListResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                           `json:"name"`
	JSON subscriptionListResponseZoneJSON `json:"-"`
}

// subscriptionListResponseZoneJSON contains the JSON metadata for the struct
// [SubscriptionListResponseZone]
type subscriptionListResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseZoneJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zones.SubscriptionGetResponseUnknown] or
// [shared.UnionString].
type SubscriptionGetResponse interface {
	ImplementsZonesSubscriptionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SubscriptionNewParams struct {
	App param.Field[SubscriptionNewParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionNewParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionNewParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionNewParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionNewParamsZone] `json:"zone"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionNewParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionNewParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionNewParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionNewParamsFrequency string

const (
	SubscriptionNewParamsFrequencyWeekly    SubscriptionNewParamsFrequency = "weekly"
	SubscriptionNewParamsFrequencyMonthly   SubscriptionNewParamsFrequency = "monthly"
	SubscriptionNewParamsFrequencyQuarterly SubscriptionNewParamsFrequency = "quarterly"
	SubscriptionNewParamsFrequencyYearly    SubscriptionNewParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type SubscriptionNewParamsRatePlan struct {
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

func (r SubscriptionNewParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionNewParamsZone struct {
}

func (r SubscriptionNewParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewResponseEnvelope struct {
	Errors   []SubscriptionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionNewResponse                   `json:"result,required"`
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

type SubscriptionNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    subscriptionNewResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionNewResponseEnvelopeErrors]
type subscriptionNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    subscriptionNewResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubscriptionNewResponseEnvelopeMessages]
type subscriptionNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionNewResponseEnvelopeSuccess bool

const (
	SubscriptionNewResponseEnvelopeSuccessTrue SubscriptionNewResponseEnvelopeSuccess = true
)

type SubscriptionListResponseEnvelope struct {
	Errors   []SubscriptionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SubscriptionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SubscriptionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SubscriptionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       subscriptionListResponseEnvelopeJSON       `json:"-"`
}

// subscriptionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionListResponseEnvelope]
type subscriptionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    subscriptionListResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionListResponseEnvelopeErrors]
type subscriptionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    subscriptionListResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubscriptionListResponseEnvelopeMessages]
type subscriptionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionListResponseEnvelopeSuccess bool

const (
	SubscriptionListResponseEnvelopeSuccessTrue SubscriptionListResponseEnvelopeSuccess = true
)

type SubscriptionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       subscriptionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// subscriptionListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [SubscriptionListResponseEnvelopeResultInfo]
type subscriptionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetResponseEnvelope struct {
	Errors   []SubscriptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SubscriptionGetResponse                   `json:"result,required"`
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
