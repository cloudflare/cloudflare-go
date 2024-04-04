// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DeploymentByScriptDetailService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDeploymentByScriptDetailService] method instead.
type DeploymentByScriptDetailService struct {
	Options []option.RequestOption
}

// NewDeploymentByScriptDetailService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDeploymentByScriptDetailService(opts ...option.RequestOption) (r *DeploymentByScriptDetailService) {
	r = &DeploymentByScriptDetailService{}
	r.Options = opts
	return
}

// Get Deployment Detail
func (r *DeploymentByScriptDetailService) Get(ctx context.Context, scriptID string, deploymentID string, query DeploymentByScriptDetailGetParams, opts ...option.RequestOption) (res *DeploymentByScriptDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeploymentByScriptDetailGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s/detail/%s", query.AccountID, scriptID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeploymentByScriptDetailGetResponse struct {
	ID        string                                  `json:"id"`
	Metadata  interface{}                             `json:"metadata"`
	Number    float64                                 `json:"number"`
	Resources interface{}                             `json:"resources"`
	JSON      deploymentByScriptDetailGetResponseJSON `json:"-"`
}

// deploymentByScriptDetailGetResponseJSON contains the JSON metadata for the
// struct [DeploymentByScriptDetailGetResponse]
type deploymentByScriptDetailGetResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentByScriptDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentByScriptDetailGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeploymentByScriptDetailGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeploymentByScriptDetailGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DeploymentByScriptDetailGetResponse                       `json:"result,required"`
	// Whether the API call was successful
	Success DeploymentByScriptDetailGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deploymentByScriptDetailGetResponseEnvelopeJSON    `json:"-"`
}

// deploymentByScriptDetailGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DeploymentByScriptDetailGetResponseEnvelope]
type deploymentByScriptDetailGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentByScriptDetailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentByScriptDetailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DeploymentByScriptDetailGetResponseEnvelopeSuccess bool

const (
	DeploymentByScriptDetailGetResponseEnvelopeSuccessTrue DeploymentByScriptDetailGetResponseEnvelopeSuccess = true
)

func (r DeploymentByScriptDetailGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeploymentByScriptDetailGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
