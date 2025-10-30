// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options          []option.RequestOption
	Top              *DNSTopService
	Summary          *DNSSummaryService
	TimeseriesGroups *DNSTimeseriesGroupService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.Top = NewDNSTopService(opts...)
	r.Summary = NewDNSSummaryService(opts...)
	r.TimeseriesGroups = NewDNSTimeseriesGroupService(opts...)
	return
}

// Retrieves the distribution of DNS queries by the specified dimension.
func (r *DNSService) SummaryV2(ctx context.Context, dimension DNSSummaryV2ParamsDimension, query DNSSummaryV2Params, opts ...option.RequestOption) (res *DNSSummaryV2Response, err error) {
	var env DNSSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/dns/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves normalized query volume to the 1.1.1.1 DNS resolver over time.
func (r *DNSService) Timeseries(ctx context.Context, query DNSTimeseriesParams, opts ...option.RequestOption) (res *DNSTimeseriesResponse, err error) {
	var env DNSTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/dns/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries grouped by dimension over time.
func (r *DNSService) TimeseriesGroupsV2(ctx context.Context, dimension DNSTimeseriesGroupsV2ParamsDimension, query DNSTimeseriesGroupsV2Params, opts ...option.RequestOption) (res *DNSTimeseriesGroupsV2Response, err error) {
	var env DNSTimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/dns/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSSummaryV2Response struct {
	// Metadata for the results.
	Meta     DNSSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string        `json:"summary_0,required"`
	JSON     dnsSummaryV2ResponseJSON `json:"-"`
}

// dnsSummaryV2ResponseJSON contains the JSON metadata for the struct
// [DNSSummaryV2Response]
type dnsSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryV2ResponseMeta struct {
	ConfidenceInfo DNSSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryV2ResponseMetaJSON   `json:"-"`
}

// dnsSummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryV2ResponseMeta]
type dnsSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                      `json:"level,required"`
	JSON  dnsSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSSummaryV2ResponseMetaConfidenceInfo]
type dnsSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                            `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSSummaryV2ResponseMetaConfidenceInfoAnnotation]
type dnsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DNSSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSSummaryV2ResponseMetaDateRange]
type dnsSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryV2ResponseMetaNormalization string

const (
	DNSSummaryV2ResponseMetaNormalizationPercentage           DNSSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryV2ResponseMetaNormalizationMin0Max              DNSSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryV2ResponseMetaNormalizationMinMax               DNSSummaryV2ResponseMetaNormalization = "MIN_MAX"
	DNSSummaryV2ResponseMetaNormalizationRawValues            DNSSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryV2ResponseMetaNormalizationPercentageChange     DNSSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryV2ResponseMetaNormalizationRollingAverage       DNSSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryV2ResponseMetaNormalizationOverlappedPercentage DNSSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryV2ResponseMetaNormalizationRatio                DNSSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryV2ResponseMetaNormalizationPercentage, DNSSummaryV2ResponseMetaNormalizationMin0Max, DNSSummaryV2ResponseMetaNormalizationMinMax, DNSSummaryV2ResponseMetaNormalizationRawValues, DNSSummaryV2ResponseMetaNormalizationPercentageChange, DNSSummaryV2ResponseMetaNormalizationRollingAverage, DNSSummaryV2ResponseMetaNormalizationOverlappedPercentage, DNSSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryV2ResponseMetaUnit struct {
	Name  string                           `json:"name,required"`
	Value string                           `json:"value,required"`
	JSON  dnsSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSSummaryV2ResponseMetaUnit]
type dnsSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponse struct {
	// Metadata for the results.
	Meta        DNSTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]DNSTimeseriesResponse `json:"-,extras"`
	JSON        dnsTimeseriesResponseJSON        `json:"-"`
}

// dnsTimeseriesResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponse]
type dnsTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    DNSTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo DNSTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  dnsTimeseriesResponseMetaJSON   `json:"-"`
}

// dnsTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMeta]
type dnsTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesResponseMetaAggInterval string

const (
	DNSTimeseriesResponseMetaAggIntervalFifteenMinutes DNSTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	DNSTimeseriesResponseMetaAggIntervalOneHour        DNSTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	DNSTimeseriesResponseMetaAggIntervalOneDay         DNSTimeseriesResponseMetaAggInterval = "ONE_DAY"
	DNSTimeseriesResponseMetaAggIntervalOneWeek        DNSTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	DNSTimeseriesResponseMetaAggIntervalOneMonth       DNSTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r DNSTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaAggIntervalFifteenMinutes, DNSTimeseriesResponseMetaAggIntervalOneHour, DNSTimeseriesResponseMetaAggIntervalOneDay, DNSTimeseriesResponseMetaAggIntervalOneWeek, DNSTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []DNSTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                       `json:"level,required"`
	JSON  dnsTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTimeseriesResponseMetaConfidenceInfo]
type dnsTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                             `json:"startDate,required" format:"date-time"`
	JSON            dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSTimeseriesResponseMetaConfidenceInfoAnnotation]
type dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      dnsTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// dnsTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaDateRange]
type dnsTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTimeseriesResponseMetaNormalization string

const (
	DNSTimeseriesResponseMetaNormalizationPercentage           DNSTimeseriesResponseMetaNormalization = "PERCENTAGE"
	DNSTimeseriesResponseMetaNormalizationMin0Max              DNSTimeseriesResponseMetaNormalization = "MIN0_MAX"
	DNSTimeseriesResponseMetaNormalizationMinMax               DNSTimeseriesResponseMetaNormalization = "MIN_MAX"
	DNSTimeseriesResponseMetaNormalizationRawValues            DNSTimeseriesResponseMetaNormalization = "RAW_VALUES"
	DNSTimeseriesResponseMetaNormalizationPercentageChange     DNSTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTimeseriesResponseMetaNormalizationRollingAverage       DNSTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTimeseriesResponseMetaNormalizationOverlappedPercentage DNSTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTimeseriesResponseMetaNormalizationRatio                DNSTimeseriesResponseMetaNormalization = "RATIO"
)

func (r DNSTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaNormalizationPercentage, DNSTimeseriesResponseMetaNormalizationMin0Max, DNSTimeseriesResponseMetaNormalizationMinMax, DNSTimeseriesResponseMetaNormalizationRawValues, DNSTimeseriesResponseMetaNormalizationPercentageChange, DNSTimeseriesResponseMetaNormalizationRollingAverage, DNSTimeseriesResponseMetaNormalizationOverlappedPercentage, DNSTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaUnit struct {
	Name  string                            `json:"name,required"`
	Value string                            `json:"value,required"`
	JSON  dnsTimeseriesResponseMetaUnitJSON `json:"-"`
}

// dnsTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaUnit]
type dnsTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   DNSTimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 DNSTimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupsV2Response]
type dnsTimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    DNSTimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo DNSTimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSTimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  dnsTimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupsV2ResponseMeta]
type dnsTimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupsV2ResponseMetaAggInterval string

const (
	DNSTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes DNSTimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneHour        DNSTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneDay         DNSTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        DNSTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       DNSTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r DNSTimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneHour, DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneDay, DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, DNSTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type DNSTimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []DNSTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupsV2ResponseMetaConfidenceInfo]
type dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      dnsTimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupsV2ResponseMetaDateRange]
type dnsTimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTimeseriesGroupsV2ResponseMetaNormalization string

const (
	DNSTimeseriesGroupsV2ResponseMetaNormalizationPercentage           DNSTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationMin0Max              DNSTimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationMinMax               DNSTimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationRawValues            DNSTimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     DNSTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       DNSTimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage DNSTimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTimeseriesGroupsV2ResponseMetaNormalizationRatio                DNSTimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r DNSTimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ResponseMetaNormalizationPercentage, DNSTimeseriesGroupsV2ResponseMetaNormalizationMin0Max, DNSTimeseriesGroupsV2ResponseMetaNormalizationMinMax, DNSTimeseriesGroupsV2ResponseMetaNormalizationRawValues, DNSTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, DNSTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, DNSTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, DNSTimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  dnsTimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupsV2ResponseMetaUnit]
type dnsTimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                             `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                     `json:"-,extras"`
	JSON        dnsTimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupsV2ResponseSerie0]
type dnsTimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryV2Params struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSSummaryV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryV2ParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryV2ParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryV2ParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryV2Params]'s query parameters as `url.Values`.
func (r DNSSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type DNSSummaryV2ParamsDimension string

const (
	DNSSummaryV2ParamsDimensionIPVersion      DNSSummaryV2ParamsDimension = "IP_VERSION"
	DNSSummaryV2ParamsDimensionCacheHit       DNSSummaryV2ParamsDimension = "CACHE_HIT"
	DNSSummaryV2ParamsDimensionDNSSEC         DNSSummaryV2ParamsDimension = "DNSSEC"
	DNSSummaryV2ParamsDimensionDNSSECAware    DNSSummaryV2ParamsDimension = "DNSSEC_AWARE"
	DNSSummaryV2ParamsDimensionDNSSECE2E      DNSSummaryV2ParamsDimension = "DNSSEC_E2E"
	DNSSummaryV2ParamsDimensionMatchingAnswer DNSSummaryV2ParamsDimension = "MATCHING_ANSWER"
	DNSSummaryV2ParamsDimensionProtocol       DNSSummaryV2ParamsDimension = "PROTOCOL"
	DNSSummaryV2ParamsDimensionQueryType      DNSSummaryV2ParamsDimension = "QUERY_TYPE"
	DNSSummaryV2ParamsDimensionResponseCode   DNSSummaryV2ParamsDimension = "RESPONSE_CODE"
	DNSSummaryV2ParamsDimensionResponseTTL    DNSSummaryV2ParamsDimension = "RESPONSE_TTL"
)

func (r DNSSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case DNSSummaryV2ParamsDimensionIPVersion, DNSSummaryV2ParamsDimensionCacheHit, DNSSummaryV2ParamsDimensionDNSSEC, DNSSummaryV2ParamsDimensionDNSSECAware, DNSSummaryV2ParamsDimensionDNSSECE2E, DNSSummaryV2ParamsDimensionMatchingAnswer, DNSSummaryV2ParamsDimensionProtocol, DNSSummaryV2ParamsDimensionQueryType, DNSSummaryV2ParamsDimensionResponseCode, DNSSummaryV2ParamsDimensionResponseTTL:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSSummaryV2ParamsFormat string

const (
	DNSSummaryV2ParamsFormatJson DNSSummaryV2ParamsFormat = "JSON"
	DNSSummaryV2ParamsFormatCsv  DNSSummaryV2ParamsFormat = "CSV"
)

func (r DNSSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryV2ParamsFormatJson, DNSSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryV2ParamsProtocol string

const (
	DNSSummaryV2ParamsProtocolUdp   DNSSummaryV2ParamsProtocol = "UDP"
	DNSSummaryV2ParamsProtocolTCP   DNSSummaryV2ParamsProtocol = "TCP"
	DNSSummaryV2ParamsProtocolHTTPS DNSSummaryV2ParamsProtocol = "HTTPS"
	DNSSummaryV2ParamsProtocolTLS   DNSSummaryV2ParamsProtocol = "TLS"
)

func (r DNSSummaryV2ParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryV2ParamsProtocolUdp, DNSSummaryV2ParamsProtocolTCP, DNSSummaryV2ParamsProtocolHTTPS, DNSSummaryV2ParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryV2ParamsQueryType string

const (
	DNSSummaryV2ParamsQueryTypeA          DNSSummaryV2ParamsQueryType = "A"
	DNSSummaryV2ParamsQueryTypeAAAA       DNSSummaryV2ParamsQueryType = "AAAA"
	DNSSummaryV2ParamsQueryTypeA6         DNSSummaryV2ParamsQueryType = "A6"
	DNSSummaryV2ParamsQueryTypeAfsdb      DNSSummaryV2ParamsQueryType = "AFSDB"
	DNSSummaryV2ParamsQueryTypeAny        DNSSummaryV2ParamsQueryType = "ANY"
	DNSSummaryV2ParamsQueryTypeApl        DNSSummaryV2ParamsQueryType = "APL"
	DNSSummaryV2ParamsQueryTypeAtma       DNSSummaryV2ParamsQueryType = "ATMA"
	DNSSummaryV2ParamsQueryTypeAXFR       DNSSummaryV2ParamsQueryType = "AXFR"
	DNSSummaryV2ParamsQueryTypeCAA        DNSSummaryV2ParamsQueryType = "CAA"
	DNSSummaryV2ParamsQueryTypeCdnskey    DNSSummaryV2ParamsQueryType = "CDNSKEY"
	DNSSummaryV2ParamsQueryTypeCds        DNSSummaryV2ParamsQueryType = "CDS"
	DNSSummaryV2ParamsQueryTypeCERT       DNSSummaryV2ParamsQueryType = "CERT"
	DNSSummaryV2ParamsQueryTypeCNAME      DNSSummaryV2ParamsQueryType = "CNAME"
	DNSSummaryV2ParamsQueryTypeCsync      DNSSummaryV2ParamsQueryType = "CSYNC"
	DNSSummaryV2ParamsQueryTypeDhcid      DNSSummaryV2ParamsQueryType = "DHCID"
	DNSSummaryV2ParamsQueryTypeDlv        DNSSummaryV2ParamsQueryType = "DLV"
	DNSSummaryV2ParamsQueryTypeDname      DNSSummaryV2ParamsQueryType = "DNAME"
	DNSSummaryV2ParamsQueryTypeDNSKEY     DNSSummaryV2ParamsQueryType = "DNSKEY"
	DNSSummaryV2ParamsQueryTypeDoa        DNSSummaryV2ParamsQueryType = "DOA"
	DNSSummaryV2ParamsQueryTypeDS         DNSSummaryV2ParamsQueryType = "DS"
	DNSSummaryV2ParamsQueryTypeEid        DNSSummaryV2ParamsQueryType = "EID"
	DNSSummaryV2ParamsQueryTypeEui48      DNSSummaryV2ParamsQueryType = "EUI48"
	DNSSummaryV2ParamsQueryTypeEui64      DNSSummaryV2ParamsQueryType = "EUI64"
	DNSSummaryV2ParamsQueryTypeGpos       DNSSummaryV2ParamsQueryType = "GPOS"
	DNSSummaryV2ParamsQueryTypeGid        DNSSummaryV2ParamsQueryType = "GID"
	DNSSummaryV2ParamsQueryTypeHinfo      DNSSummaryV2ParamsQueryType = "HINFO"
	DNSSummaryV2ParamsQueryTypeHip        DNSSummaryV2ParamsQueryType = "HIP"
	DNSSummaryV2ParamsQueryTypeHTTPS      DNSSummaryV2ParamsQueryType = "HTTPS"
	DNSSummaryV2ParamsQueryTypeIpseckey   DNSSummaryV2ParamsQueryType = "IPSECKEY"
	DNSSummaryV2ParamsQueryTypeIsdn       DNSSummaryV2ParamsQueryType = "ISDN"
	DNSSummaryV2ParamsQueryTypeIxfr       DNSSummaryV2ParamsQueryType = "IXFR"
	DNSSummaryV2ParamsQueryTypeKey        DNSSummaryV2ParamsQueryType = "KEY"
	DNSSummaryV2ParamsQueryTypeKx         DNSSummaryV2ParamsQueryType = "KX"
	DNSSummaryV2ParamsQueryTypeL32        DNSSummaryV2ParamsQueryType = "L32"
	DNSSummaryV2ParamsQueryTypeL64        DNSSummaryV2ParamsQueryType = "L64"
	DNSSummaryV2ParamsQueryTypeLOC        DNSSummaryV2ParamsQueryType = "LOC"
	DNSSummaryV2ParamsQueryTypeLp         DNSSummaryV2ParamsQueryType = "LP"
	DNSSummaryV2ParamsQueryTypeMaila      DNSSummaryV2ParamsQueryType = "MAILA"
	DNSSummaryV2ParamsQueryTypeMailb      DNSSummaryV2ParamsQueryType = "MAILB"
	DNSSummaryV2ParamsQueryTypeMB         DNSSummaryV2ParamsQueryType = "MB"
	DNSSummaryV2ParamsQueryTypeMd         DNSSummaryV2ParamsQueryType = "MD"
	DNSSummaryV2ParamsQueryTypeMf         DNSSummaryV2ParamsQueryType = "MF"
	DNSSummaryV2ParamsQueryTypeMg         DNSSummaryV2ParamsQueryType = "MG"
	DNSSummaryV2ParamsQueryTypeMinfo      DNSSummaryV2ParamsQueryType = "MINFO"
	DNSSummaryV2ParamsQueryTypeMr         DNSSummaryV2ParamsQueryType = "MR"
	DNSSummaryV2ParamsQueryTypeMX         DNSSummaryV2ParamsQueryType = "MX"
	DNSSummaryV2ParamsQueryTypeNAPTR      DNSSummaryV2ParamsQueryType = "NAPTR"
	DNSSummaryV2ParamsQueryTypeNb         DNSSummaryV2ParamsQueryType = "NB"
	DNSSummaryV2ParamsQueryTypeNbstat     DNSSummaryV2ParamsQueryType = "NBSTAT"
	DNSSummaryV2ParamsQueryTypeNid        DNSSummaryV2ParamsQueryType = "NID"
	DNSSummaryV2ParamsQueryTypeNimloc     DNSSummaryV2ParamsQueryType = "NIMLOC"
	DNSSummaryV2ParamsQueryTypeNinfo      DNSSummaryV2ParamsQueryType = "NINFO"
	DNSSummaryV2ParamsQueryTypeNS         DNSSummaryV2ParamsQueryType = "NS"
	DNSSummaryV2ParamsQueryTypeNsap       DNSSummaryV2ParamsQueryType = "NSAP"
	DNSSummaryV2ParamsQueryTypeNsec       DNSSummaryV2ParamsQueryType = "NSEC"
	DNSSummaryV2ParamsQueryTypeNsec3      DNSSummaryV2ParamsQueryType = "NSEC3"
	DNSSummaryV2ParamsQueryTypeNsec3Param DNSSummaryV2ParamsQueryType = "NSEC3PARAM"
	DNSSummaryV2ParamsQueryTypeNull       DNSSummaryV2ParamsQueryType = "NULL"
	DNSSummaryV2ParamsQueryTypeNxt        DNSSummaryV2ParamsQueryType = "NXT"
	DNSSummaryV2ParamsQueryTypeOpenpgpkey DNSSummaryV2ParamsQueryType = "OPENPGPKEY"
	DNSSummaryV2ParamsQueryTypeOpt        DNSSummaryV2ParamsQueryType = "OPT"
	DNSSummaryV2ParamsQueryTypePTR        DNSSummaryV2ParamsQueryType = "PTR"
	DNSSummaryV2ParamsQueryTypePx         DNSSummaryV2ParamsQueryType = "PX"
	DNSSummaryV2ParamsQueryTypeRkey       DNSSummaryV2ParamsQueryType = "RKEY"
	DNSSummaryV2ParamsQueryTypeRp         DNSSummaryV2ParamsQueryType = "RP"
	DNSSummaryV2ParamsQueryTypeRrsig      DNSSummaryV2ParamsQueryType = "RRSIG"
	DNSSummaryV2ParamsQueryTypeRt         DNSSummaryV2ParamsQueryType = "RT"
	DNSSummaryV2ParamsQueryTypeSig        DNSSummaryV2ParamsQueryType = "SIG"
	DNSSummaryV2ParamsQueryTypeSink       DNSSummaryV2ParamsQueryType = "SINK"
	DNSSummaryV2ParamsQueryTypeSMIMEA     DNSSummaryV2ParamsQueryType = "SMIMEA"
	DNSSummaryV2ParamsQueryTypeSOA        DNSSummaryV2ParamsQueryType = "SOA"
	DNSSummaryV2ParamsQueryTypeSPF        DNSSummaryV2ParamsQueryType = "SPF"
	DNSSummaryV2ParamsQueryTypeSRV        DNSSummaryV2ParamsQueryType = "SRV"
	DNSSummaryV2ParamsQueryTypeSSHFP      DNSSummaryV2ParamsQueryType = "SSHFP"
	DNSSummaryV2ParamsQueryTypeSVCB       DNSSummaryV2ParamsQueryType = "SVCB"
	DNSSummaryV2ParamsQueryTypeTa         DNSSummaryV2ParamsQueryType = "TA"
	DNSSummaryV2ParamsQueryTypeTalink     DNSSummaryV2ParamsQueryType = "TALINK"
	DNSSummaryV2ParamsQueryTypeTkey       DNSSummaryV2ParamsQueryType = "TKEY"
	DNSSummaryV2ParamsQueryTypeTLSA       DNSSummaryV2ParamsQueryType = "TLSA"
	DNSSummaryV2ParamsQueryTypeTSIG       DNSSummaryV2ParamsQueryType = "TSIG"
	DNSSummaryV2ParamsQueryTypeTXT        DNSSummaryV2ParamsQueryType = "TXT"
	DNSSummaryV2ParamsQueryTypeUinfo      DNSSummaryV2ParamsQueryType = "UINFO"
	DNSSummaryV2ParamsQueryTypeUID        DNSSummaryV2ParamsQueryType = "UID"
	DNSSummaryV2ParamsQueryTypeUnspec     DNSSummaryV2ParamsQueryType = "UNSPEC"
	DNSSummaryV2ParamsQueryTypeURI        DNSSummaryV2ParamsQueryType = "URI"
	DNSSummaryV2ParamsQueryTypeWks        DNSSummaryV2ParamsQueryType = "WKS"
	DNSSummaryV2ParamsQueryTypeX25        DNSSummaryV2ParamsQueryType = "X25"
	DNSSummaryV2ParamsQueryTypeZonemd     DNSSummaryV2ParamsQueryType = "ZONEMD"
)

func (r DNSSummaryV2ParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryV2ParamsQueryTypeA, DNSSummaryV2ParamsQueryTypeAAAA, DNSSummaryV2ParamsQueryTypeA6, DNSSummaryV2ParamsQueryTypeAfsdb, DNSSummaryV2ParamsQueryTypeAny, DNSSummaryV2ParamsQueryTypeApl, DNSSummaryV2ParamsQueryTypeAtma, DNSSummaryV2ParamsQueryTypeAXFR, DNSSummaryV2ParamsQueryTypeCAA, DNSSummaryV2ParamsQueryTypeCdnskey, DNSSummaryV2ParamsQueryTypeCds, DNSSummaryV2ParamsQueryTypeCERT, DNSSummaryV2ParamsQueryTypeCNAME, DNSSummaryV2ParamsQueryTypeCsync, DNSSummaryV2ParamsQueryTypeDhcid, DNSSummaryV2ParamsQueryTypeDlv, DNSSummaryV2ParamsQueryTypeDname, DNSSummaryV2ParamsQueryTypeDNSKEY, DNSSummaryV2ParamsQueryTypeDoa, DNSSummaryV2ParamsQueryTypeDS, DNSSummaryV2ParamsQueryTypeEid, DNSSummaryV2ParamsQueryTypeEui48, DNSSummaryV2ParamsQueryTypeEui64, DNSSummaryV2ParamsQueryTypeGpos, DNSSummaryV2ParamsQueryTypeGid, DNSSummaryV2ParamsQueryTypeHinfo, DNSSummaryV2ParamsQueryTypeHip, DNSSummaryV2ParamsQueryTypeHTTPS, DNSSummaryV2ParamsQueryTypeIpseckey, DNSSummaryV2ParamsQueryTypeIsdn, DNSSummaryV2ParamsQueryTypeIxfr, DNSSummaryV2ParamsQueryTypeKey, DNSSummaryV2ParamsQueryTypeKx, DNSSummaryV2ParamsQueryTypeL32, DNSSummaryV2ParamsQueryTypeL64, DNSSummaryV2ParamsQueryTypeLOC, DNSSummaryV2ParamsQueryTypeLp, DNSSummaryV2ParamsQueryTypeMaila, DNSSummaryV2ParamsQueryTypeMailb, DNSSummaryV2ParamsQueryTypeMB, DNSSummaryV2ParamsQueryTypeMd, DNSSummaryV2ParamsQueryTypeMf, DNSSummaryV2ParamsQueryTypeMg, DNSSummaryV2ParamsQueryTypeMinfo, DNSSummaryV2ParamsQueryTypeMr, DNSSummaryV2ParamsQueryTypeMX, DNSSummaryV2ParamsQueryTypeNAPTR, DNSSummaryV2ParamsQueryTypeNb, DNSSummaryV2ParamsQueryTypeNbstat, DNSSummaryV2ParamsQueryTypeNid, DNSSummaryV2ParamsQueryTypeNimloc, DNSSummaryV2ParamsQueryTypeNinfo, DNSSummaryV2ParamsQueryTypeNS, DNSSummaryV2ParamsQueryTypeNsap, DNSSummaryV2ParamsQueryTypeNsec, DNSSummaryV2ParamsQueryTypeNsec3, DNSSummaryV2ParamsQueryTypeNsec3Param, DNSSummaryV2ParamsQueryTypeNull, DNSSummaryV2ParamsQueryTypeNxt, DNSSummaryV2ParamsQueryTypeOpenpgpkey, DNSSummaryV2ParamsQueryTypeOpt, DNSSummaryV2ParamsQueryTypePTR, DNSSummaryV2ParamsQueryTypePx, DNSSummaryV2ParamsQueryTypeRkey, DNSSummaryV2ParamsQueryTypeRp, DNSSummaryV2ParamsQueryTypeRrsig, DNSSummaryV2ParamsQueryTypeRt, DNSSummaryV2ParamsQueryTypeSig, DNSSummaryV2ParamsQueryTypeSink, DNSSummaryV2ParamsQueryTypeSMIMEA, DNSSummaryV2ParamsQueryTypeSOA, DNSSummaryV2ParamsQueryTypeSPF, DNSSummaryV2ParamsQueryTypeSRV, DNSSummaryV2ParamsQueryTypeSSHFP, DNSSummaryV2ParamsQueryTypeSVCB, DNSSummaryV2ParamsQueryTypeTa, DNSSummaryV2ParamsQueryTypeTalink, DNSSummaryV2ParamsQueryTypeTkey, DNSSummaryV2ParamsQueryTypeTLSA, DNSSummaryV2ParamsQueryTypeTSIG, DNSSummaryV2ParamsQueryTypeTXT, DNSSummaryV2ParamsQueryTypeUinfo, DNSSummaryV2ParamsQueryTypeUID, DNSSummaryV2ParamsQueryTypeUnspec, DNSSummaryV2ParamsQueryTypeURI, DNSSummaryV2ParamsQueryTypeWks, DNSSummaryV2ParamsQueryTypeX25, DNSSummaryV2ParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryV2ParamsResponseCode string

const (
	DNSSummaryV2ParamsResponseCodeNoerror   DNSSummaryV2ParamsResponseCode = "NOERROR"
	DNSSummaryV2ParamsResponseCodeFormerr   DNSSummaryV2ParamsResponseCode = "FORMERR"
	DNSSummaryV2ParamsResponseCodeServfail  DNSSummaryV2ParamsResponseCode = "SERVFAIL"
	DNSSummaryV2ParamsResponseCodeNxdomain  DNSSummaryV2ParamsResponseCode = "NXDOMAIN"
	DNSSummaryV2ParamsResponseCodeNotimp    DNSSummaryV2ParamsResponseCode = "NOTIMP"
	DNSSummaryV2ParamsResponseCodeRefused   DNSSummaryV2ParamsResponseCode = "REFUSED"
	DNSSummaryV2ParamsResponseCodeYxdomain  DNSSummaryV2ParamsResponseCode = "YXDOMAIN"
	DNSSummaryV2ParamsResponseCodeYxrrset   DNSSummaryV2ParamsResponseCode = "YXRRSET"
	DNSSummaryV2ParamsResponseCodeNxrrset   DNSSummaryV2ParamsResponseCode = "NXRRSET"
	DNSSummaryV2ParamsResponseCodeNotauth   DNSSummaryV2ParamsResponseCode = "NOTAUTH"
	DNSSummaryV2ParamsResponseCodeNotzone   DNSSummaryV2ParamsResponseCode = "NOTZONE"
	DNSSummaryV2ParamsResponseCodeBadsig    DNSSummaryV2ParamsResponseCode = "BADSIG"
	DNSSummaryV2ParamsResponseCodeBadkey    DNSSummaryV2ParamsResponseCode = "BADKEY"
	DNSSummaryV2ParamsResponseCodeBadtime   DNSSummaryV2ParamsResponseCode = "BADTIME"
	DNSSummaryV2ParamsResponseCodeBadmode   DNSSummaryV2ParamsResponseCode = "BADMODE"
	DNSSummaryV2ParamsResponseCodeBadname   DNSSummaryV2ParamsResponseCode = "BADNAME"
	DNSSummaryV2ParamsResponseCodeBadalg    DNSSummaryV2ParamsResponseCode = "BADALG"
	DNSSummaryV2ParamsResponseCodeBadtrunc  DNSSummaryV2ParamsResponseCode = "BADTRUNC"
	DNSSummaryV2ParamsResponseCodeBadcookie DNSSummaryV2ParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryV2ParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryV2ParamsResponseCodeNoerror, DNSSummaryV2ParamsResponseCodeFormerr, DNSSummaryV2ParamsResponseCodeServfail, DNSSummaryV2ParamsResponseCodeNxdomain, DNSSummaryV2ParamsResponseCodeNotimp, DNSSummaryV2ParamsResponseCodeRefused, DNSSummaryV2ParamsResponseCodeYxdomain, DNSSummaryV2ParamsResponseCodeYxrrset, DNSSummaryV2ParamsResponseCodeNxrrset, DNSSummaryV2ParamsResponseCodeNotauth, DNSSummaryV2ParamsResponseCodeNotzone, DNSSummaryV2ParamsResponseCodeBadsig, DNSSummaryV2ParamsResponseCodeBadkey, DNSSummaryV2ParamsResponseCodeBadtime, DNSSummaryV2ParamsResponseCodeBadmode, DNSSummaryV2ParamsResponseCodeBadname, DNSSummaryV2ParamsResponseCodeBadalg, DNSSummaryV2ParamsResponseCodeBadtrunc, DNSSummaryV2ParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryV2ResponseEnvelope struct {
	Result  DNSSummaryV2Response             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    dnsSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryV2ResponseEnvelope]
type dnsSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesParams]'s query parameters as `url.Values`.
func (r DNSTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesParamsAggInterval string

const (
	DNSTimeseriesParamsAggInterval15m DNSTimeseriesParamsAggInterval = "15m"
	DNSTimeseriesParamsAggInterval1h  DNSTimeseriesParamsAggInterval = "1h"
	DNSTimeseriesParamsAggInterval1d  DNSTimeseriesParamsAggInterval = "1d"
	DNSTimeseriesParamsAggInterval1w  DNSTimeseriesParamsAggInterval = "1w"
)

func (r DNSTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsAggInterval15m, DNSTimeseriesParamsAggInterval1h, DNSTimeseriesParamsAggInterval1d, DNSTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesParamsFormat string

const (
	DNSTimeseriesParamsFormatJson DNSTimeseriesParamsFormat = "JSON"
	DNSTimeseriesParamsFormatCsv  DNSTimeseriesParamsFormat = "CSV"
)

func (r DNSTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsFormatJson, DNSTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesParamsProtocol string

const (
	DNSTimeseriesParamsProtocolUdp   DNSTimeseriesParamsProtocol = "UDP"
	DNSTimeseriesParamsProtocolTCP   DNSTimeseriesParamsProtocol = "TCP"
	DNSTimeseriesParamsProtocolHTTPS DNSTimeseriesParamsProtocol = "HTTPS"
	DNSTimeseriesParamsProtocolTLS   DNSTimeseriesParamsProtocol = "TLS"
)

func (r DNSTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsProtocolUdp, DNSTimeseriesParamsProtocolTCP, DNSTimeseriesParamsProtocolHTTPS, DNSTimeseriesParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesParamsQueryType string

const (
	DNSTimeseriesParamsQueryTypeA          DNSTimeseriesParamsQueryType = "A"
	DNSTimeseriesParamsQueryTypeAAAA       DNSTimeseriesParamsQueryType = "AAAA"
	DNSTimeseriesParamsQueryTypeA6         DNSTimeseriesParamsQueryType = "A6"
	DNSTimeseriesParamsQueryTypeAfsdb      DNSTimeseriesParamsQueryType = "AFSDB"
	DNSTimeseriesParamsQueryTypeAny        DNSTimeseriesParamsQueryType = "ANY"
	DNSTimeseriesParamsQueryTypeApl        DNSTimeseriesParamsQueryType = "APL"
	DNSTimeseriesParamsQueryTypeAtma       DNSTimeseriesParamsQueryType = "ATMA"
	DNSTimeseriesParamsQueryTypeAXFR       DNSTimeseriesParamsQueryType = "AXFR"
	DNSTimeseriesParamsQueryTypeCAA        DNSTimeseriesParamsQueryType = "CAA"
	DNSTimeseriesParamsQueryTypeCdnskey    DNSTimeseriesParamsQueryType = "CDNSKEY"
	DNSTimeseriesParamsQueryTypeCds        DNSTimeseriesParamsQueryType = "CDS"
	DNSTimeseriesParamsQueryTypeCERT       DNSTimeseriesParamsQueryType = "CERT"
	DNSTimeseriesParamsQueryTypeCNAME      DNSTimeseriesParamsQueryType = "CNAME"
	DNSTimeseriesParamsQueryTypeCsync      DNSTimeseriesParamsQueryType = "CSYNC"
	DNSTimeseriesParamsQueryTypeDhcid      DNSTimeseriesParamsQueryType = "DHCID"
	DNSTimeseriesParamsQueryTypeDlv        DNSTimeseriesParamsQueryType = "DLV"
	DNSTimeseriesParamsQueryTypeDname      DNSTimeseriesParamsQueryType = "DNAME"
	DNSTimeseriesParamsQueryTypeDNSKEY     DNSTimeseriesParamsQueryType = "DNSKEY"
	DNSTimeseriesParamsQueryTypeDoa        DNSTimeseriesParamsQueryType = "DOA"
	DNSTimeseriesParamsQueryTypeDS         DNSTimeseriesParamsQueryType = "DS"
	DNSTimeseriesParamsQueryTypeEid        DNSTimeseriesParamsQueryType = "EID"
	DNSTimeseriesParamsQueryTypeEui48      DNSTimeseriesParamsQueryType = "EUI48"
	DNSTimeseriesParamsQueryTypeEui64      DNSTimeseriesParamsQueryType = "EUI64"
	DNSTimeseriesParamsQueryTypeGpos       DNSTimeseriesParamsQueryType = "GPOS"
	DNSTimeseriesParamsQueryTypeGid        DNSTimeseriesParamsQueryType = "GID"
	DNSTimeseriesParamsQueryTypeHinfo      DNSTimeseriesParamsQueryType = "HINFO"
	DNSTimeseriesParamsQueryTypeHip        DNSTimeseriesParamsQueryType = "HIP"
	DNSTimeseriesParamsQueryTypeHTTPS      DNSTimeseriesParamsQueryType = "HTTPS"
	DNSTimeseriesParamsQueryTypeIpseckey   DNSTimeseriesParamsQueryType = "IPSECKEY"
	DNSTimeseriesParamsQueryTypeIsdn       DNSTimeseriesParamsQueryType = "ISDN"
	DNSTimeseriesParamsQueryTypeIxfr       DNSTimeseriesParamsQueryType = "IXFR"
	DNSTimeseriesParamsQueryTypeKey        DNSTimeseriesParamsQueryType = "KEY"
	DNSTimeseriesParamsQueryTypeKx         DNSTimeseriesParamsQueryType = "KX"
	DNSTimeseriesParamsQueryTypeL32        DNSTimeseriesParamsQueryType = "L32"
	DNSTimeseriesParamsQueryTypeL64        DNSTimeseriesParamsQueryType = "L64"
	DNSTimeseriesParamsQueryTypeLOC        DNSTimeseriesParamsQueryType = "LOC"
	DNSTimeseriesParamsQueryTypeLp         DNSTimeseriesParamsQueryType = "LP"
	DNSTimeseriesParamsQueryTypeMaila      DNSTimeseriesParamsQueryType = "MAILA"
	DNSTimeseriesParamsQueryTypeMailb      DNSTimeseriesParamsQueryType = "MAILB"
	DNSTimeseriesParamsQueryTypeMB         DNSTimeseriesParamsQueryType = "MB"
	DNSTimeseriesParamsQueryTypeMd         DNSTimeseriesParamsQueryType = "MD"
	DNSTimeseriesParamsQueryTypeMf         DNSTimeseriesParamsQueryType = "MF"
	DNSTimeseriesParamsQueryTypeMg         DNSTimeseriesParamsQueryType = "MG"
	DNSTimeseriesParamsQueryTypeMinfo      DNSTimeseriesParamsQueryType = "MINFO"
	DNSTimeseriesParamsQueryTypeMr         DNSTimeseriesParamsQueryType = "MR"
	DNSTimeseriesParamsQueryTypeMX         DNSTimeseriesParamsQueryType = "MX"
	DNSTimeseriesParamsQueryTypeNAPTR      DNSTimeseriesParamsQueryType = "NAPTR"
	DNSTimeseriesParamsQueryTypeNb         DNSTimeseriesParamsQueryType = "NB"
	DNSTimeseriesParamsQueryTypeNbstat     DNSTimeseriesParamsQueryType = "NBSTAT"
	DNSTimeseriesParamsQueryTypeNid        DNSTimeseriesParamsQueryType = "NID"
	DNSTimeseriesParamsQueryTypeNimloc     DNSTimeseriesParamsQueryType = "NIMLOC"
	DNSTimeseriesParamsQueryTypeNinfo      DNSTimeseriesParamsQueryType = "NINFO"
	DNSTimeseriesParamsQueryTypeNS         DNSTimeseriesParamsQueryType = "NS"
	DNSTimeseriesParamsQueryTypeNsap       DNSTimeseriesParamsQueryType = "NSAP"
	DNSTimeseriesParamsQueryTypeNsec       DNSTimeseriesParamsQueryType = "NSEC"
	DNSTimeseriesParamsQueryTypeNsec3      DNSTimeseriesParamsQueryType = "NSEC3"
	DNSTimeseriesParamsQueryTypeNsec3Param DNSTimeseriesParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesParamsQueryTypeNull       DNSTimeseriesParamsQueryType = "NULL"
	DNSTimeseriesParamsQueryTypeNxt        DNSTimeseriesParamsQueryType = "NXT"
	DNSTimeseriesParamsQueryTypeOpenpgpkey DNSTimeseriesParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesParamsQueryTypeOpt        DNSTimeseriesParamsQueryType = "OPT"
	DNSTimeseriesParamsQueryTypePTR        DNSTimeseriesParamsQueryType = "PTR"
	DNSTimeseriesParamsQueryTypePx         DNSTimeseriesParamsQueryType = "PX"
	DNSTimeseriesParamsQueryTypeRkey       DNSTimeseriesParamsQueryType = "RKEY"
	DNSTimeseriesParamsQueryTypeRp         DNSTimeseriesParamsQueryType = "RP"
	DNSTimeseriesParamsQueryTypeRrsig      DNSTimeseriesParamsQueryType = "RRSIG"
	DNSTimeseriesParamsQueryTypeRt         DNSTimeseriesParamsQueryType = "RT"
	DNSTimeseriesParamsQueryTypeSig        DNSTimeseriesParamsQueryType = "SIG"
	DNSTimeseriesParamsQueryTypeSink       DNSTimeseriesParamsQueryType = "SINK"
	DNSTimeseriesParamsQueryTypeSMIMEA     DNSTimeseriesParamsQueryType = "SMIMEA"
	DNSTimeseriesParamsQueryTypeSOA        DNSTimeseriesParamsQueryType = "SOA"
	DNSTimeseriesParamsQueryTypeSPF        DNSTimeseriesParamsQueryType = "SPF"
	DNSTimeseriesParamsQueryTypeSRV        DNSTimeseriesParamsQueryType = "SRV"
	DNSTimeseriesParamsQueryTypeSSHFP      DNSTimeseriesParamsQueryType = "SSHFP"
	DNSTimeseriesParamsQueryTypeSVCB       DNSTimeseriesParamsQueryType = "SVCB"
	DNSTimeseriesParamsQueryTypeTa         DNSTimeseriesParamsQueryType = "TA"
	DNSTimeseriesParamsQueryTypeTalink     DNSTimeseriesParamsQueryType = "TALINK"
	DNSTimeseriesParamsQueryTypeTkey       DNSTimeseriesParamsQueryType = "TKEY"
	DNSTimeseriesParamsQueryTypeTLSA       DNSTimeseriesParamsQueryType = "TLSA"
	DNSTimeseriesParamsQueryTypeTSIG       DNSTimeseriesParamsQueryType = "TSIG"
	DNSTimeseriesParamsQueryTypeTXT        DNSTimeseriesParamsQueryType = "TXT"
	DNSTimeseriesParamsQueryTypeUinfo      DNSTimeseriesParamsQueryType = "UINFO"
	DNSTimeseriesParamsQueryTypeUID        DNSTimeseriesParamsQueryType = "UID"
	DNSTimeseriesParamsQueryTypeUnspec     DNSTimeseriesParamsQueryType = "UNSPEC"
	DNSTimeseriesParamsQueryTypeURI        DNSTimeseriesParamsQueryType = "URI"
	DNSTimeseriesParamsQueryTypeWks        DNSTimeseriesParamsQueryType = "WKS"
	DNSTimeseriesParamsQueryTypeX25        DNSTimeseriesParamsQueryType = "X25"
	DNSTimeseriesParamsQueryTypeZonemd     DNSTimeseriesParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsQueryTypeA, DNSTimeseriesParamsQueryTypeAAAA, DNSTimeseriesParamsQueryTypeA6, DNSTimeseriesParamsQueryTypeAfsdb, DNSTimeseriesParamsQueryTypeAny, DNSTimeseriesParamsQueryTypeApl, DNSTimeseriesParamsQueryTypeAtma, DNSTimeseriesParamsQueryTypeAXFR, DNSTimeseriesParamsQueryTypeCAA, DNSTimeseriesParamsQueryTypeCdnskey, DNSTimeseriesParamsQueryTypeCds, DNSTimeseriesParamsQueryTypeCERT, DNSTimeseriesParamsQueryTypeCNAME, DNSTimeseriesParamsQueryTypeCsync, DNSTimeseriesParamsQueryTypeDhcid, DNSTimeseriesParamsQueryTypeDlv, DNSTimeseriesParamsQueryTypeDname, DNSTimeseriesParamsQueryTypeDNSKEY, DNSTimeseriesParamsQueryTypeDoa, DNSTimeseriesParamsQueryTypeDS, DNSTimeseriesParamsQueryTypeEid, DNSTimeseriesParamsQueryTypeEui48, DNSTimeseriesParamsQueryTypeEui64, DNSTimeseriesParamsQueryTypeGpos, DNSTimeseriesParamsQueryTypeGid, DNSTimeseriesParamsQueryTypeHinfo, DNSTimeseriesParamsQueryTypeHip, DNSTimeseriesParamsQueryTypeHTTPS, DNSTimeseriesParamsQueryTypeIpseckey, DNSTimeseriesParamsQueryTypeIsdn, DNSTimeseriesParamsQueryTypeIxfr, DNSTimeseriesParamsQueryTypeKey, DNSTimeseriesParamsQueryTypeKx, DNSTimeseriesParamsQueryTypeL32, DNSTimeseriesParamsQueryTypeL64, DNSTimeseriesParamsQueryTypeLOC, DNSTimeseriesParamsQueryTypeLp, DNSTimeseriesParamsQueryTypeMaila, DNSTimeseriesParamsQueryTypeMailb, DNSTimeseriesParamsQueryTypeMB, DNSTimeseriesParamsQueryTypeMd, DNSTimeseriesParamsQueryTypeMf, DNSTimeseriesParamsQueryTypeMg, DNSTimeseriesParamsQueryTypeMinfo, DNSTimeseriesParamsQueryTypeMr, DNSTimeseriesParamsQueryTypeMX, DNSTimeseriesParamsQueryTypeNAPTR, DNSTimeseriesParamsQueryTypeNb, DNSTimeseriesParamsQueryTypeNbstat, DNSTimeseriesParamsQueryTypeNid, DNSTimeseriesParamsQueryTypeNimloc, DNSTimeseriesParamsQueryTypeNinfo, DNSTimeseriesParamsQueryTypeNS, DNSTimeseriesParamsQueryTypeNsap, DNSTimeseriesParamsQueryTypeNsec, DNSTimeseriesParamsQueryTypeNsec3, DNSTimeseriesParamsQueryTypeNsec3Param, DNSTimeseriesParamsQueryTypeNull, DNSTimeseriesParamsQueryTypeNxt, DNSTimeseriesParamsQueryTypeOpenpgpkey, DNSTimeseriesParamsQueryTypeOpt, DNSTimeseriesParamsQueryTypePTR, DNSTimeseriesParamsQueryTypePx, DNSTimeseriesParamsQueryTypeRkey, DNSTimeseriesParamsQueryTypeRp, DNSTimeseriesParamsQueryTypeRrsig, DNSTimeseriesParamsQueryTypeRt, DNSTimeseriesParamsQueryTypeSig, DNSTimeseriesParamsQueryTypeSink, DNSTimeseriesParamsQueryTypeSMIMEA, DNSTimeseriesParamsQueryTypeSOA, DNSTimeseriesParamsQueryTypeSPF, DNSTimeseriesParamsQueryTypeSRV, DNSTimeseriesParamsQueryTypeSSHFP, DNSTimeseriesParamsQueryTypeSVCB, DNSTimeseriesParamsQueryTypeTa, DNSTimeseriesParamsQueryTypeTalink, DNSTimeseriesParamsQueryTypeTkey, DNSTimeseriesParamsQueryTypeTLSA, DNSTimeseriesParamsQueryTypeTSIG, DNSTimeseriesParamsQueryTypeTXT, DNSTimeseriesParamsQueryTypeUinfo, DNSTimeseriesParamsQueryTypeUID, DNSTimeseriesParamsQueryTypeUnspec, DNSTimeseriesParamsQueryTypeURI, DNSTimeseriesParamsQueryTypeWks, DNSTimeseriesParamsQueryTypeX25, DNSTimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesParamsResponseCode string

const (
	DNSTimeseriesParamsResponseCodeNoerror   DNSTimeseriesParamsResponseCode = "NOERROR"
	DNSTimeseriesParamsResponseCodeFormerr   DNSTimeseriesParamsResponseCode = "FORMERR"
	DNSTimeseriesParamsResponseCodeServfail  DNSTimeseriesParamsResponseCode = "SERVFAIL"
	DNSTimeseriesParamsResponseCodeNxdomain  DNSTimeseriesParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesParamsResponseCodeNotimp    DNSTimeseriesParamsResponseCode = "NOTIMP"
	DNSTimeseriesParamsResponseCodeRefused   DNSTimeseriesParamsResponseCode = "REFUSED"
	DNSTimeseriesParamsResponseCodeYxdomain  DNSTimeseriesParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesParamsResponseCodeYxrrset   DNSTimeseriesParamsResponseCode = "YXRRSET"
	DNSTimeseriesParamsResponseCodeNxrrset   DNSTimeseriesParamsResponseCode = "NXRRSET"
	DNSTimeseriesParamsResponseCodeNotauth   DNSTimeseriesParamsResponseCode = "NOTAUTH"
	DNSTimeseriesParamsResponseCodeNotzone   DNSTimeseriesParamsResponseCode = "NOTZONE"
	DNSTimeseriesParamsResponseCodeBadsig    DNSTimeseriesParamsResponseCode = "BADSIG"
	DNSTimeseriesParamsResponseCodeBadkey    DNSTimeseriesParamsResponseCode = "BADKEY"
	DNSTimeseriesParamsResponseCodeBadtime   DNSTimeseriesParamsResponseCode = "BADTIME"
	DNSTimeseriesParamsResponseCodeBadmode   DNSTimeseriesParamsResponseCode = "BADMODE"
	DNSTimeseriesParamsResponseCodeBadname   DNSTimeseriesParamsResponseCode = "BADNAME"
	DNSTimeseriesParamsResponseCodeBadalg    DNSTimeseriesParamsResponseCode = "BADALG"
	DNSTimeseriesParamsResponseCodeBadtrunc  DNSTimeseriesParamsResponseCode = "BADTRUNC"
	DNSTimeseriesParamsResponseCodeBadcookie DNSTimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsResponseCodeNoerror, DNSTimeseriesParamsResponseCodeFormerr, DNSTimeseriesParamsResponseCodeServfail, DNSTimeseriesParamsResponseCodeNxdomain, DNSTimeseriesParamsResponseCodeNotimp, DNSTimeseriesParamsResponseCodeRefused, DNSTimeseriesParamsResponseCodeYxdomain, DNSTimeseriesParamsResponseCodeYxrrset, DNSTimeseriesParamsResponseCodeNxrrset, DNSTimeseriesParamsResponseCodeNotauth, DNSTimeseriesParamsResponseCodeNotzone, DNSTimeseriesParamsResponseCodeBadsig, DNSTimeseriesParamsResponseCodeBadkey, DNSTimeseriesParamsResponseCodeBadtime, DNSTimeseriesParamsResponseCodeBadmode, DNSTimeseriesParamsResponseCodeBadname, DNSTimeseriesParamsResponseCodeBadalg, DNSTimeseriesParamsResponseCodeBadtrunc, DNSTimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesResponseEnvelope struct {
	Result  DNSTimeseriesResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    dnsTimeseriesResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseEnvelope]
type dnsTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupsV2ParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupsV2ParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupsV2ParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type DNSTimeseriesGroupsV2ParamsDimension string

const (
	DNSTimeseriesGroupsV2ParamsDimensionIPVersion      DNSTimeseriesGroupsV2ParamsDimension = "IP_VERSION"
	DNSTimeseriesGroupsV2ParamsDimensionCacheHit       DNSTimeseriesGroupsV2ParamsDimension = "CACHE_HIT"
	DNSTimeseriesGroupsV2ParamsDimensionDNSSEC         DNSTimeseriesGroupsV2ParamsDimension = "DNSSEC"
	DNSTimeseriesGroupsV2ParamsDimensionDNSSECAware    DNSTimeseriesGroupsV2ParamsDimension = "DNSSEC_AWARE"
	DNSTimeseriesGroupsV2ParamsDimensionDNSSECE2E      DNSTimeseriesGroupsV2ParamsDimension = "DNSSEC_E2E"
	DNSTimeseriesGroupsV2ParamsDimensionMatchingAnswer DNSTimeseriesGroupsV2ParamsDimension = "MATCHING_ANSWER"
	DNSTimeseriesGroupsV2ParamsDimensionProtocol       DNSTimeseriesGroupsV2ParamsDimension = "PROTOCOL"
	DNSTimeseriesGroupsV2ParamsDimensionQueryType      DNSTimeseriesGroupsV2ParamsDimension = "QUERY_TYPE"
	DNSTimeseriesGroupsV2ParamsDimensionResponseCode   DNSTimeseriesGroupsV2ParamsDimension = "RESPONSE_CODE"
	DNSTimeseriesGroupsV2ParamsDimensionResponseTTL    DNSTimeseriesGroupsV2ParamsDimension = "RESPONSE_TTL"
)

func (r DNSTimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ParamsDimensionIPVersion, DNSTimeseriesGroupsV2ParamsDimensionCacheHit, DNSTimeseriesGroupsV2ParamsDimensionDNSSEC, DNSTimeseriesGroupsV2ParamsDimensionDNSSECAware, DNSTimeseriesGroupsV2ParamsDimensionDNSSECE2E, DNSTimeseriesGroupsV2ParamsDimensionMatchingAnswer, DNSTimeseriesGroupsV2ParamsDimensionProtocol, DNSTimeseriesGroupsV2ParamsDimensionQueryType, DNSTimeseriesGroupsV2ParamsDimensionResponseCode, DNSTimeseriesGroupsV2ParamsDimensionResponseTTL:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupsV2ParamsAggInterval string

const (
	DNSTimeseriesGroupsV2ParamsAggInterval15m DNSTimeseriesGroupsV2ParamsAggInterval = "15m"
	DNSTimeseriesGroupsV2ParamsAggInterval1h  DNSTimeseriesGroupsV2ParamsAggInterval = "1h"
	DNSTimeseriesGroupsV2ParamsAggInterval1d  DNSTimeseriesGroupsV2ParamsAggInterval = "1d"
	DNSTimeseriesGroupsV2ParamsAggInterval1w  DNSTimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ParamsAggInterval15m, DNSTimeseriesGroupsV2ParamsAggInterval1h, DNSTimeseriesGroupsV2ParamsAggInterval1d, DNSTimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupsV2ParamsFormat string

const (
	DNSTimeseriesGroupsV2ParamsFormatJson DNSTimeseriesGroupsV2ParamsFormat = "JSON"
	DNSTimeseriesGroupsV2ParamsFormatCsv  DNSTimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ParamsFormatJson, DNSTimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupsV2ParamsProtocol string

const (
	DNSTimeseriesGroupsV2ParamsProtocolUdp   DNSTimeseriesGroupsV2ParamsProtocol = "UDP"
	DNSTimeseriesGroupsV2ParamsProtocolTCP   DNSTimeseriesGroupsV2ParamsProtocol = "TCP"
	DNSTimeseriesGroupsV2ParamsProtocolHTTPS DNSTimeseriesGroupsV2ParamsProtocol = "HTTPS"
	DNSTimeseriesGroupsV2ParamsProtocolTLS   DNSTimeseriesGroupsV2ParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupsV2ParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ParamsProtocolUdp, DNSTimeseriesGroupsV2ParamsProtocolTCP, DNSTimeseriesGroupsV2ParamsProtocolHTTPS, DNSTimeseriesGroupsV2ParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupsV2ParamsQueryType string

const (
	DNSTimeseriesGroupsV2ParamsQueryTypeA          DNSTimeseriesGroupsV2ParamsQueryType = "A"
	DNSTimeseriesGroupsV2ParamsQueryTypeAAAA       DNSTimeseriesGroupsV2ParamsQueryType = "AAAA"
	DNSTimeseriesGroupsV2ParamsQueryTypeA6         DNSTimeseriesGroupsV2ParamsQueryType = "A6"
	DNSTimeseriesGroupsV2ParamsQueryTypeAfsdb      DNSTimeseriesGroupsV2ParamsQueryType = "AFSDB"
	DNSTimeseriesGroupsV2ParamsQueryTypeAny        DNSTimeseriesGroupsV2ParamsQueryType = "ANY"
	DNSTimeseriesGroupsV2ParamsQueryTypeApl        DNSTimeseriesGroupsV2ParamsQueryType = "APL"
	DNSTimeseriesGroupsV2ParamsQueryTypeAtma       DNSTimeseriesGroupsV2ParamsQueryType = "ATMA"
	DNSTimeseriesGroupsV2ParamsQueryTypeAXFR       DNSTimeseriesGroupsV2ParamsQueryType = "AXFR"
	DNSTimeseriesGroupsV2ParamsQueryTypeCAA        DNSTimeseriesGroupsV2ParamsQueryType = "CAA"
	DNSTimeseriesGroupsV2ParamsQueryTypeCdnskey    DNSTimeseriesGroupsV2ParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeCds        DNSTimeseriesGroupsV2ParamsQueryType = "CDS"
	DNSTimeseriesGroupsV2ParamsQueryTypeCERT       DNSTimeseriesGroupsV2ParamsQueryType = "CERT"
	DNSTimeseriesGroupsV2ParamsQueryTypeCNAME      DNSTimeseriesGroupsV2ParamsQueryType = "CNAME"
	DNSTimeseriesGroupsV2ParamsQueryTypeCsync      DNSTimeseriesGroupsV2ParamsQueryType = "CSYNC"
	DNSTimeseriesGroupsV2ParamsQueryTypeDhcid      DNSTimeseriesGroupsV2ParamsQueryType = "DHCID"
	DNSTimeseriesGroupsV2ParamsQueryTypeDlv        DNSTimeseriesGroupsV2ParamsQueryType = "DLV"
	DNSTimeseriesGroupsV2ParamsQueryTypeDname      DNSTimeseriesGroupsV2ParamsQueryType = "DNAME"
	DNSTimeseriesGroupsV2ParamsQueryTypeDNSKEY     DNSTimeseriesGroupsV2ParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeDoa        DNSTimeseriesGroupsV2ParamsQueryType = "DOA"
	DNSTimeseriesGroupsV2ParamsQueryTypeDS         DNSTimeseriesGroupsV2ParamsQueryType = "DS"
	DNSTimeseriesGroupsV2ParamsQueryTypeEid        DNSTimeseriesGroupsV2ParamsQueryType = "EID"
	DNSTimeseriesGroupsV2ParamsQueryTypeEui48      DNSTimeseriesGroupsV2ParamsQueryType = "EUI48"
	DNSTimeseriesGroupsV2ParamsQueryTypeEui64      DNSTimeseriesGroupsV2ParamsQueryType = "EUI64"
	DNSTimeseriesGroupsV2ParamsQueryTypeGpos       DNSTimeseriesGroupsV2ParamsQueryType = "GPOS"
	DNSTimeseriesGroupsV2ParamsQueryTypeGid        DNSTimeseriesGroupsV2ParamsQueryType = "GID"
	DNSTimeseriesGroupsV2ParamsQueryTypeHinfo      DNSTimeseriesGroupsV2ParamsQueryType = "HINFO"
	DNSTimeseriesGroupsV2ParamsQueryTypeHip        DNSTimeseriesGroupsV2ParamsQueryType = "HIP"
	DNSTimeseriesGroupsV2ParamsQueryTypeHTTPS      DNSTimeseriesGroupsV2ParamsQueryType = "HTTPS"
	DNSTimeseriesGroupsV2ParamsQueryTypeIpseckey   DNSTimeseriesGroupsV2ParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeIsdn       DNSTimeseriesGroupsV2ParamsQueryType = "ISDN"
	DNSTimeseriesGroupsV2ParamsQueryTypeIxfr       DNSTimeseriesGroupsV2ParamsQueryType = "IXFR"
	DNSTimeseriesGroupsV2ParamsQueryTypeKey        DNSTimeseriesGroupsV2ParamsQueryType = "KEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeKx         DNSTimeseriesGroupsV2ParamsQueryType = "KX"
	DNSTimeseriesGroupsV2ParamsQueryTypeL32        DNSTimeseriesGroupsV2ParamsQueryType = "L32"
	DNSTimeseriesGroupsV2ParamsQueryTypeL64        DNSTimeseriesGroupsV2ParamsQueryType = "L64"
	DNSTimeseriesGroupsV2ParamsQueryTypeLOC        DNSTimeseriesGroupsV2ParamsQueryType = "LOC"
	DNSTimeseriesGroupsV2ParamsQueryTypeLp         DNSTimeseriesGroupsV2ParamsQueryType = "LP"
	DNSTimeseriesGroupsV2ParamsQueryTypeMaila      DNSTimeseriesGroupsV2ParamsQueryType = "MAILA"
	DNSTimeseriesGroupsV2ParamsQueryTypeMailb      DNSTimeseriesGroupsV2ParamsQueryType = "MAILB"
	DNSTimeseriesGroupsV2ParamsQueryTypeMB         DNSTimeseriesGroupsV2ParamsQueryType = "MB"
	DNSTimeseriesGroupsV2ParamsQueryTypeMd         DNSTimeseriesGroupsV2ParamsQueryType = "MD"
	DNSTimeseriesGroupsV2ParamsQueryTypeMf         DNSTimeseriesGroupsV2ParamsQueryType = "MF"
	DNSTimeseriesGroupsV2ParamsQueryTypeMg         DNSTimeseriesGroupsV2ParamsQueryType = "MG"
	DNSTimeseriesGroupsV2ParamsQueryTypeMinfo      DNSTimeseriesGroupsV2ParamsQueryType = "MINFO"
	DNSTimeseriesGroupsV2ParamsQueryTypeMr         DNSTimeseriesGroupsV2ParamsQueryType = "MR"
	DNSTimeseriesGroupsV2ParamsQueryTypeMX         DNSTimeseriesGroupsV2ParamsQueryType = "MX"
	DNSTimeseriesGroupsV2ParamsQueryTypeNAPTR      DNSTimeseriesGroupsV2ParamsQueryType = "NAPTR"
	DNSTimeseriesGroupsV2ParamsQueryTypeNb         DNSTimeseriesGroupsV2ParamsQueryType = "NB"
	DNSTimeseriesGroupsV2ParamsQueryTypeNbstat     DNSTimeseriesGroupsV2ParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupsV2ParamsQueryTypeNid        DNSTimeseriesGroupsV2ParamsQueryType = "NID"
	DNSTimeseriesGroupsV2ParamsQueryTypeNimloc     DNSTimeseriesGroupsV2ParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupsV2ParamsQueryTypeNinfo      DNSTimeseriesGroupsV2ParamsQueryType = "NINFO"
	DNSTimeseriesGroupsV2ParamsQueryTypeNS         DNSTimeseriesGroupsV2ParamsQueryType = "NS"
	DNSTimeseriesGroupsV2ParamsQueryTypeNsap       DNSTimeseriesGroupsV2ParamsQueryType = "NSAP"
	DNSTimeseriesGroupsV2ParamsQueryTypeNsec       DNSTimeseriesGroupsV2ParamsQueryType = "NSEC"
	DNSTimeseriesGroupsV2ParamsQueryTypeNsec3      DNSTimeseriesGroupsV2ParamsQueryType = "NSEC3"
	DNSTimeseriesGroupsV2ParamsQueryTypeNsec3Param DNSTimeseriesGroupsV2ParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupsV2ParamsQueryTypeNull       DNSTimeseriesGroupsV2ParamsQueryType = "NULL"
	DNSTimeseriesGroupsV2ParamsQueryTypeNxt        DNSTimeseriesGroupsV2ParamsQueryType = "NXT"
	DNSTimeseriesGroupsV2ParamsQueryTypeOpenpgpkey DNSTimeseriesGroupsV2ParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeOpt        DNSTimeseriesGroupsV2ParamsQueryType = "OPT"
	DNSTimeseriesGroupsV2ParamsQueryTypePTR        DNSTimeseriesGroupsV2ParamsQueryType = "PTR"
	DNSTimeseriesGroupsV2ParamsQueryTypePx         DNSTimeseriesGroupsV2ParamsQueryType = "PX"
	DNSTimeseriesGroupsV2ParamsQueryTypeRkey       DNSTimeseriesGroupsV2ParamsQueryType = "RKEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeRp         DNSTimeseriesGroupsV2ParamsQueryType = "RP"
	DNSTimeseriesGroupsV2ParamsQueryTypeRrsig      DNSTimeseriesGroupsV2ParamsQueryType = "RRSIG"
	DNSTimeseriesGroupsV2ParamsQueryTypeRt         DNSTimeseriesGroupsV2ParamsQueryType = "RT"
	DNSTimeseriesGroupsV2ParamsQueryTypeSig        DNSTimeseriesGroupsV2ParamsQueryType = "SIG"
	DNSTimeseriesGroupsV2ParamsQueryTypeSink       DNSTimeseriesGroupsV2ParamsQueryType = "SINK"
	DNSTimeseriesGroupsV2ParamsQueryTypeSMIMEA     DNSTimeseriesGroupsV2ParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupsV2ParamsQueryTypeSOA        DNSTimeseriesGroupsV2ParamsQueryType = "SOA"
	DNSTimeseriesGroupsV2ParamsQueryTypeSPF        DNSTimeseriesGroupsV2ParamsQueryType = "SPF"
	DNSTimeseriesGroupsV2ParamsQueryTypeSRV        DNSTimeseriesGroupsV2ParamsQueryType = "SRV"
	DNSTimeseriesGroupsV2ParamsQueryTypeSSHFP      DNSTimeseriesGroupsV2ParamsQueryType = "SSHFP"
	DNSTimeseriesGroupsV2ParamsQueryTypeSVCB       DNSTimeseriesGroupsV2ParamsQueryType = "SVCB"
	DNSTimeseriesGroupsV2ParamsQueryTypeTa         DNSTimeseriesGroupsV2ParamsQueryType = "TA"
	DNSTimeseriesGroupsV2ParamsQueryTypeTalink     DNSTimeseriesGroupsV2ParamsQueryType = "TALINK"
	DNSTimeseriesGroupsV2ParamsQueryTypeTkey       DNSTimeseriesGroupsV2ParamsQueryType = "TKEY"
	DNSTimeseriesGroupsV2ParamsQueryTypeTLSA       DNSTimeseriesGroupsV2ParamsQueryType = "TLSA"
	DNSTimeseriesGroupsV2ParamsQueryTypeTSIG       DNSTimeseriesGroupsV2ParamsQueryType = "TSIG"
	DNSTimeseriesGroupsV2ParamsQueryTypeTXT        DNSTimeseriesGroupsV2ParamsQueryType = "TXT"
	DNSTimeseriesGroupsV2ParamsQueryTypeUinfo      DNSTimeseriesGroupsV2ParamsQueryType = "UINFO"
	DNSTimeseriesGroupsV2ParamsQueryTypeUID        DNSTimeseriesGroupsV2ParamsQueryType = "UID"
	DNSTimeseriesGroupsV2ParamsQueryTypeUnspec     DNSTimeseriesGroupsV2ParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupsV2ParamsQueryTypeURI        DNSTimeseriesGroupsV2ParamsQueryType = "URI"
	DNSTimeseriesGroupsV2ParamsQueryTypeWks        DNSTimeseriesGroupsV2ParamsQueryType = "WKS"
	DNSTimeseriesGroupsV2ParamsQueryTypeX25        DNSTimeseriesGroupsV2ParamsQueryType = "X25"
	DNSTimeseriesGroupsV2ParamsQueryTypeZonemd     DNSTimeseriesGroupsV2ParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupsV2ParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ParamsQueryTypeA, DNSTimeseriesGroupsV2ParamsQueryTypeAAAA, DNSTimeseriesGroupsV2ParamsQueryTypeA6, DNSTimeseriesGroupsV2ParamsQueryTypeAfsdb, DNSTimeseriesGroupsV2ParamsQueryTypeAny, DNSTimeseriesGroupsV2ParamsQueryTypeApl, DNSTimeseriesGroupsV2ParamsQueryTypeAtma, DNSTimeseriesGroupsV2ParamsQueryTypeAXFR, DNSTimeseriesGroupsV2ParamsQueryTypeCAA, DNSTimeseriesGroupsV2ParamsQueryTypeCdnskey, DNSTimeseriesGroupsV2ParamsQueryTypeCds, DNSTimeseriesGroupsV2ParamsQueryTypeCERT, DNSTimeseriesGroupsV2ParamsQueryTypeCNAME, DNSTimeseriesGroupsV2ParamsQueryTypeCsync, DNSTimeseriesGroupsV2ParamsQueryTypeDhcid, DNSTimeseriesGroupsV2ParamsQueryTypeDlv, DNSTimeseriesGroupsV2ParamsQueryTypeDname, DNSTimeseriesGroupsV2ParamsQueryTypeDNSKEY, DNSTimeseriesGroupsV2ParamsQueryTypeDoa, DNSTimeseriesGroupsV2ParamsQueryTypeDS, DNSTimeseriesGroupsV2ParamsQueryTypeEid, DNSTimeseriesGroupsV2ParamsQueryTypeEui48, DNSTimeseriesGroupsV2ParamsQueryTypeEui64, DNSTimeseriesGroupsV2ParamsQueryTypeGpos, DNSTimeseriesGroupsV2ParamsQueryTypeGid, DNSTimeseriesGroupsV2ParamsQueryTypeHinfo, DNSTimeseriesGroupsV2ParamsQueryTypeHip, DNSTimeseriesGroupsV2ParamsQueryTypeHTTPS, DNSTimeseriesGroupsV2ParamsQueryTypeIpseckey, DNSTimeseriesGroupsV2ParamsQueryTypeIsdn, DNSTimeseriesGroupsV2ParamsQueryTypeIxfr, DNSTimeseriesGroupsV2ParamsQueryTypeKey, DNSTimeseriesGroupsV2ParamsQueryTypeKx, DNSTimeseriesGroupsV2ParamsQueryTypeL32, DNSTimeseriesGroupsV2ParamsQueryTypeL64, DNSTimeseriesGroupsV2ParamsQueryTypeLOC, DNSTimeseriesGroupsV2ParamsQueryTypeLp, DNSTimeseriesGroupsV2ParamsQueryTypeMaila, DNSTimeseriesGroupsV2ParamsQueryTypeMailb, DNSTimeseriesGroupsV2ParamsQueryTypeMB, DNSTimeseriesGroupsV2ParamsQueryTypeMd, DNSTimeseriesGroupsV2ParamsQueryTypeMf, DNSTimeseriesGroupsV2ParamsQueryTypeMg, DNSTimeseriesGroupsV2ParamsQueryTypeMinfo, DNSTimeseriesGroupsV2ParamsQueryTypeMr, DNSTimeseriesGroupsV2ParamsQueryTypeMX, DNSTimeseriesGroupsV2ParamsQueryTypeNAPTR, DNSTimeseriesGroupsV2ParamsQueryTypeNb, DNSTimeseriesGroupsV2ParamsQueryTypeNbstat, DNSTimeseriesGroupsV2ParamsQueryTypeNid, DNSTimeseriesGroupsV2ParamsQueryTypeNimloc, DNSTimeseriesGroupsV2ParamsQueryTypeNinfo, DNSTimeseriesGroupsV2ParamsQueryTypeNS, DNSTimeseriesGroupsV2ParamsQueryTypeNsap, DNSTimeseriesGroupsV2ParamsQueryTypeNsec, DNSTimeseriesGroupsV2ParamsQueryTypeNsec3, DNSTimeseriesGroupsV2ParamsQueryTypeNsec3Param, DNSTimeseriesGroupsV2ParamsQueryTypeNull, DNSTimeseriesGroupsV2ParamsQueryTypeNxt, DNSTimeseriesGroupsV2ParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupsV2ParamsQueryTypeOpt, DNSTimeseriesGroupsV2ParamsQueryTypePTR, DNSTimeseriesGroupsV2ParamsQueryTypePx, DNSTimeseriesGroupsV2ParamsQueryTypeRkey, DNSTimeseriesGroupsV2ParamsQueryTypeRp, DNSTimeseriesGroupsV2ParamsQueryTypeRrsig, DNSTimeseriesGroupsV2ParamsQueryTypeRt, DNSTimeseriesGroupsV2ParamsQueryTypeSig, DNSTimeseriesGroupsV2ParamsQueryTypeSink, DNSTimeseriesGroupsV2ParamsQueryTypeSMIMEA, DNSTimeseriesGroupsV2ParamsQueryTypeSOA, DNSTimeseriesGroupsV2ParamsQueryTypeSPF, DNSTimeseriesGroupsV2ParamsQueryTypeSRV, DNSTimeseriesGroupsV2ParamsQueryTypeSSHFP, DNSTimeseriesGroupsV2ParamsQueryTypeSVCB, DNSTimeseriesGroupsV2ParamsQueryTypeTa, DNSTimeseriesGroupsV2ParamsQueryTypeTalink, DNSTimeseriesGroupsV2ParamsQueryTypeTkey, DNSTimeseriesGroupsV2ParamsQueryTypeTLSA, DNSTimeseriesGroupsV2ParamsQueryTypeTSIG, DNSTimeseriesGroupsV2ParamsQueryTypeTXT, DNSTimeseriesGroupsV2ParamsQueryTypeUinfo, DNSTimeseriesGroupsV2ParamsQueryTypeUID, DNSTimeseriesGroupsV2ParamsQueryTypeUnspec, DNSTimeseriesGroupsV2ParamsQueryTypeURI, DNSTimeseriesGroupsV2ParamsQueryTypeWks, DNSTimeseriesGroupsV2ParamsQueryTypeX25, DNSTimeseriesGroupsV2ParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupsV2ParamsResponseCode string

const (
	DNSTimeseriesGroupsV2ParamsResponseCodeNoerror   DNSTimeseriesGroupsV2ParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupsV2ParamsResponseCodeFormerr   DNSTimeseriesGroupsV2ParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupsV2ParamsResponseCodeServfail  DNSTimeseriesGroupsV2ParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupsV2ParamsResponseCodeNxdomain  DNSTimeseriesGroupsV2ParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupsV2ParamsResponseCodeNotimp    DNSTimeseriesGroupsV2ParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupsV2ParamsResponseCodeRefused   DNSTimeseriesGroupsV2ParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupsV2ParamsResponseCodeYxdomain  DNSTimeseriesGroupsV2ParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupsV2ParamsResponseCodeYxrrset   DNSTimeseriesGroupsV2ParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupsV2ParamsResponseCodeNxrrset   DNSTimeseriesGroupsV2ParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupsV2ParamsResponseCodeNotauth   DNSTimeseriesGroupsV2ParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupsV2ParamsResponseCodeNotzone   DNSTimeseriesGroupsV2ParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadsig    DNSTimeseriesGroupsV2ParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadkey    DNSTimeseriesGroupsV2ParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadtime   DNSTimeseriesGroupsV2ParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadmode   DNSTimeseriesGroupsV2ParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadname   DNSTimeseriesGroupsV2ParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadalg    DNSTimeseriesGroupsV2ParamsResponseCode = "BADALG"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadtrunc  DNSTimeseriesGroupsV2ParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupsV2ParamsResponseCodeBadcookie DNSTimeseriesGroupsV2ParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupsV2ParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupsV2ParamsResponseCodeNoerror, DNSTimeseriesGroupsV2ParamsResponseCodeFormerr, DNSTimeseriesGroupsV2ParamsResponseCodeServfail, DNSTimeseriesGroupsV2ParamsResponseCodeNxdomain, DNSTimeseriesGroupsV2ParamsResponseCodeNotimp, DNSTimeseriesGroupsV2ParamsResponseCodeRefused, DNSTimeseriesGroupsV2ParamsResponseCodeYxdomain, DNSTimeseriesGroupsV2ParamsResponseCodeYxrrset, DNSTimeseriesGroupsV2ParamsResponseCodeNxrrset, DNSTimeseriesGroupsV2ParamsResponseCodeNotauth, DNSTimeseriesGroupsV2ParamsResponseCodeNotzone, DNSTimeseriesGroupsV2ParamsResponseCodeBadsig, DNSTimeseriesGroupsV2ParamsResponseCodeBadkey, DNSTimeseriesGroupsV2ParamsResponseCodeBadtime, DNSTimeseriesGroupsV2ParamsResponseCodeBadmode, DNSTimeseriesGroupsV2ParamsResponseCodeBadname, DNSTimeseriesGroupsV2ParamsResponseCodeBadalg, DNSTimeseriesGroupsV2ParamsResponseCodeBadtrunc, DNSTimeseriesGroupsV2ParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupsV2ResponseEnvelope struct {
	Result  DNSTimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    dnsTimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupsV2ResponseEnvelope]
type dnsTimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
