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

// ZeroTrustAccessCustomPageService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustAccessCustomPageService] method instead.
type ZeroTrustAccessCustomPageService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessCustomPageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessCustomPageService(opts ...option.RequestOption) (r *ZeroTrustAccessCustomPageService) {
	r = &ZeroTrustAccessCustomPageService{}
	r.Options = opts
	return
}

// Create a custom page
func (r *ZeroTrustAccessCustomPageService) New(ctx context.Context, identifier string, body ZeroTrustAccessCustomPageNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessCustomPageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCustomPageNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a custom page
func (r *ZeroTrustAccessCustomPageService) Update(ctx context.Context, identifier string, uuid string, body ZeroTrustAccessCustomPageUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessCustomPageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCustomPageUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List custom pages
func (r *ZeroTrustAccessCustomPageService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]ZeroTrustAccessCustomPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCustomPageListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a custom page
func (r *ZeroTrustAccessCustomPageService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZeroTrustAccessCustomPageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCustomPageDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom page and also returns its HTML.
func (r *ZeroTrustAccessCustomPageService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZeroTrustAccessCustomPageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCustomPageGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessCustomPageNewResponse struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type ZeroTrustAccessCustomPageNewResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                   `json:"uid"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessCustomPageNewResponseJSON `json:"-"`
}

// zeroTrustAccessCustomPageNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessCustomPageNewResponse]
type zeroTrustAccessCustomPageNewResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageNewResponseJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type ZeroTrustAccessCustomPageNewResponseType string

const (
	ZeroTrustAccessCustomPageNewResponseTypeIdentityDenied ZeroTrustAccessCustomPageNewResponseType = "identity_denied"
	ZeroTrustAccessCustomPageNewResponseTypeForbidden      ZeroTrustAccessCustomPageNewResponseType = "forbidden"
)

type ZeroTrustAccessCustomPageUpdateResponse struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type ZeroTrustAccessCustomPageUpdateResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                      `json:"uid"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessCustomPageUpdateResponseJSON `json:"-"`
}

// zeroTrustAccessCustomPageUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessCustomPageUpdateResponse]
type zeroTrustAccessCustomPageUpdateResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type ZeroTrustAccessCustomPageUpdateResponseType string

const (
	ZeroTrustAccessCustomPageUpdateResponseTypeIdentityDenied ZeroTrustAccessCustomPageUpdateResponseType = "identity_denied"
	ZeroTrustAccessCustomPageUpdateResponseTypeForbidden      ZeroTrustAccessCustomPageUpdateResponseType = "forbidden"
)

type ZeroTrustAccessCustomPageListResponse struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type ZeroTrustAccessCustomPageListResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                    `json:"uid"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessCustomPageListResponseJSON `json:"-"`
}

// zeroTrustAccessCustomPageListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessCustomPageListResponse]
type zeroTrustAccessCustomPageListResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageListResponseJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type ZeroTrustAccessCustomPageListResponseType string

const (
	ZeroTrustAccessCustomPageListResponseTypeIdentityDenied ZeroTrustAccessCustomPageListResponseType = "identity_denied"
	ZeroTrustAccessCustomPageListResponseTypeForbidden      ZeroTrustAccessCustomPageListResponseType = "forbidden"
)

type ZeroTrustAccessCustomPageDeleteResponse struct {
	// UUID
	ID   string                                      `json:"id"`
	JSON zeroTrustAccessCustomPageDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessCustomPageDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessCustomPageDeleteResponse]
type zeroTrustAccessCustomPageDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageGetResponse struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type ZeroTrustAccessCustomPageGetResponseType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                                   `json:"uid"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessCustomPageGetResponseJSON `json:"-"`
}

// zeroTrustAccessCustomPageGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessCustomPageGetResponse]
type zeroTrustAccessCustomPageGetResponseJSON struct {
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

func (r *ZeroTrustAccessCustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageGetResponseJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type ZeroTrustAccessCustomPageGetResponseType string

const (
	ZeroTrustAccessCustomPageGetResponseTypeIdentityDenied ZeroTrustAccessCustomPageGetResponseType = "identity_denied"
	ZeroTrustAccessCustomPageGetResponseTypeForbidden      ZeroTrustAccessCustomPageGetResponseType = "forbidden"
)

type ZeroTrustAccessCustomPageNewParams struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[ZeroTrustAccessCustomPageNewParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r ZeroTrustAccessCustomPageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type ZeroTrustAccessCustomPageNewParamsType string

const (
	ZeroTrustAccessCustomPageNewParamsTypeIdentityDenied ZeroTrustAccessCustomPageNewParamsType = "identity_denied"
	ZeroTrustAccessCustomPageNewParamsTypeForbidden      ZeroTrustAccessCustomPageNewParamsType = "forbidden"
)

type ZeroTrustAccessCustomPageNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessCustomPageNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCustomPageNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessCustomPageNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCustomPageNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCustomPageNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCustomPageNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessCustomPageNewResponseEnvelope]
type zeroTrustAccessCustomPageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageNewResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustAccessCustomPageNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCustomPageNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageNewResponseEnvelopeErrors]
type zeroTrustAccessCustomPageNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageNewResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustAccessCustomPageNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCustomPageNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageNewResponseEnvelopeMessages]
type zeroTrustAccessCustomPageNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessCustomPageNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCustomPageNewResponseEnvelopeSuccessTrue ZeroTrustAccessCustomPageNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCustomPageUpdateParams struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[ZeroTrustAccessCustomPageUpdateParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r ZeroTrustAccessCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type ZeroTrustAccessCustomPageUpdateParamsType string

const (
	ZeroTrustAccessCustomPageUpdateParamsTypeIdentityDenied ZeroTrustAccessCustomPageUpdateParamsType = "identity_denied"
	ZeroTrustAccessCustomPageUpdateParamsTypeForbidden      ZeroTrustAccessCustomPageUpdateParamsType = "forbidden"
)

type ZeroTrustAccessCustomPageUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessCustomPageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCustomPageUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessCustomPageUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCustomPageUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCustomPageUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCustomPageUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessCustomPageUpdateResponseEnvelope]
type zeroTrustAccessCustomPageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageUpdateResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessCustomPageUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCustomPageUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageUpdateResponseEnvelopeErrors]
type zeroTrustAccessCustomPageUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageUpdateResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessCustomPageUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCustomPageUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCustomPageUpdateResponseEnvelopeMessages]
type zeroTrustAccessCustomPageUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessCustomPageUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCustomPageUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessCustomPageUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCustomPageListResponseEnvelope struct {
	Errors   []ZeroTrustAccessCustomPageListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCustomPageListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessCustomPageListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessCustomPageListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessCustomPageListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessCustomPageListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessCustomPageListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessCustomPageListResponseEnvelope]
type zeroTrustAccessCustomPageListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageListResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessCustomPageListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCustomPageListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageListResponseEnvelopeErrors]
type zeroTrustAccessCustomPageListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageListResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessCustomPageListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCustomPageListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageListResponseEnvelopeMessages]
type zeroTrustAccessCustomPageListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessCustomPageListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCustomPageListResponseEnvelopeSuccessTrue ZeroTrustAccessCustomPageListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCustomPageListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       zeroTrustAccessCustomPageListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessCustomPageListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCustomPageListResponseEnvelopeResultInfo]
type zeroTrustAccessCustomPageListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessCustomPageDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCustomPageDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessCustomPageDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCustomPageDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCustomPageDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCustomPageDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessCustomPageDeleteResponseEnvelope]
type zeroTrustAccessCustomPageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageDeleteResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessCustomPageDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCustomPageDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageDeleteResponseEnvelopeErrors]
type zeroTrustAccessCustomPageDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageDeleteResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessCustomPageDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCustomPageDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCustomPageDeleteResponseEnvelopeMessages]
type zeroTrustAccessCustomPageDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessCustomPageDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCustomPageDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessCustomPageDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCustomPageGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessCustomPageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCustomPageGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessCustomPageGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCustomPageGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCustomPageGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCustomPageGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessCustomPageGetResponseEnvelope]
type zeroTrustAccessCustomPageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustAccessCustomPageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCustomPageGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageGetResponseEnvelopeErrors]
type zeroTrustAccessCustomPageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessCustomPageGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustAccessCustomPageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCustomPageGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCustomPageGetResponseEnvelopeMessages]
type zeroTrustAccessCustomPageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCustomPageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessCustomPageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessCustomPageGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCustomPageGetResponseEnvelopeSuccessTrue ZeroTrustAccessCustomPageGetResponseEnvelopeSuccess = true
)
