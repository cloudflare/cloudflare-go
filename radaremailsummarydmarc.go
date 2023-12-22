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

// RadarEmailSummaryDmarcService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSummaryDmarcService]
// method instead.
type RadarEmailSummaryDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummaryDmarcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSummaryDmarcService(opts ...option.RequestOption) (r *RadarEmailSummaryDmarcService) {
	r = &RadarEmailSummaryDmarcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DMARC validation.
func (r *RadarEmailSummaryDmarcService) List(ctx context.Context, query RadarEmailSummaryDmarcListParams, opts ...option.RequestOption) (res *RadarEmailSummaryDmarcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummaryDmarcListResponse struct {
	Result  RadarEmailSummaryDmarcListResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarEmailSummaryDmarcListResponseJSON   `json:"-"`
}

// radarEmailSummaryDmarcListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSummaryDmarcListResponse]
type radarEmailSummaryDmarcListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryDmarcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryDmarcListResponseResult struct {
	Meta     interface{}                                      `json:"meta,required"`
	Summary0 RadarEmailSummaryDmarcListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummaryDmarcListResponseResultJSON     `json:"-"`
}

// radarEmailSummaryDmarcListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSummaryDmarcListResponseResult]
type radarEmailSummaryDmarcListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryDmarcListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryDmarcListResponseResultSummary0 struct {
	Fail string                                               `json:"FAIL,required"`
	None string                                               `json:"NONE,required"`
	Pass string                                               `json:"PASS,required"`
	JSON radarEmailSummaryDmarcListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummaryDmarcListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSummaryDmarcListResponseResultSummary0]
type radarEmailSummaryDmarcListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryDmarcListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryDmarcListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSummaryDmarcListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummaryDmarcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSummaryDmarcListParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummaryDmarcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSummaryDmarcListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSummaryDmarcListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSummaryDmarcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummaryDmarcListParamsArc string

const (
	RadarEmailSummaryDmarcListParamsArcPass RadarEmailSummaryDmarcListParamsArc = "PASS"
	RadarEmailSummaryDmarcListParamsArcNone RadarEmailSummaryDmarcListParamsArc = "NONE"
	RadarEmailSummaryDmarcListParamsArcFail RadarEmailSummaryDmarcListParamsArc = "FAIL"
)

type RadarEmailSummaryDmarcListParamsDateRange string

const (
	RadarEmailSummaryDmarcListParamsDateRange1d         RadarEmailSummaryDmarcListParamsDateRange = "1d"
	RadarEmailSummaryDmarcListParamsDateRange7d         RadarEmailSummaryDmarcListParamsDateRange = "7d"
	RadarEmailSummaryDmarcListParamsDateRange14d        RadarEmailSummaryDmarcListParamsDateRange = "14d"
	RadarEmailSummaryDmarcListParamsDateRange28d        RadarEmailSummaryDmarcListParamsDateRange = "28d"
	RadarEmailSummaryDmarcListParamsDateRange12w        RadarEmailSummaryDmarcListParamsDateRange = "12w"
	RadarEmailSummaryDmarcListParamsDateRange24w        RadarEmailSummaryDmarcListParamsDateRange = "24w"
	RadarEmailSummaryDmarcListParamsDateRange52w        RadarEmailSummaryDmarcListParamsDateRange = "52w"
	RadarEmailSummaryDmarcListParamsDateRange1dControl  RadarEmailSummaryDmarcListParamsDateRange = "1dControl"
	RadarEmailSummaryDmarcListParamsDateRange7dControl  RadarEmailSummaryDmarcListParamsDateRange = "7dControl"
	RadarEmailSummaryDmarcListParamsDateRange14dControl RadarEmailSummaryDmarcListParamsDateRange = "14dControl"
	RadarEmailSummaryDmarcListParamsDateRange28dControl RadarEmailSummaryDmarcListParamsDateRange = "28dControl"
	RadarEmailSummaryDmarcListParamsDateRange12wControl RadarEmailSummaryDmarcListParamsDateRange = "12wControl"
	RadarEmailSummaryDmarcListParamsDateRange24wControl RadarEmailSummaryDmarcListParamsDateRange = "24wControl"
)

type RadarEmailSummaryDmarcListParamsDkim string

const (
	RadarEmailSummaryDmarcListParamsDkimPass RadarEmailSummaryDmarcListParamsDkim = "PASS"
	RadarEmailSummaryDmarcListParamsDkimNone RadarEmailSummaryDmarcListParamsDkim = "NONE"
	RadarEmailSummaryDmarcListParamsDkimFail RadarEmailSummaryDmarcListParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailSummaryDmarcListParamsFormat string

const (
	RadarEmailSummaryDmarcListParamsFormatJson RadarEmailSummaryDmarcListParamsFormat = "JSON"
	RadarEmailSummaryDmarcListParamsFormatCsv  RadarEmailSummaryDmarcListParamsFormat = "CSV"
)

type RadarEmailSummaryDmarcListParamsSpf string

const (
	RadarEmailSummaryDmarcListParamsSpfPass RadarEmailSummaryDmarcListParamsSpf = "PASS"
	RadarEmailSummaryDmarcListParamsSpfNone RadarEmailSummaryDmarcListParamsSpf = "NONE"
	RadarEmailSummaryDmarcListParamsSpfFail RadarEmailSummaryDmarcListParamsSpf = "FAIL"
)
