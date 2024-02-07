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
	var env AccessCustomPageNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom page and also returns its HTML.
func (r *AccessCustomPageService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessCustomPageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a custom page
func (r *AccessCustomPageService) Update(ctx context.Context, identifier string, uuid string, body AccessCustomPageUpdateParams, opts ...option.RequestOption) (res *AccessCustomPageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List custom pages
func (r *AccessCustomPageService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessCustomPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a custom page
func (r *AccessCustomPageService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessCustomPageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessCustomPageNewResponse struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageNewResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                          `json:"uid"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageNewResponseJSON `json:"-"`
}

// accessCustomPageNewResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageNewResponse]
type accessCustomPageNewResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageNewResponseType string

const (
	AccessCustomPageNewResponseTypeIdentityDenied AccessCustomPageNewResponseType = "identity_denied"
	AccessCustomPageNewResponseTypeForbidden      AccessCustomPageNewResponseType = "forbidden"
)

type AccessCustomPageGetResponse struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageGetResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                          `json:"uid"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageGetResponseJSON `json:"-"`
}

// accessCustomPageGetResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageGetResponse]
type accessCustomPageGetResponseJSON struct {
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

func (r *AccessCustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageGetResponseType string

const (
	AccessCustomPageGetResponseTypeIdentityDenied AccessCustomPageGetResponseType = "identity_denied"
	AccessCustomPageGetResponseTypeForbidden      AccessCustomPageGetResponseType = "forbidden"
)

type AccessCustomPageUpdateResponse struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageUpdateResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                             `json:"uid"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageUpdateResponseJSON `json:"-"`
}

// accessCustomPageUpdateResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageUpdateResponse]
type accessCustomPageUpdateResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageUpdateResponseType string

const (
	AccessCustomPageUpdateResponseTypeIdentityDenied AccessCustomPageUpdateResponseType = "identity_denied"
	AccessCustomPageUpdateResponseTypeForbidden      AccessCustomPageUpdateResponseType = "forbidden"
)

type AccessCustomPageListResponse struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type AccessCustomPageListResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                           `json:"uid"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      accessCustomPageListResponseJSON `json:"-"`
}

// accessCustomPageListResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageListResponse]
type accessCustomPageListResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom page type.
type AccessCustomPageListResponseType string

const (
	AccessCustomPageListResponseTypeIdentityDenied AccessCustomPageListResponseType = "identity_denied"
	AccessCustomPageListResponseTypeForbidden      AccessCustomPageListResponseType = "forbidden"
)

type AccessCustomPageDeleteResponse struct {
	// UUID
	ID   string                             `json:"id"`
	JSON accessCustomPageDeleteResponseJSON `json:"-"`
}

// accessCustomPageDeleteResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageDeleteResponse]
type accessCustomPageDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type AccessCustomPageNewResponseEnvelope struct {
	Errors   []AccessCustomPageNewResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessCustomPageNewResponseEnvelopeMessages `json:"messages"`
	Result   AccessCustomPageNewResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageNewResponseEnvelopeSuccess `json:"success"`
	JSON    accessCustomPageNewResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageNewResponseEnvelope]
type accessCustomPageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessCustomPageNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCustomPageNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessCustomPageNewResponseEnvelopeErrors]
type accessCustomPageNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessCustomPageNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCustomPageNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessCustomPageNewResponseEnvelopeMessages]
type accessCustomPageNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageNewResponseEnvelopeSuccess bool

const (
	AccessCustomPageNewResponseEnvelopeSuccessTrue AccessCustomPageNewResponseEnvelopeSuccess = true
)

type AccessCustomPageGetResponseEnvelope struct {
	Errors   []AccessCustomPageGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessCustomPageGetResponseEnvelopeMessages `json:"messages"`
	Result   AccessCustomPageGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageGetResponseEnvelopeSuccess `json:"success"`
	JSON    accessCustomPageGetResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageGetResponseEnvelope]
type accessCustomPageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessCustomPageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCustomPageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessCustomPageGetResponseEnvelopeErrors]
type accessCustomPageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessCustomPageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCustomPageGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessCustomPageGetResponseEnvelopeMessages]
type accessCustomPageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageGetResponseEnvelopeSuccess bool

const (
	AccessCustomPageGetResponseEnvelopeSuccessTrue AccessCustomPageGetResponseEnvelopeSuccess = true
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

type AccessCustomPageUpdateResponseEnvelope struct {
	Errors   []AccessCustomPageUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessCustomPageUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AccessCustomPageUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    accessCustomPageUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageUpdateResponseEnvelope]
type accessCustomPageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessCustomPageUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCustomPageUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCustomPageUpdateResponseEnvelopeErrors]
type accessCustomPageUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessCustomPageUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCustomPageUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessCustomPageUpdateResponseEnvelopeMessages]
type accessCustomPageUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageUpdateResponseEnvelopeSuccess bool

const (
	AccessCustomPageUpdateResponseEnvelopeSuccessTrue AccessCustomPageUpdateResponseEnvelopeSuccess = true
)

type AccessCustomPageListResponseEnvelope struct {
	Errors     []AccessCustomPageListResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessCustomPageListResponseEnvelopeMessages `json:"messages"`
	Result     []AccessCustomPageListResponse                 `json:"result"`
	ResultInfo AccessCustomPageListResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessCustomPageListResponseEnvelopeSuccess `json:"success"`
	JSON    accessCustomPageListResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageListResponseEnvelope]
type accessCustomPageListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessCustomPageListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCustomPageListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCustomPageListResponseEnvelopeErrors]
type accessCustomPageListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessCustomPageListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCustomPageListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessCustomPageListResponseEnvelopeMessages]
type accessCustomPageListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       accessCustomPageListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessCustomPageListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AccessCustomPageListResponseEnvelopeResultInfo]
type accessCustomPageListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageListResponseEnvelopeSuccess bool

const (
	AccessCustomPageListResponseEnvelopeSuccessTrue AccessCustomPageListResponseEnvelopeSuccess = true
)

type AccessCustomPageDeleteResponseEnvelope struct {
	Errors   []AccessCustomPageDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessCustomPageDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccessCustomPageDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessCustomPageDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    accessCustomPageDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageDeleteResponseEnvelope]
type accessCustomPageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessCustomPageDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCustomPageDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCustomPageDeleteResponseEnvelopeErrors]
type accessCustomPageDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCustomPageDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessCustomPageDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCustomPageDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessCustomPageDeleteResponseEnvelopeMessages]
type accessCustomPageDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCustomPageDeleteResponseEnvelopeSuccess bool

const (
	AccessCustomPageDeleteResponseEnvelopeSuccessTrue AccessCustomPageDeleteResponseEnvelopeSuccess = true
)
