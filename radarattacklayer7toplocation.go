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

// RadarAttackLayer7TopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopLocationService] method instead.
type RadarAttackLayer7TopLocationService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopLocationService(opts ...option.RequestOption) (r *RadarAttackLayer7TopLocationService) {
	r = &RadarAttackLayer7TopLocationService{}
	r.Options = opts
	return
}

// Get the top origin locations of and by layer 7 attacks. Values are a percentage
// out of the total layer 7 attacks. The origin location is determined by the
// client IP.
func (r *RadarAttackLayer7TopLocationService) Origin(ctx context.Context, query RadarAttackLayer7TopLocationOriginParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TopLocationOriginResponseEnvelope
	path := "radar/attacks/layer7/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top target locations of and by layer 7 attacks. Values are a percentage
// out of the total layer 7 attacks. The target location is determined by the
// attacked zone's billing country, when available.
func (r *RadarAttackLayer7TopLocationService) Target(ctx context.Context, query RadarAttackLayer7TopLocationTargetParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationTargetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TopLocationTargetResponseEnvelope
	path := "radar/attacks/layer7/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7TopLocationOriginResponse struct {
	Meta RadarAttackLayer7TopLocationOriginResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopLocationOriginResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationOriginResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopLocationOriginResponse]
type radarAttackLayer7TopLocationOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginResponseMeta struct {
	DateRange      []RadarAttackLayer7TopLocationOriginResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopLocationOriginResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationOriginResponseMeta]
type radarAttackLayer7TopLocationOriginResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopLocationOriginResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationOriginResponseMetaDateRange]
type radarAttackLayer7TopLocationOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfo]
type radarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginResponseTop0 struct {
	OriginCountryAlpha2 string                                             `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                             `json:"originCountryName,required"`
	Rank                float64                                            `json:"rank,required"`
	Value               string                                             `json:"value,required"`
	JSON                radarAttackLayer7TopLocationOriginResponseTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationOriginResponseTop0]
type radarAttackLayer7TopLocationOriginResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetResponse struct {
	Meta RadarAttackLayer7TopLocationTargetResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopLocationTargetResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationTargetResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopLocationTargetResponse]
type radarAttackLayer7TopLocationTargetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetResponseMeta struct {
	DateRange      []RadarAttackLayer7TopLocationTargetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopLocationTargetResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationTargetResponseMeta]
type radarAttackLayer7TopLocationTargetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopLocationTargetResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationTargetResponseMetaDateRange]
type radarAttackLayer7TopLocationTargetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfo]
type radarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetResponseTop0 struct {
	Rank                float64                                            `json:"rank,required"`
	TargetCountryAlpha2 string                                             `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                             `json:"targetCountryName,required"`
	Value               string                                             `json:"value,required"`
	JSON                radarAttackLayer7TopLocationTargetResponseTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationTargetResponseTop0]
type radarAttackLayer7TopLocationTargetResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopLocationOriginParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopLocationOriginParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationOriginParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopLocationOriginParamsDateRange string

const (
	RadarAttackLayer7TopLocationOriginParamsDateRange1d         RadarAttackLayer7TopLocationOriginParamsDateRange = "1d"
	RadarAttackLayer7TopLocationOriginParamsDateRange2d         RadarAttackLayer7TopLocationOriginParamsDateRange = "2d"
	RadarAttackLayer7TopLocationOriginParamsDateRange7d         RadarAttackLayer7TopLocationOriginParamsDateRange = "7d"
	RadarAttackLayer7TopLocationOriginParamsDateRange14d        RadarAttackLayer7TopLocationOriginParamsDateRange = "14d"
	RadarAttackLayer7TopLocationOriginParamsDateRange28d        RadarAttackLayer7TopLocationOriginParamsDateRange = "28d"
	RadarAttackLayer7TopLocationOriginParamsDateRange12w        RadarAttackLayer7TopLocationOriginParamsDateRange = "12w"
	RadarAttackLayer7TopLocationOriginParamsDateRange24w        RadarAttackLayer7TopLocationOriginParamsDateRange = "24w"
	RadarAttackLayer7TopLocationOriginParamsDateRange52w        RadarAttackLayer7TopLocationOriginParamsDateRange = "52w"
	RadarAttackLayer7TopLocationOriginParamsDateRange1dControl  RadarAttackLayer7TopLocationOriginParamsDateRange = "1dControl"
	RadarAttackLayer7TopLocationOriginParamsDateRange2dControl  RadarAttackLayer7TopLocationOriginParamsDateRange = "2dControl"
	RadarAttackLayer7TopLocationOriginParamsDateRange7dControl  RadarAttackLayer7TopLocationOriginParamsDateRange = "7dControl"
	RadarAttackLayer7TopLocationOriginParamsDateRange14dControl RadarAttackLayer7TopLocationOriginParamsDateRange = "14dControl"
	RadarAttackLayer7TopLocationOriginParamsDateRange28dControl RadarAttackLayer7TopLocationOriginParamsDateRange = "28dControl"
	RadarAttackLayer7TopLocationOriginParamsDateRange12wControl RadarAttackLayer7TopLocationOriginParamsDateRange = "12wControl"
	RadarAttackLayer7TopLocationOriginParamsDateRange24wControl RadarAttackLayer7TopLocationOriginParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopLocationOriginParamsFormat string

const (
	RadarAttackLayer7TopLocationOriginParamsFormatJson RadarAttackLayer7TopLocationOriginParamsFormat = "JSON"
	RadarAttackLayer7TopLocationOriginParamsFormatCsv  RadarAttackLayer7TopLocationOriginParamsFormat = "CSV"
)

type RadarAttackLayer7TopLocationOriginResponseEnvelope struct {
	Result  RadarAttackLayer7TopLocationOriginResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAttackLayer7TopLocationOriginResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopLocationOriginResponseEnvelope]
type radarAttackLayer7TopLocationOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopLocationTargetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopLocationTargetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationTargetParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationTargetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopLocationTargetParamsDateRange string

const (
	RadarAttackLayer7TopLocationTargetParamsDateRange1d         RadarAttackLayer7TopLocationTargetParamsDateRange = "1d"
	RadarAttackLayer7TopLocationTargetParamsDateRange2d         RadarAttackLayer7TopLocationTargetParamsDateRange = "2d"
	RadarAttackLayer7TopLocationTargetParamsDateRange7d         RadarAttackLayer7TopLocationTargetParamsDateRange = "7d"
	RadarAttackLayer7TopLocationTargetParamsDateRange14d        RadarAttackLayer7TopLocationTargetParamsDateRange = "14d"
	RadarAttackLayer7TopLocationTargetParamsDateRange28d        RadarAttackLayer7TopLocationTargetParamsDateRange = "28d"
	RadarAttackLayer7TopLocationTargetParamsDateRange12w        RadarAttackLayer7TopLocationTargetParamsDateRange = "12w"
	RadarAttackLayer7TopLocationTargetParamsDateRange24w        RadarAttackLayer7TopLocationTargetParamsDateRange = "24w"
	RadarAttackLayer7TopLocationTargetParamsDateRange52w        RadarAttackLayer7TopLocationTargetParamsDateRange = "52w"
	RadarAttackLayer7TopLocationTargetParamsDateRange1dControl  RadarAttackLayer7TopLocationTargetParamsDateRange = "1dControl"
	RadarAttackLayer7TopLocationTargetParamsDateRange2dControl  RadarAttackLayer7TopLocationTargetParamsDateRange = "2dControl"
	RadarAttackLayer7TopLocationTargetParamsDateRange7dControl  RadarAttackLayer7TopLocationTargetParamsDateRange = "7dControl"
	RadarAttackLayer7TopLocationTargetParamsDateRange14dControl RadarAttackLayer7TopLocationTargetParamsDateRange = "14dControl"
	RadarAttackLayer7TopLocationTargetParamsDateRange28dControl RadarAttackLayer7TopLocationTargetParamsDateRange = "28dControl"
	RadarAttackLayer7TopLocationTargetParamsDateRange12wControl RadarAttackLayer7TopLocationTargetParamsDateRange = "12wControl"
	RadarAttackLayer7TopLocationTargetParamsDateRange24wControl RadarAttackLayer7TopLocationTargetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopLocationTargetParamsFormat string

const (
	RadarAttackLayer7TopLocationTargetParamsFormatJson RadarAttackLayer7TopLocationTargetParamsFormat = "JSON"
	RadarAttackLayer7TopLocationTargetParamsFormatCsv  RadarAttackLayer7TopLocationTargetParamsFormat = "CSV"
)

type RadarAttackLayer7TopLocationTargetResponseEnvelope struct {
	Result  RadarAttackLayer7TopLocationTargetResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAttackLayer7TopLocationTargetResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopLocationTargetResponseEnvelope]
type radarAttackLayer7TopLocationTargetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
