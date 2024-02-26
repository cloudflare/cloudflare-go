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

// WorkerDeploymentByScriptService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewWorkerDeploymentByScriptService] method instead.
type WorkerDeploymentByScriptService struct {
	Options []option.RequestOption
}

// NewWorkerDeploymentByScriptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerDeploymentByScriptService(opts ...option.RequestOption) (r *WorkerDeploymentByScriptService) {
	r = &WorkerDeploymentByScriptService{}
	r.Options = opts
	return
}

// List Deployments
func (r *WorkerDeploymentByScriptService) List(ctx context.Context, scriptID string, query WorkerDeploymentByScriptListParams, opts ...option.RequestOption) (res *WorkerDeploymentByScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDeploymentByScriptListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s", query.AccountID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Deployment Detail
func (r *WorkerDeploymentByScriptService) Get(ctx context.Context, scriptID string, deploymentID string, query WorkerDeploymentByScriptGetParams, opts ...option.RequestOption) (res *WorkerDeploymentByScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDeploymentByScriptGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s/detail/%s", query.AccountID, scriptID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDeploymentByScriptListResponse struct {
	Items  []interface{}                            `json:"items"`
	Latest interface{}                              `json:"latest"`
	JSON   workerDeploymentByScriptListResponseJSON `json:"-"`
}

// workerDeploymentByScriptListResponseJSON contains the JSON metadata for the
// struct [WorkerDeploymentByScriptListResponse]
type workerDeploymentByScriptListResponseJSON struct {
	Items       apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptGetResponse struct {
	ID        string                                  `json:"id"`
	Metadata  interface{}                             `json:"metadata"`
	Number    float64                                 `json:"number"`
	Resources interface{}                             `json:"resources"`
	JSON      workerDeploymentByScriptGetResponseJSON `json:"-"`
}

// workerDeploymentByScriptGetResponseJSON contains the JSON metadata for the
// struct [WorkerDeploymentByScriptGetResponse]
type workerDeploymentByScriptGetResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDeploymentByScriptListResponseEnvelope struct {
	Errors   []WorkerDeploymentByScriptListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDeploymentByScriptListResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDeploymentByScriptListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDeploymentByScriptListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDeploymentByScriptListResponseEnvelopeJSON    `json:"-"`
}

// workerDeploymentByScriptListResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerDeploymentByScriptListResponseEnvelope]
type workerDeploymentByScriptListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    workerDeploymentByScriptListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDeploymentByScriptListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerDeploymentByScriptListResponseEnvelopeErrors]
type workerDeploymentByScriptListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    workerDeploymentByScriptListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDeploymentByScriptListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerDeploymentByScriptListResponseEnvelopeMessages]
type workerDeploymentByScriptListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDeploymentByScriptListResponseEnvelopeSuccess bool

const (
	WorkerDeploymentByScriptListResponseEnvelopeSuccessTrue WorkerDeploymentByScriptListResponseEnvelopeSuccess = true
)

type WorkerDeploymentByScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDeploymentByScriptGetResponseEnvelope struct {
	Errors   []WorkerDeploymentByScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDeploymentByScriptGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDeploymentByScriptGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDeploymentByScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDeploymentByScriptGetResponseEnvelopeJSON    `json:"-"`
}

// workerDeploymentByScriptGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerDeploymentByScriptGetResponseEnvelope]
type workerDeploymentByScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerDeploymentByScriptGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDeploymentByScriptGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerDeploymentByScriptGetResponseEnvelopeErrors]
type workerDeploymentByScriptGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    workerDeploymentByScriptGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDeploymentByScriptGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerDeploymentByScriptGetResponseEnvelopeMessages]
type workerDeploymentByScriptGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDeploymentByScriptGetResponseEnvelopeSuccess bool

const (
	WorkerDeploymentByScriptGetResponseEnvelopeSuccessTrue WorkerDeploymentByScriptGetResponseEnvelopeSuccess = true
)
