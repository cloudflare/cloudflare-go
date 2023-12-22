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

// ZoneAvailableRatePlanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAvailableRatePlanService]
// method instead.
type ZoneAvailableRatePlanService struct {
	Options []option.RequestOption
}

// NewZoneAvailableRatePlanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAvailableRatePlanService(opts ...option.RequestOption) (r *ZoneAvailableRatePlanService) {
	r = &ZoneAvailableRatePlanService{}
	r.Options = opts
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *ZoneAvailableRatePlanService) ZoneRatePlanListAvailableRatePlans(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *PlanResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PlanResponseCollection struct {
	Errors     []PlanResponseCollectionError    `json:"errors"`
	Messages   []PlanResponseCollectionMessage  `json:"messages"`
	Result     []PlanResponseCollectionResult   `json:"result"`
	ResultInfo PlanResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success PlanResponseCollectionSuccess `json:"success"`
	JSON    planResponseCollectionJSON    `json:"-"`
}

// planResponseCollectionJSON contains the JSON metadata for the struct
// [PlanResponseCollection]
type planResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    planResponseCollectionErrorJSON `json:"-"`
}

// planResponseCollectionErrorJSON contains the JSON metadata for the struct
// [PlanResponseCollectionError]
type planResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    planResponseCollectionMessageJSON `json:"-"`
}

// planResponseCollectionMessageJSON contains the JSON metadata for the struct
// [PlanResponseCollectionMessage]
type planResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanResponseCollectionResult struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []PlanResponseCollectionResultComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency PlanResponseCollectionResultFrequency `json:"frequency"`
	// The plan name.
	Name string                           `json:"name"`
	JSON planResponseCollectionResultJSON `json:"-"`
}

// planResponseCollectionResultJSON contains the JSON metadata for the struct
// [PlanResponseCollectionResult]
type planResponseCollectionResultJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanResponseCollectionResultComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name PlanResponseCollectionResultComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                                   `json:"unit_price"`
	JSON      planResponseCollectionResultComponentJSON `json:"-"`
}

// planResponseCollectionResultComponentJSON contains the JSON metadata for the
// struct [PlanResponseCollectionResultComponent]
type planResponseCollectionResultComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanResponseCollectionResultComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The unique component.
type PlanResponseCollectionResultComponentsName string

const (
	PlanResponseCollectionResultComponentsNameZones                       PlanResponseCollectionResultComponentsName = "zones"
	PlanResponseCollectionResultComponentsNamePageRules                   PlanResponseCollectionResultComponentsName = "page_rules"
	PlanResponseCollectionResultComponentsNameDedicatedCertificates       PlanResponseCollectionResultComponentsName = "dedicated_certificates"
	PlanResponseCollectionResultComponentsNameDedicatedCertificatesCustom PlanResponseCollectionResultComponentsName = "dedicated_certificates_custom"
)

// The frequency at which you will be billed for this plan.
type PlanResponseCollectionResultFrequency string

const (
	PlanResponseCollectionResultFrequencyWeekly    PlanResponseCollectionResultFrequency = "weekly"
	PlanResponseCollectionResultFrequencyMonthly   PlanResponseCollectionResultFrequency = "monthly"
	PlanResponseCollectionResultFrequencyQuarterly PlanResponseCollectionResultFrequency = "quarterly"
	PlanResponseCollectionResultFrequencyYearly    PlanResponseCollectionResultFrequency = "yearly"
)

type PlanResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       planResponseCollectionResultInfoJSON `json:"-"`
}

// planResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [PlanResponseCollectionResultInfo]
type planResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PlanResponseCollectionSuccess bool

const (
	PlanResponseCollectionSuccessTrue PlanResponseCollectionSuccess = true
)
