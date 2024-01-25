// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessBookmarkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessBookmarkService]
// method instead.
type AccountAccessBookmarkService struct {
	Options []option.RequestOption
}

// NewAccountAccessBookmarkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessBookmarkService(opts ...option.RequestOption) (r *AccountAccessBookmarkService) {
	r = &AccountAccessBookmarkService{}
	r.Options = opts
	return
}

// Create a new Bookmark application.
func (r *AccountAccessBookmarkService) New(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountAccessBookmarkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Fetches a single Bookmark application.
func (r *AccountAccessBookmarkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountAccessBookmarkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Bookmark application.
func (r *AccountAccessBookmarkService) Update(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountAccessBookmarkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Deletes a Bookmark application.
func (r *AccountAccessBookmarkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountAccessBookmarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists Bookmark applications.
func (r *AccountAccessBookmarkService) AccessBookmarkApplicationsDeprecatedListBookmarkApplications(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessBookmarkNewResponse struct {
	Errors   []AccountAccessBookmarkNewResponseError   `json:"errors"`
	Messages []AccountAccessBookmarkNewResponseMessage `json:"messages"`
	Result   AccountAccessBookmarkNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessBookmarkNewResponseSuccess `json:"success"`
	JSON    accountAccessBookmarkNewResponseJSON    `json:"-"`
}

// accountAccessBookmarkNewResponseJSON contains the JSON metadata for the struct
// [AccountAccessBookmarkNewResponse]
type accountAccessBookmarkNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkNewResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessBookmarkNewResponseErrorJSON `json:"-"`
}

// accountAccessBookmarkNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkNewResponseError]
type accountAccessBookmarkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkNewResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAccessBookmarkNewResponseMessageJSON `json:"-"`
}

// accountAccessBookmarkNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkNewResponseMessage]
type accountAccessBookmarkNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkNewResponseResult struct {
	// The unique identifier for the Bookmark application.
	ID interface{} `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accountAccessBookmarkNewResponseResultJSON `json:"-"`
}

// accountAccessBookmarkNewResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkNewResponseResult]
type accountAccessBookmarkNewResponseResultJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessBookmarkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessBookmarkNewResponseSuccess bool

const (
	AccountAccessBookmarkNewResponseSuccessTrue AccountAccessBookmarkNewResponseSuccess = true
)

type AccountAccessBookmarkGetResponse struct {
	Errors   []AccountAccessBookmarkGetResponseError   `json:"errors"`
	Messages []AccountAccessBookmarkGetResponseMessage `json:"messages"`
	Result   AccountAccessBookmarkGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessBookmarkGetResponseSuccess `json:"success"`
	JSON    accountAccessBookmarkGetResponseJSON    `json:"-"`
}

// accountAccessBookmarkGetResponseJSON contains the JSON metadata for the struct
// [AccountAccessBookmarkGetResponse]
type accountAccessBookmarkGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessBookmarkGetResponseErrorJSON `json:"-"`
}

// accountAccessBookmarkGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkGetResponseError]
type accountAccessBookmarkGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAccessBookmarkGetResponseMessageJSON `json:"-"`
}

// accountAccessBookmarkGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkGetResponseMessage]
type accountAccessBookmarkGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkGetResponseResult struct {
	// The unique identifier for the Bookmark application.
	ID interface{} `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accountAccessBookmarkGetResponseResultJSON `json:"-"`
}

// accountAccessBookmarkGetResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkGetResponseResult]
type accountAccessBookmarkGetResponseResultJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessBookmarkGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessBookmarkGetResponseSuccess bool

const (
	AccountAccessBookmarkGetResponseSuccessTrue AccountAccessBookmarkGetResponseSuccess = true
)

type AccountAccessBookmarkUpdateResponse struct {
	Errors   []AccountAccessBookmarkUpdateResponseError   `json:"errors"`
	Messages []AccountAccessBookmarkUpdateResponseMessage `json:"messages"`
	Result   AccountAccessBookmarkUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessBookmarkUpdateResponseSuccess `json:"success"`
	JSON    accountAccessBookmarkUpdateResponseJSON    `json:"-"`
}

// accountAccessBookmarkUpdateResponseJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkUpdateResponse]
type accountAccessBookmarkUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAccessBookmarkUpdateResponseErrorJSON `json:"-"`
}

// accountAccessBookmarkUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkUpdateResponseError]
type accountAccessBookmarkUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAccessBookmarkUpdateResponseMessageJSON `json:"-"`
}

// accountAccessBookmarkUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessBookmarkUpdateResponseMessage]
type accountAccessBookmarkUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkUpdateResponseResult struct {
	// The unique identifier for the Bookmark application.
	ID interface{} `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string                                        `json:"name"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      accountAccessBookmarkUpdateResponseResultJSON `json:"-"`
}

// accountAccessBookmarkUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkUpdateResponseResult]
type accountAccessBookmarkUpdateResponseResultJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessBookmarkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessBookmarkUpdateResponseSuccess bool

const (
	AccountAccessBookmarkUpdateResponseSuccessTrue AccountAccessBookmarkUpdateResponseSuccess = true
)

type AccountAccessBookmarkDeleteResponse struct {
	Errors   []AccountAccessBookmarkDeleteResponseError   `json:"errors"`
	Messages []AccountAccessBookmarkDeleteResponseMessage `json:"messages"`
	Result   AccountAccessBookmarkDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessBookmarkDeleteResponseSuccess `json:"success"`
	JSON    accountAccessBookmarkDeleteResponseJSON    `json:"-"`
}

// accountAccessBookmarkDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkDeleteResponse]
type accountAccessBookmarkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAccessBookmarkDeleteResponseErrorJSON `json:"-"`
}

// accountAccessBookmarkDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkDeleteResponseError]
type accountAccessBookmarkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAccessBookmarkDeleteResponseMessageJSON `json:"-"`
}

// accountAccessBookmarkDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessBookmarkDeleteResponseMessage]
type accountAccessBookmarkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkDeleteResponseResult struct {
	// UUID
	ID   string                                        `json:"id"`
	JSON accountAccessBookmarkDeleteResponseResultJSON `json:"-"`
}

// accountAccessBookmarkDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessBookmarkDeleteResponseResult]
type accountAccessBookmarkDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessBookmarkDeleteResponseSuccess bool

const (
	AccountAccessBookmarkDeleteResponseSuccessTrue AccountAccessBookmarkDeleteResponseSuccess = true
)

type AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse struct {
	Errors     []AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError    `json:"errors"`
	Messages   []AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage  `json:"messages"`
	Result     []AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult   `json:"result"`
	ResultInfo AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccess `json:"success"`
	JSON    accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON    `json:"-"`
}

// accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse]
type accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseErrorJSON `json:"-"`
}

// accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError]
type accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessageJSON `json:"-"`
}

// accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage]
type accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult struct {
	// The unique identifier for the Bookmark application.
	ID interface{} `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string                                                                                              `json:"name"`
	UpdatedAt time.Time                                                                                           `json:"updated_at" format:"date-time"`
	JSON      accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultJSON `json:"-"`
}

// accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult]
type accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                 `json:"total_count"`
	JSON       accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfoJSON `json:"-"`
}

// accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo]
type accountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccess bool

const (
	AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccessTrue AccountAccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccess = true
)
