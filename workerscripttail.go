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

// WorkerScriptTailService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptTailService] method
// instead.
type WorkerScriptTailService struct {
	Options []option.RequestOption
}

// NewWorkerScriptTailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptTailService(opts ...option.RequestOption) (r *WorkerScriptTailService) {
	r = &WorkerScriptTailService{}
	r.Options = opts
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *WorkerScriptTailService) New(ctx context.Context, scriptName string, body WorkerScriptTailNewParams, opts ...option.RequestOption) (res *WorkerScriptTailNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", body.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get list of tails currently deployed on a Worker.
func (r *WorkerScriptTailService) List(ctx context.Context, scriptName string, query WorkerScriptTailListParams, opts ...option.RequestOption) (res *WorkerScriptTailListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a tail from a Worker.
func (r *WorkerScriptTailService) Delete(ctx context.Context, scriptName string, id string, body WorkerScriptTailDeleteParams, opts ...option.RequestOption) (res *WorkerScriptTailDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", body.AccountID, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptTailNewResponse struct {
	ID        interface{}                     `json:"id"`
	ExpiresAt interface{}                     `json:"expires_at"`
	URL       interface{}                     `json:"url"`
	JSON      workerScriptTailNewResponseJSON `json:"-"`
}

// workerScriptTailNewResponseJSON contains the JSON metadata for the struct
// [WorkerScriptTailNewResponse]
type workerScriptTailNewResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailListResponse struct {
	ID        interface{}                      `json:"id"`
	ExpiresAt interface{}                      `json:"expires_at"`
	URL       interface{}                      `json:"url"`
	JSON      workerScriptTailListResponseJSON `json:"-"`
}

// workerScriptTailListResponseJSON contains the JSON metadata for the struct
// [WorkerScriptTailListResponse]
type workerScriptTailListResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [WorkerScriptTailDeleteResponseUnknown],
// [WorkerScriptTailDeleteResponseArray] or [shared.UnionString].
type WorkerScriptTailDeleteResponse interface {
	ImplementsWorkerScriptTailDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerScriptTailDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WorkerScriptTailDeleteResponseArray []interface{}

func (r WorkerScriptTailDeleteResponseArray) ImplementsWorkerScriptTailDeleteResponse() {}

type WorkerScriptTailNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptTailNewResponseEnvelope struct {
	Errors   []WorkerScriptTailNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptTailNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptTailNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptTailNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptTailNewResponseEnvelopeJSON    `json:"-"`
}

// workerScriptTailNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptTailNewResponseEnvelope]
type workerScriptTailNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerScriptTailNewResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptTailNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerScriptTailNewResponseEnvelopeErrors]
type workerScriptTailNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    workerScriptTailNewResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptTailNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerScriptTailNewResponseEnvelopeMessages]
type workerScriptTailNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptTailNewResponseEnvelopeSuccess bool

const (
	WorkerScriptTailNewResponseEnvelopeSuccessTrue WorkerScriptTailNewResponseEnvelopeSuccess = true
)

type WorkerScriptTailListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptTailListResponseEnvelope struct {
	Errors   []WorkerScriptTailListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptTailListResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptTailListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptTailListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptTailListResponseEnvelopeJSON    `json:"-"`
}

// workerScriptTailListResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptTailListResponseEnvelope]
type workerScriptTailListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerScriptTailListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptTailListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptTailListResponseEnvelopeErrors]
type workerScriptTailListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    workerScriptTailListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptTailListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerScriptTailListResponseEnvelopeMessages]
type workerScriptTailListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptTailListResponseEnvelopeSuccess bool

const (
	WorkerScriptTailListResponseEnvelopeSuccessTrue WorkerScriptTailListResponseEnvelopeSuccess = true
)

type WorkerScriptTailDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptTailDeleteResponseEnvelope struct {
	Errors   []WorkerScriptTailDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptTailDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptTailDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptTailDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptTailDeleteResponseEnvelopeJSON    `json:"-"`
}

// workerScriptTailDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptTailDeleteResponseEnvelope]
type workerScriptTailDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    workerScriptTailDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptTailDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptTailDeleteResponseEnvelopeErrors]
type workerScriptTailDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerScriptTailDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptTailDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptTailDeleteResponseEnvelopeMessages]
type workerScriptTailDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptTailDeleteResponseEnvelopeSuccess bool

const (
	WorkerScriptTailDeleteResponseEnvelopeSuccessTrue WorkerScriptTailDeleteResponseEnvelopeSuccess = true
)
