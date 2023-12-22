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

// Get Route
func (r *ZoneWorkerRouteService) Get(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *RouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Route
func (r *ZoneWorkerRouteService) Update(ctx context.Context, zoneID string, routeID string, body ZoneWorkerRouteUpdateParams, opts ...option.RequestOption) (res *RouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete Route
func (r *ZoneWorkerRouteService) Delete(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Route
func (r *ZoneWorkerRouteService) WorkerRoutesNewRoute(ctx context.Context, zoneID string, body ZoneWorkerRouteWorkerRoutesNewRouteParams, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Routes
func (r *ZoneWorkerRouteService) WorkerRoutesListRoutes(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *RouteResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RouteResponseCollection struct {
	Errors   []RouteResponseCollectionError   `json:"errors"`
	Messages []RouteResponseCollectionMessage `json:"messages"`
	Result   []RouteResponseCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success RouteResponseCollectionSuccess `json:"success"`
	JSON    routeResponseCollectionJSON    `json:"-"`
}

// routeResponseCollectionJSON contains the JSON metadata for the struct
// [RouteResponseCollection]
type routeResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteResponseCollectionError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    routeResponseCollectionErrorJSON `json:"-"`
}

// routeResponseCollectionErrorJSON contains the JSON metadata for the struct
// [RouteResponseCollectionError]
type routeResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteResponseCollectionMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    routeResponseCollectionMessageJSON `json:"-"`
}

// routeResponseCollectionMessageJSON contains the JSON metadata for the struct
// [RouteResponseCollectionMessage]
type routeResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteResponseCollectionResult struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script to apply when the route is matched. The route is skipped when
	// this is blank/missing.
	Script string                            `json:"script,required"`
	JSON   routeResponseCollectionResultJSON `json:"-"`
}

// routeResponseCollectionResultJSON contains the JSON metadata for the struct
// [RouteResponseCollectionResult]
type routeResponseCollectionResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RouteResponseCollectionSuccess bool

const (
	RouteResponseCollectionSuccessTrue RouteResponseCollectionSuccess = true
)

type RouteResponseSingle struct {
	Errors   []RouteResponseSingleError   `json:"errors"`
	Messages []RouteResponseSingleMessage `json:"messages"`
	Result   RouteResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success RouteResponseSingleSuccess `json:"success"`
	JSON    routeResponseSingleJSON    `json:"-"`
}

// routeResponseSingleJSON contains the JSON metadata for the struct
// [RouteResponseSingle]
type routeResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteResponseSingleError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    routeResponseSingleErrorJSON `json:"-"`
}

// routeResponseSingleErrorJSON contains the JSON metadata for the struct
// [RouteResponseSingleError]
type routeResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteResponseSingleMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    routeResponseSingleMessageJSON `json:"-"`
}

// routeResponseSingleMessageJSON contains the JSON metadata for the struct
// [RouteResponseSingleMessage]
type routeResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RouteResponseSingleResult struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script to apply when the route is matched. The route is skipped when
	// this is blank/missing.
	Script string                        `json:"script,required"`
	JSON   routeResponseSingleResultJSON `json:"-"`
}

// routeResponseSingleResultJSON contains the JSON metadata for the struct
// [RouteResponseSingleResult]
type routeResponseSingleResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RouteResponseSingleSuccess bool

const (
	RouteResponseSingleSuccessTrue RouteResponseSingleSuccess = true
)

type ZoneWorkerRouteUpdateParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script to apply when the route is matched. The route is skipped when
	// this is blank/missing.
	Script param.Field[string] `json:"script"`
}

func (r ZoneWorkerRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWorkerRouteWorkerRoutesNewRouteParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script to apply when the route is matched. The route is skipped when
	// this is blank/missing.
	Script param.Field[string] `json:"script"`
}

func (r ZoneWorkerRouteWorkerRoutesNewRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
