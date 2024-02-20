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

// Lists Bookmark applications.
func (r *AccessBookmarkService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]AccessBookmarkListResponse, err error) {
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
func (r *AccessBookmarkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkGetResponse, err error) {
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

// Updates a configured Bookmark application.
func (r *AccessBookmarkService) Replace(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessBookmarkListResponse struct {
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
	Name      string                         `json:"name"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkListResponseJSON `json:"-"`
}

// accessBookmarkListResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkListResponse]
type accessBookmarkListResponseJSON struct {
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

func (r *AccessBookmarkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type AccessBookmarkGetResponse struct {
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
	Name      string                        `json:"name"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkGetResponseJSON `json:"-"`
}

// accessBookmarkGetResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponse]
type accessBookmarkGetResponseJSON struct {
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

func (r *AccessBookmarkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkReplaceResponse struct {
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
	JSON      accessBookmarkReplaceResponseJSON `json:"-"`
}

// accessBookmarkReplaceResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkReplaceResponse]
type accessBookmarkReplaceResponseJSON struct {
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

func (r *AccessBookmarkReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkListResponseEnvelope struct {
	Errors   []AccessBookmarkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessBookmarkListResponse                 `json:"result,required,nullable"`
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

// Whether the API call was successful
type AccessBookmarkDeleteResponseEnvelopeSuccess bool

const (
	AccessBookmarkDeleteResponseEnvelopeSuccessTrue AccessBookmarkDeleteResponseEnvelopeSuccess = true
)

type AccessBookmarkGetResponseEnvelope struct {
	Errors   []AccessBookmarkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarkGetResponse                   `json:"result,required"`
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

// Whether the API call was successful
type AccessBookmarkGetResponseEnvelopeSuccess bool

const (
	AccessBookmarkGetResponseEnvelopeSuccessTrue AccessBookmarkGetResponseEnvelopeSuccess = true
)

type AccessBookmarkReplaceResponseEnvelope struct {
	Errors   []AccessBookmarkReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarkReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessBookmarkReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessBookmarkReplaceResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessBookmarkReplaceResponseEnvelope]
type accessBookmarkReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkReplaceResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessBookmarkReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessBookmarkReplaceResponseEnvelopeErrors]
type accessBookmarkReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkReplaceResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessBookmarkReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessBookmarkReplaceResponseEnvelopeMessages]
type accessBookmarkReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessBookmarkReplaceResponseEnvelopeSuccess bool

const (
	AccessBookmarkReplaceResponseEnvelopeSuccessTrue AccessBookmarkReplaceResponseEnvelopeSuccess = true
)
