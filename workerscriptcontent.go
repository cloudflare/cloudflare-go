// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiform"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerScriptContentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptContentService]
// method instead.
type WorkerScriptContentService struct {
	Options []option.RequestOption
}

// NewWorkerScriptContentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptContentService(opts ...option.RequestOption) (r *WorkerScriptContentService) {
	r = &WorkerScriptContentService{}
	r.Options = opts
	return
}

// Put script content without touching config or metadata
func (r *WorkerScriptContentService) Update(ctx context.Context, scriptName string, params WorkerScriptContentUpdateParams, opts ...option.RequestOption) (res *WorkerScriptContentUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptContentUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptContentUpdateResponse struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Deprecated. Deployment metadata for internal usage.
	PipelineHash string `json:"pipeline_hash"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptContentUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                `json:"usage_model"`
	JSON       workerScriptContentUpdateResponseJSON `json:"-"`
}

// workerScriptContentUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerScriptContentUpdateResponse]
type workerScriptContentUpdateResponseJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
	Logpush       apijson.Field
	ModifiedOn    apijson.Field
	PipelineHash  apijson.Field
	PlacementMode apijson.Field
	TailConsumers apijson.Field
	UsageModel    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptContentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptContentUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                            `json:"namespace"`
	JSON      workerScriptContentUpdateResponseTailConsumerJSON `json:"-"`
}

// workerScriptContentUpdateResponseTailConsumerJSON contains the JSON metadata for
// the struct [WorkerScriptContentUpdateResponseTailConsumer]
type workerScriptContentUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptContentUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptContentUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be
	// present. This should be referenced either in the metadata as `main_module`
	// (esm)/`body_part` (service worker) or as a header `CF-WORKER-MAIN-MODULE-PART`
	// (esm) /`CF-WORKER-BODY-PART` (service worker) by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[WorkerScriptContentUpdateParamsMetadata] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                  `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                  `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r WorkerScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type WorkerScriptContentUpdateParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r WorkerScriptContentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptContentUpdateResponseEnvelope struct {
	Errors   []WorkerScriptContentUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptContentUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptContentUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerScriptContentUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptContentUpdateResponseEnvelope]
type workerScriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptContentUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerScriptContentUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptContentUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerScriptContentUpdateResponseEnvelopeErrors]
type workerScriptContentUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptContentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptContentUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerScriptContentUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptContentUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptContentUpdateResponseEnvelopeMessages]
type workerScriptContentUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptContentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptContentUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptContentUpdateResponseEnvelopeSuccessTrue WorkerScriptContentUpdateResponseEnvelopeSuccess = true
)
