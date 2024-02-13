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

// Update Filter
func (r *WorkerFilterService) Update(ctx context.Context, zoneID string, filterID string, body WorkerFilterUpdateParams, opts ...option.RequestOption) (res *WorkerFilterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Filter
func (r *WorkerFilterService) Delete(ctx context.Context, zoneID string, filterID string, opts ...option.RequestOption) (res *WorkerFilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create Filter
func (r *WorkerFilterService) WorkerFiltersDeprecatedNewFilter(ctx context.Context, zoneID string, body WorkerFilterWorkerFiltersDeprecatedNewFilterParams, opts ...option.RequestOption) (res *WorkerFilterWorkerFiltersDeprecatedNewFilterResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Filters
func (r *WorkerFilterService) WorkerFiltersDeprecatedListFilters(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]WorkerFilterWorkerFiltersDeprecatedListFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type WorkerFilterWorkerFiltersDeprecatedNewFilterResponse struct {
	// Identifier
	ID   string                                                   `json:"id,required"`
	JSON workerFilterWorkerFiltersDeprecatedNewFilterResponseJSON `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedNewFilterResponseJSON contains the JSON
// metadata for the struct [WorkerFilterWorkerFiltersDeprecatedNewFilterResponse]
type workerFilterWorkerFiltersDeprecatedNewFilterResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedNewFilterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterWorkerFiltersDeprecatedListFiltersResponse struct {
	// Identifier
	ID      string                                                     `json:"id,required"`
	Enabled bool                                                       `json:"enabled,required"`
	Pattern string                                                     `json:"pattern,required"`
	JSON    workerFilterWorkerFiltersDeprecatedListFiltersResponseJSON `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedListFiltersResponseJSON contains the JSON
// metadata for the struct [WorkerFilterWorkerFiltersDeprecatedListFiltersResponse]
type workerFilterWorkerFiltersDeprecatedListFiltersResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedListFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterUpdateParams struct {
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

type WorkerFilterWorkerFiltersDeprecatedNewFilterParams struct {
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r WorkerFilterWorkerFiltersDeprecatedNewFilterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelope struct {
	Errors   []WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerFilterWorkerFiltersDeprecatedNewFilterResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeJSON    `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelope]
type workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrorsJSON `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrors]
type workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessagesJSON `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessages]
type workerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeSuccess bool

const (
	WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeSuccessTrue WorkerFilterWorkerFiltersDeprecatedNewFilterResponseEnvelopeSuccess = true
)

type WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelope struct {
	Errors   []WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerFilterWorkerFiltersDeprecatedListFiltersResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeJSON    `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelope]
type workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrorsJSON `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrors]
type workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessagesJSON `json:"-"`
}

// workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessages]
type workerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeSuccess bool

const (
	WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeSuccessTrue WorkerFilterWorkerFiltersDeprecatedListFiltersResponseEnvelopeSuccess = true
)
