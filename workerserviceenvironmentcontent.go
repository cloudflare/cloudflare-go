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

// WorkerServiceEnvironmentContentService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkerServiceEnvironmentContentService] method instead.
type WorkerServiceEnvironmentContentService struct {
	Options []option.RequestOption
}

// NewWorkerServiceEnvironmentContentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerServiceEnvironmentContentService(opts ...option.RequestOption) (r *WorkerServiceEnvironmentContentService) {
	r = &WorkerServiceEnvironmentContentService{}
	r.Options = opts
	return
}

// Get script content from a worker with an environment
func (r *WorkerServiceEnvironmentContentService) Get(ctx context.Context, accountID string, serviceName string, environmentName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/content", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put script content from a worker with an environment
func (r *WorkerServiceEnvironmentContentService) Replace(ctx context.Context, accountID string, serviceName string, environmentName string, params WorkerServiceEnvironmentContentReplaceParams, opts ...option.RequestOption) (res *WorkerServiceEnvironmentContentReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerServiceEnvironmentContentReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/content", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerServiceEnvironmentContentReplaceResponse struct {
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
	TailConsumers []WorkerServiceEnvironmentContentReplaceResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                             `json:"usage_model"`
	JSON       workerServiceEnvironmentContentReplaceResponseJSON `json:"-"`
}

// workerServiceEnvironmentContentReplaceResponseJSON contains the JSON metadata
// for the struct [WorkerServiceEnvironmentContentReplaceResponse]
type workerServiceEnvironmentContentReplaceResponseJSON struct {
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

func (r *WorkerServiceEnvironmentContentReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentContentReplaceResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                         `json:"namespace"`
	JSON      workerServiceEnvironmentContentReplaceResponseTailConsumerJSON `json:"-"`
}

// workerServiceEnvironmentContentReplaceResponseTailConsumerJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentContentReplaceResponseTailConsumer]
type workerServiceEnvironmentContentReplaceResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentContentReplaceResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentContentReplaceParams struct {
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be
	// present. This should be referenced either in the metadata as `main_module`
	// (esm)/`body_part` (service worker) or as a header `CF-WORKER-MAIN-MODULE-PART`
	// (esm) /`CF-WORKER-BODY-PART` (service worker) by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[WorkerServiceEnvironmentContentReplaceParamsMetadata] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                               `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                               `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r WorkerServiceEnvironmentContentReplaceParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type WorkerServiceEnvironmentContentReplaceParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r WorkerServiceEnvironmentContentReplaceParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentContentReplaceResponseEnvelope struct {
	Errors   []WorkerServiceEnvironmentContentReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerServiceEnvironmentContentReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerServiceEnvironmentContentReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerServiceEnvironmentContentReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerServiceEnvironmentContentReplaceResponseEnvelopeJSON    `json:"-"`
}

// workerServiceEnvironmentContentReplaceResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentContentReplaceResponseEnvelope]
type workerServiceEnvironmentContentReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentContentReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentContentReplaceResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    workerServiceEnvironmentContentReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// workerServiceEnvironmentContentReplaceResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentContentReplaceResponseEnvelopeErrors]
type workerServiceEnvironmentContentReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentContentReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentContentReplaceResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    workerServiceEnvironmentContentReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// workerServiceEnvironmentContentReplaceResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentContentReplaceResponseEnvelopeMessages]
type workerServiceEnvironmentContentReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentContentReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentContentReplaceResponseEnvelopeSuccess bool

const (
	WorkerServiceEnvironmentContentReplaceResponseEnvelopeSuccessTrue WorkerServiceEnvironmentContentReplaceResponseEnvelopeSuccess = true
)
