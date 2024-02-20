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
func (r *AvailableRatePlanService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]AvailableRatePlanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailableRatePlanListResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailableRatePlanListResponse struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []AvailableRatePlanListResponseComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency AvailableRatePlanListResponseFrequency `json:"frequency"`
	// The plan name.
	Name string                            `json:"name"`
	JSON availableRatePlanListResponseJSON `json:"-"`
}

// availableRatePlanListResponseJSON contains the JSON metadata for the struct
// [AvailableRatePlanListResponse]
type availableRatePlanListResponseJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanListResponseComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name AvailableRatePlanListResponseComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                                    `json:"unit_price"`
	JSON      availableRatePlanListResponseComponentJSON `json:"-"`
}

// availableRatePlanListResponseComponentJSON contains the JSON metadata for the
// struct [AvailableRatePlanListResponseComponent]
type availableRatePlanListResponseComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanListResponseComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The unique component.
type AvailableRatePlanListResponseComponentsName string

const (
	AvailableRatePlanListResponseComponentsNameZones                       AvailableRatePlanListResponseComponentsName = "zones"
	AvailableRatePlanListResponseComponentsNamePageRules                   AvailableRatePlanListResponseComponentsName = "page_rules"
	AvailableRatePlanListResponseComponentsNameDedicatedCertificates       AvailableRatePlanListResponseComponentsName = "dedicated_certificates"
	AvailableRatePlanListResponseComponentsNameDedicatedCertificatesCustom AvailableRatePlanListResponseComponentsName = "dedicated_certificates_custom"
)

// The frequency at which you will be billed for this plan.
type AvailableRatePlanListResponseFrequency string

const (
	AvailableRatePlanListResponseFrequencyWeekly    AvailableRatePlanListResponseFrequency = "weekly"
	AvailableRatePlanListResponseFrequencyMonthly   AvailableRatePlanListResponseFrequency = "monthly"
	AvailableRatePlanListResponseFrequencyQuarterly AvailableRatePlanListResponseFrequency = "quarterly"
	AvailableRatePlanListResponseFrequencyYearly    AvailableRatePlanListResponseFrequency = "yearly"
)

type AvailableRatePlanListResponseEnvelope struct {
	Errors   []AvailableRatePlanListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AvailableRatePlanListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AvailableRatePlanListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AvailableRatePlanListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AvailableRatePlanListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       availableRatePlanListResponseEnvelopeJSON       `json:"-"`
}

// availableRatePlanListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AvailableRatePlanListResponseEnvelope]
type availableRatePlanListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    availableRatePlanListResponseEnvelopeErrorsJSON `json:"-"`
}

// availableRatePlanListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AvailableRatePlanListResponseEnvelopeErrors]
type availableRatePlanListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    availableRatePlanListResponseEnvelopeMessagesJSON `json:"-"`
}

// availableRatePlanListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AvailableRatePlanListResponseEnvelopeMessages]
type availableRatePlanListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AvailableRatePlanListResponseEnvelopeSuccess bool

const (
	AvailableRatePlanListResponseEnvelopeSuccessTrue AvailableRatePlanListResponseEnvelopeSuccess = true
)

type AvailableRatePlanListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       availableRatePlanListResponseEnvelopeResultInfoJSON `json:"-"`
}

// availableRatePlanListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AvailableRatePlanListResponseEnvelopeResultInfo]
type availableRatePlanListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
