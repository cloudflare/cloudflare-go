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

// ZoneAnalyticsLatencyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAnalyticsLatencyService]
// method instead.
type ZoneAnalyticsLatencyService struct {
	Options []option.RequestOption
	Colos   *ZoneAnalyticsLatencyColoService
}

// NewZoneAnalyticsLatencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAnalyticsLatencyService(opts ...option.RequestOption) (r *ZoneAnalyticsLatencyService) {
	r = &ZoneAnalyticsLatencyService{}
	r.Options = opts
	r.Colos = NewZoneAnalyticsLatencyColoService(opts...)
	return
}

// Argo Analytics for a zone
func (r *ZoneAnalyticsLatencyService) ArgoAnalyticsForZoneArgoAnalyticsForAZone(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams, opts ...option.RequestOption) (res *ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/latency", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse struct {
	Errors   []ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseError   `json:"errors"`
	Messages []ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessage `json:"messages"`
	Result   interface{}                                                                    `json:"result"`
	// Whether the API call was successful
	Success ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseSuccess `json:"success"`
	JSON    zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseJSON    `json:"-"`
}

// zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse]
type zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseErrorJSON `json:"-"`
}

// zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseError]
type zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessageJSON `json:"-"`
}

// zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessage]
type zoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseSuccess bool

const (
	ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseSuccessTrue ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseSuccess = true
)

type ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams struct {
	Bins param.Field[string] `query:"bins"`
}

// URLQuery serializes
// [ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams]'s query
// parameters as `url.Values`.
func (r ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
