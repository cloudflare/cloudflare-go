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
func (r *ZeroTrustDLPDatasetService) New(ctx context.Context, params ZeroTrustDLPDatasetNewParams, opts ...option.RequestOption) (res *DLPDatasetCreation, err error) {
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
func (r *ZeroTrustDLPDatasetService) Update(ctx context.Context, datasetID string, params ZeroTrustDLPDatasetUpdateParams, opts ...option.RequestOption) (res *DLPDataset, err error) {
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
func (r *ZeroTrustDLPDatasetService) List(ctx context.Context, query ZeroTrustDLPDatasetListParams, opts ...option.RequestOption) (res *DLPDatasetArray, err error) {
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
func (r *ZeroTrustDLPDatasetService) Get(ctx context.Context, datasetID string, query ZeroTrustDLPDatasetGetParams, opts ...option.RequestOption) (res *DLPDataset, err error) {
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

type DLPDataset struct {
	ID          string             `json:"id,required" format:"uuid"`
	CreatedAt   time.Time          `json:"created_at,required" format:"date-time"`
	Name        string             `json:"name,required"`
	NumCells    int64              `json:"num_cells,required"`
	Secret      bool               `json:"secret,required"`
	Status      DLPDatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time          `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetUpload `json:"uploads,required"`
	Description string             `json:"description,nullable"`
	JSON        dlpDatasetJSON     `json:"-"`
}

// dlpDatasetJSON contains the JSON metadata for the struct [DLPDataset]
type dlpDatasetJSON struct {
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

func (r *DLPDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetStatus string

const (
	DLPDatasetStatusEmpty     DLPDatasetStatus = "empty"
	DLPDatasetStatusUploading DLPDatasetStatus = "uploading"
	DLPDatasetStatusFailed    DLPDatasetStatus = "failed"
	DLPDatasetStatusComplete  DLPDatasetStatus = "complete"
)

type DLPDatasetUpload struct {
	NumCells int64                   `json:"num_cells,required"`
	Status   DLPDatasetUploadsStatus `json:"status,required"`
	Version  int64                   `json:"version,required"`
	JSON     dlpDatasetUploadJSON    `json:"-"`
}

// dlpDatasetUploadJSON contains the JSON metadata for the struct
// [DLPDatasetUpload]
type dlpDatasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadsStatus string

const (
	DLPDatasetUploadsStatusEmpty     DLPDatasetUploadsStatus = "empty"
	DLPDatasetUploadsStatusUploading DLPDatasetUploadsStatus = "uploading"
	DLPDatasetUploadsStatusFailed    DLPDatasetUploadsStatus = "failed"
	DLPDatasetUploadsStatusComplete  DLPDatasetUploadsStatus = "complete"
)

type DLPDatasetArray []DLPDataset

type DLPDatasetCreation struct {
	Dataset  DLPDataset `json:"dataset,required"`
	MaxCells int64      `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                 `json:"secret" format:"password"`
	JSON   dlpDatasetCreationJSON `json:"-"`
}

// dlpDatasetCreationJSON contains the JSON metadata for the struct
// [DLPDatasetCreation]
type dlpDatasetCreationJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	Result     DLPDatasetCreation                               `json:"result"`
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
	Result     DLPDataset                                          `json:"result"`
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
	Result     DLPDatasetArray                                   `json:"result"`
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
	Result     DLPDataset                                       `json:"result"`
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
