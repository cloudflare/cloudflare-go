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

// RadarEmailTimeseryThreatCategoryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailTimeseryThreatCategoryService] method instead.
type RadarEmailTimeseryThreatCategoryService struct {
	Options []option.RequestOption
}

// NewRadarEmailTimeseryThreatCategoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTimeseryThreatCategoryService(opts ...option.RequestOption) (r *RadarEmailTimeseryThreatCategoryService) {
	r = &RadarEmailTimeseryThreatCategoryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified in Threat Categories over time.
func (r *RadarEmailTimeseryThreatCategoryService) List(ctx context.Context, query RadarEmailTimeseryThreatCategoryListParams, opts ...option.RequestOption) (res *RadarEmailTimeseryThreatCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/timeseries/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTimeseryThreatCategoryListResponse struct {
	Result  RadarEmailTimeseryThreatCategoryListResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailTimeseryThreatCategoryListResponseJSON   `json:"-"`
}

// radarEmailTimeseryThreatCategoryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailTimeseryThreatCategoryListResponse]
type radarEmailTimeseryThreatCategoryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryThreatCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryThreatCategoryListResponseResult struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 RadarEmailTimeseryThreatCategoryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailTimeseryThreatCategoryListResponseResultJSON   `json:"-"`
}

// radarEmailTimeseryThreatCategoryListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailTimeseryThreatCategoryListResponseResult]
type radarEmailTimeseryThreatCategoryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTimeseryThreatCategoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryThreatCategoryListResponseResultSerie0 struct {
	BrandImpersonation  []string                                                     `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                                     `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                                     `json:"IdentityDeception,required"`
	Link                []string                                                     `json:"Link,required"`
	JSON                radarEmailTimeseryThreatCategoryListResponseResultSerie0JSON `json:"-"`
}

// radarEmailTimeseryThreatCategoryListResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailTimeseryThreatCategoryListResponseResultSerie0]
type radarEmailTimeseryThreatCategoryListResponseResultSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTimeseryThreatCategoryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTimeseryThreatCategoryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailTimeseryThreatCategoryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTimeseryThreatCategoryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTimeseryThreatCategoryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTimeseryThreatCategoryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTimeseryThreatCategoryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTimeseryThreatCategoryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTimeseryThreatCategoryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTimeseryThreatCategoryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailTimeseryThreatCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailTimeseryThreatCategoryListParamsAggInterval string

const (
	RadarEmailTimeseryThreatCategoryListParamsAggInterval15m RadarEmailTimeseryThreatCategoryListParamsAggInterval = "15m"
	RadarEmailTimeseryThreatCategoryListParamsAggInterval1h  RadarEmailTimeseryThreatCategoryListParamsAggInterval = "1h"
	RadarEmailTimeseryThreatCategoryListParamsAggInterval1d  RadarEmailTimeseryThreatCategoryListParamsAggInterval = "1d"
	RadarEmailTimeseryThreatCategoryListParamsAggInterval1w  RadarEmailTimeseryThreatCategoryListParamsAggInterval = "1w"
)

type RadarEmailTimeseryThreatCategoryListParamsArc string

const (
	RadarEmailTimeseryThreatCategoryListParamsArcPass RadarEmailTimeseryThreatCategoryListParamsArc = "PASS"
	RadarEmailTimeseryThreatCategoryListParamsArcNone RadarEmailTimeseryThreatCategoryListParamsArc = "NONE"
	RadarEmailTimeseryThreatCategoryListParamsArcFail RadarEmailTimeseryThreatCategoryListParamsArc = "FAIL"
)

type RadarEmailTimeseryThreatCategoryListParamsDateRange string

const (
	RadarEmailTimeseryThreatCategoryListParamsDateRange1d         RadarEmailTimeseryThreatCategoryListParamsDateRange = "1d"
	RadarEmailTimeseryThreatCategoryListParamsDateRange7d         RadarEmailTimeseryThreatCategoryListParamsDateRange = "7d"
	RadarEmailTimeseryThreatCategoryListParamsDateRange14d        RadarEmailTimeseryThreatCategoryListParamsDateRange = "14d"
	RadarEmailTimeseryThreatCategoryListParamsDateRange28d        RadarEmailTimeseryThreatCategoryListParamsDateRange = "28d"
	RadarEmailTimeseryThreatCategoryListParamsDateRange12w        RadarEmailTimeseryThreatCategoryListParamsDateRange = "12w"
	RadarEmailTimeseryThreatCategoryListParamsDateRange24w        RadarEmailTimeseryThreatCategoryListParamsDateRange = "24w"
	RadarEmailTimeseryThreatCategoryListParamsDateRange52w        RadarEmailTimeseryThreatCategoryListParamsDateRange = "52w"
	RadarEmailTimeseryThreatCategoryListParamsDateRange1dControl  RadarEmailTimeseryThreatCategoryListParamsDateRange = "1dControl"
	RadarEmailTimeseryThreatCategoryListParamsDateRange7dControl  RadarEmailTimeseryThreatCategoryListParamsDateRange = "7dControl"
	RadarEmailTimeseryThreatCategoryListParamsDateRange14dControl RadarEmailTimeseryThreatCategoryListParamsDateRange = "14dControl"
	RadarEmailTimeseryThreatCategoryListParamsDateRange28dControl RadarEmailTimeseryThreatCategoryListParamsDateRange = "28dControl"
	RadarEmailTimeseryThreatCategoryListParamsDateRange12wControl RadarEmailTimeseryThreatCategoryListParamsDateRange = "12wControl"
	RadarEmailTimeseryThreatCategoryListParamsDateRange24wControl RadarEmailTimeseryThreatCategoryListParamsDateRange = "24wControl"
)

type RadarEmailTimeseryThreatCategoryListParamsDkim string

const (
	RadarEmailTimeseryThreatCategoryListParamsDkimPass RadarEmailTimeseryThreatCategoryListParamsDkim = "PASS"
	RadarEmailTimeseryThreatCategoryListParamsDkimNone RadarEmailTimeseryThreatCategoryListParamsDkim = "NONE"
	RadarEmailTimeseryThreatCategoryListParamsDkimFail RadarEmailTimeseryThreatCategoryListParamsDkim = "FAIL"
)

type RadarEmailTimeseryThreatCategoryListParamsDmarc string

const (
	RadarEmailTimeseryThreatCategoryListParamsDmarcPass RadarEmailTimeseryThreatCategoryListParamsDmarc = "PASS"
	RadarEmailTimeseryThreatCategoryListParamsDmarcNone RadarEmailTimeseryThreatCategoryListParamsDmarc = "NONE"
	RadarEmailTimeseryThreatCategoryListParamsDmarcFail RadarEmailTimeseryThreatCategoryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTimeseryThreatCategoryListParamsFormat string

const (
	RadarEmailTimeseryThreatCategoryListParamsFormatJson RadarEmailTimeseryThreatCategoryListParamsFormat = "JSON"
	RadarEmailTimeseryThreatCategoryListParamsFormatCsv  RadarEmailTimeseryThreatCategoryListParamsFormat = "CSV"
)

type RadarEmailTimeseryThreatCategoryListParamsSpf string

const (
	RadarEmailTimeseryThreatCategoryListParamsSpfPass RadarEmailTimeseryThreatCategoryListParamsSpf = "PASS"
	RadarEmailTimeseryThreatCategoryListParamsSpfNone RadarEmailTimeseryThreatCategoryListParamsSpf = "NONE"
	RadarEmailTimeseryThreatCategoryListParamsSpfFail RadarEmailTimeseryThreatCategoryListParamsSpf = "FAIL"
)
