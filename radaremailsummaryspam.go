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

// RadarEmailSummarySpamService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSummarySpamService]
// method instead.
type RadarEmailSummarySpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummarySpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSummarySpamService(opts ...option.RequestOption) (r *RadarEmailSummarySpamService) {
	r = &RadarEmailSummarySpamService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as SPAM.
func (r *RadarEmailSummarySpamService) List(ctx context.Context, query RadarEmailSummarySpamListParams, opts ...option.RequestOption) (res *RadarEmailSummarySpamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummarySpamListResponse struct {
	Result  RadarEmailSummarySpamListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarEmailSummarySpamListResponseJSON   `json:"-"`
}

// radarEmailSummarySpamListResponseJSON contains the JSON metadata for the struct
// [RadarEmailSummarySpamListResponse]
type radarEmailSummarySpamListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummarySpamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummarySpamListResponseResult struct {
	Meta     interface{}                                     `json:"meta,required"`
	Summary0 RadarEmailSummarySpamListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummarySpamListResponseResultJSON     `json:"-"`
}

// radarEmailSummarySpamListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSummarySpamListResponseResult]
type radarEmailSummarySpamListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummarySpamListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummarySpamListResponseResultSummary0 struct {
	NotSpam string                                              `json:"NOT_SPAM,required"`
	Spam    string                                              `json:"SPAM,required"`
	JSON    radarEmailSummarySpamListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummarySpamListResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSummarySpamListResponseResultSummary0]
type radarEmailSummarySpamListResponseResultSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummarySpamListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummarySpamListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSummarySpamListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummarySpamListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSummarySpamListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSummarySpamListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummarySpamListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSummarySpamListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSummarySpamListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSummarySpamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummarySpamListParamsArc string

const (
	RadarEmailSummarySpamListParamsArcPass RadarEmailSummarySpamListParamsArc = "PASS"
	RadarEmailSummarySpamListParamsArcNone RadarEmailSummarySpamListParamsArc = "NONE"
	RadarEmailSummarySpamListParamsArcFail RadarEmailSummarySpamListParamsArc = "FAIL"
)

type RadarEmailSummarySpamListParamsDateRange string

const (
	RadarEmailSummarySpamListParamsDateRange1d         RadarEmailSummarySpamListParamsDateRange = "1d"
	RadarEmailSummarySpamListParamsDateRange7d         RadarEmailSummarySpamListParamsDateRange = "7d"
	RadarEmailSummarySpamListParamsDateRange14d        RadarEmailSummarySpamListParamsDateRange = "14d"
	RadarEmailSummarySpamListParamsDateRange28d        RadarEmailSummarySpamListParamsDateRange = "28d"
	RadarEmailSummarySpamListParamsDateRange12w        RadarEmailSummarySpamListParamsDateRange = "12w"
	RadarEmailSummarySpamListParamsDateRange24w        RadarEmailSummarySpamListParamsDateRange = "24w"
	RadarEmailSummarySpamListParamsDateRange52w        RadarEmailSummarySpamListParamsDateRange = "52w"
	RadarEmailSummarySpamListParamsDateRange1dControl  RadarEmailSummarySpamListParamsDateRange = "1dControl"
	RadarEmailSummarySpamListParamsDateRange7dControl  RadarEmailSummarySpamListParamsDateRange = "7dControl"
	RadarEmailSummarySpamListParamsDateRange14dControl RadarEmailSummarySpamListParamsDateRange = "14dControl"
	RadarEmailSummarySpamListParamsDateRange28dControl RadarEmailSummarySpamListParamsDateRange = "28dControl"
	RadarEmailSummarySpamListParamsDateRange12wControl RadarEmailSummarySpamListParamsDateRange = "12wControl"
	RadarEmailSummarySpamListParamsDateRange24wControl RadarEmailSummarySpamListParamsDateRange = "24wControl"
)

type RadarEmailSummarySpamListParamsDkim string

const (
	RadarEmailSummarySpamListParamsDkimPass RadarEmailSummarySpamListParamsDkim = "PASS"
	RadarEmailSummarySpamListParamsDkimNone RadarEmailSummarySpamListParamsDkim = "NONE"
	RadarEmailSummarySpamListParamsDkimFail RadarEmailSummarySpamListParamsDkim = "FAIL"
)

type RadarEmailSummarySpamListParamsDmarc string

const (
	RadarEmailSummarySpamListParamsDmarcPass RadarEmailSummarySpamListParamsDmarc = "PASS"
	RadarEmailSummarySpamListParamsDmarcNone RadarEmailSummarySpamListParamsDmarc = "NONE"
	RadarEmailSummarySpamListParamsDmarcFail RadarEmailSummarySpamListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSummarySpamListParamsFormat string

const (
	RadarEmailSummarySpamListParamsFormatJson RadarEmailSummarySpamListParamsFormat = "JSON"
	RadarEmailSummarySpamListParamsFormatCsv  RadarEmailSummarySpamListParamsFormat = "CSV"
)

type RadarEmailSummarySpamListParamsSpf string

const (
	RadarEmailSummarySpamListParamsSpfPass RadarEmailSummarySpamListParamsSpf = "PASS"
	RadarEmailSummarySpamListParamsSpfNone RadarEmailSummarySpamListParamsSpf = "NONE"
	RadarEmailSummarySpamListParamsSpfFail RadarEmailSummarySpamListParamsSpf = "FAIL"
)
