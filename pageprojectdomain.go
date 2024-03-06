// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// PageProjectDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageProjectDomainService] method
// instead.
type PageProjectDomainService struct {
	Options []option.RequestOption
}

// NewPageProjectDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageProjectDomainService(opts ...option.RequestOption) (r *PageProjectDomainService) {
	r = &PageProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *PageProjectDomainService) New(ctx context.Context, projectName string, params PageProjectDomainNewParams, opts ...option.RequestOption) (res *PageProjectDomainNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *PageProjectDomainService) List(ctx context.Context, projectName string, query PageProjectDomainListParams, opts ...option.RequestOption) (res *[]PageProjectDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Pages project's domain.
func (r *PageProjectDomainService) Delete(ctx context.Context, projectName string, domainName string, body PageProjectDomainDeleteParams, opts ...option.RequestOption) (res *PageProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *PageProjectDomainService) Edit(ctx context.Context, projectName string, domainName string, body PageProjectDomainEditParams, opts ...option.RequestOption) (res *PageProjectDomainEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single domain.
func (r *PageProjectDomainService) Get(ctx context.Context, projectName string, domainName string, query PageProjectDomainGetParams, opts ...option.RequestOption) (res *PageProjectDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", query.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PageProjectDomainNewResponseUnknown],
// [PageProjectDomainNewResponseArray] or [shared.UnionString].
type PageProjectDomainNewResponse interface {
	ImplementsPageProjectDomainNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDomainNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDomainNewResponseArray []interface{}

func (r PageProjectDomainNewResponseArray) ImplementsPageProjectDomainNewResponse() {}

type PageProjectDomainListResponse = interface{}

type PageProjectDomainDeleteResponse = interface{}

// Union satisfied by [PageProjectDomainEditResponseUnknown],
// [PageProjectDomainEditResponseArray] or [shared.UnionString].
type PageProjectDomainEditResponse interface {
	ImplementsPageProjectDomainEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDomainEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDomainEditResponseArray []interface{}

func (r PageProjectDomainEditResponseArray) ImplementsPageProjectDomainEditResponse() {}

// Union satisfied by [PageProjectDomainGetResponseUnknown],
// [PageProjectDomainGetResponseArray] or [shared.UnionString].
type PageProjectDomainGetResponse interface {
	ImplementsPageProjectDomainGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDomainGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDomainGetResponseArray []interface{}

func (r PageProjectDomainGetResponseArray) ImplementsPageProjectDomainGetResponse() {}

type PageProjectDomainNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r PageProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PageProjectDomainNewResponseEnvelope struct {
	Errors   []PageProjectDomainNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDomainNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageProjectDomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDomainNewResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDomainNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainNewResponseEnvelope]
type pageProjectDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    pageProjectDomainNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainNewResponseEnvelopeErrors]
type pageProjectDomainNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    pageProjectDomainNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainNewResponseEnvelopeMessages]
type pageProjectDomainNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageProjectDomainNewResponseEnvelopeSuccess bool

const (
	PageProjectDomainNewResponseEnvelopeSuccessTrue PageProjectDomainNewResponseEnvelopeSuccess = true
)

type PageProjectDomainListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDomainListResponseEnvelope struct {
	Errors   []PageProjectDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageProjectDomainListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    PageProjectDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageProjectDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageProjectDomainListResponseEnvelopeJSON       `json:"-"`
}

// pageProjectDomainListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainListResponseEnvelope]
type pageProjectDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    pageProjectDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainListResponseEnvelopeErrors]
type pageProjectDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    pageProjectDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainListResponseEnvelopeMessages]
type pageProjectDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageProjectDomainListResponseEnvelopeSuccess bool

const (
	PageProjectDomainListResponseEnvelopeSuccessTrue PageProjectDomainListResponseEnvelopeSuccess = true
)

type PageProjectDomainListResponseEnvelopeResultInfo struct {
	Count      interface{}                                         `json:"count"`
	Page       interface{}                                         `json:"page"`
	PerPage    interface{}                                         `json:"per_page"`
	TotalCount interface{}                                         `json:"total_count"`
	JSON       pageProjectDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageProjectDomainListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [PageProjectDomainListResponseEnvelopeResultInfo]
type pageProjectDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDomainEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDomainEditResponseEnvelope struct {
	Errors   []PageProjectDomainEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainEditResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDomainEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageProjectDomainEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDomainEditResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDomainEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainEditResponseEnvelope]
type pageProjectDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    pageProjectDomainEditResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainEditResponseEnvelopeErrors]
type pageProjectDomainEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    pageProjectDomainEditResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainEditResponseEnvelopeMessages]
type pageProjectDomainEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageProjectDomainEditResponseEnvelopeSuccess bool

const (
	PageProjectDomainEditResponseEnvelopeSuccessTrue PageProjectDomainEditResponseEnvelopeSuccess = true
)

type PageProjectDomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDomainGetResponseEnvelope struct {
	Errors   []PageProjectDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDomainGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageProjectDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDomainGetResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDomainGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainGetResponseEnvelope]
type pageProjectDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    pageProjectDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainGetResponseEnvelopeErrors]
type pageProjectDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageProjectDomainGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    pageProjectDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainGetResponseEnvelopeMessages]
type pageProjectDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageProjectDomainGetResponseEnvelopeSuccess bool

const (
	PageProjectDomainGetResponseEnvelopeSuccessTrue PageProjectDomainGetResponseEnvelopeSuccess = true
)
