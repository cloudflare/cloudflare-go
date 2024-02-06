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

// AccessBookmarkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessBookmarkService] method
// instead.
type AccessBookmarkService struct {
	Options []option.RequestOption
}

// NewAccessBookmarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessBookmarkService(opts ...option.RequestOption) (r *AccessBookmarkService) {
	r = &AccessBookmarkService{}
	r.Options = opts
	return
}

// Fetches a single Bookmark application.
func (r *AccessBookmarkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Bookmark application.
func (r *AccessBookmarkService) Update(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Deletes a Bookmark application.
func (r *AccessBookmarkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists Bookmark applications.
func (r *AccessBookmarkService) AccessBookmarkApplicationsDeprecatedListBookmarkApplications(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessBookmarkGetResponse struct {
	Errors   []AccessBookmarkGetResponseError   `json:"errors"`
	Messages []AccessBookmarkGetResponseMessage `json:"messages"`
	Result   AccessBookmarkGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessBookmarkGetResponseSuccess `json:"success"`
	JSON    accessBookmarkGetResponseJSON    `json:"-"`
}

// accessBookmarkGetResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponse]
type accessBookmarkGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkGetResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accessBookmarkGetResponseErrorJSON `json:"-"`
}

// accessBookmarkGetResponseErrorJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponseError]
type accessBookmarkGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkGetResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accessBookmarkGetResponseMessageJSON `json:"-"`
}

// accessBookmarkGetResponseMessageJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponseMessage]
type accessBookmarkGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkGetResponseResult struct {
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
	Name      string                              `json:"name"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkGetResponseResultJSON `json:"-"`
}

// accessBookmarkGetResponseResultJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponseResult]
type accessBookmarkGetResponseResultJSON struct {
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

func (r *AccessBookmarkGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessBookmarkGetResponseSuccess bool

const (
	AccessBookmarkGetResponseSuccessTrue AccessBookmarkGetResponseSuccess = true
)

type AccessBookmarkUpdateResponse struct {
	Errors   []AccessBookmarkUpdateResponseError   `json:"errors"`
	Messages []AccessBookmarkUpdateResponseMessage `json:"messages"`
	Result   AccessBookmarkUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessBookmarkUpdateResponseSuccess `json:"success"`
	JSON    accessBookmarkUpdateResponseJSON    `json:"-"`
}

// accessBookmarkUpdateResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkUpdateResponse]
type accessBookmarkUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accessBookmarkUpdateResponseErrorJSON `json:"-"`
}

// accessBookmarkUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccessBookmarkUpdateResponseError]
type accessBookmarkUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessBookmarkUpdateResponseMessageJSON `json:"-"`
}

// accessBookmarkUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccessBookmarkUpdateResponseMessage]
type accessBookmarkUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkUpdateResponseResult struct {
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
	Name      string                                 `json:"name"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkUpdateResponseResultJSON `json:"-"`
}

// accessBookmarkUpdateResponseResultJSON contains the JSON metadata for the struct
// [AccessBookmarkUpdateResponseResult]
type accessBookmarkUpdateResponseResultJSON struct {
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

func (r *AccessBookmarkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessBookmarkUpdateResponseSuccess bool

const (
	AccessBookmarkUpdateResponseSuccessTrue AccessBookmarkUpdateResponseSuccess = true
)

type AccessBookmarkDeleteResponse struct {
	Errors   []AccessBookmarkDeleteResponseError   `json:"errors"`
	Messages []AccessBookmarkDeleteResponseMessage `json:"messages"`
	Result   AccessBookmarkDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessBookmarkDeleteResponseSuccess `json:"success"`
	JSON    accessBookmarkDeleteResponseJSON    `json:"-"`
}

// accessBookmarkDeleteResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkDeleteResponse]
type accessBookmarkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkDeleteResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accessBookmarkDeleteResponseErrorJSON `json:"-"`
}

// accessBookmarkDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccessBookmarkDeleteResponseError]
type accessBookmarkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkDeleteResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessBookmarkDeleteResponseMessageJSON `json:"-"`
}

// accessBookmarkDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccessBookmarkDeleteResponseMessage]
type accessBookmarkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkDeleteResponseResult struct {
	// UUID
	ID   string                                 `json:"id"`
	JSON accessBookmarkDeleteResponseResultJSON `json:"-"`
}

// accessBookmarkDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccessBookmarkDeleteResponseResult]
type accessBookmarkDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessBookmarkDeleteResponseSuccess bool

const (
	AccessBookmarkDeleteResponseSuccessTrue AccessBookmarkDeleteResponseSuccess = true
)

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse struct {
	Errors     []AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError    `json:"errors"`
	Messages   []AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage  `json:"messages"`
	Result     []AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult   `json:"result"`
	ResultInfo AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccess `json:"success"`
	JSON    accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON    `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseErrorJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessageJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult struct {
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
	Name      string                                                                                       `json:"name"`
	UpdatedAt time.Time                                                                                    `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultJSON struct {
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

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                          `json:"total_count"`
	JSON       accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfoJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccess bool

const (
	AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccessTrue AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseSuccess = true
)
