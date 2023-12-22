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

// AccountPageProjectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountPageProjectService] method
// instead.
type AccountPageProjectService struct {
	Options     []option.RequestOption
	Deployments *AccountPageProjectDeploymentService
	Domains     *AccountPageProjectDomainService
}

// NewAccountPageProjectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPageProjectService(opts ...option.RequestOption) (r *AccountPageProjectService) {
	r = &AccountPageProjectService{}
	r.Options = opts
	r.Deployments = NewAccountPageProjectDeploymentService(opts...)
	r.Domains = NewAccountPageProjectDomainService(opts...)
	return
}

// Fetch a project by name.
func (r *AccountPageProjectService) Get(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *ProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *AccountPageProjectService) Update(ctx context.Context, accountIdentifier string, projectName string, body AccountPageProjectUpdateParams, opts ...option.RequestOption) (res *NewProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Destroy a project by name.
func (r *AccountPageProjectService) Delete(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Make a new project.
func (r *AccountPageProjectService) PagesProjectNewProject(ctx context.Context, accountIdentifier string, body AccountPageProjectPagesProjectNewProjectParams, opts ...option.RequestOption) (res *NewProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a list of all user projects.
func (r *AccountPageProjectService) PagesProjectGetProjects(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *ProjectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NewProjectResponse struct {
	Errors   []NewProjectResponseError   `json:"errors"`
	Messages []NewProjectResponseMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success NewProjectResponseSuccess `json:"success"`
	JSON    newProjectResponseJSON    `json:"-"`
}

// newProjectResponseJSON contains the JSON metadata for the struct
// [NewProjectResponse]
type newProjectResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewProjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NewProjectResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    newProjectResponseErrorJSON `json:"-"`
}

// newProjectResponseErrorJSON contains the JSON metadata for the struct
// [NewProjectResponseError]
type newProjectResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewProjectResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NewProjectResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    newProjectResponseMessageJSON `json:"-"`
}

// newProjectResponseMessageJSON contains the JSON metadata for the struct
// [NewProjectResponseMessage]
type newProjectResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewProjectResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type NewProjectResponseSuccess bool

const (
	NewProjectResponseSuccessTrue NewProjectResponseSuccess = true
)

type ProjectResponse struct {
	Errors   []ProjectResponseError   `json:"errors"`
	Messages []ProjectResponseMessage `json:"messages"`
	Result   ProjectResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ProjectResponseSuccess `json:"success"`
	JSON    projectResponseJSON    `json:"-"`
}

// projectResponseJSON contains the JSON metadata for the struct [ProjectResponse]
type projectResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    projectResponseErrorJSON `json:"-"`
}

// projectResponseErrorJSON contains the JSON metadata for the struct
// [ProjectResponseError]
type projectResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    projectResponseMessageJSON `json:"-"`
}

// projectResponseMessageJSON contains the JSON metadata for the struct
// [ProjectResponseMessage]
type projectResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectResponseResult struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig ProjectResponseResultBuildConfig `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment ProjectResponseResultCanonicalDeployment `json:"canonical_deployment,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectResponseResultDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains []interface{} `json:"domains"`
	// Most recent deployment to the repo.
	LatestDeployment ProjectResponseResultLatestDeployment `json:"latest_deployment,nullable"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string      `json:"production_branch"`
	Source           interface{} `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                    `json:"subdomain"`
	JSON      projectResponseResultJSON `json:"-"`
}

// projectResponseResultJSON contains the JSON metadata for the struct
// [ProjectResponseResult]
type projectResponseResultJSON struct {
	ID                  apijson.Field
	BuildConfig         apijson.Field
	CanonicalDeployment apijson.Field
	CreatedOn           apijson.Field
	DeploymentConfigs   apijson.Field
	Domains             apijson.Field
	LatestDeployment    apijson.Field
	Name                apijson.Field
	ProductionBranch    apijson.Field
	Source              apijson.Field
	Subdomain           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for the project build process.
type ProjectResponseResultBuildConfig struct {
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string `json:"root_dir,nullable"`
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string                               `json:"web_analytics_token,nullable"`
	JSON              projectResponseResultBuildConfigJSON `json:"-"`
}

// projectResponseResultBuildConfigJSON contains the JSON metadata for the struct
// [ProjectResponseResultBuildConfig]
type projectResponseResultBuildConfigJSON struct {
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectResponseResultBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Most recent deployment to the repo.
type ProjectResponseResultCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectResponseResultCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []ProjectResponseResultCanonicalDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                       `json:"url"`
	JSON projectResponseResultCanonicalDeploymentJSON `json:"-"`
}

// projectResponseResultCanonicalDeploymentJSON contains the JSON metadata for the
// struct [ProjectResponseResultCanonicalDeployment]
type projectResponseResultCanonicalDeploymentJSON struct {
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

func (r *ProjectResponseResultCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type ProjectResponseResultCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectResponseResultCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                        `json:"type"`
	JSON projectResponseResultCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectResponseResultCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct
// [ProjectResponseResultCanonicalDeploymentDeploymentTrigger]
type projectResponseResultCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type ProjectResponseResultCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                `json:"commit_message"`
	JSON          projectResponseResultCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectResponseResultCanonicalDeploymentDeploymentTriggerMetadataJSON contains
// the JSON metadata for the struct
// [ProjectResponseResultCanonicalDeploymentDeploymentTriggerMetadata]
type projectResponseResultCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectResponseResultCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type ProjectResponseResultCanonicalDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                            `json:"status"`
	JSON   projectResponseResultCanonicalDeploymentStageJSON `json:"-"`
}

// projectResponseResultCanonicalDeploymentStageJSON contains the JSON metadata for
// the struct [ProjectResponseResultCanonicalDeploymentStage]
type projectResponseResultCanonicalDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultCanonicalDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for deployments in a project.
type ProjectResponseResultDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectResponseResultDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production ProjectResponseResultDeploymentConfigsProduction `json:"production"`
	JSON       projectResponseResultDeploymentConfigsJSON       `json:"-"`
}

// projectResponseResultDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectResponseResultDeploymentConfigs]
type projectResponseResultDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for preview deploys.
type ProjectResponseResultDeploymentConfigsPreview struct {
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectResponseResultDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectResponseResultDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KvNamespaces ProjectResponseResultDeploymentConfigsPreviewKvNamespaces `json:"kv_namespaces"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectResponseResultDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectResponseResultDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings ProjectResponseResultDeploymentConfigsPreviewServiceBindings `json:"service_bindings,nullable"`
	JSON            projectResponseResultDeploymentConfigsPreviewJSON            `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectResponseResultDeploymentConfigsPreview]
type projectResponseResultDeploymentConfigsPreviewJSON struct {
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	KvNamespaces            apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	ServiceBindings         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 databases used for Pages Functions.
type ProjectResponseResultDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding ProjectResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectResponseResultDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewD1DatabasesJSON contains the JSON
// metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewD1Databases]
type projectResponseResultDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 binding.
type ProjectResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                `json:"id"`
	JSON projectResponseResultDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains
// the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding]
type projectResponseResultDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespaces]
type projectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object binding.
type ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                            `json:"namespace_id"`
	JSON        projectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type projectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variables for build configs.
type ProjectResponseResultDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectResponseResultDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewEnvVarsJSON contains the JSON
// metadata for the struct [ProjectResponseResultDeploymentConfigsPreviewEnvVars]
type projectResponseResultDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable.
type ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                      `json:"value"`
	JSON  projectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type projectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of environment variable (plain text or secret)
type ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type ProjectResponseResultDeploymentConfigsPreviewKvNamespaces struct {
	// KV binding.
	KvBinding ProjectResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding `json:"KV_BINDING"`
	JSON      projectResponseResultDeploymentConfigsPreviewKvNamespacesJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewKvNamespacesJSON contains the JSON
// metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewKvNamespaces]
type projectResponseResultDeploymentConfigsPreviewKvNamespacesJSON struct {
	KvBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewKvNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// KV binding.
type ProjectResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                 `json:"namespace_id"`
	JSON        projectResponseResultDeploymentConfigsPreviewKvNamespacesKvBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewKvNamespacesKvBindingJSON contains
// the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding]
type projectResponseResultDeploymentConfigsPreviewKvNamespacesKvBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectResponseResultDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectResponseResultDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewQueueProducersJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewQueueProducers]
type projectResponseResultDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer binding.
type ProjectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                              `json:"name"`
	JSON projectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type projectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 buckets used for Pages Functions.
type ProjectResponseResultDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding ProjectResponseResultDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectResponseResultDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewR2BucketsJSON contains the JSON
// metadata for the struct [ProjectResponseResultDeploymentConfigsPreviewR2Buckets]
type projectResponseResultDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 binding.
type ProjectResponseResultDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                              `json:"name"`
	JSON projectResponseResultDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewR2BucketsR2Binding]
type projectResponseResultDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Services used for Pages Functions.
type ProjectResponseResultDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding ProjectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectResponseResultDeploymentConfigsPreviewServiceBindingsJSON           `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewServiceBindingsJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewServiceBindings]
type projectResponseResultDeploymentConfigsPreviewServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service binding.
type ProjectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                         `json:"service"`
	JSON    projectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding]
type projectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for production deploys.
type ProjectResponseResultDeploymentConfigsProduction struct {
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectResponseResultDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectResponseResultDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KvNamespaces ProjectResponseResultDeploymentConfigsProductionKvNamespaces `json:"kv_namespaces"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectResponseResultDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectResponseResultDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings ProjectResponseResultDeploymentConfigsProductionServiceBindings `json:"service_bindings,nullable"`
	JSON            projectResponseResultDeploymentConfigsProductionJSON            `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionJSON contains the JSON metadata
// for the struct [ProjectResponseResultDeploymentConfigsProduction]
type projectResponseResultDeploymentConfigsProductionJSON struct {
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	KvNamespaces            apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	ServiceBindings         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 databases used for Pages Functions.
type ProjectResponseResultDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding ProjectResponseResultDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectResponseResultDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionD1DatabasesJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionD1Databases]
type projectResponseResultDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 binding.
type ProjectResponseResultDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                   `json:"id"`
	JSON projectResponseResultDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionD1DatabasesD1BindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionD1DatabasesD1Binding]
type projectResponseResultDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectResponseResultDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespaces]
type projectResponseResultDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object binding.
type ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                               `json:"namespace_id"`
	JSON        projectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type projectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variables for build configs.
type ProjectResponseResultDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectResponseResultDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionEnvVarsJSON contains the JSON
// metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionEnvVars]
type projectResponseResultDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable.
type ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                         `json:"value"`
	JSON  projectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type projectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of environment variable (plain text or secret)
type ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type ProjectResponseResultDeploymentConfigsProductionKvNamespaces struct {
	// KV binding.
	KvBinding ProjectResponseResultDeploymentConfigsProductionKvNamespacesKvBinding `json:"KV_BINDING"`
	JSON      projectResponseResultDeploymentConfigsProductionKvNamespacesJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionKvNamespacesJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionKvNamespaces]
type projectResponseResultDeploymentConfigsProductionKvNamespacesJSON struct {
	KvBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionKvNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// KV binding.
type ProjectResponseResultDeploymentConfigsProductionKvNamespacesKvBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                    `json:"namespace_id"`
	JSON        projectResponseResultDeploymentConfigsProductionKvNamespacesKvBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionKvNamespacesKvBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionKvNamespacesKvBinding]
type projectResponseResultDeploymentConfigsProductionKvNamespacesKvBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionKvNamespacesKvBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectResponseResultDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectResponseResultDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionQueueProducersJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionQueueProducers]
type projectResponseResultDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer binding.
type ProjectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                                 `json:"name"`
	JSON projectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type projectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 buckets used for Pages Functions.
type ProjectResponseResultDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding ProjectResponseResultDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectResponseResultDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionR2BucketsJSON contains the JSON
// metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionR2Buckets]
type projectResponseResultDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 binding.
type ProjectResponseResultDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                                 `json:"name"`
	JSON projectResponseResultDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionR2BucketsR2BindingJSON contains
// the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionR2BucketsR2Binding]
type projectResponseResultDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Services used for Pages Functions.
type ProjectResponseResultDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding ProjectResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectResponseResultDeploymentConfigsProductionServiceBindingsJSON           `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionServiceBindingsJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionServiceBindings]
type projectResponseResultDeploymentConfigsProductionServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service binding.
type ProjectResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                            `json:"service"`
	JSON    projectResponseResultDeploymentConfigsProductionServiceBindingsServiceBindingJSON `json:"-"`
}

// projectResponseResultDeploymentConfigsProductionServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [ProjectResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding]
type projectResponseResultDeploymentConfigsProductionServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Most recent deployment to the repo.
type ProjectResponseResultLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectResponseResultLatestDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []ProjectResponseResultLatestDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                    `json:"url"`
	JSON projectResponseResultLatestDeploymentJSON `json:"-"`
}

// projectResponseResultLatestDeploymentJSON contains the JSON metadata for the
// struct [ProjectResponseResultLatestDeployment]
type projectResponseResultLatestDeploymentJSON struct {
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

func (r *ProjectResponseResultLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type ProjectResponseResultLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectResponseResultLatestDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                     `json:"type"`
	JSON projectResponseResultLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectResponseResultLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectResponseResultLatestDeploymentDeploymentTrigger]
type projectResponseResultLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type ProjectResponseResultLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                             `json:"commit_message"`
	JSON          projectResponseResultLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectResponseResultLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectResponseResultLatestDeploymentDeploymentTriggerMetadata]
type projectResponseResultLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectResponseResultLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type ProjectResponseResultLatestDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                         `json:"status"`
	JSON   projectResponseResultLatestDeploymentStageJSON `json:"-"`
}

// projectResponseResultLatestDeploymentStageJSON contains the JSON metadata for
// the struct [ProjectResponseResultLatestDeploymentStage]
type projectResponseResultLatestDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponseResultLatestDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ProjectResponseSuccess bool

const (
	ProjectResponseSuccessTrue ProjectResponseSuccess = true
)

type ProjectsResponse struct {
	Errors     []ProjectsResponseError    `json:"errors"`
	Messages   []ProjectsResponseMessage  `json:"messages"`
	Result     []ProjectsResponseResult   `json:"result"`
	ResultInfo ProjectsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ProjectsResponseSuccess `json:"success"`
	JSON    projectsResponseJSON    `json:"-"`
}

// projectsResponseJSON contains the JSON metadata for the struct
// [ProjectsResponse]
type projectsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectsResponseError struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    projectsResponseErrorJSON `json:"-"`
}

// projectsResponseErrorJSON contains the JSON metadata for the struct
// [ProjectsResponseError]
type projectsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectsResponseMessage struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    projectsResponseMessageJSON `json:"-"`
}

// projectsResponseMessageJSON contains the JSON metadata for the struct
// [ProjectsResponseMessage]
type projectsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectsResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectsResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []ProjectsResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                     `json:"url"`
	JSON projectsResponseResultJSON `json:"-"`
}

// projectsResponseResultJSON contains the JSON metadata for the struct
// [ProjectsResponseResult]
type projectsResponseResultJSON struct {
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

func (r *ProjectsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type ProjectsResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectsResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                      `json:"type"`
	JSON projectsResponseResultDeploymentTriggerJSON `json:"-"`
}

// projectsResponseResultDeploymentTriggerJSON contains the JSON metadata for the
// struct [ProjectsResponseResultDeploymentTrigger]
type projectsResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type ProjectsResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                              `json:"commit_message"`
	JSON          projectsResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// projectsResponseResultDeploymentTriggerMetadataJSON contains the JSON metadata
// for the struct [ProjectsResponseResultDeploymentTriggerMetadata]
type projectsResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectsResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type ProjectsResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                          `json:"status"`
	JSON   projectsResponseResultStageJSON `json:"-"`
}

// projectsResponseResultStageJSON contains the JSON metadata for the struct
// [ProjectsResponseResultStage]
type projectsResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectsResponseResultInfo struct {
	Count      interface{}                    `json:"count"`
	Page       interface{}                    `json:"page"`
	PerPage    interface{}                    `json:"per_page"`
	TotalCount interface{}                    `json:"total_count"`
	JSON       projectsResponseResultInfoJSON `json:"-"`
}

// projectsResponseResultInfoJSON contains the JSON metadata for the struct
// [ProjectsResponseResultInfo]
type projectsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ProjectsResponseSuccess bool

const (
	ProjectsResponseSuccessTrue ProjectsResponseSuccess = true
)

type AccountPageProjectDeleteResponse = interface{}

type AccountPageProjectUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountPageProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountPageProjectPagesProjectNewProjectParams struct {
	// Configs for the project build process.
	BuildConfig param.Field[AccountPageProjectPagesProjectNewProjectParamsBuildConfig] `json:"build_config"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigs] `json:"deployment_configs"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
}

func (r AccountPageProjectPagesProjectNewProjectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type AccountPageProjectPagesProjectNewProjectParamsBuildConfig struct {
	// Command used to build project.
	BuildCommand param.Field[string] `json:"build_command"`
	// Output directory of the build.
	DestinationDir param.Field[string] `json:"destination_dir"`
	// Directory to run the command.
	RootDir param.Field[string] `json:"root_dir"`
	// The classifying tag for analytics.
	WebAnalyticsTag param.Field[string] `json:"web_analytics_tag"`
	// The auth token for analytics.
	WebAnalyticsToken param.Field[string] `json:"web_analytics_token"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsBuildConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreview] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProduction] `json:"production"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreview struct {
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVars] `json:"env_vars"`
	// KV namespaces used for Pages Functions.
	KvNamespaces param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespaces] `json:"kv_namespaces"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	ServiceBindings param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindings] `json:"service_bindings"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreview) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespaces struct {
	// KV binding.
	KvBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespacesKvBinding] `json:"KV_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespacesKvBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespacesKvBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindingsServiceBinding] `json:"SERVICE_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindingsServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProduction struct {
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVars] `json:"env_vars"`
	// KV namespaces used for Pages Functions.
	KvNamespaces param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespaces] `json:"kv_namespaces"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	ServiceBindings param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindings] `json:"service_bindings"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespaces struct {
	// KV binding.
	KvBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespacesKvBinding] `json:"KV_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespacesKvBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespacesKvBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindingsServiceBinding] `json:"SERVICE_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindingsServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
