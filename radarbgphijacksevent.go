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

// RadarBgpHijacksEventService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpHijacksEventService]
// method instead.
type RadarBgpHijacksEventService struct {
	Options []option.RequestOption
}

// NewRadarBgpHijacksEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpHijacksEventService(opts ...option.RequestOption) (r *RadarBgpHijacksEventService) {
	r = &RadarBgpHijacksEventService{}
	r.Options = opts
	return
}

// Get the BGP hijack events. (Beta)
func (r *RadarBgpHijacksEventService) List(ctx context.Context, query RadarBgpHijacksEventListParams, opts ...option.RequestOption) (res *RadarBgpHijacksEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/hijacks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpHijacksEventListResponse struct {
	Result     RadarBgpHijacksEventListResponseResult     `json:"result,required"`
	ResultInfo RadarBgpHijacksEventListResponseResultInfo `json:"result_info,required"`
	Success    bool                                       `json:"success,required"`
	JSON       radarBgpHijacksEventListResponseJSON       `json:"-"`
}

// radarBgpHijacksEventListResponseJSON contains the JSON metadata for the struct
// [RadarBgpHijacksEventListResponse]
type radarBgpHijacksEventListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijacksEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpHijacksEventListResponseResult struct {
	ASNInfo       []RadarBgpHijacksEventListResponseResultASNInfo `json:"asn_info,required"`
	Events        []RadarBgpHijacksEventListResponseResultEvent   `json:"events,required"`
	TotalMonitors int64                                           `json:"total_monitors,required"`
	JSON          radarBgpHijacksEventListResponseResultJSON      `json:"-"`
}

// radarBgpHijacksEventListResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpHijacksEventListResponseResult]
type radarBgpHijacksEventListResponseResultJSON struct {
	ASNInfo       apijson.Field
	Events        apijson.Field
	TotalMonitors apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBgpHijacksEventListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpHijacksEventListResponseResultASNInfo struct {
	ASN         int64                                             `json:"asn,required"`
	CountryCode string                                            `json:"country_code,required"`
	OrgName     string                                            `json:"org_name,required"`
	JSON        radarBgpHijacksEventListResponseResultASNInfoJSON `json:"-"`
}

// radarBgpHijacksEventListResponseResultASNInfoJSON contains the JSON metadata for
// the struct [RadarBgpHijacksEventListResponseResultASNInfo]
type radarBgpHijacksEventListResponseResultASNInfoJSON struct {
	ASN         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijacksEventListResponseResultASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpHijacksEventListResponseResultEvent struct {
	ID              int64                                             `json:"id,required"`
	ConfidenceScore int64                                             `json:"confidence_score,required"`
	Duration        int64                                             `json:"duration,required"`
	EventType       int64                                             `json:"event_type,required"`
	HijackMsgsCount int64                                             `json:"hijack_msgs_count,required"`
	HijackerASN     int64                                             `json:"hijacker_asn,required"`
	HijackerCountry string                                            `json:"hijacker_country,required"`
	IsStale         bool                                              `json:"is_stale,required"`
	MaxHijackTs     string                                            `json:"max_hijack_ts,required"`
	MaxMsgTs        string                                            `json:"max_msg_ts,required"`
	MinHijackTs     string                                            `json:"min_hijack_ts,required"`
	OnGoingCount    int64                                             `json:"on_going_count,required"`
	PeerASNs        []int64                                           `json:"peer_asns,required"`
	PeerIPCount     int64                                             `json:"peer_ip_count,required"`
	Prefixes        []string                                          `json:"prefixes,required"`
	Tags            []RadarBgpHijacksEventListResponseResultEventsTag `json:"tags,required"`
	VictimASNs      []int64                                           `json:"victim_asns,required"`
	VictimCountries []string                                          `json:"victim_countries,required"`
	JSON            radarBgpHijacksEventListResponseResultEventJSON   `json:"-"`
}

// radarBgpHijacksEventListResponseResultEventJSON contains the JSON metadata for
// the struct [RadarBgpHijacksEventListResponseResultEvent]
type radarBgpHijacksEventListResponseResultEventJSON struct {
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

func (r *RadarBgpHijacksEventListResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpHijacksEventListResponseResultEventsTag struct {
	Name  string                                              `json:"name,required"`
	Score int64                                               `json:"score,required"`
	JSON  radarBgpHijacksEventListResponseResultEventsTagJSON `json:"-"`
}

// radarBgpHijacksEventListResponseResultEventsTagJSON contains the JSON metadata
// for the struct [RadarBgpHijacksEventListResponseResultEventsTag]
type radarBgpHijacksEventListResponseResultEventsTagJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijacksEventListResponseResultEventsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpHijacksEventListResponseResultInfo struct {
	Count      int64                                          `json:"count,required"`
	Page       int64                                          `json:"page,required"`
	PerPage    int64                                          `json:"per_page,required"`
	TotalCount int64                                          `json:"total_count,required"`
	JSON       radarBgpHijacksEventListResponseResultInfoJSON `json:"-"`
}

// radarBgpHijacksEventListResponseResultInfoJSON contains the JSON metadata for
// the struct [RadarBgpHijacksEventListResponseResultInfo]
type radarBgpHijacksEventListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijacksEventListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpHijacksEventListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBgpHijacksEventListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[RadarBgpHijacksEventListParamsFormat] `query:"format"`
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
	SortBy param.Field[RadarBgpHijacksEventListParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[RadarBgpHijacksEventListParamsSortOrder] `query:"sortOrder"`
	// The potential victim AS of a BGP hijack event
	VictimASN param.Field[int64] `query:"victimAsn"`
}

// URLQuery serializes [RadarBgpHijacksEventListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpHijacksEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBgpHijacksEventListParamsDateRange string

const (
	RadarBgpHijacksEventListParamsDateRange1d         RadarBgpHijacksEventListParamsDateRange = "1d"
	RadarBgpHijacksEventListParamsDateRange2d         RadarBgpHijacksEventListParamsDateRange = "2d"
	RadarBgpHijacksEventListParamsDateRange7d         RadarBgpHijacksEventListParamsDateRange = "7d"
	RadarBgpHijacksEventListParamsDateRange14d        RadarBgpHijacksEventListParamsDateRange = "14d"
	RadarBgpHijacksEventListParamsDateRange28d        RadarBgpHijacksEventListParamsDateRange = "28d"
	RadarBgpHijacksEventListParamsDateRange12w        RadarBgpHijacksEventListParamsDateRange = "12w"
	RadarBgpHijacksEventListParamsDateRange24w        RadarBgpHijacksEventListParamsDateRange = "24w"
	RadarBgpHijacksEventListParamsDateRange52w        RadarBgpHijacksEventListParamsDateRange = "52w"
	RadarBgpHijacksEventListParamsDateRange1dControl  RadarBgpHijacksEventListParamsDateRange = "1dControl"
	RadarBgpHijacksEventListParamsDateRange2dControl  RadarBgpHijacksEventListParamsDateRange = "2dControl"
	RadarBgpHijacksEventListParamsDateRange7dControl  RadarBgpHijacksEventListParamsDateRange = "7dControl"
	RadarBgpHijacksEventListParamsDateRange14dControl RadarBgpHijacksEventListParamsDateRange = "14dControl"
	RadarBgpHijacksEventListParamsDateRange28dControl RadarBgpHijacksEventListParamsDateRange = "28dControl"
	RadarBgpHijacksEventListParamsDateRange12wControl RadarBgpHijacksEventListParamsDateRange = "12wControl"
	RadarBgpHijacksEventListParamsDateRange24wControl RadarBgpHijacksEventListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBgpHijacksEventListParamsFormat string

const (
	RadarBgpHijacksEventListParamsFormatJson RadarBgpHijacksEventListParamsFormat = "JSON"
	RadarBgpHijacksEventListParamsFormatCsv  RadarBgpHijacksEventListParamsFormat = "CSV"
)

// Sort events by field
type RadarBgpHijacksEventListParamsSortBy string

const (
	RadarBgpHijacksEventListParamsSortByID         RadarBgpHijacksEventListParamsSortBy = "ID"
	RadarBgpHijacksEventListParamsSortByTime       RadarBgpHijacksEventListParamsSortBy = "TIME"
	RadarBgpHijacksEventListParamsSortByConfidence RadarBgpHijacksEventListParamsSortBy = "CONFIDENCE"
)

// Sort order
type RadarBgpHijacksEventListParamsSortOrder string

const (
	RadarBgpHijacksEventListParamsSortOrderAsc  RadarBgpHijacksEventListParamsSortOrder = "ASC"
	RadarBgpHijacksEventListParamsSortOrderDesc RadarBgpHijacksEventListParamsSortOrder = "DESC"
)
