// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPageProjectDeploymentService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPageProjectDeploymentService] method instead.
type AccountPageProjectDeploymentService struct {
	Options   []option.RequestOption
	Histories *AccountPageProjectDeploymentHistoryService
	Retries   *AccountPageProjectDeploymentRetryService
	Rollbacks *AccountPageProjectDeploymentRollbackService
}

// NewAccountPageProjectDeploymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPageProjectDeploymentService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentService) {
	r = &AccountPageProjectDeploymentService{}
	r.Options = opts
	r.Histories = NewAccountPageProjectDeploymentHistoryService(opts...)
	r.Retries = NewAccountPageProjectDeploymentRetryService(opts...)
	r.Rollbacks = NewAccountPageProjectDeploymentRollbackService(opts...)
	return
}

// Fetch a deployment.
func (r *AccountPageProjectDeploymentService) Get(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *DeploymentResponseDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Destroy a deployment.
func (r *AccountPageProjectDeploymentService) Delete(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Start a new deployment from production. The repo/account must have been
// authorized through the Pages UI dash before.
func (r *AccountPageProjectDeploymentService) PagesDeploymentNewDeployment(ctx context.Context, accountIdentifier string, projectName string, body AccountPageProjectDeploymentPagesDeploymentNewDeploymentParams, opts ...option.RequestOption) (res *DeploymentNewDeployment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a list of deployments from a project.
func (r *AccountPageProjectDeploymentService) PagesDeploymentGetDeployments(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *DeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DeploymentListResponse struct {
	Errors     []DeploymentListResponseError    `json:"errors"`
	Messages   []DeploymentListResponseMessage  `json:"messages"`
	Result     []DeploymentListResponseResult   `json:"result"`
	ResultInfo DeploymentListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DeploymentListResponseSuccess `json:"success"`
	JSON    deploymentListResponseJSON    `json:"-"`
}

// deploymentListResponseJSON contains the JSON metadata for the struct
// [DeploymentListResponse]
type deploymentListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentListResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    deploymentListResponseErrorJSON `json:"-"`
}

// deploymentListResponseErrorJSON contains the JSON metadata for the struct
// [DeploymentListResponseError]
type deploymentListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentListResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    deploymentListResponseMessageJSON `json:"-"`
}

// deploymentListResponseMessageJSON contains the JSON metadata for the struct
// [DeploymentListResponseMessage]
type deploymentListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentListResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentListResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []DeploymentListResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                           `json:"url"`
	JSON deploymentListResponseResultJSON `json:"-"`
}

// deploymentListResponseResultJSON contains the JSON metadata for the struct
// [DeploymentListResponseResult]
type deploymentListResponseResultJSON struct {
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

func (r *DeploymentListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type DeploymentListResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentListResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                            `json:"type"`
	JSON deploymentListResponseResultDeploymentTriggerJSON `json:"-"`
}

// deploymentListResponseResultDeploymentTriggerJSON contains the JSON metadata for
// the struct [DeploymentListResponseResultDeploymentTrigger]
type deploymentListResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentListResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type DeploymentListResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                    `json:"commit_message"`
	JSON          deploymentListResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentListResponseResultDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [DeploymentListResponseResultDeploymentTriggerMetadata]
type deploymentListResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentListResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type DeploymentListResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                `json:"status"`
	JSON   deploymentListResponseResultStageJSON `json:"-"`
}

// deploymentListResponseResultStageJSON contains the JSON metadata for the struct
// [DeploymentListResponseResultStage]
type deploymentListResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentListResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentListResponseResultInfo struct {
	Count      interface{}                          `json:"count"`
	Page       interface{}                          `json:"page"`
	PerPage    interface{}                          `json:"per_page"`
	TotalCount interface{}                          `json:"total_count"`
	JSON       deploymentListResponseResultInfoJSON `json:"-"`
}

// deploymentListResponseResultInfoJSON contains the JSON metadata for the struct
// [DeploymentListResponseResultInfo]
type deploymentListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeploymentListResponseSuccess bool

const (
	DeploymentListResponseSuccessTrue DeploymentListResponseSuccess = true
)

type DeploymentNewDeployment struct {
	Errors   []DeploymentNewDeploymentError   `json:"errors"`
	Messages []DeploymentNewDeploymentMessage `json:"messages"`
	Result   DeploymentNewDeploymentResult    `json:"result"`
	// Whether the API call was successful
	Success DeploymentNewDeploymentSuccess `json:"success"`
	JSON    deploymentNewDeploymentJSON    `json:"-"`
}

// deploymentNewDeploymentJSON contains the JSON metadata for the struct
// [DeploymentNewDeployment]
type deploymentNewDeploymentJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentNewDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentNewDeploymentError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    deploymentNewDeploymentErrorJSON `json:"-"`
}

// deploymentNewDeploymentErrorJSON contains the JSON metadata for the struct
// [DeploymentNewDeploymentError]
type deploymentNewDeploymentErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentNewDeploymentError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentNewDeploymentMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    deploymentNewDeploymentMessageJSON `json:"-"`
}

// deploymentNewDeploymentMessageJSON contains the JSON metadata for the struct
// [DeploymentNewDeploymentMessage]
type deploymentNewDeploymentMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentNewDeploymentMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentNewDeploymentResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentNewDeploymentResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []DeploymentNewDeploymentResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                            `json:"url"`
	JSON deploymentNewDeploymentResultJSON `json:"-"`
}

// deploymentNewDeploymentResultJSON contains the JSON metadata for the struct
// [DeploymentNewDeploymentResult]
type deploymentNewDeploymentResultJSON struct {
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

func (r *DeploymentNewDeploymentResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type DeploymentNewDeploymentResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentNewDeploymentResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                             `json:"type"`
	JSON deploymentNewDeploymentResultDeploymentTriggerJSON `json:"-"`
}

// deploymentNewDeploymentResultDeploymentTriggerJSON contains the JSON metadata
// for the struct [DeploymentNewDeploymentResultDeploymentTrigger]
type deploymentNewDeploymentResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentNewDeploymentResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type DeploymentNewDeploymentResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                     `json:"commit_message"`
	JSON          deploymentNewDeploymentResultDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentNewDeploymentResultDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [DeploymentNewDeploymentResultDeploymentTriggerMetadata]
type deploymentNewDeploymentResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentNewDeploymentResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type DeploymentNewDeploymentResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                 `json:"status"`
	JSON   deploymentNewDeploymentResultStageJSON `json:"-"`
}

// deploymentNewDeploymentResultStageJSON contains the JSON metadata for the struct
// [DeploymentNewDeploymentResultStage]
type deploymentNewDeploymentResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentNewDeploymentResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeploymentNewDeploymentSuccess bool

const (
	DeploymentNewDeploymentSuccessTrue DeploymentNewDeploymentSuccess = true
)

type DeploymentResponseDetail struct {
	Errors   []DeploymentResponseDetailError   `json:"errors"`
	Messages []DeploymentResponseDetailMessage `json:"messages"`
	Result   DeploymentResponseDetailResult    `json:"result"`
	// Whether the API call was successful
	Success DeploymentResponseDetailSuccess `json:"success"`
	JSON    deploymentResponseDetailJSON    `json:"-"`
}

// deploymentResponseDetailJSON contains the JSON metadata for the struct
// [DeploymentResponseDetail]
type deploymentResponseDetailJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentResponseDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentResponseDetailError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    deploymentResponseDetailErrorJSON `json:"-"`
}

// deploymentResponseDetailErrorJSON contains the JSON metadata for the struct
// [DeploymentResponseDetailError]
type deploymentResponseDetailErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentResponseDetailError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentResponseDetailMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    deploymentResponseDetailMessageJSON `json:"-"`
}

// deploymentResponseDetailMessageJSON contains the JSON metadata for the struct
// [DeploymentResponseDetailMessage]
type deploymentResponseDetailMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentResponseDetailMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentResponseDetailResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentResponseDetailResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []DeploymentResponseDetailResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                             `json:"url"`
	JSON deploymentResponseDetailResultJSON `json:"-"`
}

// deploymentResponseDetailResultJSON contains the JSON metadata for the struct
// [DeploymentResponseDetailResult]
type deploymentResponseDetailResultJSON struct {
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

func (r *DeploymentResponseDetailResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type DeploymentResponseDetailResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentResponseDetailResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                              `json:"type"`
	JSON deploymentResponseDetailResultDeploymentTriggerJSON `json:"-"`
}

// deploymentResponseDetailResultDeploymentTriggerJSON contains the JSON metadata
// for the struct [DeploymentResponseDetailResultDeploymentTrigger]
type deploymentResponseDetailResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentResponseDetailResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type DeploymentResponseDetailResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                      `json:"commit_message"`
	JSON          deploymentResponseDetailResultDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentResponseDetailResultDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [DeploymentResponseDetailResultDeploymentTriggerMetadata]
type deploymentResponseDetailResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentResponseDetailResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type DeploymentResponseDetailResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                  `json:"status"`
	JSON   deploymentResponseDetailResultStageJSON `json:"-"`
}

// deploymentResponseDetailResultStageJSON contains the JSON metadata for the
// struct [DeploymentResponseDetailResultStage]
type deploymentResponseDetailResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentResponseDetailResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeploymentResponseDetailSuccess bool

const (
	DeploymentResponseDetailSuccessTrue DeploymentResponseDetailSuccess = true
)

type AccountPageProjectDeploymentDeleteResponse = interface{}

type AccountPageProjectDeploymentPagesDeploymentNewDeploymentParams struct {
	// The branch to build the new deployment from. The HEAD of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r AccountPageProjectDeploymentPagesDeploymentNewDeploymentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
