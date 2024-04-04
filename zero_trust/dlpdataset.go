// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *DLPDatasetService) New(ctx context.Context, params DLPDatasetNewParams, opts ...option.RequestOption) (res *DLPDatasetCreation, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update details about a dataset.
func (r *DLPDatasetService) Update(ctx context.Context, datasetID string, params DLPDatasetUpdateParams, opts ...option.RequestOption) (res *DLPDataset, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", params.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all datasets with information about available versions.
func (r *DLPDatasetService) List(ctx context.Context, query DLPDatasetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPDataset], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *DLPDatasetService) ListAutoPaging(ctx context.Context, query DLPDatasetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPDataset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a dataset.
//
// This deletes all versions of the dataset.
func (r *DLPDatasetService) Delete(ctx context.Context, datasetID string, body DLPDatasetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetch a specific dataset with information about available versions.
func (r *DLPDatasetService) Get(ctx context.Context, datasetID string, query DLPDatasetGetParams, opts ...option.RequestOption) (res *DLPDataset, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r dlpDatasetJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetStatus string

const (
	DLPDatasetStatusEmpty     DLPDatasetStatus = "empty"
	DLPDatasetStatusUploading DLPDatasetStatus = "uploading"
	DLPDatasetStatusFailed    DLPDatasetStatus = "failed"
	DLPDatasetStatusComplete  DLPDatasetStatus = "complete"
)

func (r DLPDatasetStatus) IsKnown() bool {
	switch r {
	case DLPDatasetStatusEmpty, DLPDatasetStatusUploading, DLPDatasetStatusFailed, DLPDatasetStatusComplete:
		return true
	}
	return false
}

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

func (r dlpDatasetUploadJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadsStatus string

const (
	DLPDatasetUploadsStatusEmpty     DLPDatasetUploadsStatus = "empty"
	DLPDatasetUploadsStatusUploading DLPDatasetUploadsStatus = "uploading"
	DLPDatasetUploadsStatusFailed    DLPDatasetUploadsStatus = "failed"
	DLPDatasetUploadsStatusComplete  DLPDatasetUploadsStatus = "complete"
)

func (r DLPDatasetUploadsStatus) IsKnown() bool {
	switch r {
	case DLPDatasetUploadsStatusEmpty, DLPDatasetUploadsStatusUploading, DLPDatasetUploadsStatusFailed, DLPDatasetUploadsStatusComplete:
		return true
	}
	return false
}

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

func (r dlpDatasetCreationJSON) RawJSON() string {
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
	Errors     []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Success    bool                                                      `json:"success,required"`
	Result     DLPDatasetCreation                                        `json:"result"`
	ResultInfo DLPDatasetNewResponseEnvelopeResultInfo                   `json:"result_info"`
	JSON       dlpDatasetNewResponseEnvelopeJSON                         `json:"-"`
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
	Errors     []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Success    bool                                                      `json:"success,required"`
	Result     DLPDataset                                                `json:"result"`
	ResultInfo DLPDatasetUpdateResponseEnvelopeResultInfo                `json:"result_info"`
	JSON       dlpDatasetUpdateResponseEnvelopeJSON                      `json:"-"`
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
	Errors     []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Success    bool                                                      `json:"success,required"`
	Result     DLPDataset                                                `json:"result"`
	ResultInfo DLPDatasetGetResponseEnvelopeResultInfo                   `json:"result_info"`
	JSON       dlpDatasetGetResponseEnvelopeJSON                         `json:"-"`
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
