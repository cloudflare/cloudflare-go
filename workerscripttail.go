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

// Get list of tails currently deployed on a Worker.
func (r *WorkerScriptTailService) Get(ctx context.Context, scriptName string, query WorkerScriptTailGetParams, opts ...option.RequestOption) (res *WorkerScriptTailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r workerScriptTailNewResponseJSON) RawJSON() string {
	return r.raw
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

type WorkerScriptTailGetResponse struct {
	ID        interface{}                     `json:"id"`
	ExpiresAt interface{}                     `json:"expires_at"`
	URL       interface{}                     `json:"url"`
	JSON      workerScriptTailGetResponseJSON `json:"-"`
}

// workerScriptTailGetResponseJSON contains the JSON metadata for the struct
// [WorkerScriptTailGetResponse]
type workerScriptTailGetResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptTailGetResponseJSON) RawJSON() string {
	return r.raw
}

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

func (r workerScriptTailNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptTailNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptTailNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptTailNewResponseEnvelopeSuccess bool

const (
	WorkerScriptTailNewResponseEnvelopeSuccessTrue WorkerScriptTailNewResponseEnvelopeSuccess = true
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

func (r workerScriptTailDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptTailDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptTailDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptTailDeleteResponseEnvelopeSuccess bool

const (
	WorkerScriptTailDeleteResponseEnvelopeSuccessTrue WorkerScriptTailDeleteResponseEnvelopeSuccess = true
)

type WorkerScriptTailGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptTailGetResponseEnvelope struct {
	Errors   []WorkerScriptTailGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptTailGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptTailGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptTailGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptTailGetResponseEnvelopeJSON    `json:"-"`
}

// workerScriptTailGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptTailGetResponseEnvelope]
type workerScriptTailGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptTailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptTailGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerScriptTailGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptTailGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerScriptTailGetResponseEnvelopeErrors]
type workerScriptTailGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptTailGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptTailGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    workerScriptTailGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptTailGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerScriptTailGetResponseEnvelopeMessages]
type workerScriptTailGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptTailGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptTailGetResponseEnvelopeSuccess bool

const (
	WorkerScriptTailGetResponseEnvelopeSuccessTrue WorkerScriptTailGetResponseEnvelopeSuccess = true
)
