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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *RadarBGPHijackService) List(ctx context.Context, query RadarBGPHijackListParams, opts ...option.RequestOption) (res *shared.V4PagePagination[RadarBGPHijackListResponse], err error) {
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
func (r *RadarBGPHijackService) ListAutoPaging(ctx context.Context, query RadarBGPHijackListParams, opts ...option.RequestOption) *shared.V4PagePaginationAutoPager[RadarBGPHijackListResponse] {
	return shared.NewV4PagePaginationAutoPager(r.List(ctx, query, opts...))
}

type RadarBGPHijackListResponse struct {
	Result     RadarBGPHijackListResponseResult     `json:"result,required"`
	ResultInfo RadarBGPHijackListResponseResultInfo `json:"result_info,required"`
	Success    bool                                 `json:"success,required"`
	JSON       radarBGPHijackListResponseJSON       `json:"-"`
}

// radarBGPHijackListResponseJSON contains the JSON metadata for the struct
// [RadarBGPHijackListResponse]
type radarBGPHijackListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackListResponseResult struct {
	AsnInfo       []RadarBGPHijackListResponseResultAsnInfo `json:"asn_info,required"`
	Events        []RadarBGPHijackListResponseResultEvent   `json:"events,required"`
	TotalMonitors int64                                     `json:"total_monitors,required"`
	JSON          radarBGPHijackListResponseResultJSON      `json:"-"`
}

// radarBGPHijackListResponseResultJSON contains the JSON metadata for the struct
// [RadarBGPHijackListResponseResult]
type radarBGPHijackListResponseResultJSON struct {
	AsnInfo       apijson.Field
	Events        apijson.Field
	TotalMonitors apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBGPHijackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackListResponseResultAsnInfo struct {
	Asn         int64                                       `json:"asn,required"`
	CountryCode string                                      `json:"country_code,required"`
	OrgName     string                                      `json:"org_name,required"`
	JSON        radarBGPHijackListResponseResultAsnInfoJSON `json:"-"`
}

// radarBGPHijackListResponseResultAsnInfoJSON contains the JSON metadata for the
// struct [RadarBGPHijackListResponseResultAsnInfo]
type radarBGPHijackListResponseResultAsnInfoJSON struct {
	Asn         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackListResponseResultAsnInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackListResponseResultEvent struct {
	ID              int64                                       `json:"id,required"`
	ConfidenceScore int64                                       `json:"confidence_score,required"`
	Duration        int64                                       `json:"duration,required"`
	EventType       int64                                       `json:"event_type,required"`
	HijackMsgsCount int64                                       `json:"hijack_msgs_count,required"`
	HijackerAsn     int64                                       `json:"hijacker_asn,required"`
	HijackerCountry string                                      `json:"hijacker_country,required"`
	IsStale         bool                                        `json:"is_stale,required"`
	MaxHijackTs     string                                      `json:"max_hijack_ts,required"`
	MaxMsgTs        string                                      `json:"max_msg_ts,required"`
	MinHijackTs     string                                      `json:"min_hijack_ts,required"`
	OnGoingCount    int64                                       `json:"on_going_count,required"`
	PeerAsns        []int64                                     `json:"peer_asns,required"`
	PeerIPCount     int64                                       `json:"peer_ip_count,required"`
	Prefixes        []string                                    `json:"prefixes,required"`
	Tags            []RadarBGPHijackListResponseResultEventsTag `json:"tags,required"`
	VictimAsns      []int64                                     `json:"victim_asns,required"`
	VictimCountries []string                                    `json:"victim_countries,required"`
	JSON            radarBGPHijackListResponseResultEventJSON   `json:"-"`
}

// radarBGPHijackListResponseResultEventJSON contains the JSON metadata for the
// struct [RadarBGPHijackListResponseResultEvent]
type radarBGPHijackListResponseResultEventJSON struct {
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

func (r *RadarBGPHijackListResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackListResponseResultEventsTag struct {
	Name  string                                        `json:"name,required"`
	Score int64                                         `json:"score,required"`
	JSON  radarBGPHijackListResponseResultEventsTagJSON `json:"-"`
}

// radarBGPHijackListResponseResultEventsTagJSON contains the JSON metadata for the
// struct [RadarBGPHijackListResponseResultEventsTag]
type radarBGPHijackListResponseResultEventsTagJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackListResponseResultEventsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackListResponseResultInfo struct {
	Count      int64                                    `json:"count,required"`
	Page       int64                                    `json:"page,required"`
	PerPage    int64                                    `json:"per_page,required"`
	TotalCount int64                                    `json:"total_count,required"`
	JSON       radarBGPHijackListResponseResultInfoJSON `json:"-"`
}

// radarBGPHijackListResponseResultInfoJSON contains the JSON metadata for the
// struct [RadarBGPHijackListResponseResultInfo]
type radarBGPHijackListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPHijackListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPHijackListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBGPHijackListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[RadarBGPHijackListParamsFormat] `query:"format"`
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
	SortBy param.Field[RadarBGPHijackListParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[RadarBGPHijackListParamsSortOrder] `query:"sortOrder"`
	// The potential victim AS of a BGP hijack event
	VictimAsn param.Field[int64] `query:"victimAsn"`
}

// URLQuery serializes [RadarBGPHijackListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPHijackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBGPHijackListParamsDateRange string

const (
	RadarBGPHijackListParamsDateRange1d         RadarBGPHijackListParamsDateRange = "1d"
	RadarBGPHijackListParamsDateRange2d         RadarBGPHijackListParamsDateRange = "2d"
	RadarBGPHijackListParamsDateRange7d         RadarBGPHijackListParamsDateRange = "7d"
	RadarBGPHijackListParamsDateRange14d        RadarBGPHijackListParamsDateRange = "14d"
	RadarBGPHijackListParamsDateRange28d        RadarBGPHijackListParamsDateRange = "28d"
	RadarBGPHijackListParamsDateRange12w        RadarBGPHijackListParamsDateRange = "12w"
	RadarBGPHijackListParamsDateRange24w        RadarBGPHijackListParamsDateRange = "24w"
	RadarBGPHijackListParamsDateRange52w        RadarBGPHijackListParamsDateRange = "52w"
	RadarBGPHijackListParamsDateRange1dControl  RadarBGPHijackListParamsDateRange = "1dControl"
	RadarBGPHijackListParamsDateRange2dControl  RadarBGPHijackListParamsDateRange = "2dControl"
	RadarBGPHijackListParamsDateRange7dControl  RadarBGPHijackListParamsDateRange = "7dControl"
	RadarBGPHijackListParamsDateRange14dControl RadarBGPHijackListParamsDateRange = "14dControl"
	RadarBGPHijackListParamsDateRange28dControl RadarBGPHijackListParamsDateRange = "28dControl"
	RadarBGPHijackListParamsDateRange12wControl RadarBGPHijackListParamsDateRange = "12wControl"
	RadarBGPHijackListParamsDateRange24wControl RadarBGPHijackListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPHijackListParamsFormat string

const (
	RadarBGPHijackListParamsFormatJson RadarBGPHijackListParamsFormat = "JSON"
	RadarBGPHijackListParamsFormatCsv  RadarBGPHijackListParamsFormat = "CSV"
)

// Sort events by field
type RadarBGPHijackListParamsSortBy string

const (
	RadarBGPHijackListParamsSortByID         RadarBGPHijackListParamsSortBy = "ID"
	RadarBGPHijackListParamsSortByTime       RadarBGPHijackListParamsSortBy = "TIME"
	RadarBGPHijackListParamsSortByConfidence RadarBGPHijackListParamsSortBy = "CONFIDENCE"
)

// Sort order
type RadarBGPHijackListParamsSortOrder string

const (
	RadarBGPHijackListParamsSortOrderAsc  RadarBGPHijackListParamsSortOrder = "ASC"
	RadarBGPHijackListParamsSortOrderDesc RadarBGPHijackListParamsSortOrder = "DESC"
)
