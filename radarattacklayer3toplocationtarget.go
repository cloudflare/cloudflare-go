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

// RadarAttackLayer3TopLocationTargetService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopLocationTargetService] method instead.
type RadarAttackLayer3TopLocationTargetService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopLocationTargetService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TopLocationTargetService(opts ...option.RequestOption) (r *RadarAttackLayer3TopLocationTargetService) {
	r = &RadarAttackLayer3TopLocationTargetService{}
	r.Options = opts
	return
}

// Get the target locations of attacks.
func (r *RadarAttackLayer3TopLocationTargetService) List(ctx context.Context, query RadarAttackLayer3TopLocationTargetListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopLocationTargetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopLocationTargetListResponseEnvelope
	path := "radar/attacks/layer3/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TopLocationTargetListResponse struct {
	Meta RadarAttackLayer3TopLocationTargetListResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopLocationTargetListResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopLocationTargetListResponseJSON   `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopLocationTargetListResponse]
type radarAttackLayer3TopLocationTargetListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetListResponseMeta struct {
	DateRange      []RadarAttackLayer3TopLocationTargetListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopLocationTargetListResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationTargetListResponseMeta]
type radarAttackLayer3TopLocationTargetListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopLocationTargetListResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopLocationTargetListResponseMetaDateRange]
type radarAttackLayer3TopLocationTargetListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfo]
type radarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopLocationTargetListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetListResponseTop0 struct {
	Rank                float64                                                `json:"rank,required"`
	TargetCountryAlpha2 string                                                 `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                 `json:"targetCountryName,required"`
	Value               string                                                 `json:"value,required"`
	JSON                radarAttackLayer3TopLocationTargetListResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationTargetListResponseTop0]
type radarAttackLayer3TopLocationTargetListResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopLocationTargetListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopLocationTargetListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopLocationTargetListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopLocationTargetListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopLocationTargetListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TopLocationTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopLocationTargetListParamsDateRange string

const (
	RadarAttackLayer3TopLocationTargetListParamsDateRange1d         RadarAttackLayer3TopLocationTargetListParamsDateRange = "1d"
	RadarAttackLayer3TopLocationTargetListParamsDateRange2d         RadarAttackLayer3TopLocationTargetListParamsDateRange = "2d"
	RadarAttackLayer3TopLocationTargetListParamsDateRange7d         RadarAttackLayer3TopLocationTargetListParamsDateRange = "7d"
	RadarAttackLayer3TopLocationTargetListParamsDateRange14d        RadarAttackLayer3TopLocationTargetListParamsDateRange = "14d"
	RadarAttackLayer3TopLocationTargetListParamsDateRange28d        RadarAttackLayer3TopLocationTargetListParamsDateRange = "28d"
	RadarAttackLayer3TopLocationTargetListParamsDateRange12w        RadarAttackLayer3TopLocationTargetListParamsDateRange = "12w"
	RadarAttackLayer3TopLocationTargetListParamsDateRange24w        RadarAttackLayer3TopLocationTargetListParamsDateRange = "24w"
	RadarAttackLayer3TopLocationTargetListParamsDateRange52w        RadarAttackLayer3TopLocationTargetListParamsDateRange = "52w"
	RadarAttackLayer3TopLocationTargetListParamsDateRange1dControl  RadarAttackLayer3TopLocationTargetListParamsDateRange = "1dControl"
	RadarAttackLayer3TopLocationTargetListParamsDateRange2dControl  RadarAttackLayer3TopLocationTargetListParamsDateRange = "2dControl"
	RadarAttackLayer3TopLocationTargetListParamsDateRange7dControl  RadarAttackLayer3TopLocationTargetListParamsDateRange = "7dControl"
	RadarAttackLayer3TopLocationTargetListParamsDateRange14dControl RadarAttackLayer3TopLocationTargetListParamsDateRange = "14dControl"
	RadarAttackLayer3TopLocationTargetListParamsDateRange28dControl RadarAttackLayer3TopLocationTargetListParamsDateRange = "28dControl"
	RadarAttackLayer3TopLocationTargetListParamsDateRange12wControl RadarAttackLayer3TopLocationTargetListParamsDateRange = "12wControl"
	RadarAttackLayer3TopLocationTargetListParamsDateRange24wControl RadarAttackLayer3TopLocationTargetListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopLocationTargetListParamsFormat string

const (
	RadarAttackLayer3TopLocationTargetListParamsFormatJson RadarAttackLayer3TopLocationTargetListParamsFormat = "JSON"
	RadarAttackLayer3TopLocationTargetListParamsFormatCsv  RadarAttackLayer3TopLocationTargetListParamsFormat = "CSV"
)

type RadarAttackLayer3TopLocationTargetListParamsIPVersion string

const (
	RadarAttackLayer3TopLocationTargetListParamsIPVersionIPv4 RadarAttackLayer3TopLocationTargetListParamsIPVersion = "IPv4"
	RadarAttackLayer3TopLocationTargetListParamsIPVersionIPv6 RadarAttackLayer3TopLocationTargetListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopLocationTargetListParamsProtocol string

const (
	RadarAttackLayer3TopLocationTargetListParamsProtocolUdp  RadarAttackLayer3TopLocationTargetListParamsProtocol = "UDP"
	RadarAttackLayer3TopLocationTargetListParamsProtocolTcp  RadarAttackLayer3TopLocationTargetListParamsProtocol = "TCP"
	RadarAttackLayer3TopLocationTargetListParamsProtocolIcmp RadarAttackLayer3TopLocationTargetListParamsProtocol = "ICMP"
	RadarAttackLayer3TopLocationTargetListParamsProtocolGre  RadarAttackLayer3TopLocationTargetListParamsProtocol = "GRE"
)

type RadarAttackLayer3TopLocationTargetListResponseEnvelope struct {
	Result  RadarAttackLayer3TopLocationTargetListResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer3TopLocationTargetListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationTargetListResponseEnvelope]
type radarAttackLayer3TopLocationTargetListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
