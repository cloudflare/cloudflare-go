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

// RadarAttackLayer7SummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7SummaryService] method instead.
type RadarAttackLayer7SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7SummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7SummaryService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryService) {
	r = &RadarAttackLayer7SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate.
func (r *RadarAttackLayer7SummaryService) Bitrate(ctx context.Context, query RadarAttackLayer7SummaryBitrateParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryBitrateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryBitrateResponseEnvelope
	path := "radar/attacks/layer3/summary/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by duration.
func (r *RadarAttackLayer7SummaryService) Duration(ctx context.Context, query RadarAttackLayer7SummaryDurationParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryDurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryDurationResponseEnvelope
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
func (r *RadarAttackLayer7SummaryService) Get(ctx context.Context, query RadarAttackLayer7SummaryGetParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryGetResponseEnvelope
	path := "radar/attacks/layer3/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used.
func (r *RadarAttackLayer7SummaryService) IPVersion(ctx context.Context, query RadarAttackLayer7SummaryIPVersionParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryIPVersionResponseEnvelope
	path := "radar/attacks/layer3/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by protocol used.
func (r *RadarAttackLayer7SummaryService) Protocol(ctx context.Context, query RadarAttackLayer7SummaryProtocolParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryProtocolResponseEnvelope
	path := "radar/attacks/layer3/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vector.
func (r *RadarAttackLayer7SummaryService) Vector(ctx context.Context, query RadarAttackLayer7SummaryVectorParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryVectorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryVectorResponseEnvelope
	path := "radar/attacks/layer3/summary/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7SummaryBitrateResponse struct {
	Meta     RadarAttackLayer7SummaryBitrateResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryBitrateResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryBitrateResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryBitrateResponse]
type radarAttackLayer7SummaryBitrateResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryBitrateResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryBitrateResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	Normalization  string                                                    `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryBitrateResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryBitrateResponseMeta]
type radarAttackLayer7SummaryBitrateResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryBitrateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryBitrateResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryBitrateResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryBitrateResponseMetaDateRange]
type radarAttackLayer7SummaryBitrateResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryBitrateResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryBitrateResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryBitrateResponseSummary0 struct {
	Number1GbpsTo10Gbps   string                                              `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps string                                              `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  string                                              `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           string                                              `json:"OVER_100_GBPS,required"`
	Under500Mbps          string                                              `json:"UNDER_500_MBPS,required"`
	JSON                  radarAttackLayer7SummaryBitrateResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryBitrateResponseSummary0]
type radarAttackLayer7SummaryBitrateResponseSummary0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryBitrateResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationResponse struct {
	Meta     RadarAttackLayer7SummaryDurationResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryDurationResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryDurationResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryDurationResponse]
type radarAttackLayer7SummaryDurationResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryDurationResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryDurationResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryDurationResponseMeta]
type radarAttackLayer7SummaryDurationResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryDurationResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryDurationResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryDurationResponseMetaDateRange]
type radarAttackLayer7SummaryDurationResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryDurationResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAttackLayer7SummaryDurationResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryDurationResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryDurationResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationResponseSummary0 struct {
	Number1HourTo3Hours  string                                               `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins string                                               `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins string                                               `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  string                                               `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           string                                               `json:"OVER_3_HOURS,required"`
	Under10Mins          string                                               `json:"UNDER_10_MINS,required"`
	JSON                 radarAttackLayer7SummaryDurationResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryDurationResponseSummary0]
type radarAttackLayer7SummaryDurationResponseSummary0JSON struct {
	Number1HourTo3Hours  apijson.Field
	Number10MinsTo20Mins apijson.Field
	Number20MinsTo40Mins apijson.Field
	Number40MinsTo1Hour  apijson.Field
	Over3Hours           apijson.Field
	Under10Mins          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryDurationResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetResponse struct {
	Meta     RadarAttackLayer7SummaryGetResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryGetResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryGetResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryGetResponse]
type radarAttackLayer7SummaryGetResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryGetResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarAttackLayer7SummaryGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryGetResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryGetResponseMeta]
type radarAttackLayer7SummaryGetResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryGetResponseMetaDateRange]
type radarAttackLayer7SummaryGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarAttackLayer7SummaryGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetResponseSummary0 struct {
	Gre  string                                          `json:"gre,required"`
	Icmp string                                          `json:"icmp,required"`
	Tcp  string                                          `json:"tcp,required"`
	Udp  string                                          `json:"udp,required"`
	JSON radarAttackLayer7SummaryGetResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryGetResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryGetResponseSummary0]
type radarAttackLayer7SummaryGetResponseSummary0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionResponse struct {
	Meta     RadarAttackLayer7SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryIPVersionResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryIPVersionResponse]
type radarAttackLayer7SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryIPVersionResponseMeta]
type radarAttackLayer7SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryIPVersionResponseMetaDateRange]
type radarAttackLayer7SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                                `json:"IPv4,required"`
	IPv6 string                                                `json:"IPv6,required"`
	JSON radarAttackLayer7SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryIPVersionResponseSummary0]
type radarAttackLayer7SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolResponse struct {
	Meta     RadarAttackLayer7SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryProtocolResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryProtocolResponse]
type radarAttackLayer7SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryProtocolResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryProtocolResponseMeta]
type radarAttackLayer7SummaryProtocolResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryProtocolResponseMetaDateRange]
type radarAttackLayer7SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolResponseSummary0 struct {
	Gre  string                                               `json:"GRE,required"`
	Icmp string                                               `json:"ICMP,required"`
	Tcp  string                                               `json:"TCP,required"`
	Udp  string                                               `json:"UDP,required"`
	JSON radarAttackLayer7SummaryProtocolResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryProtocolResponseSummary0]
type radarAttackLayer7SummaryProtocolResponseSummary0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryVectorResponse struct {
	Meta     RadarAttackLayer7SummaryVectorResponseMeta `json:"meta,required"`
	Summary0 map[string][]string                        `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryVectorResponseJSON `json:"-"`
}

// radarAttackLayer7SummaryVectorResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryVectorResponse]
type radarAttackLayer7SummaryVectorResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryVectorResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryVectorResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryVectorResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryVectorResponseMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryVectorResponseMeta]
type radarAttackLayer7SummaryVectorResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryVectorResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryVectorResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryVectorResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryVectorResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryVectorResponseMetaDateRange]
type radarAttackLayer7SummaryVectorResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryVectorResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAttackLayer7SummaryVectorResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryVectorResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryVectorResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryVectorResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryBitrateParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryBitrateParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7SummaryBitrateParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryBitrateParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryBitrateParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7SummaryBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7SummaryBitrateParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7SummaryBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryBitrateParamsDateRange string

const (
	RadarAttackLayer7SummaryBitrateParamsDateRange1d         RadarAttackLayer7SummaryBitrateParamsDateRange = "1d"
	RadarAttackLayer7SummaryBitrateParamsDateRange2d         RadarAttackLayer7SummaryBitrateParamsDateRange = "2d"
	RadarAttackLayer7SummaryBitrateParamsDateRange7d         RadarAttackLayer7SummaryBitrateParamsDateRange = "7d"
	RadarAttackLayer7SummaryBitrateParamsDateRange14d        RadarAttackLayer7SummaryBitrateParamsDateRange = "14d"
	RadarAttackLayer7SummaryBitrateParamsDateRange28d        RadarAttackLayer7SummaryBitrateParamsDateRange = "28d"
	RadarAttackLayer7SummaryBitrateParamsDateRange12w        RadarAttackLayer7SummaryBitrateParamsDateRange = "12w"
	RadarAttackLayer7SummaryBitrateParamsDateRange24w        RadarAttackLayer7SummaryBitrateParamsDateRange = "24w"
	RadarAttackLayer7SummaryBitrateParamsDateRange52w        RadarAttackLayer7SummaryBitrateParamsDateRange = "52w"
	RadarAttackLayer7SummaryBitrateParamsDateRange1dControl  RadarAttackLayer7SummaryBitrateParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryBitrateParamsDateRange2dControl  RadarAttackLayer7SummaryBitrateParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryBitrateParamsDateRange7dControl  RadarAttackLayer7SummaryBitrateParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryBitrateParamsDateRange14dControl RadarAttackLayer7SummaryBitrateParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryBitrateParamsDateRange28dControl RadarAttackLayer7SummaryBitrateParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryBitrateParamsDateRange12wControl RadarAttackLayer7SummaryBitrateParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryBitrateParamsDateRange24wControl RadarAttackLayer7SummaryBitrateParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7SummaryBitrateParamsDirection string

const (
	RadarAttackLayer7SummaryBitrateParamsDirectionOrigin RadarAttackLayer7SummaryBitrateParamsDirection = "ORIGIN"
	RadarAttackLayer7SummaryBitrateParamsDirectionTarget RadarAttackLayer7SummaryBitrateParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7SummaryBitrateParamsFormat string

const (
	RadarAttackLayer7SummaryBitrateParamsFormatJson RadarAttackLayer7SummaryBitrateParamsFormat = "JSON"
	RadarAttackLayer7SummaryBitrateParamsFormatCsv  RadarAttackLayer7SummaryBitrateParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryBitrateParamsIPVersion string

const (
	RadarAttackLayer7SummaryBitrateParamsIPVersionIPv4 RadarAttackLayer7SummaryBitrateParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryBitrateParamsIPVersionIPv6 RadarAttackLayer7SummaryBitrateParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryBitrateParamsProtocol string

const (
	RadarAttackLayer7SummaryBitrateParamsProtocolUdp  RadarAttackLayer7SummaryBitrateParamsProtocol = "UDP"
	RadarAttackLayer7SummaryBitrateParamsProtocolTcp  RadarAttackLayer7SummaryBitrateParamsProtocol = "TCP"
	RadarAttackLayer7SummaryBitrateParamsProtocolIcmp RadarAttackLayer7SummaryBitrateParamsProtocol = "ICMP"
	RadarAttackLayer7SummaryBitrateParamsProtocolGre  RadarAttackLayer7SummaryBitrateParamsProtocol = "GRE"
)

type RadarAttackLayer7SummaryBitrateResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryBitrateResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAttackLayer7SummaryBitrateResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryBitrateResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryBitrateResponseEnvelope]
type radarAttackLayer7SummaryBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryDurationParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryDurationParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7SummaryDurationParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryDurationParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryDurationParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7SummaryDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7SummaryDurationParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7SummaryDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryDurationParamsDateRange string

const (
	RadarAttackLayer7SummaryDurationParamsDateRange1d         RadarAttackLayer7SummaryDurationParamsDateRange = "1d"
	RadarAttackLayer7SummaryDurationParamsDateRange2d         RadarAttackLayer7SummaryDurationParamsDateRange = "2d"
	RadarAttackLayer7SummaryDurationParamsDateRange7d         RadarAttackLayer7SummaryDurationParamsDateRange = "7d"
	RadarAttackLayer7SummaryDurationParamsDateRange14d        RadarAttackLayer7SummaryDurationParamsDateRange = "14d"
	RadarAttackLayer7SummaryDurationParamsDateRange28d        RadarAttackLayer7SummaryDurationParamsDateRange = "28d"
	RadarAttackLayer7SummaryDurationParamsDateRange12w        RadarAttackLayer7SummaryDurationParamsDateRange = "12w"
	RadarAttackLayer7SummaryDurationParamsDateRange24w        RadarAttackLayer7SummaryDurationParamsDateRange = "24w"
	RadarAttackLayer7SummaryDurationParamsDateRange52w        RadarAttackLayer7SummaryDurationParamsDateRange = "52w"
	RadarAttackLayer7SummaryDurationParamsDateRange1dControl  RadarAttackLayer7SummaryDurationParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryDurationParamsDateRange2dControl  RadarAttackLayer7SummaryDurationParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryDurationParamsDateRange7dControl  RadarAttackLayer7SummaryDurationParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryDurationParamsDateRange14dControl RadarAttackLayer7SummaryDurationParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryDurationParamsDateRange28dControl RadarAttackLayer7SummaryDurationParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryDurationParamsDateRange12wControl RadarAttackLayer7SummaryDurationParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryDurationParamsDateRange24wControl RadarAttackLayer7SummaryDurationParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7SummaryDurationParamsDirection string

const (
	RadarAttackLayer7SummaryDurationParamsDirectionOrigin RadarAttackLayer7SummaryDurationParamsDirection = "ORIGIN"
	RadarAttackLayer7SummaryDurationParamsDirectionTarget RadarAttackLayer7SummaryDurationParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7SummaryDurationParamsFormat string

const (
	RadarAttackLayer7SummaryDurationParamsFormatJson RadarAttackLayer7SummaryDurationParamsFormat = "JSON"
	RadarAttackLayer7SummaryDurationParamsFormatCsv  RadarAttackLayer7SummaryDurationParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryDurationParamsIPVersion string

const (
	RadarAttackLayer7SummaryDurationParamsIPVersionIPv4 RadarAttackLayer7SummaryDurationParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryDurationParamsIPVersionIPv6 RadarAttackLayer7SummaryDurationParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryDurationParamsProtocol string

const (
	RadarAttackLayer7SummaryDurationParamsProtocolUdp  RadarAttackLayer7SummaryDurationParamsProtocol = "UDP"
	RadarAttackLayer7SummaryDurationParamsProtocolTcp  RadarAttackLayer7SummaryDurationParamsProtocol = "TCP"
	RadarAttackLayer7SummaryDurationParamsProtocolIcmp RadarAttackLayer7SummaryDurationParamsProtocol = "ICMP"
	RadarAttackLayer7SummaryDurationParamsProtocolGre  RadarAttackLayer7SummaryDurationParamsProtocol = "GRE"
)

type RadarAttackLayer7SummaryDurationResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryDurationResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7SummaryDurationResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryDurationResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryDurationResponseEnvelope]
type radarAttackLayer7SummaryDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7SummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryGetParamsDateRange string

const (
	RadarAttackLayer7SummaryGetParamsDateRange1d         RadarAttackLayer7SummaryGetParamsDateRange = "1d"
	RadarAttackLayer7SummaryGetParamsDateRange2d         RadarAttackLayer7SummaryGetParamsDateRange = "2d"
	RadarAttackLayer7SummaryGetParamsDateRange7d         RadarAttackLayer7SummaryGetParamsDateRange = "7d"
	RadarAttackLayer7SummaryGetParamsDateRange14d        RadarAttackLayer7SummaryGetParamsDateRange = "14d"
	RadarAttackLayer7SummaryGetParamsDateRange28d        RadarAttackLayer7SummaryGetParamsDateRange = "28d"
	RadarAttackLayer7SummaryGetParamsDateRange12w        RadarAttackLayer7SummaryGetParamsDateRange = "12w"
	RadarAttackLayer7SummaryGetParamsDateRange24w        RadarAttackLayer7SummaryGetParamsDateRange = "24w"
	RadarAttackLayer7SummaryGetParamsDateRange52w        RadarAttackLayer7SummaryGetParamsDateRange = "52w"
	RadarAttackLayer7SummaryGetParamsDateRange1dControl  RadarAttackLayer7SummaryGetParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryGetParamsDateRange2dControl  RadarAttackLayer7SummaryGetParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryGetParamsDateRange7dControl  RadarAttackLayer7SummaryGetParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryGetParamsDateRange14dControl RadarAttackLayer7SummaryGetParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryGetParamsDateRange28dControl RadarAttackLayer7SummaryGetParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryGetParamsDateRange12wControl RadarAttackLayer7SummaryGetParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryGetParamsDateRange24wControl RadarAttackLayer7SummaryGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryGetParamsFormat string

const (
	RadarAttackLayer7SummaryGetParamsFormatJson RadarAttackLayer7SummaryGetParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetParamsFormatCsv  RadarAttackLayer7SummaryGetParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryGetResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryGetResponseEnvelope]
type radarAttackLayer7SummaryGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7SummaryIPVersionParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7SummaryIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7SummaryIPVersionParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryIPVersionParamsDateRange string

const (
	RadarAttackLayer7SummaryIPVersionParamsDateRange1d         RadarAttackLayer7SummaryIPVersionParamsDateRange = "1d"
	RadarAttackLayer7SummaryIPVersionParamsDateRange2d         RadarAttackLayer7SummaryIPVersionParamsDateRange = "2d"
	RadarAttackLayer7SummaryIPVersionParamsDateRange7d         RadarAttackLayer7SummaryIPVersionParamsDateRange = "7d"
	RadarAttackLayer7SummaryIPVersionParamsDateRange14d        RadarAttackLayer7SummaryIPVersionParamsDateRange = "14d"
	RadarAttackLayer7SummaryIPVersionParamsDateRange28d        RadarAttackLayer7SummaryIPVersionParamsDateRange = "28d"
	RadarAttackLayer7SummaryIPVersionParamsDateRange12w        RadarAttackLayer7SummaryIPVersionParamsDateRange = "12w"
	RadarAttackLayer7SummaryIPVersionParamsDateRange24w        RadarAttackLayer7SummaryIPVersionParamsDateRange = "24w"
	RadarAttackLayer7SummaryIPVersionParamsDateRange52w        RadarAttackLayer7SummaryIPVersionParamsDateRange = "52w"
	RadarAttackLayer7SummaryIPVersionParamsDateRange1dControl  RadarAttackLayer7SummaryIPVersionParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryIPVersionParamsDateRange2dControl  RadarAttackLayer7SummaryIPVersionParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryIPVersionParamsDateRange7dControl  RadarAttackLayer7SummaryIPVersionParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryIPVersionParamsDateRange14dControl RadarAttackLayer7SummaryIPVersionParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryIPVersionParamsDateRange28dControl RadarAttackLayer7SummaryIPVersionParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryIPVersionParamsDateRange12wControl RadarAttackLayer7SummaryIPVersionParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryIPVersionParamsDateRange24wControl RadarAttackLayer7SummaryIPVersionParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7SummaryIPVersionParamsDirection string

const (
	RadarAttackLayer7SummaryIPVersionParamsDirectionOrigin RadarAttackLayer7SummaryIPVersionParamsDirection = "ORIGIN"
	RadarAttackLayer7SummaryIPVersionParamsDirectionTarget RadarAttackLayer7SummaryIPVersionParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7SummaryIPVersionParamsFormat string

const (
	RadarAttackLayer7SummaryIPVersionParamsFormatJson RadarAttackLayer7SummaryIPVersionParamsFormat = "JSON"
	RadarAttackLayer7SummaryIPVersionParamsFormatCsv  RadarAttackLayer7SummaryIPVersionParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryIPVersionParamsProtocol string

const (
	RadarAttackLayer7SummaryIPVersionParamsProtocolUdp  RadarAttackLayer7SummaryIPVersionParamsProtocol = "UDP"
	RadarAttackLayer7SummaryIPVersionParamsProtocolTcp  RadarAttackLayer7SummaryIPVersionParamsProtocol = "TCP"
	RadarAttackLayer7SummaryIPVersionParamsProtocolIcmp RadarAttackLayer7SummaryIPVersionParamsProtocol = "ICMP"
	RadarAttackLayer7SummaryIPVersionParamsProtocolGre  RadarAttackLayer7SummaryIPVersionParamsProtocol = "GRE"
)

type RadarAttackLayer7SummaryIPVersionResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAttackLayer7SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryIPVersionResponseEnvelope]
type radarAttackLayer7SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryProtocolParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7SummaryProtocolParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryProtocolParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryProtocolParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryProtocolParamsDateRange string

const (
	RadarAttackLayer7SummaryProtocolParamsDateRange1d         RadarAttackLayer7SummaryProtocolParamsDateRange = "1d"
	RadarAttackLayer7SummaryProtocolParamsDateRange2d         RadarAttackLayer7SummaryProtocolParamsDateRange = "2d"
	RadarAttackLayer7SummaryProtocolParamsDateRange7d         RadarAttackLayer7SummaryProtocolParamsDateRange = "7d"
	RadarAttackLayer7SummaryProtocolParamsDateRange14d        RadarAttackLayer7SummaryProtocolParamsDateRange = "14d"
	RadarAttackLayer7SummaryProtocolParamsDateRange28d        RadarAttackLayer7SummaryProtocolParamsDateRange = "28d"
	RadarAttackLayer7SummaryProtocolParamsDateRange12w        RadarAttackLayer7SummaryProtocolParamsDateRange = "12w"
	RadarAttackLayer7SummaryProtocolParamsDateRange24w        RadarAttackLayer7SummaryProtocolParamsDateRange = "24w"
	RadarAttackLayer7SummaryProtocolParamsDateRange52w        RadarAttackLayer7SummaryProtocolParamsDateRange = "52w"
	RadarAttackLayer7SummaryProtocolParamsDateRange1dControl  RadarAttackLayer7SummaryProtocolParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryProtocolParamsDateRange2dControl  RadarAttackLayer7SummaryProtocolParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryProtocolParamsDateRange7dControl  RadarAttackLayer7SummaryProtocolParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryProtocolParamsDateRange14dControl RadarAttackLayer7SummaryProtocolParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryProtocolParamsDateRange28dControl RadarAttackLayer7SummaryProtocolParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryProtocolParamsDateRange12wControl RadarAttackLayer7SummaryProtocolParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryProtocolParamsDateRange24wControl RadarAttackLayer7SummaryProtocolParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7SummaryProtocolParamsDirection string

const (
	RadarAttackLayer7SummaryProtocolParamsDirectionOrigin RadarAttackLayer7SummaryProtocolParamsDirection = "ORIGIN"
	RadarAttackLayer7SummaryProtocolParamsDirectionTarget RadarAttackLayer7SummaryProtocolParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7SummaryProtocolParamsFormat string

const (
	RadarAttackLayer7SummaryProtocolParamsFormatJson RadarAttackLayer7SummaryProtocolParamsFormat = "JSON"
	RadarAttackLayer7SummaryProtocolParamsFormatCsv  RadarAttackLayer7SummaryProtocolParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryProtocolParamsIPVersion string

const (
	RadarAttackLayer7SummaryProtocolParamsIPVersionIPv4 RadarAttackLayer7SummaryProtocolParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryProtocolParamsIPVersionIPv6 RadarAttackLayer7SummaryProtocolParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryProtocolResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryProtocolResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryProtocolResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryProtocolResponseEnvelope]
type radarAttackLayer7SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryVectorParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryVectorParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7SummaryVectorParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryVectorParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryVectorParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7SummaryVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7SummaryVectorParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7SummaryVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryVectorParamsDateRange string

const (
	RadarAttackLayer7SummaryVectorParamsDateRange1d         RadarAttackLayer7SummaryVectorParamsDateRange = "1d"
	RadarAttackLayer7SummaryVectorParamsDateRange2d         RadarAttackLayer7SummaryVectorParamsDateRange = "2d"
	RadarAttackLayer7SummaryVectorParamsDateRange7d         RadarAttackLayer7SummaryVectorParamsDateRange = "7d"
	RadarAttackLayer7SummaryVectorParamsDateRange14d        RadarAttackLayer7SummaryVectorParamsDateRange = "14d"
	RadarAttackLayer7SummaryVectorParamsDateRange28d        RadarAttackLayer7SummaryVectorParamsDateRange = "28d"
	RadarAttackLayer7SummaryVectorParamsDateRange12w        RadarAttackLayer7SummaryVectorParamsDateRange = "12w"
	RadarAttackLayer7SummaryVectorParamsDateRange24w        RadarAttackLayer7SummaryVectorParamsDateRange = "24w"
	RadarAttackLayer7SummaryVectorParamsDateRange52w        RadarAttackLayer7SummaryVectorParamsDateRange = "52w"
	RadarAttackLayer7SummaryVectorParamsDateRange1dControl  RadarAttackLayer7SummaryVectorParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryVectorParamsDateRange2dControl  RadarAttackLayer7SummaryVectorParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryVectorParamsDateRange7dControl  RadarAttackLayer7SummaryVectorParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryVectorParamsDateRange14dControl RadarAttackLayer7SummaryVectorParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryVectorParamsDateRange28dControl RadarAttackLayer7SummaryVectorParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryVectorParamsDateRange12wControl RadarAttackLayer7SummaryVectorParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryVectorParamsDateRange24wControl RadarAttackLayer7SummaryVectorParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7SummaryVectorParamsDirection string

const (
	RadarAttackLayer7SummaryVectorParamsDirectionOrigin RadarAttackLayer7SummaryVectorParamsDirection = "ORIGIN"
	RadarAttackLayer7SummaryVectorParamsDirectionTarget RadarAttackLayer7SummaryVectorParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7SummaryVectorParamsFormat string

const (
	RadarAttackLayer7SummaryVectorParamsFormatJson RadarAttackLayer7SummaryVectorParamsFormat = "JSON"
	RadarAttackLayer7SummaryVectorParamsFormatCsv  RadarAttackLayer7SummaryVectorParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryVectorParamsIPVersion string

const (
	RadarAttackLayer7SummaryVectorParamsIPVersionIPv4 RadarAttackLayer7SummaryVectorParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryVectorParamsIPVersionIPv6 RadarAttackLayer7SummaryVectorParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryVectorParamsProtocol string

const (
	RadarAttackLayer7SummaryVectorParamsProtocolUdp  RadarAttackLayer7SummaryVectorParamsProtocol = "UDP"
	RadarAttackLayer7SummaryVectorParamsProtocolTcp  RadarAttackLayer7SummaryVectorParamsProtocol = "TCP"
	RadarAttackLayer7SummaryVectorParamsProtocolIcmp RadarAttackLayer7SummaryVectorParamsProtocol = "ICMP"
	RadarAttackLayer7SummaryVectorParamsProtocolGre  RadarAttackLayer7SummaryVectorParamsProtocol = "GRE"
)

type RadarAttackLayer7SummaryVectorResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryVectorResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAttackLayer7SummaryVectorResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryVectorResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryVectorResponseEnvelope]
type radarAttackLayer7SummaryVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
