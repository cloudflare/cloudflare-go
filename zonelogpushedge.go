// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogpushEdgeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogpushEdgeService] method
// instead.
type ZoneLogpushEdgeService struct {
	Options []option.RequestOption
}

// NewZoneLogpushEdgeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogpushEdgeService(opts ...option.RequestOption) (r *ZoneLogpushEdgeService) {
	r = &ZoneLogpushEdgeService{}
	r.Options = opts
	return
}

// Lists Instant Logs jobs for a zone.
func (r *ZoneLogpushEdgeService) GetZonesZoneIdentifierLogpushEdgeJobs(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new Instant Logs job for a zone.
func (r *ZoneLogpushEdgeService) PostZonesZoneIdentifierLogpushEdgeJobs(ctx context.Context, zoneIdentifier string, body ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsParams, opts ...option.RequestOption) (res *ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponse struct {
	Errors   []ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseError   `json:"errors"`
	Messages []ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessage `json:"messages"`
	Result   []ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResult  `json:"result"`
	// Whether the API call was successful
	Success ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseSuccess `json:"success"`
	JSON    zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseJSON    `json:"-"`
}

// zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseJSON contains the
// JSON metadata for the struct
// [ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponse]
type zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseErrorJSON `json:"-"`
}

// zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseError]
type zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessageJSON `json:"-"`
}

// zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessage]
type zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResult struct {
	// Unique WebSocket address that will receive messages from Cloudflare’s edge.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Comma-separated list of fields.
	Fields string `json:"fields"`
	// Filters to drill down into specific events.
	Filter string `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample int64 `json:"sample"`
	// Unique session id of the job.
	SessionID string                                                                 `json:"session_id"`
	JSON      zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResultJSON `json:"-"`
}

// zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResult]
type zoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResultJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseSuccess bool

const (
	ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseSuccessTrue ZoneLogpushEdgeGetZonesZoneIdentifierLogpushEdgeJobsResponseSuccess = true
)

type ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponse struct {
	Errors   []ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseError   `json:"errors"`
	Messages []ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessage `json:"messages"`
	Result   ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseSuccess `json:"success"`
	JSON    zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseJSON    `json:"-"`
}

// zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseJSON contains the
// JSON metadata for the struct
// [ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponse]
type zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseError struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseErrorJSON `json:"-"`
}

// zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseError]
type zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessage struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessageJSON `json:"-"`
}

// zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessage]
type zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResult struct {
	// Unique WebSocket address that will receive messages from Cloudflare’s edge.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Comma-separated list of fields.
	Fields string `json:"fields"`
	// Filters to drill down into specific events.
	Filter string `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample int64 `json:"sample"`
	// Unique session id of the job.
	SessionID string                                                                  `json:"session_id"`
	JSON      zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResultJSON `json:"-"`
}

// zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResult]
type zoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResultJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseSuccess bool

const (
	ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseSuccessTrue ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsResponseSuccess = true
)

type ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsParams struct {
	// Comma-separated list of fields.
	Fields param.Field[string] `json:"fields"`
	// Filters to drill down into specific events.
	Filter param.Field[string] `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample param.Field[int64] `json:"sample"`
}

func (r ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
