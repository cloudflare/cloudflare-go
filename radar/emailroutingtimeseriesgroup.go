// File generated from our OpenAPI spec by Stainless.

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

// EmailRoutingTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingTimeseriesGroupService] method instead.
type EmailRoutingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewEmailRoutingTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingTimeseriesGroupService(opts ...option.RequestOption) (r *EmailRoutingTimeseriesGroupService) {
	r = &EmailRoutingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *EmailRoutingTimeseriesGroupService) ARC(ctx context.Context, query EmailRoutingTimeseriesGroupARCParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingTimeseriesGroupARCResponseEnvelope
	path := "radar/email/routing/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *EmailRoutingTimeseriesGroupService) DKIM(ctx context.Context, query EmailRoutingTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingTimeseriesGroupDKIMResponseEnvelope
	path := "radar/email/routing/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *EmailRoutingTimeseriesGroupService) DMARC(ctx context.Context, query EmailRoutingTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingTimeseriesGroupDMARCResponseEnvelope
	path := "radar/email/routing/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by Encrypted over time.
func (r *EmailRoutingTimeseriesGroupService) Encrypted(ctx context.Context, query EmailRoutingTimeseriesGroupEncryptedParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupEncryptedResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingTimeseriesGroupEncryptedResponseEnvelope
	path := "radar/email/routing/timeseries_groups/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by Ip Version over time.
func (r *EmailRoutingTimeseriesGroupService) IPVersion(ctx context.Context, query EmailRoutingTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingTimeseriesGroupIPVersionResponseEnvelope
	path := "radar/email/routing/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *EmailRoutingTimeseriesGroupService) SPF(ctx context.Context, query EmailRoutingTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupSPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingTimeseriesGroupSPFResponseEnvelope
	path := "radar/email/routing/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingTimeseriesGroupARCResponse struct {
	Meta   interface{}                                  `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupARCResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupARCResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupARCResponse]
type emailRoutingTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupARCResponseSerie0 struct {
	Fail []string                                         `json:"FAIL,required"`
	None []string                                         `json:"NONE,required"`
	Pass []string                                         `json:"PASS,required"`
	JSON emailRoutingTimeseriesGroupARCResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseSerie0JSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupARCResponseSerie0]
type emailRoutingTimeseriesGroupARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupDKIMResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupDKIMResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupDKIMResponse]
type emailRoutingTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMResponseSerie0 struct {
	Fail []string                                          `json:"FAIL,required"`
	None []string                                          `json:"NONE,required"`
	Pass []string                                          `json:"PASS,required"`
	JSON emailRoutingTimeseriesGroupDKIMResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseSerie0JSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupDKIMResponseSerie0]
type emailRoutingTimeseriesGroupDKIMResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupDMARCResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupDMARCResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupDMARCResponse]
type emailRoutingTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCResponseSerie0 struct {
	Fail []string                                           `json:"FAIL,required"`
	None []string                                           `json:"NONE,required"`
	Pass []string                                           `json:"PASS,required"`
	JSON emailRoutingTimeseriesGroupDMARCResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseSerie0JSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDMARCResponseSerie0]
type emailRoutingTimeseriesGroupDMARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupEncryptedResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupEncryptedResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupEncryptedResponse]
type emailRoutingTimeseriesGroupEncryptedResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponseSerie0 struct {
	Encrypted    []string                                               `json:"ENCRYPTED,required"`
	NotEncrypted []string                                               `json:"NOT_ENCRYPTED,required"`
	JSON         emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseSerie0]
type emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupIPVersionResponse]
type emailRoutingTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                               `json:"IPv4,required"`
	IPv6 []string                                               `json:"IPv6,required"`
	JSON emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseSerie0]
type emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFResponse struct {
	Meta   interface{}                                  `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupSPFResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupSPFResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupSPFResponse]
type emailRoutingTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFResponseSerie0 struct {
	Fail []string                                         `json:"FAIL,required"`
	None []string                                         `json:"NONE,required"`
	Pass []string                                         `json:"PASS,required"`
	JSON emailRoutingTimeseriesGroupSPFResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseSerie0JSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupSPFResponseSerie0]
type emailRoutingTimeseriesGroupSPFResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailRoutingTimeseriesGroupARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingTimeseriesGroupARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingTimeseriesGroupARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupARCParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupARCParamsAggInterval15m EmailRoutingTimeseriesGroupARCParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1h  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1d  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1w  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1w"
)

type EmailRoutingTimeseriesGroupARCParamsDateRange string

const (
	EmailRoutingTimeseriesGroupARCParamsDateRange1d         EmailRoutingTimeseriesGroupARCParamsDateRange = "1d"
	EmailRoutingTimeseriesGroupARCParamsDateRange2d         EmailRoutingTimeseriesGroupARCParamsDateRange = "2d"
	EmailRoutingTimeseriesGroupARCParamsDateRange7d         EmailRoutingTimeseriesGroupARCParamsDateRange = "7d"
	EmailRoutingTimeseriesGroupARCParamsDateRange14d        EmailRoutingTimeseriesGroupARCParamsDateRange = "14d"
	EmailRoutingTimeseriesGroupARCParamsDateRange28d        EmailRoutingTimeseriesGroupARCParamsDateRange = "28d"
	EmailRoutingTimeseriesGroupARCParamsDateRange12w        EmailRoutingTimeseriesGroupARCParamsDateRange = "12w"
	EmailRoutingTimeseriesGroupARCParamsDateRange24w        EmailRoutingTimeseriesGroupARCParamsDateRange = "24w"
	EmailRoutingTimeseriesGroupARCParamsDateRange52w        EmailRoutingTimeseriesGroupARCParamsDateRange = "52w"
	EmailRoutingTimeseriesGroupARCParamsDateRange1dControl  EmailRoutingTimeseriesGroupARCParamsDateRange = "1dControl"
	EmailRoutingTimeseriesGroupARCParamsDateRange2dControl  EmailRoutingTimeseriesGroupARCParamsDateRange = "2dControl"
	EmailRoutingTimeseriesGroupARCParamsDateRange7dControl  EmailRoutingTimeseriesGroupARCParamsDateRange = "7dControl"
	EmailRoutingTimeseriesGroupARCParamsDateRange14dControl EmailRoutingTimeseriesGroupARCParamsDateRange = "14dControl"
	EmailRoutingTimeseriesGroupARCParamsDateRange28dControl EmailRoutingTimeseriesGroupARCParamsDateRange = "28dControl"
	EmailRoutingTimeseriesGroupARCParamsDateRange12wControl EmailRoutingTimeseriesGroupARCParamsDateRange = "12wControl"
	EmailRoutingTimeseriesGroupARCParamsDateRange24wControl EmailRoutingTimeseriesGroupARCParamsDateRange = "24wControl"
)

type EmailRoutingTimeseriesGroupARCParamsDKIM string

const (
	EmailRoutingTimeseriesGroupARCParamsDKIMPass EmailRoutingTimeseriesGroupARCParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupARCParamsDKIMNone EmailRoutingTimeseriesGroupARCParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupARCParamsDKIMFail EmailRoutingTimeseriesGroupARCParamsDKIM = "FAIL"
)

type EmailRoutingTimeseriesGroupARCParamsDMARC string

const (
	EmailRoutingTimeseriesGroupARCParamsDMARCPass EmailRoutingTimeseriesGroupARCParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupARCParamsDMARCNone EmailRoutingTimeseriesGroupARCParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupARCParamsDMARCFail EmailRoutingTimeseriesGroupARCParamsDMARC = "FAIL"
)

type EmailRoutingTimeseriesGroupARCParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupARCParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type EmailRoutingTimeseriesGroupARCParamsFormat string

const (
	EmailRoutingTimeseriesGroupARCParamsFormatJson EmailRoutingTimeseriesGroupARCParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupARCParamsFormatCsv  EmailRoutingTimeseriesGroupARCParamsFormat = "CSV"
)

type EmailRoutingTimeseriesGroupARCParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4 EmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6 EmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv6"
)

type EmailRoutingTimeseriesGroupARCParamsSPF string

const (
	EmailRoutingTimeseriesGroupARCParamsSPFPass EmailRoutingTimeseriesGroupARCParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupARCParamsSPFNone EmailRoutingTimeseriesGroupARCParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupARCParamsSPFFail EmailRoutingTimeseriesGroupARCParamsSPF = "FAIL"
)

type EmailRoutingTimeseriesGroupARCResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupARCResponseEnvelope]
type emailRoutingTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupDKIMParams]'s query parameters
// as `url.Values`.
func (r EmailRoutingTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDKIMParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1h  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1d  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1w  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1w"
)

type EmailRoutingTimeseriesGroupDKIMParamsARC string

const (
	EmailRoutingTimeseriesGroupDKIMParamsARCPass EmailRoutingTimeseriesGroupDKIMParamsARC = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsARCNone EmailRoutingTimeseriesGroupDKIMParamsARC = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsARCFail EmailRoutingTimeseriesGroupDKIMParamsARC = "FAIL"
)

type EmailRoutingTimeseriesGroupDKIMParamsDateRange string

const (
	EmailRoutingTimeseriesGroupDKIMParamsDateRange1d         EmailRoutingTimeseriesGroupDKIMParamsDateRange = "1d"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange2d         EmailRoutingTimeseriesGroupDKIMParamsDateRange = "2d"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange7d         EmailRoutingTimeseriesGroupDKIMParamsDateRange = "7d"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange14d        EmailRoutingTimeseriesGroupDKIMParamsDateRange = "14d"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange28d        EmailRoutingTimeseriesGroupDKIMParamsDateRange = "28d"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange12w        EmailRoutingTimeseriesGroupDKIMParamsDateRange = "12w"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange24w        EmailRoutingTimeseriesGroupDKIMParamsDateRange = "24w"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange52w        EmailRoutingTimeseriesGroupDKIMParamsDateRange = "52w"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange1dControl  EmailRoutingTimeseriesGroupDKIMParamsDateRange = "1dControl"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange2dControl  EmailRoutingTimeseriesGroupDKIMParamsDateRange = "2dControl"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange7dControl  EmailRoutingTimeseriesGroupDKIMParamsDateRange = "7dControl"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange14dControl EmailRoutingTimeseriesGroupDKIMParamsDateRange = "14dControl"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange28dControl EmailRoutingTimeseriesGroupDKIMParamsDateRange = "28dControl"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange12wControl EmailRoutingTimeseriesGroupDKIMParamsDateRange = "12wControl"
	EmailRoutingTimeseriesGroupDKIMParamsDateRange24wControl EmailRoutingTimeseriesGroupDKIMParamsDateRange = "24wControl"
)

type EmailRoutingTimeseriesGroupDKIMParamsDMARC string

const (
	EmailRoutingTimeseriesGroupDKIMParamsDMARCPass EmailRoutingTimeseriesGroupDKIMParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsDMARCNone EmailRoutingTimeseriesGroupDKIMParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsDMARCFail EmailRoutingTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

type EmailRoutingTimeseriesGroupDKIMParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupDKIMParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type EmailRoutingTimeseriesGroupDKIMParamsFormat string

const (
	EmailRoutingTimeseriesGroupDKIMParamsFormatJson EmailRoutingTimeseriesGroupDKIMParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupDKIMParamsFormatCsv  EmailRoutingTimeseriesGroupDKIMParamsFormat = "CSV"
)

type EmailRoutingTimeseriesGroupDKIMParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4 EmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6 EmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv6"
)

type EmailRoutingTimeseriesGroupDKIMParamsSPF string

const (
	EmailRoutingTimeseriesGroupDKIMParamsSPFPass EmailRoutingTimeseriesGroupDKIMParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsSPFNone EmailRoutingTimeseriesGroupDKIMParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsSPFFail EmailRoutingTimeseriesGroupDKIMParamsSPF = "FAIL"
)

type EmailRoutingTimeseriesGroupDKIMResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDKIMResponseEnvelope]
type emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupDMARCParams]'s query parameters
// as `url.Values`.
func (r EmailRoutingTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDMARCParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1h  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1d  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1w  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1w"
)

type EmailRoutingTimeseriesGroupDMARCParamsARC string

const (
	EmailRoutingTimeseriesGroupDMARCParamsARCPass EmailRoutingTimeseriesGroupDMARCParamsARC = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsARCNone EmailRoutingTimeseriesGroupDMARCParamsARC = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsARCFail EmailRoutingTimeseriesGroupDMARCParamsARC = "FAIL"
)

type EmailRoutingTimeseriesGroupDMARCParamsDateRange string

const (
	EmailRoutingTimeseriesGroupDMARCParamsDateRange1d         EmailRoutingTimeseriesGroupDMARCParamsDateRange = "1d"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange2d         EmailRoutingTimeseriesGroupDMARCParamsDateRange = "2d"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange7d         EmailRoutingTimeseriesGroupDMARCParamsDateRange = "7d"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange14d        EmailRoutingTimeseriesGroupDMARCParamsDateRange = "14d"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange28d        EmailRoutingTimeseriesGroupDMARCParamsDateRange = "28d"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange12w        EmailRoutingTimeseriesGroupDMARCParamsDateRange = "12w"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange24w        EmailRoutingTimeseriesGroupDMARCParamsDateRange = "24w"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange52w        EmailRoutingTimeseriesGroupDMARCParamsDateRange = "52w"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange1dControl  EmailRoutingTimeseriesGroupDMARCParamsDateRange = "1dControl"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange2dControl  EmailRoutingTimeseriesGroupDMARCParamsDateRange = "2dControl"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange7dControl  EmailRoutingTimeseriesGroupDMARCParamsDateRange = "7dControl"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange14dControl EmailRoutingTimeseriesGroupDMARCParamsDateRange = "14dControl"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange28dControl EmailRoutingTimeseriesGroupDMARCParamsDateRange = "28dControl"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange12wControl EmailRoutingTimeseriesGroupDMARCParamsDateRange = "12wControl"
	EmailRoutingTimeseriesGroupDMARCParamsDateRange24wControl EmailRoutingTimeseriesGroupDMARCParamsDateRange = "24wControl"
)

type EmailRoutingTimeseriesGroupDMARCParamsDKIM string

const (
	EmailRoutingTimeseriesGroupDMARCParamsDKIMPass EmailRoutingTimeseriesGroupDMARCParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsDKIMNone EmailRoutingTimeseriesGroupDMARCParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsDKIMFail EmailRoutingTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

type EmailRoutingTimeseriesGroupDMARCParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupDMARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type EmailRoutingTimeseriesGroupDMARCParamsFormat string

const (
	EmailRoutingTimeseriesGroupDMARCParamsFormatJson EmailRoutingTimeseriesGroupDMARCParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupDMARCParamsFormatCsv  EmailRoutingTimeseriesGroupDMARCParamsFormat = "CSV"
)

type EmailRoutingTimeseriesGroupDMARCParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4 EmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6 EmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv6"
)

type EmailRoutingTimeseriesGroupDMARCParamsSPF string

const (
	EmailRoutingTimeseriesGroupDMARCParamsSPFPass EmailRoutingTimeseriesGroupDMARCParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsSPFNone EmailRoutingTimeseriesGroupDMARCParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsSPFFail EmailRoutingTimeseriesGroupDMARCParamsSPF = "FAIL"
)

type EmailRoutingTimeseriesGroupDMARCResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDMARCResponseEnvelope]
type emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupEncryptedParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailRoutingTimeseriesGroupEncryptedParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupEncryptedParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingTimeseriesGroupEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupEncryptedParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1d  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1w  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1w"
)

type EmailRoutingTimeseriesGroupEncryptedParamsARC string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsARCPass EmailRoutingTimeseriesGroupEncryptedParamsARC = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsARCNone EmailRoutingTimeseriesGroupEncryptedParamsARC = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsARCFail EmailRoutingTimeseriesGroupEncryptedParamsARC = "FAIL"
)

type EmailRoutingTimeseriesGroupEncryptedParamsDateRange string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange1d         EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "1d"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange2d         EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "2d"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange7d         EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "7d"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange14d        EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "14d"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange28d        EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "28d"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange12w        EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "12w"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange24w        EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "24w"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange52w        EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "52w"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange1dControl  EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "1dControl"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange2dControl  EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "2dControl"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange7dControl  EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "7dControl"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange14dControl EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "14dControl"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange28dControl EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "28dControl"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange12wControl EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "12wControl"
	EmailRoutingTimeseriesGroupEncryptedParamsDateRange24wControl EmailRoutingTimeseriesGroupEncryptedParamsDateRange = "24wControl"
)

type EmailRoutingTimeseriesGroupEncryptedParamsDKIM string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "FAIL"
)

type EmailRoutingTimeseriesGroupEncryptedParamsDMARC string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailRoutingTimeseriesGroupEncryptedParamsFormat string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsFormatJson EmailRoutingTimeseriesGroupEncryptedParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupEncryptedParamsFormatCsv  EmailRoutingTimeseriesGroupEncryptedParamsFormat = "CSV"
)

type EmailRoutingTimeseriesGroupEncryptedParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4 EmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6 EmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv6"
)

type EmailRoutingTimeseriesGroupEncryptedParamsSPF string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsSPFPass EmailRoutingTimeseriesGroupEncryptedParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsSPFNone EmailRoutingTimeseriesGroupEncryptedParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsSPFFail EmailRoutingTimeseriesGroupEncryptedParamsSPF = "FAIL"
)

type EmailRoutingTimeseriesGroupEncryptedResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupEncryptedResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseEnvelope]
type emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupIPVersionParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1d  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1w  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type EmailRoutingTimeseriesGroupIPVersionParamsARC string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsARCPass EmailRoutingTimeseriesGroupIPVersionParamsARC = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsARCNone EmailRoutingTimeseriesGroupIPVersionParamsARC = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsARCFail EmailRoutingTimeseriesGroupIPVersionParamsARC = "FAIL"
)

type EmailRoutingTimeseriesGroupIPVersionParamsDateRange string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange1d         EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "1d"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange2d         EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "2d"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange7d         EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "7d"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange14d        EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "14d"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange28d        EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "28d"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange12w        EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "12w"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange24w        EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "24w"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange52w        EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "52w"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange1dControl  EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "1dControl"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange2dControl  EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "2dControl"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange7dControl  EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "7dControl"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange14dControl EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "14dControl"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange28dControl EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "28dControl"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange12wControl EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "12wControl"
	EmailRoutingTimeseriesGroupIPVersionParamsDateRange24wControl EmailRoutingTimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

type EmailRoutingTimeseriesGroupIPVersionParamsDKIM string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "FAIL"
)

type EmailRoutingTimeseriesGroupIPVersionParamsDMARC string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "FAIL"
)

type EmailRoutingTimeseriesGroupIPVersionParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type EmailRoutingTimeseriesGroupIPVersionParamsFormat string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsFormatJson EmailRoutingTimeseriesGroupIPVersionParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupIPVersionParamsFormatCsv  EmailRoutingTimeseriesGroupIPVersionParamsFormat = "CSV"
)

type EmailRoutingTimeseriesGroupIPVersionParamsSPF string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsSPFPass EmailRoutingTimeseriesGroupIPVersionParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsSPFNone EmailRoutingTimeseriesGroupIPVersionParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsSPFFail EmailRoutingTimeseriesGroupIPVersionParamsSPF = "FAIL"
)

type EmailRoutingTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseEnvelope]
type emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingTimeseriesGroupSPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupSPFParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupSPFParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupSPFParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupSPFParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupSPFParamsAggInterval15m EmailRoutingTimeseriesGroupSPFParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1h  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1d  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1w  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1w"
)

type EmailRoutingTimeseriesGroupSPFParamsARC string

const (
	EmailRoutingTimeseriesGroupSPFParamsARCPass EmailRoutingTimeseriesGroupSPFParamsARC = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsARCNone EmailRoutingTimeseriesGroupSPFParamsARC = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsARCFail EmailRoutingTimeseriesGroupSPFParamsARC = "FAIL"
)

type EmailRoutingTimeseriesGroupSPFParamsDateRange string

const (
	EmailRoutingTimeseriesGroupSPFParamsDateRange1d         EmailRoutingTimeseriesGroupSPFParamsDateRange = "1d"
	EmailRoutingTimeseriesGroupSPFParamsDateRange2d         EmailRoutingTimeseriesGroupSPFParamsDateRange = "2d"
	EmailRoutingTimeseriesGroupSPFParamsDateRange7d         EmailRoutingTimeseriesGroupSPFParamsDateRange = "7d"
	EmailRoutingTimeseriesGroupSPFParamsDateRange14d        EmailRoutingTimeseriesGroupSPFParamsDateRange = "14d"
	EmailRoutingTimeseriesGroupSPFParamsDateRange28d        EmailRoutingTimeseriesGroupSPFParamsDateRange = "28d"
	EmailRoutingTimeseriesGroupSPFParamsDateRange12w        EmailRoutingTimeseriesGroupSPFParamsDateRange = "12w"
	EmailRoutingTimeseriesGroupSPFParamsDateRange24w        EmailRoutingTimeseriesGroupSPFParamsDateRange = "24w"
	EmailRoutingTimeseriesGroupSPFParamsDateRange52w        EmailRoutingTimeseriesGroupSPFParamsDateRange = "52w"
	EmailRoutingTimeseriesGroupSPFParamsDateRange1dControl  EmailRoutingTimeseriesGroupSPFParamsDateRange = "1dControl"
	EmailRoutingTimeseriesGroupSPFParamsDateRange2dControl  EmailRoutingTimeseriesGroupSPFParamsDateRange = "2dControl"
	EmailRoutingTimeseriesGroupSPFParamsDateRange7dControl  EmailRoutingTimeseriesGroupSPFParamsDateRange = "7dControl"
	EmailRoutingTimeseriesGroupSPFParamsDateRange14dControl EmailRoutingTimeseriesGroupSPFParamsDateRange = "14dControl"
	EmailRoutingTimeseriesGroupSPFParamsDateRange28dControl EmailRoutingTimeseriesGroupSPFParamsDateRange = "28dControl"
	EmailRoutingTimeseriesGroupSPFParamsDateRange12wControl EmailRoutingTimeseriesGroupSPFParamsDateRange = "12wControl"
	EmailRoutingTimeseriesGroupSPFParamsDateRange24wControl EmailRoutingTimeseriesGroupSPFParamsDateRange = "24wControl"
)

type EmailRoutingTimeseriesGroupSPFParamsDKIM string

const (
	EmailRoutingTimeseriesGroupSPFParamsDKIMPass EmailRoutingTimeseriesGroupSPFParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsDKIMNone EmailRoutingTimeseriesGroupSPFParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsDKIMFail EmailRoutingTimeseriesGroupSPFParamsDKIM = "FAIL"
)

type EmailRoutingTimeseriesGroupSPFParamsDMARC string

const (
	EmailRoutingTimeseriesGroupSPFParamsDMARCPass EmailRoutingTimeseriesGroupSPFParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsDMARCNone EmailRoutingTimeseriesGroupSPFParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsDMARCFail EmailRoutingTimeseriesGroupSPFParamsDMARC = "FAIL"
)

type EmailRoutingTimeseriesGroupSPFParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupSPFParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupSPFParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type EmailRoutingTimeseriesGroupSPFParamsFormat string

const (
	EmailRoutingTimeseriesGroupSPFParamsFormatJson EmailRoutingTimeseriesGroupSPFParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupSPFParamsFormatCsv  EmailRoutingTimeseriesGroupSPFParamsFormat = "CSV"
)

type EmailRoutingTimeseriesGroupSPFParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4 EmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6 EmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv6"
)

type EmailRoutingTimeseriesGroupSPFResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupSPFResponseEnvelope]
type emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
