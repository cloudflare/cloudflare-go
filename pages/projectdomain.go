// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *ProjectDomainService) New(ctx context.Context, projectName string, params ProjectDomainNewParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union, err error) {
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
func (r *ProjectDomainService) List(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProjectDomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", query.AccountID, projectName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) ListAutoPaging(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProjectDomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, projectName, query, opts...))
}

// Delete a Pages project's domain.
func (r *ProjectDomainService) Delete(ctx context.Context, projectName string, domainName string, params ProjectDomainDeleteParams, opts ...option.RequestOption) (res *ProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", params.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *ProjectDomainService) Edit(ctx context.Context, projectName string, domainName string, params ProjectDomainEditParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDomainEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", params.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single domain.
func (r *ProjectDomainService) Get(ctx context.Context, projectName string, domainName string, query ProjectDomainGetParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", query.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDomainListResponse = interface{}

type ProjectDomainDeleteResponse = interface{}

type ProjectDomainNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union `json:"result,required,nullable"`
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

// Whether the API call was successful
type ProjectDomainNewResponseEnvelopeSuccess bool

const (
	ProjectDomainNewResponseEnvelopeSuccessTrue ProjectDomainNewResponseEnvelopeSuccess = true
)

func (r ProjectDomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDomainDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainEditParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ProjectDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union `json:"result,required,nullable"`
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

// Whether the API call was successful
type ProjectDomainEditResponseEnvelopeSuccess bool

const (
	ProjectDomainEditResponseEnvelopeSuccessTrue ProjectDomainEditResponseEnvelopeSuccess = true
)

func (r ProjectDomainEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef65e3c8c1a9c4638ec25cdbbaca7165c1Union `json:"result,required,nullable"`
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

// Whether the API call was successful
type ProjectDomainGetResponseEnvelopeSuccess bool

const (
	ProjectDomainGetResponseEnvelopeSuccessTrue ProjectDomainGetResponseEnvelopeSuccess = true
)

func (r ProjectDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
