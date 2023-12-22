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

// RadarEmailSummarySpfService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSummarySpfService]
// method instead.
type RadarEmailSummarySpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummarySpfService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSummarySpfService(opts ...option.RequestOption) (r *RadarEmailSummarySpfService) {
	r = &RadarEmailSummarySpfService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation.
func (r *RadarEmailSummarySpfService) List(ctx context.Context, query RadarEmailSummarySpfListParams, opts ...option.RequestOption) (res *RadarEmailSummarySpfListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummarySpfListResponse struct {
	Result  RadarEmailSummarySpfListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarEmailSummarySpfListResponseJSON   `json:"-"`
}

// radarEmailSummarySpfListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSummarySpfListResponse]
type radarEmailSummarySpfListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummarySpfListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummarySpfListResponseResult struct {
	Meta     interface{}                                    `json:"meta,required"`
	Summary0 RadarEmailSummarySpfListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummarySpfListResponseResultJSON     `json:"-"`
}

// radarEmailSummarySpfListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSummarySpfListResponseResult]
type radarEmailSummarySpfListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummarySpfListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummarySpfListResponseResultSummary0 struct {
	Fail string                                             `json:"FAIL,required"`
	None string                                             `json:"NONE,required"`
	Pass string                                             `json:"PASS,required"`
	JSON radarEmailSummarySpfListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummarySpfListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSummarySpfListResponseResultSummary0]
type radarEmailSummarySpfListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummarySpfListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummarySpfListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSummarySpfListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummarySpfListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSummarySpfListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSummarySpfListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummarySpfListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSummarySpfListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSummarySpfListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummarySpfListParamsArc string

const (
	RadarEmailSummarySpfListParamsArcPass RadarEmailSummarySpfListParamsArc = "PASS"
	RadarEmailSummarySpfListParamsArcNone RadarEmailSummarySpfListParamsArc = "NONE"
	RadarEmailSummarySpfListParamsArcFail RadarEmailSummarySpfListParamsArc = "FAIL"
)

type RadarEmailSummarySpfListParamsDateRange string

const (
	RadarEmailSummarySpfListParamsDateRange1d         RadarEmailSummarySpfListParamsDateRange = "1d"
	RadarEmailSummarySpfListParamsDateRange7d         RadarEmailSummarySpfListParamsDateRange = "7d"
	RadarEmailSummarySpfListParamsDateRange14d        RadarEmailSummarySpfListParamsDateRange = "14d"
	RadarEmailSummarySpfListParamsDateRange28d        RadarEmailSummarySpfListParamsDateRange = "28d"
	RadarEmailSummarySpfListParamsDateRange12w        RadarEmailSummarySpfListParamsDateRange = "12w"
	RadarEmailSummarySpfListParamsDateRange24w        RadarEmailSummarySpfListParamsDateRange = "24w"
	RadarEmailSummarySpfListParamsDateRange52w        RadarEmailSummarySpfListParamsDateRange = "52w"
	RadarEmailSummarySpfListParamsDateRange1dControl  RadarEmailSummarySpfListParamsDateRange = "1dControl"
	RadarEmailSummarySpfListParamsDateRange7dControl  RadarEmailSummarySpfListParamsDateRange = "7dControl"
	RadarEmailSummarySpfListParamsDateRange14dControl RadarEmailSummarySpfListParamsDateRange = "14dControl"
	RadarEmailSummarySpfListParamsDateRange28dControl RadarEmailSummarySpfListParamsDateRange = "28dControl"
	RadarEmailSummarySpfListParamsDateRange12wControl RadarEmailSummarySpfListParamsDateRange = "12wControl"
	RadarEmailSummarySpfListParamsDateRange24wControl RadarEmailSummarySpfListParamsDateRange = "24wControl"
)

type RadarEmailSummarySpfListParamsDkim string

const (
	RadarEmailSummarySpfListParamsDkimPass RadarEmailSummarySpfListParamsDkim = "PASS"
	RadarEmailSummarySpfListParamsDkimNone RadarEmailSummarySpfListParamsDkim = "NONE"
	RadarEmailSummarySpfListParamsDkimFail RadarEmailSummarySpfListParamsDkim = "FAIL"
)

type RadarEmailSummarySpfListParamsDmarc string

const (
	RadarEmailSummarySpfListParamsDmarcPass RadarEmailSummarySpfListParamsDmarc = "PASS"
	RadarEmailSummarySpfListParamsDmarcNone RadarEmailSummarySpfListParamsDmarc = "NONE"
	RadarEmailSummarySpfListParamsDmarcFail RadarEmailSummarySpfListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSummarySpfListParamsFormat string

const (
	RadarEmailSummarySpfListParamsFormatJson RadarEmailSummarySpfListParamsFormat = "JSON"
	RadarEmailSummarySpfListParamsFormatCsv  RadarEmailSummarySpfListParamsFormat = "CSV"
)
