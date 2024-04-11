// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_plans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RatePlanService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRatePlanService] method instead.
type RatePlanService struct {
	Options []option.RequestOption
}

// NewRatePlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatePlanService(opts ...option.RequestOption) (r *RatePlanService) {
	r = &RatePlanService{}
	r.Options = opts
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *RatePlanService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]RatePlan, err error) {
	opts = append(r.Options[:], opts...)
	var env RatePlanGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RatePlan struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []RatePlanComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency RatePlanFrequency `json:"frequency"`
	// The plan name.
	Name string       `json:"name"`
	JSON ratePlanJSON `json:"-"`
}

// ratePlanJSON contains the JSON metadata for the struct [RatePlan]
type ratePlanJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanJSON) RawJSON() string {
	return r.raw
}

type RatePlanComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name RatePlanComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64               `json:"unit_price"`
	JSON      ratePlanComponentJSON `json:"-"`
}

// ratePlanComponentJSON contains the JSON metadata for the struct
// [RatePlanComponent]
type ratePlanComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanComponentJSON) RawJSON() string {
	return r.raw
}

// The unique component.
type RatePlanComponentsName string

const (
	RatePlanComponentsNameZones                       RatePlanComponentsName = "zones"
	RatePlanComponentsNamePageRules                   RatePlanComponentsName = "page_rules"
	RatePlanComponentsNameDedicatedCertificates       RatePlanComponentsName = "dedicated_certificates"
	RatePlanComponentsNameDedicatedCertificatesCustom RatePlanComponentsName = "dedicated_certificates_custom"
)

func (r RatePlanComponentsName) IsKnown() bool {
	switch r {
	case RatePlanComponentsNameZones, RatePlanComponentsNamePageRules, RatePlanComponentsNameDedicatedCertificates, RatePlanComponentsNameDedicatedCertificatesCustom:
		return true
	}
	return false
}

// The frequency at which you will be billed for this plan.
type RatePlanFrequency string

const (
	RatePlanFrequencyWeekly    RatePlanFrequency = "weekly"
	RatePlanFrequencyMonthly   RatePlanFrequency = "monthly"
	RatePlanFrequencyQuarterly RatePlanFrequency = "quarterly"
	RatePlanFrequencyYearly    RatePlanFrequency = "yearly"
)

func (r RatePlanFrequency) IsKnown() bool {
	switch r {
	case RatePlanFrequencyWeekly, RatePlanFrequencyMonthly, RatePlanFrequencyQuarterly, RatePlanFrequencyYearly:
		return true
	}
	return false
}

type RatePlanParam struct {
	// Array of available components values for the plan.
	Components param.Field[[]RatePlanComponentParam] `json:"components"`
	// The duration of the plan subscription.
	Duration param.Field[float64] `json:"duration"`
}

func (r RatePlanParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RatePlanComponentParam struct {
	// The default amount allocated.
	Default param.Field[float64] `json:"default"`
	// The unique component.
	Name param.Field[RatePlanComponentsName] `json:"name"`
}

func (r RatePlanComponentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RatePlanGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []RatePlan            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RatePlanGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RatePlanGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ratePlanGetResponseEnvelopeJSON       `json:"-"`
}

// ratePlanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RatePlanGetResponseEnvelope]
type ratePlanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RatePlanGetResponseEnvelopeSuccess bool

const (
	RatePlanGetResponseEnvelopeSuccessTrue RatePlanGetResponseEnvelopeSuccess = true
)

func (r RatePlanGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RatePlanGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RatePlanGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       ratePlanGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ratePlanGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RatePlanGetResponseEnvelopeResultInfo]
type ratePlanGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
