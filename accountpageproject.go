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
func (r *AccountPageProjectService) Get(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *AccountPageProjectService) Update(ctx context.Context, accountIdentifier string, projectName string, body AccountPageProjectUpdateParams, opts ...option.RequestOption) (res *AccountPageProjectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a project by name.
func (r *AccountPageProjectService) Delete(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new project.
func (r *AccountPageProjectService) PagesProjectNewProject(ctx context.Context, accountIdentifier string, body AccountPageProjectPagesProjectNewProjectParams, opts ...option.RequestOption) (res *AccountPageProjectPagesProjectNewProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a list of all user projects.
func (r *AccountPageProjectService) PagesProjectGetProjects(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectPagesProjectGetProjectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPageProjectGetResponse struct {
	Errors   []AccountPageProjectGetResponseError   `json:"errors"`
	Messages []AccountPageProjectGetResponseMessage `json:"messages"`
	Result   AccountPageProjectGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectGetResponseSuccess `json:"success"`
	JSON    accountPageProjectGetResponseJSON    `json:"-"`
}

// accountPageProjectGetResponseJSON contains the JSON metadata for the struct
// [AccountPageProjectGetResponse]
type accountPageProjectGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountPageProjectGetResponseErrorJSON `json:"-"`
}

// accountPageProjectGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountPageProjectGetResponseError]
type accountPageProjectGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountPageProjectGetResponseMessageJSON `json:"-"`
}

// accountPageProjectGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountPageProjectGetResponseMessage]
type accountPageProjectGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectGetResponseResult struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig         AccountPageProjectGetResponseResultBuildConfig         `json:"build_config"`
	CanonicalDeployment AccountPageProjectGetResponseResultCanonicalDeployment `json:"canonical_deployment"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs AccountPageProjectGetResponseResultDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains          []interface{}                                       `json:"domains"`
	LatestDeployment AccountPageProjectGetResponseResultLatestDeployment `json:"latest_deployment"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string      `json:"production_branch"`
	Source           interface{} `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                                  `json:"subdomain"`
	JSON      accountPageProjectGetResponseResultJSON `json:"-"`
}

// accountPageProjectGetResponseResultJSON contains the JSON metadata for the
// struct [AccountPageProjectGetResponseResult]
type accountPageProjectGetResponseResultJSON struct {
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

func (r *AccountPageProjectGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for the project build process.
type AccountPageProjectGetResponseResultBuildConfig struct {
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string `json:"root_dir,nullable"`
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string                                             `json:"web_analytics_token,nullable"`
	JSON              accountPageProjectGetResponseResultBuildConfigJSON `json:"-"`
}

// accountPageProjectGetResponseResultBuildConfigJSON contains the JSON metadata
// for the struct [AccountPageProjectGetResponseResultBuildConfig]
type accountPageProjectGetResponseResultBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectGetResponseResultCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectGetResponseResultCanonicalDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                     `json:"url"`
	JSON accountPageProjectGetResponseResultCanonicalDeploymentJSON `json:"-"`
}

// accountPageProjectGetResponseResultCanonicalDeploymentJSON contains the JSON
// metadata for the struct [AccountPageProjectGetResponseResultCanonicalDeployment]
type accountPageProjectGetResponseResultCanonicalDeploymentJSON struct {
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

func (r *AccountPageProjectGetResponseResultCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                      `json:"type"`
	JSON accountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTrigger]
type accountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                              `json:"commit_message"`
	JSON          accountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadata]
type accountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectGetResponseResultCanonicalDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                          `json:"status"`
	JSON   accountPageProjectGetResponseResultCanonicalDeploymentStageJSON `json:"-"`
}

// accountPageProjectGetResponseResultCanonicalDeploymentStageJSON contains the
// JSON metadata for the struct
// [AccountPageProjectGetResponseResultCanonicalDeploymentStage]
type accountPageProjectGetResponseResultCanonicalDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultCanonicalDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for deployments in a project.
type AccountPageProjectGetResponseResultDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview AccountPageProjectGetResponseResultDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production AccountPageProjectGetResponseResultDeploymentConfigsProduction `json:"production"`
	JSON       accountPageProjectGetResponseResultDeploymentConfigsJSON       `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsJSON contains the JSON
// metadata for the struct [AccountPageProjectGetResponseResultDeploymentConfigs]
type accountPageProjectGetResponseResultDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for preview deploys.
type AccountPageProjectGetResponseResultDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KvNamespaces AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement AccountPageProjectGetResponseResultDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindings `json:"service_bindings,nullable"`
	JSON            accountPageProjectGetResponseResultDeploymentConfigsPreviewJSON            `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewJSON contains the
// JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreview]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	KvNamespaces            apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	ServiceBindings         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Constellation bindings used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindings]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// AI binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID interface{}                                                                        `json:"project_id"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBinding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine bindings used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   accountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasets]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                                       `json:"dataset"`
	JSON    accountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 databases used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1Databases]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                              `json:"id"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1BindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object namespaces used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespaces]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                                          `json:"namespace_id"`
	JSON        accountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variables for build configs.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                accountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsJSON contains
// the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVars]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                                    `json:"value"`
	JSON  accountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of environment variable (plain text or secret)
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText AccountPageProjectGetResponseResultDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespaces struct {
	// KV binding.
	KvBinding AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding `json:"KV_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespaces]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesJSON struct {
	KvBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// KV binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                               `json:"namespace_id"`
	JSON        accountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewKvNamespacesKvBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Placement setting used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                                   `json:"mode"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewPlacementJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewPlacement]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer bindings used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 accountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducers]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                                            `json:"name"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 buckets used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2Buckets]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                                            `json:"name"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2BindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2Binding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Services used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           accountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsJSON           `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindings]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service binding.
type AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                                       `json:"service"`
	JSON    accountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding]
type accountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsPreviewServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for production deploys.
type AccountPageProjectGetResponseResultDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases AccountPageProjectGetResponseResultDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KvNamespaces AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement AccountPageProjectGetResponseResultDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets AccountPageProjectGetResponseResultDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindings `json:"service_bindings,nullable"`
	JSON            accountPageProjectGetResponseResultDeploymentConfigsProductionJSON            `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionJSON contains the
// JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProduction]
type accountPageProjectGetResponseResultDeploymentConfigsProductionJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	KvNamespaces            apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	ServiceBindings         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Constellation bindings used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindings]
type accountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// AI binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID interface{}                                                                           `json:"project_id"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBinding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine bindings used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   accountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasets]
type accountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                                          `json:"dataset"`
	JSON    accountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 databases used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding AccountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionD1Databases]
type accountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                                 `json:"id"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1BindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1Binding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object namespaces used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespaces]
type accountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                                             `json:"namespace_id"`
	JSON        accountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variables for build configs.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                accountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVars]
type accountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                                       `json:"value"`
	JSON  accountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type accountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of environment variable (plain text or secret)
type AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText AccountPageProjectGetResponseResultDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespaces struct {
	// KV binding.
	KvBinding AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBinding `json:"KV_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespaces]
type accountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesJSON struct {
	KvBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// KV binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                                  `json:"namespace_id"`
	JSON        accountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBinding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionKvNamespacesKvBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Placement setting used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                                      `json:"mode"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionPlacementJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionPlacement]
type accountPageProjectGetResponseResultDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer bindings used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 accountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducers]
type accountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                                               `json:"name"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 buckets used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding AccountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      accountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionR2Buckets]
type accountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                                               `json:"name"`
	JSON accountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2BindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2Binding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Services used for Pages Functions.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           accountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsJSON           `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindings]
type accountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service binding.
type AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                                          `json:"service"`
	JSON    accountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBindingJSON `json:"-"`
}

// accountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding]
type accountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultDeploymentConfigsProductionServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectGetResponseResultLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectGetResponseResultLatestDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectGetResponseResultLatestDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                  `json:"url"`
	JSON accountPageProjectGetResponseResultLatestDeploymentJSON `json:"-"`
}

// accountPageProjectGetResponseResultLatestDeploymentJSON contains the JSON
// metadata for the struct [AccountPageProjectGetResponseResultLatestDeployment]
type accountPageProjectGetResponseResultLatestDeploymentJSON struct {
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

func (r *AccountPageProjectGetResponseResultLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectGetResponseResultLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                   `json:"type"`
	JSON accountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultLatestDeploymentDeploymentTrigger]
type accountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                           `json:"commit_message"`
	JSON          accountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadata]
type accountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectGetResponseResultLatestDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                       `json:"status"`
	JSON   accountPageProjectGetResponseResultLatestDeploymentStageJSON `json:"-"`
}

// accountPageProjectGetResponseResultLatestDeploymentStageJSON contains the JSON
// metadata for the struct
// [AccountPageProjectGetResponseResultLatestDeploymentStage]
type accountPageProjectGetResponseResultLatestDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectGetResponseResultLatestDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectGetResponseSuccess bool

const (
	AccountPageProjectGetResponseSuccessTrue AccountPageProjectGetResponseSuccess = true
)

type AccountPageProjectUpdateResponse struct {
	Errors   []AccountPageProjectUpdateResponseError   `json:"errors"`
	Messages []AccountPageProjectUpdateResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectUpdateResponseSuccess `json:"success"`
	JSON    accountPageProjectUpdateResponseJSON    `json:"-"`
}

// accountPageProjectUpdateResponseJSON contains the JSON metadata for the struct
// [AccountPageProjectUpdateResponse]
type accountPageProjectUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountPageProjectUpdateResponseErrorJSON `json:"-"`
}

// accountPageProjectUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountPageProjectUpdateResponseError]
type accountPageProjectUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountPageProjectUpdateResponseMessageJSON `json:"-"`
}

// accountPageProjectUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountPageProjectUpdateResponseMessage]
type accountPageProjectUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectUpdateResponseSuccess bool

const (
	AccountPageProjectUpdateResponseSuccessTrue AccountPageProjectUpdateResponseSuccess = true
)

type AccountPageProjectDeleteResponse = interface{}

type AccountPageProjectPagesProjectNewProjectResponse struct {
	Errors   []AccountPageProjectPagesProjectNewProjectResponseError   `json:"errors"`
	Messages []AccountPageProjectPagesProjectNewProjectResponseMessage `json:"messages"`
	Result   interface{}                                               `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectPagesProjectNewProjectResponseSuccess `json:"success"`
	JSON    accountPageProjectPagesProjectNewProjectResponseJSON    `json:"-"`
}

// accountPageProjectPagesProjectNewProjectResponseJSON contains the JSON metadata
// for the struct [AccountPageProjectPagesProjectNewProjectResponse]
type accountPageProjectPagesProjectNewProjectResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectNewProjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectPagesProjectNewProjectResponseError struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountPageProjectPagesProjectNewProjectResponseErrorJSON `json:"-"`
}

// accountPageProjectPagesProjectNewProjectResponseErrorJSON contains the JSON
// metadata for the struct [AccountPageProjectPagesProjectNewProjectResponseError]
type accountPageProjectPagesProjectNewProjectResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectNewProjectResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectPagesProjectNewProjectResponseMessage struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountPageProjectPagesProjectNewProjectResponseMessageJSON `json:"-"`
}

// accountPageProjectPagesProjectNewProjectResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountPageProjectPagesProjectNewProjectResponseMessage]
type accountPageProjectPagesProjectNewProjectResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectNewProjectResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectPagesProjectNewProjectResponseSuccess bool

const (
	AccountPageProjectPagesProjectNewProjectResponseSuccessTrue AccountPageProjectPagesProjectNewProjectResponseSuccess = true
)

type AccountPageProjectPagesProjectGetProjectsResponse struct {
	Errors     []AccountPageProjectPagesProjectGetProjectsResponseError    `json:"errors"`
	Messages   []AccountPageProjectPagesProjectGetProjectsResponseMessage  `json:"messages"`
	Result     []AccountPageProjectPagesProjectGetProjectsResponseResult   `json:"result"`
	ResultInfo AccountPageProjectPagesProjectGetProjectsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountPageProjectPagesProjectGetProjectsResponseSuccess `json:"success"`
	JSON    accountPageProjectPagesProjectGetProjectsResponseJSON    `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseJSON contains the JSON metadata
// for the struct [AccountPageProjectPagesProjectGetProjectsResponse]
type accountPageProjectPagesProjectGetProjectsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectPagesProjectGetProjectsResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountPageProjectPagesProjectGetProjectsResponseErrorJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseErrorJSON contains the JSON
// metadata for the struct [AccountPageProjectPagesProjectGetProjectsResponseError]
type accountPageProjectPagesProjectGetProjectsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectPagesProjectGetProjectsResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountPageProjectPagesProjectGetProjectsResponseMessageJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountPageProjectPagesProjectGetProjectsResponseMessage]
type accountPageProjectPagesProjectGetProjectsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectPagesProjectGetProjectsResponseResult struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []AccountPageProjectPagesProjectGetProjectsResponseResultStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                                      `json:"url"`
	JSON accountPageProjectPagesProjectGetProjectsResponseResultJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountPageProjectPagesProjectGetProjectsResponseResult]
type accountPageProjectPagesProjectGetProjectsResponseResultJSON struct {
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

func (r *AccountPageProjectPagesProjectGetProjectsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                                       `json:"type"`
	JSON accountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerJSON
// contains the JSON metadata for the struct
// [AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTrigger]
type accountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                               `json:"commit_message"`
	JSON          accountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadataJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadataJSON
// contains the JSON metadata for the struct
// [AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadata]
type accountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponseResultDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type AccountPageProjectPagesProjectGetProjectsResponseResultStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                                           `json:"status"`
	JSON   accountPageProjectPagesProjectGetProjectsResponseResultStageJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseResultStageJSON contains the
// JSON metadata for the struct
// [AccountPageProjectPagesProjectGetProjectsResponseResultStage]
type accountPageProjectPagesProjectGetProjectsResponseResultStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponseResultStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectPagesProjectGetProjectsResponseResultInfo struct {
	Count      interface{}                                                     `json:"count"`
	Page       interface{}                                                     `json:"page"`
	PerPage    interface{}                                                     `json:"per_page"`
	TotalCount interface{}                                                     `json:"total_count"`
	JSON       accountPageProjectPagesProjectGetProjectsResponseResultInfoJSON `json:"-"`
}

// accountPageProjectPagesProjectGetProjectsResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountPageProjectPagesProjectGetProjectsResponseResultInfo]
type accountPageProjectPagesProjectGetProjectsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPagesProjectGetProjectsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectPagesProjectGetProjectsResponseSuccess bool

const (
	AccountPageProjectPagesProjectGetProjectsResponseSuccessTrue AccountPageProjectPagesProjectGetProjectsResponseSuccess = true
)

type AccountPageProjectUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountPageProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountPageProjectPagesProjectNewProjectParams struct {
	// Configs for the project build process.
	BuildConfig         param.Field[AccountPageProjectPagesProjectNewProjectParamsBuildConfig]         `json:"build_config"`
	CanonicalDeployment param.Field[AccountPageProjectPagesProjectNewProjectParamsCanonicalDeployment] `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigs] `json:"deployment_configs"`
	LatestDeployment  param.Field[AccountPageProjectPagesProjectNewProjectParamsLatestDeployment]  `json:"latest_deployment"`
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
	// Enable build caching for the project.
	BuildCaching param.Field[bool] `json:"build_caching"`
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

type AccountPageProjectPagesProjectNewProjectParamsCanonicalDeployment struct {
}

func (r AccountPageProjectPagesProjectNewProjectParamsCanonicalDeployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata param.Field[AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentDeploymentTriggerMetadata] `json:"metadata"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentDeploymentTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentDeploymentTriggerMetadata struct {
}

func (r AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentDeploymentTriggerMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentStage struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsCanonicalDeploymentStage) MarshalJSON() (data []byte, err error) {
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
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
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
	// Placement setting used for Pages Functions.
	Placement param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
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

// Constellation bindings used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID param.Field[interface{}] `json:"project_id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
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

// Placement setting used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewPlacement) MarshalJSON() (data []byte, err error) {
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
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
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
	// Placement setting used for Pages Functions.
	Placement param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionPlacement] `json:"placement"`
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

// Constellation bindings used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID param.Field[interface{}] `json:"project_id"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
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

// Placement setting used for Pages Functions.
type AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionPlacement) MarshalJSON() (data []byte, err error) {
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

type AccountPageProjectPagesProjectNewProjectParamsLatestDeployment struct {
}

func (r AccountPageProjectPagesProjectNewProjectParamsLatestDeployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata param.Field[AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentDeploymentTriggerMetadata] `json:"metadata"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentDeploymentTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentDeploymentTriggerMetadata struct {
}

func (r AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentDeploymentTriggerMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentStage struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectPagesProjectNewProjectParamsLatestDeploymentStage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
