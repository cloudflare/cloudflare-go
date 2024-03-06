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

// RadarAttackLayer7TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer7TopService]
// method instead.
type RadarAttackLayer7TopService struct {
	Options   []option.RequestOption
	Locations *RadarAttackLayer7TopLocationService
	Ases      *RadarAttackLayer7TopAseService
}

// NewRadarAttackLayer7TopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopService(opts ...option.RequestOption) (r *RadarAttackLayer7TopService) {
	r = &RadarAttackLayer7TopService{}
	r.Options = opts
	r.Locations = NewRadarAttackLayer7TopLocationService(opts...)
	r.Ases = NewRadarAttackLayer7TopAseService(opts...)
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 7 attacks (with billing country). The attack magnitude can be
// defined by the number of mitigated requests or by the number of zones affected.
// You can optionally limit the number of attacks per origin/target location
// (useful if all the top attacks are from or to the same location).
func (r *RadarAttackLayer7TopService) Attacks(ctx context.Context, query RadarAttackLayer7TopAttacksParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TopAttacksResponseEnvelope
	path := "radar/attacks/layer7/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Industry of attacks.
func (r *RadarAttackLayer7TopService) Industry(ctx context.Context, query RadarAttackLayer7TopIndustryParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TopIndustryResponseEnvelope
	path := "radar/attacks/layer7/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Verticals of attacks.
func (r *RadarAttackLayer7TopService) Vertical(ctx context.Context, query RadarAttackLayer7TopVerticalParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TopVerticalResponseEnvelope
	path := "radar/attacks/layer7/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7TopAttacksResponse struct {
	Meta RadarAttackLayer7TopAttacksResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopAttacksResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopAttacksResponseJSON   `json:"-"`
}

// radarAttackLayer7TopAttacksResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAttacksResponse]
type radarAttackLayer7TopAttacksResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAttacksResponseMeta struct {
	DateRange      []RadarAttackLayer7TopAttacksResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopAttacksResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopAttacksResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TopAttacksResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAttacksResponseMeta]
type radarAttackLayer7TopAttacksResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttacksResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAttacksResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopAttacksResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopAttacksResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopAttacksResponseMetaDateRange]
type radarAttackLayer7TopAttacksResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttacksResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAttacksResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarAttackLayer7TopAttacksResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopAttacksResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopAttacksResponseMetaConfidenceInfo]
type radarAttackLayer7TopAttacksResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttacksResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAttacksResponseTop0 struct {
	OriginCountryAlpha2 string                                      `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                      `json:"originCountryName,required"`
	TargetCountryAlpha2 string                                      `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                      `json:"targetCountryName,required"`
	Value               string                                      `json:"value,required"`
	JSON                radarAttackLayer7TopAttacksResponseTop0JSON `json:"-"`
}

// radarAttackLayer7TopAttacksResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAttacksResponseTop0]
type radarAttackLayer7TopAttacksResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttacksResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryResponse struct {
	Meta RadarAttackLayer7TopIndustryResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopIndustryResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopIndustryResponseJSON   `json:"-"`
}

// radarAttackLayer7TopIndustryResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopIndustryResponse]
type radarAttackLayer7TopIndustryResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryResponseMeta struct {
	DateRange      []RadarAttackLayer7TopIndustryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopIndustryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopIndustryResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TopIndustryResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopIndustryResponseMeta]
type radarAttackLayer7TopIndustryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopIndustryResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopIndustryResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopIndustryResponseMetaDateRange]
type radarAttackLayer7TopIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarAttackLayer7TopIndustryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopIndustryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopIndustryResponseMetaConfidenceInfo]
type radarAttackLayer7TopIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryResponseTop0 struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  radarAttackLayer7TopIndustryResponseTop0JSON `json:"-"`
}

// radarAttackLayer7TopIndustryResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopIndustryResponseTop0]
type radarAttackLayer7TopIndustryResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopIndustryResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalResponse struct {
	Meta RadarAttackLayer7TopVerticalResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopVerticalResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopVerticalResponseJSON   `json:"-"`
}

// radarAttackLayer7TopVerticalResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopVerticalResponse]
type radarAttackLayer7TopVerticalResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalResponseMeta struct {
	DateRange      []RadarAttackLayer7TopVerticalResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopVerticalResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopVerticalResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TopVerticalResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopVerticalResponseMeta]
type radarAttackLayer7TopVerticalResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopVerticalResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopVerticalResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopVerticalResponseMetaDateRange]
type radarAttackLayer7TopVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarAttackLayer7TopVerticalResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopVerticalResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopVerticalResponseMetaConfidenceInfo]
type radarAttackLayer7TopVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalResponseTop0 struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  radarAttackLayer7TopVerticalResponseTop0JSON `json:"-"`
}

// radarAttackLayer7TopVerticalResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopVerticalResponseTop0]
type radarAttackLayer7TopVerticalResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopVerticalResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAttacksParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopAttacksParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopAttacksParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer7TopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limit the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Attack magnitude can be defined by total requests mitigated or by total zones
	// attacked.
	Magnitude param.Field[RadarAttackLayer7TopAttacksParamsMagnitude] `query:"magnitude"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopAttacksParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopAttacksParamsDateRange string

const (
	RadarAttackLayer7TopAttacksParamsDateRange1d         RadarAttackLayer7TopAttacksParamsDateRange = "1d"
	RadarAttackLayer7TopAttacksParamsDateRange2d         RadarAttackLayer7TopAttacksParamsDateRange = "2d"
	RadarAttackLayer7TopAttacksParamsDateRange7d         RadarAttackLayer7TopAttacksParamsDateRange = "7d"
	RadarAttackLayer7TopAttacksParamsDateRange14d        RadarAttackLayer7TopAttacksParamsDateRange = "14d"
	RadarAttackLayer7TopAttacksParamsDateRange28d        RadarAttackLayer7TopAttacksParamsDateRange = "28d"
	RadarAttackLayer7TopAttacksParamsDateRange12w        RadarAttackLayer7TopAttacksParamsDateRange = "12w"
	RadarAttackLayer7TopAttacksParamsDateRange24w        RadarAttackLayer7TopAttacksParamsDateRange = "24w"
	RadarAttackLayer7TopAttacksParamsDateRange52w        RadarAttackLayer7TopAttacksParamsDateRange = "52w"
	RadarAttackLayer7TopAttacksParamsDateRange1dControl  RadarAttackLayer7TopAttacksParamsDateRange = "1dControl"
	RadarAttackLayer7TopAttacksParamsDateRange2dControl  RadarAttackLayer7TopAttacksParamsDateRange = "2dControl"
	RadarAttackLayer7TopAttacksParamsDateRange7dControl  RadarAttackLayer7TopAttacksParamsDateRange = "7dControl"
	RadarAttackLayer7TopAttacksParamsDateRange14dControl RadarAttackLayer7TopAttacksParamsDateRange = "14dControl"
	RadarAttackLayer7TopAttacksParamsDateRange28dControl RadarAttackLayer7TopAttacksParamsDateRange = "28dControl"
	RadarAttackLayer7TopAttacksParamsDateRange12wControl RadarAttackLayer7TopAttacksParamsDateRange = "12wControl"
	RadarAttackLayer7TopAttacksParamsDateRange24wControl RadarAttackLayer7TopAttacksParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopAttacksParamsFormat string

const (
	RadarAttackLayer7TopAttacksParamsFormatJson RadarAttackLayer7TopAttacksParamsFormat = "JSON"
	RadarAttackLayer7TopAttacksParamsFormatCsv  RadarAttackLayer7TopAttacksParamsFormat = "CSV"
)

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer7TopAttacksParamsLimitDirection string

const (
	RadarAttackLayer7TopAttacksParamsLimitDirectionOrigin RadarAttackLayer7TopAttacksParamsLimitDirection = "ORIGIN"
	RadarAttackLayer7TopAttacksParamsLimitDirectionTarget RadarAttackLayer7TopAttacksParamsLimitDirection = "TARGET"
)

// Attack magnitude can be defined by total requests mitigated or by total zones
// attacked.
type RadarAttackLayer7TopAttacksParamsMagnitude string

const (
	RadarAttackLayer7TopAttacksParamsMagnitudeAffectedZones     RadarAttackLayer7TopAttacksParamsMagnitude = "AFFECTED_ZONES"
	RadarAttackLayer7TopAttacksParamsMagnitudeMitigatedRequests RadarAttackLayer7TopAttacksParamsMagnitude = "MITIGATED_REQUESTS"
)

type RadarAttackLayer7TopAttacksResponseEnvelope struct {
	Result  RadarAttackLayer7TopAttacksResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7TopAttacksResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TopAttacksResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopAttacksResponseEnvelope]
type radarAttackLayer7TopAttacksResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAttacksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAttacksResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopIndustryParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopIndustryParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopIndustryParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopIndustryParamsDateRange string

const (
	RadarAttackLayer7TopIndustryParamsDateRange1d         RadarAttackLayer7TopIndustryParamsDateRange = "1d"
	RadarAttackLayer7TopIndustryParamsDateRange2d         RadarAttackLayer7TopIndustryParamsDateRange = "2d"
	RadarAttackLayer7TopIndustryParamsDateRange7d         RadarAttackLayer7TopIndustryParamsDateRange = "7d"
	RadarAttackLayer7TopIndustryParamsDateRange14d        RadarAttackLayer7TopIndustryParamsDateRange = "14d"
	RadarAttackLayer7TopIndustryParamsDateRange28d        RadarAttackLayer7TopIndustryParamsDateRange = "28d"
	RadarAttackLayer7TopIndustryParamsDateRange12w        RadarAttackLayer7TopIndustryParamsDateRange = "12w"
	RadarAttackLayer7TopIndustryParamsDateRange24w        RadarAttackLayer7TopIndustryParamsDateRange = "24w"
	RadarAttackLayer7TopIndustryParamsDateRange52w        RadarAttackLayer7TopIndustryParamsDateRange = "52w"
	RadarAttackLayer7TopIndustryParamsDateRange1dControl  RadarAttackLayer7TopIndustryParamsDateRange = "1dControl"
	RadarAttackLayer7TopIndustryParamsDateRange2dControl  RadarAttackLayer7TopIndustryParamsDateRange = "2dControl"
	RadarAttackLayer7TopIndustryParamsDateRange7dControl  RadarAttackLayer7TopIndustryParamsDateRange = "7dControl"
	RadarAttackLayer7TopIndustryParamsDateRange14dControl RadarAttackLayer7TopIndustryParamsDateRange = "14dControl"
	RadarAttackLayer7TopIndustryParamsDateRange28dControl RadarAttackLayer7TopIndustryParamsDateRange = "28dControl"
	RadarAttackLayer7TopIndustryParamsDateRange12wControl RadarAttackLayer7TopIndustryParamsDateRange = "12wControl"
	RadarAttackLayer7TopIndustryParamsDateRange24wControl RadarAttackLayer7TopIndustryParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopIndustryParamsFormat string

const (
	RadarAttackLayer7TopIndustryParamsFormatJson RadarAttackLayer7TopIndustryParamsFormat = "JSON"
	RadarAttackLayer7TopIndustryParamsFormatCsv  RadarAttackLayer7TopIndustryParamsFormat = "CSV"
)

type RadarAttackLayer7TopIndustryResponseEnvelope struct {
	Result  RadarAttackLayer7TopIndustryResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer7TopIndustryResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TopIndustryResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopIndustryResponseEnvelope]
type radarAttackLayer7TopIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopVerticalParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopVerticalParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopVerticalParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TopVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopVerticalParamsDateRange string

const (
	RadarAttackLayer7TopVerticalParamsDateRange1d         RadarAttackLayer7TopVerticalParamsDateRange = "1d"
	RadarAttackLayer7TopVerticalParamsDateRange2d         RadarAttackLayer7TopVerticalParamsDateRange = "2d"
	RadarAttackLayer7TopVerticalParamsDateRange7d         RadarAttackLayer7TopVerticalParamsDateRange = "7d"
	RadarAttackLayer7TopVerticalParamsDateRange14d        RadarAttackLayer7TopVerticalParamsDateRange = "14d"
	RadarAttackLayer7TopVerticalParamsDateRange28d        RadarAttackLayer7TopVerticalParamsDateRange = "28d"
	RadarAttackLayer7TopVerticalParamsDateRange12w        RadarAttackLayer7TopVerticalParamsDateRange = "12w"
	RadarAttackLayer7TopVerticalParamsDateRange24w        RadarAttackLayer7TopVerticalParamsDateRange = "24w"
	RadarAttackLayer7TopVerticalParamsDateRange52w        RadarAttackLayer7TopVerticalParamsDateRange = "52w"
	RadarAttackLayer7TopVerticalParamsDateRange1dControl  RadarAttackLayer7TopVerticalParamsDateRange = "1dControl"
	RadarAttackLayer7TopVerticalParamsDateRange2dControl  RadarAttackLayer7TopVerticalParamsDateRange = "2dControl"
	RadarAttackLayer7TopVerticalParamsDateRange7dControl  RadarAttackLayer7TopVerticalParamsDateRange = "7dControl"
	RadarAttackLayer7TopVerticalParamsDateRange14dControl RadarAttackLayer7TopVerticalParamsDateRange = "14dControl"
	RadarAttackLayer7TopVerticalParamsDateRange28dControl RadarAttackLayer7TopVerticalParamsDateRange = "28dControl"
	RadarAttackLayer7TopVerticalParamsDateRange12wControl RadarAttackLayer7TopVerticalParamsDateRange = "12wControl"
	RadarAttackLayer7TopVerticalParamsDateRange24wControl RadarAttackLayer7TopVerticalParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopVerticalParamsFormat string

const (
	RadarAttackLayer7TopVerticalParamsFormatJson RadarAttackLayer7TopVerticalParamsFormat = "JSON"
	RadarAttackLayer7TopVerticalParamsFormatCsv  RadarAttackLayer7TopVerticalParamsFormat = "CSV"
)

type RadarAttackLayer7TopVerticalResponseEnvelope struct {
	Result  RadarAttackLayer7TopVerticalResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer7TopVerticalResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TopVerticalResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopVerticalResponseEnvelope]
type radarAttackLayer7TopVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
