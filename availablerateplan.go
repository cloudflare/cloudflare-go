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
func (r *AvailableRatePlanService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]AvailableRatePlanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailableRatePlanGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailableRatePlanGetResponse struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []AvailableRatePlanGetResponseComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency AvailableRatePlanGetResponseFrequency `json:"frequency"`
	// The plan name.
	Name string                           `json:"name"`
	JSON availableRatePlanGetResponseJSON `json:"-"`
}

// availableRatePlanGetResponseJSON contains the JSON metadata for the struct
// [AvailableRatePlanGetResponse]
type availableRatePlanGetResponseJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanGetResponseComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name AvailableRatePlanGetResponseComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                                   `json:"unit_price"`
	JSON      availableRatePlanGetResponseComponentJSON `json:"-"`
}

// availableRatePlanGetResponseComponentJSON contains the JSON metadata for the
// struct [AvailableRatePlanGetResponseComponent]
type availableRatePlanGetResponseComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanGetResponseComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The unique component.
type AvailableRatePlanGetResponseComponentsName string

const (
	AvailableRatePlanGetResponseComponentsNameZones                       AvailableRatePlanGetResponseComponentsName = "zones"
	AvailableRatePlanGetResponseComponentsNamePageRules                   AvailableRatePlanGetResponseComponentsName = "page_rules"
	AvailableRatePlanGetResponseComponentsNameDedicatedCertificates       AvailableRatePlanGetResponseComponentsName = "dedicated_certificates"
	AvailableRatePlanGetResponseComponentsNameDedicatedCertificatesCustom AvailableRatePlanGetResponseComponentsName = "dedicated_certificates_custom"
)

// The frequency at which you will be billed for this plan.
type AvailableRatePlanGetResponseFrequency string

const (
	AvailableRatePlanGetResponseFrequencyWeekly    AvailableRatePlanGetResponseFrequency = "weekly"
	AvailableRatePlanGetResponseFrequencyMonthly   AvailableRatePlanGetResponseFrequency = "monthly"
	AvailableRatePlanGetResponseFrequencyQuarterly AvailableRatePlanGetResponseFrequency = "quarterly"
	AvailableRatePlanGetResponseFrequencyYearly    AvailableRatePlanGetResponseFrequency = "yearly"
)

type AvailableRatePlanGetResponseEnvelope struct {
	Errors   []AvailableRatePlanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AvailableRatePlanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []AvailableRatePlanGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AvailableRatePlanGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AvailableRatePlanGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       availableRatePlanGetResponseEnvelopeJSON       `json:"-"`
}

// availableRatePlanGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AvailableRatePlanGetResponseEnvelope]
type availableRatePlanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    availableRatePlanGetResponseEnvelopeErrorsJSON `json:"-"`
}

// availableRatePlanGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AvailableRatePlanGetResponseEnvelopeErrors]
type availableRatePlanGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableRatePlanGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    availableRatePlanGetResponseEnvelopeMessagesJSON `json:"-"`
}

// availableRatePlanGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AvailableRatePlanGetResponseEnvelopeMessages]
type availableRatePlanGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AvailableRatePlanGetResponseEnvelopeSuccess bool

const (
	AvailableRatePlanGetResponseEnvelopeSuccessTrue AvailableRatePlanGetResponseEnvelopeSuccess = true
)

type AvailableRatePlanGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       availableRatePlanGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// availableRatePlanGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AvailableRatePlanGetResponseEnvelopeResultInfo]
type availableRatePlanGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableRatePlanGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
