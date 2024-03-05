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

// UserSubscriptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserSubscriptionService] method
// instead.
type UserSubscriptionService struct {
	Options []option.RequestOption
}

// NewUserSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSubscriptionService(opts ...option.RequestOption) (r *UserSubscriptionService) {
	r = &UserSubscriptionService{}
	r.Options = opts
	return
}

// Updates a user's subscriptions.
func (r *UserSubscriptionService) Update(ctx context.Context, identifier string, body UserSubscriptionUpdateParams, opts ...option.RequestOption) (res *UserSubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionUpdateResponseEnvelope
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a user's subscription.
func (r *UserSubscriptionService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *UserSubscriptionService) Edit(ctx context.Context, identifier string, body UserSubscriptionEditParams, opts ...option.RequestOption) (res *UserSubscriptionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of a user's subscriptions.
func (r *UserSubscriptionService) Get(ctx context.Context, opts ...option.RequestOption) (res *[]UserSubscriptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionGetResponseEnvelope
	path := "user/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserSubscriptionUpdateResponseUnknown] or
// [shared.UnionString].
type UserSubscriptionUpdateResponse interface {
	ImplementsUserSubscriptionUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSubscriptionUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserSubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                             `json:"subscription_id"`
	JSON           userSubscriptionDeleteResponseJSON `json:"-"`
}

// userSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionDeleteResponse]
type userSubscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [UserSubscriptionEditResponseUnknown] or
// [shared.UnionString].
type UserSubscriptionEditResponse interface {
	ImplementsUserSubscriptionEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSubscriptionEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserSubscriptionGetResponse struct {
	// Subscription identifier tag.
	ID  string                         `json:"id"`
	App UserSubscriptionGetResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []UserSubscriptionGetResponseComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency UserSubscriptionGetResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan UserSubscriptionGetResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State UserSubscriptionGetResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone UserSubscriptionGetResponseZone `json:"zone"`
	JSON userSubscriptionGetResponseJSON `json:"-"`
}

// userSubscriptionGetResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionGetResponse]
type userSubscriptionGetResponseJSON struct {
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

func (r *UserSubscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionGetResponseApp struct {
	// app install id.
	InstallID string                             `json:"install_id"`
	JSON      userSubscriptionGetResponseAppJSON `json:"-"`
}

// userSubscriptionGetResponseAppJSON contains the JSON metadata for the struct
// [UserSubscriptionGetResponseApp]
type userSubscriptionGetResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A component value for a subscription.
type UserSubscriptionGetResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                       `json:"value"`
	JSON  userSubscriptionGetResponseComponentValueJSON `json:"-"`
}

// userSubscriptionGetResponseComponentValueJSON contains the JSON metadata for the
// struct [UserSubscriptionGetResponseComponentValue]
type userSubscriptionGetResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionGetResponseFrequency string

const (
	UserSubscriptionGetResponseFrequencyWeekly    UserSubscriptionGetResponseFrequency = "weekly"
	UserSubscriptionGetResponseFrequencyMonthly   UserSubscriptionGetResponseFrequency = "monthly"
	UserSubscriptionGetResponseFrequencyQuarterly UserSubscriptionGetResponseFrequency = "quarterly"
	UserSubscriptionGetResponseFrequencyYearly    UserSubscriptionGetResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionGetResponseRatePlan struct {
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
	Sets []string                                `json:"sets"`
	JSON userSubscriptionGetResponseRatePlanJSON `json:"-"`
}

// userSubscriptionGetResponseRatePlanJSON contains the JSON metadata for the
// struct [UserSubscriptionGetResponseRatePlan]
type userSubscriptionGetResponseRatePlanJSON struct {
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

func (r *UserSubscriptionGetResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the subscription is in.
type UserSubscriptionGetResponseState string

const (
	UserSubscriptionGetResponseStateTrial           UserSubscriptionGetResponseState = "Trial"
	UserSubscriptionGetResponseStateProvisioned     UserSubscriptionGetResponseState = "Provisioned"
	UserSubscriptionGetResponseStatePaid            UserSubscriptionGetResponseState = "Paid"
	UserSubscriptionGetResponseStateAwaitingPayment UserSubscriptionGetResponseState = "AwaitingPayment"
	UserSubscriptionGetResponseStateCancelled       UserSubscriptionGetResponseState = "Cancelled"
	UserSubscriptionGetResponseStateFailed          UserSubscriptionGetResponseState = "Failed"
	UserSubscriptionGetResponseStateExpired         UserSubscriptionGetResponseState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionGetResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                              `json:"name"`
	JSON userSubscriptionGetResponseZoneJSON `json:"-"`
}

// userSubscriptionGetResponseZoneJSON contains the JSON metadata for the struct
// [UserSubscriptionGetResponseZone]
type userSubscriptionGetResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUpdateParams struct {
	App param.Field[UserSubscriptionUpdateParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]UserSubscriptionUpdateParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[UserSubscriptionUpdateParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[UserSubscriptionUpdateParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[UserSubscriptionUpdateParamsZone] `json:"zone"`
}

func (r UserSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionUpdateParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r UserSubscriptionUpdateParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type UserSubscriptionUpdateParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r UserSubscriptionUpdateParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionUpdateParamsFrequency string

const (
	UserSubscriptionUpdateParamsFrequencyWeekly    UserSubscriptionUpdateParamsFrequency = "weekly"
	UserSubscriptionUpdateParamsFrequencyMonthly   UserSubscriptionUpdateParamsFrequency = "monthly"
	UserSubscriptionUpdateParamsFrequencyQuarterly UserSubscriptionUpdateParamsFrequency = "quarterly"
	UserSubscriptionUpdateParamsFrequencyYearly    UserSubscriptionUpdateParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionUpdateParamsRatePlan struct {
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

func (r UserSubscriptionUpdateParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionUpdateParamsZone struct {
}

func (r UserSubscriptionUpdateParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionUpdateResponseEnvelope struct {
	Errors   []UserSubscriptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserSubscriptionUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserSubscriptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSubscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// userSubscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserSubscriptionUpdateResponseEnvelope]
type userSubscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userSubscriptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserSubscriptionUpdateResponseEnvelopeErrors]
type userSubscriptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    userSubscriptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserSubscriptionUpdateResponseEnvelopeMessages]
type userSubscriptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionUpdateResponseEnvelopeSuccess bool

const (
	UserSubscriptionUpdateResponseEnvelopeSuccessTrue UserSubscriptionUpdateResponseEnvelopeSuccess = true
)

type UserSubscriptionEditParams struct {
	App param.Field[UserSubscriptionEditParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]UserSubscriptionEditParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[UserSubscriptionEditParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[UserSubscriptionEditParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[UserSubscriptionEditParamsZone] `json:"zone"`
}

func (r UserSubscriptionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionEditParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r UserSubscriptionEditParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type UserSubscriptionEditParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r UserSubscriptionEditParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionEditParamsFrequency string

const (
	UserSubscriptionEditParamsFrequencyWeekly    UserSubscriptionEditParamsFrequency = "weekly"
	UserSubscriptionEditParamsFrequencyMonthly   UserSubscriptionEditParamsFrequency = "monthly"
	UserSubscriptionEditParamsFrequencyQuarterly UserSubscriptionEditParamsFrequency = "quarterly"
	UserSubscriptionEditParamsFrequencyYearly    UserSubscriptionEditParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionEditParamsRatePlan struct {
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

func (r UserSubscriptionEditParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionEditParamsZone struct {
}

func (r UserSubscriptionEditParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionEditResponseEnvelope struct {
	Errors   []UserSubscriptionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionEditResponseEnvelopeMessages `json:"messages,required"`
	Result   UserSubscriptionEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserSubscriptionEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSubscriptionEditResponseEnvelopeJSON    `json:"-"`
}

// userSubscriptionEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserSubscriptionEditResponseEnvelope]
type userSubscriptionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    userSubscriptionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserSubscriptionEditResponseEnvelopeErrors]
type userSubscriptionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userSubscriptionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserSubscriptionEditResponseEnvelopeMessages]
type userSubscriptionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionEditResponseEnvelopeSuccess bool

const (
	UserSubscriptionEditResponseEnvelopeSuccessTrue UserSubscriptionEditResponseEnvelopeSuccess = true
)

type UserSubscriptionGetResponseEnvelope struct {
	Errors   []UserSubscriptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserSubscriptionGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserSubscriptionGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserSubscriptionGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userSubscriptionGetResponseEnvelopeJSON       `json:"-"`
}

// userSubscriptionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserSubscriptionGetResponseEnvelope]
type userSubscriptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    userSubscriptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserSubscriptionGetResponseEnvelopeErrors]
type userSubscriptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    userSubscriptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserSubscriptionGetResponseEnvelopeMessages]
type userSubscriptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionGetResponseEnvelopeSuccess bool

const (
	UserSubscriptionGetResponseEnvelopeSuccessTrue UserSubscriptionGetResponseEnvelopeSuccess = true
)

type UserSubscriptionGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       userSubscriptionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// userSubscriptionGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [UserSubscriptionGetResponseEnvelopeResultInfo]
type userSubscriptionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
