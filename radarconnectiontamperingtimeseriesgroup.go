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

// RadarConnectionTamperingTimeseriesGroupService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarConnectionTamperingTimeseriesGroupService] method instead.
type RadarConnectionTamperingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarConnectionTamperingTimeseriesGroupService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarConnectionTamperingTimeseriesGroupService(opts ...option.RequestOption) (r *RadarConnectionTamperingTimeseriesGroupService) {
	r = &RadarConnectionTamperingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Distribution of connection tampering types over time.
func (r *RadarConnectionTamperingTimeseriesGroupService) List(ctx context.Context, query RadarConnectionTamperingTimeseriesGroupListParams, opts ...option.RequestOption) (res *RadarConnectionTamperingTimeseriesGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/connection_tampering/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarConnectionTamperingTimeseriesGroupListResponse struct {
	Result  RadarConnectionTamperingTimeseriesGroupListResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarConnectionTamperingTimeseriesGroupListResponseJSON   `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseJSON contains the JSON
// metadata for the struct [RadarConnectionTamperingTimeseriesGroupListResponse]
type radarConnectionTamperingTimeseriesGroupListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingTimeseriesGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListResponseResult struct {
	Meta   RadarConnectionTamperingTimeseriesGroupListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarConnectionTamperingTimeseriesGroupListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarConnectionTamperingTimeseriesGroupListResponseResultJSON   `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarConnectionTamperingTimeseriesGroupListResponseResult]
type radarConnectionTamperingTimeseriesGroupListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingTimeseriesGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListResponseResultMeta struct {
	AggInterval    string                                                                      `json:"aggInterval,required"`
	DateRange      []RadarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                                   `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarConnectionTamperingTimeseriesGroupListResponseResultMetaJSON           `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarConnectionTamperingTimeseriesGroupListResponseResultMeta]
type radarConnectionTamperingTimeseriesGroupListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarConnectionTamperingTimeseriesGroupListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRange]
type radarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingTimeseriesGroupListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                   `json:"level"`
	JSON        radarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfo]
type radarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                    `json:"dataSource,required"`
	Description     string                                                                                    `json:"description,required"`
	EventType       string                                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                                 `json:"startTime" format:"date-time"`
	JSON            radarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation]
type radarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarConnectionTamperingTimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListResponseResultSerie0 struct {
	// Connections matching signatures for tampering later in the connection, after the
	// transfer of multiple data packets.
	LaterInFlow []string `json:"later_in_flow,required"`
	// Connections that do not match any tampering signatures.
	NoMatch []string `json:"no_match,required"`
	// Connections matching signatures for tampering after the receipt of a SYN packet
	// and ACK packet, meaning the TCP connection was successfully established but the
	// server did not receive any subsequent packets.
	PostAck []string `json:"post_ack,required"`
	// Connections matching signatures for tampering after the receipt of a packet with
	// PSH flag set, following connection establishment.
	PostPsh []string `json:"post_psh,required"`
	// Connections matching signatures for tampering after the receipt of only a single
	// SYN packet, and before the handshake completes.
	PostSyn    []string                                                            `json:"post_syn,required"`
	Timestamps []time.Time                                                         `json:"timestamps,required" format:"date-time"`
	JSON       radarConnectionTamperingTimeseriesGroupListResponseResultSerie0JSON `json:"-"`
}

// radarConnectionTamperingTimeseriesGroupListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarConnectionTamperingTimeseriesGroupListResponseResultSerie0]
type radarConnectionTamperingTimeseriesGroupListResponseResultSerie0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingTimeseriesGroupListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingTimeseriesGroupListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarConnectionTamperingTimeseriesGroupListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarConnectionTamperingTimeseriesGroupListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarConnectionTamperingTimeseriesGroupListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarConnectionTamperingTimeseriesGroupListParams]'s query
// parameters as `url.Values`.
func (r RadarConnectionTamperingTimeseriesGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarConnectionTamperingTimeseriesGroupListParamsAggInterval string

const (
	RadarConnectionTamperingTimeseriesGroupListParamsAggInterval15m RadarConnectionTamperingTimeseriesGroupListParamsAggInterval = "15m"
	RadarConnectionTamperingTimeseriesGroupListParamsAggInterval1h  RadarConnectionTamperingTimeseriesGroupListParamsAggInterval = "1h"
	RadarConnectionTamperingTimeseriesGroupListParamsAggInterval1d  RadarConnectionTamperingTimeseriesGroupListParamsAggInterval = "1d"
	RadarConnectionTamperingTimeseriesGroupListParamsAggInterval1w  RadarConnectionTamperingTimeseriesGroupListParamsAggInterval = "1w"
)

type RadarConnectionTamperingTimeseriesGroupListParamsDateRange string

const (
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange1d         RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "1d"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange2d         RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "2d"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange7d         RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "7d"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange14d        RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "14d"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange28d        RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "28d"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange12w        RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "12w"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange24w        RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "24w"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange52w        RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "52w"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange1dControl  RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "1dControl"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange2dControl  RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "2dControl"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange7dControl  RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "7dControl"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange14dControl RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "14dControl"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange28dControl RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "28dControl"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange12wControl RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "12wControl"
	RadarConnectionTamperingTimeseriesGroupListParamsDateRange24wControl RadarConnectionTamperingTimeseriesGroupListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarConnectionTamperingTimeseriesGroupListParamsFormat string

const (
	RadarConnectionTamperingTimeseriesGroupListParamsFormatJson RadarConnectionTamperingTimeseriesGroupListParamsFormat = "JSON"
	RadarConnectionTamperingTimeseriesGroupListParamsFormatCsv  RadarConnectionTamperingTimeseriesGroupListParamsFormat = "CSV"
)
