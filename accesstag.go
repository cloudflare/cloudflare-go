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

// AccessTagService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessTagService] method instead.
type AccessTagService struct {
	Options []option.RequestOption
}

// NewAccessTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessTagService(opts ...option.RequestOption) (r *AccessTagService) {
	r = &AccessTagService{}
	r.Options = opts
	return
}

// Create a tag
func (r *AccessTagService) New(ctx context.Context, identifier string, body AccessTagNewParams, opts ...option.RequestOption) (res *AccessTagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a tag
func (r *AccessTagService) Update(ctx context.Context, identifier string, tagName string, body AccessTagUpdateParams, opts ...option.RequestOption) (res *AccessTagUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List tags
func (r *AccessTagService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessTagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a tag
func (r *AccessTagService) Delete(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *AccessTagDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a tag
func (r *AccessTagService) Get(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *AccessTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A tag
type AccessTagNewResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                    `json:"app_count"`
	CreatedAt time.Time                `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	JSON      accessTagNewResponseJSON `json:"-"`
}

// accessTagNewResponseJSON contains the JSON metadata for the struct
// [AccessTagNewResponse]
type accessTagNewResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccessTagUpdateResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                       `json:"app_count"`
	CreatedAt time.Time                   `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      accessTagUpdateResponseJSON `json:"-"`
}

// accessTagUpdateResponseJSON contains the JSON metadata for the struct
// [AccessTagUpdateResponse]
type accessTagUpdateResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccessTagListResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                     `json:"app_count"`
	CreatedAt time.Time                 `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      accessTagListResponseJSON `json:"-"`
}

// accessTagListResponseJSON contains the JSON metadata for the struct
// [AccessTagListResponse]
type accessTagListResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagDeleteResponse struct {
	// The name of the tag
	Name string                      `json:"name"`
	JSON accessTagDeleteResponseJSON `json:"-"`
}

// accessTagDeleteResponseJSON contains the JSON metadata for the struct
// [AccessTagDeleteResponse]
type accessTagDeleteResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type AccessTagGetResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                    `json:"app_count"`
	CreatedAt time.Time                `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	JSON      accessTagGetResponseJSON `json:"-"`
}

// accessTagGetResponseJSON contains the JSON metadata for the struct
// [AccessTagGetResponse]
type accessTagGetResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagNewParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r AccessTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessTagNewResponseEnvelope struct {
	Errors   []AccessTagNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagNewResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result AccessTagNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagNewResponseEnvelopeJSON    `json:"-"`
}

// accessTagNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagNewResponseEnvelope]
type accessTagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessTagNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessTagNewResponseEnvelopeErrors]
type accessTagNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessTagNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagNewResponseEnvelopeMessages]
type accessTagNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessTagNewResponseEnvelopeSuccess bool

const (
	AccessTagNewResponseEnvelopeSuccessTrue AccessTagNewResponseEnvelopeSuccess = true
)

type AccessTagUpdateParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r AccessTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessTagUpdateResponseEnvelope struct {
	Errors   []AccessTagUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result AccessTagUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessTagUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagUpdateResponseEnvelope]
type accessTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessTagUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessTagUpdateResponseEnvelopeErrors]
type accessTagUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessTagUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagUpdateResponseEnvelopeMessages]
type accessTagUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessTagUpdateResponseEnvelopeSuccess bool

const (
	AccessTagUpdateResponseEnvelopeSuccessTrue AccessTagUpdateResponseEnvelopeSuccess = true
)

type AccessTagListResponseEnvelope struct {
	Errors   []AccessTagListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessTagListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessTagListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessTagListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessTagListResponseEnvelopeJSON       `json:"-"`
}

// accessTagListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagListResponseEnvelope]
type accessTagListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessTagListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessTagListResponseEnvelopeErrors]
type accessTagListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessTagListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagListResponseEnvelopeMessages]
type accessTagListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessTagListResponseEnvelopeSuccess bool

const (
	AccessTagListResponseEnvelopeSuccessTrue AccessTagListResponseEnvelopeSuccess = true
)

type AccessTagListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       accessTagListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessTagListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessTagListResponseEnvelopeResultInfo]
type accessTagListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagDeleteResponseEnvelope struct {
	Errors   []AccessTagDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessTagDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessTagDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagDeleteResponseEnvelope]
type accessTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessTagDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessTagDeleteResponseEnvelopeErrors]
type accessTagDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessTagDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagDeleteResponseEnvelopeMessages]
type accessTagDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessTagDeleteResponseEnvelopeSuccess bool

const (
	AccessTagDeleteResponseEnvelopeSuccessTrue AccessTagDeleteResponseEnvelopeSuccess = true
)

type AccessTagGetResponseEnvelope struct {
	Errors   []AccessTagGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result AccessTagGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagGetResponseEnvelopeJSON    `json:"-"`
}

// accessTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagGetResponseEnvelope]
type accessTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessTagGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessTagGetResponseEnvelopeErrors]
type accessTagGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessTagGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagGetResponseEnvelopeMessages]
type accessTagGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessTagGetResponseEnvelopeSuccess bool

const (
	AccessTagGetResponseEnvelopeSuccessTrue AccessTagGetResponseEnvelopeSuccess = true
)
