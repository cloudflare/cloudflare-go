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

// ZoneSubscriptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSubscriptionService] method
// instead.
type ZoneSubscriptionService struct {
	Options []option.RequestOption
}

// NewZoneSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSubscriptionService(opts ...option.RequestOption) (r *ZoneSubscriptionService) {
	r = &ZoneSubscriptionService{}
	r.Options = opts
	return
}

// Create a zone subscription, either plan or add-ons.
func (r *ZoneSubscriptionService) New(ctx context.Context, identifier string, body ZoneSubscriptionNewParams, opts ...option.RequestOption) (res *ZoneSubscriptionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSubscriptionNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of an account's subscriptions.
func (r *ZoneSubscriptionService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]ZoneSubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSubscriptionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/subscriptions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists zone subscription details.
func (r *ZoneSubscriptionService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneSubscriptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSubscriptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZoneSubscriptionNewResponseUnknown] or [shared.UnionString].
type ZoneSubscriptionNewResponse interface {
	ImplementsZoneSubscriptionNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneSubscriptionNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneSubscriptionListResponse struct {
	// Subscription identifier tag.
	ID  string                          `json:"id"`
	App ZoneSubscriptionListResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []ZoneSubscriptionListResponseComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency ZoneSubscriptionListResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan ZoneSubscriptionListResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State ZoneSubscriptionListResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone ZoneSubscriptionListResponseZone `json:"zone"`
	JSON zoneSubscriptionListResponseJSON `json:"-"`
}

// zoneSubscriptionListResponseJSON contains the JSON metadata for the struct
// [ZoneSubscriptionListResponse]
type zoneSubscriptionListResponseJSON struct {
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

func (r *ZoneSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionListResponseApp struct {
	// app install id.
	InstallID string                              `json:"install_id"`
	JSON      zoneSubscriptionListResponseAppJSON `json:"-"`
}

// zoneSubscriptionListResponseAppJSON contains the JSON metadata for the struct
// [ZoneSubscriptionListResponseApp]
type zoneSubscriptionListResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseAppJSON) RawJSON() string {
	return r.raw
}

// A component value for a subscription.
type ZoneSubscriptionListResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                        `json:"value"`
	JSON  zoneSubscriptionListResponseComponentValueJSON `json:"-"`
}

// zoneSubscriptionListResponseComponentValueJSON contains the JSON metadata for
// the struct [ZoneSubscriptionListResponseComponentValue]
type zoneSubscriptionListResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseComponentValueJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type ZoneSubscriptionListResponseFrequency string

const (
	ZoneSubscriptionListResponseFrequencyWeekly    ZoneSubscriptionListResponseFrequency = "weekly"
	ZoneSubscriptionListResponseFrequencyMonthly   ZoneSubscriptionListResponseFrequency = "monthly"
	ZoneSubscriptionListResponseFrequencyQuarterly ZoneSubscriptionListResponseFrequency = "quarterly"
	ZoneSubscriptionListResponseFrequencyYearly    ZoneSubscriptionListResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type ZoneSubscriptionListResponseRatePlan struct {
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
	Sets []string                                 `json:"sets"`
	JSON zoneSubscriptionListResponseRatePlanJSON `json:"-"`
}

// zoneSubscriptionListResponseRatePlanJSON contains the JSON metadata for the
// struct [ZoneSubscriptionListResponseRatePlan]
type zoneSubscriptionListResponseRatePlanJSON struct {
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

func (r *ZoneSubscriptionListResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseRatePlanJSON) RawJSON() string {
	return r.raw
}

// The state that the subscription is in.
type ZoneSubscriptionListResponseState string

const (
	ZoneSubscriptionListResponseStateTrial           ZoneSubscriptionListResponseState = "Trial"
	ZoneSubscriptionListResponseStateProvisioned     ZoneSubscriptionListResponseState = "Provisioned"
	ZoneSubscriptionListResponseStatePaid            ZoneSubscriptionListResponseState = "Paid"
	ZoneSubscriptionListResponseStateAwaitingPayment ZoneSubscriptionListResponseState = "AwaitingPayment"
	ZoneSubscriptionListResponseStateCancelled       ZoneSubscriptionListResponseState = "Cancelled"
	ZoneSubscriptionListResponseStateFailed          ZoneSubscriptionListResponseState = "Failed"
	ZoneSubscriptionListResponseStateExpired         ZoneSubscriptionListResponseState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type ZoneSubscriptionListResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                               `json:"name"`
	JSON zoneSubscriptionListResponseZoneJSON `json:"-"`
}

// zoneSubscriptionListResponseZoneJSON contains the JSON metadata for the struct
// [ZoneSubscriptionListResponseZone]
type zoneSubscriptionListResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseZoneJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZoneSubscriptionGetResponseUnknown] or [shared.UnionString].
type ZoneSubscriptionGetResponse interface {
	ImplementsZoneSubscriptionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneSubscriptionGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneSubscriptionNewParams struct {
	App param.Field[ZoneSubscriptionNewParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]ZoneSubscriptionNewParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[ZoneSubscriptionNewParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[ZoneSubscriptionNewParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[ZoneSubscriptionNewParamsZone] `json:"zone"`
}

func (r ZoneSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSubscriptionNewParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r ZoneSubscriptionNewParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type ZoneSubscriptionNewParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r ZoneSubscriptionNewParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type ZoneSubscriptionNewParamsFrequency string

const (
	ZoneSubscriptionNewParamsFrequencyWeekly    ZoneSubscriptionNewParamsFrequency = "weekly"
	ZoneSubscriptionNewParamsFrequencyMonthly   ZoneSubscriptionNewParamsFrequency = "monthly"
	ZoneSubscriptionNewParamsFrequencyQuarterly ZoneSubscriptionNewParamsFrequency = "quarterly"
	ZoneSubscriptionNewParamsFrequencyYearly    ZoneSubscriptionNewParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type ZoneSubscriptionNewParamsRatePlan struct {
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

func (r ZoneSubscriptionNewParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type ZoneSubscriptionNewParamsZone struct {
}

func (r ZoneSubscriptionNewParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSubscriptionNewResponseEnvelope struct {
	Errors   []ZoneSubscriptionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSubscriptionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneSubscriptionNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneSubscriptionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneSubscriptionNewResponseEnvelopeJSON    `json:"-"`
}

// zoneSubscriptionNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSubscriptionNewResponseEnvelope]
type zoneSubscriptionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSubscriptionNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSubscriptionNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSubscriptionNewResponseEnvelopeErrors]
type zoneSubscriptionNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSubscriptionNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSubscriptionNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSubscriptionNewResponseEnvelopeMessages]
type zoneSubscriptionNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSubscriptionNewResponseEnvelopeSuccess bool

const (
	ZoneSubscriptionNewResponseEnvelopeSuccessTrue ZoneSubscriptionNewResponseEnvelopeSuccess = true
)

type ZoneSubscriptionListResponseEnvelope struct {
	Errors   []ZoneSubscriptionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSubscriptionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZoneSubscriptionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZoneSubscriptionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZoneSubscriptionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zoneSubscriptionListResponseEnvelopeJSON       `json:"-"`
}

// zoneSubscriptionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSubscriptionListResponseEnvelope]
type zoneSubscriptionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSubscriptionListResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSubscriptionListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSubscriptionListResponseEnvelopeErrors]
type zoneSubscriptionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSubscriptionListResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSubscriptionListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSubscriptionListResponseEnvelopeMessages]
type zoneSubscriptionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSubscriptionListResponseEnvelopeSuccess bool

const (
	ZoneSubscriptionListResponseEnvelopeSuccessTrue ZoneSubscriptionListResponseEnvelopeSuccess = true
)

type ZoneSubscriptionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       zoneSubscriptionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zoneSubscriptionListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZoneSubscriptionListResponseEnvelopeResultInfo]
type zoneSubscriptionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionGetResponseEnvelope struct {
	Errors   []ZoneSubscriptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSubscriptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneSubscriptionGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneSubscriptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneSubscriptionGetResponseEnvelopeJSON    `json:"-"`
}

// zoneSubscriptionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSubscriptionGetResponseEnvelope]
type zoneSubscriptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSubscriptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSubscriptionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSubscriptionGetResponseEnvelopeErrors]
type zoneSubscriptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSubscriptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSubscriptionGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSubscriptionGetResponseEnvelopeMessages]
type zoneSubscriptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSubscriptionGetResponseEnvelopeSuccess bool

const (
	ZoneSubscriptionGetResponseEnvelopeSuccessTrue ZoneSubscriptionGetResponseEnvelopeSuccess = true
)
