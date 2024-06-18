// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSecretUpdateResponseEnvelope
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

// Fetch secrets from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) List(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSecretListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceScriptSecretListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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

// Fetch secrets from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) ListAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSecretListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceScriptSecretListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, dispatchNamespace, scriptName, query, opts...))
}

type DispatchNamespaceScriptSecretUpdateResponse struct {
	// The name of this secret, this is what will be to access it inside the Worker.
	Name string `json:"name"`
	// The type of secret to put.
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

// The type of secret to put.
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
	// The name of this secret, this is what will be to access it inside the Worker.
	Name string `json:"name"`
	// The type of secret to put.
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

// The type of secret to put.
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

type DispatchNamespaceScriptSecretUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of this secret, this is what will be to access it inside the Worker.
	Name param.Field[string] `json:"name"`
	// The value of the secret.
	Text param.Field[string] `json:"text"`
	// The type of secret to put.
	Type param.Field[DispatchNamespaceScriptSecretUpdateParamsType] `json:"type"`
}

func (r DispatchNamespaceScriptSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of secret to put.
type DispatchNamespaceScriptSecretUpdateParamsType string

const (
	DispatchNamespaceScriptSecretUpdateParamsTypeSecretText DispatchNamespaceScriptSecretUpdateParamsType = "secret_text"
)

func (r DispatchNamespaceScriptSecretUpdateParamsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsTypeSecretText:
		return true
	}
	return false
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
