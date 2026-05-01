// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// DeviceDeploymentGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceDeploymentGroupService] method instead.
type DeviceDeploymentGroupService struct {
	Options []option.RequestOption
}

// NewDeviceDeploymentGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceDeploymentGroupService(opts ...option.RequestOption) (r *DeviceDeploymentGroupService) {
	r = &DeviceDeploymentGroupService{}
	r.Options = opts
	return
}

// Creates a new deployment group. Policy IDs must be unique across all deployment
// groups. This endpoint is in Beta.
func (r *DeviceDeploymentGroupService) New(ctx context.Context, params DeviceDeploymentGroupNewParams, opts ...option.RequestOption) (res *DeploymentGroup, err error) {
	var env DeviceDeploymentGroupNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/deployment-groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all deployment groups for an account. Use deployment groups to assign
// target WARP client versions to specific devices. This endpoint is in Beta.
func (r *DeviceDeploymentGroupService) List(ctx context.Context, params DeviceDeploymentGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DeploymentGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/deployment-groups", params.AccountID)
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

// Lists all deployment groups for an account. Use deployment groups to assign
// target WARP client versions to specific devices. This endpoint is in Beta.
func (r *DeviceDeploymentGroupService) ListAutoPaging(ctx context.Context, params DeviceDeploymentGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DeploymentGroup] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a deployment group. Associated policies no longer apply and devices stop
// receiving version targets. This endpoint is in Beta.
func (r *DeviceDeploymentGroupService) Delete(ctx context.Context, groupID string, body DeviceDeploymentGroupDeleteParams, opts ...option.RequestOption) (res *DeviceDeploymentGroupDeleteResponse, err error) {
	var env DeviceDeploymentGroupDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/deployment-groups/%s", body.AccountID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a deployment group. Returns 409 if any newly added policy IDs already
// belong to another deployment group. This endpoint is in Beta.
func (r *DeviceDeploymentGroupService) Edit(ctx context.Context, groupID string, params DeviceDeploymentGroupEditParams, opts ...option.RequestOption) (res *DeploymentGroup, err error) {
	var env DeviceDeploymentGroupEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/deployment-groups/%s", params.AccountID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches a single deployment group by its ID. This endpoint is in Beta.
func (r *DeviceDeploymentGroupService) Get(ctx context.Context, groupID string, query DeviceDeploymentGroupGetParams, opts ...option.RequestOption) (res *DeploymentGroup, err error) {
	var env DeviceDeploymentGroupGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/deployment-groups/%s", query.AccountID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DeploymentGroup struct {
	// The ID of the deployment group.
	ID string `json:"id" api:"required"`
	// The RFC3339Nano timestamp when the deployment group was created.
	CreatedAt string `json:"created_at" api:"required"`
	// A user-friendly name for the deployment group.
	Name string `json:"name" api:"required"`
	// The RFC3339Nano timestamp when the deployment group was last updated.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Contains version configurations for different target environments.
	VersionConfig []DeploymentGroupVersionConfig `json:"version_config" api:"required"`
	// Contains a list of policy IDs assigned to this deployment group.
	PolicyIDs []string            `json:"policy_ids" api:"nullable"`
	JSON      deploymentGroupJSON `json:"-"`
}

// deploymentGroupJSON contains the JSON metadata for the struct [DeploymentGroup]
type deploymentGroupJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	UpdatedAt     apijson.Field
	VersionConfig apijson.Field
	PolicyIDs     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentGroupJSON) RawJSON() string {
	return r.raw
}

type DeploymentGroupVersionConfig struct {
	// The target environment for the client version (e.g., windows, macos).
	TargetEnvironment string `json:"target_environment" api:"required,nullable"`
	// The specific client version to deploy.
	Version string                           `json:"version" api:"required"`
	JSON    deploymentGroupVersionConfigJSON `json:"-"`
}

// deploymentGroupVersionConfigJSON contains the JSON metadata for the struct
// [DeploymentGroupVersionConfig]
type deploymentGroupVersionConfigJSON struct {
	TargetEnvironment apijson.Field
	Version           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DeploymentGroupVersionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentGroupVersionConfigJSON) RawJSON() string {
	return r.raw
}

type DeviceDeploymentGroupDeleteResponse struct {
	// The ID of a deleted deployment group.
	ID   string                                  `json:"id"`
	JSON deviceDeploymentGroupDeleteResponseJSON `json:"-"`
}

// deviceDeploymentGroupDeleteResponseJSON contains the JSON metadata for the
// struct [DeviceDeploymentGroupDeleteResponse]
type deviceDeploymentGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceDeploymentGroupNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A user-friendly name for the deployment group.
	Name param.Field[string] `json:"name" api:"required"`
	// Contains at least one version configuration.
	VersionConfig param.Field[[]DeviceDeploymentGroupNewParamsVersionConfig] `json:"version_config" api:"required"`
	// Contains an optional list of policy IDs assigned to a group.
	PolicyIDs param.Field[[]string] `json:"policy_ids"`
}

func (r DeviceDeploymentGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDeploymentGroupNewParamsVersionConfig struct {
	// The target environment for the client version (e.g., windows, macos).
	TargetEnvironment param.Field[string] `json:"target_environment" api:"required"`
	// The specific client version to deploy.
	Version param.Field[string] `json:"version" api:"required"`
}

func (r DeviceDeploymentGroupNewParamsVersionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDeploymentGroupNewResponseEnvelope struct {
	Errors   []DeviceDeploymentGroupNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDeploymentGroupNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DeploymentGroup                                    `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                                         `json:"success" api:"required"`
	JSON    deviceDeploymentGroupNewResponseEnvelopeJSON `json:"-"`
}

// deviceDeploymentGroupNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDeploymentGroupNewResponseEnvelope]
type deviceDeploymentGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code" api:"required"`
	Message string                                             `json:"message" api:"required"`
	JSON    deviceDeploymentGroupNewResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeploymentGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupNewResponseEnvelopeErrors]
type deviceDeploymentGroupNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	JSON    deviceDeploymentGroupNewResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeploymentGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupNewResponseEnvelopeMessages]
type deviceDeploymentGroupNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDeploymentGroupListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The page number to return.
	Page param.Field[int64] `query:"page"`
	// The maximum number of deployment groups to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [DeviceDeploymentGroupListParams]'s query parameters as
// `url.Values`.
func (r DeviceDeploymentGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceDeploymentGroupDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DeviceDeploymentGroupDeleteResponseEnvelope struct {
	Errors   []DeviceDeploymentGroupDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDeploymentGroupDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DeviceDeploymentGroupDeleteResponse                   `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                                            `json:"success" api:"required"`
	JSON    deviceDeploymentGroupDeleteResponseEnvelopeJSON `json:"-"`
}

// deviceDeploymentGroupDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DeviceDeploymentGroupDeleteResponseEnvelope]
type deviceDeploymentGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code" api:"required"`
	Message string                                                `json:"message" api:"required"`
	JSON    deviceDeploymentGroupDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeploymentGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupDeleteResponseEnvelopeErrors]
type deviceDeploymentGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code" api:"required"`
	Message string                                                  `json:"message" api:"required"`
	JSON    deviceDeploymentGroupDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeploymentGroupDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DeviceDeploymentGroupDeleteResponseEnvelopeMessages]
type deviceDeploymentGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDeploymentGroupEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A user-friendly name for the deployment group.
	Name param.Field[string] `json:"name"`
	// Replaces the entire list of policy IDs.
	PolicyIDs param.Field[[]string] `json:"policy_ids"`
	// Replaces the entire version_config array.
	VersionConfig param.Field[[]DeviceDeploymentGroupEditParamsVersionConfig] `json:"version_config"`
}

func (r DeviceDeploymentGroupEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDeploymentGroupEditParamsVersionConfig struct {
	// The target environment for the client version (e.g., windows, macos).
	TargetEnvironment param.Field[string] `json:"target_environment" api:"required"`
	// The specific client version to deploy.
	Version param.Field[string] `json:"version" api:"required"`
}

func (r DeviceDeploymentGroupEditParamsVersionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDeploymentGroupEditResponseEnvelope struct {
	Errors   []DeviceDeploymentGroupEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDeploymentGroupEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DeploymentGroup                                     `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                                          `json:"success" api:"required"`
	JSON    deviceDeploymentGroupEditResponseEnvelopeJSON `json:"-"`
}

// deviceDeploymentGroupEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDeploymentGroupEditResponseEnvelope]
type deviceDeploymentGroupEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	JSON    deviceDeploymentGroupEditResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeploymentGroupEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupEditResponseEnvelopeErrors]
type deviceDeploymentGroupEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code" api:"required"`
	Message string                                                `json:"message" api:"required"`
	JSON    deviceDeploymentGroupEditResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeploymentGroupEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupEditResponseEnvelopeMessages]
type deviceDeploymentGroupEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDeploymentGroupGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DeviceDeploymentGroupGetResponseEnvelope struct {
	Errors   []DeviceDeploymentGroupGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceDeploymentGroupGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DeploymentGroup                                    `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                                         `json:"success" api:"required"`
	JSON    deviceDeploymentGroupGetResponseEnvelopeJSON `json:"-"`
}

// deviceDeploymentGroupGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDeploymentGroupGetResponseEnvelope]
type deviceDeploymentGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code" api:"required"`
	Message string                                             `json:"message" api:"required"`
	JSON    deviceDeploymentGroupGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeploymentGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupGetResponseEnvelopeErrors]
type deviceDeploymentGroupGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeploymentGroupGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	JSON    deviceDeploymentGroupGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeploymentGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DeviceDeploymentGroupGetResponseEnvelopeMessages]
type deviceDeploymentGroupGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeploymentGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeploymentGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
