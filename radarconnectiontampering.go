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

// RadarConnectionTamperingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarConnectionTamperingService] method instead.
type RadarConnectionTamperingService struct {
	Options          []option.RequestOption
	TimeseriesGroups *RadarConnectionTamperingTimeseriesGroupService
}

// NewRadarConnectionTamperingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarConnectionTamperingService(opts ...option.RequestOption) (r *RadarConnectionTamperingService) {
	r = &RadarConnectionTamperingService{}
	r.Options = opts
	r.TimeseriesGroups = NewRadarConnectionTamperingTimeseriesGroupService(opts...)
	return
}

// Distribution of connection tampering types over a given time period.
func (r *RadarConnectionTamperingService) List(ctx context.Context, query RadarConnectionTamperingListParams, opts ...option.RequestOption) (res *RadarConnectionTamperingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarConnectionTamperingListResponseEnvelope
	path := "radar/connection_tampering/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarConnectionTamperingListResponse struct {
	Meta     RadarConnectionTamperingListResponseMeta     `json:"meta,required"`
	Summary0 RadarConnectionTamperingListResponseSummary0 `json:"summary_0,required"`
	JSON     radarConnectionTamperingListResponseJSON     `json:"-"`
}

// radarConnectionTamperingListResponseJSON contains the JSON metadata for the
// struct [RadarConnectionTamperingListResponse]
type radarConnectionTamperingListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingListResponseMeta struct {
	DateRange      []RadarConnectionTamperingListResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarConnectionTamperingListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarConnectionTamperingListResponseMetaJSON           `json:"-"`
}

// radarConnectionTamperingListResponseMetaJSON contains the JSON metadata for the
// struct [RadarConnectionTamperingListResponseMeta]
type radarConnectionTamperingListResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarConnectionTamperingListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarConnectionTamperingListResponseMetaDateRangeJSON `json:"-"`
}

// radarConnectionTamperingListResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarConnectionTamperingListResponseMetaDateRange]
type radarConnectionTamperingListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingListResponseMetaConfidenceInfo struct {
	Annotations []RadarConnectionTamperingListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarConnectionTamperingListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarConnectionTamperingListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarConnectionTamperingListResponseMetaConfidenceInfo]
type radarConnectionTamperingListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarConnectionTamperingListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarConnectionTamperingListResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarConnectionTamperingListResponseMetaConfidenceInfoAnnotation]
type radarConnectionTamperingListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarConnectionTamperingListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingListResponseSummary0 struct {
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
	PostSyn string                                           `json:"post_syn,required"`
	JSON    radarConnectionTamperingListResponseSummary0JSON `json:"-"`
}

// radarConnectionTamperingListResponseSummary0JSON contains the JSON metadata for
// the struct [RadarConnectionTamperingListResponseSummary0]
type radarConnectionTamperingListResponseSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarConnectionTamperingListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarConnectionTamperingListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarConnectionTamperingListParams]'s query parameters as
// `url.Values`.
func (r RadarConnectionTamperingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarConnectionTamperingListParamsDateRange string

const (
	RadarConnectionTamperingListParamsDateRange1d         RadarConnectionTamperingListParamsDateRange = "1d"
	RadarConnectionTamperingListParamsDateRange2d         RadarConnectionTamperingListParamsDateRange = "2d"
	RadarConnectionTamperingListParamsDateRange7d         RadarConnectionTamperingListParamsDateRange = "7d"
	RadarConnectionTamperingListParamsDateRange14d        RadarConnectionTamperingListParamsDateRange = "14d"
	RadarConnectionTamperingListParamsDateRange28d        RadarConnectionTamperingListParamsDateRange = "28d"
	RadarConnectionTamperingListParamsDateRange12w        RadarConnectionTamperingListParamsDateRange = "12w"
	RadarConnectionTamperingListParamsDateRange24w        RadarConnectionTamperingListParamsDateRange = "24w"
	RadarConnectionTamperingListParamsDateRange52w        RadarConnectionTamperingListParamsDateRange = "52w"
	RadarConnectionTamperingListParamsDateRange1dControl  RadarConnectionTamperingListParamsDateRange = "1dControl"
	RadarConnectionTamperingListParamsDateRange2dControl  RadarConnectionTamperingListParamsDateRange = "2dControl"
	RadarConnectionTamperingListParamsDateRange7dControl  RadarConnectionTamperingListParamsDateRange = "7dControl"
	RadarConnectionTamperingListParamsDateRange14dControl RadarConnectionTamperingListParamsDateRange = "14dControl"
	RadarConnectionTamperingListParamsDateRange28dControl RadarConnectionTamperingListParamsDateRange = "28dControl"
	RadarConnectionTamperingListParamsDateRange12wControl RadarConnectionTamperingListParamsDateRange = "12wControl"
	RadarConnectionTamperingListParamsDateRange24wControl RadarConnectionTamperingListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarConnectionTamperingListParamsFormat string

const (
	RadarConnectionTamperingListParamsFormatJson RadarConnectionTamperingListParamsFormat = "JSON"
	RadarConnectionTamperingListParamsFormatCsv  RadarConnectionTamperingListParamsFormat = "CSV"
)

type RadarConnectionTamperingListResponseEnvelope struct {
	Result  RadarConnectionTamperingListResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarConnectionTamperingListResponseEnvelopeJSON `json:"-"`
}

// radarConnectionTamperingListResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarConnectionTamperingListResponseEnvelope]
type radarConnectionTamperingListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
