// File generated from our OpenAPI spec by Stainless.

package workers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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

// Creates a route that maps a URL pattern to a Worker.
func (r *RouteService) New(ctx context.Context, params RouteNewParams, opts ...option.RequestOption) (res *RouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *RouteService) Update(ctx context.Context, routeID string, params RouteUpdateParams, opts ...option.RequestOption) (res *WorkersRoutes, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", params.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns routes for a zone.
func (r *RouteService) List(ctx context.Context, query RouteListParams, opts ...option.RequestOption) (res *[]WorkersRoutes, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteListResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a route.
func (r *RouteService) Delete(ctx context.Context, routeID string, body RouteDeleteParams, opts ...option.RequestOption) (res *RouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", body.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *RouteService) Get(ctx context.Context, routeID string, query RouteGetParams, opts ...option.RequestOption) (res *WorkersRoutes, err error) {
	opts = append(r.Options[:], opts...)
	var env RouteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", query.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersRoutes struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string            `json:"script,required"`
	JSON   workersRoutesJSON `json:"-"`
}

// workersRoutesJSON contains the JSON metadata for the struct [WorkersRoutes]
type workersRoutesJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersRoutes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersRoutesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [workers.RouteNewResponseUnknown] or [shared.UnionString].
type RouteNewResponse interface {
	ImplementsWorkersRouteNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RouteNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [workers.RouteDeleteResponseUnknown] or [shared.UnionString].
type RouteDeleteResponse interface {
	ImplementsWorkersRouteDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RouteDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RouteNewParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r RouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteNewResponseEnvelope struct {
	Errors   []RouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RouteNewResponse                   `json:"result,required"`
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

type RouteNewResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    routeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// routeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelopeErrors]
type routeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteNewResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    routeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// routeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelopeMessages]
type routeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteNewResponseEnvelopeSuccess bool

const (
	RouteNewResponseEnvelopeSuccessTrue RouteNewResponseEnvelopeSuccess = true
)

type RouteUpdateParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r RouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteUpdateResponseEnvelope struct {
	Errors   []RouteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersRoutes                         `json:"result,required"`
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

type RouteUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    routeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// routeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteUpdateResponseEnvelopeErrors]
type routeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    routeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// routeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RouteUpdateResponseEnvelopeMessages]
type routeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteUpdateResponseEnvelopeSuccess bool

const (
	RouteUpdateResponseEnvelopeSuccessTrue RouteUpdateResponseEnvelopeSuccess = true
)

type RouteListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteListResponseEnvelope struct {
	Errors   []RouteListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkersRoutes                     `json:"result,required"`
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

type RouteListResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    routeListResponseEnvelopeErrorsJSON `json:"-"`
}

// routeListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteListResponseEnvelopeErrors]
type routeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteListResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    routeListResponseEnvelopeMessagesJSON `json:"-"`
}

// routeListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RouteListResponseEnvelopeMessages]
type routeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteListResponseEnvelopeSuccess bool

const (
	RouteListResponseEnvelopeSuccessTrue RouteListResponseEnvelopeSuccess = true
)

type RouteDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteDeleteResponseEnvelope struct {
	Errors   []RouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RouteDeleteResponse                   `json:"result,required"`
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

type RouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    routeDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// routeDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteDeleteResponseEnvelopeErrors]
type routeDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    routeDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// routeDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RouteDeleteResponseEnvelopeMessages]
type routeDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteDeleteResponseEnvelopeSuccess bool

const (
	RouteDeleteResponseEnvelopeSuccessTrue RouteDeleteResponseEnvelopeSuccess = true
)

type RouteGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteGetResponseEnvelope struct {
	Errors   []RouteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersRoutes                      `json:"result,required"`
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

type RouteGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    routeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// routeGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelopeErrors]
type routeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    routeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// routeGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelopeMessages]
type routeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RouteGetResponseEnvelopeSuccess bool

const (
	RouteGetResponseEnvelopeSuccessTrue RouteGetResponseEnvelopeSuccess = true
)
