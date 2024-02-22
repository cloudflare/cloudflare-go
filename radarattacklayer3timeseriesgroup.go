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

// RadarAttackLayer3TimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TimeseriesGroupService] method instead.
type RadarAttackLayer3TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupService) {
	r = &RadarAttackLayer3TimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate over time.
func (r *RadarAttackLayer3TimeseriesGroupService) Bitrate(ctx context.Context, query RadarAttackLayer3TimeseriesGroupBitrateParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupBitrateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupBitrateResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by duration over time.
func (r *RadarAttackLayer3TimeseriesGroupService) Duration(ctx context.Context, query RadarAttackLayer3TimeseriesGroupDurationParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupDurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupDurationResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a timeseries of the percentage distribution of network protocols in Layer
// 3/4 attacks.
func (r *RadarAttackLayer3TimeseriesGroupService) Get(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupGetResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by industry used over time.
func (r *RadarAttackLayer3TimeseriesGroupService) Industry(ctx context.Context, query RadarAttackLayer3TimeseriesGroupIndustryParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupIndustryResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used over time.
func (r *RadarAttackLayer3TimeseriesGroupService) IPVersion(ctx context.Context, query RadarAttackLayer3TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupIPVersionResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by protocol used over time.
func (r *RadarAttackLayer3TimeseriesGroupService) Protocol(ctx context.Context, query RadarAttackLayer3TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupProtocolResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vector used over time.
func (r *RadarAttackLayer3TimeseriesGroupService) Vector(ctx context.Context, query RadarAttackLayer3TimeseriesGroupVectorParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupVectorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupVectorResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vertical used over time.
func (r *RadarAttackLayer3TimeseriesGroupService) Vertical(ctx context.Context, query RadarAttackLayer3TimeseriesGroupVerticalParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupVerticalResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TimeseriesGroupBitrateResponse struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupBitrateResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupBitrateResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupBitrateResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupBitrateResponse]
type radarAttackLayer3TimeseriesGroupBitrateResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupBitrateResponseSerie0 struct {
	Number1GbpsTo10Gbps   []string                                                  `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps []string                                                  `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  []string                                                  `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           []string                                                  `json:"OVER_100_GBPS,required"`
	Timestamps            []string                                                  `json:"timestamps,required"`
	Under500Mbps          []string                                                  `json:"UNDER_500_MBPS,required"`
	JSON                  radarAttackLayer3TimeseriesGroupBitrateResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupBitrateResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupBitrateResponseSerie0]
type radarAttackLayer3TimeseriesGroupBitrateResponseSerie0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Timestamps            apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupBitrateResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupDurationResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupDurationResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupDurationResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupDurationResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupDurationResponse]
type radarAttackLayer3TimeseriesGroupDurationResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupDurationResponseSerie0 struct {
	Number1HourTo3Hours  []string                                                   `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins []string                                                   `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins []string                                                   `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  []string                                                   `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           []string                                                   `json:"OVER_3_HOURS,required"`
	Timestamps           []string                                                   `json:"timestamps,required"`
	Under10Mins          []string                                                   `json:"UNDER_10_MINS,required"`
	JSON                 radarAttackLayer3TimeseriesGroupDurationResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupDurationResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupDurationResponseSerie0]
type radarAttackLayer3TimeseriesGroupDurationResponseSerie0JSON struct {
	Number1HourTo3Hours  apijson.Field
	Number10MinsTo20Mins apijson.Field
	Number20MinsTo40Mins apijson.Field
	Number40MinsTo1Hour  apijson.Field
	Over3Hours           apijson.Field
	Timestamps           apijson.Field
	Under10Mins          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupDurationResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetResponse struct {
	Meta   RadarAttackLayer3TimeseriesGroupGetResponseMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TimeseriesGroupGetResponse]
type radarAttackLayer3TimeseriesGroupGetResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetResponseMeta struct {
	AggInterval    string                                                        `json:"aggInterval,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                     `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TimeseriesGroupGetResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupGetResponseMeta]
type radarAttackLayer3TimeseriesGroupGetResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetResponseMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetResponseSerie0 struct {
	GRE        []string                                              `json:"gre,required"`
	Icmp       []string                                              `json:"icmp,required"`
	Tcp        []string                                              `json:"tcp,required"`
	Timestamps []string                                              `json:"timestamps,required"`
	Udp        []string                                              `json:"udp,required"`
	JSON       radarAttackLayer3TimeseriesGroupGetResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupGetResponseSerie0]
type radarAttackLayer3TimeseriesGroupGetResponseSerie0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIndustryResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupIndustryResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupIndustryResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIndustryResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupIndustryResponse]
type radarAttackLayer3TimeseriesGroupIndustryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIndustryResponseSerie0 struct {
	Timestamps  []string                                                   `json:"timestamps,required"`
	ExtraFields map[string][]string                                        `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupIndustryResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIndustryResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupIndustryResponseSerie0]
type radarAttackLayer3TimeseriesGroupIndustryResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIndustryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIPVersionResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupIPVersionResponse]
type radarAttackLayer3TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                                    `json:"IPv4,required"`
	IPv6       []string                                                    `json:"IPv6,required"`
	Timestamps []string                                                    `json:"timestamps,required"`
	JSON       radarAttackLayer3TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIPVersionResponseSerie0]
type radarAttackLayer3TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupProtocolResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupProtocolResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupProtocolResponse]
type radarAttackLayer3TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupProtocolResponseSerie0 struct {
	GRE        []string                                                   `json:"GRE,required"`
	Icmp       []string                                                   `json:"ICMP,required"`
	Tcp        []string                                                   `json:"TCP,required"`
	Timestamps []string                                                   `json:"timestamps,required"`
	Udp        []string                                                   `json:"UDP,required"`
	JSON       radarAttackLayer3TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupProtocolResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupProtocolResponseSerie0]
type radarAttackLayer3TimeseriesGroupProtocolResponseSerie0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVectorResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupVectorResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupVectorResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVectorResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupVectorResponse]
type radarAttackLayer3TimeseriesGroupVectorResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVectorResponseSerie0 struct {
	Timestamps  []string                                                 `json:"timestamps,required"`
	ExtraFields map[string][]string                                      `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupVectorResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVectorResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupVectorResponseSerie0]
type radarAttackLayer3TimeseriesGroupVectorResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVectorResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVerticalResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupVerticalResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupVerticalResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVerticalResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseriesGroupVerticalResponse]
type radarAttackLayer3TimeseriesGroupVerticalResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVerticalResponseSerie0 struct {
	Timestamps  []string                                                   `json:"timestamps,required"`
	ExtraFields map[string][]string                                        `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupVerticalResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVerticalResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupVerticalResponseSerie0]
type radarAttackLayer3TimeseriesGroupVerticalResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVerticalResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupBitrateParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupBitrateParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupBitrateParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupBitrateParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupBitrateParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupBitrateParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval15m RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange1d         RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange2d         RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange7d         RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange14d        RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange28d        RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange12w        RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange24w        RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange52w        RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupBitrateParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupBitrateParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupBitrateParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupBitrateParamsDirectionTarget RadarAttackLayer3TimeseriesGroupBitrateParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupBitrateParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsFormatJson RadarAttackLayer3TimeseriesGroupBitrateParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupBitrateParamsFormatCsv  RadarAttackLayer3TimeseriesGroupBitrateParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupBitrateParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupBitrateParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupBitrateParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupBitrateParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupBitrateParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupBitrateParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupBitrateParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupBitrateParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupBitrateParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupBitrateParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupBitrateParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupBitrateParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupBitrateParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupBitrateParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupBitrateParamsProtocolGRE  RadarAttackLayer3TimeseriesGroupBitrateParamsProtocol = "GRE"
)

type RadarAttackLayer3TimeseriesGroupBitrateResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupBitrateResponse             `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupBitrateResponseEnvelope]
type radarAttackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupDurationParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupDurationParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupDurationParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupDurationParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupDurationParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupDurationParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupDurationParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval15m RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupDurationParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupDurationParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange1d         RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange2d         RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange7d         RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange14d        RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange28d        RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange12w        RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange24w        RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange52w        RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupDurationParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupDurationParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupDurationParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupDurationParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupDurationParamsDirectionTarget RadarAttackLayer3TimeseriesGroupDurationParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupDurationParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsFormatJson RadarAttackLayer3TimeseriesGroupDurationParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupDurationParamsFormatCsv  RadarAttackLayer3TimeseriesGroupDurationParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupDurationParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupDurationParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupDurationParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupDurationParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupDurationParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupDurationParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupDurationParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupDurationParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupDurationParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupDurationParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupDurationParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupDurationParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupDurationParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupDurationParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupDurationParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupDurationParamsProtocolGRE  RadarAttackLayer3TimeseriesGroupDurationParamsProtocol = "GRE"
)

type RadarAttackLayer3TimeseriesGroupDurationResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupDurationResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupDurationResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupDurationResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupDurationResponseEnvelope]
type radarAttackLayer3TimeseriesGroupDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupGetParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupGetParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange1d         RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange2d         RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange7d         RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange14d        RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange28d        RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange12w        RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange24w        RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange52w        RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupGetParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupGetParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetParamsFormatJson RadarAttackLayer3TimeseriesGroupGetParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupGetResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupGetResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupGetResponseEnvelope]
type radarAttackLayer3TimeseriesGroupGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIndustryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupIndustryParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupIndustryParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupIndustryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval15m RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange1d         RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange2d         RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange7d         RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange14d        RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange28d        RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange12w        RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange24w        RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange52w        RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupIndustryParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupIndustryParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupIndustryParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupIndustryParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupIndustryParamsDirectionTarget RadarAttackLayer3TimeseriesGroupIndustryParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupIndustryParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupIndustryParamsFormatJson RadarAttackLayer3TimeseriesGroupIndustryParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupIndustryParamsFormatCsv  RadarAttackLayer3TimeseriesGroupIndustryParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupIndustryParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupIndustryParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupIndustryParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupIndustryParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupIndustryParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupIndustryParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupIndustryResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupIndustryResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIndustryResponseEnvelope]
type radarAttackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupIPVersionParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval15m RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange1d         RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange2d         RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange7d         RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange14d        RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange28d        RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange12w        RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange24w        RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange52w        RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupIPVersionParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupIPVersionParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsDirectionTarget RadarAttackLayer3TimeseriesGroupIPVersionParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupIPVersionParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionParamsFormatJson RadarAttackLayer3TimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsFormatCsv  RadarAttackLayer3TimeseriesGroupIPVersionParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupIPVersionParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocolGRE  RadarAttackLayer3TimeseriesGroupIPVersionParamsProtocol = "GRE"
)

type RadarAttackLayer3TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIPVersionResponseEnvelope]
type radarAttackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupProtocolParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupProtocolParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval15m RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange1d         RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange2d         RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange7d         RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange14d        RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange28d        RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange12w        RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange24w        RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange52w        RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupProtocolParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupProtocolParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupProtocolParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupProtocolParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupProtocolParamsDirectionTarget RadarAttackLayer3TimeseriesGroupProtocolParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupProtocolParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupProtocolParamsFormatJson RadarAttackLayer3TimeseriesGroupProtocolParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupProtocolParamsFormatCsv  RadarAttackLayer3TimeseriesGroupProtocolParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupProtocolParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupProtocolParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupProtocolParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupProtocolParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupProtocolParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupProtocolParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupProtocolParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupProtocolParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupProtocolResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupProtocolResponseEnvelope]
type radarAttackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVectorParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupVectorParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupVectorParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupVectorParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupVectorParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupVectorParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupVectorParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval15m RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupVectorParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupVectorParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange1d         RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange2d         RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange7d         RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange14d        RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange28d        RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange12w        RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange24w        RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange52w        RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupVectorParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupVectorParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupVectorParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupVectorParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupVectorParamsDirectionTarget RadarAttackLayer3TimeseriesGroupVectorParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupVectorParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsFormatJson RadarAttackLayer3TimeseriesGroupVectorParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupVectorParamsFormatCsv  RadarAttackLayer3TimeseriesGroupVectorParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupVectorParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupVectorParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupVectorParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupVectorParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupVectorParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupVectorParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupVectorParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupVectorParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupVectorParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupVectorParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupVectorParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupVectorParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupVectorParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupVectorParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupVectorParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupVectorParamsProtocolGRE  RadarAttackLayer3TimeseriesGroupVectorParamsProtocol = "GRE"
)

type RadarAttackLayer3TimeseriesGroupVectorResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupVectorResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupVectorResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVectorResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupVectorResponseEnvelope]
type radarAttackLayer3TimeseriesGroupVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVerticalParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupVerticalParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupVerticalParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupVerticalParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval15m RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange1d         RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange2d         RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange7d         RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange14d        RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange28d        RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange12w        RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange24w        RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange52w        RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupVerticalParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupVerticalParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupVerticalParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupVerticalParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupVerticalParamsDirectionTarget RadarAttackLayer3TimeseriesGroupVerticalParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupVerticalParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupVerticalParamsFormatJson RadarAttackLayer3TimeseriesGroupVerticalParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupVerticalParamsFormatCsv  RadarAttackLayer3TimeseriesGroupVerticalParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupVerticalParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupVerticalParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupVerticalParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupVerticalParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupVerticalParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupVerticalParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupVerticalResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupVerticalResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupVerticalResponseEnvelope]
type radarAttackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
