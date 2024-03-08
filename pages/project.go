// File generated from our OpenAPI spec by Stainless.

package pages

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *[]ProjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a project by name.
func (r *ProjectService) Delete(ctx context.Context, projectName string, body ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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

type ProjectListResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectListResponseDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []ProjectListResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                  `json:"url"`
	JSON projectListResponseJSON `json:"-"`
}

// projectListResponseJSON contains the JSON metadata for the struct
// [ProjectListResponse]
type projectListResponseJSON struct {
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

func (r *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectListResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectListResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                   `json:"type"`
	JSON projectListResponseDeploymentTriggerJSON `json:"-"`
}

// projectListResponseDeploymentTriggerJSON contains the JSON metadata for the
// struct [ProjectListResponseDeploymentTrigger]
type projectListResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectListResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                           `json:"commit_message"`
	JSON          projectListResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectListResponseDeploymentTriggerMetadataJSON contains the JSON metadata for
// the struct [ProjectListResponseDeploymentTriggerMetadata]
type projectListResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectListResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                       `json:"status"`
	JSON   projectListResponseStageJSON `json:"-"`
}

// projectListResponseStageJSON contains the JSON metadata for the struct
// [ProjectListResponseStage]
type projectListResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseStageJSON) RawJSON() string {
	return r.raw
}

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

type ProjectGetResponse struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig         ProjectGetResponseBuildConfig         `json:"build_config"`
	CanonicalDeployment ProjectGetResponseCanonicalDeployment `json:"canonical_deployment"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectGetResponseDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains          []interface{}                      `json:"domains"`
	LatestDeployment ProjectGetResponseLatestDeployment `json:"latest_deployment"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string      `json:"production_branch"`
	Source           interface{} `json:"source"`
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

type ProjectGetResponseCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectGetResponseCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []ProjectGetResponseCanonicalDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                    `json:"url"`
	JSON projectGetResponseCanonicalDeploymentJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentJSON contains the JSON metadata for the
// struct [ProjectGetResponseCanonicalDeployment]
type projectGetResponseCanonicalDeploymentJSON struct {
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

func (r *ProjectGetResponseCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectGetResponseCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                     `json:"type"`
	JSON projectGetResponseCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectGetResponseCanonicalDeploymentDeploymentTrigger]
type projectGetResponseCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                             `json:"commit_message"`
	JSON          projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata]
type projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectGetResponseCanonicalDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                         `json:"status"`
	JSON   projectGetResponseCanonicalDeploymentStageJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentStageJSON contains the JSON metadata for
// the struct [ProjectGetResponseCanonicalDeploymentStage]
type projectGetResponseCanonicalDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentStageJSON) RawJSON() string {
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
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectGetResponseDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectGetResponseDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectGetResponseDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement ProjectGetResponseDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectGetResponseDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectGetResponseDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings ProjectGetResponseDeploymentConfigsPreviewServiceBindings `json:"service_bindings,nullable"`
	JSON            projectGetResponseDeploymentConfigsPreviewJSON            `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectGetResponseDeploymentConfigsPreview]
type projectGetResponseDeploymentConfigsPreviewJSON struct {
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
	ProjectID interface{}                                                       `json:"project_id"`
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
	// Name of the R2 bucket.
	Name string                                                           `json:"name"`
	JSON projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding]
type projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding ProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectGetResponseDeploymentConfigsPreviewServiceBindingsJSON           `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewServiceBindingsJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewServiceBindings]
type projectGetResponseDeploymentConfigsPreviewServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewServiceBindingsJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                      `json:"service"`
	JSON    projectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding]
type projectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectGetResponseDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings ProjectGetResponseDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases ProjectGetResponseDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars ProjectGetResponseDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces ProjectGetResponseDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement ProjectGetResponseDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers ProjectGetResponseDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets ProjectGetResponseDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings ProjectGetResponseDeploymentConfigsProductionServiceBindings `json:"service_bindings,nullable"`
	JSON            projectGetResponseDeploymentConfigsProductionJSON            `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionJSON contains the JSON metadata for
// the struct [ProjectGetResponseDeploymentConfigsProduction]
type projectGetResponseDeploymentConfigsProductionJSON struct {
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
	ProjectID interface{}                                                          `json:"project_id"`
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
	// Name of the R2 bucket.
	Name string                                                              `json:"name"`
	JSON projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding]
type projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding ProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           projectGetResponseDeploymentConfigsProductionServiceBindingsJSON           `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionServiceBindingsJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionServiceBindings]
type projectGetResponseDeploymentConfigsProductionServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionServiceBindingsJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                         `json:"service"`
	JSON    projectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding]
type projectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectGetResponseLatestDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []ProjectGetResponseLatestDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                 `json:"url"`
	JSON projectGetResponseLatestDeploymentJSON `json:"-"`
}

// projectGetResponseLatestDeploymentJSON contains the JSON metadata for the struct
// [ProjectGetResponseLatestDeployment]
type projectGetResponseLatestDeploymentJSON struct {
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

func (r *ProjectGetResponseLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectGetResponseLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                  `json:"type"`
	JSON projectGetResponseLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectGetResponseLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectGetResponseLatestDeploymentDeploymentTrigger]
type projectGetResponseLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                          `json:"commit_message"`
	JSON          projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata]
type projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type ProjectGetResponseLatestDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                      `json:"status"`
	JSON   projectGetResponseLatestDeploymentStageJSON `json:"-"`
}

// projectGetResponseLatestDeploymentStageJSON contains the JSON metadata for the
// struct [ProjectGetResponseLatestDeploymentStage]
type projectGetResponseLatestDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentStageJSON) RawJSON() string {
	return r.raw
}

type ProjectPurgeBuildCacheResponse = interface{}

type ProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig         param.Field[ProjectNewParamsBuildConfig]         `json:"build_config"`
	CanonicalDeployment param.Field[ProjectNewParamsCanonicalDeployment] `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectNewParamsDeploymentConfigs] `json:"deployment_configs"`
	LatestDeployment  param.Field[ProjectNewParamsLatestDeployment]  `json:"latest_deployment"`
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

type ProjectNewParamsCanonicalDeployment struct {
}

func (r ProjectNewParamsCanonicalDeployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type ProjectNewParamsCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata param.Field[ProjectNewParamsCanonicalDeploymentDeploymentTriggerMetadata] `json:"metadata"`
}

func (r ProjectNewParamsCanonicalDeploymentDeploymentTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type ProjectNewParamsCanonicalDeploymentDeploymentTriggerMetadata struct {
}

func (r ProjectNewParamsCanonicalDeploymentDeploymentTriggerMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type ProjectNewParamsCanonicalDeploymentStage struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r ProjectNewParamsCanonicalDeploymentStage) MarshalJSON() (data []byte, err error) {
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
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectNewParamsDeploymentConfigsPreviewKVNamespaces] `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectNewParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectNewParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectNewParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	ServiceBindings param.Field[ProjectNewParamsDeploymentConfigsPreviewServiceBindings] `json:"service_bindings"`
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
type ProjectNewParamsDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding param.Field[ProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding] `json:"SERVICE_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewServiceBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectNewParamsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[ProjectNewParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
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
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[ProjectNewParamsDeploymentConfigsProductionKVNamespaces] `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectNewParamsDeploymentConfigsProductionPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[ProjectNewParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[ProjectNewParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	ServiceBindings param.Field[ProjectNewParamsDeploymentConfigsProductionServiceBindings] `json:"service_bindings"`
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
type ProjectNewParamsDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding param.Field[ProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding] `json:"SERVICE_BINDING"`
}

func (r ProjectNewParamsDeploymentConfigsProductionServiceBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewParamsLatestDeployment struct {
}

func (r ProjectNewParamsLatestDeployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type ProjectNewParamsLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata param.Field[ProjectNewParamsLatestDeploymentDeploymentTriggerMetadata] `json:"metadata"`
}

func (r ProjectNewParamsLatestDeploymentDeploymentTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type ProjectNewParamsLatestDeploymentDeploymentTriggerMetadata struct {
}

func (r ProjectNewParamsLatestDeploymentDeploymentTriggerMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type ProjectNewParamsLatestDeploymentStage struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r ProjectNewParamsLatestDeploymentStage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewResponseEnvelope struct {
	Errors   []ProjectNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectNewResponse                   `json:"result,required"`
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

type ProjectNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    projectNewResponseEnvelopeErrorsJSON `json:"-"`
}

// projectNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelopeErrors]
type projectNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    projectNewResponseEnvelopeMessagesJSON `json:"-"`
}

// projectNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelopeMessages]
type projectNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectNewResponseEnvelopeSuccess bool

const (
	ProjectNewResponseEnvelopeSuccessTrue ProjectNewResponseEnvelopeSuccess = true
)

type ProjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectListResponseEnvelope struct {
	Errors   []ProjectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ProjectListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    ProjectListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ProjectListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       projectListResponseEnvelopeJSON       `json:"-"`
}

// projectListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectListResponseEnvelope]
type projectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    projectListResponseEnvelopeErrorsJSON `json:"-"`
}

// projectListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectListResponseEnvelopeErrors]
type projectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    projectListResponseEnvelopeMessagesJSON `json:"-"`
}

// projectListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectListResponseEnvelopeMessages]
type projectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectListResponseEnvelopeSuccess bool

const (
	ProjectListResponseEnvelopeSuccessTrue ProjectListResponseEnvelopeSuccess = true
)

type ProjectListResponseEnvelopeResultInfo struct {
	Count      interface{}                               `json:"count"`
	Page       interface{}                               `json:"page"`
	PerPage    interface{}                               `json:"per_page"`
	TotalCount interface{}                               `json:"total_count"`
	JSON       projectListResponseEnvelopeResultInfoJSON `json:"-"`
}

// projectListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ProjectListResponseEnvelopeResultInfo]
type projectListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ProjectDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Errors   []ProjectEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectEditResponse                   `json:"result,required"`
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

type ProjectEditResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    projectEditResponseEnvelopeErrorsJSON `json:"-"`
}

// projectEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectEditResponseEnvelopeErrors]
type projectEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    projectEditResponseEnvelopeMessagesJSON `json:"-"`
}

// projectEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectEditResponseEnvelopeMessages]
type projectEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectEditResponseEnvelopeSuccess bool

const (
	ProjectEditResponseEnvelopeSuccessTrue ProjectEditResponseEnvelopeSuccess = true
)

type ProjectGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectGetResponseEnvelope struct {
	Errors   []ProjectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectGetResponse                   `json:"result,required"`
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

type ProjectGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    projectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// projectGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelopeErrors]
type projectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    projectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// projectGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelopeMessages]
type projectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectGetResponseEnvelopeSuccess bool

const (
	ProjectGetResponseEnvelopeSuccessTrue ProjectGetResponseEnvelopeSuccess = true
)

type ProjectPurgeBuildCacheParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
