// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/workers"
)

// DispatchNamespaceScriptContentService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptContentService] method instead.
type DispatchNamespaceScriptContentService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptContentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptContentService(opts ...option.RequestOption) (r *DispatchNamespaceScriptContentService) {
	r = &DispatchNamespaceScriptContentService{}
	r.Options = opts
	return
}

// Put script content for a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptContentService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptContentUpdateParams, opts ...option.RequestOption) (res *workers.Script, err error) {
	var env DispatchNamespaceScriptContentUpdateResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch script content from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptContentService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DispatchNamespaceScriptContentUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[workers.WorkerMetadataParam] `json:"metadata,required"`
	CfWorkerBodyPart       param.Field[string]                      `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                      `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r DispatchNamespaceScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DispatchNamespaceScriptContentUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  workers.Script                                              `json:"result"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptContentUpdateResponseEnvelope]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
