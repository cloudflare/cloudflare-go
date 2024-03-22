// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PlanService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPlanService] method instead.
type PlanService struct {
	Options []option.RequestOption
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	return
}

// Lists available plans the zone can subscribe to.
func (r *PlanService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]BillSubsAPIAvailableRatePlan, err error) {
	opts = append(r.Options[:], opts...)
	var env PlanListResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details of the available plan that the zone can subscribe to.
func (r *PlanService) Get(ctx context.Context, zoneIdentifier string, planIdentifier string, opts ...option.RequestOption) (res *BillSubsAPIAvailableRatePlan, err error) {
	opts = append(r.Options[:], opts...)
	var env PlanGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_plans/%s", zoneIdentifier, planIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BillSubsAPIAvailableRatePlan struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency BillSubsAPIAvailableRatePlanFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64                          `json:"price"`
	JSON  billSubsAPIAvailableRatePlanJSON `json:"-"`
}

// billSubsAPIAvailableRatePlanJSON contains the JSON metadata for the struct
// [BillSubsAPIAvailableRatePlan]
type billSubsAPIAvailableRatePlanJSON struct {
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

func (r *BillSubsAPIAvailableRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIAvailableRatePlanJSON) RawJSON() string {
	return r.raw
}

// The frequency at which you will be billed for this plan.
type BillSubsAPIAvailableRatePlanFrequency string

const (
	BillSubsAPIAvailableRatePlanFrequencyWeekly    BillSubsAPIAvailableRatePlanFrequency = "weekly"
	BillSubsAPIAvailableRatePlanFrequencyMonthly   BillSubsAPIAvailableRatePlanFrequency = "monthly"
	BillSubsAPIAvailableRatePlanFrequencyQuarterly BillSubsAPIAvailableRatePlanFrequency = "quarterly"
	BillSubsAPIAvailableRatePlanFrequencyYearly    BillSubsAPIAvailableRatePlanFrequency = "yearly"
)

func (r BillSubsAPIAvailableRatePlanFrequency) IsKnown() bool {
	switch r {
	case BillSubsAPIAvailableRatePlanFrequencyWeekly, BillSubsAPIAvailableRatePlanFrequencyMonthly, BillSubsAPIAvailableRatePlanFrequencyQuarterly, BillSubsAPIAvailableRatePlanFrequencyYearly:
		return true
	}
	return false
}

type PlanListResponseEnvelope struct {
	Errors   []PlanListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PlanListResponseEnvelopeMessages `json:"messages,required"`
	Result   []BillSubsAPIAvailableRatePlan     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PlanListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PlanListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       planListResponseEnvelopeJSON       `json:"-"`
}

// planListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PlanListResponseEnvelope]
type planListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PlanListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    planListResponseEnvelopeErrorsJSON `json:"-"`
}

// planListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PlanListResponseEnvelopeErrors]
type planListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PlanListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    planListResponseEnvelopeMessagesJSON `json:"-"`
}

// planListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PlanListResponseEnvelopeMessages]
type planListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PlanListResponseEnvelopeSuccess bool

const (
	PlanListResponseEnvelopeSuccessTrue PlanListResponseEnvelopeSuccess = true
)

func (r PlanListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PlanListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PlanListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       planListResponseEnvelopeResultInfoJSON `json:"-"`
}

// planListResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [PlanListResponseEnvelopeResultInfo]
type planListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PlanGetResponseEnvelope struct {
	Errors   []PlanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PlanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   BillSubsAPIAvailableRatePlan      `json:"result,required"`
	// Whether the API call was successful
	Success PlanGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    planGetResponseEnvelopeJSON    `json:"-"`
}

// planGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PlanGetResponseEnvelope]
type planGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PlanGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    planGetResponseEnvelopeErrorsJSON `json:"-"`
}

// planGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PlanGetResponseEnvelopeErrors]
type planGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PlanGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    planGetResponseEnvelopeMessagesJSON `json:"-"`
}

// planGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PlanGetResponseEnvelopeMessages]
type planGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PlanGetResponseEnvelopeSuccess bool

const (
	PlanGetResponseEnvelopeSuccessTrue PlanGetResponseEnvelopeSuccess = true
)

func (r PlanGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PlanGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
