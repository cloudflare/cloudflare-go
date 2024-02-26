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

// WorkerDeploymentsByScriptService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewWorkerDeploymentsByScriptService] method instead.
type WorkerDeploymentsByScriptService struct {
	Options []option.RequestOption
}

// NewWorkerDeploymentsByScriptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerDeploymentsByScriptService(opts ...option.RequestOption) (r *WorkerDeploymentsByScriptService) {
	r = &WorkerDeploymentsByScriptService{}
	r.Options = opts
	return
}

// List Deployments
func (r *WorkerDeploymentsByScriptService) List(ctx context.Context, scriptID string, query WorkerDeploymentsByScriptListParams, opts ...option.RequestOption) (res *WorkerDeploymentsByScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDeploymentsByScriptListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s", query.AccountID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Deployment Detail
func (r *WorkerDeploymentsByScriptService) Detail(ctx context.Context, scriptID string, deploymentID string, query WorkerDeploymentsByScriptDetailParams, opts ...option.RequestOption) (res *WorkerDeploymentsByScriptDetailResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDeploymentsByScriptDetailResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s/detail/%s", query.AccountID, scriptID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDeploymentsByScriptListResponse struct {
	Items  []interface{}                             `json:"items"`
	Latest interface{}                               `json:"latest"`
	JSON   workerDeploymentsByScriptListResponseJSON `json:"-"`
}

// workerDeploymentsByScriptListResponseJSON contains the JSON metadata for the
// struct [WorkerDeploymentsByScriptListResponse]
type workerDeploymentsByScriptListResponseJSON struct {
	Items       apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentsByScriptDetailResponse struct {
	ID        string                                      `json:"id"`
	Metadata  interface{}                                 `json:"metadata"`
	Number    float64                                     `json:"number"`
	Resources interface{}                                 `json:"resources"`
	JSON      workerDeploymentsByScriptDetailResponseJSON `json:"-"`
}

// workerDeploymentsByScriptDetailResponseJSON contains the JSON metadata for the
// struct [WorkerDeploymentsByScriptDetailResponse]
type workerDeploymentsByScriptDetailResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentsByScriptListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDeploymentsByScriptListResponseEnvelope struct {
	Errors   []WorkerDeploymentsByScriptListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDeploymentsByScriptListResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDeploymentsByScriptListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDeploymentsByScriptListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDeploymentsByScriptListResponseEnvelopeJSON    `json:"-"`
}

// workerDeploymentsByScriptListResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerDeploymentsByScriptListResponseEnvelope]
type workerDeploymentsByScriptListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentsByScriptListResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    workerDeploymentsByScriptListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDeploymentsByScriptListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerDeploymentsByScriptListResponseEnvelopeErrors]
type workerDeploymentsByScriptListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentsByScriptListResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    workerDeploymentsByScriptListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDeploymentsByScriptListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerDeploymentsByScriptListResponseEnvelopeMessages]
type workerDeploymentsByScriptListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDeploymentsByScriptListResponseEnvelopeSuccess bool

const (
	WorkerDeploymentsByScriptListResponseEnvelopeSuccessTrue WorkerDeploymentsByScriptListResponseEnvelopeSuccess = true
)

type WorkerDeploymentsByScriptDetailParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDeploymentsByScriptDetailResponseEnvelope struct {
	Errors   []WorkerDeploymentsByScriptDetailResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDeploymentsByScriptDetailResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDeploymentsByScriptDetailResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDeploymentsByScriptDetailResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDeploymentsByScriptDetailResponseEnvelopeJSON    `json:"-"`
}

// workerDeploymentsByScriptDetailResponseEnvelopeJSON contains the JSON metadata
// for the struct [WorkerDeploymentsByScriptDetailResponseEnvelope]
type workerDeploymentsByScriptDetailResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptDetailResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentsByScriptDetailResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    workerDeploymentsByScriptDetailResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDeploymentsByScriptDetailResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerDeploymentsByScriptDetailResponseEnvelopeErrors]
type workerDeploymentsByScriptDetailResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptDetailResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentsByScriptDetailResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    workerDeploymentsByScriptDetailResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDeploymentsByScriptDetailResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerDeploymentsByScriptDetailResponseEnvelopeMessages]
type workerDeploymentsByScriptDetailResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentsByScriptDetailResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDeploymentsByScriptDetailResponseEnvelopeSuccess bool

const (
	WorkerDeploymentsByScriptDetailResponseEnvelopeSuccessTrue WorkerDeploymentsByScriptDetailResponseEnvelopeSuccess = true
)
