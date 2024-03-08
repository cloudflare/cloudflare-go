// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Create a new Bookmark application.
func (r *AccessBookmarkService) New(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Bookmark application.
func (r *AccessBookmarkService) Update(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Bookmark applications.
func (r *AccessBookmarkService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Bookmark application.
func (r *AccessBookmarkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Bookmark application.
func (r *AccessBookmarkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessBookmarks struct {
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
	Name      string              `json:"name"`
	UpdatedAt time.Time           `json:"updated_at" format:"date-time"`
	JSON      accessBookmarksJSON `json:"-"`
}

// accessBookmarksJSON contains the JSON metadata for the struct [AccessBookmarks]
type accessBookmarksJSON struct {
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

func (r *AccessBookmarks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarksJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponse struct {
	// UUID
	ID   string                           `json:"id"`
	JSON accessBookmarkDeleteResponseJSON `json:"-"`
}

// accessBookmarkDeleteResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkDeleteResponse]
type accessBookmarkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelope struct {
	Errors   []AccessBookmarkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarks                             `json:"result,required"`
	// Whether the API call was successful
	Success AccessBookmarkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessBookmarkNewResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessBookmarkNewResponseEnvelope]
type accessBookmarkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessBookmarkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessBookmarkNewResponseEnvelopeErrors]
type accessBookmarkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessBookmarkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessBookmarkNewResponseEnvelopeMessages]
type accessBookmarkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessBookmarkNewResponseEnvelopeSuccess bool

const (
	AccessBookmarkNewResponseEnvelopeSuccessTrue AccessBookmarkNewResponseEnvelopeSuccess = true
)

type AccessBookmarkUpdateResponseEnvelope struct {
	Errors   []AccessBookmarkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarks                                `json:"result,required"`
	// Whether the API call was successful
	Success AccessBookmarkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessBookmarkUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessBookmarkUpdateResponseEnvelope]
type accessBookmarkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessBookmarkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessBookmarkUpdateResponseEnvelopeErrors]
type accessBookmarkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessBookmarkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessBookmarkUpdateResponseEnvelopeMessages]
type accessBookmarkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessBookmarkUpdateResponseEnvelopeSuccess bool

const (
	AccessBookmarkUpdateResponseEnvelopeSuccessTrue AccessBookmarkUpdateResponseEnvelopeSuccess = true
)

type AccessBookmarkListResponseEnvelope struct {
	Errors   []AccessBookmarkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessBookmarks                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessBookmarkListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessBookmarkListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessBookmarkListResponseEnvelopeJSON       `json:"-"`
}

// accessBookmarkListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessBookmarkListResponseEnvelope]
type accessBookmarkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessBookmarkListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessBookmarkListResponseEnvelopeErrors]
type accessBookmarkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessBookmarkListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessBookmarkListResponseEnvelopeMessages]
type accessBookmarkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessBookmarkListResponseEnvelopeSuccess bool

const (
	AccessBookmarkListResponseEnvelopeSuccessTrue AccessBookmarkListResponseEnvelopeSuccess = true
)

type AccessBookmarkListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accessBookmarkListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessBookmarkListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AccessBookmarkListResponseEnvelopeResultInfo]
type accessBookmarkListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelope struct {
	Errors   []AccessBookmarkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessBookmarkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessBookmarkDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessBookmarkDeleteResponseEnvelope]
type accessBookmarkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessBookmarkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessBookmarkDeleteResponseEnvelopeErrors]
type accessBookmarkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessBookmarkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessBookmarkDeleteResponseEnvelopeMessages]
type accessBookmarkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessBookmarkDeleteResponseEnvelopeSuccess bool

const (
	AccessBookmarkDeleteResponseEnvelopeSuccessTrue AccessBookmarkDeleteResponseEnvelopeSuccess = true
)

type AccessBookmarkGetResponseEnvelope struct {
	Errors   []AccessBookmarkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarks                             `json:"result,required"`
	// Whether the API call was successful
	Success AccessBookmarkGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessBookmarkGetResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponseEnvelope]
type accessBookmarkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessBookmarkGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessBookmarkGetResponseEnvelopeErrors]
type accessBookmarkGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessBookmarkGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessBookmarkGetResponseEnvelopeMessages]
type accessBookmarkGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessBookmarkGetResponseEnvelopeSuccess bool

const (
	AccessBookmarkGetResponseEnvelopeSuccessTrue AccessBookmarkGetResponseEnvelopeSuccess = true
)
