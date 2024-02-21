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

// RadarAttackLayer3TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer3TopService]
// method instead.
type RadarAttackLayer3TopService struct {
	Options   []option.RequestOption
	Locations *RadarAttackLayer3TopLocationService
}

// NewRadarAttackLayer3TopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopService(opts ...option.RequestOption) (r *RadarAttackLayer3TopService) {
	r = &RadarAttackLayer3TopService{}
	r.Options = opts
	r.Locations = NewRadarAttackLayer3TopLocationService(opts...)
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 3 attacks (with billing country). You can optionally limit
// the number of attacks per origin/target location (useful if all the top attacks
// are from or to the same location).
func (r *RadarAttackLayer3TopService) Attacks(ctx context.Context, query RadarAttackLayer3TopAttacksParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopAttacksResponseEnvelope
	path := "radar/attacks/layer3/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Industry of attacks.
func (r *RadarAttackLayer3TopService) Industry(ctx context.Context, query RadarAttackLayer3TopIndustryParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopIndustryResponseEnvelope
	path := "radar/attacks/layer3/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Verticals of attacks.
func (r *RadarAttackLayer3TopService) Vertical(ctx context.Context, query RadarAttackLayer3TopVerticalParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TopVerticalResponseEnvelope
	path := "radar/attacks/layer3/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TopAttacksResponse struct {
	Meta RadarAttackLayer3TopAttacksResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopAttacksResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopAttacksResponseJSON   `json:"-"`
}

// radarAttackLayer3TopAttacksResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopAttacksResponse]
type radarAttackLayer3TopAttacksResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttacksResponseMeta struct {
	DateRange      []RadarAttackLayer3TopAttacksResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopAttacksResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopAttacksResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopAttacksResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopAttacksResponseMeta]
type radarAttackLayer3TopAttacksResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttacksResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttacksResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopAttacksResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopAttacksResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopAttacksResponseMetaDateRange]
type radarAttackLayer3TopAttacksResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttacksResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttacksResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarAttackLayer3TopAttacksResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopAttacksResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopAttacksResponseMetaConfidenceInfo]
type radarAttackLayer3TopAttacksResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttacksResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttacksResponseTop0 struct {
	OriginCountryAlpha2 string                                      `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                      `json:"originCountryName,required"`
	Value               string                                      `json:"value,required"`
	JSON                radarAttackLayer3TopAttacksResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopAttacksResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopAttacksResponseTop0]
type radarAttackLayer3TopAttacksResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttacksResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryResponse struct {
	Meta RadarAttackLayer3TopIndustryResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopIndustryResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopIndustryResponseJSON   `json:"-"`
}

// radarAttackLayer3TopIndustryResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopIndustryResponse]
type radarAttackLayer3TopIndustryResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryResponseMeta struct {
	DateRange      []RadarAttackLayer3TopIndustryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopIndustryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopIndustryResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopIndustryResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopIndustryResponseMeta]
type radarAttackLayer3TopIndustryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopIndustryResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopIndustryResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopIndustryResponseMetaDateRange]
type radarAttackLayer3TopIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarAttackLayer3TopIndustryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopIndustryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopIndustryResponseMetaConfidenceInfo]
type radarAttackLayer3TopIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryResponseTop0 struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  radarAttackLayer3TopIndustryResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopIndustryResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopIndustryResponseTop0]
type radarAttackLayer3TopIndustryResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalResponse struct {
	Meta RadarAttackLayer3TopVerticalResponseMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopVerticalResponseTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopVerticalResponseJSON   `json:"-"`
}

// radarAttackLayer3TopVerticalResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopVerticalResponse]
type radarAttackLayer3TopVerticalResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalResponseMeta struct {
	DateRange      []RadarAttackLayer3TopVerticalResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopVerticalResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopVerticalResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TopVerticalResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopVerticalResponseMeta]
type radarAttackLayer3TopVerticalResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopVerticalResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopVerticalResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopVerticalResponseMetaDateRange]
type radarAttackLayer3TopVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarAttackLayer3TopVerticalResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopVerticalResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopVerticalResponseMetaConfidenceInfo]
type radarAttackLayer3TopVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalResponseTop0 struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  radarAttackLayer3TopVerticalResponseTop0JSON `json:"-"`
}

// radarAttackLayer3TopVerticalResponseTop0JSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopVerticalResponseTop0]
type radarAttackLayer3TopVerticalResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttacksParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopAttacksParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopAttacksParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopAttacksParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer3TopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limit the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopAttacksParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopAttacksParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3TopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopAttacksParamsDateRange string

const (
	RadarAttackLayer3TopAttacksParamsDateRange1d         RadarAttackLayer3TopAttacksParamsDateRange = "1d"
	RadarAttackLayer3TopAttacksParamsDateRange2d         RadarAttackLayer3TopAttacksParamsDateRange = "2d"
	RadarAttackLayer3TopAttacksParamsDateRange7d         RadarAttackLayer3TopAttacksParamsDateRange = "7d"
	RadarAttackLayer3TopAttacksParamsDateRange14d        RadarAttackLayer3TopAttacksParamsDateRange = "14d"
	RadarAttackLayer3TopAttacksParamsDateRange28d        RadarAttackLayer3TopAttacksParamsDateRange = "28d"
	RadarAttackLayer3TopAttacksParamsDateRange12w        RadarAttackLayer3TopAttacksParamsDateRange = "12w"
	RadarAttackLayer3TopAttacksParamsDateRange24w        RadarAttackLayer3TopAttacksParamsDateRange = "24w"
	RadarAttackLayer3TopAttacksParamsDateRange52w        RadarAttackLayer3TopAttacksParamsDateRange = "52w"
	RadarAttackLayer3TopAttacksParamsDateRange1dControl  RadarAttackLayer3TopAttacksParamsDateRange = "1dControl"
	RadarAttackLayer3TopAttacksParamsDateRange2dControl  RadarAttackLayer3TopAttacksParamsDateRange = "2dControl"
	RadarAttackLayer3TopAttacksParamsDateRange7dControl  RadarAttackLayer3TopAttacksParamsDateRange = "7dControl"
	RadarAttackLayer3TopAttacksParamsDateRange14dControl RadarAttackLayer3TopAttacksParamsDateRange = "14dControl"
	RadarAttackLayer3TopAttacksParamsDateRange28dControl RadarAttackLayer3TopAttacksParamsDateRange = "28dControl"
	RadarAttackLayer3TopAttacksParamsDateRange12wControl RadarAttackLayer3TopAttacksParamsDateRange = "12wControl"
	RadarAttackLayer3TopAttacksParamsDateRange24wControl RadarAttackLayer3TopAttacksParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopAttacksParamsFormat string

const (
	RadarAttackLayer3TopAttacksParamsFormatJson RadarAttackLayer3TopAttacksParamsFormat = "JSON"
	RadarAttackLayer3TopAttacksParamsFormatCsv  RadarAttackLayer3TopAttacksParamsFormat = "CSV"
)

type RadarAttackLayer3TopAttacksParamsIPVersion string

const (
	RadarAttackLayer3TopAttacksParamsIPVersionIPv4 RadarAttackLayer3TopAttacksParamsIPVersion = "IPv4"
	RadarAttackLayer3TopAttacksParamsIPVersionIPv6 RadarAttackLayer3TopAttacksParamsIPVersion = "IPv6"
)

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer3TopAttacksParamsLimitDirection string

const (
	RadarAttackLayer3TopAttacksParamsLimitDirectionOrigin RadarAttackLayer3TopAttacksParamsLimitDirection = "ORIGIN"
	RadarAttackLayer3TopAttacksParamsLimitDirectionTarget RadarAttackLayer3TopAttacksParamsLimitDirection = "TARGET"
)

type RadarAttackLayer3TopAttacksParamsProtocol string

const (
	RadarAttackLayer3TopAttacksParamsProtocolUdp  RadarAttackLayer3TopAttacksParamsProtocol = "UDP"
	RadarAttackLayer3TopAttacksParamsProtocolTcp  RadarAttackLayer3TopAttacksParamsProtocol = "TCP"
	RadarAttackLayer3TopAttacksParamsProtocolIcmp RadarAttackLayer3TopAttacksParamsProtocol = "ICMP"
	RadarAttackLayer3TopAttacksParamsProtocolGre  RadarAttackLayer3TopAttacksParamsProtocol = "GRE"
)

type RadarAttackLayer3TopAttacksResponseEnvelope struct {
	Result  RadarAttackLayer3TopAttacksResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer3TopAttacksResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopAttacksResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopAttacksResponseEnvelope]
type radarAttackLayer3TopAttacksResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttacksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopIndustryParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopIndustryParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopIndustryParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopIndustryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopIndustryParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3TopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopIndustryParamsDateRange string

const (
	RadarAttackLayer3TopIndustryParamsDateRange1d         RadarAttackLayer3TopIndustryParamsDateRange = "1d"
	RadarAttackLayer3TopIndustryParamsDateRange2d         RadarAttackLayer3TopIndustryParamsDateRange = "2d"
	RadarAttackLayer3TopIndustryParamsDateRange7d         RadarAttackLayer3TopIndustryParamsDateRange = "7d"
	RadarAttackLayer3TopIndustryParamsDateRange14d        RadarAttackLayer3TopIndustryParamsDateRange = "14d"
	RadarAttackLayer3TopIndustryParamsDateRange28d        RadarAttackLayer3TopIndustryParamsDateRange = "28d"
	RadarAttackLayer3TopIndustryParamsDateRange12w        RadarAttackLayer3TopIndustryParamsDateRange = "12w"
	RadarAttackLayer3TopIndustryParamsDateRange24w        RadarAttackLayer3TopIndustryParamsDateRange = "24w"
	RadarAttackLayer3TopIndustryParamsDateRange52w        RadarAttackLayer3TopIndustryParamsDateRange = "52w"
	RadarAttackLayer3TopIndustryParamsDateRange1dControl  RadarAttackLayer3TopIndustryParamsDateRange = "1dControl"
	RadarAttackLayer3TopIndustryParamsDateRange2dControl  RadarAttackLayer3TopIndustryParamsDateRange = "2dControl"
	RadarAttackLayer3TopIndustryParamsDateRange7dControl  RadarAttackLayer3TopIndustryParamsDateRange = "7dControl"
	RadarAttackLayer3TopIndustryParamsDateRange14dControl RadarAttackLayer3TopIndustryParamsDateRange = "14dControl"
	RadarAttackLayer3TopIndustryParamsDateRange28dControl RadarAttackLayer3TopIndustryParamsDateRange = "28dControl"
	RadarAttackLayer3TopIndustryParamsDateRange12wControl RadarAttackLayer3TopIndustryParamsDateRange = "12wControl"
	RadarAttackLayer3TopIndustryParamsDateRange24wControl RadarAttackLayer3TopIndustryParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopIndustryParamsFormat string

const (
	RadarAttackLayer3TopIndustryParamsFormatJson RadarAttackLayer3TopIndustryParamsFormat = "JSON"
	RadarAttackLayer3TopIndustryParamsFormatCsv  RadarAttackLayer3TopIndustryParamsFormat = "CSV"
)

type RadarAttackLayer3TopIndustryParamsIPVersion string

const (
	RadarAttackLayer3TopIndustryParamsIPVersionIPv4 RadarAttackLayer3TopIndustryParamsIPVersion = "IPv4"
	RadarAttackLayer3TopIndustryParamsIPVersionIPv6 RadarAttackLayer3TopIndustryParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopIndustryParamsProtocol string

const (
	RadarAttackLayer3TopIndustryParamsProtocolUdp  RadarAttackLayer3TopIndustryParamsProtocol = "UDP"
	RadarAttackLayer3TopIndustryParamsProtocolTcp  RadarAttackLayer3TopIndustryParamsProtocol = "TCP"
	RadarAttackLayer3TopIndustryParamsProtocolIcmp RadarAttackLayer3TopIndustryParamsProtocol = "ICMP"
	RadarAttackLayer3TopIndustryParamsProtocolGre  RadarAttackLayer3TopIndustryParamsProtocol = "GRE"
)

type RadarAttackLayer3TopIndustryResponseEnvelope struct {
	Result  RadarAttackLayer3TopIndustryResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer3TopIndustryResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopIndustryResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopIndustryResponseEnvelope]
type radarAttackLayer3TopIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopVerticalParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopVerticalParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopVerticalParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopVerticalParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopVerticalParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3TopVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopVerticalParamsDateRange string

const (
	RadarAttackLayer3TopVerticalParamsDateRange1d         RadarAttackLayer3TopVerticalParamsDateRange = "1d"
	RadarAttackLayer3TopVerticalParamsDateRange2d         RadarAttackLayer3TopVerticalParamsDateRange = "2d"
	RadarAttackLayer3TopVerticalParamsDateRange7d         RadarAttackLayer3TopVerticalParamsDateRange = "7d"
	RadarAttackLayer3TopVerticalParamsDateRange14d        RadarAttackLayer3TopVerticalParamsDateRange = "14d"
	RadarAttackLayer3TopVerticalParamsDateRange28d        RadarAttackLayer3TopVerticalParamsDateRange = "28d"
	RadarAttackLayer3TopVerticalParamsDateRange12w        RadarAttackLayer3TopVerticalParamsDateRange = "12w"
	RadarAttackLayer3TopVerticalParamsDateRange24w        RadarAttackLayer3TopVerticalParamsDateRange = "24w"
	RadarAttackLayer3TopVerticalParamsDateRange52w        RadarAttackLayer3TopVerticalParamsDateRange = "52w"
	RadarAttackLayer3TopVerticalParamsDateRange1dControl  RadarAttackLayer3TopVerticalParamsDateRange = "1dControl"
	RadarAttackLayer3TopVerticalParamsDateRange2dControl  RadarAttackLayer3TopVerticalParamsDateRange = "2dControl"
	RadarAttackLayer3TopVerticalParamsDateRange7dControl  RadarAttackLayer3TopVerticalParamsDateRange = "7dControl"
	RadarAttackLayer3TopVerticalParamsDateRange14dControl RadarAttackLayer3TopVerticalParamsDateRange = "14dControl"
	RadarAttackLayer3TopVerticalParamsDateRange28dControl RadarAttackLayer3TopVerticalParamsDateRange = "28dControl"
	RadarAttackLayer3TopVerticalParamsDateRange12wControl RadarAttackLayer3TopVerticalParamsDateRange = "12wControl"
	RadarAttackLayer3TopVerticalParamsDateRange24wControl RadarAttackLayer3TopVerticalParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopVerticalParamsFormat string

const (
	RadarAttackLayer3TopVerticalParamsFormatJson RadarAttackLayer3TopVerticalParamsFormat = "JSON"
	RadarAttackLayer3TopVerticalParamsFormatCsv  RadarAttackLayer3TopVerticalParamsFormat = "CSV"
)

type RadarAttackLayer3TopVerticalParamsIPVersion string

const (
	RadarAttackLayer3TopVerticalParamsIPVersionIPv4 RadarAttackLayer3TopVerticalParamsIPVersion = "IPv4"
	RadarAttackLayer3TopVerticalParamsIPVersionIPv6 RadarAttackLayer3TopVerticalParamsIPVersion = "IPv6"
)

type RadarAttackLayer3TopVerticalParamsProtocol string

const (
	RadarAttackLayer3TopVerticalParamsProtocolUdp  RadarAttackLayer3TopVerticalParamsProtocol = "UDP"
	RadarAttackLayer3TopVerticalParamsProtocolTcp  RadarAttackLayer3TopVerticalParamsProtocol = "TCP"
	RadarAttackLayer3TopVerticalParamsProtocolIcmp RadarAttackLayer3TopVerticalParamsProtocol = "ICMP"
	RadarAttackLayer3TopVerticalParamsProtocolGre  RadarAttackLayer3TopVerticalParamsProtocol = "GRE"
)

type RadarAttackLayer3TopVerticalResponseEnvelope struct {
	Result  RadarAttackLayer3TopVerticalResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer3TopVerticalResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TopVerticalResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopVerticalResponseEnvelope]
type radarAttackLayer3TopVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
