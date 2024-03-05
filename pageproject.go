// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// PageProjectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageProjectService] method
// instead.
type PageProjectService struct {
	Options     []option.RequestOption
	Deployments *PageProjectDeploymentService
	Domains     *PageProjectDomainService
}

// NewPageProjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageProjectService(opts ...option.RequestOption) (r *PageProjectService) {
	r = &PageProjectService{}
	r.Options = opts
	r.Deployments = NewPageProjectDeploymentService(opts...)
	r.Domains = NewPageProjectDomainService(opts...)
	return
}

// Create a new project.
func (r *PageProjectService) New(ctx context.Context, params PageProjectNewParams, opts ...option.RequestOption) (res *PageProjectNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all user projects.
func (r *PageProjectService) List(ctx context.Context, query PageProjectListParams, opts ...option.RequestOption) (res *[]PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a project by name.
func (r *PageProjectService) Delete(ctx context.Context, projectName string, body PageProjectDeleteParams, opts ...option.RequestOption) (res *PageProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *PageProjectService) Edit(ctx context.Context, projectName string, params PageProjectEditParams, opts ...option.RequestOption) (res *PageProjectEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a project by name.
func (r *PageProjectService) Get(ctx context.Context, projectName string, query PageProjectGetParams, opts ...option.RequestOption) (res *PagesProjects, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Purge all cached build artifacts for a Pages project
func (r *PageProjectService) PurgeBuildCache(ctx context.Context, projectName string, body PageProjectPurgeBuildCacheParams, opts ...option.RequestOption) (res *PageProjectPurgeBuildCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PagesDeployments struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PagesDeploymentsDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []PagesDeploymentsStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string               `json:"url"`
	JSON pagesDeploymentsJSON `json:"-"`
}

// pagesDeploymentsJSON contains the JSON metadata for the struct
// [PagesDeployments]
type pagesDeploymentsJSON struct {
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

func (r *PagesDeployments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type PagesDeploymentsDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PagesDeploymentsDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                `json:"type"`
	JSON pagesDeploymentsDeploymentTriggerJSON `json:"-"`
}

// pagesDeploymentsDeploymentTriggerJSON contains the JSON metadata for the struct
// [PagesDeploymentsDeploymentTrigger]
type pagesDeploymentsDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesDeploymentsDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type PagesDeploymentsDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                        `json:"commit_message"`
	JSON          pagesDeploymentsDeploymentTriggerMetadataJSON `json:"-"`
}

// pagesDeploymentsDeploymentTriggerMetadataJSON contains the JSON metadata for the
// struct [PagesDeploymentsDeploymentTriggerMetadata]
type pagesDeploymentsDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PagesDeploymentsDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type PagesDeploymentsStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                    `json:"status"`
	JSON   pagesDeploymentsStageJSON `json:"-"`
}

// pagesDeploymentsStageJSON contains the JSON metadata for the struct
// [PagesDeploymentsStage]
type pagesDeploymentsStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesDeploymentsStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PagesDeploymentsParam struct {
}

func (r PagesDeploymentsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type PagesDeploymentsDeploymentTriggerParam struct {
	// Additional info about the trigger.
	Metadata param.Field[PagesDeploymentsDeploymentTriggerMetadataParam] `json:"metadata"`
}

func (r PagesDeploymentsDeploymentTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type PagesDeploymentsDeploymentTriggerMetadataParam struct {
}

func (r PagesDeploymentsDeploymentTriggerMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type PagesDeploymentsStageParam struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r PagesDeploymentsStageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PagesProjects struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig         PagesProjectsBuildConfig `json:"build_config"`
	CanonicalDeployment PagesDeployments         `json:"canonical_deployment"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs PagesProjectsDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains          []interface{}    `json:"domains"`
	LatestDeployment PagesDeployments `json:"latest_deployment"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string      `json:"production_branch"`
	Source           interface{} `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string            `json:"subdomain"`
	JSON      pagesProjectsJSON `json:"-"`
}

// pagesProjectsJSON contains the JSON metadata for the struct [PagesProjects]
type pagesProjectsJSON struct {
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

func (r *PagesProjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for the project build process.
type PagesProjectsBuildConfig struct {
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
	WebAnalyticsToken string                       `json:"web_analytics_token,nullable"`
	JSON              pagesProjectsBuildConfigJSON `json:"-"`
}

// pagesProjectsBuildConfigJSON contains the JSON metadata for the struct
// [PagesProjectsBuildConfig]
type pagesProjectsBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PagesProjectsBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for deployments in a project.
type PagesProjectsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview PagesProjectsDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production PagesProjectsDeploymentConfigsProduction `json:"production"`
	JSON       pagesProjectsDeploymentConfigsJSON       `json:"-"`
}

// pagesProjectsDeploymentConfigsJSON contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigs]
type pagesProjectsDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for preview deploys.
type PagesProjectsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings PagesProjectsDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases PagesProjectsDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces PagesProjectsDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars PagesProjectsDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces PagesProjectsDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement PagesProjectsDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers PagesProjectsDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets PagesProjectsDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings PagesProjectsDeploymentConfigsPreviewServiceBindings `json:"service_bindings,nullable"`
	JSON            pagesProjectsDeploymentConfigsPreviewJSON            `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewJSON contains the JSON metadata for the
// struct [PagesProjectsDeploymentConfigsPreview]
type pagesProjectsDeploymentConfigsPreviewJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	KVNamespaces            apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	ServiceBindings         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Constellation bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding PagesProjectsDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      pagesProjectsDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewAIBindingsJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsPreviewAIBindings]
type pagesProjectsDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// AI binding.
type PagesProjectsDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID interface{}                                                  `json:"project_id"`
	JSON      pagesProjectsDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewAIBindingsAIBindingJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewAIBindingsAIBinding]
type pagesProjectsDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasets]
type pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine binding.
type PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                 `json:"dataset"`
	JSON    pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 databases used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding PagesProjectsDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      pagesProjectsDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewD1DatabasesJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsPreviewD1Databases]
type pagesProjectsDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 binding.
type PagesProjectsDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                        `json:"id"`
	JSON pagesProjectsDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewD1DatabasesD1Binding]
type pagesProjectsDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object namespaces used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding PagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewDurableObjectNamespaces]
type pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object binding.
type PagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                    `json:"namespace_id"`
	JSON        pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variables for build configs.
type PagesProjectsDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                pagesProjectsDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewEnvVarsJSON contains the JSON metadata for
// the struct [PagesProjectsDeploymentConfigsPreviewEnvVars]
type pagesProjectsDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable.
type PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                              `json:"value"`
	JSON  pagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type pagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of environment variable (plain text or secret)
type PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding PagesProjectsDeploymentConfigsPreviewKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      pagesProjectsDeploymentConfigsPreviewKVNamespacesJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewKVNamespacesJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsPreviewKVNamespaces]
type pagesProjectsDeploymentConfigsPreviewKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// KV binding.
type PagesProjectsDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                         `json:"namespace_id"`
	JSON        pagesProjectsDeploymentConfigsPreviewKVNamespacesKVBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewKVNamespacesKVBindingJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewKVNamespacesKVBinding]
type pagesProjectsDeploymentConfigsPreviewKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Placement setting used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                             `json:"mode"`
	JSON pagesProjectsDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewPlacementJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsPreviewPlacement]
type pagesProjectsDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding PagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 pagesProjectsDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewQueueProducersJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsPreviewQueueProducers]
type pagesProjectsDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer binding.
type PagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                      `json:"name"`
	JSON pagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type pagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 buckets used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding PagesProjectsDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      pagesProjectsDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewR2BucketsJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsPreviewR2Buckets]
type pagesProjectsDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 binding.
type PagesProjectsDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                      `json:"name"`
	JSON pagesProjectsDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewR2BucketsR2Binding]
type pagesProjectsDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Services used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding PagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           pagesProjectsDeploymentConfigsPreviewServiceBindingsJSON           `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewServiceBindingsJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsPreviewServiceBindings]
type pagesProjectsDeploymentConfigsPreviewServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service binding.
type PagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                 `json:"service"`
	JSON    pagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBindingJSON contains
// the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBinding]
type pagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configs for production deploys.
type PagesProjectsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings PagesProjectsDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases PagesProjectsDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces PagesProjectsDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars PagesProjectsDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces PagesProjectsDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement PagesProjectsDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers PagesProjectsDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets PagesProjectsDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings PagesProjectsDeploymentConfigsProductionServiceBindings `json:"service_bindings,nullable"`
	JSON            pagesProjectsDeploymentConfigsProductionJSON            `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionJSON contains the JSON metadata for the
// struct [PagesProjectsDeploymentConfigsProduction]
type pagesProjectsDeploymentConfigsProductionJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	KVNamespaces            apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	ServiceBindings         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Constellation bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding PagesProjectsDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      pagesProjectsDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionAIBindingsJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsProductionAIBindings]
type pagesProjectsDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// AI binding.
type PagesProjectsDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID interface{}                                                     `json:"project_id"`
	JSON      pagesProjectsDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionAIBindingsAIBindingJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionAIBindingsAIBinding]
type pagesProjectsDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasets]
type pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine binding.
type PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                    `json:"dataset"`
	JSON    pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 databases used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding PagesProjectsDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      pagesProjectsDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionD1DatabasesJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsProductionD1Databases]
type pagesProjectsDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// D1 binding.
type PagesProjectsDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                           `json:"id"`
	JSON pagesProjectsDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionD1DatabasesD1BindingJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionD1DatabasesD1Binding]
type pagesProjectsDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object namespaces used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding PagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionDurableObjectNamespaces]
type pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Durabble Object binding.
type PagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                       `json:"namespace_id"`
	JSON        pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variables for build configs.
type PagesProjectsDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                pagesProjectsDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionEnvVarsJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsProductionEnvVars]
type pagesProjectsDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable.
type PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                 `json:"value"`
	JSON  pagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON contains
// the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type pagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of environment variable (plain text or secret)
type PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding PagesProjectsDeploymentConfigsProductionKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      pagesProjectsDeploymentConfigsProductionKVNamespacesJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionKVNamespacesJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsProductionKVNamespaces]
type pagesProjectsDeploymentConfigsProductionKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// KV binding.
type PagesProjectsDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                            `json:"namespace_id"`
	JSON        pagesProjectsDeploymentConfigsProductionKVNamespacesKVBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionKVNamespacesKVBindingJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionKVNamespacesKVBinding]
type pagesProjectsDeploymentConfigsProductionKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Placement setting used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                `json:"mode"`
	JSON pagesProjectsDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionPlacementJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsProductionPlacement]
type pagesProjectsDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding PagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 pagesProjectsDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionQueueProducersJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsProductionQueueProducers]
type pagesProjectsDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Queue Producer binding.
type PagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                         `json:"name"`
	JSON pagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type pagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 buckets used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding PagesProjectsDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      pagesProjectsDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionR2BucketsJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsProductionR2Buckets]
type pagesProjectsDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// R2 binding.
type PagesProjectsDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                         `json:"name"`
	JSON pagesProjectsDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionR2BucketsR2BindingJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsProductionR2BucketsR2Binding]
type pagesProjectsDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Services used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding PagesProjectsDeploymentConfigsProductionServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           pagesProjectsDeploymentConfigsProductionServiceBindingsJSON           `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionServiceBindingsJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsProductionServiceBindings]
type pagesProjectsDeploymentConfigsProductionServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service binding.
type PagesProjectsDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                    `json:"service"`
	JSON    pagesProjectsDeploymentConfigsProductionServiceBindingsServiceBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionServiceBindingsServiceBinding]
type pagesProjectsDeploymentConfigsProductionServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [PageProjectNewResponseUnknown],
// [PageProjectNewResponseArray] or [shared.UnionString].
type PageProjectNewResponse interface {
	ImplementsPageProjectNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectNewResponseArray []interface{}

func (r PageProjectNewResponseArray) ImplementsPageProjectNewResponse() {}

type PageProjectDeleteResponse = interface{}

// Union satisfied by [PageProjectEditResponseUnknown],
// [PageProjectEditResponseArray] or [shared.UnionString].
type PageProjectEditResponse interface {
	ImplementsPageProjectEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectEditResponseArray []interface{}

func (r PageProjectEditResponseArray) ImplementsPageProjectEditResponse() {}

type PageProjectPurgeBuildCacheResponse = interface{}

type PageProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig         param.Field[PageProjectNewParamsBuildConfig] `json:"build_config"`
	CanonicalDeployment param.Field[PagesDeploymentsParam]           `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[PageProjectNewParamsDeploymentConfigs] `json:"deployment_configs"`
	LatestDeployment  param.Field[PagesDeploymentsParam]                 `json:"latest_deployment"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
}

func (r PageProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type PageProjectNewParamsBuildConfig struct {
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

func (r PageProjectNewParamsBuildConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type PageProjectNewParamsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview param.Field[PageProjectNewParamsDeploymentConfigsPreview] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[PageProjectNewParamsDeploymentConfigsProduction] `json:"production"`
}

func (r PageProjectNewParamsDeploymentConfigs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type PageProjectNewParamsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[PageProjectNewParamsDeploymentConfigsPreviewAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[PageProjectNewParamsDeploymentConfigsPreviewD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[PageProjectNewParamsDeploymentConfigsPreviewEnvVars] `json:"env_vars"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[PageProjectNewParamsDeploymentConfigsPreviewKVNamespaces] `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement param.Field[PageProjectNewParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[PageProjectNewParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[PageProjectNewParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	ServiceBindings param.Field[PageProjectNewParamsDeploymentConfigsPreviewServiceBindings] `json:"service_bindings"`
}

func (r PageProjectNewParamsDeploymentConfigsPreview) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding param.Field[PageProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type PageProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID param.Field[interface{}] `json:"project_id"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding param.Field[PageProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type PageProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type PageProjectNewParamsDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding param.Field[PageProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding] `json:"KV_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type PageProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[PageProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type PageProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[PageProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type PageProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding param.Field[PageProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding] `json:"SERVICE_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewServiceBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type PageProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r PageProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type PageProjectNewParamsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[PageProjectNewParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[PageProjectNewParamsDeploymentConfigsProductionD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[PageProjectNewParamsDeploymentConfigsProductionEnvVars] `json:"env_vars"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[PageProjectNewParamsDeploymentConfigsProductionKVNamespaces] `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement param.Field[PageProjectNewParamsDeploymentConfigsProductionPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[PageProjectNewParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[PageProjectNewParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	ServiceBindings param.Field[PageProjectNewParamsDeploymentConfigsProductionServiceBindings] `json:"service_bindings"`
}

func (r PageProjectNewParamsDeploymentConfigsProduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding param.Field[PageProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type PageProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID param.Field[interface{}] `json:"project_id"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding param.Field[PageProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type PageProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type PageProjectNewParamsDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding param.Field[PageProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding] `json:"KV_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type PageProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[PageProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type PageProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[PageProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type PageProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type PageProjectNewParamsDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding param.Field[PageProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding] `json:"SERVICE_BINDING"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionServiceBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type PageProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r PageProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageProjectNewResponseEnvelope struct {
	Errors   []PageProjectNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectNewResponseEnvelopeJSON    `json:"-"`
}

// pageProjectNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageProjectNewResponseEnvelope]
type pageProjectNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageProjectNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageProjectNewResponseEnvelopeErrors]
type pageProjectNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pageProjectNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageProjectNewResponseEnvelopeMessages]
type pageProjectNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectNewResponseEnvelopeSuccess bool

const (
	PageProjectNewResponseEnvelopeSuccessTrue PageProjectNewResponseEnvelopeSuccess = true
)

type PageProjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectListResponseEnvelope struct {
	Errors   []PageProjectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PagesDeployments                        `json:"result,required"`
	// Whether the API call was successful
	Success    PageProjectListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageProjectListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageProjectListResponseEnvelopeJSON       `json:"-"`
}

// pageProjectListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageProjectListResponseEnvelope]
type pageProjectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    pageProjectListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageProjectListResponseEnvelopeErrors]
type pageProjectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    pageProjectListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageProjectListResponseEnvelopeMessages]
type pageProjectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectListResponseEnvelopeSuccess bool

const (
	PageProjectListResponseEnvelopeSuccessTrue PageProjectListResponseEnvelopeSuccess = true
)

type PageProjectListResponseEnvelopeResultInfo struct {
	Count      interface{}                                   `json:"count"`
	Page       interface{}                                   `json:"page"`
	PerPage    interface{}                                   `json:"per_page"`
	TotalCount interface{}                                   `json:"total_count"`
	JSON       pageProjectListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageProjectListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [PageProjectListResponseEnvelopeResultInfo]
type pageProjectListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectEditParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r PageProjectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PageProjectEditResponseEnvelope struct {
	Errors   []PageProjectEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectEditResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectEditResponseEnvelopeJSON    `json:"-"`
}

// pageProjectEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageProjectEditResponseEnvelope]
type pageProjectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    pageProjectEditResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageProjectEditResponseEnvelopeErrors]
type pageProjectEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    pageProjectEditResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageProjectEditResponseEnvelopeMessages]
type pageProjectEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectEditResponseEnvelopeSuccess bool

const (
	PageProjectEditResponseEnvelopeSuccessTrue PageProjectEditResponseEnvelopeSuccess = true
)

type PageProjectGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectGetResponseEnvelope struct {
	Errors   []PageProjectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesProjects                            `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectGetResponseEnvelopeJSON    `json:"-"`
}

// pageProjectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageProjectGetResponseEnvelope]
type pageProjectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageProjectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageProjectGetResponseEnvelopeErrors]
type pageProjectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pageProjectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageProjectGetResponseEnvelopeMessages]
type pageProjectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectGetResponseEnvelopeSuccess bool

const (
	PageProjectGetResponseEnvelopeSuccessTrue PageProjectGetResponseEnvelopeSuccess = true
)

type PageProjectPurgeBuildCacheParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
