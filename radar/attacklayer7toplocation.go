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

// AttackLayer7TopLocationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TopLocationService] method instead.
type AttackLayer7TopLocationService struct {
	Options []option.RequestOption
}

// NewAttackLayer7TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer7TopLocationService(opts ...option.RequestOption) (r *AttackLayer7TopLocationService) {
	r = &AttackLayer7TopLocationService{}
	r.Options = opts
	return
}

// Get the top origin locations of and by layer 7 attacks. Values are a percentage
// out of the total layer 7 attacks. The origin location is determined by the
// client IP.
func (r *AttackLayer7TopLocationService) Origin(ctx context.Context, query AttackLayer7TopLocationOriginParams, opts ...option.RequestOption) (res *AttackLayer7TopLocationOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TopLocationOriginResponseEnvelope
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
func (r *AttackLayer7TopLocationService) Target(ctx context.Context, query AttackLayer7TopLocationTargetParams, opts ...option.RequestOption) (res *AttackLayer7TopLocationTargetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TopLocationTargetResponseEnvelope
	path := "radar/attacks/layer7/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TopLocationOriginResponse struct {
	Meta AttackLayer7TopLocationOriginResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopLocationOriginResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopLocationOriginResponseJSON   `json:"-"`
}

// attackLayer7TopLocationOriginResponseJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationOriginResponse]
type attackLayer7TopLocationOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseMeta struct {
	DateRange      []AttackLayer7TopLocationOriginResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer7TopLocationOriginResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TopLocationOriginResponseMetaJSON           `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationOriginResponseMeta]
type attackLayer7TopLocationOriginResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopLocationOriginResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7TopLocationOriginResponseMetaDateRange]
type attackLayer7TopLocationOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7TopLocationOriginResponseMetaConfidenceInfo]
type attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseTop0 struct {
	OriginCountryAlpha2 string                                        `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                        `json:"originCountryName,required"`
	Rank                float64                                       `json:"rank,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer7TopLocationOriginResponseTop0JSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationOriginResponseTop0]
type attackLayer7TopLocationOriginResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponse struct {
	Meta AttackLayer7TopLocationTargetResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopLocationTargetResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopLocationTargetResponseJSON   `json:"-"`
}

// attackLayer7TopLocationTargetResponseJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationTargetResponse]
type attackLayer7TopLocationTargetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseMeta struct {
	DateRange      []AttackLayer7TopLocationTargetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer7TopLocationTargetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TopLocationTargetResponseMetaJSON           `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationTargetResponseMeta]
type attackLayer7TopLocationTargetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopLocationTargetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7TopLocationTargetResponseMetaDateRange]
type attackLayer7TopLocationTargetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7TopLocationTargetResponseMetaConfidenceInfo]
type attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseTop0 struct {
	Rank                float64                                       `json:"rank,required"`
	TargetCountryAlpha2 string                                        `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                        `json:"targetCountryName,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer7TopLocationTargetResponseTop0JSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationTargetResponseTop0]
type attackLayer7TopLocationTargetResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginParams struct {
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
	DateRange param.Field[[]AttackLayer7TopLocationOriginParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TopLocationOriginParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopLocationOriginParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopLocationOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AttackLayer7TopLocationOriginParamsDateRange string

const (
	AttackLayer7TopLocationOriginParamsDateRange1d         AttackLayer7TopLocationOriginParamsDateRange = "1d"
	AttackLayer7TopLocationOriginParamsDateRange2d         AttackLayer7TopLocationOriginParamsDateRange = "2d"
	AttackLayer7TopLocationOriginParamsDateRange7d         AttackLayer7TopLocationOriginParamsDateRange = "7d"
	AttackLayer7TopLocationOriginParamsDateRange14d        AttackLayer7TopLocationOriginParamsDateRange = "14d"
	AttackLayer7TopLocationOriginParamsDateRange28d        AttackLayer7TopLocationOriginParamsDateRange = "28d"
	AttackLayer7TopLocationOriginParamsDateRange12w        AttackLayer7TopLocationOriginParamsDateRange = "12w"
	AttackLayer7TopLocationOriginParamsDateRange24w        AttackLayer7TopLocationOriginParamsDateRange = "24w"
	AttackLayer7TopLocationOriginParamsDateRange52w        AttackLayer7TopLocationOriginParamsDateRange = "52w"
	AttackLayer7TopLocationOriginParamsDateRange1dControl  AttackLayer7TopLocationOriginParamsDateRange = "1dControl"
	AttackLayer7TopLocationOriginParamsDateRange2dControl  AttackLayer7TopLocationOriginParamsDateRange = "2dControl"
	AttackLayer7TopLocationOriginParamsDateRange7dControl  AttackLayer7TopLocationOriginParamsDateRange = "7dControl"
	AttackLayer7TopLocationOriginParamsDateRange14dControl AttackLayer7TopLocationOriginParamsDateRange = "14dControl"
	AttackLayer7TopLocationOriginParamsDateRange28dControl AttackLayer7TopLocationOriginParamsDateRange = "28dControl"
	AttackLayer7TopLocationOriginParamsDateRange12wControl AttackLayer7TopLocationOriginParamsDateRange = "12wControl"
	AttackLayer7TopLocationOriginParamsDateRange24wControl AttackLayer7TopLocationOriginParamsDateRange = "24wControl"
)

func (r AttackLayer7TopLocationOriginParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsDateRange1d, AttackLayer7TopLocationOriginParamsDateRange2d, AttackLayer7TopLocationOriginParamsDateRange7d, AttackLayer7TopLocationOriginParamsDateRange14d, AttackLayer7TopLocationOriginParamsDateRange28d, AttackLayer7TopLocationOriginParamsDateRange12w, AttackLayer7TopLocationOriginParamsDateRange24w, AttackLayer7TopLocationOriginParamsDateRange52w, AttackLayer7TopLocationOriginParamsDateRange1dControl, AttackLayer7TopLocationOriginParamsDateRange2dControl, AttackLayer7TopLocationOriginParamsDateRange7dControl, AttackLayer7TopLocationOriginParamsDateRange14dControl, AttackLayer7TopLocationOriginParamsDateRange28dControl, AttackLayer7TopLocationOriginParamsDateRange12wControl, AttackLayer7TopLocationOriginParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TopLocationOriginParamsFormat string

const (
	AttackLayer7TopLocationOriginParamsFormatJson AttackLayer7TopLocationOriginParamsFormat = "JSON"
	AttackLayer7TopLocationOriginParamsFormatCsv  AttackLayer7TopLocationOriginParamsFormat = "CSV"
)

func (r AttackLayer7TopLocationOriginParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsFormatJson, AttackLayer7TopLocationOriginParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginResponseEnvelope struct {
	Result  AttackLayer7TopLocationOriginResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer7TopLocationOriginResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7TopLocationOriginResponseEnvelope]
type attackLayer7TopLocationOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer7TopLocationTargetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TopLocationTargetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopLocationTargetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopLocationTargetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AttackLayer7TopLocationTargetParamsDateRange string

const (
	AttackLayer7TopLocationTargetParamsDateRange1d         AttackLayer7TopLocationTargetParamsDateRange = "1d"
	AttackLayer7TopLocationTargetParamsDateRange2d         AttackLayer7TopLocationTargetParamsDateRange = "2d"
	AttackLayer7TopLocationTargetParamsDateRange7d         AttackLayer7TopLocationTargetParamsDateRange = "7d"
	AttackLayer7TopLocationTargetParamsDateRange14d        AttackLayer7TopLocationTargetParamsDateRange = "14d"
	AttackLayer7TopLocationTargetParamsDateRange28d        AttackLayer7TopLocationTargetParamsDateRange = "28d"
	AttackLayer7TopLocationTargetParamsDateRange12w        AttackLayer7TopLocationTargetParamsDateRange = "12w"
	AttackLayer7TopLocationTargetParamsDateRange24w        AttackLayer7TopLocationTargetParamsDateRange = "24w"
	AttackLayer7TopLocationTargetParamsDateRange52w        AttackLayer7TopLocationTargetParamsDateRange = "52w"
	AttackLayer7TopLocationTargetParamsDateRange1dControl  AttackLayer7TopLocationTargetParamsDateRange = "1dControl"
	AttackLayer7TopLocationTargetParamsDateRange2dControl  AttackLayer7TopLocationTargetParamsDateRange = "2dControl"
	AttackLayer7TopLocationTargetParamsDateRange7dControl  AttackLayer7TopLocationTargetParamsDateRange = "7dControl"
	AttackLayer7TopLocationTargetParamsDateRange14dControl AttackLayer7TopLocationTargetParamsDateRange = "14dControl"
	AttackLayer7TopLocationTargetParamsDateRange28dControl AttackLayer7TopLocationTargetParamsDateRange = "28dControl"
	AttackLayer7TopLocationTargetParamsDateRange12wControl AttackLayer7TopLocationTargetParamsDateRange = "12wControl"
	AttackLayer7TopLocationTargetParamsDateRange24wControl AttackLayer7TopLocationTargetParamsDateRange = "24wControl"
)

func (r AttackLayer7TopLocationTargetParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationTargetParamsDateRange1d, AttackLayer7TopLocationTargetParamsDateRange2d, AttackLayer7TopLocationTargetParamsDateRange7d, AttackLayer7TopLocationTargetParamsDateRange14d, AttackLayer7TopLocationTargetParamsDateRange28d, AttackLayer7TopLocationTargetParamsDateRange12w, AttackLayer7TopLocationTargetParamsDateRange24w, AttackLayer7TopLocationTargetParamsDateRange52w, AttackLayer7TopLocationTargetParamsDateRange1dControl, AttackLayer7TopLocationTargetParamsDateRange2dControl, AttackLayer7TopLocationTargetParamsDateRange7dControl, AttackLayer7TopLocationTargetParamsDateRange14dControl, AttackLayer7TopLocationTargetParamsDateRange28dControl, AttackLayer7TopLocationTargetParamsDateRange12wControl, AttackLayer7TopLocationTargetParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TopLocationTargetParamsFormat string

const (
	AttackLayer7TopLocationTargetParamsFormatJson AttackLayer7TopLocationTargetParamsFormat = "JSON"
	AttackLayer7TopLocationTargetParamsFormatCsv  AttackLayer7TopLocationTargetParamsFormat = "CSV"
)

func (r AttackLayer7TopLocationTargetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationTargetParamsFormatJson, AttackLayer7TopLocationTargetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopLocationTargetResponseEnvelope struct {
	Result  AttackLayer7TopLocationTargetResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer7TopLocationTargetResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7TopLocationTargetResponseEnvelope]
type attackLayer7TopLocationTargetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
