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

// AccountPageProjectDeploymentRetryService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPageProjectDeploymentRetryService] method instead.
type AccountPageProjectDeploymentRetryService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDeploymentRetryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPageProjectDeploymentRetryService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentRetryService) {
	r = &AccountPageProjectDeploymentRetryService{}
	r.Options = opts
	return
}

// Retry a previous deployment.
func (r *AccountPageProjectDeploymentRetryService) PagesDeploymentRetryDeployment(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse struct {
	Errors   []AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseError   `json:"errors"`
	Messages []AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessage `json:"messages"`
	Result   AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseSuccess `json:"success"`
	JSON    accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseError struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseErrorJSON `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseError]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessage struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessageJSON `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessage]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                                            `json:"url"`
	JSON accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultJSON `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResult]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultJSON struct {
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

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                                             `json:"type"`
	JSON accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTrigger]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                                                     `json:"commit_message"`
	JSON          accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadata]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                                                 `json:"status"`
	JSON   accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStageJSON `json:"-"`
}

// accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStage]
type accountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseSuccess bool

const (
	AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseSuccessTrue AccountPageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseSuccess = true
)
