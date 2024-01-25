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

// AccountAccessTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessTagService] method
// instead.
type AccountAccessTagService struct {
	Options []option.RequestOption
}

// NewAccountAccessTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessTagService(opts ...option.RequestOption) (r *AccountAccessTagService) {
	r = &AccountAccessTagService{}
	r.Options = opts
	return
}

// Create a tag
func (r *AccountAccessTagService) NewTag(ctx context.Context, identifier string, body AccountAccessTagNewTagParams, opts ...option.RequestOption) (res *AccountAccessTagNewTagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a tag
func (r *AccountAccessTagService) DeleteTag(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *AccountAccessTagDeleteTagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a tag
func (r *AccountAccessTagService) GetTag(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *AccountAccessTagGetTagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List tags
func (r *AccountAccessTagService) ListTags(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessTagListTagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a tag
func (r *AccountAccessTagService) UpdateTag(ctx context.Context, identifier string, params AccountAccessTagUpdateTagParams, opts ...option.RequestOption) (res *AccountAccessTagUpdateTagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, params.PathName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type AccountAccessTagNewTagResponse struct {
	Errors   []AccountAccessTagNewTagResponseError   `json:"errors"`
	Messages []AccountAccessTagNewTagResponseMessage `json:"messages"`
	// A tag
	Result AccountAccessTagNewTagResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountAccessTagNewTagResponseSuccess `json:"success"`
	JSON    accountAccessTagNewTagResponseJSON    `json:"-"`
}

// accountAccessTagNewTagResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagNewTagResponse]
type accountAccessTagNewTagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagNewTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagNewTagResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountAccessTagNewTagResponseErrorJSON `json:"-"`
}

// accountAccessTagNewTagResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessTagNewTagResponseError]
type accountAccessTagNewTagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagNewTagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagNewTagResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessTagNewTagResponseMessageJSON `json:"-"`
}

// accountAccessTagNewTagResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessTagNewTagResponseMessage]
type accountAccessTagNewTagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagNewTagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccountAccessTagNewTagResponseResult struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                                    `json:"app_count"`
	CreatedAt time.Time                                `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      accountAccessTagNewTagResponseResultJSON `json:"-"`
}

// accountAccessTagNewTagResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessTagNewTagResponseResult]
type accountAccessTagNewTagResponseResultJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagNewTagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessTagNewTagResponseSuccess bool

const (
	AccountAccessTagNewTagResponseSuccessTrue AccountAccessTagNewTagResponseSuccess = true
)

type AccountAccessTagDeleteTagResponse struct {
	Errors   []AccountAccessTagDeleteTagResponseError   `json:"errors"`
	Messages []AccountAccessTagDeleteTagResponseMessage `json:"messages"`
	Result   AccountAccessTagDeleteTagResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessTagDeleteTagResponseSuccess `json:"success"`
	JSON    accountAccessTagDeleteTagResponseJSON    `json:"-"`
}

// accountAccessTagDeleteTagResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagDeleteTagResponse]
type accountAccessTagDeleteTagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagDeleteTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagDeleteTagResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountAccessTagDeleteTagResponseErrorJSON `json:"-"`
}

// accountAccessTagDeleteTagResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessTagDeleteTagResponseError]
type accountAccessTagDeleteTagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagDeleteTagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagDeleteTagResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAccessTagDeleteTagResponseMessageJSON `json:"-"`
}

// accountAccessTagDeleteTagResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessTagDeleteTagResponseMessage]
type accountAccessTagDeleteTagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagDeleteTagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagDeleteTagResponseResult struct {
	// The name of the tag
	Name string                                      `json:"name"`
	JSON accountAccessTagDeleteTagResponseResultJSON `json:"-"`
}

// accountAccessTagDeleteTagResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessTagDeleteTagResponseResult]
type accountAccessTagDeleteTagResponseResultJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagDeleteTagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessTagDeleteTagResponseSuccess bool

const (
	AccountAccessTagDeleteTagResponseSuccessTrue AccountAccessTagDeleteTagResponseSuccess = true
)

type AccountAccessTagGetTagResponse struct {
	Errors   []AccountAccessTagGetTagResponseError   `json:"errors"`
	Messages []AccountAccessTagGetTagResponseMessage `json:"messages"`
	// A tag
	Result AccountAccessTagGetTagResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountAccessTagGetTagResponseSuccess `json:"success"`
	JSON    accountAccessTagGetTagResponseJSON    `json:"-"`
}

// accountAccessTagGetTagResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagGetTagResponse]
type accountAccessTagGetTagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagGetTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagGetTagResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountAccessTagGetTagResponseErrorJSON `json:"-"`
}

// accountAccessTagGetTagResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessTagGetTagResponseError]
type accountAccessTagGetTagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagGetTagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagGetTagResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessTagGetTagResponseMessageJSON `json:"-"`
}

// accountAccessTagGetTagResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessTagGetTagResponseMessage]
type accountAccessTagGetTagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagGetTagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccountAccessTagGetTagResponseResult struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                                    `json:"app_count"`
	CreatedAt time.Time                                `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      accountAccessTagGetTagResponseResultJSON `json:"-"`
}

// accountAccessTagGetTagResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessTagGetTagResponseResult]
type accountAccessTagGetTagResponseResultJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagGetTagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessTagGetTagResponseSuccess bool

const (
	AccountAccessTagGetTagResponseSuccessTrue AccountAccessTagGetTagResponseSuccess = true
)

type AccountAccessTagListTagsResponse struct {
	Errors     []AccountAccessTagListTagsResponseError    `json:"errors"`
	Messages   []AccountAccessTagListTagsResponseMessage  `json:"messages"`
	Result     []AccountAccessTagListTagsResponseResult   `json:"result"`
	ResultInfo AccountAccessTagListTagsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessTagListTagsResponseSuccess `json:"success"`
	JSON    accountAccessTagListTagsResponseJSON    `json:"-"`
}

// accountAccessTagListTagsResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagListTagsResponse]
type accountAccessTagListTagsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagListTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagListTagsResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessTagListTagsResponseErrorJSON `json:"-"`
}

// accountAccessTagListTagsResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessTagListTagsResponseError]
type accountAccessTagListTagsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagListTagsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagListTagsResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAccessTagListTagsResponseMessageJSON `json:"-"`
}

// accountAccessTagListTagsResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessTagListTagsResponseMessage]
type accountAccessTagListTagsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagListTagsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccountAccessTagListTagsResponseResult struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                                      `json:"app_count"`
	CreatedAt time.Time                                  `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accountAccessTagListTagsResponseResultJSON `json:"-"`
}

// accountAccessTagListTagsResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessTagListTagsResponseResult]
type accountAccessTagListTagsResponseResultJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagListTagsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagListTagsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       accountAccessTagListTagsResponseResultInfoJSON `json:"-"`
}

// accountAccessTagListTagsResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAccessTagListTagsResponseResultInfo]
type accountAccessTagListTagsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagListTagsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessTagListTagsResponseSuccess bool

const (
	AccountAccessTagListTagsResponseSuccessTrue AccountAccessTagListTagsResponseSuccess = true
)

type AccountAccessTagUpdateTagResponse struct {
	Errors   []AccountAccessTagUpdateTagResponseError   `json:"errors"`
	Messages []AccountAccessTagUpdateTagResponseMessage `json:"messages"`
	// A tag
	Result AccountAccessTagUpdateTagResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountAccessTagUpdateTagResponseSuccess `json:"success"`
	JSON    accountAccessTagUpdateTagResponseJSON    `json:"-"`
}

// accountAccessTagUpdateTagResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagUpdateTagResponse]
type accountAccessTagUpdateTagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagUpdateTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagUpdateTagResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountAccessTagUpdateTagResponseErrorJSON `json:"-"`
}

// accountAccessTagUpdateTagResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessTagUpdateTagResponseError]
type accountAccessTagUpdateTagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagUpdateTagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessTagUpdateTagResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAccessTagUpdateTagResponseMessageJSON `json:"-"`
}

// accountAccessTagUpdateTagResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessTagUpdateTagResponseMessage]
type accountAccessTagUpdateTagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagUpdateTagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccountAccessTagUpdateTagResponseResult struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                                       `json:"app_count"`
	CreatedAt time.Time                                   `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      accountAccessTagUpdateTagResponseResultJSON `json:"-"`
}

// accountAccessTagUpdateTagResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessTagUpdateTagResponseResult]
type accountAccessTagUpdateTagResponseResultJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagUpdateTagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessTagUpdateTagResponseSuccess bool

const (
	AccountAccessTagUpdateTagResponseSuccessTrue AccountAccessTagUpdateTagResponseSuccess = true
)

type AccountAccessTagNewTagParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessTagNewTagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessTagUpdateTagParams struct {
	// The name of the tag
	PathName param.Field[string] `path:"name,required"`
	// The name of the tag
	BodyName param.Field[string] `json:"name,required"`
}

func (r AccountAccessTagUpdateTagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
