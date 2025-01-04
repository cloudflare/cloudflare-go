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

// AttackLayer7TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TimeseriesGroupService] method instead.
type AttackLayer7TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAttackLayer7TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackLayer7TimeseriesGroupService(opts ...option.RequestOption) (r *AttackLayer7TimeseriesGroupService) {
	r = &AttackLayer7TimeseriesGroupService{}
	r.Options = opts
	return
}

// Get a time series of the distribution of mitigation techniques over time.
func (r *AttackLayer7TimeseriesGroupService) Get(ctx context.Context, query AttackLayer7TimeseriesGroupGetParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupGetResponse, err error) {
	var env AttackLayer7TimeseriesGroupGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http method used over time.
func (r *AttackLayer7TimeseriesGroupService) HTTPMethod(ctx context.Context, query AttackLayer7TimeseriesGroupHTTPMethodParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupHTTPMethodResponse, err error) {
	var env AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http version used over time.
func (r *AttackLayer7TimeseriesGroupService) HTTPVersion(ctx context.Context, query AttackLayer7TimeseriesGroupHTTPVersionParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupHTTPVersionResponse, err error) {
	var env AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by targeted industry over time.
func (r *AttackLayer7TimeseriesGroupService) Industry(ctx context.Context, query AttackLayer7TimeseriesGroupIndustryParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupIndustryResponse, err error) {
	var env AttackLayer7TimeseriesGroupIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used over time.
func (r *AttackLayer7TimeseriesGroupService) IPVersion(ctx context.Context, query AttackLayer7TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupIPVersionResponse, err error) {
	var env AttackLayer7TimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by managed rules used over time.
func (r *AttackLayer7TimeseriesGroupService) ManagedRules(ctx context.Context, query AttackLayer7TimeseriesGroupManagedRulesParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupManagedRulesResponse, err error) {
	var env AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by mitigation product used over time.
func (r *AttackLayer7TimeseriesGroupService) MitigationProduct(ctx context.Context, query AttackLayer7TimeseriesGroupMitigationProductParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupMitigationProductResponse, err error) {
	var env AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by targeted vertical over time.
func (r *AttackLayer7TimeseriesGroupService) Vertical(ctx context.Context, query AttackLayer7TimeseriesGroupVerticalParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupVerticalResponse, err error) {
	var env AttackLayer7TimeseriesGroupVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TimeseriesGroupGetResponse struct {
	Meta   AttackLayer7TimeseriesGroupGetResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupGetResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupGetResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseJSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesGroupGetResponse]
type attackLayer7TimeseriesGroupGetResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupGetResponseMeta struct {
	AggInterval    string                                                   `json:"aggInterval,required"`
	DateRange      []AttackLayer7TimeseriesGroupGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TimeseriesGroupGetResponseMetaJSON           `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupGetResponseMeta]
type attackLayer7TimeseriesGroupGetResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupGetResponseMetaDateRange]
type attackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupGetResponseSerie0 struct {
	AccessRules        []string                                         `json:"ACCESS_RULES,required"`
	APIShield          []string                                         `json:"API_SHIELD,required"`
	BotManagement      []string                                         `json:"BOT_MANAGEMENT,required"`
	DataLossPrevention []string                                         `json:"DATA_LOSS_PREVENTION,required"`
	DDoS               []string                                         `json:"DDOS,required"`
	IPReputation       []string                                         `json:"IP_REPUTATION,required"`
	Timestamps         []time.Time                                      `json:"timestamps,required" format:"date-time"`
	WAF                []string                                         `json:"WAF,required"`
	JSON               attackLayer7TimeseriesGroupGetResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseSerie0JSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupGetResponseSerie0]
type attackLayer7TimeseriesGroupGetResponseSerie0JSON struct {
	AccessRules        apijson.Field
	APIShield          apijson.Field
	BotManagement      apijson.Field
	DataLossPrevention apijson.Field
	DDoS               apijson.Field
	IPReputation       apijson.Field
	Timestamps         apijson.Field
	WAF                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupGetResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPMethodResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupHTTPMethodResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupHTTPMethodResponse]
type attackLayer7TimeseriesGroupHTTPMethodResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0 struct {
	Get        []string                                                `json:"GET,required"`
	Timestamps []string                                                `json:"timestamps,required"`
	JSON       attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0]
type attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON struct {
	Get         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupHTTPVersionResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponse]
type attackLayer7TimeseriesGroupHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0 struct {
	HTTP1X     []string                                                 `json:"HTTP/1.x,required"`
	Timestamps []string                                                 `json:"timestamps,required"`
	JSON       attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0]
type attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON struct {
	HTTP1X      apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupIndustryResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupIndustryResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupIndustryResponse]
type attackLayer7TimeseriesGroupIndustryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryResponseSerie0 struct {
	Timestamps  []string                                              `json:"timestamps,required"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupIndustryResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupIndustryResponseSerie0]
type attackLayer7TimeseriesGroupIndustryResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupIPVersionResponse]
type attackLayer7TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                               `json:"IPv4,required"`
	IPv6       []string                                               `json:"IPv6,required"`
	Timestamps []string                                               `json:"timestamps,required"`
	JSON       attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIPVersionResponseSerie0]
type attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesResponse struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupManagedRulesResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupManagedRulesResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupManagedRulesResponse]
type attackLayer7TimeseriesGroupManagedRulesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesResponseSerie0 struct {
	Timestamps  []string                                                  `json:"timestamps,required"`
	ExtraFields map[string][]string                                       `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupManagedRulesResponseSerie0]
type attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductResponse struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupMitigationProductResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupMitigationProductResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupMitigationProductResponse]
type attackLayer7TimeseriesGroupMitigationProductResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductResponseSerie0 struct {
	DDoS       []string                                                       `json:"DDOS,required"`
	Timestamps []string                                                       `json:"timestamps,required"`
	JSON       attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseSerie0]
type attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON struct {
	DDoS        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupVerticalResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupVerticalResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupVerticalResponse]
type attackLayer7TimeseriesGroupVerticalResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalResponseSerie0 struct {
	Timestamps  []string                                              `json:"timestamps,required"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupVerticalResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupVerticalResponseSerie0]
type attackLayer7TimeseriesGroupVerticalResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupGetParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupGetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TimeseriesGroupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupGetParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupGetParamsAggInterval15m AttackLayer7TimeseriesGroupGetParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupGetParamsAggInterval1h  AttackLayer7TimeseriesGroupGetParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupGetParamsAggInterval1d  AttackLayer7TimeseriesGroupGetParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupGetParamsAggInterval1w  AttackLayer7TimeseriesGroupGetParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupGetParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupGetParamsAggInterval15m, AttackLayer7TimeseriesGroupGetParamsAggInterval1h, AttackLayer7TimeseriesGroupGetParamsAggInterval1d, AttackLayer7TimeseriesGroupGetParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupGetParamsFormat string

const (
	AttackLayer7TimeseriesGroupGetParamsFormatJson AttackLayer7TimeseriesGroupGetParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupGetParamsFormatCsv  AttackLayer7TimeseriesGroupGetParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupGetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupGetParamsFormatJson, AttackLayer7TimeseriesGroupGetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupGetResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupGetResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupGetResponseEnvelope]
type attackLayer7TimeseriesGroupGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPMethodParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupHTTPMethodParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupHTTPMethodParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupHTTPMethodParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1h  AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1d  AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1w  AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m, AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1h, AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1d, AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupHTTPMethodParamsFormat string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson AttackLayer7TimeseriesGroupHTTPMethodParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupHTTPMethodParamsFormatCsv  AttackLayer7TimeseriesGroupHTTPMethodParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson, AttackLayer7TimeseriesGroupHTTPMethodParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4 AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6 AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4, AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF                AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage, AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupHTTPMethodResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope]
type attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupHTTPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1h  AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1d  AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1w  AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m, AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1h, AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1d, AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupHTTPVersionParamsFormat string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson AttackLayer7TimeseriesGroupHTTPVersionParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupHTTPVersionParamsFormatCsv  AttackLayer7TimeseriesGroupHTTPVersionParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson, AttackLayer7TimeseriesGroupHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet             AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPut             AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodHead            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodACL             AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLock            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMove            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodReport          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodJson            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCook            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPut, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodHead, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodACL, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLock, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMove, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodReport, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodJson, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCook, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4 AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6 AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4, AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF                AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage, AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupHTTPVersionResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope]
type attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupIndustryParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupIndustryParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupIndustryParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval15m AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval1h  AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval1d  AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval1w  AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsAggInterval15m, AttackLayer7TimeseriesGroupIndustryParamsAggInterval1h, AttackLayer7TimeseriesGroupIndustryParamsAggInterval1d, AttackLayer7TimeseriesGroupIndustryParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupIndustryParamsFormat string

const (
	AttackLayer7TimeseriesGroupIndustryParamsFormatJson AttackLayer7TimeseriesGroupIndustryParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupIndustryParamsFormatCsv  AttackLayer7TimeseriesGroupIndustryParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsFormatJson, AttackLayer7TimeseriesGroupIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet             AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPut             AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodHead            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodACL             AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLock            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMove            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodReport          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodJson            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCook            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPut, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodHead, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodACL, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLock, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMove, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodReport, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodJson, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCook, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4 AttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6 AttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4, AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF                AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupIndustryParamsNormalization string

const (
	AttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage AttackLayer7TimeseriesGroupIndustryParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupIndustryParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupIndustryParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage, AttackLayer7TimeseriesGroupIndustryParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupIndustryResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIndustryResponseEnvelope]
type attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupIPVersionParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h  AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1d  AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1w  AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m, AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h, AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1d, AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupIPVersionParamsFormat string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsFormatJson AttackLayer7TimeseriesGroupIPVersionParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupIPVersionParamsFormatCsv  AttackLayer7TimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsFormatJson, AttackLayer7TimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet             AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPut             AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodHead            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodACL             AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLock            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMove            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodReport          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodJson            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCook            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPut, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodHead, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodACL, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLock, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMove, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodReport, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodJson, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCook, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF                AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupIPVersionParamsNormalization string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage AttackLayer7TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupIPVersionParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage, AttackLayer7TimeseriesGroupIPVersionParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIPVersionResponseEnvelope]
type attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupManagedRulesParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupManagedRulesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupManagedRulesParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupManagedRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1h  AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1d  AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1w  AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m, AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1h, AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1d, AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupManagedRulesParamsFormat string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsFormatJson AttackLayer7TimeseriesGroupManagedRulesParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupManagedRulesParamsFormatCsv  AttackLayer7TimeseriesGroupManagedRulesParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsFormatJson, AttackLayer7TimeseriesGroupManagedRulesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet             AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPut             AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodHead            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodACL             AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLock            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMove            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodReport          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodJson            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCook            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPut, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodHead, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodACL, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLock, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMove, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodReport, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodJson, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCook, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4 AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6 AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4, AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF                AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupManagedRulesParamsNormalization string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage AttackLayer7TimeseriesGroupManagedRulesParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupManagedRulesParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage, AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupManagedRulesResponse             `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope]
type attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupMitigationProductParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupMitigationProductParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupMitigationProductParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupMitigationProductParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1h  AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1d  AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1w  AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m, AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1h, AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1d, AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupMitigationProductParamsFormat string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsFormatJson AttackLayer7TimeseriesGroupMitigationProductParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupMitigationProductParamsFormatCsv  AttackLayer7TimeseriesGroupMitigationProductParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsFormatJson, AttackLayer7TimeseriesGroupMitigationProductParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet             AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPut             AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodHead            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodACL             AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLock            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMove            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodReport          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodJson            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCook            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPut, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodHead, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodACL, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLock, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMove, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodReport, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodJson, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCook, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4 AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6 AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4, AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupMitigationProductParamsNormalization string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage AttackLayer7TimeseriesGroupMitigationProductParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupMitigationProductParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage, AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupMitigationProductResponse             `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope]
type attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupVerticalParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupVerticalParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupVerticalParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval15m AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval1h  AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval1d  AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval1w  AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsAggInterval15m, AttackLayer7TimeseriesGroupVerticalParamsAggInterval1h, AttackLayer7TimeseriesGroupVerticalParamsAggInterval1d, AttackLayer7TimeseriesGroupVerticalParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesGroupVerticalParamsFormat string

const (
	AttackLayer7TimeseriesGroupVerticalParamsFormatJson AttackLayer7TimeseriesGroupVerticalParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupVerticalParamsFormatCsv  AttackLayer7TimeseriesGroupVerticalParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsFormatJson, AttackLayer7TimeseriesGroupVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet             AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPut             AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodHead            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodACL             AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLock            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMove            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodReport          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodJson            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCook            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPut, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodHead, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodACL, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLock, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMove, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodReport, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodJson, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCook, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4 AttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6 AttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4, AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF                AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupVerticalParamsNormalization string

const (
	AttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage AttackLayer7TimeseriesGroupVerticalParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupVerticalParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupVerticalParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage, AttackLayer7TimeseriesGroupVerticalParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupVerticalResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupVerticalResponseEnvelope]
type attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
