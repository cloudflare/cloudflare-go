// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AvailableRatePlanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAvailableRatePlanService] method
// instead.
type AvailableRatePlanService struct {
	Options []option.RequestOption
}

// NewAvailableRatePlanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAvailableRatePlanService(opts ...option.RequestOption) (r *AvailableRatePlanService) {
	r = &AvailableRatePlanService{}
	r.Options = opts
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *AvailableRatePlanService) ZoneRatePlanListAvailableRatePlans(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]AvailableRatePlanZoneRatePlanListAvailableRatePlansResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponse struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequency `json:"frequency"`
	// The plan name.
	Name string                                                          `json:"name"`
	JSON availableRatePlanZoneRatePlanListAvailableRatePlansResponseJSON `json:"-"`
}

// availableRatePlanZoneRatePlanListAvailableRatePlansResponseJSON contains the
// JSON metadata for the struct
// [AvailableRatePlanZoneRatePlanListAvailableRatePlansResponse]
type availableRatePlanZoneRatePlanListAvailableRatePlansResponseJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanZoneRatePlanListAvailableRatePlansResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                                                                  `json:"unit_price"`
	JSON      availableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentJSON `json:"-"`
}

// availableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentJSON
// contains the JSON metadata for the struct
// [AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponent]
type availableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The unique component.
type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsName string

const (
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsNameZones                       AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsName = "zones"
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsNamePageRules                   AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsName = "page_rules"
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsNameDedicatedCertificates       AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsName = "dedicated_certificates"
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsNameDedicatedCertificatesCustom AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseComponentsName = "dedicated_certificates_custom"
)

// The frequency at which you will be billed for this plan.
type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequency string

const (
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequencyWeekly    AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequency = "weekly"
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequencyMonthly   AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequency = "monthly"
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequencyQuarterly AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequency = "quarterly"
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequencyYearly    AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseFrequency = "yearly"
)

type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelope struct {
	Errors   []AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessages `json:"messages,required"`
	Result   []AvailableRatePlanZoneRatePlanListAvailableRatePlansResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfo `json:"result_info"`
	JSON       availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeJSON       `json:"-"`
}

// availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelope]
type availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrorsJSON `json:"-"`
}

// availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrors]
type availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessagesJSON `json:"-"`
}

// availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessages]
type availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeSuccess bool

const (
	AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeSuccessTrue AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeSuccess = true
)

type AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfoJSON `json:"-"`
}

// availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfo]
type availableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanZoneRatePlanListAvailableRatePlansResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
