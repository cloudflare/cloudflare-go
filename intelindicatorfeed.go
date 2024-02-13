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
	Options []option.RequestOption
}

// NewIntelIndicatorFeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntelIndicatorFeedService(opts ...option.RequestOption) (r *IntelIndicatorFeedService) {
	r = &IntelIndicatorFeedService{}
	r.Options = opts
	return
}

// Create new indicator feed
func (r *IntelIndicatorFeedService) New(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedNewParams, opts ...option.RequestOption) (res *IntelIndicatorFeedNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feeds owned by this account
func (r *IntelIndicatorFeedService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]IntelIndicatorFeedListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feed data
func (r *IntelIndicatorFeedService) Data(ctx context.Context, accountIdentifier string, feedID int64, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", accountIdentifier, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *IntelIndicatorFeedService) Get(ctx context.Context, accountIdentifier string, feedID int64, opts ...option.RequestOption) (res *IntelIndicatorFeedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", accountIdentifier, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Grant permission to indicator feed
func (r *IntelIndicatorFeedService) PermissionsAdd(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedPermissionsAddParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionsAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedPermissionsAddResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revoke permission to indicator feed
func (r *IntelIndicatorFeedService) PermissionsRemove(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedPermissionsRemoveParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionsRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedPermissionsRemoveResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List indicator feed permissions
func (r *IntelIndicatorFeedService) PermissionsView(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]IntelIndicatorFeedPermissionsViewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedPermissionsViewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update indicator feed data
func (r *IntelIndicatorFeedService) Snapshot(ctx context.Context, accountIdentifier string, feedID int64, body IntelIndicatorFeedSnapshotParams, opts ...option.RequestOption) (res *IntelIndicatorFeedSnapshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedSnapshotResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/snapshot", accountIdentifier, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

type IntelIndicatorFeedPermissionsAddResponse struct {
	// Whether the update succeeded or not
	Success bool                                         `json:"success"`
	JSON    intelIndicatorFeedPermissionsAddResponseJSON `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionsAddResponse]
type intelIndicatorFeedPermissionsAddResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsRemoveResponse struct {
	// Whether the update succeeded or not
	Success bool                                            `json:"success"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseJSON `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedPermissionsRemoveResponse]
type intelIndicatorFeedPermissionsRemoveResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsViewResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// The name of the indicator feed
	Name string                                        `json:"name"`
	JSON intelIndicatorFeedPermissionsViewResponseJSON `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionsViewResponse]
type intelIndicatorFeedPermissionsViewResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedSnapshotResponse struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                                 `json:"status"`
	JSON   intelIndicatorFeedSnapshotResponseJSON `json:"-"`
}

// intelIndicatorFeedSnapshotResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedSnapshotResponse]
type intelIndicatorFeedSnapshotResponseJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedNewParams struct {
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

type IntelIndicatorFeedPermissionsAddParams struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionsAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedPermissionsAddResponseEnvelope struct {
	Errors   []IntelIndicatorFeedPermissionsAddResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedPermissionsAddResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedPermissionsAddResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionsAddResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedPermissionsAddResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseEnvelopeJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsAddResponseEnvelope]
type intelIndicatorFeedPermissionsAddResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsAddResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsAddResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedPermissionsAddResponseEnvelopeErrors]
type intelIndicatorFeedPermissionsAddResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsAddResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsAddResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionsAddResponseEnvelopeMessages]
type intelIndicatorFeedPermissionsAddResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionsAddResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedPermissionsAddResponseEnvelopeSuccessTrue IntelIndicatorFeedPermissionsAddResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedPermissionsRemoveParams struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionsRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedPermissionsRemoveResponseEnvelope struct {
	Errors   []IntelIndicatorFeedPermissionsRemoveResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedPermissionsRemoveResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedPermissionsRemoveResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionsRemoveResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseEnvelopeJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedPermissionsRemoveResponseEnvelope]
type intelIndicatorFeedPermissionsRemoveResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsRemoveResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionsRemoveResponseEnvelopeErrors]
type intelIndicatorFeedPermissionsRemoveResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsRemoveResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [IntelIndicatorFeedPermissionsRemoveResponseEnvelopeMessages]
type intelIndicatorFeedPermissionsRemoveResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionsRemoveResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedPermissionsRemoveResponseEnvelopeSuccessTrue IntelIndicatorFeedPermissionsRemoveResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedPermissionsViewResponseEnvelope struct {
	Errors   []IntelIndicatorFeedPermissionsViewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedPermissionsViewResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIndicatorFeedPermissionsViewResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionsViewResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedPermissionsViewResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseEnvelopeJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsViewResponseEnvelope]
type intelIndicatorFeedPermissionsViewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsViewResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsViewResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionsViewResponseEnvelopeErrors]
type intelIndicatorFeedPermissionsViewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsViewResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsViewResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionsViewResponseEnvelopeMessages]
type intelIndicatorFeedPermissionsViewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionsViewResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedPermissionsViewResponseEnvelopeSuccessTrue IntelIndicatorFeedPermissionsViewResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedSnapshotParams struct {
	// The file to upload
	Source param.Field[string] `json:"source"`
}

func (r IntelIndicatorFeedSnapshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedSnapshotResponseEnvelope struct {
	Errors   []IntelIndicatorFeedSnapshotResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedSnapshotResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedSnapshotResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedSnapshotResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedSnapshotResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedSnapshotResponseEnvelopeJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedSnapshotResponseEnvelope]
type intelIndicatorFeedSnapshotResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedSnapshotResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    intelIndicatorFeedSnapshotResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedSnapshotResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedSnapshotResponseEnvelopeErrors]
type intelIndicatorFeedSnapshotResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedSnapshotResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    intelIndicatorFeedSnapshotResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedSnapshotResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedSnapshotResponseEnvelopeMessages]
type intelIndicatorFeedSnapshotResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedSnapshotResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedSnapshotResponseEnvelopeSuccessTrue IntelIndicatorFeedSnapshotResponseEnvelopeSuccess = true
)
