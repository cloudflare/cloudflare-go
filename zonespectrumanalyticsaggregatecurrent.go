// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSpectrumAnalyticsAggregateCurrentService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSpectrumAnalyticsAggregateCurrentService] method instead.
type ZoneSpectrumAnalyticsAggregateCurrentService struct {
	Options []option.RequestOption
}

// NewZoneSpectrumAnalyticsAggregateCurrentService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSpectrumAnalyticsAggregateCurrentService(opts ...option.RequestOption) (r *ZoneSpectrumAnalyticsAggregateCurrentService) {
	r = &ZoneSpectrumAnalyticsAggregateCurrentService{}
	r.Options = opts
	return
}

// Retrieves analytics aggregated from the last minute of usage on Spectrum
// applications underneath a given zone.
func (r *ZoneSpectrumAnalyticsAggregateCurrentService) SpectrumAggregateAnalyticsGetCurrentAggregatedAnalytics(ctx context.Context, zone string, query ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams, opts ...option.RequestOption) (res *ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/spectrum/analytics/aggregate/current", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse struct {
	Errors   []ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseError   `json:"errors"`
	Messages []ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessage `json:"messages"`
	Result   []interface{}                                                                                                 `json:"result"`
	// Whether the API call was successful
	Success ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseSuccess `json:"success"`
	JSON    zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseJSON    `json:"-"`
}

// zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseJSON
// contains the JSON metadata for the struct
// [ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse]
type zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseError struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseErrorJSON `json:"-"`
}

// zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseError]
type zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessage struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessageJSON `json:"-"`
}

// zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessage]
type zoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseSuccess bool

const (
	ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseSuccessTrue ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseSuccess = true
)

type ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams struct {
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppIDParam param.Field[string] `query:"app_id_param"`
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppID param.Field[string] `query:"appID"`
	// Co-location identifier.
	ColoName param.Field[string] `query:"colo_name"`
}

// URLQuery serializes
// [ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams]'s
// query parameters as `url.Values`.
func (r ZoneSpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
