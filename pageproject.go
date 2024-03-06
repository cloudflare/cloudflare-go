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
func (r *PageProjectService) List(ctx context.Context, query PageProjectListParams, opts ...option.RequestOption) (res *[]PageProjectListResponse, err error) {
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
func (r *PageProjectService) Get(ctx context.Context, projectName string, query PageProjectGetParams, opts ...option.RequestOption) (res *PageProjectGetResponse, err error) {
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

type PageProjectListResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectListResponseDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []PageProjectListResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                      `json:"url"`
	JSON pageProjectListResponseJSON `json:"-"`
}

// pageProjectListResponseJSON contains the JSON metadata for the struct
// [PageProjectListResponse]
type pageProjectListResponseJSON struct {
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

func (r *PageProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectListResponseJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type PageProjectListResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectListResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                       `json:"type"`
	JSON pageProjectListResponseDeploymentTriggerJSON `json:"-"`
}

// pageProjectListResponseDeploymentTriggerJSON contains the JSON metadata for the
// struct [PageProjectListResponseDeploymentTrigger]
type pageProjectListResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectListResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectListResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type PageProjectListResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                               `json:"commit_message"`
	JSON          pageProjectListResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectListResponseDeploymentTriggerMetadataJSON contains the JSON metadata
// for the struct [PageProjectListResponseDeploymentTriggerMetadata]
type pageProjectListResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectListResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectListResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type PageProjectListResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                           `json:"status"`
	JSON   pageProjectListResponseStageJSON `json:"-"`
}

// pageProjectListResponseStageJSON contains the JSON metadata for the struct
// [PageProjectListResponseStage]
type pageProjectListResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectListResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectListResponseStageJSON) RawJSON() string {
	return r.raw
}

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

type PageProjectGetResponse struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig         PageProjectGetResponseBuildConfig         `json:"build_config"`
	CanonicalDeployment PageProjectGetResponseCanonicalDeployment `json:"canonical_deployment"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs PageProjectGetResponseDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains          []interface{}                          `json:"domains"`
	LatestDeployment PageProjectGetResponseLatestDeployment `json:"latest_deployment"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string      `json:"production_branch"`
	Source           interface{} `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                     `json:"subdomain"`
	JSON      pageProjectGetResponseJSON `json:"-"`
}

// pageProjectGetResponseJSON contains the JSON metadata for the struct
// [PageProjectGetResponse]
type pageProjectGetResponseJSON struct {
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

func (r *PageProjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type PageProjectGetResponseBuildConfig struct {
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
	WebAnalyticsToken string                                `json:"web_analytics_token,nullable"`
	JSON              pageProjectGetResponseBuildConfigJSON `json:"-"`
}

// pageProjectGetResponseBuildConfigJSON contains the JSON metadata for the struct
// [PageProjectGetResponseBuildConfig]
type pageProjectGetResponseBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageProjectGetResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

type PageProjectGetResponseCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectGetResponseCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []PageProjectGetResponseCanonicalDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                        `json:"url"`
	JSON pageProjectGetResponseCanonicalDeploymentJSON `json:"-"`
}

// pageProjectGetResponseCanonicalDeploymentJSON contains the JSON metadata for the
// struct [PageProjectGetResponseCanonicalDeployment]
type pageProjectGetResponseCanonicalDeploymentJSON struct {
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

func (r *PageProjectGetResponseCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type PageProjectGetResponseCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                         `json:"type"`
	JSON pageProjectGetResponseCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// pageProjectGetResponseCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseCanonicalDeploymentDeploymentTrigger]
type pageProjectGetResponseCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseCanonicalDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type PageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                                 `json:"commit_message"`
	JSON          pageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata]
type pageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type PageProjectGetResponseCanonicalDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                             `json:"status"`
	JSON   pageProjectGetResponseCanonicalDeploymentStageJSON `json:"-"`
}

// pageProjectGetResponseCanonicalDeploymentStageJSON contains the JSON metadata
// for the struct [PageProjectGetResponseCanonicalDeploymentStage]
type pageProjectGetResponseCanonicalDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseCanonicalDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseCanonicalDeploymentStageJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type PageProjectGetResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview PageProjectGetResponseDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production PageProjectGetResponseDeploymentConfigsProduction `json:"production"`
	JSON       pageProjectGetResponseDeploymentConfigsJSON       `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [PageProjectGetResponseDeploymentConfigs]
type pageProjectGetResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type PageProjectGetResponseDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings PageProjectGetResponseDeploymentConfigsPreviewAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases PageProjectGetResponseDeploymentConfigsPreviewD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars PageProjectGetResponseDeploymentConfigsPreviewEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces PageProjectGetResponseDeploymentConfigsPreviewKVNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement PageProjectGetResponseDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers PageProjectGetResponseDeploymentConfigsPreviewQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets PageProjectGetResponseDeploymentConfigsPreviewR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings PageProjectGetResponseDeploymentConfigsPreviewServiceBindings `json:"service_bindings,nullable"`
	JSON            pageProjectGetResponseDeploymentConfigsPreviewJSON            `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewJSON contains the JSON metadata
// for the struct [PageProjectGetResponseDeploymentConfigsPreview]
type pageProjectGetResponseDeploymentConfigsPreviewJSON struct {
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

func (r *PageProjectGetResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewAIBindings struct {
	// AI binding.
	AIBinding PageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsPreviewAIBindingsJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewAIBindingsJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewAIBindings]
type pageProjectGetResponseDeploymentConfigsPreviewAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type PageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding struct {
	ProjectID interface{}                                                           `json:"project_id"`
	JSON      pageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding]
type pageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets]
type pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                          `json:"dataset"`
	JSON    pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding]
type pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewD1Databases struct {
	// D1 binding.
	D1Binding PageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewD1Databases]
type pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type PageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                 `json:"id"`
	JSON pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding]
type pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces]
type pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                             `json:"namespace_id"`
	JSON        pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding]
type pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type PageProjectGetResponseDeploymentConfigsPreviewEnvVars struct {
	// Environment variable.
	EnvironmentVariable PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                pageProjectGetResponseDeploymentConfigsPreviewEnvVarsJSON                `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewEnvVarsJSON contains the JSON
// metadata for the struct [PageProjectGetResponseDeploymentConfigsPreviewEnvVars]
type pageProjectGetResponseDeploymentConfigsPreviewEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                       `json:"value"`
	JSON  pageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable]
type pageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType string

const (
	PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText  PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "plain_text"
	PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypeSecretText PageProjectGetResponseDeploymentConfigsPreviewEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewKVNamespaces struct {
	// KV binding.
	KVBinding PageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewKVNamespaces]
type pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type PageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                  `json:"namespace_id"`
	JSON        pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding]
type pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                      `json:"mode"`
	JSON pageProjectGetResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewPlacement]
type pageProjectGetResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding PageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 pageProjectGetResponseDeploymentConfigsPreviewQueueProducersJSON                 `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewQueueProducersJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewQueueProducers]
type pageProjectGetResponseDeploymentConfigsPreviewQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type PageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                               `json:"name"`
	JSON pageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding]
type pageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewR2Buckets struct {
	// R2 binding.
	R2Binding PageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsPreviewR2BucketsJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewR2BucketsJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewR2Buckets]
type pageProjectGetResponseDeploymentConfigsPreviewR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type PageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                               `json:"name"`
	JSON pageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding]
type pageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsPreviewServiceBindings struct {
	// Service binding.
	ServiceBinding PageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsJSON           `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewServiceBindings]
type pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type PageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                          `json:"service"`
	JSON    pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding]
type pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsPreviewServiceBindingsServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type PageProjectGetResponseDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings PageProjectGetResponseDeploymentConfigsProductionAIBindings `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets `json:"analytics_engine_datasets,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []interface{} `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases PageProjectGetResponseDeploymentConfigsProductionD1Databases `json:"d1_databases,nullable"`
	// Durabble Object namespaces used for Pages Functions.
	DurableObjectNamespaces PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces `json:"durable_object_namespaces,nullable"`
	// Environment variables for build configs.
	EnvVars PageProjectGetResponseDeploymentConfigsProductionEnvVars `json:"env_vars,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces PageProjectGetResponseDeploymentConfigsProductionKVNamespaces `json:"kv_namespaces"`
	// Placement setting used for Pages Functions.
	Placement PageProjectGetResponseDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers PageProjectGetResponseDeploymentConfigsProductionQueueProducers `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets PageProjectGetResponseDeploymentConfigsProductionR2Buckets `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	ServiceBindings PageProjectGetResponseDeploymentConfigsProductionServiceBindings `json:"service_bindings,nullable"`
	JSON            pageProjectGetResponseDeploymentConfigsProductionJSON            `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionJSON contains the JSON metadata
// for the struct [PageProjectGetResponseDeploymentConfigsProduction]
type pageProjectGetResponseDeploymentConfigsProductionJSON struct {
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

func (r *PageProjectGetResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// Constellation bindings used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionAIBindings struct {
	// AI binding.
	AIBinding PageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding `json:"AI_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsProductionAIBindingsJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionAIBindingsJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionAIBindings]
type pageProjectGetResponseDeploymentConfigsProductionAIBindingsJSON struct {
	AIBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionAIBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionAIBindingsJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type PageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding struct {
	ProjectID interface{}                                                              `json:"project_id"`
	JSON      pageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding]
type pageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionAIBindingsAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine bindings used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Analytics Engine binding.
	AnalyticsEngineBinding PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding `json:"ANALYTICS_ENGINE_BINDING"`
	JSON                   pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON                   `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets]
type pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON struct {
	AnalyticsEngineBinding apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding struct {
	// Name of the dataset.
	Dataset string                                                                                             `json:"dataset"`
	JSON    pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding]
type pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingJSON) RawJSON() string {
	return r.raw
}

// D1 databases used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionD1Databases struct {
	// D1 binding.
	D1Binding PageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding `json:"D1_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsProductionD1DatabasesJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionD1DatabasesJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionD1Databases]
type pageProjectGetResponseDeploymentConfigsProductionD1DatabasesJSON struct {
	D1Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionD1Databases) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionD1DatabasesJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type PageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding struct {
	// UUID of the D1 database.
	ID   string                                                                    `json:"id"`
	JSON pageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding]
type pageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionD1DatabasesD1BindingJSON) RawJSON() string {
	return r.raw
}

// Durabble Object namespaces used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces struct {
	// Durabble Object binding.
	DoBinding PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding `json:"DO_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces]
type pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON struct {
	DoBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesJSON) RawJSON() string {
	return r.raw
}

// Durabble Object binding.
type PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding struct {
	// ID of the Durabble Object namespace.
	NamespaceID string                                                                                `json:"namespace_id"`
	JSON        pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding]
type pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionDurableObjectNamespacesDoBindingJSON) RawJSON() string {
	return r.raw
}

// Environment variables for build configs.
type PageProjectGetResponseDeploymentConfigsProductionEnvVars struct {
	// Environment variable.
	EnvironmentVariable PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable `json:"ENVIRONMENT_VARIABLE"`
	JSON                pageProjectGetResponseDeploymentConfigsProductionEnvVarsJSON                `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionEnvVarsJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionEnvVars]
type pageProjectGetResponseDeploymentConfigsProductionEnvVarsJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionEnvVars) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionEnvVarsJSON) RawJSON() string {
	return r.raw
}

// Environment variable.
type PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable struct {
	// The type of environment variable (plain text or secret)
	Type PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType `json:"type"`
	// Environment variable value.
	Value string                                                                          `json:"value"`
	JSON  pageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable]
type pageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The type of environment variable (plain text or secret)
type PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType string

const (
	PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText  PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "plain_text"
	PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableTypeSecretText PageProjectGetResponseDeploymentConfigsProductionEnvVarsEnvironmentVariableType = "secret_text"
)

// KV namespaces used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionKVNamespaces struct {
	// KV binding.
	KVBinding PageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding `json:"KV_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsProductionKVNamespacesJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionKVNamespacesJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionKVNamespaces]
type pageProjectGetResponseDeploymentConfigsProductionKVNamespacesJSON struct {
	KVBinding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionKVNamespaces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionKVNamespacesJSON) RawJSON() string {
	return r.raw
}

// KV binding.
type PageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding struct {
	// ID of the KV namespace.
	NamespaceID string                                                                     `json:"namespace_id"`
	JSON        pageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding]
type pageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionKVNamespacesKVBindingJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                         `json:"mode"`
	JSON pageProjectGetResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionPlacement]
type pageProjectGetResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer bindings used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionQueueProducers struct {
	// Queue Producer binding.
	QueueProducerBinding PageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding `json:"QUEUE_PRODUCER_BINDING"`
	JSON                 pageProjectGetResponseDeploymentConfigsProductionQueueProducersJSON                 `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionQueueProducersJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionQueueProducers]
type pageProjectGetResponseDeploymentConfigsProductionQueueProducersJSON struct {
	QueueProducerBinding apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionQueueProducers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionQueueProducersJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type PageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding struct {
	// Name of the Queue.
	Name string                                                                                  `json:"name"`
	JSON pageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding]
type pageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionQueueProducersQueueProducerBindingJSON) RawJSON() string {
	return r.raw
}

// R2 buckets used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionR2Buckets struct {
	// R2 binding.
	R2Binding PageProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding `json:"R2_BINDING"`
	JSON      pageProjectGetResponseDeploymentConfigsProductionR2BucketsJSON      `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionR2BucketsJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionR2Buckets]
type pageProjectGetResponseDeploymentConfigsProductionR2BucketsJSON struct {
	R2Binding   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionR2Buckets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionR2BucketsJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type PageProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding struct {
	// Name of the R2 bucket.
	Name string                                                                  `json:"name"`
	JSON pageProjectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding]
type pageProjectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionR2BucketsR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionR2BucketsR2BindingJSON) RawJSON() string {
	return r.raw
}

// Services used for Pages Functions.
type PageProjectGetResponseDeploymentConfigsProductionServiceBindings struct {
	// Service binding.
	ServiceBinding PageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding `json:"SERVICE_BINDING"`
	JSON           pageProjectGetResponseDeploymentConfigsProductionServiceBindingsJSON           `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionServiceBindingsJSON contains
// the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionServiceBindings]
type pageProjectGetResponseDeploymentConfigsProductionServiceBindingsJSON struct {
	ServiceBinding apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionServiceBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionServiceBindingsJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type PageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding struct {
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                                                             `json:"service"`
	JSON    pageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON `json:"-"`
}

// pageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON
// contains the JSON metadata for the struct
// [PageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding]
type pageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseDeploymentConfigsProductionServiceBindingsServiceBindingJSON) RawJSON() string {
	return r.raw
}

type PageProjectGetResponseLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectGetResponseLatestDeploymentDeploymentTrigger `json:"deployment_trigger"`
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
	Stages []PageProjectGetResponseLatestDeploymentStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                     `json:"url"`
	JSON pageProjectGetResponseLatestDeploymentJSON `json:"-"`
}

// pageProjectGetResponseLatestDeploymentJSON contains the JSON metadata for the
// struct [PageProjectGetResponseLatestDeployment]
type pageProjectGetResponseLatestDeploymentJSON struct {
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

func (r *PageProjectGetResponseLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type PageProjectGetResponseLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectGetResponseLatestDeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                      `json:"type"`
	JSON pageProjectGetResponseLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// pageProjectGetResponseLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct
// [PageProjectGetResponseLatestDeploymentDeploymentTrigger]
type pageProjectGetResponseLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseLatestDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type PageProjectGetResponseLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                              `json:"commit_message"`
	JSON          pageProjectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [PageProjectGetResponseLatestDeploymentDeploymentTriggerMetadata]
type pageProjectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectGetResponseLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// The status of the deployment.
type PageProjectGetResponseLatestDeploymentStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                          `json:"status"`
	JSON   pageProjectGetResponseLatestDeploymentStageJSON `json:"-"`
}

// pageProjectGetResponseLatestDeploymentStageJSON contains the JSON metadata for
// the struct [PageProjectGetResponseLatestDeploymentStage]
type pageProjectGetResponseLatestDeploymentStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectGetResponseLatestDeploymentStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectGetResponseLatestDeploymentStageJSON) RawJSON() string {
	return r.raw
}

type PageProjectPurgeBuildCacheResponse = interface{}

type PageProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig         param.Field[PageProjectNewParamsBuildConfig]         `json:"build_config"`
	CanonicalDeployment param.Field[PageProjectNewParamsCanonicalDeployment] `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[PageProjectNewParamsDeploymentConfigs] `json:"deployment_configs"`
	LatestDeployment  param.Field[PageProjectNewParamsLatestDeployment]  `json:"latest_deployment"`
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

type PageProjectNewParamsCanonicalDeployment struct {
}

func (r PageProjectNewParamsCanonicalDeployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type PageProjectNewParamsCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata param.Field[PageProjectNewParamsCanonicalDeploymentDeploymentTriggerMetadata] `json:"metadata"`
}

func (r PageProjectNewParamsCanonicalDeploymentDeploymentTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type PageProjectNewParamsCanonicalDeploymentDeploymentTriggerMetadata struct {
}

func (r PageProjectNewParamsCanonicalDeploymentDeploymentTriggerMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type PageProjectNewParamsCanonicalDeploymentStage struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r PageProjectNewParamsCanonicalDeploymentStage) MarshalJSON() (data []byte, err error) {
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

type PageProjectNewParamsLatestDeployment struct {
}

func (r PageProjectNewParamsLatestDeployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type PageProjectNewParamsLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata param.Field[PageProjectNewParamsLatestDeploymentDeploymentTriggerMetadata] `json:"metadata"`
}

func (r PageProjectNewParamsLatestDeploymentDeploymentTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type PageProjectNewParamsLatestDeploymentDeploymentTriggerMetadata struct {
}

func (r PageProjectNewParamsLatestDeploymentDeploymentTriggerMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type PageProjectNewParamsLatestDeploymentStage struct {
	// The current build stage.
	Name param.Field[string] `json:"name"`
}

func (r PageProjectNewParamsLatestDeploymentStage) MarshalJSON() (data []byte, err error) {
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

func (r pageProjectNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   []PageProjectListResponse                 `json:"result,required"`
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

func (r pageProjectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   PageProjectGetResponse                   `json:"result,required"`
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

func (r pageProjectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r pageProjectGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
