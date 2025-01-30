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

// RobotsTXTTopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotsTXTTopService] method instead.
type RobotsTXTTopService struct {
	Options    []option.RequestOption
	UserAgents *RobotsTXTTopUserAgentService
}

// NewRobotsTXTTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRobotsTXTTopService(opts ...option.RequestOption) (r *RobotsTXTTopService) {
	r = &RobotsTXTTopService{}
	r.Options = opts
	r.UserAgents = NewRobotsTXTTopUserAgentService(opts...)
	return
}

// Get the top domain categories by the number of robots.txt files parsed.
func (r *RobotsTXTTopService) DomainCategories(ctx context.Context, query RobotsTXTTopDomainCategoriesParams, opts ...option.RequestOption) (res *RobotsTXTTopDomainCategoriesResponse, err error) {
	var env RobotsTXTTopDomainCategoriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/robots_txt/top/domain_categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RobotsTXTTopDomainCategoriesResponse struct {
	Meta RobotsTXTTopDomainCategoriesResponseMeta   `json:"meta,required"`
	Top0 []RobotsTXTTopDomainCategoriesResponseTop0 `json:"top_0,required"`
	JSON robotsTXTTopDomainCategoriesResponseJSON   `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseJSON contains the JSON metadata for the
// struct [RobotsTXTTopDomainCategoriesResponse]
type robotsTXTTopDomainCategoriesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMeta struct {
	DateRange      []RobotsTXTTopDomainCategoriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	Units          []RobotsTXTTopDomainCategoriesResponseMetaUnit         `json:"units"`
	JSON           robotsTXTTopDomainCategoriesResponseMetaJSON           `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaJSON contains the JSON metadata for the
// struct [RobotsTXTTopDomainCategoriesResponseMeta]
type robotsTXTTopDomainCategoriesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RobotsTXTTopDomainCategoriesResponseMetaDateRange]
type robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo struct {
	Annotations []RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo]
type robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation]
type robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  robotsTXTTopDomainCategoriesResponseMetaUnitJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaUnitJSON contains the JSON metadata for
// the struct [RobotsTXTTopDomainCategoriesResponseMetaUnit]
type robotsTXTTopDomainCategoriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseTop0 struct {
	Name  string                                       `json:"name,required"`
	Value int64                                        `json:"value,required"`
	JSON  robotsTXTTopDomainCategoriesResponseTop0JSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseTop0JSON contains the JSON metadata for the
// struct [RobotsTXTTopDomainCategoriesResponseTop0]
type robotsTXTTopDomainCategoriesResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format results are returned in.
	Format param.Field[RobotsTXTTopDomainCategoriesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter by user agent category.
	UserAgentCategory param.Field[RobotsTXTTopDomainCategoriesParamsUserAgentCategory] `query:"userAgentCategory"`
}

// URLQuery serializes [RobotsTXTTopDomainCategoriesParams]'s query parameters as
// `url.Values`.
func (r RobotsTXTTopDomainCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type RobotsTXTTopDomainCategoriesParamsFormat string

const (
	RobotsTXTTopDomainCategoriesParamsFormatJson RobotsTXTTopDomainCategoriesParamsFormat = "JSON"
	RobotsTXTTopDomainCategoriesParamsFormatCsv  RobotsTXTTopDomainCategoriesParamsFormat = "CSV"
)

func (r RobotsTXTTopDomainCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RobotsTXTTopDomainCategoriesParamsFormatJson, RobotsTXTTopDomainCategoriesParamsFormatCsv:
		return true
	}
	return false
}

// Filter by user agent category.
type RobotsTXTTopDomainCategoriesParamsUserAgentCategory string

const (
	RobotsTXTTopDomainCategoriesParamsUserAgentCategoryAI RobotsTXTTopDomainCategoriesParamsUserAgentCategory = "AI"
)

func (r RobotsTXTTopDomainCategoriesParamsUserAgentCategory) IsKnown() bool {
	switch r {
	case RobotsTXTTopDomainCategoriesParamsUserAgentCategoryAI:
		return true
	}
	return false
}

type RobotsTXTTopDomainCategoriesResponseEnvelope struct {
	Result  RobotsTXTTopDomainCategoriesResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    robotsTXTTopDomainCategoriesResponseEnvelopeJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RobotsTXTTopDomainCategoriesResponseEnvelope]
type robotsTXTTopDomainCategoriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
