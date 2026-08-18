// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AssetService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r *AssetService) {
	r = &AssetService{}
	r.Options = opts
	return
}

// Check which of the provided file hashes are missing from the Pages asset store.
// Returns a list of missing hashes that need to be uploaded. Used as part of the
// Pages Direct Upload workflow.
//
// Authenticate with the JWT obtained from the upload-token endpoint: GET
// /accounts/{account_id}/pages/projects/{project_name}/upload-token
func (r *AssetService) CheckMissing(ctx context.Context, body AssetCheckMissingParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pages/assets/check-missing"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Check which of the provided file hashes are missing from the Pages asset store.
// Returns a list of missing hashes that need to be uploaded. Used as part of the
// Pages Direct Upload workflow.
//
// Authenticate with the JWT obtained from the upload-token endpoint: GET
// /accounts/{account_id}/pages/projects/{project_name}/upload-token
func (r *AssetService) CheckMissingAutoPaging(ctx context.Context, body AssetCheckMissingParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.CheckMissing(ctx, body, opts...))
}

// Upload one or more files to the Pages asset store. Each file is identified by
// its content hash and is uploaded using the same JSON shape as the Cloudflare KV
// bulk write API. Used as part of the Pages Direct Upload workflow.
//
// Authenticate with the JWT obtained from the upload-token endpoint: GET
// /accounts/{account_id}/pages/projects/{project_name}/upload-token
func (r *AssetService) Upload(ctx context.Context, body AssetUploadParams, opts ...option.RequestOption) (res *AssetUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pages/assets/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Register the provided file hashes as recently uploaded to the Pages asset store.
// Used as part of the Pages Direct Upload workflow so future deployments can avoid
// re-uploading files that are already present.
//
// Authenticate with the JWT obtained from the upload-token endpoint: GET
// /accounts/{account_id}/pages/projects/{project_name}/upload-token
func (r *AssetService) UpsertHashes(ctx context.Context, body AssetUpsertHashesParams, opts ...option.RequestOption) (res *AssetUpsertHashesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pages/assets/upsert-hashes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AssetUploadResponse struct {
	Errors   []AssetUploadResponseError   `json:"errors" api:"required"`
	Messages []AssetUploadResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AssetUploadResponseSuccess `json:"success" api:"required"`
	JSON    assetUploadResponseJSON    `json:"-"`
}

// assetUploadResponseJSON contains the JSON metadata for the struct
// [AssetUploadResponse]
type assetUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AssetUploadResponseError struct {
	Code             int64                           `json:"code" api:"required"`
	Message          string                          `json:"message" api:"required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           AssetUploadResponseErrorsSource `json:"source"`
	JSON             assetUploadResponseErrorJSON    `json:"-"`
}

// assetUploadResponseErrorJSON contains the JSON metadata for the struct
// [AssetUploadResponseError]
type assetUploadResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetUploadResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AssetUploadResponseErrorsSource struct {
	Pointer string                              `json:"pointer"`
	JSON    assetUploadResponseErrorsSourceJSON `json:"-"`
}

// assetUploadResponseErrorsSourceJSON contains the JSON metadata for the struct
// [AssetUploadResponseErrorsSource]
type assetUploadResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUploadResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AssetUploadResponseMessage struct {
	Code             int64                             `json:"code" api:"required"`
	Message          string                            `json:"message" api:"required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           AssetUploadResponseMessagesSource `json:"source"`
	JSON             assetUploadResponseMessageJSON    `json:"-"`
}

// assetUploadResponseMessageJSON contains the JSON metadata for the struct
// [AssetUploadResponseMessage]
type assetUploadResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetUploadResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AssetUploadResponseMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    assetUploadResponseMessagesSourceJSON `json:"-"`
}

// assetUploadResponseMessagesSourceJSON contains the JSON metadata for the struct
// [AssetUploadResponseMessagesSource]
type assetUploadResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUploadResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUploadResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AssetUploadResponseSuccess bool

const (
	AssetUploadResponseSuccessTrue AssetUploadResponseSuccess = true
)

func (r AssetUploadResponseSuccess) IsKnown() bool {
	switch r {
	case AssetUploadResponseSuccessTrue:
		return true
	}
	return false
}

type AssetUpsertHashesResponse struct {
	Errors   []AssetUpsertHashesResponseError   `json:"errors" api:"required"`
	Messages []AssetUpsertHashesResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AssetUpsertHashesResponseSuccess `json:"success" api:"required"`
	JSON    assetUpsertHashesResponseJSON    `json:"-"`
}

// assetUpsertHashesResponseJSON contains the JSON metadata for the struct
// [AssetUpsertHashesResponse]
type assetUpsertHashesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpsertHashesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpsertHashesResponseJSON) RawJSON() string {
	return r.raw
}

type AssetUpsertHashesResponseError struct {
	Code             int64                                 `json:"code" api:"required"`
	Message          string                                `json:"message" api:"required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AssetUpsertHashesResponseErrorsSource `json:"source"`
	JSON             assetUpsertHashesResponseErrorJSON    `json:"-"`
}

// assetUpsertHashesResponseErrorJSON contains the JSON metadata for the struct
// [AssetUpsertHashesResponseError]
type assetUpsertHashesResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetUpsertHashesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpsertHashesResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AssetUpsertHashesResponseErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    assetUpsertHashesResponseErrorsSourceJSON `json:"-"`
}

// assetUpsertHashesResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AssetUpsertHashesResponseErrorsSource]
type assetUpsertHashesResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpsertHashesResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpsertHashesResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AssetUpsertHashesResponseMessage struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AssetUpsertHashesResponseMessagesSource `json:"source"`
	JSON             assetUpsertHashesResponseMessageJSON    `json:"-"`
}

// assetUpsertHashesResponseMessageJSON contains the JSON metadata for the struct
// [AssetUpsertHashesResponseMessage]
type assetUpsertHashesResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetUpsertHashesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpsertHashesResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AssetUpsertHashesResponseMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    assetUpsertHashesResponseMessagesSourceJSON `json:"-"`
}

// assetUpsertHashesResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AssetUpsertHashesResponseMessagesSource]
type assetUpsertHashesResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpsertHashesResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpsertHashesResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AssetUpsertHashesResponseSuccess bool

const (
	AssetUpsertHashesResponseSuccessTrue AssetUpsertHashesResponseSuccess = true
)

func (r AssetUpsertHashesResponseSuccess) IsKnown() bool {
	switch r {
	case AssetUpsertHashesResponseSuccessTrue:
		return true
	}
	return false
}

type AssetCheckMissingParams struct {
	// List of file content hashes to check for existence in the asset store.
	Hashes param.Field[[]string] `json:"hashes" api:"required"`
}

func (r AssetCheckMissingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AssetUploadParams struct {
	Body []AssetUploadParamsBody `json:"body" api:"required"`
}

func (r AssetUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AssetUploadParamsBody struct {
	// Whether value is base64 encoded.
	Base64 param.Field[bool] `json:"base64" api:"required"`
	// File content hash used as the object key in the Pages asset store.
	Key      param.Field[string]                        `json:"key" api:"required"`
	Metadata param.Field[AssetUploadParamsBodyMetadata] `json:"metadata" api:"required"`
	// File content. When base64 is true, this value is base64 encoded.
	Value param.Field[string] `json:"value" api:"required"`
}

func (r AssetUploadParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AssetUploadParamsBodyMetadata struct {
	// MIME type for the uploaded file.
	ContentType param.Field[string] `json:"contentType" api:"required"`
}

func (r AssetUploadParamsBodyMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AssetUpsertHashesParams struct {
	// List of file content hashes to register in the asset store.
	Hashes param.Field[[]string] `json:"hashes" api:"required"`
}

func (r AssetUpsertHashesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
