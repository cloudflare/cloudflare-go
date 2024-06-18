// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DLPDatasetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetService] method instead.
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
func (r *DLPDatasetService) New(ctx context.Context, params DLPDatasetNewParams, opts ...option.RequestOption) (res *DatasetCreation, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetNewResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update details about a dataset.
func (r *DLPDatasetService) Update(ctx context.Context, datasetID string, params DLPDatasetUpdateParams, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUpdateResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", params.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all datasets with information about available versions.
func (r *DLPDatasetService) List(ctx context.Context, query DLPDatasetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Dataset], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetch all datasets with information about available versions.
func (r *DLPDatasetService) ListAutoPaging(ctx context.Context, query DLPDatasetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Dataset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a dataset.
//
// This deletes all versions of the dataset.
func (r *DLPDatasetService) Delete(ctx context.Context, datasetID string, body DLPDatasetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetch a specific dataset with information about available versions.
func (r *DLPDatasetService) Get(ctx context.Context, datasetID string, query DLPDatasetGetParams, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Dataset struct {
	ID          string          `json:"id,required" format:"uuid"`
	CreatedAt   time.Time       `json:"created_at,required" format:"date-time"`
	Name        string          `json:"name,required"`
	NumCells    int64           `json:"num_cells,required"`
	Secret      bool            `json:"secret,required"`
	Status      DatasetStatus   `json:"status,required"`
	UpdatedAt   time.Time       `json:"updated_at,required" format:"date-time"`
	Uploads     []DatasetUpload `json:"uploads,required"`
	Description string          `json:"description,nullable"`
	JSON        datasetJSON     `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
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

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

type DatasetStatus string

const (
	DatasetStatusEmpty     DatasetStatus = "empty"
	DatasetStatusUploading DatasetStatus = "uploading"
	DatasetStatusFailed    DatasetStatus = "failed"
	DatasetStatusComplete  DatasetStatus = "complete"
)

func (r DatasetStatus) IsKnown() bool {
	switch r {
	case DatasetStatusEmpty, DatasetStatusUploading, DatasetStatusFailed, DatasetStatusComplete:
		return true
	}
	return false
}

type DatasetUpload struct {
	NumCells int64                `json:"num_cells,required"`
	Status   DatasetUploadsStatus `json:"status,required"`
	Version  int64                `json:"version,required"`
	JSON     datasetUploadJSON    `json:"-"`
}

// datasetUploadJSON contains the JSON metadata for the struct [DatasetUpload]
type datasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetUploadJSON) RawJSON() string {
	return r.raw
}

type DatasetUploadsStatus string

const (
	DatasetUploadsStatusEmpty     DatasetUploadsStatus = "empty"
	DatasetUploadsStatusUploading DatasetUploadsStatus = "uploading"
	DatasetUploadsStatusFailed    DatasetUploadsStatus = "failed"
	DatasetUploadsStatusComplete  DatasetUploadsStatus = "complete"
)

func (r DatasetUploadsStatus) IsKnown() bool {
	switch r {
	case DatasetUploadsStatusEmpty, DatasetUploadsStatusUploading, DatasetUploadsStatusFailed, DatasetUploadsStatusComplete:
		return true
	}
	return false
}

type DatasetArray []Dataset

type DatasetCreation struct {
	Dataset  Dataset `json:"dataset,required"`
	MaxCells int64   `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string              `json:"secret" format:"password"`
	JSON   datasetCreationJSON `json:"-"`
}

// datasetCreationJSON contains the JSON metadata for the struct [DatasetCreation]
type datasetCreationJSON struct {
	Dataset     apijson.Field
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetCreationJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetNewParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
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
	Errors     []shared.ResponseInfo                   `json:"errors,required"`
	Messages   []shared.ResponseInfo                   `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     DatasetCreation                         `json:"result"`
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

func (r dlpDatasetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DLPDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDatasetUpdateResponseEnvelope struct {
	Errors     []shared.ResponseInfo                      `json:"errors,required"`
	Messages   []shared.ResponseInfo                      `json:"messages,required"`
	Success    bool                                       `json:"success,required"`
	Result     Dataset                                    `json:"result"`
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

func (r dlpDatasetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetGetResponseEnvelope struct {
	Errors     []shared.ResponseInfo                   `json:"errors,required"`
	Messages   []shared.ResponseInfo                   `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     Dataset                                 `json:"result"`
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

func (r dlpDatasetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
