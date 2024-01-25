// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAPIGatewayDiscoveryOperationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneAPIGatewayDiscoveryOperationService] method instead.
type ZoneAPIGatewayDiscoveryOperationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayDiscoveryOperationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayDiscoveryOperationService(opts ...option.RequestOption) (r *ZoneAPIGatewayDiscoveryOperationService) {
	r = &ZoneAPIGatewayDiscoveryOperationService{}
	r.Options = opts
	return
}

// Update the `state` on one or more discovered operations
func (r *ZoneAPIGatewayDiscoveryOperationService) Update(ctx context.Context, zoneID string, body ZoneAPIGatewayDiscoveryOperationUpdateParams, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryOperationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve the most up to date view of discovered operations
func (r *ZoneAPIGatewayDiscoveryOperationService) List(ctx context.Context, zoneID string, query ZoneAPIGatewayDiscoveryOperationListParams, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryOperationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update the `state` on a discovered operation
func (r *ZoneAPIGatewayDiscoveryOperationService) UpdateState(ctx context.Context, zoneID string, operationID string, body ZoneAPIGatewayDiscoveryOperationUpdateStateParams, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryOperationUpdateStateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneAPIGatewayDiscoveryOperationUpdateResponse struct {
	Errors   []ZoneAPIGatewayDiscoveryOperationUpdateResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayDiscoveryOperationUpdateResponseMessage `json:"messages"`
	Result   interface{}                                             `json:"result"`
	// Whether the API call was successful
	Success bool                                               `json:"success"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateResponseJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateResponseJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayDiscoveryOperationUpdateResponse]
type zoneAPIGatewayDiscoveryOperationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationUpdateResponseError struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateResponseErrorJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationUpdateResponseError]
type zoneAPIGatewayDiscoveryOperationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationUpdateResponseMessage struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateResponseMessageJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationUpdateResponseMessage]
type zoneAPIGatewayDiscoveryOperationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationListResponse struct {
	Errors     []ZoneAPIGatewayDiscoveryOperationListResponseError    `json:"errors"`
	Messages   []ZoneAPIGatewayDiscoveryOperationListResponseMessage  `json:"messages"`
	Result     []ZoneAPIGatewayDiscoveryOperationListResponseResult   `json:"result"`
	ResultInfo ZoneAPIGatewayDiscoveryOperationListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                                             `json:"success"`
	JSON    zoneAPIGatewayDiscoveryOperationListResponseJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayDiscoveryOperationListResponse]
type zoneAPIGatewayDiscoveryOperationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationListResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryOperationListResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseErrorJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayDiscoveryOperationListResponseError]
type zoneAPIGatewayDiscoveryOperationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationListResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryOperationListResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationListResponseMessage]
type zoneAPIGatewayDiscoveryOperationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationListResponseResult struct {
	// UUID identifier
	ID string `json:"id" format:"uuid"`
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string                                                     `json:"endpoint" format:"uri-template"`
	Features ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures `json:"features"`
	// RFC3986-compliant host.
	Host        string    `json:"host" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method ZoneAPIGatewayDiscoveryOperationListResponseResultMethod `json:"method"`
	// API discovery engine(s) that discovered this operation
	Origin []ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin `json:"origin"`
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State ZoneAPIGatewayDiscoveryOperationListResponseResultState `json:"state"`
	JSON  zoneAPIGatewayDiscoveryOperationListResponseResultJSON  `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationListResponseResult]
type zoneAPIGatewayDiscoveryOperationListResponseResultJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Features    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	Origin      apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures struct {
	TrafficStats ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats `json:"traffic_stats"`
	JSON         zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON         `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures]
type zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON struct {
	TrafficStats apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats struct {
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The period in seconds these statistics were computed over
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The average number of requests seen during this period
	Requests float64                                                                    `json:"requests,required"`
	JSON     zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats]
type zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON struct {
	LastUpdated   apijson.Field
	PeriodSeconds apijson.Field
	Requests      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to access the endpoint.
type ZoneAPIGatewayDiscoveryOperationListResponseResultMethod string

const (
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodGet     ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "GET"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPost    ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "POST"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodHead    ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "HEAD"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodOptions ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "OPTIONS"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPut     ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "PUT"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodDelete  ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "DELETE"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodConnect ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "CONNECT"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPatch   ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "PATCH"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodTrace   ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "TRACE"
)

//   - `ML` - Discovered operation was sourced using ML API Discovery \*
//     `SessionIdentifier` - Discovered operation was sourced using Session
//     Identifier API Discovery
type ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin string

const (
	ZoneAPIGatewayDiscoveryOperationListResponseResultOriginMl                ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin = "ML"
	ZoneAPIGatewayDiscoveryOperationListResponseResultOriginSessionIdentifier ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin = "SessionIdentifier"
)

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type ZoneAPIGatewayDiscoveryOperationListResponseResultState string

const (
	ZoneAPIGatewayDiscoveryOperationListResponseResultStateReview  ZoneAPIGatewayDiscoveryOperationListResponseResultState = "review"
	ZoneAPIGatewayDiscoveryOperationListResponseResultStateSaved   ZoneAPIGatewayDiscoveryOperationListResponseResultState = "saved"
	ZoneAPIGatewayDiscoveryOperationListResponseResultStateIgnored ZoneAPIGatewayDiscoveryOperationListResponseResultState = "ignored"
)

type ZoneAPIGatewayDiscoveryOperationListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationListResponseResultInfo]
type zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationUpdateStateResponse struct {
	Errors   []ZoneAPIGatewayDiscoveryOperationUpdateStateResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayDiscoveryOperationUpdateStateResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                                    `json:"success"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateStateResponseJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateStateResponseJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationUpdateStateResponse]
type zoneAPIGatewayDiscoveryOperationUpdateStateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateStateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationUpdateStateResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateStateResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateStateResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationUpdateStateResponseError]
type zoneAPIGatewayDiscoveryOperationUpdateStateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateStateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationUpdateStateResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateStateResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateStateResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationUpdateStateResponseMessage]
type zoneAPIGatewayDiscoveryOperationUpdateStateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateStateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResult struct {
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultState `json:"state"`
	JSON  zoneAPIGatewayDiscoveryOperationUpdateStateResponseResultJSON  `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateStateResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResult]
type zoneAPIGatewayDiscoveryOperationUpdateStateResponseResultJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultState string

const (
	ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultStateReview  ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultState = "review"
	ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultStateSaved   ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultState = "saved"
	ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultStateIgnored ZoneAPIGatewayDiscoveryOperationUpdateStateResponseResultState = "ignored"
)

type ZoneAPIGatewayDiscoveryOperationUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneAPIGatewayDiscoveryOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneAPIGatewayDiscoveryOperationListParams struct {
	// When `true`, only return API Discovery results that are not saved into API
	// Shield Endpoint Management
	Diff param.Field[bool] `query:"diff"`
	// Direction to order results.
	Direction param.Field[ZoneAPIGatewayDiscoveryOperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by
	Order param.Field[ZoneAPIGatewayDiscoveryOperationListParamsOrder] `query:"order"`
	// Filter results to only include discovery results sourced from a particular
	// discovery engine
	//
	//   - `ML` - Discovered operations that were sourced using ML API Discovery
	//   - `SessionIdentifier` - Discovered operations that were sourced using Session
	//     Identifier API Discovery
	Origin param.Field[ZoneAPIGatewayDiscoveryOperationListParamsOrigin] `query:"origin"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
	// Filter results to only include discovery results in a particular state. States
	// are as follows
	//
	//   - `review` - Discovered operations that are not saved into API Shield Endpoint
	//     Management
	//   - `saved` - Discovered operations that are already saved into API Shield
	//     Endpoint Management
	//   - `ignored` - Discovered operations that have been marked as ignored
	State param.Field[ZoneAPIGatewayDiscoveryOperationListParamsState] `query:"state"`
}

// URLQuery serializes [ZoneAPIGatewayDiscoveryOperationListParams]'s query
// parameters as `url.Values`.
func (r ZoneAPIGatewayDiscoveryOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type ZoneAPIGatewayDiscoveryOperationListParamsDirection string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsDirectionAsc  ZoneAPIGatewayDiscoveryOperationListParamsDirection = "asc"
	ZoneAPIGatewayDiscoveryOperationListParamsDirectionDesc ZoneAPIGatewayDiscoveryOperationListParamsDirection = "desc"
)

// Field to order by
type ZoneAPIGatewayDiscoveryOperationListParamsOrder string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsOrderHost                    ZoneAPIGatewayDiscoveryOperationListParamsOrder = "host"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderMethod                  ZoneAPIGatewayDiscoveryOperationListParamsOrder = "method"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderEndpoint                ZoneAPIGatewayDiscoveryOperationListParamsOrder = "endpoint"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderTrafficStatsRequests    ZoneAPIGatewayDiscoveryOperationListParamsOrder = "traffic_stats.requests"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderTrafficStatsLastUpdated ZoneAPIGatewayDiscoveryOperationListParamsOrder = "traffic_stats.last_updated"
)

// Filter results to only include discovery results sourced from a particular
// discovery engine
//
//   - `ML` - Discovered operations that were sourced using ML API Discovery
//   - `SessionIdentifier` - Discovered operations that were sourced using Session
//     Identifier API Discovery
type ZoneAPIGatewayDiscoveryOperationListParamsOrigin string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsOriginMl                ZoneAPIGatewayDiscoveryOperationListParamsOrigin = "ML"
	ZoneAPIGatewayDiscoveryOperationListParamsOriginSessionIdentifier ZoneAPIGatewayDiscoveryOperationListParamsOrigin = "SessionIdentifier"
)

// Filter results to only include discovery results in a particular state. States
// are as follows
//
//   - `review` - Discovered operations that are not saved into API Shield Endpoint
//     Management
//   - `saved` - Discovered operations that are already saved into API Shield
//     Endpoint Management
//   - `ignored` - Discovered operations that have been marked as ignored
type ZoneAPIGatewayDiscoveryOperationListParamsState string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsStateReview  ZoneAPIGatewayDiscoveryOperationListParamsState = "review"
	ZoneAPIGatewayDiscoveryOperationListParamsStateSaved   ZoneAPIGatewayDiscoveryOperationListParamsState = "saved"
	ZoneAPIGatewayDiscoveryOperationListParamsStateIgnored ZoneAPIGatewayDiscoveryOperationListParamsState = "ignored"
)

type ZoneAPIGatewayDiscoveryOperationUpdateStateParams struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[ZoneAPIGatewayDiscoveryOperationUpdateStateParamsState] `json:"state"`
}

func (r ZoneAPIGatewayDiscoveryOperationUpdateStateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type ZoneAPIGatewayDiscoveryOperationUpdateStateParamsState string

const (
	ZoneAPIGatewayDiscoveryOperationUpdateStateParamsStateReview  ZoneAPIGatewayDiscoveryOperationUpdateStateParamsState = "review"
	ZoneAPIGatewayDiscoveryOperationUpdateStateParamsStateIgnored ZoneAPIGatewayDiscoveryOperationUpdateStateParamsState = "ignored"
)
