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

// AccessCustomPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessCustomPageService] method
// instead.
type AccessCustomPageService struct {
	Options []option.RequestOption
}

// NewAccessCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessCustomPageService(opts ...option.RequestOption) (r *AccessCustomPageService) {
	r = &AccessCustomPageService{}
	r.Options = opts
	return
}

// Create a custom page
func (r *AccessCustomPageService) New(ctx context.Context, identifier string, body AccessCustomPageNewParams, opts ...option.RequestOption) (res *AccessCustomPageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a custom page and also returns its HTML.
func (r *AccessCustomPageService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessCustomPageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a custom page
func (r *AccessCustomPageService) Update(ctx context.Context, identifier string, uuid string, body AccessCustomPageUpdateParams, opts ...option.RequestOption) (res *AccessCustomPageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List custom pages
func (r *AccessCustomPageService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessCustomPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a custom page
func (r *AccessCustomPageService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessCustomPageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessCustomPageNewResponse struct {
	Errors   []AccessCustomPageNewResponseError   `json:"errors"`
	Messages []AccessCustomPageNewResponseMessage `json:"messages"`
	Result   AccessCustomPageNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageNewResponseSuccess `json:"success"`
	JSON    accessCustomPageNewResponseJSON    `json:"-"`
}

// accessCustomPageNewResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageNewResponse]
type accessCustomPageNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageNewResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accessCustomPageNewResponseErrorJSON `json:"-"`
}

// accessCustomPageNewResponseErrorJSON contains the JSON metadata for the struct
// [AccessCustomPageNewResponseError]
type accessCustomPageNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageNewResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessCustomPageNewResponseMessageJSON `json:"-"`
}

// accessCustomPageNewResponseMessageJSON contains the JSON metadata for the struct
// [AccessCustomPageNewResponseMessage]
type accessCustomPageNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageNewResponseResult struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageNewResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                `json:"uid"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageNewResponseResultJSON `json:"-"`
}

// accessCustomPageNewResponseResultJSON contains the JSON metadata for the struct
// [AccessCustomPageNewResponseResult]
type accessCustomPageNewResponseResultJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageNewResponseResultType string

const (
	AccessCustomPageNewResponseResultTypeIdentityDenied AccessCustomPageNewResponseResultType = "identity_denied"
	AccessCustomPageNewResponseResultTypeForbidden      AccessCustomPageNewResponseResultType = "forbidden"
)

// Whether the API call was successful
type AccessCustomPageNewResponseSuccess bool

const (
	AccessCustomPageNewResponseSuccessTrue AccessCustomPageNewResponseSuccess = true
)

type AccessCustomPageGetResponse struct {
	Errors   []AccessCustomPageGetResponseError   `json:"errors"`
	Messages []AccessCustomPageGetResponseMessage `json:"messages"`
	Result   AccessCustomPageGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageGetResponseSuccess `json:"success"`
	JSON    accessCustomPageGetResponseJSON    `json:"-"`
}

// accessCustomPageGetResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageGetResponse]
type accessCustomPageGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accessCustomPageGetResponseErrorJSON `json:"-"`
}

// accessCustomPageGetResponseErrorJSON contains the JSON metadata for the struct
// [AccessCustomPageGetResponseError]
type accessCustomPageGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessCustomPageGetResponseMessageJSON `json:"-"`
}

// accessCustomPageGetResponseMessageJSON contains the JSON metadata for the struct
// [AccessCustomPageGetResponseMessage]
type accessCustomPageGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageGetResponseResult struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageGetResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                `json:"uid"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageGetResponseResultJSON `json:"-"`
}

// accessCustomPageGetResponseResultJSON contains the JSON metadata for the struct
// [AccessCustomPageGetResponseResult]
type accessCustomPageGetResponseResultJSON struct {
	CustomHTML  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageGetResponseResultType string

const (
	AccessCustomPageGetResponseResultTypeIdentityDenied AccessCustomPageGetResponseResultType = "identity_denied"
	AccessCustomPageGetResponseResultTypeForbidden      AccessCustomPageGetResponseResultType = "forbidden"
)

// Whether the API call was successful
type AccessCustomPageGetResponseSuccess bool

const (
	AccessCustomPageGetResponseSuccessTrue AccessCustomPageGetResponseSuccess = true
)

type AccessCustomPageUpdateResponse struct {
	Errors   []AccessCustomPageUpdateResponseError   `json:"errors"`
	Messages []AccessCustomPageUpdateResponseMessage `json:"messages"`
	Result   AccessCustomPageUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageUpdateResponseSuccess `json:"success"`
	JSON    accessCustomPageUpdateResponseJSON    `json:"-"`
}

// accessCustomPageUpdateResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageUpdateResponse]
type accessCustomPageUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessCustomPageUpdateResponseErrorJSON `json:"-"`
}

// accessCustomPageUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccessCustomPageUpdateResponseError]
type accessCustomPageUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessCustomPageUpdateResponseMessageJSON `json:"-"`
}

// accessCustomPageUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccessCustomPageUpdateResponseMessage]
type accessCustomPageUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageUpdateResponseResult struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageUpdateResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                   `json:"uid"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageUpdateResponseResultJSON `json:"-"`
}

// accessCustomPageUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccessCustomPageUpdateResponseResult]
type accessCustomPageUpdateResponseResultJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageUpdateResponseResultType string

const (
	AccessCustomPageUpdateResponseResultTypeIdentityDenied AccessCustomPageUpdateResponseResultType = "identity_denied"
	AccessCustomPageUpdateResponseResultTypeForbidden      AccessCustomPageUpdateResponseResultType = "forbidden"
)

// Whether the API call was successful
type AccessCustomPageUpdateResponseSuccess bool

const (
	AccessCustomPageUpdateResponseSuccessTrue AccessCustomPageUpdateResponseSuccess = true
)

type AccessCustomPageListResponse struct {
	Errors     []AccessCustomPageListResponseError    `json:"errors"`
	Messages   []AccessCustomPageListResponseMessage  `json:"messages"`
	Result     []AccessCustomPageListResponseResult   `json:"result"`
	ResultInfo AccessCustomPageListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessCustomPageListResponseSuccess `json:"success"`
	JSON    accessCustomPageListResponseJSON    `json:"-"`
}

// accessCustomPageListResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageListResponse]
type accessCustomPageListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accessCustomPageListResponseErrorJSON `json:"-"`
}

// accessCustomPageListResponseErrorJSON contains the JSON metadata for the struct
// [AccessCustomPageListResponseError]
type accessCustomPageListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessCustomPageListResponseMessageJSON `json:"-"`
}

// accessCustomPageListResponseMessageJSON contains the JSON metadata for the
// struct [AccessCustomPageListResponseMessage]
type accessCustomPageListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageListResponseResult struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageListResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                 `json:"uid"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageListResponseResultJSON `json:"-"`
}

// accessCustomPageListResponseResultJSON contains the JSON metadata for the struct
// [AccessCustomPageListResponseResult]
type accessCustomPageListResponseResultJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageListResponseResultType string

const (
	AccessCustomPageListResponseResultTypeIdentityDenied AccessCustomPageListResponseResultType = "identity_denied"
	AccessCustomPageListResponseResultTypeForbidden      AccessCustomPageListResponseResultType = "forbidden"
)

type AccessCustomPageListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       accessCustomPageListResponseResultInfoJSON `json:"-"`
}

// accessCustomPageListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccessCustomPageListResponseResultInfo]
type accessCustomPageListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageListResponseSuccess bool

const (
	AccessCustomPageListResponseSuccessTrue AccessCustomPageListResponseSuccess = true
)

type AccessCustomPageDeleteResponse struct {
	Errors   []AccessCustomPageDeleteResponseError   `json:"errors"`
	Messages []AccessCustomPageDeleteResponseMessage `json:"messages"`
	Result   AccessCustomPageDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageDeleteResponseSuccess `json:"success"`
	JSON    accessCustomPageDeleteResponseJSON    `json:"-"`
}

// accessCustomPageDeleteResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageDeleteResponse]
type accessCustomPageDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessCustomPageDeleteResponseErrorJSON `json:"-"`
}

// accessCustomPageDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccessCustomPageDeleteResponseError]
type accessCustomPageDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessCustomPageDeleteResponseMessageJSON `json:"-"`
}

// accessCustomPageDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccessCustomPageDeleteResponseMessage]
type accessCustomPageDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageDeleteResponseResult struct {
	// UUID
	ID   string                                   `json:"id"`
	JSON accessCustomPageDeleteResponseResultJSON `json:"-"`
}

// accessCustomPageDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccessCustomPageDeleteResponseResult]
type accessCustomPageDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageDeleteResponseSuccess bool

const (
	AccessCustomPageDeleteResponseSuccessTrue AccessCustomPageDeleteResponseSuccess = true
)

type AccessCustomPageNewParams struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[AccessCustomPageNewParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r AccessCustomPageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type AccessCustomPageNewParamsType string

const (
	AccessCustomPageNewParamsTypeIdentityDenied AccessCustomPageNewParamsType = "identity_denied"
	AccessCustomPageNewParamsTypeForbidden      AccessCustomPageNewParamsType = "forbidden"
)

type AccessCustomPageUpdateParams struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[AccessCustomPageUpdateParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r AccessCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type AccessCustomPageUpdateParamsType string

const (
	AccessCustomPageUpdateParamsTypeIdentityDenied AccessCustomPageUpdateParamsType = "identity_denied"
	AccessCustomPageUpdateParamsTypeForbidden      AccessCustomPageUpdateParamsType = "forbidden"
)
