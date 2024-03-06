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

// WorkerDeploymentByScriptDetailService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkerDeploymentByScriptDetailService] method instead.
type WorkerDeploymentByScriptDetailService struct {
	Options []option.RequestOption
}

// NewWorkerDeploymentByScriptDetailService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerDeploymentByScriptDetailService(opts ...option.RequestOption) (r *WorkerDeploymentByScriptDetailService) {
	r = &WorkerDeploymentByScriptDetailService{}
	r.Options = opts
	return
}

// Get Deployment Detail
func (r *WorkerDeploymentByScriptDetailService) Get(ctx context.Context, scriptID string, deploymentID string, query WorkerDeploymentByScriptDetailGetParams, opts ...option.RequestOption) (res *WorkerDeploymentByScriptDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDeploymentByScriptDetailGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s/detail/%s", query.AccountID, scriptID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDeploymentByScriptDetailGetResponse struct {
	ID        string                                        `json:"id"`
	Metadata  interface{}                                   `json:"metadata"`
	Number    float64                                       `json:"number"`
	Resources interface{}                                   `json:"resources"`
	JSON      workerDeploymentByScriptDetailGetResponseJSON `json:"-"`
}

// workerDeploymentByScriptDetailGetResponseJSON contains the JSON metadata for the
// struct [WorkerDeploymentByScriptDetailGetResponse]
type workerDeploymentByScriptDetailGetResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerDeploymentByScriptDetailGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerDeploymentByScriptDetailGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDeploymentByScriptDetailGetResponseEnvelope struct {
	Errors   []WorkerDeploymentByScriptDetailGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDeploymentByScriptDetailGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDeploymentByScriptDetailGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDeploymentByScriptDetailGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDeploymentByScriptDetailGetResponseEnvelopeJSON    `json:"-"`
}

// workerDeploymentByScriptDetailGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [WorkerDeploymentByScriptDetailGetResponseEnvelope]
type workerDeploymentByScriptDetailGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptDetailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerDeploymentByScriptDetailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerDeploymentByScriptDetailGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    workerDeploymentByScriptDetailGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDeploymentByScriptDetailGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [WorkerDeploymentByScriptDetailGetResponseEnvelopeErrors]
type workerDeploymentByScriptDetailGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptDetailGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerDeploymentByScriptDetailGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerDeploymentByScriptDetailGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    workerDeploymentByScriptDetailGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDeploymentByScriptDetailGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerDeploymentByScriptDetailGetResponseEnvelopeMessages]
type workerDeploymentByScriptDetailGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptDetailGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerDeploymentByScriptDetailGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerDeploymentByScriptDetailGetResponseEnvelopeSuccess bool

const (
	WorkerDeploymentByScriptDetailGetResponseEnvelopeSuccessTrue WorkerDeploymentByScriptDetailGetResponseEnvelopeSuccess = true
)
