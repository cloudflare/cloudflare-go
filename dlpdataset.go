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

// DlpDatasetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDlpDatasetService] method instead.
type DlpDatasetService struct {
	Options []option.RequestOption
}

// NewDlpDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDlpDatasetService(opts ...option.RequestOption) (r *DlpDatasetService) {
	r = &DlpDatasetService{}
	r.Options = opts
	return
}

// Create a new dataset.
func (r *DlpDatasetService) New(ctx context.Context, accountID string, body DlpDatasetNewParams, opts ...option.RequestOption) (res *DlpDatasetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a specific dataset with information about available versions.
func (r *DlpDatasetService) Get(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *DlpDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update details about a dataset.
func (r *DlpDatasetService) Update(ctx context.Context, accountID string, datasetID string, body DlpDatasetUpdateParams, opts ...option.RequestOption) (res *DlpDatasetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetch all datasets with information about available versions.
func (r *DlpDatasetService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DlpDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a dataset.
//
// This deletes all versions of the dataset.
func (r *DlpDatasetService) Delete(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Upload a new version of a dataset.
func (r *DlpDatasetService) Upload(ctx context.Context, accountID string, datasetID string, version int64, opts ...option.RequestOption) (res *DlpDatasetUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", accountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Prepare to upload a new version of a dataset.
func (r *DlpDatasetService) UploadPrepare(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *DlpDatasetUploadPrepareResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type DlpDatasetNewResponse struct {
	Errors     []DlpDatasetNewResponseError    `json:"errors"`
	Messages   []DlpDatasetNewResponseMessage  `json:"messages"`
	Result     DlpDatasetNewResponseResult     `json:"result"`
	ResultInfo DlpDatasetNewResponseResultInfo `json:"result_info"`
	Success    bool                            `json:"success"`
	JSON       dlpDatasetNewResponseJSON       `json:"-"`
}

// dlpDatasetNewResponseJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponse]
type dlpDatasetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dlpDatasetNewResponseErrorJSON `json:"-"`
}

// dlpDatasetNewResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseError]
type dlpDatasetNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dlpDatasetNewResponseMessageJSON `json:"-"`
}

// dlpDatasetNewResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseMessage]
type dlpDatasetNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseResult struct {
	Dataset  DlpDatasetNewResponseResultDataset `json:"dataset,required"`
	MaxCells int64                              `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                          `json:"secret" format:"password"`
	JSON   dlpDatasetNewResponseResultJSON `json:"-"`
}

// dlpDatasetNewResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseResult]
type dlpDatasetNewResponseResultJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseResultDataset struct {
	ID          string                                     `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                                  `json:"created_at,required" format:"date-time"`
	Name        string                                     `json:"name,required"`
	NumCells    int64                                      `json:"num_cells,required"`
	Secret      bool                                       `json:"secret,required"`
	Status      DlpDatasetNewResponseResultDatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time                                  `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetNewResponseResultDatasetUpload `json:"uploads,required"`
	Description string                                     `json:"description,nullable"`
	JSON        dlpDatasetNewResponseResultDatasetJSON     `json:"-"`
}

// dlpDatasetNewResponseResultDatasetJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseResultDataset]
type dlpDatasetNewResponseResultDatasetJSON struct {
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

func (r *DlpDatasetNewResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseResultDatasetStatus string

const (
	DlpDatasetNewResponseResultDatasetStatusEmpty     DlpDatasetNewResponseResultDatasetStatus = "empty"
	DlpDatasetNewResponseResultDatasetStatusUploading DlpDatasetNewResponseResultDatasetStatus = "uploading"
	DlpDatasetNewResponseResultDatasetStatusFailed    DlpDatasetNewResponseResultDatasetStatus = "failed"
	DlpDatasetNewResponseResultDatasetStatusComplete  DlpDatasetNewResponseResultDatasetStatus = "complete"
)

type DlpDatasetNewResponseResultDatasetUpload struct {
	NumCells int64                                           `json:"num_cells,required"`
	Status   DlpDatasetNewResponseResultDatasetUploadsStatus `json:"status,required"`
	Version  int64                                           `json:"version,required"`
	JSON     dlpDatasetNewResponseResultDatasetUploadJSON    `json:"-"`
}

// dlpDatasetNewResponseResultDatasetUploadJSON contains the JSON metadata for the
// struct [DlpDatasetNewResponseResultDatasetUpload]
type dlpDatasetNewResponseResultDatasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseResultDatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseResultDatasetUploadsStatus string

const (
	DlpDatasetNewResponseResultDatasetUploadsStatusEmpty     DlpDatasetNewResponseResultDatasetUploadsStatus = "empty"
	DlpDatasetNewResponseResultDatasetUploadsStatusUploading DlpDatasetNewResponseResultDatasetUploadsStatus = "uploading"
	DlpDatasetNewResponseResultDatasetUploadsStatusFailed    DlpDatasetNewResponseResultDatasetUploadsStatus = "failed"
	DlpDatasetNewResponseResultDatasetUploadsStatusComplete  DlpDatasetNewResponseResultDatasetUploadsStatus = "complete"
)

type DlpDatasetNewResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                               `json:"total_count,required"`
	JSON       dlpDatasetNewResponseResultInfoJSON `json:"-"`
}

// dlpDatasetNewResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseResultInfo]
type dlpDatasetNewResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponse struct {
	Errors     []DlpDatasetGetResponseError    `json:"errors"`
	Messages   []DlpDatasetGetResponseMessage  `json:"messages"`
	Result     DlpDatasetGetResponseResult     `json:"result"`
	ResultInfo DlpDatasetGetResponseResultInfo `json:"result_info"`
	Success    bool                            `json:"success"`
	JSON       dlpDatasetGetResponseJSON       `json:"-"`
}

// dlpDatasetGetResponseJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponse]
type dlpDatasetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dlpDatasetGetResponseErrorJSON `json:"-"`
}

// dlpDatasetGetResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseError]
type dlpDatasetGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dlpDatasetGetResponseMessageJSON `json:"-"`
}

// dlpDatasetGetResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseMessage]
type dlpDatasetGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseResult struct {
	ID          string                              `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                           `json:"created_at,required" format:"date-time"`
	Name        string                              `json:"name,required"`
	NumCells    int64                               `json:"num_cells,required"`
	Secret      bool                                `json:"secret,required"`
	Status      DlpDatasetGetResponseResultStatus   `json:"status,required"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetGetResponseResultUpload `json:"uploads,required"`
	Description string                              `json:"description,nullable"`
	JSON        dlpDatasetGetResponseResultJSON     `json:"-"`
}

// dlpDatasetGetResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseResult]
type dlpDatasetGetResponseResultJSON struct {
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

func (r *DlpDatasetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseResultStatus string

const (
	DlpDatasetGetResponseResultStatusEmpty     DlpDatasetGetResponseResultStatus = "empty"
	DlpDatasetGetResponseResultStatusUploading DlpDatasetGetResponseResultStatus = "uploading"
	DlpDatasetGetResponseResultStatusFailed    DlpDatasetGetResponseResultStatus = "failed"
	DlpDatasetGetResponseResultStatusComplete  DlpDatasetGetResponseResultStatus = "complete"
)

type DlpDatasetGetResponseResultUpload struct {
	NumCells int64                                    `json:"num_cells,required"`
	Status   DlpDatasetGetResponseResultUploadsStatus `json:"status,required"`
	Version  int64                                    `json:"version,required"`
	JSON     dlpDatasetGetResponseResultUploadJSON    `json:"-"`
}

// dlpDatasetGetResponseResultUploadJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseResultUpload]
type dlpDatasetGetResponseResultUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseResultUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseResultUploadsStatus string

const (
	DlpDatasetGetResponseResultUploadsStatusEmpty     DlpDatasetGetResponseResultUploadsStatus = "empty"
	DlpDatasetGetResponseResultUploadsStatusUploading DlpDatasetGetResponseResultUploadsStatus = "uploading"
	DlpDatasetGetResponseResultUploadsStatusFailed    DlpDatasetGetResponseResultUploadsStatus = "failed"
	DlpDatasetGetResponseResultUploadsStatusComplete  DlpDatasetGetResponseResultUploadsStatus = "complete"
)

type DlpDatasetGetResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                               `json:"total_count,required"`
	JSON       dlpDatasetGetResponseResultInfoJSON `json:"-"`
}

// dlpDatasetGetResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseResultInfo]
type dlpDatasetGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponse struct {
	Errors     []DlpDatasetUpdateResponseError    `json:"errors"`
	Messages   []DlpDatasetUpdateResponseMessage  `json:"messages"`
	Result     DlpDatasetUpdateResponseResult     `json:"result"`
	ResultInfo DlpDatasetUpdateResponseResultInfo `json:"result_info"`
	Success    bool                               `json:"success"`
	JSON       dlpDatasetUpdateResponseJSON       `json:"-"`
}

// dlpDatasetUpdateResponseJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponse]
type dlpDatasetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    dlpDatasetUpdateResponseErrorJSON `json:"-"`
}

// dlpDatasetUpdateResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponseError]
type dlpDatasetUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    dlpDatasetUpdateResponseMessageJSON `json:"-"`
}

// dlpDatasetUpdateResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponseMessage]
type dlpDatasetUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseResult struct {
	ID          string                                 `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                              `json:"created_at,required" format:"date-time"`
	Name        string                                 `json:"name,required"`
	NumCells    int64                                  `json:"num_cells,required"`
	Secret      bool                                   `json:"secret,required"`
	Status      DlpDatasetUpdateResponseResultStatus   `json:"status,required"`
	UpdatedAt   time.Time                              `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetUpdateResponseResultUpload `json:"uploads,required"`
	Description string                                 `json:"description,nullable"`
	JSON        dlpDatasetUpdateResponseResultJSON     `json:"-"`
}

// dlpDatasetUpdateResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponseResult]
type dlpDatasetUpdateResponseResultJSON struct {
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

func (r *DlpDatasetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseResultStatus string

const (
	DlpDatasetUpdateResponseResultStatusEmpty     DlpDatasetUpdateResponseResultStatus = "empty"
	DlpDatasetUpdateResponseResultStatusUploading DlpDatasetUpdateResponseResultStatus = "uploading"
	DlpDatasetUpdateResponseResultStatusFailed    DlpDatasetUpdateResponseResultStatus = "failed"
	DlpDatasetUpdateResponseResultStatusComplete  DlpDatasetUpdateResponseResultStatus = "complete"
)

type DlpDatasetUpdateResponseResultUpload struct {
	NumCells int64                                       `json:"num_cells,required"`
	Status   DlpDatasetUpdateResponseResultUploadsStatus `json:"status,required"`
	Version  int64                                       `json:"version,required"`
	JSON     dlpDatasetUpdateResponseResultUploadJSON    `json:"-"`
}

// dlpDatasetUpdateResponseResultUploadJSON contains the JSON metadata for the
// struct [DlpDatasetUpdateResponseResultUpload]
type dlpDatasetUpdateResponseResultUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseResultUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseResultUploadsStatus string

const (
	DlpDatasetUpdateResponseResultUploadsStatusEmpty     DlpDatasetUpdateResponseResultUploadsStatus = "empty"
	DlpDatasetUpdateResponseResultUploadsStatusUploading DlpDatasetUpdateResponseResultUploadsStatus = "uploading"
	DlpDatasetUpdateResponseResultUploadsStatusFailed    DlpDatasetUpdateResponseResultUploadsStatus = "failed"
	DlpDatasetUpdateResponseResultUploadsStatusComplete  DlpDatasetUpdateResponseResultUploadsStatus = "complete"
)

type DlpDatasetUpdateResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                  `json:"total_count,required"`
	JSON       dlpDatasetUpdateResponseResultInfoJSON `json:"-"`
}

// dlpDatasetUpdateResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponseResultInfo]
type dlpDatasetUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponse struct {
	Errors     []DlpDatasetListResponseError    `json:"errors"`
	Messages   []DlpDatasetListResponseMessage  `json:"messages"`
	Result     []DlpDatasetListResponseResult   `json:"result"`
	ResultInfo DlpDatasetListResponseResultInfo `json:"result_info"`
	Success    bool                             `json:"success"`
	JSON       dlpDatasetListResponseJSON       `json:"-"`
}

// dlpDatasetListResponseJSON contains the JSON metadata for the struct
// [DlpDatasetListResponse]
type dlpDatasetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    dlpDatasetListResponseErrorJSON `json:"-"`
}

// dlpDatasetListResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseError]
type dlpDatasetListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    dlpDatasetListResponseMessageJSON `json:"-"`
}

// dlpDatasetListResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseMessage]
type dlpDatasetListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseResult struct {
	ID          string                               `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                            `json:"created_at,required" format:"date-time"`
	Name        string                               `json:"name,required"`
	NumCells    int64                                `json:"num_cells,required"`
	Secret      bool                                 `json:"secret,required"`
	Status      DlpDatasetListResponseResultStatus   `json:"status,required"`
	UpdatedAt   time.Time                            `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetListResponseResultUpload `json:"uploads,required"`
	Description string                               `json:"description,nullable"`
	JSON        dlpDatasetListResponseResultJSON     `json:"-"`
}

// dlpDatasetListResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseResult]
type dlpDatasetListResponseResultJSON struct {
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

func (r *DlpDatasetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseResultStatus string

const (
	DlpDatasetListResponseResultStatusEmpty     DlpDatasetListResponseResultStatus = "empty"
	DlpDatasetListResponseResultStatusUploading DlpDatasetListResponseResultStatus = "uploading"
	DlpDatasetListResponseResultStatusFailed    DlpDatasetListResponseResultStatus = "failed"
	DlpDatasetListResponseResultStatusComplete  DlpDatasetListResponseResultStatus = "complete"
)

type DlpDatasetListResponseResultUpload struct {
	NumCells int64                                     `json:"num_cells,required"`
	Status   DlpDatasetListResponseResultUploadsStatus `json:"status,required"`
	Version  int64                                     `json:"version,required"`
	JSON     dlpDatasetListResponseResultUploadJSON    `json:"-"`
}

// dlpDatasetListResponseResultUploadJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseResultUpload]
type dlpDatasetListResponseResultUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseResultUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseResultUploadsStatus string

const (
	DlpDatasetListResponseResultUploadsStatusEmpty     DlpDatasetListResponseResultUploadsStatus = "empty"
	DlpDatasetListResponseResultUploadsStatusUploading DlpDatasetListResponseResultUploadsStatus = "uploading"
	DlpDatasetListResponseResultUploadsStatusFailed    DlpDatasetListResponseResultUploadsStatus = "failed"
	DlpDatasetListResponseResultUploadsStatusComplete  DlpDatasetListResponseResultUploadsStatus = "complete"
)

type DlpDatasetListResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                `json:"total_count,required"`
	JSON       dlpDatasetListResponseResultInfoJSON `json:"-"`
}

// dlpDatasetListResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseResultInfo]
type dlpDatasetListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponse struct {
	Errors     []DlpDatasetUploadResponseError    `json:"errors"`
	Messages   []DlpDatasetUploadResponseMessage  `json:"messages"`
	Result     DlpDatasetUploadResponseResult     `json:"result"`
	ResultInfo DlpDatasetUploadResponseResultInfo `json:"result_info"`
	Success    bool                               `json:"success"`
	JSON       dlpDatasetUploadResponseJSON       `json:"-"`
}

// dlpDatasetUploadResponseJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponse]
type dlpDatasetUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    dlpDatasetUploadResponseErrorJSON `json:"-"`
}

// dlpDatasetUploadResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponseError]
type dlpDatasetUploadResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    dlpDatasetUploadResponseMessageJSON `json:"-"`
}

// dlpDatasetUploadResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponseMessage]
type dlpDatasetUploadResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseResult struct {
	ID          string                                 `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                              `json:"created_at,required" format:"date-time"`
	Name        string                                 `json:"name,required"`
	NumCells    int64                                  `json:"num_cells,required"`
	Secret      bool                                   `json:"secret,required"`
	Status      DlpDatasetUploadResponseResultStatus   `json:"status,required"`
	UpdatedAt   time.Time                              `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetUploadResponseResultUpload `json:"uploads,required"`
	Description string                                 `json:"description,nullable"`
	JSON        dlpDatasetUploadResponseResultJSON     `json:"-"`
}

// dlpDatasetUploadResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponseResult]
type dlpDatasetUploadResponseResultJSON struct {
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

func (r *DlpDatasetUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseResultStatus string

const (
	DlpDatasetUploadResponseResultStatusEmpty     DlpDatasetUploadResponseResultStatus = "empty"
	DlpDatasetUploadResponseResultStatusUploading DlpDatasetUploadResponseResultStatus = "uploading"
	DlpDatasetUploadResponseResultStatusFailed    DlpDatasetUploadResponseResultStatus = "failed"
	DlpDatasetUploadResponseResultStatusComplete  DlpDatasetUploadResponseResultStatus = "complete"
)

type DlpDatasetUploadResponseResultUpload struct {
	NumCells int64                                       `json:"num_cells,required"`
	Status   DlpDatasetUploadResponseResultUploadsStatus `json:"status,required"`
	Version  int64                                       `json:"version,required"`
	JSON     dlpDatasetUploadResponseResultUploadJSON    `json:"-"`
}

// dlpDatasetUploadResponseResultUploadJSON contains the JSON metadata for the
// struct [DlpDatasetUploadResponseResultUpload]
type dlpDatasetUploadResponseResultUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseResultUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseResultUploadsStatus string

const (
	DlpDatasetUploadResponseResultUploadsStatusEmpty     DlpDatasetUploadResponseResultUploadsStatus = "empty"
	DlpDatasetUploadResponseResultUploadsStatusUploading DlpDatasetUploadResponseResultUploadsStatus = "uploading"
	DlpDatasetUploadResponseResultUploadsStatusFailed    DlpDatasetUploadResponseResultUploadsStatus = "failed"
	DlpDatasetUploadResponseResultUploadsStatusComplete  DlpDatasetUploadResponseResultUploadsStatus = "complete"
)

type DlpDatasetUploadResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                  `json:"total_count,required"`
	JSON       dlpDatasetUploadResponseResultInfoJSON `json:"-"`
}

// dlpDatasetUploadResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponseResultInfo]
type dlpDatasetUploadResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponse struct {
	Errors     []DlpDatasetUploadPrepareResponseError    `json:"errors"`
	Messages   []DlpDatasetUploadPrepareResponseMessage  `json:"messages"`
	Result     DlpDatasetUploadPrepareResponseResult     `json:"result"`
	ResultInfo DlpDatasetUploadPrepareResponseResultInfo `json:"result_info"`
	Success    bool                                      `json:"success"`
	JSON       dlpDatasetUploadPrepareResponseJSON       `json:"-"`
}

// dlpDatasetUploadPrepareResponseJSON contains the JSON metadata for the struct
// [DlpDatasetUploadPrepareResponse]
type dlpDatasetUploadPrepareResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dlpDatasetUploadPrepareResponseErrorJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseErrorJSON contains the JSON metadata for the
// struct [DlpDatasetUploadPrepareResponseError]
type dlpDatasetUploadPrepareResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpDatasetUploadPrepareResponseMessageJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseMessageJSON contains the JSON metadata for the
// struct [DlpDatasetUploadPrepareResponseMessage]
type dlpDatasetUploadPrepareResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseResult struct {
	MaxCells int64                                     `json:"max_cells,required"`
	Version  int64                                     `json:"version,required"`
	Secret   string                                    `json:"secret" format:"password"`
	JSON     dlpDatasetUploadPrepareResponseResultJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseResultJSON contains the JSON metadata for the
// struct [DlpDatasetUploadPrepareResponseResult]
type dlpDatasetUploadPrepareResponseResultJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                         `json:"total_count,required"`
	JSON       dlpDatasetUploadPrepareResponseResultInfoJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseResultInfoJSON contains the JSON metadata for the
// struct [DlpDatasetUploadPrepareResponseResultInfo]
type dlpDatasetUploadPrepareResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewParams struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder. If
	// false, the response has no secret and the dataset is uploaded in plaintext.
	Secret param.Field[bool] `json:"secret"`
}

func (r DlpDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DlpDatasetUpdateParams struct {
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DlpDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
