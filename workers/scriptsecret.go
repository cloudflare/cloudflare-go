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
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ScriptSecretService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptSecretService] method instead.
type ScriptSecretService struct {
	Options []option.RequestOption
}

// NewScriptSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSecretService(opts ...option.RequestOption) (r *ScriptSecretService) {
	r = &ScriptSecretService{}
	r.Options = opts
	return
}

// Add a secret to a script.
func (r *ScriptSecretService) Update(ctx context.Context, scriptName string, params ScriptSecretUpdateParams, opts ...option.RequestOption) (res *ScriptSecretUpdateResponse, err error) {
	var env ScriptSecretUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List secrets bound to a script.
func (r *ScriptSecretService) List(ctx context.Context, scriptName string, query ScriptSecretListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ScriptSecretListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets", query.AccountID, scriptName)
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

// List secrets bound to a script.
func (r *ScriptSecretService) ListAutoPaging(ctx context.Context, scriptName string, query ScriptSecretListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ScriptSecretListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, scriptName, query, opts...))
}

// Remove a secret from a script.
func (r *ScriptSecretService) Delete(ctx context.Context, scriptName string, secretName string, body ScriptSecretDeleteParams, opts ...option.RequestOption) (res *ScriptSecretDeleteResponse, err error) {
	var env ScriptSecretDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets/%s", body.AccountID, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a given secret binding (value omitted) on a script.
func (r *ScriptSecretService) Get(ctx context.Context, scriptName string, secretName string, query ScriptSecretGetParams, opts ...option.RequestOption) (res *ScriptSecretGetResponse, err error) {
	var env ScriptSecretGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets/%s", query.AccountID, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptSecretUpdateResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type ScriptSecretUpdateResponseType `json:"type"`
	JSON scriptSecretUpdateResponseJSON `json:"-"`
}

// scriptSecretUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptSecretUpdateResponse]
type scriptSecretUpdateResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type ScriptSecretUpdateResponseType string

const (
	ScriptSecretUpdateResponseTypeSecretText ScriptSecretUpdateResponseType = "secret_text"
)

func (r ScriptSecretUpdateResponseType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretListResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type ScriptSecretListResponseType `json:"type"`
	JSON scriptSecretListResponseJSON `json:"-"`
}

// scriptSecretListResponseJSON contains the JSON metadata for the struct
// [ScriptSecretListResponse]
type scriptSecretListResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type ScriptSecretListResponseType string

const (
	ScriptSecretListResponseTypeSecretText ScriptSecretListResponseType = "secret_text"
)

func (r ScriptSecretListResponseType) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretDeleteResponse = interface{}

type ScriptSecretGetResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type ScriptSecretGetResponseType `json:"type"`
	JSON scriptSecretGetResponseJSON `json:"-"`
}

// scriptSecretGetResponseJSON contains the JSON metadata for the struct
// [ScriptSecretGetResponse]
type scriptSecretGetResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type ScriptSecretGetResponseType string

const (
	ScriptSecretGetResponseTypeSecretText ScriptSecretGetResponseType = "secret_text"
)

func (r ScriptSecretGetResponseType) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name param.Field[string] `json:"name"`
	// The value of the secret.
	Text param.Field[string] `json:"text"`
	// The type of secret to put.
	Type param.Field[ScriptSecretUpdateParamsType] `json:"type"`
}

func (r ScriptSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of secret to put.
type ScriptSecretUpdateParamsType string

const (
	ScriptSecretUpdateParamsTypeSecretText ScriptSecretUpdateParamsType = "secret_text"
)

func (r ScriptSecretUpdateParamsType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptSecretUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptSecretUpdateResponse                `json:"result"`
	JSON    scriptSecretUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptSecretUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSecretUpdateResponseEnvelope]
type scriptSecretUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSecretUpdateResponseEnvelopeSuccess bool

const (
	ScriptSecretUpdateResponseEnvelopeSuccessTrue ScriptSecretUpdateResponseEnvelopeSuccess = true
)

func (r ScriptSecretUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSecretListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSecretDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSecretDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptSecretDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptSecretDeleteResponse                `json:"result,nullable"`
	JSON    scriptSecretDeleteResponseEnvelopeJSON    `json:"-"`
}

// scriptSecretDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSecretDeleteResponseEnvelope]
type scriptSecretDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSecretDeleteResponseEnvelopeSuccess bool

const (
	ScriptSecretDeleteResponseEnvelopeSuccessTrue ScriptSecretDeleteResponseEnvelopeSuccess = true
)

func (r ScriptSecretDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSecretDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSecretGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSecretGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptSecretGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptSecretGetResponse                `json:"result"`
	JSON    scriptSecretGetResponseEnvelopeJSON    `json:"-"`
}

// scriptSecretGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSecretGetResponseEnvelope]
type scriptSecretGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSecretGetResponseEnvelopeSuccess bool

const (
	ScriptSecretGetResponseEnvelopeSuccessTrue ScriptSecretGetResponseEnvelopeSuccess = true
)

func (r ScriptSecretGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
