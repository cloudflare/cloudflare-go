// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DispatchNamespaceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceService] method instead.
type DispatchNamespaceService struct {
	Options []option.RequestOption
	Scripts *DispatchNamespaceScriptService
}

// NewDispatchNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatchNamespaceService(opts ...option.RequestOption) (r *DispatchNamespaceService) {
	r = &DispatchNamespaceService{}
	r.Options = opts
	r.Scripts = NewDispatchNamespaceScriptService(opts...)
	return
}

// Create a new Workers for Platforms namespace.
func (r *DispatchNamespaceService) New(ctx context.Context, params DispatchNamespaceNewParams, opts ...option.RequestOption) (res *DispatchNamespaceNewResponse, err error) {
	var env DispatchNamespaceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of Workers for Platforms namespaces.
func (r *DispatchNamespaceService) List(ctx context.Context, query DispatchNamespaceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", query.AccountID)
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

// Fetch a list of Workers for Platforms namespaces.
func (r *DispatchNamespaceService) ListAutoPaging(ctx context.Context, query DispatchNamespaceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Workers for Platforms namespace.
func (r *DispatchNamespaceService) Delete(ctx context.Context, dispatchNamespace string, body DispatchNamespaceDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceDeleteResponse, err error) {
	var env DispatchNamespaceDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", body.AccountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Workers for Platforms namespace.
func (r *DispatchNamespaceService) Get(ctx context.Context, dispatchNamespace string, query DispatchNamespaceGetParams, opts ...option.RequestOption) (res *DispatchNamespaceGetResponse, err error) {
	var env DispatchNamespaceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", query.AccountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceNewResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace
	ScriptCount int64                            `json:"script_count"`
	JSON        dispatchNamespaceNewResponseJSON `json:"-"`
}

// dispatchNamespaceNewResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceNewResponse]
type dispatchNamespaceNewResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	ScriptCount   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceListResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace
	ScriptCount int64                             `json:"script_count"`
	JSON        dispatchNamespaceListResponseJSON `json:"-"`
}

// dispatchNamespaceListResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceListResponse]
type dispatchNamespaceListResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	ScriptCount   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponse = interface{}

type DispatchNamespaceGetResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace
	ScriptCount int64                            `json:"script_count"`
	JSON        dispatchNamespaceGetResponseJSON `json:"-"`
}

// dispatchNamespaceGetResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceGetResponse]
type dispatchNamespaceGetResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	ScriptCount   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the dispatch namespace
	Name param.Field[string] `json:"name"`
}

func (r DispatchNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceNewResponse                `json:"result"`
	JSON    dispatchNamespaceNewResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceNewResponseEnvelope]
type dispatchNamespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceNewResponseEnvelopeSuccess bool

const (
	DispatchNamespaceNewResponseEnvelopeSuccessTrue DispatchNamespaceNewResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceDeleteResponse                `json:"result,nullable"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceDeleteResponseEnvelope]
type dispatchNamespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceDeleteResponseEnvelopeSuccessTrue DispatchNamespaceDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceGetResponse                `json:"result"`
	JSON    dispatchNamespaceGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceGetResponseEnvelope]
type dispatchNamespaceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceGetResponseEnvelopeSuccessTrue DispatchNamespaceGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
