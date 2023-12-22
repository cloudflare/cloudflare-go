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

// RadarEmailSummaryArcService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSummaryArcService]
// method instead.
type RadarEmailSummaryArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummaryArcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSummaryArcService(opts ...option.RequestOption) (r *RadarEmailSummaryArcService) {
	r = &RadarEmailSummaryArcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation.
func (r *RadarEmailSummaryArcService) List(ctx context.Context, query RadarEmailSummaryArcListParams, opts ...option.RequestOption) (res *RadarEmailSummaryArcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummaryArcListResponse struct {
	Result  RadarEmailSummaryArcListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarEmailSummaryArcListResponseJSON   `json:"-"`
}

// radarEmailSummaryArcListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSummaryArcListResponse]
type radarEmailSummaryArcListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryArcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryArcListResponseResult struct {
	Meta     interface{}                                    `json:"meta,required"`
	Summary0 RadarEmailSummaryArcListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummaryArcListResponseResultJSON     `json:"-"`
}

// radarEmailSummaryArcListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSummaryArcListResponseResult]
type radarEmailSummaryArcListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryArcListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryArcListResponseResultSummary0 struct {
	Fail string                                             `json:"FAIL,required"`
	None string                                             `json:"NONE,required"`
	Pass string                                             `json:"PASS,required"`
	JSON radarEmailSummaryArcListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummaryArcListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSummaryArcListResponseResultSummary0]
type radarEmailSummaryArcListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryArcListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryArcListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummaryArcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSummaryArcListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSummaryArcListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummaryArcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSummaryArcListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSummaryArcListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSummaryArcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummaryArcListParamsDateRange string

const (
	RadarEmailSummaryArcListParamsDateRange1d         RadarEmailSummaryArcListParamsDateRange = "1d"
	RadarEmailSummaryArcListParamsDateRange7d         RadarEmailSummaryArcListParamsDateRange = "7d"
	RadarEmailSummaryArcListParamsDateRange14d        RadarEmailSummaryArcListParamsDateRange = "14d"
	RadarEmailSummaryArcListParamsDateRange28d        RadarEmailSummaryArcListParamsDateRange = "28d"
	RadarEmailSummaryArcListParamsDateRange12w        RadarEmailSummaryArcListParamsDateRange = "12w"
	RadarEmailSummaryArcListParamsDateRange24w        RadarEmailSummaryArcListParamsDateRange = "24w"
	RadarEmailSummaryArcListParamsDateRange52w        RadarEmailSummaryArcListParamsDateRange = "52w"
	RadarEmailSummaryArcListParamsDateRange1dControl  RadarEmailSummaryArcListParamsDateRange = "1dControl"
	RadarEmailSummaryArcListParamsDateRange7dControl  RadarEmailSummaryArcListParamsDateRange = "7dControl"
	RadarEmailSummaryArcListParamsDateRange14dControl RadarEmailSummaryArcListParamsDateRange = "14dControl"
	RadarEmailSummaryArcListParamsDateRange28dControl RadarEmailSummaryArcListParamsDateRange = "28dControl"
	RadarEmailSummaryArcListParamsDateRange12wControl RadarEmailSummaryArcListParamsDateRange = "12wControl"
	RadarEmailSummaryArcListParamsDateRange24wControl RadarEmailSummaryArcListParamsDateRange = "24wControl"
)

type RadarEmailSummaryArcListParamsDkim string

const (
	RadarEmailSummaryArcListParamsDkimPass RadarEmailSummaryArcListParamsDkim = "PASS"
	RadarEmailSummaryArcListParamsDkimNone RadarEmailSummaryArcListParamsDkim = "NONE"
	RadarEmailSummaryArcListParamsDkimFail RadarEmailSummaryArcListParamsDkim = "FAIL"
)

type RadarEmailSummaryArcListParamsDmarc string

const (
	RadarEmailSummaryArcListParamsDmarcPass RadarEmailSummaryArcListParamsDmarc = "PASS"
	RadarEmailSummaryArcListParamsDmarcNone RadarEmailSummaryArcListParamsDmarc = "NONE"
	RadarEmailSummaryArcListParamsDmarcFail RadarEmailSummaryArcListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSummaryArcListParamsFormat string

const (
	RadarEmailSummaryArcListParamsFormatJson RadarEmailSummaryArcListParamsFormat = "JSON"
	RadarEmailSummaryArcListParamsFormatCsv  RadarEmailSummaryArcListParamsFormat = "CSV"
)

type RadarEmailSummaryArcListParamsSpf string

const (
	RadarEmailSummaryArcListParamsSpfPass RadarEmailSummaryArcListParamsSpf = "PASS"
	RadarEmailSummaryArcListParamsSpfNone RadarEmailSummaryArcListParamsSpf = "NONE"
	RadarEmailSummaryArcListParamsSpfFail RadarEmailSummaryArcListParamsSpf = "FAIL"
)
