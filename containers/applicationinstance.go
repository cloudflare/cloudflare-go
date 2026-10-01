// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// ApplicationInstanceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationInstanceService] method instead.
type ApplicationInstanceService struct {
	Options []option.RequestOption
}

// NewApplicationInstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationInstanceService(opts ...option.RequestOption) (r *ApplicationInstanceService) {
	r = &ApplicationInstanceService{}
	r.Options = opts
	return
}

// Lists container instances belonging to an application.
func (r *ApplicationInstanceService) List(ctx context.Context, applicationID string, params ApplicationInstanceListParams, opts ...option.RequestOption) (res *pagination.PageTokenPagination[ApplicationInstanceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s/instances-v2", params.AccountID, applicationID)
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

// Lists container instances belonging to an application.
func (r *ApplicationInstanceService) ListAutoPaging(ctx context.Context, applicationID string, params ApplicationInstanceListParams, opts ...option.RequestOption) *pagination.PageTokenPaginationAutoPager[ApplicationInstanceListResponse] {
	return pagination.NewPageTokenPaginationAutoPager(r.List(ctx, applicationID, params, opts...))
}

// Returns a container instance belonging to an application.
func (r *ApplicationInstanceService) Get(ctx context.Context, applicationID string, instanceID string, query ApplicationInstanceGetParams, opts ...option.RequestOption) (res *ApplicationInstanceGetResponse, err error) {
	var env ApplicationInstanceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s/instances/%s", query.AccountID, applicationID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deprecated: use `instances-v2` instead. Lists container instances belonging to
// an application.
//
// Deprecated: deprecated
func (r *ApplicationInstanceService) ListV1(ctx context.Context, applicationID string, params ApplicationInstanceListV1Params, opts ...option.RequestOption) (res *pagination.ContainersInstancesV1Pagination[ApplicationInstanceListV1Response], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s/instances", params.AccountID, applicationID)
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

// Deprecated: use `instances-v2` instead. Lists container instances belonging to
// an application.
//
// Deprecated: deprecated
func (r *ApplicationInstanceService) ListV1AutoPaging(ctx context.Context, applicationID string, params ApplicationInstanceListV1Params, opts ...option.RequestOption) *pagination.ContainersInstancesV1PaginationAutoPager[ApplicationInstanceListV1Response] {
	return pagination.NewContainersInstancesV1PaginationAutoPager(r.ListV1(ctx, applicationID, params, opts...))
}

// The last-reported state of a logical container instance.
type ApplicationInstanceListResponse struct {
	// A container instance ID (64-character hex Durable Object actor ID).
	ID string `json:"id" api:"required"`
	// An Application ID represents an identifier of an application.
	ApplicationID string `json:"application_id" api:"required"`
	// The image for the current container placement, when one is available.
	Image string `json:"image" api:"required"`
	// The latest known status of a container instance.
	Status ApplicationInstanceListResponseStatus `json:"status" api:"required"`
	// The resources allocated to the container instance.
	Configuration ApplicationInstanceListResponseConfiguration `json:"configuration"`
	// The location of the instance's current container placement.
	Location ApplicationInstanceListResponseLocation `json:"location"`
	// The customer-provided instance name, when available. Its UTF-8 encoding uses at
	// most 1,024 bytes.
	Name string `json:"name"`
	// The time at which the current container placement started, when one exists.
	StartedAt string                              `json:"started_at"`
	JSON      applicationInstanceListResponseJSON `json:"-"`
}

// applicationInstanceListResponseJSON contains the JSON metadata for the struct
// [ApplicationInstanceListResponse]
type applicationInstanceListResponseJSON struct {
	ID            apijson.Field
	ApplicationID apijson.Field
	Image         apijson.Field
	Status        apijson.Field
	Configuration apijson.Field
	Location      apijson.Field
	Name          apijson.Field
	StartedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ApplicationInstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListResponseJSON) RawJSON() string {
	return r.raw
}

// The latest known status of a container instance.
type ApplicationInstanceListResponseStatus struct {
	// The current lifecycle state of a container instance.
	State ApplicationInstanceListResponseStatusState `json:"state" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// The process exit code, when the runtime reports one.
	ExitCode int64                                     `json:"exit_code"`
	JSON     applicationInstanceListResponseStatusJSON `json:"-"`
}

// applicationInstanceListResponseStatusJSON contains the JSON metadata for the
// struct [ApplicationInstanceListResponseStatus]
type applicationInstanceListResponseStatusJSON struct {
	State       apijson.Field
	UpdatedAt   apijson.Field
	ExitCode    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceListResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListResponseStatusJSON) RawJSON() string {
	return r.raw
}

// The current lifecycle state of a container instance.
type ApplicationInstanceListResponseStatusState string

const (
	ApplicationInstanceListResponseStatusStateProvisioning ApplicationInstanceListResponseStatusState = "provisioning"
	ApplicationInstanceListResponseStatusStateRunning      ApplicationInstanceListResponseStatusState = "running"
	ApplicationInstanceListResponseStatusStateFailed       ApplicationInstanceListResponseStatusState = "failed"
	ApplicationInstanceListResponseStatusStateStopping     ApplicationInstanceListResponseStatusState = "stopping"
	ApplicationInstanceListResponseStatusStateStopped      ApplicationInstanceListResponseStatusState = "stopped"
	ApplicationInstanceListResponseStatusStateUnhealthy    ApplicationInstanceListResponseStatusState = "unhealthy"
	ApplicationInstanceListResponseStatusStateInactive     ApplicationInstanceListResponseStatusState = "inactive"
	ApplicationInstanceListResponseStatusStateUnknown      ApplicationInstanceListResponseStatusState = "unknown"
)

func (r ApplicationInstanceListResponseStatusState) IsKnown() bool {
	switch r {
	case ApplicationInstanceListResponseStatusStateProvisioning, ApplicationInstanceListResponseStatusStateRunning, ApplicationInstanceListResponseStatusStateFailed, ApplicationInstanceListResponseStatusStateStopping, ApplicationInstanceListResponseStatusStateStopped, ApplicationInstanceListResponseStatusStateUnhealthy, ApplicationInstanceListResponseStatusStateInactive, ApplicationInstanceListResponseStatusStateUnknown:
		return true
	}
	return false
}

// The resources allocated to the container instance.
type ApplicationInstanceListResponseConfiguration struct {
	// Disk allocated to the container instance, in decimal MB.
	Disk int64 `json:"disk" api:"required"`
	// Memory allocated to the container instance, in MiB.
	Memory int64 `json:"memory" api:"required"`
	// Number of virtual CPUs allocated to the container instance.
	Vcpu float64                                          `json:"vcpu" api:"required"`
	JSON applicationInstanceListResponseConfigurationJSON `json:"-"`
}

// applicationInstanceListResponseConfigurationJSON contains the JSON metadata for
// the struct [ApplicationInstanceListResponseConfiguration]
type applicationInstanceListResponseConfigurationJSON struct {
	Disk        apijson.Field
	Memory      apijson.Field
	Vcpu        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The location of the instance's current container placement.
type ApplicationInstanceListResponseLocation struct {
	// Unique location code used to identify locations on a logical level.
	Name string `json:"name" api:"required"`
	// Represents a group of datacenters. Choose one of "AFR", "APAC", "EEUR", "ENAM",
	// "WNAM", "ME", "OC", "SAM", or "WEUR".
	Region string                                      `json:"region" api:"required"`
	JSON   applicationInstanceListResponseLocationJSON `json:"-"`
}

// applicationInstanceListResponseLocationJSON contains the JSON metadata for the
// struct [ApplicationInstanceListResponseLocation]
type applicationInstanceListResponseLocationJSON struct {
	Name        apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceListResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListResponseLocationJSON) RawJSON() string {
	return r.raw
}

// The last-reported state of a logical container instance.
type ApplicationInstanceGetResponse struct {
	// A container instance ID (64-character hex Durable Object actor ID).
	ID string `json:"id" api:"required"`
	// An Application ID represents an identifier of an application.
	ApplicationID string `json:"application_id" api:"required"`
	// The image for the current container placement, when one is available.
	Image string `json:"image" api:"required"`
	// The latest known status of a container instance.
	Status ApplicationInstanceGetResponseStatus `json:"status" api:"required"`
	// The resources allocated to the container instance.
	Configuration ApplicationInstanceGetResponseConfiguration `json:"configuration"`
	// The location of the instance's current container placement.
	Location ApplicationInstanceGetResponseLocation `json:"location"`
	// The customer-provided instance name, when available. Its UTF-8 encoding uses at
	// most 1,024 bytes.
	Name string `json:"name"`
	// The time at which the current container placement started, when one exists.
	StartedAt string                             `json:"started_at"`
	JSON      applicationInstanceGetResponseJSON `json:"-"`
}

// applicationInstanceGetResponseJSON contains the JSON metadata for the struct
// [ApplicationInstanceGetResponse]
type applicationInstanceGetResponseJSON struct {
	ID            apijson.Field
	ApplicationID apijson.Field
	Image         apijson.Field
	Status        apijson.Field
	Configuration apijson.Field
	Location      apijson.Field
	Name          apijson.Field
	StartedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseJSON) RawJSON() string {
	return r.raw
}

// The latest known status of a container instance.
type ApplicationInstanceGetResponseStatus struct {
	// The current lifecycle state of a container instance.
	State ApplicationInstanceGetResponseStatusState `json:"state" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// The process exit code, when the runtime reports one.
	ExitCode int64                                    `json:"exit_code"`
	JSON     applicationInstanceGetResponseStatusJSON `json:"-"`
}

// applicationInstanceGetResponseStatusJSON contains the JSON metadata for the
// struct [ApplicationInstanceGetResponseStatus]
type applicationInstanceGetResponseStatusJSON struct {
	State       apijson.Field
	UpdatedAt   apijson.Field
	ExitCode    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseStatusJSON) RawJSON() string {
	return r.raw
}

// The current lifecycle state of a container instance.
type ApplicationInstanceGetResponseStatusState string

const (
	ApplicationInstanceGetResponseStatusStateProvisioning ApplicationInstanceGetResponseStatusState = "provisioning"
	ApplicationInstanceGetResponseStatusStateRunning      ApplicationInstanceGetResponseStatusState = "running"
	ApplicationInstanceGetResponseStatusStateFailed       ApplicationInstanceGetResponseStatusState = "failed"
	ApplicationInstanceGetResponseStatusStateStopping     ApplicationInstanceGetResponseStatusState = "stopping"
	ApplicationInstanceGetResponseStatusStateStopped      ApplicationInstanceGetResponseStatusState = "stopped"
	ApplicationInstanceGetResponseStatusStateUnhealthy    ApplicationInstanceGetResponseStatusState = "unhealthy"
	ApplicationInstanceGetResponseStatusStateInactive     ApplicationInstanceGetResponseStatusState = "inactive"
	ApplicationInstanceGetResponseStatusStateUnknown      ApplicationInstanceGetResponseStatusState = "unknown"
)

func (r ApplicationInstanceGetResponseStatusState) IsKnown() bool {
	switch r {
	case ApplicationInstanceGetResponseStatusStateProvisioning, ApplicationInstanceGetResponseStatusStateRunning, ApplicationInstanceGetResponseStatusStateFailed, ApplicationInstanceGetResponseStatusStateStopping, ApplicationInstanceGetResponseStatusStateStopped, ApplicationInstanceGetResponseStatusStateUnhealthy, ApplicationInstanceGetResponseStatusStateInactive, ApplicationInstanceGetResponseStatusStateUnknown:
		return true
	}
	return false
}

// The resources allocated to the container instance.
type ApplicationInstanceGetResponseConfiguration struct {
	// Disk allocated to the container instance, in decimal MB.
	Disk int64 `json:"disk" api:"required"`
	// Memory allocated to the container instance, in MiB.
	Memory int64 `json:"memory" api:"required"`
	// Number of virtual CPUs allocated to the container instance.
	Vcpu float64                                         `json:"vcpu" api:"required"`
	JSON applicationInstanceGetResponseConfigurationJSON `json:"-"`
}

// applicationInstanceGetResponseConfigurationJSON contains the JSON metadata for
// the struct [ApplicationInstanceGetResponseConfiguration]
type applicationInstanceGetResponseConfigurationJSON struct {
	Disk        apijson.Field
	Memory      apijson.Field
	Vcpu        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The location of the instance's current container placement.
type ApplicationInstanceGetResponseLocation struct {
	// Unique location code used to identify locations on a logical level.
	Name string `json:"name" api:"required"`
	// Represents a group of datacenters. Choose one of "AFR", "APAC", "EEUR", "ENAM",
	// "WNAM", "ME", "OC", "SAM", or "WEUR".
	Region string                                     `json:"region" api:"required"`
	JSON   applicationInstanceGetResponseLocationJSON `json:"-"`
}

// applicationInstanceGetResponseLocationJSON contains the JSON metadata for the
// struct [ApplicationInstanceGetResponseLocation]
type applicationInstanceGetResponseLocationJSON struct {
	Name        apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseLocationJSON) RawJSON() string {
	return r.raw
}

// The last-reported state of a logical container instance.
type ApplicationInstanceListV1Response struct {
	// A container instance ID (64-character hex Durable Object actor ID).
	ID string `json:"id" api:"required"`
	// An Application ID represents an identifier of an application.
	ApplicationID string `json:"application_id" api:"required"`
	// The image for the current container placement, when one is available.
	Image string `json:"image" api:"required"`
	// The latest known status of a container instance.
	Status ApplicationInstanceListV1ResponseStatus `json:"status" api:"required"`
	// The resources allocated to the container instance.
	Configuration ApplicationInstanceListV1ResponseConfiguration `json:"configuration"`
	// The location of the instance's current container placement.
	Location ApplicationInstanceListV1ResponseLocation `json:"location"`
	// The customer-provided instance name, when available. Its UTF-8 encoding uses at
	// most 1,024 bytes.
	Name string `json:"name"`
	// The time at which the current container placement started, when one exists.
	StartedAt string                                `json:"started_at"`
	JSON      applicationInstanceListV1ResponseJSON `json:"-"`
}

// applicationInstanceListV1ResponseJSON contains the JSON metadata for the struct
// [ApplicationInstanceListV1Response]
type applicationInstanceListV1ResponseJSON struct {
	ID            apijson.Field
	ApplicationID apijson.Field
	Image         apijson.Field
	Status        apijson.Field
	Configuration apijson.Field
	Location      apijson.Field
	Name          apijson.Field
	StartedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ApplicationInstanceListV1Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListV1ResponseJSON) RawJSON() string {
	return r.raw
}

// The latest known status of a container instance.
type ApplicationInstanceListV1ResponseStatus struct {
	// The current lifecycle state of a container instance.
	State ApplicationInstanceListV1ResponseStatusState `json:"state" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	UpdatedAt string `json:"updated_at" api:"required"`
	// The process exit code, when the runtime reports one.
	ExitCode int64                                       `json:"exit_code"`
	JSON     applicationInstanceListV1ResponseStatusJSON `json:"-"`
}

// applicationInstanceListV1ResponseStatusJSON contains the JSON metadata for the
// struct [ApplicationInstanceListV1ResponseStatus]
type applicationInstanceListV1ResponseStatusJSON struct {
	State       apijson.Field
	UpdatedAt   apijson.Field
	ExitCode    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceListV1ResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListV1ResponseStatusJSON) RawJSON() string {
	return r.raw
}

// The current lifecycle state of a container instance.
type ApplicationInstanceListV1ResponseStatusState string

const (
	ApplicationInstanceListV1ResponseStatusStateProvisioning ApplicationInstanceListV1ResponseStatusState = "provisioning"
	ApplicationInstanceListV1ResponseStatusStateRunning      ApplicationInstanceListV1ResponseStatusState = "running"
	ApplicationInstanceListV1ResponseStatusStateFailed       ApplicationInstanceListV1ResponseStatusState = "failed"
	ApplicationInstanceListV1ResponseStatusStateStopping     ApplicationInstanceListV1ResponseStatusState = "stopping"
	ApplicationInstanceListV1ResponseStatusStateStopped      ApplicationInstanceListV1ResponseStatusState = "stopped"
	ApplicationInstanceListV1ResponseStatusStateUnhealthy    ApplicationInstanceListV1ResponseStatusState = "unhealthy"
	ApplicationInstanceListV1ResponseStatusStateInactive     ApplicationInstanceListV1ResponseStatusState = "inactive"
	ApplicationInstanceListV1ResponseStatusStateUnknown      ApplicationInstanceListV1ResponseStatusState = "unknown"
)

func (r ApplicationInstanceListV1ResponseStatusState) IsKnown() bool {
	switch r {
	case ApplicationInstanceListV1ResponseStatusStateProvisioning, ApplicationInstanceListV1ResponseStatusStateRunning, ApplicationInstanceListV1ResponseStatusStateFailed, ApplicationInstanceListV1ResponseStatusStateStopping, ApplicationInstanceListV1ResponseStatusStateStopped, ApplicationInstanceListV1ResponseStatusStateUnhealthy, ApplicationInstanceListV1ResponseStatusStateInactive, ApplicationInstanceListV1ResponseStatusStateUnknown:
		return true
	}
	return false
}

// The resources allocated to the container instance.
type ApplicationInstanceListV1ResponseConfiguration struct {
	// Disk allocated to the container instance, in decimal MB.
	Disk int64 `json:"disk" api:"required"`
	// Memory allocated to the container instance, in MiB.
	Memory int64 `json:"memory" api:"required"`
	// Number of virtual CPUs allocated to the container instance.
	Vcpu float64                                            `json:"vcpu" api:"required"`
	JSON applicationInstanceListV1ResponseConfigurationJSON `json:"-"`
}

// applicationInstanceListV1ResponseConfigurationJSON contains the JSON metadata
// for the struct [ApplicationInstanceListV1ResponseConfiguration]
type applicationInstanceListV1ResponseConfigurationJSON struct {
	Disk        apijson.Field
	Memory      apijson.Field
	Vcpu        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceListV1ResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListV1ResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The location of the instance's current container placement.
type ApplicationInstanceListV1ResponseLocation struct {
	// Unique location code used to identify locations on a logical level.
	Name string `json:"name" api:"required"`
	// Represents a group of datacenters. Choose one of "AFR", "APAC", "EEUR", "ENAM",
	// "WNAM", "ME", "OC", "SAM", or "WEUR".
	Region string                                        `json:"region" api:"required"`
	JSON   applicationInstanceListV1ResponseLocationJSON `json:"-"`
}

// applicationInstanceListV1ResponseLocationJSON contains the JSON metadata for the
// struct [ApplicationInstanceListV1ResponseLocation]
type applicationInstanceListV1ResponseLocationJSON struct {
	Name        apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceListV1ResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceListV1ResponseLocationJSON) RawJSON() string {
	return r.raw
}

type ApplicationInstanceListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter instances by a case-sensitive name prefix, falling back to the actor ID
	// when no name is known. Keep the same prefix when using a page token.
	NamePrefix param.Field[string] `query:"name_prefix"`
	// Opaque token from a previous response to retrieve the next page.
	PageToken param.Field[string] `query:"page_token"`
	// Maximum number of instances to return per page. Defaults to 100.
	PerPage param.Field[int64] `query:"per_page"`
	// Filters instances by lifecycle state. `active` includes provisioning, running,
	// and stopping instances; `not-active` includes stopped and failed instances. When
	// omitted, all instances are returned.
	State param.Field[ApplicationInstanceListParamsState] `query:"state"`
}

// URLQuery serializes [ApplicationInstanceListParams]'s query parameters as
// `url.Values`.
func (r ApplicationInstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters instances by lifecycle state. `active` includes provisioning, running,
// and stopping instances; `not-active` includes stopped and failed instances. When
// omitted, all instances are returned.
type ApplicationInstanceListParamsState string

const (
	ApplicationInstanceListParamsStateActive    ApplicationInstanceListParamsState = "active"
	ApplicationInstanceListParamsStateNotActive ApplicationInstanceListParamsState = "not-active"
)

func (r ApplicationInstanceListParamsState) IsKnown() bool {
	switch r {
	case ApplicationInstanceListParamsStateActive, ApplicationInstanceListParamsStateNotActive:
		return true
	}
	return false
}

type ApplicationInstanceGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ApplicationInstanceGetResponseEnvelope struct {
	Errors   []ApplicationInstanceGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ApplicationInstanceGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// The last-reported state of a logical container instance.
	Result ApplicationInstanceGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                       `json:"success" api:"required"`
	JSON    applicationInstanceGetResponseEnvelopeJSON `json:"-"`
}

// applicationInstanceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ApplicationInstanceGetResponseEnvelope]
type applicationInstanceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ApplicationInstanceGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ApplicationInstanceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             applicationInstanceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// applicationInstanceGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ApplicationInstanceGetResponseEnvelopeErrors]
type applicationInstanceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ApplicationInstanceGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    applicationInstanceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// applicationInstanceGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ApplicationInstanceGetResponseEnvelopeErrorsSource]
type applicationInstanceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationInstanceGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ApplicationInstanceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             applicationInstanceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// applicationInstanceGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ApplicationInstanceGetResponseEnvelopeMessages]
type applicationInstanceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ApplicationInstanceGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    applicationInstanceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// applicationInstanceGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ApplicationInstanceGetResponseEnvelopeMessagesSource]
type applicationInstanceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationInstanceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationInstanceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationInstanceListV1Params struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter instances by a case-sensitive name prefix, falling back to the actor ID
	// when no name is known. Keep the same prefix when using a page token.
	NamePrefix param.Field[string] `query:"name_prefix"`
	// Opaque token from a previous response to retrieve the next page.
	PageToken param.Field[string] `query:"page_token"`
	// Maximum number of instances to return per page. Defaults to all.
	PerPage param.Field[int64] `query:"per_page"`
	// Filters instances by lifecycle state. `active` includes provisioning, running,
	// and stopping instances; `not-active` includes stopped and failed instances. When
	// omitted, all instances are returned.
	State param.Field[ApplicationInstanceListV1ParamsState] `query:"state"`
}

// URLQuery serializes [ApplicationInstanceListV1Params]'s query parameters as
// `url.Values`.
func (r ApplicationInstanceListV1Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters instances by lifecycle state. `active` includes provisioning, running,
// and stopping instances; `not-active` includes stopped and failed instances. When
// omitted, all instances are returned.
type ApplicationInstanceListV1ParamsState string

const (
	ApplicationInstanceListV1ParamsStateActive    ApplicationInstanceListV1ParamsState = "active"
	ApplicationInstanceListV1ParamsStateNotActive ApplicationInstanceListV1ParamsState = "not-active"
)

func (r ApplicationInstanceListV1ParamsState) IsKnown() bool {
	switch r {
	case ApplicationInstanceListV1ParamsStateActive, ApplicationInstanceListV1ParamsStateNotActive:
		return true
	}
	return false
}
