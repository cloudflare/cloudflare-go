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

// RadarEmailSummaryThreatCategoryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSummaryThreatCategoryService] method instead.
type RadarEmailSummaryThreatCategoryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSummaryThreatCategoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSummaryThreatCategoryService(opts ...option.RequestOption) (r *RadarEmailSummaryThreatCategoryService) {
	r = &RadarEmailSummaryThreatCategoryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified in Threat Categories.
func (r *RadarEmailSummaryThreatCategoryService) List(ctx context.Context, query RadarEmailSummaryThreatCategoryListParams, opts ...option.RequestOption) (res *RadarEmailSummaryThreatCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSummaryThreatCategoryListResponse struct {
	Result  RadarEmailSummaryThreatCategoryListResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarEmailSummaryThreatCategoryListResponseJSON   `json:"-"`
}

// radarEmailSummaryThreatCategoryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSummaryThreatCategoryListResponse]
type radarEmailSummaryThreatCategoryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryThreatCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryThreatCategoryListResponseResult struct {
	Meta     interface{}                                               `json:"meta,required"`
	Summary0 RadarEmailSummaryThreatCategoryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSummaryThreatCategoryListResponseResultJSON     `json:"-"`
}

// radarEmailSummaryThreatCategoryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSummaryThreatCategoryListResponseResult]
type radarEmailSummaryThreatCategoryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummaryThreatCategoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryThreatCategoryListResponseResultSummary0 struct {
	BrandImpersonation  string                                                        `json:"BrandImpersonation,required"`
	CredentialHarvester string                                                        `json:"CredentialHarvester,required"`
	IdentityDeception   string                                                        `json:"IdentityDeception,required"`
	Link                string                                                        `json:"Link,required"`
	JSON                radarEmailSummaryThreatCategoryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSummaryThreatCategoryListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSummaryThreatCategoryListResponseResultSummary0]
type radarEmailSummaryThreatCategoryListResponseResultSummary0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSummaryThreatCategoryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSummaryThreatCategoryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSummaryThreatCategoryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSummaryThreatCategoryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSummaryThreatCategoryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSummaryThreatCategoryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSummaryThreatCategoryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSummaryThreatCategoryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSummaryThreatCategoryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSummaryThreatCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSummaryThreatCategoryListParamsArc string

const (
	RadarEmailSummaryThreatCategoryListParamsArcPass RadarEmailSummaryThreatCategoryListParamsArc = "PASS"
	RadarEmailSummaryThreatCategoryListParamsArcNone RadarEmailSummaryThreatCategoryListParamsArc = "NONE"
	RadarEmailSummaryThreatCategoryListParamsArcFail RadarEmailSummaryThreatCategoryListParamsArc = "FAIL"
)

type RadarEmailSummaryThreatCategoryListParamsDateRange string

const (
	RadarEmailSummaryThreatCategoryListParamsDateRange1d         RadarEmailSummaryThreatCategoryListParamsDateRange = "1d"
	RadarEmailSummaryThreatCategoryListParamsDateRange7d         RadarEmailSummaryThreatCategoryListParamsDateRange = "7d"
	RadarEmailSummaryThreatCategoryListParamsDateRange14d        RadarEmailSummaryThreatCategoryListParamsDateRange = "14d"
	RadarEmailSummaryThreatCategoryListParamsDateRange28d        RadarEmailSummaryThreatCategoryListParamsDateRange = "28d"
	RadarEmailSummaryThreatCategoryListParamsDateRange12w        RadarEmailSummaryThreatCategoryListParamsDateRange = "12w"
	RadarEmailSummaryThreatCategoryListParamsDateRange24w        RadarEmailSummaryThreatCategoryListParamsDateRange = "24w"
	RadarEmailSummaryThreatCategoryListParamsDateRange52w        RadarEmailSummaryThreatCategoryListParamsDateRange = "52w"
	RadarEmailSummaryThreatCategoryListParamsDateRange1dControl  RadarEmailSummaryThreatCategoryListParamsDateRange = "1dControl"
	RadarEmailSummaryThreatCategoryListParamsDateRange7dControl  RadarEmailSummaryThreatCategoryListParamsDateRange = "7dControl"
	RadarEmailSummaryThreatCategoryListParamsDateRange14dControl RadarEmailSummaryThreatCategoryListParamsDateRange = "14dControl"
	RadarEmailSummaryThreatCategoryListParamsDateRange28dControl RadarEmailSummaryThreatCategoryListParamsDateRange = "28dControl"
	RadarEmailSummaryThreatCategoryListParamsDateRange12wControl RadarEmailSummaryThreatCategoryListParamsDateRange = "12wControl"
	RadarEmailSummaryThreatCategoryListParamsDateRange24wControl RadarEmailSummaryThreatCategoryListParamsDateRange = "24wControl"
)

type RadarEmailSummaryThreatCategoryListParamsDkim string

const (
	RadarEmailSummaryThreatCategoryListParamsDkimPass RadarEmailSummaryThreatCategoryListParamsDkim = "PASS"
	RadarEmailSummaryThreatCategoryListParamsDkimNone RadarEmailSummaryThreatCategoryListParamsDkim = "NONE"
	RadarEmailSummaryThreatCategoryListParamsDkimFail RadarEmailSummaryThreatCategoryListParamsDkim = "FAIL"
)

type RadarEmailSummaryThreatCategoryListParamsDmarc string

const (
	RadarEmailSummaryThreatCategoryListParamsDmarcPass RadarEmailSummaryThreatCategoryListParamsDmarc = "PASS"
	RadarEmailSummaryThreatCategoryListParamsDmarcNone RadarEmailSummaryThreatCategoryListParamsDmarc = "NONE"
	RadarEmailSummaryThreatCategoryListParamsDmarcFail RadarEmailSummaryThreatCategoryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSummaryThreatCategoryListParamsFormat string

const (
	RadarEmailSummaryThreatCategoryListParamsFormatJson RadarEmailSummaryThreatCategoryListParamsFormat = "JSON"
	RadarEmailSummaryThreatCategoryListParamsFormatCsv  RadarEmailSummaryThreatCategoryListParamsFormat = "CSV"
)

type RadarEmailSummaryThreatCategoryListParamsSpf string

const (
	RadarEmailSummaryThreatCategoryListParamsSpfPass RadarEmailSummaryThreatCategoryListParamsSpf = "PASS"
	RadarEmailSummaryThreatCategoryListParamsSpfNone RadarEmailSummaryThreatCategoryListParamsSpf = "NONE"
	RadarEmailSummaryThreatCategoryListParamsSpfFail RadarEmailSummaryThreatCategoryListParamsSpf = "FAIL"
)
