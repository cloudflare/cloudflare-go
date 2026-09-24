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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
// are deployed to traffic. A deployment can consist of multiple versions of a
// Worker.
func (r *ScriptDeploymentService) New(ctx context.Context, scriptName string, params ScriptDeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
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

// List Worker deployments. The first deployment in the list is the latest
// deployment actively serving traffic.
func (r *ScriptDeploymentService) List(ctx context.Context, scriptName string, params ScriptDeploymentListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ScriptDeploymentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", params.AccountID, scriptName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Worker deployments. The first deployment in the list is the latest
// deployment actively serving traffic.
func (r *ScriptDeploymentService) ListAutoPaging(ctx context.Context, scriptName string, params ScriptDeploymentListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[ScriptDeploymentListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, scriptName, params, opts...))
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
func (r *ScriptDeploymentService) Get(ctx context.Context, scriptName string, deploymentID string, query ScriptDeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
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

type Deployment struct {
	ID        string             `json:"id" api:"required" format:"uuid"`
	CreatedOn time.Time          `json:"created_on" api:"required" format:"date-time"`
	Source    string             `json:"source" api:"required"`
	Strategy  DeploymentStrategy `json:"strategy" api:"required"`
	// Worker versions included in this deployment. Each object must contain a
	// `version_id` UUID and a `percentage`; percentages across all objects must
	// total 100. In the `cf` CLI, pass the entire array as one JSON value to
	// `--versions`, either inline, for example
	// `--versions '[{"version_id":"023e105f-2a42-4f8b-a1c1-73f6a2a30c0f","percentage":100}]'`,
	// or from a JSON file with `--versions @versions.json`.
	Versions    []DeploymentVersion   `json:"versions" api:"required"`
	Annotations DeploymentAnnotations `json:"annotations"`
	AuthorEmail string                `json:"author_email" format:"email"`
	JSON        deploymentJSON        `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
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

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

type DeploymentStrategy string

const (
	DeploymentStrategyPercentage DeploymentStrategy = "percentage"
)

func (r DeploymentStrategy) IsKnown() bool {
	switch r {
	case DeploymentStrategyPercentage:
		return true
	}
	return false
}

type DeploymentVersion struct {
	// Percentage of traffic served by this version.
	Percentage float64 `json:"percentage" api:"required"`
	// Identifier of the Worker Version.
	VersionID string                `json:"version_id" api:"required" format:"uuid"`
	JSON      deploymentVersionJSON `json:"-"`
}

// deploymentVersionJSON contains the JSON metadata for the struct
// [DeploymentVersion]
type deploymentVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentVersionJSON) RawJSON() string {
	return r.raw
}

type DeploymentAnnotations struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// Operation that triggered the creation of the deployment.
	WorkersTriggeredBy string                    `json:"workers/triggered_by"`
	JSON               deploymentAnnotationsJSON `json:"-"`
}

// deploymentAnnotationsJSON contains the JSON metadata for the struct
// [DeploymentAnnotations]
type deploymentAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DeploymentAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentAnnotationsJSON) RawJSON() string {
	return r.raw
}

type DeploymentParam struct {
	Strategy param.Field[DeploymentStrategy] `json:"strategy" api:"required"`
	// Worker versions included in this deployment. Each object must contain a
	// `version_id` UUID and a `percentage`; percentages across all objects must
	// total 100. In the `cf` CLI, pass the entire array as one JSON value to
	// `--versions`, either inline, for example
	// `--versions '[{"version_id":"023e105f-2a42-4f8b-a1c1-73f6a2a30c0f","percentage":100}]'`,
	// or from a JSON file with `--versions @versions.json`.
	Versions    param.Field[[]DeploymentVersionParam]   `json:"versions" api:"required"`
	Annotations param.Field[DeploymentAnnotationsParam] `json:"annotations"`
}

func (r DeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeploymentVersionParam struct {
	// Percentage of traffic served by this version.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Identifier of the Worker Version.
	VersionID param.Field[string] `json:"version_id" api:"required" format:"uuid"`
}

func (r DeploymentVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeploymentAnnotationsParam struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	WorkersMessage param.Field[string] `json:"workers/message"`
}

func (r DeploymentAnnotationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentListResponse struct {
	Deployments []Deployment                     `json:"deployments" api:"required"`
	JSON        scriptDeploymentListResponseJSON `json:"-"`
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

type ScriptDeploymentNewParams struct {
	// Identifier.
	AccountID  param.Field[string] `path:"account_id" api:"required"`
	Deployment DeploymentParam     `json:"deployment" api:"required"`
	// If set to true, the deployment will be created even if normally blocked by
	// something such rolling back to an older version when a secret has changed.
	Force param.Field[bool] `query:"force"`
}

func (r ScriptDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Deployment)
}

// URLQuery serializes [ScriptDeploymentNewParams]'s query parameters as
// `url.Values`.
func (r ScriptDeploymentNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeploymentNewResponseEnvelope struct {
	Errors   []ScriptDeploymentNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptDeploymentNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Deployment                                    `json:"result" api:"required"`
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
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Start of the deployment creation time range, inclusive.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// End of the deployment creation time range, inclusive.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [ScriptDeploymentListParams]'s query parameters as
// `url.Values`.
func (r ScriptDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
	Result   Deployment                                    `json:"result" api:"required"`
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
