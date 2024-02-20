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

// DLPDatasetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDLPDatasetService] method instead.
type DLPDatasetService struct {
	Options []option.RequestOption
	Upload  *DLPDatasetUploadService
}

// NewDLPDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPDatasetService(opts ...option.RequestOption) (r *DLPDatasetService) {
	r = &DLPDatasetService{}
	r.Options = opts
	r.Upload = NewDLPDatasetUploadService(opts...)
	return
}

// Create a new dataset.
func (r *DLPDatasetService) New(ctx context.Context, accountID string, body DLPDatasetNewParams, opts ...option.RequestOption) (res *DLPDatasetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update details about a dataset.
func (r *DLPDatasetService) Update(ctx context.Context, accountID string, datasetID string, body DLPDatasetUpdateParams, opts ...option.RequestOption) (res *DLPDatasetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all datasets with information about available versions.
func (r *DLPDatasetService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]DLPDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a dataset.
//
// This deletes all versions of the dataset.
func (r *DLPDatasetService) Delete(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetch a specific dataset with information about available versions.
func (r *DLPDatasetService) Get(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *DLPDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPDatasetNewResponse struct {
	Dataset  DLPDatasetNewResponseDataset `json:"dataset,required"`
	MaxCells int64                        `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                    `json:"secret" format:"password"`
	JSON   dlpDatasetNewResponseJSON `json:"-"`
}

// dlpDatasetNewResponseJSON contains the JSON metadata for the struct
// [DLPDatasetNewResponse]
type dlpDatasetNewResponseJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetNewResponseDataset struct {
	ID          string                               `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                            `json:"created_at,required" format:"date-time"`
	Name        string                               `json:"name,required"`
	NumCells    int64                                `json:"num_cells,required"`
	Secret      bool                                 `json:"secret,required"`
	Status      DLPDatasetNewResponseDatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time                            `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetNewResponseDatasetUpload `json:"uploads,required"`
	Description string                               `json:"description,nullable"`
	JSON        dlpDatasetNewResponseDatasetJSON     `json:"-"`
}

// dlpDatasetNewResponseDatasetJSON contains the JSON metadata for the struct
// [DLPDatasetNewResponseDataset]
type dlpDatasetNewResponseDatasetJSON struct {
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

func (r *DLPDatasetNewResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetNewResponseDatasetStatus string

const (
	DLPDatasetNewResponseDatasetStatusEmpty     DLPDatasetNewResponseDatasetStatus = "empty"
	DLPDatasetNewResponseDatasetStatusUploading DLPDatasetNewResponseDatasetStatus = "uploading"
	DLPDatasetNewResponseDatasetStatusFailed    DLPDatasetNewResponseDatasetStatus = "failed"
	DLPDatasetNewResponseDatasetStatusComplete  DLPDatasetNewResponseDatasetStatus = "complete"
)

type DLPDatasetNewResponseDatasetUpload struct {
	NumCells int64                                     `json:"num_cells,required"`
	Status   DLPDatasetNewResponseDatasetUploadsStatus `json:"status,required"`
	Version  int64                                     `json:"version,required"`
	JSON     dlpDatasetNewResponseDatasetUploadJSON    `json:"-"`
}

// dlpDatasetNewResponseDatasetUploadJSON contains the JSON metadata for the struct
// [DLPDatasetNewResponseDatasetUpload]
type dlpDatasetNewResponseDatasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseDatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetNewResponseDatasetUploadsStatus string

const (
	DLPDatasetNewResponseDatasetUploadsStatusEmpty     DLPDatasetNewResponseDatasetUploadsStatus = "empty"
	DLPDatasetNewResponseDatasetUploadsStatusUploading DLPDatasetNewResponseDatasetUploadsStatus = "uploading"
	DLPDatasetNewResponseDatasetUploadsStatusFailed    DLPDatasetNewResponseDatasetUploadsStatus = "failed"
	DLPDatasetNewResponseDatasetUploadsStatusComplete  DLPDatasetNewResponseDatasetUploadsStatus = "complete"
)

type DLPDatasetUpdateResponse struct {
	ID          string                           `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                        `json:"created_at,required" format:"date-time"`
	Name        string                           `json:"name,required"`
	NumCells    int64                            `json:"num_cells,required"`
	Secret      bool                             `json:"secret,required"`
	Status      DLPDatasetUpdateResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                        `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetUpdateResponseUpload `json:"uploads,required"`
	Description string                           `json:"description,nullable"`
	JSON        dlpDatasetUpdateResponseJSON     `json:"-"`
}

// dlpDatasetUpdateResponseJSON contains the JSON metadata for the struct
// [DLPDatasetUpdateResponse]
type dlpDatasetUpdateResponseJSON struct {
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

func (r *DLPDatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUpdateResponseStatus string

const (
	DLPDatasetUpdateResponseStatusEmpty     DLPDatasetUpdateResponseStatus = "empty"
	DLPDatasetUpdateResponseStatusUploading DLPDatasetUpdateResponseStatus = "uploading"
	DLPDatasetUpdateResponseStatusFailed    DLPDatasetUpdateResponseStatus = "failed"
	DLPDatasetUpdateResponseStatusComplete  DLPDatasetUpdateResponseStatus = "complete"
)

type DLPDatasetUpdateResponseUpload struct {
	NumCells int64                                 `json:"num_cells,required"`
	Status   DLPDatasetUpdateResponseUploadsStatus `json:"status,required"`
	Version  int64                                 `json:"version,required"`
	JSON     dlpDatasetUpdateResponseUploadJSON    `json:"-"`
}

// dlpDatasetUpdateResponseUploadJSON contains the JSON metadata for the struct
// [DLPDatasetUpdateResponseUpload]
type dlpDatasetUpdateResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUpdateResponseUploadsStatus string

const (
	DLPDatasetUpdateResponseUploadsStatusEmpty     DLPDatasetUpdateResponseUploadsStatus = "empty"
	DLPDatasetUpdateResponseUploadsStatusUploading DLPDatasetUpdateResponseUploadsStatus = "uploading"
	DLPDatasetUpdateResponseUploadsStatusFailed    DLPDatasetUpdateResponseUploadsStatus = "failed"
	DLPDatasetUpdateResponseUploadsStatusComplete  DLPDatasetUpdateResponseUploadsStatus = "complete"
)

type DLPDatasetListResponse struct {
	ID          string                         `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                      `json:"created_at,required" format:"date-time"`
	Name        string                         `json:"name,required"`
	NumCells    int64                          `json:"num_cells,required"`
	Secret      bool                           `json:"secret,required"`
	Status      DLPDatasetListResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                      `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetListResponseUpload `json:"uploads,required"`
	Description string                         `json:"description,nullable"`
	JSON        dlpDatasetListResponseJSON     `json:"-"`
}

// dlpDatasetListResponseJSON contains the JSON metadata for the struct
// [DLPDatasetListResponse]
type dlpDatasetListResponseJSON struct {
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

func (r *DLPDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetListResponseStatus string

const (
	DLPDatasetListResponseStatusEmpty     DLPDatasetListResponseStatus = "empty"
	DLPDatasetListResponseStatusUploading DLPDatasetListResponseStatus = "uploading"
	DLPDatasetListResponseStatusFailed    DLPDatasetListResponseStatus = "failed"
	DLPDatasetListResponseStatusComplete  DLPDatasetListResponseStatus = "complete"
)

type DLPDatasetListResponseUpload struct {
	NumCells int64                               `json:"num_cells,required"`
	Status   DLPDatasetListResponseUploadsStatus `json:"status,required"`
	Version  int64                               `json:"version,required"`
	JSON     dlpDatasetListResponseUploadJSON    `json:"-"`
}

// dlpDatasetListResponseUploadJSON contains the JSON metadata for the struct
// [DLPDatasetListResponseUpload]
type dlpDatasetListResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetListResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetListResponseUploadsStatus string

const (
	DLPDatasetListResponseUploadsStatusEmpty     DLPDatasetListResponseUploadsStatus = "empty"
	DLPDatasetListResponseUploadsStatusUploading DLPDatasetListResponseUploadsStatus = "uploading"
	DLPDatasetListResponseUploadsStatusFailed    DLPDatasetListResponseUploadsStatus = "failed"
	DLPDatasetListResponseUploadsStatusComplete  DLPDatasetListResponseUploadsStatus = "complete"
)

type DLPDatasetGetResponse struct {
	ID          string                        `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                     `json:"created_at,required" format:"date-time"`
	Name        string                        `json:"name,required"`
	NumCells    int64                         `json:"num_cells,required"`
	Secret      bool                          `json:"secret,required"`
	Status      DLPDatasetGetResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                     `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetGetResponseUpload `json:"uploads,required"`
	Description string                        `json:"description,nullable"`
	JSON        dlpDatasetGetResponseJSON     `json:"-"`
}

// dlpDatasetGetResponseJSON contains the JSON metadata for the struct
// [DLPDatasetGetResponse]
type dlpDatasetGetResponseJSON struct {
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

func (r *DLPDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetGetResponseStatus string

const (
	DLPDatasetGetResponseStatusEmpty     DLPDatasetGetResponseStatus = "empty"
	DLPDatasetGetResponseStatusUploading DLPDatasetGetResponseStatus = "uploading"
	DLPDatasetGetResponseStatusFailed    DLPDatasetGetResponseStatus = "failed"
	DLPDatasetGetResponseStatusComplete  DLPDatasetGetResponseStatus = "complete"
)

type DLPDatasetGetResponseUpload struct {
	NumCells int64                              `json:"num_cells,required"`
	Status   DLPDatasetGetResponseUploadsStatus `json:"status,required"`
	Version  int64                              `json:"version,required"`
	JSON     dlpDatasetGetResponseUploadJSON    `json:"-"`
}

// dlpDatasetGetResponseUploadJSON contains the JSON metadata for the struct
// [DLPDatasetGetResponseUpload]
type dlpDatasetGetResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetGetResponseUploadsStatus string

const (
	DLPDatasetGetResponseUploadsStatusEmpty     DLPDatasetGetResponseUploadsStatus = "empty"
	DLPDatasetGetResponseUploadsStatusUploading DLPDatasetGetResponseUploadsStatus = "uploading"
	DLPDatasetGetResponseUploadsStatusFailed    DLPDatasetGetResponseUploadsStatus = "failed"
	DLPDatasetGetResponseUploadsStatusComplete  DLPDatasetGetResponseUploadsStatus = "complete"
)

type DLPDatasetNewParams struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder. If
	// false, the response has no secret and the dataset is uploaded in plaintext.
	Secret param.Field[bool] `json:"secret"`
}

func (r DLPDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDatasetNewResponseEnvelope struct {
	Errors     []DLPDatasetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetNewResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     DLPDatasetNewResponse                   `json:"result"`
	ResultInfo DLPDatasetNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetNewResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetNewResponseEnvelope]
type dlpDatasetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpDatasetNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetNewResponseEnvelopeErrors]
type dlpDatasetNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dlpDatasetNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetNewResponseEnvelopeMessages]
type dlpDatasetNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetNewResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                       `json:"total_count,required"`
	JSON       dlpDatasetNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DLPDatasetNewResponseEnvelopeResultInfo]
type dlpDatasetNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUpdateParams struct {
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DLPDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDatasetUpdateResponseEnvelope struct {
	Errors     []DLPDatasetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetUpdateResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                       `json:"success,required"`
	Result     DLPDatasetUpdateResponse                   `json:"result"`
	ResultInfo DLPDatasetUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUpdateResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetUpdateResponseEnvelope]
type dlpDatasetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpDatasetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetUpdateResponseEnvelopeErrors]
type dlpDatasetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    dlpDatasetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetUpdateResponseEnvelopeMessages]
type dlpDatasetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUpdateResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                          `json:"total_count,required"`
	JSON       dlpDatasetUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DLPDatasetUpdateResponseEnvelopeResultInfo]
type dlpDatasetUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetListResponseEnvelope struct {
	Errors     []DLPDatasetListResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetListResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                     `json:"success,required"`
	Result     []DLPDatasetListResponse                 `json:"result"`
	ResultInfo DLPDatasetListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetListResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetListResponseEnvelope]
type dlpDatasetListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dlpDatasetListResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetListResponseEnvelopeErrors]
type dlpDatasetListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpDatasetListResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetListResponseEnvelopeMessages]
type dlpDatasetListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetListResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                        `json:"total_count,required"`
	JSON       dlpDatasetListResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DLPDatasetListResponseEnvelopeResultInfo]
type dlpDatasetListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetGetResponseEnvelope struct {
	Errors     []DLPDatasetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetGetResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     DLPDatasetGetResponse                   `json:"result"`
	ResultInfo DLPDatasetGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetGetResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetGetResponseEnvelope]
type dlpDatasetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpDatasetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetGetResponseEnvelopeErrors]
type dlpDatasetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dlpDatasetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetGetResponseEnvelopeMessages]
type dlpDatasetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetGetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                       `json:"total_count,required"`
	JSON       dlpDatasetGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DLPDatasetGetResponseEnvelopeResultInfo]
type dlpDatasetGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
