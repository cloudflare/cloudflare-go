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
	var env DlpDatasetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a specific dataset with information about available versions.
func (r *DlpDatasetService) Get(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *DlpDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DlpDatasetGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update details about a dataset.
func (r *DlpDatasetService) Update(ctx context.Context, accountID string, datasetID string, body DlpDatasetUpdateParams, opts ...option.RequestOption) (res *DlpDatasetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DlpDatasetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all datasets with information about available versions.
func (r *DlpDatasetService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]DlpDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DlpDatasetListResponseEnvelope
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
	var env DlpDatasetUploadResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", accountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Prepare to upload a new version of a dataset.
func (r *DlpDatasetService) UploadPrepare(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *DlpDatasetUploadPrepareResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DlpDatasetUploadPrepareResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DlpDatasetNewResponse struct {
	Dataset  DlpDatasetNewResponseDataset `json:"dataset,required"`
	MaxCells int64                        `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                    `json:"secret" format:"password"`
	JSON   dlpDatasetNewResponseJSON `json:"-"`
}

// dlpDatasetNewResponseJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponse]
type dlpDatasetNewResponseJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseDataset struct {
	ID          string                               `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                            `json:"created_at,required" format:"date-time"`
	Name        string                               `json:"name,required"`
	NumCells    int64                                `json:"num_cells,required"`
	Secret      bool                                 `json:"secret,required"`
	Status      DlpDatasetNewResponseDatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time                            `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetNewResponseDatasetUpload `json:"uploads,required"`
	Description string                               `json:"description,nullable"`
	JSON        dlpDatasetNewResponseDatasetJSON     `json:"-"`
}

// dlpDatasetNewResponseDatasetJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseDataset]
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

func (r *DlpDatasetNewResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseDatasetStatus string

const (
	DlpDatasetNewResponseDatasetStatusEmpty     DlpDatasetNewResponseDatasetStatus = "empty"
	DlpDatasetNewResponseDatasetStatusUploading DlpDatasetNewResponseDatasetStatus = "uploading"
	DlpDatasetNewResponseDatasetStatusFailed    DlpDatasetNewResponseDatasetStatus = "failed"
	DlpDatasetNewResponseDatasetStatusComplete  DlpDatasetNewResponseDatasetStatus = "complete"
)

type DlpDatasetNewResponseDatasetUpload struct {
	NumCells int64                                     `json:"num_cells,required"`
	Status   DlpDatasetNewResponseDatasetUploadsStatus `json:"status,required"`
	Version  int64                                     `json:"version,required"`
	JSON     dlpDatasetNewResponseDatasetUploadJSON    `json:"-"`
}

// dlpDatasetNewResponseDatasetUploadJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseDatasetUpload]
type dlpDatasetNewResponseDatasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseDatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseDatasetUploadsStatus string

const (
	DlpDatasetNewResponseDatasetUploadsStatusEmpty     DlpDatasetNewResponseDatasetUploadsStatus = "empty"
	DlpDatasetNewResponseDatasetUploadsStatusUploading DlpDatasetNewResponseDatasetUploadsStatus = "uploading"
	DlpDatasetNewResponseDatasetUploadsStatusFailed    DlpDatasetNewResponseDatasetUploadsStatus = "failed"
	DlpDatasetNewResponseDatasetUploadsStatusComplete  DlpDatasetNewResponseDatasetUploadsStatus = "complete"
)

type DlpDatasetGetResponse struct {
	ID          string                        `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                     `json:"created_at,required" format:"date-time"`
	Name        string                        `json:"name,required"`
	NumCells    int64                         `json:"num_cells,required"`
	Secret      bool                          `json:"secret,required"`
	Status      DlpDatasetGetResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                     `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetGetResponseUpload `json:"uploads,required"`
	Description string                        `json:"description,nullable"`
	JSON        dlpDatasetGetResponseJSON     `json:"-"`
}

// dlpDatasetGetResponseJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponse]
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

func (r *DlpDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseStatus string

const (
	DlpDatasetGetResponseStatusEmpty     DlpDatasetGetResponseStatus = "empty"
	DlpDatasetGetResponseStatusUploading DlpDatasetGetResponseStatus = "uploading"
	DlpDatasetGetResponseStatusFailed    DlpDatasetGetResponseStatus = "failed"
	DlpDatasetGetResponseStatusComplete  DlpDatasetGetResponseStatus = "complete"
)

type DlpDatasetGetResponseUpload struct {
	NumCells int64                              `json:"num_cells,required"`
	Status   DlpDatasetGetResponseUploadsStatus `json:"status,required"`
	Version  int64                              `json:"version,required"`
	JSON     dlpDatasetGetResponseUploadJSON    `json:"-"`
}

// dlpDatasetGetResponseUploadJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseUpload]
type dlpDatasetGetResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseUploadsStatus string

const (
	DlpDatasetGetResponseUploadsStatusEmpty     DlpDatasetGetResponseUploadsStatus = "empty"
	DlpDatasetGetResponseUploadsStatusUploading DlpDatasetGetResponseUploadsStatus = "uploading"
	DlpDatasetGetResponseUploadsStatusFailed    DlpDatasetGetResponseUploadsStatus = "failed"
	DlpDatasetGetResponseUploadsStatusComplete  DlpDatasetGetResponseUploadsStatus = "complete"
)

type DlpDatasetUpdateResponse struct {
	ID          string                           `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                        `json:"created_at,required" format:"date-time"`
	Name        string                           `json:"name,required"`
	NumCells    int64                            `json:"num_cells,required"`
	Secret      bool                             `json:"secret,required"`
	Status      DlpDatasetUpdateResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                        `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetUpdateResponseUpload `json:"uploads,required"`
	Description string                           `json:"description,nullable"`
	JSON        dlpDatasetUpdateResponseJSON     `json:"-"`
}

// dlpDatasetUpdateResponseJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponse]
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

func (r *DlpDatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseStatus string

const (
	DlpDatasetUpdateResponseStatusEmpty     DlpDatasetUpdateResponseStatus = "empty"
	DlpDatasetUpdateResponseStatusUploading DlpDatasetUpdateResponseStatus = "uploading"
	DlpDatasetUpdateResponseStatusFailed    DlpDatasetUpdateResponseStatus = "failed"
	DlpDatasetUpdateResponseStatusComplete  DlpDatasetUpdateResponseStatus = "complete"
)

type DlpDatasetUpdateResponseUpload struct {
	NumCells int64                                 `json:"num_cells,required"`
	Status   DlpDatasetUpdateResponseUploadsStatus `json:"status,required"`
	Version  int64                                 `json:"version,required"`
	JSON     dlpDatasetUpdateResponseUploadJSON    `json:"-"`
}

// dlpDatasetUpdateResponseUploadJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponseUpload]
type dlpDatasetUpdateResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseUploadsStatus string

const (
	DlpDatasetUpdateResponseUploadsStatusEmpty     DlpDatasetUpdateResponseUploadsStatus = "empty"
	DlpDatasetUpdateResponseUploadsStatusUploading DlpDatasetUpdateResponseUploadsStatus = "uploading"
	DlpDatasetUpdateResponseUploadsStatusFailed    DlpDatasetUpdateResponseUploadsStatus = "failed"
	DlpDatasetUpdateResponseUploadsStatusComplete  DlpDatasetUpdateResponseUploadsStatus = "complete"
)

type DlpDatasetListResponse struct {
	ID          string                         `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                      `json:"created_at,required" format:"date-time"`
	Name        string                         `json:"name,required"`
	NumCells    int64                          `json:"num_cells,required"`
	Secret      bool                           `json:"secret,required"`
	Status      DlpDatasetListResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                      `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetListResponseUpload `json:"uploads,required"`
	Description string                         `json:"description,nullable"`
	JSON        dlpDatasetListResponseJSON     `json:"-"`
}

// dlpDatasetListResponseJSON contains the JSON metadata for the struct
// [DlpDatasetListResponse]
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

func (r *DlpDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseStatus string

const (
	DlpDatasetListResponseStatusEmpty     DlpDatasetListResponseStatus = "empty"
	DlpDatasetListResponseStatusUploading DlpDatasetListResponseStatus = "uploading"
	DlpDatasetListResponseStatusFailed    DlpDatasetListResponseStatus = "failed"
	DlpDatasetListResponseStatusComplete  DlpDatasetListResponseStatus = "complete"
)

type DlpDatasetListResponseUpload struct {
	NumCells int64                               `json:"num_cells,required"`
	Status   DlpDatasetListResponseUploadsStatus `json:"status,required"`
	Version  int64                               `json:"version,required"`
	JSON     dlpDatasetListResponseUploadJSON    `json:"-"`
}

// dlpDatasetListResponseUploadJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseUpload]
type dlpDatasetListResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseUploadsStatus string

const (
	DlpDatasetListResponseUploadsStatusEmpty     DlpDatasetListResponseUploadsStatus = "empty"
	DlpDatasetListResponseUploadsStatusUploading DlpDatasetListResponseUploadsStatus = "uploading"
	DlpDatasetListResponseUploadsStatusFailed    DlpDatasetListResponseUploadsStatus = "failed"
	DlpDatasetListResponseUploadsStatusComplete  DlpDatasetListResponseUploadsStatus = "complete"
)

type DlpDatasetUploadResponse struct {
	ID          string                           `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                        `json:"created_at,required" format:"date-time"`
	Name        string                           `json:"name,required"`
	NumCells    int64                            `json:"num_cells,required"`
	Secret      bool                             `json:"secret,required"`
	Status      DlpDatasetUploadResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                        `json:"updated_at,required" format:"date-time"`
	Uploads     []DlpDatasetUploadResponseUpload `json:"uploads,required"`
	Description string                           `json:"description,nullable"`
	JSON        dlpDatasetUploadResponseJSON     `json:"-"`
}

// dlpDatasetUploadResponseJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponse]
type dlpDatasetUploadResponseJSON struct {
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

func (r *DlpDatasetUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseStatus string

const (
	DlpDatasetUploadResponseStatusEmpty     DlpDatasetUploadResponseStatus = "empty"
	DlpDatasetUploadResponseStatusUploading DlpDatasetUploadResponseStatus = "uploading"
	DlpDatasetUploadResponseStatusFailed    DlpDatasetUploadResponseStatus = "failed"
	DlpDatasetUploadResponseStatusComplete  DlpDatasetUploadResponseStatus = "complete"
)

type DlpDatasetUploadResponseUpload struct {
	NumCells int64                                 `json:"num_cells,required"`
	Status   DlpDatasetUploadResponseUploadsStatus `json:"status,required"`
	Version  int64                                 `json:"version,required"`
	JSON     dlpDatasetUploadResponseUploadJSON    `json:"-"`
}

// dlpDatasetUploadResponseUploadJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponseUpload]
type dlpDatasetUploadResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseUploadsStatus string

const (
	DlpDatasetUploadResponseUploadsStatusEmpty     DlpDatasetUploadResponseUploadsStatus = "empty"
	DlpDatasetUploadResponseUploadsStatusUploading DlpDatasetUploadResponseUploadsStatus = "uploading"
	DlpDatasetUploadResponseUploadsStatusFailed    DlpDatasetUploadResponseUploadsStatus = "failed"
	DlpDatasetUploadResponseUploadsStatusComplete  DlpDatasetUploadResponseUploadsStatus = "complete"
)

type DlpDatasetUploadPrepareResponse struct {
	MaxCells int64                               `json:"max_cells,required"`
	Version  int64                               `json:"version,required"`
	Secret   string                              `json:"secret" format:"password"`
	JSON     dlpDatasetUploadPrepareResponseJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseJSON contains the JSON metadata for the struct
// [DlpDatasetUploadPrepareResponse]
type dlpDatasetUploadPrepareResponseJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponse) UnmarshalJSON(data []byte) (err error) {
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

type DlpDatasetNewResponseEnvelope struct {
	Errors     []DlpDatasetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DlpDatasetNewResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     DlpDatasetNewResponse                   `json:"result"`
	ResultInfo DlpDatasetNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetNewResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DlpDatasetNewResponseEnvelope]
type dlpDatasetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpDatasetNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DlpDatasetNewResponseEnvelopeErrors]
type dlpDatasetNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dlpDatasetNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DlpDatasetNewResponseEnvelopeMessages]
type dlpDatasetNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetNewResponseEnvelopeResultInfo struct {
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
// struct [DlpDatasetNewResponseEnvelopeResultInfo]
type dlpDatasetNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseEnvelope struct {
	Errors     []DlpDatasetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DlpDatasetGetResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     DlpDatasetGetResponse                   `json:"result"`
	ResultInfo DlpDatasetGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetGetResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DlpDatasetGetResponseEnvelope]
type dlpDatasetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpDatasetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DlpDatasetGetResponseEnvelopeErrors]
type dlpDatasetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dlpDatasetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DlpDatasetGetResponseEnvelopeMessages]
type dlpDatasetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetGetResponseEnvelopeResultInfo struct {
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
// struct [DlpDatasetGetResponseEnvelopeResultInfo]
type dlpDatasetGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateParams struct {
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DlpDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DlpDatasetUpdateResponseEnvelope struct {
	Errors     []DlpDatasetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DlpDatasetUpdateResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                       `json:"success,required"`
	Result     DlpDatasetUpdateResponse                   `json:"result"`
	ResultInfo DlpDatasetUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUpdateResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DlpDatasetUpdateResponseEnvelope]
type dlpDatasetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpDatasetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DlpDatasetUpdateResponseEnvelopeErrors]
type dlpDatasetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    dlpDatasetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DlpDatasetUpdateResponseEnvelopeMessages]
type dlpDatasetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUpdateResponseEnvelopeResultInfo struct {
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
// the struct [DlpDatasetUpdateResponseEnvelopeResultInfo]
type dlpDatasetUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseEnvelope struct {
	Errors     []DlpDatasetListResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DlpDatasetListResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                     `json:"success,required"`
	Result     []DlpDatasetListResponse                 `json:"result"`
	ResultInfo DlpDatasetListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetListResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DlpDatasetListResponseEnvelope]
type dlpDatasetListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dlpDatasetListResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DlpDatasetListResponseEnvelopeErrors]
type dlpDatasetListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpDatasetListResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DlpDatasetListResponseEnvelopeMessages]
type dlpDatasetListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetListResponseEnvelopeResultInfo struct {
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
// struct [DlpDatasetListResponseEnvelopeResultInfo]
type dlpDatasetListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseEnvelope struct {
	Errors     []DlpDatasetUploadResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DlpDatasetUploadResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                       `json:"success,required"`
	Result     DlpDatasetUploadResponse                   `json:"result"`
	ResultInfo DlpDatasetUploadResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadResponseEnvelopeJSON contains the JSON metadata for the struct
// [DlpDatasetUploadResponseEnvelope]
type dlpDatasetUploadResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpDatasetUploadResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUploadResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DlpDatasetUploadResponseEnvelopeErrors]
type dlpDatasetUploadResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    dlpDatasetUploadResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUploadResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DlpDatasetUploadResponseEnvelopeMessages]
type dlpDatasetUploadResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                          `json:"total_count,required"`
	JSON       dlpDatasetUploadResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DlpDatasetUploadResponseEnvelopeResultInfo]
type dlpDatasetUploadResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseEnvelope struct {
	Errors     []DlpDatasetUploadPrepareResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DlpDatasetUploadPrepareResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                              `json:"success,required"`
	Result     DlpDatasetUploadPrepareResponse                   `json:"result"`
	ResultInfo DlpDatasetUploadPrepareResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadPrepareResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadPrepareResponseEnvelopeJSON contains the JSON metadata for the
// struct [DlpDatasetUploadPrepareResponseEnvelope]
type dlpDatasetUploadPrepareResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dlpDatasetUploadPrepareResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DlpDatasetUploadPrepareResponseEnvelopeErrors]
type dlpDatasetUploadPrepareResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dlpDatasetUploadPrepareResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DlpDatasetUploadPrepareResponseEnvelopeMessages]
type dlpDatasetUploadPrepareResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DlpDatasetUploadPrepareResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                 `json:"total_count,required"`
	JSON       dlpDatasetUploadPrepareResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadPrepareResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DlpDatasetUploadPrepareResponseEnvelopeResultInfo]
type dlpDatasetUploadPrepareResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpDatasetUploadPrepareResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
