// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// AnalyticsService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options []option.RequestOption
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r *AnalyticsService) {
	r = &AnalyticsService{}
	r.Options = opts
	return
}

// Returns day-wise session and recording analytics data of an App for the
// specified time range start_date to end_date. If start_date and end_date are not
// provided, the default time range is set from 30 days ago to the current date.
func (r *AnalyticsService) GetOrgAnalytics(ctx context.Context, appID string, params AnalyticsGetOrgAnalyticsParams, opts ...option.RequestOption) (res *AnalyticsGetOrgAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/analytics/daywise", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type AnalyticsGetOrgAnalyticsResponse struct {
	Data    AnalyticsGetOrgAnalyticsResponseData `json:"data"`
	Success bool                                 `json:"success"`
	JSON    analyticsGetOrgAnalyticsResponseJSON `json:"-"`
}

// analyticsGetOrgAnalyticsResponseJSON contains the JSON metadata for the struct
// [AnalyticsGetOrgAnalyticsResponse]
type analyticsGetOrgAnalyticsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsGetOrgAnalyticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetOrgAnalyticsResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetOrgAnalyticsResponseData struct {
	// Recording statistics of an App during the range specified
	RecordingStats AnalyticsGetOrgAnalyticsResponseDataRecordingStats `json:"recording_stats"`
	// Session statistics of an App during the range specified
	SessionStats AnalyticsGetOrgAnalyticsResponseDataSessionStats `json:"session_stats"`
	JSON         analyticsGetOrgAnalyticsResponseDataJSON         `json:"-"`
}

// analyticsGetOrgAnalyticsResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsGetOrgAnalyticsResponseData]
type analyticsGetOrgAnalyticsResponseDataJSON struct {
	RecordingStats apijson.Field
	SessionStats   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AnalyticsGetOrgAnalyticsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetOrgAnalyticsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Recording statistics of an App during the range specified
type AnalyticsGetOrgAnalyticsResponseDataRecordingStats struct {
	// Day wise recording stats
	DayStats []AnalyticsGetOrgAnalyticsResponseDataRecordingStatsDayStat `json:"day_stats"`
	// Total number of recordings during the range specified
	RecordingCount int64 `json:"recording_count"`
	// Total recording minutes during the range specified
	RecordingMinutesConsumed float64                                                `json:"recording_minutes_consumed"`
	JSON                     analyticsGetOrgAnalyticsResponseDataRecordingStatsJSON `json:"-"`
}

// analyticsGetOrgAnalyticsResponseDataRecordingStatsJSON contains the JSON
// metadata for the struct [AnalyticsGetOrgAnalyticsResponseDataRecordingStats]
type analyticsGetOrgAnalyticsResponseDataRecordingStatsJSON struct {
	DayStats                 apijson.Field
	RecordingCount           apijson.Field
	RecordingMinutesConsumed apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AnalyticsGetOrgAnalyticsResponseDataRecordingStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetOrgAnalyticsResponseDataRecordingStatsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetOrgAnalyticsResponseDataRecordingStatsDayStat struct {
	Day string `json:"day"`
	// Total recording minutes for a specific day
	TotalRecordingMinutes int64 `json:"total_recording_minutes"`
	// Total number of recordings for a specific day
	TotalRecordings int64                                                         `json:"total_recordings"`
	JSON            analyticsGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON `json:"-"`
}

// analyticsGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON contains the JSON
// metadata for the struct
// [AnalyticsGetOrgAnalyticsResponseDataRecordingStatsDayStat]
type analyticsGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON struct {
	Day                   apijson.Field
	TotalRecordingMinutes apijson.Field
	TotalRecordings       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AnalyticsGetOrgAnalyticsResponseDataRecordingStatsDayStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetOrgAnalyticsResponseDataRecordingStatsDayStatJSON) RawJSON() string {
	return r.raw
}

// Session statistics of an App during the range specified
type AnalyticsGetOrgAnalyticsResponseDataSessionStats struct {
	// Day wise session stats
	DayStats []AnalyticsGetOrgAnalyticsResponseDataSessionStatsDayStat `json:"day_stats"`
	// Total number of sessions during the range specified
	SessionsCount int64 `json:"sessions_count"`
	// Total session minutes during the range specified
	SessionsMinutesConsumed float64                                              `json:"sessions_minutes_consumed"`
	JSON                    analyticsGetOrgAnalyticsResponseDataSessionStatsJSON `json:"-"`
}

// analyticsGetOrgAnalyticsResponseDataSessionStatsJSON contains the JSON metadata
// for the struct [AnalyticsGetOrgAnalyticsResponseDataSessionStats]
type analyticsGetOrgAnalyticsResponseDataSessionStatsJSON struct {
	DayStats                apijson.Field
	SessionsCount           apijson.Field
	SessionsMinutesConsumed apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AnalyticsGetOrgAnalyticsResponseDataSessionStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetOrgAnalyticsResponseDataSessionStatsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetOrgAnalyticsResponseDataSessionStatsDayStat struct {
	Day string `json:"day"`
	// Total session minutes for a specific day
	TotalSessionMinutes float64 `json:"total_session_minutes"`
	// Total number of sessions for a specific day
	TotalSessions int64                                                       `json:"total_sessions"`
	JSON          analyticsGetOrgAnalyticsResponseDataSessionStatsDayStatJSON `json:"-"`
}

// analyticsGetOrgAnalyticsResponseDataSessionStatsDayStatJSON contains the JSON
// metadata for the struct
// [AnalyticsGetOrgAnalyticsResponseDataSessionStatsDayStat]
type analyticsGetOrgAnalyticsResponseDataSessionStatsDayStatJSON struct {
	Day                 apijson.Field
	TotalSessionMinutes apijson.Field
	TotalSessions       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AnalyticsGetOrgAnalyticsResponseDataSessionStatsDayStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetOrgAnalyticsResponseDataSessionStatsDayStatJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetOrgAnalyticsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// end date in YYYY-MM-DD format
	EndDate param.Field[string] `query:"end_date"`
	// start date in YYYY-MM-DD format
	StartDate param.Field[string] `query:"start_date"`
}

// URLQuery serializes [AnalyticsGetOrgAnalyticsParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetOrgAnalyticsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
