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

// PageProjectDeploymentRollbackService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewPageProjectDeploymentRollbackService] method instead.
type PageProjectDeploymentRollbackService struct {
	Options []option.RequestOption
}

// NewPageProjectDeploymentRollbackService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPageProjectDeploymentRollbackService(opts ...option.RequestOption) (r *PageProjectDeploymentRollbackService) {
	r = &PageProjectDeploymentRollbackService{}
	r.Options = opts
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *PageProjectDeploymentRollbackService) PagesDeploymentRollbackDeployment(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                                     `json:"url"`
	JSON pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseJSON `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseJSON struct {
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

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                                      `json:"type"`
	JSON pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerJSON `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTrigger]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                                              `json:"commit_message"`
	JSON          pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadata]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                                          `json:"status"`
	JSON   pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStageJSON `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStageJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStage]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelope struct {
	Errors   []PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelope]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrors struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrors]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessages struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessages]
type pageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeSuccessTrue PageProjectDeploymentRollbackPagesDeploymentRollbackDeploymentResponseEnvelopeSuccess = true
)
