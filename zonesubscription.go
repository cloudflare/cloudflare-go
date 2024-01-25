// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *ZoneSubscriptionService) ZoneSubscriptionNewZoneSubscription(ctx context.Context, identifier string, body ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParams, opts ...option.RequestOption) (res *ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *ZoneSubscriptionService) ZoneSubscriptionUpdateZoneSubscription(ctx context.Context, identifier string, body ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParams, opts ...option.RequestOption) (res *ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists zone subscription details.
func (r *ZoneSubscriptionService) ZoneSubscriptionZoneSubscriptionDetails(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponse struct {
	Errors   []ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseError   `json:"errors"`
	Messages []ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessage `json:"messages"`
	Result   interface{}                                                          `json:"result"`
	// Whether the API call was successful
	Success ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseSuccess `json:"success"`
	JSON    zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseJSON    `json:"-"`
}

// zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseJSON contains the
// JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponse]
type zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseError struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseErrorJSON `json:"-"`
}

// zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseError]
type zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessage struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessageJSON `json:"-"`
}

// zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessage]
type zoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseSuccess bool

const (
	ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseSuccessTrue ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionResponseSuccess = true
)

type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse struct {
	Errors   []ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseError   `json:"errors"`
	Messages []ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessage `json:"messages"`
	Result   interface{}                                                             `json:"result"`
	// Whether the API call was successful
	Success ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseSuccess `json:"success"`
	JSON    zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseJSON    `json:"-"`
}

// zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseJSON contains the
// JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse]
type zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseErrorJSON `json:"-"`
}

// zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseError]
type zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessageJSON `json:"-"`
}

// zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessage]
type zoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseSuccess bool

const (
	ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseSuccessTrue ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionResponseSuccess = true
)

type ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse struct {
	Errors   []ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseError   `json:"errors"`
	Messages []ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessage `json:"messages"`
	Result   interface{}                                                              `json:"result"`
	// Whether the API call was successful
	Success ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseSuccess `json:"success"`
	JSON    zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseJSON    `json:"-"`
}

// zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseJSON contains the
// JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse]
type zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseErrorJSON `json:"-"`
}

// zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseError]
type zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessageJSON `json:"-"`
}

// zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessage]
type zoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseSuccess bool

const (
	ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseSuccessTrue ZoneSubscriptionZoneSubscriptionZoneSubscriptionDetailsResponseSuccess = true
)

type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParams struct {
	App param.Field[ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsZone] `json:"zone"`
}

func (r ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency string

const (
	ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyWeekly    ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "weekly"
	ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyMonthly   ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "monthly"
	ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyQuarterly ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "quarterly"
	ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequencyYearly    ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsRatePlan struct {
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

func (r ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsZone struct {
}

func (r ZoneSubscriptionZoneSubscriptionNewZoneSubscriptionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParams struct {
	App param.Field[ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsZone] `json:"zone"`
}

func (r ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency string

const (
	ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyWeekly    ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "weekly"
	ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyMonthly   ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "monthly"
	ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyQuarterly ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "quarterly"
	ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequencyYearly    ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsRatePlan struct {
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

func (r ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsZone struct {
}

func (r ZoneSubscriptionZoneSubscriptionUpdateZoneSubscriptionParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
