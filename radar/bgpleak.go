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

// BGPLeakService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewBGPLeakService] method instead.
type BGPLeakService struct {
	Options []option.RequestOption
}

// NewBGPLeakService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPLeakService(opts ...option.RequestOption) (r *BGPLeakService) {
	r = &BGPLeakService{}
	r.Options = opts
	return
}

// Get the BGP route leak events (Beta).
func (r *BGPLeakService) Events(ctx context.Context, query BGPLeakEventsParams, opts ...option.RequestOption) (res *BGPLeakEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BGPLeakEventsResponseEnvelope
	path := "radar/bgp/leaks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPLeakEventsResponse struct {
	ASNInfo []BGPLeakEventsResponseASNInfo `json:"asn_info,required"`
	Events  []BGPLeakEventsResponseEvent   `json:"events,required"`
	JSON    bgpLeakEventsResponseJSON      `json:"-"`
}

// bgpLeakEventsResponseJSON contains the JSON metadata for the struct
// [BGPLeakEventsResponse]
type bgpLeakEventsResponseJSON struct {
	ASNInfo     apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventsResponseJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventsResponseASNInfo struct {
	ASN         int64                            `json:"asn,required"`
	CountryCode string                           `json:"country_code,required"`
	OrgName     string                           `json:"org_name,required"`
	JSON        bgpLeakEventsResponseASNInfoJSON `json:"-"`
}

// bgpLeakEventsResponseASNInfoJSON contains the JSON metadata for the struct
// [BGPLeakEventsResponseASNInfo]
type bgpLeakEventsResponseASNInfoJSON struct {
	ASN         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventsResponseASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventsResponseASNInfoJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventsResponseEvent struct {
	ID          int64                          `json:"id,required"`
	Countries   []string                       `json:"countries,required"`
	DetectedTs  string                         `json:"detected_ts,required"`
	Finished    bool                           `json:"finished,required"`
	LeakASN     int64                          `json:"leak_asn,required"`
	LeakCount   int64                          `json:"leak_count,required"`
	LeakSeg     []int64                        `json:"leak_seg,required"`
	LeakType    int64                          `json:"leak_type,required"`
	MaxTs       string                         `json:"max_ts,required"`
	MinTs       string                         `json:"min_ts,required"`
	OriginCount int64                          `json:"origin_count,required"`
	PeerCount   int64                          `json:"peer_count,required"`
	PrefixCount int64                          `json:"prefix_count,required"`
	JSON        bgpLeakEventsResponseEventJSON `json:"-"`
}

// bgpLeakEventsResponseEventJSON contains the JSON metadata for the struct
// [BGPLeakEventsResponseEvent]
type bgpLeakEventsResponseEventJSON struct {
	ID          apijson.Field
	Countries   apijson.Field
	DetectedTs  apijson.Field
	Finished    apijson.Field
	LeakASN     apijson.Field
	LeakCount   apijson.Field
	LeakSeg     apijson.Field
	LeakType    apijson.Field
	MaxTs       apijson.Field
	MinTs       apijson.Field
	OriginCount apijson.Field
	PeerCount   apijson.Field
	PrefixCount apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventsResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventsResponseEventJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[BGPLeakEventsParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[BGPLeakEventsParamsFormat] `query:"format"`
	// ASN that is causing or affected by a route leak event
	InvolvedASN param.Field[int64] `query:"involvedAsn"`
	// Country code of a involved ASN in a route leak event
	InvolvedCountry param.Field[string] `query:"involvedCountry"`
	// The leaking AS of a route leak event
	LeakASN param.Field[int64] `query:"leakAsn"`
	// Current page number, starting from 1
	Page param.Field[int64] `query:"page"`
	// Number of entries per page
	PerPage param.Field[int64] `query:"per_page"`
	// Sort events by field
	SortBy param.Field[BGPLeakEventsParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[BGPLeakEventsParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [BGPLeakEventsParams]'s query parameters as `url.Values`.
func (r BGPLeakEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type BGPLeakEventsParamsDateRange string

const (
	BGPLeakEventsParamsDateRange1d         BGPLeakEventsParamsDateRange = "1d"
	BGPLeakEventsParamsDateRange2d         BGPLeakEventsParamsDateRange = "2d"
	BGPLeakEventsParamsDateRange7d         BGPLeakEventsParamsDateRange = "7d"
	BGPLeakEventsParamsDateRange14d        BGPLeakEventsParamsDateRange = "14d"
	BGPLeakEventsParamsDateRange28d        BGPLeakEventsParamsDateRange = "28d"
	BGPLeakEventsParamsDateRange12w        BGPLeakEventsParamsDateRange = "12w"
	BGPLeakEventsParamsDateRange24w        BGPLeakEventsParamsDateRange = "24w"
	BGPLeakEventsParamsDateRange52w        BGPLeakEventsParamsDateRange = "52w"
	BGPLeakEventsParamsDateRange1dControl  BGPLeakEventsParamsDateRange = "1dControl"
	BGPLeakEventsParamsDateRange2dControl  BGPLeakEventsParamsDateRange = "2dControl"
	BGPLeakEventsParamsDateRange7dControl  BGPLeakEventsParamsDateRange = "7dControl"
	BGPLeakEventsParamsDateRange14dControl BGPLeakEventsParamsDateRange = "14dControl"
	BGPLeakEventsParamsDateRange28dControl BGPLeakEventsParamsDateRange = "28dControl"
	BGPLeakEventsParamsDateRange12wControl BGPLeakEventsParamsDateRange = "12wControl"
	BGPLeakEventsParamsDateRange24wControl BGPLeakEventsParamsDateRange = "24wControl"
)

func (r BGPLeakEventsParamsDateRange) IsKnown() bool {
	switch r {
	case BGPLeakEventsParamsDateRange1d, BGPLeakEventsParamsDateRange2d, BGPLeakEventsParamsDateRange7d, BGPLeakEventsParamsDateRange14d, BGPLeakEventsParamsDateRange28d, BGPLeakEventsParamsDateRange12w, BGPLeakEventsParamsDateRange24w, BGPLeakEventsParamsDateRange52w, BGPLeakEventsParamsDateRange1dControl, BGPLeakEventsParamsDateRange2dControl, BGPLeakEventsParamsDateRange7dControl, BGPLeakEventsParamsDateRange14dControl, BGPLeakEventsParamsDateRange28dControl, BGPLeakEventsParamsDateRange12wControl, BGPLeakEventsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type BGPLeakEventsParamsFormat string

const (
	BGPLeakEventsParamsFormatJson BGPLeakEventsParamsFormat = "JSON"
	BGPLeakEventsParamsFormatCsv  BGPLeakEventsParamsFormat = "CSV"
)

func (r BGPLeakEventsParamsFormat) IsKnown() bool {
	switch r {
	case BGPLeakEventsParamsFormatJson, BGPLeakEventsParamsFormatCsv:
		return true
	}
	return false
}

// Sort events by field
type BGPLeakEventsParamsSortBy string

const (
	BGPLeakEventsParamsSortByID       BGPLeakEventsParamsSortBy = "ID"
	BGPLeakEventsParamsSortByLeaks    BGPLeakEventsParamsSortBy = "LEAKS"
	BGPLeakEventsParamsSortByPeers    BGPLeakEventsParamsSortBy = "PEERS"
	BGPLeakEventsParamsSortByPrefixes BGPLeakEventsParamsSortBy = "PREFIXES"
	BGPLeakEventsParamsSortByOrigins  BGPLeakEventsParamsSortBy = "ORIGINS"
	BGPLeakEventsParamsSortByTime     BGPLeakEventsParamsSortBy = "TIME"
)

func (r BGPLeakEventsParamsSortBy) IsKnown() bool {
	switch r {
	case BGPLeakEventsParamsSortByID, BGPLeakEventsParamsSortByLeaks, BGPLeakEventsParamsSortByPeers, BGPLeakEventsParamsSortByPrefixes, BGPLeakEventsParamsSortByOrigins, BGPLeakEventsParamsSortByTime:
		return true
	}
	return false
}

// Sort order
type BGPLeakEventsParamsSortOrder string

const (
	BGPLeakEventsParamsSortOrderAsc  BGPLeakEventsParamsSortOrder = "ASC"
	BGPLeakEventsParamsSortOrderDesc BGPLeakEventsParamsSortOrder = "DESC"
)

func (r BGPLeakEventsParamsSortOrder) IsKnown() bool {
	switch r {
	case BGPLeakEventsParamsSortOrderAsc, BGPLeakEventsParamsSortOrderDesc:
		return true
	}
	return false
}

type BGPLeakEventsResponseEnvelope struct {
	Result     BGPLeakEventsResponse                   `json:"result,required"`
	ResultInfo BGPLeakEventsResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                    `json:"success,required"`
	JSON       bgpLeakEventsResponseEnvelopeJSON       `json:"-"`
}

// bgpLeakEventsResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPLeakEventsResponseEnvelope]
type bgpLeakEventsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventsResponseEnvelopeResultInfo struct {
	Count      int64                                       `json:"count,required"`
	Page       int64                                       `json:"page,required"`
	PerPage    int64                                       `json:"per_page,required"`
	TotalCount int64                                       `json:"total_count,required"`
	JSON       bgpLeakEventsResponseEnvelopeResultInfoJSON `json:"-"`
}

// bgpLeakEventsResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [BGPLeakEventsResponseEnvelopeResultInfo]
type bgpLeakEventsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventsResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
