// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// ApplicationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationService] method instead.
type ApplicationService struct {
	Options   []option.RequestOption
	Instances *ApplicationInstanceService
	Rollouts  *ApplicationRolloutService
	Versions  *ApplicationVersionService
}

// NewApplicationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApplicationService(opts ...option.RequestOption) (r *ApplicationService) {
	r = &ApplicationService{}
	r.Options = opts
	r.Instances = NewApplicationInstanceService(opts...)
	r.Rollouts = NewApplicationRolloutService(opts...)
	r.Versions = NewApplicationVersionService(opts...)
	return
}

// Create a Containers application.
//
// Use `scheduling_policy: "default"` for a scheduler-backed application. The
// Containers scheduler maintains the requested instance count and manages
// deployment configuration, placement, scaling, versions, and rollouts.
//
// Use `scheduling_policy: "durable_object"` for a Durable Object-managed
// application. Each Durable Object creates and manages the lifecycle of its
// container instance. Supply `name`, `scheduling_policy`, and `durable_objects`,
// with optional `configuration` and optional top-level `observability` settings.
// Deployment configuration, scaling, constraints, versions, and rollouts do not
// apply.
func (r *ApplicationService) New(ctx context.Context, params ApplicationNewParams, opts ...option.RequestOption) (res *ApplicationNewResponse, err error) {
	var env ApplicationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all the applications that are associated with your account.
func (r *ApplicationService) List(ctx context.Context, params ApplicationListParams, opts ...option.RequestOption) (res *pagination.PageTokenPagination[ApplicationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications", params.AccountID)
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

// Lists all the applications that are associated with your account.
func (r *ApplicationService) ListAutoPaging(ctx context.Context, params ApplicationListParams, opts ...option.RequestOption) *pagination.PageTokenPaginationAutoPager[ApplicationListResponse] {
	return pagination.NewPageTokenPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a single application by id.
func (r *ApplicationService) Delete(ctx context.Context, applicationID string, body ApplicationDeleteParams, opts ...option.RequestOption) (res *ApplicationDeleteResponse, err error) {
	var env ApplicationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s", body.AccountID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Modifies a single application by id. Durable Object-managed application settings
// are published to runtime metadata without creating deployments or rollouts.
// Top-level `observability` for these applications supports only `logs.enabled`.
// The supported fields depend on the existing application's scheduling policy. For
// scheduler-backed applications, changes that replace instance deployment
// configuration, including the image, require a rollout.
func (r *ApplicationService) Edit(ctx context.Context, applicationID string, params ApplicationEditParams, opts ...option.RequestOption) (res *ApplicationEditResponse, err error) {
	var env ApplicationEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s", params.AccountID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a single application by id.
func (r *ApplicationService) Get(ctx context.Context, applicationID string, query ApplicationGetParams, opts ...option.RequestOption) (res *ApplicationGetResponse, err error) {
	var env ApplicationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s", query.AccountID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// The public Containers API returns an application.
type ApplicationNewResponse struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationNewResponseSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string `json:"active_rollout_id"`
	// This field can have the runtime type of
	// [ApplicationNewResponseCcScheduledApplicationConfiguration],
	// [ApplicationNewResponseCcDurableObjectApplicationConfiguration].
	Configuration interface{} `json:"configuration"`
	// This field can have the runtime type of
	// [ApplicationNewResponseCcScheduledApplicationConstraints].
	Constraints interface{} `json:"constraints"`
	// This field can have the runtime type of
	// [ApplicationNewResponseCcScheduledApplicationDurableObjects],
	// [ApplicationNewResponseCcDurableObjectApplicationDurableObjects].
	DurableObjects interface{} `json:"durable_objects"`
	// This field can have the runtime type of
	// [ApplicationNewResponseCcScheduledApplicationHealth],
	// [ApplicationNewResponseCcDurableObjectApplicationHealth].
	Health interface{} `json:"health"`
	// Number of deployments to create.
	Instances int64 `json:"instances"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// This field can have the runtime type of
	// [ApplicationNewResponseCcScheduledApplicationObservability],
	// [ApplicationNewResponseCcDurableObjectApplicationObservability].
	Observability interface{} `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                      `json:"rollout_active_grace_period"`
	Version                  int64                      `json:"version"`
	JSON                     applicationNewResponseJSON `json:"-"`
	union                    ApplicationNewResponseUnion
}

// applicationNewResponseJSON contains the JSON metadata for the struct
// [ApplicationNewResponse]
type applicationNewResponseJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	CreatedAt                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	ActiveRolloutID          apijson.Field
	Configuration            apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	Instances                apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	Version                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r applicationNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ApplicationNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ApplicationNewResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ApplicationNewResponseCcScheduledApplication],
// [ApplicationNewResponseCcDurableObjectApplication].
func (r ApplicationNewResponse) AsUnion() ApplicationNewResponseUnion {
	return r.union
}

// The public Containers API returns an application.
//
// Union satisfied by [ApplicationNewResponseCcScheduledApplication] or
// [ApplicationNewResponseCcDurableObjectApplication].
type ApplicationNewResponseUnion interface {
	implementsApplicationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationNewResponseCcScheduledApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationNewResponseCcDurableObjectApplication{}),
		},
	)
}

// Describes an application and the parameters that govern how it places its
// instances.
type ApplicationNewResponseCcScheduledApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// User-specified container configuration.
	Configuration ApplicationNewResponseCcScheduledApplicationConfiguration `json:"configuration" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Number of deployments to create.
	Instances int64 `json:"instances" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationNewResponseCcScheduledApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	Version   int64  `json:"version" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string                                                  `json:"active_rollout_id"`
	Constraints     ApplicationNewResponseCcScheduledApplicationConstraints `json:"constraints"`
	// Durable object configuration stored on and returned from a Cloudchamber
	// application.
	DurableObjects ApplicationNewResponseCcScheduledApplicationDurableObjects `json:"durable_objects"`
	Health         ApplicationNewResponseCcScheduledApplicationHealth         `json:"health"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// Top-level observability settings for the application. This field is mutually
	// exclusive with configuration.observability.
	Observability ApplicationNewResponseCcScheduledApplicationObservability `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                                            `json:"rollout_active_grace_period"`
	JSON                     applicationNewResponseCcScheduledApplicationJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationJSON contains the JSON metadata for
// the struct [ApplicationNewResponseCcScheduledApplication]
type applicationNewResponseCcScheduledApplicationJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	Configuration            apijson.Field
	CreatedAt                apijson.Field
	Instances                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	Version                  apijson.Field
	ActiveRolloutID          apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationNewResponseCcScheduledApplication) implementsApplicationNewResponse() {}

// User-specified container configuration.
type ApplicationNewResponseCcScheduledApplicationConfiguration struct {
	// Image url.
	Image          string                                                                   `json:"image" api:"required"`
	AuthorizedKeys []ApplicationNewResponseCcScheduledApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command []string `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint []string `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables []ApplicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariable `json:"environment_variables"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationNewResponseCcScheduledApplicationConfigurationObservability `json:"observability"`
	JSON          applicationNewResponseCcScheduledApplicationConfigurationJSON          `json:"-"`
}

// applicationNewResponseCcScheduledApplicationConfigurationJSON contains the JSON
// metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationConfiguration]
type applicationNewResponseCcScheduledApplicationConfigurationJSON struct {
	Image                apijson.Field
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationNewResponseCcScheduledApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                     `json:"name"`
	JSON applicationNewResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationConfigurationAuthorizedKey]
type applicationNewResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                                           `json:"value" api:"required"`
	JSON  applicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariable]
type applicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType string

const (
	ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeLite      ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType = "lite"
	ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeBasic     ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType = "basic"
	ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard1 ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType = "standard-1"
	ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard2 ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType = "standard-2"
	ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard3 ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType = "standard-3"
	ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard4 ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType = "standard-4"
)

func (r ApplicationNewResponseCcScheduledApplicationConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeLite, ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeBasic, ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard1, ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard2, ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard3, ApplicationNewResponseCcScheduledApplicationConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationNewResponseCcScheduledApplicationConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationNewResponseCcScheduledApplicationConfigurationObservabilityLogs `json:"logs"`
	JSON applicationNewResponseCcScheduledApplicationConfigurationObservabilityJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationConfigurationObservabilityJSON
// contains the JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationConfigurationObservability]
type applicationNewResponseCcScheduledApplicationConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationNewResponseCcScheduledApplicationConfigurationObservabilityLogs struct {
	Enabled bool                                                                           `json:"enabled"`
	JSON    applicationNewResponseCcScheduledApplicationConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationConfigurationObservabilityLogsJSON
// contains the JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationConfigurationObservabilityLogs]
type applicationNewResponseCcScheduledApplicationConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationNewResponseCcScheduledApplicationSchedulingPolicy string

const (
	ApplicationNewResponseCcScheduledApplicationSchedulingPolicyDefault       ApplicationNewResponseCcScheduledApplicationSchedulingPolicy = "default"
	ApplicationNewResponseCcScheduledApplicationSchedulingPolicyDurableObject ApplicationNewResponseCcScheduledApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationNewResponseCcScheduledApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcScheduledApplicationSchedulingPolicyDefault, ApplicationNewResponseCcScheduledApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

type ApplicationNewResponseCcScheduledApplicationConstraints struct {
	// Restricts placement to datacenters in the selected jurisdiction. Choose "eu",
	// "fedramp", or "us". When combined with regions, EU supports EEUR and WEUR while
	// FedRAMP and US support ENAM and WNAM.
	Jurisdiction string                                                      `json:"jurisdiction"`
	Regions      []string                                                    `json:"regions"`
	JSON         applicationNewResponseCcScheduledApplicationConstraintsJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationConstraintsJSON contains the JSON
// metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationConstraints]
type applicationNewResponseCcScheduledApplicationConstraintsJSON struct {
	Jurisdiction apijson.Field
	Regions      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationConstraintsJSON) RawJSON() string {
	return r.raw
}

// Durable object configuration stored on and returned from a Cloudchamber
// application.
type ApplicationNewResponseCcScheduledApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                         `json:"namespace_id" api:"required"`
	JSON        applicationNewResponseCcScheduledApplicationDurableObjectsJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationDurableObjectsJSON contains the JSON
// metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationDurableObjects]
type applicationNewResponseCcScheduledApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

type ApplicationNewResponseCcScheduledApplicationHealth struct {
	Errors []ApplicationNewResponseCcScheduledApplicationHealthError `json:"errors" api:"required"`
	// Shows a count of application instance states.
	Instances ApplicationNewResponseCcScheduledApplicationHealthInstances `json:"instances" api:"required"`
	// High-level health assessment. Only populated for "new_instances" strategy. Based
	// on a sample of target-version instances rather than a full count.
	//
	// - "pending": Zero target-version instances exist yet.
	// - "healthy": Every sampled target-version instance reports running or active.
	// - "degraded": Some sampled instances remain starting or scheduling.
	// - "unhealthy": One or more sampled instances have failed.
	Summary ApplicationNewResponseCcScheduledApplicationHealthSummary `json:"summary"`
	JSON    applicationNewResponseCcScheduledApplicationHealthJSON    `json:"-"`
}

// applicationNewResponseCcScheduledApplicationHealthJSON contains the JSON
// metadata for the struct [ApplicationNewResponseCcScheduledApplicationHealth]
type applicationNewResponseCcScheduledApplicationHealthJSON struct {
	Errors      apijson.Field
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationHealthJSON) RawJSON() string {
	return r.raw
}

type ApplicationNewResponseCcScheduledApplicationHealthError struct {
	// An event within a Placement or a Job.
	Event ApplicationNewResponseCcScheduledApplicationHealthErrorsEvent `json:"event" api:"required"`
	// An instance ID represents an identifier of an instance configuration that
	// maintains an underlying placement.
	InstanceID string                                                      `json:"instance_id" api:"required"`
	JSON       applicationNewResponseCcScheduledApplicationHealthErrorJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationHealthErrorJSON contains the JSON
// metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationHealthError]
type applicationNewResponseCcScheduledApplicationHealthErrorJSON struct {
	Event       apijson.Field
	InstanceID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationHealthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationHealthErrorJSON) RawJSON() string {
	return r.raw
}

// An event within a Placement or a Job.
type ApplicationNewResponseCcScheduledApplicationHealthErrorsEvent struct {
	ID      string                 `json:"id" api:"required"`
	Details map[string]interface{} `json:"details" api:"required"`
	Message string                 `json:"message" api:"required"`
	// Name of the event that describes the kind event that happened.
	//
	// - SchedulerPlaced: It's the first event that creates a container placement. It
	//   happens when the Containers runtime was able to retrieve deployment resources
	//   and start verifying everything is correct.
	// - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
	//   container.
	// - VMStarted: It's sent when the Containers runtime starts the VM. The container
	//   might remain unhealthy at this point.
	// - ImagePulled: It's sent when the Containers runtime pulls the image
	//   successfully.
	// - ImagePullError: It's sent when the Containers runtime is having issues pulling
	//   the image. The message and details have more information on what happened for
	//   debugging.
	// - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
	//   VM.
	// - VMStopping: It's sent when the scheduler is stopping the VM.
	// - VMStopped: It's sent when the VM finally exits.
	// - VMFailed: It's sent when the scheduling of the VM failed in the current
	//   location.
	// - RuntimeStartFailed: It's sent when the runtime hits an internal error.
	// - SSHStarted: It's sent when the container gains network connectivity and opens
	//   the SSH port. Containers only send this event when SSH keys exist.
	// - CheckUpdate: Sent when the status of a health or readiness check changes. This
	//   may also affect the health status of the placement.
	// - DurableObjectConnected: Sent when a durable object instance connects and gains
	//   control of the deployment. This event is only sent for durable object
	//   deployments. It is sent after VMStarted.
	// - ContainerStarted: It's sent when the container starts running.
	Name         ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName `json:"name" api:"required"`
	StatusChange map[string]interface{}                                            `json:"statusChange" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	Time string                                                            `json:"time" api:"required"`
	Type ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType `json:"type" api:"required"`
	JSON applicationNewResponseCcScheduledApplicationHealthErrorsEventJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationHealthErrorsEventJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationHealthErrorsEvent]
type applicationNewResponseCcScheduledApplicationHealthErrorsEventJSON struct {
	ID           apijson.Field
	Details      apijson.Field
	Message      apijson.Field
	Name         apijson.Field
	StatusChange apijson.Field
	Time         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationHealthErrorsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationHealthErrorsEventJSON) RawJSON() string {
	return r.raw
}

// Name of the event that describes the kind event that happened.
//
//   - SchedulerPlaced: It's the first event that creates a container placement. It
//     happens when the Containers runtime was able to retrieve deployment resources
//     and start verifying everything is correct.
//   - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
//     container.
//   - VMStarted: It's sent when the Containers runtime starts the VM. The container
//     might remain unhealthy at this point.
//   - ImagePulled: It's sent when the Containers runtime pulls the image
//     successfully.
//   - ImagePullError: It's sent when the Containers runtime is having issues pulling
//     the image. The message and details have more information on what happened for
//     debugging.
//   - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
//     VM.
//   - VMStopping: It's sent when the scheduler is stopping the VM.
//   - VMStopped: It's sent when the VM finally exits.
//   - VMFailed: It's sent when the scheduling of the VM failed in the current
//     location.
//   - RuntimeStartFailed: It's sent when the runtime hits an internal error.
//   - SSHStarted: It's sent when the container gains network connectivity and opens
//     the SSH port. Containers only send this event when SSH keys exist.
//   - CheckUpdate: Sent when the status of a health or readiness check changes. This
//     may also affect the health status of the placement.
//   - DurableObjectConnected: Sent when a durable object instance connects and gains
//     control of the deployment. This event is only sent for durable object
//     deployments. It is sent after VMStarted.
//   - ContainerStarted: It's sent when the container starts running.
type ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName string

const (
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced              ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "SchedulerPlaced"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned         ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssigned"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmStarted                    ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "VMStarted"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameImagePulled                  ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "ImagePulled"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameImagePullError               ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "ImagePullError"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart              ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "VMFailedToStart"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssignmentFailed"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmRunning                    ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "VMRunning"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmStopping                   ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "VMStopping"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmStopped                    ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "VMStopped"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmFailed                     ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "VMFailed"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed           ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "RuntimeStartFailed"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted                   ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "SSHStarted"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates         ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "ServiceHealthUpdates"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate                  ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "CheckUpdate"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected       ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "DurableObjectConnected"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted             ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName = "ContainerStarted"
)

func (r ApplicationNewResponseCcScheduledApplicationHealthErrorsEventName) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmStarted, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameImagePulled, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameImagePullError, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmRunning, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmStopping, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmStopped, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameVmFailed, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted:
		return true
	}
	return false
}

type ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType string

const (
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeInfo        ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType = "Info"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeError       ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType = "Error"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeWarn        ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType = "Warn"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeUserError   ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType = "UserError"
	ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeSystemError ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType = "SystemError"
)

func (r ApplicationNewResponseCcScheduledApplicationHealthErrorsEventType) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeInfo, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeError, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeWarn, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeUserError, ApplicationNewResponseCcScheduledApplicationHealthErrorsEventTypeSystemError:
		return true
	}
	return false
}

// Shows a count of application instance states.
type ApplicationNewResponseCcScheduledApplicationHealthInstances struct {
	// Number of instances whose runtime reports the container as running
	// (container_status = "running"). This is a subset of the placements that remain
	// up: an instance that is already bound to a Durable Object and serving traffic is
	// counted under "assigned" until its container_status catches up to "running", so
	// container_status can briefly lag Durable Object attachment under churn. To
	// estimate running, Durable-Object-bound instances, sum "active" + "assigned"
	// rather than reading "active" alone.
	Active int64 `json:"active" api:"required"`
	// Number of instances bound to a Durable Object with a running placement whose
	// container_status remains behind "running". These count as live, serving
	// instances; "active" + "assigned" approximates the running, Durable-Object-bound
	// count.
	Assigned int64                                                           `json:"assigned" api:"required"`
	JSON     applicationNewResponseCcScheduledApplicationHealthInstancesJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationHealthInstancesJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationHealthInstances]
type applicationNewResponseCcScheduledApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Assigned    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// High-level health assessment. Only populated for "new_instances" strategy. Based
// on a sample of target-version instances rather than a full count.
//
// - "pending": Zero target-version instances exist yet.
// - "healthy": Every sampled target-version instance reports running or active.
// - "degraded": Some sampled instances remain starting or scheduling.
// - "unhealthy": One or more sampled instances have failed.
type ApplicationNewResponseCcScheduledApplicationHealthSummary string

const (
	ApplicationNewResponseCcScheduledApplicationHealthSummaryHealthy   ApplicationNewResponseCcScheduledApplicationHealthSummary = "healthy"
	ApplicationNewResponseCcScheduledApplicationHealthSummaryDegraded  ApplicationNewResponseCcScheduledApplicationHealthSummary = "degraded"
	ApplicationNewResponseCcScheduledApplicationHealthSummaryUnhealthy ApplicationNewResponseCcScheduledApplicationHealthSummary = "unhealthy"
	ApplicationNewResponseCcScheduledApplicationHealthSummaryPending   ApplicationNewResponseCcScheduledApplicationHealthSummary = "pending"
)

func (r ApplicationNewResponseCcScheduledApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcScheduledApplicationHealthSummaryHealthy, ApplicationNewResponseCcScheduledApplicationHealthSummaryDegraded, ApplicationNewResponseCcScheduledApplicationHealthSummaryUnhealthy, ApplicationNewResponseCcScheduledApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Top-level observability settings for the application. This field is mutually
// exclusive with configuration.observability.
type ApplicationNewResponseCcScheduledApplicationObservability struct {
	// Observability logging settings.
	Logs ApplicationNewResponseCcScheduledApplicationObservabilityLogs `json:"logs"`
	JSON applicationNewResponseCcScheduledApplicationObservabilityJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationObservabilityJSON contains the JSON
// metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationObservability]
type applicationNewResponseCcScheduledApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationNewResponseCcScheduledApplicationObservabilityLogs struct {
	Enabled bool                                                              `json:"enabled"`
	JSON    applicationNewResponseCcScheduledApplicationObservabilityLogsJSON `json:"-"`
}

// applicationNewResponseCcScheduledApplicationObservabilityLogsJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcScheduledApplicationObservabilityLogs]
type applicationNewResponseCcScheduledApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcScheduledApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcScheduledApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Each Durable Object creates and manages the lifecycle of its container instance.
type ApplicationNewResponseCcDurableObjectApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Durable object configuration using a namespace ID.
	DurableObjects ApplicationNewResponseCcDurableObjectApplicationDurableObjects `json:"durable_objects" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// Selects a Durable Object-managed application. Each Durable Object creates and
	// manages the lifecycle of its container instance. Configure application-wide
	// observability settings here. Deployment configuration, scaling, placement
	// constraints, versions, and rollouts do not apply.
	SchedulingPolicy ApplicationNewResponseCcDurableObjectApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Application-wide settings for a Durable Object-managed application.
	Configuration ApplicationNewResponseCcDurableObjectApplicationConfiguration `json:"configuration"`
	// Aggregate current activity for the latest observed placement of each instance.
	// Runtime snapshots feed periodic background sweeps. Counts refresh after each
	// complete sweep. Instance listings retain their separate three-month history for
	// failure discovery.
	Health ApplicationNewResponseCcDurableObjectApplicationHealth `json:"health"`
	// Application-wide logging settings for a Durable Object-managed application. The
	// application publishes these settings to its runtime metadata. Updating them does
	// not create a deployment or rollout.
	Observability ApplicationNewResponseCcDurableObjectApplicationObservability `json:"observability"`
	JSON          applicationNewResponseCcDurableObjectApplicationJSON          `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationJSON contains the JSON metadata
// for the struct [ApplicationNewResponseCcDurableObjectApplication]
type applicationNewResponseCcDurableObjectApplicationJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	CreatedAt        apijson.Field
	DurableObjects   apijson.Field
	Name             apijson.Field
	SchedulingPolicy apijson.Field
	UpdatedAt        apijson.Field
	Configuration    apijson.Field
	Health           apijson.Field
	Observability    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationNewResponseCcDurableObjectApplication) implementsApplicationNewResponse() {}

// Durable object configuration using a namespace ID.
type ApplicationNewResponseCcDurableObjectApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                             `json:"namespace_id" api:"required"`
	JSON        applicationNewResponseCcDurableObjectApplicationDurableObjectsJSON `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationDurableObjectsJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationDurableObjects]
type applicationNewResponseCcDurableObjectApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

// Selects a Durable Object-managed application. Each Durable Object creates and
// manages the lifecycle of its container instance. Configure application-wide
// observability settings here. Deployment configuration, scaling, placement
// constraints, versions, and rollouts do not apply.
type ApplicationNewResponseCcDurableObjectApplicationSchedulingPolicy string

const (
	ApplicationNewResponseCcDurableObjectApplicationSchedulingPolicyDurableObject ApplicationNewResponseCcDurableObjectApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationNewResponseCcDurableObjectApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcDurableObjectApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// Application-wide settings for a Durable Object-managed application.
type ApplicationNewResponseCcDurableObjectApplicationConfiguration struct {
	AuthorizedKeys []ApplicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// Configuration properties for connecting with SSH to a container using Wrangler.
	WranglerSSH ApplicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSH `json:"wrangler_ssh"`
	JSON        applicationNewResponseCcDurableObjectApplicationConfigurationJSON        `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationConfigurationJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationConfiguration]
type applicationNewResponseCcDurableObjectApplicationConfigurationJSON struct {
	AuthorizedKeys apijson.Field
	WranglerSSH    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                         `json:"name"`
	JSON applicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKey]
type applicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// Configuration properties for connecting with SSH to a container using Wrangler.
type ApplicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSH struct {
	Enabled bool                                                                         `json:"enabled"`
	Port    int64                                                                        `json:"port"`
	JSON    applicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON
// contains the JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSH]
type applicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON struct {
	Enabled     apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON) RawJSON() string {
	return r.raw
}

// Aggregate current activity for the latest observed placement of each instance.
// Runtime snapshots feed periodic background sweeps. Counts refresh after each
// complete sweep. Instance listings retain their separate three-month history for
// failure discovery.
type ApplicationNewResponseCcDurableObjectApplicationHealth struct {
	// Counts of observed non-terminal instances.
	Instances ApplicationNewResponseCcDurableObjectApplicationHealthInstances `json:"instances" api:"required"`
	// Present as pending until the first activity sweep completes; omitted afterward.
	Summary ApplicationNewResponseCcDurableObjectApplicationHealthSummary `json:"summary"`
	JSON    applicationNewResponseCcDurableObjectApplicationHealthJSON    `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationHealthJSON contains the JSON
// metadata for the struct [ApplicationNewResponseCcDurableObjectApplicationHealth]
type applicationNewResponseCcDurableObjectApplicationHealthJSON struct {
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationHealthJSON) RawJSON() string {
	return r.raw
}

// Counts of observed non-terminal instances.
type ApplicationNewResponseCcDurableObjectApplicationHealthInstances struct {
	// Number of instances whose runtime reports running or stopping.
	Active int64 `json:"active" api:"required"`
	// Number of instances whose runtime reports starting.
	Starting int64                                                               `json:"starting" api:"required"`
	JSON     applicationNewResponseCcDurableObjectApplicationHealthInstancesJSON `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationHealthInstancesJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationHealthInstances]
type applicationNewResponseCcDurableObjectApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Starting    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// Present as pending until the first activity sweep completes; omitted afterward.
type ApplicationNewResponseCcDurableObjectApplicationHealthSummary string

const (
	ApplicationNewResponseCcDurableObjectApplicationHealthSummaryPending ApplicationNewResponseCcDurableObjectApplicationHealthSummary = "pending"
)

func (r ApplicationNewResponseCcDurableObjectApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationNewResponseCcDurableObjectApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Application-wide logging settings for a Durable Object-managed application. The
// application publishes these settings to its runtime metadata. Updating them does
// not create a deployment or rollout.
type ApplicationNewResponseCcDurableObjectApplicationObservability struct {
	// Application-wide logging settings.
	Logs ApplicationNewResponseCcDurableObjectApplicationObservabilityLogs `json:"logs"`
	JSON applicationNewResponseCcDurableObjectApplicationObservabilityJSON `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationObservabilityJSON contains the
// JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationObservability]
type applicationNewResponseCcDurableObjectApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Application-wide logging settings.
type ApplicationNewResponseCcDurableObjectApplicationObservabilityLogs struct {
	Enabled bool                                                                  `json:"enabled"`
	JSON    applicationNewResponseCcDurableObjectApplicationObservabilityLogsJSON `json:"-"`
}

// applicationNewResponseCcDurableObjectApplicationObservabilityLogsJSON contains
// the JSON metadata for the struct
// [ApplicationNewResponseCcDurableObjectApplicationObservabilityLogs]
type applicationNewResponseCcDurableObjectApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseCcDurableObjectApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseCcDurableObjectApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationNewResponseSchedulingPolicy string

const (
	ApplicationNewResponseSchedulingPolicyDefault       ApplicationNewResponseSchedulingPolicy = "default"
	ApplicationNewResponseSchedulingPolicyDurableObject ApplicationNewResponseSchedulingPolicy = "durable_object"
)

func (r ApplicationNewResponseSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationNewResponseSchedulingPolicyDefault, ApplicationNewResponseSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// The public Containers API returns an application.
type ApplicationListResponse struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationListResponseSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string `json:"active_rollout_id"`
	// This field can have the runtime type of
	// [ApplicationListResponseCcScheduledApplicationConfiguration],
	// [ApplicationListResponseCcDurableObjectApplicationConfiguration].
	Configuration interface{} `json:"configuration"`
	// This field can have the runtime type of
	// [ApplicationListResponseCcScheduledApplicationConstraints].
	Constraints interface{} `json:"constraints"`
	// This field can have the runtime type of
	// [ApplicationListResponseCcScheduledApplicationDurableObjects],
	// [ApplicationListResponseCcDurableObjectApplicationDurableObjects].
	DurableObjects interface{} `json:"durable_objects"`
	// This field can have the runtime type of
	// [ApplicationListResponseCcScheduledApplicationHealth],
	// [ApplicationListResponseCcDurableObjectApplicationHealth].
	Health interface{} `json:"health"`
	// Number of deployments to create.
	Instances int64 `json:"instances"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// This field can have the runtime type of
	// [ApplicationListResponseCcScheduledApplicationObservability],
	// [ApplicationListResponseCcDurableObjectApplicationObservability].
	Observability interface{} `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                       `json:"rollout_active_grace_period"`
	Version                  int64                       `json:"version"`
	JSON                     applicationListResponseJSON `json:"-"`
	union                    ApplicationListResponseUnion
}

// applicationListResponseJSON contains the JSON metadata for the struct
// [ApplicationListResponse]
type applicationListResponseJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	CreatedAt                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	ActiveRolloutID          apijson.Field
	Configuration            apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	Instances                apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	Version                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r applicationListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ApplicationListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ApplicationListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ApplicationListResponseCcScheduledApplication],
// [ApplicationListResponseCcDurableObjectApplication].
func (r ApplicationListResponse) AsUnion() ApplicationListResponseUnion {
	return r.union
}

// The public Containers API returns an application.
//
// Union satisfied by [ApplicationListResponseCcScheduledApplication] or
// [ApplicationListResponseCcDurableObjectApplication].
type ApplicationListResponseUnion interface {
	implementsApplicationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationListResponseCcScheduledApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationListResponseCcDurableObjectApplication{}),
		},
	)
}

// Describes an application and the parameters that govern how it places its
// instances.
type ApplicationListResponseCcScheduledApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// User-specified container configuration.
	Configuration ApplicationListResponseCcScheduledApplicationConfiguration `json:"configuration" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Number of deployments to create.
	Instances int64 `json:"instances" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationListResponseCcScheduledApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	Version   int64  `json:"version" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string                                                   `json:"active_rollout_id"`
	Constraints     ApplicationListResponseCcScheduledApplicationConstraints `json:"constraints"`
	// Durable object configuration stored on and returned from a Cloudchamber
	// application.
	DurableObjects ApplicationListResponseCcScheduledApplicationDurableObjects `json:"durable_objects"`
	Health         ApplicationListResponseCcScheduledApplicationHealth         `json:"health"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// Top-level observability settings for the application. This field is mutually
	// exclusive with configuration.observability.
	Observability ApplicationListResponseCcScheduledApplicationObservability `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                                             `json:"rollout_active_grace_period"`
	JSON                     applicationListResponseCcScheduledApplicationJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationJSON contains the JSON metadata for
// the struct [ApplicationListResponseCcScheduledApplication]
type applicationListResponseCcScheduledApplicationJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	Configuration            apijson.Field
	CreatedAt                apijson.Field
	Instances                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	Version                  apijson.Field
	ActiveRolloutID          apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationListResponseCcScheduledApplication) implementsApplicationListResponse() {}

// User-specified container configuration.
type ApplicationListResponseCcScheduledApplicationConfiguration struct {
	// Image url.
	Image          string                                                                    `json:"image" api:"required"`
	AuthorizedKeys []ApplicationListResponseCcScheduledApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command []string `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint []string `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables []ApplicationListResponseCcScheduledApplicationConfigurationEnvironmentVariable `json:"environment_variables"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationListResponseCcScheduledApplicationConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationListResponseCcScheduledApplicationConfigurationObservability `json:"observability"`
	JSON          applicationListResponseCcScheduledApplicationConfigurationJSON          `json:"-"`
}

// applicationListResponseCcScheduledApplicationConfigurationJSON contains the JSON
// metadata for the struct
// [ApplicationListResponseCcScheduledApplicationConfiguration]
type applicationListResponseCcScheduledApplicationConfigurationJSON struct {
	Image                apijson.Field
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationListResponseCcScheduledApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                      `json:"name"`
	JSON applicationListResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationConfigurationAuthorizedKey]
type applicationListResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationListResponseCcScheduledApplicationConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                                            `json:"value" api:"required"`
	JSON  applicationListResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationConfigurationEnvironmentVariable]
type applicationListResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationListResponseCcScheduledApplicationConfigurationInstanceType string

const (
	ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeLite      ApplicationListResponseCcScheduledApplicationConfigurationInstanceType = "lite"
	ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeBasic     ApplicationListResponseCcScheduledApplicationConfigurationInstanceType = "basic"
	ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard1 ApplicationListResponseCcScheduledApplicationConfigurationInstanceType = "standard-1"
	ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard2 ApplicationListResponseCcScheduledApplicationConfigurationInstanceType = "standard-2"
	ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard3 ApplicationListResponseCcScheduledApplicationConfigurationInstanceType = "standard-3"
	ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard4 ApplicationListResponseCcScheduledApplicationConfigurationInstanceType = "standard-4"
)

func (r ApplicationListResponseCcScheduledApplicationConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeLite, ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeBasic, ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard1, ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard2, ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard3, ApplicationListResponseCcScheduledApplicationConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationListResponseCcScheduledApplicationConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationListResponseCcScheduledApplicationConfigurationObservabilityLogs `json:"logs"`
	JSON applicationListResponseCcScheduledApplicationConfigurationObservabilityJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationConfigurationObservabilityJSON
// contains the JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationConfigurationObservability]
type applicationListResponseCcScheduledApplicationConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationListResponseCcScheduledApplicationConfigurationObservabilityLogs struct {
	Enabled bool                                                                            `json:"enabled"`
	JSON    applicationListResponseCcScheduledApplicationConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationConfigurationObservabilityLogsJSON
// contains the JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationConfigurationObservabilityLogs]
type applicationListResponseCcScheduledApplicationConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationListResponseCcScheduledApplicationSchedulingPolicy string

const (
	ApplicationListResponseCcScheduledApplicationSchedulingPolicyDefault       ApplicationListResponseCcScheduledApplicationSchedulingPolicy = "default"
	ApplicationListResponseCcScheduledApplicationSchedulingPolicyDurableObject ApplicationListResponseCcScheduledApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationListResponseCcScheduledApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcScheduledApplicationSchedulingPolicyDefault, ApplicationListResponseCcScheduledApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

type ApplicationListResponseCcScheduledApplicationConstraints struct {
	// Restricts placement to datacenters in the selected jurisdiction. Choose "eu",
	// "fedramp", or "us". When combined with regions, EU supports EEUR and WEUR while
	// FedRAMP and US support ENAM and WNAM.
	Jurisdiction string                                                       `json:"jurisdiction"`
	Regions      []string                                                     `json:"regions"`
	JSON         applicationListResponseCcScheduledApplicationConstraintsJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationConstraintsJSON contains the JSON
// metadata for the struct
// [ApplicationListResponseCcScheduledApplicationConstraints]
type applicationListResponseCcScheduledApplicationConstraintsJSON struct {
	Jurisdiction apijson.Field
	Regions      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationConstraintsJSON) RawJSON() string {
	return r.raw
}

// Durable object configuration stored on and returned from a Cloudchamber
// application.
type ApplicationListResponseCcScheduledApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                          `json:"namespace_id" api:"required"`
	JSON        applicationListResponseCcScheduledApplicationDurableObjectsJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationDurableObjectsJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationDurableObjects]
type applicationListResponseCcScheduledApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

type ApplicationListResponseCcScheduledApplicationHealth struct {
	Errors []ApplicationListResponseCcScheduledApplicationHealthError `json:"errors" api:"required"`
	// Shows a count of application instance states.
	Instances ApplicationListResponseCcScheduledApplicationHealthInstances `json:"instances" api:"required"`
	// High-level health assessment. Only populated for "new_instances" strategy. Based
	// on a sample of target-version instances rather than a full count.
	//
	// - "pending": Zero target-version instances exist yet.
	// - "healthy": Every sampled target-version instance reports running or active.
	// - "degraded": Some sampled instances remain starting or scheduling.
	// - "unhealthy": One or more sampled instances have failed.
	Summary ApplicationListResponseCcScheduledApplicationHealthSummary `json:"summary"`
	JSON    applicationListResponseCcScheduledApplicationHealthJSON    `json:"-"`
}

// applicationListResponseCcScheduledApplicationHealthJSON contains the JSON
// metadata for the struct [ApplicationListResponseCcScheduledApplicationHealth]
type applicationListResponseCcScheduledApplicationHealthJSON struct {
	Errors      apijson.Field
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationHealthJSON) RawJSON() string {
	return r.raw
}

type ApplicationListResponseCcScheduledApplicationHealthError struct {
	// An event within a Placement or a Job.
	Event ApplicationListResponseCcScheduledApplicationHealthErrorsEvent `json:"event" api:"required"`
	// An instance ID represents an identifier of an instance configuration that
	// maintains an underlying placement.
	InstanceID string                                                       `json:"instance_id" api:"required"`
	JSON       applicationListResponseCcScheduledApplicationHealthErrorJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationHealthErrorJSON contains the JSON
// metadata for the struct
// [ApplicationListResponseCcScheduledApplicationHealthError]
type applicationListResponseCcScheduledApplicationHealthErrorJSON struct {
	Event       apijson.Field
	InstanceID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationHealthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationHealthErrorJSON) RawJSON() string {
	return r.raw
}

// An event within a Placement or a Job.
type ApplicationListResponseCcScheduledApplicationHealthErrorsEvent struct {
	ID      string                 `json:"id" api:"required"`
	Details map[string]interface{} `json:"details" api:"required"`
	Message string                 `json:"message" api:"required"`
	// Name of the event that describes the kind event that happened.
	//
	// - SchedulerPlaced: It's the first event that creates a container placement. It
	//   happens when the Containers runtime was able to retrieve deployment resources
	//   and start verifying everything is correct.
	// - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
	//   container.
	// - VMStarted: It's sent when the Containers runtime starts the VM. The container
	//   might remain unhealthy at this point.
	// - ImagePulled: It's sent when the Containers runtime pulls the image
	//   successfully.
	// - ImagePullError: It's sent when the Containers runtime is having issues pulling
	//   the image. The message and details have more information on what happened for
	//   debugging.
	// - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
	//   VM.
	// - VMStopping: It's sent when the scheduler is stopping the VM.
	// - VMStopped: It's sent when the VM finally exits.
	// - VMFailed: It's sent when the scheduling of the VM failed in the current
	//   location.
	// - RuntimeStartFailed: It's sent when the runtime hits an internal error.
	// - SSHStarted: It's sent when the container gains network connectivity and opens
	//   the SSH port. Containers only send this event when SSH keys exist.
	// - CheckUpdate: Sent when the status of a health or readiness check changes. This
	//   may also affect the health status of the placement.
	// - DurableObjectConnected: Sent when a durable object instance connects and gains
	//   control of the deployment. This event is only sent for durable object
	//   deployments. It is sent after VMStarted.
	// - ContainerStarted: It's sent when the container starts running.
	Name         ApplicationListResponseCcScheduledApplicationHealthErrorsEventName `json:"name" api:"required"`
	StatusChange map[string]interface{}                                             `json:"statusChange" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	Time string                                                             `json:"time" api:"required"`
	Type ApplicationListResponseCcScheduledApplicationHealthErrorsEventType `json:"type" api:"required"`
	JSON applicationListResponseCcScheduledApplicationHealthErrorsEventJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationHealthErrorsEventJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationHealthErrorsEvent]
type applicationListResponseCcScheduledApplicationHealthErrorsEventJSON struct {
	ID           apijson.Field
	Details      apijson.Field
	Message      apijson.Field
	Name         apijson.Field
	StatusChange apijson.Field
	Time         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationHealthErrorsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationHealthErrorsEventJSON) RawJSON() string {
	return r.raw
}

// Name of the event that describes the kind event that happened.
//
//   - SchedulerPlaced: It's the first event that creates a container placement. It
//     happens when the Containers runtime was able to retrieve deployment resources
//     and start verifying everything is correct.
//   - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
//     container.
//   - VMStarted: It's sent when the Containers runtime starts the VM. The container
//     might remain unhealthy at this point.
//   - ImagePulled: It's sent when the Containers runtime pulls the image
//     successfully.
//   - ImagePullError: It's sent when the Containers runtime is having issues pulling
//     the image. The message and details have more information on what happened for
//     debugging.
//   - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
//     VM.
//   - VMStopping: It's sent when the scheduler is stopping the VM.
//   - VMStopped: It's sent when the VM finally exits.
//   - VMFailed: It's sent when the scheduling of the VM failed in the current
//     location.
//   - RuntimeStartFailed: It's sent when the runtime hits an internal error.
//   - SSHStarted: It's sent when the container gains network connectivity and opens
//     the SSH port. Containers only send this event when SSH keys exist.
//   - CheckUpdate: Sent when the status of a health or readiness check changes. This
//     may also affect the health status of the placement.
//   - DurableObjectConnected: Sent when a durable object instance connects and gains
//     control of the deployment. This event is only sent for durable object
//     deployments. It is sent after VMStarted.
//   - ContainerStarted: It's sent when the container starts running.
type ApplicationListResponseCcScheduledApplicationHealthErrorsEventName string

const (
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced              ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "SchedulerPlaced"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned         ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssigned"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmStarted                    ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "VMStarted"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameImagePulled                  ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "ImagePulled"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameImagePullError               ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "ImagePullError"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart              ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "VMFailedToStart"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssignmentFailed"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmRunning                    ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "VMRunning"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmStopping                   ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "VMStopping"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmStopped                    ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "VMStopped"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmFailed                     ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "VMFailed"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed           ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "RuntimeStartFailed"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted                   ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "SSHStarted"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates         ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "ServiceHealthUpdates"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate                  ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "CheckUpdate"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected       ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "DurableObjectConnected"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted             ApplicationListResponseCcScheduledApplicationHealthErrorsEventName = "ContainerStarted"
)

func (r ApplicationListResponseCcScheduledApplicationHealthErrorsEventName) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmStarted, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameImagePulled, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameImagePullError, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmRunning, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmStopping, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmStopped, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameVmFailed, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected, ApplicationListResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted:
		return true
	}
	return false
}

type ApplicationListResponseCcScheduledApplicationHealthErrorsEventType string

const (
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeInfo        ApplicationListResponseCcScheduledApplicationHealthErrorsEventType = "Info"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeError       ApplicationListResponseCcScheduledApplicationHealthErrorsEventType = "Error"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeWarn        ApplicationListResponseCcScheduledApplicationHealthErrorsEventType = "Warn"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeUserError   ApplicationListResponseCcScheduledApplicationHealthErrorsEventType = "UserError"
	ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeSystemError ApplicationListResponseCcScheduledApplicationHealthErrorsEventType = "SystemError"
)

func (r ApplicationListResponseCcScheduledApplicationHealthErrorsEventType) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeInfo, ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeError, ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeWarn, ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeUserError, ApplicationListResponseCcScheduledApplicationHealthErrorsEventTypeSystemError:
		return true
	}
	return false
}

// Shows a count of application instance states.
type ApplicationListResponseCcScheduledApplicationHealthInstances struct {
	// Number of instances whose runtime reports the container as running
	// (container_status = "running"). This is a subset of the placements that remain
	// up: an instance that is already bound to a Durable Object and serving traffic is
	// counted under "assigned" until its container_status catches up to "running", so
	// container_status can briefly lag Durable Object attachment under churn. To
	// estimate running, Durable-Object-bound instances, sum "active" + "assigned"
	// rather than reading "active" alone.
	Active int64 `json:"active" api:"required"`
	// Number of instances bound to a Durable Object with a running placement whose
	// container_status remains behind "running". These count as live, serving
	// instances; "active" + "assigned" approximates the running, Durable-Object-bound
	// count.
	Assigned int64                                                            `json:"assigned" api:"required"`
	JSON     applicationListResponseCcScheduledApplicationHealthInstancesJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationHealthInstancesJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationHealthInstances]
type applicationListResponseCcScheduledApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Assigned    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// High-level health assessment. Only populated for "new_instances" strategy. Based
// on a sample of target-version instances rather than a full count.
//
// - "pending": Zero target-version instances exist yet.
// - "healthy": Every sampled target-version instance reports running or active.
// - "degraded": Some sampled instances remain starting or scheduling.
// - "unhealthy": One or more sampled instances have failed.
type ApplicationListResponseCcScheduledApplicationHealthSummary string

const (
	ApplicationListResponseCcScheduledApplicationHealthSummaryHealthy   ApplicationListResponseCcScheduledApplicationHealthSummary = "healthy"
	ApplicationListResponseCcScheduledApplicationHealthSummaryDegraded  ApplicationListResponseCcScheduledApplicationHealthSummary = "degraded"
	ApplicationListResponseCcScheduledApplicationHealthSummaryUnhealthy ApplicationListResponseCcScheduledApplicationHealthSummary = "unhealthy"
	ApplicationListResponseCcScheduledApplicationHealthSummaryPending   ApplicationListResponseCcScheduledApplicationHealthSummary = "pending"
)

func (r ApplicationListResponseCcScheduledApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcScheduledApplicationHealthSummaryHealthy, ApplicationListResponseCcScheduledApplicationHealthSummaryDegraded, ApplicationListResponseCcScheduledApplicationHealthSummaryUnhealthy, ApplicationListResponseCcScheduledApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Top-level observability settings for the application. This field is mutually
// exclusive with configuration.observability.
type ApplicationListResponseCcScheduledApplicationObservability struct {
	// Observability logging settings.
	Logs ApplicationListResponseCcScheduledApplicationObservabilityLogs `json:"logs"`
	JSON applicationListResponseCcScheduledApplicationObservabilityJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationObservabilityJSON contains the JSON
// metadata for the struct
// [ApplicationListResponseCcScheduledApplicationObservability]
type applicationListResponseCcScheduledApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationListResponseCcScheduledApplicationObservabilityLogs struct {
	Enabled bool                                                               `json:"enabled"`
	JSON    applicationListResponseCcScheduledApplicationObservabilityLogsJSON `json:"-"`
}

// applicationListResponseCcScheduledApplicationObservabilityLogsJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcScheduledApplicationObservabilityLogs]
type applicationListResponseCcScheduledApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcScheduledApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcScheduledApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Each Durable Object creates and manages the lifecycle of its container instance.
type ApplicationListResponseCcDurableObjectApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Durable object configuration using a namespace ID.
	DurableObjects ApplicationListResponseCcDurableObjectApplicationDurableObjects `json:"durable_objects" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// Selects a Durable Object-managed application. Each Durable Object creates and
	// manages the lifecycle of its container instance. Configure application-wide
	// observability settings here. Deployment configuration, scaling, placement
	// constraints, versions, and rollouts do not apply.
	SchedulingPolicy ApplicationListResponseCcDurableObjectApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Application-wide settings for a Durable Object-managed application.
	Configuration ApplicationListResponseCcDurableObjectApplicationConfiguration `json:"configuration"`
	// Aggregate current activity for the latest observed placement of each instance.
	// Runtime snapshots feed periodic background sweeps. Counts refresh after each
	// complete sweep. Instance listings retain their separate three-month history for
	// failure discovery.
	Health ApplicationListResponseCcDurableObjectApplicationHealth `json:"health"`
	// Application-wide logging settings for a Durable Object-managed application. The
	// application publishes these settings to its runtime metadata. Updating them does
	// not create a deployment or rollout.
	Observability ApplicationListResponseCcDurableObjectApplicationObservability `json:"observability"`
	JSON          applicationListResponseCcDurableObjectApplicationJSON          `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationJSON contains the JSON metadata
// for the struct [ApplicationListResponseCcDurableObjectApplication]
type applicationListResponseCcDurableObjectApplicationJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	CreatedAt        apijson.Field
	DurableObjects   apijson.Field
	Name             apijson.Field
	SchedulingPolicy apijson.Field
	UpdatedAt        apijson.Field
	Configuration    apijson.Field
	Health           apijson.Field
	Observability    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationListResponseCcDurableObjectApplication) implementsApplicationListResponse() {}

// Durable object configuration using a namespace ID.
type ApplicationListResponseCcDurableObjectApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                              `json:"namespace_id" api:"required"`
	JSON        applicationListResponseCcDurableObjectApplicationDurableObjectsJSON `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationDurableObjectsJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationDurableObjects]
type applicationListResponseCcDurableObjectApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

// Selects a Durable Object-managed application. Each Durable Object creates and
// manages the lifecycle of its container instance. Configure application-wide
// observability settings here. Deployment configuration, scaling, placement
// constraints, versions, and rollouts do not apply.
type ApplicationListResponseCcDurableObjectApplicationSchedulingPolicy string

const (
	ApplicationListResponseCcDurableObjectApplicationSchedulingPolicyDurableObject ApplicationListResponseCcDurableObjectApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationListResponseCcDurableObjectApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcDurableObjectApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// Application-wide settings for a Durable Object-managed application.
type ApplicationListResponseCcDurableObjectApplicationConfiguration struct {
	AuthorizedKeys []ApplicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// Configuration properties for connecting with SSH to a container using Wrangler.
	WranglerSSH ApplicationListResponseCcDurableObjectApplicationConfigurationWranglerSSH `json:"wrangler_ssh"`
	JSON        applicationListResponseCcDurableObjectApplicationConfigurationJSON        `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationConfigurationJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationConfiguration]
type applicationListResponseCcDurableObjectApplicationConfigurationJSON struct {
	AuthorizedKeys apijson.Field
	WranglerSSH    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                          `json:"name"`
	JSON applicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKey]
type applicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// Configuration properties for connecting with SSH to a container using Wrangler.
type ApplicationListResponseCcDurableObjectApplicationConfigurationWranglerSSH struct {
	Enabled bool                                                                          `json:"enabled"`
	Port    int64                                                                         `json:"port"`
	JSON    applicationListResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON
// contains the JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationConfigurationWranglerSSH]
type applicationListResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON struct {
	Enabled     apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationConfigurationWranglerSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON) RawJSON() string {
	return r.raw
}

// Aggregate current activity for the latest observed placement of each instance.
// Runtime snapshots feed periodic background sweeps. Counts refresh after each
// complete sweep. Instance listings retain their separate three-month history for
// failure discovery.
type ApplicationListResponseCcDurableObjectApplicationHealth struct {
	// Counts of observed non-terminal instances.
	Instances ApplicationListResponseCcDurableObjectApplicationHealthInstances `json:"instances" api:"required"`
	// Present as pending until the first activity sweep completes; omitted afterward.
	Summary ApplicationListResponseCcDurableObjectApplicationHealthSummary `json:"summary"`
	JSON    applicationListResponseCcDurableObjectApplicationHealthJSON    `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationHealthJSON contains the JSON
// metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationHealth]
type applicationListResponseCcDurableObjectApplicationHealthJSON struct {
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationHealthJSON) RawJSON() string {
	return r.raw
}

// Counts of observed non-terminal instances.
type ApplicationListResponseCcDurableObjectApplicationHealthInstances struct {
	// Number of instances whose runtime reports running or stopping.
	Active int64 `json:"active" api:"required"`
	// Number of instances whose runtime reports starting.
	Starting int64                                                                `json:"starting" api:"required"`
	JSON     applicationListResponseCcDurableObjectApplicationHealthInstancesJSON `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationHealthInstancesJSON contains
// the JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationHealthInstances]
type applicationListResponseCcDurableObjectApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Starting    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// Present as pending until the first activity sweep completes; omitted afterward.
type ApplicationListResponseCcDurableObjectApplicationHealthSummary string

const (
	ApplicationListResponseCcDurableObjectApplicationHealthSummaryPending ApplicationListResponseCcDurableObjectApplicationHealthSummary = "pending"
)

func (r ApplicationListResponseCcDurableObjectApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationListResponseCcDurableObjectApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Application-wide logging settings for a Durable Object-managed application. The
// application publishes these settings to its runtime metadata. Updating them does
// not create a deployment or rollout.
type ApplicationListResponseCcDurableObjectApplicationObservability struct {
	// Application-wide logging settings.
	Logs ApplicationListResponseCcDurableObjectApplicationObservabilityLogs `json:"logs"`
	JSON applicationListResponseCcDurableObjectApplicationObservabilityJSON `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationObservabilityJSON contains the
// JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationObservability]
type applicationListResponseCcDurableObjectApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Application-wide logging settings.
type ApplicationListResponseCcDurableObjectApplicationObservabilityLogs struct {
	Enabled bool                                                                   `json:"enabled"`
	JSON    applicationListResponseCcDurableObjectApplicationObservabilityLogsJSON `json:"-"`
}

// applicationListResponseCcDurableObjectApplicationObservabilityLogsJSON contains
// the JSON metadata for the struct
// [ApplicationListResponseCcDurableObjectApplicationObservabilityLogs]
type applicationListResponseCcDurableObjectApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationListResponseCcDurableObjectApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationListResponseCcDurableObjectApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationListResponseSchedulingPolicy string

const (
	ApplicationListResponseSchedulingPolicyDefault       ApplicationListResponseSchedulingPolicy = "default"
	ApplicationListResponseSchedulingPolicyDurableObject ApplicationListResponseSchedulingPolicy = "durable_object"
)

func (r ApplicationListResponseSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationListResponseSchedulingPolicyDefault, ApplicationListResponseSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// Result of starting asynchronous deletion for a Containers application.
type ApplicationDeleteResponse struct {
	Message string                        `json:"message" api:"required"`
	JSON    applicationDeleteResponseJSON `json:"-"`
}

// applicationDeleteResponseJSON contains the JSON metadata for the struct
// [ApplicationDeleteResponse]
type applicationDeleteResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The public Containers API returns an application.
type ApplicationEditResponse struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationEditResponseSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string `json:"active_rollout_id"`
	// This field can have the runtime type of
	// [ApplicationEditResponseCcScheduledApplicationConfiguration],
	// [ApplicationEditResponseCcDurableObjectApplicationConfiguration].
	Configuration interface{} `json:"configuration"`
	// This field can have the runtime type of
	// [ApplicationEditResponseCcScheduledApplicationConstraints].
	Constraints interface{} `json:"constraints"`
	// This field can have the runtime type of
	// [ApplicationEditResponseCcScheduledApplicationDurableObjects],
	// [ApplicationEditResponseCcDurableObjectApplicationDurableObjects].
	DurableObjects interface{} `json:"durable_objects"`
	// This field can have the runtime type of
	// [ApplicationEditResponseCcScheduledApplicationHealth],
	// [ApplicationEditResponseCcDurableObjectApplicationHealth].
	Health interface{} `json:"health"`
	// Number of deployments to create.
	Instances int64 `json:"instances"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// This field can have the runtime type of
	// [ApplicationEditResponseCcScheduledApplicationObservability],
	// [ApplicationEditResponseCcDurableObjectApplicationObservability].
	Observability interface{} `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                       `json:"rollout_active_grace_period"`
	Version                  int64                       `json:"version"`
	JSON                     applicationEditResponseJSON `json:"-"`
	union                    ApplicationEditResponseUnion
}

// applicationEditResponseJSON contains the JSON metadata for the struct
// [ApplicationEditResponse]
type applicationEditResponseJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	CreatedAt                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	ActiveRolloutID          apijson.Field
	Configuration            apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	Instances                apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	Version                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r applicationEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationEditResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ApplicationEditResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ApplicationEditResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ApplicationEditResponseCcScheduledApplication],
// [ApplicationEditResponseCcDurableObjectApplication].
func (r ApplicationEditResponse) AsUnion() ApplicationEditResponseUnion {
	return r.union
}

// The public Containers API returns an application.
//
// Union satisfied by [ApplicationEditResponseCcScheduledApplication] or
// [ApplicationEditResponseCcDurableObjectApplication].
type ApplicationEditResponseUnion interface {
	implementsApplicationEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationEditResponseCcScheduledApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationEditResponseCcDurableObjectApplication{}),
		},
	)
}

// Describes an application and the parameters that govern how it places its
// instances.
type ApplicationEditResponseCcScheduledApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// User-specified container configuration.
	Configuration ApplicationEditResponseCcScheduledApplicationConfiguration `json:"configuration" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Number of deployments to create.
	Instances int64 `json:"instances" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationEditResponseCcScheduledApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	Version   int64  `json:"version" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string                                                   `json:"active_rollout_id"`
	Constraints     ApplicationEditResponseCcScheduledApplicationConstraints `json:"constraints"`
	// Durable object configuration stored on and returned from a Cloudchamber
	// application.
	DurableObjects ApplicationEditResponseCcScheduledApplicationDurableObjects `json:"durable_objects"`
	Health         ApplicationEditResponseCcScheduledApplicationHealth         `json:"health"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// Top-level observability settings for the application. This field is mutually
	// exclusive with configuration.observability.
	Observability ApplicationEditResponseCcScheduledApplicationObservability `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                                             `json:"rollout_active_grace_period"`
	JSON                     applicationEditResponseCcScheduledApplicationJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationJSON contains the JSON metadata for
// the struct [ApplicationEditResponseCcScheduledApplication]
type applicationEditResponseCcScheduledApplicationJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	Configuration            apijson.Field
	CreatedAt                apijson.Field
	Instances                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	Version                  apijson.Field
	ActiveRolloutID          apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationEditResponseCcScheduledApplication) implementsApplicationEditResponse() {}

// User-specified container configuration.
type ApplicationEditResponseCcScheduledApplicationConfiguration struct {
	// Image url.
	Image          string                                                                    `json:"image" api:"required"`
	AuthorizedKeys []ApplicationEditResponseCcScheduledApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command []string `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint []string `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables []ApplicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariable `json:"environment_variables"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationEditResponseCcScheduledApplicationConfigurationObservability `json:"observability"`
	JSON          applicationEditResponseCcScheduledApplicationConfigurationJSON          `json:"-"`
}

// applicationEditResponseCcScheduledApplicationConfigurationJSON contains the JSON
// metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationConfiguration]
type applicationEditResponseCcScheduledApplicationConfigurationJSON struct {
	Image                apijson.Field
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationEditResponseCcScheduledApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                      `json:"name"`
	JSON applicationEditResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationConfigurationAuthorizedKey]
type applicationEditResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                                            `json:"value" api:"required"`
	JSON  applicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariable]
type applicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType string

const (
	ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeLite      ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType = "lite"
	ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeBasic     ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType = "basic"
	ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard1 ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType = "standard-1"
	ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard2 ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType = "standard-2"
	ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard3 ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType = "standard-3"
	ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard4 ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType = "standard-4"
)

func (r ApplicationEditResponseCcScheduledApplicationConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeLite, ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeBasic, ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard1, ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard2, ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard3, ApplicationEditResponseCcScheduledApplicationConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationEditResponseCcScheduledApplicationConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationEditResponseCcScheduledApplicationConfigurationObservabilityLogs `json:"logs"`
	JSON applicationEditResponseCcScheduledApplicationConfigurationObservabilityJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationConfigurationObservabilityJSON
// contains the JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationConfigurationObservability]
type applicationEditResponseCcScheduledApplicationConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationEditResponseCcScheduledApplicationConfigurationObservabilityLogs struct {
	Enabled bool                                                                            `json:"enabled"`
	JSON    applicationEditResponseCcScheduledApplicationConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationConfigurationObservabilityLogsJSON
// contains the JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationConfigurationObservabilityLogs]
type applicationEditResponseCcScheduledApplicationConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationEditResponseCcScheduledApplicationSchedulingPolicy string

const (
	ApplicationEditResponseCcScheduledApplicationSchedulingPolicyDefault       ApplicationEditResponseCcScheduledApplicationSchedulingPolicy = "default"
	ApplicationEditResponseCcScheduledApplicationSchedulingPolicyDurableObject ApplicationEditResponseCcScheduledApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationEditResponseCcScheduledApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcScheduledApplicationSchedulingPolicyDefault, ApplicationEditResponseCcScheduledApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

type ApplicationEditResponseCcScheduledApplicationConstraints struct {
	// Restricts placement to datacenters in the selected jurisdiction. Choose "eu",
	// "fedramp", or "us". When combined with regions, EU supports EEUR and WEUR while
	// FedRAMP and US support ENAM and WNAM.
	Jurisdiction string                                                       `json:"jurisdiction"`
	Regions      []string                                                     `json:"regions"`
	JSON         applicationEditResponseCcScheduledApplicationConstraintsJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationConstraintsJSON contains the JSON
// metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationConstraints]
type applicationEditResponseCcScheduledApplicationConstraintsJSON struct {
	Jurisdiction apijson.Field
	Regions      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationConstraintsJSON) RawJSON() string {
	return r.raw
}

// Durable object configuration stored on and returned from a Cloudchamber
// application.
type ApplicationEditResponseCcScheduledApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                          `json:"namespace_id" api:"required"`
	JSON        applicationEditResponseCcScheduledApplicationDurableObjectsJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationDurableObjectsJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationDurableObjects]
type applicationEditResponseCcScheduledApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditResponseCcScheduledApplicationHealth struct {
	Errors []ApplicationEditResponseCcScheduledApplicationHealthError `json:"errors" api:"required"`
	// Shows a count of application instance states.
	Instances ApplicationEditResponseCcScheduledApplicationHealthInstances `json:"instances" api:"required"`
	// High-level health assessment. Only populated for "new_instances" strategy. Based
	// on a sample of target-version instances rather than a full count.
	//
	// - "pending": Zero target-version instances exist yet.
	// - "healthy": Every sampled target-version instance reports running or active.
	// - "degraded": Some sampled instances remain starting or scheduling.
	// - "unhealthy": One or more sampled instances have failed.
	Summary ApplicationEditResponseCcScheduledApplicationHealthSummary `json:"summary"`
	JSON    applicationEditResponseCcScheduledApplicationHealthJSON    `json:"-"`
}

// applicationEditResponseCcScheduledApplicationHealthJSON contains the JSON
// metadata for the struct [ApplicationEditResponseCcScheduledApplicationHealth]
type applicationEditResponseCcScheduledApplicationHealthJSON struct {
	Errors      apijson.Field
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationHealthJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditResponseCcScheduledApplicationHealthError struct {
	// An event within a Placement or a Job.
	Event ApplicationEditResponseCcScheduledApplicationHealthErrorsEvent `json:"event" api:"required"`
	// An instance ID represents an identifier of an instance configuration that
	// maintains an underlying placement.
	InstanceID string                                                       `json:"instance_id" api:"required"`
	JSON       applicationEditResponseCcScheduledApplicationHealthErrorJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationHealthErrorJSON contains the JSON
// metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationHealthError]
type applicationEditResponseCcScheduledApplicationHealthErrorJSON struct {
	Event       apijson.Field
	InstanceID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationHealthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationHealthErrorJSON) RawJSON() string {
	return r.raw
}

// An event within a Placement or a Job.
type ApplicationEditResponseCcScheduledApplicationHealthErrorsEvent struct {
	ID      string                 `json:"id" api:"required"`
	Details map[string]interface{} `json:"details" api:"required"`
	Message string                 `json:"message" api:"required"`
	// Name of the event that describes the kind event that happened.
	//
	// - SchedulerPlaced: It's the first event that creates a container placement. It
	//   happens when the Containers runtime was able to retrieve deployment resources
	//   and start verifying everything is correct.
	// - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
	//   container.
	// - VMStarted: It's sent when the Containers runtime starts the VM. The container
	//   might remain unhealthy at this point.
	// - ImagePulled: It's sent when the Containers runtime pulls the image
	//   successfully.
	// - ImagePullError: It's sent when the Containers runtime is having issues pulling
	//   the image. The message and details have more information on what happened for
	//   debugging.
	// - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
	//   VM.
	// - VMStopping: It's sent when the scheduler is stopping the VM.
	// - VMStopped: It's sent when the VM finally exits.
	// - VMFailed: It's sent when the scheduling of the VM failed in the current
	//   location.
	// - RuntimeStartFailed: It's sent when the runtime hits an internal error.
	// - SSHStarted: It's sent when the container gains network connectivity and opens
	//   the SSH port. Containers only send this event when SSH keys exist.
	// - CheckUpdate: Sent when the status of a health or readiness check changes. This
	//   may also affect the health status of the placement.
	// - DurableObjectConnected: Sent when a durable object instance connects and gains
	//   control of the deployment. This event is only sent for durable object
	//   deployments. It is sent after VMStarted.
	// - ContainerStarted: It's sent when the container starts running.
	Name         ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName `json:"name" api:"required"`
	StatusChange map[string]interface{}                                             `json:"statusChange" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	Time string                                                             `json:"time" api:"required"`
	Type ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType `json:"type" api:"required"`
	JSON applicationEditResponseCcScheduledApplicationHealthErrorsEventJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationHealthErrorsEventJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationHealthErrorsEvent]
type applicationEditResponseCcScheduledApplicationHealthErrorsEventJSON struct {
	ID           apijson.Field
	Details      apijson.Field
	Message      apijson.Field
	Name         apijson.Field
	StatusChange apijson.Field
	Time         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationHealthErrorsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationHealthErrorsEventJSON) RawJSON() string {
	return r.raw
}

// Name of the event that describes the kind event that happened.
//
//   - SchedulerPlaced: It's the first event that creates a container placement. It
//     happens when the Containers runtime was able to retrieve deployment resources
//     and start verifying everything is correct.
//   - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
//     container.
//   - VMStarted: It's sent when the Containers runtime starts the VM. The container
//     might remain unhealthy at this point.
//   - ImagePulled: It's sent when the Containers runtime pulls the image
//     successfully.
//   - ImagePullError: It's sent when the Containers runtime is having issues pulling
//     the image. The message and details have more information on what happened for
//     debugging.
//   - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
//     VM.
//   - VMStopping: It's sent when the scheduler is stopping the VM.
//   - VMStopped: It's sent when the VM finally exits.
//   - VMFailed: It's sent when the scheduling of the VM failed in the current
//     location.
//   - RuntimeStartFailed: It's sent when the runtime hits an internal error.
//   - SSHStarted: It's sent when the container gains network connectivity and opens
//     the SSH port. Containers only send this event when SSH keys exist.
//   - CheckUpdate: Sent when the status of a health or readiness check changes. This
//     may also affect the health status of the placement.
//   - DurableObjectConnected: Sent when a durable object instance connects and gains
//     control of the deployment. This event is only sent for durable object
//     deployments. It is sent after VMStarted.
//   - ContainerStarted: It's sent when the container starts running.
type ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName string

const (
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced              ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "SchedulerPlaced"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned         ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssigned"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmStarted                    ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "VMStarted"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameImagePulled                  ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "ImagePulled"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameImagePullError               ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "ImagePullError"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart              ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "VMFailedToStart"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssignmentFailed"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmRunning                    ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "VMRunning"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmStopping                   ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "VMStopping"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmStopped                    ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "VMStopped"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmFailed                     ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "VMFailed"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed           ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "RuntimeStartFailed"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted                   ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "SSHStarted"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates         ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "ServiceHealthUpdates"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate                  ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "CheckUpdate"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected       ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "DurableObjectConnected"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted             ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName = "ContainerStarted"
)

func (r ApplicationEditResponseCcScheduledApplicationHealthErrorsEventName) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmStarted, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameImagePulled, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameImagePullError, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmRunning, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmStopping, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmStopped, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameVmFailed, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted:
		return true
	}
	return false
}

type ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType string

const (
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeInfo        ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType = "Info"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeError       ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType = "Error"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeWarn        ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType = "Warn"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeUserError   ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType = "UserError"
	ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeSystemError ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType = "SystemError"
)

func (r ApplicationEditResponseCcScheduledApplicationHealthErrorsEventType) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeInfo, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeError, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeWarn, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeUserError, ApplicationEditResponseCcScheduledApplicationHealthErrorsEventTypeSystemError:
		return true
	}
	return false
}

// Shows a count of application instance states.
type ApplicationEditResponseCcScheduledApplicationHealthInstances struct {
	// Number of instances whose runtime reports the container as running
	// (container_status = "running"). This is a subset of the placements that remain
	// up: an instance that is already bound to a Durable Object and serving traffic is
	// counted under "assigned" until its container_status catches up to "running", so
	// container_status can briefly lag Durable Object attachment under churn. To
	// estimate running, Durable-Object-bound instances, sum "active" + "assigned"
	// rather than reading "active" alone.
	Active int64 `json:"active" api:"required"`
	// Number of instances bound to a Durable Object with a running placement whose
	// container_status remains behind "running". These count as live, serving
	// instances; "active" + "assigned" approximates the running, Durable-Object-bound
	// count.
	Assigned int64                                                            `json:"assigned" api:"required"`
	JSON     applicationEditResponseCcScheduledApplicationHealthInstancesJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationHealthInstancesJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationHealthInstances]
type applicationEditResponseCcScheduledApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Assigned    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// High-level health assessment. Only populated for "new_instances" strategy. Based
// on a sample of target-version instances rather than a full count.
//
// - "pending": Zero target-version instances exist yet.
// - "healthy": Every sampled target-version instance reports running or active.
// - "degraded": Some sampled instances remain starting or scheduling.
// - "unhealthy": One or more sampled instances have failed.
type ApplicationEditResponseCcScheduledApplicationHealthSummary string

const (
	ApplicationEditResponseCcScheduledApplicationHealthSummaryHealthy   ApplicationEditResponseCcScheduledApplicationHealthSummary = "healthy"
	ApplicationEditResponseCcScheduledApplicationHealthSummaryDegraded  ApplicationEditResponseCcScheduledApplicationHealthSummary = "degraded"
	ApplicationEditResponseCcScheduledApplicationHealthSummaryUnhealthy ApplicationEditResponseCcScheduledApplicationHealthSummary = "unhealthy"
	ApplicationEditResponseCcScheduledApplicationHealthSummaryPending   ApplicationEditResponseCcScheduledApplicationHealthSummary = "pending"
)

func (r ApplicationEditResponseCcScheduledApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcScheduledApplicationHealthSummaryHealthy, ApplicationEditResponseCcScheduledApplicationHealthSummaryDegraded, ApplicationEditResponseCcScheduledApplicationHealthSummaryUnhealthy, ApplicationEditResponseCcScheduledApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Top-level observability settings for the application. This field is mutually
// exclusive with configuration.observability.
type ApplicationEditResponseCcScheduledApplicationObservability struct {
	// Observability logging settings.
	Logs ApplicationEditResponseCcScheduledApplicationObservabilityLogs `json:"logs"`
	JSON applicationEditResponseCcScheduledApplicationObservabilityJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationObservabilityJSON contains the JSON
// metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationObservability]
type applicationEditResponseCcScheduledApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationEditResponseCcScheduledApplicationObservabilityLogs struct {
	Enabled bool                                                               `json:"enabled"`
	JSON    applicationEditResponseCcScheduledApplicationObservabilityLogsJSON `json:"-"`
}

// applicationEditResponseCcScheduledApplicationObservabilityLogsJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcScheduledApplicationObservabilityLogs]
type applicationEditResponseCcScheduledApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcScheduledApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcScheduledApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Each Durable Object creates and manages the lifecycle of its container instance.
type ApplicationEditResponseCcDurableObjectApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Durable object configuration using a namespace ID.
	DurableObjects ApplicationEditResponseCcDurableObjectApplicationDurableObjects `json:"durable_objects" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// Selects a Durable Object-managed application. Each Durable Object creates and
	// manages the lifecycle of its container instance. Configure application-wide
	// observability settings here. Deployment configuration, scaling, placement
	// constraints, versions, and rollouts do not apply.
	SchedulingPolicy ApplicationEditResponseCcDurableObjectApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Application-wide settings for a Durable Object-managed application.
	Configuration ApplicationEditResponseCcDurableObjectApplicationConfiguration `json:"configuration"`
	// Aggregate current activity for the latest observed placement of each instance.
	// Runtime snapshots feed periodic background sweeps. Counts refresh after each
	// complete sweep. Instance listings retain their separate three-month history for
	// failure discovery.
	Health ApplicationEditResponseCcDurableObjectApplicationHealth `json:"health"`
	// Application-wide logging settings for a Durable Object-managed application. The
	// application publishes these settings to its runtime metadata. Updating them does
	// not create a deployment or rollout.
	Observability ApplicationEditResponseCcDurableObjectApplicationObservability `json:"observability"`
	JSON          applicationEditResponseCcDurableObjectApplicationJSON          `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationJSON contains the JSON metadata
// for the struct [ApplicationEditResponseCcDurableObjectApplication]
type applicationEditResponseCcDurableObjectApplicationJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	CreatedAt        apijson.Field
	DurableObjects   apijson.Field
	Name             apijson.Field
	SchedulingPolicy apijson.Field
	UpdatedAt        apijson.Field
	Configuration    apijson.Field
	Health           apijson.Field
	Observability    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationEditResponseCcDurableObjectApplication) implementsApplicationEditResponse() {}

// Durable object configuration using a namespace ID.
type ApplicationEditResponseCcDurableObjectApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                              `json:"namespace_id" api:"required"`
	JSON        applicationEditResponseCcDurableObjectApplicationDurableObjectsJSON `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationDurableObjectsJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationDurableObjects]
type applicationEditResponseCcDurableObjectApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

// Selects a Durable Object-managed application. Each Durable Object creates and
// manages the lifecycle of its container instance. Configure application-wide
// observability settings here. Deployment configuration, scaling, placement
// constraints, versions, and rollouts do not apply.
type ApplicationEditResponseCcDurableObjectApplicationSchedulingPolicy string

const (
	ApplicationEditResponseCcDurableObjectApplicationSchedulingPolicyDurableObject ApplicationEditResponseCcDurableObjectApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationEditResponseCcDurableObjectApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcDurableObjectApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// Application-wide settings for a Durable Object-managed application.
type ApplicationEditResponseCcDurableObjectApplicationConfiguration struct {
	AuthorizedKeys []ApplicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// Configuration properties for connecting with SSH to a container using Wrangler.
	WranglerSSH ApplicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSH `json:"wrangler_ssh"`
	JSON        applicationEditResponseCcDurableObjectApplicationConfigurationJSON        `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationConfigurationJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationConfiguration]
type applicationEditResponseCcDurableObjectApplicationConfigurationJSON struct {
	AuthorizedKeys apijson.Field
	WranglerSSH    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                          `json:"name"`
	JSON applicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKey]
type applicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// Configuration properties for connecting with SSH to a container using Wrangler.
type ApplicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSH struct {
	Enabled bool                                                                          `json:"enabled"`
	Port    int64                                                                         `json:"port"`
	JSON    applicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON
// contains the JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSH]
type applicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON struct {
	Enabled     apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON) RawJSON() string {
	return r.raw
}

// Aggregate current activity for the latest observed placement of each instance.
// Runtime snapshots feed periodic background sweeps. Counts refresh after each
// complete sweep. Instance listings retain their separate three-month history for
// failure discovery.
type ApplicationEditResponseCcDurableObjectApplicationHealth struct {
	// Counts of observed non-terminal instances.
	Instances ApplicationEditResponseCcDurableObjectApplicationHealthInstances `json:"instances" api:"required"`
	// Present as pending until the first activity sweep completes; omitted afterward.
	Summary ApplicationEditResponseCcDurableObjectApplicationHealthSummary `json:"summary"`
	JSON    applicationEditResponseCcDurableObjectApplicationHealthJSON    `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationHealthJSON contains the JSON
// metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationHealth]
type applicationEditResponseCcDurableObjectApplicationHealthJSON struct {
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationHealthJSON) RawJSON() string {
	return r.raw
}

// Counts of observed non-terminal instances.
type ApplicationEditResponseCcDurableObjectApplicationHealthInstances struct {
	// Number of instances whose runtime reports running or stopping.
	Active int64 `json:"active" api:"required"`
	// Number of instances whose runtime reports starting.
	Starting int64                                                                `json:"starting" api:"required"`
	JSON     applicationEditResponseCcDurableObjectApplicationHealthInstancesJSON `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationHealthInstancesJSON contains
// the JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationHealthInstances]
type applicationEditResponseCcDurableObjectApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Starting    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// Present as pending until the first activity sweep completes; omitted afterward.
type ApplicationEditResponseCcDurableObjectApplicationHealthSummary string

const (
	ApplicationEditResponseCcDurableObjectApplicationHealthSummaryPending ApplicationEditResponseCcDurableObjectApplicationHealthSummary = "pending"
)

func (r ApplicationEditResponseCcDurableObjectApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationEditResponseCcDurableObjectApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Application-wide logging settings for a Durable Object-managed application. The
// application publishes these settings to its runtime metadata. Updating them does
// not create a deployment or rollout.
type ApplicationEditResponseCcDurableObjectApplicationObservability struct {
	// Application-wide logging settings.
	Logs ApplicationEditResponseCcDurableObjectApplicationObservabilityLogs `json:"logs"`
	JSON applicationEditResponseCcDurableObjectApplicationObservabilityJSON `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationObservabilityJSON contains the
// JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationObservability]
type applicationEditResponseCcDurableObjectApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Application-wide logging settings.
type ApplicationEditResponseCcDurableObjectApplicationObservabilityLogs struct {
	Enabled bool                                                                   `json:"enabled"`
	JSON    applicationEditResponseCcDurableObjectApplicationObservabilityLogsJSON `json:"-"`
}

// applicationEditResponseCcDurableObjectApplicationObservabilityLogsJSON contains
// the JSON metadata for the struct
// [ApplicationEditResponseCcDurableObjectApplicationObservabilityLogs]
type applicationEditResponseCcDurableObjectApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseCcDurableObjectApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseCcDurableObjectApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationEditResponseSchedulingPolicy string

const (
	ApplicationEditResponseSchedulingPolicyDefault       ApplicationEditResponseSchedulingPolicy = "default"
	ApplicationEditResponseSchedulingPolicyDurableObject ApplicationEditResponseSchedulingPolicy = "durable_object"
)

func (r ApplicationEditResponseSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationEditResponseSchedulingPolicyDefault, ApplicationEditResponseSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// The public Containers API returns an application.
type ApplicationGetResponse struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationGetResponseSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string `json:"active_rollout_id"`
	// This field can have the runtime type of
	// [ApplicationGetResponseCcScheduledApplicationConfiguration],
	// [ApplicationGetResponseCcDurableObjectApplicationConfiguration].
	Configuration interface{} `json:"configuration"`
	// This field can have the runtime type of
	// [ApplicationGetResponseCcScheduledApplicationConstraints].
	Constraints interface{} `json:"constraints"`
	// This field can have the runtime type of
	// [ApplicationGetResponseCcScheduledApplicationDurableObjects],
	// [ApplicationGetResponseCcDurableObjectApplicationDurableObjects].
	DurableObjects interface{} `json:"durable_objects"`
	// This field can have the runtime type of
	// [ApplicationGetResponseCcScheduledApplicationHealth],
	// [ApplicationGetResponseCcDurableObjectApplicationHealth].
	Health interface{} `json:"health"`
	// Number of deployments to create.
	Instances int64 `json:"instances"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// This field can have the runtime type of
	// [ApplicationGetResponseCcScheduledApplicationObservability],
	// [ApplicationGetResponseCcDurableObjectApplicationObservability].
	Observability interface{} `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                      `json:"rollout_active_grace_period"`
	Version                  int64                      `json:"version"`
	JSON                     applicationGetResponseJSON `json:"-"`
	union                    ApplicationGetResponseUnion
}

// applicationGetResponseJSON contains the JSON metadata for the struct
// [ApplicationGetResponse]
type applicationGetResponseJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	CreatedAt                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	ActiveRolloutID          apijson.Field
	Configuration            apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	Instances                apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	Version                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r applicationGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ApplicationGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ApplicationGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ApplicationGetResponseCcScheduledApplication],
// [ApplicationGetResponseCcDurableObjectApplication].
func (r ApplicationGetResponse) AsUnion() ApplicationGetResponseUnion {
	return r.union
}

// The public Containers API returns an application.
//
// Union satisfied by [ApplicationGetResponseCcScheduledApplication] or
// [ApplicationGetResponseCcDurableObjectApplication].
type ApplicationGetResponseUnion interface {
	implementsApplicationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationGetResponseCcScheduledApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationGetResponseCcDurableObjectApplication{}),
		},
	)
}

// Describes an application and the parameters that govern how it places its
// instances.
type ApplicationGetResponseCcScheduledApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// User-specified container configuration.
	Configuration ApplicationGetResponseCcScheduledApplicationConfiguration `json:"configuration" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Number of deployments to create.
	Instances int64 `json:"instances" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// The scheduling policy to use for an application.
	SchedulingPolicy ApplicationGetResponseCcScheduledApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	Version   int64  `json:"version" api:"required"`
	// An identifier for a specific rollout within an application.
	ActiveRolloutID string                                                  `json:"active_rollout_id"`
	Constraints     ApplicationGetResponseCcScheduledApplicationConstraints `json:"constraints"`
	// Durable object configuration stored on and returned from a Cloudchamber
	// application.
	DurableObjects ApplicationGetResponseCcScheduledApplicationDurableObjects `json:"durable_objects"`
	Health         ApplicationGetResponseCcScheduledApplicationHealth         `json:"health"`
	// Maximum number of instances the application allows. This is relevant for
	// applications that auto-scale.
	MaxInstances int64 `json:"max_instances"`
	// Top-level observability settings for the application. This field is mutually
	// exclusive with configuration.observability.
	Observability ApplicationGetResponseCcScheduledApplicationObservability `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod int64                                            `json:"rollout_active_grace_period"`
	JSON                     applicationGetResponseCcScheduledApplicationJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationJSON contains the JSON metadata for
// the struct [ApplicationGetResponseCcScheduledApplication]
type applicationGetResponseCcScheduledApplicationJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	Configuration            apijson.Field
	CreatedAt                apijson.Field
	Instances                apijson.Field
	Name                     apijson.Field
	SchedulingPolicy         apijson.Field
	UpdatedAt                apijson.Field
	Version                  apijson.Field
	ActiveRolloutID          apijson.Field
	Constraints              apijson.Field
	DurableObjects           apijson.Field
	Health                   apijson.Field
	MaxInstances             apijson.Field
	Observability            apijson.Field
	RolloutActiveGracePeriod apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationGetResponseCcScheduledApplication) implementsApplicationGetResponse() {}

// User-specified container configuration.
type ApplicationGetResponseCcScheduledApplicationConfiguration struct {
	// Image url.
	Image          string                                                                   `json:"image" api:"required"`
	AuthorizedKeys []ApplicationGetResponseCcScheduledApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command []string `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint []string `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables []ApplicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariable `json:"environment_variables"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationGetResponseCcScheduledApplicationConfigurationObservability `json:"observability"`
	JSON          applicationGetResponseCcScheduledApplicationConfigurationJSON          `json:"-"`
}

// applicationGetResponseCcScheduledApplicationConfigurationJSON contains the JSON
// metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationConfiguration]
type applicationGetResponseCcScheduledApplicationConfigurationJSON struct {
	Image                apijson.Field
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationGetResponseCcScheduledApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                     `json:"name"`
	JSON applicationGetResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationConfigurationAuthorizedKey]
type applicationGetResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                                           `json:"value" api:"required"`
	JSON  applicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariable]
type applicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationConfigurationEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType string

const (
	ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeLite      ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType = "lite"
	ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeBasic     ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType = "basic"
	ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard1 ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType = "standard-1"
	ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard2 ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType = "standard-2"
	ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard3 ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType = "standard-3"
	ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard4 ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType = "standard-4"
)

func (r ApplicationGetResponseCcScheduledApplicationConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeLite, ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeBasic, ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard1, ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard2, ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard3, ApplicationGetResponseCcScheduledApplicationConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationGetResponseCcScheduledApplicationConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationGetResponseCcScheduledApplicationConfigurationObservabilityLogs `json:"logs"`
	JSON applicationGetResponseCcScheduledApplicationConfigurationObservabilityJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationConfigurationObservabilityJSON
// contains the JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationConfigurationObservability]
type applicationGetResponseCcScheduledApplicationConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationGetResponseCcScheduledApplicationConfigurationObservabilityLogs struct {
	Enabled bool                                                                           `json:"enabled"`
	JSON    applicationGetResponseCcScheduledApplicationConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationConfigurationObservabilityLogsJSON
// contains the JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationConfigurationObservabilityLogs]
type applicationGetResponseCcScheduledApplicationConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationGetResponseCcScheduledApplicationSchedulingPolicy string

const (
	ApplicationGetResponseCcScheduledApplicationSchedulingPolicyDefault       ApplicationGetResponseCcScheduledApplicationSchedulingPolicy = "default"
	ApplicationGetResponseCcScheduledApplicationSchedulingPolicyDurableObject ApplicationGetResponseCcScheduledApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationGetResponseCcScheduledApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcScheduledApplicationSchedulingPolicyDefault, ApplicationGetResponseCcScheduledApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

type ApplicationGetResponseCcScheduledApplicationConstraints struct {
	// Restricts placement to datacenters in the selected jurisdiction. Choose "eu",
	// "fedramp", or "us". When combined with regions, EU supports EEUR and WEUR while
	// FedRAMP and US support ENAM and WNAM.
	Jurisdiction string                                                      `json:"jurisdiction"`
	Regions      []string                                                    `json:"regions"`
	JSON         applicationGetResponseCcScheduledApplicationConstraintsJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationConstraintsJSON contains the JSON
// metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationConstraints]
type applicationGetResponseCcScheduledApplicationConstraintsJSON struct {
	Jurisdiction apijson.Field
	Regions      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationConstraintsJSON) RawJSON() string {
	return r.raw
}

// Durable object configuration stored on and returned from a Cloudchamber
// application.
type ApplicationGetResponseCcScheduledApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                         `json:"namespace_id" api:"required"`
	JSON        applicationGetResponseCcScheduledApplicationDurableObjectsJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationDurableObjectsJSON contains the JSON
// metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationDurableObjects]
type applicationGetResponseCcScheduledApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetResponseCcScheduledApplicationHealth struct {
	Errors []ApplicationGetResponseCcScheduledApplicationHealthError `json:"errors" api:"required"`
	// Shows a count of application instance states.
	Instances ApplicationGetResponseCcScheduledApplicationHealthInstances `json:"instances" api:"required"`
	// High-level health assessment. Only populated for "new_instances" strategy. Based
	// on a sample of target-version instances rather than a full count.
	//
	// - "pending": Zero target-version instances exist yet.
	// - "healthy": Every sampled target-version instance reports running or active.
	// - "degraded": Some sampled instances remain starting or scheduling.
	// - "unhealthy": One or more sampled instances have failed.
	Summary ApplicationGetResponseCcScheduledApplicationHealthSummary `json:"summary"`
	JSON    applicationGetResponseCcScheduledApplicationHealthJSON    `json:"-"`
}

// applicationGetResponseCcScheduledApplicationHealthJSON contains the JSON
// metadata for the struct [ApplicationGetResponseCcScheduledApplicationHealth]
type applicationGetResponseCcScheduledApplicationHealthJSON struct {
	Errors      apijson.Field
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationHealthJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetResponseCcScheduledApplicationHealthError struct {
	// An event within a Placement or a Job.
	Event ApplicationGetResponseCcScheduledApplicationHealthErrorsEvent `json:"event" api:"required"`
	// An instance ID represents an identifier of an instance configuration that
	// maintains an underlying placement.
	InstanceID string                                                      `json:"instance_id" api:"required"`
	JSON       applicationGetResponseCcScheduledApplicationHealthErrorJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationHealthErrorJSON contains the JSON
// metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationHealthError]
type applicationGetResponseCcScheduledApplicationHealthErrorJSON struct {
	Event       apijson.Field
	InstanceID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationHealthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationHealthErrorJSON) RawJSON() string {
	return r.raw
}

// An event within a Placement or a Job.
type ApplicationGetResponseCcScheduledApplicationHealthErrorsEvent struct {
	ID      string                 `json:"id" api:"required"`
	Details map[string]interface{} `json:"details" api:"required"`
	Message string                 `json:"message" api:"required"`
	// Name of the event that describes the kind event that happened.
	//
	// - SchedulerPlaced: It's the first event that creates a container placement. It
	//   happens when the Containers runtime was able to retrieve deployment resources
	//   and start verifying everything is correct.
	// - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
	//   container.
	// - VMStarted: It's sent when the Containers runtime starts the VM. The container
	//   might remain unhealthy at this point.
	// - ImagePulled: It's sent when the Containers runtime pulls the image
	//   successfully.
	// - ImagePullError: It's sent when the Containers runtime is having issues pulling
	//   the image. The message and details have more information on what happened for
	//   debugging.
	// - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
	//   VM.
	// - VMStopping: It's sent when the scheduler is stopping the VM.
	// - VMStopped: It's sent when the VM finally exits.
	// - VMFailed: It's sent when the scheduling of the VM failed in the current
	//   location.
	// - RuntimeStartFailed: It's sent when the runtime hits an internal error.
	// - SSHStarted: It's sent when the container gains network connectivity and opens
	//   the SSH port. Containers only send this event when SSH keys exist.
	// - CheckUpdate: Sent when the status of a health or readiness check changes. This
	//   may also affect the health status of the placement.
	// - DurableObjectConnected: Sent when a durable object instance connects and gains
	//   control of the deployment. This event is only sent for durable object
	//   deployments. It is sent after VMStarted.
	// - ContainerStarted: It's sent when the container starts running.
	Name         ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName `json:"name" api:"required"`
	StatusChange map[string]interface{}                                            `json:"statusChange" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	Time string                                                            `json:"time" api:"required"`
	Type ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType `json:"type" api:"required"`
	JSON applicationGetResponseCcScheduledApplicationHealthErrorsEventJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationHealthErrorsEventJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationHealthErrorsEvent]
type applicationGetResponseCcScheduledApplicationHealthErrorsEventJSON struct {
	ID           apijson.Field
	Details      apijson.Field
	Message      apijson.Field
	Name         apijson.Field
	StatusChange apijson.Field
	Time         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationHealthErrorsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationHealthErrorsEventJSON) RawJSON() string {
	return r.raw
}

// Name of the event that describes the kind event that happened.
//
//   - SchedulerPlaced: It's the first event that creates a container placement. It
//     happens when the Containers runtime was able to retrieve deployment resources
//     and start verifying everything is correct.
//   - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
//     container.
//   - VMStarted: It's sent when the Containers runtime starts the VM. The container
//     might remain unhealthy at this point.
//   - ImagePulled: It's sent when the Containers runtime pulls the image
//     successfully.
//   - ImagePullError: It's sent when the Containers runtime is having issues pulling
//     the image. The message and details have more information on what happened for
//     debugging.
//   - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
//     VM.
//   - VMStopping: It's sent when the scheduler is stopping the VM.
//   - VMStopped: It's sent when the VM finally exits.
//   - VMFailed: It's sent when the scheduling of the VM failed in the current
//     location.
//   - RuntimeStartFailed: It's sent when the runtime hits an internal error.
//   - SSHStarted: It's sent when the container gains network connectivity and opens
//     the SSH port. Containers only send this event when SSH keys exist.
//   - CheckUpdate: Sent when the status of a health or readiness check changes. This
//     may also affect the health status of the placement.
//   - DurableObjectConnected: Sent when a durable object instance connects and gains
//     control of the deployment. This event is only sent for durable object
//     deployments. It is sent after VMStarted.
//   - ContainerStarted: It's sent when the container starts running.
type ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName string

const (
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced              ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "SchedulerPlaced"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned         ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssigned"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmStarted                    ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "VMStarted"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameImagePulled                  ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "ImagePulled"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameImagePullError               ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "ImagePullError"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart              ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "VMFailedToStart"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "NetworkingIPAssignmentFailed"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmRunning                    ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "VMRunning"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmStopping                   ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "VMStopping"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmStopped                    ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "VMStopped"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmFailed                     ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "VMFailed"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed           ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "RuntimeStartFailed"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted                   ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "SSHStarted"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates         ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "ServiceHealthUpdates"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate                  ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "CheckUpdate"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected       ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "DurableObjectConnected"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted             ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName = "ContainerStarted"
)

func (r ApplicationGetResponseCcScheduledApplicationHealthErrorsEventName) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameSchedulerPlaced, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssigned, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmStarted, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameImagePulled, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameImagePullError, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmFailedToStart, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameNetworkingIPAssignmentFailed, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmRunning, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmStopping, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmStopped, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameVmFailed, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameRuntimeStartFailed, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameSSHStarted, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameServiceHealthUpdates, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameCheckUpdate, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameDurableObjectConnected, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventNameContainerStarted:
		return true
	}
	return false
}

type ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType string

const (
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeInfo        ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType = "Info"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeError       ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType = "Error"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeWarn        ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType = "Warn"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeUserError   ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType = "UserError"
	ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeSystemError ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType = "SystemError"
)

func (r ApplicationGetResponseCcScheduledApplicationHealthErrorsEventType) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeInfo, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeError, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeWarn, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeUserError, ApplicationGetResponseCcScheduledApplicationHealthErrorsEventTypeSystemError:
		return true
	}
	return false
}

// Shows a count of application instance states.
type ApplicationGetResponseCcScheduledApplicationHealthInstances struct {
	// Number of instances whose runtime reports the container as running
	// (container_status = "running"). This is a subset of the placements that remain
	// up: an instance that is already bound to a Durable Object and serving traffic is
	// counted under "assigned" until its container_status catches up to "running", so
	// container_status can briefly lag Durable Object attachment under churn. To
	// estimate running, Durable-Object-bound instances, sum "active" + "assigned"
	// rather than reading "active" alone.
	Active int64 `json:"active" api:"required"`
	// Number of instances bound to a Durable Object with a running placement whose
	// container_status remains behind "running". These count as live, serving
	// instances; "active" + "assigned" approximates the running, Durable-Object-bound
	// count.
	Assigned int64                                                           `json:"assigned" api:"required"`
	JSON     applicationGetResponseCcScheduledApplicationHealthInstancesJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationHealthInstancesJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationHealthInstances]
type applicationGetResponseCcScheduledApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Assigned    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// High-level health assessment. Only populated for "new_instances" strategy. Based
// on a sample of target-version instances rather than a full count.
//
// - "pending": Zero target-version instances exist yet.
// - "healthy": Every sampled target-version instance reports running or active.
// - "degraded": Some sampled instances remain starting or scheduling.
// - "unhealthy": One or more sampled instances have failed.
type ApplicationGetResponseCcScheduledApplicationHealthSummary string

const (
	ApplicationGetResponseCcScheduledApplicationHealthSummaryHealthy   ApplicationGetResponseCcScheduledApplicationHealthSummary = "healthy"
	ApplicationGetResponseCcScheduledApplicationHealthSummaryDegraded  ApplicationGetResponseCcScheduledApplicationHealthSummary = "degraded"
	ApplicationGetResponseCcScheduledApplicationHealthSummaryUnhealthy ApplicationGetResponseCcScheduledApplicationHealthSummary = "unhealthy"
	ApplicationGetResponseCcScheduledApplicationHealthSummaryPending   ApplicationGetResponseCcScheduledApplicationHealthSummary = "pending"
)

func (r ApplicationGetResponseCcScheduledApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcScheduledApplicationHealthSummaryHealthy, ApplicationGetResponseCcScheduledApplicationHealthSummaryDegraded, ApplicationGetResponseCcScheduledApplicationHealthSummaryUnhealthy, ApplicationGetResponseCcScheduledApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Top-level observability settings for the application. This field is mutually
// exclusive with configuration.observability.
type ApplicationGetResponseCcScheduledApplicationObservability struct {
	// Observability logging settings.
	Logs ApplicationGetResponseCcScheduledApplicationObservabilityLogs `json:"logs"`
	JSON applicationGetResponseCcScheduledApplicationObservabilityJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationObservabilityJSON contains the JSON
// metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationObservability]
type applicationGetResponseCcScheduledApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationGetResponseCcScheduledApplicationObservabilityLogs struct {
	Enabled bool                                                              `json:"enabled"`
	JSON    applicationGetResponseCcScheduledApplicationObservabilityLogsJSON `json:"-"`
}

// applicationGetResponseCcScheduledApplicationObservabilityLogsJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcScheduledApplicationObservabilityLogs]
type applicationGetResponseCcScheduledApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcScheduledApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcScheduledApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Each Durable Object creates and manages the lifecycle of its container instance.
type ApplicationGetResponseCcDurableObjectApplication struct {
	// An Application ID represents an identifier of an application.
	ID string `json:"id" api:"required"`
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// Durable object configuration using a namespace ID.
	DurableObjects ApplicationGetResponseCcDurableObjectApplicationDurableObjects `json:"durable_objects" api:"required"`
	// The application name.
	Name string `json:"name" api:"required"`
	// Selects a Durable Object-managed application. Each Durable Object creates and
	// manages the lifecycle of its container instance. Configure application-wide
	// observability settings here. Deployment configuration, scaling, placement
	// constraints, versions, and rollouts do not apply.
	SchedulingPolicy ApplicationGetResponseCcDurableObjectApplicationSchedulingPolicy `json:"scheduling_policy" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Application-wide settings for a Durable Object-managed application.
	Configuration ApplicationGetResponseCcDurableObjectApplicationConfiguration `json:"configuration"`
	// Aggregate current activity for the latest observed placement of each instance.
	// Runtime snapshots feed periodic background sweeps. Counts refresh after each
	// complete sweep. Instance listings retain their separate three-month history for
	// failure discovery.
	Health ApplicationGetResponseCcDurableObjectApplicationHealth `json:"health"`
	// Application-wide logging settings for a Durable Object-managed application. The
	// application publishes these settings to its runtime metadata. Updating them does
	// not create a deployment or rollout.
	Observability ApplicationGetResponseCcDurableObjectApplicationObservability `json:"observability"`
	JSON          applicationGetResponseCcDurableObjectApplicationJSON          `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationJSON contains the JSON metadata
// for the struct [ApplicationGetResponseCcDurableObjectApplication]
type applicationGetResponseCcDurableObjectApplicationJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	CreatedAt        apijson.Field
	DurableObjects   apijson.Field
	Name             apijson.Field
	SchedulingPolicy apijson.Field
	UpdatedAt        apijson.Field
	Configuration    apijson.Field
	Health           apijson.Field
	Observability    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationGetResponseCcDurableObjectApplication) implementsApplicationGetResponse() {}

// Durable object configuration using a namespace ID.
type ApplicationGetResponseCcDurableObjectApplicationDurableObjects struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID string                                                             `json:"namespace_id" api:"required"`
	JSON        applicationGetResponseCcDurableObjectApplicationDurableObjectsJSON `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationDurableObjectsJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationDurableObjects]
type applicationGetResponseCcDurableObjectApplicationDurableObjectsJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationDurableObjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationDurableObjectsJSON) RawJSON() string {
	return r.raw
}

// Selects a Durable Object-managed application. Each Durable Object creates and
// manages the lifecycle of its container instance. Configure application-wide
// observability settings here. Deployment configuration, scaling, placement
// constraints, versions, and rollouts do not apply.
type ApplicationGetResponseCcDurableObjectApplicationSchedulingPolicy string

const (
	ApplicationGetResponseCcDurableObjectApplicationSchedulingPolicyDurableObject ApplicationGetResponseCcDurableObjectApplicationSchedulingPolicy = "durable_object"
)

func (r ApplicationGetResponseCcDurableObjectApplicationSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcDurableObjectApplicationSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// Application-wide settings for a Durable Object-managed application.
type ApplicationGetResponseCcDurableObjectApplicationConfiguration struct {
	AuthorizedKeys []ApplicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKey `json:"authorized_keys"`
	// Configuration properties for connecting with SSH to a container using Wrangler.
	WranglerSSH ApplicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSH `json:"wrangler_ssh"`
	JSON        applicationGetResponseCcDurableObjectApplicationConfigurationJSON        `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationConfigurationJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationConfiguration]
type applicationGetResponseCcDurableObjectApplicationConfigurationJSON struct {
	AuthorizedKeys apijson.Field
	WranglerSSH    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                                         `json:"name"`
	JSON applicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON
// contains the JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKey]
type applicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// Configuration properties for connecting with SSH to a container using Wrangler.
type ApplicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSH struct {
	Enabled bool                                                                         `json:"enabled"`
	Port    int64                                                                        `json:"port"`
	JSON    applicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON
// contains the JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSH]
type applicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON struct {
	Enabled     apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationConfigurationWranglerSSHJSON) RawJSON() string {
	return r.raw
}

// Aggregate current activity for the latest observed placement of each instance.
// Runtime snapshots feed periodic background sweeps. Counts refresh after each
// complete sweep. Instance listings retain their separate three-month history for
// failure discovery.
type ApplicationGetResponseCcDurableObjectApplicationHealth struct {
	// Counts of observed non-terminal instances.
	Instances ApplicationGetResponseCcDurableObjectApplicationHealthInstances `json:"instances" api:"required"`
	// Present as pending until the first activity sweep completes; omitted afterward.
	Summary ApplicationGetResponseCcDurableObjectApplicationHealthSummary `json:"summary"`
	JSON    applicationGetResponseCcDurableObjectApplicationHealthJSON    `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationHealthJSON contains the JSON
// metadata for the struct [ApplicationGetResponseCcDurableObjectApplicationHealth]
type applicationGetResponseCcDurableObjectApplicationHealthJSON struct {
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationHealthJSON) RawJSON() string {
	return r.raw
}

// Counts of observed non-terminal instances.
type ApplicationGetResponseCcDurableObjectApplicationHealthInstances struct {
	// Number of instances whose runtime reports running or stopping.
	Active int64 `json:"active" api:"required"`
	// Number of instances whose runtime reports starting.
	Starting int64                                                               `json:"starting" api:"required"`
	JSON     applicationGetResponseCcDurableObjectApplicationHealthInstancesJSON `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationHealthInstancesJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationHealthInstances]
type applicationGetResponseCcDurableObjectApplicationHealthInstancesJSON struct {
	Active      apijson.Field
	Starting    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// Present as pending until the first activity sweep completes; omitted afterward.
type ApplicationGetResponseCcDurableObjectApplicationHealthSummary string

const (
	ApplicationGetResponseCcDurableObjectApplicationHealthSummaryPending ApplicationGetResponseCcDurableObjectApplicationHealthSummary = "pending"
)

func (r ApplicationGetResponseCcDurableObjectApplicationHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationGetResponseCcDurableObjectApplicationHealthSummaryPending:
		return true
	}
	return false
}

// Application-wide logging settings for a Durable Object-managed application. The
// application publishes these settings to its runtime metadata. Updating them does
// not create a deployment or rollout.
type ApplicationGetResponseCcDurableObjectApplicationObservability struct {
	// Application-wide logging settings.
	Logs ApplicationGetResponseCcDurableObjectApplicationObservabilityLogs `json:"logs"`
	JSON applicationGetResponseCcDurableObjectApplicationObservabilityJSON `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationObservabilityJSON contains the
// JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationObservability]
type applicationGetResponseCcDurableObjectApplicationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Application-wide logging settings.
type ApplicationGetResponseCcDurableObjectApplicationObservabilityLogs struct {
	Enabled bool                                                                  `json:"enabled"`
	JSON    applicationGetResponseCcDurableObjectApplicationObservabilityLogsJSON `json:"-"`
}

// applicationGetResponseCcDurableObjectApplicationObservabilityLogsJSON contains
// the JSON metadata for the struct
// [ApplicationGetResponseCcDurableObjectApplicationObservabilityLogs]
type applicationGetResponseCcDurableObjectApplicationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseCcDurableObjectApplicationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseCcDurableObjectApplicationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// The scheduling policy to use for an application.
type ApplicationGetResponseSchedulingPolicy string

const (
	ApplicationGetResponseSchedulingPolicyDefault       ApplicationGetResponseSchedulingPolicy = "default"
	ApplicationGetResponseSchedulingPolicyDurableObject ApplicationGetResponseSchedulingPolicy = "durable_object"
)

func (r ApplicationGetResponseSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationGetResponseSchedulingPolicyDefault, ApplicationGetResponseSchedulingPolicyDurableObject:
		return true
	}
	return false
}

type ApplicationNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Create a Containers application. Set `scheduling_policy` to `default` for a
	// scheduler-backed application with deployment configuration, instance counts,
	// constraints, versions, and rollouts. Set it to `durable_object` for a Durable
	// Object-managed application where each Durable Object creates and manages the
	// lifecycle of its container instance. For `durable_object` requests, supply
	// `name`, `scheduling_policy`, and `durable_objects`, with optional
	// `configuration` and top-level `observability` settings.
	Body ApplicationNewParamsBodyUnion `json:"body" api:"required"`
}

func (r ApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Create a Containers application. Set `scheduling_policy` to `default` for a
// scheduler-backed application with deployment configuration, instance counts,
// constraints, versions, and rollouts. Set it to `durable_object` for a Durable
// Object-managed application where each Durable Object creates and manages the
// lifecycle of its container instance. For `durable_object` requests, supply
// `name`, `scheduling_policy`, and `durable_objects`, with optional
// `configuration` and top-level `observability` settings.
type ApplicationNewParamsBody struct {
	// The name for this application.
	Name param.Field[string] `json:"name" api:"required"`
	// Selects a scheduler-backed application. Use `default` when the Containers
	// scheduler should maintain the requested number of instances and manage
	// deployment configuration, placement, scaling, versions, and rollouts.
	SchedulingPolicy param.Field[ApplicationNewParamsBodySchedulingPolicy] `json:"scheduling_policy" api:"required"`
	Configuration    param.Field[interface{}]                              `json:"configuration"`
	Constraints      param.Field[interface{}]                              `json:"constraints"`
	DurableObjects   param.Field[interface{}]                              `json:"durable_objects"`
	// The initial number of deployments to create.
	Instances param.Field[int64] `json:"instances"`
	// Sets the maximum number of instances that the application can run.
	MaxInstances  param.Field[int64]       `json:"max_instances"`
	Observability param.Field[interface{}] `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod param.Field[int64] `json:"rollout_active_grace_period"`
}

func (r ApplicationNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBody) implementsApplicationNewParamsBodyUnion() {}

// Create a Containers application. Set `scheduling_policy` to `default` for a
// scheduler-backed application with deployment configuration, instance counts,
// constraints, versions, and rollouts. Set it to `durable_object` for a Durable
// Object-managed application where each Durable Object creates and manages the
// lifecycle of its container instance. For `durable_object` requests, supply
// `name`, `scheduling_policy`, and `durable_objects`, with optional
// `configuration` and top-level `observability` settings.
//
// Satisfied by
// [containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequest],
// [containers.ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequest],
// [ApplicationNewParamsBody].
type ApplicationNewParamsBodyUnion interface {
	implementsApplicationNewParamsBodyUnion()
}

// Create a scheduler-backed Containers application. The Containers scheduler
// maintains the requested instance count and applies deployment configuration,
// placement, scaling, versions, and rollouts.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequest struct {
	// Defines the deployment configuration for every deployment in this application.
	Configuration param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfiguration] `json:"configuration" api:"required"`
	// The initial number of deployments to create.
	Instances param.Field[int64] `json:"instances" api:"required"`
	// Sets the maximum number of instances that the application can run.
	MaxInstances param.Field[int64] `json:"max_instances" api:"required"`
	// The name for this application.
	Name param.Field[string] `json:"name" api:"required"`
	// Selects a scheduler-backed application. Use `default` when the Containers
	// scheduler should maintain the requested number of instances and manage
	// deployment configuration, placement, scaling, versions, and rollouts.
	SchedulingPolicy param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicy] `json:"scheduling_policy" api:"required"`
	Constraints      param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConstraints]      `json:"constraints"`
	// Optionally associates this scheduler-backed application with a Durable Object
	// namespace.
	DurableObjects param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion] `json:"durable_objects"`
	// Top-level observability settings for the application. This field is mutually
	// exclusive with configuration.observability.
	Observability param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservability] `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod param.Field[int64] `json:"rollout_active_grace_period"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequest) implementsApplicationNewParamsBodyUnion() {
}

// Defines the deployment configuration for every deployment in this application.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfiguration struct {
	// Image url.
	Image          param.Field[string]                                                                                            `json:"image" api:"required"`
	AuthorizedKeys param.Field[[]ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationAuthorizedKey] `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command param.Field[[]string] `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint param.Field[[]string] `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables param.Field[[]ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationEnvironmentVariable] `json:"environment_variables"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType] `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservability] `json:"observability"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User-provided SSH public key.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey param.Field[string] `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name param.Field[string] `json:"name"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationAuthorizedKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An environment variable with a value set.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name param.Field[string] `json:"name" api:"required"`
	// An environment variable value.
	Value param.Field[string] `json:"value" api:"required"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType string

const (
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeLite      ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType = "lite"
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeBasic     ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType = "basic"
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard1 ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType = "standard-1"
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard2 ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType = "standard-2"
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard3 ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType = "standard-3"
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard4 ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType = "standard-4"
)

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeLite, ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeBasic, ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard1, ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard2, ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard3, ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservability struct {
	// Observability logging settings.
	Logs param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservabilityLogs] `json:"logs"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability logging settings.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservabilityLogs struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Selects a scheduler-backed application. Use `default` when the Containers
// scheduler should maintain the requested number of instances and manage
// deployment configuration, placement, scaling, versions, and rollouts.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicy string

const (
	ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicyDefault ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicy = "default"
)

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicyDefault:
		return true
	}
	return false
}

type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConstraints struct {
	// Restricts placement to datacenters in the selected jurisdiction. Choose "eu",
	// "fedramp", or "us". When combined with regions, EU supports EEUR and WEUR while
	// FedRAMP and US support ENAM and WNAM.
	Jurisdiction param.Field[string]   `json:"jurisdiction"`
	Regions      param.Field[[]string] `json:"regions"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConstraints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optionally associates this scheduler-backed application with a Durable Object
// namespace.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjects struct {
	// The class name of the durable object.
	ClassName param.Field[string] `json:"class_name"`
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script name where the durable object class is defined.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjects) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjects) implementsApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion() {
}

// Optionally associates this scheduler-backed application with a Durable Object
// namespace.
//
// Satisfied by
// [containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID],
// [containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass],
// [ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjects].
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion interface {
	implementsApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion()
}

// Durable object configuration using a namespace ID.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID) implementsApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion() {
}

// Durable object configuration using script and class names.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass struct {
	// The class name of the durable object.
	ClassName param.Field[string] `json:"class_name" api:"required"`
	// The script name where the durable object class is defined.
	ScriptName param.Field[string] `json:"script_name" api:"required"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass) implementsApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion() {
}

// Top-level observability settings for the application. This field is mutually
// exclusive with configuration.observability.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservability struct {
	// Observability logging settings.
	Logs param.Field[ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservabilityLogs] `json:"logs"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability logging settings.
type ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservabilityLogs struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Create a Durable Object-managed Containers application. Each Durable Object
// creates and manages the lifecycle of its container instance. Supply `name`,
// `scheduling_policy`, and `durable_objects`, with optional `configuration` and
// top-level `observability` settings. Deployment configuration, instance counts,
// scaling, constraints, versions, and rollouts do not apply.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequest struct {
	// The customer-owned Durable Object namespace that owns this application and its
	// instances.
	DurableObjects param.Field[ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsUnion] `json:"durable_objects" api:"required"`
	// The name for this application.
	Name param.Field[string] `json:"name" api:"required"`
	// Selects a Durable Object-managed application. Each Durable Object creates and
	// manages the lifecycle of its container instance. Configure application-wide
	// observability settings here. Deployment configuration, scaling, placement
	// constraints, versions, and rollouts do not apply.
	SchedulingPolicy param.Field[ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestSchedulingPolicy] `json:"scheduling_policy" api:"required"`
	// Configuration for a Durable Object-managed application.
	Configuration param.Field[ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfiguration] `json:"configuration"`
	// Application-wide logging settings for a Durable Object-managed application. The
	// application publishes these settings to its runtime metadata. Updating them does
	// not create a deployment or rollout.
	Observability param.Field[ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestObservability] `json:"observability"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequest) implementsApplicationNewParamsBodyUnion() {
}

// The customer-owned Durable Object namespace that owns this application and its
// instances.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjects struct {
	// The class name of the durable object.
	ClassName param.Field[string] `json:"class_name"`
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script name where the durable object class is defined.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjects) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjects) implementsApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsUnion() {
}

// The customer-owned Durable Object namespace that owns this application and its
// instances.
//
// Satisfied by
// [containers.ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID],
// [containers.ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass],
// [ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjects].
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsUnion interface {
	implementsApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsUnion()
}

// Durable object configuration using a namespace ID.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID struct {
	// The namespace ID of the durable object namespace to use for this application.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID) implementsApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsUnion() {
}

// Durable object configuration using script and class names.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass struct {
	// The class name of the durable object.
	ClassName param.Field[string] `json:"class_name" api:"required"`
	// The script name where the durable object class is defined.
	ScriptName param.Field[string] `json:"script_name" api:"required"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsCcDurableObjectsConfigurationScriptAndClass) implementsApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestDurableObjectsUnion() {
}

// Selects a Durable Object-managed application. Each Durable Object creates and
// manages the lifecycle of its container instance. Configure application-wide
// observability settings here. Deployment configuration, scaling, placement
// constraints, versions, and rollouts do not apply.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestSchedulingPolicy string

const (
	ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestSchedulingPolicyDurableObject ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestSchedulingPolicy = "durable_object"
)

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestSchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestSchedulingPolicyDurableObject:
		return true
	}
	return false
}

// Configuration for a Durable Object-managed application.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfiguration struct {
	AuthorizedKeys param.Field[[]ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfigurationAuthorizedKey] `json:"authorized_keys"`
	// Configuration properties for connecting with SSH to a container using Wrangler.
	WranglerSSH param.Field[ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfigurationWranglerSSH] `json:"wrangler_ssh"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User-provided SSH public key.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey param.Field[string] `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name param.Field[string] `json:"name"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfigurationAuthorizedKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration properties for connecting with SSH to a container using Wrangler.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfigurationWranglerSSH struct {
	Enabled param.Field[bool]  `json:"enabled"`
	Port    param.Field[int64] `json:"port"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestConfigurationWranglerSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Application-wide logging settings for a Durable Object-managed application. The
// application publishes these settings to its runtime metadata. Updating them does
// not create a deployment or rollout.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestObservability struct {
	// Application-wide logging settings.
	Logs param.Field[ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestObservabilityLogs] `json:"logs"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Application-wide logging settings.
type ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestObservabilityLogs struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ApplicationNewParamsBodyCcContainersCreateDurableObjectApplicationRequestObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Selects a scheduler-backed application. Use `default` when the Containers
// scheduler should maintain the requested number of instances and manage
// deployment configuration, placement, scaling, versions, and rollouts.
type ApplicationNewParamsBodySchedulingPolicy string

const (
	ApplicationNewParamsBodySchedulingPolicyDefault       ApplicationNewParamsBodySchedulingPolicy = "default"
	ApplicationNewParamsBodySchedulingPolicyDurableObject ApplicationNewParamsBodySchedulingPolicy = "durable_object"
)

func (r ApplicationNewParamsBodySchedulingPolicy) IsKnown() bool {
	switch r {
	case ApplicationNewParamsBodySchedulingPolicyDefault, ApplicationNewParamsBodySchedulingPolicyDurableObject:
		return true
	}
	return false
}

type ApplicationNewResponseEnvelope struct {
	Errors   []ApplicationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ApplicationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// The public Containers API returns an application.
	Result ApplicationNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                               `json:"success" api:"required"`
	JSON    applicationNewResponseEnvelopeJSON `json:"-"`
}

// applicationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ApplicationNewResponseEnvelope]
type applicationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ApplicationNewResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ApplicationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             applicationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// applicationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ApplicationNewResponseEnvelopeErrors]
type applicationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ApplicationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    applicationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// applicationNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ApplicationNewResponseEnvelopeErrorsSource]
type applicationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationNewResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ApplicationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             applicationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// applicationNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ApplicationNewResponseEnvelopeMessages]
type applicationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ApplicationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    applicationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// applicationNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ApplicationNewResponseEnvelopeMessagesSource]
type applicationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter applications by image.
	Image param.Field[string] `query:"image"`
	// Filter applications by name.
	Name param.Field[string] `query:"name"`
	// Opaque token from a previous response to retrieve the next page.
	PageToken param.Field[string] `query:"page_token"`
	// Maximum number of applications to return per page. Defaults to all, or 100 when
	// `page_token` is set.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ApplicationListParams]'s query parameters as `url.Values`.
func (r ApplicationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ApplicationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ApplicationDeleteResponseEnvelope struct {
	Errors   []ApplicationDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ApplicationDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Result of starting asynchronous deletion for a Containers application.
	Result ApplicationDeleteResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                  `json:"success" api:"required"`
	JSON    applicationDeleteResponseEnvelopeJSON `json:"-"`
}

// applicationDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ApplicationDeleteResponseEnvelope]
type applicationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ApplicationDeleteResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ApplicationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             applicationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// applicationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ApplicationDeleteResponseEnvelopeErrors]
type applicationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ApplicationDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    applicationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// applicationDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ApplicationDeleteResponseEnvelopeErrorsSource]
type applicationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationDeleteResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ApplicationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             applicationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// applicationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ApplicationDeleteResponseEnvelopeMessages]
type applicationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ApplicationDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    applicationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// applicationDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ApplicationDeleteResponseEnvelopeMessagesSource]
type applicationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Application configuration fields you can change without creating a rollout.
	Configuration param.Field[ApplicationEditParamsConfiguration] `json:"configuration"`
	Constraints   param.Field[ApplicationEditParamsConstraints]   `json:"constraints"`
	// Maximum number of instances that an autoscaling application can run.
	MaxInstances param.Field[int64] `json:"max_instances"`
	// Top-level application observability settings. Scheduler-backed applications
	// hot-reload these settings across existing instances. An existing Durable
	// Object-managed application accepts only `logs.enabled` and publishes these
	// settings to runtime metadata without creating deployments or rollouts.
	Observability param.Field[ApplicationEditParamsObservability] `json:"observability"`
	// Grace period for active instances to stay alive before becoming eligible for
	// shutdown signal due to a rollout, in seconds. Defaults to 0.
	RolloutActiveGracePeriod param.Field[int64] `json:"rollout_active_grace_period"`
}

func (r ApplicationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Application configuration fields you can change without creating a rollout.
type ApplicationEditParamsConfiguration struct {
	AuthorizedKeys param.Field[[]ApplicationEditParamsConfigurationAuthorizedKey] `json:"authorized_keys"`
	// Configuration properties for connecting with SSH to a container using Wrangler.
	WranglerSSH param.Field[ApplicationEditParamsConfigurationWranglerSSH] `json:"wrangler_ssh"`
}

func (r ApplicationEditParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User-provided SSH public key.
type ApplicationEditParamsConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey param.Field[string] `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name param.Field[string] `json:"name"`
}

func (r ApplicationEditParamsConfigurationAuthorizedKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration properties for connecting with SSH to a container using Wrangler.
type ApplicationEditParamsConfigurationWranglerSSH struct {
	Enabled param.Field[bool]  `json:"enabled"`
	Port    param.Field[int64] `json:"port"`
}

func (r ApplicationEditParamsConfigurationWranglerSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationEditParamsConstraints struct {
	// Restricts placement to datacenters in the selected jurisdiction. Choose "eu",
	// "fedramp", or "us". When combined with regions, EU supports EEUR and WEUR while
	// FedRAMP and US support ENAM and WNAM.
	Jurisdiction param.Field[string]   `json:"jurisdiction"`
	Regions      param.Field[[]string] `json:"regions"`
}

func (r ApplicationEditParamsConstraints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Top-level application observability settings. Scheduler-backed applications
// hot-reload these settings across existing instances. An existing Durable
// Object-managed application accepts only `logs.enabled` and publishes these
// settings to runtime metadata without creating deployments or rollouts.
type ApplicationEditParamsObservability struct {
	// Observability logging settings.
	Logs param.Field[ApplicationEditParamsObservabilityLogs] `json:"logs"`
}

func (r ApplicationEditParamsObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability logging settings.
type ApplicationEditParamsObservabilityLogs struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ApplicationEditParamsObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationEditResponseEnvelope struct {
	Errors   []ApplicationEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ApplicationEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// The public Containers API returns an application.
	Result ApplicationEditResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                `json:"success" api:"required"`
	JSON    applicationEditResponseEnvelopeJSON `json:"-"`
}

// applicationEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ApplicationEditResponseEnvelope]
type applicationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ApplicationEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             applicationEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// applicationEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ApplicationEditResponseEnvelopeErrors]
type applicationEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    applicationEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// applicationEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ApplicationEditResponseEnvelopeErrorsSource]
type applicationEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ApplicationEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             applicationEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// applicationEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ApplicationEditResponseEnvelopeMessages]
type applicationEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ApplicationEditResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    applicationEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// applicationEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ApplicationEditResponseEnvelopeMessagesSource]
type applicationEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ApplicationGetResponseEnvelope struct {
	Errors   []ApplicationGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ApplicationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// The public Containers API returns an application.
	Result ApplicationGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                               `json:"success" api:"required"`
	JSON    applicationGetResponseEnvelopeJSON `json:"-"`
}

// applicationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ApplicationGetResponseEnvelope]
type applicationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ApplicationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             applicationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// applicationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ApplicationGetResponseEnvelopeErrors]
type applicationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    applicationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// applicationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ApplicationGetResponseEnvelopeErrorsSource]
type applicationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ApplicationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             applicationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// applicationGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ApplicationGetResponseEnvelopeMessages]
type applicationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ApplicationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    applicationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// applicationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ApplicationGetResponseEnvelopeMessagesSource]
type applicationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
