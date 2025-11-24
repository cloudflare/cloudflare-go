// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// ProjectDeploymentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDeploymentService] method instead.
type ProjectDeploymentService struct {
	Options []option.RequestOption
	History *ProjectDeploymentHistoryService
}

// NewProjectDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectDeploymentService(opts ...option.RequestOption) (r *ProjectDeploymentService) {
	r = &ProjectDeploymentService{}
	r.Options = opts
	r.History = NewProjectDeploymentHistoryService(opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *ProjectDeploymentService) New(ctx context.Context, projectName string, params ProjectDeploymentNewParams, opts ...option.RequestOption) (res *ProjectDeploymentNewResponse, err error) {
	var env ProjectDeploymentNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) List(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ProjectDeploymentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) ListAutoPaging(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ProjectDeploymentListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, projectName, params, opts...))
}

// Delete a deployment.
func (r *ProjectDeploymentService) Delete(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *ProjectDeploymentDeleteResponse, err error) {
	var env ProjectDeploymentDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch information about a deployment.
func (r *ProjectDeploymentService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentGetParams, opts ...option.RequestOption) (res *ProjectDeploymentGetResponse, err error) {
	var env ProjectDeploymentGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry a previous deployment.
func (r *ProjectDeploymentService) Retry(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentRetryParams, opts ...option.RequestOption) (res *ProjectDeploymentRetryResponse, err error) {
	var env ProjectDeploymentRetryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *ProjectDeploymentService) Rollback(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *ProjectDeploymentRollbackResponse, err error) {
	var env ProjectDeploymentRollbackResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentNewResponse struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectDeploymentNewResponseBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentNewResponseDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentNewResponseEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectDeploymentNewResponseEnvironment `json:"environment,required"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped,required"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage,required"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id,required"`
	// Name of the project.
	ProjectName string `json:"project_name,required"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id,required"`
	// Configs for the project source control.
	Source ProjectDeploymentNewResponseSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                             `json:"uses_functions,nullable"`
	JSON          projectDeploymentNewResponseJSON `json:"-"`
}

// projectDeploymentNewResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentNewResponse]
type projectDeploymentNewResponseJSON struct {
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
	UsesFunctions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectDeploymentNewResponseBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                                      `json:"root_dir,nullable"`
	JSON    projectDeploymentNewResponseBuildConfigJSON `json:"-"`
}

// projectDeploymentNewResponseBuildConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentNewResponseBuildConfig]
type projectDeploymentNewResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentNewResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentNewResponseDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectDeploymentNewResponseDeploymentTriggerType `json:"type,required"`
	JSON projectDeploymentNewResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentNewResponseDeploymentTriggerJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseDeploymentTrigger]
type projectDeploymentNewResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentNewResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                    `json:"commit_message,required"`
	JSON          projectDeploymentNewResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentNewResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [ProjectDeploymentNewResponseDeploymentTriggerMetadata]
type projectDeploymentNewResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectDeploymentNewResponseDeploymentTriggerType string

const (
	ProjectDeploymentNewResponseDeploymentTriggerTypeGitHubPush ProjectDeploymentNewResponseDeploymentTriggerType = "github:push"
	ProjectDeploymentNewResponseDeploymentTriggerTypeADHoc      ProjectDeploymentNewResponseDeploymentTriggerType = "ad_hoc"
	ProjectDeploymentNewResponseDeploymentTriggerTypeDeployHook ProjectDeploymentNewResponseDeploymentTriggerType = "deploy_hook"
)

func (r ProjectDeploymentNewResponseDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseDeploymentTriggerTypeGitHubPush, ProjectDeploymentNewResponseDeploymentTriggerTypeADHoc, ProjectDeploymentNewResponseDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectDeploymentNewResponseEnvVar struct {
	Type ProjectDeploymentNewResponseEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                 `json:"value,required"`
	JSON  projectDeploymentNewResponseEnvVarJSON `json:"-"`
	union ProjectDeploymentNewResponseEnvVarsUnion
}

// projectDeploymentNewResponseEnvVarJSON contains the JSON metadata for the struct
// [ProjectDeploymentNewResponseEnvVar]
type projectDeploymentNewResponseEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentNewResponseEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentNewResponseEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentNewResponseEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentNewResponseEnvVarsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentNewResponseEnvVar) AsUnion() ProjectDeploymentNewResponseEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar] or
// [ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentNewResponseEnvVarsUnion interface {
	implementsProjectDeploymentNewResponseEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentNewResponseEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                      `json:"value,required"`
	JSON  projectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar]
type projectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentNewResponseEnvVar() {
}

type ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                       `json:"value,required"`
	JSON  projectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar]
type projectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentNewResponseEnvVar() {
}

type ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentNewResponseEnvVarsType string

const (
	ProjectDeploymentNewResponseEnvVarsTypePlainText  ProjectDeploymentNewResponseEnvVarsType = "plain_text"
	ProjectDeploymentNewResponseEnvVarsTypeSecretText ProjectDeploymentNewResponseEnvVarsType = "secret_text"
)

func (r ProjectDeploymentNewResponseEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvVarsTypePlainText, ProjectDeploymentNewResponseEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectDeploymentNewResponseEnvironment string

const (
	ProjectDeploymentNewResponseEnvironmentPreview    ProjectDeploymentNewResponseEnvironment = "preview"
	ProjectDeploymentNewResponseEnvironmentProduction ProjectDeploymentNewResponseEnvironment = "production"
)

func (r ProjectDeploymentNewResponseEnvironment) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvironmentPreview, ProjectDeploymentNewResponseEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectDeploymentNewResponseSource struct {
	Config ProjectDeploymentNewResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectDeploymentNewResponseSourceType `json:"type,required"`
	JSON projectDeploymentNewResponseSourceJSON `json:"-"`
}

// projectDeploymentNewResponseSourceJSON contains the JSON metadata for the struct
// [ProjectDeploymentNewResponseSource]
type projectDeploymentNewResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                       `json:"repo_id"`
	JSON   projectDeploymentNewResponseSourceConfigJSON `json:"-"`
}

// projectDeploymentNewResponseSourceConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentNewResponseSourceConfig]
type projectDeploymentNewResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSettingAll    ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSettingNone   ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSettingCustom ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSettingAll, ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSettingNone, ProjectDeploymentNewResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectDeploymentNewResponseSourceType string

const (
	ProjectDeploymentNewResponseSourceTypeGitHub ProjectDeploymentNewResponseSourceType = "github"
	ProjectDeploymentNewResponseSourceTypeGitlab ProjectDeploymentNewResponseSourceType = "gitlab"
)

func (r ProjectDeploymentNewResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseSourceTypeGitHub, ProjectDeploymentNewResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectDeploymentListResponse struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectDeploymentListResponseBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentListResponseDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentListResponseEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectDeploymentListResponseEnvironment `json:"environment,required"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped,required"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage,required"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id,required"`
	// Name of the project.
	ProjectName string `json:"project_name,required"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id,required"`
	// Configs for the project source control.
	Source ProjectDeploymentListResponseSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                              `json:"uses_functions,nullable"`
	JSON          projectDeploymentListResponseJSON `json:"-"`
}

// projectDeploymentListResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentListResponse]
type projectDeploymentListResponseJSON struct {
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
	UsesFunctions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectDeploymentListResponseBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                                       `json:"root_dir,nullable"`
	JSON    projectDeploymentListResponseBuildConfigJSON `json:"-"`
}

// projectDeploymentListResponseBuildConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentListResponseBuildConfig]
type projectDeploymentListResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentListResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentListResponseDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectDeploymentListResponseDeploymentTriggerType `json:"type,required"`
	JSON projectDeploymentListResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentListResponseDeploymentTriggerJSON contains the JSON metadata
// for the struct [ProjectDeploymentListResponseDeploymentTrigger]
type projectDeploymentListResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentListResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                     `json:"commit_message,required"`
	JSON          projectDeploymentListResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentListResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [ProjectDeploymentListResponseDeploymentTriggerMetadata]
type projectDeploymentListResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectDeploymentListResponseDeploymentTriggerType string

const (
	ProjectDeploymentListResponseDeploymentTriggerTypeGitHubPush ProjectDeploymentListResponseDeploymentTriggerType = "github:push"
	ProjectDeploymentListResponseDeploymentTriggerTypeADHoc      ProjectDeploymentListResponseDeploymentTriggerType = "ad_hoc"
	ProjectDeploymentListResponseDeploymentTriggerTypeDeployHook ProjectDeploymentListResponseDeploymentTriggerType = "deploy_hook"
)

func (r ProjectDeploymentListResponseDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseDeploymentTriggerTypeGitHubPush, ProjectDeploymentListResponseDeploymentTriggerTypeADHoc, ProjectDeploymentListResponseDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectDeploymentListResponseEnvVar struct {
	Type ProjectDeploymentListResponseEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                  `json:"value,required"`
	JSON  projectDeploymentListResponseEnvVarJSON `json:"-"`
	union ProjectDeploymentListResponseEnvVarsUnion
}

// projectDeploymentListResponseEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentListResponseEnvVar]
type projectDeploymentListResponseEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentListResponseEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentListResponseEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentListResponseEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentListResponseEnvVarsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentListResponseEnvVar) AsUnion() ProjectDeploymentListResponseEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar] or
// [ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentListResponseEnvVarsUnion interface {
	implementsProjectDeploymentListResponseEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentListResponseEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                       `json:"value,required"`
	JSON  projectDeploymentListResponseEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentListResponseEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar]
type projectDeploymentListResponseEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentListResponseEnvVar() {
}

type ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                        `json:"value,required"`
	JSON  projectDeploymentListResponseEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentListResponseEnvVarsPagesSecretTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar]
type projectDeploymentListResponseEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentListResponseEnvVar() {
}

type ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentListResponseEnvVarsType string

const (
	ProjectDeploymentListResponseEnvVarsTypePlainText  ProjectDeploymentListResponseEnvVarsType = "plain_text"
	ProjectDeploymentListResponseEnvVarsTypeSecretText ProjectDeploymentListResponseEnvVarsType = "secret_text"
)

func (r ProjectDeploymentListResponseEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseEnvVarsTypePlainText, ProjectDeploymentListResponseEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectDeploymentListResponseEnvironment string

const (
	ProjectDeploymentListResponseEnvironmentPreview    ProjectDeploymentListResponseEnvironment = "preview"
	ProjectDeploymentListResponseEnvironmentProduction ProjectDeploymentListResponseEnvironment = "production"
)

func (r ProjectDeploymentListResponseEnvironment) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseEnvironmentPreview, ProjectDeploymentListResponseEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectDeploymentListResponseSource struct {
	Config ProjectDeploymentListResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectDeploymentListResponseSourceType `json:"type,required"`
	JSON projectDeploymentListResponseSourceJSON `json:"-"`
}

// projectDeploymentListResponseSourceJSON contains the JSON metadata for the
// struct [ProjectDeploymentListResponseSource]
type projectDeploymentListResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentListResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectDeploymentListResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                        `json:"repo_id"`
	JSON   projectDeploymentListResponseSourceConfigJSON `json:"-"`
}

// projectDeploymentListResponseSourceConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentListResponseSourceConfig]
type projectDeploymentListResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectDeploymentListResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentListResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectDeploymentListResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectDeploymentListResponseSourceConfigPreviewDeploymentSettingAll    ProjectDeploymentListResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectDeploymentListResponseSourceConfigPreviewDeploymentSettingNone   ProjectDeploymentListResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectDeploymentListResponseSourceConfigPreviewDeploymentSettingCustom ProjectDeploymentListResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectDeploymentListResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseSourceConfigPreviewDeploymentSettingAll, ProjectDeploymentListResponseSourceConfigPreviewDeploymentSettingNone, ProjectDeploymentListResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectDeploymentListResponseSourceType string

const (
	ProjectDeploymentListResponseSourceTypeGitHub ProjectDeploymentListResponseSourceType = "github"
	ProjectDeploymentListResponseSourceTypeGitlab ProjectDeploymentListResponseSourceType = "gitlab"
)

func (r ProjectDeploymentListResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectDeploymentListResponseSourceTypeGitHub, ProjectDeploymentListResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectDeploymentDeleteResponse = interface{}

type ProjectDeploymentGetResponse struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectDeploymentGetResponseBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentGetResponseDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentGetResponseEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectDeploymentGetResponseEnvironment `json:"environment,required"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped,required"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage,required"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id,required"`
	// Name of the project.
	ProjectName string `json:"project_name,required"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id,required"`
	// Configs for the project source control.
	Source ProjectDeploymentGetResponseSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                             `json:"uses_functions,nullable"`
	JSON          projectDeploymentGetResponseJSON `json:"-"`
}

// projectDeploymentGetResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentGetResponse]
type projectDeploymentGetResponseJSON struct {
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
	UsesFunctions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectDeploymentGetResponseBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                                      `json:"root_dir,nullable"`
	JSON    projectDeploymentGetResponseBuildConfigJSON `json:"-"`
}

// projectDeploymentGetResponseBuildConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentGetResponseBuildConfig]
type projectDeploymentGetResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentGetResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentGetResponseDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectDeploymentGetResponseDeploymentTriggerType `json:"type,required"`
	JSON projectDeploymentGetResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentGetResponseDeploymentTriggerJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseDeploymentTrigger]
type projectDeploymentGetResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentGetResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                    `json:"commit_message,required"`
	JSON          projectDeploymentGetResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentGetResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct [ProjectDeploymentGetResponseDeploymentTriggerMetadata]
type projectDeploymentGetResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectDeploymentGetResponseDeploymentTriggerType string

const (
	ProjectDeploymentGetResponseDeploymentTriggerTypeGitHubPush ProjectDeploymentGetResponseDeploymentTriggerType = "github:push"
	ProjectDeploymentGetResponseDeploymentTriggerTypeADHoc      ProjectDeploymentGetResponseDeploymentTriggerType = "ad_hoc"
	ProjectDeploymentGetResponseDeploymentTriggerTypeDeployHook ProjectDeploymentGetResponseDeploymentTriggerType = "deploy_hook"
)

func (r ProjectDeploymentGetResponseDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseDeploymentTriggerTypeGitHubPush, ProjectDeploymentGetResponseDeploymentTriggerTypeADHoc, ProjectDeploymentGetResponseDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectDeploymentGetResponseEnvVar struct {
	Type ProjectDeploymentGetResponseEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                 `json:"value,required"`
	JSON  projectDeploymentGetResponseEnvVarJSON `json:"-"`
	union ProjectDeploymentGetResponseEnvVarsUnion
}

// projectDeploymentGetResponseEnvVarJSON contains the JSON metadata for the struct
// [ProjectDeploymentGetResponseEnvVar]
type projectDeploymentGetResponseEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentGetResponseEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentGetResponseEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentGetResponseEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentGetResponseEnvVarsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentGetResponseEnvVar) AsUnion() ProjectDeploymentGetResponseEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar] or
// [ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentGetResponseEnvVarsUnion interface {
	implementsProjectDeploymentGetResponseEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentGetResponseEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                      `json:"value,required"`
	JSON  projectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar]
type projectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentGetResponseEnvVar() {
}

type ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                       `json:"value,required"`
	JSON  projectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar]
type projectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentGetResponseEnvVar() {
}

type ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentGetResponseEnvVarsType string

const (
	ProjectDeploymentGetResponseEnvVarsTypePlainText  ProjectDeploymentGetResponseEnvVarsType = "plain_text"
	ProjectDeploymentGetResponseEnvVarsTypeSecretText ProjectDeploymentGetResponseEnvVarsType = "secret_text"
)

func (r ProjectDeploymentGetResponseEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvVarsTypePlainText, ProjectDeploymentGetResponseEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectDeploymentGetResponseEnvironment string

const (
	ProjectDeploymentGetResponseEnvironmentPreview    ProjectDeploymentGetResponseEnvironment = "preview"
	ProjectDeploymentGetResponseEnvironmentProduction ProjectDeploymentGetResponseEnvironment = "production"
)

func (r ProjectDeploymentGetResponseEnvironment) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvironmentPreview, ProjectDeploymentGetResponseEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectDeploymentGetResponseSource struct {
	Config ProjectDeploymentGetResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectDeploymentGetResponseSourceType `json:"type,required"`
	JSON projectDeploymentGetResponseSourceJSON `json:"-"`
}

// projectDeploymentGetResponseSourceJSON contains the JSON metadata for the struct
// [ProjectDeploymentGetResponseSource]
type projectDeploymentGetResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                       `json:"repo_id"`
	JSON   projectDeploymentGetResponseSourceConfigJSON `json:"-"`
}

// projectDeploymentGetResponseSourceConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentGetResponseSourceConfig]
type projectDeploymentGetResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSettingAll    ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSettingNone   ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSettingCustom ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSettingAll, ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSettingNone, ProjectDeploymentGetResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectDeploymentGetResponseSourceType string

const (
	ProjectDeploymentGetResponseSourceTypeGitHub ProjectDeploymentGetResponseSourceType = "github"
	ProjectDeploymentGetResponseSourceTypeGitlab ProjectDeploymentGetResponseSourceType = "gitlab"
)

func (r ProjectDeploymentGetResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseSourceTypeGitHub, ProjectDeploymentGetResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectDeploymentRetryResponse struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectDeploymentRetryResponseBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentRetryResponseDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentRetryResponseEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectDeploymentRetryResponseEnvironment `json:"environment,required"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped,required"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage,required"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id,required"`
	// Name of the project.
	ProjectName string `json:"project_name,required"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id,required"`
	// Configs for the project source control.
	Source ProjectDeploymentRetryResponseSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                               `json:"uses_functions,nullable"`
	JSON          projectDeploymentRetryResponseJSON `json:"-"`
}

// projectDeploymentRetryResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentRetryResponse]
type projectDeploymentRetryResponseJSON struct {
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
	UsesFunctions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectDeploymentRetryResponseBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                                        `json:"root_dir,nullable"`
	JSON    projectDeploymentRetryResponseBuildConfigJSON `json:"-"`
}

// projectDeploymentRetryResponseBuildConfigJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseBuildConfig]
type projectDeploymentRetryResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentRetryResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentRetryResponseDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectDeploymentRetryResponseDeploymentTriggerType `json:"type,required"`
	JSON projectDeploymentRetryResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentRetryResponseDeploymentTriggerJSON contains the JSON metadata
// for the struct [ProjectDeploymentRetryResponseDeploymentTrigger]
type projectDeploymentRetryResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentRetryResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                      `json:"commit_message,required"`
	JSON          projectDeploymentRetryResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentRetryResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRetryResponseDeploymentTriggerMetadata]
type projectDeploymentRetryResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectDeploymentRetryResponseDeploymentTriggerType string

const (
	ProjectDeploymentRetryResponseDeploymentTriggerTypeGitHubPush ProjectDeploymentRetryResponseDeploymentTriggerType = "github:push"
	ProjectDeploymentRetryResponseDeploymentTriggerTypeADHoc      ProjectDeploymentRetryResponseDeploymentTriggerType = "ad_hoc"
	ProjectDeploymentRetryResponseDeploymentTriggerTypeDeployHook ProjectDeploymentRetryResponseDeploymentTriggerType = "deploy_hook"
)

func (r ProjectDeploymentRetryResponseDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseDeploymentTriggerTypeGitHubPush, ProjectDeploymentRetryResponseDeploymentTriggerTypeADHoc, ProjectDeploymentRetryResponseDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectDeploymentRetryResponseEnvVar struct {
	Type ProjectDeploymentRetryResponseEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                   `json:"value,required"`
	JSON  projectDeploymentRetryResponseEnvVarJSON `json:"-"`
	union ProjectDeploymentRetryResponseEnvVarsUnion
}

// projectDeploymentRetryResponseEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseEnvVar]
type projectDeploymentRetryResponseEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentRetryResponseEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentRetryResponseEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentRetryResponseEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentRetryResponseEnvVarsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentRetryResponseEnvVar) AsUnion() ProjectDeploymentRetryResponseEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar]
// or [ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentRetryResponseEnvVarsUnion interface {
	implementsProjectDeploymentRetryResponseEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentRetryResponseEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                        `json:"value,required"`
	JSON  projectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar]
type projectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentRetryResponseEnvVar() {
}

type ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                         `json:"value,required"`
	JSON  projectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar]
type projectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentRetryResponseEnvVar() {
}

type ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentRetryResponseEnvVarsType string

const (
	ProjectDeploymentRetryResponseEnvVarsTypePlainText  ProjectDeploymentRetryResponseEnvVarsType = "plain_text"
	ProjectDeploymentRetryResponseEnvVarsTypeSecretText ProjectDeploymentRetryResponseEnvVarsType = "secret_text"
)

func (r ProjectDeploymentRetryResponseEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvVarsTypePlainText, ProjectDeploymentRetryResponseEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectDeploymentRetryResponseEnvironment string

const (
	ProjectDeploymentRetryResponseEnvironmentPreview    ProjectDeploymentRetryResponseEnvironment = "preview"
	ProjectDeploymentRetryResponseEnvironmentProduction ProjectDeploymentRetryResponseEnvironment = "production"
)

func (r ProjectDeploymentRetryResponseEnvironment) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvironmentPreview, ProjectDeploymentRetryResponseEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectDeploymentRetryResponseSource struct {
	Config ProjectDeploymentRetryResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectDeploymentRetryResponseSourceType `json:"type,required"`
	JSON projectDeploymentRetryResponseSourceJSON `json:"-"`
}

// projectDeploymentRetryResponseSourceJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseSource]
type projectDeploymentRetryResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                         `json:"repo_id"`
	JSON   projectDeploymentRetryResponseSourceConfigJSON `json:"-"`
}

// projectDeploymentRetryResponseSourceConfigJSON contains the JSON metadata for
// the struct [ProjectDeploymentRetryResponseSourceConfig]
type projectDeploymentRetryResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSettingAll    ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSettingNone   ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSettingCustom ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSettingAll, ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSettingNone, ProjectDeploymentRetryResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectDeploymentRetryResponseSourceType string

const (
	ProjectDeploymentRetryResponseSourceTypeGitHub ProjectDeploymentRetryResponseSourceType = "github"
	ProjectDeploymentRetryResponseSourceTypeGitlab ProjectDeploymentRetryResponseSourceType = "gitlab"
)

func (r ProjectDeploymentRetryResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseSourceTypeGitHub, ProjectDeploymentRetryResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectDeploymentRollbackResponse struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectDeploymentRollbackResponseBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectDeploymentRollbackResponseDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentRollbackResponseEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectDeploymentRollbackResponseEnvironment `json:"environment,required"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped,required"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage,required"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id,required"`
	// Name of the project.
	ProjectName string `json:"project_name,required"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id,required"`
	// Configs for the project source control.
	Source ProjectDeploymentRollbackResponseSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                  `json:"uses_functions,nullable"`
	JSON          projectDeploymentRollbackResponseJSON `json:"-"`
}

// projectDeploymentRollbackResponseJSON contains the JSON metadata for the struct
// [ProjectDeploymentRollbackResponse]
type projectDeploymentRollbackResponseJSON struct {
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
	UsesFunctions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectDeploymentRollbackResponseBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                                           `json:"root_dir,nullable"`
	JSON    projectDeploymentRollbackResponseBuildConfigJSON `json:"-"`
}

// projectDeploymentRollbackResponseBuildConfigJSON contains the JSON metadata for
// the struct [ProjectDeploymentRollbackResponseBuildConfig]
type projectDeploymentRollbackResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectDeploymentRollbackResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectDeploymentRollbackResponseDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectDeploymentRollbackResponseDeploymentTriggerType `json:"type,required"`
	JSON projectDeploymentRollbackResponseDeploymentTriggerJSON `json:"-"`
}

// projectDeploymentRollbackResponseDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectDeploymentRollbackResponseDeploymentTrigger]
type projectDeploymentRollbackResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectDeploymentRollbackResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                         `json:"commit_message,required"`
	JSON          projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRollbackResponseDeploymentTriggerMetadata]
type projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectDeploymentRollbackResponseDeploymentTriggerType string

const (
	ProjectDeploymentRollbackResponseDeploymentTriggerTypeGitHubPush ProjectDeploymentRollbackResponseDeploymentTriggerType = "github:push"
	ProjectDeploymentRollbackResponseDeploymentTriggerTypeADHoc      ProjectDeploymentRollbackResponseDeploymentTriggerType = "ad_hoc"
	ProjectDeploymentRollbackResponseDeploymentTriggerTypeDeployHook ProjectDeploymentRollbackResponseDeploymentTriggerType = "deploy_hook"
)

func (r ProjectDeploymentRollbackResponseDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseDeploymentTriggerTypeGitHubPush, ProjectDeploymentRollbackResponseDeploymentTriggerTypeADHoc, ProjectDeploymentRollbackResponseDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectDeploymentRollbackResponseEnvVar struct {
	Type ProjectDeploymentRollbackResponseEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                      `json:"value,required"`
	JSON  projectDeploymentRollbackResponseEnvVarJSON `json:"-"`
	union ProjectDeploymentRollbackResponseEnvVarsUnion
}

// projectDeploymentRollbackResponseEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentRollbackResponseEnvVar]
type projectDeploymentRollbackResponseEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentRollbackResponseEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentRollbackResponseEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentRollbackResponseEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentRollbackResponseEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentRollbackResponseEnvVar) AsUnion() ProjectDeploymentRollbackResponseEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar] or
// [ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentRollbackResponseEnvVarsUnion interface {
	implementsProjectDeploymentRollbackResponseEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentRollbackResponseEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                           `json:"value,required"`
	JSON  projectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar]
type projectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentRollbackResponseEnvVar() {
}

type ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                            `json:"value,required"`
	JSON  projectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar]
type projectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentRollbackResponseEnvVar() {
}

type ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentRollbackResponseEnvVarsType string

const (
	ProjectDeploymentRollbackResponseEnvVarsTypePlainText  ProjectDeploymentRollbackResponseEnvVarsType = "plain_text"
	ProjectDeploymentRollbackResponseEnvVarsTypeSecretText ProjectDeploymentRollbackResponseEnvVarsType = "secret_text"
)

func (r ProjectDeploymentRollbackResponseEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvVarsTypePlainText, ProjectDeploymentRollbackResponseEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectDeploymentRollbackResponseEnvironment string

const (
	ProjectDeploymentRollbackResponseEnvironmentPreview    ProjectDeploymentRollbackResponseEnvironment = "preview"
	ProjectDeploymentRollbackResponseEnvironmentProduction ProjectDeploymentRollbackResponseEnvironment = "production"
)

func (r ProjectDeploymentRollbackResponseEnvironment) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvironmentPreview, ProjectDeploymentRollbackResponseEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectDeploymentRollbackResponseSource struct {
	Config ProjectDeploymentRollbackResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectDeploymentRollbackResponseSourceType `json:"type,required"`
	JSON projectDeploymentRollbackResponseSourceJSON `json:"-"`
}

// projectDeploymentRollbackResponseSourceJSON contains the JSON metadata for the
// struct [ProjectDeploymentRollbackResponseSource]
type projectDeploymentRollbackResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                            `json:"repo_id"`
	JSON   projectDeploymentRollbackResponseSourceConfigJSON `json:"-"`
}

// projectDeploymentRollbackResponseSourceConfigJSON contains the JSON metadata for
// the struct [ProjectDeploymentRollbackResponseSourceConfig]
type projectDeploymentRollbackResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSettingAll    ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSettingNone   ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSettingCustom ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSettingAll, ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSettingNone, ProjectDeploymentRollbackResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectDeploymentRollbackResponseSourceType string

const (
	ProjectDeploymentRollbackResponseSourceTypeGitHub ProjectDeploymentRollbackResponseSourceType = "github"
	ProjectDeploymentRollbackResponseSourceTypeGitlab ProjectDeploymentRollbackResponseSourceType = "gitlab"
)

func (r ProjectDeploymentRollbackResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseSourceTypeGitHub, ProjectDeploymentRollbackResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectDeploymentNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Headers configuration file for the deployment.
	Headers param.Field[io.Reader] `json:"_headers" format:"binary"`
	// Redirects configuration file for the deployment.
	Redirects param.Field[io.Reader] `json:"_redirects" format:"binary"`
	// Routes configuration file defining routing rules.
	RoutesJson param.Field[io.Reader] `json:"_routes.json" format:"binary"`
	// Worker bundle file in multipart/form-data format. Mutually exclusive with
	// `_worker.js`. Cannot specify both `_worker.js` and `_worker.bundle` in the same
	// request. Maximum size: 25 MiB.
	WorkerBundle param.Field[io.Reader] `json:"_worker.bundle" format:"binary"`
	// Worker JavaScript file. Mutually exclusive with `_worker.bundle`. Cannot specify
	// both `_worker.js` and `_worker.bundle` in the same request.
	WorkerJS param.Field[io.Reader] `json:"_worker.js" format:"binary"`
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
	// Boolean string indicating if the working directory has uncommitted changes.
	CommitDirty param.Field[ProjectDeploymentNewParamsCommitDirty] `json:"commit_dirty"`
	// Git commit SHA associated with this deployment.
	CommitHash param.Field[string] `json:"commit_hash"`
	// Git commit message associated with this deployment.
	CommitMessage param.Field[string] `json:"commit_message"`
	// Functions routing configuration file.
	FunctionsFilepathRoutingConfigJson param.Field[io.Reader] `json:"functions-filepath-routing-config.json" format:"binary"`
	// JSON string containing a manifest of files to deploy. Maps file paths to their
	// content hashes. Required for direct upload deployments. Maximum 20,000 entries.
	Manifest param.Field[string] `json:"manifest"`
	// The build output directory path.
	PagesBuildOutputDir param.Field[string] `json:"pages_build_output_dir"`
	// Hash of the Wrangler configuration file used for this deployment.
	WranglerConfigHash param.Field[string] `json:"wrangler_config_hash"`
}

func (r ProjectDeploymentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Boolean string indicating if the working directory has uncommitted changes.
type ProjectDeploymentNewParamsCommitDirty string

const (
	ProjectDeploymentNewParamsCommitDirtyTrue  ProjectDeploymentNewParamsCommitDirty = "true"
	ProjectDeploymentNewParamsCommitDirtyFalse ProjectDeploymentNewParamsCommitDirty = "false"
)

func (r ProjectDeploymentNewParamsCommitDirty) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewParamsCommitDirtyTrue, ProjectDeploymentNewParamsCommitDirtyFalse:
		return true
	}
	return false
}

type ProjectDeploymentNewResponseEnvelope struct {
	Errors   []ProjectDeploymentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentNewResponseEnvelope]
type projectDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ProjectDeploymentNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseEnvelopeErrors]
type projectDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    projectDeploymentNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ProjectDeploymentNewResponseEnvelopeErrorsSource]
type projectDeploymentNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectDeploymentNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentNewResponseEnvelopeMessages]
type projectDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectDeploymentNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentNewResponseEnvelopeMessagesSource]
type projectDeploymentNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentNewResponseEnvelopeSuccess bool

const (
	ProjectDeploymentNewResponseEnvelopeSuccessTrue ProjectDeploymentNewResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// What type of deployments to fetch.
	Env param.Field[ProjectDeploymentListParamsEnv] `query:"env"`
	// Which page of deployments to fetch.
	Page param.Field[int64] `query:"page"`
	// How many deployments to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ProjectDeploymentListParams]'s query parameters as
// `url.Values`.
func (r ProjectDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// What type of deployments to fetch.
type ProjectDeploymentListParamsEnv string

const (
	ProjectDeploymentListParamsEnvProduction ProjectDeploymentListParamsEnv = "production"
	ProjectDeploymentListParamsEnvPreview    ProjectDeploymentListParamsEnv = "preview"
)

func (r ProjectDeploymentListParamsEnv) IsKnown() bool {
	switch r {
	case ProjectDeploymentListParamsEnvProduction, ProjectDeploymentListParamsEnvPreview:
		return true
	}
	return false
}

type ProjectDeploymentDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentDeleteResponseEnvelope struct {
	Errors   []ProjectDeploymentDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ProjectDeploymentDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentDeleteResponseEnvelope]
type projectDeploymentDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ProjectDeploymentDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentDeleteResponseEnvelopeErrors]
type projectDeploymentDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentDeleteResponseEnvelopeErrorsSource]
type projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ProjectDeploymentDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentDeleteResponseEnvelopeMessages]
type projectDeploymentDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentDeleteResponseEnvelopeMessagesSource]
type projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeploymentDeleteResponseEnvelopeSuccessTrue ProjectDeploymentDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentGetResponseEnvelope struct {
	Errors   []ProjectDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentGetResponseEnvelope]
type projectDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ProjectDeploymentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseEnvelopeErrors]
type projectDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    projectDeploymentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ProjectDeploymentGetResponseEnvelopeErrorsSource]
type projectDeploymentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectDeploymentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDeploymentGetResponseEnvelopeMessages]
type projectDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectDeploymentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentGetResponseEnvelopeMessagesSource]
type projectDeploymentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentGetResponseEnvelopeSuccessTrue ProjectDeploymentGetResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentRetryParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentRetryResponseEnvelope struct {
	Errors   []ProjectDeploymentRetryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRetryResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentRetryResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentRetryResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentRetryResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseEnvelope]
type projectDeploymentRetryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectDeploymentRetryResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentRetryResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectDeploymentRetryResponseEnvelopeErrors]
type projectDeploymentRetryResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectDeploymentRetryResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentRetryResponseEnvelopeErrorsSource]
type projectDeploymentRetryResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ProjectDeploymentRetryResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentRetryResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentRetryResponseEnvelopeMessages]
type projectDeploymentRetryResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRetryResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    projectDeploymentRetryResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentRetryResponseEnvelopeMessagesSource]
type projectDeploymentRetryResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentRetryResponseEnvelopeSuccess bool

const (
	ProjectDeploymentRetryResponseEnvelopeSuccessTrue ProjectDeploymentRetryResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentRetryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentRollbackParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentRollbackResponseEnvelope struct {
	Errors   []ProjectDeploymentRollbackResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentRollbackResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentRollbackResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDeploymentRollbackResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentRollbackResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentRollbackResponseEnvelope]
type projectDeploymentRollbackResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ProjectDeploymentRollbackResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeploymentRollbackResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ProjectDeploymentRollbackResponseEnvelopeErrors]
type projectDeploymentRollbackResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectDeploymentRollbackResponseEnvelopeErrorsSource]
type projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           ProjectDeploymentRollbackResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeploymentRollbackResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectDeploymentRollbackResponseEnvelopeMessages]
type projectDeploymentRollbackResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentRollbackResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentRollbackResponseEnvelopeMessagesSource]
type projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeploymentRollbackResponseEnvelopeSuccess bool

const (
	ProjectDeploymentRollbackResponseEnvelopeSuccessTrue ProjectDeploymentRollbackResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentRollbackResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
