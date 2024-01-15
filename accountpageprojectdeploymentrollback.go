// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPageProjectDeploymentRollbackService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPageProjectDeploymentRollbackService] method instead.
type AccountPageProjectDeploymentRollbackService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDeploymentRollbackService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountPageProjectDeploymentRollbackService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentRollbackService) {
	r = &AccountPageProjectDeploymentRollbackService{}
	r.Options = opts
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *AccountPageProjectDeploymentRollbackService) PagesDeploymentRollbackDeployment(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse struct {
	Errors   []AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseError   `json:"errors"`
	Messages []AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessage `json:"messages"`
	Result   AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseSuccess `json:"success"`
	JSON    accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseErrorJSON `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseError]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessageJSON `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessage]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                                                  `json:"url"`
	JSON accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultJSON `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResult]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultJSON struct {
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

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                                                   `json:"type"`
	JSON accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTrigger]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                                                           `json:"commit_message"`
	JSON          accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadata]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                                                       `json:"status"`
	JSON   accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStageJSON `json:"-"`
}

// accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStage]
type accountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseSuccess bool

const (
	AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseSuccessTrue AccountPageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseSuccess = true
)
