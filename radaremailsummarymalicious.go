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

// RadarEmailSummaryMaliciousService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSummaryMaliciousService] method instead.
type RadarEmailSummaryMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummaryMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSummaryMaliciousService(opts ...option.RequestOption) (r *RadarEmailSummaryMaliciousService) {
	r = &RadarEmailSummaryMaliciousService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as MALICIOUS.
func (r *RadarEmailSummaryMaliciousService) List(ctx context.Context, query RadarEmailSummaryMaliciousListParams, opts ...option.RequestOption) (res *RadarEmailSummaryMaliciousListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummaryMaliciousListResponse struct {
	Result  RadarEmailSummaryMaliciousListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEmailSummaryMaliciousListResponseJSON   `json:"-"`
}

// radarEmailSummaryMaliciousListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSummaryMaliciousListResponse]
type radarEmailSummaryMaliciousListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryMaliciousListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryMaliciousListResponseResult struct {
	Meta     interface{}                                          `json:"meta,required"`
	Summary0 RadarEmailSummaryMaliciousListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummaryMaliciousListResponseResultJSON     `json:"-"`
}

// radarEmailSummaryMaliciousListResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSummaryMaliciousListResponseResult]
type radarEmailSummaryMaliciousListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryMaliciousListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryMaliciousListResponseResultSummary0 struct {
	Malicious    string                                                   `json:"MALICIOUS,required"`
	NotMalicious string                                                   `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSummaryMaliciousListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummaryMaliciousListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSummaryMaliciousListResponseResultSummary0]
type radarEmailSummaryMaliciousListResponseResultSummary0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSummaryMaliciousListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryMaliciousListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSummaryMaliciousListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummaryMaliciousListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSummaryMaliciousListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSummaryMaliciousListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummaryMaliciousListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSummaryMaliciousListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSummaryMaliciousListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSummaryMaliciousListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummaryMaliciousListParamsArc string

const (
	RadarEmailSummaryMaliciousListParamsArcPass RadarEmailSummaryMaliciousListParamsArc = "PASS"
	RadarEmailSummaryMaliciousListParamsArcNone RadarEmailSummaryMaliciousListParamsArc = "NONE"
	RadarEmailSummaryMaliciousListParamsArcFail RadarEmailSummaryMaliciousListParamsArc = "FAIL"
)

type RadarEmailSummaryMaliciousListParamsDateRange string

const (
	RadarEmailSummaryMaliciousListParamsDateRange1d         RadarEmailSummaryMaliciousListParamsDateRange = "1d"
	RadarEmailSummaryMaliciousListParamsDateRange7d         RadarEmailSummaryMaliciousListParamsDateRange = "7d"
	RadarEmailSummaryMaliciousListParamsDateRange14d        RadarEmailSummaryMaliciousListParamsDateRange = "14d"
	RadarEmailSummaryMaliciousListParamsDateRange28d        RadarEmailSummaryMaliciousListParamsDateRange = "28d"
	RadarEmailSummaryMaliciousListParamsDateRange12w        RadarEmailSummaryMaliciousListParamsDateRange = "12w"
	RadarEmailSummaryMaliciousListParamsDateRange24w        RadarEmailSummaryMaliciousListParamsDateRange = "24w"
	RadarEmailSummaryMaliciousListParamsDateRange52w        RadarEmailSummaryMaliciousListParamsDateRange = "52w"
	RadarEmailSummaryMaliciousListParamsDateRange1dControl  RadarEmailSummaryMaliciousListParamsDateRange = "1dControl"
	RadarEmailSummaryMaliciousListParamsDateRange7dControl  RadarEmailSummaryMaliciousListParamsDateRange = "7dControl"
	RadarEmailSummaryMaliciousListParamsDateRange14dControl RadarEmailSummaryMaliciousListParamsDateRange = "14dControl"
	RadarEmailSummaryMaliciousListParamsDateRange28dControl RadarEmailSummaryMaliciousListParamsDateRange = "28dControl"
	RadarEmailSummaryMaliciousListParamsDateRange12wControl RadarEmailSummaryMaliciousListParamsDateRange = "12wControl"
	RadarEmailSummaryMaliciousListParamsDateRange24wControl RadarEmailSummaryMaliciousListParamsDateRange = "24wControl"
)

type RadarEmailSummaryMaliciousListParamsDkim string

const (
	RadarEmailSummaryMaliciousListParamsDkimPass RadarEmailSummaryMaliciousListParamsDkim = "PASS"
	RadarEmailSummaryMaliciousListParamsDkimNone RadarEmailSummaryMaliciousListParamsDkim = "NONE"
	RadarEmailSummaryMaliciousListParamsDkimFail RadarEmailSummaryMaliciousListParamsDkim = "FAIL"
)

type RadarEmailSummaryMaliciousListParamsDmarc string

const (
	RadarEmailSummaryMaliciousListParamsDmarcPass RadarEmailSummaryMaliciousListParamsDmarc = "PASS"
	RadarEmailSummaryMaliciousListParamsDmarcNone RadarEmailSummaryMaliciousListParamsDmarc = "NONE"
	RadarEmailSummaryMaliciousListParamsDmarcFail RadarEmailSummaryMaliciousListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSummaryMaliciousListParamsFormat string

const (
	RadarEmailSummaryMaliciousListParamsFormatJson RadarEmailSummaryMaliciousListParamsFormat = "JSON"
	RadarEmailSummaryMaliciousListParamsFormatCsv  RadarEmailSummaryMaliciousListParamsFormat = "CSV"
)

type RadarEmailSummaryMaliciousListParamsSpf string

const (
	RadarEmailSummaryMaliciousListParamsSpfPass RadarEmailSummaryMaliciousListParamsSpf = "PASS"
	RadarEmailSummaryMaliciousListParamsSpfNone RadarEmailSummaryMaliciousListParamsSpf = "NONE"
	RadarEmailSummaryMaliciousListParamsSpfFail RadarEmailSummaryMaliciousListParamsSpf = "FAIL"
)
