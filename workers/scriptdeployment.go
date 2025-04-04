// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ScriptDeploymentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptDeploymentService] method instead.
type ScriptDeploymentService struct {
	Options []option.RequestOption
}

// NewScriptDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScriptDeploymentService(opts ...option.RequestOption) (r *ScriptDeploymentService) {
	r = &ScriptDeploymentService{}
	r.Options = opts
	return
}

// Deployments configure how
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions)
// are deployed to traffic. A deployment can consist of one or two versions of a
// Worker.
func (r *ScriptDeploymentService) New(ctx context.Context, scriptName string, params ScriptDeploymentNewParams, opts ...option.RequestOption) (res *ScriptDeploymentNewResponse, err error) {
	var env ScriptDeploymentNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of Worker Deployments. The first deployment in the list is the latest
// deployment actively serving traffic.
func (r *ScriptDeploymentService) Get(ctx context.Context, scriptName string, query ScriptDeploymentGetParams, opts ...option.RequestOption) (res *ScriptDeploymentGetResponse, err error) {
	var env ScriptDeploymentGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Deployment struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage string         `json:"workers/message"`
	JSON           deploymentJSON `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
	WorkersMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

type DeploymentParam struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
}

func (r DeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentNewResponse struct {
	Strategy    ScriptDeploymentNewResponseStrategy  `json:"strategy,required"`
	Versions    []ScriptDeploymentNewResponseVersion `json:"versions,required"`
	ID          string                               `json:"id"`
	Annotations Deployment                           `json:"annotations"`
	AuthorEmail string                               `json:"author_email"`
	CreatedOn   string                               `json:"created_on"`
	Source      string                               `json:"source"`
	JSON        scriptDeploymentNewResponseJSON      `json:"-"`
}

// scriptDeploymentNewResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentNewResponse]
type scriptDeploymentNewResponseJSON struct {
	Strategy    apijson.Field
	Versions    apijson.Field
	ID          apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseStrategy string

const (
	ScriptDeploymentNewResponseStrategyPercentage ScriptDeploymentNewResponseStrategy = "percentage"
)

func (r ScriptDeploymentNewResponseStrategy) IsKnown() bool {
	switch r {
	case ScriptDeploymentNewResponseStrategyPercentage:
		return true
	}
	return false
}

type ScriptDeploymentNewResponseVersion struct {
	Percentage float64                                `json:"percentage,required"`
	VersionID  string                                 `json:"version_id,required"`
	JSON       scriptDeploymentNewResponseVersionJSON `json:"-"`
}

// scriptDeploymentNewResponseVersionJSON contains the JSON metadata for the struct
// [ScriptDeploymentNewResponseVersion]
type scriptDeploymentNewResponseVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseVersionJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponse struct {
	Deployments []ScriptDeploymentGetResponseDeployment `json:"deployments"`
	JSON        scriptDeploymentGetResponseJSON         `json:"-"`
}

// scriptDeploymentGetResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentGetResponse]
type scriptDeploymentGetResponseJSON struct {
	Deployments apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseDeployment struct {
	Strategy    ScriptDeploymentGetResponseDeploymentsStrategy  `json:"strategy,required"`
	Versions    []ScriptDeploymentGetResponseDeploymentsVersion `json:"versions,required"`
	ID          string                                          `json:"id"`
	Annotations Deployment                                      `json:"annotations"`
	AuthorEmail string                                          `json:"author_email"`
	CreatedOn   string                                          `json:"created_on"`
	Source      string                                          `json:"source"`
	JSON        scriptDeploymentGetResponseDeploymentJSON       `json:"-"`
}

// scriptDeploymentGetResponseDeploymentJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseDeployment]
type scriptDeploymentGetResponseDeploymentJSON struct {
	Strategy    apijson.Field
	Versions    apijson.Field
	ID          apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseDeploymentJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseDeploymentsStrategy string

const (
	ScriptDeploymentGetResponseDeploymentsStrategyPercentage ScriptDeploymentGetResponseDeploymentsStrategy = "percentage"
)

func (r ScriptDeploymentGetResponseDeploymentsStrategy) IsKnown() bool {
	switch r {
	case ScriptDeploymentGetResponseDeploymentsStrategyPercentage:
		return true
	}
	return false
}

type ScriptDeploymentGetResponseDeploymentsVersion struct {
	Percentage float64                                           `json:"percentage,required"`
	VersionID  string                                            `json:"version_id,required"`
	JSON       scriptDeploymentGetResponseDeploymentsVersionJSON `json:"-"`
}

// scriptDeploymentGetResponseDeploymentsVersionJSON contains the JSON metadata for
// the struct [ScriptDeploymentGetResponseDeploymentsVersion]
type scriptDeploymentGetResponseDeploymentsVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseDeploymentsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseDeploymentsVersionJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewParams struct {
	// Identifier
	AccountID param.Field[string]                             `path:"account_id,required"`
	Strategy  param.Field[ScriptDeploymentNewParamsStrategy]  `json:"strategy,required"`
	Versions  param.Field[[]ScriptDeploymentNewParamsVersion] `json:"versions,required"`
	// If set to true, the deployment will be created even if normally blocked by
	// something such rolling back to an older version when a secret has changed.
	Force       param.Field[bool]            `query:"force"`
	Annotations param.Field[DeploymentParam] `json:"annotations"`
}

func (r ScriptDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ScriptDeploymentNewParams]'s query parameters as
// `url.Values`.
func (r ScriptDeploymentNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeploymentNewParamsStrategy string

const (
	ScriptDeploymentNewParamsStrategyPercentage ScriptDeploymentNewParamsStrategy = "percentage"
)

func (r ScriptDeploymentNewParamsStrategy) IsKnown() bool {
	switch r {
	case ScriptDeploymentNewParamsStrategyPercentage:
		return true
	}
	return false
}

type ScriptDeploymentNewParamsVersion struct {
	Percentage param.Field[float64] `json:"percentage,required"`
	VersionID  param.Field[string]  `json:"version_id,required"`
}

func (r ScriptDeploymentNewParamsVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptDeploymentNewResponse                `json:"result"`
	JSON    scriptDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentNewResponseEnvelope]
type scriptDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptDeploymentNewResponseEnvelopeSuccess bool

const (
	ScriptDeploymentNewResponseEnvelopeSuccessTrue ScriptDeploymentNewResponseEnvelopeSuccess = true
)

func (r ScriptDeploymentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptDeploymentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptDeploymentGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptDeploymentGetResponse                `json:"result"`
	JSON    scriptDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseEnvelope]
type scriptDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptDeploymentGetResponseEnvelopeSuccess bool

const (
	ScriptDeploymentGetResponseEnvelopeSuccessTrue ScriptDeploymentGetResponseEnvelopeSuccess = true
)

func (r ScriptDeploymentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
