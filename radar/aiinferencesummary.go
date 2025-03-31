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

// AIInferenceSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIInferenceSummaryService] method instead.
type AIInferenceSummaryService struct {
	Options []option.RequestOption
}

// NewAIInferenceSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIInferenceSummaryService(opts ...option.RequestOption) (r *AIInferenceSummaryService) {
	r = &AIInferenceSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of unique accounts by model.
func (r *AIInferenceSummaryService) Model(ctx context.Context, query AIInferenceSummaryModelParams, opts ...option.RequestOption) (res *AIInferenceSummaryModelResponse, err error) {
	var env AIInferenceSummaryModelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of unique accounts by task.
func (r *AIInferenceSummaryService) Task(ctx context.Context, query AIInferenceSummaryTaskParams, opts ...option.RequestOption) (res *AIInferenceSummaryTaskResponse, err error) {
	var env AIInferenceSummaryTaskResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIInferenceSummaryModelResponse struct {
	Meta     AIInferenceSummaryModelResponseMeta `json:"meta,required"`
	Summary0 map[string]string                   `json:"summary_0,required"`
	JSON     aiInferenceSummaryModelResponseJSON `json:"-"`
}

// aiInferenceSummaryModelResponseJSON contains the JSON metadata for the struct
// [AIInferenceSummaryModelResponse]
type aiInferenceSummaryModelResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelResponseMeta struct {
	DateRange      []AIInferenceSummaryModelResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	Normalization  string                                            `json:"normalization,required"`
	ConfidenceInfo AIInferenceSummaryModelResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           aiInferenceSummaryModelResponseMetaJSON           `json:"-"`
}

// aiInferenceSummaryModelResponseMetaJSON contains the JSON metadata for the
// struct [AIInferenceSummaryModelResponseMeta]
type aiInferenceSummaryModelResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceSummaryModelResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceSummaryModelResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIInferenceSummaryModelResponseMetaDateRange]
type aiInferenceSummaryModelResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        aiInferenceSummaryModelResponseMetaConfidenceInfoJSON         `json:"-"`
}

// aiInferenceSummaryModelResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIInferenceSummaryModelResponseMetaConfidenceInfo]
type aiInferenceSummaryModelResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation]
type aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponse struct {
	Meta     AIInferenceSummaryTaskResponseMeta `json:"meta,required"`
	Summary0 map[string]string                  `json:"summary_0,required"`
	JSON     aiInferenceSummaryTaskResponseJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseJSON contains the JSON metadata for the struct
// [AIInferenceSummaryTaskResponse]
type aiInferenceSummaryTaskResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponseMeta struct {
	DateRange      []AIInferenceSummaryTaskResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	Normalization  string                                           `json:"normalization,required"`
	ConfidenceInfo AIInferenceSummaryTaskResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           aiInferenceSummaryTaskResponseMetaJSON           `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaJSON contains the JSON metadata for the struct
// [AIInferenceSummaryTaskResponseMeta]
type aiInferenceSummaryTaskResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceSummaryTaskResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIInferenceSummaryTaskResponseMetaDateRange]
type aiInferenceSummaryTaskResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON         `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIInferenceSummaryTaskResponseMetaConfidenceInfo]
type aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation]
type aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceSummaryModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceSummaryModelParams]'s query parameters as
// `url.Values`.
func (r AIInferenceSummaryModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AIInferenceSummaryModelParamsFormat string

const (
	AIInferenceSummaryModelParamsFormatJson AIInferenceSummaryModelParamsFormat = "JSON"
	AIInferenceSummaryModelParamsFormatCsv  AIInferenceSummaryModelParamsFormat = "CSV"
)

func (r AIInferenceSummaryModelParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceSummaryModelParamsFormatJson, AIInferenceSummaryModelParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceSummaryModelResponseEnvelope struct {
	Result  AIInferenceSummaryModelResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    aiInferenceSummaryModelResponseEnvelopeJSON `json:"-"`
}

// aiInferenceSummaryModelResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIInferenceSummaryModelResponseEnvelope]
type aiInferenceSummaryModelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceSummaryTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceSummaryTaskParams]'s query parameters as
// `url.Values`.
func (r AIInferenceSummaryTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AIInferenceSummaryTaskParamsFormat string

const (
	AIInferenceSummaryTaskParamsFormatJson AIInferenceSummaryTaskParamsFormat = "JSON"
	AIInferenceSummaryTaskParamsFormatCsv  AIInferenceSummaryTaskParamsFormat = "CSV"
)

func (r AIInferenceSummaryTaskParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceSummaryTaskParamsFormatJson, AIInferenceSummaryTaskParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceSummaryTaskResponseEnvelope struct {
	Result  AIInferenceSummaryTaskResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    aiInferenceSummaryTaskResponseEnvelopeJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIInferenceSummaryTaskResponseEnvelope]
type aiInferenceSummaryTaskResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
