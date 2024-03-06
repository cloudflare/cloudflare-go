// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelIndicatorFeedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelIndicatorFeedService] method
// instead.
type IntelIndicatorFeedService struct {
	Options     []option.RequestOption
	Permissions *IntelIndicatorFeedPermissionService
}

// NewIntelIndicatorFeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntelIndicatorFeedService(opts ...option.RequestOption) (r *IntelIndicatorFeedService) {
	r = &IntelIndicatorFeedService{}
	r.Options = opts
	r.Permissions = NewIntelIndicatorFeedPermissionService(opts...)
	return
}

// Create new indicator feed
func (r *IntelIndicatorFeedService) New(ctx context.Context, params IntelIndicatorFeedNewParams, opts ...option.RequestOption) (res *IntelIndicatorFeedNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update indicator feed data
func (r *IntelIndicatorFeedService) Update(ctx context.Context, feedID int64, params IntelIndicatorFeedUpdateParams, opts ...option.RequestOption) (res *IntelIndicatorFeedUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/snapshot", params.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feeds owned by this account
func (r *IntelIndicatorFeedService) List(ctx context.Context, query IntelIndicatorFeedListParams, opts ...option.RequestOption) (res *[]IntelIndicatorFeedListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feed data
func (r *IntelIndicatorFeedService) Data(ctx context.Context, feedID int64, query IntelIndicatorFeedDataParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *IntelIndicatorFeedService) Get(ctx context.Context, feedID int64, query IntelIndicatorFeedGetParams, opts ...option.RequestOption) (res *IntelIndicatorFeedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIndicatorFeedNewResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                            `json:"name"`
	JSON intelIndicatorFeedNewResponseJSON `json:"-"`
}

// intelIndicatorFeedNewResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedNewResponse]
type intelIndicatorFeedNewResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedUpdateResponse struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                               `json:"status"`
	JSON   intelIndicatorFeedUpdateResponseJSON `json:"-"`
}

// intelIndicatorFeedUpdateResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedUpdateResponse]
type intelIndicatorFeedUpdateResponseJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedListResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                             `json:"name"`
	JSON intelIndicatorFeedListResponseJSON `json:"-"`
}

// intelIndicatorFeedListResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedListResponse]
type intelIndicatorFeedListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedGetResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Status of the latest snapshot uploaded
	LatestUploadStatus IntelIndicatorFeedGetResponseLatestUploadStatus `json:"latest_upload_status"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                            `json:"name"`
	JSON intelIndicatorFeedGetResponseJSON `json:"-"`
}

// intelIndicatorFeedGetResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedGetResponse]
type intelIndicatorFeedGetResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	LatestUploadStatus apijson.Field
	ModifiedOn         apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the latest snapshot uploaded
type IntelIndicatorFeedGetResponseLatestUploadStatus string

const (
	IntelIndicatorFeedGetResponseLatestUploadStatusMirroring    IntelIndicatorFeedGetResponseLatestUploadStatus = "Mirroring"
	IntelIndicatorFeedGetResponseLatestUploadStatusUnifying     IntelIndicatorFeedGetResponseLatestUploadStatus = "Unifying"
	IntelIndicatorFeedGetResponseLatestUploadStatusLoading      IntelIndicatorFeedGetResponseLatestUploadStatus = "Loading"
	IntelIndicatorFeedGetResponseLatestUploadStatusProvisioning IntelIndicatorFeedGetResponseLatestUploadStatus = "Provisioning"
	IntelIndicatorFeedGetResponseLatestUploadStatusComplete     IntelIndicatorFeedGetResponseLatestUploadStatus = "Complete"
	IntelIndicatorFeedGetResponseLatestUploadStatusError        IntelIndicatorFeedGetResponseLatestUploadStatus = "Error"
)

type IntelIndicatorFeedNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The description of the example test
	Description param.Field[string] `json:"description"`
	// The name of the indicator feed
	Name param.Field[string] `json:"name"`
}

func (r IntelIndicatorFeedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedNewResponseEnvelope struct {
	Errors   []IntelIndicatorFeedNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedNewResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedNewResponseEnvelope]
type intelIndicatorFeedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    intelIndicatorFeedNewResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedNewResponseEnvelopeErrors]
type intelIndicatorFeedNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    intelIndicatorFeedNewResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedNewResponseEnvelopeMessages]
type intelIndicatorFeedNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedNewResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedNewResponseEnvelopeSuccessTrue IntelIndicatorFeedNewResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The file to upload
	Source param.Field[string] `json:"source"`
}

func (r IntelIndicatorFeedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedUpdateResponseEnvelope struct {
	Errors   []IntelIndicatorFeedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedUpdateResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedUpdateResponseEnvelope]
type intelIndicatorFeedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    intelIndicatorFeedUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedUpdateResponseEnvelopeErrors]
type intelIndicatorFeedUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    intelIndicatorFeedUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedUpdateResponseEnvelopeMessages]
type intelIndicatorFeedUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedUpdateResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedUpdateResponseEnvelopeSuccessTrue IntelIndicatorFeedUpdateResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IntelIndicatorFeedListResponseEnvelope struct {
	Errors   []IntelIndicatorFeedListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIndicatorFeedListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedListResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedListResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedListResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedListResponseEnvelope]
type intelIndicatorFeedListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    intelIndicatorFeedListResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedListResponseEnvelopeErrors]
type intelIndicatorFeedListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    intelIndicatorFeedListResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedListResponseEnvelopeMessages]
type intelIndicatorFeedListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedListResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedListResponseEnvelopeSuccessTrue IntelIndicatorFeedListResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedDataParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IntelIndicatorFeedGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IntelIndicatorFeedGetResponseEnvelope struct {
	Errors   []IntelIndicatorFeedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedGetResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedGetResponseEnvelope]
type intelIndicatorFeedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    intelIndicatorFeedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedGetResponseEnvelopeErrors]
type intelIndicatorFeedGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    intelIndicatorFeedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedGetResponseEnvelopeMessages]
type intelIndicatorFeedGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedGetResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedGetResponseEnvelopeSuccessTrue IntelIndicatorFeedGetResponseEnvelopeSuccess = true
)
