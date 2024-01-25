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

// ZoneAvailablePlanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAvailablePlanService] method
// instead.
type ZoneAvailablePlanService struct {
	Options []option.RequestOption
}

// NewZoneAvailablePlanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAvailablePlanService(opts ...option.RequestOption) (r *ZoneAvailablePlanService) {
	r = &ZoneAvailablePlanService{}
	r.Options = opts
	return
}

// Details of the available plan that the zone can subscribe to.
func (r *ZoneAvailablePlanService) Get(ctx context.Context, zoneIdentifier string, planIdentifier string, opts ...option.RequestOption) (res *ZoneAvailablePlanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/available_plans/%s", zoneIdentifier, planIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists available plans the zone can subscribe to.
func (r *ZoneAvailablePlanService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneAvailablePlanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/available_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAvailablePlanGetResponse struct {
	Errors   []ZoneAvailablePlanGetResponseError   `json:"errors"`
	Messages []ZoneAvailablePlanGetResponseMessage `json:"messages"`
	Result   ZoneAvailablePlanGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAvailablePlanGetResponseSuccess `json:"success"`
	JSON    zoneAvailablePlanGetResponseJSON    `json:"-"`
}

// zoneAvailablePlanGetResponseJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanGetResponse]
type zoneAvailablePlanGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailablePlanGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneAvailablePlanGetResponseErrorJSON `json:"-"`
}

// zoneAvailablePlanGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanGetResponseError]
type zoneAvailablePlanGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailablePlanGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneAvailablePlanGetResponseMessageJSON `json:"-"`
}

// zoneAvailablePlanGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAvailablePlanGetResponseMessage]
type zoneAvailablePlanGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailablePlanGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency ZoneAvailablePlanGetResponseResultFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64                                `json:"price"`
	JSON  zoneAvailablePlanGetResponseResultJSON `json:"-"`
}

// zoneAvailablePlanGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanGetResponseResult]
type zoneAvailablePlanGetResponseResultJSON struct {
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

func (r *ZoneAvailablePlanGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which you will be billed for this plan.
type ZoneAvailablePlanGetResponseResultFrequency string

const (
	ZoneAvailablePlanGetResponseResultFrequencyWeekly    ZoneAvailablePlanGetResponseResultFrequency = "weekly"
	ZoneAvailablePlanGetResponseResultFrequencyMonthly   ZoneAvailablePlanGetResponseResultFrequency = "monthly"
	ZoneAvailablePlanGetResponseResultFrequencyQuarterly ZoneAvailablePlanGetResponseResultFrequency = "quarterly"
	ZoneAvailablePlanGetResponseResultFrequencyYearly    ZoneAvailablePlanGetResponseResultFrequency = "yearly"
)

// Whether the API call was successful
type ZoneAvailablePlanGetResponseSuccess bool

const (
	ZoneAvailablePlanGetResponseSuccessTrue ZoneAvailablePlanGetResponseSuccess = true
)

type ZoneAvailablePlanListResponse struct {
	Errors     []ZoneAvailablePlanListResponseError    `json:"errors"`
	Messages   []ZoneAvailablePlanListResponseMessage  `json:"messages"`
	Result     []ZoneAvailablePlanListResponseResult   `json:"result"`
	ResultInfo ZoneAvailablePlanListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAvailablePlanListResponseSuccess `json:"success"`
	JSON    zoneAvailablePlanListResponseJSON    `json:"-"`
}

// zoneAvailablePlanListResponseJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanListResponse]
type zoneAvailablePlanListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailablePlanListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneAvailablePlanListResponseErrorJSON `json:"-"`
}

// zoneAvailablePlanListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanListResponseError]
type zoneAvailablePlanListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailablePlanListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneAvailablePlanListResponseMessageJSON `json:"-"`
}

// zoneAvailablePlanListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAvailablePlanListResponseMessage]
type zoneAvailablePlanListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAvailablePlanListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency ZoneAvailablePlanListResponseResultFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64                                 `json:"price"`
	JSON  zoneAvailablePlanListResponseResultJSON `json:"-"`
}

// zoneAvailablePlanListResponseResultJSON contains the JSON metadata for the
// struct [ZoneAvailablePlanListResponseResult]
type zoneAvailablePlanListResponseResultJSON struct {
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

func (r *ZoneAvailablePlanListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which you will be billed for this plan.
type ZoneAvailablePlanListResponseResultFrequency string

const (
	ZoneAvailablePlanListResponseResultFrequencyWeekly    ZoneAvailablePlanListResponseResultFrequency = "weekly"
	ZoneAvailablePlanListResponseResultFrequencyMonthly   ZoneAvailablePlanListResponseResultFrequency = "monthly"
	ZoneAvailablePlanListResponseResultFrequencyQuarterly ZoneAvailablePlanListResponseResultFrequency = "quarterly"
	ZoneAvailablePlanListResponseResultFrequencyYearly    ZoneAvailablePlanListResponseResultFrequency = "yearly"
)

type ZoneAvailablePlanListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       zoneAvailablePlanListResponseResultInfoJSON `json:"-"`
}

// zoneAvailablePlanListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneAvailablePlanListResponseResultInfo]
type zoneAvailablePlanListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAvailablePlanListResponseSuccess bool

const (
	ZoneAvailablePlanListResponseSuccessTrue ZoneAvailablePlanListResponseSuccess = true
)
