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

// MagicTransitRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicTransitRouteService] method
// instead.
type MagicTransitRouteService struct {
	Options []option.RequestOption
}

// NewMagicTransitRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicTransitRouteService(opts ...option.RequestOption) (r *MagicTransitRouteService) {
	r = &MagicTransitRouteService{}
	r.Options = opts
	return
}

// Creates a new Magic static route. Use `?validate_only=true` as an optional query
// parameter to run validation only without persisting changes.
func (r *MagicTransitRouteService) New(ctx context.Context, accountIdentifier string, body MagicTransitRouteNewParams, opts ...option.RequestOption) (res *MagicTransitRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitRouteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific Magic static route. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes.
func (r *MagicTransitRouteService) Update(ctx context.Context, accountIdentifier string, routeIdentifier string, body MagicTransitRouteUpdateParams, opts ...option.RequestOption) (res *MagicTransitRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitRouteUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Magic static routes.
func (r *MagicTransitRouteService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicTransitRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitRouteListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable and remove a specific Magic static route.
func (r *MagicTransitRouteService) Delete(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *MagicTransitRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitRouteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete multiple Magic static routes.
func (r *MagicTransitRouteService) Empty(ctx context.Context, accountIdentifier string, body MagicTransitRouteEmptyParams, opts ...option.RequestOption) (res *MagicTransitRouteEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitRouteEmptyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Magic static route.
func (r *MagicTransitRouteService) Get(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *MagicTransitRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitRouteGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicTransitRouteNewResponse struct {
	Routes []MagicTransitRouteNewResponseRoute `json:"routes"`
	JSON   magicTransitRouteNewResponseJSON    `json:"-"`
}

// magicTransitRouteNewResponseJSON contains the JSON metadata for the struct
// [MagicTransitRouteNewResponse]
type magicTransitRouteNewResponseJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteNewResponseRoute struct {
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
	Scope MagicTransitRouteNewResponseRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                 `json:"weight"`
	JSON   magicTransitRouteNewResponseRouteJSON `json:"-"`
}

// magicTransitRouteNewResponseRouteJSON contains the JSON metadata for the struct
// [MagicTransitRouteNewResponseRoute]
type magicTransitRouteNewResponseRouteJSON struct {
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

func (r *MagicTransitRouteNewResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type MagicTransitRouteNewResponseRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                    `json:"colo_regions"`
	JSON        magicTransitRouteNewResponseRoutesScopeJSON `json:"-"`
}

// magicTransitRouteNewResponseRoutesScopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteNewResponseRoutesScope]
type magicTransitRouteNewResponseRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteNewResponseRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteUpdateResponse struct {
	Modified      bool                                `json:"modified"`
	ModifiedRoute interface{}                         `json:"modified_route"`
	JSON          magicTransitRouteUpdateResponseJSON `json:"-"`
}

// magicTransitRouteUpdateResponseJSON contains the JSON metadata for the struct
// [MagicTransitRouteUpdateResponse]
type magicTransitRouteUpdateResponseJSON struct {
	Modified      apijson.Field
	ModifiedRoute apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicTransitRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteListResponse struct {
	Routes []MagicTransitRouteListResponseRoute `json:"routes"`
	JSON   magicTransitRouteListResponseJSON    `json:"-"`
}

// magicTransitRouteListResponseJSON contains the JSON metadata for the struct
// [MagicTransitRouteListResponse]
type magicTransitRouteListResponseJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteListResponseRoute struct {
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
	Scope MagicTransitRouteListResponseRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                  `json:"weight"`
	JSON   magicTransitRouteListResponseRouteJSON `json:"-"`
}

// magicTransitRouteListResponseRouteJSON contains the JSON metadata for the struct
// [MagicTransitRouteListResponseRoute]
type magicTransitRouteListResponseRouteJSON struct {
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

func (r *MagicTransitRouteListResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type MagicTransitRouteListResponseRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                     `json:"colo_regions"`
	JSON        magicTransitRouteListResponseRoutesScopeJSON `json:"-"`
}

// magicTransitRouteListResponseRoutesScopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteListResponseRoutesScope]
type magicTransitRouteListResponseRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteListResponseRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteDeleteResponse struct {
	Deleted      bool                                `json:"deleted"`
	DeletedRoute interface{}                         `json:"deleted_route"`
	JSON         magicTransitRouteDeleteResponseJSON `json:"-"`
}

// magicTransitRouteDeleteResponseJSON contains the JSON metadata for the struct
// [MagicTransitRouteDeleteResponse]
type magicTransitRouteDeleteResponseJSON struct {
	Deleted      apijson.Field
	DeletedRoute apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicTransitRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteEmptyResponse struct {
	Deleted       bool                               `json:"deleted"`
	DeletedRoutes interface{}                        `json:"deleted_routes"`
	JSON          magicTransitRouteEmptyResponseJSON `json:"-"`
}

// magicTransitRouteEmptyResponseJSON contains the JSON metadata for the struct
// [MagicTransitRouteEmptyResponse]
type magicTransitRouteEmptyResponseJSON struct {
	Deleted       apijson.Field
	DeletedRoutes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicTransitRouteEmptyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteGetResponse struct {
	Route interface{}                      `json:"route"`
	JSON  magicTransitRouteGetResponseJSON `json:"-"`
}

// magicTransitRouteGetResponseJSON contains the JSON metadata for the struct
// [MagicTransitRouteGetResponse]
type magicTransitRouteGetResponseJSON struct {
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicTransitRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicTransitRouteNewResponseEnvelope struct {
	Errors   []MagicTransitRouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitRouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitRouteNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitRouteNewResponseEnvelopeJSON    `json:"-"`
}

// magicTransitRouteNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteNewResponseEnvelope]
type magicTransitRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicTransitRouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitRouteNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicTransitRouteNewResponseEnvelopeErrors]
type magicTransitRouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicTransitRouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitRouteNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicTransitRouteNewResponseEnvelopeMessages]
type magicTransitRouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitRouteNewResponseEnvelopeSuccess bool

const (
	MagicTransitRouteNewResponseEnvelopeSuccessTrue MagicTransitRouteNewResponseEnvelopeSuccess = true
)

type MagicTransitRouteUpdateParams struct {
	// The next-hop IP Address for the static route.
	Nexthop param.Field[string] `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix param.Field[string] `json:"prefix,required"`
	// Priority of the static route.
	Priority param.Field[int64] `json:"priority,required"`
	// An optional human provided description of the static route.
	Description param.Field[string] `json:"description"`
	// Used only for ECMP routes.
	Scope param.Field[MagicTransitRouteUpdateParamsScope] `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight param.Field[int64] `json:"weight"`
}

func (r MagicTransitRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Used only for ECMP routes.
type MagicTransitRouteUpdateParamsScope struct {
	// List of colo names for the ECMP scope.
	ColoNames param.Field[[]string] `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions param.Field[[]string] `json:"colo_regions"`
}

func (r MagicTransitRouteUpdateParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitRouteUpdateResponseEnvelope struct {
	Errors   []MagicTransitRouteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitRouteUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitRouteUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitRouteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitRouteUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicTransitRouteUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteUpdateResponseEnvelope]
type magicTransitRouteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    magicTransitRouteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitRouteUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicTransitRouteUpdateResponseEnvelopeErrors]
type magicTransitRouteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    magicTransitRouteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitRouteUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicTransitRouteUpdateResponseEnvelopeMessages]
type magicTransitRouteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitRouteUpdateResponseEnvelopeSuccess bool

const (
	MagicTransitRouteUpdateResponseEnvelopeSuccessTrue MagicTransitRouteUpdateResponseEnvelopeSuccess = true
)

type MagicTransitRouteListResponseEnvelope struct {
	Errors   []MagicTransitRouteListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitRouteListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitRouteListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitRouteListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitRouteListResponseEnvelopeJSON    `json:"-"`
}

// magicTransitRouteListResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteListResponseEnvelope]
type magicTransitRouteListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    magicTransitRouteListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitRouteListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicTransitRouteListResponseEnvelopeErrors]
type magicTransitRouteListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    magicTransitRouteListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitRouteListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicTransitRouteListResponseEnvelopeMessages]
type magicTransitRouteListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitRouteListResponseEnvelopeSuccess bool

const (
	MagicTransitRouteListResponseEnvelopeSuccessTrue MagicTransitRouteListResponseEnvelopeSuccess = true
)

type MagicTransitRouteDeleteResponseEnvelope struct {
	Errors   []MagicTransitRouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitRouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitRouteDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicTransitRouteDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteDeleteResponseEnvelope]
type magicTransitRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    magicTransitRouteDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitRouteDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicTransitRouteDeleteResponseEnvelopeErrors]
type magicTransitRouteDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    magicTransitRouteDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitRouteDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicTransitRouteDeleteResponseEnvelopeMessages]
type magicTransitRouteDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitRouteDeleteResponseEnvelopeSuccess bool

const (
	MagicTransitRouteDeleteResponseEnvelopeSuccessTrue MagicTransitRouteDeleteResponseEnvelopeSuccess = true
)

type MagicTransitRouteEmptyParams struct {
	Routes param.Field[[]MagicTransitRouteEmptyParamsRoute] `json:"routes,required"`
}

func (r MagicTransitRouteEmptyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitRouteEmptyParamsRoute struct {
}

func (r MagicTransitRouteEmptyParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitRouteEmptyResponseEnvelope struct {
	Errors   []MagicTransitRouteEmptyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitRouteEmptyResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitRouteEmptyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitRouteEmptyResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitRouteEmptyResponseEnvelopeJSON    `json:"-"`
}

// magicTransitRouteEmptyResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteEmptyResponseEnvelope]
type magicTransitRouteEmptyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteEmptyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteEmptyResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicTransitRouteEmptyResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitRouteEmptyResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicTransitRouteEmptyResponseEnvelopeErrors]
type magicTransitRouteEmptyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteEmptyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteEmptyResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicTransitRouteEmptyResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitRouteEmptyResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicTransitRouteEmptyResponseEnvelopeMessages]
type magicTransitRouteEmptyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteEmptyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitRouteEmptyResponseEnvelopeSuccess bool

const (
	MagicTransitRouteEmptyResponseEnvelopeSuccessTrue MagicTransitRouteEmptyResponseEnvelopeSuccess = true
)

type MagicTransitRouteGetResponseEnvelope struct {
	Errors   []MagicTransitRouteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitRouteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitRouteGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitRouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitRouteGetResponseEnvelopeJSON    `json:"-"`
}

// magicTransitRouteGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicTransitRouteGetResponseEnvelope]
type magicTransitRouteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    magicTransitRouteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitRouteGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicTransitRouteGetResponseEnvelopeErrors]
type magicTransitRouteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitRouteGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicTransitRouteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitRouteGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MagicTransitRouteGetResponseEnvelopeMessages]
type magicTransitRouteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitRouteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitRouteGetResponseEnvelopeSuccess bool

const (
	MagicTransitRouteGetResponseEnvelopeSuccessTrue MagicTransitRouteGetResponseEnvelopeSuccess = true
)
