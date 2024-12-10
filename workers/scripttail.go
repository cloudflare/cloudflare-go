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

// ScriptTailService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptTailService] method instead.
type ScriptTailService struct {
	Options []option.RequestOption
}

// NewScriptTailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptTailService(opts ...option.RequestOption) (r *ScriptTailService) {
	r = &ScriptTailService{}
	r.Options = opts
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *ScriptTailService) New(ctx context.Context, scriptName string, params ScriptTailNewParams, opts ...option.RequestOption) (res *ScriptTailNewResponse, err error) {
	var env ScriptTailNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a tail from a Worker.
func (r *ScriptTailService) Delete(ctx context.Context, scriptName string, id string, body ScriptTailDeleteParams, opts ...option.RequestOption) (res *ScriptTailDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", body.AccountID, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get list of tails currently deployed on a Worker.
func (r *ScriptTailService) Get(ctx context.Context, scriptName string, query ScriptTailGetParams, opts ...option.RequestOption) (res *ScriptTailGetResponse, err error) {
	var env ScriptTailGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A reference to a script that will consume logs from the attached Worker.
type ConsumerScript struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string             `json:"namespace"`
	JSON      consumerScriptJSON `json:"-"`
}

// consumerScriptJSON contains the JSON metadata for the struct [ConsumerScript]
type consumerScriptJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerScriptJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type ConsumerScriptParam struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ConsumerScriptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptTailNewResponse struct {
	ID        string                    `json:"id"`
	ExpiresAt string                    `json:"expires_at"`
	URL       string                    `json:"url"`
	JSON      scriptTailNewResponseJSON `json:"-"`
}

// scriptTailNewResponseJSON contains the JSON metadata for the struct
// [ScriptTailNewResponse]
type scriptTailNewResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptTailDeleteResponseSuccess `json:"success,required"`
	JSON    scriptTailDeleteResponseJSON    `json:"-"`
}

// scriptTailDeleteResponseJSON contains the JSON metadata for the struct
// [ScriptTailDeleteResponse]
type scriptTailDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptTailDeleteResponseSuccess bool

const (
	ScriptTailDeleteResponseSuccessTrue ScriptTailDeleteResponseSuccess = true
)

func (r ScriptTailDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ScriptTailDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ScriptTailGetResponse struct {
	ID        string                    `json:"id"`
	ExpiresAt string                    `json:"expires_at"`
	URL       string                    `json:"url"`
	JSON      scriptTailGetResponseJSON `json:"-"`
}

// scriptTailGetResponseJSON contains the JSON metadata for the struct
// [ScriptTailGetResponse]
type scriptTailGetResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ScriptTailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptTailNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptTailNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptTailNewResponse                `json:"result"`
	JSON    scriptTailNewResponseEnvelopeJSON    `json:"-"`
}

// scriptTailNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailNewResponseEnvelope]
type scriptTailNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptTailNewResponseEnvelopeSuccess bool

const (
	ScriptTailNewResponseEnvelopeSuccessTrue ScriptTailNewResponseEnvelopeSuccess = true
)

func (r ScriptTailNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptTailDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptTailGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptTailGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptTailGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptTailGetResponse                `json:"result"`
	JSON    scriptTailGetResponseEnvelopeJSON    `json:"-"`
}

// scriptTailGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailGetResponseEnvelope]
type scriptTailGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptTailGetResponseEnvelopeSuccess bool

const (
	ScriptTailGetResponseEnvelopeSuccessTrue ScriptTailGetResponseEnvelopeSuccess = true
)

func (r ScriptTailGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
