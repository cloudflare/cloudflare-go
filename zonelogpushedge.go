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
func (r *ZoneLogpushEdgeService) GetZonesZoneIdentifierLogpushEdgeJobs(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *InstantLogsJobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new Instant Logs job for a zone.
func (r *ZoneLogpushEdgeService) PostZonesZoneIdentifierLogpushEdgeJobs(ctx context.Context, zoneIdentifier string, body ZoneLogpushEdgePostZonesZoneIdentifierLogpushEdgeJobsParams, opts ...option.RequestOption) (res *InstantLogsJobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InstantLogsJobResponseCollection struct {
	Errors   []InstantLogsJobResponseCollectionError   `json:"errors"`
	Messages []InstantLogsJobResponseCollectionMessage `json:"messages"`
	Result   []InstantLogsJobResponseCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success InstantLogsJobResponseCollectionSuccess `json:"success"`
	JSON    instantLogsJobResponseCollectionJSON    `json:"-"`
}

// instantLogsJobResponseCollectionJSON contains the JSON metadata for the struct
// [InstantLogsJobResponseCollection]
type instantLogsJobResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstantLogsJobResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InstantLogsJobResponseCollectionError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    instantLogsJobResponseCollectionErrorJSON `json:"-"`
}

// instantLogsJobResponseCollectionErrorJSON contains the JSON metadata for the
// struct [InstantLogsJobResponseCollectionError]
type instantLogsJobResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstantLogsJobResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InstantLogsJobResponseCollectionMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    instantLogsJobResponseCollectionMessageJSON `json:"-"`
}

// instantLogsJobResponseCollectionMessageJSON contains the JSON metadata for the
// struct [InstantLogsJobResponseCollectionMessage]
type instantLogsJobResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstantLogsJobResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InstantLogsJobResponseCollectionResult struct {
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
	SessionID string                                     `json:"session_id"`
	JSON      instantLogsJobResponseCollectionResultJSON `json:"-"`
}

// instantLogsJobResponseCollectionResultJSON contains the JSON metadata for the
// struct [InstantLogsJobResponseCollectionResult]
type instantLogsJobResponseCollectionResultJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstantLogsJobResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type InstantLogsJobResponseCollectionSuccess bool

const (
	InstantLogsJobResponseCollectionSuccessTrue InstantLogsJobResponseCollectionSuccess = true
)

type InstantLogsJobResponseSingle struct {
	Errors   []InstantLogsJobResponseSingleError   `json:"errors"`
	Messages []InstantLogsJobResponseSingleMessage `json:"messages"`
	Result   InstantLogsJobResponseSingleResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success InstantLogsJobResponseSingleSuccess `json:"success"`
	JSON    instantLogsJobResponseSingleJSON    `json:"-"`
}

// instantLogsJobResponseSingleJSON contains the JSON metadata for the struct
// [InstantLogsJobResponseSingle]
type instantLogsJobResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstantLogsJobResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InstantLogsJobResponseSingleError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    instantLogsJobResponseSingleErrorJSON `json:"-"`
}

// instantLogsJobResponseSingleErrorJSON contains the JSON metadata for the struct
// [InstantLogsJobResponseSingleError]
type instantLogsJobResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstantLogsJobResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InstantLogsJobResponseSingleMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    instantLogsJobResponseSingleMessageJSON `json:"-"`
}

// instantLogsJobResponseSingleMessageJSON contains the JSON metadata for the
// struct [InstantLogsJobResponseSingleMessage]
type instantLogsJobResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstantLogsJobResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InstantLogsJobResponseSingleResult struct {
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
	SessionID string                                 `json:"session_id"`
	JSON      instantLogsJobResponseSingleResultJSON `json:"-"`
}

// instantLogsJobResponseSingleResultJSON contains the JSON metadata for the struct
// [InstantLogsJobResponseSingleResult]
type instantLogsJobResponseSingleResultJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstantLogsJobResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type InstantLogsJobResponseSingleSuccess bool

const (
	InstantLogsJobResponseSingleSuccessTrue InstantLogsJobResponseSingleSuccess = true
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
