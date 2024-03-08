// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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

// Creates a route that maps a URL pattern to a Worker.
func (r *WorkerRouteService) New(ctx context.Context, params WorkerRouteNewParams, opts ...option.RequestOption) (res *WorkerRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *WorkerRouteService) Update(ctx context.Context, routeID string, params WorkerRouteUpdateParams, opts ...option.RequestOption) (res *WorkerRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", params.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns routes for a zone.
func (r *WorkerRouteService) List(ctx context.Context, query WorkerRouteListParams, opts ...option.RequestOption) (res *[]WorkerRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteListResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a route.
func (r *WorkerRouteService) Delete(ctx context.Context, routeID string, body WorkerRouteDeleteParams, opts ...option.RequestOption) (res *WorkerRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", body.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *WorkerRouteService) Get(ctx context.Context, routeID string, query WorkerRouteGetParams, opts ...option.RequestOption) (res *WorkerRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", query.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [WorkerRouteNewResponseUnknown] or [shared.UnionString].
type WorkerRouteNewResponse interface {
	ImplementsWorkerRouteNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerRouteNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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

func (r workerRouteUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerRouteListResponse struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                      `json:"script,required"`
	JSON   workerRouteListResponseJSON `json:"-"`
}

// workerRouteListResponseJSON contains the JSON metadata for the struct
// [WorkerRouteListResponse]
type workerRouteListResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteListResponseJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerRouteNewParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r WorkerRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerRouteNewResponseEnvelope struct {
	Errors   []WorkerRouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerRouteNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteNewResponseEnvelopeJSON    `json:"-"`
}

// workerRouteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerRouteNewResponseEnvelope]
type workerRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerRouteNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    workerRouteNewResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerRouteNewResponseEnvelopeErrors]
type workerRouteNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerRouteNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerRouteNewResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerRouteNewResponseEnvelopeMessages]
type workerRouteNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerRouteNewResponseEnvelopeSuccess bool

const (
	WorkerRouteNewResponseEnvelopeSuccessTrue WorkerRouteNewResponseEnvelopeSuccess = true
)

type WorkerRouteUpdateParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
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

func (r workerRouteUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerRouteUpdateResponseEnvelopeSuccess bool

const (
	WorkerRouteUpdateResponseEnvelopeSuccessTrue WorkerRouteUpdateResponseEnvelopeSuccess = true
)

type WorkerRouteListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerRouteListResponseEnvelope struct {
	Errors   []WorkerRouteListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerRouteListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteListResponseEnvelopeJSON    `json:"-"`
}

// workerRouteListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerRouteListResponseEnvelope]
type workerRouteListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerRouteListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    workerRouteListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerRouteListResponseEnvelopeErrors]
type workerRouteListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerRouteListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerRouteListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerRouteListResponseEnvelopeMessages]
type workerRouteListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerRouteListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerRouteListResponseEnvelopeSuccess bool

const (
	WorkerRouteListResponseEnvelopeSuccessTrue WorkerRouteListResponseEnvelopeSuccess = true
)

type WorkerRouteDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

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

func (r workerRouteDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerRouteDeleteResponseEnvelopeSuccess bool

const (
	WorkerRouteDeleteResponseEnvelopeSuccessTrue WorkerRouteDeleteResponseEnvelopeSuccess = true
)

type WorkerRouteGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

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

func (r workerRouteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerRouteGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerRouteGetResponseEnvelopeSuccess bool

const (
	WorkerRouteGetResponseEnvelopeSuccessTrue WorkerRouteGetResponseEnvelopeSuccess = true
)
