// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// AttackLayer3SummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3SummaryService] method instead.
type AttackLayer3SummaryService struct {
	Options []option.RequestOption
}

// NewAttackLayer3SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer3SummaryService(opts ...option.RequestOption) (r *AttackLayer3SummaryService) {
	r = &AttackLayer3SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate.
func (r *AttackLayer3SummaryService) Bitrate(ctx context.Context, query AttackLayer3SummaryBitrateParams, opts ...option.RequestOption) (res *AttackLayer3SummaryBitrateResponse, err error) {
	var env AttackLayer3SummaryBitrateResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by duration.
func (r *AttackLayer3SummaryService) Duration(ctx context.Context, query AttackLayer3SummaryDurationParams, opts ...option.RequestOption) (res *AttackLayer3SummaryDurationResponse, err error) {
	var env AttackLayer3SummaryDurationResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of network protocols in Layer 3/4 attacks over a given
// time period.
func (r *AttackLayer3SummaryService) Get(ctx context.Context, query AttackLayer3SummaryGetParams, opts ...option.RequestOption) (res *AttackLayer3SummaryGetResponse, err error) {
	var env AttackLayer3SummaryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used.
func (r *AttackLayer3SummaryService) IPVersion(ctx context.Context, query AttackLayer3SummaryIPVersionParams, opts ...option.RequestOption) (res *AttackLayer3SummaryIPVersionResponse, err error) {
	var env AttackLayer3SummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by protocol used.
func (r *AttackLayer3SummaryService) Protocol(ctx context.Context, query AttackLayer3SummaryProtocolParams, opts ...option.RequestOption) (res *AttackLayer3SummaryProtocolResponse, err error) {
	var env AttackLayer3SummaryProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by vector.
func (r *AttackLayer3SummaryService) Vector(ctx context.Context, query AttackLayer3SummaryVectorParams, opts ...option.RequestOption) (res *AttackLayer3SummaryVectorResponse, err error) {
	var env AttackLayer3SummaryVectorResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3SummaryBitrateResponse struct {
	Meta     AttackLayer3SummaryBitrateResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryBitrateResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryBitrateResponseJSON     `json:"-"`
}

// attackLayer3SummaryBitrateResponseJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryBitrateResponse]
type attackLayer3SummaryBitrateResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseMeta struct {
	DateRange      []AttackLayer3SummaryBitrateResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                               `json:"lastUpdated,required"`
	Normalization  string                                               `json:"normalization,required"`
	ConfidenceInfo AttackLayer3SummaryBitrateResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3SummaryBitrateResponseMetaJSON           `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryBitrateResponseMeta]
type attackLayer3SummaryBitrateResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryBitrateResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryBitrateResponseMetaDateRange]
type attackLayer3SummaryBitrateResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryBitrateResponseMetaConfidenceInfo]
type attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseSummary0 struct {
	OneGBPSToTenGBPS         string                                         `json:"_1_GBPS_TO_10_GBPS,required"`
	TenGBPSToOneHundredGBPS  string                                         `json:"_10_GBPS_TO_100_GBPS,required"`
	FiveHundredMBPSToOneGBPS string                                         `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100GBPS              string                                         `json:"OVER_100_GBPS,required"`
	Under500MBPS             string                                         `json:"UNDER_500_MBPS,required"`
	JSON                     attackLayer3SummaryBitrateResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryBitrateResponseSummary0]
type attackLayer3SummaryBitrateResponseSummary0JSON struct {
	OneGBPSToTenGBPS         apijson.Field
	TenGBPSToOneHundredGBPS  apijson.Field
	FiveHundredMBPSToOneGBPS apijson.Field
	Over100GBPS              apijson.Field
	Under500MBPS             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponse struct {
	Meta     AttackLayer3SummaryDurationResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryDurationResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryDurationResponseJSON     `json:"-"`
}

// attackLayer3SummaryDurationResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryDurationResponse]
type attackLayer3SummaryDurationResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseMeta struct {
	DateRange      []AttackLayer3SummaryDurationResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	Normalization  string                                                `json:"normalization,required"`
	ConfidenceInfo AttackLayer3SummaryDurationResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3SummaryDurationResponseMetaJSON           `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryDurationResponseMeta]
type attackLayer3SummaryDurationResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryDurationResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryDurationResponseMetaDateRange]
type attackLayer3SummaryDurationResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryDurationResponseMetaConfidenceInfo]
type attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseSummary0 struct {
	OneHourToThreeHours     string                                          `json:"_1_HOUR_TO_3_HOURS,required"`
	TenMinsToTwentyMinds    string                                          `json:"_10_MINS_TO_20_MINS,required"`
	TwentyMindsToFortyMinds string                                          `json:"_20_MINS_TO_40_MINS,required"`
	FortyMinsToOneHour      string                                          `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours              string                                          `json:"OVER_3_HOURS,required"`
	Under10Mins             string                                          `json:"UNDER_10_MINS,required"`
	JSON                    attackLayer3SummaryDurationResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryDurationResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryDurationResponseSummary0]
type attackLayer3SummaryDurationResponseSummary0JSON struct {
	OneHourToThreeHours     apijson.Field
	TenMinsToTwentyMinds    apijson.Field
	TwentyMindsToFortyMinds apijson.Field
	FortyMinsToOneHour      apijson.Field
	Over3Hours              apijson.Field
	Under10Mins             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetResponse struct {
	Meta     AttackLayer3SummaryGetResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryGetResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryGetResponseJSON     `json:"-"`
}

// attackLayer3SummaryGetResponseJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryGetResponse]
type attackLayer3SummaryGetResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetResponseMeta struct {
	DateRange      []AttackLayer3SummaryGetResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo AttackLayer3SummaryGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3SummaryGetResponseMetaJSON           `json:"-"`
}

// attackLayer3SummaryGetResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryGetResponseMeta]
type attackLayer3SummaryGetResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryGetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryGetResponseMetaDateRange]
type attackLayer3SummaryGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        attackLayer3SummaryGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3SummaryGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryGetResponseMetaConfidenceInfo]
type attackLayer3SummaryGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            attackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetResponseSummary0 struct {
	GRE  string                                     `json:"gre,required"`
	Icmp string                                     `json:"icmp,required"`
	TCP  string                                     `json:"tcp,required"`
	Udp  string                                     `json:"udp,required"`
	JSON attackLayer3SummaryGetResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryGetResponseSummary0JSON contains the JSON metadata for the
// struct [AttackLayer3SummaryGetResponseSummary0]
type attackLayer3SummaryGetResponseSummary0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	TCP         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponse struct {
	Meta     AttackLayer3SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryIPVersionResponseJSON     `json:"-"`
}

// attackLayer3SummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryIPVersionResponse]
type attackLayer3SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseMeta struct {
	DateRange      []AttackLayer3SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryIPVersionResponseMeta]
type attackLayer3SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryIPVersionResponseMetaDateRange]
type attackLayer3SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo]
type attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                           `json:"IPv4,required"`
	IPv6 string                                           `json:"IPv6,required"`
	JSON attackLayer3SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIPVersionResponseSummary0]
type attackLayer3SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponse struct {
	Meta     AttackLayer3SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryProtocolResponseJSON     `json:"-"`
}

// attackLayer3SummaryProtocolResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryProtocolResponse]
type attackLayer3SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseMeta struct {
	DateRange      []AttackLayer3SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	Normalization  string                                                `json:"normalization,required"`
	ConfidenceInfo AttackLayer3SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3SummaryProtocolResponseMetaJSON           `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryProtocolResponseMeta]
type attackLayer3SummaryProtocolResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryProtocolResponseMetaDateRange]
type attackLayer3SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryProtocolResponseMetaConfidenceInfo]
type attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseSummary0 struct {
	GRE  string                                          `json:"GRE,required"`
	Icmp string                                          `json:"ICMP,required"`
	TCP  string                                          `json:"TCP,required"`
	Udp  string                                          `json:"UDP,required"`
	JSON attackLayer3SummaryProtocolResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryProtocolResponseSummary0]
type attackLayer3SummaryProtocolResponseSummary0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	TCP         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponse struct {
	Meta     AttackLayer3SummaryVectorResponseMeta `json:"meta,required"`
	Summary0 map[string][]string                   `json:"summary_0,required"`
	JSON     attackLayer3SummaryVectorResponseJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryVectorResponse]
type attackLayer3SummaryVectorResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponseMeta struct {
	DateRange      []AttackLayer3SummaryVectorResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	Normalization  string                                              `json:"normalization,required"`
	ConfidenceInfo AttackLayer3SummaryVectorResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3SummaryVectorResponseMetaJSON           `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVectorResponseMeta]
type attackLayer3SummaryVectorResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryVectorResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryVectorResponseMetaDateRange]
type attackLayer3SummaryVectorResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryVectorResponseMetaConfidenceInfo]
type attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3SummaryBitrateParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3SummaryBitrateParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3SummaryBitrateParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3SummaryBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryBitrateParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3SummaryBitrateParamsDirection string

const (
	AttackLayer3SummaryBitrateParamsDirectionOrigin AttackLayer3SummaryBitrateParamsDirection = "ORIGIN"
	AttackLayer3SummaryBitrateParamsDirectionTarget AttackLayer3SummaryBitrateParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryBitrateParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsDirectionOrigin, AttackLayer3SummaryBitrateParamsDirectionTarget:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3SummaryBitrateParamsFormat string

const (
	AttackLayer3SummaryBitrateParamsFormatJson AttackLayer3SummaryBitrateParamsFormat = "JSON"
	AttackLayer3SummaryBitrateParamsFormatCsv  AttackLayer3SummaryBitrateParamsFormat = "CSV"
)

func (r AttackLayer3SummaryBitrateParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsFormatJson, AttackLayer3SummaryBitrateParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateParamsIPVersion string

const (
	AttackLayer3SummaryBitrateParamsIPVersionIPv4 AttackLayer3SummaryBitrateParamsIPVersion = "IPv4"
	AttackLayer3SummaryBitrateParamsIPVersionIPv6 AttackLayer3SummaryBitrateParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryBitrateParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsIPVersionIPv4, AttackLayer3SummaryBitrateParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateParamsProtocol string

const (
	AttackLayer3SummaryBitrateParamsProtocolUdp  AttackLayer3SummaryBitrateParamsProtocol = "UDP"
	AttackLayer3SummaryBitrateParamsProtocolTCP  AttackLayer3SummaryBitrateParamsProtocol = "TCP"
	AttackLayer3SummaryBitrateParamsProtocolIcmp AttackLayer3SummaryBitrateParamsProtocol = "ICMP"
	AttackLayer3SummaryBitrateParamsProtocolGRE  AttackLayer3SummaryBitrateParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryBitrateParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsProtocolUdp, AttackLayer3SummaryBitrateParamsProtocolTCP, AttackLayer3SummaryBitrateParamsProtocolIcmp, AttackLayer3SummaryBitrateParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateResponseEnvelope struct {
	Result  AttackLayer3SummaryBitrateResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    attackLayer3SummaryBitrateResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryBitrateResponseEnvelope]
type attackLayer3SummaryBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3SummaryDurationParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3SummaryDurationParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3SummaryDurationParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3SummaryDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryDurationParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3SummaryDurationParamsDirection string

const (
	AttackLayer3SummaryDurationParamsDirectionOrigin AttackLayer3SummaryDurationParamsDirection = "ORIGIN"
	AttackLayer3SummaryDurationParamsDirectionTarget AttackLayer3SummaryDurationParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryDurationParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsDirectionOrigin, AttackLayer3SummaryDurationParamsDirectionTarget:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3SummaryDurationParamsFormat string

const (
	AttackLayer3SummaryDurationParamsFormatJson AttackLayer3SummaryDurationParamsFormat = "JSON"
	AttackLayer3SummaryDurationParamsFormatCsv  AttackLayer3SummaryDurationParamsFormat = "CSV"
)

func (r AttackLayer3SummaryDurationParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsFormatJson, AttackLayer3SummaryDurationParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationParamsIPVersion string

const (
	AttackLayer3SummaryDurationParamsIPVersionIPv4 AttackLayer3SummaryDurationParamsIPVersion = "IPv4"
	AttackLayer3SummaryDurationParamsIPVersionIPv6 AttackLayer3SummaryDurationParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryDurationParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsIPVersionIPv4, AttackLayer3SummaryDurationParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationParamsProtocol string

const (
	AttackLayer3SummaryDurationParamsProtocolUdp  AttackLayer3SummaryDurationParamsProtocol = "UDP"
	AttackLayer3SummaryDurationParamsProtocolTCP  AttackLayer3SummaryDurationParamsProtocol = "TCP"
	AttackLayer3SummaryDurationParamsProtocolIcmp AttackLayer3SummaryDurationParamsProtocol = "ICMP"
	AttackLayer3SummaryDurationParamsProtocolGRE  AttackLayer3SummaryDurationParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryDurationParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsProtocolUdp, AttackLayer3SummaryDurationParamsProtocolTCP, AttackLayer3SummaryDurationParamsProtocolIcmp, AttackLayer3SummaryDurationParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationResponseEnvelope struct {
	Result  AttackLayer3SummaryDurationResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer3SummaryDurationResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryDurationResponseEnvelope]
type attackLayer3SummaryDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryGetParams struct {
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
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3SummaryGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer3SummaryGetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer3SummaryGetParamsFormat string

const (
	AttackLayer3SummaryGetParamsFormatJson AttackLayer3SummaryGetParamsFormat = "JSON"
	AttackLayer3SummaryGetParamsFormatCsv  AttackLayer3SummaryGetParamsFormat = "CSV"
)

func (r AttackLayer3SummaryGetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryGetParamsFormatJson, AttackLayer3SummaryGetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryGetResponseEnvelope struct {
	Result  AttackLayer3SummaryGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer3SummaryGetResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryGetResponseEnvelope]
type attackLayer3SummaryGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3SummaryIPVersionParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3SummaryIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3SummaryIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3SummaryIPVersionParamsDirection string

const (
	AttackLayer3SummaryIPVersionParamsDirectionOrigin AttackLayer3SummaryIPVersionParamsDirection = "ORIGIN"
	AttackLayer3SummaryIPVersionParamsDirectionTarget AttackLayer3SummaryIPVersionParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryIPVersionParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionParamsDirectionOrigin, AttackLayer3SummaryIPVersionParamsDirectionTarget:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3SummaryIPVersionParamsFormat string

const (
	AttackLayer3SummaryIPVersionParamsFormatJson AttackLayer3SummaryIPVersionParamsFormat = "JSON"
	AttackLayer3SummaryIPVersionParamsFormatCsv  AttackLayer3SummaryIPVersionParamsFormat = "CSV"
)

func (r AttackLayer3SummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionParamsFormatJson, AttackLayer3SummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryIPVersionParamsProtocol string

const (
	AttackLayer3SummaryIPVersionParamsProtocolUdp  AttackLayer3SummaryIPVersionParamsProtocol = "UDP"
	AttackLayer3SummaryIPVersionParamsProtocolTCP  AttackLayer3SummaryIPVersionParamsProtocol = "TCP"
	AttackLayer3SummaryIPVersionParamsProtocolIcmp AttackLayer3SummaryIPVersionParamsProtocol = "ICMP"
	AttackLayer3SummaryIPVersionParamsProtocolGRE  AttackLayer3SummaryIPVersionParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionParamsProtocolUdp, AttackLayer3SummaryIPVersionParamsProtocolTCP, AttackLayer3SummaryIPVersionParamsProtocolIcmp, AttackLayer3SummaryIPVersionParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryIPVersionResponseEnvelope struct {
	Result  AttackLayer3SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    attackLayer3SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIPVersionResponseEnvelope]
type attackLayer3SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3SummaryProtocolParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3SummaryProtocolParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3SummaryProtocolParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer3SummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3SummaryProtocolParamsDirection string

const (
	AttackLayer3SummaryProtocolParamsDirectionOrigin AttackLayer3SummaryProtocolParamsDirection = "ORIGIN"
	AttackLayer3SummaryProtocolParamsDirectionTarget AttackLayer3SummaryProtocolParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryProtocolParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolParamsDirectionOrigin, AttackLayer3SummaryProtocolParamsDirectionTarget:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3SummaryProtocolParamsFormat string

const (
	AttackLayer3SummaryProtocolParamsFormatJson AttackLayer3SummaryProtocolParamsFormat = "JSON"
	AttackLayer3SummaryProtocolParamsFormatCsv  AttackLayer3SummaryProtocolParamsFormat = "CSV"
)

func (r AttackLayer3SummaryProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolParamsFormatJson, AttackLayer3SummaryProtocolParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryProtocolParamsIPVersion string

const (
	AttackLayer3SummaryProtocolParamsIPVersionIPv4 AttackLayer3SummaryProtocolParamsIPVersion = "IPv4"
	AttackLayer3SummaryProtocolParamsIPVersionIPv6 AttackLayer3SummaryProtocolParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolParamsIPVersionIPv4, AttackLayer3SummaryProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryProtocolResponseEnvelope struct {
	Result  AttackLayer3SummaryProtocolResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer3SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryProtocolResponseEnvelope]
type attackLayer3SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[AttackLayer3SummaryVectorParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[AttackLayer3SummaryVectorParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3SummaryVectorParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3SummaryVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryVectorParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type AttackLayer3SummaryVectorParamsDirection string

const (
	AttackLayer3SummaryVectorParamsDirectionOrigin AttackLayer3SummaryVectorParamsDirection = "ORIGIN"
	AttackLayer3SummaryVectorParamsDirectionTarget AttackLayer3SummaryVectorParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryVectorParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsDirectionOrigin, AttackLayer3SummaryVectorParamsDirectionTarget:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3SummaryVectorParamsFormat string

const (
	AttackLayer3SummaryVectorParamsFormatJson AttackLayer3SummaryVectorParamsFormat = "JSON"
	AttackLayer3SummaryVectorParamsFormatCsv  AttackLayer3SummaryVectorParamsFormat = "CSV"
)

func (r AttackLayer3SummaryVectorParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsFormatJson, AttackLayer3SummaryVectorParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorParamsIPVersion string

const (
	AttackLayer3SummaryVectorParamsIPVersionIPv4 AttackLayer3SummaryVectorParamsIPVersion = "IPv4"
	AttackLayer3SummaryVectorParamsIPVersionIPv6 AttackLayer3SummaryVectorParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryVectorParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsIPVersionIPv4, AttackLayer3SummaryVectorParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorParamsProtocol string

const (
	AttackLayer3SummaryVectorParamsProtocolUdp  AttackLayer3SummaryVectorParamsProtocol = "UDP"
	AttackLayer3SummaryVectorParamsProtocolTCP  AttackLayer3SummaryVectorParamsProtocol = "TCP"
	AttackLayer3SummaryVectorParamsProtocolIcmp AttackLayer3SummaryVectorParamsProtocol = "ICMP"
	AttackLayer3SummaryVectorParamsProtocolGRE  AttackLayer3SummaryVectorParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryVectorParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsProtocolUdp, AttackLayer3SummaryVectorParamsProtocolTCP, AttackLayer3SummaryVectorParamsProtocolIcmp, AttackLayer3SummaryVectorParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorResponseEnvelope struct {
	Result  AttackLayer3SummaryVectorResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    attackLayer3SummaryVectorResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVectorResponseEnvelope]
type attackLayer3SummaryVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
