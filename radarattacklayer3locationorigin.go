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

// RadarAttackLayer3LocationOriginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3LocationOriginService] method instead.
type RadarAttackLayer3LocationOriginService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3LocationOriginService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3LocationOriginService(opts ...option.RequestOption) (r *RadarAttackLayer3LocationOriginService) {
	r = &RadarAttackLayer3LocationOriginService{}
	r.Options = opts
	return
}

// Get the origin locations of attacks.
func (r *RadarAttackLayer3LocationOriginService) List(ctx context.Context, query RadarAttackLayer3LocationOriginListParams, opts ...option.RequestOption) (res *RadarAttackLayer3LocationOriginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3LocationOriginListResponse struct {
	Result  RadarAttackLayer3LocationOriginListResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAttackLayer3LocationOriginListResponseJSON   `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3LocationOriginListResponse]
type radarAttackLayer3LocationOriginListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationOriginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListResponseResult struct {
	Meta RadarAttackLayer3LocationOriginListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3LocationOriginListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3LocationOriginListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3LocationOriginListResponseResult]
type radarAttackLayer3LocationOriginListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationOriginListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3LocationOriginListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3LocationOriginListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3LocationOriginListResponseResultMeta]
type radarAttackLayer3LocationOriginListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationOriginListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3LocationOriginListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3LocationOriginListResponseResultMetaDateRange]
type radarAttackLayer3LocationOriginListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationOriginListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfo]
type radarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3LocationOriginListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListResponseResultTop0 struct {
	OriginCountryAlpha2 string                                                    `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                    `json:"originCountryName,required"`
	Rank                float64                                                   `json:"rank,required"`
	Value               string                                                    `json:"value,required"`
	JSON                radarAttackLayer3LocationOriginListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3LocationOriginListResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3LocationOriginListResponseResultTop0]
type radarAttackLayer3LocationOriginListResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationOriginListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationOriginListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3LocationOriginListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3LocationOriginListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3LocationOriginListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3LocationOriginListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3LocationOriginListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3LocationOriginListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3LocationOriginListParamsDateRange string

const (
	RadarAttackLayer3LocationOriginListParamsDateRange1d         RadarAttackLayer3LocationOriginListParamsDateRange = "1d"
	RadarAttackLayer3LocationOriginListParamsDateRange2d         RadarAttackLayer3LocationOriginListParamsDateRange = "2d"
	RadarAttackLayer3LocationOriginListParamsDateRange7d         RadarAttackLayer3LocationOriginListParamsDateRange = "7d"
	RadarAttackLayer3LocationOriginListParamsDateRange14d        RadarAttackLayer3LocationOriginListParamsDateRange = "14d"
	RadarAttackLayer3LocationOriginListParamsDateRange28d        RadarAttackLayer3LocationOriginListParamsDateRange = "28d"
	RadarAttackLayer3LocationOriginListParamsDateRange12w        RadarAttackLayer3LocationOriginListParamsDateRange = "12w"
	RadarAttackLayer3LocationOriginListParamsDateRange24w        RadarAttackLayer3LocationOriginListParamsDateRange = "24w"
	RadarAttackLayer3LocationOriginListParamsDateRange52w        RadarAttackLayer3LocationOriginListParamsDateRange = "52w"
	RadarAttackLayer3LocationOriginListParamsDateRange1dControl  RadarAttackLayer3LocationOriginListParamsDateRange = "1dControl"
	RadarAttackLayer3LocationOriginListParamsDateRange2dControl  RadarAttackLayer3LocationOriginListParamsDateRange = "2dControl"
	RadarAttackLayer3LocationOriginListParamsDateRange7dControl  RadarAttackLayer3LocationOriginListParamsDateRange = "7dControl"
	RadarAttackLayer3LocationOriginListParamsDateRange14dControl RadarAttackLayer3LocationOriginListParamsDateRange = "14dControl"
	RadarAttackLayer3LocationOriginListParamsDateRange28dControl RadarAttackLayer3LocationOriginListParamsDateRange = "28dControl"
	RadarAttackLayer3LocationOriginListParamsDateRange12wControl RadarAttackLayer3LocationOriginListParamsDateRange = "12wControl"
	RadarAttackLayer3LocationOriginListParamsDateRange24wControl RadarAttackLayer3LocationOriginListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3LocationOriginListParamsFormat string

const (
	RadarAttackLayer3LocationOriginListParamsFormatJson RadarAttackLayer3LocationOriginListParamsFormat = "JSON"
	RadarAttackLayer3LocationOriginListParamsFormatCsv  RadarAttackLayer3LocationOriginListParamsFormat = "CSV"
)

type RadarAttackLayer3LocationOriginListParamsIPVersion string

const (
	RadarAttackLayer3LocationOriginListParamsIPVersionIPv4 RadarAttackLayer3LocationOriginListParamsIPVersion = "IPv4"
	RadarAttackLayer3LocationOriginListParamsIPVersionIPv6 RadarAttackLayer3LocationOriginListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3LocationOriginListParamsProtocol string

const (
	RadarAttackLayer3LocationOriginListParamsProtocolUdp  RadarAttackLayer3LocationOriginListParamsProtocol = "UDP"
	RadarAttackLayer3LocationOriginListParamsProtocolTcp  RadarAttackLayer3LocationOriginListParamsProtocol = "TCP"
	RadarAttackLayer3LocationOriginListParamsProtocolIcmp RadarAttackLayer3LocationOriginListParamsProtocol = "ICMP"
	RadarAttackLayer3LocationOriginListParamsProtocolGre  RadarAttackLayer3LocationOriginListParamsProtocol = "GRE"
)
