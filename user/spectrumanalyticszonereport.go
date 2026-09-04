// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// SpectrumAnalyticsZoneReportService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpectrumAnalyticsZoneReportService] method instead.
type SpectrumAnalyticsZoneReportService struct {
	Options []option.RequestOption
}

// NewSpectrumAnalyticsZoneReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSpectrumAnalyticsZoneReportService(opts ...option.RequestOption) (r *SpectrumAnalyticsZoneReportService) {
	r = &SpectrumAnalyticsZoneReportService{}
	r.Options = opts
	return
}

// Retrieves a list of total bandwidth by zone over a given time period.
func (r *SpectrumAnalyticsZoneReportService) Get(ctx context.Context, query SpectrumAnalyticsZoneReportGetParams, opts ...option.RequestOption) (res *[]SpectrumAnalyticsZoneReportGetResponse, err error) {
	var env SpectrumAnalyticsZoneReportGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "user/spectrum_analytics/zones/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SpectrumAnalyticsZoneReportGetResponse struct {
	Totals SpectrumAnalyticsZoneReportGetResponseTotals `json:"totals" api:"required"`
	// Identifier.
	ZoneID string                                     `json:"zone_id" api:"required"`
	JSON   spectrumAnalyticsZoneReportGetResponseJSON `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseJSON contains the JSON metadata for the
// struct [SpectrumAnalyticsZoneReportGetResponse]
type spectrumAnalyticsZoneReportGetResponseJSON struct {
	Totals      apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseTotals struct {
	Bandwidth SpectrumAnalyticsZoneReportGetResponseTotalsBandwidth `json:"bandwidth" api:"required"`
	JSON      spectrumAnalyticsZoneReportGetResponseTotalsJSON      `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseTotalsJSON contains the JSON metadata for
// the struct [SpectrumAnalyticsZoneReportGetResponseTotals]
type spectrumAnalyticsZoneReportGetResponseTotalsJSON struct {
	Bandwidth   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseTotalsJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseTotalsBandwidth struct {
	// Sum of ingress and egress bytes transferred.
	All float64 `json:"all" api:"required"`
	// Sum of egress bytes transferred.
	Egress float64 `json:"egress" api:"required"`
	// Sum of ingress bytes transferred.
	Ingress float64                                                   `json:"ingress" api:"required"`
	JSON    spectrumAnalyticsZoneReportGetResponseTotalsBandwidthJSON `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseTotalsBandwidthJSON contains the JSON
// metadata for the struct [SpectrumAnalyticsZoneReportGetResponseTotalsBandwidth]
type spectrumAnalyticsZoneReportGetResponseTotalsBandwidthJSON struct {
	All         apijson.Field
	Egress      apijson.Field
	Ingress     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseTotalsBandwidthJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetParams struct {
	// Include CDN traffic in the bandwidth aggregation.
	CDNTraffic param.Field[bool] `query:"cdn_traffic"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [SpectrumAnalyticsZoneReportGetParams]'s query parameters as
// `url.Values`.
func (r SpectrumAnalyticsZoneReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SpectrumAnalyticsZoneReportGetResponseEnvelope struct {
	Errors   []SpectrumAnalyticsZoneReportGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SpectrumAnalyticsZoneReportGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Query    SpectrumAnalyticsZoneReportGetResponseEnvelopeQuery      `json:"query" api:"required"`
	Result   []SpectrumAnalyticsZoneReportGetResponse                 `json:"result" api:"required"`
	// Whether the API call was successful.
	Success    SpectrumAnalyticsZoneReportGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo SpectrumAnalyticsZoneReportGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       spectrumAnalyticsZoneReportGetResponseEnvelopeJSON       `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SpectrumAnalyticsZoneReportGetResponseEnvelope]
type spectrumAnalyticsZoneReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Query       apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           SpectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SpectrumAnalyticsZoneReportGetResponseEnvelopeErrors]
type spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [SpectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSource]
type spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           SpectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SpectrumAnalyticsZoneReportGetResponseEnvelopeMessages]
type spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SpectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSource]
type spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsZoneReportGetResponseEnvelopeQuery struct {
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since time.Time `json:"since" api:"required" format:"date-time"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until time.Time                                               `json:"until" api:"required" format:"date-time"`
	JSON  spectrumAnalyticsZoneReportGetResponseEnvelopeQueryJSON `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeQueryJSON contains the JSON
// metadata for the struct [SpectrumAnalyticsZoneReportGetResponseEnvelopeQuery]
type spectrumAnalyticsZoneReportGetResponseEnvelopeQueryJSON struct {
	Since       apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelopeQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeQueryJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SpectrumAnalyticsZoneReportGetResponseEnvelopeSuccess bool

const (
	SpectrumAnalyticsZoneReportGetResponseEnvelopeSuccessTrue SpectrumAnalyticsZoneReportGetResponseEnvelopeSuccess = true
)

func (r SpectrumAnalyticsZoneReportGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SpectrumAnalyticsZoneReportGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SpectrumAnalyticsZoneReportGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// The number of total pages in the entire result set.
	TotalPages float64                                                      `json:"total_pages"`
	JSON       spectrumAnalyticsZoneReportGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// spectrumAnalyticsZoneReportGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [SpectrumAnalyticsZoneReportGetResponseEnvelopeResultInfo]
type spectrumAnalyticsZoneReportGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsZoneReportGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsZoneReportGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
