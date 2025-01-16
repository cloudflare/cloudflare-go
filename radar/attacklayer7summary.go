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

// AttackLayer7SummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7SummaryService] method instead.
type AttackLayer7SummaryService struct {
	Options []option.RequestOption
}

// NewAttackLayer7SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer7SummaryService(opts ...option.RequestOption) (r *AttackLayer7SummaryService) {
	r = &AttackLayer7SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of mitigation techniques in Layer 7 attacks.
func (r *AttackLayer7SummaryService) Get(ctx context.Context, query AttackLayer7SummaryGetParams, opts ...option.RequestOption) (res *AttackLayer7SummaryGetResponse, err error) {
	var env AttackLayer7SummaryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http method used.
func (r *AttackLayer7SummaryService) HTTPMethod(ctx context.Context, query AttackLayer7SummaryHTTPMethodParams, opts ...option.RequestOption) (res *AttackLayer7SummaryHTTPMethodResponse, err error) {
	var env AttackLayer7SummaryHTTPMethodResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by http version used.
func (r *AttackLayer7SummaryService) HTTPVersion(ctx context.Context, query AttackLayer7SummaryHTTPVersionParams, opts ...option.RequestOption) (res *AttackLayer7SummaryHTTPVersionResponse, err error) {
	var env AttackLayer7SummaryHTTPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by ip version used.
func (r *AttackLayer7SummaryService) IPVersion(ctx context.Context, query AttackLayer7SummaryIPVersionParams, opts ...option.RequestOption) (res *AttackLayer7SummaryIPVersionResponse, err error) {
	var env AttackLayer7SummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by managed rules used.
func (r *AttackLayer7SummaryService) ManagedRules(ctx context.Context, query AttackLayer7SummaryManagedRulesParams, opts ...option.RequestOption) (res *AttackLayer7SummaryManagedRulesResponse, err error) {
	var env AttackLayer7SummaryManagedRulesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of attacks by mitigation product used.
func (r *AttackLayer7SummaryService) MitigationProduct(ctx context.Context, query AttackLayer7SummaryMitigationProductParams, opts ...option.RequestOption) (res *AttackLayer7SummaryMitigationProductResponse, err error) {
	var env AttackLayer7SummaryMitigationProductResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7SummaryGetResponse struct {
	Meta     AttackLayer7SummaryGetResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryGetResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryGetResponseJSON     `json:"-"`
}

// attackLayer7SummaryGetResponseJSON contains the JSON metadata for the struct
// [AttackLayer7SummaryGetResponse]
type attackLayer7SummaryGetResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryGetResponseMeta struct {
	DateRange      []AttackLayer7SummaryGetResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo AttackLayer7SummaryGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7SummaryGetResponseMetaJSON           `json:"-"`
}

// attackLayer7SummaryGetResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer7SummaryGetResponseMeta]
type attackLayer7SummaryGetResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryGetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryGetResponseMetaDateRange]
type attackLayer7SummaryGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryGetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        attackLayer7SummaryGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7SummaryGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryGetResponseMetaConfidenceInfo]
type attackLayer7SummaryGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            attackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryGetResponseSummary0 struct {
	AccessRules        string                                     `json:"ACCESS_RULES,required"`
	APIShield          string                                     `json:"API_SHIELD,required"`
	BotManagement      string                                     `json:"BOT_MANAGEMENT,required"`
	DataLossPrevention string                                     `json:"DATA_LOSS_PREVENTION,required"`
	DDoS               string                                     `json:"DDOS,required"`
	IPReputation       string                                     `json:"IP_REPUTATION,required"`
	WAF                string                                     `json:"WAF,required"`
	JSON               attackLayer7SummaryGetResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryGetResponseSummary0JSON contains the JSON metadata for the
// struct [AttackLayer7SummaryGetResponseSummary0]
type attackLayer7SummaryGetResponseSummary0JSON struct {
	AccessRules        apijson.Field
	APIShield          apijson.Field
	BotManagement      apijson.Field
	DataLossPrevention apijson.Field
	DDoS               apijson.Field
	IPReputation       apijson.Field
	WAF                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AttackLayer7SummaryGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponse struct {
	Meta     AttackLayer7SummaryHTTPMethodResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryHTTPMethodResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryHTTPMethodResponseJSON     `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryHTTPMethodResponse]
type attackLayer7SummaryHTTPMethodResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseMeta struct {
	DateRange      []AttackLayer7SummaryHTTPMethodResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	Normalization  string                                                  `json:"normalization,required"`
	ConfidenceInfo AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7SummaryHTTPMethodResponseMetaJSON           `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryHTTPMethodResponseMeta]
type attackLayer7SummaryHTTPMethodResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryHTTPMethodResponseMetaDateRange]
type attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo]
type attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseSummary0 struct {
	Get  string                                            `json:"GET,required"`
	Post string                                            `json:"POST,required"`
	JSON attackLayer7SummaryHTTPMethodResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer7SummaryHTTPMethodResponseSummary0]
type attackLayer7SummaryHTTPMethodResponseSummary0JSON struct {
	Get         apijson.Field
	Post        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponse struct {
	Meta     AttackLayer7SummaryHTTPVersionResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryHTTPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryHTTPVersionResponseJSON     `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryHTTPVersionResponse]
type attackLayer7SummaryHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseMeta struct {
	DateRange      []AttackLayer7SummaryHTTPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7SummaryHTTPVersionResponseMetaJSON           `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryHTTPVersionResponseMeta]
type attackLayer7SummaryHTTPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryHTTPVersionResponseMetaDateRange]
type attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo]
type attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseSummary0 struct {
	HTTP1X string                                             `json:"HTTP/1.x,required"`
	HTTP2  string                                             `json:"HTTP/2,required"`
	HTTP3  string                                             `json:"HTTP/3,required"`
	JSON   attackLayer7SummaryHTTPVersionResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseSummary0JSON contains the JSON metadata
// for the struct [AttackLayer7SummaryHTTPVersionResponseSummary0]
type attackLayer7SummaryHTTPVersionResponseSummary0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponse struct {
	Meta     AttackLayer7SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryIPVersionResponseJSON     `json:"-"`
}

// attackLayer7SummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryIPVersionResponse]
type attackLayer7SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseMeta struct {
	DateRange      []AttackLayer7SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryIPVersionResponseMeta]
type attackLayer7SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryIPVersionResponseMetaDateRange]
type attackLayer7SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo]
type attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                           `json:"IPv4,required"`
	IPv6 string                                           `json:"IPv6,required"`
	JSON attackLayer7SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIPVersionResponseSummary0]
type attackLayer7SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponse struct {
	Meta     AttackLayer7SummaryManagedRulesResponseMeta `json:"meta,required"`
	Summary0 map[string][]string                         `json:"summary_0,required"`
	JSON     attackLayer7SummaryManagedRulesResponseJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryManagedRulesResponse]
type attackLayer7SummaryManagedRulesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponseMeta struct {
	DateRange      []AttackLayer7SummaryManagedRulesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	Normalization  string                                                    `json:"normalization,required"`
	ConfidenceInfo AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7SummaryManagedRulesResponseMetaJSON           `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryManagedRulesResponseMeta]
type attackLayer7SummaryManagedRulesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryManagedRulesResponseMetaDateRange]
type attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo]
type attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponse struct {
	Meta     AttackLayer7SummaryMitigationProductResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryMitigationProductResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryMitigationProductResponseJSON     `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryMitigationProductResponse]
type attackLayer7SummaryMitigationProductResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseMeta struct {
	DateRange      []AttackLayer7SummaryMitigationProductResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	Normalization  string                                                         `json:"normalization,required"`
	ConfidenceInfo AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7SummaryMitigationProductResponseMetaJSON           `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryMitigationProductResponseMeta]
type attackLayer7SummaryMitigationProductResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryMitigationProductResponseMetaDateRange]
type attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo]
type attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseSummary0 struct {
	DDoS string                                                   `json:"DDOS,required"`
	WAF  string                                                   `json:"WAF,required"`
	JSON attackLayer7SummaryMitigationProductResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseSummary0JSON contains the JSON
// metadata for the struct [AttackLayer7SummaryMitigationProductResponseSummary0]
type attackLayer7SummaryMitigationProductResponseSummary0JSON struct {
	DDoS        apijson.Field
	WAF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryGetParams struct {
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
	Format param.Field[AttackLayer7SummaryGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryGetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer7SummaryGetParamsFormat string

const (
	AttackLayer7SummaryGetParamsFormatJson AttackLayer7SummaryGetParamsFormat = "JSON"
	AttackLayer7SummaryGetParamsFormatCsv  AttackLayer7SummaryGetParamsFormat = "CSV"
)

func (r AttackLayer7SummaryGetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryGetParamsFormatJson, AttackLayer7SummaryGetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryGetResponseEnvelope struct {
	Result  AttackLayer7SummaryGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer7SummaryGetResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryGetResponseEnvelope]
type attackLayer7SummaryGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodParams struct {
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
	Format param.Field[AttackLayer7SummaryHTTPMethodParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7SummaryHTTPMethodParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7SummaryHTTPMethodParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7SummaryHTTPMethodParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryHTTPMethodParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryHTTPMethodParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer7SummaryHTTPMethodParamsFormat string

const (
	AttackLayer7SummaryHTTPMethodParamsFormatJson AttackLayer7SummaryHTTPMethodParamsFormat = "JSON"
	AttackLayer7SummaryHTTPMethodParamsFormatCsv  AttackLayer7SummaryHTTPMethodParamsFormat = "CSV"
)

func (r AttackLayer7SummaryHTTPMethodParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsFormatJson, AttackLayer7SummaryHTTPMethodParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodParamsHTTPVersion string

const (
	AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1 AttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2 AttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3 AttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryHTTPMethodParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1, AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2, AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodParamsIPVersion string

const (
	AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4 AttackLayer7SummaryHTTPMethodParamsIPVersion = "IPv4"
	AttackLayer7SummaryHTTPMethodParamsIPVersionIPv6 AttackLayer7SummaryHTTPMethodParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryHTTPMethodParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4, AttackLayer7SummaryHTTPMethodParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodParamsMitigationProduct string

const (
	AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS               AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductWAF                AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "WAF"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement      AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductAccessRules        AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductIPReputation       AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductAPIShield          AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductDataLossPrevention AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryHTTPMethodParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS, AttackLayer7SummaryHTTPMethodParamsMitigationProductWAF, AttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement, AttackLayer7SummaryHTTPMethodParamsMitigationProductAccessRules, AttackLayer7SummaryHTTPMethodParamsMitigationProductIPReputation, AttackLayer7SummaryHTTPMethodParamsMitigationProductAPIShield, AttackLayer7SummaryHTTPMethodParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodResponseEnvelope struct {
	Result  AttackLayer7SummaryHTTPMethodResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer7SummaryHTTPMethodResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryHTTPMethodResponseEnvelope]
type attackLayer7SummaryHTTPMethodResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionParams struct {
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
	Format param.Field[AttackLayer7SummaryHTTPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7SummaryHTTPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7SummaryHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7SummaryHTTPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryHTTPVersionParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer7SummaryHTTPVersionParamsFormat string

const (
	AttackLayer7SummaryHTTPVersionParamsFormatJson AttackLayer7SummaryHTTPVersionParamsFormat = "JSON"
	AttackLayer7SummaryHTTPVersionParamsFormatCsv  AttackLayer7SummaryHTTPVersionParamsFormat = "CSV"
)

func (r AttackLayer7SummaryHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsFormatJson, AttackLayer7SummaryHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionParamsHTTPMethod string

const (
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet             AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "GET"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPost            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "POST"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPut             AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PUT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodHead            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPurge           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodOptions         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPropfind        AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkcol           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPatch           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodACL             AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "ACL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBcopy           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBdelete         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBmove           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBpropfind       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBproppatch      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckin         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckout        AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodConnect         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCopy            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "COPY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodLabel           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodLock            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMerge           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkactivity      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkworkspace     AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMove            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodNotify          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodOrderpatch      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPoll            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "POLL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodProppatch       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodReport          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodSearch          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodSubscribe       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrace           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUncheckout      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnlock          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnsubscribe     AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUpdate          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodVersioncontrol  AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBaselinecontrol AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodXmsenumatts     AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcOutData      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcInData       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodJson            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "JSON"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCook            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "COOK"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrack           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryHTTPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPost, AttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPut, AttackLayer7SummaryHTTPVersionParamsHTTPMethodHead, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPurge, AttackLayer7SummaryHTTPVersionParamsHTTPMethodOptions, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPropfind, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkcol, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodACL, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBcopy, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBdelete, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBmove, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBpropfind, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBproppatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckin, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckout, AttackLayer7SummaryHTTPVersionParamsHTTPMethodConnect, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCopy, AttackLayer7SummaryHTTPVersionParamsHTTPMethodLabel, AttackLayer7SummaryHTTPVersionParamsHTTPMethodLock, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMerge, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkactivity, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkworkspace, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMove, AttackLayer7SummaryHTTPVersionParamsHTTPMethodNotify, AttackLayer7SummaryHTTPVersionParamsHTTPMethodOrderpatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPoll, AttackLayer7SummaryHTTPVersionParamsHTTPMethodProppatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodReport, AttackLayer7SummaryHTTPVersionParamsHTTPMethodSearch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodSubscribe, AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrace, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUncheckout, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnlock, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnsubscribe, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUpdate, AttackLayer7SummaryHTTPVersionParamsHTTPMethodVersioncontrol, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryHTTPVersionParamsHTTPMethodXmsenumatts, AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcOutData, AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcInData, AttackLayer7SummaryHTTPVersionParamsHTTPMethodJson, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCook, AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionParamsIPVersion string

const (
	AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4 AttackLayer7SummaryHTTPVersionParamsIPVersion = "IPv4"
	AttackLayer7SummaryHTTPVersionParamsIPVersionIPv6 AttackLayer7SummaryHTTPVersionParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4, AttackLayer7SummaryHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionParamsMitigationProduct string

const (
	AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS               AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductWAF                AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "WAF"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement      AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductAccessRules        AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductIPReputation       AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductAPIShield          AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductDataLossPrevention AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryHTTPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS, AttackLayer7SummaryHTTPVersionParamsMitigationProductWAF, AttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement, AttackLayer7SummaryHTTPVersionParamsMitigationProductAccessRules, AttackLayer7SummaryHTTPVersionParamsMitigationProductIPReputation, AttackLayer7SummaryHTTPVersionParamsMitigationProductAPIShield, AttackLayer7SummaryHTTPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionResponseEnvelope struct {
	Result  AttackLayer7SummaryHTTPVersionResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    attackLayer7SummaryHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryHTTPVersionResponseEnvelope]
type attackLayer7SummaryHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionParams struct {
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
	Format param.Field[AttackLayer7SummaryIPVersionParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7SummaryIPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7SummaryIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7SummaryIPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer7SummaryIPVersionParamsFormat string

const (
	AttackLayer7SummaryIPVersionParamsFormatJson AttackLayer7SummaryIPVersionParamsFormat = "JSON"
	AttackLayer7SummaryIPVersionParamsFormatCsv  AttackLayer7SummaryIPVersionParamsFormat = "CSV"
)

func (r AttackLayer7SummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsFormatJson, AttackLayer7SummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionParamsHTTPMethod string

const (
	AttackLayer7SummaryIPVersionParamsHTTPMethodGet             AttackLayer7SummaryIPVersionParamsHTTPMethod = "GET"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPost            AttackLayer7SummaryIPVersionParamsHTTPMethod = "POST"
	AttackLayer7SummaryIPVersionParamsHTTPMethodDelete          AttackLayer7SummaryIPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPut             AttackLayer7SummaryIPVersionParamsHTTPMethod = "PUT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodHead            AttackLayer7SummaryIPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPurge           AttackLayer7SummaryIPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodOptions         AttackLayer7SummaryIPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPropfind        AttackLayer7SummaryIPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMkcol           AttackLayer7SummaryIPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPatch           AttackLayer7SummaryIPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodACL             AttackLayer7SummaryIPVersionParamsHTTPMethod = "ACL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBcopy           AttackLayer7SummaryIPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBdelete         AttackLayer7SummaryIPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBmove           AttackLayer7SummaryIPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBpropfind       AttackLayer7SummaryIPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBproppatch      AttackLayer7SummaryIPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCheckin         AttackLayer7SummaryIPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCheckout        AttackLayer7SummaryIPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodConnect         AttackLayer7SummaryIPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCopy            AttackLayer7SummaryIPVersionParamsHTTPMethod = "COPY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodLabel           AttackLayer7SummaryIPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodLock            AttackLayer7SummaryIPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMerge           AttackLayer7SummaryIPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMkactivity      AttackLayer7SummaryIPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMkworkspace     AttackLayer7SummaryIPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMove            AttackLayer7SummaryIPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodNotify          AttackLayer7SummaryIPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodOrderpatch      AttackLayer7SummaryIPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPoll            AttackLayer7SummaryIPVersionParamsHTTPMethod = "POLL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodProppatch       AttackLayer7SummaryIPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodReport          AttackLayer7SummaryIPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodSearch          AttackLayer7SummaryIPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodSubscribe       AttackLayer7SummaryIPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodTrace           AttackLayer7SummaryIPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUncheckout      AttackLayer7SummaryIPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUnlock          AttackLayer7SummaryIPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUnsubscribe     AttackLayer7SummaryIPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUpdate          AttackLayer7SummaryIPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodVersioncontrol  AttackLayer7SummaryIPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBaselinecontrol AttackLayer7SummaryIPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodXmsenumatts     AttackLayer7SummaryIPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryIPVersionParamsHTTPMethodRpcOutData      AttackLayer7SummaryIPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryIPVersionParamsHTTPMethodRpcInData       AttackLayer7SummaryIPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryIPVersionParamsHTTPMethodJson            AttackLayer7SummaryIPVersionParamsHTTPMethod = "JSON"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCook            AttackLayer7SummaryIPVersionParamsHTTPMethod = "COOK"
	AttackLayer7SummaryIPVersionParamsHTTPMethodTrack           AttackLayer7SummaryIPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryIPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsHTTPMethodGet, AttackLayer7SummaryIPVersionParamsHTTPMethodPost, AttackLayer7SummaryIPVersionParamsHTTPMethodDelete, AttackLayer7SummaryIPVersionParamsHTTPMethodPut, AttackLayer7SummaryIPVersionParamsHTTPMethodHead, AttackLayer7SummaryIPVersionParamsHTTPMethodPurge, AttackLayer7SummaryIPVersionParamsHTTPMethodOptions, AttackLayer7SummaryIPVersionParamsHTTPMethodPropfind, AttackLayer7SummaryIPVersionParamsHTTPMethodMkcol, AttackLayer7SummaryIPVersionParamsHTTPMethodPatch, AttackLayer7SummaryIPVersionParamsHTTPMethodACL, AttackLayer7SummaryIPVersionParamsHTTPMethodBcopy, AttackLayer7SummaryIPVersionParamsHTTPMethodBdelete, AttackLayer7SummaryIPVersionParamsHTTPMethodBmove, AttackLayer7SummaryIPVersionParamsHTTPMethodBpropfind, AttackLayer7SummaryIPVersionParamsHTTPMethodBproppatch, AttackLayer7SummaryIPVersionParamsHTTPMethodCheckin, AttackLayer7SummaryIPVersionParamsHTTPMethodCheckout, AttackLayer7SummaryIPVersionParamsHTTPMethodConnect, AttackLayer7SummaryIPVersionParamsHTTPMethodCopy, AttackLayer7SummaryIPVersionParamsHTTPMethodLabel, AttackLayer7SummaryIPVersionParamsHTTPMethodLock, AttackLayer7SummaryIPVersionParamsHTTPMethodMerge, AttackLayer7SummaryIPVersionParamsHTTPMethodMkactivity, AttackLayer7SummaryIPVersionParamsHTTPMethodMkworkspace, AttackLayer7SummaryIPVersionParamsHTTPMethodMove, AttackLayer7SummaryIPVersionParamsHTTPMethodNotify, AttackLayer7SummaryIPVersionParamsHTTPMethodOrderpatch, AttackLayer7SummaryIPVersionParamsHTTPMethodPoll, AttackLayer7SummaryIPVersionParamsHTTPMethodProppatch, AttackLayer7SummaryIPVersionParamsHTTPMethodReport, AttackLayer7SummaryIPVersionParamsHTTPMethodSearch, AttackLayer7SummaryIPVersionParamsHTTPMethodSubscribe, AttackLayer7SummaryIPVersionParamsHTTPMethodTrace, AttackLayer7SummaryIPVersionParamsHTTPMethodUncheckout, AttackLayer7SummaryIPVersionParamsHTTPMethodUnlock, AttackLayer7SummaryIPVersionParamsHTTPMethodUnsubscribe, AttackLayer7SummaryIPVersionParamsHTTPMethodUpdate, AttackLayer7SummaryIPVersionParamsHTTPMethodVersioncontrol, AttackLayer7SummaryIPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryIPVersionParamsHTTPMethodXmsenumatts, AttackLayer7SummaryIPVersionParamsHTTPMethodRpcOutData, AttackLayer7SummaryIPVersionParamsHTTPMethodRpcInData, AttackLayer7SummaryIPVersionParamsHTTPMethodJson, AttackLayer7SummaryIPVersionParamsHTTPMethodCook, AttackLayer7SummaryIPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionParamsHTTPVersion string

const (
	AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1 AttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2 AttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3 AttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1, AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2, AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionParamsMitigationProduct string

const (
	AttackLayer7SummaryIPVersionParamsMitigationProductDDoS               AttackLayer7SummaryIPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryIPVersionParamsMitigationProductWAF                AttackLayer7SummaryIPVersionParamsMitigationProduct = "WAF"
	AttackLayer7SummaryIPVersionParamsMitigationProductBotManagement      AttackLayer7SummaryIPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryIPVersionParamsMitigationProductAccessRules        AttackLayer7SummaryIPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryIPVersionParamsMitigationProductIPReputation       AttackLayer7SummaryIPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryIPVersionParamsMitigationProductAPIShield          AttackLayer7SummaryIPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryIPVersionParamsMitigationProductDataLossPrevention AttackLayer7SummaryIPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryIPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsMitigationProductDDoS, AttackLayer7SummaryIPVersionParamsMitigationProductWAF, AttackLayer7SummaryIPVersionParamsMitigationProductBotManagement, AttackLayer7SummaryIPVersionParamsMitigationProductAccessRules, AttackLayer7SummaryIPVersionParamsMitigationProductIPReputation, AttackLayer7SummaryIPVersionParamsMitigationProductAPIShield, AttackLayer7SummaryIPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionResponseEnvelope struct {
	Result  AttackLayer7SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    attackLayer7SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIPVersionResponseEnvelope]
type attackLayer7SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesParams struct {
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
	Format param.Field[AttackLayer7SummaryManagedRulesParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7SummaryManagedRulesParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7SummaryManagedRulesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7SummaryManagedRulesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]AttackLayer7SummaryManagedRulesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryManagedRulesParams]'s query parameters
// as `url.Values`.
func (r AttackLayer7SummaryManagedRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer7SummaryManagedRulesParamsFormat string

const (
	AttackLayer7SummaryManagedRulesParamsFormatJson AttackLayer7SummaryManagedRulesParamsFormat = "JSON"
	AttackLayer7SummaryManagedRulesParamsFormatCsv  AttackLayer7SummaryManagedRulesParamsFormat = "CSV"
)

func (r AttackLayer7SummaryManagedRulesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsFormatJson, AttackLayer7SummaryManagedRulesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsHTTPMethod string

const (
	AttackLayer7SummaryManagedRulesParamsHTTPMethodGet             AttackLayer7SummaryManagedRulesParamsHTTPMethod = "GET"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPost            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "POST"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodDelete          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPut             AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PUT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodHead            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPurge           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodOptions         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPropfind        AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMkcol           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPatch           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodACL             AttackLayer7SummaryManagedRulesParamsHTTPMethod = "ACL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBcopy           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBdelete         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBmove           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBpropfind       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBproppatch      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckin         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckout        AttackLayer7SummaryManagedRulesParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodConnect         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCopy            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "COPY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodLabel           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodLock            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMerge           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMkactivity      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMkworkspace     AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMove            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodNotify          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodOrderpatch      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPoll            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "POLL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodProppatch       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodReport          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodSearch          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodSubscribe       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodTrace           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUncheckout      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUnlock          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUnsubscribe     AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUpdate          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodVersioncontrol  AttackLayer7SummaryManagedRulesParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBaselinecontrol AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodXmsenumatts     AttackLayer7SummaryManagedRulesParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcOutData      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcInData       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodJson            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "JSON"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCook            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "COOK"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodTrack           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryManagedRulesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsHTTPMethodGet, AttackLayer7SummaryManagedRulesParamsHTTPMethodPost, AttackLayer7SummaryManagedRulesParamsHTTPMethodDelete, AttackLayer7SummaryManagedRulesParamsHTTPMethodPut, AttackLayer7SummaryManagedRulesParamsHTTPMethodHead, AttackLayer7SummaryManagedRulesParamsHTTPMethodPurge, AttackLayer7SummaryManagedRulesParamsHTTPMethodOptions, AttackLayer7SummaryManagedRulesParamsHTTPMethodPropfind, AttackLayer7SummaryManagedRulesParamsHTTPMethodMkcol, AttackLayer7SummaryManagedRulesParamsHTTPMethodPatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodACL, AttackLayer7SummaryManagedRulesParamsHTTPMethodBcopy, AttackLayer7SummaryManagedRulesParamsHTTPMethodBdelete, AttackLayer7SummaryManagedRulesParamsHTTPMethodBmove, AttackLayer7SummaryManagedRulesParamsHTTPMethodBpropfind, AttackLayer7SummaryManagedRulesParamsHTTPMethodBproppatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckin, AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckout, AttackLayer7SummaryManagedRulesParamsHTTPMethodConnect, AttackLayer7SummaryManagedRulesParamsHTTPMethodCopy, AttackLayer7SummaryManagedRulesParamsHTTPMethodLabel, AttackLayer7SummaryManagedRulesParamsHTTPMethodLock, AttackLayer7SummaryManagedRulesParamsHTTPMethodMerge, AttackLayer7SummaryManagedRulesParamsHTTPMethodMkactivity, AttackLayer7SummaryManagedRulesParamsHTTPMethodMkworkspace, AttackLayer7SummaryManagedRulesParamsHTTPMethodMove, AttackLayer7SummaryManagedRulesParamsHTTPMethodNotify, AttackLayer7SummaryManagedRulesParamsHTTPMethodOrderpatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodPoll, AttackLayer7SummaryManagedRulesParamsHTTPMethodProppatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodReport, AttackLayer7SummaryManagedRulesParamsHTTPMethodSearch, AttackLayer7SummaryManagedRulesParamsHTTPMethodSubscribe, AttackLayer7SummaryManagedRulesParamsHTTPMethodTrace, AttackLayer7SummaryManagedRulesParamsHTTPMethodUncheckout, AttackLayer7SummaryManagedRulesParamsHTTPMethodUnlock, AttackLayer7SummaryManagedRulesParamsHTTPMethodUnsubscribe, AttackLayer7SummaryManagedRulesParamsHTTPMethodUpdate, AttackLayer7SummaryManagedRulesParamsHTTPMethodVersioncontrol, AttackLayer7SummaryManagedRulesParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryManagedRulesParamsHTTPMethodXmsenumatts, AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcOutData, AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcInData, AttackLayer7SummaryManagedRulesParamsHTTPMethodJson, AttackLayer7SummaryManagedRulesParamsHTTPMethodCook, AttackLayer7SummaryManagedRulesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsHTTPVersion string

const (
	AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1 AttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2 AttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3 AttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryManagedRulesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1, AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2, AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsIPVersion string

const (
	AttackLayer7SummaryManagedRulesParamsIPVersionIPv4 AttackLayer7SummaryManagedRulesParamsIPVersion = "IPv4"
	AttackLayer7SummaryManagedRulesParamsIPVersionIPv6 AttackLayer7SummaryManagedRulesParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryManagedRulesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsIPVersionIPv4, AttackLayer7SummaryManagedRulesParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsMitigationProduct string

const (
	AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS               AttackLayer7SummaryManagedRulesParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryManagedRulesParamsMitigationProductWAF                AttackLayer7SummaryManagedRulesParamsMitigationProduct = "WAF"
	AttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement      AttackLayer7SummaryManagedRulesParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryManagedRulesParamsMitigationProductAccessRules        AttackLayer7SummaryManagedRulesParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryManagedRulesParamsMitigationProductIPReputation       AttackLayer7SummaryManagedRulesParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryManagedRulesParamsMitigationProductAPIShield          AttackLayer7SummaryManagedRulesParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryManagedRulesParamsMitigationProductDataLossPrevention AttackLayer7SummaryManagedRulesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryManagedRulesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS, AttackLayer7SummaryManagedRulesParamsMitigationProductWAF, AttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement, AttackLayer7SummaryManagedRulesParamsMitigationProductAccessRules, AttackLayer7SummaryManagedRulesParamsMitigationProductIPReputation, AttackLayer7SummaryManagedRulesParamsMitigationProductAPIShield, AttackLayer7SummaryManagedRulesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesResponseEnvelope struct {
	Result  AttackLayer7SummaryManagedRulesResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    attackLayer7SummaryManagedRulesResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryManagedRulesResponseEnvelope]
type attackLayer7SummaryManagedRulesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductParams struct {
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
	Format param.Field[AttackLayer7SummaryMitigationProductParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]AttackLayer7SummaryMitigationProductParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]AttackLayer7SummaryMitigationProductParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer7SummaryMitigationProductParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryMitigationProductParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7SummaryMitigationProductParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AttackLayer7SummaryMitigationProductParamsFormat string

const (
	AttackLayer7SummaryMitigationProductParamsFormatJson AttackLayer7SummaryMitigationProductParamsFormat = "JSON"
	AttackLayer7SummaryMitigationProductParamsFormatCsv  AttackLayer7SummaryMitigationProductParamsFormat = "CSV"
)

func (r AttackLayer7SummaryMitigationProductParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsFormatJson, AttackLayer7SummaryMitigationProductParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductParamsHTTPMethod string

const (
	AttackLayer7SummaryMitigationProductParamsHTTPMethodGet             AttackLayer7SummaryMitigationProductParamsHTTPMethod = "GET"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPost            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "POST"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodDelete          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPut             AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PUT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodHead            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPurge           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodOptions         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPropfind        AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMkcol           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPatch           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodACL             AttackLayer7SummaryMitigationProductParamsHTTPMethod = "ACL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBcopy           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBdelete         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBmove           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBpropfind       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBproppatch      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckin         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckout        AttackLayer7SummaryMitigationProductParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodConnect         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCopy            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "COPY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodLabel           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodLock            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMerge           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMkactivity      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMkworkspace     AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMove            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodNotify          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodOrderpatch      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPoll            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "POLL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodProppatch       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodReport          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodSearch          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodSubscribe       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodTrace           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUncheckout      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUnlock          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUnsubscribe     AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUpdate          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodVersioncontrol  AttackLayer7SummaryMitigationProductParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBaselinecontrol AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodXmsenumatts     AttackLayer7SummaryMitigationProductParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcOutData      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcInData       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodJson            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "JSON"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCook            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "COOK"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodTrack           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryMitigationProductParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsHTTPMethodGet, AttackLayer7SummaryMitigationProductParamsHTTPMethodPost, AttackLayer7SummaryMitigationProductParamsHTTPMethodDelete, AttackLayer7SummaryMitigationProductParamsHTTPMethodPut, AttackLayer7SummaryMitigationProductParamsHTTPMethodHead, AttackLayer7SummaryMitigationProductParamsHTTPMethodPurge, AttackLayer7SummaryMitigationProductParamsHTTPMethodOptions, AttackLayer7SummaryMitigationProductParamsHTTPMethodPropfind, AttackLayer7SummaryMitigationProductParamsHTTPMethodMkcol, AttackLayer7SummaryMitigationProductParamsHTTPMethodPatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodACL, AttackLayer7SummaryMitigationProductParamsHTTPMethodBcopy, AttackLayer7SummaryMitigationProductParamsHTTPMethodBdelete, AttackLayer7SummaryMitigationProductParamsHTTPMethodBmove, AttackLayer7SummaryMitigationProductParamsHTTPMethodBpropfind, AttackLayer7SummaryMitigationProductParamsHTTPMethodBproppatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckin, AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckout, AttackLayer7SummaryMitigationProductParamsHTTPMethodConnect, AttackLayer7SummaryMitigationProductParamsHTTPMethodCopy, AttackLayer7SummaryMitigationProductParamsHTTPMethodLabel, AttackLayer7SummaryMitigationProductParamsHTTPMethodLock, AttackLayer7SummaryMitigationProductParamsHTTPMethodMerge, AttackLayer7SummaryMitigationProductParamsHTTPMethodMkactivity, AttackLayer7SummaryMitigationProductParamsHTTPMethodMkworkspace, AttackLayer7SummaryMitigationProductParamsHTTPMethodMove, AttackLayer7SummaryMitigationProductParamsHTTPMethodNotify, AttackLayer7SummaryMitigationProductParamsHTTPMethodOrderpatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodPoll, AttackLayer7SummaryMitigationProductParamsHTTPMethodProppatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodReport, AttackLayer7SummaryMitigationProductParamsHTTPMethodSearch, AttackLayer7SummaryMitigationProductParamsHTTPMethodSubscribe, AttackLayer7SummaryMitigationProductParamsHTTPMethodTrace, AttackLayer7SummaryMitigationProductParamsHTTPMethodUncheckout, AttackLayer7SummaryMitigationProductParamsHTTPMethodUnlock, AttackLayer7SummaryMitigationProductParamsHTTPMethodUnsubscribe, AttackLayer7SummaryMitigationProductParamsHTTPMethodUpdate, AttackLayer7SummaryMitigationProductParamsHTTPMethodVersioncontrol, AttackLayer7SummaryMitigationProductParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryMitigationProductParamsHTTPMethodXmsenumatts, AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcOutData, AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcInData, AttackLayer7SummaryMitigationProductParamsHTTPMethodJson, AttackLayer7SummaryMitigationProductParamsHTTPMethodCook, AttackLayer7SummaryMitigationProductParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductParamsHTTPVersion string

const (
	AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1 AttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2 AttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3 AttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryMitigationProductParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1, AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2, AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductParamsIPVersion string

const (
	AttackLayer7SummaryMitigationProductParamsIPVersionIPv4 AttackLayer7SummaryMitigationProductParamsIPVersion = "IPv4"
	AttackLayer7SummaryMitigationProductParamsIPVersionIPv6 AttackLayer7SummaryMitigationProductParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryMitigationProductParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsIPVersionIPv4, AttackLayer7SummaryMitigationProductParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductResponseEnvelope struct {
	Result  AttackLayer7SummaryMitigationProductResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    attackLayer7SummaryMitigationProductResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryMitigationProductResponseEnvelope]
type attackLayer7SummaryMitigationProductResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
