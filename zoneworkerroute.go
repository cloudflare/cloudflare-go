// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneWorkerRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWorkerRouteService] method
// instead.
type ZoneWorkerRouteService struct {
	Options []option.RequestOption
}

// NewZoneWorkerRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneWorkerRouteService(opts ...option.RequestOption) (r *ZoneWorkerRouteService) {
	r = &ZoneWorkerRouteService{}
	r.Options = opts
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *ZoneWorkerRouteService) Get(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *ZoneWorkerRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *ZoneWorkerRouteService) Update(ctx context.Context, zoneID string, routeID string, body ZoneWorkerRouteUpdateParams, opts ...option.RequestOption) (res *ZoneWorkerRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a route.
func (r *ZoneWorkerRouteService) Delete(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *ZoneWorkerRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a route that maps a URL pattern to a Worker.
func (r *ZoneWorkerRouteService) WorkerRoutesNewRoute(ctx context.Context, zoneID string, body ZoneWorkerRouteWorkerRoutesNewRouteParams, opts ...option.RequestOption) (res *ZoneWorkerRouteWorkerRoutesNewRouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns routes for a zone.
func (r *ZoneWorkerRouteService) WorkerRoutesListRoutes(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneWorkerRouteWorkerRoutesListRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneWorkerRouteGetResponse struct {
	Errors   []ZoneWorkerRouteGetResponseError   `json:"errors"`
	Messages []ZoneWorkerRouteGetResponseMessage `json:"messages"`
	Result   ZoneWorkerRouteGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerRouteGetResponseSuccess `json:"success"`
	JSON    zoneWorkerRouteGetResponseJSON    `json:"-"`
}

// zoneWorkerRouteGetResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteGetResponse]
type zoneWorkerRouteGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneWorkerRouteGetResponseErrorJSON `json:"-"`
}

// zoneWorkerRouteGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteGetResponseError]
type zoneWorkerRouteGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneWorkerRouteGetResponseMessageJSON `json:"-"`
}

// zoneWorkerRouteGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteGetResponseMessage]
type zoneWorkerRouteGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteGetResponseResult struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                               `json:"script,required"`
	JSON   zoneWorkerRouteGetResponseResultJSON `json:"-"`
}

// zoneWorkerRouteGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteGetResponseResult]
type zoneWorkerRouteGetResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerRouteGetResponseSuccess bool

const (
	ZoneWorkerRouteGetResponseSuccessTrue ZoneWorkerRouteGetResponseSuccess = true
)

type ZoneWorkerRouteUpdateResponse struct {
	Errors   []ZoneWorkerRouteUpdateResponseError   `json:"errors"`
	Messages []ZoneWorkerRouteUpdateResponseMessage `json:"messages"`
	Result   ZoneWorkerRouteUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerRouteUpdateResponseSuccess `json:"success"`
	JSON    zoneWorkerRouteUpdateResponseJSON    `json:"-"`
}

// zoneWorkerRouteUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteUpdateResponse]
type zoneWorkerRouteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneWorkerRouteUpdateResponseErrorJSON `json:"-"`
}

// zoneWorkerRouteUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteUpdateResponseError]
type zoneWorkerRouteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneWorkerRouteUpdateResponseMessageJSON `json:"-"`
}

// zoneWorkerRouteUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWorkerRouteUpdateResponseMessage]
type zoneWorkerRouteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteUpdateResponseResult struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                                  `json:"script,required"`
	JSON   zoneWorkerRouteUpdateResponseResultJSON `json:"-"`
}

// zoneWorkerRouteUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneWorkerRouteUpdateResponseResult]
type zoneWorkerRouteUpdateResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerRouteUpdateResponseSuccess bool

const (
	ZoneWorkerRouteUpdateResponseSuccessTrue ZoneWorkerRouteUpdateResponseSuccess = true
)

type ZoneWorkerRouteDeleteResponse struct {
	Errors   []ZoneWorkerRouteDeleteResponseError   `json:"errors"`
	Messages []ZoneWorkerRouteDeleteResponseMessage `json:"messages"`
	Result   ZoneWorkerRouteDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerRouteDeleteResponseSuccess `json:"success"`
	JSON    zoneWorkerRouteDeleteResponseJSON    `json:"-"`
}

// zoneWorkerRouteDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteDeleteResponse]
type zoneWorkerRouteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneWorkerRouteDeleteResponseErrorJSON `json:"-"`
}

// zoneWorkerRouteDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteDeleteResponseError]
type zoneWorkerRouteDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneWorkerRouteDeleteResponseMessageJSON `json:"-"`
}

// zoneWorkerRouteDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWorkerRouteDeleteResponseMessage]
type zoneWorkerRouteDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneWorkerRouteDeleteResponseResultUnknown] or
// [shared.UnionString].
type ZoneWorkerRouteDeleteResponseResult interface {
	ImplementsZoneWorkerRouteDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneWorkerRouteDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type ZoneWorkerRouteDeleteResponseSuccess bool

const (
	ZoneWorkerRouteDeleteResponseSuccessTrue ZoneWorkerRouteDeleteResponseSuccess = true
)

type ZoneWorkerRouteWorkerRoutesNewRouteResponse struct {
	Errors   []ZoneWorkerRouteWorkerRoutesNewRouteResponseError   `json:"errors"`
	Messages []ZoneWorkerRouteWorkerRoutesNewRouteResponseMessage `json:"messages"`
	Result   ZoneWorkerRouteWorkerRoutesNewRouteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerRouteWorkerRoutesNewRouteResponseSuccess `json:"success"`
	JSON    zoneWorkerRouteWorkerRoutesNewRouteResponseJSON    `json:"-"`
}

// zoneWorkerRouteWorkerRoutesNewRouteResponseJSON contains the JSON metadata for
// the struct [ZoneWorkerRouteWorkerRoutesNewRouteResponse]
type zoneWorkerRouteWorkerRoutesNewRouteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesNewRouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteWorkerRoutesNewRouteResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneWorkerRouteWorkerRoutesNewRouteResponseErrorJSON `json:"-"`
}

// zoneWorkerRouteWorkerRoutesNewRouteResponseErrorJSON contains the JSON metadata
// for the struct [ZoneWorkerRouteWorkerRoutesNewRouteResponseError]
type zoneWorkerRouteWorkerRoutesNewRouteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesNewRouteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteWorkerRoutesNewRouteResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneWorkerRouteWorkerRoutesNewRouteResponseMessageJSON `json:"-"`
}

// zoneWorkerRouteWorkerRoutesNewRouteResponseMessageJSON contains the JSON
// metadata for the struct [ZoneWorkerRouteWorkerRoutesNewRouteResponseMessage]
type zoneWorkerRouteWorkerRoutesNewRouteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesNewRouteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneWorkerRouteWorkerRoutesNewRouteResponseResultUnknown] or
// [shared.UnionString].
type ZoneWorkerRouteWorkerRoutesNewRouteResponseResult interface {
	ImplementsZoneWorkerRouteWorkerRoutesNewRouteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneWorkerRouteWorkerRoutesNewRouteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type ZoneWorkerRouteWorkerRoutesNewRouteResponseSuccess bool

const (
	ZoneWorkerRouteWorkerRoutesNewRouteResponseSuccessTrue ZoneWorkerRouteWorkerRoutesNewRouteResponseSuccess = true
)

type ZoneWorkerRouteWorkerRoutesListRoutesResponse struct {
	Errors   []ZoneWorkerRouteWorkerRoutesListRoutesResponseError   `json:"errors"`
	Messages []ZoneWorkerRouteWorkerRoutesListRoutesResponseMessage `json:"messages"`
	Result   []ZoneWorkerRouteWorkerRoutesListRoutesResponseResult  `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerRouteWorkerRoutesListRoutesResponseSuccess `json:"success"`
	JSON    zoneWorkerRouteWorkerRoutesListRoutesResponseJSON    `json:"-"`
}

// zoneWorkerRouteWorkerRoutesListRoutesResponseJSON contains the JSON metadata for
// the struct [ZoneWorkerRouteWorkerRoutesListRoutesResponse]
type zoneWorkerRouteWorkerRoutesListRoutesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesListRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteWorkerRoutesListRoutesResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneWorkerRouteWorkerRoutesListRoutesResponseErrorJSON `json:"-"`
}

// zoneWorkerRouteWorkerRoutesListRoutesResponseErrorJSON contains the JSON
// metadata for the struct [ZoneWorkerRouteWorkerRoutesListRoutesResponseError]
type zoneWorkerRouteWorkerRoutesListRoutesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesListRoutesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteWorkerRoutesListRoutesResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneWorkerRouteWorkerRoutesListRoutesResponseMessageJSON `json:"-"`
}

// zoneWorkerRouteWorkerRoutesListRoutesResponseMessageJSON contains the JSON
// metadata for the struct [ZoneWorkerRouteWorkerRoutesListRoutesResponseMessage]
type zoneWorkerRouteWorkerRoutesListRoutesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesListRoutesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerRouteWorkerRoutesListRoutesResponseResult struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                                                  `json:"script,required"`
	JSON   zoneWorkerRouteWorkerRoutesListRoutesResponseResultJSON `json:"-"`
}

// zoneWorkerRouteWorkerRoutesListRoutesResponseResultJSON contains the JSON
// metadata for the struct [ZoneWorkerRouteWorkerRoutesListRoutesResponseResult]
type zoneWorkerRouteWorkerRoutesListRoutesResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteWorkerRoutesListRoutesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerRouteWorkerRoutesListRoutesResponseSuccess bool

const (
	ZoneWorkerRouteWorkerRoutesListRoutesResponseSuccessTrue ZoneWorkerRouteWorkerRoutesListRoutesResponseSuccess = true
)

type ZoneWorkerRouteUpdateParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r ZoneWorkerRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWorkerRouteWorkerRoutesNewRouteParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r ZoneWorkerRouteWorkerRoutesNewRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
