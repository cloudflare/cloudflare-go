// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMagicRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMagicRouteService] method
// instead.
type AccountMagicRouteService struct {
	Options []option.RequestOption
}

// NewAccountMagicRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicRouteService(opts ...option.RequestOption) (r *AccountMagicRouteService) {
	r = &AccountMagicRouteService{}
	r.Options = opts
	return
}

// Get a specific Magic static route.
func (r *AccountMagicRouteService) Get(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *RouteSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Magic static route. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes.
func (r *AccountMagicRouteService) Update(ctx context.Context, accountIdentifier string, routeIdentifier string, body AccountMagicRouteUpdateParams, opts ...option.RequestOption) (res *RouteModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disable and remove a specific Magic static route.
func (r *AccountMagicRouteService) Delete(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *RouteDeletedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Magic static route. Use `?validate_only=true` as an optional query
// parameter to run validation only without persisting changes.
func (r *AccountMagicRouteService) MagicStaticRoutesNewRoutes(ctx context.Context, accountIdentifier string, body AccountMagicRouteMagicStaticRoutesNewRoutesParams, opts ...option.RequestOption) (res *RoutesCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all Magic static routes.
func (r *AccountMagicRouteService) MagicStaticRoutesListRoutes(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RoutesCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update multiple Magic static routes. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes. Only fields
// for a route that need to be changed need be provided.
func (r *AccountMagicRouteService) MagicStaticRoutesUpdateManyRoutes(ctx context.Context, accountIdentifier string, body AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParams, opts ...option.RequestOption) (res *MultipleRouteModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type MultipleRouteModifiedResponse struct {
	Errors   []MultipleRouteModifiedResponseError   `json:"errors"`
	Messages []MultipleRouteModifiedResponseMessage `json:"messages"`
	Result   MultipleRouteModifiedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success MultipleRouteModifiedResponseSuccess `json:"success"`
	JSON    multipleRouteModifiedResponseJSON    `json:"-"`
}

// multipleRouteModifiedResponseJSON contains the JSON metadata for the struct
// [MultipleRouteModifiedResponse]
type multipleRouteModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultipleRouteModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MultipleRouteModifiedResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    multipleRouteModifiedResponseErrorJSON `json:"-"`
}

// multipleRouteModifiedResponseErrorJSON contains the JSON metadata for the struct
// [MultipleRouteModifiedResponseError]
type multipleRouteModifiedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultipleRouteModifiedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MultipleRouteModifiedResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    multipleRouteModifiedResponseMessageJSON `json:"-"`
}

// multipleRouteModifiedResponseMessageJSON contains the JSON metadata for the
// struct [MultipleRouteModifiedResponseMessage]
type multipleRouteModifiedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultipleRouteModifiedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MultipleRouteModifiedResponseResult struct {
	Modified       bool                                               `json:"modified"`
	ModifiedRoutes []MultipleRouteModifiedResponseResultModifiedRoute `json:"modified_routes"`
	JSON           multipleRouteModifiedResponseResultJSON            `json:"-"`
}

// multipleRouteModifiedResponseResultJSON contains the JSON metadata for the
// struct [MultipleRouteModifiedResponseResult]
type multipleRouteModifiedResponseResultJSON struct {
	Modified       apijson.Field
	ModifiedRoutes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MultipleRouteModifiedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MultipleRouteModifiedResponseResultModifiedRoute struct {
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
	Scope MultipleRouteModifiedResponseResultModifiedRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                                `json:"weight"`
	JSON   multipleRouteModifiedResponseResultModifiedRouteJSON `json:"-"`
}

// multipleRouteModifiedResponseResultModifiedRouteJSON contains the JSON metadata
// for the struct [MultipleRouteModifiedResponseResultModifiedRoute]
type multipleRouteModifiedResponseResultModifiedRouteJSON struct {
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

func (r *MultipleRouteModifiedResponseResultModifiedRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type MultipleRouteModifiedResponseResultModifiedRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                                   `json:"colo_regions"`
	JSON        multipleRouteModifiedResponseResultModifiedRoutesScopeJSON `json:"-"`
}

// multipleRouteModifiedResponseResultModifiedRoutesScopeJSON contains the JSON
// metadata for the struct [MultipleRouteModifiedResponseResultModifiedRoutesScope]
type multipleRouteModifiedResponseResultModifiedRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultipleRouteModifiedResponseResultModifiedRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MultipleRouteModifiedResponseSuccess bool

const (
	MultipleRouteModifiedResponseSuccessTrue MultipleRouteModifiedResponseSuccess = true
)

type RouteDeletedResponse struct {
	Errors   []RouteDeletedResponseError   `json:"errors"`
	Messages []RouteDeletedResponseMessage `json:"messages"`
	Result   RouteDeletedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RouteDeletedResponseSuccess `json:"success"`
	JSON    routeDeletedResponseJSON    `json:"-"`
}

// routeDeletedResponseJSON contains the JSON metadata for the struct
// [RouteDeletedResponse]
type routeDeletedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeletedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteDeletedResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    routeDeletedResponseErrorJSON `json:"-"`
}

// routeDeletedResponseErrorJSON contains the JSON metadata for the struct
// [RouteDeletedResponseError]
type routeDeletedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeletedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteDeletedResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    routeDeletedResponseMessageJSON `json:"-"`
}

// routeDeletedResponseMessageJSON contains the JSON metadata for the struct
// [RouteDeletedResponseMessage]
type routeDeletedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeletedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteDeletedResponseResult struct {
	Deleted      bool                           `json:"deleted"`
	DeletedRoute interface{}                    `json:"deleted_route"`
	JSON         routeDeletedResponseResultJSON `json:"-"`
}

// routeDeletedResponseResultJSON contains the JSON metadata for the struct
// [RouteDeletedResponseResult]
type routeDeletedResponseResultJSON struct {
	Deleted      apijson.Field
	DeletedRoute apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RouteDeletedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RouteDeletedResponseSuccess bool

const (
	RouteDeletedResponseSuccessTrue RouteDeletedResponseSuccess = true
)

type RouteModifiedResponse struct {
	Errors   []RouteModifiedResponseError   `json:"errors"`
	Messages []RouteModifiedResponseMessage `json:"messages"`
	Result   RouteModifiedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RouteModifiedResponseSuccess `json:"success"`
	JSON    routeModifiedResponseJSON    `json:"-"`
}

// routeModifiedResponseJSON contains the JSON metadata for the struct
// [RouteModifiedResponse]
type routeModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteModifiedResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    routeModifiedResponseErrorJSON `json:"-"`
}

// routeModifiedResponseErrorJSON contains the JSON metadata for the struct
// [RouteModifiedResponseError]
type routeModifiedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteModifiedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteModifiedResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    routeModifiedResponseMessageJSON `json:"-"`
}

// routeModifiedResponseMessageJSON contains the JSON metadata for the struct
// [RouteModifiedResponseMessage]
type routeModifiedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteModifiedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteModifiedResponseResult struct {
	Modified      bool                            `json:"modified"`
	ModifiedRoute interface{}                     `json:"modified_route"`
	JSON          routeModifiedResponseResultJSON `json:"-"`
}

// routeModifiedResponseResultJSON contains the JSON metadata for the struct
// [RouteModifiedResponseResult]
type routeModifiedResponseResultJSON struct {
	Modified      apijson.Field
	ModifiedRoute apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RouteModifiedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RouteModifiedResponseSuccess bool

const (
	RouteModifiedResponseSuccessTrue RouteModifiedResponseSuccess = true
)

type RouteSingleResponse struct {
	Errors   []RouteSingleResponseError   `json:"errors"`
	Messages []RouteSingleResponseMessage `json:"messages"`
	Result   RouteSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RouteSingleResponseSuccess `json:"success"`
	JSON    routeSingleResponseJSON    `json:"-"`
}

// routeSingleResponseJSON contains the JSON metadata for the struct
// [RouteSingleResponse]
type routeSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteSingleResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    routeSingleResponseErrorJSON `json:"-"`
}

// routeSingleResponseErrorJSON contains the JSON metadata for the struct
// [RouteSingleResponseError]
type routeSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteSingleResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    routeSingleResponseMessageJSON `json:"-"`
}

// routeSingleResponseMessageJSON contains the JSON metadata for the struct
// [RouteSingleResponseMessage]
type routeSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteSingleResponseResult struct {
	Route interface{}                   `json:"route"`
	JSON  routeSingleResponseResultJSON `json:"-"`
}

// routeSingleResponseResultJSON contains the JSON metadata for the struct
// [RouteSingleResponseResult]
type routeSingleResponseResultJSON struct {
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RouteSingleResponseSuccess bool

const (
	RouteSingleResponseSuccessTrue RouteSingleResponseSuccess = true
)

type RoutesCollectionResponse struct {
	Errors   []RoutesCollectionResponseError   `json:"errors"`
	Messages []RoutesCollectionResponseMessage `json:"messages"`
	Result   RoutesCollectionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RoutesCollectionResponseSuccess `json:"success"`
	JSON    routesCollectionResponseJSON    `json:"-"`
}

// routesCollectionResponseJSON contains the JSON metadata for the struct
// [RoutesCollectionResponse]
type routesCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutesCollectionResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    routesCollectionResponseErrorJSON `json:"-"`
}

// routesCollectionResponseErrorJSON contains the JSON metadata for the struct
// [RoutesCollectionResponseError]
type routesCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutesCollectionResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    routesCollectionResponseMessageJSON `json:"-"`
}

// routesCollectionResponseMessageJSON contains the JSON metadata for the struct
// [RoutesCollectionResponseMessage]
type routesCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutesCollectionResponseResult struct {
	Routes []RoutesCollectionResponseResultRoute `json:"routes"`
	JSON   routesCollectionResponseResultJSON    `json:"-"`
}

// routesCollectionResponseResultJSON contains the JSON metadata for the struct
// [RoutesCollectionResponseResult]
type routesCollectionResponseResultJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutesCollectionResponseResultRoute struct {
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
	Scope RoutesCollectionResponseResultRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                   `json:"weight"`
	JSON   routesCollectionResponseResultRouteJSON `json:"-"`
}

// routesCollectionResponseResultRouteJSON contains the JSON metadata for the
// struct [RoutesCollectionResponseResultRoute]
type routesCollectionResponseResultRouteJSON struct {
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

func (r *RoutesCollectionResponseResultRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type RoutesCollectionResponseResultRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                      `json:"colo_regions"`
	JSON        routesCollectionResponseResultRoutesScopeJSON `json:"-"`
}

// routesCollectionResponseResultRoutesScopeJSON contains the JSON metadata for the
// struct [RoutesCollectionResponseResultRoutesScope]
type routesCollectionResponseResultRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesCollectionResponseResultRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RoutesCollectionResponseSuccess bool

const (
	RoutesCollectionResponseSuccessTrue RoutesCollectionResponseSuccess = true
)

type AccountMagicRouteUpdateParams struct {
	// The next-hop IP Address for the static route.
	Nexthop param.Field[string] `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix param.Field[string] `json:"prefix,required"`
	// Priority of the static route.
	Priority param.Field[int64] `json:"priority,required"`
	// An optional human provided description of the static route.
	Description param.Field[string] `json:"description"`
	// Used only for ECMP routes.
	Scope param.Field[AccountMagicRouteUpdateParamsScope] `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight param.Field[int64] `json:"weight"`
}

func (r AccountMagicRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Used only for ECMP routes.
type AccountMagicRouteUpdateParamsScope struct {
	// List of colo names for the ECMP scope.
	ColoNames param.Field[[]string] `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions param.Field[[]string] `json:"colo_regions"`
}

func (r AccountMagicRouteUpdateParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteMagicStaticRoutesNewRoutesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountMagicRouteMagicStaticRoutesNewRoutesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParams struct {
	Routes param.Field[[]AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoute] `json:"routes,required"`
}

func (r AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoute struct {
	// An optional human provided description of the static route.
	Description param.Field[string] `json:"description"`
	// The next-hop IP Address for the static route.
	Nexthop param.Field[string] `json:"nexthop"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix param.Field[string] `json:"prefix"`
	// Priority of the static route.
	Priority param.Field[int64] `json:"priority"`
	// Used only for ECMP routes.
	Scope param.Field[AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoutesScope] `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight param.Field[int64] `json:"weight"`
}

func (r AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Used only for ECMP routes.
type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames param.Field[[]string] `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions param.Field[[]string] `json:"colo_regions"`
}

func (r AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoutesScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
