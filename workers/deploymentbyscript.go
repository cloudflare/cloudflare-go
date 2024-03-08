// File generated from our OpenAPI spec by Stainless.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// DeploymentByScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeploymentByScriptService] method
// instead.
type DeploymentByScriptService struct {
	Options []option.RequestOption
	Details *DeploymentByScriptDetailService
}

// NewDeploymentByScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeploymentByScriptService(opts ...option.RequestOption) (r *DeploymentByScriptService) {
	r = &DeploymentByScriptService{}
	r.Options = opts
	r.Details = NewDeploymentByScriptDetailService(opts...)
	return
}

// List Deployments
func (r *DeploymentByScriptService) Get(ctx context.Context, scriptID string, query DeploymentByScriptGetParams, opts ...option.RequestOption) (res *DeploymentByScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeploymentByScriptGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s", query.AccountID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeploymentByScriptGetResponse struct {
	Items  []interface{}                     `json:"items"`
	Latest interface{}                       `json:"latest"`
	JSON   deploymentByScriptGetResponseJSON `json:"-"`
}

// deploymentByScriptGetResponseJSON contains the JSON metadata for the struct
// [DeploymentByScriptGetResponse]
type deploymentByScriptGetResponseJSON struct {
	Items       apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentByScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentByScriptGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeploymentByScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeploymentByScriptGetResponseEnvelope struct {
	Errors   []DeploymentByScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeploymentByScriptGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DeploymentByScriptGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DeploymentByScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deploymentByScriptGetResponseEnvelopeJSON    `json:"-"`
}

// deploymentByScriptGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeploymentByScriptGetResponseEnvelope]
type deploymentByScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentByScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentByScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeploymentByScriptGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deploymentByScriptGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deploymentByScriptGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeploymentByScriptGetResponseEnvelopeErrors]
type deploymentByScriptGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentByScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentByScriptGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeploymentByScriptGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    deploymentByScriptGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deploymentByScriptGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeploymentByScriptGetResponseEnvelopeMessages]
type deploymentByScriptGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentByScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentByScriptGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DeploymentByScriptGetResponseEnvelopeSuccess bool

const (
	DeploymentByScriptGetResponseEnvelopeSuccessTrue DeploymentByScriptGetResponseEnvelopeSuccess = true
)
