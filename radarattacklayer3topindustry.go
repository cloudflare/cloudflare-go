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

// RadarAttackLayer3TopIndustryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopIndustryService] method instead.
type RadarAttackLayer3TopIndustryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopIndustryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopIndustryService(opts ...option.RequestOption) (r *RadarAttackLayer3TopIndustryService) {
	r = &RadarAttackLayer3TopIndustryService{}
	r.Options = opts
	return
}

// Get the Industry of attacks.
func (r *RadarAttackLayer3TopIndustryService) List(ctx context.Context, query RadarAttackLayer3TopIndustryListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopIndustryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopIndustryListResponseEnvelope
	path := "radar/attacks/layer3/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TopIndustryListResponse struct {
	Meta RadarAttackLayer3TopIndustryListResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopIndustryListResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopIndustryListResponseJSON   `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopIndustryListResponse]
type radarAttackLayer3TopIndustryListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryListResponseMeta struct {
	DateRange      []RadarAttackLayer3TopIndustryListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopIndustryListResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopIndustryListResponseMeta]
type radarAttackLayer3TopIndustryListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopIndustryListResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopIndustryListResponseMetaDateRange]
type radarAttackLayer3TopIndustryListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAttackLayer3TopIndustryListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfo]
type radarAttackLayer3TopIndustryListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopIndustryListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryListResponseTop0 struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  radarAttackLayer3TopIndustryListResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseTop0JSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopIndustryListResponseTop0]
type radarAttackLayer3TopIndustryListResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopIndustryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopIndustryListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopIndustryListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopIndustryListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopIndustryListParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TopIndustryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopIndustryListParamsDateRange string

const (
	RadarAttackLayer3TopIndustryListParamsDateRange1d         RadarAttackLayer3TopIndustryListParamsDateRange = "1d"
	RadarAttackLayer3TopIndustryListParamsDateRange2d         RadarAttackLayer3TopIndustryListParamsDateRange = "2d"
	RadarAttackLayer3TopIndustryListParamsDateRange7d         RadarAttackLayer3TopIndustryListParamsDateRange = "7d"
	RadarAttackLayer3TopIndustryListParamsDateRange14d        RadarAttackLayer3TopIndustryListParamsDateRange = "14d"
	RadarAttackLayer3TopIndustryListParamsDateRange28d        RadarAttackLayer3TopIndustryListParamsDateRange = "28d"
	RadarAttackLayer3TopIndustryListParamsDateRange12w        RadarAttackLayer3TopIndustryListParamsDateRange = "12w"
	RadarAttackLayer3TopIndustryListParamsDateRange24w        RadarAttackLayer3TopIndustryListParamsDateRange = "24w"
	RadarAttackLayer3TopIndustryListParamsDateRange52w        RadarAttackLayer3TopIndustryListParamsDateRange = "52w"
	RadarAttackLayer3TopIndustryListParamsDateRange1dControl  RadarAttackLayer3TopIndustryListParamsDateRange = "1dControl"
	RadarAttackLayer3TopIndustryListParamsDateRange2dControl  RadarAttackLayer3TopIndustryListParamsDateRange = "2dControl"
	RadarAttackLayer3TopIndustryListParamsDateRange7dControl  RadarAttackLayer3TopIndustryListParamsDateRange = "7dControl"
	RadarAttackLayer3TopIndustryListParamsDateRange14dControl RadarAttackLayer3TopIndustryListParamsDateRange = "14dControl"
	RadarAttackLayer3TopIndustryListParamsDateRange28dControl RadarAttackLayer3TopIndustryListParamsDateRange = "28dControl"
	RadarAttackLayer3TopIndustryListParamsDateRange12wControl RadarAttackLayer3TopIndustryListParamsDateRange = "12wControl"
	RadarAttackLayer3TopIndustryListParamsDateRange24wControl RadarAttackLayer3TopIndustryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopIndustryListParamsFormat string

const (
	RadarAttackLayer3TopIndustryListParamsFormatJson RadarAttackLayer3TopIndustryListParamsFormat = "JSON"
	RadarAttackLayer3TopIndustryListParamsFormatCsv  RadarAttackLayer3TopIndustryListParamsFormat = "CSV"
)

type RadarAttackLayer3TopIndustryListParamsIPVersion string

const (
	RadarAttackLayer3TopIndustryListParamsIPVersionIPv4 RadarAttackLayer3TopIndustryListParamsIPVersion = "IPv4"
	RadarAttackLayer3TopIndustryListParamsIPVersionIPv6 RadarAttackLayer3TopIndustryListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopIndustryListParamsProtocol string

const (
	RadarAttackLayer3TopIndustryListParamsProtocolUdp  RadarAttackLayer3TopIndustryListParamsProtocol = "UDP"
	RadarAttackLayer3TopIndustryListParamsProtocolTcp  RadarAttackLayer3TopIndustryListParamsProtocol = "TCP"
	RadarAttackLayer3TopIndustryListParamsProtocolIcmp RadarAttackLayer3TopIndustryListParamsProtocol = "ICMP"
	RadarAttackLayer3TopIndustryListParamsProtocolGre  RadarAttackLayer3TopIndustryListParamsProtocol = "GRE"
)

type RadarAttackLayer3TopIndustryListResponseEnvelope struct {
	Result  RadarAttackLayer3TopIndustryListResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer3TopIndustryListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopIndustryListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopIndustryListResponseEnvelope]
type radarAttackLayer3TopIndustryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
