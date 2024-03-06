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

// Percentage distribution of emails classified as SPOOF over time.
func (r *RadarEmailSecurityTimeseriesGroupService) Spoof(ctx context.Context, query RadarEmailSecurityTimeseriesGroupSpoofParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupSpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupSpoofResponseEnvelope
	path := "radar/email/security/timeseries_groups/spoof"
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

// Percentage distribution of emails classified per TLS Version over time.
func (r *RadarEmailSecurityTimeseriesGroupService) TLSVersion(ctx context.Context, query RadarEmailSecurityTimeseriesGroupTLSVersionParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupTLSVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelope
	path := "radar/email/security/timeseries_groups/tls_version"
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

func (r radarEmailSecurityTimeseriesGroupARCResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupARCResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupDKIMResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupDKIMResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupDMARCResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupDMARCResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupMaliciousResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupMaliciousResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupSpamResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupSpamResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupSPFResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupSPFResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupSpoofResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupSpoofResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupSpoofResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSpoofResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupSpoofResponse]
type radarEmailSecurityTimeseriesGroupSpoofResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupSpoofResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupSpoofResponseSerie0 struct {
	NotSpoof []string                                                 `json:"NOT_SPOOF,required"`
	Spoof    []string                                                 `json:"SPOOF,required"`
	JSON     radarEmailSecurityTimeseriesGroupSpoofResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSpoofResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupSpoofResponseSerie0]
type radarEmailSecurityTimeseriesGroupSpoofResponseSerie0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSpoofResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupSpoofResponseSerie0JSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
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

func (r radarEmailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupTLSVersionResponse struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupTLSVersionResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupTLSVersionResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupTLSVersionResponse]
type radarEmailSecurityTimeseriesGroupTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0 struct {
	TLS1_0 []string                                                      `json:"TLS 1.0,required"`
	TLS1_1 []string                                                      `json:"TLS 1.1,required"`
	TLS1_2 []string                                                      `json:"TLS 1.2,required"`
	TLS1_3 []string                                                      `json:"TLS 1.3,required"`
	JSON   radarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0]
type radarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON struct {
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsARC] `query:"arc"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupDMARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsARC] `query:"arc"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupMaliciousParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsARC] `query:"arc"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupSpamParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsARC] `query:"arc"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupSpamResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupSPFParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupSPFParamsARC] `query:"arc"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupSPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupSpoofParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupSpoofParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupSpoofParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupSpoofParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupSpoofParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupSpoofParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupSpoofParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupSpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval15m RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupSpoofParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsARCPass RadarEmailSecurityTimeseriesGroupSpoofParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupSpoofParamsARCNone RadarEmailSecurityTimeseriesGroupSpoofParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupSpoofParamsARCFail RadarEmailSecurityTimeseriesGroupSpoofParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange1d         RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange2d         RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange7d         RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange14d        RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange28d        RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange12w        RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange24w        RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange52w        RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupSpoofParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsDKIMPass RadarEmailSecurityTimeseriesGroupSpoofParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDKIMNone RadarEmailSecurityTimeseriesGroupSpoofParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDKIMFail RadarEmailSecurityTimeseriesGroupSpoofParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSpoofParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsDMARCPass RadarEmailSecurityTimeseriesGroupSpoofParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDMARCNone RadarEmailSecurityTimeseriesGroupSpoofParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupSpoofParamsDMARCFail RadarEmailSecurityTimeseriesGroupSpoofParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupSpoofParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsFormatJson RadarEmailSecurityTimeseriesGroupSpoofParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupSpoofParamsFormatCsv  RadarEmailSecurityTimeseriesGroupSpoofParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupSpoofParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsSPFPass RadarEmailSecurityTimeseriesGroupSpoofParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupSpoofParamsSPFNone RadarEmailSecurityTimeseriesGroupSpoofParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupSpoofParamsSPFFail RadarEmailSecurityTimeseriesGroupSpoofParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecurityTimeseriesGroupSpoofResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupSpoofResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupSpoofResponseEnvelope]
type radarEmailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupSpoofResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupThreatCategoryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC] `query:"arc"`
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
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion] `query:"tlsVersion"`
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

type RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion string

const (
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_3"
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

func (r radarEmailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupTLSVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTimeseriesGroupTLSVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupTLSVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1w"
)

type RadarEmailSecurityTimeseriesGroupTLSVersionParamsARC string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsARCPass RadarEmailSecurityTimeseriesGroupTLSVersionParamsARC = "PASS"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsARCNone RadarEmailSecurityTimeseriesGroupTLSVersionParamsARC = "NONE"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsARCFail RadarEmailSecurityTimeseriesGroupTLSVersionParamsARC = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange1d         RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "1d"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange2d         RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "2d"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange7d         RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "7d"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange14d        RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "14d"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange28d        RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "28d"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange12w        RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "12w"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange24w        RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "24w"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange52w        RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "52w"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange1dControl  RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "1dControl"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange2dControl  RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "2dControl"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange7dControl  RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "7dControl"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange14dControl RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "14dControl"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange28dControl RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "28dControl"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange12wControl RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "12wControl"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange24wControl RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "24wControl"
)

type RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIM string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "PASS"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "NONE"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARC string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "PASS"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "NONE"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormatJson RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormatCsv  RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormat = "CSV"
)

type RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPF string

const (
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPFPass RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPF = "PASS"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPFNone RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPF = "NONE"
	RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPFFail RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPF = "FAIL"
)

type RadarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelope struct {
	Result  RadarEmailSecurityTimeseriesGroupTLSVersionResponse             `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelope]
type radarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
