// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DLPDatasetUploadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetUploadService] method instead.
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

// Prepare to upload a new version of a dataset
func (r *DLPDatasetUploadService) New(ctx context.Context, datasetID string, body DLPDatasetUploadNewParams, opts ...option.RequestOption) (res *DLPDatasetUploadNewResponse, err error) {
	var env DLPDatasetUploadNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is used for single-column EDMv1 and Custom Word Lists. The EDM format can
// only be created in the Cloudflare dashboard. For other clients, this operation
// can only be used for non-secret Custom Word Lists. The body must be a UTF-8
// encoded, newline (NL or CRNL) separated list of words to be matched.
func (r *DLPDatasetUploadService) Edit(ctx context.Context, datasetID string, version int64, dataset io.Reader, body DLPDatasetUploadEditParams, opts ...option.RequestOption) (res *DLPDatasetUploadEditResponse, err error) {
	var env DLPDatasetUploadEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", dataset)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", body.AccountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPDatasetUploadNewResponse struct {
	EncodingVersion int64                               `json:"encoding_version,required"`
	MaxCells        int64                               `json:"max_cells,required"`
	Version         int64                               `json:"version,required"`
	CaseSensitive   bool                                `json:"case_sensitive"`
	Columns         []DLPDatasetUploadNewResponseColumn `json:"columns"`
	Secret          string                              `json:"secret" format:"password"`
	JSON            dlpDatasetUploadNewResponseJSON     `json:"-"`
}

// dlpDatasetUploadNewResponseJSON contains the JSON metadata for the struct
// [DLPDatasetUploadNewResponse]
type dlpDatasetUploadNewResponseJSON struct {
	EncodingVersion apijson.Field
	MaxCells        apijson.Field
	Version         apijson.Field
	CaseSensitive   apijson.Field
	Columns         apijson.Field
	Secret          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseColumn struct {
	EntryID      string                                         `json:"entry_id,required" format:"uuid"`
	HeaderName   string                                         `json:"header_name,required"`
	NumCells     int64                                          `json:"num_cells,required"`
	UploadStatus DLPDatasetUploadNewResponseColumnsUploadStatus `json:"upload_status,required"`
	JSON         dlpDatasetUploadNewResponseColumnJSON          `json:"-"`
}

// dlpDatasetUploadNewResponseColumnJSON contains the JSON metadata for the struct
// [DLPDatasetUploadNewResponseColumn]
type dlpDatasetUploadNewResponseColumnJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseColumn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseColumnJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseColumnsUploadStatus string

const (
	DLPDatasetUploadNewResponseColumnsUploadStatusEmpty      DLPDatasetUploadNewResponseColumnsUploadStatus = "empty"
	DLPDatasetUploadNewResponseColumnsUploadStatusUploading  DLPDatasetUploadNewResponseColumnsUploadStatus = "uploading"
	DLPDatasetUploadNewResponseColumnsUploadStatusPending    DLPDatasetUploadNewResponseColumnsUploadStatus = "pending"
	DLPDatasetUploadNewResponseColumnsUploadStatusProcessing DLPDatasetUploadNewResponseColumnsUploadStatus = "processing"
	DLPDatasetUploadNewResponseColumnsUploadStatusFailed     DLPDatasetUploadNewResponseColumnsUploadStatus = "failed"
	DLPDatasetUploadNewResponseColumnsUploadStatusComplete   DLPDatasetUploadNewResponseColumnsUploadStatus = "complete"
)

func (r DLPDatasetUploadNewResponseColumnsUploadStatus) IsKnown() bool {
	switch r {
	case DLPDatasetUploadNewResponseColumnsUploadStatusEmpty, DLPDatasetUploadNewResponseColumnsUploadStatusUploading, DLPDatasetUploadNewResponseColumnsUploadStatusPending, DLPDatasetUploadNewResponseColumnsUploadStatusProcessing, DLPDatasetUploadNewResponseColumnsUploadStatusFailed, DLPDatasetUploadNewResponseColumnsUploadStatusComplete:
		return true
	}
	return false
}

type DLPDatasetUploadEditResponse struct {
	ID              string                               `json:"id,required" format:"uuid"`
	Columns         []DLPDatasetUploadEditResponseColumn `json:"columns,required"`
	CreatedAt       time.Time                            `json:"created_at,required" format:"date-time"`
	EncodingVersion int64                                `json:"encoding_version,required"`
	Name            string                               `json:"name,required"`
	NumCells        int64                                `json:"num_cells,required"`
	Secret          bool                                 `json:"secret,required"`
	Status          DLPDatasetUploadEditResponseStatus   `json:"status,required"`
	// Stores when the dataset was last updated.
	//
	// This includes name or description changes as well as uploads.
	UpdatedAt     time.Time                            `json:"updated_at,required" format:"date-time"`
	Uploads       []DLPDatasetUploadEditResponseUpload `json:"uploads,required"`
	CaseSensitive bool                                 `json:"case_sensitive"`
	// The description of the dataset.
	Description string                           `json:"description,nullable"`
	JSON        dlpDatasetUploadEditResponseJSON `json:"-"`
}

// dlpDatasetUploadEditResponseJSON contains the JSON metadata for the struct
// [DLPDatasetUploadEditResponse]
type dlpDatasetUploadEditResponseJSON struct {
	ID              apijson.Field
	Columns         apijson.Field
	CreatedAt       apijson.Field
	EncodingVersion apijson.Field
	Name            apijson.Field
	NumCells        apijson.Field
	Secret          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	Uploads         apijson.Field
	CaseSensitive   apijson.Field
	Description     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseColumn struct {
	EntryID      string                                          `json:"entry_id,required" format:"uuid"`
	HeaderName   string                                          `json:"header_name,required"`
	NumCells     int64                                           `json:"num_cells,required"`
	UploadStatus DLPDatasetUploadEditResponseColumnsUploadStatus `json:"upload_status,required"`
	JSON         dlpDatasetUploadEditResponseColumnJSON          `json:"-"`
}

// dlpDatasetUploadEditResponseColumnJSON contains the JSON metadata for the struct
// [DLPDatasetUploadEditResponseColumn]
type dlpDatasetUploadEditResponseColumnJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseColumn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseColumnJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseColumnsUploadStatus string

const (
	DLPDatasetUploadEditResponseColumnsUploadStatusEmpty      DLPDatasetUploadEditResponseColumnsUploadStatus = "empty"
	DLPDatasetUploadEditResponseColumnsUploadStatusUploading  DLPDatasetUploadEditResponseColumnsUploadStatus = "uploading"
	DLPDatasetUploadEditResponseColumnsUploadStatusPending    DLPDatasetUploadEditResponseColumnsUploadStatus = "pending"
	DLPDatasetUploadEditResponseColumnsUploadStatusProcessing DLPDatasetUploadEditResponseColumnsUploadStatus = "processing"
	DLPDatasetUploadEditResponseColumnsUploadStatusFailed     DLPDatasetUploadEditResponseColumnsUploadStatus = "failed"
	DLPDatasetUploadEditResponseColumnsUploadStatusComplete   DLPDatasetUploadEditResponseColumnsUploadStatus = "complete"
)

func (r DLPDatasetUploadEditResponseColumnsUploadStatus) IsKnown() bool {
	switch r {
	case DLPDatasetUploadEditResponseColumnsUploadStatusEmpty, DLPDatasetUploadEditResponseColumnsUploadStatusUploading, DLPDatasetUploadEditResponseColumnsUploadStatusPending, DLPDatasetUploadEditResponseColumnsUploadStatusProcessing, DLPDatasetUploadEditResponseColumnsUploadStatusFailed, DLPDatasetUploadEditResponseColumnsUploadStatusComplete:
		return true
	}
	return false
}

type DLPDatasetUploadEditResponseStatus string

const (
	DLPDatasetUploadEditResponseStatusEmpty      DLPDatasetUploadEditResponseStatus = "empty"
	DLPDatasetUploadEditResponseStatusUploading  DLPDatasetUploadEditResponseStatus = "uploading"
	DLPDatasetUploadEditResponseStatusPending    DLPDatasetUploadEditResponseStatus = "pending"
	DLPDatasetUploadEditResponseStatusProcessing DLPDatasetUploadEditResponseStatus = "processing"
	DLPDatasetUploadEditResponseStatusFailed     DLPDatasetUploadEditResponseStatus = "failed"
	DLPDatasetUploadEditResponseStatusComplete   DLPDatasetUploadEditResponseStatus = "complete"
)

func (r DLPDatasetUploadEditResponseStatus) IsKnown() bool {
	switch r {
	case DLPDatasetUploadEditResponseStatusEmpty, DLPDatasetUploadEditResponseStatusUploading, DLPDatasetUploadEditResponseStatusPending, DLPDatasetUploadEditResponseStatusProcessing, DLPDatasetUploadEditResponseStatusFailed, DLPDatasetUploadEditResponseStatusComplete:
		return true
	}
	return false
}

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

func (r dlpDatasetUploadEditResponseUploadJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseUploadsStatus string

const (
	DLPDatasetUploadEditResponseUploadsStatusEmpty      DLPDatasetUploadEditResponseUploadsStatus = "empty"
	DLPDatasetUploadEditResponseUploadsStatusUploading  DLPDatasetUploadEditResponseUploadsStatus = "uploading"
	DLPDatasetUploadEditResponseUploadsStatusPending    DLPDatasetUploadEditResponseUploadsStatus = "pending"
	DLPDatasetUploadEditResponseUploadsStatusProcessing DLPDatasetUploadEditResponseUploadsStatus = "processing"
	DLPDatasetUploadEditResponseUploadsStatusFailed     DLPDatasetUploadEditResponseUploadsStatus = "failed"
	DLPDatasetUploadEditResponseUploadsStatusComplete   DLPDatasetUploadEditResponseUploadsStatus = "complete"
)

func (r DLPDatasetUploadEditResponseUploadsStatus) IsKnown() bool {
	switch r {
	case DLPDatasetUploadEditResponseUploadsStatusEmpty, DLPDatasetUploadEditResponseUploadsStatusUploading, DLPDatasetUploadEditResponseUploadsStatusPending, DLPDatasetUploadEditResponseUploadsStatusProcessing, DLPDatasetUploadEditResponseUploadsStatusFailed, DLPDatasetUploadEditResponseUploadsStatusComplete:
		return true
	}
	return false
}

type DLPDatasetUploadNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetUploadNewResponseEnvelope struct {
	Errors   []DLPDatasetUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetUploadNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetUploadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPDatasetUploadNewResponse                `json:"result"`
	JSON    dlpDatasetUploadNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelope]
type dlpDatasetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPDatasetUploadNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetUploadNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelopeErrors]
type dlpDatasetUploadNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUploadNewResponseEnvelopeErrorsSource]
type dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPDatasetUploadNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetUploadNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDatasetUploadNewResponseEnvelopeMessages]
type dlpDatasetUploadNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUploadNewResponseEnvelopeMessagesSource]
type dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetUploadNewResponseEnvelopeSuccess bool

const (
	DLPDatasetUploadNewResponseEnvelopeSuccessTrue DLPDatasetUploadNewResponseEnvelopeSuccess = true
)

func (r DLPDatasetUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDatasetUploadEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

func (r DLPDatasetUploadEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DLPDatasetUploadEditResponseEnvelope struct {
	Errors   []DLPDatasetUploadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetUploadEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetUploadEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPDatasetUploadEditResponse                `json:"result"`
	JSON    dlpDatasetUploadEditResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadEditResponseEnvelope]
type dlpDatasetUploadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPDatasetUploadEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetUploadEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPDatasetUploadEditResponseEnvelopeErrors]
type dlpDatasetUploadEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUploadEditResponseEnvelopeErrorsSource]
type dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPDatasetUploadEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetUploadEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDatasetUploadEditResponseEnvelopeMessages]
type dlpDatasetUploadEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPDatasetUploadEditResponseEnvelopeMessagesSource]
type dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetUploadEditResponseEnvelopeSuccess bool

const (
	DLPDatasetUploadEditResponseEnvelopeSuccessTrue DLPDatasetUploadEditResponseEnvelopeSuccess = true
)

func (r DLPDatasetUploadEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetUploadEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
