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

// Fetches a single Bookmark application.
func (r *AccountAccessBookmarkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *BookmarksSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Bookmark application.
func (r *AccountAccessBookmarkService) Update(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *BookmarksSingleResponse, err error) {
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
func (r *AccountAccessBookmarkService) AccessBookmarkApplicationsDeprecatedListBookmarkApplications(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *BookmarksResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/bookmarks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BookmarksResponseCollection struct {
	Errors     []BookmarksResponseCollectionError    `json:"errors"`
	Messages   []BookmarksResponseCollectionMessage  `json:"messages"`
	Result     []BookmarksResponseCollectionResult   `json:"result"`
	ResultInfo BookmarksResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success BookmarksResponseCollectionSuccess `json:"success"`
	JSON    bookmarksResponseCollectionJSON    `json:"-"`
}

// bookmarksResponseCollectionJSON contains the JSON metadata for the struct
// [BookmarksResponseCollection]
type bookmarksResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksResponseCollectionError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    bookmarksResponseCollectionErrorJSON `json:"-"`
}

// bookmarksResponseCollectionErrorJSON contains the JSON metadata for the struct
// [BookmarksResponseCollectionError]
type bookmarksResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksResponseCollectionMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    bookmarksResponseCollectionMessageJSON `json:"-"`
}

// bookmarksResponseCollectionMessageJSON contains the JSON metadata for the struct
// [BookmarksResponseCollectionMessage]
type bookmarksResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksResponseCollectionResult struct {
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
	Name      string                                `json:"name"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      bookmarksResponseCollectionResultJSON `json:"-"`
}

// bookmarksResponseCollectionResultJSON contains the JSON metadata for the struct
// [BookmarksResponseCollectionResult]
type bookmarksResponseCollectionResultJSON struct {
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

func (r *BookmarksResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       bookmarksResponseCollectionResultInfoJSON `json:"-"`
}

// bookmarksResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [BookmarksResponseCollectionResultInfo]
type bookmarksResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BookmarksResponseCollectionSuccess bool

const (
	BookmarksResponseCollectionSuccessTrue BookmarksResponseCollectionSuccess = true
)

type BookmarksSingleResponse struct {
	Errors   []BookmarksSingleResponseError   `json:"errors"`
	Messages []BookmarksSingleResponseMessage `json:"messages"`
	Result   BookmarksSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success BookmarksSingleResponseSuccess `json:"success"`
	JSON    bookmarksSingleResponseJSON    `json:"-"`
}

// bookmarksSingleResponseJSON contains the JSON metadata for the struct
// [BookmarksSingleResponse]
type bookmarksSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksSingleResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    bookmarksSingleResponseErrorJSON `json:"-"`
}

// bookmarksSingleResponseErrorJSON contains the JSON metadata for the struct
// [BookmarksSingleResponseError]
type bookmarksSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksSingleResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    bookmarksSingleResponseMessageJSON `json:"-"`
}

// bookmarksSingleResponseMessageJSON contains the JSON metadata for the struct
// [BookmarksSingleResponseMessage]
type bookmarksSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookmarksSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookmarksSingleResponseResult struct {
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
	Name      string                            `json:"name"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      bookmarksSingleResponseResultJSON `json:"-"`
}

// bookmarksSingleResponseResultJSON contains the JSON metadata for the struct
// [BookmarksSingleResponseResult]
type bookmarksSingleResponseResultJSON struct {
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

func (r *BookmarksSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BookmarksSingleResponseSuccess bool

const (
	BookmarksSingleResponseSuccessTrue BookmarksSingleResponseSuccess = true
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
