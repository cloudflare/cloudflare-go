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

// RadarBGPLeakEventService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPLeakEventService] method
// instead.
type RadarBGPLeakEventService struct {
	Options []option.RequestOption
}

// NewRadarBGPLeakEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarBGPLeakEventService(opts ...option.RequestOption) (r *RadarBGPLeakEventService) {
	r = &RadarBGPLeakEventService{}
	r.Options = opts
	return
}

// Get the BGP route leak events (Beta).
func (r *RadarBGPLeakEventService) List(ctx context.Context, query RadarBGPLeakEventListParams, opts ...option.RequestOption) (res *RadarBGPLeakEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPLeakEventListResponseEnvelope
	path := "radar/bgp/leaks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPLeakEventListResponse struct {
	AsnInfo []RadarBGPLeakEventListResponseAsnInfo `json:"asn_info,required"`
	Events  []RadarBGPLeakEventListResponseEvent   `json:"events,required"`
	JSON    radarBGPLeakEventListResponseJSON      `json:"-"`
}

// radarBGPLeakEventListResponseJSON contains the JSON metadata for the struct
// [RadarBGPLeakEventListResponse]
type radarBGPLeakEventListResponseJSON struct {
	AsnInfo     apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventListResponseAsnInfo struct {
	Asn         int64                                    `json:"asn,required"`
	CountryCode string                                   `json:"country_code,required"`
	OrgName     string                                   `json:"org_name,required"`
	JSON        radarBGPLeakEventListResponseAsnInfoJSON `json:"-"`
}

// radarBGPLeakEventListResponseAsnInfoJSON contains the JSON metadata for the
// struct [RadarBGPLeakEventListResponseAsnInfo]
type radarBGPLeakEventListResponseAsnInfoJSON struct {
	Asn         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventListResponseAsnInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventListResponseEvent struct {
	ID          int64                                  `json:"id,required"`
	Countries   []string                               `json:"countries,required"`
	DetectedTs  string                                 `json:"detected_ts,required"`
	Finished    bool                                   `json:"finished,required"`
	LeakAsn     int64                                  `json:"leak_asn,required"`
	LeakCount   int64                                  `json:"leak_count,required"`
	LeakSeg     []int64                                `json:"leak_seg,required"`
	LeakType    int64                                  `json:"leak_type,required"`
	MaxTs       string                                 `json:"max_ts,required"`
	MinTs       string                                 `json:"min_ts,required"`
	OriginCount int64                                  `json:"origin_count,required"`
	PeerCount   int64                                  `json:"peer_count,required"`
	PrefixCount int64                                  `json:"prefix_count,required"`
	JSON        radarBGPLeakEventListResponseEventJSON `json:"-"`
}

// radarBGPLeakEventListResponseEventJSON contains the JSON metadata for the struct
// [RadarBGPLeakEventListResponseEvent]
type radarBGPLeakEventListResponseEventJSON struct {
	ID          apijson.Field
	Countries   apijson.Field
	DetectedTs  apijson.Field
	Finished    apijson.Field
	LeakAsn     apijson.Field
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

func (r *RadarBGPLeakEventListResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBGPLeakEventListParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[RadarBGPLeakEventListParamsFormat] `query:"format"`
	// ASN that is causing or affected by a route leak event
	InvolvedAsn param.Field[int64] `query:"involvedAsn"`
	// Country code of a involved ASN in a route leak event
	InvolvedCountry param.Field[string] `query:"involvedCountry"`
	// The leaking AS of a route leak event
	LeakAsn param.Field[int64] `query:"leakAsn"`
	// Current page number, starting from 1
	Page param.Field[int64] `query:"page"`
	// Number of entries per page
	PerPage param.Field[int64] `query:"per_page"`
	// Sort events by field
	SortBy param.Field[RadarBGPLeakEventListParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[RadarBGPLeakEventListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [RadarBGPLeakEventListParams]'s query parameters as
// `url.Values`.
func (r RadarBGPLeakEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBGPLeakEventListParamsDateRange string

const (
	RadarBGPLeakEventListParamsDateRange1d         RadarBGPLeakEventListParamsDateRange = "1d"
	RadarBGPLeakEventListParamsDateRange2d         RadarBGPLeakEventListParamsDateRange = "2d"
	RadarBGPLeakEventListParamsDateRange7d         RadarBGPLeakEventListParamsDateRange = "7d"
	RadarBGPLeakEventListParamsDateRange14d        RadarBGPLeakEventListParamsDateRange = "14d"
	RadarBGPLeakEventListParamsDateRange28d        RadarBGPLeakEventListParamsDateRange = "28d"
	RadarBGPLeakEventListParamsDateRange12w        RadarBGPLeakEventListParamsDateRange = "12w"
	RadarBGPLeakEventListParamsDateRange24w        RadarBGPLeakEventListParamsDateRange = "24w"
	RadarBGPLeakEventListParamsDateRange52w        RadarBGPLeakEventListParamsDateRange = "52w"
	RadarBGPLeakEventListParamsDateRange1dControl  RadarBGPLeakEventListParamsDateRange = "1dControl"
	RadarBGPLeakEventListParamsDateRange2dControl  RadarBGPLeakEventListParamsDateRange = "2dControl"
	RadarBGPLeakEventListParamsDateRange7dControl  RadarBGPLeakEventListParamsDateRange = "7dControl"
	RadarBGPLeakEventListParamsDateRange14dControl RadarBGPLeakEventListParamsDateRange = "14dControl"
	RadarBGPLeakEventListParamsDateRange28dControl RadarBGPLeakEventListParamsDateRange = "28dControl"
	RadarBGPLeakEventListParamsDateRange12wControl RadarBGPLeakEventListParamsDateRange = "12wControl"
	RadarBGPLeakEventListParamsDateRange24wControl RadarBGPLeakEventListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPLeakEventListParamsFormat string

const (
	RadarBGPLeakEventListParamsFormatJson RadarBGPLeakEventListParamsFormat = "JSON"
	RadarBGPLeakEventListParamsFormatCsv  RadarBGPLeakEventListParamsFormat = "CSV"
)

// Sort events by field
type RadarBGPLeakEventListParamsSortBy string

const (
	RadarBGPLeakEventListParamsSortByID       RadarBGPLeakEventListParamsSortBy = "ID"
	RadarBGPLeakEventListParamsSortByLeaks    RadarBGPLeakEventListParamsSortBy = "LEAKS"
	RadarBGPLeakEventListParamsSortByPeers    RadarBGPLeakEventListParamsSortBy = "PEERS"
	RadarBGPLeakEventListParamsSortByPrefixes RadarBGPLeakEventListParamsSortBy = "PREFIXES"
	RadarBGPLeakEventListParamsSortByOrigins  RadarBGPLeakEventListParamsSortBy = "ORIGINS"
	RadarBGPLeakEventListParamsSortByTime     RadarBGPLeakEventListParamsSortBy = "TIME"
)

// Sort order
type RadarBGPLeakEventListParamsSortOrder string

const (
	RadarBGPLeakEventListParamsSortOrderAsc  RadarBGPLeakEventListParamsSortOrder = "ASC"
	RadarBGPLeakEventListParamsSortOrderDesc RadarBGPLeakEventListParamsSortOrder = "DESC"
)

type RadarBGPLeakEventListResponseEnvelope struct {
	Result     RadarBGPLeakEventListResponse                   `json:"result,required"`
	ResultInfo RadarBGPLeakEventListResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                            `json:"success,required"`
	JSON       radarBGPLeakEventListResponseEnvelopeJSON       `json:"-"`
}

// radarBGPLeakEventListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarBGPLeakEventListResponseEnvelope]
type radarBGPLeakEventListResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventListResponseEnvelopeResultInfo struct {
	Count      int64                                               `json:"count,required"`
	Page       int64                                               `json:"page,required"`
	PerPage    int64                                               `json:"per_page,required"`
	TotalCount int64                                               `json:"total_count,required"`
	JSON       radarBGPLeakEventListResponseEnvelopeResultInfoJSON `json:"-"`
}

// radarBGPLeakEventListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [RadarBGPLeakEventListResponseEnvelopeResultInfo]
type radarBGPLeakEventListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
