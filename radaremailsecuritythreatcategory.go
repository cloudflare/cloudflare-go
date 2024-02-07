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

// RadarEmailSecurityThreatCategoryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityThreatCategoryService] method instead.
type RadarEmailSecurityThreatCategoryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityThreatCategoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityThreatCategoryService(opts ...option.RequestOption) (r *RadarEmailSecurityThreatCategoryService) {
	r = &RadarEmailSecurityThreatCategoryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified in Threat Categories over time.
func (r *RadarEmailSecurityThreatCategoryService) List(ctx context.Context, query RadarEmailSecurityThreatCategoryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityThreatCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityThreatCategoryListResponseEnvelope
	path := "radar/email/security/timeseries_groups/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityThreatCategoryListResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarEmailSecurityThreatCategoryListResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityThreatCategoryListResponseJSON   `json:"-"`
}

// radarEmailSecurityThreatCategoryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityThreatCategoryListResponse]
type radarEmailSecurityThreatCategoryListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategoryListResponseSerie0 struct {
	BrandImpersonation  []string                                               `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                               `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                               `json:"IdentityDeception,required"`
	Link                []string                                               `json:"Link,required"`
	JSON                radarEmailSecurityThreatCategoryListResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityThreatCategoryListResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityThreatCategoryListResponseSerie0]
type radarEmailSecurityThreatCategoryListResponseSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategoryListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategoryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityThreatCategoryListParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityThreatCategoryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityThreatCategoryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityThreatCategoryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityThreatCategoryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityThreatCategoryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityThreatCategoryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityThreatCategoryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityThreatCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityThreatCategoryListParamsAggInterval string

const (
	RadarEmailSecurityThreatCategoryListParamsAggInterval15m RadarEmailSecurityThreatCategoryListParamsAggInterval = "15m"
	RadarEmailSecurityThreatCategoryListParamsAggInterval1h  RadarEmailSecurityThreatCategoryListParamsAggInterval = "1h"
	RadarEmailSecurityThreatCategoryListParamsAggInterval1d  RadarEmailSecurityThreatCategoryListParamsAggInterval = "1d"
	RadarEmailSecurityThreatCategoryListParamsAggInterval1w  RadarEmailSecurityThreatCategoryListParamsAggInterval = "1w"
)

type RadarEmailSecurityThreatCategoryListParamsArc string

const (
	RadarEmailSecurityThreatCategoryListParamsArcPass RadarEmailSecurityThreatCategoryListParamsArc = "PASS"
	RadarEmailSecurityThreatCategoryListParamsArcNone RadarEmailSecurityThreatCategoryListParamsArc = "NONE"
	RadarEmailSecurityThreatCategoryListParamsArcFail RadarEmailSecurityThreatCategoryListParamsArc = "FAIL"
)

type RadarEmailSecurityThreatCategoryListParamsDateRange string

const (
	RadarEmailSecurityThreatCategoryListParamsDateRange1d         RadarEmailSecurityThreatCategoryListParamsDateRange = "1d"
	RadarEmailSecurityThreatCategoryListParamsDateRange2d         RadarEmailSecurityThreatCategoryListParamsDateRange = "2d"
	RadarEmailSecurityThreatCategoryListParamsDateRange7d         RadarEmailSecurityThreatCategoryListParamsDateRange = "7d"
	RadarEmailSecurityThreatCategoryListParamsDateRange14d        RadarEmailSecurityThreatCategoryListParamsDateRange = "14d"
	RadarEmailSecurityThreatCategoryListParamsDateRange28d        RadarEmailSecurityThreatCategoryListParamsDateRange = "28d"
	RadarEmailSecurityThreatCategoryListParamsDateRange12w        RadarEmailSecurityThreatCategoryListParamsDateRange = "12w"
	RadarEmailSecurityThreatCategoryListParamsDateRange24w        RadarEmailSecurityThreatCategoryListParamsDateRange = "24w"
	RadarEmailSecurityThreatCategoryListParamsDateRange52w        RadarEmailSecurityThreatCategoryListParamsDateRange = "52w"
	RadarEmailSecurityThreatCategoryListParamsDateRange1dControl  RadarEmailSecurityThreatCategoryListParamsDateRange = "1dControl"
	RadarEmailSecurityThreatCategoryListParamsDateRange2dControl  RadarEmailSecurityThreatCategoryListParamsDateRange = "2dControl"
	RadarEmailSecurityThreatCategoryListParamsDateRange7dControl  RadarEmailSecurityThreatCategoryListParamsDateRange = "7dControl"
	RadarEmailSecurityThreatCategoryListParamsDateRange14dControl RadarEmailSecurityThreatCategoryListParamsDateRange = "14dControl"
	RadarEmailSecurityThreatCategoryListParamsDateRange28dControl RadarEmailSecurityThreatCategoryListParamsDateRange = "28dControl"
	RadarEmailSecurityThreatCategoryListParamsDateRange12wControl RadarEmailSecurityThreatCategoryListParamsDateRange = "12wControl"
	RadarEmailSecurityThreatCategoryListParamsDateRange24wControl RadarEmailSecurityThreatCategoryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityThreatCategoryListParamsDkim string

const (
	RadarEmailSecurityThreatCategoryListParamsDkimPass RadarEmailSecurityThreatCategoryListParamsDkim = "PASS"
	RadarEmailSecurityThreatCategoryListParamsDkimNone RadarEmailSecurityThreatCategoryListParamsDkim = "NONE"
	RadarEmailSecurityThreatCategoryListParamsDkimFail RadarEmailSecurityThreatCategoryListParamsDkim = "FAIL"
)

type RadarEmailSecurityThreatCategoryListParamsDmarc string

const (
	RadarEmailSecurityThreatCategoryListParamsDmarcPass RadarEmailSecurityThreatCategoryListParamsDmarc = "PASS"
	RadarEmailSecurityThreatCategoryListParamsDmarcNone RadarEmailSecurityThreatCategoryListParamsDmarc = "NONE"
	RadarEmailSecurityThreatCategoryListParamsDmarcFail RadarEmailSecurityThreatCategoryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityThreatCategoryListParamsFormat string

const (
	RadarEmailSecurityThreatCategoryListParamsFormatJson RadarEmailSecurityThreatCategoryListParamsFormat = "JSON"
	RadarEmailSecurityThreatCategoryListParamsFormatCsv  RadarEmailSecurityThreatCategoryListParamsFormat = "CSV"
)

type RadarEmailSecurityThreatCategoryListParamsSpf string

const (
	RadarEmailSecurityThreatCategoryListParamsSpfPass RadarEmailSecurityThreatCategoryListParamsSpf = "PASS"
	RadarEmailSecurityThreatCategoryListParamsSpfNone RadarEmailSecurityThreatCategoryListParamsSpf = "NONE"
	RadarEmailSecurityThreatCategoryListParamsSpfFail RadarEmailSecurityThreatCategoryListParamsSpf = "FAIL"
)

type RadarEmailSecurityThreatCategoryListResponseEnvelope struct {
	Result  RadarEmailSecurityThreatCategoryListResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityThreatCategoryListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityThreatCategoryListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityThreatCategoryListResponseEnvelope]
type radarEmailSecurityThreatCategoryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
