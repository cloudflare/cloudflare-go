// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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

// DispatchNamespaceScriptSecretService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptSecretService] method instead.
type DispatchNamespaceScriptSecretService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptSecretService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSecretService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSecretService) {
	r = &DispatchNamespaceScriptSecretService{}
	r.Options = opts
	return
}

// Put secrets to a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSecretUpdateParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSecretUpdateResponse, err error) {
	var env DispatchNamespaceScriptSecretUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List secrets from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) List(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSecretListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceScriptSecretListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", query.AccountID, dispatchNamespace, scriptName)
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

// List secrets from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) ListAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSecretListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceScriptSecretListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, dispatchNamespace, scriptName, query, opts...))
}

// Get secret from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) Get(ctx context.Context, dispatchNamespace string, scriptName string, secretName string, query DispatchNamespaceScriptSecretGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSecretGetResponse, err error) {
	var env DispatchNamespaceScriptSecretGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", query.AccountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersSecretModelParam struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name param.Field[string] `json:"name"`
	// The value of the secret.
	Text param.Field[string] `json:"text"`
	// The type of secret to put.
	Type param.Field[WorkersSecretModelType] `json:"type"`
}

func (r WorkersSecretModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of secret to put.
type WorkersSecretModelType string

const (
	WorkersSecretModelTypeSecretText WorkersSecretModelType = "secret_text"
)

func (r WorkersSecretModelType) IsKnown() bool {
	switch r {
	case WorkersSecretModelTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type DispatchNamespaceScriptSecretUpdateResponseType `json:"type"`
	JSON dispatchNamespaceScriptSecretUpdateResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSecretUpdateResponse]
type dispatchNamespaceScriptSecretUpdateResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type DispatchNamespaceScriptSecretUpdateResponseType string

const (
	DispatchNamespaceScriptSecretUpdateResponseTypeSecretText DispatchNamespaceScriptSecretUpdateResponseType = "secret_text"
)

func (r DispatchNamespaceScriptSecretUpdateResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretListResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type DispatchNamespaceScriptSecretListResponseType `json:"type"`
	JSON dispatchNamespaceScriptSecretListResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSecretListResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSecretListResponse]
type dispatchNamespaceScriptSecretListResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type DispatchNamespaceScriptSecretListResponseType string

const (
	DispatchNamespaceScriptSecretListResponseTypeSecretText DispatchNamespaceScriptSecretListResponseType = "secret_text"
)

func (r DispatchNamespaceScriptSecretListResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretGetResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type DispatchNamespaceScriptSecretGetResponseType `json:"type"`
	JSON dispatchNamespaceScriptSecretGetResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSecretGetResponse]
type dispatchNamespaceScriptSecretGetResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type DispatchNamespaceScriptSecretGetResponseType string

const (
	DispatchNamespaceScriptSecretGetResponseTypeSecretText DispatchNamespaceScriptSecretGetResponseType = "secret_text"
)

func (r DispatchNamespaceScriptSecretGetResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateParams struct {
	// Identifier
	AccountID          param.Field[string]     `path:"account_id,required"`
	WorkersSecretModel WorkersSecretModelParam `json:"workers_secret_model,required"`
}

func (r DispatchNamespaceScriptSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.WorkersSecretModel)
}

type DispatchNamespaceScriptSecretUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSecretUpdateResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSecretUpdateResponseEnvelope]
type dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSecretGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSecretGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSecretGetResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSecretGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSecretGetResponseEnvelope]
type dispatchNamespaceScriptSecretGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSecretGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
