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

// AccountDlpDatasetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDlpDatasetService] method
// instead.
type AccountDlpDatasetService struct {
	Options []option.RequestOption
	Upload  *AccountDlpDatasetUploadService
}

// NewAccountDlpDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpDatasetService(opts ...option.RequestOption) (r *AccountDlpDatasetService) {
	r = &AccountDlpDatasetService{}
	r.Options = opts
	r.Upload = NewAccountDlpDatasetUploadService(opts...)
	return
}

// Create a new dataset.
func (r *AccountDlpDatasetService) New(ctx context.Context, accountIdentifier string, body AccountDlpDatasetNewParams, opts ...option.RequestOption) (res *DlpDatasetCreationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a specific dataset with information about available versions.
func (r *AccountDlpDatasetService) Get(ctx context.Context, accountIdentifier string, datasetID string, opts ...option.RequestOption) (res *DlpDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountIdentifier, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update details about a dataset.
func (r *AccountDlpDatasetService) Update(ctx context.Context, accountIdentifier string, datasetID string, body AccountDlpDatasetUpdateParams, opts ...option.RequestOption) (res *DlpDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountIdentifier, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetch all datasets with information about available versions.
func (r *AccountDlpDatasetService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *DlpDatasetArrayResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a dataset.
//
// This deletes all versions of the dataset.
func (r *AccountDlpDatasetService) Delete(ctx context.Context, accountIdentifier string, datasetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountIdentifier, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type DlpDatasetArrayResponse struct {
	Errors     []DlpDatasetArrayResponseError    `json:"errors"`
	Messages   []DlpDatasetArrayResponseMessage  `json:"messages"`
	Result     []DlpDatasetArrayResponseResult   `json:"result"`
	ResultInfo DlpDatasetArrayResponseResultInfo `json:"result_info"`
	Success    bool                              `json:"success"`
	JSON       dlpDatasetArrayResponseJSON       `json:"-"`
}

// dlpDatasetArrayResponseJSON contains the JSON metadata for the struct
// [DlpDatasetArrayResponse]
type dlpDatasetArrayResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetArrayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetArrayResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dlpDatasetArrayResponseErrorJSON `json:"-"`
}

// dlpDatasetArrayResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetArrayResponseError]
type dlpDatasetArrayResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetArrayResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetArrayResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    dlpDatasetArrayResponseMessageJSON `json:"-"`
}

// dlpDatasetArrayResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetArrayResponseMessage]
type dlpDatasetArrayResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetArrayResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetArrayResponseResult struct {
	ID          string                                `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                             `json:"created_at,required" format:"date-time"`
	Name        string                                `json:"name,required"`
	NumCells    int64                                 `json:"num_cells,required"`
	Secret      bool                                  `json:"secret,required"`
	Status      DlpDatasetArrayResponseResultStatus   `json:"status,required"`
	UpdatedAt   time.Time                             `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetArrayResponseResultUpload `json:"uploads,required"`
	Description string                                `json:"description,nullable"`
	JSON        dlpDatasetArrayResponseResultJSON     `json:"-"`
}

// dlpDatasetArrayResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetArrayResponseResult]
type dlpDatasetArrayResponseResultJSON struct {
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

func (r *DlpDatasetArrayResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetArrayResponseResultStatus string

const (
	DlpDatasetArrayResponseResultStatusEmpty     DlpDatasetArrayResponseResultStatus = "empty"
	DlpDatasetArrayResponseResultStatusUploading DlpDatasetArrayResponseResultStatus = "uploading"
	DlpDatasetArrayResponseResultStatusFailed    DlpDatasetArrayResponseResultStatus = "failed"
	DlpDatasetArrayResponseResultStatusComplete  DlpDatasetArrayResponseResultStatus = "complete"
)

type DlpDatasetArrayResponseResultUpload struct {
	NumCells int64                                      `json:"num_cells,required"`
	Status   DlpDatasetArrayResponseResultUploadsStatus `json:"status,required"`
	Version  int64                                      `json:"version,required"`
	JSON     dlpDatasetArrayResponseResultUploadJSON    `json:"-"`
}

// dlpDatasetArrayResponseResultUploadJSON contains the JSON metadata for the
// struct [DlpDatasetArrayResponseResultUpload]
type dlpDatasetArrayResponseResultUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetArrayResponseResultUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetArrayResponseResultUploadsStatus string

const (
	DlpDatasetArrayResponseResultUploadsStatusEmpty     DlpDatasetArrayResponseResultUploadsStatus = "empty"
	DlpDatasetArrayResponseResultUploadsStatusUploading DlpDatasetArrayResponseResultUploadsStatus = "uploading"
	DlpDatasetArrayResponseResultUploadsStatusFailed    DlpDatasetArrayResponseResultUploadsStatus = "failed"
	DlpDatasetArrayResponseResultUploadsStatusComplete  DlpDatasetArrayResponseResultUploadsStatus = "complete"
)

type DlpDatasetArrayResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                 `json:"total_count,required"`
	JSON       dlpDatasetArrayResponseResultInfoJSON `json:"-"`
}

// dlpDatasetArrayResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetArrayResponseResultInfo]
type dlpDatasetArrayResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetArrayResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponse struct {
	Errors     []DlpDatasetCreationResponseError    `json:"errors"`
	Messages   []DlpDatasetCreationResponseMessage  `json:"messages"`
	Result     DlpDatasetCreationResponseResult     `json:"result"`
	ResultInfo DlpDatasetCreationResponseResultInfo `json:"result_info"`
	Success    bool                                 `json:"success"`
	JSON       dlpDatasetCreationResponseJSON       `json:"-"`
}

// dlpDatasetCreationResponseJSON contains the JSON metadata for the struct
// [DlpDatasetCreationResponse]
type dlpDatasetCreationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetCreationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    dlpDatasetCreationResponseErrorJSON `json:"-"`
}

// dlpDatasetCreationResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetCreationResponseError]
type dlpDatasetCreationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetCreationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dlpDatasetCreationResponseMessageJSON `json:"-"`
}

// dlpDatasetCreationResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetCreationResponseMessage]
type dlpDatasetCreationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetCreationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponseResult struct {
	Dataset  DlpDatasetCreationResponseResultDataset `json:"dataset,required"`
	MaxCells int64                                   `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                               `json:"secret" format:"password"`
	JSON   dlpDatasetCreationResponseResultJSON `json:"-"`
}

// dlpDatasetCreationResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetCreationResponseResult]
type dlpDatasetCreationResponseResultJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetCreationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponseResultDataset struct {
	ID          string                                          `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                                       `json:"created_at,required" format:"date-time"`
	Name        string                                          `json:"name,required"`
	NumCells    int64                                           `json:"num_cells,required"`
	Secret      bool                                            `json:"secret,required"`
	Status      DlpDatasetCreationResponseResultDatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time                                       `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetCreationResponseResultDatasetUpload `json:"uploads,required"`
	Description string                                          `json:"description,nullable"`
	JSON        dlpDatasetCreationResponseResultDatasetJSON     `json:"-"`
}

// dlpDatasetCreationResponseResultDatasetJSON contains the JSON metadata for the
// struct [DlpDatasetCreationResponseResultDataset]
type dlpDatasetCreationResponseResultDatasetJSON struct {
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

func (r *DlpDatasetCreationResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponseResultDatasetStatus string

const (
	DlpDatasetCreationResponseResultDatasetStatusEmpty     DlpDatasetCreationResponseResultDatasetStatus = "empty"
	DlpDatasetCreationResponseResultDatasetStatusUploading DlpDatasetCreationResponseResultDatasetStatus = "uploading"
	DlpDatasetCreationResponseResultDatasetStatusFailed    DlpDatasetCreationResponseResultDatasetStatus = "failed"
	DlpDatasetCreationResponseResultDatasetStatusComplete  DlpDatasetCreationResponseResultDatasetStatus = "complete"
)

type DlpDatasetCreationResponseResultDatasetUpload struct {
	NumCells int64                                                `json:"num_cells,required"`
	Status   DlpDatasetCreationResponseResultDatasetUploadsStatus `json:"status,required"`
	Version  int64                                                `json:"version,required"`
	JSON     dlpDatasetCreationResponseResultDatasetUploadJSON    `json:"-"`
}

// dlpDatasetCreationResponseResultDatasetUploadJSON contains the JSON metadata for
// the struct [DlpDatasetCreationResponseResultDatasetUpload]
type dlpDatasetCreationResponseResultDatasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetCreationResponseResultDatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetCreationResponseResultDatasetUploadsStatus string

const (
	DlpDatasetCreationResponseResultDatasetUploadsStatusEmpty     DlpDatasetCreationResponseResultDatasetUploadsStatus = "empty"
	DlpDatasetCreationResponseResultDatasetUploadsStatusUploading DlpDatasetCreationResponseResultDatasetUploadsStatus = "uploading"
	DlpDatasetCreationResponseResultDatasetUploadsStatusFailed    DlpDatasetCreationResponseResultDatasetUploadsStatus = "failed"
	DlpDatasetCreationResponseResultDatasetUploadsStatusComplete  DlpDatasetCreationResponseResultDatasetUploadsStatus = "complete"
)

type DlpDatasetCreationResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                    `json:"total_count,required"`
	JSON       dlpDatasetCreationResponseResultInfoJSON `json:"-"`
}

// dlpDatasetCreationResponseResultInfoJSON contains the JSON metadata for the
// struct [DlpDatasetCreationResponseResultInfo]
type dlpDatasetCreationResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetCreationResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewVersionResponse struct {
	Errors     []DlpDatasetNewVersionResponseError    `json:"errors"`
	Messages   []DlpDatasetNewVersionResponseMessage  `json:"messages"`
	Result     DlpDatasetNewVersionResponseResult     `json:"result"`
	ResultInfo DlpDatasetNewVersionResponseResultInfo `json:"result_info"`
	Success    bool                                   `json:"success"`
	JSON       dlpDatasetNewVersionResponseJSON       `json:"-"`
}

// dlpDatasetNewVersionResponseJSON contains the JSON metadata for the struct
// [DlpDatasetNewVersionResponse]
type dlpDatasetNewVersionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewVersionResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dlpDatasetNewVersionResponseErrorJSON `json:"-"`
}

// dlpDatasetNewVersionResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetNewVersionResponseError]
type dlpDatasetNewVersionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewVersionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewVersionResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpDatasetNewVersionResponseMessageJSON `json:"-"`
}

// dlpDatasetNewVersionResponseMessageJSON contains the JSON metadata for the
// struct [DlpDatasetNewVersionResponseMessage]
type dlpDatasetNewVersionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewVersionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewVersionResponseResult struct {
	MaxCells int64                                  `json:"max_cells,required"`
	Version  int64                                  `json:"version,required"`
	Secret   string                                 `json:"secret" format:"password"`
	JSON     dlpDatasetNewVersionResponseResultJSON `json:"-"`
}

// dlpDatasetNewVersionResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetNewVersionResponseResult]
type dlpDatasetNewVersionResponseResultJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewVersionResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                      `json:"total_count,required"`
	JSON       dlpDatasetNewVersionResponseResultInfoJSON `json:"-"`
}

// dlpDatasetNewVersionResponseResultInfoJSON contains the JSON metadata for the
// struct [DlpDatasetNewVersionResponseResultInfo]
type dlpDatasetNewVersionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewVersionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetResponse struct {
	Errors     []DlpDatasetResponseError    `json:"errors"`
	Messages   []DlpDatasetResponseMessage  `json:"messages"`
	Result     DlpDatasetResponseResult     `json:"result"`
	ResultInfo DlpDatasetResponseResultInfo `json:"result_info"`
	Success    bool                         `json:"success"`
	JSON       dlpDatasetResponseJSON       `json:"-"`
}

// dlpDatasetResponseJSON contains the JSON metadata for the struct
// [DlpDatasetResponse]
type dlpDatasetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    dlpDatasetResponseErrorJSON `json:"-"`
}

// dlpDatasetResponseErrorJSON contains the JSON metadata for the struct
// [DlpDatasetResponseError]
type dlpDatasetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    dlpDatasetResponseMessageJSON `json:"-"`
}

// dlpDatasetResponseMessageJSON contains the JSON metadata for the struct
// [DlpDatasetResponseMessage]
type dlpDatasetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetResponseResult struct {
	ID          string                           `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                        `json:"created_at,required" format:"date-time"`
	Name        string                           `json:"name,required"`
	NumCells    int64                            `json:"num_cells,required"`
	Secret      bool                             `json:"secret,required"`
	Status      DlpDatasetResponseResultStatus   `json:"status,required"`
	UpdatedAt   time.Time                        `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetResponseResultUpload `json:"uploads,required"`
	Description string                           `json:"description,nullable"`
	JSON        dlpDatasetResponseResultJSON     `json:"-"`
}

// dlpDatasetResponseResultJSON contains the JSON metadata for the struct
// [DlpDatasetResponseResult]
type dlpDatasetResponseResultJSON struct {
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

func (r *DlpDatasetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetResponseResultStatus string

const (
	DlpDatasetResponseResultStatusEmpty     DlpDatasetResponseResultStatus = "empty"
	DlpDatasetResponseResultStatusUploading DlpDatasetResponseResultStatus = "uploading"
	DlpDatasetResponseResultStatusFailed    DlpDatasetResponseResultStatus = "failed"
	DlpDatasetResponseResultStatusComplete  DlpDatasetResponseResultStatus = "complete"
)

type DlpDatasetResponseResultUpload struct {
	NumCells int64                                 `json:"num_cells,required"`
	Status   DlpDatasetResponseResultUploadsStatus `json:"status,required"`
	Version  int64                                 `json:"version,required"`
	JSON     dlpDatasetResponseResultUploadJSON    `json:"-"`
}

// dlpDatasetResponseResultUploadJSON contains the JSON metadata for the struct
// [DlpDatasetResponseResultUpload]
type dlpDatasetResponseResultUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetResponseResultUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetResponseResultUploadsStatus string

const (
	DlpDatasetResponseResultUploadsStatusEmpty     DlpDatasetResponseResultUploadsStatus = "empty"
	DlpDatasetResponseResultUploadsStatusUploading DlpDatasetResponseResultUploadsStatus = "uploading"
	DlpDatasetResponseResultUploadsStatusFailed    DlpDatasetResponseResultUploadsStatus = "failed"
	DlpDatasetResponseResultUploadsStatusComplete  DlpDatasetResponseResultUploadsStatus = "complete"
)

type DlpDatasetResponseResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                            `json:"total_count,required"`
	JSON       dlpDatasetResponseResultInfoJSON `json:"-"`
}

// dlpDatasetResponseResultInfoJSON contains the JSON metadata for the struct
// [DlpDatasetResponseResultInfo]
type dlpDatasetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpDatasetNewParams struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder. If
	// false, the response has no secret and the dataset is uploaded in plaintext.
	Secret param.Field[bool] `json:"secret"`
}

func (r AccountDlpDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpDatasetUpdateParams struct {
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r AccountDlpDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
