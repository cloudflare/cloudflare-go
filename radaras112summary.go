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

// RadarAs112SummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112SummaryService] method
// instead.
type RadarAs112SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryService(opts ...option.RequestOption) (r *RadarAs112SummaryService) {
	r = &RadarAs112SummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS queries to AS112 by DNSSEC support.
func (r *RadarAs112SummaryService) DNSSEC(ctx context.Context, query RadarAs112SummaryDNSSECParams, opts ...option.RequestOption) (res *RadarAs112SummaryDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112SummaryDNSSECResponseEnvelope
	path := "radar/as112/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries, to AS112, by EDNS support.
func (r *RadarAs112SummaryService) Edns(ctx context.Context, query RadarAs112SummaryEdnsParams, opts ...option.RequestOption) (res *RadarAs112SummaryEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112SummaryEdnsResponseEnvelope
	path := "radar/as112/summary/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 per IP Version.
func (r *RadarAs112SummaryService) IPVersion(ctx context.Context, query RadarAs112SummaryIPVersionParams, opts ...option.RequestOption) (res *RadarAs112SummaryIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112SummaryIPVersionResponseEnvelope
	path := "radar/as112/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 per protocol.
func (r *RadarAs112SummaryService) Protocol(ctx context.Context, query RadarAs112SummaryProtocolParams, opts ...option.RequestOption) (res *RadarAs112SummaryProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112SummaryProtocolResponseEnvelope
	path := "radar/as112/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of DNS queries to AS112 by Query Type.
func (r *RadarAs112SummaryService) QueryType(ctx context.Context, query RadarAs112SummaryQueryTypeParams, opts ...option.RequestOption) (res *RadarAs112SummaryQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112SummaryQueryTypeResponseEnvelope
	path := "radar/as112/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 dns requests classified per Response Codes.
func (r *RadarAs112SummaryService) ResponseCodes(ctx context.Context, query RadarAs112SummaryResponseCodesParams, opts ...option.RequestOption) (res *RadarAs112SummaryResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112SummaryResponseCodesResponseEnvelope
	path := "radar/as112/summary/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112SummaryDNSSECResponse struct {
	Meta     RadarAs112SummaryDNSSECResponseMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryDNSSECResponseSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryDNSSECResponseJSON     `json:"-"`
}

// radarAs112SummaryDNSSECResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryDNSSECResponse]
type radarAs112SummaryDNSSECResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDNSSECResponseMeta struct {
	DateRange      []RadarAs112SummaryDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	Normalization  string                                            `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryDNSSECResponseMetaJSON           `json:"-"`
}

// radarAs112SummaryDNSSECResponseMetaJSON contains the JSON metadata for the
// struct [RadarAs112SummaryDNSSECResponseMeta]
type radarAs112SummaryDNSSECResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryDNSSECResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAs112SummaryDNSSECResponseMetaDateRange]
type radarAs112SummaryDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDNSSECResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        radarAs112SummaryDNSSECResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAs112SummaryDNSSECResponseMetaConfidenceInfo]
type radarAs112SummaryDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotation]
type radarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDNSSECResponseSummary0 struct {
	NotSupported string                                      `json:"NOT_SUPPORTED,required"`
	Supported    string                                      `json:"SUPPORTED,required"`
	JSON         radarAs112SummaryDNSSECResponseSummary0JSON `json:"-"`
}

// radarAs112SummaryDNSSECResponseSummary0JSON contains the JSON metadata for the
// struct [RadarAs112SummaryDNSSECResponseSummary0]
type radarAs112SummaryDNSSECResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112SummaryDNSSECResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsResponse struct {
	Meta     RadarAs112SummaryEdnsResponseMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryEdnsResponseSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryEdnsResponseJSON     `json:"-"`
}

// radarAs112SummaryEdnsResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryEdnsResponse]
type radarAs112SummaryEdnsResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsResponseMeta struct {
	DateRange      []RadarAs112SummaryEdnsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                          `json:"lastUpdated,required"`
	Normalization  string                                          `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryEdnsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryEdnsResponseMetaJSON           `json:"-"`
}

// radarAs112SummaryEdnsResponseMetaJSON contains the JSON metadata for the struct
// [RadarAs112SummaryEdnsResponseMeta]
type radarAs112SummaryEdnsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryEdnsResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryEdnsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarAs112SummaryEdnsResponseMetaDateRange]
type radarAs112SummaryEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        radarAs112SummaryEdnsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAs112SummaryEdnsResponseMetaConfidenceInfo]
type radarAs112SummaryEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                        `json:"dataSource,required"`
	Description     string                                                        `json:"description,required"`
	EventType       string                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                     `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotation]
type radarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsResponseSummary0 struct {
	NotSupported string                                    `json:"NOT_SUPPORTED,required"`
	Supported    string                                    `json:"SUPPORTED,required"`
	JSON         radarAs112SummaryEdnsResponseSummary0JSON `json:"-"`
}

// radarAs112SummaryEdnsResponseSummary0JSON contains the JSON metadata for the
// struct [RadarAs112SummaryEdnsResponseSummary0]
type radarAs112SummaryEdnsResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnsResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionResponse struct {
	Meta     RadarAs112SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryIPVersionResponseJSON     `json:"-"`
}

// radarAs112SummaryIPVersionResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryIPVersionResponse]
type radarAs112SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionResponseMeta struct {
	DateRange      []RadarAs112SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                               `json:"lastUpdated,required"`
	Normalization  string                                               `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryIPVersionResponseMetaJSON           `json:"-"`
}

// radarAs112SummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [RadarAs112SummaryIPVersionResponseMeta]
type radarAs112SummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAs112SummaryIPVersionResponseMetaDateRange]
type radarAs112SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        radarAs112SummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAs112SummaryIPVersionResponseMetaConfidenceInfo]
type radarAs112SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type radarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                         `json:"IPv4,required"`
	IPv6 string                                         `json:"IPv6,required"`
	JSON radarAs112SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// radarAs112SummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAs112SummaryIPVersionResponseSummary0]
type radarAs112SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolResponse struct {
	Meta     RadarAs112SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryProtocolResponseJSON     `json:"-"`
}

// radarAs112SummaryProtocolResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryProtocolResponse]
type radarAs112SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolResponseMeta struct {
	DateRange      []RadarAs112SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	Normalization  string                                              `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryProtocolResponseMetaJSON           `json:"-"`
}

// radarAs112SummaryProtocolResponseMetaJSON contains the JSON metadata for the
// struct [RadarAs112SummaryProtocolResponseMeta]
type radarAs112SummaryProtocolResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAs112SummaryProtocolResponseMetaDateRange]
type radarAs112SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarAs112SummaryProtocolResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAs112SummaryProtocolResponseMetaConfidenceInfo]
type radarAs112SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type radarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolResponseSummary0 struct {
	Tcp  string                                        `json:"tcp,required"`
	Udp  string                                        `json:"udp,required"`
	JSON radarAs112SummaryProtocolResponseSummary0JSON `json:"-"`
}

// radarAs112SummaryProtocolResponseSummary0JSON contains the JSON metadata for the
// struct [RadarAs112SummaryProtocolResponseSummary0]
type radarAs112SummaryProtocolResponseSummary0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeResponse struct {
	Meta     RadarAs112SummaryQueryTypeResponseMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryQueryTypeResponseSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryQueryTypeResponseJSON     `json:"-"`
}

// radarAs112SummaryQueryTypeResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryQueryTypeResponse]
type radarAs112SummaryQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeResponseMeta struct {
	DateRange      []RadarAs112SummaryQueryTypeResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                               `json:"lastUpdated,required"`
	Normalization  string                                               `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryQueryTypeResponseMetaJSON           `json:"-"`
}

// radarAs112SummaryQueryTypeResponseMetaJSON contains the JSON metadata for the
// struct [RadarAs112SummaryQueryTypeResponseMeta]
type radarAs112SummaryQueryTypeResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryQueryTypeResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAs112SummaryQueryTypeResponseMetaDateRange]
type radarAs112SummaryQueryTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        radarAs112SummaryQueryTypeResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryQueryTypeResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAs112SummaryQueryTypeResponseMetaConfidenceInfo]
type radarAs112SummaryQueryTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation]
type radarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeResponseSummary0 struct {
	A    string                                         `json:"A,required"`
	Aaaa string                                         `json:"AAAA,required"`
	Ptr  string                                         `json:"PTR,required"`
	Soa  string                                         `json:"SOA,required"`
	Srv  string                                         `json:"SRV,required"`
	JSON radarAs112SummaryQueryTypeResponseSummary0JSON `json:"-"`
}

// radarAs112SummaryQueryTypeResponseSummary0JSON contains the JSON metadata for
// the struct [RadarAs112SummaryQueryTypeResponseSummary0]
type radarAs112SummaryQueryTypeResponseSummary0JSON struct {
	A           apijson.Field
	Aaaa        apijson.Field
	Ptr         apijson.Field
	Soa         apijson.Field
	Srv         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesResponse struct {
	Meta     RadarAs112SummaryResponseCodesResponseMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryResponseCodesResponseSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryResponseCodesResponseJSON     `json:"-"`
}

// radarAs112SummaryResponseCodesResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryResponseCodesResponse]
type radarAs112SummaryResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesResponseMeta struct {
	DateRange      []RadarAs112SummaryResponseCodesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryResponseCodesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryResponseCodesResponseMetaJSON           `json:"-"`
}

// radarAs112SummaryResponseCodesResponseMetaJSON contains the JSON metadata for
// the struct [RadarAs112SummaryResponseCodesResponseMeta]
type radarAs112SummaryResponseCodesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryResponseCodesResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryResponseCodesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112SummaryResponseCodesResponseMetaDateRange]
type radarAs112SummaryResponseCodesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAs112SummaryResponseCodesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryResponseCodesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryResponseCodesResponseMetaConfidenceInfo]
type radarAs112SummaryResponseCodesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation]
type radarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesResponseSummary0 struct {
	Noerror  string                                             `json:"NOERROR,required"`
	Nxdomain string                                             `json:"NXDOMAIN,required"`
	JSON     radarAs112SummaryResponseCodesResponseSummary0JSON `json:"-"`
}

// radarAs112SummaryResponseCodesResponseSummary0JSON contains the JSON metadata
// for the struct [RadarAs112SummaryResponseCodesResponseSummary0]
type radarAs112SummaryResponseCodesResponseSummary0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodesResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryDNSSECParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryDNSSECParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryDNSSECParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryDNSSECParamsDateRange string

const (
	RadarAs112SummaryDNSSECParamsDateRange1d         RadarAs112SummaryDNSSECParamsDateRange = "1d"
	RadarAs112SummaryDNSSECParamsDateRange2d         RadarAs112SummaryDNSSECParamsDateRange = "2d"
	RadarAs112SummaryDNSSECParamsDateRange7d         RadarAs112SummaryDNSSECParamsDateRange = "7d"
	RadarAs112SummaryDNSSECParamsDateRange14d        RadarAs112SummaryDNSSECParamsDateRange = "14d"
	RadarAs112SummaryDNSSECParamsDateRange28d        RadarAs112SummaryDNSSECParamsDateRange = "28d"
	RadarAs112SummaryDNSSECParamsDateRange12w        RadarAs112SummaryDNSSECParamsDateRange = "12w"
	RadarAs112SummaryDNSSECParamsDateRange24w        RadarAs112SummaryDNSSECParamsDateRange = "24w"
	RadarAs112SummaryDNSSECParamsDateRange52w        RadarAs112SummaryDNSSECParamsDateRange = "52w"
	RadarAs112SummaryDNSSECParamsDateRange1dControl  RadarAs112SummaryDNSSECParamsDateRange = "1dControl"
	RadarAs112SummaryDNSSECParamsDateRange2dControl  RadarAs112SummaryDNSSECParamsDateRange = "2dControl"
	RadarAs112SummaryDNSSECParamsDateRange7dControl  RadarAs112SummaryDNSSECParamsDateRange = "7dControl"
	RadarAs112SummaryDNSSECParamsDateRange14dControl RadarAs112SummaryDNSSECParamsDateRange = "14dControl"
	RadarAs112SummaryDNSSECParamsDateRange28dControl RadarAs112SummaryDNSSECParamsDateRange = "28dControl"
	RadarAs112SummaryDNSSECParamsDateRange12wControl RadarAs112SummaryDNSSECParamsDateRange = "12wControl"
	RadarAs112SummaryDNSSECParamsDateRange24wControl RadarAs112SummaryDNSSECParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryDNSSECParamsFormat string

const (
	RadarAs112SummaryDNSSECParamsFormatJson RadarAs112SummaryDNSSECParamsFormat = "JSON"
	RadarAs112SummaryDNSSECParamsFormatCsv  RadarAs112SummaryDNSSECParamsFormat = "CSV"
)

type RadarAs112SummaryDNSSECResponseEnvelope struct {
	Result  RadarAs112SummaryDNSSECResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAs112SummaryDNSSECResponseEnvelopeJSON `json:"-"`
}

// radarAs112SummaryDNSSECResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112SummaryDNSSECResponseEnvelope]
type radarAs112SummaryDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryEdnsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryEdnsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryEdnsParamsDateRange string

const (
	RadarAs112SummaryEdnsParamsDateRange1d         RadarAs112SummaryEdnsParamsDateRange = "1d"
	RadarAs112SummaryEdnsParamsDateRange2d         RadarAs112SummaryEdnsParamsDateRange = "2d"
	RadarAs112SummaryEdnsParamsDateRange7d         RadarAs112SummaryEdnsParamsDateRange = "7d"
	RadarAs112SummaryEdnsParamsDateRange14d        RadarAs112SummaryEdnsParamsDateRange = "14d"
	RadarAs112SummaryEdnsParamsDateRange28d        RadarAs112SummaryEdnsParamsDateRange = "28d"
	RadarAs112SummaryEdnsParamsDateRange12w        RadarAs112SummaryEdnsParamsDateRange = "12w"
	RadarAs112SummaryEdnsParamsDateRange24w        RadarAs112SummaryEdnsParamsDateRange = "24w"
	RadarAs112SummaryEdnsParamsDateRange52w        RadarAs112SummaryEdnsParamsDateRange = "52w"
	RadarAs112SummaryEdnsParamsDateRange1dControl  RadarAs112SummaryEdnsParamsDateRange = "1dControl"
	RadarAs112SummaryEdnsParamsDateRange2dControl  RadarAs112SummaryEdnsParamsDateRange = "2dControl"
	RadarAs112SummaryEdnsParamsDateRange7dControl  RadarAs112SummaryEdnsParamsDateRange = "7dControl"
	RadarAs112SummaryEdnsParamsDateRange14dControl RadarAs112SummaryEdnsParamsDateRange = "14dControl"
	RadarAs112SummaryEdnsParamsDateRange28dControl RadarAs112SummaryEdnsParamsDateRange = "28dControl"
	RadarAs112SummaryEdnsParamsDateRange12wControl RadarAs112SummaryEdnsParamsDateRange = "12wControl"
	RadarAs112SummaryEdnsParamsDateRange24wControl RadarAs112SummaryEdnsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryEdnsParamsFormat string

const (
	RadarAs112SummaryEdnsParamsFormatJson RadarAs112SummaryEdnsParamsFormat = "JSON"
	RadarAs112SummaryEdnsParamsFormatCsv  RadarAs112SummaryEdnsParamsFormat = "CSV"
)

type RadarAs112SummaryEdnsResponseEnvelope struct {
	Result  RadarAs112SummaryEdnsResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarAs112SummaryEdnsResponseEnvelopeJSON `json:"-"`
}

// radarAs112SummaryEdnsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112SummaryEdnsResponseEnvelope]
type radarAs112SummaryEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryIPVersionParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryIPVersionParamsDateRange string

const (
	RadarAs112SummaryIPVersionParamsDateRange1d         RadarAs112SummaryIPVersionParamsDateRange = "1d"
	RadarAs112SummaryIPVersionParamsDateRange2d         RadarAs112SummaryIPVersionParamsDateRange = "2d"
	RadarAs112SummaryIPVersionParamsDateRange7d         RadarAs112SummaryIPVersionParamsDateRange = "7d"
	RadarAs112SummaryIPVersionParamsDateRange14d        RadarAs112SummaryIPVersionParamsDateRange = "14d"
	RadarAs112SummaryIPVersionParamsDateRange28d        RadarAs112SummaryIPVersionParamsDateRange = "28d"
	RadarAs112SummaryIPVersionParamsDateRange12w        RadarAs112SummaryIPVersionParamsDateRange = "12w"
	RadarAs112SummaryIPVersionParamsDateRange24w        RadarAs112SummaryIPVersionParamsDateRange = "24w"
	RadarAs112SummaryIPVersionParamsDateRange52w        RadarAs112SummaryIPVersionParamsDateRange = "52w"
	RadarAs112SummaryIPVersionParamsDateRange1dControl  RadarAs112SummaryIPVersionParamsDateRange = "1dControl"
	RadarAs112SummaryIPVersionParamsDateRange2dControl  RadarAs112SummaryIPVersionParamsDateRange = "2dControl"
	RadarAs112SummaryIPVersionParamsDateRange7dControl  RadarAs112SummaryIPVersionParamsDateRange = "7dControl"
	RadarAs112SummaryIPVersionParamsDateRange14dControl RadarAs112SummaryIPVersionParamsDateRange = "14dControl"
	RadarAs112SummaryIPVersionParamsDateRange28dControl RadarAs112SummaryIPVersionParamsDateRange = "28dControl"
	RadarAs112SummaryIPVersionParamsDateRange12wControl RadarAs112SummaryIPVersionParamsDateRange = "12wControl"
	RadarAs112SummaryIPVersionParamsDateRange24wControl RadarAs112SummaryIPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryIPVersionParamsFormat string

const (
	RadarAs112SummaryIPVersionParamsFormatJson RadarAs112SummaryIPVersionParamsFormat = "JSON"
	RadarAs112SummaryIPVersionParamsFormatCsv  RadarAs112SummaryIPVersionParamsFormat = "CSV"
)

type RadarAs112SummaryIPVersionResponseEnvelope struct {
	Result  RadarAs112SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAs112SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAs112SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAs112SummaryIPVersionResponseEnvelope]
type radarAs112SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryProtocolParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryProtocolParamsDateRange string

const (
	RadarAs112SummaryProtocolParamsDateRange1d         RadarAs112SummaryProtocolParamsDateRange = "1d"
	RadarAs112SummaryProtocolParamsDateRange2d         RadarAs112SummaryProtocolParamsDateRange = "2d"
	RadarAs112SummaryProtocolParamsDateRange7d         RadarAs112SummaryProtocolParamsDateRange = "7d"
	RadarAs112SummaryProtocolParamsDateRange14d        RadarAs112SummaryProtocolParamsDateRange = "14d"
	RadarAs112SummaryProtocolParamsDateRange28d        RadarAs112SummaryProtocolParamsDateRange = "28d"
	RadarAs112SummaryProtocolParamsDateRange12w        RadarAs112SummaryProtocolParamsDateRange = "12w"
	RadarAs112SummaryProtocolParamsDateRange24w        RadarAs112SummaryProtocolParamsDateRange = "24w"
	RadarAs112SummaryProtocolParamsDateRange52w        RadarAs112SummaryProtocolParamsDateRange = "52w"
	RadarAs112SummaryProtocolParamsDateRange1dControl  RadarAs112SummaryProtocolParamsDateRange = "1dControl"
	RadarAs112SummaryProtocolParamsDateRange2dControl  RadarAs112SummaryProtocolParamsDateRange = "2dControl"
	RadarAs112SummaryProtocolParamsDateRange7dControl  RadarAs112SummaryProtocolParamsDateRange = "7dControl"
	RadarAs112SummaryProtocolParamsDateRange14dControl RadarAs112SummaryProtocolParamsDateRange = "14dControl"
	RadarAs112SummaryProtocolParamsDateRange28dControl RadarAs112SummaryProtocolParamsDateRange = "28dControl"
	RadarAs112SummaryProtocolParamsDateRange12wControl RadarAs112SummaryProtocolParamsDateRange = "12wControl"
	RadarAs112SummaryProtocolParamsDateRange24wControl RadarAs112SummaryProtocolParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryProtocolParamsFormat string

const (
	RadarAs112SummaryProtocolParamsFormatJson RadarAs112SummaryProtocolParamsFormat = "JSON"
	RadarAs112SummaryProtocolParamsFormatCsv  RadarAs112SummaryProtocolParamsFormat = "CSV"
)

type RadarAs112SummaryProtocolResponseEnvelope struct {
	Result  RadarAs112SummaryProtocolResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAs112SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAs112SummaryProtocolResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112SummaryProtocolResponseEnvelope]
type radarAs112SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryQueryTypeParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryQueryTypeParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryQueryTypeParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryQueryTypeParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryQueryTypeParamsDateRange string

const (
	RadarAs112SummaryQueryTypeParamsDateRange1d         RadarAs112SummaryQueryTypeParamsDateRange = "1d"
	RadarAs112SummaryQueryTypeParamsDateRange2d         RadarAs112SummaryQueryTypeParamsDateRange = "2d"
	RadarAs112SummaryQueryTypeParamsDateRange7d         RadarAs112SummaryQueryTypeParamsDateRange = "7d"
	RadarAs112SummaryQueryTypeParamsDateRange14d        RadarAs112SummaryQueryTypeParamsDateRange = "14d"
	RadarAs112SummaryQueryTypeParamsDateRange28d        RadarAs112SummaryQueryTypeParamsDateRange = "28d"
	RadarAs112SummaryQueryTypeParamsDateRange12w        RadarAs112SummaryQueryTypeParamsDateRange = "12w"
	RadarAs112SummaryQueryTypeParamsDateRange24w        RadarAs112SummaryQueryTypeParamsDateRange = "24w"
	RadarAs112SummaryQueryTypeParamsDateRange52w        RadarAs112SummaryQueryTypeParamsDateRange = "52w"
	RadarAs112SummaryQueryTypeParamsDateRange1dControl  RadarAs112SummaryQueryTypeParamsDateRange = "1dControl"
	RadarAs112SummaryQueryTypeParamsDateRange2dControl  RadarAs112SummaryQueryTypeParamsDateRange = "2dControl"
	RadarAs112SummaryQueryTypeParamsDateRange7dControl  RadarAs112SummaryQueryTypeParamsDateRange = "7dControl"
	RadarAs112SummaryQueryTypeParamsDateRange14dControl RadarAs112SummaryQueryTypeParamsDateRange = "14dControl"
	RadarAs112SummaryQueryTypeParamsDateRange28dControl RadarAs112SummaryQueryTypeParamsDateRange = "28dControl"
	RadarAs112SummaryQueryTypeParamsDateRange12wControl RadarAs112SummaryQueryTypeParamsDateRange = "12wControl"
	RadarAs112SummaryQueryTypeParamsDateRange24wControl RadarAs112SummaryQueryTypeParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryQueryTypeParamsFormat string

const (
	RadarAs112SummaryQueryTypeParamsFormatJson RadarAs112SummaryQueryTypeParamsFormat = "JSON"
	RadarAs112SummaryQueryTypeParamsFormatCsv  RadarAs112SummaryQueryTypeParamsFormat = "CSV"
)

type RadarAs112SummaryQueryTypeResponseEnvelope struct {
	Result  RadarAs112SummaryQueryTypeResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAs112SummaryQueryTypeResponseEnvelopeJSON `json:"-"`
}

// radarAs112SummaryQueryTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAs112SummaryQueryTypeResponseEnvelope]
type radarAs112SummaryQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryResponseCodesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryResponseCodesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryResponseCodesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryResponseCodesParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryResponseCodesParamsDateRange string

const (
	RadarAs112SummaryResponseCodesParamsDateRange1d         RadarAs112SummaryResponseCodesParamsDateRange = "1d"
	RadarAs112SummaryResponseCodesParamsDateRange2d         RadarAs112SummaryResponseCodesParamsDateRange = "2d"
	RadarAs112SummaryResponseCodesParamsDateRange7d         RadarAs112SummaryResponseCodesParamsDateRange = "7d"
	RadarAs112SummaryResponseCodesParamsDateRange14d        RadarAs112SummaryResponseCodesParamsDateRange = "14d"
	RadarAs112SummaryResponseCodesParamsDateRange28d        RadarAs112SummaryResponseCodesParamsDateRange = "28d"
	RadarAs112SummaryResponseCodesParamsDateRange12w        RadarAs112SummaryResponseCodesParamsDateRange = "12w"
	RadarAs112SummaryResponseCodesParamsDateRange24w        RadarAs112SummaryResponseCodesParamsDateRange = "24w"
	RadarAs112SummaryResponseCodesParamsDateRange52w        RadarAs112SummaryResponseCodesParamsDateRange = "52w"
	RadarAs112SummaryResponseCodesParamsDateRange1dControl  RadarAs112SummaryResponseCodesParamsDateRange = "1dControl"
	RadarAs112SummaryResponseCodesParamsDateRange2dControl  RadarAs112SummaryResponseCodesParamsDateRange = "2dControl"
	RadarAs112SummaryResponseCodesParamsDateRange7dControl  RadarAs112SummaryResponseCodesParamsDateRange = "7dControl"
	RadarAs112SummaryResponseCodesParamsDateRange14dControl RadarAs112SummaryResponseCodesParamsDateRange = "14dControl"
	RadarAs112SummaryResponseCodesParamsDateRange28dControl RadarAs112SummaryResponseCodesParamsDateRange = "28dControl"
	RadarAs112SummaryResponseCodesParamsDateRange12wControl RadarAs112SummaryResponseCodesParamsDateRange = "12wControl"
	RadarAs112SummaryResponseCodesParamsDateRange24wControl RadarAs112SummaryResponseCodesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryResponseCodesParamsFormat string

const (
	RadarAs112SummaryResponseCodesParamsFormatJson RadarAs112SummaryResponseCodesParamsFormat = "JSON"
	RadarAs112SummaryResponseCodesParamsFormatCsv  RadarAs112SummaryResponseCodesParamsFormat = "CSV"
)

type RadarAs112SummaryResponseCodesResponseEnvelope struct {
	Result  RadarAs112SummaryResponseCodesResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAs112SummaryResponseCodesResponseEnvelopeJSON `json:"-"`
}

// radarAs112SummaryResponseCodesResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAs112SummaryResponseCodesResponseEnvelope]
type radarAs112SummaryResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
