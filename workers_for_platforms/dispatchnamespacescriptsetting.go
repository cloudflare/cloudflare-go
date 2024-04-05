// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
  "context"
  "fmt"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/internal/shared"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/workers"
)

// DispatchNamespaceScriptSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptSettingService] method instead.
type DispatchNamespaceScriptSettingService struct {
Options []option.RequestOption
}

// NewDispatchNamespaceScriptSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSettingService) {
  r = &DispatchNamespaceScriptSettingService{}
  r.Options = opts
  return
}

// Patch script metadata, such as bindings
func (r *DispatchNamespaceScriptSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *workers.SettingsItem, err error) {
  opts = append(r.Options[:], opts...)
  var env workers.Setting
  path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSettingGetParams, opts ...option.RequestOption) (res *workers.SettingsItem, err error) {
  opts = append(r.Options[:], opts...)
  var env workers.Setting
  path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type DispatchNamespaceScriptSettingEditParams struct {
// Identifier
AccountID param.Field[string] `path:"account_id,required"`
Errors param.Field[[]shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param] `json:"errors,required"`
Messages param.Field[[]shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param] `json:"messages,required"`
Result param.Field[workers.SettingsItemParam] `json:"result,required"`
// Whether the API call was successful
Success param.Field[DispatchNamespaceScriptSettingEditParamsSuccess] `json:"success,required"`
}

func (r DispatchNamespaceScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingEditParamsSuccess bool

const (
  DispatchNamespaceScriptSettingEditParamsSuccessTrue DispatchNamespaceScriptSettingEditParamsSuccess = true
)

func (r DispatchNamespaceScriptSettingEditParamsSuccess) IsKnown() (bool) {
  switch r {
  case DispatchNamespaceScriptSettingEditParamsSuccessTrue:
      return true
  }
  return false
}

type Setting struct {
Errors []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
Result workers.SettingsItem `json:"result,required"`
// Whether the API call was successful
Success workers.SettingSuccess `json:"success,required"`
JSON settingJSON `json:"-"`
}

// settingJSON contains the JSON metadata for the struct [workers.Setting]
type settingJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *workers.Setting) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r settingJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type workers.SettingSuccess bool

const (
  workers.SettingSuccessTrue workers.SettingSuccess = true
)

func (r workers.SettingSuccess) IsKnown() (bool) {
  switch r {
  case workers.SettingSuccessTrue:
      return true
  }
  return false
}

type DispatchNamespaceScriptSettingGetParams struct {
// Identifier
AccountID param.Field[string] `path:"account_id,required"`
}
