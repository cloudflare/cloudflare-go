// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RouteService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRouteService] method instead.
type RouteService struct {
	Options []option.RequestOption
}

// NewRouteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRouteService(opts ...option.RequestOption) (r *RouteService) {
	r = &RouteService{}
	r.Options = opts
	return
}

// Creates a new Magic static route. Use `?validate_only=true` as an optional query
// parameter to run validation only without persisting changes.
func (r *RouteService) New(ctx context.Context, params RouteNewParams, opts ...option.RequestOption) (res *RouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific Magic static route. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes.
func (r *RouteService) Update(ctx context.Context, routeIdentifier string, params RouteUpdateParams, opts ...option.RequestOption) (res *RouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", params.AccountID, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Magic static routes.
func (r *RouteService) List(ctx context.Context, query RouteListParams, opts ...option.RequestOption) (res *RouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable and remove a specific Magic static route.
func (r *RouteService) Delete(ctx context.Context, routeIdentifier string, params RouteDeleteParams, opts ...option.RequestOption) (res *RouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", params.AccountID, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete multiple Magic static routes.
func (r *RouteService) Empty(ctx context.Context, params RouteEmptyParams, opts ...option.RequestOption) (res *RouteEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteEmptyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Magic static route.
func (r *RouteService) Get(ctx context.Context, routeIdentifier string, query RouteGetParams, opts ...option.RequestOption) (res *RouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", query.AccountID, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ColoName = string

type ColoNameParam = string

type ColoRegion = string

type ColoRegionParam = string

// Used only for ECMP routes.
type Scope struct {
	// List of colo names for the ECMP scope.
	ColoNames []ColoName `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []ColoRegion `json:"colo_regions"`
	JSON        scopeJSON    `json:"-"`
}

// scopeJSON contains the JSON metadata for the struct [Scope]
type scopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Scope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scopeJSON) RawJSON() string {
	return r.raw
}

// Used only for ECMP routes.
type ScopeParam struct {
	// List of colo names for the ECMP scope.
	ColoNames param.Field[[]ColoNameParam] `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions param.Field[[]ColoRegionParam] `json:"colo_regions"`
}

func (r ScopeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteNewResponse struct {
	Routes []RouteNewResponseRoute `json:"routes"`
	JSON   routeNewResponseJSON    `json:"-"`
}

// routeNewResponseJSON contains the JSON metadata for the struct
// [RouteNewResponse]
type routeNewResponseJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseJSON) RawJSON() string {
	return r.raw
}

type RouteNewResponseRoute struct {
	// The next-hop IP Address for the static route.
	Nexthop string `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix string `json:"prefix,required"`
	// Priority of the static route.
	Priority int64 `json:"priority,required"`
	// Identifier
	ID string `json:"id"`
	// When the route was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional human provided description of the static route.
	Description string `json:"description"`
	// When the route was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Used only for ECMP routes.
	Scope Scope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                     `json:"weight"`
	JSON   routeNewResponseRouteJSON `json:"-"`
}

// routeNewResponseRouteJSON contains the JSON metadata for the struct
// [RouteNewResponseRoute]
type routeNewResponseRouteJSON struct {
	Nexthop     apijson.Field
	Prefix      apijson.Field
	Priority    apijson.Field
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Scope       apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseRouteJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponse struct {
	Modified      bool                    `json:"modified"`
	ModifiedRoute interface{}             `json:"modified_route"`
	JSON          routeUpdateResponseJSON `json:"-"`
}

// routeUpdateResponseJSON contains the JSON metadata for the struct
// [RouteUpdateResponse]
type routeUpdateResponseJSON struct {
	Modified      apijson.Field
	ModifiedRoute apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RouteListResponse struct {
	Routes []RouteListResponseRoute `json:"routes"`
	JSON   routeListResponseJSON    `json:"-"`
}

// routeListResponseJSON contains the JSON metadata for the struct
// [RouteListResponse]
type routeListResponseJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeListResponseJSON) RawJSON() string {
	return r.raw
}

type RouteListResponseRoute struct {
	// The next-hop IP Address for the static route.
	Nexthop string `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix string `json:"prefix,required"`
	// Priority of the static route.
	Priority int64 `json:"priority,required"`
	// Identifier
	ID string `json:"id"`
	// When the route was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional human provided description of the static route.
	Description string `json:"description"`
	// When the route was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Used only for ECMP routes.
	Scope Scope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                      `json:"weight"`
	JSON   routeListResponseRouteJSON `json:"-"`
}

// routeListResponseRouteJSON contains the JSON metadata for the struct
// [RouteListResponseRoute]
type routeListResponseRouteJSON struct {
	Nexthop     apijson.Field
	Prefix      apijson.Field
	Priority    apijson.Field
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Scope       apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteListResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeListResponseRouteJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponse struct {
	Deleted      bool                    `json:"deleted"`
	DeletedRoute interface{}             `json:"deleted_route"`
	JSON         routeDeleteResponseJSON `json:"-"`
}

// routeDeleteResponseJSON contains the JSON metadata for the struct
// [RouteDeleteResponse]
type routeDeleteResponseJSON struct {
	Deleted      apijson.Field
	DeletedRoute apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RouteEmptyResponse struct {
	Deleted       bool                   `json:"deleted"`
	DeletedRoutes interface{}            `json:"deleted_routes"`
	JSON          routeEmptyResponseJSON `json:"-"`
}

// routeEmptyResponseJSON contains the JSON metadata for the struct
// [RouteEmptyResponse]
type routeEmptyResponseJSON struct {
	Deleted       apijson.Field
	DeletedRoutes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RouteEmptyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeEmptyResponseJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponse struct {
	Route interface{}          `json:"route"`
	JSON  routeGetResponseJSON `json:"-"`
}

// routeGetResponseJSON contains the JSON metadata for the struct
// [RouteGetResponse]
type routeGetResponseJSON struct {
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseJSON) RawJSON() string {
	return r.raw
}

type RouteNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r RouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RouteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteNewResponse      `json:"result,required"`
	// Whether the API call was successful
	Success RouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeNewResponseEnvelopeJSON    `json:"-"`
}

// routeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelope]
type routeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteNewResponseEnvelopeSuccess bool

const (
	RouteNewResponseEnvelopeSuccessTrue RouteNewResponseEnvelopeSuccess = true
)

func (r RouteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The next-hop IP Address for the static route.
	Nexthop param.Field[string] `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix param.Field[string] `json:"prefix,required"`
	// Priority of the static route.
	Priority param.Field[int64] `json:"priority,required"`
	// An optional human provided description of the static route.
	Description param.Field[string] `json:"description"`
	// Used only for ECMP routes.
	Scope param.Field[ScopeParam] `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight param.Field[int64] `json:"weight"`
}

func (r RouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteUpdateResponse   `json:"result,required"`
	// Whether the API call was successful
	Success RouteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeUpdateResponseEnvelopeJSON    `json:"-"`
}

// routeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteUpdateResponseEnvelope]
type routeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteUpdateResponseEnvelopeSuccess bool

const (
	RouteUpdateResponseEnvelopeSuccessTrue RouteUpdateResponseEnvelopeSuccess = true
)

func (r RouteUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type RouteListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteListResponse     `json:"result,required"`
	// Whether the API call was successful
	Success RouteListResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeListResponseEnvelopeJSON    `json:"-"`
}

// routeListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteListResponseEnvelope]
type routeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteListResponseEnvelopeSuccess bool

const (
	RouteListResponseEnvelopeSuccessTrue RouteListResponseEnvelopeSuccess = true
)

func (r RouteListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r RouteDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RouteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteDeleteResponse   `json:"result,required"`
	// Whether the API call was successful
	Success RouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeDeleteResponseEnvelopeJSON    `json:"-"`
}

// routeDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteDeleteResponseEnvelope]
type routeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteDeleteResponseEnvelopeSuccess bool

const (
	RouteDeleteResponseEnvelopeSuccessTrue RouteDeleteResponseEnvelopeSuccess = true
)

func (r RouteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteEmptyParams struct {
	// Identifier
	AccountID param.Field[string]                  `path:"account_id,required"`
	Routes    param.Field[[]RouteEmptyParamsRoute] `json:"routes,required"`
}

func (r RouteEmptyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteEmptyParamsRoute struct {
}

func (r RouteEmptyParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteEmptyResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteEmptyResponse    `json:"result,required"`
	// Whether the API call was successful
	Success RouteEmptyResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeEmptyResponseEnvelopeJSON    `json:"-"`
}

// routeEmptyResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteEmptyResponseEnvelope]
type routeEmptyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteEmptyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeEmptyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteEmptyResponseEnvelopeSuccess bool

const (
	RouteEmptyResponseEnvelopeSuccessTrue RouteEmptyResponseEnvelopeSuccess = true
)

func (r RouteEmptyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteEmptyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type RouteGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RouteGetResponse      `json:"result,required"`
	// Whether the API call was successful
	Success RouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeGetResponseEnvelopeJSON    `json:"-"`
}

// routeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelope]
type routeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteGetResponseEnvelopeSuccess bool

const (
	RouteGetResponseEnvelopeSuccessTrue RouteGetResponseEnvelopeSuccess = true
)

func (r RouteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
