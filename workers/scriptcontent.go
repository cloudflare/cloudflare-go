// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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
)

// ScriptContentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptContentService] method instead.
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
func (r *ScriptContentService) Update(ctx context.Context, scriptName string, params ScriptContentUpdateParams, opts ...option.RequestOption) (res *Script, err error) {
	var env ScriptContentUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch script content only
func (r *ScriptContentService) Get(ctx context.Context, scriptName string, query ScriptContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content/v2", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ScriptContentUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[WorkerMetadataParam] `json:"metadata,required"`
	CfWorkerBodyPart       param.Field[string]              `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]              `header:"CF-WORKER-MAIN-MODULE-PART"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Script                                     `json:"result"`
	JSON    scriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptContentUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptContentUpdateResponseEnvelope]
type scriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type ScriptContentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
