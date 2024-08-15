// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	var env ProjectNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
	var env ProjectDeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *ProjectService) Edit(ctx context.Context, projectName string, params ProjectEditParams, opts ...option.RequestOption) (res *ProjectEditResponse, err error) {
	var env ProjectEditResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	var env ProjectGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	var env ProjectPurgeBuildCacheResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Deployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,nullable"`
	// Configs for the project build process.
	BuildConfig DeploymentBuildConfig `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string           `json:"short_id"`
	Source  DeploymentSource `json:"source"`
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

// Configs for the project build process.
type DeploymentBuildConfig struct {
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
	WebAnalyticsToken string                    `json:"web_analytics_token,nullable"`
	JSON              deploymentBuildConfigJSON `json:"-"`
}

// deploymentBuildConfigJSON contains the JSON metadata for the struct
// [DeploymentBuildConfig]
type deploymentBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentBuildConfigJSON) RawJSON() string {
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

type DeploymentSource struct {
	Config DeploymentSourceConfig `json:"config"`
	Type   string                 `json:"type"`
	JSON   deploymentSourceJSON   `json:"-"`
}

// deploymentSourceJSON contains the JSON metadata for the struct
// [DeploymentSource]
type deploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentSourceJSON) RawJSON() string {
	return r.raw
}

type DeploymentSourceConfig struct {
	DeploymentsEnabled           bool                                           `json:"deployments_enabled"`
	Owner                        string                                         `json:"owner"`
	PathExcludes                 []string                                       `json:"path_excludes"`
	PathIncludes                 []string                                       `json:"path_includes"`
	PrCommentsEnabled            bool                                           `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                                       `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                                       `json:"preview_branch_includes"`
	PreviewDeploymentSetting     DeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                                         `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                           `json:"production_deployments_enabled"`
	RepoName                     string                                         `json:"repo_name"`
	JSON                         deploymentSourceConfigJSON                     `json:"-"`
}

// deploymentSourceConfigJSON contains the JSON metadata for the struct
// [DeploymentSourceConfig]
type deploymentSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

type DeploymentSourceConfigPreviewDeploymentSetting string

const (
	DeploymentSourceConfigPreviewDeploymentSettingAll    DeploymentSourceConfigPreviewDeploymentSetting = "all"
	DeploymentSourceConfigPreviewDeploymentSettingNone   DeploymentSourceConfigPreviewDeploymentSetting = "none"
	DeploymentSourceConfigPreviewDeploymentSettingCustom DeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r DeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case DeploymentSourceConfigPreviewDeploymentSettingAll, DeploymentSourceConfigPreviewDeploymentSettingNone, DeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
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

type ProjectNewResponse struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig ProjectNewResponseBuildConfig `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment Deployment `json:"canonical_deployment,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectNewResponseDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Most recent deployment to the repo.
	LatestDeployment Deployment `json:"latest_deployment,nullable"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string                   `json:"production_branch"`
	Source           ProjectNewResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                 `json:"subdomain"`
	JSON      projectNewResponseJSON `json:"-"`
}

// projectNewResponseJSON contains the JSON metadata for the struct
// [ProjectNewResponse]
type projectNewResponseJSON struct {
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

func (r *ProjectNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectNewResponseBuildConfig struct {
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
	WebAnalyticsToken string                            `json:"web_analytics_token,nullable"`
	JSON              projectNewResponseBuildConfigJSON `json:"-"`
}

// projectNewResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectNewResponseBuildConfig]
type projectNewResponseBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectNewResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectNewResponseDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production ProjectNewResponseDeploymentConfigsProduction `json:"production"`
	JSON       projectNewResponseDeploymentConfigsJSON       `json:"-"`
}

// projectNewResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectNewResponseDeploymentConfigs]
type projectNewResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectNewResponseDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectNewResponseDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectNewResponseDeploymentConfigsPreviewBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectNewResponseDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectNewResponseDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectNewResponseDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectNewResponseDeploymentConfigsPreviewMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectNewResponseDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectNewResponseDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectNewResponseDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectNewResponseDeploymentConfigsPreviewServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectNewResponseDeploymentConfigsPreviewVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectNewResponseDeploymentConfigsPreviewJSON              `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectNewResponseDeploymentConfigsPreview]
type projectNewResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectNewResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding ProjectNewResponseDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectNewResponseDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewAIBindingsJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewAIBindings]
type projectNewResponseDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectNewResponseDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID string                                                            `json:"project_id"`
	JSON      projectNewResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewAIBindingsAIBinding]
type projectNewResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasets]
type projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                      `json:"dataset"`
	JSON    projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser interface{}                                            `json:"BROWSER"`
	JSON    projectNewResponseDeploymentConfigsPreviewBrowsersJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewBrowsersJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewBrowsers]
type projectNewResponseDeploymentConfigsPreviewBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding ProjectNewResponseDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectNewResponseDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewD1DatabasesJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewD1Databases]
type projectNewResponseDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectNewResponseDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                             `json:"id"`
	JSON projectNewResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewD1DatabasesD1Binding]
type projectNewResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespaces]
type projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                         `json:"namespace_id"`
	JSON        projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectNewResponseDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectNewResponseDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewEnvVarsJSON contains the JSON metadata
// for the struct [ProjectNewResponseDeploymentConfigsPreviewEnvVars]
type projectNewResponseDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                   `json:"value"`
	JSON  projectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type projectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, ProjectNewResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsJSON       `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindings]
type projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID   string                                                                     `json:"id"`
	JSON projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive]
type projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding ProjectNewResponseDeploymentConfigsPreviewKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectNewResponseDeploymentConfigsPreviewKVNamespacesJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewKVNamespacesJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewKVNamespaces]
type projectNewResponseDeploymentConfigsPreviewKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectNewResponseDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                              `json:"namespace_id"`
	JSON        projectNewResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewKVNamespacesKVBinding]
type projectNewResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectNewResponseDeploymentConfigsPreviewMTLSCertificatesJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewMTLSCertificatesJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewMTLSCertificates]
type projectNewResponseDeploymentConfigsPreviewMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID string                                                             `json:"certificate_id"`
	JSON          projectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLS]
type projectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                  `json:"mode"`
	JSON projectNewResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewPlacement]
type projectNewResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectNewResponseDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewQueueProducersJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewQueueProducers]
type projectNewResponseDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                           `json:"name"`
	JSON projectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type projectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding ProjectNewResponseDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectNewResponseDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewR2BucketsJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewR2Buckets]
type projectNewResponseDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectNewResponseDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                           `json:"name"`
	JSON projectNewResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewR2BucketsR2Binding]
type projectNewResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding ProjectNewResponseDeploymentConfigsPreviewServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectNewResponseDeploymentConfigsPreviewServicesJSON           `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewServicesJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewServices]
type projectNewResponseDeploymentConfigsPreviewServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectNewResponseDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                               `json:"service"`
	JSON    projectNewResponseDeploymentConfigsPreviewServicesServiceBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewServicesServiceBindingJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewServicesServiceBinding]
type projectNewResponseDeploymentConfigsPreviewServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectNewResponseDeploymentConfigsPreviewVectorizeBindingsJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewVectorizeBindingsJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewVectorizeBindings]
type projectNewResponseDeploymentConfigsPreviewVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName string                                                                   `json:"index_name"`
	JSON      projectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorize]
type projectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectNewResponseDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectNewResponseDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectNewResponseDeploymentConfigsProductionBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectNewResponseDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectNewResponseDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectNewResponseDeploymentConfigsProductionHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectNewResponseDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectNewResponseDeploymentConfigsProductionMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectNewResponseDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectNewResponseDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectNewResponseDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectNewResponseDeploymentConfigsProductionServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectNewResponseDeploymentConfigsProductionVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectNewResponseDeploymentConfigsProductionJSON              `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionJSON contains the JSON metadata for
// the struct [ProjectNewResponseDeploymentConfigsProduction]
type projectNewResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectNewResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding ProjectNewResponseDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectNewResponseDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionAIBindingsJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionAIBindings]
type projectNewResponseDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectNewResponseDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID string                                                               `json:"project_id"`
	JSON      projectNewResponseDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionAIBindingsAIBindingJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionAIBindingsAIBinding]
type projectNewResponseDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasets]
type projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                         `json:"dataset"`
	JSON    projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser interface{}                                               `json:"BROWSER"`
	JSON    projectNewResponseDeploymentConfigsProductionBrowsersJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionBrowsersJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionBrowsers]
type projectNewResponseDeploymentConfigsProductionBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding ProjectNewResponseDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectNewResponseDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionD1DatabasesJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionD1Databases]
type projectNewResponseDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectNewResponseDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                `json:"id"`
	JSON projectNewResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionD1DatabasesD1Binding]
type projectNewResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespaces]
type projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                            `json:"namespace_id"`
	JSON        projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectNewResponseDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectNewResponseDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionEnvVarsJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionEnvVars]
type projectNewResponseDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                      `json:"value"`
	JSON  projectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type projectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, ProjectNewResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectNewResponseDeploymentConfigsProductionHyperdriveBindingsJSON       `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionHyperdriveBindingsJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionHyperdriveBindings]
type projectNewResponseDeploymentConfigsProductionHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID   string                                                                        `json:"id"`
	JSON projectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive]
type projectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding ProjectNewResponseDeploymentConfigsProductionKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectNewResponseDeploymentConfigsProductionKVNamespacesJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionKVNamespacesJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionKVNamespaces]
type projectNewResponseDeploymentConfigsProductionKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectNewResponseDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                 `json:"namespace_id"`
	JSON        projectNewResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionKVNamespacesKVBinding]
type projectNewResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectNewResponseDeploymentConfigsProductionMTLSCertificatesJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionMTLSCertificatesJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionMTLSCertificates]
type projectNewResponseDeploymentConfigsProductionMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID string                                                                `json:"certificate_id"`
	JSON          projectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLS]
type projectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                     `json:"mode"`
	JSON projectNewResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionPlacement]
type projectNewResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectNewResponseDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionQueueProducersJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionQueueProducers]
type projectNewResponseDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                              `json:"name"`
	JSON projectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type projectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding ProjectNewResponseDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectNewResponseDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionR2BucketsJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionR2Buckets]
type projectNewResponseDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectNewResponseDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                              `json:"name"`
	JSON projectNewResponseDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionR2BucketsR2Binding]
type projectNewResponseDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding ProjectNewResponseDeploymentConfigsProductionServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectNewResponseDeploymentConfigsProductionServicesJSON           `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionServicesJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionServices]
type projectNewResponseDeploymentConfigsProductionServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectNewResponseDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                  `json:"service"`
	JSON    projectNewResponseDeploymentConfigsProductionServicesServiceBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionServicesServiceBindingJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionServicesServiceBinding]
type projectNewResponseDeploymentConfigsProductionServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectNewResponseDeploymentConfigsProductionVectorizeBindingsJSON      `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionVectorizeBindingsJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionVectorizeBindings]
type projectNewResponseDeploymentConfigsProductionVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName string                                                                      `json:"index_name"`
	JSON      projectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorize]
type projectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseSource struct {
	Config ProjectNewResponseSourceConfig `json:"config"`
	Type   string                         `json:"type"`
	JSON   projectNewResponseSourceJSON   `json:"-"`
}

// projectNewResponseSourceJSON contains the JSON metadata for the struct
// [ProjectNewResponseSource]
type projectNewResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseSourceConfig struct {
	DeploymentsEnabled           bool                                                   `json:"deployments_enabled"`
	Owner                        string                                                 `json:"owner"`
	PathExcludes                 []string                                               `json:"path_excludes"`
	PathIncludes                 []string                                               `json:"path_includes"`
	PrCommentsEnabled            bool                                                   `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                                               `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                                               `json:"preview_branch_includes"`
	PreviewDeploymentSetting     ProjectNewResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                                                 `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                                   `json:"production_deployments_enabled"`
	RepoName                     string                                                 `json:"repo_name"`
	JSON                         projectNewResponseSourceConfigJSON                     `json:"-"`
}

// projectNewResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectNewResponseSourceConfig]
type projectNewResponseSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectNewResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectNewResponseSourceConfigPreviewDeploymentSettingAll    ProjectNewResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectNewResponseSourceConfigPreviewDeploymentSettingNone   ProjectNewResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectNewResponseSourceConfigPreviewDeploymentSettingCustom ProjectNewResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectNewResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectNewResponseSourceConfigPreviewDeploymentSettingAll, ProjectNewResponseSourceConfigPreviewDeploymentSettingNone, ProjectNewResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

type ProjectDeleteResponse = interface{}

type ProjectEditResponse struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig ProjectEditResponseBuildConfig `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment Deployment `json:"canonical_deployment,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectEditResponseDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Most recent deployment to the repo.
	LatestDeployment Deployment `json:"latest_deployment,nullable"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string                    `json:"production_branch"`
	Source           ProjectEditResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                  `json:"subdomain"`
	JSON      projectEditResponseJSON `json:"-"`
}

// projectEditResponseJSON contains the JSON metadata for the struct
// [ProjectEditResponse]
type projectEditResponseJSON struct {
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

func (r *ProjectEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectEditResponseBuildConfig struct {
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
	WebAnalyticsToken string                             `json:"web_analytics_token,nullable"`
	JSON              projectEditResponseBuildConfigJSON `json:"-"`
}

// projectEditResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectEditResponseBuildConfig]
type projectEditResponseBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectEditResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectEditResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectEditResponseDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production ProjectEditResponseDeploymentConfigsProduction `json:"production"`
	JSON       projectEditResponseDeploymentConfigsJSON       `json:"-"`
}

// projectEditResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectEditResponseDeploymentConfigs]
type projectEditResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectEditResponseDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectEditResponseDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectEditResponseDeploymentConfigsPreviewBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectEditResponseDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectEditResponseDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectEditResponseDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectEditResponseDeploymentConfigsPreviewMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectEditResponseDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectEditResponseDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectEditResponseDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectEditResponseDeploymentConfigsPreviewServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectEditResponseDeploymentConfigsPreviewVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectEditResponseDeploymentConfigsPreviewJSON              `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectEditResponseDeploymentConfigsPreview]
type projectEditResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectEditResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding ProjectEditResponseDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectEditResponseDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewAIBindingsJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewAIBindings]
type projectEditResponseDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectEditResponseDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID string                                                             `json:"project_id"`
	JSON      projectEditResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewAIBindingsAIBinding]
type projectEditResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasets]
type projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                       `json:"dataset"`
	JSON    projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser interface{}                                             `json:"BROWSER"`
	JSON    projectEditResponseDeploymentConfigsPreviewBrowsersJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewBrowsersJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewBrowsers]
type projectEditResponseDeploymentConfigsPreviewBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding ProjectEditResponseDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectEditResponseDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewD1DatabasesJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewD1Databases]
type projectEditResponseDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectEditResponseDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                              `json:"id"`
	JSON projectEditResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewD1DatabasesD1Binding]
type projectEditResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespaces]
type projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                          `json:"namespace_id"`
	JSON        projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectEditResponseDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectEditResponseDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewEnvVarsJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewEnvVars]
type projectEditResponseDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                    `json:"value"`
	JSON  projectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type projectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, ProjectEditResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsJSON       `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindings]
type projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID   string                                                                      `json:"id"`
	JSON projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive]
type projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding ProjectEditResponseDeploymentConfigsPreviewKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectEditResponseDeploymentConfigsPreviewKVNamespacesJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewKVNamespacesJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewKVNamespaces]
type projectEditResponseDeploymentConfigsPreviewKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectEditResponseDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                               `json:"namespace_id"`
	JSON        projectEditResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewKVNamespacesKVBinding]
type projectEditResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectEditResponseDeploymentConfigsPreviewMTLSCertificatesJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewMTLSCertificatesJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewMTLSCertificates]
type projectEditResponseDeploymentConfigsPreviewMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID string                                                              `json:"certificate_id"`
	JSON          projectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLS]
type projectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                   `json:"mode"`
	JSON projectEditResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewPlacement]
type projectEditResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectEditResponseDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewQueueProducersJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewQueueProducers]
type projectEditResponseDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                            `json:"name"`
	JSON projectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type projectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding ProjectEditResponseDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectEditResponseDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewR2BucketsJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewR2Buckets]
type projectEditResponseDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectEditResponseDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                            `json:"name"`
	JSON projectEditResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewR2BucketsR2Binding]
type projectEditResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding ProjectEditResponseDeploymentConfigsPreviewServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectEditResponseDeploymentConfigsPreviewServicesJSON           `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewServicesJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewServices]
type projectEditResponseDeploymentConfigsPreviewServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectEditResponseDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                `json:"service"`
	JSON    projectEditResponseDeploymentConfigsPreviewServicesServiceBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewServicesServiceBindingJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewServicesServiceBinding]
type projectEditResponseDeploymentConfigsPreviewServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectEditResponseDeploymentConfigsPreviewVectorizeBindingsJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewVectorizeBindingsJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewVectorizeBindings]
type projectEditResponseDeploymentConfigsPreviewVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName string                                                                    `json:"index_name"`
	JSON      projectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorize]
type projectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectEditResponseDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectEditResponseDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectEditResponseDeploymentConfigsProductionBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectEditResponseDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectEditResponseDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectEditResponseDeploymentConfigsProductionHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectEditResponseDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectEditResponseDeploymentConfigsProductionMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectEditResponseDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectEditResponseDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectEditResponseDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectEditResponseDeploymentConfigsProductionServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectEditResponseDeploymentConfigsProductionVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectEditResponseDeploymentConfigsProductionJSON              `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionJSON contains the JSON metadata
// for the struct [ProjectEditResponseDeploymentConfigsProduction]
type projectEditResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectEditResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding ProjectEditResponseDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectEditResponseDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionAIBindingsJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionAIBindings]
type projectEditResponseDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectEditResponseDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID string                                                                `json:"project_id"`
	JSON      projectEditResponseDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionAIBindingsAIBindingJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionAIBindingsAIBinding]
type projectEditResponseDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasets]
type projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                          `json:"dataset"`
	JSON    projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser interface{}                                                `json:"BROWSER"`
	JSON    projectEditResponseDeploymentConfigsProductionBrowsersJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionBrowsersJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionBrowsers]
type projectEditResponseDeploymentConfigsProductionBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding ProjectEditResponseDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectEditResponseDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionD1DatabasesJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionD1Databases]
type projectEditResponseDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectEditResponseDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                 `json:"id"`
	JSON projectEditResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionD1DatabasesD1Binding]
type projectEditResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespaces]
type projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                             `json:"namespace_id"`
	JSON        projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectEditResponseDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectEditResponseDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionEnvVarsJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionEnvVars]
type projectEditResponseDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                       `json:"value"`
	JSON  projectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type projectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, ProjectEditResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectEditResponseDeploymentConfigsProductionHyperdriveBindingsJSON       `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionHyperdriveBindingsJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionHyperdriveBindings]
type projectEditResponseDeploymentConfigsProductionHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID   string                                                                         `json:"id"`
	JSON projectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive]
type projectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding ProjectEditResponseDeploymentConfigsProductionKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectEditResponseDeploymentConfigsProductionKVNamespacesJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionKVNamespacesJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionKVNamespaces]
type projectEditResponseDeploymentConfigsProductionKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectEditResponseDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                  `json:"namespace_id"`
	JSON        projectEditResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionKVNamespacesKVBinding]
type projectEditResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectEditResponseDeploymentConfigsProductionMTLSCertificatesJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionMTLSCertificatesJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionMTLSCertificates]
type projectEditResponseDeploymentConfigsProductionMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID string                                                                 `json:"certificate_id"`
	JSON          projectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLS]
type projectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                      `json:"mode"`
	JSON projectEditResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionPlacement]
type projectEditResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectEditResponseDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionQueueProducersJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionQueueProducers]
type projectEditResponseDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                               `json:"name"`
	JSON projectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type projectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding ProjectEditResponseDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectEditResponseDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionR2BucketsJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionR2Buckets]
type projectEditResponseDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectEditResponseDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                               `json:"name"`
	JSON projectEditResponseDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionR2BucketsR2BindingJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionR2BucketsR2Binding]
type projectEditResponseDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding ProjectEditResponseDeploymentConfigsProductionServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectEditResponseDeploymentConfigsProductionServicesJSON           `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionServicesJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionServices]
type projectEditResponseDeploymentConfigsProductionServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectEditResponseDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                   `json:"service"`
	JSON    projectEditResponseDeploymentConfigsProductionServicesServiceBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionServicesServiceBindingJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionServicesServiceBinding]
type projectEditResponseDeploymentConfigsProductionServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectEditResponseDeploymentConfigsProductionVectorizeBindingsJSON      `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionVectorizeBindingsJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionVectorizeBindings]
type projectEditResponseDeploymentConfigsProductionVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName string                                                                       `json:"index_name"`
	JSON      projectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorize]
type projectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseSource struct {
	Config ProjectEditResponseSourceConfig `json:"config"`
	Type   string                          `json:"type"`
	JSON   projectEditResponseSourceJSON   `json:"-"`
}

// projectEditResponseSourceJSON contains the JSON metadata for the struct
// [ProjectEditResponseSource]
type projectEditResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseSourceConfig struct {
	DeploymentsEnabled           bool                                                    `json:"deployments_enabled"`
	Owner                        string                                                  `json:"owner"`
	PathExcludes                 []string                                                `json:"path_excludes"`
	PathIncludes                 []string                                                `json:"path_includes"`
	PrCommentsEnabled            bool                                                    `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                                                `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                                                `json:"preview_branch_includes"`
	PreviewDeploymentSetting     ProjectEditResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                                                  `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                                    `json:"production_deployments_enabled"`
	RepoName                     string                                                  `json:"repo_name"`
	JSON                         projectEditResponseSourceConfigJSON                     `json:"-"`
}

// projectEditResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectEditResponseSourceConfig]
type projectEditResponseSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectEditResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectEditResponseSourceConfigPreviewDeploymentSettingAll    ProjectEditResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectEditResponseSourceConfigPreviewDeploymentSettingNone   ProjectEditResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectEditResponseSourceConfigPreviewDeploymentSettingCustom ProjectEditResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectEditResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectEditResponseSourceConfigPreviewDeploymentSettingAll, ProjectEditResponseSourceConfigPreviewDeploymentSettingNone, ProjectEditResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

type ProjectGetResponse struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig ProjectGetResponseBuildConfig `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment Deployment `json:"canonical_deployment,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectGetResponseDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Most recent deployment to the repo.
	LatestDeployment Deployment `json:"latest_deployment,nullable"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string                   `json:"production_branch"`
	Source           ProjectGetResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                 `json:"subdomain"`
	JSON      projectGetResponseJSON `json:"-"`
}

// projectGetResponseJSON contains the JSON metadata for the struct
// [ProjectGetResponse]
type projectGetResponseJSON struct {
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

func (r *ProjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectGetResponseBuildConfig struct {
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
	WebAnalyticsToken string                            `json:"web_analytics_token,nullable"`
	JSON              projectGetResponseBuildConfigJSON `json:"-"`
}

// projectGetResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectGetResponseBuildConfig]
type projectGetResponseBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectGetResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectGetResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectGetResponseDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production ProjectGetResponseDeploymentConfigsProduction `json:"production"`
	JSON       projectGetResponseDeploymentConfigsJSON       `json:"-"`
}

// projectGetResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectGetResponseDeploymentConfigs]
type projectGetResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectGetResponseDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectGetResponseDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectGetResponseDeploymentConfigsPreviewBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectGetResponseDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectGetResponseDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectGetResponseDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectGetResponseDeploymentConfigsPreviewMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectGetResponseDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectGetResponseDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectGetResponseDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectGetResponseDeploymentConfigsPreviewServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectGetResponseDeploymentConfigsPreviewVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectGetResponseDeploymentConfigsPreviewJSON              `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectGetResponseDeploymentConfigsPreview]
type projectGetResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectGetResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding ProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectGetResponseDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewAIBindingsJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewAIBindings]
type projectGetResponseDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID string                                                            `json:"project_id"`
	JSON      projectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding]
type projectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets]
type projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                      `json:"dataset"`
	JSON    projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser interface{}                                            `json:"BROWSER"`
	JSON    projectGetResponseDeploymentConfigsPreviewBrowsersJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewBrowsersJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewBrowsers]
type projectGetResponseDeploymentConfigsPreviewBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding ProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectGetResponseDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewD1DatabasesJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewD1Databases]
type projectGetResponseDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                             `json:"id"`
	JSON projectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding]
type projectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces]
type projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                         `json:"namespace_id"`
	JSON        projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectGetResponseDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectGetResponseDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewEnvVarsJSON contains the JSON metadata
// for the struct [ProjectGetResponseDeploymentConfigsPreviewEnvVars]
type projectGetResponseDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                   `json:"value"`
	JSON  projectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type projectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, ProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsJSON       `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindings]
type projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID   string                                                                     `json:"id"`
	JSON projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive]
type projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding ProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectGetResponseDeploymentConfigsPreviewKVNamespacesJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewKVNamespacesJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewKVNamespaces]
type projectGetResponseDeploymentConfigsPreviewKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                              `json:"namespace_id"`
	JSON        projectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding]
type projectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectGetResponseDeploymentConfigsPreviewMTLSCertificatesJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewMTLSCertificatesJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewMTLSCertificates]
type projectGetResponseDeploymentConfigsPreviewMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID string                                                             `json:"certificate_id"`
	JSON          projectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLS]
type projectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                  `json:"mode"`
	JSON projectGetResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewPlacement]
type projectGetResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectGetResponseDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewQueueProducersJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewQueueProducers]
type projectGetResponseDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                           `json:"name"`
	JSON projectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type projectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding ProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectGetResponseDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewR2BucketsJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewR2Buckets]
type projectGetResponseDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                           `json:"name"`
	JSON projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding]
type projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding ProjectGetResponseDeploymentConfigsPreviewServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectGetResponseDeploymentConfigsPreviewServicesJSON           `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewServicesJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewServices]
type projectGetResponseDeploymentConfigsPreviewServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectGetResponseDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                               `json:"service"`
	JSON    projectGetResponseDeploymentConfigsPreviewServicesServiceBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewServicesServiceBindingJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewServicesServiceBinding]
type projectGetResponseDeploymentConfigsPreviewServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectGetResponseDeploymentConfigsPreviewVectorizeBindingsJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewVectorizeBindingsJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewVectorizeBindings]
type projectGetResponseDeploymentConfigsPreviewVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName string                                                                   `json:"index_name"`
	JSON      projectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorize]
type projectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectGetResponseDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectGetResponseDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers ProjectGetResponseDeploymentConfigsProductionBrowsers `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectGetResponseDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectGetResponseDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings ProjectGetResponseDeploymentConfigsProductionHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectGetResponseDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates ProjectGetResponseDeploymentConfigsProductionMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectGetResponseDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectGetResponseDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectGetResponseDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services ProjectGetResponseDeploymentConfigsProductionServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings ProjectGetResponseDeploymentConfigsProductionVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              projectGetResponseDeploymentConfigsProductionJSON              `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionJSON contains the JSON metadata for
// the struct [ProjectGetResponseDeploymentConfigsProduction]
type projectGetResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectGetResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding ProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      projectGetResponseDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionAIBindingsJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionAIBindings]
type projectGetResponseDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID string                                                               `json:"project_id"`
	JSON      projectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding]
type projectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets]
type projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                         `json:"dataset"`
	JSON    projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser interface{}                                               `json:"BROWSER"`
	JSON    projectGetResponseDeploymentConfigsProductionBrowsersJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionBrowsersJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionBrowsers]
type projectGetResponseDeploymentConfigsProductionBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionBrowsersJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding ProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      projectGetResponseDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionD1DatabasesJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionD1Databases]
type projectGetResponseDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                `json:"id"`
	JSON projectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding]
type projectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces]
type projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                            `json:"namespace_id"`
	JSON        projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type ProjectGetResponseDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                projectGetResponseDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionEnvVarsJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionEnvVars]
type projectGetResponseDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                      `json:"value"`
	JSON  projectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type projectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, ProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive ProjectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       projectGetResponseDeploymentConfigsProductionHyperdriveBindingsJSON       `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionHyperdriveBindingsJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionHyperdriveBindings]
type projectGetResponseDeploymentConfigsProductionHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID   string                                                                        `json:"id"`
	JSON projectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive]
type projectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

// KV namespaces used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding ProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      projectGetResponseDeploymentConfigsProductionKVNamespacesJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionKVNamespacesJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionKVNamespaces]
type projectGetResponseDeploymentConfigsProductionKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type ProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                 `json:"namespace_id"`
	JSON        projectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding]
type projectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS ProjectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLS `json:"MTLS"`
	JSON projectGetResponseDeploymentConfigsProductionMTLSCertificatesJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionMTLSCertificatesJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionMTLSCertificates]
type projectGetResponseDeploymentConfigsProductionMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID string                                                                `json:"certificate_id"`
	JSON          projectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLS]
type projectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                     `json:"mode"`
	JSON projectGetResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionPlacement]
type projectGetResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding ProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 projectGetResponseDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionQueueProducersJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionQueueProducers]
type projectGetResponseDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                              `json:"name"`
	JSON projectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type projectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding ProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      projectGetResponseDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionR2BucketsJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionR2Buckets]
type projectGetResponseDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                                              `json:"name"`
	JSON projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding]
type projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding ProjectGetResponseDeploymentConfigsProductionServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectGetResponseDeploymentConfigsProductionServicesJSON           `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionServicesJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionServices]
type projectGetResponseDeploymentConfigsProductionServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectGetResponseDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                  `json:"service"`
	JSON    projectGetResponseDeploymentConfigsProductionServicesServiceBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionServicesServiceBindingJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionServicesServiceBinding]
type projectGetResponseDeploymentConfigsProductionServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize ProjectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      projectGetResponseDeploymentConfigsProductionVectorizeBindingsJSON      `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionVectorizeBindingsJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionVectorizeBindings]
type projectGetResponseDeploymentConfigsProductionVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName string                                                                      `json:"index_name"`
	JSON      projectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorize]
type projectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseSource struct {
	Config ProjectGetResponseSourceConfig `json:"config"`
	Type   string                         `json:"type"`
	JSON   projectGetResponseSourceJSON   `json:"-"`
}

// projectGetResponseSourceJSON contains the JSON metadata for the struct
// [ProjectGetResponseSource]
type projectGetResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseSourceConfig struct {
	DeploymentsEnabled           bool                                                   `json:"deployments_enabled"`
	Owner                        string                                                 `json:"owner"`
	PathExcludes                 []string                                               `json:"path_excludes"`
	PathIncludes                 []string                                               `json:"path_includes"`
	PrCommentsEnabled            bool                                                   `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                                               `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                                               `json:"preview_branch_includes"`
	PreviewDeploymentSetting     ProjectGetResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                                                 `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                                   `json:"production_deployments_enabled"`
	RepoName                     string                                                 `json:"repo_name"`
	JSON                         projectGetResponseSourceConfigJSON                     `json:"-"`
}

// projectGetResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectGetResponseSourceConfig]
type projectGetResponseSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectGetResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectGetResponseSourceConfigPreviewDeploymentSettingAll    ProjectGetResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectGetResponseSourceConfigPreviewDeploymentSettingNone   ProjectGetResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectGetResponseSourceConfigPreviewDeploymentSettingCustom ProjectGetResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectGetResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectGetResponseSourceConfigPreviewDeploymentSettingAll, ProjectGetResponseSourceConfigPreviewDeploymentSettingNone, ProjectGetResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

type ProjectPurgeBuildCacheResponse = interface{}

type ProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig param.Field[ProjectNewParamsBuildConfig] `json:"build_config"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectNewParamsDeploymentConfigs] `json:"deployment_configs"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type ProjectNewParamsBuildConfig struct {
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

func (r ProjectNewParamsBuildConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectNewParamsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview param.Field[ProjectNewParamsDeploymentConfigsPreview] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[ProjectNewParamsDeploymentConfigsProduction] `json:"production"`
}

func (r ProjectNewParamsDeploymentConfigs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type ProjectNewParamsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectNewParamsDeploymentConfigsPreviewAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[ProjectNewParamsDeploymentConfigsPreviewBrowsers] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[ProjectNewParamsDeploymentConfigsPreviewD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[ProjectNewParamsDeploymentConfigsPreviewEnvVars] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectNewParamsDeploymentConfigsPreviewKVNamespaces] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectNewParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectNewParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectNewParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[ProjectNewParamsDeploymentConfigsPreviewServices] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings] `json:"vectorize_bindings"`
}

func (r ProjectNewParamsDeploymentConfigsPreview) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser param.Field[interface{}] `json:"BROWSER"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding param.Field[ProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type ProjectNewParamsDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive param.Field[ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive] `json:"HYPERDRIVE"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive binding.
type ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespaces used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding] `json:"KV_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type ProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS param.Field[ProjectNewParamsDeploymentConfigsPreviewMTLSCertificatesMTLS] `json:"MTLS"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectNewParamsDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewMTLSCertificatesMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[ProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewServicesServiceBinding] `json:"SERVICE_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectNewParamsDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewServicesServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize param.Field[ProjectNewParamsDeploymentConfigsPreviewVectorizeBindingsVectorize] `json:"VECTORIZE"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectNewParamsDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewVectorizeBindingsVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectNewParamsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectNewParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[ProjectNewParamsDeploymentConfigsProductionBrowsers] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[ProjectNewParamsDeploymentConfigsProductionD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[ProjectNewParamsDeploymentConfigsProductionEnvVars] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectNewParamsDeploymentConfigsProductionKVNamespaces] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[ProjectNewParamsDeploymentConfigsProductionMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectNewParamsDeploymentConfigsProductionPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectNewParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectNewParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[ProjectNewParamsDeploymentConfigsProductionServices] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[ProjectNewParamsDeploymentConfigsProductionVectorizeBindings] `json:"vectorize_bindings"`
}

func (r ProjectNewParamsDeploymentConfigsProduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding param.Field[ProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r ProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser param.Field[interface{}] `json:"BROWSER"`
}

func (r ProjectNewParamsDeploymentConfigsProductionBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding param.Field[ProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type ProjectNewParamsDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive param.Field[ProjectNewParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive] `json:"HYPERDRIVE"`
}

func (r ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive binding.
type ProjectNewParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectNewParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespaces used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding param.Field[ProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding] `json:"KV_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type ProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS param.Field[ProjectNewParamsDeploymentConfigsProductionMTLSCertificatesMTLS] `json:"MTLS"`
}

func (r ProjectNewParamsDeploymentConfigsProductionMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectNewParamsDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectNewParamsDeploymentConfigsProductionMTLSCertificatesMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectNewParamsDeploymentConfigsProductionPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[ProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[ProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding param.Field[ProjectNewParamsDeploymentConfigsProductionServicesServiceBinding] `json:"SERVICE_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectNewParamsDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectNewParamsDeploymentConfigsProductionServicesServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize bindings used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize param.Field[ProjectNewParamsDeploymentConfigsProductionVectorizeBindingsVectorize] `json:"VECTORIZE"`
}

func (r ProjectNewParamsDeploymentConfigsProductionVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectNewParamsDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectNewParamsDeploymentConfigsProductionVectorizeBindingsVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ProjectNewResponse    `json:"result,required"`
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
	ProjectNewResponseEnvelopeSuccessFalse ProjectNewResponseEnvelopeSuccess = false
	ProjectNewResponseEnvelopeSuccessTrue  ProjectNewResponseEnvelopeSuccess = true
)

func (r ProjectNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectNewResponseEnvelopeSuccessFalse, ProjectNewResponseEnvelopeSuccessTrue:
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

type ProjectDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ProjectDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDeleteResponseEnvelope]
type projectDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeleteResponseEnvelopeSuccessFalse ProjectDeleteResponseEnvelopeSuccess = false
	ProjectDeleteResponseEnvelopeSuccessTrue  ProjectDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeleteResponseEnvelopeSuccessFalse, ProjectDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig param.Field[ProjectEditParamsBuildConfig] `json:"build_config"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectEditParamsDeploymentConfigs] `json:"deployment_configs"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
}

func (r ProjectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type ProjectEditParamsBuildConfig struct {
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

func (r ProjectEditParamsBuildConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectEditParamsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview param.Field[ProjectEditParamsDeploymentConfigsPreview] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[ProjectEditParamsDeploymentConfigsProduction] `json:"production"`
}

func (r ProjectEditParamsDeploymentConfigs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type ProjectEditParamsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectEditParamsDeploymentConfigsPreviewAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[ProjectEditParamsDeploymentConfigsPreviewBrowsers] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[ProjectEditParamsDeploymentConfigsPreviewD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[ProjectEditParamsDeploymentConfigsPreviewEnvVars] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectEditParamsDeploymentConfigsPreviewKVNamespaces] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectEditParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectEditParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectEditParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[ProjectEditParamsDeploymentConfigsPreviewServices] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings] `json:"vectorize_bindings"`
}

func (r ProjectEditParamsDeploymentConfigsPreview) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding param.Field[ProjectEditParamsDeploymentConfigsPreviewAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectEditParamsDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser param.Field[interface{}] `json:"BROWSER"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding param.Field[ProjectEditParamsDeploymentConfigsPreviewD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectEditParamsDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type ProjectEditParamsDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive param.Field[ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive] `json:"HYPERDRIVE"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive binding.
type ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespaces used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding param.Field[ProjectEditParamsDeploymentConfigsPreviewKVNamespacesKVBinding] `json:"KV_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type ProjectEditParamsDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewKVNamespacesKVBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS param.Field[ProjectEditParamsDeploymentConfigsPreviewMTLSCertificatesMTLS] `json:"MTLS"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectEditParamsDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewMTLSCertificatesMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[ProjectEditParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectEditParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[ProjectEditParamsDeploymentConfigsPreviewR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectEditParamsDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding param.Field[ProjectEditParamsDeploymentConfigsPreviewServicesServiceBinding] `json:"SERVICE_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectEditParamsDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewServicesServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize param.Field[ProjectEditParamsDeploymentConfigsPreviewVectorizeBindingsVectorize] `json:"VECTORIZE"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectEditParamsDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewVectorizeBindingsVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectEditParamsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectEditParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[ProjectEditParamsDeploymentConfigsProductionBrowsers] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[ProjectEditParamsDeploymentConfigsProductionD1Databases] `json:"d1_databases"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables for build configs.
	EnvVars param.Field[ProjectEditParamsDeploymentConfigsProductionEnvVars] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectEditParamsDeploymentConfigsProductionKVNamespaces] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[ProjectEditParamsDeploymentConfigsProductionMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectEditParamsDeploymentConfigsProductionPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectEditParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectEditParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[ProjectEditParamsDeploymentConfigsProductionServices] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[ProjectEditParamsDeploymentConfigsProductionVectorizeBindings] `json:"vectorize_bindings"`
}

func (r ProjectEditParamsDeploymentConfigsProduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Constellation bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding param.Field[ProjectEditParamsDeploymentConfigsProductionAIBindingsAIBinding] `json:"AI_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectEditParamsDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r ProjectEditParamsDeploymentConfigsProductionAIBindingsAIBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding param.Field[ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding] `json:"ANALYTICS_ENGINE_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser param.Field[interface{}] `json:"BROWSER"`
}

func (r ProjectEditParamsDeploymentConfigsProductionBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 databases used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding param.Field[ProjectEditParamsDeploymentConfigsProductionD1DatabasesD1Binding] `json:"D1_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectEditParamsDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectEditParamsDeploymentConfigsProductionD1DatabasesD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object namespaces used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding param.Field[ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding] `json:"DO_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durabble Object binding.
type ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variables for build configs.
type ProjectEditParamsDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable param.Field[ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable] `json:"ENVIRONMENT_VARIABLE"`
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment variable.
type ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type param.Field[ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType] `json:"type"`
	// Environment variable value.
	Value param.Field[string] `json:"value"`
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of environment variable (plain text or secret)
type ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive param.Field[ProjectEditParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive] `json:"HYPERDRIVE"`
}

func (r ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Hyperdrive binding.
type ProjectEditParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectEditParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespaces used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding param.Field[ProjectEditParamsDeploymentConfigsProductionKVNamespacesKVBinding] `json:"KV_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV binding.
type ProjectEditParamsDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectEditParamsDeploymentConfigsProductionKVNamespacesKVBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS param.Field[ProjectEditParamsDeploymentConfigsProductionMTLSCertificatesMTLS] `json:"MTLS"`
}

func (r ProjectEditParamsDeploymentConfigsProductionMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectEditParamsDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectEditParamsDeploymentConfigsProductionMTLSCertificatesMTLS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectEditParamsDeploymentConfigsProductionPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding param.Field[ProjectEditParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding] `json:"QUEUE_PRODUCER_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectEditParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectEditParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 buckets used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding param.Field[ProjectEditParamsDeploymentConfigsProductionR2BucketsR2Binding] `json:"R2_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectEditParamsDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectEditParamsDeploymentConfigsProductionR2BucketsR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Services used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding param.Field[ProjectEditParamsDeploymentConfigsProductionServicesServiceBinding] `json:"SERVICE_BINDING"`
}

func (r ProjectEditParamsDeploymentConfigsProductionServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectEditParamsDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectEditParamsDeploymentConfigsProductionServicesServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize bindings used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize param.Field[ProjectEditParamsDeploymentConfigsProductionVectorizeBindingsVectorize] `json:"VECTORIZE"`
}

func (r ProjectEditParamsDeploymentConfigsProductionVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectEditParamsDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectEditParamsDeploymentConfigsProductionVectorizeBindingsVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ProjectEditResponse   `json:"result,required"`
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
	ProjectEditResponseEnvelopeSuccessFalse ProjectEditResponseEnvelopeSuccess = false
	ProjectEditResponseEnvelopeSuccessTrue  ProjectEditResponseEnvelopeSuccess = true
)

func (r ProjectEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectEditResponseEnvelopeSuccessFalse, ProjectEditResponseEnvelopeSuccessTrue:
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
	Result   ProjectGetResponse    `json:"result,required"`
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
	ProjectGetResponseEnvelopeSuccessFalse ProjectGetResponseEnvelopeSuccess = false
	ProjectGetResponseEnvelopeSuccessTrue  ProjectGetResponseEnvelopeSuccess = true
)

func (r ProjectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectGetResponseEnvelopeSuccessFalse, ProjectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectPurgeBuildCacheParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectPurgeBuildCacheResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ProjectPurgeBuildCacheResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectPurgeBuildCacheResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectPurgeBuildCacheResponseEnvelopeJSON    `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectPurgeBuildCacheResponseEnvelope]
type projectPurgeBuildCacheResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectPurgeBuildCacheResponseEnvelopeSuccess bool

const (
	ProjectPurgeBuildCacheResponseEnvelopeSuccessFalse ProjectPurgeBuildCacheResponseEnvelopeSuccess = false
	ProjectPurgeBuildCacheResponseEnvelopeSuccessTrue  ProjectPurgeBuildCacheResponseEnvelopeSuccess = true
)

func (r ProjectPurgeBuildCacheResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectPurgeBuildCacheResponseEnvelopeSuccessFalse, ProjectPurgeBuildCacheResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
