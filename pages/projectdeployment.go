// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ProjectDeploymentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDeploymentService] method instead.
type ProjectDeploymentService struct {
	Options []option.RequestOption
	History *ProjectDeploymentHistoryService
}

// NewProjectDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectDeploymentService(opts ...option.RequestOption) (r *ProjectDeploymentService) {
	r = &ProjectDeploymentService{}
	r.Options = opts
	r.History = NewProjectDeploymentHistoryService(opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *ProjectDeploymentService) New(ctx context.Context, projectName string, params ProjectDeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) List(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Deployment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
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

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) ListAutoPaging(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Deployment] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, projectName, params, opts...))
}

// Delete a deployment.
func (r *ProjectDeploymentService) Delete(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *ProjectDeploymentDeleteResponse, err error) {
	var env ProjectDeploymentDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch information about a deployment.
func (r *ProjectDeploymentService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry a previous deployment.
func (r *ProjectDeploymentService) Retry(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentRetryParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentRetryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *ProjectDeploymentService) Rollback(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentRollbackResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentDeleteResponse = interface{}

type ProjectDeploymentNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Headers configuration file for the deployment.
	Headers param.Field[io.Reader] `json:"_headers" format:"binary"`
	// Redirects configuration file for the deployment.
	Redirects param.Field[io.Reader] `json:"_redirects" format:"binary"`
	// Routes configuration file defining routing rules.
	RoutesJson param.Field[io.Reader] `json:"_routes.json" format:"binary"`
	// Worker bundle file in multipart/form-data format. Mutually exclusive with
	// `_worker.js`. Cannot specify both `_worker.js` and `_worker.bundle` in the same
	// request. Maximum size: 25 MiB.
	WorkerBundle param.Field[io.Reader] `json:"_worker.bundle" format:"binary"`
	// Worker JavaScript file. Mutually exclusive with `_worker.bundle`. Cannot specify
	// both `_worker.js` and `_worker.bundle` in the same request.
	WorkerJS param.Field[io.Reader] `json:"_worker.js" format:"binary"`
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
	// Boolean string indicating if the working directory has uncommitted changes.
	CommitDirty param.Field[ProjectDeploymentNewParamsCommitDirty] `json:"commit_dirty"`
	// Git commit SHA associated with this deployment.
	CommitHash param.Field[string] `json:"commit_hash"`
	// Git commit message associated with this deployment.
	CommitMessage param.Field[string] `json:"commit_message"`
	// Functions routing configuration file.
	FunctionsFilepathRoutingConfigJson param.Field[io.Reader] `json:"functions-filepath-routing-config.json" format:"binary"`
	// JSON string containing a manifest of files to deploy. Maps file paths to their
	// content hashes. Required for direct upload deployments. Maximum 20,000 entries.
	Manifest param.Field[string] `json:"manifest"`
	// The build output directory path.
	PagesBuildOutputDir param.Field[string] `json:"pages_build_output_dir"`
	// Hash of the Wrangler configuration file used for this deployment.
	WranglerConfigHash param.Field[string] `json:"wrangler_config_hash"`
}

func (r ProjectDeploymentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Boolean string indicating if the working directory has uncommitted changes.
type ProjectDeploymentNewParamsCommitDirty string

const (
	ProjectDeploymentNewParamsCommitDirtyTrue  ProjectDeploymentNewParamsCommitDirty = "true"
	ProjectDeploymentNewParamsCommitDirtyFalse ProjectDeploymentNewParamsCommitDirty = "false"
)

func (r ProjectDeploymentNewParamsCommitDirty) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewParamsCommitDirtyTrue, ProjectDeploymentNewParamsCommitDirtyFalse:
		return true
	}
	return false
}

type ProjectDeploymentNewResponseEnvelope struct {
	Errors   []ProjectDeploymentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   Deployment                                     `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentNewResponseEnvelope]
type projectDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ProjectDeploymentNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseEnvelopeErrors]
type projectDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    projectDeploymentNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ProjectDeploymentNewResponseEnvelopeErrorsSource]
type projectDeploymentNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectDeploymentNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseEnvelopeMessages]
type projectDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectDeploymentNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentNewResponseEnvelopeMessagesSource]
type projectDeploymentNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentNewResponseEnvelopeSuccess bool

const (
	ProjectDeploymentNewResponseEnvelopeSuccessTrue ProjectDeploymentNewResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// What type of deployments to fetch.
	Env param.Field[ProjectDeploymentListParamsEnv] `query:"env"`
	// Which page of deployments to fetch.
	Page param.Field[int64] `query:"page"`
	// How many deployments to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ProjectDeploymentListParams]'s query parameters as
// `url.Values`.
func (r ProjectDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// What type of deployments to fetch.
type ProjectDeploymentListParamsEnv string

const (
	ProjectDeploymentListParamsEnvProduction ProjectDeploymentListParamsEnv = "production"
	ProjectDeploymentListParamsEnvPreview    ProjectDeploymentListParamsEnv = "preview"
)

func (r ProjectDeploymentListParamsEnv) IsKnown() bool {
	switch r {
	case ProjectDeploymentListParamsEnvProduction, ProjectDeploymentListParamsEnvPreview:
		return true
	}
	return false
}

type ProjectDeploymentDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentDeleteResponseEnvelope struct {
	Errors   []ProjectDeploymentDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ProjectDeploymentDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentDeleteResponseEnvelope]
type projectDeploymentDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ProjectDeploymentDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentDeleteResponseEnvelopeErrors]
type projectDeploymentDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentDeleteResponseEnvelopeErrorsSource]
type projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ProjectDeploymentDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentDeleteResponseEnvelopeMessages]
type projectDeploymentDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentDeleteResponseEnvelopeMessagesSource]
type projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeploymentDeleteResponseEnvelopeSuccessTrue ProjectDeploymentDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentGetResponseEnvelope struct {
	Errors   []ProjectDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Deployment                                     `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentGetResponseEnvelope]
type projectDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ProjectDeploymentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseEnvelopeErrors]
type projectDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    projectDeploymentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ProjectDeploymentGetResponseEnvelopeErrorsSource]
type projectDeploymentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectDeploymentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseEnvelopeMessages]
type projectDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectDeploymentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentGetResponseEnvelopeMessagesSource]
type projectDeploymentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentGetResponseEnvelopeSuccessTrue ProjectDeploymentGetResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentRetryParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentRetryResponseEnvelope struct {
	Errors   []ProjectDeploymentRetryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRetryResponseEnvelopeMessages `json:"messages,required"`
	Result   Deployment                                       `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentRetryResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentRetryResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseEnvelope]
type projectDeploymentRetryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectDeploymentRetryResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentRetryResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentRetryResponseEnvelopeErrors]
type projectDeploymentRetryResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectDeploymentRetryResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentRetryResponseEnvelopeErrorsSource]
type projectDeploymentRetryResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ProjectDeploymentRetryResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentRetryResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentRetryResponseEnvelopeMessages]
type projectDeploymentRetryResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    projectDeploymentRetryResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentRetryResponseEnvelopeMessagesSource]
type projectDeploymentRetryResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentRetryResponseEnvelopeSuccess bool

const (
	ProjectDeploymentRetryResponseEnvelopeSuccessTrue ProjectDeploymentRetryResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentRetryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentRollbackParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentRollbackResponseEnvelope struct {
	Errors   []ProjectDeploymentRollbackResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRollbackResponseEnvelopeMessages `json:"messages,required"`
	Result   Deployment                                          `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentRollbackResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentRollbackResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentRollbackResponseEnvelope]
type projectDeploymentRollbackResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ProjectDeploymentRollbackResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentRollbackResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ProjectDeploymentRollbackResponseEnvelopeErrors]
type projectDeploymentRollbackResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentRollbackResponseEnvelopeErrorsSource]
type projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           ProjectDeploymentRollbackResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentRollbackResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentRollbackResponseEnvelopeMessages]
type projectDeploymentRollbackResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRollbackResponseEnvelopeMessagesSource]
type projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentRollbackResponseEnvelopeSuccess bool

const (
	ProjectDeploymentRollbackResponseEnvelopeSuccessTrue ProjectDeploymentRollbackResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentRollbackResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
