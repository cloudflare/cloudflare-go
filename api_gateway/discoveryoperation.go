// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
)

// DiscoveryOperationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiscoveryOperationService] method instead.
type DiscoveryOperationService struct {
	Options []option.RequestOption
}

// NewDiscoveryOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDiscoveryOperationService(opts ...option.RequestOption) (r *DiscoveryOperationService) {
	r = &DiscoveryOperationService{}
	r.Options = opts
	return
}

// Retrieve the most up to date view of discovered operations
func (r *DiscoveryOperationService) List(ctx context.Context, params DiscoveryOperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DiscoveryOperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieve the most up to date view of discovered operations
func (r *DiscoveryOperationService) ListAutoPaging(ctx context.Context, params DiscoveryOperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DiscoveryOperationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Update the `state` on one or more discovered operations
func (r *DiscoveryOperationService) BulkEdit(ctx context.Context, params DiscoveryOperationBulkEditParams, opts ...option.RequestOption) (res *DiscoveryOperationBulkEditResponse, err error) {
	var env DiscoveryOperationBulkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the `state` on a discovered operation
func (r *DiscoveryOperationService) Edit(ctx context.Context, operationID string, params DiscoveryOperationEditParams, opts ...option.RequestOption) (res *DiscoveryOperationEditResponse, err error) {
	var env DiscoveryOperationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations/%s", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DiscoveryOperationListResponse struct {
	// UUID
	ID string `json:"id,required"`
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method DiscoveryOperationListResponseMethod `json:"method,required"`
	// API discovery engine(s) that discovered this operation
	Origin []DiscoveryOperationListResponseOrigin `json:"origin,required"`
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State    DiscoveryOperationListResponseState    `json:"state,required"`
	Features DiscoveryOperationListResponseFeatures `json:"features"`
	JSON     discoveryOperationListResponseJSON     `json:"-"`
}

// discoveryOperationListResponseJSON contains the JSON metadata for the struct
// [DiscoveryOperationListResponse]
type discoveryOperationListResponseJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	Origin      apijson.Field
	State       apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationListResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type DiscoveryOperationListResponseMethod string

const (
	DiscoveryOperationListResponseMethodGet     DiscoveryOperationListResponseMethod = "GET"
	DiscoveryOperationListResponseMethodPost    DiscoveryOperationListResponseMethod = "POST"
	DiscoveryOperationListResponseMethodHead    DiscoveryOperationListResponseMethod = "HEAD"
	DiscoveryOperationListResponseMethodOptions DiscoveryOperationListResponseMethod = "OPTIONS"
	DiscoveryOperationListResponseMethodPut     DiscoveryOperationListResponseMethod = "PUT"
	DiscoveryOperationListResponseMethodDelete  DiscoveryOperationListResponseMethod = "DELETE"
	DiscoveryOperationListResponseMethodConnect DiscoveryOperationListResponseMethod = "CONNECT"
	DiscoveryOperationListResponseMethodPatch   DiscoveryOperationListResponseMethod = "PATCH"
	DiscoveryOperationListResponseMethodTrace   DiscoveryOperationListResponseMethod = "TRACE"
)

func (r DiscoveryOperationListResponseMethod) IsKnown() bool {
	switch r {
	case DiscoveryOperationListResponseMethodGet, DiscoveryOperationListResponseMethodPost, DiscoveryOperationListResponseMethodHead, DiscoveryOperationListResponseMethodOptions, DiscoveryOperationListResponseMethodPut, DiscoveryOperationListResponseMethodDelete, DiscoveryOperationListResponseMethodConnect, DiscoveryOperationListResponseMethodPatch, DiscoveryOperationListResponseMethodTrace:
		return true
	}
	return false
}

//   - `ML` - Discovered operation was sourced using ML API Discovery \*
//     `SessionIdentifier` - Discovered operation was sourced using Session
//     Identifier API Discovery
type DiscoveryOperationListResponseOrigin string

const (
	DiscoveryOperationListResponseOriginMl                DiscoveryOperationListResponseOrigin = "ML"
	DiscoveryOperationListResponseOriginSessionIdentifier DiscoveryOperationListResponseOrigin = "SessionIdentifier"
)

func (r DiscoveryOperationListResponseOrigin) IsKnown() bool {
	switch r {
	case DiscoveryOperationListResponseOriginMl, DiscoveryOperationListResponseOriginSessionIdentifier:
		return true
	}
	return false
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type DiscoveryOperationListResponseState string

const (
	DiscoveryOperationListResponseStateReview  DiscoveryOperationListResponseState = "review"
	DiscoveryOperationListResponseStateSaved   DiscoveryOperationListResponseState = "saved"
	DiscoveryOperationListResponseStateIgnored DiscoveryOperationListResponseState = "ignored"
)

func (r DiscoveryOperationListResponseState) IsKnown() bool {
	switch r {
	case DiscoveryOperationListResponseStateReview, DiscoveryOperationListResponseStateSaved, DiscoveryOperationListResponseStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationListResponseFeatures struct {
	TrafficStats DiscoveryOperationListResponseFeaturesTrafficStats `json:"traffic_stats"`
	JSON         discoveryOperationListResponseFeaturesJSON         `json:"-"`
}

// discoveryOperationListResponseFeaturesJSON contains the JSON metadata for the
// struct [DiscoveryOperationListResponseFeatures]
type discoveryOperationListResponseFeaturesJSON struct {
	TrafficStats apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DiscoveryOperationListResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationListResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

type DiscoveryOperationListResponseFeaturesTrafficStats struct {
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The period in seconds these statistics were computed over
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The average number of requests seen during this period
	Requests float64                                                `json:"requests,required"`
	JSON     discoveryOperationListResponseFeaturesTrafficStatsJSON `json:"-"`
}

// discoveryOperationListResponseFeaturesTrafficStatsJSON contains the JSON
// metadata for the struct [DiscoveryOperationListResponseFeaturesTrafficStats]
type discoveryOperationListResponseFeaturesTrafficStatsJSON struct {
	LastUpdated   apijson.Field
	PeriodSeconds apijson.Field
	Requests      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DiscoveryOperationListResponseFeaturesTrafficStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationListResponseFeaturesTrafficStatsJSON) RawJSON() string {
	return r.raw
}

type DiscoveryOperationBulkEditResponse map[string]DiscoveryOperationBulkEditResponseItem

// Mappings of discovered operations (keys) to objects describing their state
type DiscoveryOperationBulkEditResponseItem struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State DiscoveryOperationBulkEditResponseItemState `json:"state"`
	JSON  discoveryOperationBulkEditResponseItemJSON  `json:"-"`
}

// discoveryOperationBulkEditResponseItemJSON contains the JSON metadata for the
// struct [DiscoveryOperationBulkEditResponseItem]
type discoveryOperationBulkEditResponseItemJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationBulkEditResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationBulkEditResponseItemJSON) RawJSON() string {
	return r.raw
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type DiscoveryOperationBulkEditResponseItemState string

const (
	DiscoveryOperationBulkEditResponseItemStateReview  DiscoveryOperationBulkEditResponseItemState = "review"
	DiscoveryOperationBulkEditResponseItemStateIgnored DiscoveryOperationBulkEditResponseItemState = "ignored"
)

func (r DiscoveryOperationBulkEditResponseItemState) IsKnown() bool {
	switch r {
	case DiscoveryOperationBulkEditResponseItemStateReview, DiscoveryOperationBulkEditResponseItemStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationEditResponse struct {
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State DiscoveryOperationEditResponseState `json:"state"`
	JSON  discoveryOperationEditResponseJSON  `json:"-"`
}

// discoveryOperationEditResponseJSON contains the JSON metadata for the struct
// [DiscoveryOperationEditResponse]
type discoveryOperationEditResponseJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationEditResponseJSON) RawJSON() string {
	return r.raw
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type DiscoveryOperationEditResponseState string

const (
	DiscoveryOperationEditResponseStateReview  DiscoveryOperationEditResponseState = "review"
	DiscoveryOperationEditResponseStateSaved   DiscoveryOperationEditResponseState = "saved"
	DiscoveryOperationEditResponseStateIgnored DiscoveryOperationEditResponseState = "ignored"
)

func (r DiscoveryOperationEditResponseState) IsKnown() bool {
	switch r {
	case DiscoveryOperationEditResponseStateReview, DiscoveryOperationEditResponseStateSaved, DiscoveryOperationEditResponseStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// When `true`, only return API Discovery results that are not saved into API
	// Shield Endpoint Management
	Diff param.Field[bool] `query:"diff"`
	// Direction to order results.
	Direction param.Field[DiscoveryOperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by
	Order param.Field[DiscoveryOperationListParamsOrder] `query:"order"`
	// Filter results to only include discovery results sourced from a particular
	// discovery engine
	//
	//   - `ML` - Discovered operations that were sourced using ML API Discovery
	//   - `SessionIdentifier` - Discovered operations that were sourced using Session
	//     Identifier API Discovery
	Origin param.Field[DiscoveryOperationListParamsOrigin] `query:"origin"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter results to only include discovery results in a particular state. States
	// are as follows
	//
	//   - `review` - Discovered operations that are not saved into API Shield Endpoint
	//     Management
	//   - `saved` - Discovered operations that are already saved into API Shield
	//     Endpoint Management
	//   - `ignored` - Discovered operations that have been marked as ignored
	State param.Field[DiscoveryOperationListParamsState] `query:"state"`
}

// URLQuery serializes [DiscoveryOperationListParams]'s query parameters as
// `url.Values`.
func (r DiscoveryOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type DiscoveryOperationListParamsDirection string

const (
	DiscoveryOperationListParamsDirectionAsc  DiscoveryOperationListParamsDirection = "asc"
	DiscoveryOperationListParamsDirectionDesc DiscoveryOperationListParamsDirection = "desc"
)

func (r DiscoveryOperationListParamsDirection) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsDirectionAsc, DiscoveryOperationListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order by
type DiscoveryOperationListParamsOrder string

const (
	DiscoveryOperationListParamsOrderHost                    DiscoveryOperationListParamsOrder = "host"
	DiscoveryOperationListParamsOrderMethod                  DiscoveryOperationListParamsOrder = "method"
	DiscoveryOperationListParamsOrderEndpoint                DiscoveryOperationListParamsOrder = "endpoint"
	DiscoveryOperationListParamsOrderTrafficStatsRequests    DiscoveryOperationListParamsOrder = "traffic_stats.requests"
	DiscoveryOperationListParamsOrderTrafficStatsLastUpdated DiscoveryOperationListParamsOrder = "traffic_stats.last_updated"
)

func (r DiscoveryOperationListParamsOrder) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsOrderHost, DiscoveryOperationListParamsOrderMethod, DiscoveryOperationListParamsOrderEndpoint, DiscoveryOperationListParamsOrderTrafficStatsRequests, DiscoveryOperationListParamsOrderTrafficStatsLastUpdated:
		return true
	}
	return false
}

// Filter results to only include discovery results sourced from a particular
// discovery engine
//
//   - `ML` - Discovered operations that were sourced using ML API Discovery
//   - `SessionIdentifier` - Discovered operations that were sourced using Session
//     Identifier API Discovery
type DiscoveryOperationListParamsOrigin string

const (
	DiscoveryOperationListParamsOriginMl                DiscoveryOperationListParamsOrigin = "ML"
	DiscoveryOperationListParamsOriginSessionIdentifier DiscoveryOperationListParamsOrigin = "SessionIdentifier"
)

func (r DiscoveryOperationListParamsOrigin) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsOriginMl, DiscoveryOperationListParamsOriginSessionIdentifier:
		return true
	}
	return false
}

// Filter results to only include discovery results in a particular state. States
// are as follows
//
//   - `review` - Discovered operations that are not saved into API Shield Endpoint
//     Management
//   - `saved` - Discovered operations that are already saved into API Shield
//     Endpoint Management
//   - `ignored` - Discovered operations that have been marked as ignored
type DiscoveryOperationListParamsState string

const (
	DiscoveryOperationListParamsStateReview  DiscoveryOperationListParamsState = "review"
	DiscoveryOperationListParamsStateSaved   DiscoveryOperationListParamsState = "saved"
	DiscoveryOperationListParamsStateIgnored DiscoveryOperationListParamsState = "ignored"
)

func (r DiscoveryOperationListParamsState) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsStateReview, DiscoveryOperationListParamsStateSaved, DiscoveryOperationListParamsStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationBulkEditParams struct {
	// Identifier
	ZoneID param.Field[string]                             `path:"zone_id,required"`
	Body   map[string]DiscoveryOperationBulkEditParamsBody `json:"body,required"`
}

func (r DiscoveryOperationBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Mappings of discovered operations (keys) to objects describing their state
type DiscoveryOperationBulkEditParamsBody struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[DiscoveryOperationBulkEditParamsBodyState] `json:"state"`
}

func (r DiscoveryOperationBulkEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type DiscoveryOperationBulkEditParamsBodyState string

const (
	DiscoveryOperationBulkEditParamsBodyStateReview  DiscoveryOperationBulkEditParamsBodyState = "review"
	DiscoveryOperationBulkEditParamsBodyStateIgnored DiscoveryOperationBulkEditParamsBodyState = "ignored"
)

func (r DiscoveryOperationBulkEditParamsBodyState) IsKnown() bool {
	switch r {
	case DiscoveryOperationBulkEditParamsBodyStateReview, DiscoveryOperationBulkEditParamsBodyStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationBulkEditResponseEnvelope struct {
	Errors   Message                            `json:"errors,required"`
	Messages Message                            `json:"messages,required"`
	Result   DiscoveryOperationBulkEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success DiscoveryOperationBulkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    discoveryOperationBulkEditResponseEnvelopeJSON    `json:"-"`
}

// discoveryOperationBulkEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [DiscoveryOperationBulkEditResponseEnvelope]
type discoveryOperationBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DiscoveryOperationBulkEditResponseEnvelopeSuccess bool

const (
	DiscoveryOperationBulkEditResponseEnvelopeSuccessTrue DiscoveryOperationBulkEditResponseEnvelopeSuccess = true
)

func (r DiscoveryOperationBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DiscoveryOperationBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DiscoveryOperationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[DiscoveryOperationEditParamsState] `json:"state"`
}

func (r DiscoveryOperationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type DiscoveryOperationEditParamsState string

const (
	DiscoveryOperationEditParamsStateReview  DiscoveryOperationEditParamsState = "review"
	DiscoveryOperationEditParamsStateIgnored DiscoveryOperationEditParamsState = "ignored"
)

func (r DiscoveryOperationEditParamsState) IsKnown() bool {
	switch r {
	case DiscoveryOperationEditParamsStateReview, DiscoveryOperationEditParamsStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationEditResponseEnvelope struct {
	Errors   Message                        `json:"errors,required"`
	Messages Message                        `json:"messages,required"`
	Result   DiscoveryOperationEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success DiscoveryOperationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    discoveryOperationEditResponseEnvelopeJSON    `json:"-"`
}

// discoveryOperationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DiscoveryOperationEditResponseEnvelope]
type discoveryOperationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DiscoveryOperationEditResponseEnvelopeSuccess bool

const (
	DiscoveryOperationEditResponseEnvelopeSuccessTrue DiscoveryOperationEditResponseEnvelopeSuccess = true
)

func (r DiscoveryOperationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DiscoveryOperationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
