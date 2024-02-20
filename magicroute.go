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

// MagicRouteService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMagicRouteService] method instead.
type MagicRouteService struct {
	Options []option.RequestOption
}

// NewMagicRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagicRouteService(opts ...option.RequestOption) (r *MagicRouteService) {
	r = &MagicRouteService{}
	r.Options = opts
	return
}

// Creates a new Magic static route. Use `?validate_only=true` as an optional query
// parameter to run validation only without persisting changes.
func (r *MagicRouteService) New(ctx context.Context, accountIdentifier string, body MagicRouteNewParams, opts ...option.RequestOption) (res *MagicRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicRouteNewResponseEnvelope
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
func (r *MagicRouteService) Update(ctx context.Context, accountIdentifier string, routeIdentifier string, body MagicRouteUpdateParams, opts ...option.RequestOption) (res *MagicRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicRouteUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable and remove a specific Magic static route.
func (r *MagicRouteService) Delete(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *MagicRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicRouteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Magic static route.
func (r *MagicRouteService) Get(ctx context.Context, accountIdentifier string, routeIdentifier string, opts ...option.RequestOption) (res *MagicRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicRouteGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountIdentifier, routeIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Magic static routes.
func (r *MagicRouteService) MagicStaticRoutesListRoutes(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicRouteMagicStaticRoutesListRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicRouteMagicStaticRoutesListRoutesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/routes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicRouteNewResponse struct {
	Routes []MagicRouteNewResponseRoute `json:"routes"`
	JSON   magicRouteNewResponseJSON    `json:"-"`
}

// magicRouteNewResponseJSON contains the JSON metadata for the struct
// [MagicRouteNewResponse]
type magicRouteNewResponseJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteNewResponseRoute struct {
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
	Scope MagicRouteNewResponseRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                          `json:"weight"`
	JSON   magicRouteNewResponseRouteJSON `json:"-"`
}

// magicRouteNewResponseRouteJSON contains the JSON metadata for the struct
// [MagicRouteNewResponseRoute]
type magicRouteNewResponseRouteJSON struct {
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

func (r *MagicRouteNewResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type MagicRouteNewResponseRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                             `json:"colo_regions"`
	JSON        magicRouteNewResponseRoutesScopeJSON `json:"-"`
}

// magicRouteNewResponseRoutesScopeJSON contains the JSON metadata for the struct
// [MagicRouteNewResponseRoutesScope]
type magicRouteNewResponseRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteNewResponseRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteUpdateResponse struct {
	Modified      bool                         `json:"modified"`
	ModifiedRoute interface{}                  `json:"modified_route"`
	JSON          magicRouteUpdateResponseJSON `json:"-"`
}

// magicRouteUpdateResponseJSON contains the JSON metadata for the struct
// [MagicRouteUpdateResponse]
type magicRouteUpdateResponseJSON struct {
	Modified      apijson.Field
	ModifiedRoute apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteDeleteResponse struct {
	Deleted      bool                         `json:"deleted"`
	DeletedRoute interface{}                  `json:"deleted_route"`
	JSON         magicRouteDeleteResponseJSON `json:"-"`
}

// magicRouteDeleteResponseJSON contains the JSON metadata for the struct
// [MagicRouteDeleteResponse]
type magicRouteDeleteResponseJSON struct {
	Deleted      apijson.Field
	DeletedRoute apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteGetResponse struct {
	Route interface{}               `json:"route"`
	JSON  magicRouteGetResponseJSON `json:"-"`
}

// magicRouteGetResponseJSON contains the JSON metadata for the struct
// [MagicRouteGetResponse]
type magicRouteGetResponseJSON struct {
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteMagicStaticRoutesListRoutesResponse struct {
	Routes []MagicRouteMagicStaticRoutesListRoutesResponseRoute `json:"routes"`
	JSON   magicRouteMagicStaticRoutesListRoutesResponseJSON    `json:"-"`
}

// magicRouteMagicStaticRoutesListRoutesResponseJSON contains the JSON metadata for
// the struct [MagicRouteMagicStaticRoutesListRoutesResponse]
type magicRouteMagicStaticRoutesListRoutesResponseJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteMagicStaticRoutesListRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteMagicStaticRoutesListRoutesResponseRoute struct {
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
	Scope MagicRouteMagicStaticRoutesListRoutesResponseRoutesScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64                                                  `json:"weight"`
	JSON   magicRouteMagicStaticRoutesListRoutesResponseRouteJSON `json:"-"`
}

// magicRouteMagicStaticRoutesListRoutesResponseRouteJSON contains the JSON
// metadata for the struct [MagicRouteMagicStaticRoutesListRoutesResponseRoute]
type magicRouteMagicStaticRoutesListRoutesResponseRouteJSON struct {
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

func (r *MagicRouteMagicStaticRoutesListRoutesResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Used only for ECMP routes.
type MagicRouteMagicStaticRoutesListRoutesResponseRoutesScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string                                                     `json:"colo_regions"`
	JSON        magicRouteMagicStaticRoutesListRoutesResponseRoutesScopeJSON `json:"-"`
}

// magicRouteMagicStaticRoutesListRoutesResponseRoutesScopeJSON contains the JSON
// metadata for the struct
// [MagicRouteMagicStaticRoutesListRoutesResponseRoutesScope]
type magicRouteMagicStaticRoutesListRoutesResponseRoutesScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteMagicStaticRoutesListRoutesResponseRoutesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicRouteNewResponseEnvelope struct {
	Errors   []MagicRouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicRouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicRouteNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicRouteNewResponseEnvelopeJSON    `json:"-"`
}

// magicRouteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicRouteNewResponseEnvelope]
type magicRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    magicRouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicRouteNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicRouteNewResponseEnvelopeErrors]
type magicRouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    magicRouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicRouteNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicRouteNewResponseEnvelopeMessages]
type magicRouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicRouteNewResponseEnvelopeSuccess bool

const (
	MagicRouteNewResponseEnvelopeSuccessTrue MagicRouteNewResponseEnvelopeSuccess = true
)

type MagicRouteUpdateParams struct {
	// The next-hop IP Address for the static route.
	Nexthop param.Field[string] `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix param.Field[string] `json:"prefix,required"`
	// Priority of the static route.
	Priority param.Field[int64] `json:"priority,required"`
	// An optional human provided description of the static route.
	Description param.Field[string] `json:"description"`
	// Used only for ECMP routes.
	Scope param.Field[MagicRouteUpdateParamsScope] `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight param.Field[int64] `json:"weight"`
}

func (r MagicRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Used only for ECMP routes.
type MagicRouteUpdateParamsScope struct {
	// List of colo names for the ECMP scope.
	ColoNames param.Field[[]string] `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions param.Field[[]string] `json:"colo_regions"`
}

func (r MagicRouteUpdateParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicRouteUpdateResponseEnvelope struct {
	Errors   []MagicRouteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicRouteUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicRouteUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicRouteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicRouteUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicRouteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicRouteUpdateResponseEnvelope]
type magicRouteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    magicRouteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicRouteUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicRouteUpdateResponseEnvelopeErrors]
type magicRouteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    magicRouteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicRouteUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicRouteUpdateResponseEnvelopeMessages]
type magicRouteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicRouteUpdateResponseEnvelopeSuccess bool

const (
	MagicRouteUpdateResponseEnvelopeSuccessTrue MagicRouteUpdateResponseEnvelopeSuccess = true
)

type MagicRouteDeleteResponseEnvelope struct {
	Errors   []MagicRouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicRouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicRouteDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicRouteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicRouteDeleteResponseEnvelope]
type magicRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    magicRouteDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicRouteDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicRouteDeleteResponseEnvelopeErrors]
type magicRouteDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    magicRouteDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicRouteDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicRouteDeleteResponseEnvelopeMessages]
type magicRouteDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicRouteDeleteResponseEnvelopeSuccess bool

const (
	MagicRouteDeleteResponseEnvelopeSuccessTrue MagicRouteDeleteResponseEnvelopeSuccess = true
)

type MagicRouteGetResponseEnvelope struct {
	Errors   []MagicRouteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicRouteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicRouteGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicRouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicRouteGetResponseEnvelopeJSON    `json:"-"`
}

// magicRouteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MagicRouteGetResponseEnvelope]
type magicRouteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    magicRouteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicRouteGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MagicRouteGetResponseEnvelopeErrors]
type magicRouteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    magicRouteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicRouteGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MagicRouteGetResponseEnvelopeMessages]
type magicRouteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicRouteGetResponseEnvelopeSuccess bool

const (
	MagicRouteGetResponseEnvelopeSuccessTrue MagicRouteGetResponseEnvelopeSuccess = true
)

type MagicRouteMagicStaticRoutesListRoutesResponseEnvelope struct {
	Errors   []MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicRouteMagicStaticRoutesListRoutesResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicRouteMagicStaticRoutesListRoutesResponseEnvelopeJSON    `json:"-"`
}

// magicRouteMagicStaticRoutesListRoutesResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicRouteMagicStaticRoutesListRoutesResponseEnvelope]
type magicRouteMagicStaticRoutesListRoutesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteMagicStaticRoutesListRoutesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    magicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrorsJSON `json:"-"`
}

// magicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrors]
type magicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    magicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessagesJSON `json:"-"`
}

// magicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessages]
type magicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeSuccess bool

const (
	MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeSuccessTrue MagicRouteMagicStaticRoutesListRoutesResponseEnvelopeSuccess = true
)
