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

// RadarAS112SummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAS112SummaryService] method
// instead.
type RadarAS112SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAS112SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAS112SummaryService(opts ...option.RequestOption) (r *RadarAS112SummaryService) {
	r = &RadarAS112SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS queries to AS112 by DNSSEC support.
func (r *RadarAS112SummaryService) DNSSEC(ctx context.Context, query RadarAS112SummaryDNSSECParams, opts ...option.RequestOption) (res *RadarAS112SummaryDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112SummaryDNSSECResponseEnvelope
	path := "radar/as112/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries, to AS112, by EDNS support.
func (r *RadarAS112SummaryService) Edns(ctx context.Context, query RadarAS112SummaryEdnsParams, opts ...option.RequestOption) (res *RadarAS112SummaryEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112SummaryEdnsResponseEnvelope
	path := "radar/as112/summary/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 per IP Version.
func (r *RadarAS112SummaryService) IPVersion(ctx context.Context, query RadarAS112SummaryIPVersionParams, opts ...option.RequestOption) (res *RadarAS112SummaryIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112SummaryIPVersionResponseEnvelope
	path := "radar/as112/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 per protocol.
func (r *RadarAS112SummaryService) Protocol(ctx context.Context, query RadarAS112SummaryProtocolParams, opts ...option.RequestOption) (res *RadarAS112SummaryProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112SummaryProtocolResponseEnvelope
	path := "radar/as112/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 by Query Type.
func (r *RadarAS112SummaryService) QueryType(ctx context.Context, query RadarAS112SummaryQueryTypeParams, opts ...option.RequestOption) (res *RadarAS112SummaryQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112SummaryQueryTypeResponseEnvelope
	path := "radar/as112/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 dns requests classified per Response Codes.
func (r *RadarAS112SummaryService) ResponseCodes(ctx context.Context, query RadarAS112SummaryResponseCodesParams, opts ...option.RequestOption) (res *RadarAS112SummaryResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112SummaryResponseCodesResponseEnvelope
	path := "radar/as112/summary/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAS112SummaryDNSSECResponse struct {
	Meta     RadarAS112SummaryDNSSECResponseMeta     `json:"meta,required"`
	Summary0 RadarAS112SummaryDNSSECResponseSummary0 `json:"summary_0,required"`
	JSON     radarAS112SummaryDNSSECResponseJSON     `json:"-"`
}

// radarAS112SummaryDNSSECResponseJSON contains the JSON metadata for the struct
// [RadarAS112SummaryDNSSECResponse]
type radarAS112SummaryDNSSECResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryDNSSECResponseMeta struct {
	DateRange      []RadarAS112SummaryDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	Normalization  string                                            `json:"normalization,required"`
	ConfidenceInfo RadarAS112SummaryDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112SummaryDNSSECResponseMetaJSON           `json:"-"`
}

// radarAS112SummaryDNSSECResponseMetaJSON contains the JSON metadata for the
// struct [RadarAS112SummaryDNSSECResponseMeta]
type radarAS112SummaryDNSSECResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112SummaryDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAS112SummaryDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112SummaryDNSSECResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAS112SummaryDNSSECResponseMetaDateRange]
type radarAS112SummaryDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryDNSSECResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        radarAS112SummaryDNSSECResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112SummaryDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAS112SummaryDNSSECResponseMetaConfidenceInfo]
type radarAS112SummaryDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            radarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation]
type radarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryDNSSECResponseSummary0 struct {
	NotSupported string                                      `json:"NOT_SUPPORTED,required"`
	Supported    string                                      `json:"SUPPORTED,required"`
	JSON         radarAS112SummaryDNSSECResponseSummary0JSON `json:"-"`
}

// radarAS112SummaryDNSSECResponseSummary0JSON contains the JSON metadata for the
// struct [RadarAS112SummaryDNSSECResponseSummary0]
type radarAS112SummaryDNSSECResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAS112SummaryDNSSECResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsResponse struct {
	Meta     RadarAS112SummaryEdnsResponseMeta     `json:"meta,required"`
	Summary0 RadarAS112SummaryEdnsResponseSummary0 `json:"summary_0,required"`
	JSON     radarAS112SummaryEdnsResponseJSON     `json:"-"`
}

// radarAS112SummaryEdnsResponseJSON contains the JSON metadata for the struct
// [RadarAS112SummaryEdnsResponse]
type radarAS112SummaryEdnsResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsResponseMeta struct {
	DateRange      []RadarAS112SummaryEdnsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                          `json:"lastUpdated,required"`
	Normalization  string                                          `json:"normalization,required"`
	ConfidenceInfo RadarAS112SummaryEdnsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112SummaryEdnsResponseMetaJSON           `json:"-"`
}

// radarAS112SummaryEdnsResponseMetaJSON contains the JSON metadata for the struct
// [RadarAS112SummaryEdnsResponseMeta]
type radarAS112SummaryEdnsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112SummaryEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      radarAS112SummaryEdnsResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112SummaryEdnsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAS112SummaryEdnsResponseMetaDateRange]
type radarAS112SummaryEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        radarAS112SummaryEdnsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112SummaryEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAS112SummaryEdnsResponseMetaConfidenceInfo]
type radarAS112SummaryEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                        `json:"dataSource,required"`
	Description     string                                                        `json:"description,required"`
	EventType       string                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                     `json:"startTime" format:"date-time"`
	JSON            radarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotation]
type radarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsResponseSummary0 struct {
	NotSupported string                                    `json:"NOT_SUPPORTED,required"`
	Supported    string                                    `json:"SUPPORTED,required"`
	JSON         radarAS112SummaryEdnsResponseSummary0JSON `json:"-"`
}

// radarAS112SummaryEdnsResponseSummary0JSON contains the JSON metadata for the
// struct [RadarAS112SummaryEdnsResponseSummary0]
type radarAS112SummaryEdnsResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAS112SummaryEdnsResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionResponse struct {
	Meta     RadarAS112SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarAS112SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarAS112SummaryIPVersionResponseJSON     `json:"-"`
}

// radarAS112SummaryIPVersionResponseJSON contains the JSON metadata for the struct
// [RadarAS112SummaryIPVersionResponse]
type radarAS112SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionResponseMeta struct {
	DateRange      []RadarAS112SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                               `json:"lastUpdated,required"`
	Normalization  string                                               `json:"normalization,required"`
	ConfidenceInfo RadarAS112SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// radarAS112SummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [RadarAS112SummaryIPVersionResponseMeta]
type radarAS112SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAS112SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAS112SummaryIPVersionResponseMetaDateRange]
type radarAS112SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        radarAS112SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAS112SummaryIPVersionResponseMetaConfidenceInfo]
type radarAS112SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            radarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type radarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                         `json:"IPv4,required"`
	IPv6 string                                         `json:"IPv6,required"`
	JSON radarAS112SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// radarAS112SummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAS112SummaryIPVersionResponseSummary0]
type radarAS112SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolResponse struct {
	Meta     RadarAS112SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 RadarAS112SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     radarAS112SummaryProtocolResponseJSON     `json:"-"`
}

// radarAS112SummaryProtocolResponseJSON contains the JSON metadata for the struct
// [RadarAS112SummaryProtocolResponse]
type radarAS112SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolResponseMeta struct {
	DateRange      []RadarAS112SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	Normalization  string                                              `json:"normalization,required"`
	ConfidenceInfo RadarAS112SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112SummaryProtocolResponseMetaJSON           `json:"-"`
}

// radarAS112SummaryProtocolResponseMetaJSON contains the JSON metadata for the
// struct [RadarAS112SummaryProtocolResponseMeta]
type radarAS112SummaryProtocolResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAS112SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112SummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAS112SummaryProtocolResponseMetaDateRange]
type radarAS112SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarAS112SummaryProtocolResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAS112SummaryProtocolResponseMetaConfidenceInfo]
type radarAS112SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type radarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolResponseSummary0 struct {
	Tcp  string                                        `json:"tcp,required"`
	Udp  string                                        `json:"udp,required"`
	JSON radarAS112SummaryProtocolResponseSummary0JSON `json:"-"`
}

// radarAS112SummaryProtocolResponseSummary0JSON contains the JSON metadata for the
// struct [RadarAS112SummaryProtocolResponseSummary0]
type radarAS112SummaryProtocolResponseSummary0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeResponse struct {
	Meta     RadarAS112SummaryQueryTypeResponseMeta     `json:"meta,required"`
	Summary0 RadarAS112SummaryQueryTypeResponseSummary0 `json:"summary_0,required"`
	JSON     radarAS112SummaryQueryTypeResponseJSON     `json:"-"`
}

// radarAS112SummaryQueryTypeResponseJSON contains the JSON metadata for the struct
// [RadarAS112SummaryQueryTypeResponse]
type radarAS112SummaryQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeResponseMeta struct {
	DateRange      []RadarAS112SummaryQueryTypeResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                               `json:"lastUpdated,required"`
	Normalization  string                                               `json:"normalization,required"`
	ConfidenceInfo RadarAS112SummaryQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112SummaryQueryTypeResponseMetaJSON           `json:"-"`
}

// radarAS112SummaryQueryTypeResponseMetaJSON contains the JSON metadata for the
// struct [RadarAS112SummaryQueryTypeResponseMeta]
type radarAS112SummaryQueryTypeResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112SummaryQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAS112SummaryQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112SummaryQueryTypeResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAS112SummaryQueryTypeResponseMetaDateRange]
type radarAS112SummaryQueryTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        radarAS112SummaryQueryTypeResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112SummaryQueryTypeResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAS112SummaryQueryTypeResponseMetaConfidenceInfo]
type radarAS112SummaryQueryTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            radarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation]
type radarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeResponseSummary0 struct {
	A    string                                         `json:"A,required"`
	AAAA string                                         `json:"AAAA,required"`
	PTR  string                                         `json:"PTR,required"`
	Soa  string                                         `json:"SOA,required"`
	SRV  string                                         `json:"SRV,required"`
	JSON radarAS112SummaryQueryTypeResponseSummary0JSON `json:"-"`
}

// radarAS112SummaryQueryTypeResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAS112SummaryQueryTypeResponseSummary0]
type radarAS112SummaryQueryTypeResponseSummary0JSON struct {
	A           apijson.Field
	AAAA        apijson.Field
	PTR         apijson.Field
	Soa         apijson.Field
	SRV         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryQueryTypeResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesResponse struct {
	Meta     RadarAS112SummaryResponseCodesResponseMeta     `json:"meta,required"`
	Summary0 RadarAS112SummaryResponseCodesResponseSummary0 `json:"summary_0,required"`
	JSON     radarAS112SummaryResponseCodesResponseJSON     `json:"-"`
}

// radarAS112SummaryResponseCodesResponseJSON contains the JSON metadata for the
// struct [RadarAS112SummaryResponseCodesResponse]
type radarAS112SummaryResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesResponseMeta struct {
	DateRange      []RadarAS112SummaryResponseCodesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAS112SummaryResponseCodesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112SummaryResponseCodesResponseMetaJSON           `json:"-"`
}

// radarAS112SummaryResponseCodesResponseMetaJSON contains the JSON metadata for
// the struct [RadarAS112SummaryResponseCodesResponseMeta]
type radarAS112SummaryResponseCodesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112SummaryResponseCodesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAS112SummaryResponseCodesResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112SummaryResponseCodesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAS112SummaryResponseCodesResponseMetaDateRange]
type radarAS112SummaryResponseCodesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryResponseCodesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAS112SummaryResponseCodesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112SummaryResponseCodesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAS112SummaryResponseCodesResponseMetaConfidenceInfo]
type radarAS112SummaryResponseCodesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryResponseCodesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation]
type radarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesResponseSummary0 struct {
	Noerror  string                                             `json:"NOERROR,required"`
	Nxdomain string                                             `json:"NXDOMAIN,required"`
	JSON     radarAS112SummaryResponseCodesResponseSummary0JSON `json:"-"`
}

// radarAS112SummaryResponseCodesResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAS112SummaryResponseCodesResponseSummary0]
type radarAS112SummaryResponseCodesResponseSummary0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryResponseCodesResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryDNSSECParams struct {
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
	DateRange param.Field[[]RadarAS112SummaryDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112SummaryDNSSECParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112SummaryDNSSECParams]'s query parameters as
// `url.Values`.
func (r RadarAS112SummaryDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112SummaryDNSSECParamsDateRange string

const (
	RadarAS112SummaryDNSSECParamsDateRange1d         RadarAS112SummaryDNSSECParamsDateRange = "1d"
	RadarAS112SummaryDNSSECParamsDateRange2d         RadarAS112SummaryDNSSECParamsDateRange = "2d"
	RadarAS112SummaryDNSSECParamsDateRange7d         RadarAS112SummaryDNSSECParamsDateRange = "7d"
	RadarAS112SummaryDNSSECParamsDateRange14d        RadarAS112SummaryDNSSECParamsDateRange = "14d"
	RadarAS112SummaryDNSSECParamsDateRange28d        RadarAS112SummaryDNSSECParamsDateRange = "28d"
	RadarAS112SummaryDNSSECParamsDateRange12w        RadarAS112SummaryDNSSECParamsDateRange = "12w"
	RadarAS112SummaryDNSSECParamsDateRange24w        RadarAS112SummaryDNSSECParamsDateRange = "24w"
	RadarAS112SummaryDNSSECParamsDateRange52w        RadarAS112SummaryDNSSECParamsDateRange = "52w"
	RadarAS112SummaryDNSSECParamsDateRange1dControl  RadarAS112SummaryDNSSECParamsDateRange = "1dControl"
	RadarAS112SummaryDNSSECParamsDateRange2dControl  RadarAS112SummaryDNSSECParamsDateRange = "2dControl"
	RadarAS112SummaryDNSSECParamsDateRange7dControl  RadarAS112SummaryDNSSECParamsDateRange = "7dControl"
	RadarAS112SummaryDNSSECParamsDateRange14dControl RadarAS112SummaryDNSSECParamsDateRange = "14dControl"
	RadarAS112SummaryDNSSECParamsDateRange28dControl RadarAS112SummaryDNSSECParamsDateRange = "28dControl"
	RadarAS112SummaryDNSSECParamsDateRange12wControl RadarAS112SummaryDNSSECParamsDateRange = "12wControl"
	RadarAS112SummaryDNSSECParamsDateRange24wControl RadarAS112SummaryDNSSECParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112SummaryDNSSECParamsFormat string

const (
	RadarAS112SummaryDNSSECParamsFormatJson RadarAS112SummaryDNSSECParamsFormat = "JSON"
	RadarAS112SummaryDNSSECParamsFormatCsv  RadarAS112SummaryDNSSECParamsFormat = "CSV"
)

type RadarAS112SummaryDNSSECResponseEnvelope struct {
	Result  RadarAS112SummaryDNSSECResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAS112SummaryDNSSECResponseEnvelopeJSON `json:"-"`
}

// radarAS112SummaryDNSSECResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112SummaryDNSSECResponseEnvelope]
type radarAS112SummaryDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryEdnsParams struct {
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
	DateRange param.Field[[]RadarAS112SummaryEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112SummaryEdnsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112SummaryEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAS112SummaryEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112SummaryEdnsParamsDateRange string

const (
	RadarAS112SummaryEdnsParamsDateRange1d         RadarAS112SummaryEdnsParamsDateRange = "1d"
	RadarAS112SummaryEdnsParamsDateRange2d         RadarAS112SummaryEdnsParamsDateRange = "2d"
	RadarAS112SummaryEdnsParamsDateRange7d         RadarAS112SummaryEdnsParamsDateRange = "7d"
	RadarAS112SummaryEdnsParamsDateRange14d        RadarAS112SummaryEdnsParamsDateRange = "14d"
	RadarAS112SummaryEdnsParamsDateRange28d        RadarAS112SummaryEdnsParamsDateRange = "28d"
	RadarAS112SummaryEdnsParamsDateRange12w        RadarAS112SummaryEdnsParamsDateRange = "12w"
	RadarAS112SummaryEdnsParamsDateRange24w        RadarAS112SummaryEdnsParamsDateRange = "24w"
	RadarAS112SummaryEdnsParamsDateRange52w        RadarAS112SummaryEdnsParamsDateRange = "52w"
	RadarAS112SummaryEdnsParamsDateRange1dControl  RadarAS112SummaryEdnsParamsDateRange = "1dControl"
	RadarAS112SummaryEdnsParamsDateRange2dControl  RadarAS112SummaryEdnsParamsDateRange = "2dControl"
	RadarAS112SummaryEdnsParamsDateRange7dControl  RadarAS112SummaryEdnsParamsDateRange = "7dControl"
	RadarAS112SummaryEdnsParamsDateRange14dControl RadarAS112SummaryEdnsParamsDateRange = "14dControl"
	RadarAS112SummaryEdnsParamsDateRange28dControl RadarAS112SummaryEdnsParamsDateRange = "28dControl"
	RadarAS112SummaryEdnsParamsDateRange12wControl RadarAS112SummaryEdnsParamsDateRange = "12wControl"
	RadarAS112SummaryEdnsParamsDateRange24wControl RadarAS112SummaryEdnsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112SummaryEdnsParamsFormat string

const (
	RadarAS112SummaryEdnsParamsFormatJson RadarAS112SummaryEdnsParamsFormat = "JSON"
	RadarAS112SummaryEdnsParamsFormatCsv  RadarAS112SummaryEdnsParamsFormat = "CSV"
)

type RadarAS112SummaryEdnsResponseEnvelope struct {
	Result  RadarAS112SummaryEdnsResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarAS112SummaryEdnsResponseEnvelopeJSON `json:"-"`
}

// radarAS112SummaryEdnsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112SummaryEdnsResponseEnvelope]
type radarAS112SummaryEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryIPVersionParams struct {
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
	DateRange param.Field[[]RadarAS112SummaryIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112SummaryIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarAS112SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112SummaryIPVersionParamsDateRange string

const (
	RadarAS112SummaryIPVersionParamsDateRange1d         RadarAS112SummaryIPVersionParamsDateRange = "1d"
	RadarAS112SummaryIPVersionParamsDateRange2d         RadarAS112SummaryIPVersionParamsDateRange = "2d"
	RadarAS112SummaryIPVersionParamsDateRange7d         RadarAS112SummaryIPVersionParamsDateRange = "7d"
	RadarAS112SummaryIPVersionParamsDateRange14d        RadarAS112SummaryIPVersionParamsDateRange = "14d"
	RadarAS112SummaryIPVersionParamsDateRange28d        RadarAS112SummaryIPVersionParamsDateRange = "28d"
	RadarAS112SummaryIPVersionParamsDateRange12w        RadarAS112SummaryIPVersionParamsDateRange = "12w"
	RadarAS112SummaryIPVersionParamsDateRange24w        RadarAS112SummaryIPVersionParamsDateRange = "24w"
	RadarAS112SummaryIPVersionParamsDateRange52w        RadarAS112SummaryIPVersionParamsDateRange = "52w"
	RadarAS112SummaryIPVersionParamsDateRange1dControl  RadarAS112SummaryIPVersionParamsDateRange = "1dControl"
	RadarAS112SummaryIPVersionParamsDateRange2dControl  RadarAS112SummaryIPVersionParamsDateRange = "2dControl"
	RadarAS112SummaryIPVersionParamsDateRange7dControl  RadarAS112SummaryIPVersionParamsDateRange = "7dControl"
	RadarAS112SummaryIPVersionParamsDateRange14dControl RadarAS112SummaryIPVersionParamsDateRange = "14dControl"
	RadarAS112SummaryIPVersionParamsDateRange28dControl RadarAS112SummaryIPVersionParamsDateRange = "28dControl"
	RadarAS112SummaryIPVersionParamsDateRange12wControl RadarAS112SummaryIPVersionParamsDateRange = "12wControl"
	RadarAS112SummaryIPVersionParamsDateRange24wControl RadarAS112SummaryIPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112SummaryIPVersionParamsFormat string

const (
	RadarAS112SummaryIPVersionParamsFormatJson RadarAS112SummaryIPVersionParamsFormat = "JSON"
	RadarAS112SummaryIPVersionParamsFormatCsv  RadarAS112SummaryIPVersionParamsFormat = "CSV"
)

type RadarAS112SummaryIPVersionResponseEnvelope struct {
	Result  RadarAS112SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAS112SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAS112SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAS112SummaryIPVersionResponseEnvelope]
type radarAS112SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryProtocolParams struct {
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
	DateRange param.Field[[]RadarAS112SummaryProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112SummaryProtocolParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112SummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r RadarAS112SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112SummaryProtocolParamsDateRange string

const (
	RadarAS112SummaryProtocolParamsDateRange1d         RadarAS112SummaryProtocolParamsDateRange = "1d"
	RadarAS112SummaryProtocolParamsDateRange2d         RadarAS112SummaryProtocolParamsDateRange = "2d"
	RadarAS112SummaryProtocolParamsDateRange7d         RadarAS112SummaryProtocolParamsDateRange = "7d"
	RadarAS112SummaryProtocolParamsDateRange14d        RadarAS112SummaryProtocolParamsDateRange = "14d"
	RadarAS112SummaryProtocolParamsDateRange28d        RadarAS112SummaryProtocolParamsDateRange = "28d"
	RadarAS112SummaryProtocolParamsDateRange12w        RadarAS112SummaryProtocolParamsDateRange = "12w"
	RadarAS112SummaryProtocolParamsDateRange24w        RadarAS112SummaryProtocolParamsDateRange = "24w"
	RadarAS112SummaryProtocolParamsDateRange52w        RadarAS112SummaryProtocolParamsDateRange = "52w"
	RadarAS112SummaryProtocolParamsDateRange1dControl  RadarAS112SummaryProtocolParamsDateRange = "1dControl"
	RadarAS112SummaryProtocolParamsDateRange2dControl  RadarAS112SummaryProtocolParamsDateRange = "2dControl"
	RadarAS112SummaryProtocolParamsDateRange7dControl  RadarAS112SummaryProtocolParamsDateRange = "7dControl"
	RadarAS112SummaryProtocolParamsDateRange14dControl RadarAS112SummaryProtocolParamsDateRange = "14dControl"
	RadarAS112SummaryProtocolParamsDateRange28dControl RadarAS112SummaryProtocolParamsDateRange = "28dControl"
	RadarAS112SummaryProtocolParamsDateRange12wControl RadarAS112SummaryProtocolParamsDateRange = "12wControl"
	RadarAS112SummaryProtocolParamsDateRange24wControl RadarAS112SummaryProtocolParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112SummaryProtocolParamsFormat string

const (
	RadarAS112SummaryProtocolParamsFormatJson RadarAS112SummaryProtocolParamsFormat = "JSON"
	RadarAS112SummaryProtocolParamsFormatCsv  RadarAS112SummaryProtocolParamsFormat = "CSV"
)

type RadarAS112SummaryProtocolResponseEnvelope struct {
	Result  RadarAS112SummaryProtocolResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAS112SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAS112SummaryProtocolResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112SummaryProtocolResponseEnvelope]
type radarAS112SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryQueryTypeParams struct {
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
	DateRange param.Field[[]RadarAS112SummaryQueryTypeParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112SummaryQueryTypeParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112SummaryQueryTypeParams]'s query parameters as
// `url.Values`.
func (r RadarAS112SummaryQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112SummaryQueryTypeParamsDateRange string

const (
	RadarAS112SummaryQueryTypeParamsDateRange1d         RadarAS112SummaryQueryTypeParamsDateRange = "1d"
	RadarAS112SummaryQueryTypeParamsDateRange2d         RadarAS112SummaryQueryTypeParamsDateRange = "2d"
	RadarAS112SummaryQueryTypeParamsDateRange7d         RadarAS112SummaryQueryTypeParamsDateRange = "7d"
	RadarAS112SummaryQueryTypeParamsDateRange14d        RadarAS112SummaryQueryTypeParamsDateRange = "14d"
	RadarAS112SummaryQueryTypeParamsDateRange28d        RadarAS112SummaryQueryTypeParamsDateRange = "28d"
	RadarAS112SummaryQueryTypeParamsDateRange12w        RadarAS112SummaryQueryTypeParamsDateRange = "12w"
	RadarAS112SummaryQueryTypeParamsDateRange24w        RadarAS112SummaryQueryTypeParamsDateRange = "24w"
	RadarAS112SummaryQueryTypeParamsDateRange52w        RadarAS112SummaryQueryTypeParamsDateRange = "52w"
	RadarAS112SummaryQueryTypeParamsDateRange1dControl  RadarAS112SummaryQueryTypeParamsDateRange = "1dControl"
	RadarAS112SummaryQueryTypeParamsDateRange2dControl  RadarAS112SummaryQueryTypeParamsDateRange = "2dControl"
	RadarAS112SummaryQueryTypeParamsDateRange7dControl  RadarAS112SummaryQueryTypeParamsDateRange = "7dControl"
	RadarAS112SummaryQueryTypeParamsDateRange14dControl RadarAS112SummaryQueryTypeParamsDateRange = "14dControl"
	RadarAS112SummaryQueryTypeParamsDateRange28dControl RadarAS112SummaryQueryTypeParamsDateRange = "28dControl"
	RadarAS112SummaryQueryTypeParamsDateRange12wControl RadarAS112SummaryQueryTypeParamsDateRange = "12wControl"
	RadarAS112SummaryQueryTypeParamsDateRange24wControl RadarAS112SummaryQueryTypeParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112SummaryQueryTypeParamsFormat string

const (
	RadarAS112SummaryQueryTypeParamsFormatJson RadarAS112SummaryQueryTypeParamsFormat = "JSON"
	RadarAS112SummaryQueryTypeParamsFormatCsv  RadarAS112SummaryQueryTypeParamsFormat = "CSV"
)

type RadarAS112SummaryQueryTypeResponseEnvelope struct {
	Result  RadarAS112SummaryQueryTypeResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAS112SummaryQueryTypeResponseEnvelopeJSON `json:"-"`
}

// radarAS112SummaryQueryTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAS112SummaryQueryTypeResponseEnvelope]
type radarAS112SummaryQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112SummaryResponseCodesParams struct {
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
	DateRange param.Field[[]RadarAS112SummaryResponseCodesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112SummaryResponseCodesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112SummaryResponseCodesParams]'s query parameters as
// `url.Values`.
func (r RadarAS112SummaryResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAS112SummaryResponseCodesParamsDateRange string

const (
	RadarAS112SummaryResponseCodesParamsDateRange1d         RadarAS112SummaryResponseCodesParamsDateRange = "1d"
	RadarAS112SummaryResponseCodesParamsDateRange2d         RadarAS112SummaryResponseCodesParamsDateRange = "2d"
	RadarAS112SummaryResponseCodesParamsDateRange7d         RadarAS112SummaryResponseCodesParamsDateRange = "7d"
	RadarAS112SummaryResponseCodesParamsDateRange14d        RadarAS112SummaryResponseCodesParamsDateRange = "14d"
	RadarAS112SummaryResponseCodesParamsDateRange28d        RadarAS112SummaryResponseCodesParamsDateRange = "28d"
	RadarAS112SummaryResponseCodesParamsDateRange12w        RadarAS112SummaryResponseCodesParamsDateRange = "12w"
	RadarAS112SummaryResponseCodesParamsDateRange24w        RadarAS112SummaryResponseCodesParamsDateRange = "24w"
	RadarAS112SummaryResponseCodesParamsDateRange52w        RadarAS112SummaryResponseCodesParamsDateRange = "52w"
	RadarAS112SummaryResponseCodesParamsDateRange1dControl  RadarAS112SummaryResponseCodesParamsDateRange = "1dControl"
	RadarAS112SummaryResponseCodesParamsDateRange2dControl  RadarAS112SummaryResponseCodesParamsDateRange = "2dControl"
	RadarAS112SummaryResponseCodesParamsDateRange7dControl  RadarAS112SummaryResponseCodesParamsDateRange = "7dControl"
	RadarAS112SummaryResponseCodesParamsDateRange14dControl RadarAS112SummaryResponseCodesParamsDateRange = "14dControl"
	RadarAS112SummaryResponseCodesParamsDateRange28dControl RadarAS112SummaryResponseCodesParamsDateRange = "28dControl"
	RadarAS112SummaryResponseCodesParamsDateRange12wControl RadarAS112SummaryResponseCodesParamsDateRange = "12wControl"
	RadarAS112SummaryResponseCodesParamsDateRange24wControl RadarAS112SummaryResponseCodesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112SummaryResponseCodesParamsFormat string

const (
	RadarAS112SummaryResponseCodesParamsFormatJson RadarAS112SummaryResponseCodesParamsFormat = "JSON"
	RadarAS112SummaryResponseCodesParamsFormatCsv  RadarAS112SummaryResponseCodesParamsFormat = "CSV"
)

type RadarAS112SummaryResponseCodesResponseEnvelope struct {
	Result  RadarAS112SummaryResponseCodesResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAS112SummaryResponseCodesResponseEnvelopeJSON `json:"-"`
}

// radarAS112SummaryResponseCodesResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAS112SummaryResponseCodesResponseEnvelope]
type radarAS112SummaryResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112SummaryResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112SummaryResponseCodesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
