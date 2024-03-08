// File generated from our OpenAPI spec by Stainless.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AttackLayer3TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAttackLayer3TimeseriesGroupService] method instead.
type AttackLayer3TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAttackLayer3TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackLayer3TimeseriesGroupService(opts ...option.RequestOption) (r *AttackLayer3TimeseriesGroupService) {
	r = &AttackLayer3TimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate over time.
func (r *AttackLayer3TimeseriesGroupService) Bitrate(ctx context.Context, query AttackLayer3TimeseriesGroupBitrateParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupBitrateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupBitrateResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by duration over time.
func (r *AttackLayer3TimeseriesGroupService) Duration(ctx context.Context, query AttackLayer3TimeseriesGroupDurationParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupDurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupDurationResponseEnvelope
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
func (r *AttackLayer3TimeseriesGroupService) Get(ctx context.Context, query AttackLayer3TimeseriesGroupGetParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupGetResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by industry used over time.
func (r *AttackLayer3TimeseriesGroupService) Industry(ctx context.Context, query AttackLayer3TimeseriesGroupIndustryParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupIndustryResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used over time.
func (r *AttackLayer3TimeseriesGroupService) IPVersion(ctx context.Context, query AttackLayer3TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupIPVersionResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by protocol used over time.
func (r *AttackLayer3TimeseriesGroupService) Protocol(ctx context.Context, query AttackLayer3TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupProtocolResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vector used over time.
func (r *AttackLayer3TimeseriesGroupService) Vector(ctx context.Context, query AttackLayer3TimeseriesGroupVectorParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupVectorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupVectorResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vertical used over time.
func (r *AttackLayer3TimeseriesGroupService) Vertical(ctx context.Context, query AttackLayer3TimeseriesGroupVerticalParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TimeseriesGroupVerticalResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TimeseriesGroupBitrateResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupBitrateResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupBitrateResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupBitrateResponse]
type attackLayer3TimeseriesGroupBitrateResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupBitrateResponseSerie0 struct {
	Number1GbpsTo10Gbps   []string                                             `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps []string                                             `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  []string                                             `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           []string                                             `json:"OVER_100_GBPS,required"`
	Timestamps            []string                                             `json:"timestamps,required"`
	Under500Mbps          []string                                             `json:"UNDER_500_MBPS,required"`
	JSON                  attackLayer3TimeseriesGroupBitrateResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupBitrateResponseSerie0]
type attackLayer3TimeseriesGroupBitrateResponseSerie0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Timestamps            apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupDurationResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupDurationResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupDurationResponse]
type attackLayer3TimeseriesGroupDurationResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationResponseSerie0 struct {
	Number1HourTo3Hours  []string                                              `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins []string                                              `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins []string                                              `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  []string                                              `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           []string                                              `json:"OVER_3_HOURS,required"`
	Timestamps           []string                                              `json:"timestamps,required"`
	Under10Mins          []string                                              `json:"UNDER_10_MINS,required"`
	JSON                 attackLayer3TimeseriesGroupDurationResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupDurationResponseSerie0]
type attackLayer3TimeseriesGroupDurationResponseSerie0JSON struct {
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

func (r *AttackLayer3TimeseriesGroupDurationResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetResponse struct {
	Meta   AttackLayer3TimeseriesGroupGetResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupGetResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupGetResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesGroupGetResponse]
type attackLayer3TimeseriesGroupGetResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetResponseMeta struct {
	AggInterval    string                                                   `json:"aggInterval,required"`
	DateRange      []AttackLayer3TimeseriesGroupGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3TimeseriesGroupGetResponseMetaJSON           `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupGetResponseMeta]
type attackLayer3TimeseriesGroupGetResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupGetResponseMetaDateRange]
type attackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetResponseSerie0 struct {
	GRE        []string                                         `json:"gre,required"`
	Icmp       []string                                         `json:"icmp,required"`
	Tcp        []string                                         `json:"tcp,required"`
	Timestamps []string                                         `json:"timestamps,required"`
	Udp        []string                                         `json:"udp,required"`
	JSON       attackLayer3TimeseriesGroupGetResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseSerie0JSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupGetResponseSerie0]
type attackLayer3TimeseriesGroupGetResponseSerie0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupGetResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupIndustryResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupIndustryResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupIndustryResponse]
type attackLayer3TimeseriesGroupIndustryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryResponseSerie0 struct {
	Timestamps  []string                                              `json:"timestamps,required"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupIndustryResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupIndustryResponseSerie0]
type attackLayer3TimeseriesGroupIndustryResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupIPVersionResponse]
type attackLayer3TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                               `json:"IPv4,required"`
	IPv6       []string                                               `json:"IPv6,required"`
	Timestamps []string                                               `json:"timestamps,required"`
	JSON       attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIPVersionResponseSerie0]
type attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupProtocolResponse]
type attackLayer3TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolResponseSerie0 struct {
	GRE        []string                                              `json:"GRE,required"`
	Icmp       []string                                              `json:"ICMP,required"`
	Tcp        []string                                              `json:"TCP,required"`
	Timestamps []string                                              `json:"timestamps,required"`
	Udp        []string                                              `json:"UDP,required"`
	JSON       attackLayer3TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupProtocolResponseSerie0]
type attackLayer3TimeseriesGroupProtocolResponseSerie0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorResponse struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupVectorResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupVectorResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesGroupVectorResponse]
type attackLayer3TimeseriesGroupVectorResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorResponseSerie0 struct {
	Timestamps  []string                                            `json:"timestamps,required"`
	ExtraFields map[string][]string                                 `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupVectorResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVectorResponseSerie0]
type attackLayer3TimeseriesGroupVectorResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupVerticalResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupVerticalResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupVerticalResponse]
type attackLayer3TimeseriesGroupVerticalResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalResponseSerie0 struct {
	Timestamps  []string                                              `json:"timestamps,required"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupVerticalResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVerticalResponseSerie0]
type attackLayer3TimeseriesGroupVerticalResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupBitrateParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupBitrateParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupBitrateParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupBitrateParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupBitrateParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupBitrateParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupBitrateParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupBitrateParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupBitrateParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval15m AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval1h  AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval1d  AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval1w  AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupBitrateParamsDateRange string

const (
	AttackLayer3TimeseriesGroupBitrateParamsDateRange1d         AttackLayer3TimeseriesGroupBitrateParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange2d         AttackLayer3TimeseriesGroupBitrateParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange7d         AttackLayer3TimeseriesGroupBitrateParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange14d        AttackLayer3TimeseriesGroupBitrateParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange28d        AttackLayer3TimeseriesGroupBitrateParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange12w        AttackLayer3TimeseriesGroupBitrateParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange24w        AttackLayer3TimeseriesGroupBitrateParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange52w        AttackLayer3TimeseriesGroupBitrateParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange1dControl  AttackLayer3TimeseriesGroupBitrateParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange2dControl  AttackLayer3TimeseriesGroupBitrateParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange7dControl  AttackLayer3TimeseriesGroupBitrateParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange14dControl AttackLayer3TimeseriesGroupBitrateParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange28dControl AttackLayer3TimeseriesGroupBitrateParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange12wControl AttackLayer3TimeseriesGroupBitrateParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupBitrateParamsDateRange24wControl AttackLayer3TimeseriesGroupBitrateParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupBitrateParamsDirection string

const (
	AttackLayer3TimeseriesGroupBitrateParamsDirectionOrigin AttackLayer3TimeseriesGroupBitrateParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupBitrateParamsDirectionTarget AttackLayer3TimeseriesGroupBitrateParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupBitrateParamsFormat string

const (
	AttackLayer3TimeseriesGroupBitrateParamsFormatJson AttackLayer3TimeseriesGroupBitrateParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupBitrateParamsFormatCsv  AttackLayer3TimeseriesGroupBitrateParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupBitrateParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv4 AttackLayer3TimeseriesGroupBitrateParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv6 AttackLayer3TimeseriesGroupBitrateParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupBitrateParamsNormalization string

const (
	AttackLayer3TimeseriesGroupBitrateParamsNormalizationPercentage AttackLayer3TimeseriesGroupBitrateParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupBitrateParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupBitrateParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupBitrateParamsProtocol string

const (
	AttackLayer3TimeseriesGroupBitrateParamsProtocolUdp  AttackLayer3TimeseriesGroupBitrateParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupBitrateParamsProtocolTcp  AttackLayer3TimeseriesGroupBitrateParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupBitrateParamsProtocolIcmp AttackLayer3TimeseriesGroupBitrateParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupBitrateParamsProtocolGRE  AttackLayer3TimeseriesGroupBitrateParamsProtocol = "GRE"
)

type AttackLayer3TimeseriesGroupBitrateResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupBitrateResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupBitrateResponseEnvelope]
type attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupDurationParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupDurationParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupDurationParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupDurationParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupDurationParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupDurationParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupDurationParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupDurationParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupDurationParamsAggInterval15m AttackLayer3TimeseriesGroupDurationParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupDurationParamsAggInterval1h  AttackLayer3TimeseriesGroupDurationParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupDurationParamsAggInterval1d  AttackLayer3TimeseriesGroupDurationParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupDurationParamsAggInterval1w  AttackLayer3TimeseriesGroupDurationParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupDurationParamsDateRange string

const (
	AttackLayer3TimeseriesGroupDurationParamsDateRange1d         AttackLayer3TimeseriesGroupDurationParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupDurationParamsDateRange2d         AttackLayer3TimeseriesGroupDurationParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupDurationParamsDateRange7d         AttackLayer3TimeseriesGroupDurationParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupDurationParamsDateRange14d        AttackLayer3TimeseriesGroupDurationParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupDurationParamsDateRange28d        AttackLayer3TimeseriesGroupDurationParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupDurationParamsDateRange12w        AttackLayer3TimeseriesGroupDurationParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupDurationParamsDateRange24w        AttackLayer3TimeseriesGroupDurationParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupDurationParamsDateRange52w        AttackLayer3TimeseriesGroupDurationParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupDurationParamsDateRange1dControl  AttackLayer3TimeseriesGroupDurationParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupDurationParamsDateRange2dControl  AttackLayer3TimeseriesGroupDurationParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupDurationParamsDateRange7dControl  AttackLayer3TimeseriesGroupDurationParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupDurationParamsDateRange14dControl AttackLayer3TimeseriesGroupDurationParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupDurationParamsDateRange28dControl AttackLayer3TimeseriesGroupDurationParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupDurationParamsDateRange12wControl AttackLayer3TimeseriesGroupDurationParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupDurationParamsDateRange24wControl AttackLayer3TimeseriesGroupDurationParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupDurationParamsDirection string

const (
	AttackLayer3TimeseriesGroupDurationParamsDirectionOrigin AttackLayer3TimeseriesGroupDurationParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupDurationParamsDirectionTarget AttackLayer3TimeseriesGroupDurationParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupDurationParamsFormat string

const (
	AttackLayer3TimeseriesGroupDurationParamsFormatJson AttackLayer3TimeseriesGroupDurationParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupDurationParamsFormatCsv  AttackLayer3TimeseriesGroupDurationParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupDurationParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv4 AttackLayer3TimeseriesGroupDurationParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv6 AttackLayer3TimeseriesGroupDurationParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupDurationParamsNormalization string

const (
	AttackLayer3TimeseriesGroupDurationParamsNormalizationPercentage AttackLayer3TimeseriesGroupDurationParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupDurationParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupDurationParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupDurationParamsProtocol string

const (
	AttackLayer3TimeseriesGroupDurationParamsProtocolUdp  AttackLayer3TimeseriesGroupDurationParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupDurationParamsProtocolTcp  AttackLayer3TimeseriesGroupDurationParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupDurationParamsProtocolIcmp AttackLayer3TimeseriesGroupDurationParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupDurationParamsProtocolGRE  AttackLayer3TimeseriesGroupDurationParamsProtocol = "GRE"
)

type AttackLayer3TimeseriesGroupDurationResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupDurationResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupDurationResponseEnvelope]
type attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupGetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TimeseriesGroupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupGetParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupGetParamsAggInterval15m AttackLayer3TimeseriesGroupGetParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupGetParamsAggInterval1h  AttackLayer3TimeseriesGroupGetParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupGetParamsAggInterval1d  AttackLayer3TimeseriesGroupGetParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupGetParamsAggInterval1w  AttackLayer3TimeseriesGroupGetParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupGetParamsDateRange string

const (
	AttackLayer3TimeseriesGroupGetParamsDateRange1d         AttackLayer3TimeseriesGroupGetParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupGetParamsDateRange2d         AttackLayer3TimeseriesGroupGetParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupGetParamsDateRange7d         AttackLayer3TimeseriesGroupGetParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupGetParamsDateRange14d        AttackLayer3TimeseriesGroupGetParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupGetParamsDateRange28d        AttackLayer3TimeseriesGroupGetParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupGetParamsDateRange12w        AttackLayer3TimeseriesGroupGetParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupGetParamsDateRange24w        AttackLayer3TimeseriesGroupGetParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupGetParamsDateRange52w        AttackLayer3TimeseriesGroupGetParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupGetParamsDateRange1dControl  AttackLayer3TimeseriesGroupGetParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupGetParamsDateRange2dControl  AttackLayer3TimeseriesGroupGetParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupGetParamsDateRange7dControl  AttackLayer3TimeseriesGroupGetParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupGetParamsDateRange14dControl AttackLayer3TimeseriesGroupGetParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupGetParamsDateRange28dControl AttackLayer3TimeseriesGroupGetParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupGetParamsDateRange12wControl AttackLayer3TimeseriesGroupGetParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupGetParamsDateRange24wControl AttackLayer3TimeseriesGroupGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupGetParamsFormat string

const (
	AttackLayer3TimeseriesGroupGetParamsFormatJson AttackLayer3TimeseriesGroupGetParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupGetParamsFormatCsv  AttackLayer3TimeseriesGroupGetParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupGetResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupGetResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupGetResponseEnvelope]
type attackLayer3TimeseriesGroupGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupIndustryParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[AttackLayer3TimeseriesGroupIndustryParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupIndustryParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupIndustryParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval15m AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval1h  AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval1d  AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval1w  AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupIndustryParamsDateRange string

const (
	AttackLayer3TimeseriesGroupIndustryParamsDateRange1d         AttackLayer3TimeseriesGroupIndustryParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange2d         AttackLayer3TimeseriesGroupIndustryParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange7d         AttackLayer3TimeseriesGroupIndustryParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange14d        AttackLayer3TimeseriesGroupIndustryParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange28d        AttackLayer3TimeseriesGroupIndustryParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange12w        AttackLayer3TimeseriesGroupIndustryParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange24w        AttackLayer3TimeseriesGroupIndustryParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange52w        AttackLayer3TimeseriesGroupIndustryParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange1dControl  AttackLayer3TimeseriesGroupIndustryParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange2dControl  AttackLayer3TimeseriesGroupIndustryParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange7dControl  AttackLayer3TimeseriesGroupIndustryParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange14dControl AttackLayer3TimeseriesGroupIndustryParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange28dControl AttackLayer3TimeseriesGroupIndustryParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange12wControl AttackLayer3TimeseriesGroupIndustryParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupIndustryParamsDateRange24wControl AttackLayer3TimeseriesGroupIndustryParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupIndustryParamsDirection string

const (
	AttackLayer3TimeseriesGroupIndustryParamsDirectionOrigin AttackLayer3TimeseriesGroupIndustryParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupIndustryParamsDirectionTarget AttackLayer3TimeseriesGroupIndustryParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupIndustryParamsFormat string

const (
	AttackLayer3TimeseriesGroupIndustryParamsFormatJson AttackLayer3TimeseriesGroupIndustryParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupIndustryParamsFormatCsv  AttackLayer3TimeseriesGroupIndustryParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupIndustryParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv4 AttackLayer3TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv6 AttackLayer3TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupIndustryParamsNormalization string

const (
	AttackLayer3TimeseriesGroupIndustryParamsNormalizationPercentage AttackLayer3TimeseriesGroupIndustryParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupIndustryParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupIndustryParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupIndustryResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupIndustryResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIndustryResponseEnvelope]
type attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupIPVersionParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupIPVersionParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval15m AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1h  AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1d  AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1w  AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupIPVersionParamsDateRange string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange1d         AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange2d         AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange7d         AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange14d        AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange28d        AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange12w        AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange24w        AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange52w        AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange1dControl  AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange2dControl  AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange7dControl  AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange14dControl AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange28dControl AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange12wControl AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupIPVersionParamsDateRange24wControl AttackLayer3TimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupIPVersionParamsDirection string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsDirectionOrigin AttackLayer3TimeseriesGroupIPVersionParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupIPVersionParamsDirectionTarget AttackLayer3TimeseriesGroupIPVersionParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupIPVersionParamsFormat string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsFormatJson AttackLayer3TimeseriesGroupIPVersionParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupIPVersionParamsFormatCsv  AttackLayer3TimeseriesGroupIPVersionParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupIPVersionParamsNormalization string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsNormalizationPercentage AttackLayer3TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupIPVersionParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupIPVersionParamsProtocol string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolUdp  AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolTcp  AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolIcmp AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolGRE  AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "GRE"
)

type AttackLayer3TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIPVersionResponseEnvelope]
type attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupProtocolParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupProtocolParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupProtocolParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupProtocolParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval15m AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval1h  AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval1d  AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval1w  AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupProtocolParamsDateRange string

const (
	AttackLayer3TimeseriesGroupProtocolParamsDateRange1d         AttackLayer3TimeseriesGroupProtocolParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange2d         AttackLayer3TimeseriesGroupProtocolParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange7d         AttackLayer3TimeseriesGroupProtocolParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange14d        AttackLayer3TimeseriesGroupProtocolParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange28d        AttackLayer3TimeseriesGroupProtocolParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange12w        AttackLayer3TimeseriesGroupProtocolParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange24w        AttackLayer3TimeseriesGroupProtocolParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange52w        AttackLayer3TimeseriesGroupProtocolParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange1dControl  AttackLayer3TimeseriesGroupProtocolParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange2dControl  AttackLayer3TimeseriesGroupProtocolParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange7dControl  AttackLayer3TimeseriesGroupProtocolParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange14dControl AttackLayer3TimeseriesGroupProtocolParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange28dControl AttackLayer3TimeseriesGroupProtocolParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange12wControl AttackLayer3TimeseriesGroupProtocolParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupProtocolParamsDateRange24wControl AttackLayer3TimeseriesGroupProtocolParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupProtocolParamsDirection string

const (
	AttackLayer3TimeseriesGroupProtocolParamsDirectionOrigin AttackLayer3TimeseriesGroupProtocolParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupProtocolParamsDirectionTarget AttackLayer3TimeseriesGroupProtocolParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupProtocolParamsFormat string

const (
	AttackLayer3TimeseriesGroupProtocolParamsFormatJson AttackLayer3TimeseriesGroupProtocolParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupProtocolParamsFormatCsv  AttackLayer3TimeseriesGroupProtocolParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupProtocolParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv4 AttackLayer3TimeseriesGroupProtocolParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv6 AttackLayer3TimeseriesGroupProtocolParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupProtocolParamsNormalization string

const (
	AttackLayer3TimeseriesGroupProtocolParamsNormalizationPercentage AttackLayer3TimeseriesGroupProtocolParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupProtocolParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupProtocolParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupProtocolResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupProtocolResponseEnvelope]
type attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupVectorParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupVectorParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupVectorParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupVectorParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupVectorParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[AttackLayer3TimeseriesGroupVectorParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupVectorParams]'s query parameters
// as `url.Values`.
func (r AttackLayer3TimeseriesGroupVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupVectorParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupVectorParamsAggInterval15m AttackLayer3TimeseriesGroupVectorParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupVectorParamsAggInterval1h  AttackLayer3TimeseriesGroupVectorParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupVectorParamsAggInterval1d  AttackLayer3TimeseriesGroupVectorParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupVectorParamsAggInterval1w  AttackLayer3TimeseriesGroupVectorParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupVectorParamsDateRange string

const (
	AttackLayer3TimeseriesGroupVectorParamsDateRange1d         AttackLayer3TimeseriesGroupVectorParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupVectorParamsDateRange2d         AttackLayer3TimeseriesGroupVectorParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupVectorParamsDateRange7d         AttackLayer3TimeseriesGroupVectorParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupVectorParamsDateRange14d        AttackLayer3TimeseriesGroupVectorParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupVectorParamsDateRange28d        AttackLayer3TimeseriesGroupVectorParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupVectorParamsDateRange12w        AttackLayer3TimeseriesGroupVectorParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupVectorParamsDateRange24w        AttackLayer3TimeseriesGroupVectorParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupVectorParamsDateRange52w        AttackLayer3TimeseriesGroupVectorParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupVectorParamsDateRange1dControl  AttackLayer3TimeseriesGroupVectorParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupVectorParamsDateRange2dControl  AttackLayer3TimeseriesGroupVectorParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupVectorParamsDateRange7dControl  AttackLayer3TimeseriesGroupVectorParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupVectorParamsDateRange14dControl AttackLayer3TimeseriesGroupVectorParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupVectorParamsDateRange28dControl AttackLayer3TimeseriesGroupVectorParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupVectorParamsDateRange12wControl AttackLayer3TimeseriesGroupVectorParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupVectorParamsDateRange24wControl AttackLayer3TimeseriesGroupVectorParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupVectorParamsDirection string

const (
	AttackLayer3TimeseriesGroupVectorParamsDirectionOrigin AttackLayer3TimeseriesGroupVectorParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupVectorParamsDirectionTarget AttackLayer3TimeseriesGroupVectorParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupVectorParamsFormat string

const (
	AttackLayer3TimeseriesGroupVectorParamsFormatJson AttackLayer3TimeseriesGroupVectorParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupVectorParamsFormatCsv  AttackLayer3TimeseriesGroupVectorParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupVectorParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv4 AttackLayer3TimeseriesGroupVectorParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv6 AttackLayer3TimeseriesGroupVectorParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupVectorParamsNormalization string

const (
	AttackLayer3TimeseriesGroupVectorParamsNormalizationPercentage AttackLayer3TimeseriesGroupVectorParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupVectorParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupVectorParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupVectorParamsProtocol string

const (
	AttackLayer3TimeseriesGroupVectorParamsProtocolUdp  AttackLayer3TimeseriesGroupVectorParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupVectorParamsProtocolTcp  AttackLayer3TimeseriesGroupVectorParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupVectorParamsProtocolIcmp AttackLayer3TimeseriesGroupVectorParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupVectorParamsProtocolGRE  AttackLayer3TimeseriesGroupVectorParamsProtocol = "GRE"
)

type AttackLayer3TimeseriesGroupVectorResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupVectorResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVectorResponseEnvelope]
type attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TimeseriesGroupVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3TimeseriesGroupVerticalParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[AttackLayer3TimeseriesGroupVerticalParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupVerticalParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupVerticalParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval15m AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval1h  AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval1d  AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval1w  AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1w"
)

type AttackLayer3TimeseriesGroupVerticalParamsDateRange string

const (
	AttackLayer3TimeseriesGroupVerticalParamsDateRange1d         AttackLayer3TimeseriesGroupVerticalParamsDateRange = "1d"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange2d         AttackLayer3TimeseriesGroupVerticalParamsDateRange = "2d"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange7d         AttackLayer3TimeseriesGroupVerticalParamsDateRange = "7d"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange14d        AttackLayer3TimeseriesGroupVerticalParamsDateRange = "14d"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange28d        AttackLayer3TimeseriesGroupVerticalParamsDateRange = "28d"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange12w        AttackLayer3TimeseriesGroupVerticalParamsDateRange = "12w"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange24w        AttackLayer3TimeseriesGroupVerticalParamsDateRange = "24w"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange52w        AttackLayer3TimeseriesGroupVerticalParamsDateRange = "52w"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange1dControl  AttackLayer3TimeseriesGroupVerticalParamsDateRange = "1dControl"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange2dControl  AttackLayer3TimeseriesGroupVerticalParamsDateRange = "2dControl"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange7dControl  AttackLayer3TimeseriesGroupVerticalParamsDateRange = "7dControl"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange14dControl AttackLayer3TimeseriesGroupVerticalParamsDateRange = "14dControl"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange28dControl AttackLayer3TimeseriesGroupVerticalParamsDateRange = "28dControl"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange12wControl AttackLayer3TimeseriesGroupVerticalParamsDateRange = "12wControl"
	AttackLayer3TimeseriesGroupVerticalParamsDateRange24wControl AttackLayer3TimeseriesGroupVerticalParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3TimeseriesGroupVerticalParamsDirection string

const (
	AttackLayer3TimeseriesGroupVerticalParamsDirectionOrigin AttackLayer3TimeseriesGroupVerticalParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupVerticalParamsDirectionTarget AttackLayer3TimeseriesGroupVerticalParamsDirection = "TARGET"
)

// Format results are returned in.
type AttackLayer3TimeseriesGroupVerticalParamsFormat string

const (
	AttackLayer3TimeseriesGroupVerticalParamsFormatJson AttackLayer3TimeseriesGroupVerticalParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupVerticalParamsFormatCsv  AttackLayer3TimeseriesGroupVerticalParamsFormat = "CSV"
)

type AttackLayer3TimeseriesGroupVerticalParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv4 AttackLayer3TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv6 AttackLayer3TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupVerticalParamsNormalization string

const (
	AttackLayer3TimeseriesGroupVerticalParamsNormalizationPercentage AttackLayer3TimeseriesGroupVerticalParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupVerticalParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupVerticalParamsNormalization = "MIN0_MAX"
)

type AttackLayer3TimeseriesGroupVerticalResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupVerticalResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupVerticalResponseEnvelope]
type attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
