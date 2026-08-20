// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPDataTagCategoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDataTagCategoryService] method instead.
type DLPDataTagCategoryService struct {
	Options  []option.RequestOption
	DataTags *DLPDataTagCategoryDataTagService
}

// NewDLPDataTagCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPDataTagCategoryService(opts ...option.RequestOption) (r *DLPDataTagCategoryService) {
	r = &DLPDataTagCategoryService{}
	r.Options = opts
	r.DataTags = NewDLPDataTagCategoryDataTagService(opts...)
	return
}

// Creates a new data tag category.
func (r *DLPDataTagCategoryService) New(ctx context.Context, params DLPDataTagCategoryNewParams, opts ...option.RequestOption) (res *DLPDataTagCategoryNewResponse, err error) {
	var env DLPDataTagCategoryNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the attributes of a single data tag category.
func (r *DLPDataTagCategoryService) Update(ctx context.Context, categoryID string, params DLPDataTagCategoryUpdateParams, opts ...option.RequestOption) (res *DLPDataTagCategoryUpdateResponse, err error) {
	var env DLPDataTagCategoryUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s", params.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve all data tag categories in an account
func (r *DLPDataTagCategoryService) List(ctx context.Context, query DLPDataTagCategoryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPDataTagCategoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories", query.AccountID)
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

// Retrieve all data tag categories in an account
func (r *DLPDataTagCategoryService) ListAutoPaging(ctx context.Context, query DLPDataTagCategoryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPDataTagCategoryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a single data tag category.
func (r *DLPDataTagCategoryService) Delete(ctx context.Context, categoryID string, body DLPDataTagCategoryDeleteParams, opts ...option.RequestOption) (res *DLPDataTagCategoryDeleteResponse, err error) {
	var env DLPDataTagCategoryDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s", body.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a specific data tag category.
func (r *DLPDataTagCategoryService) Get(ctx context.Context, categoryID string, query DLPDataTagCategoryGetParams, opts ...option.RequestOption) (res *DLPDataTagCategoryGetResponse, err error) {
	var env DLPDataTagCategoryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s", query.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPDataTagCategoryNewResponse struct {
	ID          string                             `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                          `json:"created_at" api:"required" format:"date-time"`
	Name        string                             `json:"name" api:"required"`
	Tags        []DLPDataTagCategoryNewResponseTag `json:"tags" api:"required"`
	UpdatedAt   time.Time                          `json:"updated_at" api:"required" format:"date-time"`
	Description string                             `json:"description" api:"nullable"`
	TemplateID  string                             `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpDataTagCategoryNewResponseJSON  `json:"-"`
}

// dlpDataTagCategoryNewResponseJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryNewResponse]
type dlpDataTagCategoryNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryNewResponseTag struct {
	ID          string                               `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                            `json:"created_at" api:"required" format:"date-time"`
	Name        string                               `json:"name" api:"required"`
	UpdatedAt   time.Time                            `json:"updated_at" api:"required" format:"date-time"`
	Description string                               `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryNewResponseTagJSON `json:"-"`
}

// dlpDataTagCategoryNewResponseTagJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryNewResponseTag]
type dlpDataTagCategoryNewResponseTagJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseTagJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryUpdateResponse struct {
	ID          string                                `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                             `json:"created_at" api:"required" format:"date-time"`
	Name        string                                `json:"name" api:"required"`
	Tags        []DLPDataTagCategoryUpdateResponseTag `json:"tags" api:"required"`
	UpdatedAt   time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	Description string                                `json:"description" api:"nullable"`
	TemplateID  string                                `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpDataTagCategoryUpdateResponseJSON  `json:"-"`
}

// dlpDataTagCategoryUpdateResponseJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryUpdateResponse]
type dlpDataTagCategoryUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryUpdateResponseTag struct {
	ID          string                                  `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Name        string                                  `json:"name" api:"required"`
	UpdatedAt   time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description string                                  `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryUpdateResponseTagJSON `json:"-"`
}

// dlpDataTagCategoryUpdateResponseTagJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryUpdateResponseTag]
type dlpDataTagCategoryUpdateResponseTagJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseTagJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryListResponse struct {
	ID          string                              `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                           `json:"created_at" api:"required" format:"date-time"`
	Name        string                              `json:"name" api:"required"`
	Tags        []DLPDataTagCategoryListResponseTag `json:"tags" api:"required"`
	UpdatedAt   time.Time                           `json:"updated_at" api:"required" format:"date-time"`
	Description string                              `json:"description" api:"nullable"`
	TemplateID  string                              `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpDataTagCategoryListResponseJSON  `json:"-"`
}

// dlpDataTagCategoryListResponseJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryListResponse]
type dlpDataTagCategoryListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryListResponseTag struct {
	ID          string                                `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                             `json:"created_at" api:"required" format:"date-time"`
	Name        string                                `json:"name" api:"required"`
	UpdatedAt   time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	Description string                                `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryListResponseTagJSON `json:"-"`
}

// dlpDataTagCategoryListResponseTagJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryListResponseTag]
type dlpDataTagCategoryListResponseTagJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryListResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryListResponseTagJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDeleteResponse = interface{}

type DLPDataTagCategoryGetResponse struct {
	ID          string                             `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                          `json:"created_at" api:"required" format:"date-time"`
	Name        string                             `json:"name" api:"required"`
	Tags        []DLPDataTagCategoryGetResponseTag `json:"tags" api:"required"`
	UpdatedAt   time.Time                          `json:"updated_at" api:"required" format:"date-time"`
	Description string                             `json:"description" api:"nullable"`
	TemplateID  string                             `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpDataTagCategoryGetResponseJSON  `json:"-"`
}

// dlpDataTagCategoryGetResponseJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryGetResponse]
type dlpDataTagCategoryGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryGetResponseTag struct {
	ID          string                               `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                            `json:"created_at" api:"required" format:"date-time"`
	Name        string                               `json:"name" api:"required"`
	UpdatedAt   time.Time                            `json:"updated_at" api:"required" format:"date-time"`
	Description string                               `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryGetResponseTagJSON `json:"-"`
}

// dlpDataTagCategoryGetResponseTagJSON contains the JSON metadata for the struct
// [DLPDataTagCategoryGetResponseTag]
type dlpDataTagCategoryGetResponseTagJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseTagJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
	// Tags to create with the category. Mutually exclusive with `template_id`.
	Tags       param.Field[[]DLPDataTagCategoryNewParamsTag] `json:"tags"`
	TemplateID param.Field[string]                           `json:"template_id" format:"uuid"`
}

func (r DLPDataTagCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataTagCategoryNewParamsTag struct {
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r DLPDataTagCategoryNewParamsTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataTagCategoryNewResponseEnvelope struct {
	Errors   []DLPDataTagCategoryNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryNewResponse                `json:"result"`
	JSON    dlpDataTagCategoryNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryNewResponseEnvelope]
type dlpDataTagCategoryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPDataTagCategoryNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPDataTagCategoryNewResponseEnvelopeErrors]
type dlpDataTagCategoryNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpDataTagCategoryNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryNewResponseEnvelopeErrorsSource]
type dlpDataTagCategoryNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPDataTagCategoryNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDataTagCategoryNewResponseEnvelopeMessages]
type dlpDataTagCategoryNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpDataTagCategoryNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryNewResponseEnvelopeMessagesSource]
type dlpDataTagCategoryNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryNewResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryNewResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryNewResponseEnvelopeSuccessTrue DLPDataTagCategoryNewResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataTagCategoryUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
	// The desired final state of tags.
	//
	// - `None` (omitted): no tag changes.
	// - `Some([])`: delete all tags.
	// - `Some([...])`: desired final set + order.
	Tags param.Field[[]DLPDataTagCategoryUpdateParamsTag] `json:"tags"`
}

func (r DLPDataTagCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataTagCategoryUpdateParamsTag struct {
	// If `None` (omitted), a new tag will be created. Otherwise, an existing tag will
	// be updated.
	ID          param.Field[string] `json:"id" format:"uuid"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DLPDataTagCategoryUpdateParamsTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataTagCategoryUpdateResponseEnvelope struct {
	Errors   []DLPDataTagCategoryUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryUpdateResponse                `json:"result"`
	JSON    dlpDataTagCategoryUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryUpdateResponseEnvelope]
type dlpDataTagCategoryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryUpdateResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPDataTagCategoryUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryUpdateResponseEnvelopeErrors]
type dlpDataTagCategoryUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpDataTagCategoryUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryUpdateResponseEnvelopeErrorsSource]
type dlpDataTagCategoryUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryUpdateResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPDataTagCategoryUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryUpdateResponseEnvelopeMessages]
type dlpDataTagCategoryUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpDataTagCategoryUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryUpdateResponseEnvelopeMessagesSource]
type dlpDataTagCategoryUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryUpdateResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryUpdateResponseEnvelopeSuccessTrue DLPDataTagCategoryUpdateResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataTagCategoryListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataTagCategoryDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataTagCategoryDeleteResponseEnvelope struct {
	Errors   []DLPDataTagCategoryDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpDataTagCategoryDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryDeleteResponseEnvelope]
type dlpDataTagCategoryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPDataTagCategoryDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryDeleteResponseEnvelopeErrors]
type dlpDataTagCategoryDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpDataTagCategoryDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDeleteResponseEnvelopeErrorsSource]
type dlpDataTagCategoryDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPDataTagCategoryDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryDeleteResponseEnvelopeMessages]
type dlpDataTagCategoryDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpDataTagCategoryDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDeleteResponseEnvelopeMessagesSource]
type dlpDataTagCategoryDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryDeleteResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryDeleteResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryDeleteResponseEnvelopeSuccessTrue DLPDataTagCategoryDeleteResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataTagCategoryGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataTagCategoryGetResponseEnvelope struct {
	Errors   []DLPDataTagCategoryGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryGetResponse                `json:"result"`
	JSON    dlpDataTagCategoryGetResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryGetResponseEnvelope]
type dlpDataTagCategoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPDataTagCategoryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPDataTagCategoryGetResponseEnvelopeErrors]
type dlpDataTagCategoryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpDataTagCategoryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryGetResponseEnvelopeErrorsSource]
type dlpDataTagCategoryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPDataTagCategoryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDataTagCategoryGetResponseEnvelopeMessages]
type dlpDataTagCategoryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpDataTagCategoryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryGetResponseEnvelopeMessagesSource]
type dlpDataTagCategoryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryGetResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryGetResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryGetResponseEnvelopeSuccessTrue DLPDataTagCategoryGetResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
