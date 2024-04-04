// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ScriptContentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptContentService] method
// instead.
type ScriptContentService struct {
	Options []option.RequestOption
}

// NewScriptContentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptContentService(opts ...option.RequestOption) (r *ScriptContentService) {
	r = &ScriptContentService{}
	r.Options = opts
	return
}

// Put script content without touching config or metadata
func (r *ScriptContentService) Update(ctx context.Context, scriptName string, params ScriptContentUpdateParams, opts ...option.RequestOption) (res *WorkersScript, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptContentUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptContentUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be
	// present. This should be referenced either in the metadata as `main_module`
	// (esm)/`body_part` (service worker) or as a header `CF-WORKER-MAIN-MODULE-PART`
	// (esm) /`CF-WORKER-BODY-PART` (service worker) by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[shared.UnnamedSchemaRefEe1e79edcb234d14c4dd266880f2fd24Param] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                                       `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                                       `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r ScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type ScriptContentUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   WorkersScript                                             `json:"result,required"`
	// Whether the API call was successful
	Success ScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptContentUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptContentUpdateResponseEnvelope]
type scriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptContentUpdateResponseEnvelopeSuccess bool

const (
	ScriptContentUpdateResponseEnvelopeSuccessTrue ScriptContentUpdateResponseEnvelopeSuccess = true
)

func (r ScriptContentUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptContentUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
