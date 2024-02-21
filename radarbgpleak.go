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

// RadarBGPLeakService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPLeakService] method
// instead.
type RadarBGPLeakService struct {
	Options []option.RequestOption
}

// NewRadarBGPLeakService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPLeakService(opts ...option.RequestOption) (r *RadarBGPLeakService) {
	r = &RadarBGPLeakService{}
	r.Options = opts
	return
}

// Get the BGP route leak events (Beta).
func (r *RadarBGPLeakService) Events(ctx context.Context, query RadarBGPLeakEventsParams, opts ...option.RequestOption) (res *RadarBGPLeakEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarBGPLeakEventsResponseEnvelope
	path := "radar/bgp/leaks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarBGPLeakEventsResponse struct {
	AsnInfo []RadarBGPLeakEventsResponseAsnInfo `json:"asn_info,required"`
	Events  []RadarBGPLeakEventsResponseEvent   `json:"events,required"`
	JSON    radarBGPLeakEventsResponseJSON      `json:"-"`
}

// radarBGPLeakEventsResponseJSON contains the JSON metadata for the struct
// [RadarBGPLeakEventsResponse]
type radarBGPLeakEventsResponseJSON struct {
	AsnInfo     apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventsResponseAsnInfo struct {
	Asn         int64                                 `json:"asn,required"`
	CountryCode string                                `json:"country_code,required"`
	OrgName     string                                `json:"org_name,required"`
	JSON        radarBGPLeakEventsResponseAsnInfoJSON `json:"-"`
}

// radarBGPLeakEventsResponseAsnInfoJSON contains the JSON metadata for the struct
// [RadarBGPLeakEventsResponseAsnInfo]
type radarBGPLeakEventsResponseAsnInfoJSON struct {
	Asn         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventsResponseAsnInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventsResponseEvent struct {
	ID          int64                               `json:"id,required"`
	Countries   []string                            `json:"countries,required"`
	DetectedTs  string                              `json:"detected_ts,required"`
	Finished    bool                                `json:"finished,required"`
	LeakAsn     int64                               `json:"leak_asn,required"`
	LeakCount   int64                               `json:"leak_count,required"`
	LeakSeg     []int64                             `json:"leak_seg,required"`
	LeakType    int64                               `json:"leak_type,required"`
	MaxTs       string                              `json:"max_ts,required"`
	MinTs       string                              `json:"min_ts,required"`
	OriginCount int64                               `json:"origin_count,required"`
	PeerCount   int64                               `json:"peer_count,required"`
	PrefixCount int64                               `json:"prefix_count,required"`
	JSON        radarBGPLeakEventsResponseEventJSON `json:"-"`
}

// radarBGPLeakEventsResponseEventJSON contains the JSON metadata for the struct
// [RadarBGPLeakEventsResponseEvent]
type radarBGPLeakEventsResponseEventJSON struct {
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

func (r *RadarBGPLeakEventsResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[RadarBGPLeakEventsParamsDateRange] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event
	EventID param.Field[int64] `query:"eventId"`
	// Format results are returned in.
	Format param.Field[RadarBGPLeakEventsParamsFormat] `query:"format"`
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
	SortBy param.Field[RadarBGPLeakEventsParamsSortBy] `query:"sortBy"`
	// Sort order
	SortOrder param.Field[RadarBGPLeakEventsParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [RadarBGPLeakEventsParams]'s query parameters as
// `url.Values`.
func (r RadarBGPLeakEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Shorthand date ranges for the last X days - use when you don't need specific
// start and end dates.
type RadarBGPLeakEventsParamsDateRange string

const (
	RadarBGPLeakEventsParamsDateRange1d         RadarBGPLeakEventsParamsDateRange = "1d"
	RadarBGPLeakEventsParamsDateRange2d         RadarBGPLeakEventsParamsDateRange = "2d"
	RadarBGPLeakEventsParamsDateRange7d         RadarBGPLeakEventsParamsDateRange = "7d"
	RadarBGPLeakEventsParamsDateRange14d        RadarBGPLeakEventsParamsDateRange = "14d"
	RadarBGPLeakEventsParamsDateRange28d        RadarBGPLeakEventsParamsDateRange = "28d"
	RadarBGPLeakEventsParamsDateRange12w        RadarBGPLeakEventsParamsDateRange = "12w"
	RadarBGPLeakEventsParamsDateRange24w        RadarBGPLeakEventsParamsDateRange = "24w"
	RadarBGPLeakEventsParamsDateRange52w        RadarBGPLeakEventsParamsDateRange = "52w"
	RadarBGPLeakEventsParamsDateRange1dControl  RadarBGPLeakEventsParamsDateRange = "1dControl"
	RadarBGPLeakEventsParamsDateRange2dControl  RadarBGPLeakEventsParamsDateRange = "2dControl"
	RadarBGPLeakEventsParamsDateRange7dControl  RadarBGPLeakEventsParamsDateRange = "7dControl"
	RadarBGPLeakEventsParamsDateRange14dControl RadarBGPLeakEventsParamsDateRange = "14dControl"
	RadarBGPLeakEventsParamsDateRange28dControl RadarBGPLeakEventsParamsDateRange = "28dControl"
	RadarBGPLeakEventsParamsDateRange12wControl RadarBGPLeakEventsParamsDateRange = "12wControl"
	RadarBGPLeakEventsParamsDateRange24wControl RadarBGPLeakEventsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarBGPLeakEventsParamsFormat string

const (
	RadarBGPLeakEventsParamsFormatJson RadarBGPLeakEventsParamsFormat = "JSON"
	RadarBGPLeakEventsParamsFormatCsv  RadarBGPLeakEventsParamsFormat = "CSV"
)

// Sort events by field
type RadarBGPLeakEventsParamsSortBy string

const (
	RadarBGPLeakEventsParamsSortByID       RadarBGPLeakEventsParamsSortBy = "ID"
	RadarBGPLeakEventsParamsSortByLeaks    RadarBGPLeakEventsParamsSortBy = "LEAKS"
	RadarBGPLeakEventsParamsSortByPeers    RadarBGPLeakEventsParamsSortBy = "PEERS"
	RadarBGPLeakEventsParamsSortByPrefixes RadarBGPLeakEventsParamsSortBy = "PREFIXES"
	RadarBGPLeakEventsParamsSortByOrigins  RadarBGPLeakEventsParamsSortBy = "ORIGINS"
	RadarBGPLeakEventsParamsSortByTime     RadarBGPLeakEventsParamsSortBy = "TIME"
)

// Sort order
type RadarBGPLeakEventsParamsSortOrder string

const (
	RadarBGPLeakEventsParamsSortOrderAsc  RadarBGPLeakEventsParamsSortOrder = "ASC"
	RadarBGPLeakEventsParamsSortOrderDesc RadarBGPLeakEventsParamsSortOrder = "DESC"
)

type RadarBGPLeakEventsResponseEnvelope struct {
	Result     RadarBGPLeakEventsResponse                   `json:"result,required"`
	ResultInfo RadarBGPLeakEventsResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                         `json:"success,required"`
	JSON       radarBGPLeakEventsResponseEnvelopeJSON       `json:"-"`
}

// radarBGPLeakEventsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarBGPLeakEventsResponseEnvelope]
type radarBGPLeakEventsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarBGPLeakEventsResponseEnvelopeResultInfo struct {
	Count      int64                                            `json:"count,required"`
	Page       int64                                            `json:"page,required"`
	PerPage    int64                                            `json:"per_page,required"`
	TotalCount int64                                            `json:"total_count,required"`
	JSON       radarBGPLeakEventsResponseEnvelopeResultInfoJSON `json:"-"`
}

// radarBGPLeakEventsResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [RadarBGPLeakEventsResponseEnvelopeResultInfo]
type radarBGPLeakEventsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBGPLeakEventsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
