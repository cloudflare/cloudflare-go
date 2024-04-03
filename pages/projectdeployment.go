// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ProjectDeploymentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewProjectDeploymentService] method
// instead.
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
func (r *ProjectDeploymentService) New(ctx context.Context, projectName string, params ProjectDeploymentNewParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) List(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PagesDeployments], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *ProjectDeploymentService) ListAutoPaging(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PagesDeployments] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, projectName, params, opts...))
}

// Delete a deployment.
func (r *ProjectDeploymentService) Delete(ctx context.Context, projectName string, deploymentID string, params ProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *ProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", params.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch information about a deployment.
func (r *ProjectDeploymentService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentGetParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry a previous deployment.
func (r *ProjectDeploymentService) Retry(ctx context.Context, projectName string, deploymentID string, params ProjectDeploymentRetryParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentRetryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", params.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *ProjectDeploymentService) Rollback(ctx context.Context, projectName string, deploymentID string, params ProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentRollbackResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", params.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentDeleteResponse = interface{}

type ProjectDeploymentNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r ProjectDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDeploymentNewResponseEnvelope struct {
	Errors   []ProjectDeploymentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                               `json:"result,required"`
	// Whether the API call was successful
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
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    projectDeploymentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseEnvelopeErrors]
type projectDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    projectDeploymentNewResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseEnvelopeMessages]
type projectDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// What type of deployments to fetch.
	Env param.Field[ProjectDeploymentListParamsEnv] `query:"env"`
}

// URLQuery serializes [ProjectDeploymentListParams]'s query parameters as
// `url.Values`.
func (r ProjectDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
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
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDeploymentDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDeploymentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentGetResponseEnvelope struct {
	Errors   []ProjectDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                               `json:"result,required"`
	// Whether the API call was successful
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
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    projectDeploymentGetResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseEnvelopeErrors]
type projectDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    projectDeploymentGetResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseEnvelopeMessages]
type projectDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDeploymentRetryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDeploymentRetryResponseEnvelope struct {
	Errors   []ProjectDeploymentRetryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRetryResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                                 `json:"result,required"`
	// Whether the API call was successful
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
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    projectDeploymentRetryResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentRetryResponseEnvelopeErrors]
type projectDeploymentRetryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    projectDeploymentRetryResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentRetryResponseEnvelopeMessages]
type projectDeploymentRetryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDeploymentRollbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDeploymentRollbackResponseEnvelope struct {
	Errors   []ProjectDeploymentRollbackResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRollbackResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                                    `json:"result,required"`
	// Whether the API call was successful
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
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    projectDeploymentRollbackResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ProjectDeploymentRollbackResponseEnvelopeErrors]
type projectDeploymentRollbackResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    projectDeploymentRollbackResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentRollbackResponseEnvelopeMessages]
type projectDeploymentRollbackResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
