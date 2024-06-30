// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// EmailSecurityTimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTimeseriesGroupService] method instead.
type EmailSecurityTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTimeseriesGroupService(opts ...option.RequestOption) (r *EmailSecurityTimeseriesGroupService) {
	r = &EmailSecurityTimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *EmailSecurityTimeseriesGroupService) ARC(ctx context.Context, query EmailSecurityTimeseriesGroupARCParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupARCResponseEnvelope
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *EmailSecurityTimeseriesGroupService) DKIM(ctx context.Context, query EmailSecurityTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupDKIMResponseEnvelope
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *EmailSecurityTimeseriesGroupService) DMARC(ctx context.Context, query EmailSecurityTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupDMARCResponseEnvelope
	path := "radar/email/security/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as MALICIOUS over time.
func (r *EmailSecurityTimeseriesGroupService) Malicious(ctx context.Context, query EmailSecurityTimeseriesGroupMaliciousParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupMaliciousResponseEnvelope
	path := "radar/email/security/timeseries_groups/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as SPAM over time.
func (r *EmailSecurityTimeseriesGroupService) Spam(ctx context.Context, query EmailSecurityTimeseriesGroupSpamParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupSpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupSpamResponseEnvelope
	path := "radar/email/security/timeseries_groups/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *EmailSecurityTimeseriesGroupService) SPF(ctx context.Context, query EmailSecurityTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupSPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupSPFResponseEnvelope
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as SPOOF over time.
func (r *EmailSecurityTimeseriesGroupService) Spoof(ctx context.Context, query EmailSecurityTimeseriesGroupSpoofParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupSpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupSpoofResponseEnvelope
	path := "radar/email/security/timeseries_groups/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified in Threat Categories over time.
func (r *EmailSecurityTimeseriesGroupService) ThreatCategory(ctx context.Context, query EmailSecurityTimeseriesGroupThreatCategoryParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupThreatCategoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope
	path := "radar/email/security/timeseries_groups/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per TLS Version over time.
func (r *EmailSecurityTimeseriesGroupService) TLSVersion(ctx context.Context, query EmailSecurityTimeseriesGroupTLSVersionParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupTLSVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope
	path := "radar/email/security/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTimeseriesGroupARCResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 RadarEmailSeries                            `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupARCResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupARCResponse]
type emailSecurityTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDKIMResponse struct {
	Meta   interface{}                                  `json:"meta,required"`
	Serie0 RadarEmailSeries                             `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupDKIMResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupDKIMResponse]
type emailSecurityTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDMARCResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarEmailSeries                              `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupDMARCResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupDMARCResponse]
type emailSecurityTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupMaliciousResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupMaliciousResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupMaliciousResponse]
type emailSecurityTimeseriesGroupMaliciousResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousResponseSerie0 struct {
	Malicious    []string                                                `json:"MALICIOUS,required"`
	NotMalicious []string                                                `json:"NOT_MALICIOUS,required"`
	JSON         emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupMaliciousResponseSerie0]
type emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupSpamResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupSpamResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupSpamResponse]
type emailSecurityTimeseriesGroupSpamResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamResponseSerie0 struct {
	NotSpam []string                                           `json:"NOT_SPAM,required"`
	Spam    []string                                           `json:"SPAM,required"`
	JSON    emailSecurityTimeseriesGroupSpamResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseSerie0JSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpamResponseSerie0]
type emailSecurityTimeseriesGroupSpamResponseSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSPFResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 RadarEmailSeries                            `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupSPFResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupSPFResponse]
type emailSecurityTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofResponse struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupSpoofResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupSpoofResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupSpoofResponse]
type emailSecurityTimeseriesGroupSpoofResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofResponseSerie0 struct {
	NotSpoof []string                                            `json:"NOT_SPOOF,required"`
	Spoof    []string                                            `json:"SPOOF,required"`
	JSON     emailSecurityTimeseriesGroupSpoofResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseSerie0JSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpoofResponseSerie0]
type emailSecurityTimeseriesGroupSpoofResponseSerie0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryResponse struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupThreatCategoryResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupThreatCategoryResponse]
type emailSecurityTimeseriesGroupThreatCategoryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0 struct {
	BrandImpersonation  []string                                                     `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                                     `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                                     `json:"IdentityDeception,required"`
	Link                []string                                                     `json:"Link,required"`
	JSON                emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0]
type emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupTLSVersionResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupTLSVersionResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupTLSVersionResponse]
type emailSecurityTimeseriesGroupTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionResponseSerie0 struct {
	TLS1_0 []string                                                 `json:"TLS 1.0,required"`
	TLS1_1 []string                                                 `json:"TLS 1.1,required"`
	TLS1_2 []string                                                 `json:"TLS 1.2,required"`
	TLS1_3 []string                                                 `json:"TLS 1.3,required"`
	JSON   emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupTLSVersionResponseSerie0]
type emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON struct {
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupARCParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupARCParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupARCParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupARCParamsAggInterval15m EmailSecurityTimeseriesGroupARCParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupARCParamsAggInterval1h  EmailSecurityTimeseriesGroupARCParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupARCParamsAggInterval1d  EmailSecurityTimeseriesGroupARCParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupARCParamsAggInterval1w  EmailSecurityTimeseriesGroupARCParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsAggInterval15m, EmailSecurityTimeseriesGroupARCParamsAggInterval1h, EmailSecurityTimeseriesGroupARCParamsAggInterval1d, EmailSecurityTimeseriesGroupARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsDateRange string

const (
	EmailSecurityTimeseriesGroupARCParamsDateRange1d         EmailSecurityTimeseriesGroupARCParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupARCParamsDateRange2d         EmailSecurityTimeseriesGroupARCParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupARCParamsDateRange7d         EmailSecurityTimeseriesGroupARCParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupARCParamsDateRange14d        EmailSecurityTimeseriesGroupARCParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupARCParamsDateRange28d        EmailSecurityTimeseriesGroupARCParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupARCParamsDateRange12w        EmailSecurityTimeseriesGroupARCParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupARCParamsDateRange24w        EmailSecurityTimeseriesGroupARCParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupARCParamsDateRange52w        EmailSecurityTimeseriesGroupARCParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupARCParamsDateRange1dControl  EmailSecurityTimeseriesGroupARCParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupARCParamsDateRange2dControl  EmailSecurityTimeseriesGroupARCParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupARCParamsDateRange7dControl  EmailSecurityTimeseriesGroupARCParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupARCParamsDateRange14dControl EmailSecurityTimeseriesGroupARCParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupARCParamsDateRange28dControl EmailSecurityTimeseriesGroupARCParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupARCParamsDateRange12wControl EmailSecurityTimeseriesGroupARCParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupARCParamsDateRange24wControl EmailSecurityTimeseriesGroupARCParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupARCParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsDateRange1d, EmailSecurityTimeseriesGroupARCParamsDateRange2d, EmailSecurityTimeseriesGroupARCParamsDateRange7d, EmailSecurityTimeseriesGroupARCParamsDateRange14d, EmailSecurityTimeseriesGroupARCParamsDateRange28d, EmailSecurityTimeseriesGroupARCParamsDateRange12w, EmailSecurityTimeseriesGroupARCParamsDateRange24w, EmailSecurityTimeseriesGroupARCParamsDateRange52w, EmailSecurityTimeseriesGroupARCParamsDateRange1dControl, EmailSecurityTimeseriesGroupARCParamsDateRange2dControl, EmailSecurityTimeseriesGroupARCParamsDateRange7dControl, EmailSecurityTimeseriesGroupARCParamsDateRange14dControl, EmailSecurityTimeseriesGroupARCParamsDateRange28dControl, EmailSecurityTimeseriesGroupARCParamsDateRange12wControl, EmailSecurityTimeseriesGroupARCParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsDKIM string

const (
	EmailSecurityTimeseriesGroupARCParamsDKIMPass EmailSecurityTimeseriesGroupARCParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupARCParamsDKIMNone EmailSecurityTimeseriesGroupARCParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupARCParamsDKIMFail EmailSecurityTimeseriesGroupARCParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsDKIMPass, EmailSecurityTimeseriesGroupARCParamsDKIMNone, EmailSecurityTimeseriesGroupARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsDMARC string

const (
	EmailSecurityTimeseriesGroupARCParamsDMARCPass EmailSecurityTimeseriesGroupARCParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupARCParamsDMARCNone EmailSecurityTimeseriesGroupARCParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupARCParamsDMARCFail EmailSecurityTimeseriesGroupARCParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsDMARCPass, EmailSecurityTimeseriesGroupARCParamsDMARCNone, EmailSecurityTimeseriesGroupARCParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupARCParamsFormat string

const (
	EmailSecurityTimeseriesGroupARCParamsFormatJson EmailSecurityTimeseriesGroupARCParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupARCParamsFormatCsv  EmailSecurityTimeseriesGroupARCParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsFormatJson, EmailSecurityTimeseriesGroupARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsSPF string

const (
	EmailSecurityTimeseriesGroupARCParamsSPFPass EmailSecurityTimeseriesGroupARCParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupARCParamsSPFNone EmailSecurityTimeseriesGroupARCParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupARCParamsSPFFail EmailSecurityTimeseriesGroupARCParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsSPFPass, EmailSecurityTimeseriesGroupARCParamsSPFNone, EmailSecurityTimeseriesGroupARCParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupARCParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupARCResponseEnvelope]
type emailSecurityTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupDKIMParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupDKIMParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval15m EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval1h  EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval1d  EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval1w  EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsAggInterval15m, EmailSecurityTimeseriesGroupDKIMParamsAggInterval1h, EmailSecurityTimeseriesGroupDKIMParamsAggInterval1d, EmailSecurityTimeseriesGroupDKIMParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsARC string

const (
	EmailSecurityTimeseriesGroupDKIMParamsARCPass EmailSecurityTimeseriesGroupDKIMParamsARC = "PASS"
	EmailSecurityTimeseriesGroupDKIMParamsARCNone EmailSecurityTimeseriesGroupDKIMParamsARC = "NONE"
	EmailSecurityTimeseriesGroupDKIMParamsARCFail EmailSecurityTimeseriesGroupDKIMParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsARCPass, EmailSecurityTimeseriesGroupDKIMParamsARCNone, EmailSecurityTimeseriesGroupDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsDateRange string

const (
	EmailSecurityTimeseriesGroupDKIMParamsDateRange1d         EmailSecurityTimeseriesGroupDKIMParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange2d         EmailSecurityTimeseriesGroupDKIMParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange7d         EmailSecurityTimeseriesGroupDKIMParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange14d        EmailSecurityTimeseriesGroupDKIMParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange28d        EmailSecurityTimeseriesGroupDKIMParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange12w        EmailSecurityTimeseriesGroupDKIMParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange24w        EmailSecurityTimeseriesGroupDKIMParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange52w        EmailSecurityTimeseriesGroupDKIMParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange1dControl  EmailSecurityTimeseriesGroupDKIMParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange2dControl  EmailSecurityTimeseriesGroupDKIMParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange7dControl  EmailSecurityTimeseriesGroupDKIMParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange14dControl EmailSecurityTimeseriesGroupDKIMParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange28dControl EmailSecurityTimeseriesGroupDKIMParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange12wControl EmailSecurityTimeseriesGroupDKIMParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupDKIMParamsDateRange24wControl EmailSecurityTimeseriesGroupDKIMParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsDateRange1d, EmailSecurityTimeseriesGroupDKIMParamsDateRange2d, EmailSecurityTimeseriesGroupDKIMParamsDateRange7d, EmailSecurityTimeseriesGroupDKIMParamsDateRange14d, EmailSecurityTimeseriesGroupDKIMParamsDateRange28d, EmailSecurityTimeseriesGroupDKIMParamsDateRange12w, EmailSecurityTimeseriesGroupDKIMParamsDateRange24w, EmailSecurityTimeseriesGroupDKIMParamsDateRange52w, EmailSecurityTimeseriesGroupDKIMParamsDateRange1dControl, EmailSecurityTimeseriesGroupDKIMParamsDateRange2dControl, EmailSecurityTimeseriesGroupDKIMParamsDateRange7dControl, EmailSecurityTimeseriesGroupDKIMParamsDateRange14dControl, EmailSecurityTimeseriesGroupDKIMParamsDateRange28dControl, EmailSecurityTimeseriesGroupDKIMParamsDateRange12wControl, EmailSecurityTimeseriesGroupDKIMParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsDMARC string

const (
	EmailSecurityTimeseriesGroupDKIMParamsDMARCPass EmailSecurityTimeseriesGroupDKIMParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupDKIMParamsDMARCNone EmailSecurityTimeseriesGroupDKIMParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupDKIMParamsDMARCFail EmailSecurityTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsDMARCPass, EmailSecurityTimeseriesGroupDKIMParamsDMARCNone, EmailSecurityTimeseriesGroupDKIMParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupDKIMParamsFormat string

const (
	EmailSecurityTimeseriesGroupDKIMParamsFormatJson EmailSecurityTimeseriesGroupDKIMParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupDKIMParamsFormatCsv  EmailSecurityTimeseriesGroupDKIMParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsFormatJson, EmailSecurityTimeseriesGroupDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsSPF string

const (
	EmailSecurityTimeseriesGroupDKIMParamsSPFPass EmailSecurityTimeseriesGroupDKIMParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupDKIMParamsSPFNone EmailSecurityTimeseriesGroupDKIMParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupDKIMParamsSPFFail EmailSecurityTimeseriesGroupDKIMParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsSPFPass, EmailSecurityTimeseriesGroupDKIMParamsSPFNone, EmailSecurityTimeseriesGroupDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupDKIMResponseEnvelope]
type emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDMARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupDMARCParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupDMARCParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval15m EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval1h  EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval1d  EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval1w  EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsAggInterval15m, EmailSecurityTimeseriesGroupDMARCParamsAggInterval1h, EmailSecurityTimeseriesGroupDMARCParamsAggInterval1d, EmailSecurityTimeseriesGroupDMARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsARC string

const (
	EmailSecurityTimeseriesGroupDMARCParamsARCPass EmailSecurityTimeseriesGroupDMARCParamsARC = "PASS"
	EmailSecurityTimeseriesGroupDMARCParamsARCNone EmailSecurityTimeseriesGroupDMARCParamsARC = "NONE"
	EmailSecurityTimeseriesGroupDMARCParamsARCFail EmailSecurityTimeseriesGroupDMARCParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsARCPass, EmailSecurityTimeseriesGroupDMARCParamsARCNone, EmailSecurityTimeseriesGroupDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsDateRange string

const (
	EmailSecurityTimeseriesGroupDMARCParamsDateRange1d         EmailSecurityTimeseriesGroupDMARCParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange2d         EmailSecurityTimeseriesGroupDMARCParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange7d         EmailSecurityTimeseriesGroupDMARCParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange14d        EmailSecurityTimeseriesGroupDMARCParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange28d        EmailSecurityTimeseriesGroupDMARCParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange12w        EmailSecurityTimeseriesGroupDMARCParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange24w        EmailSecurityTimeseriesGroupDMARCParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange52w        EmailSecurityTimeseriesGroupDMARCParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange1dControl  EmailSecurityTimeseriesGroupDMARCParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange2dControl  EmailSecurityTimeseriesGroupDMARCParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange7dControl  EmailSecurityTimeseriesGroupDMARCParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange14dControl EmailSecurityTimeseriesGroupDMARCParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange28dControl EmailSecurityTimeseriesGroupDMARCParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange12wControl EmailSecurityTimeseriesGroupDMARCParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupDMARCParamsDateRange24wControl EmailSecurityTimeseriesGroupDMARCParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsDateRange1d, EmailSecurityTimeseriesGroupDMARCParamsDateRange2d, EmailSecurityTimeseriesGroupDMARCParamsDateRange7d, EmailSecurityTimeseriesGroupDMARCParamsDateRange14d, EmailSecurityTimeseriesGroupDMARCParamsDateRange28d, EmailSecurityTimeseriesGroupDMARCParamsDateRange12w, EmailSecurityTimeseriesGroupDMARCParamsDateRange24w, EmailSecurityTimeseriesGroupDMARCParamsDateRange52w, EmailSecurityTimeseriesGroupDMARCParamsDateRange1dControl, EmailSecurityTimeseriesGroupDMARCParamsDateRange2dControl, EmailSecurityTimeseriesGroupDMARCParamsDateRange7dControl, EmailSecurityTimeseriesGroupDMARCParamsDateRange14dControl, EmailSecurityTimeseriesGroupDMARCParamsDateRange28dControl, EmailSecurityTimeseriesGroupDMARCParamsDateRange12wControl, EmailSecurityTimeseriesGroupDMARCParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsDKIM string

const (
	EmailSecurityTimeseriesGroupDMARCParamsDKIMPass EmailSecurityTimeseriesGroupDMARCParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupDMARCParamsDKIMNone EmailSecurityTimeseriesGroupDMARCParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupDMARCParamsDKIMFail EmailSecurityTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsDKIMPass, EmailSecurityTimeseriesGroupDMARCParamsDKIMNone, EmailSecurityTimeseriesGroupDMARCParamsDKIMFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupDMARCParamsFormat string

const (
	EmailSecurityTimeseriesGroupDMARCParamsFormatJson EmailSecurityTimeseriesGroupDMARCParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupDMARCParamsFormatCsv  EmailSecurityTimeseriesGroupDMARCParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsFormatJson, EmailSecurityTimeseriesGroupDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsSPF string

const (
	EmailSecurityTimeseriesGroupDMARCParamsSPFPass EmailSecurityTimeseriesGroupDMARCParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupDMARCParamsSPFNone EmailSecurityTimeseriesGroupDMARCParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupDMARCParamsSPFFail EmailSecurityTimeseriesGroupDMARCParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsSPFPass, EmailSecurityTimeseriesGroupDMARCParamsSPFNone, EmailSecurityTimeseriesGroupDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupDMARCResponseEnvelope]
type emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupMaliciousParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupMaliciousParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupMaliciousParams]'s query
// parameters as `url.Values`.
func (r EmailSecurityTimeseriesGroupMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupMaliciousParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h  EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1d  EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1w  EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m, EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h, EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1d, EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsARC string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsARCPass EmailSecurityTimeseriesGroupMaliciousParamsARC = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsARCNone EmailSecurityTimeseriesGroupMaliciousParamsARC = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsARCFail EmailSecurityTimeseriesGroupMaliciousParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsARCPass, EmailSecurityTimeseriesGroupMaliciousParamsARCNone, EmailSecurityTimeseriesGroupMaliciousParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsDateRange string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange1d         EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange2d         EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange7d         EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange14d        EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange28d        EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange12w        EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange24w        EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange52w        EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange1dControl  EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange2dControl  EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange7dControl  EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange14dControl EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange28dControl EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange12wControl EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupMaliciousParamsDateRange24wControl EmailSecurityTimeseriesGroupMaliciousParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsDateRange1d, EmailSecurityTimeseriesGroupMaliciousParamsDateRange2d, EmailSecurityTimeseriesGroupMaliciousParamsDateRange7d, EmailSecurityTimeseriesGroupMaliciousParamsDateRange14d, EmailSecurityTimeseriesGroupMaliciousParamsDateRange28d, EmailSecurityTimeseriesGroupMaliciousParamsDateRange12w, EmailSecurityTimeseriesGroupMaliciousParamsDateRange24w, EmailSecurityTimeseriesGroupMaliciousParamsDateRange52w, EmailSecurityTimeseriesGroupMaliciousParamsDateRange1dControl, EmailSecurityTimeseriesGroupMaliciousParamsDateRange2dControl, EmailSecurityTimeseriesGroupMaliciousParamsDateRange7dControl, EmailSecurityTimeseriesGroupMaliciousParamsDateRange14dControl, EmailSecurityTimeseriesGroupMaliciousParamsDateRange28dControl, EmailSecurityTimeseriesGroupMaliciousParamsDateRange12wControl, EmailSecurityTimeseriesGroupMaliciousParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsDKIM string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass EmailSecurityTimeseriesGroupMaliciousParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsDKIMNone EmailSecurityTimeseriesGroupMaliciousParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsDKIMFail EmailSecurityTimeseriesGroupMaliciousParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass, EmailSecurityTimeseriesGroupMaliciousParamsDKIMNone, EmailSecurityTimeseriesGroupMaliciousParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsDMARC string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass EmailSecurityTimeseriesGroupMaliciousParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsDMARCNone EmailSecurityTimeseriesGroupMaliciousParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsDMARCFail EmailSecurityTimeseriesGroupMaliciousParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass, EmailSecurityTimeseriesGroupMaliciousParamsDMARCNone, EmailSecurityTimeseriesGroupMaliciousParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupMaliciousParamsFormat string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsFormatJson EmailSecurityTimeseriesGroupMaliciousParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupMaliciousParamsFormatCsv  EmailSecurityTimeseriesGroupMaliciousParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsFormatJson, EmailSecurityTimeseriesGroupMaliciousParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsSPF string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsSPFPass EmailSecurityTimeseriesGroupMaliciousParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsSPFNone EmailSecurityTimeseriesGroupMaliciousParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsSPFFail EmailSecurityTimeseriesGroupMaliciousParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsSPFPass, EmailSecurityTimeseriesGroupMaliciousParamsSPFNone, EmailSecurityTimeseriesGroupMaliciousParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupMaliciousResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupMaliciousResponseEnvelope]
type emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupSpamParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupSpamParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupSpamParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupSpamParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupSpamParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupSpamParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupSpamParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupSpamParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupSpamParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupSpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSpamParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupSpamParamsAggInterval15m EmailSecurityTimeseriesGroupSpamParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupSpamParamsAggInterval1h  EmailSecurityTimeseriesGroupSpamParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupSpamParamsAggInterval1d  EmailSecurityTimeseriesGroupSpamParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupSpamParamsAggInterval1w  EmailSecurityTimeseriesGroupSpamParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupSpamParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsAggInterval15m, EmailSecurityTimeseriesGroupSpamParamsAggInterval1h, EmailSecurityTimeseriesGroupSpamParamsAggInterval1d, EmailSecurityTimeseriesGroupSpamParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsARC string

const (
	EmailSecurityTimeseriesGroupSpamParamsARCPass EmailSecurityTimeseriesGroupSpamParamsARC = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsARCNone EmailSecurityTimeseriesGroupSpamParamsARC = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsARCFail EmailSecurityTimeseriesGroupSpamParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsARCPass, EmailSecurityTimeseriesGroupSpamParamsARCNone, EmailSecurityTimeseriesGroupSpamParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsDateRange string

const (
	EmailSecurityTimeseriesGroupSpamParamsDateRange1d         EmailSecurityTimeseriesGroupSpamParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupSpamParamsDateRange2d         EmailSecurityTimeseriesGroupSpamParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupSpamParamsDateRange7d         EmailSecurityTimeseriesGroupSpamParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupSpamParamsDateRange14d        EmailSecurityTimeseriesGroupSpamParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupSpamParamsDateRange28d        EmailSecurityTimeseriesGroupSpamParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupSpamParamsDateRange12w        EmailSecurityTimeseriesGroupSpamParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupSpamParamsDateRange24w        EmailSecurityTimeseriesGroupSpamParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupSpamParamsDateRange52w        EmailSecurityTimeseriesGroupSpamParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupSpamParamsDateRange1dControl  EmailSecurityTimeseriesGroupSpamParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupSpamParamsDateRange2dControl  EmailSecurityTimeseriesGroupSpamParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupSpamParamsDateRange7dControl  EmailSecurityTimeseriesGroupSpamParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupSpamParamsDateRange14dControl EmailSecurityTimeseriesGroupSpamParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupSpamParamsDateRange28dControl EmailSecurityTimeseriesGroupSpamParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupSpamParamsDateRange12wControl EmailSecurityTimeseriesGroupSpamParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupSpamParamsDateRange24wControl EmailSecurityTimeseriesGroupSpamParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupSpamParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsDateRange1d, EmailSecurityTimeseriesGroupSpamParamsDateRange2d, EmailSecurityTimeseriesGroupSpamParamsDateRange7d, EmailSecurityTimeseriesGroupSpamParamsDateRange14d, EmailSecurityTimeseriesGroupSpamParamsDateRange28d, EmailSecurityTimeseriesGroupSpamParamsDateRange12w, EmailSecurityTimeseriesGroupSpamParamsDateRange24w, EmailSecurityTimeseriesGroupSpamParamsDateRange52w, EmailSecurityTimeseriesGroupSpamParamsDateRange1dControl, EmailSecurityTimeseriesGroupSpamParamsDateRange2dControl, EmailSecurityTimeseriesGroupSpamParamsDateRange7dControl, EmailSecurityTimeseriesGroupSpamParamsDateRange14dControl, EmailSecurityTimeseriesGroupSpamParamsDateRange28dControl, EmailSecurityTimeseriesGroupSpamParamsDateRange12wControl, EmailSecurityTimeseriesGroupSpamParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsDKIM string

const (
	EmailSecurityTimeseriesGroupSpamParamsDKIMPass EmailSecurityTimeseriesGroupSpamParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsDKIMNone EmailSecurityTimeseriesGroupSpamParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsDKIMFail EmailSecurityTimeseriesGroupSpamParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsDKIMPass, EmailSecurityTimeseriesGroupSpamParamsDKIMNone, EmailSecurityTimeseriesGroupSpamParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsDMARC string

const (
	EmailSecurityTimeseriesGroupSpamParamsDMARCPass EmailSecurityTimeseriesGroupSpamParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsDMARCNone EmailSecurityTimeseriesGroupSpamParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsDMARCFail EmailSecurityTimeseriesGroupSpamParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsDMARCPass, EmailSecurityTimeseriesGroupSpamParamsDMARCNone, EmailSecurityTimeseriesGroupSpamParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupSpamParamsFormat string

const (
	EmailSecurityTimeseriesGroupSpamParamsFormatJson EmailSecurityTimeseriesGroupSpamParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupSpamParamsFormatCsv  EmailSecurityTimeseriesGroupSpamParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupSpamParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsFormatJson, EmailSecurityTimeseriesGroupSpamParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsSPF string

const (
	EmailSecurityTimeseriesGroupSpamParamsSPFPass EmailSecurityTimeseriesGroupSpamParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsSPFNone EmailSecurityTimeseriesGroupSpamParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsSPFFail EmailSecurityTimeseriesGroupSpamParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsSPFPass, EmailSecurityTimeseriesGroupSpamParamsSPFNone, EmailSecurityTimeseriesGroupSpamParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupSpamParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupSpamResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpamResponseEnvelope]
type emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSPFParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupSPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupSPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupSPFParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupSPFParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSPFParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupSPFParamsAggInterval15m EmailSecurityTimeseriesGroupSPFParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupSPFParamsAggInterval1h  EmailSecurityTimeseriesGroupSPFParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupSPFParamsAggInterval1d  EmailSecurityTimeseriesGroupSPFParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupSPFParamsAggInterval1w  EmailSecurityTimeseriesGroupSPFParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupSPFParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsAggInterval15m, EmailSecurityTimeseriesGroupSPFParamsAggInterval1h, EmailSecurityTimeseriesGroupSPFParamsAggInterval1d, EmailSecurityTimeseriesGroupSPFParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsARC string

const (
	EmailSecurityTimeseriesGroupSPFParamsARCPass EmailSecurityTimeseriesGroupSPFParamsARC = "PASS"
	EmailSecurityTimeseriesGroupSPFParamsARCNone EmailSecurityTimeseriesGroupSPFParamsARC = "NONE"
	EmailSecurityTimeseriesGroupSPFParamsARCFail EmailSecurityTimeseriesGroupSPFParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSPFParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsARCPass, EmailSecurityTimeseriesGroupSPFParamsARCNone, EmailSecurityTimeseriesGroupSPFParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsDateRange string

const (
	EmailSecurityTimeseriesGroupSPFParamsDateRange1d         EmailSecurityTimeseriesGroupSPFParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupSPFParamsDateRange2d         EmailSecurityTimeseriesGroupSPFParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupSPFParamsDateRange7d         EmailSecurityTimeseriesGroupSPFParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupSPFParamsDateRange14d        EmailSecurityTimeseriesGroupSPFParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupSPFParamsDateRange28d        EmailSecurityTimeseriesGroupSPFParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupSPFParamsDateRange12w        EmailSecurityTimeseriesGroupSPFParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupSPFParamsDateRange24w        EmailSecurityTimeseriesGroupSPFParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupSPFParamsDateRange52w        EmailSecurityTimeseriesGroupSPFParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupSPFParamsDateRange1dControl  EmailSecurityTimeseriesGroupSPFParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupSPFParamsDateRange2dControl  EmailSecurityTimeseriesGroupSPFParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupSPFParamsDateRange7dControl  EmailSecurityTimeseriesGroupSPFParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupSPFParamsDateRange14dControl EmailSecurityTimeseriesGroupSPFParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupSPFParamsDateRange28dControl EmailSecurityTimeseriesGroupSPFParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupSPFParamsDateRange12wControl EmailSecurityTimeseriesGroupSPFParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupSPFParamsDateRange24wControl EmailSecurityTimeseriesGroupSPFParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupSPFParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsDateRange1d, EmailSecurityTimeseriesGroupSPFParamsDateRange2d, EmailSecurityTimeseriesGroupSPFParamsDateRange7d, EmailSecurityTimeseriesGroupSPFParamsDateRange14d, EmailSecurityTimeseriesGroupSPFParamsDateRange28d, EmailSecurityTimeseriesGroupSPFParamsDateRange12w, EmailSecurityTimeseriesGroupSPFParamsDateRange24w, EmailSecurityTimeseriesGroupSPFParamsDateRange52w, EmailSecurityTimeseriesGroupSPFParamsDateRange1dControl, EmailSecurityTimeseriesGroupSPFParamsDateRange2dControl, EmailSecurityTimeseriesGroupSPFParamsDateRange7dControl, EmailSecurityTimeseriesGroupSPFParamsDateRange14dControl, EmailSecurityTimeseriesGroupSPFParamsDateRange28dControl, EmailSecurityTimeseriesGroupSPFParamsDateRange12wControl, EmailSecurityTimeseriesGroupSPFParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsDKIM string

const (
	EmailSecurityTimeseriesGroupSPFParamsDKIMPass EmailSecurityTimeseriesGroupSPFParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupSPFParamsDKIMNone EmailSecurityTimeseriesGroupSPFParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupSPFParamsDKIMFail EmailSecurityTimeseriesGroupSPFParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsDKIMPass, EmailSecurityTimeseriesGroupSPFParamsDKIMNone, EmailSecurityTimeseriesGroupSPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsDMARC string

const (
	EmailSecurityTimeseriesGroupSPFParamsDMARCPass EmailSecurityTimeseriesGroupSPFParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupSPFParamsDMARCNone EmailSecurityTimeseriesGroupSPFParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupSPFParamsDMARCFail EmailSecurityTimeseriesGroupSPFParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsDMARCPass, EmailSecurityTimeseriesGroupSPFParamsDMARCNone, EmailSecurityTimeseriesGroupSPFParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupSPFParamsFormat string

const (
	EmailSecurityTimeseriesGroupSPFParamsFormatJson EmailSecurityTimeseriesGroupSPFParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupSPFParamsFormatCsv  EmailSecurityTimeseriesGroupSPFParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupSPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsFormatJson, EmailSecurityTimeseriesGroupSPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupSPFParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSPFResponseEnvelope]
type emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupSpoofParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupSpoofParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupSpoofParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupSpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSpoofParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval15m EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval1h  EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval1d  EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval1w  EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsAggInterval15m, EmailSecurityTimeseriesGroupSpoofParamsAggInterval1h, EmailSecurityTimeseriesGroupSpoofParamsAggInterval1d, EmailSecurityTimeseriesGroupSpoofParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsARC string

const (
	EmailSecurityTimeseriesGroupSpoofParamsARCPass EmailSecurityTimeseriesGroupSpoofParamsARC = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsARCNone EmailSecurityTimeseriesGroupSpoofParamsARC = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsARCFail EmailSecurityTimeseriesGroupSpoofParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsARCPass, EmailSecurityTimeseriesGroupSpoofParamsARCNone, EmailSecurityTimeseriesGroupSpoofParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsDateRange string

const (
	EmailSecurityTimeseriesGroupSpoofParamsDateRange1d         EmailSecurityTimeseriesGroupSpoofParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange2d         EmailSecurityTimeseriesGroupSpoofParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange7d         EmailSecurityTimeseriesGroupSpoofParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange14d        EmailSecurityTimeseriesGroupSpoofParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange28d        EmailSecurityTimeseriesGroupSpoofParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange12w        EmailSecurityTimeseriesGroupSpoofParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange24w        EmailSecurityTimeseriesGroupSpoofParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange52w        EmailSecurityTimeseriesGroupSpoofParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange1dControl  EmailSecurityTimeseriesGroupSpoofParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange2dControl  EmailSecurityTimeseriesGroupSpoofParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange7dControl  EmailSecurityTimeseriesGroupSpoofParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange14dControl EmailSecurityTimeseriesGroupSpoofParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange28dControl EmailSecurityTimeseriesGroupSpoofParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange12wControl EmailSecurityTimeseriesGroupSpoofParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupSpoofParamsDateRange24wControl EmailSecurityTimeseriesGroupSpoofParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsDateRange1d, EmailSecurityTimeseriesGroupSpoofParamsDateRange2d, EmailSecurityTimeseriesGroupSpoofParamsDateRange7d, EmailSecurityTimeseriesGroupSpoofParamsDateRange14d, EmailSecurityTimeseriesGroupSpoofParamsDateRange28d, EmailSecurityTimeseriesGroupSpoofParamsDateRange12w, EmailSecurityTimeseriesGroupSpoofParamsDateRange24w, EmailSecurityTimeseriesGroupSpoofParamsDateRange52w, EmailSecurityTimeseriesGroupSpoofParamsDateRange1dControl, EmailSecurityTimeseriesGroupSpoofParamsDateRange2dControl, EmailSecurityTimeseriesGroupSpoofParamsDateRange7dControl, EmailSecurityTimeseriesGroupSpoofParamsDateRange14dControl, EmailSecurityTimeseriesGroupSpoofParamsDateRange28dControl, EmailSecurityTimeseriesGroupSpoofParamsDateRange12wControl, EmailSecurityTimeseriesGroupSpoofParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsDKIM string

const (
	EmailSecurityTimeseriesGroupSpoofParamsDKIMPass EmailSecurityTimeseriesGroupSpoofParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsDKIMNone EmailSecurityTimeseriesGroupSpoofParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsDKIMFail EmailSecurityTimeseriesGroupSpoofParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsDKIMPass, EmailSecurityTimeseriesGroupSpoofParamsDKIMNone, EmailSecurityTimeseriesGroupSpoofParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsDMARC string

const (
	EmailSecurityTimeseriesGroupSpoofParamsDMARCPass EmailSecurityTimeseriesGroupSpoofParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsDMARCNone EmailSecurityTimeseriesGroupSpoofParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsDMARCFail EmailSecurityTimeseriesGroupSpoofParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsDMARCPass, EmailSecurityTimeseriesGroupSpoofParamsDMARCNone, EmailSecurityTimeseriesGroupSpoofParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupSpoofParamsFormat string

const (
	EmailSecurityTimeseriesGroupSpoofParamsFormatJson EmailSecurityTimeseriesGroupSpoofParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupSpoofParamsFormatCsv  EmailSecurityTimeseriesGroupSpoofParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsFormatJson, EmailSecurityTimeseriesGroupSpoofParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsSPF string

const (
	EmailSecurityTimeseriesGroupSpoofParamsSPFPass EmailSecurityTimeseriesGroupSpoofParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsSPFNone EmailSecurityTimeseriesGroupSpoofParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsSPFFail EmailSecurityTimeseriesGroupSpoofParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsSPFPass, EmailSecurityTimeseriesGroupSpoofParamsSPFNone, EmailSecurityTimeseriesGroupSpoofParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupSpoofResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpoofResponseEnvelope]
type emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupThreatCategoryParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r EmailSecurityTimeseriesGroupThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h  EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1d  EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1w  EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m, EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h, EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1d, EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsARC string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass EmailSecurityTimeseriesGroupThreatCategoryParamsARC = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsARCNone EmailSecurityTimeseriesGroupThreatCategoryParamsARC = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsARCFail EmailSecurityTimeseriesGroupThreatCategoryParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass, EmailSecurityTimeseriesGroupThreatCategoryParamsARCNone, EmailSecurityTimeseriesGroupThreatCategoryParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1d         EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2d         EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7d         EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange14d        EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange28d        EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange12w        EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange24w        EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange52w        EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1dControl  EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2dControl  EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7dControl  EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange14dControl EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange28dControl EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange12wControl EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange24wControl EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1d, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2d, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7d, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange14d, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange28d, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange12w, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange24w, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange52w, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1dControl, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2dControl, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7dControl, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange14dControl, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange28dControl, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange12wControl, EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass, EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone, EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass, EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone, EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupThreatCategoryParamsFormat string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson EmailSecurityTimeseriesGroupThreatCategoryParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupThreatCategoryParamsFormatCsv  EmailSecurityTimeseriesGroupThreatCategoryParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson, EmailSecurityTimeseriesGroupThreatCategoryParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsSPF string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass EmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone EmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail EmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass, EmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone, EmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupThreatCategoryResponse             `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope]
type emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTimeseriesGroupTLSVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupTLSVersionParams]'s query
// parameters as `url.Values`.
func (r EmailSecurityTimeseriesGroupTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h  EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1d  EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1w  EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m, EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h, EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1d, EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsARC string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsARCPass EmailSecurityTimeseriesGroupTLSVersionParamsARC = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsARCNone EmailSecurityTimeseriesGroupTLSVersionParamsARC = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsARCFail EmailSecurityTimeseriesGroupTLSVersionParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsARCPass, EmailSecurityTimeseriesGroupTLSVersionParamsARCNone, EmailSecurityTimeseriesGroupTLSVersionParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsDateRange string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange1d         EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "1d"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange2d         EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "2d"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange7d         EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "7d"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange14d        EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "14d"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange28d        EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "28d"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange12w        EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "12w"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange24w        EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "24w"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange52w        EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "52w"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange1dControl  EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "1dControl"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange2dControl  EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "2dControl"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange7dControl  EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "7dControl"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange14dControl EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "14dControl"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange28dControl EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "28dControl"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange12wControl EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "12wControl"
	EmailSecurityTimeseriesGroupTLSVersionParamsDateRange24wControl EmailSecurityTimeseriesGroupTLSVersionParamsDateRange = "24wControl"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsDateRange1d, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange2d, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange7d, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange14d, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange28d, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange12w, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange24w, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange52w, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange1dControl, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange2dControl, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange7dControl, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange14dControl, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange28dControl, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange12wControl, EmailSecurityTimeseriesGroupTLSVersionParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsDKIM string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass EmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone EmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail EmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass, EmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone, EmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsDMARC string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass EmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone EmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail EmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass, EmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone, EmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTimeseriesGroupTLSVersionParamsFormat string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson EmailSecurityTimeseriesGroupTLSVersionParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupTLSVersionParamsFormatCsv  EmailSecurityTimeseriesGroupTLSVersionParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson, EmailSecurityTimeseriesGroupTLSVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsSPF string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass EmailSecurityTimeseriesGroupTLSVersionParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsSPFNone EmailSecurityTimeseriesGroupTLSVersionParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsSPFFail EmailSecurityTimeseriesGroupTLSVersionParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass, EmailSecurityTimeseriesGroupTLSVersionParamsSPFNone, EmailSecurityTimeseriesGroupTLSVersionParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupTLSVersionResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope]
type emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
