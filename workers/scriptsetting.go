// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ScriptSettingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptSettingService] method instead.
type ScriptSettingService struct {
	Options []option.RequestOption
}

// NewScriptSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSettingService(opts ...option.RequestOption) (r *ScriptSettingService) {
	r = &ScriptSettingService{}
	r.Options = opts
	return
}

// Patch script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Including but not limited to Logpush and Tail Consumers.
func (r *ScriptSettingService) Edit(ctx context.Context, scriptName string, params ScriptSettingEditParams, opts ...option.RequestOption) (res *ScriptSetting, err error) {
	var env ScriptSettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Includes Logpush and Tail Consumers.
func (r *ScriptSettingService) Get(ctx context.Context, scriptName string, query ScriptSettingGetParams, opts ...option.RequestOption) (res *ScriptSetting, err error) {
	var env ScriptSettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptSettingEditParams struct {
	// Identifier
	AccountID     param.Field[string] `path:"account_id,required"`
	ScriptSetting ScriptSettingParam  `json:"script_setting,required"`
}

func (r ScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ScriptSetting)
}

type ScriptSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptSetting                            `json:"result"`
	JSON    scriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponseEnvelope]
type scriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingEditResponseEnvelopeSuccess bool

const (
	ScriptSettingEditResponseEnvelopeSuccessTrue ScriptSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptSetting                           `json:"result"`
	JSON    scriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponseEnvelope]
type scriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingGetResponseEnvelopeSuccess bool

const (
	ScriptSettingGetResponseEnvelopeSuccessTrue ScriptSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
