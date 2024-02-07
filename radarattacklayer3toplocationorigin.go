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

// RadarAttackLayer3TopLocationOriginService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopLocationOriginService] method instead.
type RadarAttackLayer3TopLocationOriginService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopLocationOriginService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TopLocationOriginService(opts ...option.RequestOption) (r *RadarAttackLayer3TopLocationOriginService) {
	r = &RadarAttackLayer3TopLocationOriginService{}
	r.Options = opts
	return
}

// Get the origin locations of attacks.
func (r *RadarAttackLayer3TopLocationOriginService) List(ctx context.Context, query RadarAttackLayer3TopLocationOriginListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopLocationOriginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopLocationOriginListResponseEnvelope
	path := "radar/attacks/layer3/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TopLocationOriginListResponse struct {
	Meta RadarAttackLayer3TopLocationOriginListResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopLocationOriginListResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopLocationOriginListResponseJSON   `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopLocationOriginListResponse]
type radarAttackLayer3TopLocationOriginListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginListResponseMeta struct {
	DateRange      []RadarAttackLayer3TopLocationOriginListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopLocationOriginListResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationOriginListResponseMeta]
type radarAttackLayer3TopLocationOriginListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopLocationOriginListResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopLocationOriginListResponseMetaDateRange]
type radarAttackLayer3TopLocationOriginListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfo]
type radarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopLocationOriginListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginListResponseTop0 struct {
	OriginCountryAlpha2 string                                                 `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                 `json:"originCountryName,required"`
	Rank                float64                                                `json:"rank,required"`
	Value               string                                                 `json:"value,required"`
	JSON                radarAttackLayer3TopLocationOriginListResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationOriginListResponseTop0]
type radarAttackLayer3TopLocationOriginListResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopLocationOriginListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopLocationOriginListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopLocationOriginListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopLocationOriginListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopLocationOriginListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TopLocationOriginListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopLocationOriginListParamsDateRange string

const (
	RadarAttackLayer3TopLocationOriginListParamsDateRange1d         RadarAttackLayer3TopLocationOriginListParamsDateRange = "1d"
	RadarAttackLayer3TopLocationOriginListParamsDateRange2d         RadarAttackLayer3TopLocationOriginListParamsDateRange = "2d"
	RadarAttackLayer3TopLocationOriginListParamsDateRange7d         RadarAttackLayer3TopLocationOriginListParamsDateRange = "7d"
	RadarAttackLayer3TopLocationOriginListParamsDateRange14d        RadarAttackLayer3TopLocationOriginListParamsDateRange = "14d"
	RadarAttackLayer3TopLocationOriginListParamsDateRange28d        RadarAttackLayer3TopLocationOriginListParamsDateRange = "28d"
	RadarAttackLayer3TopLocationOriginListParamsDateRange12w        RadarAttackLayer3TopLocationOriginListParamsDateRange = "12w"
	RadarAttackLayer3TopLocationOriginListParamsDateRange24w        RadarAttackLayer3TopLocationOriginListParamsDateRange = "24w"
	RadarAttackLayer3TopLocationOriginListParamsDateRange52w        RadarAttackLayer3TopLocationOriginListParamsDateRange = "52w"
	RadarAttackLayer3TopLocationOriginListParamsDateRange1dControl  RadarAttackLayer3TopLocationOriginListParamsDateRange = "1dControl"
	RadarAttackLayer3TopLocationOriginListParamsDateRange2dControl  RadarAttackLayer3TopLocationOriginListParamsDateRange = "2dControl"
	RadarAttackLayer3TopLocationOriginListParamsDateRange7dControl  RadarAttackLayer3TopLocationOriginListParamsDateRange = "7dControl"
	RadarAttackLayer3TopLocationOriginListParamsDateRange14dControl RadarAttackLayer3TopLocationOriginListParamsDateRange = "14dControl"
	RadarAttackLayer3TopLocationOriginListParamsDateRange28dControl RadarAttackLayer3TopLocationOriginListParamsDateRange = "28dControl"
	RadarAttackLayer3TopLocationOriginListParamsDateRange12wControl RadarAttackLayer3TopLocationOriginListParamsDateRange = "12wControl"
	RadarAttackLayer3TopLocationOriginListParamsDateRange24wControl RadarAttackLayer3TopLocationOriginListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopLocationOriginListParamsFormat string

const (
	RadarAttackLayer3TopLocationOriginListParamsFormatJson RadarAttackLayer3TopLocationOriginListParamsFormat = "JSON"
	RadarAttackLayer3TopLocationOriginListParamsFormatCsv  RadarAttackLayer3TopLocationOriginListParamsFormat = "CSV"
)

type RadarAttackLayer3TopLocationOriginListParamsIPVersion string

const (
	RadarAttackLayer3TopLocationOriginListParamsIPVersionIPv4 RadarAttackLayer3TopLocationOriginListParamsIPVersion = "IPv4"
	RadarAttackLayer3TopLocationOriginListParamsIPVersionIPv6 RadarAttackLayer3TopLocationOriginListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopLocationOriginListParamsProtocol string

const (
	RadarAttackLayer3TopLocationOriginListParamsProtocolUdp  RadarAttackLayer3TopLocationOriginListParamsProtocol = "UDP"
	RadarAttackLayer3TopLocationOriginListParamsProtocolTcp  RadarAttackLayer3TopLocationOriginListParamsProtocol = "TCP"
	RadarAttackLayer3TopLocationOriginListParamsProtocolIcmp RadarAttackLayer3TopLocationOriginListParamsProtocol = "ICMP"
	RadarAttackLayer3TopLocationOriginListParamsProtocolGre  RadarAttackLayer3TopLocationOriginListParamsProtocol = "GRE"
)

type RadarAttackLayer3TopLocationOriginListResponseEnvelope struct {
	Result  RadarAttackLayer3TopLocationOriginListResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer3TopLocationOriginListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationOriginListResponseEnvelope]
type radarAttackLayer3TopLocationOriginListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
