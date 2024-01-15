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

// ZoneAnalyticsLatencyColoService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneAnalyticsLatencyColoService] method instead.
type ZoneAnalyticsLatencyColoService struct {
	Options []option.RequestOption
}

// NewZoneAnalyticsLatencyColoService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAnalyticsLatencyColoService(opts ...option.RequestOption) (r *ZoneAnalyticsLatencyColoService) {
	r = &ZoneAnalyticsLatencyColoService{}
	r.Options = opts
	return
}

// Argo Analytics for a zone at different PoPs
func (r *ZoneAnalyticsLatencyColoService) ArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPs(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/latency/colos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse struct {
	Errors   []ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseError   `json:"errors"`
	Messages []ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessage `json:"messages"`
	Result   interface{}                                                                                              `json:"result"`
	// Whether the API call was successful
	Success ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseSuccess `json:"success"`
	JSON    zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseJSON    `json:"-"`
}

// zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse]
type zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseError struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseErrorJSON `json:"-"`
}

// zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseError]
type zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessage struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessageJSON `json:"-"`
}

// zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessage]
type zoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseSuccess bool

const (
	ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseSuccessTrue ZoneAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseSuccess = true
)
