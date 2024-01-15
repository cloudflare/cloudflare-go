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
func (r *ZoneAvailableRatePlanService) ZoneRatePlanListAvailableRatePlans(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponse struct {
	Errors     []ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseError    `json:"errors"`
	Messages   []ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessage  `json:"messages"`
	Result     []ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResult   `json:"result"`
	ResultInfo ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseSuccess `json:"success"`
	JSON    zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseJSON    `json:"-"`
}

// zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseJSON contains the
// JSON metadata for the struct
// [ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponse]
type zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseErrorJSON `json:"-"`
}

// zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseError]
type zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessageJSON `json:"-"`
}

// zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessage]
type zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResult struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequency `json:"frequency"`
	// The plan name.
	Name string                                                                    `json:"name"`
	JSON zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultJSON `json:"-"`
}

// zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResult]
type zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                                                                            `json:"unit_price"`
	JSON      zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentJSON `json:"-"`
}

// zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentJSON
// contains the JSON metadata for the struct
// [ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponent]
type zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The unique component.
type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsName string

const (
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsNameZones                       ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsName = "zones"
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsNamePageRules                   ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsName = "page_rules"
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsNameDedicatedCertificates       ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsName = "dedicated_certificates"
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsNameDedicatedCertificatesCustom ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultComponentsName = "dedicated_certificates_custom"
)

// The frequency at which you will be billed for this plan.
type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequency string

const (
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequencyWeekly    ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequency = "weekly"
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequencyMonthly   ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequency = "monthly"
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequencyQuarterly ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequency = "quarterly"
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequencyYearly    ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultFrequency = "yearly"
)

type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfoJSON `json:"-"`
}

// zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfo]
type zoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseSuccess bool

const (
	ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseSuccessTrue ZoneAvailableRatePlanZoneRatePlanListAvailableRatePlansResponseSuccess = true
)
