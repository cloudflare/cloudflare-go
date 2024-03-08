// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustDLPDatasetUploadService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDLPDatasetUploadService] method instead.
type ZeroTrustDLPDatasetUploadService struct {
	Options []option.RequestOption
}

// NewZeroTrustDLPDatasetUploadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDLPDatasetUploadService(opts ...option.RequestOption) (r *ZeroTrustDLPDatasetUploadService) {
	r = &ZeroTrustDLPDatasetUploadService{}
	r.Options = opts
	return
}

// Prepare to upload a new version of a dataset.
func (r *ZeroTrustDLPDatasetUploadService) New(ctx context.Context, datasetID string, body ZeroTrustDLPDatasetUploadNewParams, opts ...option.RequestOption) (res *ZeroTrustDLPDatasetUploadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPDatasetUploadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upload a new version of a dataset.
func (r *ZeroTrustDLPDatasetUploadService) Edit(ctx context.Context, datasetID string, version int64, body ZeroTrustDLPDatasetUploadEditParams, opts ...option.RequestOption) (res *ZeroTrustDLPDatasetUploadEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPDatasetUploadEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", body.AccountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDLPDatasetUploadNewResponse struct {
	MaxCells int64                                    `json:"max_cells,required"`
	Version  int64                                    `json:"version,required"`
	Secret   string                                   `json:"secret" format:"password"`
	JSON     zeroTrustDLPDatasetUploadNewResponseJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetUploadNewResponse]
type zeroTrustDLPDatasetUploadNewResponseJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditResponse struct {
	ID          string                                        `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                                     `json:"created_at,required" format:"date-time"`
	Name        string                                        `json:"name,required"`
	NumCells    int64                                         `json:"num_cells,required"`
	Secret      bool                                          `json:"secret,required"`
	Status      ZeroTrustDLPDatasetUploadEditResponseStatus   `json:"status,required"`
	UpdatedAt   time.Time                                     `json:"updated_at,required" format:"date-time"`
	Uploads     []ZeroTrustDLPDatasetUploadEditResponseUpload `json:"uploads,required"`
	Description string                                        `json:"description,nullable"`
	JSON        zeroTrustDLPDatasetUploadEditResponseJSON     `json:"-"`
}

// zeroTrustDLPDatasetUploadEditResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPDatasetUploadEditResponse]
type zeroTrustDLPDatasetUploadEditResponseJSON struct {
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

func (r *ZeroTrustDLPDatasetUploadEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditResponseStatus string

const (
	ZeroTrustDLPDatasetUploadEditResponseStatusEmpty     ZeroTrustDLPDatasetUploadEditResponseStatus = "empty"
	ZeroTrustDLPDatasetUploadEditResponseStatusUploading ZeroTrustDLPDatasetUploadEditResponseStatus = "uploading"
	ZeroTrustDLPDatasetUploadEditResponseStatusFailed    ZeroTrustDLPDatasetUploadEditResponseStatus = "failed"
	ZeroTrustDLPDatasetUploadEditResponseStatusComplete  ZeroTrustDLPDatasetUploadEditResponseStatus = "complete"
)

type ZeroTrustDLPDatasetUploadEditResponseUpload struct {
	NumCells int64                                              `json:"num_cells,required"`
	Status   ZeroTrustDLPDatasetUploadEditResponseUploadsStatus `json:"status,required"`
	Version  int64                                              `json:"version,required"`
	JSON     zeroTrustDLPDatasetUploadEditResponseUploadJSON    `json:"-"`
}

// zeroTrustDLPDatasetUploadEditResponseUploadJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetUploadEditResponseUpload]
type zeroTrustDLPDatasetUploadEditResponseUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadEditResponseUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadEditResponseUploadJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditResponseUploadsStatus string

const (
	ZeroTrustDLPDatasetUploadEditResponseUploadsStatusEmpty     ZeroTrustDLPDatasetUploadEditResponseUploadsStatus = "empty"
	ZeroTrustDLPDatasetUploadEditResponseUploadsStatusUploading ZeroTrustDLPDatasetUploadEditResponseUploadsStatus = "uploading"
	ZeroTrustDLPDatasetUploadEditResponseUploadsStatusFailed    ZeroTrustDLPDatasetUploadEditResponseUploadsStatus = "failed"
	ZeroTrustDLPDatasetUploadEditResponseUploadsStatusComplete  ZeroTrustDLPDatasetUploadEditResponseUploadsStatus = "complete"
)

type ZeroTrustDLPDatasetUploadNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPDatasetUploadNewResponseEnvelope struct {
	Errors     []ZeroTrustDLPDatasetUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []ZeroTrustDLPDatasetUploadNewResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                                   `json:"success,required"`
	Result     ZeroTrustDLPDatasetUploadNewResponse                   `json:"result"`
	ResultInfo ZeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPDatasetUploadNewResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPDatasetUploadNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetUploadNewResponseEnvelope]
type zeroTrustDLPDatasetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadNewResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDLPDatasetUploadNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPDatasetUploadNewResponseEnvelopeErrors]
type zeroTrustDLPDatasetUploadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadNewResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDLPDatasetUploadNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDLPDatasetUploadNewResponseEnvelopeMessages]
type zeroTrustDLPDatasetUploadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                      `json:"total_count,required"`
	JSON       zeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfo]
type zeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPDatasetUploadEditResponseEnvelope struct {
	Errors     []ZeroTrustDLPDatasetUploadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []ZeroTrustDLPDatasetUploadEditResponseEnvelopeMessages `json:"messages,required"`
	Success    bool                                                    `json:"success,required"`
	Result     ZeroTrustDLPDatasetUploadEditResponse                   `json:"result"`
	ResultInfo ZeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPDatasetUploadEditResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPDatasetUploadEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDLPDatasetUploadEditResponseEnvelope]
type zeroTrustDLPDatasetUploadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustDLPDatasetUploadEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPDatasetUploadEditResponseEnvelopeErrors]
type zeroTrustDLPDatasetUploadEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDLPDatasetUploadEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDLPDatasetUploadEditResponseEnvelopeMessages]
type zeroTrustDLPDatasetUploadEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                       `json:"total_count,required"`
	JSON       zeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfo]
type zeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPDatasetUploadEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
