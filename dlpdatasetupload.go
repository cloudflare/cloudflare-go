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
func (r *DLPDatasetUploadService) Update(ctx context.Context, accountID string, datasetID string, version int64, opts ...option.RequestOption) (res *DLPDatasetUploadUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUploadUpdateResponseEnvelope
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

type DLPDatasetUploadUpdateResponse struct {
	ID          string                                 `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                              `json:"created_at,required" format:"date-time"`
	Name        string                                 `json:"name,required"`
	NumCells    int64                                  `json:"num_cells,required"`
	Secret      bool                                   `json:"secret,required"`
	Status      DLPDatasetUploadUpdateResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                              `json:"updated_at,required" format:"date-time"`
	Uploads     []DLPDatasetUploadUpdateResponseUpload `json:"uploads,required"`
	Description string                                 `json:"description,nullable"`
	JSON        dlpDatasetUploadUpdateResponseJSON     `json:"-"`
}

// dlpDatasetUploadUpdateResponseJSON contains the JSON metadata for the struct
// [DLPDatasetUploadUpdateResponse]
type dlpDatasetUploadUpdateResponseJSON struct {
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

func (r *DLPDatasetUploadUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadUpdateResponseStatus string

const (
	DLPDatasetUploadUpdateResponseStatusEmpty     DLPDatasetUploadUpdateResponseStatus = "empty"
	DLPDatasetUploadUpdateResponseStatusUploading DLPDatasetUploadUpdateResponseStatus = "uploading"
	DLPDatasetUploadUpdateResponseStatusFailed    DLPDatasetUploadUpdateResponseStatus = "failed"
	DLPDatasetUploadUpdateResponseStatusComplete  DLPDatasetUploadUpdateResponseStatus = "complete"
)

type DLPDatasetUploadUpdateResponseUpload struct {
	NumCells int64                                       `json:"num_cells,required"`
	Status   DLPDatasetUploadUpdateResponseUploadsStatus `json:"status,required"`
	Version  int64                                       `json:"version,required"`
	JSON     dlpDatasetUploadUpdateResponseUploadJSON    `json:"-"`
}

// dlpDatasetUploadUpdateResponseUploadJSON contains the JSON metadata for the
// struct [DLPDatasetUploadUpdateResponseUpload]
type dlpDatasetUploadUpdateResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadUpdateResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadUpdateResponseUploadsStatus string

const (
	DLPDatasetUploadUpdateResponseUploadsStatusEmpty     DLPDatasetUploadUpdateResponseUploadsStatus = "empty"
	DLPDatasetUploadUpdateResponseUploadsStatusUploading DLPDatasetUploadUpdateResponseUploadsStatus = "uploading"
	DLPDatasetUploadUpdateResponseUploadsStatusFailed    DLPDatasetUploadUpdateResponseUploadsStatus = "failed"
	DLPDatasetUploadUpdateResponseUploadsStatusComplete  DLPDatasetUploadUpdateResponseUploadsStatus = "complete"
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

type DLPDatasetUploadUpdateResponseEnvelope struct {
	Errors     []DLPDatasetUploadUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetUploadUpdateResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                             `json:"success,required"`
	Result     DLPDatasetUploadUpdateResponse                   `json:"result"`
	ResultInfo DLPDatasetUploadUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadUpdateResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadUpdateResponseEnvelope]
type dlpDatasetUploadUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dlpDatasetUploadUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpDatasetUploadUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPDatasetUploadUpdateResponseEnvelopeErrors]
type dlpDatasetUploadUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    dlpDatasetUploadUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpDatasetUploadUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPDatasetUploadUpdateResponseEnvelopeMessages]
type dlpDatasetUploadUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPDatasetUploadUpdateResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                `json:"total_count,required"`
	JSON       dlpDatasetUploadUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DLPDatasetUploadUpdateResponseEnvelopeResultInfo]
type dlpDatasetUploadUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
