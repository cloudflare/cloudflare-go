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

// WorkerRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerRouteService] method
// instead.
type WorkerRouteService struct {
	Options []option.RequestOption
}

// NewWorkerRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerRouteService(opts ...option.RequestOption) (r *WorkerRouteService) {
	r = &WorkerRouteService{}
	r.Options = opts
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *WorkerRouteService) Update(ctx context.Context, zoneID string, routeID string, body WorkerRouteUpdateParams, opts ...option.RequestOption) (res *WorkerRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a route.
func (r *WorkerRouteService) Delete(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *WorkerRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *WorkerRouteService) Get(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *WorkerRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a route that maps a URL pattern to a Worker.
func (r *WorkerRouteService) WorkerRoutesNewRoute(ctx context.Context, zoneID string, body WorkerRouteWorkerRoutesNewRouteParams, opts ...option.RequestOption) (res *WorkerRouteWorkerRoutesNewRouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteWorkerRoutesNewRouteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns routes for a zone.
func (r *WorkerRouteService) WorkerRoutesListRoutes(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]WorkerRouteWorkerRoutesListRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteWorkerRoutesListRoutesResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerRouteUpdateResponse struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                        `json:"script,required"`
	JSON   workerRouteUpdateResponseJSON `json:"-"`
}

// workerRouteUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerRouteUpdateResponse]
type workerRouteUpdateResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [WorkerRouteDeleteResponseUnknown] or [shared.UnionString].
type WorkerRouteDeleteResponse interface {
	ImplementsWorkerRouteDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerRouteDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WorkerRouteGetResponse struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                     `json:"script,required"`
	JSON   workerRouteGetResponseJSON `json:"-"`
}

// workerRouteGetResponseJSON contains the JSON metadata for the struct
// [WorkerRouteGetResponse]
type workerRouteGetResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [WorkerRouteWorkerRoutesNewRouteResponseUnknown] or
// [shared.UnionString].
type WorkerRouteWorkerRoutesNewRouteResponse interface {
	ImplementsWorkerRouteWorkerRoutesNewRouteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerRouteWorkerRoutesNewRouteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WorkerRouteWorkerRoutesListRoutesResponse struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                                        `json:"script,required"`
	JSON   workerRouteWorkerRoutesListRoutesResponseJSON `json:"-"`
}

// workerRouteWorkerRoutesListRoutesResponseJSON contains the JSON metadata for the
// struct [WorkerRouteWorkerRoutesListRoutesResponse]
type workerRouteWorkerRoutesListRoutesResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesListRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteUpdateParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r WorkerRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerRouteUpdateResponseEnvelope struct {
	Errors   []WorkerRouteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerRouteUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerRouteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerRouteUpdateResponseEnvelope]
type workerRouteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerRouteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerRouteUpdateResponseEnvelopeErrors]
type workerRouteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerRouteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerRouteUpdateResponseEnvelopeMessages]
type workerRouteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerRouteUpdateResponseEnvelopeSuccess bool

const (
	WorkerRouteUpdateResponseEnvelopeSuccessTrue WorkerRouteUpdateResponseEnvelopeSuccess = true
)

type WorkerRouteDeleteResponseEnvelope struct {
	Errors   []WorkerRouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerRouteDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// workerRouteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerRouteDeleteResponseEnvelope]
type workerRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerRouteDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerRouteDeleteResponseEnvelopeErrors]
type workerRouteDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerRouteDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerRouteDeleteResponseEnvelopeMessages]
type workerRouteDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerRouteDeleteResponseEnvelopeSuccess bool

const (
	WorkerRouteDeleteResponseEnvelopeSuccessTrue WorkerRouteDeleteResponseEnvelopeSuccess = true
)

type WorkerRouteGetResponseEnvelope struct {
	Errors   []WorkerRouteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerRouteGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteGetResponseEnvelopeJSON    `json:"-"`
}

// workerRouteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerRouteGetResponseEnvelope]
type workerRouteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    workerRouteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerRouteGetResponseEnvelopeErrors]
type workerRouteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerRouteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerRouteGetResponseEnvelopeMessages]
type workerRouteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerRouteGetResponseEnvelopeSuccess bool

const (
	WorkerRouteGetResponseEnvelopeSuccessTrue WorkerRouteGetResponseEnvelopeSuccess = true
)

type WorkerRouteWorkerRoutesNewRouteParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r WorkerRouteWorkerRoutesNewRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerRouteWorkerRoutesNewRouteResponseEnvelope struct {
	Errors   []WorkerRouteWorkerRoutesNewRouteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteWorkerRoutesNewRouteResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerRouteWorkerRoutesNewRouteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteWorkerRoutesNewRouteResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteWorkerRoutesNewRouteResponseEnvelopeJSON    `json:"-"`
}

// workerRouteWorkerRoutesNewRouteResponseEnvelopeJSON contains the JSON metadata
// for the struct [WorkerRouteWorkerRoutesNewRouteResponseEnvelope]
type workerRouteWorkerRoutesNewRouteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesNewRouteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteWorkerRoutesNewRouteResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    workerRouteWorkerRoutesNewRouteResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteWorkerRoutesNewRouteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerRouteWorkerRoutesNewRouteResponseEnvelopeErrors]
type workerRouteWorkerRoutesNewRouteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesNewRouteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteWorkerRoutesNewRouteResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    workerRouteWorkerRoutesNewRouteResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteWorkerRoutesNewRouteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerRouteWorkerRoutesNewRouteResponseEnvelopeMessages]
type workerRouteWorkerRoutesNewRouteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesNewRouteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerRouteWorkerRoutesNewRouteResponseEnvelopeSuccess bool

const (
	WorkerRouteWorkerRoutesNewRouteResponseEnvelopeSuccessTrue WorkerRouteWorkerRoutesNewRouteResponseEnvelopeSuccess = true
)

type WorkerRouteWorkerRoutesListRoutesResponseEnvelope struct {
	Errors   []WorkerRouteWorkerRoutesListRoutesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteWorkerRoutesListRoutesResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerRouteWorkerRoutesListRoutesResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteWorkerRoutesListRoutesResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteWorkerRoutesListRoutesResponseEnvelopeJSON    `json:"-"`
}

// workerRouteWorkerRoutesListRoutesResponseEnvelopeJSON contains the JSON metadata
// for the struct [WorkerRouteWorkerRoutesListRoutesResponseEnvelope]
type workerRouteWorkerRoutesListRoutesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesListRoutesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteWorkerRoutesListRoutesResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    workerRouteWorkerRoutesListRoutesResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteWorkerRoutesListRoutesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [WorkerRouteWorkerRoutesListRoutesResponseEnvelopeErrors]
type workerRouteWorkerRoutesListRoutesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesListRoutesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteWorkerRoutesListRoutesResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    workerRouteWorkerRoutesListRoutesResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteWorkerRoutesListRoutesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerRouteWorkerRoutesListRoutesResponseEnvelopeMessages]
type workerRouteWorkerRoutesListRoutesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteWorkerRoutesListRoutesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerRouteWorkerRoutesListRoutesResponseEnvelopeSuccess bool

const (
	WorkerRouteWorkerRoutesListRoutesResponseEnvelopeSuccessTrue WorkerRouteWorkerRoutesListRoutesResponseEnvelopeSuccess = true
)
