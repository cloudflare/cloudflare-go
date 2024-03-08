// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// AvailablePlanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAvailablePlanService] method
// instead.
type AvailablePlanService struct {
	Options []option.RequestOption
}

// NewAvailablePlanService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailablePlanService(opts ...option.RequestOption) (r *AvailablePlanService) {
	r = &AvailablePlanService{}
	r.Options = opts
	return
}

// Lists available plans the zone can subscribe to.
func (r *AvailablePlanService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]AvailablePlanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailablePlanListResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details of the available plan that the zone can subscribe to.
func (r *AvailablePlanService) Get(ctx context.Context, zoneIdentifier string, planIdentifier string, opts ...option.RequestOption) (res *AvailablePlanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailablePlanGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_plans/%s", zoneIdentifier, planIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailablePlanListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency AvailablePlanListResponseFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64                       `json:"price"`
	JSON  availablePlanListResponseJSON `json:"-"`
}

// availablePlanListResponseJSON contains the JSON metadata for the struct
// [AvailablePlanListResponse]
type availablePlanListResponseJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AvailablePlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanListResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which you will be billed for this plan.
type AvailablePlanListResponseFrequency string

const (
	AvailablePlanListResponseFrequencyWeekly    AvailablePlanListResponseFrequency = "weekly"
	AvailablePlanListResponseFrequencyMonthly   AvailablePlanListResponseFrequency = "monthly"
	AvailablePlanListResponseFrequencyQuarterly AvailablePlanListResponseFrequency = "quarterly"
	AvailablePlanListResponseFrequencyYearly    AvailablePlanListResponseFrequency = "yearly"
)

type AvailablePlanGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency AvailablePlanGetResponseFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64                      `json:"price"`
	JSON  availablePlanGetResponseJSON `json:"-"`
}

// availablePlanGetResponseJSON contains the JSON metadata for the struct
// [AvailablePlanGetResponse]
type availablePlanGetResponseJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AvailablePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanGetResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which you will be billed for this plan.
type AvailablePlanGetResponseFrequency string

const (
	AvailablePlanGetResponseFrequencyWeekly    AvailablePlanGetResponseFrequency = "weekly"
	AvailablePlanGetResponseFrequencyMonthly   AvailablePlanGetResponseFrequency = "monthly"
	AvailablePlanGetResponseFrequencyQuarterly AvailablePlanGetResponseFrequency = "quarterly"
	AvailablePlanGetResponseFrequencyYearly    AvailablePlanGetResponseFrequency = "yearly"
)

type AvailablePlanListResponseEnvelope struct {
	Errors   []AvailablePlanListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AvailablePlanListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AvailablePlanListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AvailablePlanListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AvailablePlanListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       availablePlanListResponseEnvelopeJSON       `json:"-"`
}

// availablePlanListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailablePlanListResponseEnvelope]
type availablePlanListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AvailablePlanListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    availablePlanListResponseEnvelopeErrorsJSON `json:"-"`
}

// availablePlanListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AvailablePlanListResponseEnvelopeErrors]
type availablePlanListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AvailablePlanListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    availablePlanListResponseEnvelopeMessagesJSON `json:"-"`
}

// availablePlanListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AvailablePlanListResponseEnvelopeMessages]
type availablePlanListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AvailablePlanListResponseEnvelopeSuccess bool

const (
	AvailablePlanListResponseEnvelopeSuccessTrue AvailablePlanListResponseEnvelopeSuccess = true
)

type AvailablePlanListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       availablePlanListResponseEnvelopeResultInfoJSON `json:"-"`
}

// availablePlanListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AvailablePlanListResponseEnvelopeResultInfo]
type availablePlanListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AvailablePlanGetResponseEnvelope struct {
	Errors   []AvailablePlanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AvailablePlanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AvailablePlanGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AvailablePlanGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    availablePlanGetResponseEnvelopeJSON    `json:"-"`
}

// availablePlanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailablePlanGetResponseEnvelope]
type availablePlanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AvailablePlanGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    availablePlanGetResponseEnvelopeErrorsJSON `json:"-"`
}

// availablePlanGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AvailablePlanGetResponseEnvelopeErrors]
type availablePlanGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AvailablePlanGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    availablePlanGetResponseEnvelopeMessagesJSON `json:"-"`
}

// availablePlanGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AvailablePlanGetResponseEnvelopeMessages]
type availablePlanGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailablePlanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availablePlanGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AvailablePlanGetResponseEnvelopeSuccess bool

const (
	AvailablePlanGetResponseEnvelopeSuccessTrue AvailablePlanGetResponseEnvelopeSuccess = true
)
