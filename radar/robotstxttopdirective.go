// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
  "context"
  "fmt"
  "net/http"
  "net/url"
  "time"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/option"
)

// RobotsTXTTopDirectiveService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotsTXTTopDirectiveService] method instead.
type RobotsTXTTopDirectiveService struct {
Options []option.RequestOption
}

// NewRobotsTXTTopDirectiveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRobotsTXTTopDirectiveService(opts ...option.RequestOption) (r *RobotsTXTTopDirectiveService) {
  r = &RobotsTXTTopDirectiveService{}
  r.Options = opts
  return
}

// Get the top User-Agents on robots.txt files by directive.
func (r *RobotsTXTTopDirectiveService) Get(ctx context.Context, directive RobotsTXTTopDirectiveGetParamsDirective, query RobotsTXTTopDirectiveGetParams, opts ...option.RequestOption) (res *RobotsTXTTopDirectiveGetResponse, err error) {
  var env RobotsTXTTopDirectiveGetResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("radar/robots_txt/top/%v", directive)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type RobotsTXTTopDirectiveGetResponse struct {
Meta RobotsTXTTopDirectiveGetResponseMeta `json:"meta,required"`
Top0 []RobotsTXTTopDirectiveGetResponseTop0 `json:"top_0,required"`
JSON robotsTXTTopDirectiveGetResponseJSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseJSON contains the JSON metadata for the struct
// [RobotsTXTTopDirectiveGetResponse]
type robotsTXTTopDirectiveGetResponseJSON struct {
Meta apijson.Field
Top0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTTopDirectiveGetResponseMeta struct {
DateRange []RobotsTXTTopDirectiveGetResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
ConfidenceInfo RobotsTXTTopDirectiveGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON robotsTXTTopDirectiveGetResponseMetaJSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseMetaJSON contains the JSON metadata for the
// struct [RobotsTXTTopDirectiveGetResponseMeta]
type robotsTXTTopDirectiveGetResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTTopDirectiveGetResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON robotsTXTTopDirectiveGetResponseMetaDateRangeJSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RobotsTXTTopDirectiveGetResponseMetaDateRange]
type robotsTXTTopDirectiveGetResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTTopDirectiveGetResponseMetaConfidenceInfo struct {
Annotations []RobotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON robotsTXTTopDirectiveGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RobotsTXTTopDirectiveGetResponseMetaConfidenceInfo]
type robotsTXTTopDirectiveGetResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON robotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RobotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotation]
type robotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTTopDirectiveGetResponseTop0 struct {
Date string `json:"date,required"`
Fully int64 `json:"fully,required"`
Partially int64 `json:"partially,required"`
Total int64 `json:"total,required"`
UserAgent string `json:"userAgent,required"`
JSON robotsTXTTopDirectiveGetResponseTop0JSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseTop0JSON contains the JSON metadata for the
// struct [RobotsTXTTopDirectiveGetResponseTop0]
type robotsTXTTopDirectiveGetResponseTop0JSON struct {
Date apijson.Field
Fully apijson.Field
Partially apijson.Field
Total apijson.Field
UserAgent apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseTop0JSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTTopDirectiveGetParams struct {
// Filter by user agent category.
AgentCategory param.Field[RobotsTXTTopDirectiveGetParamsAgentCategory] `query:"agentCategory"`
// Date to filter the ranking.
Date param.Field[string] `query:"date"`
// Format results are returned in.
Format param.Field[RobotsTXTTopDirectiveGetParamsFormat] `query:"format"`
// Limit the number of objects in the response.
Limit param.Field[int64] `query:"limit"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RobotsTXTTopDirectiveGetParams]'s query parameters as
// `url.Values`.
func (r RobotsTXTTopDirectiveGetParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Robots.txt directive.
type RobotsTXTTopDirectiveGetParamsDirective string

const (
  RobotsTXTTopDirectiveGetParamsDirectiveAllow RobotsTXTTopDirectiveGetParamsDirective = "ALLOW"
  RobotsTXTTopDirectiveGetParamsDirectiveDisallow RobotsTXTTopDirectiveGetParamsDirective = "DISALLOW"
)

func (r RobotsTXTTopDirectiveGetParamsDirective) IsKnown() (bool) {
  switch r {
  case RobotsTXTTopDirectiveGetParamsDirectiveAllow, RobotsTXTTopDirectiveGetParamsDirectiveDisallow:
      return true
  }
  return false
}

// Filter by user agent category.
type RobotsTXTTopDirectiveGetParamsAgentCategory string

const (
  RobotsTXTTopDirectiveGetParamsAgentCategoryAI RobotsTXTTopDirectiveGetParamsAgentCategory = "AI"
)

func (r RobotsTXTTopDirectiveGetParamsAgentCategory) IsKnown() (bool) {
  switch r {
  case RobotsTXTTopDirectiveGetParamsAgentCategoryAI:
      return true
  }
  return false
}

// Format results are returned in.
type RobotsTXTTopDirectiveGetParamsFormat string

const (
  RobotsTXTTopDirectiveGetParamsFormatJson RobotsTXTTopDirectiveGetParamsFormat = "JSON"
  RobotsTXTTopDirectiveGetParamsFormatCsv RobotsTXTTopDirectiveGetParamsFormat = "CSV"
)

func (r RobotsTXTTopDirectiveGetParamsFormat) IsKnown() (bool) {
  switch r {
  case RobotsTXTTopDirectiveGetParamsFormatJson, RobotsTXTTopDirectiveGetParamsFormatCsv:
      return true
  }
  return false
}

type RobotsTXTTopDirectiveGetResponseEnvelope struct {
Result RobotsTXTTopDirectiveGetResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON robotsTXTTopDirectiveGetResponseEnvelopeJSON `json:"-"`
}

// robotsTXTTopDirectiveGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RobotsTXTTopDirectiveGetResponseEnvelope]
type robotsTXTTopDirectiveGetResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDirectiveGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDirectiveGetResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}
