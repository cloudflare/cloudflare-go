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

// ScriptUsageModelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptUsageModelService] method
// instead.
type ScriptUsageModelService struct {
	Options []option.RequestOption
}

// NewScriptUsageModelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScriptUsageModelService(opts ...option.RequestOption) (r *ScriptUsageModelService) {
	r = &ScriptUsageModelService{}
	r.Options = opts
	return
}

// Updates the Usage Model for a given Worker. Requires a Workers Paid
// subscription.
func (r *ScriptUsageModelService) Update(ctx context.Context, scriptName string, params ScriptUsageModelUpdateParams, opts ...option.RequestOption) (res *ScriptUsageModelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptUsageModelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Usage Model for a given Worker.
func (r *ScriptUsageModelService) Get(ctx context.Context, scriptName string, query ScriptUsageModelGetParams, opts ...option.RequestOption) (res *ScriptUsageModelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptUsageModelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptUsageModelUpdateResponse struct {
	UsageModel interface{}                        `json:"usage_model"`
	JSON       scriptUsageModelUpdateResponseJSON `json:"-"`
}

// scriptUsageModelUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptUsageModelUpdateResponse]
type scriptUsageModelUpdateResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUsageModelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUsageModelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptUsageModelGetResponse struct {
	UsageModel interface{}                     `json:"usage_model"`
	JSON       scriptUsageModelGetResponseJSON `json:"-"`
}

// scriptUsageModelGetResponseJSON contains the JSON metadata for the struct
// [ScriptUsageModelGetResponse]
type scriptUsageModelGetResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUsageModelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUsageModelGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptUsageModelUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      param.Field[string] `json:"body,required"`
}

func (r ScriptUsageModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptUsageModelUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ScriptUsageModelUpdateResponse                            `json:"result,required"`
	// Whether the API call was successful
	Success ScriptUsageModelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptUsageModelUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptUsageModelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptUsageModelUpdateResponseEnvelope]
type scriptUsageModelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUsageModelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUsageModelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptUsageModelUpdateResponseEnvelopeSuccess bool

const (
	ScriptUsageModelUpdateResponseEnvelopeSuccessTrue ScriptUsageModelUpdateResponseEnvelopeSuccess = true
)

func (r ScriptUsageModelUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptUsageModelUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptUsageModelGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptUsageModelGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ScriptUsageModelGetResponse                               `json:"result,required"`
	// Whether the API call was successful
	Success ScriptUsageModelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptUsageModelGetResponseEnvelopeJSON    `json:"-"`
}

// scriptUsageModelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptUsageModelGetResponseEnvelope]
type scriptUsageModelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUsageModelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUsageModelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptUsageModelGetResponseEnvelopeSuccess bool

const (
	ScriptUsageModelGetResponseEnvelopeSuccessTrue ScriptUsageModelGetResponseEnvelopeSuccess = true
)

func (r ScriptUsageModelGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptUsageModelGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
