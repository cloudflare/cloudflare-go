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

// RadarBGPHijackService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPHijackService] method
// instead.
type RadarBGPHijackService struct {
	Options []option.RequestOption
}

// NewRadarBGPHijackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPHijackService(opts ...option.RequestOption) (r *RadarBGPHijackService) {
	r = &RadarBGPHijackService{}
	r.Options = opts
	return
}

// Get the BGP hijack events. (Beta)
func (r *RadarBGPHijackService) Events(ctx context.Context, query RadarBGPHijackEventsParams, opts ...option.RequestOption) (res *RadarBGPHijackEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPHijackEventsResponseEnvelope
	path := "radar/bgp/hijacks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPHijackEventsResponse struct {
	AsnInfo       []RadarBGPHijackEventsResponseAsnInfo `json:"asn_info,required"`
	Events        []RadarBGPHijackEventsResponseEvent   `json:"events,required"`
	TotalMonitors int64                                 `json:"total_monitors,required"`
	JSON          radarBGPHijackEventsResponseJSON      `json:"-"`
}

// radarBGPHijackEventsResponseJSON contains the JSON metadata for the struct
// [RadarBGPHijackEventsResponse]
type radarBGPHijackEventsResponseJSON struct {
	AsnInfo       apijson.Field
	Events        apijson.Field
	TotalMonitors apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBGPHijackEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackEventsResponseAsnInfo struct {
	Asn         int64                                   `json:"asn,required"`
	CountryCode string                                  `json:"country_code,required"`
	OrgName     string                                  `json:"org_name,required"`
	JSON        radarBGPHijackEventsResponseAsnInfoJSON `json:"-"`
}

// radarBGPHijackEventsResponseAsnInfoJSON contains the JSON metadata for the
// struct [RadarBGPHijackEventsResponseAsnInfo]
type radarBGPHijackEventsResponseAsnInfoJSON struct {
	Asn         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackEventsResponseAsnInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackEventsResponseEvent struct {
	ID              int64                                   `json:"id,required"`
	ConfidenceScore int64                                   `json:"confidence_score,required"`
	Duration        int64                                   `json:"duration,required"`
	EventType       int64                                   `json:"event_type,required"`
	HijackMsgsCount int64                                   `json:"hijack_msgs_count,required"`
	HijackerAsn     int64                                   `json:"hijacker_asn,required"`
	HijackerCountry string                                  `json:"hijacker_country,required"`
	IsStale         bool                                    `json:"is_stale,required"`
	MaxHijackTs     string                                  `json:"max_hijack_ts,required"`
	MaxMsgTs        string                                  `json:"max_msg_ts,required"`
	MinHijackTs     string                                  `json:"min_hijack_ts,required"`
	OnGoingCount    int64                                   `json:"on_going_count,required"`
	PeerAsns        []int64                                 `json:"peer_asns,required"`
	PeerIPCount     int64                                   `json:"peer_ip_count,required"`
	Prefixes        []string                                `json:"prefixes,required"`
	Tags            []RadarBGPHijackEventsResponseEventsTag `json:"tags,required"`
	VictimAsns      []int64                                 `json:"victim_asns,required"`
	VictimCountries []string                                `json:"victim_countries,required"`
	JSON            radarBGPHijackEventsResponseEventJSON   `json:"-"`
}

// radarBGPHijackEventsResponseEventJSON contains the JSON metadata for the struct
// [RadarBGPHijackEventsResponseEvent]
type radarBGPHijackEventsResponseEventJSON struct {
	ID              apijson.Field
	ConfidenceScore apijson.Field
	Duration        apijson.Field
	EventType       apijson.Field
	HijackMsgsCount apijson.Field
	HijackerAsn     apijson.Field
	HijackerCountry apijson.Field
	IsStale         apijson.Field
	MaxHijackTs     apijson.Field
	MaxMsgTs        apijson.Field
	MinHijackTs     apijson.Field
	OnGoingCount    apijson.Field
	PeerAsns        apijson.Field
	PeerIPCount     apijson.Field
	Prefixes        apijson.Field
	Tags            apijson.Field
	VictimAsns      apijson.Field
	VictimCountries apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarBGPHijackEventsResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackEventsResponseEventsTag struct {
	Name  string                                    `json:"name,required"`
	Score int64                                     `json:"score,required"`
	JSON  radarBGPHijackEventsResponseEventsTagJSON `json:"-"`
}

// radarBGPHijackEventsResponseEventsTagJSON contains the JSON metadata for the
// struct [RadarBGPHijackEventsResponseEventsTag]
type radarBGPHijackEventsResponseEventsTagJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackEventsResponseEventsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackEventsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBGPHijackEventsParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[RadarBGPHijackEventsParamsFormat] `query:"format"`
	// The potential hijacker AS of a BGP hijack event
	HijackerAsn param.Field[int64] `query:"hijackerAsn"`
	// The potential hijacker or victim AS of a BGP hijack event
	InvolvedAsn param.Field[int64] `query:"involvedAsn"`
	// The country code of the potential hijacker or victim AS of a BGP hijack event
	InvolvedCountry param.Field[string] `query:"involvedCountry"`
	// The maximum confidence score to filter events (1-4 low, 5-7 mid, 8+ high)
	MaxConfidence param.Field[int64] `query:"maxConfidence"`
	// The minimum confidence score to filter events (1-4 low, 5-7 mid, 8+ high)
	MinConfidence param.Field[int64] `query:"minConfidence"`
	// Current page number, starting from 1
	Page param.Field[int64] `query:"page"`
	// Number of entries per page
	PerPage param.Field[int64] `query:"per_page"`
	// The prefix hijacked during a BGP hijack event
	Prefix param.Field[string] `query:"prefix"`
	// Sort events by field
	SortBy param.Field[RadarBGPHijackEventsParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[RadarBGPHijackEventsParamsSortOrder] `query:"sortOrder"`
	// The potential victim AS of a BGP hijack event
	VictimAsn param.Field[int64] `query:"victimAsn"`
}

// URLQuery serializes [RadarBGPHijackEventsParams]'s query parameters as
// `url.Values`.
func (r RadarBGPHijackEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBGPHijackEventsParamsDateRange string

const (
	RadarBGPHijackEventsParamsDateRange1d         RadarBGPHijackEventsParamsDateRange = "1d"
	RadarBGPHijackEventsParamsDateRange2d         RadarBGPHijackEventsParamsDateRange = "2d"
	RadarBGPHijackEventsParamsDateRange7d         RadarBGPHijackEventsParamsDateRange = "7d"
	RadarBGPHijackEventsParamsDateRange14d        RadarBGPHijackEventsParamsDateRange = "14d"
	RadarBGPHijackEventsParamsDateRange28d        RadarBGPHijackEventsParamsDateRange = "28d"
	RadarBGPHijackEventsParamsDateRange12w        RadarBGPHijackEventsParamsDateRange = "12w"
	RadarBGPHijackEventsParamsDateRange24w        RadarBGPHijackEventsParamsDateRange = "24w"
	RadarBGPHijackEventsParamsDateRange52w        RadarBGPHijackEventsParamsDateRange = "52w"
	RadarBGPHijackEventsParamsDateRange1dControl  RadarBGPHijackEventsParamsDateRange = "1dControl"
	RadarBGPHijackEventsParamsDateRange2dControl  RadarBGPHijackEventsParamsDateRange = "2dControl"
	RadarBGPHijackEventsParamsDateRange7dControl  RadarBGPHijackEventsParamsDateRange = "7dControl"
	RadarBGPHijackEventsParamsDateRange14dControl RadarBGPHijackEventsParamsDateRange = "14dControl"
	RadarBGPHijackEventsParamsDateRange28dControl RadarBGPHijackEventsParamsDateRange = "28dControl"
	RadarBGPHijackEventsParamsDateRange12wControl RadarBGPHijackEventsParamsDateRange = "12wControl"
	RadarBGPHijackEventsParamsDateRange24wControl RadarBGPHijackEventsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPHijackEventsParamsFormat string

const (
	RadarBGPHijackEventsParamsFormatJson RadarBGPHijackEventsParamsFormat = "JSON"
	RadarBGPHijackEventsParamsFormatCsv  RadarBGPHijackEventsParamsFormat = "CSV"
)

// Sort events by field
type RadarBGPHijackEventsParamsSortBy string

const (
	RadarBGPHijackEventsParamsSortByID         RadarBGPHijackEventsParamsSortBy = "ID"
	RadarBGPHijackEventsParamsSortByTime       RadarBGPHijackEventsParamsSortBy = "TIME"
	RadarBGPHijackEventsParamsSortByConfidence RadarBGPHijackEventsParamsSortBy = "CONFIDENCE"
)

// Sort order
type RadarBGPHijackEventsParamsSortOrder string

const (
	RadarBGPHijackEventsParamsSortOrderAsc  RadarBGPHijackEventsParamsSortOrder = "ASC"
	RadarBGPHijackEventsParamsSortOrderDesc RadarBGPHijackEventsParamsSortOrder = "DESC"
)

type RadarBGPHijackEventsResponseEnvelope struct {
	Result     RadarBGPHijackEventsResponse                   `json:"result,required"`
	ResultInfo RadarBGPHijackEventsResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                           `json:"success,required"`
	JSON       radarBGPHijackEventsResponseEnvelopeJSON       `json:"-"`
}

// radarBGPHijackEventsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPHijackEventsResponseEnvelope]
type radarBGPHijackEventsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackEventsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackEventsResponseEnvelopeResultInfo struct {
	Count      int64                                              `json:"count,required"`
	Page       int64                                              `json:"page,required"`
	PerPage    int64                                              `json:"per_page,required"`
	TotalCount int64                                              `json:"total_count,required"`
	JSON       radarBGPHijackEventsResponseEnvelopeResultInfoJSON `json:"-"`
}

// radarBGPHijackEventsResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [RadarBGPHijackEventsResponseEnvelopeResultInfo]
type radarBGPHijackEventsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackEventsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
