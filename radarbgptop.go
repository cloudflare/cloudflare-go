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

// RadarBGPTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTopService] method
// instead.
type RadarBGPTopService struct {
	Options []option.RequestOption
	Ases    *RadarBGPTopAseService
}

// NewRadarBGPTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPTopService(opts ...option.RequestOption) (r *RadarBGPTopService) {
	r = &RadarBGPTopService{}
	r.Options = opts
	r.Ases = NewRadarBGPTopAseService(opts...)
	return
}

// Get the top network prefixes by BGP updates. Values are a percentage out of the
// total BGP updates.
func (r *RadarBGPTopService) Prefixes(ctx context.Context, query RadarBGPTopPrefixesParams, opts ...option.RequestOption) (res *RadarBGPTopPrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPTopPrefixesResponseEnvelope
	path := "radar/bgp/top/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPTopPrefixesResponse struct {
	Meta RadarBGPTopPrefixesResponseMeta   `json:"meta,required"`
	Top0 []RadarBGPTopPrefixesResponseTop0 `json:"top_0,required"`
	JSON radarBGPTopPrefixesResponseJSON   `json:"-"`
}

// radarBGPTopPrefixesResponseJSON contains the JSON metadata for the struct
// [RadarBGPTopPrefixesResponse]
type radarBGPTopPrefixesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixesResponseMeta struct {
	DateRange []RadarBGPTopPrefixesResponseMetaDateRange `json:"dateRange,required"`
	JSON      radarBGPTopPrefixesResponseMetaJSON        `json:"-"`
}

// radarBGPTopPrefixesResponseMetaJSON contains the JSON metadata for the struct
// [RadarBGPTopPrefixesResponseMeta]
type radarBGPTopPrefixesResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      radarBGPTopPrefixesResponseMetaDateRangeJSON `json:"-"`
}

// radarBGPTopPrefixesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarBGPTopPrefixesResponseMetaDateRange]
type radarBGPTopPrefixesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixesResponseTop0 struct {
	Prefix string                              `json:"prefix,required"`
	Value  string                              `json:"value,required"`
	JSON   radarBGPTopPrefixesResponseTop0JSON `json:"-"`
}

// radarBGPTopPrefixesResponseTop0JSON contains the JSON metadata for the struct
// [RadarBGPTopPrefixesResponseTop0]
type radarBGPTopPrefixesResponseTop0JSON struct {
	Prefix      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPTopPrefixesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarBGPTopPrefixesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBGPTopPrefixesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBGPTopPrefixesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBGPTopPrefixesParams]'s query parameters as
// `url.Values`.
func (r RadarBGPTopPrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarBGPTopPrefixesParamsDateRange string

const (
	RadarBGPTopPrefixesParamsDateRange1d         RadarBGPTopPrefixesParamsDateRange = "1d"
	RadarBGPTopPrefixesParamsDateRange2d         RadarBGPTopPrefixesParamsDateRange = "2d"
	RadarBGPTopPrefixesParamsDateRange7d         RadarBGPTopPrefixesParamsDateRange = "7d"
	RadarBGPTopPrefixesParamsDateRange14d        RadarBGPTopPrefixesParamsDateRange = "14d"
	RadarBGPTopPrefixesParamsDateRange28d        RadarBGPTopPrefixesParamsDateRange = "28d"
	RadarBGPTopPrefixesParamsDateRange12w        RadarBGPTopPrefixesParamsDateRange = "12w"
	RadarBGPTopPrefixesParamsDateRange24w        RadarBGPTopPrefixesParamsDateRange = "24w"
	RadarBGPTopPrefixesParamsDateRange52w        RadarBGPTopPrefixesParamsDateRange = "52w"
	RadarBGPTopPrefixesParamsDateRange1dControl  RadarBGPTopPrefixesParamsDateRange = "1dControl"
	RadarBGPTopPrefixesParamsDateRange2dControl  RadarBGPTopPrefixesParamsDateRange = "2dControl"
	RadarBGPTopPrefixesParamsDateRange7dControl  RadarBGPTopPrefixesParamsDateRange = "7dControl"
	RadarBGPTopPrefixesParamsDateRange14dControl RadarBGPTopPrefixesParamsDateRange = "14dControl"
	RadarBGPTopPrefixesParamsDateRange28dControl RadarBGPTopPrefixesParamsDateRange = "28dControl"
	RadarBGPTopPrefixesParamsDateRange12wControl RadarBGPTopPrefixesParamsDateRange = "12wControl"
	RadarBGPTopPrefixesParamsDateRange24wControl RadarBGPTopPrefixesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPTopPrefixesParamsFormat string

const (
	RadarBGPTopPrefixesParamsFormatJson RadarBGPTopPrefixesParamsFormat = "JSON"
	RadarBGPTopPrefixesParamsFormatCsv  RadarBGPTopPrefixesParamsFormat = "CSV"
)

type RadarBGPTopPrefixesParamsUpdateType string

const (
	RadarBGPTopPrefixesParamsUpdateTypeAnnouncement RadarBGPTopPrefixesParamsUpdateType = "ANNOUNCEMENT"
	RadarBGPTopPrefixesParamsUpdateTypeWithdrawal   RadarBGPTopPrefixesParamsUpdateType = "WITHDRAWAL"
)

type RadarBGPTopPrefixesResponseEnvelope struct {
	Result  RadarBGPTopPrefixesResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarBGPTopPrefixesResponseEnvelopeJSON `json:"-"`
}

// radarBGPTopPrefixesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPTopPrefixesResponseEnvelope]
type radarBGPTopPrefixesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPTopPrefixesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
