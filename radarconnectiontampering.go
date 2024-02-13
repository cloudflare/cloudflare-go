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
func (r *RadarConnectionTamperingService) Summary(ctx context.Context, query RadarConnectionTamperingSummaryParams, opts ...option.RequestOption) (res *RadarConnectionTamperingSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarConnectionTamperingSummaryResponseEnvelope
	path := "radar/connection_tampering/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarConnectionTamperingSummaryResponse struct {
	Meta     RadarConnectionTamperingSummaryResponseMeta     `json:"meta,required"`
	Summary0 RadarConnectionTamperingSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     radarConnectionTamperingSummaryResponseJSON     `json:"-"`
}

// radarConnectionTamperingSummaryResponseJSON contains the JSON metadata for the
// struct [RadarConnectionTamperingSummaryResponse]
type radarConnectionTamperingSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryResponseMeta struct {
	DateRange      []RadarConnectionTamperingSummaryResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarConnectionTamperingSummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarConnectionTamperingSummaryResponseMetaJSON           `json:"-"`
}

// radarConnectionTamperingSummaryResponseMetaJSON contains the JSON metadata for
// the struct [RadarConnectionTamperingSummaryResponseMeta]
type radarConnectionTamperingSummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarConnectionTamperingSummaryResponseMetaDateRangeJSON `json:"-"`
}

// radarConnectionTamperingSummaryResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarConnectionTamperingSummaryResponseMetaDateRange]
type radarConnectionTamperingSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryResponseMetaConfidenceInfo struct {
	Annotations []RadarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarConnectionTamperingSummaryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarConnectionTamperingSummaryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarConnectionTamperingSummaryResponseMetaConfidenceInfo]
type radarConnectionTamperingSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotation]
type radarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarConnectionTamperingSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryResponseSummary0 struct {
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
	PostSyn string                                              `json:"post_syn,required"`
	JSON    radarConnectionTamperingSummaryResponseSummary0JSON `json:"-"`
}

// radarConnectionTamperingSummaryResponseSummary0JSON contains the JSON metadata
// for the struct [RadarConnectionTamperingSummaryResponseSummary0]
type radarConnectionTamperingSummaryResponseSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarConnectionTamperingSummaryParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarConnectionTamperingSummaryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarConnectionTamperingSummaryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarConnectionTamperingSummaryParams]'s query parameters
// as `url.Values`.
func (r RadarConnectionTamperingSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarConnectionTamperingSummaryParamsDateRange string

const (
	RadarConnectionTamperingSummaryParamsDateRange1d         RadarConnectionTamperingSummaryParamsDateRange = "1d"
	RadarConnectionTamperingSummaryParamsDateRange2d         RadarConnectionTamperingSummaryParamsDateRange = "2d"
	RadarConnectionTamperingSummaryParamsDateRange7d         RadarConnectionTamperingSummaryParamsDateRange = "7d"
	RadarConnectionTamperingSummaryParamsDateRange14d        RadarConnectionTamperingSummaryParamsDateRange = "14d"
	RadarConnectionTamperingSummaryParamsDateRange28d        RadarConnectionTamperingSummaryParamsDateRange = "28d"
	RadarConnectionTamperingSummaryParamsDateRange12w        RadarConnectionTamperingSummaryParamsDateRange = "12w"
	RadarConnectionTamperingSummaryParamsDateRange24w        RadarConnectionTamperingSummaryParamsDateRange = "24w"
	RadarConnectionTamperingSummaryParamsDateRange52w        RadarConnectionTamperingSummaryParamsDateRange = "52w"
	RadarConnectionTamperingSummaryParamsDateRange1dControl  RadarConnectionTamperingSummaryParamsDateRange = "1dControl"
	RadarConnectionTamperingSummaryParamsDateRange2dControl  RadarConnectionTamperingSummaryParamsDateRange = "2dControl"
	RadarConnectionTamperingSummaryParamsDateRange7dControl  RadarConnectionTamperingSummaryParamsDateRange = "7dControl"
	RadarConnectionTamperingSummaryParamsDateRange14dControl RadarConnectionTamperingSummaryParamsDateRange = "14dControl"
	RadarConnectionTamperingSummaryParamsDateRange28dControl RadarConnectionTamperingSummaryParamsDateRange = "28dControl"
	RadarConnectionTamperingSummaryParamsDateRange12wControl RadarConnectionTamperingSummaryParamsDateRange = "12wControl"
	RadarConnectionTamperingSummaryParamsDateRange24wControl RadarConnectionTamperingSummaryParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarConnectionTamperingSummaryParamsFormat string

const (
	RadarConnectionTamperingSummaryParamsFormatJson RadarConnectionTamperingSummaryParamsFormat = "JSON"
	RadarConnectionTamperingSummaryParamsFormatCsv  RadarConnectionTamperingSummaryParamsFormat = "CSV"
)

type RadarConnectionTamperingSummaryResponseEnvelope struct {
	Result  RadarConnectionTamperingSummaryResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarConnectionTamperingSummaryResponseEnvelopeJSON `json:"-"`
}

// radarConnectionTamperingSummaryResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarConnectionTamperingSummaryResponseEnvelope]
type radarConnectionTamperingSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarConnectionTamperingSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
