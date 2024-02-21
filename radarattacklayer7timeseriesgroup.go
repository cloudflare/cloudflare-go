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

// RadarAttackLayer7TimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TimeseriesGroupService] method instead.
type RadarAttackLayer7TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer7TimeseriesGroupService) {
	r = &RadarAttackLayer7TimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate over time.
func (r *RadarAttackLayer7TimeseriesGroupService) Bitrate(ctx context.Context, query RadarAttackLayer7TimeseriesGroupBitrateParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupBitrateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupBitrateResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by duration over time.
func (r *RadarAttackLayer7TimeseriesGroupService) Duration(ctx context.Context, query RadarAttackLayer7TimeseriesGroupDurationParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupDurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupDurationResponseEnvelope
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
func (r *RadarAttackLayer7TimeseriesGroupService) Get(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupGetResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by industry used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) Industry(ctx context.Context, query RadarAttackLayer7TimeseriesGroupIndustryParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupIndustryResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) IPVersion(ctx context.Context, query RadarAttackLayer7TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupIPVersionResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by protocol used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) Protocol(ctx context.Context, query RadarAttackLayer7TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupProtocolResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vector used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) Vector(ctx context.Context, query RadarAttackLayer7TimeseriesGroupVectorParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupVectorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupVectorResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vertical used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) Vertical(ctx context.Context, query RadarAttackLayer7TimeseriesGroupVerticalParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupVerticalResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7TimeseriesGroupBitrateResponse struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupBitrateResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupBitrateResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupBitrateResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupBitrateResponse]
type radarAttackLayer7TimeseriesGroupBitrateResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupBitrateResponseSerie0 struct {
	Number1GbpsTo10Gbps   []string                                                  `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps []string                                                  `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  []string                                                  `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           []string                                                  `json:"OVER_100_GBPS,required"`
	Timestamps            []string                                                  `json:"timestamps,required"`
	Under500Mbps          []string                                                  `json:"UNDER_500_MBPS,required"`
	JSON                  radarAttackLayer7TimeseriesGroupBitrateResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupBitrateResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupBitrateResponseSerie0]
type radarAttackLayer7TimeseriesGroupBitrateResponseSerie0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Timestamps            apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupBitrateResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupDurationResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupDurationResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupDurationResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupDurationResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupDurationResponse]
type radarAttackLayer7TimeseriesGroupDurationResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupDurationResponseSerie0 struct {
	Number1HourTo3Hours  []string                                                   `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins []string                                                   `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins []string                                                   `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  []string                                                   `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           []string                                                   `json:"OVER_3_HOURS,required"`
	Timestamps           []string                                                   `json:"timestamps,required"`
	Under10Mins          []string                                                   `json:"UNDER_10_MINS,required"`
	JSON                 radarAttackLayer7TimeseriesGroupDurationResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupDurationResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupDurationResponseSerie0]
type radarAttackLayer7TimeseriesGroupDurationResponseSerie0JSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupDurationResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetResponse struct {
	Meta   RadarAttackLayer7TimeseriesGroupGetResponseMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TimeseriesGroupGetResponse]
type radarAttackLayer7TimeseriesGroupGetResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetResponseMeta struct {
	AggInterval    string                                                        `json:"aggInterval,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                     `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TimeseriesGroupGetResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupGetResponseMeta]
type radarAttackLayer7TimeseriesGroupGetResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetResponseMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetResponseSerie0 struct {
	Gre        []string                                              `json:"gre,required"`
	Icmp       []string                                              `json:"icmp,required"`
	Tcp        []string                                              `json:"tcp,required"`
	Timestamps []string                                              `json:"timestamps,required"`
	Udp        []string                                              `json:"udp,required"`
	JSON       radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupGetResponseSerie0]
type radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupIndustryResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupIndustryResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupIndustryResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupIndustryResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupIndustryResponse]
type radarAttackLayer7TimeseriesGroupIndustryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupIndustryResponseSerie0 struct {
	Timestamps  []string                                                   `json:"timestamps,required"`
	ExtraFields map[string][]string                                        `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupIndustryResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupIndustryResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupIndustryResponseSerie0]
type radarAttackLayer7TimeseriesGroupIndustryResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupIndustryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupIPVersionResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupIPVersionResponse]
type radarAttackLayer7TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                                    `json:"IPv4,required"`
	IPv6       []string                                                    `json:"IPv6,required"`
	Timestamps []string                                                    `json:"timestamps,required"`
	JSON       radarAttackLayer7TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupIPVersionResponseSerie0]
type radarAttackLayer7TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupProtocolResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupProtocolResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupProtocolResponse]
type radarAttackLayer7TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupProtocolResponseSerie0 struct {
	Gre        []string                                                   `json:"GRE,required"`
	Icmp       []string                                                   `json:"ICMP,required"`
	Tcp        []string                                                   `json:"TCP,required"`
	Timestamps []string                                                   `json:"timestamps,required"`
	Udp        []string                                                   `json:"UDP,required"`
	JSON       radarAttackLayer7TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupProtocolResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupProtocolResponseSerie0]
type radarAttackLayer7TimeseriesGroupProtocolResponseSerie0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupVectorResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupVectorResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupVectorResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupVectorResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupVectorResponse]
type radarAttackLayer7TimeseriesGroupVectorResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupVectorResponseSerie0 struct {
	Timestamps  []string                                                 `json:"timestamps,required"`
	ExtraFields map[string][]string                                      `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupVectorResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupVectorResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupVectorResponseSerie0]
type radarAttackLayer7TimeseriesGroupVectorResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupVectorResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupVerticalResponse struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupVerticalResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupVerticalResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupVerticalResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupVerticalResponse]
type radarAttackLayer7TimeseriesGroupVerticalResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupVerticalResponseSerie0 struct {
	Timestamps  []string                                                   `json:"timestamps,required"`
	ExtraFields map[string][]string                                        `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupVerticalResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupVerticalResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupVerticalResponseSerie0]
type radarAttackLayer7TimeseriesGroupVerticalResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupVerticalResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupBitrateParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupBitrateParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupBitrateParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupBitrateParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupBitrateParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval15m RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange1d         RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange2d         RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange7d         RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange14d        RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange28d        RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange12w        RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange24w        RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange52w        RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupBitrateParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupBitrateParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupBitrateParamsDirectionTarget RadarAttackLayer7TimeseriesGroupBitrateParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupBitrateParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsFormatJson RadarAttackLayer7TimeseriesGroupBitrateParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupBitrateParamsFormatCsv  RadarAttackLayer7TimeseriesGroupBitrateParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupBitrateParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupBitrateParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupBitrateParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupBitrateParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol string

const (
	RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolUdp  RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol = "UDP"
	RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolTcp  RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol = "TCP"
	RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolIcmp RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol = "ICMP"
	RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolGre  RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol = "GRE"
)

type RadarAttackLayer7TimeseriesGroupBitrateResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupBitrateResponse             `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupBitrateResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupBitrateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupBitrateResponseEnvelope]
type radarAttackLayer7TimeseriesGroupBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupDurationParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupDurationParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupDurationParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupDurationParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupDurationParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupDurationParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7TimeseriesGroupDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupDurationParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval15m RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupDurationParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange1d         RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange2d         RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange7d         RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange14d        RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange28d        RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange12w        RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange24w        RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange52w        RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupDurationParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupDurationParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupDurationParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupDurationParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupDurationParamsDirectionTarget RadarAttackLayer7TimeseriesGroupDurationParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupDurationParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsFormatJson RadarAttackLayer7TimeseriesGroupDurationParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupDurationParamsFormatCsv  RadarAttackLayer7TimeseriesGroupDurationParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupDurationParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupDurationParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupDurationParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupDurationParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupDurationParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupDurationParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupDurationParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupDurationParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupDurationParamsProtocol string

const (
	RadarAttackLayer7TimeseriesGroupDurationParamsProtocolUdp  RadarAttackLayer7TimeseriesGroupDurationParamsProtocol = "UDP"
	RadarAttackLayer7TimeseriesGroupDurationParamsProtocolTcp  RadarAttackLayer7TimeseriesGroupDurationParamsProtocol = "TCP"
	RadarAttackLayer7TimeseriesGroupDurationParamsProtocolIcmp RadarAttackLayer7TimeseriesGroupDurationParamsProtocol = "ICMP"
	RadarAttackLayer7TimeseriesGroupDurationParamsProtocolGre  RadarAttackLayer7TimeseriesGroupDurationParamsProtocol = "GRE"
)

type RadarAttackLayer7TimeseriesGroupDurationResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupDurationResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupDurationResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupDurationResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupDurationResponseEnvelope]
type radarAttackLayer7TimeseriesGroupDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupGetParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupGetParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange1d         RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange2d         RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange7d         RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange14d        RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange28d        RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange12w        RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange24w        RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange52w        RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupGetParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupGetParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetParamsFormatJson RadarAttackLayer7TimeseriesGroupGetParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupGetResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupGetResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupGetResponseEnvelope]
type radarAttackLayer7TimeseriesGroupGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupIndustryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupIndustryParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupIndustryParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupIndustryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval15m RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange1d         RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange2d         RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange7d         RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange14d        RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange28d        RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange12w        RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange24w        RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange52w        RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupIndustryParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupIndustryParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupIndustryParamsDirectionTarget RadarAttackLayer7TimeseriesGroupIndustryParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupIndustryParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsFormatJson RadarAttackLayer7TimeseriesGroupIndustryParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupIndustryParamsFormatCsv  RadarAttackLayer7TimeseriesGroupIndustryParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupIndustryParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupIndustryParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupIndustryParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupIndustryResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupIndustryResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupIndustryResponseEnvelope]
type radarAttackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange1d         RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange2d         RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange7d         RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange14d        RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange28d        RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange12w        RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange24w        RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange52w        RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupIPVersionParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupIPVersionParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsDirectionTarget RadarAttackLayer7TimeseriesGroupIPVersionParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsFormatJson RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsFormatCsv  RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolUdp  RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol = "UDP"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolTcp  RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol = "TCP"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolIcmp RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol = "ICMP"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolGre  RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol = "GRE"
)

type RadarAttackLayer7TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupIPVersionResponseEnvelope]
type radarAttackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupProtocolParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupProtocolParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval15m RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange1d         RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange2d         RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange7d         RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange14d        RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange28d        RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange12w        RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange24w        RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange52w        RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupProtocolParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupProtocolParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupProtocolParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupProtocolParamsDirectionTarget RadarAttackLayer7TimeseriesGroupProtocolParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupProtocolParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupProtocolParamsFormatJson RadarAttackLayer7TimeseriesGroupProtocolParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupProtocolParamsFormatCsv  RadarAttackLayer7TimeseriesGroupProtocolParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupProtocolParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupProtocolParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupProtocolParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupProtocolParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupProtocolParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupProtocolResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupProtocolResponseEnvelope]
type radarAttackLayer7TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupVectorParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupVectorParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupVectorParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupVectorParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupVectorParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupVectorParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7TimeseriesGroupVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupVectorParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval15m RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupVectorParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange1d         RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange2d         RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange7d         RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange14d        RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange28d        RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange12w        RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange24w        RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange52w        RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupVectorParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupVectorParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupVectorParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupVectorParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupVectorParamsDirectionTarget RadarAttackLayer7TimeseriesGroupVectorParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupVectorParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsFormatJson RadarAttackLayer7TimeseriesGroupVectorParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupVectorParamsFormatCsv  RadarAttackLayer7TimeseriesGroupVectorParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupVectorParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupVectorParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupVectorParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupVectorParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupVectorParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupVectorParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupVectorParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupVectorParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupVectorParamsProtocol string

const (
	RadarAttackLayer7TimeseriesGroupVectorParamsProtocolUdp  RadarAttackLayer7TimeseriesGroupVectorParamsProtocol = "UDP"
	RadarAttackLayer7TimeseriesGroupVectorParamsProtocolTcp  RadarAttackLayer7TimeseriesGroupVectorParamsProtocol = "TCP"
	RadarAttackLayer7TimeseriesGroupVectorParamsProtocolIcmp RadarAttackLayer7TimeseriesGroupVectorParamsProtocol = "ICMP"
	RadarAttackLayer7TimeseriesGroupVectorParamsProtocolGre  RadarAttackLayer7TimeseriesGroupVectorParamsProtocol = "GRE"
)

type RadarAttackLayer7TimeseriesGroupVectorResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupVectorResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupVectorResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupVectorResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupVectorResponseEnvelope]
type radarAttackLayer7TimeseriesGroupVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesGroupVerticalParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesGroupVerticalParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupVerticalParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupVerticalParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval15m RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange1d         RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange2d         RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange7d         RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange14d        RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange28d        RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange12w        RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange24w        RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange52w        RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesGroupVerticalParamsDirection string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsDirectionOrigin RadarAttackLayer7TimeseriesGroupVerticalParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesGroupVerticalParamsDirectionTarget RadarAttackLayer7TimeseriesGroupVerticalParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupVerticalParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsFormatJson RadarAttackLayer7TimeseriesGroupVerticalParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupVerticalParamsFormatCsv  RadarAttackLayer7TimeseriesGroupVerticalParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupVerticalParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupVerticalParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupVerticalParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupVerticalResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupVerticalResponse             `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupVerticalResponseEnvelope]
type radarAttackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
