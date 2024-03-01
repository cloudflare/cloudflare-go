// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDLPDatasetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDLPDatasetService]
// method instead.
type ZeroTrustDLPDatasetService struct {
	Options []option.RequestOption
	Upload  *ZeroTrustDLPDatasetUploadService
}

// NewZeroTrustDLPDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDLPDatasetService(opts ...option.RequestOption) (r *ZeroTrustDLPDatasetService) {
	r = &ZeroTrustDLPDatasetService{}
	r.Options = opts
	r.Upload = NewZeroTrustDLPDatasetUploadService(opts...)
	return
}

// Create a new dataset.
func (r *ZeroTrustDLPDatasetService) New(ctx context.Context, params ZeroTrustDLPDatasetNewParams, opts ...option.RequestOption) (res *ZeroTrustDLPDatasetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPDatasetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update details about a dataset.
func (r *ZeroTrustDLPDatasetService) Update(ctx context.Context, datasetID string, params ZeroTrustDLPDatasetUpdateParams, opts ...option.RequestOption) (res *ZeroTrustDLPDatasetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPDatasetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", params.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all datasets with information about available versions.
func (r *ZeroTrustDLPDatasetService) List(ctx context.Context, query ZeroTrustDLPDatasetListParams, opts ...option.RequestOption) (res *[]ZeroTrustDLPDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPDatasetListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a dataset.
//
// This deletes all versions of the dataset.
func (r *ZeroTrustDLPDatasetService) Delete(ctx context.Context, datasetID string, body ZeroTrustDLPDatasetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetch a specific dataset with information about available versions.
func (r *ZeroTrustDLPDatasetService) Get(ctx context.Context, datasetID string, query ZeroTrustDLPDatasetGetParams, opts ...option.RequestOption) (res *ZeroTrustDLPDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPDatasetGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDLPDatasetNewResponse struct {
	Dataset  ZeroTrustDLPDatasetNewResponseDataset `json:"dataset,required"`
	MaxCells int64                                 `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                             `json:"secret" format:"password"`
	JSON   zeroTrustDLPDatasetNewResponseJSON `json:"-"`
}

// zeroTrustDLPDatasetNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDLPDatasetNewResponse]
type zeroTrustDLPDatasetNewResponseJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetNewResponseDataset struct {
	ID          string                                        `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                                     `json:"created_at,required" format:"date-time"`
	Name        string                                        `json:"name,required"`
	NumCells    int64                                         `json:"num_cells,required"`
	Secret      bool                                          `json:"secret,required"`
	Status      ZeroTrustDLPDatasetNewResponseDatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time                                     `json:"updated_at,required" format:"date-time"`
	Uploads     []ZeroTrustDLPDatasetNewResponseDatasetUpload `json:"uploads,required"`
	Description string                                        `json:"description,nullable"`
	JSON        zeroTrustDLPDatasetNewResponseDatasetJSON     `json:"-"`
}

// zeroTrustDLPDatasetNewResponseDatasetJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetNewResponseDataset]
type zeroTrustDLPDatasetNewResponseDatasetJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	NumCells    apijson.Field
	Secret      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Uploads     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetNewResponseDatasetStatus string

const (
	ZeroTrustDLPDatasetNewResponseDatasetStatusEmpty     ZeroTrustDLPDatasetNewResponseDatasetStatus = "empty"
	ZeroTrustDLPDatasetNewResponseDatasetStatusUploading ZeroTrustDLPDatasetNewResponseDatasetStatus = "uploading"
	ZeroTrustDLPDatasetNewResponseDatasetStatusFailed    ZeroTrustDLPDatasetNewResponseDatasetStatus = "failed"
	ZeroTrustDLPDatasetNewResponseDatasetStatusComplete  ZeroTrustDLPDatasetNewResponseDatasetStatus = "complete"
)

type ZeroTrustDLPDatasetNewResponseDatasetUpload struct {
	NumCells int64                                              `json:"num_cells,required"`
	Status   ZeroTrustDLPDatasetNewResponseDatasetUploadsStatus `json:"status,required"`
	Version  int64                                              `json:"version,required"`
	JSON     zeroTrustDLPDatasetNewResponseDatasetUploadJSON    `json:"-"`
}

// zeroTrustDLPDatasetNewResponseDatasetUploadJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetNewResponseDatasetUpload]
type zeroTrustDLPDatasetNewResponseDatasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponseDatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetNewResponseDatasetUploadsStatus string

const (
	ZeroTrustDLPDatasetNewResponseDatasetUploadsStatusEmpty     ZeroTrustDLPDatasetNewResponseDatasetUploadsStatus = "empty"
	ZeroTrustDLPDatasetNewResponseDatasetUploadsStatusUploading ZeroTrustDLPDatasetNewResponseDatasetUploadsStatus = "uploading"
	ZeroTrustDLPDatasetNewResponseDatasetUploadsStatusFailed    ZeroTrustDLPDatasetNewResponseDatasetUploadsStatus = "failed"
	ZeroTrustDLPDatasetNewResponseDatasetUploadsStatusComplete  ZeroTrustDLPDatasetNewResponseDatasetUploadsStatus = "complete"
)

type ZeroTrustDLPDatasetUpdateResponse struct {
	ID          string                                    `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                                 `json:"created_at,required" format:"date-time"`
	Name        string                                    `json:"name,required"`
	NumCells    int64                                     `json:"num_cells,required"`
	Secret      bool                                      `json:"secret,required"`
	Status      ZeroTrustDLPDatasetUpdateResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                                 `json:"updated_at,required" format:"date-time"`
	Uploads     []ZeroTrustDLPDatasetUpdateResponseUpload `json:"uploads,required"`
	Description string                                    `json:"description,nullable"`
	JSON        zeroTrustDLPDatasetUpdateResponseJSON     `json:"-"`
}

// zeroTrustDLPDatasetUpdateResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDLPDatasetUpdateResponse]
type zeroTrustDLPDatasetUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	NumCells    apijson.Field
	Secret      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Uploads     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetUpdateResponseStatus string

const (
	ZeroTrustDLPDatasetUpdateResponseStatusEmpty     ZeroTrustDLPDatasetUpdateResponseStatus = "empty"
	ZeroTrustDLPDatasetUpdateResponseStatusUploading ZeroTrustDLPDatasetUpdateResponseStatus = "uploading"
	ZeroTrustDLPDatasetUpdateResponseStatusFailed    ZeroTrustDLPDatasetUpdateResponseStatus = "failed"
	ZeroTrustDLPDatasetUpdateResponseStatusComplete  ZeroTrustDLPDatasetUpdateResponseStatus = "complete"
)

type ZeroTrustDLPDatasetUpdateResponseUpload struct {
	NumCells int64                                          `json:"num_cells,required"`
	Status   ZeroTrustDLPDatasetUpdateResponseUploadsStatus `json:"status,required"`
	Version  int64                                          `json:"version,required"`
	JSON     zeroTrustDLPDatasetUpdateResponseUploadJSON    `json:"-"`
}

// zeroTrustDLPDatasetUpdateResponseUploadJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetUpdateResponseUpload]
type zeroTrustDLPDatasetUpdateResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUpdateResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetUpdateResponseUploadsStatus string

const (
	ZeroTrustDLPDatasetUpdateResponseUploadsStatusEmpty     ZeroTrustDLPDatasetUpdateResponseUploadsStatus = "empty"
	ZeroTrustDLPDatasetUpdateResponseUploadsStatusUploading ZeroTrustDLPDatasetUpdateResponseUploadsStatus = "uploading"
	ZeroTrustDLPDatasetUpdateResponseUploadsStatusFailed    ZeroTrustDLPDatasetUpdateResponseUploadsStatus = "failed"
	ZeroTrustDLPDatasetUpdateResponseUploadsStatusComplete  ZeroTrustDLPDatasetUpdateResponseUploadsStatus = "complete"
)

type ZeroTrustDLPDatasetListResponse struct {
	ID          string                                  `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at,required" format:"date-time"`
	Name        string                                  `json:"name,required"`
	NumCells    int64                                   `json:"num_cells,required"`
	Secret      bool                                    `json:"secret,required"`
	Status      ZeroTrustDLPDatasetListResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                               `json:"updated_at,required" format:"date-time"`
	Uploads     []ZeroTrustDLPDatasetListResponseUpload `json:"uploads,required"`
	Description string                                  `json:"description,nullable"`
	JSON        zeroTrustDLPDatasetListResponseJSON     `json:"-"`
}

// zeroTrustDLPDatasetListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDLPDatasetListResponse]
type zeroTrustDLPDatasetListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	NumCells    apijson.Field
	Secret      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Uploads     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetListResponseStatus string

const (
	ZeroTrustDLPDatasetListResponseStatusEmpty     ZeroTrustDLPDatasetListResponseStatus = "empty"
	ZeroTrustDLPDatasetListResponseStatusUploading ZeroTrustDLPDatasetListResponseStatus = "uploading"
	ZeroTrustDLPDatasetListResponseStatusFailed    ZeroTrustDLPDatasetListResponseStatus = "failed"
	ZeroTrustDLPDatasetListResponseStatusComplete  ZeroTrustDLPDatasetListResponseStatus = "complete"
)

type ZeroTrustDLPDatasetListResponseUpload struct {
	NumCells int64                                        `json:"num_cells,required"`
	Status   ZeroTrustDLPDatasetListResponseUploadsStatus `json:"status,required"`
	Version  int64                                        `json:"version,required"`
	JSON     zeroTrustDLPDatasetListResponseUploadJSON    `json:"-"`
}

// zeroTrustDLPDatasetListResponseUploadJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetListResponseUpload]
type zeroTrustDLPDatasetListResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetListResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetListResponseUploadsStatus string

const (
	ZeroTrustDLPDatasetListResponseUploadsStatusEmpty     ZeroTrustDLPDatasetListResponseUploadsStatus = "empty"
	ZeroTrustDLPDatasetListResponseUploadsStatusUploading ZeroTrustDLPDatasetListResponseUploadsStatus = "uploading"
	ZeroTrustDLPDatasetListResponseUploadsStatusFailed    ZeroTrustDLPDatasetListResponseUploadsStatus = "failed"
	ZeroTrustDLPDatasetListResponseUploadsStatusComplete  ZeroTrustDLPDatasetListResponseUploadsStatus = "complete"
)

type ZeroTrustDLPDatasetGetResponse struct {
	ID          string                                 `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                              `json:"created_at,required" format:"date-time"`
	Name        string                                 `json:"name,required"`
	NumCells    int64                                  `json:"num_cells,required"`
	Secret      bool                                   `json:"secret,required"`
	Status      ZeroTrustDLPDatasetGetResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                              `json:"updated_at,required" format:"date-time"`
	Uploads     []ZeroTrustDLPDatasetGetResponseUpload `json:"uploads,required"`
	Description string                                 `json:"description,nullable"`
	JSON        zeroTrustDLPDatasetGetResponseJSON     `json:"-"`
}

// zeroTrustDLPDatasetGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDLPDatasetGetResponse]
type zeroTrustDLPDatasetGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	NumCells    apijson.Field
	Secret      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Uploads     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetGetResponseStatus string

const (
	ZeroTrustDLPDatasetGetResponseStatusEmpty     ZeroTrustDLPDatasetGetResponseStatus = "empty"
	ZeroTrustDLPDatasetGetResponseStatusUploading ZeroTrustDLPDatasetGetResponseStatus = "uploading"
	ZeroTrustDLPDatasetGetResponseStatusFailed    ZeroTrustDLPDatasetGetResponseStatus = "failed"
	ZeroTrustDLPDatasetGetResponseStatusComplete  ZeroTrustDLPDatasetGetResponseStatus = "complete"
)

type ZeroTrustDLPDatasetGetResponseUpload struct {
	NumCells int64                                       `json:"num_cells,required"`
	Status   ZeroTrustDLPDatasetGetResponseUploadsStatus `json:"status,required"`
	Version  int64                                       `json:"version,required"`
	JSON     zeroTrustDLPDatasetGetResponseUploadJSON    `json:"-"`
}

// zeroTrustDLPDatasetGetResponseUploadJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetGetResponseUpload]
type zeroTrustDLPDatasetGetResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetGetResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetGetResponseUploadsStatus string

const (
	ZeroTrustDLPDatasetGetResponseUploadsStatusEmpty     ZeroTrustDLPDatasetGetResponseUploadsStatus = "empty"
	ZeroTrustDLPDatasetGetResponseUploadsStatusUploading ZeroTrustDLPDatasetGetResponseUploadsStatus = "uploading"
	ZeroTrustDLPDatasetGetResponseUploadsStatusFailed    ZeroTrustDLPDatasetGetResponseUploadsStatus = "failed"
	ZeroTrustDLPDatasetGetResponseUploadsStatusComplete  ZeroTrustDLPDatasetGetResponseUploadsStatus = "complete"
)

type ZeroTrustDLPDatasetNewParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder. If
	// false, the response has no secret and the dataset is uploaded in plaintext.
	Secret param.Field[bool] `json:"secret"`
}

func (r ZeroTrustDLPDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPDatasetNewResponseEnvelope struct {
	Errors     []ZeroTrustDLPDatasetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []ZeroTrustDLPDatasetNewResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                             `json:"success,required"`
	Result     ZeroTrustDLPDatasetNewResponse                   `json:"result"`
	ResultInfo ZeroTrustDLPDatasetNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPDatasetNewResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPDatasetNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetNewResponseEnvelope]
type zeroTrustDLPDatasetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustDLPDatasetNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPDatasetNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetNewResponseEnvelopeErrors]
type zeroTrustDLPDatasetNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDLPDatasetNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPDatasetNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetNewResponseEnvelopeMessages]
type zeroTrustDLPDatasetNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetNewResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                `json:"total_count,required"`
	JSON       zeroTrustDLPDatasetNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPDatasetNewResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetNewResponseEnvelopeResultInfo]
type zeroTrustDLPDatasetNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r ZeroTrustDLPDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPDatasetUpdateResponseEnvelope struct {
	Errors     []ZeroTrustDLPDatasetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []ZeroTrustDLPDatasetUpdateResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                                `json:"success,required"`
	Result     ZeroTrustDLPDatasetUpdateResponse                   `json:"result"`
	ResultInfo ZeroTrustDLPDatasetUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPDatasetUpdateResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPDatasetUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetUpdateResponseEnvelope]
type zeroTrustDLPDatasetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDLPDatasetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPDatasetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetUpdateResponseEnvelopeErrors]
type zeroTrustDLPDatasetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDLPDatasetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPDatasetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetUpdateResponseEnvelopeMessages]
type zeroTrustDLPDatasetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetUpdateResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                   `json:"total_count,required"`
	JSON       zeroTrustDLPDatasetUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPDatasetUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDLPDatasetUpdateResponseEnvelopeResultInfo]
type zeroTrustDLPDatasetUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPDatasetListResponseEnvelope struct {
	Errors     []ZeroTrustDLPDatasetListResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []ZeroTrustDLPDatasetListResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                              `json:"success,required"`
	Result     []ZeroTrustDLPDatasetListResponse                 `json:"result"`
	ResultInfo ZeroTrustDLPDatasetListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPDatasetListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPDatasetListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetListResponseEnvelope]
type zeroTrustDLPDatasetListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustDLPDatasetListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPDatasetListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetListResponseEnvelopeErrors]
type zeroTrustDLPDatasetListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDLPDatasetListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPDatasetListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetListResponseEnvelopeMessages]
type zeroTrustDLPDatasetListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetListResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                 `json:"total_count,required"`
	JSON       zeroTrustDLPDatasetListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPDatasetListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetListResponseEnvelopeResultInfo]
type zeroTrustDLPDatasetListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPDatasetGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPDatasetGetResponseEnvelope struct {
	Errors     []ZeroTrustDLPDatasetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []ZeroTrustDLPDatasetGetResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                             `json:"success,required"`
	Result     ZeroTrustDLPDatasetGetResponse                   `json:"result"`
	ResultInfo ZeroTrustDLPDatasetGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPDatasetGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPDatasetGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetGetResponseEnvelope]
type zeroTrustDLPDatasetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustDLPDatasetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPDatasetGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetGetResponseEnvelopeErrors]
type zeroTrustDLPDatasetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDLPDatasetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPDatasetGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetGetResponseEnvelopeMessages]
type zeroTrustDLPDatasetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPDatasetGetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                `json:"total_count,required"`
	JSON       zeroTrustDLPDatasetGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPDatasetGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustDLPDatasetGetResponseEnvelopeResultInfo]
type zeroTrustDLPDatasetGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
