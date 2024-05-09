// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// BGPHijackEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPHijackEventService] method instead.
type BGPHijackEventService struct {
	Options []option.RequestOption
}

// NewBGPHijackEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPHijackEventService(opts ...option.RequestOption) (r *BGPHijackEventService) {
	r = &BGPHijackEventService{}
	r.Options = opts
	return
}

// Get the BGP hijack events. (Beta)
func (r *BGPHijackEventService) List(ctx context.Context, query BGPHijackEventListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[BGPHijackEventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "radar/bgp/hijacks/events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get the BGP hijack events. (Beta)
func (r *BGPHijackEventService) ListAutoPaging(ctx context.Context, query BGPHijackEventListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[BGPHijackEventListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, query, opts...))
}

type BGPHijackEventListResponse struct {
	Result     BGPHijackEventListResponseResult     `json:"result,required"`
	ResultInfo BGPHijackEventListResponseResultInfo `json:"result_info,required"`
	Success    bool                                 `json:"success,required"`
	JSON       bgpHijackEventListResponseJSON       `json:"-"`
}

// bgpHijackEventListResponseJSON contains the JSON metadata for the struct
// [BGPHijackEventListResponse]
type bgpHijackEventListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventListResponseJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventListResponseResult struct {
	ASNInfo       []BGPHijackEventListResponseResultASNInfo `json:"asn_info,required"`
	Events        []BGPHijackEventListResponseResultEvent   `json:"events,required"`
	TotalMonitors int64                                     `json:"total_monitors,required"`
	JSON          bgpHijackEventListResponseResultJSON      `json:"-"`
}

// bgpHijackEventListResponseResultJSON contains the JSON metadata for the struct
// [BGPHijackEventListResponseResult]
type bgpHijackEventListResponseResultJSON struct {
	ASNInfo       apijson.Field
	Events        apijson.Field
	TotalMonitors apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BGPHijackEventListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventListResponseResultJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventListResponseResultASNInfo struct {
	ASN         int64                                       `json:"asn,required"`
	CountryCode string                                      `json:"country_code,required"`
	OrgName     string                                      `json:"org_name,required"`
	JSON        bgpHijackEventListResponseResultASNInfoJSON `json:"-"`
}

// bgpHijackEventListResponseResultASNInfoJSON contains the JSON metadata for the
// struct [BGPHijackEventListResponseResultASNInfo]
type bgpHijackEventListResponseResultASNInfoJSON struct {
	ASN         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventListResponseResultASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventListResponseResultASNInfoJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventListResponseResultEvent struct {
	ID              int64                                       `json:"id,required"`
	ConfidenceScore int64                                       `json:"confidence_score,required"`
	Duration        int64                                       `json:"duration,required"`
	EventType       int64                                       `json:"event_type,required"`
	HijackMsgsCount int64                                       `json:"hijack_msgs_count,required"`
	HijackerASN     int64                                       `json:"hijacker_asn,required"`
	HijackerCountry string                                      `json:"hijacker_country,required"`
	IsStale         bool                                        `json:"is_stale,required"`
	MaxHijackTs     string                                      `json:"max_hijack_ts,required"`
	MaxMsgTs        string                                      `json:"max_msg_ts,required"`
	MinHijackTs     string                                      `json:"min_hijack_ts,required"`
	OnGoingCount    int64                                       `json:"on_going_count,required"`
	PeerASNs        []int64                                     `json:"peer_asns,required"`
	PeerIPCount     int64                                       `json:"peer_ip_count,required"`
	Prefixes        []string                                    `json:"prefixes,required"`
	Tags            []BGPHijackEventListResponseResultEventsTag `json:"tags,required"`
	VictimASNs      []int64                                     `json:"victim_asns,required"`
	VictimCountries []string                                    `json:"victim_countries,required"`
	JSON            bgpHijackEventListResponseResultEventJSON   `json:"-"`
}

// bgpHijackEventListResponseResultEventJSON contains the JSON metadata for the
// struct [BGPHijackEventListResponseResultEvent]
type bgpHijackEventListResponseResultEventJSON struct {
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

func (r *BGPHijackEventListResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventListResponseResultEventJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventListResponseResultEventsTag struct {
	Name  string                                        `json:"name,required"`
	Score int64                                         `json:"score,required"`
	JSON  bgpHijackEventListResponseResultEventsTagJSON `json:"-"`
}

// bgpHijackEventListResponseResultEventsTagJSON contains the JSON metadata for the
// struct [BGPHijackEventListResponseResultEventsTag]
type bgpHijackEventListResponseResultEventsTagJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventListResponseResultEventsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventListResponseResultEventsTagJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventListResponseResultInfo struct {
	Count      int64                                    `json:"count,required"`
	Page       int64                                    `json:"page,required"`
	PerPage    int64                                    `json:"per_page,required"`
	TotalCount int64                                    `json:"total_count,required"`
	JSON       bgpHijackEventListResponseResultInfoJSON `json:"-"`
}

// bgpHijackEventListResponseResultInfoJSON contains the JSON metadata for the
// struct [BGPHijackEventListResponseResultInfo]
type bgpHijackEventListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPHijackEventListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpHijackEventListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type BGPHijackEventListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[BGPHijackEventListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[BGPHijackEventListParamsFormat] `query:"format"`
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
	SortBy param.Field[BGPHijackEventListParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[BGPHijackEventListParamsSortOrder] `query:"sortOrder"`
	// The potential victim AS of a BGP hijack event
	VictimASN param.Field[int64] `query:"victimAsn"`
}

// URLQuery serializes [BGPHijackEventListParams]'s query parameters as
// `url.Values`.
func (r BGPHijackEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type BGPHijackEventListParamsDateRange string

const (
	BGPHijackEventListParamsDateRange1d         BGPHijackEventListParamsDateRange = "1d"
	BGPHijackEventListParamsDateRange2d         BGPHijackEventListParamsDateRange = "2d"
	BGPHijackEventListParamsDateRange7d         BGPHijackEventListParamsDateRange = "7d"
	BGPHijackEventListParamsDateRange14d        BGPHijackEventListParamsDateRange = "14d"
	BGPHijackEventListParamsDateRange28d        BGPHijackEventListParamsDateRange = "28d"
	BGPHijackEventListParamsDateRange12w        BGPHijackEventListParamsDateRange = "12w"
	BGPHijackEventListParamsDateRange24w        BGPHijackEventListParamsDateRange = "24w"
	BGPHijackEventListParamsDateRange52w        BGPHijackEventListParamsDateRange = "52w"
	BGPHijackEventListParamsDateRange1dControl  BGPHijackEventListParamsDateRange = "1dControl"
	BGPHijackEventListParamsDateRange2dControl  BGPHijackEventListParamsDateRange = "2dControl"
	BGPHijackEventListParamsDateRange7dControl  BGPHijackEventListParamsDateRange = "7dControl"
	BGPHijackEventListParamsDateRange14dControl BGPHijackEventListParamsDateRange = "14dControl"
	BGPHijackEventListParamsDateRange28dControl BGPHijackEventListParamsDateRange = "28dControl"
	BGPHijackEventListParamsDateRange12wControl BGPHijackEventListParamsDateRange = "12wControl"
	BGPHijackEventListParamsDateRange24wControl BGPHijackEventListParamsDateRange = "24wControl"
)

func (r BGPHijackEventListParamsDateRange) IsKnown() bool {
	switch r {
	case BGPHijackEventListParamsDateRange1d, BGPHijackEventListParamsDateRange2d, BGPHijackEventListParamsDateRange7d, BGPHijackEventListParamsDateRange14d, BGPHijackEventListParamsDateRange28d, BGPHijackEventListParamsDateRange12w, BGPHijackEventListParamsDateRange24w, BGPHijackEventListParamsDateRange52w, BGPHijackEventListParamsDateRange1dControl, BGPHijackEventListParamsDateRange2dControl, BGPHijackEventListParamsDateRange7dControl, BGPHijackEventListParamsDateRange14dControl, BGPHijackEventListParamsDateRange28dControl, BGPHijackEventListParamsDateRange12wControl, BGPHijackEventListParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type BGPHijackEventListParamsFormat string

const (
	BGPHijackEventListParamsFormatJson BGPHijackEventListParamsFormat = "JSON"
	BGPHijackEventListParamsFormatCsv  BGPHijackEventListParamsFormat = "CSV"
)

func (r BGPHijackEventListParamsFormat) IsKnown() bool {
	switch r {
	case BGPHijackEventListParamsFormatJson, BGPHijackEventListParamsFormatCsv:
		return true
	}
	return false
}

// Sort events by field
type BGPHijackEventListParamsSortBy string

const (
	BGPHijackEventListParamsSortByID         BGPHijackEventListParamsSortBy = "ID"
	BGPHijackEventListParamsSortByTime       BGPHijackEventListParamsSortBy = "TIME"
	BGPHijackEventListParamsSortByConfidence BGPHijackEventListParamsSortBy = "CONFIDENCE"
)

func (r BGPHijackEventListParamsSortBy) IsKnown() bool {
	switch r {
	case BGPHijackEventListParamsSortByID, BGPHijackEventListParamsSortByTime, BGPHijackEventListParamsSortByConfidence:
		return true
	}
	return false
}

// Sort order
type BGPHijackEventListParamsSortOrder string

const (
	BGPHijackEventListParamsSortOrderAsc  BGPHijackEventListParamsSortOrder = "ASC"
	BGPHijackEventListParamsSortOrderDesc BGPHijackEventListParamsSortOrder = "DESC"
)

func (r BGPHijackEventListParamsSortOrder) IsKnown() bool {
	switch r {
	case BGPHijackEventListParamsSortOrderAsc, BGPHijackEventListParamsSortOrderDesc:
		return true
	}
	return false
}
