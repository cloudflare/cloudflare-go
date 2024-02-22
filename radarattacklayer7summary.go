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

// Percentage distribution of mitigation techniques in Layer 7 attacks.
func (r *RadarAttackLayer7SummaryService) Get(ctx context.Context, query RadarAttackLayer7SummaryGetParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryGetResponseEnvelope
	path := "radar/attacks/layer7/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http method used.
func (r *RadarAttackLayer7SummaryService) HTTPMethod(ctx context.Context, query RadarAttackLayer7SummaryHTTPMethodParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryHTTPMethodResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryHTTPMethodResponseEnvelope
	path := "radar/attacks/layer7/summary/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http version used.
func (r *RadarAttackLayer7SummaryService) HTTPVersion(ctx context.Context, query RadarAttackLayer7SummaryHTTPVersionParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryHTTPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryHTTPVersionResponseEnvelope
	path := "radar/attacks/layer7/summary/http_version"
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
	path := "radar/attacks/layer7/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by managed rules used.
func (r *RadarAttackLayer7SummaryService) ManagedRules(ctx context.Context, query RadarAttackLayer7SummaryManagedRulesParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryManagedRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryManagedRulesResponseEnvelope
	path := "radar/attacks/layer7/summary/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by mitigation product used.
func (r *RadarAttackLayer7SummaryService) MitigationProduct(ctx context.Context, query RadarAttackLayer7SummaryMitigationProductParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryMitigationProductResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7SummaryMitigationProductResponseEnvelope
	path := "radar/attacks/layer7/summary/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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
	AccessRules        string                                          `json:"ACCESS_RULES,required"`
	APIShield          string                                          `json:"API_SHIELD,required"`
	BotManagement      string                                          `json:"BOT_MANAGEMENT,required"`
	DataLossPrevention string                                          `json:"DATA_LOSS_PREVENTION,required"`
	DDOS               string                                          `json:"DDOS,required"`
	IPReputation       string                                          `json:"IP_REPUTATION,required"`
	WAF                string                                          `json:"WAF,required"`
	JSON               radarAttackLayer7SummaryGetResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryGetResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryGetResponseSummary0]
type radarAttackLayer7SummaryGetResponseSummary0JSON struct {
	AccessRules        apijson.Field
	APIShield          apijson.Field
	BotManagement      apijson.Field
	DataLossPrevention apijson.Field
	DDOS               apijson.Field
	IPReputation       apijson.Field
	WAF                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodResponse struct {
	Meta     RadarAttackLayer7SummaryHTTPMethodResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryHTTPMethodResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryHTTPMethodResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryHTTPMethodResponse]
type radarAttackLayer7SummaryHTTPMethodResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryHTTPMethodResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	Normalization  string                                                       `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryHTTPMethodResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryHTTPMethodResponseMeta]
type radarAttackLayer7SummaryHTTPMethodResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodResponseMetaDateRange]
type radarAttackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodResponseSummary0 struct {
	Get  string                                                 `json:"GET,required"`
	Post string                                                 `json:"POST,required"`
	JSON radarAttackLayer7SummaryHTTPMethodResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseSummary0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryHTTPMethodResponseSummary0]
type radarAttackLayer7SummaryHTTPMethodResponseSummary0JSON struct {
	Get         apijson.Field
	Post        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionResponse struct {
	Meta     RadarAttackLayer7SummaryHTTPVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryHTTPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryHTTPVersionResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryHTTPVersionResponse]
type radarAttackLayer7SummaryHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryHTTPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryHTTPVersionResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryHTTPVersionResponseMeta]
type radarAttackLayer7SummaryHTTPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionResponseMetaDateRange]
type radarAttackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionResponseSummary0 struct {
	HTTP1X string                                                  `json:"HTTP/1.x,required"`
	HTTP2  string                                                  `json:"HTTP/2,required"`
	HTTP3  string                                                  `json:"HTTP/3,required"`
	JSON   radarAttackLayer7SummaryHTTPVersionResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseSummary0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryHTTPVersionResponseSummary0]
type radarAttackLayer7SummaryHTTPVersionResponseSummary0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
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

type RadarAttackLayer7SummaryManagedRulesResponse struct {
	Meta     RadarAttackLayer7SummaryManagedRulesResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryManagedRulesResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryManagedRulesResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryManagedRulesResponse]
type radarAttackLayer7SummaryManagedRulesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRulesResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryManagedRulesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	Normalization  string                                                         `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryManagedRulesResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryManagedRulesResponseMeta]
type radarAttackLayer7SummaryManagedRulesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRulesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRulesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryManagedRulesResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryManagedRulesResponseMetaDateRange]
type radarAttackLayer7SummaryManagedRulesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRulesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRulesResponseSummary0 struct {
	Bot         string                                                   `json:"Bot,required"`
	HTTPAnomaly string                                                   `json:"HTTP Anomaly,required"`
	JSON        radarAttackLayer7SummaryManagedRulesResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseSummary0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryManagedRulesResponseSummary0]
type radarAttackLayer7SummaryManagedRulesResponseSummary0JSON struct {
	Bot         apijson.Field
	HTTPAnomaly apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRulesResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductResponse struct {
	Meta     RadarAttackLayer7SummaryMitigationProductResponseMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryMitigationProductResponseSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryMitigationProductResponseJSON     `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryMitigationProductResponse]
type radarAttackLayer7SummaryMitigationProductResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryMitigationProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductResponseMeta struct {
	DateRange      []RadarAttackLayer7SummaryMitigationProductResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	Normalization  string                                                              `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryMitigationProductResponseMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryMitigationProductResponseMeta]
type radarAttackLayer7SummaryMitigationProductResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryMitigationProductResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryMitigationProductResponseMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryMitigationProductResponseMetaDateRange]
type radarAttackLayer7SummaryMitigationProductResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryMitigationProductResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo]
type radarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductResponseSummary0 struct {
	DDOS string                                                        `json:"DDOS,required"`
	WAF  string                                                        `json:"WAF,required"`
	JSON radarAttackLayer7SummaryMitigationProductResponseSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseSummary0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryMitigationProductResponseSummary0]
type radarAttackLayer7SummaryMitigationProductResponseSummary0JSON struct {
	DDOS        apijson.Field
	WAF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryMitigationProductResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
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

type RadarAttackLayer7SummaryHTTPMethodParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryHTTPMethodParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryHTTPMethodParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryHTTPMethodParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryHTTPMethodParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryHTTPMethodParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryHTTPMethodParamsDateRange string

const (
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange1d         RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "1d"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange2d         RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "2d"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange7d         RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "7d"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange14d        RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "14d"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange28d        RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "28d"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange12w        RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "12w"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange24w        RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "24w"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange52w        RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "52w"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange1dControl  RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange2dControl  RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange7dControl  RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange14dControl RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange28dControl RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange12wControl RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryHTTPMethodParamsDateRange24wControl RadarAttackLayer7SummaryHTTPMethodParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryHTTPMethodParamsFormat string

const (
	RadarAttackLayer7SummaryHTTPMethodParamsFormatJson RadarAttackLayer7SummaryHTTPMethodParamsFormat = "JSON"
	RadarAttackLayer7SummaryHTTPMethodParamsFormatCsv  RadarAttackLayer7SummaryHTTPMethodParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryHTTPMethodParamsIPVersion string

const (
	RadarAttackLayer7SummaryHTTPMethodParamsIPVersionIPv4 RadarAttackLayer7SummaryHTTPMethodParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryHTTPMethodParamsIPVersionIPv6 RadarAttackLayer7SummaryHTTPMethodParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductDDOS               RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductWAF                RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement      RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductAccessRules        RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductIPReputation       RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductAPIShield          RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryHTTPMethodParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryHTTPMethodParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

type RadarAttackLayer7SummaryHTTPMethodResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryHTTPMethodResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAttackLayer7SummaryHTTPMethodResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryHTTPMethodResponseEnvelope]
type radarAttackLayer7SummaryHTTPMethodResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryHTTPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryHTTPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryHTTPVersionParamsDateRange string

const (
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange1d         RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "1d"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange2d         RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "2d"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange7d         RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "7d"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange14d        RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "14d"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange28d        RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "28d"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange12w        RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "12w"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange24w        RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "24w"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange52w        RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "52w"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange1dControl  RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange2dControl  RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange7dControl  RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange14dControl RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange28dControl RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange12wControl RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryHTTPVersionParamsDateRange24wControl RadarAttackLayer7SummaryHTTPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryHTTPVersionParamsFormat string

const (
	RadarAttackLayer7SummaryHTTPVersionParamsFormatJson RadarAttackLayer7SummaryHTTPVersionParamsFormat = "JSON"
	RadarAttackLayer7SummaryHTTPVersionParamsFormatCsv  RadarAttackLayer7SummaryHTTPVersionParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodGet             RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPost            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete          RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPut             RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodHead            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPurge           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodOptions         RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPropfind        RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodMkcol           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPatch           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodACL             RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodBcopy           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodBdelete         RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodBmove           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckin         RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckout        RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodConnect         RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodCopy            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodLabel           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodLock            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodMerge           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodMove            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodNotify          RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodPoll            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodProppatch       RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodReport          RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodSearch          RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodTrace           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodUnlock          RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodUpdate          RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodJson            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodCook            RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethodTrack           RadarAttackLayer7SummaryHTTPVersionParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryHTTPVersionParamsIPVersion string

const (
	RadarAttackLayer7SummaryHTTPVersionParamsIPVersionIPv4 RadarAttackLayer7SummaryHTTPVersionParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryHTTPVersionParamsIPVersionIPv6 RadarAttackLayer7SummaryHTTPVersionParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductDDOS               RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductWAF                RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement      RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductAccessRules        RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductIPReputation       RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductAPIShield          RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryHTTPVersionParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryHTTPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

type RadarAttackLayer7SummaryHTTPVersionResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryHTTPVersionResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarAttackLayer7SummaryHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryHTTPVersionResponseEnvelope]
type radarAttackLayer7SummaryHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryIPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryIPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryIPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
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

// Format results are returned in.
type RadarAttackLayer7SummaryIPVersionParamsFormat string

const (
	RadarAttackLayer7SummaryIPVersionParamsFormatJson RadarAttackLayer7SummaryIPVersionParamsFormat = "JSON"
	RadarAttackLayer7SummaryIPVersionParamsFormatCsv  RadarAttackLayer7SummaryIPVersionParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryIPVersionParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodGet             RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPost            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodDelete          RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPut             RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodHead            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPurge           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodOptions         RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPropfind        RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodMkcol           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPatch           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodACL             RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodBcopy           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodBdelete         RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodBmove           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodCheckin         RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodCheckout        RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodConnect         RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodCopy            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodLabel           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodLock            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodMerge           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodMove            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodNotify          RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodPoll            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodProppatch       RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodReport          RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodSearch          RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodTrace           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodUnlock          RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodUpdate          RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodJson            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodCook            RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryIPVersionParamsHTTPMethodTrack           RadarAttackLayer7SummaryIPVersionParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryIPVersionParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryIPVersionParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductDDOS               RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductWAF                RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductBotManagement      RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductAccessRules        RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductIPReputation       RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductAPIShield          RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryIPVersionParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryIPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
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

type RadarAttackLayer7SummaryManagedRulesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryManagedRulesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryManagedRulesParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryManagedRulesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryManagedRulesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryManagedRulesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryManagedRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryManagedRulesParamsDateRange string

const (
	RadarAttackLayer7SummaryManagedRulesParamsDateRange1d         RadarAttackLayer7SummaryManagedRulesParamsDateRange = "1d"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange2d         RadarAttackLayer7SummaryManagedRulesParamsDateRange = "2d"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange7d         RadarAttackLayer7SummaryManagedRulesParamsDateRange = "7d"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange14d        RadarAttackLayer7SummaryManagedRulesParamsDateRange = "14d"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange28d        RadarAttackLayer7SummaryManagedRulesParamsDateRange = "28d"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange12w        RadarAttackLayer7SummaryManagedRulesParamsDateRange = "12w"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange24w        RadarAttackLayer7SummaryManagedRulesParamsDateRange = "24w"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange52w        RadarAttackLayer7SummaryManagedRulesParamsDateRange = "52w"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange1dControl  RadarAttackLayer7SummaryManagedRulesParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange2dControl  RadarAttackLayer7SummaryManagedRulesParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange7dControl  RadarAttackLayer7SummaryManagedRulesParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange14dControl RadarAttackLayer7SummaryManagedRulesParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange28dControl RadarAttackLayer7SummaryManagedRulesParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange12wControl RadarAttackLayer7SummaryManagedRulesParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryManagedRulesParamsDateRange24wControl RadarAttackLayer7SummaryManagedRulesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryManagedRulesParamsFormat string

const (
	RadarAttackLayer7SummaryManagedRulesParamsFormatJson RadarAttackLayer7SummaryManagedRulesParamsFormat = "JSON"
	RadarAttackLayer7SummaryManagedRulesParamsFormatCsv  RadarAttackLayer7SummaryManagedRulesParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodGet             RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPost            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodDelete          RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPut             RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodHead            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPurge           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodOptions         RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPropfind        RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodMkcol           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPatch           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodACL             RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodBcopy           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodBdelete         RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodBmove           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodCheckin         RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodCheckout        RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodConnect         RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodCopy            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodLabel           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodLock            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodMerge           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodMove            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodNotify          RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodPoll            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodProppatch       RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodReport          RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodSearch          RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodTrace           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodUnlock          RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodUpdate          RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodJson            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodCook            RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPMethodTrack           RadarAttackLayer7SummaryManagedRulesParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryManagedRulesParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryManagedRulesParamsIPVersion string

const (
	RadarAttackLayer7SummaryManagedRulesParamsIPVersionIPv4 RadarAttackLayer7SummaryManagedRulesParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryManagedRulesParamsIPVersionIPv6 RadarAttackLayer7SummaryManagedRulesParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductDDOS               RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductWAF                RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement      RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductAccessRules        RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductIPReputation       RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductAPIShield          RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryManagedRulesParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryManagedRulesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

type RadarAttackLayer7SummaryManagedRulesResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryManagedRulesResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer7SummaryManagedRulesResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryManagedRulesResponseEnvelope]
type radarAttackLayer7SummaryManagedRulesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryMitigationProductParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryMitigationProductParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryMitigationProductParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryMitigationProductParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryMitigationProductParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryMitigationProductParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryMitigationProductParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryMitigationProductParamsDateRange string

const (
	RadarAttackLayer7SummaryMitigationProductParamsDateRange1d         RadarAttackLayer7SummaryMitigationProductParamsDateRange = "1d"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange2d         RadarAttackLayer7SummaryMitigationProductParamsDateRange = "2d"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange7d         RadarAttackLayer7SummaryMitigationProductParamsDateRange = "7d"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange14d        RadarAttackLayer7SummaryMitigationProductParamsDateRange = "14d"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange28d        RadarAttackLayer7SummaryMitigationProductParamsDateRange = "28d"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange12w        RadarAttackLayer7SummaryMitigationProductParamsDateRange = "12w"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange24w        RadarAttackLayer7SummaryMitigationProductParamsDateRange = "24w"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange52w        RadarAttackLayer7SummaryMitigationProductParamsDateRange = "52w"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange1dControl  RadarAttackLayer7SummaryMitigationProductParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange2dControl  RadarAttackLayer7SummaryMitigationProductParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange7dControl  RadarAttackLayer7SummaryMitigationProductParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange14dControl RadarAttackLayer7SummaryMitigationProductParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange28dControl RadarAttackLayer7SummaryMitigationProductParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange12wControl RadarAttackLayer7SummaryMitigationProductParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryMitigationProductParamsDateRange24wControl RadarAttackLayer7SummaryMitigationProductParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryMitigationProductParamsFormat string

const (
	RadarAttackLayer7SummaryMitigationProductParamsFormatJson RadarAttackLayer7SummaryMitigationProductParamsFormat = "JSON"
	RadarAttackLayer7SummaryMitigationProductParamsFormatCsv  RadarAttackLayer7SummaryMitigationProductParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodGet             RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPost            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodDelete          RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPut             RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodHead            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPurge           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodOptions         RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPropfind        RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodMkcol           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPatch           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodACL             RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodBcopy           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodBdelete         RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodBmove           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodCheckin         RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodCheckout        RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodConnect         RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodCopy            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodLabel           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodLock            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodMerge           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodMove            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodNotify          RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodPoll            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodProppatch       RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodReport          RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodSearch          RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodTrace           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodUnlock          RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodUpdate          RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodJson            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodCook            RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPMethodTrack           RadarAttackLayer7SummaryMitigationProductParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryMitigationProductParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryMitigationProductParamsIPVersion string

const (
	RadarAttackLayer7SummaryMitigationProductParamsIPVersionIPv4 RadarAttackLayer7SummaryMitigationProductParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryMitigationProductParamsIPVersionIPv6 RadarAttackLayer7SummaryMitigationProductParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryMitigationProductResponseEnvelope struct {
	Result  RadarAttackLayer7SummaryMitigationProductResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarAttackLayer7SummaryMitigationProductResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7SummaryMitigationProductResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryMitigationProductResponseEnvelope]
type radarAttackLayer7SummaryMitigationProductResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryMitigationProductResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
