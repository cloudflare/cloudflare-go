// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"errors"
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

// CTLogService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCTLogService] method instead.
type CTLogService struct {
	Options []option.RequestOption
}

// NewCTLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCTLogService(opts ...option.RequestOption) (r *CTLogService) {
	r = &CTLogService{}
	r.Options = opts
	return
}

// Retrieves a list of certificate logs.
func (r *CTLogService) List(ctx context.Context, query CTLogListParams, opts ...option.RequestOption) (res *CTLogListResponse, err error) {
	var env CTLogListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/ct/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested certificate log information.
func (r *CTLogService) Get(ctx context.Context, logSlug string, query CTLogGetParams, opts ...option.RequestOption) (res *CTLogGetResponse, err error) {
	var env CTLogGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if logSlug == "" {
		err = errors.New("missing required log_slug parameter")
		return
	}
	path := fmt.Sprintf("radar/ct/logs/%s", logSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CTLogListResponse struct {
	CertificateLogs []CTLogListResponseCertificateLog `json:"certificateLogs,required"`
	JSON            ctLogListResponseJSON             `json:"-"`
}

// ctLogListResponseJSON contains the JSON metadata for the struct
// [CTLogListResponse]
type ctLogListResponseJSON struct {
	CertificateLogs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CTLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogListResponseJSON) RawJSON() string {
	return r.raw
}

type CTLogListResponseCertificateLog struct {
	// The API standard that the certificate log follows.
	API CTLogListResponseCertificateLogsAPI `json:"api,required"`
	// A brief description of the certificate log.
	Description string `json:"description,required"`
	// The end date and time for when the log will stop accepting certificates.
	EndExclusive time.Time `json:"endExclusive,required" format:"date-time"`
	// The organization responsible for operating the certificate log.
	Operator string `json:"operator,required"`
	// A URL-friendly, kebab-case identifier for the certificate log.
	Slug string `json:"slug,required"`
	// The start date and time for when the log starts accepting certificates.
	StartInclusive time.Time `json:"startInclusive,required" format:"date-time"`
	// The current state of the certificate log. More details about log states can be
	// found here:
	// https://googlechrome.github.io/CertificateTransparency/log_states.html
	State CTLogListResponseCertificateLogsState `json:"state,required"`
	// Timestamp of when the log state was last updated.
	StateTimestamp time.Time `json:"stateTimestamp,required" format:"date-time"`
	// The URL for the certificate log.
	URL  string                              `json:"url,required"`
	JSON ctLogListResponseCertificateLogJSON `json:"-"`
}

// ctLogListResponseCertificateLogJSON contains the JSON metadata for the struct
// [CTLogListResponseCertificateLog]
type ctLogListResponseCertificateLogJSON struct {
	API            apijson.Field
	Description    apijson.Field
	EndExclusive   apijson.Field
	Operator       apijson.Field
	Slug           apijson.Field
	StartInclusive apijson.Field
	State          apijson.Field
	StateTimestamp apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CTLogListResponseCertificateLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogListResponseCertificateLogJSON) RawJSON() string {
	return r.raw
}

// The API standard that the certificate log follows.
type CTLogListResponseCertificateLogsAPI string

const (
	CTLogListResponseCertificateLogsAPIRfc6962 CTLogListResponseCertificateLogsAPI = "RFC6962"
	CTLogListResponseCertificateLogsAPIStatic  CTLogListResponseCertificateLogsAPI = "STATIC"
)

func (r CTLogListResponseCertificateLogsAPI) IsKnown() bool {
	switch r {
	case CTLogListResponseCertificateLogsAPIRfc6962, CTLogListResponseCertificateLogsAPIStatic:
		return true
	}
	return false
}

// The current state of the certificate log. More details about log states can be
// found here:
// https://googlechrome.github.io/CertificateTransparency/log_states.html
type CTLogListResponseCertificateLogsState string

const (
	CTLogListResponseCertificateLogsStateUsable    CTLogListResponseCertificateLogsState = "USABLE"
	CTLogListResponseCertificateLogsStatePending   CTLogListResponseCertificateLogsState = "PENDING"
	CTLogListResponseCertificateLogsStateQualified CTLogListResponseCertificateLogsState = "QUALIFIED"
	CTLogListResponseCertificateLogsStateReadOnly  CTLogListResponseCertificateLogsState = "READ_ONLY"
	CTLogListResponseCertificateLogsStateRetired   CTLogListResponseCertificateLogsState = "RETIRED"
	CTLogListResponseCertificateLogsStateRejected  CTLogListResponseCertificateLogsState = "REJECTED"
)

func (r CTLogListResponseCertificateLogsState) IsKnown() bool {
	switch r {
	case CTLogListResponseCertificateLogsStateUsable, CTLogListResponseCertificateLogsStatePending, CTLogListResponseCertificateLogsStateQualified, CTLogListResponseCertificateLogsStateReadOnly, CTLogListResponseCertificateLogsStateRetired, CTLogListResponseCertificateLogsStateRejected:
		return true
	}
	return false
}

type CTLogGetResponse struct {
	CertificateLog CTLogGetResponseCertificateLog `json:"certificateLog,required"`
	JSON           ctLogGetResponseJSON           `json:"-"`
}

// ctLogGetResponseJSON contains the JSON metadata for the struct
// [CTLogGetResponse]
type ctLogGetResponseJSON struct {
	CertificateLog apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CTLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type CTLogGetResponseCertificateLog struct {
	// The API standard that the certificate log follows.
	API CTLogGetResponseCertificateLogAPI `json:"api,required"`
	// The average throughput of the CT log, measured in certificates per hour
	// (certs/hour).
	AvgThroughput float64 `json:"avgThroughput,required"`
	// A brief description of the certificate log.
	Description string `json:"description,required"`
	// The end date and time for when the log will stop accepting certificates.
	EndExclusive time.Time `json:"endExclusive,required" format:"date-time"`
	// Timestamp of the most recent update to the CT log.
	LastUpdate time.Time `json:"lastUpdate,required" format:"date-time"`
	// The organization responsible for operating the certificate log.
	Operator string `json:"operator,required"`
	// Log performance metrics, including averages and per-endpoint details.
	Performance CTLogGetResponseCertificateLogPerformance `json:"performance,required,nullable"`
	// Logs from the same operator.
	Related []CTLogGetResponseCertificateLogRelated `json:"related,required"`
	// A URL-friendly, kebab-case identifier for the certificate log.
	Slug string `json:"slug,required"`
	// The start date and time for when the log starts accepting certificates.
	StartInclusive time.Time `json:"startInclusive,required" format:"date-time"`
	// The current state of the certificate log. More details about log states can be
	// found here:
	// https://googlechrome.github.io/CertificateTransparency/log_states.html
	State CTLogGetResponseCertificateLogState `json:"state,required"`
	// Timestamp of when the log state was last updated.
	StateTimestamp time.Time `json:"stateTimestamp,required" format:"date-time"`
	// Number of certificates that are eligible for inclusion to this log but have not
	// been included yet. Based on certificates signed by trusted root CAs within the
	// log's accepted date range.
	SubmittableCERTCount string `json:"submittableCertCount,required,nullable"`
	// Number of certificates already included in this CT log.
	SubmittedCERTCount string `json:"submittedCertCount,required,nullable"`
	// The URL for the certificate log.
	URL  string                             `json:"url,required"`
	JSON ctLogGetResponseCertificateLogJSON `json:"-"`
}

// ctLogGetResponseCertificateLogJSON contains the JSON metadata for the struct
// [CTLogGetResponseCertificateLog]
type ctLogGetResponseCertificateLogJSON struct {
	API                  apijson.Field
	AvgThroughput        apijson.Field
	Description          apijson.Field
	EndExclusive         apijson.Field
	LastUpdate           apijson.Field
	Operator             apijson.Field
	Performance          apijson.Field
	Related              apijson.Field
	Slug                 apijson.Field
	StartInclusive       apijson.Field
	State                apijson.Field
	StateTimestamp       apijson.Field
	SubmittableCERTCount apijson.Field
	SubmittedCERTCount   apijson.Field
	URL                  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CTLogGetResponseCertificateLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogJSON) RawJSON() string {
	return r.raw
}

// The API standard that the certificate log follows.
type CTLogGetResponseCertificateLogAPI string

const (
	CTLogGetResponseCertificateLogAPIRfc6962 CTLogGetResponseCertificateLogAPI = "RFC6962"
	CTLogGetResponseCertificateLogAPIStatic  CTLogGetResponseCertificateLogAPI = "STATIC"
)

func (r CTLogGetResponseCertificateLogAPI) IsKnown() bool {
	switch r {
	case CTLogGetResponseCertificateLogAPIRfc6962, CTLogGetResponseCertificateLogAPIStatic:
		return true
	}
	return false
}

// Log performance metrics, including averages and per-endpoint details.
type CTLogGetResponseCertificateLogPerformance struct {
	Endpoints    []CTLogGetResponseCertificateLogPerformanceEndpoint `json:"endpoints,required"`
	ResponseTime float64                                             `json:"responseTime,required"`
	Uptime       float64                                             `json:"uptime,required"`
	JSON         ctLogGetResponseCertificateLogPerformanceJSON       `json:"-"`
}

// ctLogGetResponseCertificateLogPerformanceJSON contains the JSON metadata for the
// struct [CTLogGetResponseCertificateLogPerformance]
type ctLogGetResponseCertificateLogPerformanceJSON struct {
	Endpoints    apijson.Field
	ResponseTime apijson.Field
	Uptime       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CTLogGetResponseCertificateLogPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogPerformanceJSON) RawJSON() string {
	return r.raw
}

type CTLogGetResponseCertificateLogPerformanceEndpoint struct {
	// The certificate log endpoint names used in performance metrics.
	Endpoint     CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint `json:"endpoint,required"`
	ResponseTime float64                                                    `json:"responseTime,required"`
	Uptime       float64                                                    `json:"uptime,required"`
	JSON         ctLogGetResponseCertificateLogPerformanceEndpointJSON      `json:"-"`
}

// ctLogGetResponseCertificateLogPerformanceEndpointJSON contains the JSON metadata
// for the struct [CTLogGetResponseCertificateLogPerformanceEndpoint]
type ctLogGetResponseCertificateLogPerformanceEndpointJSON struct {
	Endpoint     apijson.Field
	ResponseTime apijson.Field
	Uptime       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CTLogGetResponseCertificateLogPerformanceEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogPerformanceEndpointJSON) RawJSON() string {
	return r.raw
}

// The certificate log endpoint names used in performance metrics.
type CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint string

const (
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainNew    CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-chain (new)"
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainOld    CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-chain (old)"
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainNew CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-pre-chain (new)"
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainOld CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-pre-chain (old)"
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointGetEntries     CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "get-entries"
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointGetRoots       CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "get-roots"
	CTLogGetResponseCertificateLogPerformanceEndpointsEndpointGetSth         CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "get-sth"
)

func (r CTLogGetResponseCertificateLogPerformanceEndpointsEndpoint) IsKnown() bool {
	switch r {
	case CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainNew, CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainOld, CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainNew, CTLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainOld, CTLogGetResponseCertificateLogPerformanceEndpointsEndpointGetEntries, CTLogGetResponseCertificateLogPerformanceEndpointsEndpointGetRoots, CTLogGetResponseCertificateLogPerformanceEndpointsEndpointGetSth:
		return true
	}
	return false
}

type CTLogGetResponseCertificateLogRelated struct {
	// A brief description of the certificate log.
	Description string `json:"description,required"`
	// The end date and time for when the log will stop accepting certificates.
	EndExclusive time.Time `json:"endExclusive,required" format:"date-time"`
	// A URL-friendly, kebab-case identifier for the certificate log.
	Slug string `json:"slug,required"`
	// The start date and time for when the log starts accepting certificates.
	StartInclusive time.Time `json:"startInclusive,required" format:"date-time"`
	// The current state of the certificate log. More details about log states can be
	// found here:
	// https://googlechrome.github.io/CertificateTransparency/log_states.html
	State CTLogGetResponseCertificateLogRelatedState `json:"state,required"`
	JSON  ctLogGetResponseCertificateLogRelatedJSON  `json:"-"`
}

// ctLogGetResponseCertificateLogRelatedJSON contains the JSON metadata for the
// struct [CTLogGetResponseCertificateLogRelated]
type ctLogGetResponseCertificateLogRelatedJSON struct {
	Description    apijson.Field
	EndExclusive   apijson.Field
	Slug           apijson.Field
	StartInclusive apijson.Field
	State          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CTLogGetResponseCertificateLogRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogRelatedJSON) RawJSON() string {
	return r.raw
}

// The current state of the certificate log. More details about log states can be
// found here:
// https://googlechrome.github.io/CertificateTransparency/log_states.html
type CTLogGetResponseCertificateLogRelatedState string

const (
	CTLogGetResponseCertificateLogRelatedStateUsable    CTLogGetResponseCertificateLogRelatedState = "USABLE"
	CTLogGetResponseCertificateLogRelatedStatePending   CTLogGetResponseCertificateLogRelatedState = "PENDING"
	CTLogGetResponseCertificateLogRelatedStateQualified CTLogGetResponseCertificateLogRelatedState = "QUALIFIED"
	CTLogGetResponseCertificateLogRelatedStateReadOnly  CTLogGetResponseCertificateLogRelatedState = "READ_ONLY"
	CTLogGetResponseCertificateLogRelatedStateRetired   CTLogGetResponseCertificateLogRelatedState = "RETIRED"
	CTLogGetResponseCertificateLogRelatedStateRejected  CTLogGetResponseCertificateLogRelatedState = "REJECTED"
)

func (r CTLogGetResponseCertificateLogRelatedState) IsKnown() bool {
	switch r {
	case CTLogGetResponseCertificateLogRelatedStateUsable, CTLogGetResponseCertificateLogRelatedStatePending, CTLogGetResponseCertificateLogRelatedStateQualified, CTLogGetResponseCertificateLogRelatedStateReadOnly, CTLogGetResponseCertificateLogRelatedStateRetired, CTLogGetResponseCertificateLogRelatedStateRejected:
		return true
	}
	return false
}

// The current state of the certificate log. More details about log states can be
// found here:
// https://googlechrome.github.io/CertificateTransparency/log_states.html
type CTLogGetResponseCertificateLogState string

const (
	CTLogGetResponseCertificateLogStateUsable    CTLogGetResponseCertificateLogState = "USABLE"
	CTLogGetResponseCertificateLogStatePending   CTLogGetResponseCertificateLogState = "PENDING"
	CTLogGetResponseCertificateLogStateQualified CTLogGetResponseCertificateLogState = "QUALIFIED"
	CTLogGetResponseCertificateLogStateReadOnly  CTLogGetResponseCertificateLogState = "READ_ONLY"
	CTLogGetResponseCertificateLogStateRetired   CTLogGetResponseCertificateLogState = "RETIRED"
	CTLogGetResponseCertificateLogStateRejected  CTLogGetResponseCertificateLogState = "REJECTED"
)

func (r CTLogGetResponseCertificateLogState) IsKnown() bool {
	switch r {
	case CTLogGetResponseCertificateLogStateUsable, CTLogGetResponseCertificateLogStatePending, CTLogGetResponseCertificateLogStateQualified, CTLogGetResponseCertificateLogStateReadOnly, CTLogGetResponseCertificateLogStateRetired, CTLogGetResponseCertificateLogStateRejected:
		return true
	}
	return false
}

type CTLogListParams struct {
	// Format in which results will be returned.
	Format param.Field[CTLogListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CTLogListParams]'s query parameters as `url.Values`.
func (r CTLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CTLogListParamsFormat string

const (
	CTLogListParamsFormatJson CTLogListParamsFormat = "JSON"
	CTLogListParamsFormatCsv  CTLogListParamsFormat = "CSV"
)

func (r CTLogListParamsFormat) IsKnown() bool {
	switch r {
	case CTLogListParamsFormatJson, CTLogListParamsFormatCsv:
		return true
	}
	return false
}

type CTLogListResponseEnvelope struct {
	Result  CTLogListResponse             `json:"result,required"`
	Success bool                          `json:"success,required"`
	JSON    ctLogListResponseEnvelopeJSON `json:"-"`
}

// ctLogListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTLogListResponseEnvelope]
type ctLogListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTLogListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CTLogGetParams struct {
	// Format in which results will be returned.
	Format param.Field[CTLogGetParamsFormat] `query:"format"`
}

// URLQuery serializes [CTLogGetParams]'s query parameters as `url.Values`.
func (r CTLogGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CTLogGetParamsFormat string

const (
	CTLogGetParamsFormatJson CTLogGetParamsFormat = "JSON"
	CTLogGetParamsFormatCsv  CTLogGetParamsFormat = "CSV"
)

func (r CTLogGetParamsFormat) IsKnown() bool {
	switch r {
	case CTLogGetParamsFormatJson, CTLogGetParamsFormatCsv:
		return true
	}
	return false
}

type CTLogGetResponseEnvelope struct {
	Result  CTLogGetResponse             `json:"result,required"`
	Success bool                         `json:"success,required"`
	JSON    ctLogGetResponseEnvelopeJSON `json:"-"`
}

// ctLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTLogGetResponseEnvelope]
type ctLogGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
