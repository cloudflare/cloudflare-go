// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarBGPTopPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTopPrefixService] method
// instead.
type RadarBGPTopPrefixService struct {
	Options []option.RequestOption
}

// NewRadarBGPTopPrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPTopPrefixService(opts ...option.RequestOption) (r *RadarBGPTopPrefixService) {
	r = &RadarBGPTopPrefixService{}
	r.Options = opts
	return
}

// Get the top network prefixes by BGP updates. Values are a percentage out of the
// total BGP updates.
func (r *RadarBGPTopPrefixService) List(ctx context.Context, query RadarBGPTopPrefixListParams, opts ...option.RequestOption) (res *RadarBGPTopPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopPrefixListResponseEnvelope
	path := "radar/bgp/top/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTopPrefixListResponse struct {
	Meta RadarBGPTopPrefixListResponseMeta   `json:"meta,required"`
	Top0 []RadarBGPTopPrefixListResponseTop0 `json:"top_0,required"`
	JSON radarBGPTopPrefixListResponseJSON   `json:"-"`
}

// radarBGPTopPrefixListResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopPrefixListResponse]
type radarBGPTopPrefixListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixListResponseMeta struct {
	DateRange []RadarBGPTopPrefixListResponseMetaDateRange `json:"dateRange,required"`
	JSON      radarBGPTopPrefixListResponseMetaJSON        `json:"-"`
}

// radarBGPTopPrefixListResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTopPrefixListResponseMeta]
type radarBGPTopPrefixListResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      radarBGPTopPrefixListResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPTopPrefixListResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarBGPTopPrefixListResponseMetaDateRange]
type radarBGPTopPrefixListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixListResponseTop0 struct {
	Prefix string                                `json:"prefix,required"`
	Value  string                                `json:"value,required"`
	JSON   radarBGPTopPrefixListResponseTop0JSON `json:"-"`
}

// radarBGPTopPrefixListResponseTop0JSON contains the JSON metadata for the struct
// [RadarBGPTopPrefixListResponseTop0]
type radarBGPTopPrefixListResponseTop0JSON struct {
	Prefix      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBGPTopPrefixListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopPrefixListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBGPTopPrefixListParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBGPTopPrefixListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopPrefixListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarBGPTopPrefixListParamsDateRange string

const (
	RadarBGPTopPrefixListParamsDateRange1d         RadarBGPTopPrefixListParamsDateRange = "1d"
	RadarBGPTopPrefixListParamsDateRange2d         RadarBGPTopPrefixListParamsDateRange = "2d"
	RadarBGPTopPrefixListParamsDateRange7d         RadarBGPTopPrefixListParamsDateRange = "7d"
	RadarBGPTopPrefixListParamsDateRange14d        RadarBGPTopPrefixListParamsDateRange = "14d"
	RadarBGPTopPrefixListParamsDateRange28d        RadarBGPTopPrefixListParamsDateRange = "28d"
	RadarBGPTopPrefixListParamsDateRange12w        RadarBGPTopPrefixListParamsDateRange = "12w"
	RadarBGPTopPrefixListParamsDateRange24w        RadarBGPTopPrefixListParamsDateRange = "24w"
	RadarBGPTopPrefixListParamsDateRange52w        RadarBGPTopPrefixListParamsDateRange = "52w"
	RadarBGPTopPrefixListParamsDateRange1dControl  RadarBGPTopPrefixListParamsDateRange = "1dControl"
	RadarBGPTopPrefixListParamsDateRange2dControl  RadarBGPTopPrefixListParamsDateRange = "2dControl"
	RadarBGPTopPrefixListParamsDateRange7dControl  RadarBGPTopPrefixListParamsDateRange = "7dControl"
	RadarBGPTopPrefixListParamsDateRange14dControl RadarBGPTopPrefixListParamsDateRange = "14dControl"
	RadarBGPTopPrefixListParamsDateRange28dControl RadarBGPTopPrefixListParamsDateRange = "28dControl"
	RadarBGPTopPrefixListParamsDateRange12wControl RadarBGPTopPrefixListParamsDateRange = "12wControl"
	RadarBGPTopPrefixListParamsDateRange24wControl RadarBGPTopPrefixListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPTopPrefixListParamsFormat string

const (
	RadarBGPTopPrefixListParamsFormatJson RadarBGPTopPrefixListParamsFormat = "JSON"
	RadarBGPTopPrefixListParamsFormatCsv  RadarBGPTopPrefixListParamsFormat = "CSV"
)

type RadarBGPTopPrefixListParamsUpdateType string

const (
	RadarBGPTopPrefixListParamsUpdateTypeAnnouncement RadarBGPTopPrefixListParamsUpdateType = "ANNOUNCEMENT"
	RadarBGPTopPrefixListParamsUpdateTypeWithdrawal   RadarBGPTopPrefixListParamsUpdateType = "WITHDRAWAL"
)

type RadarBGPTopPrefixListResponseEnvelope struct {
	Result  RadarBGPTopPrefixListResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarBGPTopPrefixListResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopPrefixListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPTopPrefixListResponseEnvelope]
type radarBGPTopPrefixListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
