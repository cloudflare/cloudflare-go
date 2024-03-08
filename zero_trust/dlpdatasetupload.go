// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *DLPDatasetUploadService) New(ctx context.Context, datasetID string, body DLPDatasetUploadNewParams, opts ...option.RequestOption) (res *DLPDatasetNewVersion, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUploadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upload a new version of a dataset.
func (r *DLPDatasetUploadService) Edit(ctx context.Context, datasetID string, version int64, body DLPDatasetUploadEditParams, opts ...option.RequestOption) (res *DLPDataset, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPDatasetUploadEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", body.AccountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPDatasetNewVersion struct {
	MaxCells int64                    `json:"max_cells,required"`
	Version  int64                    `json:"version,required"`
	Secret   string                   `json:"secret" format:"password"`
	JSON     dlpDatasetNewVersionJSON `json:"-"`
}

// dlpDatasetNewVersionJSON contains the JSON metadata for the struct
// [DLPDatasetNewVersion]
type dlpDatasetNewVersionJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetNewVersionJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetUploadNewResponseEnvelope struct {
	Errors     []DLPDatasetUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetUploadNewResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                          `json:"success,required"`
	Result     DLPDatasetNewVersion                          `json:"result"`
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

func (r dlpDatasetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUploadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUploadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUploadNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetUploadEditResponseEnvelope struct {
	Errors     []DLPDatasetUploadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []DLPDatasetUploadEditResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                           `json:"success,required"`
	Result     DLPDataset                                     `json:"result"`
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

func (r dlpDatasetUploadEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUploadEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUploadEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dlpDatasetUploadEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
