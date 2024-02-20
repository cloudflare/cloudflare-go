// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DLPDatasetUploadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDLPDatasetUploadService] method
// instead.
type DLPDatasetUploadService struct {
	Options []option.RequestOption
}

// NewDLPDatasetUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPDatasetUploadService(opts ...option.RequestOption) (r *DLPDatasetUploadService) {
	r = &DLPDatasetUploadService{}
	r.Options = opts
	return
}

// Prepare to upload a new version of a dataset.
func (r *DLPDatasetUploadService) New(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *DLPDatasetUploadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUploadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upload a new version of a dataset.
func (r *DLPDatasetUploadService) Edit(ctx context.Context, accountID string, datasetID string, version int64, opts ...option.RequestOption) (res *DLPDatasetUploadEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUploadEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", accountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPDatasetUploadNewResponse struct {
	MaxCells int64                           `json:"max_cells,required"`
	Version  int64                           `json:"version,required"`
	Secret   string                          `json:"secret" format:"password"`
	JSON     dlpDatasetUploadNewResponseJSON `json:"-"`
}

// dlpDatasetUploadNewResponseJSON contains the JSON metadata for the struct
// [DLPDatasetUploadNewResponse]
type dlpDatasetUploadNewResponseJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponse struct {
	ID          string                               `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                            `json:"created_at,required" format:"date-time"`
	Name        string                               `json:"name,required"`
	NumCells    int64                                `json:"num_cells,required"`
	Secret      bool                                 `json:"secret,required"`
	Status      DLPDatasetUploadEditResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                            `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetUploadEditResponseUpload `json:"uploads,required"`
	Description string                               `json:"description,nullable"`
	JSON        dlpDatasetUploadEditResponseJSON     `json:"-"`
}

// dlpDatasetUploadEditResponseJSON contains the JSON metadata for the struct
// [DLPDatasetUploadEditResponse]
type dlpDatasetUploadEditResponseJSON struct {
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

func (r *DLPDatasetUploadEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponseStatus string

const (
	DLPDatasetUploadEditResponseStatusEmpty     DLPDatasetUploadEditResponseStatus = "empty"
	DLPDatasetUploadEditResponseStatusUploading DLPDatasetUploadEditResponseStatus = "uploading"
	DLPDatasetUploadEditResponseStatusFailed    DLPDatasetUploadEditResponseStatus = "failed"
	DLPDatasetUploadEditResponseStatusComplete  DLPDatasetUploadEditResponseStatus = "complete"
)

type DLPDatasetUploadEditResponseUpload struct {
	NumCells int64                                     `json:"num_cells,required"`
	Status   DLPDatasetUploadEditResponseUploadsStatus `json:"status,required"`
	Version  int64                                     `json:"version,required"`
	JSON     dlpDatasetUploadEditResponseUploadJSON    `json:"-"`
}

// dlpDatasetUploadEditResponseUploadJSON contains the JSON metadata for the struct
// [DLPDatasetUploadEditResponseUpload]
type dlpDatasetUploadEditResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponseUploadsStatus string

const (
	DLPDatasetUploadEditResponseUploadsStatusEmpty     DLPDatasetUploadEditResponseUploadsStatus = "empty"
	DLPDatasetUploadEditResponseUploadsStatusUploading DLPDatasetUploadEditResponseUploadsStatus = "uploading"
	DLPDatasetUploadEditResponseUploadsStatusFailed    DLPDatasetUploadEditResponseUploadsStatus = "failed"
	DLPDatasetUploadEditResponseUploadsStatusComplete  DLPDatasetUploadEditResponseUploadsStatus = "complete"
)

type DLPDatasetUploadNewResponseEnvelope struct {
	Errors     []DLPDatasetUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetUploadNewResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                          `json:"success,required"`
	Result     DLPDatasetUploadNewResponse                   `json:"result"`
	ResultInfo DLPDatasetUploadNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadNewResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelope]
type dlpDatasetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dlpDatasetUploadNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelopeErrors]
type dlpDatasetUploadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dlpDatasetUploadNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDatasetUploadNewResponseEnvelopeMessages]
type dlpDatasetUploadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadNewResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                             `json:"total_count,required"`
	JSON       dlpDatasetUploadNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DLPDatasetUploadNewResponseEnvelopeResultInfo]
type dlpDatasetUploadNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponseEnvelope struct {
	Errors     []DLPDatasetUploadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetUploadEditResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                           `json:"success,required"`
	Result     DLPDatasetUploadEditResponse                   `json:"result"`
	ResultInfo DLPDatasetUploadEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadEditResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadEditResponseEnvelope]
type dlpDatasetUploadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dlpDatasetUploadEditResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPDatasetUploadEditResponseEnvelopeErrors]
type dlpDatasetUploadEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dlpDatasetUploadEditResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDatasetUploadEditResponseEnvelopeMessages]
type dlpDatasetUploadEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadEditResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                              `json:"total_count,required"`
	JSON       dlpDatasetUploadEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DLPDatasetUploadEditResponseEnvelopeResultInfo]
type dlpDatasetUploadEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
