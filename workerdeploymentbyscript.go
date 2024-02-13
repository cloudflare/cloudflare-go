// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
	Details *WorkerDeploymentByScriptDetailService
}

// NewWorkerDeploymentByScriptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerDeploymentByScriptService(opts ...option.RequestOption) (r *WorkerDeploymentByScriptService) {
	r = &WorkerDeploymentByScriptService{}
	r.Options = opts
	r.Details = NewWorkerDeploymentByScriptDetailService(opts...)
	return
}

// List Deployments
func (r *WorkerDeploymentByScriptService) WorkerDeploymentsListDeployments(ctx context.Context, accountID string, scriptID string, opts ...option.RequestOption) (res *WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s", accountID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse struct {
	Items  []interface{}                                                        `json:"items"`
	Latest interface{}                                                          `json:"latest"`
	JSON   workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseJSON `json:"-"`
}

// workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseJSON contains
// the JSON metadata for the struct
// [WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse]
type workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseJSON struct {
	Items       apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelope struct {
	Errors   []WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeJSON    `json:"-"`
}

// workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelope]
type workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrors]
type workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessages]
type workerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeSuccess bool

const (
	WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeSuccessTrue WorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseEnvelopeSuccess = true
)
