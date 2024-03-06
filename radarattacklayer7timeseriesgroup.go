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

// Get a time series of the percentual distribution of mitigation techniques, over
// time.
func (r *RadarAttackLayer7TimeseriesGroupService) Get(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupGetResponseEnvelope
	path := "radar/attacks/layer7/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http method used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) HTTPMethod(ctx context.Context, query RadarAttackLayer7TimeseriesGroupHTTPMethodParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupHTTPMethodResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope
	path := "radar/attacks/layer7/timeseries_groups/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http version used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) HTTPVersion(ctx context.Context, query RadarAttackLayer7TimeseriesGroupHTTPVersionParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupHTTPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope
	path := "radar/attacks/layer7/timeseries_groups/http_version"
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
	path := "radar/attacks/layer7/timeseries_groups/industry"
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
	path := "radar/attacks/layer7/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by managed rules used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) ManagedRules(ctx context.Context, query RadarAttackLayer7TimeseriesGroupManagedRulesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupManagedRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelope
	path := "radar/attacks/layer7/timeseries_groups/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by mitigation product used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) MitigationProduct(ctx context.Context, query RadarAttackLayer7TimeseriesGroupMitigationProductParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupMitigationProductResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelope
	path := "radar/attacks/layer7/timeseries_groups/mitigation_product"
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
	path := "radar/attacks/layer7/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

func (r radarAttackLayer7TimeseriesGroupGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupGetResponseMetaJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetResponseSerie0 struct {
	AccessRules        []string                                              `json:"ACCESS_RULES,required"`
	APIShield          []string                                              `json:"API_SHIELD,required"`
	BotManagement      []string                                              `json:"BOT_MANAGEMENT,required"`
	DataLossPrevention []string                                              `json:"DATA_LOSS_PREVENTION,required"`
	DDOS               []string                                              `json:"DDOS,required"`
	IPReputation       []string                                              `json:"IP_REPUTATION,required"`
	Timestamps         []time.Time                                           `json:"timestamps,required" format:"date-time"`
	WAF                []string                                              `json:"WAF,required"`
	JSON               radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAttackLayer7TimeseriesGroupGetResponseSerie0]
type radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON struct {
	AccessRules        apijson.Field
	APIShield          apijson.Field
	BotManagement      apijson.Field
	DataLossPrevention apijson.Field
	DDOS               apijson.Field
	IPReputation       apijson.Field
	Timestamps         apijson.Field
	WAF                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupHTTPMethodResponse struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupHTTPMethodResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupHTTPMethodResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupHTTPMethodResponse]
type radarAttackLayer7TimeseriesGroupHTTPMethodResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupHTTPMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupHTTPMethodResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0 struct {
	Get        []string                                                     `json:"GET,required"`
	Timestamps []string                                                     `json:"timestamps,required"`
	JSON       radarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0]
type radarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON struct {
	Get         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupHTTPVersionResponse struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupHTTPVersionResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupHTTPVersionResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupHTTPVersionResponse]
type radarAttackLayer7TimeseriesGroupHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0 struct {
	HTTP1X     []string                                                      `json:"HTTP/1.x,required"`
	Timestamps []string                                                      `json:"timestamps,required"`
	JSON       radarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0]
type radarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON struct {
	HTTP1X      apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupIndustryResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupIndustryResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupManagedRulesResponse struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupManagedRulesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupManagedRulesResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TimeseriesGroupManagedRulesResponse]
type radarAttackLayer7TimeseriesGroupManagedRulesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupManagedRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupManagedRulesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0 struct {
	Bot        []string                                                       `json:"Bot,required"`
	Timestamps []string                                                       `json:"timestamps,required"`
	JSON       radarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0]
type radarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON struct {
	Bot         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupMitigationProductResponse struct {
	Meta   interface{}                                                     `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupMitigationProductResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupMitigationProductResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupMitigationProductResponse]
type radarAttackLayer7TimeseriesGroupMitigationProductResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupMitigationProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupMitigationProductResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0 struct {
	DDOS       []string                                                            `json:"DDOS,required"`
	Timestamps []string                                                            `json:"timestamps,required"`
	JSON       radarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0]
type radarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON struct {
	DDOS        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupVerticalResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarAttackLayer7TimeseriesGroupVerticalResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetParamsAggInterval] `query:"aggInterval"`
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

func (r radarAttackLayer7TimeseriesGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupHTTPMethodParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupHTTPMethodParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupHTTPMethodParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange1d         RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange2d         RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange7d         RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange14d        RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange28d        RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange12w        RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange24w        RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange52w        RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupHTTPMethodParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormatCsv  RadarAttackLayer7TimeseriesGroupHTTPMethodParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion = "IPv6"
)

type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDOS               RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF                RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupHTTPMethodParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupHTTPMethodResponse             `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope]
type radarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupHTTPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange1d         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange2d         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange7d         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange14d        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange28d        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange12w        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange24w        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange52w        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupHTTPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormatCsv  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion = "IPv6"
)

type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDOS               RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF                RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupHTTPVersionParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupHTTPVersionResponse             `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope]
type radarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupIndustryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct] `query:"mitigationProduct"`
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

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupIndustryParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsFormatJson RadarAttackLayer7TimeseriesGroupIndustryParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupIndustryParamsFormatCsv  RadarAttackLayer7TimeseriesGroupIndustryParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

type RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDOS               RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF                RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
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

func (r radarAttackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
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

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsFormatJson RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsFormatCsv  RadarAttackLayer7TimeseriesGroupIPVersionParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDOS               RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF                RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
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

func (r radarAttackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupManagedRulesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupManagedRulesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupManagedRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange1d         RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange2d         RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange7d         RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange14d        RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange28d        RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange12w        RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange24w        RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange52w        RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupManagedRulesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormatJson RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupManagedRulesParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupManagedRulesParamsIPVersion = "IPv6"
)

type RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDOS               RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF                RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupManagedRulesParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupManagedRulesResponse             `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelope]
type radarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupMitigationProductParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesGroupMitigationProductParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupMitigationProductParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange1d         RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange2d         RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange7d         RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange14d        RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange28d        RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange12w        RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange24w        RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange52w        RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange1dControl  RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange2dControl  RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange7dControl  RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange14dControl RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange28dControl RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange12wControl RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange24wControl RadarAttackLayer7TimeseriesGroupMitigationProductParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormatJson RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormatCsv  RadarAttackLayer7TimeseriesGroupMitigationProductParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupMitigationProductParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupMitigationProductParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesGroupMitigationProductResponse             `json:"result,required"`
	Success bool                                                                  `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelope]
type radarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupVerticalParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct] `query:"mitigationProduct"`
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

// Format results are returned in.
type RadarAttackLayer7TimeseriesGroupVerticalParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsFormatJson RadarAttackLayer7TimeseriesGroupVerticalParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupVerticalParamsFormatCsv  RadarAttackLayer7TimeseriesGroupVerticalParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

type RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDOS               RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF                RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "DATA_LOSS_PREVENTION"
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

func (r radarAttackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
