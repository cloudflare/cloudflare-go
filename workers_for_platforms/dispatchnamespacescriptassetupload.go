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
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DispatchNamespaceScriptAssetUploadService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptAssetUploadService] method instead.
type DispatchNamespaceScriptAssetUploadService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptAssetUploadService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDispatchNamespaceScriptAssetUploadService(opts ...option.RequestOption) (r *DispatchNamespaceScriptAssetUploadService) {
	r = &DispatchNamespaceScriptAssetUploadService{}
	r.Options = opts
	return
}

// Start uploading a collection of assets for use in a Worker version.
func (r *DispatchNamespaceScriptAssetUploadService) New(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptAssetUploadNewParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptAssetUploadNewResponse, err error) {
	var env DispatchNamespaceScriptAssetUploadNewResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/assets-upload-session", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptAssetUploadNewResponse struct {
	// The requests to make to upload assets.
	Buckets [][]string `json:"buckets"`
	// A JWT to use as authentication for uploading assets.
	JWT  string                                            `json:"jwt"`
	JSON dispatchNamespaceScriptAssetUploadNewResponseJSON `json:"-"`
}

// dispatchNamespaceScriptAssetUploadNewResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptAssetUploadNewResponse]
type dispatchNamespaceScriptAssetUploadNewResponseJSON struct {
	Buckets     apijson.Field
	JWT         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptAssetUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptAssetUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptAssetUploadNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A manifest ([path]: {hash, size}) map of files to upload. As an example,
	// `/blog/hello-world.html` would be a valid path key.
	Manifest param.Field[map[string]DispatchNamespaceScriptAssetUploadNewParamsManifest] `json:"manifest"`
}

func (r DispatchNamespaceScriptAssetUploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptAssetUploadNewParamsManifest struct {
	// The hash of the file.
	Hash param.Field[string] `json:"hash"`
	// The size of the file in bytes.
	Size param.Field[int64] `json:"size"`
}

func (r DispatchNamespaceScriptAssetUploadNewParamsManifest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptAssetUploadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptAssetUploadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptAssetUploadNewResponse                `json:"result"`
	JSON    dispatchNamespaceScriptAssetUploadNewResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptAssetUploadNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptAssetUploadNewResponseEnvelope]
type dispatchNamespaceScriptAssetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptAssetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptAssetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptAssetUploadNewResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptAssetUploadNewResponseEnvelopeSuccessTrue DispatchNamespaceScriptAssetUploadNewResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptAssetUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptAssetUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
