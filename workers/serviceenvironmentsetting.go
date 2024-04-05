// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ServiceEnvironmentSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewServiceEnvironmentSettingService] method instead.
type ServiceEnvironmentSettingService struct {
	Options []option.RequestOption
}

// NewServiceEnvironmentSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewServiceEnvironmentSettingService(opts ...option.RequestOption) (r *ServiceEnvironmentSettingService) {
	r = &ServiceEnvironmentSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *ServiceEnvironmentSettingService) Edit(ctx context.Context, serviceName string, environmentName string, params ServiceEnvironmentSettingEditParams, opts ...option.RequestOption) (res *SettingsItem, err error) {
	opts = append(r.Options[:], opts...)
	var env Setting
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", params.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a worker with an environment
func (r *ServiceEnvironmentSettingService) Get(ctx context.Context, serviceName string, environmentName string, query ServiceEnvironmentSettingGetParams, opts ...option.RequestOption) (res *SettingsItem, err error) {
	opts = append(r.Options[:], opts...)
	var env Setting
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", query.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceEnvironmentSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                                         `path:"account_id,required"`
	Errors    param.Field[[]shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param] `json:"errors,required"`
	Messages  param.Field[[]shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param] `json:"messages,required"`
	Result    param.Field[SettingsItemParam]                                              `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[ServiceEnvironmentSettingEditParamsSuccess] `json:"success,required"`
}

func (r ServiceEnvironmentSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type ServiceEnvironmentSettingEditParamsSuccess bool

const (
	ServiceEnvironmentSettingEditParamsSuccessTrue ServiceEnvironmentSettingEditParamsSuccess = true
)

func (r ServiceEnvironmentSettingEditParamsSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsSuccessTrue:
		return true
	}
	return false
}

type Setting struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SettingsItem                                              `json:"result,required"`
	// Whether the API call was successful
	Success SettingSuccess `json:"success,required"`
	JSON    settingJSON    `json:"-"`
}

// settingJSON contains the JSON metadata for the struct [Setting]
type settingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Setting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingSuccess bool

const (
	SettingSuccessTrue SettingSuccess = true
)

func (r SettingSuccess) IsKnown() bool {
	switch r {
	case SettingSuccessTrue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
