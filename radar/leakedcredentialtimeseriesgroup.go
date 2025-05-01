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

// LeakedCredentialTimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialTimeseriesGroupService] method instead.
type LeakedCredentialTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewLeakedCredentialTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLeakedCredentialTimeseriesGroupService(opts ...option.RequestOption) (r *LeakedCredentialTimeseriesGroupService) {
	r = &LeakedCredentialTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of HTTP authentication requests by bot class over
// time.
func (r *LeakedCredentialTimeseriesGroupService) BotClass(ctx context.Context, query LeakedCredentialTimeseriesGroupBotClassParams, opts ...option.RequestOption) (res *LeakedCredentialTimeseriesGroupBotClassResponse, err error) {
	var env LeakedCredentialTimeseriesGroupBotClassResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP authentication requests by compromised
// credential status over time.
func (r *LeakedCredentialTimeseriesGroupService) Compromised(ctx context.Context, query LeakedCredentialTimeseriesGroupCompromisedParams, opts ...option.RequestOption) (res *LeakedCredentialTimeseriesGroupCompromisedResponse, err error) {
	var env LeakedCredentialTimeseriesGroupCompromisedResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/timeseries_groups/compromised"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LeakedCredentialTimeseriesGroupBotClassResponse struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 LeakedCredentialTimeseriesGroupBotClassResponseSerie0 `json:"serie_0,required"`
	JSON   leakedCredentialTimeseriesGroupBotClassResponseJSON   `json:"-"`
}

// leakedCredentialTimeseriesGroupBotClassResponseJSON contains the JSON metadata
// for the struct [LeakedCredentialTimeseriesGroupBotClassResponse]
type leakedCredentialTimeseriesGroupBotClassResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupBotClassResponseSerie0 struct {
	Bot        []string                                                  `json:"bot,required"`
	Human      []string                                                  `json:"human,required"`
	Timestamps []string                                                  `json:"timestamps,required"`
	JSON       leakedCredentialTimeseriesGroupBotClassResponseSerie0JSON `json:"-"`
}

// leakedCredentialTimeseriesGroupBotClassResponseSerie0JSON contains the JSON
// metadata for the struct [LeakedCredentialTimeseriesGroupBotClassResponseSerie0]
type leakedCredentialTimeseriesGroupBotClassResponseSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupBotClassResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupBotClassResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupCompromisedResponse struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 LeakedCredentialTimeseriesGroupCompromisedResponseSerie0 `json:"serie_0,required"`
	JSON   leakedCredentialTimeseriesGroupCompromisedResponseJSON   `json:"-"`
}

// leakedCredentialTimeseriesGroupCompromisedResponseJSON contains the JSON
// metadata for the struct [LeakedCredentialTimeseriesGroupCompromisedResponse]
type leakedCredentialTimeseriesGroupCompromisedResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupCompromisedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupCompromisedResponseJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupCompromisedResponseSerie0 struct {
	Clean       []string                                                     `json:"CLEAN,required"`
	Compromised []string                                                     `json:"COMPROMISED,required"`
	Timestamps  []string                                                     `json:"timestamps,required"`
	JSON        leakedCredentialTimeseriesGroupCompromisedResponseSerie0JSON `json:"-"`
}

// leakedCredentialTimeseriesGroupCompromisedResponseSerie0JSON contains the JSON
// metadata for the struct
// [LeakedCredentialTimeseriesGroupCompromisedResponseSerie0]
type leakedCredentialTimeseriesGroupCompromisedResponseSerie0JSON struct {
	Clean       apijson.Field
	Compromised apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupCompromisedResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupCompromisedResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupBotClassParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[LeakedCredentialTimeseriesGroupBotClassParamsAggInterval] `query:"aggInterval"`
	// Filters results by compromised credential status (clean vs. compromised).
	Compromised param.Field[[]LeakedCredentialTimeseriesGroupBotClassParamsCompromised] `query:"compromised"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[LeakedCredentialTimeseriesGroupBotClassParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [LeakedCredentialTimeseriesGroupBotClassParams]'s query
// parameters as `url.Values`.
func (r LeakedCredentialTimeseriesGroupBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type LeakedCredentialTimeseriesGroupBotClassParamsAggInterval string

const (
	LeakedCredentialTimeseriesGroupBotClassParamsAggInterval15m LeakedCredentialTimeseriesGroupBotClassParamsAggInterval = "15m"
	LeakedCredentialTimeseriesGroupBotClassParamsAggInterval1h  LeakedCredentialTimeseriesGroupBotClassParamsAggInterval = "1h"
	LeakedCredentialTimeseriesGroupBotClassParamsAggInterval1d  LeakedCredentialTimeseriesGroupBotClassParamsAggInterval = "1d"
	LeakedCredentialTimeseriesGroupBotClassParamsAggInterval1w  LeakedCredentialTimeseriesGroupBotClassParamsAggInterval = "1w"
)

func (r LeakedCredentialTimeseriesGroupBotClassParamsAggInterval) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupBotClassParamsAggInterval15m, LeakedCredentialTimeseriesGroupBotClassParamsAggInterval1h, LeakedCredentialTimeseriesGroupBotClassParamsAggInterval1d, LeakedCredentialTimeseriesGroupBotClassParamsAggInterval1w:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupBotClassParamsCompromised string

const (
	LeakedCredentialTimeseriesGroupBotClassParamsCompromisedClean       LeakedCredentialTimeseriesGroupBotClassParamsCompromised = "CLEAN"
	LeakedCredentialTimeseriesGroupBotClassParamsCompromisedCompromised LeakedCredentialTimeseriesGroupBotClassParamsCompromised = "COMPROMISED"
)

func (r LeakedCredentialTimeseriesGroupBotClassParamsCompromised) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupBotClassParamsCompromisedClean, LeakedCredentialTimeseriesGroupBotClassParamsCompromisedCompromised:
		return true
	}
	return false
}

// Format in which results will be returned.
type LeakedCredentialTimeseriesGroupBotClassParamsFormat string

const (
	LeakedCredentialTimeseriesGroupBotClassParamsFormatJson LeakedCredentialTimeseriesGroupBotClassParamsFormat = "JSON"
	LeakedCredentialTimeseriesGroupBotClassParamsFormatCsv  LeakedCredentialTimeseriesGroupBotClassParamsFormat = "CSV"
)

func (r LeakedCredentialTimeseriesGroupBotClassParamsFormat) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupBotClassParamsFormatJson, LeakedCredentialTimeseriesGroupBotClassParamsFormatCsv:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupBotClassResponseEnvelope struct {
	Result  LeakedCredentialTimeseriesGroupBotClassResponse             `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    leakedCredentialTimeseriesGroupBotClassResponseEnvelopeJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupBotClassResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [LeakedCredentialTimeseriesGroupBotClassResponseEnvelope]
type leakedCredentialTimeseriesGroupBotClassResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupBotClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupBotClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialTimeseriesGroupCompromisedParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval] `query:"aggInterval"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]LeakedCredentialTimeseriesGroupCompromisedParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[LeakedCredentialTimeseriesGroupCompromisedParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [LeakedCredentialTimeseriesGroupCompromisedParams]'s query
// parameters as `url.Values`.
func (r LeakedCredentialTimeseriesGroupCompromisedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval string

const (
	LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval15m LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval = "15m"
	LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval1h  LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval = "1h"
	LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval1d  LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval = "1d"
	LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval1w  LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval = "1w"
)

func (r LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval15m, LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval1h, LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval1d, LeakedCredentialTimeseriesGroupCompromisedParamsAggInterval1w:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupCompromisedParamsBotClass string

const (
	LeakedCredentialTimeseriesGroupCompromisedParamsBotClassLikelyAutomated LeakedCredentialTimeseriesGroupCompromisedParamsBotClass = "LIKELY_AUTOMATED"
	LeakedCredentialTimeseriesGroupCompromisedParamsBotClassLikelyHuman     LeakedCredentialTimeseriesGroupCompromisedParamsBotClass = "LIKELY_HUMAN"
)

func (r LeakedCredentialTimeseriesGroupCompromisedParamsBotClass) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupCompromisedParamsBotClassLikelyAutomated, LeakedCredentialTimeseriesGroupCompromisedParamsBotClassLikelyHuman:
		return true
	}
	return false
}

// Format in which results will be returned.
type LeakedCredentialTimeseriesGroupCompromisedParamsFormat string

const (
	LeakedCredentialTimeseriesGroupCompromisedParamsFormatJson LeakedCredentialTimeseriesGroupCompromisedParamsFormat = "JSON"
	LeakedCredentialTimeseriesGroupCompromisedParamsFormatCsv  LeakedCredentialTimeseriesGroupCompromisedParamsFormat = "CSV"
)

func (r LeakedCredentialTimeseriesGroupCompromisedParamsFormat) IsKnown() bool {
	switch r {
	case LeakedCredentialTimeseriesGroupCompromisedParamsFormatJson, LeakedCredentialTimeseriesGroupCompromisedParamsFormatCsv:
		return true
	}
	return false
}

type LeakedCredentialTimeseriesGroupCompromisedResponseEnvelope struct {
	Result  LeakedCredentialTimeseriesGroupCompromisedResponse             `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    leakedCredentialTimeseriesGroupCompromisedResponseEnvelopeJSON `json:"-"`
}

// leakedCredentialTimeseriesGroupCompromisedResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [LeakedCredentialTimeseriesGroupCompromisedResponseEnvelope]
type leakedCredentialTimeseriesGroupCompromisedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialTimeseriesGroupCompromisedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialTimeseriesGroupCompromisedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
