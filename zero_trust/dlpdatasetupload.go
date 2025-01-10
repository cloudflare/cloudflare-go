// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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
func (r *DLPDatasetUploadService) New(ctx context.Context, datasetID string, body DLPDatasetUploadNewParams, opts ...option.RequestOption) (res *NewVersion, err error) {
	var env DLPDatasetUploadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *DLPDatasetUploadService) Edit(ctx context.Context, datasetID string, version int64, params DLPDatasetUploadEditParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env DLPDatasetUploadEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", params.AccountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NewVersion struct {
	EncodingVersion int64              `json:"encoding_version,required"`
	MaxCells        int64              `json:"max_cells,required"`
	Version         int64              `json:"version,required"`
	Columns         []NewVersionColumn `json:"columns"`
	Secret          string             `json:"secret" format:"password"`
	JSON            newVersionJSON     `json:"-"`
}

// newVersionJSON contains the JSON metadata for the struct [NewVersion]
type newVersionJSON struct {
	EncodingVersion apijson.Field
	MaxCells        apijson.Field
	Version         apijson.Field
	Columns         apijson.Field
	Secret          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NewVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVersionJSON) RawJSON() string {
	return r.raw
}

type NewVersionColumn struct {
	EntryID      string                        `json:"entry_id,required" format:"uuid"`
	HeaderName   string                        `json:"header_name,required"`
	NumCells     int64                         `json:"num_cells,required"`
	UploadStatus NewVersionColumnsUploadStatus `json:"upload_status,required"`
	JSON         newVersionColumnJSON          `json:"-"`
}

// newVersionColumnJSON contains the JSON metadata for the struct
// [NewVersionColumn]
type newVersionColumnJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NewVersionColumn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVersionColumnJSON) RawJSON() string {
	return r.raw
}

type NewVersionColumnsUploadStatus string

const (
	NewVersionColumnsUploadStatusEmpty      NewVersionColumnsUploadStatus = "empty"
	NewVersionColumnsUploadStatusUploading  NewVersionColumnsUploadStatus = "uploading"
	NewVersionColumnsUploadStatusProcessing NewVersionColumnsUploadStatus = "processing"
	NewVersionColumnsUploadStatusFailed     NewVersionColumnsUploadStatus = "failed"
	NewVersionColumnsUploadStatusComplete   NewVersionColumnsUploadStatus = "complete"
)

func (r NewVersionColumnsUploadStatus) IsKnown() bool {
	switch r {
	case NewVersionColumnsUploadStatusEmpty, NewVersionColumnsUploadStatusUploading, NewVersionColumnsUploadStatusProcessing, NewVersionColumnsUploadStatusFailed, NewVersionColumnsUploadStatusComplete:
		return true
	}
	return false
}

type DLPDatasetUploadNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetUploadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPDatasetUploadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  NewVersion                                 `json:"result"`
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

// Whether the API call was successful
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
	Body      string              `json:"body,required"`
}

func (r DLPDatasetUploadEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPDatasetUploadEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPDatasetUploadEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Dataset                                     `json:"result"`
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

// Whether the API call was successful
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
