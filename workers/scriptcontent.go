// File generated from our OpenAPI spec by Stainless.

package workers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apiform"
	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
	Metadata               param.Field[ScriptContentUpdateParamsMetadata] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                            `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                            `header:"CF-WORKER-MAIN-MODULE-PART"`
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

// JSON encoded metadata about the uploaded parts and Worker configuration.
type ScriptContentUpdateParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r ScriptContentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptContentUpdateResponseEnvelope struct {
	Errors   []ScriptContentUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptContentUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersScript                                 `json:"result,required"`
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

type ScriptContentUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    scriptContentUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptContentUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptContentUpdateResponseEnvelopeErrors]
type scriptContentUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptContentUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    scriptContentUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptContentUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptContentUpdateResponseEnvelopeMessages]
type scriptContentUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptContentUpdateResponseEnvelopeSuccess bool

const (
	ScriptContentUpdateResponseEnvelopeSuccessTrue ScriptContentUpdateResponseEnvelopeSuccess = true
)
