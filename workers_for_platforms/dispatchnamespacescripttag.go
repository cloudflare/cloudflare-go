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

// DispatchNamespaceScriptTagService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptTagService] method instead.
type DispatchNamespaceScriptTagService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptTagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptTagService(opts ...option.RequestOption) (r *DispatchNamespaceScriptTagService) {
	r = &DispatchNamespaceScriptTagService{}
	r.Options = opts
	return
}

// Put script tags for a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptTagService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptTagUpdateParams, opts ...option.RequestOption) (res *[]string, err error) {
	var env DispatchNamespaceScriptTagUpdateResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch tags from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptTagService) List(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptTagListParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags", query.AccountID, dispatchNamespace, scriptName)
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

// Fetch tags from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptTagService) ListAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptTagListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, dispatchNamespace, scriptName, query, opts...))
}

// Delete script tag for a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptTagService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, tag string, body DispatchNamespaceScriptTagDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptTagDeleteResponse, err error) {
	var env DispatchNamespaceScriptTagDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
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
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags/%s", body.AccountID, dispatchNamespace, scriptName, tag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptTagDeleteResponse = interface{}

type DispatchNamespaceScriptTagUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Tags to help you manage your Workers
	Body []string `json:"body,required"`
}

func (r DispatchNamespaceScriptTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DispatchNamespaceScriptTagUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptTagUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  []string                                                `json:"result"`
	JSON    dispatchNamespaceScriptTagUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptTagUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptTagUpdateResponseEnvelope]
type dispatchNamespaceScriptTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptTagUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptTagUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptTagUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptTagUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptTagUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptTagListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptTagDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptTagDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptTagDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptTagDeleteResponse                `json:"result,nullable"`
	JSON    dispatchNamespaceScriptTagDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptTagDeleteResponseEnvelope]
type dispatchNamespaceScriptTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptTagDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptTagDeleteResponseEnvelopeSuccessTrue DispatchNamespaceScriptTagDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptTagDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptTagDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
