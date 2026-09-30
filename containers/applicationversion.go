// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// ApplicationVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationVersionService] method instead.
type ApplicationVersionService struct {
	Options []option.RequestOption
}

// NewApplicationVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationVersionService(opts ...option.RequestOption) (r *ApplicationVersionService) {
	r = &ApplicationVersionService{}
	r.Options = opts
	return
}

// Returns all versions for a scheduler-backed application with
// `scheduling_policy: "default"`. Versions and rollouts do not apply to
// applications with `scheduling_policy: "durable_object"`.
func (r *ApplicationVersionService) List(ctx context.Context, applicationID string, query ApplicationVersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ApplicationVersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s/versions", query.AccountID, applicationID)
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

// Returns all versions for a scheduler-backed application with
// `scheduling_policy: "default"`. Versions and rollouts do not apply to
// applications with `scheduling_policy: "durable_object"`.
func (r *ApplicationVersionService) ListAutoPaging(ctx context.Context, applicationID string, query ApplicationVersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ApplicationVersionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, applicationID, query, opts...))
}

// An application with the configuration of its version.
type ApplicationVersionListResponse struct {
	// User-specified container configuration changes.
	Configuration ApplicationVersionListResponseConfiguration `json:"configuration" api:"required"`
	Percentage    int64                                       `json:"percentage" api:"required"`
	Version       int64                                       `json:"version" api:"required"`
	JSON          applicationVersionListResponseJSON          `json:"-"`
}

// applicationVersionListResponseJSON contains the JSON metadata for the struct
// [ApplicationVersionListResponse]
type applicationVersionListResponseJSON struct {
	Configuration apijson.Field
	Percentage    apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ApplicationVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// User-specified container configuration changes.
type ApplicationVersionListResponseConfiguration struct {
	AuthorizedKeys []ApplicationVersionListResponseConfigurationAuthorizedKey `json:"authorized_keys"`
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
	EnvironmentVariables []ApplicationVersionListResponseConfigurationEnvironmentVariable `json:"environment_variables"`
	// Image url.
	Image string `json:"image"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationVersionListResponseConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationVersionListResponseConfigurationObservability `json:"observability"`
	JSON          applicationVersionListResponseConfigurationJSON          `json:"-"`
}

// applicationVersionListResponseConfigurationJSON contains the JSON metadata for
// the struct [ApplicationVersionListResponseConfiguration]
type applicationVersionListResponseConfigurationJSON struct {
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	Image                apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationVersionListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationVersionListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationVersionListResponseConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                       `json:"name"`
	JSON applicationVersionListResponseConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationVersionListResponseConfigurationAuthorizedKeyJSON contains the JSON
// metadata for the struct
// [ApplicationVersionListResponseConfigurationAuthorizedKey]
type applicationVersionListResponseConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationVersionListResponseConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationVersionListResponseConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationVersionListResponseConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                             `json:"value" api:"required"`
	JSON  applicationVersionListResponseConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationVersionListResponseConfigurationEnvironmentVariableJSON contains the
// JSON metadata for the struct
// [ApplicationVersionListResponseConfigurationEnvironmentVariable]
type applicationVersionListResponseConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationVersionListResponseConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationVersionListResponseConfigurationEnvironmentVariableJSON) RawJSON() string {
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
type ApplicationVersionListResponseConfigurationInstanceType string

const (
	ApplicationVersionListResponseConfigurationInstanceTypeLite      ApplicationVersionListResponseConfigurationInstanceType = "lite"
	ApplicationVersionListResponseConfigurationInstanceTypeBasic     ApplicationVersionListResponseConfigurationInstanceType = "basic"
	ApplicationVersionListResponseConfigurationInstanceTypeStandard1 ApplicationVersionListResponseConfigurationInstanceType = "standard-1"
	ApplicationVersionListResponseConfigurationInstanceTypeStandard2 ApplicationVersionListResponseConfigurationInstanceType = "standard-2"
	ApplicationVersionListResponseConfigurationInstanceTypeStandard3 ApplicationVersionListResponseConfigurationInstanceType = "standard-3"
	ApplicationVersionListResponseConfigurationInstanceTypeStandard4 ApplicationVersionListResponseConfigurationInstanceType = "standard-4"
)

func (r ApplicationVersionListResponseConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationVersionListResponseConfigurationInstanceTypeLite, ApplicationVersionListResponseConfigurationInstanceTypeBasic, ApplicationVersionListResponseConfigurationInstanceTypeStandard1, ApplicationVersionListResponseConfigurationInstanceTypeStandard2, ApplicationVersionListResponseConfigurationInstanceTypeStandard3, ApplicationVersionListResponseConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationVersionListResponseConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationVersionListResponseConfigurationObservabilityLogs `json:"logs"`
	JSON applicationVersionListResponseConfigurationObservabilityJSON `json:"-"`
}

// applicationVersionListResponseConfigurationObservabilityJSON contains the JSON
// metadata for the struct
// [ApplicationVersionListResponseConfigurationObservability]
type applicationVersionListResponseConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationVersionListResponseConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationVersionListResponseConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationVersionListResponseConfigurationObservabilityLogs struct {
	Enabled bool                                                             `json:"enabled"`
	JSON    applicationVersionListResponseConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationVersionListResponseConfigurationObservabilityLogsJSON contains the
// JSON metadata for the struct
// [ApplicationVersionListResponseConfigurationObservabilityLogs]
type applicationVersionListResponseConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationVersionListResponseConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationVersionListResponseConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

type ApplicationVersionListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
