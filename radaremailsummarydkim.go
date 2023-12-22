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

// RadarEmailSummaryDkimService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSummaryDkimService]
// method instead.
type RadarEmailSummaryDkimService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummaryDkimService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSummaryDkimService(opts ...option.RequestOption) (r *RadarEmailSummaryDkimService) {
	r = &RadarEmailSummaryDkimService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DKIM validation.
func (r *RadarEmailSummaryDkimService) List(ctx context.Context, query RadarEmailSummaryDkimListParams, opts ...option.RequestOption) (res *RadarEmailSummaryDkimListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummaryDkimListResponse struct {
	Result  RadarEmailSummaryDkimListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarEmailSummaryDkimListResponseJSON   `json:"-"`
}

// radarEmailSummaryDkimListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSummaryDkimListResponse]
type radarEmailSummaryDkimListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryDkimListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryDkimListResponseResult struct {
	Meta     interface{}                                     `json:"meta,required"`
	Summary0 RadarEmailSummaryDkimListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummaryDkimListResponseResultJSON     `json:"-"`
}

// radarEmailSummaryDkimListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSummaryDkimListResponseResult]
type radarEmailSummaryDkimListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryDkimListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryDkimListResponseResultSummary0 struct {
	Fail string                                              `json:"FAIL,required"`
	None string                                              `json:"NONE,required"`
	Pass string                                              `json:"PASS,required"`
	JSON radarEmailSummaryDkimListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummaryDkimListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSummaryDkimListResponseResultSummary0]
type radarEmailSummaryDkimListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryDkimListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryDkimListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSummaryDkimListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummaryDkimListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSummaryDkimListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummaryDkimListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSummaryDkimListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSummaryDkimListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSummaryDkimListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummaryDkimListParamsArc string

const (
	RadarEmailSummaryDkimListParamsArcPass RadarEmailSummaryDkimListParamsArc = "PASS"
	RadarEmailSummaryDkimListParamsArcNone RadarEmailSummaryDkimListParamsArc = "NONE"
	RadarEmailSummaryDkimListParamsArcFail RadarEmailSummaryDkimListParamsArc = "FAIL"
)

type RadarEmailSummaryDkimListParamsDateRange string

const (
	RadarEmailSummaryDkimListParamsDateRange1d         RadarEmailSummaryDkimListParamsDateRange = "1d"
	RadarEmailSummaryDkimListParamsDateRange7d         RadarEmailSummaryDkimListParamsDateRange = "7d"
	RadarEmailSummaryDkimListParamsDateRange14d        RadarEmailSummaryDkimListParamsDateRange = "14d"
	RadarEmailSummaryDkimListParamsDateRange28d        RadarEmailSummaryDkimListParamsDateRange = "28d"
	RadarEmailSummaryDkimListParamsDateRange12w        RadarEmailSummaryDkimListParamsDateRange = "12w"
	RadarEmailSummaryDkimListParamsDateRange24w        RadarEmailSummaryDkimListParamsDateRange = "24w"
	RadarEmailSummaryDkimListParamsDateRange52w        RadarEmailSummaryDkimListParamsDateRange = "52w"
	RadarEmailSummaryDkimListParamsDateRange1dControl  RadarEmailSummaryDkimListParamsDateRange = "1dControl"
	RadarEmailSummaryDkimListParamsDateRange7dControl  RadarEmailSummaryDkimListParamsDateRange = "7dControl"
	RadarEmailSummaryDkimListParamsDateRange14dControl RadarEmailSummaryDkimListParamsDateRange = "14dControl"
	RadarEmailSummaryDkimListParamsDateRange28dControl RadarEmailSummaryDkimListParamsDateRange = "28dControl"
	RadarEmailSummaryDkimListParamsDateRange12wControl RadarEmailSummaryDkimListParamsDateRange = "12wControl"
	RadarEmailSummaryDkimListParamsDateRange24wControl RadarEmailSummaryDkimListParamsDateRange = "24wControl"
)

type RadarEmailSummaryDkimListParamsDmarc string

const (
	RadarEmailSummaryDkimListParamsDmarcPass RadarEmailSummaryDkimListParamsDmarc = "PASS"
	RadarEmailSummaryDkimListParamsDmarcNone RadarEmailSummaryDkimListParamsDmarc = "NONE"
	RadarEmailSummaryDkimListParamsDmarcFail RadarEmailSummaryDkimListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSummaryDkimListParamsFormat string

const (
	RadarEmailSummaryDkimListParamsFormatJson RadarEmailSummaryDkimListParamsFormat = "JSON"
	RadarEmailSummaryDkimListParamsFormatCsv  RadarEmailSummaryDkimListParamsFormat = "CSV"
)

type RadarEmailSummaryDkimListParamsSpf string

const (
	RadarEmailSummaryDkimListParamsSpfPass RadarEmailSummaryDkimListParamsSpf = "PASS"
	RadarEmailSummaryDkimListParamsSpfNone RadarEmailSummaryDkimListParamsSpf = "NONE"
	RadarEmailSummaryDkimListParamsSpfFail RadarEmailSummaryDkimListParamsSpf = "FAIL"
)
