// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
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

// IndicatorFeedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIndicatorFeedService] method
// instead.
type IndicatorFeedService struct {
	Options     []option.RequestOption
	Permissions *IndicatorFeedPermissionService
}

// NewIndicatorFeedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIndicatorFeedService(opts ...option.RequestOption) (r *IndicatorFeedService) {
	r = &IndicatorFeedService{}
	r.Options = opts
	r.Permissions = NewIndicatorFeedPermissionService(opts...)
	return
}

// Create new indicator feed
func (r *IndicatorFeedService) New(ctx context.Context, params IndicatorFeedNewParams, opts ...option.RequestOption) (res *IndicatorFeedNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update indicator feed data
func (r *IndicatorFeedService) Update(ctx context.Context, feedID int64, params IndicatorFeedUpdateParams, opts ...option.RequestOption) (res *IndicatorFeedUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/snapshot", params.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feeds owned by this account
func (r *IndicatorFeedService) List(ctx context.Context, query IndicatorFeedListParams, opts ...option.RequestOption) (res *pagination.SinglePage[IndicatorFeedListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", query.AccountID)
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

// Get indicator feeds owned by this account
func (r *IndicatorFeedService) ListAutoPaging(ctx context.Context, query IndicatorFeedListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[IndicatorFeedListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Get indicator feed data
func (r *IndicatorFeedService) Data(ctx context.Context, feedID int64, query IndicatorFeedDataParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *IndicatorFeedService) Get(ctx context.Context, feedID int64, query IndicatorFeedGetParams, opts ...option.RequestOption) (res *IndicatorFeedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndicatorFeedNewResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                       `json:"name"`
	JSON indicatorFeedNewResponseJSON `json:"-"`
}

// indicatorFeedNewResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedNewResponse]
type indicatorFeedNewResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponse struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                          `json:"status"`
	JSON   indicatorFeedUpdateResponseJSON `json:"-"`
}

// indicatorFeedUpdateResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedUpdateResponse]
type indicatorFeedUpdateResponseJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedListResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                        `json:"name"`
	JSON indicatorFeedListResponseJSON `json:"-"`
}

// indicatorFeedListResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedListResponse]
type indicatorFeedListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedListResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Status of the latest snapshot uploaded
	LatestUploadStatus IndicatorFeedGetResponseLatestUploadStatus `json:"latest_upload_status"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                       `json:"name"`
	JSON indicatorFeedGetResponseJSON `json:"-"`
}

// indicatorFeedGetResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedGetResponse]
type indicatorFeedGetResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	LatestUploadStatus apijson.Field
	ModifiedOn         apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IndicatorFeedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the latest snapshot uploaded
type IndicatorFeedGetResponseLatestUploadStatus string

const (
	IndicatorFeedGetResponseLatestUploadStatusMirroring    IndicatorFeedGetResponseLatestUploadStatus = "Mirroring"
	IndicatorFeedGetResponseLatestUploadStatusUnifying     IndicatorFeedGetResponseLatestUploadStatus = "Unifying"
	IndicatorFeedGetResponseLatestUploadStatusLoading      IndicatorFeedGetResponseLatestUploadStatus = "Loading"
	IndicatorFeedGetResponseLatestUploadStatusProvisioning IndicatorFeedGetResponseLatestUploadStatus = "Provisioning"
	IndicatorFeedGetResponseLatestUploadStatusComplete     IndicatorFeedGetResponseLatestUploadStatus = "Complete"
	IndicatorFeedGetResponseLatestUploadStatusError        IndicatorFeedGetResponseLatestUploadStatus = "Error"
)

func (r IndicatorFeedGetResponseLatestUploadStatus) IsKnown() bool {
	switch r {
	case IndicatorFeedGetResponseLatestUploadStatusMirroring, IndicatorFeedGetResponseLatestUploadStatusUnifying, IndicatorFeedGetResponseLatestUploadStatusLoading, IndicatorFeedGetResponseLatestUploadStatusProvisioning, IndicatorFeedGetResponseLatestUploadStatusComplete, IndicatorFeedGetResponseLatestUploadStatusError:
		return true
	}
	return false
}

type IndicatorFeedNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The description of the example test
	Description param.Field[string] `json:"description"`
	// The name of the indicator feed
	Name param.Field[string] `json:"name"`
}

func (r IndicatorFeedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   IndicatorFeedNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedNewResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndicatorFeedNewResponseEnvelope]
type indicatorFeedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedNewResponseEnvelopeSuccess bool

const (
	IndicatorFeedNewResponseEnvelopeSuccessTrue IndicatorFeedNewResponseEnvelopeSuccess = true
)

func (r IndicatorFeedNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The file to upload
	Source param.Field[string] `json:"source"`
}

func (r IndicatorFeedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   IndicatorFeedUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedUpdateResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndicatorFeedUpdateResponseEnvelope]
type indicatorFeedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedUpdateResponseEnvelopeSuccess bool

const (
	IndicatorFeedUpdateResponseEnvelopeSuccessTrue IndicatorFeedUpdateResponseEnvelopeSuccess = true
)

func (r IndicatorFeedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedDataParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   IndicatorFeedGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedGetResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndicatorFeedGetResponseEnvelope]
type indicatorFeedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedGetResponseEnvelopeSuccess bool

const (
	IndicatorFeedGetResponseEnvelopeSuccessTrue IndicatorFeedGetResponseEnvelopeSuccess = true
)

func (r IndicatorFeedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
