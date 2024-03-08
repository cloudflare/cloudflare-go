// File generated from our OpenAPI spec by Stainless.

package intel

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
func (r *IndicatorFeedService) List(ctx context.Context, query IndicatorFeedListParams, opts ...option.RequestOption) (res *[]IndicatorFeedListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feed data
func (r *IndicatorFeedService) Data(ctx context.Context, feedID int64, query IndicatorFeedDataParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *IndicatorFeedService) Get(ctx context.Context, feedID int64, query IndicatorFeedGetParams, opts ...option.RequestOption) (res *IndicatorFeedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
	Errors   []IndicatorFeedNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IndicatorFeedNewResponse                   `json:"result,required"`
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

type IndicatorFeedNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    indicatorFeedNewResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedNewResponseEnvelopeErrors]
type indicatorFeedNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    indicatorFeedNewResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndicatorFeedNewResponseEnvelopeMessages]
type indicatorFeedNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedNewResponseEnvelopeSuccess bool

const (
	IndicatorFeedNewResponseEnvelopeSuccessTrue IndicatorFeedNewResponseEnvelopeSuccess = true
)

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
	Errors   []IndicatorFeedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   IndicatorFeedUpdateResponse                   `json:"result,required"`
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

type IndicatorFeedUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    indicatorFeedUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedUpdateResponseEnvelopeErrors]
type indicatorFeedUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    indicatorFeedUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IndicatorFeedUpdateResponseEnvelopeMessages]
type indicatorFeedUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedUpdateResponseEnvelopeSuccess bool

const (
	IndicatorFeedUpdateResponseEnvelopeSuccessTrue IndicatorFeedUpdateResponseEnvelopeSuccess = true
)

type IndicatorFeedListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedListResponseEnvelope struct {
	Errors   []IndicatorFeedListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IndicatorFeedListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedListResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedListResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndicatorFeedListResponseEnvelope]
type indicatorFeedListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    indicatorFeedListResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedListResponseEnvelopeErrors]
type indicatorFeedListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    indicatorFeedListResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndicatorFeedListResponseEnvelopeMessages]
type indicatorFeedListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedListResponseEnvelopeSuccess bool

const (
	IndicatorFeedListResponseEnvelopeSuccessTrue IndicatorFeedListResponseEnvelopeSuccess = true
)

type IndicatorFeedDataParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedGetResponseEnvelope struct {
	Errors   []IndicatorFeedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IndicatorFeedGetResponse                   `json:"result,required"`
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

type IndicatorFeedGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    indicatorFeedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedGetResponseEnvelopeErrors]
type indicatorFeedGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    indicatorFeedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndicatorFeedGetResponseEnvelopeMessages]
type indicatorFeedGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedGetResponseEnvelopeSuccess bool

const (
	IndicatorFeedGetResponseEnvelopeSuccessTrue IndicatorFeedGetResponseEnvelopeSuccess = true
)
