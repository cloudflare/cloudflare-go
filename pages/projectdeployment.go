// File generated from our OpenAPI spec by Stainless.

package pages

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ProjectDeploymentService) New(ctx context.Context, projectName string, params ProjectDeploymentNewParams, opts ...option.RequestOption) (res *ProjectDeploymentNewResponse, err error) {
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
func (r *ProjectDeploymentService) List(ctx context.Context, projectName string, query ProjectDeploymentListParams, opts ...option.RequestOption) (res *[]ProjectDeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a deployment.
func (r *ProjectDeploymentService) Delete(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *ProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetch information about a deployment.
func (r *ProjectDeploymentService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentGetParams, opts ...option.RequestOption) (res *ProjectDeploymentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry a previous deployment.
func (r *ProjectDeploymentService) Retry(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentRetryParams, opts ...option.RequestOption) (res *ProjectDeploymentRetryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentRetryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *ProjectDeploymentService) Rollback(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *ProjectDeploymentRollbackResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentRollbackResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentNewResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentNewResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []ProjectDeploymentNewResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                           `json:"url"`
	JSON projectDeploymentNewResponseJSON `json:"-"`
}

// projectDeploymentNewResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentNewResponse]
type projectDeploymentNewResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentNewResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentNewResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                            `json:"type"`
	JSON projectDeploymentNewResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentNewResponseDeploymentTriggerJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseDeploymentTrigger]
type projectDeploymentNewResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentNewResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                    `json:"commit_message"`
	JSON          projectDeploymentNewResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentNewResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [ProjectDeploymentNewResponseDeploymentTriggerMetadata]
type projectDeploymentNewResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectDeploymentNewResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                `json:"status"`
	JSON   projectDeploymentNewResponseStageJSON `json:"-"`
}

// projectDeploymentNewResponseStageJSON contains the JSON metadata for the struct
// [ProjectDeploymentNewResponseStage]
type projectDeploymentNewResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseStageJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentListResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentListResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []ProjectDeploymentListResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                            `json:"url"`
	JSON projectDeploymentListResponseJSON `json:"-"`
}

// projectDeploymentListResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentListResponse]
type projectDeploymentListResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentListResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentListResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                             `json:"type"`
	JSON projectDeploymentListResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentListResponseDeploymentTriggerJSON contains the JSON metadata
// for the struct [ProjectDeploymentListResponseDeploymentTrigger]
type projectDeploymentListResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentListResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                     `json:"commit_message"`
	JSON          projectDeploymentListResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentListResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [ProjectDeploymentListResponseDeploymentTriggerMetadata]
type projectDeploymentListResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectDeploymentListResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                 `json:"status"`
	JSON   projectDeploymentListResponseStageJSON `json:"-"`
}

// projectDeploymentListResponseStageJSON contains the JSON metadata for the struct
// [ProjectDeploymentListResponseStage]
type projectDeploymentListResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseStageJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponse = interface{}

type ProjectDeploymentGetResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentGetResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []ProjectDeploymentGetResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                           `json:"url"`
	JSON projectDeploymentGetResponseJSON `json:"-"`
}

// projectDeploymentGetResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentGetResponse]
type projectDeploymentGetResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentGetResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentGetResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                            `json:"type"`
	JSON projectDeploymentGetResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentGetResponseDeploymentTriggerJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseDeploymentTrigger]
type projectDeploymentGetResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentGetResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                    `json:"commit_message"`
	JSON          projectDeploymentGetResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentGetResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [ProjectDeploymentGetResponseDeploymentTriggerMetadata]
type projectDeploymentGetResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectDeploymentGetResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                `json:"status"`
	JSON   projectDeploymentGetResponseStageJSON `json:"-"`
}

// projectDeploymentGetResponseStageJSON contains the JSON metadata for the struct
// [ProjectDeploymentGetResponseStage]
type projectDeploymentGetResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseStageJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentRetryResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []ProjectDeploymentRetryResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                             `json:"url"`
	JSON projectDeploymentRetryResponseJSON `json:"-"`
}

// projectDeploymentRetryResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentRetryResponse]
type projectDeploymentRetryResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentRetryResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentRetryResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                              `json:"type"`
	JSON projectDeploymentRetryResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentRetryResponseDeploymentTriggerJSON contains the JSON metadata
// for the struct [ProjectDeploymentRetryResponseDeploymentTrigger]
type projectDeploymentRetryResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentRetryResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                      `json:"commit_message"`
	JSON          projectDeploymentRetryResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentRetryResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRetryResponseDeploymentTriggerMetadata]
type projectDeploymentRetryResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectDeploymentRetryResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                  `json:"status"`
	JSON   projectDeploymentRetryResponseStageJSON `json:"-"`
}

// projectDeploymentRetryResponseStageJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseStage]
type projectDeploymentRetryResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseStageJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentRollbackResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []ProjectDeploymentRollbackResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                `json:"url"`
	JSON projectDeploymentRollbackResponseJSON `json:"-"`
}

// projectDeploymentRollbackResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentRollbackResponse]
type projectDeploymentRollbackResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentRollbackResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentRollbackResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                 `json:"type"`
	JSON projectDeploymentRollbackResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentRollbackResponseDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectDeploymentRollbackResponseDeploymentTrigger]
type projectDeploymentRollbackResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentRollbackResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                         `json:"commit_message"`
	JSON          projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRollbackResponseDeploymentTriggerMetadata]
type projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectDeploymentRollbackResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                     `json:"status"`
	JSON   projectDeploymentRollbackResponseStageJSON `json:"-"`
}

// projectDeploymentRollbackResponseStageJSON contains the JSON metadata for the
// struct [ProjectDeploymentRollbackResponseStage]
type projectDeploymentRollbackResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseStageJSON) RawJSON() string {
	return r.raw
}

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
	Result   ProjectDeploymentNewResponse                   `json:"result,required"`
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

type ProjectDeploymentListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentListResponseEnvelope struct {
	Errors   []ProjectDeploymentListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ProjectDeploymentListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    ProjectDeploymentListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ProjectDeploymentListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       projectDeploymentListResponseEnvelopeJSON       `json:"-"`
}

// projectDeploymentListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentListResponseEnvelope]
type projectDeploymentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    projectDeploymentListResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDeploymentListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentListResponseEnvelopeErrors]
type projectDeploymentListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    projectDeploymentListResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDeploymentListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentListResponseEnvelopeMessages]
type projectDeploymentListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentListResponseEnvelopeSuccess bool

const (
	ProjectDeploymentListResponseEnvelopeSuccessTrue ProjectDeploymentListResponseEnvelopeSuccess = true
)

type ProjectDeploymentListResponseEnvelopeResultInfo struct {
	Count      interface{}                                         `json:"count"`
	Page       interface{}                                         `json:"page"`
	PerPage    interface{}                                         `json:"per_page"`
	TotalCount interface{}                                         `json:"total_count"`
	JSON       projectDeploymentListResponseEnvelopeResultInfoJSON `json:"-"`
}

// projectDeploymentListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ProjectDeploymentListResponseEnvelopeResultInfo]
type projectDeploymentListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentGetResponseEnvelope struct {
	Errors   []ProjectDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentGetResponse                   `json:"result,required"`
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

type ProjectDeploymentRetryParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentRetryResponseEnvelope struct {
	Errors   []ProjectDeploymentRetryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRetryResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentRetryResponse                   `json:"result,required"`
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

type ProjectDeploymentRollbackParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentRollbackResponseEnvelope struct {
	Errors   []ProjectDeploymentRollbackResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRollbackResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentRollbackResponse                   `json:"result,required"`
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
