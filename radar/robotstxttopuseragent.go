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

// RobotsTXTTopUserAgentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotsTXTTopUserAgentService] method instead.
type RobotsTXTTopUserAgentService struct {
	Options []option.RequestOption
}

// NewRobotsTXTTopUserAgentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRobotsTXTTopUserAgentService(opts ...option.RequestOption) (r *RobotsTXTTopUserAgentService) {
	r = &RobotsTXTTopUserAgentService{}
	r.Options = opts
	return
}

// Get the top user agents on robots.txt files by directive.
func (r *RobotsTXTTopUserAgentService) Directive(ctx context.Context, query RobotsTXTTopUserAgentDirectiveParams, opts ...option.RequestOption) (res *RobotsTXTTopUserAgentDirectiveResponse, err error) {
	var env RobotsTXTTopUserAgentDirectiveResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/robots_txt/top/user_agents/directive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RobotsTXTTopUserAgentDirectiveResponse struct {
	Meta RobotsTXTTopUserAgentDirectiveResponseMeta   `json:"meta,required"`
	Top0 []RobotsTXTTopUserAgentDirectiveResponseTop0 `json:"top_0,required"`
	JSON robotsTXTTopUserAgentDirectiveResponseJSON   `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseJSON contains the JSON metadata for the
// struct [RobotsTXTTopUserAgentDirectiveResponse]
type robotsTXTTopUserAgentDirectiveResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMeta struct {
	DateRange      []RobotsTXTTopUserAgentDirectiveResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo `json:"confidenceInfo"`
	Units          []RobotsTXTTopUserAgentDirectiveResponseMetaUnit         `json:"units"`
	JSON           robotsTXTTopUserAgentDirectiveResponseMetaJSON           `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaJSON contains the JSON metadata for
// the struct [RobotsTXTTopUserAgentDirectiveResponseMeta]
type robotsTXTTopUserAgentDirectiveResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RobotsTXTTopUserAgentDirectiveResponseMetaDateRange]
type robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo struct {
	Annotations []RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON         `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo]
type robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation]
type robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON contains the JSON metadata
// for the struct [RobotsTXTTopUserAgentDirectiveResponseMetaUnit]
type robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseTop0 struct {
	Name      string                                         `json:"name,required"`
	Value     int64                                          `json:"value,required"`
	Fully     int64                                          `json:"fully"`
	Partially int64                                          `json:"partially"`
	JSON      robotsTXTTopUserAgentDirectiveResponseTop0JSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseTop0JSON contains the JSON metadata for
// the struct [RobotsTXTTopUserAgentDirectiveResponseTop0]
type robotsTXTTopUserAgentDirectiveResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Fully       apijson.Field
	Partially   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveParams struct {
	// Array of dates to filter the ranking.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Filter by directive.
	Directive param.Field[RobotsTXTTopUserAgentDirectiveParamsDirective] `query:"directive"`
	// Filter by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Format results are returned in.
	Format param.Field[RobotsTXTTopUserAgentDirectiveParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter by user agent category.
	UserAgentCategory param.Field[RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory] `query:"userAgentCategory"`
}

// URLQuery serializes [RobotsTXTTopUserAgentDirectiveParams]'s query parameters as
// `url.Values`.
func (r RobotsTXTTopUserAgentDirectiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by directive.
type RobotsTXTTopUserAgentDirectiveParamsDirective string

const (
	RobotsTXTTopUserAgentDirectiveParamsDirectiveAllow    RobotsTXTTopUserAgentDirectiveParamsDirective = "ALLOW"
	RobotsTXTTopUserAgentDirectiveParamsDirectiveDisallow RobotsTXTTopUserAgentDirectiveParamsDirective = "DISALLOW"
)

func (r RobotsTXTTopUserAgentDirectiveParamsDirective) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveParamsDirectiveAllow, RobotsTXTTopUserAgentDirectiveParamsDirectiveDisallow:
		return true
	}
	return false
}

// Format results are returned in.
type RobotsTXTTopUserAgentDirectiveParamsFormat string

const (
	RobotsTXTTopUserAgentDirectiveParamsFormatJson RobotsTXTTopUserAgentDirectiveParamsFormat = "JSON"
	RobotsTXTTopUserAgentDirectiveParamsFormatCsv  RobotsTXTTopUserAgentDirectiveParamsFormat = "CSV"
)

func (r RobotsTXTTopUserAgentDirectiveParamsFormat) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveParamsFormatJson, RobotsTXTTopUserAgentDirectiveParamsFormatCsv:
		return true
	}
	return false
}

// Filter by user agent category.
type RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory string

const (
	RobotsTXTTopUserAgentDirectiveParamsUserAgentCategoryAI RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory = "AI"
)

func (r RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveParamsUserAgentCategoryAI:
		return true
	}
	return false
}

type RobotsTXTTopUserAgentDirectiveResponseEnvelope struct {
	Result  RobotsTXTTopUserAgentDirectiveResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON contains the JSON metadata
// for the struct [RobotsTXTTopUserAgentDirectiveResponseEnvelope]
type robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
