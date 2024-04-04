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

// ScriptDeploymentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptDeploymentService] method
// instead.
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
	opts = append(r.Options[:], opts...)
	var env ScriptDeploymentNewResponseEnvelope
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
	opts = append(r.Options[:], opts...)
	var env ScriptDeploymentGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8 struct {
	// Human-readable message about the deployment.
	WorkersMessage string                                               `json:"workers/message"`
	JSON           unnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8JSON `json:"-"`
}

// unnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8]
type unnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8JSON struct {
	WorkersMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8Param struct {
	// Human-readable message about the deployment.
	WorkersMessage param.Field[string] `json:"workers/message"`
}

func (r UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentNewResponse struct {
	ID          string                                           `json:"id"`
	Annotations UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8 `json:"annotations"`
	AuthorEmail string                                           `json:"author_email"`
	CreatedOn   string                                           `json:"created_on"`
	Source      string                                           `json:"source"`
	Strategy    string                                           `json:"strategy"`
	JSON        scriptDeploymentNewResponseJSON                  `json:"-"`
}

// scriptDeploymentNewResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentNewResponse]
type scriptDeploymentNewResponseJSON struct {
	ID          apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseJSON) RawJSON() string {
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
	ID          string                                           `json:"id"`
	Annotations UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8 `json:"annotations"`
	AuthorEmail string                                           `json:"author_email"`
	CreatedOn   string                                           `json:"created_on"`
	Source      string                                           `json:"source"`
	Strategy    string                                           `json:"strategy"`
	JSON        scriptDeploymentGetResponseDeploymentJSON        `json:"-"`
}

// scriptDeploymentGetResponseDeploymentJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseDeployment]
type scriptDeploymentGetResponseDeploymentJSON struct {
	ID          apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseDeploymentJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewParams struct {
	// Identifier
	AccountID   param.Field[string]                                                `path:"account_id,required"`
	Annotations param.Field[UnnamedSchemaRefFda1c6f6758e763ae3b2964521f2fdd8Param] `json:"annotations"`
	Strategy    param.Field[string]                                                `json:"strategy"`
}

func (r ScriptDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ScriptDeploymentNewResponse                               `json:"result,required"`
	// Whether the API call was successful
	Success ScriptDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentNewResponseEnvelope]
type scriptDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ScriptDeploymentGetResponse                               `json:"result,required"`
	// Whether the API call was successful
	Success ScriptDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseEnvelope]
type scriptDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
