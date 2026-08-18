// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ProjectDeploymentTailService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDeploymentTailService] method instead.
type ProjectDeploymentTailService struct {
	Options []option.RequestOption
}

// NewProjectDeploymentTailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectDeploymentTailService(opts ...option.RequestOption) (r *ProjectDeploymentTailService) {
	r = &ProjectDeploymentTailService{}
	r.Options = opts
	return
}

// Start a tail that receives logs and exception data.
func (r *ProjectDeploymentTailService) New(ctx context.Context, projectName string, deploymentID string, params ProjectDeploymentTailNewParams, opts ...option.RequestOption) (res *ProjectDeploymentTailNewResponse, err error) {
	var env ProjectDeploymentTailNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/tails", params.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes a tail from a Pages deployment.
func (r *ProjectDeploymentTailService) Delete(ctx context.Context, projectName string, deploymentID string, tailID string, body ProjectDeploymentTailDeleteParams, opts ...option.RequestOption) (res *ProjectDeploymentTailDeleteResponse, err error) {
	var env ProjectDeploymentTailDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	if tailID == "" {
		err = errors.New("missing required tail_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/tails/%s", body.AccountID, projectName, deploymentID, tailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A tail session for streaming logs from a Pages deployment.
type ProjectDeploymentTailNewResponse struct {
	// Identifier of the tail session.
	ID string `json:"id" api:"required"`
	// Optional WebSocket URL to connect to for receiving tail events, when returned by
	// the tail service.
	URL  string                               `json:"url"`
	JSON projectDeploymentTailNewResponseJSON `json:"-"`
}

// projectDeploymentTailNewResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentTailNewResponse]
type projectDeploymentTailNewResponseJSON struct {
	ID          apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailDeleteResponse = interface{}

type ProjectDeploymentTailNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply to the tail session.
	Filters param.Field[[]map[string]interface{}] `json:"filters"`
}

func (r ProjectDeploymentTailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDeploymentTailNewResponseEnvelope struct {
	Errors   []ProjectDeploymentTailNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ProjectDeploymentTailNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// A tail session for streaming logs from a Pages deployment.
	Result ProjectDeploymentTailNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ProjectDeploymentTailNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    projectDeploymentTailNewResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentTailNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentTailNewResponseEnvelope]
type projectDeploymentTailNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailNewResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ProjectDeploymentTailNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentTailNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentTailNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ProjectDeploymentTailNewResponseEnvelopeErrors]
type projectDeploymentTailNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentTailNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    projectDeploymentTailNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentTailNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentTailNewResponseEnvelopeErrorsSource]
type projectDeploymentTailNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailNewResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ProjectDeploymentTailNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentTailNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentTailNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentTailNewResponseEnvelopeMessages]
type projectDeploymentTailNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentTailNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    projectDeploymentTailNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentTailNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentTailNewResponseEnvelopeMessagesSource]
type projectDeploymentTailNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentTailNewResponseEnvelopeSuccess bool

const (
	ProjectDeploymentTailNewResponseEnvelopeSuccessTrue ProjectDeploymentTailNewResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentTailNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentTailNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentTailDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ProjectDeploymentTailDeleteResponseEnvelope struct {
	Errors   []ProjectDeploymentTailDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ProjectDeploymentTailDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ProjectDeploymentTailDeleteResponse                   `json:"result" api:"required,nullable"`
	// Whether the API call was successful.
	Success ProjectDeploymentTailDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    projectDeploymentTailDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentTailDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ProjectDeploymentTailDeleteResponseEnvelope]
type projectDeploymentTailDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailDeleteResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           ProjectDeploymentTailDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentTailDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentTailDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ProjectDeploymentTailDeleteResponseEnvelopeErrors]
type projectDeploymentTailDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentTailDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    projectDeploymentTailDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentTailDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentTailDeleteResponseEnvelopeErrorsSource]
type projectDeploymentTailDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailDeleteResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           ProjectDeploymentTailDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentTailDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentTailDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ProjectDeploymentTailDeleteResponseEnvelopeMessages]
type projectDeploymentTailDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentTailDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentTailDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    projectDeploymentTailDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentTailDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentTailDeleteResponseEnvelopeMessagesSource]
type projectDeploymentTailDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentTailDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentTailDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentTailDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeploymentTailDeleteResponseEnvelopeSuccessTrue ProjectDeploymentTailDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentTailDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentTailDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
