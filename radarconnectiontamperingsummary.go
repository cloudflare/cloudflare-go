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

// RadarConnectionTamperingSummaryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarConnectionTamperingSummaryService] method instead.
type RadarConnectionTamperingSummaryService struct {
	Options []option.RequestOption
}

// NewRadarConnectionTamperingSummaryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarConnectionTamperingSummaryService(opts ...option.RequestOption) (r *RadarConnectionTamperingSummaryService) {
	r = &RadarConnectionTamperingSummaryService{}
	r.Options = opts
	return
}

// Distribution of connection tampering types over a given time period.
func (r *RadarConnectionTamperingSummaryService) List(ctx context.Context, query RadarConnectionTamperingSummaryListParams, opts ...option.RequestOption) (res *RadarConnectionTamperingSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/connection_tampering/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarConnectionTamperingSummaryListResponse struct {
	Result  RadarConnectionTamperingSummaryListResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarConnectionTamperingSummaryListResponseJSON   `json:"-"`
}

// radarConnectionTamperingSummaryListResponseJSON contains the JSON metadata for
// the struct [RadarConnectionTamperingSummaryListResponse]
type radarConnectionTamperingSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListResponseResult struct {
	Meta     RadarConnectionTamperingSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarConnectionTamperingSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarConnectionTamperingSummaryListResponseResultJSON     `json:"-"`
}

// radarConnectionTamperingSummaryListResponseResultJSON contains the JSON metadata
// for the struct [RadarConnectionTamperingSummaryListResponseResult]
type radarConnectionTamperingSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListResponseResultMeta struct {
	DateRange      []RadarConnectionTamperingSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarConnectionTamperingSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarConnectionTamperingSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarConnectionTamperingSummaryListResponseResultMeta]
type radarConnectionTamperingSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarConnectionTamperingSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarConnectionTamperingSummaryListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarConnectionTamperingSummaryListResponseResultMetaDateRange]
type radarConnectionTamperingSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfo]
type radarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarConnectionTamperingSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListResponseResultSummary0 struct {
	// Connections matching signatures for tampering later in the connection, after the
	// transfer of multiple data packets.
	LaterInFlow string `json:"later_in_flow,required"`
	// Connections that do not match any tampering signatures.
	NoMatch string `json:"no_match,required"`
	// Connections matching signatures for tampering after the receipt of a SYN packet
	// and ACK packet, meaning the TCP connection was successfully established but the
	// server did not receive any subsequent packets.
	PostAck string `json:"post_ack,required"`
	// Connections matching signatures for tampering after the receipt of a packet with
	// PSH flag set, following connection establishment.
	PostPsh string `json:"post_psh,required"`
	// Connections matching signatures for tampering after the receipt of only a single
	// SYN packet, and before the handshake completes.
	PostSyn string                                                        `json:"post_syn,required"`
	JSON    radarConnectionTamperingSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarConnectionTamperingSummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarConnectionTamperingSummaryListResponseResultSummary0]
type radarConnectionTamperingSummaryListResponseResultSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarConnectionTamperingSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarConnectionTamperingSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarConnectionTamperingSummaryListParams]'s query
// parameters as `url.Values`.
func (r RadarConnectionTamperingSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarConnectionTamperingSummaryListParamsDateRange string

const (
	RadarConnectionTamperingSummaryListParamsDateRange1d         RadarConnectionTamperingSummaryListParamsDateRange = "1d"
	RadarConnectionTamperingSummaryListParamsDateRange2d         RadarConnectionTamperingSummaryListParamsDateRange = "2d"
	RadarConnectionTamperingSummaryListParamsDateRange7d         RadarConnectionTamperingSummaryListParamsDateRange = "7d"
	RadarConnectionTamperingSummaryListParamsDateRange14d        RadarConnectionTamperingSummaryListParamsDateRange = "14d"
	RadarConnectionTamperingSummaryListParamsDateRange28d        RadarConnectionTamperingSummaryListParamsDateRange = "28d"
	RadarConnectionTamperingSummaryListParamsDateRange12w        RadarConnectionTamperingSummaryListParamsDateRange = "12w"
	RadarConnectionTamperingSummaryListParamsDateRange24w        RadarConnectionTamperingSummaryListParamsDateRange = "24w"
	RadarConnectionTamperingSummaryListParamsDateRange52w        RadarConnectionTamperingSummaryListParamsDateRange = "52w"
	RadarConnectionTamperingSummaryListParamsDateRange1dControl  RadarConnectionTamperingSummaryListParamsDateRange = "1dControl"
	RadarConnectionTamperingSummaryListParamsDateRange2dControl  RadarConnectionTamperingSummaryListParamsDateRange = "2dControl"
	RadarConnectionTamperingSummaryListParamsDateRange7dControl  RadarConnectionTamperingSummaryListParamsDateRange = "7dControl"
	RadarConnectionTamperingSummaryListParamsDateRange14dControl RadarConnectionTamperingSummaryListParamsDateRange = "14dControl"
	RadarConnectionTamperingSummaryListParamsDateRange28dControl RadarConnectionTamperingSummaryListParamsDateRange = "28dControl"
	RadarConnectionTamperingSummaryListParamsDateRange12wControl RadarConnectionTamperingSummaryListParamsDateRange = "12wControl"
	RadarConnectionTamperingSummaryListParamsDateRange24wControl RadarConnectionTamperingSummaryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarConnectionTamperingSummaryListParamsFormat string

const (
	RadarConnectionTamperingSummaryListParamsFormatJson RadarConnectionTamperingSummaryListParamsFormat = "JSON"
	RadarConnectionTamperingSummaryListParamsFormatCsv  RadarConnectionTamperingSummaryListParamsFormat = "CSV"
)
