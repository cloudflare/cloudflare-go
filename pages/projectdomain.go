// File generated from our OpenAPI spec by Stainless.

package pages

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ProjectDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewProjectDomainService] method
// instead.
type ProjectDomainService struct {
	Options []option.RequestOption
}

// NewProjectDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectDomainService(opts ...option.RequestOption) (r *ProjectDomainService) {
	r = &ProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *ProjectDomainService) New(ctx context.Context, projectName string, params ProjectDomainNewParams, opts ...option.RequestOption) (res *ProjectDomainNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDomainNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) List(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) (res *[]ProjectDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Pages project's domain.
func (r *ProjectDomainService) Delete(ctx context.Context, projectName string, domainName string, body ProjectDomainDeleteParams, opts ...option.RequestOption) (res *ProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *ProjectDomainService) Edit(ctx context.Context, projectName string, domainName string, body ProjectDomainEditParams, opts ...option.RequestOption) (res *ProjectDomainEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDomainEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single domain.
func (r *ProjectDomainService) Get(ctx context.Context, projectName string, domainName string, query ProjectDomainGetParams, opts ...option.RequestOption) (res *ProjectDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", query.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [pages.ProjectDomainNewResponseUnknown],
// [pages.ProjectDomainNewResponseArray] or [shared.UnionString].
type ProjectDomainNewResponse interface {
	ImplementsPagesProjectDomainNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDomainNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDomainNewResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDomainNewResponseArray []interface{}

func (r ProjectDomainNewResponseArray) ImplementsPagesProjectDomainNewResponse() {}

type ProjectDomainListResponse = interface{}

type ProjectDomainDeleteResponse = interface{}

// Union satisfied by [pages.ProjectDomainEditResponseUnknown],
// [pages.ProjectDomainEditResponseArray] or [shared.UnionString].
type ProjectDomainEditResponse interface {
	ImplementsPagesProjectDomainEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDomainEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDomainEditResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDomainEditResponseArray []interface{}

func (r ProjectDomainEditResponseArray) ImplementsPagesProjectDomainEditResponse() {}

// Union satisfied by [pages.ProjectDomainGetResponseUnknown],
// [pages.ProjectDomainGetResponseArray] or [shared.UnionString].
type ProjectDomainGetResponse interface {
	ImplementsPagesProjectDomainGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDomainGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDomainGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDomainGetResponseArray []interface{}

func (r ProjectDomainGetResponseArray) ImplementsPagesProjectDomainGetResponse() {}

type ProjectDomainNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainNewResponseEnvelope struct {
	Errors   []ProjectDomainNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainNewResponseEnvelopeJSON    `json:"-"`
}

// projectDomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainNewResponseEnvelope]
type projectDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    projectDomainNewResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseEnvelopeErrors]
type projectDomainNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    projectDomainNewResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseEnvelopeMessages]
type projectDomainNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainNewResponseEnvelopeSuccess bool

const (
	ProjectDomainNewResponseEnvelopeSuccessTrue ProjectDomainNewResponseEnvelopeSuccess = true
)

type ProjectDomainListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainListResponseEnvelope struct {
	Errors   []ProjectDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ProjectDomainListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    ProjectDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ProjectDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       projectDomainListResponseEnvelopeJSON       `json:"-"`
}

// projectDomainListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainListResponseEnvelope]
type projectDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    projectDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDomainListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainListResponseEnvelopeErrors]
type projectDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    projectDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDomainListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainListResponseEnvelopeMessages]
type projectDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainListResponseEnvelopeSuccess bool

const (
	ProjectDomainListResponseEnvelopeSuccessTrue ProjectDomainListResponseEnvelopeSuccess = true
)

type ProjectDomainListResponseEnvelopeResultInfo struct {
	Count      interface{}                                     `json:"count"`
	Page       interface{}                                     `json:"page"`
	PerPage    interface{}                                     `json:"per_page"`
	TotalCount interface{}                                     `json:"total_count"`
	JSON       projectDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// projectDomainListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [ProjectDomainListResponseEnvelopeResultInfo]
type projectDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainEditResponseEnvelope struct {
	Errors   []ProjectDomainEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainEditResponseEnvelopeJSON    `json:"-"`
}

// projectDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainEditResponseEnvelope]
type projectDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    projectDomainEditResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDomainEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseEnvelopeErrors]
type projectDomainEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    projectDomainEditResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDomainEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseEnvelopeMessages]
type projectDomainEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainEditResponseEnvelopeSuccess bool

const (
	ProjectDomainEditResponseEnvelopeSuccessTrue ProjectDomainEditResponseEnvelopeSuccess = true
)

type ProjectDomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainGetResponseEnvelope struct {
	Errors   []ProjectDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainGetResponseEnvelopeJSON    `json:"-"`
}

// projectDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainGetResponseEnvelope]
type projectDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    projectDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseEnvelopeErrors]
type projectDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    projectDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseEnvelopeMessages]
type projectDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainGetResponseEnvelopeSuccess bool

const (
	ProjectDomainGetResponseEnvelopeSuccessTrue ProjectDomainGetResponseEnvelopeSuccess = true
)
