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

// RadarBgpLeakEventService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpLeakEventService] method
// instead.
type RadarBgpLeakEventService struct {
	Options []option.RequestOption
}

// NewRadarBgpLeakEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBgpLeakEventService(opts ...option.RequestOption) (r *RadarBgpLeakEventService) {
	r = &RadarBgpLeakEventService{}
	r.Options = opts
	return
}

// Get the BGP route leak events.
func (r *RadarBgpLeakEventService) List(ctx context.Context, query RadarBgpLeakEventListParams, opts ...option.RequestOption) (res *RadarBgpLeakEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/leaks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpLeakEventListResponse struct {
	Result     RadarBgpLeakEventListResponseResult     `json:"result,required"`
	ResultInfo RadarBgpLeakEventListResponseResultInfo `json:"result_info,required"`
	Success    bool                                    `json:"success,required"`
	JSON       radarBgpLeakEventListResponseJSON       `json:"-"`
}

// radarBgpLeakEventListResponseJSON contains the JSON metadata for the struct
// [RadarBgpLeakEventListResponse]
type radarBgpLeakEventListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpLeakEventListResponseResult struct {
	ASNInfo []RadarBgpLeakEventListResponseResultASNInfo `json:"asn_info,required"`
	Events  []RadarBgpLeakEventListResponseResultEvent   `json:"events,required"`
	JSON    radarBgpLeakEventListResponseResultJSON      `json:"-"`
}

// radarBgpLeakEventListResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpLeakEventListResponseResult]
type radarBgpLeakEventListResponseResultJSON struct {
	ASNInfo     apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakEventListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpLeakEventListResponseResultASNInfo struct {
	ASN     int64                                          `json:"asn,required"`
	OrgName string                                         `json:"org_name,required"`
	JSON    radarBgpLeakEventListResponseResultASNInfoJSON `json:"-"`
}

// radarBgpLeakEventListResponseResultASNInfoJSON contains the JSON metadata for
// the struct [RadarBgpLeakEventListResponseResultASNInfo]
type radarBgpLeakEventListResponseResultASNInfoJSON struct {
	ASN         apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakEventListResponseResultASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpLeakEventListResponseResultEvent struct {
	ID          int64                                        `json:"id,required"`
	DetectedTs  string                                       `json:"detected_ts,required"`
	Finished    bool                                         `json:"finished,required"`
	LeakASN     int64                                        `json:"leak_asn,required"`
	LeakCount   int64                                        `json:"leak_count,required"`
	LeakSeg     []int64                                      `json:"leak_seg,required"`
	LeakType    int64                                        `json:"leak_type,required"`
	MaxTs       string                                       `json:"max_ts,required"`
	MinTs       string                                       `json:"min_ts,required"`
	OriginCount int64                                        `json:"origin_count,required"`
	PeerCount   int64                                        `json:"peer_count,required"`
	PrefixCount int64                                        `json:"prefix_count,required"`
	JSON        radarBgpLeakEventListResponseResultEventJSON `json:"-"`
}

// radarBgpLeakEventListResponseResultEventJSON contains the JSON metadata for the
// struct [RadarBgpLeakEventListResponseResultEvent]
type radarBgpLeakEventListResponseResultEventJSON struct {
	ID          apijson.Field
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

func (r *RadarBgpLeakEventListResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpLeakEventListResponseResultInfo struct {
	Count      int64                                       `json:"count,required"`
	Page       int64                                       `json:"page,required"`
	PerPage    int64                                       `json:"per_page,required"`
	TotalCount int64                                       `json:"total_count,required"`
	JSON       radarBgpLeakEventListResponseResultInfoJSON `json:"-"`
}

// radarBgpLeakEventListResponseResultInfoJSON contains the JSON metadata for the
// struct [RadarBgpLeakEventListResponseResultInfo]
type radarBgpLeakEventListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakEventListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBgpLeakEventListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBgpLeakEventListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarBgpLeakEventListParamsFormat] `query:"format"`
	// ASN that is causing or affected by a route leak event
	InvolvedASN param.Field[int64] `query:"involvedAsn"`
	// The leaking AS of a route leak event
	LeakASN param.Field[int64] `query:"leakAsn"`
	// Current page number, starting from 1
	Page param.Field[int64] `query:"page"`
	// Number of entries per page
	PerPage param.Field[int64] `query:"per_page"`
	// Sort events by field
	SortBy param.Field[RadarBgpLeakEventListParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[RadarBgpLeakEventListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [RadarBgpLeakEventListParams]'s query parameters as
// `url.Values`.
func (r RadarBgpLeakEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBgpLeakEventListParamsDateRange string

const (
	RadarBgpLeakEventListParamsDateRange1d         RadarBgpLeakEventListParamsDateRange = "1d"
	RadarBgpLeakEventListParamsDateRange7d         RadarBgpLeakEventListParamsDateRange = "7d"
	RadarBgpLeakEventListParamsDateRange14d        RadarBgpLeakEventListParamsDateRange = "14d"
	RadarBgpLeakEventListParamsDateRange28d        RadarBgpLeakEventListParamsDateRange = "28d"
	RadarBgpLeakEventListParamsDateRange12w        RadarBgpLeakEventListParamsDateRange = "12w"
	RadarBgpLeakEventListParamsDateRange24w        RadarBgpLeakEventListParamsDateRange = "24w"
	RadarBgpLeakEventListParamsDateRange52w        RadarBgpLeakEventListParamsDateRange = "52w"
	RadarBgpLeakEventListParamsDateRange1dControl  RadarBgpLeakEventListParamsDateRange = "1dControl"
	RadarBgpLeakEventListParamsDateRange7dControl  RadarBgpLeakEventListParamsDateRange = "7dControl"
	RadarBgpLeakEventListParamsDateRange14dControl RadarBgpLeakEventListParamsDateRange = "14dControl"
	RadarBgpLeakEventListParamsDateRange28dControl RadarBgpLeakEventListParamsDateRange = "28dControl"
	RadarBgpLeakEventListParamsDateRange12wControl RadarBgpLeakEventListParamsDateRange = "12wControl"
	RadarBgpLeakEventListParamsDateRange24wControl RadarBgpLeakEventListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBgpLeakEventListParamsFormat string

const (
	RadarBgpLeakEventListParamsFormatJson RadarBgpLeakEventListParamsFormat = "JSON"
	RadarBgpLeakEventListParamsFormatCsv  RadarBgpLeakEventListParamsFormat = "CSV"
)

// Sort events by field
type RadarBgpLeakEventListParamsSortBy string

const (
	RadarBgpLeakEventListParamsSortByID       RadarBgpLeakEventListParamsSortBy = "ID"
	RadarBgpLeakEventListParamsSortByLeaks    RadarBgpLeakEventListParamsSortBy = "LEAKS"
	RadarBgpLeakEventListParamsSortByPeers    RadarBgpLeakEventListParamsSortBy = "PEERS"
	RadarBgpLeakEventListParamsSortByPrefixes RadarBgpLeakEventListParamsSortBy = "PREFIXES"
	RadarBgpLeakEventListParamsSortByOrigins  RadarBgpLeakEventListParamsSortBy = "ORIGINS"
	RadarBgpLeakEventListParamsSortByTime     RadarBgpLeakEventListParamsSortBy = "TIME"
)

// Sort order
type RadarBgpLeakEventListParamsSortOrder string

const (
	RadarBgpLeakEventListParamsSortOrderAsc  RadarBgpLeakEventListParamsSortOrder = "ASC"
	RadarBgpLeakEventListParamsSortOrderDesc RadarBgpLeakEventListParamsSortOrder = "DESC"
)
