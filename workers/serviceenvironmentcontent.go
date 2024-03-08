// File generated from our OpenAPI spec by Stainless.

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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ServiceEnvironmentContentService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewServiceEnvironmentContentService] method instead.
type ServiceEnvironmentContentService struct {
	Options []option.RequestOption
}

// NewServiceEnvironmentContentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewServiceEnvironmentContentService(opts ...option.RequestOption) (r *ServiceEnvironmentContentService) {
	r = &ServiceEnvironmentContentService{}
	r.Options = opts
	return
}

// Put script content from a worker with an environment
func (r *ServiceEnvironmentContentService) Update(ctx context.Context, serviceName string, environmentName string, params ServiceEnvironmentContentUpdateParams, opts ...option.RequestOption) (res *WorkersScript, err error) {
	opts = append(r.Options[:], opts...)
	var env ServiceEnvironmentContentUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/content", params.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script content from a worker with an environment
func (r *ServiceEnvironmentContentService) Get(ctx context.Context, serviceName string, environmentName string, query ServiceEnvironmentContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/content", query.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ServiceEnvironmentContentUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be
	// present. This should be referenced either in the metadata as `main_module`
	// (esm)/`body_part` (service worker) or as a header `CF-WORKER-MAIN-MODULE-PART`
	// (esm) /`CF-WORKER-BODY-PART` (service worker) by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[ServiceEnvironmentContentUpdateParamsMetadata] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                        `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                        `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r ServiceEnvironmentContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type ServiceEnvironmentContentUpdateParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r ServiceEnvironmentContentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentContentUpdateResponseEnvelope struct {
	Errors   []ServiceEnvironmentContentUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ServiceEnvironmentContentUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersScript                                             `json:"result,required"`
	// Whether the API call was successful
	Success ServiceEnvironmentContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    serviceEnvironmentContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// serviceEnvironmentContentUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ServiceEnvironmentContentUpdateResponseEnvelope]
type serviceEnvironmentContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentContentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentContentUpdateResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    serviceEnvironmentContentUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// serviceEnvironmentContentUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ServiceEnvironmentContentUpdateResponseEnvelopeErrors]
type serviceEnvironmentContentUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentContentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentContentUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentContentUpdateResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    serviceEnvironmentContentUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// serviceEnvironmentContentUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ServiceEnvironmentContentUpdateResponseEnvelopeMessages]
type serviceEnvironmentContentUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentContentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentContentUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ServiceEnvironmentContentUpdateResponseEnvelopeSuccess bool

const (
	ServiceEnvironmentContentUpdateResponseEnvelopeSuccessTrue ServiceEnvironmentContentUpdateResponseEnvelopeSuccess = true
)

type ServiceEnvironmentContentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
