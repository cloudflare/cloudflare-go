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

// RadarAttackLayer7TopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopAseService] method instead.
type RadarAttackLayer7TopAseService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopAseService(opts ...option.RequestOption) (r *RadarAttackLayer7TopAseService) {
	r = &RadarAttackLayer7TopAseService{}
	r.Options = opts
	return
}

// Get the top origin Autonomous Systems of and by layer 7 attacks. Values are a
// percentage out of the total layer 7 attacks. The origin Autonomous Systems is
// determined by the client IP.
func (r *RadarAttackLayer7TopAseService) Origin(ctx context.Context, query RadarAttackLayer7TopAseOriginParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopAseOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TopAseOriginResponseEnvelope
	path := "radar/attacks/layer7/top/ases/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7TopAseOriginResponse struct {
	Meta RadarAttackLayer7TopAseOriginResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopAseOriginResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopAseOriginResponseJSON   `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAseOriginResponse]
type radarAttackLayer7TopAseOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginResponseMeta struct {
	DateRange      []RadarAttackLayer7TopAseOriginResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopAseOriginResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAseOriginResponseMeta]
type radarAttackLayer7TopAseOriginResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopAseOriginResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopAseOriginResponseMetaDateRange]
type radarAttackLayer7TopAseOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarAttackLayer7TopAseOriginResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfo]
type radarAttackLayer7TopAseOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginResponseTop0 struct {
	OriginASN     string                                        `json:"originAsn,required"`
	OriginASNName string                                        `json:"originAsnName,required"`
	Rank          float64                                       `json:"rank,required"`
	Value         string                                        `json:"value,required"`
	JSON          radarAttackLayer7TopAseOriginResponseTop0JSON `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAseOriginResponseTop0]
type radarAttackLayer7TopAseOriginResponseTop0JSON struct {
	OriginASN     apijson.Field
	OriginASNName apijson.Field
	Rank          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopAseOriginParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopAseOriginParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopAseOriginParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TopAseOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopAseOriginParamsDateRange string

const (
	RadarAttackLayer7TopAseOriginParamsDateRange1d         RadarAttackLayer7TopAseOriginParamsDateRange = "1d"
	RadarAttackLayer7TopAseOriginParamsDateRange2d         RadarAttackLayer7TopAseOriginParamsDateRange = "2d"
	RadarAttackLayer7TopAseOriginParamsDateRange7d         RadarAttackLayer7TopAseOriginParamsDateRange = "7d"
	RadarAttackLayer7TopAseOriginParamsDateRange14d        RadarAttackLayer7TopAseOriginParamsDateRange = "14d"
	RadarAttackLayer7TopAseOriginParamsDateRange28d        RadarAttackLayer7TopAseOriginParamsDateRange = "28d"
	RadarAttackLayer7TopAseOriginParamsDateRange12w        RadarAttackLayer7TopAseOriginParamsDateRange = "12w"
	RadarAttackLayer7TopAseOriginParamsDateRange24w        RadarAttackLayer7TopAseOriginParamsDateRange = "24w"
	RadarAttackLayer7TopAseOriginParamsDateRange52w        RadarAttackLayer7TopAseOriginParamsDateRange = "52w"
	RadarAttackLayer7TopAseOriginParamsDateRange1dControl  RadarAttackLayer7TopAseOriginParamsDateRange = "1dControl"
	RadarAttackLayer7TopAseOriginParamsDateRange2dControl  RadarAttackLayer7TopAseOriginParamsDateRange = "2dControl"
	RadarAttackLayer7TopAseOriginParamsDateRange7dControl  RadarAttackLayer7TopAseOriginParamsDateRange = "7dControl"
	RadarAttackLayer7TopAseOriginParamsDateRange14dControl RadarAttackLayer7TopAseOriginParamsDateRange = "14dControl"
	RadarAttackLayer7TopAseOriginParamsDateRange28dControl RadarAttackLayer7TopAseOriginParamsDateRange = "28dControl"
	RadarAttackLayer7TopAseOriginParamsDateRange12wControl RadarAttackLayer7TopAseOriginParamsDateRange = "12wControl"
	RadarAttackLayer7TopAseOriginParamsDateRange24wControl RadarAttackLayer7TopAseOriginParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopAseOriginParamsFormat string

const (
	RadarAttackLayer7TopAseOriginParamsFormatJson RadarAttackLayer7TopAseOriginParamsFormat = "JSON"
	RadarAttackLayer7TopAseOriginParamsFormatCsv  RadarAttackLayer7TopAseOriginParamsFormat = "CSV"
)

type RadarAttackLayer7TopAseOriginResponseEnvelope struct {
	Result  RadarAttackLayer7TopAseOriginResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAttackLayer7TopAseOriginResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TopAseOriginResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopAseOriginResponseEnvelope]
type radarAttackLayer7TopAseOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
