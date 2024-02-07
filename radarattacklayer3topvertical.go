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

// RadarAttackLayer3TopVerticalService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopVerticalService] method instead.
type RadarAttackLayer3TopVerticalService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopVerticalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopVerticalService(opts ...option.RequestOption) (r *RadarAttackLayer3TopVerticalService) {
	r = &RadarAttackLayer3TopVerticalService{}
	r.Options = opts
	return
}

// Get the Verticals of attacks.
func (r *RadarAttackLayer3TopVerticalService) List(ctx context.Context, query RadarAttackLayer3TopVerticalListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopVerticalListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopVerticalListResponseEnvelope
	path := "radar/attacks/layer3/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TopVerticalListResponse struct {
	Meta RadarAttackLayer3TopVerticalListResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopVerticalListResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopVerticalListResponseJSON   `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopVerticalListResponse]
type radarAttackLayer3TopVerticalListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalListResponseMeta struct {
	DateRange      []RadarAttackLayer3TopVerticalListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopVerticalListResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopVerticalListResponseMeta]
type radarAttackLayer3TopVerticalListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopVerticalListResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopVerticalListResponseMetaDateRange]
type radarAttackLayer3TopVerticalListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAttackLayer3TopVerticalListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfo]
type radarAttackLayer3TopVerticalListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopVerticalListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalListResponseTop0 struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  radarAttackLayer3TopVerticalListResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseTop0JSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopVerticalListResponseTop0]
type radarAttackLayer3TopVerticalListResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopVerticalListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopVerticalListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopVerticalListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopVerticalListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopVerticalListParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TopVerticalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopVerticalListParamsDateRange string

const (
	RadarAttackLayer3TopVerticalListParamsDateRange1d         RadarAttackLayer3TopVerticalListParamsDateRange = "1d"
	RadarAttackLayer3TopVerticalListParamsDateRange2d         RadarAttackLayer3TopVerticalListParamsDateRange = "2d"
	RadarAttackLayer3TopVerticalListParamsDateRange7d         RadarAttackLayer3TopVerticalListParamsDateRange = "7d"
	RadarAttackLayer3TopVerticalListParamsDateRange14d        RadarAttackLayer3TopVerticalListParamsDateRange = "14d"
	RadarAttackLayer3TopVerticalListParamsDateRange28d        RadarAttackLayer3TopVerticalListParamsDateRange = "28d"
	RadarAttackLayer3TopVerticalListParamsDateRange12w        RadarAttackLayer3TopVerticalListParamsDateRange = "12w"
	RadarAttackLayer3TopVerticalListParamsDateRange24w        RadarAttackLayer3TopVerticalListParamsDateRange = "24w"
	RadarAttackLayer3TopVerticalListParamsDateRange52w        RadarAttackLayer3TopVerticalListParamsDateRange = "52w"
	RadarAttackLayer3TopVerticalListParamsDateRange1dControl  RadarAttackLayer3TopVerticalListParamsDateRange = "1dControl"
	RadarAttackLayer3TopVerticalListParamsDateRange2dControl  RadarAttackLayer3TopVerticalListParamsDateRange = "2dControl"
	RadarAttackLayer3TopVerticalListParamsDateRange7dControl  RadarAttackLayer3TopVerticalListParamsDateRange = "7dControl"
	RadarAttackLayer3TopVerticalListParamsDateRange14dControl RadarAttackLayer3TopVerticalListParamsDateRange = "14dControl"
	RadarAttackLayer3TopVerticalListParamsDateRange28dControl RadarAttackLayer3TopVerticalListParamsDateRange = "28dControl"
	RadarAttackLayer3TopVerticalListParamsDateRange12wControl RadarAttackLayer3TopVerticalListParamsDateRange = "12wControl"
	RadarAttackLayer3TopVerticalListParamsDateRange24wControl RadarAttackLayer3TopVerticalListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopVerticalListParamsFormat string

const (
	RadarAttackLayer3TopVerticalListParamsFormatJson RadarAttackLayer3TopVerticalListParamsFormat = "JSON"
	RadarAttackLayer3TopVerticalListParamsFormatCsv  RadarAttackLayer3TopVerticalListParamsFormat = "CSV"
)

type RadarAttackLayer3TopVerticalListParamsIPVersion string

const (
	RadarAttackLayer3TopVerticalListParamsIPVersionIPv4 RadarAttackLayer3TopVerticalListParamsIPVersion = "IPv4"
	RadarAttackLayer3TopVerticalListParamsIPVersionIPv6 RadarAttackLayer3TopVerticalListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopVerticalListParamsProtocol string

const (
	RadarAttackLayer3TopVerticalListParamsProtocolUdp  RadarAttackLayer3TopVerticalListParamsProtocol = "UDP"
	RadarAttackLayer3TopVerticalListParamsProtocolTcp  RadarAttackLayer3TopVerticalListParamsProtocol = "TCP"
	RadarAttackLayer3TopVerticalListParamsProtocolIcmp RadarAttackLayer3TopVerticalListParamsProtocol = "ICMP"
	RadarAttackLayer3TopVerticalListParamsProtocolGre  RadarAttackLayer3TopVerticalListParamsProtocol = "GRE"
)

type RadarAttackLayer3TopVerticalListResponseEnvelope struct {
	Result  RadarAttackLayer3TopVerticalListResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer3TopVerticalListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopVerticalListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopVerticalListResponseEnvelope]
type radarAttackLayer3TopVerticalListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
