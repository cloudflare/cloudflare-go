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
func (r *ZoneAnalyticsLatencyService) ArgoAnalyticsForZoneArgoAnalyticsForAZone(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/latency", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SchemasResponseSingle struct {
	Errors   []SchemasResponseSingleError   `json:"errors"`
	Messages []SchemasResponseSingleMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success SchemasResponseSingleSuccess `json:"success"`
	JSON    schemasResponseSingleJSON    `json:"-"`
}

// schemasResponseSingleJSON contains the JSON metadata for the struct
// [SchemasResponseSingle]
type schemasResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseSingleError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    schemasResponseSingleErrorJSON `json:"-"`
}

// schemasResponseSingleErrorJSON contains the JSON metadata for the struct
// [SchemasResponseSingleError]
type schemasResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseSingleMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    schemasResponseSingleMessageJSON `json:"-"`
}

// schemasResponseSingleMessageJSON contains the JSON metadata for the struct
// [SchemasResponseSingleMessage]
type schemasResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasResponseSingleSuccess bool

const (
	SchemasResponseSingleSuccessTrue SchemasResponseSingleSuccess = true
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
