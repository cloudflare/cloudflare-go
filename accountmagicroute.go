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
func (r *AccountMagicRouteService) Get(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *AccountMagicRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Magic static route. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes.
func (r *AccountMagicRouteService) Update(ctx context.Context, accountIdentifier string, routeIdentifier string, body AccountMagicRouteUpdateParams, opts ...option.RequestOption) (res *AccountMagicRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete multiple Magic static routes.
func (r *AccountMagicRouteService) Delete(ctx context.Context, accountIdentifier string, body AccountMagicRouteDeleteParams, opts ...option.RequestOption) (res *AccountMagicRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Disable and remove a specific Magic static route.
func (r *AccountMagicRouteService) DeleteMany(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *AccountMagicRouteDeleteManyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Magic static route. Use `?validate_only=true` as an optional query
// parameter to run validation only without persisting changes.
func (r *AccountMagicRouteService) MagicStaticRoutesNewRoutes(ctx context.Context, accountIdentifier string, body AccountMagicRouteMagicStaticRoutesNewRoutesParams, opts ...option.RequestOption) (res *AccountMagicRouteMagicStaticRoutesNewRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all Magic static routes.
func (r *AccountMagicRouteService) MagicStaticRoutesListRoutes(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountMagicRouteMagicStaticRoutesListRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update multiple Magic static routes. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes. Only fields
// for a route that need to be changed need be provided.
func (r *AccountMagicRouteService) MagicStaticRoutesUpdateManyRoutes(ctx context.Context, accountIdentifier string, body AccountMagicRouteMagicStaticRoutesUpdateManyRoutesParams, opts ...option.RequestOption) (res *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountMagicRouteGetResponse struct {
	Errors   []AccountMagicRouteGetResponseError   `json:"errors"`
	Messages []AccountMagicRouteGetResponseMessage `json:"messages"`
	Result   AccountMagicRouteGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteGetResponseSuccess `json:"success"`
	JSON    accountMagicRouteGetResponseJSON    `json:"-"`
}

// accountMagicRouteGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteGetResponse]
type accountMagicRouteGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountMagicRouteGetResponseErrorJSON `json:"-"`
}

// accountMagicRouteGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountMagicRouteGetResponseError]
type accountMagicRouteGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountMagicRouteGetResponseMessageJSON `json:"-"`
}

// accountMagicRouteGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicRouteGetResponseMessage]
type accountMagicRouteGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteGetResponseResult struct {
	Route interface{}                            `json:"route"`
	JSON  accountMagicRouteGetResponseResultJSON `json:"-"`
}

// accountMagicRouteGetResponseResultJSON contains the JSON metadata for the struct
// [AccountMagicRouteGetResponseResult]
type accountMagicRouteGetResponseResultJSON struct {
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteGetResponseSuccess bool

const (
	AccountMagicRouteGetResponseSuccessTrue AccountMagicRouteGetResponseSuccess = true
)

type AccountMagicRouteUpdateResponse struct {
	Errors   []AccountMagicRouteUpdateResponseError   `json:"errors"`
	Messages []AccountMagicRouteUpdateResponseMessage `json:"messages"`
	Result   AccountMagicRouteUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteUpdateResponseSuccess `json:"success"`
	JSON    accountMagicRouteUpdateResponseJSON    `json:"-"`
}

// accountMagicRouteUpdateResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteUpdateResponse]
type accountMagicRouteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountMagicRouteUpdateResponseErrorJSON `json:"-"`
}

// accountMagicRouteUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateResponseError]
type accountMagicRouteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountMagicRouteUpdateResponseMessageJSON `json:"-"`
}

// accountMagicRouteUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateResponseMessage]
type accountMagicRouteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteUpdateResponseResult struct {
	Modified      bool                                      `json:"modified"`
	ModifiedRoute interface{}                               `json:"modified_route"`
	JSON          accountMagicRouteUpdateResponseResultJSON `json:"-"`
}

// accountMagicRouteUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateResponseResult]
type accountMagicRouteUpdateResponseResultJSON struct {
	Modified      apijson.Field
	ModifiedRoute apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteUpdateResponseSuccess bool

const (
	AccountMagicRouteUpdateResponseSuccessTrue AccountMagicRouteUpdateResponseSuccess = true
)

type AccountMagicRouteDeleteResponse struct {
	Errors   []AccountMagicRouteDeleteResponseError   `json:"errors"`
	Messages []AccountMagicRouteDeleteResponseMessage `json:"messages"`
	Result   AccountMagicRouteDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteDeleteResponseSuccess `json:"success"`
	JSON    accountMagicRouteDeleteResponseJSON    `json:"-"`
}

// accountMagicRouteDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteDeleteResponse]
type accountMagicRouteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteDeleteResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountMagicRouteDeleteResponseErrorJSON `json:"-"`
}

// accountMagicRouteDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteResponseError]
type accountMagicRouteDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteDeleteResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountMagicRouteDeleteResponseMessageJSON `json:"-"`
}

// accountMagicRouteDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteResponseMessage]
type accountMagicRouteDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteDeleteResponseResult struct {
	Deleted       bool                                      `json:"deleted"`
	DeletedRoutes interface{}                               `json:"deleted_routes"`
	JSON          accountMagicRouteDeleteResponseResultJSON `json:"-"`
}

// accountMagicRouteDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteResponseResult]
type accountMagicRouteDeleteResponseResultJSON struct {
	Deleted       apijson.Field
	DeletedRoutes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteDeleteResponseSuccess bool

const (
	AccountMagicRouteDeleteResponseSuccessTrue AccountMagicRouteDeleteResponseSuccess = true
)

type AccountMagicRouteDeleteManyResponse struct {
	Errors   []AccountMagicRouteDeleteManyResponseError   `json:"errors"`
	Messages []AccountMagicRouteDeleteManyResponseMessage `json:"messages"`
	Result   AccountMagicRouteDeleteManyResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteDeleteManyResponseSuccess `json:"success"`
	JSON    accountMagicRouteDeleteManyResponseJSON    `json:"-"`
}

// accountMagicRouteDeleteManyResponseJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteManyResponse]
type accountMagicRouteDeleteManyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteManyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteDeleteManyResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountMagicRouteDeleteManyResponseErrorJSON `json:"-"`
}

// accountMagicRouteDeleteManyResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteManyResponseError]
type accountMagicRouteDeleteManyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteManyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteDeleteManyResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMagicRouteDeleteManyResponseMessageJSON `json:"-"`
}

// accountMagicRouteDeleteManyResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicRouteDeleteManyResponseMessage]
type accountMagicRouteDeleteManyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteManyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteDeleteManyResponseResult struct {
	Deleted      bool                                          `json:"deleted"`
	DeletedRoute interface{}                                   `json:"deleted_route"`
	JSON         accountMagicRouteDeleteManyResponseResultJSON `json:"-"`
}

// accountMagicRouteDeleteManyResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteManyResponseResult]
type accountMagicRouteDeleteManyResponseResultJSON struct {
	Deleted      apijson.Field
	DeletedRoute apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteManyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteDeleteManyResponseSuccess bool

const (
	AccountMagicRouteDeleteManyResponseSuccessTrue AccountMagicRouteDeleteManyResponseSuccess = true
)

type AccountMagicRouteMagicStaticRoutesNewRoutesResponse struct {
	Errors   []AccountMagicRouteMagicStaticRoutesNewRoutesResponseError   `json:"errors"`
	Messages []AccountMagicRouteMagicStaticRoutesNewRoutesResponseMessage `json:"messages"`
	Result   AccountMagicRouteMagicStaticRoutesNewRoutesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteMagicStaticRoutesNewRoutesResponseSuccess `json:"success"`
	JSON    accountMagicRouteMagicStaticRoutesNewRoutesResponseJSON    `json:"-"`
}

// accountMagicRouteMagicStaticRoutesNewRoutesResponseJSON contains the JSON
// metadata for the struct [AccountMagicRouteMagicStaticRoutesNewRoutesResponse]
type accountMagicRouteMagicStaticRoutesNewRoutesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesNewRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesNewRoutesResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountMagicRouteMagicStaticRoutesNewRoutesResponseErrorJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesNewRoutesResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountMagicRouteMagicStaticRoutesNewRoutesResponseError]
type accountMagicRouteMagicStaticRoutesNewRoutesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesNewRoutesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesNewRoutesResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountMagicRouteMagicStaticRoutesNewRoutesResponseMessageJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesNewRoutesResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountMagicRouteMagicStaticRoutesNewRoutesResponseMessage]
type accountMagicRouteMagicStaticRoutesNewRoutesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesNewRoutesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesNewRoutesResponseResult struct {
	Routes []AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoute `json:"routes"`
	JSON   accountMagicRouteMagicStaticRoutesNewRoutesResponseResultJSON    `json:"-"`
}

// accountMagicRouteMagicStaticRoutesNewRoutesResponseResultJSON contains the JSON
// metadata for the struct
// [AccountMagicRouteMagicStaticRoutesNewRoutesResponseResult]
type accountMagicRouteMagicStaticRoutesNewRoutesResponseResultJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesNewRoutesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoute struct {
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
	Scope AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                                              `json:"weight"`
	JSON   accountMagicRouteMagicStaticRoutesNewRoutesResponseResultRouteJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesNewRoutesResponseResultRouteJSON contains the
// JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoute]
type accountMagicRouteMagicStaticRoutesNewRoutesResponseResultRouteJSON struct {
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

func (r *AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                                                 `json:"colo_regions"`
	JSON        accountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScopeJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScopeJSON
// contains the JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScope]
type accountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesNewRoutesResponseResultRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteMagicStaticRoutesNewRoutesResponseSuccess bool

const (
	AccountMagicRouteMagicStaticRoutesNewRoutesResponseSuccessTrue AccountMagicRouteMagicStaticRoutesNewRoutesResponseSuccess = true
)

type AccountMagicRouteMagicStaticRoutesListRoutesResponse struct {
	Errors   []AccountMagicRouteMagicStaticRoutesListRoutesResponseError   `json:"errors"`
	Messages []AccountMagicRouteMagicStaticRoutesListRoutesResponseMessage `json:"messages"`
	Result   AccountMagicRouteMagicStaticRoutesListRoutesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteMagicStaticRoutesListRoutesResponseSuccess `json:"success"`
	JSON    accountMagicRouteMagicStaticRoutesListRoutesResponseJSON    `json:"-"`
}

// accountMagicRouteMagicStaticRoutesListRoutesResponseJSON contains the JSON
// metadata for the struct [AccountMagicRouteMagicStaticRoutesListRoutesResponse]
type accountMagicRouteMagicStaticRoutesListRoutesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesListRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesListRoutesResponseError struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountMagicRouteMagicStaticRoutesListRoutesResponseErrorJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesListRoutesResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountMagicRouteMagicStaticRoutesListRoutesResponseError]
type accountMagicRouteMagicStaticRoutesListRoutesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesListRoutesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesListRoutesResponseMessage struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountMagicRouteMagicStaticRoutesListRoutesResponseMessageJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesListRoutesResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesListRoutesResponseMessage]
type accountMagicRouteMagicStaticRoutesListRoutesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesListRoutesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesListRoutesResponseResult struct {
	Routes []AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoute `json:"routes"`
	JSON   accountMagicRouteMagicStaticRoutesListRoutesResponseResultJSON    `json:"-"`
}

// accountMagicRouteMagicStaticRoutesListRoutesResponseResultJSON contains the JSON
// metadata for the struct
// [AccountMagicRouteMagicStaticRoutesListRoutesResponseResult]
type accountMagicRouteMagicStaticRoutesListRoutesResponseResultJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesListRoutesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoute struct {
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
	Scope AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                                               `json:"weight"`
	JSON   accountMagicRouteMagicStaticRoutesListRoutesResponseResultRouteJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesListRoutesResponseResultRouteJSON contains the
// JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoute]
type accountMagicRouteMagicStaticRoutesListRoutesResponseResultRouteJSON struct {
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

func (r *AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                                                  `json:"colo_regions"`
	JSON        accountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScopeJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScopeJSON
// contains the JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScope]
type accountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesListRoutesResponseResultRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteMagicStaticRoutesListRoutesResponseSuccess bool

const (
	AccountMagicRouteMagicStaticRoutesListRoutesResponseSuccessTrue AccountMagicRouteMagicStaticRoutesListRoutesResponseSuccess = true
)

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponse struct {
	Errors   []AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseError   `json:"errors"`
	Messages []AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessage `json:"messages"`
	Result   AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseSuccess `json:"success"`
	JSON    accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseJSON    `json:"-"`
}

// accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseJSON contains the JSON
// metadata for the struct
// [AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponse]
type accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseErrorJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseError]
type accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessageJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessage]
type accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResult struct {
	Modified       bool                                                                            `json:"modified"`
	ModifiedRoutes []AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoute `json:"modified_routes"`
	JSON           accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultJSON            `json:"-"`
}

// accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultJSON contains
// the JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResult]
type accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultJSON struct {
	Modified       apijson.Field
	ModifiedRoutes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoute struct {
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
	Scope AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                                                             `json:"weight"`
	JSON   accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRouteJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRouteJSON
// contains the JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoute]
type accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRouteJSON struct {
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

func (r *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                                                                `json:"colo_regions"`
	JSON        accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScopeJSON `json:"-"`
}

// accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScopeJSON
// contains the JSON metadata for the struct
// [AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScope]
type accountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseResultModifiedRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseSuccess bool

const (
	AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseSuccessTrue AccountMagicRouteMagicStaticRoutesUpdateManyRoutesResponseSuccess = true
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

type AccountMagicRouteDeleteParams struct {
	Routes param.Field[[]AccountMagicRouteDeleteParamsRoute] `json:"routes,required"`
}

func (r AccountMagicRouteDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteDeleteParamsRoute struct {
}

func (r AccountMagicRouteDeleteParamsRoute) MarshalJSON() (data []byte, err error) {
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
