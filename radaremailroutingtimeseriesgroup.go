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

// RadarEmailRoutingTimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailRoutingTimeseriesGroupService] method instead.
type RadarEmailRoutingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarEmailRoutingTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailRoutingTimeseriesGroupService(opts ...option.RequestOption) (r *RadarEmailRoutingTimeseriesGroupService) {
	r = &RadarEmailRoutingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per Arc validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) ARC(ctx context.Context, query RadarEmailRoutingTimeseriesGroupARCParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingTimeseriesGroupARCResponseEnvelope
	path := "radar/email/routing/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) DKIM(ctx context.Context, query RadarEmailRoutingTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingTimeseriesGroupDKIMResponseEnvelope
	path := "radar/email/routing/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) DMARC(ctx context.Context, query RadarEmailRoutingTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingTimeseriesGroupDMARCResponseEnvelope
	path := "radar/email/routing/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by Encrypted over time.
func (r *RadarEmailRoutingTimeseriesGroupService) Encrypted(ctx context.Context, query RadarEmailRoutingTimeseriesGroupEncryptedParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupEncryptedResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingTimeseriesGroupEncryptedResponseEnvelope
	path := "radar/email/routing/timeseries_groups/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by Ip Version over time.
func (r *RadarEmailRoutingTimeseriesGroupService) IPVersion(ctx context.Context, query RadarEmailRoutingTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingTimeseriesGroupIPVersionResponseEnvelope
	path := "radar/email/routing/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) SPF(ctx context.Context, query RadarEmailRoutingTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupSPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingTimeseriesGroupSPFResponseEnvelope
	path := "radar/email/routing/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailRoutingTimeseriesGroupARCResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupARCResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupARCResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupARCResponseJSON contains the JSON metadata for
// the struct [RadarEmailRoutingTimeseriesGroupARCResponse]
type radarEmailRoutingTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupARCResponseSerie0 struct {
	Fail []string                                              `json:"FAIL,required"`
	None []string                                              `json:"NONE,required"`
	Pass []string                                              `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupARCResponseSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupARCResponseSerie0JSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupARCResponseSerie0]
type radarEmailRoutingTimeseriesGroupARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupDKIMResponse struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupDKIMResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupDKIMResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupDKIMResponseJSON contains the JSON metadata for
// the struct [RadarEmailRoutingTimeseriesGroupDKIMResponse]
type radarEmailRoutingTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupDKIMResponseSerie0 struct {
	Fail []string                                               `json:"FAIL,required"`
	None []string                                               `json:"NONE,required"`
	Pass []string                                               `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupDKIMResponseSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupDKIMResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupDKIMResponseSerie0]
type radarEmailRoutingTimeseriesGroupDKIMResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupDKIMResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupDMARCResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupDMARCResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupDMARCResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupDMARCResponseJSON contains the JSON metadata for
// the struct [RadarEmailRoutingTimeseriesGroupDMARCResponse]
type radarEmailRoutingTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupDMARCResponseSerie0 struct {
	Fail []string                                                `json:"FAIL,required"`
	None []string                                                `json:"NONE,required"`
	Pass []string                                                `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupDMARCResponseSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupDMARCResponseSerie0JSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupDMARCResponseSerie0]
type radarEmailRoutingTimeseriesGroupDMARCResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupDMARCResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupEncryptedResponse struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupEncryptedResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupEncryptedResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupEncryptedResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupEncryptedResponse]
type radarEmailRoutingTimeseriesGroupEncryptedResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupEncryptedResponseSerie0 struct {
	Encrypted    []string                                                    `json:"ENCRYPTED,required"`
	NotEncrypted []string                                                    `json:"NOT_ENCRYPTED,required"`
	JSON         radarEmailRoutingTimeseriesGroupEncryptedResponseSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupEncryptedResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupEncryptedResponseSerie0]
type radarEmailRoutingTimeseriesGroupEncryptedResponseSerie0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupEncryptedResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupIPVersionResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupIPVersionResponse]
type radarEmailRoutingTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                                    `json:"IPv4,required"`
	IPv6 []string                                                    `json:"IPv6,required"`
	JSON radarEmailRoutingTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupIPVersionResponseSerie0]
type radarEmailRoutingTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupSPFResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupSPFResponseSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupSPFResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupSPFResponseJSON contains the JSON metadata for
// the struct [RadarEmailRoutingTimeseriesGroupSPFResponse]
type radarEmailRoutingTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupSPFResponseSerie0 struct {
	Fail []string                                              `json:"FAIL,required"`
	None []string                                              `json:"NONE,required"`
	Pass []string                                              `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupSPFResponseSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupSPFResponseSerie0JSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupSPFResponseSerie0]
type radarEmailRoutingTimeseriesGroupSPFResponseSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupSPFResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingTimeseriesGroupARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingTimeseriesGroupARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingTimeseriesGroupARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupARCParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupARCParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsAggInterval15m RadarEmailRoutingTimeseriesGroupARCParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupARCParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupARCParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupARCParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupARCParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupARCParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupARCParamsAggInterval = "1w"
)

type RadarEmailRoutingTimeseriesGroupARCParamsDateRange string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange1d         RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "1d"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange2d         RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "2d"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange7d         RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "7d"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange14d        RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "14d"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange28d        RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "28d"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange12w        RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "12w"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange24w        RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "24w"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange52w        RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "52w"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange1dControl  RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "1dControl"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange2dControl  RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "2dControl"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange7dControl  RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "7dControl"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange14dControl RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "14dControl"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange28dControl RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "28dControl"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange12wControl RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "12wControl"
	RadarEmailRoutingTimeseriesGroupARCParamsDateRange24wControl RadarEmailRoutingTimeseriesGroupARCParamsDateRange = "24wControl"
)

type RadarEmailRoutingTimeseriesGroupARCParamsDKIM string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsDKIMPass RadarEmailRoutingTimeseriesGroupARCParamsDKIM = "PASS"
	RadarEmailRoutingTimeseriesGroupARCParamsDKIMNone RadarEmailRoutingTimeseriesGroupARCParamsDKIM = "NONE"
	RadarEmailRoutingTimeseriesGroupARCParamsDKIMFail RadarEmailRoutingTimeseriesGroupARCParamsDKIM = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupARCParamsDMARC string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsDMARCPass RadarEmailRoutingTimeseriesGroupARCParamsDMARC = "PASS"
	RadarEmailRoutingTimeseriesGroupARCParamsDMARCNone RadarEmailRoutingTimeseriesGroupARCParamsDMARC = "NONE"
	RadarEmailRoutingTimeseriesGroupARCParamsDMARCFail RadarEmailRoutingTimeseriesGroupARCParamsDMARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupARCParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupARCParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupARCParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingTimeseriesGroupARCParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsFormatJson RadarEmailRoutingTimeseriesGroupARCParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupARCParamsFormatCsv  RadarEmailRoutingTimeseriesGroupARCParamsFormat = "CSV"
)

type RadarEmailRoutingTimeseriesGroupARCParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupARCParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv6"
)

type RadarEmailRoutingTimeseriesGroupARCParamsSPF string

const (
	RadarEmailRoutingTimeseriesGroupARCParamsSPFPass RadarEmailRoutingTimeseriesGroupARCParamsSPF = "PASS"
	RadarEmailRoutingTimeseriesGroupARCParamsSPFNone RadarEmailRoutingTimeseriesGroupARCParamsSPF = "NONE"
	RadarEmailRoutingTimeseriesGroupARCParamsSPFFail RadarEmailRoutingTimeseriesGroupARCParamsSPF = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupARCResponseEnvelope struct {
	Result  RadarEmailRoutingTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupARCResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupARCResponseEnvelope]
type radarEmailRoutingTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupDKIMParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingTimeseriesGroupDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupDKIMParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingTimeseriesGroupDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupDKIMParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval15m RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1w"
)

type RadarEmailRoutingTimeseriesGroupDKIMParamsARC string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsARCPass RadarEmailRoutingTimeseriesGroupDKIMParamsARC = "PASS"
	RadarEmailRoutingTimeseriesGroupDKIMParamsARCNone RadarEmailRoutingTimeseriesGroupDKIMParamsARC = "NONE"
	RadarEmailRoutingTimeseriesGroupDKIMParamsARCFail RadarEmailRoutingTimeseriesGroupDKIMParamsARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange1d         RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "1d"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange2d         RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "2d"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange7d         RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "7d"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange14d        RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "14d"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange28d        RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "28d"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange12w        RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "12w"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange24w        RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "24w"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange52w        RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "52w"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange1dControl  RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "1dControl"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange2dControl  RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "2dControl"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange7dControl  RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "7dControl"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange14dControl RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "14dControl"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange28dControl RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "28dControl"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange12wControl RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "12wControl"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange24wControl RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange = "24wControl"
)

type RadarEmailRoutingTimeseriesGroupDKIMParamsDMARC string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsDMARCPass RadarEmailRoutingTimeseriesGroupDKIMParamsDMARC = "PASS"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDMARCNone RadarEmailRoutingTimeseriesGroupDKIMParamsDMARC = "NONE"
	RadarEmailRoutingTimeseriesGroupDKIMParamsDMARCFail RadarEmailRoutingTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupDKIMParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupDKIMParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingTimeseriesGroupDKIMParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsFormatJson RadarEmailRoutingTimeseriesGroupDKIMParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupDKIMParamsFormatCsv  RadarEmailRoutingTimeseriesGroupDKIMParamsFormat = "CSV"
)

type RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv6"
)

type RadarEmailRoutingTimeseriesGroupDKIMParamsSPF string

const (
	RadarEmailRoutingTimeseriesGroupDKIMParamsSPFPass RadarEmailRoutingTimeseriesGroupDKIMParamsSPF = "PASS"
	RadarEmailRoutingTimeseriesGroupDKIMParamsSPFNone RadarEmailRoutingTimeseriesGroupDKIMParamsSPF = "NONE"
	RadarEmailRoutingTimeseriesGroupDKIMParamsSPFFail RadarEmailRoutingTimeseriesGroupDKIMParamsSPF = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupDKIMResponseEnvelope struct {
	Result  RadarEmailRoutingTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupDKIMResponseEnvelope]
type radarEmailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupDMARCParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupDMARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingTimeseriesGroupDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupDMARCParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval15m RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1w"
)

type RadarEmailRoutingTimeseriesGroupDMARCParamsARC string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsARCPass RadarEmailRoutingTimeseriesGroupDMARCParamsARC = "PASS"
	RadarEmailRoutingTimeseriesGroupDMARCParamsARCNone RadarEmailRoutingTimeseriesGroupDMARCParamsARC = "NONE"
	RadarEmailRoutingTimeseriesGroupDMARCParamsARCFail RadarEmailRoutingTimeseriesGroupDMARCParamsARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange1d         RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "1d"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange2d         RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "2d"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange7d         RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "7d"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange14d        RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "14d"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange28d        RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "28d"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange12w        RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "12w"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange24w        RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "24w"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange52w        RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "52w"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange1dControl  RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "1dControl"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange2dControl  RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "2dControl"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange7dControl  RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "7dControl"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange14dControl RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "14dControl"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange28dControl RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "28dControl"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange12wControl RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "12wControl"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange24wControl RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange = "24wControl"
)

type RadarEmailRoutingTimeseriesGroupDMARCParamsDKIM string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsDKIMPass RadarEmailRoutingTimeseriesGroupDMARCParamsDKIM = "PASS"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDKIMNone RadarEmailRoutingTimeseriesGroupDMARCParamsDKIM = "NONE"
	RadarEmailRoutingTimeseriesGroupDMARCParamsDKIMFail RadarEmailRoutingTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupDMARCParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupDMARCParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingTimeseriesGroupDMARCParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsFormatJson RadarEmailRoutingTimeseriesGroupDMARCParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupDMARCParamsFormatCsv  RadarEmailRoutingTimeseriesGroupDMARCParamsFormat = "CSV"
)

type RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv6"
)

type RadarEmailRoutingTimeseriesGroupDMARCParamsSPF string

const (
	RadarEmailRoutingTimeseriesGroupDMARCParamsSPFPass RadarEmailRoutingTimeseriesGroupDMARCParamsSPF = "PASS"
	RadarEmailRoutingTimeseriesGroupDMARCParamsSPFNone RadarEmailRoutingTimeseriesGroupDMARCParamsSPF = "NONE"
	RadarEmailRoutingTimeseriesGroupDMARCParamsSPFFail RadarEmailRoutingTimeseriesGroupDMARCParamsSPF = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupDMARCResponseEnvelope struct {
	Result  RadarEmailRoutingTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupDMARCResponseEnvelope]
type radarEmailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupEncryptedParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingTimeseriesGroupEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingTimeseriesGroupEncryptedParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingTimeseriesGroupEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupEncryptedParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1w"
)

type RadarEmailRoutingTimeseriesGroupEncryptedParamsARC string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsARCPass RadarEmailRoutingTimeseriesGroupEncryptedParamsARC = "PASS"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsARCNone RadarEmailRoutingTimeseriesGroupEncryptedParamsARC = "NONE"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsARCFail RadarEmailRoutingTimeseriesGroupEncryptedParamsARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange1d         RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "1d"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange2d         RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "2d"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange7d         RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "7d"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange14d        RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "14d"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange28d        RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "28d"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange12w        RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "12w"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange24w        RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "24w"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange52w        RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "52w"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange1dControl  RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "1dControl"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange2dControl  RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "2dControl"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange7dControl  RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "7dControl"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange14dControl RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "14dControl"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange28dControl RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "28dControl"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange12wControl RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "12wControl"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange24wControl RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange = "24wControl"
)

type RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIM string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIMPass RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIM = "PASS"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIMNone RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIM = "NONE"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIMFail RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIM = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARC string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARCPass RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARC = "PASS"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARCNone RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARC = "NONE"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARCFail RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailRoutingTimeseriesGroupEncryptedParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsFormatJson RadarEmailRoutingTimeseriesGroupEncryptedParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsFormatCsv  RadarEmailRoutingTimeseriesGroupEncryptedParamsFormat = "CSV"
)

type RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv6"
)

type RadarEmailRoutingTimeseriesGroupEncryptedParamsSPF string

const (
	RadarEmailRoutingTimeseriesGroupEncryptedParamsSPFPass RadarEmailRoutingTimeseriesGroupEncryptedParamsSPF = "PASS"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsSPFNone RadarEmailRoutingTimeseriesGroupEncryptedParamsSPF = "NONE"
	RadarEmailRoutingTimeseriesGroupEncryptedParamsSPFFail RadarEmailRoutingTimeseriesGroupEncryptedParamsSPF = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupEncryptedResponseEnvelope struct {
	Result  RadarEmailRoutingTimeseriesGroupEncryptedResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupEncryptedResponseEnvelope]
type radarEmailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingTimeseriesGroupIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupIPVersionParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingTimeseriesGroupIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type RadarEmailRoutingTimeseriesGroupIPVersionParamsARC string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsARCPass RadarEmailRoutingTimeseriesGroupIPVersionParamsARC = "PASS"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsARCNone RadarEmailRoutingTimeseriesGroupIPVersionParamsARC = "NONE"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsARCFail RadarEmailRoutingTimeseriesGroupIPVersionParamsARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange1d         RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "1d"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange2d         RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "2d"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange7d         RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "7d"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange14d        RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "14d"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange28d        RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "28d"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange12w        RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "12w"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange24w        RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "24w"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange52w        RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "52w"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange1dControl  RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "1dControl"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange2dControl  RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "2dControl"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange7dControl  RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "7dControl"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange14dControl RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "14dControl"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange28dControl RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "28dControl"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange12wControl RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "12wControl"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange24wControl RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

type RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIM string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIMPass RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIM = "PASS"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIMNone RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIM = "NONE"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIMFail RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIM = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARC string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARCPass RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARC = "PASS"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARCNone RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARC = "NONE"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARCFail RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupIPVersionParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingTimeseriesGroupIPVersionParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsFormatJson RadarEmailRoutingTimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsFormatCsv  RadarEmailRoutingTimeseriesGroupIPVersionParamsFormat = "CSV"
)

type RadarEmailRoutingTimeseriesGroupIPVersionParamsSPF string

const (
	RadarEmailRoutingTimeseriesGroupIPVersionParamsSPFPass RadarEmailRoutingTimeseriesGroupIPVersionParamsSPF = "PASS"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsSPFNone RadarEmailRoutingTimeseriesGroupIPVersionParamsSPF = "NONE"
	RadarEmailRoutingTimeseriesGroupIPVersionParamsSPFFail RadarEmailRoutingTimeseriesGroupIPVersionParamsSPF = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  RadarEmailRoutingTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupIPVersionResponseEnvelope]
type radarEmailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingTimeseriesGroupSPFParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingTimeseriesGroupSPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingTimeseriesGroupSPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupSPFParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupSPFParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupSPFParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval15m RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval = "1w"
)

type RadarEmailRoutingTimeseriesGroupSPFParamsARC string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsARCPass RadarEmailRoutingTimeseriesGroupSPFParamsARC = "PASS"
	RadarEmailRoutingTimeseriesGroupSPFParamsARCNone RadarEmailRoutingTimeseriesGroupSPFParamsARC = "NONE"
	RadarEmailRoutingTimeseriesGroupSPFParamsARCFail RadarEmailRoutingTimeseriesGroupSPFParamsARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupSPFParamsDateRange string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange1d         RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "1d"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange2d         RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "2d"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange7d         RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "7d"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange14d        RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "14d"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange28d        RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "28d"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange12w        RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "12w"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange24w        RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "24w"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange52w        RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "52w"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange1dControl  RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "1dControl"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange2dControl  RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "2dControl"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange7dControl  RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "7dControl"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange14dControl RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "14dControl"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange28dControl RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "28dControl"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange12wControl RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "12wControl"
	RadarEmailRoutingTimeseriesGroupSPFParamsDateRange24wControl RadarEmailRoutingTimeseriesGroupSPFParamsDateRange = "24wControl"
)

type RadarEmailRoutingTimeseriesGroupSPFParamsDKIM string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsDKIMPass RadarEmailRoutingTimeseriesGroupSPFParamsDKIM = "PASS"
	RadarEmailRoutingTimeseriesGroupSPFParamsDKIMNone RadarEmailRoutingTimeseriesGroupSPFParamsDKIM = "NONE"
	RadarEmailRoutingTimeseriesGroupSPFParamsDKIMFail RadarEmailRoutingTimeseriesGroupSPFParamsDKIM = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupSPFParamsDMARC string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsDMARCPass RadarEmailRoutingTimeseriesGroupSPFParamsDMARC = "PASS"
	RadarEmailRoutingTimeseriesGroupSPFParamsDMARCNone RadarEmailRoutingTimeseriesGroupSPFParamsDMARC = "NONE"
	RadarEmailRoutingTimeseriesGroupSPFParamsDMARCFail RadarEmailRoutingTimeseriesGroupSPFParamsDMARC = "FAIL"
)

type RadarEmailRoutingTimeseriesGroupSPFParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupSPFParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupSPFParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingTimeseriesGroupSPFParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsFormatJson RadarEmailRoutingTimeseriesGroupSPFParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupSPFParamsFormatCsv  RadarEmailRoutingTimeseriesGroupSPFParamsFormat = "CSV"
)

type RadarEmailRoutingTimeseriesGroupSPFParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv6"
)

type RadarEmailRoutingTimeseriesGroupSPFResponseEnvelope struct {
	Result  RadarEmailRoutingTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupSPFResponseEnvelope]
type radarEmailRoutingTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
