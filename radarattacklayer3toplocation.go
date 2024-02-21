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

// RadarAttackLayer3TopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopLocationService] method instead.
type RadarAttackLayer3TopLocationService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopLocationService(opts ...option.RequestOption) (r *RadarAttackLayer3TopLocationService) {
	r = &RadarAttackLayer3TopLocationService{}
	r.Options = opts
	return
}

// Get the origin locations of attacks.
func (r *RadarAttackLayer3TopLocationService) Origin(ctx context.Context, query RadarAttackLayer3TopLocationOriginParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopLocationOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopLocationOriginResponseEnvelope
	path := "radar/attacks/layer3/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the target locations of attacks.
func (r *RadarAttackLayer3TopLocationService) Target(ctx context.Context, query RadarAttackLayer3TopLocationTargetParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopLocationTargetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopLocationTargetResponseEnvelope
	path := "radar/attacks/layer3/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TopLocationOriginResponse struct {
	Meta RadarAttackLayer3TopLocationOriginResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopLocationOriginResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopLocationOriginResponseJSON   `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopLocationOriginResponse]
type radarAttackLayer3TopLocationOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginResponseMeta struct {
	DateRange      []RadarAttackLayer3TopLocationOriginResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopLocationOriginResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopLocationOriginResponseMeta]
type radarAttackLayer3TopLocationOriginResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopLocationOriginResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopLocationOriginResponseMetaDateRange]
type radarAttackLayer3TopLocationOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfo]
type radarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginResponseTop0 struct {
	OriginCountryAlpha2 string                                             `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                             `json:"originCountryName,required"`
	Rank                float64                                            `json:"rank,required"`
	Value               string                                             `json:"value,required"`
	JSON                radarAttackLayer3TopLocationOriginResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopLocationOriginResponseTop0]
type radarAttackLayer3TopLocationOriginResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetResponse struct {
	Meta RadarAttackLayer3TopLocationTargetResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopLocationTargetResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopLocationTargetResponseJSON   `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopLocationTargetResponse]
type radarAttackLayer3TopLocationTargetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetResponseMeta struct {
	DateRange      []RadarAttackLayer3TopLocationTargetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopLocationTargetResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopLocationTargetResponseMeta]
type radarAttackLayer3TopLocationTargetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopLocationTargetResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopLocationTargetResponseMetaDateRange]
type radarAttackLayer3TopLocationTargetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfo]
type radarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetResponseTop0 struct {
	Rank                float64                                            `json:"rank,required"`
	TargetCountryAlpha2 string                                             `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                             `json:"targetCountryName,required"`
	Value               string                                             `json:"value,required"`
	JSON                radarAttackLayer3TopLocationTargetResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopLocationTargetResponseTop0]
type radarAttackLayer3TopLocationTargetResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationOriginParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopLocationOriginParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopLocationOriginParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopLocationOriginParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopLocationOriginParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopLocationOriginParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TopLocationOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopLocationOriginParamsDateRange string

const (
	RadarAttackLayer3TopLocationOriginParamsDateRange1d         RadarAttackLayer3TopLocationOriginParamsDateRange = "1d"
	RadarAttackLayer3TopLocationOriginParamsDateRange2d         RadarAttackLayer3TopLocationOriginParamsDateRange = "2d"
	RadarAttackLayer3TopLocationOriginParamsDateRange7d         RadarAttackLayer3TopLocationOriginParamsDateRange = "7d"
	RadarAttackLayer3TopLocationOriginParamsDateRange14d        RadarAttackLayer3TopLocationOriginParamsDateRange = "14d"
	RadarAttackLayer3TopLocationOriginParamsDateRange28d        RadarAttackLayer3TopLocationOriginParamsDateRange = "28d"
	RadarAttackLayer3TopLocationOriginParamsDateRange12w        RadarAttackLayer3TopLocationOriginParamsDateRange = "12w"
	RadarAttackLayer3TopLocationOriginParamsDateRange24w        RadarAttackLayer3TopLocationOriginParamsDateRange = "24w"
	RadarAttackLayer3TopLocationOriginParamsDateRange52w        RadarAttackLayer3TopLocationOriginParamsDateRange = "52w"
	RadarAttackLayer3TopLocationOriginParamsDateRange1dControl  RadarAttackLayer3TopLocationOriginParamsDateRange = "1dControl"
	RadarAttackLayer3TopLocationOriginParamsDateRange2dControl  RadarAttackLayer3TopLocationOriginParamsDateRange = "2dControl"
	RadarAttackLayer3TopLocationOriginParamsDateRange7dControl  RadarAttackLayer3TopLocationOriginParamsDateRange = "7dControl"
	RadarAttackLayer3TopLocationOriginParamsDateRange14dControl RadarAttackLayer3TopLocationOriginParamsDateRange = "14dControl"
	RadarAttackLayer3TopLocationOriginParamsDateRange28dControl RadarAttackLayer3TopLocationOriginParamsDateRange = "28dControl"
	RadarAttackLayer3TopLocationOriginParamsDateRange12wControl RadarAttackLayer3TopLocationOriginParamsDateRange = "12wControl"
	RadarAttackLayer3TopLocationOriginParamsDateRange24wControl RadarAttackLayer3TopLocationOriginParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopLocationOriginParamsFormat string

const (
	RadarAttackLayer3TopLocationOriginParamsFormatJson RadarAttackLayer3TopLocationOriginParamsFormat = "JSON"
	RadarAttackLayer3TopLocationOriginParamsFormatCsv  RadarAttackLayer3TopLocationOriginParamsFormat = "CSV"
)

type RadarAttackLayer3TopLocationOriginParamsIPVersion string

const (
	RadarAttackLayer3TopLocationOriginParamsIPVersionIPv4 RadarAttackLayer3TopLocationOriginParamsIPVersion = "IPv4"
	RadarAttackLayer3TopLocationOriginParamsIPVersionIPv6 RadarAttackLayer3TopLocationOriginParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopLocationOriginParamsProtocol string

const (
	RadarAttackLayer3TopLocationOriginParamsProtocolUdp  RadarAttackLayer3TopLocationOriginParamsProtocol = "UDP"
	RadarAttackLayer3TopLocationOriginParamsProtocolTcp  RadarAttackLayer3TopLocationOriginParamsProtocol = "TCP"
	RadarAttackLayer3TopLocationOriginParamsProtocolIcmp RadarAttackLayer3TopLocationOriginParamsProtocol = "ICMP"
	RadarAttackLayer3TopLocationOriginParamsProtocolGre  RadarAttackLayer3TopLocationOriginParamsProtocol = "GRE"
)

type RadarAttackLayer3TopLocationOriginResponseEnvelope struct {
	Result  RadarAttackLayer3TopLocationOriginResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAttackLayer3TopLocationOriginResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopLocationOriginResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationOriginResponseEnvelope]
type radarAttackLayer3TopLocationOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopLocationTargetParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopLocationTargetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopLocationTargetParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopLocationTargetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopLocationTargetParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopLocationTargetParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TopLocationTargetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopLocationTargetParamsDateRange string

const (
	RadarAttackLayer3TopLocationTargetParamsDateRange1d         RadarAttackLayer3TopLocationTargetParamsDateRange = "1d"
	RadarAttackLayer3TopLocationTargetParamsDateRange2d         RadarAttackLayer3TopLocationTargetParamsDateRange = "2d"
	RadarAttackLayer3TopLocationTargetParamsDateRange7d         RadarAttackLayer3TopLocationTargetParamsDateRange = "7d"
	RadarAttackLayer3TopLocationTargetParamsDateRange14d        RadarAttackLayer3TopLocationTargetParamsDateRange = "14d"
	RadarAttackLayer3TopLocationTargetParamsDateRange28d        RadarAttackLayer3TopLocationTargetParamsDateRange = "28d"
	RadarAttackLayer3TopLocationTargetParamsDateRange12w        RadarAttackLayer3TopLocationTargetParamsDateRange = "12w"
	RadarAttackLayer3TopLocationTargetParamsDateRange24w        RadarAttackLayer3TopLocationTargetParamsDateRange = "24w"
	RadarAttackLayer3TopLocationTargetParamsDateRange52w        RadarAttackLayer3TopLocationTargetParamsDateRange = "52w"
	RadarAttackLayer3TopLocationTargetParamsDateRange1dControl  RadarAttackLayer3TopLocationTargetParamsDateRange = "1dControl"
	RadarAttackLayer3TopLocationTargetParamsDateRange2dControl  RadarAttackLayer3TopLocationTargetParamsDateRange = "2dControl"
	RadarAttackLayer3TopLocationTargetParamsDateRange7dControl  RadarAttackLayer3TopLocationTargetParamsDateRange = "7dControl"
	RadarAttackLayer3TopLocationTargetParamsDateRange14dControl RadarAttackLayer3TopLocationTargetParamsDateRange = "14dControl"
	RadarAttackLayer3TopLocationTargetParamsDateRange28dControl RadarAttackLayer3TopLocationTargetParamsDateRange = "28dControl"
	RadarAttackLayer3TopLocationTargetParamsDateRange12wControl RadarAttackLayer3TopLocationTargetParamsDateRange = "12wControl"
	RadarAttackLayer3TopLocationTargetParamsDateRange24wControl RadarAttackLayer3TopLocationTargetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopLocationTargetParamsFormat string

const (
	RadarAttackLayer3TopLocationTargetParamsFormatJson RadarAttackLayer3TopLocationTargetParamsFormat = "JSON"
	RadarAttackLayer3TopLocationTargetParamsFormatCsv  RadarAttackLayer3TopLocationTargetParamsFormat = "CSV"
)

type RadarAttackLayer3TopLocationTargetParamsIPVersion string

const (
	RadarAttackLayer3TopLocationTargetParamsIPVersionIPv4 RadarAttackLayer3TopLocationTargetParamsIPVersion = "IPv4"
	RadarAttackLayer3TopLocationTargetParamsIPVersionIPv6 RadarAttackLayer3TopLocationTargetParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopLocationTargetParamsProtocol string

const (
	RadarAttackLayer3TopLocationTargetParamsProtocolUdp  RadarAttackLayer3TopLocationTargetParamsProtocol = "UDP"
	RadarAttackLayer3TopLocationTargetParamsProtocolTcp  RadarAttackLayer3TopLocationTargetParamsProtocol = "TCP"
	RadarAttackLayer3TopLocationTargetParamsProtocolIcmp RadarAttackLayer3TopLocationTargetParamsProtocol = "ICMP"
	RadarAttackLayer3TopLocationTargetParamsProtocolGre  RadarAttackLayer3TopLocationTargetParamsProtocol = "GRE"
)

type RadarAttackLayer3TopLocationTargetResponseEnvelope struct {
	Result  RadarAttackLayer3TopLocationTargetResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAttackLayer3TopLocationTargetResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopLocationTargetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopLocationTargetResponseEnvelope]
type radarAttackLayer3TopLocationTargetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
