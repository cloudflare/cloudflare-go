// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// ProjectService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options     []option.RequestOption
	Deployments *ProjectDeploymentService
	Domains     *ProjectDomainService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.Deployments = NewProjectDeploymentService(opts...)
	r.Domains = NewProjectDomainService(opts...)
	return
}

// Create a new project.
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectNewResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all user projects.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Deployment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch a list of all user projects.
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Deployment] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a project by name.
func (r *ProjectService) Delete(ctx context.Context, projectName string, body ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *ProjectService) Edit(ctx context.Context, projectName string, params ProjectEditParams, opts ...option.RequestOption) (res *ProjectEditResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectEditResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a project by name.
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Purge all cached build artifacts for a Pages project
func (r *ProjectService) PurgeBuildCache(ctx context.Context, projectName string, body ProjectPurgeBuildCacheParams, opts ...option.RequestOption) (res *ProjectPurgeBuildCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type Deployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []Stage `json:"stages"`
	// The live URL to view this deployment.
	URL  string         `json:"url"`
	JSON deploymentJSON `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
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

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type DeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                          `json:"type"`
	JSON deploymentDeploymentTriggerJSON `json:"-"`
}

// deploymentDeploymentTriggerJSON contains the JSON metadata for the struct
// [DeploymentDeploymentTrigger]
type deploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type DeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                  `json:"commit_message"`
	JSON          deploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentDeploymentTriggerMetadataJSON contains the JSON metadata for the
// struct [DeploymentDeploymentTriggerMetadata]
type deploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

type DeploymentParam struct {
}

func (r DeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type DeploymentDeploymentTriggerParam struct {
	// Additional info about the trigger.
	Metadata param.Field[DeploymentDeploymentTriggerMetadataParam] `json:"metadata"`
}

func (r DeploymentDeploymentTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type DeploymentDeploymentTriggerMetadataParam struct {
}

func (r DeploymentDeploymentTriggerMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Project struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig         ProjectBuildConfig `json:"build_config"`
	CanonicalDeployment Deployment         `json:"canonical_deployment"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains          []interface{} `json:"domains"`
	LatestDeployment Deployment    `json:"latest_deployment"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string      `json:"production_branch"`
	Source           interface{} `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string      `json:"subdomain"`
	JSON      projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
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

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectBuildConfig struct {
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
	WebAnalyticsToken string                 `json:"web_analytics_token,nullable"`
	JSON              projectBuildConfigJSON `json:"-"`
}

// projectBuildConfigJSON contains the JSON metadata for the struct
// [ProjectBuildConfig]
type projectBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production ProjectDeploymentConfigsProduction `json:"production"`
	JSON       projectDeploymentConfigsJSON       `json:"-"`
}

// projectDeploymentConfigsJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigs]
type projectDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectDeploymentConfigsPreviewBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectDeploymentConfigsPreviewHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectDeploymentConfigsPreviewMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectDeploymentConfigsPreviewServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectDeploymentConfigsPreviewVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectDeploymentConfigsPreviewJSON              `json:"-"`
}

// projectDeploymentConfigsPreviewJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigsPreview]
type projectDeploymentConfigsPreviewJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	Browsers                apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	HyperdriveBindings      apijson.Field
	KVNamespaces            apijson.Field
	MTLSCertificates        apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	Services                apijson.Field
	VectorizeBindings       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding ProjectDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// projectDeploymentConfigsPreviewAIBindingsJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewAIBindings]
type projectDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID interface{}                                            `json:"project_id"`
	JSON      projectDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewAIBindingsAIBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewAIBindingsAIBinding]
type projectDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewAnalyticsEngineDatasets]
type projectDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                           `json:"dataset"`
	JSON    projectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser interface{}                                 `json:"BROWSER"`
	JSON    projectDeploymentConfigsPreviewBrowsersJSON `json:"-"`
}

// projectDeploymentConfigsPreviewBrowsersJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewBrowsers]
type projectDeploymentConfigsPreviewBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding ProjectDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// projectDeploymentConfigsPreviewD1DatabasesJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewD1Databases]
type projectDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                  `json:"id"`
	JSON projectDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewD1DatabasesD1Binding]
type projectDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// projectDeploymentConfigsPreviewDurableObjectNamespacesJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewDurableObjectNamespaces]
type projectDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                              `json:"namespace_id"`
	JSON        projectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type projectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// projectDeploymentConfigsPreviewEnvVarsJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewEnvVars]
type projectDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                        `json:"value"`
	JSON  projectDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type projectDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectDeploymentConfigsPreviewHyperdriveBindingsJSON       `json:"-"`
}

// projectDeploymentConfigsPreviewHyperdriveBindingsJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewHyperdriveBindings]
type projectDeploymentConfigsPreviewHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID   string                                                          `json:"id"`
	JSON projectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdrive]
type projectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding ProjectDeploymentConfigsPreviewKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectDeploymentConfigsPreviewKVNamespacesJSON      `json:"-"`
}

// projectDeploymentConfigsPreviewKVNamespacesJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewKVNamespaces]
type projectDeploymentConfigsPreviewKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                   `json:"namespace_id"`
	JSON        projectDeploymentConfigsPreviewKVNamespacesKVBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewKVNamespacesKVBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewKVNamespacesKVBinding]
type projectDeploymentConfigsPreviewKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectDeploymentConfigsPreviewMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectDeploymentConfigsPreviewMTLSCertificatesJSON `json:"-"`
}

// projectDeploymentConfigsPreviewMTLSCertificatesJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewMTLSCertificates]
type projectDeploymentConfigsPreviewMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID string                                                  `json:"certificate_id"`
	JSON          projectDeploymentConfigsPreviewMTLSCertificatesMTLSJSON `json:"-"`
}

// projectDeploymentConfigsPreviewMTLSCertificatesMTLSJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewMTLSCertificatesMTLS]
type projectDeploymentConfigsPreviewMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                       `json:"mode"`
	JSON projectDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectDeploymentConfigsPreviewPlacementJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewPlacement]
type projectDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// projectDeploymentConfigsPreviewQueueProducersJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewQueueProducers]
type projectDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                `json:"name"`
	JSON projectDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON contains
// the JSON metadata for the struct
// [ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type projectDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding ProjectDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// projectDeploymentConfigsPreviewR2BucketsJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewR2Buckets]
type projectDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                `json:"name"`
	JSON projectDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewR2BucketsR2Binding]
type projectDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding ProjectDeploymentConfigsPreviewServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectDeploymentConfigsPreviewServicesJSON           `json:"-"`
}

// projectDeploymentConfigsPreviewServicesJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewServices]
type projectDeploymentConfigsPreviewServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                    `json:"service"`
	JSON    projectDeploymentConfigsPreviewServicesServiceBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewServicesServiceBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewServicesServiceBinding]
type projectDeploymentConfigsPreviewServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectDeploymentConfigsPreviewVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectDeploymentConfigsPreviewVectorizeBindingsJSON      `json:"-"`
}

// projectDeploymentConfigsPreviewVectorizeBindingsJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewVectorizeBindings]
type projectDeploymentConfigsPreviewVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName string                                                        `json:"index_name"`
	JSON      projectDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsPreviewVectorizeBindingsVectorize]
type projectDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectDeploymentConfigsProductionBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectDeploymentConfigsProductionHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectDeploymentConfigsProductionMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectDeploymentConfigsProductionServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectDeploymentConfigsProductionVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectDeploymentConfigsProductionJSON              `json:"-"`
}

// projectDeploymentConfigsProductionJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigsProduction]
type projectDeploymentConfigsProductionJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	Browsers                apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	HyperdriveBindings      apijson.Field
	KVNamespaces            apijson.Field
	MTLSCertificates        apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	Services                apijson.Field
	VectorizeBindings       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding ProjectDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// projectDeploymentConfigsProductionAIBindingsJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionAIBindings]
type projectDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID interface{}                                               `json:"project_id"`
	JSON      projectDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionAIBindingsAIBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionAIBindingsAIBinding]
type projectDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectDeploymentConfigsProductionAnalyticsEngineDatasetsJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionAnalyticsEngineDatasets]
type projectDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                              `json:"dataset"`
	JSON    projectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser interface{}                                    `json:"BROWSER"`
	JSON    projectDeploymentConfigsProductionBrowsersJSON `json:"-"`
}

// projectDeploymentConfigsProductionBrowsersJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionBrowsers]
type projectDeploymentConfigsProductionBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding ProjectDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// projectDeploymentConfigsProductionD1DatabasesJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionD1Databases]
type projectDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                     `json:"id"`
	JSON projectDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionD1DatabasesD1BindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionD1DatabasesD1Binding]
type projectDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// projectDeploymentConfigsProductionDurableObjectNamespacesJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionDurableObjectNamespaces]
type projectDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                 `json:"namespace_id"`
	JSON        projectDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON contains
// the JSON metadata for the struct
// [ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type projectDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// projectDeploymentConfigsProductionEnvVarsJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionEnvVars]
type projectDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                           `json:"value"`
	JSON  projectDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type projectDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectDeploymentConfigsProductionHyperdriveBindingsJSON       `json:"-"`
}

// projectDeploymentConfigsProductionHyperdriveBindingsJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionHyperdriveBindings]
type projectDeploymentConfigsProductionHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID   string                                                             `json:"id"`
	JSON projectDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdrive]
type projectDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding ProjectDeploymentConfigsProductionKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectDeploymentConfigsProductionKVNamespacesJSON      `json:"-"`
}

// projectDeploymentConfigsProductionKVNamespacesJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsProductionKVNamespaces]
type projectDeploymentConfigsProductionKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                      `json:"namespace_id"`
	JSON        projectDeploymentConfigsProductionKVNamespacesKVBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionKVNamespacesKVBindingJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionKVNamespacesKVBinding]
type projectDeploymentConfigsProductionKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectDeploymentConfigsProductionMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectDeploymentConfigsProductionMTLSCertificatesJSON `json:"-"`
}

// projectDeploymentConfigsProductionMTLSCertificatesJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionMTLSCertificates]
type projectDeploymentConfigsProductionMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID string                                                     `json:"certificate_id"`
	JSON          projectDeploymentConfigsProductionMTLSCertificatesMTLSJSON `json:"-"`
}

// projectDeploymentConfigsProductionMTLSCertificatesMTLSJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionMTLSCertificatesMTLS]
type projectDeploymentConfigsProductionMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                          `json:"mode"`
	JSON projectDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectDeploymentConfigsProductionPlacementJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionPlacement]
type projectDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// projectDeploymentConfigsProductionQueueProducersJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsProductionQueueProducers]
type projectDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                   `json:"name"`
	JSON projectDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type projectDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding ProjectDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// projectDeploymentConfigsProductionR2BucketsJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionR2Buckets]
type projectDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                   `json:"name"`
	JSON projectDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionR2BucketsR2BindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionR2BucketsR2Binding]
type projectDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding ProjectDeploymentConfigsProductionServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectDeploymentConfigsProductionServicesJSON           `json:"-"`
}

// projectDeploymentConfigsProductionServicesJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionServices]
type projectDeploymentConfigsProductionServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                       `json:"service"`
	JSON    projectDeploymentConfigsProductionServicesServiceBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionServicesServiceBindingJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionServicesServiceBinding]
type projectDeploymentConfigsProductionServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectDeploymentConfigsProductionVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectDeploymentConfigsProductionVectorizeBindingsJSON      `json:"-"`
}

// projectDeploymentConfigsProductionVectorizeBindingsJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionVectorizeBindings]
type projectDeploymentConfigsProductionVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName string                                                           `json:"index_name"`
	JSON      projectDeploymentConfigsProductionVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectDeploymentConfigsProductionVectorizeBindingsVectorizeJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionVectorizeBindingsVectorize]
type projectDeploymentConfigsProductionVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

type ProjectParam struct {
	// Configs for the project build process.
	BuildConfig         param.Field[ProjectBuildConfigParam] `json:"build_config"`
	CanonicalDeployment param.Field[DeploymentParam]         `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectDeploymentConfigsParam] `json:"deployment_configs"`
	LatestDeployment  param.Field[DeploymentParam]               `json:"latest_deployment"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
}

func (r ProjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type ProjectBuildConfigParam struct {
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

func (r ProjectBuildConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectDeploymentConfigsParam struct {
	// Configs for preview deploys.
	Preview param.Field[ProjectDeploymentConfigsPreviewParam] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[ProjectDeploymentConfigsProductionParam] `json:"production"`
}

func (r ProjectDeploymentConfigsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type ProjectDeploymentConfigsPreviewParam struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectDeploymentConfigsPreviewAIBindingsParam] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsParam] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[ProjectDeploymentConfigsPreviewBrowsersParam] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[ProjectDeploymentConfigsPreviewD1DatabasesParam] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[ProjectDeploymentConfigsPreviewDurableObjectNamespacesParam] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[ProjectDeploymentConfigsPreviewEnvVarsParam] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[ProjectDeploymentConfigsPreviewHyperdriveBindingsParam] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectDeploymentConfigsPreviewKVNamespacesParam] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[ProjectDeploymentConfigsPreviewMTLSCertificatesParam] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectDeploymentConfigsPreviewPlacementParam] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectDeploymentConfigsPreviewQueueProducersParam] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectDeploymentConfigsPreviewR2BucketsParam] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[ProjectDeploymentConfigsPreviewServicesParam] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[ProjectDeploymentConfigsPreviewVectorizeBindingsParam] `json:"vectorize_bindings"`
}

func (r ProjectDeploymentConfigsPreviewParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewAIBindingsParam struct {
	// AI binding.
	AIBinding param.Field[ProjectDeploymentConfigsPreviewAIBindingsAIBindingParam] `json:"AI_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewAIBindingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectDeploymentConfigsPreviewAIBindingsAIBindingParam struct {
	ProjectID param.Field[interface{}] `json:"project_id"`
}

func (r ProjectDeploymentConfigsPreviewAIBindingsAIBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsParam struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingParam] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingParam struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewBrowsersParam struct {
	// Browser binding.
	Browser param.Field[interface{}] `json:"BROWSER"`
}

func (r ProjectDeploymentConfigsPreviewBrowsersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type ProjectDeploymentConfigsPreviewD1DatabasesParam struct {
	// D1 binding.
	D1Binding param.Field[ProjectDeploymentConfigsPreviewD1DatabasesD1BindingParam] `json:"D1_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewD1DatabasesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectDeploymentConfigsPreviewD1DatabasesD1BindingParam struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsPreviewD1DatabasesD1BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectDeploymentConfigsPreviewDurableObjectNamespacesParam struct {
	// Durabble Object binding.
	DoBinding param.Field[ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingParam] `json:"DO_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewDurableObjectNamespacesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingParam struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type ProjectDeploymentConfigsPreviewEnvVarsParam struct {
	// Environment variable.
	EnvironmentVariable param.Field[ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableParam] `json:"ENVIRONMENT_VARIABLE"`
}

func (r ProjectDeploymentConfigsPreviewEnvVarsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableParam struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewHyperdriveBindingsParam struct {
	// Hyperdrive binding.
	Hyperdrive param.Field[ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveParam] `json:"HYPERDRIVE"`
}

func (r ProjectDeploymentConfigsPreviewHyperdriveBindingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive binding.
type ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveParam struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespaces used for Pages Functions.
type ProjectDeploymentConfigsPreviewKVNamespacesParam struct {
	// KV binding.
	KVBinding param.Field[ProjectDeploymentConfigsPreviewKVNamespacesKVBindingParam] `json:"KV_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewKVNamespacesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type ProjectDeploymentConfigsPreviewKVNamespacesKVBindingParam struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsPreviewKVNamespacesKVBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewMTLSCertificatesParam struct {
	// mTLS binding.
	MTLS param.Field[ProjectDeploymentConfigsPreviewMTLSCertificatesMTLSParam] `json:"MTLS"`
}

func (r ProjectDeploymentConfigsPreviewMTLSCertificatesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectDeploymentConfigsPreviewMTLSCertificatesMTLSParam struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectDeploymentConfigsPreviewMTLSCertificatesMTLSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsPreviewPlacementParam struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectDeploymentConfigsPreviewPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewQueueProducersParam struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingParam] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewQueueProducersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingParam struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type ProjectDeploymentConfigsPreviewR2BucketsParam struct {
	// R2 binding.
	R2Binding param.Field[ProjectDeploymentConfigsPreviewR2BucketsR2BindingParam] `json:"R2_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewR2BucketsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectDeploymentConfigsPreviewR2BucketsR2BindingParam struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsPreviewR2BucketsR2BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type ProjectDeploymentConfigsPreviewServicesParam struct {
	// Service binding.
	ServiceBinding param.Field[ProjectDeploymentConfigsPreviewServicesServiceBindingParam] `json:"SERVICE_BINDING"`
}

func (r ProjectDeploymentConfigsPreviewServicesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectDeploymentConfigsPreviewServicesServiceBindingParam struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectDeploymentConfigsPreviewServicesServiceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize bindings used for Pages Functions.
type ProjectDeploymentConfigsPreviewVectorizeBindingsParam struct {
	// Vectorize binding.
	Vectorize param.Field[ProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeParam] `json:"VECTORIZE"`
}

func (r ProjectDeploymentConfigsPreviewVectorizeBindingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeParam struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectDeploymentConfigsProductionParam struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectDeploymentConfigsProductionAIBindingsParam] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsParam] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[ProjectDeploymentConfigsProductionBrowsersParam] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[ProjectDeploymentConfigsProductionD1DatabasesParam] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[ProjectDeploymentConfigsProductionDurableObjectNamespacesParam] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[ProjectDeploymentConfigsProductionEnvVarsParam] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[ProjectDeploymentConfigsProductionHyperdriveBindingsParam] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectDeploymentConfigsProductionKVNamespacesParam] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[ProjectDeploymentConfigsProductionMTLSCertificatesParam] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectDeploymentConfigsProductionPlacementParam] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectDeploymentConfigsProductionQueueProducersParam] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectDeploymentConfigsProductionR2BucketsParam] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[ProjectDeploymentConfigsProductionServicesParam] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[ProjectDeploymentConfigsProductionVectorizeBindingsParam] `json:"vectorize_bindings"`
}

func (r ProjectDeploymentConfigsProductionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionAIBindingsParam struct {
	// AI binding.
	AIBinding param.Field[ProjectDeploymentConfigsProductionAIBindingsAIBindingParam] `json:"AI_BINDING"`
}

func (r ProjectDeploymentConfigsProductionAIBindingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectDeploymentConfigsProductionAIBindingsAIBindingParam struct {
	ProjectID param.Field[interface{}] `json:"project_id"`
}

func (r ProjectDeploymentConfigsProductionAIBindingsAIBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsParam struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingParam] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingParam struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionBrowsersParam struct {
	// Browser binding.
	Browser param.Field[interface{}] `json:"BROWSER"`
}

func (r ProjectDeploymentConfigsProductionBrowsersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type ProjectDeploymentConfigsProductionD1DatabasesParam struct {
	// D1 binding.
	D1Binding param.Field[ProjectDeploymentConfigsProductionD1DatabasesD1BindingParam] `json:"D1_BINDING"`
}

func (r ProjectDeploymentConfigsProductionD1DatabasesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectDeploymentConfigsProductionD1DatabasesD1BindingParam struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsProductionD1DatabasesD1BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectDeploymentConfigsProductionDurableObjectNamespacesParam struct {
	// Durabble Object binding.
	DoBinding param.Field[ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingParam] `json:"DO_BINDING"`
}

func (r ProjectDeploymentConfigsProductionDurableObjectNamespacesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingParam struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type ProjectDeploymentConfigsProductionEnvVarsParam struct {
	// Environment variable.
	EnvironmentVariable param.Field[ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableParam] `json:"ENVIRONMENT_VARIABLE"`
}

func (r ProjectDeploymentConfigsProductionEnvVarsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableParam struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionHyperdriveBindingsParam struct {
	// Hyperdrive binding.
	Hyperdrive param.Field[ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveParam] `json:"HYPERDRIVE"`
}

func (r ProjectDeploymentConfigsProductionHyperdriveBindingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive binding.
type ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveParam struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespaces used for Pages Functions.
type ProjectDeploymentConfigsProductionKVNamespacesParam struct {
	// KV binding.
	KVBinding param.Field[ProjectDeploymentConfigsProductionKVNamespacesKVBindingParam] `json:"KV_BINDING"`
}

func (r ProjectDeploymentConfigsProductionKVNamespacesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type ProjectDeploymentConfigsProductionKVNamespacesKVBindingParam struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsProductionKVNamespacesKVBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionMTLSCertificatesParam struct {
	// mTLS binding.
	MTLS param.Field[ProjectDeploymentConfigsProductionMTLSCertificatesMTLSParam] `json:"MTLS"`
}

func (r ProjectDeploymentConfigsProductionMTLSCertificatesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectDeploymentConfigsProductionMTLSCertificatesMTLSParam struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectDeploymentConfigsProductionMTLSCertificatesMTLSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsProductionPlacementParam struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectDeploymentConfigsProductionPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionQueueProducersParam struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[ProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingParam] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r ProjectDeploymentConfigsProductionQueueProducersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingParam struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type ProjectDeploymentConfigsProductionR2BucketsParam struct {
	// R2 binding.
	R2Binding param.Field[ProjectDeploymentConfigsProductionR2BucketsR2BindingParam] `json:"R2_BINDING"`
}

func (r ProjectDeploymentConfigsProductionR2BucketsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectDeploymentConfigsProductionR2BucketsR2BindingParam struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsProductionR2BucketsR2BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type ProjectDeploymentConfigsProductionServicesParam struct {
	// Service binding.
	ServiceBinding param.Field[ProjectDeploymentConfigsProductionServicesServiceBindingParam] `json:"SERVICE_BINDING"`
}

func (r ProjectDeploymentConfigsProductionServicesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectDeploymentConfigsProductionServicesServiceBindingParam struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectDeploymentConfigsProductionServicesServiceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize bindings used for Pages Functions.
type ProjectDeploymentConfigsProductionVectorizeBindingsParam struct {
	// Vectorize binding.
	Vectorize param.Field[ProjectDeploymentConfigsProductionVectorizeBindingsVectorizeParam] `json:"VECTORIZE"`
}

func (r ProjectDeploymentConfigsProductionVectorizeBindingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectDeploymentConfigsProductionVectorizeBindingsVectorizeParam struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectDeploymentConfigsProductionVectorizeBindingsVectorizeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type Stage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string    `json:"status"`
	JSON   stageJSON `json:"-"`
}

// stageJSON contains the JSON metadata for the struct [Stage]
type stageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Stage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stageJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type StageParam struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r StageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [pages.ProjectNewResponseUnknown],
// [pages.ProjectNewResponseArray] or [shared.UnionString].
type ProjectNewResponseUnion interface {
	ImplementsPagesProjectNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectNewResponseArray []interface{}

func (r ProjectNewResponseArray) ImplementsPagesProjectNewResponseUnion() {}

type ProjectDeleteResponse = interface{}

// Union satisfied by [pages.ProjectEditResponseUnknown],
// [pages.ProjectEditResponseArray] or [shared.UnionString].
type ProjectEditResponseUnion interface {
	ImplementsPagesProjectEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectEditResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectEditResponseArray []interface{}

func (r ProjectEditResponseArray) ImplementsPagesProjectEditResponseUnion() {}

type ProjectPurgeBuildCacheResponse = interface{}

type ProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Project   ProjectParam        `json:"project,required"`
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Project)
}

type ProjectNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   ProjectNewResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success ProjectNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectNewResponseEnvelopeJSON    `json:"-"`
}

// projectNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelope]
type projectNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectNewResponseEnvelopeSuccess bool

const (
	ProjectNewResponseEnvelopeSuccessTrue ProjectNewResponseEnvelopeSuccess = true
)

func (r ProjectNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ProjectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   ProjectEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success ProjectEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectEditResponseEnvelopeJSON    `json:"-"`
}

// projectEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectEditResponseEnvelope]
type projectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectEditResponseEnvelopeSuccess bool

const (
	ProjectEditResponseEnvelopeSuccessTrue ProjectEditResponseEnvelopeSuccess = true
)

func (r ProjectEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Project               `json:"result,required"`
	// Whether the API call was successful
	Success ProjectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectGetResponseEnvelopeJSON    `json:"-"`
}

// projectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelope]
type projectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectGetResponseEnvelopeSuccess bool

const (
	ProjectGetResponseEnvelopeSuccessTrue ProjectGetResponseEnvelopeSuccess = true
)

func (r ProjectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectPurgeBuildCacheParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
