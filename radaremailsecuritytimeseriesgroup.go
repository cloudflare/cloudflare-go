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

// RadarEmailSecurityTimeseriesGroupService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTimeseriesGroupService] method instead.
type RadarEmailSecurityTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTimeseriesGroupService(opts ...option.RequestOption) (r *RadarEmailSecurityTimeseriesGroupService) {
	r = &RadarEmailSecurityTimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) ARC(ctx context.Context, query RadarEmailSecurityTimeseriesGroupARCParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupARCResponseEnvelope
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) DKIM(ctx context.Context, query RadarEmailSecurityTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) DMARC(ctx context.Context, query RadarEmailSecurityTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupDMARCResponseEnvelope
	path := "radar/email/security/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as MALICIOUS over time.
func (r *RadarEmailSecurityTimeseriesGroupService) Malicious(ctx context.Context, query RadarEmailSecurityTimeseriesGroupMaliciousParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupMaliciousResponseEnvelope
	path := "radar/email/security/timeseries_groups/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as SPAM over time.
func (r *RadarEmailSecurityTimeseriesGroupService) Spam(ctx context.Context, query RadarEmailSecurityTimeseriesGroupSpamParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupSpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupSpamResponseEnvelope
	path := "radar/email/security/timeseries_groups/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) SPF(ctx context.Context, query RadarEmailSecurityTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupSPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupSPFResponseEnvelope
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified in Threat Categories over time.
func (r *RadarEmailSecurityTimeseriesGroupService) ThreatCategory(ctx context.Context, query RadarEmailSecurityTimeseriesGroupThreatCategoryParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupThreatCategoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope
	path := "radar/email/security/timeseries_groups/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTimeseriesGroupARCResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupARCResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupARCResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupARCResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupARCResponse]
type radarEmailSecurityTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupARCResponseSerie0 struct {
	Fail []string                                               `json:"FAIL,required"`
	None []string                                               `json:"NONE,required"`
	Pass []string                                               `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupARCResponseSerie0]
type radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupDKIMResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupDKIMResponse]
type radarEmailSecurityTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0 struct {
	Fail []string                                                `json:"FAIL,required"`
	None []string                                                `json:"NONE,required"`
	Pass []string                                                `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0]
type radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDMARCResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupDMARCResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupDMARCResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDMARCResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupDMARCResponse]
type radarEmailSecurityTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDMARCResponseSerie0 struct {
	Fail []string                                                 `json:"FAIL,required"`
	None []string                                                 `json:"NONE,required"`
	Pass []string                                                 `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupDMARCResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDMARCResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupDMARCResponseSerie0]
type radarEmailSecurityTimeseriesGroupDMARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDMARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupMaliciousResponse struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupMaliciousResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupMaliciousResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupMaliciousResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupMaliciousResponse]
type radarEmailSecurityTimeseriesGroupMaliciousResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupMaliciousResponseSerie0 struct {
	Malicious    []string                                                     `json:"MALICIOUS,required"`
	NotMalicious []string                                                     `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecurityTimeseriesGroupMaliciousResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupMaliciousResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupMaliciousResponseSerie0]
type radarEmailSecurityTimeseriesGroupMaliciousResponseSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupMaliciousResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupSpamResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupSpamResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupSpamResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSpamResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupSpamResponse]
type radarEmailSecurityTimeseriesGroupSpamResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupSpamResponseSerie0 struct {
	NotSpam []string                                                `json:"NOT_SPAM,required"`
	Spam    []string                                                `json:"SPAM,required"`
	JSON    radarEmailSecurityTimeseriesGroupSpamResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSpamResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupSpamResponseSerie0]
type radarEmailSecurityTimeseriesGroupSpamResponseSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSpamResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupSPFResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupSPFResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupSPFResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSPFResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTimeseriesGroupSPFResponse]
type radarEmailSecurityTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupSPFResponseSerie0 struct {
	Fail []string                                               `json:"FAIL,required"`
	None []string                                               `json:"NONE,required"`
	Pass []string                                               `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupSPFResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSPFResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupSPFResponseSerie0]
type radarEmailSecurityTimeseriesGroupSPFResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSPFResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupThreatCategoryResponse struct {
	Meta   interface{}                                                   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupThreatCategoryResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupThreatCategoryResponseJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupThreatCategoryResponse]
type radarEmailSecurityTimeseriesGroupThreatCategoryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0 struct {
	BrandImpersonation  []string                                                          `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                                          `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                                          `json:"IdentityDeception,required"`
	Link                []string                                                          `json:"Link,required"`
	JSON                radarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0]
type radarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupARCParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupARCParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupARCParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval15m RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupARCParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupARCParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange1d         RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange2d         RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange7d         RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange14d        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange28d        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange12w        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange24w        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange52w        RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupARCParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupARCParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupARCParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsDKIMPass RadarEmailSecurityTimeseriesGroupARCParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupARCParamsDKIMNone RadarEmailSecurityTimeseriesGroupARCParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupARCParamsDKIMFail RadarEmailSecurityTimeseriesGroupARCParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupARCParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsDMARCPass RadarEmailSecurityTimeseriesGroupARCParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupARCParamsDMARCNone RadarEmailSecurityTimeseriesGroupARCParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupARCParamsDMARCFail RadarEmailSecurityTimeseriesGroupARCParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupARCParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsFormatJson RadarEmailSecurityTimeseriesGroupARCParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupARCParamsFormatCsv  RadarEmailSecurityTimeseriesGroupARCParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupARCParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsSPFPass RadarEmailSecurityTimeseriesGroupARCParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupARCParamsSPFNone RadarEmailSecurityTimeseriesGroupARCParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupARCParamsSPFFail RadarEmailSecurityTimeseriesGroupARCParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupARCResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupARCResponseEnvelope]
type radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupDKIMParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval15m RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsARCPass RadarEmailSecurityTimeseriesGroupDKIMParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsARCNone RadarEmailSecurityTimeseriesGroupDKIMParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsARCFail RadarEmailSecurityTimeseriesGroupDKIMParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange1d         RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange2d         RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange7d         RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange14d        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange28d        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange12w        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange24w        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange52w        RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCPass RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCNone RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCFail RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupDKIMParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsFormatJson RadarEmailSecurityTimeseriesGroupDKIMParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupDKIMParamsFormatCsv  RadarEmailSecurityTimeseriesGroupDKIMParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupDKIMParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsSPFPass RadarEmailSecurityTimeseriesGroupDKIMParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupDKIMParamsSPFNone RadarEmailSecurityTimeseriesGroupDKIMParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupDKIMParamsSPFFail RadarEmailSecurityTimeseriesGroupDKIMParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope]
type radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupDMARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupDMARCParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval15m RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupDMARCParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsARCPass RadarEmailSecurityTimeseriesGroupDMARCParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupDMARCParamsARCNone RadarEmailSecurityTimeseriesGroupDMARCParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupDMARCParamsARCFail RadarEmailSecurityTimeseriesGroupDMARCParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange1d         RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange2d         RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange7d         RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange14d        RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange28d        RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange12w        RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange24w        RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange52w        RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupDMARCParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsDKIMPass RadarEmailSecurityTimeseriesGroupDMARCParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDKIMNone RadarEmailSecurityTimeseriesGroupDMARCParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupDMARCParamsDKIMFail RadarEmailSecurityTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupDMARCParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsFormatJson RadarEmailSecurityTimeseriesGroupDMARCParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupDMARCParamsFormatCsv  RadarEmailSecurityTimeseriesGroupDMARCParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupDMARCParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsSPFPass RadarEmailSecurityTimeseriesGroupDMARCParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupDMARCParamsSPFNone RadarEmailSecurityTimeseriesGroupDMARCParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupDMARCParamsSPFFail RadarEmailSecurityTimeseriesGroupDMARCParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupDMARCResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupDMARCResponseEnvelope]
type radarEmailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupMaliciousParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupMaliciousParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupMaliciousParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupMaliciousParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsARCPass RadarEmailSecurityTimeseriesGroupMaliciousParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsARCNone RadarEmailSecurityTimeseriesGroupMaliciousParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsARCFail RadarEmailSecurityTimeseriesGroupMaliciousParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange1d         RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange2d         RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange7d         RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange14d        RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange28d        RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange12w        RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange24w        RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange52w        RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIMPass RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIMNone RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIMFail RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARCPass RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARCNone RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARCFail RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupMaliciousParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsFormatJson RadarEmailSecurityTimeseriesGroupMaliciousParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsFormatCsv  RadarEmailSecurityTimeseriesGroupMaliciousParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsSPFPass RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsSPFNone RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsSPFFail RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupMaliciousResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupMaliciousResponse             `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupMaliciousResponseEnvelope]
type radarEmailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupMaliciousResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupSpamParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupSpamParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupSpamParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupSpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval15m RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupSpamParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsARCPass RadarEmailSecurityTimeseriesGroupSpamParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupSpamParamsARCNone RadarEmailSecurityTimeseriesGroupSpamParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupSpamParamsARCFail RadarEmailSecurityTimeseriesGroupSpamParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSpamParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange1d         RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange2d         RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange7d         RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange14d        RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange28d        RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange12w        RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange24w        RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange52w        RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupSpamParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupSpamParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupSpamParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsDKIMPass RadarEmailSecurityTimeseriesGroupSpamParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupSpamParamsDKIMNone RadarEmailSecurityTimeseriesGroupSpamParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupSpamParamsDKIMFail RadarEmailSecurityTimeseriesGroupSpamParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSpamParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsDMARCPass RadarEmailSecurityTimeseriesGroupSpamParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupSpamParamsDMARCNone RadarEmailSecurityTimeseriesGroupSpamParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupSpamParamsDMARCFail RadarEmailSecurityTimeseriesGroupSpamParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupSpamParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsFormatJson RadarEmailSecurityTimeseriesGroupSpamParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupSpamParamsFormatCsv  RadarEmailSecurityTimeseriesGroupSpamParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupSpamParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsSPFPass RadarEmailSecurityTimeseriesGroupSpamParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupSpamParamsSPFNone RadarEmailSecurityTimeseriesGroupSpamParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupSpamParamsSPFFail RadarEmailSecurityTimeseriesGroupSpamParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSpamResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupSpamResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupSpamResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSpamResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupSpamResponseEnvelope]
type radarEmailSecurityTimeseriesGroupSpamResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSpamResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupSPFParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupSPFParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupSPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupSPFParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval15m RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupSPFParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsARCPass RadarEmailSecurityTimeseriesGroupSPFParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupSPFParamsARCNone RadarEmailSecurityTimeseriesGroupSPFParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupSPFParamsARCFail RadarEmailSecurityTimeseriesGroupSPFParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSPFParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange1d         RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange2d         RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange7d         RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange14d        RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange28d        RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange12w        RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange24w        RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange52w        RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupSPFParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupSPFParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupSPFParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsDKIMPass RadarEmailSecurityTimeseriesGroupSPFParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupSPFParamsDKIMNone RadarEmailSecurityTimeseriesGroupSPFParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupSPFParamsDKIMFail RadarEmailSecurityTimeseriesGroupSPFParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSPFParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsDMARCPass RadarEmailSecurityTimeseriesGroupSPFParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupSPFParamsDMARCNone RadarEmailSecurityTimeseriesGroupSPFParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupSPFParamsDMARCFail RadarEmailSecurityTimeseriesGroupSPFParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupSPFParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsFormatJson RadarEmailSecurityTimeseriesGroupSPFParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupSPFParamsFormatCsv  RadarEmailSecurityTimeseriesGroupSPFParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupSPFResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupSPFResponseEnvelope]
type radarEmailSecurityTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTimeseriesGroupThreatCategoryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupThreatCategoryParams]'s
// query parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARCPass RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARCNone RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARCFail RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1d         RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2d         RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7d         RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange14d        RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange28d        RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange12w        RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange24w        RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange52w        RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormatCsv  RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupThreatCategoryResponse             `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope]
type radarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
