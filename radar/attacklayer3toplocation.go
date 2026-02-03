// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
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

// AttackLayer3TopLocationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3TopLocationService] method instead.
type AttackLayer3TopLocationService struct {
	Options []option.RequestOption
}

// NewAttackLayer3TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer3TopLocationService(opts ...option.RequestOption) (r *AttackLayer3TopLocationService) {
	r = &AttackLayer3TopLocationService{}
	r.Options = opts
	return
}

// Retrieves the origin locations of layer 3 attacks.
func (r *AttackLayer3TopLocationService) Origin(ctx context.Context, query AttackLayer3TopLocationOriginParams, opts ...option.RequestOption) (res *AttackLayer3TopLocationOriginResponse, err error) {
	var env AttackLayer3TopLocationOriginResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/attacks/layer3/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the target locations of layer 3 attacks.
func (r *AttackLayer3TopLocationService) Target(ctx context.Context, query AttackLayer3TopLocationTargetParams, opts ...option.RequestOption) (res *AttackLayer3TopLocationTargetResponse, err error) {
	var env AttackLayer3TopLocationTargetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/attacks/layer3/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TopLocationOriginResponse struct {
	// Metadata for the results.
	Meta AttackLayer3TopLocationOriginResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopLocationOriginResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopLocationOriginResponseJSON   `json:"-"`
}

// attackLayer3TopLocationOriginResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationOriginResponse]
type attackLayer3TopLocationOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TopLocationOriginResponseMeta struct {
	ConfidenceInfo AttackLayer3TopLocationOriginResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer3TopLocationOriginResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TopLocationOriginResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TopLocationOriginResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TopLocationOriginResponseMetaJSON   `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationOriginResponseMeta]
type attackLayer3TopLocationOriginResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer3TopLocationOriginResponseMetaConfidenceInfo]
type attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                       `json:"description,required"`
	EndDate     time.Time                                                                    `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer3TopLocationOriginResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopLocationOriginResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer3TopLocationOriginResponseMetaDateRange]
type attackLayer3TopLocationOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TopLocationOriginResponseMetaNormalization string

const (
	AttackLayer3TopLocationOriginResponseMetaNormalizationPercentage           AttackLayer3TopLocationOriginResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TopLocationOriginResponseMetaNormalizationMin0Max              AttackLayer3TopLocationOriginResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TopLocationOriginResponseMetaNormalizationMinMax               AttackLayer3TopLocationOriginResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TopLocationOriginResponseMetaNormalizationRawValues            AttackLayer3TopLocationOriginResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TopLocationOriginResponseMetaNormalizationPercentageChange     AttackLayer3TopLocationOriginResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TopLocationOriginResponseMetaNormalizationRollingAverage       AttackLayer3TopLocationOriginResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TopLocationOriginResponseMetaNormalizationOverlappedPercentage AttackLayer3TopLocationOriginResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TopLocationOriginResponseMetaNormalizationRatio                AttackLayer3TopLocationOriginResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TopLocationOriginResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginResponseMetaNormalizationPercentage, AttackLayer3TopLocationOriginResponseMetaNormalizationMin0Max, AttackLayer3TopLocationOriginResponseMetaNormalizationMinMax, AttackLayer3TopLocationOriginResponseMetaNormalizationRawValues, AttackLayer3TopLocationOriginResponseMetaNormalizationPercentageChange, AttackLayer3TopLocationOriginResponseMetaNormalizationRollingAverage, AttackLayer3TopLocationOriginResponseMetaNormalizationOverlappedPercentage, AttackLayer3TopLocationOriginResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  attackLayer3TopLocationOriginResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3TopLocationOriginResponseMetaUnit]
type attackLayer3TopLocationOriginResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginResponseTop0 struct {
	OriginCountryAlpha2 string                                        `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                        `json:"originCountryName,required"`
	Rank                float64                                       `json:"rank,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer3TopLocationOriginResponseTop0JSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationOriginResponseTop0]
type attackLayer3TopLocationOriginResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponse struct {
	// Metadata for the results.
	Meta AttackLayer3TopLocationTargetResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopLocationTargetResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopLocationTargetResponseJSON   `json:"-"`
}

// attackLayer3TopLocationTargetResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationTargetResponse]
type attackLayer3TopLocationTargetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TopLocationTargetResponseMeta struct {
	ConfidenceInfo AttackLayer3TopLocationTargetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer3TopLocationTargetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TopLocationTargetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TopLocationTargetResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TopLocationTargetResponseMetaJSON   `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationTargetResponseMeta]
type attackLayer3TopLocationTargetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer3TopLocationTargetResponseMetaConfidenceInfo]
type attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                       `json:"description,required"`
	EndDate     time.Time                                                                    `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer3TopLocationTargetResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopLocationTargetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer3TopLocationTargetResponseMetaDateRange]
type attackLayer3TopLocationTargetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TopLocationTargetResponseMetaNormalization string

const (
	AttackLayer3TopLocationTargetResponseMetaNormalizationPercentage           AttackLayer3TopLocationTargetResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TopLocationTargetResponseMetaNormalizationMin0Max              AttackLayer3TopLocationTargetResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TopLocationTargetResponseMetaNormalizationMinMax               AttackLayer3TopLocationTargetResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TopLocationTargetResponseMetaNormalizationRawValues            AttackLayer3TopLocationTargetResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TopLocationTargetResponseMetaNormalizationPercentageChange     AttackLayer3TopLocationTargetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TopLocationTargetResponseMetaNormalizationRollingAverage       AttackLayer3TopLocationTargetResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TopLocationTargetResponseMetaNormalizationOverlappedPercentage AttackLayer3TopLocationTargetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TopLocationTargetResponseMetaNormalizationRatio                AttackLayer3TopLocationTargetResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TopLocationTargetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetResponseMetaNormalizationPercentage, AttackLayer3TopLocationTargetResponseMetaNormalizationMin0Max, AttackLayer3TopLocationTargetResponseMetaNormalizationMinMax, AttackLayer3TopLocationTargetResponseMetaNormalizationRawValues, AttackLayer3TopLocationTargetResponseMetaNormalizationPercentageChange, AttackLayer3TopLocationTargetResponseMetaNormalizationRollingAverage, AttackLayer3TopLocationTargetResponseMetaNormalizationOverlappedPercentage, AttackLayer3TopLocationTargetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  attackLayer3TopLocationTargetResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3TopLocationTargetResponseMetaUnit]
type attackLayer3TopLocationTargetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponseTop0 struct {
	Rank                float64                                       `json:"rank,required"`
	TargetCountryAlpha2 string                                        `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                        `json:"targetCountryName,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer3TopLocationTargetResponseTop0JSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationTargetResponseTop0]
type attackLayer3TopLocationTargetResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginParams struct {
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
	Format param.Field[AttackLayer3TopLocationOriginParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TopLocationOriginParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TopLocationOriginParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopLocationOriginParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopLocationOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer3TopLocationOriginParamsFormat string

const (
	AttackLayer3TopLocationOriginParamsFormatJson AttackLayer3TopLocationOriginParamsFormat = "JSON"
	AttackLayer3TopLocationOriginParamsFormatCsv  AttackLayer3TopLocationOriginParamsFormat = "CSV"
)

func (r AttackLayer3TopLocationOriginParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsFormatJson, AttackLayer3TopLocationOriginParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginParamsIPVersion string

const (
	AttackLayer3TopLocationOriginParamsIPVersionIPv4 AttackLayer3TopLocationOriginParamsIPVersion = "IPv4"
	AttackLayer3TopLocationOriginParamsIPVersionIPv6 AttackLayer3TopLocationOriginParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopLocationOriginParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsIPVersionIPv4, AttackLayer3TopLocationOriginParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginParamsProtocol string

const (
	AttackLayer3TopLocationOriginParamsProtocolUdp  AttackLayer3TopLocationOriginParamsProtocol = "UDP"
	AttackLayer3TopLocationOriginParamsProtocolTCP  AttackLayer3TopLocationOriginParamsProtocol = "TCP"
	AttackLayer3TopLocationOriginParamsProtocolIcmp AttackLayer3TopLocationOriginParamsProtocol = "ICMP"
	AttackLayer3TopLocationOriginParamsProtocolGRE  AttackLayer3TopLocationOriginParamsProtocol = "GRE"
)

func (r AttackLayer3TopLocationOriginParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsProtocolUdp, AttackLayer3TopLocationOriginParamsProtocolTCP, AttackLayer3TopLocationOriginParamsProtocolIcmp, AttackLayer3TopLocationOriginParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginResponseEnvelope struct {
	Result  AttackLayer3TopLocationOriginResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer3TopLocationOriginResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3TopLocationOriginResponseEnvelope]
type attackLayer3TopLocationOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetParams struct {
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
	Format param.Field[AttackLayer3TopLocationTargetParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TopLocationTargetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TopLocationTargetParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopLocationTargetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopLocationTargetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer3TopLocationTargetParamsFormat string

const (
	AttackLayer3TopLocationTargetParamsFormatJson AttackLayer3TopLocationTargetParamsFormat = "JSON"
	AttackLayer3TopLocationTargetParamsFormatCsv  AttackLayer3TopLocationTargetParamsFormat = "CSV"
)

func (r AttackLayer3TopLocationTargetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsFormatJson, AttackLayer3TopLocationTargetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetParamsIPVersion string

const (
	AttackLayer3TopLocationTargetParamsIPVersionIPv4 AttackLayer3TopLocationTargetParamsIPVersion = "IPv4"
	AttackLayer3TopLocationTargetParamsIPVersionIPv6 AttackLayer3TopLocationTargetParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopLocationTargetParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsIPVersionIPv4, AttackLayer3TopLocationTargetParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetParamsProtocol string

const (
	AttackLayer3TopLocationTargetParamsProtocolUdp  AttackLayer3TopLocationTargetParamsProtocol = "UDP"
	AttackLayer3TopLocationTargetParamsProtocolTCP  AttackLayer3TopLocationTargetParamsProtocol = "TCP"
	AttackLayer3TopLocationTargetParamsProtocolIcmp AttackLayer3TopLocationTargetParamsProtocol = "ICMP"
	AttackLayer3TopLocationTargetParamsProtocolGRE  AttackLayer3TopLocationTargetParamsProtocol = "GRE"
)

func (r AttackLayer3TopLocationTargetParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsProtocolUdp, AttackLayer3TopLocationTargetParamsProtocolTCP, AttackLayer3TopLocationTargetParamsProtocolIcmp, AttackLayer3TopLocationTargetParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetResponseEnvelope struct {
	Result  AttackLayer3TopLocationTargetResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer3TopLocationTargetResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3TopLocationTargetResponseEnvelope]
type attackLayer3TopLocationTargetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
