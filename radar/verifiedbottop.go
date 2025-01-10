// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// VerifiedBotTopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifiedBotTopService] method instead.
type VerifiedBotTopService struct {
	Options []option.RequestOption
}

// NewVerifiedBotTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerifiedBotTopService(opts ...option.RequestOption) (r *VerifiedBotTopService) {
	r = &VerifiedBotTopService{}
	r.Options = opts
	return
}

// Get top verified bots by HTTP requests, with owner and category.
func (r *VerifiedBotTopService) Bots(ctx context.Context, query VerifiedBotTopBotsParams, opts ...option.RequestOption) (res *VerifiedBotTopBotsResponse, err error) {
	var env VerifiedBotTopBotsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/bots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get top verified bot categories by HTTP requests, along with their corresponding
// percentage, over the total verified bot HTTP requests.
func (r *VerifiedBotTopService) Categories(ctx context.Context, query VerifiedBotTopCategoriesParams, opts ...option.RequestOption) (res *VerifiedBotTopCategoriesResponse, err error) {
	var env VerifiedBotTopCategoriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VerifiedBotTopBotsResponse struct {
	Meta VerifiedBotTopBotsResponseMeta   `json:"meta,required"`
	Top0 []VerifiedBotTopBotsResponseTop0 `json:"top_0,required"`
	JSON verifiedBotTopBotsResponseJSON   `json:"-"`
}

// verifiedBotTopBotsResponseJSON contains the JSON metadata for the struct
// [VerifiedBotTopBotsResponse]
type verifiedBotTopBotsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopBotsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopBotsResponseMeta struct {
	DateRange      []VerifiedBotTopBotsResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo VerifiedBotTopBotsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           verifiedBotTopBotsResponseMetaJSON           `json:"-"`
}

// verifiedBotTopBotsResponseMetaJSON contains the JSON metadata for the struct
// [VerifiedBotTopBotsResponseMeta]
type verifiedBotTopBotsResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VerifiedBotTopBotsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopBotsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      verifiedBotTopBotsResponseMetaDateRangeJSON `json:"-"`
}

// verifiedBotTopBotsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [VerifiedBotTopBotsResponseMetaDateRange]
type verifiedBotTopBotsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopBotsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopBotsResponseMetaConfidenceInfo struct {
	Annotations []VerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        verifiedBotTopBotsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// verifiedBotTopBotsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [VerifiedBotTopBotsResponseMetaConfidenceInfo]
type verifiedBotTopBotsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopBotsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                     `json:"dataSource,required"`
	Description     string                                                     `json:"description,required"`
	EventType       string                                                     `json:"eventType,required"`
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                  `json:"startTime" format:"date-time"`
	JSON            verifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// verifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [VerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation]
type verifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *VerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopBotsResponseTop0 struct {
	BotCategory string                             `json:"botCategory,required"`
	BotName     string                             `json:"botName,required"`
	BotOwner    string                             `json:"botOwner,required"`
	Value       string                             `json:"value,required"`
	JSON        verifiedBotTopBotsResponseTop0JSON `json:"-"`
}

// verifiedBotTopBotsResponseTop0JSON contains the JSON metadata for the struct
// [VerifiedBotTopBotsResponseTop0]
type verifiedBotTopBotsResponseTop0JSON struct {
	BotCategory apijson.Field
	BotName     apijson.Field
	BotOwner    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopBotsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesResponse struct {
	Meta VerifiedBotTopCategoriesResponseMeta   `json:"meta,required"`
	Top0 []VerifiedBotTopCategoriesResponseTop0 `json:"top_0,required"`
	JSON verifiedBotTopCategoriesResponseJSON   `json:"-"`
}

// verifiedBotTopCategoriesResponseJSON contains the JSON metadata for the struct
// [VerifiedBotTopCategoriesResponse]
type verifiedBotTopCategoriesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesResponseMeta struct {
	DateRange      []VerifiedBotTopCategoriesResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo VerifiedBotTopCategoriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           verifiedBotTopCategoriesResponseMetaJSON           `json:"-"`
}

// verifiedBotTopCategoriesResponseMetaJSON contains the JSON metadata for the
// struct [VerifiedBotTopCategoriesResponseMeta]
type verifiedBotTopCategoriesResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VerifiedBotTopCategoriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      verifiedBotTopCategoriesResponseMetaDateRangeJSON `json:"-"`
}

// verifiedBotTopCategoriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [VerifiedBotTopCategoriesResponseMetaDateRange]
type verifiedBotTopCategoriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopCategoriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesResponseMetaConfidenceInfo struct {
	Annotations []VerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        verifiedBotTopCategoriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// verifiedBotTopCategoriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [VerifiedBotTopCategoriesResponseMetaConfidenceInfo]
type verifiedBotTopCategoriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopCategoriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            verifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// verifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [VerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation]
type verifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *VerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesResponseTop0 struct {
	BotCategory string                                   `json:"botCategory,required"`
	Value       string                                   `json:"value,required"`
	JSON        verifiedBotTopCategoriesResponseTop0JSON `json:"-"`
}

// verifiedBotTopCategoriesResponseTop0JSON contains the JSON metadata for the
// struct [VerifiedBotTopCategoriesResponseTop0]
type verifiedBotTopCategoriesResponseTop0JSON struct {
	BotCategory apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopCategoriesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopBotsParams struct {
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
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[VerifiedBotTopBotsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [VerifiedBotTopBotsParams]'s query parameters as
// `url.Values`.
func (r VerifiedBotTopBotsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type VerifiedBotTopBotsParamsFormat string

const (
	VerifiedBotTopBotsParamsFormatJson VerifiedBotTopBotsParamsFormat = "JSON"
	VerifiedBotTopBotsParamsFormatCsv  VerifiedBotTopBotsParamsFormat = "CSV"
)

func (r VerifiedBotTopBotsParamsFormat) IsKnown() bool {
	switch r {
	case VerifiedBotTopBotsParamsFormatJson, VerifiedBotTopBotsParamsFormatCsv:
		return true
	}
	return false
}

type VerifiedBotTopBotsResponseEnvelope struct {
	Result  VerifiedBotTopBotsResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    verifiedBotTopBotsResponseEnvelopeJSON `json:"-"`
}

// verifiedBotTopBotsResponseEnvelopeJSON contains the JSON metadata for the struct
// [VerifiedBotTopBotsResponseEnvelope]
type verifiedBotTopBotsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopBotsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopBotsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type VerifiedBotTopCategoriesParams struct {
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
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[VerifiedBotTopCategoriesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [VerifiedBotTopCategoriesParams]'s query parameters as
// `url.Values`.
func (r VerifiedBotTopCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type VerifiedBotTopCategoriesParamsFormat string

const (
	VerifiedBotTopCategoriesParamsFormatJson VerifiedBotTopCategoriesParamsFormat = "JSON"
	VerifiedBotTopCategoriesParamsFormatCsv  VerifiedBotTopCategoriesParamsFormat = "CSV"
)

func (r VerifiedBotTopCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case VerifiedBotTopCategoriesParamsFormatJson, VerifiedBotTopCategoriesParamsFormatCsv:
		return true
	}
	return false
}

type VerifiedBotTopCategoriesResponseEnvelope struct {
	Result  VerifiedBotTopCategoriesResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    verifiedBotTopCategoriesResponseEnvelopeJSON `json:"-"`
}

// verifiedBotTopCategoriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [VerifiedBotTopCategoriesResponseEnvelope]
type verifiedBotTopCategoriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerifiedBotTopCategoriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verifiedBotTopCategoriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
