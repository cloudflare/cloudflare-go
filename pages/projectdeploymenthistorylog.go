// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ProjectDeploymentHistoryLogService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDeploymentHistoryLogService] method instead.
type ProjectDeploymentHistoryLogService struct {
	Options []option.RequestOption
}

// NewProjectDeploymentHistoryLogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectDeploymentHistoryLogService(opts ...option.RequestOption) (r *ProjectDeploymentHistoryLogService) {
	r = &ProjectDeploymentHistoryLogService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *ProjectDeploymentHistoryLogService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentHistoryLogGetParams, opts ...option.RequestOption) (res *ProjectDeploymentHistoryLogGetResponse, err error) {
	var env ProjectDeploymentHistoryLogGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentHistoryLogGetResponse struct {
	Data                  []ProjectDeploymentHistoryLogGetResponseData `json:"data,required"`
	IncludesContainerLogs bool                                         `json:"includes_container_logs,required"`
	Total                 int64                                        `json:"total,required"`
	JSON                  projectDeploymentHistoryLogGetResponseJSON   `json:"-"`
}

// projectDeploymentHistoryLogGetResponseJSON contains the JSON metadata for the
// struct [ProjectDeploymentHistoryLogGetResponse]
type projectDeploymentHistoryLogGetResponseJSON struct {
	Data                  apijson.Field
	IncludesContainerLogs apijson.Field
	Total                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseData struct {
	Line string                                         `json:"line,required"`
	Ts   string                                         `json:"ts,required"`
	JSON projectDeploymentHistoryLogGetResponseDataJSON `json:"-"`
}

// projectDeploymentHistoryLogGetResponseDataJSON contains the JSON metadata for
// the struct [ProjectDeploymentHistoryLogGetResponseData]
type projectDeploymentHistoryLogGetResponseDataJSON struct {
	Line        apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentHistoryLogGetResponseEnvelope struct {
	Errors   []ProjectDeploymentHistoryLogGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentHistoryLogGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentHistoryLogGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ProjectDeploymentHistoryLogGetResponseEnvelope]
type projectDeploymentHistoryLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           ProjectDeploymentHistoryLogGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ProjectDeploymentHistoryLogGetResponseEnvelopeErrors]
type projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentHistoryLogGetResponseEnvelopeErrorsSource]
type projectDeploymentHistoryLogGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           ProjectDeploymentHistoryLogGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ProjectDeploymentHistoryLogGetResponseEnvelopeMessages]
type projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentHistoryLogGetResponseEnvelopeMessagesSource]
type projectDeploymentHistoryLogGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
