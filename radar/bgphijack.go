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

// BGPHijackService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewBGPHijackService] method instead.
type BGPHijackService struct {
	Options []option.RequestOption
}

// NewBGPHijackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPHijackService(opts ...option.RequestOption) (r *BGPHijackService) {
	r = &BGPHijackService{}
	r.Options = opts
	return
}

// Get the BGP hijack events. (Beta)
func (r *BGPHijackService) Events(ctx context.Context, query BGPHijackEventsParams, opts ...option.RequestOption) (res *BGPHijackEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPHijackEventsResponseEnvelope
	path := "radar/bgp/hijacks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPHijackEventsResponse struct {
	ASNInfo       []BGPHijackEventsResponseASNInfo `json:"asn_info,required"`
	Events        []BGPHijackEventsResponseEvent   `json:"events,required"`
	TotalMonitors int64                            `json:"total_monitors,required"`
	JSON          bgpHijackEventsResponseJSON      `json:"-"`
}

// bgpHijackEventsResponseJSON contains the JSON metadata for the struct
// [BGPHijackEventsResponse]
type bgpHijackEventsResponseJSON struct {
	ASNInfo       apijson.Field
	Events        apijson.Field
	TotalMonitors apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BGPHijackEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventsResponseJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventsResponseASNInfo struct {
	ASN         int64                              `json:"asn,required"`
	CountryCode string                             `json:"country_code,required"`
	OrgName     string                             `json:"org_name,required"`
	JSON        bgpHijackEventsResponseASNInfoJSON `json:"-"`
}

// bgpHijackEventsResponseASNInfoJSON contains the JSON metadata for the struct
// [BGPHijackEventsResponseASNInfo]
type bgpHijackEventsResponseASNInfoJSON struct {
	ASN         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventsResponseASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventsResponseASNInfoJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventsResponseEvent struct {
	ID              int64                              `json:"id,required"`
	ConfidenceScore int64                              `json:"confidence_score,required"`
	Duration        int64                              `json:"duration,required"`
	EventType       int64                              `json:"event_type,required"`
	HijackMsgsCount int64                              `json:"hijack_msgs_count,required"`
	HijackerASN     int64                              `json:"hijacker_asn,required"`
	HijackerCountry string                             `json:"hijacker_country,required"`
	IsStale         bool                               `json:"is_stale,required"`
	MaxHijackTs     string                             `json:"max_hijack_ts,required"`
	MaxMsgTs        string                             `json:"max_msg_ts,required"`
	MinHijackTs     string                             `json:"min_hijack_ts,required"`
	OnGoingCount    int64                              `json:"on_going_count,required"`
	PeerASNs        []int64                            `json:"peer_asns,required"`
	PeerIPCount     int64                              `json:"peer_ip_count,required"`
	Prefixes        []string                           `json:"prefixes,required"`
	Tags            []BGPHijackEventsResponseEventsTag `json:"tags,required"`
	VictimASNs      []int64                            `json:"victim_asns,required"`
	VictimCountries []string                           `json:"victim_countries,required"`
	JSON            bgpHijackEventsResponseEventJSON   `json:"-"`
}

// bgpHijackEventsResponseEventJSON contains the JSON metadata for the struct
// [BGPHijackEventsResponseEvent]
type bgpHijackEventsResponseEventJSON struct {
	ID              apijson.Field
	ConfidenceScore apijson.Field
	Duration        apijson.Field
	EventType       apijson.Field
	HijackMsgsCount apijson.Field
	HijackerASN     apijson.Field
	HijackerCountry apijson.Field
	IsStale         apijson.Field
	MaxHijackTs     apijson.Field
	MaxMsgTs        apijson.Field
	MinHijackTs     apijson.Field
	OnGoingCount    apijson.Field
	PeerASNs        apijson.Field
	PeerIPCount     apijson.Field
	Prefixes        apijson.Field
	Tags            apijson.Field
	VictimASNs      apijson.Field
	VictimCountries apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BGPHijackEventsResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventsResponseEventJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventsResponseEventsTag struct {
	Name  string                               `json:"name,required"`
	Score int64                                `json:"score,required"`
	JSON  bgpHijackEventsResponseEventsTagJSON `json:"-"`
}

// bgpHijackEventsResponseEventsTagJSON contains the JSON metadata for the struct
// [BGPHijackEventsResponseEventsTag]
type bgpHijackEventsResponseEventsTagJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventsResponseEventsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventsResponseEventsTagJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[BGPHijackEventsParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[BGPHijackEventsParamsFormat] `query:"format"`
	// The potential hijacker AS of a BGP hijack event
	HijackerASN param.Field[int64] `query:"hijackerAsn"`
	// The potential hijacker or victim AS of a BGP hijack event
	InvolvedASN param.Field[int64] `query:"involvedAsn"`
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
	SortBy param.Field[BGPHijackEventsParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[BGPHijackEventsParamsSortOrder] `query:"sortOrder"`
	// The potential victim AS of a BGP hijack event
	VictimASN param.Field[int64] `query:"victimAsn"`
}

// URLQuery serializes [BGPHijackEventsParams]'s query parameters as `url.Values`.
func (r BGPHijackEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type BGPHijackEventsParamsDateRange string

const (
	BGPHijackEventsParamsDateRange1d         BGPHijackEventsParamsDateRange = "1d"
	BGPHijackEventsParamsDateRange2d         BGPHijackEventsParamsDateRange = "2d"
	BGPHijackEventsParamsDateRange7d         BGPHijackEventsParamsDateRange = "7d"
	BGPHijackEventsParamsDateRange14d        BGPHijackEventsParamsDateRange = "14d"
	BGPHijackEventsParamsDateRange28d        BGPHijackEventsParamsDateRange = "28d"
	BGPHijackEventsParamsDateRange12w        BGPHijackEventsParamsDateRange = "12w"
	BGPHijackEventsParamsDateRange24w        BGPHijackEventsParamsDateRange = "24w"
	BGPHijackEventsParamsDateRange52w        BGPHijackEventsParamsDateRange = "52w"
	BGPHijackEventsParamsDateRange1dControl  BGPHijackEventsParamsDateRange = "1dControl"
	BGPHijackEventsParamsDateRange2dControl  BGPHijackEventsParamsDateRange = "2dControl"
	BGPHijackEventsParamsDateRange7dControl  BGPHijackEventsParamsDateRange = "7dControl"
	BGPHijackEventsParamsDateRange14dControl BGPHijackEventsParamsDateRange = "14dControl"
	BGPHijackEventsParamsDateRange28dControl BGPHijackEventsParamsDateRange = "28dControl"
	BGPHijackEventsParamsDateRange12wControl BGPHijackEventsParamsDateRange = "12wControl"
	BGPHijackEventsParamsDateRange24wControl BGPHijackEventsParamsDateRange = "24wControl"
)

func (r BGPHijackEventsParamsDateRange) IsKnown() bool {
	switch r {
	case BGPHijackEventsParamsDateRange1d, BGPHijackEventsParamsDateRange2d, BGPHijackEventsParamsDateRange7d, BGPHijackEventsParamsDateRange14d, BGPHijackEventsParamsDateRange28d, BGPHijackEventsParamsDateRange12w, BGPHijackEventsParamsDateRange24w, BGPHijackEventsParamsDateRange52w, BGPHijackEventsParamsDateRange1dControl, BGPHijackEventsParamsDateRange2dControl, BGPHijackEventsParamsDateRange7dControl, BGPHijackEventsParamsDateRange14dControl, BGPHijackEventsParamsDateRange28dControl, BGPHijackEventsParamsDateRange12wControl, BGPHijackEventsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type BGPHijackEventsParamsFormat string

const (
	BGPHijackEventsParamsFormatJson BGPHijackEventsParamsFormat = "JSON"
	BGPHijackEventsParamsFormatCsv  BGPHijackEventsParamsFormat = "CSV"
)

func (r BGPHijackEventsParamsFormat) IsKnown() bool {
	switch r {
	case BGPHijackEventsParamsFormatJson, BGPHijackEventsParamsFormatCsv:
		return true
	}
	return false
}

// Sort events by field
type BGPHijackEventsParamsSortBy string

const (
	BGPHijackEventsParamsSortByID         BGPHijackEventsParamsSortBy = "ID"
	BGPHijackEventsParamsSortByTime       BGPHijackEventsParamsSortBy = "TIME"
	BGPHijackEventsParamsSortByConfidence BGPHijackEventsParamsSortBy = "CONFIDENCE"
)

func (r BGPHijackEventsParamsSortBy) IsKnown() bool {
	switch r {
	case BGPHijackEventsParamsSortByID, BGPHijackEventsParamsSortByTime, BGPHijackEventsParamsSortByConfidence:
		return true
	}
	return false
}

// Sort order
type BGPHijackEventsParamsSortOrder string

const (
	BGPHijackEventsParamsSortOrderAsc  BGPHijackEventsParamsSortOrder = "ASC"
	BGPHijackEventsParamsSortOrderDesc BGPHijackEventsParamsSortOrder = "DESC"
)

func (r BGPHijackEventsParamsSortOrder) IsKnown() bool {
	switch r {
	case BGPHijackEventsParamsSortOrderAsc, BGPHijackEventsParamsSortOrderDesc:
		return true
	}
	return false
}

type BGPHijackEventsResponseEnvelope struct {
	Result     BGPHijackEventsResponse                   `json:"result,required"`
	ResultInfo BGPHijackEventsResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                      `json:"success,required"`
	JSON       bgpHijackEventsResponseEnvelopeJSON       `json:"-"`
}

// bgpHijackEventsResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPHijackEventsResponseEnvelope]
type bgpHijackEventsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventsResponseEnvelopeResultInfo struct {
	Count      int64                                         `json:"count,required"`
	Page       int64                                         `json:"page,required"`
	PerPage    int64                                         `json:"per_page,required"`
	TotalCount int64                                         `json:"total_count,required"`
	JSON       bgpHijackEventsResponseEnvelopeResultInfoJSON `json:"-"`
}

// bgpHijackEventsResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [BGPHijackEventsResponseEnvelopeResultInfo]
type bgpHijackEventsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventsResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
