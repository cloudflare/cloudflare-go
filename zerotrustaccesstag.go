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

// ZeroTrustAccessTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessTagService] method
// instead.
type ZeroTrustAccessTagService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessTagService(opts ...option.RequestOption) (r *ZeroTrustAccessTagService) {
	r = &ZeroTrustAccessTagService{}
	r.Options = opts
	return
}

// Create a tag
func (r *ZeroTrustAccessTagService) New(ctx context.Context, identifier string, body ZeroTrustAccessTagNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessTagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessTagNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a tag
func (r *ZeroTrustAccessTagService) Update(ctx context.Context, identifier string, tagName string, body ZeroTrustAccessTagUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessTagUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessTagUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List tags
func (r *ZeroTrustAccessTagService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]ZeroTrustAccessTagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessTagListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a tag
func (r *ZeroTrustAccessTagService) Delete(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *ZeroTrustAccessTagDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessTagDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a tag
func (r *ZeroTrustAccessTagService) Get(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *ZeroTrustAccessTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A tag
type ZeroTrustAccessTagNewResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                             `json:"app_count"`
	CreatedAt time.Time                         `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessTagNewResponseJSON `json:"-"`
}

// zeroTrustAccessTagNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessTagNewResponse]
type zeroTrustAccessTagNewResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type ZeroTrustAccessTagUpdateResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                                `json:"app_count"`
	CreatedAt time.Time                            `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessTagUpdateResponseJSON `json:"-"`
}

// zeroTrustAccessTagUpdateResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessTagUpdateResponse]
type zeroTrustAccessTagUpdateResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type ZeroTrustAccessTagListResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                              `json:"app_count"`
	CreatedAt time.Time                          `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessTagListResponseJSON `json:"-"`
}

// zeroTrustAccessTagListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessTagListResponse]
type zeroTrustAccessTagListResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagDeleteResponse struct {
	// The name of the tag
	Name string                               `json:"name"`
	JSON zeroTrustAccessTagDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessTagDeleteResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessTagDeleteResponse]
type zeroTrustAccessTagDeleteResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A tag
type ZeroTrustAccessTagGetResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                             `json:"app_count"`
	CreatedAt time.Time                         `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessTagGetResponseJSON `json:"-"`
}

// zeroTrustAccessTagGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessTagGetResponse]
type zeroTrustAccessTagGetResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagNewParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessTagNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessTagNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessTagNewResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result ZeroTrustAccessTagNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessTagNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessTagNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessTagNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessTagNewResponseEnvelope]
type zeroTrustAccessTagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zeroTrustAccessTagNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessTagNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessTagNewResponseEnvelopeErrors]
type zeroTrustAccessTagNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustAccessTagNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessTagNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustAccessTagNewResponseEnvelopeMessages]
type zeroTrustAccessTagNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessTagNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessTagNewResponseEnvelopeSuccessTrue ZeroTrustAccessTagNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessTagUpdateParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessTagUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessTagUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessTagUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result ZeroTrustAccessTagUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessTagUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessTagUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessTagUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessTagUpdateResponseEnvelope]
type zeroTrustAccessTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessTagUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessTagUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessTagUpdateResponseEnvelopeErrors]
type zeroTrustAccessTagUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessTagUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessTagUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessTagUpdateResponseEnvelopeMessages]
type zeroTrustAccessTagUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessTagUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessTagUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessTagUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessTagListResponseEnvelope struct {
	Errors   []ZeroTrustAccessTagListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessTagListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessTagListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessTagListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessTagListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessTagListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessTagListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessTagListResponseEnvelope]
type zeroTrustAccessTagListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustAccessTagListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessTagListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessTagListResponseEnvelopeErrors]
type zeroTrustAccessTagListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessTagListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessTagListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessTagListResponseEnvelopeMessages]
type zeroTrustAccessTagListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessTagListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessTagListResponseEnvelopeSuccessTrue ZeroTrustAccessTagListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessTagListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       zeroTrustAccessTagListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessTagListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustAccessTagListResponseEnvelopeResultInfo]
type zeroTrustAccessTagListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessTagDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessTagDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessTagDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessTagDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessTagDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessTagDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessTagDeleteResponseEnvelope]
type zeroTrustAccessTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessTagDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessTagDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessTagDeleteResponseEnvelopeErrors]
type zeroTrustAccessTagDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessTagDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessTagDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessTagDeleteResponseEnvelopeMessages]
type zeroTrustAccessTagDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessTagDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessTagDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessTagDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessTagGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessTagGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result ZeroTrustAccessTagGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessTagGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessTagGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessTagGetResponseEnvelope]
type zeroTrustAccessTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zeroTrustAccessTagGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessTagGetResponseEnvelopeErrors]
type zeroTrustAccessTagGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessTagGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustAccessTagGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustAccessTagGetResponseEnvelopeMessages]
type zeroTrustAccessTagGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessTagGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessTagGetResponseEnvelopeSuccessTrue ZeroTrustAccessTagGetResponseEnvelopeSuccess = true
)
