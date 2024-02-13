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

// PageProjectDeploymentRetryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewPageProjectDeploymentRetryService] method instead.
type PageProjectDeploymentRetryService struct {
	Options []option.RequestOption
}

// NewPageProjectDeploymentRetryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPageProjectDeploymentRetryService(opts ...option.RequestOption) (r *PageProjectDeploymentRetryService) {
	r = &PageProjectDeploymentRetryService{}
	r.Options = opts
	return
}

// Retry a previous deployment.
func (r *PageProjectDeploymentRetryService) PagesDeploymentRetryDeployment(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                               `json:"url"`
	JSON pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseJSON `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseJSON contains
// the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseJSON struct {
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

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                                `json:"type"`
	JSON pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerJSON `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTrigger]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                                        `json:"commit_message"`
	JSON          pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadata]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                                    `json:"status"`
	JSON   pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStageJSON `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStageJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStage]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelope struct {
	Errors   []PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelope]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrors]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessages]
type pageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeSuccessTrue PageProjectDeploymentRetryPagesDeploymentRetryDeploymentResponseEnvelopeSuccess = true
)
