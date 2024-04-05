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
func (r *ServiceEnvironmentContentService) Update(ctx context.Context, serviceName string, environmentName string, params ServiceEnvironmentContentUpdateParams, opts ...option.RequestOption) (res *Script, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	Metadata               param.Field[shared.UnnamedSchemaRefEe1e79edcb234d14c4dd266880f2fd24Param] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                                       `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                                       `header:"CF-WORKER-MAIN-MODULE-PART"`
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

type ServiceEnvironmentContentUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Script                                                    `json:"result,required"`
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

// Whether the API call was successful
type ServiceEnvironmentContentUpdateResponseEnvelopeSuccess bool

const (
	ServiceEnvironmentContentUpdateResponseEnvelopeSuccessTrue ServiceEnvironmentContentUpdateResponseEnvelopeSuccess = true
)

func (r ServiceEnvironmentContentUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentContentUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ServiceEnvironmentContentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
