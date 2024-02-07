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
func (r *AccessBookmarkService) Update(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarkUpdateResponse, err error) {
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

// Lists Bookmark applications.
func (r *AccessBookmarkService) AccessBookmarkApplicationsDeprecatedListBookmarkApplications(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type AccessBookmarkUpdateResponse struct {
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
	Name      string                           `json:"name"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkUpdateResponseJSON `json:"-"`
}

// accessBookmarkUpdateResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkUpdateResponse]
type accessBookmarkUpdateResponseJSON struct {
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

func (r *AccessBookmarkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
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

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse struct {
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
	Name      string                                                                                 `json:"name"`
	UpdatedAt time.Time                                                                              `json:"updated_at" format:"date-time"`
	JSON      accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseJSON struct {
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

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkGetResponseEnvelope struct {
	Errors   []AccessBookmarkGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessBookmarkGetResponseEnvelopeMessages `json:"messages"`
	Result   AccessBookmarkGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessBookmarkGetResponseEnvelopeSuccess `json:"success"`
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

type AccessBookmarkUpdateResponseEnvelope struct {
	Errors   []AccessBookmarkUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessBookmarkUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AccessBookmarkUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessBookmarkUpdateResponseEnvelopeSuccess `json:"success"`
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

// Whether the API call was successful
type AccessBookmarkUpdateResponseEnvelopeSuccess bool

const (
	AccessBookmarkUpdateResponseEnvelopeSuccessTrue AccessBookmarkUpdateResponseEnvelopeSuccess = true
)

type AccessBookmarkDeleteResponseEnvelope struct {
	Errors   []AccessBookmarkDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessBookmarkDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccessBookmarkDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessBookmarkDeleteResponseEnvelopeSuccess `json:"success"`
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

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelope struct {
	Errors     []AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessages `json:"messages"`
	Result     []AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponse                 `json:"result"`
	ResultInfo AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeSuccess `json:"success"`
	JSON    accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelope]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrors struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrorsJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrors]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessages struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessagesJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessages]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                  `json:"total_count"`
	JSON       accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfo]
type accessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeSuccess bool

const (
	AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeSuccessTrue AccessBookmarkAccessBookmarkApplicationsDeprecatedListBookmarkApplicationsResponseEnvelopeSuccess = true
)
