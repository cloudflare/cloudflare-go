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
func (r *DispatchNamespaceScriptSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *workers.ScriptSetting, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSettingGetParams, opts ...option.RequestOption) (res *workers.ScriptSetting, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSettingGetResponseEnvelope
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
	AccountID param.Field[string]                     `path:"account_id,required"`
	Errors    param.Field[[]shared.ResponseInfoParam] `json:"errors,required"`
	Messages  param.Field[[]shared.ResponseInfoParam] `json:"messages,required"`
	Result    param.Field[workers.ScriptSettingParam] `json:"result,required"`
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

func (r DispatchNamespaceScriptSettingEditParamsSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   workers.ScriptSetting `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseEnvelope]
type dispatchNamespaceScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   workers.ScriptSetting `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseEnvelope]
type dispatchNamespaceScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
