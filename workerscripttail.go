// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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

// Deletes a tail from a Worker.
func (r *WorkerScriptTailService) Delete(ctx context.Context, accountID string, scriptName string, id string, opts ...option.RequestOption) (res *WorkerScriptTailDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", accountID, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get list of tails currently deployed on a Worker.
func (r *WorkerScriptTailService) WorkerTailLogsListTails(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptTailWorkerTailLogsListTailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailWorkerTailLogsListTailsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *WorkerScriptTailService) WorkerTailLogsStartTail(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptTailWorkerTailLogsStartTailResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptTailWorkerTailLogsStartTailResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type WorkerScriptTailWorkerTailLogsListTailsResponse struct {
	ID        interface{}                                         `json:"id"`
	ExpiresAt interface{}                                         `json:"expires_at"`
	URL       interface{}                                         `json:"url"`
	JSON      workerScriptTailWorkerTailLogsListTailsResponseJSON `json:"-"`
}

// workerScriptTailWorkerTailLogsListTailsResponseJSON contains the JSON metadata
// for the struct [WorkerScriptTailWorkerTailLogsListTailsResponse]
type workerScriptTailWorkerTailLogsListTailsResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsListTailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailWorkerTailLogsStartTailResponse struct {
	ID        interface{}                                         `json:"id"`
	ExpiresAt interface{}                                         `json:"expires_at"`
	URL       interface{}                                         `json:"url"`
	JSON      workerScriptTailWorkerTailLogsStartTailResponseJSON `json:"-"`
}

// workerScriptTailWorkerTailLogsStartTailResponseJSON contains the JSON metadata
// for the struct [WorkerScriptTailWorkerTailLogsStartTailResponse]
type workerScriptTailWorkerTailLogsStartTailResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsStartTailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WorkerScriptTailWorkerTailLogsListTailsResponseEnvelope struct {
	Errors   []WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptTailWorkerTailLogsListTailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptTailWorkerTailLogsListTailsResponseEnvelopeJSON    `json:"-"`
}

// workerScriptTailWorkerTailLogsListTailsResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [WorkerScriptTailWorkerTailLogsListTailsResponseEnvelope]
type workerScriptTailWorkerTailLogsListTailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsListTailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    workerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrors]
type workerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    workerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessages]
type workerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeSuccess bool

const (
	WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeSuccessTrue WorkerScriptTailWorkerTailLogsListTailsResponseEnvelopeSuccess = true
)

type WorkerScriptTailWorkerTailLogsStartTailResponseEnvelope struct {
	Errors   []WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptTailWorkerTailLogsStartTailResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptTailWorkerTailLogsStartTailResponseEnvelopeJSON    `json:"-"`
}

// workerScriptTailWorkerTailLogsStartTailResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [WorkerScriptTailWorkerTailLogsStartTailResponseEnvelope]
type workerScriptTailWorkerTailLogsStartTailResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsStartTailResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    workerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrors]
type workerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    workerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessages]
type workerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeSuccess bool

const (
	WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeSuccessTrue WorkerScriptTailWorkerTailLogsStartTailResponseEnvelopeSuccess = true
)
