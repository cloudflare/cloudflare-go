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

// AccountAccessCustomPageService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAccessCustomPageService] method instead.
type AccountAccessCustomPageService struct {
	Options []option.RequestOption
}

// NewAccountAccessCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessCustomPageService(opts ...option.RequestOption) (r *AccountAccessCustomPageService) {
	r = &AccountAccessCustomPageService{}
	r.Options = opts
	return
}

// Create a custom page
func (r *AccountAccessCustomPageService) New(ctx context.Context, identifier string, body AccountAccessCustomPageNewParams, opts ...option.RequestOption) (res *AccountAccessCustomPageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List custom pages
func (r *AccountAccessCustomPageService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessCustomPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a custom page
func (r *AccountAccessCustomPageService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessCustomPageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetches a custom page and also returns its HTML.
func (r *AccountAccessCustomPageService) GetCustomPage(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessCustomPageGetCustomPageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a custom page
func (r *AccountAccessCustomPageService) UpdateCustomPage(ctx context.Context, identifier string, uuid string, body AccountAccessCustomPageUpdateCustomPageParams, opts ...option.RequestOption) (res *AccountAccessCustomPageUpdateCustomPageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountAccessCustomPageNewResponse struct {
	Errors   []AccountAccessCustomPageNewResponseError   `json:"errors"`
	Messages []AccountAccessCustomPageNewResponseMessage `json:"messages"`
	Result   AccountAccessCustomPageNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCustomPageNewResponseSuccess `json:"success"`
	JSON    accountAccessCustomPageNewResponseJSON    `json:"-"`
}

// accountAccessCustomPageNewResponseJSON contains the JSON metadata for the struct
// [AccountAccessCustomPageNewResponse]
type accountAccessCustomPageNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageNewResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAccessCustomPageNewResponseErrorJSON `json:"-"`
}

// accountAccessCustomPageNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageNewResponseError]
type accountAccessCustomPageNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageNewResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountAccessCustomPageNewResponseMessageJSON `json:"-"`
}

// accountAccessCustomPageNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageNewResponseMessage]
type accountAccessCustomPageNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageNewResponseResult struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccountAccessCustomPageNewResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                       `json:"uid"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      accountAccessCustomPageNewResponseResultJSON `json:"-"`
}

// accountAccessCustomPageNewResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageNewResponseResult]
type accountAccessCustomPageNewResponseResultJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccountAccessCustomPageNewResponseResultType string

const (
	AccountAccessCustomPageNewResponseResultTypeIdentityDenied AccountAccessCustomPageNewResponseResultType = "identity_denied"
	AccountAccessCustomPageNewResponseResultTypeForbidden      AccountAccessCustomPageNewResponseResultType = "forbidden"
)

// Whether the API call was successful
type AccountAccessCustomPageNewResponseSuccess bool

const (
	AccountAccessCustomPageNewResponseSuccessTrue AccountAccessCustomPageNewResponseSuccess = true
)

type AccountAccessCustomPageListResponse struct {
	Errors     []AccountAccessCustomPageListResponseError    `json:"errors"`
	Messages   []AccountAccessCustomPageListResponseMessage  `json:"messages"`
	Result     []AccountAccessCustomPageListResponseResult   `json:"result"`
	ResultInfo AccountAccessCustomPageListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessCustomPageListResponseSuccess `json:"success"`
	JSON    accountAccessCustomPageListResponseJSON    `json:"-"`
}

// accountAccessCustomPageListResponseJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageListResponse]
type accountAccessCustomPageListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAccessCustomPageListResponseErrorJSON `json:"-"`
}

// accountAccessCustomPageListResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageListResponseError]
type accountAccessCustomPageListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAccessCustomPageListResponseMessageJSON `json:"-"`
}

// accountAccessCustomPageListResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessCustomPageListResponseMessage]
type accountAccessCustomPageListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageListResponseResult struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccountAccessCustomPageListResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                        `json:"uid"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      accountAccessCustomPageListResponseResultJSON `json:"-"`
}

// accountAccessCustomPageListResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageListResponseResult]
type accountAccessCustomPageListResponseResultJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccountAccessCustomPageListResponseResultType string

const (
	AccountAccessCustomPageListResponseResultTypeIdentityDenied AccountAccessCustomPageListResponseResultType = "identity_denied"
	AccountAccessCustomPageListResponseResultTypeForbidden      AccountAccessCustomPageListResponseResultType = "forbidden"
)

type AccountAccessCustomPageListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       accountAccessCustomPageListResponseResultInfoJSON `json:"-"`
}

// accountAccessCustomPageListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAccessCustomPageListResponseResultInfo]
type accountAccessCustomPageListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCustomPageListResponseSuccess bool

const (
	AccountAccessCustomPageListResponseSuccessTrue AccountAccessCustomPageListResponseSuccess = true
)

type AccountAccessCustomPageDeleteResponse struct {
	Errors   []AccountAccessCustomPageDeleteResponseError   `json:"errors"`
	Messages []AccountAccessCustomPageDeleteResponseMessage `json:"messages"`
	Result   AccountAccessCustomPageDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCustomPageDeleteResponseSuccess `json:"success"`
	JSON    accountAccessCustomPageDeleteResponseJSON    `json:"-"`
}

// accountAccessCustomPageDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageDeleteResponse]
type accountAccessCustomPageDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAccessCustomPageDeleteResponseErrorJSON `json:"-"`
}

// accountAccessCustomPageDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessCustomPageDeleteResponseError]
type accountAccessCustomPageDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountAccessCustomPageDeleteResponseMessageJSON `json:"-"`
}

// accountAccessCustomPageDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessCustomPageDeleteResponseMessage]
type accountAccessCustomPageDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageDeleteResponseResult struct {
	// UUID
	ID   string                                          `json:"id"`
	JSON accountAccessCustomPageDeleteResponseResultJSON `json:"-"`
}

// accountAccessCustomPageDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessCustomPageDeleteResponseResult]
type accountAccessCustomPageDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCustomPageDeleteResponseSuccess bool

const (
	AccountAccessCustomPageDeleteResponseSuccessTrue AccountAccessCustomPageDeleteResponseSuccess = true
)

type AccountAccessCustomPageGetCustomPageResponse struct {
	Errors   []AccountAccessCustomPageGetCustomPageResponseError   `json:"errors"`
	Messages []AccountAccessCustomPageGetCustomPageResponseMessage `json:"messages"`
	Result   AccountAccessCustomPageGetCustomPageResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCustomPageGetCustomPageResponseSuccess `json:"success"`
	JSON    accountAccessCustomPageGetCustomPageResponseJSON    `json:"-"`
}

// accountAccessCustomPageGetCustomPageResponseJSON contains the JSON metadata for
// the struct [AccountAccessCustomPageGetCustomPageResponse]
type accountAccessCustomPageGetCustomPageResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageGetCustomPageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageGetCustomPageResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAccessCustomPageGetCustomPageResponseErrorJSON `json:"-"`
}

// accountAccessCustomPageGetCustomPageResponseErrorJSON contains the JSON metadata
// for the struct [AccountAccessCustomPageGetCustomPageResponseError]
type accountAccessCustomPageGetCustomPageResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageGetCustomPageResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageGetCustomPageResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountAccessCustomPageGetCustomPageResponseMessageJSON `json:"-"`
}

// accountAccessCustomPageGetCustomPageResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessCustomPageGetCustomPageResponseMessage]
type accountAccessCustomPageGetCustomPageResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageGetCustomPageResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageGetCustomPageResponseResult struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccountAccessCustomPageGetCustomPageResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                                 `json:"uid"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      accountAccessCustomPageGetCustomPageResponseResultJSON `json:"-"`
}

// accountAccessCustomPageGetCustomPageResponseResultJSON contains the JSON
// metadata for the struct [AccountAccessCustomPageGetCustomPageResponseResult]
type accountAccessCustomPageGetCustomPageResponseResultJSON struct {
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

func (r *AccountAccessCustomPageGetCustomPageResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccountAccessCustomPageGetCustomPageResponseResultType string

const (
	AccountAccessCustomPageGetCustomPageResponseResultTypeIdentityDenied AccountAccessCustomPageGetCustomPageResponseResultType = "identity_denied"
	AccountAccessCustomPageGetCustomPageResponseResultTypeForbidden      AccountAccessCustomPageGetCustomPageResponseResultType = "forbidden"
)

// Whether the API call was successful
type AccountAccessCustomPageGetCustomPageResponseSuccess bool

const (
	AccountAccessCustomPageGetCustomPageResponseSuccessTrue AccountAccessCustomPageGetCustomPageResponseSuccess = true
)

type AccountAccessCustomPageUpdateCustomPageResponse struct {
	Errors   []AccountAccessCustomPageUpdateCustomPageResponseError   `json:"errors"`
	Messages []AccountAccessCustomPageUpdateCustomPageResponseMessage `json:"messages"`
	Result   AccountAccessCustomPageUpdateCustomPageResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCustomPageUpdateCustomPageResponseSuccess `json:"success"`
	JSON    accountAccessCustomPageUpdateCustomPageResponseJSON    `json:"-"`
}

// accountAccessCustomPageUpdateCustomPageResponseJSON contains the JSON metadata
// for the struct [AccountAccessCustomPageUpdateCustomPageResponse]
type accountAccessCustomPageUpdateCustomPageResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageUpdateCustomPageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageUpdateCustomPageResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountAccessCustomPageUpdateCustomPageResponseErrorJSON `json:"-"`
}

// accountAccessCustomPageUpdateCustomPageResponseErrorJSON contains the JSON
// metadata for the struct [AccountAccessCustomPageUpdateCustomPageResponseError]
type accountAccessCustomPageUpdateCustomPageResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageUpdateCustomPageResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageUpdateCustomPageResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountAccessCustomPageUpdateCustomPageResponseMessageJSON `json:"-"`
}

// accountAccessCustomPageUpdateCustomPageResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessCustomPageUpdateCustomPageResponseMessage]
type accountAccessCustomPageUpdateCustomPageResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageUpdateCustomPageResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCustomPageUpdateCustomPageResponseResult struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccountAccessCustomPageUpdateCustomPageResponseResultType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                                    `json:"uid"`
	UpdatedAt time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON      accountAccessCustomPageUpdateCustomPageResponseResultJSON `json:"-"`
}

// accountAccessCustomPageUpdateCustomPageResponseResultJSON contains the JSON
// metadata for the struct [AccountAccessCustomPageUpdateCustomPageResponseResult]
type accountAccessCustomPageUpdateCustomPageResponseResultJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageUpdateCustomPageResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccountAccessCustomPageUpdateCustomPageResponseResultType string

const (
	AccountAccessCustomPageUpdateCustomPageResponseResultTypeIdentityDenied AccountAccessCustomPageUpdateCustomPageResponseResultType = "identity_denied"
	AccountAccessCustomPageUpdateCustomPageResponseResultTypeForbidden      AccountAccessCustomPageUpdateCustomPageResponseResultType = "forbidden"
)

// Whether the API call was successful
type AccountAccessCustomPageUpdateCustomPageResponseSuccess bool

const (
	AccountAccessCustomPageUpdateCustomPageResponseSuccessTrue AccountAccessCustomPageUpdateCustomPageResponseSuccess = true
)

type AccountAccessCustomPageNewParams struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[AccountAccessCustomPageNewParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r AccountAccessCustomPageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type AccountAccessCustomPageNewParamsType string

const (
	AccountAccessCustomPageNewParamsTypeIdentityDenied AccountAccessCustomPageNewParamsType = "identity_denied"
	AccountAccessCustomPageNewParamsTypeForbidden      AccountAccessCustomPageNewParamsType = "forbidden"
)

type AccountAccessCustomPageUpdateCustomPageParams struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[AccountAccessCustomPageUpdateCustomPageParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r AccountAccessCustomPageUpdateCustomPageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type AccountAccessCustomPageUpdateCustomPageParamsType string

const (
	AccountAccessCustomPageUpdateCustomPageParamsTypeIdentityDenied AccountAccessCustomPageUpdateCustomPageParamsType = "identity_denied"
	AccountAccessCustomPageUpdateCustomPageParamsTypeForbidden      AccountAccessCustomPageUpdateCustomPageParamsType = "forbidden"
)
