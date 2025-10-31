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

// EmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingService] method instead.
type EmailRoutingService struct {
	Options          []option.RequestOption
	Summary          *EmailRoutingSummaryService
	TimeseriesGroups *EmailRoutingTimeseriesGroupService
}

// NewEmailRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailRoutingService(opts ...option.RequestOption) (r *EmailRoutingService) {
	r = &EmailRoutingService{}
	r.Options = opts
	r.Summary = NewEmailRoutingSummaryService(opts...)
	r.TimeseriesGroups = NewEmailRoutingTimeseriesGroupService(opts...)
	return
}

// Retrieves the distribution of email routing metrics by the specified dimension.
func (r *EmailRoutingService) SummaryV2(ctx context.Context, dimension EmailRoutingSummaryV2ParamsDimension, query EmailRoutingSummaryV2Params, opts ...option.RequestOption) (res *EmailRoutingSummaryV2Response, err error) {
	var env EmailRoutingSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/routing/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of email routing metrics grouped by dimension over
// time.
func (r *EmailRoutingService) TimeseriesGroupsV2(ctx context.Context, dimension EmailRoutingTimeseriesGroupsV2ParamsDimension, query EmailRoutingTimeseriesGroupsV2Params, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupsV2Response, err error) {
	var env EmailRoutingTimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/email/routing/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingSummaryV2Response struct {
	// Metadata for the results.
	Meta     EmailRoutingSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string                 `json:"summary_0,required"`
	JSON     emailRoutingSummaryV2ResponseJSON `json:"-"`
}

// emailRoutingSummaryV2ResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryV2Response]
type emailRoutingSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummaryV2ResponseMeta struct {
	ConfidenceInfo EmailRoutingSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummaryV2ResponseMetaJSON   `json:"-"`
}

// emailRoutingSummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryV2ResponseMeta]
type emailRoutingSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  emailRoutingSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryV2ResponseMetaConfidenceInfo]
type emailRoutingSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryV2ResponseMetaDateRange]
type emailRoutingSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummaryV2ResponseMetaNormalization string

const (
	EmailRoutingSummaryV2ResponseMetaNormalizationPercentage           EmailRoutingSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummaryV2ResponseMetaNormalizationMin0Max              EmailRoutingSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummaryV2ResponseMetaNormalizationMinMax               EmailRoutingSummaryV2ResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummaryV2ResponseMetaNormalizationRawValues            EmailRoutingSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummaryV2ResponseMetaNormalizationPercentageChange     EmailRoutingSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummaryV2ResponseMetaNormalizationRollingAverage       EmailRoutingSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummaryV2ResponseMetaNormalizationOverlappedPercentage EmailRoutingSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummaryV2ResponseMetaNormalizationRatio                EmailRoutingSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ResponseMetaNormalizationPercentage, EmailRoutingSummaryV2ResponseMetaNormalizationMin0Max, EmailRoutingSummaryV2ResponseMetaNormalizationMinMax, EmailRoutingSummaryV2ResponseMetaNormalizationRawValues, EmailRoutingSummaryV2ResponseMetaNormalizationPercentageChange, EmailRoutingSummaryV2ResponseMetaNormalizationRollingAverage, EmailRoutingSummaryV2ResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  emailRoutingSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryV2ResponseMetaUnit]
type emailRoutingSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupsV2Response]
type emailRoutingTimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupsV2ResponseMeta]
type emailRoutingTimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupsV2ResponseMetaDateRange]
type emailRoutingTimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupsV2ResponseMetaUnit]
type emailRoutingTimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                                      `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                              `json:"-,extras"`
	JSON        emailRoutingTimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupsV2ResponseSerie0]
type emailRoutingTimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryV2Params struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingSummaryV2ParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingSummaryV2ParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingSummaryV2ParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingSummaryV2ParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummaryV2ParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingSummaryV2ParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingSummaryV2ParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryV2Params]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type EmailRoutingSummaryV2ParamsDimension string

const (
	EmailRoutingSummaryV2ParamsDimensionIPVersion EmailRoutingSummaryV2ParamsDimension = "IP_VERSION"
	EmailRoutingSummaryV2ParamsDimensionEncrypted EmailRoutingSummaryV2ParamsDimension = "ENCRYPTED"
	EmailRoutingSummaryV2ParamsDimensionARC       EmailRoutingSummaryV2ParamsDimension = "ARC"
	EmailRoutingSummaryV2ParamsDimensionDKIM      EmailRoutingSummaryV2ParamsDimension = "DKIM"
	EmailRoutingSummaryV2ParamsDimensionDMARC     EmailRoutingSummaryV2ParamsDimension = "DMARC"
	EmailRoutingSummaryV2ParamsDimensionSPF       EmailRoutingSummaryV2ParamsDimension = "SPF"
)

func (r EmailRoutingSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsDimensionIPVersion, EmailRoutingSummaryV2ParamsDimensionEncrypted, EmailRoutingSummaryV2ParamsDimensionARC, EmailRoutingSummaryV2ParamsDimensionDKIM, EmailRoutingSummaryV2ParamsDimensionDMARC, EmailRoutingSummaryV2ParamsDimensionSPF:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ParamsARC string

const (
	EmailRoutingSummaryV2ParamsARCPass EmailRoutingSummaryV2ParamsARC = "PASS"
	EmailRoutingSummaryV2ParamsARCNone EmailRoutingSummaryV2ParamsARC = "NONE"
	EmailRoutingSummaryV2ParamsARCFail EmailRoutingSummaryV2ParamsARC = "FAIL"
)

func (r EmailRoutingSummaryV2ParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsARCPass, EmailRoutingSummaryV2ParamsARCNone, EmailRoutingSummaryV2ParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ParamsDKIM string

const (
	EmailRoutingSummaryV2ParamsDKIMPass EmailRoutingSummaryV2ParamsDKIM = "PASS"
	EmailRoutingSummaryV2ParamsDKIMNone EmailRoutingSummaryV2ParamsDKIM = "NONE"
	EmailRoutingSummaryV2ParamsDKIMFail EmailRoutingSummaryV2ParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryV2ParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsDKIMPass, EmailRoutingSummaryV2ParamsDKIMNone, EmailRoutingSummaryV2ParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ParamsDMARC string

const (
	EmailRoutingSummaryV2ParamsDMARCPass EmailRoutingSummaryV2ParamsDMARC = "PASS"
	EmailRoutingSummaryV2ParamsDMARCNone EmailRoutingSummaryV2ParamsDMARC = "NONE"
	EmailRoutingSummaryV2ParamsDMARCFail EmailRoutingSummaryV2ParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryV2ParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsDMARCPass, EmailRoutingSummaryV2ParamsDMARCNone, EmailRoutingSummaryV2ParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ParamsEncrypted string

const (
	EmailRoutingSummaryV2ParamsEncryptedEncrypted    EmailRoutingSummaryV2ParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryV2ParamsEncryptedNotEncrypted EmailRoutingSummaryV2ParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryV2ParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsEncryptedEncrypted, EmailRoutingSummaryV2ParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummaryV2ParamsFormat string

const (
	EmailRoutingSummaryV2ParamsFormatJson EmailRoutingSummaryV2ParamsFormat = "JSON"
	EmailRoutingSummaryV2ParamsFormatCsv  EmailRoutingSummaryV2ParamsFormat = "CSV"
)

func (r EmailRoutingSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsFormatJson, EmailRoutingSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ParamsIPVersion string

const (
	EmailRoutingSummaryV2ParamsIPVersionIPv4 EmailRoutingSummaryV2ParamsIPVersion = "IPv4"
	EmailRoutingSummaryV2ParamsIPVersionIPv6 EmailRoutingSummaryV2ParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsIPVersionIPv4, EmailRoutingSummaryV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ParamsSPF string

const (
	EmailRoutingSummaryV2ParamsSPFPass EmailRoutingSummaryV2ParamsSPF = "PASS"
	EmailRoutingSummaryV2ParamsSPFNone EmailRoutingSummaryV2ParamsSPF = "NONE"
	EmailRoutingSummaryV2ParamsSPFFail EmailRoutingSummaryV2ParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryV2ParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryV2ParamsSPFPass, EmailRoutingSummaryV2ParamsSPFNone, EmailRoutingSummaryV2ParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryV2ResponseEnvelope struct {
	Result  EmailRoutingSummaryV2Response             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    emailRoutingSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryV2ResponseEnvelope]
type emailRoutingSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupsV2ParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupsV2ParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupsV2ParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupsV2ParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupsV2ParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupsV2ParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type EmailRoutingTimeseriesGroupsV2ParamsDimension string

const (
	EmailRoutingTimeseriesGroupsV2ParamsDimensionIPVersion EmailRoutingTimeseriesGroupsV2ParamsDimension = "IP_VERSION"
	EmailRoutingTimeseriesGroupsV2ParamsDimensionEncrypted EmailRoutingTimeseriesGroupsV2ParamsDimension = "ENCRYPTED"
	EmailRoutingTimeseriesGroupsV2ParamsDimensionARC       EmailRoutingTimeseriesGroupsV2ParamsDimension = "ARC"
	EmailRoutingTimeseriesGroupsV2ParamsDimensionDKIM      EmailRoutingTimeseriesGroupsV2ParamsDimension = "DKIM"
	EmailRoutingTimeseriesGroupsV2ParamsDimensionDMARC     EmailRoutingTimeseriesGroupsV2ParamsDimension = "DMARC"
	EmailRoutingTimeseriesGroupsV2ParamsDimensionSPF       EmailRoutingTimeseriesGroupsV2ParamsDimension = "SPF"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsDimensionIPVersion, EmailRoutingTimeseriesGroupsV2ParamsDimensionEncrypted, EmailRoutingTimeseriesGroupsV2ParamsDimensionARC, EmailRoutingTimeseriesGroupsV2ParamsDimensionDKIM, EmailRoutingTimeseriesGroupsV2ParamsDimensionDMARC, EmailRoutingTimeseriesGroupsV2ParamsDimensionSPF:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupsV2ParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupsV2ParamsAggInterval15m EmailRoutingTimeseriesGroupsV2ParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupsV2ParamsAggInterval1h  EmailRoutingTimeseriesGroupsV2ParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupsV2ParamsAggInterval1d  EmailRoutingTimeseriesGroupsV2ParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupsV2ParamsAggInterval1w  EmailRoutingTimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsAggInterval15m, EmailRoutingTimeseriesGroupsV2ParamsAggInterval1h, EmailRoutingTimeseriesGroupsV2ParamsAggInterval1d, EmailRoutingTimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ParamsARC string

const (
	EmailRoutingTimeseriesGroupsV2ParamsARCPass EmailRoutingTimeseriesGroupsV2ParamsARC = "PASS"
	EmailRoutingTimeseriesGroupsV2ParamsARCNone EmailRoutingTimeseriesGroupsV2ParamsARC = "NONE"
	EmailRoutingTimeseriesGroupsV2ParamsARCFail EmailRoutingTimeseriesGroupsV2ParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsARCPass, EmailRoutingTimeseriesGroupsV2ParamsARCNone, EmailRoutingTimeseriesGroupsV2ParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ParamsDKIM string

const (
	EmailRoutingTimeseriesGroupsV2ParamsDKIMPass EmailRoutingTimeseriesGroupsV2ParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupsV2ParamsDKIMNone EmailRoutingTimeseriesGroupsV2ParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupsV2ParamsDKIMFail EmailRoutingTimeseriesGroupsV2ParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsDKIMPass, EmailRoutingTimeseriesGroupsV2ParamsDKIMNone, EmailRoutingTimeseriesGroupsV2ParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ParamsDMARC string

const (
	EmailRoutingTimeseriesGroupsV2ParamsDMARCPass EmailRoutingTimeseriesGroupsV2ParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupsV2ParamsDMARCNone EmailRoutingTimeseriesGroupsV2ParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupsV2ParamsDMARCFail EmailRoutingTimeseriesGroupsV2ParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsDMARCPass, EmailRoutingTimeseriesGroupsV2ParamsDMARCNone, EmailRoutingTimeseriesGroupsV2ParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupsV2ParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupsV2ParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupsV2ParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupsV2ParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupsV2ParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupsV2ParamsFormat string

const (
	EmailRoutingTimeseriesGroupsV2ParamsFormatJson EmailRoutingTimeseriesGroupsV2ParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupsV2ParamsFormatCsv  EmailRoutingTimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsFormatJson, EmailRoutingTimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupsV2ParamsIPVersionIPv4 EmailRoutingTimeseriesGroupsV2ParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupsV2ParamsIPVersionIPv6 EmailRoutingTimeseriesGroupsV2ParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsIPVersionIPv4, EmailRoutingTimeseriesGroupsV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ParamsSPF string

const (
	EmailRoutingTimeseriesGroupsV2ParamsSPFPass EmailRoutingTimeseriesGroupsV2ParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupsV2ParamsSPFNone EmailRoutingTimeseriesGroupsV2ParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupsV2ParamsSPFFail EmailRoutingTimeseriesGroupsV2ParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupsV2ParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupsV2ParamsSPFPass, EmailRoutingTimeseriesGroupsV2ParamsSPFNone, EmailRoutingTimeseriesGroupsV2ParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupsV2ResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupsV2ResponseEnvelope]
type emailRoutingTimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
