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

// ScriptVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptVersionService] method
// instead.
type ScriptVersionService struct {
	Options []option.RequestOption
}

// NewScriptVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptVersionService(opts ...option.RequestOption) (r *ScriptVersionService) {
	r = &ScriptVersionService{}
	r.Options = opts
	return
}

// Upload a Worker Version without deploying to Cloudflare's network.
func (r *ScriptVersionService) New(ctx context.Context, scriptName string, params ScriptVersionNewParams, opts ...option.RequestOption) (res *ScriptVersionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of Worker Versions. The first version in the list is the latest version.
func (r *ScriptVersionService) List(ctx context.Context, scriptName string, query ScriptVersionListParams, opts ...option.RequestOption) (res *ScriptVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Version Detail
func (r *ScriptVersionService) Get(ctx context.Context, scriptName string, versionID string, query ScriptVersionGetParams, opts ...option.RequestOption) (res *ScriptVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions/%s", query.AccountID, scriptName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptVersionNewResponse struct {
	Resources interface{}                  `json:"resources,required"`
	ID        string                       `json:"id"`
	Metadata  interface{}                  `json:"metadata"`
	Number    float64                      `json:"number"`
	JSON      scriptVersionNewResponseJSON `json:"-"`
}

// scriptVersionNewResponseJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponse]
type scriptVersionNewResponseJSON struct {
	Resources   apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionListResponse struct {
	Items []ScriptVersionListResponseItem `json:"items"`
	JSON  scriptVersionListResponseJSON   `json:"-"`
}

// scriptVersionListResponseJSON contains the JSON metadata for the struct
// [ScriptVersionListResponse]
type scriptVersionListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionListResponseItem struct {
	ID       string                            `json:"id"`
	Metadata interface{}                       `json:"metadata"`
	Number   float64                           `json:"number"`
	JSON     scriptVersionListResponseItemJSON `json:"-"`
}

// scriptVersionListResponseItemJSON contains the JSON metadata for the struct
// [ScriptVersionListResponseItem]
type scriptVersionListResponseItemJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseItemJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponse struct {
	Resources interface{}                  `json:"resources,required"`
	ID        string                       `json:"id"`
	Metadata  interface{}                  `json:"metadata"`
	Number    float64                      `json:"number"`
	JSON      scriptVersionGetResponseJSON `json:"-"`
}

// scriptVersionGetResponseJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponse]
type scriptVersionGetResponseJSON struct {
	Resources   apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module`.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptVersionNewParamsMetadata] `json:"metadata"`
}

func (r ScriptVersionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type ScriptVersionNewParamsMetadata struct {
	Annotations param.Field[ScriptVersionNewParamsMetadataAnnotations] `json:"annotations"`
	// List of bindings available to the worker.
	Bindings param.Field[[]interface{}] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Usage model to apply to invocations.
	UsageModel param.Field[ScriptVersionNewParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptVersionNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataAnnotations struct {
	// Human-readable message about the version.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r ScriptVersionNewParamsMetadataAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type ScriptVersionNewParamsMetadataUsageModel string

const (
	ScriptVersionNewParamsMetadataUsageModelStandard ScriptVersionNewParamsMetadataUsageModel = "standard"
)

func (r ScriptVersionNewParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   ScriptVersionNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionNewResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseEnvelope]
type scriptVersionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionNewResponseEnvelopeSuccess bool

const (
	ScriptVersionNewResponseEnvelopeSuccessTrue ScriptVersionNewResponseEnvelopeSuccess = true
)

func (r ScriptVersionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptVersionListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionListResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   ScriptVersionListResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionListResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionListResponseEnvelope]
type scriptVersionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionListResponseEnvelopeSuccess bool

const (
	ScriptVersionListResponseEnvelopeSuccessTrue ScriptVersionListResponseEnvelopeSuccess = true
)

func (r ScriptVersionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptVersionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   ScriptVersionGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionGetResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseEnvelope]
type scriptVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionGetResponseEnvelopeSuccess bool

const (
	ScriptVersionGetResponseEnvelopeSuccessTrue ScriptVersionGetResponseEnvelopeSuccess = true
)

func (r ScriptVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
