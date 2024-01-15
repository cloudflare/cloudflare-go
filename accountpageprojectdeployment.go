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

// Fetch information about a deployment.
func (r *AccountPageProjectDeploymentService) Get(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a deployment.
func (r *AccountPageProjectDeploymentService) Delete(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *AccountPageProjectDeploymentService) PagesDeploymentNewDeployment(ctx context.Context, accountIdentifier string, projectName string, body AccountPageProjectDeploymentPagesDeploymentNewDeploymentParams, opts ...option.RequestOption) (res *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a list of project deployments.
func (r *AccountPageProjectDeploymentService) PagesDeploymentGetDeployments(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPageProjectDeploymentGetResponse struct {
	Errors   []AccountPageProjectDeploymentGetResponseError   `json:"errors"`
	Messages []AccountPageProjectDeploymentGetResponseMessage `json:"messages"`
	Result   AccountPageProjectDeploymentGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentGetResponseSuccess `json:"success"`
	JSON    accountPageProjectDeploymentGetResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentGetResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDeploymentGetResponse]
type accountPageProjectDeploymentGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentGetResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountPageProjectDeploymentGetResponseErrorJSON `json:"-"`
}

// accountPageProjectDeploymentGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountPageProjectDeploymentGetResponseError]
type accountPageProjectDeploymentGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentGetResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountPageProjectDeploymentGetResponseMessageJSON `json:"-"`
}

// accountPageProjectDeploymentGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountPageProjectDeploymentGetResponseMessage]
type accountPageProjectDeploymentGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentGetResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectDeploymentGetResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectDeploymentGetResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                            `json:"url"`
	JSON accountPageProjectDeploymentGetResponseResultJSON `json:"-"`
}

// accountPageProjectDeploymentGetResponseResultJSON contains the JSON metadata for
// the struct [AccountPageProjectDeploymentGetResponseResult]
type accountPageProjectDeploymentGetResponseResultJSON struct {
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

func (r *AccountPageProjectDeploymentGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectDeploymentGetResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                             `json:"type"`
	JSON accountPageProjectDeploymentGetResponseResultDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectDeploymentGetResponseResultDeploymentTriggerJSON contains the
// JSON metadata for the struct
// [AccountPageProjectDeploymentGetResponseResultDeploymentTrigger]
type accountPageProjectDeploymentGetResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentGetResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                     `json:"commit_message"`
	JSON          accountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadata]
type accountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentGetResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectDeploymentGetResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                 `json:"status"`
	JSON   accountPageProjectDeploymentGetResponseResultStageJSON `json:"-"`
}

// accountPageProjectDeploymentGetResponseResultStageJSON contains the JSON
// metadata for the struct [AccountPageProjectDeploymentGetResponseResultStage]
type accountPageProjectDeploymentGetResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentGetResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDeploymentGetResponseSuccess bool

const (
	AccountPageProjectDeploymentGetResponseSuccessTrue AccountPageProjectDeploymentGetResponseSuccess = true
)

type AccountPageProjectDeploymentDeleteResponse = interface{}

type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponse struct {
	Errors   []AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseError   `json:"errors"`
	Messages []AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessage `json:"messages"`
	Result   AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseSuccess `json:"success"`
	JSON    accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseJSON contains
// the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponse]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseErrorJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseError]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessageJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessage]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                                     `json:"url"`
	JSON accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResult]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultJSON struct {
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

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                                      `json:"type"`
	JSON accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTrigger]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                                              `json:"commit_message"`
	JSON          accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadata]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                                          `json:"status"`
	JSON   accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStageJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStage]
type accountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseSuccess bool

const (
	AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseSuccessTrue AccountPageProjectDeploymentPagesDeploymentNewDeploymentResponseSuccess = true
)

type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponse struct {
	Errors     []AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseError    `json:"errors"`
	Messages   []AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessage  `json:"messages"`
	Result     []AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResult   `json:"result"`
	ResultInfo AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseSuccess `json:"success"`
	JSON    accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseJSON contains
// the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponse]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseErrorJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseError]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessageJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessage]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                                      `json:"url"`
	JSON accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResult]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultJSON struct {
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

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                                       `json:"type"`
	JSON accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTrigger]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                                               `json:"commit_message"`
	JSON          accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadata]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                                           `json:"status"`
	JSON   accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStageJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStage]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfo struct {
	Count      interface{}                                                                     `json:"count"`
	Page       interface{}                                                                     `json:"page"`
	PerPage    interface{}                                                                     `json:"per_page"`
	TotalCount interface{}                                                                     `json:"total_count"`
	JSON       accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfoJSON `json:"-"`
}

// accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfo]
type accountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseSuccess bool

const (
	AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseSuccessTrue AccountPageProjectDeploymentPagesDeploymentGetDeploymentsResponseSuccess = true
)

type AccountPageProjectDeploymentPagesDeploymentNewDeploymentParams struct {
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r AccountPageProjectDeploymentPagesDeploymentNewDeploymentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
