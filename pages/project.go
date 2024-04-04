// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ProjectService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewProjectService] method instead.
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
	opts = append(r.Options[:], opts...)
	var env ProjectNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all user projects.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PagesDeployments], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PagesDeployments] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a project by name.
func (r *ProjectService) Delete(ctx context.Context, projectName string, params ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *ProjectService) Edit(ctx context.Context, projectName string, params ProjectEditParams, opts ...option.RequestOption) (res *ProjectEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a project by name.
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *PagesProjects, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
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

func (r pagesDeploymentsJSON) RawJSON() string {
	return r.raw
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

func (r pagesDeploymentsDeploymentTriggerJSON) RawJSON() string {
	return r.raw
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

func (r pagesDeploymentsDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
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

func (r pagesDeploymentsStageJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsBuildConfigJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type PagesProjectsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings PagesProjectsDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers PagesProjectsDeploymentConfigsPreviewBrowsers `json:"browsers,nullable"`
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
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings PagesProjectsDeploymentConfigsPreviewHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces PagesProjectsDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates PagesProjectsDeploymentConfigsPreviewMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement PagesProjectsDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers PagesProjectsDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets PagesProjectsDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services PagesProjectsDeploymentConfigsPreviewServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings PagesProjectsDeploymentConfigsPreviewVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              pagesProjectsDeploymentConfigsPreviewJSON              `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewJSON contains the JSON metadata for the
// struct [PagesProjectsDeploymentConfigsPreview]
type pagesProjectsDeploymentConfigsPreviewJSON struct {
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

func (r *PagesProjectsDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewAIBindingsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewBrowsers struct {
	// Browser binding.
	Browser interface{}                                       `json:"BROWSER"`
	JSON    pagesProjectsDeploymentConfigsPreviewBrowsersJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewBrowsersJSON contains the JSON metadata for
// the struct [PagesProjectsDeploymentConfigsPreviewBrowsers]
type pagesProjectsDeploymentConfigsPreviewBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewBrowsersJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewD1DatabasesJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewEnvVarsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

func (r PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText, PagesProjectsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive PagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsJSON       `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewHyperdriveBindings]
type pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type PagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive struct {
	ID   string                                                                `json:"id"`
	JSON pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON contains
// the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive]
type pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

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

func (r pagesProjectsDeploymentConfigsPreviewKVNamespacesJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewMTLSCertificates struct {
	// mTLS binding.
	MTLS PagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLS `json:"MTLS"`
	JSON pagesProjectsDeploymentConfigsPreviewMTLSCertificatesJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewMTLSCertificatesJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsPreviewMTLSCertificates]
type pagesProjectsDeploymentConfigsPreviewMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type PagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLS struct {
	CertificateID string                                                        `json:"certificate_id"`
	JSON          pagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLSJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLSJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLS]
type pagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewQueueProducersJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewR2BucketsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewServices struct {
	// Service binding.
	ServiceBinding PagesProjectsDeploymentConfigsPreviewServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           pagesProjectsDeploymentConfigsPreviewServicesJSON           `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewServicesJSON contains the JSON metadata for
// the struct [PagesProjectsDeploymentConfigsPreviewServices]
type pagesProjectsDeploymentConfigsPreviewServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type PagesProjectsDeploymentConfigsPreviewServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                          `json:"service"`
	JSON    pagesProjectsDeploymentConfigsPreviewServicesServiceBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewServicesServiceBindingJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewServicesServiceBinding]
type pagesProjectsDeploymentConfigsPreviewServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsPreviewVectorizeBindings struct {
	// Vectorize binding.
	Vectorize PagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      pagesProjectsDeploymentConfigsPreviewVectorizeBindingsJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewVectorizeBindingsJSON contains the JSON
// metadata for the struct [PagesProjectsDeploymentConfigsPreviewVectorizeBindings]
type pagesProjectsDeploymentConfigsPreviewVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type PagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorize struct {
	IndexName string                                                              `json:"index_name"`
	JSON      pagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorize]
type pagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsPreviewVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type PagesProjectsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings PagesProjectsDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers PagesProjectsDeploymentConfigsProductionBrowsers `json:"browsers,nullable"`
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
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings PagesProjectsDeploymentConfigsProductionHyperdriveBindings `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces PagesProjectsDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates PagesProjectsDeploymentConfigsProductionMTLSCertificates `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement PagesProjectsDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers PagesProjectsDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets PagesProjectsDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services PagesProjectsDeploymentConfigsProductionServices `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings PagesProjectsDeploymentConfigsProductionVectorizeBindings `json:"vectorize_bindings,nullable"`
	JSON              pagesProjectsDeploymentConfigsProductionJSON              `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionJSON contains the JSON metadata for the
// struct [PagesProjectsDeploymentConfigsProduction]
type pagesProjectsDeploymentConfigsProductionJSON struct {
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

func (r *PagesProjectsDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionAIBindingsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// Browser bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionBrowsers struct {
	// Browser binding.
	Browser interface{}                                          `json:"BROWSER"`
	JSON    pagesProjectsDeploymentConfigsProductionBrowsersJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionBrowsersJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsProductionBrowsers]
type pagesProjectsDeploymentConfigsProductionBrowsersJSON struct {
	Browser     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionBrowsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionBrowsersJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionD1DatabasesJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionEnvVarsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

func (r PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableType) IsKnown() bool {
	switch r {
	case PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText, PagesProjectsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionHyperdriveBindings struct {
	// Hyperdrive binding.
	Hyperdrive PagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdrive `json:"HYPERDRIVE"`
	JSON       pagesProjectsDeploymentConfigsProductionHyperdriveBindingsJSON       `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionHyperdriveBindingsJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsProductionHyperdriveBindings]
type pagesProjectsDeploymentConfigsProductionHyperdriveBindingsJSON struct {
	Hyperdrive  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionHyperdriveBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionHyperdriveBindingsJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type PagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdrive struct {
	ID   string                                                                   `json:"id"`
	JSON pagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON
// contains the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdrive]
type pagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionHyperdriveBindingsHyperdriveJSON) RawJSON() string {
	return r.raw
}

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

func (r pagesProjectsDeploymentConfigsProductionKVNamespacesJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// mTLS bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionMTLSCertificates struct {
	// mTLS binding.
	MTLS PagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLS `json:"MTLS"`
	JSON pagesProjectsDeploymentConfigsProductionMTLSCertificatesJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionMTLSCertificatesJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsProductionMTLSCertificates]
type pagesProjectsDeploymentConfigsProductionMTLSCertificatesJSON struct {
	MTLS        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionMTLSCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionMTLSCertificatesJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type PagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLS struct {
	CertificateID string                                                           `json:"certificate_id"`
	JSON          pagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLSJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLSJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLS]
type pagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLSJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionMTLSCertificatesMTLSJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionQueueProducersJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionR2BucketsJSON) RawJSON() string {
	return r.raw
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

func (r pagesProjectsDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionServices struct {
	// Service binding.
	ServiceBinding PagesProjectsDeploymentConfigsProductionServicesServiceBinding `json:"SERVICE_BINDING"`
	JSON           pagesProjectsDeploymentConfigsProductionServicesJSON           `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionServicesJSON contains the JSON metadata
// for the struct [PagesProjectsDeploymentConfigsProductionServices]
type pagesProjectsDeploymentConfigsProductionServicesJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionServicesJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type PagesProjectsDeploymentConfigsProductionServicesServiceBinding struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                             `json:"service"`
	JSON    pagesProjectsDeploymentConfigsProductionServicesServiceBindingJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionServicesServiceBindingJSON contains the
// JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionServicesServiceBinding]
type pagesProjectsDeploymentConfigsProductionServicesServiceBindingJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionServicesServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionServicesServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Vectorize bindings used for Pages Functions.
type PagesProjectsDeploymentConfigsProductionVectorizeBindings struct {
	// Vectorize binding.
	Vectorize PagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorize `json:"VECTORIZE"`
	JSON      pagesProjectsDeploymentConfigsProductionVectorizeBindingsJSON      `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionVectorizeBindingsJSON contains the JSON
// metadata for the struct
// [PagesProjectsDeploymentConfigsProductionVectorizeBindings]
type pagesProjectsDeploymentConfigsProductionVectorizeBindingsJSON struct {
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionVectorizeBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionVectorizeBindingsJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type PagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorize struct {
	IndexName string                                                                 `json:"index_name"`
	JSON      pagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorizeJSON `json:"-"`
}

// pagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorizeJSON contains
// the JSON metadata for the struct
// [PagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorize]
type pagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorizeJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagesProjectsDeploymentConfigsProductionVectorizeBindingsVectorizeJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [pages.ProjectNewResponseUnknown],
// [pages.ProjectNewResponseArray] or [shared.UnionString].
type ProjectNewResponse interface {
	ImplementsPagesProjectNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponse)(nil)).Elem(),
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

func (r ProjectNewResponseArray) ImplementsPagesProjectNewResponse() {}

type ProjectDeleteResponse = interface{}

// Union satisfied by [pages.ProjectEditResponseUnknown],
// [pages.ProjectEditResponseArray] or [shared.UnionString].
type ProjectEditResponse interface {
	ImplementsPagesProjectEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectEditResponse)(nil)).Elem(),
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

func (r ProjectEditResponseArray) ImplementsPagesProjectEditResponse() {}

type ProjectPurgeBuildCacheResponse = interface{}

type ProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig         param.Field[ProjectNewParamsBuildConfig] `json:"build_config"`
	CanonicalDeployment param.Field[PagesDeploymentsParam]       `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectNewParamsDeploymentConfigs] `json:"deployment_configs"`
	LatestDeployment  param.Field[PagesDeploymentsParam]             `json:"latest_deployment"`
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
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
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
	ProjectID param.Field[interface{}] `json:"project_id"`
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
	CompatibilityFlags param.Field[[]interface{}] `json:"compatibility_flags"`
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
	ProjectID param.Field[interface{}] `json:"project_id"`
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
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectEditParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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
	Result   PagesProjects         `json:"result,required"`
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
