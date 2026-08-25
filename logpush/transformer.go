// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// TransformerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransformerService] method instead.
type TransformerService struct {
	Options  []option.RequestOption
	Content  *TransformerContentService
	Versions *TransformerVersionService
}

// NewTransformerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransformerService(opts ...option.RequestOption) (r *TransformerService) {
	r = &TransformerService{}
	r.Options = opts
	r.Content = NewTransformerContentService(opts...)
	r.Versions = NewTransformerVersionService(opts...)
	return
}

// Creates a new custom log transformer for an account.
func (r *TransformerService) New(ctx context.Context, params TransformerNewParams, opts ...option.RequestOption) (res *TransformerNewResponse, err error) {
	var env TransformerNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing custom log transformer. When `code` is provided, the SQL
// query is validated and a new version is created. When `code` is omitted, only
// the name and description are updated. Omitting `description` clears the existing
// description.
func (r *TransformerService) Update(ctx context.Context, transformerID int64, params TransformerUpdateParams, opts ...option.RequestOption) (res *TransformerUpdateResponse, err error) {
	var env TransformerUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers/%v", params.AccountID, transformerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all custom log transformers for an account.
func (r *TransformerService) List(ctx context.Context, query TransformerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TransformerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all custom log transformers for an account.
func (r *TransformerService) ListAutoPaging(ctx context.Context, query TransformerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TransformerListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a custom log transformer. Returns 409 Conflict if any active logpush
// jobs reference the transformer.
func (r *TransformerService) Delete(ctx context.Context, transformerID int64, body TransformerDeleteParams, opts ...option.RequestOption) (res *TransformerDeleteResponse, err error) {
	var env TransformerDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers/%v", body.AccountID, transformerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a single custom log transformer by ID.
func (r *TransformerService) Get(ctx context.Context, transformerID int64, query TransformerGetParams, opts ...option.RequestOption) (res *TransformerGetResponse, err error) {
	var env TransformerGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers/%v", query.AccountID, transformerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Executes a SQL transformer against a single input record and returns the
// transformed output. This is a stateless endpoint — nothing is persisted.
func (r *TransformerService) Preview(ctx context.Context, params TransformerPreviewParams, opts ...option.RequestOption) (res *TransformerPreviewResponse, err error) {
	var env TransformerPreviewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers/preview", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TransformerNewResponse struct {
	// The transformer ID.
	ID int64 `json:"id"`
	// Logpush jobs that reference this transformer.
	AssociatedJobs []TransformerNewResponseAssociatedJob `json:"associated_jobs"`
	// When the transformer was created (RFC 3339).
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The dataset this transformer operates on, derived from the SQL query's FROM
	// clause. Informational only. May be absent if the dataset cannot be determined
	// from the query.
	Dataset string `json:"dataset" api:"nullable"`
	// Optional customer-provided description.
	Description string `json:"description"`
	// Customer-provided name for identification.
	Name string `json:"name"`
	// When the transformer was last modified (RFC 3339).
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      transformerNewResponseJSON `json:"-"`
}

// transformerNewResponseJSON contains the JSON metadata for the struct
// [TransformerNewResponse]
type transformerNewResponseJSON struct {
	ID             apijson.Field
	AssociatedJobs apijson.Field
	CreatedAt      apijson.Field
	Dataset        apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransformerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerNewResponseAssociatedJob struct {
	// The logpush job ID.
	ID int64 `json:"id"`
	// The logpush job destination name.
	Name string `json:"name"`
	// The zone or account tag.
	ObjectTag string `json:"object_tag"`
	// Whether the job is zone-scoped or account-scoped.
	ObjectType TransformerNewResponseAssociatedJobsObjectType `json:"object_type"`
	JSON       transformerNewResponseAssociatedJobJSON        `json:"-"`
}

// transformerNewResponseAssociatedJobJSON contains the JSON metadata for the
// struct [TransformerNewResponseAssociatedJob]
type transformerNewResponseAssociatedJobJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectTag   apijson.Field
	ObjectType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerNewResponseAssociatedJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseAssociatedJobJSON) RawJSON() string {
	return r.raw
}

// Whether the job is zone-scoped or account-scoped.
type TransformerNewResponseAssociatedJobsObjectType string

const (
	TransformerNewResponseAssociatedJobsObjectTypeZone    TransformerNewResponseAssociatedJobsObjectType = "zone"
	TransformerNewResponseAssociatedJobsObjectTypeAccount TransformerNewResponseAssociatedJobsObjectType = "account"
)

func (r TransformerNewResponseAssociatedJobsObjectType) IsKnown() bool {
	switch r {
	case TransformerNewResponseAssociatedJobsObjectTypeZone, TransformerNewResponseAssociatedJobsObjectTypeAccount:
		return true
	}
	return false
}

type TransformerUpdateResponse struct {
	// The transformer ID.
	ID int64 `json:"id"`
	// Logpush jobs that reference this transformer.
	AssociatedJobs []TransformerUpdateResponseAssociatedJob `json:"associated_jobs"`
	// When the transformer was created (RFC 3339).
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The dataset this transformer operates on, derived from the SQL query's FROM
	// clause. Informational only. May be absent if the dataset cannot be determined
	// from the query.
	Dataset string `json:"dataset" api:"nullable"`
	// Optional customer-provided description.
	Description string `json:"description"`
	// Customer-provided name for identification.
	Name string `json:"name"`
	// When the transformer was last modified (RFC 3339).
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      transformerUpdateResponseJSON `json:"-"`
}

// transformerUpdateResponseJSON contains the JSON metadata for the struct
// [TransformerUpdateResponse]
type transformerUpdateResponseJSON struct {
	ID             apijson.Field
	AssociatedJobs apijson.Field
	CreatedAt      apijson.Field
	Dataset        apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransformerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerUpdateResponseAssociatedJob struct {
	// The logpush job ID.
	ID int64 `json:"id"`
	// The logpush job destination name.
	Name string `json:"name"`
	// The zone or account tag.
	ObjectTag string `json:"object_tag"`
	// Whether the job is zone-scoped or account-scoped.
	ObjectType TransformerUpdateResponseAssociatedJobsObjectType `json:"object_type"`
	JSON       transformerUpdateResponseAssociatedJobJSON        `json:"-"`
}

// transformerUpdateResponseAssociatedJobJSON contains the JSON metadata for the
// struct [TransformerUpdateResponseAssociatedJob]
type transformerUpdateResponseAssociatedJobJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectTag   apijson.Field
	ObjectType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerUpdateResponseAssociatedJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseAssociatedJobJSON) RawJSON() string {
	return r.raw
}

// Whether the job is zone-scoped or account-scoped.
type TransformerUpdateResponseAssociatedJobsObjectType string

const (
	TransformerUpdateResponseAssociatedJobsObjectTypeZone    TransformerUpdateResponseAssociatedJobsObjectType = "zone"
	TransformerUpdateResponseAssociatedJobsObjectTypeAccount TransformerUpdateResponseAssociatedJobsObjectType = "account"
)

func (r TransformerUpdateResponseAssociatedJobsObjectType) IsKnown() bool {
	switch r {
	case TransformerUpdateResponseAssociatedJobsObjectTypeZone, TransformerUpdateResponseAssociatedJobsObjectTypeAccount:
		return true
	}
	return false
}

type TransformerListResponse struct {
	// The transformer ID.
	ID int64 `json:"id"`
	// Logpush jobs that reference this transformer.
	AssociatedJobs []TransformerListResponseAssociatedJob `json:"associated_jobs"`
	// When the transformer was created (RFC 3339).
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The dataset this transformer operates on, derived from the SQL query's FROM
	// clause. Informational only. May be absent if the dataset cannot be determined
	// from the query.
	Dataset string `json:"dataset" api:"nullable"`
	// Optional customer-provided description.
	Description string `json:"description"`
	// Customer-provided name for identification.
	Name string `json:"name"`
	// When the transformer was last modified (RFC 3339).
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      transformerListResponseJSON `json:"-"`
}

// transformerListResponseJSON contains the JSON metadata for the struct
// [TransformerListResponse]
type transformerListResponseJSON struct {
	ID             apijson.Field
	AssociatedJobs apijson.Field
	CreatedAt      apijson.Field
	Dataset        apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransformerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerListResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerListResponseAssociatedJob struct {
	// The logpush job ID.
	ID int64 `json:"id"`
	// The logpush job destination name.
	Name string `json:"name"`
	// The zone or account tag.
	ObjectTag string `json:"object_tag"`
	// Whether the job is zone-scoped or account-scoped.
	ObjectType TransformerListResponseAssociatedJobsObjectType `json:"object_type"`
	JSON       transformerListResponseAssociatedJobJSON        `json:"-"`
}

// transformerListResponseAssociatedJobJSON contains the JSON metadata for the
// struct [TransformerListResponseAssociatedJob]
type transformerListResponseAssociatedJobJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectTag   apijson.Field
	ObjectType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerListResponseAssociatedJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerListResponseAssociatedJobJSON) RawJSON() string {
	return r.raw
}

// Whether the job is zone-scoped or account-scoped.
type TransformerListResponseAssociatedJobsObjectType string

const (
	TransformerListResponseAssociatedJobsObjectTypeZone    TransformerListResponseAssociatedJobsObjectType = "zone"
	TransformerListResponseAssociatedJobsObjectTypeAccount TransformerListResponseAssociatedJobsObjectType = "account"
)

func (r TransformerListResponseAssociatedJobsObjectType) IsKnown() bool {
	switch r {
	case TransformerListResponseAssociatedJobsObjectTypeZone, TransformerListResponseAssociatedJobsObjectTypeAccount:
		return true
	}
	return false
}

type TransformerDeleteResponse struct {
	// The deleted transformer's ID.
	ID   int64                         `json:"id"`
	JSON transformerDeleteResponseJSON `json:"-"`
}

// transformerDeleteResponseJSON contains the JSON metadata for the struct
// [TransformerDeleteResponse]
type transformerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerGetResponse struct {
	// The transformer ID.
	ID int64 `json:"id"`
	// Logpush jobs that reference this transformer.
	AssociatedJobs []TransformerGetResponseAssociatedJob `json:"associated_jobs"`
	// When the transformer was created (RFC 3339).
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The dataset this transformer operates on, derived from the SQL query's FROM
	// clause. Informational only. May be absent if the dataset cannot be determined
	// from the query.
	Dataset string `json:"dataset" api:"nullable"`
	// Optional customer-provided description.
	Description string `json:"description"`
	// Customer-provided name for identification.
	Name string `json:"name"`
	// When the transformer was last modified (RFC 3339).
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      transformerGetResponseJSON `json:"-"`
}

// transformerGetResponseJSON contains the JSON metadata for the struct
// [TransformerGetResponse]
type transformerGetResponseJSON struct {
	ID             apijson.Field
	AssociatedJobs apijson.Field
	CreatedAt      apijson.Field
	Dataset        apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransformerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerGetResponseAssociatedJob struct {
	// The logpush job ID.
	ID int64 `json:"id"`
	// The logpush job destination name.
	Name string `json:"name"`
	// The zone or account tag.
	ObjectTag string `json:"object_tag"`
	// Whether the job is zone-scoped or account-scoped.
	ObjectType TransformerGetResponseAssociatedJobsObjectType `json:"object_type"`
	JSON       transformerGetResponseAssociatedJobJSON        `json:"-"`
}

// transformerGetResponseAssociatedJobJSON contains the JSON metadata for the
// struct [TransformerGetResponseAssociatedJob]
type transformerGetResponseAssociatedJobJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectTag   apijson.Field
	ObjectType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerGetResponseAssociatedJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseAssociatedJobJSON) RawJSON() string {
	return r.raw
}

// Whether the job is zone-scoped or account-scoped.
type TransformerGetResponseAssociatedJobsObjectType string

const (
	TransformerGetResponseAssociatedJobsObjectTypeZone    TransformerGetResponseAssociatedJobsObjectType = "zone"
	TransformerGetResponseAssociatedJobsObjectTypeAccount TransformerGetResponseAssociatedJobsObjectType = "account"
)

func (r TransformerGetResponseAssociatedJobsObjectType) IsKnown() bool {
	switch r {
	case TransformerGetResponseAssociatedJobsObjectTypeZone, TransformerGetResponseAssociatedJobsObjectTypeAccount:
		return true
	}
	return false
}

type TransformerPreviewResponse map[string]interface{}

type TransformerNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The SQL transformer query. Maximum 32 KB. The query must contain a FROM clause
	// referencing a valid logpush dataset.
	Code param.Field[string] `json:"code" api:"required"`
	// Customer-provided name for identification.
	Name param.Field[string] `json:"name" api:"required"`
	// Optional customer-provided description.
	Description param.Field[string] `json:"description"`
}

func (r TransformerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransformerNewResponseEnvelope struct {
	Errors   []TransformerNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransformerNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TransformerNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TransformerNewResponse                `json:"result"`
	JSON    transformerNewResponseEnvelopeJSON    `json:"-"`
}

// transformerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TransformerNewResponseEnvelope]
type transformerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformerNewResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           TransformerNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             transformerNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// transformerNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransformerNewResponseEnvelopeErrors]
type transformerNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TransformerNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    transformerNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transformerNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [TransformerNewResponseEnvelopeErrorsSource]
type transformerNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransformerNewResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           TransformerNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             transformerNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// transformerNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TransformerNewResponseEnvelopeMessages]
type transformerNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TransformerNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    transformerNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transformerNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [TransformerNewResponseEnvelopeMessagesSource]
type transformerNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransformerNewResponseEnvelopeSuccess bool

const (
	TransformerNewResponseEnvelopeSuccessTrue TransformerNewResponseEnvelopeSuccess = true
)

func (r TransformerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransformerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TransformerUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Customer-provided name for identification.
	Name param.Field[string] `json:"name" api:"required"`
	// The SQL transformer query. Maximum 32 KB. The query must contain a FROM clause
	// referencing a valid logpush dataset.
	Code param.Field[string] `json:"code"`
	// Optional customer-provided description.
	Description param.Field[string] `json:"description"`
}

func (r TransformerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransformerUpdateResponseEnvelope struct {
	Errors   []TransformerUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransformerUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TransformerUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TransformerUpdateResponse                `json:"result"`
	JSON    transformerUpdateResponseEnvelopeJSON    `json:"-"`
}

// transformerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TransformerUpdateResponseEnvelope]
type transformerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformerUpdateResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           TransformerUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             transformerUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// transformerUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransformerUpdateResponseEnvelopeErrors]
type transformerUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TransformerUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    transformerUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transformerUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [TransformerUpdateResponseEnvelopeErrorsSource]
type transformerUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransformerUpdateResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           TransformerUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             transformerUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// transformerUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TransformerUpdateResponseEnvelopeMessages]
type transformerUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TransformerUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    transformerUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transformerUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [TransformerUpdateResponseEnvelopeMessagesSource]
type transformerUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransformerUpdateResponseEnvelopeSuccess bool

const (
	TransformerUpdateResponseEnvelopeSuccessTrue TransformerUpdateResponseEnvelopeSuccess = true
)

func (r TransformerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransformerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TransformerListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type TransformerDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type TransformerDeleteResponseEnvelope struct {
	Errors   []TransformerDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransformerDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TransformerDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TransformerDeleteResponse                `json:"result"`
	JSON    transformerDeleteResponseEnvelopeJSON    `json:"-"`
}

// transformerDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TransformerDeleteResponseEnvelope]
type transformerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformerDeleteResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           TransformerDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             transformerDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// transformerDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransformerDeleteResponseEnvelopeErrors]
type transformerDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TransformerDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    transformerDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transformerDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [TransformerDeleteResponseEnvelopeErrorsSource]
type transformerDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransformerDeleteResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           TransformerDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             transformerDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// transformerDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TransformerDeleteResponseEnvelopeMessages]
type transformerDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TransformerDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    transformerDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transformerDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [TransformerDeleteResponseEnvelopeMessagesSource]
type transformerDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransformerDeleteResponseEnvelopeSuccess bool

const (
	TransformerDeleteResponseEnvelopeSuccessTrue TransformerDeleteResponseEnvelopeSuccess = true
)

func (r TransformerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransformerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TransformerGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type TransformerGetResponseEnvelope struct {
	Errors   []TransformerGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransformerGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TransformerGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TransformerGetResponse                `json:"result"`
	JSON    transformerGetResponseEnvelopeJSON    `json:"-"`
}

// transformerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TransformerGetResponseEnvelope]
type transformerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformerGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           TransformerGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             transformerGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// transformerGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransformerGetResponseEnvelopeErrors]
type transformerGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TransformerGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    transformerGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transformerGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [TransformerGetResponseEnvelopeErrorsSource]
type transformerGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransformerGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           TransformerGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             transformerGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// transformerGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TransformerGetResponseEnvelopeMessages]
type transformerGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TransformerGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    transformerGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transformerGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [TransformerGetResponseEnvelopeMessagesSource]
type transformerGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransformerGetResponseEnvelopeSuccess bool

const (
	TransformerGetResponseEnvelopeSuccessTrue TransformerGetResponseEnvelopeSuccess = true
)

func (r TransformerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransformerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TransformerPreviewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A single log record to transform (JSON object).
	Input param.Field[map[string]interface{}] `json:"input" api:"required"`
	// The SQL transformer query. Maximum 32 KB. The query must contain a FROM clause
	// referencing a valid logpush dataset.
	Sql param.Field[string] `json:"sql" api:"required"`
}

func (r TransformerPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransformerPreviewResponseEnvelope struct {
	Errors   []TransformerPreviewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransformerPreviewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TransformerPreviewResponseEnvelopeSuccess `json:"success" api:"required"`
	// The transformed log record, or null if the query filtered it out.
	Result TransformerPreviewResponse             `json:"result" api:"nullable"`
	JSON   transformerPreviewResponseEnvelopeJSON `json:"-"`
}

// transformerPreviewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TransformerPreviewResponseEnvelope]
type transformerPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerPreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformerPreviewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           TransformerPreviewResponseEnvelopeErrorsSource `json:"source"`
	JSON             transformerPreviewResponseEnvelopeErrorsJSON   `json:"-"`
}

// transformerPreviewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransformerPreviewResponseEnvelopeErrors]
type transformerPreviewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerPreviewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerPreviewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TransformerPreviewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    transformerPreviewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transformerPreviewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [TransformerPreviewResponseEnvelopeErrorsSource]
type transformerPreviewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerPreviewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerPreviewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransformerPreviewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           TransformerPreviewResponseEnvelopeMessagesSource `json:"source"`
	JSON             transformerPreviewResponseEnvelopeMessagesJSON   `json:"-"`
}

// transformerPreviewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TransformerPreviewResponseEnvelopeMessages]
type transformerPreviewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerPreviewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerPreviewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TransformerPreviewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    transformerPreviewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transformerPreviewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [TransformerPreviewResponseEnvelopeMessagesSource]
type transformerPreviewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerPreviewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerPreviewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransformerPreviewResponseEnvelopeSuccess bool

const (
	TransformerPreviewResponseEnvelopeSuccessTrue TransformerPreviewResponseEnvelopeSuccess = true
)

func (r TransformerPreviewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransformerPreviewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
