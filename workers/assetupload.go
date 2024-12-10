// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AssetUploadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetUploadService] method instead.
type AssetUploadService struct {
	Options []option.RequestOption
}

// NewAssetUploadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetUploadService(opts ...option.RequestOption) (r *AssetUploadService) {
	r = &AssetUploadService{}
	r.Options = opts
	return
}

// Upload assets ahead of creating a Worker version.
func (r *AssetUploadService) New(ctx context.Context, params AssetUploadNewParams, opts ...option.RequestOption) (res *AssetUploadNewResponse, err error) {
	var env AssetUploadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/assets/upload", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AssetUploadNewResponse struct {
	// A "completion" JWT which can be redeemed when creating a Worker version.
	JWT  string                     `json:"jwt"`
	JSON assetUploadNewResponseJSON `json:"-"`
}

// assetUploadNewResponseJSON contains the JSON metadata for the struct
// [AssetUploadNewResponse]
type assetUploadNewResponseJSON struct {
	JWT         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type AssetUploadNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether the file contents are base64-encoded. Must be `true`.
	Base64 param.Field[AssetUploadNewParamsBase64] `query:"base64,required"`
	// Base-64 encoded contents of the file. The content type of the file should be
	// included to ensure a valid "Content-Type" header is included in asset responses.
	AnyFileHash param.Field[[]string] `json:"<any file hash>"`
}

func (r AssetUploadNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [AssetUploadNewParams]'s query parameters as `url.Values`.
func (r AssetUploadNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Whether the file contents are base64-encoded. Must be `true`.
type AssetUploadNewParamsBase64 bool

const (
	AssetUploadNewParamsBase64True AssetUploadNewParamsBase64 = true
)

func (r AssetUploadNewParamsBase64) IsKnown() bool {
	switch r {
	case AssetUploadNewParamsBase64True:
		return true
	}
	return false
}

type AssetUploadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AssetUploadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AssetUploadNewResponse                `json:"result"`
	JSON    assetUploadNewResponseEnvelopeJSON    `json:"-"`
}

// assetUploadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AssetUploadNewResponseEnvelope]
type assetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AssetUploadNewResponseEnvelopeSuccess bool

const (
	AssetUploadNewResponseEnvelopeSuccessTrue AssetUploadNewResponseEnvelopeSuccess = true
)

func (r AssetUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AssetUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
