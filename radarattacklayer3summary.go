// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarAttackLayer3SummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryService] method instead.
type RadarAttackLayer3SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryService) {
	r = &RadarAttackLayer3SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate.
func (r *RadarAttackLayer3SummaryService) Bitrate(ctx context.Context, query RadarAttackLayer3SummaryBitrateParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryBitrateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3SummaryBitrateResponseEnvelope
	path := "radar/attacks/layer3/summary/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by duration.
func (r *RadarAttackLayer3SummaryService) Duration(ctx context.Context, query RadarAttackLayer3SummaryDurationParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryDurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3SummaryDurationResponseEnvelope
	path := "radar/attacks/layer3/summary/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of network protocols in layer 3/4 attacks over a given
// time period.
func (r *RadarAttackLayer3SummaryService) Get(ctx context.Context, query RadarAttackLayer3SummaryGetParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3SummaryGetResponseEnvelope
	path := "radar/attacks/layer3/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used.
func (r *RadarAttackLayer3SummaryService) IPVersion(ctx context.Context, query RadarAttackLayer3SummaryIPVersionParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3SummaryIPVersionResponseEnvelope
	path := "radar/attacks/layer3/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by protocol used.
func (r *RadarAttackLayer3SummaryService) Protocol(ctx context.Context, query RadarAttackLayer3SummaryProtocolParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3SummaryProtocolResponseEnvelope
	path := "radar/attacks/layer3/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vector.
func (r *RadarAttackLayer3SummaryService) Vector(ctx context.Context, query RadarAttackLayer3SummaryVectorParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryVectorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3SummaryVectorResponseEnvelope
	path := "radar/attacks/layer3/summary/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3SummaryBitrateResponse struct {
	Meta     RadarAttackLayer3SummaryBitrateResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryBitrateResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryBitrateResponseJSON     `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryBitrateResponse]
type radarAttackLayer3SummaryBitrateResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryBitrateResponseMeta struct {
	DateRange      []RadarAttackLayer3SummaryBitrateResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	Normalization  string                                                    `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryBitrateResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryBitrateResponseMeta]
type radarAttackLayer3SummaryBitrateResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryBitrateResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryBitrateResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryBitrateResponseMetaDateRange]
type radarAttackLayer3SummaryBitrateResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfo]
type radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryBitrateResponseSummary0 struct {
	Number1GbpsTo10Gbps   string                                              `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps string                                              `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  string                                              `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           string                                              `json:"OVER_100_GBPS,required"`
	Under500Mbps          string                                              `json:"UNDER_500_MBPS,required"`
	JSON                  radarAttackLayer3SummaryBitrateResponseSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryBitrateResponseSummary0]
type radarAttackLayer3SummaryBitrateResponseSummary0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationResponse struct {
	Meta     RadarAttackLayer3SummaryDurationResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryDurationResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryDurationResponseJSON     `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryDurationResponse]
type radarAttackLayer3SummaryDurationResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationResponseMeta struct {
	DateRange      []RadarAttackLayer3SummaryDurationResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryDurationResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryDurationResponseMeta]
type radarAttackLayer3SummaryDurationResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryDurationResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryDurationResponseMetaDateRange]
type radarAttackLayer3SummaryDurationResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfo]
type radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationResponseSummary0 struct {
	Number1HourTo3Hours  string                                               `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins string                                               `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins string                                               `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  string                                               `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           string                                               `json:"OVER_3_HOURS,required"`
	Under10Mins          string                                               `json:"UNDER_10_MINS,required"`
	JSON                 radarAttackLayer3SummaryDurationResponseSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryDurationResponseSummary0]
type radarAttackLayer3SummaryDurationResponseSummary0JSON struct {
	Number1HourTo3Hours  apijson.Field
	Number10MinsTo20Mins apijson.Field
	Number20MinsTo40Mins apijson.Field
	Number40MinsTo1Hour  apijson.Field
	Over3Hours           apijson.Field
	Under10Mins          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetResponse struct {
	Meta     RadarAttackLayer3SummaryGetResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryGetResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetResponseJSON     `json:"-"`
}

// radarAttackLayer3SummaryGetResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryGetResponse]
type radarAttackLayer3SummaryGetResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetResponseMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryGetResponseMeta]
type radarAttackLayer3SummaryGetResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryGetResponseMetaDateRange]
type radarAttackLayer3SummaryGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarAttackLayer3SummaryGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetResponseMetaConfidenceInfo]
type radarAttackLayer3SummaryGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetResponseSummary0 struct {
	GRE  string                                          `json:"gre,required"`
	Icmp string                                          `json:"icmp,required"`
	Tcp  string                                          `json:"tcp,required"`
	Udp  string                                          `json:"udp,required"`
	JSON radarAttackLayer3SummaryGetResponseSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryGetResponseSummary0]
type radarAttackLayer3SummaryGetResponseSummary0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionResponse struct {
	Meta     RadarAttackLayer3SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryIPVersionResponseJSON     `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryIPVersionResponse]
type radarAttackLayer3SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionResponseMeta struct {
	DateRange      []RadarAttackLayer3SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryIPVersionResponseMeta]
type radarAttackLayer3SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryIPVersionResponseMetaDateRange]
type radarAttackLayer3SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfo]
type radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                                `json:"IPv4,required"`
	IPv6 string                                                `json:"IPv6,required"`
	JSON radarAttackLayer3SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryIPVersionResponseSummary0]
type radarAttackLayer3SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolResponse struct {
	Meta     RadarAttackLayer3SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryProtocolResponseJSON     `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryProtocolResponse]
type radarAttackLayer3SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolResponseMeta struct {
	DateRange      []RadarAttackLayer3SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryProtocolResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryProtocolResponseMeta]
type radarAttackLayer3SummaryProtocolResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryProtocolResponseMetaDateRange]
type radarAttackLayer3SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfo]
type radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolResponseSummary0 struct {
	GRE  string                                               `json:"GRE,required"`
	Icmp string                                               `json:"ICMP,required"`
	Tcp  string                                               `json:"TCP,required"`
	Udp  string                                               `json:"UDP,required"`
	JSON radarAttackLayer3SummaryProtocolResponseSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryProtocolResponseSummary0]
type radarAttackLayer3SummaryProtocolResponseSummary0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryVectorResponse struct {
	Meta     RadarAttackLayer3SummaryVectorResponseMeta `json:"meta,required"`
	Summary0 map[string][]string                        `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryVectorResponseJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryVectorResponse]
type radarAttackLayer3SummaryVectorResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryVectorResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryVectorResponseMeta struct {
	DateRange      []RadarAttackLayer3SummaryVectorResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryVectorResponseMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryVectorResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryVectorResponseMeta]
type radarAttackLayer3SummaryVectorResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryVectorResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryVectorResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryVectorResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryVectorResponseMetaDateRange]
type radarAttackLayer3SummaryVectorResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryVectorResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfo]
type radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryBitrateParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryBitrateParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryBitrateParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryBitrateParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryBitrateParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryBitrateParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3SummaryBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryBitrateParamsDateRange string

const (
	RadarAttackLayer3SummaryBitrateParamsDateRange1d         RadarAttackLayer3SummaryBitrateParamsDateRange = "1d"
	RadarAttackLayer3SummaryBitrateParamsDateRange2d         RadarAttackLayer3SummaryBitrateParamsDateRange = "2d"
	RadarAttackLayer3SummaryBitrateParamsDateRange7d         RadarAttackLayer3SummaryBitrateParamsDateRange = "7d"
	RadarAttackLayer3SummaryBitrateParamsDateRange14d        RadarAttackLayer3SummaryBitrateParamsDateRange = "14d"
	RadarAttackLayer3SummaryBitrateParamsDateRange28d        RadarAttackLayer3SummaryBitrateParamsDateRange = "28d"
	RadarAttackLayer3SummaryBitrateParamsDateRange12w        RadarAttackLayer3SummaryBitrateParamsDateRange = "12w"
	RadarAttackLayer3SummaryBitrateParamsDateRange24w        RadarAttackLayer3SummaryBitrateParamsDateRange = "24w"
	RadarAttackLayer3SummaryBitrateParamsDateRange52w        RadarAttackLayer3SummaryBitrateParamsDateRange = "52w"
	RadarAttackLayer3SummaryBitrateParamsDateRange1dControl  RadarAttackLayer3SummaryBitrateParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryBitrateParamsDateRange2dControl  RadarAttackLayer3SummaryBitrateParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryBitrateParamsDateRange7dControl  RadarAttackLayer3SummaryBitrateParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryBitrateParamsDateRange14dControl RadarAttackLayer3SummaryBitrateParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryBitrateParamsDateRange28dControl RadarAttackLayer3SummaryBitrateParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryBitrateParamsDateRange12wControl RadarAttackLayer3SummaryBitrateParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryBitrateParamsDateRange24wControl RadarAttackLayer3SummaryBitrateParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryBitrateParamsDirection string

const (
	RadarAttackLayer3SummaryBitrateParamsDirectionOrigin RadarAttackLayer3SummaryBitrateParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryBitrateParamsDirectionTarget RadarAttackLayer3SummaryBitrateParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryBitrateParamsFormat string

const (
	RadarAttackLayer3SummaryBitrateParamsFormatJson RadarAttackLayer3SummaryBitrateParamsFormat = "JSON"
	RadarAttackLayer3SummaryBitrateParamsFormatCsv  RadarAttackLayer3SummaryBitrateParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryBitrateParamsIPVersion string

const (
	RadarAttackLayer3SummaryBitrateParamsIPVersionIPv4 RadarAttackLayer3SummaryBitrateParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryBitrateParamsIPVersionIPv6 RadarAttackLayer3SummaryBitrateParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryBitrateParamsProtocol string

const (
	RadarAttackLayer3SummaryBitrateParamsProtocolUdp  RadarAttackLayer3SummaryBitrateParamsProtocol = "UDP"
	RadarAttackLayer3SummaryBitrateParamsProtocolTcp  RadarAttackLayer3SummaryBitrateParamsProtocol = "TCP"
	RadarAttackLayer3SummaryBitrateParamsProtocolIcmp RadarAttackLayer3SummaryBitrateParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryBitrateParamsProtocolGRE  RadarAttackLayer3SummaryBitrateParamsProtocol = "GRE"
)

type RadarAttackLayer3SummaryBitrateResponseEnvelope struct {
	Result  RadarAttackLayer3SummaryBitrateResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAttackLayer3SummaryBitrateResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryBitrateResponseEnvelope]
type radarAttackLayer3SummaryBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryBitrateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryDurationParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryDurationParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryDurationParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryDurationParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryDurationParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryDurationParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3SummaryDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryDurationParamsDateRange string

const (
	RadarAttackLayer3SummaryDurationParamsDateRange1d         RadarAttackLayer3SummaryDurationParamsDateRange = "1d"
	RadarAttackLayer3SummaryDurationParamsDateRange2d         RadarAttackLayer3SummaryDurationParamsDateRange = "2d"
	RadarAttackLayer3SummaryDurationParamsDateRange7d         RadarAttackLayer3SummaryDurationParamsDateRange = "7d"
	RadarAttackLayer3SummaryDurationParamsDateRange14d        RadarAttackLayer3SummaryDurationParamsDateRange = "14d"
	RadarAttackLayer3SummaryDurationParamsDateRange28d        RadarAttackLayer3SummaryDurationParamsDateRange = "28d"
	RadarAttackLayer3SummaryDurationParamsDateRange12w        RadarAttackLayer3SummaryDurationParamsDateRange = "12w"
	RadarAttackLayer3SummaryDurationParamsDateRange24w        RadarAttackLayer3SummaryDurationParamsDateRange = "24w"
	RadarAttackLayer3SummaryDurationParamsDateRange52w        RadarAttackLayer3SummaryDurationParamsDateRange = "52w"
	RadarAttackLayer3SummaryDurationParamsDateRange1dControl  RadarAttackLayer3SummaryDurationParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryDurationParamsDateRange2dControl  RadarAttackLayer3SummaryDurationParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryDurationParamsDateRange7dControl  RadarAttackLayer3SummaryDurationParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryDurationParamsDateRange14dControl RadarAttackLayer3SummaryDurationParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryDurationParamsDateRange28dControl RadarAttackLayer3SummaryDurationParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryDurationParamsDateRange12wControl RadarAttackLayer3SummaryDurationParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryDurationParamsDateRange24wControl RadarAttackLayer3SummaryDurationParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryDurationParamsDirection string

const (
	RadarAttackLayer3SummaryDurationParamsDirectionOrigin RadarAttackLayer3SummaryDurationParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryDurationParamsDirectionTarget RadarAttackLayer3SummaryDurationParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryDurationParamsFormat string

const (
	RadarAttackLayer3SummaryDurationParamsFormatJson RadarAttackLayer3SummaryDurationParamsFormat = "JSON"
	RadarAttackLayer3SummaryDurationParamsFormatCsv  RadarAttackLayer3SummaryDurationParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryDurationParamsIPVersion string

const (
	RadarAttackLayer3SummaryDurationParamsIPVersionIPv4 RadarAttackLayer3SummaryDurationParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryDurationParamsIPVersionIPv6 RadarAttackLayer3SummaryDurationParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryDurationParamsProtocol string

const (
	RadarAttackLayer3SummaryDurationParamsProtocolUdp  RadarAttackLayer3SummaryDurationParamsProtocol = "UDP"
	RadarAttackLayer3SummaryDurationParamsProtocolTcp  RadarAttackLayer3SummaryDurationParamsProtocol = "TCP"
	RadarAttackLayer3SummaryDurationParamsProtocolIcmp RadarAttackLayer3SummaryDurationParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryDurationParamsProtocolGRE  RadarAttackLayer3SummaryDurationParamsProtocol = "GRE"
)

type RadarAttackLayer3SummaryDurationResponseEnvelope struct {
	Result  RadarAttackLayer3SummaryDurationResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer3SummaryDurationResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3SummaryDurationResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryDurationResponseEnvelope]
type radarAttackLayer3SummaryDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryDurationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetParams struct {
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
	DateRange param.Field[[]RadarAttackLayer3SummaryGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3SummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryGetParamsDateRange string

const (
	RadarAttackLayer3SummaryGetParamsDateRange1d         RadarAttackLayer3SummaryGetParamsDateRange = "1d"
	RadarAttackLayer3SummaryGetParamsDateRange2d         RadarAttackLayer3SummaryGetParamsDateRange = "2d"
	RadarAttackLayer3SummaryGetParamsDateRange7d         RadarAttackLayer3SummaryGetParamsDateRange = "7d"
	RadarAttackLayer3SummaryGetParamsDateRange14d        RadarAttackLayer3SummaryGetParamsDateRange = "14d"
	RadarAttackLayer3SummaryGetParamsDateRange28d        RadarAttackLayer3SummaryGetParamsDateRange = "28d"
	RadarAttackLayer3SummaryGetParamsDateRange12w        RadarAttackLayer3SummaryGetParamsDateRange = "12w"
	RadarAttackLayer3SummaryGetParamsDateRange24w        RadarAttackLayer3SummaryGetParamsDateRange = "24w"
	RadarAttackLayer3SummaryGetParamsDateRange52w        RadarAttackLayer3SummaryGetParamsDateRange = "52w"
	RadarAttackLayer3SummaryGetParamsDateRange1dControl  RadarAttackLayer3SummaryGetParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryGetParamsDateRange2dControl  RadarAttackLayer3SummaryGetParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryGetParamsDateRange7dControl  RadarAttackLayer3SummaryGetParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryGetParamsDateRange14dControl RadarAttackLayer3SummaryGetParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryGetParamsDateRange28dControl RadarAttackLayer3SummaryGetParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryGetParamsDateRange12wControl RadarAttackLayer3SummaryGetParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryGetParamsDateRange24wControl RadarAttackLayer3SummaryGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3SummaryGetParamsFormat string

const (
	RadarAttackLayer3SummaryGetParamsFormatJson RadarAttackLayer3SummaryGetParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetParamsFormatCsv  RadarAttackLayer3SummaryGetParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryGetResponseEnvelope struct {
	Result  RadarAttackLayer3SummaryGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryGetResponseEnvelope]
type radarAttackLayer3SummaryGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryIPVersionParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryIPVersionParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryIPVersionParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryIPVersionParamsDateRange string

const (
	RadarAttackLayer3SummaryIPVersionParamsDateRange1d         RadarAttackLayer3SummaryIPVersionParamsDateRange = "1d"
	RadarAttackLayer3SummaryIPVersionParamsDateRange2d         RadarAttackLayer3SummaryIPVersionParamsDateRange = "2d"
	RadarAttackLayer3SummaryIPVersionParamsDateRange7d         RadarAttackLayer3SummaryIPVersionParamsDateRange = "7d"
	RadarAttackLayer3SummaryIPVersionParamsDateRange14d        RadarAttackLayer3SummaryIPVersionParamsDateRange = "14d"
	RadarAttackLayer3SummaryIPVersionParamsDateRange28d        RadarAttackLayer3SummaryIPVersionParamsDateRange = "28d"
	RadarAttackLayer3SummaryIPVersionParamsDateRange12w        RadarAttackLayer3SummaryIPVersionParamsDateRange = "12w"
	RadarAttackLayer3SummaryIPVersionParamsDateRange24w        RadarAttackLayer3SummaryIPVersionParamsDateRange = "24w"
	RadarAttackLayer3SummaryIPVersionParamsDateRange52w        RadarAttackLayer3SummaryIPVersionParamsDateRange = "52w"
	RadarAttackLayer3SummaryIPVersionParamsDateRange1dControl  RadarAttackLayer3SummaryIPVersionParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryIPVersionParamsDateRange2dControl  RadarAttackLayer3SummaryIPVersionParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryIPVersionParamsDateRange7dControl  RadarAttackLayer3SummaryIPVersionParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryIPVersionParamsDateRange14dControl RadarAttackLayer3SummaryIPVersionParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryIPVersionParamsDateRange28dControl RadarAttackLayer3SummaryIPVersionParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryIPVersionParamsDateRange12wControl RadarAttackLayer3SummaryIPVersionParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryIPVersionParamsDateRange24wControl RadarAttackLayer3SummaryIPVersionParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryIPVersionParamsDirection string

const (
	RadarAttackLayer3SummaryIPVersionParamsDirectionOrigin RadarAttackLayer3SummaryIPVersionParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryIPVersionParamsDirectionTarget RadarAttackLayer3SummaryIPVersionParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryIPVersionParamsFormat string

const (
	RadarAttackLayer3SummaryIPVersionParamsFormatJson RadarAttackLayer3SummaryIPVersionParamsFormat = "JSON"
	RadarAttackLayer3SummaryIPVersionParamsFormatCsv  RadarAttackLayer3SummaryIPVersionParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryIPVersionParamsProtocol string

const (
	RadarAttackLayer3SummaryIPVersionParamsProtocolUdp  RadarAttackLayer3SummaryIPVersionParamsProtocol = "UDP"
	RadarAttackLayer3SummaryIPVersionParamsProtocolTcp  RadarAttackLayer3SummaryIPVersionParamsProtocol = "TCP"
	RadarAttackLayer3SummaryIPVersionParamsProtocolIcmp RadarAttackLayer3SummaryIPVersionParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryIPVersionParamsProtocolGRE  RadarAttackLayer3SummaryIPVersionParamsProtocol = "GRE"
)

type RadarAttackLayer3SummaryIPVersionResponseEnvelope struct {
	Result  RadarAttackLayer3SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAttackLayer3SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryIPVersionResponseEnvelope]
type radarAttackLayer3SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryProtocolParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryProtocolParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryProtocolParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3SummaryProtocolParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryProtocolParamsDateRange string

const (
	RadarAttackLayer3SummaryProtocolParamsDateRange1d         RadarAttackLayer3SummaryProtocolParamsDateRange = "1d"
	RadarAttackLayer3SummaryProtocolParamsDateRange2d         RadarAttackLayer3SummaryProtocolParamsDateRange = "2d"
	RadarAttackLayer3SummaryProtocolParamsDateRange7d         RadarAttackLayer3SummaryProtocolParamsDateRange = "7d"
	RadarAttackLayer3SummaryProtocolParamsDateRange14d        RadarAttackLayer3SummaryProtocolParamsDateRange = "14d"
	RadarAttackLayer3SummaryProtocolParamsDateRange28d        RadarAttackLayer3SummaryProtocolParamsDateRange = "28d"
	RadarAttackLayer3SummaryProtocolParamsDateRange12w        RadarAttackLayer3SummaryProtocolParamsDateRange = "12w"
	RadarAttackLayer3SummaryProtocolParamsDateRange24w        RadarAttackLayer3SummaryProtocolParamsDateRange = "24w"
	RadarAttackLayer3SummaryProtocolParamsDateRange52w        RadarAttackLayer3SummaryProtocolParamsDateRange = "52w"
	RadarAttackLayer3SummaryProtocolParamsDateRange1dControl  RadarAttackLayer3SummaryProtocolParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryProtocolParamsDateRange2dControl  RadarAttackLayer3SummaryProtocolParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryProtocolParamsDateRange7dControl  RadarAttackLayer3SummaryProtocolParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryProtocolParamsDateRange14dControl RadarAttackLayer3SummaryProtocolParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryProtocolParamsDateRange28dControl RadarAttackLayer3SummaryProtocolParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryProtocolParamsDateRange12wControl RadarAttackLayer3SummaryProtocolParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryProtocolParamsDateRange24wControl RadarAttackLayer3SummaryProtocolParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryProtocolParamsDirection string

const (
	RadarAttackLayer3SummaryProtocolParamsDirectionOrigin RadarAttackLayer3SummaryProtocolParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryProtocolParamsDirectionTarget RadarAttackLayer3SummaryProtocolParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryProtocolParamsFormat string

const (
	RadarAttackLayer3SummaryProtocolParamsFormatJson RadarAttackLayer3SummaryProtocolParamsFormat = "JSON"
	RadarAttackLayer3SummaryProtocolParamsFormatCsv  RadarAttackLayer3SummaryProtocolParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryProtocolParamsIPVersion string

const (
	RadarAttackLayer3SummaryProtocolParamsIPVersionIPv4 RadarAttackLayer3SummaryProtocolParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryProtocolParamsIPVersionIPv6 RadarAttackLayer3SummaryProtocolParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryProtocolResponseEnvelope struct {
	Result  RadarAttackLayer3SummaryProtocolResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer3SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryProtocolResponseEnvelope]
type radarAttackLayer3SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryVectorParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryVectorParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryVectorParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryVectorParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryVectorParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryVectorParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3SummaryVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryVectorParamsDateRange string

const (
	RadarAttackLayer3SummaryVectorParamsDateRange1d         RadarAttackLayer3SummaryVectorParamsDateRange = "1d"
	RadarAttackLayer3SummaryVectorParamsDateRange2d         RadarAttackLayer3SummaryVectorParamsDateRange = "2d"
	RadarAttackLayer3SummaryVectorParamsDateRange7d         RadarAttackLayer3SummaryVectorParamsDateRange = "7d"
	RadarAttackLayer3SummaryVectorParamsDateRange14d        RadarAttackLayer3SummaryVectorParamsDateRange = "14d"
	RadarAttackLayer3SummaryVectorParamsDateRange28d        RadarAttackLayer3SummaryVectorParamsDateRange = "28d"
	RadarAttackLayer3SummaryVectorParamsDateRange12w        RadarAttackLayer3SummaryVectorParamsDateRange = "12w"
	RadarAttackLayer3SummaryVectorParamsDateRange24w        RadarAttackLayer3SummaryVectorParamsDateRange = "24w"
	RadarAttackLayer3SummaryVectorParamsDateRange52w        RadarAttackLayer3SummaryVectorParamsDateRange = "52w"
	RadarAttackLayer3SummaryVectorParamsDateRange1dControl  RadarAttackLayer3SummaryVectorParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryVectorParamsDateRange2dControl  RadarAttackLayer3SummaryVectorParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryVectorParamsDateRange7dControl  RadarAttackLayer3SummaryVectorParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryVectorParamsDateRange14dControl RadarAttackLayer3SummaryVectorParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryVectorParamsDateRange28dControl RadarAttackLayer3SummaryVectorParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryVectorParamsDateRange12wControl RadarAttackLayer3SummaryVectorParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryVectorParamsDateRange24wControl RadarAttackLayer3SummaryVectorParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryVectorParamsDirection string

const (
	RadarAttackLayer3SummaryVectorParamsDirectionOrigin RadarAttackLayer3SummaryVectorParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryVectorParamsDirectionTarget RadarAttackLayer3SummaryVectorParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryVectorParamsFormat string

const (
	RadarAttackLayer3SummaryVectorParamsFormatJson RadarAttackLayer3SummaryVectorParamsFormat = "JSON"
	RadarAttackLayer3SummaryVectorParamsFormatCsv  RadarAttackLayer3SummaryVectorParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryVectorParamsIPVersion string

const (
	RadarAttackLayer3SummaryVectorParamsIPVersionIPv4 RadarAttackLayer3SummaryVectorParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryVectorParamsIPVersionIPv6 RadarAttackLayer3SummaryVectorParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryVectorParamsProtocol string

const (
	RadarAttackLayer3SummaryVectorParamsProtocolUdp  RadarAttackLayer3SummaryVectorParamsProtocol = "UDP"
	RadarAttackLayer3SummaryVectorParamsProtocolTcp  RadarAttackLayer3SummaryVectorParamsProtocol = "TCP"
	RadarAttackLayer3SummaryVectorParamsProtocolIcmp RadarAttackLayer3SummaryVectorParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryVectorParamsProtocolGRE  RadarAttackLayer3SummaryVectorParamsProtocol = "GRE"
)

type RadarAttackLayer3SummaryVectorResponseEnvelope struct {
	Result  RadarAttackLayer3SummaryVectorResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAttackLayer3SummaryVectorResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryVectorResponseEnvelope]
type radarAttackLayer3SummaryVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryVectorResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
