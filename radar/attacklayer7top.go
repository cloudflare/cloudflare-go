// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AttackLayer7TopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TopService] method instead.
type AttackLayer7TopService struct {
	Options   []option.RequestOption
	Locations *AttackLayer7TopLocationService
	Ases      *AttackLayer7TopAseService
}

// NewAttackLayer7TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer7TopService(opts ...option.RequestOption) (r *AttackLayer7TopService) {
	r = &AttackLayer7TopService{}
	r.Options = opts
	r.Locations = NewAttackLayer7TopLocationService(opts...)
	r.Ases = NewAttackLayer7TopAseService(opts...)
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 7 attacks (with billing country). The attack magnitude can be
// defined by the number of mitigated requests or by the number of zones affected.
// You can optionally limit the number of attacks per origin/target location
// (useful if all the top attacks are from or to the same location).
func (r *AttackLayer7TopService) Attacks(ctx context.Context, query AttackLayer7TopAttacksParams, opts ...option.RequestOption) (res *AttackLayer7TopAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TopAttacksResponseEnvelope
	path := "radar/attacks/layer7/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Industry of attacks.
func (r *AttackLayer7TopService) Industry(ctx context.Context, query AttackLayer7TopIndustryParams, opts ...option.RequestOption) (res *AttackLayer7TopIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TopIndustryResponseEnvelope
	path := "radar/attacks/layer7/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Verticals of attacks.
func (r *AttackLayer7TopService) Vertical(ctx context.Context, query AttackLayer7TopVerticalParams, opts ...option.RequestOption) (res *AttackLayer7TopVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TopVerticalResponseEnvelope
	path := "radar/attacks/layer7/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TopAttacksResponse struct {
	Meta AttackLayer7TopAttacksResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopAttacksResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopAttacksResponseJSON   `json:"-"`
}

// attackLayer7TopAttacksResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopAttacksResponse]
type attackLayer7TopAttacksResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseMeta struct {
	DateRange      []AttackLayer7TopAttacksResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer7TopAttacksResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TopAttacksResponseMetaJSON           `json:"-"`
}

// attackLayer7TopAttacksResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer7TopAttacksResponseMeta]
type attackLayer7TopAttacksResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopAttacksResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopAttacksResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopAttacksResponseMetaDateRange]
type attackLayer7TopAttacksResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        attackLayer7TopAttacksResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TopAttacksResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TopAttacksResponseMetaConfidenceInfo]
type attackLayer7TopAttacksResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseTop0 struct {
	OriginCountryAlpha2 string                                 `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                 `json:"originCountryName,required"`
	TargetCountryAlpha2 string                                 `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                 `json:"targetCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                attackLayer7TopAttacksResponseTop0JSON `json:"-"`
}

// attackLayer7TopAttacksResponseTop0JSON contains the JSON metadata for the struct
// [AttackLayer7TopAttacksResponseTop0]
type attackLayer7TopAttacksResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponse struct {
	Meta AttackLayer7TopIndustryResponseMeta `json:"meta,required"`
	Top0 []Browser                           `json:"top_0,required"`
	JSON attackLayer7TopIndustryResponseJSON `json:"-"`
}

// attackLayer7TopIndustryResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopIndustryResponse]
type attackLayer7TopIndustryResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseMeta struct {
	DateRange      []AttackLayer7TopIndustryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer7TopIndustryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TopIndustryResponseMetaJSON           `json:"-"`
}

// attackLayer7TopIndustryResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopIndustryResponseMeta]
type attackLayer7TopIndustryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopIndustryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopIndustryResponseMetaDateRange]
type attackLayer7TopIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        attackLayer7TopIndustryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TopIndustryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TopIndustryResponseMetaConfidenceInfo]
type attackLayer7TopIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponse struct {
	Meta AttackLayer7TopVerticalResponseMeta `json:"meta,required"`
	Top0 []Browser                           `json:"top_0,required"`
	JSON attackLayer7TopVerticalResponseJSON `json:"-"`
}

// attackLayer7TopVerticalResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopVerticalResponse]
type attackLayer7TopVerticalResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseMeta struct {
	DateRange      []AttackLayer7TopVerticalResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer7TopVerticalResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TopVerticalResponseMetaJSON           `json:"-"`
}

// attackLayer7TopVerticalResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopVerticalResponseMeta]
type attackLayer7TopVerticalResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopVerticalResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopVerticalResponseMetaDateRange]
type attackLayer7TopVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        attackLayer7TopVerticalResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TopVerticalResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TopVerticalResponseMetaConfidenceInfo]
type attackLayer7TopVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksParams struct {
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
	DateRange param.Field[[]AttackLayer7TopAttacksParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TopAttacksParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[AttackLayer7TopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limit the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Attack magnitude can be defined by total requests mitigated or by total zones
	// attacked.
	Magnitude param.Field[AttackLayer7TopAttacksParamsMagnitude] `query:"magnitude"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TopAttacksParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TopAttacksParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer7TopAttacksParamsDateRange string

const (
	AttackLayer7TopAttacksParamsDateRange1d         AttackLayer7TopAttacksParamsDateRange = "1d"
	AttackLayer7TopAttacksParamsDateRange2d         AttackLayer7TopAttacksParamsDateRange = "2d"
	AttackLayer7TopAttacksParamsDateRange7d         AttackLayer7TopAttacksParamsDateRange = "7d"
	AttackLayer7TopAttacksParamsDateRange14d        AttackLayer7TopAttacksParamsDateRange = "14d"
	AttackLayer7TopAttacksParamsDateRange28d        AttackLayer7TopAttacksParamsDateRange = "28d"
	AttackLayer7TopAttacksParamsDateRange12w        AttackLayer7TopAttacksParamsDateRange = "12w"
	AttackLayer7TopAttacksParamsDateRange24w        AttackLayer7TopAttacksParamsDateRange = "24w"
	AttackLayer7TopAttacksParamsDateRange52w        AttackLayer7TopAttacksParamsDateRange = "52w"
	AttackLayer7TopAttacksParamsDateRange1dControl  AttackLayer7TopAttacksParamsDateRange = "1dControl"
	AttackLayer7TopAttacksParamsDateRange2dControl  AttackLayer7TopAttacksParamsDateRange = "2dControl"
	AttackLayer7TopAttacksParamsDateRange7dControl  AttackLayer7TopAttacksParamsDateRange = "7dControl"
	AttackLayer7TopAttacksParamsDateRange14dControl AttackLayer7TopAttacksParamsDateRange = "14dControl"
	AttackLayer7TopAttacksParamsDateRange28dControl AttackLayer7TopAttacksParamsDateRange = "28dControl"
	AttackLayer7TopAttacksParamsDateRange12wControl AttackLayer7TopAttacksParamsDateRange = "12wControl"
	AttackLayer7TopAttacksParamsDateRange24wControl AttackLayer7TopAttacksParamsDateRange = "24wControl"
)

func (r AttackLayer7TopAttacksParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsDateRange1d, AttackLayer7TopAttacksParamsDateRange2d, AttackLayer7TopAttacksParamsDateRange7d, AttackLayer7TopAttacksParamsDateRange14d, AttackLayer7TopAttacksParamsDateRange28d, AttackLayer7TopAttacksParamsDateRange12w, AttackLayer7TopAttacksParamsDateRange24w, AttackLayer7TopAttacksParamsDateRange52w, AttackLayer7TopAttacksParamsDateRange1dControl, AttackLayer7TopAttacksParamsDateRange2dControl, AttackLayer7TopAttacksParamsDateRange7dControl, AttackLayer7TopAttacksParamsDateRange14dControl, AttackLayer7TopAttacksParamsDateRange28dControl, AttackLayer7TopAttacksParamsDateRange12wControl, AttackLayer7TopAttacksParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TopAttacksParamsFormat string

const (
	AttackLayer7TopAttacksParamsFormatJson AttackLayer7TopAttacksParamsFormat = "JSON"
	AttackLayer7TopAttacksParamsFormatCsv  AttackLayer7TopAttacksParamsFormat = "CSV"
)

func (r AttackLayer7TopAttacksParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsFormatJson, AttackLayer7TopAttacksParamsFormatCsv:
		return true
	}
	return false
}

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type AttackLayer7TopAttacksParamsLimitDirection string

const (
	AttackLayer7TopAttacksParamsLimitDirectionOrigin AttackLayer7TopAttacksParamsLimitDirection = "ORIGIN"
	AttackLayer7TopAttacksParamsLimitDirectionTarget AttackLayer7TopAttacksParamsLimitDirection = "TARGET"
)

func (r AttackLayer7TopAttacksParamsLimitDirection) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsLimitDirectionOrigin, AttackLayer7TopAttacksParamsLimitDirectionTarget:
		return true
	}
	return false
}

// Attack magnitude can be defined by total requests mitigated or by total zones
// attacked.
type AttackLayer7TopAttacksParamsMagnitude string

const (
	AttackLayer7TopAttacksParamsMagnitudeAffectedZones     AttackLayer7TopAttacksParamsMagnitude = "AFFECTED_ZONES"
	AttackLayer7TopAttacksParamsMagnitudeMitigatedRequests AttackLayer7TopAttacksParamsMagnitude = "MITIGATED_REQUESTS"
)

func (r AttackLayer7TopAttacksParamsMagnitude) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsMagnitudeAffectedZones, AttackLayer7TopAttacksParamsMagnitudeMitigatedRequests:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopAttacksParamsNormalization string

const (
	AttackLayer7TopAttacksParamsNormalizationPercentage AttackLayer7TopAttacksParamsNormalization = "PERCENTAGE"
	AttackLayer7TopAttacksParamsNormalizationMinMax     AttackLayer7TopAttacksParamsNormalization = "MIN_MAX"
)

func (r AttackLayer7TopAttacksParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsNormalizationPercentage, AttackLayer7TopAttacksParamsNormalizationMinMax:
		return true
	}
	return false
}

type AttackLayer7TopAttacksResponseEnvelope struct {
	Result  AttackLayer7TopAttacksResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer7TopAttacksResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopAttacksResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopAttacksResponseEnvelope]
type attackLayer7TopAttacksResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryParams struct {
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
	DateRange param.Field[[]AttackLayer7TopIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TopIndustryParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopIndustryParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer7TopIndustryParamsDateRange string

const (
	AttackLayer7TopIndustryParamsDateRange1d         AttackLayer7TopIndustryParamsDateRange = "1d"
	AttackLayer7TopIndustryParamsDateRange2d         AttackLayer7TopIndustryParamsDateRange = "2d"
	AttackLayer7TopIndustryParamsDateRange7d         AttackLayer7TopIndustryParamsDateRange = "7d"
	AttackLayer7TopIndustryParamsDateRange14d        AttackLayer7TopIndustryParamsDateRange = "14d"
	AttackLayer7TopIndustryParamsDateRange28d        AttackLayer7TopIndustryParamsDateRange = "28d"
	AttackLayer7TopIndustryParamsDateRange12w        AttackLayer7TopIndustryParamsDateRange = "12w"
	AttackLayer7TopIndustryParamsDateRange24w        AttackLayer7TopIndustryParamsDateRange = "24w"
	AttackLayer7TopIndustryParamsDateRange52w        AttackLayer7TopIndustryParamsDateRange = "52w"
	AttackLayer7TopIndustryParamsDateRange1dControl  AttackLayer7TopIndustryParamsDateRange = "1dControl"
	AttackLayer7TopIndustryParamsDateRange2dControl  AttackLayer7TopIndustryParamsDateRange = "2dControl"
	AttackLayer7TopIndustryParamsDateRange7dControl  AttackLayer7TopIndustryParamsDateRange = "7dControl"
	AttackLayer7TopIndustryParamsDateRange14dControl AttackLayer7TopIndustryParamsDateRange = "14dControl"
	AttackLayer7TopIndustryParamsDateRange28dControl AttackLayer7TopIndustryParamsDateRange = "28dControl"
	AttackLayer7TopIndustryParamsDateRange12wControl AttackLayer7TopIndustryParamsDateRange = "12wControl"
	AttackLayer7TopIndustryParamsDateRange24wControl AttackLayer7TopIndustryParamsDateRange = "24wControl"
)

func (r AttackLayer7TopIndustryParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsDateRange1d, AttackLayer7TopIndustryParamsDateRange2d, AttackLayer7TopIndustryParamsDateRange7d, AttackLayer7TopIndustryParamsDateRange14d, AttackLayer7TopIndustryParamsDateRange28d, AttackLayer7TopIndustryParamsDateRange12w, AttackLayer7TopIndustryParamsDateRange24w, AttackLayer7TopIndustryParamsDateRange52w, AttackLayer7TopIndustryParamsDateRange1dControl, AttackLayer7TopIndustryParamsDateRange2dControl, AttackLayer7TopIndustryParamsDateRange7dControl, AttackLayer7TopIndustryParamsDateRange14dControl, AttackLayer7TopIndustryParamsDateRange28dControl, AttackLayer7TopIndustryParamsDateRange12wControl, AttackLayer7TopIndustryParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TopIndustryParamsFormat string

const (
	AttackLayer7TopIndustryParamsFormatJson AttackLayer7TopIndustryParamsFormat = "JSON"
	AttackLayer7TopIndustryParamsFormatCsv  AttackLayer7TopIndustryParamsFormat = "CSV"
)

func (r AttackLayer7TopIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsFormatJson, AttackLayer7TopIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopIndustryResponseEnvelope struct {
	Result  AttackLayer7TopIndustryResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer7TopIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopIndustryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopIndustryResponseEnvelope]
type attackLayer7TopIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalParams struct {
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
	DateRange param.Field[[]AttackLayer7TopVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TopVerticalParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopVerticalParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer7TopVerticalParamsDateRange string

const (
	AttackLayer7TopVerticalParamsDateRange1d         AttackLayer7TopVerticalParamsDateRange = "1d"
	AttackLayer7TopVerticalParamsDateRange2d         AttackLayer7TopVerticalParamsDateRange = "2d"
	AttackLayer7TopVerticalParamsDateRange7d         AttackLayer7TopVerticalParamsDateRange = "7d"
	AttackLayer7TopVerticalParamsDateRange14d        AttackLayer7TopVerticalParamsDateRange = "14d"
	AttackLayer7TopVerticalParamsDateRange28d        AttackLayer7TopVerticalParamsDateRange = "28d"
	AttackLayer7TopVerticalParamsDateRange12w        AttackLayer7TopVerticalParamsDateRange = "12w"
	AttackLayer7TopVerticalParamsDateRange24w        AttackLayer7TopVerticalParamsDateRange = "24w"
	AttackLayer7TopVerticalParamsDateRange52w        AttackLayer7TopVerticalParamsDateRange = "52w"
	AttackLayer7TopVerticalParamsDateRange1dControl  AttackLayer7TopVerticalParamsDateRange = "1dControl"
	AttackLayer7TopVerticalParamsDateRange2dControl  AttackLayer7TopVerticalParamsDateRange = "2dControl"
	AttackLayer7TopVerticalParamsDateRange7dControl  AttackLayer7TopVerticalParamsDateRange = "7dControl"
	AttackLayer7TopVerticalParamsDateRange14dControl AttackLayer7TopVerticalParamsDateRange = "14dControl"
	AttackLayer7TopVerticalParamsDateRange28dControl AttackLayer7TopVerticalParamsDateRange = "28dControl"
	AttackLayer7TopVerticalParamsDateRange12wControl AttackLayer7TopVerticalParamsDateRange = "12wControl"
	AttackLayer7TopVerticalParamsDateRange24wControl AttackLayer7TopVerticalParamsDateRange = "24wControl"
)

func (r AttackLayer7TopVerticalParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsDateRange1d, AttackLayer7TopVerticalParamsDateRange2d, AttackLayer7TopVerticalParamsDateRange7d, AttackLayer7TopVerticalParamsDateRange14d, AttackLayer7TopVerticalParamsDateRange28d, AttackLayer7TopVerticalParamsDateRange12w, AttackLayer7TopVerticalParamsDateRange24w, AttackLayer7TopVerticalParamsDateRange52w, AttackLayer7TopVerticalParamsDateRange1dControl, AttackLayer7TopVerticalParamsDateRange2dControl, AttackLayer7TopVerticalParamsDateRange7dControl, AttackLayer7TopVerticalParamsDateRange14dControl, AttackLayer7TopVerticalParamsDateRange28dControl, AttackLayer7TopVerticalParamsDateRange12wControl, AttackLayer7TopVerticalParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TopVerticalParamsFormat string

const (
	AttackLayer7TopVerticalParamsFormatJson AttackLayer7TopVerticalParamsFormat = "JSON"
	AttackLayer7TopVerticalParamsFormatCsv  AttackLayer7TopVerticalParamsFormat = "CSV"
)

func (r AttackLayer7TopVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsFormatJson, AttackLayer7TopVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopVerticalResponseEnvelope struct {
	Result  AttackLayer7TopVerticalResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer7TopVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopVerticalResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopVerticalResponseEnvelope]
type attackLayer7TopVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
