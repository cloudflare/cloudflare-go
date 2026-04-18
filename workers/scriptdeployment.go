// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List of Worker Deployments. The first deployment in the list is the latest
// deployment actively serving traffic.
func (r *ScriptDeploymentService) List(ctx context.Context, scriptName string, query ScriptDeploymentListParams, opts ...option.RequestOption) (res *ScriptDeploymentListResponse, err error) {
	var env ScriptDeploymentListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete a Worker Deployment. The latest deployment, which is actively serving
// traffic, cannot be deleted. All other deployments can be deleted.
func (r *ScriptDeploymentService) Delete(ctx context.Context, scriptName string, deploymentID string, body ScriptDeploymentDeleteParams, opts ...option.RequestOption) (res *ScriptDeploymentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments/%s", body.AccountID, scriptName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get information about a Worker Deployment.
func (r *ScriptDeploymentService) Get(ctx context.Context, scriptName string, deploymentID string, query ScriptDeploymentGetParams, opts ...option.RequestOption) (res *ScriptDeploymentGetResponse, err error) {
	var env ScriptDeploymentGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments/%s", query.AccountID, scriptName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ScriptDeploymentNewResponse struct {
	ID          string                                 `json:"id" api:"required" format:"uuid"`
	CreatedOn   time.Time                              `json:"created_on" api:"required" format:"date-time"`
	Source      string                                 `json:"source" api:"required"`
	Strategy    ScriptDeploymentNewResponseStrategy    `json:"strategy" api:"required"`
	Versions    []ScriptDeploymentNewResponseVersion   `json:"versions" api:"required"`
	Annotations ScriptDeploymentNewResponseAnnotations `json:"annotations"`
	AuthorEmail string                                 `json:"author_email" format:"email"`
	JSON        scriptDeploymentNewResponseJSON        `json:"-"`
}

// scriptDeploymentNewResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentNewResponse]
type scriptDeploymentNewResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	Versions    apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
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
	Percentage float64                                `json:"percentage" api:"required"`
	VersionID  string                                 `json:"version_id" api:"required" format:"uuid"`
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

type ScriptDeploymentNewResponseAnnotations struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// Operation that triggered the creation of the deployment.
	WorkersTriggeredBy string                                     `json:"workers/triggered_by"`
	JSON               scriptDeploymentNewResponseAnnotationsJSON `json:"-"`
}

// scriptDeploymentNewResponseAnnotationsJSON contains the JSON metadata for the
// struct [ScriptDeploymentNewResponseAnnotations]
type scriptDeploymentNewResponseAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseAnnotationsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponse struct {
	Deployments []ScriptDeploymentListResponseDeployment `json:"deployments" api:"required"`
	JSON        scriptDeploymentListResponseJSON         `json:"-"`
}

// scriptDeploymentListResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentListResponse]
type scriptDeploymentListResponseJSON struct {
	Deployments apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseDeployment struct {
	ID          string                                             `json:"id" api:"required" format:"uuid"`
	CreatedOn   time.Time                                          `json:"created_on" api:"required" format:"date-time"`
	Source      string                                             `json:"source" api:"required"`
	Strategy    ScriptDeploymentListResponseDeploymentsStrategy    `json:"strategy" api:"required"`
	Versions    []ScriptDeploymentListResponseDeploymentsVersion   `json:"versions" api:"required"`
	Annotations ScriptDeploymentListResponseDeploymentsAnnotations `json:"annotations"`
	AuthorEmail string                                             `json:"author_email" format:"email"`
	JSON        scriptDeploymentListResponseDeploymentJSON         `json:"-"`
}

// scriptDeploymentListResponseDeploymentJSON contains the JSON metadata for the
// struct [ScriptDeploymentListResponseDeployment]
type scriptDeploymentListResponseDeploymentJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	Versions    apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseDeploymentJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseDeploymentsStrategy string

const (
	ScriptDeploymentListResponseDeploymentsStrategyPercentage ScriptDeploymentListResponseDeploymentsStrategy = "percentage"
)

func (r ScriptDeploymentListResponseDeploymentsStrategy) IsKnown() bool {
	switch r {
	case ScriptDeploymentListResponseDeploymentsStrategyPercentage:
		return true
	}
	return false
}

type ScriptDeploymentListResponseDeploymentsVersion struct {
	Percentage float64                                            `json:"percentage" api:"required"`
	VersionID  string                                             `json:"version_id" api:"required" format:"uuid"`
	JSON       scriptDeploymentListResponseDeploymentsVersionJSON `json:"-"`
}

// scriptDeploymentListResponseDeploymentsVersionJSON contains the JSON metadata
// for the struct [ScriptDeploymentListResponseDeploymentsVersion]
type scriptDeploymentListResponseDeploymentsVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseDeploymentsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseDeploymentsVersionJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseDeploymentsAnnotations struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// Operation that triggered the creation of the deployment.
	WorkersTriggeredBy string                                                 `json:"workers/triggered_by"`
	JSON               scriptDeploymentListResponseDeploymentsAnnotationsJSON `json:"-"`
}

// scriptDeploymentListResponseDeploymentsAnnotationsJSON contains the JSON
// metadata for the struct [ScriptDeploymentListResponseDeploymentsAnnotations]
type scriptDeploymentListResponseDeploymentsAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseDeploymentsAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseDeploymentsAnnotationsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponse struct {
	Errors   []ScriptDeploymentDeleteResponseError   `json:"errors" api:"required"`
	Messages []ScriptDeploymentDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ScriptDeploymentDeleteResponseSuccess `json:"success" api:"required"`
	JSON    scriptDeploymentDeleteResponseJSON    `json:"-"`
}

// scriptDeploymentDeleteResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentDeleteResponse]
type scriptDeploymentDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseError struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScriptDeploymentDeleteResponseErrorsSource `json:"source"`
	JSON             scriptDeploymentDeleteResponseErrorJSON    `json:"-"`
}

// scriptDeploymentDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ScriptDeploymentDeleteResponseError]
type scriptDeploymentDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scriptDeploymentDeleteResponseErrorsSourceJSON `json:"-"`
}

// scriptDeploymentDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptDeploymentDeleteResponseErrorsSource]
type scriptDeploymentDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseMessage struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptDeploymentDeleteResponseMessagesSource `json:"source"`
	JSON             scriptDeploymentDeleteResponseMessageJSON    `json:"-"`
}

// scriptDeploymentDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ScriptDeploymentDeleteResponseMessage]
type scriptDeploymentDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptDeploymentDeleteResponseMessagesSourceJSON `json:"-"`
}

// scriptDeploymentDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptDeploymentDeleteResponseMessagesSource]
type scriptDeploymentDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeploymentDeleteResponseSuccess bool

const (
	ScriptDeploymentDeleteResponseSuccessTrue ScriptDeploymentDeleteResponseSuccess = true
)

func (r ScriptDeploymentDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ScriptDeploymentGetResponse struct {
	ID          string                                 `json:"id" api:"required" format:"uuid"`
	CreatedOn   time.Time                              `json:"created_on" api:"required" format:"date-time"`
	Source      string                                 `json:"source" api:"required"`
	Strategy    ScriptDeploymentGetResponseStrategy    `json:"strategy" api:"required"`
	Versions    []ScriptDeploymentGetResponseVersion   `json:"versions" api:"required"`
	Annotations ScriptDeploymentGetResponseAnnotations `json:"annotations"`
	AuthorEmail string                                 `json:"author_email" format:"email"`
	JSON        scriptDeploymentGetResponseJSON        `json:"-"`
}

// scriptDeploymentGetResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentGetResponse]
type scriptDeploymentGetResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	Versions    apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseStrategy string

const (
	ScriptDeploymentGetResponseStrategyPercentage ScriptDeploymentGetResponseStrategy = "percentage"
)

func (r ScriptDeploymentGetResponseStrategy) IsKnown() bool {
	switch r {
	case ScriptDeploymentGetResponseStrategyPercentage:
		return true
	}
	return false
}

type ScriptDeploymentGetResponseVersion struct {
	Percentage float64                                `json:"percentage" api:"required"`
	VersionID  string                                 `json:"version_id" api:"required" format:"uuid"`
	JSON       scriptDeploymentGetResponseVersionJSON `json:"-"`
}

// scriptDeploymentGetResponseVersionJSON contains the JSON metadata for the struct
// [ScriptDeploymentGetResponseVersion]
type scriptDeploymentGetResponseVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseVersionJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseAnnotations struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// Operation that triggered the creation of the deployment.
	WorkersTriggeredBy string                                     `json:"workers/triggered_by"`
	JSON               scriptDeploymentGetResponseAnnotationsJSON `json:"-"`
}

// scriptDeploymentGetResponseAnnotationsJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseAnnotations]
type scriptDeploymentGetResponseAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseAnnotationsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewParams struct {
	// Identifier.
	AccountID param.Field[string]                             `path:"account_id" api:"required"`
	Strategy  param.Field[ScriptDeploymentNewParamsStrategy]  `json:"strategy" api:"required"`
	Versions  param.Field[[]ScriptDeploymentNewParamsVersion] `json:"versions" api:"required"`
	// If set to true, the deployment will be created even if normally blocked by
	// something such rolling back to an older version when a secret has changed.
	Force       param.Field[bool]                                 `query:"force"`
	Annotations param.Field[ScriptDeploymentNewParamsAnnotations] `json:"annotations"`
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
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	VersionID  param.Field[string]  `json:"version_id" api:"required" format:"uuid"`
}

func (r ScriptDeploymentNewParamsVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentNewParamsAnnotations struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	WorkersMessage param.Field[string] `json:"workers/message"`
}

func (r ScriptDeploymentNewParamsAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentNewResponseEnvelope struct {
	Errors   []ScriptDeploymentNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptDeploymentNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptDeploymentNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptDeploymentNewResponseEnvelopeSuccess `json:"success" api:"required"`
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

type ScriptDeploymentNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptDeploymentNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeploymentNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptDeploymentNewResponseEnvelopeErrors]
type scriptDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptDeploymentNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentNewResponseEnvelopeErrorsSource]
type scriptDeploymentNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ScriptDeploymentNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeploymentNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptDeploymentNewResponseEnvelopeMessages]
type scriptDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    scriptDeploymentNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentNewResponseEnvelopeMessagesSource]
type scriptDeploymentNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type ScriptDeploymentListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ScriptDeploymentListResponseEnvelope struct {
	Errors   []ScriptDeploymentListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptDeploymentListResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptDeploymentListResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptDeploymentListResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    scriptDeploymentListResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentListResponseEnvelope]
type scriptDeploymentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptDeploymentListResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeploymentListResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeploymentListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ScriptDeploymentListResponseEnvelopeErrors]
type scriptDeploymentListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptDeploymentListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeploymentListResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentListResponseEnvelopeErrorsSource]
type scriptDeploymentListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ScriptDeploymentListResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeploymentListResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeploymentListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptDeploymentListResponseEnvelopeMessages]
type scriptDeploymentListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    scriptDeploymentListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeploymentListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ScriptDeploymentListResponseEnvelopeMessagesSource]
type scriptDeploymentListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeploymentListResponseEnvelopeSuccess bool

const (
	ScriptDeploymentListResponseEnvelopeSuccessTrue ScriptDeploymentListResponseEnvelopeSuccess = true
)

func (r ScriptDeploymentListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptDeploymentDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ScriptDeploymentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ScriptDeploymentGetResponseEnvelope struct {
	Errors   []ScriptDeploymentGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptDeploymentGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptDeploymentGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptDeploymentGetResponseEnvelopeSuccess `json:"success" api:"required"`
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

type ScriptDeploymentGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptDeploymentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeploymentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseEnvelopeErrors]
type scriptDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptDeploymentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentGetResponseEnvelopeErrorsSource]
type scriptDeploymentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ScriptDeploymentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeploymentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptDeploymentGetResponseEnvelopeMessages]
type scriptDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    scriptDeploymentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentGetResponseEnvelopeMessagesSource]
type scriptDeploymentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
