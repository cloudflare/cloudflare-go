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

// AS112SummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112SummaryService] method instead.
type AS112SummaryService struct {
	Options []option.RequestOption
}

// NewAS112SummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAS112SummaryService(opts ...option.RequestOption) (r *AS112SummaryService) {
	r = &AS112SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS queries to AS112 by DNSSEC support.
func (r *AS112SummaryService) DNSSEC(ctx context.Context, query AS112SummaryDNSSECParams, opts ...option.RequestOption) (res *AS112SummaryDNSSECResponse, err error) {
	var env AS112SummaryDNSSECResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 by EDNS support.
func (r *AS112SummaryService) Edns(ctx context.Context, query AS112SummaryEdnsParams, opts ...option.RequestOption) (res *AS112SummaryEdnsResponse, err error) {
	var env AS112SummaryEdnsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 per IP Version.
func (r *AS112SummaryService) IPVersion(ctx context.Context, query AS112SummaryIPVersionParams, opts ...option.RequestOption) (res *AS112SummaryIPVersionResponse, err error) {
	var env AS112SummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 per protocol.
func (r *AS112SummaryService) Protocol(ctx context.Context, query AS112SummaryProtocolParams, opts ...option.RequestOption) (res *AS112SummaryProtocolResponse, err error) {
	var env AS112SummaryProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 by query type.
func (r *AS112SummaryService) QueryType(ctx context.Context, query AS112SummaryQueryTypeParams, opts ...option.RequestOption) (res *AS112SummaryQueryTypeResponse, err error) {
	var env AS112SummaryQueryTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS requests classified by response code.
func (r *AS112SummaryService) ResponseCodes(ctx context.Context, query AS112SummaryResponseCodesParams, opts ...option.RequestOption) (res *AS112SummaryResponseCodesResponse, err error) {
	var env AS112SummaryResponseCodesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112SummaryDNSSECResponse struct {
	Meta     AS112SummaryDNSSECResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryDNSSECResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryDNSSECResponseJSON     `json:"-"`
}

// as112SummaryDNSSECResponseJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponse]
type as112SummaryDNSSECResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseMeta struct {
	DateRange      []AS112SummaryDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                       `json:"lastUpdated,required"`
	Normalization  string                                       `json:"normalization,required"`
	ConfidenceInfo AS112SummaryDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112SummaryDNSSECResponseMetaJSON           `json:"-"`
}

// as112SummaryDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseMeta]
type as112SummaryDNSSECResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryDNSSECResponseMetaDateRange]
type as112SummaryDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        as112SummaryDNSSECResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112SummaryDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112SummaryDNSSECResponseMetaConfidenceInfo]
type as112SummaryDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                     `json:"dataSource,required"`
	Description     string                                                     `json:"description,required"`
	EventType       string                                                     `json:"eventType,required"`
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                  `json:"startTime" format:"date-time"`
	JSON            as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation]
type as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseSummary0 struct {
	NotSupported string                                 `json:"NOT_SUPPORTED,required"`
	Supported    string                                 `json:"SUPPORTED,required"`
	JSON         as112SummaryDNSSECResponseSummary0JSON `json:"-"`
}

// as112SummaryDNSSECResponseSummary0JSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseSummary0]
type as112SummaryDNSSECResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponse struct {
	Meta     AS112SummaryEdnsResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryEdnsResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryEdnsResponseJSON     `json:"-"`
}

// as112SummaryEdnsResponseJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponse]
type as112SummaryEdnsResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseMeta struct {
	DateRange      []AS112SummaryEdnsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                     `json:"lastUpdated,required"`
	Normalization  string                                     `json:"normalization,required"`
	ConfidenceInfo AS112SummaryEdnsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112SummaryEdnsResponseMetaJSON           `json:"-"`
}

// as112SummaryEdnsResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseMeta]
type as112SummaryEdnsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                 `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryEdnsResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryEdnsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryEdnsResponseMetaDateRange]
type as112SummaryEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                  `json:"level"`
	JSON        as112SummaryEdnsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112SummaryEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112SummaryEdnsResponseMetaConfidenceInfo]
type as112SummaryEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                   `json:"dataSource,required"`
	Description     string                                                   `json:"description,required"`
	EventType       string                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                `json:"startTime" format:"date-time"`
	JSON            as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation]
type as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseSummary0 struct {
	NotSupported string                               `json:"NOT_SUPPORTED,required"`
	Supported    string                               `json:"SUPPORTED,required"`
	JSON         as112SummaryEdnsResponseSummary0JSON `json:"-"`
}

// as112SummaryEdnsResponseSummary0JSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseSummary0]
type as112SummaryEdnsResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponse struct {
	Meta     AS112SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryIPVersionResponseJSON     `json:"-"`
}

// as112SummaryIPVersionResponseJSON contains the JSON metadata for the struct
// [AS112SummaryIPVersionResponse]
type as112SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseMeta struct {
	DateRange      []AS112SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                          `json:"lastUpdated,required"`
	Normalization  string                                          `json:"normalization,required"`
	ConfidenceInfo AS112SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// as112SummaryIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryIPVersionResponseMeta]
type as112SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AS112SummaryIPVersionResponseMetaDateRange]
type as112SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        as112SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112SummaryIPVersionResponseMetaConfidenceInfo]
type as112SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                        `json:"dataSource,required"`
	Description     string                                                        `json:"description,required"`
	EventType       string                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                     `json:"startTime" format:"date-time"`
	JSON            as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                    `json:"IPv4,required"`
	IPv6 string                                    `json:"IPv6,required"`
	JSON as112SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// as112SummaryIPVersionResponseSummary0JSON contains the JSON metadata for the
// struct [AS112SummaryIPVersionResponseSummary0]
type as112SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponse struct {
	Meta     AS112SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryProtocolResponseJSON     `json:"-"`
}

// as112SummaryProtocolResponseJSON contains the JSON metadata for the struct
// [AS112SummaryProtocolResponse]
type as112SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseMeta struct {
	DateRange      []AS112SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	Normalization  string                                         `json:"normalization,required"`
	ConfidenceInfo AS112SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112SummaryProtocolResponseMetaJSON           `json:"-"`
}

// as112SummaryProtocolResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryProtocolResponseMeta]
type as112SummaryProtocolResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseMetaDateRange]
type as112SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        as112SummaryProtocolResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112SummaryProtocolResponseMetaConfidenceInfo]
type as112SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseSummary0 struct {
	TCP  string                                   `json:"tcp,required"`
	Udp  string                                   `json:"udp,required"`
	JSON as112SummaryProtocolResponseSummary0JSON `json:"-"`
}

// as112SummaryProtocolResponseSummary0JSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseSummary0]
type as112SummaryProtocolResponseSummary0JSON struct {
	TCP         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponse struct {
	Meta     AS112SummaryQueryTypeResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryQueryTypeResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryQueryTypeResponseJSON     `json:"-"`
}

// as112SummaryQueryTypeResponseJSON contains the JSON metadata for the struct
// [AS112SummaryQueryTypeResponse]
type as112SummaryQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseMeta struct {
	DateRange      []AS112SummaryQueryTypeResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                          `json:"lastUpdated,required"`
	Normalization  string                                          `json:"normalization,required"`
	ConfidenceInfo AS112SummaryQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112SummaryQueryTypeResponseMetaJSON           `json:"-"`
}

// as112SummaryQueryTypeResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryQueryTypeResponseMeta]
type as112SummaryQueryTypeResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryQueryTypeResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AS112SummaryQueryTypeResponseMetaDateRange]
type as112SummaryQueryTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        as112SummaryQueryTypeResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112SummaryQueryTypeResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112SummaryQueryTypeResponseMetaConfidenceInfo]
type as112SummaryQueryTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                        `json:"dataSource,required"`
	Description     string                                                        `json:"description,required"`
	EventType       string                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                     `json:"startTime" format:"date-time"`
	JSON            as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation]
type as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseSummary0 struct {
	A    string                                    `json:"A,required"`
	AAAA string                                    `json:"AAAA,required"`
	PTR  string                                    `json:"PTR,required"`
	SOA  string                                    `json:"SOA,required"`
	SRV  string                                    `json:"SRV,required"`
	JSON as112SummaryQueryTypeResponseSummary0JSON `json:"-"`
}

// as112SummaryQueryTypeResponseSummary0JSON contains the JSON metadata for the
// struct [AS112SummaryQueryTypeResponseSummary0]
type as112SummaryQueryTypeResponseSummary0JSON struct {
	A           apijson.Field
	AAAA        apijson.Field
	PTR         apijson.Field
	SOA         apijson.Field
	SRV         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponse struct {
	Meta     AS112SummaryResponseCodesResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryResponseCodesResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryResponseCodesResponseJSON     `json:"-"`
}

// as112SummaryResponseCodesResponseJSON contains the JSON metadata for the struct
// [AS112SummaryResponseCodesResponse]
type as112SummaryResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseMeta struct {
	DateRange      []AS112SummaryResponseCodesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	Normalization  string                                              `json:"normalization,required"`
	ConfidenceInfo AS112SummaryResponseCodesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           as112SummaryResponseCodesResponseMetaJSON           `json:"-"`
}

// as112SummaryResponseCodesResponseMetaJSON contains the JSON metadata for the
// struct [AS112SummaryResponseCodesResponseMeta]
type as112SummaryResponseCodesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryResponseCodesResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryResponseCodesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AS112SummaryResponseCodesResponseMetaDateRange]
type as112SummaryResponseCodesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        as112SummaryResponseCodesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// as112SummaryResponseCodesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AS112SummaryResponseCodesResponseMetaConfidenceInfo]
type as112SummaryResponseCodesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation]
type as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseSummary0 struct {
	Noerror  string                                        `json:"NOERROR,required"`
	Nxdomain string                                        `json:"NXDOMAIN,required"`
	JSON     as112SummaryResponseCodesResponseSummary0JSON `json:"-"`
}

// as112SummaryResponseCodesResponseSummary0JSON contains the JSON metadata for the
// struct [AS112SummaryResponseCodesResponseSummary0]
type as112SummaryResponseCodesResponseSummary0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECParams struct {
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
	Format param.Field[AS112SummaryDNSSECParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112SummaryDNSSECParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AS112SummaryDNSSECParamsFormat string

const (
	AS112SummaryDNSSECParamsFormatJson AS112SummaryDNSSECParamsFormat = "JSON"
	AS112SummaryDNSSECParamsFormatCsv  AS112SummaryDNSSECParamsFormat = "CSV"
)

func (r AS112SummaryDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryDNSSECParamsFormatJson, AS112SummaryDNSSECParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryDNSSECResponseEnvelope struct {
	Result  AS112SummaryDNSSECResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    as112SummaryDNSSECResponseEnvelopeJSON `json:"-"`
}

// as112SummaryDNSSECResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseEnvelope]
type as112SummaryDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsParams struct {
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
	Format param.Field[AS112SummaryEdnsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112SummaryEdnsParams]'s query parameters as `url.Values`.
func (r AS112SummaryEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AS112SummaryEdnsParamsFormat string

const (
	AS112SummaryEdnsParamsFormatJson AS112SummaryEdnsParamsFormat = "JSON"
	AS112SummaryEdnsParamsFormatCsv  AS112SummaryEdnsParamsFormat = "CSV"
)

func (r AS112SummaryEdnsParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryEdnsParamsFormatJson, AS112SummaryEdnsParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryEdnsResponseEnvelope struct {
	Result  AS112SummaryEdnsResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    as112SummaryEdnsResponseEnvelopeJSON `json:"-"`
}

// as112SummaryEdnsResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseEnvelope]
type as112SummaryEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionParams struct {
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
	Format param.Field[AS112SummaryIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AS112SummaryIPVersionParamsFormat string

const (
	AS112SummaryIPVersionParamsFormatJson AS112SummaryIPVersionParamsFormat = "JSON"
	AS112SummaryIPVersionParamsFormatCsv  AS112SummaryIPVersionParamsFormat = "CSV"
)

func (r AS112SummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryIPVersionParamsFormatJson, AS112SummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryIPVersionResponseEnvelope struct {
	Result  AS112SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    as112SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// as112SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryIPVersionResponseEnvelope]
type as112SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolParams struct {
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
	Format param.Field[AS112SummaryProtocolParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112SummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AS112SummaryProtocolParamsFormat string

const (
	AS112SummaryProtocolParamsFormatJson AS112SummaryProtocolParamsFormat = "JSON"
	AS112SummaryProtocolParamsFormatCsv  AS112SummaryProtocolParamsFormat = "CSV"
)

func (r AS112SummaryProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryProtocolParamsFormatJson, AS112SummaryProtocolParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryProtocolResponseEnvelope struct {
	Result  AS112SummaryProtocolResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    as112SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// as112SummaryProtocolResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseEnvelope]
type as112SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeParams struct {
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
	Format param.Field[AS112SummaryQueryTypeParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112SummaryQueryTypeParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AS112SummaryQueryTypeParamsFormat string

const (
	AS112SummaryQueryTypeParamsFormatJson AS112SummaryQueryTypeParamsFormat = "JSON"
	AS112SummaryQueryTypeParamsFormatCsv  AS112SummaryQueryTypeParamsFormat = "CSV"
)

func (r AS112SummaryQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryQueryTypeParamsFormatJson, AS112SummaryQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryQueryTypeResponseEnvelope struct {
	Result  AS112SummaryQueryTypeResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    as112SummaryQueryTypeResponseEnvelopeJSON `json:"-"`
}

// as112SummaryQueryTypeResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryQueryTypeResponseEnvelope]
type as112SummaryQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesParams struct {
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
	Format param.Field[AS112SummaryResponseCodesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112SummaryResponseCodesParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type AS112SummaryResponseCodesParamsFormat string

const (
	AS112SummaryResponseCodesParamsFormatJson AS112SummaryResponseCodesParamsFormat = "JSON"
	AS112SummaryResponseCodesParamsFormatCsv  AS112SummaryResponseCodesParamsFormat = "CSV"
)

func (r AS112SummaryResponseCodesParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryResponseCodesParamsFormatJson, AS112SummaryResponseCodesParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryResponseCodesResponseEnvelope struct {
	Result  AS112SummaryResponseCodesResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    as112SummaryResponseCodesResponseEnvelopeJSON `json:"-"`
}

// as112SummaryResponseCodesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryResponseCodesResponseEnvelope]
type as112SummaryResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
