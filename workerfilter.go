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

// WorkerFilterService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerFilterService] method
// instead.
type WorkerFilterService struct {
	Options []option.RequestOption
}

// NewWorkerFilterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerFilterService(opts ...option.RequestOption) (r *WorkerFilterService) {
	r = &WorkerFilterService{}
	r.Options = opts
	return
}

// Create Filter
func (r *WorkerFilterService) New(ctx context.Context, params WorkerFilterNewParams, opts ...option.RequestOption) (res *WorkerFilterNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update Filter
func (r *WorkerFilterService) Update(ctx context.Context, filterID string, params WorkerFilterUpdateParams, opts ...option.RequestOption) (res *WorkerFilterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters/%s", params.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Filters
func (r *WorkerFilterService) List(ctx context.Context, query WorkerFilterListParams, opts ...option.RequestOption) (res *[]WorkerFilterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterListResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Filter
func (r *WorkerFilterService) Delete(ctx context.Context, filterID string, body WorkerFilterDeleteParams, opts ...option.RequestOption) (res *WorkerFilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters/%s", body.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerFilterNewResponse struct {
	// Identifier
	ID   string                      `json:"id,required"`
	JSON workerFilterNewResponseJSON `json:"-"`
}

// workerFilterNewResponseJSON contains the JSON metadata for the struct
// [WorkerFilterNewResponse]
type workerFilterNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterUpdateResponse struct {
	// Identifier
	ID      string                         `json:"id,required"`
	Enabled bool                           `json:"enabled,required"`
	Pattern string                         `json:"pattern,required"`
	JSON    workerFilterUpdateResponseJSON `json:"-"`
}

// workerFilterUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerFilterUpdateResponse]
type workerFilterUpdateResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterListResponse struct {
	// Identifier
	ID      string                       `json:"id,required"`
	Enabled bool                         `json:"enabled,required"`
	Pattern string                       `json:"pattern,required"`
	JSON    workerFilterListResponseJSON `json:"-"`
}

// workerFilterListResponseJSON contains the JSON metadata for the struct
// [WorkerFilterListResponse]
type workerFilterListResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterDeleteResponse struct {
	// Identifier
	ID   string                         `json:"id,required"`
	JSON workerFilterDeleteResponseJSON `json:"-"`
}

// workerFilterDeleteResponseJSON contains the JSON metadata for the struct
// [WorkerFilterDeleteResponse]
type workerFilterDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterNewParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r WorkerFilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerFilterNewResponseEnvelope struct {
	Errors   []WorkerFilterNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerFilterNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerFilterNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success WorkerFilterNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerFilterNewResponseEnvelopeJSON    `json:"-"`
}

// workerFilterNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerFilterNewResponseEnvelope]
type workerFilterNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    workerFilterNewResponseEnvelopeErrorsJSON `json:"-"`
}

// workerFilterNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerFilterNewResponseEnvelopeErrors]
type workerFilterNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerFilterNewResponseEnvelopeMessagesJSON `json:"-"`
}

// workerFilterNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerFilterNewResponseEnvelopeMessages]
type workerFilterNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerFilterNewResponseEnvelopeSuccess bool

const (
	WorkerFilterNewResponseEnvelopeSuccessTrue WorkerFilterNewResponseEnvelopeSuccess = true
)

type WorkerFilterUpdateParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r WorkerFilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerFilterUpdateResponseEnvelope struct {
	Errors   []WorkerFilterUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerFilterUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerFilterUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerFilterUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerFilterUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerFilterUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerFilterUpdateResponseEnvelope]
type workerFilterUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerFilterUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerFilterUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerFilterUpdateResponseEnvelopeErrors]
type workerFilterUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerFilterUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerFilterUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerFilterUpdateResponseEnvelopeMessages]
type workerFilterUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerFilterUpdateResponseEnvelopeSuccess bool

const (
	WorkerFilterUpdateResponseEnvelopeSuccessTrue WorkerFilterUpdateResponseEnvelopeSuccess = true
)

type WorkerFilterListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerFilterListResponseEnvelope struct {
	Errors   []WorkerFilterListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerFilterListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerFilterListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerFilterListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerFilterListResponseEnvelopeJSON    `json:"-"`
}

// workerFilterListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerFilterListResponseEnvelope]
type workerFilterListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerFilterListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerFilterListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerFilterListResponseEnvelopeErrors]
type workerFilterListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerFilterListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerFilterListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerFilterListResponseEnvelopeMessages]
type workerFilterListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerFilterListResponseEnvelopeSuccess bool

const (
	WorkerFilterListResponseEnvelopeSuccessTrue WorkerFilterListResponseEnvelopeSuccess = true
)

type WorkerFilterDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerFilterDeleteResponseEnvelope struct {
	Errors   []WorkerFilterDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerFilterDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerFilterDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success WorkerFilterDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerFilterDeleteResponseEnvelopeJSON    `json:"-"`
}

// workerFilterDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerFilterDeleteResponseEnvelope]
type workerFilterDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerFilterDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// workerFilterDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerFilterDeleteResponseEnvelopeErrors]
type workerFilterDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerFilterDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// workerFilterDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerFilterDeleteResponseEnvelopeMessages]
type workerFilterDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerFilterDeleteResponseEnvelopeSuccess bool

const (
	WorkerFilterDeleteResponseEnvelopeSuccessTrue WorkerFilterDeleteResponseEnvelopeSuccess = true
)
