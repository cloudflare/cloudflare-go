// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *UserSubscriptionService) Update(ctx context.Context, identifier string, body UserSubscriptionUpdateParams, opts ...option.RequestOption) (res *UserSubscriptionResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a user's subscription.
func (r *UserSubscriptionService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists all of a user's subscriptions.
func (r *UserSubscriptionService) UserSubscriptionGetUserSubscriptions(ctx context.Context, opts ...option.RequestOption) (res *UserSubscriptionResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserSubscriptionResponseCollection struct {
	Errors     []UserSubscriptionResponseCollectionError    `json:"errors"`
	Messages   []UserSubscriptionResponseCollectionMessage  `json:"messages"`
	Result     []UserSubscriptionResponseCollectionResult   `json:"result"`
	ResultInfo UserSubscriptionResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserSubscriptionResponseCollectionSuccess `json:"success"`
	JSON    userSubscriptionResponseCollectionJSON    `json:"-"`
}

// userSubscriptionResponseCollectionJSON contains the JSON metadata for the struct
// [UserSubscriptionResponseCollection]
type userSubscriptionResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseCollectionError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    userSubscriptionResponseCollectionErrorJSON `json:"-"`
}

// userSubscriptionResponseCollectionErrorJSON contains the JSON metadata for the
// struct [UserSubscriptionResponseCollectionError]
type userSubscriptionResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseCollectionMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    userSubscriptionResponseCollectionMessageJSON `json:"-"`
}

// userSubscriptionResponseCollectionMessageJSON contains the JSON metadata for the
// struct [UserSubscriptionResponseCollectionMessage]
type userSubscriptionResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseCollectionResult struct {
	// Subscription identifier tag.
	ID  string                                      `json:"id"`
	App UserSubscriptionResponseCollectionResultApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []UserSubscriptionResponseCollectionResultComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency UserSubscriptionResponseCollectionResultFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan UserSubscriptionResponseCollectionResultRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State UserSubscriptionResponseCollectionResultState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone UserSubscriptionResponseCollectionResultZone `json:"zone"`
	JSON userSubscriptionResponseCollectionResultJSON `json:"-"`
}

// userSubscriptionResponseCollectionResultJSON contains the JSON metadata for the
// struct [UserSubscriptionResponseCollectionResult]
type userSubscriptionResponseCollectionResultJSON struct {
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

func (r *UserSubscriptionResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseCollectionResultApp struct {
	// app install id.
	InstallID string                                          `json:"install_id"`
	JSON      userSubscriptionResponseCollectionResultAppJSON `json:"-"`
}

// userSubscriptionResponseCollectionResultAppJSON contains the JSON metadata for
// the struct [UserSubscriptionResponseCollectionResultApp]
type userSubscriptionResponseCollectionResultAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollectionResultApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A component value for a subscription.
type UserSubscriptionResponseCollectionResultComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                                    `json:"value"`
	JSON  userSubscriptionResponseCollectionResultComponentValueJSON `json:"-"`
}

// userSubscriptionResponseCollectionResultComponentValueJSON contains the JSON
// metadata for the struct [UserSubscriptionResponseCollectionResultComponentValue]
type userSubscriptionResponseCollectionResultComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollectionResultComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionResponseCollectionResultFrequency string

const (
	UserSubscriptionResponseCollectionResultFrequencyWeekly    UserSubscriptionResponseCollectionResultFrequency = "weekly"
	UserSubscriptionResponseCollectionResultFrequencyMonthly   UserSubscriptionResponseCollectionResultFrequency = "monthly"
	UserSubscriptionResponseCollectionResultFrequencyQuarterly UserSubscriptionResponseCollectionResultFrequency = "quarterly"
	UserSubscriptionResponseCollectionResultFrequencyYearly    UserSubscriptionResponseCollectionResultFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionResponseCollectionResultRatePlan struct {
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
	Sets []string                                             `json:"sets"`
	JSON userSubscriptionResponseCollectionResultRatePlanJSON `json:"-"`
}

// userSubscriptionResponseCollectionResultRatePlanJSON contains the JSON metadata
// for the struct [UserSubscriptionResponseCollectionResultRatePlan]
type userSubscriptionResponseCollectionResultRatePlanJSON struct {
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

func (r *UserSubscriptionResponseCollectionResultRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the subscription is in.
type UserSubscriptionResponseCollectionResultState string

const (
	UserSubscriptionResponseCollectionResultStateTrial           UserSubscriptionResponseCollectionResultState = "Trial"
	UserSubscriptionResponseCollectionResultStateProvisioned     UserSubscriptionResponseCollectionResultState = "Provisioned"
	UserSubscriptionResponseCollectionResultStatePaid            UserSubscriptionResponseCollectionResultState = "Paid"
	UserSubscriptionResponseCollectionResultStateAwaitingPayment UserSubscriptionResponseCollectionResultState = "AwaitingPayment"
	UserSubscriptionResponseCollectionResultStateCancelled       UserSubscriptionResponseCollectionResultState = "Cancelled"
	UserSubscriptionResponseCollectionResultStateFailed          UserSubscriptionResponseCollectionResultState = "Failed"
	UserSubscriptionResponseCollectionResultStateExpired         UserSubscriptionResponseCollectionResultState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionResponseCollectionResultZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                                           `json:"name"`
	JSON userSubscriptionResponseCollectionResultZoneJSON `json:"-"`
}

// userSubscriptionResponseCollectionResultZoneJSON contains the JSON metadata for
// the struct [UserSubscriptionResponseCollectionResultZone]
type userSubscriptionResponseCollectionResultZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollectionResultZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       userSubscriptionResponseCollectionResultInfoJSON `json:"-"`
}

// userSubscriptionResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [UserSubscriptionResponseCollectionResultInfo]
type userSubscriptionResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionResponseCollectionSuccess bool

const (
	UserSubscriptionResponseCollectionSuccessTrue UserSubscriptionResponseCollectionSuccess = true
)

type UserSubscriptionResponseSingle struct {
	Errors   []UserSubscriptionResponseSingleError   `json:"errors"`
	Messages []UserSubscriptionResponseSingleMessage `json:"messages"`
	Result   interface{}                             `json:"result"`
	// Whether the API call was successful
	Success UserSubscriptionResponseSingleSuccess `json:"success"`
	JSON    userSubscriptionResponseSingleJSON    `json:"-"`
}

// userSubscriptionResponseSingleJSON contains the JSON metadata for the struct
// [UserSubscriptionResponseSingle]
type userSubscriptionResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseSingleError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    userSubscriptionResponseSingleErrorJSON `json:"-"`
}

// userSubscriptionResponseSingleErrorJSON contains the JSON metadata for the
// struct [UserSubscriptionResponseSingleError]
type userSubscriptionResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionResponseSingleMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    userSubscriptionResponseSingleMessageJSON `json:"-"`
}

// userSubscriptionResponseSingleMessageJSON contains the JSON metadata for the
// struct [UserSubscriptionResponseSingleMessage]
type userSubscriptionResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionResponseSingleSuccess bool

const (
	UserSubscriptionResponseSingleSuccessTrue UserSubscriptionResponseSingleSuccess = true
)

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
