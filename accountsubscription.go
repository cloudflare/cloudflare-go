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

// AccountSubscriptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountSubscriptionService]
// method instead.
type AccountSubscriptionService struct {
	Options []option.RequestOption
}

// NewAccountSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSubscriptionService(opts ...option.RequestOption) (r *AccountSubscriptionService) {
	r = &AccountSubscriptionService{}
	r.Options = opts
	return
}

// Updates an account subscription.
func (r *AccountSubscriptionService) Update(ctx context.Context, accountIdentifier string, subscriptionIdentifier string, body AccountSubscriptionUpdateParams, opts ...option.RequestOption) (res *AccountSubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", accountIdentifier, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an account's subscription.
func (r *AccountSubscriptionService) Delete(ctx context.Context, accountIdentifier string, subscriptionIdentifier string, opts ...option.RequestOption) (res *AccountSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", accountIdentifier, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates an account subscription.
func (r *AccountSubscriptionService) AccountSubscriptionsNewSubscription(ctx context.Context, accountIdentifier string, body AccountSubscriptionAccountSubscriptionsNewSubscriptionParams, opts ...option.RequestOption) (res *AccountSubscriptionAccountSubscriptionsNewSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/subscriptions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all of an account's subscriptions.
func (r *AccountSubscriptionService) AccountSubscriptionsListSubscriptions(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/subscriptions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountSubscriptionUpdateResponse struct {
	Errors   []AccountSubscriptionUpdateResponseError   `json:"errors"`
	Messages []AccountSubscriptionUpdateResponseMessage `json:"messages"`
	Result   interface{}                                `json:"result"`
	// Whether the API call was successful
	Success AccountSubscriptionUpdateResponseSuccess `json:"success"`
	JSON    accountSubscriptionUpdateResponseJSON    `json:"-"`
}

// accountSubscriptionUpdateResponseJSON contains the JSON metadata for the struct
// [AccountSubscriptionUpdateResponse]
type accountSubscriptionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountSubscriptionUpdateResponseErrorJSON `json:"-"`
}

// accountSubscriptionUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountSubscriptionUpdateResponseError]
type accountSubscriptionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountSubscriptionUpdateResponseMessageJSON `json:"-"`
}

// accountSubscriptionUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountSubscriptionUpdateResponseMessage]
type accountSubscriptionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSubscriptionUpdateResponseSuccess bool

const (
	AccountSubscriptionUpdateResponseSuccessTrue AccountSubscriptionUpdateResponseSuccess = true
)

type AccountSubscriptionDeleteResponse struct {
	Errors   []AccountSubscriptionDeleteResponseError   `json:"errors"`
	Messages []AccountSubscriptionDeleteResponseMessage `json:"messages"`
	Result   AccountSubscriptionDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSubscriptionDeleteResponseSuccess `json:"success"`
	JSON    accountSubscriptionDeleteResponseJSON    `json:"-"`
}

// accountSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [AccountSubscriptionDeleteResponse]
type accountSubscriptionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountSubscriptionDeleteResponseErrorJSON `json:"-"`
}

// accountSubscriptionDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountSubscriptionDeleteResponseError]
type accountSubscriptionDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountSubscriptionDeleteResponseMessageJSON `json:"-"`
}

// accountSubscriptionDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountSubscriptionDeleteResponseMessage]
type accountSubscriptionDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionDeleteResponseResult struct {
	// Subscription identifier tag.
	SubscriptionID string                                      `json:"subscription_id"`
	JSON           accountSubscriptionDeleteResponseResultJSON `json:"-"`
}

// accountSubscriptionDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountSubscriptionDeleteResponseResult]
type accountSubscriptionDeleteResponseResultJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountSubscriptionDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSubscriptionDeleteResponseSuccess bool

const (
	AccountSubscriptionDeleteResponseSuccessTrue AccountSubscriptionDeleteResponseSuccess = true
)

type AccountSubscriptionAccountSubscriptionsNewSubscriptionResponse struct {
	Errors   []AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseError   `json:"errors"`
	Messages []AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessage `json:"messages"`
	Result   interface{}                                                             `json:"result"`
	// Whether the API call was successful
	Success AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseSuccess `json:"success"`
	JSON    accountSubscriptionAccountSubscriptionsNewSubscriptionResponseJSON    `json:"-"`
}

// accountSubscriptionAccountSubscriptionsNewSubscriptionResponseJSON contains the
// JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsNewSubscriptionResponse]
type accountSubscriptionAccountSubscriptionsNewSubscriptionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsNewSubscriptionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountSubscriptionAccountSubscriptionsNewSubscriptionResponseErrorJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsNewSubscriptionResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseError]
type accountSubscriptionAccountSubscriptionsNewSubscriptionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessageJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessage]
type accountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseSuccess bool

const (
	AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseSuccessTrue AccountSubscriptionAccountSubscriptionsNewSubscriptionResponseSuccess = true
)

type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponse struct {
	Errors     []AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseError    `json:"errors"`
	Messages   []AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessage  `json:"messages"`
	Result     []AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResult   `json:"result"`
	ResultInfo AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseSuccess `json:"success"`
	JSON    accountSubscriptionAccountSubscriptionsListSubscriptionsResponseJSON    `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseJSON contains
// the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponse]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountSubscriptionAccountSubscriptionsListSubscriptionsResponseErrorJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseError]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessageJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessage]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResult struct {
	// Subscription identifier tag.
	ID  string                                                                    `json:"id"`
	App AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZone `json:"zone"`
	JSON accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResult]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultJSON struct {
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

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultApp struct {
	// app install id.
	InstallID string                                                                        `json:"install_id"`
	JSON      accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultAppJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultAppJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultApp]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A component value for a subscription.
type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                                                                  `json:"value"`
	JSON  accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValueJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValueJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValue]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription is renewed automatically.
type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequency string

const (
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequencyWeekly    AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequency = "weekly"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequencyMonthly   AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequency = "monthly"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequencyQuarterly AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequency = "quarterly"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequencyYearly    AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultFrequency = "yearly"
)

// The rate plan applied to the subscription.
type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlan struct {
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
	Sets []string                                                                           `json:"sets"`
	JSON accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlanJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlanJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlan]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlanJSON struct {
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

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the subscription is in.
type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState string

const (
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStateTrial           AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "Trial"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStateProvisioned     AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "Provisioned"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStatePaid            AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "Paid"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStateAwaitingPayment AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "AwaitingPayment"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStateCancelled       AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "Cancelled"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStateFailed          AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "Failed"
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultStateExpired         AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                                                                         `json:"name"`
	JSON accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZoneJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZoneJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZone]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                        `json:"total_count"`
	JSON       accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfoJSON `json:"-"`
}

// accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfo]
type accountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseSuccess bool

const (
	AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseSuccessTrue AccountSubscriptionAccountSubscriptionsListSubscriptionsResponseSuccess = true
)

type AccountSubscriptionUpdateParams struct {
	App param.Field[AccountSubscriptionUpdateParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]AccountSubscriptionUpdateParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[AccountSubscriptionUpdateParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[AccountSubscriptionUpdateParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[AccountSubscriptionUpdateParamsZone] `json:"zone"`
}

func (r AccountSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSubscriptionUpdateParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r AccountSubscriptionUpdateParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type AccountSubscriptionUpdateParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r AccountSubscriptionUpdateParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type AccountSubscriptionUpdateParamsFrequency string

const (
	AccountSubscriptionUpdateParamsFrequencyWeekly    AccountSubscriptionUpdateParamsFrequency = "weekly"
	AccountSubscriptionUpdateParamsFrequencyMonthly   AccountSubscriptionUpdateParamsFrequency = "monthly"
	AccountSubscriptionUpdateParamsFrequencyQuarterly AccountSubscriptionUpdateParamsFrequency = "quarterly"
	AccountSubscriptionUpdateParamsFrequencyYearly    AccountSubscriptionUpdateParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type AccountSubscriptionUpdateParamsRatePlan struct {
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

func (r AccountSubscriptionUpdateParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type AccountSubscriptionUpdateParamsZone struct {
}

func (r AccountSubscriptionUpdateParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSubscriptionAccountSubscriptionsNewSubscriptionParams struct {
	App param.Field[AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsZone] `json:"zone"`
}

func (r AccountSubscriptionAccountSubscriptionsNewSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency string

const (
	AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyWeekly    AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "weekly"
	AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyMonthly   AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "monthly"
	AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyQuarterly AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "quarterly"
	AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequencyYearly    AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsRatePlan struct {
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

func (r AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsZone struct {
}

func (r AccountSubscriptionAccountSubscriptionsNewSubscriptionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
