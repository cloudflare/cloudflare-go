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

// Creates a route that maps a URL pattern to a Worker.
func (r *WorkerRouteService) New(ctx context.Context, zoneID string, body WorkerRouteNewParams, opts ...option.RequestOption) (res *WorkerRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns routes for a zone.
func (r *WorkerRouteService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]WorkerRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteListResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

// Updates the URL pattern or Worker associated with a route.
func (r *WorkerRouteService) Replace(ctx context.Context, zoneID string, routeID string, body WorkerRouteReplaceParams, opts ...option.RequestOption) (res *WorkerRouteReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerRouteReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

type WorkerRouteReplaceResponse struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string                         `json:"script,required"`
	JSON   workerRouteReplaceResponseJSON `json:"-"`
}

// workerRouteReplaceResponseJSON contains the JSON metadata for the struct
// [WorkerRouteReplaceResponse]
type workerRouteReplaceResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteNewParams struct {
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

// Whether the API call was successful
type WorkerRouteNewResponseEnvelopeSuccess bool

const (
	WorkerRouteNewResponseEnvelopeSuccessTrue WorkerRouteNewResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type WorkerRouteListResponseEnvelopeSuccess bool

const (
	WorkerRouteListResponseEnvelopeSuccessTrue WorkerRouteListResponseEnvelopeSuccess = true
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

type WorkerRouteReplaceParams struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r WorkerRouteReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerRouteReplaceResponseEnvelope struct {
	Errors   []WorkerRouteReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerRouteReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerRouteReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerRouteReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerRouteReplaceResponseEnvelopeJSON    `json:"-"`
}

// workerRouteReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerRouteReplaceResponseEnvelope]
type workerRouteReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteReplaceResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerRouteReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// workerRouteReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerRouteReplaceResponseEnvelopeErrors]
type workerRouteReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerRouteReplaceResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerRouteReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// workerRouteReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerRouteReplaceResponseEnvelopeMessages]
type workerRouteReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerRouteReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerRouteReplaceResponseEnvelopeSuccess bool

const (
	WorkerRouteReplaceResponseEnvelopeSuccessTrue WorkerRouteReplaceResponseEnvelopeSuccess = true
)
