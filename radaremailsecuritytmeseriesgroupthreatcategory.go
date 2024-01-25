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

// RadarEmailSecurityTmeseriesGroupThreatCategoryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityTmeseriesGroupThreatCategoryService] method instead.
type RadarEmailSecurityTmeseriesGroupThreatCategoryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTmeseriesGroupThreatCategoryService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTmeseriesGroupThreatCategoryService(opts ...option.RequestOption) (r *RadarEmailSecurityTmeseriesGroupThreatCategoryService) {
	r = &RadarEmailSecurityTmeseriesGroupThreatCategoryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified in Threat Categories over time.
func (r *RadarEmailSecurityTmeseriesGroupThreatCategoryService) List(ctx context.Context, query RadarEmailSecurityTmeseriesGroupThreatCategoryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTmeseriesGroupThreatCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTmeseriesGroupThreatCategoryListResponse struct {
	Result  RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResult `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    radarEmailSecurityTmeseriesGroupThreatCategoryListResponseJSON   `json:"-"`
}

// radarEmailSecurityTmeseriesGroupThreatCategoryListResponseJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTmeseriesGroupThreatCategoryListResponse]
type radarEmailSecurityTmeseriesGroupThreatCategoryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTmeseriesGroupThreatCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResult struct {
	Meta   interface{}                                                            `json:"meta,required"`
	Serie0 RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResult]
type radarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0 struct {
	BrandImpersonation  []string                                                                   `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                                                   `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                                                   `json:"IdentityDeception,required"`
	Link                []string                                                                   `json:"Link,required"`
	JSON                radarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0]
type radarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTmeseriesGroupThreatCategoryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTmeseriesGroupThreatCategoryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTmeseriesGroupThreatCategoryListParams]'s
// query parameters as `url.Values`.
func (r RadarEmailSecurityTmeseriesGroupThreatCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval15m RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval = "15m"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval1h  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval = "1h"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval1d  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval = "1d"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval1w  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsAggInterval = "1w"
)

type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArc string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArcPass RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArc = "PASS"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArcNone RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArc = "NONE"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArcFail RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsArc = "FAIL"
)

type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange1d         RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "1d"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange2d         RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "2d"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange7d         RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "7d"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange14d        RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "14d"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange28d        RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "28d"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange12w        RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "12w"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange24w        RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "24w"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange52w        RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "52w"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange1dControl  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "1dControl"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange2dControl  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "2dControl"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange7dControl  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "7dControl"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange14dControl RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "14dControl"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange28dControl RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "28dControl"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange12wControl RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "12wControl"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange24wControl RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkim string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkimPass RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkim = "PASS"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkimNone RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkim = "NONE"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkimFail RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDkim = "FAIL"
)

type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarc string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarcPass RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarc = "PASS"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarcNone RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarc = "NONE"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarcFail RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsFormat string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsFormatJson RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsFormat = "JSON"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsFormatCsv  RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsFormat = "CSV"
)

type RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpf string

const (
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpfPass RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpf = "PASS"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpfNone RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpf = "NONE"
	RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpfFail RadarEmailSecurityTmeseriesGroupThreatCategoryListParamsSpf = "FAIL"
)
