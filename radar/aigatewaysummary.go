// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// AIGatewaySummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewaySummaryService] method instead.
type AIGatewaySummaryService struct {
	Options []option.RequestOption
}

// NewAIGatewaySummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIGatewaySummaryService(opts ...option.RequestOption) (r *AIGatewaySummaryService) {
	r = &AIGatewaySummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of unique accounts per model.
func (r *AIGatewaySummaryService) Model(ctx context.Context, query AIGatewaySummaryModelParams, opts ...option.RequestOption) (res *AIGatewaySummaryModelResponse, err error) {
	var env AIGatewaySummaryModelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of unique accounts per provider.
func (r *AIGatewaySummaryService) Provider(ctx context.Context, query AIGatewaySummaryProviderParams, opts ...option.RequestOption) (res *AIGatewaySummaryProviderResponse, err error) {
	var env AIGatewaySummaryProviderResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/gateway/summary/provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of unique accounts per task.
func (r *AIGatewaySummaryService) Task(ctx context.Context, query AIGatewaySummaryTaskParams, opts ...option.RequestOption) (res *AIGatewaySummaryTaskResponse, err error) {
	var env AIGatewaySummaryTaskResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIGatewaySummaryModelResponse struct {
	Meta     AIGatewaySummaryModelResponseMeta `json:"meta,required"`
	Summary0 map[string][]string               `json:"summary_0,required"`
	JSON     aiGatewaySummaryModelResponseJSON `json:"-"`
}

// aiGatewaySummaryModelResponseJSON contains the JSON metadata for the struct
// [AIGatewaySummaryModelResponse]
type aiGatewaySummaryModelResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryModelResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryModelResponseMeta struct {
	DateRange      []AIGatewaySummaryModelResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                          `json:"lastUpdated,required"`
	Normalization  string                                          `json:"normalization,required"`
	ConfidenceInfo AIGatewaySummaryModelResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           aiGatewaySummaryModelResponseMetaJSON           `json:"-"`
}

// aiGatewaySummaryModelResponseMetaJSON contains the JSON metadata for the struct
// [AIGatewaySummaryModelResponseMeta]
type aiGatewaySummaryModelResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIGatewaySummaryModelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryModelResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryModelResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      aiGatewaySummaryModelResponseMetaDateRangeJSON `json:"-"`
}

// aiGatewaySummaryModelResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIGatewaySummaryModelResponseMetaDateRange]
type aiGatewaySummaryModelResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryModelResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryModelResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryModelResponseMetaConfidenceInfo struct {
	Annotations []AIGatewaySummaryModelResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        aiGatewaySummaryModelResponseMetaConfidenceInfoJSON         `json:"-"`
}

// aiGatewaySummaryModelResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIGatewaySummaryModelResponseMetaConfidenceInfo]
type aiGatewaySummaryModelResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryModelResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryModelResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryModelResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                        `json:"dataSource,required"`
	Description     string                                                        `json:"description,required"`
	EventType       string                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                     `json:"startTime" format:"date-time"`
	JSON            aiGatewaySummaryModelResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiGatewaySummaryModelResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIGatewaySummaryModelResponseMetaConfidenceInfoAnnotation]
type aiGatewaySummaryModelResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIGatewaySummaryModelResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryModelResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryProviderResponse struct {
	Meta     AIGatewaySummaryProviderResponseMeta `json:"meta,required"`
	Summary0 map[string][]string                  `json:"summary_0,required"`
	JSON     aiGatewaySummaryProviderResponseJSON `json:"-"`
}

// aiGatewaySummaryProviderResponseJSON contains the JSON metadata for the struct
// [AIGatewaySummaryProviderResponse]
type aiGatewaySummaryProviderResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryProviderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryProviderResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryProviderResponseMeta struct {
	DateRange      []AIGatewaySummaryProviderResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	Normalization  string                                             `json:"normalization,required"`
	ConfidenceInfo AIGatewaySummaryProviderResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           aiGatewaySummaryProviderResponseMetaJSON           `json:"-"`
}

// aiGatewaySummaryProviderResponseMetaJSON contains the JSON metadata for the
// struct [AIGatewaySummaryProviderResponseMeta]
type aiGatewaySummaryProviderResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIGatewaySummaryProviderResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryProviderResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryProviderResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      aiGatewaySummaryProviderResponseMetaDateRangeJSON `json:"-"`
}

// aiGatewaySummaryProviderResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIGatewaySummaryProviderResponseMetaDateRange]
type aiGatewaySummaryProviderResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryProviderResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryProviderResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryProviderResponseMetaConfidenceInfo struct {
	Annotations []AIGatewaySummaryProviderResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        aiGatewaySummaryProviderResponseMetaConfidenceInfoJSON         `json:"-"`
}

// aiGatewaySummaryProviderResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AIGatewaySummaryProviderResponseMetaConfidenceInfo]
type aiGatewaySummaryProviderResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryProviderResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryProviderResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryProviderResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            aiGatewaySummaryProviderResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiGatewaySummaryProviderResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AIGatewaySummaryProviderResponseMetaConfidenceInfoAnnotation]
type aiGatewaySummaryProviderResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIGatewaySummaryProviderResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryProviderResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryTaskResponse struct {
	Meta     AIGatewaySummaryTaskResponseMeta `json:"meta,required"`
	Summary0 map[string][]string              `json:"summary_0,required"`
	JSON     aiGatewaySummaryTaskResponseJSON `json:"-"`
}

// aiGatewaySummaryTaskResponseJSON contains the JSON metadata for the struct
// [AIGatewaySummaryTaskResponse]
type aiGatewaySummaryTaskResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryTaskResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryTaskResponseMeta struct {
	DateRange      []AIGatewaySummaryTaskResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	Normalization  string                                         `json:"normalization,required"`
	ConfidenceInfo AIGatewaySummaryTaskResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           aiGatewaySummaryTaskResponseMetaJSON           `json:"-"`
}

// aiGatewaySummaryTaskResponseMetaJSON contains the JSON metadata for the struct
// [AIGatewaySummaryTaskResponseMeta]
type aiGatewaySummaryTaskResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIGatewaySummaryTaskResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryTaskResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryTaskResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      aiGatewaySummaryTaskResponseMetaDateRangeJSON `json:"-"`
}

// aiGatewaySummaryTaskResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AIGatewaySummaryTaskResponseMetaDateRange]
type aiGatewaySummaryTaskResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryTaskResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryTaskResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryTaskResponseMetaConfidenceInfo struct {
	Annotations []AIGatewaySummaryTaskResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        aiGatewaySummaryTaskResponseMetaConfidenceInfoJSON         `json:"-"`
}

// aiGatewaySummaryTaskResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIGatewaySummaryTaskResponseMetaConfidenceInfo]
type aiGatewaySummaryTaskResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryTaskResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryTaskResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryTaskResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            aiGatewaySummaryTaskResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiGatewaySummaryTaskResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIGatewaySummaryTaskResponseMetaConfidenceInfoAnnotation]
type aiGatewaySummaryTaskResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIGatewaySummaryTaskResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryTaskResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryModelParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AIGatewaySummaryModelParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIGatewaySummaryModelParams]'s query parameters as
// `url.Values`.
func (r AIGatewaySummaryModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AIGatewaySummaryModelParamsFormat string

const (
	AIGatewaySummaryModelParamsFormatJson AIGatewaySummaryModelParamsFormat = "JSON"
	AIGatewaySummaryModelParamsFormatCsv  AIGatewaySummaryModelParamsFormat = "CSV"
)

func (r AIGatewaySummaryModelParamsFormat) IsKnown() bool {
	switch r {
	case AIGatewaySummaryModelParamsFormatJson, AIGatewaySummaryModelParamsFormatCsv:
		return true
	}
	return false
}

type AIGatewaySummaryModelResponseEnvelope struct {
	Result  AIGatewaySummaryModelResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    aiGatewaySummaryModelResponseEnvelopeJSON `json:"-"`
}

// aiGatewaySummaryModelResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIGatewaySummaryModelResponseEnvelope]
type aiGatewaySummaryModelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryModelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryProviderParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AIGatewaySummaryProviderParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIGatewaySummaryProviderParams]'s query parameters as
// `url.Values`.
func (r AIGatewaySummaryProviderParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AIGatewaySummaryProviderParamsFormat string

const (
	AIGatewaySummaryProviderParamsFormatJson AIGatewaySummaryProviderParamsFormat = "JSON"
	AIGatewaySummaryProviderParamsFormatCsv  AIGatewaySummaryProviderParamsFormat = "CSV"
)

func (r AIGatewaySummaryProviderParamsFormat) IsKnown() bool {
	switch r {
	case AIGatewaySummaryProviderParamsFormatJson, AIGatewaySummaryProviderParamsFormatCsv:
		return true
	}
	return false
}

type AIGatewaySummaryProviderResponseEnvelope struct {
	Result  AIGatewaySummaryProviderResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    aiGatewaySummaryProviderResponseEnvelopeJSON `json:"-"`
}

// aiGatewaySummaryProviderResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIGatewaySummaryProviderResponseEnvelope]
type aiGatewaySummaryProviderResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryProviderResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryProviderResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewaySummaryTaskParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AIGatewaySummaryTaskParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIGatewaySummaryTaskParams]'s query parameters as
// `url.Values`.
func (r AIGatewaySummaryTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AIGatewaySummaryTaskParamsFormat string

const (
	AIGatewaySummaryTaskParamsFormatJson AIGatewaySummaryTaskParamsFormat = "JSON"
	AIGatewaySummaryTaskParamsFormatCsv  AIGatewaySummaryTaskParamsFormat = "CSV"
)

func (r AIGatewaySummaryTaskParamsFormat) IsKnown() bool {
	switch r {
	case AIGatewaySummaryTaskParamsFormatJson, AIGatewaySummaryTaskParamsFormatCsv:
		return true
	}
	return false
}

type AIGatewaySummaryTaskResponseEnvelope struct {
	Result  AIGatewaySummaryTaskResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    aiGatewaySummaryTaskResponseEnvelopeJSON `json:"-"`
}

// aiGatewaySummaryTaskResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIGatewaySummaryTaskResponseEnvelope]
type aiGatewaySummaryTaskResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewaySummaryTaskResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewaySummaryTaskResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
